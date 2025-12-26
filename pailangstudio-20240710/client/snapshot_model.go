// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSnapshot interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *Snapshot
	GetAccessibility() *string
	SetCreationType(v string) *Snapshot
	GetCreationType() *string
	SetCreator(v string) *Snapshot
	GetCreator() *string
	SetDescription(v string) *Snapshot
	GetDescription() *string
	SetGmtCreateTime(v string) *Snapshot
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *Snapshot
	GetGmtModifiedTime() *string
	SetSnapshotId(v string) *Snapshot
	GetSnapshotId() *string
	SetSnapshotName(v string) *Snapshot
	GetSnapshotName() *string
	SetSnapshotResourceId(v string) *Snapshot
	GetSnapshotResourceId() *string
	SetSnapshotResourceType(v string) *Snapshot
	GetSnapshotResourceType() *string
	SetSnapshotStoragePath(v string) *Snapshot
	GetSnapshotStoragePath() *string
	SetSnapshotUrl(v string) *Snapshot
	GetSnapshotUrl() *string
	SetWorkDir(v string) *Snapshot
	GetWorkDir() *string
	SetWorkspaceId(v string) *Snapshot
	GetWorkspaceId() *string
}

type Snapshot struct {
	Accessibility        *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	CreationType         *string `json:"CreationType,omitempty" xml:"CreationType,omitempty"`
	Creator              *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime        *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime      *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	SnapshotId           *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SnapshotName         *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	SnapshotResourceId   *string `json:"SnapshotResourceId,omitempty" xml:"SnapshotResourceId,omitempty"`
	SnapshotResourceType *string `json:"SnapshotResourceType,omitempty" xml:"SnapshotResourceType,omitempty"`
	SnapshotStoragePath  *string `json:"SnapshotStoragePath,omitempty" xml:"SnapshotStoragePath,omitempty"`
	SnapshotUrl          *string `json:"SnapshotUrl,omitempty" xml:"SnapshotUrl,omitempty"`
	WorkDir              *string `json:"WorkDir,omitempty" xml:"WorkDir,omitempty"`
	WorkspaceId          *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Snapshot) String() string {
	return dara.Prettify(s)
}

func (s Snapshot) GoString() string {
	return s.String()
}

func (s *Snapshot) GetAccessibility() *string {
	return s.Accessibility
}

func (s *Snapshot) GetCreationType() *string {
	return s.CreationType
}

func (s *Snapshot) GetCreator() *string {
	return s.Creator
}

func (s *Snapshot) GetDescription() *string {
	return s.Description
}

func (s *Snapshot) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *Snapshot) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *Snapshot) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *Snapshot) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *Snapshot) GetSnapshotResourceId() *string {
	return s.SnapshotResourceId
}

func (s *Snapshot) GetSnapshotResourceType() *string {
	return s.SnapshotResourceType
}

func (s *Snapshot) GetSnapshotStoragePath() *string {
	return s.SnapshotStoragePath
}

func (s *Snapshot) GetSnapshotUrl() *string {
	return s.SnapshotUrl
}

func (s *Snapshot) GetWorkDir() *string {
	return s.WorkDir
}

func (s *Snapshot) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *Snapshot) SetAccessibility(v string) *Snapshot {
	s.Accessibility = &v
	return s
}

func (s *Snapshot) SetCreationType(v string) *Snapshot {
	s.CreationType = &v
	return s
}

func (s *Snapshot) SetCreator(v string) *Snapshot {
	s.Creator = &v
	return s
}

func (s *Snapshot) SetDescription(v string) *Snapshot {
	s.Description = &v
	return s
}

func (s *Snapshot) SetGmtCreateTime(v string) *Snapshot {
	s.GmtCreateTime = &v
	return s
}

func (s *Snapshot) SetGmtModifiedTime(v string) *Snapshot {
	s.GmtModifiedTime = &v
	return s
}

func (s *Snapshot) SetSnapshotId(v string) *Snapshot {
	s.SnapshotId = &v
	return s
}

func (s *Snapshot) SetSnapshotName(v string) *Snapshot {
	s.SnapshotName = &v
	return s
}

func (s *Snapshot) SetSnapshotResourceId(v string) *Snapshot {
	s.SnapshotResourceId = &v
	return s
}

func (s *Snapshot) SetSnapshotResourceType(v string) *Snapshot {
	s.SnapshotResourceType = &v
	return s
}

func (s *Snapshot) SetSnapshotStoragePath(v string) *Snapshot {
	s.SnapshotStoragePath = &v
	return s
}

func (s *Snapshot) SetSnapshotUrl(v string) *Snapshot {
	s.SnapshotUrl = &v
	return s
}

func (s *Snapshot) SetWorkDir(v string) *Snapshot {
	s.WorkDir = &v
	return s
}

func (s *Snapshot) SetWorkspaceId(v string) *Snapshot {
	s.WorkspaceId = &v
	return s
}

func (s *Snapshot) Validate() error {
	return dara.Validate(s)
}
