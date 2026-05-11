// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportResp interface {
	dara.Model
	String() string
	GoString() string
	SetCopiedCount(v int64) *GetReportResp
	GetCopiedCount() *int64
	SetErrorMessage(v string) *GetReportResp
	GetErrorMessage() *string
	SetFailedCount(v int64) *GetReportResp
	GetFailedCount() *int64
	SetFailedListPrefix(v string) *GetReportResp
	GetFailedListPrefix() *string
	SetJobCreateTime(v string) *GetReportResp
	GetJobCreateTime() *string
	SetJobEndTime(v string) *GetReportResp
	GetJobEndTime() *string
	SetJobExecuteTime(v string) *GetReportResp
	GetJobExecuteTime() *string
	SetReportCreateTime(v string) *GetReportResp
	GetReportCreateTime() *string
	SetReportEndTime(v string) *GetReportResp
	GetReportEndTime() *string
	SetSkippedCount(v int64) *GetReportResp
	GetSkippedCount() *int64
	SetSkippedListPrefix(v string) *GetReportResp
	GetSkippedListPrefix() *string
	SetStatus(v string) *GetReportResp
	GetStatus() *string
	SetTotalCount(v int64) *GetReportResp
	GetTotalCount() *int64
	SetTotalListPrefix(v string) *GetReportResp
	GetTotalListPrefix() *string
}

type GetReportResp struct {
	// The number of files that are migrated.
	//
	// example:
	//
	// 800
	CopiedCount *int64 `json:"CopiedCount,omitempty" xml:"CopiedCount,omitempty"`
	// The error message.
	//
	// example:
	//
	// test error msg.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The number of files that failed to be migrated.
	//
	// example:
	//
	// 100
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The path of files that failed to be migrated.
	//
	// example:
	//
	// test_failed_prefix/
	FailedListPrefix *string `json:"FailedListPrefix,omitempty" xml:"FailedListPrefix,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2024-05-01 12:00:00
	JobCreateTime *string `json:"JobCreateTime,omitempty" xml:"JobCreateTime,omitempty"`
	// The time when the task ended.
	//
	// example:
	//
	// 2024-05-01 12:00:00
	JobEndTime *string `json:"JobEndTime,omitempty" xml:"JobEndTime,omitempty"`
	// The task duration.
	//
	// example:
	//
	// 1000
	JobExecuteTime *string `json:"JobExecuteTime,omitempty" xml:"JobExecuteTime,omitempty"`
	// The time when the migration report was created.
	//
	// example:
	//
	// 2024-05-01 12:00:00
	ReportCreateTime *string `json:"ReportCreateTime,omitempty" xml:"ReportCreateTime,omitempty"`
	// The time when the migration report was generated.
	//
	// example:
	//
	// 2024-05-01 12:00:00
	ReportEndTime *string `json:"ReportEndTime,omitempty" xml:"ReportEndTime,omitempty"`
	// The number of files that are skipped.
	//
	// example:
	//
	// 100
	SkippedCount *int64 `json:"SkippedCount,omitempty" xml:"SkippedCount,omitempty"`
	// The path of files that are skipped.
	//
	// example:
	//
	// test_skipped_prefix/
	SkippedListPrefix *string `json:"SkippedListPrefix,omitempty" xml:"SkippedListPrefix,omitempty"`
	// The state of the migration report.\\
	//
	// Valid values: Running, NotExist, Finished, and Failed.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of files.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The path of files.
	//
	// example:
	//
	// test_total_prefix/
	TotalListPrefix *string `json:"TotalListPrefix,omitempty" xml:"TotalListPrefix,omitempty"`
}

func (s GetReportResp) String() string {
	return dara.Prettify(s)
}

func (s GetReportResp) GoString() string {
	return s.String()
}

func (s *GetReportResp) GetCopiedCount() *int64 {
	return s.CopiedCount
}

func (s *GetReportResp) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetReportResp) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *GetReportResp) GetFailedListPrefix() *string {
	return s.FailedListPrefix
}

func (s *GetReportResp) GetJobCreateTime() *string {
	return s.JobCreateTime
}

func (s *GetReportResp) GetJobEndTime() *string {
	return s.JobEndTime
}

func (s *GetReportResp) GetJobExecuteTime() *string {
	return s.JobExecuteTime
}

func (s *GetReportResp) GetReportCreateTime() *string {
	return s.ReportCreateTime
}

func (s *GetReportResp) GetReportEndTime() *string {
	return s.ReportEndTime
}

func (s *GetReportResp) GetSkippedCount() *int64 {
	return s.SkippedCount
}

func (s *GetReportResp) GetSkippedListPrefix() *string {
	return s.SkippedListPrefix
}

func (s *GetReportResp) GetStatus() *string {
	return s.Status
}

func (s *GetReportResp) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetReportResp) GetTotalListPrefix() *string {
	return s.TotalListPrefix
}

func (s *GetReportResp) SetCopiedCount(v int64) *GetReportResp {
	s.CopiedCount = &v
	return s
}

func (s *GetReportResp) SetErrorMessage(v string) *GetReportResp {
	s.ErrorMessage = &v
	return s
}

func (s *GetReportResp) SetFailedCount(v int64) *GetReportResp {
	s.FailedCount = &v
	return s
}

func (s *GetReportResp) SetFailedListPrefix(v string) *GetReportResp {
	s.FailedListPrefix = &v
	return s
}

func (s *GetReportResp) SetJobCreateTime(v string) *GetReportResp {
	s.JobCreateTime = &v
	return s
}

func (s *GetReportResp) SetJobEndTime(v string) *GetReportResp {
	s.JobEndTime = &v
	return s
}

func (s *GetReportResp) SetJobExecuteTime(v string) *GetReportResp {
	s.JobExecuteTime = &v
	return s
}

func (s *GetReportResp) SetReportCreateTime(v string) *GetReportResp {
	s.ReportCreateTime = &v
	return s
}

func (s *GetReportResp) SetReportEndTime(v string) *GetReportResp {
	s.ReportEndTime = &v
	return s
}

func (s *GetReportResp) SetSkippedCount(v int64) *GetReportResp {
	s.SkippedCount = &v
	return s
}

func (s *GetReportResp) SetSkippedListPrefix(v string) *GetReportResp {
	s.SkippedListPrefix = &v
	return s
}

func (s *GetReportResp) SetStatus(v string) *GetReportResp {
	s.Status = &v
	return s
}

func (s *GetReportResp) SetTotalCount(v int64) *GetReportResp {
	s.TotalCount = &v
	return s
}

func (s *GetReportResp) SetTotalListPrefix(v string) *GetReportResp {
	s.TotalListPrefix = &v
	return s
}

func (s *GetReportResp) Validate() error {
	return dara.Validate(s)
}
