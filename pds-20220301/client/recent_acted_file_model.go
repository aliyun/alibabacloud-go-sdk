// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecentActedFile interface {
	dara.Model
	String() string
	GoString() string
	SetActionList(v []*string) *RecentActedFile
	GetActionList() []*string
	SetCategory(v string) *RecentActedFile
	GetCategory() *string
	SetDeleted(v bool) *RecentActedFile
	GetDeleted() *bool
	SetDriveId(v string) *RecentActedFile
	GetDriveId() *string
	SetDriveIsHandover(v bool) *RecentActedFile
	GetDriveIsHandover() *bool
	SetDriveName(v string) *RecentActedFile
	GetDriveName() *string
	SetDriveOwnerId(v string) *RecentActedFile
	GetDriveOwnerId() *string
	SetDriveOwnerType(v string) *RecentActedFile
	GetDriveOwnerType() *string
	SetFileId(v string) *RecentActedFile
	GetFileId() *string
	SetFileName(v string) *RecentActedFile
	GetFileName() *string
	SetSize(v int64) *RecentActedFile
	GetSize() *int64
	SetThumbnail(v string) *RecentActedFile
	GetThumbnail() *string
	SetTrashed(v bool) *RecentActedFile
	GetTrashed() *bool
}

type RecentActedFile struct {
	ActionList []*string `json:"action_list,omitempty" xml:"action_list,omitempty" type:"Repeated"`
	// example:
	//
	// doc
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// true
	Deleted *bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// example:
	//
	// 50d6f2aaa16525c7d053998e48fac265962f585f
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// example:
	//
	// false
	DriveIsHandover *bool `json:"drive_is_handover,omitempty" xml:"drive_is_handover,omitempty"`
	// example:
	//
	// group drive
	DriveName *string `json:"drive_name,omitempty" xml:"drive_name,omitempty"`
	// example:
	//
	// 50d6f2aaa16525c7d053998e48fac265962f585f
	DriveOwnerId *string `json:"drive_owner_id,omitempty" xml:"drive_owner_id,omitempty"`
	// example:
	//
	// group
	DriveOwnerType *string `json:"drive_owner_type,omitempty" xml:"drive_owner_type,omitempty"`
	// example:
	//
	// 50d6f2aaa16525c7d053998e48fac265962f585f
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// example:
	//
	// a.jpg
	FileName *string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// example:
	//
	// 100
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// https://xxx.jpg
	Thumbnail *string `json:"thumbnail,omitempty" xml:"thumbnail,omitempty"`
	// example:
	//
	// true
	Trashed *bool `json:"trashed,omitempty" xml:"trashed,omitempty"`
}

func (s RecentActedFile) String() string {
	return dara.Prettify(s)
}

func (s RecentActedFile) GoString() string {
	return s.String()
}

func (s *RecentActedFile) GetActionList() []*string {
	return s.ActionList
}

func (s *RecentActedFile) GetCategory() *string {
	return s.Category
}

func (s *RecentActedFile) GetDeleted() *bool {
	return s.Deleted
}

func (s *RecentActedFile) GetDriveId() *string {
	return s.DriveId
}

func (s *RecentActedFile) GetDriveIsHandover() *bool {
	return s.DriveIsHandover
}

func (s *RecentActedFile) GetDriveName() *string {
	return s.DriveName
}

func (s *RecentActedFile) GetDriveOwnerId() *string {
	return s.DriveOwnerId
}

func (s *RecentActedFile) GetDriveOwnerType() *string {
	return s.DriveOwnerType
}

func (s *RecentActedFile) GetFileId() *string {
	return s.FileId
}

func (s *RecentActedFile) GetFileName() *string {
	return s.FileName
}

func (s *RecentActedFile) GetSize() *int64 {
	return s.Size
}

func (s *RecentActedFile) GetThumbnail() *string {
	return s.Thumbnail
}

func (s *RecentActedFile) GetTrashed() *bool {
	return s.Trashed
}

func (s *RecentActedFile) SetActionList(v []*string) *RecentActedFile {
	s.ActionList = v
	return s
}

func (s *RecentActedFile) SetCategory(v string) *RecentActedFile {
	s.Category = &v
	return s
}

func (s *RecentActedFile) SetDeleted(v bool) *RecentActedFile {
	s.Deleted = &v
	return s
}

func (s *RecentActedFile) SetDriveId(v string) *RecentActedFile {
	s.DriveId = &v
	return s
}

func (s *RecentActedFile) SetDriveIsHandover(v bool) *RecentActedFile {
	s.DriveIsHandover = &v
	return s
}

func (s *RecentActedFile) SetDriveName(v string) *RecentActedFile {
	s.DriveName = &v
	return s
}

func (s *RecentActedFile) SetDriveOwnerId(v string) *RecentActedFile {
	s.DriveOwnerId = &v
	return s
}

func (s *RecentActedFile) SetDriveOwnerType(v string) *RecentActedFile {
	s.DriveOwnerType = &v
	return s
}

func (s *RecentActedFile) SetFileId(v string) *RecentActedFile {
	s.FileId = &v
	return s
}

func (s *RecentActedFile) SetFileName(v string) *RecentActedFile {
	s.FileName = &v
	return s
}

func (s *RecentActedFile) SetSize(v int64) *RecentActedFile {
	s.Size = &v
	return s
}

func (s *RecentActedFile) SetThumbnail(v string) *RecentActedFile {
	s.Thumbnail = &v
	return s
}

func (s *RecentActedFile) SetTrashed(v bool) *RecentActedFile {
	s.Trashed = &v
	return s
}

func (s *RecentActedFile) Validate() error {
	return dara.Validate(s)
}
