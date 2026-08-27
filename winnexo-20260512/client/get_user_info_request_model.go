// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *GetUserInfoRequest
	GetTenantId() *string
}

type GetUserInfoRequest struct {
	// The effective tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetUserInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserInfoRequest) GoString() string {
	return s.String()
}

func (s *GetUserInfoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetUserInfoRequest) SetTenantId(v string) *GetUserInfoRequest {
	s.TenantId = &v
	return s
}

func (s *GetUserInfoRequest) Validate() error {
	return dara.Validate(s)
}
