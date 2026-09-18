// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UpdateAgentResponseBody
	GetAgentId() *string
	SetName(v string) *UpdateAgentResponseBody
	GetName() *string
	SetRequestId(v string) *UpdateAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAgentResponseBody
	GetSuccess() *bool
	SetUpdatedAt(v int64) *UpdateAgentResponseBody
	GetUpdatedAt() *int64
}

type UpdateAgentResponseBody struct {
	// Agent ID。
	//
	// example:
	//
	// agent_00000000000000000000000000000001
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The Agent name.
	//
	// example:
	//
	// code-review-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID, used for Tracing Analysis and troubleshooting.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the Agent was successfully updated. A successful response always returns `true`. A failure returns an error response.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The most recent update time, as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1788332700000
	UpdatedAt *int64 `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s UpdateAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdateAgentResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAgentResponseBody) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *UpdateAgentResponseBody) SetAgentId(v string) *UpdateAgentResponseBody {
	s.AgentId = &v
	return s
}

func (s *UpdateAgentResponseBody) SetName(v string) *UpdateAgentResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateAgentResponseBody) SetRequestId(v string) *UpdateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgentResponseBody) SetSuccess(v bool) *UpdateAgentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAgentResponseBody) SetUpdatedAt(v int64) *UpdateAgentResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
