// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayIsolationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListGatewayIsolationRuleResponseBodyData) *ListGatewayIsolationRuleResponseBody
	GetData() *ListGatewayIsolationRuleResponseBodyData
	SetRequestId(v string) *ListGatewayIsolationRuleResponseBody
	GetRequestId() *string
}

type ListGatewayIsolationRuleResponseBody struct {
	Data *ListGatewayIsolationRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 6F025D43-8632-5716-AE9B-7EDDF16C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGatewayIsolationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayIsolationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayIsolationRuleResponseBody) GetData() *ListGatewayIsolationRuleResponseBodyData {
	return s.Data
}

func (s *ListGatewayIsolationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayIsolationRuleResponseBody) SetData(v *ListGatewayIsolationRuleResponseBodyData) *ListGatewayIsolationRuleResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayIsolationRuleResponseBody) SetRequestId(v string) *ListGatewayIsolationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayIsolationRuleResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result   []*ListGatewayIsolationRuleResponseBodyDataResult  `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Results  []*ListGatewayIsolationRuleResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListGatewayIsolationRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayIsolationRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayIsolationRuleResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayIsolationRuleResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayIsolationRuleResponseBodyData) GetResult() []*ListGatewayIsolationRuleResponseBodyDataResult {
	return s.Result
}

func (s *ListGatewayIsolationRuleResponseBodyData) GetResults() []*ListGatewayIsolationRuleResponseBodyDataResults {
	return s.Results
}

func (s *ListGatewayIsolationRuleResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListGatewayIsolationRuleResponseBodyData) SetPageNumber(v int32) *ListGatewayIsolationRuleResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyData) SetPageSize(v int32) *ListGatewayIsolationRuleResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyData) SetResult(v []*ListGatewayIsolationRuleResponseBodyDataResult) *ListGatewayIsolationRuleResponseBodyData {
	s.Result = v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyData) SetResults(v []*ListGatewayIsolationRuleResponseBodyDataResults) *ListGatewayIsolationRuleResponseBodyData {
	s.Results = v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyData) SetTotalSize(v int32) *ListGatewayIsolationRuleResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyData) Validate() error {
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

type ListGatewayIsolationRuleResponseBodyDataResult struct {
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
	// 358
	Id     *int64   `json:"Id,omitempty" xml:"Id,omitempty"`
	IdList []*int64 `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	LimitMode *int32 `json:"LimitMode,omitempty" xml:"LimitMode,omitempty"`
	// example:
	//
	// 100
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// example:
	//
	// key:value
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
	// 52853
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	// example:
	//
	// routeName
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
}

func (s ListGatewayIsolationRuleResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayIsolationRuleResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) GetBehaviorType() *int32 {
	return s.BehaviorType
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) GetBodyEncoding() *int32 {
	return s.BodyEncoding
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) GetEnable() *int32 {
	return s.Enable
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) GetId() *int64 {
	return s.Id
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) GetIdList() []*int64 {
	return s.IdList
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) GetLimitMode() *int32 {
	return s.LimitMode
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) GetMaxConcurrency() *int32 {
	return s.MaxConcurrency
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) GetResponseAdditionalHeaders() *string {
	return s.ResponseAdditionalHeaders
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) GetResponseContentBody() *string {
	return s.ResponseContentBody
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) GetResponseRedirectUrl() *string {
	return s.ResponseRedirectUrl
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) GetResponseStatusCode() *int32 {
	return s.ResponseStatusCode
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) GetRouteId() *int64 {
	return s.RouteId
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) GetRouteName() *string {
	return s.RouteName
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) SetBehaviorType(v int32) *ListGatewayIsolationRuleResponseBodyDataResult {
	s.BehaviorType = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) SetBodyEncoding(v int32) *ListGatewayIsolationRuleResponseBodyDataResult {
	s.BodyEncoding = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) SetEnable(v int32) *ListGatewayIsolationRuleResponseBodyDataResult {
	s.Enable = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) SetGatewayId(v int64) *ListGatewayIsolationRuleResponseBodyDataResult {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) SetGatewayUniqueId(v string) *ListGatewayIsolationRuleResponseBodyDataResult {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) SetId(v int64) *ListGatewayIsolationRuleResponseBodyDataResult {
	s.Id = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) SetIdList(v []*int64) *ListGatewayIsolationRuleResponseBodyDataResult {
	s.IdList = v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) SetLimitMode(v int32) *ListGatewayIsolationRuleResponseBodyDataResult {
	s.LimitMode = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) SetMaxConcurrency(v int32) *ListGatewayIsolationRuleResponseBodyDataResult {
	s.MaxConcurrency = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) SetResponseAdditionalHeaders(v string) *ListGatewayIsolationRuleResponseBodyDataResult {
	s.ResponseAdditionalHeaders = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) SetResponseContentBody(v string) *ListGatewayIsolationRuleResponseBodyDataResult {
	s.ResponseContentBody = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) SetResponseRedirectUrl(v string) *ListGatewayIsolationRuleResponseBodyDataResult {
	s.ResponseRedirectUrl = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) SetResponseStatusCode(v int32) *ListGatewayIsolationRuleResponseBodyDataResult {
	s.ResponseStatusCode = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) SetRouteId(v int64) *ListGatewayIsolationRuleResponseBodyDataResult {
	s.RouteId = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) SetRouteName(v string) *ListGatewayIsolationRuleResponseBodyDataResult {
	s.RouteName = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}

type ListGatewayIsolationRuleResponseBodyDataResults struct {
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
	// 358
	Id     *int64   `json:"Id,omitempty" xml:"Id,omitempty"`
	IdList []*int64 `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	LimitMode *int32 `json:"LimitMode,omitempty" xml:"LimitMode,omitempty"`
	// example:
	//
	// 100
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
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
	// 52853
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	// example:
	//
	// routeName
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
}

func (s ListGatewayIsolationRuleResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayIsolationRuleResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) GetBehaviorType() *int32 {
	return s.BehaviorType
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) GetBodyEncoding() *int32 {
	return s.BodyEncoding
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) GetEnable() *int32 {
	return s.Enable
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) GetId() *int64 {
	return s.Id
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) GetIdList() []*int64 {
	return s.IdList
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) GetLimitMode() *int32 {
	return s.LimitMode
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) GetMaxConcurrency() *int32 {
	return s.MaxConcurrency
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) GetResponseAdditionalHeaders() *string {
	return s.ResponseAdditionalHeaders
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) GetResponseContentBody() *string {
	return s.ResponseContentBody
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) GetResponseRedirectUrl() *string {
	return s.ResponseRedirectUrl
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) GetResponseStatusCode() *int32 {
	return s.ResponseStatusCode
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) GetRouteId() *int64 {
	return s.RouteId
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) GetRouteName() *string {
	return s.RouteName
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) SetBehaviorType(v int32) *ListGatewayIsolationRuleResponseBodyDataResults {
	s.BehaviorType = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) SetBodyEncoding(v int32) *ListGatewayIsolationRuleResponseBodyDataResults {
	s.BodyEncoding = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) SetEnable(v int32) *ListGatewayIsolationRuleResponseBodyDataResults {
	s.Enable = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) SetGatewayId(v int64) *ListGatewayIsolationRuleResponseBodyDataResults {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) SetGatewayUniqueId(v string) *ListGatewayIsolationRuleResponseBodyDataResults {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) SetId(v int64) *ListGatewayIsolationRuleResponseBodyDataResults {
	s.Id = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) SetIdList(v []*int64) *ListGatewayIsolationRuleResponseBodyDataResults {
	s.IdList = v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) SetLimitMode(v int32) *ListGatewayIsolationRuleResponseBodyDataResults {
	s.LimitMode = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) SetMaxConcurrency(v int32) *ListGatewayIsolationRuleResponseBodyDataResults {
	s.MaxConcurrency = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) SetResponseAdditionalHeaders(v string) *ListGatewayIsolationRuleResponseBodyDataResults {
	s.ResponseAdditionalHeaders = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) SetResponseContentBody(v string) *ListGatewayIsolationRuleResponseBodyDataResults {
	s.ResponseContentBody = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) SetResponseRedirectUrl(v string) *ListGatewayIsolationRuleResponseBodyDataResults {
	s.ResponseRedirectUrl = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) SetResponseStatusCode(v int32) *ListGatewayIsolationRuleResponseBodyDataResults {
	s.ResponseStatusCode = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) SetRouteId(v int64) *ListGatewayIsolationRuleResponseBodyDataResults {
	s.RouteId = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) SetRouteName(v string) *ListGatewayIsolationRuleResponseBodyDataResults {
	s.RouteName = &v
	return s
}

func (s *ListGatewayIsolationRuleResponseBodyDataResults) Validate() error {
	return dara.Validate(s)
}
