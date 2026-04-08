// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrafficControlTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTrafficControlTasksResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListTrafficControlTasksResponseBody
	GetTotalCount() *string
	SetTrafficControlTasks(v []*ListTrafficControlTasksResponseBodyTrafficControlTasks) *ListTrafficControlTasksResponseBody
	GetTrafficControlTasks() []*ListTrafficControlTasksResponseBodyTrafficControlTasks
}

type ListTrafficControlTasksResponseBody struct {
	// example:
	//
	// 3AAA45F6-0798-5461-9360-81D133823CE7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalCount          *string                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TrafficControlTasks []*ListTrafficControlTasksResponseBodyTrafficControlTasks `json:"TrafficControlTasks,omitempty" xml:"TrafficControlTasks,omitempty" type:"Repeated"`
}

func (s ListTrafficControlTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficControlTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrafficControlTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTrafficControlTasksResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListTrafficControlTasksResponseBody) GetTrafficControlTasks() []*ListTrafficControlTasksResponseBodyTrafficControlTasks {
	return s.TrafficControlTasks
}

func (s *ListTrafficControlTasksResponseBody) SetRequestId(v string) *ListTrafficControlTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrafficControlTasksResponseBody) SetTotalCount(v string) *ListTrafficControlTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTrafficControlTasksResponseBody) SetTrafficControlTasks(v []*ListTrafficControlTasksResponseBodyTrafficControlTasks) *ListTrafficControlTasksResponseBody {
	s.TrafficControlTasks = v
	return s
}

func (s *ListTrafficControlTasksResponseBody) Validate() error {
	if s.TrafficControlTasks != nil {
		for _, item := range s.TrafficControlTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTrafficControlTasksResponseBodyTrafficControlTasks struct {
	// example:
	//
	// 1
	BehaviorTableMetaId *string `json:"BehaviorTableMetaId,omitempty" xml:"BehaviorTableMetaId,omitempty"`
	// example:
	//
	// Global
	ControlGranularity *string `json:"ControlGranularity,omitempty" xml:"ControlGranularity,omitempty"`
	// example:
	//
	// Guaranteed
	ControlLogic *string `json:"ControlLogic,omitempty" xml:"ControlLogic,omitempty"`
	// example:
	//
	// Percent
	ControlType *string `json:"ControlType,omitempty" xml:"ControlType,omitempty"`
	// example:
	//
	// this is a test task
	Description            *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	EffectiveSceneIds      []*int32  `json:"EffectiveSceneIds,omitempty" xml:"EffectiveSceneIds,omitempty" type:"Repeated"`
	EffectiveSceneNameList []*string `json:"EffectiveSceneNameList,omitempty" xml:"EffectiveSceneNameList,omitempty" type:"Repeated"`
	EffectiveSceneNames    []*int32  `json:"EffectiveSceneNames,omitempty" xml:"EffectiveSceneNames,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-04-25
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// false
	EverPublished *bool `json:"EverPublished,omitempty" xml:"EverPublished,omitempty"`
	// example:
	//
	// Permanent
	ExecutionTime *string `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	// example:
	//
	// resource-xxx
	FlinkResourceId *string `json:"FlinkResourceId,omitempty" xml:"FlinkResourceId,omitempty"`
	// example:
	//
	// flink_test
	FlinkResourceName *string `json:"FlinkResourceName,omitempty" xml:"FlinkResourceName,omitempty"`
	// example:
	//
	// 2024-03-25T06:04:28.441Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2024-03-25T06:04:28.441Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// [{"field":"status","option":"=","value":"1"}]
	ItemConditionArray *string `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	// example:
	//
	// status=1
	ItemConditionExpress *string `json:"ItemConditionExpress,omitempty" xml:"ItemConditionExpress,omitempty"`
	// example:
	//
	// Array
	ItemConditionType *string `json:"ItemConditionType,omitempty" xml:"ItemConditionType,omitempty"`
	// example:
	//
	// 3
	ItemTableMetaId *string `json:"ItemTableMetaId,omitempty" xml:"ItemTableMetaId,omitempty"`
	// example:
	//
	// task_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1,2
	PreExperimentIds *string `json:"PreExperimentIds,omitempty" xml:"PreExperimentIds,omitempty"`
	// example:
	//
	// Runnning
	PrepubStatus *string `json:"PrepubStatus,omitempty" xml:"PrepubStatus,omitempty"`
	// example:
	//
	// 6,7
	ProdExperimentIds *string `json:"ProdExperimentIds,omitempty" xml:"ProdExperimentIds,omitempty"`
	// example:
	//
	// Runnning
	ProductStatus *string `json:"ProductStatus,omitempty" xml:"ProductStatus,omitempty"`
	// example:
	//
	// 4
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// home_feed
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// 1
	ServiceId     *string   `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceIdList []*int32  `json:"ServiceIdList,omitempty" xml:"ServiceIdList,omitempty" type:"Repeated"`
	ServiceIds    []*string `json:"ServiceIds,omitempty" xml:"ServiceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-03-25
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// event=click
	StatisBahaviorConditionExpress *string `json:"StatisBahaviorConditionExpress,omitempty" xml:"StatisBahaviorConditionExpress,omitempty"`
	// example:
	//
	// [{"field":"click","option":"<=","value":"30"}]
	StatisBehaviorConditionArray *string `json:"StatisBehaviorConditionArray,omitempty" xml:"StatisBehaviorConditionArray,omitempty"`
	// example:
	//
	// event=exposure
	StatisBehaviorConditionExpress *string `json:"StatisBehaviorConditionExpress,omitempty" xml:"StatisBehaviorConditionExpress,omitempty"`
	// example:
	//
	// Array
	StatisBehaviorConditionType *string                                                                        `json:"StatisBehaviorConditionType,omitempty" xml:"StatisBehaviorConditionType,omitempty"`
	TrafficControlTargets       []*ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets `json:"TrafficControlTargets,omitempty" xml:"TrafficControlTargets,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TrafficControlTaskId *string `json:"TrafficControlTaskId,omitempty" xml:"TrafficControlTaskId,omitempty"`
	// example:
	//
	// [{"field":"gender","option":"=","value":"male"}]
	UserConditionArray *string `json:"UserConditionArray,omitempty" xml:"UserConditionArray,omitempty"`
	// example:
	//
	// age<=30&&(3<=level<=6)&&gender=male
	UserConditionExpress *string `json:"UserConditionExpress,omitempty" xml:"UserConditionExpress,omitempty"`
	// example:
	//
	// Array
	UserConditionType *string `json:"UserConditionType,omitempty" xml:"UserConditionType,omitempty"`
	// example:
	//
	// 2
	UserTableMetaId *string `json:"UserTableMetaId,omitempty" xml:"UserTableMetaId,omitempty"`
}

func (s ListTrafficControlTasksResponseBodyTrafficControlTasks) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficControlTasksResponseBodyTrafficControlTasks) GoString() string {
	return s.String()
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetBehaviorTableMetaId() *string {
	return s.BehaviorTableMetaId
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetControlGranularity() *string {
	return s.ControlGranularity
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetControlLogic() *string {
	return s.ControlLogic
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetControlType() *string {
	return s.ControlType
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetDescription() *string {
	return s.Description
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetEffectiveSceneIds() []*int32 {
	return s.EffectiveSceneIds
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetEffectiveSceneNameList() []*string {
	return s.EffectiveSceneNameList
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetEffectiveSceneNames() []*int32 {
	return s.EffectiveSceneNames
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetEverPublished() *bool {
	return s.EverPublished
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetExecutionTime() *string {
	return s.ExecutionTime
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetFlinkResourceId() *string {
	return s.FlinkResourceId
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetFlinkResourceName() *string {
	return s.FlinkResourceName
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetItemConditionArray() *string {
	return s.ItemConditionArray
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetItemConditionExpress() *string {
	return s.ItemConditionExpress
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetItemConditionType() *string {
	return s.ItemConditionType
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetItemTableMetaId() *string {
	return s.ItemTableMetaId
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetName() *string {
	return s.Name
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetPreExperimentIds() *string {
	return s.PreExperimentIds
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetPrepubStatus() *string {
	return s.PrepubStatus
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetProdExperimentIds() *string {
	return s.ProdExperimentIds
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetProductStatus() *string {
	return s.ProductStatus
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetSceneId() *string {
	return s.SceneId
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetSceneName() *string {
	return s.SceneName
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetServiceIdList() []*int32 {
	return s.ServiceIdList
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetServiceIds() []*string {
	return s.ServiceIds
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetStatisBahaviorConditionExpress() *string {
	return s.StatisBahaviorConditionExpress
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetStatisBehaviorConditionArray() *string {
	return s.StatisBehaviorConditionArray
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetStatisBehaviorConditionExpress() *string {
	return s.StatisBehaviorConditionExpress
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetStatisBehaviorConditionType() *string {
	return s.StatisBehaviorConditionType
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetTrafficControlTargets() []*ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	return s.TrafficControlTargets
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetTrafficControlTaskId() *string {
	return s.TrafficControlTaskId
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetUserConditionArray() *string {
	return s.UserConditionArray
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetUserConditionExpress() *string {
	return s.UserConditionExpress
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetUserConditionType() *string {
	return s.UserConditionType
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) GetUserTableMetaId() *string {
	return s.UserTableMetaId
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetBehaviorTableMetaId(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.BehaviorTableMetaId = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetControlGranularity(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ControlGranularity = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetControlLogic(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ControlLogic = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetControlType(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ControlType = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetDescription(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.Description = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetEffectiveSceneIds(v []*int32) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.EffectiveSceneIds = v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetEffectiveSceneNameList(v []*string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.EffectiveSceneNameList = v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetEffectiveSceneNames(v []*int32) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.EffectiveSceneNames = v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetEndTime(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.EndTime = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetEverPublished(v bool) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.EverPublished = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetExecutionTime(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ExecutionTime = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetFlinkResourceId(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.FlinkResourceId = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetFlinkResourceName(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.FlinkResourceName = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetGmtCreateTime(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.GmtCreateTime = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetGmtModifiedTime(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetItemConditionArray(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ItemConditionArray = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetItemConditionExpress(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ItemConditionExpress = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetItemConditionType(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ItemConditionType = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetItemTableMetaId(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ItemTableMetaId = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetName(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.Name = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetPreExperimentIds(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.PreExperimentIds = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetPrepubStatus(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.PrepubStatus = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetProdExperimentIds(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ProdExperimentIds = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetProductStatus(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ProductStatus = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetSceneId(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.SceneId = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetSceneName(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.SceneName = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetServiceId(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ServiceId = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetServiceIdList(v []*int32) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ServiceIdList = v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetServiceIds(v []*string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.ServiceIds = v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetStartTime(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.StartTime = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetStatisBahaviorConditionExpress(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.StatisBahaviorConditionExpress = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetStatisBehaviorConditionArray(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.StatisBehaviorConditionArray = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetStatisBehaviorConditionExpress(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.StatisBehaviorConditionExpress = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetStatisBehaviorConditionType(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.StatisBehaviorConditionType = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetTrafficControlTargets(v []*ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.TrafficControlTargets = v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetTrafficControlTaskId(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.TrafficControlTaskId = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetUserConditionArray(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.UserConditionArray = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetUserConditionExpress(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.UserConditionExpress = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetUserConditionType(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.UserConditionType = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) SetUserTableMetaId(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasks {
	s.UserTableMetaId = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasks) Validate() error {
	if s.TrafficControlTargets != nil {
		for _, item := range s.TrafficControlTargets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets struct {
	// example:
	//
	// 2024-04-25
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// click
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 2025-06-03T15:08:10.58+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2025-11-21T16:26:58.57+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// [{"field":"status","option":"=","value":"1"}]
	ItemConditionArray *string `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	// example:
	//
	// status=1
	ItemConditionExpress *string `json:"ItemConditionExpress,omitempty" xml:"ItemConditionExpress,omitempty"`
	// example:
	//
	// Array
	ItemConditionType *string `json:"ItemConditionType,omitempty" xml:"ItemConditionType,omitempty"`
	// example:
	//
	// target-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// false
	NewProductRegulation *bool `json:"NewProductRegulation,omitempty" xml:"NewProductRegulation,omitempty"`
	// example:
	//
	// recall-1
	RecallName *string                                                                                `json:"RecallName,omitempty" xml:"RecallName,omitempty"`
	SplitParts *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts `json:"SplitParts,omitempty" xml:"SplitParts,omitempty" type:"Struct"`
	// example:
	//
	// 2024-03-25
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Daily
	StatisPeriod *string `json:"StatisPeriod,omitempty" xml:"StatisPeriod,omitempty"`
	// example:
	//
	// Opened
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 10
	ToleranceValue *int64 `json:"ToleranceValue,omitempty" xml:"ToleranceValue,omitempty"`
	// example:
	//
	// 1
	TrafficControlTargetId *string `json:"TrafficControlTargetId,omitempty" xml:"TrafficControlTargetId,omitempty"`
	// example:
	//
	// 1
	TrafficControlTaskId *string `json:"TrafficControlTaskId,omitempty" xml:"TrafficControlTaskId,omitempty"`
	// example:
	//
	// 30
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) GoString() string {
	return s.String()
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) GetEvent() *string {
	return s.Event
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) GetItemConditionArray() *string {
	return s.ItemConditionArray
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) GetItemConditionExpress() *string {
	return s.ItemConditionExpress
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) GetItemConditionType() *string {
	return s.ItemConditionType
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) GetName() *string {
	return s.Name
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) GetNewProductRegulation() *bool {
	return s.NewProductRegulation
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) GetRecallName() *string {
	return s.RecallName
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) GetSplitParts() *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts {
	return s.SplitParts
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) GetStatisPeriod() *string {
	return s.StatisPeriod
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) GetStatus() *string {
	return s.Status
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) GetToleranceValue() *int64 {
	return s.ToleranceValue
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) GetTrafficControlTargetId() *string {
	return s.TrafficControlTargetId
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) GetTrafficControlTaskId() *string {
	return s.TrafficControlTaskId
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) GetValue() *float32 {
	return s.Value
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetEndTime(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.EndTime = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetEvent(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.Event = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetGmtCreateTime(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.GmtCreateTime = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetGmtModifiedTime(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetItemConditionArray(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.ItemConditionArray = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetItemConditionExpress(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.ItemConditionExpress = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetItemConditionType(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.ItemConditionType = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetName(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.Name = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetNewProductRegulation(v bool) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.NewProductRegulation = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetRecallName(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.RecallName = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetSplitParts(v *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.SplitParts = v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetStartTime(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.StartTime = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetStatisPeriod(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.StatisPeriod = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetStatus(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.Status = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetToleranceValue(v int64) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.ToleranceValue = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetTrafficControlTargetId(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.TrafficControlTargetId = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetTrafficControlTaskId(v string) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.TrafficControlTaskId = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) SetValue(v float32) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets {
	s.Value = &v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) Validate() error {
	if s.SplitParts != nil {
		if err := s.SplitParts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts struct {
	SetValues  []*int64 `json:"SetValues,omitempty" xml:"SetValues,omitempty" type:"Repeated"`
	TimePoints []*int64 `json:"TimePoints,omitempty" xml:"TimePoints,omitempty" type:"Repeated"`
}

func (s ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts) GoString() string {
	return s.String()
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts) GetSetValues() []*int64 {
	return s.SetValues
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts) GetTimePoints() []*int64 {
	return s.TimePoints
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts) SetSetValues(v []*int64) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts {
	s.SetValues = v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts) SetTimePoints(v []*int64) *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts {
	s.TimePoints = v
	return s
}

func (s *ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts) Validate() error {
	return dara.Validate(s)
}
