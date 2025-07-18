// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAttackProtectionCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentType(v string) *DescribeAttackProtectionCountRequest
	GetAgentType() *string
	SetEndTimestamp(v int32) *DescribeAttackProtectionCountRequest
	GetEndTimestamp() *int32
	SetStartTimestamp(v int32) *DescribeAttackProtectionCountRequest
	GetStartTimestamp() *int32
}

type DescribeAttackProtectionCountRequest struct {
	// example:
	//
	// sas
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1767456000000
	EndTimestamp *int32 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1739289981765
	StartTimestamp *int32 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeAttackProtectionCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttackProtectionCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeAttackProtectionCountRequest) GetAgentType() *string {
	return s.AgentType
}

func (s *DescribeAttackProtectionCountRequest) GetEndTimestamp() *int32 {
	return s.EndTimestamp
}

func (s *DescribeAttackProtectionCountRequest) GetStartTimestamp() *int32 {
	return s.StartTimestamp
}

func (s *DescribeAttackProtectionCountRequest) SetAgentType(v string) *DescribeAttackProtectionCountRequest {
	s.AgentType = &v
	return s
}

func (s *DescribeAttackProtectionCountRequest) SetEndTimestamp(v int32) *DescribeAttackProtectionCountRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeAttackProtectionCountRequest) SetStartTimestamp(v int32) *DescribeAttackProtectionCountRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeAttackProtectionCountRequest) Validate() error {
	return dara.Validate(s)
}
