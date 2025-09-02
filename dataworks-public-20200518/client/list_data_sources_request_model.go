// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceType(v string) *ListDataSourcesRequest
	GetDataSourceType() *string
	SetEnvType(v int32) *ListDataSourcesRequest
	GetEnvType() *int32
	SetName(v string) *ListDataSourcesRequest
	GetName() *string
	SetPageNumber(v int32) *ListDataSourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataSourcesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDataSourcesRequest
	GetProjectId() *int64
	SetStatus(v string) *ListDataSourcesRequest
	GetStatus() *string
	SetSubType(v string) *ListDataSourcesRequest
	GetSubType() *string
}

type ListDataSourcesRequest struct {
	// The type of the data source. Valid values:
	//
	// 	- odps
	//
	// 	- mysql
	//
	// 	- rds
	//
	// 	- oss
	//
	// 	- sqlserver
	//
	// 	- polardb
	//
	// 	- oracle
	//
	// 	- mongodb
	//
	// 	- emr
	//
	// 	- postgresql
	//
	// 	- analyticdb_for_mysql
	//
	// 	- hybriddb_for_postgresql
	//
	// 	- holo
	//
	// example:
	//
	// rds
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The environment in which the data source is used. Valid values: 0 and 1. The value 0 indicates development environment. The value 1 indicates production environment.
	//
	// example:
	//
	// 1
	EnvType *int32 `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The name of the data source that you want to query.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the DataWorks workspace to which the data sources belong. You can call the [ListProjects](https://help.aliyun.com/document_detail/2780068.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The status of the data source. Valid values:
	//
	// 	- ENABLED
	//
	// 	- DISABLED
	//
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The subtype of the data source. This parameter takes effect only when the DataSourceType parameter is set to rds.
	//
	// If the value of the DataSourceType parameter is rds, the value of this parameter can be mysql, sqlserver, or postgresql.
	//
	// example:
	//
	// mysql
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s ListDataSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourcesRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ListDataSourcesRequest) GetEnvType() *int32 {
	return s.EnvType
}

func (s *ListDataSourcesRequest) GetName() *string {
	return s.Name
}

func (s *ListDataSourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataSourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataSourcesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataSourcesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListDataSourcesRequest) GetSubType() *string {
	return s.SubType
}

func (s *ListDataSourcesRequest) SetDataSourceType(v string) *ListDataSourcesRequest {
	s.DataSourceType = &v
	return s
}

func (s *ListDataSourcesRequest) SetEnvType(v int32) *ListDataSourcesRequest {
	s.EnvType = &v
	return s
}

func (s *ListDataSourcesRequest) SetName(v string) *ListDataSourcesRequest {
	s.Name = &v
	return s
}

func (s *ListDataSourcesRequest) SetPageNumber(v int32) *ListDataSourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataSourcesRequest) SetPageSize(v int32) *ListDataSourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataSourcesRequest) SetProjectId(v int64) *ListDataSourcesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataSourcesRequest) SetStatus(v string) *ListDataSourcesRequest {
	s.Status = &v
	return s
}

func (s *ListDataSourcesRequest) SetSubType(v string) *ListDataSourcesRequest {
	s.SubType = &v
	return s
}

func (s *ListDataSourcesRequest) Validate() error {
	return dara.Validate(s)
}
