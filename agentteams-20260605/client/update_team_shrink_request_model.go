// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTeamShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateTeamShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateTeamShrinkRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateTeamShrinkRequest
	GetInstanceId() *string
	SetName(v string) *UpdateTeamShrinkRequest
	GetName() *string
	SetTeamMembersShrink(v string) *UpdateTeamShrinkRequest
	GetTeamMembersShrink() *string
}

type UpdateTeamShrinkRequest struct {
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TeamMembersShrink *string `json:"TeamMembers,omitempty" xml:"TeamMembers,omitempty"`
}

func (s UpdateTeamShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTeamShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTeamShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTeamShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTeamShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateTeamShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTeamShrinkRequest) GetTeamMembersShrink() *string {
	return s.TeamMembersShrink
}

func (s *UpdateTeamShrinkRequest) SetClientToken(v string) *UpdateTeamShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTeamShrinkRequest) SetDescription(v string) *UpdateTeamShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateTeamShrinkRequest) SetInstanceId(v string) *UpdateTeamShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateTeamShrinkRequest) SetName(v string) *UpdateTeamShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateTeamShrinkRequest) SetTeamMembersShrink(v string) *UpdateTeamShrinkRequest {
	s.TeamMembersShrink = &v
	return s
}

func (s *UpdateTeamShrinkRequest) Validate() error {
	return dara.Validate(s)
}
