// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaColumnLineageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetMetaColumnLineageRequest
	GetClusterId() *string
	SetColumnGuid(v string) *GetMetaColumnLineageRequest
	GetColumnGuid() *string
	SetColumnName(v string) *GetMetaColumnLineageRequest
	GetColumnName() *string
	SetDataSourceType(v string) *GetMetaColumnLineageRequest
	GetDataSourceType() *string
	SetDatabaseName(v string) *GetMetaColumnLineageRequest
	GetDatabaseName() *string
	SetDirection(v string) *GetMetaColumnLineageRequest
	GetDirection() *string
	SetPageNum(v int32) *GetMetaColumnLineageRequest
	GetPageNum() *int32
	SetPageSize(v int32) *GetMetaColumnLineageRequest
	GetPageSize() *int32
	SetTableName(v string) *GetMetaColumnLineageRequest
	GetTableName() *string
}

type GetMetaColumnLineageRequest struct {
	// The ID of the E-MapReduce (EMR) cluster. Configure this parameter only when you query data in an EMR compute engine instance.
	//
	// example:
	//
	// abc
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The unique identifier of the field.
	//
	// example:
	//
	// odps.engine_name.table_name.column_name
	ColumnGuid *string `json:"ColumnGuid,omitempty" xml:"ColumnGuid,omitempty"`
	// The name of the field.
	//
	// example:
	//
	// abc
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The type of the data source. Valid values: odps and emr.
	//
	// example:
	//
	// emr
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// abc
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// Specifies whether to query the ancestor or descendant lineage of the field. The value up indicates the ancestor lineage. The value down indicates the descendant lineage.
	//
	// This parameter is required.
	//
	// example:
	//
	// up
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the metatable.
	//
	// example:
	//
	// abc
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetMetaColumnLineageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaColumnLineageRequest) GoString() string {
	return s.String()
}

func (s *GetMetaColumnLineageRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetMetaColumnLineageRequest) GetColumnGuid() *string {
	return s.ColumnGuid
}

func (s *GetMetaColumnLineageRequest) GetColumnName() *string {
	return s.ColumnName
}

func (s *GetMetaColumnLineageRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *GetMetaColumnLineageRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetMetaColumnLineageRequest) GetDirection() *string {
	return s.Direction
}

func (s *GetMetaColumnLineageRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetMetaColumnLineageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMetaColumnLineageRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetMetaColumnLineageRequest) SetClusterId(v string) *GetMetaColumnLineageRequest {
	s.ClusterId = &v
	return s
}

func (s *GetMetaColumnLineageRequest) SetColumnGuid(v string) *GetMetaColumnLineageRequest {
	s.ColumnGuid = &v
	return s
}

func (s *GetMetaColumnLineageRequest) SetColumnName(v string) *GetMetaColumnLineageRequest {
	s.ColumnName = &v
	return s
}

func (s *GetMetaColumnLineageRequest) SetDataSourceType(v string) *GetMetaColumnLineageRequest {
	s.DataSourceType = &v
	return s
}

func (s *GetMetaColumnLineageRequest) SetDatabaseName(v string) *GetMetaColumnLineageRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetMetaColumnLineageRequest) SetDirection(v string) *GetMetaColumnLineageRequest {
	s.Direction = &v
	return s
}

func (s *GetMetaColumnLineageRequest) SetPageNum(v int32) *GetMetaColumnLineageRequest {
	s.PageNum = &v
	return s
}

func (s *GetMetaColumnLineageRequest) SetPageSize(v int32) *GetMetaColumnLineageRequest {
	s.PageSize = &v
	return s
}

func (s *GetMetaColumnLineageRequest) SetTableName(v string) *GetMetaColumnLineageRequest {
	s.TableName = &v
	return s
}

func (s *GetMetaColumnLineageRequest) Validate() error {
	return dara.Validate(s)
}
