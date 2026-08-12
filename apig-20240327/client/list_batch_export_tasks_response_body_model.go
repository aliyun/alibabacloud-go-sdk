// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchExportTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBatchExportTasksResponseBody
	GetCode() *string
	SetData(v *ListBatchExportTasksResponseBodyData) *ListBatchExportTasksResponseBody
	GetData() *ListBatchExportTasksResponseBodyData
	SetMaxResults(v int32) *ListBatchExportTasksResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListBatchExportTasksResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListBatchExportTasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListBatchExportTasksResponseBody
	GetRequestId() *string
}

type ListBatchExportTasksResponseBody struct {
	// example:
	//
	// Ok
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListBatchExportTasksResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// token-xxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// CE534E1D-FCE4-5930-B784-E055EC1AEE6F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListBatchExportTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBatchExportTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListBatchExportTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBatchExportTasksResponseBody) GetData() *ListBatchExportTasksResponseBodyData {
	return s.Data
}

func (s *ListBatchExportTasksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBatchExportTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBatchExportTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBatchExportTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBatchExportTasksResponseBody) SetCode(v string) *ListBatchExportTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListBatchExportTasksResponseBody) SetData(v *ListBatchExportTasksResponseBodyData) *ListBatchExportTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListBatchExportTasksResponseBody) SetMaxResults(v int32) *ListBatchExportTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListBatchExportTasksResponseBody) SetMessage(v string) *ListBatchExportTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListBatchExportTasksResponseBody) SetNextToken(v string) *ListBatchExportTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBatchExportTasksResponseBody) SetRequestId(v string) *ListBatchExportTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBatchExportTasksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBatchExportTasksResponseBodyData struct {
	Items []*ListBatchExportTasksResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListBatchExportTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBatchExportTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBatchExportTasksResponseBodyData) GetItems() []*ListBatchExportTasksResponseBodyDataItems {
	return s.Items
}

func (s *ListBatchExportTasksResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBatchExportTasksResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBatchExportTasksResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListBatchExportTasksResponseBodyData) SetItems(v []*ListBatchExportTasksResponseBodyDataItems) *ListBatchExportTasksResponseBodyData {
	s.Items = v
	return s
}

func (s *ListBatchExportTasksResponseBodyData) SetPageNumber(v int32) *ListBatchExportTasksResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListBatchExportTasksResponseBodyData) SetPageSize(v int32) *ListBatchExportTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListBatchExportTasksResponseBodyData) SetTotalSize(v int32) *ListBatchExportTasksResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListBatchExportTasksResponseBodyData) Validate() error {
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

type ListBatchExportTasksResponseBodyDataItems struct {
	// example:
	//
	// 2026-05-26T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// some apis failed
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// gw-xxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// 5
	ProcessedCount *int32 `json:"processedCount,omitempty" xml:"processedCount,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// task-xxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// BatchExport
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListBatchExportTasksResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListBatchExportTasksResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListBatchExportTasksResponseBodyDataItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListBatchExportTasksResponseBodyDataItems) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListBatchExportTasksResponseBodyDataItems) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListBatchExportTasksResponseBodyDataItems) GetProcessedCount() *int32 {
	return s.ProcessedCount
}

func (s *ListBatchExportTasksResponseBodyDataItems) GetStatus() *string {
	return s.Status
}

func (s *ListBatchExportTasksResponseBodyDataItems) GetTaskId() *string {
	return s.TaskId
}

func (s *ListBatchExportTasksResponseBodyDataItems) GetTaskType() *string {
	return s.TaskType
}

func (s *ListBatchExportTasksResponseBodyDataItems) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBatchExportTasksResponseBodyDataItems) SetCreateTime(v string) *ListBatchExportTasksResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *ListBatchExportTasksResponseBodyDataItems) SetErrorMessage(v string) *ListBatchExportTasksResponseBodyDataItems {
	s.ErrorMessage = &v
	return s
}

func (s *ListBatchExportTasksResponseBodyDataItems) SetGatewayId(v string) *ListBatchExportTasksResponseBodyDataItems {
	s.GatewayId = &v
	return s
}

func (s *ListBatchExportTasksResponseBodyDataItems) SetProcessedCount(v int32) *ListBatchExportTasksResponseBodyDataItems {
	s.ProcessedCount = &v
	return s
}

func (s *ListBatchExportTasksResponseBodyDataItems) SetStatus(v string) *ListBatchExportTasksResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *ListBatchExportTasksResponseBodyDataItems) SetTaskId(v string) *ListBatchExportTasksResponseBodyDataItems {
	s.TaskId = &v
	return s
}

func (s *ListBatchExportTasksResponseBodyDataItems) SetTaskType(v string) *ListBatchExportTasksResponseBodyDataItems {
	s.TaskType = &v
	return s
}

func (s *ListBatchExportTasksResponseBodyDataItems) SetTotalCount(v int32) *ListBatchExportTasksResponseBodyDataItems {
	s.TotalCount = &v
	return s
}

func (s *ListBatchExportTasksResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
