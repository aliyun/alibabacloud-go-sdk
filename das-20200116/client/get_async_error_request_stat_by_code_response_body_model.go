// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncErrorRequestStatByCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetAsyncErrorRequestStatByCodeResponseBody
	GetCode() *int64
	SetData(v *GetAsyncErrorRequestStatByCodeResponseBodyData) *GetAsyncErrorRequestStatByCodeResponseBody
	GetData() *GetAsyncErrorRequestStatByCodeResponseBodyData
	SetMessage(v string) *GetAsyncErrorRequestStatByCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAsyncErrorRequestStatByCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAsyncErrorRequestStatByCodeResponseBody
	GetSuccess() *bool
}

type GetAsyncErrorRequestStatByCodeResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// {     "fail": false,     "data": [       {         "instanceId": "rm-2ze8g2am97624****",         "count": 1,         "errorCode": "1062"       },       {         "instanceId": "rm-2ze8g2am97624****",         "count": 2,         "errorCode": "1064"      }     ],     "resultId": "async__fcd7c35788e62324622c3b4a03de****",     "isFinish": true,     "state": "SUCCESS",     "complete": true,     "timestamp": 1644560866961   }
	Data *GetAsyncErrorRequestStatByCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 840F51F7-9C01-538D-94F6-AE712905****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAsyncErrorRequestStatByCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncErrorRequestStatByCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestStatByCodeResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetAsyncErrorRequestStatByCodeResponseBody) GetData() *GetAsyncErrorRequestStatByCodeResponseBodyData {
	return s.Data
}

func (s *GetAsyncErrorRequestStatByCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAsyncErrorRequestStatByCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAsyncErrorRequestStatByCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAsyncErrorRequestStatByCodeResponseBody) SetCode(v int64) *GetAsyncErrorRequestStatByCodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponseBody) SetData(v *GetAsyncErrorRequestStatByCodeResponseBodyData) *GetAsyncErrorRequestStatByCodeResponseBody {
	s.Data = v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponseBody) SetMessage(v string) *GetAsyncErrorRequestStatByCodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponseBody) SetRequestId(v string) *GetAsyncErrorRequestStatByCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponseBody) SetSuccess(v bool) *GetAsyncErrorRequestStatByCodeResponseBody {
	s.Success = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAsyncErrorRequestStatByCodeResponseBodyData struct {
	// Indicates whether the asynchronous request was complete.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Complete *bool `json:"complete,omitempty" xml:"complete,omitempty"`
	// Indicates whether the asynchronous request failed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Fail *bool `json:"fail,omitempty" xml:"fail,omitempty"`
	// Indicates whether the asynchronous request was complete. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsFinish *bool `json:"isFinish,omitempty" xml:"isFinish,omitempty"`
	// The number of SQL queries corresponding to the error code.
	Result []*GetAsyncErrorRequestStatByCodeResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The ID of the asynchronous request.
	//
	// example:
	//
	// async__fcd7c35788e62324622c3b4a03de****
	ResultId *string `json:"resultId,omitempty" xml:"resultId,omitempty"`
	// The state of the asynchronous request. Valid values:
	//
	// 	- **RUNNING**
	//
	// 	- **SUCCESS**
	//
	// 	- **FAIL**
	//
	// example:
	//
	// SUCCESS
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The time when the asynchronous request was made. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1644560866961
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s GetAsyncErrorRequestStatByCodeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncErrorRequestStatByCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyData) GetComplete() *bool {
	return s.Complete
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyData) GetFail() *bool {
	return s.Fail
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyData) GetIsFinish() *bool {
	return s.IsFinish
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyData) GetResult() []*GetAsyncErrorRequestStatByCodeResponseBodyDataResult {
	return s.Result
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyData) GetResultId() *string {
	return s.ResultId
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyData) SetComplete(v bool) *GetAsyncErrorRequestStatByCodeResponseBodyData {
	s.Complete = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyData) SetFail(v bool) *GetAsyncErrorRequestStatByCodeResponseBodyData {
	s.Fail = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyData) SetIsFinish(v bool) *GetAsyncErrorRequestStatByCodeResponseBodyData {
	s.IsFinish = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyData) SetResult(v []*GetAsyncErrorRequestStatByCodeResponseBodyDataResult) *GetAsyncErrorRequestStatByCodeResponseBodyData {
	s.Result = v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyData) SetResultId(v string) *GetAsyncErrorRequestStatByCodeResponseBodyData {
	s.ResultId = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyData) SetState(v string) *GetAsyncErrorRequestStatByCodeResponseBodyData {
	s.State = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyData) SetTimestamp(v int64) *GetAsyncErrorRequestStatByCodeResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetAsyncErrorRequestStatByCodeResponseBodyDataResult struct {
	// The number of SQL queries corresponding to the error code.
	//
	// example:
	//
	// 1
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// 1062
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s GetAsyncErrorRequestStatByCodeResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncErrorRequestStatByCodeResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyDataResult) GetCount() *int32 {
	return s.Count
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyDataResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyDataResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyDataResult) SetCount(v int32) *GetAsyncErrorRequestStatByCodeResponseBodyDataResult {
	s.Count = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyDataResult) SetErrorCode(v string) *GetAsyncErrorRequestStatByCodeResponseBodyDataResult {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyDataResult) SetInstanceId(v string) *GetAsyncErrorRequestStatByCodeResponseBodyDataResult {
	s.InstanceId = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
