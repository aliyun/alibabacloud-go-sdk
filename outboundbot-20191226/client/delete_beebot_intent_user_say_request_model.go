// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBeebotIntentUserSayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteBeebotIntentUserSayRequest
	GetInstanceId() *string
	SetIntentId(v string) *DeleteBeebotIntentUserSayRequest
	GetIntentId() *string
	SetScriptId(v string) *DeleteBeebotIntentUserSayRequest
	GetScriptId() *string
	SetUserSayId(v string) *DeleteBeebotIntentUserSayRequest
	GetUserSayId() *string
}

type DeleteBeebotIntentUserSayRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10717802
	IntentId *string `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c5c5d8c0-c0f1-48a7-be2b-dc46006d888a
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 17448458
	UserSayId *string `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s DeleteBeebotIntentUserSayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBeebotIntentUserSayRequest) GoString() string {
	return s.String()
}

func (s *DeleteBeebotIntentUserSayRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteBeebotIntentUserSayRequest) GetIntentId() *string {
	return s.IntentId
}

func (s *DeleteBeebotIntentUserSayRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DeleteBeebotIntentUserSayRequest) GetUserSayId() *string {
	return s.UserSayId
}

func (s *DeleteBeebotIntentUserSayRequest) SetInstanceId(v string) *DeleteBeebotIntentUserSayRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteBeebotIntentUserSayRequest) SetIntentId(v string) *DeleteBeebotIntentUserSayRequest {
	s.IntentId = &v
	return s
}

func (s *DeleteBeebotIntentUserSayRequest) SetScriptId(v string) *DeleteBeebotIntentUserSayRequest {
	s.ScriptId = &v
	return s
}

func (s *DeleteBeebotIntentUserSayRequest) SetUserSayId(v string) *DeleteBeebotIntentUserSayRequest {
	s.UserSayId = &v
	return s
}

func (s *DeleteBeebotIntentUserSayRequest) Validate() error {
	return dara.Validate(s)
}
