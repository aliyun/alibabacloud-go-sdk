// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTeamShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdminName(v string) *CreateTeamShrinkRequest
	GetAdminName() *string
	SetClientToken(v string) *CreateTeamShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *CreateTeamShrinkRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateTeamShrinkRequest
	GetInstanceId() *string
	SetName(v string) *CreateTeamShrinkRequest
	GetName() *string
	SetTeamMembersShrink(v string) *CreateTeamShrinkRequest
	GetTeamMembersShrink() *string
}

type CreateTeamShrinkRequest struct {
	AdminName         *string `json:"AdminName,omitempty" xml:"AdminName,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TeamMembersShrink *string `json:"TeamMembers,omitempty" xml:"TeamMembers,omitempty"`
}

func (s CreateTeamShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTeamShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTeamShrinkRequest) GetAdminName() *string {
	return s.AdminName
}

func (s *CreateTeamShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTeamShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTeamShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTeamShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateTeamShrinkRequest) GetTeamMembersShrink() *string {
	return s.TeamMembersShrink
}

func (s *CreateTeamShrinkRequest) SetAdminName(v string) *CreateTeamShrinkRequest {
	s.AdminName = &v
	return s
}

func (s *CreateTeamShrinkRequest) SetClientToken(v string) *CreateTeamShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTeamShrinkRequest) SetDescription(v string) *CreateTeamShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateTeamShrinkRequest) SetInstanceId(v string) *CreateTeamShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTeamShrinkRequest) SetName(v string) *CreateTeamShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateTeamShrinkRequest) SetTeamMembersShrink(v string) *CreateTeamShrinkRequest {
	s.TeamMembersShrink = &v
	return s
}

func (s *CreateTeamShrinkRequest) Validate() error {
	return dara.Validate(s)
}
