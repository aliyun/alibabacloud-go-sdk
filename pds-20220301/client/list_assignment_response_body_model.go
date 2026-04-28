// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssignmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssignmentList(v []*ListAssignmentResponseBodyAssignmentList) *ListAssignmentResponseBody
	GetAssignmentList() []*ListAssignmentResponseBodyAssignmentList
	SetNextMarker(v string) *ListAssignmentResponseBody
	GetNextMarker() *string
}

type ListAssignmentResponseBody struct {
	// The assigned roles.
	AssignmentList []*ListAssignmentResponseBodyAssignmentList `json:"assignment_list,omitempty" xml:"assignment_list,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListAssignmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAssignmentResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssignmentResponseBody) GetAssignmentList() []*ListAssignmentResponseBodyAssignmentList {
	return s.AssignmentList
}

func (s *ListAssignmentResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListAssignmentResponseBody) SetAssignmentList(v []*ListAssignmentResponseBodyAssignmentList) *ListAssignmentResponseBody {
	s.AssignmentList = v
	return s
}

func (s *ListAssignmentResponseBody) SetNextMarker(v string) *ListAssignmentResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListAssignmentResponseBody) Validate() error {
	if s.AssignmentList != nil {
		for _, item := range s.AssignmentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAssignmentResponseBodyAssignmentList struct {
	// The time when the role was assigned. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1622682267564
	CreatedAt *int64 `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The ID of the user who assigned the role.
	//
	// example:
	//
	// 216***c83
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// hz1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The identity to whom the role is assigned, which is a user or a group.
	Identity *Identity `json:"identity,omitempty" xml:"identity,omitempty"`
	// The ID of the managed resource, such as a group ID.
	//
	// example:
	//
	// 105***b82
	ManageResourceId *string `json:"manage_resource_id,omitempty" xml:"manage_resource_id,omitempty"`
	// The type of the managed resource. For example, a value of RT_Group indicates group.
	//
	// example:
	//
	// RT_Group
	ManageResourceType *string `json:"manage_resource_type,omitempty" xml:"manage_resource_type,omitempty"`
	// The ID of the role assigned to the identity.
	//
	// example:
	//
	// SystemGroupAdmin
	RoleId *string `json:"role_id,omitempty" xml:"role_id,omitempty"`
}

func (s ListAssignmentResponseBodyAssignmentList) String() string {
	return dara.Prettify(s)
}

func (s ListAssignmentResponseBodyAssignmentList) GoString() string {
	return s.String()
}

func (s *ListAssignmentResponseBodyAssignmentList) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *ListAssignmentResponseBodyAssignmentList) GetCreator() *string {
	return s.Creator
}

func (s *ListAssignmentResponseBodyAssignmentList) GetDomainId() *string {
	return s.DomainId
}

func (s *ListAssignmentResponseBodyAssignmentList) GetIdentity() *Identity {
	return s.Identity
}

func (s *ListAssignmentResponseBodyAssignmentList) GetManageResourceId() *string {
	return s.ManageResourceId
}

func (s *ListAssignmentResponseBodyAssignmentList) GetManageResourceType() *string {
	return s.ManageResourceType
}

func (s *ListAssignmentResponseBodyAssignmentList) GetRoleId() *string {
	return s.RoleId
}

func (s *ListAssignmentResponseBodyAssignmentList) SetCreatedAt(v int64) *ListAssignmentResponseBodyAssignmentList {
	s.CreatedAt = &v
	return s
}

func (s *ListAssignmentResponseBodyAssignmentList) SetCreator(v string) *ListAssignmentResponseBodyAssignmentList {
	s.Creator = &v
	return s
}

func (s *ListAssignmentResponseBodyAssignmentList) SetDomainId(v string) *ListAssignmentResponseBodyAssignmentList {
	s.DomainId = &v
	return s
}

func (s *ListAssignmentResponseBodyAssignmentList) SetIdentity(v *Identity) *ListAssignmentResponseBodyAssignmentList {
	s.Identity = v
	return s
}

func (s *ListAssignmentResponseBodyAssignmentList) SetManageResourceId(v string) *ListAssignmentResponseBodyAssignmentList {
	s.ManageResourceId = &v
	return s
}

func (s *ListAssignmentResponseBodyAssignmentList) SetManageResourceType(v string) *ListAssignmentResponseBodyAssignmentList {
	s.ManageResourceType = &v
	return s
}

func (s *ListAssignmentResponseBodyAssignmentList) SetRoleId(v string) *ListAssignmentResponseBodyAssignmentList {
	s.RoleId = &v
	return s
}

func (s *ListAssignmentResponseBodyAssignmentList) Validate() error {
	if s.Identity != nil {
		if err := s.Identity.Validate(); err != nil {
			return err
		}
	}
	return nil
}
