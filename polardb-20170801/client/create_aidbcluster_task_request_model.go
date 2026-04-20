// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIDBClusterTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateAIDBClusterTaskRequest
	GetDBClusterId() *string
	SetDBInstanceClass(v string) *CreateAIDBClusterTaskRequest
	GetDBInstanceClass() *string
	SetDatasetPath(v string) *CreateAIDBClusterTaskRequest
	GetDatasetPath() *string
	SetEvalDatasetPath(v string) *CreateAIDBClusterTaskRequest
	GetEvalDatasetPath() *string
	SetKubeType(v string) *CreateAIDBClusterTaskRequest
	GetKubeType() *string
	SetModelName(v string) *CreateAIDBClusterTaskRequest
	GetModelName() *string
	SetModelSource(v string) *CreateAIDBClusterTaskRequest
	GetModelSource() *string
	SetModelType(v string) *CreateAIDBClusterTaskRequest
	GetModelType() *string
	SetOwnerAccount(v string) *CreateAIDBClusterTaskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateAIDBClusterTaskRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateAIDBClusterTaskRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateAIDBClusterTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAIDBClusterTaskRequest
	GetResourceOwnerId() *int64
	SetRunningParameter(v string) *CreateAIDBClusterTaskRequest
	GetRunningParameter() *string
	SetSecurityGroupId(v string) *CreateAIDBClusterTaskRequest
	GetSecurityGroupId() *string
	SetTaskName(v string) *CreateAIDBClusterTaskRequest
	GetTaskName() *string
	SetVPCId(v string) *CreateAIDBClusterTaskRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateAIDBClusterTaskRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CreateAIDBClusterTaskRequest
	GetZoneId() *string
}

type CreateAIDBClusterTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// polar.pg.g6.4xlarge.guh
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// example:
	//
	// pds-2zetrain***
	DatasetPath *string `json:"DatasetPath,omitempty" xml:"DatasetPath,omitempty"`
	// example:
	//
	// pds-2zetrain***
	EvalDatasetPath *string `json:"EvalDatasetPath,omitempty" xml:"EvalDatasetPath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aitrain
	KubeType *string `json:"KubeType,omitempty" xml:"KubeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Qwen3-8B
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// public
	ModelSource *string `json:"ModelSource,omitempty" xml:"ModelSource,omitempty"`
	// example:
	//
	// qwen3
	ModelType    *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"split_dataset_ratio": 0.1,"train_mode": "grpo","train_type": "lora","num_epochs": 1,"batch_size": 2,"eval_batch_size": 2, "num_generations": 2, "learning_rate": "1e-6", "data_file_list": "test-00000-of-00001.jsonl#1000", "lora_rank": 8, "lora_alpha": 32, "external_plugins": "/plugin/train/plugin.py", "reward_funcs": "format,external_countdown", "gpu_memory_utilization": 0.4}
	RunningParameter *string `json:"RunningParameter,omitempty" xml:"RunningParameter,omitempty"`
	// example:
	//
	// sg-bp**************
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// xxxx
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// vpc-*******************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-*********************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateAIDBClusterTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAIDBClusterTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAIDBClusterTaskRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAIDBClusterTaskRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *CreateAIDBClusterTaskRequest) GetDatasetPath() *string {
	return s.DatasetPath
}

func (s *CreateAIDBClusterTaskRequest) GetEvalDatasetPath() *string {
	return s.EvalDatasetPath
}

func (s *CreateAIDBClusterTaskRequest) GetKubeType() *string {
	return s.KubeType
}

func (s *CreateAIDBClusterTaskRequest) GetModelName() *string {
	return s.ModelName
}

func (s *CreateAIDBClusterTaskRequest) GetModelSource() *string {
	return s.ModelSource
}

func (s *CreateAIDBClusterTaskRequest) GetModelType() *string {
	return s.ModelType
}

func (s *CreateAIDBClusterTaskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateAIDBClusterTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAIDBClusterTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAIDBClusterTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAIDBClusterTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAIDBClusterTaskRequest) GetRunningParameter() *string {
	return s.RunningParameter
}

func (s *CreateAIDBClusterTaskRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateAIDBClusterTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateAIDBClusterTaskRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateAIDBClusterTaskRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateAIDBClusterTaskRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateAIDBClusterTaskRequest) SetDBClusterId(v string) *CreateAIDBClusterTaskRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAIDBClusterTaskRequest) SetDBInstanceClass(v string) *CreateAIDBClusterTaskRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CreateAIDBClusterTaskRequest) SetDatasetPath(v string) *CreateAIDBClusterTaskRequest {
	s.DatasetPath = &v
	return s
}

func (s *CreateAIDBClusterTaskRequest) SetEvalDatasetPath(v string) *CreateAIDBClusterTaskRequest {
	s.EvalDatasetPath = &v
	return s
}

func (s *CreateAIDBClusterTaskRequest) SetKubeType(v string) *CreateAIDBClusterTaskRequest {
	s.KubeType = &v
	return s
}

func (s *CreateAIDBClusterTaskRequest) SetModelName(v string) *CreateAIDBClusterTaskRequest {
	s.ModelName = &v
	return s
}

func (s *CreateAIDBClusterTaskRequest) SetModelSource(v string) *CreateAIDBClusterTaskRequest {
	s.ModelSource = &v
	return s
}

func (s *CreateAIDBClusterTaskRequest) SetModelType(v string) *CreateAIDBClusterTaskRequest {
	s.ModelType = &v
	return s
}

func (s *CreateAIDBClusterTaskRequest) SetOwnerAccount(v string) *CreateAIDBClusterTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAIDBClusterTaskRequest) SetOwnerId(v int64) *CreateAIDBClusterTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAIDBClusterTaskRequest) SetRegionId(v string) *CreateAIDBClusterTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAIDBClusterTaskRequest) SetResourceOwnerAccount(v string) *CreateAIDBClusterTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAIDBClusterTaskRequest) SetResourceOwnerId(v int64) *CreateAIDBClusterTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAIDBClusterTaskRequest) SetRunningParameter(v string) *CreateAIDBClusterTaskRequest {
	s.RunningParameter = &v
	return s
}

func (s *CreateAIDBClusterTaskRequest) SetSecurityGroupId(v string) *CreateAIDBClusterTaskRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateAIDBClusterTaskRequest) SetTaskName(v string) *CreateAIDBClusterTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateAIDBClusterTaskRequest) SetVPCId(v string) *CreateAIDBClusterTaskRequest {
	s.VPCId = &v
	return s
}

func (s *CreateAIDBClusterTaskRequest) SetVSwitchId(v string) *CreateAIDBClusterTaskRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateAIDBClusterTaskRequest) SetZoneId(v string) *CreateAIDBClusterTaskRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateAIDBClusterTaskRequest) Validate() error {
	return dara.Validate(s)
}
