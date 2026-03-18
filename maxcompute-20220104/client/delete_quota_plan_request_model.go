// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQuotaPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegion(v string) *DeleteQuotaPlanRequest
	GetRegion() *string
	SetTenantId(v string) *DeleteQuotaPlanRequest
	GetTenantId() *string
}

type DeleteQuotaPlanRequest struct {
	Region   *string `json:"region,omitempty" xml:"region,omitempty"`
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteQuotaPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQuotaPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteQuotaPlanRequest) GetRegion() *string {
	return s.Region
}

func (s *DeleteQuotaPlanRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteQuotaPlanRequest) SetRegion(v string) *DeleteQuotaPlanRequest {
	s.Region = &v
	return s
}

func (s *DeleteQuotaPlanRequest) SetTenantId(v string) *DeleteQuotaPlanRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteQuotaPlanRequest) Validate() error {
	return dara.Validate(s)
}
