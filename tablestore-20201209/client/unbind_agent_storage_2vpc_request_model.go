// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAgentStorage2VpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentStorageName(v string) *UnbindAgentStorage2VpcRequest
	GetAgentStorageName() *string
	SetAgentStorageVpcName(v string) *UnbindAgentStorage2VpcRequest
	GetAgentStorageVpcName() *string
}

type UnbindAgentStorage2VpcRequest struct {
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
}

func (s UnbindAgentStorage2VpcRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindAgentStorage2VpcRequest) GoString() string {
	return s.String()
}

func (s *UnbindAgentStorage2VpcRequest) GetAgentStorageName() *string {
	return s.AgentStorageName
}

func (s *UnbindAgentStorage2VpcRequest) GetAgentStorageVpcName() *string {
	return s.AgentStorageVpcName
}

func (s *UnbindAgentStorage2VpcRequest) SetAgentStorageName(v string) *UnbindAgentStorage2VpcRequest {
	s.AgentStorageName = &v
	return s
}

func (s *UnbindAgentStorage2VpcRequest) SetAgentStorageVpcName(v string) *UnbindAgentStorage2VpcRequest {
	s.AgentStorageVpcName = &v
	return s
}

func (s *UnbindAgentStorage2VpcRequest) Validate() error {
	return dara.Validate(s)
}
