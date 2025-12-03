// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPlayer(v *UpdateAppMemberRequestPlayer) *UpdateAppMemberRequest
	GetPlayer() *UpdateAppMemberRequestPlayer
	SetRoleNames(v []*string) *UpdateAppMemberRequest
	GetRoleNames() []*string
	SetOrganizationId(v string) *UpdateAppMemberRequest
	GetOrganizationId() *string
}

type UpdateAppMemberRequest struct {
	Player    *UpdateAppMemberRequestPlayer `json:"player,omitempty" xml:"player,omitempty" type:"Struct"`
	RoleNames []*string                     `json:"roleNames,omitempty" xml:"roleNames,omitempty" type:"Repeated"`
	// example:
	//
	// 66c0c9fffeb86b450c199fcd
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s UpdateAppMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppMemberRequest) GetPlayer() *UpdateAppMemberRequestPlayer {
	return s.Player
}

func (s *UpdateAppMemberRequest) GetRoleNames() []*string {
	return s.RoleNames
}

func (s *UpdateAppMemberRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *UpdateAppMemberRequest) SetPlayer(v *UpdateAppMemberRequestPlayer) *UpdateAppMemberRequest {
	s.Player = v
	return s
}

func (s *UpdateAppMemberRequest) SetRoleNames(v []*string) *UpdateAppMemberRequest {
	s.RoleNames = v
	return s
}

func (s *UpdateAppMemberRequest) SetOrganizationId(v string) *UpdateAppMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *UpdateAppMemberRequest) Validate() error {
	if s.Player != nil {
		if err := s.Player.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAppMemberRequestPlayer struct {
	// example:
	//
	// 1332695887xxxxxx
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// User
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateAppMemberRequestPlayer) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppMemberRequestPlayer) GoString() string {
	return s.String()
}

func (s *UpdateAppMemberRequestPlayer) GetId() *string {
	return s.Id
}

func (s *UpdateAppMemberRequestPlayer) GetType() *string {
	return s.Type
}

func (s *UpdateAppMemberRequestPlayer) SetId(v string) *UpdateAppMemberRequestPlayer {
	s.Id = &v
	return s
}

func (s *UpdateAppMemberRequestPlayer) SetType(v string) *UpdateAppMemberRequestPlayer {
	s.Type = &v
	return s
}

func (s *UpdateAppMemberRequestPlayer) Validate() error {
	return dara.Validate(s)
}
