// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTeamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdminName(v string) *CreateTeamRequest
	GetAdminName() *string
	SetClientToken(v string) *CreateTeamRequest
	GetClientToken() *string
	SetDescription(v string) *CreateTeamRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateTeamRequest
	GetInstanceId() *string
	SetName(v string) *CreateTeamRequest
	GetName() *string
	SetTeamMembers(v []*CreateTeamRequestTeamMembers) *CreateTeamRequest
	GetTeamMembers() []*CreateTeamRequestTeamMembers
}

type CreateTeamRequest struct {
	AdminName   *string                         `json:"AdminName,omitempty" xml:"AdminName,omitempty"`
	ClientToken *string                         `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string                         `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId  *string                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
	TeamMembers []*CreateTeamRequestTeamMembers `json:"TeamMembers,omitempty" xml:"TeamMembers,omitempty" type:"Repeated"`
}

func (s CreateTeamRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTeamRequest) GoString() string {
	return s.String()
}

func (s *CreateTeamRequest) GetAdminName() *string {
	return s.AdminName
}

func (s *CreateTeamRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTeamRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTeamRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTeamRequest) GetName() *string {
	return s.Name
}

func (s *CreateTeamRequest) GetTeamMembers() []*CreateTeamRequestTeamMembers {
	return s.TeamMembers
}

func (s *CreateTeamRequest) SetAdminName(v string) *CreateTeamRequest {
	s.AdminName = &v
	return s
}

func (s *CreateTeamRequest) SetClientToken(v string) *CreateTeamRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTeamRequest) SetDescription(v string) *CreateTeamRequest {
	s.Description = &v
	return s
}

func (s *CreateTeamRequest) SetInstanceId(v string) *CreateTeamRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTeamRequest) SetName(v string) *CreateTeamRequest {
	s.Name = &v
	return s
}

func (s *CreateTeamRequest) SetTeamMembers(v []*CreateTeamRequestTeamMembers) *CreateTeamRequest {
	s.TeamMembers = v
	return s
}

func (s *CreateTeamRequest) Validate() error {
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

type CreateTeamRequestTeamMembers struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateTeamRequestTeamMembers) String() string {
	return dara.Prettify(s)
}

func (s CreateTeamRequestTeamMembers) GoString() string {
	return s.String()
}

func (s *CreateTeamRequestTeamMembers) GetName() *string {
	return s.Name
}

func (s *CreateTeamRequestTeamMembers) SetName(v string) *CreateTeamRequestTeamMembers {
	s.Name = &v
	return s
}

func (s *CreateTeamRequestTeamMembers) Validate() error {
	return dara.Validate(s)
}
