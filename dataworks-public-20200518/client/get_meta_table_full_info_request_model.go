// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableFullInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetMetaTableFullInfoRequest
	GetClusterId() *string
	SetDataSourceType(v string) *GetMetaTableFullInfoRequest
	GetDataSourceType() *string
	SetDatabaseName(v string) *GetMetaTableFullInfoRequest
	GetDatabaseName() *string
	SetPageNum(v int32) *GetMetaTableFullInfoRequest
	GetPageNum() *int32
	SetPageSize(v int32) *GetMetaTableFullInfoRequest
	GetPageSize() *int32
	SetTableGuid(v string) *GetMetaTableFullInfoRequest
	GetTableGuid() *string
	SetTableName(v string) *GetMetaTableFullInfoRequest
	GetTableName() *string
}

type GetMetaTableFullInfoRequest struct {
	// The ID of the EMR cluster. This parameter is required only if you set the DataSourceType parameter to emr.
	//
	// You can log on to the [EMR console](https://emr.console.aliyun.com/?spm=a2c4g.11186623.0.0.965cc5c2GeiHet#/cn-hangzhou) to query the ID.
	//
	// example:
	//
	// C-010A704DA760****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the data source. Set the value to emr.
	//
	// example:
	//
	// emr
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The name of the database. This parameter is required only if you set the DataSourceType parameter to emr.
	//
	// You can call the [ListMetaDB](https://help.aliyun.com/document_detail/185662.html) operation to query the database name.
	//
	// example:
	//
	// abc
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The page number requested for pagination.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of items per page, with a default of 10 and a maximum of 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique identifier of the table. You can call the [GetMetaDBTableList](https://help.aliyun.com/document_detail/173916.html) operation to query the unique identifier of the table.
	//
	// example:
	//
	// odps.engine_name.table_name
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The name of the table in the EMR cluster. This parameter is required only if you set the DataSourceType parameter to emr.
	//
	// You can call the [GetMetaDBTableList](https://help.aliyun.com/document_detail/173916.html) operation to query the table name.
	//
	// example:
	//
	// abc
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetMetaTableFullInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableFullInfoRequest) GoString() string {
	return s.String()
}

func (s *GetMetaTableFullInfoRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetMetaTableFullInfoRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *GetMetaTableFullInfoRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetMetaTableFullInfoRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetMetaTableFullInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMetaTableFullInfoRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetMetaTableFullInfoRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetMetaTableFullInfoRequest) SetClusterId(v string) *GetMetaTableFullInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *GetMetaTableFullInfoRequest) SetDataSourceType(v string) *GetMetaTableFullInfoRequest {
	s.DataSourceType = &v
	return s
}

func (s *GetMetaTableFullInfoRequest) SetDatabaseName(v string) *GetMetaTableFullInfoRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetMetaTableFullInfoRequest) SetPageNum(v int32) *GetMetaTableFullInfoRequest {
	s.PageNum = &v
	return s
}

func (s *GetMetaTableFullInfoRequest) SetPageSize(v int32) *GetMetaTableFullInfoRequest {
	s.PageSize = &v
	return s
}

func (s *GetMetaTableFullInfoRequest) SetTableGuid(v string) *GetMetaTableFullInfoRequest {
	s.TableGuid = &v
	return s
}

func (s *GetMetaTableFullInfoRequest) SetTableName(v string) *GetMetaTableFullInfoRequest {
	s.TableName = &v
	return s
}

func (s *GetMetaTableFullInfoRequest) Validate() error {
	return dara.Validate(s)
}
