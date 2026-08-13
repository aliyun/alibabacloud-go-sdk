// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *GetUserRequest
	GetTenantId() *string
	SetWnAccountId(v string) *GetUserRequest
	GetWnAccountId() *string
	SetWnUserId(v string) *GetUserRequest
	GetWnUserId() *string
}

type GetUserRequest struct {
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// WINNEXO 登录账号（与 wnUserId 二选一）
	//
	// example:
	//
	// exampleAccountId
	WnAccountId *string `json:"wnAccountId,omitempty" xml:"wnAccountId,omitempty"`
	// WINNEXO 平台用户ID（与 accountId 二选一）
	//
	// example:
	//
	// 1
	WnUserId *string `json:"wnUserId,omitempty" xml:"wnUserId,omitempty"`
}

func (s GetUserRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetUserRequest) GetWnAccountId() *string {
	return s.WnAccountId
}

func (s *GetUserRequest) GetWnUserId() *string {
	return s.WnUserId
}

func (s *GetUserRequest) SetTenantId(v string) *GetUserRequest {
	s.TenantId = &v
	return s
}

func (s *GetUserRequest) SetWnAccountId(v string) *GetUserRequest {
	s.WnAccountId = &v
	return s
}

func (s *GetUserRequest) SetWnUserId(v string) *GetUserRequest {
	s.WnUserId = &v
	return s
}

func (s *GetUserRequest) Validate() error {
	return dara.Validate(s)
}
