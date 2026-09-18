// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSqlPatternCompareReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompareEndTime(v string) *CreateSqlPatternCompareReportRequest
	GetCompareEndTime() *string
	SetCompareStartTime(v string) *CreateSqlPatternCompareReportRequest
	GetCompareStartTime() *string
	SetDBClusterId(v string) *CreateSqlPatternCompareReportRequest
	GetDBClusterId() *string
	SetEndTime(v string) *CreateSqlPatternCompareReportRequest
	GetEndTime() *string
	SetPatternType(v string) *CreateSqlPatternCompareReportRequest
	GetPatternType() *string
	SetRegionId(v string) *CreateSqlPatternCompareReportRequest
	GetRegionId() *string
	SetStartTime(v string) *CreateSqlPatternCompareReportRequest
	GetStartTime() *string
}

type CreateSqlPatternCompareReportRequest struct {
	// The end time of time window 2. Specify the time in UTC in the yyyy-MM-ddTHH:mmZ or yyyy-MM-ddTHH:mm:ssZ format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-08T01:00Z
	CompareEndTime *string `json:"CompareEndTime,omitempty" xml:"CompareEndTime,omitempty"`
	// The start time of time window 2. Specify the time in UTC in the yyyy-MM-ddTHH:mmZ or yyyy-MM-ddTHH:mm:ssZ format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-08T00:00Z
	CompareStartTime *string `json:"CompareStartTime,omitempty" xml:"CompareStartTime,omitempty"`
	// The ID of the AnalyticDB for MySQL cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-2ze1234567890****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of time window 1. Specify the time in UTC in the yyyy-MM-ddTHH:mmZ or yyyy-MM-ddTHH:mm:ssZ format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-07T01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The report type. Valid values:
	//
	// - `NEW`: Patterns that are new in time window 2.
	//
	// - `CHANGED`: Patterns that exist in both time windows and have an increased average value in at least one metric.
	//
	// This parameter is required.
	//
	// example:
	//
	// CHANGED
	PatternType *string `json:"PatternType,omitempty" xml:"PatternType,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of time window 1. Specify the time in UTC in the yyyy-MM-ddTHH:mmZ or yyyy-MM-ddTHH:mm:ssZ format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-07T00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateSqlPatternCompareReportRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlPatternCompareReportRequest) GoString() string {
	return s.String()
}

func (s *CreateSqlPatternCompareReportRequest) GetCompareEndTime() *string {
	return s.CompareEndTime
}

func (s *CreateSqlPatternCompareReportRequest) GetCompareStartTime() *string {
	return s.CompareStartTime
}

func (s *CreateSqlPatternCompareReportRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateSqlPatternCompareReportRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateSqlPatternCompareReportRequest) GetPatternType() *string {
	return s.PatternType
}

func (s *CreateSqlPatternCompareReportRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSqlPatternCompareReportRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateSqlPatternCompareReportRequest) SetCompareEndTime(v string) *CreateSqlPatternCompareReportRequest {
	s.CompareEndTime = &v
	return s
}

func (s *CreateSqlPatternCompareReportRequest) SetCompareStartTime(v string) *CreateSqlPatternCompareReportRequest {
	s.CompareStartTime = &v
	return s
}

func (s *CreateSqlPatternCompareReportRequest) SetDBClusterId(v string) *CreateSqlPatternCompareReportRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateSqlPatternCompareReportRequest) SetEndTime(v string) *CreateSqlPatternCompareReportRequest {
	s.EndTime = &v
	return s
}

func (s *CreateSqlPatternCompareReportRequest) SetPatternType(v string) *CreateSqlPatternCompareReportRequest {
	s.PatternType = &v
	return s
}

func (s *CreateSqlPatternCompareReportRequest) SetRegionId(v string) *CreateSqlPatternCompareReportRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSqlPatternCompareReportRequest) SetStartTime(v string) *CreateSqlPatternCompareReportRequest {
	s.StartTime = &v
	return s
}

func (s *CreateSqlPatternCompareReportRequest) Validate() error {
	return dara.Validate(s)
}
