// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceReportAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeInstanceReportAttributeRequest
	GetRegionId() *string
	SetReportId(v string) *DescribeInstanceReportAttributeRequest
	GetReportId() *string
}

type DescribeInstanceReportAttributeRequest struct {
	// The ID of the region where the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the instance inspection report to query.
	//
	// example:
	//
	// report-001
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s DescribeInstanceReportAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceReportAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceReportAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceReportAttributeRequest) GetReportId() *string {
	return s.ReportId
}

func (s *DescribeInstanceReportAttributeRequest) SetRegionId(v string) *DescribeInstanceReportAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceReportAttributeRequest) SetReportId(v string) *DescribeInstanceReportAttributeRequest {
	s.ReportId = &v
	return s
}

func (s *DescribeInstanceReportAttributeRequest) Validate() error {
	return dara.Validate(s)
}
