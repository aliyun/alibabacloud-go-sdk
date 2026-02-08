// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBeebotIntentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteBeebotIntentRequest
	GetInstanceId() *string
	SetIntentId(v int64) *DeleteBeebotIntentRequest
	GetIntentId() *int64
	SetScriptId(v string) *DeleteBeebotIntentRequest
	GetScriptId() *string
}

type DeleteBeebotIntentRequest struct {
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
	// Scenario ID
	//
	// This parameter is required.
	//
	// example:
	//
	// c5c5d8c0-c0f1-48a7-be2b-dc46006d888a
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s DeleteBeebotIntentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBeebotIntentRequest) GoString() string {
	return s.String()
}

func (s *DeleteBeebotIntentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteBeebotIntentRequest) GetIntentId() *int64 {
	return s.IntentId
}

func (s *DeleteBeebotIntentRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DeleteBeebotIntentRequest) SetInstanceId(v string) *DeleteBeebotIntentRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteBeebotIntentRequest) SetIntentId(v int64) *DeleteBeebotIntentRequest {
	s.IntentId = &v
	return s
}

func (s *DeleteBeebotIntentRequest) SetScriptId(v string) *DeleteBeebotIntentRequest {
	s.ScriptId = &v
	return s
}

func (s *DeleteBeebotIntentRequest) Validate() error {
	return dara.Validate(s)
}
