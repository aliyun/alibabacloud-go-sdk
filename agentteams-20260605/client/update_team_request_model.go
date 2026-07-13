// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTeamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateTeamRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateTeamRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateTeamRequest
	GetInstanceId() *string
	SetName(v string) *UpdateTeamRequest
	GetName() *string
	SetTeamMembers(v []*UpdateTeamRequestTeamMembers) *UpdateTeamRequest
	GetTeamMembers() []*UpdateTeamRequestTeamMembers
}

type UpdateTeamRequest struct {
	ClientToken *string                         `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string                         `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId  *string                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
	TeamMembers []*UpdateTeamRequestTeamMembers `json:"TeamMembers,omitempty" xml:"TeamMembers,omitempty" type:"Repeated"`
}

func (s UpdateTeamRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTeamRequest) GoString() string {
	return s.String()
}

func (s *UpdateTeamRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTeamRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTeamRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateTeamRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTeamRequest) GetTeamMembers() []*UpdateTeamRequestTeamMembers {
	return s.TeamMembers
}

func (s *UpdateTeamRequest) SetClientToken(v string) *UpdateTeamRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTeamRequest) SetDescription(v string) *UpdateTeamRequest {
	s.Description = &v
	return s
}

func (s *UpdateTeamRequest) SetInstanceId(v string) *UpdateTeamRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateTeamRequest) SetName(v string) *UpdateTeamRequest {
	s.Name = &v
	return s
}

func (s *UpdateTeamRequest) SetTeamMembers(v []*UpdateTeamRequestTeamMembers) *UpdateTeamRequest {
	s.TeamMembers = v
	return s
}

func (s *UpdateTeamRequest) Validate() error {
	if s.TeamMembers != nil {
		for _, item := range s.TeamMembers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateTeamRequestTeamMembers struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateTeamRequestTeamMembers) String() string {
	return dara.Prettify(s)
}

func (s UpdateTeamRequestTeamMembers) GoString() string {
	return s.String()
}

func (s *UpdateTeamRequestTeamMembers) GetName() *string {
	return s.Name
}

func (s *UpdateTeamRequestTeamMembers) SetName(v string) *UpdateTeamRequestTeamMembers {
	s.Name = &v
	return s
}

func (s *UpdateTeamRequestTeamMembers) Validate() error {
	return dara.Validate(s)
}
