// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayCircuitBreakerRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayCircuitBreakerRuleRequest
	GetAcceptLanguage() *string
	SetBehaviorType(v int32) *UpdateGatewayCircuitBreakerRuleRequest
	GetBehaviorType() *int32
	SetBodyEncoding(v int32) *UpdateGatewayCircuitBreakerRuleRequest
	GetBodyEncoding() *int32
	SetEnable(v int32) *UpdateGatewayCircuitBreakerRuleRequest
	GetEnable() *int32
	SetGatewayId(v int64) *UpdateGatewayCircuitBreakerRuleRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdateGatewayCircuitBreakerRuleRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *UpdateGatewayCircuitBreakerRuleRequest
	GetId() *int64
	SetMaxAllowedMs(v int32) *UpdateGatewayCircuitBreakerRuleRequest
	GetMaxAllowedMs() *int32
	SetMinRequestAmount(v int32) *UpdateGatewayCircuitBreakerRuleRequest
	GetMinRequestAmount() *int32
	SetRecoveryTimeoutSec(v int32) *UpdateGatewayCircuitBreakerRuleRequest
	GetRecoveryTimeoutSec() *int32
	SetResponseContentBody(v string) *UpdateGatewayCircuitBreakerRuleRequest
	GetResponseContentBody() *string
	SetResponseRedirectUrl(v string) *UpdateGatewayCircuitBreakerRuleRequest
	GetResponseRedirectUrl() *string
	SetResponseStatusCode(v int32) *UpdateGatewayCircuitBreakerRuleRequest
	GetResponseStatusCode() *int32
	SetRouteId(v int64) *UpdateGatewayCircuitBreakerRuleRequest
	GetRouteId() *int64
	SetRouteName(v string) *UpdateGatewayCircuitBreakerRuleRequest
	GetRouteName() *string
	SetStatDurationSec(v int32) *UpdateGatewayCircuitBreakerRuleRequest
	GetStatDurationSec() *int32
	SetStrategy(v int32) *UpdateGatewayCircuitBreakerRuleRequest
	GetStrategy() *int32
	SetTriggerRatio(v int32) *UpdateGatewayCircuitBreakerRuleRequest
	GetTriggerRatio() *int32
}

type UpdateGatewayCircuitBreakerRuleRequest struct {
	// The language of the returned results. Valid values: zh (default): Chinese, en: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The web fallback behavior type. Valid values:
	//
	// - 0: Return specified text.
	//
	// - 1: Redirect to a specified page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	BehaviorType *int32 `json:"BehaviorType,omitempty" xml:"BehaviorType,omitempty"`
	// The encoding type. Valid values:
	//
	// - 0: Plain text.
	//
	// - 1: JSON.
	//
	// example:
	//
	// 0
	BodyEncoding *int32 `json:"BodyEncoding,omitempty" xml:"BodyEncoding,omitempty"`
	// Specifies whether to enable the rule. Valid values:
	//
	// - 0: Disabled.
	//
	// - 1: Enabled.
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
	// gw-c9bc5afd61014165bd58f621b491****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the circuit breaker rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 369
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The slow call threshold in slow call mode.
	//
	// example:
	//
	// 10
	MaxAllowedMs *int32 `json:"MaxAllowedMs,omitempty" xml:"MaxAllowedMs,omitempty"`
	// The minimum number of requests that triggers circuit breaking.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	MinRequestAmount *int32 `json:"MinRequestAmount,omitempty" xml:"MinRequestAmount,omitempty"`
	// The recovery duration for circuit breaking.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	RecoveryTimeoutSec *int32 `json:"RecoveryTimeoutSec,omitempty" xml:"RecoveryTimeoutSec,omitempty"`
	// The HTTP response text.
	//
	// example:
	//
	// text
	ResponseContentBody *string `json:"ResponseContentBody,omitempty" xml:"ResponseContentBody,omitempty"`
	// The redirect URL.
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
	// 645
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	// The name of the route.
	//
	// This parameter is required.
	//
	// example:
	//
	// routeName
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The statistical duration for circuit breaking.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	StatDurationSec *int32 `json:"StatDurationSec,omitempty" xml:"StatDurationSec,omitempty"`
	// The threshold type. Valid values:
	//
	// - 0: Slow call.
	//
	// - 1: Exception.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Strategy *int32 `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// The percentage that triggers circuit breaking.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	TriggerRatio *int32 `json:"TriggerRatio,omitempty" xml:"TriggerRatio,omitempty"`
}

func (s UpdateGatewayCircuitBreakerRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayCircuitBreakerRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) GetBehaviorType() *int32 {
	return s.BehaviorType
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) GetBodyEncoding() *int32 {
	return s.BodyEncoding
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) GetEnable() *int32 {
	return s.Enable
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) GetMaxAllowedMs() *int32 {
	return s.MaxAllowedMs
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) GetMinRequestAmount() *int32 {
	return s.MinRequestAmount
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) GetRecoveryTimeoutSec() *int32 {
	return s.RecoveryTimeoutSec
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) GetResponseContentBody() *string {
	return s.ResponseContentBody
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) GetResponseRedirectUrl() *string {
	return s.ResponseRedirectUrl
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) GetResponseStatusCode() *int32 {
	return s.ResponseStatusCode
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) GetRouteId() *int64 {
	return s.RouteId
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) GetRouteName() *string {
	return s.RouteName
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) GetStatDurationSec() *int32 {
	return s.StatDurationSec
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) GetStrategy() *int32 {
	return s.Strategy
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) GetTriggerRatio() *int32 {
	return s.TriggerRatio
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) SetAcceptLanguage(v string) *UpdateGatewayCircuitBreakerRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) SetBehaviorType(v int32) *UpdateGatewayCircuitBreakerRuleRequest {
	s.BehaviorType = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) SetBodyEncoding(v int32) *UpdateGatewayCircuitBreakerRuleRequest {
	s.BodyEncoding = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) SetEnable(v int32) *UpdateGatewayCircuitBreakerRuleRequest {
	s.Enable = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) SetGatewayId(v int64) *UpdateGatewayCircuitBreakerRuleRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) SetGatewayUniqueId(v string) *UpdateGatewayCircuitBreakerRuleRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) SetId(v int64) *UpdateGatewayCircuitBreakerRuleRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) SetMaxAllowedMs(v int32) *UpdateGatewayCircuitBreakerRuleRequest {
	s.MaxAllowedMs = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) SetMinRequestAmount(v int32) *UpdateGatewayCircuitBreakerRuleRequest {
	s.MinRequestAmount = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) SetRecoveryTimeoutSec(v int32) *UpdateGatewayCircuitBreakerRuleRequest {
	s.RecoveryTimeoutSec = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) SetResponseContentBody(v string) *UpdateGatewayCircuitBreakerRuleRequest {
	s.ResponseContentBody = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) SetResponseRedirectUrl(v string) *UpdateGatewayCircuitBreakerRuleRequest {
	s.ResponseRedirectUrl = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) SetResponseStatusCode(v int32) *UpdateGatewayCircuitBreakerRuleRequest {
	s.ResponseStatusCode = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) SetRouteId(v int64) *UpdateGatewayCircuitBreakerRuleRequest {
	s.RouteId = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) SetRouteName(v string) *UpdateGatewayCircuitBreakerRuleRequest {
	s.RouteName = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) SetStatDurationSec(v int32) *UpdateGatewayCircuitBreakerRuleRequest {
	s.StatDurationSec = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) SetStrategy(v int32) *UpdateGatewayCircuitBreakerRuleRequest {
	s.Strategy = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) SetTriggerRatio(v int32) *UpdateGatewayCircuitBreakerRuleRequest {
	s.TriggerRatio = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleRequest) Validate() error {
	return dara.Validate(s)
}
