// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCodeBundleId(v int64) *CreateScanResponseBody
	GetCodeBundleId() *int64
	SetCreatedAt(v string) *CreateScanResponseBody
	GetCreatedAt() *string
	SetCreatedBy(v string) *CreateScanResponseBody
	GetCreatedBy() *string
	SetEngineSnapshot(v *CreateScanResponseBodyEngineSnapshot) *CreateScanResponseBody
	GetEngineSnapshot() *CreateScanResponseBodyEngineSnapshot
	SetFinishedAt(v string) *CreateScanResponseBody
	GetFinishedAt() *string
	SetId(v int64) *CreateScanResponseBody
	GetId() *int64
	SetKind(v string) *CreateScanResponseBody
	GetKind() *string
	SetProjectId(v int64) *CreateScanResponseBody
	GetProjectId() *int64
	SetRequestId(v string) *CreateScanResponseBody
	GetRequestId() *string
	SetScanMetrics(v *CreateScanResponseBodyScanMetrics) *CreateScanResponseBody
	GetScanMetrics() *CreateScanResponseBodyScanMetrics
	SetScanProgress(v int64) *CreateScanResponseBody
	GetScanProgress() *int64
	SetStartedAt(v string) *CreateScanResponseBody
	GetStartedAt() *string
	SetStatus(v string) *CreateScanResponseBody
	GetStatus() *string
	SetTaskName(v string) *CreateScanResponseBody
	GetTaskName() *string
	SetUpdatedAt(v string) *CreateScanResponseBody
	GetUpdatedAt() *string
	SetWorkerId(v string) *CreateScanResponseBody
	GetWorkerId() *string
}

type CreateScanResponseBody struct {
	// The code package ID.
	//
	// example:
	//
	// 111
	CodeBundleId *int64 `json:"codeBundleId,omitempty" xml:"codeBundleId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-08-27T00:53:46.774Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The user ID of the task creator.
	//
	// example:
	//
	// 3221
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// The supported types.
	EngineSnapshot *CreateScanResponseBodyEngineSnapshot `json:"engineSnapshot,omitempty" xml:"engineSnapshot,omitempty" type:"Struct"`
	// The scan end time.
	//
	// example:
	//
	// 2026-08-27T00:53:46.774Z
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 111
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The type. Valid values:
	//
	// 	- full: full data
	//
	// 	- incremental: incremental
	//
	// example:
	//
	// full
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 11
	ProjectId *int64 `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9A1F403F-0A85-5578-8B7C-55E3E9408659
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The scan information.
	ScanMetrics *CreateScanResponseBodyScanMetrics `json:"scanMetrics,omitempty" xml:"scanMetrics,omitempty" type:"Struct"`
	// The task progress.
	//
	// example:
	//
	// 40
	ScanProgress *int64 `json:"scanProgress,omitempty" xml:"scanProgress,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2026-08-27T00:53:46.774Z
	StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	// The status. Valid values:
	//
	// 	- running: Running.
	//
	// 	- completed: Completed.
	//
	// 	- failed: Failed.
	//
	// 	- canceling: Being canceled.
	//
	// 	- canceled: Canceled.
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The task name.
	//
	// example:
	//
	// name
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2026-08-27T00:53:46.774Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// Deprecated.
	//
	// example:
	//
	// 1
	WorkerId *string `json:"workerId,omitempty" xml:"workerId,omitempty"`
}

func (s CreateScanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScanResponseBody) GetCodeBundleId() *int64 {
	return s.CodeBundleId
}

func (s *CreateScanResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateScanResponseBody) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *CreateScanResponseBody) GetEngineSnapshot() *CreateScanResponseBodyEngineSnapshot {
	return s.EngineSnapshot
}

func (s *CreateScanResponseBody) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *CreateScanResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateScanResponseBody) GetKind() *string {
	return s.Kind
}

func (s *CreateScanResponseBody) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateScanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScanResponseBody) GetScanMetrics() *CreateScanResponseBodyScanMetrics {
	return s.ScanMetrics
}

func (s *CreateScanResponseBody) GetScanProgress() *int64 {
	return s.ScanProgress
}

func (s *CreateScanResponseBody) GetStartedAt() *string {
	return s.StartedAt
}

func (s *CreateScanResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateScanResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateScanResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateScanResponseBody) GetWorkerId() *string {
	return s.WorkerId
}

func (s *CreateScanResponseBody) SetCodeBundleId(v int64) *CreateScanResponseBody {
	s.CodeBundleId = &v
	return s
}

func (s *CreateScanResponseBody) SetCreatedAt(v string) *CreateScanResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateScanResponseBody) SetCreatedBy(v string) *CreateScanResponseBody {
	s.CreatedBy = &v
	return s
}

func (s *CreateScanResponseBody) SetEngineSnapshot(v *CreateScanResponseBodyEngineSnapshot) *CreateScanResponseBody {
	s.EngineSnapshot = v
	return s
}

func (s *CreateScanResponseBody) SetFinishedAt(v string) *CreateScanResponseBody {
	s.FinishedAt = &v
	return s
}

func (s *CreateScanResponseBody) SetId(v int64) *CreateScanResponseBody {
	s.Id = &v
	return s
}

func (s *CreateScanResponseBody) SetKind(v string) *CreateScanResponseBody {
	s.Kind = &v
	return s
}

func (s *CreateScanResponseBody) SetProjectId(v int64) *CreateScanResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateScanResponseBody) SetRequestId(v string) *CreateScanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScanResponseBody) SetScanMetrics(v *CreateScanResponseBodyScanMetrics) *CreateScanResponseBody {
	s.ScanMetrics = v
	return s
}

func (s *CreateScanResponseBody) SetScanProgress(v int64) *CreateScanResponseBody {
	s.ScanProgress = &v
	return s
}

func (s *CreateScanResponseBody) SetStartedAt(v string) *CreateScanResponseBody {
	s.StartedAt = &v
	return s
}

func (s *CreateScanResponseBody) SetStatus(v string) *CreateScanResponseBody {
	s.Status = &v
	return s
}

func (s *CreateScanResponseBody) SetTaskName(v string) *CreateScanResponseBody {
	s.TaskName = &v
	return s
}

func (s *CreateScanResponseBody) SetUpdatedAt(v string) *CreateScanResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *CreateScanResponseBody) SetWorkerId(v string) *CreateScanResponseBody {
	s.WorkerId = &v
	return s
}

func (s *CreateScanResponseBody) Validate() error {
	if s.EngineSnapshot != nil {
		if err := s.EngineSnapshot.Validate(); err != nil {
			return err
		}
	}
	if s.ScanMetrics != nil {
		if err := s.ScanMetrics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateScanResponseBodyEngineSnapshot struct {
	// Indicates whether SAST is supported.
	//
	// example:
	//
	// true
	Sast *bool `json:"sast,omitempty" xml:"sast,omitempty"`
	// Indicates whether SCA is supported.
	//
	// example:
	//
	// true
	Sca *bool `json:"sca,omitempty" xml:"sca,omitempty"`
}

func (s CreateScanResponseBodyEngineSnapshot) String() string {
	return dara.Prettify(s)
}

func (s CreateScanResponseBodyEngineSnapshot) GoString() string {
	return s.String()
}

func (s *CreateScanResponseBodyEngineSnapshot) GetSast() *bool {
	return s.Sast
}

func (s *CreateScanResponseBodyEngineSnapshot) GetSca() *bool {
	return s.Sca
}

func (s *CreateScanResponseBodyEngineSnapshot) SetSast(v bool) *CreateScanResponseBodyEngineSnapshot {
	s.Sast = &v
	return s
}

func (s *CreateScanResponseBodyEngineSnapshot) SetSca(v bool) *CreateScanResponseBodyEngineSnapshot {
	s.Sca = &v
	return s
}

func (s *CreateScanResponseBodyEngineSnapshot) Validate() error {
	return dara.Validate(s)
}

type CreateScanResponseBodyScanMetrics struct {
	// The number of files.
	//
	// example:
	//
	// 1
	FileCount *int64 `json:"fileCount,omitempty" xml:"fileCount,omitempty"`
	// The number of lines of code.
	//
	// example:
	//
	// 11
	LinesOfCode *int64 `json:"linesOfCode,omitempty" xml:"linesOfCode,omitempty"`
	// Deprecated.
	//
	// example:
	//
	// 11
	TokenTotal *int64 `json:"tokenTotal,omitempty" xml:"tokenTotal,omitempty"`
}

func (s CreateScanResponseBodyScanMetrics) String() string {
	return dara.Prettify(s)
}

func (s CreateScanResponseBodyScanMetrics) GoString() string {
	return s.String()
}

func (s *CreateScanResponseBodyScanMetrics) GetFileCount() *int64 {
	return s.FileCount
}

func (s *CreateScanResponseBodyScanMetrics) GetLinesOfCode() *int64 {
	return s.LinesOfCode
}

func (s *CreateScanResponseBodyScanMetrics) GetTokenTotal() *int64 {
	return s.TokenTotal
}

func (s *CreateScanResponseBodyScanMetrics) SetFileCount(v int64) *CreateScanResponseBodyScanMetrics {
	s.FileCount = &v
	return s
}

func (s *CreateScanResponseBodyScanMetrics) SetLinesOfCode(v int64) *CreateScanResponseBodyScanMetrics {
	s.LinesOfCode = &v
	return s
}

func (s *CreateScanResponseBodyScanMetrics) SetTokenTotal(v int64) *CreateScanResponseBodyScanMetrics {
	s.TokenTotal = &v
	return s
}

func (s *CreateScanResponseBodyScanMetrics) Validate() error {
	return dara.Validate(s)
}
