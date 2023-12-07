// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
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

type AlgorithmSpec struct {
	CodeDir              *Location                         `json:"CodeDir,omitempty" xml:"CodeDir,omitempty"`
	Command              []*string                         `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	ComputeResource      *AlgorithmSpecComputeResource     `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	Customization        *AlgorithmSpecCustomization       `json:"Customization,omitempty" xml:"Customization,omitempty" type:"Struct"`
	HyperParameters      []*HyperParameterDefinition       `json:"HyperParameters,omitempty" xml:"HyperParameters,omitempty" type:"Repeated"`
	Image                *string                           `json:"Image,omitempty" xml:"Image,omitempty"`
	InputChannels        []*Channel                        `json:"InputChannels,omitempty" xml:"InputChannels,omitempty" type:"Repeated"`
	JobType              *string                           `json:"JobType,omitempty" xml:"JobType,omitempty"`
	MetricDefinitions    []*MetricDefinition               `json:"MetricDefinitions,omitempty" xml:"MetricDefinitions,omitempty" type:"Repeated"`
	OutputChannels       []*Channel                        `json:"OutputChannels,omitempty" xml:"OutputChannels,omitempty" type:"Repeated"`
	ProgressDefinitions  *AlgorithmSpecProgressDefinitions `json:"ProgressDefinitions,omitempty" xml:"ProgressDefinitions,omitempty" type:"Struct"`
	ResourceRequirements []*ConditionExpression            `json:"ResourceRequirements,omitempty" xml:"ResourceRequirements,omitempty" type:"Repeated"`
	// Deprecated
	SupportedInstanceTypes      []*string `json:"SupportedInstanceTypes,omitempty" xml:"SupportedInstanceTypes,omitempty" type:"Repeated"`
	SupportsDistributedTraining *bool     `json:"SupportsDistributedTraining,omitempty" xml:"SupportsDistributedTraining,omitempty"`
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
	Value   *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Description           *string                `json:"Description,omitempty" xml:"Description,omitempty"`
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
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	CodeDir              *Location                   `json:"CodeDir,omitempty" xml:"CodeDir,omitempty"`
	Command              *string                     `json:"Command,omitempty" xml:"Command,omitempty"`
	HyperParameters      []*HyperParameterDefinition `json:"HyperParameters,omitempty" xml:"HyperParameters,omitempty" type:"Repeated"`
	Image                *string                     `json:"Image,omitempty" xml:"Image,omitempty"`
	InputChannels        []*Channel                  `json:"InputChannels,omitempty" xml:"InputChannels,omitempty" type:"Repeated"`
	JobType              *string                     `json:"JobType,omitempty" xml:"JobType,omitempty"`
	MetricDefinitions    []*MetricDefinition         `json:"MetricDefinitions,omitempty" xml:"MetricDefinitions,omitempty" type:"Repeated"`
	OutputChannels       []*Channel                  `json:"OutputChannels,omitempty" xml:"OutputChannels,omitempty" type:"Repeated"`
	ResourceRequirements []*ConditionExpression      `json:"ResourceRequirements,omitempty" xml:"ResourceRequirements,omitempty" type:"Repeated"`
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
	Key      *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Operator *string   `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Values   []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
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

type HyperParameterDefinition struct {
	DefaultValue *string              `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description  *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName  *string              `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Name         *string              `json:"Name,omitempty" xml:"Name,omitempty"`
	Range        *HyperParameterRange `json:"Range,omitempty" xml:"Range,omitempty"`
	Required     *bool                `json:"Required,omitempty" xml:"Required,omitempty"`
	Type         *string              `json:"Type,omitempty" xml:"Type,omitempty"`
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
	ResourceGroupID   *string   `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	TotalCPU          *int32    `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	TotalGPU          *int32    `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	TotalMemory       *int64    `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	UserId            *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	CreatorID           *string   `json:"CreatorID,omitempty" xml:"CreatorID,omitempty"`
	DefaultDriver       *string   `json:"DefaultDriver,omitempty" xml:"DefaultDriver,omitempty"`
	EcsCount            *int64    `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	EcsSpec             *string   `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	GmtCreatedTime      *string   `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	GmtExpiredTime      *string   `json:"GmtExpiredTime,omitempty" xml:"GmtExpiredTime,omitempty"`
	GmtModifiedTime     *string   `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	GmtStartedTime      *string   `json:"GmtStartedTime,omitempty" xml:"GmtStartedTime,omitempty"`
	MachineGroupID      *string   `json:"MachineGroupID,omitempty" xml:"MachineGroupID,omitempty"`
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
	Time  *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
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
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Regex       *string `json:"Regex,omitempty" xml:"Regex,omitempty"`
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
	GPUType           *string        `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	GmtCreateTime     *string        `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtExpiredTime    *string        `json:"GmtExpiredTime,omitempty" xml:"GmtExpiredTime,omitempty"`
	GmtModifiedTime   *string        `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	IsBound           *bool          `json:"IsBound,omitempty" xml:"IsBound,omitempty"`
	MachineGroupId    *string        `json:"MachineGroupId,omitempty" xml:"MachineGroupId,omitempty"`
	Memory            *string        `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NodeName          *string        `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeStatus        *string        `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	NodeType          *string        `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OrderStatus       *string        `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	ReasonCode        *string        `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage     *string        `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	ResourceGroupId   *string        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceGroupName *string        `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
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

func (s *Node) SetReasonCode(v string) *Node {
	s.ReasonCode = &v
	return s
}

func (s *Node) SetReasonMessage(v string) *Node {
	s.ReasonMessage = &v
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

type NodeMetric struct {
	GPUType *string   `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	Metrics []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	NodeID  *string   `json:"NodeID,omitempty" xml:"NodeID,omitempty"`
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

type NodeSpec struct {
	Count *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	CPU             *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	GPU             *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	GPUType         *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	Memory          *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NodeType        *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
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
	CanBeBoundCount *int32  `json:"CanBeBoundCount,omitempty" xml:"CanBeBoundCount,omitempty"`
	NodeType        *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	TotalCount      *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	CPUUsageRate      *string                `json:"CPUUsageRate,omitempty" xml:"CPUUsageRate,omitempty"`
	CreatedTime       *string                `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DiskReadRate      *string                `json:"DiskReadRate,omitempty" xml:"DiskReadRate,omitempty"`
	DiskWriteRate     *string                `json:"DiskWriteRate,omitempty" xml:"DiskWriteRate,omitempty"`
	GPUType           *string                `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	MachineGroupID    *string                `json:"MachineGroupID,omitempty" xml:"MachineGroupID,omitempty"`
	MemoryUsageRate   *string                `json:"MemoryUsageRate,omitempty" xml:"MemoryUsageRate,omitempty"`
	NetworkInputRate  *string                `json:"NetworkInputRate,omitempty" xml:"NetworkInputRate,omitempty"`
	NetworkOutputRate *string                `json:"NetworkOutputRate,omitempty" xml:"NetworkOutputRate,omitempty"`
	NodeID            *string                `json:"NodeID,omitempty" xml:"NodeID,omitempty"`
	NodeStatus        *string                `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	NodeType          *string                `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	RequestCPU        *int64                 `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	RequestGPU        *int64                 `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	RequestMemory     *int64                 `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	TaskIdMap         map[string]interface{} `json:"TaskIdMap,omitempty" xml:"TaskIdMap,omitempty"`
	TotalCPU          *int64                 `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	TotalGPU          *int64                 `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	TotalMemory       *int64                 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	TotalTasks        *int64                 `json:"TotalTasks,omitempty" xml:"TotalTasks,omitempty"`
	UserIDs           []*string              `json:"UserIDs,omitempty" xml:"UserIDs,omitempty" type:"Repeated"`
	UserNumber        *string                `json:"UserNumber,omitempty" xml:"UserNumber,omitempty"`
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

type Quota struct {
	AllocateStrategy  *string            `json:"AllocateStrategy,omitempty" xml:"AllocateStrategy,omitempty"`
	CreatorId         *string            `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Description       *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreatedTime    *string            `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	GmtModifiedTime   *string            `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels            []*Label           `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LatestOperationId *string            `json:"LatestOperationId,omitempty" xml:"LatestOperationId,omitempty"`
	Min               *ResourceSpec      `json:"Min,omitempty" xml:"Min,omitempty"`
	ParentQuotaId     *string            `json:"ParentQuotaId,omitempty" xml:"ParentQuotaId,omitempty"`
	QueueStrategy     *string            `json:"QueueStrategy,omitempty" xml:"QueueStrategy,omitempty"`
	QuotaConfig       *QuotaConfig       `json:"QuotaConfig,omitempty" xml:"QuotaConfig,omitempty"`
	QuotaDetails      *QuotaDetails      `json:"QuotaDetails,omitempty" xml:"QuotaDetails,omitempty"`
	QuotaId           *string            `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	QuotaName         *string            `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	ReasonCode        *string            `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage     *string            `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	ResourceGroupIds  []*string          `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
	ResourceType      *string            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status            *string            `json:"Status,omitempty" xml:"Status,omitempty"`
	SubQuotas         []*QuotaIdName     `json:"SubQuotas,omitempty" xml:"SubQuotas,omitempty" type:"Repeated"`
	Workspaces        []*WorkspaceIdName `json:"Workspaces,omitempty" xml:"Workspaces,omitempty" type:"Repeated"`
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
	ACS               *ACS      `json:"ACS,omitempty" xml:"ACS,omitempty"`
	ClusterId         *string   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DefaultGPUDriver  *string   `json:"DefaultGPUDriver,omitempty" xml:"DefaultGPUDriver,omitempty"`
	SupportGPUDrivers []*string `json:"SupportGPUDrivers,omitempty" xml:"SupportGPUDrivers,omitempty" type:"Repeated"`
	SupportRDMA       *bool     `json:"SupportRDMA,omitempty" xml:"SupportRDMA,omitempty"`
	UserVpc           *UserVpc  `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
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
	ActualMinQuota  *ResourceAmount `json:"ActualMinQuota,omitempty" xml:"ActualMinQuota,omitempty"`
	DesiredMinQuota *ResourceAmount `json:"DesiredMinQuota,omitempty" xml:"DesiredMinQuota,omitempty"`
	RequestedQuota  *ResourceAmount `json:"RequestedQuota,omitempty" xml:"RequestedQuota,omitempty"`
	UsedQuota       *ResourceAmount `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty"`
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

func (s *QuotaDetails) SetDesiredMinQuota(v *ResourceAmount) *QuotaDetails {
	s.DesiredMinQuota = v
	return s
}

func (s *QuotaDetails) SetRequestedQuota(v *ResourceAmount) *QuotaDetails {
	s.RequestedQuota = v
	return s
}

func (s *QuotaDetails) SetUsedQuota(v *ResourceAmount) *QuotaDetails {
	s.UsedQuota = v
	return s
}

type QuotaIdName struct {
	QuotaId   *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
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
	GPUType *string   `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
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
	CPUUsageRate      *string                `json:"CPUUsageRate,omitempty" xml:"CPUUsageRate,omitempty"`
	CreatedTime       *string                `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DiskReadRate      *string                `json:"DiskReadRate,omitempty" xml:"DiskReadRate,omitempty"`
	DiskWriteRate     *string                `json:"DiskWriteRate,omitempty" xml:"DiskWriteRate,omitempty"`
	GPUType           *string                `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	MemoryUsageRate   *string                `json:"MemoryUsageRate,omitempty" xml:"MemoryUsageRate,omitempty"`
	NetworkInputRate  *string                `json:"NetworkInputRate,omitempty" xml:"NetworkInputRate,omitempty"`
	NetworkOutputRate *string                `json:"NetworkOutputRate,omitempty" xml:"NetworkOutputRate,omitempty"`
	NodeID            *string                `json:"NodeID,omitempty" xml:"NodeID,omitempty"`
	NodeStatus        *string                `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	NodeType          *string                `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	QuotaId           *string                `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	RequestCPU        *int64                 `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	RequestGPU        *int64                 `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	RequestMemory     *int64                 `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	TaskIdMap         map[string]interface{} `json:"TaskIdMap,omitempty" xml:"TaskIdMap,omitempty"`
	TotalCPU          *int64                 `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	TotalGPU          *int64                 `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	TotalMemory       *int64                 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	TotalTasks        *int64                 `json:"TotalTasks,omitempty" xml:"TotalTasks,omitempty"`
	UserIDs           []*string              `json:"UserIDs,omitempty" xml:"UserIDs,omitempty" type:"Repeated"`
	UserNumber        *string                `json:"UserNumber,omitempty" xml:"UserNumber,omitempty"`
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
	CPU     *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	GPU     *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	Memory  *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
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

type ResourceGroup struct {
	CreatorID       *string  `json:"CreatorID,omitempty" xml:"CreatorID,omitempty"`
	GmtCreatedTime  *string  `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	GmtModifiedTime *string  `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Name            *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	NodeCount       *int32   `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	ResourceGroupID *string  `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	UserVpc         *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
	WorkspaceID     *string  `json:"WorkspaceID,omitempty" xml:"WorkspaceID,omitempty"`
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
	GpuType         *string   `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
	Metrics         []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	ResourceGroupID *string   `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
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
	ResourceGroupId   *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TotalCPU          *int32    `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	TotalGPU          *int32    `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	TotalMemory       *int64    `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	UserId            *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	DefaultRoute    *string   `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	ExtendedCIDRs   []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	RoleArn         *string   `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SwitchId        *string   `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	VpcId           *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UserVpc) String() string {
	return tea.Prettify(s)
}

func (s UserVpc) GoString() string {
	return s.String()
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

type WorkspaceIdName struct {
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

type CreateAlgorithmRequest struct {
	AlgorithmDescription *string `json:"AlgorithmDescription,omitempty" xml:"AlgorithmDescription,omitempty"`
	AlgorithmName        *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	DisplayName          *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	WorkspaceId          *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AlgorithmId      *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAlgorithmVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateQuotaRequest struct {
	AllocateStrategy *string       `json:"AllocateStrategy,omitempty" xml:"AllocateStrategy,omitempty"`
	Description      *string       `json:"Description,omitempty" xml:"Description,omitempty"`
	Labels           []*Label      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Min              *ResourceSpec `json:"Min,omitempty" xml:"Min,omitempty"`
	ParentQuotaId    *string       `json:"ParentQuotaId,omitempty" xml:"ParentQuotaId,omitempty"`
	QueueStrategy    *string       `json:"QueueStrategy,omitempty" xml:"QueueStrategy,omitempty"`
	QuotaConfig      *QuotaConfig  `json:"QuotaConfig,omitempty" xml:"QuotaConfig,omitempty"`
	QuotaName        *string       `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	ResourceGroupIds []*string     `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
	ResourceType     *string       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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
	QuotaId   *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ComputingResourceProvider *string  `json:"ComputingResourceProvider,omitempty" xml:"ComputingResourceProvider,omitempty"`
	Description               *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                      *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceType              *string  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	UserVpc                   *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
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

func (s *CreateResourceGroupRequest) SetUserVpc(v *UserVpc) *CreateResourceGroupRequest {
	s.UserVpc = v
	return s
}

type CreateResourceGroupResponseBody struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AlgorithmName          *string                                    `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	AlgorithmProvider      *string                                    `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	AlgorithmSpec          *AlgorithmSpec                             `json:"AlgorithmSpec,omitempty" xml:"AlgorithmSpec,omitempty"`
	AlgorithmVersion       *string                                    `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	CodeDir                *Location                                  `json:"CodeDir,omitempty" xml:"CodeDir,omitempty"`
	ComputeResource        *CreateTrainingJobRequestComputeResource   `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	HyperParameters        []*CreateTrainingJobRequestHyperParameters `json:"HyperParameters,omitempty" xml:"HyperParameters,omitempty" type:"Repeated"`
	InputChannels          []*CreateTrainingJobRequestInputChannels   `json:"InputChannels,omitempty" xml:"InputChannels,omitempty" type:"Repeated"`
	Labels                 []*CreateTrainingJobRequestLabels          `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	OutputChannels         []*CreateTrainingJobRequestOutputChannels  `json:"OutputChannels,omitempty" xml:"OutputChannels,omitempty" type:"Repeated"`
	RoleArn                *string                                    `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	Scheduler              *CreateTrainingJobRequestScheduler         `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Struct"`
	TrainingJobDescription *string                                    `json:"TrainingJobDescription,omitempty" xml:"TrainingJobDescription,omitempty"`
	TrainingJobName        *string                                    `json:"TrainingJobName,omitempty" xml:"TrainingJobName,omitempty"`
	UserVpc                *CreateTrainingJobRequestUserVpc           `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	WorkspaceId            *string                                    `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

func (s *CreateTrainingJobRequest) SetRoleArn(v string) *CreateTrainingJobRequest {
	s.RoleArn = &v
	return s
}

func (s *CreateTrainingJobRequest) SetScheduler(v *CreateTrainingJobRequestScheduler) *CreateTrainingJobRequest {
	s.Scheduler = v
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
	EcsCount      *int64                                               `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	EcsSpec       *string                                              `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	InstanceCount *int64                                               `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	InstanceSpec  *CreateTrainingJobRequestComputeResourceInstanceSpec `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" type:"Struct"`
	ResourceId    *string                                              `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
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

type CreateTrainingJobRequestComputeResourceInstanceSpec struct {
	CPU          *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	GPU          *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	GPUType      *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	Memory       *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
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

type CreateTrainingJobRequestHyperParameters struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	InputUri  *string `json:"InputUri,omitempty" xml:"InputUri,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

type CreateTrainingJobRequestLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OutputUri *string `json:"OutputUri,omitempty" xml:"OutputUri,omitempty"`
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

type CreateTrainingJobRequestScheduler struct {
	MaxRunningTimeInSeconds *int64 `json:"MaxRunningTimeInSeconds,omitempty" xml:"MaxRunningTimeInSeconds,omitempty"`
}

func (s CreateTrainingJobRequestScheduler) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainingJobRequestScheduler) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestScheduler) SetMaxRunningTimeInSeconds(v int64) *CreateTrainingJobRequestScheduler {
	s.MaxRunningTimeInSeconds = &v
	return s
}

type CreateTrainingJobRequestUserVpc struct {
	ExtendedCIDRs   []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SwitchId        *string   `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	VpcId           *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateTrainingJobRequestUserVpc) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainingJobRequestUserVpc) GoString() string {
	return s.String()
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
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTrainingJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteMachineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	QuotaId   *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteResourceGroupMachineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetAlgorithmResponseBody struct {
	AlgorithmDescription *string `json:"AlgorithmDescription,omitempty" xml:"AlgorithmDescription,omitempty"`
	AlgorithmId          *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	AlgorithmName        *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	AlgorithmProvider    *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	DisplayName          *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	GmtCreateTime        *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime      *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TenantId             *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UserId               *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId          *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AlgorithmId       *string        `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	AlgorithmName     *string        `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	AlgorithmProvider *string        `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	AlgorithmSpec     *AlgorithmSpec `json:"AlgorithmSpec,omitempty" xml:"AlgorithmSpec,omitempty"`
	AlgorithmVersion  *string        `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	GmtCreateTime     *string        `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime   *string        `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	TenantId          *string        `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UserId            *string        `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAlgorithmVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMachineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GPUType   *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeStep  *string `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
	Verbose   *bool   `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
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
	MetricType      *string       `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	NodesMetrics    []*NodeMetric `json:"NodesMetrics,omitempty" xml:"NodesMetrics,omitempty" type:"Repeated"`
	ResourceGroupID *string       `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNodeMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetQuotaResponseBody struct {
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
	// Quota Id
	QuotaId          *string            `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	QuotaName        *string            `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	ReasonCode       *string            `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage    *string            `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	RequestId        *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupIds []*string          `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
	ResourceType     *string            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status           *string            `json:"Status,omitempty" xml:"Status,omitempty"`
	SubQuotas        []*QuotaIdName     `json:"SubQuotas,omitempty" xml:"SubQuotas,omitempty" type:"Repeated"`
	Workspaces       []*WorkspaceIdName `json:"Workspaces,omitempty" xml:"Workspaces,omitempty" type:"Repeated"`
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	IsAIWorkspaceDataEnabled *bool `json:"IsAIWorkspaceDataEnabled,omitempty" xml:"IsAIWorkspaceDataEnabled,omitempty"`
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

type GetResourceGroupResponseBody struct {
	ClusterID                 *string  `json:"ClusterID,omitempty" xml:"ClusterID,omitempty"`
	ComputingResourceProvider *string  `json:"ComputingResourceProvider,omitempty" xml:"ComputingResourceProvider,omitempty"`
	CreatorID                 *string  `json:"CreatorID,omitempty" xml:"CreatorID,omitempty"`
	GmtCreatedTime            *string  `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	GmtModifiedTime           *string  `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Name                      *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId                 *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceType              *string  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status                    *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportRDMA               *bool    `json:"SupportRDMA,omitempty" xml:"SupportRDMA,omitempty"`
	UserVpc                   *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
	WorkspaceID               *string  `json:"WorkspaceID,omitempty" xml:"WorkspaceID,omitempty"`
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

func (s *GetResourceGroupResponseBody) SetUserVpc(v *UserVpc) *GetResourceGroupResponseBody {
	s.UserVpc = v
	return s
}

func (s *GetResourceGroupResponseBody) SetWorkspaceID(v string) *GetResourceGroupResponseBody {
	s.WorkspaceID = &v
	return s
}

type GetResourceGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetResourceGroupMachineGroupResponseBody struct {
	Cpu                 *string   `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	DefaultDriver       *string   `json:"DefaultDriver,omitempty" xml:"DefaultDriver,omitempty"`
	EcsCount            *int64    `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	EcsSpec             *string   `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	GmtCreatedTime      *string   `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	GmtExpiredTime      *string   `json:"GmtExpiredTime,omitempty" xml:"GmtExpiredTime,omitempty"`
	GmtModifiedTime     *string   `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	GmtStartedTime      *string   `json:"GmtStartedTime,omitempty" xml:"GmtStartedTime,omitempty"`
	Gpu                 *string   `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	GpuType             *string   `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
	MachineGroupID      *string   `json:"MachineGroupID,omitempty" xml:"MachineGroupID,omitempty"`
	Memory              *string   `json:"Memory,omitempty" xml:"Memory,omitempty"`
	PaymentDuration     *string   `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	PaymentDurationUnit *string   `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	PaymentType         *string   `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	RequestId           *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupID     *string   `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	Status              *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportedDrivers    []*string `json:"SupportedDrivers,omitempty" xml:"SupportedDrivers,omitempty" type:"Repeated"`
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

type GetResourceGroupMachineGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceGroupMachineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PodStatus       *string `json:"PodStatus,omitempty" xml:"PodStatus,omitempty"`
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
	RequestCPU      *int32     `json:"requestCPU,omitempty" xml:"requestCPU,omitempty"`
	RequestGPU      *int32     `json:"requestGPU,omitempty" xml:"requestGPU,omitempty"`
	RequestGPUInfos []*GPUInfo `json:"requestGPUInfos,omitempty" xml:"requestGPUInfos,omitempty" type:"Repeated"`
	RequestMemory   *int32     `json:"requestMemory,omitempty" xml:"requestMemory,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceGroupRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	TotalCPU      *int32     `json:"totalCPU,omitempty" xml:"totalCPU,omitempty"`
	TotalGPU      *int32     `json:"totalGPU,omitempty" xml:"totalGPU,omitempty"`
	TotalGPUInfos []*GPUInfo `json:"totalGPUInfos,omitempty" xml:"totalGPUInfos,omitempty" type:"Repeated"`
	TotalMemory   *int32     `json:"totalMemory,omitempty" xml:"totalMemory,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceGroupTotalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetTrainingJobResponseBody struct {
	AlgorithmId            *string                                        `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	AlgorithmName          *string                                        `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	AlgorithmProvider      *string                                        `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	AlgorithmSpec          *AlgorithmSpec                                 `json:"AlgorithmSpec,omitempty" xml:"AlgorithmSpec,omitempty"`
	AlgorithmVersion       *string                                        `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	ComputeResource        *GetTrainingJobResponseBodyComputeResource     `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	GmtCreateTime          *string                                        `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime        *string                                        `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	HyperParameters        []*GetTrainingJobResponseBodyHyperParameters   `json:"HyperParameters,omitempty" xml:"HyperParameters,omitempty" type:"Repeated"`
	InputChannels          []*GetTrainingJobResponseBodyInputChannels     `json:"InputChannels,omitempty" xml:"InputChannels,omitempty" type:"Repeated"`
	Instances              []*GetTrainingJobResponseBodyInstances         `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	IsTempAlgo             *bool                                          `json:"IsTempAlgo,omitempty" xml:"IsTempAlgo,omitempty"`
	Labels                 []*GetTrainingJobResponseBodyLabels            `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LatestMetrics          []*GetTrainingJobResponseBodyLatestMetrics     `json:"LatestMetrics,omitempty" xml:"LatestMetrics,omitempty" type:"Repeated"`
	LatestProgress         *GetTrainingJobResponseBodyLatestProgress      `json:"LatestProgress,omitempty" xml:"LatestProgress,omitempty" type:"Struct"`
	OutputChannels         []*GetTrainingJobResponseBodyOutputChannels    `json:"OutputChannels,omitempty" xml:"OutputChannels,omitempty" type:"Repeated"`
	OutputModel            *GetTrainingJobResponseBodyOutputModel         `json:"OutputModel,omitempty" xml:"OutputModel,omitempty" type:"Struct"`
	ReasonCode             *string                                        `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage          *string                                        `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	RequestId              *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoleArn                *string                                        `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	Scheduler              *GetTrainingJobResponseBodyScheduler           `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Struct"`
	Status                 *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusTransitions      []*GetTrainingJobResponseBodyStatusTransitions `json:"StatusTransitions,omitempty" xml:"StatusTransitions,omitempty" type:"Repeated"`
	TrainingJobDescription *string                                        `json:"TrainingJobDescription,omitempty" xml:"TrainingJobDescription,omitempty"`
	TrainingJobId          *string                                        `json:"TrainingJobId,omitempty" xml:"TrainingJobId,omitempty"`
	TrainingJobName        *string                                        `json:"TrainingJobName,omitempty" xml:"TrainingJobName,omitempty"`
	TrainingJobUrl         *string                                        `json:"TrainingJobUrl,omitempty" xml:"TrainingJobUrl,omitempty"`
	UserId                 *string                                        `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserVpc                *GetTrainingJobResponseBodyUserVpc             `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	WorkspaceId            *string                                        `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	EcsCount      *int64                                                 `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	EcsSpec       *string                                                `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	InstanceCount *int64                                                 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	InstanceSpec  *GetTrainingJobResponseBodyComputeResourceInstanceSpec `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" type:"Struct"`
	ResourceId    *string                                                `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
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

type GetTrainingJobResponseBodyComputeResourceInstanceSpec struct {
	CPU          *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	GPU          *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	GPUType      *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	Memory       *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
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

type GetTrainingJobResponseBodyHyperParameters struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	InputUri  *string `json:"InputUri,omitempty" xml:"InputUri,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

type GetTrainingJobResponseBodyInstances struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Role   *string `json:"Role,omitempty" xml:"Role,omitempty"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Name      *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Timestamp *string  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Value     *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Timestamp *string  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Value     *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Value     *int64  `json:"Value,omitempty" xml:"Value,omitempty"`
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
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OutputUri *string `json:"OutputUri,omitempty" xml:"OutputUri,omitempty"`
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

type GetTrainingJobResponseBodyOutputModel struct {
	OutputChannelName *string `json:"OutputChannelName,omitempty" xml:"OutputChannelName,omitempty"`
	Uri               *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
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
	MaxRunningTimeInSeconds *int64 `json:"MaxRunningTimeInSeconds,omitempty" xml:"MaxRunningTimeInSeconds,omitempty"`
}

func (s GetTrainingJobResponseBodyScheduler) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBodyScheduler) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyScheduler) SetMaxRunningTimeInSeconds(v int64) *GetTrainingJobResponseBodyScheduler {
	s.MaxRunningTimeInSeconds = &v
	return s
}

type GetTrainingJobResponseBodyStatusTransitions struct {
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ReasonCode    *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	ExtendedCIDRs   []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SwitchId        *string   `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	VpcId           *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTrainingJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetUserViewMetricsRequest struct {
	Order       *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber  *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy      *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	TimeStep    *string `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	ResourceGroupId *string           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Summary         *UserViewMetric   `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Total           *int32            `json:"Total,omitempty" xml:"Total,omitempty"`
	UserMetrics     []*UserViewMetric `json:"UserMetrics,omitempty" xml:"UserMetrics,omitempty" type:"Repeated"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserViewMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount        *int64                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AlgorithmId       *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	AlgorithmName     *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	AlgorithmVersion  *string `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	GmtCreateTime     *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime   *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	TenantId          *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAlgorithmVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AlgorithmId       *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	AlgorithmName     *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	PageNumber        *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	WorkspaceId       *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AlgorithmId          *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	AlgorithmName        *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	AlgorithmProvider    *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	DisplayName          *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	GmtCreateTime        *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime      *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	UserId               *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId          *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAlgorithmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListQuotasRequest struct {
	Labels        *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Order         *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ParentQuotaId *string `json:"ParentQuotaId,omitempty" xml:"ParentQuotaId,omitempty"`
	QuotaIds      *string `json:"QuotaIds,omitempty" xml:"QuotaIds,omitempty"`
	QuotaName     *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SortBy        *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Statuses      *string `json:"Statuses,omitempty" xml:"Statuses,omitempty"`
	WorkspaceIds  *string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty"`
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

func (s *ListQuotasRequest) SetWorkspaceIds(v string) *ListQuotasRequest {
	s.WorkspaceIds = &v
	return s
}

type ListQuotasResponseBody struct {
	Quotas    []*Quota `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type ListQuotasResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CreatorID           *string `json:"CreatorID,omitempty" xml:"CreatorID,omitempty"`
	EcsSpec             *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order               *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber          *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PaymentDuration     *string `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	PaymentDurationUnit *string `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	PaymentType         *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	SortBy              *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	RequestId     *string         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *string         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourceGroupMachineGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ComputingResourceProvider *string `json:"ComputingResourceProvider,omitempty" xml:"ComputingResourceProvider,omitempty"`
	Name                      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order                     *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber                *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                  *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceType              *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ShowAll                   *bool   `json:"ShowAll,omitempty" xml:"ShowAll,omitempty"`
	SortBy                    *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	RequestId      *string          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroups []*ResourceGroup `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Repeated"`
	TotalCount     *int64           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListTrainingJobLogsRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	WorkerId   *string `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
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
	Logs       []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTrainingJobLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order      *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	Metrics   []*ListTrainingJobMetricsResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type ListTrainingJobMetricsResponseBodyMetrics struct {
	Name      *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Timestamp *string  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Value     *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTrainingJobMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListTrainingJobsRequest struct {
	AlgorithmName     *string                `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	AlgorithmProvider *string                `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	EndTime           *string                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IsTempAlgo        *bool                  `json:"IsTempAlgo,omitempty" xml:"IsTempAlgo,omitempty"`
	Labels            map[string]interface{} `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Order             *string                `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber        *int64                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int64                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy            *string                `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	StartTime         *string                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status            *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	TrainingJobId     *string                `json:"TrainingJobId,omitempty" xml:"TrainingJobId,omitempty"`
	TrainingJobName   *string                `json:"TrainingJobName,omitempty" xml:"TrainingJobName,omitempty"`
	WorkspaceId       *string                `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	AlgorithmName     *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IsTempAlgo        *bool   `json:"IsTempAlgo,omitempty" xml:"IsTempAlgo,omitempty"`
	LabelsShrink      *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Order             *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber        *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy            *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TrainingJobId     *string `json:"TrainingJobId,omitempty" xml:"TrainingJobId,omitempty"`
	TrainingJobName   *string `json:"TrainingJobName,omitempty" xml:"TrainingJobName,omitempty"`
	WorkspaceId       *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AlgorithmName          *string                                                      `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	AlgorithmProvider      *string                                                      `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	AlgorithmVersion       *string                                                      `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	ComputeResource        *ListTrainingJobsResponseBodyTrainingJobsComputeResource     `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	GmtCreateTime          *string                                                      `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime        *string                                                      `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	HyperParameters        []*ListTrainingJobsResponseBodyTrainingJobsHyperParameters   `json:"HyperParameters,omitempty" xml:"HyperParameters,omitempty" type:"Repeated"`
	InputChannels          []*ListTrainingJobsResponseBodyTrainingJobsInputChannels     `json:"InputChannels,omitempty" xml:"InputChannels,omitempty" type:"Repeated"`
	IsTempAlgo             *bool                                                        `json:"IsTempAlgo,omitempty" xml:"IsTempAlgo,omitempty"`
	Labels                 []*ListTrainingJobsResponseBodyTrainingJobsLabels            `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	OutputChannels         []*ListTrainingJobsResponseBodyTrainingJobsOutputChannels    `json:"OutputChannels,omitempty" xml:"OutputChannels,omitempty" type:"Repeated"`
	ReasonCode             *string                                                      `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage          *string                                                      `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	RoleArn                *string                                                      `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	Scheduler              *ListTrainingJobsResponseBodyTrainingJobsScheduler           `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Struct"`
	Status                 *string                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusTransitions      []*ListTrainingJobsResponseBodyTrainingJobsStatusTransitions `json:"StatusTransitions,omitempty" xml:"StatusTransitions,omitempty" type:"Repeated"`
	TrainingJobDescription *string                                                      `json:"TrainingJobDescription,omitempty" xml:"TrainingJobDescription,omitempty"`
	TrainingJobId          *string                                                      `json:"TrainingJobId,omitempty" xml:"TrainingJobId,omitempty"`
	TrainingJobName        *string                                                      `json:"TrainingJobName,omitempty" xml:"TrainingJobName,omitempty"`
	UserId                 *string                                                      `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserVpc                *ListTrainingJobsResponseBodyTrainingJobsUserVpc             `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	WorkspaceId            *string                                                      `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	EcsCount      *int64                                                               `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	EcsSpec       *string                                                              `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	InstanceCount *int64                                                               `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	InstanceSpec  *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" type:"Struct"`
	ResourceId    *string                                                              `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
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

type ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec struct {
	CPU          *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	GPU          *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	GPUType      *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	Memory       *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
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

type ListTrainingJobsResponseBodyTrainingJobsHyperParameters struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	InputUri  *string `json:"InputUri,omitempty" xml:"InputUri,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

type ListTrainingJobsResponseBodyTrainingJobsLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OutputUri *string `json:"OutputUri,omitempty" xml:"OutputUri,omitempty"`
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

type ListTrainingJobsResponseBodyTrainingJobsScheduler struct {
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
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ReasonCode    *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	ExtendedCIDRs   []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SwitchId        *string   `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	VpcId           *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsUserVpc) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsUserVpc) GoString() string {
	return s.String()
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTrainingJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	QuotaId   *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ScaleQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopTrainingJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DisplayName          *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
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
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AlgorithmId      *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAlgorithmVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Description *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Labels      []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
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

type UpdateQuotaResponseBody struct {
	// Quota Id
	QuotaId   *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Unbind  *bool    `json:"Unbind,omitempty" xml:"Unbind,omitempty"`
	UserVpc *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
}

func (s UpdateResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupRequest) GoString() string {
	return s.String()
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
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	RequestId       *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTrainingJobLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	_result = &CreateAlgorithmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateAlgorithmVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

	if !tea.BoolValue(util.IsUnset(request.RoleArn)) {
		body["RoleArn"] = request.RoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.Scheduler)) {
		body["Scheduler"] = request.Scheduler
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
	_result = &CreateTrainingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DeleteMachineGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DeleteQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DeleteResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DeleteResourceGroupMachineGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetAlgorithmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetAlgorithmVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetMachineGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetNodeMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) GetQuotaWithOptions(QuotaId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetQuotaResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
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
	_result = &GetQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQuota(QuotaId *string) (_result *GetQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetQuotaResponse{}
	_body, _err := client.GetQuotaWithOptions(QuotaId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourceGroupWithOptions(ResourceGroupID *string, request *GetResourceGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsAIWorkspaceDataEnabled)) {
		query["IsAIWorkspaceDataEnabled"] = request.IsAIWorkspaceDataEnabled
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
	_result = &GetResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) GetResourceGroupMachineGroupWithOptions(MachineGroupID *string, ResourceGroupID *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetResourceGroupMachineGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
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
	_result = &GetResourceGroupMachineGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceGroupMachineGroup(MachineGroupID *string, ResourceGroupID *string) (_result *GetResourceGroupMachineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceGroupMachineGroupResponse{}
	_body, _err := client.GetResourceGroupMachineGroupWithOptions(MachineGroupID, ResourceGroupID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
	_result = &GetResourceGroupRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetResourceGroupTotalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetTrainingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetUserViewMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListAlgorithmVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListAlgorithmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ListQuotasWithOptions(request *ListQuotasRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListQuotasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
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

	if !tea.BoolValue(util.IsUnset(request.WorkspaceIds)) {
		query["WorkspaceIds"] = request.WorkspaceIds
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
	_result = &ListQuotasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListResourceGroupMachineGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListResourceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListTrainingJobLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListTrainingJobMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListTrainingJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ScaleQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &StopTrainingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpdateAlgorithmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpdateAlgorithmVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpdateQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) UpdateResourceGroupWithOptions(ResourceGroupID *string, request *UpdateResourceGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
	_result = &UpdateResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpdateTrainingJobLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
