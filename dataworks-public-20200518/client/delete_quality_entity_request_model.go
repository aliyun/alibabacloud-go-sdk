// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityId(v int64) *DeleteQualityEntityRequest
	GetEntityId() *int64
	SetEnvType(v string) *DeleteQualityEntityRequest
	GetEnvType() *string
	SetProjectId(v int64) *DeleteQualityEntityRequest
	GetProjectId() *int64
	SetProjectName(v string) *DeleteQualityEntityRequest
	GetProjectName() *string
}

type DeleteQualityEntityRequest struct {
	// The ID of the partition filter expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The type of the compute engine or data source. The following types are supported: E-MapReduce (EMR), Hologres, AnalyticDB for PostgreSQL, CDH, MaxCompute, Kafka, and DataHub.
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
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the compute engine or data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// autotest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DeleteQualityEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityEntityRequest) GoString() string {
	return s.String()
}

func (s *DeleteQualityEntityRequest) GetEntityId() *int64 {
	return s.EntityId
}

func (s *DeleteQualityEntityRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *DeleteQualityEntityRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteQualityEntityRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeleteQualityEntityRequest) SetEntityId(v int64) *DeleteQualityEntityRequest {
	s.EntityId = &v
	return s
}

func (s *DeleteQualityEntityRequest) SetEnvType(v string) *DeleteQualityEntityRequest {
	s.EnvType = &v
	return s
}

func (s *DeleteQualityEntityRequest) SetProjectId(v int64) *DeleteQualityEntityRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteQualityEntityRequest) SetProjectName(v string) *DeleteQualityEntityRequest {
	s.ProjectName = &v
	return s
}

func (s *DeleteQualityEntityRequest) Validate() error {
	return dara.Validate(s)
}
