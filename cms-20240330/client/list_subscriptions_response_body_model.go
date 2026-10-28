// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubscriptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSubscriptionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListSubscriptionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSubscriptionsResponseBody
	GetRequestId() *string
	SetSubscriptionList(v []*SubscriptionForView) *ListSubscriptionsResponseBody
	GetSubscriptionList() []*SubscriptionForView
}

type ListSubscriptionsResponseBody struct {
	// The maximum number of records returned for the current request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// 2-3408-4595-9d80-****
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 123-123-234-345-123
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of subscriptions.
	SubscriptionList []*SubscriptionForView `json:"subscriptionList,omitempty" xml:"subscriptionList,omitempty" type:"Repeated"`
}

func (s ListSubscriptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubscriptionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSubscriptionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSubscriptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSubscriptionsResponseBody) GetSubscriptionList() []*SubscriptionForView {
	return s.SubscriptionList
}

func (s *ListSubscriptionsResponseBody) SetMaxResults(v int32) *ListSubscriptionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSubscriptionsResponseBody) SetNextToken(v string) *ListSubscriptionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSubscriptionsResponseBody) SetRequestId(v string) *ListSubscriptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubscriptionsResponseBody) SetSubscriptionList(v []*SubscriptionForView) *ListSubscriptionsResponseBody {
	s.SubscriptionList = v
	return s
}

func (s *ListSubscriptionsResponseBody) Validate() error {
	if s.SubscriptionList != nil {
		for _, item := range s.SubscriptionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
