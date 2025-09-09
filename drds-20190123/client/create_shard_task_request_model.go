// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateShardTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *CreateShardTaskRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *CreateShardTaskRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *CreateShardTaskRequest
	GetRegionId() *string
	SetSourceTableName(v string) *CreateShardTaskRequest
	GetSourceTableName() *string
	SetTargetTableName(v string) *CreateShardTaskRequest
	GetTargetTableName() *string
	SetTaskType(v string) *CreateShardTaskRequest
	GetTaskType() *string
}

type CreateShardTaskRequest struct {
	// The name of the DRDS database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region where the resource group resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the source table.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_tb1
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	// The name of the destination table.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_tb2
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
	// The type of the task. Valid values:`  SHARD_TO_SINGLE `,`  SINGLE_TO_SHARD `,`  SHARD_TO_SHARD `.
	//
	// This parameter is required.
	//
	// example:
	//
	// SINGLE_TO_SHARD
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateShardTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateShardTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateShardTaskRequest) GetDbName() *string {
	return s.DbName
}

func (s *CreateShardTaskRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *CreateShardTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateShardTaskRequest) GetSourceTableName() *string {
	return s.SourceTableName
}

func (s *CreateShardTaskRequest) GetTargetTableName() *string {
	return s.TargetTableName
}

func (s *CreateShardTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateShardTaskRequest) SetDbName(v string) *CreateShardTaskRequest {
	s.DbName = &v
	return s
}

func (s *CreateShardTaskRequest) SetDrdsInstanceId(v string) *CreateShardTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CreateShardTaskRequest) SetRegionId(v string) *CreateShardTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateShardTaskRequest) SetSourceTableName(v string) *CreateShardTaskRequest {
	s.SourceTableName = &v
	return s
}

func (s *CreateShardTaskRequest) SetTargetTableName(v string) *CreateShardTaskRequest {
	s.TargetTableName = &v
	return s
}

func (s *CreateShardTaskRequest) SetTaskType(v string) *CreateShardTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateShardTaskRequest) Validate() error {
	return dara.Validate(s)
}
