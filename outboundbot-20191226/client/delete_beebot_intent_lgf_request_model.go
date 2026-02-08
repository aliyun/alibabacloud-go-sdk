// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBeebotIntentLgfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteBeebotIntentLgfRequest
	GetInstanceId() *string
	SetIntentId(v int64) *DeleteBeebotIntentLgfRequest
	GetIntentId() *int64
	SetLgfId(v int64) *DeleteBeebotIntentLgfRequest
	GetLgfId() *int64
	SetScriptId(v string) *DeleteBeebotIntentLgfRequest
	GetScriptId() *string
}

type DeleteBeebotIntentLgfRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Intent ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 10717802
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// Utterance template ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 5666117
	LgfId *int64 `json:"LgfId,omitempty" xml:"LgfId,omitempty"`
	// Scenario ID
	//
	// This parameter is required.
	//
	// example:
	//
	// c5c5d8c0-c0f1-48a7-be2b-dc46006d888a
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s DeleteBeebotIntentLgfRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBeebotIntentLgfRequest) GoString() string {
	return s.String()
}

func (s *DeleteBeebotIntentLgfRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteBeebotIntentLgfRequest) GetIntentId() *int64 {
	return s.IntentId
}

func (s *DeleteBeebotIntentLgfRequest) GetLgfId() *int64 {
	return s.LgfId
}

func (s *DeleteBeebotIntentLgfRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DeleteBeebotIntentLgfRequest) SetInstanceId(v string) *DeleteBeebotIntentLgfRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteBeebotIntentLgfRequest) SetIntentId(v int64) *DeleteBeebotIntentLgfRequest {
	s.IntentId = &v
	return s
}

func (s *DeleteBeebotIntentLgfRequest) SetLgfId(v int64) *DeleteBeebotIntentLgfRequest {
	s.LgfId = &v
	return s
}

func (s *DeleteBeebotIntentLgfRequest) SetScriptId(v string) *DeleteBeebotIntentLgfRequest {
	s.ScriptId = &v
	return s
}

func (s *DeleteBeebotIntentLgfRequest) Validate() error {
	return dara.Validate(s)
}
