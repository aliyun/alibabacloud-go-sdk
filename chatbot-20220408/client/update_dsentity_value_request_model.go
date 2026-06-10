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
	// The key for the business space. If you omit this parameter, the default business space is used. You can find this key on the Business Management page of your primary account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The new content for the entity value. For an entity type of `synonyms`, this is the normalized value. For an entity type of `regex`, this is the regular expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// 书本类型
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The entity ID. You can leave this parameter empty when modifying an entity value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 223423423
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The ID of the entity value to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2342377423
	EntityValueId *int64 `json:"EntityValueId,omitempty" xml:"EntityValueId,omitempty"`
	// The bot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The synonym list for the normalized value.
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
