// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmAddressPoolReferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeCloudGtmAddressPoolReferenceRequest
	GetAcceptLanguage() *string
	SetAddressPoolId(v string) *DescribeCloudGtmAddressPoolReferenceRequest
	GetAddressPoolId() *string
	SetClientToken(v string) *DescribeCloudGtmAddressPoolReferenceRequest
	GetClientToken() *string
}

type DescribeCloudGtmAddressPoolReferenceRequest struct {
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
}

func (s DescribeCloudGtmAddressPoolReferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressPoolReferenceRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressPoolReferenceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeCloudGtmAddressPoolReferenceRequest) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *DescribeCloudGtmAddressPoolReferenceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeCloudGtmAddressPoolReferenceRequest) SetAcceptLanguage(v string) *DescribeCloudGtmAddressPoolReferenceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceRequest) SetAddressPoolId(v string) *DescribeCloudGtmAddressPoolReferenceRequest {
	s.AddressPoolId = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceRequest) SetClientToken(v string) *DescribeCloudGtmAddressPoolReferenceRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceRequest) Validate() error {
	return dara.Validate(s)
}
