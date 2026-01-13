// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrafficControlTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBehaviorTableMetaId(v string) *GetTrafficControlTaskResponseBody
	GetBehaviorTableMetaId() *string
	SetControlGranularity(v string) *GetTrafficControlTaskResponseBody
	GetControlGranularity() *string
	SetControlLogic(v string) *GetTrafficControlTaskResponseBody
	GetControlLogic() *string
	SetControlType(v string) *GetTrafficControlTaskResponseBody
	GetControlType() *string
	SetDescription(v string) *GetTrafficControlTaskResponseBody
	GetDescription() *string
	SetEffectiveSceneIds(v []*int32) *GetTrafficControlTaskResponseBody
	GetEffectiveSceneIds() []*int32
	SetEffectiveSceneNames(v []*string) *GetTrafficControlTaskResponseBody
	GetEffectiveSceneNames() []*string
	SetEndTime(v string) *GetTrafficControlTaskResponseBody
	GetEndTime() *string
	SetEverPublished(v bool) *GetTrafficControlTaskResponseBody
	GetEverPublished() *bool
	SetExecutionTime(v string) *GetTrafficControlTaskResponseBody
	GetExecutionTime() *string
	SetFlinkResourceId(v string) *GetTrafficControlTaskResponseBody
	GetFlinkResourceId() *string
	SetFlinkResourceName(v string) *GetTrafficControlTaskResponseBody
	GetFlinkResourceName() *string
	SetGmtCreateTime(v string) *GetTrafficControlTaskResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetTrafficControlTaskResponseBody
	GetGmtModifiedTime() *string
	SetItemConditionArray(v string) *GetTrafficControlTaskResponseBody
	GetItemConditionArray() *string
	SetItemConditionExpress(v string) *GetTrafficControlTaskResponseBody
	GetItemConditionExpress() *string
	SetItemConditionType(v string) *GetTrafficControlTaskResponseBody
	GetItemConditionType() *string
	SetItemTableMetaId(v string) *GetTrafficControlTaskResponseBody
	GetItemTableMetaId() *string
	SetName(v string) *GetTrafficControlTaskResponseBody
	GetName() *string
	SetPreExperimentIds(v string) *GetTrafficControlTaskResponseBody
	GetPreExperimentIds() *string
	SetPrepubStatus(v string) *GetTrafficControlTaskResponseBody
	GetPrepubStatus() *string
	SetProdExperimentIds(v string) *GetTrafficControlTaskResponseBody
	GetProdExperimentIds() *string
	SetProductStatus(v string) *GetTrafficControlTaskResponseBody
	GetProductStatus() *string
	SetRequestId(v string) *GetTrafficControlTaskResponseBody
	GetRequestId() *string
	SetSceneId(v string) *GetTrafficControlTaskResponseBody
	GetSceneId() *string
	SetSceneName(v string) *GetTrafficControlTaskResponseBody
	GetSceneName() *string
	SetServiceId(v string) *GetTrafficControlTaskResponseBody
	GetServiceId() *string
	SetServiceIds(v []*int32) *GetTrafficControlTaskResponseBody
	GetServiceIds() []*int32
	SetStartTime(v string) *GetTrafficControlTaskResponseBody
	GetStartTime() *string
	SetStatisBehaviorConditionArray(v string) *GetTrafficControlTaskResponseBody
	GetStatisBehaviorConditionArray() *string
	SetStatisBehaviorConditionExpress(v string) *GetTrafficControlTaskResponseBody
	GetStatisBehaviorConditionExpress() *string
	SetStatisBehaviorConditionType(v string) *GetTrafficControlTaskResponseBody
	GetStatisBehaviorConditionType() *string
	SetTrafficControlTargets(v []*GetTrafficControlTaskResponseBodyTrafficControlTargets) *GetTrafficControlTaskResponseBody
	GetTrafficControlTargets() []*GetTrafficControlTaskResponseBodyTrafficControlTargets
	SetTrafficControlTaskId(v string) *GetTrafficControlTaskResponseBody
	GetTrafficControlTaskId() *string
	SetUserConditionArray(v string) *GetTrafficControlTaskResponseBody
	GetUserConditionArray() *string
	SetUserConditionExpress(v string) *GetTrafficControlTaskResponseBody
	GetUserConditionExpress() *string
	SetUserConditionType(v string) *GetTrafficControlTaskResponseBody
	GetUserConditionType() *string
	SetUserTableMetaId(v string) *GetTrafficControlTaskResponseBody
	GetUserTableMetaId() *string
}

type GetTrafficControlTaskResponseBody struct {
	BehaviorTableMetaId            *string                                                   `json:"BehaviorTableMetaId,omitempty" xml:"BehaviorTableMetaId,omitempty"`
	ControlGranularity             *string                                                   `json:"ControlGranularity,omitempty" xml:"ControlGranularity,omitempty"`
	ControlLogic                   *string                                                   `json:"ControlLogic,omitempty" xml:"ControlLogic,omitempty"`
	ControlType                    *string                                                   `json:"ControlType,omitempty" xml:"ControlType,omitempty"`
	Description                    *string                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	EffectiveSceneIds              []*int32                                                  `json:"EffectiveSceneIds,omitempty" xml:"EffectiveSceneIds,omitempty" type:"Repeated"`
	EffectiveSceneNames            []*string                                                 `json:"EffectiveSceneNames,omitempty" xml:"EffectiveSceneNames,omitempty" type:"Repeated"`
	EndTime                        *string                                                   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EverPublished                  *bool                                                     `json:"EverPublished,omitempty" xml:"EverPublished,omitempty"`
	ExecutionTime                  *string                                                   `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	FlinkResourceId                *string                                                   `json:"FlinkResourceId,omitempty" xml:"FlinkResourceId,omitempty"`
	FlinkResourceName              *string                                                   `json:"FlinkResourceName,omitempty" xml:"FlinkResourceName,omitempty"`
	GmtCreateTime                  *string                                                   `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime                *string                                                   `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	ItemConditionArray             *string                                                   `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	ItemConditionExpress           *string                                                   `json:"ItemConditionExpress,omitempty" xml:"ItemConditionExpress,omitempty"`
	ItemConditionType              *string                                                   `json:"ItemConditionType,omitempty" xml:"ItemConditionType,omitempty"`
	ItemTableMetaId                *string                                                   `json:"ItemTableMetaId,omitempty" xml:"ItemTableMetaId,omitempty"`
	Name                           *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	PreExperimentIds               *string                                                   `json:"PreExperimentIds,omitempty" xml:"PreExperimentIds,omitempty"`
	PrepubStatus                   *string                                                   `json:"PrepubStatus,omitempty" xml:"PrepubStatus,omitempty"`
	ProdExperimentIds              *string                                                   `json:"ProdExperimentIds,omitempty" xml:"ProdExperimentIds,omitempty"`
	ProductStatus                  *string                                                   `json:"ProductStatus,omitempty" xml:"ProductStatus,omitempty"`
	RequestId                      *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SceneId                        *string                                                   `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName                      *string                                                   `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	ServiceId                      *string                                                   `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceIds                     []*int32                                                  `json:"ServiceIds,omitempty" xml:"ServiceIds,omitempty" type:"Repeated"`
	StartTime                      *string                                                   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatisBehaviorConditionArray   *string                                                   `json:"StatisBehaviorConditionArray,omitempty" xml:"StatisBehaviorConditionArray,omitempty"`
	StatisBehaviorConditionExpress *string                                                   `json:"StatisBehaviorConditionExpress,omitempty" xml:"StatisBehaviorConditionExpress,omitempty"`
	StatisBehaviorConditionType    *string                                                   `json:"StatisBehaviorConditionType,omitempty" xml:"StatisBehaviorConditionType,omitempty"`
	TrafficControlTargets          []*GetTrafficControlTaskResponseBodyTrafficControlTargets `json:"TrafficControlTargets,omitempty" xml:"TrafficControlTargets,omitempty" type:"Repeated"`
	TrafficControlTaskId           *string                                                   `json:"TrafficControlTaskId,omitempty" xml:"TrafficControlTaskId,omitempty"`
	UserConditionArray             *string                                                   `json:"UserConditionArray,omitempty" xml:"UserConditionArray,omitempty"`
	UserConditionExpress           *string                                                   `json:"UserConditionExpress,omitempty" xml:"UserConditionExpress,omitempty"`
	UserConditionType              *string                                                   `json:"UserConditionType,omitempty" xml:"UserConditionType,omitempty"`
	UserTableMetaId                *string                                                   `json:"UserTableMetaId,omitempty" xml:"UserTableMetaId,omitempty"`
}

func (s GetTrafficControlTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTrafficControlTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTaskResponseBody) GetBehaviorTableMetaId() *string {
	return s.BehaviorTableMetaId
}

func (s *GetTrafficControlTaskResponseBody) GetControlGranularity() *string {
	return s.ControlGranularity
}

func (s *GetTrafficControlTaskResponseBody) GetControlLogic() *string {
	return s.ControlLogic
}

func (s *GetTrafficControlTaskResponseBody) GetControlType() *string {
	return s.ControlType
}

func (s *GetTrafficControlTaskResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetTrafficControlTaskResponseBody) GetEffectiveSceneIds() []*int32 {
	return s.EffectiveSceneIds
}

func (s *GetTrafficControlTaskResponseBody) GetEffectiveSceneNames() []*string {
	return s.EffectiveSceneNames
}

func (s *GetTrafficControlTaskResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetTrafficControlTaskResponseBody) GetEverPublished() *bool {
	return s.EverPublished
}

func (s *GetTrafficControlTaskResponseBody) GetExecutionTime() *string {
	return s.ExecutionTime
}

func (s *GetTrafficControlTaskResponseBody) GetFlinkResourceId() *string {
	return s.FlinkResourceId
}

func (s *GetTrafficControlTaskResponseBody) GetFlinkResourceName() *string {
	return s.FlinkResourceName
}

func (s *GetTrafficControlTaskResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetTrafficControlTaskResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetTrafficControlTaskResponseBody) GetItemConditionArray() *string {
	return s.ItemConditionArray
}

func (s *GetTrafficControlTaskResponseBody) GetItemConditionExpress() *string {
	return s.ItemConditionExpress
}

func (s *GetTrafficControlTaskResponseBody) GetItemConditionType() *string {
	return s.ItemConditionType
}

func (s *GetTrafficControlTaskResponseBody) GetItemTableMetaId() *string {
	return s.ItemTableMetaId
}

func (s *GetTrafficControlTaskResponseBody) GetName() *string {
	return s.Name
}

func (s *GetTrafficControlTaskResponseBody) GetPreExperimentIds() *string {
	return s.PreExperimentIds
}

func (s *GetTrafficControlTaskResponseBody) GetPrepubStatus() *string {
	return s.PrepubStatus
}

func (s *GetTrafficControlTaskResponseBody) GetProdExperimentIds() *string {
	return s.ProdExperimentIds
}

func (s *GetTrafficControlTaskResponseBody) GetProductStatus() *string {
	return s.ProductStatus
}

func (s *GetTrafficControlTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTrafficControlTaskResponseBody) GetSceneId() *string {
	return s.SceneId
}

func (s *GetTrafficControlTaskResponseBody) GetSceneName() *string {
	return s.SceneName
}

func (s *GetTrafficControlTaskResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetTrafficControlTaskResponseBody) GetServiceIds() []*int32 {
	return s.ServiceIds
}

func (s *GetTrafficControlTaskResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetTrafficControlTaskResponseBody) GetStatisBehaviorConditionArray() *string {
	return s.StatisBehaviorConditionArray
}

func (s *GetTrafficControlTaskResponseBody) GetStatisBehaviorConditionExpress() *string {
	return s.StatisBehaviorConditionExpress
}

func (s *GetTrafficControlTaskResponseBody) GetStatisBehaviorConditionType() *string {
	return s.StatisBehaviorConditionType
}

func (s *GetTrafficControlTaskResponseBody) GetTrafficControlTargets() []*GetTrafficControlTaskResponseBodyTrafficControlTargets {
	return s.TrafficControlTargets
}

func (s *GetTrafficControlTaskResponseBody) GetTrafficControlTaskId() *string {
	return s.TrafficControlTaskId
}

func (s *GetTrafficControlTaskResponseBody) GetUserConditionArray() *string {
	return s.UserConditionArray
}

func (s *GetTrafficControlTaskResponseBody) GetUserConditionExpress() *string {
	return s.UserConditionExpress
}

func (s *GetTrafficControlTaskResponseBody) GetUserConditionType() *string {
	return s.UserConditionType
}

func (s *GetTrafficControlTaskResponseBody) GetUserTableMetaId() *string {
	return s.UserTableMetaId
}

func (s *GetTrafficControlTaskResponseBody) SetBehaviorTableMetaId(v string) *GetTrafficControlTaskResponseBody {
	s.BehaviorTableMetaId = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetControlGranularity(v string) *GetTrafficControlTaskResponseBody {
	s.ControlGranularity = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetControlLogic(v string) *GetTrafficControlTaskResponseBody {
	s.ControlLogic = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetControlType(v string) *GetTrafficControlTaskResponseBody {
	s.ControlType = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetDescription(v string) *GetTrafficControlTaskResponseBody {
	s.Description = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetEffectiveSceneIds(v []*int32) *GetTrafficControlTaskResponseBody {
	s.EffectiveSceneIds = v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetEffectiveSceneNames(v []*string) *GetTrafficControlTaskResponseBody {
	s.EffectiveSceneNames = v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetEndTime(v string) *GetTrafficControlTaskResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetEverPublished(v bool) *GetTrafficControlTaskResponseBody {
	s.EverPublished = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetExecutionTime(v string) *GetTrafficControlTaskResponseBody {
	s.ExecutionTime = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetFlinkResourceId(v string) *GetTrafficControlTaskResponseBody {
	s.FlinkResourceId = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetFlinkResourceName(v string) *GetTrafficControlTaskResponseBody {
	s.FlinkResourceName = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetGmtCreateTime(v string) *GetTrafficControlTaskResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetGmtModifiedTime(v string) *GetTrafficControlTaskResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetItemConditionArray(v string) *GetTrafficControlTaskResponseBody {
	s.ItemConditionArray = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetItemConditionExpress(v string) *GetTrafficControlTaskResponseBody {
	s.ItemConditionExpress = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetItemConditionType(v string) *GetTrafficControlTaskResponseBody {
	s.ItemConditionType = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetItemTableMetaId(v string) *GetTrafficControlTaskResponseBody {
	s.ItemTableMetaId = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetName(v string) *GetTrafficControlTaskResponseBody {
	s.Name = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetPreExperimentIds(v string) *GetTrafficControlTaskResponseBody {
	s.PreExperimentIds = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetPrepubStatus(v string) *GetTrafficControlTaskResponseBody {
	s.PrepubStatus = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetProdExperimentIds(v string) *GetTrafficControlTaskResponseBody {
	s.ProdExperimentIds = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetProductStatus(v string) *GetTrafficControlTaskResponseBody {
	s.ProductStatus = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetRequestId(v string) *GetTrafficControlTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetSceneId(v string) *GetTrafficControlTaskResponseBody {
	s.SceneId = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetSceneName(v string) *GetTrafficControlTaskResponseBody {
	s.SceneName = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetServiceId(v string) *GetTrafficControlTaskResponseBody {
	s.ServiceId = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetServiceIds(v []*int32) *GetTrafficControlTaskResponseBody {
	s.ServiceIds = v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetStartTime(v string) *GetTrafficControlTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetStatisBehaviorConditionArray(v string) *GetTrafficControlTaskResponseBody {
	s.StatisBehaviorConditionArray = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetStatisBehaviorConditionExpress(v string) *GetTrafficControlTaskResponseBody {
	s.StatisBehaviorConditionExpress = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetStatisBehaviorConditionType(v string) *GetTrafficControlTaskResponseBody {
	s.StatisBehaviorConditionType = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetTrafficControlTargets(v []*GetTrafficControlTaskResponseBodyTrafficControlTargets) *GetTrafficControlTaskResponseBody {
	s.TrafficControlTargets = v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetTrafficControlTaskId(v string) *GetTrafficControlTaskResponseBody {
	s.TrafficControlTaskId = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetUserConditionArray(v string) *GetTrafficControlTaskResponseBody {
	s.UserConditionArray = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetUserConditionExpress(v string) *GetTrafficControlTaskResponseBody {
	s.UserConditionExpress = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetUserConditionType(v string) *GetTrafficControlTaskResponseBody {
	s.UserConditionType = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) SetUserTableMetaId(v string) *GetTrafficControlTaskResponseBody {
	s.UserTableMetaId = &v
	return s
}

func (s *GetTrafficControlTaskResponseBody) Validate() error {
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

type GetTrafficControlTaskResponseBodyTrafficControlTargets struct {
	EndTime                *string                                                           `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Event                  *string                                                           `json:"Event,omitempty" xml:"Event,omitempty"`
	GmtCreateTime          *string                                                           `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime        *string                                                           `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	ItemConditionArray     *string                                                           `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	ItemConditionExpress   *string                                                           `json:"ItemConditionExpress,omitempty" xml:"ItemConditionExpress,omitempty"`
	ItemConditionType      *string                                                           `json:"ItemConditionType,omitempty" xml:"ItemConditionType,omitempty"`
	Name                   *string                                                           `json:"Name,omitempty" xml:"Name,omitempty"`
	NewProductRegulation   *bool                                                             `json:"NewProductRegulation,omitempty" xml:"NewProductRegulation,omitempty"`
	RecallName             *string                                                           `json:"RecallName,omitempty" xml:"RecallName,omitempty"`
	SplitParts             *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts `json:"SplitParts,omitempty" xml:"SplitParts,omitempty" type:"Struct"`
	StartTime              *string                                                           `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatisPeriod           *string                                                           `json:"StatisPeriod,omitempty" xml:"StatisPeriod,omitempty"`
	Status                 *string                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	ToleranceValue         *int64                                                            `json:"ToleranceValue,omitempty" xml:"ToleranceValue,omitempty"`
	TrafficControlTargetId *string                                                           `json:"TrafficControlTargetId,omitempty" xml:"TrafficControlTargetId,omitempty"`
	TrafficControlTaskId   *string                                                           `json:"TrafficControlTaskId,omitempty" xml:"TrafficControlTaskId,omitempty"`
	Value                  *float32                                                          `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTrafficControlTaskResponseBodyTrafficControlTargets) String() string {
	return dara.Prettify(s)
}

func (s GetTrafficControlTaskResponseBodyTrafficControlTargets) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) GetEndTime() *string {
	return s.EndTime
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) GetEvent() *string {
	return s.Event
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) GetItemConditionArray() *string {
	return s.ItemConditionArray
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) GetItemConditionExpress() *string {
	return s.ItemConditionExpress
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) GetItemConditionType() *string {
	return s.ItemConditionType
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) GetName() *string {
	return s.Name
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) GetNewProductRegulation() *bool {
	return s.NewProductRegulation
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) GetRecallName() *string {
	return s.RecallName
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) GetSplitParts() *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts {
	return s.SplitParts
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) GetStartTime() *string {
	return s.StartTime
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) GetStatisPeriod() *string {
	return s.StatisPeriod
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) GetStatus() *string {
	return s.Status
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) GetToleranceValue() *int64 {
	return s.ToleranceValue
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) GetTrafficControlTargetId() *string {
	return s.TrafficControlTargetId
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) GetTrafficControlTaskId() *string {
	return s.TrafficControlTaskId
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) GetValue() *float32 {
	return s.Value
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetEndTime(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.EndTime = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetEvent(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.Event = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetGmtCreateTime(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.GmtCreateTime = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetGmtModifiedTime(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetItemConditionArray(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.ItemConditionArray = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetItemConditionExpress(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.ItemConditionExpress = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetItemConditionType(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.ItemConditionType = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetName(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.Name = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetNewProductRegulation(v bool) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.NewProductRegulation = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetRecallName(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.RecallName = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetSplitParts(v *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.SplitParts = v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetStartTime(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.StartTime = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetStatisPeriod(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.StatisPeriod = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetStatus(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.Status = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetToleranceValue(v int64) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.ToleranceValue = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetTrafficControlTargetId(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.TrafficControlTargetId = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetTrafficControlTaskId(v string) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.TrafficControlTaskId = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) SetValue(v float32) *GetTrafficControlTaskResponseBodyTrafficControlTargets {
	s.Value = &v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargets) Validate() error {
	if s.SplitParts != nil {
		if err := s.SplitParts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts struct {
	SetPoints  []*int32 `json:"SetPoints,omitempty" xml:"SetPoints,omitempty" type:"Repeated"`
	SetValues  []*int64 `json:"SetValues,omitempty" xml:"SetValues,omitempty" type:"Repeated"`
	TimePoints []*int32 `json:"TimePoints,omitempty" xml:"TimePoints,omitempty" type:"Repeated"`
}

func (s GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts) String() string {
	return dara.Prettify(s)
}

func (s GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts) GetSetPoints() []*int32 {
	return s.SetPoints
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts) GetSetValues() []*int64 {
	return s.SetValues
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts) GetTimePoints() []*int32 {
	return s.TimePoints
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts) SetSetPoints(v []*int32) *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts {
	s.SetPoints = v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts) SetSetValues(v []*int64) *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts {
	s.SetValues = v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts) SetTimePoints(v []*int32) *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts {
	s.TimePoints = v
	return s
}

func (s *GetTrafficControlTaskResponseBodyTrafficControlTargetsSplitParts) Validate() error {
	return dara.Validate(s)
}
