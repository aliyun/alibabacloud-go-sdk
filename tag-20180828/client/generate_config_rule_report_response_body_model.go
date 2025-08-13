// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateConfigRuleReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *GenerateConfigRuleReportResponseBody
	GetReportId() *string
	SetRequestId(v string) *GenerateConfigRuleReportResponseBody
	GetRequestId() *string
}

type GenerateConfigRuleReportResponseBody struct {
	// The ID of the resource non-compliance report.
	//
	// example:
	//
	// crp-ao0786618088006c****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 37B2AC06-89D8-5D95-98DF-3E68C12BDE05
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateConfigRuleReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateConfigRuleReportResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateConfigRuleReportResponseBody) GetReportId() *string {
	return s.ReportId
}

func (s *GenerateConfigRuleReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateConfigRuleReportResponseBody) SetReportId(v string) *GenerateConfigRuleReportResponseBody {
	s.ReportId = &v
	return s
}

func (s *GenerateConfigRuleReportResponseBody) SetRequestId(v string) *GenerateConfigRuleReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateConfigRuleReportResponseBody) Validate() error {
	return dara.Validate(s)
}
