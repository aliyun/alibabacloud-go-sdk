// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayType(v string) *ListConsumerGroupsRequest
	GetGatewayType() *string
	SetNameLike(v string) *ListConsumerGroupsRequest
	GetNameLike() *string
	SetPageNumber(v int32) *ListConsumerGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListConsumerGroupsRequest
	GetPageSize() *int32
}

type ListConsumerGroupsRequest struct {
	// example:
	//
	// API
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// example:
	//
	// api
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

func (s ListConsumerGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupsRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListConsumerGroupsRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListConsumerGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConsumerGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConsumerGroupsRequest) SetGatewayType(v string) *ListConsumerGroupsRequest {
	s.GatewayType = &v
	return s
}

func (s *ListConsumerGroupsRequest) SetNameLike(v string) *ListConsumerGroupsRequest {
	s.NameLike = &v
	return s
}

func (s *ListConsumerGroupsRequest) SetPageNumber(v int32) *ListConsumerGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConsumerGroupsRequest) SetPageSize(v int32) *ListConsumerGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListConsumerGroupsRequest) Validate() error {
	return dara.Validate(s)
}
