// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceVpcEndpointLinkedVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteInstanceVpcEndpointLinkedVpcRequest
	GetInstanceId() *string
	SetModuleName(v string) *DeleteInstanceVpcEndpointLinkedVpcRequest
	GetModuleName() *string
	SetVpcId(v string) *DeleteInstanceVpcEndpointLinkedVpcRequest
	GetVpcId() *string
	SetVswitchId(v string) *DeleteInstanceVpcEndpointLinkedVpcRequest
	GetVswitchId() *string
}

type DeleteInstanceVpcEndpointLinkedVpcRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// 	- `Registry`: the image repository.
	//
	// 	- `Chart`: a Helm chart.
	//
	// example:
	//
	// Chart
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The ID of the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-uf6pa68zxnnlc48dd****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-uf6pa68zxnnlc48dd****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DeleteInstanceVpcEndpointLinkedVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceVpcEndpointLinkedVpcRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceVpcEndpointLinkedVpcRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteInstanceVpcEndpointLinkedVpcRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *DeleteInstanceVpcEndpointLinkedVpcRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DeleteInstanceVpcEndpointLinkedVpcRequest) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DeleteInstanceVpcEndpointLinkedVpcRequest) SetInstanceId(v string) *DeleteInstanceVpcEndpointLinkedVpcRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceVpcEndpointLinkedVpcRequest) SetModuleName(v string) *DeleteInstanceVpcEndpointLinkedVpcRequest {
	s.ModuleName = &v
	return s
}

func (s *DeleteInstanceVpcEndpointLinkedVpcRequest) SetVpcId(v string) *DeleteInstanceVpcEndpointLinkedVpcRequest {
	s.VpcId = &v
	return s
}

func (s *DeleteInstanceVpcEndpointLinkedVpcRequest) SetVswitchId(v string) *DeleteInstanceVpcEndpointLinkedVpcRequest {
	s.VswitchId = &v
	return s
}

func (s *DeleteInstanceVpcEndpointLinkedVpcRequest) Validate() error {
	return dara.Validate(s)
}
