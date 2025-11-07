// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNisInspectionReportStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInspectionReportId(v string) *DescribeNisInspectionReportStatusRequest
	GetInspectionReportId() *string
}

type DescribeNisInspectionReportStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// nir-2ca527b8de114ba4afb9
	InspectionReportId *string `json:"InspectionReportId,omitempty" xml:"InspectionReportId,omitempty"`
}

func (s DescribeNisInspectionReportStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionReportStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionReportStatusRequest) GetInspectionReportId() *string {
	return s.InspectionReportId
}

func (s *DescribeNisInspectionReportStatusRequest) SetInspectionReportId(v string) *DescribeNisInspectionReportStatusRequest {
	s.InspectionReportId = &v
	return s
}

func (s *DescribeNisInspectionReportStatusRequest) Validate() error {
	return dara.Validate(s)
}
