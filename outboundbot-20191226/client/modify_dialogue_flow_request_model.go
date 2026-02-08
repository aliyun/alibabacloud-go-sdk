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
	// Conversation flow Data
	//
	// This parameter is required.
	//
	// example:
	//
	// {"transitions":[{"id":"a91c4023","index":1,"source":"cc31e02b","sourceAnchor":0,"target":"946045be","targetAnchor":0}],"nodes":[{"collectedNumberEnabled":false,"content":{"branches":[{"branchId":"f5450420-09ab-11ea-b107-e9059c6a79d8","branchName":"发起对话"}]},"coordinates":{"x":180,"y":134},"id":"cc31e02b","index":0,"interruptible":false,"nodeIndex":0,"shape":"start-html","size":"170*34","type":"start","x":180,"y":134},{"collectedNumberEnabled":false,"content":{"actionParams":"","action":"Hangup"},"coordinates":{"x":487.65625,"y":155},"id":"946045be","index":2,"interruptible":false,"labels":[],"name":"功能节点","nodeIndex":1,"questions":["好的同学，您的情况已了解了，稍后我们会安排资深顾问老师给您做更详细的留学评估以及升学指导，请留意电话接听"],"script":"好的同学，您的情况已了解了，稍后我们会安排资深顾问老师给您做更详细的留学评估以及升学指导，请留意电话接听","shape":"function-html","size":"170*34","type":"transfer","x":500,"y":182}]}
	DialogueFlowDefinition *string `json:"DialogueFlowDefinition,omitempty" xml:"DialogueFlowDefinition,omitempty"`
	// Conversation flow ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 390515b5-6115-4ccf-83e2-52d5bfaf2ddf
	DialogueFlowId *string `json:"DialogueFlowId,omitempty" xml:"DialogueFlowId,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the flow is in Draft status
	//
	// example:
	//
	// true
	IsDrafted *bool `json:"IsDrafted,omitempty" xml:"IsDrafted,omitempty"`
	// Script ID
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
