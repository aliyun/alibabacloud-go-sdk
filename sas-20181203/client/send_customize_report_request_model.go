// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendCustomizeReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v int64) *SendCustomizeReportRequest
	GetReportId() *int64
}

type SendCustomizeReportRequest struct {
	// The ID of the security report.
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

func (s SendCustomizeReportRequest) String() string {
	return dara.Prettify(s)
}

func (s SendCustomizeReportRequest) GoString() string {
	return s.String()
}

func (s *SendCustomizeReportRequest) GetReportId() *int64 {
	return s.ReportId
}

func (s *SendCustomizeReportRequest) SetReportId(v int64) *SendCustomizeReportRequest {
	s.ReportId = &v
	return s
}

func (s *SendCustomizeReportRequest) Validate() error {
	return dara.Validate(s)
}
