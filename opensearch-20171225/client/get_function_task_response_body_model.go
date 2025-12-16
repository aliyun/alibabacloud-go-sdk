// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetFunctionTaskResponseBody
	GetCode() *string
	SetHttpCode(v int64) *GetFunctionTaskResponseBody
	GetHttpCode() *int64
	SetLatency(v int64) *GetFunctionTaskResponseBody
	GetLatency() *int64
	SetMessage(v string) *GetFunctionTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFunctionTaskResponseBody
	GetRequestId() *string
	SetResult(v *GetFunctionTaskResponseBodyResult) *GetFunctionTaskResponseBody
	GetResult() *GetFunctionTaskResponseBodyResult
	SetStatus(v string) *GetFunctionTaskResponseBody
	GetStatus() *string
}

type GetFunctionTaskResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Task.NotExist
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	//
	// example:
	//
	// 123
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A4D487A9-A456-5AA5-A9C6-B7BF2889CF74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	Result *GetFunctionTaskResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The status of the request.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFunctionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFunctionTaskResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *GetFunctionTaskResponseBody) GetLatency() *int64 {
	return s.Latency
}

func (s *GetFunctionTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFunctionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFunctionTaskResponseBody) GetResult() *GetFunctionTaskResponseBodyResult {
	return s.Result
}

func (s *GetFunctionTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetFunctionTaskResponseBody) SetCode(v string) *GetFunctionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetFunctionTaskResponseBody) SetHttpCode(v int64) *GetFunctionTaskResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetFunctionTaskResponseBody) SetLatency(v int64) *GetFunctionTaskResponseBody {
	s.Latency = &v
	return s
}

func (s *GetFunctionTaskResponseBody) SetMessage(v string) *GetFunctionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetFunctionTaskResponseBody) SetRequestId(v string) *GetFunctionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFunctionTaskResponseBody) SetResult(v *GetFunctionTaskResponseBodyResult) *GetFunctionTaskResponseBody {
	s.Result = v
	return s
}

func (s *GetFunctionTaskResponseBody) SetStatus(v string) *GetFunctionTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetFunctionTaskResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFunctionTaskResponseBodyResult struct {
	// The timestamp that indicates the end time of the task. Unit: milliseconds.
	//
	// example:
	//
	// 1647213406267
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The extended information, which is a JSON string.
	//
	// example:
	//
	// {\\"recall\\":91,\\"errors\\":[]}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The name of the feature.
	//
	// example:
	//
	// ctr
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The number of iterations.
	//
	// example:
	//
	// 1
	Generation *string `json:"Generation,omitempty" xml:"Generation,omitempty"`
	// The progress. 90 indicates 90%.
	//
	// example:
	//
	// 90
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// trigger__2021-03-05T12:18:41
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// The timestamp that indicates the start time of the task. Unit: milliseconds.
	//
	// example:
	//
	// 1647212220000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- success
	//
	// 	- failed
	//
	// 	- running
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFunctionTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFunctionTaskResponseBodyResult) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetFunctionTaskResponseBodyResult) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *GetFunctionTaskResponseBodyResult) GetFunctionName() *string {
	return s.FunctionName
}

func (s *GetFunctionTaskResponseBodyResult) GetGeneration() *string {
	return s.Generation
}

func (s *GetFunctionTaskResponseBodyResult) GetProgress() *int64 {
	return s.Progress
}

func (s *GetFunctionTaskResponseBodyResult) GetRunId() *string {
	return s.RunId
}

func (s *GetFunctionTaskResponseBodyResult) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetFunctionTaskResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetFunctionTaskResponseBodyResult) SetEndTime(v int64) *GetFunctionTaskResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *GetFunctionTaskResponseBodyResult) SetExtendInfo(v string) *GetFunctionTaskResponseBodyResult {
	s.ExtendInfo = &v
	return s
}

func (s *GetFunctionTaskResponseBodyResult) SetFunctionName(v string) *GetFunctionTaskResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *GetFunctionTaskResponseBodyResult) SetGeneration(v string) *GetFunctionTaskResponseBodyResult {
	s.Generation = &v
	return s
}

func (s *GetFunctionTaskResponseBodyResult) SetProgress(v int64) *GetFunctionTaskResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *GetFunctionTaskResponseBodyResult) SetRunId(v string) *GetFunctionTaskResponseBodyResult {
	s.RunId = &v
	return s
}

func (s *GetFunctionTaskResponseBodyResult) SetStartTime(v int64) *GetFunctionTaskResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *GetFunctionTaskResponseBodyResult) SetStatus(v string) *GetFunctionTaskResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetFunctionTaskResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
