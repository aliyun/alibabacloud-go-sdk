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
	// The ID of the tenant to which the task belongs.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The WINNEXO logon account. This is a unique identifier and cannot be empty.
	//
	// example:
	//
	// exampleAccountId
	WnAccountId *string `json:"wnAccountId,omitempty" xml:"wnAccountId,omitempty"`
	// The WINNEXO platform user ID. Specify either this parameter or accountId.
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
