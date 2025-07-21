// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayType(v string) *ListConsumersRequest
	GetGatewayType() *string
	SetNameLike(v string) *ListConsumersRequest
	GetNameLike() *string
	SetPageNumber(v int32) *ListConsumersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListConsumersRequest
	GetPageSize() *int32
}

type ListConsumersRequest struct {
	// example:
	//
	// AI
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// example:
	//
	// UI-test
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

func (s ListConsumersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConsumersRequest) GoString() string {
	return s.String()
}

func (s *ListConsumersRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListConsumersRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListConsumersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConsumersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConsumersRequest) SetGatewayType(v string) *ListConsumersRequest {
	s.GatewayType = &v
	return s
}

func (s *ListConsumersRequest) SetNameLike(v string) *ListConsumersRequest {
	s.NameLike = &v
	return s
}

func (s *ListConsumersRequest) SetPageNumber(v int32) *ListConsumersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConsumersRequest) SetPageSize(v int32) *ListConsumersRequest {
	s.PageSize = &v
	return s
}

func (s *ListConsumersRequest) Validate() error {
	return dara.Validate(s)
}
