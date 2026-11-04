// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsolationId(v string) *AddFaceUserGroupRequest
	GetIsolationId() *string
	SetUserGroupName(v string) *AddFaceUserGroupRequest
	GetUserGroupName() *string
}

type AddFaceUserGroupRequest struct {
	// The business isolation ID. It is used to isolate data between different business applications under the same Alibaba Cloud account. It supports uppercase and lowercase English letters and digits.
	//
	// For example, if you have facial recognition businesses in Region A and Region B, you can set different business isolation IDs to distinguish the facial recognition businesses in Region A and Region B.
	//
	// 	Notice:
	//
	// Save the business isolation ID properly for use when you call other API operations in this business isolation zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZheJiangHZ
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	// The name of the user group. Specify a custom value. The value can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestUserGroup
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s AddFaceUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFaceUserGroupRequest) GoString() string {
	return s.String()
}

func (s *AddFaceUserGroupRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *AddFaceUserGroupRequest) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *AddFaceUserGroupRequest) SetIsolationId(v string) *AddFaceUserGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *AddFaceUserGroupRequest) SetUserGroupName(v string) *AddFaceUserGroupRequest {
	s.UserGroupName = &v
	return s
}

func (s *AddFaceUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
