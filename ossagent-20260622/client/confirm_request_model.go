// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfirmed(v bool) *ConfirmRequest
	GetConfirmed() *bool
	SetPhase(v string) *ConfirmRequest
	GetPhase() *string
	SetReason(v string) *ConfirmRequest
	GetReason() *string
	SetSessionId(v string) *ConfirmRequest
	GetSessionId() *string
	SetToolCalls(v []*ConfirmRequestToolCalls) *ConfirmRequest
	GetToolCalls() []*ConfirmRequestToolCalls
}

type ConfirmRequest struct {
	// Specifies whether to approve the tool execution.
	//
	// example:
	//
	// true
	Confirmed *bool `json:"confirmed,omitempty" xml:"confirmed,omitempty"`
	// The current execution phase.
	//
	// example:
	//
	// PARAM_INPUT
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
	// The reason for whether to call the tool.
	//
	// example:
	//
	// null
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// The Q&A session ID.
	//
	// example:
	//
	// UUID
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The tool invocations.
	ToolCalls []*ConfirmRequestToolCalls `json:"toolCalls,omitempty" xml:"toolCalls,omitempty" type:"Repeated"`
}

func (s ConfirmRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfirmRequest) GoString() string {
	return s.String()
}

func (s *ConfirmRequest) GetConfirmed() *bool {
	return s.Confirmed
}

func (s *ConfirmRequest) GetPhase() *string {
	return s.Phase
}

func (s *ConfirmRequest) GetReason() *string {
	return s.Reason
}

func (s *ConfirmRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ConfirmRequest) GetToolCalls() []*ConfirmRequestToolCalls {
	return s.ToolCalls
}

func (s *ConfirmRequest) SetConfirmed(v bool) *ConfirmRequest {
	s.Confirmed = &v
	return s
}

func (s *ConfirmRequest) SetPhase(v string) *ConfirmRequest {
	s.Phase = &v
	return s
}

func (s *ConfirmRequest) SetReason(v string) *ConfirmRequest {
	s.Reason = &v
	return s
}

func (s *ConfirmRequest) SetSessionId(v string) *ConfirmRequest {
	s.SessionId = &v
	return s
}

func (s *ConfirmRequest) SetToolCalls(v []*ConfirmRequestToolCalls) *ConfirmRequest {
	s.ToolCalls = v
	return s
}

func (s *ConfirmRequest) Validate() error {
	if s.ToolCalls != nil {
		for _, item := range s.ToolCalls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ConfirmRequestToolCalls struct {
	// The tool ID, returned by the Chat operation.
	//
	// example:
	//
	// call_662cc029b3444d8d923a7ea6
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The command to execute for the tool calling operation, returned by the Chat operation.
	//
	// example:
	//
	// {
	//
	//     "command": "api put-bucket-acl --bucket xxx --acl private",
	//
	//     "region": "cn-hangzhou"
	//
	// }
	ModifiedInput map[string]interface{} `json:"modifiedInput,omitempty" xml:"modifiedInput,omitempty"`
	// The consumer name.
	//
	// example:
	//
	// ossutil_safe
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ConfirmRequestToolCalls) String() string {
	return dara.Prettify(s)
}

func (s ConfirmRequestToolCalls) GoString() string {
	return s.String()
}

func (s *ConfirmRequestToolCalls) GetId() *string {
	return s.Id
}

func (s *ConfirmRequestToolCalls) GetModifiedInput() map[string]interface{} {
	return s.ModifiedInput
}

func (s *ConfirmRequestToolCalls) GetName() *string {
	return s.Name
}

func (s *ConfirmRequestToolCalls) SetId(v string) *ConfirmRequestToolCalls {
	s.Id = &v
	return s
}

func (s *ConfirmRequestToolCalls) SetModifiedInput(v map[string]interface{}) *ConfirmRequestToolCalls {
	s.ModifiedInput = v
	return s
}

func (s *ConfirmRequestToolCalls) SetName(v string) *ConfirmRequestToolCalls {
	s.Name = &v
	return s
}

func (s *ConfirmRequestToolCalls) Validate() error {
	return dara.Validate(s)
}
