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
	SetInputRegionId(v string) *DeleteFileSystemRequest
	GetInputRegionId() *string
}

type DeleteFileSystemRequest struct {
	// This parameter is required.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
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

func (s *DeleteFileSystemRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *DeleteFileSystemRequest) SetFileSystemId(v string) *DeleteFileSystemRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteFileSystemRequest) SetInputRegionId(v string) *DeleteFileSystemRequest {
	s.InputRegionId = &v
	return s
}

func (s *DeleteFileSystemRequest) Validate() error {
	return dara.Validate(s)
}
