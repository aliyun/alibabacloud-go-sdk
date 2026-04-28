// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBaseAssignmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetAssociatedRoleTagId(v string) *BaseAssignmentResponse
	GetAssociatedRoleTagId() *string
	SetCreatedAt(v string) *BaseAssignmentResponse
	GetCreatedAt() *string
	SetCreator(v string) *BaseAssignmentResponse
	GetCreator() *string
	SetDisinheritSubGroup(v bool) *BaseAssignmentResponse
	GetDisinheritSubGroup() *bool
	SetDomainId(v string) *BaseAssignmentResponse
	GetDomainId() *string
	SetIdentity(v *Identity) *BaseAssignmentResponse
	GetIdentity() *Identity
	SetManageResourceId(v string) *BaseAssignmentResponse
	GetManageResourceId() *string
	SetManageResourceType(v string) *BaseAssignmentResponse
	GetManageResourceType() *string
	SetRoleId(v string) *BaseAssignmentResponse
	GetRoleId() *string
	SetUpdatedAt(v string) *BaseAssignmentResponse
	GetUpdatedAt() *string
}

type BaseAssignmentResponse struct {
	AssociatedRoleTagId *string   `json:"associated_role_tag_id,omitempty" xml:"associated_role_tag_id,omitempty"`
	CreatedAt           *string   `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Creator             *string   `json:"creator,omitempty" xml:"creator,omitempty"`
	DisinheritSubGroup  *bool     `json:"disinherit_sub_group,omitempty" xml:"disinherit_sub_group,omitempty"`
	DomainId            *string   `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	Identity            *Identity `json:"identity,omitempty" xml:"identity,omitempty"`
	ManageResourceId    *string   `json:"manage_resource_id,omitempty" xml:"manage_resource_id,omitempty"`
	ManageResourceType  *string   `json:"manage_resource_type,omitempty" xml:"manage_resource_type,omitempty"`
	RoleId              *string   `json:"role_id,omitempty" xml:"role_id,omitempty"`
	UpdatedAt           *string   `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s BaseAssignmentResponse) String() string {
	return dara.Prettify(s)
}

func (s BaseAssignmentResponse) GoString() string {
	return s.String()
}

func (s *BaseAssignmentResponse) GetAssociatedRoleTagId() *string {
	return s.AssociatedRoleTagId
}

func (s *BaseAssignmentResponse) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *BaseAssignmentResponse) GetCreator() *string {
	return s.Creator
}

func (s *BaseAssignmentResponse) GetDisinheritSubGroup() *bool {
	return s.DisinheritSubGroup
}

func (s *BaseAssignmentResponse) GetDomainId() *string {
	return s.DomainId
}

func (s *BaseAssignmentResponse) GetIdentity() *Identity {
	return s.Identity
}

func (s *BaseAssignmentResponse) GetManageResourceId() *string {
	return s.ManageResourceId
}

func (s *BaseAssignmentResponse) GetManageResourceType() *string {
	return s.ManageResourceType
}

func (s *BaseAssignmentResponse) GetRoleId() *string {
	return s.RoleId
}

func (s *BaseAssignmentResponse) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *BaseAssignmentResponse) SetAssociatedRoleTagId(v string) *BaseAssignmentResponse {
	s.AssociatedRoleTagId = &v
	return s
}

func (s *BaseAssignmentResponse) SetCreatedAt(v string) *BaseAssignmentResponse {
	s.CreatedAt = &v
	return s
}

func (s *BaseAssignmentResponse) SetCreator(v string) *BaseAssignmentResponse {
	s.Creator = &v
	return s
}

func (s *BaseAssignmentResponse) SetDisinheritSubGroup(v bool) *BaseAssignmentResponse {
	s.DisinheritSubGroup = &v
	return s
}

func (s *BaseAssignmentResponse) SetDomainId(v string) *BaseAssignmentResponse {
	s.DomainId = &v
	return s
}

func (s *BaseAssignmentResponse) SetIdentity(v *Identity) *BaseAssignmentResponse {
	s.Identity = v
	return s
}

func (s *BaseAssignmentResponse) SetManageResourceId(v string) *BaseAssignmentResponse {
	s.ManageResourceId = &v
	return s
}

func (s *BaseAssignmentResponse) SetManageResourceType(v string) *BaseAssignmentResponse {
	s.ManageResourceType = &v
	return s
}

func (s *BaseAssignmentResponse) SetRoleId(v string) *BaseAssignmentResponse {
	s.RoleId = &v
	return s
}

func (s *BaseAssignmentResponse) SetUpdatedAt(v string) *BaseAssignmentResponse {
	s.UpdatedAt = &v
	return s
}

func (s *BaseAssignmentResponse) Validate() error {
	if s.Identity != nil {
		if err := s.Identity.Validate(); err != nil {
			return err
		}
	}
	return nil
}
