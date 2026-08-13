// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *ResetTokenRequest
	GetTenantId() *string
	SetWnUserId(v string) *ResetTokenRequest
	GetWnUserId() *string
}

type ResetTokenRequest struct {
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

func (s ResetTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetTokenRequest) GoString() string {
	return s.String()
}

func (s *ResetTokenRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ResetTokenRequest) GetWnUserId() *string {
	return s.WnUserId
}

func (s *ResetTokenRequest) SetTenantId(v string) *ResetTokenRequest {
	s.TenantId = &v
	return s
}

func (s *ResetTokenRequest) SetWnUserId(v string) *ResetTokenRequest {
	s.WnUserId = &v
	return s
}

func (s *ResetTokenRequest) Validate() error {
	return dara.Validate(s)
}
