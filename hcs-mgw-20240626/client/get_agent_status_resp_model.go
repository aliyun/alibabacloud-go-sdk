// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentStatusResp interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIP(v string) *GetAgentStatusResp
	GetAgentIP() *string
	SetAgentVersion(v string) *GetAgentStatusResp
	GetAgentVersion() *string
	SetStatus(v string) *GetAgentStatusResp
	GetStatus() *string
}

type GetAgentStatusResp struct {
	// example:
	//
	// 192.168.0.2
	AgentIP *string `json:"AgentIP,omitempty" xml:"AgentIP,omitempty"`
	// example:
	//
	// 1.5.0
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAgentStatusResp) String() string {
	return dara.Prettify(s)
}

func (s GetAgentStatusResp) GoString() string {
	return s.String()
}

func (s *GetAgentStatusResp) GetAgentIP() *string {
	return s.AgentIP
}

func (s *GetAgentStatusResp) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *GetAgentStatusResp) GetStatus() *string {
	return s.Status
}

func (s *GetAgentStatusResp) SetAgentIP(v string) *GetAgentStatusResp {
	s.AgentIP = &v
	return s
}

func (s *GetAgentStatusResp) SetAgentVersion(v string) *GetAgentStatusResp {
	s.AgentVersion = &v
	return s
}

func (s *GetAgentStatusResp) SetStatus(v string) *GetAgentStatusResp {
	s.Status = &v
	return s
}

func (s *GetAgentStatusResp) Validate() error {
	return dara.Validate(s)
}
