// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSorter(v *ListMmsTablesRequestSorter) *ListMmsTablesRequest
	GetSorter() *ListMmsTablesRequestSorter
	SetDbId(v int64) *ListMmsTablesRequest
	GetDbId() *int64
	SetDbName(v string) *ListMmsTablesRequest
	GetDbName() *string
	SetDstName(v string) *ListMmsTablesRequest
	GetDstName() *string
	SetDstProjectName(v string) *ListMmsTablesRequest
	GetDstProjectName() *string
	SetDstSchemaName(v string) *ListMmsTablesRequest
	GetDstSchemaName() *string
	SetHasPartitions(v bool) *ListMmsTablesRequest
	GetHasPartitions() *bool
	SetLastDdlTimeEnd(v string) *ListMmsTablesRequest
	GetLastDdlTimeEnd() *string
	SetLastDdlTimeStart(v string) *ListMmsTablesRequest
	GetLastDdlTimeStart() *string
	SetName(v string) *ListMmsTablesRequest
	GetName() *string
	SetOnlyName(v bool) *ListMmsTablesRequest
	GetOnlyName() *bool
	SetPageNum(v int32) *ListMmsTablesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMmsTablesRequest
	GetPageSize() *int32
	SetStatus(v []*string) *ListMmsTablesRequest
	GetStatus() []*string
	SetType(v string) *ListMmsTablesRequest
	GetType() *string
}

type ListMmsTablesRequest struct {
	Sorter *ListMmsTablesRequestSorter `json:"sorter,omitempty" xml:"sorter,omitempty" type:"Struct"`
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
	// test
	DstName *string `json:"dstName,omitempty" xml:"dstName,omitempty"`
	// example:
	//
	// mms_test
	DstProjectName *string `json:"dstProjectName,omitempty" xml:"dstProjectName,omitempty"`
	// example:
	//
	// default
	DstSchemaName *string `json:"dstSchemaName,omitempty" xml:"dstSchemaName,omitempty"`
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
	PageSize *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Status   []*string `json:"status,omitempty" xml:"status,omitempty" type:"Repeated"`
	// example:
	//
	// MANAGED_TABLE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMmsTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTablesRequest) GoString() string {
	return s.String()
}

func (s *ListMmsTablesRequest) GetSorter() *ListMmsTablesRequestSorter {
	return s.Sorter
}

func (s *ListMmsTablesRequest) GetDbId() *int64 {
	return s.DbId
}

func (s *ListMmsTablesRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListMmsTablesRequest) GetDstName() *string {
	return s.DstName
}

func (s *ListMmsTablesRequest) GetDstProjectName() *string {
	return s.DstProjectName
}

func (s *ListMmsTablesRequest) GetDstSchemaName() *string {
	return s.DstSchemaName
}

func (s *ListMmsTablesRequest) GetHasPartitions() *bool {
	return s.HasPartitions
}

func (s *ListMmsTablesRequest) GetLastDdlTimeEnd() *string {
	return s.LastDdlTimeEnd
}

func (s *ListMmsTablesRequest) GetLastDdlTimeStart() *string {
	return s.LastDdlTimeStart
}

func (s *ListMmsTablesRequest) GetName() *string {
	return s.Name
}

func (s *ListMmsTablesRequest) GetOnlyName() *bool {
	return s.OnlyName
}

func (s *ListMmsTablesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMmsTablesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmsTablesRequest) GetStatus() []*string {
	return s.Status
}

func (s *ListMmsTablesRequest) GetType() *string {
	return s.Type
}

func (s *ListMmsTablesRequest) SetSorter(v *ListMmsTablesRequestSorter) *ListMmsTablesRequest {
	s.Sorter = v
	return s
}

func (s *ListMmsTablesRequest) SetDbId(v int64) *ListMmsTablesRequest {
	s.DbId = &v
	return s
}

func (s *ListMmsTablesRequest) SetDbName(v string) *ListMmsTablesRequest {
	s.DbName = &v
	return s
}

func (s *ListMmsTablesRequest) SetDstName(v string) *ListMmsTablesRequest {
	s.DstName = &v
	return s
}

func (s *ListMmsTablesRequest) SetDstProjectName(v string) *ListMmsTablesRequest {
	s.DstProjectName = &v
	return s
}

func (s *ListMmsTablesRequest) SetDstSchemaName(v string) *ListMmsTablesRequest {
	s.DstSchemaName = &v
	return s
}

func (s *ListMmsTablesRequest) SetHasPartitions(v bool) *ListMmsTablesRequest {
	s.HasPartitions = &v
	return s
}

func (s *ListMmsTablesRequest) SetLastDdlTimeEnd(v string) *ListMmsTablesRequest {
	s.LastDdlTimeEnd = &v
	return s
}

func (s *ListMmsTablesRequest) SetLastDdlTimeStart(v string) *ListMmsTablesRequest {
	s.LastDdlTimeStart = &v
	return s
}

func (s *ListMmsTablesRequest) SetName(v string) *ListMmsTablesRequest {
	s.Name = &v
	return s
}

func (s *ListMmsTablesRequest) SetOnlyName(v bool) *ListMmsTablesRequest {
	s.OnlyName = &v
	return s
}

func (s *ListMmsTablesRequest) SetPageNum(v int32) *ListMmsTablesRequest {
	s.PageNum = &v
	return s
}

func (s *ListMmsTablesRequest) SetPageSize(v int32) *ListMmsTablesRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmsTablesRequest) SetStatus(v []*string) *ListMmsTablesRequest {
	s.Status = v
	return s
}

func (s *ListMmsTablesRequest) SetType(v string) *ListMmsTablesRequest {
	s.Type = &v
	return s
}

func (s *ListMmsTablesRequest) Validate() error {
	if s.Sorter != nil {
		if err := s.Sorter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMmsTablesRequestSorter struct {
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

func (s ListMmsTablesRequestSorter) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTablesRequestSorter) GoString() string {
	return s.String()
}

func (s *ListMmsTablesRequestSorter) GetLastDdlTime() *string {
	return s.LastDdlTime
}

func (s *ListMmsTablesRequestSorter) GetNumRows() *string {
	return s.NumRows
}

func (s *ListMmsTablesRequestSorter) GetSize() *string {
	return s.Size
}

func (s *ListMmsTablesRequestSorter) SetLastDdlTime(v string) *ListMmsTablesRequestSorter {
	s.LastDdlTime = &v
	return s
}

func (s *ListMmsTablesRequestSorter) SetNumRows(v string) *ListMmsTablesRequestSorter {
	s.NumRows = &v
	return s
}

func (s *ListMmsTablesRequestSorter) SetSize(v string) *ListMmsTablesRequestSorter {
	s.Size = &v
	return s
}

func (s *ListMmsTablesRequestSorter) Validate() error {
	return dara.Validate(s)
}
