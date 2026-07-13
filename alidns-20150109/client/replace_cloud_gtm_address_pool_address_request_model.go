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
	// The response language. Valid values:
	//
	// - **zh-CN**: Chinese
	//
	// - **en-US*	- (Default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The unique ID of the address pool to update.
	//
	// example:
	//
	// pool-89618921167339****
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	// The list of addresses.
	Addresses []*ReplaceCloudGtmAddressPoolAddressRequestAddresses `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. Ensure the client token is unique for each request. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
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
	// The unique ID of the address.
	//
	// - If you specify this parameter, all existing addresses in the address pool are deleted and replaced with the specified addresses.
	//
	// - If you leave this parameter empty, all existing addresses in the address pool are deleted.
	//
	// example:
	//
	// addr-89636516932803****
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// The list of DNS request sources.
	RequestSource []*string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty" type:"Repeated"`
	// The serial number, which determines the priority of the address. A smaller number indicates a higher priority. This setting applies to the updated addresses.
	//
	// example:
	//
	// 1
	SerialNumber *int32 `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The weight of the address. Valid values: 1 to 100. You can set a different weight for each address. DNS queries are then resolved based on the weight ratio. This setting applies to the updated addresses.
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
