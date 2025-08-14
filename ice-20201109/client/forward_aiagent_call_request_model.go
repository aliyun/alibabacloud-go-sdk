// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForwardAIAgentCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalledNumber(v string) *ForwardAIAgentCallRequest
	GetCalledNumber() *string
	SetInstanceId(v string) *ForwardAIAgentCallRequest
	GetInstanceId() *string
}

type ForwardAIAgentCallRequest struct {
	// example:
	//
	// 13**********
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// example:
	//
	// call_instance_202******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ForwardAIAgentCallRequest) String() string {
	return dara.Prettify(s)
}

func (s ForwardAIAgentCallRequest) GoString() string {
	return s.String()
}

func (s *ForwardAIAgentCallRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *ForwardAIAgentCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ForwardAIAgentCallRequest) SetCalledNumber(v string) *ForwardAIAgentCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *ForwardAIAgentCallRequest) SetInstanceId(v string) *ForwardAIAgentCallRequest {
	s.InstanceId = &v
	return s
}

func (s *ForwardAIAgentCallRequest) Validate() error {
	return dara.Validate(s)
}
