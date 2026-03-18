// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStoragePartitionsInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAscOrder(v bool) *ListStoragePartitionsInfoRequest
	GetAscOrder() *bool
	SetDate(v string) *ListStoragePartitionsInfoRequest
	GetDate() *string
	SetOrderColumn(v string) *ListStoragePartitionsInfoRequest
	GetOrderColumn() *string
	SetPageNumber(v int64) *ListStoragePartitionsInfoRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListStoragePartitionsInfoRequest
	GetPageSize() *int64
	SetPartitionPrefix(v string) *ListStoragePartitionsInfoRequest
	GetPartitionPrefix() *string
	SetRegion(v string) *ListStoragePartitionsInfoRequest
	GetRegion() *string
	SetSchema(v string) *ListStoragePartitionsInfoRequest
	GetSchema() *string
	SetTenantId(v string) *ListStoragePartitionsInfoRequest
	GetTenantId() *string
	SetTypes(v []*string) *ListStoragePartitionsInfoRequest
	GetTypes() []*string
}

type ListStoragePartitionsInfoRequest struct {
	AscOrder *bool `json:"ascOrder,omitempty" xml:"ascOrder,omitempty"`
	// This parameter is required.
	Date            *string   `json:"date,omitempty" xml:"date,omitempty"`
	OrderColumn     *string   `json:"orderColumn,omitempty" xml:"orderColumn,omitempty"`
	PageNumber      *int64    `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize        *int64    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	PartitionPrefix *string   `json:"partitionPrefix,omitempty" xml:"partitionPrefix,omitempty"`
	Region          *string   `json:"region,omitempty" xml:"region,omitempty"`
	Schema          *string   `json:"schema,omitempty" xml:"schema,omitempty"`
	TenantId        *string   `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	Types           []*string `json:"types,omitempty" xml:"types,omitempty" type:"Repeated"`
}

func (s ListStoragePartitionsInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStoragePartitionsInfoRequest) GoString() string {
	return s.String()
}

func (s *ListStoragePartitionsInfoRequest) GetAscOrder() *bool {
	return s.AscOrder
}

func (s *ListStoragePartitionsInfoRequest) GetDate() *string {
	return s.Date
}

func (s *ListStoragePartitionsInfoRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *ListStoragePartitionsInfoRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListStoragePartitionsInfoRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListStoragePartitionsInfoRequest) GetPartitionPrefix() *string {
	return s.PartitionPrefix
}

func (s *ListStoragePartitionsInfoRequest) GetRegion() *string {
	return s.Region
}

func (s *ListStoragePartitionsInfoRequest) GetSchema() *string {
	return s.Schema
}

func (s *ListStoragePartitionsInfoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListStoragePartitionsInfoRequest) GetTypes() []*string {
	return s.Types
}

func (s *ListStoragePartitionsInfoRequest) SetAscOrder(v bool) *ListStoragePartitionsInfoRequest {
	s.AscOrder = &v
	return s
}

func (s *ListStoragePartitionsInfoRequest) SetDate(v string) *ListStoragePartitionsInfoRequest {
	s.Date = &v
	return s
}

func (s *ListStoragePartitionsInfoRequest) SetOrderColumn(v string) *ListStoragePartitionsInfoRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListStoragePartitionsInfoRequest) SetPageNumber(v int64) *ListStoragePartitionsInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStoragePartitionsInfoRequest) SetPageSize(v int64) *ListStoragePartitionsInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListStoragePartitionsInfoRequest) SetPartitionPrefix(v string) *ListStoragePartitionsInfoRequest {
	s.PartitionPrefix = &v
	return s
}

func (s *ListStoragePartitionsInfoRequest) SetRegion(v string) *ListStoragePartitionsInfoRequest {
	s.Region = &v
	return s
}

func (s *ListStoragePartitionsInfoRequest) SetSchema(v string) *ListStoragePartitionsInfoRequest {
	s.Schema = &v
	return s
}

func (s *ListStoragePartitionsInfoRequest) SetTenantId(v string) *ListStoragePartitionsInfoRequest {
	s.TenantId = &v
	return s
}

func (s *ListStoragePartitionsInfoRequest) SetTypes(v []*string) *ListStoragePartitionsInfoRequest {
	s.Types = v
	return s
}

func (s *ListStoragePartitionsInfoRequest) Validate() error {
	return dara.Validate(s)
}
