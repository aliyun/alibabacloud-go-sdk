// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachVscMountPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *AttachVscMountPointRequest
	GetDescription() *string
	SetFileSystemId(v string) *AttachVscMountPointRequest
	GetFileSystemId() *string
	SetInputRegionId(v string) *AttachVscMountPointRequest
	GetInputRegionId() *string
	SetInstanceIds(v []*string) *AttachVscMountPointRequest
	GetInstanceIds() []*string
	SetMountPointId(v string) *AttachVscMountPointRequest
	GetMountPointId() *string
	SetUseAssumeRoleChkServerPerm(v bool) *AttachVscMountPointRequest
	GetUseAssumeRoleChkServerPerm() *bool
	SetVscIds(v []*string) *AttachVscMountPointRequest
	GetVscIds() []*string
	SetVscName(v string) *AttachVscMountPointRequest
	GetVscName() *string
	SetVscType(v string) *AttachVscMountPointRequest
	GetVscType() *string
}

type AttachVscMountPointRequest struct {
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
	VscName                    *string   `json:"VscName,omitempty" xml:"VscName,omitempty"`
	VscType                    *string   `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s AttachVscMountPointRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachVscMountPointRequest) GoString() string {
	return s.String()
}

func (s *AttachVscMountPointRequest) GetDescription() *string {
	return s.Description
}

func (s *AttachVscMountPointRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *AttachVscMountPointRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *AttachVscMountPointRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *AttachVscMountPointRequest) GetMountPointId() *string {
	return s.MountPointId
}

func (s *AttachVscMountPointRequest) GetUseAssumeRoleChkServerPerm() *bool {
	return s.UseAssumeRoleChkServerPerm
}

func (s *AttachVscMountPointRequest) GetVscIds() []*string {
	return s.VscIds
}

func (s *AttachVscMountPointRequest) GetVscName() *string {
	return s.VscName
}

func (s *AttachVscMountPointRequest) GetVscType() *string {
	return s.VscType
}

func (s *AttachVscMountPointRequest) SetDescription(v string) *AttachVscMountPointRequest {
	s.Description = &v
	return s
}

func (s *AttachVscMountPointRequest) SetFileSystemId(v string) *AttachVscMountPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *AttachVscMountPointRequest) SetInputRegionId(v string) *AttachVscMountPointRequest {
	s.InputRegionId = &v
	return s
}

func (s *AttachVscMountPointRequest) SetInstanceIds(v []*string) *AttachVscMountPointRequest {
	s.InstanceIds = v
	return s
}

func (s *AttachVscMountPointRequest) SetMountPointId(v string) *AttachVscMountPointRequest {
	s.MountPointId = &v
	return s
}

func (s *AttachVscMountPointRequest) SetUseAssumeRoleChkServerPerm(v bool) *AttachVscMountPointRequest {
	s.UseAssumeRoleChkServerPerm = &v
	return s
}

func (s *AttachVscMountPointRequest) SetVscIds(v []*string) *AttachVscMountPointRequest {
	s.VscIds = v
	return s
}

func (s *AttachVscMountPointRequest) SetVscName(v string) *AttachVscMountPointRequest {
	s.VscName = &v
	return s
}

func (s *AttachVscMountPointRequest) SetVscType(v string) *AttachVscMountPointRequest {
	s.VscType = &v
	return s
}

func (s *AttachVscMountPointRequest) Validate() error {
	return dara.Validate(s)
}
