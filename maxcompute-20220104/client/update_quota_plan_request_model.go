// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQuotaPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdateQuotaPlanRequest
	GetBody() *string
	SetRegion(v string) *UpdateQuotaPlanRequest
	GetRegion() *string
	SetTenantId(v string) *UpdateQuotaPlanRequest
	GetTenantId() *string
}

type UpdateQuotaPlanRequest struct {
	Body     *string `json:"body,omitempty" xml:"body,omitempty"`
	Region   *string `json:"region,omitempty" xml:"region,omitempty"`
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateQuotaPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateQuotaPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateQuotaPlanRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateQuotaPlanRequest) GetRegion() *string {
	return s.Region
}

func (s *UpdateQuotaPlanRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateQuotaPlanRequest) SetBody(v string) *UpdateQuotaPlanRequest {
	s.Body = &v
	return s
}

func (s *UpdateQuotaPlanRequest) SetRegion(v string) *UpdateQuotaPlanRequest {
	s.Region = &v
	return s
}

func (s *UpdateQuotaPlanRequest) SetTenantId(v string) *UpdateQuotaPlanRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateQuotaPlanRequest) Validate() error {
	return dara.Validate(s)
}
