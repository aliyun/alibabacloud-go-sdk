// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigRulesReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *GetConfigRulesReportRequest
	GetReportId() *string
}

type GetConfigRulesReportRequest struct {
	// The ID of the compliance evaluation report.
	//
	// example:
	//
	// crp-88176457e0d900c9****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s GetConfigRulesReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRulesReportRequest) GoString() string {
	return s.String()
}

func (s *GetConfigRulesReportRequest) GetReportId() *string {
	return s.ReportId
}

func (s *GetConfigRulesReportRequest) SetReportId(v string) *GetConfigRulesReportRequest {
	s.ReportId = &v
	return s
}

func (s *GetConfigRulesReportRequest) Validate() error {
	return dara.Validate(s)
}
