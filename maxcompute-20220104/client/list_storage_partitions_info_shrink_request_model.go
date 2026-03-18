// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStoragePartitionsInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAscOrder(v bool) *ListStoragePartitionsInfoShrinkRequest
	GetAscOrder() *bool
	SetDate(v string) *ListStoragePartitionsInfoShrinkRequest
	GetDate() *string
	SetOrderColumn(v string) *ListStoragePartitionsInfoShrinkRequest
	GetOrderColumn() *string
	SetPageNumber(v int64) *ListStoragePartitionsInfoShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListStoragePartitionsInfoShrinkRequest
	GetPageSize() *int64
	SetPartitionPrefix(v string) *ListStoragePartitionsInfoShrinkRequest
	GetPartitionPrefix() *string
	SetRegion(v string) *ListStoragePartitionsInfoShrinkRequest
	GetRegion() *string
	SetSchema(v string) *ListStoragePartitionsInfoShrinkRequest
	GetSchema() *string
	SetTenantId(v string) *ListStoragePartitionsInfoShrinkRequest
	GetTenantId() *string
	SetTypesShrink(v string) *ListStoragePartitionsInfoShrinkRequest
	GetTypesShrink() *string
}

type ListStoragePartitionsInfoShrinkRequest struct {
	AscOrder *bool `json:"ascOrder,omitempty" xml:"ascOrder,omitempty"`
	// This parameter is required.
	Date            *string `json:"date,omitempty" xml:"date,omitempty"`
	OrderColumn     *string `json:"orderColumn,omitempty" xml:"orderColumn,omitempty"`
	PageNumber      *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize        *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	PartitionPrefix *string `json:"partitionPrefix,omitempty" xml:"partitionPrefix,omitempty"`
	Region          *string `json:"region,omitempty" xml:"region,omitempty"`
	Schema          *string `json:"schema,omitempty" xml:"schema,omitempty"`
	TenantId        *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	TypesShrink     *string `json:"types,omitempty" xml:"types,omitempty"`
}

func (s ListStoragePartitionsInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStoragePartitionsInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListStoragePartitionsInfoShrinkRequest) GetAscOrder() *bool {
	return s.AscOrder
}

func (s *ListStoragePartitionsInfoShrinkRequest) GetDate() *string {
	return s.Date
}

func (s *ListStoragePartitionsInfoShrinkRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *ListStoragePartitionsInfoShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListStoragePartitionsInfoShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListStoragePartitionsInfoShrinkRequest) GetPartitionPrefix() *string {
	return s.PartitionPrefix
}

func (s *ListStoragePartitionsInfoShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *ListStoragePartitionsInfoShrinkRequest) GetSchema() *string {
	return s.Schema
}

func (s *ListStoragePartitionsInfoShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListStoragePartitionsInfoShrinkRequest) GetTypesShrink() *string {
	return s.TypesShrink
}

func (s *ListStoragePartitionsInfoShrinkRequest) SetAscOrder(v bool) *ListStoragePartitionsInfoShrinkRequest {
	s.AscOrder = &v
	return s
}

func (s *ListStoragePartitionsInfoShrinkRequest) SetDate(v string) *ListStoragePartitionsInfoShrinkRequest {
	s.Date = &v
	return s
}

func (s *ListStoragePartitionsInfoShrinkRequest) SetOrderColumn(v string) *ListStoragePartitionsInfoShrinkRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListStoragePartitionsInfoShrinkRequest) SetPageNumber(v int64) *ListStoragePartitionsInfoShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStoragePartitionsInfoShrinkRequest) SetPageSize(v int64) *ListStoragePartitionsInfoShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListStoragePartitionsInfoShrinkRequest) SetPartitionPrefix(v string) *ListStoragePartitionsInfoShrinkRequest {
	s.PartitionPrefix = &v
	return s
}

func (s *ListStoragePartitionsInfoShrinkRequest) SetRegion(v string) *ListStoragePartitionsInfoShrinkRequest {
	s.Region = &v
	return s
}

func (s *ListStoragePartitionsInfoShrinkRequest) SetSchema(v string) *ListStoragePartitionsInfoShrinkRequest {
	s.Schema = &v
	return s
}

func (s *ListStoragePartitionsInfoShrinkRequest) SetTenantId(v string) *ListStoragePartitionsInfoShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListStoragePartitionsInfoShrinkRequest) SetTypesShrink(v string) *ListStoragePartitionsInfoShrinkRequest {
	s.TypesShrink = &v
	return s
}

func (s *ListStoragePartitionsInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
