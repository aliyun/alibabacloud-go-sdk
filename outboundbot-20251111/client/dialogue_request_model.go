// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDialogueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtras(v string) *DialogueRequest
	GetExtras() *string
	SetInstanceId(v string) *DialogueRequest
	GetInstanceId() *string
	SetScriptId(v string) *DialogueRequest
	GetScriptId() *string
	SetSessionId(v string) *DialogueRequest
	GetSessionId() *string
	SetUtterance(v string) *DialogueRequest
	GetUtterance() *string
}

type DialogueRequest struct {
	// The JSON array string of additional input events. Supports DTMF key press events (dataType=DTMF, dtmf is the key content), Silence events (dataType=Silence), and RuntimePrompts events (dataType=RuntimePrompts, prompts is the prompt text, persistent specifies whether to persist). For example, a silence event is [{"dataType":"Silence"}]. This parameter is optional when no additional events are present.
	//
	// example:
	//
	// [{"dataType":"Silence"}]
	Extras *string `json:"Extras,omitempty" xml:"Extras,omitempty"`
	// The instance identity of the associated product. You must use an instance that the invoker has permission to access and that actually exists. It must be consistent with the instance to which the scenario, channel, and session belong. The maximum length is subject to the current POP parameter definition.
	//
	// example:
	//
	// instance-example-001
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the conversation scenario under the instance. Specify an existing scenario. Subsequent conversations and end requests for the same text session must use the scenario specified when the session was initiated.
	//
	// example:
	//
	// script-example-001
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The session ID used to uniquely associate this interaction. For a text session, it is the text- identifier returned when the session is initiated. For web voice, it is the ws- identifier. For real number outbound call tests, it is the test- identifier. Subsequent calls must use the actual returned value.
	//
	// example:
	//
	// text-11111111-2222-4333-8444-555555555555
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The text input content of the user for the current round. For normal conversations, pass the actual text of the user. Events such as key presses or silence can be passed through Extras. The processing method is determined by the scenario and the engine.
	//
	// example:
	//
	// I would like to inquire about the business handling process
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
}

func (s DialogueRequest) String() string {
	return dara.Prettify(s)
}

func (s DialogueRequest) GoString() string {
	return s.String()
}

func (s *DialogueRequest) GetExtras() *string {
	return s.Extras
}

func (s *DialogueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DialogueRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DialogueRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DialogueRequest) GetUtterance() *string {
	return s.Utterance
}

func (s *DialogueRequest) SetExtras(v string) *DialogueRequest {
	s.Extras = &v
	return s
}

func (s *DialogueRequest) SetInstanceId(v string) *DialogueRequest {
	s.InstanceId = &v
	return s
}

func (s *DialogueRequest) SetScriptId(v string) *DialogueRequest {
	s.ScriptId = &v
	return s
}

func (s *DialogueRequest) SetSessionId(v string) *DialogueRequest {
	s.SessionId = &v
	return s
}

func (s *DialogueRequest) SetUtterance(v string) *DialogueRequest {
	s.Utterance = &v
	return s
}

func (s *DialogueRequest) Validate() error {
	return dara.Validate(s)
}
