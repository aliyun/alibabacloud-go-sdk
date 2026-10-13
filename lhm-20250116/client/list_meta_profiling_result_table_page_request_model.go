// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaProfilingResultTablePageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentType(v int32) *ListMetaProfilingResultTablePageRequest
	GetComponentType() *int32
	SetDataRange(v string) *ListMetaProfilingResultTablePageRequest
	GetDataRange() *string
	SetDataSize(v float64) *ListMetaProfilingResultTablePageRequest
	GetDataSize() *float64
	SetDataUnit(v string) *ListMetaProfilingResultTablePageRequest
	GetDataUnit() *string
	SetDbName(v string) *ListMetaProfilingResultTablePageRequest
	GetDbName() *string
	SetGroupBy(v string) *ListMetaProfilingResultTablePageRequest
	GetGroupBy() *string
	SetIsPartitioned(v int32) *ListMetaProfilingResultTablePageRequest
	GetIsPartitioned() *int32
	SetIsSummary(v bool) *ListMetaProfilingResultTablePageRequest
	GetIsSummary() *bool
	SetMigrationStatus(v string) *ListMetaProfilingResultTablePageRequest
	GetMigrationStatus() *string
	SetNeedTotalCount(v bool) *ListMetaProfilingResultTablePageRequest
	GetNeedTotalCount() *bool
	SetOrderBy(v string) *ListMetaProfilingResultTablePageRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListMetaProfilingResultTablePageRequest
	GetOrderDirection() *string
	SetPageIndex(v int32) *ListMetaProfilingResultTablePageRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListMetaProfilingResultTablePageRequest
	GetPageSize() *int32
	SetTableName(v string) *ListMetaProfilingResultTablePageRequest
	GetTableName() *string
	SetTaskId(v int64) *ListMetaProfilingResultTablePageRequest
	GetTaskId() *int64
	SetType(v string) *ListMetaProfilingResultTablePageRequest
	GetType() *string
}

type ListMetaProfilingResultTablePageRequest struct {
	// The identifier of the component (data source) type.
	//
	// example:
	//
	// 1
	ComponentType *int32 `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// The filter condition for the data range.
	//
	// example:
	//
	// LAST_7_DAYS
	DataRange *string `json:"dataRange,omitempty" xml:"dataRange,omitempty"`
	// The threshold for filtering by data size. This parameter must be used together with dataUnit.
	//
	// example:
	//
	// 1024.0
	DataSize *float64 `json:"dataSize,omitempty" xml:"dataSize,omitempty"`
	// The unit of data size, such as GB or TB. This parameter must be used together with dataSize.
	//
	// example:
	//
	// GB
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// The name of the database, which is used to filter results by database.
	//
	// example:
	//
	// ods
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// The name of the group-by field, which is used for aggregate statistics.
	//
	// example:
	//
	// dbName
	GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	// Specifies whether to filter only partitioned tables.
	//
	// example:
	//
	// 1
	IsPartitioned *int32 `json:"isPartitioned,omitempty" xml:"isPartitioned,omitempty"`
	// Specifies whether to return only summary statistics.
	IsSummary *bool `json:"isSummary,omitempty" xml:"isSummary,omitempty"`
	// The migration status, which is used to filter by migration progress or indicate the migration status of an object.
	//
	// example:
	//
	// DONE
	MigrationStatus *string `json:"migrationStatus,omitempty" xml:"migrationStatus,omitempty"`
	// Specifies whether to return the total number of records.
	NeedTotalCount *bool `json:"needTotalCount,omitempty" xml:"needTotalCount,omitempty"`
	// The name of the order-by field.
	//
	// example:
	//
	// gmtCreate
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// The sort direction, such as ascending (ASC) or descending (DESC).
	//
	// example:
	//
	// DESC
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// The current page number, starting from 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of records to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The name of the table, which is used to filter results by table name.
	//
	// example:
	//
	// ods_order_detail
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// The ID of the metadata profiling task, which is used to locate the profiling task to which this query belongs.
	//
	// example:
	//
	// 100256
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The identifier of the object type.
	//
	// example:
	//
	// TABLE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMetaProfilingResultTablePageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMetaProfilingResultTablePageRequest) GoString() string {
	return s.String()
}

func (s *ListMetaProfilingResultTablePageRequest) GetComponentType() *int32 {
	return s.ComponentType
}

func (s *ListMetaProfilingResultTablePageRequest) GetDataRange() *string {
	return s.DataRange
}

func (s *ListMetaProfilingResultTablePageRequest) GetDataSize() *float64 {
	return s.DataSize
}

func (s *ListMetaProfilingResultTablePageRequest) GetDataUnit() *string {
	return s.DataUnit
}

func (s *ListMetaProfilingResultTablePageRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListMetaProfilingResultTablePageRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *ListMetaProfilingResultTablePageRequest) GetIsPartitioned() *int32 {
	return s.IsPartitioned
}

func (s *ListMetaProfilingResultTablePageRequest) GetIsSummary() *bool {
	return s.IsSummary
}

func (s *ListMetaProfilingResultTablePageRequest) GetMigrationStatus() *string {
	return s.MigrationStatus
}

func (s *ListMetaProfilingResultTablePageRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *ListMetaProfilingResultTablePageRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMetaProfilingResultTablePageRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListMetaProfilingResultTablePageRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListMetaProfilingResultTablePageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMetaProfilingResultTablePageRequest) GetTableName() *string {
	return s.TableName
}

func (s *ListMetaProfilingResultTablePageRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListMetaProfilingResultTablePageRequest) GetType() *string {
	return s.Type
}

func (s *ListMetaProfilingResultTablePageRequest) SetComponentType(v int32) *ListMetaProfilingResultTablePageRequest {
	s.ComponentType = &v
	return s
}

func (s *ListMetaProfilingResultTablePageRequest) SetDataRange(v string) *ListMetaProfilingResultTablePageRequest {
	s.DataRange = &v
	return s
}

func (s *ListMetaProfilingResultTablePageRequest) SetDataSize(v float64) *ListMetaProfilingResultTablePageRequest {
	s.DataSize = &v
	return s
}

func (s *ListMetaProfilingResultTablePageRequest) SetDataUnit(v string) *ListMetaProfilingResultTablePageRequest {
	s.DataUnit = &v
	return s
}

func (s *ListMetaProfilingResultTablePageRequest) SetDbName(v string) *ListMetaProfilingResultTablePageRequest {
	s.DbName = &v
	return s
}

func (s *ListMetaProfilingResultTablePageRequest) SetGroupBy(v string) *ListMetaProfilingResultTablePageRequest {
	s.GroupBy = &v
	return s
}

func (s *ListMetaProfilingResultTablePageRequest) SetIsPartitioned(v int32) *ListMetaProfilingResultTablePageRequest {
	s.IsPartitioned = &v
	return s
}

func (s *ListMetaProfilingResultTablePageRequest) SetIsSummary(v bool) *ListMetaProfilingResultTablePageRequest {
	s.IsSummary = &v
	return s
}

func (s *ListMetaProfilingResultTablePageRequest) SetMigrationStatus(v string) *ListMetaProfilingResultTablePageRequest {
	s.MigrationStatus = &v
	return s
}

func (s *ListMetaProfilingResultTablePageRequest) SetNeedTotalCount(v bool) *ListMetaProfilingResultTablePageRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *ListMetaProfilingResultTablePageRequest) SetOrderBy(v string) *ListMetaProfilingResultTablePageRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMetaProfilingResultTablePageRequest) SetOrderDirection(v string) *ListMetaProfilingResultTablePageRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListMetaProfilingResultTablePageRequest) SetPageIndex(v int32) *ListMetaProfilingResultTablePageRequest {
	s.PageIndex = &v
	return s
}

func (s *ListMetaProfilingResultTablePageRequest) SetPageSize(v int32) *ListMetaProfilingResultTablePageRequest {
	s.PageSize = &v
	return s
}

func (s *ListMetaProfilingResultTablePageRequest) SetTableName(v string) *ListMetaProfilingResultTablePageRequest {
	s.TableName = &v
	return s
}

func (s *ListMetaProfilingResultTablePageRequest) SetTaskId(v int64) *ListMetaProfilingResultTablePageRequest {
	s.TaskId = &v
	return s
}

func (s *ListMetaProfilingResultTablePageRequest) SetType(v string) *ListMetaProfilingResultTablePageRequest {
	s.Type = &v
	return s
}

func (s *ListMetaProfilingResultTablePageRequest) Validate() error {
	return dara.Validate(s)
}
