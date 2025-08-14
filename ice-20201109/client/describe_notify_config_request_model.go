// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNotifyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIAgentId(v string) *DescribeNotifyConfigRequest
	GetAIAgentId() *string
}

type DescribeNotifyConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	AIAgentId *string `json:"AIAgentId,omitempty" xml:"AIAgentId,omitempty"`
}

func (s DescribeNotifyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeNotifyConfigRequest) GetAIAgentId() *string {
	return s.AIAgentId
}

func (s *DescribeNotifyConfigRequest) SetAIAgentId(v string) *DescribeNotifyConfigRequest {
	s.AIAgentId = &v
	return s
}

func (s *DescribeNotifyConfigRequest) Validate() error {
	return dara.Validate(s)
}
