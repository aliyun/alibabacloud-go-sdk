// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateRecommendManagedRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecommendedManagedRules(v *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules) *ListAggregateRecommendManagedRulesResponseBody
	GetRecommendedManagedRules() *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules
	SetRequestId(v string) *ListAggregateRecommendManagedRulesResponseBody
	GetRequestId() *string
}

type ListAggregateRecommendManagedRulesResponseBody struct {
	// The list of rules.
	RecommendedManagedRules *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules `json:"RecommendedManagedRules,omitempty" xml:"RecommendedManagedRules,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6CE4ABA1-9A57-41A9-8EA9-E8B17D46****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAggregateRecommendManagedRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateRecommendManagedRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggregateRecommendManagedRulesResponseBody) GetRecommendedManagedRules() *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules {
	return s.RecommendedManagedRules
}

func (s *ListAggregateRecommendManagedRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAggregateRecommendManagedRulesResponseBody) SetRecommendedManagedRules(v *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules) *ListAggregateRecommendManagedRulesResponseBody {
	s.RecommendedManagedRules = v
	return s
}

func (s *ListAggregateRecommendManagedRulesResponseBody) SetRequestId(v string) *ListAggregateRecommendManagedRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAggregateRecommendManagedRulesResponseBody) Validate() error {
	if s.RecommendedManagedRules != nil {
		if err := s.RecommendedManagedRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules struct {
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 200
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// If the return results are truncated, you can use NextToken to initiate another request to retrieve the remaining results.
	//
	// example:
	//
	// zXZXbg4Mra0kOrhpwl21****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of managed rules.
	RecommendedManagedRuleList []*ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList `json:"RecommendedManagedRuleList,omitempty" xml:"RecommendedManagedRuleList,omitempty" type:"Repeated"`
	// The total number of rule templates.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules) GoString() string {
	return s.String()
}

func (s *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules) GetRecommendedManagedRuleList() []*ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList {
	return s.RecommendedManagedRuleList
}

func (s *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules) SetMaxResults(v int32) *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules {
	s.MaxResults = &v
	return s
}

func (s *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules) SetNextToken(v string) *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules {
	s.NextToken = &v
	return s
}

func (s *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules) SetRecommendedManagedRuleList(v []*ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules {
	s.RecommendedManagedRuleList = v
	return s
}

func (s *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules) SetTotalCount(v int64) *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules {
	s.TotalCount = &v
	return s
}

func (s *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRules) Validate() error {
	if s.RecommendedManagedRuleList != nil {
		for _, item := range s.RecommendedManagedRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList struct {
	// The name of the rule.
	//
	// example:
	//
	// ram-policy-in-use-check
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// If the permission policy is bound to one or more RAM user groups, RAM roles, or RAM users, the configuration is considered compliant.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The identifier of the rule.
	//
	// example:
	//
	// ram-policy-in-use-check
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::RAM::Policy
	ResourceTypeScope *string `json:"ResourceTypeScope,omitempty" xml:"ResourceTypeScope,omitempty"`
}

func (s ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) GoString() string {
	return s.String()
}

func (s *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) GetDescription() *string {
	return s.Description
}

func (s *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) GetResourceTypeScope() *string {
	return s.ResourceTypeScope
}

func (s *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) SetConfigRuleName(v string) *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList {
	s.ConfigRuleName = &v
	return s
}

func (s *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) SetDescription(v string) *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList {
	s.Description = &v
	return s
}

func (s *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) SetIdentifier(v string) *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList {
	s.Identifier = &v
	return s
}

func (s *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) SetResourceTypeScope(v string) *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList {
	s.ResourceTypeScope = &v
	return s
}

func (s *ListAggregateRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) Validate() error {
	return dara.Validate(s)
}
