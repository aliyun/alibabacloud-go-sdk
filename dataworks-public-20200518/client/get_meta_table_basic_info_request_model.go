// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableBasicInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetMetaTableBasicInfoRequest
	GetClusterId() *string
	SetDataSourceType(v string) *GetMetaTableBasicInfoRequest
	GetDataSourceType() *string
	SetDatabaseName(v string) *GetMetaTableBasicInfoRequest
	GetDatabaseName() *string
	SetExtension(v bool) *GetMetaTableBasicInfoRequest
	GetExtension() *bool
	SetTableGuid(v string) *GetMetaTableBasicInfoRequest
	GetTableGuid() *string
	SetTableName(v string) *GetMetaTableBasicInfoRequest
	GetTableName() *string
}

type GetMetaTableBasicInfoRequest struct {
	// The ID of the EMR cluster. This parameter is required only if you set the DataSourceType parameter to emr.
	//
	// You can log on to the [EMR console](https://emr.console.aliyun.com/?spm=a2c4g.11186623.0.0.965cc5c2GeiHet#/cn-hangzhou) to query the ID.
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
	// The name of the metadatabase. This parameter is required only if you set the DataSourceType parameter to emr.
	//
	// You can call the [ListMetaDB](https://help.aliyun.com/document_detail/2780105.html) operation to query the name.
	//
	// example:
	//
	// abc
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// Specifies whether to include extended fields in query results. The extended fields include ReadCount, FavoriteCount, and ViewCount. This parameter takes effect only if you set the DataSourceType parameter to odps.
	//
	// example:
	//
	// false
	Extension *bool `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The GUID of the MaxCompute table. Specify the GUID in the odps.projectName.tableName format.
	//
	// > This parameter is optional for E-MapReduce (EMR) tables.
	//
	// example:
	//
	// odps.engine_name.table_name
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The name of the metatable in the EMR cluster. This parameter is required only if you set the DataSourceType parameter to emr.
	//
	// You can call the [GetMetaDBTableList](https://help.aliyun.com/document_detail/2780086.html) operation to query the name.
	//
	// example:
	//
	// abc
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetMetaTableBasicInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *GetMetaTableBasicInfoRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetMetaTableBasicInfoRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *GetMetaTableBasicInfoRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetMetaTableBasicInfoRequest) GetExtension() *bool {
	return s.Extension
}

func (s *GetMetaTableBasicInfoRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetMetaTableBasicInfoRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetMetaTableBasicInfoRequest) SetClusterId(v string) *GetMetaTableBasicInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *GetMetaTableBasicInfoRequest) SetDataSourceType(v string) *GetMetaTableBasicInfoRequest {
	s.DataSourceType = &v
	return s
}

func (s *GetMetaTableBasicInfoRequest) SetDatabaseName(v string) *GetMetaTableBasicInfoRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetMetaTableBasicInfoRequest) SetExtension(v bool) *GetMetaTableBasicInfoRequest {
	s.Extension = &v
	return s
}

func (s *GetMetaTableBasicInfoRequest) SetTableGuid(v string) *GetMetaTableBasicInfoRequest {
	s.TableGuid = &v
	return s
}

func (s *GetMetaTableBasicInfoRequest) SetTableName(v string) *GetMetaTableBasicInfoRequest {
	s.TableName = &v
	return s
}

func (s *GetMetaTableBasicInfoRequest) Validate() error {
	return dara.Validate(s)
}
