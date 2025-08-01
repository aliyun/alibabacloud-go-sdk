// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteCORSShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayRouteCORSShrinkRequest
	GetAcceptLanguage() *string
	SetCorsJSONShrink(v string) *UpdateGatewayRouteCORSShrinkRequest
	GetCorsJSONShrink() *string
	SetGatewayId(v int64) *UpdateGatewayRouteCORSShrinkRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdateGatewayRouteCORSShrinkRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *UpdateGatewayRouteCORSShrinkRequest
	GetId() *int64
}

type UpdateGatewayRouteCORSShrinkRequest struct {
	// The language of the response. In compliance with [RFC 7231](https://tools.ietf.org/html/rfc7231), the backend service must return a response based on the language used by the user.
	//
	// 	- No default value.
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The information about the CORS policy.
	CorsJSONShrink *string `json:"CorsJSON,omitempty" xml:"CorsJSON,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// 85
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-f70a6ddf2f0941a2bb997b2d16028f37
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the associated record.
	//
	// example:
	//
	// 55
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateGatewayRouteCORSShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteCORSShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteCORSShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayRouteCORSShrinkRequest) GetCorsJSONShrink() *string {
	return s.CorsJSONShrink
}

func (s *UpdateGatewayRouteCORSShrinkRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayRouteCORSShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayRouteCORSShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayRouteCORSShrinkRequest) SetAcceptLanguage(v string) *UpdateGatewayRouteCORSShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayRouteCORSShrinkRequest) SetCorsJSONShrink(v string) *UpdateGatewayRouteCORSShrinkRequest {
	s.CorsJSONShrink = &v
	return s
}

func (s *UpdateGatewayRouteCORSShrinkRequest) SetGatewayId(v int64) *UpdateGatewayRouteCORSShrinkRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayRouteCORSShrinkRequest) SetGatewayUniqueId(v string) *UpdateGatewayRouteCORSShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayRouteCORSShrinkRequest) SetId(v int64) *UpdateGatewayRouteCORSShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayRouteCORSShrinkRequest) Validate() error {
	return dara.Validate(s)
}
