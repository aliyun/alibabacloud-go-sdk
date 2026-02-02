// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *GetCustomAgentResponseBody
	GetCreatedAt() *string
	SetEnableTools(v bool) *GetCustomAgentResponseBody
	GetEnableTools() *bool
	SetId(v string) *GetCustomAgentResponseBody
	GetId() *string
	SetName(v string) *GetCustomAgentResponseBody
	GetName() *string
	SetRequestId(v string) *GetCustomAgentResponseBody
	GetRequestId() *string
	SetSkills(v []*GetCustomAgentResponseBodySkills) *GetCustomAgentResponseBody
	GetSkills() []*GetCustomAgentResponseBodySkills
	SetSystemPrompt(v string) *GetCustomAgentResponseBody
	GetSystemPrompt() *string
	SetTools(v []*string) *GetCustomAgentResponseBody
	GetTools() []*string
	SetUpdatedAt(v string) *GetCustomAgentResponseBody
	GetUpdatedAt() *string
}

type GetCustomAgentResponseBody struct {
	// The creation time of the agent.
	//
	// example:
	//
	// 2025-06-04T02:25:43Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// Indicates whether tools are enabled.
	//
	// example:
	//
	// true
	EnableTools *bool `json:"EnableTools,omitempty" xml:"EnableTools,omitempty"`
	// The ID of the backup set.
	//
	// example:
	//
	// 17053
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the dedicated agent.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Skills    []*GetCustomAgentResponseBodySkills `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	// The system prompts.
	SystemPrompt *string `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	// The details of the tools.
	Tools []*string `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
	// The modification time of the agent.
	//
	// example:
	//
	// 2020-11-27 16:02:28
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s GetCustomAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetCustomAgentResponseBody) GetEnableTools() *bool {
	return s.EnableTools
}

func (s *GetCustomAgentResponseBody) GetId() *string {
	return s.Id
}

func (s *GetCustomAgentResponseBody) GetName() *string {
	return s.Name
}

func (s *GetCustomAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomAgentResponseBody) GetSkills() []*GetCustomAgentResponseBodySkills {
	return s.Skills
}

func (s *GetCustomAgentResponseBody) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *GetCustomAgentResponseBody) GetTools() []*string {
	return s.Tools
}

func (s *GetCustomAgentResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetCustomAgentResponseBody) SetCreatedAt(v string) *GetCustomAgentResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetCustomAgentResponseBody) SetEnableTools(v bool) *GetCustomAgentResponseBody {
	s.EnableTools = &v
	return s
}

func (s *GetCustomAgentResponseBody) SetId(v string) *GetCustomAgentResponseBody {
	s.Id = &v
	return s
}

func (s *GetCustomAgentResponseBody) SetName(v string) *GetCustomAgentResponseBody {
	s.Name = &v
	return s
}

func (s *GetCustomAgentResponseBody) SetRequestId(v string) *GetCustomAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomAgentResponseBody) SetSkills(v []*GetCustomAgentResponseBodySkills) *GetCustomAgentResponseBody {
	s.Skills = v
	return s
}

func (s *GetCustomAgentResponseBody) SetSystemPrompt(v string) *GetCustomAgentResponseBody {
	s.SystemPrompt = &v
	return s
}

func (s *GetCustomAgentResponseBody) SetTools(v []*string) *GetCustomAgentResponseBody {
	s.Tools = v
	return s
}

func (s *GetCustomAgentResponseBody) SetUpdatedAt(v string) *GetCustomAgentResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *GetCustomAgentResponseBody) Validate() error {
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

type GetCustomAgentResponseBodySkills struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SkillType   *string `json:"SkillType,omitempty" xml:"SkillType,omitempty"`
}

func (s GetCustomAgentResponseBodySkills) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentResponseBodySkills) GoString() string {
	return s.String()
}

func (s *GetCustomAgentResponseBodySkills) GetDescription() *string {
	return s.Description
}

func (s *GetCustomAgentResponseBodySkills) GetId() *string {
	return s.Id
}

func (s *GetCustomAgentResponseBodySkills) GetName() *string {
	return s.Name
}

func (s *GetCustomAgentResponseBodySkills) GetSkillType() *string {
	return s.SkillType
}

func (s *GetCustomAgentResponseBodySkills) SetDescription(v string) *GetCustomAgentResponseBodySkills {
	s.Description = &v
	return s
}

func (s *GetCustomAgentResponseBodySkills) SetId(v string) *GetCustomAgentResponseBodySkills {
	s.Id = &v
	return s
}

func (s *GetCustomAgentResponseBodySkills) SetName(v string) *GetCustomAgentResponseBodySkills {
	s.Name = &v
	return s
}

func (s *GetCustomAgentResponseBodySkills) SetSkillType(v string) *GetCustomAgentResponseBodySkills {
	s.SkillType = &v
	return s
}

func (s *GetCustomAgentResponseBodySkills) Validate() error {
	return dara.Validate(s)
}
