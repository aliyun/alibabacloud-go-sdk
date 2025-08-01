// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteAuthShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayRouteAuthShrinkRequest
	GetAcceptLanguage() *string
	SetAuthJSONShrink(v string) *UpdateGatewayRouteAuthShrinkRequest
	GetAuthJSONShrink() *string
	SetGatewayId(v int64) *UpdateGatewayRouteAuthShrinkRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdateGatewayRouteAuthShrinkRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *UpdateGatewayRouteAuthShrinkRequest
	GetId() *int64
}

type UpdateGatewayRouteAuthShrinkRequest struct {
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
	// The authentication configurations.
	//
	// This parameter is required.
	AuthJSONShrink *string `json:"AuthJSON,omitempty" xml:"AuthJSON,omitempty"`
	// The gateway ID.
	//
	// example:
	//
	// 102
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-0adf3ad751284cc69fcf9669fba*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The route ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 109
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateGatewayRouteAuthShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteAuthShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteAuthShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayRouteAuthShrinkRequest) GetAuthJSONShrink() *string {
	return s.AuthJSONShrink
}

func (s *UpdateGatewayRouteAuthShrinkRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayRouteAuthShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayRouteAuthShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayRouteAuthShrinkRequest) SetAcceptLanguage(v string) *UpdateGatewayRouteAuthShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayRouteAuthShrinkRequest) SetAuthJSONShrink(v string) *UpdateGatewayRouteAuthShrinkRequest {
	s.AuthJSONShrink = &v
	return s
}

func (s *UpdateGatewayRouteAuthShrinkRequest) SetGatewayId(v int64) *UpdateGatewayRouteAuthShrinkRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayRouteAuthShrinkRequest) SetGatewayUniqueId(v string) *UpdateGatewayRouteAuthShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayRouteAuthShrinkRequest) SetId(v int64) *UpdateGatewayRouteAuthShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayRouteAuthShrinkRequest) Validate() error {
	return dara.Validate(s)
}
