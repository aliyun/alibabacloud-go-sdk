// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgent(v *CreateAgentResponseBodyAgent) *CreateAgentResponseBody
	GetAgent() *CreateAgentResponseBodyAgent
	SetRequestId(v string) *CreateAgentResponseBody
	GetRequestId() *string
}

type CreateAgentResponseBody struct {
	// The details of the created agent.
	Agent *CreateAgentResponseBodyAgent `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 824F80BA-1778-5D8A-BAFF-668A4D9C4CC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentResponseBody) GetAgent() *CreateAgentResponseBodyAgent {
	return s.Agent
}

func (s *CreateAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentResponseBody) SetAgent(v *CreateAgentResponseBodyAgent) *CreateAgentResponseBody {
	s.Agent = v
	return s
}

func (s *CreateAgentResponseBody) SetRequestId(v string) *CreateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentResponseBody) Validate() error {
	if s.Agent != nil {
		if err := s.Agent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentResponseBodyAgent struct {
	// 已保存的套件引用。官方资源返回来源及目录 ID；发布后的制品版本由 PackageVersion 表示。
	ExpertSuites []*CreateAgentResponseBodyAgentExpertSuites `json:"ExpertSuites,omitempty" xml:"ExpertSuites,omitempty" type:"Repeated"`
	// The creation time. The value is a UNIX timestamp in milliseconds.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1780555634000
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The last modification time. The value is a UNIX timestamp in milliseconds.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1780555634000
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The name of the agent.
	//
	// example:
	//
	// my-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 已保存的 Skill 引用；官方 Skill 在 Agent 发布后返回固定的 PackageVersion。
	Skills []*CreateAgentResponseBodyAgentSkills `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
}

func (s CreateAgentResponseBodyAgent) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentResponseBodyAgent) GoString() string {
	return s.String()
}

func (s *CreateAgentResponseBodyAgent) GetExpertSuites() []*CreateAgentResponseBodyAgentExpertSuites {
	return s.ExpertSuites
}

func (s *CreateAgentResponseBodyAgent) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *CreateAgentResponseBodyAgent) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *CreateAgentResponseBodyAgent) GetName() *string {
	return s.Name
}

func (s *CreateAgentResponseBodyAgent) GetSkills() []*CreateAgentResponseBodyAgentSkills {
	return s.Skills
}

func (s *CreateAgentResponseBodyAgent) SetExpertSuites(v []*CreateAgentResponseBodyAgentExpertSuites) *CreateAgentResponseBodyAgent {
	s.ExpertSuites = v
	return s
}

func (s *CreateAgentResponseBodyAgent) SetGmtCreateTime(v string) *CreateAgentResponseBodyAgent {
	s.GmtCreateTime = &v
	return s
}

func (s *CreateAgentResponseBodyAgent) SetGmtModifiedTime(v string) *CreateAgentResponseBodyAgent {
	s.GmtModifiedTime = &v
	return s
}

func (s *CreateAgentResponseBodyAgent) SetName(v string) *CreateAgentResponseBodyAgent {
	s.Name = &v
	return s
}

func (s *CreateAgentResponseBodyAgent) SetSkills(v []*CreateAgentResponseBodyAgentSkills) *CreateAgentResponseBodyAgent {
	s.Skills = v
	return s
}

func (s *CreateAgentResponseBodyAgent) Validate() error {
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

type CreateAgentResponseBodyAgentExpertSuites struct {
	// 创建时不返回；请使用 GetAgent 显式查询已发布版本获取固定的官方套件制品版本。
	//
	// example:
	//
	// 0.1.0.202610211408
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// example:
	//
	// OFFICIAL
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// dataworks-data-analysis-extension
	SuiteId *string `json:"SuiteId,omitempty" xml:"SuiteId,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateAgentResponseBodyAgentExpertSuites) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentResponseBodyAgentExpertSuites) GoString() string {
	return s.String()
}

func (s *CreateAgentResponseBodyAgentExpertSuites) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *CreateAgentResponseBodyAgentExpertSuites) GetSource() *string {
	return s.Source
}

func (s *CreateAgentResponseBodyAgentExpertSuites) GetSuiteId() *string {
	return s.SuiteId
}

func (s *CreateAgentResponseBodyAgentExpertSuites) GetVersion() *int32 {
	return s.Version
}

func (s *CreateAgentResponseBodyAgentExpertSuites) SetPackageVersion(v string) *CreateAgentResponseBodyAgentExpertSuites {
	s.PackageVersion = &v
	return s
}

func (s *CreateAgentResponseBodyAgentExpertSuites) SetSource(v string) *CreateAgentResponseBodyAgentExpertSuites {
	s.Source = &v
	return s
}

func (s *CreateAgentResponseBodyAgentExpertSuites) SetSuiteId(v string) *CreateAgentResponseBodyAgentExpertSuites {
	s.SuiteId = &v
	return s
}

func (s *CreateAgentResponseBodyAgentExpertSuites) SetVersion(v int32) *CreateAgentResponseBodyAgentExpertSuites {
	s.Version = &v
	return s
}

func (s *CreateAgentResponseBodyAgentExpertSuites) Validate() error {
	return dara.Validate(s)
}

type CreateAgentResponseBodyAgentSkills struct {
	// example:
	//
	// my-business-skill
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 创建时不返回；请使用 GetAgent 显式查询已发布版本获取固定的官方制品版本。
	//
	// example:
	//
	// 0.6.0.202608201110
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// example:
	//
	// dataworks-semantic
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// example:
	//
	// OFFICIAL
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateAgentResponseBodyAgentSkills) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentResponseBodyAgentSkills) GoString() string {
	return s.String()
}

func (s *CreateAgentResponseBodyAgentSkills) GetName() *string {
	return s.Name
}

func (s *CreateAgentResponseBodyAgentSkills) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *CreateAgentResponseBodyAgentSkills) GetSkillId() *string {
	return s.SkillId
}

func (s *CreateAgentResponseBodyAgentSkills) GetSource() *string {
	return s.Source
}

func (s *CreateAgentResponseBodyAgentSkills) GetVersion() *int32 {
	return s.Version
}

func (s *CreateAgentResponseBodyAgentSkills) SetName(v string) *CreateAgentResponseBodyAgentSkills {
	s.Name = &v
	return s
}

func (s *CreateAgentResponseBodyAgentSkills) SetPackageVersion(v string) *CreateAgentResponseBodyAgentSkills {
	s.PackageVersion = &v
	return s
}

func (s *CreateAgentResponseBodyAgentSkills) SetSkillId(v string) *CreateAgentResponseBodyAgentSkills {
	s.SkillId = &v
	return s
}

func (s *CreateAgentResponseBodyAgentSkills) SetSource(v string) *CreateAgentResponseBodyAgentSkills {
	s.Source = &v
	return s
}

func (s *CreateAgentResponseBodyAgentSkills) SetVersion(v int32) *CreateAgentResponseBodyAgentSkills {
	s.Version = &v
	return s
}

func (s *CreateAgentResponseBodyAgentSkills) Validate() error {
	return dara.Validate(s)
}
