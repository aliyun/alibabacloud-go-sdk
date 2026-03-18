// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQuotaPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *CreateQuotaPlanRequest
	GetBody() *string
	SetRegion(v string) *CreateQuotaPlanRequest
	GetRegion() *string
	SetTenantId(v string) *CreateQuotaPlanRequest
	GetTenantId() *string
}

type CreateQuotaPlanRequest struct {
	Body     *string `json:"body,omitempty" xml:"body,omitempty"`
	Region   *string `json:"region,omitempty" xml:"region,omitempty"`
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateQuotaPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateQuotaPlanRequest) GetBody() *string {
	return s.Body
}

func (s *CreateQuotaPlanRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateQuotaPlanRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateQuotaPlanRequest) SetBody(v string) *CreateQuotaPlanRequest {
	s.Body = &v
	return s
}

func (s *CreateQuotaPlanRequest) SetRegion(v string) *CreateQuotaPlanRequest {
	s.Region = &v
	return s
}

func (s *CreateQuotaPlanRequest) SetTenantId(v string) *CreateQuotaPlanRequest {
	s.TenantId = &v
	return s
}

func (s *CreateQuotaPlanRequest) Validate() error {
	return dara.Validate(s)
}
