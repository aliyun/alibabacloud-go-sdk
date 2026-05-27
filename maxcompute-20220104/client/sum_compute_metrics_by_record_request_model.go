// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumComputeMetricsByRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v int64) *SumComputeMetricsByRecordRequest
	GetEndDate() *int64
	SetProjectNames(v []*string) *SumComputeMetricsByRecordRequest
	GetProjectNames() []*string
	SetStartDate(v int64) *SumComputeMetricsByRecordRequest
	GetStartDate() *int64
}

type SumComputeMetricsByRecordRequest struct {
	// example:
	//
	// 1718590596556
	EndDate      *int64    `json:"endDate,omitempty" xml:"endDate,omitempty"`
	ProjectNames []*string `json:"projectNames,omitempty" xml:"projectNames,omitempty" type:"Repeated"`
	// example:
	//
	// 1715393576201
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s SumComputeMetricsByRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s SumComputeMetricsByRecordRequest) GoString() string {
	return s.String()
}

func (s *SumComputeMetricsByRecordRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *SumComputeMetricsByRecordRequest) GetProjectNames() []*string {
	return s.ProjectNames
}

func (s *SumComputeMetricsByRecordRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *SumComputeMetricsByRecordRequest) SetEndDate(v int64) *SumComputeMetricsByRecordRequest {
	s.EndDate = &v
	return s
}

func (s *SumComputeMetricsByRecordRequest) SetProjectNames(v []*string) *SumComputeMetricsByRecordRequest {
	s.ProjectNames = v
	return s
}

func (s *SumComputeMetricsByRecordRequest) SetStartDate(v int64) *SumComputeMetricsByRecordRequest {
	s.StartDate = &v
	return s
}

func (s *SumComputeMetricsByRecordRequest) Validate() error {
	return dara.Validate(s)
}
