// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLosslessRuleByAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetLosslessRuleByAppResponseBody
	GetCode() *int32
	SetData(v *GetLosslessRuleByAppResponseBodyData) *GetLosslessRuleByAppResponseBody
	GetData() *GetLosslessRuleByAppResponseBodyData
	SetErrorCode(v string) *GetLosslessRuleByAppResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *GetLosslessRuleByAppResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetLosslessRuleByAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLosslessRuleByAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLosslessRuleByAppResponseBody
	GetSuccess() *bool
}

type GetLosslessRuleByAppResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetLosslessRuleByAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned.
	//
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
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
	// The request ID.
	//
	// example:
	//
	// DC34E4A3-5F1C-4E40-86EA-02EDF967****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLosslessRuleByAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLosslessRuleByAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetLosslessRuleByAppResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetLosslessRuleByAppResponseBody) GetData() *GetLosslessRuleByAppResponseBodyData {
	return s.Data
}

func (s *GetLosslessRuleByAppResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetLosslessRuleByAppResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetLosslessRuleByAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLosslessRuleByAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLosslessRuleByAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLosslessRuleByAppResponseBody) SetCode(v int32) *GetLosslessRuleByAppResponseBody {
	s.Code = &v
	return s
}

func (s *GetLosslessRuleByAppResponseBody) SetData(v *GetLosslessRuleByAppResponseBodyData) *GetLosslessRuleByAppResponseBody {
	s.Data = v
	return s
}

func (s *GetLosslessRuleByAppResponseBody) SetErrorCode(v string) *GetLosslessRuleByAppResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetLosslessRuleByAppResponseBody) SetHttpStatusCode(v int32) *GetLosslessRuleByAppResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetLosslessRuleByAppResponseBody) SetMessage(v string) *GetLosslessRuleByAppResponseBody {
	s.Message = &v
	return s
}

func (s *GetLosslessRuleByAppResponseBody) SetRequestId(v string) *GetLosslessRuleByAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLosslessRuleByAppResponseBody) SetSuccess(v bool) *GetLosslessRuleByAppResponseBody {
	s.Success = &v
	return s
}

func (s *GetLosslessRuleByAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLosslessRuleByAppResponseBodyData struct {
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
	// hkhon1po62@24810bf4364a***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The number of instances.
	//
	// example:
	//
	// 1
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
	// true
	LossLessDetail *bool `json:"LossLessDetail,omitempty" xml:"LossLessDetail,omitempty"`
	// Indicates whether notification is enabled.
	//
	// example:
	//
	// false
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
	// 120
	WarmupTime *int32 `json:"WarmupTime,omitempty" xml:"WarmupTime,omitempty"`
}

func (s GetLosslessRuleByAppResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetLosslessRuleByAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLosslessRuleByAppResponseBodyData) GetAligned() *bool {
	return s.Aligned
}

func (s *GetLosslessRuleByAppResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *GetLosslessRuleByAppResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *GetLosslessRuleByAppResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *GetLosslessRuleByAppResponseBodyData) GetDelayTime() *int32 {
	return s.DelayTime
}

func (s *GetLosslessRuleByAppResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *GetLosslessRuleByAppResponseBodyData) GetFuncType() *int32 {
	return s.FuncType
}

func (s *GetLosslessRuleByAppResponseBodyData) GetLossLessDetail() *bool {
	return s.LossLessDetail
}

func (s *GetLosslessRuleByAppResponseBodyData) GetNotice() *bool {
	return s.Notice
}

func (s *GetLosslessRuleByAppResponseBodyData) GetRelated() *bool {
	return s.Related
}

func (s *GetLosslessRuleByAppResponseBodyData) GetWarmupTime() *int32 {
	return s.WarmupTime
}

func (s *GetLosslessRuleByAppResponseBodyData) SetAligned(v bool) *GetLosslessRuleByAppResponseBodyData {
	s.Aligned = &v
	return s
}

func (s *GetLosslessRuleByAppResponseBodyData) SetAppId(v string) *GetLosslessRuleByAppResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetLosslessRuleByAppResponseBodyData) SetAppName(v string) *GetLosslessRuleByAppResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetLosslessRuleByAppResponseBodyData) SetCount(v int32) *GetLosslessRuleByAppResponseBodyData {
	s.Count = &v
	return s
}

func (s *GetLosslessRuleByAppResponseBodyData) SetDelayTime(v int32) *GetLosslessRuleByAppResponseBodyData {
	s.DelayTime = &v
	return s
}

func (s *GetLosslessRuleByAppResponseBodyData) SetEnable(v bool) *GetLosslessRuleByAppResponseBodyData {
	s.Enable = &v
	return s
}

func (s *GetLosslessRuleByAppResponseBodyData) SetFuncType(v int32) *GetLosslessRuleByAppResponseBodyData {
	s.FuncType = &v
	return s
}

func (s *GetLosslessRuleByAppResponseBodyData) SetLossLessDetail(v bool) *GetLosslessRuleByAppResponseBodyData {
	s.LossLessDetail = &v
	return s
}

func (s *GetLosslessRuleByAppResponseBodyData) SetNotice(v bool) *GetLosslessRuleByAppResponseBodyData {
	s.Notice = &v
	return s
}

func (s *GetLosslessRuleByAppResponseBodyData) SetRelated(v bool) *GetLosslessRuleByAppResponseBodyData {
	s.Related = &v
	return s
}

func (s *GetLosslessRuleByAppResponseBodyData) SetWarmupTime(v int32) *GetLosslessRuleByAppResponseBodyData {
	s.WarmupTime = &v
	return s
}

func (s *GetLosslessRuleByAppResponseBodyData) Validate() error {
	return dara.Validate(s)
}
