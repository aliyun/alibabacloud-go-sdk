// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartComputeInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StartComputeInstanceShrinkRequest
	GetInstanceId() *string
	SetInstanceName(v string) *StartComputeInstanceShrinkRequest
	GetInstanceName() *string
	SetRegionId(v string) *StartComputeInstanceShrinkRequest
	GetRegionId() *string
	SetVSwitchIdsShrink(v string) *StartComputeInstanceShrinkRequest
	GetVSwitchIdsShrink() *string
	SetVpcId(v string) *StartComputeInstanceShrinkRequest
	GetVpcId() *string
}

type StartComputeInstanceShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// alikafka_streaming-cn-pe333xxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// streaming-prod
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	VSwitchIdsShrink *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1abcdefg
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s StartComputeInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartComputeInstanceShrinkRequest) GoString() string {
	return s.String()
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

func (s *StartComputeInstanceShrinkRequest) GetVSwitchIdsShrink() *string {
	return s.VSwitchIdsShrink
}

func (s *StartComputeInstanceShrinkRequest) GetVpcId() *string {
	return s.VpcId
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
