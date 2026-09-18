// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumStorageMetricsByDateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v int64) *SumStorageMetricsByDateRequest
	GetEndDate() *int64
	SetProjectNames(v []*string) *SumStorageMetricsByDateRequest
	GetProjectNames() []*string
	SetStartDate(v int64) *SumStorageMetricsByDateRequest
	GetStartDate() *int64
	SetStatsType(v string) *SumStorageMetricsByDateRequest
	GetStatsType() *string
}

type SumStorageMetricsByDateRequest struct {
	// The end of the time range to query. The value must be a timestamp in milliseconds. This parameter is required.
	//
	// example:
	//
	// 1718590596556
	EndDate *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// The project names. If this parameter is not specified, the data of all projects is returned.
	ProjectNames []*string `json:"projectNames,omitempty" xml:"projectNames,omitempty" type:"Repeated"`
	// The beginning of the time range to query. The value must be a timestamp in milliseconds. This parameter is required.
	//
	// example:
	//
	// 1715393576201
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// The dimension for aggregating storage metrics. This parameter is required. Valid values:
	//
	// - `PROJECT`
	//
	// - `STORAGE_TYPE`
	//
	// example:
	//
	// PROJECT
	StatsType *string `json:"statsType,omitempty" xml:"statsType,omitempty"`
}

func (s SumStorageMetricsByDateRequest) String() string {
	return dara.Prettify(s)
}

func (s SumStorageMetricsByDateRequest) GoString() string {
	return s.String()
}

func (s *SumStorageMetricsByDateRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *SumStorageMetricsByDateRequest) GetProjectNames() []*string {
	return s.ProjectNames
}

func (s *SumStorageMetricsByDateRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *SumStorageMetricsByDateRequest) GetStatsType() *string {
	return s.StatsType
}

func (s *SumStorageMetricsByDateRequest) SetEndDate(v int64) *SumStorageMetricsByDateRequest {
	s.EndDate = &v
	return s
}

func (s *SumStorageMetricsByDateRequest) SetProjectNames(v []*string) *SumStorageMetricsByDateRequest {
	s.ProjectNames = v
	return s
}

func (s *SumStorageMetricsByDateRequest) SetStartDate(v int64) *SumStorageMetricsByDateRequest {
	s.StartDate = &v
	return s
}

func (s *SumStorageMetricsByDateRequest) SetStatsType(v string) *SumStorageMetricsByDateRequest {
	s.StatsType = &v
	return s
}

func (s *SumStorageMetricsByDateRequest) Validate() error {
	return dara.Validate(s)
}
