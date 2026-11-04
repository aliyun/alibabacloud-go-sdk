// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsolationId(v string) *DeleteFaceUserGroupRequest
	GetIsolationId() *string
	SetUserGroupId(v string) *DeleteFaceUserGroupRequest
	GetUserGroupId() *string
}

type DeleteFaceUserGroupRequest struct {
	// The business isolation ID, which is used to isolate data between different business applications under the same Alibaba Cloud account. This value must be the same as the **IsolationId*	- specified when you call the [AddFaceUserGroup](https://help.aliyun.com/document_detail/288341.html) operation to add a user.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZheJiangHZ
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	// The ID of the user group. Specify the user group ID that you obtained and saved locally after calling the [AddFaceUserGroup](https://help.aliyun.com/document_detail/288341.html) operation to add the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// kvjy****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DeleteFaceUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceUserGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserGroupRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *DeleteFaceUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *DeleteFaceUserGroupRequest) SetIsolationId(v string) *DeleteFaceUserGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *DeleteFaceUserGroupRequest) SetUserGroupId(v string) *DeleteFaceUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *DeleteFaceUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
