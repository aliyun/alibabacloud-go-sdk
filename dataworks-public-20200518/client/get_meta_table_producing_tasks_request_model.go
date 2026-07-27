// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableProducingTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetMetaTableProducingTasksRequest
	GetClusterId() *string
	SetDataSourceType(v string) *GetMetaTableProducingTasksRequest
	GetDataSourceType() *string
	SetDbName(v string) *GetMetaTableProducingTasksRequest
	GetDbName() *string
	SetSchemaName(v string) *GetMetaTableProducingTasksRequest
	GetSchemaName() *string
	SetTableGuid(v string) *GetMetaTableProducingTasksRequest
	GetTableGuid() *string
	SetTableName(v string) *GetMetaTableProducingTasksRequest
	GetTableName() *string
}

type GetMetaTableProducingTasksRequest struct {
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	SchemaName     *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// This parameter is required.
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetMetaTableProducingTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableProducingTasksRequest) GoString() string {
	return s.String()
}

func (s *GetMetaTableProducingTasksRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetMetaTableProducingTasksRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *GetMetaTableProducingTasksRequest) GetDbName() *string {
	return s.DbName
}

func (s *GetMetaTableProducingTasksRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetMetaTableProducingTasksRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetMetaTableProducingTasksRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetMetaTableProducingTasksRequest) SetClusterId(v string) *GetMetaTableProducingTasksRequest {
	s.ClusterId = &v
	return s
}

func (s *GetMetaTableProducingTasksRequest) SetDataSourceType(v string) *GetMetaTableProducingTasksRequest {
	s.DataSourceType = &v
	return s
}

func (s *GetMetaTableProducingTasksRequest) SetDbName(v string) *GetMetaTableProducingTasksRequest {
	s.DbName = &v
	return s
}

func (s *GetMetaTableProducingTasksRequest) SetSchemaName(v string) *GetMetaTableProducingTasksRequest {
	s.SchemaName = &v
	return s
}

func (s *GetMetaTableProducingTasksRequest) SetTableGuid(v string) *GetMetaTableProducingTasksRequest {
	s.TableGuid = &v
	return s
}

func (s *GetMetaTableProducingTasksRequest) SetTableName(v string) *GetMetaTableProducingTasksRequest {
	s.TableName = &v
	return s
}

func (s *GetMetaTableProducingTasksRequest) Validate() error {
	return dara.Validate(s)
}
