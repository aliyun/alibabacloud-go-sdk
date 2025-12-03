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
	SetSystemPrompt(v string) *GetCustomAgentResponseBody
	GetSystemPrompt() *string
	SetTools(v []*string) *GetCustomAgentResponseBody
	GetTools() []*string
	SetUpdatedAt(v string) *GetCustomAgentResponseBody
	GetUpdatedAt() *string
}

type GetCustomAgentResponseBody struct {
	// example:
	//
	// 2025-06-04T02:25:43Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// true
	EnableTools *bool `json:"EnableTools,omitempty" xml:"EnableTools,omitempty"`
	// example:
	//
	// 17053
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId    *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SystemPrompt *string   `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	Tools        []*string `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}
