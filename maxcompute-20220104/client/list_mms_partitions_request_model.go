// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsPartitionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSorter(v *ListMmsPartitionsRequestSorter) *ListMmsPartitionsRequest
	GetSorter() *ListMmsPartitionsRequestSorter
	SetDbId(v int64) *ListMmsPartitionsRequest
	GetDbId() *int64
	SetDbName(v string) *ListMmsPartitionsRequest
	GetDbName() *string
	SetLastDdlTimeEnd(v string) *ListMmsPartitionsRequest
	GetLastDdlTimeEnd() *string
	SetLastDdlTimeStart(v string) *ListMmsPartitionsRequest
	GetLastDdlTimeStart() *string
	SetPageNum(v int32) *ListMmsPartitionsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMmsPartitionsRequest
	GetPageSize() *int32
	SetStatus(v []*string) *ListMmsPartitionsRequest
	GetStatus() []*string
	SetTableId(v int64) *ListMmsPartitionsRequest
	GetTableId() *int64
	SetTableName(v string) *ListMmsPartitionsRequest
	GetTableName() *string
	SetUpdated(v bool) *ListMmsPartitionsRequest
	GetUpdated() *bool
	SetValue(v string) *ListMmsPartitionsRequest
	GetValue() *string
}

type ListMmsPartitionsRequest struct {
	Sorter *ListMmsPartitionsRequestSorter `json:"sorter,omitempty" xml:"sorter,omitempty" type:"Struct"`
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
	PageSize *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Status   []*string `json:"status,omitempty" xml:"status,omitempty" type:"Repeated"`
	// example:
	//
	// 20323
	TableId *int64 `json:"tableId,omitempty" xml:"tableId,omitempty"`
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

func (s ListMmsPartitionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMmsPartitionsRequest) GoString() string {
	return s.String()
}

func (s *ListMmsPartitionsRequest) GetSorter() *ListMmsPartitionsRequestSorter {
	return s.Sorter
}

func (s *ListMmsPartitionsRequest) GetDbId() *int64 {
	return s.DbId
}

func (s *ListMmsPartitionsRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListMmsPartitionsRequest) GetLastDdlTimeEnd() *string {
	return s.LastDdlTimeEnd
}

func (s *ListMmsPartitionsRequest) GetLastDdlTimeStart() *string {
	return s.LastDdlTimeStart
}

func (s *ListMmsPartitionsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMmsPartitionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmsPartitionsRequest) GetStatus() []*string {
	return s.Status
}

func (s *ListMmsPartitionsRequest) GetTableId() *int64 {
	return s.TableId
}

func (s *ListMmsPartitionsRequest) GetTableName() *string {
	return s.TableName
}

func (s *ListMmsPartitionsRequest) GetUpdated() *bool {
	return s.Updated
}

func (s *ListMmsPartitionsRequest) GetValue() *string {
	return s.Value
}

func (s *ListMmsPartitionsRequest) SetSorter(v *ListMmsPartitionsRequestSorter) *ListMmsPartitionsRequest {
	s.Sorter = v
	return s
}

func (s *ListMmsPartitionsRequest) SetDbId(v int64) *ListMmsPartitionsRequest {
	s.DbId = &v
	return s
}

func (s *ListMmsPartitionsRequest) SetDbName(v string) *ListMmsPartitionsRequest {
	s.DbName = &v
	return s
}

func (s *ListMmsPartitionsRequest) SetLastDdlTimeEnd(v string) *ListMmsPartitionsRequest {
	s.LastDdlTimeEnd = &v
	return s
}

func (s *ListMmsPartitionsRequest) SetLastDdlTimeStart(v string) *ListMmsPartitionsRequest {
	s.LastDdlTimeStart = &v
	return s
}

func (s *ListMmsPartitionsRequest) SetPageNum(v int32) *ListMmsPartitionsRequest {
	s.PageNum = &v
	return s
}

func (s *ListMmsPartitionsRequest) SetPageSize(v int32) *ListMmsPartitionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmsPartitionsRequest) SetStatus(v []*string) *ListMmsPartitionsRequest {
	s.Status = v
	return s
}

func (s *ListMmsPartitionsRequest) SetTableId(v int64) *ListMmsPartitionsRequest {
	s.TableId = &v
	return s
}

func (s *ListMmsPartitionsRequest) SetTableName(v string) *ListMmsPartitionsRequest {
	s.TableName = &v
	return s
}

func (s *ListMmsPartitionsRequest) SetUpdated(v bool) *ListMmsPartitionsRequest {
	s.Updated = &v
	return s
}

func (s *ListMmsPartitionsRequest) SetValue(v string) *ListMmsPartitionsRequest {
	s.Value = &v
	return s
}

func (s *ListMmsPartitionsRequest) Validate() error {
	if s.Sorter != nil {
		if err := s.Sorter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMmsPartitionsRequestSorter struct {
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

func (s ListMmsPartitionsRequestSorter) String() string {
	return dara.Prettify(s)
}

func (s ListMmsPartitionsRequestSorter) GoString() string {
	return s.String()
}

func (s *ListMmsPartitionsRequestSorter) GetLastDdlTime() *string {
	return s.LastDdlTime
}

func (s *ListMmsPartitionsRequestSorter) GetNumRows() *string {
	return s.NumRows
}

func (s *ListMmsPartitionsRequestSorter) GetSize() *string {
	return s.Size
}

func (s *ListMmsPartitionsRequestSorter) SetLastDdlTime(v string) *ListMmsPartitionsRequestSorter {
	s.LastDdlTime = &v
	return s
}

func (s *ListMmsPartitionsRequestSorter) SetNumRows(v string) *ListMmsPartitionsRequestSorter {
	s.NumRows = &v
	return s
}

func (s *ListMmsPartitionsRequestSorter) SetSize(v string) *ListMmsPartitionsRequestSorter {
	s.Size = &v
	return s
}

func (s *ListMmsPartitionsRequestSorter) Validate() error {
	return dara.Validate(s)
}
