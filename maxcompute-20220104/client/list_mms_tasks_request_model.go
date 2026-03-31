// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSorter(v *ListMmsTasksRequestSorter) *ListMmsTasksRequest
	GetSorter() *ListMmsTasksRequestSorter
	SetDstDbName(v string) *ListMmsTasksRequest
	GetDstDbName() *string
	SetDstTableName(v string) *ListMmsTasksRequest
	GetDstTableName() *string
	SetJobId(v int64) *ListMmsTasksRequest
	GetJobId() *int64
	SetJobName(v string) *ListMmsTasksRequest
	GetJobName() *string
	SetPageNum(v int32) *ListMmsTasksRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMmsTasksRequest
	GetPageSize() *int32
	SetPartition(v string) *ListMmsTasksRequest
	GetPartition() *string
	SetSrcDbName(v string) *ListMmsTasksRequest
	GetSrcDbName() *string
	SetSrcTableName(v string) *ListMmsTasksRequest
	GetSrcTableName() *string
	SetStatus(v string) *ListMmsTasksRequest
	GetStatus() *string
}

type ListMmsTasksRequest struct {
	Sorter *ListMmsTasksRequestSorter `json:"sorter,omitempty" xml:"sorter,omitempty" type:"Struct"`
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
	// 10
	JobId *int64 `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// test1
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
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
	// p1=1/p2=abc
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
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
	// DATA_DOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListMmsTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTasksRequest) GoString() string {
	return s.String()
}

func (s *ListMmsTasksRequest) GetSorter() *ListMmsTasksRequestSorter {
	return s.Sorter
}

func (s *ListMmsTasksRequest) GetDstDbName() *string {
	return s.DstDbName
}

func (s *ListMmsTasksRequest) GetDstTableName() *string {
	return s.DstTableName
}

func (s *ListMmsTasksRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *ListMmsTasksRequest) GetJobName() *string {
	return s.JobName
}

func (s *ListMmsTasksRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMmsTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmsTasksRequest) GetPartition() *string {
	return s.Partition
}

func (s *ListMmsTasksRequest) GetSrcDbName() *string {
	return s.SrcDbName
}

func (s *ListMmsTasksRequest) GetSrcTableName() *string {
	return s.SrcTableName
}

func (s *ListMmsTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *ListMmsTasksRequest) SetSorter(v *ListMmsTasksRequestSorter) *ListMmsTasksRequest {
	s.Sorter = v
	return s
}

func (s *ListMmsTasksRequest) SetDstDbName(v string) *ListMmsTasksRequest {
	s.DstDbName = &v
	return s
}

func (s *ListMmsTasksRequest) SetDstTableName(v string) *ListMmsTasksRequest {
	s.DstTableName = &v
	return s
}

func (s *ListMmsTasksRequest) SetJobId(v int64) *ListMmsTasksRequest {
	s.JobId = &v
	return s
}

func (s *ListMmsTasksRequest) SetJobName(v string) *ListMmsTasksRequest {
	s.JobName = &v
	return s
}

func (s *ListMmsTasksRequest) SetPageNum(v int32) *ListMmsTasksRequest {
	s.PageNum = &v
	return s
}

func (s *ListMmsTasksRequest) SetPageSize(v int32) *ListMmsTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmsTasksRequest) SetPartition(v string) *ListMmsTasksRequest {
	s.Partition = &v
	return s
}

func (s *ListMmsTasksRequest) SetSrcDbName(v string) *ListMmsTasksRequest {
	s.SrcDbName = &v
	return s
}

func (s *ListMmsTasksRequest) SetSrcTableName(v string) *ListMmsTasksRequest {
	s.SrcTableName = &v
	return s
}

func (s *ListMmsTasksRequest) SetStatus(v string) *ListMmsTasksRequest {
	s.Status = &v
	return s
}

func (s *ListMmsTasksRequest) Validate() error {
	if s.Sorter != nil {
		if err := s.Sorter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMmsTasksRequestSorter struct {
	// example:
	//
	// desc
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// asc
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListMmsTasksRequestSorter) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTasksRequestSorter) GoString() string {
	return s.String()
}

func (s *ListMmsTasksRequestSorter) GetStartTime() *string {
	return s.StartTime
}

func (s *ListMmsTasksRequestSorter) GetStatus() *string {
	return s.Status
}

func (s *ListMmsTasksRequestSorter) SetStartTime(v string) *ListMmsTasksRequestSorter {
	s.StartTime = &v
	return s
}

func (s *ListMmsTasksRequestSorter) SetStatus(v string) *ListMmsTasksRequestSorter {
	s.Status = &v
	return s
}

func (s *ListMmsTasksRequestSorter) Validate() error {
	return dara.Validate(s)
}
