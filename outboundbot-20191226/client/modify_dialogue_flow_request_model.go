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
	// The dialog flow data.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"transitions":[{"id":"a91c4023","index":1,"source":"cc31e02b","sourceAnchor":0,"target":"946045be","targetAnchor":0}],"nodes":[{"collectedNumberEnabled":false,"content":{"branches":[{"branchId":"f5450420-09ab-11ea-b107-e9059c6a79d8","branchName":"Initiate conversation"}]},"coordinates":{"x":180,"y":134},"id":"cc31e02b","index":0,"interruptible":false,"nodeIndex":0,"shape":"start-html","size":"170*34","type":"start","x":180,"y":134},{"collectedNumberEnabled":false,"content":{"actionParams":"","action":"Hangup"},"coordinates":{"x":487.65625,"y":155},"id":"946045be","index":2,"interruptible":false,"labels":[],"name":"Function node","nodeIndex":1,"questions":["OK, we have noted your situation. We will arrange a senior consultant to provide you with a more detailed study abroad assessment and academic guidance. Please stay available for the call."],"script":"OK, we have noted your situation. We will arrange a senior consultant to provide you with a more detailed study abroad assessment and academic guidance. Please stay available for the call.","shape":"function-html","size":"170*34","type":"transfer","x":500,"y":182}]}
	DialogueFlowDefinition *string `json:"DialogueFlowDefinition,omitempty" xml:"DialogueFlowDefinition,omitempty"`
	// The dialog flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 390515b5-6115-4ccf-83e2-52d5bfaf2ddf
	DialogueFlowId *string `json:"DialogueFlowId,omitempty" xml:"DialogueFlowId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the dialog flow is in draft state.
	//
	// example:
	//
	// true
	IsDrafted *bool `json:"IsDrafted,omitempty" xml:"IsDrafted,omitempty"`
	// The script ID.
	//
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
