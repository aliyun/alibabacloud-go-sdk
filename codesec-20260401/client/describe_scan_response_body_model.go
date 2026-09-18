// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCodeBundleId(v int64) *DescribeScanResponseBody
	GetCodeBundleId() *int64
	SetCreatedAt(v string) *DescribeScanResponseBody
	GetCreatedAt() *string
	SetCreatedBy(v string) *DescribeScanResponseBody
	GetCreatedBy() *string
	SetCurrentPhase(v string) *DescribeScanResponseBody
	GetCurrentPhase() *string
	SetEngineSnapshot(v *DescribeScanResponseBodyEngineSnapshot) *DescribeScanResponseBody
	GetEngineSnapshot() *DescribeScanResponseBodyEngineSnapshot
	SetFinishedAt(v string) *DescribeScanResponseBody
	GetFinishedAt() *string
	SetId(v int64) *DescribeScanResponseBody
	GetId() *int64
	SetKind(v string) *DescribeScanResponseBody
	GetKind() *string
	SetProjectId(v int64) *DescribeScanResponseBody
	GetProjectId() *int64
	SetRequestId(v string) *DescribeScanResponseBody
	GetRequestId() *string
	SetScanMetrics(v *DescribeScanResponseBodyScanMetrics) *DescribeScanResponseBody
	GetScanMetrics() *DescribeScanResponseBodyScanMetrics
	SetScanProgress(v int64) *DescribeScanResponseBody
	GetScanProgress() *int64
	SetSecurityCredits(v float32) *DescribeScanResponseBody
	GetSecurityCredits() *float32
	SetStartedAt(v string) *DescribeScanResponseBody
	GetStartedAt() *string
	SetStatus(v string) *DescribeScanResponseBody
	GetStatus() *string
	SetTaskName(v string) *DescribeScanResponseBody
	GetTaskName() *string
	SetUpdatedAt(v string) *DescribeScanResponseBody
	GetUpdatedAt() *string
	SetWorkerId(v string) *DescribeScanResponseBody
	GetWorkerId() *string
}

type DescribeScanResponseBody struct {
	// The function code package ID.
	//
	// example:
	//
	// 111
	CodeBundleId *int64 `json:"codeBundleId,omitempty" xml:"codeBundleId,omitempty"`
	// The time when the task was created.
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
	// The scan phase. Valid values:
	//
	// 	- threat_model: threat modeling.
	//
	// 	- discovery: vulnerability discovery.
	//
	// 	- panel: vulnerability review.
	//
	// 	- adversarial: adversarial verification.
	//
	// 	- finalize: report compilation.
	//
	// example:
	//
	// discovery
	CurrentPhase *string `json:"currentPhase,omitempty" xml:"currentPhase,omitempty"`
	// The supported engine types.
	EngineSnapshot *DescribeScanResponseBodyEngineSnapshot `json:"engineSnapshot,omitempty" xml:"engineSnapshot,omitempty" type:"Struct"`
	// The time when the scan finished.
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
	// The scan type. Valid values:
	//
	// 	- full: full
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
	// 1111
	ProjectId *int64 `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 9A1F403F-0A85-5578-8B7C-55E3E9408659
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The scan results.
	ScanMetrics *DescribeScanResponseBodyScanMetrics `json:"scanMetrics,omitempty" xml:"scanMetrics,omitempty" type:"Struct"`
	// The task progress.
	//
	// example:
	//
	// 40
	ScanProgress *int64 `json:"scanProgress,omitempty" xml:"scanProgress,omitempty"`
	// **[Deprecated]*	- This parameter is no longer used.
	//
	// example:
	//
	// 1
	SecurityCredits *float32 `json:"securityCredits,omitempty" xml:"securityCredits,omitempty"`
	// The time when the task started.
	//
	// example:
	//
	// 2026-08-27T00:53:46.774Z
	StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	// The task status. Valid values:
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
	// completed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The task name.
	//
	// example:
	//
	// name
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// The time when the task was last updated.
	//
	// example:
	//
	// 2026-08-27T00:53:46.774Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// **[Deprecated]*	- This parameter is no longer used.
	//
	// example:
	//
	// 1
	WorkerId *string `json:"workerId,omitempty" xml:"workerId,omitempty"`
}

func (s DescribeScanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScanResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScanResponseBody) GetCodeBundleId() *int64 {
	return s.CodeBundleId
}

func (s *DescribeScanResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DescribeScanResponseBody) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *DescribeScanResponseBody) GetCurrentPhase() *string {
	return s.CurrentPhase
}

func (s *DescribeScanResponseBody) GetEngineSnapshot() *DescribeScanResponseBodyEngineSnapshot {
	return s.EngineSnapshot
}

func (s *DescribeScanResponseBody) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *DescribeScanResponseBody) GetId() *int64 {
	return s.Id
}

func (s *DescribeScanResponseBody) GetKind() *string {
	return s.Kind
}

func (s *DescribeScanResponseBody) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DescribeScanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScanResponseBody) GetScanMetrics() *DescribeScanResponseBodyScanMetrics {
	return s.ScanMetrics
}

func (s *DescribeScanResponseBody) GetScanProgress() *int64 {
	return s.ScanProgress
}

func (s *DescribeScanResponseBody) GetSecurityCredits() *float32 {
	return s.SecurityCredits
}

func (s *DescribeScanResponseBody) GetStartedAt() *string {
	return s.StartedAt
}

func (s *DescribeScanResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeScanResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeScanResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *DescribeScanResponseBody) GetWorkerId() *string {
	return s.WorkerId
}

func (s *DescribeScanResponseBody) SetCodeBundleId(v int64) *DescribeScanResponseBody {
	s.CodeBundleId = &v
	return s
}

func (s *DescribeScanResponseBody) SetCreatedAt(v string) *DescribeScanResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *DescribeScanResponseBody) SetCreatedBy(v string) *DescribeScanResponseBody {
	s.CreatedBy = &v
	return s
}

func (s *DescribeScanResponseBody) SetCurrentPhase(v string) *DescribeScanResponseBody {
	s.CurrentPhase = &v
	return s
}

func (s *DescribeScanResponseBody) SetEngineSnapshot(v *DescribeScanResponseBodyEngineSnapshot) *DescribeScanResponseBody {
	s.EngineSnapshot = v
	return s
}

func (s *DescribeScanResponseBody) SetFinishedAt(v string) *DescribeScanResponseBody {
	s.FinishedAt = &v
	return s
}

func (s *DescribeScanResponseBody) SetId(v int64) *DescribeScanResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeScanResponseBody) SetKind(v string) *DescribeScanResponseBody {
	s.Kind = &v
	return s
}

func (s *DescribeScanResponseBody) SetProjectId(v int64) *DescribeScanResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeScanResponseBody) SetRequestId(v string) *DescribeScanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScanResponseBody) SetScanMetrics(v *DescribeScanResponseBodyScanMetrics) *DescribeScanResponseBody {
	s.ScanMetrics = v
	return s
}

func (s *DescribeScanResponseBody) SetScanProgress(v int64) *DescribeScanResponseBody {
	s.ScanProgress = &v
	return s
}

func (s *DescribeScanResponseBody) SetSecurityCredits(v float32) *DescribeScanResponseBody {
	s.SecurityCredits = &v
	return s
}

func (s *DescribeScanResponseBody) SetStartedAt(v string) *DescribeScanResponseBody {
	s.StartedAt = &v
	return s
}

func (s *DescribeScanResponseBody) SetStatus(v string) *DescribeScanResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeScanResponseBody) SetTaskName(v string) *DescribeScanResponseBody {
	s.TaskName = &v
	return s
}

func (s *DescribeScanResponseBody) SetUpdatedAt(v string) *DescribeScanResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeScanResponseBody) SetWorkerId(v string) *DescribeScanResponseBody {
	s.WorkerId = &v
	return s
}

func (s *DescribeScanResponseBody) Validate() error {
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

type DescribeScanResponseBodyEngineSnapshot struct {
	// Indicates whether Static Application Security Testing (SAST) is supported.
	//
	// example:
	//
	// true
	Sast *bool `json:"sast,omitempty" xml:"sast,omitempty"`
	// Indicates whether Software Composition Analysis (SCA) is supported.
	//
	// example:
	//
	// true
	Sca *bool `json:"sca,omitempty" xml:"sca,omitempty"`
}

func (s DescribeScanResponseBodyEngineSnapshot) String() string {
	return dara.Prettify(s)
}

func (s DescribeScanResponseBodyEngineSnapshot) GoString() string {
	return s.String()
}

func (s *DescribeScanResponseBodyEngineSnapshot) GetSast() *bool {
	return s.Sast
}

func (s *DescribeScanResponseBodyEngineSnapshot) GetSca() *bool {
	return s.Sca
}

func (s *DescribeScanResponseBodyEngineSnapshot) SetSast(v bool) *DescribeScanResponseBodyEngineSnapshot {
	s.Sast = &v
	return s
}

func (s *DescribeScanResponseBodyEngineSnapshot) SetSca(v bool) *DescribeScanResponseBodyEngineSnapshot {
	s.Sca = &v
	return s
}

func (s *DescribeScanResponseBodyEngineSnapshot) Validate() error {
	return dara.Validate(s)
}

type DescribeScanResponseBodyScanMetrics struct {
	// The number of credits consumed by the task.
	//
	// example:
	//
	// 1
	Credit *float32 `json:"credit,omitempty" xml:"credit,omitempty"`
	// The number of files.
	//
	// example:
	//
	// 73894
	FileCount *int64 `json:"fileCount,omitempty" xml:"fileCount,omitempty"`
	// The number of lines of code.
	//
	// example:
	//
	// 1
	LinesOfCode *int64 `json:"linesOfCode,omitempty" xml:"linesOfCode,omitempty"`
	// **[Deprecated]*	- This parameter is no longer used.
	//
	// example:
	//
	// 1
	TokenTotal *int64 `json:"tokenTotal,omitempty" xml:"tokenTotal,omitempty"`
}

func (s DescribeScanResponseBodyScanMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeScanResponseBodyScanMetrics) GoString() string {
	return s.String()
}

func (s *DescribeScanResponseBodyScanMetrics) GetCredit() *float32 {
	return s.Credit
}

func (s *DescribeScanResponseBodyScanMetrics) GetFileCount() *int64 {
	return s.FileCount
}

func (s *DescribeScanResponseBodyScanMetrics) GetLinesOfCode() *int64 {
	return s.LinesOfCode
}

func (s *DescribeScanResponseBodyScanMetrics) GetTokenTotal() *int64 {
	return s.TokenTotal
}

func (s *DescribeScanResponseBodyScanMetrics) SetCredit(v float32) *DescribeScanResponseBodyScanMetrics {
	s.Credit = &v
	return s
}

func (s *DescribeScanResponseBodyScanMetrics) SetFileCount(v int64) *DescribeScanResponseBodyScanMetrics {
	s.FileCount = &v
	return s
}

func (s *DescribeScanResponseBodyScanMetrics) SetLinesOfCode(v int64) *DescribeScanResponseBodyScanMetrics {
	s.LinesOfCode = &v
	return s
}

func (s *DescribeScanResponseBodyScanMetrics) SetTokenTotal(v int64) *DescribeScanResponseBodyScanMetrics {
	s.TokenTotal = &v
	return s
}

func (s *DescribeScanResponseBodyScanMetrics) Validate() error {
	return dara.Validate(s)
}
