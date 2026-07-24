// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartComputeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartComputeInstanceRequest
	GetClientToken() *string
	SetInstanceId(v string) *StartComputeInstanceRequest
	GetInstanceId() *string
	SetInstanceName(v string) *StartComputeInstanceRequest
	GetInstanceName() *string
	SetRegionId(v string) *StartComputeInstanceRequest
	GetRegionId() *string
	SetSelectedZones(v string) *StartComputeInstanceRequest
	GetSelectedZones() *string
	SetServiceVersion(v string) *StartComputeInstanceRequest
	GetServiceVersion() *string
	SetVSwitchIds(v []*string) *StartComputeInstanceRequest
	GetVSwitchIds() []*string
	SetVpcId(v string) *StartComputeInstanceRequest
	GetVpcId() *string
}

type StartComputeInstanceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SelectedZones  *string `json:"SelectedZones,omitempty" xml:"SelectedZones,omitempty"`
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// This parameter is required.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// This parameter is required.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s StartComputeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartComputeInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartComputeInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartComputeInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartComputeInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *StartComputeInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartComputeInstanceRequest) GetSelectedZones() *string {
	return s.SelectedZones
}

func (s *StartComputeInstanceRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *StartComputeInstanceRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *StartComputeInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *StartComputeInstanceRequest) SetClientToken(v string) *StartComputeInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StartComputeInstanceRequest) SetInstanceId(v string) *StartComputeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StartComputeInstanceRequest) SetInstanceName(v string) *StartComputeInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *StartComputeInstanceRequest) SetRegionId(v string) *StartComputeInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartComputeInstanceRequest) SetSelectedZones(v string) *StartComputeInstanceRequest {
	s.SelectedZones = &v
	return s
}

func (s *StartComputeInstanceRequest) SetServiceVersion(v string) *StartComputeInstanceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *StartComputeInstanceRequest) SetVSwitchIds(v []*string) *StartComputeInstanceRequest {
	s.VSwitchIds = v
	return s
}

func (s *StartComputeInstanceRequest) SetVpcId(v string) *StartComputeInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *StartComputeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
