// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPlayerList(v []*CreateAppMembersRequestPlayerList) *CreateAppMembersRequest
	GetPlayerList() []*CreateAppMembersRequestPlayerList
	SetRoleNames(v []*string) *CreateAppMembersRequest
	GetRoleNames() []*string
	SetOrganizationId(v string) *CreateAppMembersRequest
	GetOrganizationId() *string
}

type CreateAppMembersRequest struct {
	PlayerList []*CreateAppMembersRequestPlayerList `json:"playerList,omitempty" xml:"playerList,omitempty" type:"Repeated"`
	RoleNames  []*string                            `json:"roleNames,omitempty" xml:"roleNames,omitempty" type:"Repeated"`
	// example:
	//
	// 66c0c9fffeb86b450c199fcd
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s CreateAppMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppMembersRequest) GoString() string {
	return s.String()
}

func (s *CreateAppMembersRequest) GetPlayerList() []*CreateAppMembersRequestPlayerList {
	return s.PlayerList
}

func (s *CreateAppMembersRequest) GetRoleNames() []*string {
	return s.RoleNames
}

func (s *CreateAppMembersRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateAppMembersRequest) SetPlayerList(v []*CreateAppMembersRequestPlayerList) *CreateAppMembersRequest {
	s.PlayerList = v
	return s
}

func (s *CreateAppMembersRequest) SetRoleNames(v []*string) *CreateAppMembersRequest {
	s.RoleNames = v
	return s
}

func (s *CreateAppMembersRequest) SetOrganizationId(v string) *CreateAppMembersRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateAppMembersRequest) Validate() error {
	if s.PlayerList != nil {
		for _, item := range s.PlayerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAppMembersRequestPlayerList struct {
	// example:
	//
	// 1332695887xxxxxx
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// User
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAppMembersRequestPlayerList) String() string {
	return dara.Prettify(s)
}

func (s CreateAppMembersRequestPlayerList) GoString() string {
	return s.String()
}

func (s *CreateAppMembersRequestPlayerList) GetId() *string {
	return s.Id
}

func (s *CreateAppMembersRequestPlayerList) GetType() *string {
	return s.Type
}

func (s *CreateAppMembersRequestPlayerList) SetId(v string) *CreateAppMembersRequestPlayerList {
	s.Id = &v
	return s
}

func (s *CreateAppMembersRequestPlayerList) SetType(v string) *CreateAppMembersRequestPlayerList {
	s.Type = &v
	return s
}

func (s *CreateAppMembersRequestPlayerList) Validate() error {
	return dara.Validate(s)
}
