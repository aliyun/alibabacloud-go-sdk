// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *GetAgentRequest
	GetXDebugId() *string
	SetAgentId(v string) *GetAgentRequest
	GetAgentId() *string
	SetXSysomInvokeSource(v string) *GetAgentRequest
	GetXSysomInvokeSource() *string
}

type GetAgentRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The component ID.
	//
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	AgentId            *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s GetAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentRequest) GoString() string {
	return s.String()
}

func (s *GetAgentRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *GetAgentRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *GetAgentRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *GetAgentRequest) SetXDebugId(v string) *GetAgentRequest {
	s.XDebugId = &v
	return s
}

func (s *GetAgentRequest) SetAgentId(v string) *GetAgentRequest {
	s.AgentId = &v
	return s
}

func (s *GetAgentRequest) SetXSysomInvokeSource(v string) *GetAgentRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *GetAgentRequest) Validate() error {
	return dara.Validate(s)
}
