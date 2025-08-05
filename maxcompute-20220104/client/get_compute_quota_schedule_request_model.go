// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeQuotaScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayTimezone(v string) *GetComputeQuotaScheduleRequest
	GetDisplayTimezone() *string
}

type GetComputeQuotaScheduleRequest struct {
	// Display time zone.
	//
	// example:
	//
	// UTC+8
	DisplayTimezone *string `json:"displayTimezone,omitempty" xml:"displayTimezone,omitempty"`
}

func (s GetComputeQuotaScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetComputeQuotaScheduleRequest) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaScheduleRequest) GetDisplayTimezone() *string {
	return s.DisplayTimezone
}

func (s *GetComputeQuotaScheduleRequest) SetDisplayTimezone(v string) *GetComputeQuotaScheduleRequest {
	s.DisplayTimezone = &v
	return s
}

func (s *GetComputeQuotaScheduleRequest) Validate() error {
	return dara.Validate(s)
}
