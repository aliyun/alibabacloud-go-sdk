// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStorageTablesInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAscOrder(v bool) *ListStorageTablesInfoRequest
	GetAscOrder() *bool
	SetDate(v string) *ListStorageTablesInfoRequest
	GetDate() *string
	SetOrderColumn(v string) *ListStorageTablesInfoRequest
	GetOrderColumn() *string
	SetPageNumber(v int64) *ListStorageTablesInfoRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListStorageTablesInfoRequest
	GetPageSize() *int64
	SetRecentDays(v int32) *ListStorageTablesInfoRequest
	GetRecentDays() *int32
	SetRegion(v string) *ListStorageTablesInfoRequest
	GetRegion() *string
	SetSchema(v string) *ListStorageTablesInfoRequest
	GetSchema() *string
	SetTablePrefix(v string) *ListStorageTablesInfoRequest
	GetTablePrefix() *string
	SetTenantId(v string) *ListStorageTablesInfoRequest
	GetTenantId() *string
	SetTypes(v []*string) *ListStorageTablesInfoRequest
	GetTypes() []*string
}

type ListStorageTablesInfoRequest struct {
	// Specifies whether to sort data in ascending order.
	//
	// example:
	//
	// false
	AscOrder *bool `json:"ascOrder,omitempty" xml:"ascOrder,omitempty"`
	// The date on which the statistics are collected, in days. Set this parameter to a value in the `YYYYMMdd` format.
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
	// The number of recent days for comparison.
	//
	// example:
	//
	// 1
	RecentDays *int32 `json:"recentDays,omitempty" xml:"recentDays,omitempty"`
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
	// The name of the table that you want to use for fuzzy match.
	//
	// example:
	//
	// bank
	TablePrefix *string `json:"tablePrefix,omitempty" xml:"tablePrefix,omitempty"`
	// The ID of the tenant. You can log on to the MaxCompute console, and choose **Tenants*	- > **Tenant Property*	- from the left-side navigation pane to view the tenant ID.
	//
	// example:
	//
	// 28074710977****
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The storage types.
	Types []*string `json:"types,omitempty" xml:"types,omitempty" type:"Repeated"`
}

func (s ListStorageTablesInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStorageTablesInfoRequest) GoString() string {
	return s.String()
}

func (s *ListStorageTablesInfoRequest) GetAscOrder() *bool {
	return s.AscOrder
}

func (s *ListStorageTablesInfoRequest) GetDate() *string {
	return s.Date
}

func (s *ListStorageTablesInfoRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *ListStorageTablesInfoRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListStorageTablesInfoRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListStorageTablesInfoRequest) GetRecentDays() *int32 {
	return s.RecentDays
}

func (s *ListStorageTablesInfoRequest) GetRegion() *string {
	return s.Region
}

func (s *ListStorageTablesInfoRequest) GetSchema() *string {
	return s.Schema
}

func (s *ListStorageTablesInfoRequest) GetTablePrefix() *string {
	return s.TablePrefix
}

func (s *ListStorageTablesInfoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListStorageTablesInfoRequest) GetTypes() []*string {
	return s.Types
}

func (s *ListStorageTablesInfoRequest) SetAscOrder(v bool) *ListStorageTablesInfoRequest {
	s.AscOrder = &v
	return s
}

func (s *ListStorageTablesInfoRequest) SetDate(v string) *ListStorageTablesInfoRequest {
	s.Date = &v
	return s
}

func (s *ListStorageTablesInfoRequest) SetOrderColumn(v string) *ListStorageTablesInfoRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListStorageTablesInfoRequest) SetPageNumber(v int64) *ListStorageTablesInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStorageTablesInfoRequest) SetPageSize(v int64) *ListStorageTablesInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListStorageTablesInfoRequest) SetRecentDays(v int32) *ListStorageTablesInfoRequest {
	s.RecentDays = &v
	return s
}

func (s *ListStorageTablesInfoRequest) SetRegion(v string) *ListStorageTablesInfoRequest {
	s.Region = &v
	return s
}

func (s *ListStorageTablesInfoRequest) SetSchema(v string) *ListStorageTablesInfoRequest {
	s.Schema = &v
	return s
}

func (s *ListStorageTablesInfoRequest) SetTablePrefix(v string) *ListStorageTablesInfoRequest {
	s.TablePrefix = &v
	return s
}

func (s *ListStorageTablesInfoRequest) SetTenantId(v string) *ListStorageTablesInfoRequest {
	s.TenantId = &v
	return s
}

func (s *ListStorageTablesInfoRequest) SetTypes(v []*string) *ListStorageTablesInfoRequest {
	s.Types = v
	return s
}

func (s *ListStorageTablesInfoRequest) Validate() error {
	return dara.Validate(s)
}
