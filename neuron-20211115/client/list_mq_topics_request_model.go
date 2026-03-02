// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMqTopicsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderBy(v string) *ListMqTopicsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListMqTopicsRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListMqTopicsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMqTopicsRequest
	GetPageSize() *int32
	SetTopicName(v string) *ListMqTopicsRequest
	GetTopicName() *string
}

type ListMqTopicsRequest struct {
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TopicName      *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s ListMqTopicsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMqTopicsRequest) GoString() string {
	return s.String()
}

func (s *ListMqTopicsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMqTopicsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListMqTopicsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMqTopicsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMqTopicsRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *ListMqTopicsRequest) SetOrderBy(v string) *ListMqTopicsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMqTopicsRequest) SetOrderDirection(v string) *ListMqTopicsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListMqTopicsRequest) SetPageNumber(v int32) *ListMqTopicsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMqTopicsRequest) SetPageSize(v int32) *ListMqTopicsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMqTopicsRequest) SetTopicName(v string) *ListMqTopicsRequest {
	s.TopicName = &v
	return s
}

func (s *ListMqTopicsRequest) Validate() error {
	return dara.Validate(s)
}
