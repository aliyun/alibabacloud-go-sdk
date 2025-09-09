// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSqlStatementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodeContent(v string) *CreateSqlStatementRequest
	GetCodeContent() *string
	SetDefaultCatalog(v string) *CreateSqlStatementRequest
	GetDefaultCatalog() *string
	SetDefaultDatabase(v string) *CreateSqlStatementRequest
	GetDefaultDatabase() *string
	SetLimit(v int32) *CreateSqlStatementRequest
	GetLimit() *int32
	SetSqlComputeId(v string) *CreateSqlStatementRequest
	GetSqlComputeId() *string
	SetTaskBizId(v string) *CreateSqlStatementRequest
	GetTaskBizId() *string
	SetRegionId(v string) *CreateSqlStatementRequest
	GetRegionId() *string
}

type CreateSqlStatementRequest struct {
	// The SQL code. You can specify one or more SQL statements.
	//
	// example:
	//
	// SHOW TABLES
	CodeContent *string `json:"codeContent,omitempty" xml:"codeContent,omitempty"`
	// The default Data Lake Formation (DLF) catalog ID.
	//
	// example:
	//
	// default_catalog
	DefaultCatalog *string `json:"defaultCatalog,omitempty" xml:"defaultCatalog,omitempty"`
	// The name of the default database.
	//
	// example:
	//
	// default
	DefaultDatabase *string `json:"defaultDatabase,omitempty" xml:"defaultDatabase,omitempty"`
	// The maximum number of entries to return. Valid values: 1 to 10000.
	//
	// example:
	//
	// 1000
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The SQL session ID. You can create an SQL session in the workspace created in EMR Serverless Spark.
	//
	// example:
	//
	// sc-dfahdfjafhajd****
	SqlComputeId *string `json:"sqlComputeId,omitempty" xml:"sqlComputeId,omitempty"`
	TaskBizId    *string `json:"taskBizId,omitempty" xml:"taskBizId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateSqlStatementRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlStatementRequest) GoString() string {
	return s.String()
}

func (s *CreateSqlStatementRequest) GetCodeContent() *string {
	return s.CodeContent
}

func (s *CreateSqlStatementRequest) GetDefaultCatalog() *string {
	return s.DefaultCatalog
}

func (s *CreateSqlStatementRequest) GetDefaultDatabase() *string {
	return s.DefaultDatabase
}

func (s *CreateSqlStatementRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *CreateSqlStatementRequest) GetSqlComputeId() *string {
	return s.SqlComputeId
}

func (s *CreateSqlStatementRequest) GetTaskBizId() *string {
	return s.TaskBizId
}

func (s *CreateSqlStatementRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSqlStatementRequest) SetCodeContent(v string) *CreateSqlStatementRequest {
	s.CodeContent = &v
	return s
}

func (s *CreateSqlStatementRequest) SetDefaultCatalog(v string) *CreateSqlStatementRequest {
	s.DefaultCatalog = &v
	return s
}

func (s *CreateSqlStatementRequest) SetDefaultDatabase(v string) *CreateSqlStatementRequest {
	s.DefaultDatabase = &v
	return s
}

func (s *CreateSqlStatementRequest) SetLimit(v int32) *CreateSqlStatementRequest {
	s.Limit = &v
	return s
}

func (s *CreateSqlStatementRequest) SetSqlComputeId(v string) *CreateSqlStatementRequest {
	s.SqlComputeId = &v
	return s
}

func (s *CreateSqlStatementRequest) SetTaskBizId(v string) *CreateSqlStatementRequest {
	s.TaskBizId = &v
	return s
}

func (s *CreateSqlStatementRequest) SetRegionId(v string) *CreateSqlStatementRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSqlStatementRequest) Validate() error {
	return dara.Validate(s)
}
