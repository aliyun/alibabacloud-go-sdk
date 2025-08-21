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
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 223423423
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2342377423
	EntityValueId *int64 `json:"EntityValueId,omitempty" xml:"EntityValueId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
