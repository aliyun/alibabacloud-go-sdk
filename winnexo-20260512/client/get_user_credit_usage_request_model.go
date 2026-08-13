// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserCreditUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *GetUserCreditUsageRequest
	GetTenantId() *string
}

type GetUserCreditUsageRequest struct {
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetUserCreditUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserCreditUsageRequest) GoString() string {
	return s.String()
}

func (s *GetUserCreditUsageRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetUserCreditUsageRequest) SetTenantId(v string) *GetUserCreditUsageRequest {
	s.TenantId = &v
	return s
}

func (s *GetUserCreditUsageRequest) Validate() error {
	return dara.Validate(s)
}
