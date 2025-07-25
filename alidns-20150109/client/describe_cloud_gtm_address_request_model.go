// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeCloudGtmAddressRequest
	GetAcceptLanguage() *string
	SetAddressId(v string) *DescribeCloudGtmAddressRequest
	GetAddressId() *string
	SetClientToken(v string) *DescribeCloudGtmAddressRequest
	GetClientToken() *string
}

type DescribeCloudGtmAddressRequest struct {
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
	// addr-89518218114368**92
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribeCloudGtmAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeCloudGtmAddressRequest) GetAddressId() *string {
	return s.AddressId
}

func (s *DescribeCloudGtmAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeCloudGtmAddressRequest) SetAcceptLanguage(v string) *DescribeCloudGtmAddressRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeCloudGtmAddressRequest) SetAddressId(v string) *DescribeCloudGtmAddressRequest {
	s.AddressId = &v
	return s
}

func (s *DescribeCloudGtmAddressRequest) SetClientToken(v string) *DescribeCloudGtmAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeCloudGtmAddressRequest) Validate() error {
	return dara.Validate(s)
}
