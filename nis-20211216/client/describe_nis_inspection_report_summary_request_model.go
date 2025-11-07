// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNisInspectionReportSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInspectionReportId(v string) *DescribeNisInspectionReportSummaryRequest
	GetInspectionReportId() *string
}

type DescribeNisInspectionReportSummaryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// nir-38abb318b27b49cc9a01
	InspectionReportId *string `json:"InspectionReportId,omitempty" xml:"InspectionReportId,omitempty"`
}

func (s DescribeNisInspectionReportSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionReportSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionReportSummaryRequest) GetInspectionReportId() *string {
	return s.InspectionReportId
}

func (s *DescribeNisInspectionReportSummaryRequest) SetInspectionReportId(v string) *DescribeNisInspectionReportSummaryRequest {
	s.InspectionReportId = &v
	return s
}

func (s *DescribeNisInspectionReportSummaryRequest) Validate() error {
	return dara.Validate(s)
}
