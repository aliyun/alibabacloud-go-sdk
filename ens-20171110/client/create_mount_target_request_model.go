// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMountTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *CreateMountTargetRequest
	GetEnsRegionId() *string
	SetFileSystemId(v string) *CreateMountTargetRequest
	GetFileSystemId() *string
	SetMountTargetName(v string) *CreateMountTargetRequest
	GetMountTargetName() *string
	SetNetWorkId(v string) *CreateMountTargetRequest
	GetNetWorkId() *string
}

type CreateMountTargetRequest struct {
	// The ID of the region.
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
	// The ID of the network.
	//
	// This parameter is required.
	//
	// example:
	//
	// n-*****
	NetWorkId *string `json:"NetWorkId,omitempty" xml:"NetWorkId,omitempty"`
}

func (s CreateMountTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMountTargetRequest) GoString() string {
	return s.String()
}

func (s *CreateMountTargetRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateMountTargetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateMountTargetRequest) GetMountTargetName() *string {
	return s.MountTargetName
}

func (s *CreateMountTargetRequest) GetNetWorkId() *string {
	return s.NetWorkId
}

func (s *CreateMountTargetRequest) SetEnsRegionId(v string) *CreateMountTargetRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateMountTargetRequest) SetFileSystemId(v string) *CreateMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateMountTargetRequest) SetMountTargetName(v string) *CreateMountTargetRequest {
	s.MountTargetName = &v
	return s
}

func (s *CreateMountTargetRequest) SetNetWorkId(v string) *CreateMountTargetRequest {
	s.NetWorkId = &v
	return s
}

func (s *CreateMountTargetRequest) Validate() error {
	return dara.Validate(s)
}
