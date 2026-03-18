// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQuotaScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdateQuotaScheduleRequest
	GetBody() *string
	SetRegion(v string) *UpdateQuotaScheduleRequest
	GetRegion() *string
	SetTenantId(v string) *UpdateQuotaScheduleRequest
	GetTenantId() *string
}

type UpdateQuotaScheduleRequest struct {
	Body     *string `json:"body,omitempty" xml:"body,omitempty"`
	Region   *string `json:"region,omitempty" xml:"region,omitempty"`
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateQuotaScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateQuotaScheduleRequest) GoString() string {
	return s.String()
}

func (s *UpdateQuotaScheduleRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateQuotaScheduleRequest) GetRegion() *string {
	return s.Region
}

func (s *UpdateQuotaScheduleRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateQuotaScheduleRequest) SetBody(v string) *UpdateQuotaScheduleRequest {
	s.Body = &v
	return s
}

func (s *UpdateQuotaScheduleRequest) SetRegion(v string) *UpdateQuotaScheduleRequest {
	s.Region = &v
	return s
}

func (s *UpdateQuotaScheduleRequest) SetTenantId(v string) *UpdateQuotaScheduleRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateQuotaScheduleRequest) Validate() error {
	return dara.Validate(s)
}
