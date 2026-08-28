// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPatrolReportDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v string) *GetPatrolReportDetailRequest
	GetDate() *string
	SetReportId(v string) *GetPatrolReportDetailRequest
	GetReportId() *string
	SetTimezone(v string) *GetPatrolReportDetailRequest
	GetTimezone() *string
}

type GetPatrolReportDetailRequest struct {
	// The report date in ISO format such as 2026-08-13, or a special value such as today or yesterday. Specify either this parameter or reportId. This parameter is used to retrieve the latest report for the specified date.
	//
	// example:
	//
	// 20260416
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The report ID. Specify either this parameter or date. The reportId parameter takes priority over the date parameter.
	//
	// example:
	//
	// d72edfd9f3ef475ca310456ecdb09fe8
	ReportId *string `json:"reportId,omitempty" xml:"reportId,omitempty"`
	// The time zone. Used together with the date parameter. Default value: UTC.
	//
	// example:
	//
	// UTC
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetPatrolReportDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPatrolReportDetailRequest) GoString() string {
	return s.String()
}

func (s *GetPatrolReportDetailRequest) GetDate() *string {
	return s.Date
}

func (s *GetPatrolReportDetailRequest) GetReportId() *string {
	return s.ReportId
}

func (s *GetPatrolReportDetailRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *GetPatrolReportDetailRequest) SetDate(v string) *GetPatrolReportDetailRequest {
	s.Date = &v
	return s
}

func (s *GetPatrolReportDetailRequest) SetReportId(v string) *GetPatrolReportDetailRequest {
	s.ReportId = &v
	return s
}

func (s *GetPatrolReportDetailRequest) SetTimezone(v string) *GetPatrolReportDetailRequest {
	s.Timezone = &v
	return s
}

func (s *GetPatrolReportDetailRequest) Validate() error {
	return dara.Validate(s)
}
