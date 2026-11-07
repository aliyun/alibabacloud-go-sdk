// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayCircuitBreakerRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateGatewayCircuitBreakerRuleRequest
	GetAcceptLanguage() *string
	SetBehaviorType(v int32) *CreateGatewayCircuitBreakerRuleRequest
	GetBehaviorType() *int32
	SetBodyEncoding(v int32) *CreateGatewayCircuitBreakerRuleRequest
	GetBodyEncoding() *int32
	SetEnable(v int32) *CreateGatewayCircuitBreakerRuleRequest
	GetEnable() *int32
	SetGatewayId(v int64) *CreateGatewayCircuitBreakerRuleRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *CreateGatewayCircuitBreakerRuleRequest
	GetGatewayUniqueId() *string
	SetMaxAllowedMs(v int32) *CreateGatewayCircuitBreakerRuleRequest
	GetMaxAllowedMs() *int32
	SetMinRequestAmount(v int32) *CreateGatewayCircuitBreakerRuleRequest
	GetMinRequestAmount() *int32
	SetRecoveryTimeoutSec(v int32) *CreateGatewayCircuitBreakerRuleRequest
	GetRecoveryTimeoutSec() *int32
	SetResponseContentBody(v string) *CreateGatewayCircuitBreakerRuleRequest
	GetResponseContentBody() *string
	SetResponseRedirectUrl(v string) *CreateGatewayCircuitBreakerRuleRequest
	GetResponseRedirectUrl() *string
	SetResponseStatusCode(v int32) *CreateGatewayCircuitBreakerRuleRequest
	GetResponseStatusCode() *int32
	SetRouteId(v int64) *CreateGatewayCircuitBreakerRuleRequest
	GetRouteId() *int64
	SetRouteName(v string) *CreateGatewayCircuitBreakerRuleRequest
	GetRouteName() *string
	SetStatDurationSec(v int32) *CreateGatewayCircuitBreakerRuleRequest
	GetStatDurationSec() *int32
	SetStrategy(v int32) *CreateGatewayCircuitBreakerRuleRequest
	GetStrategy() *int32
	SetTriggerRatio(v int32) *CreateGatewayCircuitBreakerRuleRequest
	GetTriggerRatio() *int32
}

type CreateGatewayCircuitBreakerRuleRequest struct {
	// The language in which you want the results to be displayed. Valid values: zh (default): Chinese. en: English.
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
	// 0: disables the rule.
	//
	// 1: enables the rule.
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
	// gw-c9bc5afd61014165bd58f621b491*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The threshold in slow call proportion mode.
	//
	// example:
	//
	// 10
	MaxAllowedMs *int32 `json:"MaxAllowedMs,omitempty" xml:"MaxAllowedMs,omitempty"`
	// The minimum number of requests to trigger circuit breaking.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	MinRequestAmount *int32 `json:"MinRequestAmount,omitempty" xml:"MinRequestAmount,omitempty"`
	// The circuit breaking timeout period.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	RecoveryTimeoutSec *int32 `json:"RecoveryTimeoutSec,omitempty" xml:"RecoveryTimeoutSec,omitempty"`
	// The HTTP text to be returned.
	//
	// example:
	//
	// text
	ResponseContentBody *string `json:"ResponseContentBody,omitempty" xml:"ResponseContentBody,omitempty"`
	// The redirection URL.
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
	// 645
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	// The route name.
	//
	// This parameter is required.
	//
	// example:
	//
	// routeName
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The statistical duration for triggering circuit breaking.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	StatDurationSec *int32 `json:"StatDurationSec,omitempty" xml:"StatDurationSec,omitempty"`
	// The type of the threshold.
	//
	// 0: slow call proportion
	//
	// 1: failure proportion
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Strategy *int32 `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// The request percentage that triggers circuit breaking.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	TriggerRatio *int32 `json:"TriggerRatio,omitempty" xml:"TriggerRatio,omitempty"`
}

func (s CreateGatewayCircuitBreakerRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayCircuitBreakerRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayCircuitBreakerRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateGatewayCircuitBreakerRuleRequest) GetBehaviorType() *int32 {
	return s.BehaviorType
}

func (s *CreateGatewayCircuitBreakerRuleRequest) GetBodyEncoding() *int32 {
	return s.BodyEncoding
}

func (s *CreateGatewayCircuitBreakerRuleRequest) GetEnable() *int32 {
	return s.Enable
}

func (s *CreateGatewayCircuitBreakerRuleRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *CreateGatewayCircuitBreakerRuleRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *CreateGatewayCircuitBreakerRuleRequest) GetMaxAllowedMs() *int32 {
	return s.MaxAllowedMs
}

func (s *CreateGatewayCircuitBreakerRuleRequest) GetMinRequestAmount() *int32 {
	return s.MinRequestAmount
}

func (s *CreateGatewayCircuitBreakerRuleRequest) GetRecoveryTimeoutSec() *int32 {
	return s.RecoveryTimeoutSec
}

func (s *CreateGatewayCircuitBreakerRuleRequest) GetResponseContentBody() *string {
	return s.ResponseContentBody
}

func (s *CreateGatewayCircuitBreakerRuleRequest) GetResponseRedirectUrl() *string {
	return s.ResponseRedirectUrl
}

func (s *CreateGatewayCircuitBreakerRuleRequest) GetResponseStatusCode() *int32 {
	return s.ResponseStatusCode
}

func (s *CreateGatewayCircuitBreakerRuleRequest) GetRouteId() *int64 {
	return s.RouteId
}

func (s *CreateGatewayCircuitBreakerRuleRequest) GetRouteName() *string {
	return s.RouteName
}

func (s *CreateGatewayCircuitBreakerRuleRequest) GetStatDurationSec() *int32 {
	return s.StatDurationSec
}

func (s *CreateGatewayCircuitBreakerRuleRequest) GetStrategy() *int32 {
	return s.Strategy
}

func (s *CreateGatewayCircuitBreakerRuleRequest) GetTriggerRatio() *int32 {
	return s.TriggerRatio
}

func (s *CreateGatewayCircuitBreakerRuleRequest) SetAcceptLanguage(v string) *CreateGatewayCircuitBreakerRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleRequest) SetBehaviorType(v int32) *CreateGatewayCircuitBreakerRuleRequest {
	s.BehaviorType = &v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleRequest) SetBodyEncoding(v int32) *CreateGatewayCircuitBreakerRuleRequest {
	s.BodyEncoding = &v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleRequest) SetEnable(v int32) *CreateGatewayCircuitBreakerRuleRequest {
	s.Enable = &v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleRequest) SetGatewayId(v int64) *CreateGatewayCircuitBreakerRuleRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleRequest) SetGatewayUniqueId(v string) *CreateGatewayCircuitBreakerRuleRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleRequest) SetMaxAllowedMs(v int32) *CreateGatewayCircuitBreakerRuleRequest {
	s.MaxAllowedMs = &v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleRequest) SetMinRequestAmount(v int32) *CreateGatewayCircuitBreakerRuleRequest {
	s.MinRequestAmount = &v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleRequest) SetRecoveryTimeoutSec(v int32) *CreateGatewayCircuitBreakerRuleRequest {
	s.RecoveryTimeoutSec = &v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleRequest) SetResponseContentBody(v string) *CreateGatewayCircuitBreakerRuleRequest {
	s.ResponseContentBody = &v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleRequest) SetResponseRedirectUrl(v string) *CreateGatewayCircuitBreakerRuleRequest {
	s.ResponseRedirectUrl = &v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleRequest) SetResponseStatusCode(v int32) *CreateGatewayCircuitBreakerRuleRequest {
	s.ResponseStatusCode = &v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleRequest) SetRouteId(v int64) *CreateGatewayCircuitBreakerRuleRequest {
	s.RouteId = &v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleRequest) SetRouteName(v string) *CreateGatewayCircuitBreakerRuleRequest {
	s.RouteName = &v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleRequest) SetStatDurationSec(v int32) *CreateGatewayCircuitBreakerRuleRequest {
	s.StatDurationSec = &v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleRequest) SetStrategy(v int32) *CreateGatewayCircuitBreakerRuleRequest {
	s.Strategy = &v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleRequest) SetTriggerRatio(v int32) *CreateGatewayCircuitBreakerRuleRequest {
	s.TriggerRatio = &v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleRequest) Validate() error {
	return dara.Validate(s)
}
