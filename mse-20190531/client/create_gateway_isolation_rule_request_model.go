// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayIsolationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateGatewayIsolationRuleRequest
	GetAcceptLanguage() *string
	SetBehaviorType(v int32) *CreateGatewayIsolationRuleRequest
	GetBehaviorType() *int32
	SetBodyEncoding(v int32) *CreateGatewayIsolationRuleRequest
	GetBodyEncoding() *int32
	SetEnable(v int32) *CreateGatewayIsolationRuleRequest
	GetEnable() *int32
	SetGatewayId(v int64) *CreateGatewayIsolationRuleRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *CreateGatewayIsolationRuleRequest
	GetGatewayUniqueId() *string
	SetMaxConcurrency(v int32) *CreateGatewayIsolationRuleRequest
	GetMaxConcurrency() *int32
	SetResponseContentBody(v string) *CreateGatewayIsolationRuleRequest
	GetResponseContentBody() *string
	SetResponseRedirectUrl(v string) *CreateGatewayIsolationRuleRequest
	GetResponseRedirectUrl() *string
	SetResponseStatusCode(v int32) *CreateGatewayIsolationRuleRequest
	GetResponseStatusCode() *int32
	SetRouteId(v int64) *CreateGatewayIsolationRuleRequest
	GetRouteId() *int64
	SetRouteName(v string) *CreateGatewayIsolationRuleRequest
	GetRouteName() *string
}

type CreateGatewayIsolationRuleRequest struct {
	// The language in which you want to display the results. Valid values: zh and en. zh is the default value, which indicates Chinese. en indicates English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The type of a web fallback behavior.
	//
	// 0: The specified content is returned.
	//
	// 1: The specified page is displayed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	BehaviorType *int32 `json:"BehaviorType,omitempty" xml:"BehaviorType,omitempty"`
	// The encoding type.
	//
	// 0: plain text
	//
	// 1: JSON
	//
	// example:
	//
	// 0
	BodyEncoding *int32 `json:"BodyEncoding,omitempty" xml:"BodyEncoding,omitempty"`
	// Specifies whether to enable the rule.
	//
	// 0: disabled.
	//
	// 1: enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The gateway ID.
	//
	// example:
	//
	// 14407
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-e2d226bba4b2445c9e29fa7f8216****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The maximum number of concurrent requests.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// The HTTP text to be returned.
	//
	// example:
	//
	// text
	ResponseContentBody *string `json:"ResponseContentBody,omitempty" xml:"ResponseContentBody,omitempty"`
	// The address to be redirected to.
	//
	// example:
	//
	// www.******.com
	ResponseRedirectUrl *string `json:"ResponseRedirectUrl,omitempty" xml:"ResponseRedirectUrl,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 429
	ResponseStatusCode *int32 `json:"ResponseStatusCode,omitempty" xml:"ResponseStatusCode,omitempty"`
	// The route ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 52853
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	// The name of the route.
	//
	// This parameter is required.
	//
	// example:
	//
	// routeName
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
}

func (s CreateGatewayIsolationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayIsolationRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayIsolationRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateGatewayIsolationRuleRequest) GetBehaviorType() *int32 {
	return s.BehaviorType
}

func (s *CreateGatewayIsolationRuleRequest) GetBodyEncoding() *int32 {
	return s.BodyEncoding
}

func (s *CreateGatewayIsolationRuleRequest) GetEnable() *int32 {
	return s.Enable
}

func (s *CreateGatewayIsolationRuleRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *CreateGatewayIsolationRuleRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *CreateGatewayIsolationRuleRequest) GetMaxConcurrency() *int32 {
	return s.MaxConcurrency
}

func (s *CreateGatewayIsolationRuleRequest) GetResponseContentBody() *string {
	return s.ResponseContentBody
}

func (s *CreateGatewayIsolationRuleRequest) GetResponseRedirectUrl() *string {
	return s.ResponseRedirectUrl
}

func (s *CreateGatewayIsolationRuleRequest) GetResponseStatusCode() *int32 {
	return s.ResponseStatusCode
}

func (s *CreateGatewayIsolationRuleRequest) GetRouteId() *int64 {
	return s.RouteId
}

func (s *CreateGatewayIsolationRuleRequest) GetRouteName() *string {
	return s.RouteName
}

func (s *CreateGatewayIsolationRuleRequest) SetAcceptLanguage(v string) *CreateGatewayIsolationRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateGatewayIsolationRuleRequest) SetBehaviorType(v int32) *CreateGatewayIsolationRuleRequest {
	s.BehaviorType = &v
	return s
}

func (s *CreateGatewayIsolationRuleRequest) SetBodyEncoding(v int32) *CreateGatewayIsolationRuleRequest {
	s.BodyEncoding = &v
	return s
}

func (s *CreateGatewayIsolationRuleRequest) SetEnable(v int32) *CreateGatewayIsolationRuleRequest {
	s.Enable = &v
	return s
}

func (s *CreateGatewayIsolationRuleRequest) SetGatewayId(v int64) *CreateGatewayIsolationRuleRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateGatewayIsolationRuleRequest) SetGatewayUniqueId(v string) *CreateGatewayIsolationRuleRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *CreateGatewayIsolationRuleRequest) SetMaxConcurrency(v int32) *CreateGatewayIsolationRuleRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *CreateGatewayIsolationRuleRequest) SetResponseContentBody(v string) *CreateGatewayIsolationRuleRequest {
	s.ResponseContentBody = &v
	return s
}

func (s *CreateGatewayIsolationRuleRequest) SetResponseRedirectUrl(v string) *CreateGatewayIsolationRuleRequest {
	s.ResponseRedirectUrl = &v
	return s
}

func (s *CreateGatewayIsolationRuleRequest) SetResponseStatusCode(v int32) *CreateGatewayIsolationRuleRequest {
	s.ResponseStatusCode = &v
	return s
}

func (s *CreateGatewayIsolationRuleRequest) SetRouteId(v int64) *CreateGatewayIsolationRuleRequest {
	s.RouteId = &v
	return s
}

func (s *CreateGatewayIsolationRuleRequest) SetRouteName(v string) *CreateGatewayIsolationRuleRequest {
	s.RouteName = &v
	return s
}

func (s *CreateGatewayIsolationRuleRequest) Validate() error {
	return dara.Validate(s)
}
