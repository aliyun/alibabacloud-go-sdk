// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumStorageMetricsByTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v int64) *SumStorageMetricsByTypeRequest
	GetEndDate() *int64
	SetProjectNames(v []*string) *SumStorageMetricsByTypeRequest
	GetProjectNames() []*string
	SetStartDate(v int64) *SumStorageMetricsByTypeRequest
	GetStartDate() *int64
	SetStatsType(v string) *SumStorageMetricsByTypeRequest
	GetStatsType() *string
}

type SumStorageMetricsByTypeRequest struct {
	// Required. The query end time, specified as a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1718590596556
	EndDate *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// An array of project names.
	ProjectNames []*string `json:"projectNames,omitempty" xml:"projectNames,omitempty" type:"Repeated"`
	// Required. The query start time, specified as a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1715393576201
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// The dimension for aggregating statistics. Valid values: `PROJECT` (by project) and `STORAGE_TYPE` (by storage type).
	//
	// example:
	//
	// PROJECT
	StatsType *string `json:"statsType,omitempty" xml:"statsType,omitempty"`
}

func (s SumStorageMetricsByTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s SumStorageMetricsByTypeRequest) GoString() string {
	return s.String()
}

func (s *SumStorageMetricsByTypeRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *SumStorageMetricsByTypeRequest) GetProjectNames() []*string {
	return s.ProjectNames
}

func (s *SumStorageMetricsByTypeRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *SumStorageMetricsByTypeRequest) GetStatsType() *string {
	return s.StatsType
}

func (s *SumStorageMetricsByTypeRequest) SetEndDate(v int64) *SumStorageMetricsByTypeRequest {
	s.EndDate = &v
	return s
}

func (s *SumStorageMetricsByTypeRequest) SetProjectNames(v []*string) *SumStorageMetricsByTypeRequest {
	s.ProjectNames = v
	return s
}

func (s *SumStorageMetricsByTypeRequest) SetStartDate(v int64) *SumStorageMetricsByTypeRequest {
	s.StartDate = &v
	return s
}

func (s *SumStorageMetricsByTypeRequest) SetStatsType(v string) *SumStorageMetricsByTypeRequest {
	s.StatsType = &v
	return s
}

func (s *SumStorageMetricsByTypeRequest) Validate() error {
	return dara.Validate(s)
}
