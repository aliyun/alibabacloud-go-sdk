// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApiImpactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDataServiceApiImpactsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListDataServiceApiImpactsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDataServiceApiImpactsResponseBody
	GetMessage() *string
	SetPageResult(v *ListDataServiceApiImpactsResponseBodyPageResult) *ListDataServiceApiImpactsResponseBody
	GetPageResult() *ListDataServiceApiImpactsResponseBodyPageResult
	SetRequestId(v string) *ListDataServiceApiImpactsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataServiceApiImpactsResponseBody
	GetSuccess() *bool
}

type ListDataServiceApiImpactsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message    *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListDataServiceApiImpactsResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataServiceApiImpactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiImpactsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiImpactsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDataServiceApiImpactsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDataServiceApiImpactsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDataServiceApiImpactsResponseBody) GetPageResult() *ListDataServiceApiImpactsResponseBodyPageResult {
	return s.PageResult
}

func (s *ListDataServiceApiImpactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataServiceApiImpactsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataServiceApiImpactsResponseBody) SetCode(v string) *ListDataServiceApiImpactsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBody) SetHttpStatusCode(v int32) *ListDataServiceApiImpactsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBody) SetMessage(v string) *ListDataServiceApiImpactsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBody) SetPageResult(v *ListDataServiceApiImpactsResponseBodyPageResult) *ListDataServiceApiImpactsResponseBody {
	s.PageResult = v
	return s
}

func (s *ListDataServiceApiImpactsResponseBody) SetRequestId(v string) *ListDataServiceApiImpactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBody) SetSuccess(v bool) *ListDataServiceApiImpactsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApiImpactsResponseBodyPageResult struct {
	ImpactList []*ListDataServiceApiImpactsResponseBodyPageResultImpactList `json:"ImpactList,omitempty" xml:"ImpactList,omitempty" type:"Repeated"`
	// example:
	//
	// 68
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataServiceApiImpactsResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiImpactsResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiImpactsResponseBodyPageResult) GetImpactList() []*ListDataServiceApiImpactsResponseBodyPageResultImpactList {
	return s.ImpactList
}

func (s *ListDataServiceApiImpactsResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataServiceApiImpactsResponseBodyPageResult) SetImpactList(v []*ListDataServiceApiImpactsResponseBodyPageResultImpactList) *ListDataServiceApiImpactsResponseBodyPageResult {
	s.ImpactList = v
	return s
}

func (s *ListDataServiceApiImpactsResponseBodyPageResult) SetTotalCount(v int32) *ListDataServiceApiImpactsResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBodyPageResult) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApiImpactsResponseBodyPageResultImpactList struct {
	// apiNo
	//
	// example:
	//
	// 2011
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// appKey
	//
	// example:
	//
	// 1101
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 11
	CallCount *int64 `json:"CallCount,omitempty" xml:"CallCount,omitempty"`
	// example:
	//
	// 2
	ClientFailCount *int64 `json:"ClientFailCount,omitempty" xml:"ClientFailCount,omitempty"`
	// example:
	//
	// 192.168.1.1
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// 1
	ErrorApiCount *int64 `json:"ErrorApiCount,omitempty" xml:"ErrorApiCount,omitempty"`
	// example:
	//
	// 1
	ErrorCount *int64 `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	// example:
	//
	// 10.0
	ErrorRate *string `json:"ErrorRate,omitempty" xml:"ErrorRate,omitempty"`
	// example:
	//
	// 2025-06-30 08:00:00
	LastCallTime *string `json:"LastCallTime,omitempty" xml:"LastCallTime,omitempty"`
	// example:
	//
	// 2025_0611_1011
	Minute *string `json:"Minute,omitempty" xml:"Minute,omitempty"`
	// example:
	//
	// 1
	OfflineCount *int64 `json:"OfflineCount,omitempty" xml:"OfflineCount,omitempty"`
	// example:
	//
	// 99
	SuccessTimeCost *string `json:"SuccessTimeCost,omitempty" xml:"SuccessTimeCost,omitempty"`
	// example:
	//
	// 88
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 231
	TotalTimeCost *string `json:"TotalTimeCost,omitempty" xml:"TotalTimeCost,omitempty"`
}

func (s ListDataServiceApiImpactsResponseBodyPageResultImpactList) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiImpactsResponseBodyPageResultImpactList) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) GetApiId() *int64 {
	return s.ApiId
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) GetAppKey() *int64 {
	return s.AppKey
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) GetAppName() *string {
	return s.AppName
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) GetCallCount() *int64 {
	return s.CallCount
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) GetClientFailCount() *int64 {
	return s.ClientFailCount
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) GetClientIp() *string {
	return s.ClientIp
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) GetErrorApiCount() *int64 {
	return s.ErrorApiCount
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) GetErrorCount() *int64 {
	return s.ErrorCount
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) GetErrorRate() *string {
	return s.ErrorRate
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) GetLastCallTime() *string {
	return s.LastCallTime
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) GetMinute() *string {
	return s.Minute
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) GetOfflineCount() *int64 {
	return s.OfflineCount
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) GetSuccessTimeCost() *string {
	return s.SuccessTimeCost
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) GetTotalTimeCost() *string {
	return s.TotalTimeCost
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) SetApiId(v int64) *ListDataServiceApiImpactsResponseBodyPageResultImpactList {
	s.ApiId = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) SetAppKey(v int64) *ListDataServiceApiImpactsResponseBodyPageResultImpactList {
	s.AppKey = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) SetAppName(v string) *ListDataServiceApiImpactsResponseBodyPageResultImpactList {
	s.AppName = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) SetCallCount(v int64) *ListDataServiceApiImpactsResponseBodyPageResultImpactList {
	s.CallCount = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) SetClientFailCount(v int64) *ListDataServiceApiImpactsResponseBodyPageResultImpactList {
	s.ClientFailCount = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) SetClientIp(v string) *ListDataServiceApiImpactsResponseBodyPageResultImpactList {
	s.ClientIp = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) SetErrorApiCount(v int64) *ListDataServiceApiImpactsResponseBodyPageResultImpactList {
	s.ErrorApiCount = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) SetErrorCount(v int64) *ListDataServiceApiImpactsResponseBodyPageResultImpactList {
	s.ErrorCount = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) SetErrorRate(v string) *ListDataServiceApiImpactsResponseBodyPageResultImpactList {
	s.ErrorRate = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) SetLastCallTime(v string) *ListDataServiceApiImpactsResponseBodyPageResultImpactList {
	s.LastCallTime = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) SetMinute(v string) *ListDataServiceApiImpactsResponseBodyPageResultImpactList {
	s.Minute = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) SetOfflineCount(v int64) *ListDataServiceApiImpactsResponseBodyPageResultImpactList {
	s.OfflineCount = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) SetSuccessTimeCost(v string) *ListDataServiceApiImpactsResponseBodyPageResultImpactList {
	s.SuccessTimeCost = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) SetTotalCount(v int64) *ListDataServiceApiImpactsResponseBodyPageResultImpactList {
	s.TotalCount = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) SetTotalTimeCost(v string) *ListDataServiceApiImpactsResponseBodyPageResultImpactList {
	s.TotalTimeCost = &v
	return s
}

func (s *ListDataServiceApiImpactsResponseBodyPageResultImpactList) Validate() error {
	return dara.Validate(s)
}
