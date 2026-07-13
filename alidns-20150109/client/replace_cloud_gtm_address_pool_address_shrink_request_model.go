// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceCloudGtmAddressPoolAddressShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReplaceCloudGtmAddressPoolAddressShrinkRequest
	GetAcceptLanguage() *string
	SetAddressPoolId(v string) *ReplaceCloudGtmAddressPoolAddressShrinkRequest
	GetAddressPoolId() *string
	SetAddressesShrink(v string) *ReplaceCloudGtmAddressPoolAddressShrinkRequest
	GetAddressesShrink() *string
	SetClientToken(v string) *ReplaceCloudGtmAddressPoolAddressShrinkRequest
	GetClientToken() *string
}

type ReplaceCloudGtmAddressPoolAddressShrinkRequest struct {
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
	AddressesShrink *string `json:"Addresses,omitempty" xml:"Addresses,omitempty"`
	// The client token that is used to ensure the idempotence of the request. Ensure the client token is unique for each request. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ReplaceCloudGtmAddressPoolAddressShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceCloudGtmAddressPoolAddressShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReplaceCloudGtmAddressPoolAddressShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReplaceCloudGtmAddressPoolAddressShrinkRequest) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *ReplaceCloudGtmAddressPoolAddressShrinkRequest) GetAddressesShrink() *string {
	return s.AddressesShrink
}

func (s *ReplaceCloudGtmAddressPoolAddressShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ReplaceCloudGtmAddressPoolAddressShrinkRequest) SetAcceptLanguage(v string) *ReplaceCloudGtmAddressPoolAddressShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReplaceCloudGtmAddressPoolAddressShrinkRequest) SetAddressPoolId(v string) *ReplaceCloudGtmAddressPoolAddressShrinkRequest {
	s.AddressPoolId = &v
	return s
}

func (s *ReplaceCloudGtmAddressPoolAddressShrinkRequest) SetAddressesShrink(v string) *ReplaceCloudGtmAddressPoolAddressShrinkRequest {
	s.AddressesShrink = &v
	return s
}

func (s *ReplaceCloudGtmAddressPoolAddressShrinkRequest) SetClientToken(v string) *ReplaceCloudGtmAddressPoolAddressShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ReplaceCloudGtmAddressPoolAddressShrinkRequest) Validate() error {
	return dara.Validate(s)
}
