// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileSystemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DeleteFileSystemRequest
	GetEnsRegionId() *string
	SetFileSystemId(v string) *DeleteFileSystemRequest
	GetFileSystemId() *string
}

type DeleteFileSystemRequest struct {
	// The ID of the edge node.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu-telecom-4
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the file system that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// e42640****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DeleteFileSystemRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileSystemRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileSystemRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DeleteFileSystemRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteFileSystemRequest) SetEnsRegionId(v string) *DeleteFileSystemRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DeleteFileSystemRequest) SetFileSystemId(v string) *DeleteFileSystemRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteFileSystemRequest) Validate() error {
	return dara.Validate(s)
}
