// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateConfigRuleEvaluationStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *ListAggregateConfigRuleEvaluationStatisticsRequest
	GetAggregatorId() *string
}

type ListAggregateConfigRuleEvaluationStatisticsRequest struct {
	// This parameter is required.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
}

func (s ListAggregateConfigRuleEvaluationStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRuleEvaluationStatisticsRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRuleEvaluationStatisticsRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListAggregateConfigRuleEvaluationStatisticsRequest) SetAggregatorId(v string) *ListAggregateConfigRuleEvaluationStatisticsRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
