// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDSEntityValueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateDSEntityValueRequest
	GetAgentKey() *string
	SetContent(v string) *CreateDSEntityValueRequest
	GetContent() *string
	SetEntityId(v int64) *CreateDSEntityValueRequest
	GetEntityId() *int64
	SetInstanceId(v string) *CreateDSEntityValueRequest
	GetInstanceId() *string
	SetSynonyms(v []*string) *CreateDSEntityValueRequest
	GetSynonyms() []*string
}

type CreateDSEntityValueRequest struct {
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
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Synonyms   []*string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty" type:"Repeated"`
}

func (s CreateDSEntityValueRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDSEntityValueRequest) GoString() string {
	return s.String()
}

func (s *CreateDSEntityValueRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateDSEntityValueRequest) GetContent() *string {
	return s.Content
}

func (s *CreateDSEntityValueRequest) GetEntityId() *int64 {
	return s.EntityId
}

func (s *CreateDSEntityValueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDSEntityValueRequest) GetSynonyms() []*string {
	return s.Synonyms
}

func (s *CreateDSEntityValueRequest) SetAgentKey(v string) *CreateDSEntityValueRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateDSEntityValueRequest) SetContent(v string) *CreateDSEntityValueRequest {
	s.Content = &v
	return s
}

func (s *CreateDSEntityValueRequest) SetEntityId(v int64) *CreateDSEntityValueRequest {
	s.EntityId = &v
	return s
}

func (s *CreateDSEntityValueRequest) SetInstanceId(v string) *CreateDSEntityValueRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDSEntityValueRequest) SetSynonyms(v []*string) *CreateDSEntityValueRequest {
	s.Synonyms = v
	return s
}

func (s *CreateDSEntityValueRequest) Validate() error {
	return dara.Validate(s)
}
