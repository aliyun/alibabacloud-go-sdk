// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDasAgentSSERequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *GetDasAgentSSERequest
	GetAgentId() *string
	SetInstanceId(v string) *GetDasAgentSSERequest
	GetInstanceId() *string
	SetQuery(v string) *GetDasAgentSSERequest
	GetQuery() *string
	SetSessionId(v string) *GetDasAgentSSERequest
	GetSessionId() *string
}

type GetDasAgentSSERequest struct {
	// example:
	//
	// ag-472T0DxtmjIxxxxx
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// rm-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Are there any issues or abnormalities with my instance rm-xxx?
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// xxx-xxx-xxx
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetDasAgentSSERequest) String() string {
	return dara.Prettify(s)
}

func (s GetDasAgentSSERequest) GoString() string {
	return s.String()
}

func (s *GetDasAgentSSERequest) GetAgentId() *string {
	return s.AgentId
}

func (s *GetDasAgentSSERequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDasAgentSSERequest) GetQuery() *string {
	return s.Query
}

func (s *GetDasAgentSSERequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetDasAgentSSERequest) SetAgentId(v string) *GetDasAgentSSERequest {
	s.AgentId = &v
	return s
}

func (s *GetDasAgentSSERequest) SetInstanceId(v string) *GetDasAgentSSERequest {
	s.InstanceId = &v
	return s
}

func (s *GetDasAgentSSERequest) SetQuery(v string) *GetDasAgentSSERequest {
	s.Query = &v
	return s
}

func (s *GetDasAgentSSERequest) SetSessionId(v string) *GetDasAgentSSERequest {
	s.SessionId = &v
	return s
}

func (s *GetDasAgentSSERequest) Validate() error {
	return dara.Validate(s)
}
