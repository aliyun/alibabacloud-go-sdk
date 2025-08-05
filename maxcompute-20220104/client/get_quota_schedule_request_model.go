// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayTimezone(v string) *GetQuotaScheduleRequest
	GetDisplayTimezone() *string
	SetRegion(v string) *GetQuotaScheduleRequest
	GetRegion() *string
	SetTenantId(v string) *GetQuotaScheduleRequest
	GetTenantId() *string
}

type GetQuotaScheduleRequest struct {
	// The time zone.
	//
	// example:
	//
	// UTC+8
	DisplayTimezone *string `json:"displayTimezone,omitempty" xml:"displayTimezone,omitempty"`
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

func (s GetQuotaScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaScheduleRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaScheduleRequest) GetDisplayTimezone() *string {
	return s.DisplayTimezone
}

func (s *GetQuotaScheduleRequest) GetRegion() *string {
	return s.Region
}

func (s *GetQuotaScheduleRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetQuotaScheduleRequest) SetDisplayTimezone(v string) *GetQuotaScheduleRequest {
	s.DisplayTimezone = &v
	return s
}

func (s *GetQuotaScheduleRequest) SetRegion(v string) *GetQuotaScheduleRequest {
	s.Region = &v
	return s
}

func (s *GetQuotaScheduleRequest) SetTenantId(v string) *GetQuotaScheduleRequest {
	s.TenantId = &v
	return s
}

func (s *GetQuotaScheduleRequest) Validate() error {
	return dara.Validate(s)
}
