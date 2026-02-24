// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecommendManagedRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecommendedManagedRules(v *ListRecommendManagedRulesResponseBodyRecommendedManagedRules) *ListRecommendManagedRulesResponseBody
	GetRecommendedManagedRules() *ListRecommendManagedRulesResponseBodyRecommendedManagedRules
	SetRequestId(v string) *ListRecommendManagedRulesResponseBody
	GetRequestId() *string
}

type ListRecommendManagedRulesResponseBody struct {
	// The list of rules.
	RecommendedManagedRules *ListRecommendManagedRulesResponseBodyRecommendedManagedRules `json:"RecommendedManagedRules,omitempty" xml:"RecommendedManagedRules,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DC300244-FCE3-5061-8214-C27ECB66****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRecommendManagedRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendManagedRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecommendManagedRulesResponseBody) GetRecommendedManagedRules() *ListRecommendManagedRulesResponseBodyRecommendedManagedRules {
	return s.RecommendedManagedRules
}

func (s *ListRecommendManagedRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecommendManagedRulesResponseBody) SetRecommendedManagedRules(v *ListRecommendManagedRulesResponseBodyRecommendedManagedRules) *ListRecommendManagedRulesResponseBody {
	s.RecommendedManagedRules = v
	return s
}

func (s *ListRecommendManagedRulesResponseBody) SetRequestId(v string) *ListRecommendManagedRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecommendManagedRulesResponseBody) Validate() error {
	if s.RecommendedManagedRules != nil {
		if err := s.RecommendedManagedRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRecommendManagedRulesResponseBodyRecommendedManagedRules struct {
	// The maximum number of entries returned for the request.
	//
	// example:
	//
	// 200
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to initiate the next query.
	//
	// > If this parameter is left empty, no more results are returned.
	//
	// example:
	//
	// zXZXbg4Mra0kOrhpwl21****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of managed rules.
	RecommendedManagedRuleList []*ListRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList `json:"RecommendedManagedRuleList,omitempty" xml:"RecommendedManagedRuleList,omitempty" type:"Repeated"`
	// The total number of rules.
	//
	// example:
	//
	// 39
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRecommendManagedRulesResponseBodyRecommendedManagedRules) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendManagedRulesResponseBodyRecommendedManagedRules) GoString() string {
	return s.String()
}

func (s *ListRecommendManagedRulesResponseBodyRecommendedManagedRules) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRecommendManagedRulesResponseBodyRecommendedManagedRules) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRecommendManagedRulesResponseBodyRecommendedManagedRules) GetRecommendedManagedRuleList() []*ListRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList {
	return s.RecommendedManagedRuleList
}

func (s *ListRecommendManagedRulesResponseBodyRecommendedManagedRules) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRecommendManagedRulesResponseBodyRecommendedManagedRules) SetMaxResults(v int32) *ListRecommendManagedRulesResponseBodyRecommendedManagedRules {
	s.MaxResults = &v
	return s
}

func (s *ListRecommendManagedRulesResponseBodyRecommendedManagedRules) SetNextToken(v string) *ListRecommendManagedRulesResponseBodyRecommendedManagedRules {
	s.NextToken = &v
	return s
}

func (s *ListRecommendManagedRulesResponseBodyRecommendedManagedRules) SetRecommendedManagedRuleList(v []*ListRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) *ListRecommendManagedRulesResponseBodyRecommendedManagedRules {
	s.RecommendedManagedRuleList = v
	return s
}

func (s *ListRecommendManagedRulesResponseBodyRecommendedManagedRules) SetTotalCount(v int64) *ListRecommendManagedRulesResponseBodyRecommendedManagedRules {
	s.TotalCount = &v
	return s
}

func (s *ListRecommendManagedRulesResponseBodyRecommendedManagedRules) Validate() error {
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

type ListRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList struct {
	// The rule name.
	//
	// example:
	//
	// oss-bucket-referer-limit
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The description of the managed rule.
	//
	// example:
	//
	// If the hotlink protection feature is enabled for the OSS bucket and the Referer is added to a specified whitelist, the configuration is considered compliant.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The rule identifier.
	//
	// example:
	//
	// oss-bucket-referer-limit
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::OSS::Bucket
	ResourceTypeScope *string `json:"ResourceTypeScope,omitempty" xml:"ResourceTypeScope,omitempty"`
}

func (s ListRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) GoString() string {
	return s.String()
}

func (s *ListRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *ListRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) GetDescription() *string {
	return s.Description
}

func (s *ListRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) GetResourceTypeScope() *string {
	return s.ResourceTypeScope
}

func (s *ListRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) SetConfigRuleName(v string) *ListRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList {
	s.ConfigRuleName = &v
	return s
}

func (s *ListRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) SetDescription(v string) *ListRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList {
	s.Description = &v
	return s
}

func (s *ListRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) SetIdentifier(v string) *ListRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList {
	s.Identifier = &v
	return s
}

func (s *ListRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) SetResourceTypeScope(v string) *ListRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList {
	s.ResourceTypeScope = &v
	return s
}

func (s *ListRecommendManagedRulesResponseBodyRecommendedManagedRulesRecommendedManagedRuleList) Validate() error {
	return dara.Validate(s)
}
