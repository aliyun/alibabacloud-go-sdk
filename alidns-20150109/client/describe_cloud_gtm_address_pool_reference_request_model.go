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
	// - zh-CN: Chinese.
	//
	// - en-US: English. This is the default value.
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The unique ID of the address pool.
	//
	// example:
	//
	// pool-89528023225442****
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. Generate a token to make sure that it is unique among different requests. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
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
