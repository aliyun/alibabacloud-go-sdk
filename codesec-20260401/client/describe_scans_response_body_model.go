// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeScansResponseBodyItems) *DescribeScansResponseBody
	GetItems() []*DescribeScansResponseBodyItems
	SetMaxResults(v int64) *DescribeScansResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeScansResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeScansResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeScansResponseBody
	GetTotalCount() *int64
}

type DescribeScansResponseBody struct {
	// The task list.
	Items []*DescribeScansResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The page size.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. An empty value indicates the last page.
	//
	// example:
	//
	// eyJ0IjoiMjAyNi0wNy0xNlQwNzo1MzozOC4wMjFaIiwiaSI6MTAwMDQ0OH0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9A1F403F-0A85-5578-8B7C-55E3E9408659
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s DescribeScansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScansResponseBody) GetItems() []*DescribeScansResponseBodyItems {
	return s.Items
}

func (s *DescribeScansResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeScansResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeScansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScansResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeScansResponseBody) SetItems(v []*DescribeScansResponseBodyItems) *DescribeScansResponseBody {
	s.Items = v
	return s
}

func (s *DescribeScansResponseBody) SetMaxResults(v int64) *DescribeScansResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeScansResponseBody) SetNextToken(v string) *DescribeScansResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeScansResponseBody) SetRequestId(v string) *DescribeScansResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScansResponseBody) SetTotalCount(v int64) *DescribeScansResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeScansResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeScansResponseBodyItems struct {
	// The code bundle ID.
	//
	// example:
	//
	// 11
	CodeBundleId *int64 `json:"codeBundleId,omitempty" xml:"codeBundleId,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2026-07-28T03:36:31.573Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The user ID of the task creator.
	//
	// example:
	//
	// 11111
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// The scan phase. Valid values:
	//
	// 	- threat_model: Threat modeling.
	//
	// 	- discovery: Vulnerability discovery.
	//
	// 	- panel: Vulnerability review.
	//
	// 	- adversarial: Adversarial verification.
	//
	// 	- finalize: Report generation.
	//
	// example:
	//
	// finalize
	CurrentPhase *string `json:"currentPhase,omitempty" xml:"currentPhase,omitempty"`
	// The supported scan types.
	EngineSnapshot *DescribeScansResponseBodyItemsEngineSnapshot `json:"engineSnapshot,omitempty" xml:"engineSnapshot,omitempty" type:"Struct"`
	// The time when the scan finished.
	//
	// example:
	//
	// 2026-07-28T03:36:31.573Z
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 934
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The scan type. Valid values:
	//
	// 	- full: Full scan.
	//
	// 	- incremental: Incremental scan.
	//
	// example:
	//
	// full
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The project ID.
	//
	// example:
	//
	// p-n72k9yrkq81ny7z
	ProjectId *int64 `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The scan result statistics information.
	ScanMetrics *DescribeScansResponseBodyItemsScanMetrics `json:"scanMetrics,omitempty" xml:"scanMetrics,omitempty" type:"Struct"`
	// The task progress.
	//
	// example:
	//
	// 100
	ScanProgress *int64 `json:"scanProgress,omitempty" xml:"scanProgress,omitempty"`
	// The time when the task started.
	//
	// example:
	//
	// 2026-07-28T03:36:31.573Z
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
	// 1648622222394847-ha-cn-lm64p7tby01_dsl_kb_video_1773817008236_full
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// The time when the task was last updated.
	//
	// example:
	//
	// 2026-07-28T03:36:31.573Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// Deprecated.
	//
	// example:
	//
	// 1
	WorkerId *string `json:"workerId,omitempty" xml:"workerId,omitempty"`
}

func (s DescribeScansResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeScansResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeScansResponseBodyItems) GetCodeBundleId() *int64 {
	return s.CodeBundleId
}

func (s *DescribeScansResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DescribeScansResponseBodyItems) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *DescribeScansResponseBodyItems) GetCurrentPhase() *string {
	return s.CurrentPhase
}

func (s *DescribeScansResponseBodyItems) GetEngineSnapshot() *DescribeScansResponseBodyItemsEngineSnapshot {
	return s.EngineSnapshot
}

func (s *DescribeScansResponseBodyItems) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *DescribeScansResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *DescribeScansResponseBodyItems) GetKind() *string {
	return s.Kind
}

func (s *DescribeScansResponseBodyItems) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DescribeScansResponseBodyItems) GetScanMetrics() *DescribeScansResponseBodyItemsScanMetrics {
	return s.ScanMetrics
}

func (s *DescribeScansResponseBodyItems) GetScanProgress() *int64 {
	return s.ScanProgress
}

func (s *DescribeScansResponseBodyItems) GetStartedAt() *string {
	return s.StartedAt
}

func (s *DescribeScansResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeScansResponseBodyItems) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeScansResponseBodyItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *DescribeScansResponseBodyItems) GetWorkerId() *string {
	return s.WorkerId
}

func (s *DescribeScansResponseBodyItems) SetCodeBundleId(v int64) *DescribeScansResponseBodyItems {
	s.CodeBundleId = &v
	return s
}

func (s *DescribeScansResponseBodyItems) SetCreatedAt(v string) *DescribeScansResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *DescribeScansResponseBodyItems) SetCreatedBy(v string) *DescribeScansResponseBodyItems {
	s.CreatedBy = &v
	return s
}

func (s *DescribeScansResponseBodyItems) SetCurrentPhase(v string) *DescribeScansResponseBodyItems {
	s.CurrentPhase = &v
	return s
}

func (s *DescribeScansResponseBodyItems) SetEngineSnapshot(v *DescribeScansResponseBodyItemsEngineSnapshot) *DescribeScansResponseBodyItems {
	s.EngineSnapshot = v
	return s
}

func (s *DescribeScansResponseBodyItems) SetFinishedAt(v string) *DescribeScansResponseBodyItems {
	s.FinishedAt = &v
	return s
}

func (s *DescribeScansResponseBodyItems) SetId(v int64) *DescribeScansResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeScansResponseBodyItems) SetKind(v string) *DescribeScansResponseBodyItems {
	s.Kind = &v
	return s
}

func (s *DescribeScansResponseBodyItems) SetProjectId(v int64) *DescribeScansResponseBodyItems {
	s.ProjectId = &v
	return s
}

func (s *DescribeScansResponseBodyItems) SetScanMetrics(v *DescribeScansResponseBodyItemsScanMetrics) *DescribeScansResponseBodyItems {
	s.ScanMetrics = v
	return s
}

func (s *DescribeScansResponseBodyItems) SetScanProgress(v int64) *DescribeScansResponseBodyItems {
	s.ScanProgress = &v
	return s
}

func (s *DescribeScansResponseBodyItems) SetStartedAt(v string) *DescribeScansResponseBodyItems {
	s.StartedAt = &v
	return s
}

func (s *DescribeScansResponseBodyItems) SetStatus(v string) *DescribeScansResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeScansResponseBodyItems) SetTaskName(v string) *DescribeScansResponseBodyItems {
	s.TaskName = &v
	return s
}

func (s *DescribeScansResponseBodyItems) SetUpdatedAt(v string) *DescribeScansResponseBodyItems {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeScansResponseBodyItems) SetWorkerId(v string) *DescribeScansResponseBodyItems {
	s.WorkerId = &v
	return s
}

func (s *DescribeScansResponseBodyItems) Validate() error {
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

type DescribeScansResponseBodyItemsEngineSnapshot struct {
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

func (s DescribeScansResponseBodyItemsEngineSnapshot) String() string {
	return dara.Prettify(s)
}

func (s DescribeScansResponseBodyItemsEngineSnapshot) GoString() string {
	return s.String()
}

func (s *DescribeScansResponseBodyItemsEngineSnapshot) GetSast() *bool {
	return s.Sast
}

func (s *DescribeScansResponseBodyItemsEngineSnapshot) GetSca() *bool {
	return s.Sca
}

func (s *DescribeScansResponseBodyItemsEngineSnapshot) SetSast(v bool) *DescribeScansResponseBodyItemsEngineSnapshot {
	s.Sast = &v
	return s
}

func (s *DescribeScansResponseBodyItemsEngineSnapshot) SetSca(v bool) *DescribeScansResponseBodyItemsEngineSnapshot {
	s.Sca = &v
	return s
}

func (s *DescribeScansResponseBodyItemsEngineSnapshot) Validate() error {
	return dara.Validate(s)
}

type DescribeScansResponseBodyItemsScanMetrics struct {
	// The number of credits consumed by the task.
	//
	// example:
	//
	// 1.25
	Credit *float32 `json:"credit,omitempty" xml:"credit,omitempty"`
	// The number of files.
	//
	// example:
	//
	// 459
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
	// 1
	TokenTotal *int64 `json:"tokenTotal,omitempty" xml:"tokenTotal,omitempty"`
}

func (s DescribeScansResponseBodyItemsScanMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeScansResponseBodyItemsScanMetrics) GoString() string {
	return s.String()
}

func (s *DescribeScansResponseBodyItemsScanMetrics) GetCredit() *float32 {
	return s.Credit
}

func (s *DescribeScansResponseBodyItemsScanMetrics) GetFileCount() *int64 {
	return s.FileCount
}

func (s *DescribeScansResponseBodyItemsScanMetrics) GetLinesOfCode() *int64 {
	return s.LinesOfCode
}

func (s *DescribeScansResponseBodyItemsScanMetrics) GetTokenTotal() *int64 {
	return s.TokenTotal
}

func (s *DescribeScansResponseBodyItemsScanMetrics) SetCredit(v float32) *DescribeScansResponseBodyItemsScanMetrics {
	s.Credit = &v
	return s
}

func (s *DescribeScansResponseBodyItemsScanMetrics) SetFileCount(v int64) *DescribeScansResponseBodyItemsScanMetrics {
	s.FileCount = &v
	return s
}

func (s *DescribeScansResponseBodyItemsScanMetrics) SetLinesOfCode(v int64) *DescribeScansResponseBodyItemsScanMetrics {
	s.LinesOfCode = &v
	return s
}

func (s *DescribeScansResponseBodyItemsScanMetrics) SetTokenTotal(v int64) *DescribeScansResponseBodyItemsScanMetrics {
	s.TokenTotal = &v
	return s
}

func (s *DescribeScansResponseBodyItemsScanMetrics) Validate() error {
	return dara.Validate(s)
}
