// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableMetricRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableMetricRulesResponseBody
	GetCode() *string
	SetMessage(v string) *DisableMetricRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableMetricRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableMetricRulesResponseBody
	GetSuccess() *bool
}

type DisableMetricRulesResponseBody struct {
	// The responses code.
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
	// RuleId is mandatory for this action.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FF38D33A-67C1-40EB-AB65-FAEE51EDB644
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

func (s DisableMetricRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableMetricRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DisableMetricRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableMetricRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableMetricRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableMetricRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableMetricRulesResponseBody) SetCode(v string) *DisableMetricRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DisableMetricRulesResponseBody) SetMessage(v string) *DisableMetricRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DisableMetricRulesResponseBody) SetRequestId(v string) *DisableMetricRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableMetricRulesResponseBody) SetSuccess(v bool) *DisableMetricRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DisableMetricRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
