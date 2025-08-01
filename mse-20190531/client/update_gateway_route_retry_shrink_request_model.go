// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteRetryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayRouteRetryShrinkRequest
	GetAcceptLanguage() *string
	SetGatewayId(v int64) *UpdateGatewayRouteRetryShrinkRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdateGatewayRouteRetryShrinkRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *UpdateGatewayRouteRetryShrinkRequest
	GetId() *int64
	SetRetryJSONShrink(v string) *UpdateGatewayRouteRetryShrinkRequest
	GetRetryJSONShrink() *string
}

type UpdateGatewayRouteRetryShrinkRequest struct {
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
	// 501
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-3f97e2989c344f35ab3fd62b19f1d10a
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the associated record.
	//
	// example:
	//
	// 508
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The information about the retry policy.
	RetryJSONShrink *string `json:"RetryJSON,omitempty" xml:"RetryJSON,omitempty"`
}

func (s UpdateGatewayRouteRetryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteRetryShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteRetryShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayRouteRetryShrinkRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayRouteRetryShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayRouteRetryShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayRouteRetryShrinkRequest) GetRetryJSONShrink() *string {
	return s.RetryJSONShrink
}

func (s *UpdateGatewayRouteRetryShrinkRequest) SetAcceptLanguage(v string) *UpdateGatewayRouteRetryShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayRouteRetryShrinkRequest) SetGatewayId(v int64) *UpdateGatewayRouteRetryShrinkRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayRouteRetryShrinkRequest) SetGatewayUniqueId(v string) *UpdateGatewayRouteRetryShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayRouteRetryShrinkRequest) SetId(v int64) *UpdateGatewayRouteRetryShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayRouteRetryShrinkRequest) SetRetryJSONShrink(v string) *UpdateGatewayRouteRetryShrinkRequest {
	s.RetryJSONShrink = &v
	return s
}

func (s *UpdateGatewayRouteRetryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
