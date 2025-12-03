// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachVscMountPointShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *AttachVscMountPointShrinkRequest
	GetDescription() *string
	SetFileSystemId(v string) *AttachVscMountPointShrinkRequest
	GetFileSystemId() *string
	SetInputRegionId(v string) *AttachVscMountPointShrinkRequest
	GetInputRegionId() *string
	SetInstanceIdsShrink(v string) *AttachVscMountPointShrinkRequest
	GetInstanceIdsShrink() *string
	SetMountPointId(v string) *AttachVscMountPointShrinkRequest
	GetMountPointId() *string
	SetUseAssumeRoleChkServerPerm(v bool) *AttachVscMountPointShrinkRequest
	GetUseAssumeRoleChkServerPerm() *bool
	SetVscIdsShrink(v string) *AttachVscMountPointShrinkRequest
	GetVscIdsShrink() *string
	SetVscName(v string) *AttachVscMountPointShrinkRequest
	GetVscName() *string
	SetVscType(v string) *AttachVscMountPointShrinkRequest
	GetVscType() *string
}

type AttachVscMountPointShrinkRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 037****e1d
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// example:
	//
	// ["ecs-instance1", "ecs-instance2"]
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 037cb49e1d-c***5
	MountPointId               *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
	UseAssumeRoleChkServerPerm *bool   `json:"UseAssumeRoleChkServerPerm,omitempty" xml:"UseAssumeRoleChkServerPerm,omitempty"`
	VscIdsShrink               *string `json:"VscIds,omitempty" xml:"VscIds,omitempty"`
	VscName                    *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
	VscType                    *string `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s AttachVscMountPointShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachVscMountPointShrinkRequest) GoString() string {
	return s.String()
}

func (s *AttachVscMountPointShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *AttachVscMountPointShrinkRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *AttachVscMountPointShrinkRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *AttachVscMountPointShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *AttachVscMountPointShrinkRequest) GetMountPointId() *string {
	return s.MountPointId
}

func (s *AttachVscMountPointShrinkRequest) GetUseAssumeRoleChkServerPerm() *bool {
	return s.UseAssumeRoleChkServerPerm
}

func (s *AttachVscMountPointShrinkRequest) GetVscIdsShrink() *string {
	return s.VscIdsShrink
}

func (s *AttachVscMountPointShrinkRequest) GetVscName() *string {
	return s.VscName
}

func (s *AttachVscMountPointShrinkRequest) GetVscType() *string {
	return s.VscType
}

func (s *AttachVscMountPointShrinkRequest) SetDescription(v string) *AttachVscMountPointShrinkRequest {
	s.Description = &v
	return s
}

func (s *AttachVscMountPointShrinkRequest) SetFileSystemId(v string) *AttachVscMountPointShrinkRequest {
	s.FileSystemId = &v
	return s
}

func (s *AttachVscMountPointShrinkRequest) SetInputRegionId(v string) *AttachVscMountPointShrinkRequest {
	s.InputRegionId = &v
	return s
}

func (s *AttachVscMountPointShrinkRequest) SetInstanceIdsShrink(v string) *AttachVscMountPointShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *AttachVscMountPointShrinkRequest) SetMountPointId(v string) *AttachVscMountPointShrinkRequest {
	s.MountPointId = &v
	return s
}

func (s *AttachVscMountPointShrinkRequest) SetUseAssumeRoleChkServerPerm(v bool) *AttachVscMountPointShrinkRequest {
	s.UseAssumeRoleChkServerPerm = &v
	return s
}

func (s *AttachVscMountPointShrinkRequest) SetVscIdsShrink(v string) *AttachVscMountPointShrinkRequest {
	s.VscIdsShrink = &v
	return s
}

func (s *AttachVscMountPointShrinkRequest) SetVscName(v string) *AttachVscMountPointShrinkRequest {
	s.VscName = &v
	return s
}

func (s *AttachVscMountPointShrinkRequest) SetVscType(v string) *AttachVscMountPointShrinkRequest {
	s.VscType = &v
	return s
}

func (s *AttachVscMountPointShrinkRequest) Validate() error {
	return dara.Validate(s)
}
