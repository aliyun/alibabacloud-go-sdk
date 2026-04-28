// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckNameMode(v string) *UpdateFileRequest
	GetCheckNameMode() *string
	SetDescription(v string) *UpdateFileRequest
	GetDescription() *string
	SetDriveId(v string) *UpdateFileRequest
	GetDriveId() *string
	SetFileId(v string) *UpdateFileRequest
	GetFileId() *string
	SetHidden(v bool) *UpdateFileRequest
	GetHidden() *bool
	SetLabels(v []*string) *UpdateFileRequest
	GetLabels() []*string
	SetLocalModifiedAt(v string) *UpdateFileRequest
	GetLocalModifiedAt() *string
	SetName(v string) *UpdateFileRequest
	GetName() *string
	SetStarred(v bool) *UpdateFileRequest
	GetStarred() *bool
}

type UpdateFileRequest struct {
	// The processing method that is used if the file that you want to modify has the same name as an existing file on the cloud. Valid values:
	//
	// ignore: allows you to modify the file by using the same name as an existing file on the cloud.
	//
	// auto_rename: automatically renames the file that has the same name on the cloud. By default, the current point in time is added to the end of the file name. Example: xxx_20060102_150405.
	//
	// refuse: does not modify the file that you want to modify but returns the information about the file that has the same name on the cloud.
	//
	// Default value: ignore.
	//
	// example:
	//
	// ignore
	CheckNameMode *string `json:"check_name_mode,omitempty" xml:"check_name_mode,omitempty"`
	// The description of the file. The description can be up to 1,024 characters in length.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// Specifies whether to hide the file.
	//
	// example:
	//
	// true
	Hidden *bool `json:"hidden,omitempty" xml:"hidden,omitempty"`
	// The tags of the file. You can specify up to 100 tags.
	Labels []*string `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// The local time when the file was modified. The time is in the yyyy-MM-ddTHH:mm:ssZ format based on the UTC+0 time zone.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	LocalModifiedAt *string `json:"local_modified_at,omitempty" xml:"local_modified_at,omitempty"`
	// The name of the file. The name can be up to 1,024 bytes in length based on the UTF-8 encoding rule.
	//
	// example:
	//
	// a.jpg
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specifies whether to add the file to favorites.
	//
	// example:
	//
	// true
	Starred *bool `json:"starred,omitempty" xml:"starred,omitempty"`
}

func (s UpdateFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileRequest) GetCheckNameMode() *string {
	return s.CheckNameMode
}

func (s *UpdateFileRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateFileRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *UpdateFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *UpdateFileRequest) GetHidden() *bool {
	return s.Hidden
}

func (s *UpdateFileRequest) GetLabels() []*string {
	return s.Labels
}

func (s *UpdateFileRequest) GetLocalModifiedAt() *string {
	return s.LocalModifiedAt
}

func (s *UpdateFileRequest) GetName() *string {
	return s.Name
}

func (s *UpdateFileRequest) GetStarred() *bool {
	return s.Starred
}

func (s *UpdateFileRequest) SetCheckNameMode(v string) *UpdateFileRequest {
	s.CheckNameMode = &v
	return s
}

func (s *UpdateFileRequest) SetDescription(v string) *UpdateFileRequest {
	s.Description = &v
	return s
}

func (s *UpdateFileRequest) SetDriveId(v string) *UpdateFileRequest {
	s.DriveId = &v
	return s
}

func (s *UpdateFileRequest) SetFileId(v string) *UpdateFileRequest {
	s.FileId = &v
	return s
}

func (s *UpdateFileRequest) SetHidden(v bool) *UpdateFileRequest {
	s.Hidden = &v
	return s
}

func (s *UpdateFileRequest) SetLabels(v []*string) *UpdateFileRequest {
	s.Labels = v
	return s
}

func (s *UpdateFileRequest) SetLocalModifiedAt(v string) *UpdateFileRequest {
	s.LocalModifiedAt = &v
	return s
}

func (s *UpdateFileRequest) SetName(v string) *UpdateFileRequest {
	s.Name = &v
	return s
}

func (s *UpdateFileRequest) SetStarred(v bool) *UpdateFileRequest {
	s.Starred = &v
	return s
}

func (s *UpdateFileRequest) Validate() error {
	return dara.Validate(s)
}
