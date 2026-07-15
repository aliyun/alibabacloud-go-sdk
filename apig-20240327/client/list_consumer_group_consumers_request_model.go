// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerGroupConsumersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNameLike(v string) *ListConsumerGroupConsumersRequest
	GetNameLike() *string
	SetPageNumber(v int32) *ListConsumerGroupConsumersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListConsumerGroupConsumersRequest
	GetPageSize() *int32
}

type ListConsumerGroupConsumersRequest struct {
	// example:
	//
	// consumer
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListConsumerGroupConsumersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupConsumersRequest) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupConsumersRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListConsumerGroupConsumersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConsumerGroupConsumersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConsumerGroupConsumersRequest) SetNameLike(v string) *ListConsumerGroupConsumersRequest {
	s.NameLike = &v
	return s
}

func (s *ListConsumerGroupConsumersRequest) SetPageNumber(v int32) *ListConsumerGroupConsumersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConsumerGroupConsumersRequest) SetPageSize(v int32) *ListConsumerGroupConsumersRequest {
	s.PageSize = &v
	return s
}

func (s *ListConsumerGroupConsumersRequest) Validate() error {
	return dara.Validate(s)
}
