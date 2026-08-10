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
	Items      []*DescribeScansResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	MaxResults *int64                            `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string                           `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId  *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *int64                            `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	CodeBundleId *int64 `json:"codeBundleId,omitempty" xml:"codeBundleId,omitempty"`
	// 扫描任务创建时间（RFC3339）
	CreatedAt      *string                                       `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy      *string                                       `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	EngineSnapshot *DescribeScansResponseBodyItemsEngineSnapshot `json:"engineSnapshot,omitempty" xml:"engineSnapshot,omitempty" type:"Struct"`
	// 扫描结束时间（RFC3339）
	FinishedAt   *string                                    `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	Id           *int64                                     `json:"id,omitempty" xml:"id,omitempty"`
	Kind         *string                                    `json:"kind,omitempty" xml:"kind,omitempty"`
	ProjectId    *int64                                     `json:"projectId,omitempty" xml:"projectId,omitempty"`
	ScanMetrics  *DescribeScansResponseBodyItemsScanMetrics `json:"scanMetrics,omitempty" xml:"scanMetrics,omitempty" type:"Struct"`
	ScanProgress *int64                                     `json:"scanProgress,omitempty" xml:"scanProgress,omitempty"`
	// 扫描开始时间（RFC3339）
	StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
	TaskName  *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// 扫描任务更新时间（RFC3339）
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	WorkerId  *string `json:"workerId,omitempty" xml:"workerId,omitempty"`
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
	Sast *bool `json:"sast,omitempty" xml:"sast,omitempty"`
	Sca  *bool `json:"sca,omitempty" xml:"sca,omitempty"`
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
	Credit      *float32 `json:"credit,omitempty" xml:"credit,omitempty"`
	FileCount   *int64   `json:"fileCount,omitempty" xml:"fileCount,omitempty"`
	LinesOfCode *int64   `json:"linesOfCode,omitempty" xml:"linesOfCode,omitempty"`
	TokenTotal  *int64   `json:"tokenTotal,omitempty" xml:"tokenTotal,omitempty"`
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
