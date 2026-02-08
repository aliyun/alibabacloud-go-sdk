// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDialogueFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDialogueFlowId(v string) *DeleteDialogueFlowRequest
	GetDialogueFlowId() *string
	SetInstanceId(v string) *DeleteDialogueFlowRequest
	GetInstanceId() *string
	SetScriptId(v string) *DeleteDialogueFlowRequest
	GetScriptId() *string
}

type DeleteDialogueFlowRequest struct {
	// Conversation flow ID
	//
	// This parameter is required.
	//
	// example:
	//
	// caab25d1-1f30-4996-8135-0036f5661b43
	DialogueFlowId *string `json:"DialogueFlowId,omitempty" xml:"DialogueFlowId,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 361c8a53-0e29-42f3-8aa7-c7752d010399
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Scenario ID
	//
	// This parameter is required.
	//
	// example:
	//
	// d7c28efb-47f7-4a85-a522-5038e30a0b98
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s DeleteDialogueFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDialogueFlowRequest) GoString() string {
	return s.String()
}

func (s *DeleteDialogueFlowRequest) GetDialogueFlowId() *string {
	return s.DialogueFlowId
}

func (s *DeleteDialogueFlowRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDialogueFlowRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DeleteDialogueFlowRequest) SetDialogueFlowId(v string) *DeleteDialogueFlowRequest {
	s.DialogueFlowId = &v
	return s
}

func (s *DeleteDialogueFlowRequest) SetInstanceId(v string) *DeleteDialogueFlowRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDialogueFlowRequest) SetScriptId(v string) *DeleteDialogueFlowRequest {
	s.ScriptId = &v
	return s
}

func (s *DeleteDialogueFlowRequest) Validate() error {
	return dara.Validate(s)
}
