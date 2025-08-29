// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficControlTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBehaviorTableMetaId(v string) *CreateTrafficControlTaskRequest
	GetBehaviorTableMetaId() *string
	SetControlGranularity(v string) *CreateTrafficControlTaskRequest
	GetControlGranularity() *string
	SetControlLogic(v string) *CreateTrafficControlTaskRequest
	GetControlLogic() *string
	SetControlType(v string) *CreateTrafficControlTaskRequest
	GetControlType() *string
	SetDescription(v string) *CreateTrafficControlTaskRequest
	GetDescription() *string
	SetEffectiveSceneIds(v []*int32) *CreateTrafficControlTaskRequest
	GetEffectiveSceneIds() []*int32
	SetEndTime(v string) *CreateTrafficControlTaskRequest
	GetEndTime() *string
	SetExecutionTime(v string) *CreateTrafficControlTaskRequest
	GetExecutionTime() *string
	SetFlinkResourceId(v string) *CreateTrafficControlTaskRequest
	GetFlinkResourceId() *string
	SetInstanceId(v string) *CreateTrafficControlTaskRequest
	GetInstanceId() *string
	SetItemConditionArray(v string) *CreateTrafficControlTaskRequest
	GetItemConditionArray() *string
	SetItemConditionExpress(v string) *CreateTrafficControlTaskRequest
	GetItemConditionExpress() *string
	SetItemConditionType(v string) *CreateTrafficControlTaskRequest
	GetItemConditionType() *string
	SetItemTableMetaId(v string) *CreateTrafficControlTaskRequest
	GetItemTableMetaId() *string
	SetName(v string) *CreateTrafficControlTaskRequest
	GetName() *string
	SetPreExperimentIds(v string) *CreateTrafficControlTaskRequest
	GetPreExperimentIds() *string
	SetProdExperimentIds(v string) *CreateTrafficControlTaskRequest
	GetProdExperimentIds() *string
	SetSceneId(v string) *CreateTrafficControlTaskRequest
	GetSceneId() *string
	SetServiceId(v string) *CreateTrafficControlTaskRequest
	GetServiceId() *string
	SetServiceIds(v []*int32) *CreateTrafficControlTaskRequest
	GetServiceIds() []*int32
	SetStartTime(v string) *CreateTrafficControlTaskRequest
	GetStartTime() *string
	SetStatisBehaviorConditionArray(v string) *CreateTrafficControlTaskRequest
	GetStatisBehaviorConditionArray() *string
	SetStatisBehaviorConditionExpress(v string) *CreateTrafficControlTaskRequest
	GetStatisBehaviorConditionExpress() *string
	SetStatisBehaviorConditionType(v string) *CreateTrafficControlTaskRequest
	GetStatisBehaviorConditionType() *string
	SetTrafficControlTargets(v []*CreateTrafficControlTaskRequestTrafficControlTargets) *CreateTrafficControlTaskRequest
	GetTrafficControlTargets() []*CreateTrafficControlTaskRequestTrafficControlTargets
	SetUserConditionArray(v string) *CreateTrafficControlTaskRequest
	GetUserConditionArray() *string
	SetUserConditionExpress(v string) *CreateTrafficControlTaskRequest
	GetUserConditionExpress() *string
	SetUserConditionType(v string) *CreateTrafficControlTaskRequest
	GetUserConditionType() *string
	SetUserTableMetaId(v string) *CreateTrafficControlTaskRequest
	GetUserTableMetaId() *string
}

type CreateTrafficControlTaskRequest struct {
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
	Description       *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	EffectiveSceneIds []*int32 `json:"EffectiveSceneIds,omitempty" xml:"EffectiveSceneIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-03-26
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// TimeRange
	ExecutionTime   *string `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	FlinkResourceId *string `json:"FlinkResourceId,omitempty" xml:"FlinkResourceId,omitempty"`
	// example:
	//
	// pairec-cn-***test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// [{\\"field\\":\\"status\\",\\"option\\":\\"=\\",\\"value\\":\\"1\\"}]
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
	// task-1
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PreExperimentIds  *string `json:"PreExperimentIds,omitempty" xml:"PreExperimentIds,omitempty"`
	ProdExperimentIds *string `json:"ProdExperimentIds,omitempty" xml:"ProdExperimentIds,omitempty"`
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// This parameter is required.
	ServiceId  *string  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceIds []*int32 `json:"ServiceIds,omitempty" xml:"ServiceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-03-25
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// [{\\"field\\":\\"click\\",\\"option\\":\\"<=\\",\\"value\\":\\"30\\"}]
	StatisBehaviorConditionArray *string `json:"StatisBehaviorConditionArray,omitempty" xml:"StatisBehaviorConditionArray,omitempty"`
	// example:
	//
	// click=30
	StatisBehaviorConditionExpress *string `json:"StatisBehaviorConditionExpress,omitempty" xml:"StatisBehaviorConditionExpress,omitempty"`
	// example:
	//
	// Array
	StatisBehaviorConditionType *string                                                 `json:"StatisBehaviorConditionType,omitempty" xml:"StatisBehaviorConditionType,omitempty"`
	TrafficControlTargets       []*CreateTrafficControlTaskRequestTrafficControlTargets `json:"TrafficControlTargets,omitempty" xml:"TrafficControlTargets,omitempty" type:"Repeated"`
	// example:
	//
	// [{\\"field\\":\\"gender\\",\\"option\\":\\"=\\",\\"value\\":\\"male\\"}]
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

func (s CreateTrafficControlTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficControlTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTrafficControlTaskRequest) GetBehaviorTableMetaId() *string {
	return s.BehaviorTableMetaId
}

func (s *CreateTrafficControlTaskRequest) GetControlGranularity() *string {
	return s.ControlGranularity
}

func (s *CreateTrafficControlTaskRequest) GetControlLogic() *string {
	return s.ControlLogic
}

func (s *CreateTrafficControlTaskRequest) GetControlType() *string {
	return s.ControlType
}

func (s *CreateTrafficControlTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTrafficControlTaskRequest) GetEffectiveSceneIds() []*int32 {
	return s.EffectiveSceneIds
}

func (s *CreateTrafficControlTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateTrafficControlTaskRequest) GetExecutionTime() *string {
	return s.ExecutionTime
}

func (s *CreateTrafficControlTaskRequest) GetFlinkResourceId() *string {
	return s.FlinkResourceId
}

func (s *CreateTrafficControlTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTrafficControlTaskRequest) GetItemConditionArray() *string {
	return s.ItemConditionArray
}

func (s *CreateTrafficControlTaskRequest) GetItemConditionExpress() *string {
	return s.ItemConditionExpress
}

func (s *CreateTrafficControlTaskRequest) GetItemConditionType() *string {
	return s.ItemConditionType
}

func (s *CreateTrafficControlTaskRequest) GetItemTableMetaId() *string {
	return s.ItemTableMetaId
}

func (s *CreateTrafficControlTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateTrafficControlTaskRequest) GetPreExperimentIds() *string {
	return s.PreExperimentIds
}

func (s *CreateTrafficControlTaskRequest) GetProdExperimentIds() *string {
	return s.ProdExperimentIds
}

func (s *CreateTrafficControlTaskRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *CreateTrafficControlTaskRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateTrafficControlTaskRequest) GetServiceIds() []*int32 {
	return s.ServiceIds
}

func (s *CreateTrafficControlTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateTrafficControlTaskRequest) GetStatisBehaviorConditionArray() *string {
	return s.StatisBehaviorConditionArray
}

func (s *CreateTrafficControlTaskRequest) GetStatisBehaviorConditionExpress() *string {
	return s.StatisBehaviorConditionExpress
}

func (s *CreateTrafficControlTaskRequest) GetStatisBehaviorConditionType() *string {
	return s.StatisBehaviorConditionType
}

func (s *CreateTrafficControlTaskRequest) GetTrafficControlTargets() []*CreateTrafficControlTaskRequestTrafficControlTargets {
	return s.TrafficControlTargets
}

func (s *CreateTrafficControlTaskRequest) GetUserConditionArray() *string {
	return s.UserConditionArray
}

func (s *CreateTrafficControlTaskRequest) GetUserConditionExpress() *string {
	return s.UserConditionExpress
}

func (s *CreateTrafficControlTaskRequest) GetUserConditionType() *string {
	return s.UserConditionType
}

func (s *CreateTrafficControlTaskRequest) GetUserTableMetaId() *string {
	return s.UserTableMetaId
}

func (s *CreateTrafficControlTaskRequest) SetBehaviorTableMetaId(v string) *CreateTrafficControlTaskRequest {
	s.BehaviorTableMetaId = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetControlGranularity(v string) *CreateTrafficControlTaskRequest {
	s.ControlGranularity = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetControlLogic(v string) *CreateTrafficControlTaskRequest {
	s.ControlLogic = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetControlType(v string) *CreateTrafficControlTaskRequest {
	s.ControlType = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetDescription(v string) *CreateTrafficControlTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetEffectiveSceneIds(v []*int32) *CreateTrafficControlTaskRequest {
	s.EffectiveSceneIds = v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetEndTime(v string) *CreateTrafficControlTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetExecutionTime(v string) *CreateTrafficControlTaskRequest {
	s.ExecutionTime = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetFlinkResourceId(v string) *CreateTrafficControlTaskRequest {
	s.FlinkResourceId = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetInstanceId(v string) *CreateTrafficControlTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetItemConditionArray(v string) *CreateTrafficControlTaskRequest {
	s.ItemConditionArray = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetItemConditionExpress(v string) *CreateTrafficControlTaskRequest {
	s.ItemConditionExpress = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetItemConditionType(v string) *CreateTrafficControlTaskRequest {
	s.ItemConditionType = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetItemTableMetaId(v string) *CreateTrafficControlTaskRequest {
	s.ItemTableMetaId = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetName(v string) *CreateTrafficControlTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetPreExperimentIds(v string) *CreateTrafficControlTaskRequest {
	s.PreExperimentIds = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetProdExperimentIds(v string) *CreateTrafficControlTaskRequest {
	s.ProdExperimentIds = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetSceneId(v string) *CreateTrafficControlTaskRequest {
	s.SceneId = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetServiceId(v string) *CreateTrafficControlTaskRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetServiceIds(v []*int32) *CreateTrafficControlTaskRequest {
	s.ServiceIds = v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetStartTime(v string) *CreateTrafficControlTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetStatisBehaviorConditionArray(v string) *CreateTrafficControlTaskRequest {
	s.StatisBehaviorConditionArray = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetStatisBehaviorConditionExpress(v string) *CreateTrafficControlTaskRequest {
	s.StatisBehaviorConditionExpress = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetStatisBehaviorConditionType(v string) *CreateTrafficControlTaskRequest {
	s.StatisBehaviorConditionType = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetTrafficControlTargets(v []*CreateTrafficControlTaskRequestTrafficControlTargets) *CreateTrafficControlTaskRequest {
	s.TrafficControlTargets = v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetUserConditionArray(v string) *CreateTrafficControlTaskRequest {
	s.UserConditionArray = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetUserConditionExpress(v string) *CreateTrafficControlTaskRequest {
	s.UserConditionExpress = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetUserConditionType(v string) *CreateTrafficControlTaskRequest {
	s.UserConditionType = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) SetUserTableMetaId(v string) *CreateTrafficControlTaskRequest {
	s.UserTableMetaId = &v
	return s
}

func (s *CreateTrafficControlTaskRequest) Validate() error {
	return dara.Validate(s)
}

type CreateTrafficControlTaskRequestTrafficControlTargets struct {
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
	// [{\\"field\\":\\"status\\",\\"option\\":\\"=\\",\\"value\\":\\"1\\"}]
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
	// target_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// false
	NewProductRegulation *bool `json:"NewProductRegulation,omitempty" xml:"NewProductRegulation,omitempty"`
	// example:
	//
	// recall_1
	RecallName *string `json:"RecallName,omitempty" xml:"RecallName,omitempty"`
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
	// 20
	ToleranceValue *int64 `json:"ToleranceValue,omitempty" xml:"ToleranceValue,omitempty"`
	// example:
	//
	// 100
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTrafficControlTaskRequestTrafficControlTargets) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficControlTaskRequestTrafficControlTargets) GoString() string {
	return s.String()
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) GetEvent() *string {
	return s.Event
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) GetItemConditionArray() *string {
	return s.ItemConditionArray
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) GetItemConditionExpress() *string {
	return s.ItemConditionExpress
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) GetItemConditionType() *string {
	return s.ItemConditionType
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) GetName() *string {
	return s.Name
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) GetNewProductRegulation() *bool {
	return s.NewProductRegulation
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) GetRecallName() *string {
	return s.RecallName
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) GetStatisPeriod() *string {
	return s.StatisPeriod
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) GetStatus() *string {
	return s.Status
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) GetToleranceValue() *int64 {
	return s.ToleranceValue
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) GetValue() *float32 {
	return s.Value
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetEndTime(v string) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.EndTime = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetEvent(v string) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.Event = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetItemConditionArray(v string) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.ItemConditionArray = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetItemConditionExpress(v string) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.ItemConditionExpress = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetItemConditionType(v string) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.ItemConditionType = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetName(v string) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.Name = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetNewProductRegulation(v bool) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.NewProductRegulation = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetRecallName(v string) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.RecallName = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetStartTime(v string) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.StartTime = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetStatisPeriod(v string) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.StatisPeriod = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetStatus(v string) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.Status = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetToleranceValue(v int64) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.ToleranceValue = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) SetValue(v float32) *CreateTrafficControlTaskRequestTrafficControlTargets {
	s.Value = &v
	return s
}

func (s *CreateTrafficControlTaskRequestTrafficControlTargets) Validate() error {
	return dara.Validate(s)
}
