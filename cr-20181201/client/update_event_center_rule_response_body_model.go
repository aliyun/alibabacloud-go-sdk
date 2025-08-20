// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventCenterRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateEventCenterRuleResponseBody
	GetCode() *int32
	SetRequestId(v string) *UpdateEventCenterRuleResponseBody
	GetRequestId() *string
	SetRuleId(v string) *UpdateEventCenterRuleResponseBody
	GetRuleId() *string
}

type UpdateEventCenterRuleResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 031572FA-7D8F-4C05-B790-1071E0E05DE6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the event notification rule.
	//
	// example:
	//
	// crecr-n6pbhgjxt*****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s UpdateEventCenterRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventCenterRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEventCenterRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateEventCenterRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEventCenterRuleResponseBody) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateEventCenterRuleResponseBody) SetCode(v int32) *UpdateEventCenterRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEventCenterRuleResponseBody) SetRequestId(v string) *UpdateEventCenterRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEventCenterRuleResponseBody) SetRuleId(v string) *UpdateEventCenterRuleResponseBody {
	s.RuleId = &v
	return s
}

func (s *UpdateEventCenterRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
