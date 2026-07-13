// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressPoolLbStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCloudGtmAddressPoolLbStrategyRequest
	GetAcceptLanguage() *string
	SetAddressLbStrategy(v string) *UpdateCloudGtmAddressPoolLbStrategyRequest
	GetAddressLbStrategy() *string
	SetAddressPoolId(v string) *UpdateCloudGtmAddressPoolLbStrategyRequest
	GetAddressPoolId() *string
	SetClientToken(v string) *UpdateCloudGtmAddressPoolLbStrategyRequest
	GetClientToken() *string
	SetSequenceLbStrategyMode(v string) *UpdateCloudGtmAddressPoolLbStrategyRequest
	GetSequenceLbStrategyMode() *string
}

type UpdateCloudGtmAddressPoolLbStrategyRequest struct {
	// The language of the response. Valid values:
	//
	// - zh-CN: Chinese
	//
	// - en-US (default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The load balancing policy for the addresses in the address pool.
	//
	// - round_robin: Round robin. For each DNS query, all addresses are returned in a rotating order.
	//
	// - sequence: Sequence. The address with the highest priority is returned. Priority is determined by the ordinal number of an address. A smaller ordinal number indicates a higher priority. If an address is unavailable, the address with the next highest priority is returned.
	//
	// - weight: Weight. DNS queries are resolved based on the weight that you set for each address.
	//
	// - source_nearest: Source nearest. This is an intelligent DNS resolution feature. GTM returns an address based on the source of the DNS query. This directs users to the nearest resource.
	//
	// example:
	//
	// sequence
	AddressLbStrategy *string `json:"AddressLbStrategy,omitempty" xml:"AddressLbStrategy,omitempty"`
	// The unique ID of the address pool.
	//
	// example:
	//
	// pool-89528023225442****
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	// A client-generated token that is used to ensure the idempotence of the request. The token must be unique for each request and can contain up to 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The recovery mode when the load balancing policy is \\`sequence\\`.
	//
	// - preemptive: Preemptive mode. If a higher-priority address recovers, it is used preferentially.
	//
	// - non_preemptive: Non-preemptive mode. If a higher-priority address recovers, the current address continues to be used.
	//
	// example:
	//
	// preemptive
	SequenceLbStrategyMode *string `json:"SequenceLbStrategyMode,omitempty" xml:"SequenceLbStrategyMode,omitempty"`
}

func (s UpdateCloudGtmAddressPoolLbStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressPoolLbStrategyRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressPoolLbStrategyRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCloudGtmAddressPoolLbStrategyRequest) GetAddressLbStrategy() *string {
	return s.AddressLbStrategy
}

func (s *UpdateCloudGtmAddressPoolLbStrategyRequest) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *UpdateCloudGtmAddressPoolLbStrategyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudGtmAddressPoolLbStrategyRequest) GetSequenceLbStrategyMode() *string {
	return s.SequenceLbStrategyMode
}

func (s *UpdateCloudGtmAddressPoolLbStrategyRequest) SetAcceptLanguage(v string) *UpdateCloudGtmAddressPoolLbStrategyRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolLbStrategyRequest) SetAddressLbStrategy(v string) *UpdateCloudGtmAddressPoolLbStrategyRequest {
	s.AddressLbStrategy = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolLbStrategyRequest) SetAddressPoolId(v string) *UpdateCloudGtmAddressPoolLbStrategyRequest {
	s.AddressPoolId = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolLbStrategyRequest) SetClientToken(v string) *UpdateCloudGtmAddressPoolLbStrategyRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolLbStrategyRequest) SetSequenceLbStrategyMode(v string) *UpdateCloudGtmAddressPoolLbStrategyRequest {
	s.SequenceLbStrategyMode = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolLbStrategyRequest) Validate() error {
	return dara.Validate(s)
}
