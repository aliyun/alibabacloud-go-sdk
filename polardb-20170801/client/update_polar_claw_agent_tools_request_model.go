// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarClawAgentToolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UpdatePolarClawAgentToolsRequest
	GetAgentId() *string
	SetAllow(v []*string) *UpdatePolarClawAgentToolsRequest
	GetAllow() []*string
	SetAlsoAllow(v []*string) *UpdatePolarClawAgentToolsRequest
	GetAlsoAllow() []*string
	SetApplicationId(v string) *UpdatePolarClawAgentToolsRequest
	GetApplicationId() *string
	SetDeny(v []*string) *UpdatePolarClawAgentToolsRequest
	GetDeny() []*string
	SetProfile(v string) *UpdatePolarClawAgentToolsRequest
	GetProfile() *string
}

type UpdatePolarClawAgentToolsRequest struct {
	// Agent ID
	//
	// This parameter is required.
	//
	// example:
	//
	// work
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The list of explicitly allowed tools.
	//
	// example:
	//
	// ["read","write","exec"]
	Allow []*string `json:"Allow,omitempty" xml:"Allow,omitempty" type:"Repeated"`
	// The list of additionally allowed tools.
	//
	// example:
	//
	// ["send_message"]
	AlsoAllow []*string `json:"AlsoAllow,omitempty" xml:"AlsoAllow,omitempty" type:"Repeated"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The list of denied tools.
	//
	// example:
	//
	// ["exec"]
	Deny []*string `json:"Deny,omitempty" xml:"Deny,omitempty" type:"Repeated"`
	// The tool profile.
	//
	// example:
	//
	// coding
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
}

func (s UpdatePolarClawAgentToolsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawAgentToolsRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawAgentToolsRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdatePolarClawAgentToolsRequest) GetAllow() []*string {
	return s.Allow
}

func (s *UpdatePolarClawAgentToolsRequest) GetAlsoAllow() []*string {
	return s.AlsoAllow
}

func (s *UpdatePolarClawAgentToolsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdatePolarClawAgentToolsRequest) GetDeny() []*string {
	return s.Deny
}

func (s *UpdatePolarClawAgentToolsRequest) GetProfile() *string {
	return s.Profile
}

func (s *UpdatePolarClawAgentToolsRequest) SetAgentId(v string) *UpdatePolarClawAgentToolsRequest {
	s.AgentId = &v
	return s
}

func (s *UpdatePolarClawAgentToolsRequest) SetAllow(v []*string) *UpdatePolarClawAgentToolsRequest {
	s.Allow = v
	return s
}

func (s *UpdatePolarClawAgentToolsRequest) SetAlsoAllow(v []*string) *UpdatePolarClawAgentToolsRequest {
	s.AlsoAllow = v
	return s
}

func (s *UpdatePolarClawAgentToolsRequest) SetApplicationId(v string) *UpdatePolarClawAgentToolsRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePolarClawAgentToolsRequest) SetDeny(v []*string) *UpdatePolarClawAgentToolsRequest {
	s.Deny = v
	return s
}

func (s *UpdatePolarClawAgentToolsRequest) SetProfile(v string) *UpdatePolarClawAgentToolsRequest {
	s.Profile = &v
	return s
}

func (s *UpdatePolarClawAgentToolsRequest) Validate() error {
	return dara.Validate(s)
}
