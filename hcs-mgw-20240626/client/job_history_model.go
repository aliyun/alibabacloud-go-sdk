// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobHistory interface {
	dara.Model
	String() string
	GoString() string
	SetCommitId(v string) *JobHistory
	GetCommitId() *string
	SetCopiedCount(v int64) *JobHistory
	GetCopiedCount() *int64
	SetCopiedSize(v int64) *JobHistory
	GetCopiedSize() *int64
	SetEndTime(v string) *JobHistory
	GetEndTime() *string
	SetFailedCount(v int64) *JobHistory
	GetFailedCount() *int64
	SetJobVersion(v string) *JobHistory
	GetJobVersion() *string
	SetListStatus(v string) *JobHistory
	GetListStatus() *string
	SetMessage(v string) *JobHistory
	GetMessage() *string
	SetName(v string) *JobHistory
	GetName() *string
	SetOperator(v string) *JobHistory
	GetOperator() *string
	SetRuntimeId(v string) *JobHistory
	GetRuntimeId() *string
	SetRuntimeState(v string) *JobHistory
	GetRuntimeState() *string
	SetSkippedCount(v int64) *JobHistory
	GetSkippedCount() *int64
	SetSkippedSize(v int64) *JobHistory
	GetSkippedSize() *int64
	SetStartTime(v string) *JobHistory
	GetStartTime() *string
	SetStatus(v string) *JobHistory
	GetStatus() *string
	SetTotalCount(v int64) *JobHistory
	GetTotalCount() *int64
	SetTotalSize(v int64) *JobHistory
	GetTotalSize() *int64
}

type JobHistory struct {
	// example:
	//
	// 2
	CommitId *string `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	// example:
	//
	// 900
	CopiedCount *int64 `json:"CopiedCount,omitempty" xml:"CopiedCount,omitempty"`
	// example:
	//
	// 1000
	CopiedSize *int64 `json:"CopiedSize,omitempty" xml:"CopiedSize,omitempty"`
	// example:
	//
	// 2025-07-04T08:09:26.000Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 100
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// example:
	//
	// ****-188f-***-b266-****
	JobVersion *string `json:"JobVersion,omitempty" xml:"JobVersion,omitempty"`
	// example:
	//
	// Listing
	ListStatus *string `json:"ListStatus,omitempty" xml:"ListStatus,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// <your-job-name>
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// user
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 1
	RuntimeId *string `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
	// example:
	//
	// Normal
	RuntimeState *string `json:"RuntimeState,omitempty" xml:"RuntimeState,omitempty"`
	// example:
	//
	// 1000
	SkippedCount *int64 `json:"SkippedCount,omitempty" xml:"SkippedCount,omitempty"`
	// example:
	//
	// 100000
	SkippedSize *int64 `json:"SkippedSize,omitempty" xml:"SkippedSize,omitempty"`
	// example:
	//
	// 2025-07-04T07:09:26.000Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// IMPORT_JOB_DOING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 1000
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s JobHistory) String() string {
	return dara.Prettify(s)
}

func (s JobHistory) GoString() string {
	return s.String()
}

func (s *JobHistory) GetCommitId() *string {
	return s.CommitId
}

func (s *JobHistory) GetCopiedCount() *int64 {
	return s.CopiedCount
}

func (s *JobHistory) GetCopiedSize() *int64 {
	return s.CopiedSize
}

func (s *JobHistory) GetEndTime() *string {
	return s.EndTime
}

func (s *JobHistory) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *JobHistory) GetJobVersion() *string {
	return s.JobVersion
}

func (s *JobHistory) GetListStatus() *string {
	return s.ListStatus
}

func (s *JobHistory) GetMessage() *string {
	return s.Message
}

func (s *JobHistory) GetName() *string {
	return s.Name
}

func (s *JobHistory) GetOperator() *string {
	return s.Operator
}

func (s *JobHistory) GetRuntimeId() *string {
	return s.RuntimeId
}

func (s *JobHistory) GetRuntimeState() *string {
	return s.RuntimeState
}

func (s *JobHistory) GetSkippedCount() *int64 {
	return s.SkippedCount
}

func (s *JobHistory) GetSkippedSize() *int64 {
	return s.SkippedSize
}

func (s *JobHistory) GetStartTime() *string {
	return s.StartTime
}

func (s *JobHistory) GetStatus() *string {
	return s.Status
}

func (s *JobHistory) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *JobHistory) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *JobHistory) SetCommitId(v string) *JobHistory {
	s.CommitId = &v
	return s
}

func (s *JobHistory) SetCopiedCount(v int64) *JobHistory {
	s.CopiedCount = &v
	return s
}

func (s *JobHistory) SetCopiedSize(v int64) *JobHistory {
	s.CopiedSize = &v
	return s
}

func (s *JobHistory) SetEndTime(v string) *JobHistory {
	s.EndTime = &v
	return s
}

func (s *JobHistory) SetFailedCount(v int64) *JobHistory {
	s.FailedCount = &v
	return s
}

func (s *JobHistory) SetJobVersion(v string) *JobHistory {
	s.JobVersion = &v
	return s
}

func (s *JobHistory) SetListStatus(v string) *JobHistory {
	s.ListStatus = &v
	return s
}

func (s *JobHistory) SetMessage(v string) *JobHistory {
	s.Message = &v
	return s
}

func (s *JobHistory) SetName(v string) *JobHistory {
	s.Name = &v
	return s
}

func (s *JobHistory) SetOperator(v string) *JobHistory {
	s.Operator = &v
	return s
}

func (s *JobHistory) SetRuntimeId(v string) *JobHistory {
	s.RuntimeId = &v
	return s
}

func (s *JobHistory) SetRuntimeState(v string) *JobHistory {
	s.RuntimeState = &v
	return s
}

func (s *JobHistory) SetSkippedCount(v int64) *JobHistory {
	s.SkippedCount = &v
	return s
}

func (s *JobHistory) SetSkippedSize(v int64) *JobHistory {
	s.SkippedSize = &v
	return s
}

func (s *JobHistory) SetStartTime(v string) *JobHistory {
	s.StartTime = &v
	return s
}

func (s *JobHistory) SetStatus(v string) *JobHistory {
	s.Status = &v
	return s
}

func (s *JobHistory) SetTotalCount(v int64) *JobHistory {
	s.TotalCount = &v
	return s
}

func (s *JobHistory) SetTotalSize(v int64) *JobHistory {
	s.TotalSize = &v
	return s
}

func (s *JobHistory) Validate() error {
	return dara.Validate(s)
}
