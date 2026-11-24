// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchImportTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBatchImportTasksResponseBody
	GetCode() *string
	SetData(v *ListBatchImportTasksResponseBodyData) *ListBatchImportTasksResponseBody
	GetData() *ListBatchImportTasksResponseBodyData
	SetMaxResults(v int32) *ListBatchImportTasksResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListBatchImportTasksResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListBatchImportTasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListBatchImportTasksResponseBody
	GetRequestId() *string
}

type ListBatchImportTasksResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ListBatchImportTasksResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The maximum number of records returned in the current request.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next query. This parameter is empty if no more results are available.
	//
	// example:
	//
	// token-xxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CE534E1D-FCE4-5930-B784-E055EC1AEE6F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListBatchImportTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBatchImportTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListBatchImportTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBatchImportTasksResponseBody) GetData() *ListBatchImportTasksResponseBodyData {
	return s.Data
}

func (s *ListBatchImportTasksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBatchImportTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBatchImportTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBatchImportTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBatchImportTasksResponseBody) SetCode(v string) *ListBatchImportTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListBatchImportTasksResponseBody) SetData(v *ListBatchImportTasksResponseBodyData) *ListBatchImportTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListBatchImportTasksResponseBody) SetMaxResults(v int32) *ListBatchImportTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListBatchImportTasksResponseBody) SetMessage(v string) *ListBatchImportTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListBatchImportTasksResponseBody) SetNextToken(v string) *ListBatchImportTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBatchImportTasksResponseBody) SetRequestId(v string) *ListBatchImportTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBatchImportTasksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBatchImportTasksResponseBodyData struct {
	// The list of tasks. Tasks are returned in descending order of creation time, with the latest tasks first.
	Items []*ListBatchImportTasksResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of tasks.
	//
	// example:
	//
	// 1
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListBatchImportTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBatchImportTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBatchImportTasksResponseBodyData) GetItems() []*ListBatchImportTasksResponseBodyDataItems {
	return s.Items
}

func (s *ListBatchImportTasksResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBatchImportTasksResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBatchImportTasksResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListBatchImportTasksResponseBodyData) SetItems(v []*ListBatchImportTasksResponseBodyDataItems) *ListBatchImportTasksResponseBodyData {
	s.Items = v
	return s
}

func (s *ListBatchImportTasksResponseBodyData) SetPageNumber(v int32) *ListBatchImportTasksResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListBatchImportTasksResponseBodyData) SetPageSize(v int32) *ListBatchImportTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListBatchImportTasksResponseBodyData) SetTotalSize(v int32) *ListBatchImportTasksResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListBatchImportTasksResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBatchImportTasksResponseBodyDataItems struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2026-05-26T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The error message.
	//
	// example:
	//
	// some apis failed
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The gateway instance ID associated with the task.
	//
	// example:
	//
	// gw-xxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The number of processed items.
	//
	// example:
	//
	// 5
	ProcessedCount *int32 `json:"processedCount,omitempty" xml:"processedCount,omitempty"`
	// The status of the task. Valid values: Pending, Running, Success, PartialSuccess, and Failed.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// task-xxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The type of the task.
	//
	// example:
	//
	// BatchImport
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// The total number of tasks.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListBatchImportTasksResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListBatchImportTasksResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListBatchImportTasksResponseBodyDataItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListBatchImportTasksResponseBodyDataItems) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListBatchImportTasksResponseBodyDataItems) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListBatchImportTasksResponseBodyDataItems) GetProcessedCount() *int32 {
	return s.ProcessedCount
}

func (s *ListBatchImportTasksResponseBodyDataItems) GetStatus() *string {
	return s.Status
}

func (s *ListBatchImportTasksResponseBodyDataItems) GetTaskId() *string {
	return s.TaskId
}

func (s *ListBatchImportTasksResponseBodyDataItems) GetTaskType() *string {
	return s.TaskType
}

func (s *ListBatchImportTasksResponseBodyDataItems) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBatchImportTasksResponseBodyDataItems) SetCreateTime(v string) *ListBatchImportTasksResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *ListBatchImportTasksResponseBodyDataItems) SetErrorMessage(v string) *ListBatchImportTasksResponseBodyDataItems {
	s.ErrorMessage = &v
	return s
}

func (s *ListBatchImportTasksResponseBodyDataItems) SetGatewayId(v string) *ListBatchImportTasksResponseBodyDataItems {
	s.GatewayId = &v
	return s
}

func (s *ListBatchImportTasksResponseBodyDataItems) SetProcessedCount(v int32) *ListBatchImportTasksResponseBodyDataItems {
	s.ProcessedCount = &v
	return s
}

func (s *ListBatchImportTasksResponseBodyDataItems) SetStatus(v string) *ListBatchImportTasksResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *ListBatchImportTasksResponseBodyDataItems) SetTaskId(v string) *ListBatchImportTasksResponseBodyDataItems {
	s.TaskId = &v
	return s
}

func (s *ListBatchImportTasksResponseBodyDataItems) SetTaskType(v string) *ListBatchImportTasksResponseBodyDataItems {
	s.TaskType = &v
	return s
}

func (s *ListBatchImportTasksResponseBodyDataItems) SetTotalCount(v int32) *ListBatchImportTasksResponseBodyDataItems {
	s.TotalCount = &v
	return s
}

func (s *ListBatchImportTasksResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
