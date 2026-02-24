// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAggregateConfigRulesReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GenerateAggregateConfigRulesReportResponseBody
	GetAggregatorId() *string
	SetReportId(v string) *GenerateAggregateConfigRulesReportResponseBody
	GetReportId() *string
	SetRequestId(v string) *GenerateAggregateConfigRulesReportResponseBody
	GetRequestId() *string
}

type GenerateAggregateConfigRulesReportResponseBody struct {
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ReportId     *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateAggregateConfigRulesReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateAggregateConfigRulesReportResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateAggregateConfigRulesReportResponseBody) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GenerateAggregateConfigRulesReportResponseBody) GetReportId() *string {
	return s.ReportId
}

func (s *GenerateAggregateConfigRulesReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateAggregateConfigRulesReportResponseBody) SetAggregatorId(v string) *GenerateAggregateConfigRulesReportResponseBody {
	s.AggregatorId = &v
	return s
}

func (s *GenerateAggregateConfigRulesReportResponseBody) SetReportId(v string) *GenerateAggregateConfigRulesReportResponseBody {
	s.ReportId = &v
	return s
}

func (s *GenerateAggregateConfigRulesReportResponseBody) SetRequestId(v string) *GenerateAggregateConfigRulesReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateAggregateConfigRulesReportResponseBody) Validate() error {
	return dara.Validate(s)
}
