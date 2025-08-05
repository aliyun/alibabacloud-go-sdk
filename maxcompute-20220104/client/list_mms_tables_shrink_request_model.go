// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsTablesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSorter(v *ListMmsTablesShrinkRequestSorter) *ListMmsTablesShrinkRequest
	GetSorter() *ListMmsTablesShrinkRequestSorter
	SetDbId(v int64) *ListMmsTablesShrinkRequest
	GetDbId() *int64
	SetDbName(v string) *ListMmsTablesShrinkRequest
	GetDbName() *string
	SetHasPartitions(v bool) *ListMmsTablesShrinkRequest
	GetHasPartitions() *bool
	SetLastDdlTimeEnd(v string) *ListMmsTablesShrinkRequest
	GetLastDdlTimeEnd() *string
	SetLastDdlTimeStart(v string) *ListMmsTablesShrinkRequest
	GetLastDdlTimeStart() *string
	SetName(v string) *ListMmsTablesShrinkRequest
	GetName() *string
	SetOnlyName(v bool) *ListMmsTablesShrinkRequest
	GetOnlyName() *bool
	SetPageNum(v int32) *ListMmsTablesShrinkRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMmsTablesShrinkRequest
	GetPageSize() *int32
	SetStatusShrink(v string) *ListMmsTablesShrinkRequest
	GetStatusShrink() *string
	SetType(v string) *ListMmsTablesShrinkRequest
	GetType() *string
}

type ListMmsTablesShrinkRequest struct {
	Sorter *ListMmsTablesShrinkRequestSorter `json:"sorter,omitempty" xml:"sorter,omitempty" type:"Struct"`
	// example:
	//
	// 197
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// example:
	//
	// mms_test
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// example:
	//
	// true
	HasPartitions *bool `json:"hasPartitions,omitempty" xml:"hasPartitions,omitempty"`
	// example:
	//
	// 2024-12-19 15:44:42
	LastDdlTimeEnd *string `json:"lastDdlTimeEnd,omitempty" xml:"lastDdlTimeEnd,omitempty"`
	// example:
	//
	// 2024-12-17 15:44:42
	LastDdlTimeStart *string `json:"lastDdlTimeStart,omitempty" xml:"lastDdlTimeStart,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// false
	OnlyName *bool `json:"onlyName,omitempty" xml:"onlyName,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 10
	PageSize     *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StatusShrink *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// MANAGED_TABLE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMmsTablesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTablesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMmsTablesShrinkRequest) GetSorter() *ListMmsTablesShrinkRequestSorter {
	return s.Sorter
}

func (s *ListMmsTablesShrinkRequest) GetDbId() *int64 {
	return s.DbId
}

func (s *ListMmsTablesShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListMmsTablesShrinkRequest) GetHasPartitions() *bool {
	return s.HasPartitions
}

func (s *ListMmsTablesShrinkRequest) GetLastDdlTimeEnd() *string {
	return s.LastDdlTimeEnd
}

func (s *ListMmsTablesShrinkRequest) GetLastDdlTimeStart() *string {
	return s.LastDdlTimeStart
}

func (s *ListMmsTablesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListMmsTablesShrinkRequest) GetOnlyName() *bool {
	return s.OnlyName
}

func (s *ListMmsTablesShrinkRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMmsTablesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmsTablesShrinkRequest) GetStatusShrink() *string {
	return s.StatusShrink
}

func (s *ListMmsTablesShrinkRequest) GetType() *string {
	return s.Type
}

func (s *ListMmsTablesShrinkRequest) SetSorter(v *ListMmsTablesShrinkRequestSorter) *ListMmsTablesShrinkRequest {
	s.Sorter = v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetDbId(v int64) *ListMmsTablesShrinkRequest {
	s.DbId = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetDbName(v string) *ListMmsTablesShrinkRequest {
	s.DbName = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetHasPartitions(v bool) *ListMmsTablesShrinkRequest {
	s.HasPartitions = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetLastDdlTimeEnd(v string) *ListMmsTablesShrinkRequest {
	s.LastDdlTimeEnd = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetLastDdlTimeStart(v string) *ListMmsTablesShrinkRequest {
	s.LastDdlTimeStart = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetName(v string) *ListMmsTablesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetOnlyName(v bool) *ListMmsTablesShrinkRequest {
	s.OnlyName = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetPageNum(v int32) *ListMmsTablesShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetPageSize(v int32) *ListMmsTablesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetStatusShrink(v string) *ListMmsTablesShrinkRequest {
	s.StatusShrink = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) SetType(v string) *ListMmsTablesShrinkRequest {
	s.Type = &v
	return s
}

func (s *ListMmsTablesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type ListMmsTablesShrinkRequestSorter struct {
	// example:
	//
	// desc
	LastDdlTime *string `json:"lastDdlTime,omitempty" xml:"lastDdlTime,omitempty"`
	// example:
	//
	// desc
	NumRows *string `json:"numRows,omitempty" xml:"numRows,omitempty"`
	// example:
	//
	// asc
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListMmsTablesShrinkRequestSorter) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTablesShrinkRequestSorter) GoString() string {
	return s.String()
}

func (s *ListMmsTablesShrinkRequestSorter) GetLastDdlTime() *string {
	return s.LastDdlTime
}

func (s *ListMmsTablesShrinkRequestSorter) GetNumRows() *string {
	return s.NumRows
}

func (s *ListMmsTablesShrinkRequestSorter) GetSize() *string {
	return s.Size
}

func (s *ListMmsTablesShrinkRequestSorter) SetLastDdlTime(v string) *ListMmsTablesShrinkRequestSorter {
	s.LastDdlTime = &v
	return s
}

func (s *ListMmsTablesShrinkRequestSorter) SetNumRows(v string) *ListMmsTablesShrinkRequestSorter {
	s.NumRows = &v
	return s
}

func (s *ListMmsTablesShrinkRequestSorter) SetSize(v string) *ListMmsTablesShrinkRequestSorter {
	s.Size = &v
	return s
}

func (s *ListMmsTablesShrinkRequestSorter) Validate() error {
	return dara.Validate(s)
}
