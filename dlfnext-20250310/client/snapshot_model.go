// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSnapshot interface {
	dara.Model
	String() string
	GoString() string
	SetBaseManifestList(v string) *Snapshot
	GetBaseManifestList() *string
	SetChangelogManifestList(v string) *Snapshot
	GetChangelogManifestList() *string
	SetChangelogRecordCount(v int64) *Snapshot
	GetChangelogRecordCount() *int64
	SetCommitIdentifier(v int64) *Snapshot
	GetCommitIdentifier() *int64
	SetCommitKind(v string) *Snapshot
	GetCommitKind() *string
	SetCommitUser(v string) *Snapshot
	GetCommitUser() *string
	SetDeltaManifestList(v string) *Snapshot
	GetDeltaManifestList() *string
	SetDeltaRecordCount(v int64) *Snapshot
	GetDeltaRecordCount() *int64
	SetId(v int64) *Snapshot
	GetId() *int64
	SetIndexManifest(v string) *Snapshot
	GetIndexManifest() *string
	SetLogOffsets(v map[string]*int64) *Snapshot
	GetLogOffsets() map[string]*int64
	SetSchemaId(v int64) *Snapshot
	GetSchemaId() *int64
	SetStatistics(v string) *Snapshot
	GetStatistics() *string
	SetTimeMillis(v int64) *Snapshot
	GetTimeMillis() *int64
	SetTotalRecordCount(v int64) *Snapshot
	GetTotalRecordCount() *int64
	SetVersion(v int32) *Snapshot
	GetVersion() *int32
	SetWatermark(v int64) *Snapshot
	GetWatermark() *int64
}

type Snapshot struct {
	BaseManifestList *string `json:"baseManifestList,omitempty" xml:"baseManifestList,omitempty"`
	// if can be null:
	// true
	ChangelogManifestList *string           `json:"changelogManifestList,omitempty" xml:"changelogManifestList,omitempty"`
	ChangelogRecordCount  *int64            `json:"changelogRecordCount,omitempty" xml:"changelogRecordCount,omitempty"`
	CommitIdentifier      *int64            `json:"commitIdentifier,omitempty" xml:"commitIdentifier,omitempty"`
	CommitKind            *string           `json:"commitKind,omitempty" xml:"commitKind,omitempty"`
	CommitUser            *string           `json:"commitUser,omitempty" xml:"commitUser,omitempty"`
	DeltaManifestList     *string           `json:"deltaManifestList,omitempty" xml:"deltaManifestList,omitempty"`
	DeltaRecordCount      *int64            `json:"deltaRecordCount,omitempty" xml:"deltaRecordCount,omitempty"`
	Id                    *int64            `json:"id,omitempty" xml:"id,omitempty"`
	IndexManifest         *string           `json:"indexManifest,omitempty" xml:"indexManifest,omitempty"`
	LogOffsets            map[string]*int64 `json:"logOffsets,omitempty" xml:"logOffsets,omitempty"`
	SchemaId              *int64            `json:"schemaId,omitempty" xml:"schemaId,omitempty"`
	Statistics            *string           `json:"statistics,omitempty" xml:"statistics,omitempty"`
	TimeMillis            *int64            `json:"timeMillis,omitempty" xml:"timeMillis,omitempty"`
	TotalRecordCount      *int64            `json:"totalRecordCount,omitempty" xml:"totalRecordCount,omitempty"`
	// if can be null:
	// true
	Version   *int32 `json:"version,omitempty" xml:"version,omitempty"`
	Watermark *int64 `json:"watermark,omitempty" xml:"watermark,omitempty"`
}

func (s Snapshot) String() string {
	return dara.Prettify(s)
}

func (s Snapshot) GoString() string {
	return s.String()
}

func (s *Snapshot) GetBaseManifestList() *string {
	return s.BaseManifestList
}

func (s *Snapshot) GetChangelogManifestList() *string {
	return s.ChangelogManifestList
}

func (s *Snapshot) GetChangelogRecordCount() *int64 {
	return s.ChangelogRecordCount
}

func (s *Snapshot) GetCommitIdentifier() *int64 {
	return s.CommitIdentifier
}

func (s *Snapshot) GetCommitKind() *string {
	return s.CommitKind
}

func (s *Snapshot) GetCommitUser() *string {
	return s.CommitUser
}

func (s *Snapshot) GetDeltaManifestList() *string {
	return s.DeltaManifestList
}

func (s *Snapshot) GetDeltaRecordCount() *int64 {
	return s.DeltaRecordCount
}

func (s *Snapshot) GetId() *int64 {
	return s.Id
}

func (s *Snapshot) GetIndexManifest() *string {
	return s.IndexManifest
}

func (s *Snapshot) GetLogOffsets() map[string]*int64 {
	return s.LogOffsets
}

func (s *Snapshot) GetSchemaId() *int64 {
	return s.SchemaId
}

func (s *Snapshot) GetStatistics() *string {
	return s.Statistics
}

func (s *Snapshot) GetTimeMillis() *int64 {
	return s.TimeMillis
}

func (s *Snapshot) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *Snapshot) GetVersion() *int32 {
	return s.Version
}

func (s *Snapshot) GetWatermark() *int64 {
	return s.Watermark
}

func (s *Snapshot) SetBaseManifestList(v string) *Snapshot {
	s.BaseManifestList = &v
	return s
}

func (s *Snapshot) SetChangelogManifestList(v string) *Snapshot {
	s.ChangelogManifestList = &v
	return s
}

func (s *Snapshot) SetChangelogRecordCount(v int64) *Snapshot {
	s.ChangelogRecordCount = &v
	return s
}

func (s *Snapshot) SetCommitIdentifier(v int64) *Snapshot {
	s.CommitIdentifier = &v
	return s
}

func (s *Snapshot) SetCommitKind(v string) *Snapshot {
	s.CommitKind = &v
	return s
}

func (s *Snapshot) SetCommitUser(v string) *Snapshot {
	s.CommitUser = &v
	return s
}

func (s *Snapshot) SetDeltaManifestList(v string) *Snapshot {
	s.DeltaManifestList = &v
	return s
}

func (s *Snapshot) SetDeltaRecordCount(v int64) *Snapshot {
	s.DeltaRecordCount = &v
	return s
}

func (s *Snapshot) SetId(v int64) *Snapshot {
	s.Id = &v
	return s
}

func (s *Snapshot) SetIndexManifest(v string) *Snapshot {
	s.IndexManifest = &v
	return s
}

func (s *Snapshot) SetLogOffsets(v map[string]*int64) *Snapshot {
	s.LogOffsets = v
	return s
}

func (s *Snapshot) SetSchemaId(v int64) *Snapshot {
	s.SchemaId = &v
	return s
}

func (s *Snapshot) SetStatistics(v string) *Snapshot {
	s.Statistics = &v
	return s
}

func (s *Snapshot) SetTimeMillis(v int64) *Snapshot {
	s.TimeMillis = &v
	return s
}

func (s *Snapshot) SetTotalRecordCount(v int64) *Snapshot {
	s.TotalRecordCount = &v
	return s
}

func (s *Snapshot) SetVersion(v int32) *Snapshot {
	s.Version = &v
	return s
}

func (s *Snapshot) SetWatermark(v int64) *Snapshot {
	s.Watermark = &v
	return s
}

func (s *Snapshot) Validate() error {
	return dara.Validate(s)
}
