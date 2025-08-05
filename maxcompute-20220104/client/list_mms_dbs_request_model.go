// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsDbsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListMmsDbsRequest
	GetName() *string
	SetPageNum(v int32) *ListMmsDbsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMmsDbsRequest
	GetPageSize() *int32
	SetSorter(v *ListMmsDbsRequestSorter) *ListMmsDbsRequest
	GetSorter() *ListMmsDbsRequestSorter
	SetStatus(v string) *ListMmsDbsRequest
	GetStatus() *string
}

type ListMmsDbsRequest struct {
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Sorter   *ListMmsDbsRequestSorter `json:"sorter,omitempty" xml:"sorter,omitempty" type:"Struct"`
	// example:
	//
	// STARTED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListMmsDbsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMmsDbsRequest) GoString() string {
	return s.String()
}

func (s *ListMmsDbsRequest) GetName() *string {
	return s.Name
}

func (s *ListMmsDbsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMmsDbsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmsDbsRequest) GetSorter() *ListMmsDbsRequestSorter {
	return s.Sorter
}

func (s *ListMmsDbsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListMmsDbsRequest) SetName(v string) *ListMmsDbsRequest {
	s.Name = &v
	return s
}

func (s *ListMmsDbsRequest) SetPageNum(v int32) *ListMmsDbsRequest {
	s.PageNum = &v
	return s
}

func (s *ListMmsDbsRequest) SetPageSize(v int32) *ListMmsDbsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmsDbsRequest) SetSorter(v *ListMmsDbsRequestSorter) *ListMmsDbsRequest {
	s.Sorter = v
	return s
}

func (s *ListMmsDbsRequest) SetStatus(v string) *ListMmsDbsRequest {
	s.Status = &v
	return s
}

func (s *ListMmsDbsRequest) Validate() error {
	return dara.Validate(s)
}

type ListMmsDbsRequestSorter struct {
	// example:
	//
	// desc
	NumRows *string `json:"numRows,omitempty" xml:"numRows,omitempty"`
	// example:
	//
	// asc
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// 2024-12-17 15:44:17
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListMmsDbsRequestSorter) String() string {
	return dara.Prettify(s)
}

func (s ListMmsDbsRequestSorter) GoString() string {
	return s.String()
}

func (s *ListMmsDbsRequestSorter) GetNumRows() *string {
	return s.NumRows
}

func (s *ListMmsDbsRequestSorter) GetSize() *string {
	return s.Size
}

func (s *ListMmsDbsRequestSorter) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListMmsDbsRequestSorter) SetNumRows(v string) *ListMmsDbsRequestSorter {
	s.NumRows = &v
	return s
}

func (s *ListMmsDbsRequestSorter) SetSize(v string) *ListMmsDbsRequestSorter {
	s.Size = &v
	return s
}

func (s *ListMmsDbsRequestSorter) SetUpdateTime(v string) *ListMmsDbsRequestSorter {
	s.UpdateTime = &v
	return s
}

func (s *ListMmsDbsRequestSorter) Validate() error {
	return dara.Validate(s)
}
