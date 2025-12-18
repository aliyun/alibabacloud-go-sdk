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
	// The request ID.
	//
	// example:
	//
	// ABC0FFF8-0B44-40C6-8BBF-3A185EFDD212
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the involved resources were evaluated. Valid values:
	//
	// 	- true: The involved resources were evaluated.
	//
	// 	- false: The involved resources were not evaluated
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
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
