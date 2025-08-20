// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsCopyWorkloadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateApsCopyWorkloadRequest
	GetDBClusterId() *string
	SetDatasourceId(v int32) *CreateApsCopyWorkloadRequest
	GetDatasourceId() *int32
	SetDbName(v string) *CreateApsCopyWorkloadRequest
	GetDbName() *string
	SetRegionId(v string) *CreateApsCopyWorkloadRequest
	GetRegionId() *string
	SetTableName(v string) *CreateApsCopyWorkloadRequest
	GetTableName() *string
	SetWorkloadId(v string) *CreateApsCopyWorkloadRequest
	GetWorkloadId() *string
	SetWorkloadType(v string) *CreateApsCopyWorkloadRequest
	GetWorkloadType() *string
}

type CreateApsCopyWorkloadRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1xxxxxxxx47
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 123
	DatasourceId *int32 `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// dbName
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aps-******
	WorkloadId *string `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// The type of the job.
	//
	// This parameter is required.
	//
	// example:
	//
	// SLS_INGESTION_ADB
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
}

func (s CreateApsCopyWorkloadRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApsCopyWorkloadRequest) GoString() string {
	return s.String()
}

func (s *CreateApsCopyWorkloadRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateApsCopyWorkloadRequest) GetDatasourceId() *int32 {
	return s.DatasourceId
}

func (s *CreateApsCopyWorkloadRequest) GetDbName() *string {
	return s.DbName
}

func (s *CreateApsCopyWorkloadRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApsCopyWorkloadRequest) GetTableName() *string {
	return s.TableName
}

func (s *CreateApsCopyWorkloadRequest) GetWorkloadId() *string {
	return s.WorkloadId
}

func (s *CreateApsCopyWorkloadRequest) GetWorkloadType() *string {
	return s.WorkloadType
}

func (s *CreateApsCopyWorkloadRequest) SetDBClusterId(v string) *CreateApsCopyWorkloadRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateApsCopyWorkloadRequest) SetDatasourceId(v int32) *CreateApsCopyWorkloadRequest {
	s.DatasourceId = &v
	return s
}

func (s *CreateApsCopyWorkloadRequest) SetDbName(v string) *CreateApsCopyWorkloadRequest {
	s.DbName = &v
	return s
}

func (s *CreateApsCopyWorkloadRequest) SetRegionId(v string) *CreateApsCopyWorkloadRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApsCopyWorkloadRequest) SetTableName(v string) *CreateApsCopyWorkloadRequest {
	s.TableName = &v
	return s
}

func (s *CreateApsCopyWorkloadRequest) SetWorkloadId(v string) *CreateApsCopyWorkloadRequest {
	s.WorkloadId = &v
	return s
}

func (s *CreateApsCopyWorkloadRequest) SetWorkloadType(v string) *CreateApsCopyWorkloadRequest {
	s.WorkloadType = &v
	return s
}

func (s *CreateApsCopyWorkloadRequest) Validate() error {
	return dara.Validate(s)
}
