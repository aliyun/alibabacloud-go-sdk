// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateShardTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *ValidateShardTaskRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *ValidateShardTaskRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *ValidateShardTaskRequest
	GetRegionId() *string
	SetSourceTableName(v string) *ValidateShardTaskRequest
	GetSourceTableName() *string
	SetTargetTableName(v string) *ValidateShardTaskRequest
	GetTargetTableName() *string
	SetTaskType(v string) *ValidateShardTaskRequest
	GetTaskType() *string
}

type ValidateShardTaskRequest struct {
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds23ds****
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region where the PolarDB-X 1.0 instance is created.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the table or table shard on which you want to perform the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// buyer
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	// The name of the table or table shard on which you perform the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// buyer_new
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
	// The type of the task. Valid values:
	//
	// 	- **SINGLE_TO_SHARD**: converts a single table to a table shard.
	//
	// 	- **SHARD_TO_SINGLE**: converts a table shard to a single table.
	//
	// 	- **SHARD_TO_SHARD**: converts a table shard to another table shard.
	//
	// This parameter is required.
	//
	// example:
	//
	// SINGLE_TO_SHARD
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ValidateShardTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateShardTaskRequest) GoString() string {
	return s.String()
}

func (s *ValidateShardTaskRequest) GetDbName() *string {
	return s.DbName
}

func (s *ValidateShardTaskRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *ValidateShardTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ValidateShardTaskRequest) GetSourceTableName() *string {
	return s.SourceTableName
}

func (s *ValidateShardTaskRequest) GetTargetTableName() *string {
	return s.TargetTableName
}

func (s *ValidateShardTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *ValidateShardTaskRequest) SetDbName(v string) *ValidateShardTaskRequest {
	s.DbName = &v
	return s
}

func (s *ValidateShardTaskRequest) SetDrdsInstanceId(v string) *ValidateShardTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ValidateShardTaskRequest) SetRegionId(v string) *ValidateShardTaskRequest {
	s.RegionId = &v
	return s
}

func (s *ValidateShardTaskRequest) SetSourceTableName(v string) *ValidateShardTaskRequest {
	s.SourceTableName = &v
	return s
}

func (s *ValidateShardTaskRequest) SetTargetTableName(v string) *ValidateShardTaskRequest {
	s.TargetTableName = &v
	return s
}

func (s *ValidateShardTaskRequest) SetTaskType(v string) *ValidateShardTaskRequest {
	s.TaskType = &v
	return s
}

func (s *ValidateShardTaskRequest) Validate() error {
	return dara.Validate(s)
}
