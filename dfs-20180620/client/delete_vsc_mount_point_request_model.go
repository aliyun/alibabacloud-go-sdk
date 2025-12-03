// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVscMountPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DeleteVscMountPointRequest
	GetFileSystemId() *string
	SetInputRegionId(v string) *DeleteVscMountPointRequest
	GetInputRegionId() *string
	SetMountPointId(v string) *DeleteVscMountPointRequest
	GetMountPointId() *string
}

type DeleteVscMountPointRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 037c****1d
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 037cb49e1d-c***5
	MountPointId *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
}

func (s DeleteVscMountPointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVscMountPointRequest) GoString() string {
	return s.String()
}

func (s *DeleteVscMountPointRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteVscMountPointRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *DeleteVscMountPointRequest) GetMountPointId() *string {
	return s.MountPointId
}

func (s *DeleteVscMountPointRequest) SetFileSystemId(v string) *DeleteVscMountPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteVscMountPointRequest) SetInputRegionId(v string) *DeleteVscMountPointRequest {
	s.InputRegionId = &v
	return s
}

func (s *DeleteVscMountPointRequest) SetMountPointId(v string) *DeleteVscMountPointRequest {
	s.MountPointId = &v
	return s
}

func (s *DeleteVscMountPointRequest) Validate() error {
	return dara.Validate(s)
}
