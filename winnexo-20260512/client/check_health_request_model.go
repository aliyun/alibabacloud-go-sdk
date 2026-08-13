// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckHealthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *CheckHealthRequest
	GetTenantId() *string
}

type CheckHealthRequest struct {
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CheckHealthRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckHealthRequest) GoString() string {
	return s.String()
}

func (s *CheckHealthRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CheckHealthRequest) SetTenantId(v string) *CheckHealthRequest {
	s.TenantId = &v
	return s
}

func (s *CheckHealthRequest) Validate() error {
	return dara.Validate(s)
}
