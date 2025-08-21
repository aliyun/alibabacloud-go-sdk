// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DescribeInstanceRequest
	GetAgentKey() *string
	SetInstanceId(v string) *DescribeInstanceRequest
	GetInstanceId() *string
}

type DescribeInstanceRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DescribeInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceRequest) SetAgentKey(v string) *DescribeInstanceRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeInstanceRequest) SetInstanceId(v string) *DescribeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
