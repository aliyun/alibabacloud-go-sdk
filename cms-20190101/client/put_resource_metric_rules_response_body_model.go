// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutResourceMetricRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PutResourceMetricRulesResponseBody
	GetCode() *string
	SetFailedListResult(v *PutResourceMetricRulesResponseBodyFailedListResult) *PutResourceMetricRulesResponseBody
	GetFailedListResult() *PutResourceMetricRulesResponseBodyFailedListResult
	SetMessage(v string) *PutResourceMetricRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutResourceMetricRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PutResourceMetricRulesResponseBody
	GetSuccess() *bool
}

type PutResourceMetricRulesResponseBody struct {
	// The response code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The alert rules that failed to be created for the resource.
	FailedListResult *PutResourceMetricRulesResponseBodyFailedListResult `json:"FailedListResult,omitempty" xml:"FailedListResult,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15D1440E-BF24-5A41-93E4-36864635179E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s PutResourceMetricRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRulesResponseBody) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *PutResourceMetricRulesResponseBody) GetFailedListResult() *PutResourceMetricRulesResponseBodyFailedListResult {
	return s.FailedListResult
}

func (s *PutResourceMetricRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutResourceMetricRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutResourceMetricRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutResourceMetricRulesResponseBody) SetCode(v string) *PutResourceMetricRulesResponseBody {
	s.Code = &v
	return s
}

func (s *PutResourceMetricRulesResponseBody) SetFailedListResult(v *PutResourceMetricRulesResponseBodyFailedListResult) *PutResourceMetricRulesResponseBody {
	s.FailedListResult = v
	return s
}

func (s *PutResourceMetricRulesResponseBody) SetMessage(v string) *PutResourceMetricRulesResponseBody {
	s.Message = &v
	return s
}

func (s *PutResourceMetricRulesResponseBody) SetRequestId(v string) *PutResourceMetricRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutResourceMetricRulesResponseBody) SetSuccess(v bool) *PutResourceMetricRulesResponseBody {
	s.Success = &v
	return s
}

func (s *PutResourceMetricRulesResponseBody) Validate() error {
	if s.FailedListResult != nil {
		if err := s.FailedListResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PutResourceMetricRulesResponseBodyFailedListResult struct {
	Target []*PutResourceMetricRulesResponseBodyFailedListResultTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Repeated"`
}

func (s PutResourceMetricRulesResponseBodyFailedListResult) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRulesResponseBodyFailedListResult) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesResponseBodyFailedListResult) GetTarget() []*PutResourceMetricRulesResponseBodyFailedListResultTarget {
	return s.Target
}

func (s *PutResourceMetricRulesResponseBodyFailedListResult) SetTarget(v []*PutResourceMetricRulesResponseBodyFailedListResultTarget) *PutResourceMetricRulesResponseBodyFailedListResult {
	s.Target = v
	return s
}

func (s *PutResourceMetricRulesResponseBodyFailedListResult) Validate() error {
	if s.Target != nil {
		for _, item := range s.Target {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutResourceMetricRulesResponseBodyFailedListResultTarget struct {
	// The alert rule that failed to be created.
	Result *PutResourceMetricRulesResponseBodyFailedListResultTargetResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The ID of the alert rule.
	//
	// example:
	//
	// a151cd6023eacee2f0978e03863cc1697c89508****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s PutResourceMetricRulesResponseBodyFailedListResultTarget) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRulesResponseBodyFailedListResultTarget) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesResponseBodyFailedListResultTarget) GetResult() *PutResourceMetricRulesResponseBodyFailedListResultTargetResult {
	return s.Result
}

func (s *PutResourceMetricRulesResponseBodyFailedListResultTarget) GetRuleId() *string {
	return s.RuleId
}

func (s *PutResourceMetricRulesResponseBodyFailedListResultTarget) SetResult(v *PutResourceMetricRulesResponseBodyFailedListResultTargetResult) *PutResourceMetricRulesResponseBodyFailedListResultTarget {
	s.Result = v
	return s
}

func (s *PutResourceMetricRulesResponseBodyFailedListResultTarget) SetRuleId(v string) *PutResourceMetricRulesResponseBodyFailedListResultTarget {
	s.RuleId = &v
	return s
}

func (s *PutResourceMetricRulesResponseBodyFailedListResultTarget) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PutResourceMetricRulesResponseBodyFailedListResultTargetResult struct {
	// The response code.
	//
	// example:
	//
	// 404
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutResourceMetricRulesResponseBodyFailedListResultTargetResult) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRulesResponseBodyFailedListResultTargetResult) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesResponseBodyFailedListResultTargetResult) GetCode() *string {
	return s.Code
}

func (s *PutResourceMetricRulesResponseBodyFailedListResultTargetResult) GetMessage() *string {
	return s.Message
}

func (s *PutResourceMetricRulesResponseBodyFailedListResultTargetResult) GetSuccess() *bool {
	return s.Success
}

func (s *PutResourceMetricRulesResponseBodyFailedListResultTargetResult) SetCode(v string) *PutResourceMetricRulesResponseBodyFailedListResultTargetResult {
	s.Code = &v
	return s
}

func (s *PutResourceMetricRulesResponseBodyFailedListResultTargetResult) SetMessage(v string) *PutResourceMetricRulesResponseBodyFailedListResultTargetResult {
	s.Message = &v
	return s
}

func (s *PutResourceMetricRulesResponseBodyFailedListResultTargetResult) SetSuccess(v bool) *PutResourceMetricRulesResponseBodyFailedListResultTargetResult {
	s.Success = &v
	return s
}

func (s *PutResourceMetricRulesResponseBodyFailedListResultTargetResult) Validate() error {
	return dara.Validate(s)
}
