// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnableTools(v string) *UpdateCustomAgentResponseBody
	GetEnableTools() *string
	SetId(v string) *UpdateCustomAgentResponseBody
	GetId() *string
	SetName(v string) *UpdateCustomAgentResponseBody
	GetName() *string
	SetRequestId(v string) *UpdateCustomAgentResponseBody
	GetRequestId() *string
	SetSystemPrompt(v string) *UpdateCustomAgentResponseBody
	GetSystemPrompt() *string
	SetTools(v []*string) *UpdateCustomAgentResponseBody
	GetTools() []*string
}

type UpdateCustomAgentResponseBody struct {
	// example:
	//
	// true
	EnableTools *string `json:"EnableTools,omitempty" xml:"EnableTools,omitempty"`
	// AgentId。
	//
	// example:
	//
	// 82cf3d62-0add-47bd-869f-877131f7****
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId    *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SystemPrompt *string   `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	Tools        []*string `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
}

func (s UpdateCustomAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponseBody) GetEnableTools() *string {
	return s.EnableTools
}

func (s *UpdateCustomAgentResponseBody) GetId() *string {
	return s.Id
}

func (s *UpdateCustomAgentResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateCustomAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomAgentResponseBody) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *UpdateCustomAgentResponseBody) GetTools() []*string {
	return s.Tools
}

func (s *UpdateCustomAgentResponseBody) SetEnableTools(v string) *UpdateCustomAgentResponseBody {
	s.EnableTools = &v
	return s
}

func (s *UpdateCustomAgentResponseBody) SetId(v string) *UpdateCustomAgentResponseBody {
	s.Id = &v
	return s
}

func (s *UpdateCustomAgentResponseBody) SetName(v string) *UpdateCustomAgentResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateCustomAgentResponseBody) SetRequestId(v string) *UpdateCustomAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomAgentResponseBody) SetSystemPrompt(v string) *UpdateCustomAgentResponseBody {
	s.SystemPrompt = &v
	return s
}

func (s *UpdateCustomAgentResponseBody) SetTools(v []*string) *UpdateCustomAgentResponseBody {
	s.Tools = v
	return s
}

func (s *UpdateCustomAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
