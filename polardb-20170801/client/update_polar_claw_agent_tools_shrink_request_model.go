// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarClawAgentToolsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UpdatePolarClawAgentToolsShrinkRequest
	GetAgentId() *string
	SetAllowShrink(v string) *UpdatePolarClawAgentToolsShrinkRequest
	GetAllowShrink() *string
	SetAlsoAllowShrink(v string) *UpdatePolarClawAgentToolsShrinkRequest
	GetAlsoAllowShrink() *string
	SetApplicationId(v string) *UpdatePolarClawAgentToolsShrinkRequest
	GetApplicationId() *string
	SetDenyShrink(v string) *UpdatePolarClawAgentToolsShrinkRequest
	GetDenyShrink() *string
	SetProfile(v string) *UpdatePolarClawAgentToolsShrinkRequest
	GetProfile() *string
}

type UpdatePolarClawAgentToolsShrinkRequest struct {
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
	AllowShrink *string `json:"Allow,omitempty" xml:"Allow,omitempty"`
	// The list of additionally allowed tools.
	//
	// example:
	//
	// ["send_message"]
	AlsoAllowShrink *string `json:"AlsoAllow,omitempty" xml:"AlsoAllow,omitempty"`
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
	DenyShrink *string `json:"Deny,omitempty" xml:"Deny,omitempty"`
	// The tool profile.
	//
	// example:
	//
	// coding
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
}

func (s UpdatePolarClawAgentToolsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawAgentToolsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawAgentToolsShrinkRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdatePolarClawAgentToolsShrinkRequest) GetAllowShrink() *string {
	return s.AllowShrink
}

func (s *UpdatePolarClawAgentToolsShrinkRequest) GetAlsoAllowShrink() *string {
	return s.AlsoAllowShrink
}

func (s *UpdatePolarClawAgentToolsShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdatePolarClawAgentToolsShrinkRequest) GetDenyShrink() *string {
	return s.DenyShrink
}

func (s *UpdatePolarClawAgentToolsShrinkRequest) GetProfile() *string {
	return s.Profile
}

func (s *UpdatePolarClawAgentToolsShrinkRequest) SetAgentId(v string) *UpdatePolarClawAgentToolsShrinkRequest {
	s.AgentId = &v
	return s
}

func (s *UpdatePolarClawAgentToolsShrinkRequest) SetAllowShrink(v string) *UpdatePolarClawAgentToolsShrinkRequest {
	s.AllowShrink = &v
	return s
}

func (s *UpdatePolarClawAgentToolsShrinkRequest) SetAlsoAllowShrink(v string) *UpdatePolarClawAgentToolsShrinkRequest {
	s.AlsoAllowShrink = &v
	return s
}

func (s *UpdatePolarClawAgentToolsShrinkRequest) SetApplicationId(v string) *UpdatePolarClawAgentToolsShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePolarClawAgentToolsShrinkRequest) SetDenyShrink(v string) *UpdatePolarClawAgentToolsShrinkRequest {
	s.DenyShrink = &v
	return s
}

func (s *UpdatePolarClawAgentToolsShrinkRequest) SetProfile(v string) *UpdatePolarClawAgentToolsShrinkRequest {
	s.Profile = &v
	return s
}

func (s *UpdatePolarClawAgentToolsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
