// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartConfigRuleEvaluationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartConfigRuleEvaluationResponseBody
	GetRequestId() *string
	SetResult(v bool) *StartConfigRuleEvaluationResponseBody
	GetResult() *bool
}

type StartConfigRuleEvaluationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D31EEAD7-BF1E-5927-977A-AFF9342A7273
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s StartConfigRuleEvaluationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartConfigRuleEvaluationResponseBody) GoString() string {
	return s.String()
}

func (s *StartConfigRuleEvaluationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartConfigRuleEvaluationResponseBody) GetResult() *bool {
	return s.Result
}

func (s *StartConfigRuleEvaluationResponseBody) SetRequestId(v string) *StartConfigRuleEvaluationResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartConfigRuleEvaluationResponseBody) SetResult(v bool) *StartConfigRuleEvaluationResponseBody {
	s.Result = &v
	return s
}

func (s *StartConfigRuleEvaluationResponseBody) Validate() error {
	return dara.Validate(s)
}
