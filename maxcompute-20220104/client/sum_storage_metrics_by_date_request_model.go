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
	SetRegion(v string) *SumStorageMetricsByDateRequest
	GetRegion() *string
	SetStartDate(v int64) *SumStorageMetricsByDateRequest
	GetStartDate() *int64
	SetStatsType(v string) *SumStorageMetricsByDateRequest
	GetStatsType() *string
	SetUserId(v string) *SumStorageMetricsByDateRequest
	GetUserId() *string
}

type SumStorageMetricsByDateRequest struct {
	// example:
	//
	// 1718590596556
	EndDate      *int64    `json:"endDate,omitempty" xml:"endDate,omitempty"`
	ProjectNames []*string `json:"projectNames,omitempty" xml:"projectNames,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 1715393576201
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// example:
	//
	// PROJECT
	StatsType *string `json:"statsType,omitempty" xml:"statsType,omitempty"`
	// example:
	//
	// 12345
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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

func (s *SumStorageMetricsByDateRequest) GetRegion() *string {
	return s.Region
}

func (s *SumStorageMetricsByDateRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *SumStorageMetricsByDateRequest) GetStatsType() *string {
	return s.StatsType
}

func (s *SumStorageMetricsByDateRequest) GetUserId() *string {
	return s.UserId
}

func (s *SumStorageMetricsByDateRequest) SetEndDate(v int64) *SumStorageMetricsByDateRequest {
	s.EndDate = &v
	return s
}

func (s *SumStorageMetricsByDateRequest) SetProjectNames(v []*string) *SumStorageMetricsByDateRequest {
	s.ProjectNames = v
	return s
}

func (s *SumStorageMetricsByDateRequest) SetRegion(v string) *SumStorageMetricsByDateRequest {
	s.Region = &v
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

func (s *SumStorageMetricsByDateRequest) SetUserId(v string) *SumStorageMetricsByDateRequest {
	s.UserId = &v
	return s
}

func (s *SumStorageMetricsByDateRequest) Validate() error {
	return dara.Validate(s)
}
