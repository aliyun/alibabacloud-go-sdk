// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenStructDlfSnapshotDto interface {
	dara.Model
	String() string
	GoString() string
	SetChangelogRecordCount(v int64) *OpenStructDlfSnapshotDto
	GetChangelogRecordCount() *int64
	SetCommitKind(v string) *OpenStructDlfSnapshotDto
	GetCommitKind() *string
	SetDeltaRecordCount(v int64) *OpenStructDlfSnapshotDto
	GetDeltaRecordCount() *int64
	SetSchemaId(v int64) *OpenStructDlfSnapshotDto
	GetSchemaId() *int64
	SetSnapshotId(v int64) *OpenStructDlfSnapshotDto
	GetSnapshotId() *int64
	SetTimeMillis(v int64) *OpenStructDlfSnapshotDto
	GetTimeMillis() *int64
	SetTotalRecordCount(v int64) *OpenStructDlfSnapshotDto
	GetTotalRecordCount() *int64
}

type OpenStructDlfSnapshotDto struct {
	// example:
	//
	// 500
	ChangelogRecordCount *int64 `json:"changelogRecordCount,omitempty" xml:"changelogRecordCount,omitempty"`
	// example:
	//
	// APPEND
	CommitKind *string `json:"commitKind,omitempty" xml:"commitKind,omitempty"`
	// example:
	//
	// 1000
	DeltaRecordCount *int64 `json:"deltaRecordCount,omitempty" xml:"deltaRecordCount,omitempty"`
	// example:
	//
	// 1
	SchemaId *int64 `json:"schemaId,omitempty" xml:"schemaId,omitempty"`
	// example:
	//
	// 123456789
	SnapshotId *int64 `json:"snapshotId,omitempty" xml:"snapshotId,omitempty"`
	// example:
	//
	// 1711334400000
	TimeMillis *int64 `json:"timeMillis,omitempty" xml:"timeMillis,omitempty"`
	// example:
	//
	// 100000
	TotalRecordCount *int64 `json:"totalRecordCount,omitempty" xml:"totalRecordCount,omitempty"`
}

func (s OpenStructDlfSnapshotDto) String() string {
	return dara.Prettify(s)
}

func (s OpenStructDlfSnapshotDto) GoString() string {
	return s.String()
}

func (s *OpenStructDlfSnapshotDto) GetChangelogRecordCount() *int64 {
	return s.ChangelogRecordCount
}

func (s *OpenStructDlfSnapshotDto) GetCommitKind() *string {
	return s.CommitKind
}

func (s *OpenStructDlfSnapshotDto) GetDeltaRecordCount() *int64 {
	return s.DeltaRecordCount
}

func (s *OpenStructDlfSnapshotDto) GetSchemaId() *int64 {
	return s.SchemaId
}

func (s *OpenStructDlfSnapshotDto) GetSnapshotId() *int64 {
	return s.SnapshotId
}

func (s *OpenStructDlfSnapshotDto) GetTimeMillis() *int64 {
	return s.TimeMillis
}

func (s *OpenStructDlfSnapshotDto) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *OpenStructDlfSnapshotDto) SetChangelogRecordCount(v int64) *OpenStructDlfSnapshotDto {
	s.ChangelogRecordCount = &v
	return s
}

func (s *OpenStructDlfSnapshotDto) SetCommitKind(v string) *OpenStructDlfSnapshotDto {
	s.CommitKind = &v
	return s
}

func (s *OpenStructDlfSnapshotDto) SetDeltaRecordCount(v int64) *OpenStructDlfSnapshotDto {
	s.DeltaRecordCount = &v
	return s
}

func (s *OpenStructDlfSnapshotDto) SetSchemaId(v int64) *OpenStructDlfSnapshotDto {
	s.SchemaId = &v
	return s
}

func (s *OpenStructDlfSnapshotDto) SetSnapshotId(v int64) *OpenStructDlfSnapshotDto {
	s.SnapshotId = &v
	return s
}

func (s *OpenStructDlfSnapshotDto) SetTimeMillis(v int64) *OpenStructDlfSnapshotDto {
	s.TimeMillis = &v
	return s
}

func (s *OpenStructDlfSnapshotDto) SetTotalRecordCount(v int64) *OpenStructDlfSnapshotDto {
	s.TotalRecordCount = &v
	return s
}

func (s *OpenStructDlfSnapshotDto) Validate() error {
	return dara.Validate(s)
}
