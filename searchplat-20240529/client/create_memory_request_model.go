// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *CreateMemoryRequest
	GetAgentId() *string
	SetEnhancements(v map[string]interface{}) *CreateMemoryRequest
	GetEnhancements() map[string]interface{}
	SetMessages(v interface{}) *CreateMemoryRequest
	GetMessages() interface{}
	SetRunId(v string) *CreateMemoryRequest
	GetRunId() *string
	SetUserId(v string) *CreateMemoryRequest
	GetUserId() *string
}

type CreateMemoryRequest struct {
	AgentId      *string                `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	Enhancements map[string]interface{} `json:"enhancements,omitempty" xml:"enhancements,omitempty"`
	Messages     interface{}            `json:"messages,omitempty" xml:"messages,omitempty"`
	RunId        *string                `json:"run_id,omitempty" xml:"run_id,omitempty"`
	UserId       *string                `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CreateMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryRequest) GoString() string {
	return s.String()
}

func (s *CreateMemoryRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateMemoryRequest) GetEnhancements() map[string]interface{} {
	return s.Enhancements
}

func (s *CreateMemoryRequest) GetMessages() interface{} {
	return s.Messages
}

func (s *CreateMemoryRequest) GetRunId() *string {
	return s.RunId
}

func (s *CreateMemoryRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateMemoryRequest) SetAgentId(v string) *CreateMemoryRequest {
	s.AgentId = &v
	return s
}

func (s *CreateMemoryRequest) SetEnhancements(v map[string]interface{}) *CreateMemoryRequest {
	s.Enhancements = v
	return s
}

func (s *CreateMemoryRequest) SetMessages(v interface{}) *CreateMemoryRequest {
	s.Messages = v
	return s
}

func (s *CreateMemoryRequest) SetRunId(v string) *CreateMemoryRequest {
	s.RunId = &v
	return s
}

func (s *CreateMemoryRequest) SetUserId(v string) *CreateMemoryRequest {
	s.UserId = &v
	return s
}

func (s *CreateMemoryRequest) Validate() error {
	return dara.Validate(s)
}
