// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductSubscriptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerName(v string) *ListProductSubscriptionsRequest
	GetConsumerName() *string
	SetPageNumber(v int32) *ListProductSubscriptionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProductSubscriptionsRequest
	GetPageSize() *int32
	SetStatus(v string) *ListProductSubscriptionsRequest
	GetStatus() *string
}

type ListProductSubscriptionsRequest struct {
	// The keyword of the consumer name used to search the subscription records of the API product by consumer name.
	//
	// example:
	//
	// consumer-1
	ConsumerName *string `json:"consumerName,omitempty" xml:"consumerName,omitempty"`
	// The page number. Pages start from 1.
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
	// The subscription status used to filter subscriptions. Valid values: PENDING and APPROVED.
	//
	// example:
	//
	// APPROVED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListProductSubscriptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductSubscriptionsRequest) GoString() string {
	return s.String()
}

func (s *ListProductSubscriptionsRequest) GetConsumerName() *string {
	return s.ConsumerName
}

func (s *ListProductSubscriptionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProductSubscriptionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProductSubscriptionsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListProductSubscriptionsRequest) SetConsumerName(v string) *ListProductSubscriptionsRequest {
	s.ConsumerName = &v
	return s
}

func (s *ListProductSubscriptionsRequest) SetPageNumber(v int32) *ListProductSubscriptionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProductSubscriptionsRequest) SetPageSize(v int32) *ListProductSubscriptionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProductSubscriptionsRequest) SetStatus(v string) *ListProductSubscriptionsRequest {
	s.Status = &v
	return s
}

func (s *ListProductSubscriptionsRequest) Validate() error {
	return dara.Validate(s)
}
