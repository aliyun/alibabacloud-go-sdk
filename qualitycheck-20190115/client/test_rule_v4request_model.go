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
	// example:
	//
	// 1
	IsSchemeData *int32 `json:"IsSchemeData,omitempty" xml:"IsSchemeData,omitempty"`
	// This parameter is required.
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
