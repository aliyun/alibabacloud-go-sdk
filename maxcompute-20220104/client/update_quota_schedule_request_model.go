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
	// The request body parameters.
	//
	// example:
	//
	// \\# The quota plan immediately takes effect. [ { "type": "once", "plan": "planA", "operator":"userA" } ] # The quota plan is scheduled on a regular basis. [ { "id": "etl_time", "type": "daily", "condition": { "at": "0800", "after": "2022-04-25T04:23:04Z" // optional }, "plan": "planA" }, { "id": "bi", "type": "daily", "condition": { "at": "0900", "after": "2022-04-25T04:23:04Z" // optional }, "plan": "planB" }, ]
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-chengdu
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 478403690625249
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
