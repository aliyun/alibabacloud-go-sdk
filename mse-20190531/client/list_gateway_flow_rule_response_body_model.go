// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayFlowRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListGatewayFlowRuleResponseBodyData) *ListGatewayFlowRuleResponseBody
	GetData() *ListGatewayFlowRuleResponseBodyData
	SetRequestId(v string) *ListGatewayFlowRuleResponseBody
	GetRequestId() *string
}

type ListGatewayFlowRuleResponseBody struct {
	Data *ListGatewayFlowRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 0951EBF0-798E-5E0B-8D38-460A14AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGatewayFlowRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayFlowRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayFlowRuleResponseBody) GetData() *ListGatewayFlowRuleResponseBodyData {
	return s.Data
}

func (s *ListGatewayFlowRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayFlowRuleResponseBody) SetData(v *ListGatewayFlowRuleResponseBodyData) *ListGatewayFlowRuleResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayFlowRuleResponseBody) SetRequestId(v string) *ListGatewayFlowRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListGatewayFlowRuleResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result   []*ListGatewayFlowRuleResponseBodyDataResult  `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Results  []*ListGatewayFlowRuleResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListGatewayFlowRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayFlowRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayFlowRuleResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayFlowRuleResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayFlowRuleResponseBodyData) GetResult() []*ListGatewayFlowRuleResponseBodyDataResult {
	return s.Result
}

func (s *ListGatewayFlowRuleResponseBodyData) GetResults() []*ListGatewayFlowRuleResponseBodyDataResults {
	return s.Results
}

func (s *ListGatewayFlowRuleResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListGatewayFlowRuleResponseBodyData) SetPageNumber(v int32) *ListGatewayFlowRuleResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyData) SetPageSize(v int32) *ListGatewayFlowRuleResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyData) SetResult(v []*ListGatewayFlowRuleResponseBodyDataResult) *ListGatewayFlowRuleResponseBodyData {
	s.Result = v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyData) SetResults(v []*ListGatewayFlowRuleResponseBodyDataResults) *ListGatewayFlowRuleResponseBodyData {
	s.Results = v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyData) SetTotalSize(v int32) *ListGatewayFlowRuleResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListGatewayFlowRuleResponseBodyDataResult struct {
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
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ListGatewayFlowRuleResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayFlowRuleResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) GetBehaviorType() *int32 {
	return s.BehaviorType
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) GetBodyEncoding() *int32 {
	return s.BodyEncoding
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) GetEnable() *int32 {
	return s.Enable
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) GetId() *int64 {
	return s.Id
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) GetResponseContentBody() *string {
	return s.ResponseContentBody
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) GetResponseRedirectUrl() *string {
	return s.ResponseRedirectUrl
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) GetResponseStatusCode() *int32 {
	return s.ResponseStatusCode
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) GetRouteId() *int64 {
	return s.RouteId
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) GetRouteName() *string {
	return s.RouteName
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) GetThreshold() *int32 {
	return s.Threshold
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) SetBehaviorType(v int32) *ListGatewayFlowRuleResponseBodyDataResult {
	s.BehaviorType = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) SetBodyEncoding(v int32) *ListGatewayFlowRuleResponseBodyDataResult {
	s.BodyEncoding = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) SetEnable(v int32) *ListGatewayFlowRuleResponseBodyDataResult {
	s.Enable = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) SetGatewayId(v int64) *ListGatewayFlowRuleResponseBodyDataResult {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) SetGatewayUniqueId(v string) *ListGatewayFlowRuleResponseBodyDataResult {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) SetId(v int64) *ListGatewayFlowRuleResponseBodyDataResult {
	s.Id = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) SetResponseContentBody(v string) *ListGatewayFlowRuleResponseBodyDataResult {
	s.ResponseContentBody = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) SetResponseRedirectUrl(v string) *ListGatewayFlowRuleResponseBodyDataResult {
	s.ResponseRedirectUrl = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) SetResponseStatusCode(v int32) *ListGatewayFlowRuleResponseBodyDataResult {
	s.ResponseStatusCode = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) SetRouteId(v int64) *ListGatewayFlowRuleResponseBodyDataResult {
	s.RouteId = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) SetRouteName(v string) *ListGatewayFlowRuleResponseBodyDataResult {
	s.RouteName = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) SetThreshold(v int32) *ListGatewayFlowRuleResponseBodyDataResult {
	s.Threshold = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}

type ListGatewayFlowRuleResponseBodyDataResults struct {
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

func (s ListGatewayFlowRuleResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayFlowRuleResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) GetBehaviorType() *int32 {
	return s.BehaviorType
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) GetBodyEncoding() *int32 {
	return s.BodyEncoding
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) GetEnable() *int32 {
	return s.Enable
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) GetId() *int64 {
	return s.Id
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) GetIdList() []*int64 {
	return s.IdList
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) GetLimitMode() *int32 {
	return s.LimitMode
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) GetResponseAdditionalHeaders() *string {
	return s.ResponseAdditionalHeaders
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) GetResponseContentBody() *string {
	return s.ResponseContentBody
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) GetResponseRedirectUrl() *string {
	return s.ResponseRedirectUrl
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) GetResponseStatusCode() *int32 {
	return s.ResponseStatusCode
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) GetRouteId() *int64 {
	return s.RouteId
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) GetRouteName() *string {
	return s.RouteName
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) GetStatDurationMs() *int32 {
	return s.StatDurationMs
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) GetThreshold() *int32 {
	return s.Threshold
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) SetBehaviorType(v int32) *ListGatewayFlowRuleResponseBodyDataResults {
	s.BehaviorType = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) SetBodyEncoding(v int32) *ListGatewayFlowRuleResponseBodyDataResults {
	s.BodyEncoding = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) SetEnable(v int32) *ListGatewayFlowRuleResponseBodyDataResults {
	s.Enable = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) SetGatewayId(v int64) *ListGatewayFlowRuleResponseBodyDataResults {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) SetGatewayUniqueId(v string) *ListGatewayFlowRuleResponseBodyDataResults {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) SetId(v int64) *ListGatewayFlowRuleResponseBodyDataResults {
	s.Id = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) SetIdList(v []*int64) *ListGatewayFlowRuleResponseBodyDataResults {
	s.IdList = v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) SetLimitMode(v int32) *ListGatewayFlowRuleResponseBodyDataResults {
	s.LimitMode = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) SetResponseAdditionalHeaders(v string) *ListGatewayFlowRuleResponseBodyDataResults {
	s.ResponseAdditionalHeaders = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) SetResponseContentBody(v string) *ListGatewayFlowRuleResponseBodyDataResults {
	s.ResponseContentBody = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) SetResponseRedirectUrl(v string) *ListGatewayFlowRuleResponseBodyDataResults {
	s.ResponseRedirectUrl = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) SetResponseStatusCode(v int32) *ListGatewayFlowRuleResponseBodyDataResults {
	s.ResponseStatusCode = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) SetRouteId(v int64) *ListGatewayFlowRuleResponseBodyDataResults {
	s.RouteId = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) SetRouteName(v string) *ListGatewayFlowRuleResponseBodyDataResults {
	s.RouteName = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) SetStatDurationMs(v int32) *ListGatewayFlowRuleResponseBodyDataResults {
	s.StatDurationMs = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) SetThreshold(v int32) *ListGatewayFlowRuleResponseBodyDataResults {
	s.Threshold = &v
	return s
}

func (s *ListGatewayFlowRuleResponseBodyDataResults) Validate() error {
	return dara.Validate(s)
}
