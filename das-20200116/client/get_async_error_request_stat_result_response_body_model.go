// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncErrorRequestStatResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetAsyncErrorRequestStatResultResponseBody
	GetCode() *int64
	SetData(v *GetAsyncErrorRequestStatResultResponseBodyData) *GetAsyncErrorRequestStatResultResponseBody
	GetData() *GetAsyncErrorRequestStatResultResponseBodyData
	SetMessage(v string) *GetAsyncErrorRequestStatResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAsyncErrorRequestStatResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAsyncErrorRequestStatResultResponseBody
	GetSuccess() *bool
}

type GetAsyncErrorRequestStatResultResponseBody struct {
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
	// {         "fail": false,         "data": {             "ad78a4e7d3ce81590c9dc2d5f4bc****": {                 "sqlId": "ad78a4e7d3ce81590c9dc2d5f4bc****",                 "instanceId": "rm-2ze8g2am97624****",                 "count": 1             },             "0f92feacd92c048b06a16617a633****": {                 "sqlId": "0f92feacd92c048b06a16617a633****",                 "instanceId": "rm-2ze8g2am97624****",                 "count": 2             }         },         "resultId": "async__61f45ee381b2fa4e8a6545e3bee9****",         "isFinish": true,         "state": "SUCCESS",         "complete": true,         "timestamp": 1644558576717     }
	Data *GetAsyncErrorRequestStatResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 3FC3F8EB-3564-5D1A-B187-3B03E5B0****
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

func (s GetAsyncErrorRequestStatResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncErrorRequestStatResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestStatResultResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetAsyncErrorRequestStatResultResponseBody) GetData() *GetAsyncErrorRequestStatResultResponseBodyData {
	return s.Data
}

func (s *GetAsyncErrorRequestStatResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAsyncErrorRequestStatResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAsyncErrorRequestStatResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAsyncErrorRequestStatResultResponseBody) SetCode(v int64) *GetAsyncErrorRequestStatResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultResponseBody) SetData(v *GetAsyncErrorRequestStatResultResponseBodyData) *GetAsyncErrorRequestStatResultResponseBody {
	s.Data = v
	return s
}

func (s *GetAsyncErrorRequestStatResultResponseBody) SetMessage(v string) *GetAsyncErrorRequestStatResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultResponseBody) SetRequestId(v string) *GetAsyncErrorRequestStatResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultResponseBody) SetSuccess(v bool) *GetAsyncErrorRequestStatResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAsyncErrorRequestStatResultResponseBodyData struct {
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
	// Indicates whether the request failed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Fail *bool `json:"fail,omitempty" xml:"fail,omitempty"`
	// Indicates whether the asynchronous request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsFinish *bool `json:"isFinish,omitempty" xml:"isFinish,omitempty"`
	// The returned data of the asynchronous request.
	Result map[string]*DataResultValue `json:"result,omitempty" xml:"result,omitempty"`
	// The ID of the asynchronous request.
	//
	// example:
	//
	// async__61f45ee381b2fa4e8a6545e3bee9****
	ResultId *string `json:"resultId,omitempty" xml:"resultId,omitempty"`
	// The state of the asynchronous request. Valid values:
	//
	// 	- **RUNNING**: The asynchronous request is running.
	//
	// 	- **SUCCESS**: The asynchronous request is successful.
	//
	// 	- **FAIL**: The asynchronous request fails.
	//
	// example:
	//
	// SUCCESS
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The time when the asynchronous request was made. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1644558576717
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s GetAsyncErrorRequestStatResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncErrorRequestStatResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestStatResultResponseBodyData) GetComplete() *bool {
	return s.Complete
}

func (s *GetAsyncErrorRequestStatResultResponseBodyData) GetFail() *bool {
	return s.Fail
}

func (s *GetAsyncErrorRequestStatResultResponseBodyData) GetIsFinish() *bool {
	return s.IsFinish
}

func (s *GetAsyncErrorRequestStatResultResponseBodyData) GetResult() map[string]*DataResultValue {
	return s.Result
}

func (s *GetAsyncErrorRequestStatResultResponseBodyData) GetResultId() *string {
	return s.ResultId
}

func (s *GetAsyncErrorRequestStatResultResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetAsyncErrorRequestStatResultResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetAsyncErrorRequestStatResultResponseBodyData) SetComplete(v bool) *GetAsyncErrorRequestStatResultResponseBodyData {
	s.Complete = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultResponseBodyData) SetFail(v bool) *GetAsyncErrorRequestStatResultResponseBodyData {
	s.Fail = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultResponseBodyData) SetIsFinish(v bool) *GetAsyncErrorRequestStatResultResponseBodyData {
	s.IsFinish = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultResponseBodyData) SetResult(v map[string]*DataResultValue) *GetAsyncErrorRequestStatResultResponseBodyData {
	s.Result = v
	return s
}

func (s *GetAsyncErrorRequestStatResultResponseBodyData) SetResultId(v string) *GetAsyncErrorRequestStatResultResponseBodyData {
	s.ResultId = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultResponseBodyData) SetState(v string) *GetAsyncErrorRequestStatResultResponseBodyData {
	s.State = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultResponseBodyData) SetTimestamp(v int64) *GetAsyncErrorRequestStatResultResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
