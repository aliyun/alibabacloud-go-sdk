// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIconUrl(v string) *CreateAgentRequest
	GetAgentIconUrl() *string
	SetAgentName(v string) *CreateAgentRequest
	GetAgentName() *string
	SetAgentScene(v string) *CreateAgentRequest
	GetAgentScene() *string
	SetCharacterAgeStage(v string) *CreateAgentRequest
	GetCharacterAgeStage() *string
	SetCharacterGender(v string) *CreateAgentRequest
	GetCharacterGender() *string
	SetCharacterName(v string) *CreateAgentRequest
	GetCharacterName() *string
	SetExtraDescription(v string) *CreateAgentRequest
	GetExtraDescription() *string
	SetIndustry(v string) *CreateAgentRequest
	GetIndustry() *string
}

type CreateAgentRequest struct {
	// example:
	//
	// http://img.com
	AgentIconUrl *string `json:"agentIconUrl,omitempty" xml:"agentIconUrl,omitempty"`
	// example:
	//
	// AgentAlpha
	AgentName *string `json:"agentName,omitempty" xml:"agentName,omitempty"`
	// example:
	//
	// aiCoachPractice
	AgentScene *string `json:"agentScene,omitempty" xml:"agentScene,omitempty"`
	// example:
	//
	// 18-22
	CharacterAgeStage *string `json:"characterAgeStage,omitempty" xml:"characterAgeStage,omitempty"`
	CharacterGender   *string `json:"characterGender,omitempty" xml:"characterGender,omitempty"`
	// example:
	//
	// Tom
	CharacterName    *string `json:"characterName,omitempty" xml:"characterName,omitempty"`
	ExtraDescription *string `json:"extraDescription,omitempty" xml:"extraDescription,omitempty"`
	// example:
	//
	// Common
	Industry *string `json:"industry,omitempty" xml:"industry,omitempty"`
}

func (s CreateAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentRequest) GetAgentIconUrl() *string {
	return s.AgentIconUrl
}

func (s *CreateAgentRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *CreateAgentRequest) GetAgentScene() *string {
	return s.AgentScene
}

func (s *CreateAgentRequest) GetCharacterAgeStage() *string {
	return s.CharacterAgeStage
}

func (s *CreateAgentRequest) GetCharacterGender() *string {
	return s.CharacterGender
}

func (s *CreateAgentRequest) GetCharacterName() *string {
	return s.CharacterName
}

func (s *CreateAgentRequest) GetExtraDescription() *string {
	return s.ExtraDescription
}

func (s *CreateAgentRequest) GetIndustry() *string {
	return s.Industry
}

func (s *CreateAgentRequest) SetAgentIconUrl(v string) *CreateAgentRequest {
	s.AgentIconUrl = &v
	return s
}

func (s *CreateAgentRequest) SetAgentName(v string) *CreateAgentRequest {
	s.AgentName = &v
	return s
}

func (s *CreateAgentRequest) SetAgentScene(v string) *CreateAgentRequest {
	s.AgentScene = &v
	return s
}

func (s *CreateAgentRequest) SetCharacterAgeStage(v string) *CreateAgentRequest {
	s.CharacterAgeStage = &v
	return s
}

func (s *CreateAgentRequest) SetCharacterGender(v string) *CreateAgentRequest {
	s.CharacterGender = &v
	return s
}

func (s *CreateAgentRequest) SetCharacterName(v string) *CreateAgentRequest {
	s.CharacterName = &v
	return s
}

func (s *CreateAgentRequest) SetExtraDescription(v string) *CreateAgentRequest {
	s.ExtraDescription = &v
	return s
}

func (s *CreateAgentRequest) SetIndustry(v string) *CreateAgentRequest {
	s.Industry = &v
	return s
}

func (s *CreateAgentRequest) Validate() error {
	return dara.Validate(s)
}
