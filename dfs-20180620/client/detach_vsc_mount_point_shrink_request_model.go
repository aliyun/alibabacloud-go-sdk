// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachVscMountPointShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DetachVscMountPointShrinkRequest
	GetDescription() *string
	SetFileSystemId(v string) *DetachVscMountPointShrinkRequest
	GetFileSystemId() *string
	SetInputRegionId(v string) *DetachVscMountPointShrinkRequest
	GetInputRegionId() *string
	SetInstanceIdsShrink(v string) *DetachVscMountPointShrinkRequest
	GetInstanceIdsShrink() *string
	SetMountPointId(v string) *DetachVscMountPointShrinkRequest
	GetMountPointId() *string
	SetUseAssumeRoleChkServerPerm(v bool) *DetachVscMountPointShrinkRequest
	GetUseAssumeRoleChkServerPerm() *bool
	SetVscIdsShrink(v string) *DetachVscMountPointShrinkRequest
	GetVscIdsShrink() *string
}

type DetachVscMountPointShrinkRequest struct {
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
}

func (s DetachVscMountPointShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachVscMountPointShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetachVscMountPointShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *DetachVscMountPointShrinkRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DetachVscMountPointShrinkRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *DetachVscMountPointShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *DetachVscMountPointShrinkRequest) GetMountPointId() *string {
	return s.MountPointId
}

func (s *DetachVscMountPointShrinkRequest) GetUseAssumeRoleChkServerPerm() *bool {
	return s.UseAssumeRoleChkServerPerm
}

func (s *DetachVscMountPointShrinkRequest) GetVscIdsShrink() *string {
	return s.VscIdsShrink
}

func (s *DetachVscMountPointShrinkRequest) SetDescription(v string) *DetachVscMountPointShrinkRequest {
	s.Description = &v
	return s
}

func (s *DetachVscMountPointShrinkRequest) SetFileSystemId(v string) *DetachVscMountPointShrinkRequest {
	s.FileSystemId = &v
	return s
}

func (s *DetachVscMountPointShrinkRequest) SetInputRegionId(v string) *DetachVscMountPointShrinkRequest {
	s.InputRegionId = &v
	return s
}

func (s *DetachVscMountPointShrinkRequest) SetInstanceIdsShrink(v string) *DetachVscMountPointShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *DetachVscMountPointShrinkRequest) SetMountPointId(v string) *DetachVscMountPointShrinkRequest {
	s.MountPointId = &v
	return s
}

func (s *DetachVscMountPointShrinkRequest) SetUseAssumeRoleChkServerPerm(v bool) *DetachVscMountPointShrinkRequest {
	s.UseAssumeRoleChkServerPerm = &v
	return s
}

func (s *DetachVscMountPointShrinkRequest) SetVscIdsShrink(v string) *DetachVscMountPointShrinkRequest {
	s.VscIdsShrink = &v
	return s
}

func (s *DetachVscMountPointShrinkRequest) Validate() error {
	return dara.Validate(s)
}
