// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserGroupResponseBody
	GetRequestId() *string
	SetUserGroup(v *GetUserGroupResponseBodyUserGroup) *GetUserGroupResponseBody
	GetUserGroup() *GetUserGroupResponseBodyUserGroup
}

type GetUserGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the user group returned.
	UserGroup *GetUserGroupResponseBodyUserGroup `json:"UserGroup,omitempty" xml:"UserGroup,omitempty" type:"Struct"`
}

func (s GetUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserGroupResponseBody) GetUserGroup() *GetUserGroupResponseBodyUserGroup {
	return s.UserGroup
}

func (s *GetUserGroupResponseBody) SetRequestId(v string) *GetUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserGroupResponseBody) SetUserGroup(v *GetUserGroupResponseBodyUserGroup) *GetUserGroupResponseBody {
	s.UserGroup = v
	return s
}

func (s *GetUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUserGroupResponseBodyUserGroup struct {
	// The description of the user group.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the group.
	//
	// example:
	//
	// 1
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The name of the user group.
	//
	// example:
	//
	// UserGroup01
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s GetUserGroupResponseBodyUserGroup) String() string {
	return dara.Prettify(s)
}

func (s GetUserGroupResponseBodyUserGroup) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponseBodyUserGroup) GetComment() *string {
	return s.Comment
}

func (s *GetUserGroupResponseBodyUserGroup) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *GetUserGroupResponseBodyUserGroup) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *GetUserGroupResponseBodyUserGroup) SetComment(v string) *GetUserGroupResponseBodyUserGroup {
	s.Comment = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroup) SetUserGroupId(v string) *GetUserGroupResponseBodyUserGroup {
	s.UserGroupId = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroup) SetUserGroupName(v string) *GetUserGroupResponseBodyUserGroup {
	s.UserGroupName = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroup) Validate() error {
	return dara.Validate(s)
}
