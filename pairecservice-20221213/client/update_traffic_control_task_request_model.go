// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficControlTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBehaviorTableMetaId(v string) *UpdateTrafficControlTaskRequest
	GetBehaviorTableMetaId() *string
	SetControlGranularity(v string) *UpdateTrafficControlTaskRequest
	GetControlGranularity() *string
	SetControlLogic(v string) *UpdateTrafficControlTaskRequest
	GetControlLogic() *string
	SetControlType(v string) *UpdateTrafficControlTaskRequest
	GetControlType() *string
	SetDescription(v string) *UpdateTrafficControlTaskRequest
	GetDescription() *string
	SetEffectiveSceneIds(v []*int32) *UpdateTrafficControlTaskRequest
	GetEffectiveSceneIds() []*int32
	SetEndTime(v string) *UpdateTrafficControlTaskRequest
	GetEndTime() *string
	SetExecutionTime(v string) *UpdateTrafficControlTaskRequest
	GetExecutionTime() *string
	SetFlinkResourceId(v string) *UpdateTrafficControlTaskRequest
	GetFlinkResourceId() *string
	SetInstanceId(v string) *UpdateTrafficControlTaskRequest
	GetInstanceId() *string
	SetItemConditionArray(v string) *UpdateTrafficControlTaskRequest
	GetItemConditionArray() *string
	SetItemConditionExpress(v string) *UpdateTrafficControlTaskRequest
	GetItemConditionExpress() *string
	SetItemConditionType(v string) *UpdateTrafficControlTaskRequest
	GetItemConditionType() *string
	SetItemTableMetaId(v string) *UpdateTrafficControlTaskRequest
	GetItemTableMetaId() *string
	SetName(v string) *UpdateTrafficControlTaskRequest
	GetName() *string
	SetPreExperimentIds(v string) *UpdateTrafficControlTaskRequest
	GetPreExperimentIds() *string
	SetProdExperimentIds(v string) *UpdateTrafficControlTaskRequest
	GetProdExperimentIds() *string
	SetSceneId(v string) *UpdateTrafficControlTaskRequest
	GetSceneId() *string
	SetServiceId(v string) *UpdateTrafficControlTaskRequest
	GetServiceId() *string
	SetServiceIds(v []*int32) *UpdateTrafficControlTaskRequest
	GetServiceIds() []*int32
	SetStartTime(v string) *UpdateTrafficControlTaskRequest
	GetStartTime() *string
	SetStatisBaeaviorConditionArray(v string) *UpdateTrafficControlTaskRequest
	GetStatisBaeaviorConditionArray() *string
	SetStatisBehaviorConditionExpress(v string) *UpdateTrafficControlTaskRequest
	GetStatisBehaviorConditionExpress() *string
	SetStatisBehaviorConditionType(v string) *UpdateTrafficControlTaskRequest
	GetStatisBehaviorConditionType() *string
	SetTrafficControlTargets(v []*UpdateTrafficControlTaskRequestTrafficControlTargets) *UpdateTrafficControlTaskRequest
	GetTrafficControlTargets() []*UpdateTrafficControlTaskRequestTrafficControlTargets
	SetUserConditionArray(v string) *UpdateTrafficControlTaskRequest
	GetUserConditionArray() *string
	SetUserConditionExpress(v string) *UpdateTrafficControlTaskRequest
	GetUserConditionExpress() *string
	SetUserConditionType(v string) *UpdateTrafficControlTaskRequest
	GetUserConditionType() *string
	SetUserTableMetaId(v string) *UpdateTrafficControlTaskRequest
	GetUserTableMetaId() *string
}

type UpdateTrafficControlTaskRequest struct {
	BehaviorTableMetaId            *string                                                 `json:"BehaviorTableMetaId,omitempty" xml:"BehaviorTableMetaId,omitempty"`
	ControlGranularity             *string                                                 `json:"ControlGranularity,omitempty" xml:"ControlGranularity,omitempty"`
	ControlLogic                   *string                                                 `json:"ControlLogic,omitempty" xml:"ControlLogic,omitempty"`
	ControlType                    *string                                                 `json:"ControlType,omitempty" xml:"ControlType,omitempty"`
	Description                    *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	EffectiveSceneIds              []*int32                                                `json:"EffectiveSceneIds,omitempty" xml:"EffectiveSceneIds,omitempty" type:"Repeated"`
	EndTime                        *string                                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExecutionTime                  *string                                                 `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	FlinkResourceId                *string                                                 `json:"FlinkResourceId,omitempty" xml:"FlinkResourceId,omitempty"`
	InstanceId                     *string                                                 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemConditionArray             *string                                                 `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	ItemConditionExpress           *string                                                 `json:"ItemConditionExpress,omitempty" xml:"ItemConditionExpress,omitempty"`
	ItemConditionType              *string                                                 `json:"ItemConditionType,omitempty" xml:"ItemConditionType,omitempty"`
	ItemTableMetaId                *string                                                 `json:"ItemTableMetaId,omitempty" xml:"ItemTableMetaId,omitempty"`
	Name                           *string                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	PreExperimentIds               *string                                                 `json:"PreExperimentIds,omitempty" xml:"PreExperimentIds,omitempty"`
	ProdExperimentIds              *string                                                 `json:"ProdExperimentIds,omitempty" xml:"ProdExperimentIds,omitempty"`
	SceneId                        *string                                                 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	ServiceId                      *string                                                 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceIds                     []*int32                                                `json:"ServiceIds,omitempty" xml:"ServiceIds,omitempty" type:"Repeated"`
	StartTime                      *string                                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatisBaeaviorConditionArray   *string                                                 `json:"StatisBaeaviorConditionArray,omitempty" xml:"StatisBaeaviorConditionArray,omitempty"`
	StatisBehaviorConditionExpress *string                                                 `json:"StatisBehaviorConditionExpress,omitempty" xml:"StatisBehaviorConditionExpress,omitempty"`
	StatisBehaviorConditionType    *string                                                 `json:"StatisBehaviorConditionType,omitempty" xml:"StatisBehaviorConditionType,omitempty"`
	TrafficControlTargets          []*UpdateTrafficControlTaskRequestTrafficControlTargets `json:"TrafficControlTargets,omitempty" xml:"TrafficControlTargets,omitempty" type:"Repeated"`
	UserConditionArray             *string                                                 `json:"UserConditionArray,omitempty" xml:"UserConditionArray,omitempty"`
	UserConditionExpress           *string                                                 `json:"UserConditionExpress,omitempty" xml:"UserConditionExpress,omitempty"`
	UserConditionType              *string                                                 `json:"UserConditionType,omitempty" xml:"UserConditionType,omitempty"`
	UserTableMetaId                *string                                                 `json:"UserTableMetaId,omitempty" xml:"UserTableMetaId,omitempty"`
}

func (s UpdateTrafficControlTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficControlTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTaskRequest) GetBehaviorTableMetaId() *string {
	return s.BehaviorTableMetaId
}

func (s *UpdateTrafficControlTaskRequest) GetControlGranularity() *string {
	return s.ControlGranularity
}

func (s *UpdateTrafficControlTaskRequest) GetControlLogic() *string {
	return s.ControlLogic
}

func (s *UpdateTrafficControlTaskRequest) GetControlType() *string {
	return s.ControlType
}

func (s *UpdateTrafficControlTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTrafficControlTaskRequest) GetEffectiveSceneIds() []*int32 {
	return s.EffectiveSceneIds
}

func (s *UpdateTrafficControlTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateTrafficControlTaskRequest) GetExecutionTime() *string {
	return s.ExecutionTime
}

func (s *UpdateTrafficControlTaskRequest) GetFlinkResourceId() *string {
	return s.FlinkResourceId
}

func (s *UpdateTrafficControlTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateTrafficControlTaskRequest) GetItemConditionArray() *string {
	return s.ItemConditionArray
}

func (s *UpdateTrafficControlTaskRequest) GetItemConditionExpress() *string {
	return s.ItemConditionExpress
}

func (s *UpdateTrafficControlTaskRequest) GetItemConditionType() *string {
	return s.ItemConditionType
}

func (s *UpdateTrafficControlTaskRequest) GetItemTableMetaId() *string {
	return s.ItemTableMetaId
}

func (s *UpdateTrafficControlTaskRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTrafficControlTaskRequest) GetPreExperimentIds() *string {
	return s.PreExperimentIds
}

func (s *UpdateTrafficControlTaskRequest) GetProdExperimentIds() *string {
	return s.ProdExperimentIds
}

func (s *UpdateTrafficControlTaskRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *UpdateTrafficControlTaskRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateTrafficControlTaskRequest) GetServiceIds() []*int32 {
	return s.ServiceIds
}

func (s *UpdateTrafficControlTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateTrafficControlTaskRequest) GetStatisBaeaviorConditionArray() *string {
	return s.StatisBaeaviorConditionArray
}

func (s *UpdateTrafficControlTaskRequest) GetStatisBehaviorConditionExpress() *string {
	return s.StatisBehaviorConditionExpress
}

func (s *UpdateTrafficControlTaskRequest) GetStatisBehaviorConditionType() *string {
	return s.StatisBehaviorConditionType
}

func (s *UpdateTrafficControlTaskRequest) GetTrafficControlTargets() []*UpdateTrafficControlTaskRequestTrafficControlTargets {
	return s.TrafficControlTargets
}

func (s *UpdateTrafficControlTaskRequest) GetUserConditionArray() *string {
	return s.UserConditionArray
}

func (s *UpdateTrafficControlTaskRequest) GetUserConditionExpress() *string {
	return s.UserConditionExpress
}

func (s *UpdateTrafficControlTaskRequest) GetUserConditionType() *string {
	return s.UserConditionType
}

func (s *UpdateTrafficControlTaskRequest) GetUserTableMetaId() *string {
	return s.UserTableMetaId
}

func (s *UpdateTrafficControlTaskRequest) SetBehaviorTableMetaId(v string) *UpdateTrafficControlTaskRequest {
	s.BehaviorTableMetaId = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetControlGranularity(v string) *UpdateTrafficControlTaskRequest {
	s.ControlGranularity = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetControlLogic(v string) *UpdateTrafficControlTaskRequest {
	s.ControlLogic = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetControlType(v string) *UpdateTrafficControlTaskRequest {
	s.ControlType = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetDescription(v string) *UpdateTrafficControlTaskRequest {
	s.Description = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetEffectiveSceneIds(v []*int32) *UpdateTrafficControlTaskRequest {
	s.EffectiveSceneIds = v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetEndTime(v string) *UpdateTrafficControlTaskRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetExecutionTime(v string) *UpdateTrafficControlTaskRequest {
	s.ExecutionTime = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetFlinkResourceId(v string) *UpdateTrafficControlTaskRequest {
	s.FlinkResourceId = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetInstanceId(v string) *UpdateTrafficControlTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetItemConditionArray(v string) *UpdateTrafficControlTaskRequest {
	s.ItemConditionArray = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetItemConditionExpress(v string) *UpdateTrafficControlTaskRequest {
	s.ItemConditionExpress = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetItemConditionType(v string) *UpdateTrafficControlTaskRequest {
	s.ItemConditionType = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetItemTableMetaId(v string) *UpdateTrafficControlTaskRequest {
	s.ItemTableMetaId = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetName(v string) *UpdateTrafficControlTaskRequest {
	s.Name = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetPreExperimentIds(v string) *UpdateTrafficControlTaskRequest {
	s.PreExperimentIds = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetProdExperimentIds(v string) *UpdateTrafficControlTaskRequest {
	s.ProdExperimentIds = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetSceneId(v string) *UpdateTrafficControlTaskRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetServiceId(v string) *UpdateTrafficControlTaskRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetServiceIds(v []*int32) *UpdateTrafficControlTaskRequest {
	s.ServiceIds = v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetStartTime(v string) *UpdateTrafficControlTaskRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetStatisBaeaviorConditionArray(v string) *UpdateTrafficControlTaskRequest {
	s.StatisBaeaviorConditionArray = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetStatisBehaviorConditionExpress(v string) *UpdateTrafficControlTaskRequest {
	s.StatisBehaviorConditionExpress = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetStatisBehaviorConditionType(v string) *UpdateTrafficControlTaskRequest {
	s.StatisBehaviorConditionType = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetTrafficControlTargets(v []*UpdateTrafficControlTaskRequestTrafficControlTargets) *UpdateTrafficControlTaskRequest {
	s.TrafficControlTargets = v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetUserConditionArray(v string) *UpdateTrafficControlTaskRequest {
	s.UserConditionArray = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetUserConditionExpress(v string) *UpdateTrafficControlTaskRequest {
	s.UserConditionExpress = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetUserConditionType(v string) *UpdateTrafficControlTaskRequest {
	s.UserConditionType = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) SetUserTableMetaId(v string) *UpdateTrafficControlTaskRequest {
	s.UserTableMetaId = &v
	return s
}

func (s *UpdateTrafficControlTaskRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateTrafficControlTaskRequestTrafficControlTargets struct {
	EndTime              *string  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Event                *string  `json:"Event,omitempty" xml:"Event,omitempty"`
	ItemConditionArray   *string  `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	ItemConditionExpress *string  `json:"ItemConditionExpress,omitempty" xml:"ItemConditionExpress,omitempty"`
	ItemConditionType    *string  `json:"ItemConditionType,omitempty" xml:"ItemConditionType,omitempty"`
	Name                 *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	NewProductRegulation *bool    `json:"NewProductRegulation,omitempty" xml:"NewProductRegulation,omitempty"`
	RecallName           *string  `json:"RecallName,omitempty" xml:"RecallName,omitempty"`
	StartTime            *string  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatisPeriod         *string  `json:"StatisPeriod,omitempty" xml:"StatisPeriod,omitempty"`
	Status               *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	ToleranceValue       *int64   `json:"ToleranceValue,omitempty" xml:"ToleranceValue,omitempty"`
	Value                *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateTrafficControlTaskRequestTrafficControlTargets) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficControlTaskRequestTrafficControlTargets) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) GetEvent() *string {
	return s.Event
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) GetItemConditionArray() *string {
	return s.ItemConditionArray
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) GetItemConditionExpress() *string {
	return s.ItemConditionExpress
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) GetItemConditionType() *string {
	return s.ItemConditionType
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) GetName() *string {
	return s.Name
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) GetNewProductRegulation() *bool {
	return s.NewProductRegulation
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) GetRecallName() *string {
	return s.RecallName
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) GetStatisPeriod() *string {
	return s.StatisPeriod
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) GetStatus() *string {
	return s.Status
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) GetToleranceValue() *int64 {
	return s.ToleranceValue
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) GetValue() *float32 {
	return s.Value
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetEndTime(v string) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.EndTime = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetEvent(v string) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.Event = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetItemConditionArray(v string) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.ItemConditionArray = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetItemConditionExpress(v string) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.ItemConditionExpress = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetItemConditionType(v string) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.ItemConditionType = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetName(v string) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.Name = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetNewProductRegulation(v bool) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.NewProductRegulation = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetRecallName(v string) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.RecallName = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetStartTime(v string) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.StartTime = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetStatisPeriod(v string) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.StatisPeriod = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetStatus(v string) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.Status = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetToleranceValue(v int64) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.ToleranceValue = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) SetValue(v float32) *UpdateTrafficControlTaskRequestTrafficControlTargets {
	s.Value = &v
	return s
}

func (s *UpdateTrafficControlTaskRequestTrafficControlTargets) Validate() error {
	return dara.Validate(s)
}
