// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResolverEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointId(v string) *DescribeResolverEndpointRequest
	GetEndpointId() *string
	SetLang(v string) *DescribeResolverEndpointRequest
	GetLang() *string
}

type DescribeResolverEndpointRequest struct {
	// The endpoint ID. This ID uniquely identifies the endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// hr****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeResolverEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverEndpointRequest) GoString() string {
	return s.String()
}

func (s *DescribeResolverEndpointRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DescribeResolverEndpointRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeResolverEndpointRequest) SetEndpointId(v string) *DescribeResolverEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *DescribeResolverEndpointRequest) SetLang(v string) *DescribeResolverEndpointRequest {
	s.Lang = &v
	return s
}

func (s *DescribeResolverEndpointRequest) Validate() error {
	return dara.Validate(s)
}
