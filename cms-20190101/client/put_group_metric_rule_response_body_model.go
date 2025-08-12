// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutGroupMetricRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PutGroupMetricRuleResponseBody
	GetCode() *string
	SetMessage(v string) *PutGroupMetricRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutGroupMetricRuleResponseBody
	GetRequestId() *string
	SetResult(v *PutGroupMetricRuleResponseBodyResult) *PutGroupMetricRuleResponseBody
	GetResult() *PutGroupMetricRuleResponseBodyResult
	SetSuccess(v bool) *PutGroupMetricRuleResponseBody
	GetSuccess() *bool
}

type PutGroupMetricRuleResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 461CF2CD-2FC3-4B26-8645-7BD27E7D0F1D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	Result *PutGroupMetricRuleResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutGroupMetricRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutGroupMetricRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PutGroupMetricRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *PutGroupMetricRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutGroupMetricRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutGroupMetricRuleResponseBody) GetResult() *PutGroupMetricRuleResponseBodyResult {
	return s.Result
}

func (s *PutGroupMetricRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutGroupMetricRuleResponseBody) SetCode(v string) *PutGroupMetricRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PutGroupMetricRuleResponseBody) SetMessage(v string) *PutGroupMetricRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PutGroupMetricRuleResponseBody) SetRequestId(v string) *PutGroupMetricRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutGroupMetricRuleResponseBody) SetResult(v *PutGroupMetricRuleResponseBodyResult) *PutGroupMetricRuleResponseBody {
	s.Result = v
	return s
}

func (s *PutGroupMetricRuleResponseBody) SetSuccess(v bool) *PutGroupMetricRuleResponseBody {
	s.Success = &v
	return s
}

func (s *PutGroupMetricRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type PutGroupMetricRuleResponseBodyResult struct {
	// The ID of the alert rule.
	//
	// example:
	//
	// 123456
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s PutGroupMetricRuleResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s PutGroupMetricRuleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PutGroupMetricRuleResponseBodyResult) GetRuleId() *string {
	return s.RuleId
}

func (s *PutGroupMetricRuleResponseBodyResult) SetRuleId(v string) *PutGroupMetricRuleResponseBodyResult {
	s.RuleId = &v
	return s
}

func (s *PutGroupMetricRuleResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
