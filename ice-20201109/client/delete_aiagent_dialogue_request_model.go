// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIAgentDialogueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDialogueId(v string) *DeleteAIAgentDialogueRequest
	GetDialogueId() *string
	SetNodeId(v string) *DeleteAIAgentDialogueRequest
	GetNodeId() *string
	SetSessionId(v string) *DeleteAIAgentDialogueRequest
	GetSessionId() *string
}

type DeleteAIAgentDialogueRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// f27f9b9be28642a88e18*******
	DialogueId *string `json:"DialogueId,omitempty" xml:"DialogueId,omitempty"`
	NodeId     *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6d594e7f55624c47a48789******
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DeleteAIAgentDialogueRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIAgentDialogueRequest) GoString() string {
	return s.String()
}

func (s *DeleteAIAgentDialogueRequest) GetDialogueId() *string {
	return s.DialogueId
}

func (s *DeleteAIAgentDialogueRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DeleteAIAgentDialogueRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DeleteAIAgentDialogueRequest) SetDialogueId(v string) *DeleteAIAgentDialogueRequest {
	s.DialogueId = &v
	return s
}

func (s *DeleteAIAgentDialogueRequest) SetNodeId(v string) *DeleteAIAgentDialogueRequest {
	s.NodeId = &v
	return s
}

func (s *DeleteAIAgentDialogueRequest) SetSessionId(v string) *DeleteAIAgentDialogueRequest {
	s.SessionId = &v
	return s
}

func (s *DeleteAIAgentDialogueRequest) Validate() error {
	return dara.Validate(s)
}
