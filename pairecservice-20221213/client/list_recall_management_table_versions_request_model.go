// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecallManagementTableVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListRecallManagementTableVersionsRequest
	GetInstanceId() *string
	SetOrder(v string) *ListRecallManagementTableVersionsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListRecallManagementTableVersionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRecallManagementTableVersionsRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListRecallManagementTableVersionsRequest
	GetSortBy() *string
}

type ListRecallManagementTableVersionsRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Order      *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListRecallManagementTableVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementTableVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListRecallManagementTableVersionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRecallManagementTableVersionsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListRecallManagementTableVersionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRecallManagementTableVersionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRecallManagementTableVersionsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListRecallManagementTableVersionsRequest) SetInstanceId(v string) *ListRecallManagementTableVersionsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRecallManagementTableVersionsRequest) SetOrder(v string) *ListRecallManagementTableVersionsRequest {
	s.Order = &v
	return s
}

func (s *ListRecallManagementTableVersionsRequest) SetPageNumber(v int32) *ListRecallManagementTableVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRecallManagementTableVersionsRequest) SetPageSize(v int32) *ListRecallManagementTableVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecallManagementTableVersionsRequest) SetSortBy(v string) *ListRecallManagementTableVersionsRequest {
	s.SortBy = &v
	return s
}

func (s *ListRecallManagementTableVersionsRequest) Validate() error {
	return dara.Validate(s)
}
