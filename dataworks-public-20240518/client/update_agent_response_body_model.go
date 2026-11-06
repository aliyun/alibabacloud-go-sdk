// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgent(v *UpdateAgentResponseBodyAgent) *UpdateAgentResponseBody
	GetAgent() *UpdateAgentResponseBodyAgent
	SetRequestId(v string) *UpdateAgentResponseBody
	GetRequestId() *string
}

type UpdateAgentResponseBody struct {
	// The details of the updated agent.
	Agent *UpdateAgentResponseBodyAgent `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 824F80BA-1778-5D8A-BAFF-668A4D9C4CC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentResponseBody) GetAgent() *UpdateAgentResponseBodyAgent {
	return s.Agent
}

func (s *UpdateAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAgentResponseBody) SetAgent(v *UpdateAgentResponseBodyAgent) *UpdateAgentResponseBody {
	s.Agent = v
	return s
}

func (s *UpdateAgentResponseBody) SetRequestId(v string) *UpdateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgentResponseBody) Validate() error {
	if s.Agent != nil {
		if err := s.Agent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAgentResponseBodyAgent struct {
	// The version of the updated agent definition.
	//
	// example:
	//
	// 3
	AgentVersion *int32 `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// The bound expert suites.
	ExpertSuites []*UpdateAgentResponseBodyAgentExpertSuites `json:"ExpertSuites,omitempty" xml:"ExpertSuites,omitempty" type:"Repeated"`
	// The creation time of the agent. It is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1780555634000
	GmtCreateTime *int64 `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The last modified time of the agent. It is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1780555634000
	GmtModifiedTime *int64 `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The name of the agent.
	//
	// example:
	//
	// my-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of bound skill references after the update.
	Skills []*UpdateAgentResponseBodyAgentSkills `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
}

func (s UpdateAgentResponseBodyAgent) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentResponseBodyAgent) GoString() string {
	return s.String()
}

func (s *UpdateAgentResponseBodyAgent) GetAgentVersion() *int32 {
	return s.AgentVersion
}

func (s *UpdateAgentResponseBodyAgent) GetExpertSuites() []*UpdateAgentResponseBodyAgentExpertSuites {
	return s.ExpertSuites
}

func (s *UpdateAgentResponseBodyAgent) GetGmtCreateTime() *int64 {
	return s.GmtCreateTime
}

func (s *UpdateAgentResponseBodyAgent) GetGmtModifiedTime() *int64 {
	return s.GmtModifiedTime
}

func (s *UpdateAgentResponseBodyAgent) GetName() *string {
	return s.Name
}

func (s *UpdateAgentResponseBodyAgent) GetSkills() []*UpdateAgentResponseBodyAgentSkills {
	return s.Skills
}

func (s *UpdateAgentResponseBodyAgent) SetAgentVersion(v int32) *UpdateAgentResponseBodyAgent {
	s.AgentVersion = &v
	return s
}

func (s *UpdateAgentResponseBodyAgent) SetExpertSuites(v []*UpdateAgentResponseBodyAgentExpertSuites) *UpdateAgentResponseBodyAgent {
	s.ExpertSuites = v
	return s
}

func (s *UpdateAgentResponseBodyAgent) SetGmtCreateTime(v int64) *UpdateAgentResponseBodyAgent {
	s.GmtCreateTime = &v
	return s
}

func (s *UpdateAgentResponseBodyAgent) SetGmtModifiedTime(v int64) *UpdateAgentResponseBodyAgent {
	s.GmtModifiedTime = &v
	return s
}

func (s *UpdateAgentResponseBodyAgent) SetName(v string) *UpdateAgentResponseBodyAgent {
	s.Name = &v
	return s
}

func (s *UpdateAgentResponseBodyAgent) SetSkills(v []*UpdateAgentResponseBodyAgentSkills) *UpdateAgentResponseBodyAgent {
	s.Skills = v
	return s
}

func (s *UpdateAgentResponseBodyAgent) Validate() error {
	if s.ExpertSuites != nil {
		for _, item := range s.ExpertSuites {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Skills != nil {
		for _, item := range s.Skills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateAgentResponseBodyAgentExpertSuites struct {
	// The suite description of the bound version.
	//
	// example:
	//
	// Collection of data analytics expert capabilities
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the bound version.
	//
	// example:
	//
	// Data analytics expert suite
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The suite name of the bound version.
	//
	// example:
	//
	// analysis-suite
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The semantic version of the bound version.
	//
	// example:
	//
	// 1.2.0
	SemanticVersion *string `json:"SemanticVersion,omitempty" xml:"SemanticVersion,omitempty"`
	// The source of the expert suite.
	//
	// example:
	//
	// OFFICIAL
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the expert suite.
	//
	// example:
	//
	// suite_example
	SuiteId *string `json:"SuiteId,omitempty" xml:"SuiteId,omitempty"`
	// The bound published version number.
	//
	// example:
	//
	// 3
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
	// The version selection expression.
	//
	// example:
	//
	// ^1.2.3
	VersionSelector *string `json:"VersionSelector,omitempty" xml:"VersionSelector,omitempty"`
}

func (s UpdateAgentResponseBodyAgentExpertSuites) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentResponseBodyAgentExpertSuites) GoString() string {
	return s.String()
}

func (s *UpdateAgentResponseBodyAgentExpertSuites) GetDescription() *string {
	return s.Description
}

func (s *UpdateAgentResponseBodyAgentExpertSuites) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateAgentResponseBodyAgentExpertSuites) GetName() *string {
	return s.Name
}

func (s *UpdateAgentResponseBodyAgentExpertSuites) GetSemanticVersion() *string {
	return s.SemanticVersion
}

func (s *UpdateAgentResponseBodyAgentExpertSuites) GetSource() *string {
	return s.Source
}

func (s *UpdateAgentResponseBodyAgentExpertSuites) GetSuiteId() *string {
	return s.SuiteId
}

func (s *UpdateAgentResponseBodyAgentExpertSuites) GetVersion() *int32 {
	return s.Version
}

func (s *UpdateAgentResponseBodyAgentExpertSuites) GetVersionSelector() *string {
	return s.VersionSelector
}

func (s *UpdateAgentResponseBodyAgentExpertSuites) SetDescription(v string) *UpdateAgentResponseBodyAgentExpertSuites {
	s.Description = &v
	return s
}

func (s *UpdateAgentResponseBodyAgentExpertSuites) SetDisplayName(v string) *UpdateAgentResponseBodyAgentExpertSuites {
	s.DisplayName = &v
	return s
}

func (s *UpdateAgentResponseBodyAgentExpertSuites) SetName(v string) *UpdateAgentResponseBodyAgentExpertSuites {
	s.Name = &v
	return s
}

func (s *UpdateAgentResponseBodyAgentExpertSuites) SetSemanticVersion(v string) *UpdateAgentResponseBodyAgentExpertSuites {
	s.SemanticVersion = &v
	return s
}

func (s *UpdateAgentResponseBodyAgentExpertSuites) SetSource(v string) *UpdateAgentResponseBodyAgentExpertSuites {
	s.Source = &v
	return s
}

func (s *UpdateAgentResponseBodyAgentExpertSuites) SetSuiteId(v string) *UpdateAgentResponseBodyAgentExpertSuites {
	s.SuiteId = &v
	return s
}

func (s *UpdateAgentResponseBodyAgentExpertSuites) SetVersion(v int32) *UpdateAgentResponseBodyAgentExpertSuites {
	s.Version = &v
	return s
}

func (s *UpdateAgentResponseBodyAgentExpertSuites) SetVersionSelector(v string) *UpdateAgentResponseBodyAgentExpertSuites {
	s.VersionSelector = &v
	return s
}

func (s *UpdateAgentResponseBodyAgentExpertSuites) Validate() error {
	return dara.Validate(s)
}

type UpdateAgentResponseBodyAgentSkills struct {
	// The name of the bound custom skill.
	//
	// example:
	//
	// my-business-skill
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the official skill.
	//
	// example:
	//
	// dataworks-semantic
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The source of the skill.
	//
	// example:
	//
	// OFFICIAL
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The published version number of the bound custom skill.
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
	// The version selection expression.
	//
	// example:
	//
	// ^1.2.3
	VersionSelector *string `json:"VersionSelector,omitempty" xml:"VersionSelector,omitempty"`
}

func (s UpdateAgentResponseBodyAgentSkills) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentResponseBodyAgentSkills) GoString() string {
	return s.String()
}

func (s *UpdateAgentResponseBodyAgentSkills) GetName() *string {
	return s.Name
}

func (s *UpdateAgentResponseBodyAgentSkills) GetSkillId() *string {
	return s.SkillId
}

func (s *UpdateAgentResponseBodyAgentSkills) GetSource() *string {
	return s.Source
}

func (s *UpdateAgentResponseBodyAgentSkills) GetVersion() *int32 {
	return s.Version
}

func (s *UpdateAgentResponseBodyAgentSkills) GetVersionSelector() *string {
	return s.VersionSelector
}

func (s *UpdateAgentResponseBodyAgentSkills) SetName(v string) *UpdateAgentResponseBodyAgentSkills {
	s.Name = &v
	return s
}

func (s *UpdateAgentResponseBodyAgentSkills) SetSkillId(v string) *UpdateAgentResponseBodyAgentSkills {
	s.SkillId = &v
	return s
}

func (s *UpdateAgentResponseBodyAgentSkills) SetSource(v string) *UpdateAgentResponseBodyAgentSkills {
	s.Source = &v
	return s
}

func (s *UpdateAgentResponseBodyAgentSkills) SetVersion(v int32) *UpdateAgentResponseBodyAgentSkills {
	s.Version = &v
	return s
}

func (s *UpdateAgentResponseBodyAgentSkills) SetVersionSelector(v string) *UpdateAgentResponseBodyAgentSkills {
	s.VersionSelector = &v
	return s
}

func (s *UpdateAgentResponseBodyAgentSkills) Validate() error {
	return dara.Validate(s)
}
