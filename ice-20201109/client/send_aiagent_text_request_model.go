// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAIAgentTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SendAIAgentTextRequest
	GetInstanceId() *string
	SetText(v string) *SendAIAgentTextRequest
	GetText() *string
}

type SendAIAgentTextRequest struct {
	// The ID of the AI agent instance.
	//
	// > InstanceId is a unique identifier that is returned when the AI agent is started. To start an AI agent, call [StartAIAgentInstance ](https://help.aliyun.com/document_detail/2846201.html) or [GenerateAIAgentCall](https://help.aliyun.com/document_detail/2846209.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// f27f9b9be28642a88e18****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The input to the LLM.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hello
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s SendAIAgentTextRequest) String() string {
	return dara.Prettify(s)
}

func (s SendAIAgentTextRequest) GoString() string {
	return s.String()
}

func (s *SendAIAgentTextRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SendAIAgentTextRequest) GetText() *string {
	return s.Text
}

func (s *SendAIAgentTextRequest) SetInstanceId(v string) *SendAIAgentTextRequest {
	s.InstanceId = &v
	return s
}

func (s *SendAIAgentTextRequest) SetText(v string) *SendAIAgentTextRequest {
	s.Text = &v
	return s
}

func (s *SendAIAgentTextRequest) Validate() error {
	return dara.Validate(s)
}
