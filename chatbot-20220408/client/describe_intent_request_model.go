// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIntentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DescribeIntentRequest
	GetAgentKey() *string
	SetInstanceId(v string) *DescribeIntentRequest
	GetInstanceId() *string
	SetIntentId(v int64) *DescribeIntentRequest
	GetIntentId() *int64
}

type DescribeIntentRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
}

func (s DescribeIntentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntentRequest) GoString() string {
	return s.String()
}

func (s *DescribeIntentRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DescribeIntentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeIntentRequest) GetIntentId() *int64 {
	return s.IntentId
}

func (s *DescribeIntentRequest) SetAgentKey(v string) *DescribeIntentRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeIntentRequest) SetInstanceId(v string) *DescribeIntentRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeIntentRequest) SetIntentId(v int64) *DescribeIntentRequest {
	s.IntentId = &v
	return s
}

func (s *DescribeIntentRequest) Validate() error {
	return dara.Validate(s)
}
