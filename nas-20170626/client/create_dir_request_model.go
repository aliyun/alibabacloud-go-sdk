// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDirRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *CreateDirRequest
	GetFileSystemId() *string
	SetOwnerGroupId(v int32) *CreateDirRequest
	GetOwnerGroupId() *int32
	SetOwnerUserId(v int32) *CreateDirRequest
	GetOwnerUserId() *int32
	SetPermission(v string) *CreateDirRequest
	GetPermission() *string
	SetRecursion(v bool) *CreateDirRequest
	GetRecursion() *bool
	SetRootDirectory(v string) *CreateDirRequest
	GetRootDirectory() *string
}

type CreateDirRequest struct {
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The user group ID of the directory owner. Valid values: 0 to 4294967295.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	OwnerGroupId *int32 `json:"OwnerGroupId,omitempty" xml:"OwnerGroupId,omitempty"`
	// The user ID of the directory owner. Valid values: 0 to 4294967295.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	OwnerUserId *int32 `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The POSIX permissions to apply to the root directory path. The value is a valid octal number, such as 0755.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0755
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// Specifies whether to create a multi-level directory. Valid values:
	//
	// - true (default): If the multi-level directory does not exist, the system creates the directories level by level.
	//
	// - false: Only the last-level directory is created. If the parent directory does not exist, an error message is returned.
	//
	// example:
	//
	// true
	Recursion *bool `json:"Recursion,omitempty" xml:"Recursion,omitempty"`
	// The directory name.
	//
	// - The name must start with a forward slash (/).
	//
	// - The name can contain digits and uppercase and lowercase letters.
	//
	// - The name can contain underscores (_), hyphens (-), or periods (.).
	//
	// - The path cannot contain symbolic links, such as the current directory (.), the parent directory (..), or other symbolic links.
	//
	// >  - If the root directory does not exist, configure the automatic creation information. The system automatically creates the specified root directory based on the settings.
	//
	//  > - If the root directory already exists, you do not need to configure the automatic creation information. If you configure the information, the automatic creation settings are ignored.
	//
	// This parameter is required.
	//
	// example:
	//
	// /test
	RootDirectory *string `json:"RootDirectory,omitempty" xml:"RootDirectory,omitempty"`
}

func (s CreateDirRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDirRequest) GoString() string {
	return s.String()
}

func (s *CreateDirRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateDirRequest) GetOwnerGroupId() *int32 {
	return s.OwnerGroupId
}

func (s *CreateDirRequest) GetOwnerUserId() *int32 {
	return s.OwnerUserId
}

func (s *CreateDirRequest) GetPermission() *string {
	return s.Permission
}

func (s *CreateDirRequest) GetRecursion() *bool {
	return s.Recursion
}

func (s *CreateDirRequest) GetRootDirectory() *string {
	return s.RootDirectory
}

func (s *CreateDirRequest) SetFileSystemId(v string) *CreateDirRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateDirRequest) SetOwnerGroupId(v int32) *CreateDirRequest {
	s.OwnerGroupId = &v
	return s
}

func (s *CreateDirRequest) SetOwnerUserId(v int32) *CreateDirRequest {
	s.OwnerUserId = &v
	return s
}

func (s *CreateDirRequest) SetPermission(v string) *CreateDirRequest {
	s.Permission = &v
	return s
}

func (s *CreateDirRequest) SetRecursion(v bool) *CreateDirRequest {
	s.Recursion = &v
	return s
}

func (s *CreateDirRequest) SetRootDirectory(v string) *CreateDirRequest {
	s.RootDirectory = &v
	return s
}

func (s *CreateDirRequest) Validate() error {
	return dara.Validate(s)
}
