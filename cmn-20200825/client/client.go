// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type OrderStep struct {
	// 步骤展示方式
	DisplayMethod *string `json:"DisplayMethod,omitempty" xml:"DisplayMethod,omitempty"`
	// 创建时间
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 修改时间
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// 步骤任务参数
	JobKwargs *string `json:"JobKwargs,omitempty" xml:"JobKwargs,omitempty"`
	// 步骤标题
	JobMessage *string `json:"JobMessage,omitempty" xml:"JobMessage,omitempty"`
	// 步骤标题
	JobReturnStatus *string `json:"JobReturnStatus,omitempty" xml:"JobReturnStatus,omitempty"`
	// 步骤任务返回
	JobReturnValues *string `json:"JobReturnValues,omitempty" xml:"JobReturnValues,omitempty"`
	// 步骤任务系统
	JobSystem *string `json:"JobSystem,omitempty" xml:"JobSystem,omitempty"`
	// 工单id
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// 工单步骤id
	OrderStepId *string `json:"OrderStepId,omitempty" xml:"OrderStepId,omitempty"`
	// 下一步步骤名
	RealNextStep *string `json:"RealNextStep,omitempty" xml:"RealNextStep,omitempty"`
	// 下一步步骤可选列表
	Restriction []*OrderStepRestriction `json:"Restriction,omitempty" xml:"Restriction,omitempty" type:"Repeated"`
	// 步骤名
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// 步骤状态
	StepStatus *string `json:"StepStatus,omitempty" xml:"StepStatus,omitempty"`
	// 步骤标题
	StepTitle *string `json:"StepTitle,omitempty" xml:"StepTitle,omitempty"`
	// 步骤类型
	StepType *string `json:"StepType,omitempty" xml:"StepType,omitempty"`
}

func (s OrderStep) String() string {
	return tea.Prettify(s)
}

func (s OrderStep) GoString() string {
	return s.String()
}

func (s *OrderStep) SetDisplayMethod(v string) *OrderStep {
	s.DisplayMethod = &v
	return s
}

func (s *OrderStep) SetGmtCreate(v string) *OrderStep {
	s.GmtCreate = &v
	return s
}

func (s *OrderStep) SetGmtModify(v string) *OrderStep {
	s.GmtModify = &v
	return s
}

func (s *OrderStep) SetJobKwargs(v string) *OrderStep {
	s.JobKwargs = &v
	return s
}

func (s *OrderStep) SetJobMessage(v string) *OrderStep {
	s.JobMessage = &v
	return s
}

func (s *OrderStep) SetJobReturnStatus(v string) *OrderStep {
	s.JobReturnStatus = &v
	return s
}

func (s *OrderStep) SetJobReturnValues(v string) *OrderStep {
	s.JobReturnValues = &v
	return s
}

func (s *OrderStep) SetJobSystem(v string) *OrderStep {
	s.JobSystem = &v
	return s
}

func (s *OrderStep) SetOrderId(v string) *OrderStep {
	s.OrderId = &v
	return s
}

func (s *OrderStep) SetOrderStepId(v string) *OrderStep {
	s.OrderStepId = &v
	return s
}

func (s *OrderStep) SetRealNextStep(v string) *OrderStep {
	s.RealNextStep = &v
	return s
}

func (s *OrderStep) SetRestriction(v []*OrderStepRestriction) *OrderStep {
	s.Restriction = v
	return s
}

func (s *OrderStep) SetStepName(v string) *OrderStep {
	s.StepName = &v
	return s
}

func (s *OrderStep) SetStepStatus(v string) *OrderStep {
	s.StepStatus = &v
	return s
}

func (s *OrderStep) SetStepTitle(v string) *OrderStep {
	s.StepTitle = &v
	return s
}

func (s *OrderStep) SetStepType(v string) *OrderStep {
	s.StepType = &v
	return s
}

type OrderStepRestriction struct {
	// 步骤标题
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// 步骤名
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OrderStepRestriction) String() string {
	return tea.Prettify(s)
}

func (s OrderStepRestriction) GoString() string {
	return s.String()
}

func (s *OrderStepRestriction) SetLabel(v string) *OrderStepRestriction {
	s.Label = &v
	return s
}

func (s *OrderStepRestriction) SetValue(v string) *OrderStepRestriction {
	s.Value = &v
	return s
}

type DeviceForm struct {
	// 是否需要配置账号信息
	AccountConfig *string `json:"AccountConfig,omitempty" xml:"AccountConfig,omitempty"`
	// 是否需要展示配置备份
	ConfigCompare *string `json:"ConfigCompare,omitempty" xml:"ConfigCompare,omitempty"`
	// 设备形态ID
	FormId *string `json:"FormId,omitempty" xml:"FormId,omitempty"`
	// 设备形态名称
	FormName *string `json:"FormName,omitempty" xml:"FormName,omitempty"`
	// 设备形态属性列表
	PropertiesList []*DeviceFormProperty `json:"PropertiesList,omitempty" xml:"PropertiesList,omitempty" type:"Repeated"`
}

func (s DeviceForm) String() string {
	return tea.Prettify(s)
}

func (s DeviceForm) GoString() string {
	return s.String()
}

func (s *DeviceForm) SetAccountConfig(v string) *DeviceForm {
	s.AccountConfig = &v
	return s
}

func (s *DeviceForm) SetConfigCompare(v string) *DeviceForm {
	s.ConfigCompare = &v
	return s
}

func (s *DeviceForm) SetFormId(v string) *DeviceForm {
	s.FormId = &v
	return s
}

func (s *DeviceForm) SetFormName(v string) *DeviceForm {
	s.FormName = &v
	return s
}

func (s *DeviceForm) SetPropertiesList(v []*DeviceFormProperty) *DeviceForm {
	s.PropertiesList = v
	return s
}

type ScheduleType struct {
	// 资源一级ID
	ScheduleTypeId *string `json:"ScheduleTypeId,omitempty" xml:"ScheduleTypeId,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 值班类型key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 值班类型value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// 值班类型状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// relatedWorkerStr
	RelatedWorker []*string `json:"RelatedWorker,omitempty" xml:"RelatedWorker,omitempty" type:"Repeated"`
}

func (s ScheduleType) String() string {
	return tea.Prettify(s)
}

func (s ScheduleType) GoString() string {
	return s.String()
}

func (s *ScheduleType) SetScheduleTypeId(v string) *ScheduleType {
	s.ScheduleTypeId = &v
	return s
}

func (s *ScheduleType) SetCreateTime(v string) *ScheduleType {
	s.CreateTime = &v
	return s
}

func (s *ScheduleType) SetUpdateTime(v string) *ScheduleType {
	s.UpdateTime = &v
	return s
}

func (s *ScheduleType) SetKey(v string) *ScheduleType {
	s.Key = &v
	return s
}

func (s *ScheduleType) SetValue(v string) *ScheduleType {
	s.Value = &v
	return s
}

func (s *ScheduleType) SetStatus(v string) *ScheduleType {
	s.Status = &v
	return s
}

func (s *ScheduleType) SetRelatedWorker(v []*string) *ScheduleType {
	s.RelatedWorker = v
	return s
}

type Task struct {
	// 模板类别
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 创建时间
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 修改时间
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// 任务参数
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// 任务错误码
	ResponseCode *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	// 任务返回
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// 任务状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 任务id
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 模板id
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 模板名称
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s Task) String() string {
	return tea.Prettify(s)
}

func (s Task) GoString() string {
	return s.String()
}

func (s *Task) SetCategory(v string) *Task {
	s.Category = &v
	return s
}

func (s *Task) SetGmtCreate(v string) *Task {
	s.GmtCreate = &v
	return s
}

func (s *Task) SetGmtModify(v string) *Task {
	s.GmtModify = &v
	return s
}

func (s *Task) SetParams(v string) *Task {
	s.Params = &v
	return s
}

func (s *Task) SetResponseCode(v string) *Task {
	s.ResponseCode = &v
	return s
}

func (s *Task) SetResult(v string) *Task {
	s.Result = &v
	return s
}

func (s *Task) SetStatus(v string) *Task {
	s.Status = &v
	return s
}

func (s *Task) SetTaskId(v string) *Task {
	s.TaskId = &v
	return s
}

func (s *Task) SetTemplateId(v string) *Task {
	s.TemplateId = &v
	return s
}

func (s *Task) SetTemplateName(v string) *Task {
	s.TemplateName = &v
	return s
}

type Scheme struct {
	// 方案类型
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 方案内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 方案说明
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 创建时间
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 修改时间
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// 方案入参
	Input []*SchemeInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Repeated"`
	// 方案出参
	Output []*SchemeOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Repeated"`
	// 方案id
	SchemeId *string `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	// 方案名称
	SchemeName *string `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
	// 方案状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 方案展示
	View *string `json:"View,omitempty" xml:"View,omitempty"`
}

func (s Scheme) String() string {
	return tea.Prettify(s)
}

func (s Scheme) GoString() string {
	return s.String()
}

func (s *Scheme) SetCategory(v string) *Scheme {
	s.Category = &v
	return s
}

func (s *Scheme) SetContent(v string) *Scheme {
	s.Content = &v
	return s
}

func (s *Scheme) SetDescription(v string) *Scheme {
	s.Description = &v
	return s
}

func (s *Scheme) SetGmtCreate(v string) *Scheme {
	s.GmtCreate = &v
	return s
}

func (s *Scheme) SetGmtModify(v string) *Scheme {
	s.GmtModify = &v
	return s
}

func (s *Scheme) SetInput(v []*SchemeInput) *Scheme {
	s.Input = v
	return s
}

func (s *Scheme) SetOutput(v []*SchemeOutput) *Scheme {
	s.Output = v
	return s
}

func (s *Scheme) SetSchemeId(v string) *Scheme {
	s.SchemeId = &v
	return s
}

func (s *Scheme) SetSchemeName(v string) *Scheme {
	s.SchemeName = &v
	return s
}

func (s *Scheme) SetStatus(v string) *Scheme {
	s.Status = &v
	return s
}

func (s *Scheme) SetView(v string) *Scheme {
	s.View = &v
	return s
}

type SchemeInput struct {
	// 参数说明
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 参数名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 参数示例
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// 参数类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SchemeInput) String() string {
	return tea.Prettify(s)
}

func (s SchemeInput) GoString() string {
	return s.String()
}

func (s *SchemeInput) SetDescription(v string) *SchemeInput {
	s.Description = &v
	return s
}

func (s *SchemeInput) SetName(v string) *SchemeInput {
	s.Name = &v
	return s
}

func (s *SchemeInput) SetSample(v string) *SchemeInput {
	s.Sample = &v
	return s
}

func (s *SchemeInput) SetType(v string) *SchemeInput {
	s.Type = &v
	return s
}

type SchemeOutput struct {
	// 参数说明
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 参数名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 参数示例
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// 参数类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SchemeOutput) String() string {
	return tea.Prettify(s)
}

func (s SchemeOutput) GoString() string {
	return s.String()
}

func (s *SchemeOutput) SetDescription(v string) *SchemeOutput {
	s.Description = &v
	return s
}

func (s *SchemeOutput) SetName(v string) *SchemeOutput {
	s.Name = &v
	return s
}

func (s *SchemeOutput) SetSample(v string) *SchemeOutput {
	s.Sample = &v
	return s
}

func (s *SchemeOutput) SetType(v string) *SchemeOutput {
	s.Type = &v
	return s
}

type AggregateData struct {
	// 创建时间
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 修改时间
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 数据项
	DataItem *string `json:"DataItem,omitempty" xml:"DataItem,omitempty"`
	// 聚合设备ID列表
	DeviceIdList []*string `json:"DeviceIdList,omitempty" xml:"DeviceIdList,omitempty" type:"Repeated"`
	// 聚合方式列表
	AggregateModeList []*string `json:"AggregateModeList,omitempty" xml:"AggregateModeList,omitempty" type:"Repeated"`
	// 聚合数据名称
	AggregateDataName *string `json:"AggregateDataName,omitempty" xml:"AggregateDataName,omitempty"`
	// 描述
	AggregateDataDescription *string `json:"AggregateDataDescription,omitempty" xml:"AggregateDataDescription,omitempty"`
	// 是否聚合全部设备
	IsAllDevice *int32 `json:"IsAllDevice,omitempty" xml:"IsAllDevice,omitempty"`
	// 监控项ID
	MonitorItemId *string `json:"MonitorItemId,omitempty" xml:"MonitorItemId,omitempty"`
	// 聚合数据ID
	AggregateDataId *string `json:"AggregateDataId,omitempty" xml:"AggregateDataId,omitempty"`
}

func (s AggregateData) String() string {
	return tea.Prettify(s)
}

func (s AggregateData) GoString() string {
	return s.String()
}

func (s *AggregateData) SetGmtCreate(v string) *AggregateData {
	s.GmtCreate = &v
	return s
}

func (s *AggregateData) SetGmtModified(v string) *AggregateData {
	s.GmtModified = &v
	return s
}

func (s *AggregateData) SetDataItem(v string) *AggregateData {
	s.DataItem = &v
	return s
}

func (s *AggregateData) SetDeviceIdList(v []*string) *AggregateData {
	s.DeviceIdList = v
	return s
}

func (s *AggregateData) SetAggregateModeList(v []*string) *AggregateData {
	s.AggregateModeList = v
	return s
}

func (s *AggregateData) SetAggregateDataName(v string) *AggregateData {
	s.AggregateDataName = &v
	return s
}

func (s *AggregateData) SetAggregateDataDescription(v string) *AggregateData {
	s.AggregateDataDescription = &v
	return s
}

func (s *AggregateData) SetIsAllDevice(v int32) *AggregateData {
	s.IsAllDevice = &v
	return s
}

func (s *AggregateData) SetMonitorItemId(v string) *AggregateData {
	s.MonitorItemId = &v
	return s
}

func (s *AggregateData) SetAggregateDataId(v string) *AggregateData {
	s.AggregateDataId = &v
	return s
}

type SpaceModel struct {
	// 资源名称
	SpaceModelName *string `json:"SpaceModelName,omitempty" xml:"SpaceModelName,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 资源一级ID
	SpaceModelId *string `json:"SpaceModelId,omitempty" xml:"SpaceModelId,omitempty"`
	// 物理空间类型
	SpaceType *string `json:"SpaceType,omitempty" xml:"SpaceType,omitempty"`
	// 层级
	Sort *SpaceModelSort `json:"Sort,omitempty" xml:"Sort,omitempty" type:"Struct"`
	// 物理空间模型
	SpaceModel *SpaceModelSpaceModel `json:"SpaceModel,omitempty" xml:"SpaceModel,omitempty" type:"Struct"`
	// 物理空间id
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// 物理空间实例
	Instance *string `json:"Instance,omitempty" xml:"Instance,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 操作类型
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
}

func (s SpaceModel) String() string {
	return tea.Prettify(s)
}

func (s SpaceModel) GoString() string {
	return s.String()
}

func (s *SpaceModel) SetSpaceModelName(v string) *SpaceModel {
	s.SpaceModelName = &v
	return s
}

func (s *SpaceModel) SetCreateTime(v string) *SpaceModel {
	s.CreateTime = &v
	return s
}

func (s *SpaceModel) SetSpaceModelId(v string) *SpaceModel {
	s.SpaceModelId = &v
	return s
}

func (s *SpaceModel) SetSpaceType(v string) *SpaceModel {
	s.SpaceType = &v
	return s
}

func (s *SpaceModel) SetSort(v *SpaceModelSort) *SpaceModel {
	s.Sort = v
	return s
}

func (s *SpaceModel) SetSpaceModel(v *SpaceModelSpaceModel) *SpaceModel {
	s.SpaceModel = v
	return s
}

func (s *SpaceModel) SetSpaceId(v string) *SpaceModel {
	s.SpaceId = &v
	return s
}

func (s *SpaceModel) SetInstance(v string) *SpaceModel {
	s.Instance = &v
	return s
}

func (s *SpaceModel) SetStatus(v string) *SpaceModel {
	s.Status = &v
	return s
}

func (s *SpaceModel) SetOperateType(v string) *SpaceModel {
	s.OperateType = &v
	return s
}

type SpaceModelSort struct {
	// 分层名称
	LevelName *string `json:"LevelName,omitempty" xml:"LevelName,omitempty"`
	// 层次
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s SpaceModelSort) String() string {
	return tea.Prettify(s)
}

func (s SpaceModelSort) GoString() string {
	return s.String()
}

func (s *SpaceModelSort) SetLevelName(v string) *SpaceModelSort {
	s.LevelName = &v
	return s
}

func (s *SpaceModelSort) SetLevel(v int64) *SpaceModelSort {
	s.Level = &v
	return s
}

type SpaceModelSpaceModel struct {
	// 物理空间模型id
	SpaceModelId *string `json:"SpaceModelId,omitempty" xml:"SpaceModelId,omitempty"`
	// 模型状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 物理空间类型
	SpaceType *string `json:"SpaceType,omitempty" xml:"SpaceType,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 模型实例
	Sort *SpaceModelSpaceModelSort `json:"Sort,omitempty" xml:"Sort,omitempty" type:"Struct"`
}

func (s SpaceModelSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s SpaceModelSpaceModel) GoString() string {
	return s.String()
}

func (s *SpaceModelSpaceModel) SetSpaceModelId(v string) *SpaceModelSpaceModel {
	s.SpaceModelId = &v
	return s
}

func (s *SpaceModelSpaceModel) SetStatus(v string) *SpaceModelSpaceModel {
	s.Status = &v
	return s
}

func (s *SpaceModelSpaceModel) SetSpaceType(v string) *SpaceModelSpaceModel {
	s.SpaceType = &v
	return s
}

func (s *SpaceModelSpaceModel) SetCreateTime(v string) *SpaceModelSpaceModel {
	s.CreateTime = &v
	return s
}

func (s *SpaceModelSpaceModel) SetUpdateTime(v string) *SpaceModelSpaceModel {
	s.UpdateTime = &v
	return s
}

func (s *SpaceModelSpaceModel) SetSort(v *SpaceModelSpaceModelSort) *SpaceModelSpaceModel {
	s.Sort = v
	return s
}

type SpaceModelSpaceModelSort struct {
	// 层级
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// 层级名称
	LevleName *string `json:"LevleName,omitempty" xml:"LevleName,omitempty"`
}

func (s SpaceModelSpaceModelSort) String() string {
	return tea.Prettify(s)
}

func (s SpaceModelSpaceModelSort) GoString() string {
	return s.String()
}

func (s *SpaceModelSpaceModelSort) SetLevel(v int64) *SpaceModelSpaceModelSort {
	s.Level = &v
	return s
}

func (s *SpaceModelSpaceModelSort) SetLevleName(v string) *SpaceModelSpaceModelSort {
	s.LevleName = &v
	return s
}

type Port struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 创建时间
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 修改时间
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 端口集ID
	PortCollectionId *string `json:"PortCollectionId,omitempty" xml:"PortCollectionId,omitempty"`
	// 端口名称
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
}

func (s Port) String() string {
	return tea.Prettify(s)
}

func (s Port) GoString() string {
	return s.String()
}

func (s *Port) SetDeviceId(v string) *Port {
	s.DeviceId = &v
	return s
}

func (s *Port) SetGmtCreate(v string) *Port {
	s.GmtCreate = &v
	return s
}

func (s *Port) SetGmtModified(v string) *Port {
	s.GmtModified = &v
	return s
}

func (s *Port) SetPortCollectionId(v string) *Port {
	s.PortCollectionId = &v
	return s
}

func (s *Port) SetPortName(v string) *Port {
	s.PortName = &v
	return s
}

type BusinessType struct {
	// 配置规范对象
	BusinessTypeId *string `json:"BusinessTypeId,omitempty" xml:"BusinessTypeId,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 业务类型名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 业务类型缩写
	Abbr *string `json:"Abbr,omitempty" xml:"Abbr,omitempty"`
	// 掩码
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// 网关地址位置，正数为正数序号，负数为倒数序号
	Gateway *int64 `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// 是否复用 reuse/single
	Sharing *string `json:"Sharing,omitempty" xml:"Sharing,omitempty"`
	// 分配方向，0表示正向，1表示反向
	Direction *int64 `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// 保留地址数目
	ReserveNumber *int64 `json:"ReserveNumber,omitempty" xml:"ReserveNumber,omitempty"`
	// 业务类型大类
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 绑定的园区类型
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
	// 有效时间
	LeaseTime *string `json:"LeaseTime,omitempty" xml:"LeaseTime,omitempty"`
	// Vlan
	Vlan *string `json:"Vlan,omitempty" xml:"Vlan,omitempty"`
	// 业务类型地址申请完对应的动作，DHCP表示需要触发DHCP变更
	ActionFlag *string `json:"ActionFlag,omitempty" xml:"ActionFlag,omitempty"`
}

func (s BusinessType) String() string {
	return tea.Prettify(s)
}

func (s BusinessType) GoString() string {
	return s.String()
}

func (s *BusinessType) SetBusinessTypeId(v string) *BusinessType {
	s.BusinessTypeId = &v
	return s
}

func (s *BusinessType) SetCreateTime(v string) *BusinessType {
	s.CreateTime = &v
	return s
}

func (s *BusinessType) SetUpdateTime(v string) *BusinessType {
	s.UpdateTime = &v
	return s
}

func (s *BusinessType) SetName(v string) *BusinessType {
	s.Name = &v
	return s
}

func (s *BusinessType) SetAbbr(v string) *BusinessType {
	s.Abbr = &v
	return s
}

func (s *BusinessType) SetMask(v string) *BusinessType {
	s.Mask = &v
	return s
}

func (s *BusinessType) SetGateway(v int64) *BusinessType {
	s.Gateway = &v
	return s
}

func (s *BusinessType) SetSharing(v string) *BusinessType {
	s.Sharing = &v
	return s
}

func (s *BusinessType) SetDirection(v int64) *BusinessType {
	s.Direction = &v
	return s
}

func (s *BusinessType) SetReserveNumber(v int64) *BusinessType {
	s.ReserveNumber = &v
	return s
}

func (s *BusinessType) SetType(v string) *BusinessType {
	s.Type = &v
	return s
}

func (s *BusinessType) SetZoneType(v string) *BusinessType {
	s.ZoneType = &v
	return s
}

func (s *BusinessType) SetLeaseTime(v string) *BusinessType {
	s.LeaseTime = &v
	return s
}

func (s *BusinessType) SetVlan(v string) *BusinessType {
	s.Vlan = &v
	return s
}

func (s *BusinessType) SetActionFlag(v string) *BusinessType {
	s.ActionFlag = &v
	return s
}

type NotificationGroup struct {
	// 创建时间
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 修改时间
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 描述
	NotificationGroupDescription *string `json:"NotificationGroupDescription,omitempty" xml:"NotificationGroupDescription,omitempty"`
	// 通知组ID
	NotificationGroupId *string `json:"NotificationGroupId,omitempty" xml:"NotificationGroupId,omitempty"`
	// 通知组名称
	NotificationGroupName *string `json:"NotificationGroupName,omitempty" xml:"NotificationGroupName,omitempty"`
	// 通知组类型
	NotificationGroupType *string `json:"NotificationGroupType,omitempty" xml:"NotificationGroupType,omitempty"`
	// 钉钉群webhook
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s NotificationGroup) String() string {
	return tea.Prettify(s)
}

func (s NotificationGroup) GoString() string {
	return s.String()
}

func (s *NotificationGroup) SetGmtCreate(v string) *NotificationGroup {
	s.GmtCreate = &v
	return s
}

func (s *NotificationGroup) SetGmtModified(v string) *NotificationGroup {
	s.GmtModified = &v
	return s
}

func (s *NotificationGroup) SetNotificationGroupDescription(v string) *NotificationGroup {
	s.NotificationGroupDescription = &v
	return s
}

func (s *NotificationGroup) SetNotificationGroupId(v string) *NotificationGroup {
	s.NotificationGroupId = &v
	return s
}

func (s *NotificationGroup) SetNotificationGroupName(v string) *NotificationGroup {
	s.NotificationGroupName = &v
	return s
}

func (s *NotificationGroup) SetNotificationGroupType(v string) *NotificationGroup {
	s.NotificationGroupType = &v
	return s
}

func (s *NotificationGroup) SetWebhook(v string) *NotificationGroup {
	s.Webhook = &v
	return s
}

type Template struct {
	// 模板类型
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 模板说明
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// 创建时间
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 修改时间
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// 模板入参
	Input []*TemplateInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Repeated"`
	// 模板出参
	Output []*TemplateOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Repeated"`
	// 模板名称
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 模板类型
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s Template) String() string {
	return tea.Prettify(s)
}

func (s Template) GoString() string {
	return s.String()
}

func (s *Template) SetCategory(v string) *Template {
	s.Category = &v
	return s
}

func (s *Template) SetComment(v string) *Template {
	s.Comment = &v
	return s
}

func (s *Template) SetGmtCreate(v string) *Template {
	s.GmtCreate = &v
	return s
}

func (s *Template) SetGmtModify(v string) *Template {
	s.GmtModify = &v
	return s
}

func (s *Template) SetInput(v []*TemplateInput) *Template {
	s.Input = v
	return s
}

func (s *Template) SetOutput(v []*TemplateOutput) *Template {
	s.Output = v
	return s
}

func (s *Template) SetTemplateName(v string) *Template {
	s.TemplateName = &v
	return s
}

func (s *Template) SetTemplateType(v string) *Template {
	s.TemplateType = &v
	return s
}

type TemplateInput struct {
	// 参数说明
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 参数名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 参数示例
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// 参数类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s TemplateInput) String() string {
	return tea.Prettify(s)
}

func (s TemplateInput) GoString() string {
	return s.String()
}

func (s *TemplateInput) SetDescription(v string) *TemplateInput {
	s.Description = &v
	return s
}

func (s *TemplateInput) SetName(v string) *TemplateInput {
	s.Name = &v
	return s
}

func (s *TemplateInput) SetSample(v string) *TemplateInput {
	s.Sample = &v
	return s
}

func (s *TemplateInput) SetType(v string) *TemplateInput {
	s.Type = &v
	return s
}

type TemplateOutput struct {
	// 参数说明
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 参数名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 参数示例
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// 参数类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s TemplateOutput) String() string {
	return tea.Prettify(s)
}

func (s TemplateOutput) GoString() string {
	return s.String()
}

func (s *TemplateOutput) SetDescription(v string) *TemplateOutput {
	s.Description = &v
	return s
}

func (s *TemplateOutput) SetName(v string) *TemplateOutput {
	s.Name = &v
	return s
}

func (s *TemplateOutput) SetSample(v string) *TemplateOutput {
	s.Sample = &v
	return s
}

func (s *TemplateOutput) SetType(v string) *TemplateOutput {
	s.Type = &v
	return s
}

type DedicatedLine struct {
	// 宽带（Mbps）
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// 关联设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 关联设备端口
	DevicePort *string `json:"DevicePort,omitempty" xml:"DevicePort,omitempty"`
	// 运营商
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// 专线网关
	LineGateway *string `json:"LineGateway,omitempty" xml:"LineGateway,omitempty"`
	// 物理空间专线ID
	LineId *string `json:"LineId,omitempty" xml:"LineId,omitempty"`
	// 专线IP
	LineIp *string `json:"LineIp,omitempty" xml:"LineIp,omitempty"`
	// 专线角色
	LineRole *string `json:"LineRole,omitempty" xml:"LineRole,omitempty"`
}

func (s DedicatedLine) String() string {
	return tea.Prettify(s)
}

func (s DedicatedLine) GoString() string {
	return s.String()
}

func (s *DedicatedLine) SetBandwidth(v int32) *DedicatedLine {
	s.Bandwidth = &v
	return s
}

func (s *DedicatedLine) SetDeviceId(v string) *DedicatedLine {
	s.DeviceId = &v
	return s
}

func (s *DedicatedLine) SetDevicePort(v string) *DedicatedLine {
	s.DevicePort = &v
	return s
}

func (s *DedicatedLine) SetIsp(v string) *DedicatedLine {
	s.Isp = &v
	return s
}

func (s *DedicatedLine) SetLineGateway(v string) *DedicatedLine {
	s.LineGateway = &v
	return s
}

func (s *DedicatedLine) SetLineId(v string) *DedicatedLine {
	s.LineId = &v
	return s
}

func (s *DedicatedLine) SetLineIp(v string) *DedicatedLine {
	s.LineIp = &v
	return s
}

func (s *DedicatedLine) SetLineRole(v string) *DedicatedLine {
	s.LineRole = &v
	return s
}

type DeviceResource struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 资源一级ID
	DeviceResourceIds []*string `json:"DeviceResourceIds,omitempty" xml:"DeviceResourceIds,omitempty" type:"Repeated"`
	// 设备资源id
	DeviceResourceId *string `json:"DeviceResourceId,omitempty" xml:"DeviceResourceId,omitempty"`
	// 建设项目id
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// 操作类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 更新数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 操作类型
	DownloadType *string `json:"DownloadType,omitempty" xml:"DownloadType,omitempty"`
	// 业务类型
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// list类型
	ListType *string `json:"ListType,omitempty" xml:"ListType,omitempty"`
	// 架构资源ID
	ArchId *string `json:"ArchId,omitempty" xml:"ArchId,omitempty"`
	// 设备资源
	DeviceResource []*DeviceResourceDeviceResource `json:"DeviceResource,omitempty" xml:"DeviceResource,omitempty" type:"Repeated"`
	// ip类型
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// 位置
	NetLocation *string `json:"NetLocation,omitempty" xml:"NetLocation,omitempty"`
	// 业务参数
	BusinessTypeParams *string `json:"BusinessTypeParams,omitempty" xml:"BusinessTypeParams,omitempty"`
	// Loopback口
	LoopbackPort *string `json:"LoopbackPort,omitempty" xml:"LoopbackPort,omitempty"`
	// 业务类型id
	BusinessTypeId *string `json:"BusinessTypeId,omitempty" xml:"BusinessTypeId,omitempty"`
}

func (s DeviceResource) String() string {
	return tea.Prettify(s)
}

func (s DeviceResource) GoString() string {
	return s.String()
}

func (s *DeviceResource) SetCreateTime(v string) *DeviceResource {
	s.CreateTime = &v
	return s
}

func (s *DeviceResource) SetDeviceResourceIds(v []*string) *DeviceResource {
	s.DeviceResourceIds = v
	return s
}

func (s *DeviceResource) SetDeviceResourceId(v string) *DeviceResource {
	s.DeviceResourceId = &v
	return s
}

func (s *DeviceResource) SetSetupProjectId(v string) *DeviceResource {
	s.SetupProjectId = &v
	return s
}

func (s *DeviceResource) SetType(v string) *DeviceResource {
	s.Type = &v
	return s
}

func (s *DeviceResource) SetData(v string) *DeviceResource {
	s.Data = &v
	return s
}

func (s *DeviceResource) SetDownloadType(v string) *DeviceResource {
	s.DownloadType = &v
	return s
}

func (s *DeviceResource) SetBusinessType(v string) *DeviceResource {
	s.BusinessType = &v
	return s
}

func (s *DeviceResource) SetListType(v string) *DeviceResource {
	s.ListType = &v
	return s
}

func (s *DeviceResource) SetArchId(v string) *DeviceResource {
	s.ArchId = &v
	return s
}

func (s *DeviceResource) SetDeviceResource(v []*DeviceResourceDeviceResource) *DeviceResource {
	s.DeviceResource = v
	return s
}

func (s *DeviceResource) SetIpType(v string) *DeviceResource {
	s.IpType = &v
	return s
}

func (s *DeviceResource) SetNetLocation(v string) *DeviceResource {
	s.NetLocation = &v
	return s
}

func (s *DeviceResource) SetBusinessTypeParams(v string) *DeviceResource {
	s.BusinessTypeParams = &v
	return s
}

func (s *DeviceResource) SetLoopbackPort(v string) *DeviceResource {
	s.LoopbackPort = &v
	return s
}

func (s *DeviceResource) SetBusinessTypeId(v string) *DeviceResource {
	s.BusinessTypeId = &v
	return s
}

type DeviceResourceDeviceResource struct {
	// 设备资源ID
	DeviceResourceId *string `json:"DeviceResourceId,omitempty" xml:"DeviceResourceId,omitempty"`
	// 建设项目id
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 组号
	BlockNumber *string `json:"BlockNumber,omitempty" xml:"BlockNumber,omitempty"`
	// 设备号
	DeviceNumber *string `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 模型
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 主机名
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// 设备sn号
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// 物理空间位置
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// 带内管理地址
	ManagerIp *string `json:"ManagerIp,omitempty" xml:"ManagerIp,omitempty"`
	// 交付登录地址
	DeliveryIp *string `json:"DeliveryIp,omitempty" xml:"DeliveryIp,omitempty"`
	// 配置生成
	ConfigGenerate *bool `json:"ConfigGenerate,omitempty" xml:"ConfigGenerate,omitempty"`
	// 配置下发状态
	ConfigTaskStatus *string `json:"ConfigTaskStatus,omitempty" xml:"ConfigTaskStatus,omitempty"`
	// 设备配置
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// loopback地址
	Loopback *string `json:"Loopback,omitempty" xml:"Loopback,omitempty"`
	// 设备互联地址
	InterConnection *string `json:"InterConnection,omitempty" xml:"InterConnection,omitempty"`
	// 设备业务地址
	Business *string `json:"Business,omitempty" xml:"Business,omitempty"`
	// 堆叠状态
	Stack *bool `json:"Stack,omitempty" xml:"Stack,omitempty"`
}

func (s DeviceResourceDeviceResource) String() string {
	return tea.Prettify(s)
}

func (s DeviceResourceDeviceResource) GoString() string {
	return s.String()
}

func (s *DeviceResourceDeviceResource) SetDeviceResourceId(v string) *DeviceResourceDeviceResource {
	s.DeviceResourceId = &v
	return s
}

func (s *DeviceResourceDeviceResource) SetSetupProjectId(v string) *DeviceResourceDeviceResource {
	s.SetupProjectId = &v
	return s
}

func (s *DeviceResourceDeviceResource) SetRole(v string) *DeviceResourceDeviceResource {
	s.Role = &v
	return s
}

func (s *DeviceResourceDeviceResource) SetBlockNumber(v string) *DeviceResourceDeviceResource {
	s.BlockNumber = &v
	return s
}

func (s *DeviceResourceDeviceResource) SetDeviceNumber(v string) *DeviceResourceDeviceResource {
	s.DeviceNumber = &v
	return s
}

func (s *DeviceResourceDeviceResource) SetVendor(v string) *DeviceResourceDeviceResource {
	s.Vendor = &v
	return s
}

func (s *DeviceResourceDeviceResource) SetModel(v string) *DeviceResourceDeviceResource {
	s.Model = &v
	return s
}

func (s *DeviceResourceDeviceResource) SetHostName(v string) *DeviceResourceDeviceResource {
	s.HostName = &v
	return s
}

func (s *DeviceResourceDeviceResource) SetSn(v string) *DeviceResourceDeviceResource {
	s.Sn = &v
	return s
}

func (s *DeviceResourceDeviceResource) SetLocation(v string) *DeviceResourceDeviceResource {
	s.Location = &v
	return s
}

func (s *DeviceResourceDeviceResource) SetManagerIp(v string) *DeviceResourceDeviceResource {
	s.ManagerIp = &v
	return s
}

func (s *DeviceResourceDeviceResource) SetDeliveryIp(v string) *DeviceResourceDeviceResource {
	s.DeliveryIp = &v
	return s
}

func (s *DeviceResourceDeviceResource) SetConfigGenerate(v bool) *DeviceResourceDeviceResource {
	s.ConfigGenerate = &v
	return s
}

func (s *DeviceResourceDeviceResource) SetConfigTaskStatus(v string) *DeviceResourceDeviceResource {
	s.ConfigTaskStatus = &v
	return s
}

func (s *DeviceResourceDeviceResource) SetConfig(v string) *DeviceResourceDeviceResource {
	s.Config = &v
	return s
}

func (s *DeviceResourceDeviceResource) SetLoopback(v string) *DeviceResourceDeviceResource {
	s.Loopback = &v
	return s
}

func (s *DeviceResourceDeviceResource) SetInterConnection(v string) *DeviceResourceDeviceResource {
	s.InterConnection = &v
	return s
}

func (s *DeviceResourceDeviceResource) SetBusiness(v string) *DeviceResourceDeviceResource {
	s.Business = &v
	return s
}

func (s *DeviceResourceDeviceResource) SetStack(v bool) *DeviceResourceDeviceResource {
	s.Stack = &v
	return s
}

type ModelToRole struct {
	// 资源一级ID
	ModelToRoleId *string `json:"ModelToRoleId,omitempty" xml:"ModelToRoleId,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 角色型号对应关系uid
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 架构迭代uid
	NetworkArchitectureIterationId *string `json:"NetworkArchitectureIterationId,omitempty" xml:"NetworkArchitectureIterationId,omitempty"`
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 设备厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 设备型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s ModelToRole) String() string {
	return tea.Prettify(s)
}

func (s ModelToRole) GoString() string {
	return s.String()
}

func (s *ModelToRole) SetModelToRoleId(v string) *ModelToRole {
	s.ModelToRoleId = &v
	return s
}

func (s *ModelToRole) SetCreateTime(v string) *ModelToRole {
	s.CreateTime = &v
	return s
}

func (s *ModelToRole) SetUpdateTime(v string) *ModelToRole {
	s.UpdateTime = &v
	return s
}

func (s *ModelToRole) SetId(v string) *ModelToRole {
	s.Id = &v
	return s
}

func (s *ModelToRole) SetNetworkArchitectureIterationId(v string) *ModelToRole {
	s.NetworkArchitectureIterationId = &v
	return s
}

func (s *ModelToRole) SetRole(v string) *ModelToRole {
	s.Role = &v
	return s
}

func (s *ModelToRole) SetVendor(v string) *ModelToRole {
	s.Vendor = &v
	return s
}

func (s *ModelToRole) SetModel(v string) *ModelToRole {
	s.Model = &v
	return s
}

type ScheduleWorker struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 资源一级ID
	ScheduleWorkerId *string `json:"ScheduleWorkerId,omitempty" xml:"ScheduleWorkerId,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 值班人员工号
	WorkerId *string `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
	// 值班人员姓名
	WorkerName *string `json:"WorkerName,omitempty" xml:"WorkerName,omitempty"`
	// 联系方式
	WorkerContact *string `json:"WorkerContact,omitempty" xml:"WorkerContact,omitempty"`
}

func (s ScheduleWorker) String() string {
	return tea.Prettify(s)
}

func (s ScheduleWorker) GoString() string {
	return s.String()
}

func (s *ScheduleWorker) SetCreateTime(v string) *ScheduleWorker {
	s.CreateTime = &v
	return s
}

func (s *ScheduleWorker) SetScheduleWorkerId(v string) *ScheduleWorker {
	s.ScheduleWorkerId = &v
	return s
}

func (s *ScheduleWorker) SetUpdateTime(v string) *ScheduleWorker {
	s.UpdateTime = &v
	return s
}

func (s *ScheduleWorker) SetWorkerId(v string) *ScheduleWorker {
	s.WorkerId = &v
	return s
}

func (s *ScheduleWorker) SetWorkerName(v string) *ScheduleWorker {
	s.WorkerName = &v
	return s
}

func (s *ScheduleWorker) SetWorkerContact(v string) *ScheduleWorker {
	s.WorkerContact = &v
	return s
}

type IpBlock struct {
	// 资源组ID
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 资源一级ID
	IpBlockId *string `json:"IpBlockId,omitempty" xml:"IpBlockId,omitempty"`
	// IP段地址
	Block *string `json:"Block,omitempty" xml:"Block,omitempty"`
	// IP段掩码
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// 父地址段UID
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// 公网私网标志 PUBLIC PRIVATE
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// 地址类别 IPV4
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 业务类型UID
	BusinessTypeId *string `json:"BusinessTypeId,omitempty" xml:"BusinessTypeId,omitempty"`
	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// 园区名称
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	// 备份设备名称
	BackupDeviceName *string `json:"BackupDeviceName,omitempty" xml:"BackupDeviceName,omitempty"`
	// 公网地址类型 INC GUEST VIP
	NetBusiness *string `json:"NetBusiness,omitempty" xml:"NetBusiness,omitempty"`
	// IP归属
	Ownership *string `json:"Ownership,omitempty" xml:"Ownership,omitempty"`
	// IP用途
	Application *string `json:"Application,omitempty" xml:"Application,omitempty"`
	// 备注
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 是否保留父段 true 是 false 否
	ReserveParentBlock *string `json:"ReserveParentBlock,omitempty" xml:"ReserveParentBlock,omitempty"`
	// 更新类型 update 更新 split 拆分
	UpdateType *string `json:"UpdateType,omitempty" xml:"UpdateType,omitempty"`
	// 园区层级
	ZoneLayer []*IpBlockZoneLayer `json:"ZoneLayer,omitempty" xml:"ZoneLayer,omitempty" type:"Repeated"`
}

func (s IpBlock) String() string {
	return tea.Prettify(s)
}

func (s IpBlock) GoString() string {
	return s.String()
}

func (s *IpBlock) SetResourceGroupId(v string) *IpBlock {
	s.ResourceGroupId = &v
	return s
}

func (s *IpBlock) SetCreateTime(v string) *IpBlock {
	s.CreateTime = &v
	return s
}

func (s *IpBlock) SetIpBlockId(v string) *IpBlock {
	s.IpBlockId = &v
	return s
}

func (s *IpBlock) SetBlock(v string) *IpBlock {
	s.Block = &v
	return s
}

func (s *IpBlock) SetMask(v string) *IpBlock {
	s.Mask = &v
	return s
}

func (s *IpBlock) SetParentId(v string) *IpBlock {
	s.ParentId = &v
	return s
}

func (s *IpBlock) SetNetType(v string) *IpBlock {
	s.NetType = &v
	return s
}

func (s *IpBlock) SetCategory(v string) *IpBlock {
	s.Category = &v
	return s
}

func (s *IpBlock) SetBusinessTypeId(v string) *IpBlock {
	s.BusinessTypeId = &v
	return s
}

func (s *IpBlock) SetDeviceName(v string) *IpBlock {
	s.DeviceName = &v
	return s
}

func (s *IpBlock) SetZoneName(v string) *IpBlock {
	s.ZoneName = &v
	return s
}

func (s *IpBlock) SetBackupDeviceName(v string) *IpBlock {
	s.BackupDeviceName = &v
	return s
}

func (s *IpBlock) SetNetBusiness(v string) *IpBlock {
	s.NetBusiness = &v
	return s
}

func (s *IpBlock) SetOwnership(v string) *IpBlock {
	s.Ownership = &v
	return s
}

func (s *IpBlock) SetApplication(v string) *IpBlock {
	s.Application = &v
	return s
}

func (s *IpBlock) SetDescription(v string) *IpBlock {
	s.Description = &v
	return s
}

func (s *IpBlock) SetReserveParentBlock(v string) *IpBlock {
	s.ReserveParentBlock = &v
	return s
}

func (s *IpBlock) SetUpdateType(v string) *IpBlock {
	s.UpdateType = &v
	return s
}

func (s *IpBlock) SetZoneLayer(v []*IpBlockZoneLayer) *IpBlock {
	s.ZoneLayer = v
	return s
}

type IpBlockZoneLayer struct {
	// 园区层级名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 园区层级值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s IpBlockZoneLayer) String() string {
	return tea.Prettify(s)
}

func (s IpBlockZoneLayer) GoString() string {
	return s.String()
}

func (s *IpBlockZoneLayer) SetName(v string) *IpBlockZoneLayer {
	s.Name = &v
	return s
}

func (s *IpBlockZoneLayer) SetValue(v string) *IpBlockZoneLayer {
	s.Value = &v
	return s
}

type Order struct {
	// 创建时间
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 修改时间
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// 工单id
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// 工单返回
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// 工单参数
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// 方案id
	SchemeId *string `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	// 方案名
	SchemeName *string `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
	// 任务状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 工单标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s Order) String() string {
	return tea.Prettify(s)
}

func (s Order) GoString() string {
	return s.String()
}

func (s *Order) SetGmtCreate(v string) *Order {
	s.GmtCreate = &v
	return s
}

func (s *Order) SetGmtModify(v string) *Order {
	s.GmtModify = &v
	return s
}

func (s *Order) SetOrderId(v string) *Order {
	s.OrderId = &v
	return s
}

func (s *Order) SetOutput(v string) *Order {
	s.Output = &v
	return s
}

func (s *Order) SetParams(v string) *Order {
	s.Params = &v
	return s
}

func (s *Order) SetSchemeId(v string) *Order {
	s.SchemeId = &v
	return s
}

func (s *Order) SetSchemeName(v string) *Order {
	s.SchemeName = &v
	return s
}

func (s *Order) SetStatus(v string) *Order {
	s.Status = &v
	return s
}

func (s *Order) SetTitle(v string) *Order {
	s.Title = &v
	return s
}

type Script struct {
	// 脚本代码
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 创建时间
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 修改时间
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// 模板入参
	Input []*ScriptInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Repeated"`
	// 模板出参
	Output []*ScriptOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Repeated"`
	// 规则列表
	Rules []*ScriptRule `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// 脚本id
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// 模板id
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 版本id
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s Script) String() string {
	return tea.Prettify(s)
}

func (s Script) GoString() string {
	return s.String()
}

func (s *Script) SetContent(v string) *Script {
	s.Content = &v
	return s
}

func (s *Script) SetGmtCreate(v string) *Script {
	s.GmtCreate = &v
	return s
}

func (s *Script) SetGmtModify(v string) *Script {
	s.GmtModify = &v
	return s
}

func (s *Script) SetInput(v []*ScriptInput) *Script {
	s.Input = v
	return s
}

func (s *Script) SetOutput(v []*ScriptOutput) *Script {
	s.Output = v
	return s
}

func (s *Script) SetRules(v []*ScriptRule) *Script {
	s.Rules = v
	return s
}

func (s *Script) SetScriptId(v string) *Script {
	s.ScriptId = &v
	return s
}

func (s *Script) SetTemplateId(v string) *Script {
	s.TemplateId = &v
	return s
}

func (s *Script) SetVersionId(v string) *Script {
	s.VersionId = &v
	return s
}

type ScriptInput struct {
	// 参数说明
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 参数名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 参数示例
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// 参数类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ScriptInput) String() string {
	return tea.Prettify(s)
}

func (s ScriptInput) GoString() string {
	return s.String()
}

func (s *ScriptInput) SetDescription(v string) *ScriptInput {
	s.Description = &v
	return s
}

func (s *ScriptInput) SetName(v string) *ScriptInput {
	s.Name = &v
	return s
}

func (s *ScriptInput) SetSample(v string) *ScriptInput {
	s.Sample = &v
	return s
}

func (s *ScriptInput) SetType(v string) *ScriptInput {
	s.Type = &v
	return s
}

type ScriptOutput struct {
	// 参数说明
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 参数名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 参数示例
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// 参数类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ScriptOutput) String() string {
	return tea.Prettify(s)
}

func (s ScriptOutput) GoString() string {
	return s.String()
}

func (s *ScriptOutput) SetDescription(v string) *ScriptOutput {
	s.Description = &v
	return s
}

func (s *ScriptOutput) SetName(v string) *ScriptOutput {
	s.Name = &v
	return s
}

func (s *ScriptOutput) SetSample(v string) *ScriptOutput {
	s.Sample = &v
	return s
}

func (s *ScriptOutput) SetType(v string) *ScriptOutput {
	s.Type = &v
	return s
}

type Agent struct {
	// 探针Id
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 探针名称
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// 探针版本
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// cpu使用率
	CpuUsage *string `json:"CpuUsage,omitempty" xml:"CpuUsage,omitempty"`
	// 磁盘利用率
	DiskUsage *string `json:"DiskUsage,omitempty" xml:"DiskUsage,omitempty"`
	// 更新时间
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// 探针IP
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// 系统版本
	KernelVersion *string `json:"KernelVersion,omitempty" xml:"KernelVersion,omitempty"`
	// 内存使用率
	MemoryUsage *string `json:"MemoryUsage,omitempty" xml:"MemoryUsage,omitempty"`
	// 安全域
	SecurityDomain *string `json:"SecurityDomain,omitempty" xml:"SecurityDomain,omitempty"`
	// 探针状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s Agent) String() string {
	return tea.Prettify(s)
}

func (s Agent) GoString() string {
	return s.String()
}

func (s *Agent) SetAgentId(v string) *Agent {
	s.AgentId = &v
	return s
}

func (s *Agent) SetAgentName(v string) *Agent {
	s.AgentName = &v
	return s
}

func (s *Agent) SetAgentVersion(v string) *Agent {
	s.AgentVersion = &v
	return s
}

func (s *Agent) SetCpuUsage(v string) *Agent {
	s.CpuUsage = &v
	return s
}

func (s *Agent) SetDiskUsage(v string) *Agent {
	s.DiskUsage = &v
	return s
}

func (s *Agent) SetGmtModify(v string) *Agent {
	s.GmtModify = &v
	return s
}

func (s *Agent) SetIp(v string) *Agent {
	s.Ip = &v
	return s
}

func (s *Agent) SetKernelVersion(v string) *Agent {
	s.KernelVersion = &v
	return s
}

func (s *Agent) SetMemoryUsage(v string) *Agent {
	s.MemoryUsage = &v
	return s
}

func (s *Agent) SetSecurityDomain(v string) *Agent {
	s.SecurityDomain = &v
	return s
}

func (s *Agent) SetStatus(v string) *Agent {
	s.Status = &v
	return s
}

type ResourceInformation struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 资源一级ID
	ResourceInformationId *string `json:"ResourceInformationId,omitempty" xml:"ResourceInformationId,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 资源属性
	ResourceAttribute *string `json:"ResourceAttribute,omitempty" xml:"ResourceAttribute,omitempty"`
	// 建设项目资源id
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// 信息
	Information []*ResourceInformationInformation `json:"Information,omitempty" xml:"Information,omitempty" type:"Repeated"`
	// 架构id
	ArchitectureId *string `json:"ArchitectureId,omitempty" xml:"ArchitectureId,omitempty"`
}

func (s ResourceInformation) String() string {
	return tea.Prettify(s)
}

func (s ResourceInformation) GoString() string {
	return s.String()
}

func (s *ResourceInformation) SetCreateTime(v string) *ResourceInformation {
	s.CreateTime = &v
	return s
}

func (s *ResourceInformation) SetResourceInformationId(v string) *ResourceInformation {
	s.ResourceInformationId = &v
	return s
}

func (s *ResourceInformation) SetUpdateTime(v string) *ResourceInformation {
	s.UpdateTime = &v
	return s
}

func (s *ResourceInformation) SetResourceType(v string) *ResourceInformation {
	s.ResourceType = &v
	return s
}

func (s *ResourceInformation) SetResourceAttribute(v string) *ResourceInformation {
	s.ResourceAttribute = &v
	return s
}

func (s *ResourceInformation) SetSetupProjectId(v string) *ResourceInformation {
	s.SetupProjectId = &v
	return s
}

func (s *ResourceInformation) SetInformation(v []*ResourceInformationInformation) *ResourceInformation {
	s.Information = v
	return s
}

func (s *ResourceInformation) SetArchitectureId(v string) *ResourceInformation {
	s.ArchitectureId = &v
	return s
}

type ResourceInformationInformation struct {
	// 键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 键属性
	KeyAttribute *string `json:"KeyAttribute,omitempty" xml:"KeyAttribute,omitempty"`
	// 键动作
	KeyAction *string `json:"KeyAction,omitempty" xml:"KeyAction,omitempty"`
	// 键描述
	KeyDescription *string `json:"KeyDescription,omitempty" xml:"KeyDescription,omitempty"`
}

func (s ResourceInformationInformation) String() string {
	return tea.Prettify(s)
}

func (s ResourceInformationInformation) GoString() string {
	return s.String()
}

func (s *ResourceInformationInformation) SetKey(v string) *ResourceInformationInformation {
	s.Key = &v
	return s
}

func (s *ResourceInformationInformation) SetKeyAttribute(v string) *ResourceInformationInformation {
	s.KeyAttribute = &v
	return s
}

func (s *ResourceInformationInformation) SetKeyAction(v string) *ResourceInformationInformation {
	s.KeyAction = &v
	return s
}

func (s *ResourceInformationInformation) SetKeyDescription(v string) *ResourceInformationInformation {
	s.KeyDescription = &v
	return s
}

type ScriptRule struct {
	// 设备架构
	Arch *string `json:"Arch,omitempty" xml:"Arch,omitempty"`
	// 设备安全域
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 设备型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 设备OS版本
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// 设备角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 规则id
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// 脚本id
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// 设备厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ScriptRule) String() string {
	return tea.Prettify(s)
}

func (s ScriptRule) GoString() string {
	return s.String()
}

func (s *ScriptRule) SetArch(v string) *ScriptRule {
	s.Arch = &v
	return s
}

func (s *ScriptRule) SetDomain(v string) *ScriptRule {
	s.Domain = &v
	return s
}

func (s *ScriptRule) SetModel(v string) *ScriptRule {
	s.Model = &v
	return s
}

func (s *ScriptRule) SetOs(v string) *ScriptRule {
	s.Os = &v
	return s
}

func (s *ScriptRule) SetRole(v string) *ScriptRule {
	s.Role = &v
	return s
}

func (s *ScriptRule) SetRuleId(v string) *ScriptRule {
	s.RuleId = &v
	return s
}

func (s *ScriptRule) SetScriptId(v string) *ScriptRule {
	s.ScriptId = &v
	return s
}

func (s *ScriptRule) SetVendor(v string) *ScriptRule {
	s.Vendor = &v
	return s
}

type ScheduleDuty struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 资源一级ID
	ScheduleDutyId *string `json:"ScheduleDutyId,omitempty" xml:"ScheduleDutyId,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 值班表日期
	WorkDate *string `json:"WorkDate,omitempty" xml:"WorkDate,omitempty"`
	// 值班表类型
	WorkType *string `json:"WorkType,omitempty" xml:"WorkType,omitempty"`
	// 值班人员姓名
	WorkerName *string `json:"WorkerName,omitempty" xml:"WorkerName,omitempty"`
	// 值班人员工号
	WorkerId *string `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
	// 开始时间
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束时间
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// scheduleTypeIds
	ScheduleTypeIds []*string `json:"ScheduleTypeIds,omitempty" xml:"ScheduleTypeIds,omitempty" type:"Repeated"`
	// typeWorkerList
	TypeWorkerList []*ScheduleDutyTypeWorkerList `json:"TypeWorkerList,omitempty" xml:"TypeWorkerList,omitempty" type:"Repeated"`
}

func (s ScheduleDuty) String() string {
	return tea.Prettify(s)
}

func (s ScheduleDuty) GoString() string {
	return s.String()
}

func (s *ScheduleDuty) SetCreateTime(v string) *ScheduleDuty {
	s.CreateTime = &v
	return s
}

func (s *ScheduleDuty) SetScheduleDutyId(v string) *ScheduleDuty {
	s.ScheduleDutyId = &v
	return s
}

func (s *ScheduleDuty) SetUpdateTime(v string) *ScheduleDuty {
	s.UpdateTime = &v
	return s
}

func (s *ScheduleDuty) SetWorkDate(v string) *ScheduleDuty {
	s.WorkDate = &v
	return s
}

func (s *ScheduleDuty) SetWorkType(v string) *ScheduleDuty {
	s.WorkType = &v
	return s
}

func (s *ScheduleDuty) SetWorkerName(v string) *ScheduleDuty {
	s.WorkerName = &v
	return s
}

func (s *ScheduleDuty) SetWorkerId(v string) *ScheduleDuty {
	s.WorkerId = &v
	return s
}

func (s *ScheduleDuty) SetStartDate(v string) *ScheduleDuty {
	s.StartDate = &v
	return s
}

func (s *ScheduleDuty) SetEndDate(v string) *ScheduleDuty {
	s.EndDate = &v
	return s
}

func (s *ScheduleDuty) SetScheduleTypeIds(v []*string) *ScheduleDuty {
	s.ScheduleTypeIds = v
	return s
}

func (s *ScheduleDuty) SetTypeWorkerList(v []*ScheduleDutyTypeWorkerList) *ScheduleDuty {
	s.TypeWorkerList = v
	return s
}

type ScheduleDutyTypeWorkerList struct {
	// scheduleTypeId
	ScheduleTypeId *string `json:"ScheduleTypeId,omitempty" xml:"ScheduleTypeId,omitempty"`
	// 值班人员姓名
	ScheduleWorkerName []*string `json:"ScheduleWorkerName,omitempty" xml:"ScheduleWorkerName,omitempty" type:"Repeated"`
}

func (s ScheduleDutyTypeWorkerList) String() string {
	return tea.Prettify(s)
}

func (s ScheduleDutyTypeWorkerList) GoString() string {
	return s.String()
}

func (s *ScheduleDutyTypeWorkerList) SetScheduleTypeId(v string) *ScheduleDutyTypeWorkerList {
	s.ScheduleTypeId = &v
	return s
}

func (s *ScheduleDutyTypeWorkerList) SetScheduleWorkerName(v []*string) *ScheduleDutyTypeWorkerList {
	s.ScheduleWorkerName = v
	return s
}

type Device struct {
	// 账号类型
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// snmp版本号
	AccountVersion *string `json:"AccountVersion,omitempty" xml:"AccountVersion,omitempty"`
	// Auth PassPhrase
	AuthPassPhrase *string `json:"AuthPassPhrase,omitempty" xml:"AuthPassPhrase,omitempty"`
	// Auth Protocol
	AuthProtocol *string `json:"AuthProtocol,omitempty" xml:"AuthProtocol,omitempty"`
	// community
	Community *string `json:"Community,omitempty" xml:"Community,omitempty"`
	// 设备形态
	DeviceForm *string `json:"DeviceForm,omitempty" xml:"DeviceForm,omitempty"`
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 设备IP
	DeviceIp *string `json:"DeviceIp,omitempty" xml:"DeviceIp,omitempty"`
	// 设备MAC地址
	DeviceMac *string `json:"DeviceMac,omitempty" xml:"DeviceMac,omitempty"`
	// 设备SN
	DeviceSn *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	// 主机名
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// 设备型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Privacy PassPhrase
	PrivacyPassPhrase *string `json:"PrivacyPassPhrase,omitempty" xml:"PrivacyPassPhrase,omitempty"`
	// Privacy Protocol
	PrivacyProtocol *string `json:"PrivacyProtocol,omitempty" xml:"PrivacyProtocol,omitempty"`
	// 设备安全域
	SecurityDomain *string `json:"SecurityDomain,omitempty" xml:"SecurityDomain,omitempty"`
	// 安全等级
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// 设备所属物理空间
	Space *string `json:"Space,omitempty" xml:"Space,omitempty"`
	// SSH登录账号
	SshAccount *string `json:"SshAccount,omitempty" xml:"SshAccount,omitempty"`
	// SSH登录密码
	SshPassword *string `json:"SshPassword,omitempty" xml:"SshPassword,omitempty"`
	// 设备状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// TELNET登录账号
	TelnetAccount *string `json:"TelnetAccount,omitempty" xml:"TelnetAccount,omitempty"`
	// TELNET登录密码
	TelnetPassword *string `json:"TelnetPassword,omitempty" xml:"TelnetPassword,omitempty"`
	// 用户名
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// 设备厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s Device) String() string {
	return tea.Prettify(s)
}

func (s Device) GoString() string {
	return s.String()
}

func (s *Device) SetAccountType(v string) *Device {
	s.AccountType = &v
	return s
}

func (s *Device) SetAccountVersion(v string) *Device {
	s.AccountVersion = &v
	return s
}

func (s *Device) SetAuthPassPhrase(v string) *Device {
	s.AuthPassPhrase = &v
	return s
}

func (s *Device) SetAuthProtocol(v string) *Device {
	s.AuthProtocol = &v
	return s
}

func (s *Device) SetCommunity(v string) *Device {
	s.Community = &v
	return s
}

func (s *Device) SetDeviceForm(v string) *Device {
	s.DeviceForm = &v
	return s
}

func (s *Device) SetDeviceId(v string) *Device {
	s.DeviceId = &v
	return s
}

func (s *Device) SetDeviceIp(v string) *Device {
	s.DeviceIp = &v
	return s
}

func (s *Device) SetDeviceMac(v string) *Device {
	s.DeviceMac = &v
	return s
}

func (s *Device) SetDeviceSn(v string) *Device {
	s.DeviceSn = &v
	return s
}

func (s *Device) SetHostname(v string) *Device {
	s.Hostname = &v
	return s
}

func (s *Device) SetModel(v string) *Device {
	s.Model = &v
	return s
}

func (s *Device) SetPrivacyPassPhrase(v string) *Device {
	s.PrivacyPassPhrase = &v
	return s
}

func (s *Device) SetPrivacyProtocol(v string) *Device {
	s.PrivacyProtocol = &v
	return s
}

func (s *Device) SetSecurityDomain(v string) *Device {
	s.SecurityDomain = &v
	return s
}

func (s *Device) SetSecurityLevel(v string) *Device {
	s.SecurityLevel = &v
	return s
}

func (s *Device) SetSpace(v string) *Device {
	s.Space = &v
	return s
}

func (s *Device) SetSshAccount(v string) *Device {
	s.SshAccount = &v
	return s
}

func (s *Device) SetSshPassword(v string) *Device {
	s.SshPassword = &v
	return s
}

func (s *Device) SetStatus(v string) *Device {
	s.Status = &v
	return s
}

func (s *Device) SetTelnetAccount(v string) *Device {
	s.TelnetAccount = &v
	return s
}

func (s *Device) SetTelnetPassword(v string) *Device {
	s.TelnetPassword = &v
	return s
}

func (s *Device) SetUserName(v string) *Device {
	s.UserName = &v
	return s
}

func (s *Device) SetVendor(v string) *Device {
	s.Vendor = &v
	return s
}

type TimePeriod struct {
	// 创建时间
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 修改时间
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 描述
	TimePeriodDescription *string `json:"TimePeriodDescription,omitempty" xml:"TimePeriodDescription,omitempty"`
	// 时间段名称
	TimePeriodName *string `json:"TimePeriodName,omitempty" xml:"TimePeriodName,omitempty"`
	// 时间段ID
	TimePeriodId *string `json:"TimePeriodId,omitempty" xml:"TimePeriodId,omitempty"`
	// Cron表达式
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// 来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s TimePeriod) String() string {
	return tea.Prettify(s)
}

func (s TimePeriod) GoString() string {
	return s.String()
}

func (s *TimePeriod) SetGmtCreate(v string) *TimePeriod {
	s.GmtCreate = &v
	return s
}

func (s *TimePeriod) SetGmtModified(v string) *TimePeriod {
	s.GmtModified = &v
	return s
}

func (s *TimePeriod) SetTimePeriodDescription(v string) *TimePeriod {
	s.TimePeriodDescription = &v
	return s
}

func (s *TimePeriod) SetTimePeriodName(v string) *TimePeriod {
	s.TimePeriodName = &v
	return s
}

func (s *TimePeriod) SetTimePeriodId(v string) *TimePeriod {
	s.TimePeriodId = &v
	return s
}

func (s *TimePeriod) SetCronExpression(v string) *TimePeriod {
	s.CronExpression = &v
	return s
}

func (s *TimePeriod) SetSource(v string) *TimePeriod {
	s.Source = &v
	return s
}

type AtomicStep struct {
	// 步骤说明
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 步骤入参
	Input []*AtomicStepInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Repeated"`
	// 步骤出参
	Output []*AtomicStepOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Repeated"`
	// 步骤id
	StepId *string `json:"StepId,omitempty" xml:"StepId,omitempty"`
	// 步骤名称
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// 步骤类型
	StepType *string `json:"StepType,omitempty" xml:"StepType,omitempty"`
}

func (s AtomicStep) String() string {
	return tea.Prettify(s)
}

func (s AtomicStep) GoString() string {
	return s.String()
}

func (s *AtomicStep) SetDescription(v string) *AtomicStep {
	s.Description = &v
	return s
}

func (s *AtomicStep) SetInput(v []*AtomicStepInput) *AtomicStep {
	s.Input = v
	return s
}

func (s *AtomicStep) SetOutput(v []*AtomicStepOutput) *AtomicStep {
	s.Output = v
	return s
}

func (s *AtomicStep) SetStepId(v string) *AtomicStep {
	s.StepId = &v
	return s
}

func (s *AtomicStep) SetStepName(v string) *AtomicStep {
	s.StepName = &v
	return s
}

func (s *AtomicStep) SetStepType(v string) *AtomicStep {
	s.StepType = &v
	return s
}

type AtomicStepInput struct {
	// 参数说明
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 参数名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 参数示例
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// 参数类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AtomicStepInput) String() string {
	return tea.Prettify(s)
}

func (s AtomicStepInput) GoString() string {
	return s.String()
}

func (s *AtomicStepInput) SetDescription(v string) *AtomicStepInput {
	s.Description = &v
	return s
}

func (s *AtomicStepInput) SetName(v string) *AtomicStepInput {
	s.Name = &v
	return s
}

func (s *AtomicStepInput) SetSample(v string) *AtomicStepInput {
	s.Sample = &v
	return s
}

func (s *AtomicStepInput) SetType(v string) *AtomicStepInput {
	s.Type = &v
	return s
}

type AtomicStepOutput struct {
	// 参数说明
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 参数名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 参数示例
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// 参数类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AtomicStepOutput) String() string {
	return tea.Prettify(s)
}

func (s AtomicStepOutput) GoString() string {
	return s.String()
}

func (s *AtomicStepOutput) SetDescription(v string) *AtomicStepOutput {
	s.Description = &v
	return s
}

func (s *AtomicStepOutput) SetName(v string) *AtomicStepOutput {
	s.Name = &v
	return s
}

func (s *AtomicStepOutput) SetSample(v string) *AtomicStepOutput {
	s.Sample = &v
	return s
}

func (s *AtomicStepOutput) SetType(v string) *AtomicStepOutput {
	s.Type = &v
	return s
}

type InspectionItem struct {
	// 巡检项ID
	ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// 巡检项名字
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// 巡检项描述
	ItemDescription *string `json:"ItemDescription,omitempty" xml:"ItemDescription,omitempty"`
	// 巡检定时表达式
	InspectionCrontab *string `json:"InspectionCrontab,omitempty" xml:"InspectionCrontab,omitempty"`
}

func (s InspectionItem) String() string {
	return tea.Prettify(s)
}

func (s InspectionItem) GoString() string {
	return s.String()
}

func (s *InspectionItem) SetItemId(v string) *InspectionItem {
	s.ItemId = &v
	return s
}

func (s *InspectionItem) SetItemName(v string) *InspectionItem {
	s.ItemName = &v
	return s
}

func (s *InspectionItem) SetItemDescription(v string) *InspectionItem {
	s.ItemDescription = &v
	return s
}

func (s *InspectionItem) SetInspectionCrontab(v string) *InspectionItem {
	s.InspectionCrontab = &v
	return s
}

type DeviceProperty struct {
	// 属性值
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 设备形态
	DeviceForm *string `json:"DeviceForm,omitempty" xml:"DeviceForm,omitempty"`
	// 属性格式，包括JSON和SPLITTER（分隔符）
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// 属性展示名称
	NameCn *string `json:"NameCn,omitempty" xml:"NameCn,omitempty"`
	// 属性英文主键
	NameEn *string `json:"NameEn,omitempty" xml:"NameEn,omitempty"`
	// 设备属性ID
	PropertyId *string `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
}

func (s DeviceProperty) String() string {
	return tea.Prettify(s)
}

func (s DeviceProperty) GoString() string {
	return s.String()
}

func (s *DeviceProperty) SetContent(v string) *DeviceProperty {
	s.Content = &v
	return s
}

func (s *DeviceProperty) SetDeviceForm(v string) *DeviceProperty {
	s.DeviceForm = &v
	return s
}

func (s *DeviceProperty) SetFormat(v string) *DeviceProperty {
	s.Format = &v
	return s
}

func (s *DeviceProperty) SetNameCn(v string) *DeviceProperty {
	s.NameCn = &v
	return s
}

func (s *DeviceProperty) SetNameEn(v string) *DeviceProperty {
	s.NameEn = &v
	return s
}

func (s *DeviceProperty) SetPropertyId(v string) *DeviceProperty {
	s.PropertyId = &v
	return s
}

type AgentsTask struct {
	// 操作类型
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// 探针类型
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// 创建时间
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 更新时间
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// 任务参数
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// 任务状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 任务ID
	AgentsTaskId *string `json:"AgentsTaskId,omitempty" xml:"AgentsTaskId,omitempty"`
}

func (s AgentsTask) String() string {
	return tea.Prettify(s)
}

func (s AgentsTask) GoString() string {
	return s.String()
}

func (s *AgentsTask) SetActionType(v string) *AgentsTask {
	s.ActionType = &v
	return s
}

func (s *AgentsTask) SetAgentType(v string) *AgentsTask {
	s.AgentType = &v
	return s
}

func (s *AgentsTask) SetGmtCreate(v string) *AgentsTask {
	s.GmtCreate = &v
	return s
}

func (s *AgentsTask) SetGmtModify(v string) *AgentsTask {
	s.GmtModify = &v
	return s
}

func (s *AgentsTask) SetParams(v string) *AgentsTask {
	s.Params = &v
	return s
}

func (s *AgentsTask) SetStatus(v string) *AgentsTask {
	s.Status = &v
	return s
}

func (s *AgentsTask) SetAgentsTaskId(v string) *AgentsTask {
	s.AgentsTaskId = &v
	return s
}

type Module struct {
	// 资源一级ID
	ModuleId *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 模块uuid
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 父模块uuid
	ParentModuleId *string `json:"ParentModuleId,omitempty" xml:"ParentModuleId,omitempty"`
	// 模块名字
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 模块类型
	ModuleType *string `json:"ModuleType,omitempty" xml:"ModuleType,omitempty"`
	// 最小建设模块数量
	MinCount *int64 `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	// 最大建设模块数量
	MaxCount *int64 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// 堆叠
	Stack *bool `json:"Stack,omitempty" xml:"Stack,omitempty"`
	// 设备信息
	Device *ModuleDevice `json:"Device,omitempty" xml:"Device,omitempty" type:"Struct"`
	// 模块详情
	ModuleDetails []*ModuleModuleDetails `json:"ModuleDetails,omitempty" xml:"ModuleDetails,omitempty" type:"Repeated"`
}

func (s Module) String() string {
	return tea.Prettify(s)
}

func (s Module) GoString() string {
	return s.String()
}

func (s *Module) SetModuleId(v string) *Module {
	s.ModuleId = &v
	return s
}

func (s *Module) SetCreateTime(v string) *Module {
	s.CreateTime = &v
	return s
}

func (s *Module) SetUpdateTime(v string) *Module {
	s.UpdateTime = &v
	return s
}

func (s *Module) SetId(v string) *Module {
	s.Id = &v
	return s
}

func (s *Module) SetParentModuleId(v string) *Module {
	s.ParentModuleId = &v
	return s
}

func (s *Module) SetName(v string) *Module {
	s.Name = &v
	return s
}

func (s *Module) SetModuleType(v string) *Module {
	s.ModuleType = &v
	return s
}

func (s *Module) SetMinCount(v int64) *Module {
	s.MinCount = &v
	return s
}

func (s *Module) SetMaxCount(v int64) *Module {
	s.MaxCount = &v
	return s
}

func (s *Module) SetStack(v bool) *Module {
	s.Stack = &v
	return s
}

func (s *Module) SetDevice(v *ModuleDevice) *Module {
	s.Device = v
	return s
}

func (s *Module) SetModuleDetails(v []*ModuleModuleDetails) *Module {
	s.ModuleDetails = v
	return s
}

type ModuleDevice struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 设备uuid
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 设备角色
	DeviceRole *string `json:"DeviceRole,omitempty" xml:"DeviceRole,omitempty"`
	// 区块内设备数量
	DeviceCount *int64 `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty"`
	// 设备x坐标
	DeviceX *string `json:"DeviceX,omitempty" xml:"DeviceX,omitempty"`
	// 设备y坐标
	DeviceY *string `json:"DeviceY,omitempty" xml:"DeviceY,omitempty"`
	// 设备角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 区块内设备数量
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// 设备x坐标
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// 设备y坐标
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ModuleDevice) String() string {
	return tea.Prettify(s)
}

func (s ModuleDevice) GoString() string {
	return s.String()
}

func (s *ModuleDevice) SetCreateTime(v string) *ModuleDevice {
	s.CreateTime = &v
	return s
}

func (s *ModuleDevice) SetUpdateTime(v string) *ModuleDevice {
	s.UpdateTime = &v
	return s
}

func (s *ModuleDevice) SetId(v string) *ModuleDevice {
	s.Id = &v
	return s
}

func (s *ModuleDevice) SetDeviceRole(v string) *ModuleDevice {
	s.DeviceRole = &v
	return s
}

func (s *ModuleDevice) SetDeviceCount(v int64) *ModuleDevice {
	s.DeviceCount = &v
	return s
}

func (s *ModuleDevice) SetDeviceX(v string) *ModuleDevice {
	s.DeviceX = &v
	return s
}

func (s *ModuleDevice) SetDeviceY(v string) *ModuleDevice {
	s.DeviceY = &v
	return s
}

func (s *ModuleDevice) SetRole(v string) *ModuleDevice {
	s.Role = &v
	return s
}

func (s *ModuleDevice) SetCount(v int64) *ModuleDevice {
	s.Count = &v
	return s
}

func (s *ModuleDevice) SetX(v string) *ModuleDevice {
	s.X = &v
	return s
}

func (s *ModuleDevice) SetY(v string) *ModuleDevice {
	s.Y = &v
	return s
}

type ModuleModuleDetails struct {
	// 模块序号
	OrderNumber *int64 `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	// 设备详情
	DeviceDetails []*ModuleModuleDetailsDeviceDetails `json:"DeviceDetails,omitempty" xml:"DeviceDetails,omitempty" type:"Repeated"`
}

func (s ModuleModuleDetails) String() string {
	return tea.Prettify(s)
}

func (s ModuleModuleDetails) GoString() string {
	return s.String()
}

func (s *ModuleModuleDetails) SetOrderNumber(v int64) *ModuleModuleDetails {
	s.OrderNumber = &v
	return s
}

func (s *ModuleModuleDetails) SetDeviceDetails(v []*ModuleModuleDetailsDeviceDetails) *ModuleModuleDetails {
	s.DeviceDetails = v
	return s
}

type ModuleModuleDetailsDeviceDetails struct {
	// 设备序号
	OrderNumber *int64 `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
}

func (s ModuleModuleDetailsDeviceDetails) String() string {
	return tea.Prettify(s)
}

func (s ModuleModuleDetailsDeviceDetails) GoString() string {
	return s.String()
}

func (s *ModuleModuleDetailsDeviceDetails) SetOrderNumber(v int64) *ModuleModuleDetailsDeviceDetails {
	s.OrderNumber = &v
	return s
}

type DataViewChart struct {
	// 创建时间
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 修改时间
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 数据视图ID
	DataViewId *string `json:"DataViewId,omitempty" xml:"DataViewId,omitempty"`
	// 图表类型
	ChartType *string `json:"ChartType,omitempty" xml:"ChartType,omitempty"`
	// 数据源类型
	DataViewSource *string `json:"DataViewSource,omitempty" xml:"DataViewSource,omitempty"`
	// 布局配置
	Grid *string `json:"Grid,omitempty" xml:"Grid,omitempty"`
}

func (s DataViewChart) String() string {
	return tea.Prettify(s)
}

func (s DataViewChart) GoString() string {
	return s.String()
}

func (s *DataViewChart) SetGmtCreate(v string) *DataViewChart {
	s.GmtCreate = &v
	return s
}

func (s *DataViewChart) SetGmtModified(v string) *DataViewChart {
	s.GmtModified = &v
	return s
}

func (s *DataViewChart) SetDataViewId(v string) *DataViewChart {
	s.DataViewId = &v
	return s
}

func (s *DataViewChart) SetChartType(v string) *DataViewChart {
	s.ChartType = &v
	return s
}

func (s *DataViewChart) SetDataViewSource(v string) *DataViewChart {
	s.DataViewSource = &v
	return s
}

func (s *DataViewChart) SetGrid(v string) *DataViewChart {
	s.Grid = &v
	return s
}

type ScriptHistory struct {
	// 版本说明
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// 脚本代码
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 创建时间
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 修改时间
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// 模板入参
	Input []*ScriptHistoryInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Repeated"`
	// 模板出参
	Output []*ScriptHistoryOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Repeated"`
	// 脚本id
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// 版本id
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ScriptHistory) String() string {
	return tea.Prettify(s)
}

func (s ScriptHistory) GoString() string {
	return s.String()
}

func (s *ScriptHistory) SetComment(v string) *ScriptHistory {
	s.Comment = &v
	return s
}

func (s *ScriptHistory) SetContent(v string) *ScriptHistory {
	s.Content = &v
	return s
}

func (s *ScriptHistory) SetGmtCreate(v string) *ScriptHistory {
	s.GmtCreate = &v
	return s
}

func (s *ScriptHistory) SetGmtModify(v string) *ScriptHistory {
	s.GmtModify = &v
	return s
}

func (s *ScriptHistory) SetInput(v []*ScriptHistoryInput) *ScriptHistory {
	s.Input = v
	return s
}

func (s *ScriptHistory) SetOutput(v []*ScriptHistoryOutput) *ScriptHistory {
	s.Output = v
	return s
}

func (s *ScriptHistory) SetScriptId(v string) *ScriptHistory {
	s.ScriptId = &v
	return s
}

func (s *ScriptHistory) SetVersionId(v string) *ScriptHistory {
	s.VersionId = &v
	return s
}

type ScriptHistoryInput struct {
	// 参数说明
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 参数名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 参数示例
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// 参数类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ScriptHistoryInput) String() string {
	return tea.Prettify(s)
}

func (s ScriptHistoryInput) GoString() string {
	return s.String()
}

func (s *ScriptHistoryInput) SetDescription(v string) *ScriptHistoryInput {
	s.Description = &v
	return s
}

func (s *ScriptHistoryInput) SetName(v string) *ScriptHistoryInput {
	s.Name = &v
	return s
}

func (s *ScriptHistoryInput) SetSample(v string) *ScriptHistoryInput {
	s.Sample = &v
	return s
}

func (s *ScriptHistoryInput) SetType(v string) *ScriptHistoryInput {
	s.Type = &v
	return s
}

type ScriptHistoryOutput struct {
	// 参数说明
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 参数名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 参数示例
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// 参数类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ScriptHistoryOutput) String() string {
	return tea.Prettify(s)
}

func (s ScriptHistoryOutput) GoString() string {
	return s.String()
}

func (s *ScriptHistoryOutput) SetDescription(v string) *ScriptHistoryOutput {
	s.Description = &v
	return s
}

func (s *ScriptHistoryOutput) SetName(v string) *ScriptHistoryOutput {
	s.Name = &v
	return s
}

func (s *ScriptHistoryOutput) SetSample(v string) *ScriptHistoryOutput {
	s.Sample = &v
	return s
}

func (s *ScriptHistoryOutput) SetType(v string) *ScriptHistoryOutput {
	s.Type = &v
	return s
}

type NetworkArchitecture struct {
	// 资源一级ID
	NetworkArchitectureId *string `json:"NetworkArchitectureId,omitempty" xml:"NetworkArchitectureId,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 架构内容
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 组数
	GroupNumber *int64 `json:"GroupNumber,omitempty" xml:"GroupNumber,omitempty"`
	// 设备数
	DeviceNumber *int64 `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
	// 可用
	Availabe *bool `json:"Availabe,omitempty" xml:"Availabe,omitempty"`
	// 堆叠
	Stack *bool `json:"Stack,omitempty" xml:"Stack,omitempty"`
	// 可选
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
	// 子节点
	Children []*string `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// 架构资源id
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 架构版本
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// 架构名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 架构描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 架构状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 架构最新版本uuid
	ArchVersionIterationId *string `json:"ArchVersionIterationId,omitempty" xml:"ArchVersionIterationId,omitempty"`
}

func (s NetworkArchitecture) String() string {
	return tea.Prettify(s)
}

func (s NetworkArchitecture) GoString() string {
	return s.String()
}

func (s *NetworkArchitecture) SetNetworkArchitectureId(v string) *NetworkArchitecture {
	s.NetworkArchitectureId = &v
	return s
}

func (s *NetworkArchitecture) SetCreateTime(v string) *NetworkArchitecture {
	s.CreateTime = &v
	return s
}

func (s *NetworkArchitecture) SetUpdateTime(v string) *NetworkArchitecture {
	s.UpdateTime = &v
	return s
}

func (s *NetworkArchitecture) SetRole(v string) *NetworkArchitecture {
	s.Role = &v
	return s
}

func (s *NetworkArchitecture) SetGroupNumber(v int64) *NetworkArchitecture {
	s.GroupNumber = &v
	return s
}

func (s *NetworkArchitecture) SetDeviceNumber(v int64) *NetworkArchitecture {
	s.DeviceNumber = &v
	return s
}

func (s *NetworkArchitecture) SetAvailabe(v bool) *NetworkArchitecture {
	s.Availabe = &v
	return s
}

func (s *NetworkArchitecture) SetStack(v bool) *NetworkArchitecture {
	s.Stack = &v
	return s
}

func (s *NetworkArchitecture) SetSelected(v bool) *NetworkArchitecture {
	s.Selected = &v
	return s
}

func (s *NetworkArchitecture) SetChildren(v []*string) *NetworkArchitecture {
	s.Children = v
	return s
}

func (s *NetworkArchitecture) SetId(v string) *NetworkArchitecture {
	s.Id = &v
	return s
}

func (s *NetworkArchitecture) SetVersion(v string) *NetworkArchitecture {
	s.Version = &v
	return s
}

func (s *NetworkArchitecture) SetName(v string) *NetworkArchitecture {
	s.Name = &v
	return s
}

func (s *NetworkArchitecture) SetDescription(v string) *NetworkArchitecture {
	s.Description = &v
	return s
}

func (s *NetworkArchitecture) SetStatus(v string) *NetworkArchitecture {
	s.Status = &v
	return s
}

func (s *NetworkArchitecture) SetArchVersionIterationId(v string) *NetworkArchitecture {
	s.ArchVersionIterationId = &v
	return s
}

type ConfigurationVariate struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 资源一级ID
	ConfigurationVariateId *string `json:"ConfigurationVariateId,omitempty" xml:"ConfigurationVariateId,omitempty"`
	// 变量名字
	VariateName *string `json:"VariateName,omitempty" xml:"VariateName,omitempty"`
	// 描述变量
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// python转换函数
	FormatFunction *string `json:"FormatFunction,omitempty" xml:"FormatFunction,omitempty"`
}

func (s ConfigurationVariate) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationVariate) GoString() string {
	return s.String()
}

func (s *ConfigurationVariate) SetCreateTime(v string) *ConfigurationVariate {
	s.CreateTime = &v
	return s
}

func (s *ConfigurationVariate) SetUpdateTime(v string) *ConfigurationVariate {
	s.UpdateTime = &v
	return s
}

func (s *ConfigurationVariate) SetConfigurationVariateId(v string) *ConfigurationVariate {
	s.ConfigurationVariateId = &v
	return s
}

func (s *ConfigurationVariate) SetVariateName(v string) *ConfigurationVariate {
	s.VariateName = &v
	return s
}

func (s *ConfigurationVariate) SetComment(v string) *ConfigurationVariate {
	s.Comment = &v
	return s
}

func (s *ConfigurationVariate) SetFormatFunction(v string) *ConfigurationVariate {
	s.FormatFunction = &v
	return s
}

type Ip struct {
	// 资源一级ID
	IpId *string `json:"IpId,omitempty" xml:"IpId,omitempty"`
	// 资源名称
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// IP地址
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// 地址段
	ParentIpBlock *string `json:"ParentIpBlock,omitempty" xml:"ParentIpBlock,omitempty"`
	// 业务类型UID
	BusinessTypeId *string `json:"BusinessTypeId,omitempty" xml:"BusinessTypeId,omitempty"`
	// 业务类型名称
	BusinessTypeName *string `json:"BusinessTypeName,omitempty" xml:"BusinessTypeName,omitempty"`
	// 状态 using available lock
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// 设备端口名称
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// 设备MAC
	DeviceMac *string `json:"DeviceMac,omitempty" xml:"DeviceMac,omitempty"`
	// 园区层级
	ZoneLayer []*IpZoneLayer `json:"ZoneLayer,omitempty" xml:"ZoneLayer,omitempty" type:"Repeated"`
}

func (s Ip) String() string {
	return tea.Prettify(s)
}

func (s Ip) GoString() string {
	return s.String()
}

func (s *Ip) SetIpId(v string) *Ip {
	s.IpId = &v
	return s
}

func (s *Ip) SetIpName(v string) *Ip {
	s.IpName = &v
	return s
}

func (s *Ip) SetCreateTime(v string) *Ip {
	s.CreateTime = &v
	return s
}

func (s *Ip) SetIpAddress(v string) *Ip {
	s.IpAddress = &v
	return s
}

func (s *Ip) SetParentIpBlock(v string) *Ip {
	s.ParentIpBlock = &v
	return s
}

func (s *Ip) SetBusinessTypeId(v string) *Ip {
	s.BusinessTypeId = &v
	return s
}

func (s *Ip) SetBusinessTypeName(v string) *Ip {
	s.BusinessTypeName = &v
	return s
}

func (s *Ip) SetStatus(v string) *Ip {
	s.Status = &v
	return s
}

func (s *Ip) SetDeviceName(v string) *Ip {
	s.DeviceName = &v
	return s
}

func (s *Ip) SetPort(v string) *Ip {
	s.Port = &v
	return s
}

func (s *Ip) SetDeviceMac(v string) *Ip {
	s.DeviceMac = &v
	return s
}

func (s *Ip) SetZoneLayer(v []*IpZoneLayer) *Ip {
	s.ZoneLayer = v
	return s
}

type IpZoneLayer struct {
	// 园区层级名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 园区层级值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s IpZoneLayer) String() string {
	return tea.Prettify(s)
}

func (s IpZoneLayer) GoString() string {
	return s.String()
}

func (s *IpZoneLayer) SetName(v string) *IpZoneLayer {
	s.Name = &v
	return s
}

func (s *IpZoneLayer) SetValue(v string) *IpZoneLayer {
	s.Value = &v
	return s
}

type OsVersion struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 版本
	OsVersion *string `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// file
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 文件路径
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// 资源id
	OsVersionId *string `json:"OsVersionId,omitempty" xml:"OsVersionId,omitempty"`
	// 用户名
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// 策略
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// 签名
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// 目录
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// 主机
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
}

func (s OsVersion) String() string {
	return tea.Prettify(s)
}

func (s OsVersion) GoString() string {
	return s.String()
}

func (s *OsVersion) SetCreateTime(v string) *OsVersion {
	s.CreateTime = &v
	return s
}

func (s *OsVersion) SetVendor(v string) *OsVersion {
	s.Vendor = &v
	return s
}

func (s *OsVersion) SetModel(v string) *OsVersion {
	s.Model = &v
	return s
}

func (s *OsVersion) SetOsVersion(v string) *OsVersion {
	s.OsVersion = &v
	return s
}

func (s *OsVersion) SetStatus(v string) *OsVersion {
	s.Status = &v
	return s
}

func (s *OsVersion) SetFileName(v string) *OsVersion {
	s.FileName = &v
	return s
}

func (s *OsVersion) SetFilePath(v string) *OsVersion {
	s.FilePath = &v
	return s
}

func (s *OsVersion) SetOsVersionId(v string) *OsVersion {
	s.OsVersionId = &v
	return s
}

func (s *OsVersion) SetAccessId(v string) *OsVersion {
	s.AccessId = &v
	return s
}

func (s *OsVersion) SetPolicy(v string) *OsVersion {
	s.Policy = &v
	return s
}

func (s *OsVersion) SetSignature(v string) *OsVersion {
	s.Signature = &v
	return s
}

func (s *OsVersion) SetDirectory(v string) *OsVersion {
	s.Directory = &v
	return s
}

func (s *OsVersion) SetHost(v string) *OsVersion {
	s.Host = &v
	return s
}

func (s *OsVersion) SetExpireTime(v string) *OsVersion {
	s.ExpireTime = &v
	return s
}

type TaskLog struct {
	// 函数名
	FuncName *string `json:"FuncName,omitempty" xml:"FuncName,omitempty"`
	// 记录时间
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 日志等级
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// 行数
	LineNo *int32 `json:"LineNo,omitempty" xml:"LineNo,omitempty"`
	// 日志id
	LogId *string `json:"LogId,omitempty" xml:"LogId,omitempty"`
	// 日志信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 任务id
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TaskLog) String() string {
	return tea.Prettify(s)
}

func (s TaskLog) GoString() string {
	return s.String()
}

func (s *TaskLog) SetFuncName(v string) *TaskLog {
	s.FuncName = &v
	return s
}

func (s *TaskLog) SetGmtCreate(v string) *TaskLog {
	s.GmtCreate = &v
	return s
}

func (s *TaskLog) SetLevel(v string) *TaskLog {
	s.Level = &v
	return s
}

func (s *TaskLog) SetLineNo(v int32) *TaskLog {
	s.LineNo = &v
	return s
}

func (s *TaskLog) SetLogId(v string) *TaskLog {
	s.LogId = &v
	return s
}

func (s *TaskLog) SetMessage(v string) *TaskLog {
	s.Message = &v
	return s
}

func (s *TaskLog) SetTaskId(v string) *TaskLog {
	s.TaskId = &v
	return s
}

type ConfigurationSpecification struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 配置规范名字
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// 配置规范uid
	ConfigurationSpecificationId *string `json:"ConfigurationSpecificationId,omitempty" xml:"ConfigurationSpecificationId,omitempty"`
	// 架构类型
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 配置规范内容
	SpecificationContent *string `json:"SpecificationContent,omitempty" xml:"SpecificationContent,omitempty"`
	// 相关变量
	RelatedVariate []*string `json:"RelatedVariate,omitempty" xml:"RelatedVariate,omitempty" type:"Repeated"`
}

func (s ConfigurationSpecification) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationSpecification) GoString() string {
	return s.String()
}

func (s *ConfigurationSpecification) SetCreateTime(v string) *ConfigurationSpecification {
	s.CreateTime = &v
	return s
}

func (s *ConfigurationSpecification) SetUpdateTime(v string) *ConfigurationSpecification {
	s.UpdateTime = &v
	return s
}

func (s *ConfigurationSpecification) SetSpecificationName(v string) *ConfigurationSpecification {
	s.SpecificationName = &v
	return s
}

func (s *ConfigurationSpecification) SetConfigurationSpecificationId(v string) *ConfigurationSpecification {
	s.ConfigurationSpecificationId = &v
	return s
}

func (s *ConfigurationSpecification) SetArchitecture(v string) *ConfigurationSpecification {
	s.Architecture = &v
	return s
}

func (s *ConfigurationSpecification) SetRole(v string) *ConfigurationSpecification {
	s.Role = &v
	return s
}

func (s *ConfigurationSpecification) SetVendor(v string) *ConfigurationSpecification {
	s.Vendor = &v
	return s
}

func (s *ConfigurationSpecification) SetModel(v string) *ConfigurationSpecification {
	s.Model = &v
	return s
}

func (s *ConfigurationSpecification) SetSpecificationContent(v string) *ConfigurationSpecification {
	s.SpecificationContent = &v
	return s
}

func (s *ConfigurationSpecification) SetRelatedVariate(v []*string) *ConfigurationSpecification {
	s.RelatedVariate = v
	return s
}

type ZoneType struct {
	// 资源名称
	ZoneTypeName *string `json:"ZoneTypeName,omitempty" xml:"ZoneTypeName,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 资源一级ID
	ZoneTypeId *string `json:"ZoneTypeId,omitempty" xml:"ZoneTypeId,omitempty"`
	// 园区类型名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 园区层级
	ZoneTypeLayer []*ZoneTypeZoneTypeLayer `json:"ZoneTypeLayer,omitempty" xml:"ZoneTypeLayer,omitempty" type:"Repeated"`
}

func (s ZoneType) String() string {
	return tea.Prettify(s)
}

func (s ZoneType) GoString() string {
	return s.String()
}

func (s *ZoneType) SetZoneTypeName(v string) *ZoneType {
	s.ZoneTypeName = &v
	return s
}

func (s *ZoneType) SetCreateTime(v string) *ZoneType {
	s.CreateTime = &v
	return s
}

func (s *ZoneType) SetZoneTypeId(v string) *ZoneType {
	s.ZoneTypeId = &v
	return s
}

func (s *ZoneType) SetName(v string) *ZoneType {
	s.Name = &v
	return s
}

func (s *ZoneType) SetZoneTypeLayer(v []*ZoneTypeZoneTypeLayer) *ZoneType {
	s.ZoneTypeLayer = v
	return s
}

type ZoneTypeZoneTypeLayer struct {
	// 园区层级名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 园区层级掩码
	Mask *int64 `json:"Mask,omitempty" xml:"Mask,omitempty"`
}

func (s ZoneTypeZoneTypeLayer) String() string {
	return tea.Prettify(s)
}

func (s ZoneTypeZoneTypeLayer) GoString() string {
	return s.String()
}

func (s *ZoneTypeZoneTypeLayer) SetName(v string) *ZoneTypeZoneTypeLayer {
	s.Name = &v
	return s
}

func (s *ZoneTypeZoneTypeLayer) SetMask(v int64) *ZoneTypeZoneTypeLayer {
	s.Mask = &v
	return s
}

type SetupProject struct {
	// 资源名称
	SetupProjectName *string `json:"SetupProjectName,omitempty" xml:"SetupProjectName,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 资源一级ID
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// 物理空间uId
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 架构id
	ArchId *string `json:"ArchId,omitempty" xml:"ArchId,omitempty"`
	// 预计交付时间
	DeliveryTime *string `json:"DeliveryTime,omitempty" xml:"DeliveryTime,omitempty"`
	// 节点
	Nodes *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	// 架构id
	ArchitectureId *string `json:"ArchitectureId,omitempty" xml:"ArchitectureId,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 物理空间类型
	SpaceType *string `json:"SpaceType,omitempty" xml:"SpaceType,omitempty"`
	// 物理空间名称
	SpaceName *string `json:"SpaceName,omitempty" xml:"SpaceName,omitempty"`
	// 架构版本
	ArchVersion *string `json:"ArchVersion,omitempty" xml:"ArchVersion,omitempty"`
	// 套餐
	Packages []*SetupProjectPackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Repeated"`
}

func (s SetupProject) String() string {
	return tea.Prettify(s)
}

func (s SetupProject) GoString() string {
	return s.String()
}

func (s *SetupProject) SetSetupProjectName(v string) *SetupProject {
	s.SetupProjectName = &v
	return s
}

func (s *SetupProject) SetCreateTime(v string) *SetupProject {
	s.CreateTime = &v
	return s
}

func (s *SetupProject) SetSetupProjectId(v string) *SetupProject {
	s.SetupProjectId = &v
	return s
}

func (s *SetupProject) SetSpaceId(v string) *SetupProject {
	s.SpaceId = &v
	return s
}

func (s *SetupProject) SetDescription(v string) *SetupProject {
	s.Description = &v
	return s
}

func (s *SetupProject) SetArchId(v string) *SetupProject {
	s.ArchId = &v
	return s
}

func (s *SetupProject) SetDeliveryTime(v string) *SetupProject {
	s.DeliveryTime = &v
	return s
}

func (s *SetupProject) SetNodes(v string) *SetupProject {
	s.Nodes = &v
	return s
}

func (s *SetupProject) SetArchitectureId(v string) *SetupProject {
	s.ArchitectureId = &v
	return s
}

func (s *SetupProject) SetStatus(v string) *SetupProject {
	s.Status = &v
	return s
}

func (s *SetupProject) SetSpaceType(v string) *SetupProject {
	s.SpaceType = &v
	return s
}

func (s *SetupProject) SetSpaceName(v string) *SetupProject {
	s.SpaceName = &v
	return s
}

func (s *SetupProject) SetArchVersion(v string) *SetupProject {
	s.ArchVersion = &v
	return s
}

func (s *SetupProject) SetPackages(v []*SetupProjectPackages) *SetupProject {
	s.Packages = v
	return s
}

type SetupProjectPackages struct {
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 设备号
	DeviceNumber *string `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s SetupProjectPackages) String() string {
	return tea.Prettify(s)
}

func (s SetupProjectPackages) GoString() string {
	return s.String()
}

func (s *SetupProjectPackages) SetRole(v string) *SetupProjectPackages {
	s.Role = &v
	return s
}

func (s *SetupProjectPackages) SetDeviceNumber(v string) *SetupProjectPackages {
	s.DeviceNumber = &v
	return s
}

func (s *SetupProjectPackages) SetVendor(v string) *SetupProjectPackages {
	s.Vendor = &v
	return s
}

func (s *SetupProjectPackages) SetModel(v string) *SetupProjectPackages {
	s.Model = &v
	return s
}

type InspectionTask struct {
	// 巡检模板ID
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 巡检项名字
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// 巡检项ID
	ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// 物理空间
	Space *string `json:"Space,omitempty" xml:"Space,omitempty"`
	// 主机名
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// 设备IP
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 任务状态
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 执行结果
	InspectionResult *string `json:"InspectionResult,omitempty" xml:"InspectionResult,omitempty"`
	// 执行开始时间
	ExecutionBeginTime *string `json:"ExecutionBeginTime,omitempty" xml:"ExecutionBeginTime,omitempty"`
	// 执行结束时间
	ExecutionEndTime *string `json:"ExecutionEndTime,omitempty" xml:"ExecutionEndTime,omitempty"`
	// 告警规则
	InspectionAlarmRules []*InspectionTaskInspectionAlarmRules `json:"InspectionAlarmRules,omitempty" xml:"InspectionAlarmRules,omitempty" type:"Repeated"`
}

func (s InspectionTask) String() string {
	return tea.Prettify(s)
}

func (s InspectionTask) GoString() string {
	return s.String()
}

func (s *InspectionTask) SetTemplateId(v string) *InspectionTask {
	s.TemplateId = &v
	return s
}

func (s *InspectionTask) SetItemName(v string) *InspectionTask {
	s.ItemName = &v
	return s
}

func (s *InspectionTask) SetItemId(v string) *InspectionTask {
	s.ItemId = &v
	return s
}

func (s *InspectionTask) SetSpace(v string) *InspectionTask {
	s.Space = &v
	return s
}

func (s *InspectionTask) SetHostname(v string) *InspectionTask {
	s.Hostname = &v
	return s
}

func (s *InspectionTask) SetIP(v string) *InspectionTask {
	s.IP = &v
	return s
}

func (s *InspectionTask) SetVendor(v string) *InspectionTask {
	s.Vendor = &v
	return s
}

func (s *InspectionTask) SetModel(v string) *InspectionTask {
	s.Model = &v
	return s
}

func (s *InspectionTask) SetRole(v string) *InspectionTask {
	s.Role = &v
	return s
}

func (s *InspectionTask) SetTaskStatus(v string) *InspectionTask {
	s.TaskStatus = &v
	return s
}

func (s *InspectionTask) SetDeviceId(v string) *InspectionTask {
	s.DeviceId = &v
	return s
}

func (s *InspectionTask) SetErrorCode(v string) *InspectionTask {
	s.ErrorCode = &v
	return s
}

func (s *InspectionTask) SetInspectionResult(v string) *InspectionTask {
	s.InspectionResult = &v
	return s
}

func (s *InspectionTask) SetExecutionBeginTime(v string) *InspectionTask {
	s.ExecutionBeginTime = &v
	return s
}

func (s *InspectionTask) SetExecutionEndTime(v string) *InspectionTask {
	s.ExecutionEndTime = &v
	return s
}

func (s *InspectionTask) SetInspectionAlarmRules(v []*InspectionTaskInspectionAlarmRules) *InspectionTask {
	s.InspectionAlarmRules = v
	return s
}

type InspectionTaskInspectionAlarmRules struct {
	// 告警表达式
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// 告警操作符
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// 告警值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// 告警实际值
	ActualValue *string `json:"ActualValue,omitempty" xml:"ActualValue,omitempty"`
	// 告警级别
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s InspectionTaskInspectionAlarmRules) String() string {
	return tea.Prettify(s)
}

func (s InspectionTaskInspectionAlarmRules) GoString() string {
	return s.String()
}

func (s *InspectionTaskInspectionAlarmRules) SetExpression(v string) *InspectionTaskInspectionAlarmRules {
	s.Expression = &v
	return s
}

func (s *InspectionTaskInspectionAlarmRules) SetOperator(v string) *InspectionTaskInspectionAlarmRules {
	s.Operator = &v
	return s
}

func (s *InspectionTaskInspectionAlarmRules) SetValue(v string) *InspectionTaskInspectionAlarmRules {
	s.Value = &v
	return s
}

func (s *InspectionTaskInspectionAlarmRules) SetActualValue(v string) *InspectionTaskInspectionAlarmRules {
	s.ActualValue = &v
	return s
}

func (s *InspectionTaskInspectionAlarmRules) SetLevel(v string) *InspectionTaskInspectionAlarmRules {
	s.Level = &v
	return s
}

type CliTask struct {
	// agent IP
	AgentIp *string `json:"AgentIp,omitempty" xml:"AgentIp,omitempty"`
	// cli任务id
	CliTaskId *string `json:"CliTaskId,omitempty" xml:"CliTaskId,omitempty"`
	// cli命令
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// 设备id
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 创建时间
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 修改时间
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// 设备回显
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// 协议
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// 任务错误码
	ResponseCode *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	// 任务结果
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// 会话id
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// cli任务状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 超时参数
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CliTask) String() string {
	return tea.Prettify(s)
}

func (s CliTask) GoString() string {
	return s.String()
}

func (s *CliTask) SetAgentIp(v string) *CliTask {
	s.AgentIp = &v
	return s
}

func (s *CliTask) SetCliTaskId(v string) *CliTask {
	s.CliTaskId = &v
	return s
}

func (s *CliTask) SetCommand(v string) *CliTask {
	s.Command = &v
	return s
}

func (s *CliTask) SetDeviceId(v string) *CliTask {
	s.DeviceId = &v
	return s
}

func (s *CliTask) SetGmtCreate(v string) *CliTask {
	s.GmtCreate = &v
	return s
}

func (s *CliTask) SetGmtModify(v string) *CliTask {
	s.GmtModify = &v
	return s
}

func (s *CliTask) SetOutput(v string) *CliTask {
	s.Output = &v
	return s
}

func (s *CliTask) SetProtocol(v string) *CliTask {
	s.Protocol = &v
	return s
}

func (s *CliTask) SetResponseCode(v string) *CliTask {
	s.ResponseCode = &v
	return s
}

func (s *CliTask) SetResult(v string) *CliTask {
	s.Result = &v
	return s
}

func (s *CliTask) SetSessionId(v string) *CliTask {
	s.SessionId = &v
	return s
}

func (s *CliTask) SetStatus(v string) *CliTask {
	s.Status = &v
	return s
}

func (s *CliTask) SetTimeout(v int32) *CliTask {
	s.Timeout = &v
	return s
}

type IpBlockRecord struct {
	// 工单uuid
	IpBlockRecordId *string `json:"IpBlockRecordId,omitempty" xml:"IpBlockRecordId,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 工单名称
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 创建人
	Creator *int64 `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// 工号
	WorkNo *string `json:"WorkNo,omitempty" xml:"WorkNo,omitempty"`
	// 公网私网标志 PUBLIC PRIVATE
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// 工单状态 running complete fail cancel lock approving
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 园区名称
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	// 工单备注
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 工单详情
	Detail []*IpBlockRecordDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// 工单类型 Apply 申请工单 Recycle 释放工单
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// 要释放的IP段
	IpBlockCode []*string `json:"IpBlockCode,omitempty" xml:"IpBlockCode,omitempty" type:"Repeated"`
}

func (s IpBlockRecord) String() string {
	return tea.Prettify(s)
}

func (s IpBlockRecord) GoString() string {
	return s.String()
}

func (s *IpBlockRecord) SetIpBlockRecordId(v string) *IpBlockRecord {
	s.IpBlockRecordId = &v
	return s
}

func (s *IpBlockRecord) SetCreateTime(v string) *IpBlockRecord {
	s.CreateTime = &v
	return s
}

func (s *IpBlockRecord) SetUpdateTime(v string) *IpBlockRecord {
	s.UpdateTime = &v
	return s
}

func (s *IpBlockRecord) SetTitle(v string) *IpBlockRecord {
	s.Title = &v
	return s
}

func (s *IpBlockRecord) SetCreator(v int64) *IpBlockRecord {
	s.Creator = &v
	return s
}

func (s *IpBlockRecord) SetWorkNo(v string) *IpBlockRecord {
	s.WorkNo = &v
	return s
}

func (s *IpBlockRecord) SetNetType(v string) *IpBlockRecord {
	s.NetType = &v
	return s
}

func (s *IpBlockRecord) SetStatus(v string) *IpBlockRecord {
	s.Status = &v
	return s
}

func (s *IpBlockRecord) SetZoneName(v string) *IpBlockRecord {
	s.ZoneName = &v
	return s
}

func (s *IpBlockRecord) SetDescription(v string) *IpBlockRecord {
	s.Description = &v
	return s
}

func (s *IpBlockRecord) SetDetail(v []*IpBlockRecordDetail) *IpBlockRecord {
	s.Detail = v
	return s
}

func (s *IpBlockRecord) SetRecordType(v string) *IpBlockRecord {
	s.RecordType = &v
	return s
}

func (s *IpBlockRecord) SetIpBlockCode(v []*string) *IpBlockRecord {
	s.IpBlockCode = v
	return s
}

type IpBlockRecordDetail struct {
	// 业务类型
	BusinessTypeName *string `json:"BusinessTypeName,omitempty" xml:"BusinessTypeName,omitempty"`
	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// 网关
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// 父地址段列表
	ParentIpBlocks []*string `json:"ParentIpBlocks,omitempty" xml:"ParentIpBlocks,omitempty" type:"Repeated"`
	// 园区层级
	ZoneLayer []*IpBlockRecordDetailZoneLayer `json:"ZoneLayer,omitempty" xml:"ZoneLayer,omitempty" type:"Repeated"`
}

func (s IpBlockRecordDetail) String() string {
	return tea.Prettify(s)
}

func (s IpBlockRecordDetail) GoString() string {
	return s.String()
}

func (s *IpBlockRecordDetail) SetBusinessTypeName(v string) *IpBlockRecordDetail {
	s.BusinessTypeName = &v
	return s
}

func (s *IpBlockRecordDetail) SetDeviceName(v string) *IpBlockRecordDetail {
	s.DeviceName = &v
	return s
}

func (s *IpBlockRecordDetail) SetGateway(v string) *IpBlockRecordDetail {
	s.Gateway = &v
	return s
}

func (s *IpBlockRecordDetail) SetParentIpBlocks(v []*string) *IpBlockRecordDetail {
	s.ParentIpBlocks = v
	return s
}

func (s *IpBlockRecordDetail) SetZoneLayer(v []*IpBlockRecordDetailZoneLayer) *IpBlockRecordDetail {
	s.ZoneLayer = v
	return s
}

type IpBlockRecordDetailZoneLayer struct {
	// 园区层级名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 园区层级值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s IpBlockRecordDetailZoneLayer) String() string {
	return tea.Prettify(s)
}

func (s IpBlockRecordDetailZoneLayer) GoString() string {
	return s.String()
}

func (s *IpBlockRecordDetailZoneLayer) SetName(v string) *IpBlockRecordDetailZoneLayer {
	s.Name = &v
	return s
}

func (s *IpBlockRecordDetailZoneLayer) SetValue(v string) *IpBlockRecordDetailZoneLayer {
	s.Value = &v
	return s
}

type IpRecord struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 资源一级ID
	IpRecordId *string `json:"IpRecordId,omitempty" xml:"IpRecordId,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 业务类型名称
	BusinessTypeName *string `json:"BusinessTypeName,omitempty" xml:"BusinessTypeName,omitempty"`
	// 创建人
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// 工单状态 running complete fail
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 园区名
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	// 备注
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 地址段
	IpBlock *string `json:"IpBlock,omitempty" xml:"IpBlock,omitempty"`
	// 工单详情
	Detail []*IpRecordDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// 工单类型 Apply 申请工单 Recycle 释放工单
	RecodeType *string `json:"RecodeType,omitempty" xml:"RecodeType,omitempty"`
	// IP地址列表
	IpCode []*string `json:"IpCode,omitempty" xml:"IpCode,omitempty" type:"Repeated"`
}

func (s IpRecord) String() string {
	return tea.Prettify(s)
}

func (s IpRecord) GoString() string {
	return s.String()
}

func (s *IpRecord) SetCreateTime(v string) *IpRecord {
	s.CreateTime = &v
	return s
}

func (s *IpRecord) SetIpRecordId(v string) *IpRecord {
	s.IpRecordId = &v
	return s
}

func (s *IpRecord) SetUpdateTime(v string) *IpRecord {
	s.UpdateTime = &v
	return s
}

func (s *IpRecord) SetBusinessTypeName(v string) *IpRecord {
	s.BusinessTypeName = &v
	return s
}

func (s *IpRecord) SetCreator(v string) *IpRecord {
	s.Creator = &v
	return s
}

func (s *IpRecord) SetStatus(v string) *IpRecord {
	s.Status = &v
	return s
}

func (s *IpRecord) SetZoneName(v string) *IpRecord {
	s.ZoneName = &v
	return s
}

func (s *IpRecord) SetDescription(v string) *IpRecord {
	s.Description = &v
	return s
}

func (s *IpRecord) SetIpBlock(v string) *IpRecord {
	s.IpBlock = &v
	return s
}

func (s *IpRecord) SetDetail(v []*IpRecordDetail) *IpRecord {
	s.Detail = v
	return s
}

func (s *IpRecord) SetRecodeType(v string) *IpRecord {
	s.RecodeType = &v
	return s
}

func (s *IpRecord) SetIpCode(v []*string) *IpRecord {
	s.IpCode = v
	return s
}

type IpRecordDetail struct {
	// 申请到的Ip
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// 设备端口
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// 对端IP
	RemoteIp *string `json:"RemoteIp,omitempty" xml:"RemoteIp,omitempty"`
	// 对端设备名称
	RemoteDeviceName *string `json:"RemoteDeviceName,omitempty" xml:"RemoteDeviceName,omitempty"`
	// 对端设备端口
	RemotePort *string `json:"RemotePort,omitempty" xml:"RemotePort,omitempty"`
	// 设备MAC
	DeviceMac *string `json:"DeviceMac,omitempty" xml:"DeviceMac,omitempty"`
	// 网关
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// 父地址段列表
	ParentIpBlocks []*string `json:"ParentIpBlocks,omitempty" xml:"ParentIpBlocks,omitempty" type:"Repeated"`
	// 园区层级
	ZoneLayer []*IpRecordDetailZoneLayer `json:"ZoneLayer,omitempty" xml:"ZoneLayer,omitempty" type:"Repeated"`
}

func (s IpRecordDetail) String() string {
	return tea.Prettify(s)
}

func (s IpRecordDetail) GoString() string {
	return s.String()
}

func (s *IpRecordDetail) SetIp(v string) *IpRecordDetail {
	s.Ip = &v
	return s
}

func (s *IpRecordDetail) SetDeviceName(v string) *IpRecordDetail {
	s.DeviceName = &v
	return s
}

func (s *IpRecordDetail) SetPort(v string) *IpRecordDetail {
	s.Port = &v
	return s
}

func (s *IpRecordDetail) SetRemoteIp(v string) *IpRecordDetail {
	s.RemoteIp = &v
	return s
}

func (s *IpRecordDetail) SetRemoteDeviceName(v string) *IpRecordDetail {
	s.RemoteDeviceName = &v
	return s
}

func (s *IpRecordDetail) SetRemotePort(v string) *IpRecordDetail {
	s.RemotePort = &v
	return s
}

func (s *IpRecordDetail) SetDeviceMac(v string) *IpRecordDetail {
	s.DeviceMac = &v
	return s
}

func (s *IpRecordDetail) SetGateway(v string) *IpRecordDetail {
	s.Gateway = &v
	return s
}

func (s *IpRecordDetail) SetParentIpBlocks(v []*string) *IpRecordDetail {
	s.ParentIpBlocks = v
	return s
}

func (s *IpRecordDetail) SetZoneLayer(v []*IpRecordDetailZoneLayer) *IpRecordDetail {
	s.ZoneLayer = v
	return s
}

type IpRecordDetailZoneLayer struct {
	// 园区层级名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 园区层级值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s IpRecordDetailZoneLayer) String() string {
	return tea.Prettify(s)
}

func (s IpRecordDetailZoneLayer) GoString() string {
	return s.String()
}

func (s *IpRecordDetailZoneLayer) SetName(v string) *IpRecordDetailZoneLayer {
	s.Name = &v
	return s
}

func (s *IpRecordDetailZoneLayer) SetValue(v string) *IpRecordDetailZoneLayer {
	s.Value = &v
	return s
}

type MonitorItem struct {
	// 解析代码
	AnalysisCode *string `json:"AnalysisCode,omitempty" xml:"AnalysisCode,omitempty"`
	// 采集类型
	CollectionType *string `json:"CollectionType,omitempty" xml:"CollectionType,omitempty"`
	// 采集配置
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// 描述
	MonitorItemDescription *string `json:"MonitorItemDescription,omitempty" xml:"MonitorItemDescription,omitempty"`
	// 是否启用
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// 采集间隔
	ExecInterval *string `json:"ExecInterval,omitempty" xml:"ExecInterval,omitempty"`
	// 创建时间
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 更新时间
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 监控项名称
	MonitorItemName *string `json:"MonitorItemName,omitempty" xml:"MonitorItemName,omitempty"`
	// 安全域
	SecurityDomain *string `json:"SecurityDomain,omitempty" xml:"SecurityDomain,omitempty"`
	// 监控项ID
	MonitorItemId *string `json:"MonitorItemId,omitempty" xml:"MonitorItemId,omitempty"`
}

func (s MonitorItem) String() string {
	return tea.Prettify(s)
}

func (s MonitorItem) GoString() string {
	return s.String()
}

func (s *MonitorItem) SetAnalysisCode(v string) *MonitorItem {
	s.AnalysisCode = &v
	return s
}

func (s *MonitorItem) SetCollectionType(v string) *MonitorItem {
	s.CollectionType = &v
	return s
}

func (s *MonitorItem) SetConfig(v string) *MonitorItem {
	s.Config = &v
	return s
}

func (s *MonitorItem) SetMonitorItemDescription(v string) *MonitorItem {
	s.MonitorItemDescription = &v
	return s
}

func (s *MonitorItem) SetEnable(v int32) *MonitorItem {
	s.Enable = &v
	return s
}

func (s *MonitorItem) SetExecInterval(v string) *MonitorItem {
	s.ExecInterval = &v
	return s
}

func (s *MonitorItem) SetGmtCreate(v string) *MonitorItem {
	s.GmtCreate = &v
	return s
}

func (s *MonitorItem) SetGmtModified(v string) *MonitorItem {
	s.GmtModified = &v
	return s
}

func (s *MonitorItem) SetMonitorItemName(v string) *MonitorItem {
	s.MonitorItemName = &v
	return s
}

func (s *MonitorItem) SetSecurityDomain(v string) *MonitorItem {
	s.SecurityDomain = &v
	return s
}

func (s *MonitorItem) SetMonitorItemId(v string) *MonitorItem {
	s.MonitorItemId = &v
	return s
}

type PhysicalSpace struct {
	// 具体所在地址
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// 所属城市
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// 所属国家
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// 所属省份
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// 物理空间ID
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// 物理空间名称
	SpaceName *string `json:"SpaceName,omitempty" xml:"SpaceName,omitempty"`
}

func (s PhysicalSpace) String() string {
	return tea.Prettify(s)
}

func (s PhysicalSpace) GoString() string {
	return s.String()
}

func (s *PhysicalSpace) SetAddress(v string) *PhysicalSpace {
	s.Address = &v
	return s
}

func (s *PhysicalSpace) SetCity(v string) *PhysicalSpace {
	s.City = &v
	return s
}

func (s *PhysicalSpace) SetCountry(v string) *PhysicalSpace {
	s.Country = &v
	return s
}

func (s *PhysicalSpace) SetProvince(v string) *PhysicalSpace {
	s.Province = &v
	return s
}

func (s *PhysicalSpace) SetSpaceId(v string) *PhysicalSpace {
	s.SpaceId = &v
	return s
}

func (s *PhysicalSpace) SetSpaceName(v string) *PhysicalSpace {
	s.SpaceName = &v
	return s
}

type InspectionScript struct {
	// 巡检模板ID
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// 巡检项ID
	ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// 巡检项名字
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// 巡检项描述
	ItemDescription *string `json:"ItemDescription,omitempty" xml:"ItemDescription,omitempty"`
	// 巡检项定时表达式
	InspectionCrontab *string `json:"InspectionCrontab,omitempty" xml:"InspectionCrontab,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 模板状态
	ScriptStatus *string `json:"ScriptStatus,omitempty" xml:"ScriptStatus,omitempty"`
	// 模板执行内容
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// 巡检告警规则
	InspectionAlarmRules []*InspectionScriptInspectionAlarmRules `json:"InspectionAlarmRules,omitempty" xml:"InspectionAlarmRules,omitempty" type:"Repeated"`
}

func (s InspectionScript) String() string {
	return tea.Prettify(s)
}

func (s InspectionScript) GoString() string {
	return s.String()
}

func (s *InspectionScript) SetScriptId(v string) *InspectionScript {
	s.ScriptId = &v
	return s
}

func (s *InspectionScript) SetItemId(v string) *InspectionScript {
	s.ItemId = &v
	return s
}

func (s *InspectionScript) SetItemName(v string) *InspectionScript {
	s.ItemName = &v
	return s
}

func (s *InspectionScript) SetItemDescription(v string) *InspectionScript {
	s.ItemDescription = &v
	return s
}

func (s *InspectionScript) SetInspectionCrontab(v string) *InspectionScript {
	s.InspectionCrontab = &v
	return s
}

func (s *InspectionScript) SetVendor(v string) *InspectionScript {
	s.Vendor = &v
	return s
}

func (s *InspectionScript) SetModel(v string) *InspectionScript {
	s.Model = &v
	return s
}

func (s *InspectionScript) SetRole(v string) *InspectionScript {
	s.Role = &v
	return s
}

func (s *InspectionScript) SetScriptStatus(v string) *InspectionScript {
	s.ScriptStatus = &v
	return s
}

func (s *InspectionScript) SetScript(v string) *InspectionScript {
	s.Script = &v
	return s
}

func (s *InspectionScript) SetInspectionAlarmRules(v []*InspectionScriptInspectionAlarmRules) *InspectionScript {
	s.InspectionAlarmRules = v
	return s
}

type InspectionScriptInspectionAlarmRules struct {
	// 告警表达式
	AlarmExpression *string `json:"AlarmExpression,omitempty" xml:"AlarmExpression,omitempty"`
	// 告警符号
	AlarmOperator *string `json:"AlarmOperator,omitempty" xml:"AlarmOperator,omitempty"`
	// 告警值
	AlarmValue *string `json:"AlarmValue,omitempty" xml:"AlarmValue,omitempty"`
	// 告警级别
	AlarmLevel *string `json:"AlarmLevel,omitempty" xml:"AlarmLevel,omitempty"`
}

func (s InspectionScriptInspectionAlarmRules) String() string {
	return tea.Prettify(s)
}

func (s InspectionScriptInspectionAlarmRules) GoString() string {
	return s.String()
}

func (s *InspectionScriptInspectionAlarmRules) SetAlarmExpression(v string) *InspectionScriptInspectionAlarmRules {
	s.AlarmExpression = &v
	return s
}

func (s *InspectionScriptInspectionAlarmRules) SetAlarmOperator(v string) *InspectionScriptInspectionAlarmRules {
	s.AlarmOperator = &v
	return s
}

func (s *InspectionScriptInspectionAlarmRules) SetAlarmValue(v string) *InspectionScriptInspectionAlarmRules {
	s.AlarmValue = &v
	return s
}

func (s *InspectionScriptInspectionAlarmRules) SetAlarmLevel(v string) *InspectionScriptInspectionAlarmRules {
	s.AlarmLevel = &v
	return s
}

type SubscriptionItem struct {
	// 告警状态
	AlarmStatus *string `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	// 通知方式
	NotificationMode *string `json:"NotificationMode,omitempty" xml:"NotificationMode,omitempty"`
	// 抑制策略
	SuppressionStrategy *string `json:"SuppressionStrategy,omitempty" xml:"SuppressionStrategy,omitempty"`
	// 通知组ID
	NotificationGroupId *string `json:"NotificationGroupId,omitempty" xml:"NotificationGroupId,omitempty"`
	// 订阅类型
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	// 连续触发次数
	TriggerTimes *int32 `json:"TriggerTimes,omitempty" xml:"TriggerTimes,omitempty"`
	// 监控项ID
	MonitorItemId *string `json:"MonitorItemId,omitempty" xml:"MonitorItemId,omitempty"`
	// 发送通知的语言
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// 是否发送通知
	RecoveryNotice *int32 `json:"RecoveryNotice,omitempty" xml:"RecoveryNotice,omitempty"`
	// 订阅项ID
	SubscriptionItemId *string `json:"SubscriptionItemId,omitempty" xml:"SubscriptionItemId,omitempty"`
}

func (s SubscriptionItem) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionItem) GoString() string {
	return s.String()
}

func (s *SubscriptionItem) SetAlarmStatus(v string) *SubscriptionItem {
	s.AlarmStatus = &v
	return s
}

func (s *SubscriptionItem) SetNotificationMode(v string) *SubscriptionItem {
	s.NotificationMode = &v
	return s
}

func (s *SubscriptionItem) SetSuppressionStrategy(v string) *SubscriptionItem {
	s.SuppressionStrategy = &v
	return s
}

func (s *SubscriptionItem) SetNotificationGroupId(v string) *SubscriptionItem {
	s.NotificationGroupId = &v
	return s
}

func (s *SubscriptionItem) SetSubscriptionType(v string) *SubscriptionItem {
	s.SubscriptionType = &v
	return s
}

func (s *SubscriptionItem) SetTriggerTimes(v int32) *SubscriptionItem {
	s.TriggerTimes = &v
	return s
}

func (s *SubscriptionItem) SetMonitorItemId(v string) *SubscriptionItem {
	s.MonitorItemId = &v
	return s
}

func (s *SubscriptionItem) SetLanguage(v string) *SubscriptionItem {
	s.Language = &v
	return s
}

func (s *SubscriptionItem) SetRecoveryNotice(v int32) *SubscriptionItem {
	s.RecoveryNotice = &v
	return s
}

func (s *SubscriptionItem) SetSubscriptionItemId(v string) *SubscriptionItem {
	s.SubscriptionItemId = &v
	return s
}

type DataView struct {
	// 创建时间
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 修改时间
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 数据视图名称
	DataViewName *string `json:"DataViewName,omitempty" xml:"DataViewName,omitempty"`
	// 描述
	DataViewDescription *string `json:"DataViewDescription,omitempty" xml:"DataViewDescription,omitempty"`
	// 聚合数据
	DataViewId *string `json:"DataViewId,omitempty" xml:"DataViewId,omitempty"`
	// 图表列表
	DataViewChartList []*DataViewChart `json:"DataViewChartList,omitempty" xml:"DataViewChartList,omitempty" type:"Repeated"`
}

func (s DataView) String() string {
	return tea.Prettify(s)
}

func (s DataView) GoString() string {
	return s.String()
}

func (s *DataView) SetGmtCreate(v string) *DataView {
	s.GmtCreate = &v
	return s
}

func (s *DataView) SetGmtModified(v string) *DataView {
	s.GmtModified = &v
	return s
}

func (s *DataView) SetDataViewName(v string) *DataView {
	s.DataViewName = &v
	return s
}

func (s *DataView) SetDataViewDescription(v string) *DataView {
	s.DataViewDescription = &v
	return s
}

func (s *DataView) SetDataViewId(v string) *DataView {
	s.DataViewId = &v
	return s
}

func (s *DataView) SetDataViewChartList(v []*DataViewChart) *DataView {
	s.DataViewChartList = v
	return s
}

type DeviceTask struct {
	// 设备id
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 设备ip
	DeviceIp *string `json:"DeviceIp,omitempty" xml:"DeviceIp,omitempty"`
	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// 设备任务id
	DeviceTaskId *string `json:"DeviceTaskId,omitempty" xml:"DeviceTaskId,omitempty"`
	// 创建时间
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 修改时间
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// 设备任务回显
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// 设备任务参数
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// 设备任务错误码
	ResponseCode *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	// 设备任务返回
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// 脚本id
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// 版本id
	ScriptVersion *string `json:"ScriptVersion,omitempty" xml:"ScriptVersion,omitempty"`
	// 设备任务状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板id
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 模板名称
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DeviceTask) String() string {
	return tea.Prettify(s)
}

func (s DeviceTask) GoString() string {
	return s.String()
}

func (s *DeviceTask) SetDeviceId(v string) *DeviceTask {
	s.DeviceId = &v
	return s
}

func (s *DeviceTask) SetDeviceIp(v string) *DeviceTask {
	s.DeviceIp = &v
	return s
}

func (s *DeviceTask) SetDeviceName(v string) *DeviceTask {
	s.DeviceName = &v
	return s
}

func (s *DeviceTask) SetDeviceTaskId(v string) *DeviceTask {
	s.DeviceTaskId = &v
	return s
}

func (s *DeviceTask) SetGmtCreate(v string) *DeviceTask {
	s.GmtCreate = &v
	return s
}

func (s *DeviceTask) SetGmtModify(v string) *DeviceTask {
	s.GmtModify = &v
	return s
}

func (s *DeviceTask) SetOutput(v string) *DeviceTask {
	s.Output = &v
	return s
}

func (s *DeviceTask) SetParams(v string) *DeviceTask {
	s.Params = &v
	return s
}

func (s *DeviceTask) SetResponseCode(v string) *DeviceTask {
	s.ResponseCode = &v
	return s
}

func (s *DeviceTask) SetResult(v string) *DeviceTask {
	s.Result = &v
	return s
}

func (s *DeviceTask) SetScriptId(v string) *DeviceTask {
	s.ScriptId = &v
	return s
}

func (s *DeviceTask) SetScriptVersion(v string) *DeviceTask {
	s.ScriptVersion = &v
	return s
}

func (s *DeviceTask) SetStatus(v string) *DeviceTask {
	s.Status = &v
	return s
}

func (s *DeviceTask) SetTemplateId(v string) *DeviceTask {
	s.TemplateId = &v
	return s
}

func (s *DeviceTask) SetTemplateName(v string) *DeviceTask {
	s.TemplateName = &v
	return s
}

type PortCollection struct {
	// 创建时间
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 修改时间
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 描述
	PortCollectionDescription *string `json:"PortCollectionDescription,omitempty" xml:"PortCollectionDescription,omitempty"`
	// 端口集ID
	PortCollectionId *string `json:"PortCollectionId,omitempty" xml:"PortCollectionId,omitempty"`
	// 端口集名称
	PortCollectionName *string `json:"PortCollectionName,omitempty" xml:"PortCollectionName,omitempty"`
	// 端口列表
	PortList []*Port `json:"PortList,omitempty" xml:"PortList,omitempty" type:"Repeated"`
}

func (s PortCollection) String() string {
	return tea.Prettify(s)
}

func (s PortCollection) GoString() string {
	return s.String()
}

func (s *PortCollection) SetGmtCreate(v string) *PortCollection {
	s.GmtCreate = &v
	return s
}

func (s *PortCollection) SetGmtModified(v string) *PortCollection {
	s.GmtModified = &v
	return s
}

func (s *PortCollection) SetPortCollectionDescription(v string) *PortCollection {
	s.PortCollectionDescription = &v
	return s
}

func (s *PortCollection) SetPortCollectionId(v string) *PortCollection {
	s.PortCollectionId = &v
	return s
}

func (s *PortCollection) SetPortCollectionName(v string) *PortCollection {
	s.PortCollectionName = &v
	return s
}

func (s *PortCollection) SetPortList(v []*Port) *PortCollection {
	s.PortList = v
	return s
}

type ConnectionPolicy struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 资源一级ID
	ConnectionPolicyId *string `json:"ConnectionPolicyId,omitempty" xml:"ConnectionPolicyId,omitempty"`
	// 架构迭代uid
	NetworkArchitectureIterationId *string `json:"NetworkArchitectureIterationId,omitempty" xml:"NetworkArchitectureIterationId,omitempty"`
	// 连接策略名字
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 连接策略算法
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// 连接数
	LinkCount *int64 `json:"LinkCount,omitempty" xml:"LinkCount,omitempty"`
	// 上联模块uid
	UplinkModelId *string `json:"UplinkModelId,omitempty" xml:"UplinkModelId,omitempty"`
	// 下联模块uid
	DownlinkModuleId *string `json:"DownlinkModuleId,omitempty" xml:"DownlinkModuleId,omitempty"`
	// 上联设备uid
	UplinkDeviceId *string `json:"UplinkDeviceId,omitempty" xml:"UplinkDeviceId,omitempty"`
	// 下联设备uid
	DownlinkDeviceId *string `json:"DownlinkDeviceId,omitempty" xml:"DownlinkDeviceId,omitempty"`
	// 连接策略uid
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ConnectionPolicy) String() string {
	return tea.Prettify(s)
}

func (s ConnectionPolicy) GoString() string {
	return s.String()
}

func (s *ConnectionPolicy) SetCreateTime(v string) *ConnectionPolicy {
	s.CreateTime = &v
	return s
}

func (s *ConnectionPolicy) SetConnectionPolicyId(v string) *ConnectionPolicy {
	s.ConnectionPolicyId = &v
	return s
}

func (s *ConnectionPolicy) SetNetworkArchitectureIterationId(v string) *ConnectionPolicy {
	s.NetworkArchitectureIterationId = &v
	return s
}

func (s *ConnectionPolicy) SetName(v string) *ConnectionPolicy {
	s.Name = &v
	return s
}

func (s *ConnectionPolicy) SetAlgorithm(v string) *ConnectionPolicy {
	s.Algorithm = &v
	return s
}

func (s *ConnectionPolicy) SetLinkCount(v int64) *ConnectionPolicy {
	s.LinkCount = &v
	return s
}

func (s *ConnectionPolicy) SetUplinkModelId(v string) *ConnectionPolicy {
	s.UplinkModelId = &v
	return s
}

func (s *ConnectionPolicy) SetDownlinkModuleId(v string) *ConnectionPolicy {
	s.DownlinkModuleId = &v
	return s
}

func (s *ConnectionPolicy) SetUplinkDeviceId(v string) *ConnectionPolicy {
	s.UplinkDeviceId = &v
	return s
}

func (s *ConnectionPolicy) SetDownlinkDeviceId(v string) *ConnectionPolicy {
	s.DownlinkDeviceId = &v
	return s
}

func (s *ConnectionPolicy) SetId(v string) *ConnectionPolicy {
	s.Id = &v
	return s
}

func (s *ConnectionPolicy) SetUpdateTime(v string) *ConnectionPolicy {
	s.UpdateTime = &v
	return s
}

type DeviceFormProperty struct {
	// 属性描述
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 属性关键词
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// 前端界面控件占位符文字
	Placeholder *bool `json:"Placeholder,omitempty" xml:"Placeholder,omitempty"`
	// 属性是否必填
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// 属性是否作为界面查询条件
	SearchSupported *bool `json:"SearchSupported,omitempty" xml:"SearchSupported,omitempty"`
	// 属性展示的次序
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// 前端界面是否展示为表格列
	TableVisible *bool `json:"TableVisible,omitempty" xml:"TableVisible,omitempty"`
	// 属性是否需要唯一检查
	Uniqueness *bool `json:"Uniqueness,omitempty" xml:"Uniqueness,omitempty"`
	// 属性值来源具体的方式
	ValueReference *string `json:"ValueReference,omitempty" xml:"ValueReference,omitempty"`
	// 属性值来源类型：枚举、接口等
	ValueSource *string `json:"ValueSource,omitempty" xml:"ValueSource,omitempty"`
	// 属性类型，JSON或者分隔符
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s DeviceFormProperty) String() string {
	return tea.Prettify(s)
}

func (s DeviceFormProperty) GoString() string {
	return s.String()
}

func (s *DeviceFormProperty) SetContent(v string) *DeviceFormProperty {
	s.Content = &v
	return s
}

func (s *DeviceFormProperty) SetKeyword(v string) *DeviceFormProperty {
	s.Keyword = &v
	return s
}

func (s *DeviceFormProperty) SetPlaceholder(v bool) *DeviceFormProperty {
	s.Placeholder = &v
	return s
}

func (s *DeviceFormProperty) SetRequired(v bool) *DeviceFormProperty {
	s.Required = &v
	return s
}

func (s *DeviceFormProperty) SetSearchSupported(v bool) *DeviceFormProperty {
	s.SearchSupported = &v
	return s
}

func (s *DeviceFormProperty) SetSequence(v int32) *DeviceFormProperty {
	s.Sequence = &v
	return s
}

func (s *DeviceFormProperty) SetTableVisible(v bool) *DeviceFormProperty {
	s.TableVisible = &v
	return s
}

func (s *DeviceFormProperty) SetUniqueness(v bool) *DeviceFormProperty {
	s.Uniqueness = &v
	return s
}

func (s *DeviceFormProperty) SetValueReference(v string) *DeviceFormProperty {
	s.ValueReference = &v
	return s
}

func (s *DeviceFormProperty) SetValueSource(v string) *DeviceFormProperty {
	s.ValueSource = &v
	return s
}

func (s *DeviceFormProperty) SetValueType(v string) *DeviceFormProperty {
	s.ValueType = &v
	return s
}

type CreateConfigurationSpecificationRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 配置规范名字
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// 架构类型
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 配置规范内容
	SpecificationContent *string `json:"SpecificationContent,omitempty" xml:"SpecificationContent,omitempty"`
	// 相关变量
	RelatedVariate [][]byte `json:"RelatedVariate,omitempty" xml:"RelatedVariate,omitempty" type:"Repeated"`
}

func (s CreateConfigurationSpecificationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigurationSpecificationRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigurationSpecificationRequest) SetInstanceId(v string) *CreateConfigurationSpecificationRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateConfigurationSpecificationRequest) SetClientToken(v string) *CreateConfigurationSpecificationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateConfigurationSpecificationRequest) SetSpecificationName(v string) *CreateConfigurationSpecificationRequest {
	s.SpecificationName = &v
	return s
}

func (s *CreateConfigurationSpecificationRequest) SetArchitecture(v string) *CreateConfigurationSpecificationRequest {
	s.Architecture = &v
	return s
}

func (s *CreateConfigurationSpecificationRequest) SetRole(v string) *CreateConfigurationSpecificationRequest {
	s.Role = &v
	return s
}

func (s *CreateConfigurationSpecificationRequest) SetVendor(v string) *CreateConfigurationSpecificationRequest {
	s.Vendor = &v
	return s
}

func (s *CreateConfigurationSpecificationRequest) SetModel(v string) *CreateConfigurationSpecificationRequest {
	s.Model = &v
	return s
}

func (s *CreateConfigurationSpecificationRequest) SetSpecificationContent(v string) *CreateConfigurationSpecificationRequest {
	s.SpecificationContent = &v
	return s
}

func (s *CreateConfigurationSpecificationRequest) SetRelatedVariate(v [][]byte) *CreateConfigurationSpecificationRequest {
	s.RelatedVariate = v
	return s
}

type CreateConfigurationSpecificationShrinkRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 配置规范名字
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// 架构类型
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 配置规范内容
	SpecificationContent *string `json:"SpecificationContent,omitempty" xml:"SpecificationContent,omitempty"`
	// 相关变量
	RelatedVariateShrink *string `json:"RelatedVariate,omitempty" xml:"RelatedVariate,omitempty"`
}

func (s CreateConfigurationSpecificationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigurationSpecificationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigurationSpecificationShrinkRequest) SetInstanceId(v string) *CreateConfigurationSpecificationShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateConfigurationSpecificationShrinkRequest) SetClientToken(v string) *CreateConfigurationSpecificationShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateConfigurationSpecificationShrinkRequest) SetSpecificationName(v string) *CreateConfigurationSpecificationShrinkRequest {
	s.SpecificationName = &v
	return s
}

func (s *CreateConfigurationSpecificationShrinkRequest) SetArchitecture(v string) *CreateConfigurationSpecificationShrinkRequest {
	s.Architecture = &v
	return s
}

func (s *CreateConfigurationSpecificationShrinkRequest) SetRole(v string) *CreateConfigurationSpecificationShrinkRequest {
	s.Role = &v
	return s
}

func (s *CreateConfigurationSpecificationShrinkRequest) SetVendor(v string) *CreateConfigurationSpecificationShrinkRequest {
	s.Vendor = &v
	return s
}

func (s *CreateConfigurationSpecificationShrinkRequest) SetModel(v string) *CreateConfigurationSpecificationShrinkRequest {
	s.Model = &v
	return s
}

func (s *CreateConfigurationSpecificationShrinkRequest) SetSpecificationContent(v string) *CreateConfigurationSpecificationShrinkRequest {
	s.SpecificationContent = &v
	return s
}

func (s *CreateConfigurationSpecificationShrinkRequest) SetRelatedVariateShrink(v string) *CreateConfigurationSpecificationShrinkRequest {
	s.RelatedVariateShrink = &v
	return s
}

type CreateConfigurationSpecificationResponseBody struct {
	// 资源uuid
	ConfigurationSpecificationId *string `json:"ConfigurationSpecificationId,omitempty" xml:"ConfigurationSpecificationId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateConfigurationSpecificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigurationSpecificationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConfigurationSpecificationResponseBody) SetConfigurationSpecificationId(v string) *CreateConfigurationSpecificationResponseBody {
	s.ConfigurationSpecificationId = &v
	return s
}

func (s *CreateConfigurationSpecificationResponseBody) SetRequestId(v string) *CreateConfigurationSpecificationResponseBody {
	s.RequestId = &v
	return s
}

type CreateConfigurationSpecificationResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateConfigurationSpecificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateConfigurationSpecificationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigurationSpecificationResponse) GoString() string {
	return s.String()
}

func (s *CreateConfigurationSpecificationResponse) SetHeaders(v map[string]*string) *CreateConfigurationSpecificationResponse {
	s.Headers = v
	return s
}

func (s *CreateConfigurationSpecificationResponse) SetBody(v *CreateConfigurationSpecificationResponseBody) *CreateConfigurationSpecificationResponse {
	s.Body = v
	return s
}

type CreateIpBlockRequest struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源一级ID
	IpBlockId *string `json:"IpBlockId,omitempty" xml:"IpBlockId,omitempty"`
	// IP段地址
	Block *string `json:"Block,omitempty" xml:"Block,omitempty"`
	// IP段掩码
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// 父地址段UID
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// 公网私网标志 PUBLIC PRIVATE
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// 地址类别 IPV4
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 地址状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 业务类型UID
	BusinessTypeId *string `json:"BusinessTypeId,omitempty" xml:"BusinessTypeId,omitempty"`
	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// 园区名称
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	// 备份设备名称
	BackupDeviceName *string `json:"BackupDeviceName,omitempty" xml:"BackupDeviceName,omitempty"`
	// 公网地址类型 INC GUEST VIP
	NetBusiness *string `json:"NetBusiness,omitempty" xml:"NetBusiness,omitempty"`
	// IP归属
	Ownership *string `json:"Ownership,omitempty" xml:"Ownership,omitempty"`
	// IP用途
	Application *string `json:"Application,omitempty" xml:"Application,omitempty"`
	// 备注
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 是否保留父段 true 是 false 否
	ReserveParentBlock *string `json:"ReserveParentBlock,omitempty" xml:"ReserveParentBlock,omitempty"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateIpBlockRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIpBlockRequest) GoString() string {
	return s.String()
}

func (s *CreateIpBlockRequest) SetInstanceId(v string) *CreateIpBlockRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateIpBlockRequest) SetIpBlockId(v string) *CreateIpBlockRequest {
	s.IpBlockId = &v
	return s
}

func (s *CreateIpBlockRequest) SetBlock(v string) *CreateIpBlockRequest {
	s.Block = &v
	return s
}

func (s *CreateIpBlockRequest) SetMask(v string) *CreateIpBlockRequest {
	s.Mask = &v
	return s
}

func (s *CreateIpBlockRequest) SetParentId(v string) *CreateIpBlockRequest {
	s.ParentId = &v
	return s
}

func (s *CreateIpBlockRequest) SetNetType(v string) *CreateIpBlockRequest {
	s.NetType = &v
	return s
}

func (s *CreateIpBlockRequest) SetCategory(v string) *CreateIpBlockRequest {
	s.Category = &v
	return s
}

func (s *CreateIpBlockRequest) SetStatus(v string) *CreateIpBlockRequest {
	s.Status = &v
	return s
}

func (s *CreateIpBlockRequest) SetBusinessTypeId(v string) *CreateIpBlockRequest {
	s.BusinessTypeId = &v
	return s
}

func (s *CreateIpBlockRequest) SetDeviceName(v string) *CreateIpBlockRequest {
	s.DeviceName = &v
	return s
}

func (s *CreateIpBlockRequest) SetZoneName(v string) *CreateIpBlockRequest {
	s.ZoneName = &v
	return s
}

func (s *CreateIpBlockRequest) SetBackupDeviceName(v string) *CreateIpBlockRequest {
	s.BackupDeviceName = &v
	return s
}

func (s *CreateIpBlockRequest) SetNetBusiness(v string) *CreateIpBlockRequest {
	s.NetBusiness = &v
	return s
}

func (s *CreateIpBlockRequest) SetOwnership(v string) *CreateIpBlockRequest {
	s.Ownership = &v
	return s
}

func (s *CreateIpBlockRequest) SetApplication(v string) *CreateIpBlockRequest {
	s.Application = &v
	return s
}

func (s *CreateIpBlockRequest) SetDescription(v string) *CreateIpBlockRequest {
	s.Description = &v
	return s
}

func (s *CreateIpBlockRequest) SetReserveParentBlock(v string) *CreateIpBlockRequest {
	s.ReserveParentBlock = &v
	return s
}

func (s *CreateIpBlockRequest) SetClientToken(v string) *CreateIpBlockRequest {
	s.ClientToken = &v
	return s
}

type CreateIpBlockResponseBody struct {
	// 资源实例ID，如ECS实例的创建接口CreateInstance应返回InstanceId。
	IpBlockId *string `json:"IpBlockId,omitempty" xml:"IpBlockId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIpBlockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIpBlockResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpBlockResponseBody) SetIpBlockId(v string) *CreateIpBlockResponseBody {
	s.IpBlockId = &v
	return s
}

func (s *CreateIpBlockResponseBody) SetRequestId(v string) *CreateIpBlockResponseBody {
	s.RequestId = &v
	return s
}

type CreateIpBlockResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateIpBlockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIpBlockResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIpBlockResponse) GoString() string {
	return s.String()
}

func (s *CreateIpBlockResponse) SetHeaders(v map[string]*string) *CreateIpBlockResponse {
	s.Headers = v
	return s
}

func (s *CreateIpBlockResponse) SetBody(v *CreateIpBlockResponseBody) *CreateIpBlockResponse {
	s.Body = v
	return s
}

type UpdateConfigurationVariateRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源一级ID
	ConfigurationVariateId *string `json:"ConfigurationVariateId,omitempty" xml:"ConfigurationVariateId,omitempty"`
	// 变量名字
	VariateName *string `json:"VariateName,omitempty" xml:"VariateName,omitempty"`
	// 描述变量
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// python转换函数
	FormatFunction *string `json:"FormatFunction,omitempty" xml:"FormatFunction,omitempty"`
}

func (s UpdateConfigurationVariateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigurationVariateRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigurationVariateRequest) SetInstanceId(v string) *UpdateConfigurationVariateRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateConfigurationVariateRequest) SetConfigurationVariateId(v string) *UpdateConfigurationVariateRequest {
	s.ConfigurationVariateId = &v
	return s
}

func (s *UpdateConfigurationVariateRequest) SetVariateName(v string) *UpdateConfigurationVariateRequest {
	s.VariateName = &v
	return s
}

func (s *UpdateConfigurationVariateRequest) SetComment(v string) *UpdateConfigurationVariateRequest {
	s.Comment = &v
	return s
}

func (s *UpdateConfigurationVariateRequest) SetFormatFunction(v string) *UpdateConfigurationVariateRequest {
	s.FormatFunction = &v
	return s
}

type UpdateConfigurationVariateResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConfigurationVariateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigurationVariateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigurationVariateResponseBody) SetRequestId(v string) *UpdateConfigurationVariateResponseBody {
	s.RequestId = &v
	return s
}

type UpdateConfigurationVariateResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateConfigurationVariateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateConfigurationVariateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigurationVariateResponse) GoString() string {
	return s.String()
}

func (s *UpdateConfigurationVariateResponse) SetHeaders(v map[string]*string) *UpdateConfigurationVariateResponse {
	s.Headers = v
	return s
}

func (s *UpdateConfigurationVariateResponse) SetBody(v *UpdateConfigurationVariateResponseBody) *UpdateConfigurationVariateResponse {
	s.Body = v
	return s
}

type GetScheduleTypeRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例 ID。
	ScheduleTypeId *string `json:"ScheduleTypeId,omitempty" xml:"ScheduleTypeId,omitempty"`
}

func (s GetScheduleTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleTypeRequest) GoString() string {
	return s.String()
}

func (s *GetScheduleTypeRequest) SetInstanceId(v string) *GetScheduleTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetScheduleTypeRequest) SetScheduleTypeId(v string) *GetScheduleTypeRequest {
	s.ScheduleTypeId = &v
	return s
}

type GetScheduleTypeResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 值班类型对象
	ScheduleType *GetScheduleTypeResponseBodyScheduleType `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty" type:"Struct"`
}

func (s GetScheduleTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetScheduleTypeResponseBody) SetRequestId(v string) *GetScheduleTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScheduleTypeResponseBody) SetScheduleType(v *GetScheduleTypeResponseBodyScheduleType) *GetScheduleTypeResponseBody {
	s.ScheduleType = v
	return s
}

type GetScheduleTypeResponseBodyScheduleType struct {
	// 值班类型状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 资源一级ID
	ScheduleTypeId *string `json:"ScheduleTypeId,omitempty" xml:"ScheduleTypeId,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 值班类型value
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// relatedWorker
	RelatedWorker []*string `json:"RelatedWorker,omitempty" xml:"RelatedWorker,omitempty" type:"Repeated"`
}

func (s GetScheduleTypeResponseBodyScheduleType) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleTypeResponseBodyScheduleType) GoString() string {
	return s.String()
}

func (s *GetScheduleTypeResponseBodyScheduleType) SetStatus(v string) *GetScheduleTypeResponseBodyScheduleType {
	s.Status = &v
	return s
}

func (s *GetScheduleTypeResponseBodyScheduleType) SetScheduleTypeId(v string) *GetScheduleTypeResponseBodyScheduleType {
	s.ScheduleTypeId = &v
	return s
}

func (s *GetScheduleTypeResponseBodyScheduleType) SetCreateTime(v string) *GetScheduleTypeResponseBodyScheduleType {
	s.CreateTime = &v
	return s
}

func (s *GetScheduleTypeResponseBodyScheduleType) SetUpdateTime(v string) *GetScheduleTypeResponseBodyScheduleType {
	s.UpdateTime = &v
	return s
}

func (s *GetScheduleTypeResponseBodyScheduleType) SetScheduleType(v string) *GetScheduleTypeResponseBodyScheduleType {
	s.ScheduleType = &v
	return s
}

func (s *GetScheduleTypeResponseBodyScheduleType) SetRelatedWorker(v []*string) *GetScheduleTypeResponseBodyScheduleType {
	s.RelatedWorker = v
	return s
}

type GetScheduleTypeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetScheduleTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetScheduleTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleTypeResponse) GoString() string {
	return s.String()
}

func (s *GetScheduleTypeResponse) SetHeaders(v map[string]*string) *GetScheduleTypeResponse {
	s.Headers = v
	return s
}

func (s *GetScheduleTypeResponse) SetBody(v *GetScheduleTypeResponseBody) *GetScheduleTypeResponse {
	s.Body = v
	return s
}

type DeleteDeviceRequest struct {
	// 实例 ID。
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeviceRequest) SetDeviceId(v string) *DeleteDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *DeleteDeviceRequest) SetInstanceId(v string) *DeleteDeviceRequest {
	s.InstanceId = &v
	return s
}

type DeleteDeviceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeviceResponseBody) SetRequestId(v string) *DeleteDeviceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDeviceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeviceResponse) SetHeaders(v map[string]*string) *DeleteDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeviceResponse) SetBody(v *DeleteDeviceResponseBody) *DeleteDeviceResponse {
	s.Body = v
	return s
}

type ListDeviceValuesRequest struct {
	// 设备形态ID
	DeviceFormId *string `json:"DeviceFormId,omitempty" xml:"DeviceFormId,omitempty"`
	// 设备形态名称
	DeviceFormName *string `json:"DeviceFormName,omitempty" xml:"DeviceFormName,omitempty"`
	// 查询属性主键
	AttributeKeyword *string `json:"AttributeKeyword,omitempty" xml:"AttributeKeyword,omitempty"`
	// 查询属性对应JSON中主键
	AttributeGroup *string `json:"AttributeGroup,omitempty" xml:"AttributeGroup,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListDeviceValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceValuesRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceValuesRequest) SetDeviceFormId(v string) *ListDeviceValuesRequest {
	s.DeviceFormId = &v
	return s
}

func (s *ListDeviceValuesRequest) SetDeviceFormName(v string) *ListDeviceValuesRequest {
	s.DeviceFormName = &v
	return s
}

func (s *ListDeviceValuesRequest) SetAttributeKeyword(v string) *ListDeviceValuesRequest {
	s.AttributeKeyword = &v
	return s
}

func (s *ListDeviceValuesRequest) SetAttributeGroup(v string) *ListDeviceValuesRequest {
	s.AttributeGroup = &v
	return s
}

func (s *ListDeviceValuesRequest) SetInstanceId(v string) *ListDeviceValuesRequest {
	s.InstanceId = &v
	return s
}

type ListDeviceValuesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 数组，返回示例目录。
	DeviceValues []*string `json:"DeviceValues,omitempty" xml:"DeviceValues,omitempty" type:"Repeated"`
}

func (s ListDeviceValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceValuesResponseBody) SetRequestId(v string) *ListDeviceValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceValuesResponseBody) SetDeviceValues(v []*string) *ListDeviceValuesResponseBody {
	s.DeviceValues = v
	return s
}

type ListDeviceValuesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeviceValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeviceValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceValuesResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceValuesResponse) SetHeaders(v map[string]*string) *ListDeviceValuesResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceValuesResponse) SetBody(v *ListDeviceValuesResponseBody) *ListDeviceValuesResponse {
	s.Body = v
	return s
}

type GetIpBlockRecordRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例 ID。
	IpBlockRecordId *string `json:"IpBlockRecordId,omitempty" xml:"IpBlockRecordId,omitempty"`
}

func (s GetIpBlockRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIpBlockRecordRequest) GoString() string {
	return s.String()
}

func (s *GetIpBlockRecordRequest) SetInstanceId(v string) *GetIpBlockRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *GetIpBlockRecordRequest) SetIpBlockRecordId(v string) *GetIpBlockRecordRequest {
	s.IpBlockRecordId = &v
	return s
}

type GetIpBlockRecordResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 工单类型对象
	IpBlockRecord *GetIpBlockRecordResponseBodyIpBlockRecord `json:"IpBlockRecord,omitempty" xml:"IpBlockRecord,omitempty" type:"Struct"`
}

func (s GetIpBlockRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIpBlockRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetIpBlockRecordResponseBody) SetRequestId(v string) *GetIpBlockRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIpBlockRecordResponseBody) SetIpBlockRecord(v *GetIpBlockRecordResponseBodyIpBlockRecord) *GetIpBlockRecordResponseBody {
	s.IpBlockRecord = v
	return s
}

type GetIpBlockRecordResponseBodyIpBlockRecord struct {
	// 工单uuid
	IpBlockRecordId *string `json:"IpBlockRecordId,omitempty" xml:"IpBlockRecordId,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 2020-12-22 10:39:17
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 工单名称
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 创建人
	Creator *int64 `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// 工号
	WorkNo *string `json:"WorkNo,omitempty" xml:"WorkNo,omitempty"`
	// 公网私网标志 PUBLIC PRIVATE
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// 工单状态 running complete fail cancel lock approving
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 园区名称
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	// 工单备注
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 工单详情
	Detail []*GetIpBlockRecordResponseBodyIpBlockRecordDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// 工单类型 Apply 申请工单 Recycle 释放工单
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// 要释放的IP段
	IpBlockCode []*string `json:"IpBlockCode,omitempty" xml:"IpBlockCode,omitempty" type:"Repeated"`
}

func (s GetIpBlockRecordResponseBodyIpBlockRecord) String() string {
	return tea.Prettify(s)
}

func (s GetIpBlockRecordResponseBodyIpBlockRecord) GoString() string {
	return s.String()
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecord) SetIpBlockRecordId(v string) *GetIpBlockRecordResponseBodyIpBlockRecord {
	s.IpBlockRecordId = &v
	return s
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecord) SetCreateTime(v string) *GetIpBlockRecordResponseBodyIpBlockRecord {
	s.CreateTime = &v
	return s
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecord) SetUpdateTime(v string) *GetIpBlockRecordResponseBodyIpBlockRecord {
	s.UpdateTime = &v
	return s
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecord) SetTitle(v string) *GetIpBlockRecordResponseBodyIpBlockRecord {
	s.Title = &v
	return s
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecord) SetCreator(v int64) *GetIpBlockRecordResponseBodyIpBlockRecord {
	s.Creator = &v
	return s
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecord) SetWorkNo(v string) *GetIpBlockRecordResponseBodyIpBlockRecord {
	s.WorkNo = &v
	return s
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecord) SetNetType(v string) *GetIpBlockRecordResponseBodyIpBlockRecord {
	s.NetType = &v
	return s
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecord) SetStatus(v string) *GetIpBlockRecordResponseBodyIpBlockRecord {
	s.Status = &v
	return s
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecord) SetZoneName(v string) *GetIpBlockRecordResponseBodyIpBlockRecord {
	s.ZoneName = &v
	return s
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecord) SetDescription(v string) *GetIpBlockRecordResponseBodyIpBlockRecord {
	s.Description = &v
	return s
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecord) SetDetail(v []*GetIpBlockRecordResponseBodyIpBlockRecordDetail) *GetIpBlockRecordResponseBodyIpBlockRecord {
	s.Detail = v
	return s
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecord) SetRecordType(v string) *GetIpBlockRecordResponseBodyIpBlockRecord {
	s.RecordType = &v
	return s
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecord) SetIpBlockCode(v []*string) *GetIpBlockRecordResponseBodyIpBlockRecord {
	s.IpBlockCode = v
	return s
}

type GetIpBlockRecordResponseBodyIpBlockRecordDetail struct {
	// 业务类型
	BusinessTypeName *string `json:"BusinessTypeName,omitempty" xml:"BusinessTypeName,omitempty"`
	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// 网关
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// 父地址段列表
	ParentIpBlocks []*string `json:"ParentIpBlocks,omitempty" xml:"ParentIpBlocks,omitempty" type:"Repeated"`
	// 园区层级
	ZoneLayer []*GetIpBlockRecordResponseBodyIpBlockRecordDetailZoneLayer `json:"ZoneLayer,omitempty" xml:"ZoneLayer,omitempty" type:"Repeated"`
	// 地址段
	IpBlockCode *string `json:"IpBlockCode,omitempty" xml:"IpBlockCode,omitempty"`
}

func (s GetIpBlockRecordResponseBodyIpBlockRecordDetail) String() string {
	return tea.Prettify(s)
}

func (s GetIpBlockRecordResponseBodyIpBlockRecordDetail) GoString() string {
	return s.String()
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecordDetail) SetBusinessTypeName(v string) *GetIpBlockRecordResponseBodyIpBlockRecordDetail {
	s.BusinessTypeName = &v
	return s
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecordDetail) SetDeviceName(v string) *GetIpBlockRecordResponseBodyIpBlockRecordDetail {
	s.DeviceName = &v
	return s
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecordDetail) SetGateway(v string) *GetIpBlockRecordResponseBodyIpBlockRecordDetail {
	s.Gateway = &v
	return s
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecordDetail) SetParentIpBlocks(v []*string) *GetIpBlockRecordResponseBodyIpBlockRecordDetail {
	s.ParentIpBlocks = v
	return s
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecordDetail) SetZoneLayer(v []*GetIpBlockRecordResponseBodyIpBlockRecordDetailZoneLayer) *GetIpBlockRecordResponseBodyIpBlockRecordDetail {
	s.ZoneLayer = v
	return s
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecordDetail) SetIpBlockCode(v string) *GetIpBlockRecordResponseBodyIpBlockRecordDetail {
	s.IpBlockCode = &v
	return s
}

type GetIpBlockRecordResponseBodyIpBlockRecordDetailZoneLayer struct {
	// 园区层级名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 园区层级值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetIpBlockRecordResponseBodyIpBlockRecordDetailZoneLayer) String() string {
	return tea.Prettify(s)
}

func (s GetIpBlockRecordResponseBodyIpBlockRecordDetailZoneLayer) GoString() string {
	return s.String()
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecordDetailZoneLayer) SetName(v string) *GetIpBlockRecordResponseBodyIpBlockRecordDetailZoneLayer {
	s.Name = &v
	return s
}

func (s *GetIpBlockRecordResponseBodyIpBlockRecordDetailZoneLayer) SetValue(v string) *GetIpBlockRecordResponseBodyIpBlockRecordDetailZoneLayer {
	s.Value = &v
	return s
}

type GetIpBlockRecordResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetIpBlockRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIpBlockRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIpBlockRecordResponse) GoString() string {
	return s.String()
}

func (s *GetIpBlockRecordResponse) SetHeaders(v map[string]*string) *GetIpBlockRecordResponse {
	s.Headers = v
	return s
}

func (s *GetIpBlockRecordResponse) SetBody(v *GetIpBlockRecordResponseBody) *GetIpBlockRecordResponse {
	s.Body = v
	return s
}

type ListSpaceModelsRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 物理空间类型
	SpaceType *string `json:"SpaceType,omitempty" xml:"SpaceType,omitempty"`
	// 物理空间状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSpaceModelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSpaceModelsRequest) GoString() string {
	return s.String()
}

func (s *ListSpaceModelsRequest) SetInstanceId(v string) *ListSpaceModelsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSpaceModelsRequest) SetMaxResults(v int32) *ListSpaceModelsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSpaceModelsRequest) SetNextToken(v string) *ListSpaceModelsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSpaceModelsRequest) SetSpaceType(v string) *ListSpaceModelsRequest {
	s.SpaceType = &v
	return s
}

func (s *ListSpaceModelsRequest) SetStatus(v string) *ListSpaceModelsRequest {
	s.Status = &v
	return s
}

type ListSpaceModelsResponseBody struct {
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据量
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 数组，返回示例目录。
	SpaceModel []*ListSpaceModelsResponseBodySpaceModel `json:"SpaceModel,omitempty" xml:"SpaceModel,omitempty" type:"Repeated"`
}

func (s ListSpaceModelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSpaceModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSpaceModelsResponseBody) SetTotalCount(v int32) *ListSpaceModelsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSpaceModelsResponseBody) SetRequestId(v string) *ListSpaceModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSpaceModelsResponseBody) SetNextToken(v string) *ListSpaceModelsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSpaceModelsResponseBody) SetMaxResults(v int64) *ListSpaceModelsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSpaceModelsResponseBody) SetSpaceModel(v []*ListSpaceModelsResponseBodySpaceModel) *ListSpaceModelsResponseBody {
	s.SpaceModel = v
	return s
}

type ListSpaceModelsResponseBodySpaceModel struct {
	// 物理空间类型
	SpaceType *string `json:"SpaceType,omitempty" xml:"SpaceType,omitempty"`
	// 层级
	Sort []*ListSpaceModelsResponseBodySpaceModelSort `json:"Sort,omitempty" xml:"Sort,omitempty" type:"Repeated"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 资源一级ID
	SpaceModelId *string `json:"SpaceModelId,omitempty" xml:"SpaceModelId,omitempty"`
	// 模型状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListSpaceModelsResponseBodySpaceModel) String() string {
	return tea.Prettify(s)
}

func (s ListSpaceModelsResponseBodySpaceModel) GoString() string {
	return s.String()
}

func (s *ListSpaceModelsResponseBodySpaceModel) SetSpaceType(v string) *ListSpaceModelsResponseBodySpaceModel {
	s.SpaceType = &v
	return s
}

func (s *ListSpaceModelsResponseBodySpaceModel) SetSort(v []*ListSpaceModelsResponseBodySpaceModelSort) *ListSpaceModelsResponseBodySpaceModel {
	s.Sort = v
	return s
}

func (s *ListSpaceModelsResponseBodySpaceModel) SetCreateTime(v string) *ListSpaceModelsResponseBodySpaceModel {
	s.CreateTime = &v
	return s
}

func (s *ListSpaceModelsResponseBodySpaceModel) SetSpaceModelId(v string) *ListSpaceModelsResponseBodySpaceModel {
	s.SpaceModelId = &v
	return s
}

func (s *ListSpaceModelsResponseBodySpaceModel) SetStatus(v string) *ListSpaceModelsResponseBodySpaceModel {
	s.Status = &v
	return s
}

func (s *ListSpaceModelsResponseBodySpaceModel) SetUpdateTime(v string) *ListSpaceModelsResponseBodySpaceModel {
	s.UpdateTime = &v
	return s
}

type ListSpaceModelsResponseBodySpaceModelSort struct {
	// 层级名称
	LevelName *string `json:"LevelName,omitempty" xml:"LevelName,omitempty"`
	// 层级
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s ListSpaceModelsResponseBodySpaceModelSort) String() string {
	return tea.Prettify(s)
}

func (s ListSpaceModelsResponseBodySpaceModelSort) GoString() string {
	return s.String()
}

func (s *ListSpaceModelsResponseBodySpaceModelSort) SetLevelName(v string) *ListSpaceModelsResponseBodySpaceModelSort {
	s.LevelName = &v
	return s
}

func (s *ListSpaceModelsResponseBodySpaceModelSort) SetLevel(v int64) *ListSpaceModelsResponseBodySpaceModelSort {
	s.Level = &v
	return s
}

type ListSpaceModelsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSpaceModelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSpaceModelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSpaceModelsResponse) GoString() string {
	return s.String()
}

func (s *ListSpaceModelsResponse) SetHeaders(v map[string]*string) *ListSpaceModelsResponse {
	s.Headers = v
	return s
}

func (s *ListSpaceModelsResponse) SetBody(v *ListSpaceModelsResponseBody) *ListSpaceModelsResponse {
	s.Body = v
	return s
}

type ListDevicePropertiesRequest struct {
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 设备形态ID
	DeviceFormId *string `json:"DeviceFormId,omitempty" xml:"DeviceFormId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListDevicePropertiesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDevicePropertiesRequest) GoString() string {
	return s.String()
}

func (s *ListDevicePropertiesRequest) SetMaxResults(v int32) *ListDevicePropertiesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDevicePropertiesRequest) SetNextToken(v string) *ListDevicePropertiesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDevicePropertiesRequest) SetDeviceFormId(v string) *ListDevicePropertiesRequest {
	s.DeviceFormId = &v
	return s
}

func (s *ListDevicePropertiesRequest) SetInstanceId(v string) *ListDevicePropertiesRequest {
	s.InstanceId = &v
	return s
}

type ListDevicePropertiesResponseBody struct {
	// 数组，返回示例目录。
	DeviceProperties []*ListDevicePropertiesResponseBodyDeviceProperties `json:"DeviceProperties,omitempty" xml:"DeviceProperties,omitempty" type:"Repeated"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 每页数量。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListDevicePropertiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevicePropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevicePropertiesResponseBody) SetDeviceProperties(v []*ListDevicePropertiesResponseBodyDeviceProperties) *ListDevicePropertiesResponseBody {
	s.DeviceProperties = v
	return s
}

func (s *ListDevicePropertiesResponseBody) SetNextToken(v int32) *ListDevicePropertiesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDevicePropertiesResponseBody) SetRequestId(v string) *ListDevicePropertiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDevicePropertiesResponseBody) SetTotalCount(v int32) *ListDevicePropertiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDevicePropertiesResponseBody) SetMaxResults(v int32) *ListDevicePropertiesResponseBody {
	s.MaxResults = &v
	return s
}

type ListDevicePropertiesResponseBodyDeviceProperties struct {
	// 设备属性ID
	DevicePropertyId *string `json:"DevicePropertyId,omitempty" xml:"DevicePropertyId,omitempty"`
	// 设备形态ID
	DeviceFormId *string `json:"DeviceFormId,omitempty" xml:"DeviceFormId,omitempty"`
	// 设备形态名称
	DeviceFormName *string `json:"DeviceFormName,omitempty" xml:"DeviceFormName,omitempty"`
	// 属性名称
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	// 属性主键
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// 属性格式
	PropertyFormat *string `json:"PropertyFormat,omitempty" xml:"PropertyFormat,omitempty"`
	// 属性内容
	PropertyContent *string `json:"PropertyContent,omitempty" xml:"PropertyContent,omitempty"`
	// 是否内置属性
	BuiltIn *bool `json:"BuiltIn,omitempty" xml:"BuiltIn,omitempty"`
}

func (s ListDevicePropertiesResponseBodyDeviceProperties) String() string {
	return tea.Prettify(s)
}

func (s ListDevicePropertiesResponseBodyDeviceProperties) GoString() string {
	return s.String()
}

func (s *ListDevicePropertiesResponseBodyDeviceProperties) SetDevicePropertyId(v string) *ListDevicePropertiesResponseBodyDeviceProperties {
	s.DevicePropertyId = &v
	return s
}

func (s *ListDevicePropertiesResponseBodyDeviceProperties) SetDeviceFormId(v string) *ListDevicePropertiesResponseBodyDeviceProperties {
	s.DeviceFormId = &v
	return s
}

func (s *ListDevicePropertiesResponseBodyDeviceProperties) SetDeviceFormName(v string) *ListDevicePropertiesResponseBodyDeviceProperties {
	s.DeviceFormName = &v
	return s
}

func (s *ListDevicePropertiesResponseBodyDeviceProperties) SetPropertyName(v string) *ListDevicePropertiesResponseBodyDeviceProperties {
	s.PropertyName = &v
	return s
}

func (s *ListDevicePropertiesResponseBodyDeviceProperties) SetPropertyKey(v string) *ListDevicePropertiesResponseBodyDeviceProperties {
	s.PropertyKey = &v
	return s
}

func (s *ListDevicePropertiesResponseBodyDeviceProperties) SetPropertyFormat(v string) *ListDevicePropertiesResponseBodyDeviceProperties {
	s.PropertyFormat = &v
	return s
}

func (s *ListDevicePropertiesResponseBodyDeviceProperties) SetPropertyContent(v string) *ListDevicePropertiesResponseBodyDeviceProperties {
	s.PropertyContent = &v
	return s
}

func (s *ListDevicePropertiesResponseBodyDeviceProperties) SetBuiltIn(v bool) *ListDevicePropertiesResponseBodyDeviceProperties {
	s.BuiltIn = &v
	return s
}

type ListDevicePropertiesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDevicePropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDevicePropertiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDevicePropertiesResponse) GoString() string {
	return s.String()
}

func (s *ListDevicePropertiesResponse) SetHeaders(v map[string]*string) *ListDevicePropertiesResponse {
	s.Headers = v
	return s
}

func (s *ListDevicePropertiesResponse) SetBody(v *ListDevicePropertiesResponseBody) *ListDevicePropertiesResponse {
	s.Body = v
	return s
}

type ListInspectionDevicesRequest struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model []*string `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
}

func (s ListInspectionDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInspectionDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListInspectionDevicesRequest) SetInstanceId(v string) *ListInspectionDevicesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInspectionDevicesRequest) SetRole(v string) *ListInspectionDevicesRequest {
	s.Role = &v
	return s
}

func (s *ListInspectionDevicesRequest) SetVendor(v string) *ListInspectionDevicesRequest {
	s.Vendor = &v
	return s
}

func (s *ListInspectionDevicesRequest) SetModel(v []*string) *ListInspectionDevicesRequest {
	s.Model = v
	return s
}

type ListInspectionDevicesShrinkRequest struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	ModelShrink *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s ListInspectionDevicesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInspectionDevicesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListInspectionDevicesShrinkRequest) SetInstanceId(v string) *ListInspectionDevicesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInspectionDevicesShrinkRequest) SetRole(v string) *ListInspectionDevicesShrinkRequest {
	s.Role = &v
	return s
}

func (s *ListInspectionDevicesShrinkRequest) SetVendor(v string) *ListInspectionDevicesShrinkRequest {
	s.Vendor = &v
	return s
}

func (s *ListInspectionDevicesShrinkRequest) SetModelShrink(v string) *ListInspectionDevicesShrinkRequest {
	s.ModelShrink = &v
	return s
}

type ListInspectionDevicesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// content
	InspectionScripts []*ListInspectionDevicesResponseBodyInspectionScripts `json:"InspectionScripts,omitempty" xml:"InspectionScripts,omitempty" type:"Repeated"`
}

func (s ListInspectionDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInspectionDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInspectionDevicesResponseBody) SetRequestId(v string) *ListInspectionDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInspectionDevicesResponseBody) SetInspectionScripts(v []*ListInspectionDevicesResponseBodyInspectionScripts) *ListInspectionDevicesResponseBody {
	s.InspectionScripts = v
	return s
}

type ListInspectionDevicesResponseBodyInspectionScripts struct {
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 主机名
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// 管理ip
	ManageIp *string `json:"ManageIp,omitempty" xml:"ManageIp,omitempty"`
	// 设备状态
	DeviceState *string `json:"DeviceState,omitempty" xml:"DeviceState,omitempty"`
	// 物理空间
	Space *string `json:"Space,omitempty" xml:"Space,omitempty"`
	// 设备id
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s ListInspectionDevicesResponseBodyInspectionScripts) String() string {
	return tea.Prettify(s)
}

func (s ListInspectionDevicesResponseBodyInspectionScripts) GoString() string {
	return s.String()
}

func (s *ListInspectionDevicesResponseBodyInspectionScripts) SetRole(v string) *ListInspectionDevicesResponseBodyInspectionScripts {
	s.Role = &v
	return s
}

func (s *ListInspectionDevicesResponseBodyInspectionScripts) SetVendor(v string) *ListInspectionDevicesResponseBodyInspectionScripts {
	s.Vendor = &v
	return s
}

func (s *ListInspectionDevicesResponseBodyInspectionScripts) SetModel(v string) *ListInspectionDevicesResponseBodyInspectionScripts {
	s.Model = &v
	return s
}

func (s *ListInspectionDevicesResponseBodyInspectionScripts) SetHostName(v string) *ListInspectionDevicesResponseBodyInspectionScripts {
	s.HostName = &v
	return s
}

func (s *ListInspectionDevicesResponseBodyInspectionScripts) SetManageIp(v string) *ListInspectionDevicesResponseBodyInspectionScripts {
	s.ManageIp = &v
	return s
}

func (s *ListInspectionDevicesResponseBodyInspectionScripts) SetDeviceState(v string) *ListInspectionDevicesResponseBodyInspectionScripts {
	s.DeviceState = &v
	return s
}

func (s *ListInspectionDevicesResponseBodyInspectionScripts) SetSpace(v string) *ListInspectionDevicesResponseBodyInspectionScripts {
	s.Space = &v
	return s
}

func (s *ListInspectionDevicesResponseBodyInspectionScripts) SetDeviceId(v string) *ListInspectionDevicesResponseBodyInspectionScripts {
	s.DeviceId = &v
	return s
}

type ListInspectionDevicesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInspectionDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInspectionDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInspectionDevicesResponse) GoString() string {
	return s.String()
}

func (s *ListInspectionDevicesResponse) SetHeaders(v map[string]*string) *ListInspectionDevicesResponse {
	s.Headers = v
	return s
}

func (s *ListInspectionDevicesResponse) SetBody(v *ListInspectionDevicesResponseBody) *ListInspectionDevicesResponse {
	s.Body = v
	return s
}

type ListInspectionTasksRequest struct {
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 巡检项ID
	ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// 主机名
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// 设备IP
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// 巡检状态
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 物理空间
	Space *string `json:"Space,omitempty" xml:"Space,omitempty"`
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s ListInspectionTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInspectionTasksRequest) GoString() string {
	return s.String()
}

func (s *ListInspectionTasksRequest) SetMaxResults(v int32) *ListInspectionTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInspectionTasksRequest) SetNextToken(v string) *ListInspectionTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListInspectionTasksRequest) SetItemId(v string) *ListInspectionTasksRequest {
	s.ItemId = &v
	return s
}

func (s *ListInspectionTasksRequest) SetHostName(v string) *ListInspectionTasksRequest {
	s.HostName = &v
	return s
}

func (s *ListInspectionTasksRequest) SetIP(v string) *ListInspectionTasksRequest {
	s.IP = &v
	return s
}

func (s *ListInspectionTasksRequest) SetTaskStatus(v string) *ListInspectionTasksRequest {
	s.TaskStatus = &v
	return s
}

func (s *ListInspectionTasksRequest) SetInstanceId(v string) *ListInspectionTasksRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInspectionTasksRequest) SetSpace(v string) *ListInspectionTasksRequest {
	s.Space = &v
	return s
}

func (s *ListInspectionTasksRequest) SetRole(v string) *ListInspectionTasksRequest {
	s.Role = &v
	return s
}

func (s *ListInspectionTasksRequest) SetVendor(v string) *ListInspectionTasksRequest {
	s.Vendor = &v
	return s
}

func (s *ListInspectionTasksRequest) SetModel(v string) *ListInspectionTasksRequest {
	s.Model = &v
	return s
}

type ListInspectionTasksResponseBody struct {
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 数组，返回示例目录。
	InspectionTasks []*ListInspectionTasksResponseBodyInspectionTasks `json:"InspectionTasks,omitempty" xml:"InspectionTasks,omitempty" type:"Repeated"`
}

func (s ListInspectionTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInspectionTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListInspectionTasksResponseBody) SetTotalCount(v int32) *ListInspectionTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInspectionTasksResponseBody) SetRequestId(v string) *ListInspectionTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInspectionTasksResponseBody) SetNextToken(v int32) *ListInspectionTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInspectionTasksResponseBody) SetInspectionTasks(v []*ListInspectionTasksResponseBodyInspectionTasks) *ListInspectionTasksResponseBody {
	s.InspectionTasks = v
	return s
}

type ListInspectionTasksResponseBodyInspectionTasks struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 巡检项ID
	ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// 巡检结束时间
	ExecutionEndTime *string `json:"ExecutionEndTime,omitempty" xml:"ExecutionEndTime,omitempty"`
	// 巡检开始时间
	ExecutionBeginTime *string `json:"ExecutionBeginTime,omitempty" xml:"ExecutionBeginTime,omitempty"`
	// 巡检项名字
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// 模板ID
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// 物理空间
	Space *string `json:"Space,omitempty" xml:"Space,omitempty"`
	// 巡检结果
	InspectionResult *string `json:"InspectionResult,omitempty" xml:"InspectionResult,omitempty"`
	// 告警规则
	InspectionAlarmRules []*ListInspectionTasksResponseBodyInspectionTasksInspectionAlarmRules `json:"InspectionAlarmRules,omitempty" xml:"InspectionAlarmRules,omitempty" type:"Repeated"`
	// IP地址
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// 主机名
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 任务状态
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// 型号
	Model []*string `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 巡检信息
	InspectionMessage *string `json:"InspectionMessage,omitempty" xml:"InspectionMessage,omitempty"`
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 设备回显
	DeviceDisplay *string `json:"DeviceDisplay,omitempty" xml:"DeviceDisplay,omitempty"`
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ListInspectionTasksResponseBodyInspectionTasks) String() string {
	return tea.Prettify(s)
}

func (s ListInspectionTasksResponseBodyInspectionTasks) GoString() string {
	return s.String()
}

func (s *ListInspectionTasksResponseBodyInspectionTasks) SetDeviceId(v string) *ListInspectionTasksResponseBodyInspectionTasks {
	s.DeviceId = &v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasks) SetItemId(v string) *ListInspectionTasksResponseBodyInspectionTasks {
	s.ItemId = &v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasks) SetExecutionEndTime(v string) *ListInspectionTasksResponseBodyInspectionTasks {
	s.ExecutionEndTime = &v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasks) SetExecutionBeginTime(v string) *ListInspectionTasksResponseBodyInspectionTasks {
	s.ExecutionBeginTime = &v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasks) SetItemName(v string) *ListInspectionTasksResponseBodyInspectionTasks {
	s.ItemName = &v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasks) SetScriptId(v string) *ListInspectionTasksResponseBodyInspectionTasks {
	s.ScriptId = &v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasks) SetSpace(v string) *ListInspectionTasksResponseBodyInspectionTasks {
	s.Space = &v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasks) SetInspectionResult(v string) *ListInspectionTasksResponseBodyInspectionTasks {
	s.InspectionResult = &v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasks) SetInspectionAlarmRules(v []*ListInspectionTasksResponseBodyInspectionTasksInspectionAlarmRules) *ListInspectionTasksResponseBodyInspectionTasks {
	s.InspectionAlarmRules = v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasks) SetIP(v string) *ListInspectionTasksResponseBodyInspectionTasks {
	s.IP = &v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasks) SetHostName(v string) *ListInspectionTasksResponseBodyInspectionTasks {
	s.HostName = &v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasks) SetVendor(v string) *ListInspectionTasksResponseBodyInspectionTasks {
	s.Vendor = &v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasks) SetTaskStatus(v string) *ListInspectionTasksResponseBodyInspectionTasks {
	s.TaskStatus = &v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasks) SetModel(v []*string) *ListInspectionTasksResponseBodyInspectionTasks {
	s.Model = v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasks) SetErrorCode(v string) *ListInspectionTasksResponseBodyInspectionTasks {
	s.ErrorCode = &v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasks) SetInspectionMessage(v string) *ListInspectionTasksResponseBodyInspectionTasks {
	s.InspectionMessage = &v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasks) SetTaskId(v string) *ListInspectionTasksResponseBodyInspectionTasks {
	s.TaskId = &v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasks) SetDeviceDisplay(v string) *ListInspectionTasksResponseBodyInspectionTasks {
	s.DeviceDisplay = &v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasks) SetRole(v string) *ListInspectionTasksResponseBodyInspectionTasks {
	s.Role = &v
	return s
}

type ListInspectionTasksResponseBodyInspectionTasksInspectionAlarmRules struct {
	// 告警符号
	AlarmExpression *string `json:"AlarmExpression,omitempty" xml:"AlarmExpression,omitempty"`
	// 告警变量
	AlarmOperator *string `json:"AlarmOperator,omitempty" xml:"AlarmOperator,omitempty"`
	// 告警值
	AlarmValue *string `json:"AlarmValue,omitempty" xml:"AlarmValue,omitempty"`
	// 告警实际值
	ActualValue *string `json:"ActualValue,omitempty" xml:"ActualValue,omitempty"`
	// 告警级别
	AlarmLevel *string `json:"AlarmLevel,omitempty" xml:"AlarmLevel,omitempty"`
}

func (s ListInspectionTasksResponseBodyInspectionTasksInspectionAlarmRules) String() string {
	return tea.Prettify(s)
}

func (s ListInspectionTasksResponseBodyInspectionTasksInspectionAlarmRules) GoString() string {
	return s.String()
}

func (s *ListInspectionTasksResponseBodyInspectionTasksInspectionAlarmRules) SetAlarmExpression(v string) *ListInspectionTasksResponseBodyInspectionTasksInspectionAlarmRules {
	s.AlarmExpression = &v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasksInspectionAlarmRules) SetAlarmOperator(v string) *ListInspectionTasksResponseBodyInspectionTasksInspectionAlarmRules {
	s.AlarmOperator = &v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasksInspectionAlarmRules) SetAlarmValue(v string) *ListInspectionTasksResponseBodyInspectionTasksInspectionAlarmRules {
	s.AlarmValue = &v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasksInspectionAlarmRules) SetActualValue(v string) *ListInspectionTasksResponseBodyInspectionTasksInspectionAlarmRules {
	s.ActualValue = &v
	return s
}

func (s *ListInspectionTasksResponseBodyInspectionTasksInspectionAlarmRules) SetAlarmLevel(v string) *ListInspectionTasksResponseBodyInspectionTasksInspectionAlarmRules {
	s.AlarmLevel = &v
	return s
}

type ListInspectionTasksResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInspectionTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInspectionTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInspectionTasksResponse) GoString() string {
	return s.String()
}

func (s *ListInspectionTasksResponse) SetHeaders(v map[string]*string) *ListInspectionTasksResponse {
	s.Headers = v
	return s
}

func (s *ListInspectionTasksResponse) SetBody(v *ListInspectionTasksResponseBody) *ListInspectionTasksResponse {
	s.Body = v
	return s
}

type ListScheduleWorkersRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListScheduleWorkersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleWorkersRequest) GoString() string {
	return s.String()
}

func (s *ListScheduleWorkersRequest) SetInstanceId(v string) *ListScheduleWorkersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScheduleWorkersRequest) SetMaxResults(v int32) *ListScheduleWorkersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListScheduleWorkersRequest) SetNextToken(v string) *ListScheduleWorkersRequest {
	s.NextToken = &v
	return s
}

type ListScheduleWorkersResponseBody struct {
	// 数组，返回示例目录。
	ScheduleWorker []*ListScheduleWorkersResponseBodyScheduleWorker `json:"ScheduleWorker,omitempty" xml:"ScheduleWorker,omitempty" type:"Repeated"`
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 本次读取的最大数据量
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListScheduleWorkersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleWorkersResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduleWorkersResponseBody) SetScheduleWorker(v []*ListScheduleWorkersResponseBodyScheduleWorker) *ListScheduleWorkersResponseBody {
	s.ScheduleWorker = v
	return s
}

func (s *ListScheduleWorkersResponseBody) SetTotalCount(v int32) *ListScheduleWorkersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListScheduleWorkersResponseBody) SetMaxResults(v int64) *ListScheduleWorkersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListScheduleWorkersResponseBody) SetRequestId(v string) *ListScheduleWorkersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduleWorkersResponseBody) SetNextToken(v int32) *ListScheduleWorkersResponseBody {
	s.NextToken = &v
	return s
}

type ListScheduleWorkersResponseBodyScheduleWorker struct {
	// 资源一级ID
	ScheduleWorkerId *string `json:"ScheduleWorkerId,omitempty" xml:"ScheduleWorkerId,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 值班人员工号
	WorkerId *string `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
	// 值班人员姓名
	WorkerName *string `json:"WorkerName,omitempty" xml:"WorkerName,omitempty"`
	// 联系方式
	WorkerContact *string `json:"WorkerContact,omitempty" xml:"WorkerContact,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListScheduleWorkersResponseBodyScheduleWorker) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleWorkersResponseBodyScheduleWorker) GoString() string {
	return s.String()
}

func (s *ListScheduleWorkersResponseBodyScheduleWorker) SetScheduleWorkerId(v string) *ListScheduleWorkersResponseBodyScheduleWorker {
	s.ScheduleWorkerId = &v
	return s
}

func (s *ListScheduleWorkersResponseBodyScheduleWorker) SetCreateTime(v string) *ListScheduleWorkersResponseBodyScheduleWorker {
	s.CreateTime = &v
	return s
}

func (s *ListScheduleWorkersResponseBodyScheduleWorker) SetWorkerId(v string) *ListScheduleWorkersResponseBodyScheduleWorker {
	s.WorkerId = &v
	return s
}

func (s *ListScheduleWorkersResponseBodyScheduleWorker) SetWorkerName(v string) *ListScheduleWorkersResponseBodyScheduleWorker {
	s.WorkerName = &v
	return s
}

func (s *ListScheduleWorkersResponseBodyScheduleWorker) SetWorkerContact(v string) *ListScheduleWorkersResponseBodyScheduleWorker {
	s.WorkerContact = &v
	return s
}

func (s *ListScheduleWorkersResponseBodyScheduleWorker) SetUpdateTime(v string) *ListScheduleWorkersResponseBodyScheduleWorker {
	s.UpdateTime = &v
	return s
}

type ListScheduleWorkersResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListScheduleWorkersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListScheduleWorkersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleWorkersResponse) GoString() string {
	return s.String()
}

func (s *ListScheduleWorkersResponse) SetHeaders(v map[string]*string) *ListScheduleWorkersResponse {
	s.Headers = v
	return s
}

func (s *ListScheduleWorkersResponse) SetBody(v *ListScheduleWorkersResponseBody) *ListScheduleWorkersResponse {
	s.Body = v
	return s
}

type UpdateProjectProgressRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源id
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// 建设进展
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s UpdateProjectProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectProgressRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectProgressRequest) SetInstanceId(v string) *UpdateProjectProgressRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateProjectProgressRequest) SetSetupProjectId(v string) *UpdateProjectProgressRequest {
	s.SetupProjectId = &v
	return s
}

func (s *UpdateProjectProgressRequest) SetProgress(v string) *UpdateProjectProgressRequest {
	s.Progress = &v
	return s
}

type UpdateProjectProgressResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProjectProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectProgressResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectProgressResponseBody) SetRequestId(v string) *UpdateProjectProgressResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProjectProgressResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProjectProgressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProjectProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectProgressResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectProgressResponse) SetHeaders(v map[string]*string) *UpdateProjectProgressResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectProgressResponse) SetBody(v *UpdateProjectProgressResponseBody) *UpdateProjectProgressResponse {
	s.Body = v
	return s
}

type UpdateDeviceResourceRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源一级ID
	DeviceResourceId *string `json:"DeviceResourceId,omitempty" xml:"DeviceResourceId,omitempty"`
	// 资源一级ID
	DeviceResourceIds []*string `json:"DeviceResourceIds,omitempty" xml:"DeviceResourceIds,omitempty" type:"Repeated"`
	// 建设项目id
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// 操作类型
	UpdateType *string `json:"UpdateType,omitempty" xml:"UpdateType,omitempty"`
	// 更新数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s UpdateDeviceResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeviceResourceRequest) SetInstanceId(v string) *UpdateDeviceResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateDeviceResourceRequest) SetDeviceResourceId(v string) *UpdateDeviceResourceRequest {
	s.DeviceResourceId = &v
	return s
}

func (s *UpdateDeviceResourceRequest) SetDeviceResourceIds(v []*string) *UpdateDeviceResourceRequest {
	s.DeviceResourceIds = v
	return s
}

func (s *UpdateDeviceResourceRequest) SetSetupProjectId(v string) *UpdateDeviceResourceRequest {
	s.SetupProjectId = &v
	return s
}

func (s *UpdateDeviceResourceRequest) SetUpdateType(v string) *UpdateDeviceResourceRequest {
	s.UpdateType = &v
	return s
}

func (s *UpdateDeviceResourceRequest) SetData(v string) *UpdateDeviceResourceRequest {
	s.Data = &v
	return s
}

type UpdateDeviceResourceShrinkRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源一级ID
	DeviceResourceId *string `json:"DeviceResourceId,omitempty" xml:"DeviceResourceId,omitempty"`
	// 资源一级ID
	DeviceResourceIdsShrink *string `json:"DeviceResourceIds,omitempty" xml:"DeviceResourceIds,omitempty"`
	// 建设项目id
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// 操作类型
	UpdateType *string `json:"UpdateType,omitempty" xml:"UpdateType,omitempty"`
	// 更新数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s UpdateDeviceResourceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeviceResourceShrinkRequest) SetInstanceId(v string) *UpdateDeviceResourceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateDeviceResourceShrinkRequest) SetDeviceResourceId(v string) *UpdateDeviceResourceShrinkRequest {
	s.DeviceResourceId = &v
	return s
}

func (s *UpdateDeviceResourceShrinkRequest) SetDeviceResourceIdsShrink(v string) *UpdateDeviceResourceShrinkRequest {
	s.DeviceResourceIdsShrink = &v
	return s
}

func (s *UpdateDeviceResourceShrinkRequest) SetSetupProjectId(v string) *UpdateDeviceResourceShrinkRequest {
	s.SetupProjectId = &v
	return s
}

func (s *UpdateDeviceResourceShrinkRequest) SetUpdateType(v string) *UpdateDeviceResourceShrinkRequest {
	s.UpdateType = &v
	return s
}

func (s *UpdateDeviceResourceShrinkRequest) SetData(v string) *UpdateDeviceResourceShrinkRequest {
	s.Data = &v
	return s
}

type UpdateDeviceResourceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDeviceResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeviceResourceResponseBody) SetRequestId(v string) *UpdateDeviceResourceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDeviceResourceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDeviceResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDeviceResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeviceResourceResponse) SetHeaders(v map[string]*string) *UpdateDeviceResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeviceResourceResponse) SetBody(v *UpdateDeviceResourceResponseBody) *UpdateDeviceResourceResponse {
	s.Body = v
	return s
}

type GetDevicePropertyRequest struct {
	// 实例 ID。
	DevicePropertyId *string `json:"DevicePropertyId,omitempty" xml:"DevicePropertyId,omitempty"`
	// 属性主键
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// 设备形态ID
	DeviceFormId *string `json:"DeviceFormId,omitempty" xml:"DeviceFormId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetDevicePropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDevicePropertyRequest) GoString() string {
	return s.String()
}

func (s *GetDevicePropertyRequest) SetDevicePropertyId(v string) *GetDevicePropertyRequest {
	s.DevicePropertyId = &v
	return s
}

func (s *GetDevicePropertyRequest) SetPropertyKey(v string) *GetDevicePropertyRequest {
	s.PropertyKey = &v
	return s
}

func (s *GetDevicePropertyRequest) SetDeviceFormId(v string) *GetDevicePropertyRequest {
	s.DeviceFormId = &v
	return s
}

func (s *GetDevicePropertyRequest) SetInstanceId(v string) *GetDevicePropertyRequest {
	s.InstanceId = &v
	return s
}

type GetDevicePropertyResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 设备属性详情
	DeviceProperty *GetDevicePropertyResponseBodyDeviceProperty `json:"DeviceProperty,omitempty" xml:"DeviceProperty,omitempty" type:"Struct"`
}

func (s GetDevicePropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDevicePropertyResponseBody) GoString() string {
	return s.String()
}

func (s *GetDevicePropertyResponseBody) SetRequestId(v string) *GetDevicePropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDevicePropertyResponseBody) SetDeviceProperty(v *GetDevicePropertyResponseBodyDeviceProperty) *GetDevicePropertyResponseBody {
	s.DeviceProperty = v
	return s
}

type GetDevicePropertyResponseBodyDeviceProperty struct {
	// 设备属性ID
	DevicePropertyId *string `json:"DevicePropertyId,omitempty" xml:"DevicePropertyId,omitempty"`
	// 设备形态ID
	DeviceFormId *string `json:"DeviceFormId,omitempty" xml:"DeviceFormId,omitempty"`
	// 设备形态名称
	DeviceFormName *string `json:"DeviceFormName,omitempty" xml:"DeviceFormName,omitempty"`
	// 属性名称
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	// 属性主键
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// 属性格式
	PropertyFormat *string `json:"PropertyFormat,omitempty" xml:"PropertyFormat,omitempty"`
	// 属性内容
	PropertyContent *string `json:"PropertyContent,omitempty" xml:"PropertyContent,omitempty"`
	// 是否内置属性
	BuiltIn *bool `json:"BuiltIn,omitempty" xml:"BuiltIn,omitempty"`
}

func (s GetDevicePropertyResponseBodyDeviceProperty) String() string {
	return tea.Prettify(s)
}

func (s GetDevicePropertyResponseBodyDeviceProperty) GoString() string {
	return s.String()
}

func (s *GetDevicePropertyResponseBodyDeviceProperty) SetDevicePropertyId(v string) *GetDevicePropertyResponseBodyDeviceProperty {
	s.DevicePropertyId = &v
	return s
}

func (s *GetDevicePropertyResponseBodyDeviceProperty) SetDeviceFormId(v string) *GetDevicePropertyResponseBodyDeviceProperty {
	s.DeviceFormId = &v
	return s
}

func (s *GetDevicePropertyResponseBodyDeviceProperty) SetDeviceFormName(v string) *GetDevicePropertyResponseBodyDeviceProperty {
	s.DeviceFormName = &v
	return s
}

func (s *GetDevicePropertyResponseBodyDeviceProperty) SetPropertyName(v string) *GetDevicePropertyResponseBodyDeviceProperty {
	s.PropertyName = &v
	return s
}

func (s *GetDevicePropertyResponseBodyDeviceProperty) SetPropertyKey(v string) *GetDevicePropertyResponseBodyDeviceProperty {
	s.PropertyKey = &v
	return s
}

func (s *GetDevicePropertyResponseBodyDeviceProperty) SetPropertyFormat(v string) *GetDevicePropertyResponseBodyDeviceProperty {
	s.PropertyFormat = &v
	return s
}

func (s *GetDevicePropertyResponseBodyDeviceProperty) SetPropertyContent(v string) *GetDevicePropertyResponseBodyDeviceProperty {
	s.PropertyContent = &v
	return s
}

func (s *GetDevicePropertyResponseBodyDeviceProperty) SetBuiltIn(v bool) *GetDevicePropertyResponseBodyDeviceProperty {
	s.BuiltIn = &v
	return s
}

type GetDevicePropertyResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDevicePropertyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDevicePropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDevicePropertyResponse) GoString() string {
	return s.String()
}

func (s *GetDevicePropertyResponse) SetHeaders(v map[string]*string) *GetDevicePropertyResponse {
	s.Headers = v
	return s
}

func (s *GetDevicePropertyResponse) SetBody(v *GetDevicePropertyResponseBody) *GetDevicePropertyResponse {
	s.Body = v
	return s
}

type ListResourceTypesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 资源类型
	ResourceType []*ListResourceTypesResponseBodyResourceType `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBody) SetRequestId(v string) *ListResourceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceTypesResponseBody) SetResourceType(v []*ListResourceTypesResponseBodyResourceType) *ListResourceTypesResponseBody {
	s.ResourceType = v
	return s
}

type ListResourceTypesResponseBodyResourceType struct {
	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 资源名称
	ResourceTypeName *string `json:"ResourceTypeName,omitempty" xml:"ResourceTypeName,omitempty"`
	Key              *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListResourceTypesResponseBodyResourceType) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceType) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceType) SetResourceType(v string) *ListResourceTypesResponseBodyResourceType {
	s.ResourceType = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceType) SetResourceTypeName(v string) *ListResourceTypesResponseBodyResourceType {
	s.ResourceTypeName = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceType) SetKey(v string) *ListResourceTypesResponseBodyResourceType {
	s.Key = &v
	return s
}

type ListResourceTypesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListResourceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourceTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponse) SetHeaders(v map[string]*string) *ListResourceTypesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceTypesResponse) SetBody(v *ListResourceTypesResponseBody) *ListResourceTypesResponse {
	s.Body = v
	return s
}

type GetSetupProjectRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例 ID。
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
}

func (s GetSetupProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSetupProjectRequest) GoString() string {
	return s.String()
}

func (s *GetSetupProjectRequest) SetInstanceId(v string) *GetSetupProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSetupProjectRequest) SetSetupProjectId(v string) *GetSetupProjectRequest {
	s.SetupProjectId = &v
	return s
}

type GetSetupProjectResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 建设项目
	SetupProject *GetSetupProjectResponseBodySetupProject `json:"SetupProject,omitempty" xml:"SetupProject,omitempty" type:"Struct"`
}

func (s GetSetupProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSetupProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetSetupProjectResponseBody) SetRequestId(v string) *GetSetupProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSetupProjectResponseBody) SetSetupProject(v *GetSetupProjectResponseBodySetupProject) *GetSetupProjectResponseBody {
	s.SetupProject = v
	return s
}

type GetSetupProjectResponseBodySetupProject struct {
	// 预计交付时间
	DeliveryTime *string `json:"DeliveryTime,omitempty" xml:"DeliveryTime,omitempty"`
	// 套餐
	Packages []*GetSetupProjectResponseBodySetupProjectPackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Repeated"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 物理空间uId
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// 物理空间名称
	SpaceName *string `json:"SpaceName,omitempty" xml:"SpaceName,omitempty"`
	// 物理空间类型
	SpaceType *string `json:"SpaceType,omitempty" xml:"SpaceType,omitempty"`
	// 资源一级ID
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// 节点
	Nodes *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	// 项目进展
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s GetSetupProjectResponseBodySetupProject) String() string {
	return tea.Prettify(s)
}

func (s GetSetupProjectResponseBodySetupProject) GoString() string {
	return s.String()
}

func (s *GetSetupProjectResponseBodySetupProject) SetDeliveryTime(v string) *GetSetupProjectResponseBodySetupProject {
	s.DeliveryTime = &v
	return s
}

func (s *GetSetupProjectResponseBodySetupProject) SetPackages(v []*GetSetupProjectResponseBodySetupProjectPackages) *GetSetupProjectResponseBodySetupProject {
	s.Packages = v
	return s
}

func (s *GetSetupProjectResponseBodySetupProject) SetCreateTime(v string) *GetSetupProjectResponseBodySetupProject {
	s.CreateTime = &v
	return s
}

func (s *GetSetupProjectResponseBodySetupProject) SetSpaceId(v string) *GetSetupProjectResponseBodySetupProject {
	s.SpaceId = &v
	return s
}

func (s *GetSetupProjectResponseBodySetupProject) SetSpaceName(v string) *GetSetupProjectResponseBodySetupProject {
	s.SpaceName = &v
	return s
}

func (s *GetSetupProjectResponseBodySetupProject) SetSpaceType(v string) *GetSetupProjectResponseBodySetupProject {
	s.SpaceType = &v
	return s
}

func (s *GetSetupProjectResponseBodySetupProject) SetSetupProjectId(v string) *GetSetupProjectResponseBodySetupProject {
	s.SetupProjectId = &v
	return s
}

func (s *GetSetupProjectResponseBodySetupProject) SetNodes(v string) *GetSetupProjectResponseBodySetupProject {
	s.Nodes = &v
	return s
}

func (s *GetSetupProjectResponseBodySetupProject) SetProgress(v string) *GetSetupProjectResponseBodySetupProject {
	s.Progress = &v
	return s
}

type GetSetupProjectResponseBodySetupProjectPackages struct {
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 设备号
	DeviceNumber *string `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s GetSetupProjectResponseBodySetupProjectPackages) String() string {
	return tea.Prettify(s)
}

func (s GetSetupProjectResponseBodySetupProjectPackages) GoString() string {
	return s.String()
}

func (s *GetSetupProjectResponseBodySetupProjectPackages) SetRole(v string) *GetSetupProjectResponseBodySetupProjectPackages {
	s.Role = &v
	return s
}

func (s *GetSetupProjectResponseBodySetupProjectPackages) SetDeviceNumber(v string) *GetSetupProjectResponseBodySetupProjectPackages {
	s.DeviceNumber = &v
	return s
}

func (s *GetSetupProjectResponseBodySetupProjectPackages) SetVendor(v string) *GetSetupProjectResponseBodySetupProjectPackages {
	s.Vendor = &v
	return s
}

func (s *GetSetupProjectResponseBodySetupProjectPackages) SetModel(v string) *GetSetupProjectResponseBodySetupProjectPackages {
	s.Model = &v
	return s
}

type GetSetupProjectResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSetupProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSetupProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSetupProjectResponse) GoString() string {
	return s.String()
}

func (s *GetSetupProjectResponse) SetHeaders(v map[string]*string) *GetSetupProjectResponse {
	s.Headers = v
	return s
}

func (s *GetSetupProjectResponse) SetBody(v *GetSetupProjectResponseBody) *GetSetupProjectResponse {
	s.Body = v
	return s
}

type ListConfigurationVariateRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 变量名字
	VariateName *string `json:"VariateName,omitempty" xml:"VariateName,omitempty"`
}

func (s ListConfigurationVariateRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConfigurationVariateRequest) GoString() string {
	return s.String()
}

func (s *ListConfigurationVariateRequest) SetInstanceId(v string) *ListConfigurationVariateRequest {
	s.InstanceId = &v
	return s
}

func (s *ListConfigurationVariateRequest) SetMaxResults(v int32) *ListConfigurationVariateRequest {
	s.MaxResults = &v
	return s
}

func (s *ListConfigurationVariateRequest) SetNextToken(v string) *ListConfigurationVariateRequest {
	s.NextToken = &v
	return s
}

func (s *ListConfigurationVariateRequest) SetVariateName(v string) *ListConfigurationVariateRequest {
	s.VariateName = &v
	return s
}

type ListConfigurationVariateResponseBody struct {
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 数组，返回示例目录。
	ConfigurationVariate []*ListConfigurationVariateResponseBodyConfigurationVariate `json:"ConfigurationVariate,omitempty" xml:"ConfigurationVariate,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据量
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListConfigurationVariateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConfigurationVariateResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigurationVariateResponseBody) SetTotalCount(v int32) *ListConfigurationVariateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListConfigurationVariateResponseBody) SetConfigurationVariate(v []*ListConfigurationVariateResponseBodyConfigurationVariate) *ListConfigurationVariateResponseBody {
	s.ConfigurationVariate = v
	return s
}

func (s *ListConfigurationVariateResponseBody) SetRequestId(v string) *ListConfigurationVariateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConfigurationVariateResponseBody) SetNextToken(v int32) *ListConfigurationVariateResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListConfigurationVariateResponseBody) SetMaxResults(v int64) *ListConfigurationVariateResponseBody {
	s.MaxResults = &v
	return s
}

type ListConfigurationVariateResponseBodyConfigurationVariate struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 配置变量id
	ConfigurationVariateId *string `json:"ConfigurationVariateId,omitempty" xml:"ConfigurationVariateId,omitempty"`
	// 配置变量名称
	VariateName *string `json:"VariateName,omitempty" xml:"VariateName,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 配置变量描述
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// 配置变量转换函数
	FormatFunction *string `json:"FormatFunction,omitempty" xml:"FormatFunction,omitempty"`
}

func (s ListConfigurationVariateResponseBodyConfigurationVariate) String() string {
	return tea.Prettify(s)
}

func (s ListConfigurationVariateResponseBodyConfigurationVariate) GoString() string {
	return s.String()
}

func (s *ListConfigurationVariateResponseBodyConfigurationVariate) SetCreateTime(v string) *ListConfigurationVariateResponseBodyConfigurationVariate {
	s.CreateTime = &v
	return s
}

func (s *ListConfigurationVariateResponseBodyConfigurationVariate) SetConfigurationVariateId(v string) *ListConfigurationVariateResponseBodyConfigurationVariate {
	s.ConfigurationVariateId = &v
	return s
}

func (s *ListConfigurationVariateResponseBodyConfigurationVariate) SetVariateName(v string) *ListConfigurationVariateResponseBodyConfigurationVariate {
	s.VariateName = &v
	return s
}

func (s *ListConfigurationVariateResponseBodyConfigurationVariate) SetUpdateTime(v string) *ListConfigurationVariateResponseBodyConfigurationVariate {
	s.UpdateTime = &v
	return s
}

func (s *ListConfigurationVariateResponseBodyConfigurationVariate) SetComment(v string) *ListConfigurationVariateResponseBodyConfigurationVariate {
	s.Comment = &v
	return s
}

func (s *ListConfigurationVariateResponseBodyConfigurationVariate) SetFormatFunction(v string) *ListConfigurationVariateResponseBodyConfigurationVariate {
	s.FormatFunction = &v
	return s
}

type ListConfigurationVariateResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConfigurationVariateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConfigurationVariateResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConfigurationVariateResponse) GoString() string {
	return s.String()
}

func (s *ListConfigurationVariateResponse) SetHeaders(v map[string]*string) *ListConfigurationVariateResponse {
	s.Headers = v
	return s
}

func (s *ListConfigurationVariateResponse) SetBody(v *ListConfigurationVariateResponseBody) *ListConfigurationVariateResponse {
	s.Body = v
	return s
}

type CreateSpaceModelRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 物理空间类型
	SpaceType *string `json:"SpaceType,omitempty" xml:"SpaceType,omitempty"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateSpaceModelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSpaceModelRequest) GoString() string {
	return s.String()
}

func (s *CreateSpaceModelRequest) SetInstanceId(v string) *CreateSpaceModelRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSpaceModelRequest) SetSpaceType(v string) *CreateSpaceModelRequest {
	s.SpaceType = &v
	return s
}

func (s *CreateSpaceModelRequest) SetClientToken(v string) *CreateSpaceModelRequest {
	s.ClientToken = &v
	return s
}

type CreateSpaceModelResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 资源实例ID，如ECS实例的创建接口CreateInstance应返回InstanceId。
	SpaceModelId *string `json:"SpaceModelId,omitempty" xml:"SpaceModelId,omitempty"`
}

func (s CreateSpaceModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSpaceModelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSpaceModelResponseBody) SetRequestId(v string) *CreateSpaceModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSpaceModelResponseBody) SetSpaceModelId(v string) *CreateSpaceModelResponseBody {
	s.SpaceModelId = &v
	return s
}

type CreateSpaceModelResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSpaceModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSpaceModelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSpaceModelResponse) GoString() string {
	return s.String()
}

func (s *CreateSpaceModelResponse) SetHeaders(v map[string]*string) *CreateSpaceModelResponse {
	s.Headers = v
	return s
}

func (s *CreateSpaceModelResponse) SetBody(v *CreateSpaceModelResponseBody) *CreateSpaceModelResponse {
	s.Body = v
	return s
}

type DeleteScheduleWorkerRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例 ID。
	ScheduleWorkerId *string `json:"ScheduleWorkerId,omitempty" xml:"ScheduleWorkerId,omitempty"`
}

func (s DeleteScheduleWorkerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleWorkerRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduleWorkerRequest) SetInstanceId(v string) *DeleteScheduleWorkerRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteScheduleWorkerRequest) SetScheduleWorkerId(v string) *DeleteScheduleWorkerRequest {
	s.ScheduleWorkerId = &v
	return s
}

type DeleteScheduleWorkerResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScheduleWorkerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleWorkerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduleWorkerResponseBody) SetRequestId(v string) *DeleteScheduleWorkerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteScheduleWorkerResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteScheduleWorkerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteScheduleWorkerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleWorkerResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduleWorkerResponse) SetHeaders(v map[string]*string) *DeleteScheduleWorkerResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduleWorkerResponse) SetBody(v *DeleteScheduleWorkerResponseBody) *DeleteScheduleWorkerResponse {
	s.Body = v
	return s
}

type ListDedicatedLinesRequest struct {
	// 物理空间ID
	PhysicalSpaceId *string `json:"PhysicalSpaceId,omitempty" xml:"PhysicalSpaceId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListDedicatedLinesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDedicatedLinesRequest) GoString() string {
	return s.String()
}

func (s *ListDedicatedLinesRequest) SetPhysicalSpaceId(v string) *ListDedicatedLinesRequest {
	s.PhysicalSpaceId = &v
	return s
}

func (s *ListDedicatedLinesRequest) SetInstanceId(v string) *ListDedicatedLinesRequest {
	s.InstanceId = &v
	return s
}

type ListDedicatedLinesResponseBody struct {
	// 数组，返回示例目录。
	DedicatedLines []*ListDedicatedLinesResponseBodyDedicatedLines `json:"DedicatedLines,omitempty" xml:"DedicatedLines,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDedicatedLinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDedicatedLinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDedicatedLinesResponseBody) SetDedicatedLines(v []*ListDedicatedLinesResponseBodyDedicatedLines) *ListDedicatedLinesResponseBody {
	s.DedicatedLines = v
	return s
}

func (s *ListDedicatedLinesResponseBody) SetRequestId(v string) *ListDedicatedLinesResponseBody {
	s.RequestId = &v
	return s
}

type ListDedicatedLinesResponseBodyDedicatedLines struct {
	// 物理空间专线ID
	DedicatedLineId *string `json:"DedicatedLineId,omitempty" xml:"DedicatedLineId,omitempty"`
	// 运营商
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// 宽带（Mbps）
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// 专线IP
	DedicatedLineIp *string `json:"DedicatedLineIp,omitempty" xml:"DedicatedLineIp,omitempty"`
	// 专线网关
	DedicatedLineGateway *string `json:"DedicatedLineGateway,omitempty" xml:"DedicatedLineGateway,omitempty"`
	// 专线角色
	DedicatedLineRole *string `json:"DedicatedLineRole,omitempty" xml:"DedicatedLineRole,omitempty"`
	// 关联设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 关联设备端口
	DevicePort *string `json:"DevicePort,omitempty" xml:"DevicePort,omitempty"`
	// 关联设备名称
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// 关联物理空间ID
	PhysicalSpaceId *string `json:"PhysicalSpaceId,omitempty" xml:"PhysicalSpaceId,omitempty"`
}

func (s ListDedicatedLinesResponseBodyDedicatedLines) String() string {
	return tea.Prettify(s)
}

func (s ListDedicatedLinesResponseBodyDedicatedLines) GoString() string {
	return s.String()
}

func (s *ListDedicatedLinesResponseBodyDedicatedLines) SetDedicatedLineId(v string) *ListDedicatedLinesResponseBodyDedicatedLines {
	s.DedicatedLineId = &v
	return s
}

func (s *ListDedicatedLinesResponseBodyDedicatedLines) SetIsp(v string) *ListDedicatedLinesResponseBodyDedicatedLines {
	s.Isp = &v
	return s
}

func (s *ListDedicatedLinesResponseBodyDedicatedLines) SetBandwidth(v int32) *ListDedicatedLinesResponseBodyDedicatedLines {
	s.Bandwidth = &v
	return s
}

func (s *ListDedicatedLinesResponseBodyDedicatedLines) SetDedicatedLineIp(v string) *ListDedicatedLinesResponseBodyDedicatedLines {
	s.DedicatedLineIp = &v
	return s
}

func (s *ListDedicatedLinesResponseBodyDedicatedLines) SetDedicatedLineGateway(v string) *ListDedicatedLinesResponseBodyDedicatedLines {
	s.DedicatedLineGateway = &v
	return s
}

func (s *ListDedicatedLinesResponseBodyDedicatedLines) SetDedicatedLineRole(v string) *ListDedicatedLinesResponseBodyDedicatedLines {
	s.DedicatedLineRole = &v
	return s
}

func (s *ListDedicatedLinesResponseBodyDedicatedLines) SetDeviceId(v string) *ListDedicatedLinesResponseBodyDedicatedLines {
	s.DeviceId = &v
	return s
}

func (s *ListDedicatedLinesResponseBodyDedicatedLines) SetDevicePort(v string) *ListDedicatedLinesResponseBodyDedicatedLines {
	s.DevicePort = &v
	return s
}

func (s *ListDedicatedLinesResponseBodyDedicatedLines) SetDeviceName(v string) *ListDedicatedLinesResponseBodyDedicatedLines {
	s.DeviceName = &v
	return s
}

func (s *ListDedicatedLinesResponseBodyDedicatedLines) SetPhysicalSpaceId(v string) *ListDedicatedLinesResponseBodyDedicatedLines {
	s.PhysicalSpaceId = &v
	return s
}

type ListDedicatedLinesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDedicatedLinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDedicatedLinesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDedicatedLinesResponse) GoString() string {
	return s.String()
}

func (s *ListDedicatedLinesResponse) SetHeaders(v map[string]*string) *ListDedicatedLinesResponse {
	s.Headers = v
	return s
}

func (s *ListDedicatedLinesResponse) SetBody(v *ListDedicatedLinesResponseBody) *ListDedicatedLinesResponse {
	s.Body = v
	return s
}

type UpdateInformationKeyActionRequest struct {
	// 资源id
	ResourceInformationId *string `json:"ResourceInformationId,omitempty" xml:"ResourceInformationId,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 键值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// 动作
	KeyAction *string `json:"KeyAction,omitempty" xml:"KeyAction,omitempty"`
	// 建设项目id
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
}

func (s UpdateInformationKeyActionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInformationKeyActionRequest) GoString() string {
	return s.String()
}

func (s *UpdateInformationKeyActionRequest) SetResourceInformationId(v string) *UpdateInformationKeyActionRequest {
	s.ResourceInformationId = &v
	return s
}

func (s *UpdateInformationKeyActionRequest) SetInstanceId(v string) *UpdateInformationKeyActionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInformationKeyActionRequest) SetKey(v string) *UpdateInformationKeyActionRequest {
	s.Key = &v
	return s
}

func (s *UpdateInformationKeyActionRequest) SetValue(v string) *UpdateInformationKeyActionRequest {
	s.Value = &v
	return s
}

func (s *UpdateInformationKeyActionRequest) SetKeyAction(v string) *UpdateInformationKeyActionRequest {
	s.KeyAction = &v
	return s
}

func (s *UpdateInformationKeyActionRequest) SetSetupProjectId(v string) *UpdateInformationKeyActionRequest {
	s.SetupProjectId = &v
	return s
}

type UpdateInformationKeyActionResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回信息
	ActionMessage *string `json:"ActionMessage,omitempty" xml:"ActionMessage,omitempty"`
}

func (s UpdateInformationKeyActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInformationKeyActionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInformationKeyActionResponseBody) SetRequestId(v string) *UpdateInformationKeyActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInformationKeyActionResponseBody) SetActionMessage(v string) *UpdateInformationKeyActionResponseBody {
	s.ActionMessage = &v
	return s
}

type UpdateInformationKeyActionResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInformationKeyActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInformationKeyActionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInformationKeyActionResponse) GoString() string {
	return s.String()
}

func (s *UpdateInformationKeyActionResponse) SetHeaders(v map[string]*string) *UpdateInformationKeyActionResponse {
	s.Headers = v
	return s
}

func (s *UpdateInformationKeyActionResponse) SetBody(v *UpdateInformationKeyActionResponseBody) *UpdateInformationKeyActionResponse {
	s.Body = v
	return s
}

type GetRealtimeTaskRequest struct {
	// 实时任务ID
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetRealtimeTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRealtimeTaskRequest) GoString() string {
	return s.String()
}

func (s *GetRealtimeTaskRequest) SetTaskId(v string) *GetRealtimeTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetRealtimeTaskRequest) SetInstanceId(v string) *GetRealtimeTaskRequest {
	s.InstanceId = &v
	return s
}

type GetRealtimeTaskResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求任务结果
	InspectionTask *GetRealtimeTaskResponseBodyInspectionTask `json:"InspectionTask,omitempty" xml:"InspectionTask,omitempty" type:"Struct"`
}

func (s GetRealtimeTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRealtimeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetRealtimeTaskResponseBody) SetRequestId(v string) *GetRealtimeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRealtimeTaskResponseBody) SetInspectionTask(v *GetRealtimeTaskResponseBodyInspectionTask) *GetRealtimeTaskResponseBody {
	s.InspectionTask = v
	return s
}

type GetRealtimeTaskResponseBodyInspectionTask struct {
	// 巡检状态
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// 巡检输出
	InspectionResult *string `json:"InspectionResult,omitempty" xml:"InspectionResult,omitempty"`
	// 巡检错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 巡检错误信息
	InspectionMessage *string `json:"InspectionMessage,omitempty" xml:"InspectionMessage,omitempty"`
}

func (s GetRealtimeTaskResponseBodyInspectionTask) String() string {
	return tea.Prettify(s)
}

func (s GetRealtimeTaskResponseBodyInspectionTask) GoString() string {
	return s.String()
}

func (s *GetRealtimeTaskResponseBodyInspectionTask) SetTaskStatus(v string) *GetRealtimeTaskResponseBodyInspectionTask {
	s.TaskStatus = &v
	return s
}

func (s *GetRealtimeTaskResponseBodyInspectionTask) SetInspectionResult(v string) *GetRealtimeTaskResponseBodyInspectionTask {
	s.InspectionResult = &v
	return s
}

func (s *GetRealtimeTaskResponseBodyInspectionTask) SetErrorCode(v string) *GetRealtimeTaskResponseBodyInspectionTask {
	s.ErrorCode = &v
	return s
}

func (s *GetRealtimeTaskResponseBodyInspectionTask) SetInspectionMessage(v string) *GetRealtimeTaskResponseBodyInspectionTask {
	s.InspectionMessage = &v
	return s
}

type GetRealtimeTaskResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRealtimeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRealtimeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRealtimeTaskResponse) GoString() string {
	return s.String()
}

func (s *GetRealtimeTaskResponse) SetHeaders(v map[string]*string) *GetRealtimeTaskResponse {
	s.Headers = v
	return s
}

func (s *GetRealtimeTaskResponse) SetBody(v *GetRealtimeTaskResponseBody) *GetRealtimeTaskResponse {
	s.Body = v
	return s
}

type ListScheduleTypesRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListScheduleTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleTypesRequest) GoString() string {
	return s.String()
}

func (s *ListScheduleTypesRequest) SetInstanceId(v string) *ListScheduleTypesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScheduleTypesRequest) SetMaxResults(v int32) *ListScheduleTypesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListScheduleTypesRequest) SetNextToken(v string) *ListScheduleTypesRequest {
	s.NextToken = &v
	return s
}

type ListScheduleTypesResponseBody struct {
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 数组，返回示例目录。
	ScheduleType []*ListScheduleTypesResponseBodyScheduleType `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据量
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListScheduleTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduleTypesResponseBody) SetTotalCount(v int32) *ListScheduleTypesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListScheduleTypesResponseBody) SetScheduleType(v []*ListScheduleTypesResponseBodyScheduleType) *ListScheduleTypesResponseBody {
	s.ScheduleType = v
	return s
}

func (s *ListScheduleTypesResponseBody) SetRequestId(v string) *ListScheduleTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduleTypesResponseBody) SetNextToken(v int32) *ListScheduleTypesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListScheduleTypesResponseBody) SetMaxResults(v int64) *ListScheduleTypesResponseBody {
	s.MaxResults = &v
	return s
}

type ListScheduleTypesResponseBodyScheduleType struct {
	// 值班类型状态
	Status        *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	RelatedWorker []*string `json:"RelatedWorker,omitempty" xml:"RelatedWorker,omitempty" type:"Repeated"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 资源一级ID
	ScheduleTypeId *string `json:"ScheduleTypeId,omitempty" xml:"ScheduleTypeId,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 值班类型value
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
}

func (s ListScheduleTypesResponseBodyScheduleType) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleTypesResponseBodyScheduleType) GoString() string {
	return s.String()
}

func (s *ListScheduleTypesResponseBodyScheduleType) SetStatus(v string) *ListScheduleTypesResponseBodyScheduleType {
	s.Status = &v
	return s
}

func (s *ListScheduleTypesResponseBodyScheduleType) SetRelatedWorker(v []*string) *ListScheduleTypesResponseBodyScheduleType {
	s.RelatedWorker = v
	return s
}

func (s *ListScheduleTypesResponseBodyScheduleType) SetCreateTime(v string) *ListScheduleTypesResponseBodyScheduleType {
	s.CreateTime = &v
	return s
}

func (s *ListScheduleTypesResponseBodyScheduleType) SetScheduleTypeId(v string) *ListScheduleTypesResponseBodyScheduleType {
	s.ScheduleTypeId = &v
	return s
}

func (s *ListScheduleTypesResponseBodyScheduleType) SetUpdateTime(v string) *ListScheduleTypesResponseBodyScheduleType {
	s.UpdateTime = &v
	return s
}

func (s *ListScheduleTypesResponseBodyScheduleType) SetScheduleType(v string) *ListScheduleTypesResponseBodyScheduleType {
	s.ScheduleType = &v
	return s
}

type ListScheduleTypesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListScheduleTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListScheduleTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleTypesResponse) GoString() string {
	return s.String()
}

func (s *ListScheduleTypesResponse) SetHeaders(v map[string]*string) *ListScheduleTypesResponse {
	s.Headers = v
	return s
}

func (s *ListScheduleTypesResponse) SetBody(v *ListScheduleTypesResponseBody) *ListScheduleTypesResponse {
	s.Body = v
	return s
}

type CreateScheduleTypeRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 值班类型key
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// relatedWorkerStr
	RelatedWorker []*string `json:"RelatedWorker,omitempty" xml:"RelatedWorker,omitempty" type:"Repeated"`
	// 值班类型状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateScheduleTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleTypeRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduleTypeRequest) SetInstanceId(v string) *CreateScheduleTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateScheduleTypeRequest) SetScheduleType(v string) *CreateScheduleTypeRequest {
	s.ScheduleType = &v
	return s
}

func (s *CreateScheduleTypeRequest) SetRelatedWorker(v []*string) *CreateScheduleTypeRequest {
	s.RelatedWorker = v
	return s
}

func (s *CreateScheduleTypeRequest) SetStatus(v string) *CreateScheduleTypeRequest {
	s.Status = &v
	return s
}

func (s *CreateScheduleTypeRequest) SetClientToken(v string) *CreateScheduleTypeRequest {
	s.ClientToken = &v
	return s
}

type CreateScheduleTypeShrinkRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 值班类型key
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// relatedWorkerStr
	RelatedWorkerShrink *string `json:"RelatedWorker,omitempty" xml:"RelatedWorker,omitempty"`
	// 值班类型状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateScheduleTypeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleTypeShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduleTypeShrinkRequest) SetInstanceId(v string) *CreateScheduleTypeShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateScheduleTypeShrinkRequest) SetScheduleType(v string) *CreateScheduleTypeShrinkRequest {
	s.ScheduleType = &v
	return s
}

func (s *CreateScheduleTypeShrinkRequest) SetRelatedWorkerShrink(v string) *CreateScheduleTypeShrinkRequest {
	s.RelatedWorkerShrink = &v
	return s
}

func (s *CreateScheduleTypeShrinkRequest) SetStatus(v string) *CreateScheduleTypeShrinkRequest {
	s.Status = &v
	return s
}

func (s *CreateScheduleTypeShrinkRequest) SetClientToken(v string) *CreateScheduleTypeShrinkRequest {
	s.ClientToken = &v
	return s
}

type CreateScheduleTypeResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 资源实例ID，如ECS实例的创建接口CreateInstance应返回InstanceId。
	ScheduleTypeId *string `json:"ScheduleTypeId,omitempty" xml:"ScheduleTypeId,omitempty"`
}

func (s CreateScheduleTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleTypeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduleTypeResponseBody) SetRequestId(v string) *CreateScheduleTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduleTypeResponseBody) SetScheduleTypeId(v string) *CreateScheduleTypeResponseBody {
	s.ScheduleTypeId = &v
	return s
}

type CreateScheduleTypeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateScheduleTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateScheduleTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleTypeResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduleTypeResponse) SetHeaders(v map[string]*string) *CreateScheduleTypeResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduleTypeResponse) SetBody(v *CreateScheduleTypeResponseBody) *CreateScheduleTypeResponse {
	s.Body = v
	return s
}

type GetScheduleWorkerRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例 ID。
	ScheduleWorkerId *string `json:"ScheduleWorkerId,omitempty" xml:"ScheduleWorkerId,omitempty"`
}

func (s GetScheduleWorkerRequest) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleWorkerRequest) GoString() string {
	return s.String()
}

func (s *GetScheduleWorkerRequest) SetInstanceId(v string) *GetScheduleWorkerRequest {
	s.InstanceId = &v
	return s
}

func (s *GetScheduleWorkerRequest) SetScheduleWorkerId(v string) *GetScheduleWorkerRequest {
	s.ScheduleWorkerId = &v
	return s
}

type GetScheduleWorkerResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 值班人员对象
	ScheduleWorker *GetScheduleWorkerResponseBodyScheduleWorker `json:"ScheduleWorker,omitempty" xml:"ScheduleWorker,omitempty" type:"Struct"`
}

func (s GetScheduleWorkerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleWorkerResponseBody) GoString() string {
	return s.String()
}

func (s *GetScheduleWorkerResponseBody) SetRequestId(v string) *GetScheduleWorkerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScheduleWorkerResponseBody) SetScheduleWorker(v *GetScheduleWorkerResponseBodyScheduleWorker) *GetScheduleWorkerResponseBody {
	s.ScheduleWorker = v
	return s
}

type GetScheduleWorkerResponseBodyScheduleWorker struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 资源一级ID
	ScheduleWorkerId *string `json:"ScheduleWorkerId,omitempty" xml:"ScheduleWorkerId,omitempty"`
	// 值班人员工号
	WorkerId *string `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
	// 值班人员姓名
	WorkerName *string `json:"WorkerName,omitempty" xml:"WorkerName,omitempty"`
	// 联系方式
	WorkerContact *string `json:"WorkerContact,omitempty" xml:"WorkerContact,omitempty"`
}

func (s GetScheduleWorkerResponseBodyScheduleWorker) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleWorkerResponseBodyScheduleWorker) GoString() string {
	return s.String()
}

func (s *GetScheduleWorkerResponseBodyScheduleWorker) SetCreateTime(v string) *GetScheduleWorkerResponseBodyScheduleWorker {
	s.CreateTime = &v
	return s
}

func (s *GetScheduleWorkerResponseBodyScheduleWorker) SetUpdateTime(v string) *GetScheduleWorkerResponseBodyScheduleWorker {
	s.UpdateTime = &v
	return s
}

func (s *GetScheduleWorkerResponseBodyScheduleWorker) SetScheduleWorkerId(v string) *GetScheduleWorkerResponseBodyScheduleWorker {
	s.ScheduleWorkerId = &v
	return s
}

func (s *GetScheduleWorkerResponseBodyScheduleWorker) SetWorkerId(v string) *GetScheduleWorkerResponseBodyScheduleWorker {
	s.WorkerId = &v
	return s
}

func (s *GetScheduleWorkerResponseBodyScheduleWorker) SetWorkerName(v string) *GetScheduleWorkerResponseBodyScheduleWorker {
	s.WorkerName = &v
	return s
}

func (s *GetScheduleWorkerResponseBodyScheduleWorker) SetWorkerContact(v string) *GetScheduleWorkerResponseBodyScheduleWorker {
	s.WorkerContact = &v
	return s
}

type GetScheduleWorkerResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetScheduleWorkerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetScheduleWorkerResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleWorkerResponse) GoString() string {
	return s.String()
}

func (s *GetScheduleWorkerResponse) SetHeaders(v map[string]*string) *GetScheduleWorkerResponse {
	s.Headers = v
	return s
}

func (s *GetScheduleWorkerResponse) SetBody(v *GetScheduleWorkerResponseBody) *GetScheduleWorkerResponse {
	s.Body = v
	return s
}

type CreateScheduleWorkerRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 值班人员工号
	WorkerId *string `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
	// 值班人员姓名
	WorkerName *string `json:"WorkerName,omitempty" xml:"WorkerName,omitempty"`
	// 联系方式
	WorkerContact *string `json:"WorkerContact,omitempty" xml:"WorkerContact,omitempty"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateScheduleWorkerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleWorkerRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduleWorkerRequest) SetInstanceId(v string) *CreateScheduleWorkerRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateScheduleWorkerRequest) SetWorkerId(v string) *CreateScheduleWorkerRequest {
	s.WorkerId = &v
	return s
}

func (s *CreateScheduleWorkerRequest) SetWorkerName(v string) *CreateScheduleWorkerRequest {
	s.WorkerName = &v
	return s
}

func (s *CreateScheduleWorkerRequest) SetWorkerContact(v string) *CreateScheduleWorkerRequest {
	s.WorkerContact = &v
	return s
}

func (s *CreateScheduleWorkerRequest) SetClientToken(v string) *CreateScheduleWorkerRequest {
	s.ClientToken = &v
	return s
}

type CreateScheduleWorkerResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 资源实例ID，如ECS实例的创建接口CreateInstance应返回InstanceId。
	ScheduleWorkerId *string `json:"ScheduleWorkerId,omitempty" xml:"ScheduleWorkerId,omitempty"`
}

func (s CreateScheduleWorkerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleWorkerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduleWorkerResponseBody) SetRequestId(v string) *CreateScheduleWorkerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduleWorkerResponseBody) SetScheduleWorkerId(v string) *CreateScheduleWorkerResponseBody {
	s.ScheduleWorkerId = &v
	return s
}

type CreateScheduleWorkerResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateScheduleWorkerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateScheduleWorkerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleWorkerResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduleWorkerResponse) SetHeaders(v map[string]*string) *CreateScheduleWorkerResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduleWorkerResponse) SetBody(v *CreateScheduleWorkerResponseBody) *CreateScheduleWorkerResponse {
	s.Body = v
	return s
}

type CreateConfigurationVariateRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源名称
	VariateName *string `json:"VariateName,omitempty" xml:"VariateName,omitempty"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 描述变量
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// python转换函数
	FormatFunction *string `json:"FormatFunction,omitempty" xml:"FormatFunction,omitempty"`
}

func (s CreateConfigurationVariateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigurationVariateRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigurationVariateRequest) SetInstanceId(v string) *CreateConfigurationVariateRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateConfigurationVariateRequest) SetVariateName(v string) *CreateConfigurationVariateRequest {
	s.VariateName = &v
	return s
}

func (s *CreateConfigurationVariateRequest) SetClientToken(v string) *CreateConfigurationVariateRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateConfigurationVariateRequest) SetComment(v string) *CreateConfigurationVariateRequest {
	s.Comment = &v
	return s
}

func (s *CreateConfigurationVariateRequest) SetFormatFunction(v string) *CreateConfigurationVariateRequest {
	s.FormatFunction = &v
	return s
}

type CreateConfigurationVariateResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 资源实例ID，如ECS实例的创建接口CreateInstance应返回InstanceId。
	ConfigurationVariateId *string `json:"ConfigurationVariateId,omitempty" xml:"ConfigurationVariateId,omitempty"`
}

func (s CreateConfigurationVariateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigurationVariateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConfigurationVariateResponseBody) SetRequestId(v string) *CreateConfigurationVariateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConfigurationVariateResponseBody) SetConfigurationVariateId(v string) *CreateConfigurationVariateResponseBody {
	s.ConfigurationVariateId = &v
	return s
}

type CreateConfigurationVariateResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateConfigurationVariateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateConfigurationVariateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigurationVariateResponse) GoString() string {
	return s.String()
}

func (s *CreateConfigurationVariateResponse) SetHeaders(v map[string]*string) *CreateConfigurationVariateResponse {
	s.Headers = v
	return s
}

func (s *CreateConfigurationVariateResponse) SetBody(v *CreateConfigurationVariateResponseBody) *CreateConfigurationVariateResponse {
	s.Body = v
	return s
}

type GetSpaceModelSortRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 物理空间类型
	SpaceType *string `json:"SpaceType,omitempty" xml:"SpaceType,omitempty"`
	// 操作类型
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
}

func (s GetSpaceModelSortRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceModelSortRequest) GoString() string {
	return s.String()
}

func (s *GetSpaceModelSortRequest) SetInstanceId(v string) *GetSpaceModelSortRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSpaceModelSortRequest) SetSpaceType(v string) *GetSpaceModelSortRequest {
	s.SpaceType = &v
	return s
}

func (s *GetSpaceModelSortRequest) SetOperateType(v string) *GetSpaceModelSortRequest {
	s.OperateType = &v
	return s
}

type GetSpaceModelSortResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 建设项目
	SpaceModel []*GetSpaceModelSortResponseBodySpaceModel `json:"SpaceModel,omitempty" xml:"SpaceModel,omitempty" type:"Repeated"`
}

func (s GetSpaceModelSortResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceModelSortResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpaceModelSortResponseBody) SetRequestId(v string) *GetSpaceModelSortResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSpaceModelSortResponseBody) SetSpaceModel(v []*GetSpaceModelSortResponseBodySpaceModel) *GetSpaceModelSortResponseBody {
	s.SpaceModel = v
	return s
}

type GetSpaceModelSortResponseBodySpaceModel struct {
	// 层级名称
	LevelName *string `json:"LevelName,omitempty" xml:"LevelName,omitempty"`
	// 层级
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s GetSpaceModelSortResponseBodySpaceModel) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceModelSortResponseBodySpaceModel) GoString() string {
	return s.String()
}

func (s *GetSpaceModelSortResponseBodySpaceModel) SetLevelName(v string) *GetSpaceModelSortResponseBodySpaceModel {
	s.LevelName = &v
	return s
}

func (s *GetSpaceModelSortResponseBodySpaceModel) SetLevel(v int64) *GetSpaceModelSortResponseBodySpaceModel {
	s.Level = &v
	return s
}

type GetSpaceModelSortResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSpaceModelSortResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSpaceModelSortResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceModelSortResponse) GoString() string {
	return s.String()
}

func (s *GetSpaceModelSortResponse) SetHeaders(v map[string]*string) *GetSpaceModelSortResponse {
	s.Headers = v
	return s
}

func (s *GetSpaceModelSortResponse) SetBody(v *GetSpaceModelSortResponseBody) *GetSpaceModelSortResponse {
	s.Body = v
	return s
}

type CreateRealtimeTaskRequest struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 模板执行脚本
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateRealtimeTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRealtimeTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRealtimeTaskRequest) SetDeviceId(v string) *CreateRealtimeTaskRequest {
	s.DeviceId = &v
	return s
}

func (s *CreateRealtimeTaskRequest) SetScript(v string) *CreateRealtimeTaskRequest {
	s.Script = &v
	return s
}

func (s *CreateRealtimeTaskRequest) SetInstanceId(v string) *CreateRealtimeTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRealtimeTaskRequest) SetClientToken(v string) *CreateRealtimeTaskRequest {
	s.ClientToken = &v
	return s
}

type CreateRealtimeTaskResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实时任务ID
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateRealtimeTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRealtimeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRealtimeTaskResponseBody) SetRequestId(v string) *CreateRealtimeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRealtimeTaskResponseBody) SetTaskId(v string) *CreateRealtimeTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateRealtimeTaskResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRealtimeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRealtimeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRealtimeTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateRealtimeTaskResponse) SetHeaders(v map[string]*string) *CreateRealtimeTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateRealtimeTaskResponse) SetBody(v *CreateRealtimeTaskResponseBody) *CreateRealtimeTaskResponse {
	s.Body = v
	return s
}

type GetDeviceFormRequest struct {
	// 实例 ID。
	DeviceFormId *string `json:"DeviceFormId,omitempty" xml:"DeviceFormId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetDeviceFormRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceFormRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceFormRequest) SetDeviceFormId(v string) *GetDeviceFormRequest {
	s.DeviceFormId = &v
	return s
}

func (s *GetDeviceFormRequest) SetInstanceId(v string) *GetDeviceFormRequest {
	s.InstanceId = &v
	return s
}

type GetDeviceFormResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 设备详情
	DeviceForm *GetDeviceFormResponseBodyDeviceForm `json:"DeviceForm,omitempty" xml:"DeviceForm,omitempty" type:"Struct"`
}

func (s GetDeviceFormResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceFormResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceFormResponseBody) SetRequestId(v string) *GetDeviceFormResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceFormResponseBody) SetDeviceForm(v *GetDeviceFormResponseBodyDeviceForm) *GetDeviceFormResponseBody {
	s.DeviceForm = v
	return s
}

type GetDeviceFormResponseBodyDeviceForm struct {
	// 是否支持配置生成
	ConfigCompare *bool `json:"ConfigCompare,omitempty" xml:"ConfigCompare,omitempty"`
	// 设备形态属性列表
	AttributeList []*GetDeviceFormResponseBodyDeviceFormAttributeList `json:"AttributeList,omitempty" xml:"AttributeList,omitempty" type:"Repeated"`
	// 设备形态ID
	DeviceFormId *string `json:"DeviceFormId,omitempty" xml:"DeviceFormId,omitempty"`
	// 设备形态名称
	DeviceFormName *string `json:"DeviceFormName,omitempty" xml:"DeviceFormName,omitempty"`
	// 设备形态是否内置
	FormBuiltIn *bool `json:"FormBuiltIn,omitempty" xml:"FormBuiltIn,omitempty"`
	// 是否需要账号配置
	AccountConfig *bool `json:"AccountConfig,omitempty" xml:"AccountConfig,omitempty"`
	// 是否展示设备详情
	DetailDisplay *bool `json:"DetailDisplay,omitempty" xml:"DetailDisplay,omitempty"`
	// 设备形态主键
	UniqueKey *string `json:"UniqueKey,omitempty" xml:"UniqueKey,omitempty"`
}

func (s GetDeviceFormResponseBodyDeviceForm) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceFormResponseBodyDeviceForm) GoString() string {
	return s.String()
}

func (s *GetDeviceFormResponseBodyDeviceForm) SetConfigCompare(v bool) *GetDeviceFormResponseBodyDeviceForm {
	s.ConfigCompare = &v
	return s
}

func (s *GetDeviceFormResponseBodyDeviceForm) SetAttributeList(v []*GetDeviceFormResponseBodyDeviceFormAttributeList) *GetDeviceFormResponseBodyDeviceForm {
	s.AttributeList = v
	return s
}

func (s *GetDeviceFormResponseBodyDeviceForm) SetDeviceFormId(v string) *GetDeviceFormResponseBodyDeviceForm {
	s.DeviceFormId = &v
	return s
}

func (s *GetDeviceFormResponseBodyDeviceForm) SetDeviceFormName(v string) *GetDeviceFormResponseBodyDeviceForm {
	s.DeviceFormName = &v
	return s
}

func (s *GetDeviceFormResponseBodyDeviceForm) SetFormBuiltIn(v bool) *GetDeviceFormResponseBodyDeviceForm {
	s.FormBuiltIn = &v
	return s
}

func (s *GetDeviceFormResponseBodyDeviceForm) SetAccountConfig(v bool) *GetDeviceFormResponseBodyDeviceForm {
	s.AccountConfig = &v
	return s
}

func (s *GetDeviceFormResponseBodyDeviceForm) SetDetailDisplay(v bool) *GetDeviceFormResponseBodyDeviceForm {
	s.DetailDisplay = &v
	return s
}

func (s *GetDeviceFormResponseBodyDeviceForm) SetUniqueKey(v string) *GetDeviceFormResponseBodyDeviceForm {
	s.UniqueKey = &v
	return s
}

type GetDeviceFormResponseBodyDeviceFormAttributeList struct {
	// 设备形态属性主键
	AttributeKey *string `json:"AttributeKey,omitempty" xml:"AttributeKey,omitempty"`
	// 设备形态属性名称
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	// 设备形态属性是否必填
	AttributeRequirement *bool `json:"AttributeRequirement,omitempty" xml:"AttributeRequirement,omitempty"`
	// 设备形态属性是否唯一
	AttributeUniqueness *bool `json:"AttributeUniqueness,omitempty" xml:"AttributeUniqueness,omitempty"`
	// 设备形态属性值格式
	AttributeFormat *string `json:"AttributeFormat,omitempty" xml:"AttributeFormat,omitempty"`
	// 设备形态属性值类型
	AttributeType *string `json:"AttributeType,omitempty" xml:"AttributeType,omitempty"`
	// 设备形态属性关联对象
	AttributeReference *string `json:"AttributeReference,omitempty" xml:"AttributeReference,omitempty"`
	// 设备形态属性是否表格可见
	AttributeTableDisplay *bool `json:"AttributeTableDisplay,omitempty" xml:"AttributeTableDisplay,omitempty"`
	// 前端查询控件占位符
	AttributePlaceholder *string `json:"AttributePlaceholder,omitempty" xml:"AttributePlaceholder,omitempty"`
	// 前端是否展示对应的查询控件
	AttributeQuery *bool `json:"AttributeQuery,omitempty" xml:"AttributeQuery,omitempty"`
	// 前端查询控件是否支持模糊搜索
	AttributeFuzzyQuery *bool `json:"AttributeFuzzyQuery,omitempty" xml:"AttributeFuzzyQuery,omitempty"`
	// 设备形态属性是否内置
	AttributeBuiltIn *bool `json:"AttributeBuiltIn,omitempty" xml:"AttributeBuiltIn,omitempty"`
}

func (s GetDeviceFormResponseBodyDeviceFormAttributeList) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceFormResponseBodyDeviceFormAttributeList) GoString() string {
	return s.String()
}

func (s *GetDeviceFormResponseBodyDeviceFormAttributeList) SetAttributeKey(v string) *GetDeviceFormResponseBodyDeviceFormAttributeList {
	s.AttributeKey = &v
	return s
}

func (s *GetDeviceFormResponseBodyDeviceFormAttributeList) SetAttributeName(v string) *GetDeviceFormResponseBodyDeviceFormAttributeList {
	s.AttributeName = &v
	return s
}

func (s *GetDeviceFormResponseBodyDeviceFormAttributeList) SetAttributeRequirement(v bool) *GetDeviceFormResponseBodyDeviceFormAttributeList {
	s.AttributeRequirement = &v
	return s
}

func (s *GetDeviceFormResponseBodyDeviceFormAttributeList) SetAttributeUniqueness(v bool) *GetDeviceFormResponseBodyDeviceFormAttributeList {
	s.AttributeUniqueness = &v
	return s
}

func (s *GetDeviceFormResponseBodyDeviceFormAttributeList) SetAttributeFormat(v string) *GetDeviceFormResponseBodyDeviceFormAttributeList {
	s.AttributeFormat = &v
	return s
}

func (s *GetDeviceFormResponseBodyDeviceFormAttributeList) SetAttributeType(v string) *GetDeviceFormResponseBodyDeviceFormAttributeList {
	s.AttributeType = &v
	return s
}

func (s *GetDeviceFormResponseBodyDeviceFormAttributeList) SetAttributeReference(v string) *GetDeviceFormResponseBodyDeviceFormAttributeList {
	s.AttributeReference = &v
	return s
}

func (s *GetDeviceFormResponseBodyDeviceFormAttributeList) SetAttributeTableDisplay(v bool) *GetDeviceFormResponseBodyDeviceFormAttributeList {
	s.AttributeTableDisplay = &v
	return s
}

func (s *GetDeviceFormResponseBodyDeviceFormAttributeList) SetAttributePlaceholder(v string) *GetDeviceFormResponseBodyDeviceFormAttributeList {
	s.AttributePlaceholder = &v
	return s
}

func (s *GetDeviceFormResponseBodyDeviceFormAttributeList) SetAttributeQuery(v bool) *GetDeviceFormResponseBodyDeviceFormAttributeList {
	s.AttributeQuery = &v
	return s
}

func (s *GetDeviceFormResponseBodyDeviceFormAttributeList) SetAttributeFuzzyQuery(v bool) *GetDeviceFormResponseBodyDeviceFormAttributeList {
	s.AttributeFuzzyQuery = &v
	return s
}

func (s *GetDeviceFormResponseBodyDeviceFormAttributeList) SetAttributeBuiltIn(v bool) *GetDeviceFormResponseBodyDeviceFormAttributeList {
	s.AttributeBuiltIn = &v
	return s
}

type GetDeviceFormResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeviceFormResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceFormResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceFormResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceFormResponse) SetHeaders(v map[string]*string) *GetDeviceFormResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceFormResponse) SetBody(v *GetDeviceFormResponseBody) *GetDeviceFormResponse {
	s.Body = v
	return s
}

type CreateDeviceRequest struct {
	// 设备形态ID
	DeviceFormId *string `json:"DeviceFormId,omitempty" xml:"DeviceFormId,omitempty"`
	// 物理空间ID
	PhysicalSpaceId *string `json:"PhysicalSpaceId,omitempty" xml:"PhysicalSpaceId,omitempty"`
	// 主机名
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// 设备IP
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// 设备SN
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// 设备MAC地址
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// 设备厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 设备型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 设备状态
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// 设备安全域
	SecurityDomain *string `json:"SecurityDomain,omitempty" xml:"SecurityDomain,omitempty"`
	// 登录类型
	LoginType *string `json:"LoginType,omitempty" xml:"LoginType,omitempty"`
	// 登录账号
	LoginUsername *string `json:"LoginUsername,omitempty" xml:"LoginUsername,omitempty"`
	// 登录密码
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// enable密码
	EnablePassword *string `json:"EnablePassword,omitempty" xml:"EnablePassword,omitempty"`
	// SNMP 版本号
	SnmpAccountVersion *string `json:"SnmpAccountVersion,omitempty" xml:"SnmpAccountVersion,omitempty"`
	// SNMP Community
	SnmpCommunity *string `json:"SnmpCommunity,omitempty" xml:"SnmpCommunity,omitempty"`
	// SNMP 账号类型
	SnmpAccountType *string `json:"SnmpAccountType,omitempty" xml:"SnmpAccountType,omitempty"`
	// SNMP 安全级别
	SnmpSecurityLevel *string `json:"SnmpSecurityLevel,omitempty" xml:"SnmpSecurityLevel,omitempty"`
	// SNMP 用户名
	SnmpUsername *string `json:"SnmpUsername,omitempty" xml:"SnmpUsername,omitempty"`
	// SNMP Auth PassPhrase
	SnmpAuthPassphrase *string `json:"SnmpAuthPassphrase,omitempty" xml:"SnmpAuthPassphrase,omitempty"`
	// SNMP Auth Protocol
	SnmpAuthProtocol *string `json:"SnmpAuthProtocol,omitempty" xml:"SnmpAuthProtocol,omitempty"`
	// SNMP Privacy Passphase
	SnmpPrivacyPassphase *string `json:"SnmpPrivacyPassphase,omitempty" xml:"SnmpPrivacyPassphase,omitempty"`
	// SNMP Privacy Protocol
	SnmpPrivacyProtocol *string `json:"SnmpPrivacyProtocol,omitempty" xml:"SnmpPrivacyProtocol,omitempty"`
	// 设备额外属性
	ExtAttributes *string `json:"ExtAttributes,omitempty" xml:"ExtAttributes,omitempty"`
	// 幂等校验 token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceRequest) GoString() string {
	return s.String()
}

func (s *CreateDeviceRequest) SetDeviceFormId(v string) *CreateDeviceRequest {
	s.DeviceFormId = &v
	return s
}

func (s *CreateDeviceRequest) SetPhysicalSpaceId(v string) *CreateDeviceRequest {
	s.PhysicalSpaceId = &v
	return s
}

func (s *CreateDeviceRequest) SetHostName(v string) *CreateDeviceRequest {
	s.HostName = &v
	return s
}

func (s *CreateDeviceRequest) SetIp(v string) *CreateDeviceRequest {
	s.Ip = &v
	return s
}

func (s *CreateDeviceRequest) SetSn(v string) *CreateDeviceRequest {
	s.Sn = &v
	return s
}

func (s *CreateDeviceRequest) SetMac(v string) *CreateDeviceRequest {
	s.Mac = &v
	return s
}

func (s *CreateDeviceRequest) SetVendor(v string) *CreateDeviceRequest {
	s.Vendor = &v
	return s
}

func (s *CreateDeviceRequest) SetModel(v string) *CreateDeviceRequest {
	s.Model = &v
	return s
}

func (s *CreateDeviceRequest) SetServiceStatus(v string) *CreateDeviceRequest {
	s.ServiceStatus = &v
	return s
}

func (s *CreateDeviceRequest) SetSecurityDomain(v string) *CreateDeviceRequest {
	s.SecurityDomain = &v
	return s
}

func (s *CreateDeviceRequest) SetLoginType(v string) *CreateDeviceRequest {
	s.LoginType = &v
	return s
}

func (s *CreateDeviceRequest) SetLoginUsername(v string) *CreateDeviceRequest {
	s.LoginUsername = &v
	return s
}

func (s *CreateDeviceRequest) SetLoginPassword(v string) *CreateDeviceRequest {
	s.LoginPassword = &v
	return s
}

func (s *CreateDeviceRequest) SetEnablePassword(v string) *CreateDeviceRequest {
	s.EnablePassword = &v
	return s
}

func (s *CreateDeviceRequest) SetSnmpAccountVersion(v string) *CreateDeviceRequest {
	s.SnmpAccountVersion = &v
	return s
}

func (s *CreateDeviceRequest) SetSnmpCommunity(v string) *CreateDeviceRequest {
	s.SnmpCommunity = &v
	return s
}

func (s *CreateDeviceRequest) SetSnmpAccountType(v string) *CreateDeviceRequest {
	s.SnmpAccountType = &v
	return s
}

func (s *CreateDeviceRequest) SetSnmpSecurityLevel(v string) *CreateDeviceRequest {
	s.SnmpSecurityLevel = &v
	return s
}

func (s *CreateDeviceRequest) SetSnmpUsername(v string) *CreateDeviceRequest {
	s.SnmpUsername = &v
	return s
}

func (s *CreateDeviceRequest) SetSnmpAuthPassphrase(v string) *CreateDeviceRequest {
	s.SnmpAuthPassphrase = &v
	return s
}

func (s *CreateDeviceRequest) SetSnmpAuthProtocol(v string) *CreateDeviceRequest {
	s.SnmpAuthProtocol = &v
	return s
}

func (s *CreateDeviceRequest) SetSnmpPrivacyPassphase(v string) *CreateDeviceRequest {
	s.SnmpPrivacyPassphase = &v
	return s
}

func (s *CreateDeviceRequest) SetSnmpPrivacyProtocol(v string) *CreateDeviceRequest {
	s.SnmpPrivacyProtocol = &v
	return s
}

func (s *CreateDeviceRequest) SetExtAttributes(v string) *CreateDeviceRequest {
	s.ExtAttributes = &v
	return s
}

func (s *CreateDeviceRequest) SetClientToken(v string) *CreateDeviceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDeviceRequest) SetInstanceId(v string) *CreateDeviceRequest {
	s.InstanceId = &v
	return s
}

type CreateDeviceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 资源实例ID，如ECS实例的创建接口CreateInstance应返回InstanceId。
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s CreateDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeviceResponseBody) SetRequestId(v string) *CreateDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeviceResponseBody) SetDeviceId(v string) *CreateDeviceResponseBody {
	s.DeviceId = &v
	return s
}

type CreateDeviceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceResponse) GoString() string {
	return s.String()
}

func (s *CreateDeviceResponse) SetHeaders(v map[string]*string) *CreateDeviceResponse {
	s.Headers = v
	return s
}

func (s *CreateDeviceResponse) SetBody(v *CreateDeviceResponseBody) *CreateDeviceResponse {
	s.Body = v
	return s
}

type CreateSetupProjectRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 创建时间
	DeliveryTime *string `json:"DeliveryTime,omitempty" xml:"DeliveryTime,omitempty"`
	// 物理空间uId
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateSetupProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSetupProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateSetupProjectRequest) SetInstanceId(v string) *CreateSetupProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSetupProjectRequest) SetDeliveryTime(v string) *CreateSetupProjectRequest {
	s.DeliveryTime = &v
	return s
}

func (s *CreateSetupProjectRequest) SetSpaceId(v string) *CreateSetupProjectRequest {
	s.SpaceId = &v
	return s
}

func (s *CreateSetupProjectRequest) SetDescription(v string) *CreateSetupProjectRequest {
	s.Description = &v
	return s
}

type CreateSetupProjectResponseBody struct {
	// 资源实例ID，如ECS实例的创建接口CreateInstance应返回InstanceId。
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSetupProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSetupProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSetupProjectResponseBody) SetSetupProjectId(v string) *CreateSetupProjectResponseBody {
	s.SetupProjectId = &v
	return s
}

func (s *CreateSetupProjectResponseBody) SetRequestId(v string) *CreateSetupProjectResponseBody {
	s.RequestId = &v
	return s
}

type CreateSetupProjectResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSetupProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSetupProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSetupProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateSetupProjectResponse) SetHeaders(v map[string]*string) *CreateSetupProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateSetupProjectResponse) SetBody(v *CreateSetupProjectResponseBody) *CreateSetupProjectResponse {
	s.Body = v
	return s
}

type UpdateDedicatedLineRequest struct {
	// 实例 ID。
	DedicatedLineId *string `json:"DedicatedLineId,omitempty" xml:"DedicatedLineId,omitempty"`
	// 运营商
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// 宽带（Mbps）
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// 专线IP
	DedicatedLineIp *string `json:"DedicatedLineIp,omitempty" xml:"DedicatedLineIp,omitempty"`
	// 专线网关
	DedicatedLineGateway *string `json:"DedicatedLineGateway,omitempty" xml:"DedicatedLineGateway,omitempty"`
	// 专线角色
	DedicatedLineRole *string `json:"DedicatedLineRole,omitempty" xml:"DedicatedLineRole,omitempty"`
	// 关联设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 关联设备端口名称
	DevicePort *string `json:"DevicePort,omitempty" xml:"DevicePort,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateDedicatedLineRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDedicatedLineRequest) GoString() string {
	return s.String()
}

func (s *UpdateDedicatedLineRequest) SetDedicatedLineId(v string) *UpdateDedicatedLineRequest {
	s.DedicatedLineId = &v
	return s
}

func (s *UpdateDedicatedLineRequest) SetIsp(v string) *UpdateDedicatedLineRequest {
	s.Isp = &v
	return s
}

func (s *UpdateDedicatedLineRequest) SetBandwidth(v int32) *UpdateDedicatedLineRequest {
	s.Bandwidth = &v
	return s
}

func (s *UpdateDedicatedLineRequest) SetDedicatedLineIp(v string) *UpdateDedicatedLineRequest {
	s.DedicatedLineIp = &v
	return s
}

func (s *UpdateDedicatedLineRequest) SetDedicatedLineGateway(v string) *UpdateDedicatedLineRequest {
	s.DedicatedLineGateway = &v
	return s
}

func (s *UpdateDedicatedLineRequest) SetDedicatedLineRole(v string) *UpdateDedicatedLineRequest {
	s.DedicatedLineRole = &v
	return s
}

func (s *UpdateDedicatedLineRequest) SetDeviceId(v string) *UpdateDedicatedLineRequest {
	s.DeviceId = &v
	return s
}

func (s *UpdateDedicatedLineRequest) SetDevicePort(v string) *UpdateDedicatedLineRequest {
	s.DevicePort = &v
	return s
}

func (s *UpdateDedicatedLineRequest) SetInstanceId(v string) *UpdateDedicatedLineRequest {
	s.InstanceId = &v
	return s
}

type UpdateDedicatedLineResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDedicatedLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDedicatedLineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDedicatedLineResponseBody) SetRequestId(v string) *UpdateDedicatedLineResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDedicatedLineResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDedicatedLineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDedicatedLineResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDedicatedLineResponse) GoString() string {
	return s.String()
}

func (s *UpdateDedicatedLineResponse) SetHeaders(v map[string]*string) *UpdateDedicatedLineResponse {
	s.Headers = v
	return s
}

func (s *UpdateDedicatedLineResponse) SetBody(v *UpdateDedicatedLineResponseBody) *UpdateDedicatedLineResponse {
	s.Body = v
	return s
}

type GetOsVersionRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例 ID。
	OsVersionId *string `json:"OsVersionId,omitempty" xml:"OsVersionId,omitempty"`
}

func (s GetOsVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOsVersionRequest) GoString() string {
	return s.String()
}

func (s *GetOsVersionRequest) SetInstanceId(v string) *GetOsVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetOsVersionRequest) SetOsVersionId(v string) *GetOsVersionRequest {
	s.OsVersionId = &v
	return s
}

type GetOsVersionResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 数组，返回示例目录。
	OsVersion []*GetOsVersionResponseBodyOsVersion `json:"OsVersion,omitempty" xml:"OsVersion,omitempty" type:"Repeated"`
}

func (s GetOsVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOsVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetOsVersionResponseBody) SetRequestId(v string) *GetOsVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOsVersionResponseBody) SetOsVersion(v []*GetOsVersionResponseBodyOsVersion) *GetOsVersionResponseBody {
	s.OsVersion = v
	return s
}

type GetOsVersionResponseBodyOsVersion struct {
	// 下载路径
	DownloadPath *string `json:"DownloadPath,omitempty" xml:"DownloadPath,omitempty"`
}

func (s GetOsVersionResponseBodyOsVersion) String() string {
	return tea.Prettify(s)
}

func (s GetOsVersionResponseBodyOsVersion) GoString() string {
	return s.String()
}

func (s *GetOsVersionResponseBodyOsVersion) SetDownloadPath(v string) *GetOsVersionResponseBodyOsVersion {
	s.DownloadPath = &v
	return s
}

type GetOsVersionResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOsVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOsVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOsVersionResponse) GoString() string {
	return s.String()
}

func (s *GetOsVersionResponse) SetHeaders(v map[string]*string) *GetOsVersionResponse {
	s.Headers = v
	return s
}

func (s *GetOsVersionResponse) SetBody(v *GetOsVersionResponseBody) *GetOsVersionResponse {
	s.Body = v
	return s
}

type UpdateScheduleDutyRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源一级ID
	ScheduleDutyId *string `json:"ScheduleDutyId,omitempty" xml:"ScheduleDutyId,omitempty"`
	// typeWorkerList
	TypeWorkerList []*UpdateScheduleDutyRequestTypeWorkerList `json:"TypeWorkerList,omitempty" xml:"TypeWorkerList,omitempty" type:"Repeated"`
}

func (s UpdateScheduleDutyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleDutyRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduleDutyRequest) SetInstanceId(v string) *UpdateScheduleDutyRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateScheduleDutyRequest) SetScheduleDutyId(v string) *UpdateScheduleDutyRequest {
	s.ScheduleDutyId = &v
	return s
}

func (s *UpdateScheduleDutyRequest) SetTypeWorkerList(v []*UpdateScheduleDutyRequestTypeWorkerList) *UpdateScheduleDutyRequest {
	s.TypeWorkerList = v
	return s
}

type UpdateScheduleDutyRequestTypeWorkerList struct {
	// scheduleTypeId
	ScheduleTypeId *string `json:"ScheduleTypeId,omitempty" xml:"ScheduleTypeId,omitempty"`
	// 值班人员姓名
	ScheduleWorkerName []*string `json:"ScheduleWorkerName,omitempty" xml:"ScheduleWorkerName,omitempty" type:"Repeated"`
}

func (s UpdateScheduleDutyRequestTypeWorkerList) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleDutyRequestTypeWorkerList) GoString() string {
	return s.String()
}

func (s *UpdateScheduleDutyRequestTypeWorkerList) SetScheduleTypeId(v string) *UpdateScheduleDutyRequestTypeWorkerList {
	s.ScheduleTypeId = &v
	return s
}

func (s *UpdateScheduleDutyRequestTypeWorkerList) SetScheduleWorkerName(v []*string) *UpdateScheduleDutyRequestTypeWorkerList {
	s.ScheduleWorkerName = v
	return s
}

type UpdateScheduleDutyShrinkRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源一级ID
	ScheduleDutyId *string `json:"ScheduleDutyId,omitempty" xml:"ScheduleDutyId,omitempty"`
	// typeWorkerList
	TypeWorkerListShrink *string `json:"TypeWorkerList,omitempty" xml:"TypeWorkerList,omitempty"`
}

func (s UpdateScheduleDutyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleDutyShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduleDutyShrinkRequest) SetInstanceId(v string) *UpdateScheduleDutyShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateScheduleDutyShrinkRequest) SetScheduleDutyId(v string) *UpdateScheduleDutyShrinkRequest {
	s.ScheduleDutyId = &v
	return s
}

func (s *UpdateScheduleDutyShrinkRequest) SetTypeWorkerListShrink(v string) *UpdateScheduleDutyShrinkRequest {
	s.TypeWorkerListShrink = &v
	return s
}

type UpdateScheduleDutyResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateScheduleDutyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleDutyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScheduleDutyResponseBody) SetRequestId(v string) *UpdateScheduleDutyResponseBody {
	s.RequestId = &v
	return s
}

type UpdateScheduleDutyResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateScheduleDutyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateScheduleDutyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleDutyResponse) GoString() string {
	return s.String()
}

func (s *UpdateScheduleDutyResponse) SetHeaders(v map[string]*string) *UpdateScheduleDutyResponse {
	s.Headers = v
	return s
}

func (s *UpdateScheduleDutyResponse) SetBody(v *UpdateScheduleDutyResponseBody) *UpdateScheduleDutyResponse {
	s.Body = v
	return s
}

type CreateIpRecordRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源名称
	IpRecordName *string `json:"IpRecordName,omitempty" xml:"IpRecordName,omitempty"`
	// 业务类型uuid
	BusinessTypeId *string `json:"BusinessTypeId,omitempty" xml:"BusinessTypeId,omitempty"`
	// 创建人
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// 设备列表
	Device []*CreateIpRecordRequestDevice `json:"Device,omitempty" xml:"Device,omitempty" type:"Repeated"`
	// 工单类型 Apply 申请工单 Recycle 释放工单
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// IP地址列表
	IpCode []*string `json:"IpCode,omitempty" xml:"IpCode,omitempty" type:"Repeated"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateIpRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIpRecordRequest) GoString() string {
	return s.String()
}

func (s *CreateIpRecordRequest) SetInstanceId(v string) *CreateIpRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateIpRecordRequest) SetIpRecordName(v string) *CreateIpRecordRequest {
	s.IpRecordName = &v
	return s
}

func (s *CreateIpRecordRequest) SetBusinessTypeId(v string) *CreateIpRecordRequest {
	s.BusinessTypeId = &v
	return s
}

func (s *CreateIpRecordRequest) SetCreator(v string) *CreateIpRecordRequest {
	s.Creator = &v
	return s
}

func (s *CreateIpRecordRequest) SetDevice(v []*CreateIpRecordRequestDevice) *CreateIpRecordRequest {
	s.Device = v
	return s
}

func (s *CreateIpRecordRequest) SetRecordType(v string) *CreateIpRecordRequest {
	s.RecordType = &v
	return s
}

func (s *CreateIpRecordRequest) SetIpCode(v []*string) *CreateIpRecordRequest {
	s.IpCode = v
	return s
}

func (s *CreateIpRecordRequest) SetClientToken(v string) *CreateIpRecordRequest {
	s.ClientToken = &v
	return s
}

type CreateIpRecordRequestDevice struct {
	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// 设备端口
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// 对端设备名称
	RemoteDeviceName *string `json:"RemoteDeviceName,omitempty" xml:"RemoteDeviceName,omitempty"`
	// 对端设备端口
	RemotePort *string `json:"RemotePort,omitempty" xml:"RemotePort,omitempty"`
	// 设备MAC
	DeviceMac *string `json:"DeviceMac,omitempty" xml:"DeviceMac,omitempty"`
	// 园区层级
	ZoneLayer []*CreateIpRecordRequestDeviceZoneLayer `json:"ZoneLayer,omitempty" xml:"ZoneLayer,omitempty" type:"Repeated"`
}

func (s CreateIpRecordRequestDevice) String() string {
	return tea.Prettify(s)
}

func (s CreateIpRecordRequestDevice) GoString() string {
	return s.String()
}

func (s *CreateIpRecordRequestDevice) SetDeviceName(v string) *CreateIpRecordRequestDevice {
	s.DeviceName = &v
	return s
}

func (s *CreateIpRecordRequestDevice) SetPort(v string) *CreateIpRecordRequestDevice {
	s.Port = &v
	return s
}

func (s *CreateIpRecordRequestDevice) SetRemoteDeviceName(v string) *CreateIpRecordRequestDevice {
	s.RemoteDeviceName = &v
	return s
}

func (s *CreateIpRecordRequestDevice) SetRemotePort(v string) *CreateIpRecordRequestDevice {
	s.RemotePort = &v
	return s
}

func (s *CreateIpRecordRequestDevice) SetDeviceMac(v string) *CreateIpRecordRequestDevice {
	s.DeviceMac = &v
	return s
}

func (s *CreateIpRecordRequestDevice) SetZoneLayer(v []*CreateIpRecordRequestDeviceZoneLayer) *CreateIpRecordRequestDevice {
	s.ZoneLayer = v
	return s
}

type CreateIpRecordRequestDeviceZoneLayer struct {
	// 园区层级名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 园区层级值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateIpRecordRequestDeviceZoneLayer) String() string {
	return tea.Prettify(s)
}

func (s CreateIpRecordRequestDeviceZoneLayer) GoString() string {
	return s.String()
}

func (s *CreateIpRecordRequestDeviceZoneLayer) SetName(v string) *CreateIpRecordRequestDeviceZoneLayer {
	s.Name = &v
	return s
}

func (s *CreateIpRecordRequestDeviceZoneLayer) SetValue(v string) *CreateIpRecordRequestDeviceZoneLayer {
	s.Value = &v
	return s
}

type CreateIpRecordShrinkRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源名称
	IpRecordName *string `json:"IpRecordName,omitempty" xml:"IpRecordName,omitempty"`
	// 业务类型uuid
	BusinessTypeId *string `json:"BusinessTypeId,omitempty" xml:"BusinessTypeId,omitempty"`
	// 创建人
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// 设备列表
	DeviceShrink *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// 工单类型 Apply 申请工单 Recycle 释放工单
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// IP地址列表
	IpCodeShrink *string `json:"IpCode,omitempty" xml:"IpCode,omitempty"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateIpRecordShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIpRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateIpRecordShrinkRequest) SetInstanceId(v string) *CreateIpRecordShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateIpRecordShrinkRequest) SetIpRecordName(v string) *CreateIpRecordShrinkRequest {
	s.IpRecordName = &v
	return s
}

func (s *CreateIpRecordShrinkRequest) SetBusinessTypeId(v string) *CreateIpRecordShrinkRequest {
	s.BusinessTypeId = &v
	return s
}

func (s *CreateIpRecordShrinkRequest) SetCreator(v string) *CreateIpRecordShrinkRequest {
	s.Creator = &v
	return s
}

func (s *CreateIpRecordShrinkRequest) SetDeviceShrink(v string) *CreateIpRecordShrinkRequest {
	s.DeviceShrink = &v
	return s
}

func (s *CreateIpRecordShrinkRequest) SetRecordType(v string) *CreateIpRecordShrinkRequest {
	s.RecordType = &v
	return s
}

func (s *CreateIpRecordShrinkRequest) SetIpCodeShrink(v string) *CreateIpRecordShrinkRequest {
	s.IpCodeShrink = &v
	return s
}

func (s *CreateIpRecordShrinkRequest) SetClientToken(v string) *CreateIpRecordShrinkRequest {
	s.ClientToken = &v
	return s
}

type CreateIpRecordResponseBody struct {
	// 资源实例ID，如ECS实例的创建接口CreateInstance应返回InstanceId。
	IpRecordId *string `json:"IpRecordId,omitempty" xml:"IpRecordId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 若创建接口为异步实现，则需返回明确的JobId。
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CreateIpRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIpRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpRecordResponseBody) SetIpRecordId(v string) *CreateIpRecordResponseBody {
	s.IpRecordId = &v
	return s
}

func (s *CreateIpRecordResponseBody) SetRequestId(v string) *CreateIpRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIpRecordResponseBody) SetJobId(v string) *CreateIpRecordResponseBody {
	s.JobId = &v
	return s
}

type CreateIpRecordResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateIpRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIpRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIpRecordResponse) GoString() string {
	return s.String()
}

func (s *CreateIpRecordResponse) SetHeaders(v map[string]*string) *CreateIpRecordResponse {
	s.Headers = v
	return s
}

func (s *CreateIpRecordResponse) SetBody(v *CreateIpRecordResponseBody) *CreateIpRecordResponse {
	s.Body = v
	return s
}

type ListInstancesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例列表
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

type ListInstancesResponseBodyInstances struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// 实例规格
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// 实例开通时间
	InstanceOpenDate *string `json:"InstanceOpenDate,omitempty" xml:"InstanceOpenDate,omitempty"`
	// 实例到期时间
	InstanceEndDate *string `json:"InstanceEndDate,omitempty" xml:"InstanceEndDate,omitempty"`
	// 最大纳管设备数量
	InstanceDeviceMaxCount *string `json:"InstanceDeviceMaxCount,omitempty" xml:"InstanceDeviceMaxCount,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceName(v string) *ListInstancesResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceSpec(v string) *ListInstancesResponseBodyInstances {
	s.InstanceSpec = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceOpenDate(v string) *ListInstancesResponseBodyInstances {
	s.InstanceOpenDate = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceEndDate(v string) *ListInstancesResponseBodyInstances {
	s.InstanceEndDate = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceDeviceMaxCount(v string) *ListInstancesResponseBodyInstances {
	s.InstanceDeviceMaxCount = &v
	return s
}

type ListInstancesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type ListZoneTypesRequest struct {
	// 园区类型名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListZoneTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListZoneTypesRequest) GoString() string {
	return s.String()
}

func (s *ListZoneTypesRequest) SetName(v string) *ListZoneTypesRequest {
	s.Name = &v
	return s
}

func (s *ListZoneTypesRequest) SetMaxResults(v int32) *ListZoneTypesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListZoneTypesRequest) SetNextToken(v string) *ListZoneTypesRequest {
	s.NextToken = &v
	return s
}

func (s *ListZoneTypesRequest) SetInstanceId(v string) *ListZoneTypesRequest {
	s.InstanceId = &v
	return s
}

type ListZoneTypesResponseBody struct {
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 本次读取的最大数据量
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 数组，返回示例目录。
	ZoneType []*ListZoneTypesResponseBodyZoneType `json:"ZoneType,omitempty" xml:"ZoneType,omitempty" type:"Repeated"`
}

func (s ListZoneTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListZoneTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListZoneTypesResponseBody) SetTotalCount(v int32) *ListZoneTypesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListZoneTypesResponseBody) SetMaxResults(v int64) *ListZoneTypesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListZoneTypesResponseBody) SetRequestId(v string) *ListZoneTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListZoneTypesResponseBody) SetNextToken(v int32) *ListZoneTypesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListZoneTypesResponseBody) SetZoneType(v []*ListZoneTypesResponseBodyZoneType) *ListZoneTypesResponseBody {
	s.ZoneType = v
	return s
}

type ListZoneTypesResponseBodyZoneType struct {
	// 园区类型名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 资源名称
	ZoneTypeName *string `json:"ZoneTypeName,omitempty" xml:"ZoneTypeName,omitempty"`
	// 园区层级
	ZoneTypeLayer []*ListZoneTypesResponseBodyZoneTypeZoneTypeLayer `json:"ZoneTypeLayer,omitempty" xml:"ZoneTypeLayer,omitempty" type:"Repeated"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 资源一级ID
	ZoneTypeId *string `json:"ZoneTypeId,omitempty" xml:"ZoneTypeId,omitempty"`
}

func (s ListZoneTypesResponseBodyZoneType) String() string {
	return tea.Prettify(s)
}

func (s ListZoneTypesResponseBodyZoneType) GoString() string {
	return s.String()
}

func (s *ListZoneTypesResponseBodyZoneType) SetName(v string) *ListZoneTypesResponseBodyZoneType {
	s.Name = &v
	return s
}

func (s *ListZoneTypesResponseBodyZoneType) SetZoneTypeName(v string) *ListZoneTypesResponseBodyZoneType {
	s.ZoneTypeName = &v
	return s
}

func (s *ListZoneTypesResponseBodyZoneType) SetZoneTypeLayer(v []*ListZoneTypesResponseBodyZoneTypeZoneTypeLayer) *ListZoneTypesResponseBodyZoneType {
	s.ZoneTypeLayer = v
	return s
}

func (s *ListZoneTypesResponseBodyZoneType) SetCreateTime(v string) *ListZoneTypesResponseBodyZoneType {
	s.CreateTime = &v
	return s
}

func (s *ListZoneTypesResponseBodyZoneType) SetZoneTypeId(v string) *ListZoneTypesResponseBodyZoneType {
	s.ZoneTypeId = &v
	return s
}

type ListZoneTypesResponseBodyZoneTypeZoneTypeLayer struct {
	// 园区层级名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 园区层级掩码
	Mask *int64 `json:"Mask,omitempty" xml:"Mask,omitempty"`
}

func (s ListZoneTypesResponseBodyZoneTypeZoneTypeLayer) String() string {
	return tea.Prettify(s)
}

func (s ListZoneTypesResponseBodyZoneTypeZoneTypeLayer) GoString() string {
	return s.String()
}

func (s *ListZoneTypesResponseBodyZoneTypeZoneTypeLayer) SetName(v string) *ListZoneTypesResponseBodyZoneTypeZoneTypeLayer {
	s.Name = &v
	return s
}

func (s *ListZoneTypesResponseBodyZoneTypeZoneTypeLayer) SetMask(v int64) *ListZoneTypesResponseBodyZoneTypeZoneTypeLayer {
	s.Mask = &v
	return s
}

type ListZoneTypesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListZoneTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListZoneTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListZoneTypesResponse) GoString() string {
	return s.String()
}

func (s *ListZoneTypesResponse) SetHeaders(v map[string]*string) *ListZoneTypesResponse {
	s.Headers = v
	return s
}

func (s *ListZoneTypesResponse) SetBody(v *ListZoneTypesResponseBody) *ListZoneTypesResponse {
	s.Body = v
	return s
}

type UpdatePhysicalSpaceRequest struct {
	// 实例 ID。
	PhysicalSpaceId *string `json:"PhysicalSpaceId,omitempty" xml:"PhysicalSpaceId,omitempty"`
	// 物理空间名称
	PhysicalSpaceName *string `json:"PhysicalSpaceName,omitempty" xml:"PhysicalSpaceName,omitempty"`
	// 所属国家
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// 所属省份
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// 所属城市
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// 具体地址
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 负责人
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// 缩写
	SpaceAbbreviation *string `json:"SpaceAbbreviation,omitempty" xml:"SpaceAbbreviation,omitempty"`
	// 模型
	SpaceType *string `json:"SpaceType,omitempty" xml:"SpaceType,omitempty"`
}

func (s UpdatePhysicalSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhysicalSpaceRequest) GoString() string {
	return s.String()
}

func (s *UpdatePhysicalSpaceRequest) SetPhysicalSpaceId(v string) *UpdatePhysicalSpaceRequest {
	s.PhysicalSpaceId = &v
	return s
}

func (s *UpdatePhysicalSpaceRequest) SetPhysicalSpaceName(v string) *UpdatePhysicalSpaceRequest {
	s.PhysicalSpaceName = &v
	return s
}

func (s *UpdatePhysicalSpaceRequest) SetCountry(v string) *UpdatePhysicalSpaceRequest {
	s.Country = &v
	return s
}

func (s *UpdatePhysicalSpaceRequest) SetProvince(v string) *UpdatePhysicalSpaceRequest {
	s.Province = &v
	return s
}

func (s *UpdatePhysicalSpaceRequest) SetCity(v string) *UpdatePhysicalSpaceRequest {
	s.City = &v
	return s
}

func (s *UpdatePhysicalSpaceRequest) SetAddress(v string) *UpdatePhysicalSpaceRequest {
	s.Address = &v
	return s
}

func (s *UpdatePhysicalSpaceRequest) SetInstanceId(v string) *UpdatePhysicalSpaceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdatePhysicalSpaceRequest) SetOwner(v string) *UpdatePhysicalSpaceRequest {
	s.Owner = &v
	return s
}

func (s *UpdatePhysicalSpaceRequest) SetSpaceAbbreviation(v string) *UpdatePhysicalSpaceRequest {
	s.SpaceAbbreviation = &v
	return s
}

func (s *UpdatePhysicalSpaceRequest) SetSpaceType(v string) *UpdatePhysicalSpaceRequest {
	s.SpaceType = &v
	return s
}

type UpdatePhysicalSpaceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePhysicalSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhysicalSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePhysicalSpaceResponseBody) SetRequestId(v string) *UpdatePhysicalSpaceResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePhysicalSpaceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePhysicalSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePhysicalSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhysicalSpaceResponse) GoString() string {
	return s.String()
}

func (s *UpdatePhysicalSpaceResponse) SetHeaders(v map[string]*string) *UpdatePhysicalSpaceResponse {
	s.Headers = v
	return s
}

func (s *UpdatePhysicalSpaceResponse) SetBody(v *UpdatePhysicalSpaceResponseBody) *UpdatePhysicalSpaceResponse {
	s.Body = v
	return s
}

type UpdateResourceInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源id
	ResourceInformationId *string `json:"ResourceInformationId,omitempty" xml:"ResourceInformationId,omitempty"`
	// 建设项目资源uuid
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// 数据
	ResourceInformation []*UpdateResourceInstanceRequestResourceInformation `json:"ResourceInformation,omitempty" xml:"ResourceInformation,omitempty" type:"Repeated"`
}

func (s UpdateResourceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceInstanceRequest) SetInstanceId(v string) *UpdateResourceInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateResourceInstanceRequest) SetResourceInformationId(v string) *UpdateResourceInstanceRequest {
	s.ResourceInformationId = &v
	return s
}

func (s *UpdateResourceInstanceRequest) SetSetupProjectId(v string) *UpdateResourceInstanceRequest {
	s.SetupProjectId = &v
	return s
}

func (s *UpdateResourceInstanceRequest) SetResourceInformation(v []*UpdateResourceInstanceRequestResourceInformation) *UpdateResourceInstanceRequest {
	s.ResourceInformation = v
	return s
}

type UpdateResourceInstanceRequestResourceInformation struct {
	// 键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateResourceInstanceRequestResourceInformation) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceInstanceRequestResourceInformation) GoString() string {
	return s.String()
}

func (s *UpdateResourceInstanceRequestResourceInformation) SetKey(v string) *UpdateResourceInstanceRequestResourceInformation {
	s.Key = &v
	return s
}

func (s *UpdateResourceInstanceRequestResourceInformation) SetValue(v string) *UpdateResourceInstanceRequestResourceInformation {
	s.Value = &v
	return s
}

type UpdateResourceInstanceShrinkRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源id
	ResourceInformationId *string `json:"ResourceInformationId,omitempty" xml:"ResourceInformationId,omitempty"`
	// 建设项目资源uuid
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// 数据
	ResourceInformationShrink *string `json:"ResourceInformation,omitempty" xml:"ResourceInformation,omitempty"`
}

func (s UpdateResourceInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceInstanceShrinkRequest) SetInstanceId(v string) *UpdateResourceInstanceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateResourceInstanceShrinkRequest) SetResourceInformationId(v string) *UpdateResourceInstanceShrinkRequest {
	s.ResourceInformationId = &v
	return s
}

func (s *UpdateResourceInstanceShrinkRequest) SetSetupProjectId(v string) *UpdateResourceInstanceShrinkRequest {
	s.SetupProjectId = &v
	return s
}

func (s *UpdateResourceInstanceShrinkRequest) SetResourceInformationShrink(v string) *UpdateResourceInstanceShrinkRequest {
	s.ResourceInformationShrink = &v
	return s
}

type UpdateResourceInstanceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceInstanceResponseBody) SetRequestId(v string) *UpdateResourceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateResourceInstanceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateResourceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateResourceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceInstanceResponse) SetHeaders(v map[string]*string) *UpdateResourceInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceInstanceResponse) SetBody(v *UpdateResourceInstanceResponseBody) *UpdateResourceInstanceResponse {
	s.Body = v
	return s
}

type GetScheduleDutyRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// scheduleDutyId
	ScheduleDutyId *string `json:"ScheduleDutyId,omitempty" xml:"ScheduleDutyId,omitempty"`
}

func (s GetScheduleDutyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleDutyRequest) GoString() string {
	return s.String()
}

func (s *GetScheduleDutyRequest) SetInstanceId(v string) *GetScheduleDutyRequest {
	s.InstanceId = &v
	return s
}

func (s *GetScheduleDutyRequest) SetScheduleDutyId(v string) *GetScheduleDutyRequest {
	s.ScheduleDutyId = &v
	return s
}

type GetScheduleDutyResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 值班表对象
	ScheduleDuty *GetScheduleDutyResponseBodyScheduleDuty `json:"ScheduleDuty,omitempty" xml:"ScheduleDuty,omitempty" type:"Struct"`
}

func (s GetScheduleDutyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleDutyResponseBody) GoString() string {
	return s.String()
}

func (s *GetScheduleDutyResponseBody) SetRequestId(v string) *GetScheduleDutyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScheduleDutyResponseBody) SetScheduleDuty(v *GetScheduleDutyResponseBodyScheduleDuty) *GetScheduleDutyResponseBody {
	s.ScheduleDuty = v
	return s
}

type GetScheduleDutyResponseBodyScheduleDuty struct {
	// 值班表类型
	WorkType *string `json:"WorkType,omitempty" xml:"WorkType,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 值班表日期
	WorkDate *string `json:"WorkDate,omitempty" xml:"WorkDate,omitempty"`
	// 值班人员工号
	WorkerId *string `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
	// 资源一级ID
	ScheduleDutyId *string `json:"ScheduleDutyId,omitempty" xml:"ScheduleDutyId,omitempty"`
	// 值班人员姓名
	WorkerName *string `json:"WorkerName,omitempty" xml:"WorkerName,omitempty"`
}

func (s GetScheduleDutyResponseBodyScheduleDuty) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleDutyResponseBodyScheduleDuty) GoString() string {
	return s.String()
}

func (s *GetScheduleDutyResponseBodyScheduleDuty) SetWorkType(v string) *GetScheduleDutyResponseBodyScheduleDuty {
	s.WorkType = &v
	return s
}

func (s *GetScheduleDutyResponseBodyScheduleDuty) SetCreateTime(v string) *GetScheduleDutyResponseBodyScheduleDuty {
	s.CreateTime = &v
	return s
}

func (s *GetScheduleDutyResponseBodyScheduleDuty) SetUpdateTime(v string) *GetScheduleDutyResponseBodyScheduleDuty {
	s.UpdateTime = &v
	return s
}

func (s *GetScheduleDutyResponseBodyScheduleDuty) SetWorkDate(v string) *GetScheduleDutyResponseBodyScheduleDuty {
	s.WorkDate = &v
	return s
}

func (s *GetScheduleDutyResponseBodyScheduleDuty) SetWorkerId(v string) *GetScheduleDutyResponseBodyScheduleDuty {
	s.WorkerId = &v
	return s
}

func (s *GetScheduleDutyResponseBodyScheduleDuty) SetScheduleDutyId(v string) *GetScheduleDutyResponseBodyScheduleDuty {
	s.ScheduleDutyId = &v
	return s
}

func (s *GetScheduleDutyResponseBodyScheduleDuty) SetWorkerName(v string) *GetScheduleDutyResponseBodyScheduleDuty {
	s.WorkerName = &v
	return s
}

type GetScheduleDutyResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetScheduleDutyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetScheduleDutyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleDutyResponse) GoString() string {
	return s.String()
}

func (s *GetScheduleDutyResponse) SetHeaders(v map[string]*string) *GetScheduleDutyResponse {
	s.Headers = v
	return s
}

func (s *GetScheduleDutyResponse) SetBody(v *GetScheduleDutyResponseBody) *GetScheduleDutyResponse {
	s.Body = v
	return s
}

type GetConfigurationVariateRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例 ID。
	ConfigurationVariateId *string `json:"ConfigurationVariateId,omitempty" xml:"ConfigurationVariateId,omitempty"`
}

func (s GetConfigurationVariateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConfigurationVariateRequest) GoString() string {
	return s.String()
}

func (s *GetConfigurationVariateRequest) SetInstanceId(v string) *GetConfigurationVariateRequest {
	s.InstanceId = &v
	return s
}

func (s *GetConfigurationVariateRequest) SetConfigurationVariateId(v string) *GetConfigurationVariateRequest {
	s.ConfigurationVariateId = &v
	return s
}

type GetConfigurationVariateResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 资源对象
	ConfigurationVariate *GetConfigurationVariateResponseBodyConfigurationVariate `json:"ConfigurationVariate,omitempty" xml:"ConfigurationVariate,omitempty" type:"Struct"`
}

func (s GetConfigurationVariateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConfigurationVariateResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfigurationVariateResponseBody) SetRequestId(v string) *GetConfigurationVariateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConfigurationVariateResponseBody) SetConfigurationVariate(v *GetConfigurationVariateResponseBodyConfigurationVariate) *GetConfigurationVariateResponseBody {
	s.ConfigurationVariate = v
	return s
}

type GetConfigurationVariateResponseBodyConfigurationVariate struct {
	// project
	VariateName *string `json:"VariateName,omitempty" xml:"VariateName,omitempty"`
	// 变量描述
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// 转换函数
	FormatFunction *string `json:"FormatFunction,omitempty" xml:"FormatFunction,omitempty"`
}

func (s GetConfigurationVariateResponseBodyConfigurationVariate) String() string {
	return tea.Prettify(s)
}

func (s GetConfigurationVariateResponseBodyConfigurationVariate) GoString() string {
	return s.String()
}

func (s *GetConfigurationVariateResponseBodyConfigurationVariate) SetVariateName(v string) *GetConfigurationVariateResponseBodyConfigurationVariate {
	s.VariateName = &v
	return s
}

func (s *GetConfigurationVariateResponseBodyConfigurationVariate) SetComment(v string) *GetConfigurationVariateResponseBodyConfigurationVariate {
	s.Comment = &v
	return s
}

func (s *GetConfigurationVariateResponseBodyConfigurationVariate) SetFormatFunction(v string) *GetConfigurationVariateResponseBodyConfigurationVariate {
	s.FormatFunction = &v
	return s
}

type GetConfigurationVariateResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetConfigurationVariateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConfigurationVariateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConfigurationVariateResponse) GoString() string {
	return s.String()
}

func (s *GetConfigurationVariateResponse) SetHeaders(v map[string]*string) *GetConfigurationVariateResponse {
	s.Headers = v
	return s
}

func (s *GetConfigurationVariateResponse) SetBody(v *GetConfigurationVariateResponseBody) *GetConfigurationVariateResponse {
	s.Body = v
	return s
}

type GetAlarmStatusRequest struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 监控项ID
	MonitorItemId *string `json:"MonitorItemId,omitempty" xml:"MonitorItemId,omitempty"`
	// 数据类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 聚合数据ID
	AggregateDataId *string `json:"AggregateDataId,omitempty" xml:"AggregateDataId,omitempty"`
	// 专线ID
	DedicatedLineId *string `json:"DedicatedLineId,omitempty" xml:"DedicatedLineId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 端口集ID
	PortCollectionId *string `json:"PortCollectionId,omitempty" xml:"PortCollectionId,omitempty"`
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetAlarmStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAlarmStatusRequest) SetDeviceId(v string) *GetAlarmStatusRequest {
	s.DeviceId = &v
	return s
}

func (s *GetAlarmStatusRequest) SetMonitorItemId(v string) *GetAlarmStatusRequest {
	s.MonitorItemId = &v
	return s
}

func (s *GetAlarmStatusRequest) SetType(v string) *GetAlarmStatusRequest {
	s.Type = &v
	return s
}

func (s *GetAlarmStatusRequest) SetAggregateDataId(v string) *GetAlarmStatusRequest {
	s.AggregateDataId = &v
	return s
}

func (s *GetAlarmStatusRequest) SetDedicatedLineId(v string) *GetAlarmStatusRequest {
	s.DedicatedLineId = &v
	return s
}

func (s *GetAlarmStatusRequest) SetInstanceId(v string) *GetAlarmStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAlarmStatusRequest) SetPortCollectionId(v string) *GetAlarmStatusRequest {
	s.PortCollectionId = &v
	return s
}

func (s *GetAlarmStatusRequest) SetAppId(v string) *GetAlarmStatusRequest {
	s.AppId = &v
	return s
}

type GetAlarmStatusResponseBody struct {
	// request Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 告警状态
	AlarmStatus *GetAlarmStatusResponseBodyAlarmStatus `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty" type:"Struct"`
}

func (s GetAlarmStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlarmStatusResponseBody) SetRequestId(v string) *GetAlarmStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlarmStatusResponseBody) SetAlarmStatus(v *GetAlarmStatusResponseBodyAlarmStatus) *GetAlarmStatusResponseBody {
	s.AlarmStatus = v
	return s
}

type GetAlarmStatusResponseBodyAlarmStatus struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 监控项ID
	MonitorItemId *string `json:"MonitorItemId,omitempty" xml:"MonitorItemId,omitempty"`
	// 采集时间
	CollectionTime *string `json:"CollectionTime,omitempty" xml:"CollectionTime,omitempty"`
	// 接收时间
	ReceiveTime *string `json:"ReceiveTime,omitempty" xml:"ReceiveTime,omitempty"`
	// 命中告警规则
	AlarmRule *string `json:"AlarmRule,omitempty" xml:"AlarmRule,omitempty"`
	// 告警状态
	AlarmStatus *string `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	// 采集结果
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// 异常数据项
	AbnormalDataItem *string `json:"AbnormalDataItem,omitempty" xml:"AbnormalDataItem,omitempty"`
	// 索引
	UniqueKey *string `json:"UniqueKey,omitempty" xml:"UniqueKey,omitempty"`
	// 采集状态码
	ResponseCode *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	// 设备
	ResourceDevice *GetAlarmStatusResponseBodyAlarmStatusResourceDevice `json:"ResourceDevice,omitempty" xml:"ResourceDevice,omitempty" type:"Struct"`
	// 监控项
	MonitorItem *GetAlarmStatusResponseBodyAlarmStatusMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Struct"`
	// 首次异常时间
	FirstAbnormalTime *string `json:"FirstAbnormalTime,omitempty" xml:"FirstAbnormalTime,omitempty"`
	// 告警开关
	NotificationSwitch *GetAlarmStatusResponseBodyAlarmStatusNotificationSwitch `json:"NotificationSwitch,omitempty" xml:"NotificationSwitch,omitempty" type:"Struct"`
	// 聚合数据ID
	AggregateDataId *string `json:"AggregateDataId,omitempty" xml:"AggregateDataId,omitempty"`
	// 专线ID
	DedicatedLineId *string `json:"DedicatedLineId,omitempty" xml:"DedicatedLineId,omitempty"`
	// 聚合数据详情
	AggregateData *GetAlarmStatusResponseBodyAlarmStatusAggregateData `json:"AggregateData,omitempty" xml:"AggregateData,omitempty" type:"Struct"`
	// 专线详情
	DedicatedLine *GetAlarmStatusResponseBodyAlarmStatusDedicatedLine `json:"DedicatedLine,omitempty" xml:"DedicatedLine,omitempty" type:"Struct"`
	// 端口集ID
	PortCollectionId *string                                              `json:"PortCollectionId,omitempty" xml:"PortCollectionId,omitempty"`
	PortCollection   *GetAlarmStatusResponseBodyAlarmStatusPortCollection `json:"PortCollection,omitempty" xml:"PortCollection,omitempty" type:"Struct"`
	// 采集探针IP
	AgentIp *string `json:"AgentIp,omitempty" xml:"AgentIp,omitempty"`
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用
	ResourceApp *GetAlarmStatusResponseBodyAlarmStatusResourceApp `json:"ResourceApp,omitempty" xml:"ResourceApp,omitempty" type:"Struct"`
}

func (s GetAlarmStatusResponseBodyAlarmStatus) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmStatusResponseBodyAlarmStatus) GoString() string {
	return s.String()
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetDeviceId(v string) *GetAlarmStatusResponseBodyAlarmStatus {
	s.DeviceId = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetMonitorItemId(v string) *GetAlarmStatusResponseBodyAlarmStatus {
	s.MonitorItemId = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetCollectionTime(v string) *GetAlarmStatusResponseBodyAlarmStatus {
	s.CollectionTime = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetReceiveTime(v string) *GetAlarmStatusResponseBodyAlarmStatus {
	s.ReceiveTime = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetAlarmRule(v string) *GetAlarmStatusResponseBodyAlarmStatus {
	s.AlarmRule = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetAlarmStatus(v string) *GetAlarmStatusResponseBodyAlarmStatus {
	s.AlarmStatus = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetResult(v string) *GetAlarmStatusResponseBodyAlarmStatus {
	s.Result = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetAbnormalDataItem(v string) *GetAlarmStatusResponseBodyAlarmStatus {
	s.AbnormalDataItem = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetUniqueKey(v string) *GetAlarmStatusResponseBodyAlarmStatus {
	s.UniqueKey = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetResponseCode(v string) *GetAlarmStatusResponseBodyAlarmStatus {
	s.ResponseCode = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetResourceDevice(v *GetAlarmStatusResponseBodyAlarmStatusResourceDevice) *GetAlarmStatusResponseBodyAlarmStatus {
	s.ResourceDevice = v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetMonitorItem(v *GetAlarmStatusResponseBodyAlarmStatusMonitorItem) *GetAlarmStatusResponseBodyAlarmStatus {
	s.MonitorItem = v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetFirstAbnormalTime(v string) *GetAlarmStatusResponseBodyAlarmStatus {
	s.FirstAbnormalTime = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetNotificationSwitch(v *GetAlarmStatusResponseBodyAlarmStatusNotificationSwitch) *GetAlarmStatusResponseBodyAlarmStatus {
	s.NotificationSwitch = v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetAggregateDataId(v string) *GetAlarmStatusResponseBodyAlarmStatus {
	s.AggregateDataId = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetDedicatedLineId(v string) *GetAlarmStatusResponseBodyAlarmStatus {
	s.DedicatedLineId = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetAggregateData(v *GetAlarmStatusResponseBodyAlarmStatusAggregateData) *GetAlarmStatusResponseBodyAlarmStatus {
	s.AggregateData = v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetDedicatedLine(v *GetAlarmStatusResponseBodyAlarmStatusDedicatedLine) *GetAlarmStatusResponseBodyAlarmStatus {
	s.DedicatedLine = v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetPortCollectionId(v string) *GetAlarmStatusResponseBodyAlarmStatus {
	s.PortCollectionId = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetPortCollection(v *GetAlarmStatusResponseBodyAlarmStatusPortCollection) *GetAlarmStatusResponseBodyAlarmStatus {
	s.PortCollection = v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetAgentIp(v string) *GetAlarmStatusResponseBodyAlarmStatus {
	s.AgentIp = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetAppId(v string) *GetAlarmStatusResponseBodyAlarmStatus {
	s.AppId = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatus) SetResourceApp(v *GetAlarmStatusResponseBodyAlarmStatusResourceApp) *GetAlarmStatusResponseBodyAlarmStatus {
	s.ResourceApp = v
	return s
}

type GetAlarmStatusResponseBodyAlarmStatusResourceDevice struct {
	// 设备名
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// IP
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// sn
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// 物理空间
	Space *string `json:"Space,omitempty" xml:"Space,omitempty"`
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 安全域
	SecurityDomain *string `json:"SecurityDomain,omitempty" xml:"SecurityDomain,omitempty"`
	// 设备形态
	DeviceForm *string `json:"DeviceForm,omitempty" xml:"DeviceForm,omitempty"`
}

func (s GetAlarmStatusResponseBodyAlarmStatusResourceDevice) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmStatusResponseBodyAlarmStatusResourceDevice) GoString() string {
	return s.String()
}

func (s *GetAlarmStatusResponseBodyAlarmStatusResourceDevice) SetHostName(v string) *GetAlarmStatusResponseBodyAlarmStatusResourceDevice {
	s.HostName = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusResourceDevice) SetIp(v string) *GetAlarmStatusResponseBodyAlarmStatusResourceDevice {
	s.Ip = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusResourceDevice) SetVendor(v string) *GetAlarmStatusResponseBodyAlarmStatusResourceDevice {
	s.Vendor = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusResourceDevice) SetModel(v string) *GetAlarmStatusResponseBodyAlarmStatusResourceDevice {
	s.Model = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusResourceDevice) SetStatus(v string) *GetAlarmStatusResponseBodyAlarmStatusResourceDevice {
	s.Status = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusResourceDevice) SetSn(v string) *GetAlarmStatusResponseBodyAlarmStatusResourceDevice {
	s.Sn = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusResourceDevice) SetSpace(v string) *GetAlarmStatusResponseBodyAlarmStatusResourceDevice {
	s.Space = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusResourceDevice) SetDeviceId(v string) *GetAlarmStatusResponseBodyAlarmStatusResourceDevice {
	s.DeviceId = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusResourceDevice) SetSecurityDomain(v string) *GetAlarmStatusResponseBodyAlarmStatusResourceDevice {
	s.SecurityDomain = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusResourceDevice) SetDeviceForm(v string) *GetAlarmStatusResponseBodyAlarmStatusResourceDevice {
	s.DeviceForm = &v
	return s
}

type GetAlarmStatusResponseBodyAlarmStatusMonitorItem struct {
	// 监控项名称
	MonitorItemName *string `json:"MonitorItemName,omitempty" xml:"MonitorItemName,omitempty"`
	// 描述
	MonitorItemDescription *string `json:"MonitorItemDescription,omitempty" xml:"MonitorItemDescription,omitempty"`
	// 安全域
	SecurityDomain *string `json:"SecurityDomain,omitempty" xml:"SecurityDomain,omitempty"`
	// 采集类型
	CollectionType *string `json:"CollectionType,omitempty" xml:"CollectionType,omitempty"`
	// 执行间隔
	ExecInterval *string `json:"ExecInterval,omitempty" xml:"ExecInterval,omitempty"`
	// 监控项ID
	MonitorItemId *string `json:"MonitorItemId,omitempty" xml:"MonitorItemId,omitempty"`
	// 设备形态
	DeviceForm *string `json:"DeviceForm,omitempty" xml:"DeviceForm,omitempty"`
	// 是否启用
	Effective *int64 `json:"Effective,omitempty" xml:"Effective,omitempty"`
}

func (s GetAlarmStatusResponseBodyAlarmStatusMonitorItem) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmStatusResponseBodyAlarmStatusMonitorItem) GoString() string {
	return s.String()
}

func (s *GetAlarmStatusResponseBodyAlarmStatusMonitorItem) SetMonitorItemName(v string) *GetAlarmStatusResponseBodyAlarmStatusMonitorItem {
	s.MonitorItemName = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusMonitorItem) SetMonitorItemDescription(v string) *GetAlarmStatusResponseBodyAlarmStatusMonitorItem {
	s.MonitorItemDescription = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusMonitorItem) SetSecurityDomain(v string) *GetAlarmStatusResponseBodyAlarmStatusMonitorItem {
	s.SecurityDomain = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusMonitorItem) SetCollectionType(v string) *GetAlarmStatusResponseBodyAlarmStatusMonitorItem {
	s.CollectionType = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusMonitorItem) SetExecInterval(v string) *GetAlarmStatusResponseBodyAlarmStatusMonitorItem {
	s.ExecInterval = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusMonitorItem) SetMonitorItemId(v string) *GetAlarmStatusResponseBodyAlarmStatusMonitorItem {
	s.MonitorItemId = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusMonitorItem) SetDeviceForm(v string) *GetAlarmStatusResponseBodyAlarmStatusMonitorItem {
	s.DeviceForm = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusMonitorItem) SetEffective(v int64) *GetAlarmStatusResponseBodyAlarmStatusMonitorItem {
	s.Effective = &v
	return s
}

type GetAlarmStatusResponseBodyAlarmStatusNotificationSwitch struct {
	// 关闭原因
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// 关闭到期时间
	ExpiryTime *string `json:"ExpiryTime,omitempty" xml:"ExpiryTime,omitempty"`
}

func (s GetAlarmStatusResponseBodyAlarmStatusNotificationSwitch) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmStatusResponseBodyAlarmStatusNotificationSwitch) GoString() string {
	return s.String()
}

func (s *GetAlarmStatusResponseBodyAlarmStatusNotificationSwitch) SetReason(v string) *GetAlarmStatusResponseBodyAlarmStatusNotificationSwitch {
	s.Reason = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusNotificationSwitch) SetExpiryTime(v string) *GetAlarmStatusResponseBodyAlarmStatusNotificationSwitch {
	s.ExpiryTime = &v
	return s
}

type GetAlarmStatusResponseBodyAlarmStatusAggregateData struct {
	// 聚合方式
	AggregateMode *string `json:"AggregateMode,omitempty" xml:"AggregateMode,omitempty"`
	// 描述
	AggregateDataDescription *string `json:"AggregateDataDescription,omitempty" xml:"AggregateDataDescription,omitempty"`
	// 数据项
	DataItem *string `json:"DataItem,omitempty" xml:"DataItem,omitempty"`
	// 聚合数据名称
	AggregateDataName *string `json:"AggregateDataName,omitempty" xml:"AggregateDataName,omitempty"`
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 是否聚合全部设备
	IsAllDevice *int32 `json:"IsAllDevice,omitempty" xml:"IsAllDevice,omitempty"`
	// 聚合数据ID
	AggregateDataId *string `json:"AggregateDataId,omitempty" xml:"AggregateDataId,omitempty"`
	// 监控项ID
	MonitorItemId *string `json:"MonitorItemId,omitempty" xml:"MonitorItemId,omitempty"`
}

func (s GetAlarmStatusResponseBodyAlarmStatusAggregateData) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmStatusResponseBodyAlarmStatusAggregateData) GoString() string {
	return s.String()
}

func (s *GetAlarmStatusResponseBodyAlarmStatusAggregateData) SetAggregateMode(v string) *GetAlarmStatusResponseBodyAlarmStatusAggregateData {
	s.AggregateMode = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusAggregateData) SetAggregateDataDescription(v string) *GetAlarmStatusResponseBodyAlarmStatusAggregateData {
	s.AggregateDataDescription = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusAggregateData) SetDataItem(v string) *GetAlarmStatusResponseBodyAlarmStatusAggregateData {
	s.DataItem = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusAggregateData) SetAggregateDataName(v string) *GetAlarmStatusResponseBodyAlarmStatusAggregateData {
	s.AggregateDataName = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusAggregateData) SetDeviceId(v string) *GetAlarmStatusResponseBodyAlarmStatusAggregateData {
	s.DeviceId = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusAggregateData) SetIsAllDevice(v int32) *GetAlarmStatusResponseBodyAlarmStatusAggregateData {
	s.IsAllDevice = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusAggregateData) SetAggregateDataId(v string) *GetAlarmStatusResponseBodyAlarmStatusAggregateData {
	s.AggregateDataId = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusAggregateData) SetMonitorItemId(v string) *GetAlarmStatusResponseBodyAlarmStatusAggregateData {
	s.MonitorItemId = &v
	return s
}

type GetAlarmStatusResponseBodyAlarmStatusDedicatedLine struct {
	// 专线名称
	DedicatedLineName *string `json:"DedicatedLineName,omitempty" xml:"DedicatedLineName,omitempty"`
	// 物理空间
	Space *string `json:"Space,omitempty" xml:"Space,omitempty"`
	// 端口名
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 带宽
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// IP
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// 对端IP
	DedicatedLineGateway *string `json:"DedicatedLineGateway,omitempty" xml:"DedicatedLineGateway,omitempty"`
}

func (s GetAlarmStatusResponseBodyAlarmStatusDedicatedLine) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmStatusResponseBodyAlarmStatusDedicatedLine) GoString() string {
	return s.String()
}

func (s *GetAlarmStatusResponseBodyAlarmStatusDedicatedLine) SetDedicatedLineName(v string) *GetAlarmStatusResponseBodyAlarmStatusDedicatedLine {
	s.DedicatedLineName = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusDedicatedLine) SetSpace(v string) *GetAlarmStatusResponseBodyAlarmStatusDedicatedLine {
	s.Space = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusDedicatedLine) SetPortName(v string) *GetAlarmStatusResponseBodyAlarmStatusDedicatedLine {
	s.PortName = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusDedicatedLine) SetDeviceId(v string) *GetAlarmStatusResponseBodyAlarmStatusDedicatedLine {
	s.DeviceId = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusDedicatedLine) SetBandwidth(v string) *GetAlarmStatusResponseBodyAlarmStatusDedicatedLine {
	s.Bandwidth = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusDedicatedLine) SetIp(v string) *GetAlarmStatusResponseBodyAlarmStatusDedicatedLine {
	s.Ip = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusDedicatedLine) SetDedicatedLineGateway(v string) *GetAlarmStatusResponseBodyAlarmStatusDedicatedLine {
	s.DedicatedLineGateway = &v
	return s
}

type GetAlarmStatusResponseBodyAlarmStatusPortCollection struct {
	// 端口集名称
	PortCollectionName *string `json:"PortCollectionName,omitempty" xml:"PortCollectionName,omitempty"`
	// 端口集ID
	PortCollectionId *string `json:"PortCollectionId,omitempty" xml:"PortCollectionId,omitempty"`
	// 端口集描述
	PortCollectionDescription *string `json:"PortCollectionDescription,omitempty" xml:"PortCollectionDescription,omitempty"`
	// 端口列表
	PortList []*GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortList `json:"PortList,omitempty" xml:"PortList,omitempty" type:"Repeated"`
}

func (s GetAlarmStatusResponseBodyAlarmStatusPortCollection) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmStatusResponseBodyAlarmStatusPortCollection) GoString() string {
	return s.String()
}

func (s *GetAlarmStatusResponseBodyAlarmStatusPortCollection) SetPortCollectionName(v string) *GetAlarmStatusResponseBodyAlarmStatusPortCollection {
	s.PortCollectionName = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusPortCollection) SetPortCollectionId(v string) *GetAlarmStatusResponseBodyAlarmStatusPortCollection {
	s.PortCollectionId = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusPortCollection) SetPortCollectionDescription(v string) *GetAlarmStatusResponseBodyAlarmStatusPortCollection {
	s.PortCollectionDescription = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusPortCollection) SetPortList(v []*GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortList) *GetAlarmStatusResponseBodyAlarmStatusPortCollection {
	s.PortList = v
	return s
}

type GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortList struct {
	// 端口名
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 设备详情
	ResourceDevice *GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortListResourceDevice `json:"ResourceDevice,omitempty" xml:"ResourceDevice,omitempty" type:"Struct"`
}

func (s GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortList) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortList) GoString() string {
	return s.String()
}

func (s *GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortList) SetPortName(v string) *GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortList {
	s.PortName = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortList) SetDeviceId(v string) *GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortList {
	s.DeviceId = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortList) SetResourceDevice(v *GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortListResourceDevice) *GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortList {
	s.ResourceDevice = v
	return s
}

type GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortListResourceDevice struct {
	// 设备名
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// IP
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// 安全域
	SecurityDomain *string `json:"SecurityDomain,omitempty" xml:"SecurityDomain,omitempty"`
}

func (s GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortListResourceDevice) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortListResourceDevice) GoString() string {
	return s.String()
}

func (s *GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortListResourceDevice) SetHostName(v string) *GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortListResourceDevice {
	s.HostName = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortListResourceDevice) SetIp(v string) *GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortListResourceDevice {
	s.Ip = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortListResourceDevice) SetSecurityDomain(v string) *GetAlarmStatusResponseBodyAlarmStatusPortCollectionPortListResourceDevice {
	s.SecurityDomain = &v
	return s
}

type GetAlarmStatusResponseBodyAlarmStatusResourceApp struct {
	// 监控域名
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 端口
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// 资源类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 所属探针
	SecurityDomain *string `json:"SecurityDomain,omitempty" xml:"SecurityDomain,omitempty"`
}

func (s GetAlarmStatusResponseBodyAlarmStatusResourceApp) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmStatusResponseBodyAlarmStatusResourceApp) GoString() string {
	return s.String()
}

func (s *GetAlarmStatusResponseBodyAlarmStatusResourceApp) SetDomain(v string) *GetAlarmStatusResponseBodyAlarmStatusResourceApp {
	s.Domain = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusResourceApp) SetAppId(v string) *GetAlarmStatusResponseBodyAlarmStatusResourceApp {
	s.AppId = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusResourceApp) SetPort(v string) *GetAlarmStatusResponseBodyAlarmStatusResourceApp {
	s.Port = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusResourceApp) SetType(v string) *GetAlarmStatusResponseBodyAlarmStatusResourceApp {
	s.Type = &v
	return s
}

func (s *GetAlarmStatusResponseBodyAlarmStatusResourceApp) SetSecurityDomain(v string) *GetAlarmStatusResponseBodyAlarmStatusResourceApp {
	s.SecurityDomain = &v
	return s
}

type GetAlarmStatusResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAlarmStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAlarmStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAlarmStatusResponse) SetHeaders(v map[string]*string) *GetAlarmStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAlarmStatusResponse) SetBody(v *GetAlarmStatusResponseBody) *GetAlarmStatusResponse {
	s.Body = v
	return s
}

type GetIpRecordRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例 ID。
	IpRecordId *string `json:"IpRecordId,omitempty" xml:"IpRecordId,omitempty"`
}

func (s GetIpRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIpRecordRequest) GoString() string {
	return s.String()
}

func (s *GetIpRecordRequest) SetInstanceId(v string) *GetIpRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *GetIpRecordRequest) SetIpRecordId(v string) *GetIpRecordRequest {
	s.IpRecordId = &v
	return s
}

type GetIpRecordResponseBody struct {
	// 工单类型对象
	IpRecord *GetIpRecordResponseBodyIpRecord `json:"IpRecord,omitempty" xml:"IpRecord,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIpRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIpRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetIpRecordResponseBody) SetIpRecord(v *GetIpRecordResponseBodyIpRecord) *GetIpRecordResponseBody {
	s.IpRecord = v
	return s
}

func (s *GetIpRecordResponseBody) SetRequestId(v string) *GetIpRecordResponseBody {
	s.RequestId = &v
	return s
}

type GetIpRecordResponseBodyIpRecord struct {
	// 工单状态 running complete fail
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 备注
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// IP地址列表
	IpCode []*string `json:"IpCode,omitempty" xml:"IpCode,omitempty" type:"Repeated"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 创建人
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// 资源一级ID
	IpRecordId *string `json:"IpRecordId,omitempty" xml:"IpRecordId,omitempty"`
	// 资源名称
	IpRecordName *string `json:"IpRecordName,omitempty" xml:"IpRecordName,omitempty"`
	// 园区名
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	// 业务类型名称
	BusinessTypeName *string `json:"BusinessTypeName,omitempty" xml:"BusinessTypeName,omitempty"`
	// 工单类型 Apply 申请工单 Recycle 释放工单
	RecodeType *string `json:"RecodeType,omitempty" xml:"RecodeType,omitempty"`
	// 地址段
	IpBlock *string `json:"IpBlock,omitempty" xml:"IpBlock,omitempty"`
	// 工单详情
	Detail []*GetIpRecordResponseBodyIpRecordDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
}

func (s GetIpRecordResponseBodyIpRecord) String() string {
	return tea.Prettify(s)
}

func (s GetIpRecordResponseBodyIpRecord) GoString() string {
	return s.String()
}

func (s *GetIpRecordResponseBodyIpRecord) SetStatus(v string) *GetIpRecordResponseBodyIpRecord {
	s.Status = &v
	return s
}

func (s *GetIpRecordResponseBodyIpRecord) SetDescription(v string) *GetIpRecordResponseBodyIpRecord {
	s.Description = &v
	return s
}

func (s *GetIpRecordResponseBodyIpRecord) SetIpCode(v []*string) *GetIpRecordResponseBodyIpRecord {
	s.IpCode = v
	return s
}

func (s *GetIpRecordResponseBodyIpRecord) SetCreateTime(v string) *GetIpRecordResponseBodyIpRecord {
	s.CreateTime = &v
	return s
}

func (s *GetIpRecordResponseBodyIpRecord) SetUpdateTime(v string) *GetIpRecordResponseBodyIpRecord {
	s.UpdateTime = &v
	return s
}

func (s *GetIpRecordResponseBodyIpRecord) SetCreator(v string) *GetIpRecordResponseBodyIpRecord {
	s.Creator = &v
	return s
}

func (s *GetIpRecordResponseBodyIpRecord) SetIpRecordId(v string) *GetIpRecordResponseBodyIpRecord {
	s.IpRecordId = &v
	return s
}

func (s *GetIpRecordResponseBodyIpRecord) SetIpRecordName(v string) *GetIpRecordResponseBodyIpRecord {
	s.IpRecordName = &v
	return s
}

func (s *GetIpRecordResponseBodyIpRecord) SetZoneName(v string) *GetIpRecordResponseBodyIpRecord {
	s.ZoneName = &v
	return s
}

func (s *GetIpRecordResponseBodyIpRecord) SetBusinessTypeName(v string) *GetIpRecordResponseBodyIpRecord {
	s.BusinessTypeName = &v
	return s
}

func (s *GetIpRecordResponseBodyIpRecord) SetRecodeType(v string) *GetIpRecordResponseBodyIpRecord {
	s.RecodeType = &v
	return s
}

func (s *GetIpRecordResponseBodyIpRecord) SetIpBlock(v string) *GetIpRecordResponseBodyIpRecord {
	s.IpBlock = &v
	return s
}

func (s *GetIpRecordResponseBodyIpRecord) SetDetail(v []*GetIpRecordResponseBodyIpRecordDetail) *GetIpRecordResponseBodyIpRecord {
	s.Detail = v
	return s
}

type GetIpRecordResponseBodyIpRecordDetail struct {
	// 申请到的Ip
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// 设备端口
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// 对端IP
	RemoteIp *string `json:"RemoteIp,omitempty" xml:"RemoteIp,omitempty"`
	// 对端设备名称
	RemoteDeviceName *string `json:"RemoteDeviceName,omitempty" xml:"RemoteDeviceName,omitempty"`
	// 对端设备端口
	RemotePort *string `json:"RemotePort,omitempty" xml:"RemotePort,omitempty"`
	// 设备MAC
	DeviceMac *string `json:"DeviceMac,omitempty" xml:"DeviceMac,omitempty"`
	// 网关
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// 父地址段列表
	ParentIpBlocks []*string `json:"ParentIpBlocks,omitempty" xml:"ParentIpBlocks,omitempty" type:"Repeated"`
	// 园区层级
	ZoneLayer []*GetIpRecordResponseBodyIpRecordDetailZoneLayer `json:"ZoneLayer,omitempty" xml:"ZoneLayer,omitempty" type:"Repeated"`
}

func (s GetIpRecordResponseBodyIpRecordDetail) String() string {
	return tea.Prettify(s)
}

func (s GetIpRecordResponseBodyIpRecordDetail) GoString() string {
	return s.String()
}

func (s *GetIpRecordResponseBodyIpRecordDetail) SetIp(v string) *GetIpRecordResponseBodyIpRecordDetail {
	s.Ip = &v
	return s
}

func (s *GetIpRecordResponseBodyIpRecordDetail) SetDeviceName(v string) *GetIpRecordResponseBodyIpRecordDetail {
	s.DeviceName = &v
	return s
}

func (s *GetIpRecordResponseBodyIpRecordDetail) SetPort(v string) *GetIpRecordResponseBodyIpRecordDetail {
	s.Port = &v
	return s
}

func (s *GetIpRecordResponseBodyIpRecordDetail) SetRemoteIp(v string) *GetIpRecordResponseBodyIpRecordDetail {
	s.RemoteIp = &v
	return s
}

func (s *GetIpRecordResponseBodyIpRecordDetail) SetRemoteDeviceName(v string) *GetIpRecordResponseBodyIpRecordDetail {
	s.RemoteDeviceName = &v
	return s
}

func (s *GetIpRecordResponseBodyIpRecordDetail) SetRemotePort(v string) *GetIpRecordResponseBodyIpRecordDetail {
	s.RemotePort = &v
	return s
}

func (s *GetIpRecordResponseBodyIpRecordDetail) SetDeviceMac(v string) *GetIpRecordResponseBodyIpRecordDetail {
	s.DeviceMac = &v
	return s
}

func (s *GetIpRecordResponseBodyIpRecordDetail) SetGateway(v string) *GetIpRecordResponseBodyIpRecordDetail {
	s.Gateway = &v
	return s
}

func (s *GetIpRecordResponseBodyIpRecordDetail) SetParentIpBlocks(v []*string) *GetIpRecordResponseBodyIpRecordDetail {
	s.ParentIpBlocks = v
	return s
}

func (s *GetIpRecordResponseBodyIpRecordDetail) SetZoneLayer(v []*GetIpRecordResponseBodyIpRecordDetailZoneLayer) *GetIpRecordResponseBodyIpRecordDetail {
	s.ZoneLayer = v
	return s
}

type GetIpRecordResponseBodyIpRecordDetailZoneLayer struct {
	// 园区层级名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 园区层级值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetIpRecordResponseBodyIpRecordDetailZoneLayer) String() string {
	return tea.Prettify(s)
}

func (s GetIpRecordResponseBodyIpRecordDetailZoneLayer) GoString() string {
	return s.String()
}

func (s *GetIpRecordResponseBodyIpRecordDetailZoneLayer) SetName(v string) *GetIpRecordResponseBodyIpRecordDetailZoneLayer {
	s.Name = &v
	return s
}

func (s *GetIpRecordResponseBodyIpRecordDetailZoneLayer) SetValue(v string) *GetIpRecordResponseBodyIpRecordDetailZoneLayer {
	s.Value = &v
	return s
}

type GetIpRecordResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetIpRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIpRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIpRecordResponse) GoString() string {
	return s.String()
}

func (s *GetIpRecordResponse) SetHeaders(v map[string]*string) *GetIpRecordResponse {
	s.Headers = v
	return s
}

func (s *GetIpRecordResponse) SetBody(v *GetIpRecordResponseBody) *GetIpRecordResponse {
	s.Body = v
	return s
}

type ListScheduleDutiesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s ListScheduleDutiesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleDutiesRequest) GoString() string {
	return s.String()
}

func (s *ListScheduleDutiesRequest) SetInstanceId(v string) *ListScheduleDutiesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScheduleDutiesRequest) SetMaxResults(v int32) *ListScheduleDutiesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListScheduleDutiesRequest) SetNextToken(v string) *ListScheduleDutiesRequest {
	s.NextToken = &v
	return s
}

func (s *ListScheduleDutiesRequest) SetStartDate(v string) *ListScheduleDutiesRequest {
	s.StartDate = &v
	return s
}

func (s *ListScheduleDutiesRequest) SetEndDate(v string) *ListScheduleDutiesRequest {
	s.EndDate = &v
	return s
}

type ListScheduleDutiesResponseBody struct {
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 数组，返回示例目录。
	ScheduleDuties []*ListScheduleDutiesResponseBodyScheduleDuties `json:"ScheduleDuties,omitempty" xml:"ScheduleDuties,omitempty" type:"Repeated"`
}

func (s ListScheduleDutiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleDutiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduleDutiesResponseBody) SetTotalCount(v int32) *ListScheduleDutiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListScheduleDutiesResponseBody) SetRequestId(v string) *ListScheduleDutiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduleDutiesResponseBody) SetNextToken(v int32) *ListScheduleDutiesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListScheduleDutiesResponseBody) SetScheduleDuties(v []*ListScheduleDutiesResponseBodyScheduleDuties) *ListScheduleDutiesResponseBody {
	s.ScheduleDuties = v
	return s
}

type ListScheduleDutiesResponseBodyScheduleDuties struct {
	// 值班表日期
	ScheduleDutyDate *string `json:"ScheduleDutyDate,omitempty" xml:"ScheduleDutyDate,omitempty"`
	// 资源一级ID
	ScheduleDutyId *string `json:"ScheduleDutyId,omitempty" xml:"ScheduleDutyId,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 是否节假日
	Holiday        *bool                                                         `json:"Holiday,omitempty" xml:"Holiday,omitempty"`
	TypeWorkerList []*ListScheduleDutiesResponseBodyScheduleDutiesTypeWorkerList `json:"TypeWorkerList,omitempty" xml:"TypeWorkerList,omitempty" type:"Repeated"`
}

func (s ListScheduleDutiesResponseBodyScheduleDuties) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleDutiesResponseBodyScheduleDuties) GoString() string {
	return s.String()
}

func (s *ListScheduleDutiesResponseBodyScheduleDuties) SetScheduleDutyDate(v string) *ListScheduleDutiesResponseBodyScheduleDuties {
	s.ScheduleDutyDate = &v
	return s
}

func (s *ListScheduleDutiesResponseBodyScheduleDuties) SetScheduleDutyId(v string) *ListScheduleDutiesResponseBodyScheduleDuties {
	s.ScheduleDutyId = &v
	return s
}

func (s *ListScheduleDutiesResponseBodyScheduleDuties) SetCreateTime(v string) *ListScheduleDutiesResponseBodyScheduleDuties {
	s.CreateTime = &v
	return s
}

func (s *ListScheduleDutiesResponseBodyScheduleDuties) SetUpdateTime(v string) *ListScheduleDutiesResponseBodyScheduleDuties {
	s.UpdateTime = &v
	return s
}

func (s *ListScheduleDutiesResponseBodyScheduleDuties) SetHoliday(v bool) *ListScheduleDutiesResponseBodyScheduleDuties {
	s.Holiday = &v
	return s
}

func (s *ListScheduleDutiesResponseBodyScheduleDuties) SetTypeWorkerList(v []*ListScheduleDutiesResponseBodyScheduleDutiesTypeWorkerList) *ListScheduleDutiesResponseBodyScheduleDuties {
	s.TypeWorkerList = v
	return s
}

type ListScheduleDutiesResponseBodyScheduleDutiesTypeWorkerList struct {
	// 值班表id
	ScheduleDutyId *string `json:"ScheduleDutyId,omitempty" xml:"ScheduleDutyId,omitempty"`
	// 值班类型Id
	ScheduleTypeId *string `json:"ScheduleTypeId,omitempty" xml:"ScheduleTypeId,omitempty"`
	// 值班类型名称
	ScheduleTypeName *string `json:"ScheduleTypeName,omitempty" xml:"ScheduleTypeName,omitempty"`
	// 值班人员名称
	ScheduleWorkerName []*string `json:"ScheduleWorkerName,omitempty" xml:"ScheduleWorkerName,omitempty" type:"Repeated"`
}

func (s ListScheduleDutiesResponseBodyScheduleDutiesTypeWorkerList) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleDutiesResponseBodyScheduleDutiesTypeWorkerList) GoString() string {
	return s.String()
}

func (s *ListScheduleDutiesResponseBodyScheduleDutiesTypeWorkerList) SetScheduleDutyId(v string) *ListScheduleDutiesResponseBodyScheduleDutiesTypeWorkerList {
	s.ScheduleDutyId = &v
	return s
}

func (s *ListScheduleDutiesResponseBodyScheduleDutiesTypeWorkerList) SetScheduleTypeId(v string) *ListScheduleDutiesResponseBodyScheduleDutiesTypeWorkerList {
	s.ScheduleTypeId = &v
	return s
}

func (s *ListScheduleDutiesResponseBodyScheduleDutiesTypeWorkerList) SetScheduleTypeName(v string) *ListScheduleDutiesResponseBodyScheduleDutiesTypeWorkerList {
	s.ScheduleTypeName = &v
	return s
}

func (s *ListScheduleDutiesResponseBodyScheduleDutiesTypeWorkerList) SetScheduleWorkerName(v []*string) *ListScheduleDutiesResponseBodyScheduleDutiesTypeWorkerList {
	s.ScheduleWorkerName = v
	return s
}

type ListScheduleDutiesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListScheduleDutiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListScheduleDutiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScheduleDutiesResponse) GoString() string {
	return s.String()
}

func (s *ListScheduleDutiesResponse) SetHeaders(v map[string]*string) *ListScheduleDutiesResponse {
	s.Headers = v
	return s
}

func (s *ListScheduleDutiesResponse) SetBody(v *ListScheduleDutiesResponseBody) *ListScheduleDutiesResponse {
	s.Body = v
	return s
}

type LockSpaceModelRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源id
	SpaceModelId *string `json:"SpaceModelId,omitempty" xml:"SpaceModelId,omitempty"`
}

func (s LockSpaceModelRequest) String() string {
	return tea.Prettify(s)
}

func (s LockSpaceModelRequest) GoString() string {
	return s.String()
}

func (s *LockSpaceModelRequest) SetInstanceId(v string) *LockSpaceModelRequest {
	s.InstanceId = &v
	return s
}

func (s *LockSpaceModelRequest) SetSpaceModelId(v string) *LockSpaceModelRequest {
	s.SpaceModelId = &v
	return s
}

type LockSpaceModelResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LockSpaceModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LockSpaceModelResponseBody) GoString() string {
	return s.String()
}

func (s *LockSpaceModelResponseBody) SetRequestId(v string) *LockSpaceModelResponseBody {
	s.RequestId = &v
	return s
}

type LockSpaceModelResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LockSpaceModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LockSpaceModelResponse) String() string {
	return tea.Prettify(s)
}

func (s LockSpaceModelResponse) GoString() string {
	return s.String()
}

func (s *LockSpaceModelResponse) SetHeaders(v map[string]*string) *LockSpaceModelResponse {
	s.Headers = v
	return s
}

func (s *LockSpaceModelResponse) SetBody(v *LockSpaceModelResponseBody) *LockSpaceModelResponse {
	s.Body = v
	return s
}

type UpdateResourceInformationRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源信息Id
	ResourceInformationId *string `json:"ResourceInformationId,omitempty" xml:"ResourceInformationId,omitempty"`
	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 资源属性
	ResourceAttribute *string `json:"ResourceAttribute,omitempty" xml:"ResourceAttribute,omitempty"`
	// 信息
	Information []*UpdateResourceInformationRequestInformation `json:"Information,omitempty" xml:"Information,omitempty" type:"Repeated"`
}

func (s UpdateResourceInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceInformationRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceInformationRequest) SetInstanceId(v string) *UpdateResourceInformationRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateResourceInformationRequest) SetResourceInformationId(v string) *UpdateResourceInformationRequest {
	s.ResourceInformationId = &v
	return s
}

func (s *UpdateResourceInformationRequest) SetResourceType(v string) *UpdateResourceInformationRequest {
	s.ResourceType = &v
	return s
}

func (s *UpdateResourceInformationRequest) SetResourceAttribute(v string) *UpdateResourceInformationRequest {
	s.ResourceAttribute = &v
	return s
}

func (s *UpdateResourceInformationRequest) SetInformation(v []*UpdateResourceInformationRequestInformation) *UpdateResourceInformationRequest {
	s.Information = v
	return s
}

type UpdateResourceInformationRequestInformation struct {
	// 键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 键属性
	KeyAttribute *string `json:"KeyAttribute,omitempty" xml:"KeyAttribute,omitempty"`
	// 键动作
	KeyAction *string `json:"KeyAction,omitempty" xml:"KeyAction,omitempty"`
	// 键描述
	KeyDescription *string `json:"KeyDescription,omitempty" xml:"KeyDescription,omitempty"`
	// 建设项目id
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
}

func (s UpdateResourceInformationRequestInformation) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceInformationRequestInformation) GoString() string {
	return s.String()
}

func (s *UpdateResourceInformationRequestInformation) SetKey(v string) *UpdateResourceInformationRequestInformation {
	s.Key = &v
	return s
}

func (s *UpdateResourceInformationRequestInformation) SetKeyAttribute(v string) *UpdateResourceInformationRequestInformation {
	s.KeyAttribute = &v
	return s
}

func (s *UpdateResourceInformationRequestInformation) SetKeyAction(v string) *UpdateResourceInformationRequestInformation {
	s.KeyAction = &v
	return s
}

func (s *UpdateResourceInformationRequestInformation) SetKeyDescription(v string) *UpdateResourceInformationRequestInformation {
	s.KeyDescription = &v
	return s
}

func (s *UpdateResourceInformationRequestInformation) SetSetupProjectId(v string) *UpdateResourceInformationRequestInformation {
	s.SetupProjectId = &v
	return s
}

type UpdateResourceInformationShrinkRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源信息Id
	ResourceInformationId *string `json:"ResourceInformationId,omitempty" xml:"ResourceInformationId,omitempty"`
	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 资源属性
	ResourceAttribute *string `json:"ResourceAttribute,omitempty" xml:"ResourceAttribute,omitempty"`
	// 信息
	InformationShrink *string `json:"Information,omitempty" xml:"Information,omitempty"`
}

func (s UpdateResourceInformationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceInformationShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceInformationShrinkRequest) SetInstanceId(v string) *UpdateResourceInformationShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateResourceInformationShrinkRequest) SetResourceInformationId(v string) *UpdateResourceInformationShrinkRequest {
	s.ResourceInformationId = &v
	return s
}

func (s *UpdateResourceInformationShrinkRequest) SetResourceType(v string) *UpdateResourceInformationShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *UpdateResourceInformationShrinkRequest) SetResourceAttribute(v string) *UpdateResourceInformationShrinkRequest {
	s.ResourceAttribute = &v
	return s
}

func (s *UpdateResourceInformationShrinkRequest) SetInformationShrink(v string) *UpdateResourceInformationShrinkRequest {
	s.InformationShrink = &v
	return s
}

type UpdateResourceInformationResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourceInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceInformationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceInformationResponseBody) SetRequestId(v string) *UpdateResourceInformationResponseBody {
	s.RequestId = &v
	return s
}

type UpdateResourceInformationResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateResourceInformationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateResourceInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceInformationResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceInformationResponse) SetHeaders(v map[string]*string) *UpdateResourceInformationResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceInformationResponse) SetBody(v *UpdateResourceInformationResponseBody) *UpdateResourceInformationResponse {
	s.Body = v
	return s
}

type CreateTimePeriodRequest struct {
	// 时间段名称
	TimePeriodName *string `json:"TimePeriodName,omitempty" xml:"TimePeriodName,omitempty"`
	// 描述
	TimePeriodDescription *string `json:"TimePeriodDescription,omitempty" xml:"TimePeriodDescription,omitempty"`
	// cron表达式
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateTimePeriodRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTimePeriodRequest) GoString() string {
	return s.String()
}

func (s *CreateTimePeriodRequest) SetTimePeriodName(v string) *CreateTimePeriodRequest {
	s.TimePeriodName = &v
	return s
}

func (s *CreateTimePeriodRequest) SetTimePeriodDescription(v string) *CreateTimePeriodRequest {
	s.TimePeriodDescription = &v
	return s
}

func (s *CreateTimePeriodRequest) SetExpression(v string) *CreateTimePeriodRequest {
	s.Expression = &v
	return s
}

func (s *CreateTimePeriodRequest) SetClientToken(v string) *CreateTimePeriodRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTimePeriodRequest) SetInstanceId(v string) *CreateTimePeriodRequest {
	s.InstanceId = &v
	return s
}

type CreateTimePeriodResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 时间段ID
	TimePeriodId *string `json:"TimePeriodId,omitempty" xml:"TimePeriodId,omitempty"`
}

func (s CreateTimePeriodResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTimePeriodResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTimePeriodResponseBody) SetRequestId(v string) *CreateTimePeriodResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTimePeriodResponseBody) SetTimePeriodId(v string) *CreateTimePeriodResponseBody {
	s.TimePeriodId = &v
	return s
}

type CreateTimePeriodResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTimePeriodResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTimePeriodResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTimePeriodResponse) GoString() string {
	return s.String()
}

func (s *CreateTimePeriodResponse) SetHeaders(v map[string]*string) *CreateTimePeriodResponse {
	s.Headers = v
	return s
}

func (s *CreateTimePeriodResponse) SetBody(v *CreateTimePeriodResponseBody) *CreateTimePeriodResponse {
	s.Body = v
	return s
}

type ListDevicesRequest struct {
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 设备形态ID
	DeviceFormId *string `json:"DeviceFormId,omitempty" xml:"DeviceFormId,omitempty"`
	// 设备形态名称
	DeviceFormName *string `json:"DeviceFormName,omitempty" xml:"DeviceFormName,omitempty"`
	// 物理空间ID
	PhysicalSpaceId *string `json:"PhysicalSpaceId,omitempty" xml:"PhysicalSpaceId,omitempty"`
	// 设备主机名
	HostName []*string `json:"HostName,omitempty" xml:"HostName,omitempty" type:"Repeated"`
	// 设备IP
	Ip []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
	// 设备SN
	Sn []*string `json:"Sn,omitempty" xml:"Sn,omitempty" type:"Repeated"`
	// 设备MAC
	Mac []*string `json:"Mac,omitempty" xml:"Mac,omitempty" type:"Repeated"`
	// 设备厂商
	Vendor []*string `json:"Vendor,omitempty" xml:"Vendor,omitempty" type:"Repeated"`
	// 设备型号
	Model []*string `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	// 设备服务状态
	ServiceStatus []*string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty" type:"Repeated"`
	// 安全域
	SecurityDomain []*string `json:"SecurityDomain,omitempty" xml:"SecurityDomain,omitempty" type:"Repeated"`
	// 设备额外属性
	ExtAttributes *string `json:"ExtAttributes,omitempty" xml:"ExtAttributes,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListDevicesRequest) SetMaxResults(v int32) *ListDevicesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDevicesRequest) SetNextToken(v string) *ListDevicesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDevicesRequest) SetDeviceFormId(v string) *ListDevicesRequest {
	s.DeviceFormId = &v
	return s
}

func (s *ListDevicesRequest) SetDeviceFormName(v string) *ListDevicesRequest {
	s.DeviceFormName = &v
	return s
}

func (s *ListDevicesRequest) SetPhysicalSpaceId(v string) *ListDevicesRequest {
	s.PhysicalSpaceId = &v
	return s
}

func (s *ListDevicesRequest) SetHostName(v []*string) *ListDevicesRequest {
	s.HostName = v
	return s
}

func (s *ListDevicesRequest) SetIp(v []*string) *ListDevicesRequest {
	s.Ip = v
	return s
}

func (s *ListDevicesRequest) SetSn(v []*string) *ListDevicesRequest {
	s.Sn = v
	return s
}

func (s *ListDevicesRequest) SetMac(v []*string) *ListDevicesRequest {
	s.Mac = v
	return s
}

func (s *ListDevicesRequest) SetVendor(v []*string) *ListDevicesRequest {
	s.Vendor = v
	return s
}

func (s *ListDevicesRequest) SetModel(v []*string) *ListDevicesRequest {
	s.Model = v
	return s
}

func (s *ListDevicesRequest) SetServiceStatus(v []*string) *ListDevicesRequest {
	s.ServiceStatus = v
	return s
}

func (s *ListDevicesRequest) SetSecurityDomain(v []*string) *ListDevicesRequest {
	s.SecurityDomain = v
	return s
}

func (s *ListDevicesRequest) SetExtAttributes(v string) *ListDevicesRequest {
	s.ExtAttributes = &v
	return s
}

func (s *ListDevicesRequest) SetInstanceId(v string) *ListDevicesRequest {
	s.InstanceId = &v
	return s
}

type ListDevicesResponseBody struct {
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 数组，返回示例目录。
	Devices []*ListDevicesResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// 每页数量。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBody) SetTotalCount(v int32) *ListDevicesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDevicesResponseBody) SetRequestId(v string) *ListDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDevicesResponseBody) SetNextToken(v int32) *ListDevicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDevicesResponseBody) SetDevices(v []*ListDevicesResponseBodyDevices) *ListDevicesResponseBody {
	s.Devices = v
	return s
}

func (s *ListDevicesResponseBody) SetMaxResults(v int32) *ListDevicesResponseBody {
	s.MaxResults = &v
	return s
}

type ListDevicesResponseBodyDevices struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 设备形态名称
	DeviceFormName *string `json:"DeviceFormName,omitempty" xml:"DeviceFormName,omitempty"`
	// 设备形态ID
	DeviceFormId *string `json:"DeviceFormId,omitempty" xml:"DeviceFormId,omitempty"`
	// 物理空间ID
	PhysicalSpaceId *string `json:"PhysicalSpaceId,omitempty" xml:"PhysicalSpaceId,omitempty"`
	// 物理空间名称
	PhysicalSpaceName *string `json:"PhysicalSpaceName,omitempty" xml:"PhysicalSpaceName,omitempty"`
	// 主机名
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// 设备IP
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// 设备SN
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// 设备MAC地址
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// 设备厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 设备型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 设备安全域
	SecurityDomain *string `json:"SecurityDomain,omitempty" xml:"SecurityDomain,omitempty"`
	// 设备状态
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// 登录类型
	LoginType *string `json:"LoginType,omitempty" xml:"LoginType,omitempty"`
	// 登录账号
	LoginUsername *string `json:"LoginUsername,omitempty" xml:"LoginUsername,omitempty"`
	// 登录密码
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// enable密码
	EnablePassword *string `json:"EnablePassword,omitempty" xml:"EnablePassword,omitempty"`
	// SNMP版本号
	SnmpAccountVersion *string `json:"SnmpAccountVersion,omitempty" xml:"SnmpAccountVersion,omitempty"`
	// SNMP Community
	SnmpCommunity *string `json:"SnmpCommunity,omitempty" xml:"SnmpCommunity,omitempty"`
	// SNMP 账号类型
	SnmpAccountType *string `json:"SnmpAccountType,omitempty" xml:"SnmpAccountType,omitempty"`
	// SNMP 安全级别
	SnmpSecurityLevel *string `json:"SnmpSecurityLevel,omitempty" xml:"SnmpSecurityLevel,omitempty"`
	// SNMP 用户名
	SnmpUsername *string `json:"SnmpUsername,omitempty" xml:"SnmpUsername,omitempty"`
	// SNMP Auth PassPhrase
	SnmpAuthPassPhrase *string `json:"SnmpAuthPassPhrase,omitempty" xml:"SnmpAuthPassPhrase,omitempty"`
	// SNMP Auth Protocol
	SnmpAuthProtocol *string `json:"SnmpAuthProtocol,omitempty" xml:"SnmpAuthProtocol,omitempty"`
	// SNMP Privacy Passphase
	SnmpPrivacyPassphase *string `json:"SnmpPrivacyPassphase,omitempty" xml:"SnmpPrivacyPassphase,omitempty"`
	// SNMP Privacy Protocol
	SnmpPrivacyProtocol *string `json:"SnmpPrivacyProtocol,omitempty" xml:"SnmpPrivacyProtocol,omitempty"`
	// 设备额外属性
	ExtAttributes *string `json:"ExtAttributes,omitempty" xml:"ExtAttributes,omitempty"`
}

func (s ListDevicesResponseBodyDevices) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyDevices) SetDeviceId(v string) *ListDevicesResponseBodyDevices {
	s.DeviceId = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetDeviceFormName(v string) *ListDevicesResponseBodyDevices {
	s.DeviceFormName = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetDeviceFormId(v string) *ListDevicesResponseBodyDevices {
	s.DeviceFormId = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetPhysicalSpaceId(v string) *ListDevicesResponseBodyDevices {
	s.PhysicalSpaceId = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetPhysicalSpaceName(v string) *ListDevicesResponseBodyDevices {
	s.PhysicalSpaceName = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetHostName(v string) *ListDevicesResponseBodyDevices {
	s.HostName = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetIp(v string) *ListDevicesResponseBodyDevices {
	s.Ip = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetSn(v string) *ListDevicesResponseBodyDevices {
	s.Sn = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetMac(v string) *ListDevicesResponseBodyDevices {
	s.Mac = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetVendor(v string) *ListDevicesResponseBodyDevices {
	s.Vendor = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetModel(v string) *ListDevicesResponseBodyDevices {
	s.Model = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetSecurityDomain(v string) *ListDevicesResponseBodyDevices {
	s.SecurityDomain = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetServiceStatus(v string) *ListDevicesResponseBodyDevices {
	s.ServiceStatus = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetLoginType(v string) *ListDevicesResponseBodyDevices {
	s.LoginType = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetLoginUsername(v string) *ListDevicesResponseBodyDevices {
	s.LoginUsername = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetLoginPassword(v string) *ListDevicesResponseBodyDevices {
	s.LoginPassword = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetEnablePassword(v string) *ListDevicesResponseBodyDevices {
	s.EnablePassword = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetSnmpAccountVersion(v string) *ListDevicesResponseBodyDevices {
	s.SnmpAccountVersion = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetSnmpCommunity(v string) *ListDevicesResponseBodyDevices {
	s.SnmpCommunity = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetSnmpAccountType(v string) *ListDevicesResponseBodyDevices {
	s.SnmpAccountType = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetSnmpSecurityLevel(v string) *ListDevicesResponseBodyDevices {
	s.SnmpSecurityLevel = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetSnmpUsername(v string) *ListDevicesResponseBodyDevices {
	s.SnmpUsername = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetSnmpAuthPassPhrase(v string) *ListDevicesResponseBodyDevices {
	s.SnmpAuthPassPhrase = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetSnmpAuthProtocol(v string) *ListDevicesResponseBodyDevices {
	s.SnmpAuthProtocol = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetSnmpPrivacyPassphase(v string) *ListDevicesResponseBodyDevices {
	s.SnmpPrivacyPassphase = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetSnmpPrivacyProtocol(v string) *ListDevicesResponseBodyDevices {
	s.SnmpPrivacyProtocol = &v
	return s
}

func (s *ListDevicesResponseBodyDevices) SetExtAttributes(v string) *ListDevicesResponseBodyDevices {
	s.ExtAttributes = &v
	return s
}

type ListDevicesResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponse) GoString() string {
	return s.String()
}

func (s *ListDevicesResponse) SetHeaders(v map[string]*string) *ListDevicesResponse {
	s.Headers = v
	return s
}

func (s *ListDevicesResponse) SetBody(v *ListDevicesResponseBody) *ListDevicesResponse {
	s.Body = v
	return s
}

type ListInspectionTaskReportsRequest struct {
	// 巡检项id
	InspectionItemId *string `json:"InspectionItemId,omitempty" xml:"InspectionItemId,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListInspectionTaskReportsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInspectionTaskReportsRequest) GoString() string {
	return s.String()
}

func (s *ListInspectionTaskReportsRequest) SetInspectionItemId(v string) *ListInspectionTaskReportsRequest {
	s.InspectionItemId = &v
	return s
}

func (s *ListInspectionTaskReportsRequest) SetVendor(v string) *ListInspectionTaskReportsRequest {
	s.Vendor = &v
	return s
}

func (s *ListInspectionTaskReportsRequest) SetInstanceId(v string) *ListInspectionTaskReportsRequest {
	s.InstanceId = &v
	return s
}

type ListInspectionTaskReportsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 内容
	InspectionTask []*ListInspectionTaskReportsResponseBodyInspectionTask `json:"InspectionTask,omitempty" xml:"InspectionTask,omitempty" type:"Repeated"`
}

func (s ListInspectionTaskReportsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInspectionTaskReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInspectionTaskReportsResponseBody) SetRequestId(v string) *ListInspectionTaskReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBody) SetInspectionTask(v []*ListInspectionTaskReportsResponseBodyInspectionTask) *ListInspectionTaskReportsResponseBody {
	s.InspectionTask = v
	return s
}

type ListInspectionTaskReportsResponseBodyInspectionTask struct {
	// 总设备数
	TotalDeviceNumber *int64 `json:"TotalDeviceNumber,omitempty" xml:"TotalDeviceNumber,omitempty"`
	// 实际设备数
	ActualDeviceNumber *int64 `json:"ActualDeviceNumber,omitempty" xml:"ActualDeviceNumber,omitempty"`
	// 设备占比
	DeviceRate *string `json:"DeviceRate,omitempty" xml:"DeviceRate,omitempty"`
	// 任务数
	TaskNumber *int64 `json:"TaskNumber,omitempty" xml:"TaskNumber,omitempty"`
	// 正在运行任务数
	RunningTaskNumber *int64 `json:"RunningTaskNumber,omitempty" xml:"RunningTaskNumber,omitempty"`
	// 运行占比
	RunningTaskRate *string `json:"RunningTaskRate,omitempty" xml:"RunningTaskRate,omitempty"`
	// 成功任务数
	SuccessTaskNumber *int64 `json:"SuccessTaskNumber,omitempty" xml:"SuccessTaskNumber,omitempty"`
	// 成功占比
	SuccessTaskRate *string `json:"SuccessTaskRate,omitempty" xml:"SuccessTaskRate,omitempty"`
	// 失败任务数
	FailureTaskNumber *int64 `json:"FailureTaskNumber,omitempty" xml:"FailureTaskNumber,omitempty"`
	// 失败占比
	FailureTaskRate *string `json:"FailureTaskRate,omitempty" xml:"FailureTaskRate,omitempty"`
	// 失败详情
	FailureStatistic *string `json:"FailureStatistic,omitempty" xml:"FailureStatistic,omitempty"`
	// 告警数量
	AlarmNumber *int64 `json:"AlarmNumber,omitempty" xml:"AlarmNumber,omitempty"`
	// 高危数量
	CriticalNumber *int64 `json:"CriticalNumber,omitempty" xml:"CriticalNumber,omitempty"`
	// 高危占比
	CriticalRate *string `json:"CriticalRate,omitempty" xml:"CriticalRate,omitempty"`
	// 中危数量
	WarningNumber *int64 `json:"WarningNumber,omitempty" xml:"WarningNumber,omitempty"`
	// 中危占比
	WarningRate *string `json:"WarningRate,omitempty" xml:"WarningRate,omitempty"`
	// 正常数量
	NormalNumber *int64 `json:"NormalNumber,omitempty" xml:"NormalNumber,omitempty"`
	// 正常占比
	NormalRate     *string                                                              `json:"NormalRate,omitempty" xml:"NormalRate,omitempty"`
	AlarmStatistic []*ListInspectionTaskReportsResponseBodyInspectionTaskAlarmStatistic `json:"AlarmStatistic,omitempty" xml:"AlarmStatistic,omitempty" type:"Repeated"`
}

func (s ListInspectionTaskReportsResponseBodyInspectionTask) String() string {
	return tea.Prettify(s)
}

func (s ListInspectionTaskReportsResponseBodyInspectionTask) GoString() string {
	return s.String()
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTask) SetTotalDeviceNumber(v int64) *ListInspectionTaskReportsResponseBodyInspectionTask {
	s.TotalDeviceNumber = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTask) SetActualDeviceNumber(v int64) *ListInspectionTaskReportsResponseBodyInspectionTask {
	s.ActualDeviceNumber = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTask) SetDeviceRate(v string) *ListInspectionTaskReportsResponseBodyInspectionTask {
	s.DeviceRate = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTask) SetTaskNumber(v int64) *ListInspectionTaskReportsResponseBodyInspectionTask {
	s.TaskNumber = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTask) SetRunningTaskNumber(v int64) *ListInspectionTaskReportsResponseBodyInspectionTask {
	s.RunningTaskNumber = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTask) SetRunningTaskRate(v string) *ListInspectionTaskReportsResponseBodyInspectionTask {
	s.RunningTaskRate = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTask) SetSuccessTaskNumber(v int64) *ListInspectionTaskReportsResponseBodyInspectionTask {
	s.SuccessTaskNumber = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTask) SetSuccessTaskRate(v string) *ListInspectionTaskReportsResponseBodyInspectionTask {
	s.SuccessTaskRate = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTask) SetFailureTaskNumber(v int64) *ListInspectionTaskReportsResponseBodyInspectionTask {
	s.FailureTaskNumber = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTask) SetFailureTaskRate(v string) *ListInspectionTaskReportsResponseBodyInspectionTask {
	s.FailureTaskRate = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTask) SetFailureStatistic(v string) *ListInspectionTaskReportsResponseBodyInspectionTask {
	s.FailureStatistic = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTask) SetAlarmNumber(v int64) *ListInspectionTaskReportsResponseBodyInspectionTask {
	s.AlarmNumber = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTask) SetCriticalNumber(v int64) *ListInspectionTaskReportsResponseBodyInspectionTask {
	s.CriticalNumber = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTask) SetCriticalRate(v string) *ListInspectionTaskReportsResponseBodyInspectionTask {
	s.CriticalRate = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTask) SetWarningNumber(v int64) *ListInspectionTaskReportsResponseBodyInspectionTask {
	s.WarningNumber = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTask) SetWarningRate(v string) *ListInspectionTaskReportsResponseBodyInspectionTask {
	s.WarningRate = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTask) SetNormalNumber(v int64) *ListInspectionTaskReportsResponseBodyInspectionTask {
	s.NormalNumber = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTask) SetNormalRate(v string) *ListInspectionTaskReportsResponseBodyInspectionTask {
	s.NormalRate = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTask) SetAlarmStatistic(v []*ListInspectionTaskReportsResponseBodyInspectionTaskAlarmStatistic) *ListInspectionTaskReportsResponseBodyInspectionTask {
	s.AlarmStatistic = v
	return s
}

type ListInspectionTaskReportsResponseBodyInspectionTaskAlarmStatistic struct {
	// 巡检项
	InspectionItem *string `json:"InspectionItem,omitempty" xml:"InspectionItem,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 高危数量
	CriticalNumber *int64 `json:"CriticalNumber,omitempty" xml:"CriticalNumber,omitempty"`
	// 中危数量
	WarningNumber *int64 `json:"WarningNumber,omitempty" xml:"WarningNumber,omitempty"`
}

func (s ListInspectionTaskReportsResponseBodyInspectionTaskAlarmStatistic) String() string {
	return tea.Prettify(s)
}

func (s ListInspectionTaskReportsResponseBodyInspectionTaskAlarmStatistic) GoString() string {
	return s.String()
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTaskAlarmStatistic) SetInspectionItem(v string) *ListInspectionTaskReportsResponseBodyInspectionTaskAlarmStatistic {
	s.InspectionItem = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTaskAlarmStatistic) SetVendor(v string) *ListInspectionTaskReportsResponseBodyInspectionTaskAlarmStatistic {
	s.Vendor = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTaskAlarmStatistic) SetModel(v string) *ListInspectionTaskReportsResponseBodyInspectionTaskAlarmStatistic {
	s.Model = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTaskAlarmStatistic) SetCriticalNumber(v int64) *ListInspectionTaskReportsResponseBodyInspectionTaskAlarmStatistic {
	s.CriticalNumber = &v
	return s
}

func (s *ListInspectionTaskReportsResponseBodyInspectionTaskAlarmStatistic) SetWarningNumber(v int64) *ListInspectionTaskReportsResponseBodyInspectionTaskAlarmStatistic {
	s.WarningNumber = &v
	return s
}

type ListInspectionTaskReportsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInspectionTaskReportsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInspectionTaskReportsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInspectionTaskReportsResponse) GoString() string {
	return s.String()
}

func (s *ListInspectionTaskReportsResponse) SetHeaders(v map[string]*string) *ListInspectionTaskReportsResponse {
	s.Headers = v
	return s
}

func (s *ListInspectionTaskReportsResponse) SetBody(v *ListInspectionTaskReportsResponseBody) *ListInspectionTaskReportsResponse {
	s.Body = v
	return s
}

type CreateMonitorItemRequest struct {
	// 监控项名称
	MonitorItemName *string `json:"MonitorItemName,omitempty" xml:"MonitorItemName,omitempty"`
	// 监控项描述
	MonitorItemDescription *string `json:"MonitorItemDescription,omitempty" xml:"MonitorItemDescription,omitempty"`
	// 数据项
	DataItem *string `json:"DataItem,omitempty" xml:"DataItem,omitempty"`
	// 安全域
	SecurityDomain *string `json:"SecurityDomain,omitempty" xml:"SecurityDomain,omitempty"`
	// 解析代码
	AnalysisCode *string `json:"AnalysisCode,omitempty" xml:"AnalysisCode,omitempty"`
	// 采集类型
	CollectionType *string `json:"CollectionType,omitempty" xml:"CollectionType,omitempty"`
	// 是否启用
	Effective *int32 `json:"Effective,omitempty" xml:"Effective,omitempty"`
	// 监控项参数配置
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// 执行间隔(s)
	ExecInterval *int32 `json:"ExecInterval,omitempty" xml:"ExecInterval,omitempty"`
	// 设备形态
	DeviceForm *string `json:"DeviceForm,omitempty" xml:"DeviceForm,omitempty"`
	// 告警规则列表
	AlarmRuleList []*CreateMonitorItemRequestAlarmRuleList `json:"AlarmRuleList,omitempty" xml:"AlarmRuleList,omitempty" type:"Repeated"`
	// 类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateMonitorItemRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorItemRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorItemRequest) SetMonitorItemName(v string) *CreateMonitorItemRequest {
	s.MonitorItemName = &v
	return s
}

func (s *CreateMonitorItemRequest) SetMonitorItemDescription(v string) *CreateMonitorItemRequest {
	s.MonitorItemDescription = &v
	return s
}

func (s *CreateMonitorItemRequest) SetDataItem(v string) *CreateMonitorItemRequest {
	s.DataItem = &v
	return s
}

func (s *CreateMonitorItemRequest) SetSecurityDomain(v string) *CreateMonitorItemRequest {
	s.SecurityDomain = &v
	return s
}

func (s *CreateMonitorItemRequest) SetAnalysisCode(v string) *CreateMonitorItemRequest {
	s.AnalysisCode = &v
	return s
}

func (s *CreateMonitorItemRequest) SetCollectionType(v string) *CreateMonitorItemRequest {
	s.CollectionType = &v
	return s
}

func (s *CreateMonitorItemRequest) SetEffective(v int32) *CreateMonitorItemRequest {
	s.Effective = &v
	return s
}

func (s *CreateMonitorItemRequest) SetConfig(v string) *CreateMonitorItemRequest {
	s.Config = &v
	return s
}

func (s *CreateMonitorItemRequest) SetExecInterval(v int32) *CreateMonitorItemRequest {
	s.ExecInterval = &v
	return s
}

func (s *CreateMonitorItemRequest) SetDeviceForm(v string) *CreateMonitorItemRequest {
	s.DeviceForm = &v
	return s
}

func (s *CreateMonitorItemRequest) SetAlarmRuleList(v []*CreateMonitorItemRequestAlarmRuleList) *CreateMonitorItemRequest {
	s.AlarmRuleList = v
	return s
}

func (s *CreateMonitorItemRequest) SetType(v string) *CreateMonitorItemRequest {
	s.Type = &v
	return s
}

func (s *CreateMonitorItemRequest) SetClientToken(v string) *CreateMonitorItemRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateMonitorItemRequest) SetInstanceId(v string) *CreateMonitorItemRequest {
	s.InstanceId = &v
	return s
}

type CreateMonitorItemRequestAlarmRuleList struct {
	// 告警状态
	AlarmStatus *string `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	// 指标名
	Variable *string `json:"Variable,omitempty" xml:"Variable,omitempty"`
	// 表达式
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// 比较值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateMonitorItemRequestAlarmRuleList) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorItemRequestAlarmRuleList) GoString() string {
	return s.String()
}

func (s *CreateMonitorItemRequestAlarmRuleList) SetAlarmStatus(v string) *CreateMonitorItemRequestAlarmRuleList {
	s.AlarmStatus = &v
	return s
}

func (s *CreateMonitorItemRequestAlarmRuleList) SetVariable(v string) *CreateMonitorItemRequestAlarmRuleList {
	s.Variable = &v
	return s
}

func (s *CreateMonitorItemRequestAlarmRuleList) SetExpression(v string) *CreateMonitorItemRequestAlarmRuleList {
	s.Expression = &v
	return s
}

func (s *CreateMonitorItemRequestAlarmRuleList) SetValue(v string) *CreateMonitorItemRequestAlarmRuleList {
	s.Value = &v
	return s
}

type CreateMonitorItemResponseBody struct {
	// 监控项ID
	MonitorItemId *string `json:"MonitorItemId,omitempty" xml:"MonitorItemId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMonitorItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorItemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMonitorItemResponseBody) SetMonitorItemId(v string) *CreateMonitorItemResponseBody {
	s.MonitorItemId = &v
	return s
}

func (s *CreateMonitorItemResponseBody) SetRequestId(v string) *CreateMonitorItemResponseBody {
	s.RequestId = &v
	return s
}

type CreateMonitorItemResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMonitorItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMonitorItemResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorItemResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorItemResponse) SetHeaders(v map[string]*string) *CreateMonitorItemResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorItemResponse) SetBody(v *CreateMonitorItemResponseBody) *CreateMonitorItemResponse {
	s.Body = v
	return s
}

type CreatePhysicalSpaceRequest struct {
	// 物理空间名称
	PhysicalSpaceName *string `json:"PhysicalSpaceName,omitempty" xml:"PhysicalSpaceName,omitempty"`
	// 所属国家
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// 所属省份
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// 所属城市
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// 具体地址
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// 幂等校验 token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 物理空间模型
	SpaceType *string `json:"SpaceType,omitempty" xml:"SpaceType,omitempty"`
	// 物理空间缩写
	SpaceAbbreviation *string `json:"SpaceAbbreviation,omitempty" xml:"SpaceAbbreviation,omitempty"`
	// 负责人
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
}

func (s CreatePhysicalSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePhysicalSpaceRequest) GoString() string {
	return s.String()
}

func (s *CreatePhysicalSpaceRequest) SetPhysicalSpaceName(v string) *CreatePhysicalSpaceRequest {
	s.PhysicalSpaceName = &v
	return s
}

func (s *CreatePhysicalSpaceRequest) SetCountry(v string) *CreatePhysicalSpaceRequest {
	s.Country = &v
	return s
}

func (s *CreatePhysicalSpaceRequest) SetProvince(v string) *CreatePhysicalSpaceRequest {
	s.Province = &v
	return s
}

func (s *CreatePhysicalSpaceRequest) SetCity(v string) *CreatePhysicalSpaceRequest {
	s.City = &v
	return s
}

func (s *CreatePhysicalSpaceRequest) SetAddress(v string) *CreatePhysicalSpaceRequest {
	s.Address = &v
	return s
}

func (s *CreatePhysicalSpaceRequest) SetClientToken(v string) *CreatePhysicalSpaceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePhysicalSpaceRequest) SetInstanceId(v string) *CreatePhysicalSpaceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreatePhysicalSpaceRequest) SetSpaceType(v string) *CreatePhysicalSpaceRequest {
	s.SpaceType = &v
	return s
}

func (s *CreatePhysicalSpaceRequest) SetSpaceAbbreviation(v string) *CreatePhysicalSpaceRequest {
	s.SpaceAbbreviation = &v
	return s
}

func (s *CreatePhysicalSpaceRequest) SetOwner(v string) *CreatePhysicalSpaceRequest {
	s.Owner = &v
	return s
}

type CreatePhysicalSpaceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 物理空间ID
	PhysicalSpaceId *string `json:"PhysicalSpaceId,omitempty" xml:"PhysicalSpaceId,omitempty"`
}

func (s CreatePhysicalSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePhysicalSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePhysicalSpaceResponseBody) SetRequestId(v string) *CreatePhysicalSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePhysicalSpaceResponseBody) SetPhysicalSpaceId(v string) *CreatePhysicalSpaceResponseBody {
	s.PhysicalSpaceId = &v
	return s
}

type CreatePhysicalSpaceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePhysicalSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePhysicalSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePhysicalSpaceResponse) GoString() string {
	return s.String()
}

func (s *CreatePhysicalSpaceResponse) SetHeaders(v map[string]*string) *CreatePhysicalSpaceResponse {
	s.Headers = v
	return s
}

func (s *CreatePhysicalSpaceResponse) SetBody(v *CreatePhysicalSpaceResponseBody) *CreatePhysicalSpaceResponse {
	s.Body = v
	return s
}

type UpdateDevicesRequest struct {
	// 设备ID
	DeviceIds []*string `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty" type:"Repeated"`
	// 登录类型
	LoginType *string `json:"LoginType,omitempty" xml:"LoginType,omitempty"`
	// 登录账号
	LoginUsername *string `json:"LoginUsername,omitempty" xml:"LoginUsername,omitempty"`
	// 登录密码
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// enable密码
	EnablePassword *string `json:"EnablePassword,omitempty" xml:"EnablePassword,omitempty"`
	// SNMP 版本号
	SnmpAccountVersion *string `json:"SnmpAccountVersion,omitempty" xml:"SnmpAccountVersion,omitempty"`
	// SNMP Community
	SnmpCommunity *string `json:"SnmpCommunity,omitempty" xml:"SnmpCommunity,omitempty"`
	// SNMP 账号类型
	SnmpAccountType *string `json:"SnmpAccountType,omitempty" xml:"SnmpAccountType,omitempty"`
	// SNMP 安全级别
	SnmpSecurityLevel *string `json:"SnmpSecurityLevel,omitempty" xml:"SnmpSecurityLevel,omitempty"`
	// SNMP 用户名
	SnmpUserName *string `json:"SnmpUserName,omitempty" xml:"SnmpUserName,omitempty"`
	// SNMP Auth PassPhrase
	SnmpAuthPassphrase *string `json:"SnmpAuthPassphrase,omitempty" xml:"SnmpAuthPassphrase,omitempty"`
	// SNMP Auth Protocol
	SnmpAuthProtocol *string `json:"SnmpAuthProtocol,omitempty" xml:"SnmpAuthProtocol,omitempty"`
	// SNMP Privacy Passphase
	SnmpPrivacyPassphase *string `json:"SnmpPrivacyPassphase,omitempty" xml:"SnmpPrivacyPassphase,omitempty"`
	// SNMP Privacy Protocol
	SnmpPrivacyProtocol *string `json:"SnmpPrivacyProtocol,omitempty" xml:"SnmpPrivacyProtocol,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 物理空间id
	PhysicalSpaceId *string `json:"PhysicalSpaceId,omitempty" xml:"PhysicalSpaceId,omitempty"`
	// 物理空间名称
	PhysicalSpaceName *string `json:"PhysicalSpaceName,omitempty" xml:"PhysicalSpaceName,omitempty"`
	// 服务状态
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s UpdateDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevicesRequest) GoString() string {
	return s.String()
}

func (s *UpdateDevicesRequest) SetDeviceIds(v []*string) *UpdateDevicesRequest {
	s.DeviceIds = v
	return s
}

func (s *UpdateDevicesRequest) SetLoginType(v string) *UpdateDevicesRequest {
	s.LoginType = &v
	return s
}

func (s *UpdateDevicesRequest) SetLoginUsername(v string) *UpdateDevicesRequest {
	s.LoginUsername = &v
	return s
}

func (s *UpdateDevicesRequest) SetLoginPassword(v string) *UpdateDevicesRequest {
	s.LoginPassword = &v
	return s
}

func (s *UpdateDevicesRequest) SetEnablePassword(v string) *UpdateDevicesRequest {
	s.EnablePassword = &v
	return s
}

func (s *UpdateDevicesRequest) SetSnmpAccountVersion(v string) *UpdateDevicesRequest {
	s.SnmpAccountVersion = &v
	return s
}

func (s *UpdateDevicesRequest) SetSnmpCommunity(v string) *UpdateDevicesRequest {
	s.SnmpCommunity = &v
	return s
}

func (s *UpdateDevicesRequest) SetSnmpAccountType(v string) *UpdateDevicesRequest {
	s.SnmpAccountType = &v
	return s
}

func (s *UpdateDevicesRequest) SetSnmpSecurityLevel(v string) *UpdateDevicesRequest {
	s.SnmpSecurityLevel = &v
	return s
}

func (s *UpdateDevicesRequest) SetSnmpUserName(v string) *UpdateDevicesRequest {
	s.SnmpUserName = &v
	return s
}

func (s *UpdateDevicesRequest) SetSnmpAuthPassphrase(v string) *UpdateDevicesRequest {
	s.SnmpAuthPassphrase = &v
	return s
}

func (s *UpdateDevicesRequest) SetSnmpAuthProtocol(v string) *UpdateDevicesRequest {
	s.SnmpAuthProtocol = &v
	return s
}

func (s *UpdateDevicesRequest) SetSnmpPrivacyPassphase(v string) *UpdateDevicesRequest {
	s.SnmpPrivacyPassphase = &v
	return s
}

func (s *UpdateDevicesRequest) SetSnmpPrivacyProtocol(v string) *UpdateDevicesRequest {
	s.SnmpPrivacyProtocol = &v
	return s
}

func (s *UpdateDevicesRequest) SetInstanceId(v string) *UpdateDevicesRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateDevicesRequest) SetPhysicalSpaceId(v string) *UpdateDevicesRequest {
	s.PhysicalSpaceId = &v
	return s
}

func (s *UpdateDevicesRequest) SetPhysicalSpaceName(v string) *UpdateDevicesRequest {
	s.PhysicalSpaceName = &v
	return s
}

func (s *UpdateDevicesRequest) SetServiceStatus(v string) *UpdateDevicesRequest {
	s.ServiceStatus = &v
	return s
}

func (s *UpdateDevicesRequest) SetVendor(v string) *UpdateDevicesRequest {
	s.Vendor = &v
	return s
}

func (s *UpdateDevicesRequest) SetModel(v string) *UpdateDevicesRequest {
	s.Model = &v
	return s
}

type UpdateDevicesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDevicesResponseBody) SetRequestId(v string) *UpdateDevicesResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDevicesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevicesResponse) GoString() string {
	return s.String()
}

func (s *UpdateDevicesResponse) SetHeaders(v map[string]*string) *UpdateDevicesResponse {
	s.Headers = v
	return s
}

func (s *UpdateDevicesResponse) SetBody(v *UpdateDevicesResponseBody) *UpdateDevicesResponse {
	s.Body = v
	return s
}

type UpdateScheduleTypeRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源一级ID
	ScheduleTypeId *string `json:"ScheduleTypeId,omitempty" xml:"ScheduleTypeId,omitempty"`
	// 值班类型key
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// relatedWorkerStr
	RelatedWorker []*string `json:"RelatedWorker,omitempty" xml:"RelatedWorker,omitempty" type:"Repeated"`
	// 值班类型状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateScheduleTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleTypeRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduleTypeRequest) SetInstanceId(v string) *UpdateScheduleTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateScheduleTypeRequest) SetScheduleTypeId(v string) *UpdateScheduleTypeRequest {
	s.ScheduleTypeId = &v
	return s
}

func (s *UpdateScheduleTypeRequest) SetScheduleType(v string) *UpdateScheduleTypeRequest {
	s.ScheduleType = &v
	return s
}

func (s *UpdateScheduleTypeRequest) SetRelatedWorker(v []*string) *UpdateScheduleTypeRequest {
	s.RelatedWorker = v
	return s
}

func (s *UpdateScheduleTypeRequest) SetStatus(v string) *UpdateScheduleTypeRequest {
	s.Status = &v
	return s
}

type UpdateScheduleTypeShrinkRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源一级ID
	ScheduleTypeId *string `json:"ScheduleTypeId,omitempty" xml:"ScheduleTypeId,omitempty"`
	// 值班类型key
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// relatedWorkerStr
	RelatedWorkerShrink *string `json:"RelatedWorker,omitempty" xml:"RelatedWorker,omitempty"`
	// 值班类型状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateScheduleTypeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleTypeShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduleTypeShrinkRequest) SetInstanceId(v string) *UpdateScheduleTypeShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateScheduleTypeShrinkRequest) SetScheduleTypeId(v string) *UpdateScheduleTypeShrinkRequest {
	s.ScheduleTypeId = &v
	return s
}

func (s *UpdateScheduleTypeShrinkRequest) SetScheduleType(v string) *UpdateScheduleTypeShrinkRequest {
	s.ScheduleType = &v
	return s
}

func (s *UpdateScheduleTypeShrinkRequest) SetRelatedWorkerShrink(v string) *UpdateScheduleTypeShrinkRequest {
	s.RelatedWorkerShrink = &v
	return s
}

func (s *UpdateScheduleTypeShrinkRequest) SetStatus(v string) *UpdateScheduleTypeShrinkRequest {
	s.Status = &v
	return s
}

type UpdateScheduleTypeResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateScheduleTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleTypeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScheduleTypeResponseBody) SetRequestId(v string) *UpdateScheduleTypeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateScheduleTypeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateScheduleTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateScheduleTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleTypeResponse) GoString() string {
	return s.String()
}

func (s *UpdateScheduleTypeResponse) SetHeaders(v map[string]*string) *UpdateScheduleTypeResponse {
	s.Headers = v
	return s
}

func (s *UpdateScheduleTypeResponse) SetBody(v *UpdateScheduleTypeResponseBody) *UpdateScheduleTypeResponse {
	s.Body = v
	return s
}

type DownloadDeviceResourceRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// deviceResourceId
	DeviceResourceId *string `json:"DeviceResourceId,omitempty" xml:"DeviceResourceId,omitempty"`
	// 资源uuid
	DeviceResourceIds []*string `json:"DeviceResourceIds,omitempty" xml:"DeviceResourceIds,omitempty" type:"Repeated"`
	// 操作类型
	DownloadType *string `json:"DownloadType,omitempty" xml:"DownloadType,omitempty"`
	// 建设项目id
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
}

func (s DownloadDeviceResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadDeviceResourceRequest) GoString() string {
	return s.String()
}

func (s *DownloadDeviceResourceRequest) SetInstanceId(v string) *DownloadDeviceResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DownloadDeviceResourceRequest) SetDeviceResourceId(v string) *DownloadDeviceResourceRequest {
	s.DeviceResourceId = &v
	return s
}

func (s *DownloadDeviceResourceRequest) SetDeviceResourceIds(v []*string) *DownloadDeviceResourceRequest {
	s.DeviceResourceIds = v
	return s
}

func (s *DownloadDeviceResourceRequest) SetDownloadType(v string) *DownloadDeviceResourceRequest {
	s.DownloadType = &v
	return s
}

func (s *DownloadDeviceResourceRequest) SetSetupProjectId(v string) *DownloadDeviceResourceRequest {
	s.SetupProjectId = &v
	return s
}

type DownloadDeviceResourceShrinkRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// deviceResourceId
	DeviceResourceId *string `json:"DeviceResourceId,omitempty" xml:"DeviceResourceId,omitempty"`
	// 资源uuid
	DeviceResourceIdsShrink *string `json:"DeviceResourceIds,omitempty" xml:"DeviceResourceIds,omitempty"`
	// 操作类型
	DownloadType *string `json:"DownloadType,omitempty" xml:"DownloadType,omitempty"`
	// 建设项目id
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
}

func (s DownloadDeviceResourceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadDeviceResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DownloadDeviceResourceShrinkRequest) SetInstanceId(v string) *DownloadDeviceResourceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DownloadDeviceResourceShrinkRequest) SetDeviceResourceId(v string) *DownloadDeviceResourceShrinkRequest {
	s.DeviceResourceId = &v
	return s
}

func (s *DownloadDeviceResourceShrinkRequest) SetDeviceResourceIdsShrink(v string) *DownloadDeviceResourceShrinkRequest {
	s.DeviceResourceIdsShrink = &v
	return s
}

func (s *DownloadDeviceResourceShrinkRequest) SetDownloadType(v string) *DownloadDeviceResourceShrinkRequest {
	s.DownloadType = &v
	return s
}

func (s *DownloadDeviceResourceShrinkRequest) SetSetupProjectId(v string) *DownloadDeviceResourceShrinkRequest {
	s.SetupProjectId = &v
	return s
}

type DownloadDeviceResourceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 下载链接
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
}

func (s DownloadDeviceResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadDeviceResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadDeviceResourceResponseBody) SetRequestId(v string) *DownloadDeviceResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadDeviceResourceResponseBody) SetDownloadUrl(v string) *DownloadDeviceResourceResponseBody {
	s.DownloadUrl = &v
	return s
}

type DownloadDeviceResourceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DownloadDeviceResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DownloadDeviceResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadDeviceResourceResponse) GoString() string {
	return s.String()
}

func (s *DownloadDeviceResourceResponse) SetHeaders(v map[string]*string) *DownloadDeviceResourceResponse {
	s.Headers = v
	return s
}

func (s *DownloadDeviceResourceResponse) SetBody(v *DownloadDeviceResourceResponseBody) *DownloadDeviceResourceResponse {
	s.Body = v
	return s
}

type GetOsDownloadPathRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 架构资源uuid
	OsVersionId *string `json:"OsVersionId,omitempty" xml:"OsVersionId,omitempty"`
}

func (s GetOsDownloadPathRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOsDownloadPathRequest) GoString() string {
	return s.String()
}

func (s *GetOsDownloadPathRequest) SetInstanceId(v string) *GetOsDownloadPathRequest {
	s.InstanceId = &v
	return s
}

func (s *GetOsDownloadPathRequest) SetOsVersionId(v string) *GetOsDownloadPathRequest {
	s.OsVersionId = &v
	return s
}

type GetOsDownloadPathResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 系统版本
	OsVersion *GetOsDownloadPathResponseBodyOsVersion `json:"OsVersion,omitempty" xml:"OsVersion,omitempty" type:"Struct"`
}

func (s GetOsDownloadPathResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOsDownloadPathResponseBody) GoString() string {
	return s.String()
}

func (s *GetOsDownloadPathResponseBody) SetRequestId(v string) *GetOsDownloadPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOsDownloadPathResponseBody) SetOsVersion(v *GetOsDownloadPathResponseBodyOsVersion) *GetOsDownloadPathResponseBody {
	s.OsVersion = v
	return s
}

type GetOsDownloadPathResponseBodyOsVersion struct {
	// 系统版本下载路径
	DownloadPath *string `json:"DownloadPath,omitempty" xml:"DownloadPath,omitempty"`
}

func (s GetOsDownloadPathResponseBodyOsVersion) String() string {
	return tea.Prettify(s)
}

func (s GetOsDownloadPathResponseBodyOsVersion) GoString() string {
	return s.String()
}

func (s *GetOsDownloadPathResponseBodyOsVersion) SetDownloadPath(v string) *GetOsDownloadPathResponseBodyOsVersion {
	s.DownloadPath = &v
	return s
}

type GetOsDownloadPathResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOsDownloadPathResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOsDownloadPathResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOsDownloadPathResponse) GoString() string {
	return s.String()
}

func (s *GetOsDownloadPathResponse) SetHeaders(v map[string]*string) *GetOsDownloadPathResponse {
	s.Headers = v
	return s
}

func (s *GetOsDownloadPathResponse) SetBody(v *GetOsDownloadPathResponseBody) *GetOsDownloadPathResponse {
	s.Body = v
	return s
}

type ListConnectionPoliciesRequest struct {
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 架构迭代uid
	ArchitectureIterationId *string `json:"ArchitectureIterationId,omitempty" xml:"ArchitectureIterationId,omitempty"`
	// 连接策略id
	ConnectionPolicyId *string `json:"ConnectionPolicyId,omitempty" xml:"ConnectionPolicyId,omitempty"`
	// 上联模块uid
	UplinkArchitectureModuleId *string `json:"UplinkArchitectureModuleId,omitempty" xml:"UplinkArchitectureModuleId,omitempty"`
	// 下联模块uid
	DownlinkArchitectureModuleId *string `json:"DownlinkArchitectureModuleId,omitempty" xml:"DownlinkArchitectureModuleId,omitempty"`
	// 上联设备uid
	UplinkArchitectureDeviceId *string `json:"UplinkArchitectureDeviceId,omitempty" xml:"UplinkArchitectureDeviceId,omitempty"`
	// 下联设备uid
	DownlinkArchitectureDeviceId *string `json:"DownlinkArchitectureDeviceId,omitempty" xml:"DownlinkArchitectureDeviceId,omitempty"`
}

func (s ListConnectionPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListConnectionPoliciesRequest) SetMaxResults(v int32) *ListConnectionPoliciesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListConnectionPoliciesRequest) SetNextToken(v string) *ListConnectionPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListConnectionPoliciesRequest) SetInstanceId(v string) *ListConnectionPoliciesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListConnectionPoliciesRequest) SetArchitectureIterationId(v string) *ListConnectionPoliciesRequest {
	s.ArchitectureIterationId = &v
	return s
}

func (s *ListConnectionPoliciesRequest) SetConnectionPolicyId(v string) *ListConnectionPoliciesRequest {
	s.ConnectionPolicyId = &v
	return s
}

func (s *ListConnectionPoliciesRequest) SetUplinkArchitectureModuleId(v string) *ListConnectionPoliciesRequest {
	s.UplinkArchitectureModuleId = &v
	return s
}

func (s *ListConnectionPoliciesRequest) SetDownlinkArchitectureModuleId(v string) *ListConnectionPoliciesRequest {
	s.DownlinkArchitectureModuleId = &v
	return s
}

func (s *ListConnectionPoliciesRequest) SetUplinkArchitectureDeviceId(v string) *ListConnectionPoliciesRequest {
	s.UplinkArchitectureDeviceId = &v
	return s
}

func (s *ListConnectionPoliciesRequest) SetDownlinkArchitectureDeviceId(v string) *ListConnectionPoliciesRequest {
	s.DownlinkArchitectureDeviceId = &v
	return s
}

type ListConnectionPoliciesResponseBody struct {
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 数组，返回示例目录。
	ConnectionPolicy []*ListConnectionPoliciesResponseBodyConnectionPolicy `json:"ConnectionPolicy,omitempty" xml:"ConnectionPolicy,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据量
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListConnectionPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectionPoliciesResponseBody) SetTotalCount(v int32) *ListConnectionPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListConnectionPoliciesResponseBody) SetConnectionPolicy(v []*ListConnectionPoliciesResponseBodyConnectionPolicy) *ListConnectionPoliciesResponseBody {
	s.ConnectionPolicy = v
	return s
}

func (s *ListConnectionPoliciesResponseBody) SetRequestId(v string) *ListConnectionPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConnectionPoliciesResponseBody) SetNextToken(v int32) *ListConnectionPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListConnectionPoliciesResponseBody) SetMaxResults(v int64) *ListConnectionPoliciesResponseBody {
	s.MaxResults = &v
	return s
}

type ListConnectionPoliciesResponseBodyConnectionPolicy struct {
	// 连接策略uid
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 架构迭代uid
	ArchitectureIterationId *string `json:"ArchitectureIterationId,omitempty" xml:"ArchitectureIterationId,omitempty"`
	// 连接策略名字
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 连接数
	LinkCount *int32 `json:"LinkCount,omitempty" xml:"LinkCount,omitempty"`
	// 连接策略算法
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 更新是啊金
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 下联模块uid
	UplinkArchitectureModuleId *string `json:"UplinkArchitectureModuleId,omitempty" xml:"UplinkArchitectureModuleId,omitempty"`
	// 下联设备uid
	DownlinkArchitectureModuleId *string `json:"DownlinkArchitectureModuleId,omitempty" xml:"DownlinkArchitectureModuleId,omitempty"`
	// 上联模块uid
	UplinkArchitectureDeviceId *string `json:"UplinkArchitectureDeviceId,omitempty" xml:"UplinkArchitectureDeviceId,omitempty"`
	// 上联设备uid
	DownlinkArchitectureDeviceId *string `json:"DownlinkArchitectureDeviceId,omitempty" xml:"DownlinkArchitectureDeviceId,omitempty"`
}

func (s ListConnectionPoliciesResponseBodyConnectionPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionPoliciesResponseBodyConnectionPolicy) GoString() string {
	return s.String()
}

func (s *ListConnectionPoliciesResponseBodyConnectionPolicy) SetId(v string) *ListConnectionPoliciesResponseBodyConnectionPolicy {
	s.Id = &v
	return s
}

func (s *ListConnectionPoliciesResponseBodyConnectionPolicy) SetArchitectureIterationId(v string) *ListConnectionPoliciesResponseBodyConnectionPolicy {
	s.ArchitectureIterationId = &v
	return s
}

func (s *ListConnectionPoliciesResponseBodyConnectionPolicy) SetName(v string) *ListConnectionPoliciesResponseBodyConnectionPolicy {
	s.Name = &v
	return s
}

func (s *ListConnectionPoliciesResponseBodyConnectionPolicy) SetLinkCount(v int32) *ListConnectionPoliciesResponseBodyConnectionPolicy {
	s.LinkCount = &v
	return s
}

func (s *ListConnectionPoliciesResponseBodyConnectionPolicy) SetAlgorithm(v string) *ListConnectionPoliciesResponseBodyConnectionPolicy {
	s.Algorithm = &v
	return s
}

func (s *ListConnectionPoliciesResponseBodyConnectionPolicy) SetCreateTime(v string) *ListConnectionPoliciesResponseBodyConnectionPolicy {
	s.CreateTime = &v
	return s
}

func (s *ListConnectionPoliciesResponseBodyConnectionPolicy) SetUpdateTime(v string) *ListConnectionPoliciesResponseBodyConnectionPolicy {
	s.UpdateTime = &v
	return s
}

func (s *ListConnectionPoliciesResponseBodyConnectionPolicy) SetUplinkArchitectureModuleId(v string) *ListConnectionPoliciesResponseBodyConnectionPolicy {
	s.UplinkArchitectureModuleId = &v
	return s
}

func (s *ListConnectionPoliciesResponseBodyConnectionPolicy) SetDownlinkArchitectureModuleId(v string) *ListConnectionPoliciesResponseBodyConnectionPolicy {
	s.DownlinkArchitectureModuleId = &v
	return s
}

func (s *ListConnectionPoliciesResponseBodyConnectionPolicy) SetUplinkArchitectureDeviceId(v string) *ListConnectionPoliciesResponseBodyConnectionPolicy {
	s.UplinkArchitectureDeviceId = &v
	return s
}

func (s *ListConnectionPoliciesResponseBodyConnectionPolicy) SetDownlinkArchitectureDeviceId(v string) *ListConnectionPoliciesResponseBodyConnectionPolicy {
	s.DownlinkArchitectureDeviceId = &v
	return s
}

type ListConnectionPoliciesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConnectionPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConnectionPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListConnectionPoliciesResponse) SetHeaders(v map[string]*string) *ListConnectionPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListConnectionPoliciesResponse) SetBody(v *ListConnectionPoliciesResponseBody) *ListConnectionPoliciesResponse {
	s.Body = v
	return s
}

type UpdateScheduleWorkerRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源一级ID
	ScheduleWorkerId *string `json:"ScheduleWorkerId,omitempty" xml:"ScheduleWorkerId,omitempty"`
	// 值班人员工号
	WorkerId *string `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
	// 值班人员姓名
	WorkerName *string `json:"WorkerName,omitempty" xml:"WorkerName,omitempty"`
	// 联系方式
	WorkerContact *string `json:"WorkerContact,omitempty" xml:"WorkerContact,omitempty"`
}

func (s UpdateScheduleWorkerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleWorkerRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduleWorkerRequest) SetInstanceId(v string) *UpdateScheduleWorkerRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateScheduleWorkerRequest) SetScheduleWorkerId(v string) *UpdateScheduleWorkerRequest {
	s.ScheduleWorkerId = &v
	return s
}

func (s *UpdateScheduleWorkerRequest) SetWorkerId(v string) *UpdateScheduleWorkerRequest {
	s.WorkerId = &v
	return s
}

func (s *UpdateScheduleWorkerRequest) SetWorkerName(v string) *UpdateScheduleWorkerRequest {
	s.WorkerName = &v
	return s
}

func (s *UpdateScheduleWorkerRequest) SetWorkerContact(v string) *UpdateScheduleWorkerRequest {
	s.WorkerContact = &v
	return s
}

type UpdateScheduleWorkerResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateScheduleWorkerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleWorkerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScheduleWorkerResponseBody) SetRequestId(v string) *UpdateScheduleWorkerResponseBody {
	s.RequestId = &v
	return s
}

type UpdateScheduleWorkerResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateScheduleWorkerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateScheduleWorkerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleWorkerResponse) GoString() string {
	return s.String()
}

func (s *UpdateScheduleWorkerResponse) SetHeaders(v map[string]*string) *UpdateScheduleWorkerResponse {
	s.Headers = v
	return s
}

func (s *UpdateScheduleWorkerResponse) SetBody(v *UpdateScheduleWorkerResponseBody) *UpdateScheduleWorkerResponse {
	s.Body = v
	return s
}

type DeleteConfigurationVariateRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例 ID。
	ConfigurationVariateId *string `json:"ConfigurationVariateId,omitempty" xml:"ConfigurationVariateId,omitempty"`
}

func (s DeleteConfigurationVariateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigurationVariateRequest) GoString() string {
	return s.String()
}

func (s *DeleteConfigurationVariateRequest) SetInstanceId(v string) *DeleteConfigurationVariateRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteConfigurationVariateRequest) SetConfigurationVariateId(v string) *DeleteConfigurationVariateRequest {
	s.ConfigurationVariateId = &v
	return s
}

type DeleteConfigurationVariateResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConfigurationVariateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigurationVariateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConfigurationVariateResponseBody) SetRequestId(v string) *DeleteConfigurationVariateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteConfigurationVariateResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteConfigurationVariateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteConfigurationVariateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigurationVariateResponse) GoString() string {
	return s.String()
}

func (s *DeleteConfigurationVariateResponse) SetHeaders(v map[string]*string) *DeleteConfigurationVariateResponse {
	s.Headers = v
	return s
}

func (s *DeleteConfigurationVariateResponse) SetBody(v *DeleteConfigurationVariateResponseBody) *DeleteConfigurationVariateResponse {
	s.Body = v
	return s
}

type CreateScheduleDutyRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始时间
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束时间
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 值班表类型
	ScheduleTypeIds []*string `json:"ScheduleTypeIds,omitempty" xml:"ScheduleTypeIds,omitempty" type:"Repeated"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateScheduleDutyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleDutyRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduleDutyRequest) SetInstanceId(v string) *CreateScheduleDutyRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateScheduleDutyRequest) SetStartDate(v string) *CreateScheduleDutyRequest {
	s.StartDate = &v
	return s
}

func (s *CreateScheduleDutyRequest) SetEndDate(v string) *CreateScheduleDutyRequest {
	s.EndDate = &v
	return s
}

func (s *CreateScheduleDutyRequest) SetScheduleTypeIds(v []*string) *CreateScheduleDutyRequest {
	s.ScheduleTypeIds = v
	return s
}

func (s *CreateScheduleDutyRequest) SetClientToken(v string) *CreateScheduleDutyRequest {
	s.ClientToken = &v
	return s
}

type CreateScheduleDutyShrinkRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 开始时间
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 结束时间
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 值班表类型
	ScheduleTypeIdsShrink *string `json:"ScheduleTypeIds,omitempty" xml:"ScheduleTypeIds,omitempty"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateScheduleDutyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleDutyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduleDutyShrinkRequest) SetInstanceId(v string) *CreateScheduleDutyShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateScheduleDutyShrinkRequest) SetStartDate(v string) *CreateScheduleDutyShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *CreateScheduleDutyShrinkRequest) SetEndDate(v string) *CreateScheduleDutyShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *CreateScheduleDutyShrinkRequest) SetScheduleTypeIdsShrink(v string) *CreateScheduleDutyShrinkRequest {
	s.ScheduleTypeIdsShrink = &v
	return s
}

func (s *CreateScheduleDutyShrinkRequest) SetClientToken(v string) *CreateScheduleDutyShrinkRequest {
	s.ClientToken = &v
	return s
}

type CreateScheduleDutyResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 资源实例ID，如ECS实例的创建接口CreateInstance应返回InstanceId。
	ScheduleDutyId *string `json:"ScheduleDutyId,omitempty" xml:"ScheduleDutyId,omitempty"`
}

func (s CreateScheduleDutyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleDutyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduleDutyResponseBody) SetRequestId(v string) *CreateScheduleDutyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduleDutyResponseBody) SetScheduleDutyId(v string) *CreateScheduleDutyResponseBody {
	s.ScheduleDutyId = &v
	return s
}

type CreateScheduleDutyResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateScheduleDutyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateScheduleDutyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleDutyResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduleDutyResponse) SetHeaders(v map[string]*string) *CreateScheduleDutyResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduleDutyResponse) SetBody(v *CreateScheduleDutyResponseBody) *CreateScheduleDutyResponse {
	s.Body = v
	return s
}

type GetPhysicalSpaceRequest struct {
	// 实例 ID。
	PhysicalSpaceId *string `json:"PhysicalSpaceId,omitempty" xml:"PhysicalSpaceId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetPhysicalSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalSpaceRequest) GoString() string {
	return s.String()
}

func (s *GetPhysicalSpaceRequest) SetPhysicalSpaceId(v string) *GetPhysicalSpaceRequest {
	s.PhysicalSpaceId = &v
	return s
}

func (s *GetPhysicalSpaceRequest) SetInstanceId(v string) *GetPhysicalSpaceRequest {
	s.InstanceId = &v
	return s
}

type GetPhysicalSpaceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 物理空间详情
	PhysicalSpace *GetPhysicalSpaceResponseBodyPhysicalSpace `json:"PhysicalSpace,omitempty" xml:"PhysicalSpace,omitempty" type:"Struct"`
}

func (s GetPhysicalSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhysicalSpaceResponseBody) SetRequestId(v string) *GetPhysicalSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhysicalSpaceResponseBody) SetPhysicalSpace(v *GetPhysicalSpaceResponseBodyPhysicalSpace) *GetPhysicalSpaceResponseBody {
	s.PhysicalSpace = v
	return s
}

type GetPhysicalSpaceResponseBodyPhysicalSpace struct {
	// 物理空间ID
	PhysicalSpaceId *string `json:"PhysicalSpaceId,omitempty" xml:"PhysicalSpaceId,omitempty"`
	// 物理空间名称
	PhysicalSpaceName *string `json:"PhysicalSpaceName,omitempty" xml:"PhysicalSpaceName,omitempty"`
	// 所属国家
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// 所属省份
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// 所属城市
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// 具体地址
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// 模型
	SpaceType *string `json:"SpaceType,omitempty" xml:"SpaceType,omitempty"`
	// 缩写
	SpaceAbbreviation *string `json:"SpaceAbbreviation,omitempty" xml:"SpaceAbbreviation,omitempty"`
	// 负责人
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
}

func (s GetPhysicalSpaceResponseBodyPhysicalSpace) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalSpaceResponseBodyPhysicalSpace) GoString() string {
	return s.String()
}

func (s *GetPhysicalSpaceResponseBodyPhysicalSpace) SetPhysicalSpaceId(v string) *GetPhysicalSpaceResponseBodyPhysicalSpace {
	s.PhysicalSpaceId = &v
	return s
}

func (s *GetPhysicalSpaceResponseBodyPhysicalSpace) SetPhysicalSpaceName(v string) *GetPhysicalSpaceResponseBodyPhysicalSpace {
	s.PhysicalSpaceName = &v
	return s
}

func (s *GetPhysicalSpaceResponseBodyPhysicalSpace) SetCountry(v string) *GetPhysicalSpaceResponseBodyPhysicalSpace {
	s.Country = &v
	return s
}

func (s *GetPhysicalSpaceResponseBodyPhysicalSpace) SetProvince(v string) *GetPhysicalSpaceResponseBodyPhysicalSpace {
	s.Province = &v
	return s
}

func (s *GetPhysicalSpaceResponseBodyPhysicalSpace) SetCity(v string) *GetPhysicalSpaceResponseBodyPhysicalSpace {
	s.City = &v
	return s
}

func (s *GetPhysicalSpaceResponseBodyPhysicalSpace) SetAddress(v string) *GetPhysicalSpaceResponseBodyPhysicalSpace {
	s.Address = &v
	return s
}

func (s *GetPhysicalSpaceResponseBodyPhysicalSpace) SetSpaceType(v string) *GetPhysicalSpaceResponseBodyPhysicalSpace {
	s.SpaceType = &v
	return s
}

func (s *GetPhysicalSpaceResponseBodyPhysicalSpace) SetSpaceAbbreviation(v string) *GetPhysicalSpaceResponseBodyPhysicalSpace {
	s.SpaceAbbreviation = &v
	return s
}

func (s *GetPhysicalSpaceResponseBodyPhysicalSpace) SetOwner(v string) *GetPhysicalSpaceResponseBodyPhysicalSpace {
	s.Owner = &v
	return s
}

type GetPhysicalSpaceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPhysicalSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPhysicalSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalSpaceResponse) GoString() string {
	return s.String()
}

func (s *GetPhysicalSpaceResponse) SetHeaders(v map[string]*string) *GetPhysicalSpaceResponse {
	s.Headers = v
	return s
}

func (s *GetPhysicalSpaceResponse) SetBody(v *GetPhysicalSpaceResponseBody) *GetPhysicalSpaceResponse {
	s.Body = v
	return s
}

type DeleteResourceInformationRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例 ID。
	ResourceInformationId *string `json:"ResourceInformationId,omitempty" xml:"ResourceInformationId,omitempty"`
}

func (s DeleteResourceInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceInformationRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceInformationRequest) SetInstanceId(v string) *DeleteResourceInformationRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteResourceInformationRequest) SetResourceInformationId(v string) *DeleteResourceInformationRequest {
	s.ResourceInformationId = &v
	return s
}

type DeleteResourceInformationResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceInformationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceInformationResponseBody) SetRequestId(v string) *DeleteResourceInformationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteResourceInformationResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteResourceInformationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteResourceInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceInformationResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceInformationResponse) SetHeaders(v map[string]*string) *DeleteResourceInformationResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceInformationResponse) SetBody(v *DeleteResourceInformationResponseBody) *DeleteResourceInformationResponse {
	s.Body = v
	return s
}

type DeleteSetupProjectRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例 ID。
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
}

func (s DeleteSetupProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSetupProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteSetupProjectRequest) SetInstanceId(v string) *DeleteSetupProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSetupProjectRequest) SetSetupProjectId(v string) *DeleteSetupProjectRequest {
	s.SetupProjectId = &v
	return s
}

type DeleteSetupProjectResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSetupProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSetupProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSetupProjectResponseBody) SetRequestId(v string) *DeleteSetupProjectResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSetupProjectResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSetupProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSetupProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSetupProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteSetupProjectResponse) SetHeaders(v map[string]*string) *DeleteSetupProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteSetupProjectResponse) SetBody(v *DeleteSetupProjectResponseBody) *DeleteSetupProjectResponse {
	s.Body = v
	return s
}

type CreateDedicatedLineRequest struct {
	// 物理空间ID
	PhysicalSpaceId *string `json:"PhysicalSpaceId,omitempty" xml:"PhysicalSpaceId,omitempty"`
	// 运营商
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// 宽带（Mbps）
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// 专线IP
	DedicatedLineIp *string `json:"DedicatedLineIp,omitempty" xml:"DedicatedLineIp,omitempty"`
	// 专线网关
	DedicatedLineGateway *string `json:"DedicatedLineGateway,omitempty" xml:"DedicatedLineGateway,omitempty"`
	// 专线角色
	DedicatedLineRole *string `json:"DedicatedLineRole,omitempty" xml:"DedicatedLineRole,omitempty"`
	// 关联设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 关联设备端口名称
	DevicePort *string `json:"DevicePort,omitempty" xml:"DevicePort,omitempty"`
	// 幂等校验 token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateDedicatedLineRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedLineRequest) GoString() string {
	return s.String()
}

func (s *CreateDedicatedLineRequest) SetPhysicalSpaceId(v string) *CreateDedicatedLineRequest {
	s.PhysicalSpaceId = &v
	return s
}

func (s *CreateDedicatedLineRequest) SetIsp(v string) *CreateDedicatedLineRequest {
	s.Isp = &v
	return s
}

func (s *CreateDedicatedLineRequest) SetBandwidth(v int32) *CreateDedicatedLineRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateDedicatedLineRequest) SetDedicatedLineIp(v string) *CreateDedicatedLineRequest {
	s.DedicatedLineIp = &v
	return s
}

func (s *CreateDedicatedLineRequest) SetDedicatedLineGateway(v string) *CreateDedicatedLineRequest {
	s.DedicatedLineGateway = &v
	return s
}

func (s *CreateDedicatedLineRequest) SetDedicatedLineRole(v string) *CreateDedicatedLineRequest {
	s.DedicatedLineRole = &v
	return s
}

func (s *CreateDedicatedLineRequest) SetDeviceId(v string) *CreateDedicatedLineRequest {
	s.DeviceId = &v
	return s
}

func (s *CreateDedicatedLineRequest) SetDevicePort(v string) *CreateDedicatedLineRequest {
	s.DevicePort = &v
	return s
}

func (s *CreateDedicatedLineRequest) SetClientToken(v string) *CreateDedicatedLineRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDedicatedLineRequest) SetInstanceId(v string) *CreateDedicatedLineRequest {
	s.InstanceId = &v
	return s
}

type CreateDedicatedLineResponseBody struct {
	// 资源实例ID，如ECS实例的创建接口CreateInstance应返回InstanceId。
	DedicatedLineId *string `json:"DedicatedLineId,omitempty" xml:"DedicatedLineId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDedicatedLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedLineResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDedicatedLineResponseBody) SetDedicatedLineId(v string) *CreateDedicatedLineResponseBody {
	s.DedicatedLineId = &v
	return s
}

func (s *CreateDedicatedLineResponseBody) SetRequestId(v string) *CreateDedicatedLineResponseBody {
	s.RequestId = &v
	return s
}

type CreateDedicatedLineResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDedicatedLineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDedicatedLineResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedLineResponse) GoString() string {
	return s.String()
}

func (s *CreateDedicatedLineResponse) SetHeaders(v map[string]*string) *CreateDedicatedLineResponse {
	s.Headers = v
	return s
}

func (s *CreateDedicatedLineResponse) SetBody(v *CreateDedicatedLineResponseBody) *CreateDedicatedLineResponse {
	s.Body = v
	return s
}

type ApplyIPRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// ip地址类型
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// 建设项目id
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// 业务类型id
	BusinessTypeId *string `json:"BusinessTypeId,omitempty" xml:"BusinessTypeId,omitempty"`
	// 设备uuid列表
	DeviceResourceIds []*string `json:"DeviceResourceIds,omitempty" xml:"DeviceResourceIds,omitempty" type:"Repeated"`
	// Loopback端口
	LoopbackPort *string `json:"LoopbackPort,omitempty" xml:"LoopbackPort,omitempty"`
	// 位置空间
	NetLocation *string `json:"NetLocation,omitempty" xml:"NetLocation,omitempty"`
	// 业务参数
	BusinessTypeParams *string `json:"BusinessTypeParams,omitempty" xml:"BusinessTypeParams,omitempty"`
	// deviceResourceId
	DeviceResourceId *string `json:"DeviceResourceId,omitempty" xml:"DeviceResourceId,omitempty"`
}

func (s ApplyIPRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyIPRequest) GoString() string {
	return s.String()
}

func (s *ApplyIPRequest) SetInstanceId(v string) *ApplyIPRequest {
	s.InstanceId = &v
	return s
}

func (s *ApplyIPRequest) SetIpType(v string) *ApplyIPRequest {
	s.IpType = &v
	return s
}

func (s *ApplyIPRequest) SetSetupProjectId(v string) *ApplyIPRequest {
	s.SetupProjectId = &v
	return s
}

func (s *ApplyIPRequest) SetBusinessTypeId(v string) *ApplyIPRequest {
	s.BusinessTypeId = &v
	return s
}

func (s *ApplyIPRequest) SetDeviceResourceIds(v []*string) *ApplyIPRequest {
	s.DeviceResourceIds = v
	return s
}

func (s *ApplyIPRequest) SetLoopbackPort(v string) *ApplyIPRequest {
	s.LoopbackPort = &v
	return s
}

func (s *ApplyIPRequest) SetNetLocation(v string) *ApplyIPRequest {
	s.NetLocation = &v
	return s
}

func (s *ApplyIPRequest) SetBusinessTypeParams(v string) *ApplyIPRequest {
	s.BusinessTypeParams = &v
	return s
}

func (s *ApplyIPRequest) SetDeviceResourceId(v string) *ApplyIPRequest {
	s.DeviceResourceId = &v
	return s
}

type ApplyIPShrinkRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// ip地址类型
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// 建设项目id
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// 业务类型id
	BusinessTypeId *string `json:"BusinessTypeId,omitempty" xml:"BusinessTypeId,omitempty"`
	// 设备uuid列表
	DeviceResourceIdsShrink *string `json:"DeviceResourceIds,omitempty" xml:"DeviceResourceIds,omitempty"`
	// Loopback端口
	LoopbackPort *string `json:"LoopbackPort,omitempty" xml:"LoopbackPort,omitempty"`
	// 位置空间
	NetLocation *string `json:"NetLocation,omitempty" xml:"NetLocation,omitempty"`
	// 业务参数
	BusinessTypeParams *string `json:"BusinessTypeParams,omitempty" xml:"BusinessTypeParams,omitempty"`
	// deviceResourceId
	DeviceResourceId *string `json:"DeviceResourceId,omitempty" xml:"DeviceResourceId,omitempty"`
}

func (s ApplyIPShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyIPShrinkRequest) GoString() string {
	return s.String()
}

func (s *ApplyIPShrinkRequest) SetInstanceId(v string) *ApplyIPShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ApplyIPShrinkRequest) SetIpType(v string) *ApplyIPShrinkRequest {
	s.IpType = &v
	return s
}

func (s *ApplyIPShrinkRequest) SetSetupProjectId(v string) *ApplyIPShrinkRequest {
	s.SetupProjectId = &v
	return s
}

func (s *ApplyIPShrinkRequest) SetBusinessTypeId(v string) *ApplyIPShrinkRequest {
	s.BusinessTypeId = &v
	return s
}

func (s *ApplyIPShrinkRequest) SetDeviceResourceIdsShrink(v string) *ApplyIPShrinkRequest {
	s.DeviceResourceIdsShrink = &v
	return s
}

func (s *ApplyIPShrinkRequest) SetLoopbackPort(v string) *ApplyIPShrinkRequest {
	s.LoopbackPort = &v
	return s
}

func (s *ApplyIPShrinkRequest) SetNetLocation(v string) *ApplyIPShrinkRequest {
	s.NetLocation = &v
	return s
}

func (s *ApplyIPShrinkRequest) SetBusinessTypeParams(v string) *ApplyIPShrinkRequest {
	s.BusinessTypeParams = &v
	return s
}

func (s *ApplyIPShrinkRequest) SetDeviceResourceId(v string) *ApplyIPShrinkRequest {
	s.DeviceResourceId = &v
	return s
}

type ApplyIPResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyIPResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyIPResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyIPResponseBody) SetRequestId(v string) *ApplyIPResponseBody {
	s.RequestId = &v
	return s
}

type ApplyIPResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApplyIPResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyIPResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyIPResponse) GoString() string {
	return s.String()
}

func (s *ApplyIPResponse) SetHeaders(v map[string]*string) *ApplyIPResponse {
	s.Headers = v
	return s
}

func (s *ApplyIPResponse) SetBody(v *ApplyIPResponseBody) *ApplyIPResponse {
	s.Body = v
	return s
}

type UpdateOsVersionRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// osVersionId
	OsVersionId *string `json:"OsVersionId,omitempty" xml:"OsVersionId,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 系统版本
	OsVersion *string `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 文件名
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 文件路径
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
}

func (s UpdateOsVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateOsVersionRequest) SetInstanceId(v string) *UpdateOsVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateOsVersionRequest) SetOsVersionId(v string) *UpdateOsVersionRequest {
	s.OsVersionId = &v
	return s
}

func (s *UpdateOsVersionRequest) SetVendor(v string) *UpdateOsVersionRequest {
	s.Vendor = &v
	return s
}

func (s *UpdateOsVersionRequest) SetModel(v string) *UpdateOsVersionRequest {
	s.Model = &v
	return s
}

func (s *UpdateOsVersionRequest) SetOsVersion(v string) *UpdateOsVersionRequest {
	s.OsVersion = &v
	return s
}

func (s *UpdateOsVersionRequest) SetStatus(v string) *UpdateOsVersionRequest {
	s.Status = &v
	return s
}

func (s *UpdateOsVersionRequest) SetFileName(v string) *UpdateOsVersionRequest {
	s.FileName = &v
	return s
}

func (s *UpdateOsVersionRequest) SetFilePath(v string) *UpdateOsVersionRequest {
	s.FilePath = &v
	return s
}

type UpdateOsVersionResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOsVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOsVersionResponseBody) SetRequestId(v string) *UpdateOsVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateOsVersionResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateOsVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateOsVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOsVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateOsVersionResponse) SetHeaders(v map[string]*string) *UpdateOsVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateOsVersionResponse) SetBody(v *UpdateOsVersionResponseBody) *UpdateOsVersionResponse {
	s.Body = v
	return s
}

type GetSpaceModelInstanceRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源id
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// 物理空间类型
	SpaceType *string `json:"SpaceType,omitempty" xml:"SpaceType,omitempty"`
	// 操作类型
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
}

func (s GetSpaceModelInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceModelInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetSpaceModelInstanceRequest) SetInstanceId(v string) *GetSpaceModelInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSpaceModelInstanceRequest) SetSpaceId(v string) *GetSpaceModelInstanceRequest {
	s.SpaceId = &v
	return s
}

func (s *GetSpaceModelInstanceRequest) SetSpaceType(v string) *GetSpaceModelInstanceRequest {
	s.SpaceType = &v
	return s
}

func (s *GetSpaceModelInstanceRequest) SetOperateType(v string) *GetSpaceModelInstanceRequest {
	s.OperateType = &v
	return s
}

type GetSpaceModelInstanceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 建设项目
	SpaceModel *GetSpaceModelInstanceResponseBodySpaceModel `json:"SpaceModel,omitempty" xml:"SpaceModel,omitempty" type:"Struct"`
}

func (s GetSpaceModelInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceModelInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpaceModelInstanceResponseBody) SetRequestId(v string) *GetSpaceModelInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSpaceModelInstanceResponseBody) SetSpaceModel(v *GetSpaceModelInstanceResponseBodySpaceModel) *GetSpaceModelInstanceResponseBody {
	s.SpaceModel = v
	return s
}

type GetSpaceModelInstanceResponseBodySpaceModel struct {
	// 物理空间实例
	Instance *string `json:"Instance,omitempty" xml:"Instance,omitempty"`
}

func (s GetSpaceModelInstanceResponseBodySpaceModel) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceModelInstanceResponseBodySpaceModel) GoString() string {
	return s.String()
}

func (s *GetSpaceModelInstanceResponseBodySpaceModel) SetInstance(v string) *GetSpaceModelInstanceResponseBodySpaceModel {
	s.Instance = &v
	return s
}

type GetSpaceModelInstanceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSpaceModelInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSpaceModelInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceModelInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetSpaceModelInstanceResponse) SetHeaders(v map[string]*string) *GetSpaceModelInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetSpaceModelInstanceResponse) SetBody(v *GetSpaceModelInstanceResponseBody) *GetSpaceModelInstanceResponse {
	s.Body = v
	return s
}

type ListOsVersionsRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListOsVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOsVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListOsVersionsRequest) SetInstanceId(v string) *ListOsVersionsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOsVersionsRequest) SetMaxResults(v int32) *ListOsVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOsVersionsRequest) SetNextToken(v string) *ListOsVersionsRequest {
	s.NextToken = &v
	return s
}

type ListOsVersionsResponseBody struct {
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 本次读取的最大数据量
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 数组，返回示例目录。
	OsVersion []*ListOsVersionsResponseBodyOsVersion `json:"OsVersion,omitempty" xml:"OsVersion,omitempty" type:"Repeated"`
}

func (s ListOsVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOsVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOsVersionsResponseBody) SetTotalCount(v int32) *ListOsVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOsVersionsResponseBody) SetMaxResults(v int64) *ListOsVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListOsVersionsResponseBody) SetRequestId(v string) *ListOsVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOsVersionsResponseBody) SetNextToken(v int32) *ListOsVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOsVersionsResponseBody) SetOsVersion(v []*ListOsVersionsResponseBodyOsVersion) *ListOsVersionsResponseBody {
	s.OsVersion = v
	return s
}

type ListOsVersionsResponseBodyOsVersion struct {
	// 系统版本
	OsVersion *string `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	// 文件名
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 文件路径
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 资源uuid
	OsVersionId *string `json:"OsVersionId,omitempty" xml:"OsVersionId,omitempty"`
	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListOsVersionsResponseBodyOsVersion) String() string {
	return tea.Prettify(s)
}

func (s ListOsVersionsResponseBodyOsVersion) GoString() string {
	return s.String()
}

func (s *ListOsVersionsResponseBodyOsVersion) SetOsVersion(v string) *ListOsVersionsResponseBodyOsVersion {
	s.OsVersion = &v
	return s
}

func (s *ListOsVersionsResponseBodyOsVersion) SetFileName(v string) *ListOsVersionsResponseBodyOsVersion {
	s.FileName = &v
	return s
}

func (s *ListOsVersionsResponseBodyOsVersion) SetStatus(v string) *ListOsVersionsResponseBodyOsVersion {
	s.Status = &v
	return s
}

func (s *ListOsVersionsResponseBodyOsVersion) SetFilePath(v string) *ListOsVersionsResponseBodyOsVersion {
	s.FilePath = &v
	return s
}

func (s *ListOsVersionsResponseBodyOsVersion) SetCreateTime(v string) *ListOsVersionsResponseBodyOsVersion {
	s.CreateTime = &v
	return s
}

func (s *ListOsVersionsResponseBodyOsVersion) SetVendor(v string) *ListOsVersionsResponseBodyOsVersion {
	s.Vendor = &v
	return s
}

func (s *ListOsVersionsResponseBodyOsVersion) SetModel(v string) *ListOsVersionsResponseBodyOsVersion {
	s.Model = &v
	return s
}

func (s *ListOsVersionsResponseBodyOsVersion) SetOsVersionId(v string) *ListOsVersionsResponseBodyOsVersion {
	s.OsVersionId = &v
	return s
}

func (s *ListOsVersionsResponseBodyOsVersion) SetUpdateTime(v string) *ListOsVersionsResponseBodyOsVersion {
	s.UpdateTime = &v
	return s
}

type ListOsVersionsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOsVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOsVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOsVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListOsVersionsResponse) SetHeaders(v map[string]*string) *ListOsVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListOsVersionsResponse) SetBody(v *ListOsVersionsResponseBody) *ListOsVersionsResponse {
	s.Body = v
	return s
}

type GetDeviceConfigRequest struct {
	// 实例 ID。
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 查询日期，格式 yyyy-MM-dd
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetDeviceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceConfigRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigRequest) SetDeviceId(v string) *GetDeviceConfigRequest {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceConfigRequest) SetDate(v string) *GetDeviceConfigRequest {
	s.Date = &v
	return s
}

func (s *GetDeviceConfigRequest) SetInstanceId(v string) *GetDeviceConfigRequest {
	s.InstanceId = &v
	return s
}

type GetDeviceConfigResponseBody struct {
	// 设备配置内容
	DeviceConfig *string `json:"DeviceConfig,omitempty" xml:"DeviceConfig,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigResponseBody) SetDeviceConfig(v string) *GetDeviceConfigResponseBody {
	s.DeviceConfig = &v
	return s
}

func (s *GetDeviceConfigResponseBody) SetRequestId(v string) *GetDeviceConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetDeviceConfigResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeviceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceConfigResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigResponse) SetHeaders(v map[string]*string) *GetDeviceConfigResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceConfigResponse) SetBody(v *GetDeviceConfigResponseBody) *GetDeviceConfigResponse {
	s.Body = v
	return s
}

type GetDedicatedLineRequest struct {
	// 实例 ID。
	DedicatedLineId *string `json:"DedicatedLineId,omitempty" xml:"DedicatedLineId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetDedicatedLineRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDedicatedLineRequest) GoString() string {
	return s.String()
}

func (s *GetDedicatedLineRequest) SetDedicatedLineId(v string) *GetDedicatedLineRequest {
	s.DedicatedLineId = &v
	return s
}

func (s *GetDedicatedLineRequest) SetInstanceId(v string) *GetDedicatedLineRequest {
	s.InstanceId = &v
	return s
}

type GetDedicatedLineResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 物理空间专线详情
	DedicatedLine *GetDedicatedLineResponseBodyDedicatedLine `json:"DedicatedLine,omitempty" xml:"DedicatedLine,omitempty" type:"Struct"`
}

func (s GetDedicatedLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDedicatedLineResponseBody) GoString() string {
	return s.String()
}

func (s *GetDedicatedLineResponseBody) SetRequestId(v string) *GetDedicatedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDedicatedLineResponseBody) SetDedicatedLine(v *GetDedicatedLineResponseBodyDedicatedLine) *GetDedicatedLineResponseBody {
	s.DedicatedLine = v
	return s
}

type GetDedicatedLineResponseBodyDedicatedLine struct {
	// 专线ID
	DedicatedLineId *string `json:"DedicatedLineId,omitempty" xml:"DedicatedLineId,omitempty"`
	// 运营商
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// 宽带（Mbps）
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// 专线IP
	DedicatedLineIp *string `json:"DedicatedLineIp,omitempty" xml:"DedicatedLineIp,omitempty"`
	// 专线网关
	DedicatedLineGateway *string `json:"DedicatedLineGateway,omitempty" xml:"DedicatedLineGateway,omitempty"`
	// 专线角色
	DedicatedLineRole *string `json:"DedicatedLineRole,omitempty" xml:"DedicatedLineRole,omitempty"`
	// 关联设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 关联设备端口名称
	DevicePort *string `json:"DevicePort,omitempty" xml:"DevicePort,omitempty"`
	// 关联设备名称
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// 物理空间ID
	PhysicalSpaceId *string `json:"PhysicalSpaceId,omitempty" xml:"PhysicalSpaceId,omitempty"`
}

func (s GetDedicatedLineResponseBodyDedicatedLine) String() string {
	return tea.Prettify(s)
}

func (s GetDedicatedLineResponseBodyDedicatedLine) GoString() string {
	return s.String()
}

func (s *GetDedicatedLineResponseBodyDedicatedLine) SetDedicatedLineId(v string) *GetDedicatedLineResponseBodyDedicatedLine {
	s.DedicatedLineId = &v
	return s
}

func (s *GetDedicatedLineResponseBodyDedicatedLine) SetIsp(v string) *GetDedicatedLineResponseBodyDedicatedLine {
	s.Isp = &v
	return s
}

func (s *GetDedicatedLineResponseBodyDedicatedLine) SetBandwidth(v int32) *GetDedicatedLineResponseBodyDedicatedLine {
	s.Bandwidth = &v
	return s
}

func (s *GetDedicatedLineResponseBodyDedicatedLine) SetDedicatedLineIp(v string) *GetDedicatedLineResponseBodyDedicatedLine {
	s.DedicatedLineIp = &v
	return s
}

func (s *GetDedicatedLineResponseBodyDedicatedLine) SetDedicatedLineGateway(v string) *GetDedicatedLineResponseBodyDedicatedLine {
	s.DedicatedLineGateway = &v
	return s
}

func (s *GetDedicatedLineResponseBodyDedicatedLine) SetDedicatedLineRole(v string) *GetDedicatedLineResponseBodyDedicatedLine {
	s.DedicatedLineRole = &v
	return s
}

func (s *GetDedicatedLineResponseBodyDedicatedLine) SetDeviceId(v string) *GetDedicatedLineResponseBodyDedicatedLine {
	s.DeviceId = &v
	return s
}

func (s *GetDedicatedLineResponseBodyDedicatedLine) SetDevicePort(v string) *GetDedicatedLineResponseBodyDedicatedLine {
	s.DevicePort = &v
	return s
}

func (s *GetDedicatedLineResponseBodyDedicatedLine) SetDeviceName(v string) *GetDedicatedLineResponseBodyDedicatedLine {
	s.DeviceName = &v
	return s
}

func (s *GetDedicatedLineResponseBodyDedicatedLine) SetPhysicalSpaceId(v string) *GetDedicatedLineResponseBodyDedicatedLine {
	s.PhysicalSpaceId = &v
	return s
}

type GetDedicatedLineResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDedicatedLineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDedicatedLineResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDedicatedLineResponse) GoString() string {
	return s.String()
}

func (s *GetDedicatedLineResponse) SetHeaders(v map[string]*string) *GetDedicatedLineResponse {
	s.Headers = v
	return s
}

func (s *GetDedicatedLineResponse) SetBody(v *GetDedicatedLineResponseBody) *GetDedicatedLineResponse {
	s.Body = v
	return s
}

type GetDeviceResourceRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例 ID。
	DeviceResourceId *string `json:"DeviceResourceId,omitempty" xml:"DeviceResourceId,omitempty"`
}

func (s GetDeviceResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceResourceRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceResourceRequest) SetInstanceId(v string) *GetDeviceResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDeviceResourceRequest) SetDeviceResourceId(v string) *GetDeviceResourceRequest {
	s.DeviceResourceId = &v
	return s
}

type GetDeviceResourceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 设备资源
	DeviceResource *GetDeviceResourceResponseBodyDeviceResource `json:"DeviceResource,omitempty" xml:"DeviceResource,omitempty" type:"Struct"`
}

func (s GetDeviceResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceResourceResponseBody) SetRequestId(v string) *GetDeviceResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceResourceResponseBody) SetDeviceResource(v *GetDeviceResourceResponseBodyDeviceResource) *GetDeviceResourceResponseBody {
	s.DeviceResource = v
	return s
}

type GetDeviceResourceResponseBodyDeviceResource struct {
	// 资源一级ID
	DeviceResourceId *string `json:"DeviceResourceId,omitempty" xml:"DeviceResourceId,omitempty"`
	// 建设项目id
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 组号
	BlockNumber *string `json:"BlockNumber,omitempty" xml:"BlockNumber,omitempty"`
	// 设备号
	DeviceNum *string `json:"DeviceNum,omitempty" xml:"DeviceNum,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 模型
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 主机名
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// 设备sn号
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// 物理空间位置
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// 带内管理地址
	ManagerIp *string `json:"ManagerIp,omitempty" xml:"ManagerIp,omitempty"`
	// 交付登录地址
	DeliveryIp *string `json:"DeliveryIp,omitempty" xml:"DeliveryIp,omitempty"`
	// 配置生成
	GenerateConfig *string `json:"GenerateConfig,omitempty" xml:"GenerateConfig,omitempty"`
	// 配置下发状态
	ConfigTaskStatus *string `json:"ConfigTaskStatus,omitempty" xml:"ConfigTaskStatus,omitempty"`
	// 设备配置
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// loopback地址
	Loopback *string `json:"Loopback,omitempty" xml:"Loopback,omitempty"`
	// 设备互联地址
	InterConnection *string `json:"InterConnection,omitempty" xml:"InterConnection,omitempty"`
	// 设备业务地址
	Business *string `json:"Business,omitempty" xml:"Business,omitempty"`
	// 堆叠状态
	Stack *bool `json:"Stack,omitempty" xml:"Stack,omitempty"`
}

func (s GetDeviceResourceResponseBodyDeviceResource) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceResourceResponseBodyDeviceResource) GoString() string {
	return s.String()
}

func (s *GetDeviceResourceResponseBodyDeviceResource) SetDeviceResourceId(v string) *GetDeviceResourceResponseBodyDeviceResource {
	s.DeviceResourceId = &v
	return s
}

func (s *GetDeviceResourceResponseBodyDeviceResource) SetSetupProjectId(v string) *GetDeviceResourceResponseBodyDeviceResource {
	s.SetupProjectId = &v
	return s
}

func (s *GetDeviceResourceResponseBodyDeviceResource) SetRole(v string) *GetDeviceResourceResponseBodyDeviceResource {
	s.Role = &v
	return s
}

func (s *GetDeviceResourceResponseBodyDeviceResource) SetBlockNumber(v string) *GetDeviceResourceResponseBodyDeviceResource {
	s.BlockNumber = &v
	return s
}

func (s *GetDeviceResourceResponseBodyDeviceResource) SetDeviceNum(v string) *GetDeviceResourceResponseBodyDeviceResource {
	s.DeviceNum = &v
	return s
}

func (s *GetDeviceResourceResponseBodyDeviceResource) SetVendor(v string) *GetDeviceResourceResponseBodyDeviceResource {
	s.Vendor = &v
	return s
}

func (s *GetDeviceResourceResponseBodyDeviceResource) SetModel(v string) *GetDeviceResourceResponseBodyDeviceResource {
	s.Model = &v
	return s
}

func (s *GetDeviceResourceResponseBodyDeviceResource) SetHostName(v string) *GetDeviceResourceResponseBodyDeviceResource {
	s.HostName = &v
	return s
}

func (s *GetDeviceResourceResponseBodyDeviceResource) SetSn(v string) *GetDeviceResourceResponseBodyDeviceResource {
	s.Sn = &v
	return s
}

func (s *GetDeviceResourceResponseBodyDeviceResource) SetLocation(v string) *GetDeviceResourceResponseBodyDeviceResource {
	s.Location = &v
	return s
}

func (s *GetDeviceResourceResponseBodyDeviceResource) SetManagerIp(v string) *GetDeviceResourceResponseBodyDeviceResource {
	s.ManagerIp = &v
	return s
}

func (s *GetDeviceResourceResponseBodyDeviceResource) SetDeliveryIp(v string) *GetDeviceResourceResponseBodyDeviceResource {
	s.DeliveryIp = &v
	return s
}

func (s *GetDeviceResourceResponseBodyDeviceResource) SetGenerateConfig(v string) *GetDeviceResourceResponseBodyDeviceResource {
	s.GenerateConfig = &v
	return s
}

func (s *GetDeviceResourceResponseBodyDeviceResource) SetConfigTaskStatus(v string) *GetDeviceResourceResponseBodyDeviceResource {
	s.ConfigTaskStatus = &v
	return s
}

func (s *GetDeviceResourceResponseBodyDeviceResource) SetConfig(v string) *GetDeviceResourceResponseBodyDeviceResource {
	s.Config = &v
	return s
}

func (s *GetDeviceResourceResponseBodyDeviceResource) SetLoopback(v string) *GetDeviceResourceResponseBodyDeviceResource {
	s.Loopback = &v
	return s
}

func (s *GetDeviceResourceResponseBodyDeviceResource) SetInterConnection(v string) *GetDeviceResourceResponseBodyDeviceResource {
	s.InterConnection = &v
	return s
}

func (s *GetDeviceResourceResponseBodyDeviceResource) SetBusiness(v string) *GetDeviceResourceResponseBodyDeviceResource {
	s.Business = &v
	return s
}

func (s *GetDeviceResourceResponseBodyDeviceResource) SetStack(v bool) *GetDeviceResourceResponseBodyDeviceResource {
	s.Stack = &v
	return s
}

type GetDeviceResourceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeviceResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceResourceResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceResourceResponse) SetHeaders(v map[string]*string) *GetDeviceResourceResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceResourceResponse) SetBody(v *GetDeviceResourceResponseBody) *GetDeviceResourceResponse {
	s.Body = v
	return s
}

type DeleteDedicatedLineRequest struct {
	// 实例 ID。
	DedicatedLineId *string `json:"DedicatedLineId,omitempty" xml:"DedicatedLineId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteDedicatedLineRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDedicatedLineRequest) GoString() string {
	return s.String()
}

func (s *DeleteDedicatedLineRequest) SetDedicatedLineId(v string) *DeleteDedicatedLineRequest {
	s.DedicatedLineId = &v
	return s
}

func (s *DeleteDedicatedLineRequest) SetInstanceId(v string) *DeleteDedicatedLineRequest {
	s.InstanceId = &v
	return s
}

type DeleteDedicatedLineResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDedicatedLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDedicatedLineResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDedicatedLineResponseBody) SetRequestId(v string) *DeleteDedicatedLineResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDedicatedLineResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDedicatedLineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDedicatedLineResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDedicatedLineResponse) GoString() string {
	return s.String()
}

func (s *DeleteDedicatedLineResponse) SetHeaders(v map[string]*string) *DeleteDedicatedLineResponse {
	s.Headers = v
	return s
}

func (s *DeleteDedicatedLineResponse) SetBody(v *DeleteDedicatedLineResponseBody) *DeleteDedicatedLineResponse {
	s.Body = v
	return s
}

type ListIpRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Ip地址
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// IP地址段的UID
	IpBlockId *string `json:"IpBlockId,omitempty" xml:"IpBlockId,omitempty"`
}

func (s ListIpRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIpRequest) GoString() string {
	return s.String()
}

func (s *ListIpRequest) SetInstanceId(v string) *ListIpRequest {
	s.InstanceId = &v
	return s
}

func (s *ListIpRequest) SetMaxResults(v int32) *ListIpRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpRequest) SetNextToken(v string) *ListIpRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpRequest) SetIpAddress(v string) *ListIpRequest {
	s.IpAddress = &v
	return s
}

func (s *ListIpRequest) SetIpBlockId(v string) *ListIpRequest {
	s.IpBlockId = &v
	return s
}

type ListIpResponseBody struct {
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据量
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 数组，返回示例目录。
	Ip []*ListIpResponseBodyIp `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
}

func (s ListIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIpResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpResponseBody) SetTotalCount(v int32) *ListIpResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIpResponseBody) SetRequestId(v string) *ListIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpResponseBody) SetNextToken(v string) *ListIpResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpResponseBody) SetMaxResults(v int64) *ListIpResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpResponseBody) SetIp(v []*ListIpResponseBodyIp) *ListIpResponseBody {
	s.Ip = v
	return s
}

type ListIpResponseBodyIp struct {
	// 设备端口名称
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// 设备MAC
	DeviceMac *string `json:"DeviceMac,omitempty" xml:"DeviceMac,omitempty"`
	// IP地址
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// 资源一级ID
	IpId *string `json:"IpId,omitempty" xml:"IpId,omitempty"`
	// 状态 using available lock
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 地址段
	ParentIpBlock *string `json:"ParentIpBlock,omitempty" xml:"ParentIpBlock,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 园区层级
	ZoneLayer []*ListIpResponseBodyIpZoneLayer `json:"ZoneLayer,omitempty" xml:"ZoneLayer,omitempty" type:"Repeated"`
	// 业务类型名称
	BusinessTypeName *string `json:"BusinessTypeName,omitempty" xml:"BusinessTypeName,omitempty"`
	// 业务类型UID
	BusinessTypeId *string `json:"BusinessTypeId,omitempty" xml:"BusinessTypeId,omitempty"`
	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
}

func (s ListIpResponseBodyIp) String() string {
	return tea.Prettify(s)
}

func (s ListIpResponseBodyIp) GoString() string {
	return s.String()
}

func (s *ListIpResponseBodyIp) SetPort(v string) *ListIpResponseBodyIp {
	s.Port = &v
	return s
}

func (s *ListIpResponseBodyIp) SetDeviceMac(v string) *ListIpResponseBodyIp {
	s.DeviceMac = &v
	return s
}

func (s *ListIpResponseBodyIp) SetIpAddress(v string) *ListIpResponseBodyIp {
	s.IpAddress = &v
	return s
}

func (s *ListIpResponseBodyIp) SetIpId(v string) *ListIpResponseBodyIp {
	s.IpId = &v
	return s
}

func (s *ListIpResponseBodyIp) SetStatus(v string) *ListIpResponseBodyIp {
	s.Status = &v
	return s
}

func (s *ListIpResponseBodyIp) SetParentIpBlock(v string) *ListIpResponseBodyIp {
	s.ParentIpBlock = &v
	return s
}

func (s *ListIpResponseBodyIp) SetCreateTime(v string) *ListIpResponseBodyIp {
	s.CreateTime = &v
	return s
}

func (s *ListIpResponseBodyIp) SetZoneLayer(v []*ListIpResponseBodyIpZoneLayer) *ListIpResponseBodyIp {
	s.ZoneLayer = v
	return s
}

func (s *ListIpResponseBodyIp) SetBusinessTypeName(v string) *ListIpResponseBodyIp {
	s.BusinessTypeName = &v
	return s
}

func (s *ListIpResponseBodyIp) SetBusinessTypeId(v string) *ListIpResponseBodyIp {
	s.BusinessTypeId = &v
	return s
}

func (s *ListIpResponseBodyIp) SetDeviceName(v string) *ListIpResponseBodyIp {
	s.DeviceName = &v
	return s
}

type ListIpResponseBodyIpZoneLayer struct {
	// 园区层级名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 园区层级值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIpResponseBodyIpZoneLayer) String() string {
	return tea.Prettify(s)
}

func (s ListIpResponseBodyIpZoneLayer) GoString() string {
	return s.String()
}

func (s *ListIpResponseBodyIpZoneLayer) SetName(v string) *ListIpResponseBodyIpZoneLayer {
	s.Name = &v
	return s
}

func (s *ListIpResponseBodyIpZoneLayer) SetValue(v string) *ListIpResponseBodyIpZoneLayer {
	s.Value = &v
	return s
}

type ListIpResponse struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIpResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIpResponse) GoString() string {
	return s.String()
}

func (s *ListIpResponse) SetHeaders(v map[string]*string) *ListIpResponse {
	s.Headers = v
	return s
}

func (s *ListIpResponse) SetBody(v *ListIpResponseBody) *ListIpResponse {
	s.Body = v
	return s
}

type ListConfigurationSpecificationsRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 架构类型
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 配置规范名字
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
}

func (s ListConfigurationSpecificationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConfigurationSpecificationsRequest) GoString() string {
	return s.String()
}

func (s *ListConfigurationSpecificationsRequest) SetInstanceId(v string) *ListConfigurationSpecificationsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListConfigurationSpecificationsRequest) SetMaxResults(v int32) *ListConfigurationSpecificationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListConfigurationSpecificationsRequest) SetNextToken(v string) *ListConfigurationSpecificationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListConfigurationSpecificationsRequest) SetArchitecture(v string) *ListConfigurationSpecificationsRequest {
	s.Architecture = &v
	return s
}

func (s *ListConfigurationSpecificationsRequest) SetRole(v string) *ListConfigurationSpecificationsRequest {
	s.Role = &v
	return s
}

func (s *ListConfigurationSpecificationsRequest) SetVendor(v string) *ListConfigurationSpecificationsRequest {
	s.Vendor = &v
	return s
}

func (s *ListConfigurationSpecificationsRequest) SetModel(v string) *ListConfigurationSpecificationsRequest {
	s.Model = &v
	return s
}

func (s *ListConfigurationSpecificationsRequest) SetSpecificationName(v string) *ListConfigurationSpecificationsRequest {
	s.SpecificationName = &v
	return s
}

type ListConfigurationSpecificationsResponseBody struct {
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 数组，返回示例目录。
	ConfigurationSpecification []*ListConfigurationSpecificationsResponseBodyConfigurationSpecification `json:"ConfigurationSpecification,omitempty" xml:"ConfigurationSpecification,omitempty" type:"Repeated"`
	// 本次读取的最大数据量
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListConfigurationSpecificationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConfigurationSpecificationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigurationSpecificationsResponseBody) SetTotalCount(v int32) *ListConfigurationSpecificationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListConfigurationSpecificationsResponseBody) SetRequestId(v string) *ListConfigurationSpecificationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConfigurationSpecificationsResponseBody) SetNextToken(v int32) *ListConfigurationSpecificationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListConfigurationSpecificationsResponseBody) SetConfigurationSpecification(v []*ListConfigurationSpecificationsResponseBodyConfigurationSpecification) *ListConfigurationSpecificationsResponseBody {
	s.ConfigurationSpecification = v
	return s
}

func (s *ListConfigurationSpecificationsResponseBody) SetMaxResults(v int64) *ListConfigurationSpecificationsResponseBody {
	s.MaxResults = &v
	return s
}

type ListConfigurationSpecificationsResponseBodyConfigurationSpecification struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 配置规范名字
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// 架构
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 规范内容
	SpecificationContent *string `json:"SpecificationContent,omitempty" xml:"SpecificationContent,omitempty"`
	// 相关变量
	RelatedVariate []*string `json:"RelatedVariate,omitempty" xml:"RelatedVariate,omitempty" type:"Repeated"`
	// 配置规范id
	ConfigurationSpecificationId *string `json:"ConfigurationSpecificationId,omitempty" xml:"ConfigurationSpecificationId,omitempty"`
}

func (s ListConfigurationSpecificationsResponseBodyConfigurationSpecification) String() string {
	return tea.Prettify(s)
}

func (s ListConfigurationSpecificationsResponseBodyConfigurationSpecification) GoString() string {
	return s.String()
}

func (s *ListConfigurationSpecificationsResponseBodyConfigurationSpecification) SetCreateTime(v string) *ListConfigurationSpecificationsResponseBodyConfigurationSpecification {
	s.CreateTime = &v
	return s
}

func (s *ListConfigurationSpecificationsResponseBodyConfigurationSpecification) SetUpdateTime(v string) *ListConfigurationSpecificationsResponseBodyConfigurationSpecification {
	s.UpdateTime = &v
	return s
}

func (s *ListConfigurationSpecificationsResponseBodyConfigurationSpecification) SetSpecificationName(v string) *ListConfigurationSpecificationsResponseBodyConfigurationSpecification {
	s.SpecificationName = &v
	return s
}

func (s *ListConfigurationSpecificationsResponseBodyConfigurationSpecification) SetArchitecture(v string) *ListConfigurationSpecificationsResponseBodyConfigurationSpecification {
	s.Architecture = &v
	return s
}

func (s *ListConfigurationSpecificationsResponseBodyConfigurationSpecification) SetRole(v string) *ListConfigurationSpecificationsResponseBodyConfigurationSpecification {
	s.Role = &v
	return s
}

func (s *ListConfigurationSpecificationsResponseBodyConfigurationSpecification) SetModel(v string) *ListConfigurationSpecificationsResponseBodyConfigurationSpecification {
	s.Model = &v
	return s
}

func (s *ListConfigurationSpecificationsResponseBodyConfigurationSpecification) SetVendor(v string) *ListConfigurationSpecificationsResponseBodyConfigurationSpecification {
	s.Vendor = &v
	return s
}

func (s *ListConfigurationSpecificationsResponseBodyConfigurationSpecification) SetSpecificationContent(v string) *ListConfigurationSpecificationsResponseBodyConfigurationSpecification {
	s.SpecificationContent = &v
	return s
}

func (s *ListConfigurationSpecificationsResponseBodyConfigurationSpecification) SetRelatedVariate(v []*string) *ListConfigurationSpecificationsResponseBodyConfigurationSpecification {
	s.RelatedVariate = v
	return s
}

func (s *ListConfigurationSpecificationsResponseBodyConfigurationSpecification) SetConfigurationSpecificationId(v string) *ListConfigurationSpecificationsResponseBodyConfigurationSpecification {
	s.ConfigurationSpecificationId = &v
	return s
}

type ListConfigurationSpecificationsResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConfigurationSpecificationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConfigurationSpecificationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConfigurationSpecificationsResponse) GoString() string {
	return s.String()
}

func (s *ListConfigurationSpecificationsResponse) SetHeaders(v map[string]*string) *ListConfigurationSpecificationsResponse {
	s.Headers = v
	return s
}

func (s *ListConfigurationSpecificationsResponse) SetBody(v *ListConfigurationSpecificationsResponseBody) *ListConfigurationSpecificationsResponse {
	s.Body = v
	return s
}

type EnableNotificationRequest struct {
	// 通知对象
	List []*EnableNotificationRequestList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableNotificationRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableNotificationRequest) GoString() string {
	return s.String()
}

func (s *EnableNotificationRequest) SetList(v []*EnableNotificationRequestList) *EnableNotificationRequest {
	s.List = v
	return s
}

func (s *EnableNotificationRequest) SetInstanceId(v string) *EnableNotificationRequest {
	s.InstanceId = &v
	return s
}

type EnableNotificationRequestList struct {
	// 类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 监控项ID
	MonitorItemId *string `json:"MonitorItemId,omitempty" xml:"MonitorItemId,omitempty"`
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 聚合数据ID
	AggregateDataId *string `json:"AggregateDataId,omitempty" xml:"AggregateDataId,omitempty"`
	// 专线ID
	DedicatedLineId *string `json:"DedicatedLineId,omitempty" xml:"DedicatedLineId,omitempty"`
	// 端口集ID
	PortCollectionId *string `json:"PortCollectionId,omitempty" xml:"PortCollectionId,omitempty"`
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s EnableNotificationRequestList) String() string {
	return tea.Prettify(s)
}

func (s EnableNotificationRequestList) GoString() string {
	return s.String()
}

func (s *EnableNotificationRequestList) SetType(v string) *EnableNotificationRequestList {
	s.Type = &v
	return s
}

func (s *EnableNotificationRequestList) SetMonitorItemId(v string) *EnableNotificationRequestList {
	s.MonitorItemId = &v
	return s
}

func (s *EnableNotificationRequestList) SetDeviceId(v string) *EnableNotificationRequestList {
	s.DeviceId = &v
	return s
}

func (s *EnableNotificationRequestList) SetAggregateDataId(v string) *EnableNotificationRequestList {
	s.AggregateDataId = &v
	return s
}

func (s *EnableNotificationRequestList) SetDedicatedLineId(v string) *EnableNotificationRequestList {
	s.DedicatedLineId = &v
	return s
}

func (s *EnableNotificationRequestList) SetPortCollectionId(v string) *EnableNotificationRequestList {
	s.PortCollectionId = &v
	return s
}

func (s *EnableNotificationRequestList) SetAppId(v string) *EnableNotificationRequestList {
	s.AppId = &v
	return s
}

type EnableNotificationResponseBody struct {
	// request id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableNotificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableNotificationResponseBody) GoString() string {
	return s.String()
}

func (s *EnableNotificationResponseBody) SetRequestId(v string) *EnableNotificationResponseBody {
	s.RequestId = &v
	return s
}

type EnableNotificationResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableNotificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableNotificationResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableNotificationResponse) GoString() string {
	return s.String()
}

func (s *EnableNotificationResponse) SetHeaders(v map[string]*string) *EnableNotificationResponse {
	s.Headers = v
	return s
}

func (s *EnableNotificationResponse) SetBody(v *EnableNotificationResponseBody) *EnableNotificationResponse {
	s.Body = v
	return s
}

type UpdateDevicePropertyRequest struct {
	// 实例 ID。
	DevicePropertyId *string `json:"DevicePropertyId,omitempty" xml:"DevicePropertyId,omitempty"`
	// 属性格式
	PropertyFormat *string `json:"PropertyFormat,omitempty" xml:"PropertyFormat,omitempty"`
	// 属性内容
	PropertyContent *string `json:"PropertyContent,omitempty" xml:"PropertyContent,omitempty"`
	// 属性名称
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateDevicePropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevicePropertyRequest) GoString() string {
	return s.String()
}

func (s *UpdateDevicePropertyRequest) SetDevicePropertyId(v string) *UpdateDevicePropertyRequest {
	s.DevicePropertyId = &v
	return s
}

func (s *UpdateDevicePropertyRequest) SetPropertyFormat(v string) *UpdateDevicePropertyRequest {
	s.PropertyFormat = &v
	return s
}

func (s *UpdateDevicePropertyRequest) SetPropertyContent(v string) *UpdateDevicePropertyRequest {
	s.PropertyContent = &v
	return s
}

func (s *UpdateDevicePropertyRequest) SetPropertyName(v string) *UpdateDevicePropertyRequest {
	s.PropertyName = &v
	return s
}

func (s *UpdateDevicePropertyRequest) SetInstanceId(v string) *UpdateDevicePropertyRequest {
	s.InstanceId = &v
	return s
}

type UpdateDevicePropertyResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDevicePropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevicePropertyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDevicePropertyResponseBody) SetRequestId(v string) *UpdateDevicePropertyResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDevicePropertyResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDevicePropertyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDevicePropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevicePropertyResponse) GoString() string {
	return s.String()
}

func (s *UpdateDevicePropertyResponse) SetHeaders(v map[string]*string) *UpdateDevicePropertyResponse {
	s.Headers = v
	return s
}

func (s *UpdateDevicePropertyResponse) SetBody(v *UpdateDevicePropertyResponseBody) *UpdateDevicePropertyResponse {
	s.Body = v
	return s
}

type ListNotificationHistoriesRequest struct {
	// 标记当前开始读取的位置，置空表示从头开始
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 监控项ID
	MonitorItemId *string `json:"MonitorItemId,omitempty" xml:"MonitorItemId,omitempty"`
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 专线ID
	DedicatedLineId *string `json:"DedicatedLineId,omitempty" xml:"DedicatedLineId,omitempty"`
	// 聚合数据ID
	AggregateDataId *string `json:"AggregateDataId,omitempty" xml:"AggregateDataId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 端口集ID
	PortCollectionId *string `json:"PortCollectionId,omitempty" xml:"PortCollectionId,omitempty"`
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListNotificationHistoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationHistoriesRequest) GoString() string {
	return s.String()
}

func (s *ListNotificationHistoriesRequest) SetNextToken(v string) *ListNotificationHistoriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListNotificationHistoriesRequest) SetMaxResults(v int32) *ListNotificationHistoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNotificationHistoriesRequest) SetMonitorItemId(v string) *ListNotificationHistoriesRequest {
	s.MonitorItemId = &v
	return s
}

func (s *ListNotificationHistoriesRequest) SetDeviceId(v string) *ListNotificationHistoriesRequest {
	s.DeviceId = &v
	return s
}

func (s *ListNotificationHistoriesRequest) SetType(v string) *ListNotificationHistoriesRequest {
	s.Type = &v
	return s
}

func (s *ListNotificationHistoriesRequest) SetDedicatedLineId(v string) *ListNotificationHistoriesRequest {
	s.DedicatedLineId = &v
	return s
}

func (s *ListNotificationHistoriesRequest) SetAggregateDataId(v string) *ListNotificationHistoriesRequest {
	s.AggregateDataId = &v
	return s
}

func (s *ListNotificationHistoriesRequest) SetInstanceId(v string) *ListNotificationHistoriesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListNotificationHistoriesRequest) SetPortCollectionId(v string) *ListNotificationHistoriesRequest {
	s.PortCollectionId = &v
	return s
}

func (s *ListNotificationHistoriesRequest) SetAppId(v string) *ListNotificationHistoriesRequest {
	s.AppId = &v
	return s
}

type ListNotificationHistoriesResponseBody struct {
	// TotalCount本次请求条件下的数据总量，此参数为可选参数，默认可不返回
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// request Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// MaxResults本次请求所返回的最大记录条数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 数据列表
	NotificationHistories []*ListNotificationHistoriesResponseBodyNotificationHistories `json:"NotificationHistories,omitempty" xml:"NotificationHistories,omitempty" type:"Repeated"`
}

func (s ListNotificationHistoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNotificationHistoriesResponseBody) SetTotalCount(v int32) *ListNotificationHistoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNotificationHistoriesResponseBody) SetRequestId(v string) *ListNotificationHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNotificationHistoriesResponseBody) SetNextToken(v string) *ListNotificationHistoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNotificationHistoriesResponseBody) SetMaxResults(v int32) *ListNotificationHistoriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNotificationHistoriesResponseBody) SetNotificationHistories(v []*ListNotificationHistoriesResponseBodyNotificationHistories) *ListNotificationHistoriesResponseBody {
	s.NotificationHistories = v
	return s
}

type ListNotificationHistoriesResponseBodyNotificationHistories struct {
	// 发送时间
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// 发送方式
	NotificationMode *string `json:"NotificationMode,omitempty" xml:"NotificationMode,omitempty"`
	// 发送状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 输出内容
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// 发送内容
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 监控项ID
	MonitorItemId *string `json:"MonitorItemId,omitempty" xml:"MonitorItemId,omitempty"`
	// 通知组ID
	NotificationGroupId *string `json:"NotificationGroupId,omitempty" xml:"NotificationGroupId,omitempty"`
	// 通知组名称
	NotificationGroupName *string `json:"NotificationGroupName,omitempty" xml:"NotificationGroupName,omitempty"`
	// 专线ID
	DedicatedLineId *string `json:"DedicatedLineId,omitempty" xml:"DedicatedLineId,omitempty"`
	// 聚合数据ID
	AggregateDataId *string `json:"AggregateDataId,omitempty" xml:"AggregateDataId,omitempty"`
	// 端口集ID
	PortCollectionId *string `json:"PortCollectionId,omitempty" xml:"PortCollectionId,omitempty"`
}

func (s ListNotificationHistoriesResponseBodyNotificationHistories) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationHistoriesResponseBodyNotificationHistories) GoString() string {
	return s.String()
}

func (s *ListNotificationHistoriesResponseBodyNotificationHistories) SetTime(v string) *ListNotificationHistoriesResponseBodyNotificationHistories {
	s.Time = &v
	return s
}

func (s *ListNotificationHistoriesResponseBodyNotificationHistories) SetNotificationMode(v string) *ListNotificationHistoriesResponseBodyNotificationHistories {
	s.NotificationMode = &v
	return s
}

func (s *ListNotificationHistoriesResponseBodyNotificationHistories) SetStatus(v string) *ListNotificationHistoriesResponseBodyNotificationHistories {
	s.Status = &v
	return s
}

func (s *ListNotificationHistoriesResponseBodyNotificationHistories) SetOutput(v string) *ListNotificationHistoriesResponseBodyNotificationHistories {
	s.Output = &v
	return s
}

func (s *ListNotificationHistoriesResponseBodyNotificationHistories) SetMessage(v string) *ListNotificationHistoriesResponseBodyNotificationHistories {
	s.Message = &v
	return s
}

func (s *ListNotificationHistoriesResponseBodyNotificationHistories) SetDeviceId(v string) *ListNotificationHistoriesResponseBodyNotificationHistories {
	s.DeviceId = &v
	return s
}

func (s *ListNotificationHistoriesResponseBodyNotificationHistories) SetMonitorItemId(v string) *ListNotificationHistoriesResponseBodyNotificationHistories {
	s.MonitorItemId = &v
	return s
}

func (s *ListNotificationHistoriesResponseBodyNotificationHistories) SetNotificationGroupId(v string) *ListNotificationHistoriesResponseBodyNotificationHistories {
	s.NotificationGroupId = &v
	return s
}

func (s *ListNotificationHistoriesResponseBodyNotificationHistories) SetNotificationGroupName(v string) *ListNotificationHistoriesResponseBodyNotificationHistories {
	s.NotificationGroupName = &v
	return s
}

func (s *ListNotificationHistoriesResponseBodyNotificationHistories) SetDedicatedLineId(v string) *ListNotificationHistoriesResponseBodyNotificationHistories {
	s.DedicatedLineId = &v
	return s
}

func (s *ListNotificationHistoriesResponseBodyNotificationHistories) SetAggregateDataId(v string) *ListNotificationHistoriesResponseBodyNotificationHistories {
	s.AggregateDataId = &v
	return s
}

func (s *ListNotificationHistoriesResponseBodyNotificationHistories) SetPortCollectionId(v string) *ListNotificationHistoriesResponseBodyNotificationHistories {
	s.PortCollectionId = &v
	return s
}

type ListNotificationHistoriesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNotificationHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNotificationHistoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationHistoriesResponse) GoString() string {
	return s.String()
}

func (s *ListNotificationHistoriesResponse) SetHeaders(v map[string]*string) *ListNotificationHistoriesResponse {
	s.Headers = v
	return s
}

func (s *ListNotificationHistoriesResponse) SetBody(v *ListNotificationHistoriesResponseBody) *ListNotificationHistoriesResponse {
	s.Body = v
	return s
}

type DeleteDevicePropertyRequest struct {
	// 实例 ID。
	DevicePropertyId *string `json:"DevicePropertyId,omitempty" xml:"DevicePropertyId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteDevicePropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevicePropertyRequest) GoString() string {
	return s.String()
}

func (s *DeleteDevicePropertyRequest) SetDevicePropertyId(v string) *DeleteDevicePropertyRequest {
	s.DevicePropertyId = &v
	return s
}

func (s *DeleteDevicePropertyRequest) SetInstanceId(v string) *DeleteDevicePropertyRequest {
	s.InstanceId = &v
	return s
}

type DeleteDevicePropertyResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDevicePropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevicePropertyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDevicePropertyResponseBody) SetRequestId(v string) *DeleteDevicePropertyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDevicePropertyResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDevicePropertyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDevicePropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevicePropertyResponse) GoString() string {
	return s.String()
}

func (s *DeleteDevicePropertyResponse) SetHeaders(v map[string]*string) *DeleteDevicePropertyResponse {
	s.Headers = v
	return s
}

func (s *DeleteDevicePropertyResponse) SetBody(v *DeleteDevicePropertyResponseBody) *DeleteDevicePropertyResponse {
	s.Body = v
	return s
}

type ListResourceInstancesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 建设项目资源id
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
}

func (s ListResourceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceInstancesRequest) SetInstanceId(v string) *ListResourceInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListResourceInstancesRequest) SetSetupProjectId(v string) *ListResourceInstancesRequest {
	s.SetupProjectId = &v
	return s
}

type ListResourceInstancesResponseBody struct {
	LogicResource []*ListResourceInstancesResponseBodyLogicResource `json:"LogicResource,omitempty" xml:"LogicResource,omitempty" type:"Repeated"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListResourceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceInstancesResponseBody) SetLogicResource(v []*ListResourceInstancesResponseBodyLogicResource) *ListResourceInstancesResponseBody {
	s.LogicResource = v
	return s
}

func (s *ListResourceInstancesResponseBody) SetRequestId(v string) *ListResourceInstancesResponseBody {
	s.RequestId = &v
	return s
}

type ListResourceInstancesResponseBodyLogicResource struct {
	ResourceType      *string                                                  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceAttribute *string                                                  `json:"ResourceAttribute,omitempty" xml:"ResourceAttribute,omitempty"`
	KeyList           []*ListResourceInstancesResponseBodyLogicResourceKeyList `json:"KeyList,omitempty" xml:"KeyList,omitempty" type:"Repeated"`
}

func (s ListResourceInstancesResponseBodyLogicResource) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInstancesResponseBodyLogicResource) GoString() string {
	return s.String()
}

func (s *ListResourceInstancesResponseBodyLogicResource) SetResourceType(v string) *ListResourceInstancesResponseBodyLogicResource {
	s.ResourceType = &v
	return s
}

func (s *ListResourceInstancesResponseBodyLogicResource) SetResourceAttribute(v string) *ListResourceInstancesResponseBodyLogicResource {
	s.ResourceAttribute = &v
	return s
}

func (s *ListResourceInstancesResponseBodyLogicResource) SetKeyList(v []*ListResourceInstancesResponseBodyLogicResourceKeyList) *ListResourceInstancesResponseBodyLogicResource {
	s.KeyList = v
	return s
}

type ListResourceInstancesResponseBodyLogicResourceKeyList struct {
	Key            *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value          *string `json:"Value,omitempty" xml:"Value,omitempty"`
	KeyAttribute   *string `json:"KeyAttribute,omitempty" xml:"KeyAttribute,omitempty"`
	KeyAction      *string `json:"KeyAction,omitempty" xml:"KeyAction,omitempty"`
	KeyDescription *string `json:"KeyDescription,omitempty" xml:"KeyDescription,omitempty"`
}

func (s ListResourceInstancesResponseBodyLogicResourceKeyList) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInstancesResponseBodyLogicResourceKeyList) GoString() string {
	return s.String()
}

func (s *ListResourceInstancesResponseBodyLogicResourceKeyList) SetKey(v string) *ListResourceInstancesResponseBodyLogicResourceKeyList {
	s.Key = &v
	return s
}

func (s *ListResourceInstancesResponseBodyLogicResourceKeyList) SetValue(v string) *ListResourceInstancesResponseBodyLogicResourceKeyList {
	s.Value = &v
	return s
}

func (s *ListResourceInstancesResponseBodyLogicResourceKeyList) SetKeyAttribute(v string) *ListResourceInstancesResponseBodyLogicResourceKeyList {
	s.KeyAttribute = &v
	return s
}

func (s *ListResourceInstancesResponseBodyLogicResourceKeyList) SetKeyAction(v string) *ListResourceInstancesResponseBodyLogicResourceKeyList {
	s.KeyAction = &v
	return s
}

func (s *ListResourceInstancesResponseBodyLogicResourceKeyList) SetKeyDescription(v string) *ListResourceInstancesResponseBodyLogicResourceKeyList {
	s.KeyDescription = &v
	return s
}

type ListResourceInstancesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListResourceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceInstancesResponse) SetHeaders(v map[string]*string) *ListResourceInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceInstancesResponse) SetBody(v *ListResourceInstancesResponseBody) *ListResourceInstancesResponse {
	s.Body = v
	return s
}

type ListIpBlocksRequest struct {
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 网络类型 PRIVATE PUBLIC
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// 是否树状展示地址段
	TreeType *bool `json:"TreeType,omitempty" xml:"TreeType,omitempty"`
	// 地址段
	IpBlock *string `json:"IpBlock,omitempty" xml:"IpBlock,omitempty"`
	// IP地址
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// 地址段状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 园区名称，NetType为PUBLIC有效
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	// 公网地址类型，NetType为PUBLIC有效
	NetBusiness *string `json:"NetBusiness,omitempty" xml:"NetBusiness,omitempty"`
}

func (s ListIpBlocksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIpBlocksRequest) GoString() string {
	return s.String()
}

func (s *ListIpBlocksRequest) SetMaxResults(v int32) *ListIpBlocksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpBlocksRequest) SetNextToken(v string) *ListIpBlocksRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpBlocksRequest) SetInstanceId(v string) *ListIpBlocksRequest {
	s.InstanceId = &v
	return s
}

func (s *ListIpBlocksRequest) SetNetType(v string) *ListIpBlocksRequest {
	s.NetType = &v
	return s
}

func (s *ListIpBlocksRequest) SetTreeType(v bool) *ListIpBlocksRequest {
	s.TreeType = &v
	return s
}

func (s *ListIpBlocksRequest) SetIpBlock(v string) *ListIpBlocksRequest {
	s.IpBlock = &v
	return s
}

func (s *ListIpBlocksRequest) SetIp(v string) *ListIpBlocksRequest {
	s.Ip = &v
	return s
}

func (s *ListIpBlocksRequest) SetStatus(v string) *ListIpBlocksRequest {
	s.Status = &v
	return s
}

func (s *ListIpBlocksRequest) SetZoneName(v string) *ListIpBlocksRequest {
	s.ZoneName = &v
	return s
}

func (s *ListIpBlocksRequest) SetNetBusiness(v string) *ListIpBlocksRequest {
	s.NetBusiness = &v
	return s
}

type ListIpBlocksResponseBody struct {
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 每页数量。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 地址段对象
	IpBlock []*ListIpBlocksResponseBodyIpBlock `json:"IpBlock,omitempty" xml:"IpBlock,omitempty" type:"Repeated"`
}

func (s ListIpBlocksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIpBlocksResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpBlocksResponseBody) SetTotalCount(v int32) *ListIpBlocksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIpBlocksResponseBody) SetRequestId(v string) *ListIpBlocksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpBlocksResponseBody) SetNextToken(v int32) *ListIpBlocksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpBlocksResponseBody) SetMaxResults(v int32) *ListIpBlocksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpBlocksResponseBody) SetIpBlock(v []*ListIpBlocksResponseBodyIpBlock) *ListIpBlocksResponseBody {
	s.IpBlock = v
	return s
}

type ListIpBlocksResponseBodyIpBlock struct {
	// IP地址段UID
	IpBlockId *string `json:"IpBlockId,omitempty" xml:"IpBlockId,omitempty"`
	// IP段
	IpBlockCode *string `json:"IpBlockCode,omitempty" xml:"IpBlockCode,omitempty"`
	// 父地址段UID
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// 公网私网标志 PUBLIC PRIVATE
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// 状态： using available lock
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// IP归属
	Ownership *string `json:"Ownership,omitempty" xml:"Ownership,omitempty"`
	// 地址类别 IPV4
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 备注
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 备份设备名称
	BackupDeviceName *string `json:"BackupDeviceName,omitempty" xml:"BackupDeviceName,omitempty"`
	// 园区层级
	ZoneLayer []*ListIpBlocksResponseBodyIpBlockZoneLayer `json:"ZoneLayer,omitempty" xml:"ZoneLayer,omitempty" type:"Repeated"`
	// 业务类型UID
	BusinessTypeId *string `json:"BusinessTypeId,omitempty" xml:"BusinessTypeId,omitempty"`
	// 业务类型名称
	BusinessTypeName *string `json:"BusinessTypeName,omitempty" xml:"BusinessTypeName,omitempty"`
	// 公网地址类型 INC GUEST VIP
	NetBusiness *string `json:"NetBusiness,omitempty" xml:"NetBusiness,omitempty"`
	// IP用途
	Application *string `json:"Application,omitempty" xml:"Application,omitempty"`
	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// 园区名称
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	// 地址段子段列表
	IpBlocks []*string `json:"IpBlocks,omitempty" xml:"IpBlocks,omitempty" type:"Repeated"`
}

func (s ListIpBlocksResponseBodyIpBlock) String() string {
	return tea.Prettify(s)
}

func (s ListIpBlocksResponseBodyIpBlock) GoString() string {
	return s.String()
}

func (s *ListIpBlocksResponseBodyIpBlock) SetIpBlockId(v string) *ListIpBlocksResponseBodyIpBlock {
	s.IpBlockId = &v
	return s
}

func (s *ListIpBlocksResponseBodyIpBlock) SetIpBlockCode(v string) *ListIpBlocksResponseBodyIpBlock {
	s.IpBlockCode = &v
	return s
}

func (s *ListIpBlocksResponseBodyIpBlock) SetParentId(v string) *ListIpBlocksResponseBodyIpBlock {
	s.ParentId = &v
	return s
}

func (s *ListIpBlocksResponseBodyIpBlock) SetNetType(v string) *ListIpBlocksResponseBodyIpBlock {
	s.NetType = &v
	return s
}

func (s *ListIpBlocksResponseBodyIpBlock) SetStatus(v string) *ListIpBlocksResponseBodyIpBlock {
	s.Status = &v
	return s
}

func (s *ListIpBlocksResponseBodyIpBlock) SetOwnership(v string) *ListIpBlocksResponseBodyIpBlock {
	s.Ownership = &v
	return s
}

func (s *ListIpBlocksResponseBodyIpBlock) SetCategory(v string) *ListIpBlocksResponseBodyIpBlock {
	s.Category = &v
	return s
}

func (s *ListIpBlocksResponseBodyIpBlock) SetDescription(v string) *ListIpBlocksResponseBodyIpBlock {
	s.Description = &v
	return s
}

func (s *ListIpBlocksResponseBodyIpBlock) SetBackupDeviceName(v string) *ListIpBlocksResponseBodyIpBlock {
	s.BackupDeviceName = &v
	return s
}

func (s *ListIpBlocksResponseBodyIpBlock) SetZoneLayer(v []*ListIpBlocksResponseBodyIpBlockZoneLayer) *ListIpBlocksResponseBodyIpBlock {
	s.ZoneLayer = v
	return s
}

func (s *ListIpBlocksResponseBodyIpBlock) SetBusinessTypeId(v string) *ListIpBlocksResponseBodyIpBlock {
	s.BusinessTypeId = &v
	return s
}

func (s *ListIpBlocksResponseBodyIpBlock) SetBusinessTypeName(v string) *ListIpBlocksResponseBodyIpBlock {
	s.BusinessTypeName = &v
	return s
}

func (s *ListIpBlocksResponseBodyIpBlock) SetNetBusiness(v string) *ListIpBlocksResponseBodyIpBlock {
	s.NetBusiness = &v
	return s
}

func (s *ListIpBlocksResponseBodyIpBlock) SetApplication(v string) *ListIpBlocksResponseBodyIpBlock {
	s.Application = &v
	return s
}

func (s *ListIpBlocksResponseBodyIpBlock) SetDeviceName(v string) *ListIpBlocksResponseBodyIpBlock {
	s.DeviceName = &v
	return s
}

func (s *ListIpBlocksResponseBodyIpBlock) SetZoneName(v string) *ListIpBlocksResponseBodyIpBlock {
	s.ZoneName = &v
	return s
}

func (s *ListIpBlocksResponseBodyIpBlock) SetIpBlocks(v []*string) *ListIpBlocksResponseBodyIpBlock {
	s.IpBlocks = v
	return s
}

type ListIpBlocksResponseBodyIpBlockZoneLayer struct {
	// 园区层级名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 园区层级值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIpBlocksResponseBodyIpBlockZoneLayer) String() string {
	return tea.Prettify(s)
}

func (s ListIpBlocksResponseBodyIpBlockZoneLayer) GoString() string {
	return s.String()
}

func (s *ListIpBlocksResponseBodyIpBlockZoneLayer) SetName(v string) *ListIpBlocksResponseBodyIpBlockZoneLayer {
	s.Name = &v
	return s
}

func (s *ListIpBlocksResponseBodyIpBlockZoneLayer) SetValue(v string) *ListIpBlocksResponseBodyIpBlockZoneLayer {
	s.Value = &v
	return s
}

type ListIpBlocksResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListIpBlocksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIpBlocksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIpBlocksResponse) GoString() string {
	return s.String()
}

func (s *ListIpBlocksResponse) SetHeaders(v map[string]*string) *ListIpBlocksResponse {
	s.Headers = v
	return s
}

func (s *ListIpBlocksResponse) SetBody(v *ListIpBlocksResponseBody) *ListIpBlocksResponse {
	s.Body = v
	return s
}

type ListDeviceResourcesRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 业务类型
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// List类型
	ListType *string `json:"ListType,omitempty" xml:"ListType,omitempty"`
	// 建设项目id
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
}

func (s ListDeviceResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceResourcesRequest) SetInstanceId(v string) *ListDeviceResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDeviceResourcesRequest) SetMaxResults(v int32) *ListDeviceResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDeviceResourcesRequest) SetNextToken(v string) *ListDeviceResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDeviceResourcesRequest) SetBusinessType(v string) *ListDeviceResourcesRequest {
	s.BusinessType = &v
	return s
}

func (s *ListDeviceResourcesRequest) SetListType(v string) *ListDeviceResourcesRequest {
	s.ListType = &v
	return s
}

func (s *ListDeviceResourcesRequest) SetSetupProjectId(v string) *ListDeviceResourcesRequest {
	s.SetupProjectId = &v
	return s
}

type ListDeviceResourcesResponseBody struct {
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 数组，返回示例目录。
	DeviceResource []*ListDeviceResourcesResponseBodyDeviceResource `json:"DeviceResource,omitempty" xml:"DeviceResource,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据量
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListDeviceResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceResourcesResponseBody) SetTotalCount(v int32) *ListDeviceResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDeviceResourcesResponseBody) SetDeviceResource(v []*ListDeviceResourcesResponseBodyDeviceResource) *ListDeviceResourcesResponseBody {
	s.DeviceResource = v
	return s
}

func (s *ListDeviceResourcesResponseBody) SetRequestId(v string) *ListDeviceResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceResourcesResponseBody) SetNextToken(v int32) *ListDeviceResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDeviceResourcesResponseBody) SetMaxResults(v int64) *ListDeviceResourcesResponseBody {
	s.MaxResults = &v
	return s
}

type ListDeviceResourcesResponseBodyDeviceResource struct {
	// 建设项目id
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// 资源一级ID
	DeviceResourceId *string `json:"DeviceResourceId,omitempty" xml:"DeviceResourceId,omitempty"`
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 组号
	BlockNumber *string `json:"BlockNumber,omitempty" xml:"BlockNumber,omitempty"`
	// 设备号
	DeviceNumber *string `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 模型
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 主机名
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// 设备sn号
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// 配置任务Id
	ConfigTaskId *string `json:"ConfigTaskId,omitempty" xml:"ConfigTaskId,omitempty"`
	// 物理空间位置
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// 带内管理地址
	ManagerIp *string `json:"ManagerIp,omitempty" xml:"ManagerIp,omitempty"`
	// 交付登录地址
	DeliveryIp *string `json:"DeliveryIp,omitempty" xml:"DeliveryIp,omitempty"`
	// 配置生成
	GenerateConfig *string `json:"GenerateConfig,omitempty" xml:"GenerateConfig,omitempty"`
	// 配置下发状态
	ConfigTaskStatus *string `json:"ConfigTaskStatus,omitempty" xml:"ConfigTaskStatus,omitempty"`
	// 设备配置
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// 配置规范
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// 配置入参
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// loopback地址
	Loopback *string `json:"Loopback,omitempty" xml:"Loopback,omitempty"`
	// 设备互联地址
	InterConnection *string `json:"InterConnection,omitempty" xml:"InterConnection,omitempty"`
	// 设备业务地址
	Business *string `json:"Business,omitempty" xml:"Business,omitempty"`
	// 堆叠状态
	Stack *bool `json:"Stack,omitempty" xml:"Stack,omitempty"`
}

func (s ListDeviceResourcesResponseBodyDeviceResource) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceResourcesResponseBodyDeviceResource) GoString() string {
	return s.String()
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetSetupProjectId(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.SetupProjectId = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetDeviceResourceId(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.DeviceResourceId = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetRole(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.Role = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetBlockNumber(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.BlockNumber = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetDeviceNumber(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.DeviceNumber = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetVendor(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.Vendor = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetModel(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.Model = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetHostName(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.HostName = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetSn(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.Sn = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetConfigTaskId(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.ConfigTaskId = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetLocation(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.Location = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetManagerIp(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.ManagerIp = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetDeliveryIp(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.DeliveryIp = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetGenerateConfig(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.GenerateConfig = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetConfigTaskStatus(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.ConfigTaskStatus = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetConfig(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.Config = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetSpecification(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.Specification = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetParams(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.Params = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetLoopback(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.Loopback = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetInterConnection(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.InterConnection = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetBusiness(v string) *ListDeviceResourcesResponseBodyDeviceResource {
	s.Business = &v
	return s
}

func (s *ListDeviceResourcesResponseBodyDeviceResource) SetStack(v bool) *ListDeviceResourcesResponseBodyDeviceResource {
	s.Stack = &v
	return s
}

type ListDeviceResourcesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeviceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeviceResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceResourcesResponse) SetHeaders(v map[string]*string) *ListDeviceResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceResourcesResponse) SetBody(v *ListDeviceResourcesResponseBody) *ListDeviceResourcesResponse {
	s.Body = v
	return s
}

type ListResourceInformationsRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 架构id
	ArchitectureId *string `json:"ArchitectureId,omitempty" xml:"ArchitectureId,omitempty"`
}

func (s ListResourceInformationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInformationsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceInformationsRequest) SetInstanceId(v string) *ListResourceInformationsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListResourceInformationsRequest) SetMaxResults(v int32) *ListResourceInformationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceInformationsRequest) SetNextToken(v string) *ListResourceInformationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceInformationsRequest) SetArchitectureId(v string) *ListResourceInformationsRequest {
	s.ArchitectureId = &v
	return s
}

type ListResourceInformationsResponseBody struct {
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据量
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 数组，返回示例目录。
	ResourceInformation []*ListResourceInformationsResponseBodyResourceInformation `json:"ResourceInformation,omitempty" xml:"ResourceInformation,omitempty" type:"Repeated"`
}

func (s ListResourceInformationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInformationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceInformationsResponseBody) SetTotalCount(v int32) *ListResourceInformationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourceInformationsResponseBody) SetRequestId(v string) *ListResourceInformationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceInformationsResponseBody) SetNextToken(v int32) *ListResourceInformationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceInformationsResponseBody) SetMaxResults(v int64) *ListResourceInformationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListResourceInformationsResponseBody) SetResourceInformation(v []*ListResourceInformationsResponseBodyResourceInformation) *ListResourceInformationsResponseBody {
	s.ResourceInformation = v
	return s
}

type ListResourceInformationsResponseBodyResourceInformation struct {
	// 信息
	Information []*ListResourceInformationsResponseBodyResourceInformationInformation `json:"Information,omitempty" xml:"Information,omitempty" type:"Repeated"`
	// 资源一级ID
	ResourceInformationId *string `json:"ResourceInformationId,omitempty" xml:"ResourceInformationId,omitempty"`
	// 资源属性
	ResourceAttribute *string `json:"ResourceAttribute,omitempty" xml:"ResourceAttribute,omitempty"`
	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourceInformationsResponseBodyResourceInformation) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInformationsResponseBodyResourceInformation) GoString() string {
	return s.String()
}

func (s *ListResourceInformationsResponseBodyResourceInformation) SetInformation(v []*ListResourceInformationsResponseBodyResourceInformationInformation) *ListResourceInformationsResponseBodyResourceInformation {
	s.Information = v
	return s
}

func (s *ListResourceInformationsResponseBodyResourceInformation) SetResourceInformationId(v string) *ListResourceInformationsResponseBodyResourceInformation {
	s.ResourceInformationId = &v
	return s
}

func (s *ListResourceInformationsResponseBodyResourceInformation) SetResourceAttribute(v string) *ListResourceInformationsResponseBodyResourceInformation {
	s.ResourceAttribute = &v
	return s
}

func (s *ListResourceInformationsResponseBodyResourceInformation) SetResourceType(v string) *ListResourceInformationsResponseBodyResourceInformation {
	s.ResourceType = &v
	return s
}

type ListResourceInformationsResponseBodyResourceInformationInformation struct {
	// 键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 键属性
	KeyAttribute *string `json:"KeyAttribute,omitempty" xml:"KeyAttribute,omitempty"`
	// 键动作
	KeyAction *string `json:"KeyAction,omitempty" xml:"KeyAction,omitempty"`
	// 键描述
	KeyDescription *string `json:"KeyDescription,omitempty" xml:"KeyDescription,omitempty"`
}

func (s ListResourceInformationsResponseBodyResourceInformationInformation) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInformationsResponseBodyResourceInformationInformation) GoString() string {
	return s.String()
}

func (s *ListResourceInformationsResponseBodyResourceInformationInformation) SetKey(v string) *ListResourceInformationsResponseBodyResourceInformationInformation {
	s.Key = &v
	return s
}

func (s *ListResourceInformationsResponseBodyResourceInformationInformation) SetKeyAttribute(v string) *ListResourceInformationsResponseBodyResourceInformationInformation {
	s.KeyAttribute = &v
	return s
}

func (s *ListResourceInformationsResponseBodyResourceInformationInformation) SetKeyAction(v string) *ListResourceInformationsResponseBodyResourceInformationInformation {
	s.KeyAction = &v
	return s
}

func (s *ListResourceInformationsResponseBodyResourceInformationInformation) SetKeyDescription(v string) *ListResourceInformationsResponseBodyResourceInformationInformation {
	s.KeyDescription = &v
	return s
}

type ListResourceInformationsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListResourceInformationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourceInformationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInformationsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceInformationsResponse) SetHeaders(v map[string]*string) *ListResourceInformationsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceInformationsResponse) SetBody(v *ListResourceInformationsResponseBody) *ListResourceInformationsResponse {
	s.Body = v
	return s
}

type ListDeviceFormsRequest struct {
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListDeviceFormsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceFormsRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceFormsRequest) SetMaxResults(v int32) *ListDeviceFormsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDeviceFormsRequest) SetNextToken(v string) *ListDeviceFormsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDeviceFormsRequest) SetInstanceId(v string) *ListDeviceFormsRequest {
	s.InstanceId = &v
	return s
}

type ListDeviceFormsResponseBody struct {
	// 数组，返回示例目录。
	DeviceForms []*ListDeviceFormsResponseBodyDeviceForms `json:"DeviceForms,omitempty" xml:"DeviceForms,omitempty" type:"Repeated"`
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 每页数量。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListDeviceFormsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceFormsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceFormsResponseBody) SetDeviceForms(v []*ListDeviceFormsResponseBodyDeviceForms) *ListDeviceFormsResponseBody {
	s.DeviceForms = v
	return s
}

func (s *ListDeviceFormsResponseBody) SetTotalCount(v int32) *ListDeviceFormsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDeviceFormsResponseBody) SetRequestId(v string) *ListDeviceFormsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceFormsResponseBody) SetNextToken(v int32) *ListDeviceFormsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDeviceFormsResponseBody) SetMaxResults(v int32) *ListDeviceFormsResponseBody {
	s.MaxResults = &v
	return s
}

type ListDeviceFormsResponseBodyDeviceForms struct {
	// 是否支持配置生成
	ConfigCompare *bool `json:"ConfigCompare,omitempty" xml:"ConfigCompare,omitempty"`
	// 设备形态属性列表
	AttributeList []*ListDeviceFormsResponseBodyDeviceFormsAttributeList `json:"AttributeList,omitempty" xml:"AttributeList,omitempty" type:"Repeated"`
	// 是否需要账号配置
	AccountConfig *bool `json:"AccountConfig,omitempty" xml:"AccountConfig,omitempty"`
	// 是否展示详情
	DetailDisplay *bool `json:"DetailDisplay,omitempty" xml:"DetailDisplay,omitempty"`
	// 设备形态是否内置
	FormBuiltIn *bool `json:"FormBuiltIn,omitempty" xml:"FormBuiltIn,omitempty"`
	// 设备形态主键
	UniqueKey *string `json:"UniqueKey,omitempty" xml:"UniqueKey,omitempty"`
	// 设备形态ID
	DeviceFormId *string `json:"DeviceFormId,omitempty" xml:"DeviceFormId,omitempty"`
	// 设备形态名称
	DeviceFormName *string `json:"DeviceFormName,omitempty" xml:"DeviceFormName,omitempty"`
}

func (s ListDeviceFormsResponseBodyDeviceForms) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceFormsResponseBodyDeviceForms) GoString() string {
	return s.String()
}

func (s *ListDeviceFormsResponseBodyDeviceForms) SetConfigCompare(v bool) *ListDeviceFormsResponseBodyDeviceForms {
	s.ConfigCompare = &v
	return s
}

func (s *ListDeviceFormsResponseBodyDeviceForms) SetAttributeList(v []*ListDeviceFormsResponseBodyDeviceFormsAttributeList) *ListDeviceFormsResponseBodyDeviceForms {
	s.AttributeList = v
	return s
}

func (s *ListDeviceFormsResponseBodyDeviceForms) SetAccountConfig(v bool) *ListDeviceFormsResponseBodyDeviceForms {
	s.AccountConfig = &v
	return s
}

func (s *ListDeviceFormsResponseBodyDeviceForms) SetDetailDisplay(v bool) *ListDeviceFormsResponseBodyDeviceForms {
	s.DetailDisplay = &v
	return s
}

func (s *ListDeviceFormsResponseBodyDeviceForms) SetFormBuiltIn(v bool) *ListDeviceFormsResponseBodyDeviceForms {
	s.FormBuiltIn = &v
	return s
}

func (s *ListDeviceFormsResponseBodyDeviceForms) SetUniqueKey(v string) *ListDeviceFormsResponseBodyDeviceForms {
	s.UniqueKey = &v
	return s
}

func (s *ListDeviceFormsResponseBodyDeviceForms) SetDeviceFormId(v string) *ListDeviceFormsResponseBodyDeviceForms {
	s.DeviceFormId = &v
	return s
}

func (s *ListDeviceFormsResponseBodyDeviceForms) SetDeviceFormName(v string) *ListDeviceFormsResponseBodyDeviceForms {
	s.DeviceFormName = &v
	return s
}

type ListDeviceFormsResponseBodyDeviceFormsAttributeList struct {
	// 设备形态属性主键
	AttributeKey *string `json:"AttributeKey,omitempty" xml:"AttributeKey,omitempty"`
	// 设备形态属性名称
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	// 设备形态属性是否必填
	AttributeRequirement *bool `json:"AttributeRequirement,omitempty" xml:"AttributeRequirement,omitempty"`
	// 设备形态属性是否唯一
	AttributeUniqueness *bool `json:"AttributeUniqueness,omitempty" xml:"AttributeUniqueness,omitempty"`
	// 设备形态属性值格式
	AttributeFormat *string `json:"AttributeFormat,omitempty" xml:"AttributeFormat,omitempty"`
	// 设备形态属性值类型
	AttributeType *string `json:"AttributeType,omitempty" xml:"AttributeType,omitempty"`
	// 设备形态属性关联对象
	AttributeReference *string `json:"AttributeReference,omitempty" xml:"AttributeReference,omitempty"`
	// 设备形态属性是否表格可见
	AttributeTableDisplay *bool `json:"AttributeTableDisplay,omitempty" xml:"AttributeTableDisplay,omitempty"`
	// 前端查询控件占位符
	AttributePlaceholder *string `json:"AttributePlaceholder,omitempty" xml:"AttributePlaceholder,omitempty"`
	// 前端是否展示对应的查询控件
	AttributeQuery *bool `json:"AttributeQuery,omitempty" xml:"AttributeQuery,omitempty"`
	// 前端查询控件是否支持模糊搜索
	AttributeFuzzyQuery *bool `json:"AttributeFuzzyQuery,omitempty" xml:"AttributeFuzzyQuery,omitempty"`
	// 设备形态属性是否内置
	AttributeBuiltIn *bool `json:"AttributeBuiltIn,omitempty" xml:"AttributeBuiltIn,omitempty"`
}

func (s ListDeviceFormsResponseBodyDeviceFormsAttributeList) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceFormsResponseBodyDeviceFormsAttributeList) GoString() string {
	return s.String()
}

func (s *ListDeviceFormsResponseBodyDeviceFormsAttributeList) SetAttributeKey(v string) *ListDeviceFormsResponseBodyDeviceFormsAttributeList {
	s.AttributeKey = &v
	return s
}

func (s *ListDeviceFormsResponseBodyDeviceFormsAttributeList) SetAttributeName(v string) *ListDeviceFormsResponseBodyDeviceFormsAttributeList {
	s.AttributeName = &v
	return s
}

func (s *ListDeviceFormsResponseBodyDeviceFormsAttributeList) SetAttributeRequirement(v bool) *ListDeviceFormsResponseBodyDeviceFormsAttributeList {
	s.AttributeRequirement = &v
	return s
}

func (s *ListDeviceFormsResponseBodyDeviceFormsAttributeList) SetAttributeUniqueness(v bool) *ListDeviceFormsResponseBodyDeviceFormsAttributeList {
	s.AttributeUniqueness = &v
	return s
}

func (s *ListDeviceFormsResponseBodyDeviceFormsAttributeList) SetAttributeFormat(v string) *ListDeviceFormsResponseBodyDeviceFormsAttributeList {
	s.AttributeFormat = &v
	return s
}

func (s *ListDeviceFormsResponseBodyDeviceFormsAttributeList) SetAttributeType(v string) *ListDeviceFormsResponseBodyDeviceFormsAttributeList {
	s.AttributeType = &v
	return s
}

func (s *ListDeviceFormsResponseBodyDeviceFormsAttributeList) SetAttributeReference(v string) *ListDeviceFormsResponseBodyDeviceFormsAttributeList {
	s.AttributeReference = &v
	return s
}

func (s *ListDeviceFormsResponseBodyDeviceFormsAttributeList) SetAttributeTableDisplay(v bool) *ListDeviceFormsResponseBodyDeviceFormsAttributeList {
	s.AttributeTableDisplay = &v
	return s
}

func (s *ListDeviceFormsResponseBodyDeviceFormsAttributeList) SetAttributePlaceholder(v string) *ListDeviceFormsResponseBodyDeviceFormsAttributeList {
	s.AttributePlaceholder = &v
	return s
}

func (s *ListDeviceFormsResponseBodyDeviceFormsAttributeList) SetAttributeQuery(v bool) *ListDeviceFormsResponseBodyDeviceFormsAttributeList {
	s.AttributeQuery = &v
	return s
}

func (s *ListDeviceFormsResponseBodyDeviceFormsAttributeList) SetAttributeFuzzyQuery(v bool) *ListDeviceFormsResponseBodyDeviceFormsAttributeList {
	s.AttributeFuzzyQuery = &v
	return s
}

func (s *ListDeviceFormsResponseBodyDeviceFormsAttributeList) SetAttributeBuiltIn(v bool) *ListDeviceFormsResponseBodyDeviceFormsAttributeList {
	s.AttributeBuiltIn = &v
	return s
}

type ListDeviceFormsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeviceFormsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeviceFormsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceFormsResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceFormsResponse) SetHeaders(v map[string]*string) *ListDeviceFormsResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceFormsResponse) SetBody(v *ListDeviceFormsResponseBody) *ListDeviceFormsResponse {
	s.Body = v
	return s
}

type GetConfigurationSpecificationRequest struct {
	// 实例 ID。
	ConfigurationSpecificationId *string `json:"ConfigurationSpecificationId,omitempty" xml:"ConfigurationSpecificationId,omitempty"`
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetConfigurationSpecificationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConfigurationSpecificationRequest) GoString() string {
	return s.String()
}

func (s *GetConfigurationSpecificationRequest) SetConfigurationSpecificationId(v string) *GetConfigurationSpecificationRequest {
	s.ConfigurationSpecificationId = &v
	return s
}

func (s *GetConfigurationSpecificationRequest) SetInstanceId(v string) *GetConfigurationSpecificationRequest {
	s.InstanceId = &v
	return s
}

type GetConfigurationSpecificationResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 配置规范对象
	ConfigurationSpecification *GetConfigurationSpecificationResponseBodyConfigurationSpecification `json:"ConfigurationSpecification,omitempty" xml:"ConfigurationSpecification,omitempty" type:"Struct"`
}

func (s GetConfigurationSpecificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConfigurationSpecificationResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfigurationSpecificationResponseBody) SetRequestId(v string) *GetConfigurationSpecificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConfigurationSpecificationResponseBody) SetConfigurationSpecification(v *GetConfigurationSpecificationResponseBodyConfigurationSpecification) *GetConfigurationSpecificationResponseBody {
	s.ConfigurationSpecification = v
	return s
}

type GetConfigurationSpecificationResponseBodyConfigurationSpecification struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 配置规范资源ID
	ConfigurationSpecificationId *string `json:"ConfigurationSpecificationId,omitempty" xml:"ConfigurationSpecificationId,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 配置规范名字
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// 架构
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 型号
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 规范内容
	SpecificationContent *string `json:"SpecificationContent,omitempty" xml:"SpecificationContent,omitempty"`
	// 相关变量
	RelatedVariate []*string `json:"RelatedVariate,omitempty" xml:"RelatedVariate,omitempty" type:"Repeated"`
}

func (s GetConfigurationSpecificationResponseBodyConfigurationSpecification) String() string {
	return tea.Prettify(s)
}

func (s GetConfigurationSpecificationResponseBodyConfigurationSpecification) GoString() string {
	return s.String()
}

func (s *GetConfigurationSpecificationResponseBodyConfigurationSpecification) SetCreateTime(v string) *GetConfigurationSpecificationResponseBodyConfigurationSpecification {
	s.CreateTime = &v
	return s
}

func (s *GetConfigurationSpecificationResponseBodyConfigurationSpecification) SetConfigurationSpecificationId(v string) *GetConfigurationSpecificationResponseBodyConfigurationSpecification {
	s.ConfigurationSpecificationId = &v
	return s
}

func (s *GetConfigurationSpecificationResponseBodyConfigurationSpecification) SetUpdateTime(v string) *GetConfigurationSpecificationResponseBodyConfigurationSpecification {
	s.UpdateTime = &v
	return s
}

func (s *GetConfigurationSpecificationResponseBodyConfigurationSpecification) SetSpecificationName(v string) *GetConfigurationSpecificationResponseBodyConfigurationSpecification {
	s.SpecificationName = &v
	return s
}

func (s *GetConfigurationSpecificationResponseBodyConfigurationSpecification) SetArchitecture(v string) *GetConfigurationSpecificationResponseBodyConfigurationSpecification {
	s.Architecture = &v
	return s
}

func (s *GetConfigurationSpecificationResponseBodyConfigurationSpecification) SetRole(v string) *GetConfigurationSpecificationResponseBodyConfigurationSpecification {
	s.Role = &v
	return s
}

func (s *GetConfigurationSpecificationResponseBodyConfigurationSpecification) SetMode(v string) *GetConfigurationSpecificationResponseBodyConfigurationSpecification {
	s.Mode = &v
	return s
}

func (s *GetConfigurationSpecificationResponseBodyConfigurationSpecification) SetVendor(v string) *GetConfigurationSpecificationResponseBodyConfigurationSpecification {
	s.Vendor = &v
	return s
}

func (s *GetConfigurationSpecificationResponseBodyConfigurationSpecification) SetSpecificationContent(v string) *GetConfigurationSpecificationResponseBodyConfigurationSpecification {
	s.SpecificationContent = &v
	return s
}

func (s *GetConfigurationSpecificationResponseBodyConfigurationSpecification) SetRelatedVariate(v []*string) *GetConfigurationSpecificationResponseBodyConfigurationSpecification {
	s.RelatedVariate = v
	return s
}

type GetConfigurationSpecificationResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetConfigurationSpecificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConfigurationSpecificationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConfigurationSpecificationResponse) GoString() string {
	return s.String()
}

func (s *GetConfigurationSpecificationResponse) SetHeaders(v map[string]*string) *GetConfigurationSpecificationResponse {
	s.Headers = v
	return s
}

func (s *GetConfigurationSpecificationResponse) SetBody(v *GetConfigurationSpecificationResponseBody) *GetConfigurationSpecificationResponse {
	s.Body = v
	return s
}

type DeleteScheduleDutyRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例 ID。
	ScheduleDutyId *string `json:"ScheduleDutyId,omitempty" xml:"ScheduleDutyId,omitempty"`
	// typeWorkerList
	TypeWorkerList []*DeleteScheduleDutyRequestTypeWorkerList `json:"TypeWorkerList,omitempty" xml:"TypeWorkerList,omitempty" type:"Repeated"`
}

func (s DeleteScheduleDutyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleDutyRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduleDutyRequest) SetInstanceId(v string) *DeleteScheduleDutyRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteScheduleDutyRequest) SetScheduleDutyId(v string) *DeleteScheduleDutyRequest {
	s.ScheduleDutyId = &v
	return s
}

func (s *DeleteScheduleDutyRequest) SetTypeWorkerList(v []*DeleteScheduleDutyRequestTypeWorkerList) *DeleteScheduleDutyRequest {
	s.TypeWorkerList = v
	return s
}

type DeleteScheduleDutyRequestTypeWorkerList struct {
	ScheduleTypeId     *string `json:"ScheduleTypeId,omitempty" xml:"ScheduleTypeId,omitempty"`
	ScheduleTypeName   *string `json:"ScheduleTypeName,omitempty" xml:"ScheduleTypeName,omitempty"`
	ScheduleWorkerName *string `json:"ScheduleWorkerName,omitempty" xml:"ScheduleWorkerName,omitempty"`
}

func (s DeleteScheduleDutyRequestTypeWorkerList) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleDutyRequestTypeWorkerList) GoString() string {
	return s.String()
}

func (s *DeleteScheduleDutyRequestTypeWorkerList) SetScheduleTypeId(v string) *DeleteScheduleDutyRequestTypeWorkerList {
	s.ScheduleTypeId = &v
	return s
}

func (s *DeleteScheduleDutyRequestTypeWorkerList) SetScheduleTypeName(v string) *DeleteScheduleDutyRequestTypeWorkerList {
	s.ScheduleTypeName = &v
	return s
}

func (s *DeleteScheduleDutyRequestTypeWorkerList) SetScheduleWorkerName(v string) *DeleteScheduleDutyRequestTypeWorkerList {
	s.ScheduleWorkerName = &v
	return s
}

type DeleteScheduleDutyShrinkRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例 ID。
	ScheduleDutyId *string `json:"ScheduleDutyId,omitempty" xml:"ScheduleDutyId,omitempty"`
	// typeWorkerList
	TypeWorkerListShrink *string `json:"TypeWorkerList,omitempty" xml:"TypeWorkerList,omitempty"`
}

func (s DeleteScheduleDutyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleDutyShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduleDutyShrinkRequest) SetInstanceId(v string) *DeleteScheduleDutyShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteScheduleDutyShrinkRequest) SetScheduleDutyId(v string) *DeleteScheduleDutyShrinkRequest {
	s.ScheduleDutyId = &v
	return s
}

func (s *DeleteScheduleDutyShrinkRequest) SetTypeWorkerListShrink(v string) *DeleteScheduleDutyShrinkRequest {
	s.TypeWorkerListShrink = &v
	return s
}

type DeleteScheduleDutyResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScheduleDutyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleDutyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduleDutyResponseBody) SetRequestId(v string) *DeleteScheduleDutyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteScheduleDutyResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteScheduleDutyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteScheduleDutyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleDutyResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduleDutyResponse) SetHeaders(v map[string]*string) *DeleteScheduleDutyResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduleDutyResponse) SetBody(v *DeleteScheduleDutyResponseBody) *DeleteScheduleDutyResponse {
	s.Body = v
	return s
}

type UploadScheduleDutyRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// scheduleDuty
	ScheduleDuty []*UploadScheduleDutyRequestScheduleDuty `json:"ScheduleDuty,omitempty" xml:"ScheduleDuty,omitempty" type:"Repeated"`
}

func (s UploadScheduleDutyRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadScheduleDutyRequest) GoString() string {
	return s.String()
}

func (s *UploadScheduleDutyRequest) SetInstanceId(v string) *UploadScheduleDutyRequest {
	s.InstanceId = &v
	return s
}

func (s *UploadScheduleDutyRequest) SetScheduleDuty(v []*UploadScheduleDutyRequestScheduleDuty) *UploadScheduleDutyRequest {
	s.ScheduleDuty = v
	return s
}

type UploadScheduleDutyRequestScheduleDuty struct {
	// 值班表日期
	WorkDate *string `json:"WorkDate,omitempty" xml:"WorkDate,omitempty"`
	// worker
	Worker []*UploadScheduleDutyRequestScheduleDutyWorker `json:"Worker,omitempty" xml:"Worker,omitempty" type:"Repeated"`
}

func (s UploadScheduleDutyRequestScheduleDuty) String() string {
	return tea.Prettify(s)
}

func (s UploadScheduleDutyRequestScheduleDuty) GoString() string {
	return s.String()
}

func (s *UploadScheduleDutyRequestScheduleDuty) SetWorkDate(v string) *UploadScheduleDutyRequestScheduleDuty {
	s.WorkDate = &v
	return s
}

func (s *UploadScheduleDutyRequestScheduleDuty) SetWorker(v []*UploadScheduleDutyRequestScheduleDutyWorker) *UploadScheduleDutyRequestScheduleDuty {
	s.Worker = v
	return s
}

type UploadScheduleDutyRequestScheduleDutyWorker struct {
	// 值班表类型
	WorkType *string `json:"WorkType,omitempty" xml:"WorkType,omitempty"`
	// 值班人员姓名
	WorkerName []*string `json:"WorkerName,omitempty" xml:"WorkerName,omitempty" type:"Repeated"`
}

func (s UploadScheduleDutyRequestScheduleDutyWorker) String() string {
	return tea.Prettify(s)
}

func (s UploadScheduleDutyRequestScheduleDutyWorker) GoString() string {
	return s.String()
}

func (s *UploadScheduleDutyRequestScheduleDutyWorker) SetWorkType(v string) *UploadScheduleDutyRequestScheduleDutyWorker {
	s.WorkType = &v
	return s
}

func (s *UploadScheduleDutyRequestScheduleDutyWorker) SetWorkerName(v []*string) *UploadScheduleDutyRequestScheduleDutyWorker {
	s.WorkerName = v
	return s
}

type UploadScheduleDutyShrinkRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// scheduleDuty
	ScheduleDutyShrink *string `json:"ScheduleDuty,omitempty" xml:"ScheduleDuty,omitempty"`
}

func (s UploadScheduleDutyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadScheduleDutyShrinkRequest) GoString() string {
	return s.String()
}

func (s *UploadScheduleDutyShrinkRequest) SetInstanceId(v string) *UploadScheduleDutyShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UploadScheduleDutyShrinkRequest) SetScheduleDutyShrink(v string) *UploadScheduleDutyShrinkRequest {
	s.ScheduleDutyShrink = &v
	return s
}

type UploadScheduleDutyResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadScheduleDutyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadScheduleDutyResponseBody) GoString() string {
	return s.String()
}

func (s *UploadScheduleDutyResponseBody) SetRequestId(v string) *UploadScheduleDutyResponseBody {
	s.RequestId = &v
	return s
}

type UploadScheduleDutyResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadScheduleDutyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadScheduleDutyResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadScheduleDutyResponse) GoString() string {
	return s.String()
}

func (s *UploadScheduleDutyResponse) SetHeaders(v map[string]*string) *UploadScheduleDutyResponse {
	s.Headers = v
	return s
}

func (s *UploadScheduleDutyResponse) SetBody(v *UploadScheduleDutyResponseBody) *UploadScheduleDutyResponse {
	s.Body = v
	return s
}

type ListAlarmStatusHistoriesRequest struct {
	// 开始时间秒级时间戳
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// 结束时间秒级时间戳
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 监控项ID
	MonitorItemId *string `json:"MonitorItemId,omitempty" xml:"MonitorItemId,omitempty"`
	// 类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 聚合数据ID
	AggregateDataId *string `json:"AggregateDataId,omitempty" xml:"AggregateDataId,omitempty"`
	// 专线ID
	DedicatedLineId *string `json:"DedicatedLineId,omitempty" xml:"DedicatedLineId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 端口集ID
	PortCollectionId *string `json:"PortCollectionId,omitempty" xml:"PortCollectionId,omitempty"`
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListAlarmStatusHistoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmStatusHistoriesRequest) GoString() string {
	return s.String()
}

func (s *ListAlarmStatusHistoriesRequest) SetStart(v int64) *ListAlarmStatusHistoriesRequest {
	s.Start = &v
	return s
}

func (s *ListAlarmStatusHistoriesRequest) SetEnd(v int64) *ListAlarmStatusHistoriesRequest {
	s.End = &v
	return s
}

func (s *ListAlarmStatusHistoriesRequest) SetDeviceId(v string) *ListAlarmStatusHistoriesRequest {
	s.DeviceId = &v
	return s
}

func (s *ListAlarmStatusHistoriesRequest) SetMonitorItemId(v string) *ListAlarmStatusHistoriesRequest {
	s.MonitorItemId = &v
	return s
}

func (s *ListAlarmStatusHistoriesRequest) SetType(v string) *ListAlarmStatusHistoriesRequest {
	s.Type = &v
	return s
}

func (s *ListAlarmStatusHistoriesRequest) SetAggregateDataId(v string) *ListAlarmStatusHistoriesRequest {
	s.AggregateDataId = &v
	return s
}

func (s *ListAlarmStatusHistoriesRequest) SetDedicatedLineId(v string) *ListAlarmStatusHistoriesRequest {
	s.DedicatedLineId = &v
	return s
}

func (s *ListAlarmStatusHistoriesRequest) SetInstanceId(v string) *ListAlarmStatusHistoriesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAlarmStatusHistoriesRequest) SetPortCollectionId(v string) *ListAlarmStatusHistoriesRequest {
	s.PortCollectionId = &v
	return s
}

func (s *ListAlarmStatusHistoriesRequest) SetAppId(v string) *ListAlarmStatusHistoriesRequest {
	s.AppId = &v
	return s
}

type ListAlarmStatusHistoriesResponseBody struct {
	// request id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 数据列表
	AlarmStatusHistories []*ListAlarmStatusHistoriesResponseBodyAlarmStatusHistories `json:"AlarmStatusHistories,omitempty" xml:"AlarmStatusHistories,omitempty" type:"Repeated"`
}

func (s ListAlarmStatusHistoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmStatusHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlarmStatusHistoriesResponseBody) SetRequestId(v string) *ListAlarmStatusHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlarmStatusHistoriesResponseBody) SetAlarmStatusHistories(v []*ListAlarmStatusHistoriesResponseBodyAlarmStatusHistories) *ListAlarmStatusHistoriesResponseBody {
	s.AlarmStatusHistories = v
	return s
}

type ListAlarmStatusHistoriesResponseBodyAlarmStatusHistories struct {
	// 时间戳
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// 数值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAlarmStatusHistoriesResponseBodyAlarmStatusHistories) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmStatusHistoriesResponseBodyAlarmStatusHistories) GoString() string {
	return s.String()
}

func (s *ListAlarmStatusHistoriesResponseBodyAlarmStatusHistories) SetTimestamp(v int64) *ListAlarmStatusHistoriesResponseBodyAlarmStatusHistories {
	s.Timestamp = &v
	return s
}

func (s *ListAlarmStatusHistoriesResponseBodyAlarmStatusHistories) SetValue(v string) *ListAlarmStatusHistoriesResponseBodyAlarmStatusHistories {
	s.Value = &v
	return s
}

type ListAlarmStatusHistoriesResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAlarmStatusHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlarmStatusHistoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmStatusHistoriesResponse) GoString() string {
	return s.String()
}

func (s *ListAlarmStatusHistoriesResponse) SetHeaders(v map[string]*string) *ListAlarmStatusHistoriesResponse {
	s.Headers = v
	return s
}

func (s *ListAlarmStatusHistoriesResponse) SetBody(v *ListAlarmStatusHistoriesResponseBody) *ListAlarmStatusHistoriesResponse {
	s.Body = v
	return s
}

type GetSpaceModelRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例 ID。
	SpaceModelId *string `json:"SpaceModelId,omitempty" xml:"SpaceModelId,omitempty"`
}

func (s GetSpaceModelRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceModelRequest) GoString() string {
	return s.String()
}

func (s *GetSpaceModelRequest) SetInstanceId(v string) *GetSpaceModelRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSpaceModelRequest) SetSpaceModelId(v string) *GetSpaceModelRequest {
	s.SpaceModelId = &v
	return s
}

type GetSpaceModelResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 物理空间模型
	SpaceModel *GetSpaceModelResponseBodySpaceModel `json:"SpaceModel,omitempty" xml:"SpaceModel,omitempty" type:"Struct"`
}

func (s GetSpaceModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceModelResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpaceModelResponseBody) SetRequestId(v string) *GetSpaceModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSpaceModelResponseBody) SetSpaceModel(v *GetSpaceModelResponseBodySpaceModel) *GetSpaceModelResponseBody {
	s.SpaceModel = v
	return s
}

type GetSpaceModelResponseBodySpaceModel struct {
	// 物理空间模型id
	SpaceModelId *string `json:"SpaceModelId,omitempty" xml:"SpaceModelId,omitempty"`
	// 模型状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 物理空间类型
	SpaceType *string `json:"SpaceType,omitempty" xml:"SpaceType,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 模型实例
	Sort []*GetSpaceModelResponseBodySpaceModelSort `json:"Sort,omitempty" xml:"Sort,omitempty" type:"Repeated"`
}

func (s GetSpaceModelResponseBodySpaceModel) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceModelResponseBodySpaceModel) GoString() string {
	return s.String()
}

func (s *GetSpaceModelResponseBodySpaceModel) SetSpaceModelId(v string) *GetSpaceModelResponseBodySpaceModel {
	s.SpaceModelId = &v
	return s
}

func (s *GetSpaceModelResponseBodySpaceModel) SetStatus(v string) *GetSpaceModelResponseBodySpaceModel {
	s.Status = &v
	return s
}

func (s *GetSpaceModelResponseBodySpaceModel) SetSpaceType(v string) *GetSpaceModelResponseBodySpaceModel {
	s.SpaceType = &v
	return s
}

func (s *GetSpaceModelResponseBodySpaceModel) SetCreateTime(v string) *GetSpaceModelResponseBodySpaceModel {
	s.CreateTime = &v
	return s
}

func (s *GetSpaceModelResponseBodySpaceModel) SetUpdateTime(v string) *GetSpaceModelResponseBodySpaceModel {
	s.UpdateTime = &v
	return s
}

func (s *GetSpaceModelResponseBodySpaceModel) SetSort(v []*GetSpaceModelResponseBodySpaceModelSort) *GetSpaceModelResponseBodySpaceModel {
	s.Sort = v
	return s
}

type GetSpaceModelResponseBodySpaceModelSort struct {
	// 层级名称
	LevelName *string `json:"LevelName,omitempty" xml:"LevelName,omitempty"`
	// 层级
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s GetSpaceModelResponseBodySpaceModelSort) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceModelResponseBodySpaceModelSort) GoString() string {
	return s.String()
}

func (s *GetSpaceModelResponseBodySpaceModelSort) SetLevelName(v string) *GetSpaceModelResponseBodySpaceModelSort {
	s.LevelName = &v
	return s
}

func (s *GetSpaceModelResponseBodySpaceModelSort) SetLevel(v int64) *GetSpaceModelResponseBodySpaceModelSort {
	s.Level = &v
	return s
}

type GetSpaceModelResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSpaceModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSpaceModelResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceModelResponse) GoString() string {
	return s.String()
}

func (s *GetSpaceModelResponse) SetHeaders(v map[string]*string) *GetSpaceModelResponse {
	s.Headers = v
	return s
}

func (s *GetSpaceModelResponse) SetBody(v *GetSpaceModelResponseBody) *GetSpaceModelResponse {
	s.Body = v
	return s
}

type CreateDeviceFormRequest struct {
	// 设备形态名称
	DeviceFormName *string `json:"DeviceFormName,omitempty" xml:"DeviceFormName,omitempty"`
	// 是否支持配置生成
	ConfigCompare *bool `json:"ConfigCompare,omitempty" xml:"ConfigCompare,omitempty"`
	// 是否需要账号配置
	AccountConfig *bool `json:"AccountConfig,omitempty" xml:"AccountConfig,omitempty"`
	// 是否展示设备详情
	DetailDisplay *bool `json:"DetailDisplay,omitempty" xml:"DetailDisplay,omitempty"`
	// 设备形态的主键
	UniqueKey *string `json:"UniqueKey,omitempty" xml:"UniqueKey,omitempty"`
	// 幂等校验 token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateDeviceFormRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceFormRequest) GoString() string {
	return s.String()
}

func (s *CreateDeviceFormRequest) SetDeviceFormName(v string) *CreateDeviceFormRequest {
	s.DeviceFormName = &v
	return s
}

func (s *CreateDeviceFormRequest) SetConfigCompare(v bool) *CreateDeviceFormRequest {
	s.ConfigCompare = &v
	return s
}

func (s *CreateDeviceFormRequest) SetAccountConfig(v bool) *CreateDeviceFormRequest {
	s.AccountConfig = &v
	return s
}

func (s *CreateDeviceFormRequest) SetDetailDisplay(v bool) *CreateDeviceFormRequest {
	s.DetailDisplay = &v
	return s
}

func (s *CreateDeviceFormRequest) SetUniqueKey(v string) *CreateDeviceFormRequest {
	s.UniqueKey = &v
	return s
}

func (s *CreateDeviceFormRequest) SetClientToken(v string) *CreateDeviceFormRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDeviceFormRequest) SetInstanceId(v string) *CreateDeviceFormRequest {
	s.InstanceId = &v
	return s
}

type CreateDeviceFormResponseBody struct {
	// 资源实例ID，如ECS实例的创建接口CreateInstance应返回InstanceId。
	DeviceFormId *string `json:"DeviceFormId,omitempty" xml:"DeviceFormId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDeviceFormResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceFormResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeviceFormResponseBody) SetDeviceFormId(v string) *CreateDeviceFormResponseBody {
	s.DeviceFormId = &v
	return s
}

func (s *CreateDeviceFormResponseBody) SetRequestId(v string) *CreateDeviceFormResponseBody {
	s.RequestId = &v
	return s
}

type CreateDeviceFormResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDeviceFormResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDeviceFormResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceFormResponse) GoString() string {
	return s.String()
}

func (s *CreateDeviceFormResponse) SetHeaders(v map[string]*string) *CreateDeviceFormResponse {
	s.Headers = v
	return s
}

func (s *CreateDeviceFormResponse) SetBody(v *CreateDeviceFormResponseBody) *CreateDeviceFormResponse {
	s.Body = v
	return s
}

type ListPhysicalSpacesRequest struct {
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 物理空间ID
	PhysicalSpaceIds []*string `json:"PhysicalSpaceIds,omitempty" xml:"PhysicalSpaceIds,omitempty" type:"Repeated"`
	// 物理空间名称，支持模糊搜索。
	PhysicalSpaceName *string `json:"PhysicalSpaceName,omitempty" xml:"PhysicalSpaceName,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListPhysicalSpacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPhysicalSpacesRequest) GoString() string {
	return s.String()
}

func (s *ListPhysicalSpacesRequest) SetMaxResults(v int32) *ListPhysicalSpacesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPhysicalSpacesRequest) SetNextToken(v string) *ListPhysicalSpacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPhysicalSpacesRequest) SetPhysicalSpaceIds(v []*string) *ListPhysicalSpacesRequest {
	s.PhysicalSpaceIds = v
	return s
}

func (s *ListPhysicalSpacesRequest) SetPhysicalSpaceName(v string) *ListPhysicalSpacesRequest {
	s.PhysicalSpaceName = &v
	return s
}

func (s *ListPhysicalSpacesRequest) SetInstanceId(v string) *ListPhysicalSpacesRequest {
	s.InstanceId = &v
	return s
}

type ListPhysicalSpacesResponseBody struct {
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 数组，返回示例目录。
	PhysicalSpaces []*ListPhysicalSpacesResponseBodyPhysicalSpaces `json:"PhysicalSpaces,omitempty" xml:"PhysicalSpaces,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 每页数量。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListPhysicalSpacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPhysicalSpacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPhysicalSpacesResponseBody) SetNextToken(v int32) *ListPhysicalSpacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPhysicalSpacesResponseBody) SetPhysicalSpaces(v []*ListPhysicalSpacesResponseBodyPhysicalSpaces) *ListPhysicalSpacesResponseBody {
	s.PhysicalSpaces = v
	return s
}

func (s *ListPhysicalSpacesResponseBody) SetRequestId(v string) *ListPhysicalSpacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPhysicalSpacesResponseBody) SetTotalCount(v int32) *ListPhysicalSpacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPhysicalSpacesResponseBody) SetMaxResults(v int32) *ListPhysicalSpacesResponseBody {
	s.MaxResults = &v
	return s
}

type ListPhysicalSpacesResponseBodyPhysicalSpaces struct {
	// 物理空间ID
	PhysicalSpaceId *string `json:"PhysicalSpaceId,omitempty" xml:"PhysicalSpaceId,omitempty"`
	// 物理空间名称
	PhysicalSpaceName *string `json:"PhysicalSpaceName,omitempty" xml:"PhysicalSpaceName,omitempty"`
	// 所属国家
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// 所属省份
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// 所属城市
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// 具体地址
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// 缩写
	SpaceAbbreviation *string `json:"SpaceAbbreviation,omitempty" xml:"SpaceAbbreviation,omitempty"`
	// 模型
	SpaceType *string `json:"SpaceType,omitempty" xml:"SpaceType,omitempty"`
	// 负责人
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// 实例
	Instance *string `json:"Instance,omitempty" xml:"Instance,omitempty"`
}

func (s ListPhysicalSpacesResponseBodyPhysicalSpaces) String() string {
	return tea.Prettify(s)
}

func (s ListPhysicalSpacesResponseBodyPhysicalSpaces) GoString() string {
	return s.String()
}

func (s *ListPhysicalSpacesResponseBodyPhysicalSpaces) SetPhysicalSpaceId(v string) *ListPhysicalSpacesResponseBodyPhysicalSpaces {
	s.PhysicalSpaceId = &v
	return s
}

func (s *ListPhysicalSpacesResponseBodyPhysicalSpaces) SetPhysicalSpaceName(v string) *ListPhysicalSpacesResponseBodyPhysicalSpaces {
	s.PhysicalSpaceName = &v
	return s
}

func (s *ListPhysicalSpacesResponseBodyPhysicalSpaces) SetCountry(v string) *ListPhysicalSpacesResponseBodyPhysicalSpaces {
	s.Country = &v
	return s
}

func (s *ListPhysicalSpacesResponseBodyPhysicalSpaces) SetProvince(v string) *ListPhysicalSpacesResponseBodyPhysicalSpaces {
	s.Province = &v
	return s
}

func (s *ListPhysicalSpacesResponseBodyPhysicalSpaces) SetCity(v string) *ListPhysicalSpacesResponseBodyPhysicalSpaces {
	s.City = &v
	return s
}

func (s *ListPhysicalSpacesResponseBodyPhysicalSpaces) SetAddress(v string) *ListPhysicalSpacesResponseBodyPhysicalSpaces {
	s.Address = &v
	return s
}

func (s *ListPhysicalSpacesResponseBodyPhysicalSpaces) SetSpaceAbbreviation(v string) *ListPhysicalSpacesResponseBodyPhysicalSpaces {
	s.SpaceAbbreviation = &v
	return s
}

func (s *ListPhysicalSpacesResponseBodyPhysicalSpaces) SetSpaceType(v string) *ListPhysicalSpacesResponseBodyPhysicalSpaces {
	s.SpaceType = &v
	return s
}

func (s *ListPhysicalSpacesResponseBodyPhysicalSpaces) SetOwner(v string) *ListPhysicalSpacesResponseBodyPhysicalSpaces {
	s.Owner = &v
	return s
}

func (s *ListPhysicalSpacesResponseBodyPhysicalSpaces) SetInstance(v string) *ListPhysicalSpacesResponseBodyPhysicalSpaces {
	s.Instance = &v
	return s
}

type ListPhysicalSpacesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPhysicalSpacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPhysicalSpacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPhysicalSpacesResponse) GoString() string {
	return s.String()
}

func (s *ListPhysicalSpacesResponse) SetHeaders(v map[string]*string) *ListPhysicalSpacesResponse {
	s.Headers = v
	return s
}

func (s *ListPhysicalSpacesResponse) SetBody(v *ListPhysicalSpacesResponseBody) *ListPhysicalSpacesResponse {
	s.Body = v
	return s
}

type ListMonitorDataRequest struct {
	// 开始时间
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// 结束时间
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// 数据类型
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// 数据项
	DataItem *string `json:"DataItem,omitempty" xml:"DataItem,omitempty"`
	// 监控项ID
	MonitorItemId *string `json:"MonitorItemId,omitempty" xml:"MonitorItemId,omitempty"`
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 聚合数据ID
	AggregateDataId *string `json:"AggregateDataId,omitempty" xml:"AggregateDataId,omitempty"`
	// 端口集ID
	PortCollectionId *string `json:"PortCollectionId,omitempty" xml:"PortCollectionId,omitempty"`
	// 专线ID
	DedicatedLineId *string `json:"DedicatedLineId,omitempty" xml:"DedicatedLineId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListMonitorDataRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *ListMonitorDataRequest) SetStart(v int64) *ListMonitorDataRequest {
	s.Start = &v
	return s
}

func (s *ListMonitorDataRequest) SetEnd(v int64) *ListMonitorDataRequest {
	s.End = &v
	return s
}

func (s *ListMonitorDataRequest) SetDataType(v string) *ListMonitorDataRequest {
	s.DataType = &v
	return s
}

func (s *ListMonitorDataRequest) SetDataItem(v string) *ListMonitorDataRequest {
	s.DataItem = &v
	return s
}

func (s *ListMonitorDataRequest) SetMonitorItemId(v string) *ListMonitorDataRequest {
	s.MonitorItemId = &v
	return s
}

func (s *ListMonitorDataRequest) SetDeviceId(v string) *ListMonitorDataRequest {
	s.DeviceId = &v
	return s
}

func (s *ListMonitorDataRequest) SetKey(v string) *ListMonitorDataRequest {
	s.Key = &v
	return s
}

func (s *ListMonitorDataRequest) SetAggregateDataId(v string) *ListMonitorDataRequest {
	s.AggregateDataId = &v
	return s
}

func (s *ListMonitorDataRequest) SetPortCollectionId(v string) *ListMonitorDataRequest {
	s.PortCollectionId = &v
	return s
}

func (s *ListMonitorDataRequest) SetDedicatedLineId(v string) *ListMonitorDataRequest {
	s.DedicatedLineId = &v
	return s
}

func (s *ListMonitorDataRequest) SetInstanceId(v string) *ListMonitorDataRequest {
	s.InstanceId = &v
	return s
}

func (s *ListMonitorDataRequest) SetAppId(v string) *ListMonitorDataRequest {
	s.AppId = &v
	return s
}

type ListMonitorDataResponseBody struct {
	// Request Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 数据列表
	MonitorData []*ListMonitorDataResponseBodyMonitorData `json:"MonitorData,omitempty" xml:"MonitorData,omitempty" type:"Repeated"`
}

func (s ListMonitorDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *ListMonitorDataResponseBody) SetRequestId(v string) *ListMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMonitorDataResponseBody) SetMonitorData(v []*ListMonitorDataResponseBodyMonitorData) *ListMonitorDataResponseBody {
	s.MonitorData = v
	return s
}

type ListMonitorDataResponseBodyMonitorData struct {
	// 时间戳
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// 数值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// 数据项
	DataItem *string `json:"DataItem,omitempty" xml:"DataItem,omitempty"`
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListMonitorDataResponseBodyMonitorData) String() string {
	return tea.Prettify(s)
}

func (s ListMonitorDataResponseBodyMonitorData) GoString() string {
	return s.String()
}

func (s *ListMonitorDataResponseBodyMonitorData) SetTimestamp(v int64) *ListMonitorDataResponseBodyMonitorData {
	s.Timestamp = &v
	return s
}

func (s *ListMonitorDataResponseBodyMonitorData) SetValue(v string) *ListMonitorDataResponseBodyMonitorData {
	s.Value = &v
	return s
}

func (s *ListMonitorDataResponseBodyMonitorData) SetDataItem(v string) *ListMonitorDataResponseBodyMonitorData {
	s.DataItem = &v
	return s
}

func (s *ListMonitorDataResponseBodyMonitorData) SetKey(v string) *ListMonitorDataResponseBodyMonitorData {
	s.Key = &v
	return s
}

type ListMonitorDataResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMonitorDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMonitorDataResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMonitorDataResponse) GoString() string {
	return s.String()
}

func (s *ListMonitorDataResponse) SetHeaders(v map[string]*string) *ListMonitorDataResponse {
	s.Headers = v
	return s
}

func (s *ListMonitorDataResponse) SetBody(v *ListMonitorDataResponseBody) *ListMonitorDataResponse {
	s.Body = v
	return s
}

type CreateResourceInformationRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 资源属性
	ResourceAttribute *string `json:"ResourceAttribute,omitempty" xml:"ResourceAttribute,omitempty"`
	// 架构id
	ArchitectureId *string `json:"ArchitectureId,omitempty" xml:"ArchitectureId,omitempty"`
	// 信息
	Information []*CreateResourceInformationRequestInformation `json:"Information,omitempty" xml:"Information,omitempty" type:"Repeated"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateResourceInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceInformationRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceInformationRequest) SetInstanceId(v string) *CreateResourceInformationRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateResourceInformationRequest) SetResourceType(v string) *CreateResourceInformationRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateResourceInformationRequest) SetResourceAttribute(v string) *CreateResourceInformationRequest {
	s.ResourceAttribute = &v
	return s
}

func (s *CreateResourceInformationRequest) SetArchitectureId(v string) *CreateResourceInformationRequest {
	s.ArchitectureId = &v
	return s
}

func (s *CreateResourceInformationRequest) SetInformation(v []*CreateResourceInformationRequestInformation) *CreateResourceInformationRequest {
	s.Information = v
	return s
}

func (s *CreateResourceInformationRequest) SetClientToken(v string) *CreateResourceInformationRequest {
	s.ClientToken = &v
	return s
}

type CreateResourceInformationRequestInformation struct {
	// 键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 键属性
	KeyAttribute *string `json:"KeyAttribute,omitempty" xml:"KeyAttribute,omitempty"`
	// 键动作
	KeyAction *string `json:"KeyAction,omitempty" xml:"KeyAction,omitempty"`
	// 键描述
	KeyDescription *string `json:"KeyDescription,omitempty" xml:"KeyDescription,omitempty"`
}

func (s CreateResourceInformationRequestInformation) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceInformationRequestInformation) GoString() string {
	return s.String()
}

func (s *CreateResourceInformationRequestInformation) SetKey(v string) *CreateResourceInformationRequestInformation {
	s.Key = &v
	return s
}

func (s *CreateResourceInformationRequestInformation) SetKeyAttribute(v string) *CreateResourceInformationRequestInformation {
	s.KeyAttribute = &v
	return s
}

func (s *CreateResourceInformationRequestInformation) SetKeyAction(v string) *CreateResourceInformationRequestInformation {
	s.KeyAction = &v
	return s
}

func (s *CreateResourceInformationRequestInformation) SetKeyDescription(v string) *CreateResourceInformationRequestInformation {
	s.KeyDescription = &v
	return s
}

type CreateResourceInformationResponseBody struct {
	// 资源实例ID，如ECS实例的创建接口CreateInstance应返回InstanceId。
	ResourceInformationId *string `json:"ResourceInformationId,omitempty" xml:"ResourceInformationId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateResourceInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceInformationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceInformationResponseBody) SetResourceInformationId(v string) *CreateResourceInformationResponseBody {
	s.ResourceInformationId = &v
	return s
}

func (s *CreateResourceInformationResponseBody) SetRequestId(v string) *CreateResourceInformationResponseBody {
	s.RequestId = &v
	return s
}

type CreateResourceInformationResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateResourceInformationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateResourceInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceInformationResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceInformationResponse) SetHeaders(v map[string]*string) *CreateResourceInformationResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceInformationResponse) SetBody(v *CreateResourceInformationResponseBody) *CreateResourceInformationResponse {
	s.Body = v
	return s
}

type UpdateSpaceModelInstanceRequest struct {
	// 物理空间id
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// 物理空间实例
	Instance *string `json:"Instance,omitempty" xml:"Instance,omitempty"`
}

func (s UpdateSpaceModelInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpaceModelInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateSpaceModelInstanceRequest) SetSpaceId(v string) *UpdateSpaceModelInstanceRequest {
	s.SpaceId = &v
	return s
}

func (s *UpdateSpaceModelInstanceRequest) SetInstance(v string) *UpdateSpaceModelInstanceRequest {
	s.Instance = &v
	return s
}

type UpdateSpaceModelInstanceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSpaceModelInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpaceModelInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSpaceModelInstanceResponseBody) SetRequestId(v string) *UpdateSpaceModelInstanceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSpaceModelInstanceResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSpaceModelInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSpaceModelInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpaceModelInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateSpaceModelInstanceResponse) SetHeaders(v map[string]*string) *UpdateSpaceModelInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateSpaceModelInstanceResponse) SetBody(v *UpdateSpaceModelInstanceResponseBody) *UpdateSpaceModelInstanceResponse {
	s.Body = v
	return s
}

type UpdateIpRecordRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 用于保证请求的幂等性。由客户端生成该参数值，要保证在不同请求间唯一。只支持 ASCII 字符，且不能超过 64 个字符
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 资源一级ID
	IpRecordId *string `json:"IpRecordId,omitempty" xml:"IpRecordId,omitempty"`
	// 工单状态 running complete fail
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateIpRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpRecordRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpRecordRequest) SetInstanceId(v string) *UpdateIpRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateIpRecordRequest) SetClientToken(v string) *UpdateIpRecordRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIpRecordRequest) SetIpRecordId(v string) *UpdateIpRecordRequest {
	s.IpRecordId = &v
	return s
}

func (s *UpdateIpRecordRequest) SetStatus(v string) *UpdateIpRecordRequest {
	s.Status = &v
	return s
}

type UpdateIpRecordResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpRecordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpRecordResponseBody) SetRequestId(v string) *UpdateIpRecordResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIpRecordResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateIpRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateIpRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpRecordResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpRecordResponse) SetHeaders(v map[string]*string) *UpdateIpRecordResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpRecordResponse) SetBody(v *UpdateIpRecordResponseBody) *UpdateIpRecordResponse {
	s.Body = v
	return s
}

type ReleaseIPRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// ip地址类型
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// 建设项目id
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// deviceResourceIdStr
	DeviceResourceIds []*string `json:"DeviceResourceIds,omitempty" xml:"DeviceResourceIds,omitempty" type:"Repeated"`
	// deviceResourceId
	DeviceResourceId *string `json:"DeviceResourceId,omitempty" xml:"DeviceResourceId,omitempty"`
}

func (s ReleaseIPRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseIPRequest) GoString() string {
	return s.String()
}

func (s *ReleaseIPRequest) SetInstanceId(v string) *ReleaseIPRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseIPRequest) SetIpType(v string) *ReleaseIPRequest {
	s.IpType = &v
	return s
}

func (s *ReleaseIPRequest) SetSetupProjectId(v string) *ReleaseIPRequest {
	s.SetupProjectId = &v
	return s
}

func (s *ReleaseIPRequest) SetDeviceResourceIds(v []*string) *ReleaseIPRequest {
	s.DeviceResourceIds = v
	return s
}

func (s *ReleaseIPRequest) SetDeviceResourceId(v string) *ReleaseIPRequest {
	s.DeviceResourceId = &v
	return s
}

type ReleaseIPShrinkRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// ip地址类型
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// 建设项目id
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// deviceResourceIdStr
	DeviceResourceIdsShrink *string `json:"DeviceResourceIds,omitempty" xml:"DeviceResourceIds,omitempty"`
	// deviceResourceId
	DeviceResourceId *string `json:"DeviceResourceId,omitempty" xml:"DeviceResourceId,omitempty"`
}

func (s ReleaseIPShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseIPShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReleaseIPShrinkRequest) SetInstanceId(v string) *ReleaseIPShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseIPShrinkRequest) SetIpType(v string) *ReleaseIPShrinkRequest {
	s.IpType = &v
	return s
}

func (s *ReleaseIPShrinkRequest) SetSetupProjectId(v string) *ReleaseIPShrinkRequest {
	s.SetupProjectId = &v
	return s
}

func (s *ReleaseIPShrinkRequest) SetDeviceResourceIdsShrink(v string) *ReleaseIPShrinkRequest {
	s.DeviceResourceIdsShrink = &v
	return s
}

func (s *ReleaseIPShrinkRequest) SetDeviceResourceId(v string) *ReleaseIPShrinkRequest {
	s.DeviceResourceId = &v
	return s
}

type ReleaseIPResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseIPResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseIPResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseIPResponseBody) SetRequestId(v string) *ReleaseIPResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseIPResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseIPResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseIPResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseIPResponse) GoString() string {
	return s.String()
}

func (s *ReleaseIPResponse) SetHeaders(v map[string]*string) *ReleaseIPResponse {
	s.Headers = v
	return s
}

func (s *ReleaseIPResponse) SetBody(v *ReleaseIPResponseBody) *ReleaseIPResponse {
	s.Body = v
	return s
}

type DeleteDeviceResourceRequest struct {
	// 实例 ID。
	DeviceResourceId *string `json:"DeviceResourceId,omitempty" xml:"DeviceResourceId,omitempty"`
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteDeviceResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeviceResourceRequest) SetDeviceResourceId(v string) *DeleteDeviceResourceRequest {
	s.DeviceResourceId = &v
	return s
}

func (s *DeleteDeviceResourceRequest) SetInstanceId(v string) *DeleteDeviceResourceRequest {
	s.InstanceId = &v
	return s
}

type DeleteDeviceResourceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDeviceResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeviceResourceResponseBody) SetRequestId(v string) *DeleteDeviceResourceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDeviceResourceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDeviceResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDeviceResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeviceResourceResponse) SetHeaders(v map[string]*string) *DeleteDeviceResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeviceResourceResponse) SetBody(v *DeleteDeviceResourceResponseBody) *DeleteDeviceResourceResponse {
	s.Body = v
	return s
}

type GetIpBlockRequest struct {
	// 资源ID
	IpBlockId *string `json:"IpBlockId,omitempty" xml:"IpBlockId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetIpBlockRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIpBlockRequest) GoString() string {
	return s.String()
}

func (s *GetIpBlockRequest) SetIpBlockId(v string) *GetIpBlockRequest {
	s.IpBlockId = &v
	return s
}

func (s *GetIpBlockRequest) SetInstanceId(v string) *GetIpBlockRequest {
	s.InstanceId = &v
	return s
}

type GetIpBlockResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 园区类型对象
	IpBlock *GetIpBlockResponseBodyIpBlock `json:"IpBlock,omitempty" xml:"IpBlock,omitempty" type:"Struct"`
}

func (s GetIpBlockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIpBlockResponseBody) GoString() string {
	return s.String()
}

func (s *GetIpBlockResponseBody) SetRequestId(v string) *GetIpBlockResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIpBlockResponseBody) SetIpBlock(v *GetIpBlockResponseBodyIpBlock) *GetIpBlockResponseBody {
	s.IpBlock = v
	return s
}

type GetIpBlockResponseBodyIpBlock struct {
	// IP地址段UID
	IpBlockId *string `json:"IpBlockId,omitempty" xml:"IpBlockId,omitempty"`
	// IP段
	IpBlockCode *string `json:"IpBlockCode,omitempty" xml:"IpBlockCode,omitempty"`
	// 父地址段UID
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// 公网私网标志 PUBLIC PRIVATE
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// 状态： using available lock
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// IP归属
	Ownership *string `json:"Ownership,omitempty" xml:"Ownership,omitempty"`
	// 地址类别 IPV4
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 备注
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 备份设备名称
	BackupDeviceName *string `json:"BackupDeviceName,omitempty" xml:"BackupDeviceName,omitempty"`
	// 园区层级
	ZoneLayer []*GetIpBlockResponseBodyIpBlockZoneLayer `json:"ZoneLayer,omitempty" xml:"ZoneLayer,omitempty" type:"Repeated"`
	// 业务类型UID
	BusinessTypeId *string `json:"BusinessTypeId,omitempty" xml:"BusinessTypeId,omitempty"`
	// 业务类型名称
	BusinessTypeName *string `json:"BusinessTypeName,omitempty" xml:"BusinessTypeName,omitempty"`
	// 公网地址类型 INC GUEST VIP
	NetBusiness *string `json:"NetBusiness,omitempty" xml:"NetBusiness,omitempty"`
	// IP用途
	Application *string `json:"Application,omitempty" xml:"Application,omitempty"`
	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// 园区名称
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s GetIpBlockResponseBodyIpBlock) String() string {
	return tea.Prettify(s)
}

func (s GetIpBlockResponseBodyIpBlock) GoString() string {
	return s.String()
}

func (s *GetIpBlockResponseBodyIpBlock) SetIpBlockId(v string) *GetIpBlockResponseBodyIpBlock {
	s.IpBlockId = &v
	return s
}

func (s *GetIpBlockResponseBodyIpBlock) SetIpBlockCode(v string) *GetIpBlockResponseBodyIpBlock {
	s.IpBlockCode = &v
	return s
}

func (s *GetIpBlockResponseBodyIpBlock) SetParentId(v string) *GetIpBlockResponseBodyIpBlock {
	s.ParentId = &v
	return s
}

func (s *GetIpBlockResponseBodyIpBlock) SetNetType(v string) *GetIpBlockResponseBodyIpBlock {
	s.NetType = &v
	return s
}

func (s *GetIpBlockResponseBodyIpBlock) SetStatus(v string) *GetIpBlockResponseBodyIpBlock {
	s.Status = &v
	return s
}

func (s *GetIpBlockResponseBodyIpBlock) SetOwnership(v string) *GetIpBlockResponseBodyIpBlock {
	s.Ownership = &v
	return s
}

func (s *GetIpBlockResponseBodyIpBlock) SetCategory(v string) *GetIpBlockResponseBodyIpBlock {
	s.Category = &v
	return s
}

func (s *GetIpBlockResponseBodyIpBlock) SetDescription(v string) *GetIpBlockResponseBodyIpBlock {
	s.Description = &v
	return s
}

func (s *GetIpBlockResponseBodyIpBlock) SetBackupDeviceName(v string) *GetIpBlockResponseBodyIpBlock {
	s.BackupDeviceName = &v
	return s
}

func (s *GetIpBlockResponseBodyIpBlock) SetZoneLayer(v []*GetIpBlockResponseBodyIpBlockZoneLayer) *GetIpBlockResponseBodyIpBlock {
	s.ZoneLayer = v
	return s
}

func (s *GetIpBlockResponseBodyIpBlock) SetBusinessTypeId(v string) *GetIpBlockResponseBodyIpBlock {
	s.BusinessTypeId = &v
	return s
}

func (s *GetIpBlockResponseBodyIpBlock) SetBusinessTypeName(v string) *GetIpBlockResponseBodyIpBlock {
	s.BusinessTypeName = &v
	return s
}

func (s *GetIpBlockResponseBodyIpBlock) SetNetBusiness(v string) *GetIpBlockResponseBodyIpBlock {
	s.NetBusiness = &v
	return s
}

func (s *GetIpBlockResponseBodyIpBlock) SetApplication(v string) *GetIpBlockResponseBodyIpBlock {
	s.Application = &v
	return s
}

func (s *GetIpBlockResponseBodyIpBlock) SetDeviceName(v string) *GetIpBlockResponseBodyIpBlock {
	s.DeviceName = &v
	return s
}

func (s *GetIpBlockResponseBodyIpBlock) SetZoneName(v string) *GetIpBlockResponseBodyIpBlock {
	s.ZoneName = &v
	return s
}

type GetIpBlockResponseBodyIpBlockZoneLayer struct {
	// 园区层级名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 园区层级值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetIpBlockResponseBodyIpBlockZoneLayer) String() string {
	return tea.Prettify(s)
}

func (s GetIpBlockResponseBodyIpBlockZoneLayer) GoString() string {
	return s.String()
}

func (s *GetIpBlockResponseBodyIpBlockZoneLayer) SetName(v string) *GetIpBlockResponseBodyIpBlockZoneLayer {
	s.Name = &v
	return s
}

func (s *GetIpBlockResponseBodyIpBlockZoneLayer) SetValue(v string) *GetIpBlockResponseBodyIpBlockZoneLayer {
	s.Value = &v
	return s
}

type GetIpBlockResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetIpBlockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIpBlockResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIpBlockResponse) GoString() string {
	return s.String()
}

func (s *GetIpBlockResponse) SetHeaders(v map[string]*string) *GetIpBlockResponse {
	s.Headers = v
	return s
}

func (s *GetIpBlockResponse) SetBody(v *GetIpBlockResponseBody) *GetIpBlockResponse {
	s.Body = v
	return s
}

type DeleteIpBlockRequest struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源ID
	IpBlockId *string `json:"IpBlockId,omitempty" xml:"IpBlockId,omitempty"`
}

func (s DeleteIpBlockRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpBlockRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpBlockRequest) SetInstanceId(v string) *DeleteIpBlockRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteIpBlockRequest) SetIpBlockId(v string) *DeleteIpBlockRequest {
	s.IpBlockId = &v
	return s
}

type DeleteIpBlockResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpBlockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpBlockResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpBlockResponseBody) SetRequestId(v string) *DeleteIpBlockResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIpBlockResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteIpBlockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIpBlockResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpBlockResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpBlockResponse) SetHeaders(v map[string]*string) *DeleteIpBlockResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpBlockResponse) SetBody(v *DeleteIpBlockResponseBody) *DeleteIpBlockResponse {
	s.Body = v
	return s
}

type DeleteInspectionTaskRequest struct {
	// 周期性任务的ID
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteInspectionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInspectionTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteInspectionTaskRequest) SetTaskId(v string) *DeleteInspectionTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DeleteInspectionTaskRequest) SetInstanceId(v string) *DeleteInspectionTaskRequest {
	s.InstanceId = &v
	return s
}

type DeleteInspectionTaskResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInspectionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInspectionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInspectionTaskResponseBody) SetRequestId(v string) *DeleteInspectionTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInspectionTaskResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInspectionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInspectionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInspectionTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteInspectionTaskResponse) SetHeaders(v map[string]*string) *DeleteInspectionTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteInspectionTaskResponse) SetBody(v *DeleteInspectionTaskResponseBody) *DeleteInspectionTaskResponse {
	s.Body = v
	return s
}

type DeleteConfigurationSpecificationRequest struct {
	// 实例 ID。
	ConfigurationSpecificationId *string `json:"ConfigurationSpecificationId,omitempty" xml:"ConfigurationSpecificationId,omitempty"`
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteConfigurationSpecificationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigurationSpecificationRequest) GoString() string {
	return s.String()
}

func (s *DeleteConfigurationSpecificationRequest) SetConfigurationSpecificationId(v string) *DeleteConfigurationSpecificationRequest {
	s.ConfigurationSpecificationId = &v
	return s
}

func (s *DeleteConfigurationSpecificationRequest) SetInstanceId(v string) *DeleteConfigurationSpecificationRequest {
	s.InstanceId = &v
	return s
}

type DeleteConfigurationSpecificationResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConfigurationSpecificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigurationSpecificationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConfigurationSpecificationResponseBody) SetRequestId(v string) *DeleteConfigurationSpecificationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteConfigurationSpecificationResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteConfigurationSpecificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteConfigurationSpecificationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigurationSpecificationResponse) GoString() string {
	return s.String()
}

func (s *DeleteConfigurationSpecificationResponse) SetHeaders(v map[string]*string) *DeleteConfigurationSpecificationResponse {
	s.Headers = v
	return s
}

func (s *DeleteConfigurationSpecificationResponse) SetBody(v *DeleteConfigurationSpecificationResponseBody) *DeleteConfigurationSpecificationResponse {
	s.Body = v
	return s
}

type ListBusinessTypesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 业务类型名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 保留地址数目
	ReserveNumber *int64 `json:"ReserveNumber,omitempty" xml:"ReserveNumber,omitempty"`
	// 业务类型大类
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 绑定的园区类型
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s ListBusinessTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBusinessTypesRequest) GoString() string {
	return s.String()
}

func (s *ListBusinessTypesRequest) SetInstanceId(v string) *ListBusinessTypesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListBusinessTypesRequest) SetMaxResults(v int32) *ListBusinessTypesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBusinessTypesRequest) SetNextToken(v string) *ListBusinessTypesRequest {
	s.NextToken = &v
	return s
}

func (s *ListBusinessTypesRequest) SetName(v string) *ListBusinessTypesRequest {
	s.Name = &v
	return s
}

func (s *ListBusinessTypesRequest) SetReserveNumber(v int64) *ListBusinessTypesRequest {
	s.ReserveNumber = &v
	return s
}

func (s *ListBusinessTypesRequest) SetType(v string) *ListBusinessTypesRequest {
	s.Type = &v
	return s
}

func (s *ListBusinessTypesRequest) SetZoneType(v string) *ListBusinessTypesRequest {
	s.ZoneType = &v
	return s
}

type ListBusinessTypesResponseBody struct {
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据量
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 数组，返回示例目录。
	BusinessType []*ListBusinessTypesResponseBodyBusinessType `json:"BusinessType,omitempty" xml:"BusinessType,omitempty" type:"Repeated"`
}

func (s ListBusinessTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBusinessTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBusinessTypesResponseBody) SetTotalCount(v int32) *ListBusinessTypesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBusinessTypesResponseBody) SetRequestId(v string) *ListBusinessTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBusinessTypesResponseBody) SetNextToken(v int32) *ListBusinessTypesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBusinessTypesResponseBody) SetMaxResults(v int64) *ListBusinessTypesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListBusinessTypesResponseBody) SetBusinessType(v []*ListBusinessTypesResponseBodyBusinessType) *ListBusinessTypesResponseBody {
	s.BusinessType = v
	return s
}

type ListBusinessTypesResponseBodyBusinessType struct {
	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 资源一级ID
	BusinessTypeId *string `json:"BusinessTypeId,omitempty" xml:"BusinessTypeId,omitempty"`
	// 业务类型名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 业务类型缩写
	Abbreviation *string `json:"Abbreviation,omitempty" xml:"Abbreviation,omitempty"`
	// 掩码
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// 网关地址位置，正数为正数序号，负数为倒数序号
	Gateway *int64 `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// 是否复用 reuse/single
	Sharing *string `json:"Sharing,omitempty" xml:"Sharing,omitempty"`
	// 分配方向，0表示正向，1表示反向
	Direction *int64 `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// 保留地址数目
	ReserveNumber *int64 `json:"ReserveNumber,omitempty" xml:"ReserveNumber,omitempty"`
	// 业务类型大类
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 绑定的园区类型
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
	// 有效时间
	LeaseTime *string `json:"LeaseTime,omitempty" xml:"LeaseTime,omitempty"`
	// Vlan
	Vlan *string `json:"Vlan,omitempty" xml:"Vlan,omitempty"`
	// 业务类型地址申请完对应的动作，DHCP表示需要触发DHCP变更
	ActionFlag *string `json:"ActionFlag,omitempty" xml:"ActionFlag,omitempty"`
}

func (s ListBusinessTypesResponseBodyBusinessType) String() string {
	return tea.Prettify(s)
}

func (s ListBusinessTypesResponseBodyBusinessType) GoString() string {
	return s.String()
}

func (s *ListBusinessTypesResponseBodyBusinessType) SetUpdateTime(v string) *ListBusinessTypesResponseBodyBusinessType {
	s.UpdateTime = &v
	return s
}

func (s *ListBusinessTypesResponseBodyBusinessType) SetCreateTime(v string) *ListBusinessTypesResponseBodyBusinessType {
	s.CreateTime = &v
	return s
}

func (s *ListBusinessTypesResponseBodyBusinessType) SetBusinessTypeId(v string) *ListBusinessTypesResponseBodyBusinessType {
	s.BusinessTypeId = &v
	return s
}

func (s *ListBusinessTypesResponseBodyBusinessType) SetName(v string) *ListBusinessTypesResponseBodyBusinessType {
	s.Name = &v
	return s
}

func (s *ListBusinessTypesResponseBodyBusinessType) SetAbbreviation(v string) *ListBusinessTypesResponseBodyBusinessType {
	s.Abbreviation = &v
	return s
}

func (s *ListBusinessTypesResponseBodyBusinessType) SetMask(v string) *ListBusinessTypesResponseBodyBusinessType {
	s.Mask = &v
	return s
}

func (s *ListBusinessTypesResponseBodyBusinessType) SetGateway(v int64) *ListBusinessTypesResponseBodyBusinessType {
	s.Gateway = &v
	return s
}

func (s *ListBusinessTypesResponseBodyBusinessType) SetSharing(v string) *ListBusinessTypesResponseBodyBusinessType {
	s.Sharing = &v
	return s
}

func (s *ListBusinessTypesResponseBodyBusinessType) SetDirection(v int64) *ListBusinessTypesResponseBodyBusinessType {
	s.Direction = &v
	return s
}

func (s *ListBusinessTypesResponseBodyBusinessType) SetReserveNumber(v int64) *ListBusinessTypesResponseBodyBusinessType {
	s.ReserveNumber = &v
	return s
}

func (s *ListBusinessTypesResponseBodyBusinessType) SetType(v string) *ListBusinessTypesResponseBodyBusinessType {
	s.Type = &v
	return s
}

func (s *ListBusinessTypesResponseBodyBusinessType) SetZoneType(v string) *ListBusinessTypesResponseBodyBusinessType {
	s.ZoneType = &v
	return s
}

func (s *ListBusinessTypesResponseBodyBusinessType) SetLeaseTime(v string) *ListBusinessTypesResponseBodyBusinessType {
	s.LeaseTime = &v
	return s
}

func (s *ListBusinessTypesResponseBodyBusinessType) SetVlan(v string) *ListBusinessTypesResponseBodyBusinessType {
	s.Vlan = &v
	return s
}

func (s *ListBusinessTypesResponseBodyBusinessType) SetActionFlag(v string) *ListBusinessTypesResponseBodyBusinessType {
	s.ActionFlag = &v
	return s
}

type ListBusinessTypesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBusinessTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBusinessTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBusinessTypesResponse) GoString() string {
	return s.String()
}

func (s *ListBusinessTypesResponse) SetHeaders(v map[string]*string) *ListBusinessTypesResponse {
	s.Headers = v
	return s
}

func (s *ListBusinessTypesResponse) SetBody(v *ListBusinessTypesResponseBody) *ListBusinessTypesResponse {
	s.Body = v
	return s
}

type ListSetupProjectsRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 集群名
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 物理空间id
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s ListSetupProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSetupProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListSetupProjectsRequest) SetInstanceId(v string) *ListSetupProjectsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSetupProjectsRequest) SetMaxResults(v int32) *ListSetupProjectsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSetupProjectsRequest) SetNextToken(v string) *ListSetupProjectsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSetupProjectsRequest) SetStatus(v string) *ListSetupProjectsRequest {
	s.Status = &v
	return s
}

func (s *ListSetupProjectsRequest) SetSpaceId(v string) *ListSetupProjectsRequest {
	s.SpaceId = &v
	return s
}

type ListSetupProjectsResponseBody struct {
	// 数组，返回示例目录。
	SetupProject []*ListSetupProjectsResponseBodySetupProject `json:"SetupProject,omitempty" xml:"SetupProject,omitempty" type:"Repeated"`
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据量
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListSetupProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSetupProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSetupProjectsResponseBody) SetSetupProject(v []*ListSetupProjectsResponseBodySetupProject) *ListSetupProjectsResponseBody {
	s.SetupProject = v
	return s
}

func (s *ListSetupProjectsResponseBody) SetTotalCount(v int32) *ListSetupProjectsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSetupProjectsResponseBody) SetRequestId(v string) *ListSetupProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSetupProjectsResponseBody) SetNextToken(v int32) *ListSetupProjectsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSetupProjectsResponseBody) SetMaxResults(v int64) *ListSetupProjectsResponseBody {
	s.MaxResults = &v
	return s
}

type ListSetupProjectsResponseBodySetupProject struct {
	// 项目进展
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// 物理空间类型
	SpaceType *string `json:"SpaceType,omitempty" xml:"SpaceType,omitempty"`
	// 物理空间名称
	SpaceName *string `json:"SpaceName,omitempty" xml:"SpaceName,omitempty"`
	// 架构版本
	ArchVersion *string `json:"ArchVersion,omitempty" xml:"ArchVersion,omitempty"`
	// 架构id
	ArchitectureId *string `json:"ArchitectureId,omitempty" xml:"ArchitectureId,omitempty"`
	// 套餐
	Packages []*ListSetupProjectsResponseBodySetupProjectPackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Repeated"`
	// 预计交付时间
	DeliveryTime *string `json:"DeliveryTime,omitempty" xml:"DeliveryTime,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 物理空间uId
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// 资源一级ID
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// 节点
	Nodes *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ListSetupProjectsResponseBodySetupProject) String() string {
	return tea.Prettify(s)
}

func (s ListSetupProjectsResponseBodySetupProject) GoString() string {
	return s.String()
}

func (s *ListSetupProjectsResponseBodySetupProject) SetProgress(v string) *ListSetupProjectsResponseBodySetupProject {
	s.Progress = &v
	return s
}

func (s *ListSetupProjectsResponseBodySetupProject) SetSpaceType(v string) *ListSetupProjectsResponseBodySetupProject {
	s.SpaceType = &v
	return s
}

func (s *ListSetupProjectsResponseBodySetupProject) SetSpaceName(v string) *ListSetupProjectsResponseBodySetupProject {
	s.SpaceName = &v
	return s
}

func (s *ListSetupProjectsResponseBodySetupProject) SetArchVersion(v string) *ListSetupProjectsResponseBodySetupProject {
	s.ArchVersion = &v
	return s
}

func (s *ListSetupProjectsResponseBodySetupProject) SetArchitectureId(v string) *ListSetupProjectsResponseBodySetupProject {
	s.ArchitectureId = &v
	return s
}

func (s *ListSetupProjectsResponseBodySetupProject) SetPackages(v []*ListSetupProjectsResponseBodySetupProjectPackages) *ListSetupProjectsResponseBodySetupProject {
	s.Packages = v
	return s
}

func (s *ListSetupProjectsResponseBodySetupProject) SetDeliveryTime(v string) *ListSetupProjectsResponseBodySetupProject {
	s.DeliveryTime = &v
	return s
}

func (s *ListSetupProjectsResponseBodySetupProject) SetCreateTime(v string) *ListSetupProjectsResponseBodySetupProject {
	s.CreateTime = &v
	return s
}

func (s *ListSetupProjectsResponseBodySetupProject) SetSpaceId(v string) *ListSetupProjectsResponseBodySetupProject {
	s.SpaceId = &v
	return s
}

func (s *ListSetupProjectsResponseBodySetupProject) SetSetupProjectId(v string) *ListSetupProjectsResponseBodySetupProject {
	s.SetupProjectId = &v
	return s
}

func (s *ListSetupProjectsResponseBodySetupProject) SetNodes(v string) *ListSetupProjectsResponseBodySetupProject {
	s.Nodes = &v
	return s
}

func (s *ListSetupProjectsResponseBodySetupProject) SetDescription(v string) *ListSetupProjectsResponseBodySetupProject {
	s.Description = &v
	return s
}

type ListSetupProjectsResponseBodySetupProjectPackages struct {
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 设备号
	DeviceNumber *int64 `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s ListSetupProjectsResponseBodySetupProjectPackages) String() string {
	return tea.Prettify(s)
}

func (s ListSetupProjectsResponseBodySetupProjectPackages) GoString() string {
	return s.String()
}

func (s *ListSetupProjectsResponseBodySetupProjectPackages) SetRole(v string) *ListSetupProjectsResponseBodySetupProjectPackages {
	s.Role = &v
	return s
}

func (s *ListSetupProjectsResponseBodySetupProjectPackages) SetDeviceNumber(v int64) *ListSetupProjectsResponseBodySetupProjectPackages {
	s.DeviceNumber = &v
	return s
}

func (s *ListSetupProjectsResponseBodySetupProjectPackages) SetVendor(v string) *ListSetupProjectsResponseBodySetupProjectPackages {
	s.Vendor = &v
	return s
}

func (s *ListSetupProjectsResponseBodySetupProjectPackages) SetModel(v string) *ListSetupProjectsResponseBodySetupProjectPackages {
	s.Model = &v
	return s
}

type ListSetupProjectsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSetupProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSetupProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSetupProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListSetupProjectsResponse) SetHeaders(v map[string]*string) *ListSetupProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListSetupProjectsResponse) SetBody(v *ListSetupProjectsResponseBody) *ListSetupProjectsResponse {
	s.Body = v
	return s
}

type ListTasksHistoriesRequest struct {
	// 返回结果的最大个数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 巡检项ID
	ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListTasksHistoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksHistoriesRequest) GoString() string {
	return s.String()
}

func (s *ListTasksHistoriesRequest) SetMaxResults(v int32) *ListTasksHistoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTasksHistoriesRequest) SetNextToken(v string) *ListTasksHistoriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTasksHistoriesRequest) SetItemId(v string) *ListTasksHistoriesRequest {
	s.ItemId = &v
	return s
}

func (s *ListTasksHistoriesRequest) SetDeviceId(v string) *ListTasksHistoriesRequest {
	s.DeviceId = &v
	return s
}

func (s *ListTasksHistoriesRequest) SetInstanceId(v string) *ListTasksHistoriesRequest {
	s.InstanceId = &v
	return s
}

type ListTasksHistoriesResponseBody struct {
	// 总记录数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token。
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 数组，返回示例目录。
	InspectionTasks []*ListTasksHistoriesResponseBodyInspectionTasks `json:"InspectionTasks,omitempty" xml:"InspectionTasks,omitempty" type:"Repeated"`
}

func (s ListTasksHistoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTasksHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksHistoriesResponseBody) SetTotalCount(v int32) *ListTasksHistoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTasksHistoriesResponseBody) SetRequestId(v string) *ListTasksHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksHistoriesResponseBody) SetNextToken(v int32) *ListTasksHistoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTasksHistoriesResponseBody) SetInspectionTasks(v []*ListTasksHistoriesResponseBodyInspectionTasks) *ListTasksHistoriesResponseBody {
	s.InspectionTasks = v
	return s
}

type ListTasksHistoriesResponseBodyInspectionTasks struct {
	// 巡检结束时间
	ExecutionEndTime *string `json:"ExecutionEndTime,omitempty" xml:"ExecutionEndTime,omitempty"`
	// 巡检开始时间
	ExecutionBeginTime *string `json:"ExecutionBeginTime,omitempty" xml:"ExecutionBeginTime,omitempty"`
	// 巡检结果
	InspectionResult *string `json:"InspectionResult,omitempty" xml:"InspectionResult,omitempty"`
	// 告警规则
	InspectionAlarmRules []*ListTasksHistoriesResponseBodyInspectionTasksInspectionAlarmRules `json:"InspectionAlarmRules,omitempty" xml:"InspectionAlarmRules,omitempty" type:"Repeated"`
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListTasksHistoriesResponseBodyInspectionTasks) String() string {
	return tea.Prettify(s)
}

func (s ListTasksHistoriesResponseBodyInspectionTasks) GoString() string {
	return s.String()
}

func (s *ListTasksHistoriesResponseBodyInspectionTasks) SetExecutionEndTime(v string) *ListTasksHistoriesResponseBodyInspectionTasks {
	s.ExecutionEndTime = &v
	return s
}

func (s *ListTasksHistoriesResponseBodyInspectionTasks) SetExecutionBeginTime(v string) *ListTasksHistoriesResponseBodyInspectionTasks {
	s.ExecutionBeginTime = &v
	return s
}

func (s *ListTasksHistoriesResponseBodyInspectionTasks) SetInspectionResult(v string) *ListTasksHistoriesResponseBodyInspectionTasks {
	s.InspectionResult = &v
	return s
}

func (s *ListTasksHistoriesResponseBodyInspectionTasks) SetInspectionAlarmRules(v []*ListTasksHistoriesResponseBodyInspectionTasksInspectionAlarmRules) *ListTasksHistoriesResponseBodyInspectionTasks {
	s.InspectionAlarmRules = v
	return s
}

func (s *ListTasksHistoriesResponseBodyInspectionTasks) SetTaskId(v string) *ListTasksHistoriesResponseBodyInspectionTasks {
	s.TaskId = &v
	return s
}

type ListTasksHistoriesResponseBodyInspectionTasksInspectionAlarmRules struct {
	// 告警表达式
	AlarmExpression *string `json:"AlarmExpression,omitempty" xml:"AlarmExpression,omitempty"`
	// 告警操作符
	AlarmOperator *string `json:"AlarmOperator,omitempty" xml:"AlarmOperator,omitempty"`
	// 告警值
	AlarmValue *string `json:"AlarmValue,omitempty" xml:"AlarmValue,omitempty"`
	// 告警实际值
	ActualValue *string `json:"ActualValue,omitempty" xml:"ActualValue,omitempty"`
	// 告警级别
	AlarmLevel *string `json:"AlarmLevel,omitempty" xml:"AlarmLevel,omitempty"`
}

func (s ListTasksHistoriesResponseBodyInspectionTasksInspectionAlarmRules) String() string {
	return tea.Prettify(s)
}

func (s ListTasksHistoriesResponseBodyInspectionTasksInspectionAlarmRules) GoString() string {
	return s.String()
}

func (s *ListTasksHistoriesResponseBodyInspectionTasksInspectionAlarmRules) SetAlarmExpression(v string) *ListTasksHistoriesResponseBodyInspectionTasksInspectionAlarmRules {
	s.AlarmExpression = &v
	return s
}

func (s *ListTasksHistoriesResponseBodyInspectionTasksInspectionAlarmRules) SetAlarmOperator(v string) *ListTasksHistoriesResponseBodyInspectionTasksInspectionAlarmRules {
	s.AlarmOperator = &v
	return s
}

func (s *ListTasksHistoriesResponseBodyInspectionTasksInspectionAlarmRules) SetAlarmValue(v string) *ListTasksHistoriesResponseBodyInspectionTasksInspectionAlarmRules {
	s.AlarmValue = &v
	return s
}

func (s *ListTasksHistoriesResponseBodyInspectionTasksInspectionAlarmRules) SetActualValue(v string) *ListTasksHistoriesResponseBodyInspectionTasksInspectionAlarmRules {
	s.ActualValue = &v
	return s
}

func (s *ListTasksHistoriesResponseBodyInspectionTasksInspectionAlarmRules) SetAlarmLevel(v string) *ListTasksHistoriesResponseBodyInspectionTasksInspectionAlarmRules {
	s.AlarmLevel = &v
	return s
}

type ListTasksHistoriesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTasksHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTasksHistoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTasksHistoriesResponse) GoString() string {
	return s.String()
}

func (s *ListTasksHistoriesResponse) SetHeaders(v map[string]*string) *ListTasksHistoriesResponse {
	s.Headers = v
	return s
}

func (s *ListTasksHistoriesResponse) SetBody(v *ListTasksHistoriesResponseBody) *ListTasksHistoriesResponse {
	s.Body = v
	return s
}

type UpdateConfigurationSpecificationRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 配置规范名字
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// 配置规范id
	ConfigurationSpecificationId *string `json:"ConfigurationSpecificationId,omitempty" xml:"ConfigurationSpecificationId,omitempty"`
	// 架构类型
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 配置规范内容
	SpecificationContent *string `json:"SpecificationContent,omitempty" xml:"SpecificationContent,omitempty"`
	// 相关变量
	RelatedVariate [][]byte `json:"RelatedVariate,omitempty" xml:"RelatedVariate,omitempty" type:"Repeated"`
}

func (s UpdateConfigurationSpecificationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigurationSpecificationRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigurationSpecificationRequest) SetInstanceId(v string) *UpdateConfigurationSpecificationRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateConfigurationSpecificationRequest) SetSpecificationName(v string) *UpdateConfigurationSpecificationRequest {
	s.SpecificationName = &v
	return s
}

func (s *UpdateConfigurationSpecificationRequest) SetConfigurationSpecificationId(v string) *UpdateConfigurationSpecificationRequest {
	s.ConfigurationSpecificationId = &v
	return s
}

func (s *UpdateConfigurationSpecificationRequest) SetArchitecture(v string) *UpdateConfigurationSpecificationRequest {
	s.Architecture = &v
	return s
}

func (s *UpdateConfigurationSpecificationRequest) SetRole(v string) *UpdateConfigurationSpecificationRequest {
	s.Role = &v
	return s
}

func (s *UpdateConfigurationSpecificationRequest) SetVendor(v string) *UpdateConfigurationSpecificationRequest {
	s.Vendor = &v
	return s
}

func (s *UpdateConfigurationSpecificationRequest) SetModel(v string) *UpdateConfigurationSpecificationRequest {
	s.Model = &v
	return s
}

func (s *UpdateConfigurationSpecificationRequest) SetSpecificationContent(v string) *UpdateConfigurationSpecificationRequest {
	s.SpecificationContent = &v
	return s
}

func (s *UpdateConfigurationSpecificationRequest) SetRelatedVariate(v [][]byte) *UpdateConfigurationSpecificationRequest {
	s.RelatedVariate = v
	return s
}

type UpdateConfigurationSpecificationShrinkRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 配置规范名字
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// 配置规范id
	ConfigurationSpecificationId *string `json:"ConfigurationSpecificationId,omitempty" xml:"ConfigurationSpecificationId,omitempty"`
	// 架构类型
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 配置规范内容
	SpecificationContent *string `json:"SpecificationContent,omitempty" xml:"SpecificationContent,omitempty"`
	// 相关变量
	RelatedVariateShrink *string `json:"RelatedVariate,omitempty" xml:"RelatedVariate,omitempty"`
}

func (s UpdateConfigurationSpecificationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigurationSpecificationShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigurationSpecificationShrinkRequest) SetInstanceId(v string) *UpdateConfigurationSpecificationShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateConfigurationSpecificationShrinkRequest) SetSpecificationName(v string) *UpdateConfigurationSpecificationShrinkRequest {
	s.SpecificationName = &v
	return s
}

func (s *UpdateConfigurationSpecificationShrinkRequest) SetConfigurationSpecificationId(v string) *UpdateConfigurationSpecificationShrinkRequest {
	s.ConfigurationSpecificationId = &v
	return s
}

func (s *UpdateConfigurationSpecificationShrinkRequest) SetArchitecture(v string) *UpdateConfigurationSpecificationShrinkRequest {
	s.Architecture = &v
	return s
}

func (s *UpdateConfigurationSpecificationShrinkRequest) SetRole(v string) *UpdateConfigurationSpecificationShrinkRequest {
	s.Role = &v
	return s
}

func (s *UpdateConfigurationSpecificationShrinkRequest) SetVendor(v string) *UpdateConfigurationSpecificationShrinkRequest {
	s.Vendor = &v
	return s
}

func (s *UpdateConfigurationSpecificationShrinkRequest) SetModel(v string) *UpdateConfigurationSpecificationShrinkRequest {
	s.Model = &v
	return s
}

func (s *UpdateConfigurationSpecificationShrinkRequest) SetSpecificationContent(v string) *UpdateConfigurationSpecificationShrinkRequest {
	s.SpecificationContent = &v
	return s
}

func (s *UpdateConfigurationSpecificationShrinkRequest) SetRelatedVariateShrink(v string) *UpdateConfigurationSpecificationShrinkRequest {
	s.RelatedVariateShrink = &v
	return s
}

type UpdateConfigurationSpecificationResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConfigurationSpecificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigurationSpecificationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigurationSpecificationResponseBody) SetRequestId(v string) *UpdateConfigurationSpecificationResponseBody {
	s.RequestId = &v
	return s
}

type UpdateConfigurationSpecificationResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateConfigurationSpecificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateConfigurationSpecificationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigurationSpecificationResponse) GoString() string {
	return s.String()
}

func (s *UpdateConfigurationSpecificationResponse) SetHeaders(v map[string]*string) *UpdateConfigurationSpecificationResponse {
	s.Headers = v
	return s
}

func (s *UpdateConfigurationSpecificationResponse) SetBody(v *UpdateConfigurationSpecificationResponseBody) *UpdateConfigurationSpecificationResponse {
	s.Body = v
	return s
}

type CreateOsVersionRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 系统版本
	OsVersion *string `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 文件名
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 文件路径
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateOsVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOsVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateOsVersionRequest) SetInstanceId(v string) *CreateOsVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateOsVersionRequest) SetCreateTime(v string) *CreateOsVersionRequest {
	s.CreateTime = &v
	return s
}

func (s *CreateOsVersionRequest) SetVendor(v string) *CreateOsVersionRequest {
	s.Vendor = &v
	return s
}

func (s *CreateOsVersionRequest) SetModel(v string) *CreateOsVersionRequest {
	s.Model = &v
	return s
}

func (s *CreateOsVersionRequest) SetOsVersion(v string) *CreateOsVersionRequest {
	s.OsVersion = &v
	return s
}

func (s *CreateOsVersionRequest) SetStatus(v string) *CreateOsVersionRequest {
	s.Status = &v
	return s
}

func (s *CreateOsVersionRequest) SetFileName(v string) *CreateOsVersionRequest {
	s.FileName = &v
	return s
}

func (s *CreateOsVersionRequest) SetFilePath(v string) *CreateOsVersionRequest {
	s.FilePath = &v
	return s
}

func (s *CreateOsVersionRequest) SetClientToken(v string) *CreateOsVersionRequest {
	s.ClientToken = &v
	return s
}

type CreateOsVersionResponseBody struct {
	// 资源实例ID，如ECS实例的创建接口CreateInstance应返回InstanceId。
	OsVersionId *string `json:"OsVersionId,omitempty" xml:"OsVersionId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOsVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOsVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOsVersionResponseBody) SetOsVersionId(v string) *CreateOsVersionResponseBody {
	s.OsVersionId = &v
	return s
}

func (s *CreateOsVersionResponseBody) SetRequestId(v string) *CreateOsVersionResponseBody {
	s.RequestId = &v
	return s
}

type CreateOsVersionResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateOsVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOsVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOsVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateOsVersionResponse) SetHeaders(v map[string]*string) *CreateOsVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateOsVersionResponse) SetBody(v *CreateOsVersionResponseBody) *CreateOsVersionResponse {
	s.Body = v
	return s
}

type CreateIpBlockRecordRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 工单名称
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 创建人
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// 工号
	WorkNo *string `json:"WorkNo,omitempty" xml:"WorkNo,omitempty"`
	// 公网私网标志 PUBLIC PRIVATE
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// 园区层级
	ZoneLayer []*CreateIpBlockRecordRequestZoneLayer `json:"ZoneLayer,omitempty" xml:"ZoneLayer,omitempty" type:"Repeated"`
	// 地址业务类型
	BusinessType []*CreateIpBlockRecordRequestBusinessType `json:"BusinessType,omitempty" xml:"BusinessType,omitempty" type:"Repeated"`
	// 公网地址类型 INC GUEST VIP
	NetBusiness *string `json:"NetBusiness,omitempty" xml:"NetBusiness,omitempty"`
	// 父段地址
	ParentIpBlock *string `json:"ParentIpBlock,omitempty" xml:"ParentIpBlock,omitempty"`
	// 申请公网地址的掩码大小
	Mask *int64 `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// 工单类型 Apply 申请工单 Recycle 释放工单
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// 要释放的IP段
	IpBlockCode []*string `json:"IpBlockCode,omitempty" xml:"IpBlockCode,omitempty" type:"Repeated"`
}

func (s CreateIpBlockRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIpBlockRecordRequest) GoString() string {
	return s.String()
}

func (s *CreateIpBlockRecordRequest) SetInstanceId(v string) *CreateIpBlockRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateIpBlockRecordRequest) SetClientToken(v string) *CreateIpBlockRecordRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIpBlockRecordRequest) SetTitle(v string) *CreateIpBlockRecordRequest {
	s.Title = &v
	return s
}

func (s *CreateIpBlockRecordRequest) SetCreator(v string) *CreateIpBlockRecordRequest {
	s.Creator = &v
	return s
}

func (s *CreateIpBlockRecordRequest) SetWorkNo(v string) *CreateIpBlockRecordRequest {
	s.WorkNo = &v
	return s
}

func (s *CreateIpBlockRecordRequest) SetNetType(v string) *CreateIpBlockRecordRequest {
	s.NetType = &v
	return s
}

func (s *CreateIpBlockRecordRequest) SetZoneLayer(v []*CreateIpBlockRecordRequestZoneLayer) *CreateIpBlockRecordRequest {
	s.ZoneLayer = v
	return s
}

func (s *CreateIpBlockRecordRequest) SetBusinessType(v []*CreateIpBlockRecordRequestBusinessType) *CreateIpBlockRecordRequest {
	s.BusinessType = v
	return s
}

func (s *CreateIpBlockRecordRequest) SetNetBusiness(v string) *CreateIpBlockRecordRequest {
	s.NetBusiness = &v
	return s
}

func (s *CreateIpBlockRecordRequest) SetParentIpBlock(v string) *CreateIpBlockRecordRequest {
	s.ParentIpBlock = &v
	return s
}

func (s *CreateIpBlockRecordRequest) SetMask(v int64) *CreateIpBlockRecordRequest {
	s.Mask = &v
	return s
}

func (s *CreateIpBlockRecordRequest) SetRecordType(v string) *CreateIpBlockRecordRequest {
	s.RecordType = &v
	return s
}

func (s *CreateIpBlockRecordRequest) SetIpBlockCode(v []*string) *CreateIpBlockRecordRequest {
	s.IpBlockCode = v
	return s
}

type CreateIpBlockRecordRequestZoneLayer struct {
	// 园区层级名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 园区层级值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateIpBlockRecordRequestZoneLayer) String() string {
	return tea.Prettify(s)
}

func (s CreateIpBlockRecordRequestZoneLayer) GoString() string {
	return s.String()
}

func (s *CreateIpBlockRecordRequestZoneLayer) SetName(v string) *CreateIpBlockRecordRequestZoneLayer {
	s.Name = &v
	return s
}

func (s *CreateIpBlockRecordRequestZoneLayer) SetValue(v string) *CreateIpBlockRecordRequestZoneLayer {
	s.Value = &v
	return s
}

type CreateIpBlockRecordRequestBusinessType struct {
	// 地址业务类型名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 地址业务类型数量
	Number *int64 `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s CreateIpBlockRecordRequestBusinessType) String() string {
	return tea.Prettify(s)
}

func (s CreateIpBlockRecordRequestBusinessType) GoString() string {
	return s.String()
}

func (s *CreateIpBlockRecordRequestBusinessType) SetName(v string) *CreateIpBlockRecordRequestBusinessType {
	s.Name = &v
	return s
}

func (s *CreateIpBlockRecordRequestBusinessType) SetNumber(v int64) *CreateIpBlockRecordRequestBusinessType {
	s.Number = &v
	return s
}

type CreateIpBlockRecordShrinkRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 工单名称
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 创建人
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// 工号
	WorkNo *string `json:"WorkNo,omitempty" xml:"WorkNo,omitempty"`
	// 公网私网标志 PUBLIC PRIVATE
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// 园区层级
	ZoneLayerShrink *string `json:"ZoneLayer,omitempty" xml:"ZoneLayer,omitempty"`
	// 地址业务类型
	BusinessTypeShrink *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// 公网地址类型 INC GUEST VIP
	NetBusiness *string `json:"NetBusiness,omitempty" xml:"NetBusiness,omitempty"`
	// 父段地址
	ParentIpBlock *string `json:"ParentIpBlock,omitempty" xml:"ParentIpBlock,omitempty"`
	// 申请公网地址的掩码大小
	Mask *int64 `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// 工单类型 Apply 申请工单 Recycle 释放工单
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// 要释放的IP段
	IpBlockCodeShrink *string `json:"IpBlockCode,omitempty" xml:"IpBlockCode,omitempty"`
}

func (s CreateIpBlockRecordShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIpBlockRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateIpBlockRecordShrinkRequest) SetInstanceId(v string) *CreateIpBlockRecordShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateIpBlockRecordShrinkRequest) SetClientToken(v string) *CreateIpBlockRecordShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIpBlockRecordShrinkRequest) SetTitle(v string) *CreateIpBlockRecordShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateIpBlockRecordShrinkRequest) SetCreator(v string) *CreateIpBlockRecordShrinkRequest {
	s.Creator = &v
	return s
}

func (s *CreateIpBlockRecordShrinkRequest) SetWorkNo(v string) *CreateIpBlockRecordShrinkRequest {
	s.WorkNo = &v
	return s
}

func (s *CreateIpBlockRecordShrinkRequest) SetNetType(v string) *CreateIpBlockRecordShrinkRequest {
	s.NetType = &v
	return s
}

func (s *CreateIpBlockRecordShrinkRequest) SetZoneLayerShrink(v string) *CreateIpBlockRecordShrinkRequest {
	s.ZoneLayerShrink = &v
	return s
}

func (s *CreateIpBlockRecordShrinkRequest) SetBusinessTypeShrink(v string) *CreateIpBlockRecordShrinkRequest {
	s.BusinessTypeShrink = &v
	return s
}

func (s *CreateIpBlockRecordShrinkRequest) SetNetBusiness(v string) *CreateIpBlockRecordShrinkRequest {
	s.NetBusiness = &v
	return s
}

func (s *CreateIpBlockRecordShrinkRequest) SetParentIpBlock(v string) *CreateIpBlockRecordShrinkRequest {
	s.ParentIpBlock = &v
	return s
}

func (s *CreateIpBlockRecordShrinkRequest) SetMask(v int64) *CreateIpBlockRecordShrinkRequest {
	s.Mask = &v
	return s
}

func (s *CreateIpBlockRecordShrinkRequest) SetRecordType(v string) *CreateIpBlockRecordShrinkRequest {
	s.RecordType = &v
	return s
}

func (s *CreateIpBlockRecordShrinkRequest) SetIpBlockCodeShrink(v string) *CreateIpBlockRecordShrinkRequest {
	s.IpBlockCodeShrink = &v
	return s
}

type CreateIpBlockRecordResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 资源实例ID，如ECS实例的创建接口CreateInstance应返回InstanceId。
	IpBlockRecordId *string `json:"IpBlockRecordId,omitempty" xml:"IpBlockRecordId,omitempty"`
}

func (s CreateIpBlockRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIpBlockRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpBlockRecordResponseBody) SetRequestId(v string) *CreateIpBlockRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIpBlockRecordResponseBody) SetIpBlockRecordId(v string) *CreateIpBlockRecordResponseBody {
	s.IpBlockRecordId = &v
	return s
}

type CreateIpBlockRecordResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateIpBlockRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIpBlockRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIpBlockRecordResponse) GoString() string {
	return s.String()
}

func (s *CreateIpBlockRecordResponse) SetHeaders(v map[string]*string) *CreateIpBlockRecordResponse {
	s.Headers = v
	return s
}

func (s *CreateIpBlockRecordResponse) SetBody(v *CreateIpBlockRecordResponseBody) *CreateIpBlockRecordResponse {
	s.Body = v
	return s
}

type UpdateIpBlockRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 用于保证请求的幂等性。由客户端生成该参数值，要保证在不同请求间唯一。只支持 ASCII 字符，且不能超过 64 个字符
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 资源一级ID
	IpBlockId *string `json:"IpBlockId,omitempty" xml:"IpBlockId,omitempty"`
	// 业务类型UID
	BusinessTypeId *string `json:"BusinessTypeId,omitempty" xml:"BusinessTypeId,omitempty"`
	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// 园区名称
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	// 备份设备名称
	BackupDeviceName *string `json:"BackupDeviceName,omitempty" xml:"BackupDeviceName,omitempty"`
	// 公网地址类型 INC GUEST VIP
	NetBusiness *string `json:"NetBusiness,omitempty" xml:"NetBusiness,omitempty"`
	// IP归属
	Ownership *string `json:"Ownership,omitempty" xml:"Ownership,omitempty"`
	// IP用途
	Application *string `json:"Application,omitempty" xml:"Application,omitempty"`
	// 备注
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 更新类型 update 更新 split 拆分
	UpdateType *string `json:"UpdateType,omitempty" xml:"UpdateType,omitempty"`
	// 状态： using available lock
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateIpBlockRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpBlockRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpBlockRequest) SetInstanceId(v string) *UpdateIpBlockRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateIpBlockRequest) SetClientToken(v string) *UpdateIpBlockRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIpBlockRequest) SetIpBlockId(v string) *UpdateIpBlockRequest {
	s.IpBlockId = &v
	return s
}

func (s *UpdateIpBlockRequest) SetBusinessTypeId(v string) *UpdateIpBlockRequest {
	s.BusinessTypeId = &v
	return s
}

func (s *UpdateIpBlockRequest) SetDeviceName(v string) *UpdateIpBlockRequest {
	s.DeviceName = &v
	return s
}

func (s *UpdateIpBlockRequest) SetZoneName(v string) *UpdateIpBlockRequest {
	s.ZoneName = &v
	return s
}

func (s *UpdateIpBlockRequest) SetBackupDeviceName(v string) *UpdateIpBlockRequest {
	s.BackupDeviceName = &v
	return s
}

func (s *UpdateIpBlockRequest) SetNetBusiness(v string) *UpdateIpBlockRequest {
	s.NetBusiness = &v
	return s
}

func (s *UpdateIpBlockRequest) SetOwnership(v string) *UpdateIpBlockRequest {
	s.Ownership = &v
	return s
}

func (s *UpdateIpBlockRequest) SetApplication(v string) *UpdateIpBlockRequest {
	s.Application = &v
	return s
}

func (s *UpdateIpBlockRequest) SetDescription(v string) *UpdateIpBlockRequest {
	s.Description = &v
	return s
}

func (s *UpdateIpBlockRequest) SetUpdateType(v string) *UpdateIpBlockRequest {
	s.UpdateType = &v
	return s
}

func (s *UpdateIpBlockRequest) SetStatus(v string) *UpdateIpBlockRequest {
	s.Status = &v
	return s
}

type UpdateIpBlockResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpBlockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpBlockResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpBlockResponseBody) SetRequestId(v string) *UpdateIpBlockResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIpBlockResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateIpBlockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateIpBlockResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpBlockResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpBlockResponse) SetHeaders(v map[string]*string) *UpdateIpBlockResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpBlockResponse) SetBody(v *UpdateIpBlockResponseBody) *UpdateIpBlockResponse {
	s.Body = v
	return s
}

type CreateDevicePropertyRequest struct {
	// 设备形态ID
	DeviceFormId *string `json:"DeviceFormId,omitempty" xml:"DeviceFormId,omitempty"`
	// 属性名称
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	// 属性主键
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// 属性格式
	PropertyFormat *string `json:"PropertyFormat,omitempty" xml:"PropertyFormat,omitempty"`
	// 属性内容
	PropertyContent *string `json:"PropertyContent,omitempty" xml:"PropertyContent,omitempty"`
	// 幂等校验 token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateDevicePropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDevicePropertyRequest) GoString() string {
	return s.String()
}

func (s *CreateDevicePropertyRequest) SetDeviceFormId(v string) *CreateDevicePropertyRequest {
	s.DeviceFormId = &v
	return s
}

func (s *CreateDevicePropertyRequest) SetPropertyName(v string) *CreateDevicePropertyRequest {
	s.PropertyName = &v
	return s
}

func (s *CreateDevicePropertyRequest) SetPropertyKey(v string) *CreateDevicePropertyRequest {
	s.PropertyKey = &v
	return s
}

func (s *CreateDevicePropertyRequest) SetPropertyFormat(v string) *CreateDevicePropertyRequest {
	s.PropertyFormat = &v
	return s
}

func (s *CreateDevicePropertyRequest) SetPropertyContent(v string) *CreateDevicePropertyRequest {
	s.PropertyContent = &v
	return s
}

func (s *CreateDevicePropertyRequest) SetClientToken(v string) *CreateDevicePropertyRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDevicePropertyRequest) SetInstanceId(v string) *CreateDevicePropertyRequest {
	s.InstanceId = &v
	return s
}

type CreateDevicePropertyResponseBody struct {
	// 资源实例ID，如ECS实例的创建接口CreateInstance应返回InstanceId。
	DevicePropertyId *string `json:"DevicePropertyId,omitempty" xml:"DevicePropertyId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDevicePropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDevicePropertyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDevicePropertyResponseBody) SetDevicePropertyId(v string) *CreateDevicePropertyResponseBody {
	s.DevicePropertyId = &v
	return s
}

func (s *CreateDevicePropertyResponseBody) SetRequestId(v string) *CreateDevicePropertyResponseBody {
	s.RequestId = &v
	return s
}

type CreateDevicePropertyResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDevicePropertyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDevicePropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDevicePropertyResponse) GoString() string {
	return s.String()
}

func (s *CreateDevicePropertyResponse) SetHeaders(v map[string]*string) *CreateDevicePropertyResponse {
	s.Headers = v
	return s
}

func (s *CreateDevicePropertyResponse) SetBody(v *CreateDevicePropertyResponseBody) *CreateDevicePropertyResponse {
	s.Body = v
	return s
}

type UpdateIpBlockRecordRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 用于保证请求的幂等性。由客户端生成该参数值，要保证在不同请求间唯一。只支持 ASCII 字符，且不能超过 64 个字符
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 资源一级ID
	IpBlockRecordId *string `json:"IpBlockRecordId,omitempty" xml:"IpBlockRecordId,omitempty"`
	// 工单状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateIpBlockRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpBlockRecordRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpBlockRecordRequest) SetInstanceId(v string) *UpdateIpBlockRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateIpBlockRecordRequest) SetClientToken(v string) *UpdateIpBlockRecordRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIpBlockRecordRequest) SetIpBlockRecordId(v string) *UpdateIpBlockRecordRequest {
	s.IpBlockRecordId = &v
	return s
}

func (s *UpdateIpBlockRecordRequest) SetStatus(v string) *UpdateIpBlockRecordRequest {
	s.Status = &v
	return s
}

type UpdateIpBlockRecordResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpBlockRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpBlockRecordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpBlockRecordResponseBody) SetRequestId(v string) *UpdateIpBlockRecordResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIpBlockRecordResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateIpBlockRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateIpBlockRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpBlockRecordResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpBlockRecordResponse) SetHeaders(v map[string]*string) *UpdateIpBlockRecordResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpBlockRecordResponse) SetBody(v *UpdateIpBlockRecordResponseBody) *UpdateIpBlockRecordResponse {
	s.Body = v
	return s
}

type RetryTasksRequest struct {
	// 重执行任务的数组
	RetryTasks []*RetryTasksRequestRetryTasks `json:"RetryTasks,omitempty" xml:"RetryTasks,omitempty" type:"Repeated"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RetryTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s RetryTasksRequest) GoString() string {
	return s.String()
}

func (s *RetryTasksRequest) SetRetryTasks(v []*RetryTasksRequestRetryTasks) *RetryTasksRequest {
	s.RetryTasks = v
	return s
}

func (s *RetryTasksRequest) SetInstanceId(v string) *RetryTasksRequest {
	s.InstanceId = &v
	return s
}

type RetryTasksRequestRetryTasks struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 脚本ID
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s RetryTasksRequestRetryTasks) String() string {
	return tea.Prettify(s)
}

func (s RetryTasksRequestRetryTasks) GoString() string {
	return s.String()
}

func (s *RetryTasksRequestRetryTasks) SetDeviceId(v string) *RetryTasksRequestRetryTasks {
	s.DeviceId = &v
	return s
}

func (s *RetryTasksRequestRetryTasks) SetScriptId(v string) *RetryTasksRequestRetryTasks {
	s.ScriptId = &v
	return s
}

type RetryTasksShrinkRequest struct {
	// 重执行任务的数组
	RetryTasksShrink *string `json:"RetryTasks,omitempty" xml:"RetryTasks,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RetryTasksShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RetryTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *RetryTasksShrinkRequest) SetRetryTasksShrink(v string) *RetryTasksShrinkRequest {
	s.RetryTasksShrink = &v
	return s
}

func (s *RetryTasksShrinkRequest) SetInstanceId(v string) *RetryTasksShrinkRequest {
	s.InstanceId = &v
	return s
}

type RetryTasksResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetryTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetryTasksResponseBody) GoString() string {
	return s.String()
}

func (s *RetryTasksResponseBody) SetRequestId(v string) *RetryTasksResponseBody {
	s.RequestId = &v
	return s
}

type RetryTasksResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RetryTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RetryTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s RetryTasksResponse) GoString() string {
	return s.String()
}

func (s *RetryTasksResponse) SetHeaders(v map[string]*string) *RetryTasksResponse {
	s.Headers = v
	return s
}

func (s *RetryTasksResponse) SetBody(v *RetryTasksResponseBody) *RetryTasksResponse {
	s.Body = v
	return s
}

type GetPhysicalSpaceTopoRequest struct {
	// 物理空间id
	PhysicalSpaceId *string `json:"PhysicalSpaceId,omitempty" xml:"PhysicalSpaceId,omitempty"`
	// 拓扑类型
	TopoType *string `json:"TopoType,omitempty" xml:"TopoType,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetPhysicalSpaceTopoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalSpaceTopoRequest) GoString() string {
	return s.String()
}

func (s *GetPhysicalSpaceTopoRequest) SetPhysicalSpaceId(v string) *GetPhysicalSpaceTopoRequest {
	s.PhysicalSpaceId = &v
	return s
}

func (s *GetPhysicalSpaceTopoRequest) SetTopoType(v string) *GetPhysicalSpaceTopoRequest {
	s.TopoType = &v
	return s
}

func (s *GetPhysicalSpaceTopoRequest) SetInstanceId(v string) *GetPhysicalSpaceTopoRequest {
	s.InstanceId = &v
	return s
}

type GetPhysicalSpaceTopoResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 拓扑数据
	TopoData *GetPhysicalSpaceTopoResponseBodyTopoData `json:"TopoData,omitempty" xml:"TopoData,omitempty" type:"Struct"`
}

func (s GetPhysicalSpaceTopoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalSpaceTopoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhysicalSpaceTopoResponseBody) SetRequestId(v string) *GetPhysicalSpaceTopoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhysicalSpaceTopoResponseBody) SetTopoData(v *GetPhysicalSpaceTopoResponseBodyTopoData) *GetPhysicalSpaceTopoResponseBody {
	s.TopoData = v
	return s
}

type GetPhysicalSpaceTopoResponseBodyTopoData struct {
	// 更新时间
	UpdateTime *string                                            `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Links      []*GetPhysicalSpaceTopoResponseBodyTopoDataLinks   `json:"Links,omitempty" xml:"Links,omitempty" type:"Repeated"`
	Devices    []*GetPhysicalSpaceTopoResponseBodyTopoDataDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
}

func (s GetPhysicalSpaceTopoResponseBodyTopoData) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalSpaceTopoResponseBodyTopoData) GoString() string {
	return s.String()
}

func (s *GetPhysicalSpaceTopoResponseBodyTopoData) SetUpdateTime(v string) *GetPhysicalSpaceTopoResponseBodyTopoData {
	s.UpdateTime = &v
	return s
}

func (s *GetPhysicalSpaceTopoResponseBodyTopoData) SetLinks(v []*GetPhysicalSpaceTopoResponseBodyTopoDataLinks) *GetPhysicalSpaceTopoResponseBodyTopoData {
	s.Links = v
	return s
}

func (s *GetPhysicalSpaceTopoResponseBodyTopoData) SetDevices(v []*GetPhysicalSpaceTopoResponseBodyTopoDataDevices) *GetPhysicalSpaceTopoResponseBodyTopoData {
	s.Devices = v
	return s
}

type GetPhysicalSpaceTopoResponseBodyTopoDataLinks struct {
	// 源设备id
	SourceDeviceId *string `json:"SourceDeviceId,omitempty" xml:"SourceDeviceId,omitempty"`
	// 源设备端口
	SourcePort *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// 目标设备id
	TargetDeviceId *string `json:"TargetDeviceId,omitempty" xml:"TargetDeviceId,omitempty"`
	// 目标设备端口
	TargetPort *string `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	// 源设备名
	SourceDeviceName *string `json:"SourceDeviceName,omitempty" xml:"SourceDeviceName,omitempty"`
	// 目标设备名
	TargetDeviceName *string `json:"TargetDeviceName,omitempty" xml:"TargetDeviceName,omitempty"`
}

func (s GetPhysicalSpaceTopoResponseBodyTopoDataLinks) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalSpaceTopoResponseBodyTopoDataLinks) GoString() string {
	return s.String()
}

func (s *GetPhysicalSpaceTopoResponseBodyTopoDataLinks) SetSourceDeviceId(v string) *GetPhysicalSpaceTopoResponseBodyTopoDataLinks {
	s.SourceDeviceId = &v
	return s
}

func (s *GetPhysicalSpaceTopoResponseBodyTopoDataLinks) SetSourcePort(v string) *GetPhysicalSpaceTopoResponseBodyTopoDataLinks {
	s.SourcePort = &v
	return s
}

func (s *GetPhysicalSpaceTopoResponseBodyTopoDataLinks) SetTargetDeviceId(v string) *GetPhysicalSpaceTopoResponseBodyTopoDataLinks {
	s.TargetDeviceId = &v
	return s
}

func (s *GetPhysicalSpaceTopoResponseBodyTopoDataLinks) SetTargetPort(v string) *GetPhysicalSpaceTopoResponseBodyTopoDataLinks {
	s.TargetPort = &v
	return s
}

func (s *GetPhysicalSpaceTopoResponseBodyTopoDataLinks) SetSourceDeviceName(v string) *GetPhysicalSpaceTopoResponseBodyTopoDataLinks {
	s.SourceDeviceName = &v
	return s
}

func (s *GetPhysicalSpaceTopoResponseBodyTopoDataLinks) SetTargetDeviceName(v string) *GetPhysicalSpaceTopoResponseBodyTopoDataLinks {
	s.TargetDeviceName = &v
	return s
}

type GetPhysicalSpaceTopoResponseBodyTopoDataDevices struct {
	// 设备id
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 设备角色
	DeviceRole *string `json:"DeviceRole,omitempty" xml:"DeviceRole,omitempty"`
	// 设备主机名
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// 设备ip
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s GetPhysicalSpaceTopoResponseBodyTopoDataDevices) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalSpaceTopoResponseBodyTopoDataDevices) GoString() string {
	return s.String()
}

func (s *GetPhysicalSpaceTopoResponseBodyTopoDataDevices) SetDeviceId(v string) *GetPhysicalSpaceTopoResponseBodyTopoDataDevices {
	s.DeviceId = &v
	return s
}

func (s *GetPhysicalSpaceTopoResponseBodyTopoDataDevices) SetDeviceRole(v string) *GetPhysicalSpaceTopoResponseBodyTopoDataDevices {
	s.DeviceRole = &v
	return s
}

func (s *GetPhysicalSpaceTopoResponseBodyTopoDataDevices) SetHostName(v string) *GetPhysicalSpaceTopoResponseBodyTopoDataDevices {
	s.HostName = &v
	return s
}

func (s *GetPhysicalSpaceTopoResponseBodyTopoDataDevices) SetIp(v string) *GetPhysicalSpaceTopoResponseBodyTopoDataDevices {
	s.Ip = &v
	return s
}

type GetPhysicalSpaceTopoResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPhysicalSpaceTopoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPhysicalSpaceTopoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalSpaceTopoResponse) GoString() string {
	return s.String()
}

func (s *GetPhysicalSpaceTopoResponse) SetHeaders(v map[string]*string) *GetPhysicalSpaceTopoResponse {
	s.Headers = v
	return s
}

func (s *GetPhysicalSpaceTopoResponse) SetBody(v *GetPhysicalSpaceTopoResponseBody) *GetPhysicalSpaceTopoResponse {
	s.Body = v
	return s
}

type GetOssPolicyRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetOssPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOssPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetOssPolicyRequest) SetInstanceId(v string) *GetOssPolicyRequest {
	s.InstanceId = &v
	return s
}

type GetOssPolicyResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 资源上传规则
	OssPolicy *GetOssPolicyResponseBodyOssPolicy `json:"OssPolicy,omitempty" xml:"OssPolicy,omitempty" type:"Struct"`
}

func (s GetOssPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOssPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssPolicyResponseBody) SetRequestId(v string) *GetOssPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetOssPolicy(v *GetOssPolicyResponseBodyOssPolicy) *GetOssPolicyResponseBody {
	s.OssPolicy = v
	return s
}

type GetOssPolicyResponseBodyOssPolicy struct {
	// 通行id
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// 通行规则
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// 签名
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// 目录
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// 主机名
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
}

func (s GetOssPolicyResponseBodyOssPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetOssPolicyResponseBodyOssPolicy) GoString() string {
	return s.String()
}

func (s *GetOssPolicyResponseBodyOssPolicy) SetAccessId(v string) *GetOssPolicyResponseBodyOssPolicy {
	s.AccessId = &v
	return s
}

func (s *GetOssPolicyResponseBodyOssPolicy) SetPolicy(v string) *GetOssPolicyResponseBodyOssPolicy {
	s.Policy = &v
	return s
}

func (s *GetOssPolicyResponseBodyOssPolicy) SetSignature(v string) *GetOssPolicyResponseBodyOssPolicy {
	s.Signature = &v
	return s
}

func (s *GetOssPolicyResponseBodyOssPolicy) SetDirectory(v string) *GetOssPolicyResponseBodyOssPolicy {
	s.Directory = &v
	return s
}

func (s *GetOssPolicyResponseBodyOssPolicy) SetHost(v string) *GetOssPolicyResponseBodyOssPolicy {
	s.Host = &v
	return s
}

func (s *GetOssPolicyResponseBodyOssPolicy) SetExpireTime(v string) *GetOssPolicyResponseBodyOssPolicy {
	s.ExpireTime = &v
	return s
}

type GetOssPolicyResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOssPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOssPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOssPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetOssPolicyResponse) SetHeaders(v map[string]*string) *GetOssPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetOssPolicyResponse) SetBody(v *GetOssPolicyResponseBody) *GetOssPolicyResponse {
	s.Body = v
	return s
}

type DeleteDeviceFormRequest struct {
	// 实例 ID。
	DeviceFormId *string `json:"DeviceFormId,omitempty" xml:"DeviceFormId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteDeviceFormRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceFormRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeviceFormRequest) SetDeviceFormId(v string) *DeleteDeviceFormRequest {
	s.DeviceFormId = &v
	return s
}

func (s *DeleteDeviceFormRequest) SetInstanceId(v string) *DeleteDeviceFormRequest {
	s.InstanceId = &v
	return s
}

type DeleteDeviceFormResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDeviceFormResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceFormResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeviceFormResponseBody) SetRequestId(v string) *DeleteDeviceFormResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDeviceFormResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDeviceFormResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDeviceFormResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceFormResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeviceFormResponse) SetHeaders(v map[string]*string) *DeleteDeviceFormResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeviceFormResponse) SetBody(v *DeleteDeviceFormResponseBody) *DeleteDeviceFormResponse {
	s.Body = v
	return s
}

type UpdateDeviceFormRequest struct {
	// 设备形态ID
	DeviceFormId *string `json:"DeviceFormId,omitempty" xml:"DeviceFormId,omitempty"`
	// 是否支持配置生成
	ConfigCompare *bool `json:"ConfigCompare,omitempty" xml:"ConfigCompare,omitempty"`
	// 是否需要账号配置
	AccountConfig *bool `json:"AccountConfig,omitempty" xml:"AccountConfig,omitempty"`
	// 设备形态属性列表
	AttributeList []*UpdateDeviceFormRequestAttributeList `json:"AttributeList,omitempty" xml:"AttributeList,omitempty" type:"Repeated"`
	// 是否展示设备详情
	DetailDisplay *bool `json:"DetailDisplay,omitempty" xml:"DetailDisplay,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateDeviceFormRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceFormRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeviceFormRequest) SetDeviceFormId(v string) *UpdateDeviceFormRequest {
	s.DeviceFormId = &v
	return s
}

func (s *UpdateDeviceFormRequest) SetConfigCompare(v bool) *UpdateDeviceFormRequest {
	s.ConfigCompare = &v
	return s
}

func (s *UpdateDeviceFormRequest) SetAccountConfig(v bool) *UpdateDeviceFormRequest {
	s.AccountConfig = &v
	return s
}

func (s *UpdateDeviceFormRequest) SetAttributeList(v []*UpdateDeviceFormRequestAttributeList) *UpdateDeviceFormRequest {
	s.AttributeList = v
	return s
}

func (s *UpdateDeviceFormRequest) SetDetailDisplay(v bool) *UpdateDeviceFormRequest {
	s.DetailDisplay = &v
	return s
}

func (s *UpdateDeviceFormRequest) SetInstanceId(v string) *UpdateDeviceFormRequest {
	s.InstanceId = &v
	return s
}

type UpdateDeviceFormRequestAttributeList struct {
	// 设备形态属性主键
	AttributeKey *string `json:"AttributeKey,omitempty" xml:"AttributeKey,omitempty"`
	// 设备形态属性名称
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	// 设备形态属性是否必填
	AttributeRequirement *bool `json:"AttributeRequirement,omitempty" xml:"AttributeRequirement,omitempty"`
	// 设备形态属性是否唯一
	AttributeUniqueness *bool `json:"AttributeUniqueness,omitempty" xml:"AttributeUniqueness,omitempty"`
	// 设备形态属性值格式
	AttributeFormat *string `json:"AttributeFormat,omitempty" xml:"AttributeFormat,omitempty"`
	// 设备形态属性值类型
	AttributeType *string `json:"AttributeType,omitempty" xml:"AttributeType,omitempty"`
	// 设备形态属性关联对象
	AttributeReference *string `json:"AttributeReference,omitempty" xml:"AttributeReference,omitempty"`
	// 设备形态属性是否表格可见
	AttributeTableDisplay *bool `json:"AttributeTableDisplay,omitempty" xml:"AttributeTableDisplay,omitempty"`
	// 前端查询控件占位符
	AttributePlaceholder *string `json:"AttributePlaceholder,omitempty" xml:"AttributePlaceholder,omitempty"`
	// 前端展示搜索控件
	AttributeQuery *bool `json:"AttributeQuery,omitempty" xml:"AttributeQuery,omitempty"`
	// 查询支持模糊搜索
	AttributeFuzzyQuery *bool `json:"AttributeFuzzyQuery,omitempty" xml:"AttributeFuzzyQuery,omitempty"`
}

func (s UpdateDeviceFormRequestAttributeList) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceFormRequestAttributeList) GoString() string {
	return s.String()
}

func (s *UpdateDeviceFormRequestAttributeList) SetAttributeKey(v string) *UpdateDeviceFormRequestAttributeList {
	s.AttributeKey = &v
	return s
}

func (s *UpdateDeviceFormRequestAttributeList) SetAttributeName(v string) *UpdateDeviceFormRequestAttributeList {
	s.AttributeName = &v
	return s
}

func (s *UpdateDeviceFormRequestAttributeList) SetAttributeRequirement(v bool) *UpdateDeviceFormRequestAttributeList {
	s.AttributeRequirement = &v
	return s
}

func (s *UpdateDeviceFormRequestAttributeList) SetAttributeUniqueness(v bool) *UpdateDeviceFormRequestAttributeList {
	s.AttributeUniqueness = &v
	return s
}

func (s *UpdateDeviceFormRequestAttributeList) SetAttributeFormat(v string) *UpdateDeviceFormRequestAttributeList {
	s.AttributeFormat = &v
	return s
}

func (s *UpdateDeviceFormRequestAttributeList) SetAttributeType(v string) *UpdateDeviceFormRequestAttributeList {
	s.AttributeType = &v
	return s
}

func (s *UpdateDeviceFormRequestAttributeList) SetAttributeReference(v string) *UpdateDeviceFormRequestAttributeList {
	s.AttributeReference = &v
	return s
}

func (s *UpdateDeviceFormRequestAttributeList) SetAttributeTableDisplay(v bool) *UpdateDeviceFormRequestAttributeList {
	s.AttributeTableDisplay = &v
	return s
}

func (s *UpdateDeviceFormRequestAttributeList) SetAttributePlaceholder(v string) *UpdateDeviceFormRequestAttributeList {
	s.AttributePlaceholder = &v
	return s
}

func (s *UpdateDeviceFormRequestAttributeList) SetAttributeQuery(v bool) *UpdateDeviceFormRequestAttributeList {
	s.AttributeQuery = &v
	return s
}

func (s *UpdateDeviceFormRequestAttributeList) SetAttributeFuzzyQuery(v bool) *UpdateDeviceFormRequestAttributeList {
	s.AttributeFuzzyQuery = &v
	return s
}

type UpdateDeviceFormShrinkRequest struct {
	// 设备形态ID
	DeviceFormId *string `json:"DeviceFormId,omitempty" xml:"DeviceFormId,omitempty"`
	// 是否支持配置生成
	ConfigCompare *bool `json:"ConfigCompare,omitempty" xml:"ConfigCompare,omitempty"`
	// 是否需要账号配置
	AccountConfig *bool `json:"AccountConfig,omitempty" xml:"AccountConfig,omitempty"`
	// 设备形态属性列表
	AttributeListShrink *string `json:"AttributeList,omitempty" xml:"AttributeList,omitempty"`
	// 是否展示设备详情
	DetailDisplay *bool `json:"DetailDisplay,omitempty" xml:"DetailDisplay,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateDeviceFormShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceFormShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeviceFormShrinkRequest) SetDeviceFormId(v string) *UpdateDeviceFormShrinkRequest {
	s.DeviceFormId = &v
	return s
}

func (s *UpdateDeviceFormShrinkRequest) SetConfigCompare(v bool) *UpdateDeviceFormShrinkRequest {
	s.ConfigCompare = &v
	return s
}

func (s *UpdateDeviceFormShrinkRequest) SetAccountConfig(v bool) *UpdateDeviceFormShrinkRequest {
	s.AccountConfig = &v
	return s
}

func (s *UpdateDeviceFormShrinkRequest) SetAttributeListShrink(v string) *UpdateDeviceFormShrinkRequest {
	s.AttributeListShrink = &v
	return s
}

func (s *UpdateDeviceFormShrinkRequest) SetDetailDisplay(v bool) *UpdateDeviceFormShrinkRequest {
	s.DetailDisplay = &v
	return s
}

func (s *UpdateDeviceFormShrinkRequest) SetInstanceId(v string) *UpdateDeviceFormShrinkRequest {
	s.InstanceId = &v
	return s
}

type UpdateDeviceFormResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDeviceFormResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceFormResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeviceFormResponseBody) SetRequestId(v string) *UpdateDeviceFormResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDeviceFormResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDeviceFormResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDeviceFormResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceFormResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeviceFormResponse) SetHeaders(v map[string]*string) *UpdateDeviceFormResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeviceFormResponse) SetBody(v *UpdateDeviceFormResponseBody) *UpdateDeviceFormResponse {
	s.Body = v
	return s
}

type UpdateDeviceRequest struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 物理空间
	PhysicalSpaceId *string `json:"PhysicalSpaceId,omitempty" xml:"PhysicalSpaceId,omitempty"`
	// 主机名
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// 设备IP
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// 设备SN
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// 设备MAC地址
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// 设备厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 设备型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 设备状态
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// 设备安全域
	SecurityDomain *string `json:"SecurityDomain,omitempty" xml:"SecurityDomain,omitempty"`
	// 登录类型
	LoginType *string `json:"LoginType,omitempty" xml:"LoginType,omitempty"`
	// 登录账号
	LoginUsername *string `json:"LoginUsername,omitempty" xml:"LoginUsername,omitempty"`
	// 登录密码
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// enable密码
	EnablePassword *string `json:"EnablePassword,omitempty" xml:"EnablePassword,omitempty"`
	// SNMP 版本号
	SnmpAccountVersion *string `json:"SnmpAccountVersion,omitempty" xml:"SnmpAccountVersion,omitempty"`
	// SNMP Community
	SnmpCommunity *string `json:"SnmpCommunity,omitempty" xml:"SnmpCommunity,omitempty"`
	// SNMP 账号类型
	SnmpAccountType *string `json:"SnmpAccountType,omitempty" xml:"SnmpAccountType,omitempty"`
	// SNMP 安全级别
	SnmpSecurityLevel *string `json:"SnmpSecurityLevel,omitempty" xml:"SnmpSecurityLevel,omitempty"`
	// SNMP 用户名
	SnmpUsername *string `json:"SnmpUsername,omitempty" xml:"SnmpUsername,omitempty"`
	// SNMP Auth PassPhrase
	SnmpAuthPassphrase *string `json:"SnmpAuthPassphrase,omitempty" xml:"SnmpAuthPassphrase,omitempty"`
	// Auth Protocol
	SnmpAuthProtocol *string `json:"SnmpAuthProtocol,omitempty" xml:"SnmpAuthProtocol,omitempty"`
	// Privacy Passphase
	SnmpPrivacyPassphase *string `json:"SnmpPrivacyPassphase,omitempty" xml:"SnmpPrivacyPassphase,omitempty"`
	// Privacy Protocol
	SnmpPrivacyProtocol *string `json:"SnmpPrivacyProtocol,omitempty" xml:"SnmpPrivacyProtocol,omitempty"`
	// 设备额外属性
	ExtAttributes *string `json:"ExtAttributes,omitempty" xml:"ExtAttributes,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeviceRequest) SetDeviceId(v string) *UpdateDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *UpdateDeviceRequest) SetPhysicalSpaceId(v string) *UpdateDeviceRequest {
	s.PhysicalSpaceId = &v
	return s
}

func (s *UpdateDeviceRequest) SetHostName(v string) *UpdateDeviceRequest {
	s.HostName = &v
	return s
}

func (s *UpdateDeviceRequest) SetIp(v string) *UpdateDeviceRequest {
	s.Ip = &v
	return s
}

func (s *UpdateDeviceRequest) SetSn(v string) *UpdateDeviceRequest {
	s.Sn = &v
	return s
}

func (s *UpdateDeviceRequest) SetMac(v string) *UpdateDeviceRequest {
	s.Mac = &v
	return s
}

func (s *UpdateDeviceRequest) SetVendor(v string) *UpdateDeviceRequest {
	s.Vendor = &v
	return s
}

func (s *UpdateDeviceRequest) SetModel(v string) *UpdateDeviceRequest {
	s.Model = &v
	return s
}

func (s *UpdateDeviceRequest) SetServiceStatus(v string) *UpdateDeviceRequest {
	s.ServiceStatus = &v
	return s
}

func (s *UpdateDeviceRequest) SetSecurityDomain(v string) *UpdateDeviceRequest {
	s.SecurityDomain = &v
	return s
}

func (s *UpdateDeviceRequest) SetLoginType(v string) *UpdateDeviceRequest {
	s.LoginType = &v
	return s
}

func (s *UpdateDeviceRequest) SetLoginUsername(v string) *UpdateDeviceRequest {
	s.LoginUsername = &v
	return s
}

func (s *UpdateDeviceRequest) SetLoginPassword(v string) *UpdateDeviceRequest {
	s.LoginPassword = &v
	return s
}

func (s *UpdateDeviceRequest) SetEnablePassword(v string) *UpdateDeviceRequest {
	s.EnablePassword = &v
	return s
}

func (s *UpdateDeviceRequest) SetSnmpAccountVersion(v string) *UpdateDeviceRequest {
	s.SnmpAccountVersion = &v
	return s
}

func (s *UpdateDeviceRequest) SetSnmpCommunity(v string) *UpdateDeviceRequest {
	s.SnmpCommunity = &v
	return s
}

func (s *UpdateDeviceRequest) SetSnmpAccountType(v string) *UpdateDeviceRequest {
	s.SnmpAccountType = &v
	return s
}

func (s *UpdateDeviceRequest) SetSnmpSecurityLevel(v string) *UpdateDeviceRequest {
	s.SnmpSecurityLevel = &v
	return s
}

func (s *UpdateDeviceRequest) SetSnmpUsername(v string) *UpdateDeviceRequest {
	s.SnmpUsername = &v
	return s
}

func (s *UpdateDeviceRequest) SetSnmpAuthPassphrase(v string) *UpdateDeviceRequest {
	s.SnmpAuthPassphrase = &v
	return s
}

func (s *UpdateDeviceRequest) SetSnmpAuthProtocol(v string) *UpdateDeviceRequest {
	s.SnmpAuthProtocol = &v
	return s
}

func (s *UpdateDeviceRequest) SetSnmpPrivacyPassphase(v string) *UpdateDeviceRequest {
	s.SnmpPrivacyPassphase = &v
	return s
}

func (s *UpdateDeviceRequest) SetSnmpPrivacyProtocol(v string) *UpdateDeviceRequest {
	s.SnmpPrivacyProtocol = &v
	return s
}

func (s *UpdateDeviceRequest) SetExtAttributes(v string) *UpdateDeviceRequest {
	s.ExtAttributes = &v
	return s
}

func (s *UpdateDeviceRequest) SetInstanceId(v string) *UpdateDeviceRequest {
	s.InstanceId = &v
	return s
}

type UpdateDeviceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeviceResponseBody) SetRequestId(v string) *UpdateDeviceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDeviceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeviceResponse) SetHeaders(v map[string]*string) *UpdateDeviceResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeviceResponse) SetBody(v *UpdateDeviceResponseBody) *UpdateDeviceResponse {
	s.Body = v
	return s
}

type GetDeviceRequest struct {
	// 实例 ID。
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceRequest) SetDeviceId(v string) *GetDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceRequest) SetInstanceId(v string) *GetDeviceRequest {
	s.InstanceId = &v
	return s
}

type GetDeviceResponseBody struct {
	// 设备详情
	Device *GetDeviceResponseBodyDevice `json:"Device,omitempty" xml:"Device,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceResponseBody) SetDevice(v *GetDeviceResponseBodyDevice) *GetDeviceResponseBody {
	s.Device = v
	return s
}

func (s *GetDeviceResponseBody) SetRequestId(v string) *GetDeviceResponseBody {
	s.RequestId = &v
	return s
}

type GetDeviceResponseBodyDevice struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 设备形态ID
	DeviceFormId *string `json:"DeviceFormId,omitempty" xml:"DeviceFormId,omitempty"`
	// 设备形态名称
	DeviceFormName *string `json:"DeviceFormName,omitempty" xml:"DeviceFormName,omitempty"`
	// 物理空间ID
	PhysicalSpaceId *string `json:"PhysicalSpaceId,omitempty" xml:"PhysicalSpaceId,omitempty"`
	// 物理空间名称
	PhysicalSpaceName *string `json:"PhysicalSpaceName,omitempty" xml:"PhysicalSpaceName,omitempty"`
	// 主机名
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// 设备IP
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// 设备SN
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// 设备MAC地址
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// 设备厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 设备型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 设备安全域
	SecurityDomain *string `json:"SecurityDomain,omitempty" xml:"SecurityDomain,omitempty"`
	// 设备状态
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// 登录类型
	LoginType *string `json:"LoginType,omitempty" xml:"LoginType,omitempty"`
	// 登录账号
	LoginUsername *string `json:"LoginUsername,omitempty" xml:"LoginUsername,omitempty"`
	// 登录密码
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// enable密码
	EnablePassword *string `json:"EnablePassword,omitempty" xml:"EnablePassword,omitempty"`
	// SNMP版本号
	SnmpAccountVersion *string `json:"SnmpAccountVersion,omitempty" xml:"SnmpAccountVersion,omitempty"`
	// SNMP Community
	SnmpCommunity *string `json:"SnmpCommunity,omitempty" xml:"SnmpCommunity,omitempty"`
	// SNMP 账号类型
	SnmpAccountType *string `json:"SnmpAccountType,omitempty" xml:"SnmpAccountType,omitempty"`
	// SNMP 安全级别
	SnmpSecurityLevel *string `json:"SnmpSecurityLevel,omitempty" xml:"SnmpSecurityLevel,omitempty"`
	// SNMP 用户名
	SnmpUsername *string `json:"SnmpUsername,omitempty" xml:"SnmpUsername,omitempty"`
	// SNMP Auth PassPhrase
	SnmpAuthPassphrase *string `json:"SnmpAuthPassphrase,omitempty" xml:"SnmpAuthPassphrase,omitempty"`
	// SNMP Auth Protocol
	SnmpAuthProtocol *string `json:"SnmpAuthProtocol,omitempty" xml:"SnmpAuthProtocol,omitempty"`
	// SNMP Privacy Passphase
	SnmpPrivacyPassphase *string `json:"SnmpPrivacyPassphase,omitempty" xml:"SnmpPrivacyPassphase,omitempty"`
	// SNMP Privacy Protocol
	SnmpPrivacyProtocol *string `json:"SnmpPrivacyProtocol,omitempty" xml:"SnmpPrivacyProtocol,omitempty"`
	// 设备额外属性
	ExtAttributes *string `json:"ExtAttributes,omitempty" xml:"ExtAttributes,omitempty"`
}

func (s GetDeviceResponseBodyDevice) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceResponseBodyDevice) GoString() string {
	return s.String()
}

func (s *GetDeviceResponseBodyDevice) SetDeviceId(v string) *GetDeviceResponseBodyDevice {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetDeviceFormId(v string) *GetDeviceResponseBodyDevice {
	s.DeviceFormId = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetDeviceFormName(v string) *GetDeviceResponseBodyDevice {
	s.DeviceFormName = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetPhysicalSpaceId(v string) *GetDeviceResponseBodyDevice {
	s.PhysicalSpaceId = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetPhysicalSpaceName(v string) *GetDeviceResponseBodyDevice {
	s.PhysicalSpaceName = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetHostName(v string) *GetDeviceResponseBodyDevice {
	s.HostName = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetIp(v string) *GetDeviceResponseBodyDevice {
	s.Ip = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetSn(v string) *GetDeviceResponseBodyDevice {
	s.Sn = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetMac(v string) *GetDeviceResponseBodyDevice {
	s.Mac = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetVendor(v string) *GetDeviceResponseBodyDevice {
	s.Vendor = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetModel(v string) *GetDeviceResponseBodyDevice {
	s.Model = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetSecurityDomain(v string) *GetDeviceResponseBodyDevice {
	s.SecurityDomain = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetServiceStatus(v string) *GetDeviceResponseBodyDevice {
	s.ServiceStatus = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetLoginType(v string) *GetDeviceResponseBodyDevice {
	s.LoginType = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetLoginUsername(v string) *GetDeviceResponseBodyDevice {
	s.LoginUsername = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetLoginPassword(v string) *GetDeviceResponseBodyDevice {
	s.LoginPassword = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetEnablePassword(v string) *GetDeviceResponseBodyDevice {
	s.EnablePassword = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetSnmpAccountVersion(v string) *GetDeviceResponseBodyDevice {
	s.SnmpAccountVersion = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetSnmpCommunity(v string) *GetDeviceResponseBodyDevice {
	s.SnmpCommunity = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetSnmpAccountType(v string) *GetDeviceResponseBodyDevice {
	s.SnmpAccountType = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetSnmpSecurityLevel(v string) *GetDeviceResponseBodyDevice {
	s.SnmpSecurityLevel = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetSnmpUsername(v string) *GetDeviceResponseBodyDevice {
	s.SnmpUsername = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetSnmpAuthPassphrase(v string) *GetDeviceResponseBodyDevice {
	s.SnmpAuthPassphrase = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetSnmpAuthProtocol(v string) *GetDeviceResponseBodyDevice {
	s.SnmpAuthProtocol = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetSnmpPrivacyPassphase(v string) *GetDeviceResponseBodyDevice {
	s.SnmpPrivacyPassphase = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetSnmpPrivacyProtocol(v string) *GetDeviceResponseBodyDevice {
	s.SnmpPrivacyProtocol = &v
	return s
}

func (s *GetDeviceResponseBodyDevice) SetExtAttributes(v string) *GetDeviceResponseBodyDevice {
	s.ExtAttributes = &v
	return s
}

type GetDeviceResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceResponse) SetHeaders(v map[string]*string) *GetDeviceResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceResponse) SetBody(v *GetDeviceResponseBody) *GetDeviceResponse {
	s.Body = v
	return s
}

type UpdateSetupProjectRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源一级ID
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// 物理空间uId
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 架构id
	ArchitectureId *string `json:"ArchitectureId,omitempty" xml:"ArchitectureId,omitempty"`
	// 预计交付时间
	DeliveryTime *string `json:"DeliveryTime,omitempty" xml:"DeliveryTime,omitempty"`
	// 节点
	Nodes *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	// 套餐
	Packages []*UpdateSetupProjectRequestPackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Repeated"`
}

func (s UpdateSetupProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSetupProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateSetupProjectRequest) SetInstanceId(v string) *UpdateSetupProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateSetupProjectRequest) SetSetupProjectId(v string) *UpdateSetupProjectRequest {
	s.SetupProjectId = &v
	return s
}

func (s *UpdateSetupProjectRequest) SetSpaceId(v string) *UpdateSetupProjectRequest {
	s.SpaceId = &v
	return s
}

func (s *UpdateSetupProjectRequest) SetDescription(v string) *UpdateSetupProjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateSetupProjectRequest) SetArchitectureId(v string) *UpdateSetupProjectRequest {
	s.ArchitectureId = &v
	return s
}

func (s *UpdateSetupProjectRequest) SetDeliveryTime(v string) *UpdateSetupProjectRequest {
	s.DeliveryTime = &v
	return s
}

func (s *UpdateSetupProjectRequest) SetNodes(v string) *UpdateSetupProjectRequest {
	s.Nodes = &v
	return s
}

func (s *UpdateSetupProjectRequest) SetPackages(v []*UpdateSetupProjectRequestPackages) *UpdateSetupProjectRequest {
	s.Packages = v
	return s
}

type UpdateSetupProjectRequestPackages struct {
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 设备号
	DeviceNumber *int64 `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 型号
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s UpdateSetupProjectRequestPackages) String() string {
	return tea.Prettify(s)
}

func (s UpdateSetupProjectRequestPackages) GoString() string {
	return s.String()
}

func (s *UpdateSetupProjectRequestPackages) SetRole(v string) *UpdateSetupProjectRequestPackages {
	s.Role = &v
	return s
}

func (s *UpdateSetupProjectRequestPackages) SetDeviceNumber(v int64) *UpdateSetupProjectRequestPackages {
	s.DeviceNumber = &v
	return s
}

func (s *UpdateSetupProjectRequestPackages) SetVendor(v string) *UpdateSetupProjectRequestPackages {
	s.Vendor = &v
	return s
}

func (s *UpdateSetupProjectRequestPackages) SetModel(v string) *UpdateSetupProjectRequestPackages {
	s.Model = &v
	return s
}

type UpdateSetupProjectShrinkRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源一级ID
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// 物理空间uId
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 架构id
	ArchitectureId *string `json:"ArchitectureId,omitempty" xml:"ArchitectureId,omitempty"`
	// 预计交付时间
	DeliveryTime *string `json:"DeliveryTime,omitempty" xml:"DeliveryTime,omitempty"`
	// 节点
	Nodes *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	// 套餐
	PackagesShrink *string `json:"Packages,omitempty" xml:"Packages,omitempty"`
}

func (s UpdateSetupProjectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSetupProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSetupProjectShrinkRequest) SetInstanceId(v string) *UpdateSetupProjectShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateSetupProjectShrinkRequest) SetSetupProjectId(v string) *UpdateSetupProjectShrinkRequest {
	s.SetupProjectId = &v
	return s
}

func (s *UpdateSetupProjectShrinkRequest) SetSpaceId(v string) *UpdateSetupProjectShrinkRequest {
	s.SpaceId = &v
	return s
}

func (s *UpdateSetupProjectShrinkRequest) SetDescription(v string) *UpdateSetupProjectShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateSetupProjectShrinkRequest) SetArchitectureId(v string) *UpdateSetupProjectShrinkRequest {
	s.ArchitectureId = &v
	return s
}

func (s *UpdateSetupProjectShrinkRequest) SetDeliveryTime(v string) *UpdateSetupProjectShrinkRequest {
	s.DeliveryTime = &v
	return s
}

func (s *UpdateSetupProjectShrinkRequest) SetNodes(v string) *UpdateSetupProjectShrinkRequest {
	s.Nodes = &v
	return s
}

func (s *UpdateSetupProjectShrinkRequest) SetPackagesShrink(v string) *UpdateSetupProjectShrinkRequest {
	s.PackagesShrink = &v
	return s
}

type UpdateSetupProjectResponseBody struct {
	// 实例的名称
	SetupProjectName *string `json:"SetupProjectName,omitempty" xml:"SetupProjectName,omitempty"`
	// 资源实例ID，如ECS实例的创建接口CreateInstance应返回InstanceId。
	SetupProjectId *string `json:"SetupProjectId,omitempty" xml:"SetupProjectId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例的规格。
	SetupProjectSpecification *string `json:"SetupProjectSpecification,omitempty" xml:"SetupProjectSpecification,omitempty"`
}

func (s UpdateSetupProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSetupProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSetupProjectResponseBody) SetSetupProjectName(v string) *UpdateSetupProjectResponseBody {
	s.SetupProjectName = &v
	return s
}

func (s *UpdateSetupProjectResponseBody) SetSetupProjectId(v string) *UpdateSetupProjectResponseBody {
	s.SetupProjectId = &v
	return s
}

func (s *UpdateSetupProjectResponseBody) SetRequestId(v string) *UpdateSetupProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSetupProjectResponseBody) SetSetupProjectSpecification(v string) *UpdateSetupProjectResponseBody {
	s.SetupProjectSpecification = &v
	return s
}

type UpdateSetupProjectResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSetupProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSetupProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSetupProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateSetupProjectResponse) SetHeaders(v map[string]*string) *UpdateSetupProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateSetupProjectResponse) SetBody(v *UpdateSetupProjectResponseBody) *UpdateSetupProjectResponse {
	s.Body = v
	return s
}

type UpdateSpaceModelRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 物理空间类型
	SpaceType *string `json:"SpaceType,omitempty" xml:"SpaceType,omitempty"`
	// 资源uuid
	SpaceModelId *string `json:"SpaceModelId,omitempty" xml:"SpaceModelId,omitempty"`
	// 层级
	Sort []*UpdateSpaceModelRequestSort `json:"Sort,omitempty" xml:"Sort,omitempty" type:"Repeated"`
}

func (s UpdateSpaceModelRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpaceModelRequest) GoString() string {
	return s.String()
}

func (s *UpdateSpaceModelRequest) SetInstanceId(v string) *UpdateSpaceModelRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateSpaceModelRequest) SetSpaceType(v string) *UpdateSpaceModelRequest {
	s.SpaceType = &v
	return s
}

func (s *UpdateSpaceModelRequest) SetSpaceModelId(v string) *UpdateSpaceModelRequest {
	s.SpaceModelId = &v
	return s
}

func (s *UpdateSpaceModelRequest) SetSort(v []*UpdateSpaceModelRequestSort) *UpdateSpaceModelRequest {
	s.Sort = v
	return s
}

type UpdateSpaceModelRequestSort struct {
	// 层级名称
	LevelName *string `json:"LevelName,omitempty" xml:"LevelName,omitempty"`
	// 层级
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s UpdateSpaceModelRequestSort) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpaceModelRequestSort) GoString() string {
	return s.String()
}

func (s *UpdateSpaceModelRequestSort) SetLevelName(v string) *UpdateSpaceModelRequestSort {
	s.LevelName = &v
	return s
}

func (s *UpdateSpaceModelRequestSort) SetLevel(v int64) *UpdateSpaceModelRequestSort {
	s.Level = &v
	return s
}

type UpdateSpaceModelShrinkRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 物理空间类型
	SpaceType *string `json:"SpaceType,omitempty" xml:"SpaceType,omitempty"`
	// 资源uuid
	SpaceModelId *string `json:"SpaceModelId,omitempty" xml:"SpaceModelId,omitempty"`
	// 层级
	SortShrink *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
}

func (s UpdateSpaceModelShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpaceModelShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSpaceModelShrinkRequest) SetInstanceId(v string) *UpdateSpaceModelShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateSpaceModelShrinkRequest) SetSpaceType(v string) *UpdateSpaceModelShrinkRequest {
	s.SpaceType = &v
	return s
}

func (s *UpdateSpaceModelShrinkRequest) SetSpaceModelId(v string) *UpdateSpaceModelShrinkRequest {
	s.SpaceModelId = &v
	return s
}

func (s *UpdateSpaceModelShrinkRequest) SetSortShrink(v string) *UpdateSpaceModelShrinkRequest {
	s.SortShrink = &v
	return s
}

type UpdateSpaceModelResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSpaceModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpaceModelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSpaceModelResponseBody) SetRequestId(v string) *UpdateSpaceModelResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSpaceModelResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSpaceModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSpaceModelResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpaceModelResponse) GoString() string {
	return s.String()
}

func (s *UpdateSpaceModelResponse) SetHeaders(v map[string]*string) *UpdateSpaceModelResponse {
	s.Headers = v
	return s
}

func (s *UpdateSpaceModelResponse) SetBody(v *UpdateSpaceModelResponseBody) *UpdateSpaceModelResponse {
	s.Body = v
	return s
}

type ListRegionsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegionsResponseBody) SetData(v string) *ListRegionsResponseBody {
	s.Data = &v
	return s
}

func (s *ListRegionsResponseBody) SetSuccess(v bool) *ListRegionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListRegionsResponseBody) SetMessage(v string) *ListRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListRegionsResponseBody) SetCode(v string) *ListRegionsResponseBody {
	s.Code = &v
	return s
}

type ListRegionsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type DeleteSpaceModelRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例 ID。
	SpaceModelId *string `json:"SpaceModelId,omitempty" xml:"SpaceModelId,omitempty"`
}

func (s DeleteSpaceModelRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpaceModelRequest) GoString() string {
	return s.String()
}

func (s *DeleteSpaceModelRequest) SetInstanceId(v string) *DeleteSpaceModelRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSpaceModelRequest) SetSpaceModelId(v string) *DeleteSpaceModelRequest {
	s.SpaceModelId = &v
	return s
}

type DeleteSpaceModelResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSpaceModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpaceModelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSpaceModelResponseBody) SetRequestId(v string) *DeleteSpaceModelResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSpaceModelResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSpaceModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSpaceModelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpaceModelResponse) GoString() string {
	return s.String()
}

func (s *DeleteSpaceModelResponse) SetHeaders(v map[string]*string) *DeleteSpaceModelResponse {
	s.Headers = v
	return s
}

func (s *DeleteSpaceModelResponse) SetBody(v *DeleteSpaceModelResponseBody) *DeleteSpaceModelResponse {
	s.Body = v
	return s
}

type DisableNotificationRequest struct {
	// 到期时间
	ExpiryTime *string `json:"ExpiryTime,omitempty" xml:"ExpiryTime,omitempty"`
	// 关闭原因
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// 关闭通知的对象
	List []*DisableNotificationRequestList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableNotificationRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableNotificationRequest) GoString() string {
	return s.String()
}

func (s *DisableNotificationRequest) SetExpiryTime(v string) *DisableNotificationRequest {
	s.ExpiryTime = &v
	return s
}

func (s *DisableNotificationRequest) SetReason(v string) *DisableNotificationRequest {
	s.Reason = &v
	return s
}

func (s *DisableNotificationRequest) SetList(v []*DisableNotificationRequestList) *DisableNotificationRequest {
	s.List = v
	return s
}

func (s *DisableNotificationRequest) SetInstanceId(v string) *DisableNotificationRequest {
	s.InstanceId = &v
	return s
}

type DisableNotificationRequestList struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 监控项ID
	MonitorItemId *string `json:"MonitorItemId,omitempty" xml:"MonitorItemId,omitempty"`
	// 聚合数据ID
	AggregateDataId *string `json:"AggregateDataId,omitempty" xml:"AggregateDataId,omitempty"`
	// 类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 专线ID
	DedicatedLineId *string `json:"DedicatedLineId,omitempty" xml:"DedicatedLineId,omitempty"`
	// 端口集ID
	PortCollectionId *string `json:"PortCollectionId,omitempty" xml:"PortCollectionId,omitempty"`
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DisableNotificationRequestList) String() string {
	return tea.Prettify(s)
}

func (s DisableNotificationRequestList) GoString() string {
	return s.String()
}

func (s *DisableNotificationRequestList) SetDeviceId(v string) *DisableNotificationRequestList {
	s.DeviceId = &v
	return s
}

func (s *DisableNotificationRequestList) SetMonitorItemId(v string) *DisableNotificationRequestList {
	s.MonitorItemId = &v
	return s
}

func (s *DisableNotificationRequestList) SetAggregateDataId(v string) *DisableNotificationRequestList {
	s.AggregateDataId = &v
	return s
}

func (s *DisableNotificationRequestList) SetType(v string) *DisableNotificationRequestList {
	s.Type = &v
	return s
}

func (s *DisableNotificationRequestList) SetDedicatedLineId(v string) *DisableNotificationRequestList {
	s.DedicatedLineId = &v
	return s
}

func (s *DisableNotificationRequestList) SetPortCollectionId(v string) *DisableNotificationRequestList {
	s.PortCollectionId = &v
	return s
}

func (s *DisableNotificationRequestList) SetAppId(v string) *DisableNotificationRequestList {
	s.AppId = &v
	return s
}

type DisableNotificationResponseBody struct {
	// request id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableNotificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableNotificationResponseBody) GoString() string {
	return s.String()
}

func (s *DisableNotificationResponseBody) SetRequestId(v string) *DisableNotificationResponseBody {
	s.RequestId = &v
	return s
}

type DisableNotificationResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableNotificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableNotificationResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableNotificationResponse) GoString() string {
	return s.String()
}

func (s *DisableNotificationResponse) SetHeaders(v map[string]*string) *DisableNotificationResponse {
	s.Headers = v
	return s
}

func (s *DisableNotificationResponse) SetBody(v *DisableNotificationResponseBody) *DisableNotificationResponse {
	s.Body = v
	return s
}

type GetDeviceConfigDiffRequest struct {
	// 实例 ID。
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 查询日期1，格式 yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 查询日期2，格式 yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetDeviceConfigDiffRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceConfigDiffRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigDiffRequest) SetDeviceId(v string) *GetDeviceConfigDiffRequest {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceConfigDiffRequest) SetStartDate(v string) *GetDeviceConfigDiffRequest {
	s.StartDate = &v
	return s
}

func (s *GetDeviceConfigDiffRequest) SetEndDate(v string) *GetDeviceConfigDiffRequest {
	s.EndDate = &v
	return s
}

func (s *GetDeviceConfigDiffRequest) SetInstanceId(v string) *GetDeviceConfigDiffRequest {
	s.InstanceId = &v
	return s
}

type GetDeviceConfigDiffResponseBody struct {
	// Id of the request
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceConfigDiff *GetDeviceConfigDiffResponseBodyDeviceConfigDiff `json:"DeviceConfigDiff,omitempty" xml:"DeviceConfigDiff,omitempty" type:"Struct"`
}

func (s GetDeviceConfigDiffResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceConfigDiffResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigDiffResponseBody) SetRequestId(v string) *GetDeviceConfigDiffResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceConfigDiffResponseBody) SetDeviceConfigDiff(v *GetDeviceConfigDiffResponseBodyDeviceConfigDiff) *GetDeviceConfigDiffResponseBody {
	s.DeviceConfigDiff = v
	return s
}

type GetDeviceConfigDiffResponseBodyDeviceConfigDiff struct {
	// 差异提取
	ExtractDiff *string `json:"ExtractDiff,omitempty" xml:"ExtractDiff,omitempty"`
	// 全量比对
	TotalDiff *string `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
}

func (s GetDeviceConfigDiffResponseBodyDeviceConfigDiff) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceConfigDiffResponseBodyDeviceConfigDiff) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigDiffResponseBodyDeviceConfigDiff) SetExtractDiff(v string) *GetDeviceConfigDiffResponseBodyDeviceConfigDiff {
	s.ExtractDiff = &v
	return s
}

func (s *GetDeviceConfigDiffResponseBodyDeviceConfigDiff) SetTotalDiff(v string) *GetDeviceConfigDiffResponseBodyDeviceConfigDiff {
	s.TotalDiff = &v
	return s
}

type GetDeviceConfigDiffResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeviceConfigDiffResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceConfigDiffResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceConfigDiffResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigDiffResponse) SetHeaders(v map[string]*string) *GetDeviceConfigDiffResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceConfigDiffResponse) SetBody(v *GetDeviceConfigDiffResponseBody) *GetDeviceConfigDiffResponse {
	s.Body = v
	return s
}

type GetInspectionTaskRequest struct {
	// 巡检项ID
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInspectionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInspectionTaskRequest) GoString() string {
	return s.String()
}

func (s *GetInspectionTaskRequest) SetTaskId(v string) *GetInspectionTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetInspectionTaskRequest) SetInstanceId(v string) *GetInspectionTaskRequest {
	s.InstanceId = &v
	return s
}

type GetInspectionTaskResponseBody struct {
	// 物理空间
	Space *string `json:"Space,omitempty" xml:"Space,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 巡检结果
	InspectionResult *string `json:"InspectionResult,omitempty" xml:"InspectionResult,omitempty"`
	// 告警规则
	InspectionAlarmRules []*GetInspectionTaskResponseBodyInspectionAlarmRules `json:"InspectionAlarmRules,omitempty" xml:"InspectionAlarmRules,omitempty" type:"Repeated"`
	// IP地址
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// 主机名
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// 任务状态
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// 巡检项ID
	ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// 巡检结束时间
	ExecutionEndTime *string `json:"ExecutionEndTime,omitempty" xml:"ExecutionEndTime,omitempty"`
	// 巡检开始时间
	ExecutionBeginTime *string `json:"ExecutionBeginTime,omitempty" xml:"ExecutionBeginTime,omitempty"`
	// 型号
	Model []*string `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	// 巡检项名字
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 模板ID
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetInspectionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInspectionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetInspectionTaskResponseBody) SetSpace(v string) *GetInspectionTaskResponseBody {
	s.Space = &v
	return s
}

func (s *GetInspectionTaskResponseBody) SetRequestId(v string) *GetInspectionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInspectionTaskResponseBody) SetDeviceId(v string) *GetInspectionTaskResponseBody {
	s.DeviceId = &v
	return s
}

func (s *GetInspectionTaskResponseBody) SetInspectionResult(v string) *GetInspectionTaskResponseBody {
	s.InspectionResult = &v
	return s
}

func (s *GetInspectionTaskResponseBody) SetInspectionAlarmRules(v []*GetInspectionTaskResponseBodyInspectionAlarmRules) *GetInspectionTaskResponseBody {
	s.InspectionAlarmRules = v
	return s
}

func (s *GetInspectionTaskResponseBody) SetIP(v string) *GetInspectionTaskResponseBody {
	s.IP = &v
	return s
}

func (s *GetInspectionTaskResponseBody) SetHostName(v string) *GetInspectionTaskResponseBody {
	s.HostName = &v
	return s
}

func (s *GetInspectionTaskResponseBody) SetVendor(v string) *GetInspectionTaskResponseBody {
	s.Vendor = &v
	return s
}

func (s *GetInspectionTaskResponseBody) SetTaskStatus(v string) *GetInspectionTaskResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetInspectionTaskResponseBody) SetItemId(v string) *GetInspectionTaskResponseBody {
	s.ItemId = &v
	return s
}

func (s *GetInspectionTaskResponseBody) SetExecutionEndTime(v string) *GetInspectionTaskResponseBody {
	s.ExecutionEndTime = &v
	return s
}

func (s *GetInspectionTaskResponseBody) SetExecutionBeginTime(v string) *GetInspectionTaskResponseBody {
	s.ExecutionBeginTime = &v
	return s
}

func (s *GetInspectionTaskResponseBody) SetModel(v []*string) *GetInspectionTaskResponseBody {
	s.Model = v
	return s
}

func (s *GetInspectionTaskResponseBody) SetItemName(v string) *GetInspectionTaskResponseBody {
	s.ItemName = &v
	return s
}

func (s *GetInspectionTaskResponseBody) SetErrorCode(v string) *GetInspectionTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetInspectionTaskResponseBody) SetScriptId(v string) *GetInspectionTaskResponseBody {
	s.ScriptId = &v
	return s
}

func (s *GetInspectionTaskResponseBody) SetTaskId(v string) *GetInspectionTaskResponseBody {
	s.TaskId = &v
	return s
}

type GetInspectionTaskResponseBodyInspectionAlarmRules struct {
	// 告警符号
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// 告警操作符
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// 告警值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// 告警实际值
	ActualValue *string `json:"ActualValue,omitempty" xml:"ActualValue,omitempty"`
	// 告警级别
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s GetInspectionTaskResponseBodyInspectionAlarmRules) String() string {
	return tea.Prettify(s)
}

func (s GetInspectionTaskResponseBodyInspectionAlarmRules) GoString() string {
	return s.String()
}

func (s *GetInspectionTaskResponseBodyInspectionAlarmRules) SetExpression(v string) *GetInspectionTaskResponseBodyInspectionAlarmRules {
	s.Expression = &v
	return s
}

func (s *GetInspectionTaskResponseBodyInspectionAlarmRules) SetOperator(v string) *GetInspectionTaskResponseBodyInspectionAlarmRules {
	s.Operator = &v
	return s
}

func (s *GetInspectionTaskResponseBodyInspectionAlarmRules) SetValue(v string) *GetInspectionTaskResponseBodyInspectionAlarmRules {
	s.Value = &v
	return s
}

func (s *GetInspectionTaskResponseBodyInspectionAlarmRules) SetActualValue(v string) *GetInspectionTaskResponseBodyInspectionAlarmRules {
	s.ActualValue = &v
	return s
}

func (s *GetInspectionTaskResponseBodyInspectionAlarmRules) SetLevel(v string) *GetInspectionTaskResponseBodyInspectionAlarmRules {
	s.Level = &v
	return s
}

type GetInspectionTaskResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInspectionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInspectionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInspectionTaskResponse) GoString() string {
	return s.String()
}

func (s *GetInspectionTaskResponse) SetHeaders(v map[string]*string) *GetInspectionTaskResponse {
	s.Headers = v
	return s
}

func (s *GetInspectionTaskResponse) SetBody(v *GetInspectionTaskResponseBody) *GetInspectionTaskResponse {
	s.Body = v
	return s
}

type ListAlarmStatusRequest struct {
	// 标记当前开始读取的位置，置空表示从头开始
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 物理空间
	Space *string `json:"Space,omitempty" xml:"Space,omitempty"`
	// 设备形态
	DeviceForm *string `json:"DeviceForm,omitempty" xml:"DeviceForm,omitempty"`
	// 告警状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 监控项ID
	MonitorItemId *string `json:"MonitorItemId,omitempty" xml:"MonitorItemId,omitempty"`
	// 数据类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 专线名称
	DedicatedLineName *string `json:"DedicatedLineName,omitempty" xml:"DedicatedLineName,omitempty"`
	// 设备所属地域
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ListAlarmStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmStatusRequest) GoString() string {
	return s.String()
}

func (s *ListAlarmStatusRequest) SetNextToken(v string) *ListAlarmStatusRequest {
	s.NextToken = &v
	return s
}

func (s *ListAlarmStatusRequest) SetMaxResults(v int32) *ListAlarmStatusRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAlarmStatusRequest) SetSpace(v string) *ListAlarmStatusRequest {
	s.Space = &v
	return s
}

func (s *ListAlarmStatusRequest) SetDeviceForm(v string) *ListAlarmStatusRequest {
	s.DeviceForm = &v
	return s
}

func (s *ListAlarmStatusRequest) SetStatus(v string) *ListAlarmStatusRequest {
	s.Status = &v
	return s
}

func (s *ListAlarmStatusRequest) SetDeviceId(v string) *ListAlarmStatusRequest {
	s.DeviceId = &v
	return s
}

func (s *ListAlarmStatusRequest) SetMonitorItemId(v string) *ListAlarmStatusRequest {
	s.MonitorItemId = &v
	return s
}

func (s *ListAlarmStatusRequest) SetType(v string) *ListAlarmStatusRequest {
	s.Type = &v
	return s
}

func (s *ListAlarmStatusRequest) SetInstanceId(v string) *ListAlarmStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAlarmStatusRequest) SetDedicatedLineName(v string) *ListAlarmStatusRequest {
	s.DedicatedLineName = &v
	return s
}

func (s *ListAlarmStatusRequest) SetRegion(v string) *ListAlarmStatusRequest {
	s.Region = &v
	return s
}

type ListAlarmStatusResponseBody struct {
	// TotalCount本次请求条件下的数据总量，此参数为可选参数，默认可不返回
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// MaxResults本次请求所返回的最大记录条数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 告警状态列表
	AlarmStatus []*ListAlarmStatusResponseBodyAlarmStatus `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty" type:"Repeated"`
	// 告警状态统计
	Statistics []*ListAlarmStatusResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s ListAlarmStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlarmStatusResponseBody) SetTotalCount(v int32) *ListAlarmStatusResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAlarmStatusResponseBody) SetRequestId(v string) *ListAlarmStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlarmStatusResponseBody) SetNextToken(v string) *ListAlarmStatusResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAlarmStatusResponseBody) SetMaxResults(v int32) *ListAlarmStatusResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAlarmStatusResponseBody) SetAlarmStatus(v []*ListAlarmStatusResponseBodyAlarmStatus) *ListAlarmStatusResponseBody {
	s.AlarmStatus = v
	return s
}

func (s *ListAlarmStatusResponseBody) SetStatistics(v []*ListAlarmStatusResponseBodyStatistics) *ListAlarmStatusResponseBody {
	s.Statistics = v
	return s
}

type ListAlarmStatusResponseBodyAlarmStatus struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 监控项ID
	MonitorItemId *string `json:"MonitorItemId,omitempty" xml:"MonitorItemId,omitempty"`
	// 检测时间
	CollectionTime *string `json:"CollectionTime,omitempty" xml:"CollectionTime,omitempty"`
	// 接收时间
	ReceiveTime *string `json:"ReceiveTime,omitempty" xml:"ReceiveTime,omitempty"`
	// 命中告警规则
	AlarmRule *string `json:"AlarmRule,omitempty" xml:"AlarmRule,omitempty"`
	// 告警状态
	AlarmStatus *string `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	// 采集结果
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// 异常数据项
	AbnormalDataItem *string `json:"AbnormalDataItem,omitempty" xml:"AbnormalDataItem,omitempty"`
	// 索引
	UniqueKey *string `json:"UniqueKey,omitempty" xml:"UniqueKey,omitempty"`
	// 采集状态码
	ResponseCode *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	// 设备
	ResourceDevice *ListAlarmStatusResponseBodyAlarmStatusResourceDevice `json:"ResourceDevice,omitempty" xml:"ResourceDevice,omitempty" type:"Struct"`
	// 监控项
	MonitorItem *ListAlarmStatusResponseBodyAlarmStatusMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Struct"`
	// 首次异常时间
	FirstAbnormalTime *string `json:"FirstAbnormalTime,omitempty" xml:"FirstAbnormalTime,omitempty"`
	// 告警开关配置
	NotificationSwitch *ListAlarmStatusResponseBodyAlarmStatusNotificationSwitch `json:"NotificationSwitch,omitempty" xml:"NotificationSwitch,omitempty" type:"Struct"`
	// 聚合数据ID
	AggregateDataId *string `json:"AggregateDataId,omitempty" xml:"AggregateDataId,omitempty"`
	// 聚合数据
	AggregateData *ListAlarmStatusResponseBodyAlarmStatusAggregateData `json:"AggregateData,omitempty" xml:"AggregateData,omitempty" type:"Struct"`
	// 专线ID
	DedicatedLineId *string `json:"DedicatedLineId,omitempty" xml:"DedicatedLineId,omitempty"`
	// 专线
	DedicatedLine *ListAlarmStatusResponseBodyAlarmStatusDedicatedLine `json:"DedicatedLine,omitempty" xml:"DedicatedLine,omitempty" type:"Struct"`
	// 端口集ID
	PortCollectionId *string `json:"PortCollectionId,omitempty" xml:"PortCollectionId,omitempty"`
	// 端口集
	PortCollection *ListAlarmStatusResponseBodyAlarmStatusPortCollection `json:"PortCollection,omitempty" xml:"PortCollection,omitempty" type:"Struct"`
	// 采集探针IP
	AgentIp *string `json:"AgentIp,omitempty" xml:"AgentIp,omitempty"`
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用资源
	ResourceApp *ListAlarmStatusResponseBodyAlarmStatusResourceApp `json:"ResourceApp,omitempty" xml:"ResourceApp,omitempty" type:"Struct"`
}

func (s ListAlarmStatusResponseBodyAlarmStatus) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmStatusResponseBodyAlarmStatus) GoString() string {
	return s.String()
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetDeviceId(v string) *ListAlarmStatusResponseBodyAlarmStatus {
	s.DeviceId = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetMonitorItemId(v string) *ListAlarmStatusResponseBodyAlarmStatus {
	s.MonitorItemId = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetCollectionTime(v string) *ListAlarmStatusResponseBodyAlarmStatus {
	s.CollectionTime = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetReceiveTime(v string) *ListAlarmStatusResponseBodyAlarmStatus {
	s.ReceiveTime = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetAlarmRule(v string) *ListAlarmStatusResponseBodyAlarmStatus {
	s.AlarmRule = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetAlarmStatus(v string) *ListAlarmStatusResponseBodyAlarmStatus {
	s.AlarmStatus = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetResult(v string) *ListAlarmStatusResponseBodyAlarmStatus {
	s.Result = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetAbnormalDataItem(v string) *ListAlarmStatusResponseBodyAlarmStatus {
	s.AbnormalDataItem = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetUniqueKey(v string) *ListAlarmStatusResponseBodyAlarmStatus {
	s.UniqueKey = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetResponseCode(v string) *ListAlarmStatusResponseBodyAlarmStatus {
	s.ResponseCode = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetResourceDevice(v *ListAlarmStatusResponseBodyAlarmStatusResourceDevice) *ListAlarmStatusResponseBodyAlarmStatus {
	s.ResourceDevice = v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetMonitorItem(v *ListAlarmStatusResponseBodyAlarmStatusMonitorItem) *ListAlarmStatusResponseBodyAlarmStatus {
	s.MonitorItem = v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetFirstAbnormalTime(v string) *ListAlarmStatusResponseBodyAlarmStatus {
	s.FirstAbnormalTime = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetNotificationSwitch(v *ListAlarmStatusResponseBodyAlarmStatusNotificationSwitch) *ListAlarmStatusResponseBodyAlarmStatus {
	s.NotificationSwitch = v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetAggregateDataId(v string) *ListAlarmStatusResponseBodyAlarmStatus {
	s.AggregateDataId = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetAggregateData(v *ListAlarmStatusResponseBodyAlarmStatusAggregateData) *ListAlarmStatusResponseBodyAlarmStatus {
	s.AggregateData = v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetDedicatedLineId(v string) *ListAlarmStatusResponseBodyAlarmStatus {
	s.DedicatedLineId = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetDedicatedLine(v *ListAlarmStatusResponseBodyAlarmStatusDedicatedLine) *ListAlarmStatusResponseBodyAlarmStatus {
	s.DedicatedLine = v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetPortCollectionId(v string) *ListAlarmStatusResponseBodyAlarmStatus {
	s.PortCollectionId = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetPortCollection(v *ListAlarmStatusResponseBodyAlarmStatusPortCollection) *ListAlarmStatusResponseBodyAlarmStatus {
	s.PortCollection = v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetAgentIp(v string) *ListAlarmStatusResponseBodyAlarmStatus {
	s.AgentIp = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetAppId(v string) *ListAlarmStatusResponseBodyAlarmStatus {
	s.AppId = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatus) SetResourceApp(v *ListAlarmStatusResponseBodyAlarmStatusResourceApp) *ListAlarmStatusResponseBodyAlarmStatus {
	s.ResourceApp = v
	return s
}

type ListAlarmStatusResponseBodyAlarmStatusResourceDevice struct {
	// 设备名
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
}

func (s ListAlarmStatusResponseBodyAlarmStatusResourceDevice) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmStatusResponseBodyAlarmStatusResourceDevice) GoString() string {
	return s.String()
}

func (s *ListAlarmStatusResponseBodyAlarmStatusResourceDevice) SetHostName(v string) *ListAlarmStatusResponseBodyAlarmStatusResourceDevice {
	s.HostName = &v
	return s
}

type ListAlarmStatusResponseBodyAlarmStatusMonitorItem struct {
	// 监控项名称
	MonitorItemName *string `json:"MonitorItemName,omitempty" xml:"MonitorItemName,omitempty"`
}

func (s ListAlarmStatusResponseBodyAlarmStatusMonitorItem) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmStatusResponseBodyAlarmStatusMonitorItem) GoString() string {
	return s.String()
}

func (s *ListAlarmStatusResponseBodyAlarmStatusMonitorItem) SetMonitorItemName(v string) *ListAlarmStatusResponseBodyAlarmStatusMonitorItem {
	s.MonitorItemName = &v
	return s
}

type ListAlarmStatusResponseBodyAlarmStatusNotificationSwitch struct {
	// 关闭原因
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// 关闭到期时间
	ExpiryTime *string `json:"ExpiryTime,omitempty" xml:"ExpiryTime,omitempty"`
}

func (s ListAlarmStatusResponseBodyAlarmStatusNotificationSwitch) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmStatusResponseBodyAlarmStatusNotificationSwitch) GoString() string {
	return s.String()
}

func (s *ListAlarmStatusResponseBodyAlarmStatusNotificationSwitch) SetReason(v string) *ListAlarmStatusResponseBodyAlarmStatusNotificationSwitch {
	s.Reason = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatusNotificationSwitch) SetExpiryTime(v string) *ListAlarmStatusResponseBodyAlarmStatusNotificationSwitch {
	s.ExpiryTime = &v
	return s
}

type ListAlarmStatusResponseBodyAlarmStatusAggregateData struct {
	// 聚合数据名称
	AggregateDataName *string `json:"AggregateDataName,omitempty" xml:"AggregateDataName,omitempty"`
	// 数据项
	DataItem *string `json:"DataItem,omitempty" xml:"DataItem,omitempty"`
}

func (s ListAlarmStatusResponseBodyAlarmStatusAggregateData) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmStatusResponseBodyAlarmStatusAggregateData) GoString() string {
	return s.String()
}

func (s *ListAlarmStatusResponseBodyAlarmStatusAggregateData) SetAggregateDataName(v string) *ListAlarmStatusResponseBodyAlarmStatusAggregateData {
	s.AggregateDataName = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatusAggregateData) SetDataItem(v string) *ListAlarmStatusResponseBodyAlarmStatusAggregateData {
	s.DataItem = &v
	return s
}

type ListAlarmStatusResponseBodyAlarmStatusDedicatedLine struct {
	// 专线名称
	DedicatedLineName *string `json:"DedicatedLineName,omitempty" xml:"DedicatedLineName,omitempty"`
}

func (s ListAlarmStatusResponseBodyAlarmStatusDedicatedLine) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmStatusResponseBodyAlarmStatusDedicatedLine) GoString() string {
	return s.String()
}

func (s *ListAlarmStatusResponseBodyAlarmStatusDedicatedLine) SetDedicatedLineName(v string) *ListAlarmStatusResponseBodyAlarmStatusDedicatedLine {
	s.DedicatedLineName = &v
	return s
}

type ListAlarmStatusResponseBodyAlarmStatusPortCollection struct {
	// 端口集名称
	PortCollectionName *string `json:"PortCollectionName,omitempty" xml:"PortCollectionName,omitempty"`
}

func (s ListAlarmStatusResponseBodyAlarmStatusPortCollection) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmStatusResponseBodyAlarmStatusPortCollection) GoString() string {
	return s.String()
}

func (s *ListAlarmStatusResponseBodyAlarmStatusPortCollection) SetPortCollectionName(v string) *ListAlarmStatusResponseBodyAlarmStatusPortCollection {
	s.PortCollectionName = &v
	return s
}

type ListAlarmStatusResponseBodyAlarmStatusResourceApp struct {
	// 监控域名
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 端口
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// 资源类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 所属探针
	SecurityDomain *string `json:"SecurityDomain,omitempty" xml:"SecurityDomain,omitempty"`
}

func (s ListAlarmStatusResponseBodyAlarmStatusResourceApp) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmStatusResponseBodyAlarmStatusResourceApp) GoString() string {
	return s.String()
}

func (s *ListAlarmStatusResponseBodyAlarmStatusResourceApp) SetDomain(v string) *ListAlarmStatusResponseBodyAlarmStatusResourceApp {
	s.Domain = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatusResourceApp) SetAppId(v string) *ListAlarmStatusResponseBodyAlarmStatusResourceApp {
	s.AppId = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatusResourceApp) SetPort(v string) *ListAlarmStatusResponseBodyAlarmStatusResourceApp {
	s.Port = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatusResourceApp) SetType(v string) *ListAlarmStatusResponseBodyAlarmStatusResourceApp {
	s.Type = &v
	return s
}

func (s *ListAlarmStatusResponseBodyAlarmStatusResourceApp) SetSecurityDomain(v string) *ListAlarmStatusResponseBodyAlarmStatusResourceApp {
	s.SecurityDomain = &v
	return s
}

type ListAlarmStatusResponseBodyStatistics struct {
	// 数量
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// 告警状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAlarmStatusResponseBodyStatistics) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmStatusResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *ListAlarmStatusResponseBodyStatistics) SetCount(v int64) *ListAlarmStatusResponseBodyStatistics {
	s.Count = &v
	return s
}

func (s *ListAlarmStatusResponseBodyStatistics) SetStatus(v string) *ListAlarmStatusResponseBodyStatistics {
	s.Status = &v
	return s
}

type ListAlarmStatusResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAlarmStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlarmStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmStatusResponse) GoString() string {
	return s.String()
}

func (s *ListAlarmStatusResponse) SetHeaders(v map[string]*string) *ListAlarmStatusResponse {
	s.Headers = v
	return s
}

func (s *ListAlarmStatusResponse) SetBody(v *ListAlarmStatusResponseBody) *ListAlarmStatusResponse {
	s.Body = v
	return s
}

type ListArchitectureAttributeRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 架构id
	ArchitectureId *string `json:"ArchitectureId,omitempty" xml:"ArchitectureId,omitempty"`
	// 角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 厂商
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ListArchitectureAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListArchitectureAttributeRequest) GoString() string {
	return s.String()
}

func (s *ListArchitectureAttributeRequest) SetInstanceId(v string) *ListArchitectureAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ListArchitectureAttributeRequest) SetArchitectureId(v string) *ListArchitectureAttributeRequest {
	s.ArchitectureId = &v
	return s
}

func (s *ListArchitectureAttributeRequest) SetRole(v string) *ListArchitectureAttributeRequest {
	s.Role = &v
	return s
}

func (s *ListArchitectureAttributeRequest) SetVendor(v string) *ListArchitectureAttributeRequest {
	s.Vendor = &v
	return s
}

type ListArchitectureAttributeResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 架构对象
	Architecture []*ListArchitectureAttributeResponseBodyArchitecture `json:"Architecture,omitempty" xml:"Architecture,omitempty" type:"Repeated"`
}

func (s ListArchitectureAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListArchitectureAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ListArchitectureAttributeResponseBody) SetRequestId(v string) *ListArchitectureAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListArchitectureAttributeResponseBody) SetArchitecture(v []*ListArchitectureAttributeResponseBodyArchitecture) *ListArchitectureAttributeResponseBody {
	s.Architecture = v
	return s
}

type ListArchitectureAttributeResponseBodyArchitecture struct {
	// 角色
	Role []*string `json:"Role,omitempty" xml:"Role,omitempty" type:"Repeated"`
	// 厂商
	Vendor []*string `json:"Vendor,omitempty" xml:"Vendor,omitempty" type:"Repeated"`
	// 型号
	Model []*string `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
}

func (s ListArchitectureAttributeResponseBodyArchitecture) String() string {
	return tea.Prettify(s)
}

func (s ListArchitectureAttributeResponseBodyArchitecture) GoString() string {
	return s.String()
}

func (s *ListArchitectureAttributeResponseBodyArchitecture) SetRole(v []*string) *ListArchitectureAttributeResponseBodyArchitecture {
	s.Role = v
	return s
}

func (s *ListArchitectureAttributeResponseBodyArchitecture) SetVendor(v []*string) *ListArchitectureAttributeResponseBodyArchitecture {
	s.Vendor = v
	return s
}

func (s *ListArchitectureAttributeResponseBodyArchitecture) SetModel(v []*string) *ListArchitectureAttributeResponseBodyArchitecture {
	s.Model = v
	return s
}

type ListArchitectureAttributeResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListArchitectureAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListArchitectureAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListArchitectureAttributeResponse) GoString() string {
	return s.String()
}

func (s *ListArchitectureAttributeResponse) SetHeaders(v map[string]*string) *ListArchitectureAttributeResponse {
	s.Headers = v
	return s
}

func (s *ListArchitectureAttributeResponse) SetBody(v *ListArchitectureAttributeResponseBody) *ListArchitectureAttributeResponse {
	s.Body = v
	return s
}

type DeleteOsVersionRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例 ID。
	OsVersionId *string `json:"OsVersionId,omitempty" xml:"OsVersionId,omitempty"`
}

func (s DeleteOsVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteOsVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteOsVersionRequest) SetInstanceId(v string) *DeleteOsVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteOsVersionRequest) SetOsVersionId(v string) *DeleteOsVersionRequest {
	s.OsVersionId = &v
	return s
}

type DeleteOsVersionResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOsVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteOsVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOsVersionResponseBody) SetRequestId(v string) *DeleteOsVersionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteOsVersionResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteOsVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteOsVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOsVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteOsVersionResponse) SetHeaders(v map[string]*string) *DeleteOsVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteOsVersionResponse) SetBody(v *DeleteOsVersionResponseBody) *DeleteOsVersionResponse {
	s.Body = v
	return s
}

type DeleteScheduleTypeRequest struct {
	// 实例 ID。
	ScheduleTypeId *string `json:"ScheduleTypeId,omitempty" xml:"ScheduleTypeId,omitempty"`
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteScheduleTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleTypeRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTypeRequest) SetScheduleTypeId(v string) *DeleteScheduleTypeRequest {
	s.ScheduleTypeId = &v
	return s
}

func (s *DeleteScheduleTypeRequest) SetInstanceId(v string) *DeleteScheduleTypeRequest {
	s.InstanceId = &v
	return s
}

type DeleteScheduleTypeResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScheduleTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTypeResponseBody) SetRequestId(v string) *DeleteScheduleTypeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteScheduleTypeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteScheduleTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteScheduleTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleTypeResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTypeResponse) SetHeaders(v map[string]*string) *DeleteScheduleTypeResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduleTypeResponse) SetBody(v *DeleteScheduleTypeResponseBody) *DeleteScheduleTypeResponse {
	s.Body = v
	return s
}

type DeletePhysicalSpaceRequest struct {
	// 实例 ID。
	PhysicalSpaceId *string `json:"PhysicalSpaceId,omitempty" xml:"PhysicalSpaceId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeletePhysicalSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePhysicalSpaceRequest) GoString() string {
	return s.String()
}

func (s *DeletePhysicalSpaceRequest) SetPhysicalSpaceId(v string) *DeletePhysicalSpaceRequest {
	s.PhysicalSpaceId = &v
	return s
}

func (s *DeletePhysicalSpaceRequest) SetInstanceId(v string) *DeletePhysicalSpaceRequest {
	s.InstanceId = &v
	return s
}

type DeletePhysicalSpaceResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePhysicalSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePhysicalSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePhysicalSpaceResponseBody) SetRequestId(v string) *DeletePhysicalSpaceResponseBody {
	s.RequestId = &v
	return s
}

type DeletePhysicalSpaceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePhysicalSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePhysicalSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePhysicalSpaceResponse) GoString() string {
	return s.String()
}

func (s *DeletePhysicalSpaceResponse) SetHeaders(v map[string]*string) *DeletePhysicalSpaceResponse {
	s.Headers = v
	return s
}

func (s *DeletePhysicalSpaceResponse) SetBody(v *DeletePhysicalSpaceResponseBody) *DeletePhysicalSpaceResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cmn"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConfigurationSpecificationWithOptions(tmpReq *CreateConfigurationSpecificationRequest, runtime *util.RuntimeOptions) (_result *CreateConfigurationSpecificationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateConfigurationSpecificationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RelatedVariate)) {
		request.RelatedVariateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedVariate, tea.String("RelatedVariate"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateConfigurationSpecificationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateConfigurationSpecification"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConfigurationSpecification(request *CreateConfigurationSpecificationRequest) (_result *CreateConfigurationSpecificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateConfigurationSpecificationResponse{}
	_body, _err := client.CreateConfigurationSpecificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIpBlockWithOptions(request *CreateIpBlockRequest, runtime *util.RuntimeOptions) (_result *CreateIpBlockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateIpBlockResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateIpBlock"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIpBlock(request *CreateIpBlockRequest) (_result *CreateIpBlockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIpBlockResponse{}
	_body, _err := client.CreateIpBlockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateConfigurationVariateWithOptions(request *UpdateConfigurationVariateRequest, runtime *util.RuntimeOptions) (_result *UpdateConfigurationVariateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateConfigurationVariateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateConfigurationVariate"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateConfigurationVariate(request *UpdateConfigurationVariateRequest) (_result *UpdateConfigurationVariateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateConfigurationVariateResponse{}
	_body, _err := client.UpdateConfigurationVariateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetScheduleTypeWithOptions(request *GetScheduleTypeRequest, runtime *util.RuntimeOptions) (_result *GetScheduleTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetScheduleTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetScheduleType"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetScheduleType(request *GetScheduleTypeRequest) (_result *GetScheduleTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetScheduleTypeResponse{}
	_body, _err := client.GetScheduleTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDeviceWithOptions(request *DeleteDeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDevice"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDevice(request *DeleteDeviceRequest) (_result *DeleteDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDeviceResponse{}
	_body, _err := client.DeleteDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeviceValuesWithOptions(request *ListDeviceValuesRequest, runtime *util.RuntimeOptions) (_result *ListDeviceValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListDeviceValuesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDeviceValues"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceValues(request *ListDeviceValuesRequest) (_result *ListDeviceValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeviceValuesResponse{}
	_body, _err := client.ListDeviceValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIpBlockRecordWithOptions(request *GetIpBlockRecordRequest, runtime *util.RuntimeOptions) (_result *GetIpBlockRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetIpBlockRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetIpBlockRecord"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIpBlockRecord(request *GetIpBlockRecordRequest) (_result *GetIpBlockRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIpBlockRecordResponse{}
	_body, _err := client.GetIpBlockRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSpaceModelsWithOptions(request *ListSpaceModelsRequest, runtime *util.RuntimeOptions) (_result *ListSpaceModelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListSpaceModelsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSpaceModels"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSpaceModels(request *ListSpaceModelsRequest) (_result *ListSpaceModelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSpaceModelsResponse{}
	_body, _err := client.ListSpaceModelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDevicePropertiesWithOptions(request *ListDevicePropertiesRequest, runtime *util.RuntimeOptions) (_result *ListDevicePropertiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListDevicePropertiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDeviceProperties"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceProperties(request *ListDevicePropertiesRequest) (_result *ListDevicePropertiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDevicePropertiesResponse{}
	_body, _err := client.ListDevicePropertiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInspectionDevicesWithOptions(tmpReq *ListInspectionDevicesRequest, runtime *util.RuntimeOptions) (_result *ListInspectionDevicesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListInspectionDevicesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Model)) {
		request.ModelShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Model, tea.String("Model"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListInspectionDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInspectionDevices"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInspectionDevices(request *ListInspectionDevicesRequest) (_result *ListInspectionDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInspectionDevicesResponse{}
	_body, _err := client.ListInspectionDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInspectionTasksWithOptions(request *ListInspectionTasksRequest, runtime *util.RuntimeOptions) (_result *ListInspectionTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListInspectionTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInspectionTasks"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInspectionTasks(request *ListInspectionTasksRequest) (_result *ListInspectionTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInspectionTasksResponse{}
	_body, _err := client.ListInspectionTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListScheduleWorkersWithOptions(request *ListScheduleWorkersRequest, runtime *util.RuntimeOptions) (_result *ListScheduleWorkersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListScheduleWorkersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListScheduleWorkers"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListScheduleWorkers(request *ListScheduleWorkersRequest) (_result *ListScheduleWorkersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListScheduleWorkersResponse{}
	_body, _err := client.ListScheduleWorkersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProjectProgressWithOptions(request *UpdateProjectProgressRequest, runtime *util.RuntimeOptions) (_result *UpdateProjectProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateProjectProgressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateProjectProgress"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProjectProgress(request *UpdateProjectProgressRequest) (_result *UpdateProjectProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateProjectProgressResponse{}
	_body, _err := client.UpdateProjectProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDeviceResourceWithOptions(tmpReq *UpdateDeviceResourceRequest, runtime *util.RuntimeOptions) (_result *UpdateDeviceResourceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateDeviceResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceResourceIds)) {
		request.DeviceResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceResourceIds, tea.String("DeviceResourceIds"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDeviceResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDeviceResource"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDeviceResource(request *UpdateDeviceResourceRequest) (_result *UpdateDeviceResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDeviceResourceResponse{}
	_body, _err := client.UpdateDeviceResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDevicePropertyWithOptions(request *GetDevicePropertyRequest, runtime *util.RuntimeOptions) (_result *GetDevicePropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetDevicePropertyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDeviceProperty"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceProperty(request *GetDevicePropertyRequest) (_result *GetDevicePropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDevicePropertyResponse{}
	_body, _err := client.GetDevicePropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourceTypesWithOptions(runtime *util.RuntimeOptions) (_result *ListResourceTypesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListResourceTypes"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourceTypes() (_result *ListResourceTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.ListResourceTypesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSetupProjectWithOptions(request *GetSetupProjectRequest, runtime *util.RuntimeOptions) (_result *GetSetupProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetSetupProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSetupProject"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSetupProject(request *GetSetupProjectRequest) (_result *GetSetupProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSetupProjectResponse{}
	_body, _err := client.GetSetupProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConfigurationVariateWithOptions(request *ListConfigurationVariateRequest, runtime *util.RuntimeOptions) (_result *ListConfigurationVariateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListConfigurationVariateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListConfigurationVariate"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConfigurationVariate(request *ListConfigurationVariateRequest) (_result *ListConfigurationVariateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConfigurationVariateResponse{}
	_body, _err := client.ListConfigurationVariateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSpaceModelWithOptions(request *CreateSpaceModelRequest, runtime *util.RuntimeOptions) (_result *CreateSpaceModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSpaceModelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSpaceModel"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSpaceModel(request *CreateSpaceModelRequest) (_result *CreateSpaceModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSpaceModelResponse{}
	_body, _err := client.CreateSpaceModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteScheduleWorkerWithOptions(request *DeleteScheduleWorkerRequest, runtime *util.RuntimeOptions) (_result *DeleteScheduleWorkerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteScheduleWorkerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteScheduleWorker"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteScheduleWorker(request *DeleteScheduleWorkerRequest) (_result *DeleteScheduleWorkerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScheduleWorkerResponse{}
	_body, _err := client.DeleteScheduleWorkerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDedicatedLinesWithOptions(request *ListDedicatedLinesRequest, runtime *util.RuntimeOptions) (_result *ListDedicatedLinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListDedicatedLinesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDedicatedLines"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDedicatedLines(request *ListDedicatedLinesRequest) (_result *ListDedicatedLinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDedicatedLinesResponse{}
	_body, _err := client.ListDedicatedLinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInformationKeyActionWithOptions(request *UpdateInformationKeyActionRequest, runtime *util.RuntimeOptions) (_result *UpdateInformationKeyActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateInformationKeyActionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateInformationKeyAction"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInformationKeyAction(request *UpdateInformationKeyActionRequest) (_result *UpdateInformationKeyActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInformationKeyActionResponse{}
	_body, _err := client.UpdateInformationKeyActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRealtimeTaskWithOptions(request *GetRealtimeTaskRequest, runtime *util.RuntimeOptions) (_result *GetRealtimeTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetRealtimeTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRealtimeTask"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRealtimeTask(request *GetRealtimeTaskRequest) (_result *GetRealtimeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRealtimeTaskResponse{}
	_body, _err := client.GetRealtimeTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListScheduleTypesWithOptions(request *ListScheduleTypesRequest, runtime *util.RuntimeOptions) (_result *ListScheduleTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListScheduleTypesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListScheduleTypes"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListScheduleTypes(request *ListScheduleTypesRequest) (_result *ListScheduleTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListScheduleTypesResponse{}
	_body, _err := client.ListScheduleTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateScheduleTypeWithOptions(tmpReq *CreateScheduleTypeRequest, runtime *util.RuntimeOptions) (_result *CreateScheduleTypeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateScheduleTypeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RelatedWorker)) {
		request.RelatedWorkerShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedWorker, tea.String("RelatedWorker"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateScheduleTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateScheduleType"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateScheduleType(request *CreateScheduleTypeRequest) (_result *CreateScheduleTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScheduleTypeResponse{}
	_body, _err := client.CreateScheduleTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetScheduleWorkerWithOptions(request *GetScheduleWorkerRequest, runtime *util.RuntimeOptions) (_result *GetScheduleWorkerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetScheduleWorkerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetScheduleWorker"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetScheduleWorker(request *GetScheduleWorkerRequest) (_result *GetScheduleWorkerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetScheduleWorkerResponse{}
	_body, _err := client.GetScheduleWorkerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateScheduleWorkerWithOptions(request *CreateScheduleWorkerRequest, runtime *util.RuntimeOptions) (_result *CreateScheduleWorkerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateScheduleWorkerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateScheduleWorker"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateScheduleWorker(request *CreateScheduleWorkerRequest) (_result *CreateScheduleWorkerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScheduleWorkerResponse{}
	_body, _err := client.CreateScheduleWorkerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConfigurationVariateWithOptions(request *CreateConfigurationVariateRequest, runtime *util.RuntimeOptions) (_result *CreateConfigurationVariateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateConfigurationVariateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateConfigurationVariate"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConfigurationVariate(request *CreateConfigurationVariateRequest) (_result *CreateConfigurationVariateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateConfigurationVariateResponse{}
	_body, _err := client.CreateConfigurationVariateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSpaceModelSortWithOptions(request *GetSpaceModelSortRequest, runtime *util.RuntimeOptions) (_result *GetSpaceModelSortResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetSpaceModelSortResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSpaceModelSort"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSpaceModelSort(request *GetSpaceModelSortRequest) (_result *GetSpaceModelSortResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSpaceModelSortResponse{}
	_body, _err := client.GetSpaceModelSortWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRealtimeTaskWithOptions(request *CreateRealtimeTaskRequest, runtime *util.RuntimeOptions) (_result *CreateRealtimeTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRealtimeTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRealtimeTask"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRealtimeTask(request *CreateRealtimeTaskRequest) (_result *CreateRealtimeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRealtimeTaskResponse{}
	_body, _err := client.CreateRealtimeTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceFormWithOptions(request *GetDeviceFormRequest, runtime *util.RuntimeOptions) (_result *GetDeviceFormResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetDeviceFormResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDeviceForm"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceForm(request *GetDeviceFormRequest) (_result *GetDeviceFormResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceFormResponse{}
	_body, _err := client.GetDeviceFormWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDeviceWithOptions(request *CreateDeviceRequest, runtime *util.RuntimeOptions) (_result *CreateDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDevice"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDevice(request *CreateDeviceRequest) (_result *CreateDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDeviceResponse{}
	_body, _err := client.CreateDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSetupProjectWithOptions(request *CreateSetupProjectRequest, runtime *util.RuntimeOptions) (_result *CreateSetupProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSetupProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSetupProject"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSetupProject(request *CreateSetupProjectRequest) (_result *CreateSetupProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSetupProjectResponse{}
	_body, _err := client.CreateSetupProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDedicatedLineWithOptions(request *UpdateDedicatedLineRequest, runtime *util.RuntimeOptions) (_result *UpdateDedicatedLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDedicatedLineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDedicatedLine"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDedicatedLine(request *UpdateDedicatedLineRequest) (_result *UpdateDedicatedLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDedicatedLineResponse{}
	_body, _err := client.UpdateDedicatedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOsVersionWithOptions(request *GetOsVersionRequest, runtime *util.RuntimeOptions) (_result *GetOsVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetOsVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOsVersion"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOsVersion(request *GetOsVersionRequest) (_result *GetOsVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOsVersionResponse{}
	_body, _err := client.GetOsVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateScheduleDutyWithOptions(tmpReq *UpdateScheduleDutyRequest, runtime *util.RuntimeOptions) (_result *UpdateScheduleDutyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateScheduleDutyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TypeWorkerList)) {
		request.TypeWorkerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TypeWorkerList, tea.String("TypeWorkerList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateScheduleDutyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateScheduleDuty"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateScheduleDuty(request *UpdateScheduleDutyRequest) (_result *UpdateScheduleDutyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateScheduleDutyResponse{}
	_body, _err := client.UpdateScheduleDutyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIpRecordWithOptions(tmpReq *CreateIpRecordRequest, runtime *util.RuntimeOptions) (_result *CreateIpRecordResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateIpRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Device)) {
		request.DeviceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Device, tea.String("Device"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.IpCode)) {
		request.IpCodeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IpCode, tea.String("IpCode"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateIpRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateIpRecord"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIpRecord(request *CreateIpRecordRequest) (_result *CreateIpRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIpRecordResponse{}
	_body, _err := client.CreateIpRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstancesWithOptions(runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &ListInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInstances"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstances() (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListZoneTypesWithOptions(request *ListZoneTypesRequest, runtime *util.RuntimeOptions) (_result *ListZoneTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListZoneTypesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListZoneTypes"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListZoneTypes(request *ListZoneTypesRequest) (_result *ListZoneTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListZoneTypesResponse{}
	_body, _err := client.ListZoneTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePhysicalSpaceWithOptions(request *UpdatePhysicalSpaceRequest, runtime *util.RuntimeOptions) (_result *UpdatePhysicalSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdatePhysicalSpaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdatePhysicalSpace"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePhysicalSpace(request *UpdatePhysicalSpaceRequest) (_result *UpdatePhysicalSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePhysicalSpaceResponse{}
	_body, _err := client.UpdatePhysicalSpaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResourceInstanceWithOptions(tmpReq *UpdateResourceInstanceRequest, runtime *util.RuntimeOptions) (_result *UpdateResourceInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateResourceInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceInformation)) {
		request.ResourceInformationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceInformation, tea.String("ResourceInformation"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateResourceInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateResourceInstance"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateResourceInstance(request *UpdateResourceInstanceRequest) (_result *UpdateResourceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateResourceInstanceResponse{}
	_body, _err := client.UpdateResourceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetScheduleDutyWithOptions(request *GetScheduleDutyRequest, runtime *util.RuntimeOptions) (_result *GetScheduleDutyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetScheduleDutyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetScheduleDuty"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetScheduleDuty(request *GetScheduleDutyRequest) (_result *GetScheduleDutyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetScheduleDutyResponse{}
	_body, _err := client.GetScheduleDutyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConfigurationVariateWithOptions(request *GetConfigurationVariateRequest, runtime *util.RuntimeOptions) (_result *GetConfigurationVariateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetConfigurationVariateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetConfigurationVariate"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConfigurationVariate(request *GetConfigurationVariateRequest) (_result *GetConfigurationVariateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConfigurationVariateResponse{}
	_body, _err := client.GetConfigurationVariateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAlarmStatusWithOptions(request *GetAlarmStatusRequest, runtime *util.RuntimeOptions) (_result *GetAlarmStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetAlarmStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAlarmStatus"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAlarmStatus(request *GetAlarmStatusRequest) (_result *GetAlarmStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAlarmStatusResponse{}
	_body, _err := client.GetAlarmStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIpRecordWithOptions(request *GetIpRecordRequest, runtime *util.RuntimeOptions) (_result *GetIpRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetIpRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetIpRecord"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIpRecord(request *GetIpRecordRequest) (_result *GetIpRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIpRecordResponse{}
	_body, _err := client.GetIpRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListScheduleDutiesWithOptions(request *ListScheduleDutiesRequest, runtime *util.RuntimeOptions) (_result *ListScheduleDutiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListScheduleDutiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListScheduleDuties"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListScheduleDuties(request *ListScheduleDutiesRequest) (_result *ListScheduleDutiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListScheduleDutiesResponse{}
	_body, _err := client.ListScheduleDutiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LockSpaceModelWithOptions(request *LockSpaceModelRequest, runtime *util.RuntimeOptions) (_result *LockSpaceModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &LockSpaceModelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("LockSpaceModel"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LockSpaceModel(request *LockSpaceModelRequest) (_result *LockSpaceModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LockSpaceModelResponse{}
	_body, _err := client.LockSpaceModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResourceInformationWithOptions(tmpReq *UpdateResourceInformationRequest, runtime *util.RuntimeOptions) (_result *UpdateResourceInformationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateResourceInformationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Information)) {
		request.InformationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Information, tea.String("Information"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateResourceInformationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateResourceInformation"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateResourceInformation(request *UpdateResourceInformationRequest) (_result *UpdateResourceInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateResourceInformationResponse{}
	_body, _err := client.UpdateResourceInformationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTimePeriodWithOptions(request *CreateTimePeriodRequest, runtime *util.RuntimeOptions) (_result *CreateTimePeriodResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTimePeriodResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateTimePeriod"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTimePeriod(request *CreateTimePeriodRequest) (_result *CreateTimePeriodResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTimePeriodResponse{}
	_body, _err := client.CreateTimePeriodWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDevicesWithOptions(request *ListDevicesRequest, runtime *util.RuntimeOptions) (_result *ListDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDevices"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDevices(request *ListDevicesRequest) (_result *ListDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDevicesResponse{}
	_body, _err := client.ListDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInspectionTaskReportsWithOptions(request *ListInspectionTaskReportsRequest, runtime *util.RuntimeOptions) (_result *ListInspectionTaskReportsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListInspectionTaskReportsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInspectionTaskReports"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInspectionTaskReports(request *ListInspectionTaskReportsRequest) (_result *ListInspectionTaskReportsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInspectionTaskReportsResponse{}
	_body, _err := client.ListInspectionTaskReportsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMonitorItemWithOptions(request *CreateMonitorItemRequest, runtime *util.RuntimeOptions) (_result *CreateMonitorItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateMonitorItemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateMonitorItem"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMonitorItem(request *CreateMonitorItemRequest) (_result *CreateMonitorItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMonitorItemResponse{}
	_body, _err := client.CreateMonitorItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePhysicalSpaceWithOptions(request *CreatePhysicalSpaceRequest, runtime *util.RuntimeOptions) (_result *CreatePhysicalSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreatePhysicalSpaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreatePhysicalSpace"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePhysicalSpace(request *CreatePhysicalSpaceRequest) (_result *CreatePhysicalSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePhysicalSpaceResponse{}
	_body, _err := client.CreatePhysicalSpaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDevicesWithOptions(request *UpdateDevicesRequest, runtime *util.RuntimeOptions) (_result *UpdateDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDevices"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDevices(request *UpdateDevicesRequest) (_result *UpdateDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDevicesResponse{}
	_body, _err := client.UpdateDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateScheduleTypeWithOptions(tmpReq *UpdateScheduleTypeRequest, runtime *util.RuntimeOptions) (_result *UpdateScheduleTypeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateScheduleTypeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RelatedWorker)) {
		request.RelatedWorkerShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedWorker, tea.String("RelatedWorker"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateScheduleTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateScheduleType"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateScheduleType(request *UpdateScheduleTypeRequest) (_result *UpdateScheduleTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateScheduleTypeResponse{}
	_body, _err := client.UpdateScheduleTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadDeviceResourceWithOptions(tmpReq *DownloadDeviceResourceRequest, runtime *util.RuntimeOptions) (_result *DownloadDeviceResourceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DownloadDeviceResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceResourceIds)) {
		request.DeviceResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceResourceIds, tea.String("DeviceResourceIds"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DownloadDeviceResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DownloadDeviceResource"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadDeviceResource(request *DownloadDeviceResourceRequest) (_result *DownloadDeviceResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadDeviceResourceResponse{}
	_body, _err := client.DownloadDeviceResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOsDownloadPathWithOptions(request *GetOsDownloadPathRequest, runtime *util.RuntimeOptions) (_result *GetOsDownloadPathResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetOsDownloadPathResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOsDownloadPath"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOsDownloadPath(request *GetOsDownloadPathRequest) (_result *GetOsDownloadPathResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOsDownloadPathResponse{}
	_body, _err := client.GetOsDownloadPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConnectionPoliciesWithOptions(request *ListConnectionPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListConnectionPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListConnectionPoliciesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListConnectionPolicies"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConnectionPolicies(request *ListConnectionPoliciesRequest) (_result *ListConnectionPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConnectionPoliciesResponse{}
	_body, _err := client.ListConnectionPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateScheduleWorkerWithOptions(request *UpdateScheduleWorkerRequest, runtime *util.RuntimeOptions) (_result *UpdateScheduleWorkerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateScheduleWorkerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateScheduleWorker"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateScheduleWorker(request *UpdateScheduleWorkerRequest) (_result *UpdateScheduleWorkerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateScheduleWorkerResponse{}
	_body, _err := client.UpdateScheduleWorkerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConfigurationVariateWithOptions(request *DeleteConfigurationVariateRequest, runtime *util.RuntimeOptions) (_result *DeleteConfigurationVariateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteConfigurationVariateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteConfigurationVariate"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConfigurationVariate(request *DeleteConfigurationVariateRequest) (_result *DeleteConfigurationVariateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteConfigurationVariateResponse{}
	_body, _err := client.DeleteConfigurationVariateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateScheduleDutyWithOptions(tmpReq *CreateScheduleDutyRequest, runtime *util.RuntimeOptions) (_result *CreateScheduleDutyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateScheduleDutyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ScheduleTypeIds)) {
		request.ScheduleTypeIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleTypeIds, tea.String("ScheduleTypeIds"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateScheduleDutyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateScheduleDuty"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateScheduleDuty(request *CreateScheduleDutyRequest) (_result *CreateScheduleDutyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScheduleDutyResponse{}
	_body, _err := client.CreateScheduleDutyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPhysicalSpaceWithOptions(request *GetPhysicalSpaceRequest, runtime *util.RuntimeOptions) (_result *GetPhysicalSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetPhysicalSpaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPhysicalSpace"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPhysicalSpace(request *GetPhysicalSpaceRequest) (_result *GetPhysicalSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPhysicalSpaceResponse{}
	_body, _err := client.GetPhysicalSpaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteResourceInformationWithOptions(request *DeleteResourceInformationRequest, runtime *util.RuntimeOptions) (_result *DeleteResourceInformationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteResourceInformationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteResourceInformation"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteResourceInformation(request *DeleteResourceInformationRequest) (_result *DeleteResourceInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteResourceInformationResponse{}
	_body, _err := client.DeleteResourceInformationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSetupProjectWithOptions(request *DeleteSetupProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteSetupProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSetupProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSetupProject"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSetupProject(request *DeleteSetupProjectRequest) (_result *DeleteSetupProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSetupProjectResponse{}
	_body, _err := client.DeleteSetupProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDedicatedLineWithOptions(request *CreateDedicatedLineRequest, runtime *util.RuntimeOptions) (_result *CreateDedicatedLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDedicatedLineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDedicatedLine"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDedicatedLine(request *CreateDedicatedLineRequest) (_result *CreateDedicatedLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDedicatedLineResponse{}
	_body, _err := client.CreateDedicatedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyIPWithOptions(tmpReq *ApplyIPRequest, runtime *util.RuntimeOptions) (_result *ApplyIPResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ApplyIPShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceResourceIds)) {
		request.DeviceResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceResourceIds, tea.String("DeviceResourceIds"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ApplyIPResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ApplyIP"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyIP(request *ApplyIPRequest) (_result *ApplyIPResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyIPResponse{}
	_body, _err := client.ApplyIPWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOsVersionWithOptions(request *UpdateOsVersionRequest, runtime *util.RuntimeOptions) (_result *UpdateOsVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateOsVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateOsVersion"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOsVersion(request *UpdateOsVersionRequest) (_result *UpdateOsVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateOsVersionResponse{}
	_body, _err := client.UpdateOsVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSpaceModelInstanceWithOptions(request *GetSpaceModelInstanceRequest, runtime *util.RuntimeOptions) (_result *GetSpaceModelInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetSpaceModelInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSpaceModelInstance"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSpaceModelInstance(request *GetSpaceModelInstanceRequest) (_result *GetSpaceModelInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSpaceModelInstanceResponse{}
	_body, _err := client.GetSpaceModelInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOsVersionsWithOptions(request *ListOsVersionsRequest, runtime *util.RuntimeOptions) (_result *ListOsVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListOsVersionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOsVersions"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOsVersions(request *ListOsVersionsRequest) (_result *ListOsVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOsVersionsResponse{}
	_body, _err := client.ListOsVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceConfigWithOptions(request *GetDeviceConfigRequest, runtime *util.RuntimeOptions) (_result *GetDeviceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetDeviceConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDeviceConfig"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceConfig(request *GetDeviceConfigRequest) (_result *GetDeviceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceConfigResponse{}
	_body, _err := client.GetDeviceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDedicatedLineWithOptions(request *GetDedicatedLineRequest, runtime *util.RuntimeOptions) (_result *GetDedicatedLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetDedicatedLineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDedicatedLine"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDedicatedLine(request *GetDedicatedLineRequest) (_result *GetDedicatedLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDedicatedLineResponse{}
	_body, _err := client.GetDedicatedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceResourceWithOptions(request *GetDeviceResourceRequest, runtime *util.RuntimeOptions) (_result *GetDeviceResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetDeviceResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDeviceResource"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceResource(request *GetDeviceResourceRequest) (_result *GetDeviceResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceResourceResponse{}
	_body, _err := client.GetDeviceResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDedicatedLineWithOptions(request *DeleteDedicatedLineRequest, runtime *util.RuntimeOptions) (_result *DeleteDedicatedLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDedicatedLineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDedicatedLine"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDedicatedLine(request *DeleteDedicatedLineRequest) (_result *DeleteDedicatedLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDedicatedLineResponse{}
	_body, _err := client.DeleteDedicatedLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIpWithOptions(request *ListIpRequest, runtime *util.RuntimeOptions) (_result *ListIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListIpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListIp"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIp(request *ListIpRequest) (_result *ListIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIpResponse{}
	_body, _err := client.ListIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConfigurationSpecificationsWithOptions(request *ListConfigurationSpecificationsRequest, runtime *util.RuntimeOptions) (_result *ListConfigurationSpecificationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListConfigurationSpecificationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListConfigurationSpecifications"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConfigurationSpecifications(request *ListConfigurationSpecificationsRequest) (_result *ListConfigurationSpecificationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConfigurationSpecificationsResponse{}
	_body, _err := client.ListConfigurationSpecificationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableNotificationWithOptions(request *EnableNotificationRequest, runtime *util.RuntimeOptions) (_result *EnableNotificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableNotificationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableNotification"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableNotification(request *EnableNotificationRequest) (_result *EnableNotificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableNotificationResponse{}
	_body, _err := client.EnableNotificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDevicePropertyWithOptions(request *UpdateDevicePropertyRequest, runtime *util.RuntimeOptions) (_result *UpdateDevicePropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDevicePropertyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDeviceProperty"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDeviceProperty(request *UpdateDevicePropertyRequest) (_result *UpdateDevicePropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDevicePropertyResponse{}
	_body, _err := client.UpdateDevicePropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNotificationHistoriesWithOptions(request *ListNotificationHistoriesRequest, runtime *util.RuntimeOptions) (_result *ListNotificationHistoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListNotificationHistoriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListNotificationHistories"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNotificationHistories(request *ListNotificationHistoriesRequest) (_result *ListNotificationHistoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNotificationHistoriesResponse{}
	_body, _err := client.ListNotificationHistoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDevicePropertyWithOptions(request *DeleteDevicePropertyRequest, runtime *util.RuntimeOptions) (_result *DeleteDevicePropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDevicePropertyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDeviceProperty"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDeviceProperty(request *DeleteDevicePropertyRequest) (_result *DeleteDevicePropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDevicePropertyResponse{}
	_body, _err := client.DeleteDevicePropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourceInstancesWithOptions(request *ListResourceInstancesRequest, runtime *util.RuntimeOptions) (_result *ListResourceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListResourceInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListResourceInstances"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourceInstances(request *ListResourceInstancesRequest) (_result *ListResourceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceInstancesResponse{}
	_body, _err := client.ListResourceInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIpBlocksWithOptions(request *ListIpBlocksRequest, runtime *util.RuntimeOptions) (_result *ListIpBlocksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListIpBlocksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListIpBlocks"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIpBlocks(request *ListIpBlocksRequest) (_result *ListIpBlocksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIpBlocksResponse{}
	_body, _err := client.ListIpBlocksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeviceResourcesWithOptions(request *ListDeviceResourcesRequest, runtime *util.RuntimeOptions) (_result *ListDeviceResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListDeviceResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDeviceResources"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceResources(request *ListDeviceResourcesRequest) (_result *ListDeviceResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeviceResourcesResponse{}
	_body, _err := client.ListDeviceResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourceInformationsWithOptions(request *ListResourceInformationsRequest, runtime *util.RuntimeOptions) (_result *ListResourceInformationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListResourceInformationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListResourceInformations"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourceInformations(request *ListResourceInformationsRequest) (_result *ListResourceInformationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceInformationsResponse{}
	_body, _err := client.ListResourceInformationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeviceFormsWithOptions(request *ListDeviceFormsRequest, runtime *util.RuntimeOptions) (_result *ListDeviceFormsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListDeviceFormsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDeviceForms"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceForms(request *ListDeviceFormsRequest) (_result *ListDeviceFormsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeviceFormsResponse{}
	_body, _err := client.ListDeviceFormsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConfigurationSpecificationWithOptions(request *GetConfigurationSpecificationRequest, runtime *util.RuntimeOptions) (_result *GetConfigurationSpecificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetConfigurationSpecificationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetConfigurationSpecification"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConfigurationSpecification(request *GetConfigurationSpecificationRequest) (_result *GetConfigurationSpecificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConfigurationSpecificationResponse{}
	_body, _err := client.GetConfigurationSpecificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteScheduleDutyWithOptions(tmpReq *DeleteScheduleDutyRequest, runtime *util.RuntimeOptions) (_result *DeleteScheduleDutyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteScheduleDutyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TypeWorkerList)) {
		request.TypeWorkerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TypeWorkerList, tea.String("TypeWorkerList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteScheduleDutyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteScheduleDuty"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteScheduleDuty(request *DeleteScheduleDutyRequest) (_result *DeleteScheduleDutyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScheduleDutyResponse{}
	_body, _err := client.DeleteScheduleDutyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadScheduleDutyWithOptions(tmpReq *UploadScheduleDutyRequest, runtime *util.RuntimeOptions) (_result *UploadScheduleDutyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UploadScheduleDutyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ScheduleDuty)) {
		request.ScheduleDutyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleDuty, tea.String("ScheduleDuty"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UploadScheduleDutyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UploadScheduleDuty"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadScheduleDuty(request *UploadScheduleDutyRequest) (_result *UploadScheduleDutyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadScheduleDutyResponse{}
	_body, _err := client.UploadScheduleDutyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAlarmStatusHistoriesWithOptions(request *ListAlarmStatusHistoriesRequest, runtime *util.RuntimeOptions) (_result *ListAlarmStatusHistoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListAlarmStatusHistoriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAlarmStatusHistories"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAlarmStatusHistories(request *ListAlarmStatusHistoriesRequest) (_result *ListAlarmStatusHistoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAlarmStatusHistoriesResponse{}
	_body, _err := client.ListAlarmStatusHistoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSpaceModelWithOptions(request *GetSpaceModelRequest, runtime *util.RuntimeOptions) (_result *GetSpaceModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetSpaceModelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSpaceModel"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSpaceModel(request *GetSpaceModelRequest) (_result *GetSpaceModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSpaceModelResponse{}
	_body, _err := client.GetSpaceModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDeviceFormWithOptions(request *CreateDeviceFormRequest, runtime *util.RuntimeOptions) (_result *CreateDeviceFormResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDeviceFormResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDeviceForm"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDeviceForm(request *CreateDeviceFormRequest) (_result *CreateDeviceFormResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDeviceFormResponse{}
	_body, _err := client.CreateDeviceFormWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPhysicalSpacesWithOptions(request *ListPhysicalSpacesRequest, runtime *util.RuntimeOptions) (_result *ListPhysicalSpacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListPhysicalSpacesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPhysicalSpaces"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPhysicalSpaces(request *ListPhysicalSpacesRequest) (_result *ListPhysicalSpacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPhysicalSpacesResponse{}
	_body, _err := client.ListPhysicalSpacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMonitorDataWithOptions(request *ListMonitorDataRequest, runtime *util.RuntimeOptions) (_result *ListMonitorDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListMonitorDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListMonitorData"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMonitorData(request *ListMonitorDataRequest) (_result *ListMonitorDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMonitorDataResponse{}
	_body, _err := client.ListMonitorDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateResourceInformationWithOptions(request *CreateResourceInformationRequest, runtime *util.RuntimeOptions) (_result *CreateResourceInformationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateResourceInformationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateResourceInformation"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateResourceInformation(request *CreateResourceInformationRequest) (_result *CreateResourceInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateResourceInformationResponse{}
	_body, _err := client.CreateResourceInformationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSpaceModelInstanceWithOptions(request *UpdateSpaceModelInstanceRequest, runtime *util.RuntimeOptions) (_result *UpdateSpaceModelInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateSpaceModelInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateSpaceModelInstance"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSpaceModelInstance(request *UpdateSpaceModelInstanceRequest) (_result *UpdateSpaceModelInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSpaceModelInstanceResponse{}
	_body, _err := client.UpdateSpaceModelInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIpRecordWithOptions(request *UpdateIpRecordRequest, runtime *util.RuntimeOptions) (_result *UpdateIpRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateIpRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateIpRecord"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIpRecord(request *UpdateIpRecordRequest) (_result *UpdateIpRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIpRecordResponse{}
	_body, _err := client.UpdateIpRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseIPWithOptions(tmpReq *ReleaseIPRequest, runtime *util.RuntimeOptions) (_result *ReleaseIPResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ReleaseIPShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceResourceIds)) {
		request.DeviceResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceResourceIds, tea.String("DeviceResourceIds"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleaseIPResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleaseIP"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseIP(request *ReleaseIPRequest) (_result *ReleaseIPResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseIPResponse{}
	_body, _err := client.ReleaseIPWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDeviceResourceWithOptions(request *DeleteDeviceResourceRequest, runtime *util.RuntimeOptions) (_result *DeleteDeviceResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDeviceResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDeviceResource"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDeviceResource(request *DeleteDeviceResourceRequest) (_result *DeleteDeviceResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDeviceResourceResponse{}
	_body, _err := client.DeleteDeviceResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIpBlockWithOptions(request *GetIpBlockRequest, runtime *util.RuntimeOptions) (_result *GetIpBlockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetIpBlockResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetIpBlock"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIpBlock(request *GetIpBlockRequest) (_result *GetIpBlockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIpBlockResponse{}
	_body, _err := client.GetIpBlockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIpBlockWithOptions(request *DeleteIpBlockRequest, runtime *util.RuntimeOptions) (_result *DeleteIpBlockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteIpBlockResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteIpBlock"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIpBlock(request *DeleteIpBlockRequest) (_result *DeleteIpBlockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIpBlockResponse{}
	_body, _err := client.DeleteIpBlockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInspectionTaskWithOptions(request *DeleteInspectionTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteInspectionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteInspectionTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteInspectionTask"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInspectionTask(request *DeleteInspectionTaskRequest) (_result *DeleteInspectionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInspectionTaskResponse{}
	_body, _err := client.DeleteInspectionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConfigurationSpecificationWithOptions(request *DeleteConfigurationSpecificationRequest, runtime *util.RuntimeOptions) (_result *DeleteConfigurationSpecificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteConfigurationSpecificationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteConfigurationSpecification"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConfigurationSpecification(request *DeleteConfigurationSpecificationRequest) (_result *DeleteConfigurationSpecificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteConfigurationSpecificationResponse{}
	_body, _err := client.DeleteConfigurationSpecificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBusinessTypesWithOptions(request *ListBusinessTypesRequest, runtime *util.RuntimeOptions) (_result *ListBusinessTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListBusinessTypesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListBusinessTypes"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBusinessTypes(request *ListBusinessTypesRequest) (_result *ListBusinessTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBusinessTypesResponse{}
	_body, _err := client.ListBusinessTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSetupProjectsWithOptions(request *ListSetupProjectsRequest, runtime *util.RuntimeOptions) (_result *ListSetupProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListSetupProjectsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSetupProjects"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSetupProjects(request *ListSetupProjectsRequest) (_result *ListSetupProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSetupProjectsResponse{}
	_body, _err := client.ListSetupProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTasksHistoriesWithOptions(request *ListTasksHistoriesRequest, runtime *util.RuntimeOptions) (_result *ListTasksHistoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListTasksHistoriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTasksHistories"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTasksHistories(request *ListTasksHistoriesRequest) (_result *ListTasksHistoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTasksHistoriesResponse{}
	_body, _err := client.ListTasksHistoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateConfigurationSpecificationWithOptions(tmpReq *UpdateConfigurationSpecificationRequest, runtime *util.RuntimeOptions) (_result *UpdateConfigurationSpecificationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateConfigurationSpecificationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RelatedVariate)) {
		request.RelatedVariateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedVariate, tea.String("RelatedVariate"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateConfigurationSpecificationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateConfigurationSpecification"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateConfigurationSpecification(request *UpdateConfigurationSpecificationRequest) (_result *UpdateConfigurationSpecificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateConfigurationSpecificationResponse{}
	_body, _err := client.UpdateConfigurationSpecificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOsVersionWithOptions(request *CreateOsVersionRequest, runtime *util.RuntimeOptions) (_result *CreateOsVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateOsVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateOsVersion"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOsVersion(request *CreateOsVersionRequest) (_result *CreateOsVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOsVersionResponse{}
	_body, _err := client.CreateOsVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIpBlockRecordWithOptions(tmpReq *CreateIpBlockRecordRequest, runtime *util.RuntimeOptions) (_result *CreateIpBlockRecordResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateIpBlockRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ZoneLayer)) {
		request.ZoneLayerShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ZoneLayer, tea.String("ZoneLayer"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.BusinessType)) {
		request.BusinessTypeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BusinessType, tea.String("BusinessType"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.IpBlockCode)) {
		request.IpBlockCodeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IpBlockCode, tea.String("IpBlockCode"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateIpBlockRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateIpBlockRecord"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIpBlockRecord(request *CreateIpBlockRecordRequest) (_result *CreateIpBlockRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIpBlockRecordResponse{}
	_body, _err := client.CreateIpBlockRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIpBlockWithOptions(request *UpdateIpBlockRequest, runtime *util.RuntimeOptions) (_result *UpdateIpBlockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateIpBlockResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateIpBlock"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIpBlock(request *UpdateIpBlockRequest) (_result *UpdateIpBlockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIpBlockResponse{}
	_body, _err := client.UpdateIpBlockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDevicePropertyWithOptions(request *CreateDevicePropertyRequest, runtime *util.RuntimeOptions) (_result *CreateDevicePropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDevicePropertyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDeviceProperty"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDeviceProperty(request *CreateDevicePropertyRequest) (_result *CreateDevicePropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDevicePropertyResponse{}
	_body, _err := client.CreateDevicePropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIpBlockRecordWithOptions(request *UpdateIpBlockRecordRequest, runtime *util.RuntimeOptions) (_result *UpdateIpBlockRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateIpBlockRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateIpBlockRecord"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIpBlockRecord(request *UpdateIpBlockRecordRequest) (_result *UpdateIpBlockRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIpBlockRecordResponse{}
	_body, _err := client.UpdateIpBlockRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RetryTasksWithOptions(tmpReq *RetryTasksRequest, runtime *util.RuntimeOptions) (_result *RetryTasksResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RetryTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RetryTasks)) {
		request.RetryTasksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RetryTasks, tea.String("RetryTasks"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RetryTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RetryTasks"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RetryTasks(request *RetryTasksRequest) (_result *RetryTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RetryTasksResponse{}
	_body, _err := client.RetryTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPhysicalSpaceTopoWithOptions(request *GetPhysicalSpaceTopoRequest, runtime *util.RuntimeOptions) (_result *GetPhysicalSpaceTopoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPhysicalSpaceTopoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPhysicalSpaceTopo"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPhysicalSpaceTopo(request *GetPhysicalSpaceTopoRequest) (_result *GetPhysicalSpaceTopoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPhysicalSpaceTopoResponse{}
	_body, _err := client.GetPhysicalSpaceTopoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOssPolicyWithOptions(request *GetOssPolicyRequest, runtime *util.RuntimeOptions) (_result *GetOssPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetOssPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOssPolicy"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOssPolicy(request *GetOssPolicyRequest) (_result *GetOssPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOssPolicyResponse{}
	_body, _err := client.GetOssPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDeviceFormWithOptions(request *DeleteDeviceFormRequest, runtime *util.RuntimeOptions) (_result *DeleteDeviceFormResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDeviceFormResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDeviceForm"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDeviceForm(request *DeleteDeviceFormRequest) (_result *DeleteDeviceFormResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDeviceFormResponse{}
	_body, _err := client.DeleteDeviceFormWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDeviceFormWithOptions(tmpReq *UpdateDeviceFormRequest, runtime *util.RuntimeOptions) (_result *UpdateDeviceFormResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateDeviceFormShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AttributeList)) {
		request.AttributeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AttributeList, tea.String("AttributeList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDeviceFormResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDeviceForm"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDeviceForm(request *UpdateDeviceFormRequest) (_result *UpdateDeviceFormResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDeviceFormResponse{}
	_body, _err := client.UpdateDeviceFormWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDeviceWithOptions(request *UpdateDeviceRequest, runtime *util.RuntimeOptions) (_result *UpdateDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDevice"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDevice(request *UpdateDeviceRequest) (_result *UpdateDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDeviceResponse{}
	_body, _err := client.UpdateDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceWithOptions(request *GetDeviceRequest, runtime *util.RuntimeOptions) (_result *GetDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDevice"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDevice(request *GetDeviceRequest) (_result *GetDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceResponse{}
	_body, _err := client.GetDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSetupProjectWithOptions(tmpReq *UpdateSetupProjectRequest, runtime *util.RuntimeOptions) (_result *UpdateSetupProjectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateSetupProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Packages)) {
		request.PackagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Packages, tea.String("Packages"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateSetupProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateSetupProject"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSetupProject(request *UpdateSetupProjectRequest) (_result *UpdateSetupProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSetupProjectResponse{}
	_body, _err := client.UpdateSetupProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSpaceModelWithOptions(tmpReq *UpdateSpaceModelRequest, runtime *util.RuntimeOptions) (_result *UpdateSpaceModelResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateSpaceModelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Sort)) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, tea.String("Sort"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateSpaceModelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateSpaceModel"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSpaceModel(request *UpdateSpaceModelRequest) (_result *UpdateSpaceModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSpaceModelResponse{}
	_body, _err := client.UpdateSpaceModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRegionsWithOptions(runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &ListRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRegions"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSpaceModelWithOptions(request *DeleteSpaceModelRequest, runtime *util.RuntimeOptions) (_result *DeleteSpaceModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSpaceModelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSpaceModel"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSpaceModel(request *DeleteSpaceModelRequest) (_result *DeleteSpaceModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSpaceModelResponse{}
	_body, _err := client.DeleteSpaceModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableNotificationWithOptions(request *DisableNotificationRequest, runtime *util.RuntimeOptions) (_result *DisableNotificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableNotificationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableNotification"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableNotification(request *DisableNotificationRequest) (_result *DisableNotificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableNotificationResponse{}
	_body, _err := client.DisableNotificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceConfigDiffWithOptions(request *GetDeviceConfigDiffRequest, runtime *util.RuntimeOptions) (_result *GetDeviceConfigDiffResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetDeviceConfigDiffResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDeviceConfigDiff"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceConfigDiff(request *GetDeviceConfigDiffRequest) (_result *GetDeviceConfigDiffResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceConfigDiffResponse{}
	_body, _err := client.GetDeviceConfigDiffWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInspectionTaskWithOptions(request *GetInspectionTaskRequest, runtime *util.RuntimeOptions) (_result *GetInspectionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetInspectionTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetInspectionTask"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInspectionTask(request *GetInspectionTaskRequest) (_result *GetInspectionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInspectionTaskResponse{}
	_body, _err := client.GetInspectionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAlarmStatusWithOptions(request *ListAlarmStatusRequest, runtime *util.RuntimeOptions) (_result *ListAlarmStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListAlarmStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAlarmStatus"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAlarmStatus(request *ListAlarmStatusRequest) (_result *ListAlarmStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAlarmStatusResponse{}
	_body, _err := client.ListAlarmStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListArchitectureAttributeWithOptions(request *ListArchitectureAttributeRequest, runtime *util.RuntimeOptions) (_result *ListArchitectureAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListArchitectureAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListArchitectureAttribute"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListArchitectureAttribute(request *ListArchitectureAttributeRequest) (_result *ListArchitectureAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListArchitectureAttributeResponse{}
	_body, _err := client.ListArchitectureAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteOsVersionWithOptions(request *DeleteOsVersionRequest, runtime *util.RuntimeOptions) (_result *DeleteOsVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteOsVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteOsVersion"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteOsVersion(request *DeleteOsVersionRequest) (_result *DeleteOsVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteOsVersionResponse{}
	_body, _err := client.DeleteOsVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteScheduleTypeWithOptions(request *DeleteScheduleTypeRequest, runtime *util.RuntimeOptions) (_result *DeleteScheduleTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteScheduleTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteScheduleType"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteScheduleType(request *DeleteScheduleTypeRequest) (_result *DeleteScheduleTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScheduleTypeResponse{}
	_body, _err := client.DeleteScheduleTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePhysicalSpaceWithOptions(request *DeletePhysicalSpaceRequest, runtime *util.RuntimeOptions) (_result *DeletePhysicalSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeletePhysicalSpaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeletePhysicalSpace"), tea.String("2020-08-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePhysicalSpace(request *DeletePhysicalSpaceRequest) (_result *DeletePhysicalSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePhysicalSpaceResponse{}
	_body, _err := client.DeletePhysicalSpaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
