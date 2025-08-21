// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDSEntityValueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeleteDSEntityValueRequest
	GetAgentKey() *string
	SetEntityId(v int64) *DeleteDSEntityValueRequest
	GetEntityId() *int64
	SetEntityValueId(v int64) *DeleteDSEntityValueRequest
	GetEntityValueId() *int64
	SetInstanceId(v string) *DeleteDSEntityValueRequest
	GetInstanceId() *string
}

type DeleteDSEntityValueRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 345346223452
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3453453452
	EntityValueId *int64 `json:"EntityValueId,omitempty" xml:"EntityValueId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteDSEntityValueRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDSEntityValueRequest) GoString() string {
	return s.String()
}

func (s *DeleteDSEntityValueRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeleteDSEntityValueRequest) GetEntityId() *int64 {
	return s.EntityId
}

func (s *DeleteDSEntityValueRequest) GetEntityValueId() *int64 {
	return s.EntityValueId
}

func (s *DeleteDSEntityValueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDSEntityValueRequest) SetAgentKey(v string) *DeleteDSEntityValueRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteDSEntityValueRequest) SetEntityId(v int64) *DeleteDSEntityValueRequest {
	s.EntityId = &v
	return s
}

func (s *DeleteDSEntityValueRequest) SetEntityValueId(v int64) *DeleteDSEntityValueRequest {
	s.EntityValueId = &v
	return s
}

func (s *DeleteDSEntityValueRequest) SetInstanceId(v string) *DeleteDSEntityValueRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDSEntityValueRequest) Validate() error {
	return dara.Validate(s)
}
