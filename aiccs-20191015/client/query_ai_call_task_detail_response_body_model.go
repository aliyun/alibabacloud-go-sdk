// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAiCallTaskDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryAiCallTaskDetailResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryAiCallTaskDetailResponseBody
	GetCode() *string
	SetData(v *QueryAiCallTaskDetailResponseBodyData) *QueryAiCallTaskDetailResponseBody
	GetData() *QueryAiCallTaskDetailResponseBodyData
	SetMessage(v string) *QueryAiCallTaskDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAiCallTaskDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAiCallTaskDetailResponseBody
	GetSuccess() *bool
}

type QueryAiCallTaskDetailResponseBody struct {
	// example:
	//
	// Access Denied
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryAiCallTaskDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 23822ECB-8CAA-5C52-9C9E-807FD82A5A7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAiCallTaskDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAiCallTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAiCallTaskDetailResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryAiCallTaskDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAiCallTaskDetailResponseBody) GetData() *QueryAiCallTaskDetailResponseBodyData {
	return s.Data
}

func (s *QueryAiCallTaskDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAiCallTaskDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAiCallTaskDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAiCallTaskDetailResponseBody) SetAccessDeniedDetail(v string) *QueryAiCallTaskDetailResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBody) SetCode(v string) *QueryAiCallTaskDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBody) SetData(v *QueryAiCallTaskDetailResponseBodyData) *QueryAiCallTaskDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryAiCallTaskDetailResponseBody) SetMessage(v string) *QueryAiCallTaskDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBody) SetRequestId(v string) *QueryAiCallTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBody) SetSuccess(v bool) *QueryAiCallTaskDetailResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAiCallTaskDetailResponseBodyData struct {
	// example:
	//
	// 示例值示例值
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 示例值示例值
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// example:
	//
	// 示例值
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	// example:
	//
	// 示例值
	ApplicationName *string                                           `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	CallDays        []*string                                         `json:"CallDays,omitempty" xml:"CallDays,omitempty" type:"Repeated"`
	CallTimes       []*QueryAiCallTaskDetailResponseBodyDataCallTimes `json:"CallTimes,omitempty" xml:"CallTimes,omitempty" type:"Repeated"`
	// example:
	//
	// 05370124****
	CallerNumber *string `json:"CallerNumber,omitempty" xml:"CallerNumber,omitempty"`
	// example:
	//
	// 10
	ConcurrentCount *int64 `json:"ConcurrentCount,omitempty" xml:"ConcurrentCount,omitempty"`
	// example:
	//
	// 示例值示例值
	LineEncoding *string `json:"LineEncoding,omitempty" xml:"LineEncoding,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	LinePhoneNum *string `json:"LinePhoneNum,omitempty" xml:"LinePhoneNum,omitempty"`
	// example:
	//
	// 17
	PhoneType *int64 `json:"PhoneType,omitempty" xml:"PhoneType,omitempty"`
	// example:
	//
	// 1748932499000
	RealStartTime *int64 `json:"RealStartTime,omitempty" xml:"RealStartTime,omitempty"`
	// example:
	//
	// 2
	RetryCount *int64 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	// example:
	//
	// true
	RetryEnable *bool `json:"RetryEnable,omitempty" xml:"RetryEnable,omitempty"`
	// example:
	//
	// 1
	RetryInterval *int64    `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	RetryReasons  []*string `json:"RetryReasons,omitempty" xml:"RetryReasons,omitempty" type:"Repeated"`
	// example:
	//
	// 1748932499000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// IMMEDIATE
	StartType *string `json:"StartType,omitempty" xml:"StartType,omitempty"`
	// example:
	//
	// 11121232222****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 示例值示例值
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s QueryAiCallTaskDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAiCallTaskDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAiCallTaskDetailResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *QueryAiCallTaskDetailResponseBodyData) GetAgentName() *string {
	return s.AgentName
}

func (s *QueryAiCallTaskDetailResponseBodyData) GetApplicationCode() *string {
	return s.ApplicationCode
}

func (s *QueryAiCallTaskDetailResponseBodyData) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *QueryAiCallTaskDetailResponseBodyData) GetCallDays() []*string {
	return s.CallDays
}

func (s *QueryAiCallTaskDetailResponseBodyData) GetCallTimes() []*QueryAiCallTaskDetailResponseBodyDataCallTimes {
	return s.CallTimes
}

func (s *QueryAiCallTaskDetailResponseBodyData) GetCallerNumber() *string {
	return s.CallerNumber
}

func (s *QueryAiCallTaskDetailResponseBodyData) GetConcurrentCount() *int64 {
	return s.ConcurrentCount
}

func (s *QueryAiCallTaskDetailResponseBodyData) GetLineEncoding() *string {
	return s.LineEncoding
}

func (s *QueryAiCallTaskDetailResponseBodyData) GetLinePhoneNum() *string {
	return s.LinePhoneNum
}

func (s *QueryAiCallTaskDetailResponseBodyData) GetPhoneType() *int64 {
	return s.PhoneType
}

func (s *QueryAiCallTaskDetailResponseBodyData) GetRealStartTime() *int64 {
	return s.RealStartTime
}

func (s *QueryAiCallTaskDetailResponseBodyData) GetRetryCount() *int64 {
	return s.RetryCount
}

func (s *QueryAiCallTaskDetailResponseBodyData) GetRetryEnable() *bool {
	return s.RetryEnable
}

func (s *QueryAiCallTaskDetailResponseBodyData) GetRetryInterval() *int64 {
	return s.RetryInterval
}

func (s *QueryAiCallTaskDetailResponseBodyData) GetRetryReasons() []*string {
	return s.RetryReasons
}

func (s *QueryAiCallTaskDetailResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryAiCallTaskDetailResponseBodyData) GetStartType() *string {
	return s.StartType
}

func (s *QueryAiCallTaskDetailResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryAiCallTaskDetailResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *QueryAiCallTaskDetailResponseBodyData) SetAgentId(v string) *QueryAiCallTaskDetailResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyData) SetAgentName(v string) *QueryAiCallTaskDetailResponseBodyData {
	s.AgentName = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyData) SetApplicationCode(v string) *QueryAiCallTaskDetailResponseBodyData {
	s.ApplicationCode = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyData) SetApplicationName(v string) *QueryAiCallTaskDetailResponseBodyData {
	s.ApplicationName = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyData) SetCallDays(v []*string) *QueryAiCallTaskDetailResponseBodyData {
	s.CallDays = v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyData) SetCallTimes(v []*QueryAiCallTaskDetailResponseBodyDataCallTimes) *QueryAiCallTaskDetailResponseBodyData {
	s.CallTimes = v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyData) SetCallerNumber(v string) *QueryAiCallTaskDetailResponseBodyData {
	s.CallerNumber = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyData) SetConcurrentCount(v int64) *QueryAiCallTaskDetailResponseBodyData {
	s.ConcurrentCount = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyData) SetLineEncoding(v string) *QueryAiCallTaskDetailResponseBodyData {
	s.LineEncoding = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyData) SetLinePhoneNum(v string) *QueryAiCallTaskDetailResponseBodyData {
	s.LinePhoneNum = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyData) SetPhoneType(v int64) *QueryAiCallTaskDetailResponseBodyData {
	s.PhoneType = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyData) SetRealStartTime(v int64) *QueryAiCallTaskDetailResponseBodyData {
	s.RealStartTime = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyData) SetRetryCount(v int64) *QueryAiCallTaskDetailResponseBodyData {
	s.RetryCount = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyData) SetRetryEnable(v bool) *QueryAiCallTaskDetailResponseBodyData {
	s.RetryEnable = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyData) SetRetryInterval(v int64) *QueryAiCallTaskDetailResponseBodyData {
	s.RetryInterval = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyData) SetRetryReasons(v []*string) *QueryAiCallTaskDetailResponseBodyData {
	s.RetryReasons = v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyData) SetStartTime(v int64) *QueryAiCallTaskDetailResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyData) SetStartType(v string) *QueryAiCallTaskDetailResponseBodyData {
	s.StartType = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyData) SetTaskId(v string) *QueryAiCallTaskDetailResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyData) SetTaskName(v string) *QueryAiCallTaskDetailResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyData) Validate() error {
	if s.CallTimes != nil {
		for _, item := range s.CallTimes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAiCallTaskDetailResponseBodyDataCallTimes struct {
	// example:
	//
	// 09:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 12:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryAiCallTaskDetailResponseBodyDataCallTimes) String() string {
	return dara.Prettify(s)
}

func (s QueryAiCallTaskDetailResponseBodyDataCallTimes) GoString() string {
	return s.String()
}

func (s *QueryAiCallTaskDetailResponseBodyDataCallTimes) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryAiCallTaskDetailResponseBodyDataCallTimes) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryAiCallTaskDetailResponseBodyDataCallTimes) SetEndTime(v string) *QueryAiCallTaskDetailResponseBodyDataCallTimes {
	s.EndTime = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyDataCallTimes) SetStartTime(v string) *QueryAiCallTaskDetailResponseBodyDataCallTimes {
	s.StartTime = &v
	return s
}

func (s *QueryAiCallTaskDetailResponseBodyDataCallTimes) Validate() error {
	return dara.Validate(s)
}
