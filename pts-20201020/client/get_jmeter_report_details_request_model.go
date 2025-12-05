// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJMeterReportDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *GetJMeterReportDetailsRequest
	GetReportId() *string
}

type GetJMeterReportDetailsRequest struct {
	// The report ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// KS2YE3J2
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s GetJMeterReportDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJMeterReportDetailsRequest) GoString() string {
	return s.String()
}

func (s *GetJMeterReportDetailsRequest) GetReportId() *string {
	return s.ReportId
}

func (s *GetJMeterReportDetailsRequest) SetReportId(v string) *GetJMeterReportDetailsRequest {
	s.ReportId = &v
	return s
}

func (s *GetJMeterReportDetailsRequest) Validate() error {
	return dara.Validate(s)
}
