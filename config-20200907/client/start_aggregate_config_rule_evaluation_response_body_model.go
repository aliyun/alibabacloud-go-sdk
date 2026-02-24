// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAggregateConfigRuleEvaluationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartAggregateConfigRuleEvaluationResponseBody
	GetRequestId() *string
	SetResult(v bool) *StartAggregateConfigRuleEvaluationResponseBody
	GetResult() *bool
}

type StartAggregateConfigRuleEvaluationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s StartAggregateConfigRuleEvaluationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartAggregateConfigRuleEvaluationResponseBody) GoString() string {
	return s.String()
}

func (s *StartAggregateConfigRuleEvaluationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartAggregateConfigRuleEvaluationResponseBody) GetResult() *bool {
	return s.Result
}

func (s *StartAggregateConfigRuleEvaluationResponseBody) SetRequestId(v string) *StartAggregateConfigRuleEvaluationResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartAggregateConfigRuleEvaluationResponseBody) SetResult(v bool) *StartAggregateConfigRuleEvaluationResponseBody {
	s.Result = &v
	return s
}

func (s *StartAggregateConfigRuleEvaluationResponseBody) Validate() error {
	return dara.Validate(s)
}
