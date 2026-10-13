// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDSEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateDSEntityRequest
	GetAgentKey() *string
	SetEntityName(v string) *CreateDSEntityRequest
	GetEntityName() *string
	SetEntityType(v string) *CreateDSEntityRequest
	GetEntityType() *string
	SetInstanceId(v string) *CreateDSEntityRequest
	GetInstanceId() *string
}

type CreateDSEntityRequest struct {
	// The business space key. If this parameter is not set, the default business space is accessed. You can obtain the key value on the business management page of the main account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The entity name. Only Chinese characters, uppercase and lowercase letters, digits, and underscores are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// EntityName
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// The entity type. Valid values:
	//
	// - synonyms: synonym.
	//
	// - regex: regular expression.
	//
	// example:
	//
	// synonyms
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The chatbot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateDSEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDSEntityRequest) GoString() string {
	return s.String()
}

func (s *CreateDSEntityRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateDSEntityRequest) GetEntityName() *string {
	return s.EntityName
}

func (s *CreateDSEntityRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *CreateDSEntityRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDSEntityRequest) SetAgentKey(v string) *CreateDSEntityRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateDSEntityRequest) SetEntityName(v string) *CreateDSEntityRequest {
	s.EntityName = &v
	return s
}

func (s *CreateDSEntityRequest) SetEntityType(v string) *CreateDSEntityRequest {
	s.EntityType = &v
	return s
}

func (s *CreateDSEntityRequest) SetInstanceId(v string) *CreateDSEntityRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDSEntityRequest) Validate() error {
	return dara.Validate(s)
}
