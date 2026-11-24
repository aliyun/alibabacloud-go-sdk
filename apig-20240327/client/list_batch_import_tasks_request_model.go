// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchImportTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListBatchImportTasksRequest
	GetEndTime() *string
	SetGatewayId(v string) *ListBatchImportTasksRequest
	GetGatewayId() *string
	SetMaxResults(v int32) *ListBatchImportTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListBatchImportTasksRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListBatchImportTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBatchImportTasksRequest
	GetPageSize() *int32
	SetStartTime(v string) *ListBatchImportTasksRequest
	GetStartTime() *string
	SetStatuses(v string) *ListBatchImportTasksRequest
	GetStatuses() *string
}

type ListBatchImportTasksRequest struct {
	// The end time of the creation time range. Specify the time in RFC 3339 format.
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
	// The start mark for the paging query.
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
	// The start time of the creation time range. Specify the time in RFC 3339 format.
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

func (s ListBatchImportTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBatchImportTasksRequest) GoString() string {
	return s.String()
}

func (s *ListBatchImportTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListBatchImportTasksRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListBatchImportTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBatchImportTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBatchImportTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBatchImportTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBatchImportTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListBatchImportTasksRequest) GetStatuses() *string {
	return s.Statuses
}

func (s *ListBatchImportTasksRequest) SetEndTime(v string) *ListBatchImportTasksRequest {
	s.EndTime = &v
	return s
}

func (s *ListBatchImportTasksRequest) SetGatewayId(v string) *ListBatchImportTasksRequest {
	s.GatewayId = &v
	return s
}

func (s *ListBatchImportTasksRequest) SetMaxResults(v int32) *ListBatchImportTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBatchImportTasksRequest) SetNextToken(v string) *ListBatchImportTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListBatchImportTasksRequest) SetPageNumber(v int32) *ListBatchImportTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBatchImportTasksRequest) SetPageSize(v int32) *ListBatchImportTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListBatchImportTasksRequest) SetStartTime(v string) *ListBatchImportTasksRequest {
	s.StartTime = &v
	return s
}

func (s *ListBatchImportTasksRequest) SetStatuses(v string) *ListBatchImportTasksRequest {
	s.Statuses = &v
	return s
}

func (s *ListBatchImportTasksRequest) Validate() error {
	return dara.Validate(s)
}
