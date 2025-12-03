// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMountPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DeleteMountPointRequest
	GetFileSystemId() *string
	SetInputRegionId(v string) *DeleteMountPointRequest
	GetInputRegionId() *string
	SetMountPointId(v string) *DeleteMountPointRequest
	GetMountPointId() *string
}

type DeleteMountPointRequest struct {
	// This parameter is required.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// This parameter is required.
	MountPointId *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
}

func (s DeleteMountPointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMountPointRequest) GoString() string {
	return s.String()
}

func (s *DeleteMountPointRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteMountPointRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *DeleteMountPointRequest) GetMountPointId() *string {
	return s.MountPointId
}

func (s *DeleteMountPointRequest) SetFileSystemId(v string) *DeleteMountPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteMountPointRequest) SetInputRegionId(v string) *DeleteMountPointRequest {
	s.InputRegionId = &v
	return s
}

func (s *DeleteMountPointRequest) SetMountPointId(v string) *DeleteMountPointRequest {
	s.MountPointId = &v
	return s
}

func (s *DeleteMountPointRequest) Validate() error {
	return dara.Validate(s)
}
