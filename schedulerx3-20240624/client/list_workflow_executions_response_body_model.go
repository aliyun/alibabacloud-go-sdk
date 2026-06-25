// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowExecutionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListWorkflowExecutionsResponseBody
	GetCode() *int32
	SetData(v *ListWorkflowExecutionsResponseBodyData) *ListWorkflowExecutionsResponseBody
	GetData() *ListWorkflowExecutionsResponseBodyData
	SetMaxResults(v int32) *ListWorkflowExecutionsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListWorkflowExecutionsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListWorkflowExecutionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWorkflowExecutionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWorkflowExecutionsResponseBody
	GetSuccess() *bool
}

type ListWorkflowExecutionsResponseBody struct {
	// The response code. A value of `200` indicates a successful request.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *ListWorkflowExecutionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The maximum number of entries to return per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The response message. If the request fails, this parameter provides details about the error.
	//
	// This parameter is required.
	//
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token to retrieve the next page of results. If this parameter is not returned, all results have been returned.
	//
	// example:
	//
	// eCKqVlS5FKF5EWGGOo8EgQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The unique request ID. Use this ID to troubleshoot issues.
	//
	// example:
	//
	// BAC1ADB5-EEB5-5834-93D8-522E067AF8D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - `true`: The request was successful.
	//
	// - `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWorkflowExecutionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowExecutionsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListWorkflowExecutionsResponseBody) GetData() *ListWorkflowExecutionsResponseBodyData {
	return s.Data
}

func (s *ListWorkflowExecutionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkflowExecutionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWorkflowExecutionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkflowExecutionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkflowExecutionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkflowExecutionsResponseBody) SetCode(v int32) *ListWorkflowExecutionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListWorkflowExecutionsResponseBody) SetData(v *ListWorkflowExecutionsResponseBodyData) *ListWorkflowExecutionsResponseBody {
	s.Data = v
	return s
}

func (s *ListWorkflowExecutionsResponseBody) SetMaxResults(v int32) *ListWorkflowExecutionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkflowExecutionsResponseBody) SetMessage(v string) *ListWorkflowExecutionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListWorkflowExecutionsResponseBody) SetNextToken(v string) *ListWorkflowExecutionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowExecutionsResponseBody) SetRequestId(v string) *ListWorkflowExecutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkflowExecutionsResponseBody) SetSuccess(v bool) *ListWorkflowExecutionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkflowExecutionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkflowExecutionsResponseBodyData struct {
	// The page number.
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
	// -
	Records []*ListWorkflowExecutionsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The total number of matching workflow executions.
	//
	// example:
	//
	// 65
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListWorkflowExecutionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowExecutionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWorkflowExecutionsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWorkflowExecutionsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkflowExecutionsResponseBodyData) GetRecords() []*ListWorkflowExecutionsResponseBodyDataRecords {
	return s.Records
}

func (s *ListWorkflowExecutionsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListWorkflowExecutionsResponseBodyData) SetPageNumber(v int32) *ListWorkflowExecutionsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListWorkflowExecutionsResponseBodyData) SetPageSize(v int32) *ListWorkflowExecutionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowExecutionsResponseBodyData) SetRecords(v []*ListWorkflowExecutionsResponseBodyDataRecords) *ListWorkflowExecutionsResponseBodyData {
	s.Records = v
	return s
}

func (s *ListWorkflowExecutionsResponseBodyData) SetTotal(v int32) *ListWorkflowExecutionsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListWorkflowExecutionsResponseBodyData) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkflowExecutionsResponseBodyDataRecords struct {
	// The application name.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The data time of the workflow execution.
	//
	// example:
	//
	// 2024-11-12 14:52:42
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// The end time of the workflow execution.
	//
	// example:
	//
	// 2024-11-12 14:52:42
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The executor ID.
	//
	// example:
	//
	// 1827811800526000
	Executor *string `json:"Executor,omitempty" xml:"Executor,omitempty"`
	// The schedule time of the workflow execution.
	//
	// example:
	//
	// 2024-11-12 14:52:42
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The start time of the workflow execution.
	//
	// example:
	//
	// 2025-11-04 01:09:27
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the workflow execution. Valid values:
	//
	// - 0: UNKNOWN
	//
	// - 1: WAITING
	//
	// - 2: READY
	//
	// - 3: RUNNING
	//
	// - 4: SUCCESS
	//
	// - 5: FAILED
	//
	// - 6: PAUSED
	//
	// - 7: SUBMITTED
	//
	// - 8: REJECTED
	//
	// - 9: ACCEPTED
	//
	// - 10: PARTIAL_FAILED
	//
	// - 11: SKIPPED
	//
	// - 12: REMOVED
	//
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The workflow execution ID.
	//
	// example:
	//
	// 100
	WorkflowExecutionId *string `json:"WorkflowExecutionId,omitempty" xml:"WorkflowExecutionId,omitempty"`
	// The workflow ID.
	//
	// example:
	//
	// 10
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
	// The workflow name.
	//
	// example:
	//
	// myWorkflow
	WorkflowName *string `json:"WorkflowName,omitempty" xml:"WorkflowName,omitempty"`
}

func (s ListWorkflowExecutionsResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowExecutionsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) GetAppName() *string {
	return s.AppName
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) GetDataTime() *string {
	return s.DataTime
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) GetEndTime() *string {
	return s.EndTime
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) GetExecutor() *string {
	return s.Executor
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) GetStartTime() *string {
	return s.StartTime
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) GetStatus() *int32 {
	return s.Status
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) GetWorkflowExecutionId() *string {
	return s.WorkflowExecutionId
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) SetAppName(v string) *ListWorkflowExecutionsResponseBodyDataRecords {
	s.AppName = &v
	return s
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) SetDataTime(v string) *ListWorkflowExecutionsResponseBodyDataRecords {
	s.DataTime = &v
	return s
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) SetEndTime(v string) *ListWorkflowExecutionsResponseBodyDataRecords {
	s.EndTime = &v
	return s
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) SetExecutor(v string) *ListWorkflowExecutionsResponseBodyDataRecords {
	s.Executor = &v
	return s
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) SetScheduleTime(v string) *ListWorkflowExecutionsResponseBodyDataRecords {
	s.ScheduleTime = &v
	return s
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) SetStartTime(v string) *ListWorkflowExecutionsResponseBodyDataRecords {
	s.StartTime = &v
	return s
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) SetStatus(v int32) *ListWorkflowExecutionsResponseBodyDataRecords {
	s.Status = &v
	return s
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) SetWorkflowExecutionId(v string) *ListWorkflowExecutionsResponseBodyDataRecords {
	s.WorkflowExecutionId = &v
	return s
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) SetWorkflowId(v string) *ListWorkflowExecutionsResponseBodyDataRecords {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) SetWorkflowName(v string) *ListWorkflowExecutionsResponseBodyDataRecords {
	s.WorkflowName = &v
	return s
}

func (s *ListWorkflowExecutionsResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
