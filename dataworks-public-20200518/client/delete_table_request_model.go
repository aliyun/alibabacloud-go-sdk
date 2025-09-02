// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppGuid(v string) *DeleteTableRequest
	GetAppGuid() *string
	SetEnvType(v int32) *DeleteTableRequest
	GetEnvType() *int32
	SetProjectId(v int64) *DeleteTableRequest
	GetProjectId() *int64
	SetSchema(v string) *DeleteTableRequest
	GetSchema() *string
	SetTableName(v string) *DeleteTableRequest
	GetTableName() *string
}

type DeleteTableRequest struct {
	// The globally unique identifier (GUID) of the MaxCompute project. Specify the GUID in the odps.{projectName} format.
	//
	// example:
	//
	// odps.test
	AppGuid *string `json:"AppGuid,omitempty" xml:"AppGuid,omitempty"`
	// The type of the compute engine or data source. Valid values:
	//
	// 	- cdh
	//
	// 	- analyticdb_for_mysql
	//
	// 	- odps
	//
	// 	- emr
	//
	// 	- hadoop
	//
	// 	- holodb
	//
	// 	- hybriddb_for_postgresql
	//
	// example:
	//
	// 1
	EnvType *int32 `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The ID of the DataWorks workspace.
	//
	// example:
	//
	// 101
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The schema information of the table. You need to enter the schema information of the table if you enable the table schema in MaxCompute.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// default
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The name of the MaxCompute table.
	//
	// This parameter is required.
	//
	// example:
	//
	// table1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DeleteTableRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTableRequest) GoString() string {
	return s.String()
}

func (s *DeleteTableRequest) GetAppGuid() *string {
	return s.AppGuid
}

func (s *DeleteTableRequest) GetEnvType() *int32 {
	return s.EnvType
}

func (s *DeleteTableRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteTableRequest) GetSchema() *string {
	return s.Schema
}

func (s *DeleteTableRequest) GetTableName() *string {
	return s.TableName
}

func (s *DeleteTableRequest) SetAppGuid(v string) *DeleteTableRequest {
	s.AppGuid = &v
	return s
}

func (s *DeleteTableRequest) SetEnvType(v int32) *DeleteTableRequest {
	s.EnvType = &v
	return s
}

func (s *DeleteTableRequest) SetProjectId(v int64) *DeleteTableRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteTableRequest) SetSchema(v string) *DeleteTableRequest {
	s.Schema = &v
	return s
}

func (s *DeleteTableRequest) SetTableName(v string) *DeleteTableRequest {
	s.TableName = &v
	return s
}

func (s *DeleteTableRequest) Validate() error {
	return dara.Validate(s)
}
