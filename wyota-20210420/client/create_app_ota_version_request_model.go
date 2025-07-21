// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppOtaVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppVersion(v string) *CreateAppOtaVersionRequest
	GetAppVersion() *string
	SetArch(v string) *CreateAppOtaVersionRequest
	GetArch() *string
	SetChannel(v string) *CreateAppOtaVersionRequest
	GetChannel() *string
	SetClientType(v int32) *CreateAppOtaVersionRequest
	GetClientType() *int32
	SetCreator(v string) *CreateAppOtaVersionRequest
	GetCreator() *string
	SetDownloadUrl(v string) *CreateAppOtaVersionRequest
	GetDownloadUrl() *string
	SetMd5(v string) *CreateAppOtaVersionRequest
	GetMd5() *string
	SetOs(v string) *CreateAppOtaVersionRequest
	GetOs() *string
	SetOsType(v string) *CreateAppOtaVersionRequest
	GetOsType() *string
	SetOtaType(v int32) *CreateAppOtaVersionRequest
	GetOtaType() *int32
	SetProject(v string) *CreateAppOtaVersionRequest
	GetProject() *string
	SetRelationVersionUids(v []*string) *CreateAppOtaVersionRequest
	GetRelationVersionUids() []*string
	SetReleaseNote(v string) *CreateAppOtaVersionRequest
	GetReleaseNote() *string
	SetReleaseNoteEn(v string) *CreateAppOtaVersionRequest
	GetReleaseNoteEn() *string
	SetReleaseNoteJp(v string) *CreateAppOtaVersionRequest
	GetReleaseNoteJp() *string
	SetSize(v int64) *CreateAppOtaVersionRequest
	GetSize() *int64
	SetSnapshotId(v string) *CreateAppOtaVersionRequest
	GetSnapshotId() *string
	SetSnapshotRegionId(v string) *CreateAppOtaVersionRequest
	GetSnapshotRegionId() *string
	SetStatus(v int32) *CreateAppOtaVersionRequest
	GetStatus() *int32
	SetVersionType(v string) *CreateAppOtaVersionRequest
	GetVersionType() *string
}

type CreateAppOtaVersionRequest struct {
	AppVersion          *string   `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	Arch                *string   `json:"Arch,omitempty" xml:"Arch,omitempty"`
	Channel             *string   `json:"Channel,omitempty" xml:"Channel,omitempty"`
	ClientType          *int32    `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	Creator             *string   `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DownloadUrl         *string   `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	Md5                 *string   `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Os                  *string   `json:"Os,omitempty" xml:"Os,omitempty"`
	OsType              *string   `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OtaType             *int32    `json:"OtaType,omitempty" xml:"OtaType,omitempty"`
	Project             *string   `json:"Project,omitempty" xml:"Project,omitempty"`
	RelationVersionUids []*string `json:"RelationVersionUids,omitempty" xml:"RelationVersionUids,omitempty" type:"Repeated"`
	ReleaseNote         *string   `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	ReleaseNoteEn       *string   `json:"ReleaseNoteEn,omitempty" xml:"ReleaseNoteEn,omitempty"`
	ReleaseNoteJp       *string   `json:"ReleaseNoteJp,omitempty" xml:"ReleaseNoteJp,omitempty"`
	Size                *int64    `json:"Size,omitempty" xml:"Size,omitempty"`
	SnapshotId          *string   `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SnapshotRegionId    *string   `json:"SnapshotRegionId,omitempty" xml:"SnapshotRegionId,omitempty"`
	Status              *int32    `json:"Status,omitempty" xml:"Status,omitempty"`
	VersionType         *string   `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s CreateAppOtaVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppOtaVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateAppOtaVersionRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *CreateAppOtaVersionRequest) GetArch() *string {
	return s.Arch
}

func (s *CreateAppOtaVersionRequest) GetChannel() *string {
	return s.Channel
}

func (s *CreateAppOtaVersionRequest) GetClientType() *int32 {
	return s.ClientType
}

func (s *CreateAppOtaVersionRequest) GetCreator() *string {
	return s.Creator
}

func (s *CreateAppOtaVersionRequest) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *CreateAppOtaVersionRequest) GetMd5() *string {
	return s.Md5
}

func (s *CreateAppOtaVersionRequest) GetOs() *string {
	return s.Os
}

func (s *CreateAppOtaVersionRequest) GetOsType() *string {
	return s.OsType
}

func (s *CreateAppOtaVersionRequest) GetOtaType() *int32 {
	return s.OtaType
}

func (s *CreateAppOtaVersionRequest) GetProject() *string {
	return s.Project
}

func (s *CreateAppOtaVersionRequest) GetRelationVersionUids() []*string {
	return s.RelationVersionUids
}

func (s *CreateAppOtaVersionRequest) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *CreateAppOtaVersionRequest) GetReleaseNoteEn() *string {
	return s.ReleaseNoteEn
}

func (s *CreateAppOtaVersionRequest) GetReleaseNoteJp() *string {
	return s.ReleaseNoteJp
}

func (s *CreateAppOtaVersionRequest) GetSize() *int64 {
	return s.Size
}

func (s *CreateAppOtaVersionRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateAppOtaVersionRequest) GetSnapshotRegionId() *string {
	return s.SnapshotRegionId
}

func (s *CreateAppOtaVersionRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CreateAppOtaVersionRequest) GetVersionType() *string {
	return s.VersionType
}

func (s *CreateAppOtaVersionRequest) SetAppVersion(v string) *CreateAppOtaVersionRequest {
	s.AppVersion = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetArch(v string) *CreateAppOtaVersionRequest {
	s.Arch = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetChannel(v string) *CreateAppOtaVersionRequest {
	s.Channel = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetClientType(v int32) *CreateAppOtaVersionRequest {
	s.ClientType = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetCreator(v string) *CreateAppOtaVersionRequest {
	s.Creator = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetDownloadUrl(v string) *CreateAppOtaVersionRequest {
	s.DownloadUrl = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetMd5(v string) *CreateAppOtaVersionRequest {
	s.Md5 = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetOs(v string) *CreateAppOtaVersionRequest {
	s.Os = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetOsType(v string) *CreateAppOtaVersionRequest {
	s.OsType = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetOtaType(v int32) *CreateAppOtaVersionRequest {
	s.OtaType = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetProject(v string) *CreateAppOtaVersionRequest {
	s.Project = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetRelationVersionUids(v []*string) *CreateAppOtaVersionRequest {
	s.RelationVersionUids = v
	return s
}

func (s *CreateAppOtaVersionRequest) SetReleaseNote(v string) *CreateAppOtaVersionRequest {
	s.ReleaseNote = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetReleaseNoteEn(v string) *CreateAppOtaVersionRequest {
	s.ReleaseNoteEn = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetReleaseNoteJp(v string) *CreateAppOtaVersionRequest {
	s.ReleaseNoteJp = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetSize(v int64) *CreateAppOtaVersionRequest {
	s.Size = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetSnapshotId(v string) *CreateAppOtaVersionRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetSnapshotRegionId(v string) *CreateAppOtaVersionRequest {
	s.SnapshotRegionId = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetStatus(v int32) *CreateAppOtaVersionRequest {
	s.Status = &v
	return s
}

func (s *CreateAppOtaVersionRequest) SetVersionType(v string) *CreateAppOtaVersionRequest {
	s.VersionType = &v
	return s
}

func (s *CreateAppOtaVersionRequest) Validate() error {
	return dara.Validate(s)
}
