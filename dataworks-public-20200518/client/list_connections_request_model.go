// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionType(v string) *ListConnectionsRequest
	GetConnectionType() *string
	SetEnvType(v int32) *ListConnectionsRequest
	GetEnvType() *int32
	SetName(v string) *ListConnectionsRequest
	GetName() *string
	SetPageNumber(v int32) *ListConnectionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListConnectionsRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListConnectionsRequest
	GetProjectId() *int64
	SetStatus(v string) *ListConnectionsRequest
	GetStatus() *string
	SetSubType(v string) *ListConnectionsRequest
	GetSubType() *string
}

type ListConnectionsRequest struct {
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
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// The environment in which the data source is used. Valid values: 0 and 1. The value 0 indicates the development environment. The value 1 indicates the production environment.
	//
	// example:
	//
	// 1
	EnvType *int32 `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The name of the data source that you want to query.
	//
	// example:
	//
	// abc
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
	// The ID of the workspace to which the data source belongs. You can call the [ListProjects](https://help.aliyun.com/document_detail/178393.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 76086
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The status of the data source. Valid values:
	//
	// 	- ENABLED
	//
	// 	- DISABLED
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The subtype of the data source. This parameter is used in scenarios where a type includes subtypes. The following type and subtypes are supported:
	//
	// 	- Type: `rds`
	//
	// 	- Subtypes: `mysql`, `sqlserver`, and `postgresql`
	//
	// example:
	//
	// mysql
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s ListConnectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsRequest) GoString() string {
	return s.String()
}

func (s *ListConnectionsRequest) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *ListConnectionsRequest) GetEnvType() *int32 {
	return s.EnvType
}

func (s *ListConnectionsRequest) GetName() *string {
	return s.Name
}

func (s *ListConnectionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConnectionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConnectionsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListConnectionsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListConnectionsRequest) GetSubType() *string {
	return s.SubType
}

func (s *ListConnectionsRequest) SetConnectionType(v string) *ListConnectionsRequest {
	s.ConnectionType = &v
	return s
}

func (s *ListConnectionsRequest) SetEnvType(v int32) *ListConnectionsRequest {
	s.EnvType = &v
	return s
}

func (s *ListConnectionsRequest) SetName(v string) *ListConnectionsRequest {
	s.Name = &v
	return s
}

func (s *ListConnectionsRequest) SetPageNumber(v int32) *ListConnectionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConnectionsRequest) SetPageSize(v int32) *ListConnectionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListConnectionsRequest) SetProjectId(v int64) *ListConnectionsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListConnectionsRequest) SetStatus(v string) *ListConnectionsRequest {
	s.Status = &v
	return s
}

func (s *ListConnectionsRequest) SetSubType(v string) *ListConnectionsRequest {
	s.SubType = &v
	return s
}

func (s *ListConnectionsRequest) Validate() error {
	return dara.Validate(s)
}
