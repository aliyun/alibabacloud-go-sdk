// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNASDefaultMountTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *ModifyNASDefaultMountTargetRequest
	GetFileSystemId() *string
	SetMountTargetDomain(v string) *ModifyNASDefaultMountTargetRequest
	GetMountTargetDomain() *string
	SetRegionId(v string) *ModifyNASDefaultMountTargetRequest
	GetRegionId() *string
}

type ModifyNASDefaultMountTargetRequest struct {
	// The ID of the NAS file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 04f314****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The domain name of the mount target.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0bf744****-xo***.cn-hangzhou.nas.aliyuncs.com
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyNASDefaultMountTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNASDefaultMountTargetRequest) GoString() string {
	return s.String()
}

func (s *ModifyNASDefaultMountTargetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyNASDefaultMountTargetRequest) GetMountTargetDomain() *string {
	return s.MountTargetDomain
}

func (s *ModifyNASDefaultMountTargetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyNASDefaultMountTargetRequest) SetFileSystemId(v string) *ModifyNASDefaultMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyNASDefaultMountTargetRequest) SetMountTargetDomain(v string) *ModifyNASDefaultMountTargetRequest {
	s.MountTargetDomain = &v
	return s
}

func (s *ModifyNASDefaultMountTargetRequest) SetRegionId(v string) *ModifyNASDefaultMountTargetRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNASDefaultMountTargetRequest) Validate() error {
	return dara.Validate(s)
}
