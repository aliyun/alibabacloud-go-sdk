// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIconUrl(v string) *UpdateAgentRequest
	GetAgentIconUrl() *string
	SetAgentId(v string) *UpdateAgentRequest
	GetAgentId() *string
	SetAgentName(v string) *UpdateAgentRequest
	GetAgentName() *string
	SetCharacterAgeStage(v string) *UpdateAgentRequest
	GetCharacterAgeStage() *string
	SetCharacterGender(v string) *UpdateAgentRequest
	GetCharacterGender() *string
	SetCharacterName(v string) *UpdateAgentRequest
	GetCharacterName() *string
	SetExtraDescription(v string) *UpdateAgentRequest
	GetExtraDescription() *string
	SetIndustry(v string) *UpdateAgentRequest
	GetIndustry() *string
}

type UpdateAgentRequest struct {
	AgentIconUrl      *string `json:"agentIconUrl,omitempty" xml:"agentIconUrl,omitempty"`
	AgentId           *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	AgentName         *string `json:"agentName,omitempty" xml:"agentName,omitempty"`
	CharacterAgeStage *string `json:"characterAgeStage,omitempty" xml:"characterAgeStage,omitempty"`
	CharacterGender   *string `json:"characterGender,omitempty" xml:"characterGender,omitempty"`
	CharacterName     *string `json:"characterName,omitempty" xml:"characterName,omitempty"`
	ExtraDescription  *string `json:"extraDescription,omitempty" xml:"extraDescription,omitempty"`
	Industry          *string `json:"industry,omitempty" xml:"industry,omitempty"`
}

func (s UpdateAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentRequest) GetAgentIconUrl() *string {
	return s.AgentIconUrl
}

func (s *UpdateAgentRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdateAgentRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *UpdateAgentRequest) GetCharacterAgeStage() *string {
	return s.CharacterAgeStage
}

func (s *UpdateAgentRequest) GetCharacterGender() *string {
	return s.CharacterGender
}

func (s *UpdateAgentRequest) GetCharacterName() *string {
	return s.CharacterName
}

func (s *UpdateAgentRequest) GetExtraDescription() *string {
	return s.ExtraDescription
}

func (s *UpdateAgentRequest) GetIndustry() *string {
	return s.Industry
}

func (s *UpdateAgentRequest) SetAgentIconUrl(v string) *UpdateAgentRequest {
	s.AgentIconUrl = &v
	return s
}

func (s *UpdateAgentRequest) SetAgentId(v string) *UpdateAgentRequest {
	s.AgentId = &v
	return s
}

func (s *UpdateAgentRequest) SetAgentName(v string) *UpdateAgentRequest {
	s.AgentName = &v
	return s
}

func (s *UpdateAgentRequest) SetCharacterAgeStage(v string) *UpdateAgentRequest {
	s.CharacterAgeStage = &v
	return s
}

func (s *UpdateAgentRequest) SetCharacterGender(v string) *UpdateAgentRequest {
	s.CharacterGender = &v
	return s
}

func (s *UpdateAgentRequest) SetCharacterName(v string) *UpdateAgentRequest {
	s.CharacterName = &v
	return s
}

func (s *UpdateAgentRequest) SetExtraDescription(v string) *UpdateAgentRequest {
	s.ExtraDescription = &v
	return s
}

func (s *UpdateAgentRequest) SetIndustry(v string) *UpdateAgentRequest {
	s.Industry = &v
	return s
}

func (s *UpdateAgentRequest) Validate() error {
	return dara.Validate(s)
}
