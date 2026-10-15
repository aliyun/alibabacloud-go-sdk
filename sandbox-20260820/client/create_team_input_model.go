// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTeamInput interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateTeamInput
	GetDescription() *string
	SetPlan(v string) *CreateTeamInput
	GetPlan() *string
	SetResourceGroupID(v string) *CreateTeamInput
	GetResourceGroupID() *string
	SetTeamName(v string) *CreateTeamInput
	GetTeamName() *string
}

type CreateTeamInput struct {
	// The description of the team.
	//
	// example:
	//
	// Sandbox environment for the algorithm team
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The plan specification of the team. Valid values: eco (economy), std (standard), pro (professional), ent (enterprise).
	//
	// example:
	//
	// std
	Plan *string `json:"plan,omitempty" xml:"plan,omitempty"`
	// The ID of the Alibaba Cloud resource group to which the team belongs. If this parameter is not specified, the default resource group of the account is used.
	//
	// example:
	//
	// rg-****
	ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
	// The name of the team.
	//
	// example:
	//
	// sandbox-dev
	TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
}

func (s CreateTeamInput) String() string {
	return dara.Prettify(s)
}

func (s CreateTeamInput) GoString() string {
	return s.String()
}

func (s *CreateTeamInput) GetDescription() *string {
	return s.Description
}

func (s *CreateTeamInput) GetPlan() *string {
	return s.Plan
}

func (s *CreateTeamInput) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *CreateTeamInput) GetTeamName() *string {
	return s.TeamName
}

func (s *CreateTeamInput) SetDescription(v string) *CreateTeamInput {
	s.Description = &v
	return s
}

func (s *CreateTeamInput) SetPlan(v string) *CreateTeamInput {
	s.Plan = &v
	return s
}

func (s *CreateTeamInput) SetResourceGroupID(v string) *CreateTeamInput {
	s.ResourceGroupID = &v
	return s
}

func (s *CreateTeamInput) SetTeamName(v string) *CreateTeamInput {
	s.TeamName = &v
	return s
}

func (s *CreateTeamInput) Validate() error {
	return dara.Validate(s)
}
