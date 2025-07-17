// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnoseReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeDiagnoseReportsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDiagnoseReportsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDiagnoseReportsRequest
	GetRegionId() *string
	SetReportIds(v []*string) *DescribeDiagnoseReportsRequest
	GetReportIds() []*string
	SetScalingGroupId(v string) *DescribeDiagnoseReportsRequest
	GetScalingGroupId() *string
}

type DescribeDiagnoseReportsRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the diagnostic reports. You can specify at most 20 IDs.
	ReportIds []*string `json:"ReportIds,omitempty" xml:"ReportIds,omitempty" type:"Repeated"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-2vcis7yglxtm*****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeDiagnoseReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnoseReportsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDiagnoseReportsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDiagnoseReportsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiagnoseReportsRequest) GetReportIds() []*string {
	return s.ReportIds
}

func (s *DescribeDiagnoseReportsRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeDiagnoseReportsRequest) SetPageNumber(v int32) *DescribeDiagnoseReportsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiagnoseReportsRequest) SetPageSize(v int32) *DescribeDiagnoseReportsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDiagnoseReportsRequest) SetRegionId(v string) *DescribeDiagnoseReportsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnoseReportsRequest) SetReportIds(v []*string) *DescribeDiagnoseReportsRequest {
	s.ReportIds = v
	return s
}

func (s *DescribeDiagnoseReportsRequest) SetScalingGroupId(v string) *DescribeDiagnoseReportsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeDiagnoseReportsRequest) Validate() error {
	return dara.Validate(s)
}
