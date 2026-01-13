// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindInstance2VpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *BindInstance2VpcRequest
	GetInstanceName() *string
	SetInstanceVpcName(v string) *BindInstance2VpcRequest
	GetInstanceVpcName() *string
	SetVirtualSwitchId(v string) *BindInstance2VpcRequest
	GetVirtualSwitchId() *string
	SetVpcId(v string) *BindInstance2VpcRequest
	GetVpcId() *string
}

type BindInstance2VpcRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mkdd-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xu6666
	InstanceVpcName *string `json:"InstanceVpcName,omitempty" xml:"InstanceVpcName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf6***********ez6ge
	VirtualSwitchId *string `json:"VirtualSwitchId,omitempty" xml:"VirtualSwitchId,omitempty"`
	// VPC ID
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-2ze***********g31n7
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s BindInstance2VpcRequest) String() string {
	return dara.Prettify(s)
}

func (s BindInstance2VpcRequest) GoString() string {
	return s.String()
}

func (s *BindInstance2VpcRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *BindInstance2VpcRequest) GetInstanceVpcName() *string {
	return s.InstanceVpcName
}

func (s *BindInstance2VpcRequest) GetVirtualSwitchId() *string {
	return s.VirtualSwitchId
}

func (s *BindInstance2VpcRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *BindInstance2VpcRequest) SetInstanceName(v string) *BindInstance2VpcRequest {
	s.InstanceName = &v
	return s
}

func (s *BindInstance2VpcRequest) SetInstanceVpcName(v string) *BindInstance2VpcRequest {
	s.InstanceVpcName = &v
	return s
}

func (s *BindInstance2VpcRequest) SetVirtualSwitchId(v string) *BindInstance2VpcRequest {
	s.VirtualSwitchId = &v
	return s
}

func (s *BindInstance2VpcRequest) SetVpcId(v string) *BindInstance2VpcRequest {
	s.VpcId = &v
	return s
}

func (s *BindInstance2VpcRequest) Validate() error {
	return dara.Validate(s)
}
