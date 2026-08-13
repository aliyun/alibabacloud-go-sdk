// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *GetTokenInfoRequest
	GetTenantId() *string
	SetWnUserId(v string) *GetTokenInfoRequest
	GetWnUserId() *string
}

type GetTokenInfoRequest struct {
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 目标用户 ID（WINNEXO 平台用户ID，空则操作自身，管理员可传入他人 ID 代操作）
	//
	// example:
	//
	// 1
	WnUserId *string `json:"wnUserId,omitempty" xml:"wnUserId,omitempty"`
}

func (s GetTokenInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTokenInfoRequest) GoString() string {
	return s.String()
}

func (s *GetTokenInfoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetTokenInfoRequest) GetWnUserId() *string {
	return s.WnUserId
}

func (s *GetTokenInfoRequest) SetTenantId(v string) *GetTokenInfoRequest {
	s.TenantId = &v
	return s
}

func (s *GetTokenInfoRequest) SetWnUserId(v string) *GetTokenInfoRequest {
	s.WnUserId = &v
	return s
}

func (s *GetTokenInfoRequest) Validate() error {
	return dara.Validate(s)
}
