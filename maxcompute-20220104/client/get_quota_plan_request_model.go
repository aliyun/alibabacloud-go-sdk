// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegion(v string) *GetQuotaPlanRequest
	GetRegion() *string
	SetTenantId(v string) *GetQuotaPlanRequest
	GetTenantId() *string
}

type GetQuotaPlanRequest struct {
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 483212237127906
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetQuotaPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaPlanRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanRequest) GetRegion() *string {
	return s.Region
}

func (s *GetQuotaPlanRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetQuotaPlanRequest) SetRegion(v string) *GetQuotaPlanRequest {
	s.Region = &v
	return s
}

func (s *GetQuotaPlanRequest) SetTenantId(v string) *GetQuotaPlanRequest {
	s.TenantId = &v
	return s
}

func (s *GetQuotaPlanRequest) Validate() error {
	return dara.Validate(s)
}
