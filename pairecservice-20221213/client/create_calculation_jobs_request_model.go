// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCalculationJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetABMetricIds(v string) *CreateCalculationJobsRequest
	GetABMetricIds() *string
	SetEndDate(v string) *CreateCalculationJobsRequest
	GetEndDate() *string
	SetInstanceId(v string) *CreateCalculationJobsRequest
	GetInstanceId() *string
	SetStartDate(v string) *CreateCalculationJobsRequest
	GetStartDate() *string
}

type CreateCalculationJobsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2,3,4
	ABMetricIds *string `json:"ABMetricIds,omitempty" xml:"ABMetricIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-03
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s CreateCalculationJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCalculationJobsRequest) GoString() string {
	return s.String()
}

func (s *CreateCalculationJobsRequest) GetABMetricIds() *string {
	return s.ABMetricIds
}

func (s *CreateCalculationJobsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *CreateCalculationJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCalculationJobsRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *CreateCalculationJobsRequest) SetABMetricIds(v string) *CreateCalculationJobsRequest {
	s.ABMetricIds = &v
	return s
}

func (s *CreateCalculationJobsRequest) SetEndDate(v string) *CreateCalculationJobsRequest {
	s.EndDate = &v
	return s
}

func (s *CreateCalculationJobsRequest) SetInstanceId(v string) *CreateCalculationJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCalculationJobsRequest) SetStartDate(v string) *CreateCalculationJobsRequest {
	s.StartDate = &v
	return s
}

func (s *CreateCalculationJobsRequest) Validate() error {
	return dara.Validate(s)
}
