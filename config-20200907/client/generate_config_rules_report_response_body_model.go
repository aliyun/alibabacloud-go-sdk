// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateConfigRulesReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *GenerateConfigRulesReportResponseBody
	GetReportId() *string
	SetRequestId(v string) *GenerateConfigRulesReportResponseBody
	GetRequestId() *string
}

type GenerateConfigRulesReportResponseBody struct {
	// The ID of the compliance evaluation report.
	//
	// example:
	//
	// crp-88176457e0d900c9****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6EC7AED1-172F-42AE-9C12-295BC2ADB751
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateConfigRulesReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateConfigRulesReportResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateConfigRulesReportResponseBody) GetReportId() *string {
	return s.ReportId
}

func (s *GenerateConfigRulesReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateConfigRulesReportResponseBody) SetReportId(v string) *GenerateConfigRulesReportResponseBody {
	s.ReportId = &v
	return s
}

func (s *GenerateConfigRulesReportResponseBody) SetRequestId(v string) *GenerateConfigRulesReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateConfigRulesReportResponseBody) Validate() error {
	return dara.Validate(s)
}
