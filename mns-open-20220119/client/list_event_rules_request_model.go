// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListEventRulesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListEventRulesRequest
	GetNextToken() *string
	SetPageNum(v int64) *ListEventRulesRequest
	GetPageNum() *int64
	SetPageSize(v int64) *ListEventRulesRequest
	GetPageSize() *int64
	SetProductName(v string) *ListEventRulesRequest
	GetProductName() *string
	SetResourceName(v string) *ListEventRulesRequest
	GetResourceName() *string
	SetRuleName(v string) *ListEventRulesRequest
	GetRuleName() *string
	SetSubscription(v *ListEventRulesRequestSubscription) *ListEventRulesRequest
	GetSubscription() *ListEventRulesRequestSubscription
	SetTopicName(v string) *ListEventRulesRequest
	GetTopicName() *string
}

type ListEventRulesRequest struct {
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
	Subscription *ListEventRulesRequestSubscription `json:"Subscription,omitempty" xml:"Subscription,omitempty" type:"Struct"`
	// The name of the topic.
	//
	// example:
	//
	// test-topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s ListEventRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEventRulesRequest) GoString() string {
	return s.String()
}

func (s *ListEventRulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEventRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEventRulesRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListEventRulesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListEventRulesRequest) GetProductName() *string {
	return s.ProductName
}

func (s *ListEventRulesRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListEventRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListEventRulesRequest) GetSubscription() *ListEventRulesRequestSubscription {
	return s.Subscription
}

func (s *ListEventRulesRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *ListEventRulesRequest) SetMaxResults(v int32) *ListEventRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEventRulesRequest) SetNextToken(v string) *ListEventRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListEventRulesRequest) SetPageNum(v int64) *ListEventRulesRequest {
	s.PageNum = &v
	return s
}

func (s *ListEventRulesRequest) SetPageSize(v int64) *ListEventRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListEventRulesRequest) SetProductName(v string) *ListEventRulesRequest {
	s.ProductName = &v
	return s
}

func (s *ListEventRulesRequest) SetResourceName(v string) *ListEventRulesRequest {
	s.ResourceName = &v
	return s
}

func (s *ListEventRulesRequest) SetRuleName(v string) *ListEventRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListEventRulesRequest) SetSubscription(v *ListEventRulesRequestSubscription) *ListEventRulesRequest {
	s.Subscription = v
	return s
}

func (s *ListEventRulesRequest) SetTopicName(v string) *ListEventRulesRequest {
	s.TopicName = &v
	return s
}

func (s *ListEventRulesRequest) Validate() error {
	if s.Subscription != nil {
		if err := s.Subscription.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEventRulesRequestSubscription struct {
	// The endpoint type. Valid values:
	//
	// - **topic**: The topic type. A topic supports delivery to multiple subscription endpoints. You can add or remove subscription endpoints later.
	//
	// - **queue**: The queue type. Messages are delivered directly to a queue. This reduces the link path, but you cannot add new subscription endpoints later.
	//
	// example:
	//
	// topic
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The actual value of the receiving endpoint. When EndpointType is set to topic, this is the topic name. When EndpointType is set to queue, this is the queue name.
	//
	// example:
	//
	// test-topic
	EndpointValue *string `json:"EndpointValue,omitempty" xml:"EndpointValue,omitempty"`
}

func (s ListEventRulesRequestSubscription) String() string {
	return dara.Prettify(s)
}

func (s ListEventRulesRequestSubscription) GoString() string {
	return s.String()
}

func (s *ListEventRulesRequestSubscription) GetEndpointType() *string {
	return s.EndpointType
}

func (s *ListEventRulesRequestSubscription) GetEndpointValue() *string {
	return s.EndpointValue
}

func (s *ListEventRulesRequestSubscription) SetEndpointType(v string) *ListEventRulesRequestSubscription {
	s.EndpointType = &v
	return s
}

func (s *ListEventRulesRequestSubscription) SetEndpointValue(v string) *ListEventRulesRequestSubscription {
	s.EndpointValue = &v
	return s
}

func (s *ListEventRulesRequestSubscription) Validate() error {
	return dara.Validate(s)
}
