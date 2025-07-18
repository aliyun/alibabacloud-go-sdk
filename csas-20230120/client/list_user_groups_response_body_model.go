// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListUserGroupsResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *ListUserGroupsResponseBody
	GetTotalNum() *int32
	SetUserGroups(v []*ListUserGroupsResponseBodyUserGroups) *ListUserGroupsResponseBody
	GetUserGroups() []*ListUserGroupsResponseBodyUserGroups
}

type ListUserGroupsResponseBody struct {
	// example:
	//
	// 4AB972E2-D702-5464-B132-B1911498B8BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum   *int32                                  `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	UserGroups []*ListUserGroupsResponseBodyUserGroups `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
}

func (s ListUserGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserGroupsResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListUserGroupsResponseBody) GetUserGroups() []*ListUserGroupsResponseBodyUserGroups {
	return s.UserGroups
}

func (s *ListUserGroupsResponseBody) SetRequestId(v string) *ListUserGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetTotalNum(v int32) *ListUserGroupsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetUserGroups(v []*ListUserGroupsResponseBodyUserGroups) *ListUserGroupsResponseBody {
	s.UserGroups = v
	return s
}

func (s *ListUserGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUserGroupsResponseBodyUserGroups struct {
	Attributes []*ListUserGroupsResponseBodyUserGroupsAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-10-10 11:39:22
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// user_group_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// usergroup-6f1ef2fc56b6****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListUserGroupsResponseBodyUserGroups) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsResponseBodyUserGroups) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBodyUserGroups) GetAttributes() []*ListUserGroupsResponseBodyUserGroupsAttributes {
	return s.Attributes
}

func (s *ListUserGroupsResponseBodyUserGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListUserGroupsResponseBodyUserGroups) GetDescription() *string {
	return s.Description
}

func (s *ListUserGroupsResponseBodyUserGroups) GetName() *string {
	return s.Name
}

func (s *ListUserGroupsResponseBodyUserGroups) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListUserGroupsResponseBodyUserGroups) SetAttributes(v []*ListUserGroupsResponseBodyUserGroupsAttributes) *ListUserGroupsResponseBodyUserGroups {
	s.Attributes = v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroups) SetCreateTime(v string) *ListUserGroupsResponseBodyUserGroups {
	s.CreateTime = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroups) SetDescription(v string) *ListUserGroupsResponseBodyUserGroups {
	s.Description = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroups) SetName(v string) *ListUserGroupsResponseBodyUserGroups {
	s.Name = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroups) SetUserGroupId(v string) *ListUserGroupsResponseBodyUserGroups {
	s.UserGroupId = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroups) Validate() error {
	return dara.Validate(s)
}

type ListUserGroupsResponseBodyUserGroupsAttributes struct {
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListUserGroupsResponseBodyUserGroupsAttributes) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsResponseBodyUserGroupsAttributes) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBodyUserGroupsAttributes) GetIdpId() *int32 {
	return s.IdpId
}

func (s *ListUserGroupsResponseBodyUserGroupsAttributes) GetRelation() *string {
	return s.Relation
}

func (s *ListUserGroupsResponseBodyUserGroupsAttributes) GetUserGroupType() *string {
	return s.UserGroupType
}

func (s *ListUserGroupsResponseBodyUserGroupsAttributes) GetValue() *string {
	return s.Value
}

func (s *ListUserGroupsResponseBodyUserGroupsAttributes) SetIdpId(v int32) *ListUserGroupsResponseBodyUserGroupsAttributes {
	s.IdpId = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroupsAttributes) SetRelation(v string) *ListUserGroupsResponseBodyUserGroupsAttributes {
	s.Relation = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroupsAttributes) SetUserGroupType(v string) *ListUserGroupsResponseBodyUserGroupsAttributes {
	s.UserGroupType = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroupsAttributes) SetValue(v string) *ListUserGroupsResponseBodyUserGroupsAttributes {
	s.Value = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroupsAttributes) Validate() error {
	return dara.Validate(s)
}
