// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDSEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeleteDSEntityRequest
	GetAgentKey() *string
	SetEntityId(v int64) *DeleteDSEntityRequest
	GetEntityId() *int64
	SetInstanceId(v string) *DeleteDSEntityRequest
	GetInstanceId() *string
}

type DeleteDSEntityRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteDSEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDSEntityRequest) GoString() string {
	return s.String()
}

func (s *DeleteDSEntityRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeleteDSEntityRequest) GetEntityId() *int64 {
	return s.EntityId
}

func (s *DeleteDSEntityRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDSEntityRequest) SetAgentKey(v string) *DeleteDSEntityRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteDSEntityRequest) SetEntityId(v int64) *DeleteDSEntityRequest {
	s.EntityId = &v
	return s
}

func (s *DeleteDSEntityRequest) SetInstanceId(v string) *DeleteDSEntityRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDSEntityRequest) Validate() error {
	return dara.Validate(s)
}
