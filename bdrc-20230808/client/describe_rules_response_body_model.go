// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeRulesResponseBodyData) *DescribeRulesResponseBody
	GetData() *DescribeRulesResponseBodyData
	SetRequestId(v string) *DescribeRulesResponseBody
	GetRequestId() *string
}

type DescribeRulesResponseBody struct {
	// The returned data.
	Data *DescribeRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique ID of the request.
	//
	// example:
	//
	// 86DEBAC9-AB6A-59AB-9E5C-A540E579ECC9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponseBody) GetData() *DescribeRulesResponseBodyData {
	return s.Data
}

func (s *DescribeRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRulesResponseBody) SetData(v *DescribeRulesResponseBodyData) *DescribeRulesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeRulesResponseBody) SetRequestId(v string) *DescribeRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRulesResponseBodyData struct {
	// The list of returned records.
	Content []*DescribeRulesResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The maximum number of entries returned on the current page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. If this parameter is empty, all results have been returned.
	//
	// example:
	//
	// 0975951c75d7b41464c8d08ae17043ca
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries that meet the filter criteria. This parameter is optional and is not returned by default.
	//
	// example:
	//
	// 42
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponseBodyData) GetContent() []*DescribeRulesResponseBodyDataContent {
	return s.Content
}

func (s *DescribeRulesResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeRulesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRulesResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeRulesResponseBodyData) SetContent(v []*DescribeRulesResponseBodyDataContent) *DescribeRulesResponseBodyData {
	s.Content = v
	return s
}

func (s *DescribeRulesResponseBodyData) SetMaxResults(v int32) *DescribeRulesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeRulesResponseBodyData) SetNextToken(v string) *DescribeRulesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeRulesResponseBodyData) SetTotalCount(v int64) *DescribeRulesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeRulesResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRulesResponseBodyDataContent struct {
	// The number of resources for which the check failed.
	//
	// example:
	//
	// 0
	CheckFailedResourceCount *int64 `json:"CheckFailedResourceCount,omitempty" xml:"CheckFailedResourceCount,omitempty"`
	// The check status. Valid values: NOT_CHECKED (Not checked), PASSED (Passed), FAILED (Failed), CHECKING (Checking), and CHECK_FAILED (Check failed).
	//
	// example:
	//
	// PASSED
	CheckStatus *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// The UNIX timestamp that indicates when the check was performed.
	//
	// example:
	//
	// 1704157635
	CheckTime *int64 `json:"CheckTime,omitempty" xml:"CheckTime,omitempty"`
	// The product type to which the rule applies.
	//
	// example:
	//
	// ecs
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The resource type to which the rule applies.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The number of at-risk resources.
	//
	// example:
	//
	// 0
	RiskyResourceCount *int64 `json:"RiskyResourceCount,omitempty" xml:"RiskyResourceCount,omitempty"`
	// The unique ID of the rule.
	//
	// example:
	//
	// rule-bp11ggd8wr762
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The rule template.
	//
	// example:
	//
	// ecs-backup
	RuleTemplate *string `json:"RuleTemplate,omitempty" xml:"RuleTemplate,omitempty"`
	// The total number of resources that were checked.
	//
	// example:
	//
	// 1
	TotalResourceCount *int64 `json:"TotalResourceCount,omitempty" xml:"TotalResourceCount,omitempty"`
}

func (s DescribeRulesResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeRulesResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponseBodyDataContent) GetCheckFailedResourceCount() *int64 {
	return s.CheckFailedResourceCount
}

func (s *DescribeRulesResponseBodyDataContent) GetCheckStatus() *string {
	return s.CheckStatus
}

func (s *DescribeRulesResponseBodyDataContent) GetCheckTime() *int64 {
	return s.CheckTime
}

func (s *DescribeRulesResponseBodyDataContent) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeRulesResponseBodyDataContent) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeRulesResponseBodyDataContent) GetRiskyResourceCount() *int64 {
	return s.RiskyResourceCount
}

func (s *DescribeRulesResponseBodyDataContent) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeRulesResponseBodyDataContent) GetRuleTemplate() *string {
	return s.RuleTemplate
}

func (s *DescribeRulesResponseBodyDataContent) GetTotalResourceCount() *int64 {
	return s.TotalResourceCount
}

func (s *DescribeRulesResponseBodyDataContent) SetCheckFailedResourceCount(v int64) *DescribeRulesResponseBodyDataContent {
	s.CheckFailedResourceCount = &v
	return s
}

func (s *DescribeRulesResponseBodyDataContent) SetCheckStatus(v string) *DescribeRulesResponseBodyDataContent {
	s.CheckStatus = &v
	return s
}

func (s *DescribeRulesResponseBodyDataContent) SetCheckTime(v int64) *DescribeRulesResponseBodyDataContent {
	s.CheckTime = &v
	return s
}

func (s *DescribeRulesResponseBodyDataContent) SetProductType(v string) *DescribeRulesResponseBodyDataContent {
	s.ProductType = &v
	return s
}

func (s *DescribeRulesResponseBodyDataContent) SetResourceType(v string) *DescribeRulesResponseBodyDataContent {
	s.ResourceType = &v
	return s
}

func (s *DescribeRulesResponseBodyDataContent) SetRiskyResourceCount(v int64) *DescribeRulesResponseBodyDataContent {
	s.RiskyResourceCount = &v
	return s
}

func (s *DescribeRulesResponseBodyDataContent) SetRuleId(v string) *DescribeRulesResponseBodyDataContent {
	s.RuleId = &v
	return s
}

func (s *DescribeRulesResponseBodyDataContent) SetRuleTemplate(v string) *DescribeRulesResponseBodyDataContent {
	s.RuleTemplate = &v
	return s
}

func (s *DescribeRulesResponseBodyDataContent) SetTotalResourceCount(v int64) *DescribeRulesResponseBodyDataContent {
	s.TotalResourceCount = &v
	return s
}

func (s *DescribeRulesResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
