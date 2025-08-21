// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserSayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeleteUserSayRequest
	GetAgentKey() *string
	SetInstanceId(v string) *DeleteUserSayRequest
	GetInstanceId() *string
	SetIntentId(v int64) *DeleteUserSayRequest
	GetIntentId() *int64
	SetUserSayId(v int64) *DeleteUserSayRequest
	GetUserSayId() *int64
}

type DeleteUserSayRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 343df2sdf23
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5564564546
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4562121234
	UserSayId *int64 `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s DeleteUserSayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserSayRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserSayRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeleteUserSayRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteUserSayRequest) GetIntentId() *int64 {
	return s.IntentId
}

func (s *DeleteUserSayRequest) GetUserSayId() *int64 {
	return s.UserSayId
}

func (s *DeleteUserSayRequest) SetAgentKey(v string) *DeleteUserSayRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteUserSayRequest) SetInstanceId(v string) *DeleteUserSayRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteUserSayRequest) SetIntentId(v int64) *DeleteUserSayRequest {
	s.IntentId = &v
	return s
}

func (s *DeleteUserSayRequest) SetUserSayId(v int64) *DeleteUserSayRequest {
	s.UserSayId = &v
	return s
}

func (s *DeleteUserSayRequest) Validate() error {
	return dara.Validate(s)
}
