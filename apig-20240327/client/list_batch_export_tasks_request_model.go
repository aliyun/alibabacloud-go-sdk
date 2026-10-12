// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchExportTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListBatchExportTasksRequest
	GetEndTime() *string
	SetGatewayId(v string) *ListBatchExportTasksRequest
	GetGatewayId() *string
	SetMaxResults(v int32) *ListBatchExportTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListBatchExportTasksRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListBatchExportTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBatchExportTasksRequest
	GetPageSize() *int32
	SetStartTime(v string) *ListBatchExportTasksRequest
	GetStartTime() *string
	SetStatuses(v string) *ListBatchExportTasksRequest
	GetStatuses() *string
}

type ListBatchExportTasksRequest struct {
	// The end of the creation time range. Format: yyyy-MM-dd\\"T\\"HH:mm:ss\\"Z\\".
	//
	// example:
	//
	// 2026-05-26T11:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The gateway instance ID. If specified, only tasks under this gateway are returned.
	//
	// example:
	//
	// gw-xxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The maximum number of tasks to return in a single request.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The paging token that indicates the start position of the next query.
	//
	// example:
	//
	// token-xxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	// The start of the creation time range. Format: yyyy-MM-dd\\"T\\"HH:mm:ss\\"Z\\".
	//
	// example:
	//
	// 2026-05-26T10:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The list of task statuses. Separate multiple statuses with commas (,).
	//
	// example:
	//
	// Pending,Running
	Statuses *string `json:"statuses,omitempty" xml:"statuses,omitempty"`
}

func (s ListBatchExportTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBatchExportTasksRequest) GoString() string {
	return s.String()
}

func (s *ListBatchExportTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListBatchExportTasksRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListBatchExportTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBatchExportTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBatchExportTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBatchExportTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBatchExportTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListBatchExportTasksRequest) GetStatuses() *string {
	return s.Statuses
}

func (s *ListBatchExportTasksRequest) SetEndTime(v string) *ListBatchExportTasksRequest {
	s.EndTime = &v
	return s
}

func (s *ListBatchExportTasksRequest) SetGatewayId(v string) *ListBatchExportTasksRequest {
	s.GatewayId = &v
	return s
}

func (s *ListBatchExportTasksRequest) SetMaxResults(v int32) *ListBatchExportTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBatchExportTasksRequest) SetNextToken(v string) *ListBatchExportTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListBatchExportTasksRequest) SetPageNumber(v int32) *ListBatchExportTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBatchExportTasksRequest) SetPageSize(v int32) *ListBatchExportTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListBatchExportTasksRequest) SetStartTime(v string) *ListBatchExportTasksRequest {
	s.StartTime = &v
	return s
}

func (s *ListBatchExportTasksRequest) SetStatuses(v string) *ListBatchExportTasksRequest {
	s.Statuses = &v
	return s
}

func (s *ListBatchExportTasksRequest) Validate() error {
	return dara.Validate(s)
}
