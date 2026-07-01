// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIAgentInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeAIAgentInstanceRequest
	GetInstanceId() *string
}

type DescribeAIAgentInstanceRequest struct {
	// The ID of the AI agent instance.
	//
	// > The `InstanceId` is the unique ID returned when an AI agent instance starts successfully. For APIs that start an AI agent, see [StartAIAgentInstance](https://help.aliyun.com/document_detail/2846201.html) and [GenerateAIAgentCall](https://help.aliyun.com/document_detail/2846209.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeAIAgentInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIAgentInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAIAgentInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAIAgentInstanceRequest) SetInstanceId(v string) *DescribeAIAgentInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAIAgentInstanceRequest) Validate() error {
	return dara.Validate(s)
}
