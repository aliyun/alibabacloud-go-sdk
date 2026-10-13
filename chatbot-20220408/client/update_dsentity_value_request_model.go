// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDSEntityValueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateDSEntityValueRequest
	GetAgentKey() *string
	SetContent(v string) *UpdateDSEntityValueRequest
	GetContent() *string
	SetEntityId(v int64) *UpdateDSEntityValueRequest
	GetEntityId() *int64
	SetEntityValueId(v int64) *UpdateDSEntityValueRequest
	GetEntityValueId() *int64
	SetInstanceId(v string) *UpdateDSEntityValueRequest
	GetInstanceId() *string
	SetSynonyms(v []*string) *UpdateDSEntityValueRequest
	GetSynonyms() []*string
}

type UpdateDSEntityValueRequest struct {
	// The business space key. If this parameter is not set, the default business space is accessed. Obtain the key value from the business management page of the main account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// When the entity type is synonyms, this parameter specifies the normalized value of the entity. When the entity type is regex, this parameter specifies the regular expression text.
	//
	// This parameter is required.
	//
	// example:
	//
	// Book type
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The entity ID. This parameter can be left empty when you modify an entity member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 223423423
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The entity member ID. This parameter is left empty when you create an entity member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2342377423
	EntityValueId *int64 `json:"EntityValueId,omitempty" xml:"EntityValueId,omitempty"`
	// The chatbot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of synonyms.
	Synonyms []*string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty" type:"Repeated"`
}

func (s UpdateDSEntityValueRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDSEntityValueRequest) GoString() string {
	return s.String()
}

func (s *UpdateDSEntityValueRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateDSEntityValueRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateDSEntityValueRequest) GetEntityId() *int64 {
	return s.EntityId
}

func (s *UpdateDSEntityValueRequest) GetEntityValueId() *int64 {
	return s.EntityValueId
}

func (s *UpdateDSEntityValueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateDSEntityValueRequest) GetSynonyms() []*string {
	return s.Synonyms
}

func (s *UpdateDSEntityValueRequest) SetAgentKey(v string) *UpdateDSEntityValueRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateDSEntityValueRequest) SetContent(v string) *UpdateDSEntityValueRequest {
	s.Content = &v
	return s
}

func (s *UpdateDSEntityValueRequest) SetEntityId(v int64) *UpdateDSEntityValueRequest {
	s.EntityId = &v
	return s
}

func (s *UpdateDSEntityValueRequest) SetEntityValueId(v int64) *UpdateDSEntityValueRequest {
	s.EntityValueId = &v
	return s
}

func (s *UpdateDSEntityValueRequest) SetInstanceId(v string) *UpdateDSEntityValueRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateDSEntityValueRequest) SetSynonyms(v []*string) *UpdateDSEntityValueRequest {
	s.Synonyms = v
	return s
}

func (s *UpdateDSEntityValueRequest) Validate() error {
	return dara.Validate(s)
}
