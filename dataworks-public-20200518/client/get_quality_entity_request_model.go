// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvType(v string) *GetQualityEntityRequest
	GetEnvType() *string
	SetMatchExpression(v string) *GetQualityEntityRequest
	GetMatchExpression() *string
	SetProjectId(v int64) *GetQualityEntityRequest
	GetProjectId() *int64
	SetProjectName(v string) *GetQualityEntityRequest
	GetProjectName() *string
	SetTableName(v string) *GetQualityEntityRequest
	GetTableName() *string
}

type GetQualityEntityRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// odps
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The partition filter expression.
	//
	// example:
	//
	// dt=$[yyyymmdd]
	MatchExpression *string `json:"MatchExpression,omitempty" xml:"MatchExpression,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the compute engine instance or data source. You can obtain the name from data source configurations.
	//
	// This parameter is required.
	//
	// example:
	//
	// autotest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The name of the partitioned table. You can call the [GetMetaTablePartition](https://help.aliyun.com/document_detail/173923.html) operation to obtain the name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dual
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetQualityEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityEntityRequest) GoString() string {
	return s.String()
}

func (s *GetQualityEntityRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *GetQualityEntityRequest) GetMatchExpression() *string {
	return s.MatchExpression
}

func (s *GetQualityEntityRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetQualityEntityRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetQualityEntityRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetQualityEntityRequest) SetEnvType(v string) *GetQualityEntityRequest {
	s.EnvType = &v
	return s
}

func (s *GetQualityEntityRequest) SetMatchExpression(v string) *GetQualityEntityRequest {
	s.MatchExpression = &v
	return s
}

func (s *GetQualityEntityRequest) SetProjectId(v int64) *GetQualityEntityRequest {
	s.ProjectId = &v
	return s
}

func (s *GetQualityEntityRequest) SetProjectName(v string) *GetQualityEntityRequest {
	s.ProjectName = &v
	return s
}

func (s *GetQualityEntityRequest) SetTableName(v string) *GetQualityEntityRequest {
	s.TableName = &v
	return s
}

func (s *GetQualityEntityRequest) Validate() error {
	return dara.Validate(s)
}
