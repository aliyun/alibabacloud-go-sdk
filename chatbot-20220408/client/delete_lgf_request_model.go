// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLgfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeleteLgfRequest
	GetAgentKey() *string
	SetInstanceId(v string) *DeleteLgfRequest
	GetInstanceId() *string
	SetIntentId(v int64) *DeleteLgfRequest
	GetIntentId() *int64
	SetLgfId(v int64) *DeleteLgfRequest
	GetLgfId() *int64
}

type DeleteLgfRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 23242342
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// lgf Id
	//
	// This parameter is required.
	//
	// example:
	//
	// 2342424
	LgfId *int64 `json:"LgfId,omitempty" xml:"LgfId,omitempty"`
}

func (s DeleteLgfRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLgfRequest) GoString() string {
	return s.String()
}

func (s *DeleteLgfRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeleteLgfRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteLgfRequest) GetIntentId() *int64 {
	return s.IntentId
}

func (s *DeleteLgfRequest) GetLgfId() *int64 {
	return s.LgfId
}

func (s *DeleteLgfRequest) SetAgentKey(v string) *DeleteLgfRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteLgfRequest) SetInstanceId(v string) *DeleteLgfRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteLgfRequest) SetIntentId(v int64) *DeleteLgfRequest {
	s.IntentId = &v
	return s
}

func (s *DeleteLgfRequest) SetLgfId(v int64) *DeleteLgfRequest {
	s.LgfId = &v
	return s
}

func (s *DeleteLgfRequest) Validate() error {
	return dara.Validate(s)
}
