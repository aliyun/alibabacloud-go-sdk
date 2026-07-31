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
	Confirmed *bool                      `json:"confirmed,omitempty" xml:"confirmed,omitempty"`
	Phase     *string                    `json:"phase,omitempty" xml:"phase,omitempty"`
	Reason    *string                    `json:"reason,omitempty" xml:"reason,omitempty"`
	SessionId *string                    `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
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
	Id            *string                `json:"id,omitempty" xml:"id,omitempty"`
	ModifiedInput map[string]interface{} `json:"modifiedInput,omitempty" xml:"modifiedInput,omitempty"`
	Name          *string                `json:"name,omitempty" xml:"name,omitempty"`
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
