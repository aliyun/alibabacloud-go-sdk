// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListByUserGroupIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListByUserGroupIdResponseBody
	GetRequestId() *string
	SetResult(v *ListByUserGroupIdResponseBodyResult) *ListByUserGroupIdResponseBody
	GetResult() *ListByUserGroupIdResponseBodyResult
	SetSuccess(v bool) *ListByUserGroupIdResponseBody
	GetSuccess() *bool
}

type ListByUserGroupIdResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The user group query result is returned.
	Result *ListByUserGroupIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s ListByUserGroupIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListByUserGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListByUserGroupIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListByUserGroupIdResponseBody) GetResult() *ListByUserGroupIdResponseBodyResult {
	return s.Result
}

func (s *ListByUserGroupIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListByUserGroupIdResponseBody) SetRequestId(v string) *ListByUserGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListByUserGroupIdResponseBody) SetResult(v *ListByUserGroupIdResponseBodyResult) *ListByUserGroupIdResponseBody {
	s.Result = v
	return s
}

func (s *ListByUserGroupIdResponseBody) SetSuccess(v bool) *ListByUserGroupIdResponseBody {
	s.Success = &v
	return s
}

func (s *ListByUserGroupIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListByUserGroupIdResponseBodyResult struct {
	// List of failed user groups.
	FailedUserGroupIds []*string `json:"FailedUserGroupIds,omitempty" xml:"FailedUserGroupIds,omitempty" type:"Repeated"`
	// The details of the user group that was queried.
	UserGroupModels []*ListByUserGroupIdResponseBodyResultUserGroupModels `json:"UserGroupModels,omitempty" xml:"UserGroupModels,omitempty" type:"Repeated"`
}

func (s ListByUserGroupIdResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListByUserGroupIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListByUserGroupIdResponseBodyResult) GetFailedUserGroupIds() []*string {
	return s.FailedUserGroupIds
}

func (s *ListByUserGroupIdResponseBodyResult) GetUserGroupModels() []*ListByUserGroupIdResponseBodyResultUserGroupModels {
	return s.UserGroupModels
}

func (s *ListByUserGroupIdResponseBodyResult) SetFailedUserGroupIds(v []*string) *ListByUserGroupIdResponseBodyResult {
	s.FailedUserGroupIds = v
	return s
}

func (s *ListByUserGroupIdResponseBodyResult) SetUserGroupModels(v []*ListByUserGroupIdResponseBodyResultUserGroupModels) *ListByUserGroupIdResponseBodyResult {
	s.UserGroupModels = v
	return s
}

func (s *ListByUserGroupIdResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListByUserGroupIdResponseBodyResultUserGroupModels struct {
	// The time when the Secret was created.
	//
	// example:
	//
	// 2021-03-15 17:13:55
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The UserID of the creator in the Quick BI.
	//
	// example:
	//
	// 46e5*******ee22e2a292704c8
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The path of the user group.
	//
	// example:
	//
	// 2fe4fbd8-****-af083ea/34fd1-****-dcbc33f
	IdentifiedPath *string `json:"IdentifiedPath,omitempty" xml:"IdentifiedPath,omitempty"`
	// The time when the protection policy was last modified.
	//
	// example:
	//
	// 2021-03-15 20:36:40
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The UserID of the modifier in the Quick BI.
	//
	// example:
	//
	// 46e5*******ee22e2a292704c8
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The ID of the parent user group.
	//
	// example:
	//
	// 2fe4fbd8-588f-489a-b3e1-e92c7af083ea
	ParentUsergroupId *string `json:"ParentUsergroupId,omitempty" xml:"ParentUsergroupId,omitempty"`
	// The description of the user group.
	//
	// example:
	//
	// Description
	UsergroupDesc *string `json:"UsergroupDesc,omitempty" xml:"UsergroupDesc,omitempty"`
	// The ID of the user group.
	//
	// example:
	//
	// 34fd141d-****-4093-8c33-8e066dcbc33f
	UsergroupId *string `json:"UsergroupId,omitempty" xml:"UsergroupId,omitempty"`
	// The name of the user group.
	//
	// example:
	//
	// Test user group
	UsergroupName *string `json:"UsergroupName,omitempty" xml:"UsergroupName,omitempty"`
}

func (s ListByUserGroupIdResponseBodyResultUserGroupModels) String() string {
	return dara.Prettify(s)
}

func (s ListByUserGroupIdResponseBodyResultUserGroupModels) GoString() string {
	return s.String()
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) GetIdentifiedPath() *string {
	return s.IdentifiedPath
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) GetParentUsergroupId() *string {
	return s.ParentUsergroupId
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) GetUsergroupDesc() *string {
	return s.UsergroupDesc
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) GetUsergroupId() *string {
	return s.UsergroupId
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) GetUsergroupName() *string {
	return s.UsergroupName
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) SetCreateTime(v string) *ListByUserGroupIdResponseBodyResultUserGroupModels {
	s.CreateTime = &v
	return s
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) SetCreateUser(v string) *ListByUserGroupIdResponseBodyResultUserGroupModels {
	s.CreateUser = &v
	return s
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) SetIdentifiedPath(v string) *ListByUserGroupIdResponseBodyResultUserGroupModels {
	s.IdentifiedPath = &v
	return s
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) SetModifiedTime(v string) *ListByUserGroupIdResponseBodyResultUserGroupModels {
	s.ModifiedTime = &v
	return s
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) SetModifyUser(v string) *ListByUserGroupIdResponseBodyResultUserGroupModels {
	s.ModifyUser = &v
	return s
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) SetParentUsergroupId(v string) *ListByUserGroupIdResponseBodyResultUserGroupModels {
	s.ParentUsergroupId = &v
	return s
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) SetUsergroupDesc(v string) *ListByUserGroupIdResponseBodyResultUserGroupModels {
	s.UsergroupDesc = &v
	return s
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) SetUsergroupId(v string) *ListByUserGroupIdResponseBodyResultUserGroupModels {
	s.UsergroupId = &v
	return s
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) SetUsergroupName(v string) *ListByUserGroupIdResponseBodyResultUserGroupModels {
	s.UsergroupName = &v
	return s
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) Validate() error {
	return dara.Validate(s)
}
