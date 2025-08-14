// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeepCopyRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeepCopyRuleResponseBody
	GetRequestId() *string
	SetResultObject(v *DeepCopyRuleResponseBodyResultObject) *DeepCopyRuleResponseBody
	GetResultObject() *DeepCopyRuleResponseBodyResultObject
}

type DeepCopyRuleResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information
	ResultObject *DeepCopyRuleResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DeepCopyRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeepCopyRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeepCopyRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeepCopyRuleResponseBody) GetResultObject() *DeepCopyRuleResponseBodyResultObject {
	return s.ResultObject
}

func (s *DeepCopyRuleResponseBody) SetRequestId(v string) *DeepCopyRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeepCopyRuleResponseBody) SetResultObject(v *DeepCopyRuleResponseBodyResultObject) *DeepCopyRuleResponseBody {
	s.ResultObject = v
	return s
}

func (s *DeepCopyRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeepCopyRuleResponseBodyResultObject struct {
	// Whether to redirect to details
	//
	// example:
	//
	// true
	BatchCopyFlag *bool `json:"BatchCopyFlag,omitempty" xml:"BatchCopyFlag,omitempty"`
	// Primary key of the policy
	//
	// example:
	//
	// 2346
	ConsoleRuleId *int64 `json:"ConsoleRuleId,omitempty" xml:"ConsoleRuleId,omitempty"`
	// Policy ID
	//
	// example:
	//
	// 104556
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// Primary key of the policy version
	//
	// example:
	//
	// 1135
	RuleVersionId *int64 `json:"RuleVersionId,omitempty" xml:"RuleVersionId,omitempty"`
}

func (s DeepCopyRuleResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DeepCopyRuleResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DeepCopyRuleResponseBodyResultObject) GetBatchCopyFlag() *bool {
	return s.BatchCopyFlag
}

func (s *DeepCopyRuleResponseBodyResultObject) GetConsoleRuleId() *int64 {
	return s.ConsoleRuleId
}

func (s *DeepCopyRuleResponseBodyResultObject) GetRuleId() *string {
	return s.RuleId
}

func (s *DeepCopyRuleResponseBodyResultObject) GetRuleVersionId() *int64 {
	return s.RuleVersionId
}

func (s *DeepCopyRuleResponseBodyResultObject) SetBatchCopyFlag(v bool) *DeepCopyRuleResponseBodyResultObject {
	s.BatchCopyFlag = &v
	return s
}

func (s *DeepCopyRuleResponseBodyResultObject) SetConsoleRuleId(v int64) *DeepCopyRuleResponseBodyResultObject {
	s.ConsoleRuleId = &v
	return s
}

func (s *DeepCopyRuleResponseBodyResultObject) SetRuleId(v string) *DeepCopyRuleResponseBodyResultObject {
	s.RuleId = &v
	return s
}

func (s *DeepCopyRuleResponseBodyResultObject) SetRuleVersionId(v int64) *DeepCopyRuleResponseBodyResultObject {
	s.RuleVersionId = &v
	return s
}

func (s *DeepCopyRuleResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
