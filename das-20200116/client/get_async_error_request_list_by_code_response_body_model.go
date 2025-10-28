// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncErrorRequestListByCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetAsyncErrorRequestListByCodeResponseBody
	GetCode() *int64
	SetData(v *GetAsyncErrorRequestListByCodeResponseBodyData) *GetAsyncErrorRequestListByCodeResponseBody
	GetData() *GetAsyncErrorRequestListByCodeResponseBodyData
	SetMessage(v string) *GetAsyncErrorRequestListByCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAsyncErrorRequestListByCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAsyncErrorRequestListByCodeResponseBody
	GetSuccess() *bool
}

type GetAsyncErrorRequestListByCodeResponseBody struct {
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
	// {         "fail": false,         "data": [             {                 "sqlId": "ad78a4e7d3ce81590c9dc2d5f4bc****",                 "instanceId": "rm-2ze8g2am97624****"             },             {                 "sqlId": "0f92feacd92c048b06a16617a633****",                 "instanceId": "rm-2ze8g2am97624****"             }         ],         "resultId": "async__c39d43ece52d35267cc4b92a0c26****",         "isFinish": true,         "state": "SUCCESS",         "complete": true,         "timestamp": 1644559407740     }
	Data *GetAsyncErrorRequestListByCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 9CB97BC4-6479-55D0-B9D0-EA925AFE****
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

func (s GetAsyncErrorRequestListByCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncErrorRequestListByCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestListByCodeResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetAsyncErrorRequestListByCodeResponseBody) GetData() *GetAsyncErrorRequestListByCodeResponseBodyData {
	return s.Data
}

func (s *GetAsyncErrorRequestListByCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAsyncErrorRequestListByCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAsyncErrorRequestListByCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAsyncErrorRequestListByCodeResponseBody) SetCode(v int64) *GetAsyncErrorRequestListByCodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponseBody) SetData(v *GetAsyncErrorRequestListByCodeResponseBodyData) *GetAsyncErrorRequestListByCodeResponseBody {
	s.Data = v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponseBody) SetMessage(v string) *GetAsyncErrorRequestListByCodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponseBody) SetRequestId(v string) *GetAsyncErrorRequestListByCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponseBody) SetSuccess(v bool) *GetAsyncErrorRequestListByCodeResponseBody {
	s.Success = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAsyncErrorRequestListByCodeResponseBodyData struct {
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
	// The instance ID.
	Result []*GetAsyncErrorRequestListByCodeResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The ID of the asynchronous request.
	//
	// example:
	//
	// async__c39d43ece52d35267cc4b92a0c26****
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
	// 1644559407740
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s GetAsyncErrorRequestListByCodeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncErrorRequestListByCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestListByCodeResponseBodyData) GetComplete() *bool {
	return s.Complete
}

func (s *GetAsyncErrorRequestListByCodeResponseBodyData) GetFail() *bool {
	return s.Fail
}

func (s *GetAsyncErrorRequestListByCodeResponseBodyData) GetIsFinish() *bool {
	return s.IsFinish
}

func (s *GetAsyncErrorRequestListByCodeResponseBodyData) GetResult() []*GetAsyncErrorRequestListByCodeResponseBodyDataResult {
	return s.Result
}

func (s *GetAsyncErrorRequestListByCodeResponseBodyData) GetResultId() *string {
	return s.ResultId
}

func (s *GetAsyncErrorRequestListByCodeResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetAsyncErrorRequestListByCodeResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetAsyncErrorRequestListByCodeResponseBodyData) SetComplete(v bool) *GetAsyncErrorRequestListByCodeResponseBodyData {
	s.Complete = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponseBodyData) SetFail(v bool) *GetAsyncErrorRequestListByCodeResponseBodyData {
	s.Fail = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponseBodyData) SetIsFinish(v bool) *GetAsyncErrorRequestListByCodeResponseBodyData {
	s.IsFinish = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponseBodyData) SetResult(v []*GetAsyncErrorRequestListByCodeResponseBodyDataResult) *GetAsyncErrorRequestListByCodeResponseBodyData {
	s.Result = v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponseBodyData) SetResultId(v string) *GetAsyncErrorRequestListByCodeResponseBodyData {
	s.ResultId = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponseBodyData) SetState(v string) *GetAsyncErrorRequestListByCodeResponseBodyData {
	s.State = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponseBodyData) SetTimestamp(v int64) *GetAsyncErrorRequestListByCodeResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAsyncErrorRequestListByCodeResponseBodyDataResult struct {
	// The instance ID
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// SQL ID.
	//
	// example:
	//
	// ad78a4e7d3ce81590c9dc2d5f4bc****
	SqlId *string `json:"sqlId,omitempty" xml:"sqlId,omitempty"`
}

func (s GetAsyncErrorRequestListByCodeResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncErrorRequestListByCodeResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestListByCodeResponseBodyDataResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAsyncErrorRequestListByCodeResponseBodyDataResult) GetSqlId() *string {
	return s.SqlId
}

func (s *GetAsyncErrorRequestListByCodeResponseBodyDataResult) SetInstanceId(v string) *GetAsyncErrorRequestListByCodeResponseBodyDataResult {
	s.InstanceId = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponseBodyDataResult) SetSqlId(v string) *GetAsyncErrorRequestListByCodeResponseBodyDataResult {
	s.SqlId = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
