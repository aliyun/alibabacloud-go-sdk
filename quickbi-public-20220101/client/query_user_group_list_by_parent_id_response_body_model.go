// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserGroupListByParentIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryUserGroupListByParentIdResponseBody
	GetRequestId() *string
	SetResult(v []*QueryUserGroupListByParentIdResponseBodyResult) *QueryUserGroupListByParentIdResponseBody
	GetResult() []*QueryUserGroupListByParentIdResponseBodyResult
	SetSuccess(v bool) *QueryUserGroupListByParentIdResponseBody
	GetSuccess() *bool
}

type QueryUserGroupListByParentIdResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 72B19D61-B37A-5C7A-9389-0856CD7935B3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the sub-user group.
	Result []*QueryUserGroupListByParentIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserGroupListByParentIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUserGroupListByParentIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserGroupListByParentIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUserGroupListByParentIdResponseBody) GetResult() []*QueryUserGroupListByParentIdResponseBodyResult {
	return s.Result
}

func (s *QueryUserGroupListByParentIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUserGroupListByParentIdResponseBody) SetRequestId(v string) *QueryUserGroupListByParentIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBody) SetResult(v []*QueryUserGroupListByParentIdResponseBodyResult) *QueryUserGroupListByParentIdResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBody) SetSuccess(v bool) *QueryUserGroupListByParentIdResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryUserGroupListByParentIdResponseBodyResult struct {
	// The time when the sub-user group was created.
	//
	// example:
	//
	// 2020-10-30 10:03:09
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator of the sub-user group. The UserID of the Quick BI is used instead of the UID of Alibaba Cloud.
	//
	// example:
	//
	// 136516262323****
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// Directory level of the sub-user group.
	IdentifiedPath *string `json:"IdentifiedPath,omitempty" xml:"IdentifiedPath,omitempty"`
	// The time when the sub-user group was last modified.
	//
	// example:
	//
	// 2020-11-16 15:49:08
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The user who modified the subgroup. The UserID of the Quick BI is used instead of the UID of Alibaba Cloud.
	//
	// example:
	//
	// 136516262323****
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The ID of the parent user group.
	//
	// example:
	//
	// 3d2c23d4-2b41-4af8-a1f5-f6390f32****
	ParentUserGroupId *string `json:"ParentUserGroupId,omitempty" xml:"ParentUserGroupId,omitempty"`
	// The description of the sub-user group.
	//
	// example:
	//
	// User Group for Testing
	UserGroupDescription *string `json:"UserGroupDescription,omitempty" xml:"UserGroupDescription,omitempty"`
	// The ID of the sub-user group.
	//
	// example:
	//
	// f5eeb52e-d9c2-4a8b-80e3-47ab55c2****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The name of the sub-user group.
	//
	// example:
	//
	// popapi test group
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s QueryUserGroupListByParentIdResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryUserGroupListByParentIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) GetCreateUser() *string {
	return s.CreateUser
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) GetIdentifiedPath() *string {
	return s.IdentifiedPath
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) GetParentUserGroupId() *string {
	return s.ParentUserGroupId
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) GetUserGroupDescription() *string {
	return s.UserGroupDescription
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) SetCreateTime(v string) *QueryUserGroupListByParentIdResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) SetCreateUser(v string) *QueryUserGroupListByParentIdResponseBodyResult {
	s.CreateUser = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) SetIdentifiedPath(v string) *QueryUserGroupListByParentIdResponseBodyResult {
	s.IdentifiedPath = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) SetModifiedTime(v string) *QueryUserGroupListByParentIdResponseBodyResult {
	s.ModifiedTime = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) SetModifyUser(v string) *QueryUserGroupListByParentIdResponseBodyResult {
	s.ModifyUser = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) SetParentUserGroupId(v string) *QueryUserGroupListByParentIdResponseBodyResult {
	s.ParentUserGroupId = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) SetUserGroupDescription(v string) *QueryUserGroupListByParentIdResponseBodyResult {
	s.UserGroupDescription = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) SetUserGroupId(v string) *QueryUserGroupListByParentIdResponseBodyResult {
	s.UserGroupId = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) SetUserGroupName(v string) *QueryUserGroupListByParentIdResponseBodyResult {
	s.UserGroupName = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
