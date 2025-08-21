// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMountTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DeleteMountTargetRequest
	GetEnsRegionId() *string
	SetFileSystemId(v string) *DeleteMountTargetRequest
	GetFileSystemId() *string
	SetMountTargetName(v string) *DeleteMountTargetRequest
	GetMountTargetName() *string
}

type DeleteMountTargetRequest struct {
	// The ID of the edge node.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing-cmcc
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// c50f8*****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The name of the mount target.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestMountPath
	MountTargetName *string `json:"MountTargetName,omitempty" xml:"MountTargetName,omitempty"`
}

func (s DeleteMountTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMountTargetRequest) GoString() string {
	return s.String()
}

func (s *DeleteMountTargetRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DeleteMountTargetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteMountTargetRequest) GetMountTargetName() *string {
	return s.MountTargetName
}

func (s *DeleteMountTargetRequest) SetEnsRegionId(v string) *DeleteMountTargetRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DeleteMountTargetRequest) SetFileSystemId(v string) *DeleteMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteMountTargetRequest) SetMountTargetName(v string) *DeleteMountTargetRequest {
	s.MountTargetName = &v
	return s
}

func (s *DeleteMountTargetRequest) Validate() error {
	return dara.Validate(s)
}
