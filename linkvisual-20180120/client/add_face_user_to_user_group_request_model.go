// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceUserToUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsolationId(v string) *AddFaceUserToUserGroupRequest
	GetIsolationId() *string
	SetUserGroupId(v string) *AddFaceUserToUserGroupRequest
	GetUserGroupId() *string
	SetUserId(v string) *AddFaceUserToUserGroupRequest
	GetUserId() *string
}

type AddFaceUserToUserGroupRequest struct {
	// The business isolation ID. It is used to isolate data between different business applications under the same Alibaba Cloud account.
	//
	// 	Notice:
	//
	// The user and the user group must be under the same isolation ID. Make sure that this parameter is consistent with the **IsolationId*	- parameter that you specified when you called the [AddFaceUser](https://help.aliyun.com/document_detail/288338.html) operation to add the user and the [AddFaceUserGroup](https://help.aliyun.com/document_detail/288341.html) operation to create the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZheJiangHZ
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	// The user group ID. Set this parameter to the user group ID that you obtained and saved locally after you called the [AddFaceUserGroup](https://help.aliyun.com/document_detail/288341.html) operation to create the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// kvjy****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The user ID. Set this parameter to the user ID that you obtained and saved locally after you called the [AddFaceUser](https://help.aliyun.com/document_detail/288338.html) operation to add the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3mtjaqebeq4d****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddFaceUserToUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFaceUserToUserGroupRequest) GoString() string {
	return s.String()
}

func (s *AddFaceUserToUserGroupRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *AddFaceUserToUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *AddFaceUserToUserGroupRequest) GetUserId() *string {
	return s.UserId
}

func (s *AddFaceUserToUserGroupRequest) SetIsolationId(v string) *AddFaceUserToUserGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *AddFaceUserToUserGroupRequest) SetUserGroupId(v string) *AddFaceUserToUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *AddFaceUserToUserGroupRequest) SetUserId(v string) *AddFaceUserToUserGroupRequest {
	s.UserId = &v
	return s
}

func (s *AddFaceUserToUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
