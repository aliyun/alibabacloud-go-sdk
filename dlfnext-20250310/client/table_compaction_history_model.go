// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTableCompactionHistory interface {
	dara.Model
	String() string
	GoString() string
	SetAfterFilesCount(v int64) *TableCompactionHistory
	GetAfterFilesCount() *int64
	SetAfterFilesSize(v int64) *TableCompactionHistory
	GetAfterFilesSize() *int64
	SetBeforeFilesCount(v int64) *TableCompactionHistory
	GetBeforeFilesCount() *int64
	SetBeforeFilesLastCreationTime(v int64) *TableCompactionHistory
	GetBeforeFilesLastCreationTime() *int64
	SetBeforeFilesSize(v int64) *TableCompactionHistory
	GetBeforeFilesSize() *int64
	SetCatalogId(v string) *TableCompactionHistory
	GetCatalogId() *string
	SetCommitTime(v int64) *TableCompactionHistory
	GetCommitTime() *int64
	SetSnapshotId(v int64) *TableCompactionHistory
	GetSnapshotId() *int64
	SetTableId(v string) *TableCompactionHistory
	GetTableId() *string
	SetUpdatedAt(v int64) *TableCompactionHistory
	GetUpdatedAt() *int64
	SetUpdatedBy(v string) *TableCompactionHistory
	GetUpdatedBy() *string
}

type TableCompactionHistory struct {
	AfterFilesCount             *int64  `json:"afterFilesCount,omitempty" xml:"afterFilesCount,omitempty"`
	AfterFilesSize              *int64  `json:"afterFilesSize,omitempty" xml:"afterFilesSize,omitempty"`
	BeforeFilesCount            *int64  `json:"beforeFilesCount,omitempty" xml:"beforeFilesCount,omitempty"`
	BeforeFilesLastCreationTime *int64  `json:"beforeFilesLastCreationTime,omitempty" xml:"beforeFilesLastCreationTime,omitempty"`
	BeforeFilesSize             *int64  `json:"beforeFilesSize,omitempty" xml:"beforeFilesSize,omitempty"`
	CatalogId                   *string `json:"catalogId,omitempty" xml:"catalogId,omitempty"`
	CommitTime                  *int64  `json:"commitTime,omitempty" xml:"commitTime,omitempty"`
	SnapshotId                  *int64  `json:"snapshotId,omitempty" xml:"snapshotId,omitempty"`
	TableId                     *string `json:"tableId,omitempty" xml:"tableId,omitempty"`
	UpdatedAt                   *int64  `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UpdatedBy                   *string `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
}

func (s TableCompactionHistory) String() string {
	return dara.Prettify(s)
}

func (s TableCompactionHistory) GoString() string {
	return s.String()
}

func (s *TableCompactionHistory) GetAfterFilesCount() *int64 {
	return s.AfterFilesCount
}

func (s *TableCompactionHistory) GetAfterFilesSize() *int64 {
	return s.AfterFilesSize
}

func (s *TableCompactionHistory) GetBeforeFilesCount() *int64 {
	return s.BeforeFilesCount
}

func (s *TableCompactionHistory) GetBeforeFilesLastCreationTime() *int64 {
	return s.BeforeFilesLastCreationTime
}

func (s *TableCompactionHistory) GetBeforeFilesSize() *int64 {
	return s.BeforeFilesSize
}

func (s *TableCompactionHistory) GetCatalogId() *string {
	return s.CatalogId
}

func (s *TableCompactionHistory) GetCommitTime() *int64 {
	return s.CommitTime
}

func (s *TableCompactionHistory) GetSnapshotId() *int64 {
	return s.SnapshotId
}

func (s *TableCompactionHistory) GetTableId() *string {
	return s.TableId
}

func (s *TableCompactionHistory) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *TableCompactionHistory) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *TableCompactionHistory) SetAfterFilesCount(v int64) *TableCompactionHistory {
	s.AfterFilesCount = &v
	return s
}

func (s *TableCompactionHistory) SetAfterFilesSize(v int64) *TableCompactionHistory {
	s.AfterFilesSize = &v
	return s
}

func (s *TableCompactionHistory) SetBeforeFilesCount(v int64) *TableCompactionHistory {
	s.BeforeFilesCount = &v
	return s
}

func (s *TableCompactionHistory) SetBeforeFilesLastCreationTime(v int64) *TableCompactionHistory {
	s.BeforeFilesLastCreationTime = &v
	return s
}

func (s *TableCompactionHistory) SetBeforeFilesSize(v int64) *TableCompactionHistory {
	s.BeforeFilesSize = &v
	return s
}

func (s *TableCompactionHistory) SetCatalogId(v string) *TableCompactionHistory {
	s.CatalogId = &v
	return s
}

func (s *TableCompactionHistory) SetCommitTime(v int64) *TableCompactionHistory {
	s.CommitTime = &v
	return s
}

func (s *TableCompactionHistory) SetSnapshotId(v int64) *TableCompactionHistory {
	s.SnapshotId = &v
	return s
}

func (s *TableCompactionHistory) SetTableId(v string) *TableCompactionHistory {
	s.TableId = &v
	return s
}

func (s *TableCompactionHistory) SetUpdatedAt(v int64) *TableCompactionHistory {
	s.UpdatedAt = &v
	return s
}

func (s *TableCompactionHistory) SetUpdatedBy(v string) *TableCompactionHistory {
	s.UpdatedBy = &v
	return s
}

func (s *TableCompactionHistory) Validate() error {
	return dara.Validate(s)
}
