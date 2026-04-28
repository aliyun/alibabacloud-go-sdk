// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileAddPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *FileAddPermissionRequest
	GetDriveId() *string
	SetFileId(v string) *FileAddPermissionRequest
	GetFileId() *string
	SetMemberList(v []*FilePermissionMember) *FileAddPermissionRequest
	GetMemberList() []*FilePermissionMember
}

type FileAddPermissionRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The ID of the folder. If you want to authorize a user or group to access a team drive, set this parameter to root. If you want to authorize a user or group to access an individual drive, you cannot set this parameter to root.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4221bf6e6ab43c255edc4463bf3a6f5f5d317406
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The members that are authorized to access files.
	//
	// This parameter is required.
	MemberList []*FilePermissionMember `json:"member_list,omitempty" xml:"member_list,omitempty" type:"Repeated"`
}

func (s FileAddPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s FileAddPermissionRequest) GoString() string {
	return s.String()
}

func (s *FileAddPermissionRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *FileAddPermissionRequest) GetFileId() *string {
	return s.FileId
}

func (s *FileAddPermissionRequest) GetMemberList() []*FilePermissionMember {
	return s.MemberList
}

func (s *FileAddPermissionRequest) SetDriveId(v string) *FileAddPermissionRequest {
	s.DriveId = &v
	return s
}

func (s *FileAddPermissionRequest) SetFileId(v string) *FileAddPermissionRequest {
	s.FileId = &v
	return s
}

func (s *FileAddPermissionRequest) SetMemberList(v []*FilePermissionMember) *FileAddPermissionRequest {
	s.MemberList = v
	return s
}

func (s *FileAddPermissionRequest) Validate() error {
	if s.MemberList != nil {
		for _, item := range s.MemberList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
