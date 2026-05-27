// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumComputeMetricsByUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v int64) *SumComputeMetricsByUsageRequest
	GetEndDate() *int64
	SetProjectNames(v []*string) *SumComputeMetricsByUsageRequest
	GetProjectNames() []*string
	SetStartDate(v int64) *SumComputeMetricsByUsageRequest
	GetStartDate() *int64
	SetUsageType(v string) *SumComputeMetricsByUsageRequest
	GetUsageType() *string
}

type SumComputeMetricsByUsageRequest struct {
	// example:
	//
	// 1718590596556
	EndDate      *int64    `json:"endDate,omitempty" xml:"endDate,omitempty"`
	ProjectNames []*string `json:"projectNames,omitempty" xml:"projectNames,omitempty" type:"Repeated"`
	// example:
	//
	// 1715393576201
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// example:
	//
	// SCAN
	UsageType *string `json:"usageType,omitempty" xml:"usageType,omitempty"`
}

func (s SumComputeMetricsByUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s SumComputeMetricsByUsageRequest) GoString() string {
	return s.String()
}

func (s *SumComputeMetricsByUsageRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *SumComputeMetricsByUsageRequest) GetProjectNames() []*string {
	return s.ProjectNames
}

func (s *SumComputeMetricsByUsageRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *SumComputeMetricsByUsageRequest) GetUsageType() *string {
	return s.UsageType
}

func (s *SumComputeMetricsByUsageRequest) SetEndDate(v int64) *SumComputeMetricsByUsageRequest {
	s.EndDate = &v
	return s
}

func (s *SumComputeMetricsByUsageRequest) SetProjectNames(v []*string) *SumComputeMetricsByUsageRequest {
	s.ProjectNames = v
	return s
}

func (s *SumComputeMetricsByUsageRequest) SetStartDate(v int64) *SumComputeMetricsByUsageRequest {
	s.StartDate = &v
	return s
}

func (s *SumComputeMetricsByUsageRequest) SetUsageType(v string) *SumComputeMetricsByUsageRequest {
	s.UsageType = &v
	return s
}

func (s *SumComputeMetricsByUsageRequest) Validate() error {
	return dara.Validate(s)
}
