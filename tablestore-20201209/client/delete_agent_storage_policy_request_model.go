// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentStoragePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentStorageName(v string) *DeleteAgentStoragePolicyRequest
	GetAgentStorageName() *string
	SetPolicyVersion(v int64) *DeleteAgentStoragePolicyRequest
	GetPolicyVersion() *int64
}

type DeleteAgentStoragePolicyRequest struct {
	// The name of the agent storage.
	//
	// This parameter is required.
	//
	// example:
	//
	// first-agent
	AgentStorageName *string `json:"AgentStorageName,omitempty" xml:"AgentStorageName,omitempty"`
	// The version of the access control policy for agent storage.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	PolicyVersion *int64 `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
}

func (s DeleteAgentStoragePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentStoragePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentStoragePolicyRequest) GetAgentStorageName() *string {
	return s.AgentStorageName
}

func (s *DeleteAgentStoragePolicyRequest) GetPolicyVersion() *int64 {
	return s.PolicyVersion
}

func (s *DeleteAgentStoragePolicyRequest) SetAgentStorageName(v string) *DeleteAgentStoragePolicyRequest {
	s.AgentStorageName = &v
	return s
}

func (s *DeleteAgentStoragePolicyRequest) SetPolicyVersion(v int64) *DeleteAgentStoragePolicyRequest {
	s.PolicyVersion = &v
	return s
}

func (s *DeleteAgentStoragePolicyRequest) Validate() error {
	return dara.Validate(s)
}
