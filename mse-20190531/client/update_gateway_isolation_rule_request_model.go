// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayIsolationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayIsolationRuleRequest
	GetAcceptLanguage() *string
	SetBehaviorType(v int32) *UpdateGatewayIsolationRuleRequest
	GetBehaviorType() *int32
	SetBodyEncoding(v int32) *UpdateGatewayIsolationRuleRequest
	GetBodyEncoding() *int32
	SetEnable(v int32) *UpdateGatewayIsolationRuleRequest
	GetEnable() *int32
	SetGatewayId(v int64) *UpdateGatewayIsolationRuleRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdateGatewayIsolationRuleRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *UpdateGatewayIsolationRuleRequest
	GetId() *int64
	SetMaxConcurrency(v int32) *UpdateGatewayIsolationRuleRequest
	GetMaxConcurrency() *int32
	SetResponseContentBody(v string) *UpdateGatewayIsolationRuleRequest
	GetResponseContentBody() *string
	SetResponseRedirectUrl(v string) *UpdateGatewayIsolationRuleRequest
	GetResponseRedirectUrl() *string
	SetResponseStatusCode(v int32) *UpdateGatewayIsolationRuleRequest
	GetResponseStatusCode() *int32
	SetRouteId(v int64) *UpdateGatewayIsolationRuleRequest
	GetRouteId() *int64
	SetRouteName(v string) *UpdateGatewayIsolationRuleRequest
	GetRouteName() *string
}

type UpdateGatewayIsolationRuleRequest struct {
	// The language in which you want to display the results. Valid values: zh (default): Chinese. en: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The type of the web fallback behavior.
	//
	// 0: The specified content is returned.
	//
	// 1: Redirect to the specified page.
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
	// 1: JSON text
	//
	// example:
	//
	// 0
	BodyEncoding *int32 `json:"BodyEncoding,omitempty" xml:"BodyEncoding,omitempty"`
	// Specifies whether to enable the rule.
	//
	// 0: disable
	//
	// 1: enable
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
	// The ID of the concurrency rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 358
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// The URL to be redirected to.
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
	// The route name.
	//
	// This parameter is required.
	//
	// example:
	//
	// routeName
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
}

func (s UpdateGatewayIsolationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayIsolationRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayIsolationRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayIsolationRuleRequest) GetBehaviorType() *int32 {
	return s.BehaviorType
}

func (s *UpdateGatewayIsolationRuleRequest) GetBodyEncoding() *int32 {
	return s.BodyEncoding
}

func (s *UpdateGatewayIsolationRuleRequest) GetEnable() *int32 {
	return s.Enable
}

func (s *UpdateGatewayIsolationRuleRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayIsolationRuleRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayIsolationRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayIsolationRuleRequest) GetMaxConcurrency() *int32 {
	return s.MaxConcurrency
}

func (s *UpdateGatewayIsolationRuleRequest) GetResponseContentBody() *string {
	return s.ResponseContentBody
}

func (s *UpdateGatewayIsolationRuleRequest) GetResponseRedirectUrl() *string {
	return s.ResponseRedirectUrl
}

func (s *UpdateGatewayIsolationRuleRequest) GetResponseStatusCode() *int32 {
	return s.ResponseStatusCode
}

func (s *UpdateGatewayIsolationRuleRequest) GetRouteId() *int64 {
	return s.RouteId
}

func (s *UpdateGatewayIsolationRuleRequest) GetRouteName() *string {
	return s.RouteName
}

func (s *UpdateGatewayIsolationRuleRequest) SetAcceptLanguage(v string) *UpdateGatewayIsolationRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayIsolationRuleRequest) SetBehaviorType(v int32) *UpdateGatewayIsolationRuleRequest {
	s.BehaviorType = &v
	return s
}

func (s *UpdateGatewayIsolationRuleRequest) SetBodyEncoding(v int32) *UpdateGatewayIsolationRuleRequest {
	s.BodyEncoding = &v
	return s
}

func (s *UpdateGatewayIsolationRuleRequest) SetEnable(v int32) *UpdateGatewayIsolationRuleRequest {
	s.Enable = &v
	return s
}

func (s *UpdateGatewayIsolationRuleRequest) SetGatewayId(v int64) *UpdateGatewayIsolationRuleRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayIsolationRuleRequest) SetGatewayUniqueId(v string) *UpdateGatewayIsolationRuleRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayIsolationRuleRequest) SetId(v int64) *UpdateGatewayIsolationRuleRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayIsolationRuleRequest) SetMaxConcurrency(v int32) *UpdateGatewayIsolationRuleRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *UpdateGatewayIsolationRuleRequest) SetResponseContentBody(v string) *UpdateGatewayIsolationRuleRequest {
	s.ResponseContentBody = &v
	return s
}

func (s *UpdateGatewayIsolationRuleRequest) SetResponseRedirectUrl(v string) *UpdateGatewayIsolationRuleRequest {
	s.ResponseRedirectUrl = &v
	return s
}

func (s *UpdateGatewayIsolationRuleRequest) SetResponseStatusCode(v int32) *UpdateGatewayIsolationRuleRequest {
	s.ResponseStatusCode = &v
	return s
}

func (s *UpdateGatewayIsolationRuleRequest) SetRouteId(v int64) *UpdateGatewayIsolationRuleRequest {
	s.RouteId = &v
	return s
}

func (s *UpdateGatewayIsolationRuleRequest) SetRouteName(v string) *UpdateGatewayIsolationRuleRequest {
	s.RouteName = &v
	return s
}

func (s *UpdateGatewayIsolationRuleRequest) Validate() error {
	return dara.Validate(s)
}
