// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDSEntityValueShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateDSEntityValueShrinkRequest
	GetAgentKey() *string
	SetContent(v string) *CreateDSEntityValueShrinkRequest
	GetContent() *string
	SetEntityId(v int64) *CreateDSEntityValueShrinkRequest
	GetEntityId() *int64
	SetInstanceId(v string) *CreateDSEntityValueShrinkRequest
	GetInstanceId() *string
	SetSynonymsShrink(v string) *CreateDSEntityValueShrinkRequest
	GetSynonymsShrink() *string
}

type CreateDSEntityValueShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ada
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SynonymsShrink *string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty"`
}

func (s CreateDSEntityValueShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDSEntityValueShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDSEntityValueShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateDSEntityValueShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *CreateDSEntityValueShrinkRequest) GetEntityId() *int64 {
	return s.EntityId
}

func (s *CreateDSEntityValueShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDSEntityValueShrinkRequest) GetSynonymsShrink() *string {
	return s.SynonymsShrink
}

func (s *CreateDSEntityValueShrinkRequest) SetAgentKey(v string) *CreateDSEntityValueShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateDSEntityValueShrinkRequest) SetContent(v string) *CreateDSEntityValueShrinkRequest {
	s.Content = &v
	return s
}

func (s *CreateDSEntityValueShrinkRequest) SetEntityId(v int64) *CreateDSEntityValueShrinkRequest {
	s.EntityId = &v
	return s
}

func (s *CreateDSEntityValueShrinkRequest) SetInstanceId(v string) *CreateDSEntityValueShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDSEntityValueShrinkRequest) SetSynonymsShrink(v string) *CreateDSEntityValueShrinkRequest {
	s.SynonymsShrink = &v
	return s
}

func (s *CreateDSEntityValueShrinkRequest) Validate() error {
	return dara.Validate(s)
}
