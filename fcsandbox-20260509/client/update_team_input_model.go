// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTeamInput interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateTeamInput
	GetDescription() *string
	SetResourceGroupID(v string) *UpdateTeamInput
	GetResourceGroupID() *string
	SetTeamName(v string) *UpdateTeamInput
	GetTeamName() *string
}

type UpdateTeamInput struct {
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
	TeamName        *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
}

func (s UpdateTeamInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateTeamInput) GoString() string {
	return s.String()
}

func (s *UpdateTeamInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateTeamInput) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *UpdateTeamInput) GetTeamName() *string {
	return s.TeamName
}

func (s *UpdateTeamInput) SetDescription(v string) *UpdateTeamInput {
	s.Description = &v
	return s
}

func (s *UpdateTeamInput) SetResourceGroupID(v string) *UpdateTeamInput {
	s.ResourceGroupID = &v
	return s
}

func (s *UpdateTeamInput) SetTeamName(v string) *UpdateTeamInput {
	s.TeamName = &v
	return s
}

func (s *UpdateTeamInput) Validate() error {
	return dara.Validate(s)
}
