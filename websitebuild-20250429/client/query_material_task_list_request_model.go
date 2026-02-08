// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *QueryMaterialTaskListRequest
	GetMaxResults() *int32
	SetNextToken(v string) *QueryMaterialTaskListRequest
	GetNextToken() *string
	SetOrderColumn(v string) *QueryMaterialTaskListRequest
	GetOrderColumn() *string
	SetOrderType(v string) *QueryMaterialTaskListRequest
	GetOrderType() *string
	SetPageNum(v int32) *QueryMaterialTaskListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryMaterialTaskListRequest
	GetPageSize() *int32
	SetStatusList(v []*string) *QueryMaterialTaskListRequest
	GetStatusList() []*string
	SetTaskTypeList(v []*string) *QueryMaterialTaskListRequest
	GetTaskTypeList() []*string
}

type QueryMaterialTaskListRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFh3Xqm+JgZ/U9Jyb7wdVr9LWk80Tghn5UZjbcWEVEderBcbVF+Y6PS0i8PpCL4PQZ3e0C9oEH0Asd4tJEuGtkl2WuKdiWZpEwadNydQdJPFM=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// gmtCreated
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// example:
	//
	// DESC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize     *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StatusList   []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	TaskTypeList []*string `json:"TaskTypeList,omitempty" xml:"TaskTypeList,omitempty" type:"Repeated"`
}

func (s QueryMaterialTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialTaskListRequest) GoString() string {
	return s.String()
}

func (s *QueryMaterialTaskListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryMaterialTaskListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryMaterialTaskListRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *QueryMaterialTaskListRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *QueryMaterialTaskListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryMaterialTaskListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMaterialTaskListRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *QueryMaterialTaskListRequest) GetTaskTypeList() []*string {
	return s.TaskTypeList
}

func (s *QueryMaterialTaskListRequest) SetMaxResults(v int32) *QueryMaterialTaskListRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryMaterialTaskListRequest) SetNextToken(v string) *QueryMaterialTaskListRequest {
	s.NextToken = &v
	return s
}

func (s *QueryMaterialTaskListRequest) SetOrderColumn(v string) *QueryMaterialTaskListRequest {
	s.OrderColumn = &v
	return s
}

func (s *QueryMaterialTaskListRequest) SetOrderType(v string) *QueryMaterialTaskListRequest {
	s.OrderType = &v
	return s
}

func (s *QueryMaterialTaskListRequest) SetPageNum(v int32) *QueryMaterialTaskListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryMaterialTaskListRequest) SetPageSize(v int32) *QueryMaterialTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMaterialTaskListRequest) SetStatusList(v []*string) *QueryMaterialTaskListRequest {
	s.StatusList = v
	return s
}

func (s *QueryMaterialTaskListRequest) SetTaskTypeList(v []*string) *QueryMaterialTaskListRequest {
	s.TaskTypeList = v
	return s
}

func (s *QueryMaterialTaskListRequest) Validate() error {
	return dara.Validate(s)
}
