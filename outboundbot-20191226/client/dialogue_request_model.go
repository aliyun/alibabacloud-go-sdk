// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDialogueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionKey(v string) *DialogueRequest
	GetActionKey() *string
	SetActionParams(v string) *DialogueRequest
	GetActionParams() *string
	SetCallId(v string) *DialogueRequest
	GetCallId() *string
	SetCallType(v string) *DialogueRequest
	GetCallType() *string
	SetCalledNumber(v string) *DialogueRequest
	GetCalledNumber() *string
	SetCallingNumber(v string) *DialogueRequest
	GetCallingNumber() *string
	SetInstanceId(v string) *DialogueRequest
	GetInstanceId() *string
	SetScenarioId(v string) *DialogueRequest
	GetScenarioId() *string
	SetScriptId(v string) *DialogueRequest
	GetScriptId() *string
	SetTaskId(v string) *DialogueRequest
	GetTaskId() *string
	SetUtterance(v string) *DialogueRequest
	GetUtterance() *string
}

type DialogueRequest struct {
	// example:
	//
	// broadcast
	ActionKey    *string `json:"ActionKey,omitempty" xml:"ActionKey,omitempty"`
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1528189846043
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Outbound
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 135****4353
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1***6
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 6cea9bed-63e6-439e-ae4c-b3333efff53d
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// 场景id
	//
	// example:
	//
	// c5c5d8c0-c0f1-48a7-be2b-dc46006d888a
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// ff44709e-39a6-43ba-959b-20fcabe3e496
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
}

func (s DialogueRequest) String() string {
	return dara.Prettify(s)
}

func (s DialogueRequest) GoString() string {
	return s.String()
}

func (s *DialogueRequest) GetActionKey() *string {
	return s.ActionKey
}

func (s *DialogueRequest) GetActionParams() *string {
	return s.ActionParams
}

func (s *DialogueRequest) GetCallId() *string {
	return s.CallId
}

func (s *DialogueRequest) GetCallType() *string {
	return s.CallType
}

func (s *DialogueRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *DialogueRequest) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *DialogueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DialogueRequest) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *DialogueRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DialogueRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DialogueRequest) GetUtterance() *string {
	return s.Utterance
}

func (s *DialogueRequest) SetActionKey(v string) *DialogueRequest {
	s.ActionKey = &v
	return s
}

func (s *DialogueRequest) SetActionParams(v string) *DialogueRequest {
	s.ActionParams = &v
	return s
}

func (s *DialogueRequest) SetCallId(v string) *DialogueRequest {
	s.CallId = &v
	return s
}

func (s *DialogueRequest) SetCallType(v string) *DialogueRequest {
	s.CallType = &v
	return s
}

func (s *DialogueRequest) SetCalledNumber(v string) *DialogueRequest {
	s.CalledNumber = &v
	return s
}

func (s *DialogueRequest) SetCallingNumber(v string) *DialogueRequest {
	s.CallingNumber = &v
	return s
}

func (s *DialogueRequest) SetInstanceId(v string) *DialogueRequest {
	s.InstanceId = &v
	return s
}

func (s *DialogueRequest) SetScenarioId(v string) *DialogueRequest {
	s.ScenarioId = &v
	return s
}

func (s *DialogueRequest) SetScriptId(v string) *DialogueRequest {
	s.ScriptId = &v
	return s
}

func (s *DialogueRequest) SetTaskId(v string) *DialogueRequest {
	s.TaskId = &v
	return s
}

func (s *DialogueRequest) SetUtterance(v string) *DialogueRequest {
	s.Utterance = &v
	return s
}

func (s *DialogueRequest) Validate() error {
	return dara.Validate(s)
}
