// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-pop/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ACS struct {
	ACSQuotaId         *string   `json:"ACSQuotaId,omitempty" xml:"ACSQuotaId,omitempty"`
	AssociatedProducts []*string `json:"AssociatedProducts,omitempty" xml:"AssociatedProducts,omitempty" type:"Repeated"`
}

func (s ACS) String() string {
	return tea.Prettify(s)
}

func (s ACS) GoString() string {
	return s.String()
}

func (s *ACS) SetACSQuotaId(v string) *ACS {
	s.ACSQuotaId = &v
	return s
}

func (s *ACS) SetAssociatedProducts(v []*string) *ACS {
	s.AssociatedProducts = v
	return s
}

type Action struct {
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
}

func (s Action) String() string {
	return tea.Prettify(s)
}

func (s Action) GoString() string {
	return s.String()
}

func (s *Action) SetActionType(v string) *Action {
	s.ActionType = &v
	return s
}

type AlgorithmSpec struct {
	CodeDir *Location `json:"CodeDir,omitempty" xml:"CodeDir,omitempty"`
	// This parameter is required.
	Command         []*string                     `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	ComputeResource *AlgorithmSpecComputeResource `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	Customization   *AlgorithmSpecCustomization   `json:"Customization,omitempty" xml:"Customization,omitempty" type:"Struct"`
	HyperParameters []*HyperParameterDefinition   `json:"HyperParameters,omitempty" xml:"HyperParameters,omitempty" type:"Repeated"`
	// This parameter is required.
	Image         *string    `json:"Image,omitempty" xml:"Image,omitempty"`
	InputChannels []*Channel `json:"InputChannels,omitempty" xml:"InputChannels,omitempty" type:"Repeated"`
	// This parameter is required.
	JobType                     *string                           `json:"JobType,omitempty" xml:"JobType,omitempty"`
	MetricDefinitions           []*MetricDefinition               `json:"MetricDefinitions,omitempty" xml:"MetricDefinitions,omitempty" type:"Repeated"`
	OutputChannels              []*Channel                        `json:"OutputChannels,omitempty" xml:"OutputChannels,omitempty" type:"Repeated"`
	ProgressDefinitions         *AlgorithmSpecProgressDefinitions `json:"ProgressDefinitions,omitempty" xml:"ProgressDefinitions,omitempty" type:"Struct"`
	ResourceRequirements        []*ConditionExpression            `json:"ResourceRequirements,omitempty" xml:"ResourceRequirements,omitempty" type:"Repeated"`
	SupportedInstanceTypes      []*string                         `json:"SupportedInstanceTypes,omitempty" xml:"SupportedInstanceTypes,omitempty" type:"Repeated"`
	SupportsDistributedTraining *bool                             `json:"SupportsDistributedTraining,omitempty" xml:"SupportsDistributedTraining,omitempty"`
}

func (s AlgorithmSpec) String() string {
	return tea.Prettify(s)
}

func (s AlgorithmSpec) GoString() string {
	return s.String()
}

func (s *AlgorithmSpec) SetCodeDir(v *Location) *AlgorithmSpec {
	s.CodeDir = v
	return s
}

func (s *AlgorithmSpec) SetCommand(v []*string) *AlgorithmSpec {
	s.Command = v
	return s
}

func (s *AlgorithmSpec) SetComputeResource(v *AlgorithmSpecComputeResource) *AlgorithmSpec {
	s.ComputeResource = v
	return s
}

func (s *AlgorithmSpec) SetCustomization(v *AlgorithmSpecCustomization) *AlgorithmSpec {
	s.Customization = v
	return s
}

func (s *AlgorithmSpec) SetHyperParameters(v []*HyperParameterDefinition) *AlgorithmSpec {
	s.HyperParameters = v
	return s
}

func (s *AlgorithmSpec) SetImage(v string) *AlgorithmSpec {
	s.Image = &v
	return s
}

func (s *AlgorithmSpec) SetInputChannels(v []*Channel) *AlgorithmSpec {
	s.InputChannels = v
	return s
}

func (s *AlgorithmSpec) SetJobType(v string) *AlgorithmSpec {
	s.JobType = &v
	return s
}

func (s *AlgorithmSpec) SetMetricDefinitions(v []*MetricDefinition) *AlgorithmSpec {
	s.MetricDefinitions = v
	return s
}

func (s *AlgorithmSpec) SetOutputChannels(v []*Channel) *AlgorithmSpec {
	s.OutputChannels = v
	return s
}

func (s *AlgorithmSpec) SetProgressDefinitions(v *AlgorithmSpecProgressDefinitions) *AlgorithmSpec {
	s.ProgressDefinitions = v
	return s
}

func (s *AlgorithmSpec) SetResourceRequirements(v []*ConditionExpression) *AlgorithmSpec {
	s.ResourceRequirements = v
	return s
}

func (s *AlgorithmSpec) SetSupportedInstanceTypes(v []*string) *AlgorithmSpec {
	s.SupportedInstanceTypes = v
	return s
}

func (s *AlgorithmSpec) SetSupportsDistributedTraining(v bool) *AlgorithmSpec {
	s.SupportsDistributedTraining = &v
	return s
}

type AlgorithmSpecComputeResource struct {
	// This parameter is required.
	Policy *AlgorithmSpecComputeResourcePolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
}

func (s AlgorithmSpecComputeResource) String() string {
	return tea.Prettify(s)
}

func (s AlgorithmSpecComputeResource) GoString() string {
	return s.String()
}

func (s *AlgorithmSpecComputeResource) SetPolicy(v *AlgorithmSpecComputeResourcePolicy) *AlgorithmSpecComputeResource {
	s.Policy = v
	return s
}

type AlgorithmSpecComputeResourcePolicy struct {
	// This parameter is required.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// This parameter is required.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s AlgorithmSpecComputeResourcePolicy) String() string {
	return tea.Prettify(s)
}

func (s AlgorithmSpecComputeResourcePolicy) GoString() string {
	return s.String()
}

func (s *AlgorithmSpecComputeResourcePolicy) SetValue(v string) *AlgorithmSpecComputeResourcePolicy {
	s.Value = &v
	return s
}

func (s *AlgorithmSpecComputeResourcePolicy) SetVersion(v string) *AlgorithmSpecComputeResourcePolicy {
	s.Version = &v
	return s
}

type AlgorithmSpecCustomization struct {
	CodeDir *bool `json:"CodeDir,omitempty" xml:"CodeDir,omitempty"`
}

func (s AlgorithmSpecCustomization) String() string {
	return tea.Prettify(s)
}

func (s AlgorithmSpecCustomization) GoString() string {
	return s.String()
}

func (s *AlgorithmSpecCustomization) SetCodeDir(v bool) *AlgorithmSpecCustomization {
	s.CodeDir = &v
	return s
}

type AlgorithmSpecProgressDefinitions struct {
	OverallProgress *AlgorithmSpecProgressDefinitionsOverallProgress `json:"OverallProgress,omitempty" xml:"OverallProgress,omitempty" type:"Struct"`
	RemainingTime   *AlgorithmSpecProgressDefinitionsRemainingTime   `json:"RemainingTime,omitempty" xml:"RemainingTime,omitempty" type:"Struct"`
}

func (s AlgorithmSpecProgressDefinitions) String() string {
	return tea.Prettify(s)
}

func (s AlgorithmSpecProgressDefinitions) GoString() string {
	return s.String()
}

func (s *AlgorithmSpecProgressDefinitions) SetOverallProgress(v *AlgorithmSpecProgressDefinitionsOverallProgress) *AlgorithmSpecProgressDefinitions {
	s.OverallProgress = v
	return s
}

func (s *AlgorithmSpecProgressDefinitions) SetRemainingTime(v *AlgorithmSpecProgressDefinitionsRemainingTime) *AlgorithmSpecProgressDefinitions {
	s.RemainingTime = v
	return s
}

type AlgorithmSpecProgressDefinitionsOverallProgress struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Regex       *string `json:"Regex,omitempty" xml:"Regex,omitempty"`
}

func (s AlgorithmSpecProgressDefinitionsOverallProgress) String() string {
	return tea.Prettify(s)
}

func (s AlgorithmSpecProgressDefinitionsOverallProgress) GoString() string {
	return s.String()
}

func (s *AlgorithmSpecProgressDefinitionsOverallProgress) SetDescription(v string) *AlgorithmSpecProgressDefinitionsOverallProgress {
	s.Description = &v
	return s
}

func (s *AlgorithmSpecProgressDefinitionsOverallProgress) SetRegex(v string) *AlgorithmSpecProgressDefinitionsOverallProgress {
	s.Regex = &v
	return s
}

type AlgorithmSpecProgressDefinitionsRemainingTime struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Regex       *string `json:"Regex,omitempty" xml:"Regex,omitempty"`
}

func (s AlgorithmSpecProgressDefinitionsRemainingTime) String() string {
	return tea.Prettify(s)
}

func (s AlgorithmSpecProgressDefinitionsRemainingTime) GoString() string {
	return s.String()
}

func (s *AlgorithmSpecProgressDefinitionsRemainingTime) SetDescription(v string) *AlgorithmSpecProgressDefinitionsRemainingTime {
	s.Description = &v
	return s
}

func (s *AlgorithmSpecProgressDefinitionsRemainingTime) SetRegex(v string) *AlgorithmSpecProgressDefinitionsRemainingTime {
	s.Regex = &v
	return s
}

type AllocateStrategySpec struct {
	NodeSpecs []*NodeSpec `json:"NodeSpecs,omitempty" xml:"NodeSpecs,omitempty" type:"Repeated"`
}

func (s AllocateStrategySpec) String() string {
	return tea.Prettify(s)
}

func (s AllocateStrategySpec) GoString() string {
	return s.String()
}

func (s *AllocateStrategySpec) SetNodeSpecs(v []*NodeSpec) *AllocateStrategySpec {
	s.NodeSpecs = v
	return s
}

type Channel struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name                  *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Properties            map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	Required              *bool                  `json:"Required,omitempty" xml:"Required,omitempty"`
	SupportedChannelTypes []*string              `json:"SupportedChannelTypes,omitempty" xml:"SupportedChannelTypes,omitempty" type:"Repeated"`
}

func (s Channel) String() string {
	return tea.Prettify(s)
}

func (s Channel) GoString() string {
	return s.String()
}

func (s *Channel) SetDescription(v string) *Channel {
	s.Description = &v
	return s
}

func (s *Channel) SetName(v string) *Channel {
	s.Name = &v
	return s
}

func (s *Channel) SetProperties(v map[string]interface{}) *Channel {
	s.Properties = v
	return s
}

func (s *Channel) SetRequired(v bool) *Channel {
	s.Required = &v
	return s
}

func (s *Channel) SetSupportedChannelTypes(v []*string) *Channel {
	s.SupportedChannelTypes = v
	return s
}

type ChannelProperty struct {
	// This parameter is required.
	//
	// example:
	//
	// SKlearn
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Framework
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ChannelProperty) String() string {
	return tea.Prettify(s)
}

func (s ChannelProperty) GoString() string {
	return s.String()
}

func (s *ChannelProperty) SetName(v string) *ChannelProperty {
	s.Name = &v
	return s
}

func (s *ChannelProperty) SetValue(v string) *ChannelProperty {
	s.Value = &v
	return s
}

type ComponentSpec struct {
	CodeDir *Location `json:"CodeDir,omitempty" xml:"CodeDir,omitempty"`
	// This parameter is required.
	Command         *string                     `json:"Command,omitempty" xml:"Command,omitempty"`
	HyperParameters []*HyperParameterDefinition `json:"HyperParameters,omitempty" xml:"HyperParameters,omitempty" type:"Repeated"`
	// This parameter is required.
	Image         *string    `json:"Image,omitempty" xml:"Image,omitempty"`
	InputChannels []*Channel `json:"InputChannels,omitempty" xml:"InputChannels,omitempty" type:"Repeated"`
	// This parameter is required.
	JobType              *string                `json:"JobType,omitempty" xml:"JobType,omitempty"`
	MetricDefinitions    []*MetricDefinition    `json:"MetricDefinitions,omitempty" xml:"MetricDefinitions,omitempty" type:"Repeated"`
	OutputChannels       []*Channel             `json:"OutputChannels,omitempty" xml:"OutputChannels,omitempty" type:"Repeated"`
	ResourceRequirements []*ConditionExpression `json:"ResourceRequirements,omitempty" xml:"ResourceRequirements,omitempty" type:"Repeated"`
}

func (s ComponentSpec) String() string {
	return tea.Prettify(s)
}

func (s ComponentSpec) GoString() string {
	return s.String()
}

func (s *ComponentSpec) SetCodeDir(v *Location) *ComponentSpec {
	s.CodeDir = v
	return s
}

func (s *ComponentSpec) SetCommand(v string) *ComponentSpec {
	s.Command = &v
	return s
}

func (s *ComponentSpec) SetHyperParameters(v []*HyperParameterDefinition) *ComponentSpec {
	s.HyperParameters = v
	return s
}

func (s *ComponentSpec) SetImage(v string) *ComponentSpec {
	s.Image = &v
	return s
}

func (s *ComponentSpec) SetInputChannels(v []*Channel) *ComponentSpec {
	s.InputChannels = v
	return s
}

func (s *ComponentSpec) SetJobType(v string) *ComponentSpec {
	s.JobType = &v
	return s
}

func (s *ComponentSpec) SetMetricDefinitions(v []*MetricDefinition) *ComponentSpec {
	s.MetricDefinitions = v
	return s
}

func (s *ComponentSpec) SetOutputChannels(v []*Channel) *ComponentSpec {
	s.OutputChannels = v
	return s
}

func (s *ComponentSpec) SetResourceRequirements(v []*ConditionExpression) *ComponentSpec {
	s.ResourceRequirements = v
	return s
}

type ConditionExpression struct {
	// This parameter is required.
	//
	// example:
	//
	// SupportedMachineTypes
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// in
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// This parameter is required.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ConditionExpression) String() string {
	return tea.Prettify(s)
}

func (s ConditionExpression) GoString() string {
	return s.String()
}

func (s *ConditionExpression) SetKey(v string) *ConditionExpression {
	s.Key = &v
	return s
}

func (s *ConditionExpression) SetOperator(v string) *ConditionExpression {
	s.Operator = &v
	return s
}

func (s *ConditionExpression) SetValues(v []*string) *ConditionExpression {
	s.Values = v
	return s
}

type Features struct {
	Quota *FeaturesQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
}

func (s Features) String() string {
	return tea.Prettify(s)
}

func (s Features) GoString() string {
	return s.String()
}

func (s *Features) SetQuota(v *FeaturesQuota) *Features {
	s.Quota = v
	return s
}

type FeaturesQuota struct {
	// example:
	//
	// true
	IsEnabled *bool `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
}

func (s FeaturesQuota) String() string {
	return tea.Prettify(s)
}

func (s FeaturesQuota) GoString() string {
	return s.String()
}

func (s *FeaturesQuota) SetIsEnabled(v bool) *FeaturesQuota {
	s.IsEnabled = &v
	return s
}

type ForwardInfo struct {
	EipAllocationId *string `json:"EipAllocationId,omitempty" xml:"EipAllocationId,omitempty"`
	NatGatewayId    *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
}

func (s ForwardInfo) String() string {
	return tea.Prettify(s)
}

func (s ForwardInfo) GoString() string {
	return s.String()
}

func (s *ForwardInfo) SetEipAllocationId(v string) *ForwardInfo {
	s.EipAllocationId = &v
	return s
}

func (s *ForwardInfo) SetNatGatewayId(v string) *ForwardInfo {
	s.NatGatewayId = &v
	return s
}

type GPUInfo struct {
	Count *int64  `json:"count,omitempty" xml:"count,omitempty"`
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GPUInfo) String() string {
	return tea.Prettify(s)
}

func (s GPUInfo) GoString() string {
	return s.String()
}

func (s *GPUInfo) SetCount(v int64) *GPUInfo {
	s.Count = &v
	return s
}

func (s *GPUInfo) SetType(v string) *GPUInfo {
	s.Type = &v
	return s
}

type GPUMetric struct {
	Index *int64  `json:"Index,omitempty" xml:"Index,omitempty"`
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 0：异常；1：正常
	Status    *int64   `json:"Status,omitempty" xml:"Status,omitempty"`
	UsageRate *float32 `json:"UsageRate,omitempty" xml:"UsageRate,omitempty"`
}

func (s GPUMetric) String() string {
	return tea.Prettify(s)
}

func (s GPUMetric) GoString() string {
	return s.String()
}

func (s *GPUMetric) SetIndex(v int64) *GPUMetric {
	s.Index = &v
	return s
}

func (s *GPUMetric) SetModel(v string) *GPUMetric {
	s.Model = &v
	return s
}

func (s *GPUMetric) SetStatus(v int64) *GPUMetric {
	s.Status = &v
	return s
}

func (s *GPUMetric) SetUsageRate(v float32) *GPUMetric {
	s.UsageRate = &v
	return s
}

type HyperParameterDefinition struct {
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	Name     *string              `json:"Name,omitempty" xml:"Name,omitempty"`
	Range    *HyperParameterRange `json:"Range,omitempty" xml:"Range,omitempty"`
	Required *bool                `json:"Required,omitempty" xml:"Required,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s HyperParameterDefinition) String() string {
	return tea.Prettify(s)
}

func (s HyperParameterDefinition) GoString() string {
	return s.String()
}

func (s *HyperParameterDefinition) SetDefaultValue(v string) *HyperParameterDefinition {
	s.DefaultValue = &v
	return s
}

func (s *HyperParameterDefinition) SetDescription(v string) *HyperParameterDefinition {
	s.Description = &v
	return s
}

func (s *HyperParameterDefinition) SetDisplayName(v string) *HyperParameterDefinition {
	s.DisplayName = &v
	return s
}

func (s *HyperParameterDefinition) SetName(v string) *HyperParameterDefinition {
	s.Name = &v
	return s
}

func (s *HyperParameterDefinition) SetRange(v *HyperParameterRange) *HyperParameterDefinition {
	s.Range = v
	return s
}

func (s *HyperParameterDefinition) SetRequired(v bool) *HyperParameterDefinition {
	s.Required = &v
	return s
}

func (s *HyperParameterDefinition) SetType(v string) *HyperParameterDefinition {
	s.Type = &v
	return s
}

type HyperParameterRange struct {
	Enum             []*string `json:"Enum,omitempty" xml:"Enum,omitempty" type:"Repeated"`
	ExclusiveMaximum *bool     `json:"ExclusiveMaximum,omitempty" xml:"ExclusiveMaximum,omitempty"`
	ExclusiveMinimum *bool     `json:"ExclusiveMinimum,omitempty" xml:"ExclusiveMinimum,omitempty"`
	MaxLength        *int64    `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	Maximum          *string   `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	MinLength        *int64    `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	Minimum          *string   `json:"Minimum,omitempty" xml:"Minimum,omitempty"`
	Pattern          *string   `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
}

func (s HyperParameterRange) String() string {
	return tea.Prettify(s)
}

func (s HyperParameterRange) GoString() string {
	return s.String()
}

func (s *HyperParameterRange) SetEnum(v []*string) *HyperParameterRange {
	s.Enum = v
	return s
}

func (s *HyperParameterRange) SetExclusiveMaximum(v bool) *HyperParameterRange {
	s.ExclusiveMaximum = &v
	return s
}

func (s *HyperParameterRange) SetExclusiveMinimum(v bool) *HyperParameterRange {
	s.ExclusiveMinimum = &v
	return s
}

func (s *HyperParameterRange) SetMaxLength(v int64) *HyperParameterRange {
	s.MaxLength = &v
	return s
}

func (s *HyperParameterRange) SetMaximum(v string) *HyperParameterRange {
	s.Maximum = &v
	return s
}

func (s *HyperParameterRange) SetMinLength(v int64) *HyperParameterRange {
	s.MinLength = &v
	return s
}

func (s *HyperParameterRange) SetMinimum(v string) *HyperParameterRange {
	s.Minimum = &v
	return s
}

func (s *HyperParameterRange) SetPattern(v string) *HyperParameterRange {
	s.Pattern = &v
	return s
}

type JobSettings struct {
	AdvancedSettings map[string]interface{} `json:"AdvancedSettings,omitempty" xml:"AdvancedSettings,omitempty"`
	// example:
	//
	// 166924
	BusinessUserId *string `json:"BusinessUserId,omitempty" xml:"BusinessUserId,omitempty"`
	// example:
	//
	// SilkFlow
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// example:
	//
	// 535.54.03
	Driver *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
	// example:
	//
	// false
	EnableErrorMonitoringInAIMaster *bool `json:"EnableErrorMonitoringInAIMaster,omitempty" xml:"EnableErrorMonitoringInAIMaster,omitempty"`
	// example:
	//
	// true
	EnableOssAppend *bool `json:"EnableOssAppend,omitempty" xml:"EnableOssAppend,omitempty"`
	// example:
	//
	// true
	EnableRDMA *bool `json:"EnableRDMA,omitempty" xml:"EnableRDMA,omitempty"`
	// example:
	//
	// true
	EnableSanityCheck *bool `json:"EnableSanityCheck,omitempty" xml:"EnableSanityCheck,omitempty"`
	// example:
	//
	// true
	EnableTideResource *bool `json:"EnableTideResource,omitempty" xml:"EnableTideResource,omitempty"`
	// example:
	//
	// --enable-log-hang-detection true
	ErrorMonitoringArgs *string `json:"ErrorMonitoringArgs,omitempty" xml:"ErrorMonitoringArgs,omitempty"`
	// example:
	//
	// 30
	JobReservedMinutes *int32 `json:"JobReservedMinutes,omitempty" xml:"JobReservedMinutes,omitempty"`
	// example:
	//
	// Always
	JobReservedPolicy *string `json:"JobReservedPolicy,omitempty" xml:"JobReservedPolicy,omitempty"`
	// example:
	//
	// AcceptQuotaOverSold
	OversoldType *string `json:"OversoldType,omitempty" xml:"OversoldType,omitempty"`
	// example:
	//
	// pid-123456
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// example:
	//
	// --sanity-check-timing=AfterJobFaultTolerant --sanity-check-timeout-ops=MarkJobFai
	SanityCheckArgs *string            `json:"SanityCheckArgs,omitempty" xml:"SanityCheckArgs,omitempty"`
	Tags            map[string]*string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s JobSettings) String() string {
	return tea.Prettify(s)
}

func (s JobSettings) GoString() string {
	return s.String()
}

func (s *JobSettings) SetAdvancedSettings(v map[string]interface{}) *JobSettings {
	s.AdvancedSettings = v
	return s
}

func (s *JobSettings) SetBusinessUserId(v string) *JobSettings {
	s.BusinessUserId = &v
	return s
}

func (s *JobSettings) SetCaller(v string) *JobSettings {
	s.Caller = &v
	return s
}

func (s *JobSettings) SetDriver(v string) *JobSettings {
	s.Driver = &v
	return s
}

func (s *JobSettings) SetEnableErrorMonitoringInAIMaster(v bool) *JobSettings {
	s.EnableErrorMonitoringInAIMaster = &v
	return s
}

func (s *JobSettings) SetEnableOssAppend(v bool) *JobSettings {
	s.EnableOssAppend = &v
	return s
}

func (s *JobSettings) SetEnableRDMA(v bool) *JobSettings {
	s.EnableRDMA = &v
	return s
}

func (s *JobSettings) SetEnableSanityCheck(v bool) *JobSettings {
	s.EnableSanityCheck = &v
	return s
}

func (s *JobSettings) SetEnableTideResource(v bool) *JobSettings {
	s.EnableTideResource = &v
	return s
}

func (s *JobSettings) SetErrorMonitoringArgs(v string) *JobSettings {
	s.ErrorMonitoringArgs = &v
	return s
}

func (s *JobSettings) SetJobReservedMinutes(v int32) *JobSettings {
	s.JobReservedMinutes = &v
	return s
}

func (s *JobSettings) SetJobReservedPolicy(v string) *JobSettings {
	s.JobReservedPolicy = &v
	return s
}

func (s *JobSettings) SetOversoldType(v string) *JobSettings {
	s.OversoldType = &v
	return s
}

func (s *JobSettings) SetPipelineId(v string) *JobSettings {
	s.PipelineId = &v
	return s
}

func (s *JobSettings) SetSanityCheckArgs(v string) *JobSettings {
	s.SanityCheckArgs = &v
	return s
}

func (s *JobSettings) SetTags(v map[string]*string) *JobSettings {
	s.Tags = v
	return s
}

type JobViewMetric struct {
	CPUUsageRate      *string   `json:"CPUUsageRate,omitempty" xml:"CPUUsageRate,omitempty"`
	DiskReadRate      *string   `json:"DiskReadRate,omitempty" xml:"DiskReadRate,omitempty"`
	DiskWriteRate     *string   `json:"DiskWriteRate,omitempty" xml:"DiskWriteRate,omitempty"`
	GPUUsageRate      *string   `json:"GPUUsageRate,omitempty" xml:"GPUUsageRate,omitempty"`
	JobId             *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobType           *string   `json:"JobType,omitempty" xml:"JobType,omitempty"`
	MemoryUsageRate   *string   `json:"MemoryUsageRate,omitempty" xml:"MemoryUsageRate,omitempty"`
	NetworkInputRate  *string   `json:"NetworkInputRate,omitempty" xml:"NetworkInputRate,omitempty"`
	NetworkOutputRate *string   `json:"NetworkOutputRate,omitempty" xml:"NetworkOutputRate,omitempty"`
	NodeNames         []*string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty" type:"Repeated"`
	RequestCPU        *int32    `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	RequestGPU        *int32    `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	RequestMemory     *int64    `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	// example:
	//
	// rg17tmvwiokhzaxg
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	TotalCPU        *int32  `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	TotalGPU        *int32  `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	TotalMemory     *int64  `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s JobViewMetric) String() string {
	return tea.Prettify(s)
}

func (s JobViewMetric) GoString() string {
	return s.String()
}

func (s *JobViewMetric) SetCPUUsageRate(v string) *JobViewMetric {
	s.CPUUsageRate = &v
	return s
}

func (s *JobViewMetric) SetDiskReadRate(v string) *JobViewMetric {
	s.DiskReadRate = &v
	return s
}

func (s *JobViewMetric) SetDiskWriteRate(v string) *JobViewMetric {
	s.DiskWriteRate = &v
	return s
}

func (s *JobViewMetric) SetGPUUsageRate(v string) *JobViewMetric {
	s.GPUUsageRate = &v
	return s
}

func (s *JobViewMetric) SetJobId(v string) *JobViewMetric {
	s.JobId = &v
	return s
}

func (s *JobViewMetric) SetJobType(v string) *JobViewMetric {
	s.JobType = &v
	return s
}

func (s *JobViewMetric) SetMemoryUsageRate(v string) *JobViewMetric {
	s.MemoryUsageRate = &v
	return s
}

func (s *JobViewMetric) SetNetworkInputRate(v string) *JobViewMetric {
	s.NetworkInputRate = &v
	return s
}

func (s *JobViewMetric) SetNetworkOutputRate(v string) *JobViewMetric {
	s.NetworkOutputRate = &v
	return s
}

func (s *JobViewMetric) SetNodeNames(v []*string) *JobViewMetric {
	s.NodeNames = v
	return s
}

func (s *JobViewMetric) SetRequestCPU(v int32) *JobViewMetric {
	s.RequestCPU = &v
	return s
}

func (s *JobViewMetric) SetRequestGPU(v int32) *JobViewMetric {
	s.RequestGPU = &v
	return s
}

func (s *JobViewMetric) SetRequestMemory(v int64) *JobViewMetric {
	s.RequestMemory = &v
	return s
}

func (s *JobViewMetric) SetResourceGroupID(v string) *JobViewMetric {
	s.ResourceGroupID = &v
	return s
}

func (s *JobViewMetric) SetTotalCPU(v int32) *JobViewMetric {
	s.TotalCPU = &v
	return s
}

func (s *JobViewMetric) SetTotalGPU(v int32) *JobViewMetric {
	s.TotalGPU = &v
	return s
}

func (s *JobViewMetric) SetTotalMemory(v int64) *JobViewMetric {
	s.TotalMemory = &v
	return s
}

func (s *JobViewMetric) SetUserId(v string) *JobViewMetric {
	s.UserId = &v
	return s
}

type Label struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s Label) String() string {
	return tea.Prettify(s)
}

func (s Label) GoString() string {
	return s.String()
}

func (s *Label) SetKey(v string) *Label {
	s.Key = &v
	return s
}

func (s *Label) SetValue(v string) *Label {
	s.Value = &v
	return s
}

type Location struct {
	LocationType  *string                `json:"LocationType,omitempty" xml:"LocationType,omitempty"`
	LocationValue map[string]interface{} `json:"LocationValue,omitempty" xml:"LocationValue,omitempty"`
}

func (s Location) String() string {
	return tea.Prettify(s)
}

func (s Location) GoString() string {
	return s.String()
}

func (s *Location) SetLocationType(v string) *Location {
	s.LocationType = &v
	return s
}

func (s *Location) SetLocationValue(v map[string]interface{}) *Location {
	s.LocationValue = v
	return s
}

type MachineGroup struct {
	CreatorID *string `json:"CreatorID,omitempty" xml:"CreatorID,omitempty"`
	// example:
	//
	// 470.199.02
	DefaultDriver   *string `json:"DefaultDriver,omitempty" xml:"DefaultDriver,omitempty"`
	EcsCount        *int64  `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	EcsSpec         *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	GmtCreatedTime  *string `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	GmtExpiredTime  *string `json:"GmtExpiredTime,omitempty" xml:"GmtExpiredTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	GmtStartedTime  *string `json:"GmtStartedTime,omitempty" xml:"GmtStartedTime,omitempty"`
	// example:
	//
	// mg1
	MachineGroupID      *string   `json:"MachineGroupID,omitempty" xml:"MachineGroupID,omitempty"`
	OrderInstanceId     *string   `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	PaymentDuration     *string   `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	PaymentDurationUnit *string   `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	PaymentType         *string   `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	ReasonCode          *string   `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage       *string   `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	ResourceGroupID     *string   `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	Status              *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportedDrivers    []*string `json:"SupportedDrivers,omitempty" xml:"SupportedDrivers,omitempty" type:"Repeated"`
}

func (s MachineGroup) String() string {
	return tea.Prettify(s)
}

func (s MachineGroup) GoString() string {
	return s.String()
}

func (s *MachineGroup) SetCreatorID(v string) *MachineGroup {
	s.CreatorID = &v
	return s
}

func (s *MachineGroup) SetDefaultDriver(v string) *MachineGroup {
	s.DefaultDriver = &v
	return s
}

func (s *MachineGroup) SetEcsCount(v int64) *MachineGroup {
	s.EcsCount = &v
	return s
}

func (s *MachineGroup) SetEcsSpec(v string) *MachineGroup {
	s.EcsSpec = &v
	return s
}

func (s *MachineGroup) SetGmtCreatedTime(v string) *MachineGroup {
	s.GmtCreatedTime = &v
	return s
}

func (s *MachineGroup) SetGmtExpiredTime(v string) *MachineGroup {
	s.GmtExpiredTime = &v
	return s
}

func (s *MachineGroup) SetGmtModifiedTime(v string) *MachineGroup {
	s.GmtModifiedTime = &v
	return s
}

func (s *MachineGroup) SetGmtStartedTime(v string) *MachineGroup {
	s.GmtStartedTime = &v
	return s
}

func (s *MachineGroup) SetMachineGroupID(v string) *MachineGroup {
	s.MachineGroupID = &v
	return s
}

func (s *MachineGroup) SetOrderInstanceId(v string) *MachineGroup {
	s.OrderInstanceId = &v
	return s
}

func (s *MachineGroup) SetPaymentDuration(v string) *MachineGroup {
	s.PaymentDuration = &v
	return s
}

func (s *MachineGroup) SetPaymentDurationUnit(v string) *MachineGroup {
	s.PaymentDurationUnit = &v
	return s
}

func (s *MachineGroup) SetPaymentType(v string) *MachineGroup {
	s.PaymentType = &v
	return s
}

func (s *MachineGroup) SetReasonCode(v string) *MachineGroup {
	s.ReasonCode = &v
	return s
}

func (s *MachineGroup) SetReasonMessage(v string) *MachineGroup {
	s.ReasonMessage = &v
	return s
}

func (s *MachineGroup) SetResourceGroupID(v string) *MachineGroup {
	s.ResourceGroupID = &v
	return s
}

func (s *MachineGroup) SetStatus(v string) *MachineGroup {
	s.Status = &v
	return s
}

func (s *MachineGroup) SetSupportedDrivers(v []*string) *MachineGroup {
	s.SupportedDrivers = v
	return s
}

type Metric struct {
	// example:
	//
	// rg17tmvwiokhzaxg
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// example:
	//
	// 23000
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s Metric) String() string {
	return tea.Prettify(s)
}

func (s Metric) GoString() string {
	return s.String()
}

func (s *Metric) SetTime(v int64) *Metric {
	s.Time = &v
	return s
}

func (s *Metric) SetValue(v string) *Metric {
	s.Value = &v
	return s
}

type MetricDefinition struct {
	// example:
	//
	// train dataset oob score
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// train:oob_score
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// .*train:oob_score=([-+]?[0-9]*\\.?[0-9]+(?:[eE][-+]?[0-9]+)?).*
	Regex *string `json:"Regex,omitempty" xml:"Regex,omitempty"`
}

func (s MetricDefinition) String() string {
	return tea.Prettify(s)
}

func (s MetricDefinition) GoString() string {
	return s.String()
}

func (s *MetricDefinition) SetDescription(v string) *MetricDefinition {
	s.Description = &v
	return s
}

func (s *MetricDefinition) SetName(v string) *MetricDefinition {
	s.Name = &v
	return s
}

func (s *MetricDefinition) SetRegex(v string) *MetricDefinition {
	s.Regex = &v
	return s
}

type Node struct {
	AcceleratorType   *string        `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	BoundQuotas       []*QuotaIdName `json:"BoundQuotas,omitempty" xml:"BoundQuotas,omitempty" type:"Repeated"`
	CPU               *string        `json:"CPU,omitempty" xml:"CPU,omitempty"`
	CreatorId         *string        `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	GPU               *string        `json:"GPU,omitempty" xml:"GPU,omitempty"`
	GPUMemory         *string        `json:"GPUMemory,omitempty" xml:"GPUMemory,omitempty"`
	GPUType           *string        `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	GmtCreateTime     *string        `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtExpiredTime    *string        `json:"GmtExpiredTime,omitempty" xml:"GmtExpiredTime,omitempty"`
	GmtModifiedTime   *string        `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	IsBound           *bool          `json:"IsBound,omitempty" xml:"IsBound,omitempty"`
	LimitCPU          *string        `json:"LimitCPU,omitempty" xml:"LimitCPU,omitempty"`
	LimitGPU          *string        `json:"LimitGPU,omitempty" xml:"LimitGPU,omitempty"`
	LimitMemory       *string        `json:"LimitMemory,omitempty" xml:"LimitMemory,omitempty"`
	MachineGroupId    *string        `json:"MachineGroupId,omitempty" xml:"MachineGroupId,omitempty"`
	Memory            *string        `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NodeName          *string        `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeStatus        *string        `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	NodeType          *string        `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OrderStatus       *string        `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	PodNum            *int64         `json:"PodNum,omitempty" xml:"PodNum,omitempty"`
	ReasonCode        *string        `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage     *string        `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	RequestCPU        *string        `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	RequestGPU        *string        `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	RequestMemory     *string        `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	ResourceGroupId   *string        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceGroupName *string        `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	Users             []*UserInfo    `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
	WorkloadNum       *int64         `json:"WorkloadNum,omitempty" xml:"WorkloadNum,omitempty"`
}

func (s Node) String() string {
	return tea.Prettify(s)
}

func (s Node) GoString() string {
	return s.String()
}

func (s *Node) SetAcceleratorType(v string) *Node {
	s.AcceleratorType = &v
	return s
}

func (s *Node) SetBoundQuotas(v []*QuotaIdName) *Node {
	s.BoundQuotas = v
	return s
}

func (s *Node) SetCPU(v string) *Node {
	s.CPU = &v
	return s
}

func (s *Node) SetCreatorId(v string) *Node {
	s.CreatorId = &v
	return s
}

func (s *Node) SetGPU(v string) *Node {
	s.GPU = &v
	return s
}

func (s *Node) SetGPUMemory(v string) *Node {
	s.GPUMemory = &v
	return s
}

func (s *Node) SetGPUType(v string) *Node {
	s.GPUType = &v
	return s
}

func (s *Node) SetGmtCreateTime(v string) *Node {
	s.GmtCreateTime = &v
	return s
}

func (s *Node) SetGmtExpiredTime(v string) *Node {
	s.GmtExpiredTime = &v
	return s
}

func (s *Node) SetGmtModifiedTime(v string) *Node {
	s.GmtModifiedTime = &v
	return s
}

func (s *Node) SetIsBound(v bool) *Node {
	s.IsBound = &v
	return s
}

func (s *Node) SetLimitCPU(v string) *Node {
	s.LimitCPU = &v
	return s
}

func (s *Node) SetLimitGPU(v string) *Node {
	s.LimitGPU = &v
	return s
}

func (s *Node) SetLimitMemory(v string) *Node {
	s.LimitMemory = &v
	return s
}

func (s *Node) SetMachineGroupId(v string) *Node {
	s.MachineGroupId = &v
	return s
}

func (s *Node) SetMemory(v string) *Node {
	s.Memory = &v
	return s
}

func (s *Node) SetNodeName(v string) *Node {
	s.NodeName = &v
	return s
}

func (s *Node) SetNodeStatus(v string) *Node {
	s.NodeStatus = &v
	return s
}

func (s *Node) SetNodeType(v string) *Node {
	s.NodeType = &v
	return s
}

func (s *Node) SetOrderStatus(v string) *Node {
	s.OrderStatus = &v
	return s
}

func (s *Node) SetPodNum(v int64) *Node {
	s.PodNum = &v
	return s
}

func (s *Node) SetReasonCode(v string) *Node {
	s.ReasonCode = &v
	return s
}

func (s *Node) SetReasonMessage(v string) *Node {
	s.ReasonMessage = &v
	return s
}

func (s *Node) SetRequestCPU(v string) *Node {
	s.RequestCPU = &v
	return s
}

func (s *Node) SetRequestGPU(v string) *Node {
	s.RequestGPU = &v
	return s
}

func (s *Node) SetRequestMemory(v string) *Node {
	s.RequestMemory = &v
	return s
}

func (s *Node) SetResourceGroupId(v string) *Node {
	s.ResourceGroupId = &v
	return s
}

func (s *Node) SetResourceGroupName(v string) *Node {
	s.ResourceGroupName = &v
	return s
}

func (s *Node) SetUsers(v []*UserInfo) *Node {
	s.Users = v
	return s
}

func (s *Node) SetWorkloadNum(v int64) *Node {
	s.WorkloadNum = &v
	return s
}

type NodeGPUMetric struct {
	AcceleratorType *string      `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	GPUCount        *int32       `json:"GPUCount,omitempty" xml:"GPUCount,omitempty"`
	GPUMetrics      []*GPUMetric `json:"GPUMetrics,omitempty" xml:"GPUMetrics,omitempty" type:"Repeated"`
	GPUType         *string      `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	MemoryUtil      *float32     `json:"MemoryUtil,omitempty" xml:"MemoryUtil,omitempty"`
	NodeId          *string      `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeType        *string      `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	TotalMemory     *float32     `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	UsedMemory      *float32     `json:"UsedMemory,omitempty" xml:"UsedMemory,omitempty"`
}

func (s NodeGPUMetric) String() string {
	return tea.Prettify(s)
}

func (s NodeGPUMetric) GoString() string {
	return s.String()
}

func (s *NodeGPUMetric) SetAcceleratorType(v string) *NodeGPUMetric {
	s.AcceleratorType = &v
	return s
}

func (s *NodeGPUMetric) SetGPUCount(v int32) *NodeGPUMetric {
	s.GPUCount = &v
	return s
}

func (s *NodeGPUMetric) SetGPUMetrics(v []*GPUMetric) *NodeGPUMetric {
	s.GPUMetrics = v
	return s
}

func (s *NodeGPUMetric) SetGPUType(v string) *NodeGPUMetric {
	s.GPUType = &v
	return s
}

func (s *NodeGPUMetric) SetMemoryUtil(v float32) *NodeGPUMetric {
	s.MemoryUtil = &v
	return s
}

func (s *NodeGPUMetric) SetNodeId(v string) *NodeGPUMetric {
	s.NodeId = &v
	return s
}

func (s *NodeGPUMetric) SetNodeType(v string) *NodeGPUMetric {
	s.NodeType = &v
	return s
}

func (s *NodeGPUMetric) SetTotalMemory(v float32) *NodeGPUMetric {
	s.TotalMemory = &v
	return s
}

func (s *NodeGPUMetric) SetUsedMemory(v float32) *NodeGPUMetric {
	s.UsedMemory = &v
	return s
}

type NodeMetric struct {
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// example:
	//
	// 23000
	Metrics []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// example:
	//
	// -i121212node
	NodeID *string `json:"NodeID,omitempty" xml:"NodeID,omitempty"`
}

func (s NodeMetric) String() string {
	return tea.Prettify(s)
}

func (s NodeMetric) GoString() string {
	return s.String()
}

func (s *NodeMetric) SetGPUType(v string) *NodeMetric {
	s.GPUType = &v
	return s
}

func (s *NodeMetric) SetMetrics(v []*Metric) *NodeMetric {
	s.Metrics = v
	return s
}

func (s *NodeMetric) SetNodeID(v string) *NodeMetric {
	s.NodeID = &v
	return s
}

type NodePodInfo struct {
	// example:
	//
	// Running
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// 192.168.2.2
	PodIP *string `json:"PodIP,omitempty" xml:"PodIP,omitempty"`
	// example:
	//
	// test
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// example:
	//
	// test
	PodNamespace *string         `json:"PodNamespace,omitempty" xml:"PodNamespace,omitempty"`
	ResourceSpec *ResourceAmount `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty"`
	// example:
	//
	// dlc19de9s6vn3acr
	WorkloadId *string `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// example:
	//
	// dlc
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
}

func (s NodePodInfo) String() string {
	return tea.Prettify(s)
}

func (s NodePodInfo) GoString() string {
	return s.String()
}

func (s *NodePodInfo) SetPhase(v string) *NodePodInfo {
	s.Phase = &v
	return s
}

func (s *NodePodInfo) SetPodIP(v string) *NodePodInfo {
	s.PodIP = &v
	return s
}

func (s *NodePodInfo) SetPodName(v string) *NodePodInfo {
	s.PodName = &v
	return s
}

func (s *NodePodInfo) SetPodNamespace(v string) *NodePodInfo {
	s.PodNamespace = &v
	return s
}

func (s *NodePodInfo) SetResourceSpec(v *ResourceAmount) *NodePodInfo {
	s.ResourceSpec = v
	return s
}

func (s *NodePodInfo) SetWorkloadId(v string) *NodePodInfo {
	s.WorkloadId = &v
	return s
}

func (s *NodePodInfo) SetWorkloadType(v string) *NodePodInfo {
	s.WorkloadType = &v
	return s
}

type NodeSnapshot struct {
	AncestorQuotaWorkloadNum   *int32                   `json:"AncestorQuotaWorkloadNum,omitempty" xml:"AncestorQuotaWorkloadNum,omitempty"`
	DescendantQuotaWorkloadNum *int32                   `json:"DescendantQuotaWorkloadNum,omitempty" xml:"DescendantQuotaWorkloadNum,omitempty"`
	NodeName                   *string                  `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	RequestCPU                 *string                  `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	RequestGPU                 *string                  `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	RequestMemory              *string                  `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	SelfQuotaWorkloadNum       *int32                   `json:"SelfQuotaWorkloadNum,omitempty" xml:"SelfQuotaWorkloadNum,omitempty"`
	WorkloadNum                *int32                   `json:"WorkloadNum,omitempty" xml:"WorkloadNum,omitempty"`
	Workloads                  []*NodeSnapshotWorkloads `json:"Workloads,omitempty" xml:"Workloads,omitempty" type:"Repeated"`
}

func (s NodeSnapshot) String() string {
	return tea.Prettify(s)
}

func (s NodeSnapshot) GoString() string {
	return s.String()
}

func (s *NodeSnapshot) SetAncestorQuotaWorkloadNum(v int32) *NodeSnapshot {
	s.AncestorQuotaWorkloadNum = &v
	return s
}

func (s *NodeSnapshot) SetDescendantQuotaWorkloadNum(v int32) *NodeSnapshot {
	s.DescendantQuotaWorkloadNum = &v
	return s
}

func (s *NodeSnapshot) SetNodeName(v string) *NodeSnapshot {
	s.NodeName = &v
	return s
}

func (s *NodeSnapshot) SetRequestCPU(v string) *NodeSnapshot {
	s.RequestCPU = &v
	return s
}

func (s *NodeSnapshot) SetRequestGPU(v string) *NodeSnapshot {
	s.RequestGPU = &v
	return s
}

func (s *NodeSnapshot) SetRequestMemory(v string) *NodeSnapshot {
	s.RequestMemory = &v
	return s
}

func (s *NodeSnapshot) SetSelfQuotaWorkloadNum(v int32) *NodeSnapshot {
	s.SelfQuotaWorkloadNum = &v
	return s
}

func (s *NodeSnapshot) SetWorkloadNum(v int32) *NodeSnapshot {
	s.WorkloadNum = &v
	return s
}

func (s *NodeSnapshot) SetWorkloads(v []*NodeSnapshotWorkloads) *NodeSnapshot {
	s.Workloads = v
	return s
}

type NodeSnapshotWorkloads struct {
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	WorkloadId   *string `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
}

func (s NodeSnapshotWorkloads) String() string {
	return tea.Prettify(s)
}

func (s NodeSnapshotWorkloads) GoString() string {
	return s.String()
}

func (s *NodeSnapshotWorkloads) SetName(v string) *NodeSnapshotWorkloads {
	s.Name = &v
	return s
}

func (s *NodeSnapshotWorkloads) SetWorkloadId(v string) *NodeSnapshotWorkloads {
	s.WorkloadId = &v
	return s
}

func (s *NodeSnapshotWorkloads) SetWorkloadType(v string) *NodeSnapshotWorkloads {
	s.WorkloadType = &v
	return s
}

type NodeSpec struct {
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// ecs.g6.4xlarge
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s NodeSpec) String() string {
	return tea.Prettify(s)
}

func (s NodeSpec) GoString() string {
	return s.String()
}

func (s *NodeSpec) SetCount(v int64) *NodeSpec {
	s.Count = &v
	return s
}

func (s *NodeSpec) SetType(v string) *NodeSpec {
	s.Type = &v
	return s
}

type NodeType struct {
	// example:
	//
	// CPU
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// example:
	//
	// 16
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// 0
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// example:
	//
	// 80G
	GPUMemory *string `json:"GPUMemory,omitempty" xml:"GPUMemory,omitempty"`
	GPUType   *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// example:
	//
	// 64Gi
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// ecs.g6.4xlarge
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s NodeType) String() string {
	return tea.Prettify(s)
}

func (s NodeType) GoString() string {
	return s.String()
}

func (s *NodeType) SetAcceleratorType(v string) *NodeType {
	s.AcceleratorType = &v
	return s
}

func (s *NodeType) SetCPU(v string) *NodeType {
	s.CPU = &v
	return s
}

func (s *NodeType) SetGPU(v string) *NodeType {
	s.GPU = &v
	return s
}

func (s *NodeType) SetGPUMemory(v string) *NodeType {
	s.GPUMemory = &v
	return s
}

func (s *NodeType) SetGPUType(v string) *NodeType {
	s.GPUType = &v
	return s
}

func (s *NodeType) SetMemory(v string) *NodeType {
	s.Memory = &v
	return s
}

func (s *NodeType) SetNodeType(v string) *NodeType {
	s.NodeType = &v
	return s
}

type NodeTypeStatistic struct {
	// example:
	//
	// 4
	CanBeBoundCount *int32 `json:"CanBeBoundCount,omitempty" xml:"CanBeBoundCount,omitempty"`
	// example:
	//
	// ecs.g6.4xlarge
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s NodeTypeStatistic) String() string {
	return tea.Prettify(s)
}

func (s NodeTypeStatistic) GoString() string {
	return s.String()
}

func (s *NodeTypeStatistic) SetCanBeBoundCount(v int32) *NodeTypeStatistic {
	s.CanBeBoundCount = &v
	return s
}

func (s *NodeTypeStatistic) SetNodeType(v string) *NodeTypeStatistic {
	s.NodeType = &v
	return s
}

func (s *NodeTypeStatistic) SetTotalCount(v int32) *NodeTypeStatistic {
	s.TotalCount = &v
	return s
}

type NodeViewMetric struct {
	CPUUsageRate      *string `json:"CPUUsageRate,omitempty" xml:"CPUUsageRate,omitempty"`
	CreatedTime       *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DiskReadRate      *string `json:"DiskReadRate,omitempty" xml:"DiskReadRate,omitempty"`
	DiskWriteRate     *string `json:"DiskWriteRate,omitempty" xml:"DiskWriteRate,omitempty"`
	GPUType           *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	MachineGroupID    *string `json:"MachineGroupID,omitempty" xml:"MachineGroupID,omitempty"`
	MemoryUsageRate   *string `json:"MemoryUsageRate,omitempty" xml:"MemoryUsageRate,omitempty"`
	NetworkInputRate  *string `json:"NetworkInputRate,omitempty" xml:"NetworkInputRate,omitempty"`
	NetworkOutputRate *string `json:"NetworkOutputRate,omitempty" xml:"NetworkOutputRate,omitempty"`
	// example:
	//
	// -i121212node
	NodeID        *string                `json:"NodeID,omitempty" xml:"NodeID,omitempty"`
	NodeStatus    *string                `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	NodeType      *string                `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	RequestCPU    *int64                 `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	RequestGPU    *int64                 `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	RequestMemory *int64                 `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	TaskIdMap     map[string]interface{} `json:"TaskIdMap,omitempty" xml:"TaskIdMap,omitempty"`
	TotalCPU      *int64                 `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	TotalGPU      *int64                 `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	TotalMemory   *int64                 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	TotalTasks    *int64                 `json:"TotalTasks,omitempty" xml:"TotalTasks,omitempty"`
	UserIDs       []*string              `json:"UserIDs,omitempty" xml:"UserIDs,omitempty" type:"Repeated"`
	UserNumber    *string                `json:"UserNumber,omitempty" xml:"UserNumber,omitempty"`
}

func (s NodeViewMetric) String() string {
	return tea.Prettify(s)
}

func (s NodeViewMetric) GoString() string {
	return s.String()
}

func (s *NodeViewMetric) SetCPUUsageRate(v string) *NodeViewMetric {
	s.CPUUsageRate = &v
	return s
}

func (s *NodeViewMetric) SetCreatedTime(v string) *NodeViewMetric {
	s.CreatedTime = &v
	return s
}

func (s *NodeViewMetric) SetDiskReadRate(v string) *NodeViewMetric {
	s.DiskReadRate = &v
	return s
}

func (s *NodeViewMetric) SetDiskWriteRate(v string) *NodeViewMetric {
	s.DiskWriteRate = &v
	return s
}

func (s *NodeViewMetric) SetGPUType(v string) *NodeViewMetric {
	s.GPUType = &v
	return s
}

func (s *NodeViewMetric) SetMachineGroupID(v string) *NodeViewMetric {
	s.MachineGroupID = &v
	return s
}

func (s *NodeViewMetric) SetMemoryUsageRate(v string) *NodeViewMetric {
	s.MemoryUsageRate = &v
	return s
}

func (s *NodeViewMetric) SetNetworkInputRate(v string) *NodeViewMetric {
	s.NetworkInputRate = &v
	return s
}

func (s *NodeViewMetric) SetNetworkOutputRate(v string) *NodeViewMetric {
	s.NetworkOutputRate = &v
	return s
}

func (s *NodeViewMetric) SetNodeID(v string) *NodeViewMetric {
	s.NodeID = &v
	return s
}

func (s *NodeViewMetric) SetNodeStatus(v string) *NodeViewMetric {
	s.NodeStatus = &v
	return s
}

func (s *NodeViewMetric) SetNodeType(v string) *NodeViewMetric {
	s.NodeType = &v
	return s
}

func (s *NodeViewMetric) SetRequestCPU(v int64) *NodeViewMetric {
	s.RequestCPU = &v
	return s
}

func (s *NodeViewMetric) SetRequestGPU(v int64) *NodeViewMetric {
	s.RequestGPU = &v
	return s
}

func (s *NodeViewMetric) SetRequestMemory(v int64) *NodeViewMetric {
	s.RequestMemory = &v
	return s
}

func (s *NodeViewMetric) SetTaskIdMap(v map[string]interface{}) *NodeViewMetric {
	s.TaskIdMap = v
	return s
}

func (s *NodeViewMetric) SetTotalCPU(v int64) *NodeViewMetric {
	s.TotalCPU = &v
	return s
}

func (s *NodeViewMetric) SetTotalGPU(v int64) *NodeViewMetric {
	s.TotalGPU = &v
	return s
}

func (s *NodeViewMetric) SetTotalMemory(v int64) *NodeViewMetric {
	s.TotalMemory = &v
	return s
}

func (s *NodeViewMetric) SetTotalTasks(v int64) *NodeViewMetric {
	s.TotalTasks = &v
	return s
}

func (s *NodeViewMetric) SetUserIDs(v []*string) *NodeViewMetric {
	s.UserIDs = v
	return s
}

func (s *NodeViewMetric) SetUserNumber(v string) *NodeViewMetric {
	s.UserNumber = &v
	return s
}

type Permission struct {
	IsEnabled    *bool   `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s Permission) String() string {
	return tea.Prettify(s)
}

func (s Permission) GoString() string {
	return s.String()
}

func (s *Permission) SetIsEnabled(v bool) *Permission {
	s.IsEnabled = &v
	return s
}

func (s *Permission) SetResourceType(v string) *Permission {
	s.ResourceType = &v
	return s
}

type QueueInfo struct {
	// example:
	//
	// roleMaximumResource
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// ConfigRule
	CodeType       *string `json:"CodeType,omitempty" xml:"CodeType,omitempty"`
	GmtCreatedTime *string `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	// example:
	//
	// "2023-06-22T00:00:00Z"
	GmtDequeuedTime *string `json:"GmtDequeuedTime,omitempty" xml:"GmtDequeuedTime,omitempty"`
	// example:
	//
	// “2023-06-22T00:00:00Z”
	GmtEnqueuedTime *string `json:"GmtEnqueuedTime,omitempty" xml:"GmtEnqueuedTime,omitempty"`
	// example:
	//
	// "2023-06-22T00:00:00Z"
	GmtPositionModifiedTime *string `json:"GmtPositionModifiedTime,omitempty" xml:"GmtPositionModifiedTime,omitempty"`
	// example:
	//
	// test-label-79f5498dd-9qrzs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 10
	Position *int64 `json:"Position,omitempty" xml:"Position,omitempty"`
	// example:
	//
	// 2
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// PaiStrategyIntelligent
	QueueStrategy *string `json:"QueueStrategy,omitempty" xml:"QueueStrategy,omitempty"`
	// example:
	//
	// “quotamtl37ge7gkvdz”
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// example:
	//
	// Current GPU Limit is 5, limited by Role PAI.AlgoDeveloper
	Reason   *string         `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Resource *ResourceAmount `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// example:
	//
	// Enqueued
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// example:
	//
	// dlcxxxx
	WorkloadId   *string `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
	// example:
	//
	// PreAllocation
	WorkloadStatus *string `json:"WorkloadStatus,omitempty" xml:"WorkloadStatus,omitempty"`
	// example:
	//
	// dlc
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
	// example:
	//
	// “432524”
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueueInfo) String() string {
	return tea.Prettify(s)
}

func (s QueueInfo) GoString() string {
	return s.String()
}

func (s *QueueInfo) SetCode(v string) *QueueInfo {
	s.Code = &v
	return s
}

func (s *QueueInfo) SetCodeType(v string) *QueueInfo {
	s.CodeType = &v
	return s
}

func (s *QueueInfo) SetGmtCreatedTime(v string) *QueueInfo {
	s.GmtCreatedTime = &v
	return s
}

func (s *QueueInfo) SetGmtDequeuedTime(v string) *QueueInfo {
	s.GmtDequeuedTime = &v
	return s
}

func (s *QueueInfo) SetGmtEnqueuedTime(v string) *QueueInfo {
	s.GmtEnqueuedTime = &v
	return s
}

func (s *QueueInfo) SetGmtPositionModifiedTime(v string) *QueueInfo {
	s.GmtPositionModifiedTime = &v
	return s
}

func (s *QueueInfo) SetName(v string) *QueueInfo {
	s.Name = &v
	return s
}

func (s *QueueInfo) SetPosition(v int64) *QueueInfo {
	s.Position = &v
	return s
}

func (s *QueueInfo) SetPriority(v int64) *QueueInfo {
	s.Priority = &v
	return s
}

func (s *QueueInfo) SetQueueStrategy(v string) *QueueInfo {
	s.QueueStrategy = &v
	return s
}

func (s *QueueInfo) SetQuotaId(v string) *QueueInfo {
	s.QuotaId = &v
	return s
}

func (s *QueueInfo) SetReason(v string) *QueueInfo {
	s.Reason = &v
	return s
}

func (s *QueueInfo) SetResource(v *ResourceAmount) *QueueInfo {
	s.Resource = v
	return s
}

func (s *QueueInfo) SetStatus(v string) *QueueInfo {
	s.Status = &v
	return s
}

func (s *QueueInfo) SetUserId(v string) *QueueInfo {
	s.UserId = &v
	return s
}

func (s *QueueInfo) SetUserName(v string) *QueueInfo {
	s.UserName = &v
	return s
}

func (s *QueueInfo) SetWorkloadId(v string) *QueueInfo {
	s.WorkloadId = &v
	return s
}

func (s *QueueInfo) SetWorkloadName(v string) *QueueInfo {
	s.WorkloadName = &v
	return s
}

func (s *QueueInfo) SetWorkloadStatus(v string) *QueueInfo {
	s.WorkloadStatus = &v
	return s
}

func (s *QueueInfo) SetWorkloadType(v string) *QueueInfo {
	s.WorkloadType = &v
	return s
}

func (s *QueueInfo) SetWorkspaceId(v string) *QueueInfo {
	s.WorkspaceId = &v
	return s
}

type Quota struct {
	AllocateStrategy  *string       `json:"AllocateStrategy,omitempty" xml:"AllocateStrategy,omitempty"`
	CreatorId         *string       `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Description       *string       `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreatedTime    *string       `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	GmtModifiedTime   *string       `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels            []*Label      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LatestOperationId *string       `json:"LatestOperationId,omitempty" xml:"LatestOperationId,omitempty"`
	Min               *ResourceSpec `json:"Min,omitempty" xml:"Min,omitempty"`
	ParentQuotaId     *string       `json:"ParentQuotaId,omitempty" xml:"ParentQuotaId,omitempty"`
	QueueStrategy     *string       `json:"QueueStrategy,omitempty" xml:"QueueStrategy,omitempty"`
	QuotaConfig       *QuotaConfig  `json:"QuotaConfig,omitempty" xml:"QuotaConfig,omitempty"`
	QuotaDetails      *QuotaDetails `json:"QuotaDetails,omitempty" xml:"QuotaDetails,omitempty"`
	// example:
	//
	// quota12345
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// example:
	//
	// dlc-quota
	QuotaName        *string            `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	ReasonCode       *string            `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage    *string            `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	ResourceGroupIds []*string          `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
	ResourceType     *string            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status           *string            `json:"Status,omitempty" xml:"Status,omitempty"`
	SubQuotas        []*QuotaIdName     `json:"SubQuotas,omitempty" xml:"SubQuotas,omitempty" type:"Repeated"`
	Workspaces       []*WorkspaceIdName `json:"Workspaces,omitempty" xml:"Workspaces,omitempty" type:"Repeated"`
}

func (s Quota) String() string {
	return tea.Prettify(s)
}

func (s Quota) GoString() string {
	return s.String()
}

func (s *Quota) SetAllocateStrategy(v string) *Quota {
	s.AllocateStrategy = &v
	return s
}

func (s *Quota) SetCreatorId(v string) *Quota {
	s.CreatorId = &v
	return s
}

func (s *Quota) SetDescription(v string) *Quota {
	s.Description = &v
	return s
}

func (s *Quota) SetGmtCreatedTime(v string) *Quota {
	s.GmtCreatedTime = &v
	return s
}

func (s *Quota) SetGmtModifiedTime(v string) *Quota {
	s.GmtModifiedTime = &v
	return s
}

func (s *Quota) SetLabels(v []*Label) *Quota {
	s.Labels = v
	return s
}

func (s *Quota) SetLatestOperationId(v string) *Quota {
	s.LatestOperationId = &v
	return s
}

func (s *Quota) SetMin(v *ResourceSpec) *Quota {
	s.Min = v
	return s
}

func (s *Quota) SetParentQuotaId(v string) *Quota {
	s.ParentQuotaId = &v
	return s
}

func (s *Quota) SetQueueStrategy(v string) *Quota {
	s.QueueStrategy = &v
	return s
}

func (s *Quota) SetQuotaConfig(v *QuotaConfig) *Quota {
	s.QuotaConfig = v
	return s
}

func (s *Quota) SetQuotaDetails(v *QuotaDetails) *Quota {
	s.QuotaDetails = v
	return s
}

func (s *Quota) SetQuotaId(v string) *Quota {
	s.QuotaId = &v
	return s
}

func (s *Quota) SetQuotaName(v string) *Quota {
	s.QuotaName = &v
	return s
}

func (s *Quota) SetReasonCode(v string) *Quota {
	s.ReasonCode = &v
	return s
}

func (s *Quota) SetReasonMessage(v string) *Quota {
	s.ReasonMessage = &v
	return s
}

func (s *Quota) SetResourceGroupIds(v []*string) *Quota {
	s.ResourceGroupIds = v
	return s
}

func (s *Quota) SetResourceType(v string) *Quota {
	s.ResourceType = &v
	return s
}

func (s *Quota) SetStatus(v string) *Quota {
	s.Status = &v
	return s
}

func (s *Quota) SetSubQuotas(v []*QuotaIdName) *Quota {
	s.SubQuotas = v
	return s
}

func (s *Quota) SetWorkspaces(v []*WorkspaceIdName) *Quota {
	s.Workspaces = v
	return s
}

type QuotaConfig struct {
	ACS *ACS `json:"ACS,omitempty" xml:"ACS,omitempty"`
	// example:
	//
	// ceeb37xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 470.199.02
	DefaultGPUDriver               *string           `json:"DefaultGPUDriver,omitempty" xml:"DefaultGPUDriver,omitempty"`
	EnablePreemptSubquotaWorkloads *bool             `json:"EnablePreemptSubquotaWorkloads,omitempty" xml:"EnablePreemptSubquotaWorkloads,omitempty"`
	ResourceSpecs                  []*WorkspaceSpecs `json:"ResourceSpecs,omitempty" xml:"ResourceSpecs,omitempty" type:"Repeated"`
	SupportGPUDrivers              []*string         `json:"SupportGPUDrivers,omitempty" xml:"SupportGPUDrivers,omitempty" type:"Repeated"`
	// example:
	//
	// false
	SupportRDMA *bool    `json:"SupportRDMA,omitempty" xml:"SupportRDMA,omitempty"`
	UserVpc     *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
}

func (s QuotaConfig) String() string {
	return tea.Prettify(s)
}

func (s QuotaConfig) GoString() string {
	return s.String()
}

func (s *QuotaConfig) SetACS(v *ACS) *QuotaConfig {
	s.ACS = v
	return s
}

func (s *QuotaConfig) SetClusterId(v string) *QuotaConfig {
	s.ClusterId = &v
	return s
}

func (s *QuotaConfig) SetDefaultGPUDriver(v string) *QuotaConfig {
	s.DefaultGPUDriver = &v
	return s
}

func (s *QuotaConfig) SetEnablePreemptSubquotaWorkloads(v bool) *QuotaConfig {
	s.EnablePreemptSubquotaWorkloads = &v
	return s
}

func (s *QuotaConfig) SetResourceSpecs(v []*WorkspaceSpecs) *QuotaConfig {
	s.ResourceSpecs = v
	return s
}

func (s *QuotaConfig) SetSupportGPUDrivers(v []*string) *QuotaConfig {
	s.SupportGPUDrivers = v
	return s
}

func (s *QuotaConfig) SetSupportRDMA(v bool) *QuotaConfig {
	s.SupportRDMA = &v
	return s
}

func (s *QuotaConfig) SetUserVpc(v *UserVpc) *QuotaConfig {
	s.UserVpc = v
	return s
}

type QuotaDetails struct {
	ActualMinQuota            *ResourceAmount `json:"ActualMinQuota,omitempty" xml:"ActualMinQuota,omitempty"`
	AllocatedQuota            *ResourceAmount `json:"AllocatedQuota,omitempty" xml:"AllocatedQuota,omitempty"`
	AncestorsAllocatedQuota   *ResourceAmount `json:"AncestorsAllocatedQuota,omitempty" xml:"AncestorsAllocatedQuota,omitempty"`
	DescendantsAllocatedQuota *ResourceAmount `json:"DescendantsAllocatedQuota,omitempty" xml:"DescendantsAllocatedQuota,omitempty"`
	DesiredMinQuota           *ResourceAmount `json:"DesiredMinQuota,omitempty" xml:"DesiredMinQuota,omitempty"`
	RequestedQuota            *ResourceAmount `json:"RequestedQuota,omitempty" xml:"RequestedQuota,omitempty"`
	SelfAllocatedQuota        *ResourceAmount `json:"SelfAllocatedQuota,omitempty" xml:"SelfAllocatedQuota,omitempty"`
	UsedQuota                 *ResourceAmount `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty"`
}

func (s QuotaDetails) String() string {
	return tea.Prettify(s)
}

func (s QuotaDetails) GoString() string {
	return s.String()
}

func (s *QuotaDetails) SetActualMinQuota(v *ResourceAmount) *QuotaDetails {
	s.ActualMinQuota = v
	return s
}

func (s *QuotaDetails) SetAllocatedQuota(v *ResourceAmount) *QuotaDetails {
	s.AllocatedQuota = v
	return s
}

func (s *QuotaDetails) SetAncestorsAllocatedQuota(v *ResourceAmount) *QuotaDetails {
	s.AncestorsAllocatedQuota = v
	return s
}

func (s *QuotaDetails) SetDescendantsAllocatedQuota(v *ResourceAmount) *QuotaDetails {
	s.DescendantsAllocatedQuota = v
	return s
}

func (s *QuotaDetails) SetDesiredMinQuota(v *ResourceAmount) *QuotaDetails {
	s.DesiredMinQuota = v
	return s
}

func (s *QuotaDetails) SetRequestedQuota(v *ResourceAmount) *QuotaDetails {
	s.RequestedQuota = v
	return s
}

func (s *QuotaDetails) SetSelfAllocatedQuota(v *ResourceAmount) *QuotaDetails {
	s.SelfAllocatedQuota = v
	return s
}

func (s *QuotaDetails) SetUsedQuota(v *ResourceAmount) *QuotaDetails {
	s.UsedQuota = v
	return s
}

type QuotaIdName struct {
	// example:
	//
	// quota12345
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// example:
	//
	// dlc-quota
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
}

func (s QuotaIdName) String() string {
	return tea.Prettify(s)
}

func (s QuotaIdName) GoString() string {
	return s.String()
}

func (s *QuotaIdName) SetQuotaId(v string) *QuotaIdName {
	s.QuotaId = &v
	return s
}

func (s *QuotaIdName) SetQuotaName(v string) *QuotaIdName {
	s.QuotaName = &v
	return s
}

type QuotaJob struct {
	Queuing *int64 `json:"Queuing,omitempty" xml:"Queuing,omitempty"`
	Running *int64 `json:"Running,omitempty" xml:"Running,omitempty"`
	Total   *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QuotaJob) String() string {
	return tea.Prettify(s)
}

func (s QuotaJob) GoString() string {
	return s.String()
}

func (s *QuotaJob) SetQueuing(v int64) *QuotaJob {
	s.Queuing = &v
	return s
}

func (s *QuotaJob) SetRunning(v int64) *QuotaJob {
	s.Running = &v
	return s
}

func (s *QuotaJob) SetTotal(v int64) *QuotaJob {
	s.Total = &v
	return s
}

type QuotaJobViewMetric struct {
	CPUUsageRate      *string   `json:"CPUUsageRate,omitempty" xml:"CPUUsageRate,omitempty"`
	DiskReadRate      *string   `json:"DiskReadRate,omitempty" xml:"DiskReadRate,omitempty"`
	DiskWriteRate     *string   `json:"DiskWriteRate,omitempty" xml:"DiskWriteRate,omitempty"`
	GPUUsageRate      *string   `json:"GPUUsageRate,omitempty" xml:"GPUUsageRate,omitempty"`
	JobId             *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobType           *string   `json:"JobType,omitempty" xml:"JobType,omitempty"`
	MemoryUsageRate   *string   `json:"MemoryUsageRate,omitempty" xml:"MemoryUsageRate,omitempty"`
	NetworkInputRate  *string   `json:"NetworkInputRate,omitempty" xml:"NetworkInputRate,omitempty"`
	NetworkOutputRate *string   `json:"NetworkOutputRate,omitempty" xml:"NetworkOutputRate,omitempty"`
	NodeNames         []*string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty" type:"Repeated"`
	RequestCPU        *int32    `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	RequestGPU        *int32    `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	RequestMemory     *int64    `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	TotalCPU          *int32    `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	TotalGPU          *int32    `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	TotalMemory       *int64    `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	UserId            *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QuotaJobViewMetric) String() string {
	return tea.Prettify(s)
}

func (s QuotaJobViewMetric) GoString() string {
	return s.String()
}

func (s *QuotaJobViewMetric) SetCPUUsageRate(v string) *QuotaJobViewMetric {
	s.CPUUsageRate = &v
	return s
}

func (s *QuotaJobViewMetric) SetDiskReadRate(v string) *QuotaJobViewMetric {
	s.DiskReadRate = &v
	return s
}

func (s *QuotaJobViewMetric) SetDiskWriteRate(v string) *QuotaJobViewMetric {
	s.DiskWriteRate = &v
	return s
}

func (s *QuotaJobViewMetric) SetGPUUsageRate(v string) *QuotaJobViewMetric {
	s.GPUUsageRate = &v
	return s
}

func (s *QuotaJobViewMetric) SetJobId(v string) *QuotaJobViewMetric {
	s.JobId = &v
	return s
}

func (s *QuotaJobViewMetric) SetJobType(v string) *QuotaJobViewMetric {
	s.JobType = &v
	return s
}

func (s *QuotaJobViewMetric) SetMemoryUsageRate(v string) *QuotaJobViewMetric {
	s.MemoryUsageRate = &v
	return s
}

func (s *QuotaJobViewMetric) SetNetworkInputRate(v string) *QuotaJobViewMetric {
	s.NetworkInputRate = &v
	return s
}

func (s *QuotaJobViewMetric) SetNetworkOutputRate(v string) *QuotaJobViewMetric {
	s.NetworkOutputRate = &v
	return s
}

func (s *QuotaJobViewMetric) SetNodeNames(v []*string) *QuotaJobViewMetric {
	s.NodeNames = v
	return s
}

func (s *QuotaJobViewMetric) SetRequestCPU(v int32) *QuotaJobViewMetric {
	s.RequestCPU = &v
	return s
}

func (s *QuotaJobViewMetric) SetRequestGPU(v int32) *QuotaJobViewMetric {
	s.RequestGPU = &v
	return s
}

func (s *QuotaJobViewMetric) SetRequestMemory(v int64) *QuotaJobViewMetric {
	s.RequestMemory = &v
	return s
}

func (s *QuotaJobViewMetric) SetTotalCPU(v int32) *QuotaJobViewMetric {
	s.TotalCPU = &v
	return s
}

func (s *QuotaJobViewMetric) SetTotalGPU(v int32) *QuotaJobViewMetric {
	s.TotalGPU = &v
	return s
}

func (s *QuotaJobViewMetric) SetTotalMemory(v int64) *QuotaJobViewMetric {
	s.TotalMemory = &v
	return s
}

func (s *QuotaJobViewMetric) SetUserId(v string) *QuotaJobViewMetric {
	s.UserId = &v
	return s
}

type QuotaMetric struct {
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// example:
	//
	// 23000
	Metrics []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
}

func (s QuotaMetric) String() string {
	return tea.Prettify(s)
}

func (s QuotaMetric) GoString() string {
	return s.String()
}

func (s *QuotaMetric) SetGPUType(v string) *QuotaMetric {
	s.GPUType = &v
	return s
}

func (s *QuotaMetric) SetMetrics(v []*Metric) *QuotaMetric {
	s.Metrics = v
	return s
}

type QuotaNodeViewMetric struct {
	CPUUsageRate      *string `json:"CPUUsageRate,omitempty" xml:"CPUUsageRate,omitempty"`
	CreatedTime       *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DiskReadRate      *string `json:"DiskReadRate,omitempty" xml:"DiskReadRate,omitempty"`
	DiskWriteRate     *string `json:"DiskWriteRate,omitempty" xml:"DiskWriteRate,omitempty"`
	GPUType           *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	MemoryUsageRate   *string `json:"MemoryUsageRate,omitempty" xml:"MemoryUsageRate,omitempty"`
	NetworkInputRate  *string `json:"NetworkInputRate,omitempty" xml:"NetworkInputRate,omitempty"`
	NetworkOutputRate *string `json:"NetworkOutputRate,omitempty" xml:"NetworkOutputRate,omitempty"`
	// example:
	//
	// -i121212node
	NodeID        *string                `json:"NodeID,omitempty" xml:"NodeID,omitempty"`
	NodeStatus    *string                `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	NodeType      *string                `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	QuotaId       *string                `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	RequestCPU    *int64                 `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	RequestGPU    *int64                 `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	RequestMemory *int64                 `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	TaskIdMap     map[string]interface{} `json:"TaskIdMap,omitempty" xml:"TaskIdMap,omitempty"`
	TotalCPU      *int64                 `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	TotalGPU      *int64                 `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	TotalMemory   *int64                 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	TotalTasks    *int64                 `json:"TotalTasks,omitempty" xml:"TotalTasks,omitempty"`
	UserIDs       []*string              `json:"UserIDs,omitempty" xml:"UserIDs,omitempty" type:"Repeated"`
	UserNumber    *string                `json:"UserNumber,omitempty" xml:"UserNumber,omitempty"`
}

func (s QuotaNodeViewMetric) String() string {
	return tea.Prettify(s)
}

func (s QuotaNodeViewMetric) GoString() string {
	return s.String()
}

func (s *QuotaNodeViewMetric) SetCPUUsageRate(v string) *QuotaNodeViewMetric {
	s.CPUUsageRate = &v
	return s
}

func (s *QuotaNodeViewMetric) SetCreatedTime(v string) *QuotaNodeViewMetric {
	s.CreatedTime = &v
	return s
}

func (s *QuotaNodeViewMetric) SetDiskReadRate(v string) *QuotaNodeViewMetric {
	s.DiskReadRate = &v
	return s
}

func (s *QuotaNodeViewMetric) SetDiskWriteRate(v string) *QuotaNodeViewMetric {
	s.DiskWriteRate = &v
	return s
}

func (s *QuotaNodeViewMetric) SetGPUType(v string) *QuotaNodeViewMetric {
	s.GPUType = &v
	return s
}

func (s *QuotaNodeViewMetric) SetMemoryUsageRate(v string) *QuotaNodeViewMetric {
	s.MemoryUsageRate = &v
	return s
}

func (s *QuotaNodeViewMetric) SetNetworkInputRate(v string) *QuotaNodeViewMetric {
	s.NetworkInputRate = &v
	return s
}

func (s *QuotaNodeViewMetric) SetNetworkOutputRate(v string) *QuotaNodeViewMetric {
	s.NetworkOutputRate = &v
	return s
}

func (s *QuotaNodeViewMetric) SetNodeID(v string) *QuotaNodeViewMetric {
	s.NodeID = &v
	return s
}

func (s *QuotaNodeViewMetric) SetNodeStatus(v string) *QuotaNodeViewMetric {
	s.NodeStatus = &v
	return s
}

func (s *QuotaNodeViewMetric) SetNodeType(v string) *QuotaNodeViewMetric {
	s.NodeType = &v
	return s
}

func (s *QuotaNodeViewMetric) SetQuotaId(v string) *QuotaNodeViewMetric {
	s.QuotaId = &v
	return s
}

func (s *QuotaNodeViewMetric) SetRequestCPU(v int64) *QuotaNodeViewMetric {
	s.RequestCPU = &v
	return s
}

func (s *QuotaNodeViewMetric) SetRequestGPU(v int64) *QuotaNodeViewMetric {
	s.RequestGPU = &v
	return s
}

func (s *QuotaNodeViewMetric) SetRequestMemory(v int64) *QuotaNodeViewMetric {
	s.RequestMemory = &v
	return s
}

func (s *QuotaNodeViewMetric) SetTaskIdMap(v map[string]interface{}) *QuotaNodeViewMetric {
	s.TaskIdMap = v
	return s
}

func (s *QuotaNodeViewMetric) SetTotalCPU(v int64) *QuotaNodeViewMetric {
	s.TotalCPU = &v
	return s
}

func (s *QuotaNodeViewMetric) SetTotalGPU(v int64) *QuotaNodeViewMetric {
	s.TotalGPU = &v
	return s
}

func (s *QuotaNodeViewMetric) SetTotalMemory(v int64) *QuotaNodeViewMetric {
	s.TotalMemory = &v
	return s
}

func (s *QuotaNodeViewMetric) SetTotalTasks(v int64) *QuotaNodeViewMetric {
	s.TotalTasks = &v
	return s
}

func (s *QuotaNodeViewMetric) SetUserIDs(v []*string) *QuotaNodeViewMetric {
	s.UserIDs = v
	return s
}

func (s *QuotaNodeViewMetric) SetUserNumber(v string) *QuotaNodeViewMetric {
	s.UserNumber = &v
	return s
}

type QuotaTopo struct {
	Depth           *string          `json:"Depth,omitempty" xml:"Depth,omitempty"`
	ParentQuotaId   *string          `json:"ParentQuotaId,omitempty" xml:"ParentQuotaId,omitempty"`
	QuotaDetails    *QuotaDetails    `json:"QuotaDetails,omitempty" xml:"QuotaDetails,omitempty"`
	QuotaId         *string          `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	QuotaName       *string          `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	ResourceType    *string          `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	WorkloadDetails *WorkloadDetails `json:"WorkloadDetails,omitempty" xml:"WorkloadDetails,omitempty"`
}

func (s QuotaTopo) String() string {
	return tea.Prettify(s)
}

func (s QuotaTopo) GoString() string {
	return s.String()
}

func (s *QuotaTopo) SetDepth(v string) *QuotaTopo {
	s.Depth = &v
	return s
}

func (s *QuotaTopo) SetParentQuotaId(v string) *QuotaTopo {
	s.ParentQuotaId = &v
	return s
}

func (s *QuotaTopo) SetQuotaDetails(v *QuotaDetails) *QuotaTopo {
	s.QuotaDetails = v
	return s
}

func (s *QuotaTopo) SetQuotaId(v string) *QuotaTopo {
	s.QuotaId = &v
	return s
}

func (s *QuotaTopo) SetQuotaName(v string) *QuotaTopo {
	s.QuotaName = &v
	return s
}

func (s *QuotaTopo) SetResourceType(v string) *QuotaTopo {
	s.ResourceType = &v
	return s
}

func (s *QuotaTopo) SetWorkloadDetails(v *WorkloadDetails) *QuotaTopo {
	s.WorkloadDetails = v
	return s
}

type QuotaUser struct {
	Resources     *QuotaUserResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	UserId        *string             `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Username      *string             `json:"Username,omitempty" xml:"Username,omitempty"`
	WorkloadCount *int32              `json:"WorkloadCount,omitempty" xml:"WorkloadCount,omitempty"`
}

func (s QuotaUser) String() string {
	return tea.Prettify(s)
}

func (s QuotaUser) GoString() string {
	return s.String()
}

func (s *QuotaUser) SetResources(v *QuotaUserResources) *QuotaUser {
	s.Resources = v
	return s
}

func (s *QuotaUser) SetUserId(v string) *QuotaUser {
	s.UserId = &v
	return s
}

func (s *QuotaUser) SetUsername(v string) *QuotaUser {
	s.Username = &v
	return s
}

func (s *QuotaUser) SetWorkloadCount(v int32) *QuotaUser {
	s.WorkloadCount = &v
	return s
}

type QuotaUserResources struct {
	Submitted *ResourceAmount `json:"Submitted,omitempty" xml:"Submitted,omitempty"`
	Used      *ResourceAmount `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s QuotaUserResources) String() string {
	return tea.Prettify(s)
}

func (s QuotaUserResources) GoString() string {
	return s.String()
}

func (s *QuotaUserResources) SetSubmitted(v *ResourceAmount) *QuotaUserResources {
	s.Submitted = v
	return s
}

func (s *QuotaUserResources) SetUsed(v *ResourceAmount) *QuotaUserResources {
	s.Used = v
	return s
}

type QuotaUserViewMetric struct {
	CPUNodeNumber     *int32    `json:"CPUNodeNumber,omitempty" xml:"CPUNodeNumber,omitempty"`
	CPUUsageRate      *string   `json:"CPUUsageRate,omitempty" xml:"CPUUsageRate,omitempty"`
	CpuJobNames       []*string `json:"CpuJobNames,omitempty" xml:"CpuJobNames,omitempty" type:"Repeated"`
	CpuNodeNames      []*string `json:"CpuNodeNames,omitempty" xml:"CpuNodeNames,omitempty" type:"Repeated"`
	DiskReadRate      *string   `json:"DiskReadRate,omitempty" xml:"DiskReadRate,omitempty"`
	DiskWriteRate     *string   `json:"DiskWriteRate,omitempty" xml:"DiskWriteRate,omitempty"`
	GPUNodeNumber     *int32    `json:"GPUNodeNumber,omitempty" xml:"GPUNodeNumber,omitempty"`
	GPUUsageRate      *string   `json:"GPUUsageRate,omitempty" xml:"GPUUsageRate,omitempty"`
	GpuJobNames       []*string `json:"GpuJobNames,omitempty" xml:"GpuJobNames,omitempty" type:"Repeated"`
	GpuNodeNames      []*string `json:"GpuNodeNames,omitempty" xml:"GpuNodeNames,omitempty" type:"Repeated"`
	JobType           *string   `json:"JobType,omitempty" xml:"JobType,omitempty"`
	MemoryUsageRate   *string   `json:"MemoryUsageRate,omitempty" xml:"MemoryUsageRate,omitempty"`
	NetworkInputRate  *string   `json:"NetworkInputRate,omitempty" xml:"NetworkInputRate,omitempty"`
	NetworkOutputRate *string   `json:"NetworkOutputRate,omitempty" xml:"NetworkOutputRate,omitempty"`
	NodeNames         []*string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty" type:"Repeated"`
	RequestCPU        *int32    `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	RequestGPU        *int32    `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	RequestMemory     *int64    `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	TotalCPU          *int32    `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	TotalGPU          *int32    `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	TotalMemory       *int64    `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	UserId            *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QuotaUserViewMetric) String() string {
	return tea.Prettify(s)
}

func (s QuotaUserViewMetric) GoString() string {
	return s.String()
}

func (s *QuotaUserViewMetric) SetCPUNodeNumber(v int32) *QuotaUserViewMetric {
	s.CPUNodeNumber = &v
	return s
}

func (s *QuotaUserViewMetric) SetCPUUsageRate(v string) *QuotaUserViewMetric {
	s.CPUUsageRate = &v
	return s
}

func (s *QuotaUserViewMetric) SetCpuJobNames(v []*string) *QuotaUserViewMetric {
	s.CpuJobNames = v
	return s
}

func (s *QuotaUserViewMetric) SetCpuNodeNames(v []*string) *QuotaUserViewMetric {
	s.CpuNodeNames = v
	return s
}

func (s *QuotaUserViewMetric) SetDiskReadRate(v string) *QuotaUserViewMetric {
	s.DiskReadRate = &v
	return s
}

func (s *QuotaUserViewMetric) SetDiskWriteRate(v string) *QuotaUserViewMetric {
	s.DiskWriteRate = &v
	return s
}

func (s *QuotaUserViewMetric) SetGPUNodeNumber(v int32) *QuotaUserViewMetric {
	s.GPUNodeNumber = &v
	return s
}

func (s *QuotaUserViewMetric) SetGPUUsageRate(v string) *QuotaUserViewMetric {
	s.GPUUsageRate = &v
	return s
}

func (s *QuotaUserViewMetric) SetGpuJobNames(v []*string) *QuotaUserViewMetric {
	s.GpuJobNames = v
	return s
}

func (s *QuotaUserViewMetric) SetGpuNodeNames(v []*string) *QuotaUserViewMetric {
	s.GpuNodeNames = v
	return s
}

func (s *QuotaUserViewMetric) SetJobType(v string) *QuotaUserViewMetric {
	s.JobType = &v
	return s
}

func (s *QuotaUserViewMetric) SetMemoryUsageRate(v string) *QuotaUserViewMetric {
	s.MemoryUsageRate = &v
	return s
}

func (s *QuotaUserViewMetric) SetNetworkInputRate(v string) *QuotaUserViewMetric {
	s.NetworkInputRate = &v
	return s
}

func (s *QuotaUserViewMetric) SetNetworkOutputRate(v string) *QuotaUserViewMetric {
	s.NetworkOutputRate = &v
	return s
}

func (s *QuotaUserViewMetric) SetNodeNames(v []*string) *QuotaUserViewMetric {
	s.NodeNames = v
	return s
}

func (s *QuotaUserViewMetric) SetRequestCPU(v int32) *QuotaUserViewMetric {
	s.RequestCPU = &v
	return s
}

func (s *QuotaUserViewMetric) SetRequestGPU(v int32) *QuotaUserViewMetric {
	s.RequestGPU = &v
	return s
}

func (s *QuotaUserViewMetric) SetRequestMemory(v int64) *QuotaUserViewMetric {
	s.RequestMemory = &v
	return s
}

func (s *QuotaUserViewMetric) SetTotalCPU(v int32) *QuotaUserViewMetric {
	s.TotalCPU = &v
	return s
}

func (s *QuotaUserViewMetric) SetTotalGPU(v int32) *QuotaUserViewMetric {
	s.TotalGPU = &v
	return s
}

func (s *QuotaUserViewMetric) SetTotalMemory(v int64) *QuotaUserViewMetric {
	s.TotalMemory = &v
	return s
}

func (s *QuotaUserViewMetric) SetUserId(v string) *QuotaUserViewMetric {
	s.UserId = &v
	return s
}

type ResourceAmount struct {
	// example:
	//
	// 100
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// 16
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// example:
	//
	// GPU
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// example:
	//
	// 100Gi
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ResourceAmount) String() string {
	return tea.Prettify(s)
}

func (s ResourceAmount) GoString() string {
	return s.String()
}

func (s *ResourceAmount) SetCPU(v string) *ResourceAmount {
	s.CPU = &v
	return s
}

func (s *ResourceAmount) SetGPU(v string) *ResourceAmount {
	s.GPU = &v
	return s
}

func (s *ResourceAmount) SetGPUType(v string) *ResourceAmount {
	s.GPUType = &v
	return s
}

func (s *ResourceAmount) SetMemory(v string) *ResourceAmount {
	s.Memory = &v
	return s
}

type ResourceDiagnosisDetail struct {
	ExceedResources []*string       `json:"ExceedResources,omitempty" xml:"ExceedResources,omitempty" type:"Repeated"`
	Limit           *ResourceAmount `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Status          *string         `json:"Status,omitempty" xml:"Status,omitempty"`
	Used            *ResourceAmount `json:"Used,omitempty" xml:"Used,omitempty"`
	WorkloadIds     []*string       `json:"WorkloadIds,omitempty" xml:"WorkloadIds,omitempty" type:"Repeated"`
}

func (s ResourceDiagnosisDetail) String() string {
	return tea.Prettify(s)
}

func (s ResourceDiagnosisDetail) GoString() string {
	return s.String()
}

func (s *ResourceDiagnosisDetail) SetExceedResources(v []*string) *ResourceDiagnosisDetail {
	s.ExceedResources = v
	return s
}

func (s *ResourceDiagnosisDetail) SetLimit(v *ResourceAmount) *ResourceDiagnosisDetail {
	s.Limit = v
	return s
}

func (s *ResourceDiagnosisDetail) SetStatus(v string) *ResourceDiagnosisDetail {
	s.Status = &v
	return s
}

func (s *ResourceDiagnosisDetail) SetUsed(v *ResourceAmount) *ResourceDiagnosisDetail {
	s.Used = v
	return s
}

func (s *ResourceDiagnosisDetail) SetWorkloadIds(v []*string) *ResourceDiagnosisDetail {
	s.WorkloadIds = v
	return s
}

type ResourceGroup struct {
	CreatorID       *string `json:"CreatorID,omitempty" xml:"CreatorID,omitempty"`
	GmtCreatedTime  *string `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// RG1
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NodeCount *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// example:
	//
	// rg17tmvwiokhzaxg
	ResourceGroupID *string  `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	UserVpc         *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
	// example:
	//
	// 23000
	WorkspaceID *string `json:"WorkspaceID,omitempty" xml:"WorkspaceID,omitempty"`
}

func (s ResourceGroup) String() string {
	return tea.Prettify(s)
}

func (s ResourceGroup) GoString() string {
	return s.String()
}

func (s *ResourceGroup) SetCreatorID(v string) *ResourceGroup {
	s.CreatorID = &v
	return s
}

func (s *ResourceGroup) SetGmtCreatedTime(v string) *ResourceGroup {
	s.GmtCreatedTime = &v
	return s
}

func (s *ResourceGroup) SetGmtModifiedTime(v string) *ResourceGroup {
	s.GmtModifiedTime = &v
	return s
}

func (s *ResourceGroup) SetName(v string) *ResourceGroup {
	s.Name = &v
	return s
}

func (s *ResourceGroup) SetNodeCount(v int32) *ResourceGroup {
	s.NodeCount = &v
	return s
}

func (s *ResourceGroup) SetResourceGroupID(v string) *ResourceGroup {
	s.ResourceGroupID = &v
	return s
}

func (s *ResourceGroup) SetUserVpc(v *UserVpc) *ResourceGroup {
	s.UserVpc = v
	return s
}

func (s *ResourceGroup) SetWorkspaceID(v string) *ResourceGroup {
	s.WorkspaceID = &v
	return s
}

type ResourceGroupMetric struct {
	GpuType *string `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
	// example:
	//
	// 23000
	Metrics []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// example:
	//
	// rg17tmvwiokhzaxg
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
}

func (s ResourceGroupMetric) String() string {
	return tea.Prettify(s)
}

func (s ResourceGroupMetric) GoString() string {
	return s.String()
}

func (s *ResourceGroupMetric) SetGpuType(v string) *ResourceGroupMetric {
	s.GpuType = &v
	return s
}

func (s *ResourceGroupMetric) SetMetrics(v []*Metric) *ResourceGroupMetric {
	s.Metrics = v
	return s
}

func (s *ResourceGroupMetric) SetResourceGroupID(v string) *ResourceGroupMetric {
	s.ResourceGroupID = &v
	return s
}

type ResourceOperation struct {
	CreatorId            *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	GmtCreatedTime       *string `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	GmtEndTime           *string `json:"GmtEndTime,omitempty" xml:"GmtEndTime,omitempty"`
	GmtModifiedTime      *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	GmtStartTime         *string `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	ObjectId             *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ObjectType           *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	OperationId          *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	OperationSpecJson    *string `json:"OperationSpecJson,omitempty" xml:"OperationSpecJson,omitempty"`
	OperationType        *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	ReasonCode           *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage        *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ResourceOperation) String() string {
	return tea.Prettify(s)
}

func (s ResourceOperation) GoString() string {
	return s.String()
}

func (s *ResourceOperation) SetCreatorId(v string) *ResourceOperation {
	s.CreatorId = &v
	return s
}

func (s *ResourceOperation) SetGmtCreatedTime(v string) *ResourceOperation {
	s.GmtCreatedTime = &v
	return s
}

func (s *ResourceOperation) SetGmtEndTime(v string) *ResourceOperation {
	s.GmtEndTime = &v
	return s
}

func (s *ResourceOperation) SetGmtModifiedTime(v string) *ResourceOperation {
	s.GmtModifiedTime = &v
	return s
}

func (s *ResourceOperation) SetGmtStartTime(v string) *ResourceOperation {
	s.GmtStartTime = &v
	return s
}

func (s *ResourceOperation) SetObjectId(v string) *ResourceOperation {
	s.ObjectId = &v
	return s
}

func (s *ResourceOperation) SetObjectType(v string) *ResourceOperation {
	s.ObjectType = &v
	return s
}

func (s *ResourceOperation) SetOperationDescription(v string) *ResourceOperation {
	s.OperationDescription = &v
	return s
}

func (s *ResourceOperation) SetOperationId(v string) *ResourceOperation {
	s.OperationId = &v
	return s
}

func (s *ResourceOperation) SetOperationSpecJson(v string) *ResourceOperation {
	s.OperationSpecJson = &v
	return s
}

func (s *ResourceOperation) SetOperationType(v string) *ResourceOperation {
	s.OperationType = &v
	return s
}

func (s *ResourceOperation) SetReasonCode(v string) *ResourceOperation {
	s.ReasonCode = &v
	return s
}

func (s *ResourceOperation) SetReasonMessage(v string) *ResourceOperation {
	s.ReasonMessage = &v
	return s
}

func (s *ResourceOperation) SetStatus(v string) *ResourceOperation {
	s.Status = &v
	return s
}

type ResourceSpec struct {
	NodeSpecs []*NodeSpec `json:"NodeSpecs,omitempty" xml:"NodeSpecs,omitempty" type:"Repeated"`
}

func (s ResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s ResourceSpec) GoString() string {
	return s.String()
}

func (s *ResourceSpec) SetNodeSpecs(v []*NodeSpec) *ResourceSpec {
	s.NodeSpecs = v
	return s
}

type Rules struct {
	SchedulingRule *SchedulingRule `json:"SchedulingRule,omitempty" xml:"SchedulingRule,omitempty"`
}

func (s Rules) String() string {
	return tea.Prettify(s)
}

func (s Rules) GoString() string {
	return s.String()
}

func (s *Rules) SetSchedulingRule(v *SchedulingRule) *Rules {
	s.SchedulingRule = v
	return s
}

type SchedulingRule struct {
	// if can be null:
	// true
	CronTab *string `json:"CronTab,omitempty" xml:"CronTab,omitempty"`
	// if can be null:
	// true
	EndAt *string `json:"EndAt,omitempty" xml:"EndAt,omitempty"`
	// if can be null:
	// true
	ExecuteOnce *bool `json:"ExecuteOnce,omitempty" xml:"ExecuteOnce,omitempty"`
	// if can be null:
	// true
	StartAt *string `json:"StartAt,omitempty" xml:"StartAt,omitempty"`
}

func (s SchedulingRule) String() string {
	return tea.Prettify(s)
}

func (s SchedulingRule) GoString() string {
	return s.String()
}

func (s *SchedulingRule) SetCronTab(v string) *SchedulingRule {
	s.CronTab = &v
	return s
}

func (s *SchedulingRule) SetEndAt(v string) *SchedulingRule {
	s.EndAt = &v
	return s
}

func (s *SchedulingRule) SetExecuteOnce(v bool) *SchedulingRule {
	s.ExecuteOnce = &v
	return s
}

func (s *SchedulingRule) SetStartAt(v string) *SchedulingRule {
	s.StartAt = &v
	return s
}

type SpotPriceItem struct {
	// example:
	//
	// ml.gu8xf.8xlarge-gu108
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 0.1
	SpotDiscount *float32 `json:"SpotDiscount,omitempty" xml:"SpotDiscount,omitempty"`
	// example:
	//
	// 2024-01-17T06:00:00Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// cn-wulanchabu-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SpotPriceItem) String() string {
	return tea.Prettify(s)
}

func (s SpotPriceItem) GoString() string {
	return s.String()
}

func (s *SpotPriceItem) SetInstanceType(v string) *SpotPriceItem {
	s.InstanceType = &v
	return s
}

func (s *SpotPriceItem) SetSpotDiscount(v float32) *SpotPriceItem {
	s.SpotDiscount = &v
	return s
}

func (s *SpotPriceItem) SetTimestamp(v string) *SpotPriceItem {
	s.Timestamp = &v
	return s
}

func (s *SpotPriceItem) SetZoneId(v string) *SpotPriceItem {
	s.ZoneId = &v
	return s
}

type SpotStockPreview struct {
	// example:
	//
	// ml.gu8xf.8xlarge-gu108
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 0.1
	SpotDiscount *float32 `json:"SpotDiscount,omitempty" xml:"SpotDiscount,omitempty"`
	// example:
	//
	// WithStock
	StockStatus *string `json:"StockStatus,omitempty" xml:"StockStatus,omitempty"`
}

func (s SpotStockPreview) String() string {
	return tea.Prettify(s)
}

func (s SpotStockPreview) GoString() string {
	return s.String()
}

func (s *SpotStockPreview) SetInstanceType(v string) *SpotStockPreview {
	s.InstanceType = &v
	return s
}

func (s *SpotStockPreview) SetSpotDiscount(v float32) *SpotStockPreview {
	s.SpotDiscount = &v
	return s
}

func (s *SpotStockPreview) SetStockStatus(v string) *SpotStockPreview {
	s.StockStatus = &v
	return s
}

type Task struct {
	Actions          []*Action `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	Description      *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtActivatedTime *string   `json:"GmtActivatedTime,omitempty" xml:"GmtActivatedTime,omitempty"`
	GmtCreatedTime   *string   `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	GmtModifiedTime  *string   `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	GmtStoppedTime   *string   `json:"GmtStoppedTime,omitempty" xml:"GmtStoppedTime,omitempty"`
	QuotaId          *string   `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	Rules            *Rules    `json:"Rules,omitempty" xml:"Rules,omitempty"`
	Status           *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId           *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName         *string   `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	UserId           *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName         *string   `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Task) String() string {
	return tea.Prettify(s)
}

func (s Task) GoString() string {
	return s.String()
}

func (s *Task) SetActions(v []*Action) *Task {
	s.Actions = v
	return s
}

func (s *Task) SetDescription(v string) *Task {
	s.Description = &v
	return s
}

func (s *Task) SetGmtActivatedTime(v string) *Task {
	s.GmtActivatedTime = &v
	return s
}

func (s *Task) SetGmtCreatedTime(v string) *Task {
	s.GmtCreatedTime = &v
	return s
}

func (s *Task) SetGmtModifiedTime(v string) *Task {
	s.GmtModifiedTime = &v
	return s
}

func (s *Task) SetGmtStoppedTime(v string) *Task {
	s.GmtStoppedTime = &v
	return s
}

func (s *Task) SetQuotaId(v string) *Task {
	s.QuotaId = &v
	return s
}

func (s *Task) SetRules(v *Rules) *Task {
	s.Rules = v
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

func (s *Task) SetTaskName(v string) *Task {
	s.TaskName = &v
	return s
}

func (s *Task) SetUserId(v string) *Task {
	s.UserId = &v
	return s
}

func (s *Task) SetUserName(v string) *Task {
	s.UserName = &v
	return s
}

type TaskInstance struct {
	GmtCreatedTime *string `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	GmtEndTime     *string `json:"GmtEndTime,omitempty" xml:"GmtEndTime,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskInstanceId *string `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
	TenantId       *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s TaskInstance) String() string {
	return tea.Prettify(s)
}

func (s TaskInstance) GoString() string {
	return s.String()
}

func (s *TaskInstance) SetGmtCreatedTime(v string) *TaskInstance {
	s.GmtCreatedTime = &v
	return s
}

func (s *TaskInstance) SetGmtEndTime(v string) *TaskInstance {
	s.GmtEndTime = &v
	return s
}

func (s *TaskInstance) SetStatus(v string) *TaskInstance {
	s.Status = &v
	return s
}

func (s *TaskInstance) SetTaskId(v string) *TaskInstance {
	s.TaskId = &v
	return s
}

func (s *TaskInstance) SetTaskInstanceId(v string) *TaskInstance {
	s.TaskInstanceId = &v
	return s
}

func (s *TaskInstance) SetTenantId(v string) *TaskInstance {
	s.TenantId = &v
	return s
}

func (s *TaskInstance) SetUserId(v string) *TaskInstance {
	s.UserId = &v
	return s
}

type TaskInstanceEvent struct {
	GmtEndTime   *string `json:"GmtEndTime,omitempty" xml:"GmtEndTime,omitempty"`
	GmtStartTime *string `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PodName      *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
}

func (s TaskInstanceEvent) String() string {
	return tea.Prettify(s)
}

func (s TaskInstanceEvent) GoString() string {
	return s.String()
}

func (s *TaskInstanceEvent) SetGmtEndTime(v string) *TaskInstanceEvent {
	s.GmtEndTime = &v
	return s
}

func (s *TaskInstanceEvent) SetGmtStartTime(v string) *TaskInstanceEvent {
	s.GmtStartTime = &v
	return s
}

func (s *TaskInstanceEvent) SetMessage(v string) *TaskInstanceEvent {
	s.Message = &v
	return s
}

func (s *TaskInstanceEvent) SetPodName(v string) *TaskInstanceEvent {
	s.PodName = &v
	return s
}

func (s *TaskInstanceEvent) SetStatus(v string) *TaskInstanceEvent {
	s.Status = &v
	return s
}

func (s *TaskInstanceEvent) SetWorkloadType(v string) *TaskInstanceEvent {
	s.WorkloadType = &v
	return s
}

type TimeRangeFilter struct {
	// example:
	//
	// 2023-06-22T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2023-06-22T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s TimeRangeFilter) String() string {
	return tea.Prettify(s)
}

func (s TimeRangeFilter) GoString() string {
	return s.String()
}

func (s *TimeRangeFilter) SetEndTime(v string) *TimeRangeFilter {
	s.EndTime = &v
	return s
}

func (s *TimeRangeFilter) SetStartTime(v string) *TimeRangeFilter {
	s.StartTime = &v
	return s
}

type UnschedulableNodeDetail struct {
	Nodes  []*string `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	Reason *string   `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s UnschedulableNodeDetail) String() string {
	return tea.Prettify(s)
}

func (s UnschedulableNodeDetail) GoString() string {
	return s.String()
}

func (s *UnschedulableNodeDetail) SetNodes(v []*string) *UnschedulableNodeDetail {
	s.Nodes = v
	return s
}

func (s *UnschedulableNodeDetail) SetReason(v string) *UnschedulableNodeDetail {
	s.Reason = &v
	return s
}

type UserInfo struct {
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UserInfo) String() string {
	return tea.Prettify(s)
}

func (s UserInfo) GoString() string {
	return s.String()
}

func (s *UserInfo) SetUserId(v string) *UserInfo {
	s.UserId = &v
	return s
}

func (s *UserInfo) SetUserName(v string) *UserInfo {
	s.UserName = &v
	return s
}

type UserQuotaPermission struct {
	Permissions []*string `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	QuotaId     *string   `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
}

func (s UserQuotaPermission) String() string {
	return tea.Prettify(s)
}

func (s UserQuotaPermission) GoString() string {
	return s.String()
}

func (s *UserQuotaPermission) SetPermissions(v []*string) *UserQuotaPermission {
	s.Permissions = v
	return s
}

func (s *UserQuotaPermission) SetQuotaId(v string) *UserQuotaPermission {
	s.QuotaId = &v
	return s
}

type UserViewMetric struct {
	CPUNodeNumber     *int32    `json:"CPUNodeNumber,omitempty" xml:"CPUNodeNumber,omitempty"`
	CPUUsageRate      *string   `json:"CPUUsageRate,omitempty" xml:"CPUUsageRate,omitempty"`
	CpuJobNames       []*string `json:"CpuJobNames,omitempty" xml:"CpuJobNames,omitempty" type:"Repeated"`
	CpuNodeNames      []*string `json:"CpuNodeNames,omitempty" xml:"CpuNodeNames,omitempty" type:"Repeated"`
	DiskReadRate      *string   `json:"DiskReadRate,omitempty" xml:"DiskReadRate,omitempty"`
	DiskWriteRate     *string   `json:"DiskWriteRate,omitempty" xml:"DiskWriteRate,omitempty"`
	GPUNodeNumber     *int32    `json:"GPUNodeNumber,omitempty" xml:"GPUNodeNumber,omitempty"`
	GPUUsageRate      *string   `json:"GPUUsageRate,omitempty" xml:"GPUUsageRate,omitempty"`
	GpuJobNames       []*string `json:"GpuJobNames,omitempty" xml:"GpuJobNames,omitempty" type:"Repeated"`
	GpuNodeNames      []*string `json:"GpuNodeNames,omitempty" xml:"GpuNodeNames,omitempty" type:"Repeated"`
	JobType           *string   `json:"JobType,omitempty" xml:"JobType,omitempty"`
	MemoryUsageRate   *string   `json:"MemoryUsageRate,omitempty" xml:"MemoryUsageRate,omitempty"`
	NetworkInputRate  *string   `json:"NetworkInputRate,omitempty" xml:"NetworkInputRate,omitempty"`
	NetworkOutputRate *string   `json:"NetworkOutputRate,omitempty" xml:"NetworkOutputRate,omitempty"`
	NodeNames         []*string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty" type:"Repeated"`
	RequestCPU        *int32    `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	RequestGPU        *int32    `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	RequestMemory     *int64    `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	// example:
	//
	// rg17tmvwiokhzaxg
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TotalCPU        *int32  `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	TotalGPU        *int32  `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	TotalMemory     *int64  `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UserViewMetric) String() string {
	return tea.Prettify(s)
}

func (s UserViewMetric) GoString() string {
	return s.String()
}

func (s *UserViewMetric) SetCPUNodeNumber(v int32) *UserViewMetric {
	s.CPUNodeNumber = &v
	return s
}

func (s *UserViewMetric) SetCPUUsageRate(v string) *UserViewMetric {
	s.CPUUsageRate = &v
	return s
}

func (s *UserViewMetric) SetCpuJobNames(v []*string) *UserViewMetric {
	s.CpuJobNames = v
	return s
}

func (s *UserViewMetric) SetCpuNodeNames(v []*string) *UserViewMetric {
	s.CpuNodeNames = v
	return s
}

func (s *UserViewMetric) SetDiskReadRate(v string) *UserViewMetric {
	s.DiskReadRate = &v
	return s
}

func (s *UserViewMetric) SetDiskWriteRate(v string) *UserViewMetric {
	s.DiskWriteRate = &v
	return s
}

func (s *UserViewMetric) SetGPUNodeNumber(v int32) *UserViewMetric {
	s.GPUNodeNumber = &v
	return s
}

func (s *UserViewMetric) SetGPUUsageRate(v string) *UserViewMetric {
	s.GPUUsageRate = &v
	return s
}

func (s *UserViewMetric) SetGpuJobNames(v []*string) *UserViewMetric {
	s.GpuJobNames = v
	return s
}

func (s *UserViewMetric) SetGpuNodeNames(v []*string) *UserViewMetric {
	s.GpuNodeNames = v
	return s
}

func (s *UserViewMetric) SetJobType(v string) *UserViewMetric {
	s.JobType = &v
	return s
}

func (s *UserViewMetric) SetMemoryUsageRate(v string) *UserViewMetric {
	s.MemoryUsageRate = &v
	return s
}

func (s *UserViewMetric) SetNetworkInputRate(v string) *UserViewMetric {
	s.NetworkInputRate = &v
	return s
}

func (s *UserViewMetric) SetNetworkOutputRate(v string) *UserViewMetric {
	s.NetworkOutputRate = &v
	return s
}

func (s *UserViewMetric) SetNodeNames(v []*string) *UserViewMetric {
	s.NodeNames = v
	return s
}

func (s *UserViewMetric) SetRequestCPU(v int32) *UserViewMetric {
	s.RequestCPU = &v
	return s
}

func (s *UserViewMetric) SetRequestGPU(v int32) *UserViewMetric {
	s.RequestGPU = &v
	return s
}

func (s *UserViewMetric) SetRequestMemory(v int64) *UserViewMetric {
	s.RequestMemory = &v
	return s
}

func (s *UserViewMetric) SetResourceGroupId(v string) *UserViewMetric {
	s.ResourceGroupId = &v
	return s
}

func (s *UserViewMetric) SetTotalCPU(v int32) *UserViewMetric {
	s.TotalCPU = &v
	return s
}

func (s *UserViewMetric) SetTotalGPU(v int32) *UserViewMetric {
	s.TotalGPU = &v
	return s
}

func (s *UserViewMetric) SetTotalMemory(v int64) *UserViewMetric {
	s.TotalMemory = &v
	return s
}

func (s *UserViewMetric) SetUserId(v string) *UserViewMetric {
	s.UserId = &v
	return s
}

type UserVpc struct {
	DefaultForwardInfo *ForwardInfo `json:"DefaultForwardInfo,omitempty" xml:"DefaultForwardInfo,omitempty"`
	DefaultRoute       *string      `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	ExtendedCIDRs      []*string    `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	RoleArn            *string      `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	SecurityGroupId    *string      `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SwitchId           *string      `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	VpcId              *string      `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UserVpc) String() string {
	return tea.Prettify(s)
}

func (s UserVpc) GoString() string {
	return s.String()
}

func (s *UserVpc) SetDefaultForwardInfo(v *ForwardInfo) *UserVpc {
	s.DefaultForwardInfo = v
	return s
}

func (s *UserVpc) SetDefaultRoute(v string) *UserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *UserVpc) SetExtendedCIDRs(v []*string) *UserVpc {
	s.ExtendedCIDRs = v
	return s
}

func (s *UserVpc) SetRoleArn(v string) *UserVpc {
	s.RoleArn = &v
	return s
}

func (s *UserVpc) SetSecurityGroupId(v string) *UserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *UserVpc) SetSwitchId(v string) *UserVpc {
	s.SwitchId = &v
	return s
}

func (s *UserVpc) SetVpcId(v string) *UserVpc {
	s.VpcId = &v
	return s
}

type WorkloadDetails struct {
	DLC     *QuotaJob `json:"DLC,omitempty" xml:"DLC,omitempty"`
	DSW     *QuotaJob `json:"DSW,omitempty" xml:"DSW,omitempty"`
	EAS     *QuotaJob `json:"EAS,omitempty" xml:"EAS,omitempty"`
	Summary *QuotaJob `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s WorkloadDetails) String() string {
	return tea.Prettify(s)
}

func (s WorkloadDetails) GoString() string {
	return s.String()
}

func (s *WorkloadDetails) SetDLC(v *QuotaJob) *WorkloadDetails {
	s.DLC = v
	return s
}

func (s *WorkloadDetails) SetDSW(v *QuotaJob) *WorkloadDetails {
	s.DSW = v
	return s
}

func (s *WorkloadDetails) SetEAS(v *QuotaJob) *WorkloadDetails {
	s.EAS = v
	return s
}

func (s *WorkloadDetails) SetSummary(v *QuotaJob) *WorkloadDetails {
	s.Summary = v
	return s
}

type WorkspaceIdName struct {
	// example:
	//
	// ws123456
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s WorkspaceIdName) String() string {
	return tea.Prettify(s)
}

func (s WorkspaceIdName) GoString() string {
	return s.String()
}

func (s *WorkspaceIdName) SetWorkspaceId(v string) *WorkspaceIdName {
	s.WorkspaceId = &v
	return s
}

type WorkspaceSpec struct {
	Code              *string         `json:"Code,omitempty" xml:"Code,omitempty"`
	CodeType          *string         `json:"CodeType,omitempty" xml:"CodeType,omitempty"`
	IsGuaranteedValid *bool           `json:"IsGuaranteedValid,omitempty" xml:"IsGuaranteedValid,omitempty"`
	IsOverSoldValid   *bool           `json:"IsOverSoldValid,omitempty" xml:"IsOverSoldValid,omitempty"`
	Reason            *string         `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Spec              *ResourceAmount `json:"Spec,omitempty" xml:"Spec,omitempty"`
	SpecName          *string         `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
}

func (s WorkspaceSpec) String() string {
	return tea.Prettify(s)
}

func (s WorkspaceSpec) GoString() string {
	return s.String()
}

func (s *WorkspaceSpec) SetCode(v string) *WorkspaceSpec {
	s.Code = &v
	return s
}

func (s *WorkspaceSpec) SetCodeType(v string) *WorkspaceSpec {
	s.CodeType = &v
	return s
}

func (s *WorkspaceSpec) SetIsGuaranteedValid(v bool) *WorkspaceSpec {
	s.IsGuaranteedValid = &v
	return s
}

func (s *WorkspaceSpec) SetIsOverSoldValid(v bool) *WorkspaceSpec {
	s.IsOverSoldValid = &v
	return s
}

func (s *WorkspaceSpec) SetReason(v string) *WorkspaceSpec {
	s.Reason = &v
	return s
}

func (s *WorkspaceSpec) SetSpec(v *ResourceAmount) *WorkspaceSpec {
	s.Spec = v
	return s
}

func (s *WorkspaceSpec) SetSpecName(v string) *WorkspaceSpec {
	s.SpecName = &v
	return s
}

type WorkspaceSpecs struct {
	Product     *string          `json:"Product,omitempty" xml:"Product,omitempty"`
	Specs       []*WorkspaceSpec `json:"Specs,omitempty" xml:"Specs,omitempty" type:"Repeated"`
	WorkspaceId *string          `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s WorkspaceSpecs) String() string {
	return tea.Prettify(s)
}

func (s WorkspaceSpecs) GoString() string {
	return s.String()
}

func (s *WorkspaceSpecs) SetProduct(v string) *WorkspaceSpecs {
	s.Product = &v
	return s
}

func (s *WorkspaceSpecs) SetSpecs(v []*WorkspaceSpec) *WorkspaceSpecs {
	s.Specs = v
	return s
}

func (s *WorkspaceSpecs) SetWorkspaceId(v string) *WorkspaceSpecs {
	s.WorkspaceId = &v
	return s
}

type CheckInstanceWebTerminalRequest struct {
	// example:
	//
	// wss://pai-dlc-proxy-cn-shanghai.aliyun.com/terminal/t1157703270994901/dlcmjzjt1dxbmx4h/dlcmjzjt1dxbmx4h-worker-0?Token=******
	CheckInfo *string `json:"CheckInfo,omitempty" xml:"CheckInfo,omitempty"`
}

func (s CheckInstanceWebTerminalRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckInstanceWebTerminalRequest) GoString() string {
	return s.String()
}

func (s *CheckInstanceWebTerminalRequest) SetCheckInfo(v string) *CheckInstanceWebTerminalRequest {
	s.CheckInfo = &v
	return s
}

type CheckInstanceWebTerminalResponseBody struct {
	// example:
	//
	// F2D0392B-D749-5C48-A98A-3FAE5C9444A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckInstanceWebTerminalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckInstanceWebTerminalResponseBody) GoString() string {
	return s.String()
}

func (s *CheckInstanceWebTerminalResponseBody) SetRequestId(v string) *CheckInstanceWebTerminalResponseBody {
	s.RequestId = &v
	return s
}

type CheckInstanceWebTerminalResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckInstanceWebTerminalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckInstanceWebTerminalResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckInstanceWebTerminalResponse) GoString() string {
	return s.String()
}

func (s *CheckInstanceWebTerminalResponse) SetHeaders(v map[string]*string) *CheckInstanceWebTerminalResponse {
	s.Headers = v
	return s
}

func (s *CheckInstanceWebTerminalResponse) SetStatusCode(v int32) *CheckInstanceWebTerminalResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckInstanceWebTerminalResponse) SetBody(v *CheckInstanceWebTerminalResponseBody) *CheckInstanceWebTerminalResponse {
	s.Body = v
	return s
}

type CreateAlgorithmRequest struct {
	AlgorithmDescription *string `json:"AlgorithmDescription,omitempty" xml:"AlgorithmDescription,omitempty"`
	// example:
	//
	// llm_training
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	DisplayName   *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateAlgorithmRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlgorithmRequest) GoString() string {
	return s.String()
}

func (s *CreateAlgorithmRequest) SetAlgorithmDescription(v string) *CreateAlgorithmRequest {
	s.AlgorithmDescription = &v
	return s
}

func (s *CreateAlgorithmRequest) SetAlgorithmName(v string) *CreateAlgorithmRequest {
	s.AlgorithmName = &v
	return s
}

func (s *CreateAlgorithmRequest) SetDisplayName(v string) *CreateAlgorithmRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateAlgorithmRequest) SetWorkspaceId(v string) *CreateAlgorithmRequest {
	s.WorkspaceId = &v
	return s
}

type CreateAlgorithmResponseBody struct {
	// example:
	//
	// algo-xsldfvu1334
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	// example:
	//
	// E7C42CC7-2E85-508A-84F4-923B605FD10F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAlgorithmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlgorithmResponseBody) SetAlgorithmId(v string) *CreateAlgorithmResponseBody {
	s.AlgorithmId = &v
	return s
}

func (s *CreateAlgorithmResponseBody) SetRequestId(v string) *CreateAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

type CreateAlgorithmResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAlgorithmResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *CreateAlgorithmResponse) SetHeaders(v map[string]*string) *CreateAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *CreateAlgorithmResponse) SetStatusCode(v int32) *CreateAlgorithmResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlgorithmResponse) SetBody(v *CreateAlgorithmResponseBody) *CreateAlgorithmResponse {
	s.Body = v
	return s
}

type CreateAlgorithmVersionRequest struct {
	AlgorithmSpec *AlgorithmSpec `json:"AlgorithmSpec,omitempty" xml:"AlgorithmSpec,omitempty"`
}

func (s CreateAlgorithmVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlgorithmVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateAlgorithmVersionRequest) SetAlgorithmSpec(v *AlgorithmSpec) *CreateAlgorithmVersionRequest {
	s.AlgorithmSpec = v
	return s
}

type CreateAlgorithmVersionShrinkRequest struct {
	AlgorithmSpecShrink *string `json:"AlgorithmSpec,omitempty" xml:"AlgorithmSpec,omitempty"`
}

func (s CreateAlgorithmVersionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlgorithmVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAlgorithmVersionShrinkRequest) SetAlgorithmSpecShrink(v string) *CreateAlgorithmVersionShrinkRequest {
	s.AlgorithmSpecShrink = &v
	return s
}

type CreateAlgorithmVersionResponseBody struct {
	// example:
	//
	// algo-xsldfvu1334
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	// example:
	//
	// v0.0.1
	AlgorithmVersion *string `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
}

func (s CreateAlgorithmVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAlgorithmVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlgorithmVersionResponseBody) SetAlgorithmId(v string) *CreateAlgorithmVersionResponseBody {
	s.AlgorithmId = &v
	return s
}

func (s *CreateAlgorithmVersionResponseBody) SetAlgorithmVersion(v string) *CreateAlgorithmVersionResponseBody {
	s.AlgorithmVersion = &v
	return s
}

type CreateAlgorithmVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAlgorithmVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAlgorithmVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAlgorithmVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateAlgorithmVersionResponse) SetHeaders(v map[string]*string) *CreateAlgorithmVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateAlgorithmVersionResponse) SetStatusCode(v int32) *CreateAlgorithmVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlgorithmVersionResponse) SetBody(v *CreateAlgorithmVersionResponseBody) *CreateAlgorithmVersionResponse {
	s.Body = v
	return s
}

type CreateInstanceWebTerminalResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// wss://pai-dlc-proxy-cn-shanghai.aliyun.com/terminal/t1157703270994901/dlcmjzjt1dxbmx4h/dlcmjzjt1dxbmx4h-worker-0?Token=******
	WebTerminalId *string `json:"WebTerminalId,omitempty" xml:"WebTerminalId,omitempty"`
}

func (s CreateInstanceWebTerminalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceWebTerminalResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceWebTerminalResponseBody) SetRequestId(v string) *CreateInstanceWebTerminalResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceWebTerminalResponseBody) SetWebTerminalId(v string) *CreateInstanceWebTerminalResponseBody {
	s.WebTerminalId = &v
	return s
}

type CreateInstanceWebTerminalResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceWebTerminalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceWebTerminalResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceWebTerminalResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceWebTerminalResponse) SetHeaders(v map[string]*string) *CreateInstanceWebTerminalResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceWebTerminalResponse) SetStatusCode(v int32) *CreateInstanceWebTerminalResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceWebTerminalResponse) SetBody(v *CreateInstanceWebTerminalResponseBody) *CreateInstanceWebTerminalResponse {
	s.Body = v
	return s
}

type CreateQuotaRequest struct {
	// example:
	//
	// ByNodeSpecs
	AllocateStrategy *string `json:"AllocateStrategy,omitempty" xml:"AllocateStrategy,omitempty"`
	// example:
	//
	// this is a test quota
	Description *string       `json:"Description,omitempty" xml:"Description,omitempty"`
	Labels      []*Label      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Min         *ResourceSpec `json:"Min,omitempty" xml:"Min,omitempty"`
	// example:
	//
	// quota1ci8g793pgm
	ParentQuotaId *string `json:"ParentQuotaId,omitempty" xml:"ParentQuotaId,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// PaiStrategyIntelligent
	QueueStrategy *string      `json:"QueueStrategy,omitempty" xml:"QueueStrategy,omitempty"`
	QuotaConfig   *QuotaConfig `json:"QuotaConfig,omitempty" xml:"QuotaConfig,omitempty"`
	// example:
	//
	// test-quota
	QuotaName        *string   `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
	// example:
	//
	// ECS
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaRequest) GoString() string {
	return s.String()
}

func (s *CreateQuotaRequest) SetAllocateStrategy(v string) *CreateQuotaRequest {
	s.AllocateStrategy = &v
	return s
}

func (s *CreateQuotaRequest) SetDescription(v string) *CreateQuotaRequest {
	s.Description = &v
	return s
}

func (s *CreateQuotaRequest) SetLabels(v []*Label) *CreateQuotaRequest {
	s.Labels = v
	return s
}

func (s *CreateQuotaRequest) SetMin(v *ResourceSpec) *CreateQuotaRequest {
	s.Min = v
	return s
}

func (s *CreateQuotaRequest) SetParentQuotaId(v string) *CreateQuotaRequest {
	s.ParentQuotaId = &v
	return s
}

func (s *CreateQuotaRequest) SetQueueStrategy(v string) *CreateQuotaRequest {
	s.QueueStrategy = &v
	return s
}

func (s *CreateQuotaRequest) SetQuotaConfig(v *QuotaConfig) *CreateQuotaRequest {
	s.QuotaConfig = v
	return s
}

func (s *CreateQuotaRequest) SetQuotaName(v string) *CreateQuotaRequest {
	s.QuotaName = &v
	return s
}

func (s *CreateQuotaRequest) SetResourceGroupIds(v []*string) *CreateQuotaRequest {
	s.ResourceGroupIds = v
	return s
}

func (s *CreateQuotaRequest) SetResourceType(v string) *CreateQuotaRequest {
	s.ResourceType = &v
	return s
}

type CreateQuotaResponseBody struct {
	// Quota Id
	//
	// example:
	//
	// quotad2kd8ljpsno
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// example:
	//
	// CBF05F13-B24C-5129-9048-4FA684DCD579
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQuotaResponseBody) SetQuotaId(v string) *CreateQuotaResponseBody {
	s.QuotaId = &v
	return s
}

func (s *CreateQuotaResponseBody) SetRequestId(v string) *CreateQuotaResponseBody {
	s.RequestId = &v
	return s
}

type CreateQuotaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaResponse) GoString() string {
	return s.String()
}

func (s *CreateQuotaResponse) SetHeaders(v map[string]*string) *CreateQuotaResponse {
	s.Headers = v
	return s
}

func (s *CreateQuotaResponse) SetStatusCode(v int32) *CreateQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQuotaResponse) SetBody(v *CreateQuotaResponseBody) *CreateQuotaResponse {
	s.Body = v
	return s
}

type CreateResourceGroupRequest struct {
	// example:
	//
	// Ecs
	ComputingResourceProvider *string `json:"ComputingResourceProvider,omitempty" xml:"ComputingResourceProvider,omitempty"`
	// example:
	//
	// test_api_report
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// testResourceGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Ecs
	ResourceType *string                          `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*CreateResourceGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	UserVpc      *UserVpc                         `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
}

func (s CreateResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupRequest) SetComputingResourceProvider(v string) *CreateResourceGroupRequest {
	s.ComputingResourceProvider = &v
	return s
}

func (s *CreateResourceGroupRequest) SetDescription(v string) *CreateResourceGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateResourceGroupRequest) SetName(v string) *CreateResourceGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateResourceGroupRequest) SetResourceType(v string) *CreateResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateResourceGroupRequest) SetTag(v []*CreateResourceGroupRequestTag) *CreateResourceGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateResourceGroupRequest) SetUserVpc(v *UserVpc) *CreateResourceGroupRequest {
	s.UserVpc = v
	return s
}

type CreateResourceGroupRequestTag struct {
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateResourceGroupRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupRequestTag) SetKey(v string) *CreateResourceGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateResourceGroupRequestTag) SetValue(v string) *CreateResourceGroupRequestTag {
	s.Value = &v
	return s
}

type CreateResourceGroupResponseBody struct {
	// example:
	//
	// 18D5A1C6-14B8-545E-8408-0A7DDB4C6B5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// ResourceGroup ID。
	//
	// example:
	//
	// rgf0zhfqn1d4ity2
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
}

func (s CreateResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponseBody) SetRequestId(v string) *CreateResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceGroupResponseBody) SetResourceGroupID(v string) *CreateResourceGroupResponseBody {
	s.ResourceGroupID = &v
	return s
}

type CreateResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponse) SetHeaders(v map[string]*string) *CreateResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceGroupResponse) SetStatusCode(v int32) *CreateResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceGroupResponse) SetBody(v *CreateResourceGroupResponseBody) *CreateResourceGroupResponse {
	s.Body = v
	return s
}

type CreateTrainingJobRequest struct {
	// example:
	//
	// ev_classification
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// example:
	//
	// pai
	AlgorithmProvider *string        `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	AlgorithmSpec     *AlgorithmSpec `json:"AlgorithmSpec,omitempty" xml:"AlgorithmSpec,omitempty"`
	// example:
	//
	// v1.0.0
	AlgorithmVersion   *string                                    `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	CodeDir            *Location                                  `json:"CodeDir,omitempty" xml:"CodeDir,omitempty"`
	ComputeResource    *CreateTrainingJobRequestComputeResource   `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	Environments       map[string]*string                         `json:"Environments,omitempty" xml:"Environments,omitempty"`
	ExperimentConfig   *CreateTrainingJobRequestExperimentConfig  `json:"ExperimentConfig,omitempty" xml:"ExperimentConfig,omitempty" type:"Struct"`
	HyperParameters    []*CreateTrainingJobRequestHyperParameters `json:"HyperParameters,omitempty" xml:"HyperParameters,omitempty" type:"Repeated"`
	InputChannels      []*CreateTrainingJobRequestInputChannels   `json:"InputChannels,omitempty" xml:"InputChannels,omitempty" type:"Repeated"`
	Labels             []*CreateTrainingJobRequestLabels          `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	OutputChannels     []*CreateTrainingJobRequestOutputChannels  `json:"OutputChannels,omitempty" xml:"OutputChannels,omitempty" type:"Repeated"`
	Priority           *int32                                     `json:"Priority,omitempty" xml:"Priority,omitempty"`
	PythonRequirements []*string                                  `json:"PythonRequirements,omitempty" xml:"PythonRequirements,omitempty" type:"Repeated"`
	// example:
	//
	// acs:ram::1157703270994901:role/aliyunserviceroleforpaiworkspace
	RoleArn   *string                            `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	Scheduler *CreateTrainingJobRequestScheduler `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Struct"`
	Settings  *JobSettings                       `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// example:
	//
	// qwen large language model training
	TrainingJobDescription *string `json:"TrainingJobDescription,omitempty" xml:"TrainingJobDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qwen_llm
	TrainingJobName *string                          `json:"TrainingJobName,omitempty" xml:"TrainingJobName,omitempty"`
	UserVpc         *CreateTrainingJobRequestUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateTrainingJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainingJobRequest) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequest) SetAlgorithmName(v string) *CreateTrainingJobRequest {
	s.AlgorithmName = &v
	return s
}

func (s *CreateTrainingJobRequest) SetAlgorithmProvider(v string) *CreateTrainingJobRequest {
	s.AlgorithmProvider = &v
	return s
}

func (s *CreateTrainingJobRequest) SetAlgorithmSpec(v *AlgorithmSpec) *CreateTrainingJobRequest {
	s.AlgorithmSpec = v
	return s
}

func (s *CreateTrainingJobRequest) SetAlgorithmVersion(v string) *CreateTrainingJobRequest {
	s.AlgorithmVersion = &v
	return s
}

func (s *CreateTrainingJobRequest) SetCodeDir(v *Location) *CreateTrainingJobRequest {
	s.CodeDir = v
	return s
}

func (s *CreateTrainingJobRequest) SetComputeResource(v *CreateTrainingJobRequestComputeResource) *CreateTrainingJobRequest {
	s.ComputeResource = v
	return s
}

func (s *CreateTrainingJobRequest) SetEnvironments(v map[string]*string) *CreateTrainingJobRequest {
	s.Environments = v
	return s
}

func (s *CreateTrainingJobRequest) SetExperimentConfig(v *CreateTrainingJobRequestExperimentConfig) *CreateTrainingJobRequest {
	s.ExperimentConfig = v
	return s
}

func (s *CreateTrainingJobRequest) SetHyperParameters(v []*CreateTrainingJobRequestHyperParameters) *CreateTrainingJobRequest {
	s.HyperParameters = v
	return s
}

func (s *CreateTrainingJobRequest) SetInputChannels(v []*CreateTrainingJobRequestInputChannels) *CreateTrainingJobRequest {
	s.InputChannels = v
	return s
}

func (s *CreateTrainingJobRequest) SetLabels(v []*CreateTrainingJobRequestLabels) *CreateTrainingJobRequest {
	s.Labels = v
	return s
}

func (s *CreateTrainingJobRequest) SetOutputChannels(v []*CreateTrainingJobRequestOutputChannels) *CreateTrainingJobRequest {
	s.OutputChannels = v
	return s
}

func (s *CreateTrainingJobRequest) SetPriority(v int32) *CreateTrainingJobRequest {
	s.Priority = &v
	return s
}

func (s *CreateTrainingJobRequest) SetPythonRequirements(v []*string) *CreateTrainingJobRequest {
	s.PythonRequirements = v
	return s
}

func (s *CreateTrainingJobRequest) SetRoleArn(v string) *CreateTrainingJobRequest {
	s.RoleArn = &v
	return s
}

func (s *CreateTrainingJobRequest) SetScheduler(v *CreateTrainingJobRequestScheduler) *CreateTrainingJobRequest {
	s.Scheduler = v
	return s
}

func (s *CreateTrainingJobRequest) SetSettings(v *JobSettings) *CreateTrainingJobRequest {
	s.Settings = v
	return s
}

func (s *CreateTrainingJobRequest) SetTrainingJobDescription(v string) *CreateTrainingJobRequest {
	s.TrainingJobDescription = &v
	return s
}

func (s *CreateTrainingJobRequest) SetTrainingJobName(v string) *CreateTrainingJobRequest {
	s.TrainingJobName = &v
	return s
}

func (s *CreateTrainingJobRequest) SetUserVpc(v *CreateTrainingJobRequestUserVpc) *CreateTrainingJobRequest {
	s.UserVpc = v
	return s
}

func (s *CreateTrainingJobRequest) SetWorkspaceId(v string) *CreateTrainingJobRequest {
	s.WorkspaceId = &v
	return s
}

type CreateTrainingJobRequestComputeResource struct {
	// example:
	//
	// 1
	EcsCount *int64 `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	// example:
	//
	// ecs.gn5-c8g1.2xlarge
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// example:
	//
	// 1
	InstanceCount *int64                                               `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	InstanceSpec  *CreateTrainingJobRequestComputeResourceInstanceSpec `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" type:"Struct"`
	// example:
	//
	// quotam670lixikcs
	ResourceId *string                                          `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	SpotSpec   *CreateTrainingJobRequestComputeResourceSpotSpec `json:"SpotSpec,omitempty" xml:"SpotSpec,omitempty" type:"Struct"`
	// example:
	//
	// true
	UseSpotInstance *bool `json:"UseSpotInstance,omitempty" xml:"UseSpotInstance,omitempty"`
}

func (s CreateTrainingJobRequestComputeResource) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainingJobRequestComputeResource) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestComputeResource) SetEcsCount(v int64) *CreateTrainingJobRequestComputeResource {
	s.EcsCount = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResource) SetEcsSpec(v string) *CreateTrainingJobRequestComputeResource {
	s.EcsSpec = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResource) SetInstanceCount(v int64) *CreateTrainingJobRequestComputeResource {
	s.InstanceCount = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResource) SetInstanceSpec(v *CreateTrainingJobRequestComputeResourceInstanceSpec) *CreateTrainingJobRequestComputeResource {
	s.InstanceSpec = v
	return s
}

func (s *CreateTrainingJobRequestComputeResource) SetResourceId(v string) *CreateTrainingJobRequestComputeResource {
	s.ResourceId = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResource) SetSpotSpec(v *CreateTrainingJobRequestComputeResourceSpotSpec) *CreateTrainingJobRequestComputeResource {
	s.SpotSpec = v
	return s
}

func (s *CreateTrainingJobRequestComputeResource) SetUseSpotInstance(v bool) *CreateTrainingJobRequestComputeResource {
	s.UseSpotInstance = &v
	return s
}

type CreateTrainingJobRequestComputeResourceInstanceSpec struct {
	// example:
	//
	// 8
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// 1
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// example:
	//
	// V100
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// example:
	//
	// 32
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// 32
	SharedMemory *string `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
}

func (s CreateTrainingJobRequestComputeResourceInstanceSpec) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainingJobRequestComputeResourceInstanceSpec) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestComputeResourceInstanceSpec) SetCPU(v string) *CreateTrainingJobRequestComputeResourceInstanceSpec {
	s.CPU = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResourceInstanceSpec) SetGPU(v string) *CreateTrainingJobRequestComputeResourceInstanceSpec {
	s.GPU = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResourceInstanceSpec) SetGPUType(v string) *CreateTrainingJobRequestComputeResourceInstanceSpec {
	s.GPUType = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResourceInstanceSpec) SetMemory(v string) *CreateTrainingJobRequestComputeResourceInstanceSpec {
	s.Memory = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResourceInstanceSpec) SetSharedMemory(v string) *CreateTrainingJobRequestComputeResourceInstanceSpec {
	s.SharedMemory = &v
	return s
}

type CreateTrainingJobRequestComputeResourceSpotSpec struct {
	// example:
	//
	// 9
	SpotDiscountLimit *float32 `json:"SpotDiscountLimit,omitempty" xml:"SpotDiscountLimit,omitempty"`
	// example:
	//
	// SpotWithPriceLimit
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
}

func (s CreateTrainingJobRequestComputeResourceSpotSpec) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainingJobRequestComputeResourceSpotSpec) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestComputeResourceSpotSpec) SetSpotDiscountLimit(v float32) *CreateTrainingJobRequestComputeResourceSpotSpec {
	s.SpotDiscountLimit = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResourceSpotSpec) SetSpotStrategy(v string) *CreateTrainingJobRequestComputeResourceSpotSpec {
	s.SpotStrategy = &v
	return s
}

type CreateTrainingJobRequestExperimentConfig struct {
	// example:
	//
	// exp-ds9aefia90v
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
}

func (s CreateTrainingJobRequestExperimentConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainingJobRequestExperimentConfig) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestExperimentConfig) SetExperimentId(v string) *CreateTrainingJobRequestExperimentConfig {
	s.ExperimentId = &v
	return s
}

type CreateTrainingJobRequestHyperParameters struct {
	// example:
	//
	// learning_rate
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.0001
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTrainingJobRequestHyperParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainingJobRequestHyperParameters) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestHyperParameters) SetName(v string) *CreateTrainingJobRequestHyperParameters {
	s.Name = &v
	return s
}

func (s *CreateTrainingJobRequestHyperParameters) SetValue(v string) *CreateTrainingJobRequestHyperParameters {
	s.Value = &v
	return s
}

type CreateTrainingJobRequestInputChannels struct {
	// example:
	//
	// d-475megosidivjfgfq6
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// oss://pai-quickstart-cn-hangzhou.oss-cn-hangzhou-internal.aliyuncs.com/modelscope/models/qwen2-0.5b/main/
	InputUri *string `json:"InputUri,omitempty" xml:"InputUri,omitempty"`
	// example:
	//
	// model
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Options     *string `json:"Options,omitempty" xml:"Options,omitempty"`
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateTrainingJobRequestInputChannels) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainingJobRequestInputChannels) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestInputChannels) SetDatasetId(v string) *CreateTrainingJobRequestInputChannels {
	s.DatasetId = &v
	return s
}

func (s *CreateTrainingJobRequestInputChannels) SetInputUri(v string) *CreateTrainingJobRequestInputChannels {
	s.InputUri = &v
	return s
}

func (s *CreateTrainingJobRequestInputChannels) SetName(v string) *CreateTrainingJobRequestInputChannels {
	s.Name = &v
	return s
}

func (s *CreateTrainingJobRequestInputChannels) SetOptions(v string) *CreateTrainingJobRequestInputChannels {
	s.Options = &v
	return s
}

func (s *CreateTrainingJobRequestInputChannels) SetVersionName(v string) *CreateTrainingJobRequestInputChannels {
	s.VersionName = &v
	return s
}

type CreateTrainingJobRequestLabels struct {
	// example:
	//
	// CreatedBy
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// QuickStart
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTrainingJobRequestLabels) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainingJobRequestLabels) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestLabels) SetKey(v string) *CreateTrainingJobRequestLabels {
	s.Key = &v
	return s
}

func (s *CreateTrainingJobRequestLabels) SetValue(v string) *CreateTrainingJobRequestLabels {
	s.Value = &v
	return s
}

type CreateTrainingJobRequestOutputChannels struct {
	// example:
	//
	// d-475megosidivjfgfq6
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// model
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// oss://pai-quickstart-cn-hangzhou.oss-cn-hangzhou-internal.aliyuncs.com/modelscope/models/qwen2-0.5b/main/
	OutputUri   *string `json:"OutputUri,omitempty" xml:"OutputUri,omitempty"`
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateTrainingJobRequestOutputChannels) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainingJobRequestOutputChannels) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestOutputChannels) SetDatasetId(v string) *CreateTrainingJobRequestOutputChannels {
	s.DatasetId = &v
	return s
}

func (s *CreateTrainingJobRequestOutputChannels) SetName(v string) *CreateTrainingJobRequestOutputChannels {
	s.Name = &v
	return s
}

func (s *CreateTrainingJobRequestOutputChannels) SetOutputUri(v string) *CreateTrainingJobRequestOutputChannels {
	s.OutputUri = &v
	return s
}

func (s *CreateTrainingJobRequestOutputChannels) SetVersionName(v string) *CreateTrainingJobRequestOutputChannels {
	s.VersionName = &v
	return s
}

type CreateTrainingJobRequestScheduler struct {
	MaxRunningTimeInMinutes *int64 `json:"MaxRunningTimeInMinutes,omitempty" xml:"MaxRunningTimeInMinutes,omitempty"`
	// example:
	//
	// 0
	MaxRunningTimeInSeconds *int64 `json:"MaxRunningTimeInSeconds,omitempty" xml:"MaxRunningTimeInSeconds,omitempty"`
}

func (s CreateTrainingJobRequestScheduler) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainingJobRequestScheduler) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestScheduler) SetMaxRunningTimeInMinutes(v int64) *CreateTrainingJobRequestScheduler {
	s.MaxRunningTimeInMinutes = &v
	return s
}

func (s *CreateTrainingJobRequestScheduler) SetMaxRunningTimeInSeconds(v int64) *CreateTrainingJobRequestScheduler {
	s.MaxRunningTimeInSeconds = &v
	return s
}

type CreateTrainingJobRequestUserVpc struct {
	// example:
	//
	// eth0
	DefaultRoute  *string   `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	ExtendedCIDRs []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	// example:
	//
	// sg-qdfasd13sdasf
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// vs-icrc813vdsfol
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-dxiflssjx978sl
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateTrainingJobRequestUserVpc) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainingJobRequestUserVpc) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestUserVpc) SetDefaultRoute(v string) *CreateTrainingJobRequestUserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *CreateTrainingJobRequestUserVpc) SetExtendedCIDRs(v []*string) *CreateTrainingJobRequestUserVpc {
	s.ExtendedCIDRs = v
	return s
}

func (s *CreateTrainingJobRequestUserVpc) SetSecurityGroupId(v string) *CreateTrainingJobRequestUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateTrainingJobRequestUserVpc) SetSwitchId(v string) *CreateTrainingJobRequestUserVpc {
	s.SwitchId = &v
	return s
}

func (s *CreateTrainingJobRequestUserVpc) SetVpcId(v string) *CreateTrainingJobRequestUserVpc {
	s.VpcId = &v
	return s
}

type CreateTrainingJobResponseBody struct {
	// example:
	//
	// E7C42CC7-2E85-508A-84F4-923B605FD10F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// traineyfz0m2hsfv
	TrainingJobId *string `json:"TrainingJobId,omitempty" xml:"TrainingJobId,omitempty"`
}

func (s CreateTrainingJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainingJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobResponseBody) SetRequestId(v string) *CreateTrainingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrainingJobResponseBody) SetTrainingJobId(v string) *CreateTrainingJobResponseBody {
	s.TrainingJobId = &v
	return s
}

type CreateTrainingJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrainingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrainingJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainingJobResponse) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobResponse) SetHeaders(v map[string]*string) *CreateTrainingJobResponse {
	s.Headers = v
	return s
}

func (s *CreateTrainingJobResponse) SetStatusCode(v int32) *CreateTrainingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrainingJobResponse) SetBody(v *CreateTrainingJobResponseBody) *CreateTrainingJobResponse {
	s.Body = v
	return s
}

type DeleteAlgorithmResponseBody struct {
	// example:
	//
	// FFB1D4B4-B253-540A-9B3B-AA711C48A1B7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteAlgorithmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlgorithmResponseBody) SetRequestId(v string) *DeleteAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAlgorithmResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAlgorithmResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlgorithmResponse) SetHeaders(v map[string]*string) *DeleteAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlgorithmResponse) SetStatusCode(v int32) *DeleteAlgorithmResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlgorithmResponse) SetBody(v *DeleteAlgorithmResponseBody) *DeleteAlgorithmResponse {
	s.Body = v
	return s
}

type DeleteAlgorithmVersionResponseBody struct {
	// example:
	//
	// F082BD0D-21E1-5F9B-81A0-AB07485B03CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAlgorithmVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlgorithmVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlgorithmVersionResponseBody) SetRequestId(v string) *DeleteAlgorithmVersionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAlgorithmVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAlgorithmVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAlgorithmVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlgorithmVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlgorithmVersionResponse) SetHeaders(v map[string]*string) *DeleteAlgorithmVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlgorithmVersionResponse) SetStatusCode(v int32) *DeleteAlgorithmVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlgorithmVersionResponse) SetBody(v *DeleteAlgorithmVersionResponseBody) *DeleteAlgorithmVersionResponse {
	s.Body = v
	return s
}

type DeleteMachineGroupResponseBody struct {
	MachineGroupID *string `json:"MachineGroupID,omitempty" xml:"MachineGroupID,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMachineGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMachineGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMachineGroupResponseBody) SetMachineGroupID(v string) *DeleteMachineGroupResponseBody {
	s.MachineGroupID = &v
	return s
}

func (s *DeleteMachineGroupResponseBody) SetRequestId(v string) *DeleteMachineGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMachineGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMachineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMachineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteMachineGroupResponse) SetHeaders(v map[string]*string) *DeleteMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteMachineGroupResponse) SetStatusCode(v int32) *DeleteMachineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMachineGroupResponse) SetBody(v *DeleteMachineGroupResponseBody) *DeleteMachineGroupResponse {
	s.Body = v
	return s
}

type DeleteQuotaResponseBody struct {
	// Quota Id
	//
	// example:
	//
	// quotamtl37ge7gkvdz
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQuotaResponseBody) SetQuotaId(v string) *DeleteQuotaResponseBody {
	s.QuotaId = &v
	return s
}

func (s *DeleteQuotaResponseBody) SetRequestId(v string) *DeleteQuotaResponseBody {
	s.RequestId = &v
	return s
}

type DeleteQuotaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteQuotaResponse) GoString() string {
	return s.String()
}

func (s *DeleteQuotaResponse) SetHeaders(v map[string]*string) *DeleteQuotaResponse {
	s.Headers = v
	return s
}

func (s *DeleteQuotaResponse) SetStatusCode(v int32) *DeleteQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQuotaResponse) SetBody(v *DeleteQuotaResponseBody) *DeleteQuotaResponse {
	s.Body = v
	return s
}

type DeleteResourceGroupResponseBody struct {
	// example:
	//
	// F082BD0D-21E1-5F9B-81A0-AB07485B03CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rgvl9d6utwcscukh
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
}

func (s DeleteResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponseBody) SetRequestId(v string) *DeleteResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceGroupResponseBody) SetResourceGroupID(v string) *DeleteResourceGroupResponseBody {
	s.ResourceGroupID = &v
	return s
}

type DeleteResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponse) SetHeaders(v map[string]*string) *DeleteResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceGroupResponse) SetStatusCode(v int32) *DeleteResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceGroupResponse) SetBody(v *DeleteResourceGroupResponseBody) *DeleteResourceGroupResponse {
	s.Body = v
	return s
}

type DeleteResourceGroupMachineGroupResponseBody struct {
	MachineGroupID *string `json:"MachineGroupID,omitempty" xml:"MachineGroupID,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceGroupMachineGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupMachineGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupMachineGroupResponseBody) SetMachineGroupID(v string) *DeleteResourceGroupMachineGroupResponseBody {
	s.MachineGroupID = &v
	return s
}

func (s *DeleteResourceGroupMachineGroupResponseBody) SetRequestId(v string) *DeleteResourceGroupMachineGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteResourceGroupMachineGroupResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceGroupMachineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceGroupMachineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupMachineGroupResponse) SetHeaders(v map[string]*string) *DeleteResourceGroupMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceGroupMachineGroupResponse) SetStatusCode(v int32) *DeleteResourceGroupMachineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceGroupMachineGroupResponse) SetBody(v *DeleteResourceGroupMachineGroupResponseBody) *DeleteResourceGroupMachineGroupResponse {
	s.Body = v
	return s
}

type DeleteTrainingJobResponseBody struct {
	// example:
	//
	// 4cc83062-9bcb-4ab3-979e-2e571a35834f
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrainingJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrainingJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrainingJobResponseBody) SetRequestId(v string) *DeleteTrainingJobResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTrainingJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrainingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTrainingJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrainingJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrainingJobResponse) SetHeaders(v map[string]*string) *DeleteTrainingJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrainingJobResponse) SetStatusCode(v int32) *DeleteTrainingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrainingJobResponse) SetBody(v *DeleteTrainingJobResponseBody) *DeleteTrainingJobResponse {
	s.Body = v
	return s
}

type DeleteTrainingJobLabelsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// RootModelID
	Keys *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
}

func (s DeleteTrainingJobLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrainingJobLabelsRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrainingJobLabelsRequest) SetKeys(v string) *DeleteTrainingJobLabelsRequest {
	s.Keys = &v
	return s
}

type DeleteTrainingJobLabelsResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrainingJobLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrainingJobLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrainingJobLabelsResponseBody) SetRequestId(v string) *DeleteTrainingJobLabelsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTrainingJobLabelsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrainingJobLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTrainingJobLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrainingJobLabelsResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrainingJobLabelsResponse) SetHeaders(v map[string]*string) *DeleteTrainingJobLabelsResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrainingJobLabelsResponse) SetStatusCode(v int32) *DeleteTrainingJobLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrainingJobLabelsResponse) SetBody(v *DeleteTrainingJobLabelsResponseBody) *DeleteTrainingJobLabelsResponse {
	s.Body = v
	return s
}

type GetAlgorithmResponseBody struct {
	AlgorithmDescription *string `json:"AlgorithmDescription,omitempty" xml:"AlgorithmDescription,omitempty"`
	// example:
	//
	// algo-xsldfvu1334
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	// example:
	//
	// llm_training
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// example:
	//
	// pai
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	// example:
	//
	// llm_training
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 2024-07-10T11:49:47Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2024-07-10T11:49:47Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123456789
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetAlgorithmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlgorithmResponseBody) SetAlgorithmDescription(v string) *GetAlgorithmResponseBody {
	s.AlgorithmDescription = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetAlgorithmId(v string) *GetAlgorithmResponseBody {
	s.AlgorithmId = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetAlgorithmName(v string) *GetAlgorithmResponseBody {
	s.AlgorithmName = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetAlgorithmProvider(v string) *GetAlgorithmResponseBody {
	s.AlgorithmProvider = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetDisplayName(v string) *GetAlgorithmResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetGmtCreateTime(v string) *GetAlgorithmResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetGmtModifiedTime(v string) *GetAlgorithmResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetRequestId(v string) *GetAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetTenantId(v string) *GetAlgorithmResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetUserId(v string) *GetAlgorithmResponseBody {
	s.UserId = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetWorkspaceId(v string) *GetAlgorithmResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetAlgorithmResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlgorithmResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *GetAlgorithmResponse) SetHeaders(v map[string]*string) *GetAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *GetAlgorithmResponse) SetStatusCode(v int32) *GetAlgorithmResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlgorithmResponse) SetBody(v *GetAlgorithmResponseBody) *GetAlgorithmResponse {
	s.Body = v
	return s
}

type GetAlgorithmVersionResponseBody struct {
	// example:
	//
	// algo-xsldfvu1334
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	// example:
	//
	// llm_training
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// example:
	//
	// pai
	AlgorithmProvider *string        `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	AlgorithmSpec     *AlgorithmSpec `json:"AlgorithmSpec,omitempty" xml:"AlgorithmSpec,omitempty"`
	// example:
	//
	// v0.0.1
	AlgorithmVersion *string `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	// example:
	//
	// 2024-07-10T11:49:47Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2024-07-10T11:49:47Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 123456789
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetAlgorithmVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlgorithmVersionResponseBody) SetAlgorithmId(v string) *GetAlgorithmVersionResponseBody {
	s.AlgorithmId = &v
	return s
}

func (s *GetAlgorithmVersionResponseBody) SetAlgorithmName(v string) *GetAlgorithmVersionResponseBody {
	s.AlgorithmName = &v
	return s
}

func (s *GetAlgorithmVersionResponseBody) SetAlgorithmProvider(v string) *GetAlgorithmVersionResponseBody {
	s.AlgorithmProvider = &v
	return s
}

func (s *GetAlgorithmVersionResponseBody) SetAlgorithmSpec(v *AlgorithmSpec) *GetAlgorithmVersionResponseBody {
	s.AlgorithmSpec = v
	return s
}

func (s *GetAlgorithmVersionResponseBody) SetAlgorithmVersion(v string) *GetAlgorithmVersionResponseBody {
	s.AlgorithmVersion = &v
	return s
}

func (s *GetAlgorithmVersionResponseBody) SetGmtCreateTime(v string) *GetAlgorithmVersionResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetAlgorithmVersionResponseBody) SetGmtModifiedTime(v string) *GetAlgorithmVersionResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetAlgorithmVersionResponseBody) SetTenantId(v string) *GetAlgorithmVersionResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetAlgorithmVersionResponseBody) SetUserId(v string) *GetAlgorithmVersionResponseBody {
	s.UserId = &v
	return s
}

type GetAlgorithmVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlgorithmVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlgorithmVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmVersionResponse) GoString() string {
	return s.String()
}

func (s *GetAlgorithmVersionResponse) SetHeaders(v map[string]*string) *GetAlgorithmVersionResponse {
	s.Headers = v
	return s
}

func (s *GetAlgorithmVersionResponse) SetStatusCode(v int32) *GetAlgorithmVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlgorithmVersionResponse) SetBody(v *GetAlgorithmVersionResponseBody) *GetAlgorithmVersionResponse {
	s.Body = v
	return s
}

type GetMachineGroupResponseBody struct {
	Count            *int64    `json:"Count,omitempty" xml:"Count,omitempty"`
	DefaultDriver    *string   `json:"DefaultDriver,omitempty" xml:"DefaultDriver,omitempty"`
	Duration         *string   `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EcsType          *string   `json:"EcsType,omitempty" xml:"EcsType,omitempty"`
	GmtCreated       *string   `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtExpired       *string   `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	GmtModified      *string   `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GmtStarted       *string   `json:"GmtStarted,omitempty" xml:"GmtStarted,omitempty"`
	MachineGroupID   *string   `json:"MachineGroupID,omitempty" xml:"MachineGroupID,omitempty"`
	OrderID          *string   `json:"OrderID,omitempty" xml:"OrderID,omitempty"`
	OrderInstanceId  *string   `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	PAIResourceID    *string   `json:"PAIResourceID,omitempty" xml:"PAIResourceID,omitempty"`
	PayType          *string   `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PricingCycle     *string   `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	RegionID         *string   `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	RequestId        *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status           *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportedDrivers []*string `json:"SupportedDrivers,omitempty" xml:"SupportedDrivers,omitempty" type:"Repeated"`
}

func (s GetMachineGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMachineGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetMachineGroupResponseBody) SetCount(v int64) *GetMachineGroupResponseBody {
	s.Count = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetDefaultDriver(v string) *GetMachineGroupResponseBody {
	s.DefaultDriver = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetDuration(v string) *GetMachineGroupResponseBody {
	s.Duration = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetEcsType(v string) *GetMachineGroupResponseBody {
	s.EcsType = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetGmtCreated(v string) *GetMachineGroupResponseBody {
	s.GmtCreated = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetGmtExpired(v string) *GetMachineGroupResponseBody {
	s.GmtExpired = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetGmtModified(v string) *GetMachineGroupResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetGmtStarted(v string) *GetMachineGroupResponseBody {
	s.GmtStarted = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetMachineGroupID(v string) *GetMachineGroupResponseBody {
	s.MachineGroupID = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetOrderID(v string) *GetMachineGroupResponseBody {
	s.OrderID = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetOrderInstanceId(v string) *GetMachineGroupResponseBody {
	s.OrderInstanceId = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetPAIResourceID(v string) *GetMachineGroupResponseBody {
	s.PAIResourceID = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetPayType(v string) *GetMachineGroupResponseBody {
	s.PayType = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetPricingCycle(v string) *GetMachineGroupResponseBody {
	s.PricingCycle = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetRegionID(v string) *GetMachineGroupResponseBody {
	s.RegionID = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetRequestId(v string) *GetMachineGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetStatus(v string) *GetMachineGroupResponseBody {
	s.Status = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetSupportedDrivers(v []*string) *GetMachineGroupResponseBody {
	s.SupportedDrivers = v
	return s
}

type GetMachineGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMachineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMachineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *GetMachineGroupResponse) SetHeaders(v map[string]*string) *GetMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *GetMachineGroupResponse) SetStatusCode(v int32) *GetMachineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMachineGroupResponse) SetBody(v *GetMachineGroupResponseBody) *GetMachineGroupResponse {
	s.Body = v
	return s
}

type GetNodeMetricsRequest struct {
	// example:
	//
	// 2024-07-10T10:17:06
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// V100
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// example:
	//
	// 2024-07-08T02:23:30.292Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1h
	TimeStep *string `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
	// example:
	//
	// true
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetNodeMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetNodeMetricsRequest) SetEndTime(v string) *GetNodeMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *GetNodeMetricsRequest) SetGPUType(v string) *GetNodeMetricsRequest {
	s.GPUType = &v
	return s
}

func (s *GetNodeMetricsRequest) SetStartTime(v string) *GetNodeMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *GetNodeMetricsRequest) SetTimeStep(v string) *GetNodeMetricsRequest {
	s.TimeStep = &v
	return s
}

func (s *GetNodeMetricsRequest) SetVerbose(v bool) *GetNodeMetricsRequest {
	s.Verbose = &v
	return s
}

type GetNodeMetricsResponseBody struct {
	// example:
	//
	// DiskWriteRate
	MetricType   *string       `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	NodesMetrics []*NodeMetric `json:"NodesMetrics,omitempty" xml:"NodesMetrics,omitempty" type:"Repeated"`
	// example:
	//
	// rgf0zhfqn1d4ity2
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
}

func (s GetNodeMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeMetricsResponseBody) SetMetricType(v string) *GetNodeMetricsResponseBody {
	s.MetricType = &v
	return s
}

func (s *GetNodeMetricsResponseBody) SetNodesMetrics(v []*NodeMetric) *GetNodeMetricsResponseBody {
	s.NodesMetrics = v
	return s
}

func (s *GetNodeMetricsResponseBody) SetResourceGroupID(v string) *GetNodeMetricsResponseBody {
	s.ResourceGroupID = &v
	return s
}

type GetNodeMetricsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetNodeMetricsResponse) SetHeaders(v map[string]*string) *GetNodeMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetNodeMetricsResponse) SetStatusCode(v int32) *GetNodeMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeMetricsResponse) SetBody(v *GetNodeMetricsResponseBody) *GetNodeMetricsResponse {
	s.Body = v
	return s
}

type GetQuotaRequest struct {
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaRequest) SetVerbose(v bool) *GetQuotaRequest {
	s.Verbose = &v
	return s
}

type GetQuotaResponseBody struct {
	// example:
	//
	// ByNodeSpec
	AllocateStrategy *string `json:"AllocateStrategy,omitempty" xml:"AllocateStrategy,omitempty"`
	// example:
	//
	// 18846926616
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// example:
	//
	// this is a test quota
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtCreatedTime *string `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtModifiedTime *string  `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels          []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// example:
	//
	// operation1234
	LatestOperationId *string       `json:"LatestOperationId,omitempty" xml:"LatestOperationId,omitempty"`
	Min               *ResourceSpec `json:"Min,omitempty" xml:"Min,omitempty"`
	// example:
	//
	// quota1ci8g793pgm
	ParentQuotaId *string `json:"ParentQuotaId,omitempty" xml:"ParentQuotaId,omitempty"`
	// example:
	//
	// PaiStrategyIntelligent
	QueueStrategy *string       `json:"QueueStrategy,omitempty" xml:"QueueStrategy,omitempty"`
	QuotaConfig   *QuotaConfig  `json:"QuotaConfig,omitempty" xml:"QuotaConfig,omitempty"`
	QuotaDetails  *QuotaDetails `json:"QuotaDetails,omitempty" xml:"QuotaDetails,omitempty"`
	// Quota Id
	//
	// example:
	//
	// quotajradxh43rgb
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// example:
	//
	// test-quota
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// example:
	//
	// “”
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// example:
	//
	// “”
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// example:
	//
	// 18D5A1C6-14B8-545E-8408-0A7DDB4C6B5E
	RequestId        *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
	// example:
	//
	// ECS
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// Ready
	Status     *string            `json:"Status,omitempty" xml:"Status,omitempty"`
	SubQuotas  []*QuotaIdName     `json:"SubQuotas,omitempty" xml:"SubQuotas,omitempty" type:"Repeated"`
	Workspaces []*WorkspaceIdName `json:"Workspaces,omitempty" xml:"Workspaces,omitempty" type:"Repeated"`
}

func (s GetQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBody) SetAllocateStrategy(v string) *GetQuotaResponseBody {
	s.AllocateStrategy = &v
	return s
}

func (s *GetQuotaResponseBody) SetCreatorId(v string) *GetQuotaResponseBody {
	s.CreatorId = &v
	return s
}

func (s *GetQuotaResponseBody) SetDescription(v string) *GetQuotaResponseBody {
	s.Description = &v
	return s
}

func (s *GetQuotaResponseBody) SetGmtCreatedTime(v string) *GetQuotaResponseBody {
	s.GmtCreatedTime = &v
	return s
}

func (s *GetQuotaResponseBody) SetGmtModifiedTime(v string) *GetQuotaResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetQuotaResponseBody) SetLabels(v []*Label) *GetQuotaResponseBody {
	s.Labels = v
	return s
}

func (s *GetQuotaResponseBody) SetLatestOperationId(v string) *GetQuotaResponseBody {
	s.LatestOperationId = &v
	return s
}

func (s *GetQuotaResponseBody) SetMin(v *ResourceSpec) *GetQuotaResponseBody {
	s.Min = v
	return s
}

func (s *GetQuotaResponseBody) SetParentQuotaId(v string) *GetQuotaResponseBody {
	s.ParentQuotaId = &v
	return s
}

func (s *GetQuotaResponseBody) SetQueueStrategy(v string) *GetQuotaResponseBody {
	s.QueueStrategy = &v
	return s
}

func (s *GetQuotaResponseBody) SetQuotaConfig(v *QuotaConfig) *GetQuotaResponseBody {
	s.QuotaConfig = v
	return s
}

func (s *GetQuotaResponseBody) SetQuotaDetails(v *QuotaDetails) *GetQuotaResponseBody {
	s.QuotaDetails = v
	return s
}

func (s *GetQuotaResponseBody) SetQuotaId(v string) *GetQuotaResponseBody {
	s.QuotaId = &v
	return s
}

func (s *GetQuotaResponseBody) SetQuotaName(v string) *GetQuotaResponseBody {
	s.QuotaName = &v
	return s
}

func (s *GetQuotaResponseBody) SetReasonCode(v string) *GetQuotaResponseBody {
	s.ReasonCode = &v
	return s
}

func (s *GetQuotaResponseBody) SetReasonMessage(v string) *GetQuotaResponseBody {
	s.ReasonMessage = &v
	return s
}

func (s *GetQuotaResponseBody) SetRequestId(v string) *GetQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaResponseBody) SetResourceGroupIds(v []*string) *GetQuotaResponseBody {
	s.ResourceGroupIds = v
	return s
}

func (s *GetQuotaResponseBody) SetResourceType(v string) *GetQuotaResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetQuotaResponseBody) SetStatus(v string) *GetQuotaResponseBody {
	s.Status = &v
	return s
}

func (s *GetQuotaResponseBody) SetSubQuotas(v []*QuotaIdName) *GetQuotaResponseBody {
	s.SubQuotas = v
	return s
}

func (s *GetQuotaResponseBody) SetWorkspaces(v []*WorkspaceIdName) *GetQuotaResponseBody {
	s.Workspaces = v
	return s
}

type GetQuotaResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaResponse) SetHeaders(v map[string]*string) *GetQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaResponse) SetStatusCode(v int32) *GetQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaResponse) SetBody(v *GetQuotaResponseBody) *GetQuotaResponse {
	s.Body = v
	return s
}

type GetResourceGroupRequest struct {
	// example:
	//
	// true
	IsAIWorkspaceDataEnabled *bool                         `json:"IsAIWorkspaceDataEnabled,omitempty" xml:"IsAIWorkspaceDataEnabled,omitempty"`
	Tag                      []*GetResourceGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupRequest) SetIsAIWorkspaceDataEnabled(v bool) *GetResourceGroupRequest {
	s.IsAIWorkspaceDataEnabled = &v
	return s
}

func (s *GetResourceGroupRequest) SetTag(v []*GetResourceGroupRequestTag) *GetResourceGroupRequest {
	s.Tag = v
	return s
}

type GetResourceGroupRequestTag struct {
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetResourceGroupRequestTag) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupRequestTag) GoString() string {
	return s.String()
}

func (s *GetResourceGroupRequestTag) SetKey(v string) *GetResourceGroupRequestTag {
	s.Key = &v
	return s
}

func (s *GetResourceGroupRequestTag) SetValue(v string) *GetResourceGroupRequestTag {
	s.Value = &v
	return s
}

type GetResourceGroupShrinkRequest struct {
	// example:
	//
	// true
	IsAIWorkspaceDataEnabled *bool   `json:"IsAIWorkspaceDataEnabled,omitempty" xml:"IsAIWorkspaceDataEnabled,omitempty"`
	TagShrink                *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetResourceGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupShrinkRequest) SetIsAIWorkspaceDataEnabled(v bool) *GetResourceGroupShrinkRequest {
	s.IsAIWorkspaceDataEnabled = &v
	return s
}

func (s *GetResourceGroupShrinkRequest) SetTagShrink(v string) *GetResourceGroupShrinkRequest {
	s.TagShrink = &v
	return s
}

type GetResourceGroupResponseBody struct {
	// example:
	//
	// cb2c7bde30b774e46a329c
	ClusterID *string `json:"ClusterID,omitempty" xml:"ClusterID,omitempty"`
	// example:
	//
	// ECS
	ComputingResourceProvider *string `json:"ComputingResourceProvider,omitempty" xml:"ComputingResourceProvider,omitempty"`
	// example:
	//
	// 1612285282502324
	CreatorID *string `json:"CreatorID,omitempty" xml:"CreatorID,omitempty"`
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtCreatedTime *string `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// TestResourceGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 18D5A1C6-14B8-545E-8408-0A7DDB4C6B5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Ecs
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// Ready
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	SupportRDMA *bool                               `json:"SupportRDMA,omitempty" xml:"SupportRDMA,omitempty"`
	Tags        []*GetResourceGroupResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	UserVpc     *UserVpc                            `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
	// example:
	//
	// 35201
	WorkspaceID *string `json:"WorkspaceID,omitempty" xml:"WorkspaceID,omitempty"`
}

func (s GetResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBody) SetClusterID(v string) *GetResourceGroupResponseBody {
	s.ClusterID = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetComputingResourceProvider(v string) *GetResourceGroupResponseBody {
	s.ComputingResourceProvider = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetCreatorID(v string) *GetResourceGroupResponseBody {
	s.CreatorID = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetDescription(v string) *GetResourceGroupResponseBody {
	s.Description = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetGmtCreatedTime(v string) *GetResourceGroupResponseBody {
	s.GmtCreatedTime = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetGmtModifiedTime(v string) *GetResourceGroupResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetName(v string) *GetResourceGroupResponseBody {
	s.Name = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetRequestId(v string) *GetResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetResourceType(v string) *GetResourceGroupResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetStatus(v string) *GetResourceGroupResponseBody {
	s.Status = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetSupportRDMA(v bool) *GetResourceGroupResponseBody {
	s.SupportRDMA = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetTags(v []*GetResourceGroupResponseBodyTags) *GetResourceGroupResponseBody {
	s.Tags = v
	return s
}

func (s *GetResourceGroupResponseBody) SetUserVpc(v *UserVpc) *GetResourceGroupResponseBody {
	s.UserVpc = v
	return s
}

func (s *GetResourceGroupResponseBody) SetWorkspaceID(v string) *GetResourceGroupResponseBody {
	s.WorkspaceID = &v
	return s
}

type GetResourceGroupResponseBodyTags struct {
	// example:
	//
	// testKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// testValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetResourceGroupResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBodyTags) SetTagKey(v string) *GetResourceGroupResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *GetResourceGroupResponseBodyTags) SetTagValue(v string) *GetResourceGroupResponseBodyTags {
	s.TagValue = &v
	return s
}

type GetResourceGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponse) SetHeaders(v map[string]*string) *GetResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *GetResourceGroupResponse) SetStatusCode(v int32) *GetResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceGroupResponse) SetBody(v *GetResourceGroupResponseBody) *GetResourceGroupResponse {
	s.Body = v
	return s
}

type GetResourceGroupMachineGroupRequest struct {
	Tag []*GetResourceGroupMachineGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetResourceGroupMachineGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupMachineGroupRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupMachineGroupRequest) SetTag(v []*GetResourceGroupMachineGroupRequestTag) *GetResourceGroupMachineGroupRequest {
	s.Tag = v
	return s
}

type GetResourceGroupMachineGroupRequestTag struct {
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetResourceGroupMachineGroupRequestTag) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupMachineGroupRequestTag) GoString() string {
	return s.String()
}

func (s *GetResourceGroupMachineGroupRequestTag) SetKey(v string) *GetResourceGroupMachineGroupRequestTag {
	s.Key = &v
	return s
}

func (s *GetResourceGroupMachineGroupRequestTag) SetValue(v string) *GetResourceGroupMachineGroupRequestTag {
	s.Value = &v
	return s
}

type GetResourceGroupMachineGroupShrinkRequest struct {
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetResourceGroupMachineGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupMachineGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupMachineGroupShrinkRequest) SetTagShrink(v string) *GetResourceGroupMachineGroupShrinkRequest {
	s.TagShrink = &v
	return s
}

type GetResourceGroupMachineGroupResponseBody struct {
	// example:
	//
	// 2
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 535
	DefaultDriver *string `json:"DefaultDriver,omitempty" xml:"DefaultDriver,omitempty"`
	// example:
	//
	// 1
	EcsCount *int64 `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	// example:
	//
	// ecs.c6.large
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtCreatedTime *string `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtExpiredTime *string `json:"GmtExpiredTime,omitempty" xml:"GmtExpiredTime,omitempty"`
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtStartedTime *string `json:"GmtStartedTime,omitempty" xml:"GmtStartedTime,omitempty"`
	// example:
	//
	// 8
	Gpu *string `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// example:
	//
	// A100
	GpuType *string `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
	// example:
	//
	// mgmioirqjgw6c5lg
	MachineGroupID *string `json:"MachineGroupID,omitempty" xml:"MachineGroupID,omitempty"`
	// example:
	//
	// 64
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// testMachineGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PaymentDuration *string `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	// example:
	//
	// Month
	PaymentDurationUnit *string `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	// example:
	//
	// PREPAY
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// example:
	//
	// 18D5A1C6-14B8-545E-8408-0A7DDB4C6B5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rgf0zhfqn1d4ity2
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	// example:
	//
	// Ready
	Status           *string                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportedDrivers []*string                                       `json:"SupportedDrivers,omitempty" xml:"SupportedDrivers,omitempty" type:"Repeated"`
	Tags             []*GetResourceGroupMachineGroupResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetResourceGroupMachineGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupMachineGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceGroupMachineGroupResponseBody) SetCpu(v string) *GetResourceGroupMachineGroupResponseBody {
	s.Cpu = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetDefaultDriver(v string) *GetResourceGroupMachineGroupResponseBody {
	s.DefaultDriver = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetEcsCount(v int64) *GetResourceGroupMachineGroupResponseBody {
	s.EcsCount = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetEcsSpec(v string) *GetResourceGroupMachineGroupResponseBody {
	s.EcsSpec = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetGmtCreatedTime(v string) *GetResourceGroupMachineGroupResponseBody {
	s.GmtCreatedTime = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetGmtExpiredTime(v string) *GetResourceGroupMachineGroupResponseBody {
	s.GmtExpiredTime = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetGmtModifiedTime(v string) *GetResourceGroupMachineGroupResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetGmtStartedTime(v string) *GetResourceGroupMachineGroupResponseBody {
	s.GmtStartedTime = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetGpu(v string) *GetResourceGroupMachineGroupResponseBody {
	s.Gpu = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetGpuType(v string) *GetResourceGroupMachineGroupResponseBody {
	s.GpuType = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetMachineGroupID(v string) *GetResourceGroupMachineGroupResponseBody {
	s.MachineGroupID = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetMemory(v string) *GetResourceGroupMachineGroupResponseBody {
	s.Memory = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetName(v string) *GetResourceGroupMachineGroupResponseBody {
	s.Name = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetPaymentDuration(v string) *GetResourceGroupMachineGroupResponseBody {
	s.PaymentDuration = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetPaymentDurationUnit(v string) *GetResourceGroupMachineGroupResponseBody {
	s.PaymentDurationUnit = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetPaymentType(v string) *GetResourceGroupMachineGroupResponseBody {
	s.PaymentType = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetRequestId(v string) *GetResourceGroupMachineGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetResourceGroupID(v string) *GetResourceGroupMachineGroupResponseBody {
	s.ResourceGroupID = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetStatus(v string) *GetResourceGroupMachineGroupResponseBody {
	s.Status = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetSupportedDrivers(v []*string) *GetResourceGroupMachineGroupResponseBody {
	s.SupportedDrivers = v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBody) SetTags(v []*GetResourceGroupMachineGroupResponseBodyTags) *GetResourceGroupMachineGroupResponseBody {
	s.Tags = v
	return s
}

type GetResourceGroupMachineGroupResponseBodyTags struct {
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetResourceGroupMachineGroupResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupMachineGroupResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetResourceGroupMachineGroupResponseBodyTags) SetTagKey(v string) *GetResourceGroupMachineGroupResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponseBodyTags) SetTagValue(v string) *GetResourceGroupMachineGroupResponseBodyTags {
	s.TagValue = &v
	return s
}

type GetResourceGroupMachineGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceGroupMachineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceGroupMachineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *GetResourceGroupMachineGroupResponse) SetHeaders(v map[string]*string) *GetResourceGroupMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *GetResourceGroupMachineGroupResponse) SetStatusCode(v int32) *GetResourceGroupMachineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponse) SetBody(v *GetResourceGroupMachineGroupResponseBody) *GetResourceGroupMachineGroupResponse {
	s.Body = v
	return s
}

type GetResourceGroupRequestRequest struct {
	// if can be null:
	// true
	//
	// example:
	//
	// Running
	PodStatus *string `json:"PodStatus,omitempty" xml:"PodStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rgf0zhfqn1d4ity2
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
}

func (s GetResourceGroupRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupRequestRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupRequestRequest) SetPodStatus(v string) *GetResourceGroupRequestRequest {
	s.PodStatus = &v
	return s
}

func (s *GetResourceGroupRequestRequest) SetResourceGroupID(v string) *GetResourceGroupRequestRequest {
	s.ResourceGroupID = &v
	return s
}

type GetResourceGroupRequestResponseBody struct {
	// example:
	//
	// 1
	RequestCPU *int32 `json:"requestCPU,omitempty" xml:"requestCPU,omitempty"`
	// example:
	//
	// 8
	RequestGPU      *int32     `json:"requestGPU,omitempty" xml:"requestGPU,omitempty"`
	RequestGPUInfos []*GPUInfo `json:"requestGPUInfos,omitempty" xml:"requestGPUInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	RequestMemory *int32 `json:"requestMemory,omitempty" xml:"requestMemory,omitempty"`
}

func (s GetResourceGroupRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupRequestResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceGroupRequestResponseBody) SetRequestCPU(v int32) *GetResourceGroupRequestResponseBody {
	s.RequestCPU = &v
	return s
}

func (s *GetResourceGroupRequestResponseBody) SetRequestGPU(v int32) *GetResourceGroupRequestResponseBody {
	s.RequestGPU = &v
	return s
}

func (s *GetResourceGroupRequestResponseBody) SetRequestGPUInfos(v []*GPUInfo) *GetResourceGroupRequestResponseBody {
	s.RequestGPUInfos = v
	return s
}

func (s *GetResourceGroupRequestResponseBody) SetRequestMemory(v int32) *GetResourceGroupRequestResponseBody {
	s.RequestMemory = &v
	return s
}

type GetResourceGroupRequestResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceGroupRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceGroupRequestResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupRequestResponse) GoString() string {
	return s.String()
}

func (s *GetResourceGroupRequestResponse) SetHeaders(v map[string]*string) *GetResourceGroupRequestResponse {
	s.Headers = v
	return s
}

func (s *GetResourceGroupRequestResponse) SetStatusCode(v int32) *GetResourceGroupRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceGroupRequestResponse) SetBody(v *GetResourceGroupRequestResponseBody) *GetResourceGroupRequestResponse {
	s.Body = v
	return s
}

type GetResourceGroupTotalRequest struct {
	// example:
	//
	// rgf0zhfqn1d4ity2
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
}

func (s GetResourceGroupTotalRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupTotalRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupTotalRequest) SetResourceGroupID(v string) *GetResourceGroupTotalRequest {
	s.ResourceGroupID = &v
	return s
}

type GetResourceGroupTotalResponseBody struct {
	// example:
	//
	// 100
	TotalCPU *int32 `json:"totalCPU,omitempty" xml:"totalCPU,omitempty"`
	// example:
	//
	// 24
	TotalGPU      *int32     `json:"totalGPU,omitempty" xml:"totalGPU,omitempty"`
	TotalGPUInfos []*GPUInfo `json:"totalGPUInfos,omitempty" xml:"totalGPUInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 300
	TotalMemory *int32 `json:"totalMemory,omitempty" xml:"totalMemory,omitempty"`
}

func (s GetResourceGroupTotalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupTotalResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceGroupTotalResponseBody) SetTotalCPU(v int32) *GetResourceGroupTotalResponseBody {
	s.TotalCPU = &v
	return s
}

func (s *GetResourceGroupTotalResponseBody) SetTotalGPU(v int32) *GetResourceGroupTotalResponseBody {
	s.TotalGPU = &v
	return s
}

func (s *GetResourceGroupTotalResponseBody) SetTotalGPUInfos(v []*GPUInfo) *GetResourceGroupTotalResponseBody {
	s.TotalGPUInfos = v
	return s
}

func (s *GetResourceGroupTotalResponseBody) SetTotalMemory(v int32) *GetResourceGroupTotalResponseBody {
	s.TotalMemory = &v
	return s
}

type GetResourceGroupTotalResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceGroupTotalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceGroupTotalResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupTotalResponse) GoString() string {
	return s.String()
}

func (s *GetResourceGroupTotalResponse) SetHeaders(v map[string]*string) *GetResourceGroupTotalResponse {
	s.Headers = v
	return s
}

func (s *GetResourceGroupTotalResponse) SetStatusCode(v int32) *GetResourceGroupTotalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceGroupTotalResponse) SetBody(v *GetResourceGroupTotalResponseBody) *GetResourceGroupTotalResponse {
	s.Body = v
	return s
}

type GetTokenRequest struct {
	// example:
	//
	// 60
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// traincclrt205dcs
	TrainingJobId *string `json:"TrainingJobId,omitempty" xml:"TrainingJobId,omitempty"`
}

func (s GetTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) SetExpireTime(v int64) *GetTokenRequest {
	s.ExpireTime = &v
	return s
}

func (s *GetTokenRequest) SetTrainingJobId(v string) *GetTokenRequest {
	s.TrainingJobId = &v
	return s
}

type GetTokenResponseBody struct {
	// example:
	//
	// F2D0392B-D749-5C48-A98A-3FAE5C9444A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ql4OU830nJaF17LP6KTry4a9DvnjIXHP
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBody) SetRequestId(v string) *GetTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenResponseBody) SetToken(v string) *GetTokenResponseBody {
	s.Token = &v
	return s
}

type GetTokenResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponse) GoString() string {
	return s.String()
}

func (s *GetTokenResponse) SetHeaders(v map[string]*string) *GetTokenResponse {
	s.Headers = v
	return s
}

func (s *GetTokenResponse) SetStatusCode(v int32) *GetTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTokenResponse) SetBody(v *GetTokenResponseBody) *GetTokenResponse {
	s.Body = v
	return s
}

type GetTrainingJobResponseBody struct {
	// example:
	//
	// algo-xsldfvu1334
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	// example:
	//
	// llm_training
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// example:
	//
	// pai
	AlgorithmProvider *string        `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	AlgorithmSpec     *AlgorithmSpec `json:"AlgorithmSpec,omitempty" xml:"AlgorithmSpec,omitempty"`
	// example:
	//
	// v0.0.1
	AlgorithmVersion *string                                    `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	ComputeResource  *GetTrainingJobResponseBodyComputeResource `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	// example:
	//
	// 7200
	Duration         *int64                                      `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Environments     map[string]*string                          `json:"Environments,omitempty" xml:"Environments,omitempty"`
	ExperimentConfig *GetTrainingJobResponseBodyExperimentConfig `json:"ExperimentConfig,omitempty" xml:"ExperimentConfig,omitempty" type:"Struct"`
	// example:
	//
	// 2024-07-10T11:49:47Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2024-07-10T11:49:47Z
	GmtModifiedTime *string                                      `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	HyperParameters []*GetTrainingJobResponseBodyHyperParameters `json:"HyperParameters,omitempty" xml:"HyperParameters,omitempty" type:"Repeated"`
	InputChannels   []*GetTrainingJobResponseBodyInputChannels   `json:"InputChannels,omitempty" xml:"InputChannels,omitempty" type:"Repeated"`
	Instances       []*GetTrainingJobResponseBodyInstances       `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// true
	IsTempAlgo         *bool                                       `json:"IsTempAlgo,omitempty" xml:"IsTempAlgo,omitempty"`
	Labels             []*GetTrainingJobResponseBodyLabels         `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LatestMetrics      []*GetTrainingJobResponseBodyLatestMetrics  `json:"LatestMetrics,omitempty" xml:"LatestMetrics,omitempty" type:"Repeated"`
	LatestProgress     *GetTrainingJobResponseBodyLatestProgress   `json:"LatestProgress,omitempty" xml:"LatestProgress,omitempty" type:"Struct"`
	OutputChannels     []*GetTrainingJobResponseBodyOutputChannels `json:"OutputChannels,omitempty" xml:"OutputChannels,omitempty" type:"Repeated"`
	OutputModel        *GetTrainingJobResponseBodyOutputModel      `json:"OutputModel,omitempty" xml:"OutputModel,omitempty" type:"Struct"`
	Priority           *int32                                      `json:"Priority,omitempty" xml:"Priority,omitempty"`
	PythonRequirements []*string                                   `json:"PythonRequirements,omitempty" xml:"PythonRequirements,omitempty" type:"Repeated"`
	// example:
	//
	// TrainingJobSucceed
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// example:
	//
	// None
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// acs:ram::{accountID}:role/{roleName}
	RoleArn   *string                              `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	Scheduler *GetTrainingJobResponseBodyScheduler `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Struct"`
	Settings  *JobSettings                         `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// example:
	//
	// Running
	Status                 *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusTransitions      []*GetTrainingJobResponseBodyStatusTransitions `json:"StatusTransitions,omitempty" xml:"StatusTransitions,omitempty" type:"Repeated"`
	TrainingJobDescription *string                                        `json:"TrainingJobDescription,omitempty" xml:"TrainingJobDescription,omitempty"`
	// example:
	//
	// traini6hhxiq69eo
	TrainingJobId *string `json:"TrainingJobId,omitempty" xml:"TrainingJobId,omitempty"`
	// example:
	//
	// qwen_llm
	TrainingJobName *string `json:"TrainingJobName,omitempty" xml:"TrainingJobName,omitempty"`
	// example:
	//
	// https://pai.console.aliyun.com/?regionId=cn-hangzhou&workspaceId=1234#/training/jobs/train1ouyadsl8n4
	TrainingJobUrl *string `json:"TrainingJobUrl,omitempty" xml:"TrainingJobUrl,omitempty"`
	// example:
	//
	// 123456789
	UserId  *string                            `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserVpc *GetTrainingJobResponseBodyUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// example:
	//
	// 86995
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetTrainingJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBody) SetAlgorithmId(v string) *GetTrainingJobResponseBody {
	s.AlgorithmId = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetAlgorithmName(v string) *GetTrainingJobResponseBody {
	s.AlgorithmName = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetAlgorithmProvider(v string) *GetTrainingJobResponseBody {
	s.AlgorithmProvider = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetAlgorithmSpec(v *AlgorithmSpec) *GetTrainingJobResponseBody {
	s.AlgorithmSpec = v
	return s
}

func (s *GetTrainingJobResponseBody) SetAlgorithmVersion(v string) *GetTrainingJobResponseBody {
	s.AlgorithmVersion = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetComputeResource(v *GetTrainingJobResponseBodyComputeResource) *GetTrainingJobResponseBody {
	s.ComputeResource = v
	return s
}

func (s *GetTrainingJobResponseBody) SetDuration(v int64) *GetTrainingJobResponseBody {
	s.Duration = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetEnvironments(v map[string]*string) *GetTrainingJobResponseBody {
	s.Environments = v
	return s
}

func (s *GetTrainingJobResponseBody) SetExperimentConfig(v *GetTrainingJobResponseBodyExperimentConfig) *GetTrainingJobResponseBody {
	s.ExperimentConfig = v
	return s
}

func (s *GetTrainingJobResponseBody) SetGmtCreateTime(v string) *GetTrainingJobResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetGmtModifiedTime(v string) *GetTrainingJobResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetHyperParameters(v []*GetTrainingJobResponseBodyHyperParameters) *GetTrainingJobResponseBody {
	s.HyperParameters = v
	return s
}

func (s *GetTrainingJobResponseBody) SetInputChannels(v []*GetTrainingJobResponseBodyInputChannels) *GetTrainingJobResponseBody {
	s.InputChannels = v
	return s
}

func (s *GetTrainingJobResponseBody) SetInstances(v []*GetTrainingJobResponseBodyInstances) *GetTrainingJobResponseBody {
	s.Instances = v
	return s
}

func (s *GetTrainingJobResponseBody) SetIsTempAlgo(v bool) *GetTrainingJobResponseBody {
	s.IsTempAlgo = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetLabels(v []*GetTrainingJobResponseBodyLabels) *GetTrainingJobResponseBody {
	s.Labels = v
	return s
}

func (s *GetTrainingJobResponseBody) SetLatestMetrics(v []*GetTrainingJobResponseBodyLatestMetrics) *GetTrainingJobResponseBody {
	s.LatestMetrics = v
	return s
}

func (s *GetTrainingJobResponseBody) SetLatestProgress(v *GetTrainingJobResponseBodyLatestProgress) *GetTrainingJobResponseBody {
	s.LatestProgress = v
	return s
}

func (s *GetTrainingJobResponseBody) SetOutputChannels(v []*GetTrainingJobResponseBodyOutputChannels) *GetTrainingJobResponseBody {
	s.OutputChannels = v
	return s
}

func (s *GetTrainingJobResponseBody) SetOutputModel(v *GetTrainingJobResponseBodyOutputModel) *GetTrainingJobResponseBody {
	s.OutputModel = v
	return s
}

func (s *GetTrainingJobResponseBody) SetPriority(v int32) *GetTrainingJobResponseBody {
	s.Priority = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetPythonRequirements(v []*string) *GetTrainingJobResponseBody {
	s.PythonRequirements = v
	return s
}

func (s *GetTrainingJobResponseBody) SetReasonCode(v string) *GetTrainingJobResponseBody {
	s.ReasonCode = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetReasonMessage(v string) *GetTrainingJobResponseBody {
	s.ReasonMessage = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetRequestId(v string) *GetTrainingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetRoleArn(v string) *GetTrainingJobResponseBody {
	s.RoleArn = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetScheduler(v *GetTrainingJobResponseBodyScheduler) *GetTrainingJobResponseBody {
	s.Scheduler = v
	return s
}

func (s *GetTrainingJobResponseBody) SetSettings(v *JobSettings) *GetTrainingJobResponseBody {
	s.Settings = v
	return s
}

func (s *GetTrainingJobResponseBody) SetStatus(v string) *GetTrainingJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetStatusTransitions(v []*GetTrainingJobResponseBodyStatusTransitions) *GetTrainingJobResponseBody {
	s.StatusTransitions = v
	return s
}

func (s *GetTrainingJobResponseBody) SetTrainingJobDescription(v string) *GetTrainingJobResponseBody {
	s.TrainingJobDescription = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetTrainingJobId(v string) *GetTrainingJobResponseBody {
	s.TrainingJobId = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetTrainingJobName(v string) *GetTrainingJobResponseBody {
	s.TrainingJobName = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetTrainingJobUrl(v string) *GetTrainingJobResponseBody {
	s.TrainingJobUrl = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetUserId(v string) *GetTrainingJobResponseBody {
	s.UserId = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetUserVpc(v *GetTrainingJobResponseBodyUserVpc) *GetTrainingJobResponseBody {
	s.UserVpc = v
	return s
}

func (s *GetTrainingJobResponseBody) SetWorkspaceId(v string) *GetTrainingJobResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetTrainingJobResponseBodyComputeResource struct {
	// example:
	//
	// 1
	EcsCount *int64 `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	// example:
	//
	// ecs.gn5-c8g1.2xlarge
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// example:
	//
	// 1
	InstanceCount *int64                                                 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	InstanceSpec  *GetTrainingJobResponseBodyComputeResourceInstanceSpec `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" type:"Struct"`
	// example:
	//
	// quotam670lixikcl
	ResourceId   *string                                            `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName *string                                            `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	SpotSpec     *GetTrainingJobResponseBodyComputeResourceSpotSpec `json:"SpotSpec,omitempty" xml:"SpotSpec,omitempty" type:"Struct"`
	// example:
	//
	// true
	UseSpotInstance *bool `json:"UseSpotInstance,omitempty" xml:"UseSpotInstance,omitempty"`
}

func (s GetTrainingJobResponseBodyComputeResource) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBodyComputeResource) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyComputeResource) SetEcsCount(v int64) *GetTrainingJobResponseBodyComputeResource {
	s.EcsCount = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResource) SetEcsSpec(v string) *GetTrainingJobResponseBodyComputeResource {
	s.EcsSpec = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResource) SetInstanceCount(v int64) *GetTrainingJobResponseBodyComputeResource {
	s.InstanceCount = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResource) SetInstanceSpec(v *GetTrainingJobResponseBodyComputeResourceInstanceSpec) *GetTrainingJobResponseBodyComputeResource {
	s.InstanceSpec = v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResource) SetResourceId(v string) *GetTrainingJobResponseBodyComputeResource {
	s.ResourceId = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResource) SetResourceName(v string) *GetTrainingJobResponseBodyComputeResource {
	s.ResourceName = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResource) SetSpotSpec(v *GetTrainingJobResponseBodyComputeResourceSpotSpec) *GetTrainingJobResponseBodyComputeResource {
	s.SpotSpec = v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResource) SetUseSpotInstance(v bool) *GetTrainingJobResponseBodyComputeResource {
	s.UseSpotInstance = &v
	return s
}

type GetTrainingJobResponseBodyComputeResourceInstanceSpec struct {
	// example:
	//
	// 8
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// 1
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// example:
	//
	// V100
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// example:
	//
	// 32
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// 32
	SharedMemory *string `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
}

func (s GetTrainingJobResponseBodyComputeResourceInstanceSpec) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBodyComputeResourceInstanceSpec) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyComputeResourceInstanceSpec) SetCPU(v string) *GetTrainingJobResponseBodyComputeResourceInstanceSpec {
	s.CPU = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResourceInstanceSpec) SetGPU(v string) *GetTrainingJobResponseBodyComputeResourceInstanceSpec {
	s.GPU = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResourceInstanceSpec) SetGPUType(v string) *GetTrainingJobResponseBodyComputeResourceInstanceSpec {
	s.GPUType = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResourceInstanceSpec) SetMemory(v string) *GetTrainingJobResponseBodyComputeResourceInstanceSpec {
	s.Memory = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResourceInstanceSpec) SetSharedMemory(v string) *GetTrainingJobResponseBodyComputeResourceInstanceSpec {
	s.SharedMemory = &v
	return s
}

type GetTrainingJobResponseBodyComputeResourceSpotSpec struct {
	// example:
	//
	// 0.9
	SpotDiscountLimit *float32 `json:"SpotDiscountLimit,omitempty" xml:"SpotDiscountLimit,omitempty"`
	// example:
	//
	// SpotWithPriceLimit
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
}

func (s GetTrainingJobResponseBodyComputeResourceSpotSpec) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBodyComputeResourceSpotSpec) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyComputeResourceSpotSpec) SetSpotDiscountLimit(v float32) *GetTrainingJobResponseBodyComputeResourceSpotSpec {
	s.SpotDiscountLimit = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResourceSpotSpec) SetSpotStrategy(v string) *GetTrainingJobResponseBodyComputeResourceSpotSpec {
	s.SpotStrategy = &v
	return s
}

type GetTrainingJobResponseBodyExperimentConfig struct {
	// example:
	//
	// exp-ds9aefia90v
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// large_language_model_train
	ExperimentName *string `json:"ExperimentName,omitempty" xml:"ExperimentName,omitempty"`
}

func (s GetTrainingJobResponseBodyExperimentConfig) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBodyExperimentConfig) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyExperimentConfig) SetExperimentId(v string) *GetTrainingJobResponseBodyExperimentConfig {
	s.ExperimentId = &v
	return s
}

func (s *GetTrainingJobResponseBodyExperimentConfig) SetExperimentName(v string) *GetTrainingJobResponseBodyExperimentConfig {
	s.ExperimentName = &v
	return s
}

type GetTrainingJobResponseBodyHyperParameters struct {
	// example:
	//
	// learning_rate
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.0001
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTrainingJobResponseBodyHyperParameters) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBodyHyperParameters) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyHyperParameters) SetName(v string) *GetTrainingJobResponseBodyHyperParameters {
	s.Name = &v
	return s
}

func (s *GetTrainingJobResponseBodyHyperParameters) SetValue(v string) *GetTrainingJobResponseBodyHyperParameters {
	s.Value = &v
	return s
}

type GetTrainingJobResponseBodyInputChannels struct {
	// example:
	//
	// d-475megosidivjfgfq6
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// oss://test-bucket.oss-cn-hangzhou-internal.aliyuncs.com/path/to/input/model/
	InputUri *string `json:"InputUri,omitempty" xml:"InputUri,omitempty"`
	// example:
	//
	// model
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Options     *string `json:"Options,omitempty" xml:"Options,omitempty"`
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetTrainingJobResponseBodyInputChannels) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBodyInputChannels) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyInputChannels) SetDatasetId(v string) *GetTrainingJobResponseBodyInputChannels {
	s.DatasetId = &v
	return s
}

func (s *GetTrainingJobResponseBodyInputChannels) SetInputUri(v string) *GetTrainingJobResponseBodyInputChannels {
	s.InputUri = &v
	return s
}

func (s *GetTrainingJobResponseBodyInputChannels) SetName(v string) *GetTrainingJobResponseBodyInputChannels {
	s.Name = &v
	return s
}

func (s *GetTrainingJobResponseBodyInputChannels) SetOptions(v string) *GetTrainingJobResponseBodyInputChannels {
	s.Options = &v
	return s
}

func (s *GetTrainingJobResponseBodyInputChannels) SetVersionName(v string) *GetTrainingJobResponseBodyInputChannels {
	s.VersionName = &v
	return s
}

type GetTrainingJobResponseBodyInstances struct {
	// example:
	//
	// train1oug3yehan4-master-0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// master
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTrainingJobResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyInstances) SetName(v string) *GetTrainingJobResponseBodyInstances {
	s.Name = &v
	return s
}

func (s *GetTrainingJobResponseBodyInstances) SetRole(v string) *GetTrainingJobResponseBodyInstances {
	s.Role = &v
	return s
}

func (s *GetTrainingJobResponseBodyInstances) SetStatus(v string) *GetTrainingJobResponseBodyInstances {
	s.Status = &v
	return s
}

type GetTrainingJobResponseBodyLabels struct {
	// example:
	//
	// CreatedBy
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// QuickStart
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTrainingJobResponseBodyLabels) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyLabels) SetKey(v string) *GetTrainingJobResponseBodyLabels {
	s.Key = &v
	return s
}

func (s *GetTrainingJobResponseBodyLabels) SetValue(v string) *GetTrainingJobResponseBodyLabels {
	s.Value = &v
	return s
}

type GetTrainingJobResponseBodyLatestMetrics struct {
	// example:
	//
	// loss
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2024-07-10T11:49:47Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 0.11
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTrainingJobResponseBodyLatestMetrics) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBodyLatestMetrics) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyLatestMetrics) SetName(v string) *GetTrainingJobResponseBodyLatestMetrics {
	s.Name = &v
	return s
}

func (s *GetTrainingJobResponseBodyLatestMetrics) SetTimestamp(v string) *GetTrainingJobResponseBodyLatestMetrics {
	s.Timestamp = &v
	return s
}

func (s *GetTrainingJobResponseBodyLatestMetrics) SetValue(v float64) *GetTrainingJobResponseBodyLatestMetrics {
	s.Value = &v
	return s
}

type GetTrainingJobResponseBodyLatestProgress struct {
	OverallProgress *GetTrainingJobResponseBodyLatestProgressOverallProgress `json:"OverallProgress,omitempty" xml:"OverallProgress,omitempty" type:"Struct"`
	RemainingTime   *GetTrainingJobResponseBodyLatestProgressRemainingTime   `json:"RemainingTime,omitempty" xml:"RemainingTime,omitempty" type:"Struct"`
}

func (s GetTrainingJobResponseBodyLatestProgress) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBodyLatestProgress) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyLatestProgress) SetOverallProgress(v *GetTrainingJobResponseBodyLatestProgressOverallProgress) *GetTrainingJobResponseBodyLatestProgress {
	s.OverallProgress = v
	return s
}

func (s *GetTrainingJobResponseBodyLatestProgress) SetRemainingTime(v *GetTrainingJobResponseBodyLatestProgressRemainingTime) *GetTrainingJobResponseBodyLatestProgress {
	s.RemainingTime = v
	return s
}

type GetTrainingJobResponseBodyLatestProgressOverallProgress struct {
	// example:
	//
	// 2023-07-04T13:20:18Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 0.75
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTrainingJobResponseBodyLatestProgressOverallProgress) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBodyLatestProgressOverallProgress) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyLatestProgressOverallProgress) SetTimestamp(v string) *GetTrainingJobResponseBodyLatestProgressOverallProgress {
	s.Timestamp = &v
	return s
}

func (s *GetTrainingJobResponseBodyLatestProgressOverallProgress) SetValue(v float32) *GetTrainingJobResponseBodyLatestProgressOverallProgress {
	s.Value = &v
	return s
}

type GetTrainingJobResponseBodyLatestProgressRemainingTime struct {
	// example:
	//
	// 2023-07-04T13:20:18Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 3600
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTrainingJobResponseBodyLatestProgressRemainingTime) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBodyLatestProgressRemainingTime) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyLatestProgressRemainingTime) SetTimestamp(v string) *GetTrainingJobResponseBodyLatestProgressRemainingTime {
	s.Timestamp = &v
	return s
}

func (s *GetTrainingJobResponseBodyLatestProgressRemainingTime) SetValue(v int64) *GetTrainingJobResponseBodyLatestProgressRemainingTime {
	s.Value = &v
	return s
}

type GetTrainingJobResponseBodyOutputChannels struct {
	// example:
	//
	// d-8o0hh35po15ejcdq2p
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// model
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// oss://test-bucket.oss-cn-hangzhou-internal.aliyuncs.com/path/to/output/model/
	OutputUri   *string `json:"OutputUri,omitempty" xml:"OutputUri,omitempty"`
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetTrainingJobResponseBodyOutputChannels) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBodyOutputChannels) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyOutputChannels) SetDatasetId(v string) *GetTrainingJobResponseBodyOutputChannels {
	s.DatasetId = &v
	return s
}

func (s *GetTrainingJobResponseBodyOutputChannels) SetName(v string) *GetTrainingJobResponseBodyOutputChannels {
	s.Name = &v
	return s
}

func (s *GetTrainingJobResponseBodyOutputChannels) SetOutputUri(v string) *GetTrainingJobResponseBodyOutputChannels {
	s.OutputUri = &v
	return s
}

func (s *GetTrainingJobResponseBodyOutputChannels) SetVersionName(v string) *GetTrainingJobResponseBodyOutputChannels {
	s.VersionName = &v
	return s
}

type GetTrainingJobResponseBodyOutputModel struct {
	// example:
	//
	// model
	OutputChannelName *string `json:"OutputChannelName,omitempty" xml:"OutputChannelName,omitempty"`
	// example:
	//
	// oss://test-bucket.oss-cn-hangzhou-internal.aliyuncs.com/path/to/model/output/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s GetTrainingJobResponseBodyOutputModel) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBodyOutputModel) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyOutputModel) SetOutputChannelName(v string) *GetTrainingJobResponseBodyOutputModel {
	s.OutputChannelName = &v
	return s
}

func (s *GetTrainingJobResponseBodyOutputModel) SetUri(v string) *GetTrainingJobResponseBodyOutputModel {
	s.Uri = &v
	return s
}

type GetTrainingJobResponseBodyScheduler struct {
	MaxRunningTimeInMinutes *string `json:"MaxRunningTimeInMinutes,omitempty" xml:"MaxRunningTimeInMinutes,omitempty"`
	// example:
	//
	// 0
	MaxRunningTimeInSeconds *string `json:"MaxRunningTimeInSeconds,omitempty" xml:"MaxRunningTimeInSeconds,omitempty"`
}

func (s GetTrainingJobResponseBodyScheduler) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBodyScheduler) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyScheduler) SetMaxRunningTimeInMinutes(v string) *GetTrainingJobResponseBodyScheduler {
	s.MaxRunningTimeInMinutes = &v
	return s
}

func (s *GetTrainingJobResponseBodyScheduler) SetMaxRunningTimeInSeconds(v string) *GetTrainingJobResponseBodyScheduler {
	s.MaxRunningTimeInSeconds = &v
	return s
}

type GetTrainingJobResponseBodyStatusTransitions struct {
	// example:
	//
	// 2024-07-10T11:49:47Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// TrainingJobSucceed
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// example:
	//
	// KubeDL job runs successfully
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// example:
	//
	// 2024-07-10T11:49:47Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTrainingJobResponseBodyStatusTransitions) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBodyStatusTransitions) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyStatusTransitions) SetEndTime(v string) *GetTrainingJobResponseBodyStatusTransitions {
	s.EndTime = &v
	return s
}

func (s *GetTrainingJobResponseBodyStatusTransitions) SetReasonCode(v string) *GetTrainingJobResponseBodyStatusTransitions {
	s.ReasonCode = &v
	return s
}

func (s *GetTrainingJobResponseBodyStatusTransitions) SetReasonMessage(v string) *GetTrainingJobResponseBodyStatusTransitions {
	s.ReasonMessage = &v
	return s
}

func (s *GetTrainingJobResponseBodyStatusTransitions) SetStartTime(v string) *GetTrainingJobResponseBodyStatusTransitions {
	s.StartTime = &v
	return s
}

func (s *GetTrainingJobResponseBodyStatusTransitions) SetStatus(v string) *GetTrainingJobResponseBodyStatusTransitions {
	s.Status = &v
	return s
}

type GetTrainingJobResponseBodyUserVpc struct {
	ExtendedCIDRs []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	// example:
	//
	// sg-abcdef****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// vs-abcdef****
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-abcdef****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetTrainingJobResponseBodyUserVpc) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBodyUserVpc) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyUserVpc) SetExtendedCIDRs(v []*string) *GetTrainingJobResponseBodyUserVpc {
	s.ExtendedCIDRs = v
	return s
}

func (s *GetTrainingJobResponseBodyUserVpc) SetSecurityGroupId(v string) *GetTrainingJobResponseBodyUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *GetTrainingJobResponseBodyUserVpc) SetSwitchId(v string) *GetTrainingJobResponseBodyUserVpc {
	s.SwitchId = &v
	return s
}

func (s *GetTrainingJobResponseBodyUserVpc) SetVpcId(v string) *GetTrainingJobResponseBodyUserVpc {
	s.VpcId = &v
	return s
}

type GetTrainingJobResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrainingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrainingJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponse) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponse) SetHeaders(v map[string]*string) *GetTrainingJobResponse {
	s.Headers = v
	return s
}

func (s *GetTrainingJobResponse) SetStatusCode(v int32) *GetTrainingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrainingJobResponse) SetBody(v *GetTrainingJobResponseBody) *GetTrainingJobResponse {
	s.Body = v
	return s
}

type GetTrainingJobErrorInfoResponseBody struct {
	ErrorInfo *GetTrainingJobErrorInfoResponseBodyErrorInfo `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty" type:"Struct"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTrainingJobErrorInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobErrorInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrainingJobErrorInfoResponseBody) SetErrorInfo(v *GetTrainingJobErrorInfoResponseBodyErrorInfo) *GetTrainingJobErrorInfoResponseBody {
	s.ErrorInfo = v
	return s
}

func (s *GetTrainingJobErrorInfoResponseBody) SetRequestId(v string) *GetTrainingJobErrorInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetTrainingJobErrorInfoResponseBodyErrorInfo struct {
	// example:
	//
	// additional info
	AdditionalInfo *string `json:"AdditionalInfo,omitempty" xml:"AdditionalInfo,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GetTrainingJobErrorInfoResponseBodyErrorInfo) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobErrorInfoResponseBodyErrorInfo) GoString() string {
	return s.String()
}

func (s *GetTrainingJobErrorInfoResponseBodyErrorInfo) SetAdditionalInfo(v string) *GetTrainingJobErrorInfoResponseBodyErrorInfo {
	s.AdditionalInfo = &v
	return s
}

func (s *GetTrainingJobErrorInfoResponseBodyErrorInfo) SetCode(v string) *GetTrainingJobErrorInfoResponseBodyErrorInfo {
	s.Code = &v
	return s
}

func (s *GetTrainingJobErrorInfoResponseBodyErrorInfo) SetMessage(v string) *GetTrainingJobErrorInfoResponseBodyErrorInfo {
	s.Message = &v
	return s
}

type GetTrainingJobErrorInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrainingJobErrorInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrainingJobErrorInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobErrorInfoResponse) GoString() string {
	return s.String()
}

func (s *GetTrainingJobErrorInfoResponse) SetHeaders(v map[string]*string) *GetTrainingJobErrorInfoResponse {
	s.Headers = v
	return s
}

func (s *GetTrainingJobErrorInfoResponse) SetStatusCode(v int32) *GetTrainingJobErrorInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrainingJobErrorInfoResponse) SetBody(v *GetTrainingJobErrorInfoResponseBody) *GetTrainingJobErrorInfoResponse {
	s.Body = v
	return s
}

type GetTrainingJobLatestMetricsRequest struct {
	// example:
	//
	// loss
	Names *string `json:"Names,omitempty" xml:"Names,omitempty"`
}

func (s GetTrainingJobLatestMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobLatestMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetTrainingJobLatestMetricsRequest) SetNames(v string) *GetTrainingJobLatestMetricsRequest {
	s.Names = &v
	return s
}

type GetTrainingJobLatestMetricsResponseBody struct {
	Metrics []*GetTrainingJobLatestMetricsResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// example:
	//
	// 18D5A1C6-14B8-545E-8408-0A7DDB4C6B5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTrainingJobLatestMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobLatestMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrainingJobLatestMetricsResponseBody) SetMetrics(v []*GetTrainingJobLatestMetricsResponseBodyMetrics) *GetTrainingJobLatestMetricsResponseBody {
	s.Metrics = v
	return s
}

func (s *GetTrainingJobLatestMetricsResponseBody) SetRequestId(v string) *GetTrainingJobLatestMetricsResponseBody {
	s.RequestId = &v
	return s
}

type GetTrainingJobLatestMetricsResponseBodyMetrics struct {
	// example:
	//
	// loss
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-04-18T22:20:55Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 0.97
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTrainingJobLatestMetricsResponseBodyMetrics) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobLatestMetricsResponseBodyMetrics) GoString() string {
	return s.String()
}

func (s *GetTrainingJobLatestMetricsResponseBodyMetrics) SetName(v string) *GetTrainingJobLatestMetricsResponseBodyMetrics {
	s.Name = &v
	return s
}

func (s *GetTrainingJobLatestMetricsResponseBodyMetrics) SetTimestamp(v string) *GetTrainingJobLatestMetricsResponseBodyMetrics {
	s.Timestamp = &v
	return s
}

func (s *GetTrainingJobLatestMetricsResponseBodyMetrics) SetValue(v float64) *GetTrainingJobLatestMetricsResponseBodyMetrics {
	s.Value = &v
	return s
}

type GetTrainingJobLatestMetricsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrainingJobLatestMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrainingJobLatestMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobLatestMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetTrainingJobLatestMetricsResponse) SetHeaders(v map[string]*string) *GetTrainingJobLatestMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetTrainingJobLatestMetricsResponse) SetStatusCode(v int32) *GetTrainingJobLatestMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrainingJobLatestMetricsResponse) SetBody(v *GetTrainingJobLatestMetricsResponseBody) *GetTrainingJobLatestMetricsResponse {
	s.Body = v
	return s
}

type GetUserViewMetricsRequest struct {
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtModified
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// 1h
	TimeStep *string `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
	// example:
	//
	// 123456
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 86995
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetUserViewMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserViewMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetUserViewMetricsRequest) SetOrder(v string) *GetUserViewMetricsRequest {
	s.Order = &v
	return s
}

func (s *GetUserViewMetricsRequest) SetPageNumber(v string) *GetUserViewMetricsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetUserViewMetricsRequest) SetPageSize(v string) *GetUserViewMetricsRequest {
	s.PageSize = &v
	return s
}

func (s *GetUserViewMetricsRequest) SetSortBy(v string) *GetUserViewMetricsRequest {
	s.SortBy = &v
	return s
}

func (s *GetUserViewMetricsRequest) SetTimeStep(v string) *GetUserViewMetricsRequest {
	s.TimeStep = &v
	return s
}

func (s *GetUserViewMetricsRequest) SetUserId(v string) *GetUserViewMetricsRequest {
	s.UserId = &v
	return s
}

func (s *GetUserViewMetricsRequest) SetWorkspaceId(v string) *GetUserViewMetricsRequest {
	s.WorkspaceId = &v
	return s
}

type GetUserViewMetricsResponseBody struct {
	// example:
	//
	// rgf0zhfqn1d4ity2
	ResourceGroupId *string         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Summary         *UserViewMetric `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 2
	Total       *int32            `json:"Total,omitempty" xml:"Total,omitempty"`
	UserMetrics []*UserViewMetric `json:"UserMetrics,omitempty" xml:"UserMetrics,omitempty" type:"Repeated"`
}

func (s GetUserViewMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserViewMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserViewMetricsResponseBody) SetResourceGroupId(v string) *GetUserViewMetricsResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetUserViewMetricsResponseBody) SetSummary(v *UserViewMetric) *GetUserViewMetricsResponseBody {
	s.Summary = v
	return s
}

func (s *GetUserViewMetricsResponseBody) SetTotal(v int32) *GetUserViewMetricsResponseBody {
	s.Total = &v
	return s
}

func (s *GetUserViewMetricsResponseBody) SetUserMetrics(v []*UserViewMetric) *GetUserViewMetricsResponseBody {
	s.UserMetrics = v
	return s
}

type GetUserViewMetricsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserViewMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserViewMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserViewMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetUserViewMetricsResponse) SetHeaders(v map[string]*string) *GetUserViewMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetUserViewMetricsResponse) SetStatusCode(v int32) *GetUserViewMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserViewMetricsResponse) SetBody(v *GetUserViewMetricsResponseBody) *GetUserViewMetricsResponse {
	s.Body = v
	return s
}

type ListAlgorithmVersionsRequest struct {
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAlgorithmVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlgorithmVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListAlgorithmVersionsRequest) SetPageNumber(v int64) *ListAlgorithmVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlgorithmVersionsRequest) SetPageSize(v int64) *ListAlgorithmVersionsRequest {
	s.PageSize = &v
	return s
}

type ListAlgorithmVersionsResponseBody struct {
	AlgorithmVersions []*ListAlgorithmVersionsResponseBodyAlgorithmVersions `json:"AlgorithmVersions,omitempty" xml:"AlgorithmVersions,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 23
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAlgorithmVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlgorithmVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlgorithmVersionsResponseBody) SetAlgorithmVersions(v []*ListAlgorithmVersionsResponseBodyAlgorithmVersions) *ListAlgorithmVersionsResponseBody {
	s.AlgorithmVersions = v
	return s
}

func (s *ListAlgorithmVersionsResponseBody) SetRequestId(v string) *ListAlgorithmVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlgorithmVersionsResponseBody) SetTotalCount(v int64) *ListAlgorithmVersionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAlgorithmVersionsResponseBodyAlgorithmVersions struct {
	// example:
	//
	// algo-sidjc8134hv
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	// example:
	//
	// llm_train
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// example:
	//
	// pai
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	// example:
	//
	// v0.1.0
	AlgorithmVersion *string `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	// example:
	//
	// 2024-01-19T02:00:26Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2024-01-22T02:00:59Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 123456789
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListAlgorithmVersionsResponseBodyAlgorithmVersions) String() string {
	return tea.Prettify(s)
}

func (s ListAlgorithmVersionsResponseBodyAlgorithmVersions) GoString() string {
	return s.String()
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) SetAlgorithmId(v string) *ListAlgorithmVersionsResponseBodyAlgorithmVersions {
	s.AlgorithmId = &v
	return s
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) SetAlgorithmName(v string) *ListAlgorithmVersionsResponseBodyAlgorithmVersions {
	s.AlgorithmName = &v
	return s
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) SetAlgorithmProvider(v string) *ListAlgorithmVersionsResponseBodyAlgorithmVersions {
	s.AlgorithmProvider = &v
	return s
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) SetAlgorithmVersion(v string) *ListAlgorithmVersionsResponseBodyAlgorithmVersions {
	s.AlgorithmVersion = &v
	return s
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) SetGmtCreateTime(v string) *ListAlgorithmVersionsResponseBodyAlgorithmVersions {
	s.GmtCreateTime = &v
	return s
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) SetGmtModifiedTime(v string) *ListAlgorithmVersionsResponseBodyAlgorithmVersions {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) SetTenantId(v string) *ListAlgorithmVersionsResponseBodyAlgorithmVersions {
	s.TenantId = &v
	return s
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) SetUserId(v string) *ListAlgorithmVersionsResponseBodyAlgorithmVersions {
	s.UserId = &v
	return s
}

type ListAlgorithmVersionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlgorithmVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlgorithmVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlgorithmVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListAlgorithmVersionsResponse) SetHeaders(v map[string]*string) *ListAlgorithmVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListAlgorithmVersionsResponse) SetStatusCode(v int32) *ListAlgorithmVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlgorithmVersionsResponse) SetBody(v *ListAlgorithmVersionsResponseBody) *ListAlgorithmVersionsResponse {
	s.Body = v
	return s
}

type ListAlgorithmsRequest struct {
	// example:
	//
	// algo-xsldfvu1334
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	// example:
	//
	// llm_training
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// example:
	//
	// pai
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListAlgorithmsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlgorithmsRequest) GoString() string {
	return s.String()
}

func (s *ListAlgorithmsRequest) SetAlgorithmId(v string) *ListAlgorithmsRequest {
	s.AlgorithmId = &v
	return s
}

func (s *ListAlgorithmsRequest) SetAlgorithmName(v string) *ListAlgorithmsRequest {
	s.AlgorithmName = &v
	return s
}

func (s *ListAlgorithmsRequest) SetAlgorithmProvider(v string) *ListAlgorithmsRequest {
	s.AlgorithmProvider = &v
	return s
}

func (s *ListAlgorithmsRequest) SetPageNumber(v int64) *ListAlgorithmsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlgorithmsRequest) SetPageSize(v int64) *ListAlgorithmsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlgorithmsRequest) SetWorkspaceId(v string) *ListAlgorithmsRequest {
	s.WorkspaceId = &v
	return s
}

type ListAlgorithmsResponseBody struct {
	Algorithms []*ListAlgorithmsResponseBodyAlgorithms `json:"Algorithms,omitempty" xml:"Algorithms,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 4
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAlgorithmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlgorithmsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlgorithmsResponseBody) SetAlgorithms(v []*ListAlgorithmsResponseBodyAlgorithms) *ListAlgorithmsResponseBody {
	s.Algorithms = v
	return s
}

func (s *ListAlgorithmsResponseBody) SetRequestId(v string) *ListAlgorithmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlgorithmsResponseBody) SetTotalCount(v int64) *ListAlgorithmsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAlgorithmsResponseBodyAlgorithms struct {
	AlgorithmDescription *string `json:"AlgorithmDescription,omitempty" xml:"AlgorithmDescription,omitempty"`
	// example:
	//
	// algo-sidjc8134hv
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	// example:
	//
	// llm_train
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// example:
	//
	// pai
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	// example:
	//
	// LLM Train
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 2023-07-21T03:35:24Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-07-25T02:15:40Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListAlgorithmsResponseBodyAlgorithms) String() string {
	return tea.Prettify(s)
}

func (s ListAlgorithmsResponseBodyAlgorithms) GoString() string {
	return s.String()
}

func (s *ListAlgorithmsResponseBodyAlgorithms) SetAlgorithmDescription(v string) *ListAlgorithmsResponseBodyAlgorithms {
	s.AlgorithmDescription = &v
	return s
}

func (s *ListAlgorithmsResponseBodyAlgorithms) SetAlgorithmId(v string) *ListAlgorithmsResponseBodyAlgorithms {
	s.AlgorithmId = &v
	return s
}

func (s *ListAlgorithmsResponseBodyAlgorithms) SetAlgorithmName(v string) *ListAlgorithmsResponseBodyAlgorithms {
	s.AlgorithmName = &v
	return s
}

func (s *ListAlgorithmsResponseBodyAlgorithms) SetAlgorithmProvider(v string) *ListAlgorithmsResponseBodyAlgorithms {
	s.AlgorithmProvider = &v
	return s
}

func (s *ListAlgorithmsResponseBodyAlgorithms) SetDisplayName(v string) *ListAlgorithmsResponseBodyAlgorithms {
	s.DisplayName = &v
	return s
}

func (s *ListAlgorithmsResponseBodyAlgorithms) SetGmtCreateTime(v string) *ListAlgorithmsResponseBodyAlgorithms {
	s.GmtCreateTime = &v
	return s
}

func (s *ListAlgorithmsResponseBodyAlgorithms) SetGmtModifiedTime(v string) *ListAlgorithmsResponseBodyAlgorithms {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListAlgorithmsResponseBodyAlgorithms) SetUserId(v string) *ListAlgorithmsResponseBodyAlgorithms {
	s.UserId = &v
	return s
}

func (s *ListAlgorithmsResponseBodyAlgorithms) SetWorkspaceId(v string) *ListAlgorithmsResponseBodyAlgorithms {
	s.WorkspaceId = &v
	return s
}

type ListAlgorithmsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlgorithmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlgorithmsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlgorithmsResponse) GoString() string {
	return s.String()
}

func (s *ListAlgorithmsResponse) SetHeaders(v map[string]*string) *ListAlgorithmsResponse {
	s.Headers = v
	return s
}

func (s *ListAlgorithmsResponse) SetStatusCode(v int32) *ListAlgorithmsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlgorithmsResponse) SetBody(v *ListAlgorithmsResponseBody) *ListAlgorithmsResponse {
	s.Body = v
	return s
}

type ListNodesRequest struct {
	// example:
	//
	// CPU
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// example:
	//
	// quotamtl37ge7gkvdz
	FilterByQuotaId *string `json:"FilterByQuotaId,omitempty" xml:"FilterByQuotaId,omitempty"`
	// example:
	//
	// rg69rj0leslwdnbe
	FilterByResourceGroupIds *string `json:"FilterByResourceGroupIds,omitempty" xml:"FilterByResourceGroupIds,omitempty"`
	// example:
	//
	// T4
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// example:
	//
	// lingjxxxx
	NodeNames *string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty"`
	// example:
	//
	// Ready
	NodeStatuses *string `json:"NodeStatuses,omitempty" xml:"NodeStatuses,omitempty"`
	// example:
	//
	// ecs.c6.xlarge
	NodeTypes *string `json:"NodeTypes,omitempty" xml:"NodeTypes,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// Ready
	OrderStatuses *string `json:"OrderStatuses,omitempty" xml:"OrderStatuses,omitempty"`
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// quotamtl37ge7gkvdz
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// example:
	//
	// rg69rj0leslwdnbe
	ResourceGroupIds *string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// false
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s ListNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) SetAcceleratorType(v string) *ListNodesRequest {
	s.AcceleratorType = &v
	return s
}

func (s *ListNodesRequest) SetFilterByQuotaId(v string) *ListNodesRequest {
	s.FilterByQuotaId = &v
	return s
}

func (s *ListNodesRequest) SetFilterByResourceGroupIds(v string) *ListNodesRequest {
	s.FilterByResourceGroupIds = &v
	return s
}

func (s *ListNodesRequest) SetGPUType(v string) *ListNodesRequest {
	s.GPUType = &v
	return s
}

func (s *ListNodesRequest) SetNodeNames(v string) *ListNodesRequest {
	s.NodeNames = &v
	return s
}

func (s *ListNodesRequest) SetNodeStatuses(v string) *ListNodesRequest {
	s.NodeStatuses = &v
	return s
}

func (s *ListNodesRequest) SetNodeTypes(v string) *ListNodesRequest {
	s.NodeTypes = &v
	return s
}

func (s *ListNodesRequest) SetOrder(v string) *ListNodesRequest {
	s.Order = &v
	return s
}

func (s *ListNodesRequest) SetOrderStatuses(v string) *ListNodesRequest {
	s.OrderStatuses = &v
	return s
}

func (s *ListNodesRequest) SetPageNumber(v int32) *ListNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodesRequest) SetPageSize(v int32) *ListNodesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodesRequest) SetQuotaId(v string) *ListNodesRequest {
	s.QuotaId = &v
	return s
}

func (s *ListNodesRequest) SetResourceGroupIds(v string) *ListNodesRequest {
	s.ResourceGroupIds = &v
	return s
}

func (s *ListNodesRequest) SetSortBy(v string) *ListNodesRequest {
	s.SortBy = &v
	return s
}

func (s *ListNodesRequest) SetVerbose(v bool) *ListNodesRequest {
	s.Verbose = &v
	return s
}

type ListNodesResponseBody struct {
	Nodes []*Node `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 23
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) SetNodes(v []*Node) *ListNodesResponseBody {
	s.Nodes = v
	return s
}

func (s *ListNodesResponseBody) SetRequestId(v string) *ListNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesResponseBody) SetTotalCount(v int32) *ListNodesResponseBody {
	s.TotalCount = &v
	return s
}

type ListNodesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponse) GoString() string {
	return s.String()
}

func (s *ListNodesResponse) SetHeaders(v map[string]*string) *ListNodesResponse {
	s.Headers = v
	return s
}

func (s *ListNodesResponse) SetStatusCode(v int32) *ListNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodesResponse) SetBody(v *ListNodesResponseBody) *ListNodesResponse {
	s.Body = v
	return s
}

type ListQuotaWorkloadsRequest struct {
	// example:
	//
	// dsw65443322
	BeforeWorkloadId             *string          `json:"BeforeWorkloadId,omitempty" xml:"BeforeWorkloadId,omitempty"`
	GmtDequeuedTimeRange         *TimeRangeFilter `json:"GmtDequeuedTimeRange,omitempty" xml:"GmtDequeuedTimeRange,omitempty"`
	GmtEnqueuedTimeRange         *TimeRangeFilter `json:"GmtEnqueuedTimeRange,omitempty" xml:"GmtEnqueuedTimeRange,omitempty"`
	GmtPositionModifiedTimeRange *TimeRangeFilter `json:"GmtPositionModifiedTimeRange,omitempty" xml:"GmtPositionModifiedTimeRange,omitempty"`
	// example:
	//
	// lrn48278127617
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// true
	ShowOwn *bool `json:"ShowOwn,omitempty" xml:"ShowOwn,omitempty"`
	// example:
	//
	// GmtCreatedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// Enqueued
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// quota12344666,quota64432233
	SubQuotaIds *string `json:"SubQuotaIds,omitempty" xml:"SubQuotaIds,omitempty"`
	// example:
	//
	// 29043893812,23829093093
	UserIds                  *string          `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
	WorkloadCreatedTimeRange *TimeRangeFilter `json:"WorkloadCreatedTimeRange,omitempty" xml:"WorkloadCreatedTimeRange,omitempty"`
	// example:
	//
	// dlc12344556
	WorkloadIds *string `json:"WorkloadIds,omitempty" xml:"WorkloadIds,omitempty"`
	// example:
	//
	// Pending
	WorkloadStatuses *string `json:"WorkloadStatuses,omitempty" xml:"WorkloadStatuses,omitempty"`
	// example:
	//
	// dlc
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
	// example:
	//
	// 186692
	WorkspaceIds *string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty"`
}

func (s ListQuotaWorkloadsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaWorkloadsRequest) GoString() string {
	return s.String()
}

func (s *ListQuotaWorkloadsRequest) SetBeforeWorkloadId(v string) *ListQuotaWorkloadsRequest {
	s.BeforeWorkloadId = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetGmtDequeuedTimeRange(v *TimeRangeFilter) *ListQuotaWorkloadsRequest {
	s.GmtDequeuedTimeRange = v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetGmtEnqueuedTimeRange(v *TimeRangeFilter) *ListQuotaWorkloadsRequest {
	s.GmtEnqueuedTimeRange = v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetGmtPositionModifiedTimeRange(v *TimeRangeFilter) *ListQuotaWorkloadsRequest {
	s.GmtPositionModifiedTimeRange = v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetNodeName(v string) *ListQuotaWorkloadsRequest {
	s.NodeName = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetOrder(v string) *ListQuotaWorkloadsRequest {
	s.Order = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetPageNumber(v int32) *ListQuotaWorkloadsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetPageSize(v int32) *ListQuotaWorkloadsRequest {
	s.PageSize = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetShowOwn(v bool) *ListQuotaWorkloadsRequest {
	s.ShowOwn = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetSortBy(v string) *ListQuotaWorkloadsRequest {
	s.SortBy = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetStatus(v string) *ListQuotaWorkloadsRequest {
	s.Status = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetSubQuotaIds(v string) *ListQuotaWorkloadsRequest {
	s.SubQuotaIds = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetUserIds(v string) *ListQuotaWorkloadsRequest {
	s.UserIds = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetWorkloadCreatedTimeRange(v *TimeRangeFilter) *ListQuotaWorkloadsRequest {
	s.WorkloadCreatedTimeRange = v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetWorkloadIds(v string) *ListQuotaWorkloadsRequest {
	s.WorkloadIds = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetWorkloadStatuses(v string) *ListQuotaWorkloadsRequest {
	s.WorkloadStatuses = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetWorkloadType(v string) *ListQuotaWorkloadsRequest {
	s.WorkloadType = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetWorkspaceIds(v string) *ListQuotaWorkloadsRequest {
	s.WorkspaceIds = &v
	return s
}

type ListQuotaWorkloadsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 42F23B58-3684-5443-848A-8DA81FF99712
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 23
	TotalCount *int64       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Workloads  []*QueueInfo `json:"Workloads,omitempty" xml:"Workloads,omitempty" type:"Repeated"`
}

func (s ListQuotaWorkloadsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaWorkloadsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotaWorkloadsResponseBody) SetRequestId(v string) *ListQuotaWorkloadsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotaWorkloadsResponseBody) SetTotalCount(v int64) *ListQuotaWorkloadsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListQuotaWorkloadsResponseBody) SetWorkloads(v []*QueueInfo) *ListQuotaWorkloadsResponseBody {
	s.Workloads = v
	return s
}

type ListQuotaWorkloadsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotaWorkloadsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotaWorkloadsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaWorkloadsResponse) GoString() string {
	return s.String()
}

func (s *ListQuotaWorkloadsResponse) SetHeaders(v map[string]*string) *ListQuotaWorkloadsResponse {
	s.Headers = v
	return s
}

func (s *ListQuotaWorkloadsResponse) SetStatusCode(v int32) *ListQuotaWorkloadsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotaWorkloadsResponse) SetBody(v *ListQuotaWorkloadsResponseBody) *ListQuotaWorkloadsResponse {
	s.Body = v
	return s
}

type ListQuotasRequest struct {
	// example:
	//
	// official=true,gpu=false
	Labels     *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	LayoutMode *string `json:"LayoutMode,omitempty" xml:"LayoutMode,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// quotajradxh43rgb
	ParentQuotaId *string `json:"ParentQuotaId,omitempty" xml:"ParentQuotaId,omitempty"`
	// example:
	//
	// quota1ci8g793pgm,quotajradxh43rgb
	QuotaIds *string `json:"QuotaIds,omitempty" xml:"QuotaIds,omitempty"`
	// example:
	//
	// quotajradxh43rgb
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// example:
	//
	// ECS
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// status
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// Creating
	Statuses *string `json:"Statuses,omitempty" xml:"Statuses,omitempty"`
	Verbose  *bool   `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// example:
	//
	// 21345,38727
	WorkspaceIds  *string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty"`
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListQuotasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasRequest) GoString() string {
	return s.String()
}

func (s *ListQuotasRequest) SetLabels(v string) *ListQuotasRequest {
	s.Labels = &v
	return s
}

func (s *ListQuotasRequest) SetLayoutMode(v string) *ListQuotasRequest {
	s.LayoutMode = &v
	return s
}

func (s *ListQuotasRequest) SetOrder(v string) *ListQuotasRequest {
	s.Order = &v
	return s
}

func (s *ListQuotasRequest) SetPageNumber(v int32) *ListQuotasRequest {
	s.PageNumber = &v
	return s
}

func (s *ListQuotasRequest) SetPageSize(v int32) *ListQuotasRequest {
	s.PageSize = &v
	return s
}

func (s *ListQuotasRequest) SetParentQuotaId(v string) *ListQuotasRequest {
	s.ParentQuotaId = &v
	return s
}

func (s *ListQuotasRequest) SetQuotaIds(v string) *ListQuotasRequest {
	s.QuotaIds = &v
	return s
}

func (s *ListQuotasRequest) SetQuotaName(v string) *ListQuotasRequest {
	s.QuotaName = &v
	return s
}

func (s *ListQuotasRequest) SetResourceType(v string) *ListQuotasRequest {
	s.ResourceType = &v
	return s
}

func (s *ListQuotasRequest) SetSortBy(v string) *ListQuotasRequest {
	s.SortBy = &v
	return s
}

func (s *ListQuotasRequest) SetStatuses(v string) *ListQuotasRequest {
	s.Statuses = &v
	return s
}

func (s *ListQuotasRequest) SetVerbose(v bool) *ListQuotasRequest {
	s.Verbose = &v
	return s
}

func (s *ListQuotasRequest) SetWorkspaceIds(v string) *ListQuotasRequest {
	s.WorkspaceIds = &v
	return s
}

func (s *ListQuotasRequest) SetWorkspaceName(v string) *ListQuotasRequest {
	s.WorkspaceName = &v
	return s
}

type ListQuotasResponseBody struct {
	Quotas []*Quota `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// example:
	//
	// F082BD0D-21E1-5F9B-81A0-AB07485B03CD
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQuotasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBody) SetQuotas(v []*Quota) *ListQuotasResponseBody {
	s.Quotas = v
	return s
}

func (s *ListQuotasResponseBody) SetRequestId(v string) *ListQuotasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotasResponseBody) SetTotalCount(v int32) *ListQuotasResponseBody {
	s.TotalCount = &v
	return s
}

type ListQuotasResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponse) GoString() string {
	return s.String()
}

func (s *ListQuotasResponse) SetHeaders(v map[string]*string) *ListQuotasResponse {
	s.Headers = v
	return s
}

func (s *ListQuotasResponse) SetStatusCode(v int32) *ListQuotasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotasResponse) SetBody(v *ListQuotasResponseBody) *ListQuotasResponse {
	s.Body = v
	return s
}

type ListResourceGroupMachineGroupsRequest struct {
	// example:
	//
	// 1612285282502326
	CreatorID *string `json:"CreatorID,omitempty" xml:"CreatorID,omitempty"`
	// example:
	//
	// ecs.c6.large
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 236553689400333
	OrderInstanceId *string `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	PaymentDuration *string `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	// example:
	//
	// Month
	PaymentDurationUnit *string `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	// example:
	//
	// PREPAY
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// example:
	//
	// GmtCreatedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// Ready
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListResourceGroupMachineGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupMachineGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupMachineGroupsRequest) SetCreatorID(v string) *ListResourceGroupMachineGroupsRequest {
	s.CreatorID = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetEcsSpec(v string) *ListResourceGroupMachineGroupsRequest {
	s.EcsSpec = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetName(v string) *ListResourceGroupMachineGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetOrder(v string) *ListResourceGroupMachineGroupsRequest {
	s.Order = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetOrderInstanceId(v string) *ListResourceGroupMachineGroupsRequest {
	s.OrderInstanceId = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetPageNumber(v int32) *ListResourceGroupMachineGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetPageSize(v int32) *ListResourceGroupMachineGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetPaymentDuration(v string) *ListResourceGroupMachineGroupsRequest {
	s.PaymentDuration = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetPaymentDurationUnit(v string) *ListResourceGroupMachineGroupsRequest {
	s.PaymentDurationUnit = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetPaymentType(v string) *ListResourceGroupMachineGroupsRequest {
	s.PaymentType = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetSortBy(v string) *ListResourceGroupMachineGroupsRequest {
	s.SortBy = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetStatus(v string) *ListResourceGroupMachineGroupsRequest {
	s.Status = &v
	return s
}

type ListResourceGroupMachineGroupsResponseBody struct {
	MachineGroups []*MachineGroup `json:"MachineGroups,omitempty" xml:"MachineGroups,omitempty" type:"Repeated"`
	// example:
	//
	// E7C42CC7-2E85-508A-84F4-923B605FD10F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 4
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceGroupMachineGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupMachineGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceGroupMachineGroupsResponseBody) SetMachineGroups(v []*MachineGroup) *ListResourceGroupMachineGroupsResponseBody {
	s.MachineGroups = v
	return s
}

func (s *ListResourceGroupMachineGroupsResponseBody) SetRequestId(v string) *ListResourceGroupMachineGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceGroupMachineGroupsResponseBody) SetTotalCount(v string) *ListResourceGroupMachineGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourceGroupMachineGroupsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceGroupMachineGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceGroupMachineGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupMachineGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceGroupMachineGroupsResponse) SetHeaders(v map[string]*string) *ListResourceGroupMachineGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceGroupMachineGroupsResponse) SetStatusCode(v int32) *ListResourceGroupMachineGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceGroupMachineGroupsResponse) SetBody(v *ListResourceGroupMachineGroupsResponseBody) *ListResourceGroupMachineGroupsResponse {
	s.Body = v
	return s
}

type ListResourceGroupsRequest struct {
	// example:
	//
	// Ecs
	ComputingResourceProvider *string `json:"ComputingResourceProvider,omitempty" xml:"ComputingResourceProvider,omitempty"`
	// example:
	//
	// rgf0zhfqn1d4ity2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 2
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// Lingjun
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// false
	ShowAll *bool `json:"ShowAll,omitempty" xml:"ShowAll,omitempty"`
	// example:
	//
	// DisplayName
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListResourceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsRequest) SetComputingResourceProvider(v string) *ListResourceGroupsRequest {
	s.ComputingResourceProvider = &v
	return s
}

func (s *ListResourceGroupsRequest) SetName(v string) *ListResourceGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListResourceGroupsRequest) SetOrder(v string) *ListResourceGroupsRequest {
	s.Order = &v
	return s
}

func (s *ListResourceGroupsRequest) SetPageNumber(v int64) *ListResourceGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceGroupsRequest) SetPageSize(v int64) *ListResourceGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceGroupsRequest) SetResourceType(v string) *ListResourceGroupsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourceGroupsRequest) SetShowAll(v bool) *ListResourceGroupsRequest {
	s.ShowAll = &v
	return s
}

func (s *ListResourceGroupsRequest) SetSortBy(v string) *ListResourceGroupsRequest {
	s.SortBy = &v
	return s
}

func (s *ListResourceGroupsRequest) SetStatus(v string) *ListResourceGroupsRequest {
	s.Status = &v
	return s
}

type ListResourceGroupsResponseBody struct {
	// example:
	//
	// 9CFA2665-1FFE-5929-8468-C14C25890486
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// RG1
	ResourceGroups []*ResourceGroup `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBody) SetRequestId(v string) *ListResourceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetResourceGroups(v []*ResourceGroup) *ListResourceGroupsResponseBody {
	s.ResourceGroups = v
	return s
}

func (s *ListResourceGroupsResponseBody) SetTotalCount(v int64) *ListResourceGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourceGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponse) SetHeaders(v map[string]*string) *ListResourceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceGroupsResponse) SetStatusCode(v int32) *ListResourceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceGroupsResponse) SetBody(v *ListResourceGroupsResponseBody) *ListResourceGroupsResponse {
	s.Body = v
	return s
}

type ListTrainingJobEventsRequest struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListTrainingJobEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobEventsRequest) GoString() string {
	return s.String()
}

func (s *ListTrainingJobEventsRequest) SetEndTime(v string) *ListTrainingJobEventsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTrainingJobEventsRequest) SetPageNumber(v int64) *ListTrainingJobEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrainingJobEventsRequest) SetPageSize(v int64) *ListTrainingJobEventsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTrainingJobEventsRequest) SetStartTime(v string) *ListTrainingJobEventsRequest {
	s.StartTime = &v
	return s
}

type ListTrainingJobEventsResponseBody struct {
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 23
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTrainingJobEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrainingJobEventsResponseBody) SetEvents(v []*string) *ListTrainingJobEventsResponseBody {
	s.Events = v
	return s
}

func (s *ListTrainingJobEventsResponseBody) SetRequestId(v string) *ListTrainingJobEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrainingJobEventsResponseBody) SetTotalCount(v string) *ListTrainingJobEventsResponseBody {
	s.TotalCount = &v
	return s
}

type ListTrainingJobEventsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrainingJobEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrainingJobEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobEventsResponse) GoString() string {
	return s.String()
}

func (s *ListTrainingJobEventsResponse) SetHeaders(v map[string]*string) *ListTrainingJobEventsResponse {
	s.Headers = v
	return s
}

func (s *ListTrainingJobEventsResponse) SetStatusCode(v int32) *ListTrainingJobEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrainingJobEventsResponse) SetBody(v *ListTrainingJobEventsResponseBody) *ListTrainingJobEventsResponse {
	s.Body = v
	return s
}

type ListTrainingJobInstanceEventsRequest struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListTrainingJobInstanceEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobInstanceEventsRequest) GoString() string {
	return s.String()
}

func (s *ListTrainingJobInstanceEventsRequest) SetEndTime(v string) *ListTrainingJobInstanceEventsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTrainingJobInstanceEventsRequest) SetPageNumber(v int64) *ListTrainingJobInstanceEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrainingJobInstanceEventsRequest) SetPageSize(v int64) *ListTrainingJobInstanceEventsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTrainingJobInstanceEventsRequest) SetStartTime(v string) *ListTrainingJobInstanceEventsRequest {
	s.StartTime = &v
	return s
}

type ListTrainingJobInstanceEventsResponseBody struct {
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTrainingJobInstanceEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobInstanceEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrainingJobInstanceEventsResponseBody) SetEvents(v []*string) *ListTrainingJobInstanceEventsResponseBody {
	s.Events = v
	return s
}

func (s *ListTrainingJobInstanceEventsResponseBody) SetRequestId(v string) *ListTrainingJobInstanceEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrainingJobInstanceEventsResponseBody) SetTotalCount(v string) *ListTrainingJobInstanceEventsResponseBody {
	s.TotalCount = &v
	return s
}

type ListTrainingJobInstanceEventsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrainingJobInstanceEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrainingJobInstanceEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobInstanceEventsResponse) GoString() string {
	return s.String()
}

func (s *ListTrainingJobInstanceEventsResponse) SetHeaders(v map[string]*string) *ListTrainingJobInstanceEventsResponse {
	s.Headers = v
	return s
}

func (s *ListTrainingJobInstanceEventsResponse) SetStatusCode(v int32) *ListTrainingJobInstanceEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrainingJobInstanceEventsResponse) SetBody(v *ListTrainingJobInstanceEventsResponseBody) *ListTrainingJobInstanceEventsResponse {
	s.Body = v
	return s
}

type ListTrainingJobInstanceMetricsRequest struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// trains930928remn-master-0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// GpuCoreUsage
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 10s
	TimeStep *string `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
}

func (s ListTrainingJobInstanceMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobInstanceMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListTrainingJobInstanceMetricsRequest) SetEndTime(v string) *ListTrainingJobInstanceMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTrainingJobInstanceMetricsRequest) SetInstanceId(v string) *ListTrainingJobInstanceMetricsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTrainingJobInstanceMetricsRequest) SetMetricType(v string) *ListTrainingJobInstanceMetricsRequest {
	s.MetricType = &v
	return s
}

func (s *ListTrainingJobInstanceMetricsRequest) SetStartTime(v string) *ListTrainingJobInstanceMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *ListTrainingJobInstanceMetricsRequest) SetTimeStep(v string) *ListTrainingJobInstanceMetricsRequest {
	s.TimeStep = &v
	return s
}

type ListTrainingJobInstanceMetricsResponseBody struct {
	InstanceMetrics []*ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics `json:"InstanceMetrics,omitempty" xml:"InstanceMetrics,omitempty" type:"Repeated"`
	// example:
	//
	// F082BD0D-21E1-5F9B-81A0-AB07485B03CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTrainingJobInstanceMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobInstanceMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrainingJobInstanceMetricsResponseBody) SetInstanceMetrics(v []*ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics) *ListTrainingJobInstanceMetricsResponseBody {
	s.InstanceMetrics = v
	return s
}

func (s *ListTrainingJobInstanceMetricsResponseBody) SetRequestId(v string) *ListTrainingJobInstanceMetricsResponseBody {
	s.RequestId = &v
	return s
}

type ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics struct {
	InstanceId *string                                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Metrics    []*ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	NodeName   *string                                                             `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics) GoString() string {
	return s.String()
}

func (s *ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics) SetInstanceId(v string) *ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics {
	s.InstanceId = &v
	return s
}

func (s *ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics) SetMetrics(v []*ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics) *ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics {
	s.Metrics = v
	return s
}

func (s *ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics) SetNodeName(v string) *ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics {
	s.NodeName = &v
	return s
}

type ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics struct {
	Time  *string  `json:"Time,omitempty" xml:"Time,omitempty"`
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics) GoString() string {
	return s.String()
}

func (s *ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics) SetTime(v string) *ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics {
	s.Time = &v
	return s
}

func (s *ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics) SetValue(v float64) *ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics {
	s.Value = &v
	return s
}

type ListTrainingJobInstanceMetricsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrainingJobInstanceMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrainingJobInstanceMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobInstanceMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListTrainingJobInstanceMetricsResponse) SetHeaders(v map[string]*string) *ListTrainingJobInstanceMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListTrainingJobInstanceMetricsResponse) SetStatusCode(v int32) *ListTrainingJobInstanceMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrainingJobInstanceMetricsResponse) SetBody(v *ListTrainingJobInstanceMetricsResponseBody) *ListTrainingJobInstanceMetricsResponse {
	s.Body = v
	return s
}

type ListTrainingJobLogsRequest struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// train129f212o89d-master-0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// train129f212o89d-master-0
	WorkerId *string `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
}

func (s ListTrainingJobLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobLogsRequest) GoString() string {
	return s.String()
}

func (s *ListTrainingJobLogsRequest) SetEndTime(v string) *ListTrainingJobLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTrainingJobLogsRequest) SetInstanceId(v string) *ListTrainingJobLogsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTrainingJobLogsRequest) SetPageNumber(v int64) *ListTrainingJobLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrainingJobLogsRequest) SetPageSize(v int64) *ListTrainingJobLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTrainingJobLogsRequest) SetStartTime(v string) *ListTrainingJobLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListTrainingJobLogsRequest) SetWorkerId(v string) *ListTrainingJobLogsRequest {
	s.WorkerId = &v
	return s
}

type ListTrainingJobLogsResponseBody struct {
	Logs []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// example:
	//
	// CBF05F13-B24C-5129-9048-4FA684DCD579
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 23
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTrainingJobLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrainingJobLogsResponseBody) SetLogs(v []*string) *ListTrainingJobLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListTrainingJobLogsResponseBody) SetRequestId(v string) *ListTrainingJobLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrainingJobLogsResponseBody) SetTotalCount(v string) *ListTrainingJobLogsResponseBody {
	s.TotalCount = &v
	return s
}

type ListTrainingJobLogsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrainingJobLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrainingJobLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobLogsResponse) GoString() string {
	return s.String()
}

func (s *ListTrainingJobLogsResponse) SetHeaders(v map[string]*string) *ListTrainingJobLogsResponse {
	s.Headers = v
	return s
}

func (s *ListTrainingJobLogsResponse) SetStatusCode(v int32) *ListTrainingJobLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrainingJobLogsResponse) SetBody(v *ListTrainingJobLogsResponseBody) *ListTrainingJobLogsResponse {
	s.Body = v
	return s
}

type ListTrainingJobMetricsRequest struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// accuracy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListTrainingJobMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListTrainingJobMetricsRequest) SetEndTime(v string) *ListTrainingJobMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTrainingJobMetricsRequest) SetName(v string) *ListTrainingJobMetricsRequest {
	s.Name = &v
	return s
}

func (s *ListTrainingJobMetricsRequest) SetOrder(v string) *ListTrainingJobMetricsRequest {
	s.Order = &v
	return s
}

func (s *ListTrainingJobMetricsRequest) SetPageNumber(v int64) *ListTrainingJobMetricsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrainingJobMetricsRequest) SetPageSize(v int64) *ListTrainingJobMetricsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTrainingJobMetricsRequest) SetStartTime(v string) *ListTrainingJobMetricsRequest {
	s.StartTime = &v
	return s
}

type ListTrainingJobMetricsResponseBody struct {
	Metrics []*ListTrainingJobMetricsResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTrainingJobMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrainingJobMetricsResponseBody) SetMetrics(v []*ListTrainingJobMetricsResponseBodyMetrics) *ListTrainingJobMetricsResponseBody {
	s.Metrics = v
	return s
}

func (s *ListTrainingJobMetricsResponseBody) SetRequestId(v string) *ListTrainingJobMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrainingJobMetricsResponseBody) SetTotalCount(v int64) *ListTrainingJobMetricsResponseBody {
	s.TotalCount = &v
	return s
}

type ListTrainingJobMetricsResponseBodyMetrics struct {
	// example:
	//
	// accuracy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-04-18T22:20:55Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 0.97
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTrainingJobMetricsResponseBodyMetrics) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobMetricsResponseBodyMetrics) GoString() string {
	return s.String()
}

func (s *ListTrainingJobMetricsResponseBodyMetrics) SetName(v string) *ListTrainingJobMetricsResponseBodyMetrics {
	s.Name = &v
	return s
}

func (s *ListTrainingJobMetricsResponseBodyMetrics) SetTimestamp(v string) *ListTrainingJobMetricsResponseBodyMetrics {
	s.Timestamp = &v
	return s
}

func (s *ListTrainingJobMetricsResponseBodyMetrics) SetValue(v float64) *ListTrainingJobMetricsResponseBodyMetrics {
	s.Value = &v
	return s
}

type ListTrainingJobMetricsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrainingJobMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrainingJobMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListTrainingJobMetricsResponse) SetHeaders(v map[string]*string) *ListTrainingJobMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListTrainingJobMetricsResponse) SetStatusCode(v int32) *ListTrainingJobMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrainingJobMetricsResponse) SetBody(v *ListTrainingJobMetricsResponseBody) *ListTrainingJobMetricsResponse {
	s.Body = v
	return s
}

type ListTrainingJobOutputModelsResponseBody struct {
	OutputModels []*ListTrainingJobOutputModelsResponseBodyOutputModels `json:"OutputModels,omitempty" xml:"OutputModels,omitempty" type:"Repeated"`
}

func (s ListTrainingJobOutputModelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobOutputModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrainingJobOutputModelsResponseBody) SetOutputModels(v []*ListTrainingJobOutputModelsResponseBodyOutputModels) *ListTrainingJobOutputModelsResponseBody {
	s.OutputModels = v
	return s
}

type ListTrainingJobOutputModelsResponseBodyOutputModels struct {
	CompressionSpec map[string]interface{} `json:"CompressionSpec,omitempty" xml:"CompressionSpec,omitempty"`
	// example:
	//
	// {}
	EvaluationSpec map[string]interface{} `json:"EvaluationSpec,omitempty" xml:"EvaluationSpec,omitempty"`
	// example:
	//
	// {}
	InferenceSpec map[string]interface{}                                       `json:"InferenceSpec,omitempty" xml:"InferenceSpec,omitempty"`
	Labels        []*ListTrainingJobOutputModelsResponseBodyOutputModelsLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// example:
	//
	// {
	//
	//       "lr": 0.000001,
	//
	//       "train_loss": 2.6345
	//
	// }
	Metrics map[string]interface{} `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// example:
	//
	// model
	OutputChannelName *string `json:"OutputChannelName,omitempty" xml:"OutputChannelName,omitempty"`
	// example:
	//
	// region=cn-shanghai,workspaceId=1345,kind=PipelineRun,id=run-sakdbaskjdf
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// PAIFlow
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// {}
	TrainingSpec map[string]interface{} `json:"TrainingSpec,omitempty" xml:"TrainingSpec,omitempty"`
	// example:
	//
	// oss://test-bucket.oss-cn-hangzhou.aliyuncs.com/path/to/output/channel/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s ListTrainingJobOutputModelsResponseBodyOutputModels) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobOutputModelsResponseBodyOutputModels) GoString() string {
	return s.String()
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) SetCompressionSpec(v map[string]interface{}) *ListTrainingJobOutputModelsResponseBodyOutputModels {
	s.CompressionSpec = v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) SetEvaluationSpec(v map[string]interface{}) *ListTrainingJobOutputModelsResponseBodyOutputModels {
	s.EvaluationSpec = v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) SetInferenceSpec(v map[string]interface{}) *ListTrainingJobOutputModelsResponseBodyOutputModels {
	s.InferenceSpec = v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) SetLabels(v []*ListTrainingJobOutputModelsResponseBodyOutputModelsLabels) *ListTrainingJobOutputModelsResponseBodyOutputModels {
	s.Labels = v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) SetMetrics(v map[string]interface{}) *ListTrainingJobOutputModelsResponseBodyOutputModels {
	s.Metrics = v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) SetOutputChannelName(v string) *ListTrainingJobOutputModelsResponseBodyOutputModels {
	s.OutputChannelName = &v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) SetSourceId(v string) *ListTrainingJobOutputModelsResponseBodyOutputModels {
	s.SourceId = &v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) SetSourceType(v string) *ListTrainingJobOutputModelsResponseBodyOutputModels {
	s.SourceType = &v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) SetTrainingSpec(v map[string]interface{}) *ListTrainingJobOutputModelsResponseBodyOutputModels {
	s.TrainingSpec = v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) SetUri(v string) *ListTrainingJobOutputModelsResponseBodyOutputModels {
	s.Uri = &v
	return s
}

type ListTrainingJobOutputModelsResponseBodyOutputModelsLabels struct {
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// StableDiffusion
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTrainingJobOutputModelsResponseBodyOutputModelsLabels) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobOutputModelsResponseBodyOutputModelsLabels) GoString() string {
	return s.String()
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModelsLabels) SetKey(v string) *ListTrainingJobOutputModelsResponseBodyOutputModelsLabels {
	s.Key = &v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModelsLabels) SetValue(v string) *ListTrainingJobOutputModelsResponseBodyOutputModelsLabels {
	s.Value = &v
	return s
}

type ListTrainingJobOutputModelsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrainingJobOutputModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrainingJobOutputModelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobOutputModelsResponse) GoString() string {
	return s.String()
}

func (s *ListTrainingJobOutputModelsResponse) SetHeaders(v map[string]*string) *ListTrainingJobOutputModelsResponse {
	s.Headers = v
	return s
}

func (s *ListTrainingJobOutputModelsResponse) SetStatusCode(v int32) *ListTrainingJobOutputModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrainingJobOutputModelsResponse) SetBody(v *ListTrainingJobOutputModelsResponseBody) *ListTrainingJobOutputModelsResponse {
	s.Body = v
	return s
}

type ListTrainingJobsRequest struct {
	// example:
	//
	// llm_train
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// example:
	//
	// pai
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	// example:
	//
	// 2023-12-27T02:10:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// false
	IsTempAlgo *bool `json:"IsTempAlgo,omitempty" xml:"IsTempAlgo,omitempty"`
	// example:
	//
	// {"project": "sd-s3"}
	Labels map[string]interface{} `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtModifiedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// 2024-06-22T01:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// trains930928remn
	TrainingJobId *string `json:"TrainingJobId,omitempty" xml:"TrainingJobId,omitempty"`
	// example:
	//
	// large_language_model_training
	TrainingJobName *string `json:"TrainingJobName,omitempty" xml:"TrainingJobName,omitempty"`
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListTrainingJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsRequest) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsRequest) SetAlgorithmName(v string) *ListTrainingJobsRequest {
	s.AlgorithmName = &v
	return s
}

func (s *ListTrainingJobsRequest) SetAlgorithmProvider(v string) *ListTrainingJobsRequest {
	s.AlgorithmProvider = &v
	return s
}

func (s *ListTrainingJobsRequest) SetEndTime(v string) *ListTrainingJobsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTrainingJobsRequest) SetIsTempAlgo(v bool) *ListTrainingJobsRequest {
	s.IsTempAlgo = &v
	return s
}

func (s *ListTrainingJobsRequest) SetLabels(v map[string]interface{}) *ListTrainingJobsRequest {
	s.Labels = v
	return s
}

func (s *ListTrainingJobsRequest) SetOrder(v string) *ListTrainingJobsRequest {
	s.Order = &v
	return s
}

func (s *ListTrainingJobsRequest) SetPageNumber(v int64) *ListTrainingJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrainingJobsRequest) SetPageSize(v int64) *ListTrainingJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTrainingJobsRequest) SetSortBy(v string) *ListTrainingJobsRequest {
	s.SortBy = &v
	return s
}

func (s *ListTrainingJobsRequest) SetStartTime(v string) *ListTrainingJobsRequest {
	s.StartTime = &v
	return s
}

func (s *ListTrainingJobsRequest) SetStatus(v string) *ListTrainingJobsRequest {
	s.Status = &v
	return s
}

func (s *ListTrainingJobsRequest) SetTrainingJobId(v string) *ListTrainingJobsRequest {
	s.TrainingJobId = &v
	return s
}

func (s *ListTrainingJobsRequest) SetTrainingJobName(v string) *ListTrainingJobsRequest {
	s.TrainingJobName = &v
	return s
}

func (s *ListTrainingJobsRequest) SetWorkspaceId(v string) *ListTrainingJobsRequest {
	s.WorkspaceId = &v
	return s
}

type ListTrainingJobsShrinkRequest struct {
	// example:
	//
	// llm_train
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// example:
	//
	// pai
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	// example:
	//
	// 2023-12-27T02:10:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// false
	IsTempAlgo *bool `json:"IsTempAlgo,omitempty" xml:"IsTempAlgo,omitempty"`
	// example:
	//
	// {"project": "sd-s3"}
	LabelsShrink *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtModifiedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// 2024-06-22T01:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// trains930928remn
	TrainingJobId *string `json:"TrainingJobId,omitempty" xml:"TrainingJobId,omitempty"`
	// example:
	//
	// large_language_model_training
	TrainingJobName *string `json:"TrainingJobName,omitempty" xml:"TrainingJobName,omitempty"`
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListTrainingJobsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsShrinkRequest) SetAlgorithmName(v string) *ListTrainingJobsShrinkRequest {
	s.AlgorithmName = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetAlgorithmProvider(v string) *ListTrainingJobsShrinkRequest {
	s.AlgorithmProvider = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetEndTime(v string) *ListTrainingJobsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetIsTempAlgo(v bool) *ListTrainingJobsShrinkRequest {
	s.IsTempAlgo = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetLabelsShrink(v string) *ListTrainingJobsShrinkRequest {
	s.LabelsShrink = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetOrder(v string) *ListTrainingJobsShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetPageNumber(v int64) *ListTrainingJobsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetPageSize(v int64) *ListTrainingJobsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetSortBy(v string) *ListTrainingJobsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetStartTime(v string) *ListTrainingJobsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetStatus(v string) *ListTrainingJobsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetTrainingJobId(v string) *ListTrainingJobsShrinkRequest {
	s.TrainingJobId = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetTrainingJobName(v string) *ListTrainingJobsShrinkRequest {
	s.TrainingJobName = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetWorkspaceId(v string) *ListTrainingJobsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

type ListTrainingJobsResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount   *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TrainingJobs []*ListTrainingJobsResponseBodyTrainingJobs `json:"TrainingJobs,omitempty" xml:"TrainingJobs,omitempty" type:"Repeated"`
}

func (s ListTrainingJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBody) SetRequestId(v string) *ListTrainingJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrainingJobsResponseBody) SetTotalCount(v int64) *ListTrainingJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTrainingJobsResponseBody) SetTrainingJobs(v []*ListTrainingJobsResponseBodyTrainingJobs) *ListTrainingJobsResponseBody {
	s.TrainingJobs = v
	return s
}

type ListTrainingJobsResponseBodyTrainingJobs struct {
	// example:
	//
	// llm_train
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// example:
	//
	// pai
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	// example:
	//
	// v0.0.1
	AlgorithmVersion *string                                                   `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	ComputeResource  *ListTrainingJobsResponseBodyTrainingJobsComputeResource  `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	DlcJobId         *string                                                   `json:"DlcJobId,omitempty" xml:"DlcJobId,omitempty"`
	Environments     map[string]*string                                        `json:"Environments,omitempty" xml:"Environments,omitempty"`
	ExperimentConfig *ListTrainingJobsResponseBodyTrainingJobsExperimentConfig `json:"ExperimentConfig,omitempty" xml:"ExperimentConfig,omitempty" type:"Struct"`
	// example:
	//
	// 2024-07-10T11:49:47Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2024-07-10T11:49:47Z
	GmtModifiedTime *string                                                    `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	HyperParameters []*ListTrainingJobsResponseBodyTrainingJobsHyperParameters `json:"HyperParameters,omitempty" xml:"HyperParameters,omitempty" type:"Repeated"`
	InputChannels   []*ListTrainingJobsResponseBodyTrainingJobsInputChannels   `json:"InputChannels,omitempty" xml:"InputChannels,omitempty" type:"Repeated"`
	// example:
	//
	// true
	IsTempAlgo         *bool                                                     `json:"IsTempAlgo,omitempty" xml:"IsTempAlgo,omitempty"`
	Labels             []*ListTrainingJobsResponseBodyTrainingJobsLabels         `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	OutputChannels     []*ListTrainingJobsResponseBodyTrainingJobsOutputChannels `json:"OutputChannels,omitempty" xml:"OutputChannels,omitempty" type:"Repeated"`
	PythonRequirements []*string                                                 `json:"PythonRequirements,omitempty" xml:"PythonRequirements,omitempty" type:"Repeated"`
	// example:
	//
	// TrainingJobSucceed
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// example:
	//
	// None
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// example:
	//
	// acs:ram::{accountID}:role/{roleName}
	RoleArn   *string                                            `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	Scheduler *ListTrainingJobsResponseBodyTrainingJobsScheduler `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Struct"`
	// example:
	//
	// Running
	Status                 *string                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusTransitions      []*ListTrainingJobsResponseBodyTrainingJobsStatusTransitions `json:"StatusTransitions,omitempty" xml:"StatusTransitions,omitempty" type:"Repeated"`
	TrainingJobDescription *string                                                      `json:"TrainingJobDescription,omitempty" xml:"TrainingJobDescription,omitempty"`
	// example:
	//
	// train1layo6js8ra
	TrainingJobId *string `json:"TrainingJobId,omitempty" xml:"TrainingJobId,omitempty"`
	// example:
	//
	// qwen2-7b
	TrainingJobName *string `json:"TrainingJobName,omitempty" xml:"TrainingJobName,omitempty"`
	// example:
	//
	// 123456789
	UserId  *string                                          `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserVpc *ListTrainingJobsResponseBodyTrainingJobsUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// example:
	//
	// 1234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobs) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobs) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetAlgorithmName(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.AlgorithmName = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetAlgorithmProvider(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.AlgorithmProvider = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetAlgorithmVersion(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.AlgorithmVersion = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetComputeResource(v *ListTrainingJobsResponseBodyTrainingJobsComputeResource) *ListTrainingJobsResponseBodyTrainingJobs {
	s.ComputeResource = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetDlcJobId(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.DlcJobId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetEnvironments(v map[string]*string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.Environments = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetExperimentConfig(v *ListTrainingJobsResponseBodyTrainingJobsExperimentConfig) *ListTrainingJobsResponseBodyTrainingJobs {
	s.ExperimentConfig = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetGmtCreateTime(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.GmtCreateTime = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetGmtModifiedTime(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetHyperParameters(v []*ListTrainingJobsResponseBodyTrainingJobsHyperParameters) *ListTrainingJobsResponseBodyTrainingJobs {
	s.HyperParameters = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetInputChannels(v []*ListTrainingJobsResponseBodyTrainingJobsInputChannels) *ListTrainingJobsResponseBodyTrainingJobs {
	s.InputChannels = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetIsTempAlgo(v bool) *ListTrainingJobsResponseBodyTrainingJobs {
	s.IsTempAlgo = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetLabels(v []*ListTrainingJobsResponseBodyTrainingJobsLabels) *ListTrainingJobsResponseBodyTrainingJobs {
	s.Labels = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetOutputChannels(v []*ListTrainingJobsResponseBodyTrainingJobsOutputChannels) *ListTrainingJobsResponseBodyTrainingJobs {
	s.OutputChannels = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetPythonRequirements(v []*string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.PythonRequirements = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetReasonCode(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.ReasonCode = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetReasonMessage(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.ReasonMessage = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetRoleArn(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.RoleArn = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetScheduler(v *ListTrainingJobsResponseBodyTrainingJobsScheduler) *ListTrainingJobsResponseBodyTrainingJobs {
	s.Scheduler = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetStatus(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.Status = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetStatusTransitions(v []*ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) *ListTrainingJobsResponseBodyTrainingJobs {
	s.StatusTransitions = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetTrainingJobDescription(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.TrainingJobDescription = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetTrainingJobId(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.TrainingJobId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetTrainingJobName(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.TrainingJobName = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetUserId(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.UserId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetUserVpc(v *ListTrainingJobsResponseBodyTrainingJobsUserVpc) *ListTrainingJobsResponseBodyTrainingJobs {
	s.UserVpc = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetWorkspaceId(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.WorkspaceId = &v
	return s
}

type ListTrainingJobsResponseBodyTrainingJobsComputeResource struct {
	// example:
	//
	// 1
	EcsCount *int64 `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	// example:
	//
	// ecs.gn5-c8g1.2xlarge
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// example:
	//
	// 1
	InstanceCount *int64                                                               `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	InstanceSpec  *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" type:"Struct"`
	// example:
	//
	// quotam670lixikcl
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsComputeResource) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsComputeResource) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResource) SetEcsCount(v int64) *ListTrainingJobsResponseBodyTrainingJobsComputeResource {
	s.EcsCount = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResource) SetEcsSpec(v string) *ListTrainingJobsResponseBodyTrainingJobsComputeResource {
	s.EcsSpec = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResource) SetInstanceCount(v int64) *ListTrainingJobsResponseBodyTrainingJobsComputeResource {
	s.InstanceCount = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResource) SetInstanceSpec(v *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) *ListTrainingJobsResponseBodyTrainingJobsComputeResource {
	s.InstanceSpec = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResource) SetResourceId(v string) *ListTrainingJobsResponseBodyTrainingJobsComputeResource {
	s.ResourceId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResource) SetResourceName(v string) *ListTrainingJobsResponseBodyTrainingJobsComputeResource {
	s.ResourceName = &v
	return s
}

type ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec struct {
	// example:
	//
	// 8
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// 1
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// example:
	//
	// V100
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// example:
	//
	// 32
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// 32
	SharedMemory *string `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) SetCPU(v string) *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec {
	s.CPU = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) SetGPU(v string) *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec {
	s.GPU = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) SetGPUType(v string) *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec {
	s.GPUType = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) SetMemory(v string) *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec {
	s.Memory = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) SetSharedMemory(v string) *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec {
	s.SharedMemory = &v
	return s
}

type ListTrainingJobsResponseBodyTrainingJobsExperimentConfig struct {
	// example:
	//
	// exp-ds9aefia90v
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// large_language_model
	ExperimentName *string `json:"ExperimentName,omitempty" xml:"ExperimentName,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsExperimentConfig) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsExperimentConfig) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobsExperimentConfig) SetExperimentId(v string) *ListTrainingJobsResponseBodyTrainingJobsExperimentConfig {
	s.ExperimentId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsExperimentConfig) SetExperimentName(v string) *ListTrainingJobsResponseBodyTrainingJobsExperimentConfig {
	s.ExperimentName = &v
	return s
}

type ListTrainingJobsResponseBodyTrainingJobsHyperParameters struct {
	// example:
	//
	// learning_rate
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.001
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsHyperParameters) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsHyperParameters) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobsHyperParameters) SetName(v string) *ListTrainingJobsResponseBodyTrainingJobsHyperParameters {
	s.Name = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsHyperParameters) SetValue(v string) *ListTrainingJobsResponseBodyTrainingJobsHyperParameters {
	s.Value = &v
	return s
}

type ListTrainingJobsResponseBodyTrainingJobsInputChannels struct {
	// example:
	//
	// d-475megosidivjfgfq6
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// oss://test-bucket.oss-cn-hangzhou-internal.aliyuncs.com/path/to/input/channel/
	InputUri *string `json:"InputUri,omitempty" xml:"InputUri,omitempty"`
	// example:
	//
	// model
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsInputChannels) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsInputChannels) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobsInputChannels) SetDatasetId(v string) *ListTrainingJobsResponseBodyTrainingJobsInputChannels {
	s.DatasetId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsInputChannels) SetInputUri(v string) *ListTrainingJobsResponseBodyTrainingJobsInputChannels {
	s.InputUri = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsInputChannels) SetName(v string) *ListTrainingJobsResponseBodyTrainingJobsInputChannels {
	s.Name = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsInputChannels) SetVersionName(v string) *ListTrainingJobsResponseBodyTrainingJobsInputChannels {
	s.VersionName = &v
	return s
}

type ListTrainingJobsResponseBodyTrainingJobsLabels struct {
	// example:
	//
	// CreatedBy
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// QuickStart
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsLabels) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsLabels) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobsLabels) SetKey(v string) *ListTrainingJobsResponseBodyTrainingJobsLabels {
	s.Key = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsLabels) SetValue(v string) *ListTrainingJobsResponseBodyTrainingJobsLabels {
	s.Value = &v
	return s
}

type ListTrainingJobsResponseBodyTrainingJobsOutputChannels struct {
	// example:
	//
	// d-8o0hh35po15ejcdq2p
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// model
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// oss://test-bucket.oss-cn-hangzhou-internal.aliyuncs.com/path/to/output/channel/
	OutputUri   *string `json:"OutputUri,omitempty" xml:"OutputUri,omitempty"`
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsOutputChannels) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsOutputChannels) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobsOutputChannels) SetDatasetId(v string) *ListTrainingJobsResponseBodyTrainingJobsOutputChannels {
	s.DatasetId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsOutputChannels) SetName(v string) *ListTrainingJobsResponseBodyTrainingJobsOutputChannels {
	s.Name = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsOutputChannels) SetOutputUri(v string) *ListTrainingJobsResponseBodyTrainingJobsOutputChannels {
	s.OutputUri = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsOutputChannels) SetVersionName(v string) *ListTrainingJobsResponseBodyTrainingJobsOutputChannels {
	s.VersionName = &v
	return s
}

type ListTrainingJobsResponseBodyTrainingJobsScheduler struct {
	// example:
	//
	// 0
	MaxRunningTimeInSeconds *int64 `json:"MaxRunningTimeInSeconds,omitempty" xml:"MaxRunningTimeInSeconds,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsScheduler) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsScheduler) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobsScheduler) SetMaxRunningTimeInSeconds(v int64) *ListTrainingJobsResponseBodyTrainingJobsScheduler {
	s.MaxRunningTimeInSeconds = &v
	return s
}

type ListTrainingJobsResponseBodyTrainingJobsStatusTransitions struct {
	// example:
	//
	// 2024-07-10T11:49:47Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// TrainingJobSucceed
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// example:
	//
	// KubeDL job runs successfully
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// example:
	//
	// 2024-07-10T11:49:47Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) SetEndTime(v string) *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions {
	s.EndTime = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) SetReasonCode(v string) *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions {
	s.ReasonCode = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) SetReasonMessage(v string) *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions {
	s.ReasonMessage = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) SetStartTime(v string) *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions {
	s.StartTime = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) SetStatus(v string) *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions {
	s.Status = &v
	return s
}

type ListTrainingJobsResponseBodyTrainingJobsUserVpc struct {
	// example:
	//
	// eth1
	DefaultRoute  *string   `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	ExtendedCIDRs []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	// example:
	//
	// sg-abcdef****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// vs-abcdef****
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-abcdef****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsUserVpc) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsUserVpc) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobsUserVpc) SetDefaultRoute(v string) *ListTrainingJobsResponseBodyTrainingJobsUserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsUserVpc) SetExtendedCIDRs(v []*string) *ListTrainingJobsResponseBodyTrainingJobsUserVpc {
	s.ExtendedCIDRs = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsUserVpc) SetSecurityGroupId(v string) *ListTrainingJobsResponseBodyTrainingJobsUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsUserVpc) SetSwitchId(v string) *ListTrainingJobsResponseBodyTrainingJobsUserVpc {
	s.SwitchId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsUserVpc) SetVpcId(v string) *ListTrainingJobsResponseBodyTrainingJobsUserVpc {
	s.VpcId = &v
	return s
}

type ListTrainingJobsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrainingJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrainingJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsResponse) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponse) SetHeaders(v map[string]*string) *ListTrainingJobsResponse {
	s.Headers = v
	return s
}

func (s *ListTrainingJobsResponse) SetStatusCode(v int32) *ListTrainingJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrainingJobsResponse) SetBody(v *ListTrainingJobsResponseBody) *ListTrainingJobsResponse {
	s.Body = v
	return s
}

type ScaleQuotaRequest struct {
	Min              *ResourceSpec `json:"Min,omitempty" xml:"Min,omitempty"`
	ResourceGroupIds []*string     `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
}

func (s ScaleQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s ScaleQuotaRequest) GoString() string {
	return s.String()
}

func (s *ScaleQuotaRequest) SetMin(v *ResourceSpec) *ScaleQuotaRequest {
	s.Min = v
	return s
}

func (s *ScaleQuotaRequest) SetResourceGroupIds(v []*string) *ScaleQuotaRequest {
	s.ResourceGroupIds = v
	return s
}

type ScaleQuotaResponseBody struct {
	// Quota Id
	//
	// example:
	//
	// quotamtl37ge7gkvdz
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// example:
	//
	// F2D0392B-D749-5C48-A98A-3FAE5C9444A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScaleQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleQuotaResponseBody) SetQuotaId(v string) *ScaleQuotaResponseBody {
	s.QuotaId = &v
	return s
}

func (s *ScaleQuotaResponseBody) SetRequestId(v string) *ScaleQuotaResponseBody {
	s.RequestId = &v
	return s
}

type ScaleQuotaResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScaleQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScaleQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s ScaleQuotaResponse) GoString() string {
	return s.String()
}

func (s *ScaleQuotaResponse) SetHeaders(v map[string]*string) *ScaleQuotaResponse {
	s.Headers = v
	return s
}

func (s *ScaleQuotaResponse) SetStatusCode(v int32) *ScaleQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *ScaleQuotaResponse) SetBody(v *ScaleQuotaResponseBody) *ScaleQuotaResponse {
	s.Body = v
	return s
}

type StopTrainingJobResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopTrainingJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopTrainingJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopTrainingJobResponseBody) SetRequestId(v string) *StopTrainingJobResponseBody {
	s.RequestId = &v
	return s
}

type StopTrainingJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopTrainingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTrainingJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StopTrainingJobResponse) GoString() string {
	return s.String()
}

func (s *StopTrainingJobResponse) SetHeaders(v map[string]*string) *StopTrainingJobResponse {
	s.Headers = v
	return s
}

func (s *StopTrainingJobResponse) SetStatusCode(v int32) *StopTrainingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTrainingJobResponse) SetBody(v *StopTrainingJobResponseBody) *StopTrainingJobResponse {
	s.Body = v
	return s
}

type UpdateAlgorithmRequest struct {
	AlgorithmDescription *string `json:"AlgorithmDescription,omitempty" xml:"AlgorithmDescription,omitempty"`
	// example:
	//
	// LLM Train
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
}

func (s UpdateAlgorithmRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlgorithmRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlgorithmRequest) SetAlgorithmDescription(v string) *UpdateAlgorithmRequest {
	s.AlgorithmDescription = &v
	return s
}

func (s *UpdateAlgorithmRequest) SetDisplayName(v string) *UpdateAlgorithmRequest {
	s.DisplayName = &v
	return s
}

type UpdateAlgorithmResponseBody struct {
	// example:
	//
	// algo-sidjc8134hv
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	// example:
	//
	// E7C42CC7-2E85-508A-84F4-923B605FD10F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAlgorithmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlgorithmResponseBody) SetAlgorithmId(v string) *UpdateAlgorithmResponseBody {
	s.AlgorithmId = &v
	return s
}

func (s *UpdateAlgorithmResponseBody) SetRequestId(v string) *UpdateAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAlgorithmResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlgorithmResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlgorithmResponse) SetHeaders(v map[string]*string) *UpdateAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlgorithmResponse) SetStatusCode(v int32) *UpdateAlgorithmResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlgorithmResponse) SetBody(v *UpdateAlgorithmResponseBody) *UpdateAlgorithmResponse {
	s.Body = v
	return s
}

type UpdateAlgorithmVersionRequest struct {
	AlgorithmSpec *AlgorithmSpec `json:"AlgorithmSpec,omitempty" xml:"AlgorithmSpec,omitempty"`
}

func (s UpdateAlgorithmVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlgorithmVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlgorithmVersionRequest) SetAlgorithmSpec(v *AlgorithmSpec) *UpdateAlgorithmVersionRequest {
	s.AlgorithmSpec = v
	return s
}

type UpdateAlgorithmVersionShrinkRequest struct {
	AlgorithmSpecShrink *string `json:"AlgorithmSpec,omitempty" xml:"AlgorithmSpec,omitempty"`
}

func (s UpdateAlgorithmVersionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlgorithmVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlgorithmVersionShrinkRequest) SetAlgorithmSpecShrink(v string) *UpdateAlgorithmVersionShrinkRequest {
	s.AlgorithmSpecShrink = &v
	return s
}

type UpdateAlgorithmVersionResponseBody struct {
	// example:
	//
	// algo-sidjc8134hv
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	// example:
	//
	// v0.1.0
	AlgorithmVersion *string `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
}

func (s UpdateAlgorithmVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlgorithmVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlgorithmVersionResponseBody) SetAlgorithmId(v string) *UpdateAlgorithmVersionResponseBody {
	s.AlgorithmId = &v
	return s
}

func (s *UpdateAlgorithmVersionResponseBody) SetAlgorithmVersion(v string) *UpdateAlgorithmVersionResponseBody {
	s.AlgorithmVersion = &v
	return s
}

type UpdateAlgorithmVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlgorithmVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlgorithmVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlgorithmVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlgorithmVersionResponse) SetHeaders(v map[string]*string) *UpdateAlgorithmVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlgorithmVersionResponse) SetStatusCode(v int32) *UpdateAlgorithmVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlgorithmVersionResponse) SetBody(v *UpdateAlgorithmVersionResponseBody) *UpdateAlgorithmVersionResponse {
	s.Body = v
	return s
}

type UpdateQuotaRequest struct {
	// example:
	//
	// this is a test quota
	Description   *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Labels        []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	QueueStrategy *string  `json:"QueueStrategy,omitempty" xml:"QueueStrategy,omitempty"`
	// if can be null:
	// true
	QuotaConfig *QuotaConfig `json:"QuotaConfig,omitempty" xml:"QuotaConfig,omitempty"`
	QuotaName   *string      `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
}

func (s UpdateQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaRequest) GoString() string {
	return s.String()
}

func (s *UpdateQuotaRequest) SetDescription(v string) *UpdateQuotaRequest {
	s.Description = &v
	return s
}

func (s *UpdateQuotaRequest) SetLabels(v []*Label) *UpdateQuotaRequest {
	s.Labels = v
	return s
}

func (s *UpdateQuotaRequest) SetQueueStrategy(v string) *UpdateQuotaRequest {
	s.QueueStrategy = &v
	return s
}

func (s *UpdateQuotaRequest) SetQuotaConfig(v *QuotaConfig) *UpdateQuotaRequest {
	s.QuotaConfig = v
	return s
}

func (s *UpdateQuotaRequest) SetQuotaName(v string) *UpdateQuotaRequest {
	s.QuotaName = &v
	return s
}

type UpdateQuotaResponseBody struct {
	// Quota Id
	//
	// example:
	//
	// quota-20210126170216-mtl37ge7gkvdz
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// example:
	//
	// 96496E6E-00B4-5F55-80F6-1844FA9E92DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQuotaResponseBody) SetQuotaId(v string) *UpdateQuotaResponseBody {
	s.QuotaId = &v
	return s
}

func (s *UpdateQuotaResponseBody) SetRequestId(v string) *UpdateQuotaResponseBody {
	s.RequestId = &v
	return s
}

type UpdateQuotaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaResponse) GoString() string {
	return s.String()
}

func (s *UpdateQuotaResponse) SetHeaders(v map[string]*string) *UpdateQuotaResponse {
	s.Headers = v
	return s
}

func (s *UpdateQuotaResponse) SetStatusCode(v int32) *UpdateQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQuotaResponse) SetBody(v *UpdateQuotaResponseBody) *UpdateQuotaResponse {
	s.Body = v
	return s
}

type UpdateResourceGroupRequest struct {
	// example:
	//
	// test_new_havpn_tf
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// prophet
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// true
	Unbind  *bool    `json:"Unbind,omitempty" xml:"Unbind,omitempty"`
	UserVpc *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
}

func (s UpdateResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupRequest) SetDescription(v string) *UpdateResourceGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateResourceGroupRequest) SetName(v string) *UpdateResourceGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateResourceGroupRequest) SetUnbind(v bool) *UpdateResourceGroupRequest {
	s.Unbind = &v
	return s
}

func (s *UpdateResourceGroupRequest) SetUserVpc(v *UserVpc) *UpdateResourceGroupRequest {
	s.UserVpc = v
	return s
}

type UpdateResourceGroupResponseBody struct {
	// example:
	//
	// rgf0zhfqn1d4ity2
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	// example:
	//
	// FFB1D4B4-B253-540A-9B3B-AA711C48A1B7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupResponseBody) SetResourceGroupID(v string) *UpdateResourceGroupResponseBody {
	s.ResourceGroupID = &v
	return s
}

func (s *UpdateResourceGroupResponseBody) SetRequestId(v string) *UpdateResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupResponse) SetHeaders(v map[string]*string) *UpdateResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceGroupResponse) SetStatusCode(v int32) *UpdateResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceGroupResponse) SetBody(v *UpdateResourceGroupResponseBody) *UpdateResourceGroupResponse {
	s.Body = v
	return s
}

type UpdateTrainingJobLabelsRequest struct {
	Labels []*UpdateTrainingJobLabelsRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s UpdateTrainingJobLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrainingJobLabelsRequest) GoString() string {
	return s.String()
}

func (s *UpdateTrainingJobLabelsRequest) SetLabels(v []*UpdateTrainingJobLabelsRequestLabels) *UpdateTrainingJobLabelsRequest {
	s.Labels = v
	return s
}

type UpdateTrainingJobLabelsRequestLabels struct {
	// example:
	//
	// RootModelID
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// model-ad8cv770kl
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateTrainingJobLabelsRequestLabels) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrainingJobLabelsRequestLabels) GoString() string {
	return s.String()
}

func (s *UpdateTrainingJobLabelsRequestLabels) SetKey(v string) *UpdateTrainingJobLabelsRequestLabels {
	s.Key = &v
	return s
}

func (s *UpdateTrainingJobLabelsRequestLabels) SetValue(v string) *UpdateTrainingJobLabelsRequestLabels {
	s.Value = &v
	return s
}

type UpdateTrainingJobLabelsResponseBody struct {
	// example:
	//
	// F082BD0D-21E1-5F9B-81A0-AB07485B03CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTrainingJobLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrainingJobLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTrainingJobLabelsResponseBody) SetRequestId(v string) *UpdateTrainingJobLabelsResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTrainingJobLabelsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTrainingJobLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTrainingJobLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrainingJobLabelsResponse) GoString() string {
	return s.String()
}

func (s *UpdateTrainingJobLabelsResponse) SetHeaders(v map[string]*string) *UpdateTrainingJobLabelsResponse {
	s.Headers = v
	return s
}

func (s *UpdateTrainingJobLabelsResponse) SetStatusCode(v int32) *UpdateTrainingJobLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTrainingJobLabelsResponse) SetBody(v *UpdateTrainingJobLabelsResponseBody) *UpdateTrainingJobLabelsResponse {
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
	client.ProductId = tea.String("PaiStudio")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-beijing":     tea.String("pai.cn-beijing.aliyuncs.com"),
		"cn-hangzhou":    tea.String("pai.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":    tea.String("pai.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":    tea.String("pai.cn-shenzhen.aliyuncs.com"),
		"cn-hongkong":    tea.String("pai.cn-hongkong.aliyuncs.com"),
		"ap-southeast-1": tea.String("pai.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2": tea.String("pai.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3": tea.String("pai.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5": tea.String("pai.ap-southeast-5.aliyuncs.com"),
		"us-west-1":      tea.String("pai.us-west-1.aliyuncs.com"),
		"us-east-1":      tea.String("pai.us-east-1.aliyuncs.com"),
		"eu-central-1":   tea.String("pai.eu-central-1.aliyuncs.com"),
		"me-east-1":      tea.String("pai.me-east-1.aliyuncs.com"),
		"ap-south-1":     tea.String("pai.ap-south-1.aliyuncs.com"),
		"cn-qingdao":     tea.String("pai.cn-qingdao.aliyuncs.com"),
		"cn-zhangjiakou": tea.String("pai.cn-zhangjiakou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("paistudio"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// 检查WebTerminal
//
// @param request - CheckInstanceWebTerminalRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckInstanceWebTerminalResponse
func (client *Client) CheckInstanceWebTerminalWithOptions(TrainingJobId *string, InstanceId *string, request *CheckInstanceWebTerminalRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckInstanceWebTerminalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckInfo)) {
		body["CheckInfo"] = request.CheckInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckInstanceWebTerminal"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trainingjobs/" + tea.StringValue(openapiutil.GetEncodeParam(TrainingJobId)) + "/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/webterminals/action/check"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CheckInstanceWebTerminalResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CheckInstanceWebTerminalResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 检查WebTerminal
//
// @param request - CheckInstanceWebTerminalRequest
//
// @return CheckInstanceWebTerminalResponse
func (client *Client) CheckInstanceWebTerminal(TrainingJobId *string, InstanceId *string, request *CheckInstanceWebTerminalRequest) (_result *CheckInstanceWebTerminalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckInstanceWebTerminalResponse{}
	_body, _err := client.CheckInstanceWebTerminalWithOptions(TrainingJobId, InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建新的算法
//
// @param request - CreateAlgorithmRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlgorithmResponse
func (client *Client) CreateAlgorithmWithOptions(request *CreateAlgorithmRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAlgorithmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlgorithmDescription)) {
		body["AlgorithmDescription"] = request.AlgorithmDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AlgorithmName)) {
		body["AlgorithmName"] = request.AlgorithmName
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAlgorithm"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/algorithms"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateAlgorithmResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateAlgorithmResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建新的算法
//
// @param request - CreateAlgorithmRequest
//
// @return CreateAlgorithmResponse
func (client *Client) CreateAlgorithm(request *CreateAlgorithmRequest) (_result *CreateAlgorithmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAlgorithmResponse{}
	_body, _err := client.CreateAlgorithmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建一个新的算法版本
//
// @param tmpReq - CreateAlgorithmVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlgorithmVersionResponse
func (client *Client) CreateAlgorithmVersionWithOptions(AlgorithmId *string, AlgorithmVersion *string, tmpReq *CreateAlgorithmVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAlgorithmVersionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAlgorithmVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AlgorithmSpec)) {
		request.AlgorithmSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlgorithmSpec, tea.String("AlgorithmSpec"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlgorithmSpecShrink)) {
		body["AlgorithmSpec"] = request.AlgorithmSpecShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAlgorithmVersion"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/algorithms/" + tea.StringValue(openapiutil.GetEncodeParam(AlgorithmId)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(AlgorithmVersion))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateAlgorithmVersionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateAlgorithmVersionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建一个新的算法版本
//
// @param request - CreateAlgorithmVersionRequest
//
// @return CreateAlgorithmVersionResponse
func (client *Client) CreateAlgorithmVersion(AlgorithmId *string, AlgorithmVersion *string, request *CreateAlgorithmVersionRequest) (_result *CreateAlgorithmVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAlgorithmVersionResponse{}
	_body, _err := client.CreateAlgorithmVersionWithOptions(AlgorithmId, AlgorithmVersion, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建WebTerminal
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceWebTerminalResponse
func (client *Client) CreateInstanceWebTerminalWithOptions(TrainingJobId *string, InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceWebTerminalResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceWebTerminal"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trainingjobs/" + tea.StringValue(openapiutil.GetEncodeParam(TrainingJobId)) + "/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/webterminals"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateInstanceWebTerminalResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateInstanceWebTerminalResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建WebTerminal
//
// @return CreateInstanceWebTerminalResponse
func (client *Client) CreateInstanceWebTerminal(TrainingJobId *string, InstanceId *string) (_result *CreateInstanceWebTerminalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceWebTerminalResponse{}
	_body, _err := client.CreateInstanceWebTerminalWithOptions(TrainingJobId, InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Quota
//
// @param request - CreateQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQuotaResponse
func (client *Client) CreateQuotaWithOptions(request *CreateQuotaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllocateStrategy)) {
		body["AllocateStrategy"] = request.AllocateStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Min)) {
		body["Min"] = request.Min
	}

	if !tea.BoolValue(util.IsUnset(request.ParentQuotaId)) {
		body["ParentQuotaId"] = request.ParentQuotaId
	}

	if !tea.BoolValue(util.IsUnset(request.QueueStrategy)) {
		body["QueueStrategy"] = request.QueueStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaConfig)) {
		body["QuotaConfig"] = request.QuotaConfig
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaName)) {
		body["QuotaName"] = request.QuotaName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupIds)) {
		body["ResourceGroupIds"] = request.ResourceGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateQuota"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateQuotaResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateQuotaResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建Quota
//
// @param request - CreateQuotaRequest
//
// @return CreateQuotaResponse
func (client *Client) CreateQuota(request *CreateQuotaRequest) (_result *CreateQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateQuotaResponse{}
	_body, _err := client.CreateQuotaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建资源组
//
// @param request - CreateResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceGroupResponse
func (client *Client) CreateResourceGroupWithOptions(request *CreateResourceGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComputingResourceProvider)) {
		body["ComputingResourceProvider"] = request.ComputingResourceProvider
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.UserVpc)) {
		body["UserVpc"] = request.UserVpc
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResourceGroup"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateResourceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateResourceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建资源组
//
// @param request - CreateResourceGroupRequest
//
// @return CreateResourceGroupResponse
func (client *Client) CreateResourceGroup(request *CreateResourceGroupRequest) (_result *CreateResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResourceGroupResponse{}
	_body, _err := client.CreateResourceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建TrainingJob
//
// @param request - CreateTrainingJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrainingJobResponse
func (client *Client) CreateTrainingJobWithOptions(request *CreateTrainingJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTrainingJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlgorithmName)) {
		body["AlgorithmName"] = request.AlgorithmName
	}

	if !tea.BoolValue(util.IsUnset(request.AlgorithmProvider)) {
		body["AlgorithmProvider"] = request.AlgorithmProvider
	}

	if !tea.BoolValue(util.IsUnset(request.AlgorithmSpec)) {
		body["AlgorithmSpec"] = request.AlgorithmSpec
	}

	if !tea.BoolValue(util.IsUnset(request.AlgorithmVersion)) {
		body["AlgorithmVersion"] = request.AlgorithmVersion
	}

	if !tea.BoolValue(util.IsUnset(request.CodeDir)) {
		body["CodeDir"] = request.CodeDir
	}

	if !tea.BoolValue(util.IsUnset(request.ComputeResource)) {
		body["ComputeResource"] = request.ComputeResource
	}

	if !tea.BoolValue(util.IsUnset(request.Environments)) {
		body["Environments"] = request.Environments
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentConfig)) {
		body["ExperimentConfig"] = request.ExperimentConfig
	}

	if !tea.BoolValue(util.IsUnset(request.HyperParameters)) {
		body["HyperParameters"] = request.HyperParameters
	}

	if !tea.BoolValue(util.IsUnset(request.InputChannels)) {
		body["InputChannels"] = request.InputChannels
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.OutputChannels)) {
		body["OutputChannels"] = request.OutputChannels
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.PythonRequirements)) {
		body["PythonRequirements"] = request.PythonRequirements
	}

	if !tea.BoolValue(util.IsUnset(request.RoleArn)) {
		body["RoleArn"] = request.RoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.Scheduler)) {
		body["Scheduler"] = request.Scheduler
	}

	if !tea.BoolValue(util.IsUnset(request.Settings)) {
		body["Settings"] = request.Settings
	}

	if !tea.BoolValue(util.IsUnset(request.TrainingJobDescription)) {
		body["TrainingJobDescription"] = request.TrainingJobDescription
	}

	if !tea.BoolValue(util.IsUnset(request.TrainingJobName)) {
		body["TrainingJobName"] = request.TrainingJobName
	}

	if !tea.BoolValue(util.IsUnset(request.UserVpc)) {
		body["UserVpc"] = request.UserVpc
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTrainingJob"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trainingjobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateTrainingJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateTrainingJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建TrainingJob
//
// @param request - CreateTrainingJobRequest
//
// @return CreateTrainingJobResponse
func (client *Client) CreateTrainingJob(request *CreateTrainingJobRequest) (_result *CreateTrainingJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTrainingJobResponse{}
	_body, _err := client.CreateTrainingJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除算法
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlgorithmResponse
func (client *Client) DeleteAlgorithmWithOptions(AlgorithmId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAlgorithmResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAlgorithm"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/algorithms/" + tea.StringValue(openapiutil.GetEncodeParam(AlgorithmId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteAlgorithmResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteAlgorithmResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除算法
//
// @return DeleteAlgorithmResponse
func (client *Client) DeleteAlgorithm(AlgorithmId *string) (_result *DeleteAlgorithmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAlgorithmResponse{}
	_body, _err := client.DeleteAlgorithmWithOptions(AlgorithmId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除算法版本
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlgorithmVersionResponse
func (client *Client) DeleteAlgorithmVersionWithOptions(AlgorithmId *string, AlgorithmVersion *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAlgorithmVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAlgorithmVersion"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/algorithms/" + tea.StringValue(openapiutil.GetEncodeParam(AlgorithmId)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(AlgorithmVersion))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteAlgorithmVersionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteAlgorithmVersionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除算法版本
//
// @return DeleteAlgorithmVersionResponse
func (client *Client) DeleteAlgorithmVersion(AlgorithmId *string, AlgorithmVersion *string) (_result *DeleteAlgorithmVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAlgorithmVersionResponse{}
	_body, _err := client.DeleteAlgorithmVersionWithOptions(AlgorithmId, AlgorithmVersion, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteMachineGroup is deprecated
//
// Summary:
//
// delete machine group
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMachineGroupResponse
// Deprecated
func (client *Client) DeleteMachineGroupWithOptions(MachineGroupID *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteMachineGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMachineGroup"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resources/machinegroups/" + tea.StringValue(openapiutil.GetEncodeParam(MachineGroupID))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteMachineGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteMachineGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Deprecated: OpenAPI DeleteMachineGroup is deprecated
//
// Summary:
//
// delete machine group
//
// @return DeleteMachineGroupResponse
// Deprecated
func (client *Client) DeleteMachineGroup(MachineGroupID *string) (_result *DeleteMachineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMachineGroupResponse{}
	_body, _err := client.DeleteMachineGroupWithOptions(MachineGroupID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除Quota
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQuotaResponse
func (client *Client) DeleteQuotaWithOptions(QuotaId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteQuotaResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteQuota"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(QuotaId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteQuotaResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteQuotaResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除Quota
//
// @return DeleteQuotaResponse
func (client *Client) DeleteQuota(QuotaId *string) (_result *DeleteQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteQuotaResponse{}
	_body, _err := client.DeleteQuotaWithOptions(QuotaId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除资源组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceGroupResponse
func (client *Client) DeleteResourceGroupWithOptions(ResourceGroupID *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteResourceGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResourceGroup"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceGroupID))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteResourceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteResourceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除资源组
//
// @return DeleteResourceGroupResponse
func (client *Client) DeleteResourceGroup(ResourceGroupID *string) (_result *DeleteResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceGroupResponse{}
	_body, _err := client.DeleteResourceGroupWithOptions(ResourceGroupID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteResourceGroupMachineGroup is deprecated
//
// Summary:
//
// delete machine group
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceGroupMachineGroupResponse
// Deprecated
func (client *Client) DeleteResourceGroupMachineGroupWithOptions(MachineGroupID *string, ResourceGroupID *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteResourceGroupMachineGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResourceGroupMachineGroup"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceGroupID)) + "/machinegroups/" + tea.StringValue(openapiutil.GetEncodeParam(MachineGroupID))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteResourceGroupMachineGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteResourceGroupMachineGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Deprecated: OpenAPI DeleteResourceGroupMachineGroup is deprecated
//
// Summary:
//
// delete machine group
//
// @return DeleteResourceGroupMachineGroupResponse
// Deprecated
func (client *Client) DeleteResourceGroupMachineGroup(MachineGroupID *string, ResourceGroupID *string) (_result *DeleteResourceGroupMachineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceGroupMachineGroupResponse{}
	_body, _err := client.DeleteResourceGroupMachineGroupWithOptions(MachineGroupID, ResourceGroupID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除一个TrainingJob
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrainingJobResponse
func (client *Client) DeleteTrainingJobWithOptions(TrainingJobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTrainingJobResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTrainingJob"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trainingjobs/" + tea.StringValue(openapiutil.GetEncodeParam(TrainingJobId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteTrainingJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteTrainingJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除一个TrainingJob
//
// @return DeleteTrainingJobResponse
func (client *Client) DeleteTrainingJob(TrainingJobId *string) (_result *DeleteTrainingJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTrainingJobResponse{}
	_body, _err := client.DeleteTrainingJobWithOptions(TrainingJobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除TrainingJob的Labels
//
// @param request - DeleteTrainingJobLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrainingJobLabelsResponse
func (client *Client) DeleteTrainingJobLabelsWithOptions(TrainingJobId *string, request *DeleteTrainingJobLabelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTrainingJobLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keys)) {
		query["Keys"] = request.Keys
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTrainingJobLabels"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trainingjobs/" + tea.StringValue(openapiutil.GetEncodeParam(TrainingJobId)) + "/labels"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteTrainingJobLabelsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteTrainingJobLabelsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除TrainingJob的Labels
//
// @param request - DeleteTrainingJobLabelsRequest
//
// @return DeleteTrainingJobLabelsResponse
func (client *Client) DeleteTrainingJobLabels(TrainingJobId *string, request *DeleteTrainingJobLabelsRequest) (_result *DeleteTrainingJobLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTrainingJobLabelsResponse{}
	_body, _err := client.DeleteTrainingJobLabelsWithOptions(TrainingJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取一个算法信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlgorithmResponse
func (client *Client) GetAlgorithmWithOptions(AlgorithmId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAlgorithmResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAlgorithm"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/algorithms/" + tea.StringValue(openapiutil.GetEncodeParam(AlgorithmId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetAlgorithmResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetAlgorithmResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取一个算法信息
//
// @return GetAlgorithmResponse
func (client *Client) GetAlgorithm(AlgorithmId *string) (_result *GetAlgorithmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAlgorithmResponse{}
	_body, _err := client.GetAlgorithmWithOptions(AlgorithmId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建一个新的算法版本
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlgorithmVersionResponse
func (client *Client) GetAlgorithmVersionWithOptions(AlgorithmId *string, AlgorithmVersion *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAlgorithmVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAlgorithmVersion"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/algorithms/" + tea.StringValue(openapiutil.GetEncodeParam(AlgorithmId)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(AlgorithmVersion))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetAlgorithmVersionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetAlgorithmVersionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建一个新的算法版本
//
// @return GetAlgorithmVersionResponse
func (client *Client) GetAlgorithmVersion(AlgorithmId *string, AlgorithmVersion *string) (_result *GetAlgorithmVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAlgorithmVersionResponse{}
	_body, _err := client.GetAlgorithmVersionWithOptions(AlgorithmId, AlgorithmVersion, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetMachineGroup is deprecated
//
// Summary:
//
// get machine group
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMachineGroupResponse
// Deprecated
func (client *Client) GetMachineGroupWithOptions(MachineGroupID *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMachineGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetMachineGroup"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resources/machinegroups/" + tea.StringValue(openapiutil.GetEncodeParam(MachineGroupID))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetMachineGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetMachineGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Deprecated: OpenAPI GetMachineGroup is deprecated
//
// Summary:
//
// get machine group
//
// @return GetMachineGroupResponse
// Deprecated
func (client *Client) GetMachineGroup(MachineGroupID *string) (_result *GetMachineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMachineGroupResponse{}
	_body, _err := client.GetMachineGroupWithOptions(MachineGroupID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetNodeMetrics is deprecated
//
// Summary:
//
// get resource group node metrics
//
// @param request - GetNodeMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeMetricsResponse
// Deprecated
func (client *Client) GetNodeMetricsWithOptions(ResourceGroupID *string, MetricType *string, request *GetNodeMetricsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetNodeMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.GPUType)) {
		query["GPUType"] = request.GPUType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TimeStep)) {
		query["TimeStep"] = request.TimeStep
	}

	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		query["Verbose"] = request.Verbose
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNodeMetrics"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceGroupID)) + "/nodemetrics/" + tea.StringValue(openapiutil.GetEncodeParam(MetricType))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetNodeMetricsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetNodeMetricsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Deprecated: OpenAPI GetNodeMetrics is deprecated
//
// Summary:
//
// get resource group node metrics
//
// @param request - GetNodeMetricsRequest
//
// @return GetNodeMetricsResponse
// Deprecated
func (client *Client) GetNodeMetrics(ResourceGroupID *string, MetricType *string, request *GetNodeMetricsRequest) (_result *GetNodeMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetNodeMetricsResponse{}
	_body, _err := client.GetNodeMetricsWithOptions(ResourceGroupID, MetricType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Quota
//
// @param request - GetQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaResponse
func (client *Client) GetQuotaWithOptions(QuotaId *string, request *GetQuotaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		query["Verbose"] = request.Verbose
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQuota"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(QuotaId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetQuotaResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetQuotaResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取Quota
//
// @param request - GetQuotaRequest
//
// @return GetQuotaResponse
func (client *Client) GetQuota(QuotaId *string, request *GetQuotaRequest) (_result *GetQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetQuotaResponse{}
	_body, _err := client.GetQuotaWithOptions(QuotaId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// get resource group by group id
//
// @param tmpReq - GetResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupResponse
func (client *Client) GetResourceGroupWithOptions(ResourceGroupID *string, tmpReq *GetResourceGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetResourceGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetResourceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("Tag"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsAIWorkspaceDataEnabled)) {
		query["IsAIWorkspaceDataEnabled"] = request.IsAIWorkspaceDataEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["Tag"] = request.TagShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceGroup"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceGroupID))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetResourceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetResourceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// get resource group by group id
//
// @param request - GetResourceGroupRequest
//
// @return GetResourceGroupResponse
func (client *Client) GetResourceGroup(ResourceGroupID *string, request *GetResourceGroupRequest) (_result *GetResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceGroupResponse{}
	_body, _err := client.GetResourceGroupWithOptions(ResourceGroupID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// get machine group
//
// @param tmpReq - GetResourceGroupMachineGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupMachineGroupResponse
func (client *Client) GetResourceGroupMachineGroupWithOptions(MachineGroupID *string, ResourceGroupID *string, tmpReq *GetResourceGroupMachineGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetResourceGroupMachineGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetResourceGroupMachineGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("Tag"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["Tag"] = request.TagShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceGroupMachineGroup"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceGroupID)) + "/machinegroups/" + tea.StringValue(openapiutil.GetEncodeParam(MachineGroupID))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetResourceGroupMachineGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetResourceGroupMachineGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// get machine group
//
// @param request - GetResourceGroupMachineGroupRequest
//
// @return GetResourceGroupMachineGroupResponse
func (client *Client) GetResourceGroupMachineGroup(MachineGroupID *string, ResourceGroupID *string, request *GetResourceGroupMachineGroupRequest) (_result *GetResourceGroupMachineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceGroupMachineGroupResponse{}
	_body, _err := client.GetResourceGroupMachineGroupWithOptions(MachineGroupID, ResourceGroupID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetResourceGroupRequest is deprecated
//
// Summary:
//
// get resource group requested resource by resource group id
//
// @param request - GetResourceGroupRequestRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupRequestResponse
// Deprecated
func (client *Client) GetResourceGroupRequestWithOptions(request *GetResourceGroupRequestRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetResourceGroupRequestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PodStatus)) {
		query["PodStatus"] = request.PodStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupID)) {
		query["ResourceGroupID"] = request.ResourceGroupID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceGroupRequest"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resources/data/request"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetResourceGroupRequestResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetResourceGroupRequestResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Deprecated: OpenAPI GetResourceGroupRequest is deprecated
//
// Summary:
//
// get resource group requested resource by resource group id
//
// @param request - GetResourceGroupRequestRequest
//
// @return GetResourceGroupRequestResponse
// Deprecated
func (client *Client) GetResourceGroupRequest(request *GetResourceGroupRequestRequest) (_result *GetResourceGroupRequestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceGroupRequestResponse{}
	_body, _err := client.GetResourceGroupRequestWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// get resource group total resource by group id
//
// @param request - GetResourceGroupTotalRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupTotalResponse
func (client *Client) GetResourceGroupTotalWithOptions(request *GetResourceGroupTotalRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetResourceGroupTotalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupID)) {
		query["ResourceGroupID"] = request.ResourceGroupID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceGroupTotal"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resources/data/total"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetResourceGroupTotalResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetResourceGroupTotalResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// get resource group total resource by group id
//
// @param request - GetResourceGroupTotalRequest
//
// @return GetResourceGroupTotalResponse
func (client *Client) GetResourceGroupTotal(request *GetResourceGroupTotalRequest) (_result *GetResourceGroupTotalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceGroupTotalResponse{}
	_body, _err := client.GetResourceGroupTotalWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用GetToken获取临时鉴权信息
//
// @param request - GetTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenResponse
func (client *Client) GetTokenWithOptions(request *GetTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.TrainingJobId)) {
		query["TrainingJobId"] = request.TrainingJobId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetToken"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tokens"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTokenResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTokenResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 调用GetToken获取临时鉴权信息
//
// @param request - GetTokenRequest
//
// @return GetTokenResponse
func (client *Client) GetToken(request *GetTokenRequest) (_result *GetTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTokenResponse{}
	_body, _err := client.GetTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取TrainingJob的详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrainingJobResponse
func (client *Client) GetTrainingJobWithOptions(TrainingJobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTrainingJobResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrainingJob"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trainingjobs/" + tea.StringValue(openapiutil.GetEncodeParam(TrainingJobId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTrainingJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTrainingJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取TrainingJob的详情
//
// @return GetTrainingJobResponse
func (client *Client) GetTrainingJob(TrainingJobId *string) (_result *GetTrainingJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTrainingJobResponse{}
	_body, _err := client.GetTrainingJobWithOptions(TrainingJobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Training Job的算法错误信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrainingJobErrorInfoResponse
func (client *Client) GetTrainingJobErrorInfoWithOptions(TrainingJobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTrainingJobErrorInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrainingJobErrorInfo"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trainingjobs/" + tea.StringValue(openapiutil.GetEncodeParam(TrainingJobId)) + "/errorinfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTrainingJobErrorInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTrainingJobErrorInfoResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取Training Job的算法错误信息
//
// @return GetTrainingJobErrorInfoResponse
func (client *Client) GetTrainingJobErrorInfo(TrainingJobId *string) (_result *GetTrainingJobErrorInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTrainingJobErrorInfoResponse{}
	_body, _err := client.GetTrainingJobErrorInfoWithOptions(TrainingJobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取TrainingJob最近的Metrics
//
// @param request - GetTrainingJobLatestMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrainingJobLatestMetricsResponse
func (client *Client) GetTrainingJobLatestMetricsWithOptions(TrainingJobId *string, request *GetTrainingJobLatestMetricsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTrainingJobLatestMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Names)) {
		query["Names"] = request.Names
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrainingJobLatestMetrics"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trainingjobs/" + tea.StringValue(openapiutil.GetEncodeParam(TrainingJobId)) + "/latestmetrics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTrainingJobLatestMetricsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTrainingJobLatestMetricsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取TrainingJob最近的Metrics
//
// @param request - GetTrainingJobLatestMetricsRequest
//
// @return GetTrainingJobLatestMetricsResponse
func (client *Client) GetTrainingJobLatestMetrics(TrainingJobId *string, request *GetTrainingJobLatestMetricsRequest) (_result *GetTrainingJobLatestMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTrainingJobLatestMetricsResponse{}
	_body, _err := client.GetTrainingJobLatestMetricsWithOptions(TrainingJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetUserViewMetrics is deprecated
//
// Summary:
//
// get user view  metrics
//
// @param request - GetUserViewMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserViewMetricsResponse
// Deprecated
func (client *Client) GetUserViewMetricsWithOptions(ResourceGroupID *string, request *GetUserViewMetricsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserViewMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.TimeStep)) {
		query["TimeStep"] = request.TimeStep
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserViewMetrics"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceGroupID)) + "/usermetrics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetUserViewMetricsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetUserViewMetricsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Deprecated: OpenAPI GetUserViewMetrics is deprecated
//
// Summary:
//
// get user view  metrics
//
// @param request - GetUserViewMetricsRequest
//
// @return GetUserViewMetricsResponse
// Deprecated
func (client *Client) GetUserViewMetrics(ResourceGroupID *string, request *GetUserViewMetricsRequest) (_result *GetUserViewMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserViewMetricsResponse{}
	_body, _err := client.GetUserViewMetricsWithOptions(ResourceGroupID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取算法的所有版本信息
//
// @param request - ListAlgorithmVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlgorithmVersionsResponse
func (client *Client) ListAlgorithmVersionsWithOptions(AlgorithmId *string, request *ListAlgorithmVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAlgorithmVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlgorithmVersions"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/algorithms/" + tea.StringValue(openapiutil.GetEncodeParam(AlgorithmId)) + "/versions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAlgorithmVersionsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAlgorithmVersionsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取算法的所有版本信息
//
// @param request - ListAlgorithmVersionsRequest
//
// @return ListAlgorithmVersionsResponse
func (client *Client) ListAlgorithmVersions(AlgorithmId *string, request *ListAlgorithmVersionsRequest) (_result *ListAlgorithmVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAlgorithmVersionsResponse{}
	_body, _err := client.ListAlgorithmVersionsWithOptions(AlgorithmId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取算法列表
//
// @param request - ListAlgorithmsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlgorithmsResponse
func (client *Client) ListAlgorithmsWithOptions(request *ListAlgorithmsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAlgorithmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlgorithmId)) {
		query["AlgorithmId"] = request.AlgorithmId
	}

	if !tea.BoolValue(util.IsUnset(request.AlgorithmName)) {
		query["AlgorithmName"] = request.AlgorithmName
	}

	if !tea.BoolValue(util.IsUnset(request.AlgorithmProvider)) {
		query["AlgorithmProvider"] = request.AlgorithmProvider
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlgorithms"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/algorithms"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAlgorithmsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAlgorithmsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取算法列表
//
// @param request - ListAlgorithmsRequest
//
// @return ListAlgorithmsResponse
func (client *Client) ListAlgorithms(request *ListAlgorithmsRequest) (_result *ListAlgorithmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAlgorithmsResponse{}
	_body, _err := client.ListAlgorithmsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取资源节点列表
//
// @param request - ListNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithOptions(request *ListNodesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorType)) {
		query["AcceleratorType"] = request.AcceleratorType
	}

	if !tea.BoolValue(util.IsUnset(request.FilterByQuotaId)) {
		query["FilterByQuotaId"] = request.FilterByQuotaId
	}

	if !tea.BoolValue(util.IsUnset(request.FilterByResourceGroupIds)) {
		query["FilterByResourceGroupIds"] = request.FilterByResourceGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.GPUType)) {
		query["GPUType"] = request.GPUType
	}

	if !tea.BoolValue(util.IsUnset(request.NodeNames)) {
		query["NodeNames"] = request.NodeNames
	}

	if !tea.BoolValue(util.IsUnset(request.NodeStatuses)) {
		query["NodeStatuses"] = request.NodeStatuses
	}

	if !tea.BoolValue(util.IsUnset(request.NodeTypes)) {
		query["NodeTypes"] = request.NodeTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderStatuses)) {
		query["OrderStatuses"] = request.OrderStatuses
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaId)) {
		query["QuotaId"] = request.QuotaId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupIds)) {
		query["ResourceGroupIds"] = request.ResourceGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		query["Verbose"] = request.Verbose
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodes"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/nodes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListNodesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListNodesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取资源节点列表
//
// @param request - ListNodesRequest
//
// @return ListNodesResponse
func (client *Client) ListNodes(request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 您可以通过此API获取Quota上的任务信息列表
//
// @param request - ListQuotaWorkloadsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaWorkloadsResponse
func (client *Client) ListQuotaWorkloadsWithOptions(QuotaId *string, request *ListQuotaWorkloadsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListQuotaWorkloadsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeforeWorkloadId)) {
		query["BeforeWorkloadId"] = request.BeforeWorkloadId
	}

	if !tea.BoolValue(util.IsUnset(request.GmtDequeuedTimeRange)) {
		query["GmtDequeuedTimeRange"] = request.GmtDequeuedTimeRange
	}

	if !tea.BoolValue(util.IsUnset(request.GmtEnqueuedTimeRange)) {
		query["GmtEnqueuedTimeRange"] = request.GmtEnqueuedTimeRange
	}

	if !tea.BoolValue(util.IsUnset(request.GmtPositionModifiedTimeRange)) {
		query["GmtPositionModifiedTimeRange"] = request.GmtPositionModifiedTimeRange
	}

	if !tea.BoolValue(util.IsUnset(request.NodeName)) {
		query["NodeName"] = request.NodeName
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ShowOwn)) {
		query["ShowOwn"] = request.ShowOwn
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.SubQuotaIds)) {
		query["SubQuotaIds"] = request.SubQuotaIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		query["UserIds"] = request.UserIds
	}

	if !tea.BoolValue(util.IsUnset(request.WorkloadCreatedTimeRange)) {
		query["WorkloadCreatedTimeRange"] = request.WorkloadCreatedTimeRange
	}

	if !tea.BoolValue(util.IsUnset(request.WorkloadIds)) {
		query["WorkloadIds"] = request.WorkloadIds
	}

	if !tea.BoolValue(util.IsUnset(request.WorkloadStatuses)) {
		query["WorkloadStatuses"] = request.WorkloadStatuses
	}

	if !tea.BoolValue(util.IsUnset(request.WorkloadType)) {
		query["WorkloadType"] = request.WorkloadType
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceIds)) {
		query["WorkspaceIds"] = request.WorkspaceIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQuotaWorkloads"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(QuotaId)) + "/workloads"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListQuotaWorkloadsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListQuotaWorkloadsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 您可以通过此API获取Quota上的任务信息列表
//
// @param request - ListQuotaWorkloadsRequest
//
// @return ListQuotaWorkloadsResponse
func (client *Client) ListQuotaWorkloads(QuotaId *string, request *ListQuotaWorkloadsRequest) (_result *ListQuotaWorkloadsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQuotaWorkloadsResponse{}
	_body, _err := client.ListQuotaWorkloadsWithOptions(QuotaId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Quota列表
//
// @param request - ListQuotasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotasResponse
func (client *Client) ListQuotasWithOptions(request *ListQuotasRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListQuotasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.LayoutMode)) {
		query["LayoutMode"] = request.LayoutMode
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParentQuotaId)) {
		query["ParentQuotaId"] = request.ParentQuotaId
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaIds)) {
		query["QuotaIds"] = request.QuotaIds
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaName)) {
		query["QuotaName"] = request.QuotaName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Statuses)) {
		query["Statuses"] = request.Statuses
	}

	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		query["Verbose"] = request.Verbose
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceIds)) {
		query["WorkspaceIds"] = request.WorkspaceIds
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceName)) {
		query["WorkspaceName"] = request.WorkspaceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQuotas"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListQuotasResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListQuotasResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取Quota列表
//
// @param request - ListQuotasRequest
//
// @return ListQuotasResponse
func (client *Client) ListQuotas(request *ListQuotasRequest) (_result *ListQuotasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQuotasResponse{}
	_body, _err := client.ListQuotasWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// list machine groups
//
// @param request - ListResourceGroupMachineGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupMachineGroupsResponse
func (client *Client) ListResourceGroupMachineGroupsWithOptions(ResourceGroupID *string, request *ListResourceGroupMachineGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourceGroupMachineGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatorID)) {
		query["CreatorID"] = request.CreatorID
	}

	if !tea.BoolValue(util.IsUnset(request.EcsSpec)) {
		query["EcsSpec"] = request.EcsSpec
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderInstanceId)) {
		query["OrderInstanceId"] = request.OrderInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentDuration)) {
		query["PaymentDuration"] = request.PaymentDuration
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentDurationUnit)) {
		query["PaymentDurationUnit"] = request.PaymentDurationUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentType)) {
		query["PaymentType"] = request.PaymentType
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceGroupMachineGroups"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceGroupID)) + "/machinegroups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListResourceGroupMachineGroupsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListResourceGroupMachineGroupsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// list machine groups
//
// @param request - ListResourceGroupMachineGroupsRequest
//
// @return ListResourceGroupMachineGroupsResponse
func (client *Client) ListResourceGroupMachineGroups(ResourceGroupID *string, request *ListResourceGroupMachineGroupsRequest) (_result *ListResourceGroupMachineGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceGroupMachineGroupsResponse{}
	_body, _err := client.ListResourceGroupMachineGroupsWithOptions(ResourceGroupID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// list resource group
//
// @param request - ListResourceGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupsResponse
func (client *Client) ListResourceGroupsWithOptions(request *ListResourceGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComputingResourceProvider)) {
		query["ComputingResourceProvider"] = request.ComputingResourceProvider
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ShowAll)) {
		query["ShowAll"] = request.ShowAll
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceGroups"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListResourceGroupsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListResourceGroupsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// list resource group
//
// @param request - ListResourceGroupsRequest
//
// @return ListResourceGroupsResponse
func (client *Client) ListResourceGroups(request *ListResourceGroupsRequest) (_result *ListResourceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceGroupsResponse{}
	_body, _err := client.ListResourceGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定TrainingJob的事件。
//
// @param request - ListTrainingJobEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobEventsResponse
func (client *Client) ListTrainingJobEventsWithOptions(TrainingJobId *string, request *ListTrainingJobEventsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTrainingJobEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTrainingJobEvents"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trainingjobs/" + tea.StringValue(openapiutil.GetEncodeParam(TrainingJobId)) + "/events"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTrainingJobEventsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTrainingJobEventsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取指定TrainingJob的事件。
//
// @param request - ListTrainingJobEventsRequest
//
// @return ListTrainingJobEventsResponse
func (client *Client) ListTrainingJobEvents(TrainingJobId *string, request *ListTrainingJobEventsRequest) (_result *ListTrainingJobEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTrainingJobEventsResponse{}
	_body, _err := client.ListTrainingJobEventsWithOptions(TrainingJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定Instance（TrainingJob的运行单元）的日志。
//
// @param request - ListTrainingJobInstanceEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobInstanceEventsResponse
func (client *Client) ListTrainingJobInstanceEventsWithOptions(TrainingJobId *string, InstanceId *string, request *ListTrainingJobInstanceEventsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTrainingJobInstanceEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTrainingJobInstanceEvents"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trainingjobs/" + tea.StringValue(openapiutil.GetEncodeParam(TrainingJobId)) + "/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/events"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTrainingJobInstanceEventsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTrainingJobInstanceEventsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取指定Instance（TrainingJob的运行单元）的日志。
//
// @param request - ListTrainingJobInstanceEventsRequest
//
// @return ListTrainingJobInstanceEventsResponse
func (client *Client) ListTrainingJobInstanceEvents(TrainingJobId *string, InstanceId *string, request *ListTrainingJobInstanceEventsRequest) (_result *ListTrainingJobInstanceEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTrainingJobInstanceEventsResponse{}
	_body, _err := client.ListTrainingJobInstanceEventsWithOptions(TrainingJobId, InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Training Job实例的Metrics
//
// @param request - ListTrainingJobInstanceMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobInstanceMetricsResponse
func (client *Client) ListTrainingJobInstanceMetricsWithOptions(TrainingJobId *string, request *ListTrainingJobInstanceMetricsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTrainingJobInstanceMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TimeStep)) {
		query["TimeStep"] = request.TimeStep
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTrainingJobInstanceMetrics"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trainingjobs/" + tea.StringValue(openapiutil.GetEncodeParam(TrainingJobId)) + "/instancemetrics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTrainingJobInstanceMetricsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTrainingJobInstanceMetricsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取Training Job实例的Metrics
//
// @param request - ListTrainingJobInstanceMetricsRequest
//
// @return ListTrainingJobInstanceMetricsResponse
func (client *Client) ListTrainingJobInstanceMetrics(TrainingJobId *string, request *ListTrainingJobInstanceMetricsRequest) (_result *ListTrainingJobInstanceMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTrainingJobInstanceMetricsResponse{}
	_body, _err := client.ListTrainingJobInstanceMetricsWithOptions(TrainingJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Training Job的日志
//
// @param request - ListTrainingJobLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobLogsResponse
func (client *Client) ListTrainingJobLogsWithOptions(TrainingJobId *string, request *ListTrainingJobLogsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTrainingJobLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.WorkerId)) {
		query["WorkerId"] = request.WorkerId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTrainingJobLogs"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trainingjobs/" + tea.StringValue(openapiutil.GetEncodeParam(TrainingJobId)) + "/logs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTrainingJobLogsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTrainingJobLogsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取Training Job的日志
//
// @param request - ListTrainingJobLogsRequest
//
// @return ListTrainingJobLogsResponse
func (client *Client) ListTrainingJobLogs(TrainingJobId *string, request *ListTrainingJobLogsRequest) (_result *ListTrainingJobLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTrainingJobLogsResponse{}
	_body, _err := client.ListTrainingJobLogsWithOptions(TrainingJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Training Job的Metrics
//
// @param request - ListTrainingJobMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobMetricsResponse
func (client *Client) ListTrainingJobMetricsWithOptions(TrainingJobId *string, request *ListTrainingJobMetricsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTrainingJobMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTrainingJobMetrics"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trainingjobs/" + tea.StringValue(openapiutil.GetEncodeParam(TrainingJobId)) + "/metrics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTrainingJobMetricsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTrainingJobMetricsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取Training Job的Metrics
//
// @param request - ListTrainingJobMetricsRequest
//
// @return ListTrainingJobMetricsResponse
func (client *Client) ListTrainingJobMetrics(TrainingJobId *string, request *ListTrainingJobMetricsRequest) (_result *ListTrainingJobMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTrainingJobMetricsResponse{}
	_body, _err := client.ListTrainingJobMetricsWithOptions(TrainingJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Training Job 产出的所有模型信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobOutputModelsResponse
func (client *Client) ListTrainingJobOutputModelsWithOptions(TrainingJobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTrainingJobOutputModelsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListTrainingJobOutputModels"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trainingjobs/" + tea.StringValue(openapiutil.GetEncodeParam(TrainingJobId)) + "/outputmodels"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTrainingJobOutputModelsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTrainingJobOutputModelsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取Training Job 产出的所有模型信息
//
// @return ListTrainingJobOutputModelsResponse
func (client *Client) ListTrainingJobOutputModels(TrainingJobId *string) (_result *ListTrainingJobOutputModelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTrainingJobOutputModelsResponse{}
	_body, _err := client.ListTrainingJobOutputModelsWithOptions(TrainingJobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取TrainingJob的列表
//
// @param tmpReq - ListTrainingJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobsResponse
func (client *Client) ListTrainingJobsWithOptions(tmpReq *ListTrainingJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTrainingJobsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTrainingJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Labels)) {
		request.LabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Labels, tea.String("Labels"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlgorithmName)) {
		query["AlgorithmName"] = request.AlgorithmName
	}

	if !tea.BoolValue(util.IsUnset(request.AlgorithmProvider)) {
		query["AlgorithmProvider"] = request.AlgorithmProvider
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IsTempAlgo)) {
		query["IsTempAlgo"] = request.IsTempAlgo
	}

	if !tea.BoolValue(util.IsUnset(request.LabelsShrink)) {
		query["Labels"] = request.LabelsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TrainingJobId)) {
		query["TrainingJobId"] = request.TrainingJobId
	}

	if !tea.BoolValue(util.IsUnset(request.TrainingJobName)) {
		query["TrainingJobName"] = request.TrainingJobName
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTrainingJobs"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trainingjobs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTrainingJobsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTrainingJobsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取TrainingJob的列表
//
// @param request - ListTrainingJobsRequest
//
// @return ListTrainingJobsResponse
func (client *Client) ListTrainingJobs(request *ListTrainingJobsRequest) (_result *ListTrainingJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTrainingJobsResponse{}
	_body, _err := client.ListTrainingJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 扩缩容Quota
//
// @param request - ScaleQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScaleQuotaResponse
func (client *Client) ScaleQuotaWithOptions(QuotaId *string, request *ScaleQuotaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ScaleQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Min)) {
		body["Min"] = request.Min
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupIds)) {
		body["ResourceGroupIds"] = request.ResourceGroupIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ScaleQuota"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(QuotaId)) + "/action/scale"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ScaleQuotaResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ScaleQuotaResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 扩缩容Quota
//
// @param request - ScaleQuotaRequest
//
// @return ScaleQuotaResponse
func (client *Client) ScaleQuota(QuotaId *string, request *ScaleQuotaRequest) (_result *ScaleQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ScaleQuotaResponse{}
	_body, _err := client.ScaleQuotaWithOptions(QuotaId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止一个TrainingJob
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTrainingJobResponse
func (client *Client) StopTrainingJobWithOptions(TrainingJobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopTrainingJobResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopTrainingJob"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trainingjobs/" + tea.StringValue(openapiutil.GetEncodeParam(TrainingJobId)) + "/stop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StopTrainingJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StopTrainingJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 停止一个TrainingJob
//
// @return StopTrainingJobResponse
func (client *Client) StopTrainingJob(TrainingJobId *string) (_result *StopTrainingJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopTrainingJobResponse{}
	_body, _err := client.StopTrainingJobWithOptions(TrainingJobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新算法
//
// @param request - UpdateAlgorithmRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlgorithmResponse
func (client *Client) UpdateAlgorithmWithOptions(AlgorithmId *string, request *UpdateAlgorithmRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAlgorithmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlgorithmDescription)) {
		body["AlgorithmDescription"] = request.AlgorithmDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAlgorithm"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/algorithms/" + tea.StringValue(openapiutil.GetEncodeParam(AlgorithmId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateAlgorithmResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateAlgorithmResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新算法
//
// @param request - UpdateAlgorithmRequest
//
// @return UpdateAlgorithmResponse
func (client *Client) UpdateAlgorithm(AlgorithmId *string, request *UpdateAlgorithmRequest) (_result *UpdateAlgorithmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAlgorithmResponse{}
	_body, _err := client.UpdateAlgorithmWithOptions(AlgorithmId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新算法
//
// @param tmpReq - UpdateAlgorithmVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlgorithmVersionResponse
func (client *Client) UpdateAlgorithmVersionWithOptions(AlgorithmId *string, AlgorithmVersion *string, tmpReq *UpdateAlgorithmVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAlgorithmVersionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateAlgorithmVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AlgorithmSpec)) {
		request.AlgorithmSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlgorithmSpec, tea.String("AlgorithmSpec"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlgorithmSpecShrink)) {
		body["AlgorithmSpec"] = request.AlgorithmSpecShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAlgorithmVersion"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/algorithms/" + tea.StringValue(openapiutil.GetEncodeParam(AlgorithmId)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(AlgorithmVersion))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateAlgorithmVersionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateAlgorithmVersionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新算法
//
// @param request - UpdateAlgorithmVersionRequest
//
// @return UpdateAlgorithmVersionResponse
func (client *Client) UpdateAlgorithmVersion(AlgorithmId *string, AlgorithmVersion *string, request *UpdateAlgorithmVersionRequest) (_result *UpdateAlgorithmVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAlgorithmVersionResponse{}
	_body, _err := client.UpdateAlgorithmVersionWithOptions(AlgorithmId, AlgorithmVersion, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新Quota
//
// @param request - UpdateQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQuotaResponse
func (client *Client) UpdateQuotaWithOptions(QuotaId *string, request *UpdateQuotaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.QueueStrategy)) {
		body["QueueStrategy"] = request.QueueStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaConfig)) {
		body["QuotaConfig"] = request.QuotaConfig
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaName)) {
		body["QuotaName"] = request.QuotaName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateQuota"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(QuotaId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateQuotaResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateQuotaResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新Quota
//
// @param request - UpdateQuotaRequest
//
// @return UpdateQuotaResponse
func (client *Client) UpdateQuota(QuotaId *string, request *UpdateQuotaRequest) (_result *UpdateQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateQuotaResponse{}
	_body, _err := client.UpdateQuotaWithOptions(QuotaId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新Resource Group
//
// @param request - UpdateResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceGroupResponse
func (client *Client) UpdateResourceGroupWithOptions(ResourceGroupID *string, request *UpdateResourceGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Unbind)) {
		body["Unbind"] = request.Unbind
	}

	if !tea.BoolValue(util.IsUnset(request.UserVpc)) {
		body["UserVpc"] = request.UserVpc
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResourceGroup"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceGroupID))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateResourceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateResourceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新Resource Group
//
// @param request - UpdateResourceGroupRequest
//
// @return UpdateResourceGroupResponse
func (client *Client) UpdateResourceGroup(ResourceGroupID *string, request *UpdateResourceGroupRequest) (_result *UpdateResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourceGroupResponse{}
	_body, _err := client.UpdateResourceGroupWithOptions(ResourceGroupID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新一个TrainingJob的Labels
//
// @param request - UpdateTrainingJobLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTrainingJobLabelsResponse
func (client *Client) UpdateTrainingJobLabelsWithOptions(TrainingJobId *string, request *UpdateTrainingJobLabelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTrainingJobLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTrainingJobLabels"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/trainingjobs/" + tea.StringValue(openapiutil.GetEncodeParam(TrainingJobId)) + "/labels"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateTrainingJobLabelsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateTrainingJobLabelsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新一个TrainingJob的Labels
//
// @param request - UpdateTrainingJobLabelsRequest
//
// @return UpdateTrainingJobLabelsResponse
func (client *Client) UpdateTrainingJobLabels(TrainingJobId *string, request *UpdateTrainingJobLabelsRequest) (_result *UpdateTrainingJobLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTrainingJobLabelsResponse{}
	_body, _err := client.UpdateTrainingJobLabelsWithOptions(TrainingJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
