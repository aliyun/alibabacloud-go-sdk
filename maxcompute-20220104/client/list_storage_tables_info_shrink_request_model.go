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
	AscOrder *bool `json:"ascOrder,omitempty" xml:"ascOrder,omitempty"`
	// This parameter is required.
	Date        *string `json:"date,omitempty" xml:"date,omitempty"`
	OrderColumn *string `json:"orderColumn,omitempty" xml:"orderColumn,omitempty"`
	PageNumber  *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize    *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RecentDays  *int32  `json:"recentDays,omitempty" xml:"recentDays,omitempty"`
	Region      *string `json:"region,omitempty" xml:"region,omitempty"`
	Schema      *string `json:"schema,omitempty" xml:"schema,omitempty"`
	TablePrefix *string `json:"tablePrefix,omitempty" xml:"tablePrefix,omitempty"`
	TenantId    *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
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
