// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDataQualityRulesFromEvaluationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachDataQualityRulesFromEvaluationTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DetachDataQualityRulesFromEvaluationTaskResponseBody
	GetSuccess() *bool
}

type DetachDataQualityRulesFromEvaluationTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8abcb91f-d266-4073-b907-2ed670378ed1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call is successful. The values are as follows:
	//
	// - true: The call is successful.
	//
	// - false: the call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetachDataQualityRulesFromEvaluationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachDataQualityRulesFromEvaluationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DetachDataQualityRulesFromEvaluationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachDataQualityRulesFromEvaluationTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DetachDataQualityRulesFromEvaluationTaskResponseBody) SetRequestId(v string) *DetachDataQualityRulesFromEvaluationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachDataQualityRulesFromEvaluationTaskResponseBody) SetSuccess(v bool) *DetachDataQualityRulesFromEvaluationTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DetachDataQualityRulesFromEvaluationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
