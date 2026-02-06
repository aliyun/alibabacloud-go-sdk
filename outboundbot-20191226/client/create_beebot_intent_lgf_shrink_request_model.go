// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBeebotIntentLgfShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateBeebotIntentLgfShrinkRequest
	GetInstanceId() *string
	SetLgfDefinitionShrink(v string) *CreateBeebotIntentLgfShrinkRequest
	GetLgfDefinitionShrink() *string
	SetScriptId(v string) *CreateBeebotIntentLgfShrinkRequest
	GetScriptId() *string
}

type CreateBeebotIntentLgfShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	LgfDefinitionShrink *string `json:"LgfDefinition,omitempty" xml:"LgfDefinition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c5c5d8c0-c0f1-48a7-be2b-dc46006d888a
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s CreateBeebotIntentLgfShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBeebotIntentLgfShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateBeebotIntentLgfShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateBeebotIntentLgfShrinkRequest) GetLgfDefinitionShrink() *string {
	return s.LgfDefinitionShrink
}

func (s *CreateBeebotIntentLgfShrinkRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateBeebotIntentLgfShrinkRequest) SetInstanceId(v string) *CreateBeebotIntentLgfShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBeebotIntentLgfShrinkRequest) SetLgfDefinitionShrink(v string) *CreateBeebotIntentLgfShrinkRequest {
	s.LgfDefinitionShrink = &v
	return s
}

func (s *CreateBeebotIntentLgfShrinkRequest) SetScriptId(v string) *CreateBeebotIntentLgfShrinkRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateBeebotIntentLgfShrinkRequest) Validate() error {
	return dara.Validate(s)
}
