// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeJobStatusResponseBody
	GetCode() *string
	SetData(v *DescribeJobStatusResponseBodyData) *DescribeJobStatusResponseBody
	GetData() *DescribeJobStatusResponseBodyData
	SetErrorCode(v string) *DescribeJobStatusResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeJobStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeJobStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeJobStatusResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeJobStatusResponseBody
	GetTraceId() *string
}

type DescribeJobStatusResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DescribeJobStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned. Take note of the following rules:
	//
	// 	- **ErrorCode*	- is not returned if the request succeeds.
	//
	// 	- **ErrorCode*	- is returned if the request fails. For more information, see the "**Error codes**" section in this topic.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message. Take note of the following rules:
	//
	// 	- If the call is successful, **success*	- is returned.
	//
	// 	- If the call fails, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 53F15A18-8079-5992-810C-0211A5AE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of the request.
	//
	// example:
	//
	// 0b1639af16575057857241351e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeJobStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeJobStatusResponseBody) GetData() *DescribeJobStatusResponseBodyData {
	return s.Data
}

func (s *DescribeJobStatusResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeJobStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeJobStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeJobStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeJobStatusResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeJobStatusResponseBody) SetCode(v string) *DescribeJobStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeJobStatusResponseBody) SetData(v *DescribeJobStatusResponseBodyData) *DescribeJobStatusResponseBody {
	s.Data = v
	return s
}

func (s *DescribeJobStatusResponseBody) SetErrorCode(v string) *DescribeJobStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeJobStatusResponseBody) SetMessage(v string) *DescribeJobStatusResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeJobStatusResponseBody) SetRequestId(v string) *DescribeJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJobStatusResponseBody) SetSuccess(v bool) *DescribeJobStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeJobStatusResponseBody) SetTraceId(v string) *DescribeJobStatusResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeJobStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeJobStatusResponseBodyData struct {
	// The number of running instances.
	//
	// example:
	//
	// 0
	Active *int64 `json:"Active,omitempty" xml:"Active,omitempty"`
	// The time when the job was executed.
	//
	// example:
	//
	// 1657522839
	CompletionTime *int64 `json:"CompletionTime,omitempty" xml:"CompletionTime,omitempty"`
	// The number of instances that failed to run.
	//
	// example:
	//
	// 0
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The job ID.
	//
	// example:
	//
	// event-b798157b-40a2-4388-b578-71fb897103**-**
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The returned message. Take note of the following rules:
	//
	// 	- If the call is successful, **success*	- is returned.
	//
	// 	- If the call fails, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The time when the job was created.
	//
	// example:
	//
	// 1657522800
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the job. Valid values:
	//
	// 	- **0**: The job is not executed.
	//
	// 	- **1**: The job was executed.
	//
	// 	- **2**: The job failed to be executed.
	//
	// 	- **3**: The job is being executed.
	//
	// example:
	//
	// 1
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The number of instances that are successfully run.
	//
	// example:
	//
	// 3
	Succeeded *int64 `json:"Succeeded,omitempty" xml:"Succeeded,omitempty"`
}

func (s DescribeJobStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeJobStatusResponseBodyData) GetActive() *int64 {
	return s.Active
}

func (s *DescribeJobStatusResponseBodyData) GetCompletionTime() *int64 {
	return s.CompletionTime
}

func (s *DescribeJobStatusResponseBodyData) GetFailed() *int64 {
	return s.Failed
}

func (s *DescribeJobStatusResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *DescribeJobStatusResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DescribeJobStatusResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeJobStatusResponseBodyData) GetState() *string {
	return s.State
}

func (s *DescribeJobStatusResponseBodyData) GetSucceeded() *int64 {
	return s.Succeeded
}

func (s *DescribeJobStatusResponseBodyData) SetActive(v int64) *DescribeJobStatusResponseBodyData {
	s.Active = &v
	return s
}

func (s *DescribeJobStatusResponseBodyData) SetCompletionTime(v int64) *DescribeJobStatusResponseBodyData {
	s.CompletionTime = &v
	return s
}

func (s *DescribeJobStatusResponseBodyData) SetFailed(v int64) *DescribeJobStatusResponseBodyData {
	s.Failed = &v
	return s
}

func (s *DescribeJobStatusResponseBodyData) SetJobId(v string) *DescribeJobStatusResponseBodyData {
	s.JobId = &v
	return s
}

func (s *DescribeJobStatusResponseBodyData) SetMessage(v string) *DescribeJobStatusResponseBodyData {
	s.Message = &v
	return s
}

func (s *DescribeJobStatusResponseBodyData) SetStartTime(v int64) *DescribeJobStatusResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeJobStatusResponseBodyData) SetState(v string) *DescribeJobStatusResponseBodyData {
	s.State = &v
	return s
}

func (s *DescribeJobStatusResponseBodyData) SetSucceeded(v int64) *DescribeJobStatusResponseBodyData {
	s.Succeeded = &v
	return s
}

func (s *DescribeJobStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
