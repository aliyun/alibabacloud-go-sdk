// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceTestTaskLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServiceTestTaskLogsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceTestTaskLogsRequest
	GetNextToken() *string
	SetSortOrder(v string) *ListServiceTestTaskLogsRequest
	GetSortOrder() *string
	SetTaskId(v string) *ListServiceTestTaskLogsRequest
	GetTaskId() *string
}

type ListServiceTestTaskLogsRequest struct {
	// The number of items to return per page when paginating results. The maximum is 100, and the default is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Sort Order. Possible values:
	//
	// + Ascending: Ascending order
	//
	// + Descending (default value): Descending order
	//
	// example:
	//
	// Ascending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The task ID.
	//
	// example:
	//
	// stt-568c2c5a687a409b977e
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListServiceTestTaskLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceTestTaskLogsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceTestTaskLogsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceTestTaskLogsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceTestTaskLogsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListServiceTestTaskLogsRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListServiceTestTaskLogsRequest) SetMaxResults(v int32) *ListServiceTestTaskLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceTestTaskLogsRequest) SetNextToken(v string) *ListServiceTestTaskLogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceTestTaskLogsRequest) SetSortOrder(v string) *ListServiceTestTaskLogsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListServiceTestTaskLogsRequest) SetTaskId(v string) *ListServiceTestTaskLogsRequest {
	s.TaskId = &v
	return s
}

func (s *ListServiceTestTaskLogsRequest) Validate() error {
	return dara.Validate(s)
}
