// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartComputeInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartComputeInstanceShrinkRequest
	GetClientToken() *string
	SetInstanceId(v string) *StartComputeInstanceShrinkRequest
	GetInstanceId() *string
	SetInstanceName(v string) *StartComputeInstanceShrinkRequest
	GetInstanceName() *string
	SetRegionId(v string) *StartComputeInstanceShrinkRequest
	GetRegionId() *string
	SetSelectedZones(v string) *StartComputeInstanceShrinkRequest
	GetSelectedZones() *string
	SetServiceVersion(v string) *StartComputeInstanceShrinkRequest
	GetServiceVersion() *string
	SetVSwitchIdsShrink(v string) *StartComputeInstanceShrinkRequest
	GetVSwitchIdsShrink() *string
	SetVpcId(v string) *StartComputeInstanceShrinkRequest
	GetVpcId() *string
}

type StartComputeInstanceShrinkRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SelectedZones  *string `json:"SelectedZones,omitempty" xml:"SelectedZones,omitempty"`
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// This parameter is required.
	VSwitchIdsShrink *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// This parameter is required.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s StartComputeInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartComputeInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartComputeInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartComputeInstanceShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartComputeInstanceShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *StartComputeInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartComputeInstanceShrinkRequest) GetSelectedZones() *string {
	return s.SelectedZones
}

func (s *StartComputeInstanceShrinkRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *StartComputeInstanceShrinkRequest) GetVSwitchIdsShrink() *string {
	return s.VSwitchIdsShrink
}

func (s *StartComputeInstanceShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *StartComputeInstanceShrinkRequest) SetClientToken(v string) *StartComputeInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *StartComputeInstanceShrinkRequest) SetInstanceId(v string) *StartComputeInstanceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *StartComputeInstanceShrinkRequest) SetInstanceName(v string) *StartComputeInstanceShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *StartComputeInstanceShrinkRequest) SetRegionId(v string) *StartComputeInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *StartComputeInstanceShrinkRequest) SetSelectedZones(v string) *StartComputeInstanceShrinkRequest {
	s.SelectedZones = &v
	return s
}

func (s *StartComputeInstanceShrinkRequest) SetServiceVersion(v string) *StartComputeInstanceShrinkRequest {
	s.ServiceVersion = &v
	return s
}

func (s *StartComputeInstanceShrinkRequest) SetVSwitchIdsShrink(v string) *StartComputeInstanceShrinkRequest {
	s.VSwitchIdsShrink = &v
	return s
}

func (s *StartComputeInstanceShrinkRequest) SetVpcId(v string) *StartComputeInstanceShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *StartComputeInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
