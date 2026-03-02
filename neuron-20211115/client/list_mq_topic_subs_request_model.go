// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMqTopicSubsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderBy(v string) *ListMqTopicSubsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListMqTopicSubsRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListMqTopicSubsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMqTopicSubsRequest
	GetPageSize() *int32
	SetServiceName(v string) *ListMqTopicSubsRequest
	GetServiceName() *string
}

type ListMqTopicSubsRequest struct {
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ServiceName    *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s ListMqTopicSubsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMqTopicSubsRequest) GoString() string {
	return s.String()
}

func (s *ListMqTopicSubsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMqTopicSubsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListMqTopicSubsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMqTopicSubsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMqTopicSubsRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListMqTopicSubsRequest) SetOrderBy(v string) *ListMqTopicSubsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMqTopicSubsRequest) SetOrderDirection(v string) *ListMqTopicSubsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListMqTopicSubsRequest) SetPageNumber(v int32) *ListMqTopicSubsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMqTopicSubsRequest) SetPageSize(v int32) *ListMqTopicSubsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMqTopicSubsRequest) SetServiceName(v string) *ListMqTopicSubsRequest {
	s.ServiceName = &v
	return s
}

func (s *ListMqTopicSubsRequest) Validate() error {
	return dara.Validate(s)
}
