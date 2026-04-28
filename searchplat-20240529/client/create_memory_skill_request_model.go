// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemorySkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *CreateMemorySkillRequest
	GetAgentId() *string
	SetUserId(v string) *CreateMemorySkillRequest
	GetUserId() *string
	SetZipBase64(v string) *CreateMemorySkillRequest
	GetZipBase64() *string
}

type CreateMemorySkillRequest struct {
	AgentId   *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	UserId    *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	ZipBase64 *string `json:"zip_base64,omitempty" xml:"zip_base64,omitempty"`
}

func (s CreateMemorySkillRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMemorySkillRequest) GoString() string {
	return s.String()
}

func (s *CreateMemorySkillRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateMemorySkillRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateMemorySkillRequest) GetZipBase64() *string {
	return s.ZipBase64
}

func (s *CreateMemorySkillRequest) SetAgentId(v string) *CreateMemorySkillRequest {
	s.AgentId = &v
	return s
}

func (s *CreateMemorySkillRequest) SetUserId(v string) *CreateMemorySkillRequest {
	s.UserId = &v
	return s
}

func (s *CreateMemorySkillRequest) SetZipBase64(v string) *CreateMemorySkillRequest {
	s.ZipBase64 = &v
	return s
}

func (s *CreateMemorySkillRequest) Validate() error {
	return dara.Validate(s)
}
