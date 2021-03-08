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
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// 巡检状态
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s *ListInspectionTasksRequest) SetIp(v string) *ListInspectionTasksRequest {
	s.Ip = &v
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

type CreateRealtimeTaskRequest struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 模板执行脚本
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

type UpdateDevicesRequest struct {
	// 设备ID
	DeviceIds []*string `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty" type:"Repeated"`
	// 登录类型
	LoginType *string `json:"LoginType,omitempty" xml:"LoginType,omitempty"`
	// 登录账号
	LoginUsername *string `json:"LoginUsername,omitempty" xml:"LoginUsername,omitempty"`
	// 登录密码
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
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
