// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAgentStorage2VpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentStorageName(v string) *BindAgentStorage2VpcRequest
	GetAgentStorageName() *string
	SetAgentStorageVpcName(v string) *BindAgentStorage2VpcRequest
	GetAgentStorageVpcName() *string
	SetVirtualSwitchId(v string) *BindAgentStorage2VpcRequest
	GetVirtualSwitchId() *string
	SetVpcId(v string) *BindAgentStorage2VpcRequest
	GetVpcId() *string
}

type BindAgentStorage2VpcRequest struct {
	// The agent storage name.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent-test
	AgentStorageName *string `json:"AgentStorageName,omitempty" xml:"AgentStorageName,omitempty"`
	// The VPC name.
	//
	// This parameter is required.
	//
	// example:
	//
	// remua
	AgentStorageVpcName *string `json:"AgentStorageVpcName,omitempty" xml:"AgentStorageVpcName,omitempty"`
	// The vSwitch ID.
	//
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

func (s BindAgentStorage2VpcRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAgentStorage2VpcRequest) GoString() string {
	return s.String()
}

func (s *BindAgentStorage2VpcRequest) GetAgentStorageName() *string {
	return s.AgentStorageName
}

func (s *BindAgentStorage2VpcRequest) GetAgentStorageVpcName() *string {
	return s.AgentStorageVpcName
}

func (s *BindAgentStorage2VpcRequest) GetVirtualSwitchId() *string {
	return s.VirtualSwitchId
}

func (s *BindAgentStorage2VpcRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *BindAgentStorage2VpcRequest) SetAgentStorageName(v string) *BindAgentStorage2VpcRequest {
	s.AgentStorageName = &v
	return s
}

func (s *BindAgentStorage2VpcRequest) SetAgentStorageVpcName(v string) *BindAgentStorage2VpcRequest {
	s.AgentStorageVpcName = &v
	return s
}

func (s *BindAgentStorage2VpcRequest) SetVirtualSwitchId(v string) *BindAgentStorage2VpcRequest {
	s.VirtualSwitchId = &v
	return s
}

func (s *BindAgentStorage2VpcRequest) SetVpcId(v string) *BindAgentStorage2VpcRequest {
	s.VpcId = &v
	return s
}

func (s *BindAgentStorage2VpcRequest) Validate() error {
	return dara.Validate(s)
}
