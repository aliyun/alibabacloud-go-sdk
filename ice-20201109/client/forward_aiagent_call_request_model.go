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
	// example:
	//
	// 13**********
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CallerNumber *string `json:"CallerNumber,omitempty" xml:"CallerNumber,omitempty"`
	ErrorPrompt  *string `json:"ErrorPrompt,omitempty" xml:"ErrorPrompt,omitempty"`
	// example:
	//
	// call_instance_202******
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
