// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResolverEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointId(v string) *DeleteResolverEndpointRequest
	GetEndpointId() *string
	SetLang(v string) *DeleteResolverEndpointRequest
	GetLang() *string
}

type DeleteResolverEndpointRequest struct {
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

func (s DeleteResolverEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResolverEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteResolverEndpointRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DeleteResolverEndpointRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteResolverEndpointRequest) SetEndpointId(v string) *DeleteResolverEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *DeleteResolverEndpointRequest) SetLang(v string) *DeleteResolverEndpointRequest {
	s.Lang = &v
	return s
}

func (s *DeleteResolverEndpointRequest) Validate() error {
	return dara.Validate(s)
}
