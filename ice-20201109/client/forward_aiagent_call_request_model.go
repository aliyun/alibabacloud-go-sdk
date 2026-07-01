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
	SetCallerNumber(v string) *ForwardAIAgentCallRequest
	GetCallerNumber() *string
	SetErrorPrompt(v string) *ForwardAIAgentCallRequest
	GetErrorPrompt() *string
	SetInstanceId(v string) *ForwardAIAgentCallRequest
	GetInstanceId() *string
	SetTransferPrompt(v string) *ForwardAIAgentCallRequest
	GetTransferPrompt() *string
}

type ForwardAIAgentCallRequest struct {
	// The target phone number for call transfer.
	//
	// example:
	//
	// 13**********
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The caller phone number for the transferred call. Optional.
	//
	// 	Notice:
	//
	// By default, the CallerNumber is the agent\\"s phone number after the call starts:
	//
	// 1\\. For inbound lines, the agent number is the agent’s seat number.
	//
	// 2\\. For outbound lines, the agent number is the original caller number.
	//
	//
	//
	//
	// 	Warning:
	//
	// Alibaba Cloud lines do not support this parameter.
	//
	// example:
	//
	// 13**********
	CallerNumber *string `json:"CallerNumber,omitempty" xml:"CallerNumber,omitempty"`
	// Abnormal prompt text played when the transfer fails. Default is empty.
	//
	// example:
	//
	// We’re sorry, we’re unable to transfer your call at the moment. Please try again later.
	ErrorPrompt *string `json:"ErrorPrompt,omitempty" xml:"ErrorPrompt,omitempty"`
	// Current call instance ID, used only in inbound call transfer scenarios.
	//
	// example:
	//
	// call_instance_202******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Prompt message played before initiating the transfer. If empty, the system skips the prompt and plays the ringing tone directly. Default is empty.
	//
	// example:
	//
	// Please hold while I transfer your call.
	TransferPrompt *string `json:"TransferPrompt,omitempty" xml:"TransferPrompt,omitempty"`
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

func (s *ForwardAIAgentCallRequest) GetCallerNumber() *string {
	return s.CallerNumber
}

func (s *ForwardAIAgentCallRequest) GetErrorPrompt() *string {
	return s.ErrorPrompt
}

func (s *ForwardAIAgentCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ForwardAIAgentCallRequest) GetTransferPrompt() *string {
	return s.TransferPrompt
}

func (s *ForwardAIAgentCallRequest) SetCalledNumber(v string) *ForwardAIAgentCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *ForwardAIAgentCallRequest) SetCallerNumber(v string) *ForwardAIAgentCallRequest {
	s.CallerNumber = &v
	return s
}

func (s *ForwardAIAgentCallRequest) SetErrorPrompt(v string) *ForwardAIAgentCallRequest {
	s.ErrorPrompt = &v
	return s
}

func (s *ForwardAIAgentCallRequest) SetInstanceId(v string) *ForwardAIAgentCallRequest {
	s.InstanceId = &v
	return s
}

func (s *ForwardAIAgentCallRequest) SetTransferPrompt(v string) *ForwardAIAgentCallRequest {
	s.TransferPrompt = &v
	return s
}

func (s *ForwardAIAgentCallRequest) Validate() error {
	return dara.Validate(s)
}
