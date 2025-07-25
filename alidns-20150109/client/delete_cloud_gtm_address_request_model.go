// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudGtmAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteCloudGtmAddressRequest
	GetAcceptLanguage() *string
	SetAddressId(v string) *DeleteCloudGtmAddressRequest
	GetAddressId() *string
	SetClientToken(v string) *DeleteCloudGtmAddressRequest
	GetClientToken() *string
}

type DeleteCloudGtmAddressRequest struct {
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
	// The address ID. This ID uniquely identifies the address.
	//
	// This parameter is required.
	//
	// example:
	//
	// addr-895182181143688192
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can specify a custom value for this parameter, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DeleteCloudGtmAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudGtmAddressRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudGtmAddressRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteCloudGtmAddressRequest) GetAddressId() *string {
	return s.AddressId
}

func (s *DeleteCloudGtmAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteCloudGtmAddressRequest) SetAcceptLanguage(v string) *DeleteCloudGtmAddressRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteCloudGtmAddressRequest) SetAddressId(v string) *DeleteCloudGtmAddressRequest {
	s.AddressId = &v
	return s
}

func (s *DeleteCloudGtmAddressRequest) SetClientToken(v string) *DeleteCloudGtmAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCloudGtmAddressRequest) Validate() error {
	return dara.Validate(s)
}
