// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateIntentRequest
	GetInstanceId() *string
	SetIntentDescription(v string) *CreateIntentRequest
	GetIntentDescription() *string
	SetIntentName(v string) *CreateIntentRequest
	GetIntentName() *string
	SetKeywords(v string) *CreateIntentRequest
	GetKeywords() *string
	SetScriptId(v string) *CreateIntentRequest
	GetScriptId() *string
	SetUtterances(v string) *CreateIntentRequest
	GetUtterances() *string
}

type CreateIntentRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 361c8a53-0e29-42f3-8aa7-c7752d010399
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Intent description
	//
	// example:
	//
	// 确定是本人的意图
	IntentDescription *string `json:"IntentDescription,omitempty" xml:"IntentDescription,omitempty"`
	// Intent name
	//
	// This parameter is required.
	//
	// example:
	//
	// 是本人
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// Keyword
	//
	// example:
	//
	// ["是","是的"]
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// Scenario ID
	//
	// This parameter is required.
	//
	// example:
	//
	// c5c5d8c0-c0f1-48a7-be2b-dc46006d888a
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// User utterances
	//
	// This parameter is required.
	//
	// example:
	//
	// ["是","是的","是啊","嗯嗯","是我","是我，有事吗","对，是","对的","对是我",""]
	Utterances *string `json:"Utterances,omitempty" xml:"Utterances,omitempty"`
}

func (s CreateIntentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIntentRequest) GoString() string {
	return s.String()
}

func (s *CreateIntentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateIntentRequest) GetIntentDescription() *string {
	return s.IntentDescription
}

func (s *CreateIntentRequest) GetIntentName() *string {
	return s.IntentName
}

func (s *CreateIntentRequest) GetKeywords() *string {
	return s.Keywords
}

func (s *CreateIntentRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateIntentRequest) GetUtterances() *string {
	return s.Utterances
}

func (s *CreateIntentRequest) SetInstanceId(v string) *CreateIntentRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateIntentRequest) SetIntentDescription(v string) *CreateIntentRequest {
	s.IntentDescription = &v
	return s
}

func (s *CreateIntentRequest) SetIntentName(v string) *CreateIntentRequest {
	s.IntentName = &v
	return s
}

func (s *CreateIntentRequest) SetKeywords(v string) *CreateIntentRequest {
	s.Keywords = &v
	return s
}

func (s *CreateIntentRequest) SetScriptId(v string) *CreateIntentRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateIntentRequest) SetUtterances(v string) *CreateIntentRequest {
	s.Utterances = &v
	return s
}

func (s *CreateIntentRequest) Validate() error {
	return dara.Validate(s)
}
