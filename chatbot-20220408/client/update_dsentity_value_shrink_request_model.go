// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDSEntityValueShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateDSEntityValueShrinkRequest
	GetAgentKey() *string
	SetContent(v string) *UpdateDSEntityValueShrinkRequest
	GetContent() *string
	SetEntityId(v int64) *UpdateDSEntityValueShrinkRequest
	GetEntityId() *int64
	SetEntityValueId(v int64) *UpdateDSEntityValueShrinkRequest
	GetEntityValueId() *int64
	SetInstanceId(v string) *UpdateDSEntityValueShrinkRequest
	GetInstanceId() *string
	SetSynonymsShrink(v string) *UpdateDSEntityValueShrinkRequest
	GetSynonymsShrink() *string
}

type UpdateDSEntityValueShrinkRequest struct {
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
	SynonymsShrink *string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty"`
}

func (s UpdateDSEntityValueShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDSEntityValueShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDSEntityValueShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateDSEntityValueShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateDSEntityValueShrinkRequest) GetEntityId() *int64 {
	return s.EntityId
}

func (s *UpdateDSEntityValueShrinkRequest) GetEntityValueId() *int64 {
	return s.EntityValueId
}

func (s *UpdateDSEntityValueShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateDSEntityValueShrinkRequest) GetSynonymsShrink() *string {
	return s.SynonymsShrink
}

func (s *UpdateDSEntityValueShrinkRequest) SetAgentKey(v string) *UpdateDSEntityValueShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateDSEntityValueShrinkRequest) SetContent(v string) *UpdateDSEntityValueShrinkRequest {
	s.Content = &v
	return s
}

func (s *UpdateDSEntityValueShrinkRequest) SetEntityId(v int64) *UpdateDSEntityValueShrinkRequest {
	s.EntityId = &v
	return s
}

func (s *UpdateDSEntityValueShrinkRequest) SetEntityValueId(v int64) *UpdateDSEntityValueShrinkRequest {
	s.EntityValueId = &v
	return s
}

func (s *UpdateDSEntityValueShrinkRequest) SetInstanceId(v string) *UpdateDSEntityValueShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateDSEntityValueShrinkRequest) SetSynonymsShrink(v string) *UpdateDSEntityValueShrinkRequest {
	s.SynonymsShrink = &v
	return s
}

func (s *UpdateDSEntityValueShrinkRequest) Validate() error {
	return dara.Validate(s)
}
