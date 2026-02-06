// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDialogueFlowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListDialogueFlowsRequest
	GetInstanceId() *string
	SetScriptId(v string) *ListDialogueFlowsRequest
	GetScriptId() *string
}

type ListDialogueFlowsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 361c8a53-0e29-42f3-8aa7-c7752d010399
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d7c28efb-47f7-4a85-a522-5038e30a0b98
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s ListDialogueFlowsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDialogueFlowsRequest) GoString() string {
	return s.String()
}

func (s *ListDialogueFlowsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDialogueFlowsRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListDialogueFlowsRequest) SetInstanceId(v string) *ListDialogueFlowsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDialogueFlowsRequest) SetScriptId(v string) *ListDialogueFlowsRequest {
	s.ScriptId = &v
	return s
}

func (s *ListDialogueFlowsRequest) Validate() error {
	return dara.Validate(s)
}
