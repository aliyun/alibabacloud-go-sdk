// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayFlowRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateGatewayFlowRuleRequest
	GetAcceptLanguage() *string
	SetBehaviorType(v int32) *CreateGatewayFlowRuleRequest
	GetBehaviorType() *int32
	SetBodyEncoding(v int32) *CreateGatewayFlowRuleRequest
	GetBodyEncoding() *int32
	SetEnable(v int32) *CreateGatewayFlowRuleRequest
	GetEnable() *int32
	SetGatewayId(v int64) *CreateGatewayFlowRuleRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *CreateGatewayFlowRuleRequest
	GetGatewayUniqueId() *string
	SetResponseContentBody(v string) *CreateGatewayFlowRuleRequest
	GetResponseContentBody() *string
	SetResponseRedirectUrl(v string) *CreateGatewayFlowRuleRequest
	GetResponseRedirectUrl() *string
	SetResponseStatusCode(v int32) *CreateGatewayFlowRuleRequest
	GetResponseStatusCode() *int32
	SetRouteId(v int64) *CreateGatewayFlowRuleRequest
	GetRouteId() *int64
	SetRouteName(v string) *CreateGatewayFlowRuleRequest
	GetRouteName() *string
	SetThreshold(v int32) *CreateGatewayFlowRuleRequest
	GetThreshold() *int32
}

type CreateGatewayFlowRuleRequest struct {
	// The language in which you want to display the results. Valid values: zh and en. zh indicates Chinese, which is the default value. en indicates English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The type of the web fallback behavior.
	//
	// 0: returns the specified content.
	//
	// 1: redirects to the specified page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	BehaviorType *int32 `json:"BehaviorType,omitempty" xml:"BehaviorType,omitempty"`
	// The encoding format.
	//
	// 0: normal text
	//
	// 1: JSON
	//
	// example:
	//
	// 0
	BodyEncoding *int32 `json:"BodyEncoding,omitempty" xml:"BodyEncoding,omitempty"`
	// Specifies whether to enable the throttling rule.
	//
	// 0: disables the throttling rule.
	//
	// 1: enables the throttling rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The ID of the gateway.
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
	// The ID of the route.
	//
	// This parameter is required.
	//
	// example:
	//
	// 52853
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	// The name of the routing rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// routeName
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The overall queries per second (QPS) threshold.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateGatewayFlowRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayFlowRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayFlowRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateGatewayFlowRuleRequest) GetBehaviorType() *int32 {
	return s.BehaviorType
}

func (s *CreateGatewayFlowRuleRequest) GetBodyEncoding() *int32 {
	return s.BodyEncoding
}

func (s *CreateGatewayFlowRuleRequest) GetEnable() *int32 {
	return s.Enable
}

func (s *CreateGatewayFlowRuleRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *CreateGatewayFlowRuleRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *CreateGatewayFlowRuleRequest) GetResponseContentBody() *string {
	return s.ResponseContentBody
}

func (s *CreateGatewayFlowRuleRequest) GetResponseRedirectUrl() *string {
	return s.ResponseRedirectUrl
}

func (s *CreateGatewayFlowRuleRequest) GetResponseStatusCode() *int32 {
	return s.ResponseStatusCode
}

func (s *CreateGatewayFlowRuleRequest) GetRouteId() *int64 {
	return s.RouteId
}

func (s *CreateGatewayFlowRuleRequest) GetRouteName() *string {
	return s.RouteName
}

func (s *CreateGatewayFlowRuleRequest) GetThreshold() *int32 {
	return s.Threshold
}

func (s *CreateGatewayFlowRuleRequest) SetAcceptLanguage(v string) *CreateGatewayFlowRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateGatewayFlowRuleRequest) SetBehaviorType(v int32) *CreateGatewayFlowRuleRequest {
	s.BehaviorType = &v
	return s
}

func (s *CreateGatewayFlowRuleRequest) SetBodyEncoding(v int32) *CreateGatewayFlowRuleRequest {
	s.BodyEncoding = &v
	return s
}

func (s *CreateGatewayFlowRuleRequest) SetEnable(v int32) *CreateGatewayFlowRuleRequest {
	s.Enable = &v
	return s
}

func (s *CreateGatewayFlowRuleRequest) SetGatewayId(v int64) *CreateGatewayFlowRuleRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateGatewayFlowRuleRequest) SetGatewayUniqueId(v string) *CreateGatewayFlowRuleRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *CreateGatewayFlowRuleRequest) SetResponseContentBody(v string) *CreateGatewayFlowRuleRequest {
	s.ResponseContentBody = &v
	return s
}

func (s *CreateGatewayFlowRuleRequest) SetResponseRedirectUrl(v string) *CreateGatewayFlowRuleRequest {
	s.ResponseRedirectUrl = &v
	return s
}

func (s *CreateGatewayFlowRuleRequest) SetResponseStatusCode(v int32) *CreateGatewayFlowRuleRequest {
	s.ResponseStatusCode = &v
	return s
}

func (s *CreateGatewayFlowRuleRequest) SetRouteId(v int64) *CreateGatewayFlowRuleRequest {
	s.RouteId = &v
	return s
}

func (s *CreateGatewayFlowRuleRequest) SetRouteName(v string) *CreateGatewayFlowRuleRequest {
	s.RouteName = &v
	return s
}

func (s *CreateGatewayFlowRuleRequest) SetThreshold(v int32) *CreateGatewayFlowRuleRequest {
	s.Threshold = &v
	return s
}

func (s *CreateGatewayFlowRuleRequest) Validate() error {
	return dara.Validate(s)
}
