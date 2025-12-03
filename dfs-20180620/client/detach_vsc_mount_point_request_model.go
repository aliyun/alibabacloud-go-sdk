// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachVscMountPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DetachVscMountPointRequest
	GetDescription() *string
	SetFileSystemId(v string) *DetachVscMountPointRequest
	GetFileSystemId() *string
	SetInputRegionId(v string) *DetachVscMountPointRequest
	GetInputRegionId() *string
	SetInstanceIds(v []*string) *DetachVscMountPointRequest
	GetInstanceIds() []*string
	SetMountPointId(v string) *DetachVscMountPointRequest
	GetMountPointId() *string
	SetUseAssumeRoleChkServerPerm(v bool) *DetachVscMountPointRequest
	GetUseAssumeRoleChkServerPerm() *bool
	SetVscIds(v []*string) *DetachVscMountPointRequest
	GetVscIds() []*string
}

type DetachVscMountPointRequest struct {
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
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 037cb49e1d-c***5
	MountPointId               *string   `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
	UseAssumeRoleChkServerPerm *bool     `json:"UseAssumeRoleChkServerPerm,omitempty" xml:"UseAssumeRoleChkServerPerm,omitempty"`
	VscIds                     []*string `json:"VscIds,omitempty" xml:"VscIds,omitempty" type:"Repeated"`
}

func (s DetachVscMountPointRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachVscMountPointRequest) GoString() string {
	return s.String()
}

func (s *DetachVscMountPointRequest) GetDescription() *string {
	return s.Description
}

func (s *DetachVscMountPointRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DetachVscMountPointRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *DetachVscMountPointRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DetachVscMountPointRequest) GetMountPointId() *string {
	return s.MountPointId
}

func (s *DetachVscMountPointRequest) GetUseAssumeRoleChkServerPerm() *bool {
	return s.UseAssumeRoleChkServerPerm
}

func (s *DetachVscMountPointRequest) GetVscIds() []*string {
	return s.VscIds
}

func (s *DetachVscMountPointRequest) SetDescription(v string) *DetachVscMountPointRequest {
	s.Description = &v
	return s
}

func (s *DetachVscMountPointRequest) SetFileSystemId(v string) *DetachVscMountPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *DetachVscMountPointRequest) SetInputRegionId(v string) *DetachVscMountPointRequest {
	s.InputRegionId = &v
	return s
}

func (s *DetachVscMountPointRequest) SetInstanceIds(v []*string) *DetachVscMountPointRequest {
	s.InstanceIds = v
	return s
}

func (s *DetachVscMountPointRequest) SetMountPointId(v string) *DetachVscMountPointRequest {
	s.MountPointId = &v
	return s
}

func (s *DetachVscMountPointRequest) SetUseAssumeRoleChkServerPerm(v bool) *DetachVscMountPointRequest {
	s.UseAssumeRoleChkServerPerm = &v
	return s
}

func (s *DetachVscMountPointRequest) SetVscIds(v []*string) *DetachVscMountPointRequest {
	s.VscIds = v
	return s
}

func (s *DetachVscMountPointRequest) Validate() error {
	return dara.Validate(s)
}
