// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSorter(v *ListMmsJobsRequestSorter) *ListMmsJobsRequest
	GetSorter() *ListMmsJobsRequestSorter
	SetDstDbName(v string) *ListMmsJobsRequest
	GetDstDbName() *string
	SetDstTableName(v string) *ListMmsJobsRequest
	GetDstTableName() *string
	SetName(v string) *ListMmsJobsRequest
	GetName() *string
	SetPageNum(v int32) *ListMmsJobsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMmsJobsRequest
	GetPageSize() *int32
	SetSrcDbName(v string) *ListMmsJobsRequest
	GetSrcDbName() *string
	SetSrcTableName(v string) *ListMmsJobsRequest
	GetSrcTableName() *string
	SetStatus(v string) *ListMmsJobsRequest
	GetStatus() *string
	SetStopped(v int64) *ListMmsJobsRequest
	GetStopped() *int64
}

type ListMmsJobsRequest struct {
	Sorter *ListMmsJobsRequestSorter `json:"sorter,omitempty" xml:"sorter,omitempty" type:"Struct"`
	// example:
	//
	// mms_test
	DstDbName *string `json:"dstDbName,omitempty" xml:"dstDbName,omitempty"`
	// example:
	//
	// test_table_1
	DstTableName *string `json:"dstTableName,omitempty" xml:"dstTableName,omitempty"`
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
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// test_db_1
	SrcDbName *string `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	// example:
	//
	// test_table_1
	SrcTableName *string `json:"srcTableName,omitempty" xml:"srcTableName,omitempty"`
	// example:
	//
	// DOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// false
	Stopped *int64 `json:"stopped,omitempty" xml:"stopped,omitempty"`
}

func (s ListMmsJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMmsJobsRequest) GoString() string {
	return s.String()
}

func (s *ListMmsJobsRequest) GetSorter() *ListMmsJobsRequestSorter {
	return s.Sorter
}

func (s *ListMmsJobsRequest) GetDstDbName() *string {
	return s.DstDbName
}

func (s *ListMmsJobsRequest) GetDstTableName() *string {
	return s.DstTableName
}

func (s *ListMmsJobsRequest) GetName() *string {
	return s.Name
}

func (s *ListMmsJobsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMmsJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmsJobsRequest) GetSrcDbName() *string {
	return s.SrcDbName
}

func (s *ListMmsJobsRequest) GetSrcTableName() *string {
	return s.SrcTableName
}

func (s *ListMmsJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListMmsJobsRequest) GetStopped() *int64 {
	return s.Stopped
}

func (s *ListMmsJobsRequest) SetSorter(v *ListMmsJobsRequestSorter) *ListMmsJobsRequest {
	s.Sorter = v
	return s
}

func (s *ListMmsJobsRequest) SetDstDbName(v string) *ListMmsJobsRequest {
	s.DstDbName = &v
	return s
}

func (s *ListMmsJobsRequest) SetDstTableName(v string) *ListMmsJobsRequest {
	s.DstTableName = &v
	return s
}

func (s *ListMmsJobsRequest) SetName(v string) *ListMmsJobsRequest {
	s.Name = &v
	return s
}

func (s *ListMmsJobsRequest) SetPageNum(v int32) *ListMmsJobsRequest {
	s.PageNum = &v
	return s
}

func (s *ListMmsJobsRequest) SetPageSize(v int32) *ListMmsJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmsJobsRequest) SetSrcDbName(v string) *ListMmsJobsRequest {
	s.SrcDbName = &v
	return s
}

func (s *ListMmsJobsRequest) SetSrcTableName(v string) *ListMmsJobsRequest {
	s.SrcTableName = &v
	return s
}

func (s *ListMmsJobsRequest) SetStatus(v string) *ListMmsJobsRequest {
	s.Status = &v
	return s
}

func (s *ListMmsJobsRequest) SetStopped(v int64) *ListMmsJobsRequest {
	s.Stopped = &v
	return s
}

func (s *ListMmsJobsRequest) Validate() error {
	if s.Sorter != nil {
		if err := s.Sorter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMmsJobsRequestSorter struct {
	// example:
	//
	// desc
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListMmsJobsRequestSorter) String() string {
	return dara.Prettify(s)
}

func (s ListMmsJobsRequestSorter) GoString() string {
	return s.String()
}

func (s *ListMmsJobsRequestSorter) GetStatus() *string {
	return s.Status
}

func (s *ListMmsJobsRequestSorter) SetStatus(v string) *ListMmsJobsRequestSorter {
	s.Status = &v
	return s
}

func (s *ListMmsJobsRequestSorter) Validate() error {
	return dara.Validate(s)
}
