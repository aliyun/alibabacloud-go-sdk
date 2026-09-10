// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v int64) *ListDataCheckReportRequest
	GetBatchId() *int64
	SetCheckResult(v int32) *ListDataCheckReportRequest
	GetCheckResult() *int32
	SetJobStatus(v int32) *ListDataCheckReportRequest
	GetJobStatus() *int32
	SetPageIndex(v int32) *ListDataCheckReportRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListDataCheckReportRequest
	GetPageSize() *int32
	SetTableName(v string) *ListDataCheckReportRequest
	GetTableName() *string
}

type ListDataCheckReportRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20001
	BatchId *int64 `json:"batchId,omitempty" xml:"batchId,omitempty"`
	// example:
	//
	// 0
	CheckResult *int32 `json:"checkResult,omitempty" xml:"checkResult,omitempty"`
	// example:
	//
	// 0
	JobStatus *int32 `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// dim_func_with_diff_area_data_d
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
}

func (s ListDataCheckReportRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckReportRequest) GoString() string {
	return s.String()
}

func (s *ListDataCheckReportRequest) GetBatchId() *int64 {
	return s.BatchId
}

func (s *ListDataCheckReportRequest) GetCheckResult() *int32 {
	return s.CheckResult
}

func (s *ListDataCheckReportRequest) GetJobStatus() *int32 {
	return s.JobStatus
}

func (s *ListDataCheckReportRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListDataCheckReportRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataCheckReportRequest) GetTableName() *string {
	return s.TableName
}

func (s *ListDataCheckReportRequest) SetBatchId(v int64) *ListDataCheckReportRequest {
	s.BatchId = &v
	return s
}

func (s *ListDataCheckReportRequest) SetCheckResult(v int32) *ListDataCheckReportRequest {
	s.CheckResult = &v
	return s
}

func (s *ListDataCheckReportRequest) SetJobStatus(v int32) *ListDataCheckReportRequest {
	s.JobStatus = &v
	return s
}

func (s *ListDataCheckReportRequest) SetPageIndex(v int32) *ListDataCheckReportRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDataCheckReportRequest) SetPageSize(v int32) *ListDataCheckReportRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataCheckReportRequest) SetTableName(v string) *ListDataCheckReportRequest {
	s.TableName = &v
	return s
}

func (s *ListDataCheckReportRequest) Validate() error {
	return dara.Validate(s)
}
