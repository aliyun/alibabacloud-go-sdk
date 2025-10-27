// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAntiBruteForceRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateAntiBruteForceRule(v *CreateAntiBruteForceRuleResponseBodyCreateAntiBruteForceRule) *CreateAntiBruteForceRuleResponseBody
	GetCreateAntiBruteForceRule() *CreateAntiBruteForceRuleResponseBodyCreateAntiBruteForceRule
	SetRequestId(v string) *CreateAntiBruteForceRuleResponseBody
	GetRequestId() *string
}

type CreateAntiBruteForceRuleResponseBody struct {
	// The information about the defense rule.
	CreateAntiBruteForceRule *CreateAntiBruteForceRuleResponseBodyCreateAntiBruteForceRule `json:"CreateAntiBruteForceRule,omitempty" xml:"CreateAntiBruteForceRule,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F35F45B0-5D6B-4238-BE02-A62D0760E840
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAntiBruteForceRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAntiBruteForceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAntiBruteForceRuleResponseBody) GetCreateAntiBruteForceRule() *CreateAntiBruteForceRuleResponseBodyCreateAntiBruteForceRule {
	return s.CreateAntiBruteForceRule
}

func (s *CreateAntiBruteForceRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAntiBruteForceRuleResponseBody) SetCreateAntiBruteForceRule(v *CreateAntiBruteForceRuleResponseBodyCreateAntiBruteForceRule) *CreateAntiBruteForceRuleResponseBody {
	s.CreateAntiBruteForceRule = v
	return s
}

func (s *CreateAntiBruteForceRuleResponseBody) SetRequestId(v string) *CreateAntiBruteForceRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAntiBruteForceRuleResponseBody) Validate() error {
	if s.CreateAntiBruteForceRule != nil {
		if err := s.CreateAntiBruteForceRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAntiBruteForceRuleResponseBodyCreateAntiBruteForceRule struct {
	// The ID of the defense rule.
	//
	// example:
	//
	// 65778
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateAntiBruteForceRuleResponseBodyCreateAntiBruteForceRule) String() string {
	return dara.Prettify(s)
}

func (s CreateAntiBruteForceRuleResponseBodyCreateAntiBruteForceRule) GoString() string {
	return s.String()
}

func (s *CreateAntiBruteForceRuleResponseBodyCreateAntiBruteForceRule) GetRuleId() *int64 {
	return s.RuleId
}

func (s *CreateAntiBruteForceRuleResponseBodyCreateAntiBruteForceRule) SetRuleId(v int64) *CreateAntiBruteForceRuleResponseBodyCreateAntiBruteForceRule {
	s.RuleId = &v
	return s
}

func (s *CreateAntiBruteForceRuleResponseBodyCreateAntiBruteForceRule) Validate() error {
	return dara.Validate(s)
}
