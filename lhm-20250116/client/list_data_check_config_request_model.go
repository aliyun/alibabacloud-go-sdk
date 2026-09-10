// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageIndex(v int32) *ListDataCheckConfigRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListDataCheckConfigRequest
	GetPageSize() *int32
	SetSrcTable(v string) *ListDataCheckConfigRequest
	GetSrcTable() *string
	SetTaskId(v int64) *ListDataCheckConfigRequest
	GetTaskId() *int64
}

type ListDataCheckConfigRequest struct {
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// table_demo
	SrcTable *string `json:"srcTable,omitempty" xml:"srcTable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10001
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListDataCheckConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckConfigRequest) GoString() string {
	return s.String()
}

func (s *ListDataCheckConfigRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListDataCheckConfigRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataCheckConfigRequest) GetSrcTable() *string {
	return s.SrcTable
}

func (s *ListDataCheckConfigRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListDataCheckConfigRequest) SetPageIndex(v int32) *ListDataCheckConfigRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDataCheckConfigRequest) SetPageSize(v int32) *ListDataCheckConfigRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataCheckConfigRequest) SetSrcTable(v string) *ListDataCheckConfigRequest {
	s.SrcTable = &v
	return s
}

func (s *ListDataCheckConfigRequest) SetTaskId(v int64) *ListDataCheckConfigRequest {
	s.TaskId = &v
	return s
}

func (s *ListDataCheckConfigRequest) Validate() error {
	return dara.Validate(s)
}
