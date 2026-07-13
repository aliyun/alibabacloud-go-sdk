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
	// The language of the response. Valid values:
	//
	// - **zh-CN**: Chinese.
	//
	// - **en-US**: English. This is the default value.
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The load balancing policy for the address pools. Valid values:
	//
	// - round_robin: Returns all address pools for any DNS request. The address pools are rotated for each request.
	//
	// - sequence: Returns the address pool with the smallest ordinal number. The smaller the ordinal number, the higher the priority. If the primary address pool is unavailable, the next address pool in the sequence is used.
	//
	// - weight: Distributes DNS requests to address pools based on their configured weights.
	//
	// - source_nearest: Returns an address pool based on the proximity of the DNS request source. This implements intelligent DNS resolution and directs users to the nearest access point.
	//
	// example:
	//
	// sequence
	AddressPoolLbStrategy *string `json:"AddressPoolLbStrategy,omitempty" xml:"AddressPoolLbStrategy,omitempty"`
	// A client-generated token that is used to ensure the idempotence of the request. The token must be unique among different requests. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the instance configuration. A GTM instance can have multiple configurations for the same domain name, such as one for A records and another for AAAA records. The ConfigId uniquely identifies the configuration that you want to modify.
	//
	// For more information, see [ListCloudGtmInstanceConfigs](https://help.aliyun.com/document_detail/2797349.html).
	//
	// example:
	//
	// Config-000****
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the Global Traffic Manager (GTM) 3.0 instance.
	//
	// example:
	//
	// gtm-cn-wwo3a3h***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The recovery mode for a primary address pool when the load balancing policy is set to sequence. Valid values:
	//
	// - preemptive: The system switches back to the primary address pool as soon as it recovers.
	//
	// - non_preemptive: The system continues to use the current address pool even after the primary address pool recovers.
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
