// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayCircuitBreakerRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListGatewayCircuitBreakerRuleResponseBodyData) *ListGatewayCircuitBreakerRuleResponseBody
	GetData() *ListGatewayCircuitBreakerRuleResponseBodyData
	SetRequestId(v string) *ListGatewayCircuitBreakerRuleResponseBody
	GetRequestId() *string
}

type ListGatewayCircuitBreakerRuleResponseBody struct {
	Data *ListGatewayCircuitBreakerRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// DC34E4A3-5F1C-4E40-86EA-02EDF967****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGatewayCircuitBreakerRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayCircuitBreakerRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayCircuitBreakerRuleResponseBody) GetData() *ListGatewayCircuitBreakerRuleResponseBodyData {
	return s.Data
}

func (s *ListGatewayCircuitBreakerRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayCircuitBreakerRuleResponseBody) SetData(v *ListGatewayCircuitBreakerRuleResponseBodyData) *ListGatewayCircuitBreakerRuleResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBody) SetRequestId(v string) *ListGatewayCircuitBreakerRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayCircuitBreakerRuleResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result   []*ListGatewayCircuitBreakerRuleResponseBodyDataResult  `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Results  []*ListGatewayCircuitBreakerRuleResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// example:
	//
	// 11
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListGatewayCircuitBreakerRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayCircuitBreakerRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyData) GetResult() []*ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	return s.Result
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyData) GetResults() []*ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	return s.Results
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyData) SetPageNumber(v int32) *ListGatewayCircuitBreakerRuleResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyData) SetPageSize(v int32) *ListGatewayCircuitBreakerRuleResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyData) SetResult(v []*ListGatewayCircuitBreakerRuleResponseBodyDataResult) *ListGatewayCircuitBreakerRuleResponseBodyData {
	s.Result = v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyData) SetResults(v []*ListGatewayCircuitBreakerRuleResponseBodyDataResults) *ListGatewayCircuitBreakerRuleResponseBodyData {
	s.Results = v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyData) SetTotalSize(v int32) *ListGatewayCircuitBreakerRuleResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGatewayCircuitBreakerRuleResponseBodyDataResult struct {
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
	// 11919
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// gw-5017305290e14centbrveca****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// example:
	//
	// 467
	Id     *int64   `json:"Id,omitempty" xml:"Id,omitempty"`
	IdList []*int64 `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	LimitMode *int32 `json:"LimitMode,omitempty" xml:"LimitMode,omitempty"`
	// example:
	//
	// 14
	MaxAllowedMs *int32 `json:"MaxAllowedMs,omitempty" xml:"MaxAllowedMs,omitempty"`
	// example:
	//
	// 10
	MinRequestAmount *int32 `json:"MinRequestAmount,omitempty" xml:"MinRequestAmount,omitempty"`
	// example:
	//
	// 12
	RecoveryTimeoutSec *int32 `json:"RecoveryTimeoutSec,omitempty" xml:"RecoveryTimeoutSec,omitempty"`
	// example:
	//
	// key=value
	ResponseAdditionalHeaders *string `json:"ResponseAdditionalHeaders,omitempty" xml:"ResponseAdditionalHeaders,omitempty"`
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
	// 204
	ResponseStatusCode *int32 `json:"ResponseStatusCode,omitempty" xml:"ResponseStatusCode,omitempty"`
	// example:
	//
	// 3091
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	// example:
	//
	// routeName
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// example:
	//
	// 11
	StatDurationSec *int32 `json:"StatDurationSec,omitempty" xml:"StatDurationSec,omitempty"`
	// example:
	//
	// 0
	Strategy *int32 `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// example:
	//
	// 10
	TriggerRatio *int32 `json:"TriggerRatio,omitempty" xml:"TriggerRatio,omitempty"`
}

func (s ListGatewayCircuitBreakerRuleResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayCircuitBreakerRuleResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) GetBehaviorType() *int32 {
	return s.BehaviorType
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) GetBodyEncoding() *int32 {
	return s.BodyEncoding
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) GetEnable() *int32 {
	return s.Enable
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) GetId() *int64 {
	return s.Id
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) GetIdList() []*int64 {
	return s.IdList
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) GetLimitMode() *int32 {
	return s.LimitMode
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) GetMaxAllowedMs() *int32 {
	return s.MaxAllowedMs
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) GetMinRequestAmount() *int32 {
	return s.MinRequestAmount
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) GetRecoveryTimeoutSec() *int32 {
	return s.RecoveryTimeoutSec
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) GetResponseAdditionalHeaders() *string {
	return s.ResponseAdditionalHeaders
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) GetResponseContentBody() *string {
	return s.ResponseContentBody
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) GetResponseRedirectUrl() *string {
	return s.ResponseRedirectUrl
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) GetResponseStatusCode() *int32 {
	return s.ResponseStatusCode
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) GetRouteId() *int64 {
	return s.RouteId
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) GetRouteName() *string {
	return s.RouteName
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) GetStatDurationSec() *int32 {
	return s.StatDurationSec
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) GetStrategy() *int32 {
	return s.Strategy
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) GetTriggerRatio() *int32 {
	return s.TriggerRatio
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) SetBehaviorType(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	s.BehaviorType = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) SetBodyEncoding(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	s.BodyEncoding = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) SetEnable(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	s.Enable = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) SetGatewayId(v int64) *ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) SetGatewayUniqueId(v string) *ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) SetId(v int64) *ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	s.Id = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) SetIdList(v []*int64) *ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	s.IdList = v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) SetLimitMode(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	s.LimitMode = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) SetMaxAllowedMs(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	s.MaxAllowedMs = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) SetMinRequestAmount(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	s.MinRequestAmount = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) SetRecoveryTimeoutSec(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	s.RecoveryTimeoutSec = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) SetResponseAdditionalHeaders(v string) *ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	s.ResponseAdditionalHeaders = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) SetResponseContentBody(v string) *ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	s.ResponseContentBody = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) SetResponseRedirectUrl(v string) *ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	s.ResponseRedirectUrl = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) SetResponseStatusCode(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	s.ResponseStatusCode = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) SetRouteId(v int64) *ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	s.RouteId = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) SetRouteName(v string) *ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	s.RouteName = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) SetStatDurationSec(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	s.StatDurationSec = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) SetStrategy(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	s.Strategy = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) SetTriggerRatio(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResult {
	s.TriggerRatio = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}

type ListGatewayCircuitBreakerRuleResponseBodyDataResults struct {
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
	// 11919
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// gw-1ee34548c68f4778a25c05abd657****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// example:
	//
	// 467
	Id     *int64   `json:"Id,omitempty" xml:"Id,omitempty"`
	IdList []*int64 `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	LimitMode *int32 `json:"LimitMode,omitempty" xml:"LimitMode,omitempty"`
	// example:
	//
	// 10
	MaxAllowedMs *int32 `json:"MaxAllowedMs,omitempty" xml:"MaxAllowedMs,omitempty"`
	// example:
	//
	// 10
	MinRequestAmount *int32 `json:"MinRequestAmount,omitempty" xml:"MinRequestAmount,omitempty"`
	// example:
	//
	// 12
	RecoveryTimeoutSec *int32 `json:"RecoveryTimeoutSec,omitempty" xml:"RecoveryTimeoutSec,omitempty"`
	// example:
	//
	// key=value
	ResponseAdditionalHeaders *string `json:"ResponseAdditionalHeaders,omitempty" xml:"ResponseAdditionalHeaders,omitempty"`
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
	// 204
	ResponseStatusCode *int32 `json:"ResponseStatusCode,omitempty" xml:"ResponseStatusCode,omitempty"`
	// example:
	//
	// 3450
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	// example:
	//
	// routeName
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// example:
	//
	// 11
	StatDurationSec *int32 `json:"StatDurationSec,omitempty" xml:"StatDurationSec,omitempty"`
	// example:
	//
	// 0
	Strategy *int32 `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// example:
	//
	// 10
	TriggerRatio *int32 `json:"TriggerRatio,omitempty" xml:"TriggerRatio,omitempty"`
}

func (s ListGatewayCircuitBreakerRuleResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayCircuitBreakerRuleResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) GetBehaviorType() *int32 {
	return s.BehaviorType
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) GetBodyEncoding() *int32 {
	return s.BodyEncoding
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) GetEnable() *int32 {
	return s.Enable
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) GetId() *int64 {
	return s.Id
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) GetIdList() []*int64 {
	return s.IdList
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) GetLimitMode() *int32 {
	return s.LimitMode
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) GetMaxAllowedMs() *int32 {
	return s.MaxAllowedMs
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) GetMinRequestAmount() *int32 {
	return s.MinRequestAmount
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) GetRecoveryTimeoutSec() *int32 {
	return s.RecoveryTimeoutSec
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) GetResponseAdditionalHeaders() *string {
	return s.ResponseAdditionalHeaders
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) GetResponseContentBody() *string {
	return s.ResponseContentBody
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) GetResponseRedirectUrl() *string {
	return s.ResponseRedirectUrl
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) GetResponseStatusCode() *int32 {
	return s.ResponseStatusCode
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) GetRouteId() *int64 {
	return s.RouteId
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) GetRouteName() *string {
	return s.RouteName
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) GetStatDurationSec() *int32 {
	return s.StatDurationSec
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) GetStrategy() *int32 {
	return s.Strategy
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) GetTriggerRatio() *int32 {
	return s.TriggerRatio
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) SetBehaviorType(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	s.BehaviorType = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) SetBodyEncoding(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	s.BodyEncoding = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) SetEnable(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	s.Enable = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) SetGatewayId(v int64) *ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) SetGatewayUniqueId(v string) *ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) SetId(v int64) *ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	s.Id = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) SetIdList(v []*int64) *ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	s.IdList = v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) SetLimitMode(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	s.LimitMode = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) SetMaxAllowedMs(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	s.MaxAllowedMs = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) SetMinRequestAmount(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	s.MinRequestAmount = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) SetRecoveryTimeoutSec(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	s.RecoveryTimeoutSec = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) SetResponseAdditionalHeaders(v string) *ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	s.ResponseAdditionalHeaders = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) SetResponseContentBody(v string) *ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	s.ResponseContentBody = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) SetResponseRedirectUrl(v string) *ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	s.ResponseRedirectUrl = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) SetResponseStatusCode(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	s.ResponseStatusCode = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) SetRouteId(v int64) *ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	s.RouteId = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) SetRouteName(v string) *ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	s.RouteName = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) SetStatDurationSec(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	s.StatDurationSec = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) SetStrategy(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	s.Strategy = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) SetTriggerRatio(v int32) *ListGatewayCircuitBreakerRuleResponseBodyDataResults {
	s.TriggerRatio = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleResponseBodyDataResults) Validate() error {
	return dara.Validate(s)
}
