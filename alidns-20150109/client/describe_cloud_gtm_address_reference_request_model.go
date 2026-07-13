// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmAddressReferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeCloudGtmAddressReferenceRequest
	GetAcceptLanguage() *string
	SetAddressId(v string) *DescribeCloudGtmAddressReferenceRequest
	GetAddressId() *string
	SetClientToken(v string) *DescribeCloudGtmAddressReferenceRequest
	GetClientToken() *string
}

type DescribeCloudGtmAddressReferenceRequest struct {
	// The language of the response. Valid values:
	//
	// - zh-CN: Chinese
	//
	// - en-US: English (default)
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The unique ID of the address.
	//
	// example:
	//
	// addr-89518218114368****
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. Make sure that the client generates a unique token for each request. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribeCloudGtmAddressReferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressReferenceRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressReferenceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeCloudGtmAddressReferenceRequest) GetAddressId() *string {
	return s.AddressId
}

func (s *DescribeCloudGtmAddressReferenceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeCloudGtmAddressReferenceRequest) SetAcceptLanguage(v string) *DescribeCloudGtmAddressReferenceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceRequest) SetAddressId(v string) *DescribeCloudGtmAddressReferenceRequest {
	s.AddressId = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceRequest) SetClientToken(v string) *DescribeCloudGtmAddressReferenceRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceRequest) Validate() error {
	return dara.Validate(s)
}
