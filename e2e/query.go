package e2e

import (
	"encoding/json"
	"fmt"
	"net/http"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	wasmstoragetypes "github.com/sedaprotocol/seda-chain/x/wasm-storage/types"
)

func queryTx(endpoint, txHash string) error {
	resp, err := http.Get(fmt.Sprintf("%s/cosmos/tx/v1beta1/txs/%s", endpoint, txHash))
	if err != nil {
		return fmt.Errorf("failed to execute HTTP request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("tx query returned non-200 status: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	txResp := result["tx_response"].(map[string]interface{})
	if v := txResp["code"]; v.(float64) != 0 {
		return fmt.Errorf("tx %s failed with status code %v", txHash, v)
	}

	return nil
}

func queryGovProposal(endpoint string, proposalID int) (govtypes.QueryProposalResponse, error) {
	var govProposalResp govtypes.QueryProposalResponse

	path := fmt.Sprintf("%s/cosmos/gov/v1/proposals/%d", endpoint, proposalID)

	body, err := httpGet(path)
	if err != nil {
		return govProposalResp, fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	if err := cdc.UnmarshalJSON(body, &govProposalResp); err != nil {
		return govProposalResp, err
	}
	return govProposalResp, nil
}

func queryDataRequestWasm(endpoint, drHash string) (wasmstoragetypes.QueryDataRequestWasmResponse, error) {
	var res wasmstoragetypes.QueryDataRequestWasmResponse

	body, err := httpGet(fmt.Sprintf("%s/seda-chain/wasm-storage/data_request_wasm/%s", endpoint, drHash))
	if err != nil {
		return res, err
	}

	if err = cdc.UnmarshalJSON(body, &res); err != nil {
		return res, err
	}
	return res, nil
}

func queryOverlayWasm(endpoint, hash string) (wasmstoragetypes.QueryOverlayWasmResponse, error) {
	var res wasmstoragetypes.QueryOverlayWasmResponse

	body, err := httpGet(fmt.Sprintf("%s/seda-chain/wasm-storage/overlay_wasm/%s", endpoint, hash))
	if err != nil {
		return res, err
	}

	if err = cdc.UnmarshalJSON(body, &res); err != nil {
		return res, err
	}
	return res, nil
}

func queryDataRequestWasms(endpoint string) (wasmstoragetypes.QueryDataRequestWasmsResponse, error) {
	var res wasmstoragetypes.QueryDataRequestWasmsResponse

	body, err := httpGet(fmt.Sprintf("%s/seda-chain/wasm-storage/data_request_wasms", endpoint))
	if err != nil {
		return res, err
	}

	if err = cdc.UnmarshalJSON(body, &res); err != nil {
		return res, err
	}
	return res, nil
}

func queryOverlayWasms(endpoint string) (wasmstoragetypes.QueryOverlayWasmsResponse, error) {
	var res wasmstoragetypes.QueryOverlayWasmsResponse

	body, err := httpGet(fmt.Sprintf("%s/seda-chain/wasm-storage/overlay_wasms", endpoint))
	if err != nil {
		return res, err
	}

	if err = cdc.UnmarshalJSON(body, &res); err != nil {
		return res, err
	}
	return res, nil
}

func queryProxyContractRegistry(endpoint string) (wasmstoragetypes.QueryProxyContractRegistryResponse, error) {
	var res wasmstoragetypes.QueryProxyContractRegistryResponse

	body, err := httpGet(fmt.Sprintf("%s/seda-chain/wasm-storage/proxy_contract_registry", endpoint))
	if err != nil {
		return res, err
	}

	if err = cdc.UnmarshalJSON(body, &res); err != nil {
		return res, err
	}
	return res, nil
}
