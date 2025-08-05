// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsPartitionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSorter(v *ListMmsPartitionsShrinkRequestSorter) *ListMmsPartitionsShrinkRequest
	GetSorter() *ListMmsPartitionsShrinkRequestSorter
	SetDbId(v int64) *ListMmsPartitionsShrinkRequest
	GetDbId() *int64
	SetDbName(v string) *ListMmsPartitionsShrinkRequest
	GetDbName() *string
	SetLastDdlTimeEnd(v string) *ListMmsPartitionsShrinkRequest
	GetLastDdlTimeEnd() *string
	SetLastDdlTimeStart(v string) *ListMmsPartitionsShrinkRequest
	GetLastDdlTimeStart() *string
	SetPageNum(v int32) *ListMmsPartitionsShrinkRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMmsPartitionsShrinkRequest
	GetPageSize() *int32
	SetStatusShrink(v string) *ListMmsPartitionsShrinkRequest
	GetStatusShrink() *string
	SetTableName(v string) *ListMmsPartitionsShrinkRequest
	GetTableName() *string
	SetUpdated(v bool) *ListMmsPartitionsShrinkRequest
	GetUpdated() *bool
	SetValue(v string) *ListMmsPartitionsShrinkRequest
	GetValue() *string
}

type ListMmsPartitionsShrinkRequest struct {
	Sorter *ListMmsPartitionsShrinkRequestSorter `json:"sorter,omitempty" xml:"sorter,omitempty" type:"Struct"`
	// example:
	//
	// 2
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// example:
	//
	// d1
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// example:
	//
	// 2024-12-17 19:44:42
	LastDdlTimeEnd *string `json:"lastDdlTimeEnd,omitempty" xml:"lastDdlTimeEnd,omitempty"`
	// example:
	//
	// 2024-12-17 15:44:42
	LastDdlTimeStart *string `json:"lastDdlTimeStart,omitempty" xml:"lastDdlTimeStart,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 100
	PageSize     *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StatusShrink *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// t1
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// example:
	//
	// false
	Updated *bool `json:"updated,omitempty" xml:"updated,omitempty"`
	// example:
	//
	// p1=1/p2=abc
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMmsPartitionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMmsPartitionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMmsPartitionsShrinkRequest) GetSorter() *ListMmsPartitionsShrinkRequestSorter {
	return s.Sorter
}

func (s *ListMmsPartitionsShrinkRequest) GetDbId() *int64 {
	return s.DbId
}

func (s *ListMmsPartitionsShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListMmsPartitionsShrinkRequest) GetLastDdlTimeEnd() *string {
	return s.LastDdlTimeEnd
}

func (s *ListMmsPartitionsShrinkRequest) GetLastDdlTimeStart() *string {
	return s.LastDdlTimeStart
}

func (s *ListMmsPartitionsShrinkRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMmsPartitionsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmsPartitionsShrinkRequest) GetStatusShrink() *string {
	return s.StatusShrink
}

func (s *ListMmsPartitionsShrinkRequest) GetTableName() *string {
	return s.TableName
}

func (s *ListMmsPartitionsShrinkRequest) GetUpdated() *bool {
	return s.Updated
}

func (s *ListMmsPartitionsShrinkRequest) GetValue() *string {
	return s.Value
}

func (s *ListMmsPartitionsShrinkRequest) SetSorter(v *ListMmsPartitionsShrinkRequestSorter) *ListMmsPartitionsShrinkRequest {
	s.Sorter = v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) SetDbId(v int64) *ListMmsPartitionsShrinkRequest {
	s.DbId = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) SetDbName(v string) *ListMmsPartitionsShrinkRequest {
	s.DbName = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) SetLastDdlTimeEnd(v string) *ListMmsPartitionsShrinkRequest {
	s.LastDdlTimeEnd = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) SetLastDdlTimeStart(v string) *ListMmsPartitionsShrinkRequest {
	s.LastDdlTimeStart = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) SetPageNum(v int32) *ListMmsPartitionsShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) SetPageSize(v int32) *ListMmsPartitionsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) SetStatusShrink(v string) *ListMmsPartitionsShrinkRequest {
	s.StatusShrink = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) SetTableName(v string) *ListMmsPartitionsShrinkRequest {
	s.TableName = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) SetUpdated(v bool) *ListMmsPartitionsShrinkRequest {
	s.Updated = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) SetValue(v string) *ListMmsPartitionsShrinkRequest {
	s.Value = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type ListMmsPartitionsShrinkRequestSorter struct {
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

func (s ListMmsPartitionsShrinkRequestSorter) String() string {
	return dara.Prettify(s)
}

func (s ListMmsPartitionsShrinkRequestSorter) GoString() string {
	return s.String()
}

func (s *ListMmsPartitionsShrinkRequestSorter) GetLastDdlTime() *string {
	return s.LastDdlTime
}

func (s *ListMmsPartitionsShrinkRequestSorter) GetNumRows() *string {
	return s.NumRows
}

func (s *ListMmsPartitionsShrinkRequestSorter) GetSize() *string {
	return s.Size
}

func (s *ListMmsPartitionsShrinkRequestSorter) SetLastDdlTime(v string) *ListMmsPartitionsShrinkRequestSorter {
	s.LastDdlTime = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequestSorter) SetNumRows(v string) *ListMmsPartitionsShrinkRequestSorter {
	s.NumRows = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequestSorter) SetSize(v string) *ListMmsPartitionsShrinkRequestSorter {
	s.Size = &v
	return s
}

func (s *ListMmsPartitionsShrinkRequestSorter) Validate() error {
	return dara.Validate(s)
}
