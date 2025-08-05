// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotasPlansRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegion(v string) *ListQuotasPlansRequest
	GetRegion() *string
	SetTenantId(v string) *ListQuotasPlansRequest
	GetTenantId() *string
}

type ListQuotasPlansRequest struct {
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
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListQuotasPlansRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasPlansRequest) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansRequest) GetRegion() *string {
	return s.Region
}

func (s *ListQuotasPlansRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListQuotasPlansRequest) SetRegion(v string) *ListQuotasPlansRequest {
	s.Region = &v
	return s
}

func (s *ListQuotasPlansRequest) SetTenantId(v string) *ListQuotasPlansRequest {
	s.TenantId = &v
	return s
}

func (s *ListQuotasPlansRequest) Validate() error {
	return dara.Validate(s)
}
