// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStorageTablesInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAscOrder(v bool) *ListStorageTablesInfoShrinkRequest
	GetAscOrder() *bool
	SetDate(v string) *ListStorageTablesInfoShrinkRequest
	GetDate() *string
	SetOrderColumn(v string) *ListStorageTablesInfoShrinkRequest
	GetOrderColumn() *string
	SetPageNumber(v int64) *ListStorageTablesInfoShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListStorageTablesInfoShrinkRequest
	GetPageSize() *int64
	SetRecentDays(v int32) *ListStorageTablesInfoShrinkRequest
	GetRecentDays() *int32
	SetRegion(v string) *ListStorageTablesInfoShrinkRequest
	GetRegion() *string
	SetSchema(v string) *ListStorageTablesInfoShrinkRequest
	GetSchema() *string
	SetTablePrefix(v string) *ListStorageTablesInfoShrinkRequest
	GetTablePrefix() *string
	SetTenantId(v string) *ListStorageTablesInfoShrinkRequest
	GetTenantId() *string
	SetTypesShrink(v string) *ListStorageTablesInfoShrinkRequest
	GetTypesShrink() *string
}

type ListStorageTablesInfoShrinkRequest struct {
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
	TypesShrink *string `json:"types,omitempty" xml:"types,omitempty"`
}

func (s ListStorageTablesInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStorageTablesInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListStorageTablesInfoShrinkRequest) GetAscOrder() *bool {
	return s.AscOrder
}

func (s *ListStorageTablesInfoShrinkRequest) GetDate() *string {
	return s.Date
}

func (s *ListStorageTablesInfoShrinkRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *ListStorageTablesInfoShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListStorageTablesInfoShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListStorageTablesInfoShrinkRequest) GetRecentDays() *int32 {
	return s.RecentDays
}

func (s *ListStorageTablesInfoShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *ListStorageTablesInfoShrinkRequest) GetSchema() *string {
	return s.Schema
}

func (s *ListStorageTablesInfoShrinkRequest) GetTablePrefix() *string {
	return s.TablePrefix
}

func (s *ListStorageTablesInfoShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListStorageTablesInfoShrinkRequest) GetTypesShrink() *string {
	return s.TypesShrink
}

func (s *ListStorageTablesInfoShrinkRequest) SetAscOrder(v bool) *ListStorageTablesInfoShrinkRequest {
	s.AscOrder = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) SetDate(v string) *ListStorageTablesInfoShrinkRequest {
	s.Date = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) SetOrderColumn(v string) *ListStorageTablesInfoShrinkRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) SetPageNumber(v int64) *ListStorageTablesInfoShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) SetPageSize(v int64) *ListStorageTablesInfoShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) SetRecentDays(v int32) *ListStorageTablesInfoShrinkRequest {
	s.RecentDays = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) SetRegion(v string) *ListStorageTablesInfoShrinkRequest {
	s.Region = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) SetSchema(v string) *ListStorageTablesInfoShrinkRequest {
	s.Schema = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) SetTablePrefix(v string) *ListStorageTablesInfoShrinkRequest {
	s.TablePrefix = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) SetTenantId(v string) *ListStorageTablesInfoShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) SetTypesShrink(v string) *ListStorageTablesInfoShrinkRequest {
	s.TypesShrink = &v
	return s
}

func (s *ListStorageTablesInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
