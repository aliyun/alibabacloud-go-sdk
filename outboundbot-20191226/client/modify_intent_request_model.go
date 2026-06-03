// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIntentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyIntentRequest
	GetInstanceId() *string
	SetIntentDescription(v string) *ModifyIntentRequest
	GetIntentDescription() *string
	SetIntentId(v string) *ModifyIntentRequest
	GetIntentId() *string
	SetIntentName(v string) *ModifyIntentRequest
	GetIntentName() *string
	SetKeywords(v string) *ModifyIntentRequest
	GetKeywords() *string
	SetScriptId(v string) *ModifyIntentRequest
	GetScriptId() *string
	SetUtterances(v string) *ModifyIntentRequest
	GetUtterances() *string
}

type ModifyIntentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 8fa1953f-4a84-46d8-b80c-8ce9cf684fb3
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IntentDescription *string `json:"IntentDescription,omitempty" xml:"IntentDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3b9a2b33-50d4-4576-8c68-22498f4bf731
	IntentId *string `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// This parameter is required.
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// example:
	//
	// []
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9b70486f-b1c2-429c-8a24-62798015ab1b
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// This parameter is required.
	Utterances *string `json:"Utterances,omitempty" xml:"Utterances,omitempty"`
}

func (s ModifyIntentRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyIntentRequest) GoString() string {
	return s.String()
}

func (s *ModifyIntentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyIntentRequest) GetIntentDescription() *string {
	return s.IntentDescription
}

func (s *ModifyIntentRequest) GetIntentId() *string {
	return s.IntentId
}

func (s *ModifyIntentRequest) GetIntentName() *string {
	return s.IntentName
}

func (s *ModifyIntentRequest) GetKeywords() *string {
	return s.Keywords
}

func (s *ModifyIntentRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyIntentRequest) GetUtterances() *string {
	return s.Utterances
}

func (s *ModifyIntentRequest) SetInstanceId(v string) *ModifyIntentRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyIntentRequest) SetIntentDescription(v string) *ModifyIntentRequest {
	s.IntentDescription = &v
	return s
}

func (s *ModifyIntentRequest) SetIntentId(v string) *ModifyIntentRequest {
	s.IntentId = &v
	return s
}

func (s *ModifyIntentRequest) SetIntentName(v string) *ModifyIntentRequest {
	s.IntentName = &v
	return s
}

func (s *ModifyIntentRequest) SetKeywords(v string) *ModifyIntentRequest {
	s.Keywords = &v
	return s
}

func (s *ModifyIntentRequest) SetScriptId(v string) *ModifyIntentRequest {
	s.ScriptId = &v
	return s
}

func (s *ModifyIntentRequest) SetUtterances(v string) *ModifyIntentRequest {
	s.Utterances = &v
	return s
}

func (s *ModifyIntentRequest) Validate() error {
	return dara.Validate(s)
}
