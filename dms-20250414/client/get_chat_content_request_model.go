// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *GetChatContentRequest
	GetAgentId() *string
	SetCheckpoint(v string) *GetChatContentRequest
	GetCheckpoint() *string
	SetDMSUnit(v string) *GetChatContentRequest
	GetDMSUnit() *string
	SetSessionId(v string) *GetChatContentRequest
	GetSessionId() *string
}

type GetChatContentRequest struct {
	// example:
	//
	// xxxx-xxxx-xxxx
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 0
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// example:
	//
	// sess_12345
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetChatContentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChatContentRequest) GoString() string {
	return s.String()
}

func (s *GetChatContentRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *GetChatContentRequest) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *GetChatContentRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *GetChatContentRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetChatContentRequest) SetAgentId(v string) *GetChatContentRequest {
	s.AgentId = &v
	return s
}

func (s *GetChatContentRequest) SetCheckpoint(v string) *GetChatContentRequest {
	s.Checkpoint = &v
	return s
}

func (s *GetChatContentRequest) SetDMSUnit(v string) *GetChatContentRequest {
	s.DMSUnit = &v
	return s
}

func (s *GetChatContentRequest) SetSessionId(v string) *GetChatContentRequest {
	s.SessionId = &v
	return s
}

func (s *GetChatContentRequest) Validate() error {
	return dara.Validate(s)
}
