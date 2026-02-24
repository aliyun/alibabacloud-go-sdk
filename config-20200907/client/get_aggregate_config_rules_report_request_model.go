// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateConfigRulesReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateConfigRulesReportRequest
	GetAggregatorId() *string
	SetReportId(v string) *GetAggregateConfigRulesReportRequest
	GetReportId() *string
}

type GetAggregateConfigRulesReportRequest struct {
	// This parameter is required.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ReportId     *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s GetAggregateConfigRulesReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRulesReportRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRulesReportRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateConfigRulesReportRequest) GetReportId() *string {
	return s.ReportId
}

func (s *GetAggregateConfigRulesReportRequest) SetAggregatorId(v string) *GetAggregateConfigRulesReportRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateConfigRulesReportRequest) SetReportId(v string) *GetAggregateConfigRulesReportRequest {
	s.ReportId = &v
	return s
}

func (s *GetAggregateConfigRulesReportRequest) Validate() error {
	return dara.Validate(s)
}
