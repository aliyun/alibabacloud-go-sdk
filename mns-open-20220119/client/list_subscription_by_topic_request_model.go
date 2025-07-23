// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubscriptionByTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointType(v string) *ListSubscriptionByTopicRequest
	GetEndpointType() *string
	SetEndpointValue(v string) *ListSubscriptionByTopicRequest
	GetEndpointValue() *string
	SetPageNum(v int64) *ListSubscriptionByTopicRequest
	GetPageNum() *int64
	SetPageSize(v int64) *ListSubscriptionByTopicRequest
	GetPageSize() *int64
	SetSubscriptionName(v string) *ListSubscriptionByTopicRequest
	GetSubscriptionName() *string
	SetTopicName(v string) *ListSubscriptionByTopicRequest
	GetTopicName() *string
}

type ListSubscriptionByTopicRequest struct {
	EndpointType  *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	EndpointValue *string `json:"EndpointValue,omitempty" xml:"EndpointValue,omitempty"`
	// The page number. Valid values: 1 to 100000000. If you set this parameter to a value smaller than 1, the value of this parameter is 1 by default. If you set this parameter to a value greater than 100000000, the value of this parameter is 100000000 by default.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Value values: 10 to 50. If you set this parameter to a value smaller than 10, the value of this parameter is 10 by default. If you set this parameter to a value greater than 50, the value of this parameter is 50 by default.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the subscription.
	//
	// example:
	//
	// demo-subscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The topic name.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s ListSubscriptionByTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionByTopicRequest) GoString() string {
	return s.String()
}

func (s *ListSubscriptionByTopicRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *ListSubscriptionByTopicRequest) GetEndpointValue() *string {
	return s.EndpointValue
}

func (s *ListSubscriptionByTopicRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListSubscriptionByTopicRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSubscriptionByTopicRequest) GetSubscriptionName() *string {
	return s.SubscriptionName
}

func (s *ListSubscriptionByTopicRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *ListSubscriptionByTopicRequest) SetEndpointType(v string) *ListSubscriptionByTopicRequest {
	s.EndpointType = &v
	return s
}

func (s *ListSubscriptionByTopicRequest) SetEndpointValue(v string) *ListSubscriptionByTopicRequest {
	s.EndpointValue = &v
	return s
}

func (s *ListSubscriptionByTopicRequest) SetPageNum(v int64) *ListSubscriptionByTopicRequest {
	s.PageNum = &v
	return s
}

func (s *ListSubscriptionByTopicRequest) SetPageSize(v int64) *ListSubscriptionByTopicRequest {
	s.PageSize = &v
	return s
}

func (s *ListSubscriptionByTopicRequest) SetSubscriptionName(v string) *ListSubscriptionByTopicRequest {
	s.SubscriptionName = &v
	return s
}

func (s *ListSubscriptionByTopicRequest) SetTopicName(v string) *ListSubscriptionByTopicRequest {
	s.TopicName = &v
	return s
}

func (s *ListSubscriptionByTopicRequest) Validate() error {
	return dara.Validate(s)
}
