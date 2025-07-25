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
	// 	- zh-CN: Chinese
	//
	// 	- en-US (default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Load balancing policy among addresses in the address pool:
	//
	// - round_robin: Round-robin, for any source of DNS resolution requests, all addresses are returned. The order of all addresses is rotated each time.
	//
	// - sequence: Sequential, for any source of DNS resolution requests, the address with the smaller sequence number (the sequence number indicates the priority of address returns, with smaller numbers having higher priority) is returned. If the address with the smaller sequence number is unavailable, the next address with a smaller sequence number is returned.
	//
	// - weight: Weighted, supports setting different weight values for each address, realizing the return of addresses according to the weight ratio for resolution queries.
	//
	// - source_nearest: Source-nearest, i.e., intelligent resolution function, where GTM can return different addresses based on the source of different DNS resolution requests, achieving the effect of users accessing nearby.
	//
	// example:
	//
	// sequence
	AddressLbStrategy *string `json:"AddressLbStrategy,omitempty" xml:"AddressLbStrategy,omitempty"`
	// The ID of the address pool. This ID uniquely identifies the address pool.
	//
	// example:
	//
	// pool-89528023225442**16
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The mode used if the address with the smallest sequence number is recovered. This parameter is required only when AddressLbStrategy is set to sequence. Valid values:
	//
	// 	- preemptive: The address with the smallest sequence number is preferentially used if this address is recovered.
	//
	// 	- non_preemptive: The current address is still used even if the address with the smallest sequence number is recovered.
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
