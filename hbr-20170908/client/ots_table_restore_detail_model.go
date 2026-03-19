// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOtsTableRestoreDetail interface {
	dara.Model
	String() string
	GoString() string
	SetBatchChannelCount(v int32) *OtsTableRestoreDetail
	GetBatchChannelCount() *int32
	SetIndexNameSuffix(v string) *OtsTableRestoreDetail
	GetIndexNameSuffix() *string
	SetOverwriteExisting(v bool) *OtsTableRestoreDetail
	GetOverwriteExisting() *bool
	SetReGenerateAutoIncrementPK(v bool) *OtsTableRestoreDetail
	GetReGenerateAutoIncrementPK() *bool
	SetRestoreIndex(v bool) *OtsTableRestoreDetail
	GetRestoreIndex() *bool
	SetRestoreSearchIndex(v bool) *OtsTableRestoreDetail
	GetRestoreSearchIndex() *bool
	SetSearchIndexNameSuffix(v string) *OtsTableRestoreDetail
	GetSearchIndexNameSuffix() *string
}

type OtsTableRestoreDetail struct {
	// The number of concurrent processes for each restore job.
	//
	// example:
	//
	// 0
	BatchChannelCount *int32 `json:"BatchChannelCount,omitempty" xml:"BatchChannelCount,omitempty"`
	// The name prefixes of the indexes that you want to restore.
	//
	// example:
	//
	// 2022xxxx143933
	IndexNameSuffix *string `json:"IndexNameSuffix,omitempty" xml:"IndexNameSuffix,omitempty"`
	// Specifies whether to overwrite existing tables.
	//
	// example:
	//
	// true
	OverwriteExisting *bool `json:"OverwriteExisting,omitempty" xml:"OverwriteExisting,omitempty"`
	// Specifies whether to regenerate auto-increment primary keys.
	//
	// example:
	//
	// true
	ReGenerateAutoIncrementPK *bool `json:"ReGenerateAutoIncrementPK,omitempty" xml:"ReGenerateAutoIncrementPK,omitempty"`
	// Specifies whether to restore indexes.
	//
	// example:
	//
	// true
	RestoreIndex *bool `json:"RestoreIndex,omitempty" xml:"RestoreIndex,omitempty"`
	// Specifies whether to restore search indexes.
	//
	// example:
	//
	// true
	RestoreSearchIndex *bool `json:"RestoreSearchIndex,omitempty" xml:"RestoreSearchIndex,omitempty"`
	// The name prefixes of the search indexes that you want to restore.
	//
	// example:
	//
	// 2022xxxx143933
	SearchIndexNameSuffix *string `json:"SearchIndexNameSuffix,omitempty" xml:"SearchIndexNameSuffix,omitempty"`
}

func (s OtsTableRestoreDetail) String() string {
	return dara.Prettify(s)
}

func (s OtsTableRestoreDetail) GoString() string {
	return s.String()
}

func (s *OtsTableRestoreDetail) GetBatchChannelCount() *int32 {
	return s.BatchChannelCount
}

func (s *OtsTableRestoreDetail) GetIndexNameSuffix() *string {
	return s.IndexNameSuffix
}

func (s *OtsTableRestoreDetail) GetOverwriteExisting() *bool {
	return s.OverwriteExisting
}

func (s *OtsTableRestoreDetail) GetReGenerateAutoIncrementPK() *bool {
	return s.ReGenerateAutoIncrementPK
}

func (s *OtsTableRestoreDetail) GetRestoreIndex() *bool {
	return s.RestoreIndex
}

func (s *OtsTableRestoreDetail) GetRestoreSearchIndex() *bool {
	return s.RestoreSearchIndex
}

func (s *OtsTableRestoreDetail) GetSearchIndexNameSuffix() *string {
	return s.SearchIndexNameSuffix
}

func (s *OtsTableRestoreDetail) SetBatchChannelCount(v int32) *OtsTableRestoreDetail {
	s.BatchChannelCount = &v
	return s
}

func (s *OtsTableRestoreDetail) SetIndexNameSuffix(v string) *OtsTableRestoreDetail {
	s.IndexNameSuffix = &v
	return s
}

func (s *OtsTableRestoreDetail) SetOverwriteExisting(v bool) *OtsTableRestoreDetail {
	s.OverwriteExisting = &v
	return s
}

func (s *OtsTableRestoreDetail) SetReGenerateAutoIncrementPK(v bool) *OtsTableRestoreDetail {
	s.ReGenerateAutoIncrementPK = &v
	return s
}

func (s *OtsTableRestoreDetail) SetRestoreIndex(v bool) *OtsTableRestoreDetail {
	s.RestoreIndex = &v
	return s
}

func (s *OtsTableRestoreDetail) SetRestoreSearchIndex(v bool) *OtsTableRestoreDetail {
	s.RestoreSearchIndex = &v
	return s
}

func (s *OtsTableRestoreDetail) SetSearchIndexNameSuffix(v string) *OtsTableRestoreDetail {
	s.SearchIndexNameSuffix = &v
	return s
}

func (s *OtsTableRestoreDetail) Validate() error {
	return dara.Validate(s)
}
