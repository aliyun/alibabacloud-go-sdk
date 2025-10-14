// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceCloudGtmAddressPoolAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReplaceCloudGtmAddressPoolAddressRequest
	GetAcceptLanguage() *string
	SetAddressPoolId(v string) *ReplaceCloudGtmAddressPoolAddressRequest
	GetAddressPoolId() *string
	SetAddresses(v []*ReplaceCloudGtmAddressPoolAddressRequestAddresses) *ReplaceCloudGtmAddressPoolAddressRequest
	GetAddresses() []*ReplaceCloudGtmAddressPoolAddressRequestAddresses
	SetClientToken(v string) *ReplaceCloudGtmAddressPoolAddressRequest
	GetClientToken() *string
}

type ReplaceCloudGtmAddressPoolAddressRequest struct {
	// The language of the response. Valid values:
	//
	// 	- **zh-CN**: Chinese
	//
	// 	- **en-US (default)**: English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the address pool for which you want to replace addresses. This ID uniquely identifies the address pool.
	//
	// example:
	//
	// pool-89618921167339**24
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	// The addresses.
	Addresses []*ReplaceCloudGtmAddressPoolAddressRequestAddresses `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ReplaceCloudGtmAddressPoolAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceCloudGtmAddressPoolAddressRequest) GoString() string {
	return s.String()
}

func (s *ReplaceCloudGtmAddressPoolAddressRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReplaceCloudGtmAddressPoolAddressRequest) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *ReplaceCloudGtmAddressPoolAddressRequest) GetAddresses() []*ReplaceCloudGtmAddressPoolAddressRequestAddresses {
	return s.Addresses
}

func (s *ReplaceCloudGtmAddressPoolAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ReplaceCloudGtmAddressPoolAddressRequest) SetAcceptLanguage(v string) *ReplaceCloudGtmAddressPoolAddressRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReplaceCloudGtmAddressPoolAddressRequest) SetAddressPoolId(v string) *ReplaceCloudGtmAddressPoolAddressRequest {
	s.AddressPoolId = &v
	return s
}

func (s *ReplaceCloudGtmAddressPoolAddressRequest) SetAddresses(v []*ReplaceCloudGtmAddressPoolAddressRequestAddresses) *ReplaceCloudGtmAddressPoolAddressRequest {
	s.Addresses = v
	return s
}

func (s *ReplaceCloudGtmAddressPoolAddressRequest) SetClientToken(v string) *ReplaceCloudGtmAddressPoolAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *ReplaceCloudGtmAddressPoolAddressRequest) Validate() error {
	if s.Addresses != nil {
		for _, item := range s.Addresses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReplaceCloudGtmAddressPoolAddressRequestAddresses struct {
	// The ID of the new address. This ID uniquely identifies the address.
	//
	// 	- If you specify this parameter, the original addresses in the address pool will be deleted and replaced with new addresses.
	//
	// 	- If you do not specify this parameter, all addresses in the address pool will be deleted and the address pool will be left empty.
	//
	// example:
	//
	// addr-89636516932803**44
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// The DNS request sources.
	RequestSource []*string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty" type:"Repeated"`
	// The sequence number that specifies the priority for returning the new address. A smaller sequence number specifies a higher priority. This setting takes effect for new addresses.
	//
	// example:
	//
	// 1
	SerialNumber *int32 `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The weight value of the new address. You can set a different weight value for each address. This way, addresses are returned based on the weight values for Domain Name System (DNS) requests. A weight value must be an integer that ranges from 1 to 100. This setting takes effect for new addresses.
	//
	// example:
	//
	// 1
	WeightValue *int32 `json:"WeightValue,omitempty" xml:"WeightValue,omitempty"`
}

func (s ReplaceCloudGtmAddressPoolAddressRequestAddresses) String() string {
	return dara.Prettify(s)
}

func (s ReplaceCloudGtmAddressPoolAddressRequestAddresses) GoString() string {
	return s.String()
}

func (s *ReplaceCloudGtmAddressPoolAddressRequestAddresses) GetAddressId() *string {
	return s.AddressId
}

func (s *ReplaceCloudGtmAddressPoolAddressRequestAddresses) GetRequestSource() []*string {
	return s.RequestSource
}

func (s *ReplaceCloudGtmAddressPoolAddressRequestAddresses) GetSerialNumber() *int32 {
	return s.SerialNumber
}

func (s *ReplaceCloudGtmAddressPoolAddressRequestAddresses) GetWeightValue() *int32 {
	return s.WeightValue
}

func (s *ReplaceCloudGtmAddressPoolAddressRequestAddresses) SetAddressId(v string) *ReplaceCloudGtmAddressPoolAddressRequestAddresses {
	s.AddressId = &v
	return s
}

func (s *ReplaceCloudGtmAddressPoolAddressRequestAddresses) SetRequestSource(v []*string) *ReplaceCloudGtmAddressPoolAddressRequestAddresses {
	s.RequestSource = v
	return s
}

func (s *ReplaceCloudGtmAddressPoolAddressRequestAddresses) SetSerialNumber(v int32) *ReplaceCloudGtmAddressPoolAddressRequestAddresses {
	s.SerialNumber = &v
	return s
}

func (s *ReplaceCloudGtmAddressPoolAddressRequestAddresses) SetWeightValue(v int32) *ReplaceCloudGtmAddressPoolAddressRequestAddresses {
	s.WeightValue = &v
	return s
}

func (s *ReplaceCloudGtmAddressPoolAddressRequestAddresses) Validate() error {
	return dara.Validate(s)
}
