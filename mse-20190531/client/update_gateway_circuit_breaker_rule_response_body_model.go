// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayCircuitBreakerRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateGatewayCircuitBreakerRuleResponseBodyData) *UpdateGatewayCircuitBreakerRuleResponseBody
	GetData() *UpdateGatewayCircuitBreakerRuleResponseBodyData
	SetRequestId(v string) *UpdateGatewayCircuitBreakerRuleResponseBody
	GetRequestId() *string
}

type UpdateGatewayCircuitBreakerRuleResponseBody struct {
	// The returned data.
	Data *UpdateGatewayCircuitBreakerRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// DC34E4A3-5F1C-4E40-86EA-02EDF967****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGatewayCircuitBreakerRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayCircuitBreakerRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBody) GetData() *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	return s.Data
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBody) SetData(v *UpdateGatewayCircuitBreakerRuleResponseBodyData) *UpdateGatewayCircuitBreakerRuleResponseBody {
	s.Data = v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBody) SetRequestId(v string) *UpdateGatewayCircuitBreakerRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateGatewayCircuitBreakerRuleResponseBodyData struct {
	// The web fallback behavior type. Valid values:
	//
	// - 0: Return specified text.
	//
	// - 1: Redirect to a specified page.
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
	// example:
	//
	// gw-c9bc5afd61014165bd58f621b491****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the circuit breaker rule.
	//
	// example:
	//
	// 369
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IDs of the filter rules.
	IdList []*int64 `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
	// The threshold conversion mode. Valid values:
	//
	// - 0: Standalone.
	//
	// - 1: Cluster.
	//
	// - 2: Cluster to standalone.
	//
	// example:
	//
	// 0
	LimitMode *int32 `json:"LimitMode,omitempty" xml:"LimitMode,omitempty"`
	// The slow call threshold in slow call mode.
	//
	// example:
	//
	// 10
	MaxAllowedMs *int32 `json:"MaxAllowedMs,omitempty" xml:"MaxAllowedMs,omitempty"`
	// The minimum number of requests that triggers circuit breaking.
	//
	// example:
	//
	// 10
	MinRequestAmount *int32 `json:"MinRequestAmount,omitempty" xml:"MinRequestAmount,omitempty"`
	// The recovery duration for circuit breaking.
	//
	// example:
	//
	// 12
	RecoveryTimeoutSec *int32 `json:"RecoveryTimeoutSec,omitempty" xml:"RecoveryTimeoutSec,omitempty"`
	// The custom headers, which are entered as key-value pairs.
	//
	// example:
	//
	// key=value
	ResponseAdditionalHeaders *string `json:"ResponseAdditionalHeaders,omitempty" xml:"ResponseAdditionalHeaders,omitempty"`
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
	// 201
	ResponseStatusCode *int32 `json:"ResponseStatusCode,omitempty" xml:"ResponseStatusCode,omitempty"`
	// The ID of the route.
	//
	// example:
	//
	// 645
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	// The name of the route.
	//
	// example:
	//
	// routeName
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The statistical duration for circuit breaking.
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
	// example:
	//
	// 0
	Strategy *int32 `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// The percentage that triggers circuit breaking.
	//
	// example:
	//
	// 10
	TriggerRatio *int32 `json:"TriggerRatio,omitempty" xml:"TriggerRatio,omitempty"`
}

func (s UpdateGatewayCircuitBreakerRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayCircuitBreakerRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) GetBehaviorType() *int32 {
	return s.BehaviorType
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) GetBodyEncoding() *int32 {
	return s.BodyEncoding
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) GetEnable() *int32 {
	return s.Enable
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) GetIdList() []*int64 {
	return s.IdList
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) GetLimitMode() *int32 {
	return s.LimitMode
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) GetMaxAllowedMs() *int32 {
	return s.MaxAllowedMs
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) GetMinRequestAmount() *int32 {
	return s.MinRequestAmount
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) GetRecoveryTimeoutSec() *int32 {
	return s.RecoveryTimeoutSec
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) GetResponseAdditionalHeaders() *string {
	return s.ResponseAdditionalHeaders
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) GetResponseContentBody() *string {
	return s.ResponseContentBody
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) GetResponseRedirectUrl() *string {
	return s.ResponseRedirectUrl
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) GetResponseStatusCode() *int32 {
	return s.ResponseStatusCode
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) GetRouteId() *int64 {
	return s.RouteId
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) GetRouteName() *string {
	return s.RouteName
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) GetStatDurationSec() *int32 {
	return s.StatDurationSec
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) GetStrategy() *int32 {
	return s.Strategy
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) GetTriggerRatio() *int32 {
	return s.TriggerRatio
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) SetBehaviorType(v int32) *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	s.BehaviorType = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) SetBodyEncoding(v int32) *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	s.BodyEncoding = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) SetEnable(v int32) *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) SetGatewayId(v int64) *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) SetGatewayUniqueId(v string) *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) SetId(v int64) *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) SetIdList(v []*int64) *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	s.IdList = v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) SetLimitMode(v int32) *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	s.LimitMode = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) SetMaxAllowedMs(v int32) *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	s.MaxAllowedMs = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) SetMinRequestAmount(v int32) *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	s.MinRequestAmount = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) SetRecoveryTimeoutSec(v int32) *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	s.RecoveryTimeoutSec = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) SetResponseAdditionalHeaders(v string) *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	s.ResponseAdditionalHeaders = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) SetResponseContentBody(v string) *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	s.ResponseContentBody = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) SetResponseRedirectUrl(v string) *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	s.ResponseRedirectUrl = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) SetResponseStatusCode(v int32) *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	s.ResponseStatusCode = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) SetRouteId(v int64) *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	s.RouteId = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) SetRouteName(v string) *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	s.RouteName = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) SetStatDurationSec(v int32) *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	s.StatDurationSec = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) SetStrategy(v int32) *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	s.Strategy = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) SetTriggerRatio(v int32) *UpdateGatewayCircuitBreakerRuleResponseBodyData {
	s.TriggerRatio = &v
	return s
}

func (s *UpdateGatewayCircuitBreakerRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
