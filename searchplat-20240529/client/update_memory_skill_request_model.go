// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemorySkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UpdateMemorySkillRequest
	GetAgentId() *string
	SetFiles(v map[string]*string) *UpdateMemorySkillRequest
	GetFiles() map[string]*string
	SetName(v string) *UpdateMemorySkillRequest
	GetName() *string
	SetUserId(v string) *UpdateMemorySkillRequest
	GetUserId() *string
	SetVersion(v string) *UpdateMemorySkillRequest
	GetVersion() *string
}

type UpdateMemorySkillRequest struct {
	AgentId *string            `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	Files   map[string]*string `json:"files,omitempty" xml:"files,omitempty"`
	Name    *string            `json:"name,omitempty" xml:"name,omitempty"`
	UserId  *string            `json:"user_id,omitempty" xml:"user_id,omitempty"`
	Version *string            `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpdateMemorySkillRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemorySkillRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemorySkillRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdateMemorySkillRequest) GetFiles() map[string]*string {
	return s.Files
}

func (s *UpdateMemorySkillRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMemorySkillRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateMemorySkillRequest) GetVersion() *string {
	return s.Version
}

func (s *UpdateMemorySkillRequest) SetAgentId(v string) *UpdateMemorySkillRequest {
	s.AgentId = &v
	return s
}

func (s *UpdateMemorySkillRequest) SetFiles(v map[string]*string) *UpdateMemorySkillRequest {
	s.Files = v
	return s
}

func (s *UpdateMemorySkillRequest) SetName(v string) *UpdateMemorySkillRequest {
	s.Name = &v
	return s
}

func (s *UpdateMemorySkillRequest) SetUserId(v string) *UpdateMemorySkillRequest {
	s.UserId = &v
	return s
}

func (s *UpdateMemorySkillRequest) SetVersion(v string) *UpdateMemorySkillRequest {
	s.Version = &v
	return s
}

func (s *UpdateMemorySkillRequest) Validate() error {
	return dara.Validate(s)
}
