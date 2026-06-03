// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDialogueFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDialogueFlowDefinition(v string) *ModifyDialogueFlowRequest
	GetDialogueFlowDefinition() *string
	SetDialogueFlowId(v string) *ModifyDialogueFlowRequest
	GetDialogueFlowId() *string
	SetInstanceId(v string) *ModifyDialogueFlowRequest
	GetInstanceId() *string
	SetIsDrafted(v bool) *ModifyDialogueFlowRequest
	GetIsDrafted() *bool
	SetScriptId(v string) *ModifyDialogueFlowRequest
	GetScriptId() *string
}

type ModifyDialogueFlowRequest struct {
	// This parameter is required.
	DialogueFlowDefinition *string `json:"DialogueFlowDefinition,omitempty" xml:"DialogueFlowDefinition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 390515b5-6115-4ccf-83e2-52d5bfaf2ddf
	DialogueFlowId *string `json:"DialogueFlowId,omitempty" xml:"DialogueFlowId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	IsDrafted *bool `json:"IsDrafted,omitempty" xml:"IsDrafted,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// b0f35dd1-0337-402e-9c4f-3a6c2426950a
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s ModifyDialogueFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDialogueFlowRequest) GoString() string {
	return s.String()
}

func (s *ModifyDialogueFlowRequest) GetDialogueFlowDefinition() *string {
	return s.DialogueFlowDefinition
}

func (s *ModifyDialogueFlowRequest) GetDialogueFlowId() *string {
	return s.DialogueFlowId
}

func (s *ModifyDialogueFlowRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDialogueFlowRequest) GetIsDrafted() *bool {
	return s.IsDrafted
}

func (s *ModifyDialogueFlowRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyDialogueFlowRequest) SetDialogueFlowDefinition(v string) *ModifyDialogueFlowRequest {
	s.DialogueFlowDefinition = &v
	return s
}

func (s *ModifyDialogueFlowRequest) SetDialogueFlowId(v string) *ModifyDialogueFlowRequest {
	s.DialogueFlowId = &v
	return s
}

func (s *ModifyDialogueFlowRequest) SetInstanceId(v string) *ModifyDialogueFlowRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDialogueFlowRequest) SetIsDrafted(v bool) *ModifyDialogueFlowRequest {
	s.IsDrafted = &v
	return s
}

func (s *ModifyDialogueFlowRequest) SetScriptId(v string) *ModifyDialogueFlowRequest {
	s.ScriptId = &v
	return s
}

func (s *ModifyDialogueFlowRequest) Validate() error {
	return dara.Validate(s)
}
