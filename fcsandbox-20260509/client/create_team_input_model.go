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
	SetResourceGroupID(v string) *CreateTeamInput
	GetResourceGroupID() *string
	SetTeamName(v string) *CreateTeamInput
	GetTeamName() *string
}

type CreateTeamInput struct {
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
	TeamName        *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
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
