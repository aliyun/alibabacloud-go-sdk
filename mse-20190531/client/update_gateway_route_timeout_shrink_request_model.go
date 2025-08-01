// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteTimeoutShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayRouteTimeoutShrinkRequest
	GetAcceptLanguage() *string
	SetGatewayId(v int64) *UpdateGatewayRouteTimeoutShrinkRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdateGatewayRouteTimeoutShrinkRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *UpdateGatewayRouteTimeoutShrinkRequest
	GetId() *int64
	SetTimeoutJSONShrink(v string) *UpdateGatewayRouteTimeoutShrinkRequest
	GetTimeoutJSONShrink() *string
}

type UpdateGatewayRouteTimeoutShrinkRequest struct {
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
	// 85
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-533290d279c1405f9628c64f7c8272ee
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the associated record.
	//
	// example:
	//
	// 567
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The timeout period.
	TimeoutJSONShrink *string `json:"TimeoutJSON,omitempty" xml:"TimeoutJSON,omitempty"`
}

func (s UpdateGatewayRouteTimeoutShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteTimeoutShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteTimeoutShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayRouteTimeoutShrinkRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayRouteTimeoutShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayRouteTimeoutShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayRouteTimeoutShrinkRequest) GetTimeoutJSONShrink() *string {
	return s.TimeoutJSONShrink
}

func (s *UpdateGatewayRouteTimeoutShrinkRequest) SetAcceptLanguage(v string) *UpdateGatewayRouteTimeoutShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayRouteTimeoutShrinkRequest) SetGatewayId(v int64) *UpdateGatewayRouteTimeoutShrinkRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayRouteTimeoutShrinkRequest) SetGatewayUniqueId(v string) *UpdateGatewayRouteTimeoutShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayRouteTimeoutShrinkRequest) SetId(v int64) *UpdateGatewayRouteTimeoutShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayRouteTimeoutShrinkRequest) SetTimeoutJSONShrink(v string) *UpdateGatewayRouteTimeoutShrinkRequest {
	s.TimeoutJSONShrink = &v
	return s
}

func (s *UpdateGatewayRouteTimeoutShrinkRequest) Validate() error {
	return dara.Validate(s)
}
