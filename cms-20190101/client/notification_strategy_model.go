// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotificationStrategy interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *NotificationStrategy
	GetCreateTime() *string
	SetDescription(v string) *NotificationStrategy
	GetDescription() *string
	SetEscalationSetting(v *NotificationStrategyEscalationSetting) *NotificationStrategy
	GetEscalationSetting() *NotificationStrategyEscalationSetting
	SetFilterSetting(v *NotificationStrategyFilterSetting) *NotificationStrategy
	GetFilterSetting() *NotificationStrategyFilterSetting
	SetGroupingSetting(v *NotificationStrategyGroupingSetting) *NotificationStrategy
	GetGroupingSetting() *NotificationStrategyGroupingSetting
	SetName(v string) *NotificationStrategy
	GetName() *string
	SetProduct(v string) *NotificationStrategy
	GetProduct() *string
	SetPushingSetting(v *NotificationStrategyPushingSetting) *NotificationStrategy
	GetPushingSetting() *NotificationStrategyPushingSetting
	SetRouteSetting(v *NotificationStrategyRouteSetting) *NotificationStrategy
	GetRouteSetting() *NotificationStrategyRouteSetting
	SetUpdateTime(v string) *NotificationStrategy
	GetUpdateTime() *string
	SetUserId(v string) *NotificationStrategy
	GetUserId() *string
	SetUuid(v string) *NotificationStrategy
	GetUuid() *string
}

type NotificationStrategy struct {
	CreateTime        *string                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description       *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	EscalationSetting *NotificationStrategyEscalationSetting `json:"EscalationSetting,omitempty" xml:"EscalationSetting,omitempty" type:"Struct"`
	FilterSetting     *NotificationStrategyFilterSetting     `json:"FilterSetting,omitempty" xml:"FilterSetting,omitempty" type:"Struct"`
	GroupingSetting   *NotificationStrategyGroupingSetting   `json:"GroupingSetting,omitempty" xml:"GroupingSetting,omitempty" type:"Struct"`
	// This parameter is required.
	Name           *string                             `json:"Name,omitempty" xml:"Name,omitempty"`
	Product        *string                             `json:"Product,omitempty" xml:"Product,omitempty"`
	PushingSetting *NotificationStrategyPushingSetting `json:"PushingSetting,omitempty" xml:"PushingSetting,omitempty" type:"Struct"`
	RouteSetting   *NotificationStrategyRouteSetting   `json:"RouteSetting,omitempty" xml:"RouteSetting,omitempty" type:"Struct"`
	UpdateTime     *string                             `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId         *string                             `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Uuid           *string                             `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s NotificationStrategy) String() string {
	return dara.Prettify(s)
}

func (s NotificationStrategy) GoString() string {
	return s.String()
}

func (s *NotificationStrategy) GetCreateTime() *string {
	return s.CreateTime
}

func (s *NotificationStrategy) GetDescription() *string {
	return s.Description
}

func (s *NotificationStrategy) GetEscalationSetting() *NotificationStrategyEscalationSetting {
	return s.EscalationSetting
}

func (s *NotificationStrategy) GetFilterSetting() *NotificationStrategyFilterSetting {
	return s.FilterSetting
}

func (s *NotificationStrategy) GetGroupingSetting() *NotificationStrategyGroupingSetting {
	return s.GroupingSetting
}

func (s *NotificationStrategy) GetName() *string {
	return s.Name
}

func (s *NotificationStrategy) GetProduct() *string {
	return s.Product
}

func (s *NotificationStrategy) GetPushingSetting() *NotificationStrategyPushingSetting {
	return s.PushingSetting
}

func (s *NotificationStrategy) GetRouteSetting() *NotificationStrategyRouteSetting {
	return s.RouteSetting
}

func (s *NotificationStrategy) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *NotificationStrategy) GetUserId() *string {
	return s.UserId
}

func (s *NotificationStrategy) GetUuid() *string {
	return s.Uuid
}

func (s *NotificationStrategy) SetCreateTime(v string) *NotificationStrategy {
	s.CreateTime = &v
	return s
}

func (s *NotificationStrategy) SetDescription(v string) *NotificationStrategy {
	s.Description = &v
	return s
}

func (s *NotificationStrategy) SetEscalationSetting(v *NotificationStrategyEscalationSetting) *NotificationStrategy {
	s.EscalationSetting = v
	return s
}

func (s *NotificationStrategy) SetFilterSetting(v *NotificationStrategyFilterSetting) *NotificationStrategy {
	s.FilterSetting = v
	return s
}

func (s *NotificationStrategy) SetGroupingSetting(v *NotificationStrategyGroupingSetting) *NotificationStrategy {
	s.GroupingSetting = v
	return s
}

func (s *NotificationStrategy) SetName(v string) *NotificationStrategy {
	s.Name = &v
	return s
}

func (s *NotificationStrategy) SetProduct(v string) *NotificationStrategy {
	s.Product = &v
	return s
}

func (s *NotificationStrategy) SetPushingSetting(v *NotificationStrategyPushingSetting) *NotificationStrategy {
	s.PushingSetting = v
	return s
}

func (s *NotificationStrategy) SetRouteSetting(v *NotificationStrategyRouteSetting) *NotificationStrategy {
	s.RouteSetting = v
	return s
}

func (s *NotificationStrategy) SetUpdateTime(v string) *NotificationStrategy {
	s.UpdateTime = &v
	return s
}

func (s *NotificationStrategy) SetUserId(v string) *NotificationStrategy {
	s.UserId = &v
	return s
}

func (s *NotificationStrategy) SetUuid(v string) *NotificationStrategy {
	s.Uuid = &v
	return s
}

func (s *NotificationStrategy) Validate() error {
	return dara.Validate(s)
}

type NotificationStrategyEscalationSetting struct {
	AutoResolveMin  *int64                                                 `json:"AutoResolveMin,omitempty" xml:"AutoResolveMin,omitempty"`
	CustomChannels  []*NotificationStrategyEscalationSettingCustomChannels `json:"CustomChannels,omitempty" xml:"CustomChannels,omitempty" type:"Repeated"`
	EscalationLevel *string                                                `json:"EscalationLevel,omitempty" xml:"EscalationLevel,omitempty"`
	EscalationUuid  *string                                                `json:"EscalationUuid,omitempty" xml:"EscalationUuid,omitempty"`
	Range           *string                                                `json:"Range,omitempty" xml:"Range,omitempty"`
	RetriggerMin    *int64                                                 `json:"RetriggerMin,omitempty" xml:"RetriggerMin,omitempty"`
}

func (s NotificationStrategyEscalationSetting) String() string {
	return dara.Prettify(s)
}

func (s NotificationStrategyEscalationSetting) GoString() string {
	return s.String()
}

func (s *NotificationStrategyEscalationSetting) GetAutoResolveMin() *int64 {
	return s.AutoResolveMin
}

func (s *NotificationStrategyEscalationSetting) GetCustomChannels() []*NotificationStrategyEscalationSettingCustomChannels {
	return s.CustomChannels
}

func (s *NotificationStrategyEscalationSetting) GetEscalationLevel() *string {
	return s.EscalationLevel
}

func (s *NotificationStrategyEscalationSetting) GetEscalationUuid() *string {
	return s.EscalationUuid
}

func (s *NotificationStrategyEscalationSetting) GetRange() *string {
	return s.Range
}

func (s *NotificationStrategyEscalationSetting) GetRetriggerMin() *int64 {
	return s.RetriggerMin
}

func (s *NotificationStrategyEscalationSetting) SetAutoResolveMin(v int64) *NotificationStrategyEscalationSetting {
	s.AutoResolveMin = &v
	return s
}

func (s *NotificationStrategyEscalationSetting) SetCustomChannels(v []*NotificationStrategyEscalationSettingCustomChannels) *NotificationStrategyEscalationSetting {
	s.CustomChannels = v
	return s
}

func (s *NotificationStrategyEscalationSetting) SetEscalationLevel(v string) *NotificationStrategyEscalationSetting {
	s.EscalationLevel = &v
	return s
}

func (s *NotificationStrategyEscalationSetting) SetEscalationUuid(v string) *NotificationStrategyEscalationSetting {
	s.EscalationUuid = &v
	return s
}

func (s *NotificationStrategyEscalationSetting) SetRange(v string) *NotificationStrategyEscalationSetting {
	s.Range = &v
	return s
}

func (s *NotificationStrategyEscalationSetting) SetRetriggerMin(v int64) *NotificationStrategyEscalationSetting {
	s.RetriggerMin = &v
	return s
}

func (s *NotificationStrategyEscalationSetting) Validate() error {
	return dara.Validate(s)
}

type NotificationStrategyEscalationSettingCustomChannels struct {
	// This parameter is required.
	ChannelType  *string   `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	Severities   []*string `json:"Severities,omitempty" xml:"Severities,omitempty" type:"Repeated"`
	TemplateUuid *string   `json:"TemplateUuid,omitempty" xml:"TemplateUuid,omitempty"`
}

func (s NotificationStrategyEscalationSettingCustomChannels) String() string {
	return dara.Prettify(s)
}

func (s NotificationStrategyEscalationSettingCustomChannels) GoString() string {
	return s.String()
}

func (s *NotificationStrategyEscalationSettingCustomChannels) GetChannelType() *string {
	return s.ChannelType
}

func (s *NotificationStrategyEscalationSettingCustomChannels) GetSeverities() []*string {
	return s.Severities
}

func (s *NotificationStrategyEscalationSettingCustomChannels) GetTemplateUuid() *string {
	return s.TemplateUuid
}

func (s *NotificationStrategyEscalationSettingCustomChannels) SetChannelType(v string) *NotificationStrategyEscalationSettingCustomChannels {
	s.ChannelType = &v
	return s
}

func (s *NotificationStrategyEscalationSettingCustomChannels) SetSeverities(v []*string) *NotificationStrategyEscalationSettingCustomChannels {
	s.Severities = v
	return s
}

func (s *NotificationStrategyEscalationSettingCustomChannels) SetTemplateUuid(v string) *NotificationStrategyEscalationSettingCustomChannels {
	s.TemplateUuid = &v
	return s
}

func (s *NotificationStrategyEscalationSettingCustomChannels) Validate() error {
	return dara.Validate(s)
}

type NotificationStrategyFilterSetting struct {
	BlackList [][]*NotificationStrategyFilterSettingBlackList `json:"BlackList,omitempty" xml:"BlackList,omitempty" type:"Repeated"`
	WhiteList [][]*NotificationStrategyFilterSettingWhiteList `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s NotificationStrategyFilterSetting) String() string {
	return dara.Prettify(s)
}

func (s NotificationStrategyFilterSetting) GoString() string {
	return s.String()
}

func (s *NotificationStrategyFilterSetting) GetBlackList() [][]*NotificationStrategyFilterSettingBlackList {
	return s.BlackList
}

func (s *NotificationStrategyFilterSetting) GetWhiteList() [][]*NotificationStrategyFilterSettingWhiteList {
	return s.WhiteList
}

func (s *NotificationStrategyFilterSetting) SetBlackList(v [][]*NotificationStrategyFilterSettingBlackList) *NotificationStrategyFilterSetting {
	s.BlackList = v
	return s
}

func (s *NotificationStrategyFilterSetting) SetWhiteList(v [][]*NotificationStrategyFilterSettingWhiteList) *NotificationStrategyFilterSetting {
	s.WhiteList = v
	return s
}

func (s *NotificationStrategyFilterSetting) Validate() error {
	return dara.Validate(s)
}

type NotificationStrategyFilterSettingBlackList struct {
	// This parameter is required.
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// This parameter is required.
	Op *string `json:"Op,omitempty" xml:"Op,omitempty"`
	// This parameter is required.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s NotificationStrategyFilterSettingBlackList) String() string {
	return dara.Prettify(s)
}

func (s NotificationStrategyFilterSettingBlackList) GoString() string {
	return s.String()
}

func (s *NotificationStrategyFilterSettingBlackList) GetField() *string {
	return s.Field
}

func (s *NotificationStrategyFilterSettingBlackList) GetOp() *string {
	return s.Op
}

func (s *NotificationStrategyFilterSettingBlackList) GetValue() *string {
	return s.Value
}

func (s *NotificationStrategyFilterSettingBlackList) SetField(v string) *NotificationStrategyFilterSettingBlackList {
	s.Field = &v
	return s
}

func (s *NotificationStrategyFilterSettingBlackList) SetOp(v string) *NotificationStrategyFilterSettingBlackList {
	s.Op = &v
	return s
}

func (s *NotificationStrategyFilterSettingBlackList) SetValue(v string) *NotificationStrategyFilterSettingBlackList {
	s.Value = &v
	return s
}

func (s *NotificationStrategyFilterSettingBlackList) Validate() error {
	return dara.Validate(s)
}

type NotificationStrategyFilterSettingWhiteList struct {
	// This parameter is required.
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// This parameter is required.
	Op *string `json:"Op,omitempty" xml:"Op,omitempty"`
	// This parameter is required.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s NotificationStrategyFilterSettingWhiteList) String() string {
	return dara.Prettify(s)
}

func (s NotificationStrategyFilterSettingWhiteList) GoString() string {
	return s.String()
}

func (s *NotificationStrategyFilterSettingWhiteList) GetField() *string {
	return s.Field
}

func (s *NotificationStrategyFilterSettingWhiteList) GetOp() *string {
	return s.Op
}

func (s *NotificationStrategyFilterSettingWhiteList) GetValue() *string {
	return s.Value
}

func (s *NotificationStrategyFilterSettingWhiteList) SetField(v string) *NotificationStrategyFilterSettingWhiteList {
	s.Field = &v
	return s
}

func (s *NotificationStrategyFilterSettingWhiteList) SetOp(v string) *NotificationStrategyFilterSettingWhiteList {
	s.Op = &v
	return s
}

func (s *NotificationStrategyFilterSettingWhiteList) SetValue(v string) *NotificationStrategyFilterSettingWhiteList {
	s.Value = &v
	return s
}

func (s *NotificationStrategyFilterSettingWhiteList) Validate() error {
	return dara.Validate(s)
}

type NotificationStrategyGroupingSetting struct {
	EnableRawAlertDispatching *bool                                               `json:"EnableRawAlertDispatching,omitempty" xml:"EnableRawAlertDispatching,omitempty"`
	GroupingItems             []*NotificationStrategyGroupingSettingGroupingItems `json:"GroupingItems,omitempty" xml:"GroupingItems,omitempty" type:"Repeated"`
	PeriodMin                 *int32                                              `json:"PeriodMin,omitempty" xml:"PeriodMin,omitempty"`
	SilenceSec                *int32                                              `json:"SilenceSec,omitempty" xml:"SilenceSec,omitempty"`
	Times                     *int32                                              `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s NotificationStrategyGroupingSetting) String() string {
	return dara.Prettify(s)
}

func (s NotificationStrategyGroupingSetting) GoString() string {
	return s.String()
}

func (s *NotificationStrategyGroupingSetting) GetEnableRawAlertDispatching() *bool {
	return s.EnableRawAlertDispatching
}

func (s *NotificationStrategyGroupingSetting) GetGroupingItems() []*NotificationStrategyGroupingSettingGroupingItems {
	return s.GroupingItems
}

func (s *NotificationStrategyGroupingSetting) GetPeriodMin() *int32 {
	return s.PeriodMin
}

func (s *NotificationStrategyGroupingSetting) GetSilenceSec() *int32 {
	return s.SilenceSec
}

func (s *NotificationStrategyGroupingSetting) GetTimes() *int32 {
	return s.Times
}

func (s *NotificationStrategyGroupingSetting) SetEnableRawAlertDispatching(v bool) *NotificationStrategyGroupingSetting {
	s.EnableRawAlertDispatching = &v
	return s
}

func (s *NotificationStrategyGroupingSetting) SetGroupingItems(v []*NotificationStrategyGroupingSettingGroupingItems) *NotificationStrategyGroupingSetting {
	s.GroupingItems = v
	return s
}

func (s *NotificationStrategyGroupingSetting) SetPeriodMin(v int32) *NotificationStrategyGroupingSetting {
	s.PeriodMin = &v
	return s
}

func (s *NotificationStrategyGroupingSetting) SetSilenceSec(v int32) *NotificationStrategyGroupingSetting {
	s.SilenceSec = &v
	return s
}

func (s *NotificationStrategyGroupingSetting) SetTimes(v int32) *NotificationStrategyGroupingSetting {
	s.Times = &v
	return s
}

func (s *NotificationStrategyGroupingSetting) Validate() error {
	return dara.Validate(s)
}

type NotificationStrategyGroupingSettingGroupingItems struct {
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	Type *string   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s NotificationStrategyGroupingSettingGroupingItems) String() string {
	return dara.Prettify(s)
}

func (s NotificationStrategyGroupingSettingGroupingItems) GoString() string {
	return s.String()
}

func (s *NotificationStrategyGroupingSettingGroupingItems) GetKeys() []*string {
	return s.Keys
}

func (s *NotificationStrategyGroupingSettingGroupingItems) GetType() *string {
	return s.Type
}

func (s *NotificationStrategyGroupingSettingGroupingItems) SetKeys(v []*string) *NotificationStrategyGroupingSettingGroupingItems {
	s.Keys = v
	return s
}

func (s *NotificationStrategyGroupingSettingGroupingItems) SetType(v string) *NotificationStrategyGroupingSettingGroupingItems {
	s.Type = &v
	return s
}

func (s *NotificationStrategyGroupingSettingGroupingItems) Validate() error {
	return dara.Validate(s)
}

type NotificationStrategyPushingSetting struct {
	PushingDataFormat *string   `json:"PushingDataFormat,omitempty" xml:"PushingDataFormat,omitempty"`
	Range             *string   `json:"Range,omitempty" xml:"Range,omitempty"`
	TargetUuids       []*string `json:"TargetUuids,omitempty" xml:"TargetUuids,omitempty" type:"Repeated"`
	TemplateUuid      *string   `json:"TemplateUuid,omitempty" xml:"TemplateUuid,omitempty"`
}

func (s NotificationStrategyPushingSetting) String() string {
	return dara.Prettify(s)
}

func (s NotificationStrategyPushingSetting) GoString() string {
	return s.String()
}

func (s *NotificationStrategyPushingSetting) GetPushingDataFormat() *string {
	return s.PushingDataFormat
}

func (s *NotificationStrategyPushingSetting) GetRange() *string {
	return s.Range
}

func (s *NotificationStrategyPushingSetting) GetTargetUuids() []*string {
	return s.TargetUuids
}

func (s *NotificationStrategyPushingSetting) GetTemplateUuid() *string {
	return s.TemplateUuid
}

func (s *NotificationStrategyPushingSetting) SetPushingDataFormat(v string) *NotificationStrategyPushingSetting {
	s.PushingDataFormat = &v
	return s
}

func (s *NotificationStrategyPushingSetting) SetRange(v string) *NotificationStrategyPushingSetting {
	s.Range = &v
	return s
}

func (s *NotificationStrategyPushingSetting) SetTargetUuids(v []*string) *NotificationStrategyPushingSetting {
	s.TargetUuids = v
	return s
}

func (s *NotificationStrategyPushingSetting) SetTemplateUuid(v string) *NotificationStrategyPushingSetting {
	s.TemplateUuid = &v
	return s
}

func (s *NotificationStrategyPushingSetting) Validate() error {
	return dara.Validate(s)
}

type NotificationStrategyRouteSetting struct {
	Routes []*NotificationStrategyRouteSettingRoutes `json:"Routes,omitempty" xml:"Routes,omitempty" type:"Repeated"`
}

func (s NotificationStrategyRouteSetting) String() string {
	return dara.Prettify(s)
}

func (s NotificationStrategyRouteSetting) GoString() string {
	return s.String()
}

func (s *NotificationStrategyRouteSetting) GetRoutes() []*NotificationStrategyRouteSettingRoutes {
	return s.Routes
}

func (s *NotificationStrategyRouteSetting) SetRoutes(v []*NotificationStrategyRouteSettingRoutes) *NotificationStrategyRouteSetting {
	s.Routes = v
	return s
}

func (s *NotificationStrategyRouteSetting) Validate() error {
	return dara.Validate(s)
}

type NotificationStrategyRouteSettingRoutes struct {
	Conditions     []*NotificationStrategyRouteSettingRoutesConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	EscalationUuid *string                                             `json:"EscalationUuid,omitempty" xml:"EscalationUuid,omitempty"`
}

func (s NotificationStrategyRouteSettingRoutes) String() string {
	return dara.Prettify(s)
}

func (s NotificationStrategyRouteSettingRoutes) GoString() string {
	return s.String()
}

func (s *NotificationStrategyRouteSettingRoutes) GetConditions() []*NotificationStrategyRouteSettingRoutesConditions {
	return s.Conditions
}

func (s *NotificationStrategyRouteSettingRoutes) GetEscalationUuid() *string {
	return s.EscalationUuid
}

func (s *NotificationStrategyRouteSettingRoutes) SetConditions(v []*NotificationStrategyRouteSettingRoutesConditions) *NotificationStrategyRouteSettingRoutes {
	s.Conditions = v
	return s
}

func (s *NotificationStrategyRouteSettingRoutes) SetEscalationUuid(v string) *NotificationStrategyRouteSettingRoutes {
	s.EscalationUuid = &v
	return s
}

func (s *NotificationStrategyRouteSettingRoutes) Validate() error {
	return dara.Validate(s)
}

type NotificationStrategyRouteSettingRoutesConditions struct {
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	Op    *string `json:"Op,omitempty" xml:"Op,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s NotificationStrategyRouteSettingRoutesConditions) String() string {
	return dara.Prettify(s)
}

func (s NotificationStrategyRouteSettingRoutesConditions) GoString() string {
	return s.String()
}

func (s *NotificationStrategyRouteSettingRoutesConditions) GetField() *string {
	return s.Field
}

func (s *NotificationStrategyRouteSettingRoutesConditions) GetOp() *string {
	return s.Op
}

func (s *NotificationStrategyRouteSettingRoutesConditions) GetValue() *string {
	return s.Value
}

func (s *NotificationStrategyRouteSettingRoutesConditions) SetField(v string) *NotificationStrategyRouteSettingRoutesConditions {
	s.Field = &v
	return s
}

func (s *NotificationStrategyRouteSettingRoutesConditions) SetOp(v string) *NotificationStrategyRouteSettingRoutesConditions {
	s.Op = &v
	return s
}

func (s *NotificationStrategyRouteSettingRoutesConditions) SetValue(v string) *NotificationStrategyRouteSettingRoutesConditions {
	s.Value = &v
	return s
}

func (s *NotificationStrategyRouteSettingRoutesConditions) Validate() error {
	return dara.Validate(s)
}
