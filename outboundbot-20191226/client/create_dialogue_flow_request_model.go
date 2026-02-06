// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDialogueFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDialogueFlowType(v string) *CreateDialogueFlowRequest
	GetDialogueFlowType() *string
	SetDialogueName(v string) *CreateDialogueFlowRequest
	GetDialogueName() *string
	SetInstanceId(v string) *CreateDialogueFlowRequest
	GetInstanceId() *string
	SetScriptId(v string) *CreateDialogueFlowRequest
	GetScriptId() *string
}

type CreateDialogueFlowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// MainFlow
	DialogueFlowType *string `json:"DialogueFlowType,omitempty" xml:"DialogueFlowType,omitempty"`
	// This parameter is required.
	DialogueName *string `json:"DialogueName,omitempty" xml:"DialogueName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 174952ab-9825-4cc9-a5e2-de82d7fa4cdd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d0bf93dd-1a54-4f00-819e-c75502d38681
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s CreateDialogueFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDialogueFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateDialogueFlowRequest) GetDialogueFlowType() *string {
	return s.DialogueFlowType
}

func (s *CreateDialogueFlowRequest) GetDialogueName() *string {
	return s.DialogueName
}

func (s *CreateDialogueFlowRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDialogueFlowRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateDialogueFlowRequest) SetDialogueFlowType(v string) *CreateDialogueFlowRequest {
	s.DialogueFlowType = &v
	return s
}

func (s *CreateDialogueFlowRequest) SetDialogueName(v string) *CreateDialogueFlowRequest {
	s.DialogueName = &v
	return s
}

func (s *CreateDialogueFlowRequest) SetInstanceId(v string) *CreateDialogueFlowRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDialogueFlowRequest) SetScriptId(v string) *CreateDialogueFlowRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateDialogueFlowRequest) Validate() error {
	return dara.Validate(s)
}
