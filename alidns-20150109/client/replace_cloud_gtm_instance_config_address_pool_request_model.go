// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceCloudGtmInstanceConfigAddressPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReplaceCloudGtmInstanceConfigAddressPoolRequest
	GetAcceptLanguage() *string
	SetAddressPools(v []*ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools) *ReplaceCloudGtmInstanceConfigAddressPoolRequest
	GetAddressPools() []*ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools
	SetClientToken(v string) *ReplaceCloudGtmInstanceConfigAddressPoolRequest
	GetClientToken() *string
	SetConfigId(v string) *ReplaceCloudGtmInstanceConfigAddressPoolRequest
	GetConfigId() *string
	SetInstanceId(v string) *ReplaceCloudGtmInstanceConfigAddressPoolRequest
	GetInstanceId() *string
}

type ReplaceCloudGtmInstanceConfigAddressPoolRequest struct {
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
	// A list of address pools.
	AddressPools []*ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools `json:"AddressPools,omitempty" xml:"AddressPools,omitempty" type:"Repeated"`
	// A client-generated token that you use to ensure the idempotence of the request. Make sure that the token is unique among different requests. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the instance configuration. For the same access domain name and GTM instance, you can configure both A and AAAA records. In this case, the GTM instance has two instance configurations. The ConfigId parameter uniquely identifies an instance configuration.
	//
	// Call the [ListCloudGtmInstanceConfigs](https://help.aliyun.com/document_detail/2797349.html) operation to query the ConfigId of the instance configuration.
	//
	// example:
	//
	// Config-000****
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the GTM 3.0 instance for which you want to replace address pools.
	//
	// example:
	//
	// gtm-cn-wwo3a3hb****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReplaceCloudGtmInstanceConfigAddressPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceCloudGtmInstanceConfigAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolRequest) GetAddressPools() []*ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools {
	return s.AddressPools
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolRequest) SetAcceptLanguage(v string) *ReplaceCloudGtmInstanceConfigAddressPoolRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolRequest) SetAddressPools(v []*ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools) *ReplaceCloudGtmInstanceConfigAddressPoolRequest {
	s.AddressPools = v
	return s
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolRequest) SetClientToken(v string) *ReplaceCloudGtmInstanceConfigAddressPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolRequest) SetConfigId(v string) *ReplaceCloudGtmInstanceConfigAddressPoolRequest {
	s.ConfigId = &v
	return s
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolRequest) SetInstanceId(v string) *ReplaceCloudGtmInstanceConfigAddressPoolRequest {
	s.InstanceId = &v
	return s
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolRequest) Validate() error {
	if s.AddressPools != nil {
		for _, item := range s.AddressPools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools struct {
	// The unique ID of the address pool.
	//
	// - If you specify this parameter, the existing address pools associated with the target instance are deleted and replaced with the address pools that you specify.
	//
	// - If you leave this parameter empty, all address pools associated with the target instance are deleted.
	//
	// example:
	//
	// pool-89564542105737****
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	// A list of request sources.
	RequestSource []*string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty" type:"Repeated"`
	// The ordinal number. For DNS requests from any source, address pools with smaller ordinal numbers are returned first. A smaller ordinal number indicates a higher priority. This parameter takes effect for the updated address pools.
	//
	// example:
	//
	// 1
	SerialNumber *int32 `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The weight of the address pool. Valid values are integers from 1 to 100. You can set a different weight for each address pool. DNS queries are resolved based on the specified weights. This parameter takes effect for the updated address pools.
	//
	// example:
	//
	// 1
	WeightValue *int32 `json:"WeightValue,omitempty" xml:"WeightValue,omitempty"`
}

func (s ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools) String() string {
	return dara.Prettify(s)
}

func (s ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools) GoString() string {
	return s.String()
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools) GetRequestSource() []*string {
	return s.RequestSource
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools) GetSerialNumber() *int32 {
	return s.SerialNumber
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools) GetWeightValue() *int32 {
	return s.WeightValue
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools) SetAddressPoolId(v string) *ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools {
	s.AddressPoolId = &v
	return s
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools) SetRequestSource(v []*string) *ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools {
	s.RequestSource = v
	return s
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools) SetSerialNumber(v int32) *ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools {
	s.SerialNumber = &v
	return s
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools) SetWeightValue(v int32) *ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools {
	s.WeightValue = &v
	return s
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools) Validate() error {
	return dara.Validate(s)
}
