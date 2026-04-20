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
	// example:
	//
	// [{
	//
	// 	"dataType": "Silence"
	//
	// }]
	Extras *string `json:"Extras,omitempty" xml:"Extras,omitempty"`
	// example:
	//
	// 36e9a4cd-a821-4f78-86fa-a9a4aefeea2e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 0a88e269-01f5-49ac-97af-5830f0ccb271
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// 07d72ea0-6e38-48d4-8371-14deaaba996f
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
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
