// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsByUserIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListUserGroupsByUserIdResponseBody
	GetRequestId() *string
	SetResult(v []*ListUserGroupsByUserIdResponseBodyResult) *ListUserGroupsByUserIdResponseBody
	GetResult() []*ListUserGroupsByUserIdResponseBodyResult
	SetSuccess(v bool) *ListUserGroupsByUserIdResponseBody
	GetSuccess() *bool
}

type ListUserGroupsByUserIdResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E2440604-3059-561A-AD68-DEDBC870EB2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the group.
	Result []*ListUserGroupsByUserIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
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

func (s ListUserGroupsByUserIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserGroupsByUserIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserGroupsByUserIdResponseBody) GetResult() []*ListUserGroupsByUserIdResponseBodyResult {
	return s.Result
}

func (s *ListUserGroupsByUserIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUserGroupsByUserIdResponseBody) SetRequestId(v string) *ListUserGroupsByUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserGroupsByUserIdResponseBody) SetResult(v []*ListUserGroupsByUserIdResponseBodyResult) *ListUserGroupsByUserIdResponseBody {
	s.Result = v
	return s
}

func (s *ListUserGroupsByUserIdResponseBody) SetSuccess(v bool) *ListUserGroupsByUserIdResponseBody {
	s.Success = &v
	return s
}

func (s *ListUserGroupsByUserIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUserGroupsByUserIdResponseBodyResult struct {
	// The time when the user group was created.
	//
	// example:
	//
	// 2021-03-15 17:13:55
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The user group creator. The UserID of the Quick BI is used instead of the UID of Alibaba Cloud.
	//
	// example:
	//
	// 46e5374665ba4b679ee22e2a2927****
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// Directory level of the user group.
	IdentifiedPath *string `json:"IdentifiedPath,omitempty" xml:"IdentifiedPath,omitempty"`
	// The time when the user group was last modified.
	//
	// example:
	//
	// 2021-03-15 20:36:40
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The user group modifier. The UserID of the Quick BI is used instead of the UID of Alibaba Cloud.
	//
	// example:
	//
	// 46e5374665ba4b679ee22e2a2927****
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The ID of the user group.
	//
	// example:
	//
	// 2fe4fbd8-588f-489a-b3e1-e92c7af0****
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
	// 34fd141d-4598-4093-8c33-8e066dcb****
	UsergroupId *string `json:"UsergroupId,omitempty" xml:"UsergroupId,omitempty"`
	// The name of the user group.
	//
	// example:
	//
	// Test user group
	UsergroupName *string `json:"UsergroupName,omitempty" xml:"UsergroupName,omitempty"`
}

func (s ListUserGroupsByUserIdResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsByUserIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListUserGroupsByUserIdResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListUserGroupsByUserIdResponseBodyResult) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListUserGroupsByUserIdResponseBodyResult) GetIdentifiedPath() *string {
	return s.IdentifiedPath
}

func (s *ListUserGroupsByUserIdResponseBodyResult) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListUserGroupsByUserIdResponseBodyResult) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListUserGroupsByUserIdResponseBodyResult) GetParentUsergroupId() *string {
	return s.ParentUsergroupId
}

func (s *ListUserGroupsByUserIdResponseBodyResult) GetUsergroupDesc() *string {
	return s.UsergroupDesc
}

func (s *ListUserGroupsByUserIdResponseBodyResult) GetUsergroupId() *string {
	return s.UsergroupId
}

func (s *ListUserGroupsByUserIdResponseBodyResult) GetUsergroupName() *string {
	return s.UsergroupName
}

func (s *ListUserGroupsByUserIdResponseBodyResult) SetCreateTime(v string) *ListUserGroupsByUserIdResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListUserGroupsByUserIdResponseBodyResult) SetCreateUser(v string) *ListUserGroupsByUserIdResponseBodyResult {
	s.CreateUser = &v
	return s
}

func (s *ListUserGroupsByUserIdResponseBodyResult) SetIdentifiedPath(v string) *ListUserGroupsByUserIdResponseBodyResult {
	s.IdentifiedPath = &v
	return s
}

func (s *ListUserGroupsByUserIdResponseBodyResult) SetModifiedTime(v string) *ListUserGroupsByUserIdResponseBodyResult {
	s.ModifiedTime = &v
	return s
}

func (s *ListUserGroupsByUserIdResponseBodyResult) SetModifyUser(v string) *ListUserGroupsByUserIdResponseBodyResult {
	s.ModifyUser = &v
	return s
}

func (s *ListUserGroupsByUserIdResponseBodyResult) SetParentUsergroupId(v string) *ListUserGroupsByUserIdResponseBodyResult {
	s.ParentUsergroupId = &v
	return s
}

func (s *ListUserGroupsByUserIdResponseBodyResult) SetUsergroupDesc(v string) *ListUserGroupsByUserIdResponseBodyResult {
	s.UsergroupDesc = &v
	return s
}

func (s *ListUserGroupsByUserIdResponseBodyResult) SetUsergroupId(v string) *ListUserGroupsByUserIdResponseBodyResult {
	s.UsergroupId = &v
	return s
}

func (s *ListUserGroupsByUserIdResponseBodyResult) SetUsergroupName(v string) *ListUserGroupsByUserIdResponseBodyResult {
	s.UsergroupName = &v
	return s
}

func (s *ListUserGroupsByUserIdResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
