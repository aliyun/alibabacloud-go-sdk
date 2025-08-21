// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDSEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateDSEntityRequest
	GetAgentKey() *string
	SetEntityId(v int64) *UpdateDSEntityRequest
	GetEntityId() *int64
	SetEntityName(v string) *UpdateDSEntityRequest
	GetEntityName() *string
	SetEntityType(v string) *UpdateDSEntityRequest
	GetEntityType() *string
	SetInstanceId(v string) *UpdateDSEntityRequest
	GetInstanceId() *string
}

type UpdateDSEntityRequest struct {
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
	// 实体名称
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// example:
	//
	// synonyms
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateDSEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDSEntityRequest) GoString() string {
	return s.String()
}

func (s *UpdateDSEntityRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateDSEntityRequest) GetEntityId() *int64 {
	return s.EntityId
}

func (s *UpdateDSEntityRequest) GetEntityName() *string {
	return s.EntityName
}

func (s *UpdateDSEntityRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *UpdateDSEntityRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateDSEntityRequest) SetAgentKey(v string) *UpdateDSEntityRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateDSEntityRequest) SetEntityId(v int64) *UpdateDSEntityRequest {
	s.EntityId = &v
	return s
}

func (s *UpdateDSEntityRequest) SetEntityName(v string) *UpdateDSEntityRequest {
	s.EntityName = &v
	return s
}

func (s *UpdateDSEntityRequest) SetEntityType(v string) *UpdateDSEntityRequest {
	s.EntityType = &v
	return s
}

func (s *UpdateDSEntityRequest) SetInstanceId(v string) *UpdateDSEntityRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateDSEntityRequest) Validate() error {
	return dara.Validate(s)
}
