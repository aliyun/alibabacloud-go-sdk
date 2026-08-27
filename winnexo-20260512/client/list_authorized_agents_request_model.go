// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedAgentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPermission(v string) *ListAuthorizedAgentsRequest
	GetPermission() *string
	SetTargetUserId(v int64) *ListAuthorizedAgentsRequest
	GetTargetUserId() *int64
	SetTenantId(v string) *ListAuthorizedAgentsRequest
	GetTenantId() *string
}

type ListAuthorizedAgentsRequest struct {
	// The userId of the responsible user.
	//
	// example:
	//
	// USE
	Permission *string `json:"permission,omitempty" xml:"permission,omitempty"`
	// The target user ID.
	//
	// example:
	//
	// 1
	TargetUserId *int64 `json:"targetUserId,omitempty" xml:"targetUserId,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListAuthorizedAgentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedAgentsRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedAgentsRequest) GetPermission() *string {
	return s.Permission
}

func (s *ListAuthorizedAgentsRequest) GetTargetUserId() *int64 {
	return s.TargetUserId
}

func (s *ListAuthorizedAgentsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAuthorizedAgentsRequest) SetPermission(v string) *ListAuthorizedAgentsRequest {
	s.Permission = &v
	return s
}

func (s *ListAuthorizedAgentsRequest) SetTargetUserId(v int64) *ListAuthorizedAgentsRequest {
	s.TargetUserId = &v
	return s
}

func (s *ListAuthorizedAgentsRequest) SetTenantId(v string) *ListAuthorizedAgentsRequest {
	s.TenantId = &v
	return s
}

func (s *ListAuthorizedAgentsRequest) Validate() error {
	return dara.Validate(s)
}
