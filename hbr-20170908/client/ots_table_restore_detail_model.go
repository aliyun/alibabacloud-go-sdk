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
	BatchChannelCount         *int32  `json:"BatchChannelCount,omitempty" xml:"BatchChannelCount,omitempty"`
	IndexNameSuffix           *string `json:"IndexNameSuffix,omitempty" xml:"IndexNameSuffix,omitempty"`
	OverwriteExisting         *bool   `json:"OverwriteExisting,omitempty" xml:"OverwriteExisting,omitempty"`
	ReGenerateAutoIncrementPK *bool   `json:"ReGenerateAutoIncrementPK,omitempty" xml:"ReGenerateAutoIncrementPK,omitempty"`
	RestoreIndex              *bool   `json:"RestoreIndex,omitempty" xml:"RestoreIndex,omitempty"`
	RestoreSearchIndex        *bool   `json:"RestoreSearchIndex,omitempty" xml:"RestoreSearchIndex,omitempty"`
	SearchIndexNameSuffix     *string `json:"SearchIndexNameSuffix,omitempty" xml:"SearchIndexNameSuffix,omitempty"`
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
