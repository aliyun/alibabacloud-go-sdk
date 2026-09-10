// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLhmAgentStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentType(v int32) *GetLhmAgentStatusRequest
	GetAgentType() *int32
	SetSkillName(v string) *GetLhmAgentStatusRequest
	GetSkillName() *string
}

type GetLhmAgentStatusRequest struct {
	// The Agent type. Valid values:
	//
	// - 0: data validation (the only type currently supported).
	//
	// - 1: metadata.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	AgentType *int32 `json:"agentType,omitempty" xml:"agentType,omitempty"`
	// The skill name. This parameter is optional.
	//
	// example:
	//
	// lhm-data-validation-skill
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
}

func (s GetLhmAgentStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLhmAgentStatusRequest) GoString() string {
	return s.String()
}

func (s *GetLhmAgentStatusRequest) GetAgentType() *int32 {
	return s.AgentType
}

func (s *GetLhmAgentStatusRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *GetLhmAgentStatusRequest) SetAgentType(v int32) *GetLhmAgentStatusRequest {
	s.AgentType = &v
	return s
}

func (s *GetLhmAgentStatusRequest) SetSkillName(v string) *GetLhmAgentStatusRequest {
	s.SkillName = &v
	return s
}

func (s *GetLhmAgentStatusRequest) Validate() error {
	return dara.Validate(s)
}
