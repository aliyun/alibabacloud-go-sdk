// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBaseRoleMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetAssignmentList(v []*BaseAssignmentResponse) *BaseRoleMemberResponse
	GetAssignmentList() []*BaseAssignmentResponse
	SetCreatedAt(v string) *BaseRoleMemberResponse
	GetCreatedAt() *string
	SetCreator(v string) *BaseRoleMemberResponse
	GetCreator() *string
	SetDomainId(v string) *BaseRoleMemberResponse
	GetDomainId() *string
	SetIdentity(v *Identity) *BaseRoleMemberResponse
	GetIdentity() *Identity
	SetIdentityName(v string) *BaseRoleMemberResponse
	GetIdentityName() *string
	SetIsAdmin(v bool) *BaseRoleMemberResponse
	GetIsAdmin() *bool
	SetSubdomainId(v string) *BaseRoleMemberResponse
	GetSubdomainId() *string
}

type BaseRoleMemberResponse struct {
	AssignmentList []*BaseAssignmentResponse `json:"assignment_list,omitempty" xml:"assignment_list,omitempty" type:"Repeated"`
	CreatedAt      *string                   `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Creator        *string                   `json:"creator,omitempty" xml:"creator,omitempty"`
	DomainId       *string                   `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	Identity       *Identity                 `json:"identity,omitempty" xml:"identity,omitempty"`
	IdentityName   *string                   `json:"identity_name,omitempty" xml:"identity_name,omitempty"`
	IsAdmin        *bool                     `json:"is_admin,omitempty" xml:"is_admin,omitempty"`
	SubdomainId    *string                   `json:"subdomain_id,omitempty" xml:"subdomain_id,omitempty"`
}

func (s BaseRoleMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s BaseRoleMemberResponse) GoString() string {
	return s.String()
}

func (s *BaseRoleMemberResponse) GetAssignmentList() []*BaseAssignmentResponse {
	return s.AssignmentList
}

func (s *BaseRoleMemberResponse) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *BaseRoleMemberResponse) GetCreator() *string {
	return s.Creator
}

func (s *BaseRoleMemberResponse) GetDomainId() *string {
	return s.DomainId
}

func (s *BaseRoleMemberResponse) GetIdentity() *Identity {
	return s.Identity
}

func (s *BaseRoleMemberResponse) GetIdentityName() *string {
	return s.IdentityName
}

func (s *BaseRoleMemberResponse) GetIsAdmin() *bool {
	return s.IsAdmin
}

func (s *BaseRoleMemberResponse) GetSubdomainId() *string {
	return s.SubdomainId
}

func (s *BaseRoleMemberResponse) SetAssignmentList(v []*BaseAssignmentResponse) *BaseRoleMemberResponse {
	s.AssignmentList = v
	return s
}

func (s *BaseRoleMemberResponse) SetCreatedAt(v string) *BaseRoleMemberResponse {
	s.CreatedAt = &v
	return s
}

func (s *BaseRoleMemberResponse) SetCreator(v string) *BaseRoleMemberResponse {
	s.Creator = &v
	return s
}

func (s *BaseRoleMemberResponse) SetDomainId(v string) *BaseRoleMemberResponse {
	s.DomainId = &v
	return s
}

func (s *BaseRoleMemberResponse) SetIdentity(v *Identity) *BaseRoleMemberResponse {
	s.Identity = v
	return s
}

func (s *BaseRoleMemberResponse) SetIdentityName(v string) *BaseRoleMemberResponse {
	s.IdentityName = &v
	return s
}

func (s *BaseRoleMemberResponse) SetIsAdmin(v bool) *BaseRoleMemberResponse {
	s.IsAdmin = &v
	return s
}

func (s *BaseRoleMemberResponse) SetSubdomainId(v string) *BaseRoleMemberResponse {
	s.SubdomainId = &v
	return s
}

func (s *BaseRoleMemberResponse) Validate() error {
	if s.AssignmentList != nil {
		for _, item := range s.AssignmentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Identity != nil {
		if err := s.Identity.Validate(); err != nil {
			return err
		}
	}
	return nil
}
