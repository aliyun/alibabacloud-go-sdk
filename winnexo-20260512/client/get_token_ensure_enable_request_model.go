// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenEnsureEnableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *GetTokenEnsureEnableRequest
	GetTenantId() *string
	SetWnUserId(v string) *GetTokenEnsureEnableRequest
	GetWnUserId() *string
}

type GetTokenEnsureEnableRequest struct {
	// The ID of the tenant to which the task belongs.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The WinNexo user ID.
	//
	// example:
	//
	// 1
	WnUserId *string `json:"wnUserId,omitempty" xml:"wnUserId,omitempty"`
}

func (s GetTokenEnsureEnableRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTokenEnsureEnableRequest) GoString() string {
	return s.String()
}

func (s *GetTokenEnsureEnableRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetTokenEnsureEnableRequest) GetWnUserId() *string {
	return s.WnUserId
}

func (s *GetTokenEnsureEnableRequest) SetTenantId(v string) *GetTokenEnsureEnableRequest {
	s.TenantId = &v
	return s
}

func (s *GetTokenEnsureEnableRequest) SetWnUserId(v string) *GetTokenEnsureEnableRequest {
	s.WnUserId = &v
	return s
}

func (s *GetTokenEnsureEnableRequest) Validate() error {
	return dara.Validate(s)
}
