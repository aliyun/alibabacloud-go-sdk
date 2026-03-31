// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsTruncated(v bool) *ListRolesResponseBody
	GetIsTruncated() *bool
	SetMarker(v string) *ListRolesResponseBody
	GetMarker() *string
	SetRequestId(v string) *ListRolesResponseBody
	GetRequestId() *string
	SetRoles(v *ListRolesResponseBodyRoles) *ListRolesResponseBody
	GetRoles() *ListRolesResponseBodyRoles
}

type ListRolesResponseBody struct {
	// Indicates whether the response is truncated.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The marker. This parameter is returned only if the value of `IsTruncated` is `true`. If the parameter is returned, you can call this operation again and set this parameter to obtain the truncated part.````
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Roles     *ListRolesResponseBodyRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Struct"`
}

func (s ListRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListRolesResponseBody) GetMarker() *string {
	return s.Marker
}

func (s *ListRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRolesResponseBody) GetRoles() *ListRolesResponseBodyRoles {
	return s.Roles
}

func (s *ListRolesResponseBody) SetIsTruncated(v bool) *ListRolesResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListRolesResponseBody) SetMarker(v string) *ListRolesResponseBody {
	s.Marker = &v
	return s
}

func (s *ListRolesResponseBody) SetRequestId(v string) *ListRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRolesResponseBody) SetRoles(v *ListRolesResponseBodyRoles) *ListRolesResponseBody {
	s.Roles = v
	return s
}

func (s *ListRolesResponseBody) Validate() error {
	if s.Roles != nil {
		if err := s.Roles.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRolesResponseBodyRoles struct {
	Role []*ListRolesResponseBodyRolesRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Repeated"`
}

func (s ListRolesResponseBodyRoles) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBodyRoles) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyRoles) GetRole() []*ListRolesResponseBodyRolesRole {
	return s.Role
}

func (s *ListRolesResponseBodyRoles) SetRole(v []*ListRolesResponseBodyRolesRole) *ListRolesResponseBodyRoles {
	s.Role = v
	return s
}

func (s *ListRolesResponseBodyRoles) Validate() error {
	if s.Role != nil {
		for _, item := range s.Role {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRolesResponseBodyRolesRole struct {
	Arn                *string                             `json:"Arn,omitempty" xml:"Arn,omitempty"`
	CreateDate         *string                             `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Description        *string                             `json:"Description,omitempty" xml:"Description,omitempty"`
	MaxSessionDuration *int64                              `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	RoleId             *string                             `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	RoleName           *string                             `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	Tags               *ListRolesResponseBodyRolesRoleTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UpdateDate         *string                             `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListRolesResponseBodyRolesRole) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBodyRolesRole) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyRolesRole) GetArn() *string {
	return s.Arn
}

func (s *ListRolesResponseBodyRolesRole) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListRolesResponseBodyRolesRole) GetDescription() *string {
	return s.Description
}

func (s *ListRolesResponseBodyRolesRole) GetMaxSessionDuration() *int64 {
	return s.MaxSessionDuration
}

func (s *ListRolesResponseBodyRolesRole) GetRoleId() *string {
	return s.RoleId
}

func (s *ListRolesResponseBodyRolesRole) GetRoleName() *string {
	return s.RoleName
}

func (s *ListRolesResponseBodyRolesRole) GetTags() *ListRolesResponseBodyRolesRoleTags {
	return s.Tags
}

func (s *ListRolesResponseBodyRolesRole) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ListRolesResponseBodyRolesRole) SetArn(v string) *ListRolesResponseBodyRolesRole {
	s.Arn = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetCreateDate(v string) *ListRolesResponseBodyRolesRole {
	s.CreateDate = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetDescription(v string) *ListRolesResponseBodyRolesRole {
	s.Description = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetMaxSessionDuration(v int64) *ListRolesResponseBodyRolesRole {
	s.MaxSessionDuration = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetRoleId(v string) *ListRolesResponseBodyRolesRole {
	s.RoleId = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetRoleName(v string) *ListRolesResponseBodyRolesRole {
	s.RoleName = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetTags(v *ListRolesResponseBodyRolesRoleTags) *ListRolesResponseBodyRolesRole {
	s.Tags = v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetUpdateDate(v string) *ListRolesResponseBodyRolesRole {
	s.UpdateDate = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRolesResponseBodyRolesRoleTags struct {
	Tag []*ListRolesResponseBodyRolesRoleTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListRolesResponseBodyRolesRoleTags) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBodyRolesRoleTags) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyRolesRoleTags) GetTag() []*ListRolesResponseBodyRolesRoleTagsTag {
	return s.Tag
}

func (s *ListRolesResponseBodyRolesRoleTags) SetTag(v []*ListRolesResponseBodyRolesRoleTagsTag) *ListRolesResponseBodyRolesRoleTags {
	s.Tag = v
	return s
}

func (s *ListRolesResponseBodyRolesRoleTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRolesResponseBodyRolesRoleTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListRolesResponseBodyRolesRoleTagsTag) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBodyRolesRoleTagsTag) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyRolesRoleTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *ListRolesResponseBodyRolesRoleTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *ListRolesResponseBodyRolesRoleTagsTag) SetTagKey(v string) *ListRolesResponseBodyRolesRoleTagsTag {
	s.TagKey = &v
	return s
}

func (s *ListRolesResponseBodyRolesRoleTagsTag) SetTagValue(v string) *ListRolesResponseBodyRolesRoleTagsTag {
	s.TagValue = &v
	return s
}

func (s *ListRolesResponseBodyRolesRoleTagsTag) Validate() error {
	return dara.Validate(s)
}
