// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteHTTPRewriteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayRouteHTTPRewriteRequest
	GetAcceptLanguage() *string
	SetGatewayId(v int64) *UpdateGatewayRouteHTTPRewriteRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdateGatewayRouteHTTPRewriteRequest
	GetGatewayUniqueId() *string
	SetHttpRewriteJSON(v string) *UpdateGatewayRouteHTTPRewriteRequest
	GetHttpRewriteJSON() *string
	SetId(v int64) *UpdateGatewayRouteHTTPRewriteRequest
	GetId() *int64
}

type UpdateGatewayRouteHTTPRewriteRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// 430
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-90392d768a3847a7b804c505254da96d
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The information about the rewrite policy. The JSON format is supported.
	//
	// example:
	//
	// {"pathType":"PRE","path":"/","status":"off"}
	HttpRewriteJSON *string `json:"HttpRewriteJSON,omitempty" xml:"HttpRewriteJSON,omitempty"`
	// The ID of the route.
	//
	// example:
	//
	// 238
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateGatewayRouteHTTPRewriteRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteHTTPRewriteRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteHTTPRewriteRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayRouteHTTPRewriteRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayRouteHTTPRewriteRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayRouteHTTPRewriteRequest) GetHttpRewriteJSON() *string {
	return s.HttpRewriteJSON
}

func (s *UpdateGatewayRouteHTTPRewriteRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayRouteHTTPRewriteRequest) SetAcceptLanguage(v string) *UpdateGatewayRouteHTTPRewriteRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteRequest) SetGatewayId(v int64) *UpdateGatewayRouteHTTPRewriteRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteRequest) SetGatewayUniqueId(v string) *UpdateGatewayRouteHTTPRewriteRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteRequest) SetHttpRewriteJSON(v string) *UpdateGatewayRouteHTTPRewriteRequest {
	s.HttpRewriteJSON = &v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteRequest) SetId(v int64) *UpdateGatewayRouteHTTPRewriteRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteRequest) Validate() error {
	return dara.Validate(s)
}
