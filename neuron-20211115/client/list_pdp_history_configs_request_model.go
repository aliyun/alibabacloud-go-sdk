// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpHistoryConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListPdpHistoryConfigsRequest
	GetConfigId() *int64
	SetOrderBy(v string) *ListPdpHistoryConfigsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListPdpHistoryConfigsRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListPdpHistoryConfigsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPdpHistoryConfigsRequest
	GetPageSize() *int32
	SetServiceGroupId(v int64) *ListPdpHistoryConfigsRequest
	GetServiceGroupId() *int64
	SetType(v string) *ListPdpHistoryConfigsRequest
	GetType() *string
}

type ListPdpHistoryConfigsRequest struct {
	ConfigId       *int64  `json:"configId,omitempty" xml:"configId,omitempty"`
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListPdpHistoryConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPdpHistoryConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListPdpHistoryConfigsRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListPdpHistoryConfigsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListPdpHistoryConfigsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListPdpHistoryConfigsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPdpHistoryConfigsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPdpHistoryConfigsRequest) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *ListPdpHistoryConfigsRequest) GetType() *string {
	return s.Type
}

func (s *ListPdpHistoryConfigsRequest) SetConfigId(v int64) *ListPdpHistoryConfigsRequest {
	s.ConfigId = &v
	return s
}

func (s *ListPdpHistoryConfigsRequest) SetOrderBy(v string) *ListPdpHistoryConfigsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPdpHistoryConfigsRequest) SetOrderDirection(v string) *ListPdpHistoryConfigsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListPdpHistoryConfigsRequest) SetPageNumber(v int32) *ListPdpHistoryConfigsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPdpHistoryConfigsRequest) SetPageSize(v int32) *ListPdpHistoryConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPdpHistoryConfigsRequest) SetServiceGroupId(v int64) *ListPdpHistoryConfigsRequest {
	s.ServiceGroupId = &v
	return s
}

func (s *ListPdpHistoryConfigsRequest) SetType(v string) *ListPdpHistoryConfigsRequest {
	s.Type = &v
	return s
}

func (s *ListPdpHistoryConfigsRequest) Validate() error {
	return dara.Validate(s)
}
