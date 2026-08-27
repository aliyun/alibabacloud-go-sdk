// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *GetUserGroupRequest
	GetTenantId() *string
	SetUserGroupId(v string) *GetUserGroupRequest
	GetUserGroupId() *string
}

type GetUserGroupRequest struct {
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this parameter explicitly by using `--tenant-id`.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the target user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7ea8973f-7a5c-4e8a-956b-4fe0e7e2eb11
	UserGroupId *string `json:"userGroupId,omitempty" xml:"userGroupId,omitempty"`
}

func (s GetUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserGroupRequest) GoString() string {
	return s.String()
}

func (s *GetUserGroupRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *GetUserGroupRequest) SetTenantId(v string) *GetUserGroupRequest {
	s.TenantId = &v
	return s
}

func (s *GetUserGroupRequest) SetUserGroupId(v string) *GetUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *GetUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
