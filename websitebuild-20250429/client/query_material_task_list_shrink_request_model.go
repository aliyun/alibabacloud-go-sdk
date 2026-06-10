// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialTaskListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *QueryMaterialTaskListShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *QueryMaterialTaskListShrinkRequest
	GetNextToken() *string
	SetOrderColumn(v string) *QueryMaterialTaskListShrinkRequest
	GetOrderColumn() *string
	SetOrderType(v string) *QueryMaterialTaskListShrinkRequest
	GetOrderType() *string
	SetPageNum(v int32) *QueryMaterialTaskListShrinkRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryMaterialTaskListShrinkRequest
	GetPageSize() *int32
	SetStatusListShrink(v string) *QueryMaterialTaskListShrinkRequest
	GetStatusListShrink() *string
	SetTaskTypeListShrink(v string) *QueryMaterialTaskListShrinkRequest
	GetTaskTypeListShrink() *string
}

type QueryMaterialTaskListShrinkRequest struct {
	// Number of results per query.
	//
	// Valid values: 10 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Token indicating the start of the next query. This field is empty if there is no next query.
	//
	// example:
	//
	// FFh3Xqm+JgZ/U9Jyb7wdVr9LWk80Tghn5UZjbcWEVEderBcbVF+Y6PS0i8PpCL4PQZ3e0C9oEH0Asd4tJEuGtkl2WuKdiWZpEwadNydQdJPFM=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Sorting field.
	//
	// example:
	//
	// gmtCreated
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// Sorting type: ASC or DESC
	//
	// example:
	//
	// DESC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// Page number. Default value is 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Page size. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// List of task statuses.
	StatusListShrink *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	// List of task types.
	TaskTypeListShrink *string `json:"TaskTypeList,omitempty" xml:"TaskTypeList,omitempty"`
}

func (s QueryMaterialTaskListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialTaskListShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryMaterialTaskListShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryMaterialTaskListShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryMaterialTaskListShrinkRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *QueryMaterialTaskListShrinkRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *QueryMaterialTaskListShrinkRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryMaterialTaskListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMaterialTaskListShrinkRequest) GetStatusListShrink() *string {
	return s.StatusListShrink
}

func (s *QueryMaterialTaskListShrinkRequest) GetTaskTypeListShrink() *string {
	return s.TaskTypeListShrink
}

func (s *QueryMaterialTaskListShrinkRequest) SetMaxResults(v int32) *QueryMaterialTaskListShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryMaterialTaskListShrinkRequest) SetNextToken(v string) *QueryMaterialTaskListShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *QueryMaterialTaskListShrinkRequest) SetOrderColumn(v string) *QueryMaterialTaskListShrinkRequest {
	s.OrderColumn = &v
	return s
}

func (s *QueryMaterialTaskListShrinkRequest) SetOrderType(v string) *QueryMaterialTaskListShrinkRequest {
	s.OrderType = &v
	return s
}

func (s *QueryMaterialTaskListShrinkRequest) SetPageNum(v int32) *QueryMaterialTaskListShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *QueryMaterialTaskListShrinkRequest) SetPageSize(v int32) *QueryMaterialTaskListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMaterialTaskListShrinkRequest) SetStatusListShrink(v string) *QueryMaterialTaskListShrinkRequest {
	s.StatusListShrink = &v
	return s
}

func (s *QueryMaterialTaskListShrinkRequest) SetTaskTypeListShrink(v string) *QueryMaterialTaskListShrinkRequest {
	s.TaskTypeListShrink = &v
	return s
}

func (s *QueryMaterialTaskListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
