// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmAddressPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeCloudGtmAddressPoolRequest
	GetAcceptLanguage() *string
	SetAddressPoolId(v string) *DescribeCloudGtmAddressPoolRequest
	GetAddressPoolId() *string
	SetClientToken(v string) *DescribeCloudGtmAddressPoolRequest
	GetClientToken() *string
}

type DescribeCloudGtmAddressPoolRequest struct {
	// The language of the response. Valid values:
	//
	// - zh-CN: Chinese.
	//
	// - en-US (default): English.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The unique ID of the address pool.
	//
	// example:
	//
	// pool-89564674533755****
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	// A client-generated token that ensures the idempotence of the request. Note: Make sure that the token is unique among different requests. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribeCloudGtmAddressPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressPoolRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeCloudGtmAddressPoolRequest) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *DescribeCloudGtmAddressPoolRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeCloudGtmAddressPoolRequest) SetAcceptLanguage(v string) *DescribeCloudGtmAddressPoolRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolRequest) SetAddressPoolId(v string) *DescribeCloudGtmAddressPoolRequest {
	s.AddressPoolId = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolRequest) SetClientToken(v string) *DescribeCloudGtmAddressPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolRequest) Validate() error {
	return dara.Validate(s)
}
