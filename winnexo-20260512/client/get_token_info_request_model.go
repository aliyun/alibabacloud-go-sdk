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
	// The ID of the tenant to which the task belongs.
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
