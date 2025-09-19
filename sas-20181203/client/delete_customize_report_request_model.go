// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomizeReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v int64) *DeleteCustomizeReportRequest
	GetReportId() *int64
}

type DeleteCustomizeReportRequest struct {
	// The ID of the report.
	//
	// >  You can call the [DescribeCustomizeReportList](~~DescribeCustomizeReportList~~) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s DeleteCustomizeReportRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomizeReportRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomizeReportRequest) GetReportId() *int64 {
	return s.ReportId
}

func (s *DeleteCustomizeReportRequest) SetReportId(v int64) *DeleteCustomizeReportRequest {
	s.ReportId = &v
	return s
}

func (s *DeleteCustomizeReportRequest) Validate() error {
	return dara.Validate(s)
}
