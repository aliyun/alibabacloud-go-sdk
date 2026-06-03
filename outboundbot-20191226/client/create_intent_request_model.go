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
	// This parameter is required.
	//
	// example:
	//
	// 361c8a53-0e29-42f3-8aa7-c7752d010399
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IntentDescription *string `json:"IntentDescription,omitempty" xml:"IntentDescription,omitempty"`
	// This parameter is required.
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	Keywords   *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// b06fad9a-cc74-4ab6-b3a5-8d062adebf2c
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// This parameter is required.
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
