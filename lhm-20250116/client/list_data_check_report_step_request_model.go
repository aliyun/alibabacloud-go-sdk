// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckReportStepRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckResult(v int32) *ListDataCheckReportStepRequest
	GetCheckResult() *int32
	SetJobId(v int64) *ListDataCheckReportStepRequest
	GetJobId() *int64
	SetJobStatus(v int32) *ListDataCheckReportStepRequest
	GetJobStatus() *int32
	SetPageIndex(v int32) *ListDataCheckReportStepRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListDataCheckReportStepRequest
	GetPageSize() *int32
}

type ListDataCheckReportStepRequest struct {
	// The verification result filter. Valid values:
	//
	// - 0: no record.
	//
	// - 1: passed.
	//
	// - 2: failed.
	//
	// example:
	//
	// 0
	CheckResult *int32 `json:"checkResult,omitempty" xml:"checkResult,omitempty"`
	// The job database ID (integer) that identifies a verification sub-job. This parameter differs in format from the UUID-format sub-job ID (string) used in the operation that queries step details by UUID. The two are not interchangeable.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	JobId *int64 `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// The step status filter. Valid values:
	//
	// - 0: INIT.
	//
	// - 1: RUNNING.
	//
	// - 2: FINISHED.
	//
	// - 3: STOPPED.
	//
	// - 4: FAIL.
	//
	// - 6: READY.
	//
	// - 7: SKIPPED.
	//
	// example:
	//
	// 0
	JobStatus *int32 `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListDataCheckReportStepRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckReportStepRequest) GoString() string {
	return s.String()
}

func (s *ListDataCheckReportStepRequest) GetCheckResult() *int32 {
	return s.CheckResult
}

func (s *ListDataCheckReportStepRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *ListDataCheckReportStepRequest) GetJobStatus() *int32 {
	return s.JobStatus
}

func (s *ListDataCheckReportStepRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListDataCheckReportStepRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataCheckReportStepRequest) SetCheckResult(v int32) *ListDataCheckReportStepRequest {
	s.CheckResult = &v
	return s
}

func (s *ListDataCheckReportStepRequest) SetJobId(v int64) *ListDataCheckReportStepRequest {
	s.JobId = &v
	return s
}

func (s *ListDataCheckReportStepRequest) SetJobStatus(v int32) *ListDataCheckReportStepRequest {
	s.JobStatus = &v
	return s
}

func (s *ListDataCheckReportStepRequest) SetPageIndex(v int32) *ListDataCheckReportStepRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDataCheckReportStepRequest) SetPageSize(v int32) *ListDataCheckReportStepRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataCheckReportStepRequest) Validate() error {
	return dara.Validate(s)
}
