// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayFlowRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayFlowRuleRequest
	GetAcceptLanguage() *string
	SetBehaviorType(v int32) *UpdateGatewayFlowRuleRequest
	GetBehaviorType() *int32
	SetBodyEncoding(v int32) *UpdateGatewayFlowRuleRequest
	GetBodyEncoding() *int32
	SetEnable(v int32) *UpdateGatewayFlowRuleRequest
	GetEnable() *int32
	SetGatewayId(v int64) *UpdateGatewayFlowRuleRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdateGatewayFlowRuleRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *UpdateGatewayFlowRuleRequest
	GetId() *int64
	SetResponseContentBody(v string) *UpdateGatewayFlowRuleRequest
	GetResponseContentBody() *string
	SetResponseRedirectUrl(v string) *UpdateGatewayFlowRuleRequest
	GetResponseRedirectUrl() *string
	SetResponseStatusCode(v int32) *UpdateGatewayFlowRuleRequest
	GetResponseStatusCode() *int32
	SetRouteId(v int64) *UpdateGatewayFlowRuleRequest
	GetRouteId() *int64
	SetRouteName(v string) *UpdateGatewayFlowRuleRequest
	GetRouteName() *string
	SetThreshold(v int32) *UpdateGatewayFlowRuleRequest
	GetThreshold() *int32
}

type UpdateGatewayFlowRuleRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	BehaviorType *int32 `json:"BehaviorType,omitempty" xml:"BehaviorType,omitempty"`
	// example:
	//
	// 0
	BodyEncoding *int32 `json:"BodyEncoding,omitempty" xml:"BodyEncoding,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 14407
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gw-e2d226bba4b2445c9e29fa7f8216****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 549
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// text
	ResponseContentBody *string `json:"ResponseContentBody,omitempty" xml:"ResponseContentBody,omitempty"`
	// example:
	//
	// www.******.com
	ResponseRedirectUrl *string `json:"ResponseRedirectUrl,omitempty" xml:"ResponseRedirectUrl,omitempty"`
	// example:
	//
	// 429
	ResponseStatusCode *int32 `json:"ResponseStatusCode,omitempty" xml:"ResponseStatusCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 48811
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// routeA
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s UpdateGatewayFlowRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayFlowRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayFlowRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayFlowRuleRequest) GetBehaviorType() *int32 {
	return s.BehaviorType
}

func (s *UpdateGatewayFlowRuleRequest) GetBodyEncoding() *int32 {
	return s.BodyEncoding
}

func (s *UpdateGatewayFlowRuleRequest) GetEnable() *int32 {
	return s.Enable
}

func (s *UpdateGatewayFlowRuleRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayFlowRuleRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayFlowRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayFlowRuleRequest) GetResponseContentBody() *string {
	return s.ResponseContentBody
}

func (s *UpdateGatewayFlowRuleRequest) GetResponseRedirectUrl() *string {
	return s.ResponseRedirectUrl
}

func (s *UpdateGatewayFlowRuleRequest) GetResponseStatusCode() *int32 {
	return s.ResponseStatusCode
}

func (s *UpdateGatewayFlowRuleRequest) GetRouteId() *int64 {
	return s.RouteId
}

func (s *UpdateGatewayFlowRuleRequest) GetRouteName() *string {
	return s.RouteName
}

func (s *UpdateGatewayFlowRuleRequest) GetThreshold() *int32 {
	return s.Threshold
}

func (s *UpdateGatewayFlowRuleRequest) SetAcceptLanguage(v string) *UpdateGatewayFlowRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayFlowRuleRequest) SetBehaviorType(v int32) *UpdateGatewayFlowRuleRequest {
	s.BehaviorType = &v
	return s
}

func (s *UpdateGatewayFlowRuleRequest) SetBodyEncoding(v int32) *UpdateGatewayFlowRuleRequest {
	s.BodyEncoding = &v
	return s
}

func (s *UpdateGatewayFlowRuleRequest) SetEnable(v int32) *UpdateGatewayFlowRuleRequest {
	s.Enable = &v
	return s
}

func (s *UpdateGatewayFlowRuleRequest) SetGatewayId(v int64) *UpdateGatewayFlowRuleRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayFlowRuleRequest) SetGatewayUniqueId(v string) *UpdateGatewayFlowRuleRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayFlowRuleRequest) SetId(v int64) *UpdateGatewayFlowRuleRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayFlowRuleRequest) SetResponseContentBody(v string) *UpdateGatewayFlowRuleRequest {
	s.ResponseContentBody = &v
	return s
}

func (s *UpdateGatewayFlowRuleRequest) SetResponseRedirectUrl(v string) *UpdateGatewayFlowRuleRequest {
	s.ResponseRedirectUrl = &v
	return s
}

func (s *UpdateGatewayFlowRuleRequest) SetResponseStatusCode(v int32) *UpdateGatewayFlowRuleRequest {
	s.ResponseStatusCode = &v
	return s
}

func (s *UpdateGatewayFlowRuleRequest) SetRouteId(v int64) *UpdateGatewayFlowRuleRequest {
	s.RouteId = &v
	return s
}

func (s *UpdateGatewayFlowRuleRequest) SetRouteName(v string) *UpdateGatewayFlowRuleRequest {
	s.RouteName = &v
	return s
}

func (s *UpdateGatewayFlowRuleRequest) SetThreshold(v int32) *UpdateGatewayFlowRuleRequest {
	s.Threshold = &v
	return s
}

func (s *UpdateGatewayFlowRuleRequest) Validate() error {
	return dara.Validate(s)
}
