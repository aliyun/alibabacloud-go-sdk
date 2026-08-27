// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *DisableTokenRequest
	GetTenantId() *string
	SetWnUserId(v string) *DisableTokenRequest
	GetWnUserId() *string
}

type DisableTokenRequest struct {
	// The ID of the effective tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the target user (WINNEXO platform user ID). If left empty, the operation is performed on the caller. Administrators can specify another user\\"s ID to perform the operation on their behalf.
	//
	// example:
	//
	// 1
	WnUserId *string `json:"wnUserId,omitempty" xml:"wnUserId,omitempty"`
}

func (s DisableTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableTokenRequest) GoString() string {
	return s.String()
}

func (s *DisableTokenRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DisableTokenRequest) GetWnUserId() *string {
	return s.WnUserId
}

func (s *DisableTokenRequest) SetTenantId(v string) *DisableTokenRequest {
	s.TenantId = &v
	return s
}

func (s *DisableTokenRequest) SetWnUserId(v string) *DisableTokenRequest {
	s.WnUserId = &v
	return s
}

func (s *DisableTokenRequest) Validate() error {
	return dara.Validate(s)
}
