// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSaasDataUnitUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetSaasDataUnitUsageRequest
	GetEndTime() *string
	SetRegionId(v string) *GetSaasDataUnitUsageRequest
	GetRegionId() *string
	SetServiceId(v string) *GetSaasDataUnitUsageRequest
	GetServiceId() *string
	SetStartTime(v string) *GetSaasDataUnitUsageRequest
	GetStartTime() *string
	SetTimeUnit(v string) *GetSaasDataUnitUsageRequest
	GetTimeUnit() *string
}

type GetSaasDataUnitUsageRequest struct {
	// The end time of the query in UTC. The format is yyyy-MM-ddTHH:mm:ssZ, and the seconds must be 00. The end time must be later than the start time and no later than the request time. The query range is not automatically truncated.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-10-20T02:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID of the undeleted SaaS service. A standard DBInstanceId cannot override it.
	//
	// This parameter is required.
	//
	// example:
	//
	// agdb-example
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The start time of the query in UTC. The format is yyyy-MM-ddTHH:mm:ssZ, and the seconds must be 00. The start time cannot be earlier than 30 days before the request time. The start time and the end time form a left-closed and right-open interval, and the span cannot exceed 7 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-10-20T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Specifies the time granularity. Valid values: hour (UTC hour) and day (UTC day). Default value: hour. This parameter only determines the trend aggregation method and is not returned in the response.
	//
	// example:
	//
	// hour
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s GetSaasDataUnitUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSaasDataUnitUsageRequest) GoString() string {
	return s.String()
}

func (s *GetSaasDataUnitUsageRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetSaasDataUnitUsageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSaasDataUnitUsageRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetSaasDataUnitUsageRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetSaasDataUnitUsageRequest) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *GetSaasDataUnitUsageRequest) SetEndTime(v string) *GetSaasDataUnitUsageRequest {
	s.EndTime = &v
	return s
}

func (s *GetSaasDataUnitUsageRequest) SetRegionId(v string) *GetSaasDataUnitUsageRequest {
	s.RegionId = &v
	return s
}

func (s *GetSaasDataUnitUsageRequest) SetServiceId(v string) *GetSaasDataUnitUsageRequest {
	s.ServiceId = &v
	return s
}

func (s *GetSaasDataUnitUsageRequest) SetStartTime(v string) *GetSaasDataUnitUsageRequest {
	s.StartTime = &v
	return s
}

func (s *GetSaasDataUnitUsageRequest) SetTimeUnit(v string) *GetSaasDataUnitUsageRequest {
	s.TimeUnit = &v
	return s
}

func (s *GetSaasDataUnitUsageRequest) Validate() error {
	return dara.Validate(s)
}
