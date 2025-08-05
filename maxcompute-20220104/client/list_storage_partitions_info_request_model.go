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
	// Specifies whether to sort data in ascending order.
	//
	// example:
	//
	// false
	AscOrder *bool `json:"ascOrder,omitempty" xml:"ascOrder,omitempty"`
	// The date on which the statistics are collected, in days. Set this parameter to a value in the YYYYMMdd format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20241205
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The sorting column.
	//
	// example:
	//
	// totalFrequency
	OrderColumn *string `json:"orderColumn,omitempty" xml:"orderColumn,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The name of the partition that you want to use for fuzzy match.
	//
	// example:
	//
	// ds=20241201
	PartitionPrefix *string `json:"partitionPrefix,omitempty" xml:"partitionPrefix,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The name of the schema.
	//
	// example:
	//
	// schema
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
	// The ID of the tenant. You can log on to the MaxCompute console, and choose **Tenants*	- > **Tenant Property*	- from the left-side navigation pane to view the tenant ID.
	//
	// example:
	//
	// 40713753659****
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The storage types.
	Types []*string `json:"types,omitempty" xml:"types,omitempty" type:"Repeated"`
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
