// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *DeleteUserGroupRequest
	GetOpTenantId() *int64
	SetUserGroupId(v string) *DeleteUserGroupRequest
	GetUserGroupId() *string
}

type DeleteUserGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 232131
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DeleteUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *DeleteUserGroupRequest) SetOpTenantId(v int64) *DeleteUserGroupRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteUserGroupRequest) SetUserGroupId(v string) *DeleteUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *DeleteUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
