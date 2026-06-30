// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestRuleV4Request interface {
	dara.Model
	String() string
	GoString() string
	SetIsSchemeData(v int32) *TestRuleV4Request
	GetIsSchemeData() *int32
	SetTestJson(v string) *TestRuleV4Request
	GetTestJson() *string
}

type TestRuleV4Request struct {
	// Whether this is the new quality check version. Valid values: 0 (legacy version) and 1 (new version). Default value: 1.
	//
	// example:
	//
	// 1
	IsSchemeData *int32 `json:"IsSchemeData,omitempty" xml:"IsSchemeData,omitempty"`
	// JSON request parameters for rule testing. For details, see the supplemental description of request parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"ruleList":[9771],"dialogues":[{"begin":0,"end":760,"hourMinSec":"00:00","role":"客户","identity":"客户","words":"123"},{"begin":21004,"end":21494,"hourMinSec":"00:21","role":"客服","identity":"客服","words":"123"}]}
	TestJson *string `json:"TestJson,omitempty" xml:"TestJson,omitempty"`
}

func (s TestRuleV4Request) String() string {
	return dara.Prettify(s)
}

func (s TestRuleV4Request) GoString() string {
	return s.String()
}

func (s *TestRuleV4Request) GetIsSchemeData() *int32 {
	return s.IsSchemeData
}

func (s *TestRuleV4Request) GetTestJson() *string {
	return s.TestJson
}

func (s *TestRuleV4Request) SetIsSchemeData(v int32) *TestRuleV4Request {
	s.IsSchemeData = &v
	return s
}

func (s *TestRuleV4Request) SetTestJson(v string) *TestRuleV4Request {
	s.TestJson = &v
	return s
}

func (s *TestRuleV4Request) Validate() error {
	return dara.Validate(s)
}
