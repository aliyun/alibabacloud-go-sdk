// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *DescribeAIDBClusterTasksResponseBody
	GetEngine() *string
	SetEngineVersion(v string) *DescribeAIDBClusterTasksResponseBody
	GetEngineVersion() *string
	SetItems(v []*DescribeAIDBClusterTasksResponseBodyItems) *DescribeAIDBClusterTasksResponseBody
	GetItems() []*DescribeAIDBClusterTasksResponseBodyItems
	SetRelativeDBClusterId(v string) *DescribeAIDBClusterTasksResponseBody
	GetRelativeDBClusterId() *string
	SetRequestId(v string) *DescribeAIDBClusterTasksResponseBody
	GetRequestId() *string
	SetTaskType(v string) *DescribeAIDBClusterTasksResponseBody
	GetTaskType() *string
}

type DescribeAIDBClusterTasksResponseBody struct {
	// The cluster engine.
	//
	// example:
	//
	// polardb_ai
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version.
	//
	// example:
	//
	// 3.1
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The cluster endpoint details.
	Items []*DescribeAIDBClusterTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The ID of the PolarDB cluster.
	//
	// example:
	//
	// pc-2ze***
	RelativeDBClusterId *string `json:"RelativeDBClusterId,omitempty" xml:"RelativeDBClusterId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 925B84D9-CA72-432C-95CF-738C22******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task type.
	//
	// example:
	//
	// train
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeAIDBClusterTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterTasksResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeAIDBClusterTasksResponseBody) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeAIDBClusterTasksResponseBody) GetItems() []*DescribeAIDBClusterTasksResponseBodyItems {
	return s.Items
}

func (s *DescribeAIDBClusterTasksResponseBody) GetRelativeDBClusterId() *string {
	return s.RelativeDBClusterId
}

func (s *DescribeAIDBClusterTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAIDBClusterTasksResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeAIDBClusterTasksResponseBody) SetEngine(v string) *DescribeAIDBClusterTasksResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBody) SetEngineVersion(v string) *DescribeAIDBClusterTasksResponseBody {
	s.EngineVersion = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBody) SetItems(v []*DescribeAIDBClusterTasksResponseBodyItems) *DescribeAIDBClusterTasksResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBody) SetRelativeDBClusterId(v string) *DescribeAIDBClusterTasksResponseBody {
	s.RelativeDBClusterId = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBody) SetRequestId(v string) *DescribeAIDBClusterTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBody) SetTaskType(v string) *DescribeAIDBClusterTasksResponseBody {
	s.TaskType = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAIDBClusterTasksResponseBodyItems struct {
	// The task completion time.
	//
	// example:
	//
	// 2020-06-09T18:00:00Z
	CompletedTime *string `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-03-25T09:37:10Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The node description.
	//
	// example:
	//
	// test
	DBNodeDescription *string `json:"DBNodeDescription,omitempty" xml:"DBNodeDescription,omitempty"`
	// The instance ID of the model operator.
	//
	// example:
	//
	// pm-2ze***
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The instance status. This parameter may not be returned.
	//
	// example:
	//
	// RUNNING
	DBNodeStatus *string `json:"DBNodeStatus,omitempty" xml:"DBNodeStatus,omitempty"`
	// The instance status.
	//
	// example:
	//
	// RUNNING
	DBNodeStatusDesc *string `json:"DBNodeStatusDesc,omitempty" xml:"DBNodeStatusDesc,omitempty"`
	// The zone.
	//
	// example:
	//
	// cn-beijing-i
	DataZoneId *string `json:"DataZoneId,omitempty" xml:"DataZoneId,omitempty"`
	// The cluster engine.
	//
	// example:
	//
	// polardb_ai
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version.
	//
	// example:
	//
	// 3.1
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The model name.
	//
	// example:
	//
	// ***
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The path.
	//
	// example:
	//
	// ***
	ModelPath *string `json:"ModelPath,omitempty" xml:"ModelPath,omitempty"`
	// The model source.
	//
	// example:
	//
	// public
	ModelSource *string `json:"ModelSource,omitempty" xml:"ModelSource,omitempty"`
	// The running parameters.
	//
	// example:
	//
	// xxx
	RunningTimes *string `json:"RunningTimes,omitempty" xml:"RunningTimes,omitempty"`
	// The task start time.
	//
	// example:
	//
	// 2020-06-09T18:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The mode.
	//
	// example:
	//
	// sft
	TrainMode *string `json:"TrainMode,omitempty" xml:"TrainMode,omitempty"`
	// The type.
	//
	// example:
	//
	// lora
	TrainType *string `json:"TrainType,omitempty" xml:"TrainType,omitempty"`
	TuneArch  *string `json:"TuneArch,omitempty" xml:"TuneArch,omitempty"`
}

func (s DescribeAIDBClusterTasksResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) GetCompletedTime() *string {
	return s.CompletedTime
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) GetDBNodeDescription() *string {
	return s.DBNodeDescription
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) GetDBNodeStatus() *string {
	return s.DBNodeStatus
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) GetDBNodeStatusDesc() *string {
	return s.DBNodeStatusDesc
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) GetDataZoneId() *string {
	return s.DataZoneId
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) GetEngine() *string {
	return s.Engine
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) GetModelName() *string {
	return s.ModelName
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) GetModelPath() *string {
	return s.ModelPath
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) GetModelSource() *string {
	return s.ModelSource
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) GetRunningTimes() *string {
	return s.RunningTimes
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) GetTrainMode() *string {
	return s.TrainMode
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) GetTrainType() *string {
	return s.TrainType
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) GetTuneArch() *string {
	return s.TuneArch
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) SetCompletedTime(v string) *DescribeAIDBClusterTasksResponseBodyItems {
	s.CompletedTime = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) SetCreationTime(v string) *DescribeAIDBClusterTasksResponseBodyItems {
	s.CreationTime = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) SetDBNodeDescription(v string) *DescribeAIDBClusterTasksResponseBodyItems {
	s.DBNodeDescription = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) SetDBNodeId(v string) *DescribeAIDBClusterTasksResponseBodyItems {
	s.DBNodeId = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) SetDBNodeStatus(v string) *DescribeAIDBClusterTasksResponseBodyItems {
	s.DBNodeStatus = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) SetDBNodeStatusDesc(v string) *DescribeAIDBClusterTasksResponseBodyItems {
	s.DBNodeStatusDesc = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) SetDataZoneId(v string) *DescribeAIDBClusterTasksResponseBodyItems {
	s.DataZoneId = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) SetEngine(v string) *DescribeAIDBClusterTasksResponseBodyItems {
	s.Engine = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) SetEngineVersion(v string) *DescribeAIDBClusterTasksResponseBodyItems {
	s.EngineVersion = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) SetModelName(v string) *DescribeAIDBClusterTasksResponseBodyItems {
	s.ModelName = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) SetModelPath(v string) *DescribeAIDBClusterTasksResponseBodyItems {
	s.ModelPath = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) SetModelSource(v string) *DescribeAIDBClusterTasksResponseBodyItems {
	s.ModelSource = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) SetRunningTimes(v string) *DescribeAIDBClusterTasksResponseBodyItems {
	s.RunningTimes = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) SetStartTime(v string) *DescribeAIDBClusterTasksResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) SetTrainMode(v string) *DescribeAIDBClusterTasksResponseBodyItems {
	s.TrainMode = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) SetTrainType(v string) *DescribeAIDBClusterTasksResponseBodyItems {
	s.TrainType = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) SetTuneArch(v string) *DescribeAIDBClusterTasksResponseBodyItems {
	s.TuneArch = &v
	return s
}

func (s *DescribeAIDBClusterTasksResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
