// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileSystemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DeleteFileSystemRequest
	GetFileSystemId() *string
}

type DeleteFileSystemRequest struct {
	// The ID of the file system that you want to delete.
	//
	// - General-purpose NAS: 31a8e4\\*\\*\\*\\*.
	//
	// - Extreme NAS: The ID must start with `extreme-`, such as extreme-0015\\*\\*\\*\\*.
	//
	// - Cloud Parallel File Storage (CPFS): The ID must start with `cpfs-`, such as cpfs-00cb6fa094ca\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DeleteFileSystemRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileSystemRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileSystemRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteFileSystemRequest) SetFileSystemId(v string) *DeleteFileSystemRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteFileSystemRequest) Validate() error {
	return dara.Validate(s)
}
