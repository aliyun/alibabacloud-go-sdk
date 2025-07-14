// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserGroupListByParentIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParentUserGroupId(v string) *QueryUserGroupListByParentIdRequest
	GetParentUserGroupId() *string
}

type QueryUserGroupListByParentIdRequest struct {
	// The ID of the parent user group.
	//
	// 	- If you enter the ID of the parent user group, you can obtain the information of the child user group under this ID.
	//
	// 	- If you enter -1, you can obtain the sub-user group information under the root directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3d2c23d4-2b41-4af8-a1f5-f6390f32****
	ParentUserGroupId *string `json:"ParentUserGroupId,omitempty" xml:"ParentUserGroupId,omitempty"`
}

func (s QueryUserGroupListByParentIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUserGroupListByParentIdRequest) GoString() string {
	return s.String()
}

func (s *QueryUserGroupListByParentIdRequest) GetParentUserGroupId() *string {
	return s.ParentUserGroupId
}

func (s *QueryUserGroupListByParentIdRequest) SetParentUserGroupId(v string) *QueryUserGroupListByParentIdRequest {
	s.ParentUserGroupId = &v
	return s
}

func (s *QueryUserGroupListByParentIdRequest) Validate() error {
	return dara.Validate(s)
}
