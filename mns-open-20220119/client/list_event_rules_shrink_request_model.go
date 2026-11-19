// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListEventRulesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListEventRulesShrinkRequest
	GetNextToken() *string
	SetPageNum(v int64) *ListEventRulesShrinkRequest
	GetPageNum() *int64
	SetPageSize(v int64) *ListEventRulesShrinkRequest
	GetPageSize() *int64
	SetProductName(v string) *ListEventRulesShrinkRequest
	GetProductName() *string
	SetResourceName(v string) *ListEventRulesShrinkRequest
	GetResourceName() *string
	SetRuleName(v string) *ListEventRulesShrinkRequest
	GetRuleName() *string
	SetSubscriptionShrink(v string) *ListEventRulesShrinkRequest
	GetSubscriptionShrink() *string
	SetTopicName(v string) *ListEventRulesShrinkRequest
	GetTopicName() *string
}

type ListEventRulesShrinkRequest struct {
	// **[Deprecated]*	- Use PageSize for paging.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// **[Deprecated]*	- Use PageNum for paging.
	//
	// example:
	//
	// cd7NlPlX4kgKCdsCWMiMR/+HnVzPLQ4/XLvjR64jZ7F9AQ+Mr3T59J6IVkuXeV3w
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number of the returned results. Valid values: 1 to 100000. If the specified value is less than 1, the system defaults to 1. If the specified value is greater than 100000, the system defaults to 100000.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Valid values: 10 to 50. If the specified value is less than 10, the system defaults to 10. If the specified value is greater than 50, the system defaults to 50.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the cloud service for event notifications.
	//
	// example:
	//
	// oss
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The resource name in the matching rules, used to filter rules. For example, in OSS, this is the OSS bucket name.
	//
	// example:
	//
	// test-bucket
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The filter condition for the rule name.
	//
	// example:
	//
	// rule-xsXDW
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The subscription endpoint.
	SubscriptionShrink *string `json:"Subscription,omitempty" xml:"Subscription,omitempty"`
	// The name of the topic.
	//
	// example:
	//
	// test-topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s ListEventRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEventRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListEventRulesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEventRulesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEventRulesShrinkRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListEventRulesShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListEventRulesShrinkRequest) GetProductName() *string {
	return s.ProductName
}

func (s *ListEventRulesShrinkRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListEventRulesShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListEventRulesShrinkRequest) GetSubscriptionShrink() *string {
	return s.SubscriptionShrink
}

func (s *ListEventRulesShrinkRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *ListEventRulesShrinkRequest) SetMaxResults(v int32) *ListEventRulesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEventRulesShrinkRequest) SetNextToken(v string) *ListEventRulesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListEventRulesShrinkRequest) SetPageNum(v int64) *ListEventRulesShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *ListEventRulesShrinkRequest) SetPageSize(v int64) *ListEventRulesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListEventRulesShrinkRequest) SetProductName(v string) *ListEventRulesShrinkRequest {
	s.ProductName = &v
	return s
}

func (s *ListEventRulesShrinkRequest) SetResourceName(v string) *ListEventRulesShrinkRequest {
	s.ResourceName = &v
	return s
}

func (s *ListEventRulesShrinkRequest) SetRuleName(v string) *ListEventRulesShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *ListEventRulesShrinkRequest) SetSubscriptionShrink(v string) *ListEventRulesShrinkRequest {
	s.SubscriptionShrink = &v
	return s
}

func (s *ListEventRulesShrinkRequest) SetTopicName(v string) *ListEventRulesShrinkRequest {
	s.TopicName = &v
	return s
}

func (s *ListEventRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
