// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMountPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupId(v string) *CreateMountPointRequest
	GetAccessGroupId() *string
	SetDescription(v string) *CreateMountPointRequest
	GetDescription() *string
	SetFileSystemId(v string) *CreateMountPointRequest
	GetFileSystemId() *string
	SetInputRegionId(v string) *CreateMountPointRequest
	GetInputRegionId() *string
	SetNetworkType(v string) *CreateMountPointRequest
	GetNetworkType() *string
	SetUsePerformanceMode(v bool) *CreateMountPointRequest
	GetUsePerformanceMode() *bool
	SetVSwitchId(v string) *CreateMountPointRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateMountPointRequest
	GetVpcId() *string
}

type CreateMountPointRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// acg-e3755fb0-358d-4286-9942-8d461048****
	AccessGroupId *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
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
	// VPC
	NetworkType        *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	UsePerformanceMode *bool   `json:"UsePerformanceMode,omitempty" xml:"UsePerformanceMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-iq8fymi327krd14mt****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-iq8hhsk3ymzv9m4wn****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateMountPointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMountPointRequest) GoString() string {
	return s.String()
}

func (s *CreateMountPointRequest) GetAccessGroupId() *string {
	return s.AccessGroupId
}

func (s *CreateMountPointRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateMountPointRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateMountPointRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *CreateMountPointRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateMountPointRequest) GetUsePerformanceMode() *bool {
	return s.UsePerformanceMode
}

func (s *CreateMountPointRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateMountPointRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateMountPointRequest) SetAccessGroupId(v string) *CreateMountPointRequest {
	s.AccessGroupId = &v
	return s
}

func (s *CreateMountPointRequest) SetDescription(v string) *CreateMountPointRequest {
	s.Description = &v
	return s
}

func (s *CreateMountPointRequest) SetFileSystemId(v string) *CreateMountPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateMountPointRequest) SetInputRegionId(v string) *CreateMountPointRequest {
	s.InputRegionId = &v
	return s
}

func (s *CreateMountPointRequest) SetNetworkType(v string) *CreateMountPointRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateMountPointRequest) SetUsePerformanceMode(v bool) *CreateMountPointRequest {
	s.UsePerformanceMode = &v
	return s
}

func (s *CreateMountPointRequest) SetVSwitchId(v string) *CreateMountPointRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateMountPointRequest) SetVpcId(v string) *CreateMountPointRequest {
	s.VpcId = &v
	return s
}

func (s *CreateMountPointRequest) Validate() error {
	return dara.Validate(s)
}
