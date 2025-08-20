// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventCenterRuleNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListEventCenterRuleNameResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListEventCenterRuleNameResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *ListEventCenterRuleNameResponseBody
	GetRequestId() *string
	SetRuleNames(v []*ListEventCenterRuleNameResponseBodyRuleNames) *ListEventCenterRuleNameResponseBody
	GetRuleNames() []*ListEventCenterRuleNameResponseBodyRuleNames
}

type ListEventCenterRuleNameResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 031572FA-7D8F-4C05-B790-1071E0E05DE6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of names of event notification rules.
	//
	// example:
	//
	// [{\\"RuleName\\": \\"mlf\\", \\"RuleId\\": \\"crecr-73q93pgljm1pc2fp\\"}]
	RuleNames []*ListEventCenterRuleNameResponseBodyRuleNames `json:"RuleNames,omitempty" xml:"RuleNames,omitempty" type:"Repeated"`
}

func (s ListEventCenterRuleNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventCenterRuleNameResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventCenterRuleNameResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListEventCenterRuleNameResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListEventCenterRuleNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEventCenterRuleNameResponseBody) GetRuleNames() []*ListEventCenterRuleNameResponseBodyRuleNames {
	return s.RuleNames
}

func (s *ListEventCenterRuleNameResponseBody) SetCode(v string) *ListEventCenterRuleNameResponseBody {
	s.Code = &v
	return s
}

func (s *ListEventCenterRuleNameResponseBody) SetIsSuccess(v bool) *ListEventCenterRuleNameResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListEventCenterRuleNameResponseBody) SetRequestId(v string) *ListEventCenterRuleNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEventCenterRuleNameResponseBody) SetRuleNames(v []*ListEventCenterRuleNameResponseBodyRuleNames) *ListEventCenterRuleNameResponseBody {
	s.RuleNames = v
	return s
}

func (s *ListEventCenterRuleNameResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEventCenterRuleNameResponseBodyRuleNames struct {
	// The ID of the event notification rule.
	//
	// example:
	//
	// crecr-n6pbhgjxtl*****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the event notification rule.
	//
	// example:
	//
	// test-chain
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListEventCenterRuleNameResponseBodyRuleNames) String() string {
	return dara.Prettify(s)
}

func (s ListEventCenterRuleNameResponseBodyRuleNames) GoString() string {
	return s.String()
}

func (s *ListEventCenterRuleNameResponseBodyRuleNames) GetRuleId() *string {
	return s.RuleId
}

func (s *ListEventCenterRuleNameResponseBodyRuleNames) GetRuleName() *string {
	return s.RuleName
}

func (s *ListEventCenterRuleNameResponseBodyRuleNames) SetRuleId(v string) *ListEventCenterRuleNameResponseBodyRuleNames {
	s.RuleId = &v
	return s
}

func (s *ListEventCenterRuleNameResponseBodyRuleNames) SetRuleName(v string) *ListEventCenterRuleNameResponseBodyRuleNames {
	s.RuleName = &v
	return s
}

func (s *ListEventCenterRuleNameResponseBodyRuleNames) Validate() error {
	return dara.Validate(s)
}
