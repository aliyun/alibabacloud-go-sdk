// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLakebaseMountParameterGroupChangeLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeLakebaseMountParameterGroupChangeLogsRequest
	GetEndTime() *string
	SetName(v string) *DescribeLakebaseMountParameterGroupChangeLogsRequest
	GetName() *string
	SetParameterGroupId(v string) *DescribeLakebaseMountParameterGroupChangeLogsRequest
	GetParameterGroupId() *string
	SetPolarFsInstanceId(v string) *DescribeLakebaseMountParameterGroupChangeLogsRequest
	GetPolarFsInstanceId() *string
	SetRegionId(v string) *DescribeLakebaseMountParameterGroupChangeLogsRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLakebaseMountParameterGroupChangeLogsRequest
	GetStartTime() *string
}

type DescribeLakebaseMountParameterGroupChangeLogsRequest struct {
	// The end time of the query (UTC).
	//
	// example:
	//
	// 2026-01-02T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The parameter name filter.
	//
	// example:
	//
	// param1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-xxx
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The PolarFS instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pfs-xxx
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of the query (UTC).
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLakebaseMountParameterGroupChangeLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakebaseMountParameterGroupChangeLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsRequest) SetEndTime(v string) *DescribeLakebaseMountParameterGroupChangeLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsRequest) SetName(v string) *DescribeLakebaseMountParameterGroupChangeLogsRequest {
	s.Name = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsRequest) SetParameterGroupId(v string) *DescribeLakebaseMountParameterGroupChangeLogsRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsRequest) SetPolarFsInstanceId(v string) *DescribeLakebaseMountParameterGroupChangeLogsRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsRequest) SetRegionId(v string) *DescribeLakebaseMountParameterGroupChangeLogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsRequest) SetStartTime(v string) *DescribeLakebaseMountParameterGroupChangeLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsRequest) Validate() error {
	return dara.Validate(s)
}
