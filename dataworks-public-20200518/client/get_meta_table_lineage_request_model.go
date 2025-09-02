// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableLineageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetMetaTableLineageRequest
	GetClusterId() *string
	SetDataSourceType(v string) *GetMetaTableLineageRequest
	GetDataSourceType() *string
	SetDatabaseName(v string) *GetMetaTableLineageRequest
	GetDatabaseName() *string
	SetDirection(v string) *GetMetaTableLineageRequest
	GetDirection() *string
	SetNextPrimaryKey(v string) *GetMetaTableLineageRequest
	GetNextPrimaryKey() *string
	SetPageSize(v int32) *GetMetaTableLineageRequest
	GetPageSize() *int32
	SetTableGuid(v string) *GetMetaTableLineageRequest
	GetTableGuid() *string
	SetTableName(v string) *GetMetaTableLineageRequest
	GetTableName() *string
}

type GetMetaTableLineageRequest struct {
	// The ID of the E-MapReduce (EMR) cluster. Configure this parameter only if you want to query the lineage of an EMR table.
	//
	// example:
	//
	// abc
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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
	// Specifies the ancestor or descendant lineage that you want to query for a field. Valid values: up and down. The value up indicates the ancestor lineage. The value down indicates the descendant lineage.
	//
	// This parameter is required.
	//
	// example:
	//
	// up
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The logic of paging. Configure this parameter based on the value of the response parameter NextPrimaryKey when the value of the response parameter HasNext is true in the previous request.
	//
	// example:
	//
	// odps.engine_name.table_name1
	NextPrimaryKey *string `json:"NextPrimaryKey,omitempty" xml:"NextPrimaryKey,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique identifier of the table.
	//
	// example:
	//
	// odps.engine_name.table_name
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// abc
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetMetaTableLineageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableLineageRequest) GoString() string {
	return s.String()
}

func (s *GetMetaTableLineageRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetMetaTableLineageRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *GetMetaTableLineageRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetMetaTableLineageRequest) GetDirection() *string {
	return s.Direction
}

func (s *GetMetaTableLineageRequest) GetNextPrimaryKey() *string {
	return s.NextPrimaryKey
}

func (s *GetMetaTableLineageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMetaTableLineageRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetMetaTableLineageRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetMetaTableLineageRequest) SetClusterId(v string) *GetMetaTableLineageRequest {
	s.ClusterId = &v
	return s
}

func (s *GetMetaTableLineageRequest) SetDataSourceType(v string) *GetMetaTableLineageRequest {
	s.DataSourceType = &v
	return s
}

func (s *GetMetaTableLineageRequest) SetDatabaseName(v string) *GetMetaTableLineageRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetMetaTableLineageRequest) SetDirection(v string) *GetMetaTableLineageRequest {
	s.Direction = &v
	return s
}

func (s *GetMetaTableLineageRequest) SetNextPrimaryKey(v string) *GetMetaTableLineageRequest {
	s.NextPrimaryKey = &v
	return s
}

func (s *GetMetaTableLineageRequest) SetPageSize(v int32) *GetMetaTableLineageRequest {
	s.PageSize = &v
	return s
}

func (s *GetMetaTableLineageRequest) SetTableGuid(v string) *GetMetaTableLineageRequest {
	s.TableGuid = &v
	return s
}

func (s *GetMetaTableLineageRequest) SetTableName(v string) *GetMetaTableLineageRequest {
	s.TableName = &v
	return s
}

func (s *GetMetaTableLineageRequest) Validate() error {
	return dara.Validate(s)
}
