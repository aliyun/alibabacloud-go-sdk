// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmInstanceConfigLbStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCloudGtmInstanceConfigLbStrategyRequest
	GetAcceptLanguage() *string
	SetAddressPoolLbStrategy(v string) *UpdateCloudGtmInstanceConfigLbStrategyRequest
	GetAddressPoolLbStrategy() *string
	SetClientToken(v string) *UpdateCloudGtmInstanceConfigLbStrategyRequest
	GetClientToken() *string
	SetConfigId(v string) *UpdateCloudGtmInstanceConfigLbStrategyRequest
	GetConfigId() *string
	SetInstanceId(v string) *UpdateCloudGtmInstanceConfigLbStrategyRequest
	GetInstanceId() *string
	SetSequenceLbStrategyMode(v string) *UpdateCloudGtmInstanceConfigLbStrategyRequest
	GetSequenceLbStrategyMode() *string
}

type UpdateCloudGtmInstanceConfigLbStrategyRequest struct {
	// The language in which the returned results are displayed. Valid values:
	//
	// 	- **zh-CN**: Chinese
	//
	// 	- **en-US*	- (default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The new policy for load balancing between address pools. Valid values:
	//
	// 	- round_robin: All address pools are returned for Domain Name System (DNS) requests from any source. All address pools are sorted in round-robin mode each time they are returned.
	//
	// 	- sequence: The address pool with the smallest sequence number is preferentially returned for DNS requests from any source. The sequence number indicates the priority for returning the address pool. A smaller sequence number indicates a higher priority. If the address pool with the smallest sequence number is unavailable, the address pool with the second smallest sequence number is returned.
	//
	// 	- weight: You can set a different weight value for each address pool. This way, address pools are returned based on the weight values.
	//
	// 	- source_nearest: GTM returns different address pools based on the sources of DNS requests. This way, users can access nearby address pools.
	//
	// example:
	//
	// sequence
	AddressPoolLbStrategy *string `json:"AddressPoolLbStrategy,omitempty" xml:"AddressPoolLbStrategy,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configuration ID of the access domain name. Two configuration IDs exist when the access domain name is bound to the same GTM instance but an A record and an AAAA record are configured for the access domain name. The configuration ID uniquely identifies a configuration.
	//
	// You can call the [ListCloudGtmInstanceConfigs](~~ListCloudGtmInstanceConfigs~~) operation to query the configuration ID of the desired access domain name.
	//
	// example:
	//
	// Config-000**11
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the GTM 3.0 instance for which you want to modify the load balancing policy.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The mode used if the address pool with the smallest sequence number is recovered. This parameter is required when AddressPoolLbStrategy is set to sequence. Valid values:
	//
	// 	- preemptive: The address pool with the smallest sequence number is preferentially used if this address pool is recovered.
	//
	// 	- non_preemptive: The current address pool is still used even if the address pool with the smallest sequence number is recovered.
	//
	// example:
	//
	// preemptive
	SequenceLbStrategyMode *string `json:"SequenceLbStrategyMode,omitempty" xml:"SequenceLbStrategyMode,omitempty"`
}

func (s UpdateCloudGtmInstanceConfigLbStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmInstanceConfigLbStrategyRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyRequest) GetAddressPoolLbStrategy() *string {
	return s.AddressPoolLbStrategy
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyRequest) GetSequenceLbStrategyMode() *string {
	return s.SequenceLbStrategyMode
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyRequest) SetAcceptLanguage(v string) *UpdateCloudGtmInstanceConfigLbStrategyRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyRequest) SetAddressPoolLbStrategy(v string) *UpdateCloudGtmInstanceConfigLbStrategyRequest {
	s.AddressPoolLbStrategy = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyRequest) SetClientToken(v string) *UpdateCloudGtmInstanceConfigLbStrategyRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyRequest) SetConfigId(v string) *UpdateCloudGtmInstanceConfigLbStrategyRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyRequest) SetInstanceId(v string) *UpdateCloudGtmInstanceConfigLbStrategyRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyRequest) SetSequenceLbStrategyMode(v string) *UpdateCloudGtmInstanceConfigLbStrategyRequest {
	s.SequenceLbStrategyMode = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyRequest) Validate() error {
	return dara.Validate(s)
}
