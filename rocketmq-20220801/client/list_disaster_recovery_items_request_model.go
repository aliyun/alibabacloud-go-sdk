// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDisasterRecoveryItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *ListDisasterRecoveryItemsRequest
	GetFilter() *string
	SetPageNumber(v int32) *ListDisasterRecoveryItemsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDisasterRecoveryItemsRequest
	GetPageSize() *int32
	SetTopicName(v string) *ListDisasterRecoveryItemsRequest
	GetTopicName() *string
}

type ListDisasterRecoveryItemsRequest struct {
	// The filter condition. Topics are filtered by topic name.
	//
	// example:
	//
	// topic_test
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The topic name.
	//
	// example:
	//
	// topic-test920
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s ListDisasterRecoveryItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryItemsRequest) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryItemsRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListDisasterRecoveryItemsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDisasterRecoveryItemsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDisasterRecoveryItemsRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *ListDisasterRecoveryItemsRequest) SetFilter(v string) *ListDisasterRecoveryItemsRequest {
	s.Filter = &v
	return s
}

func (s *ListDisasterRecoveryItemsRequest) SetPageNumber(v int32) *ListDisasterRecoveryItemsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDisasterRecoveryItemsRequest) SetPageSize(v int32) *ListDisasterRecoveryItemsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDisasterRecoveryItemsRequest) SetTopicName(v string) *ListDisasterRecoveryItemsRequest {
	s.TopicName = &v
	return s
}

func (s *ListDisasterRecoveryItemsRequest) Validate() error {
	return dara.Validate(s)
}
