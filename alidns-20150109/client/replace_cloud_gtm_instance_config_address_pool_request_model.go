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
	// 	- zh-CN: Chinese
	//
	// 	- en-US (default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The address pools.
	AddressPools []*ReplaceCloudGtmInstanceConfigAddressPoolRequestAddressPools `json:"AddressPools,omitempty" xml:"AddressPools,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configuration ID of the access domain name. Two configuration IDs exist when the access domain name is bound to the same GTM instance but an A record and an AAAA record are configured for the access domain name. The configuration ID uniquely identifies a configuration.
	//
	// You can call the [ListCloudGtmInstanceConfigs](~~ListCloudGtmInstanceConfigs~~) operation to query the configuration ID of the access domain name.
	//
	// example:
	//
	// Config-000**11
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the GTM 3.0 instance for which you want to change address pools.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**
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
	// The ID of the address pool. This ID uniquely identifies the address pool.
	//
	// 	- If you specify this parameter, the address pools that are associated with the desired instance are removed and the instance is associated with new address pools.
	//
	// 	- If this parameter is left empty, the address pools that are associated with the desired instance are removed and no address pool is associated with the instance.
	//
	// example:
	//
	// pool-89564542105737**12
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	// The DNS request sources.
	RequestSource []*string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty" type:"Repeated"`
	// The sequence number of the new address pool. The address pool with the smallest sequence number is preferentially returned for DNS requests from any source. The sequence number specifies the priority for returning the address pool. A smaller sequence number specifies a higher priority.
	//
	// example:
	//
	// 1
	SerialNumber *int32 `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The weight value of the new address pool. You can set a different weight value for each address pool. This way, address pools are returned based on the weight values for Domain Name System (DNS) requests. A weight value must be an integer that ranges from 1 to 100.
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
