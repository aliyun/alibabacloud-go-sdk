// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayFlowRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateGatewayFlowRuleResponseBodyData) *UpdateGatewayFlowRuleResponseBody
	GetData() *UpdateGatewayFlowRuleResponseBodyData
	SetRequestId(v string) *UpdateGatewayFlowRuleResponseBody
	GetRequestId() *string
}

type UpdateGatewayFlowRuleResponseBody struct {
	Data *UpdateGatewayFlowRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 2B74E7F7-DF54-5AB1-B8F2-67391B83****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGatewayFlowRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayFlowRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayFlowRuleResponseBody) GetData() *UpdateGatewayFlowRuleResponseBodyData {
	return s.Data
}

func (s *UpdateGatewayFlowRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayFlowRuleResponseBody) SetData(v *UpdateGatewayFlowRuleResponseBodyData) *UpdateGatewayFlowRuleResponseBody {
	s.Data = v
	return s
}

func (s *UpdateGatewayFlowRuleResponseBody) SetRequestId(v string) *UpdateGatewayFlowRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayFlowRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayFlowRuleResponseBodyData struct {
	// example:
	//
	// 0
	BehaviorType *int32 `json:"BehaviorType,omitempty" xml:"BehaviorType,omitempty"`
	// example:
	//
	// 0
	BodyEncoding *int32 `json:"BodyEncoding,omitempty" xml:"BodyEncoding,omitempty"`
	// example:
	//
	// 0
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 14407
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// gw-e2d226bba4b2445c9e29fa7f8216****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// example:
	//
	// 549
	Id     *int64   `json:"Id,omitempty" xml:"Id,omitempty"`
	IdList []*int64 `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	LimitMode *int32 `json:"LimitMode,omitempty" xml:"LimitMode,omitempty"`
	// example:
	//
	// key=value
	ResponseAdditionalHeaders *string `json:"ResponseAdditionalHeaders,omitempty" xml:"ResponseAdditionalHeaders,omitempty"`
	// example:
	//
	// Text
	ResponseContentBody *string `json:"ResponseContentBody,omitempty" xml:"ResponseContentBody,omitempty"`
	// example:
	//
	// www.******.com
	ResponseRedirectUrl *string `json:"ResponseRedirectUrl,omitempty" xml:"ResponseRedirectUrl,omitempty"`
	// example:
	//
	// 429
	ResponseStatusCode *int32 `json:"ResponseStatusCode,omitempty" xml:"ResponseStatusCode,omitempty"`
	// example:
	//
	// 48811
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	// example:
	//
	// routeA
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// example:
	//
	// 10
	StatDurationMs *int32 `json:"StatDurationMs,omitempty" xml:"StatDurationMs,omitempty"`
	// example:
	//
	// 10
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s UpdateGatewayFlowRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayFlowRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateGatewayFlowRuleResponseBodyData) GetBehaviorType() *int32 {
	return s.BehaviorType
}

func (s *UpdateGatewayFlowRuleResponseBodyData) GetBodyEncoding() *int32 {
	return s.BodyEncoding
}

func (s *UpdateGatewayFlowRuleResponseBodyData) GetEnable() *int32 {
	return s.Enable
}

func (s *UpdateGatewayFlowRuleResponseBodyData) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayFlowRuleResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayFlowRuleResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayFlowRuleResponseBodyData) GetIdList() []*int64 {
	return s.IdList
}

func (s *UpdateGatewayFlowRuleResponseBodyData) GetLimitMode() *int32 {
	return s.LimitMode
}

func (s *UpdateGatewayFlowRuleResponseBodyData) GetResponseAdditionalHeaders() *string {
	return s.ResponseAdditionalHeaders
}

func (s *UpdateGatewayFlowRuleResponseBodyData) GetResponseContentBody() *string {
	return s.ResponseContentBody
}

func (s *UpdateGatewayFlowRuleResponseBodyData) GetResponseRedirectUrl() *string {
	return s.ResponseRedirectUrl
}

func (s *UpdateGatewayFlowRuleResponseBodyData) GetResponseStatusCode() *int32 {
	return s.ResponseStatusCode
}

func (s *UpdateGatewayFlowRuleResponseBodyData) GetRouteId() *int64 {
	return s.RouteId
}

func (s *UpdateGatewayFlowRuleResponseBodyData) GetRouteName() *string {
	return s.RouteName
}

func (s *UpdateGatewayFlowRuleResponseBodyData) GetStatDurationMs() *int32 {
	return s.StatDurationMs
}

func (s *UpdateGatewayFlowRuleResponseBodyData) GetThreshold() *int32 {
	return s.Threshold
}

func (s *UpdateGatewayFlowRuleResponseBodyData) SetBehaviorType(v int32) *UpdateGatewayFlowRuleResponseBodyData {
	s.BehaviorType = &v
	return s
}

func (s *UpdateGatewayFlowRuleResponseBodyData) SetBodyEncoding(v int32) *UpdateGatewayFlowRuleResponseBodyData {
	s.BodyEncoding = &v
	return s
}

func (s *UpdateGatewayFlowRuleResponseBodyData) SetEnable(v int32) *UpdateGatewayFlowRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *UpdateGatewayFlowRuleResponseBodyData) SetGatewayId(v int64) *UpdateGatewayFlowRuleResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayFlowRuleResponseBodyData) SetGatewayUniqueId(v string) *UpdateGatewayFlowRuleResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayFlowRuleResponseBodyData) SetId(v int64) *UpdateGatewayFlowRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateGatewayFlowRuleResponseBodyData) SetIdList(v []*int64) *UpdateGatewayFlowRuleResponseBodyData {
	s.IdList = v
	return s
}

func (s *UpdateGatewayFlowRuleResponseBodyData) SetLimitMode(v int32) *UpdateGatewayFlowRuleResponseBodyData {
	s.LimitMode = &v
	return s
}

func (s *UpdateGatewayFlowRuleResponseBodyData) SetResponseAdditionalHeaders(v string) *UpdateGatewayFlowRuleResponseBodyData {
	s.ResponseAdditionalHeaders = &v
	return s
}

func (s *UpdateGatewayFlowRuleResponseBodyData) SetResponseContentBody(v string) *UpdateGatewayFlowRuleResponseBodyData {
	s.ResponseContentBody = &v
	return s
}

func (s *UpdateGatewayFlowRuleResponseBodyData) SetResponseRedirectUrl(v string) *UpdateGatewayFlowRuleResponseBodyData {
	s.ResponseRedirectUrl = &v
	return s
}

func (s *UpdateGatewayFlowRuleResponseBodyData) SetResponseStatusCode(v int32) *UpdateGatewayFlowRuleResponseBodyData {
	s.ResponseStatusCode = &v
	return s
}

func (s *UpdateGatewayFlowRuleResponseBodyData) SetRouteId(v int64) *UpdateGatewayFlowRuleResponseBodyData {
	s.RouteId = &v
	return s
}

func (s *UpdateGatewayFlowRuleResponseBodyData) SetRouteName(v string) *UpdateGatewayFlowRuleResponseBodyData {
	s.RouteName = &v
	return s
}

func (s *UpdateGatewayFlowRuleResponseBodyData) SetStatDurationMs(v int32) *UpdateGatewayFlowRuleResponseBodyData {
	s.StatDurationMs = &v
	return s
}

func (s *UpdateGatewayFlowRuleResponseBodyData) SetThreshold(v int32) *UpdateGatewayFlowRuleResponseBodyData {
	s.Threshold = &v
	return s
}

func (s *UpdateGatewayFlowRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
