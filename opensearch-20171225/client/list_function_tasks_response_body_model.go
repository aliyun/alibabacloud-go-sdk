// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFunctionTasksResponseBody
	GetCode() *string
	SetHttpCode(v int64) *ListFunctionTasksResponseBody
	GetHttpCode() *int64
	SetLatency(v int64) *ListFunctionTasksResponseBody
	GetLatency() *int64
	SetMessage(v string) *ListFunctionTasksResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListFunctionTasksResponseBody
	GetRequestId() *string
	SetResult(v []*ListFunctionTasksResponseBodyResult) *ListFunctionTasksResponseBody
	GetResult() []*ListFunctionTasksResponseBodyResult
	SetStatus(v string) *ListFunctionTasksResponseBody
	GetStatus() *string
	SetTotalCount(v int64) *ListFunctionTasksResponseBody
	GetTotalCount() *int64
}

type ListFunctionTasksResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
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
	// fail
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1638157479281
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	//
	// example:
	//
	// [         {             "functionName": "ctr",             "progress": 100,             "status": "success",             "startTime": 100010,             "endTime": 200020,             "extendInfo": "{\\"recall\\":91,\\"errors\\":[]}",             "runId": "trigger__2021-03-05T12:18:41"         }     ]
	Result []*ListFunctionTasksResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The status of the request.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of records that meet the requirements.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFunctionTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListFunctionTasksResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *ListFunctionTasksResponseBody) GetLatency() *int64 {
	return s.Latency
}

func (s *ListFunctionTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFunctionTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFunctionTasksResponseBody) GetResult() []*ListFunctionTasksResponseBodyResult {
	return s.Result
}

func (s *ListFunctionTasksResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListFunctionTasksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListFunctionTasksResponseBody) SetCode(v string) *ListFunctionTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListFunctionTasksResponseBody) SetHttpCode(v int64) *ListFunctionTasksResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListFunctionTasksResponseBody) SetLatency(v int64) *ListFunctionTasksResponseBody {
	s.Latency = &v
	return s
}

func (s *ListFunctionTasksResponseBody) SetMessage(v string) *ListFunctionTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListFunctionTasksResponseBody) SetRequestId(v string) *ListFunctionTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFunctionTasksResponseBody) SetResult(v []*ListFunctionTasksResponseBodyResult) *ListFunctionTasksResponseBody {
	s.Result = v
	return s
}

func (s *ListFunctionTasksResponseBody) SetStatus(v string) *ListFunctionTasksResponseBody {
	s.Status = &v
	return s
}

func (s *ListFunctionTasksResponseBody) SetTotalCount(v int64) *ListFunctionTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFunctionTasksResponseBody) Validate() error {
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

type ListFunctionTasksResponseBodyResult struct {
	// The timestamp that indicates the end time. Unit: milliseconds. 0 indicates that the task has not ended.
	//
	// example:
	//
	// 100010
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The value is a JSON string. It includes model evaluation information and training error information.
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
	// 2
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
	// The timestamp that indicates the start time. Unit: milliseconds.
	//
	// example:
	//
	// 100010
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

func (s ListFunctionTasksResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListFunctionTasksResponseBodyResult) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListFunctionTasksResponseBodyResult) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *ListFunctionTasksResponseBodyResult) GetFunctionName() *string {
	return s.FunctionName
}

func (s *ListFunctionTasksResponseBodyResult) GetGeneration() *string {
	return s.Generation
}

func (s *ListFunctionTasksResponseBodyResult) GetProgress() *int64 {
	return s.Progress
}

func (s *ListFunctionTasksResponseBodyResult) GetRunId() *string {
	return s.RunId
}

func (s *ListFunctionTasksResponseBodyResult) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListFunctionTasksResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListFunctionTasksResponseBodyResult) SetEndTime(v int64) *ListFunctionTasksResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) SetExtendInfo(v string) *ListFunctionTasksResponseBodyResult {
	s.ExtendInfo = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) SetFunctionName(v string) *ListFunctionTasksResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) SetGeneration(v string) *ListFunctionTasksResponseBodyResult {
	s.Generation = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) SetProgress(v int64) *ListFunctionTasksResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) SetRunId(v string) *ListFunctionTasksResponseBodyResult {
	s.RunId = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) SetStartTime(v int64) *ListFunctionTasksResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) SetStatus(v string) *ListFunctionTasksResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
