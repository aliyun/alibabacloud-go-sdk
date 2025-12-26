// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRelatedDefenseRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeRelatedDefenseRulesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeRelatedDefenseRulesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeRelatedDefenseRulesResponseBody
	GetRequestId() *string
	SetRules(v []*DescribeRelatedDefenseRulesResponseBodyRules) *DescribeRelatedDefenseRulesResponseBody
	GetRules() []*DescribeRelatedDefenseRulesResponseBodyRules
	SetTotalCount(v int64) *DescribeRelatedDefenseRulesResponseBody
	GetTotalCount() *int64
}

type DescribeRelatedDefenseRulesResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAGBgV9tolsLfijC4wam2htS*****D/46H3X2wIS
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// F35F45B0-5D6B-4238-BE02-A62D****E840
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     []*DescribeRelatedDefenseRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRelatedDefenseRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRelatedDefenseRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRelatedDefenseRulesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeRelatedDefenseRulesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRelatedDefenseRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRelatedDefenseRulesResponseBody) GetRules() []*DescribeRelatedDefenseRulesResponseBodyRules {
	return s.Rules
}

func (s *DescribeRelatedDefenseRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeRelatedDefenseRulesResponseBody) SetMaxResults(v int32) *DescribeRelatedDefenseRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeRelatedDefenseRulesResponseBody) SetNextToken(v string) *DescribeRelatedDefenseRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeRelatedDefenseRulesResponseBody) SetRequestId(v string) *DescribeRelatedDefenseRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRelatedDefenseRulesResponseBody) SetRules(v []*DescribeRelatedDefenseRulesResponseBodyRules) *DescribeRelatedDefenseRulesResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeRelatedDefenseRulesResponseBody) SetTotalCount(v int64) *DescribeRelatedDefenseRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRelatedDefenseRulesResponseBody) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRelatedDefenseRulesResponseBodyRules struct {
	// example:
	//
	// custom_acl
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// example:
	//
	// template
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// example:
	//
	// 2456789
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// ruleTest
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// 81501
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeRelatedDefenseRulesResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeRelatedDefenseRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeRelatedDefenseRulesResponseBodyRules) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeRelatedDefenseRulesResponseBodyRules) GetDefenseType() *string {
	return s.DefenseType
}

func (s *DescribeRelatedDefenseRulesResponseBodyRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeRelatedDefenseRulesResponseBodyRules) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeRelatedDefenseRulesResponseBodyRules) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeRelatedDefenseRulesResponseBodyRules) SetDefenseScene(v string) *DescribeRelatedDefenseRulesResponseBodyRules {
	s.DefenseScene = &v
	return s
}

func (s *DescribeRelatedDefenseRulesResponseBodyRules) SetDefenseType(v string) *DescribeRelatedDefenseRulesResponseBodyRules {
	s.DefenseType = &v
	return s
}

func (s *DescribeRelatedDefenseRulesResponseBodyRules) SetRuleId(v int64) *DescribeRelatedDefenseRulesResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *DescribeRelatedDefenseRulesResponseBodyRules) SetRuleName(v string) *DescribeRelatedDefenseRulesResponseBodyRules {
	s.RuleName = &v
	return s
}

func (s *DescribeRelatedDefenseRulesResponseBodyRules) SetTemplateId(v int64) *DescribeRelatedDefenseRulesResponseBodyRules {
	s.TemplateId = &v
	return s
}

func (s *DescribeRelatedDefenseRulesResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
