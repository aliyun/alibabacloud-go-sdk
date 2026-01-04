// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubscriptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v *ListSubscriptionsResponseBodyList) *ListSubscriptionsResponseBody
	GetList() *ListSubscriptionsResponseBodyList
	SetMaxResults(v int32) *ListSubscriptionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListSubscriptionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSubscriptionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSubscriptionsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListSubscriptionsResponseBody
	GetTotalCount() *int32
}

type ListSubscriptionsResponseBody struct {
	List *ListSubscriptionsResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 9892074a2a89600ae4b0d5a34fb99a3f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// A20A7093-8FE0-058C-BE0C-3C8057D5F1A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSubscriptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubscriptionsResponseBody) GetList() *ListSubscriptionsResponseBodyList {
	return s.List
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

func (s *ListSubscriptionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSubscriptionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSubscriptionsResponseBody) SetList(v *ListSubscriptionsResponseBodyList) *ListSubscriptionsResponseBody {
	s.List = v
	return s
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

func (s *ListSubscriptionsResponseBody) SetSuccess(v bool) *ListSubscriptionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSubscriptionsResponseBody) SetTotalCount(v int32) *ListSubscriptionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSubscriptionsResponseBody) Validate() error {
	if s.List != nil {
		if err := s.List.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSubscriptionsResponseBodyList struct {
	Subscription []*ListSubscriptionsResponseBodyListSubscription `json:"Subscription,omitempty" xml:"Subscription,omitempty" type:"Repeated"`
}

func (s ListSubscriptionsResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionsResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListSubscriptionsResponseBodyList) GetSubscription() []*ListSubscriptionsResponseBodyListSubscription {
	return s.Subscription
}

func (s *ListSubscriptionsResponseBodyList) SetSubscription(v []*ListSubscriptionsResponseBodyListSubscription) *ListSubscriptionsResponseBodyList {
	s.Subscription = v
	return s
}

func (s *ListSubscriptionsResponseBodyList) Validate() error {
	if s.Subscription != nil {
		for _, item := range s.Subscription {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSubscriptionsResponseBodyListSubscription struct {
	// example:
	//
	// test_application_name
	Application *string `json:"Application,omitempty" xml:"Application,omitempty"`
	// example:
	//
	// test_comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 1708171905000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1048133943212399
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 1
	State *int32 `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 1745824636429WZ2EE
	SubscriptionId *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
	// example:
	//
	// test_topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// example:
	//
	// USER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1708171905000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListSubscriptionsResponseBodyListSubscription) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionsResponseBodyListSubscription) GoString() string {
	return s.String()
}

func (s *ListSubscriptionsResponseBodyListSubscription) GetApplication() *string {
	return s.Application
}

func (s *ListSubscriptionsResponseBodyListSubscription) GetComment() *string {
	return s.Comment
}

func (s *ListSubscriptionsResponseBodyListSubscription) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListSubscriptionsResponseBodyListSubscription) GetCreator() *string {
	return s.Creator
}

func (s *ListSubscriptionsResponseBodyListSubscription) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListSubscriptionsResponseBodyListSubscription) GetState() *int32 {
	return s.State
}

func (s *ListSubscriptionsResponseBodyListSubscription) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *ListSubscriptionsResponseBodyListSubscription) GetTopicName() *string {
	return s.TopicName
}

func (s *ListSubscriptionsResponseBodyListSubscription) GetType() *string {
	return s.Type
}

func (s *ListSubscriptionsResponseBodyListSubscription) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListSubscriptionsResponseBodyListSubscription) SetApplication(v string) *ListSubscriptionsResponseBodyListSubscription {
	s.Application = &v
	return s
}

func (s *ListSubscriptionsResponseBodyListSubscription) SetComment(v string) *ListSubscriptionsResponseBodyListSubscription {
	s.Comment = &v
	return s
}

func (s *ListSubscriptionsResponseBodyListSubscription) SetCreateTime(v int64) *ListSubscriptionsResponseBodyListSubscription {
	s.CreateTime = &v
	return s
}

func (s *ListSubscriptionsResponseBodyListSubscription) SetCreator(v string) *ListSubscriptionsResponseBodyListSubscription {
	s.Creator = &v
	return s
}

func (s *ListSubscriptionsResponseBodyListSubscription) SetProjectName(v string) *ListSubscriptionsResponseBodyListSubscription {
	s.ProjectName = &v
	return s
}

func (s *ListSubscriptionsResponseBodyListSubscription) SetState(v int32) *ListSubscriptionsResponseBodyListSubscription {
	s.State = &v
	return s
}

func (s *ListSubscriptionsResponseBodyListSubscription) SetSubscriptionId(v string) *ListSubscriptionsResponseBodyListSubscription {
	s.SubscriptionId = &v
	return s
}

func (s *ListSubscriptionsResponseBodyListSubscription) SetTopicName(v string) *ListSubscriptionsResponseBodyListSubscription {
	s.TopicName = &v
	return s
}

func (s *ListSubscriptionsResponseBodyListSubscription) SetType(v string) *ListSubscriptionsResponseBodyListSubscription {
	s.Type = &v
	return s
}

func (s *ListSubscriptionsResponseBodyListSubscription) SetUpdateTime(v int64) *ListSubscriptionsResponseBodyListSubscription {
	s.UpdateTime = &v
	return s
}

func (s *ListSubscriptionsResponseBodyListSubscription) Validate() error {
	return dara.Validate(s)
}
