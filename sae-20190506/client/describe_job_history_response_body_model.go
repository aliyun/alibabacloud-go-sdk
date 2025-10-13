// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeJobHistoryResponseBody
	GetCode() *string
	SetData(v *DescribeJobHistoryResponseBodyData) *DescribeJobHistoryResponseBody
	GetData() *DescribeJobHistoryResponseBodyData
	SetErrorCode(v string) *DescribeJobHistoryResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeJobHistoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeJobHistoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeJobHistoryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeJobHistoryResponseBody
	GetTraceId() *string
}

type DescribeJobHistoryResponseBody struct {
	// The HTTP status code returned. Valid values:
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
	Data *DescribeJobHistoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned. Take note of the following rules:
	//
	// 	- If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the "**Error codes**" section in this topic.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned. Take note of the following rules:
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

func (s DescribeJobHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobHistoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeJobHistoryResponseBody) GetData() *DescribeJobHistoryResponseBodyData {
	return s.Data
}

func (s *DescribeJobHistoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeJobHistoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeJobHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeJobHistoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeJobHistoryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeJobHistoryResponseBody) SetCode(v string) *DescribeJobHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeJobHistoryResponseBody) SetData(v *DescribeJobHistoryResponseBodyData) *DescribeJobHistoryResponseBody {
	s.Data = v
	return s
}

func (s *DescribeJobHistoryResponseBody) SetErrorCode(v string) *DescribeJobHistoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeJobHistoryResponseBody) SetMessage(v string) *DescribeJobHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeJobHistoryResponseBody) SetRequestId(v string) *DescribeJobHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJobHistoryResponseBody) SetSuccess(v bool) *DescribeJobHistoryResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeJobHistoryResponseBody) SetTraceId(v string) *DescribeJobHistoryResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeJobHistoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeJobHistoryResponseBodyData struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The jobs.
	Jobs []*DescribeJobHistoryResponseBodyDataJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// The number of entries to return on each page. Valid values: 0 to 10000.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of jobs.
	//
	// example:
	//
	// 20
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeJobHistoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobHistoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeJobHistoryResponseBodyData) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribeJobHistoryResponseBodyData) GetJobs() []*DescribeJobHistoryResponseBodyDataJobs {
	return s.Jobs
}

func (s *DescribeJobHistoryResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeJobHistoryResponseBodyData) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *DescribeJobHistoryResponseBodyData) SetCurrentPage(v int64) *DescribeJobHistoryResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *DescribeJobHistoryResponseBodyData) SetJobs(v []*DescribeJobHistoryResponseBodyDataJobs) *DescribeJobHistoryResponseBodyData {
	s.Jobs = v
	return s
}

func (s *DescribeJobHistoryResponseBodyData) SetPageSize(v int64) *DescribeJobHistoryResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeJobHistoryResponseBodyData) SetTotalSize(v int64) *DescribeJobHistoryResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *DescribeJobHistoryResponseBodyData) Validate() error {
	if s.Jobs != nil {
		for _, item := range s.Jobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeJobHistoryResponseBodyDataJobs struct {
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
	// manual-3db7a8fa-5d40-4edc-92e4-49d50eab****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The message returned if exceptions occur during job running.
	//
	// example:
	//
	// Null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The time when the job was created.
	//
	// example:
	//
	// 1657522800
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the job. Valid values:
	//
	// 	- **0**: The job was not executed.
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

func (s DescribeJobHistoryResponseBodyDataJobs) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobHistoryResponseBodyDataJobs) GoString() string {
	return s.String()
}

func (s *DescribeJobHistoryResponseBodyDataJobs) GetActive() *int64 {
	return s.Active
}

func (s *DescribeJobHistoryResponseBodyDataJobs) GetCompletionTime() *int64 {
	return s.CompletionTime
}

func (s *DescribeJobHistoryResponseBodyDataJobs) GetFailed() *int64 {
	return s.Failed
}

func (s *DescribeJobHistoryResponseBodyDataJobs) GetJobId() *string {
	return s.JobId
}

func (s *DescribeJobHistoryResponseBodyDataJobs) GetMessage() *string {
	return s.Message
}

func (s *DescribeJobHistoryResponseBodyDataJobs) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeJobHistoryResponseBodyDataJobs) GetState() *string {
	return s.State
}

func (s *DescribeJobHistoryResponseBodyDataJobs) GetSucceeded() *int64 {
	return s.Succeeded
}

func (s *DescribeJobHistoryResponseBodyDataJobs) SetActive(v int64) *DescribeJobHistoryResponseBodyDataJobs {
	s.Active = &v
	return s
}

func (s *DescribeJobHistoryResponseBodyDataJobs) SetCompletionTime(v int64) *DescribeJobHistoryResponseBodyDataJobs {
	s.CompletionTime = &v
	return s
}

func (s *DescribeJobHistoryResponseBodyDataJobs) SetFailed(v int64) *DescribeJobHistoryResponseBodyDataJobs {
	s.Failed = &v
	return s
}

func (s *DescribeJobHistoryResponseBodyDataJobs) SetJobId(v string) *DescribeJobHistoryResponseBodyDataJobs {
	s.JobId = &v
	return s
}

func (s *DescribeJobHistoryResponseBodyDataJobs) SetMessage(v string) *DescribeJobHistoryResponseBodyDataJobs {
	s.Message = &v
	return s
}

func (s *DescribeJobHistoryResponseBodyDataJobs) SetStartTime(v int64) *DescribeJobHistoryResponseBodyDataJobs {
	s.StartTime = &v
	return s
}

func (s *DescribeJobHistoryResponseBodyDataJobs) SetState(v string) *DescribeJobHistoryResponseBodyDataJobs {
	s.State = &v
	return s
}

func (s *DescribeJobHistoryResponseBodyDataJobs) SetSucceeded(v int64) *DescribeJobHistoryResponseBodyDataJobs {
	s.Succeeded = &v
	return s
}

func (s *DescribeJobHistoryResponseBodyDataJobs) Validate() error {
	return dara.Validate(s)
}
