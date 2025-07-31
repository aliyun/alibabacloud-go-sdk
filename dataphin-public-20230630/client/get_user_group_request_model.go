// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetUserGroupRequest
	GetOpTenantId() *int64
	SetUserGroupId(v string) *GetUserGroupRequest
	GetUserGroupId() *string
}

type GetUserGroupRequest struct {
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
	// 1313213
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s GetUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserGroupRequest) GoString() string {
	return s.String()
}

func (s *GetUserGroupRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *GetUserGroupRequest) SetOpTenantId(v int64) *GetUserGroupRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetUserGroupRequest) SetUserGroupId(v string) *GetUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *GetUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
