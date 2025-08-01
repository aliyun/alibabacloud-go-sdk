// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchLosslessRuleListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *FetchLosslessRuleListResponseBody
	GetCode() *int32
	SetData(v *FetchLosslessRuleListResponseBodyData) *FetchLosslessRuleListResponseBody
	GetData() *FetchLosslessRuleListResponseBodyData
	SetErrorCode(v string) *FetchLosslessRuleListResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *FetchLosslessRuleListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *FetchLosslessRuleListResponseBody
	GetMessage() *string
	SetRequestId(v string) *FetchLosslessRuleListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FetchLosslessRuleListResponseBody
	GetSuccess() *bool
}

type FetchLosslessRuleListResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *FetchLosslessRuleListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 56D9E600-6348-4260-B35F-583413F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FetchLosslessRuleListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FetchLosslessRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *FetchLosslessRuleListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *FetchLosslessRuleListResponseBody) GetData() *FetchLosslessRuleListResponseBodyData {
	return s.Data
}

func (s *FetchLosslessRuleListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *FetchLosslessRuleListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *FetchLosslessRuleListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FetchLosslessRuleListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FetchLosslessRuleListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FetchLosslessRuleListResponseBody) SetCode(v int32) *FetchLosslessRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *FetchLosslessRuleListResponseBody) SetData(v *FetchLosslessRuleListResponseBodyData) *FetchLosslessRuleListResponseBody {
	s.Data = v
	return s
}

func (s *FetchLosslessRuleListResponseBody) SetErrorCode(v string) *FetchLosslessRuleListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *FetchLosslessRuleListResponseBody) SetHttpStatusCode(v int32) *FetchLosslessRuleListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *FetchLosslessRuleListResponseBody) SetMessage(v string) *FetchLosslessRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *FetchLosslessRuleListResponseBody) SetRequestId(v string) *FetchLosslessRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchLosslessRuleListResponseBody) SetSuccess(v bool) *FetchLosslessRuleListResponseBody {
	s.Success = &v
	return s
}

func (s *FetchLosslessRuleListResponseBody) Validate() error {
	return dara.Validate(s)
}

type FetchLosslessRuleListResponseBodyData struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The returned data.
	Results []*FetchLosslessRuleListResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 36
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s FetchLosslessRuleListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FetchLosslessRuleListResponseBodyData) GoString() string {
	return s.String()
}

func (s *FetchLosslessRuleListResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *FetchLosslessRuleListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *FetchLosslessRuleListResponseBodyData) GetResults() []*FetchLosslessRuleListResponseBodyDataResults {
	return s.Results
}

func (s *FetchLosslessRuleListResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *FetchLosslessRuleListResponseBodyData) SetPageNumber(v int32) *FetchLosslessRuleListResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *FetchLosslessRuleListResponseBodyData) SetPageSize(v int32) *FetchLosslessRuleListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *FetchLosslessRuleListResponseBodyData) SetResults(v []*FetchLosslessRuleListResponseBodyDataResults) *FetchLosslessRuleListResponseBodyData {
	s.Results = v
	return s
}

func (s *FetchLosslessRuleListResponseBodyData) SetTotalSize(v int32) *FetchLosslessRuleListResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *FetchLosslessRuleListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type FetchLosslessRuleListResponseBodyDataResults struct {
	// Indicates whether service registration is complete before readiness probe.
	//
	// example:
	//
	// true
	Aligned *bool `json:"Aligned,omitempty" xml:"Aligned,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// hkhon1po62@24810bf4364aea1
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// echo-demo
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The number of instances.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The registration latency. Unit: seconds.
	//
	// example:
	//
	// 60
	DelayTime *int32 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// Indicates whether graceful start is enabled. Valid values:
	//
	// 	- `true`: enabled
	//
	// 	- `false`: disabled
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The slope of the prefetching curve.
	//
	// example:
	//
	// 2
	FuncType *int32 `json:"FuncType,omitempty" xml:"FuncType,omitempty"`
	// Indicates whether online and offline processing details are displayed.
	//
	// example:
	//
	// false
	LossLessDetail *bool `json:"LossLessDetail,omitempty" xml:"LossLessDetail,omitempty"`
	// Indicates whether notification is enabled.
	//
	// example:
	//
	// true
	Notice *bool `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// Indicates whether service prefetching is complete before readiness probe.
	//
	// example:
	//
	// false
	Related *bool `json:"Related,omitempty" xml:"Related,omitempty"`
	// The prefetching duration. Unit: seconds.
	//
	// example:
	//
	// 60
	WarmupTime *int32 `json:"WarmupTime,omitempty" xml:"WarmupTime,omitempty"`
}

func (s FetchLosslessRuleListResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s FetchLosslessRuleListResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *FetchLosslessRuleListResponseBodyDataResults) GetAligned() *bool {
	return s.Aligned
}

func (s *FetchLosslessRuleListResponseBodyDataResults) GetAppId() *string {
	return s.AppId
}

func (s *FetchLosslessRuleListResponseBodyDataResults) GetAppName() *string {
	return s.AppName
}

func (s *FetchLosslessRuleListResponseBodyDataResults) GetCount() *int32 {
	return s.Count
}

func (s *FetchLosslessRuleListResponseBodyDataResults) GetDelayTime() *int32 {
	return s.DelayTime
}

func (s *FetchLosslessRuleListResponseBodyDataResults) GetEnable() *bool {
	return s.Enable
}

func (s *FetchLosslessRuleListResponseBodyDataResults) GetFuncType() *int32 {
	return s.FuncType
}

func (s *FetchLosslessRuleListResponseBodyDataResults) GetLossLessDetail() *bool {
	return s.LossLessDetail
}

func (s *FetchLosslessRuleListResponseBodyDataResults) GetNotice() *bool {
	return s.Notice
}

func (s *FetchLosslessRuleListResponseBodyDataResults) GetRelated() *bool {
	return s.Related
}

func (s *FetchLosslessRuleListResponseBodyDataResults) GetWarmupTime() *int32 {
	return s.WarmupTime
}

func (s *FetchLosslessRuleListResponseBodyDataResults) SetAligned(v bool) *FetchLosslessRuleListResponseBodyDataResults {
	s.Aligned = &v
	return s
}

func (s *FetchLosslessRuleListResponseBodyDataResults) SetAppId(v string) *FetchLosslessRuleListResponseBodyDataResults {
	s.AppId = &v
	return s
}

func (s *FetchLosslessRuleListResponseBodyDataResults) SetAppName(v string) *FetchLosslessRuleListResponseBodyDataResults {
	s.AppName = &v
	return s
}

func (s *FetchLosslessRuleListResponseBodyDataResults) SetCount(v int32) *FetchLosslessRuleListResponseBodyDataResults {
	s.Count = &v
	return s
}

func (s *FetchLosslessRuleListResponseBodyDataResults) SetDelayTime(v int32) *FetchLosslessRuleListResponseBodyDataResults {
	s.DelayTime = &v
	return s
}

func (s *FetchLosslessRuleListResponseBodyDataResults) SetEnable(v bool) *FetchLosslessRuleListResponseBodyDataResults {
	s.Enable = &v
	return s
}

func (s *FetchLosslessRuleListResponseBodyDataResults) SetFuncType(v int32) *FetchLosslessRuleListResponseBodyDataResults {
	s.FuncType = &v
	return s
}

func (s *FetchLosslessRuleListResponseBodyDataResults) SetLossLessDetail(v bool) *FetchLosslessRuleListResponseBodyDataResults {
	s.LossLessDetail = &v
	return s
}

func (s *FetchLosslessRuleListResponseBodyDataResults) SetNotice(v bool) *FetchLosslessRuleListResponseBodyDataResults {
	s.Notice = &v
	return s
}

func (s *FetchLosslessRuleListResponseBodyDataResults) SetRelated(v bool) *FetchLosslessRuleListResponseBodyDataResults {
	s.Related = &v
	return s
}

func (s *FetchLosslessRuleListResponseBodyDataResults) SetWarmupTime(v int32) *FetchLosslessRuleListResponseBodyDataResults {
	s.WarmupTime = &v
	return s
}

func (s *FetchLosslessRuleListResponseBodyDataResults) Validate() error {
	return dara.Validate(s)
}
