// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIntentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeleteIntentRequest
	GetAgentKey() *string
	SetInstanceId(v string) *DeleteIntentRequest
	GetInstanceId() *string
	SetIntentId(v int64) *DeleteIntentRequest
	GetIntentId() *int64
}

type DeleteIntentRequest struct {
	// The key that identifies the business space. If you omit this parameter, the service uses the default business space. You can get this key from the Business Management page of your primary account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The robot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The intent ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
}

func (s DeleteIntentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIntentRequest) GoString() string {
	return s.String()
}

func (s *DeleteIntentRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeleteIntentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteIntentRequest) GetIntentId() *int64 {
	return s.IntentId
}

func (s *DeleteIntentRequest) SetAgentKey(v string) *DeleteIntentRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteIntentRequest) SetInstanceId(v string) *DeleteIntentRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteIntentRequest) SetIntentId(v int64) *DeleteIntentRequest {
	s.IntentId = &v
	return s
}

func (s *DeleteIntentRequest) Validate() error {
	return dara.Validate(s)
}
