// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *CreateCustomAgentResponseBody
	GetCreatedAt() *string
	SetEnableTools(v bool) *CreateCustomAgentResponseBody
	GetEnableTools() *bool
	SetId(v string) *CreateCustomAgentResponseBody
	GetId() *string
	SetName(v string) *CreateCustomAgentResponseBody
	GetName() *string
	SetRequestId(v string) *CreateCustomAgentResponseBody
	GetRequestId() *string
	SetSystemPrompt(v string) *CreateCustomAgentResponseBody
	GetSystemPrompt() *string
	SetTools(v []*string) *CreateCustomAgentResponseBody
	GetTools() []*string
}

type CreateCustomAgentResponseBody struct {
	// example:
	//
	// 2020-11-27 16:01:28
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// true
	EnableTools *bool `json:"EnableTools,omitempty" xml:"EnableTools,omitempty"`
	// AgentId
	//
	// example:
	//
	// d1b7d639-f34e-44c7-8231-987da14d****
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId    *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SystemPrompt *string   `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	Tools        []*string `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
}

func (s CreateCustomAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateCustomAgentResponseBody) GetEnableTools() *bool {
	return s.EnableTools
}

func (s *CreateCustomAgentResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateCustomAgentResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateCustomAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomAgentResponseBody) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *CreateCustomAgentResponseBody) GetTools() []*string {
	return s.Tools
}

func (s *CreateCustomAgentResponseBody) SetCreatedAt(v string) *CreateCustomAgentResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateCustomAgentResponseBody) SetEnableTools(v bool) *CreateCustomAgentResponseBody {
	s.EnableTools = &v
	return s
}

func (s *CreateCustomAgentResponseBody) SetId(v string) *CreateCustomAgentResponseBody {
	s.Id = &v
	return s
}

func (s *CreateCustomAgentResponseBody) SetName(v string) *CreateCustomAgentResponseBody {
	s.Name = &v
	return s
}

func (s *CreateCustomAgentResponseBody) SetRequestId(v string) *CreateCustomAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomAgentResponseBody) SetSystemPrompt(v string) *CreateCustomAgentResponseBody {
	s.SystemPrompt = &v
	return s
}

func (s *CreateCustomAgentResponseBody) SetTools(v []*string) *CreateCustomAgentResponseBody {
	s.Tools = v
	return s
}

func (s *CreateCustomAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
