// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityLevel(v int32) *CreateQualityEntityRequest
	GetEntityLevel() *int32
	SetEnvType(v string) *CreateQualityEntityRequest
	GetEnvType() *string
	SetMatchExpression(v string) *CreateQualityEntityRequest
	GetMatchExpression() *string
	SetProjectId(v int64) *CreateQualityEntityRequest
	GetProjectId() *int64
	SetProjectName(v string) *CreateQualityEntityRequest
	GetProjectName() *string
	SetTableName(v string) *CreateQualityEntityRequest
	GetTableName() *string
}

type CreateQualityEntityRequest struct {
	// This parameter is deprecated.
	//
	// The value 0 indicates that the partition filter expression is at the SQL level, and the system checks data quality after each SQL statement is executed.
	//
	// 	- 0
	//
	// 	- 1
	//
	// example:
	//
	// 0
	EntityLevel *int32 `json:"EntityLevel,omitempty" xml:"EntityLevel,omitempty"`
	// The type of the compute engine or data source.
	//
	// Valid values:
	//
	// 	- odps
	//
	// 	- emr
	//
	// 	- hadoop
	//
	// 	- cdh
	//
	// 	- analyticdb_for_mysql
	//
	// 	- hybriddb_for_postgresql
	//
	// 	- holodb
	//
	// This parameter is required.
	//
	// example:
	//
	// ODPS
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The partition filter expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// dt=$[yyyymmdd]
	MatchExpression *string `json:"MatchExpression,omitempty" xml:"MatchExpression,omitempty"`
	// The DataWorks workspace ID. You can log on to the DataWorks console to query the ID.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the compute engine or data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// autotest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The name of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// dual
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s CreateQualityEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityEntityRequest) GoString() string {
	return s.String()
}

func (s *CreateQualityEntityRequest) GetEntityLevel() *int32 {
	return s.EntityLevel
}

func (s *CreateQualityEntityRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *CreateQualityEntityRequest) GetMatchExpression() *string {
	return s.MatchExpression
}

func (s *CreateQualityEntityRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateQualityEntityRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateQualityEntityRequest) GetTableName() *string {
	return s.TableName
}

func (s *CreateQualityEntityRequest) SetEntityLevel(v int32) *CreateQualityEntityRequest {
	s.EntityLevel = &v
	return s
}

func (s *CreateQualityEntityRequest) SetEnvType(v string) *CreateQualityEntityRequest {
	s.EnvType = &v
	return s
}

func (s *CreateQualityEntityRequest) SetMatchExpression(v string) *CreateQualityEntityRequest {
	s.MatchExpression = &v
	return s
}

func (s *CreateQualityEntityRequest) SetProjectId(v int64) *CreateQualityEntityRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateQualityEntityRequest) SetProjectName(v string) *CreateQualityEntityRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateQualityEntityRequest) SetTableName(v string) *CreateQualityEntityRequest {
	s.TableName = &v
	return s
}

func (s *CreateQualityEntityRequest) Validate() error {
	return dara.Validate(s)
}
