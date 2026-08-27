// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantAgentUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireDate(v int64) *GrantAgentUsersRequest
	GetExpireDate() *int64
	SetOperatingObjectName(v string) *GrantAgentUsersRequest
	GetOperatingObjectName() *string
	SetPermissions(v []*string) *GrantAgentUsersRequest
	GetPermissions() []*string
	SetTenantId(v string) *GrantAgentUsersRequest
	GetTenantId() *string
	SetUserGroupIds(v []*string) *GrantAgentUsersRequest
	GetUserGroupIds() []*string
	SetUserIds(v []*string) *GrantAgentUsersRequest
	GetUserIds() []*string
}

type GrantAgentUsersRequest struct {
	// The authorization expiration timestamp in milliseconds. If this parameter is not specified, the authorization never expires.
	//
	// example:
	//
	// 1
	ExpireDate *int64 `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
	// The name of the digital human.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The permission items.
	//
	// example:
	//
	// string_value
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// example:
	//
	// 676577544219585
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The list of user group IDs.
	//
	// example:
	//
	// string_value
	UserGroupIds []*string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty" type:"Repeated"`
	// The list of user IDs to be authorized.
	//
	// example:
	//
	// 1
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s GrantAgentUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantAgentUsersRequest) GoString() string {
	return s.String()
}

func (s *GrantAgentUsersRequest) GetExpireDate() *int64 {
	return s.ExpireDate
}

func (s *GrantAgentUsersRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *GrantAgentUsersRequest) GetPermissions() []*string {
	return s.Permissions
}

func (s *GrantAgentUsersRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GrantAgentUsersRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *GrantAgentUsersRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *GrantAgentUsersRequest) SetExpireDate(v int64) *GrantAgentUsersRequest {
	s.ExpireDate = &v
	return s
}

func (s *GrantAgentUsersRequest) SetOperatingObjectName(v string) *GrantAgentUsersRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *GrantAgentUsersRequest) SetPermissions(v []*string) *GrantAgentUsersRequest {
	s.Permissions = v
	return s
}

func (s *GrantAgentUsersRequest) SetTenantId(v string) *GrantAgentUsersRequest {
	s.TenantId = &v
	return s
}

func (s *GrantAgentUsersRequest) SetUserGroupIds(v []*string) *GrantAgentUsersRequest {
	s.UserGroupIds = v
	return s
}

func (s *GrantAgentUsersRequest) SetUserIds(v []*string) *GrantAgentUsersRequest {
	s.UserIds = v
	return s
}

func (s *GrantAgentUsersRequest) Validate() error {
	return dara.Validate(s)
}
