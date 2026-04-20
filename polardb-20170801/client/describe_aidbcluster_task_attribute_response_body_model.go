// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterTaskAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessInfo(v string) *DescribeAIDBClusterTaskAttributeResponseBody
	GetAccessInfo() *string
	SetClusterNetworkType(v string) *DescribeAIDBClusterTaskAttributeResponseBody
	GetClusterNetworkType() *string
	SetCreateTime(v string) *DescribeAIDBClusterTaskAttributeResponseBody
	GetCreateTime() *string
	SetDBClusterDescription(v string) *DescribeAIDBClusterTaskAttributeResponseBody
	GetDBClusterDescription() *string
	SetDBClusterId(v string) *DescribeAIDBClusterTaskAttributeResponseBody
	GetDBClusterId() *string
	SetDBClusterStatus(v string) *DescribeAIDBClusterTaskAttributeResponseBody
	GetDBClusterStatus() *string
	SetDBClusterStatusDesc(v string) *DescribeAIDBClusterTaskAttributeResponseBody
	GetDBClusterStatusDesc() *string
	SetDBType(v string) *DescribeAIDBClusterTaskAttributeResponseBody
	GetDBType() *string
	SetDBVersion(v string) *DescribeAIDBClusterTaskAttributeResponseBody
	GetDBVersion() *string
	SetDataSets(v []*DescribeAIDBClusterTaskAttributeResponseBodyDataSets) *DescribeAIDBClusterTaskAttributeResponseBody
	GetDataSets() []*DescribeAIDBClusterTaskAttributeResponseBodyDataSets
	SetExtraInfo(v []map[string]interface{}) *DescribeAIDBClusterTaskAttributeResponseBody
	GetExtraInfo() []map[string]interface{}
	SetKindCode(v int64) *DescribeAIDBClusterTaskAttributeResponseBody
	GetKindCode() *int64
	SetLockMode(v string) *DescribeAIDBClusterTaskAttributeResponseBody
	GetLockMode() *string
	SetMaintainEndTime(v string) *DescribeAIDBClusterTaskAttributeResponseBody
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *DescribeAIDBClusterTaskAttributeResponseBody
	GetMaintainStartTime() *string
	SetModelPath(v string) *DescribeAIDBClusterTaskAttributeResponseBody
	GetModelPath() *string
	SetRequestId(v string) *DescribeAIDBClusterTaskAttributeResponseBody
	GetRequestId() *string
	SetTaskInfo(v []*DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) *DescribeAIDBClusterTaskAttributeResponseBody
	GetTaskInfo() []*DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo
	SetVPCId(v string) *DescribeAIDBClusterTaskAttributeResponseBody
	GetVPCId() *string
	SetVSwitchId(v string) *DescribeAIDBClusterTaskAttributeResponseBody
	GetVSwitchId() *string
}

type DescribeAIDBClusterTaskAttributeResponseBody struct {
	// example:
	//
	// {\\"networkInterfaceId\\":\\"eni-2zea***\\",\\"port\\":\\"8000\\",\\"host\\":\\"192.**.**.**\\"}
	AccessInfo *string `json:"AccessInfo,omitempty" xml:"AccessInfo,omitempty"`
	// example:
	//
	// VPC
	ClusterNetworkType *string `json:"ClusterNetworkType,omitempty" xml:"ClusterNetworkType,omitempty"`
	// example:
	//
	// 2025-11-12T03:45:13Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// task01
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// example:
	//
	// pm-2ze99***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// COMPLETED
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// example:
	//
	// COMPLETED
	DBClusterStatusDesc *string `json:"DBClusterStatusDesc,omitempty" xml:"DBClusterStatusDesc,omitempty"`
	// example:
	//
	// polardb_ai
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// example:
	//
	// 3.1
	DBVersion *string                                                 `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	DataSets  []*DescribeAIDBClusterTaskAttributeResponseBodyDataSets `json:"DataSets,omitempty" xml:"DataSets,omitempty" type:"Repeated"`
	ExtraInfo []map[string]interface{}                                `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty" type:"Repeated"`
	// example:
	//
	// 18
	KindCode *int64 `json:"KindCode,omitempty" xml:"KindCode,omitempty"`
	// example:
	//
	// 1
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// example:
	//
	// 12:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// example:
	//
	// 8:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	ModelPath         *string `json:"ModelPath,omitempty" xml:"ModelPath,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 45D24263-7E3A-4140-9472-************
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskInfo  []*DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-**********
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-*********************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeAIDBClusterTaskAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterTaskAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) GetAccessInfo() *string {
	return s.AccessInfo
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) GetClusterNetworkType() *string {
	return s.ClusterNetworkType
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) GetDBClusterStatusDesc() *string {
	return s.DBClusterStatusDesc
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) GetDBType() *string {
	return s.DBType
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) GetDataSets() []*DescribeAIDBClusterTaskAttributeResponseBodyDataSets {
	return s.DataSets
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) GetExtraInfo() []map[string]interface{} {
	return s.ExtraInfo
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) GetKindCode() *int64 {
	return s.KindCode
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) GetModelPath() *string {
	return s.ModelPath
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) GetTaskInfo() []*DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo {
	return s.TaskInfo
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) SetAccessInfo(v string) *DescribeAIDBClusterTaskAttributeResponseBody {
	s.AccessInfo = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) SetClusterNetworkType(v string) *DescribeAIDBClusterTaskAttributeResponseBody {
	s.ClusterNetworkType = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) SetCreateTime(v string) *DescribeAIDBClusterTaskAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) SetDBClusterDescription(v string) *DescribeAIDBClusterTaskAttributeResponseBody {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) SetDBClusterId(v string) *DescribeAIDBClusterTaskAttributeResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) SetDBClusterStatus(v string) *DescribeAIDBClusterTaskAttributeResponseBody {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) SetDBClusterStatusDesc(v string) *DescribeAIDBClusterTaskAttributeResponseBody {
	s.DBClusterStatusDesc = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) SetDBType(v string) *DescribeAIDBClusterTaskAttributeResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) SetDBVersion(v string) *DescribeAIDBClusterTaskAttributeResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) SetDataSets(v []*DescribeAIDBClusterTaskAttributeResponseBodyDataSets) *DescribeAIDBClusterTaskAttributeResponseBody {
	s.DataSets = v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) SetExtraInfo(v []map[string]interface{}) *DescribeAIDBClusterTaskAttributeResponseBody {
	s.ExtraInfo = v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) SetKindCode(v int64) *DescribeAIDBClusterTaskAttributeResponseBody {
	s.KindCode = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) SetLockMode(v string) *DescribeAIDBClusterTaskAttributeResponseBody {
	s.LockMode = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) SetMaintainEndTime(v string) *DescribeAIDBClusterTaskAttributeResponseBody {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) SetMaintainStartTime(v string) *DescribeAIDBClusterTaskAttributeResponseBody {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) SetModelPath(v string) *DescribeAIDBClusterTaskAttributeResponseBody {
	s.ModelPath = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) SetRequestId(v string) *DescribeAIDBClusterTaskAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) SetTaskInfo(v []*DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) *DescribeAIDBClusterTaskAttributeResponseBody {
	s.TaskInfo = v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) SetVPCId(v string) *DescribeAIDBClusterTaskAttributeResponseBody {
	s.VPCId = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) SetVSwitchId(v string) *DescribeAIDBClusterTaskAttributeResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBody) Validate() error {
	if s.DataSets != nil {
		for _, item := range s.DataSets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TaskInfo != nil {
		for _, item := range s.TaskInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAIDBClusterTaskAttributeResponseBodyDataSets struct {
	// example:
	//
	// dataset02
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// example:
	//
	// polardb_ai/datasets/train/grpo/dataset02/test-**.jsonl#1000
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// 0.1
	SplitDatasetRatio *string `json:"SplitDatasetRatio,omitempty" xml:"SplitDatasetRatio,omitempty"`
	// example:
	//
	// train
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAIDBClusterTaskAttributeResponseBodyDataSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterTaskAttributeResponseBodyDataSets) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyDataSets) GetDatasetName() *string {
	return s.DatasetName
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyDataSets) GetPath() *string {
	return s.Path
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyDataSets) GetSplitDatasetRatio() *string {
	return s.SplitDatasetRatio
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyDataSets) GetType() *string {
	return s.Type
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyDataSets) SetDatasetName(v string) *DescribeAIDBClusterTaskAttributeResponseBodyDataSets {
	s.DatasetName = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyDataSets) SetPath(v string) *DescribeAIDBClusterTaskAttributeResponseBodyDataSets {
	s.Path = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyDataSets) SetSplitDatasetRatio(v string) *DescribeAIDBClusterTaskAttributeResponseBodyDataSets {
	s.SplitDatasetRatio = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyDataSets) SetType(v string) *DescribeAIDBClusterTaskAttributeResponseBodyDataSets {
	s.Type = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyDataSets) Validate() error {
	return dara.Validate(s)
}

type DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo struct {
	// example:
	//
	// 2025-09-10T01:56:00Z
	CompletedTime *string `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	// example:
	//
	// Qwen-1.7B
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// example:
	//
	// Qwen-1.7B
	ModelPath *string `json:"ModelPath,omitempty" xml:"ModelPath,omitempty"`
	// example:
	//
	// public
	ModelSource *string `json:"ModelSource,omitempty" xml:"ModelSource,omitempty"`
	// example:
	//
	// {"split_dataset_ratio": 0.1,"train_mode": "grpo","train_type": "lora","num_epochs": 1,"batch_size": 2,"eval_batch_size": 2, "num_generations": 2, "learning_rate": "1e-6", "data_file_list": "test-00000-of-00001.jsonl#1000", "lora_rank": 8, "lora_alpha": 32, "external_plugins": "/plugin/train/plugin.py", "reward_funcs": "format,external_countdown", "gpu_memory_utilization": 0.4}
	RunningTimes *string `json:"RunningTimes,omitempty" xml:"RunningTimes,omitempty"`
	// example:
	//
	// 2025-09-10T01:56:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// stf
	TrainMode *string `json:"TrainMode,omitempty" xml:"TrainMode,omitempty"`
	// example:
	//
	// lora
	TrainType *string `json:"TrainType,omitempty" xml:"TrainType,omitempty"`
}

func (s DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) GetCompletedTime() *string {
	return s.CompletedTime
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) GetModelName() *string {
	return s.ModelName
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) GetModelPath() *string {
	return s.ModelPath
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) GetModelSource() *string {
	return s.ModelSource
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) GetRunningTimes() *string {
	return s.RunningTimes
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) GetTrainMode() *string {
	return s.TrainMode
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) GetTrainType() *string {
	return s.TrainType
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) SetCompletedTime(v string) *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo {
	s.CompletedTime = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) SetModelName(v string) *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo {
	s.ModelName = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) SetModelPath(v string) *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo {
	s.ModelPath = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) SetModelSource(v string) *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo {
	s.ModelSource = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) SetRunningTimes(v string) *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo {
	s.RunningTimes = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) SetStartTime(v string) *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) SetTrainMode(v string) *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo {
	s.TrainMode = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) SetTrainType(v string) *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo {
	s.TrainType = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponseBodyTaskInfo) Validate() error {
	return dara.Validate(s)
}
