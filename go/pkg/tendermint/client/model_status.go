/*
Tendermint RPC

Tendermint supports the following RPC protocols:  * URI over HTTP * JSONRPC over HTTP * JSONRPC over websockets  ## Configuration  RPC can be configured by tuning parameters under `[rpc]` table in the `$TMHOME/config/config.toml` file or by using the `--rpc.X` command-line flags.  Default rpc listen address is `tcp://0.0.0.0:26657`. To set another address, set the `laddr` config parameter to desired value. CORS (Cross-Origin Resource Sharing) can be enabled by setting `cors_allowed_origins`, `cors_allowed_methods`, `cors_allowed_headers` config parameters.  ## Arguments  Arguments which expect strings or byte arrays may be passed as quoted strings, like `\"abc\"` or as `0x`-prefixed strings, like `0x616263`.  ## URI/HTTP  A REST like interface.      curl localhost:26657/block?height=5  ## JSONRPC/HTTP  JSONRPC requests can be POST'd to the root RPC endpoint via HTTP.      curl --header \"Content-Type: application/json\" --request POST --data '{\"method\": \"block\", \"params\": [\"5\"], \"id\": 1}' localhost:26657  ## JSONRPC/websockets  JSONRPC requests can be also made via websocket. The websocket endpoint is at `/websocket`, e.g. `localhost:26657/websocket`. Asynchronous RPC functions like event `subscribe` and `unsubscribe` are only available via websockets.  Example using https://github.com/hashrocket/ws:      ws ws://localhost:26657/websocket     > { \"jsonrpc\": \"2.0\", \"method\": \"subscribe\", \"params\": [\"tm.event='NewBlock'\"], \"id\": 1 } 

API version: Master
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Status Status Response
type Status struct {
	NodeInfo *NodeInfo `json:"node_info,omitempty"`
	SyncInfo *SyncInfo `json:"sync_info,omitempty"`
	ValidatorInfo *ValidatorInfo `json:"validator_info,omitempty"`
}

// NewStatus instantiates a new Status object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatus() *Status {
	this := Status{}
	return &this
}

// NewStatusWithDefaults instantiates a new Status object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusWithDefaults() *Status {
	this := Status{}
	return &this
}

// GetNodeInfo returns the NodeInfo field value if set, zero value otherwise.
func (o *Status) GetNodeInfo() NodeInfo {
	if o == nil || o.NodeInfo == nil {
		var ret NodeInfo
		return ret
	}
	return *o.NodeInfo
}

// GetNodeInfoOk returns a tuple with the NodeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Status) GetNodeInfoOk() (*NodeInfo, bool) {
	if o == nil || o.NodeInfo == nil {
		return nil, false
	}
	return o.NodeInfo, true
}

// HasNodeInfo returns a boolean if a field has been set.
func (o *Status) HasNodeInfo() bool {
	if o != nil && o.NodeInfo != nil {
		return true
	}

	return false
}

// SetNodeInfo gets a reference to the given NodeInfo and assigns it to the NodeInfo field.
func (o *Status) SetNodeInfo(v NodeInfo) {
	o.NodeInfo = &v
}

// GetSyncInfo returns the SyncInfo field value if set, zero value otherwise.
func (o *Status) GetSyncInfo() SyncInfo {
	if o == nil || o.SyncInfo == nil {
		var ret SyncInfo
		return ret
	}
	return *o.SyncInfo
}

// GetSyncInfoOk returns a tuple with the SyncInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Status) GetSyncInfoOk() (*SyncInfo, bool) {
	if o == nil || o.SyncInfo == nil {
		return nil, false
	}
	return o.SyncInfo, true
}

// HasSyncInfo returns a boolean if a field has been set.
func (o *Status) HasSyncInfo() bool {
	if o != nil && o.SyncInfo != nil {
		return true
	}

	return false
}

// SetSyncInfo gets a reference to the given SyncInfo and assigns it to the SyncInfo field.
func (o *Status) SetSyncInfo(v SyncInfo) {
	o.SyncInfo = &v
}

// GetValidatorInfo returns the ValidatorInfo field value if set, zero value otherwise.
func (o *Status) GetValidatorInfo() ValidatorInfo {
	if o == nil || o.ValidatorInfo == nil {
		var ret ValidatorInfo
		return ret
	}
	return *o.ValidatorInfo
}

// GetValidatorInfoOk returns a tuple with the ValidatorInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Status) GetValidatorInfoOk() (*ValidatorInfo, bool) {
	if o == nil || o.ValidatorInfo == nil {
		return nil, false
	}
	return o.ValidatorInfo, true
}

// HasValidatorInfo returns a boolean if a field has been set.
func (o *Status) HasValidatorInfo() bool {
	if o != nil && o.ValidatorInfo != nil {
		return true
	}

	return false
}

// SetValidatorInfo gets a reference to the given ValidatorInfo and assigns it to the ValidatorInfo field.
func (o *Status) SetValidatorInfo(v ValidatorInfo) {
	o.ValidatorInfo = &v
}

func (o Status) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NodeInfo != nil {
		toSerialize["node_info"] = o.NodeInfo
	}
	if o.SyncInfo != nil {
		toSerialize["sync_info"] = o.SyncInfo
	}
	if o.ValidatorInfo != nil {
		toSerialize["validator_info"] = o.ValidatorInfo
	}
	return json.Marshal(toSerialize)
}

type NullableStatus struct {
	value *Status
	isSet bool
}

func (v NullableStatus) Get() *Status {
	return v.value
}

func (v *NullableStatus) Set(val *Status) {
	v.value = val
	v.isSet = true
}

func (v NullableStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatus(val *Status) *NullableStatus {
	return &NullableStatus{value: val, isSet: true}
}

func (v NullableStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

