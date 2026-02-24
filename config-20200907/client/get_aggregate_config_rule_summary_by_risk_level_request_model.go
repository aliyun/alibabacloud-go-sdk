// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateConfigRuleSummaryByRiskLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateConfigRuleSummaryByRiskLevelRequest
	GetAggregatorId() *string
}

type GetAggregateConfigRuleSummaryByRiskLevelRequest struct {
	// This parameter is required.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
}

func (s GetAggregateConfigRuleSummaryByRiskLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleSummaryByRiskLevelRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelRequest) SetAggregatorId(v string) *GetAggregateConfigRuleSummaryByRiskLevelRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelRequest) Validate() error {
	return dara.Validate(s)
}
