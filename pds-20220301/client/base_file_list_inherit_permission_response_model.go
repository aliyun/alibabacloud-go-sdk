// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBaseFileListInheritPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v string) *BaseFileListInheritPermissionResponse
	GetFileId() *string
	SetMember(v *FilePermissionMember) *BaseFileListInheritPermissionResponse
	GetMember() *FilePermissionMember
}

type BaseFileListInheritPermissionResponse struct {
	FileId *string               `json:"file_id,omitempty" xml:"file_id,omitempty"`
	Member *FilePermissionMember `json:"member,omitempty" xml:"member,omitempty"`
}

func (s BaseFileListInheritPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s BaseFileListInheritPermissionResponse) GoString() string {
	return s.String()
}

func (s *BaseFileListInheritPermissionResponse) GetFileId() *string {
	return s.FileId
}

func (s *BaseFileListInheritPermissionResponse) GetMember() *FilePermissionMember {
	return s.Member
}

func (s *BaseFileListInheritPermissionResponse) SetFileId(v string) *BaseFileListInheritPermissionResponse {
	s.FileId = &v
	return s
}

func (s *BaseFileListInheritPermissionResponse) SetMember(v *FilePermissionMember) *BaseFileListInheritPermissionResponse {
	s.Member = v
	return s
}

func (s *BaseFileListInheritPermissionResponse) Validate() error {
	if s.Member != nil {
		if err := s.Member.Validate(); err != nil {
			return err
		}
	}
	return nil
}
