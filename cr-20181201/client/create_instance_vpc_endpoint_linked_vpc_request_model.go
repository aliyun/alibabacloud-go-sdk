// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceVpcEndpointLinkedVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableCreateDNSRecordInPvzt(v bool) *CreateInstanceVpcEndpointLinkedVpcRequest
	GetEnableCreateDNSRecordInPvzt() *bool
	SetInstanceId(v string) *CreateInstanceVpcEndpointLinkedVpcRequest
	GetInstanceId() *string
	SetModuleName(v string) *CreateInstanceVpcEndpointLinkedVpcRequest
	GetModuleName() *string
	SetVpcId(v string) *CreateInstanceVpcEndpointLinkedVpcRequest
	GetVpcId() *string
	SetVswitchId(v string) *CreateInstanceVpcEndpointLinkedVpcRequest
	GetVswitchId() *string
}

type CreateInstanceVpcEndpointLinkedVpcRequest struct {
	// Specifies whether to automatically create Alibaba Cloud DNS PrivateZone records. Valid values:
	//
	// >  If you enable automatic creation of PrivateZone records, a PrivateZone record is automatically created when you associate a VPC with the instance.
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	EnableCreateDNSRecordInPvzt *bool `json:"EnableCreateDNSRecordInPvzt,omitempty" xml:"EnableCreateDNSRecordInPvzt,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// 	- `Registry`: image repositories.
	//
	// 	- `Chart`: Helm charts.
	//
	// example:
	//
	// Registry
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The VPC ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-uf6pa68zxnnlc48dd****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch that is associated with the specified VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf6u0kn8x2gbzxfn2****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s CreateInstanceVpcEndpointLinkedVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceVpcEndpointLinkedVpcRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceVpcEndpointLinkedVpcRequest) GetEnableCreateDNSRecordInPvzt() *bool {
	return s.EnableCreateDNSRecordInPvzt
}

func (s *CreateInstanceVpcEndpointLinkedVpcRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceVpcEndpointLinkedVpcRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *CreateInstanceVpcEndpointLinkedVpcRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateInstanceVpcEndpointLinkedVpcRequest) GetVswitchId() *string {
	return s.VswitchId
}

func (s *CreateInstanceVpcEndpointLinkedVpcRequest) SetEnableCreateDNSRecordInPvzt(v bool) *CreateInstanceVpcEndpointLinkedVpcRequest {
	s.EnableCreateDNSRecordInPvzt = &v
	return s
}

func (s *CreateInstanceVpcEndpointLinkedVpcRequest) SetInstanceId(v string) *CreateInstanceVpcEndpointLinkedVpcRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceVpcEndpointLinkedVpcRequest) SetModuleName(v string) *CreateInstanceVpcEndpointLinkedVpcRequest {
	s.ModuleName = &v
	return s
}

func (s *CreateInstanceVpcEndpointLinkedVpcRequest) SetVpcId(v string) *CreateInstanceVpcEndpointLinkedVpcRequest {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceVpcEndpointLinkedVpcRequest) SetVswitchId(v string) *CreateInstanceVpcEndpointLinkedVpcRequest {
	s.VswitchId = &v
	return s
}

func (s *CreateInstanceVpcEndpointLinkedVpcRequest) Validate() error {
	return dara.Validate(s)
}
