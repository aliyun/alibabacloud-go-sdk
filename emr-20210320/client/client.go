// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AckNode struct {
	NodeId       *string              `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeSelector *AckNodeNodeSelector `json:"NodeSelector,omitempty" xml:"NodeSelector,omitempty" type:"Struct"`
}

func (s AckNode) String() string {
	return tea.Prettify(s)
}

func (s AckNode) GoString() string {
	return s.String()
}

func (s *AckNode) SetNodeId(v string) *AckNode {
	s.NodeId = &v
	return s
}

func (s *AckNode) SetNodeSelector(v *AckNodeNodeSelector) *AckNode {
	s.NodeSelector = v
	return s
}

type AckNodeNodeSelector struct {
	Labels []*AckNodeNodeSelectorLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Taints []*AckNodeNodeSelectorTaints `json:"Taints,omitempty" xml:"Taints,omitempty" type:"Repeated"`
}

func (s AckNodeNodeSelector) String() string {
	return tea.Prettify(s)
}

func (s AckNodeNodeSelector) GoString() string {
	return s.String()
}

func (s *AckNodeNodeSelector) SetLabels(v []*AckNodeNodeSelectorLabels) *AckNodeNodeSelector {
	s.Labels = v
	return s
}

func (s *AckNodeNodeSelector) SetTaints(v []*AckNodeNodeSelectorTaints) *AckNodeNodeSelector {
	s.Taints = v
	return s
}

type AckNodeNodeSelectorLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AckNodeNodeSelectorLabels) String() string {
	return tea.Prettify(s)
}

func (s AckNodeNodeSelectorLabels) GoString() string {
	return s.String()
}

func (s *AckNodeNodeSelectorLabels) SetKey(v string) *AckNodeNodeSelectorLabels {
	s.Key = &v
	return s
}

func (s *AckNodeNodeSelectorLabels) SetValue(v string) *AckNodeNodeSelectorLabels {
	s.Value = &v
	return s
}

type AckNodeNodeSelectorTaints struct {
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	Key    *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value  *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AckNodeNodeSelectorTaints) String() string {
	return tea.Prettify(s)
}

func (s AckNodeNodeSelectorTaints) GoString() string {
	return s.String()
}

func (s *AckNodeNodeSelectorTaints) SetEffect(v string) *AckNodeNodeSelectorTaints {
	s.Effect = &v
	return s
}

func (s *AckNodeNodeSelectorTaints) SetKey(v string) *AckNodeNodeSelectorTaints {
	s.Key = &v
	return s
}

func (s *AckNodeNodeSelectorTaints) SetValue(v string) *AckNodeNodeSelectorTaints {
	s.Value = &v
	return s
}

type AckNodePool struct {
	NodePoolId   *string                  `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
	NodeSelector *AckNodePoolNodeSelector `json:"NodeSelector,omitempty" xml:"NodeSelector,omitempty" type:"Struct"`
}

func (s AckNodePool) String() string {
	return tea.Prettify(s)
}

func (s AckNodePool) GoString() string {
	return s.String()
}

func (s *AckNodePool) SetNodePoolId(v string) *AckNodePool {
	s.NodePoolId = &v
	return s
}

func (s *AckNodePool) SetNodeSelector(v *AckNodePoolNodeSelector) *AckNodePool {
	s.NodeSelector = v
	return s
}

type AckNodePoolNodeSelector struct {
	Labels []*AckNodePoolNodeSelectorLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Taints []*AckNodePoolNodeSelectorTaints `json:"Taints,omitempty" xml:"Taints,omitempty" type:"Repeated"`
}

func (s AckNodePoolNodeSelector) String() string {
	return tea.Prettify(s)
}

func (s AckNodePoolNodeSelector) GoString() string {
	return s.String()
}

func (s *AckNodePoolNodeSelector) SetLabels(v []*AckNodePoolNodeSelectorLabels) *AckNodePoolNodeSelector {
	s.Labels = v
	return s
}

func (s *AckNodePoolNodeSelector) SetTaints(v []*AckNodePoolNodeSelectorTaints) *AckNodePoolNodeSelector {
	s.Taints = v
	return s
}

type AckNodePoolNodeSelectorLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AckNodePoolNodeSelectorLabels) String() string {
	return tea.Prettify(s)
}

func (s AckNodePoolNodeSelectorLabels) GoString() string {
	return s.String()
}

func (s *AckNodePoolNodeSelectorLabels) SetKey(v string) *AckNodePoolNodeSelectorLabels {
	s.Key = &v
	return s
}

func (s *AckNodePoolNodeSelectorLabels) SetValue(v string) *AckNodePoolNodeSelectorLabels {
	s.Value = &v
	return s
}

type AckNodePoolNodeSelectorTaints struct {
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	Key    *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value  *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AckNodePoolNodeSelectorTaints) String() string {
	return tea.Prettify(s)
}

func (s AckNodePoolNodeSelectorTaints) GoString() string {
	return s.String()
}

func (s *AckNodePoolNodeSelectorTaints) SetEffect(v string) *AckNodePoolNodeSelectorTaints {
	s.Effect = &v
	return s
}

func (s *AckNodePoolNodeSelectorTaints) SetKey(v string) *AckNodePoolNodeSelectorTaints {
	s.Key = &v
	return s
}

func (s *AckNodePoolNodeSelectorTaints) SetValue(v string) *AckNodePoolNodeSelectorTaints {
	s.Value = &v
	return s
}

type AckNodeSelector struct {
	Labels []*AckNodeSelectorLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Taints []*AckNodeSelectorTaints `json:"Taints,omitempty" xml:"Taints,omitempty" type:"Repeated"`
}

func (s AckNodeSelector) String() string {
	return tea.Prettify(s)
}

func (s AckNodeSelector) GoString() string {
	return s.String()
}

func (s *AckNodeSelector) SetLabels(v []*AckNodeSelectorLabels) *AckNodeSelector {
	s.Labels = v
	return s
}

func (s *AckNodeSelector) SetTaints(v []*AckNodeSelectorTaints) *AckNodeSelector {
	s.Taints = v
	return s
}

type AckNodeSelectorLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AckNodeSelectorLabels) String() string {
	return tea.Prettify(s)
}

func (s AckNodeSelectorLabels) GoString() string {
	return s.String()
}

func (s *AckNodeSelectorLabels) SetKey(v string) *AckNodeSelectorLabels {
	s.Key = &v
	return s
}

func (s *AckNodeSelectorLabels) SetValue(v string) *AckNodeSelectorLabels {
	s.Value = &v
	return s
}

type AckNodeSelectorTaints struct {
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	Key    *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value  *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AckNodeSelectorTaints) String() string {
	return tea.Prettify(s)
}

func (s AckNodeSelectorTaints) GoString() string {
	return s.String()
}

func (s *AckNodeSelectorTaints) SetEffect(v string) *AckNodeSelectorTaints {
	s.Effect = &v
	return s
}

func (s *AckNodeSelectorTaints) SetKey(v string) *AckNodeSelectorTaints {
	s.Key = &v
	return s
}

func (s *AckNodeSelectorTaints) SetValue(v string) *AckNodeSelectorTaints {
	s.Value = &v
	return s
}

type Application struct {
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
}

func (s Application) String() string {
	return tea.Prettify(s)
}

func (s Application) GoString() string {
	return s.String()
}

func (s *Application) SetApplicationName(v string) *Application {
	s.ApplicationName = &v
	return s
}

type ApplicationConfig struct {
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	ConfigFileName  *string `json:"ConfigFileName,omitempty" xml:"ConfigFileName,omitempty"`
	ConfigItemKey   *string `json:"ConfigItemKey,omitempty" xml:"ConfigItemKey,omitempty"`
	ConfigItemValue *string `json:"ConfigItemValue,omitempty" xml:"ConfigItemValue,omitempty"`
	ConfigScope     *string `json:"ConfigScope,omitempty" xml:"ConfigScope,omitempty"`
	NodeGroupId     *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	NodeGroupName   *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
}

func (s ApplicationConfig) String() string {
	return tea.Prettify(s)
}

func (s ApplicationConfig) GoString() string {
	return s.String()
}

func (s *ApplicationConfig) SetApplicationName(v string) *ApplicationConfig {
	s.ApplicationName = &v
	return s
}

func (s *ApplicationConfig) SetConfigFileName(v string) *ApplicationConfig {
	s.ConfigFileName = &v
	return s
}

func (s *ApplicationConfig) SetConfigItemKey(v string) *ApplicationConfig {
	s.ConfigItemKey = &v
	return s
}

func (s *ApplicationConfig) SetConfigItemValue(v string) *ApplicationConfig {
	s.ConfigItemValue = &v
	return s
}

func (s *ApplicationConfig) SetConfigScope(v string) *ApplicationConfig {
	s.ConfigScope = &v
	return s
}

func (s *ApplicationConfig) SetNodeGroupId(v string) *ApplicationConfig {
	s.NodeGroupId = &v
	return s
}

func (s *ApplicationConfig) SetNodeGroupName(v string) *ApplicationConfig {
	s.NodeGroupName = &v
	return s
}

type ApplicationConfigParam struct {
	ConfigAction          *string `json:"ConfigAction,omitempty" xml:"ConfigAction,omitempty"`
	ConfigFileName        *string `json:"ConfigFileName,omitempty" xml:"ConfigFileName,omitempty"`
	ConfigItemDescription *string `json:"ConfigItemDescription,omitempty" xml:"ConfigItemDescription,omitempty"`
	ConfigItemKey         *string `json:"ConfigItemKey,omitempty" xml:"ConfigItemKey,omitempty"`
	ConfigItemValue       *string `json:"ConfigItemValue,omitempty" xml:"ConfigItemValue,omitempty"`
	ConfigScope           *string `json:"ConfigScope,omitempty" xml:"ConfigScope,omitempty"`
	EffectiveActions      *string `json:"EffectiveActions,omitempty" xml:"EffectiveActions,omitempty"`
	EffectiveType         *string `json:"EffectiveType,omitempty" xml:"EffectiveType,omitempty"`
	NodeGroupId           *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	NodeId                *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ApplicationConfigParam) String() string {
	return tea.Prettify(s)
}

func (s ApplicationConfigParam) GoString() string {
	return s.String()
}

func (s *ApplicationConfigParam) SetConfigAction(v string) *ApplicationConfigParam {
	s.ConfigAction = &v
	return s
}

func (s *ApplicationConfigParam) SetConfigFileName(v string) *ApplicationConfigParam {
	s.ConfigFileName = &v
	return s
}

func (s *ApplicationConfigParam) SetConfigItemDescription(v string) *ApplicationConfigParam {
	s.ConfigItemDescription = &v
	return s
}

func (s *ApplicationConfigParam) SetConfigItemKey(v string) *ApplicationConfigParam {
	s.ConfigItemKey = &v
	return s
}

func (s *ApplicationConfigParam) SetConfigItemValue(v string) *ApplicationConfigParam {
	s.ConfigItemValue = &v
	return s
}

func (s *ApplicationConfigParam) SetConfigScope(v string) *ApplicationConfigParam {
	s.ConfigScope = &v
	return s
}

func (s *ApplicationConfigParam) SetEffectiveActions(v string) *ApplicationConfigParam {
	s.EffectiveActions = &v
	return s
}

func (s *ApplicationConfigParam) SetEffectiveType(v string) *ApplicationConfigParam {
	s.EffectiveType = &v
	return s
}

func (s *ApplicationConfigParam) SetNodeGroupId(v string) *ApplicationConfigParam {
	s.NodeGroupId = &v
	return s
}

func (s *ApplicationConfigParam) SetNodeId(v string) *ApplicationConfigParam {
	s.NodeId = &v
	return s
}

type AutoRenewInstance struct {
	AutoRenew             *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewDuration     *int32  `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	AutoRenewDurationUnit *string `json:"AutoRenewDurationUnit,omitempty" xml:"AutoRenewDurationUnit,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s AutoRenewInstance) String() string {
	return tea.Prettify(s)
}

func (s AutoRenewInstance) GoString() string {
	return s.String()
}

func (s *AutoRenewInstance) SetAutoRenew(v bool) *AutoRenewInstance {
	s.AutoRenew = &v
	return s
}

func (s *AutoRenewInstance) SetAutoRenewDuration(v int32) *AutoRenewInstance {
	s.AutoRenewDuration = &v
	return s
}

func (s *AutoRenewInstance) SetAutoRenewDurationUnit(v string) *AutoRenewInstance {
	s.AutoRenewDurationUnit = &v
	return s
}

func (s *AutoRenewInstance) SetInstanceId(v string) *AutoRenewInstance {
	s.InstanceId = &v
	return s
}

type AutoRenewInstanceParam struct {
	AutoRenew             *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewDuration     *string `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	AutoRenewDurationUnit *string `json:"AutoRenewDurationUnit,omitempty" xml:"AutoRenewDurationUnit,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s AutoRenewInstanceParam) String() string {
	return tea.Prettify(s)
}

func (s AutoRenewInstanceParam) GoString() string {
	return s.String()
}

func (s *AutoRenewInstanceParam) SetAutoRenew(v string) *AutoRenewInstanceParam {
	s.AutoRenew = &v
	return s
}

func (s *AutoRenewInstanceParam) SetAutoRenewDuration(v string) *AutoRenewInstanceParam {
	s.AutoRenewDuration = &v
	return s
}

func (s *AutoRenewInstanceParam) SetAutoRenewDurationUnit(v string) *AutoRenewInstanceParam {
	s.AutoRenewDurationUnit = &v
	return s
}

func (s *AutoRenewInstanceParam) SetInstanceId(v string) *AutoRenewInstanceParam {
	s.InstanceId = &v
	return s
}

type ByLoadScalingRule struct {
	ComparisonOperator *string  `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	CoolDownInterval   *int32   `json:"CoolDownInterval,omitempty" xml:"CoolDownInterval,omitempty"`
	EvaluationCount    *int32   `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	MetricName         *string  `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Statistics         *string  `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *float64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	TimeWindow         *int32   `json:"TimeWindow,omitempty" xml:"TimeWindow,omitempty"`
}

func (s ByLoadScalingRule) String() string {
	return tea.Prettify(s)
}

func (s ByLoadScalingRule) GoString() string {
	return s.String()
}

func (s *ByLoadScalingRule) SetComparisonOperator(v string) *ByLoadScalingRule {
	s.ComparisonOperator = &v
	return s
}

func (s *ByLoadScalingRule) SetCoolDownInterval(v int32) *ByLoadScalingRule {
	s.CoolDownInterval = &v
	return s
}

func (s *ByLoadScalingRule) SetEvaluationCount(v int32) *ByLoadScalingRule {
	s.EvaluationCount = &v
	return s
}

func (s *ByLoadScalingRule) SetMetricName(v string) *ByLoadScalingRule {
	s.MetricName = &v
	return s
}

func (s *ByLoadScalingRule) SetStatistics(v string) *ByLoadScalingRule {
	s.Statistics = &v
	return s
}

func (s *ByLoadScalingRule) SetThreshold(v float64) *ByLoadScalingRule {
	s.Threshold = &v
	return s
}

func (s *ByLoadScalingRule) SetTimeWindow(v int32) *ByLoadScalingRule {
	s.TimeWindow = &v
	return s
}

type ByLoadScalingRuleSpec struct {
	ComparisonOperator *string  `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	EvaluationCount    *int32   `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	MetricName         *string  `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Statistics         *string  `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *float64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	TimeWindow         *int32   `json:"TimeWindow,omitempty" xml:"TimeWindow,omitempty"`
}

func (s ByLoadScalingRuleSpec) String() string {
	return tea.Prettify(s)
}

func (s ByLoadScalingRuleSpec) GoString() string {
	return s.String()
}

func (s *ByLoadScalingRuleSpec) SetComparisonOperator(v string) *ByLoadScalingRuleSpec {
	s.ComparisonOperator = &v
	return s
}

func (s *ByLoadScalingRuleSpec) SetEvaluationCount(v int32) *ByLoadScalingRuleSpec {
	s.EvaluationCount = &v
	return s
}

func (s *ByLoadScalingRuleSpec) SetMetricName(v string) *ByLoadScalingRuleSpec {
	s.MetricName = &v
	return s
}

func (s *ByLoadScalingRuleSpec) SetStatistics(v string) *ByLoadScalingRuleSpec {
	s.Statistics = &v
	return s
}

func (s *ByLoadScalingRuleSpec) SetThreshold(v float64) *ByLoadScalingRuleSpec {
	s.Threshold = &v
	return s
}

func (s *ByLoadScalingRuleSpec) SetTimeWindow(v int32) *ByLoadScalingRuleSpec {
	s.TimeWindow = &v
	return s
}

type ByTimeScalingRule struct {
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LaunchTime      *int64  `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	RecurrenceType  *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
}

func (s ByTimeScalingRule) String() string {
	return tea.Prettify(s)
}

func (s ByTimeScalingRule) GoString() string {
	return s.String()
}

func (s *ByTimeScalingRule) SetEndTime(v int64) *ByTimeScalingRule {
	s.EndTime = &v
	return s
}

func (s *ByTimeScalingRule) SetLaunchTime(v int64) *ByTimeScalingRule {
	s.LaunchTime = &v
	return s
}

func (s *ByTimeScalingRule) SetRecurrenceType(v string) *ByTimeScalingRule {
	s.RecurrenceType = &v
	return s
}

func (s *ByTimeScalingRule) SetRecurrenceValue(v string) *ByTimeScalingRule {
	s.RecurrenceValue = &v
	return s
}

type ByTimeScalingRuleSpec struct {
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LaunchTime      *int64  `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	RecurrenceType  *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
}

func (s ByTimeScalingRuleSpec) String() string {
	return tea.Prettify(s)
}

func (s ByTimeScalingRuleSpec) GoString() string {
	return s.String()
}

func (s *ByTimeScalingRuleSpec) SetEndTime(v int64) *ByTimeScalingRuleSpec {
	s.EndTime = &v
	return s
}

func (s *ByTimeScalingRuleSpec) SetLaunchTime(v int64) *ByTimeScalingRuleSpec {
	s.LaunchTime = &v
	return s
}

func (s *ByTimeScalingRuleSpec) SetRecurrenceType(v string) *ByTimeScalingRuleSpec {
	s.RecurrenceType = &v
	return s
}

func (s *ByTimeScalingRuleSpec) SetRecurrenceValue(v string) *ByTimeScalingRuleSpec {
	s.RecurrenceValue = &v
	return s
}

type ClickhouseConf struct {
	InitialReplica *int32  `json:"InitialReplica,omitempty" xml:"InitialReplica,omitempty"`
	InitialShard   *int32  `json:"InitialShard,omitempty" xml:"InitialShard,omitempty"`
	NewNodeCount   *int32  `json:"NewNodeCount,omitempty" xml:"NewNodeCount,omitempty"`
	ResizeType     *string `json:"ResizeType,omitempty" xml:"ResizeType,omitempty"`
}

func (s ClickhouseConf) String() string {
	return tea.Prettify(s)
}

func (s ClickhouseConf) GoString() string {
	return s.String()
}

func (s *ClickhouseConf) SetInitialReplica(v int32) *ClickhouseConf {
	s.InitialReplica = &v
	return s
}

func (s *ClickhouseConf) SetInitialShard(v int32) *ClickhouseConf {
	s.InitialShard = &v
	return s
}

func (s *ClickhouseConf) SetNewNodeCount(v int32) *ClickhouseConf {
	s.NewNodeCount = &v
	return s
}

func (s *ClickhouseConf) SetResizeType(v string) *ClickhouseConf {
	s.ResizeType = &v
	return s
}

type Cluster struct {
	ClusterId          *string                   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName        *string                   `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterState       *string                   `json:"ClusterState,omitempty" xml:"ClusterState,omitempty"`
	ClusterType        *string                   `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	CreateTime         *int64                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeployMode         *string                   `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	EmrDefaultRole     *string                   `json:"EmrDefaultRole,omitempty" xml:"EmrDefaultRole,omitempty"`
	EndTime            *int64                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExpireTime         *int64                    `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	NodeAttributes     *NodeAttributes           `json:"NodeAttributes,omitempty" xml:"NodeAttributes,omitempty"`
	PaymentType        *string                   `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	ReadyTime          *int64                    `json:"ReadyTime,omitempty" xml:"ReadyTime,omitempty"`
	RegionId           *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReleaseVersion     *string                   `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	ResourceGroupId    *string                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityMode       *string                   `json:"SecurityMode,omitempty" xml:"SecurityMode,omitempty"`
	StateChangeReason  *ClusterStateChangeReason `json:"StateChangeReason,omitempty" xml:"StateChangeReason,omitempty"`
	SubscriptionConfig *SubscriptionConfig       `json:"SubscriptionConfig,omitempty" xml:"SubscriptionConfig,omitempty"`
	Tags               []*Tag                    `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	VpcId              *string                   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s Cluster) String() string {
	return tea.Prettify(s)
}

func (s Cluster) GoString() string {
	return s.String()
}

func (s *Cluster) SetClusterId(v string) *Cluster {
	s.ClusterId = &v
	return s
}

func (s *Cluster) SetClusterName(v string) *Cluster {
	s.ClusterName = &v
	return s
}

func (s *Cluster) SetClusterState(v string) *Cluster {
	s.ClusterState = &v
	return s
}

func (s *Cluster) SetClusterType(v string) *Cluster {
	s.ClusterType = &v
	return s
}

func (s *Cluster) SetCreateTime(v int64) *Cluster {
	s.CreateTime = &v
	return s
}

func (s *Cluster) SetDeployMode(v string) *Cluster {
	s.DeployMode = &v
	return s
}

func (s *Cluster) SetEmrDefaultRole(v string) *Cluster {
	s.EmrDefaultRole = &v
	return s
}

func (s *Cluster) SetEndTime(v int64) *Cluster {
	s.EndTime = &v
	return s
}

func (s *Cluster) SetExpireTime(v int64) *Cluster {
	s.ExpireTime = &v
	return s
}

func (s *Cluster) SetNodeAttributes(v *NodeAttributes) *Cluster {
	s.NodeAttributes = v
	return s
}

func (s *Cluster) SetPaymentType(v string) *Cluster {
	s.PaymentType = &v
	return s
}

func (s *Cluster) SetReadyTime(v int64) *Cluster {
	s.ReadyTime = &v
	return s
}

func (s *Cluster) SetRegionId(v string) *Cluster {
	s.RegionId = &v
	return s
}

func (s *Cluster) SetReleaseVersion(v string) *Cluster {
	s.ReleaseVersion = &v
	return s
}

func (s *Cluster) SetResourceGroupId(v string) *Cluster {
	s.ResourceGroupId = &v
	return s
}

func (s *Cluster) SetSecurityMode(v string) *Cluster {
	s.SecurityMode = &v
	return s
}

func (s *Cluster) SetStateChangeReason(v *ClusterStateChangeReason) *Cluster {
	s.StateChangeReason = v
	return s
}

func (s *Cluster) SetSubscriptionConfig(v *SubscriptionConfig) *Cluster {
	s.SubscriptionConfig = v
	return s
}

func (s *Cluster) SetTags(v []*Tag) *Cluster {
	s.Tags = v
	return s
}

func (s *Cluster) SetVpcId(v string) *Cluster {
	s.VpcId = &v
	return s
}

type ClusterScript struct {
	ExecutionFailStrategy *string       `json:"ExecutionFailStrategy,omitempty" xml:"ExecutionFailStrategy,omitempty"`
	ExecutionMoment       *string       `json:"ExecutionMoment,omitempty" xml:"ExecutionMoment,omitempty"`
	NodeSelect            *NodeSelector `json:"NodeSelect,omitempty" xml:"NodeSelect,omitempty"`
	Priority              *int32        `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ScriptArgs            *string       `json:"ScriptArgs,omitempty" xml:"ScriptArgs,omitempty"`
	ScriptName            *string       `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	ScriptPath            *string       `json:"ScriptPath,omitempty" xml:"ScriptPath,omitempty"`
}

func (s ClusterScript) String() string {
	return tea.Prettify(s)
}

func (s ClusterScript) GoString() string {
	return s.String()
}

func (s *ClusterScript) SetExecutionFailStrategy(v string) *ClusterScript {
	s.ExecutionFailStrategy = &v
	return s
}

func (s *ClusterScript) SetExecutionMoment(v string) *ClusterScript {
	s.ExecutionMoment = &v
	return s
}

func (s *ClusterScript) SetNodeSelect(v *NodeSelector) *ClusterScript {
	s.NodeSelect = v
	return s
}

func (s *ClusterScript) SetPriority(v int32) *ClusterScript {
	s.Priority = &v
	return s
}

func (s *ClusterScript) SetScriptArgs(v string) *ClusterScript {
	s.ScriptArgs = &v
	return s
}

func (s *ClusterScript) SetScriptName(v string) *ClusterScript {
	s.ScriptName = &v
	return s
}

func (s *ClusterScript) SetScriptPath(v string) *ClusterScript {
	s.ScriptPath = &v
	return s
}

type ClusterStateChangeReason struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ClusterStateChangeReason) String() string {
	return tea.Prettify(s)
}

func (s ClusterStateChangeReason) GoString() string {
	return s.String()
}

func (s *ClusterStateChangeReason) SetCode(v string) *ClusterStateChangeReason {
	s.Code = &v
	return s
}

func (s *ClusterStateChangeReason) SetMessage(v string) *ClusterStateChangeReason {
	s.Message = &v
	return s
}

type ClusterSummary struct {
	ClusterId         *string                   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName       *string                   `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterState      *string                   `json:"ClusterState,omitempty" xml:"ClusterState,omitempty"`
	ClusterType       *string                   `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	CreateTime        *int64                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EmrDefaultRole    *string                   `json:"EmrDefaultRole,omitempty" xml:"EmrDefaultRole,omitempty"`
	EndTime           *int64                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExpireTime        *int64                    `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	PaymentType       *string                   `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	ReadyTime         *int64                    `json:"ReadyTime,omitempty" xml:"ReadyTime,omitempty"`
	ReleaseVersion    *string                   `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	ResourceGroupId   *string                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StateChangeReason *ClusterStateChangeReason `json:"StateChangeReason,omitempty" xml:"StateChangeReason,omitempty"`
	Tags              []*Tag                    `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ClusterSummary) String() string {
	return tea.Prettify(s)
}

func (s ClusterSummary) GoString() string {
	return s.String()
}

func (s *ClusterSummary) SetClusterId(v string) *ClusterSummary {
	s.ClusterId = &v
	return s
}

func (s *ClusterSummary) SetClusterName(v string) *ClusterSummary {
	s.ClusterName = &v
	return s
}

func (s *ClusterSummary) SetClusterState(v string) *ClusterSummary {
	s.ClusterState = &v
	return s
}

func (s *ClusterSummary) SetClusterType(v string) *ClusterSummary {
	s.ClusterType = &v
	return s
}

func (s *ClusterSummary) SetCreateTime(v int64) *ClusterSummary {
	s.CreateTime = &v
	return s
}

func (s *ClusterSummary) SetEmrDefaultRole(v string) *ClusterSummary {
	s.EmrDefaultRole = &v
	return s
}

func (s *ClusterSummary) SetEndTime(v int64) *ClusterSummary {
	s.EndTime = &v
	return s
}

func (s *ClusterSummary) SetExpireTime(v int64) *ClusterSummary {
	s.ExpireTime = &v
	return s
}

func (s *ClusterSummary) SetPaymentType(v string) *ClusterSummary {
	s.PaymentType = &v
	return s
}

func (s *ClusterSummary) SetReadyTime(v int64) *ClusterSummary {
	s.ReadyTime = &v
	return s
}

func (s *ClusterSummary) SetReleaseVersion(v string) *ClusterSummary {
	s.ReleaseVersion = &v
	return s
}

func (s *ClusterSummary) SetResourceGroupId(v string) *ClusterSummary {
	s.ResourceGroupId = &v
	return s
}

func (s *ClusterSummary) SetStateChangeReason(v *ClusterStateChangeReason) *ClusterSummary {
	s.StateChangeReason = v
	return s
}

func (s *ClusterSummary) SetTags(v []*Tag) *ClusterSummary {
	s.Tags = v
	return s
}

type ComponentInstanceSelector struct {
	ActionScope        *string                                        `json:"ActionScope,omitempty" xml:"ActionScope,omitempty"`
	ApplicationNames   []*string                                      `json:"ApplicationNames,omitempty" xml:"ApplicationNames,omitempty" type:"Repeated"`
	ComponentInstances []*ComponentInstanceSelectorComponentInstances `json:"ComponentInstances,omitempty" xml:"ComponentInstances,omitempty" type:"Repeated"`
	Components         []*ComponentInstanceSelectorComponents         `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	NodeGroupIds       []*string                                      `json:"NodeGroupIds,omitempty" xml:"NodeGroupIds,omitempty" type:"Repeated"`
	NodeIds            []*string                                      `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	RunActionScope     *string                                        `json:"RunActionScope,omitempty" xml:"RunActionScope,omitempty"`
}

func (s ComponentInstanceSelector) String() string {
	return tea.Prettify(s)
}

func (s ComponentInstanceSelector) GoString() string {
	return s.String()
}

func (s *ComponentInstanceSelector) SetActionScope(v string) *ComponentInstanceSelector {
	s.ActionScope = &v
	return s
}

func (s *ComponentInstanceSelector) SetApplicationNames(v []*string) *ComponentInstanceSelector {
	s.ApplicationNames = v
	return s
}

func (s *ComponentInstanceSelector) SetComponentInstances(v []*ComponentInstanceSelectorComponentInstances) *ComponentInstanceSelector {
	s.ComponentInstances = v
	return s
}

func (s *ComponentInstanceSelector) SetComponents(v []*ComponentInstanceSelectorComponents) *ComponentInstanceSelector {
	s.Components = v
	return s
}

func (s *ComponentInstanceSelector) SetNodeGroupIds(v []*string) *ComponentInstanceSelector {
	s.NodeGroupIds = v
	return s
}

func (s *ComponentInstanceSelector) SetNodeIds(v []*string) *ComponentInstanceSelector {
	s.NodeIds = v
	return s
}

func (s *ComponentInstanceSelector) SetRunActionScope(v string) *ComponentInstanceSelector {
	s.RunActionScope = &v
	return s
}

type ComponentInstanceSelectorComponentInstances struct {
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	ComponentName   *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	NodeId          *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ComponentInstanceSelectorComponentInstances) String() string {
	return tea.Prettify(s)
}

func (s ComponentInstanceSelectorComponentInstances) GoString() string {
	return s.String()
}

func (s *ComponentInstanceSelectorComponentInstances) SetApplicationName(v string) *ComponentInstanceSelectorComponentInstances {
	s.ApplicationName = &v
	return s
}

func (s *ComponentInstanceSelectorComponentInstances) SetComponentName(v string) *ComponentInstanceSelectorComponentInstances {
	s.ComponentName = &v
	return s
}

func (s *ComponentInstanceSelectorComponentInstances) SetNodeId(v string) *ComponentInstanceSelectorComponentInstances {
	s.NodeId = &v
	return s
}

type ComponentInstanceSelectorComponents struct {
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	ComponentName   *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
}

func (s ComponentInstanceSelectorComponents) String() string {
	return tea.Prettify(s)
}

func (s ComponentInstanceSelectorComponents) GoString() string {
	return s.String()
}

func (s *ComponentInstanceSelectorComponents) SetApplicationName(v string) *ComponentInstanceSelectorComponents {
	s.ApplicationName = &v
	return s
}

func (s *ComponentInstanceSelectorComponents) SetComponentName(v string) *ComponentInstanceSelectorComponents {
	s.ComponentName = &v
	return s
}

type ComponentLayout struct {
	ApplicationName *string                      `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	ComponentName   *string                      `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	NodeSelector    *ComponentLayoutNodeSelector `json:"NodeSelector,omitempty" xml:"NodeSelector,omitempty" type:"Struct"`
}

func (s ComponentLayout) String() string {
	return tea.Prettify(s)
}

func (s ComponentLayout) GoString() string {
	return s.String()
}

func (s *ComponentLayout) SetApplicationName(v string) *ComponentLayout {
	s.ApplicationName = &v
	return s
}

func (s *ComponentLayout) SetComponentName(v string) *ComponentLayout {
	s.ComponentName = &v
	return s
}

func (s *ComponentLayout) SetNodeSelector(v *ComponentLayoutNodeSelector) *ComponentLayout {
	s.NodeSelector = v
	return s
}

type ComponentLayoutNodeSelector struct {
	NodeEndIndex   *int32    `json:"NodeEndIndex,omitempty" xml:"NodeEndIndex,omitempty"`
	NodeGroupId    *string   `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	NodeGroupIndex *int32    `json:"NodeGroupIndex,omitempty" xml:"NodeGroupIndex,omitempty"`
	NodeGroupName  *string   `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	NodeGroupTypes []*string `json:"NodeGroupTypes,omitempty" xml:"NodeGroupTypes,omitempty" type:"Repeated"`
	NodeNames      []*string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty" type:"Repeated"`
	NodeSelectType *string   `json:"NodeSelectType,omitempty" xml:"NodeSelectType,omitempty"`
	NodeStartIndex *int32    `json:"NodeStartIndex,omitempty" xml:"NodeStartIndex,omitempty"`
}

func (s ComponentLayoutNodeSelector) String() string {
	return tea.Prettify(s)
}

func (s ComponentLayoutNodeSelector) GoString() string {
	return s.String()
}

func (s *ComponentLayoutNodeSelector) SetNodeEndIndex(v int32) *ComponentLayoutNodeSelector {
	s.NodeEndIndex = &v
	return s
}

func (s *ComponentLayoutNodeSelector) SetNodeGroupId(v string) *ComponentLayoutNodeSelector {
	s.NodeGroupId = &v
	return s
}

func (s *ComponentLayoutNodeSelector) SetNodeGroupIndex(v int32) *ComponentLayoutNodeSelector {
	s.NodeGroupIndex = &v
	return s
}

func (s *ComponentLayoutNodeSelector) SetNodeGroupName(v string) *ComponentLayoutNodeSelector {
	s.NodeGroupName = &v
	return s
}

func (s *ComponentLayoutNodeSelector) SetNodeGroupTypes(v []*string) *ComponentLayoutNodeSelector {
	s.NodeGroupTypes = v
	return s
}

func (s *ComponentLayoutNodeSelector) SetNodeNames(v []*string) *ComponentLayoutNodeSelector {
	s.NodeNames = v
	return s
}

func (s *ComponentLayoutNodeSelector) SetNodeSelectType(v string) *ComponentLayoutNodeSelector {
	s.NodeSelectType = &v
	return s
}

func (s *ComponentLayoutNodeSelector) SetNodeStartIndex(v int32) *ComponentLayoutNodeSelector {
	s.NodeStartIndex = &v
	return s
}

type ConvertNodeGroup struct {
	NodeGroupId         *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	PaymentDuration     *int32  `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	PaymentDurationUnit *string `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	PaymentType         *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
}

func (s ConvertNodeGroup) String() string {
	return tea.Prettify(s)
}

func (s ConvertNodeGroup) GoString() string {
	return s.String()
}

func (s *ConvertNodeGroup) SetNodeGroupId(v string) *ConvertNodeGroup {
	s.NodeGroupId = &v
	return s
}

func (s *ConvertNodeGroup) SetPaymentDuration(v int32) *ConvertNodeGroup {
	s.PaymentDuration = &v
	return s
}

func (s *ConvertNodeGroup) SetPaymentDurationUnit(v string) *ConvertNodeGroup {
	s.PaymentDurationUnit = &v
	return s
}

func (s *ConvertNodeGroup) SetPaymentType(v string) *ConvertNodeGroup {
	s.PaymentType = &v
	return s
}

type ConvertNodeGroupParam struct {
	NodeGroupId         *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	PaymentDuration     *int64  `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	PaymentDurationUnit *string `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	PaymentType         *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
}

func (s ConvertNodeGroupParam) String() string {
	return tea.Prettify(s)
}

func (s ConvertNodeGroupParam) GoString() string {
	return s.String()
}

func (s *ConvertNodeGroupParam) SetNodeGroupId(v string) *ConvertNodeGroupParam {
	s.NodeGroupId = &v
	return s
}

func (s *ConvertNodeGroupParam) SetPaymentDuration(v int64) *ConvertNodeGroupParam {
	s.PaymentDuration = &v
	return s
}

func (s *ConvertNodeGroupParam) SetPaymentDurationUnit(v string) *ConvertNodeGroupParam {
	s.PaymentDurationUnit = &v
	return s
}

func (s *ConvertNodeGroupParam) SetPaymentType(v string) *ConvertNodeGroupParam {
	s.PaymentType = &v
	return s
}

type CostOptimizedConfig struct {
	OnDemandBaseCapacity                *int32 `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	OnDemandPercentageAboveBaseCapacity *int32 `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	SpotInstancePools                   *int32 `json:"SpotInstancePools,omitempty" xml:"SpotInstancePools,omitempty"`
}

func (s CostOptimizedConfig) String() string {
	return tea.Prettify(s)
}

func (s CostOptimizedConfig) GoString() string {
	return s.String()
}

func (s *CostOptimizedConfig) SetOnDemandBaseCapacity(v int32) *CostOptimizedConfig {
	s.OnDemandBaseCapacity = &v
	return s
}

func (s *CostOptimizedConfig) SetOnDemandPercentageAboveBaseCapacity(v int32) *CostOptimizedConfig {
	s.OnDemandPercentageAboveBaseCapacity = &v
	return s
}

func (s *CostOptimizedConfig) SetSpotInstancePools(v int32) *CostOptimizedConfig {
	s.SpotInstancePools = &v
	return s
}

type CreateNodeGroupParam struct {
	AutoRenew             *bool            `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewDuration     *int32           `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	AutoRenewDurationUnit *string          `json:"AutoRenewDurationUnit,omitempty" xml:"AutoRenewDurationUnit,omitempty"`
	DataDisks             []*DiskInfo      `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	InstanceTypes         []*string        `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	NodeCount             *int32           `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	NodeGroupName         *string          `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	NodeGroupType         *string          `json:"NodeGroupType,omitempty" xml:"NodeGroupType,omitempty"`
	NodeKeyPairName       *string          `json:"NodeKeyPairName,omitempty" xml:"NodeKeyPairName,omitempty"`
	NodeRamRole           *string          `json:"NodeRamRole,omitempty" xml:"NodeRamRole,omitempty"`
	NodeRootPassword      *string          `json:"NodeRootPassword,omitempty" xml:"NodeRootPassword,omitempty"`
	PaymentDuration       *int32           `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	PaymentDurationUnit   *string          `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	PaymentType           *string          `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	SecurityGroupId       *string          `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SpotStrategy          *string          `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDisk            *SystemDiskParam `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
	VSwitchIds            []*string        `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	WithPublicIp          *bool            `json:"WithPublicIp,omitempty" xml:"WithPublicIp,omitempty"`
	ZoneId                *string          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateNodeGroupParam) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeGroupParam) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupParam) SetAutoRenew(v bool) *CreateNodeGroupParam {
	s.AutoRenew = &v
	return s
}

func (s *CreateNodeGroupParam) SetAutoRenewDuration(v int32) *CreateNodeGroupParam {
	s.AutoRenewDuration = &v
	return s
}

func (s *CreateNodeGroupParam) SetAutoRenewDurationUnit(v string) *CreateNodeGroupParam {
	s.AutoRenewDurationUnit = &v
	return s
}

func (s *CreateNodeGroupParam) SetDataDisks(v []*DiskInfo) *CreateNodeGroupParam {
	s.DataDisks = v
	return s
}

func (s *CreateNodeGroupParam) SetInstanceTypes(v []*string) *CreateNodeGroupParam {
	s.InstanceTypes = v
	return s
}

func (s *CreateNodeGroupParam) SetNodeCount(v int32) *CreateNodeGroupParam {
	s.NodeCount = &v
	return s
}

func (s *CreateNodeGroupParam) SetNodeGroupName(v string) *CreateNodeGroupParam {
	s.NodeGroupName = &v
	return s
}

func (s *CreateNodeGroupParam) SetNodeGroupType(v string) *CreateNodeGroupParam {
	s.NodeGroupType = &v
	return s
}

func (s *CreateNodeGroupParam) SetNodeKeyPairName(v string) *CreateNodeGroupParam {
	s.NodeKeyPairName = &v
	return s
}

func (s *CreateNodeGroupParam) SetNodeRamRole(v string) *CreateNodeGroupParam {
	s.NodeRamRole = &v
	return s
}

func (s *CreateNodeGroupParam) SetNodeRootPassword(v string) *CreateNodeGroupParam {
	s.NodeRootPassword = &v
	return s
}

func (s *CreateNodeGroupParam) SetPaymentDuration(v int32) *CreateNodeGroupParam {
	s.PaymentDuration = &v
	return s
}

func (s *CreateNodeGroupParam) SetPaymentDurationUnit(v string) *CreateNodeGroupParam {
	s.PaymentDurationUnit = &v
	return s
}

func (s *CreateNodeGroupParam) SetPaymentType(v string) *CreateNodeGroupParam {
	s.PaymentType = &v
	return s
}

func (s *CreateNodeGroupParam) SetSecurityGroupId(v string) *CreateNodeGroupParam {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateNodeGroupParam) SetSpotStrategy(v string) *CreateNodeGroupParam {
	s.SpotStrategy = &v
	return s
}

func (s *CreateNodeGroupParam) SetSystemDisk(v *SystemDiskParam) *CreateNodeGroupParam {
	s.SystemDisk = v
	return s
}

func (s *CreateNodeGroupParam) SetVSwitchIds(v []*string) *CreateNodeGroupParam {
	s.VSwitchIds = v
	return s
}

func (s *CreateNodeGroupParam) SetWithPublicIp(v bool) *CreateNodeGroupParam {
	s.WithPublicIp = &v
	return s
}

func (s *CreateNodeGroupParam) SetZoneId(v string) *CreateNodeGroupParam {
	s.ZoneId = &v
	return s
}

type DataDisk struct {
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Count            *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Size             *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DataDisk) String() string {
	return tea.Prettify(s)
}

func (s DataDisk) GoString() string {
	return s.String()
}

func (s *DataDisk) SetCategory(v string) *DataDisk {
	s.Category = &v
	return s
}

func (s *DataDisk) SetCount(v int32) *DataDisk {
	s.Count = &v
	return s
}

func (s *DataDisk) SetPerformanceLevel(v string) *DataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DataDisk) SetSize(v int32) *DataDisk {
	s.Size = &v
	return s
}

type DecreaseNodeGroupParam struct {
	NodeGroupId        *string   `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	ReleaseInstanceIds []*string `json:"ReleaseInstanceIds,omitempty" xml:"ReleaseInstanceIds,omitempty" type:"Repeated"`
}

func (s DecreaseNodeGroupParam) String() string {
	return tea.Prettify(s)
}

func (s DecreaseNodeGroupParam) GoString() string {
	return s.String()
}

func (s *DecreaseNodeGroupParam) SetNodeGroupId(v string) *DecreaseNodeGroupParam {
	s.NodeGroupId = &v
	return s
}

func (s *DecreaseNodeGroupParam) SetReleaseInstanceIds(v []*string) *DecreaseNodeGroupParam {
	s.ReleaseInstanceIds = v
	return s
}

type DeploymentLayout struct {
	ApplicationName *string       `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	ComponentName   *string       `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	NodeSelector    *NodeSelector `json:"NodeSelector,omitempty" xml:"NodeSelector,omitempty"`
}

func (s DeploymentLayout) String() string {
	return tea.Prettify(s)
}

func (s DeploymentLayout) GoString() string {
	return s.String()
}

func (s *DeploymentLayout) SetApplicationName(v string) *DeploymentLayout {
	s.ApplicationName = &v
	return s
}

func (s *DeploymentLayout) SetComponentName(v string) *DeploymentLayout {
	s.ComponentName = &v
	return s
}

func (s *DeploymentLayout) SetNodeSelector(v *NodeSelector) *DeploymentLayout {
	s.NodeSelector = v
	return s
}

type Disk struct {
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Count            *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Size             *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s Disk) String() string {
	return tea.Prettify(s)
}

func (s Disk) GoString() string {
	return s.String()
}

func (s *Disk) SetCategory(v string) *Disk {
	s.Category = &v
	return s
}

func (s *Disk) SetCount(v int32) *Disk {
	s.Count = &v
	return s
}

func (s *Disk) SetPerformanceLevel(v string) *Disk {
	s.PerformanceLevel = &v
	return s
}

func (s *Disk) SetSize(v int32) *Disk {
	s.Size = &v
	return s
}

type DiskInfo struct {
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Count            *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Size             *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DiskInfo) String() string {
	return tea.Prettify(s)
}

func (s DiskInfo) GoString() string {
	return s.String()
}

func (s *DiskInfo) SetCategory(v string) *DiskInfo {
	s.Category = &v
	return s
}

func (s *DiskInfo) SetCount(v int32) *DiskInfo {
	s.Count = &v
	return s
}

func (s *DiskInfo) SetPerformanceLevel(v string) *DiskInfo {
	s.PerformanceLevel = &v
	return s
}

func (s *DiskInfo) SetSize(v int32) *DiskInfo {
	s.Size = &v
	return s
}

type DiskSize struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Size     *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DiskSize) String() string {
	return tea.Prettify(s)
}

func (s DiskSize) GoString() string {
	return s.String()
}

func (s *DiskSize) SetCategory(v string) *DiskSize {
	s.Category = &v
	return s
}

func (s *DiskSize) SetSize(v int32) *DiskSize {
	s.Size = &v
	return s
}

type FailedReason struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorMsg     *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Service      *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s FailedReason) String() string {
	return tea.Prettify(s)
}

func (s FailedReason) GoString() string {
	return s.String()
}

func (s *FailedReason) SetErrorCode(v string) *FailedReason {
	s.ErrorCode = &v
	return s
}

func (s *FailedReason) SetErrorMessage(v string) *FailedReason {
	s.ErrorMessage = &v
	return s
}

func (s *FailedReason) SetErrorMsg(v string) *FailedReason {
	s.ErrorMsg = &v
	return s
}

func (s *FailedReason) SetRequestId(v string) *FailedReason {
	s.RequestId = &v
	return s
}

func (s *FailedReason) SetService(v string) *FailedReason {
	s.Service = &v
	return s
}

type HealthSummary struct {
	BadCount     *int64 `json:"BadCount,omitempty" xml:"BadCount,omitempty"`
	GoodCount    *int64 `json:"GoodCount,omitempty" xml:"GoodCount,omitempty"`
	NoneCount    *int64 `json:"NoneCount,omitempty" xml:"NoneCount,omitempty"`
	StoppedCount *int64 `json:"StoppedCount,omitempty" xml:"StoppedCount,omitempty"`
	TotalCount   *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UnknownCount *int64 `json:"UnknownCount,omitempty" xml:"UnknownCount,omitempty"`
	WarningCount *int64 `json:"WarningCount,omitempty" xml:"WarningCount,omitempty"`
}

func (s HealthSummary) String() string {
	return tea.Prettify(s)
}

func (s HealthSummary) GoString() string {
	return s.String()
}

func (s *HealthSummary) SetBadCount(v int64) *HealthSummary {
	s.BadCount = &v
	return s
}

func (s *HealthSummary) SetGoodCount(v int64) *HealthSummary {
	s.GoodCount = &v
	return s
}

func (s *HealthSummary) SetNoneCount(v int64) *HealthSummary {
	s.NoneCount = &v
	return s
}

func (s *HealthSummary) SetStoppedCount(v int64) *HealthSummary {
	s.StoppedCount = &v
	return s
}

func (s *HealthSummary) SetTotalCount(v int64) *HealthSummary {
	s.TotalCount = &v
	return s
}

func (s *HealthSummary) SetUnknownCount(v int64) *HealthSummary {
	s.UnknownCount = &v
	return s
}

func (s *HealthSummary) SetWarningCount(v int64) *HealthSummary {
	s.WarningCount = &v
	return s
}

type IncreaseNodeGroup struct {
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NodeCount           *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	NodeGroupId         *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	PaymentDuration     *int32  `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	PaymentDurationUnit *string `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	VSwitchId           *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s IncreaseNodeGroup) String() string {
	return tea.Prettify(s)
}

func (s IncreaseNodeGroup) GoString() string {
	return s.String()
}

func (s *IncreaseNodeGroup) SetDescription(v string) *IncreaseNodeGroup {
	s.Description = &v
	return s
}

func (s *IncreaseNodeGroup) SetNodeCount(v int32) *IncreaseNodeGroup {
	s.NodeCount = &v
	return s
}

func (s *IncreaseNodeGroup) SetNodeGroupId(v string) *IncreaseNodeGroup {
	s.NodeGroupId = &v
	return s
}

func (s *IncreaseNodeGroup) SetPaymentDuration(v int32) *IncreaseNodeGroup {
	s.PaymentDuration = &v
	return s
}

func (s *IncreaseNodeGroup) SetPaymentDurationUnit(v string) *IncreaseNodeGroup {
	s.PaymentDurationUnit = &v
	return s
}

func (s *IncreaseNodeGroup) SetVSwitchId(v string) *IncreaseNodeGroup {
	s.VSwitchId = &v
	return s
}

type IncreaseNodeGroupParam struct {
	NodeCount           *int64  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	NodeGroupId         *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	PaymentDuration     *int32  `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	PaymentDurationUnit *string `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	VSwitchId           *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s IncreaseNodeGroupParam) String() string {
	return tea.Prettify(s)
}

func (s IncreaseNodeGroupParam) GoString() string {
	return s.String()
}

func (s *IncreaseNodeGroupParam) SetNodeCount(v int64) *IncreaseNodeGroupParam {
	s.NodeCount = &v
	return s
}

func (s *IncreaseNodeGroupParam) SetNodeGroupId(v string) *IncreaseNodeGroupParam {
	s.NodeGroupId = &v
	return s
}

func (s *IncreaseNodeGroupParam) SetPaymentDuration(v int32) *IncreaseNodeGroupParam {
	s.PaymentDuration = &v
	return s
}

func (s *IncreaseNodeGroupParam) SetPaymentDurationUnit(v string) *IncreaseNodeGroupParam {
	s.PaymentDurationUnit = &v
	return s
}

func (s *IncreaseNodeGroupParam) SetVSwitchId(v string) *IncreaseNodeGroupParam {
	s.VSwitchId = &v
	return s
}

type KeyValue struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s KeyValue) String() string {
	return tea.Prettify(s)
}

func (s KeyValue) GoString() string {
	return s.String()
}

func (s *KeyValue) SetKey(v string) *KeyValue {
	s.Key = &v
	return s
}

func (s *KeyValue) SetValue(v string) *KeyValue {
	s.Value = &v
	return s
}

type MetaStoreConf struct {
	DbPassword *string `json:"DbPassword,omitempty" xml:"DbPassword,omitempty"`
	DbUrl      *string `json:"DbUrl,omitempty" xml:"DbUrl,omitempty"`
	DbUserName *string `json:"DbUserName,omitempty" xml:"DbUserName,omitempty"`
}

func (s MetaStoreConf) String() string {
	return tea.Prettify(s)
}

func (s MetaStoreConf) GoString() string {
	return s.String()
}

func (s *MetaStoreConf) SetDbPassword(v string) *MetaStoreConf {
	s.DbPassword = &v
	return s
}

func (s *MetaStoreConf) SetDbUrl(v string) *MetaStoreConf {
	s.DbUrl = &v
	return s
}

func (s *MetaStoreConf) SetDbUserName(v string) *MetaStoreConf {
	s.DbUserName = &v
	return s
}

type Node struct {
	AutoRenew             *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewDuration     *int32  `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	AutoRenewDurationUnit *string `json:"AutoRenewDurationUnit,omitempty" xml:"AutoRenewDurationUnit,omitempty"`
	ExpireTime            *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceType          *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	NodeGroupId           *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	NodeGroupType         *string `json:"NodeGroupType,omitempty" xml:"NodeGroupType,omitempty"`
	NodeId                *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName              *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeState             *string `json:"NodeState,omitempty" xml:"NodeState,omitempty"`
	PrivateIp             *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	PublicIp              *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	ZoneId                *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s Node) String() string {
	return tea.Prettify(s)
}

func (s Node) GoString() string {
	return s.String()
}

func (s *Node) SetAutoRenew(v bool) *Node {
	s.AutoRenew = &v
	return s
}

func (s *Node) SetAutoRenewDuration(v int32) *Node {
	s.AutoRenewDuration = &v
	return s
}

func (s *Node) SetAutoRenewDurationUnit(v string) *Node {
	s.AutoRenewDurationUnit = &v
	return s
}

func (s *Node) SetExpireTime(v int64) *Node {
	s.ExpireTime = &v
	return s
}

func (s *Node) SetInstanceType(v string) *Node {
	s.InstanceType = &v
	return s
}

func (s *Node) SetNodeGroupId(v string) *Node {
	s.NodeGroupId = &v
	return s
}

func (s *Node) SetNodeGroupType(v string) *Node {
	s.NodeGroupType = &v
	return s
}

func (s *Node) SetNodeId(v string) *Node {
	s.NodeId = &v
	return s
}

func (s *Node) SetNodeName(v string) *Node {
	s.NodeName = &v
	return s
}

func (s *Node) SetNodeState(v string) *Node {
	s.NodeState = &v
	return s
}

func (s *Node) SetPrivateIp(v string) *Node {
	s.PrivateIp = &v
	return s
}

func (s *Node) SetPublicIp(v string) *Node {
	s.PublicIp = &v
	return s
}

func (s *Node) SetZoneId(v string) *Node {
	s.ZoneId = &v
	return s
}

type NodeAttributes struct {
	KeyPairName     *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	RamRole         *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId          *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s NodeAttributes) String() string {
	return tea.Prettify(s)
}

func (s NodeAttributes) GoString() string {
	return s.String()
}

func (s *NodeAttributes) SetKeyPairName(v string) *NodeAttributes {
	s.KeyPairName = &v
	return s
}

func (s *NodeAttributes) SetRamRole(v string) *NodeAttributes {
	s.RamRole = &v
	return s
}

func (s *NodeAttributes) SetSecurityGroupId(v string) *NodeAttributes {
	s.SecurityGroupId = &v
	return s
}

func (s *NodeAttributes) SetVpcId(v string) *NodeAttributes {
	s.VpcId = &v
	return s
}

func (s *NodeAttributes) SetZoneId(v string) *NodeAttributes {
	s.ZoneId = &v
	return s
}

type NodeCountConstraint struct {
	Max    *int32   `json:"Max,omitempty" xml:"Max,omitempty"`
	Min    *int32   `json:"Min,omitempty" xml:"Min,omitempty"`
	Type   *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Values []*int32 `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s NodeCountConstraint) String() string {
	return tea.Prettify(s)
}

func (s NodeCountConstraint) GoString() string {
	return s.String()
}

func (s *NodeCountConstraint) SetMax(v int32) *NodeCountConstraint {
	s.Max = &v
	return s
}

func (s *NodeCountConstraint) SetMin(v int32) *NodeCountConstraint {
	s.Min = &v
	return s
}

func (s *NodeCountConstraint) SetType(v string) *NodeCountConstraint {
	s.Type = &v
	return s
}

func (s *NodeCountConstraint) SetValues(v []*int32) *NodeCountConstraint {
	s.Values = v
	return s
}

type NodeGroup struct {
	AdditionalSecurityGroupIds []*string                   `json:"AdditionalSecurityGroupIds,omitempty" xml:"AdditionalSecurityGroupIds,omitempty" type:"Repeated"`
	CostOptimizedConfig        *CostOptimizedConfig        `json:"CostOptimizedConfig,omitempty" xml:"CostOptimizedConfig,omitempty"`
	DataDisks                  []*DataDisk                 `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	DeploymentSetStrategy      *string                     `json:"DeploymentSetStrategy,omitempty" xml:"DeploymentSetStrategy,omitempty"`
	GracefulShutdown           *bool                       `json:"GracefulShutdown,omitempty" xml:"GracefulShutdown,omitempty"`
	InstanceTypes              []*string                   `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	NodeGroupId                *string                     `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	NodeGroupName              *string                     `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	NodeGroupState             *string                     `json:"NodeGroupState,omitempty" xml:"NodeGroupState,omitempty"`
	NodeGroupType              *string                     `json:"NodeGroupType,omitempty" xml:"NodeGroupType,omitempty"`
	NodeResizeStrategy         *string                     `json:"NodeResizeStrategy,omitempty" xml:"NodeResizeStrategy,omitempty"`
	PaymentType                *string                     `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	RunningNodeCount           *int32                      `json:"RunningNodeCount,omitempty" xml:"RunningNodeCount,omitempty"`
	SpotBidPrices              []*SpotBidPrice             `json:"SpotBidPrices,omitempty" xml:"SpotBidPrices,omitempty" type:"Repeated"`
	SpotInstanceRemedy         *bool                       `json:"SpotInstanceRemedy,omitempty" xml:"SpotInstanceRemedy,omitempty"`
	SpotStrategy               *string                     `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	StateChangeReason          *NodeGroupStateChangeReason `json:"StateChangeReason,omitempty" xml:"StateChangeReason,omitempty"`
	SystemDisk                 *SystemDisk                 `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
	VSwitchIds                 []*string                   `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	WithPublicIp               *bool                       `json:"WithPublicIp,omitempty" xml:"WithPublicIp,omitempty"`
	ZoneId                     *string                     `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s NodeGroup) String() string {
	return tea.Prettify(s)
}

func (s NodeGroup) GoString() string {
	return s.String()
}

func (s *NodeGroup) SetAdditionalSecurityGroupIds(v []*string) *NodeGroup {
	s.AdditionalSecurityGroupIds = v
	return s
}

func (s *NodeGroup) SetCostOptimizedConfig(v *CostOptimizedConfig) *NodeGroup {
	s.CostOptimizedConfig = v
	return s
}

func (s *NodeGroup) SetDataDisks(v []*DataDisk) *NodeGroup {
	s.DataDisks = v
	return s
}

func (s *NodeGroup) SetDeploymentSetStrategy(v string) *NodeGroup {
	s.DeploymentSetStrategy = &v
	return s
}

func (s *NodeGroup) SetGracefulShutdown(v bool) *NodeGroup {
	s.GracefulShutdown = &v
	return s
}

func (s *NodeGroup) SetInstanceTypes(v []*string) *NodeGroup {
	s.InstanceTypes = v
	return s
}

func (s *NodeGroup) SetNodeGroupId(v string) *NodeGroup {
	s.NodeGroupId = &v
	return s
}

func (s *NodeGroup) SetNodeGroupName(v string) *NodeGroup {
	s.NodeGroupName = &v
	return s
}

func (s *NodeGroup) SetNodeGroupState(v string) *NodeGroup {
	s.NodeGroupState = &v
	return s
}

func (s *NodeGroup) SetNodeGroupType(v string) *NodeGroup {
	s.NodeGroupType = &v
	return s
}

func (s *NodeGroup) SetNodeResizeStrategy(v string) *NodeGroup {
	s.NodeResizeStrategy = &v
	return s
}

func (s *NodeGroup) SetPaymentType(v string) *NodeGroup {
	s.PaymentType = &v
	return s
}

func (s *NodeGroup) SetRunningNodeCount(v int32) *NodeGroup {
	s.RunningNodeCount = &v
	return s
}

func (s *NodeGroup) SetSpotBidPrices(v []*SpotBidPrice) *NodeGroup {
	s.SpotBidPrices = v
	return s
}

func (s *NodeGroup) SetSpotInstanceRemedy(v bool) *NodeGroup {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *NodeGroup) SetSpotStrategy(v string) *NodeGroup {
	s.SpotStrategy = &v
	return s
}

func (s *NodeGroup) SetStateChangeReason(v *NodeGroupStateChangeReason) *NodeGroup {
	s.StateChangeReason = v
	return s
}

func (s *NodeGroup) SetSystemDisk(v *SystemDisk) *NodeGroup {
	s.SystemDisk = v
	return s
}

func (s *NodeGroup) SetVSwitchIds(v []*string) *NodeGroup {
	s.VSwitchIds = v
	return s
}

func (s *NodeGroup) SetWithPublicIp(v bool) *NodeGroup {
	s.WithPublicIp = &v
	return s
}

func (s *NodeGroup) SetZoneId(v string) *NodeGroup {
	s.ZoneId = &v
	return s
}

type NodeGroupConfig struct {
	AdditionalSecurityGroupIds []*string            `json:"AdditionalSecurityGroupIds,omitempty" xml:"AdditionalSecurityGroupIds,omitempty" type:"Repeated"`
	CostOptimizedConfig        *CostOptimizedConfig `json:"CostOptimizedConfig,omitempty" xml:"CostOptimizedConfig,omitempty"`
	DataDisks                  []*DataDisk          `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	DeploymentSetStrategy      *string              `json:"DeploymentSetStrategy,omitempty" xml:"DeploymentSetStrategy,omitempty"`
	GracefulShutdown           *bool                `json:"GracefulShutdown,omitempty" xml:"GracefulShutdown,omitempty"`
	InstanceTypes              []*string            `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	NodeCount                  *int32               `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	NodeGroupName              *string              `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	NodeGroupType              *string              `json:"NodeGroupType,omitempty" xml:"NodeGroupType,omitempty"`
	NodeResizeStrategy         *string              `json:"NodeResizeStrategy,omitempty" xml:"NodeResizeStrategy,omitempty"`
	PaymentType                *string              `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	SpotBidPrices              []*SpotBidPrice      `json:"SpotBidPrices,omitempty" xml:"SpotBidPrices,omitempty" type:"Repeated"`
	SpotInstanceRemedy         *bool                `json:"SpotInstanceRemedy,omitempty" xml:"SpotInstanceRemedy,omitempty"`
	SpotStrategy               *string              `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SubscriptionConfig         *SubscriptionConfig  `json:"SubscriptionConfig,omitempty" xml:"SubscriptionConfig,omitempty"`
	SystemDisk                 *SystemDisk          `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
	VSwitchIds                 []*string            `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	WithPublicIp               *bool                `json:"WithPublicIp,omitempty" xml:"WithPublicIp,omitempty"`
}

func (s NodeGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s NodeGroupConfig) GoString() string {
	return s.String()
}

func (s *NodeGroupConfig) SetAdditionalSecurityGroupIds(v []*string) *NodeGroupConfig {
	s.AdditionalSecurityGroupIds = v
	return s
}

func (s *NodeGroupConfig) SetCostOptimizedConfig(v *CostOptimizedConfig) *NodeGroupConfig {
	s.CostOptimizedConfig = v
	return s
}

func (s *NodeGroupConfig) SetDataDisks(v []*DataDisk) *NodeGroupConfig {
	s.DataDisks = v
	return s
}

func (s *NodeGroupConfig) SetDeploymentSetStrategy(v string) *NodeGroupConfig {
	s.DeploymentSetStrategy = &v
	return s
}

func (s *NodeGroupConfig) SetGracefulShutdown(v bool) *NodeGroupConfig {
	s.GracefulShutdown = &v
	return s
}

func (s *NodeGroupConfig) SetInstanceTypes(v []*string) *NodeGroupConfig {
	s.InstanceTypes = v
	return s
}

func (s *NodeGroupConfig) SetNodeCount(v int32) *NodeGroupConfig {
	s.NodeCount = &v
	return s
}

func (s *NodeGroupConfig) SetNodeGroupName(v string) *NodeGroupConfig {
	s.NodeGroupName = &v
	return s
}

func (s *NodeGroupConfig) SetNodeGroupType(v string) *NodeGroupConfig {
	s.NodeGroupType = &v
	return s
}

func (s *NodeGroupConfig) SetNodeResizeStrategy(v string) *NodeGroupConfig {
	s.NodeResizeStrategy = &v
	return s
}

func (s *NodeGroupConfig) SetPaymentType(v string) *NodeGroupConfig {
	s.PaymentType = &v
	return s
}

func (s *NodeGroupConfig) SetSpotBidPrices(v []*SpotBidPrice) *NodeGroupConfig {
	s.SpotBidPrices = v
	return s
}

func (s *NodeGroupConfig) SetSpotInstanceRemedy(v bool) *NodeGroupConfig {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *NodeGroupConfig) SetSpotStrategy(v string) *NodeGroupConfig {
	s.SpotStrategy = &v
	return s
}

func (s *NodeGroupConfig) SetSubscriptionConfig(v *SubscriptionConfig) *NodeGroupConfig {
	s.SubscriptionConfig = v
	return s
}

func (s *NodeGroupConfig) SetSystemDisk(v *SystemDisk) *NodeGroupConfig {
	s.SystemDisk = v
	return s
}

func (s *NodeGroupConfig) SetVSwitchIds(v []*string) *NodeGroupConfig {
	s.VSwitchIds = v
	return s
}

func (s *NodeGroupConfig) SetWithPublicIp(v bool) *NodeGroupConfig {
	s.WithPublicIp = &v
	return s
}

type NodeGroupParam struct {
	AutoPayOrder          *bool            `json:"AutoPayOrder,omitempty" xml:"AutoPayOrder,omitempty"`
	AutoRenew             *bool            `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewDuration     *int32           `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	AutoRenewDurationUnit *string          `json:"AutoRenewDurationUnit,omitempty" xml:"AutoRenewDurationUnit,omitempty"`
	DataDisks             []*DiskInfo      `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	Description           *string          `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceTypes         []*string        `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	NodeCount             *int32           `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	NodeGroupIndex        *int32           `json:"NodeGroupIndex,omitempty" xml:"NodeGroupIndex,omitempty"`
	NodeGroupName         *string          `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	NodeGroupType         *string          `json:"NodeGroupType,omitempty" xml:"NodeGroupType,omitempty"`
	PaymentDuration       *int32           `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	PaymentDurationUnit   *string          `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	PaymentType           *string          `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	SystemDisk            *SystemDiskParam `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
	VSwitchIds            []*string        `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	ZoneId                *string          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s NodeGroupParam) String() string {
	return tea.Prettify(s)
}

func (s NodeGroupParam) GoString() string {
	return s.String()
}

func (s *NodeGroupParam) SetAutoPayOrder(v bool) *NodeGroupParam {
	s.AutoPayOrder = &v
	return s
}

func (s *NodeGroupParam) SetAutoRenew(v bool) *NodeGroupParam {
	s.AutoRenew = &v
	return s
}

func (s *NodeGroupParam) SetAutoRenewDuration(v int32) *NodeGroupParam {
	s.AutoRenewDuration = &v
	return s
}

func (s *NodeGroupParam) SetAutoRenewDurationUnit(v string) *NodeGroupParam {
	s.AutoRenewDurationUnit = &v
	return s
}

func (s *NodeGroupParam) SetDataDisks(v []*DiskInfo) *NodeGroupParam {
	s.DataDisks = v
	return s
}

func (s *NodeGroupParam) SetDescription(v string) *NodeGroupParam {
	s.Description = &v
	return s
}

func (s *NodeGroupParam) SetInstanceTypes(v []*string) *NodeGroupParam {
	s.InstanceTypes = v
	return s
}

func (s *NodeGroupParam) SetNodeCount(v int32) *NodeGroupParam {
	s.NodeCount = &v
	return s
}

func (s *NodeGroupParam) SetNodeGroupIndex(v int32) *NodeGroupParam {
	s.NodeGroupIndex = &v
	return s
}

func (s *NodeGroupParam) SetNodeGroupName(v string) *NodeGroupParam {
	s.NodeGroupName = &v
	return s
}

func (s *NodeGroupParam) SetNodeGroupType(v string) *NodeGroupParam {
	s.NodeGroupType = &v
	return s
}

func (s *NodeGroupParam) SetPaymentDuration(v int32) *NodeGroupParam {
	s.PaymentDuration = &v
	return s
}

func (s *NodeGroupParam) SetPaymentDurationUnit(v string) *NodeGroupParam {
	s.PaymentDurationUnit = &v
	return s
}

func (s *NodeGroupParam) SetPaymentType(v string) *NodeGroupParam {
	s.PaymentType = &v
	return s
}

func (s *NodeGroupParam) SetSystemDisk(v *SystemDiskParam) *NodeGroupParam {
	s.SystemDisk = v
	return s
}

func (s *NodeGroupParam) SetVSwitchIds(v []*string) *NodeGroupParam {
	s.VSwitchIds = v
	return s
}

func (s *NodeGroupParam) SetZoneId(v string) *NodeGroupParam {
	s.ZoneId = &v
	return s
}

type NodeGroupStateChangeReason struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s NodeGroupStateChangeReason) String() string {
	return tea.Prettify(s)
}

func (s NodeGroupStateChangeReason) GoString() string {
	return s.String()
}

func (s *NodeGroupStateChangeReason) SetCode(v string) *NodeGroupStateChangeReason {
	s.Code = &v
	return s
}

func (s *NodeGroupStateChangeReason) SetMessage(v string) *NodeGroupStateChangeReason {
	s.Message = &v
	return s
}

type NodeSelector struct {
	NodeGroupId    *string   `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	NodeGroupName  *string   `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	NodeGroupTypes []*string `json:"NodeGroupTypes,omitempty" xml:"NodeGroupTypes,omitempty" type:"Repeated"`
	NodeNames      []*string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty" type:"Repeated"`
	NodeSelectType *string   `json:"NodeSelectType,omitempty" xml:"NodeSelectType,omitempty"`
}

func (s NodeSelector) String() string {
	return tea.Prettify(s)
}

func (s NodeSelector) GoString() string {
	return s.String()
}

func (s *NodeSelector) SetNodeGroupId(v string) *NodeSelector {
	s.NodeGroupId = &v
	return s
}

func (s *NodeSelector) SetNodeGroupName(v string) *NodeSelector {
	s.NodeGroupName = &v
	return s
}

func (s *NodeSelector) SetNodeGroupTypes(v []*string) *NodeSelector {
	s.NodeGroupTypes = v
	return s
}

func (s *NodeSelector) SetNodeNames(v []*string) *NodeSelector {
	s.NodeNames = v
	return s
}

func (s *NodeSelector) SetNodeSelectType(v string) *NodeSelector {
	s.NodeSelectType = &v
	return s
}

type OSUser struct {
	Group    *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	User     *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s OSUser) String() string {
	return tea.Prettify(s)
}

func (s OSUser) GoString() string {
	return s.String()
}

func (s *OSUser) SetGroup(v string) *OSUser {
	s.Group = &v
	return s
}

func (s *OSUser) SetPassword(v string) *OSUser {
	s.Password = &v
	return s
}

func (s *OSUser) SetUser(v string) *OSUser {
	s.User = &v
	return s
}

type OnKubeClusterResource struct {
	Cpu    *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s OnKubeClusterResource) String() string {
	return tea.Prettify(s)
}

func (s OnKubeClusterResource) GoString() string {
	return s.String()
}

func (s *OnKubeClusterResource) SetCpu(v string) *OnKubeClusterResource {
	s.Cpu = &v
	return s
}

func (s *OnKubeClusterResource) SetMemory(v string) *OnKubeClusterResource {
	s.Memory = &v
	return s
}

type Operation struct {
	ClusterId         *string                     `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreateTime        *int64                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description       *string                     `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime           *int64                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OperationId       *string                     `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	OperationState    *string                     `json:"OperationState,omitempty" xml:"OperationState,omitempty"`
	OperationType     *string                     `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	StartTime         *int64                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StateChangeReason *OperationStateChangeReason `json:"StateChangeReason,omitempty" xml:"StateChangeReason,omitempty"`
}

func (s Operation) String() string {
	return tea.Prettify(s)
}

func (s Operation) GoString() string {
	return s.String()
}

func (s *Operation) SetClusterId(v string) *Operation {
	s.ClusterId = &v
	return s
}

func (s *Operation) SetCreateTime(v int64) *Operation {
	s.CreateTime = &v
	return s
}

func (s *Operation) SetDescription(v string) *Operation {
	s.Description = &v
	return s
}

func (s *Operation) SetEndTime(v int64) *Operation {
	s.EndTime = &v
	return s
}

func (s *Operation) SetOperationId(v string) *Operation {
	s.OperationId = &v
	return s
}

func (s *Operation) SetOperationState(v string) *Operation {
	s.OperationState = &v
	return s
}

func (s *Operation) SetOperationType(v string) *Operation {
	s.OperationType = &v
	return s
}

func (s *Operation) SetStartTime(v int64) *Operation {
	s.StartTime = &v
	return s
}

func (s *Operation) SetStateChangeReason(v *OperationStateChangeReason) *Operation {
	s.StateChangeReason = v
	return s
}

type OperationStateChangeReason struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s OperationStateChangeReason) String() string {
	return tea.Prettify(s)
}

func (s OperationStateChangeReason) GoString() string {
	return s.String()
}

func (s *OperationStateChangeReason) SetCode(v string) *OperationStateChangeReason {
	s.Code = &v
	return s
}

func (s *OperationStateChangeReason) SetMessage(v string) *OperationStateChangeReason {
	s.Message = &v
	return s
}

type Order struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	OrderId    *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s Order) String() string {
	return tea.Prettify(s)
}

func (s Order) GoString() string {
	return s.String()
}

func (s *Order) SetCreateTime(v string) *Order {
	s.CreateTime = &v
	return s
}

func (s *Order) SetOrderId(v string) *Order {
	s.OrderId = &v
	return s
}

type Page struct {
	Items      []*string `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	MaxResults *int64    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TotalCount *int64    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s Page) String() string {
	return tea.Prettify(s)
}

func (s Page) GoString() string {
	return s.String()
}

func (s *Page) SetItems(v []*string) *Page {
	s.Items = v
	return s
}

func (s *Page) SetMaxResults(v int64) *Page {
	s.MaxResults = &v
	return s
}

func (s *Page) SetNextToken(v string) *Page {
	s.NextToken = &v
	return s
}

func (s *Page) SetTotalCount(v int64) *Page {
	s.TotalCount = &v
	return s
}

type Pod struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PodName   *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	PodStatus *string `json:"PodStatus,omitempty" xml:"PodStatus,omitempty"`
	Reason    *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s Pod) String() string {
	return tea.Prettify(s)
}

func (s Pod) GoString() string {
	return s.String()
}

func (s *Pod) SetMessage(v string) *Pod {
	s.Message = &v
	return s
}

func (s *Pod) SetPodName(v string) *Pod {
	s.PodName = &v
	return s
}

func (s *Pod) SetPodStatus(v string) *Pod {
	s.PodStatus = &v
	return s
}

func (s *Pod) SetReason(v string) *Pod {
	s.Reason = &v
	return s
}

type PriceInfo struct {
	Currency                      *string          `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice                 *string          `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice                 *string          `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	PayType                       *string          `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PromotionResults              []*PromotionInfo `json:"PromotionResults,omitempty" xml:"PromotionResults,omitempty" type:"Repeated"`
	ResourceType                  *string          `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SpotInstanceTypeOriginalPrice *string          `json:"SpotInstanceTypeOriginalPrice,omitempty" xml:"SpotInstanceTypeOriginalPrice,omitempty"`
	SpotInstanceTypePrice         *string          `json:"SpotInstanceTypePrice,omitempty" xml:"SpotInstanceTypePrice,omitempty"`
	SpotOriginalPrice             *string          `json:"SpotOriginalPrice,omitempty" xml:"SpotOriginalPrice,omitempty"`
	SpotPrice                     *string          `json:"SpotPrice,omitempty" xml:"SpotPrice,omitempty"`
	TaxPrice                      *string          `json:"TaxPrice,omitempty" xml:"TaxPrice,omitempty"`
	TradePrice                    *string          `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s PriceInfo) String() string {
	return tea.Prettify(s)
}

func (s PriceInfo) GoString() string {
	return s.String()
}

func (s *PriceInfo) SetCurrency(v string) *PriceInfo {
	s.Currency = &v
	return s
}

func (s *PriceInfo) SetDiscountPrice(v string) *PriceInfo {
	s.DiscountPrice = &v
	return s
}

func (s *PriceInfo) SetOriginalPrice(v string) *PriceInfo {
	s.OriginalPrice = &v
	return s
}

func (s *PriceInfo) SetPayType(v string) *PriceInfo {
	s.PayType = &v
	return s
}

func (s *PriceInfo) SetPromotionResults(v []*PromotionInfo) *PriceInfo {
	s.PromotionResults = v
	return s
}

func (s *PriceInfo) SetResourceType(v string) *PriceInfo {
	s.ResourceType = &v
	return s
}

func (s *PriceInfo) SetSpotInstanceTypeOriginalPrice(v string) *PriceInfo {
	s.SpotInstanceTypeOriginalPrice = &v
	return s
}

func (s *PriceInfo) SetSpotInstanceTypePrice(v string) *PriceInfo {
	s.SpotInstanceTypePrice = &v
	return s
}

func (s *PriceInfo) SetSpotOriginalPrice(v string) *PriceInfo {
	s.SpotOriginalPrice = &v
	return s
}

func (s *PriceInfo) SetSpotPrice(v string) *PriceInfo {
	s.SpotPrice = &v
	return s
}

func (s *PriceInfo) SetTaxPrice(v string) *PriceInfo {
	s.TaxPrice = &v
	return s
}

func (s *PriceInfo) SetTradePrice(v string) *PriceInfo {
	s.TradePrice = &v
	return s
}

type Promotion struct {
	ProductCode         *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	PromotionOptionCode *string `json:"PromotionOptionCode,omitempty" xml:"PromotionOptionCode,omitempty"`
	PromotionOptionNo   *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
}

func (s Promotion) String() string {
	return tea.Prettify(s)
}

func (s Promotion) GoString() string {
	return s.String()
}

func (s *Promotion) SetProductCode(v string) *Promotion {
	s.ProductCode = &v
	return s
}

func (s *Promotion) SetPromotionOptionCode(v string) *Promotion {
	s.PromotionOptionCode = &v
	return s
}

func (s *Promotion) SetPromotionOptionNo(v string) *Promotion {
	s.PromotionOptionNo = &v
	return s
}

type PromotionInfo struct {
	CanPromFee          *string `json:"CanPromFee,omitempty" xml:"CanPromFee,omitempty"`
	IsSelected          *string `json:"IsSelected,omitempty" xml:"IsSelected,omitempty"`
	PromotionDesc       *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	PromotionName       *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	PromotionOptionCode *string `json:"PromotionOptionCode,omitempty" xml:"PromotionOptionCode,omitempty"`
	PromotionOptionNo   *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
}

func (s PromotionInfo) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfo) GoString() string {
	return s.String()
}

func (s *PromotionInfo) SetCanPromFee(v string) *PromotionInfo {
	s.CanPromFee = &v
	return s
}

func (s *PromotionInfo) SetIsSelected(v string) *PromotionInfo {
	s.IsSelected = &v
	return s
}

func (s *PromotionInfo) SetPromotionDesc(v string) *PromotionInfo {
	s.PromotionDesc = &v
	return s
}

func (s *PromotionInfo) SetPromotionName(v string) *PromotionInfo {
	s.PromotionName = &v
	return s
}

func (s *PromotionInfo) SetPromotionOptionCode(v string) *PromotionInfo {
	s.PromotionOptionCode = &v
	return s
}

func (s *PromotionInfo) SetPromotionOptionNo(v string) *PromotionInfo {
	s.PromotionOptionNo = &v
	return s
}

type PromotionParam struct {
	ProductCode         *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	PromotionOptionCode *string `json:"PromotionOptionCode,omitempty" xml:"PromotionOptionCode,omitempty"`
	PromotionOptionNo   *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
}

func (s PromotionParam) String() string {
	return tea.Prettify(s)
}

func (s PromotionParam) GoString() string {
	return s.String()
}

func (s *PromotionParam) SetProductCode(v string) *PromotionParam {
	s.ProductCode = &v
	return s
}

func (s *PromotionParam) SetPromotionOptionCode(v string) *PromotionParam {
	s.PromotionOptionCode = &v
	return s
}

func (s *PromotionParam) SetPromotionOptionNo(v string) *PromotionParam {
	s.PromotionOptionNo = &v
	return s
}

type RenewInstance struct {
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RenewDuration     *int32  `json:"RenewDuration,omitempty" xml:"RenewDuration,omitempty"`
	RenewDurationUnit *string `json:"RenewDurationUnit,omitempty" xml:"RenewDurationUnit,omitempty"`
}

func (s RenewInstance) String() string {
	return tea.Prettify(s)
}

func (s RenewInstance) GoString() string {
	return s.String()
}

func (s *RenewInstance) SetInstanceId(v string) *RenewInstance {
	s.InstanceId = &v
	return s
}

func (s *RenewInstance) SetRenewDuration(v int32) *RenewInstance {
	s.RenewDuration = &v
	return s
}

func (s *RenewInstance) SetRenewDurationUnit(v string) *RenewInstance {
	s.RenewDurationUnit = &v
	return s
}

type RenewInstanceParam struct {
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RenewDuration     *int64  `json:"RenewDuration,omitempty" xml:"RenewDuration,omitempty"`
	RenewDurationUnit *string `json:"RenewDurationUnit,omitempty" xml:"RenewDurationUnit,omitempty"`
}

func (s RenewInstanceParam) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceParam) GoString() string {
	return s.String()
}

func (s *RenewInstanceParam) SetInstanceId(v string) *RenewInstanceParam {
	s.InstanceId = &v
	return s
}

func (s *RenewInstanceParam) SetRenewDuration(v int64) *RenewInstanceParam {
	s.RenewDuration = &v
	return s
}

func (s *RenewInstanceParam) SetRenewDurationUnit(v string) *RenewInstanceParam {
	s.RenewDurationUnit = &v
	return s
}

type RequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RequestTag) String() string {
	return tea.Prettify(s)
}

func (s RequestTag) GoString() string {
	return s.String()
}

func (s *RequestTag) SetKey(v string) *RequestTag {
	s.Key = &v
	return s
}

func (s *RequestTag) SetValue(v string) *RequestTag {
	s.Value = &v
	return s
}

type ResizeDiskNodeGroupParam struct {
	DataDiskCapacity *int64  `json:"DataDiskCapacity,omitempty" xml:"DataDiskCapacity,omitempty"`
	NodeGroupId      *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	RollingRestart   *bool   `json:"RollingRestart,omitempty" xml:"RollingRestart,omitempty"`
}

func (s ResizeDiskNodeGroupParam) String() string {
	return tea.Prettify(s)
}

func (s ResizeDiskNodeGroupParam) GoString() string {
	return s.String()
}

func (s *ResizeDiskNodeGroupParam) SetDataDiskCapacity(v int64) *ResizeDiskNodeGroupParam {
	s.DataDiskCapacity = &v
	return s
}

func (s *ResizeDiskNodeGroupParam) SetNodeGroupId(v string) *ResizeDiskNodeGroupParam {
	s.NodeGroupId = &v
	return s
}

func (s *ResizeDiskNodeGroupParam) SetRollingRestart(v bool) *ResizeDiskNodeGroupParam {
	s.RollingRestart = &v
	return s
}

type ScalingActivity struct {
	Cause            *string `json:"Cause,omitempty" xml:"Cause,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime          *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EssScalingRuleId *string `json:"EssScalingRuleId,omitempty" xml:"EssScalingRuleId,omitempty"`
	ExpectNum        *int32  `json:"ExpectNum,omitempty" xml:"ExpectNum,omitempty"`
	HostGroupName    *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceIds      *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	ScalingGroupId   *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScalingRuleName  *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	StartTime        *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalCapacity    *int32  `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
	Transition       *string `json:"Transition,omitempty" xml:"Transition,omitempty"`
}

func (s ScalingActivity) String() string {
	return tea.Prettify(s)
}

func (s ScalingActivity) GoString() string {
	return s.String()
}

func (s *ScalingActivity) SetCause(v string) *ScalingActivity {
	s.Cause = &v
	return s
}

func (s *ScalingActivity) SetDescription(v string) *ScalingActivity {
	s.Description = &v
	return s
}

func (s *ScalingActivity) SetEndTime(v int64) *ScalingActivity {
	s.EndTime = &v
	return s
}

func (s *ScalingActivity) SetEssScalingRuleId(v string) *ScalingActivity {
	s.EssScalingRuleId = &v
	return s
}

func (s *ScalingActivity) SetExpectNum(v int32) *ScalingActivity {
	s.ExpectNum = &v
	return s
}

func (s *ScalingActivity) SetHostGroupName(v string) *ScalingActivity {
	s.HostGroupName = &v
	return s
}

func (s *ScalingActivity) SetId(v string) *ScalingActivity {
	s.Id = &v
	return s
}

func (s *ScalingActivity) SetInstanceIds(v string) *ScalingActivity {
	s.InstanceIds = &v
	return s
}

func (s *ScalingActivity) SetScalingGroupId(v string) *ScalingActivity {
	s.ScalingGroupId = &v
	return s
}

func (s *ScalingActivity) SetScalingRuleName(v string) *ScalingActivity {
	s.ScalingRuleName = &v
	return s
}

func (s *ScalingActivity) SetStartTime(v int64) *ScalingActivity {
	s.StartTime = &v
	return s
}

func (s *ScalingActivity) SetStatus(v string) *ScalingActivity {
	s.Status = &v
	return s
}

func (s *ScalingActivity) SetTotalCapacity(v int32) *ScalingActivity {
	s.TotalCapacity = &v
	return s
}

func (s *ScalingActivity) SetTransition(v string) *ScalingActivity {
	s.Transition = &v
	return s
}

type ScalingGroupConfig struct {
	DataDiskCategory     *string                                 `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	DataDiskCount        *int32                                  `json:"DataDiskCount,omitempty" xml:"DataDiskCount,omitempty"`
	DataDiskSize         *int64                                  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	DefaultCoolDownTime  *int64                                  `json:"DefaultCoolDownTime,omitempty" xml:"DefaultCoolDownTime,omitempty"`
	InstanceTypeList     []*ScalingGroupConfigInstanceTypeList   `json:"InstanceTypeList,omitempty" xml:"InstanceTypeList,omitempty" type:"Repeated"`
	MultiAvailablePolicy *ScalingGroupConfigMultiAvailablePolicy `json:"MultiAvailablePolicy,omitempty" xml:"MultiAvailablePolicy,omitempty" type:"Struct"`
	NodeOfflinePolicy    *ScalingGroupConfigNodeOfflinePolicy    `json:"NodeOfflinePolicy,omitempty" xml:"NodeOfflinePolicy,omitempty" type:"Struct"`
	PrivatePoolOptions   *ScalingGroupConfigPrivatePoolOptions   `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	ScalingMaxSize       *int32                                  `json:"ScalingMaxSize,omitempty" xml:"ScalingMaxSize,omitempty"`
	ScalingMinSize       *int32                                  `json:"ScalingMinSize,omitempty" xml:"ScalingMinSize,omitempty"`
	SpotStrategy         *string                                 `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SysDiskCategory      *string                                 `json:"SysDiskCategory,omitempty" xml:"SysDiskCategory,omitempty"`
	SysDiskSize          *int64                                  `json:"SysDiskSize,omitempty" xml:"SysDiskSize,omitempty"`
	TriggerMode          *string                                 `json:"TriggerMode,omitempty" xml:"TriggerMode,omitempty"`
}

func (s ScalingGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s ScalingGroupConfig) GoString() string {
	return s.String()
}

func (s *ScalingGroupConfig) SetDataDiskCategory(v string) *ScalingGroupConfig {
	s.DataDiskCategory = &v
	return s
}

func (s *ScalingGroupConfig) SetDataDiskCount(v int32) *ScalingGroupConfig {
	s.DataDiskCount = &v
	return s
}

func (s *ScalingGroupConfig) SetDataDiskSize(v int64) *ScalingGroupConfig {
	s.DataDiskSize = &v
	return s
}

func (s *ScalingGroupConfig) SetDefaultCoolDownTime(v int64) *ScalingGroupConfig {
	s.DefaultCoolDownTime = &v
	return s
}

func (s *ScalingGroupConfig) SetInstanceTypeList(v []*ScalingGroupConfigInstanceTypeList) *ScalingGroupConfig {
	s.InstanceTypeList = v
	return s
}

func (s *ScalingGroupConfig) SetMultiAvailablePolicy(v *ScalingGroupConfigMultiAvailablePolicy) *ScalingGroupConfig {
	s.MultiAvailablePolicy = v
	return s
}

func (s *ScalingGroupConfig) SetNodeOfflinePolicy(v *ScalingGroupConfigNodeOfflinePolicy) *ScalingGroupConfig {
	s.NodeOfflinePolicy = v
	return s
}

func (s *ScalingGroupConfig) SetPrivatePoolOptions(v *ScalingGroupConfigPrivatePoolOptions) *ScalingGroupConfig {
	s.PrivatePoolOptions = v
	return s
}

func (s *ScalingGroupConfig) SetScalingMaxSize(v int32) *ScalingGroupConfig {
	s.ScalingMaxSize = &v
	return s
}

func (s *ScalingGroupConfig) SetScalingMinSize(v int32) *ScalingGroupConfig {
	s.ScalingMinSize = &v
	return s
}

func (s *ScalingGroupConfig) SetSpotStrategy(v string) *ScalingGroupConfig {
	s.SpotStrategy = &v
	return s
}

func (s *ScalingGroupConfig) SetSysDiskCategory(v string) *ScalingGroupConfig {
	s.SysDiskCategory = &v
	return s
}

func (s *ScalingGroupConfig) SetSysDiskSize(v int64) *ScalingGroupConfig {
	s.SysDiskSize = &v
	return s
}

func (s *ScalingGroupConfig) SetTriggerMode(v string) *ScalingGroupConfig {
	s.TriggerMode = &v
	return s
}

type ScalingGroupConfigInstanceTypeList struct {
	InstanceType   *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
}

func (s ScalingGroupConfigInstanceTypeList) String() string {
	return tea.Prettify(s)
}

func (s ScalingGroupConfigInstanceTypeList) GoString() string {
	return s.String()
}

func (s *ScalingGroupConfigInstanceTypeList) SetInstanceType(v string) *ScalingGroupConfigInstanceTypeList {
	s.InstanceType = &v
	return s
}

func (s *ScalingGroupConfigInstanceTypeList) SetSpotPriceLimit(v float32) *ScalingGroupConfigInstanceTypeList {
	s.SpotPriceLimit = &v
	return s
}

type ScalingGroupConfigMultiAvailablePolicy struct {
	PolicyParam *ScalingGroupConfigMultiAvailablePolicyPolicyParam `json:"PolicyParam,omitempty" xml:"PolicyParam,omitempty" type:"Struct"`
	PolicyType  *string                                            `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ScalingGroupConfigMultiAvailablePolicy) String() string {
	return tea.Prettify(s)
}

func (s ScalingGroupConfigMultiAvailablePolicy) GoString() string {
	return s.String()
}

func (s *ScalingGroupConfigMultiAvailablePolicy) SetPolicyParam(v *ScalingGroupConfigMultiAvailablePolicyPolicyParam) *ScalingGroupConfigMultiAvailablePolicy {
	s.PolicyParam = v
	return s
}

func (s *ScalingGroupConfigMultiAvailablePolicy) SetPolicyType(v string) *ScalingGroupConfigMultiAvailablePolicy {
	s.PolicyType = &v
	return s
}

type ScalingGroupConfigMultiAvailablePolicyPolicyParam struct {
	OnDemandBaseCapacity                *int32 `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	OnDemandPercentageAboveBaseCapacity *int32 `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	SpotInstancePools                   *int32 `json:"SpotInstancePools,omitempty" xml:"SpotInstancePools,omitempty"`
	SpotInstanceRemedy                  *bool  `json:"SpotInstanceRemedy,omitempty" xml:"SpotInstanceRemedy,omitempty"`
}

func (s ScalingGroupConfigMultiAvailablePolicyPolicyParam) String() string {
	return tea.Prettify(s)
}

func (s ScalingGroupConfigMultiAvailablePolicyPolicyParam) GoString() string {
	return s.String()
}

func (s *ScalingGroupConfigMultiAvailablePolicyPolicyParam) SetOnDemandBaseCapacity(v int32) *ScalingGroupConfigMultiAvailablePolicyPolicyParam {
	s.OnDemandBaseCapacity = &v
	return s
}

func (s *ScalingGroupConfigMultiAvailablePolicyPolicyParam) SetOnDemandPercentageAboveBaseCapacity(v int32) *ScalingGroupConfigMultiAvailablePolicyPolicyParam {
	s.OnDemandPercentageAboveBaseCapacity = &v
	return s
}

func (s *ScalingGroupConfigMultiAvailablePolicyPolicyParam) SetSpotInstancePools(v int32) *ScalingGroupConfigMultiAvailablePolicyPolicyParam {
	s.SpotInstancePools = &v
	return s
}

func (s *ScalingGroupConfigMultiAvailablePolicyPolicyParam) SetSpotInstanceRemedy(v bool) *ScalingGroupConfigMultiAvailablePolicyPolicyParam {
	s.SpotInstanceRemedy = &v
	return s
}

type ScalingGroupConfigNodeOfflinePolicy struct {
	Mode      *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	TimeoutMs *int64  `json:"TimeoutMs,omitempty" xml:"TimeoutMs,omitempty"`
}

func (s ScalingGroupConfigNodeOfflinePolicy) String() string {
	return tea.Prettify(s)
}

func (s ScalingGroupConfigNodeOfflinePolicy) GoString() string {
	return s.String()
}

func (s *ScalingGroupConfigNodeOfflinePolicy) SetMode(v string) *ScalingGroupConfigNodeOfflinePolicy {
	s.Mode = &v
	return s
}

func (s *ScalingGroupConfigNodeOfflinePolicy) SetTimeoutMs(v int64) *ScalingGroupConfigNodeOfflinePolicy {
	s.TimeoutMs = &v
	return s
}

type ScalingGroupConfigPrivatePoolOptions struct {
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
}

func (s ScalingGroupConfigPrivatePoolOptions) String() string {
	return tea.Prettify(s)
}

func (s ScalingGroupConfigPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *ScalingGroupConfigPrivatePoolOptions) SetId(v string) *ScalingGroupConfigPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *ScalingGroupConfigPrivatePoolOptions) SetMatchCriteria(v string) *ScalingGroupConfigPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

type ScalingRule struct {
	AdjustmentValue     *int32                        `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	ByLoadScalingRule   *ScalingRuleByLoadScalingRule `json:"ByLoadScalingRule,omitempty" xml:"ByLoadScalingRule,omitempty" type:"Struct"`
	ByTimeScalingRule   *ScalingRuleByTimeScalingRule `json:"ByTimeScalingRule,omitempty" xml:"ByTimeScalingRule,omitempty" type:"Struct"`
	CoolDownInterval    *int32                        `json:"CoolDownInterval,omitempty" xml:"CoolDownInterval,omitempty"`
	ScalingActivityType *string                       `json:"ScalingActivityType,omitempty" xml:"ScalingActivityType,omitempty"`
	ScalingRuleName     *string                       `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	ScalingRuleType     *string                       `json:"ScalingRuleType,omitempty" xml:"ScalingRuleType,omitempty"`
}

func (s ScalingRule) String() string {
	return tea.Prettify(s)
}

func (s ScalingRule) GoString() string {
	return s.String()
}

func (s *ScalingRule) SetAdjustmentValue(v int32) *ScalingRule {
	s.AdjustmentValue = &v
	return s
}

func (s *ScalingRule) SetByLoadScalingRule(v *ScalingRuleByLoadScalingRule) *ScalingRule {
	s.ByLoadScalingRule = v
	return s
}

func (s *ScalingRule) SetByTimeScalingRule(v *ScalingRuleByTimeScalingRule) *ScalingRule {
	s.ByTimeScalingRule = v
	return s
}

func (s *ScalingRule) SetCoolDownInterval(v int32) *ScalingRule {
	s.CoolDownInterval = &v
	return s
}

func (s *ScalingRule) SetScalingActivityType(v string) *ScalingRule {
	s.ScalingActivityType = &v
	return s
}

func (s *ScalingRule) SetScalingRuleName(v string) *ScalingRule {
	s.ScalingRuleName = &v
	return s
}

func (s *ScalingRule) SetScalingRuleType(v string) *ScalingRule {
	s.ScalingRuleType = &v
	return s
}

type ScalingRuleByLoadScalingRule struct {
	ComparisonOperator *string  `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	CoolDownInterval   *int32   `json:"CoolDownInterval,omitempty" xml:"CoolDownInterval,omitempty"`
	EvaluationCount    *int32   `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	MetricName         *string  `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Statistics         *string  `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *float64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	TimeWindow         *int32   `json:"TimeWindow,omitempty" xml:"TimeWindow,omitempty"`
}

func (s ScalingRuleByLoadScalingRule) String() string {
	return tea.Prettify(s)
}

func (s ScalingRuleByLoadScalingRule) GoString() string {
	return s.String()
}

func (s *ScalingRuleByLoadScalingRule) SetComparisonOperator(v string) *ScalingRuleByLoadScalingRule {
	s.ComparisonOperator = &v
	return s
}

func (s *ScalingRuleByLoadScalingRule) SetCoolDownInterval(v int32) *ScalingRuleByLoadScalingRule {
	s.CoolDownInterval = &v
	return s
}

func (s *ScalingRuleByLoadScalingRule) SetEvaluationCount(v int32) *ScalingRuleByLoadScalingRule {
	s.EvaluationCount = &v
	return s
}

func (s *ScalingRuleByLoadScalingRule) SetMetricName(v string) *ScalingRuleByLoadScalingRule {
	s.MetricName = &v
	return s
}

func (s *ScalingRuleByLoadScalingRule) SetStatistics(v string) *ScalingRuleByLoadScalingRule {
	s.Statistics = &v
	return s
}

func (s *ScalingRuleByLoadScalingRule) SetThreshold(v float64) *ScalingRuleByLoadScalingRule {
	s.Threshold = &v
	return s
}

func (s *ScalingRuleByLoadScalingRule) SetTimeWindow(v int32) *ScalingRuleByLoadScalingRule {
	s.TimeWindow = &v
	return s
}

type ScalingRuleByTimeScalingRule struct {
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LaunchTime      *int64  `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	RecurrenceType  *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
}

func (s ScalingRuleByTimeScalingRule) String() string {
	return tea.Prettify(s)
}

func (s ScalingRuleByTimeScalingRule) GoString() string {
	return s.String()
}

func (s *ScalingRuleByTimeScalingRule) SetEndTime(v int64) *ScalingRuleByTimeScalingRule {
	s.EndTime = &v
	return s
}

func (s *ScalingRuleByTimeScalingRule) SetLaunchTime(v int64) *ScalingRuleByTimeScalingRule {
	s.LaunchTime = &v
	return s
}

func (s *ScalingRuleByTimeScalingRule) SetRecurrenceType(v string) *ScalingRuleByTimeScalingRule {
	s.RecurrenceType = &v
	return s
}

func (s *ScalingRuleByTimeScalingRule) SetRecurrenceValue(v string) *ScalingRuleByTimeScalingRule {
	s.RecurrenceValue = &v
	return s
}

type ScalingRuleSpec struct {
	AdjustmentValue       *int32                                `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	ByLoadScalingRuleSpec *ScalingRuleSpecByLoadScalingRuleSpec `json:"ByLoadScalingRuleSpec,omitempty" xml:"ByLoadScalingRuleSpec,omitempty" type:"Struct"`
	ByTimeScalingRuleSpec *ScalingRuleSpecByTimeScalingRuleSpec `json:"ByTimeScalingRuleSpec,omitempty" xml:"ByTimeScalingRuleSpec,omitempty" type:"Struct"`
	CoolDownInterval      *int32                                `json:"CoolDownInterval,omitempty" xml:"CoolDownInterval,omitempty"`
	ScalingActivityType   *string                               `json:"ScalingActivityType,omitempty" xml:"ScalingActivityType,omitempty"`
	ScalingRuleName       *string                               `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	ScalingRuleType       *string                               `json:"ScalingRuleType,omitempty" xml:"ScalingRuleType,omitempty"`
}

func (s ScalingRuleSpec) String() string {
	return tea.Prettify(s)
}

func (s ScalingRuleSpec) GoString() string {
	return s.String()
}

func (s *ScalingRuleSpec) SetAdjustmentValue(v int32) *ScalingRuleSpec {
	s.AdjustmentValue = &v
	return s
}

func (s *ScalingRuleSpec) SetByLoadScalingRuleSpec(v *ScalingRuleSpecByLoadScalingRuleSpec) *ScalingRuleSpec {
	s.ByLoadScalingRuleSpec = v
	return s
}

func (s *ScalingRuleSpec) SetByTimeScalingRuleSpec(v *ScalingRuleSpecByTimeScalingRuleSpec) *ScalingRuleSpec {
	s.ByTimeScalingRuleSpec = v
	return s
}

func (s *ScalingRuleSpec) SetCoolDownInterval(v int32) *ScalingRuleSpec {
	s.CoolDownInterval = &v
	return s
}

func (s *ScalingRuleSpec) SetScalingActivityType(v string) *ScalingRuleSpec {
	s.ScalingActivityType = &v
	return s
}

func (s *ScalingRuleSpec) SetScalingRuleName(v string) *ScalingRuleSpec {
	s.ScalingRuleName = &v
	return s
}

func (s *ScalingRuleSpec) SetScalingRuleType(v string) *ScalingRuleSpec {
	s.ScalingRuleType = &v
	return s
}

type ScalingRuleSpecByLoadScalingRuleSpec struct {
	ComparisonOperator *string  `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	EvaluationCount    *int32   `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	MetricName         *string  `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Statistics         *string  `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *float64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	TimeWindow         *int32   `json:"TimeWindow,omitempty" xml:"TimeWindow,omitempty"`
}

func (s ScalingRuleSpecByLoadScalingRuleSpec) String() string {
	return tea.Prettify(s)
}

func (s ScalingRuleSpecByLoadScalingRuleSpec) GoString() string {
	return s.String()
}

func (s *ScalingRuleSpecByLoadScalingRuleSpec) SetComparisonOperator(v string) *ScalingRuleSpecByLoadScalingRuleSpec {
	s.ComparisonOperator = &v
	return s
}

func (s *ScalingRuleSpecByLoadScalingRuleSpec) SetEvaluationCount(v int32) *ScalingRuleSpecByLoadScalingRuleSpec {
	s.EvaluationCount = &v
	return s
}

func (s *ScalingRuleSpecByLoadScalingRuleSpec) SetMetricName(v string) *ScalingRuleSpecByLoadScalingRuleSpec {
	s.MetricName = &v
	return s
}

func (s *ScalingRuleSpecByLoadScalingRuleSpec) SetStatistics(v string) *ScalingRuleSpecByLoadScalingRuleSpec {
	s.Statistics = &v
	return s
}

func (s *ScalingRuleSpecByLoadScalingRuleSpec) SetThreshold(v float64) *ScalingRuleSpecByLoadScalingRuleSpec {
	s.Threshold = &v
	return s
}

func (s *ScalingRuleSpecByLoadScalingRuleSpec) SetTimeWindow(v int32) *ScalingRuleSpecByLoadScalingRuleSpec {
	s.TimeWindow = &v
	return s
}

type ScalingRuleSpecByTimeScalingRuleSpec struct {
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LaunchTime      *int64  `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	RecurrenceType  *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
}

func (s ScalingRuleSpecByTimeScalingRuleSpec) String() string {
	return tea.Prettify(s)
}

func (s ScalingRuleSpecByTimeScalingRuleSpec) GoString() string {
	return s.String()
}

func (s *ScalingRuleSpecByTimeScalingRuleSpec) SetEndTime(v int64) *ScalingRuleSpecByTimeScalingRuleSpec {
	s.EndTime = &v
	return s
}

func (s *ScalingRuleSpecByTimeScalingRuleSpec) SetLaunchTime(v int64) *ScalingRuleSpecByTimeScalingRuleSpec {
	s.LaunchTime = &v
	return s
}

func (s *ScalingRuleSpecByTimeScalingRuleSpec) SetRecurrenceType(v string) *ScalingRuleSpecByTimeScalingRuleSpec {
	s.RecurrenceType = &v
	return s
}

func (s *ScalingRuleSpecByTimeScalingRuleSpec) SetRecurrenceValue(v string) *ScalingRuleSpecByTimeScalingRuleSpec {
	s.RecurrenceValue = &v
	return s
}

type ScalingRuleV1 struct {
	AdjustmentType     *string                 `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	AdjustmentValue    *int32                  `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	CoolDownTime       *int32                  `json:"CoolDownTime,omitempty" xml:"CoolDownTime,omitempty"`
	RuleName           *string                 `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleParam          *ScalingRuleV1RuleParam `json:"RuleParam,omitempty" xml:"RuleParam,omitempty" type:"Struct"`
	RuleType           *string                 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	ScalingConfigBizId *string                 `json:"ScalingConfigBizId,omitempty" xml:"ScalingConfigBizId,omitempty"`
}

func (s ScalingRuleV1) String() string {
	return tea.Prettify(s)
}

func (s ScalingRuleV1) GoString() string {
	return s.String()
}

func (s *ScalingRuleV1) SetAdjustmentType(v string) *ScalingRuleV1 {
	s.AdjustmentType = &v
	return s
}

func (s *ScalingRuleV1) SetAdjustmentValue(v int32) *ScalingRuleV1 {
	s.AdjustmentValue = &v
	return s
}

func (s *ScalingRuleV1) SetCoolDownTime(v int32) *ScalingRuleV1 {
	s.CoolDownTime = &v
	return s
}

func (s *ScalingRuleV1) SetRuleName(v string) *ScalingRuleV1 {
	s.RuleName = &v
	return s
}

func (s *ScalingRuleV1) SetRuleParam(v *ScalingRuleV1RuleParam) *ScalingRuleV1 {
	s.RuleParam = v
	return s
}

func (s *ScalingRuleV1) SetRuleType(v string) *ScalingRuleV1 {
	s.RuleType = &v
	return s
}

func (s *ScalingRuleV1) SetScalingConfigBizId(v string) *ScalingRuleV1 {
	s.ScalingConfigBizId = &v
	return s
}

type ScalingRuleV1RuleParam struct {
	ComparisonOperator   *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	EvaluationCount      *int32  `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	LaunchExpirationTime *int32  `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
	LaunchTime           *string `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	MetricName           *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Period               *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	RecurrenceEndTime    *string `json:"RecurrenceEndTime,omitempty" xml:"RecurrenceEndTime,omitempty"`
	RecurrenceType       *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValue      *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	Statistics           *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold            *int32  `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ScalingRuleV1RuleParam) String() string {
	return tea.Prettify(s)
}

func (s ScalingRuleV1RuleParam) GoString() string {
	return s.String()
}

func (s *ScalingRuleV1RuleParam) SetComparisonOperator(v string) *ScalingRuleV1RuleParam {
	s.ComparisonOperator = &v
	return s
}

func (s *ScalingRuleV1RuleParam) SetEvaluationCount(v int32) *ScalingRuleV1RuleParam {
	s.EvaluationCount = &v
	return s
}

func (s *ScalingRuleV1RuleParam) SetLaunchExpirationTime(v int32) *ScalingRuleV1RuleParam {
	s.LaunchExpirationTime = &v
	return s
}

func (s *ScalingRuleV1RuleParam) SetLaunchTime(v string) *ScalingRuleV1RuleParam {
	s.LaunchTime = &v
	return s
}

func (s *ScalingRuleV1RuleParam) SetMetricName(v string) *ScalingRuleV1RuleParam {
	s.MetricName = &v
	return s
}

func (s *ScalingRuleV1RuleParam) SetPeriod(v int32) *ScalingRuleV1RuleParam {
	s.Period = &v
	return s
}

func (s *ScalingRuleV1RuleParam) SetRecurrenceEndTime(v string) *ScalingRuleV1RuleParam {
	s.RecurrenceEndTime = &v
	return s
}

func (s *ScalingRuleV1RuleParam) SetRecurrenceType(v string) *ScalingRuleV1RuleParam {
	s.RecurrenceType = &v
	return s
}

func (s *ScalingRuleV1RuleParam) SetRecurrenceValue(v string) *ScalingRuleV1RuleParam {
	s.RecurrenceValue = &v
	return s
}

func (s *ScalingRuleV1RuleParam) SetStatistics(v string) *ScalingRuleV1RuleParam {
	s.Statistics = &v
	return s
}

func (s *ScalingRuleV1RuleParam) SetThreshold(v int32) *ScalingRuleV1RuleParam {
	s.Threshold = &v
	return s
}

type Script struct {
	ExecutionFailStrategy *string       `json:"ExecutionFailStrategy,omitempty" xml:"ExecutionFailStrategy,omitempty"`
	ExecutionMoment       *string       `json:"ExecutionMoment,omitempty" xml:"ExecutionMoment,omitempty"`
	NodeSelector          *NodeSelector `json:"NodeSelector,omitempty" xml:"NodeSelector,omitempty"`
	Priority              *int32        `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ScriptArgs            *string       `json:"ScriptArgs,omitempty" xml:"ScriptArgs,omitempty"`
	ScriptName            *string       `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	ScriptPath            *string       `json:"ScriptPath,omitempty" xml:"ScriptPath,omitempty"`
}

func (s Script) String() string {
	return tea.Prettify(s)
}

func (s Script) GoString() string {
	return s.String()
}

func (s *Script) SetExecutionFailStrategy(v string) *Script {
	s.ExecutionFailStrategy = &v
	return s
}

func (s *Script) SetExecutionMoment(v string) *Script {
	s.ExecutionMoment = &v
	return s
}

func (s *Script) SetNodeSelector(v *NodeSelector) *Script {
	s.NodeSelector = v
	return s
}

func (s *Script) SetPriority(v int32) *Script {
	s.Priority = &v
	return s
}

func (s *Script) SetScriptArgs(v string) *Script {
	s.ScriptArgs = &v
	return s
}

func (s *Script) SetScriptName(v string) *Script {
	s.ScriptName = &v
	return s
}

func (s *Script) SetScriptPath(v string) *Script {
	s.ScriptPath = &v
	return s
}

type SpotBidPrice struct {
	BidPrice     *float64 `json:"BidPrice,omitempty" xml:"BidPrice,omitempty"`
	InstanceType *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s SpotBidPrice) String() string {
	return tea.Prettify(s)
}

func (s SpotBidPrice) GoString() string {
	return s.String()
}

func (s *SpotBidPrice) SetBidPrice(v float64) *SpotBidPrice {
	s.BidPrice = &v
	return s
}

func (s *SpotBidPrice) SetInstanceType(v string) *SpotBidPrice {
	s.InstanceType = &v
	return s
}

type SpotPriceLimit struct {
	InstanceType *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	PriceLimit   *float64 `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
}

func (s SpotPriceLimit) String() string {
	return tea.Prettify(s)
}

func (s SpotPriceLimit) GoString() string {
	return s.String()
}

func (s *SpotPriceLimit) SetInstanceType(v string) *SpotPriceLimit {
	s.InstanceType = &v
	return s
}

func (s *SpotPriceLimit) SetPriceLimit(v float64) *SpotPriceLimit {
	s.PriceLimit = &v
	return s
}

type StateChangeReason struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s StateChangeReason) String() string {
	return tea.Prettify(s)
}

func (s StateChangeReason) GoString() string {
	return s.String()
}

func (s *StateChangeReason) SetCode(v string) *StateChangeReason {
	s.Code = &v
	return s
}

func (s *StateChangeReason) SetMessage(v string) *StateChangeReason {
	s.Message = &v
	return s
}

type SubscriptionConfig struct {
	AutoRenew             *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewDuration     *int32  `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	AutoRenewDurationUnit *string `json:"AutoRenewDurationUnit,omitempty" xml:"AutoRenewDurationUnit,omitempty"`
	PaymentDuration       *int32  `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	PaymentDurationUnit   *string `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
}

func (s SubscriptionConfig) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionConfig) GoString() string {
	return s.String()
}

func (s *SubscriptionConfig) SetAutoRenew(v bool) *SubscriptionConfig {
	s.AutoRenew = &v
	return s
}

func (s *SubscriptionConfig) SetAutoRenewDuration(v int32) *SubscriptionConfig {
	s.AutoRenewDuration = &v
	return s
}

func (s *SubscriptionConfig) SetAutoRenewDurationUnit(v string) *SubscriptionConfig {
	s.AutoRenewDurationUnit = &v
	return s
}

func (s *SubscriptionConfig) SetPaymentDuration(v int32) *SubscriptionConfig {
	s.PaymentDuration = &v
	return s
}

func (s *SubscriptionConfig) SetPaymentDurationUnit(v string) *SubscriptionConfig {
	s.PaymentDurationUnit = &v
	return s
}

type SystemDisk struct {
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Count            *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Size             *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s SystemDisk) String() string {
	return tea.Prettify(s)
}

func (s SystemDisk) GoString() string {
	return s.String()
}

func (s *SystemDisk) SetCategory(v string) *SystemDisk {
	s.Category = &v
	return s
}

func (s *SystemDisk) SetCount(v int32) *SystemDisk {
	s.Count = &v
	return s
}

func (s *SystemDisk) SetPerformanceLevel(v string) *SystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *SystemDisk) SetSize(v int32) *SystemDisk {
	s.Size = &v
	return s
}

type SystemDiskParam struct {
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Size             *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s SystemDiskParam) String() string {
	return tea.Prettify(s)
}

func (s SystemDiskParam) GoString() string {
	return s.String()
}

func (s *SystemDiskParam) SetCategory(v string) *SystemDiskParam {
	s.Category = &v
	return s
}

func (s *SystemDiskParam) SetPerformanceLevel(v string) *SystemDiskParam {
	s.PerformanceLevel = &v
	return s
}

func (s *SystemDiskParam) SetSize(v int32) *SystemDiskParam {
	s.Size = &v
	return s
}

type Tag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s Tag) String() string {
	return tea.Prettify(s)
}

func (s Tag) GoString() string {
	return s.String()
}

func (s *Tag) SetKey(v string) *Tag {
	s.Key = &v
	return s
}

func (s *Tag) SetValue(v string) *Tag {
	s.Value = &v
	return s
}

type TagResource struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s TagResource) String() string {
	return tea.Prettify(s)
}

func (s TagResource) GoString() string {
	return s.String()
}

func (s *TagResource) SetResourceId(v string) *TagResource {
	s.ResourceId = &v
	return s
}

func (s *TagResource) SetResourceType(v string) *TagResource {
	s.ResourceType = &v
	return s
}

func (s *TagResource) SetTagKey(v string) *TagResource {
	s.TagKey = &v
	return s
}

func (s *TagResource) SetTagValue(v string) *TagResource {
	s.TagValue = &v
	return s
}

type UpdateApplicationConfig struct {
	ConfigAction      *string `json:"ConfigAction,omitempty" xml:"ConfigAction,omitempty"`
	ConfigDescription *string `json:"ConfigDescription,omitempty" xml:"ConfigDescription,omitempty"`
	ConfigFileName    *string `json:"ConfigFileName,omitempty" xml:"ConfigFileName,omitempty"`
	ConfigItemKey     *string `json:"ConfigItemKey,omitempty" xml:"ConfigItemKey,omitempty"`
	ConfigItemValue   *string `json:"ConfigItemValue,omitempty" xml:"ConfigItemValue,omitempty"`
	ConfigScope       *string `json:"ConfigScope,omitempty" xml:"ConfigScope,omitempty"`
	EffectiveActions  *string `json:"EffectiveActions,omitempty" xml:"EffectiveActions,omitempty"`
	EffectiveType     *string `json:"EffectiveType,omitempty" xml:"EffectiveType,omitempty"`
	NodeGroupId       *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	NodeId            *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s UpdateApplicationConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationConfig) SetConfigAction(v string) *UpdateApplicationConfig {
	s.ConfigAction = &v
	return s
}

func (s *UpdateApplicationConfig) SetConfigDescription(v string) *UpdateApplicationConfig {
	s.ConfigDescription = &v
	return s
}

func (s *UpdateApplicationConfig) SetConfigFileName(v string) *UpdateApplicationConfig {
	s.ConfigFileName = &v
	return s
}

func (s *UpdateApplicationConfig) SetConfigItemKey(v string) *UpdateApplicationConfig {
	s.ConfigItemKey = &v
	return s
}

func (s *UpdateApplicationConfig) SetConfigItemValue(v string) *UpdateApplicationConfig {
	s.ConfigItemValue = &v
	return s
}

func (s *UpdateApplicationConfig) SetConfigScope(v string) *UpdateApplicationConfig {
	s.ConfigScope = &v
	return s
}

func (s *UpdateApplicationConfig) SetEffectiveActions(v string) *UpdateApplicationConfig {
	s.EffectiveActions = &v
	return s
}

func (s *UpdateApplicationConfig) SetEffectiveType(v string) *UpdateApplicationConfig {
	s.EffectiveType = &v
	return s
}

func (s *UpdateApplicationConfig) SetNodeGroupId(v string) *UpdateApplicationConfig {
	s.NodeGroupId = &v
	return s
}

func (s *UpdateApplicationConfig) SetNodeId(v string) *UpdateApplicationConfig {
	s.NodeId = &v
	return s
}

type UpdateSpecNodeGroup struct {
	NewInstanceType *string `json:"NewInstanceType,omitempty" xml:"NewInstanceType,omitempty"`
	NodeGroupId     *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s UpdateSpecNodeGroup) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpecNodeGroup) GoString() string {
	return s.String()
}

func (s *UpdateSpecNodeGroup) SetNewInstanceType(v string) *UpdateSpecNodeGroup {
	s.NewInstanceType = &v
	return s
}

func (s *UpdateSpecNodeGroup) SetNodeGroupId(v string) *UpdateSpecNodeGroup {
	s.NodeGroupId = &v
	return s
}

type UpdateSpecNodeGroupParam struct {
	NewInstanceType *string `json:"NewInstanceType,omitempty" xml:"NewInstanceType,omitempty"`
	NodeGroupId     *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s UpdateSpecNodeGroupParam) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpecNodeGroupParam) GoString() string {
	return s.String()
}

func (s *UpdateSpecNodeGroupParam) SetNewInstanceType(v string) *UpdateSpecNodeGroupParam {
	s.NewInstanceType = &v
	return s
}

func (s *UpdateSpecNodeGroupParam) SetNodeGroupId(v string) *UpdateSpecNodeGroupParam {
	s.NodeGroupId = &v
	return s
}

type User struct {
	Group    *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s User) String() string {
	return tea.Prettify(s)
}

func (s User) GoString() string {
	return s.String()
}

func (s *User) SetGroup(v string) *User {
	s.Group = &v
	return s
}

func (s *User) SetPassword(v string) *User {
	s.Password = &v
	return s
}

func (s *User) SetUserId(v string) *User {
	s.UserId = &v
	return s
}

func (s *User) SetUserName(v string) *User {
	s.UserName = &v
	return s
}

func (s *User) SetUserType(v string) *User {
	s.UserType = &v
	return s
}

type UserParam struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UserParam) String() string {
	return tea.Prettify(s)
}

func (s UserParam) GoString() string {
	return s.String()
}

func (s *UserParam) SetPassword(v string) *UserParam {
	s.Password = &v
	return s
}

func (s *UserParam) SetUserId(v string) *UserParam {
	s.UserId = &v
	return s
}

func (s *UserParam) SetUserName(v string) *UserParam {
	s.UserName = &v
	return s
}

type CreateClusterRequest struct {
	ApplicationConfigs []*ApplicationConfig `json:"ApplicationConfigs,omitempty" xml:"ApplicationConfigs,omitempty" type:"Repeated"`
	Applications       []*Application       `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	BootstrapScripts   []*Script            `json:"BootstrapScripts,omitempty" xml:"BootstrapScripts,omitempty" type:"Repeated"`
	ClientToken        *string              `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClusterName        *string              `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterType        *string              `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	DeployMode         *string              `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	NodeAttributes     *NodeAttributes      `json:"NodeAttributes,omitempty" xml:"NodeAttributes,omitempty"`
	NodeGroups         []*NodeGroupConfig   `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
	PaymentType        *string              `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	RegionId           *string              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReleaseVersion     *string              `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	ResourceGroupId    *string              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityMode       *string              `json:"SecurityMode,omitempty" xml:"SecurityMode,omitempty"`
	SubscriptionConfig *SubscriptionConfig  `json:"SubscriptionConfig,omitempty" xml:"SubscriptionConfig,omitempty"`
	Tags               []*Tag               `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetApplicationConfigs(v []*ApplicationConfig) *CreateClusterRequest {
	s.ApplicationConfigs = v
	return s
}

func (s *CreateClusterRequest) SetApplications(v []*Application) *CreateClusterRequest {
	s.Applications = v
	return s
}

func (s *CreateClusterRequest) SetBootstrapScripts(v []*Script) *CreateClusterRequest {
	s.BootstrapScripts = v
	return s
}

func (s *CreateClusterRequest) SetClientToken(v string) *CreateClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateClusterRequest) SetClusterName(v string) *CreateClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterRequest) SetClusterType(v string) *CreateClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterRequest) SetDeployMode(v string) *CreateClusterRequest {
	s.DeployMode = &v
	return s
}

func (s *CreateClusterRequest) SetNodeAttributes(v *NodeAttributes) *CreateClusterRequest {
	s.NodeAttributes = v
	return s
}

func (s *CreateClusterRequest) SetNodeGroups(v []*NodeGroupConfig) *CreateClusterRequest {
	s.NodeGroups = v
	return s
}

func (s *CreateClusterRequest) SetPaymentType(v string) *CreateClusterRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateClusterRequest) SetRegionId(v string) *CreateClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateClusterRequest) SetReleaseVersion(v string) *CreateClusterRequest {
	s.ReleaseVersion = &v
	return s
}

func (s *CreateClusterRequest) SetResourceGroupId(v string) *CreateClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetSecurityMode(v string) *CreateClusterRequest {
	s.SecurityMode = &v
	return s
}

func (s *CreateClusterRequest) SetSubscriptionConfig(v *SubscriptionConfig) *CreateClusterRequest {
	s.SubscriptionConfig = v
	return s
}

func (s *CreateClusterRequest) SetTags(v []*Tag) *CreateClusterRequest {
	s.Tags = v
	return s
}

type CreateClusterResponseBody struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterResponseBody) SetOperationId(v string) *CreateClusterResponseBody {
	s.OperationId = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

type CreateClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) SetHeaders(v map[string]*string) *CreateClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterResponse) SetStatusCode(v int32) *CreateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type DecreaseNodesRequest struct {
	ClusterId         *string   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DecreaseNodeCount *int32    `json:"DecreaseNodeCount,omitempty" xml:"DecreaseNodeCount,omitempty"`
	NodeGroupId       *string   `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	NodeIds           []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	RegionId          *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DecreaseNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DecreaseNodesRequest) GoString() string {
	return s.String()
}

func (s *DecreaseNodesRequest) SetClusterId(v string) *DecreaseNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *DecreaseNodesRequest) SetDecreaseNodeCount(v int32) *DecreaseNodesRequest {
	s.DecreaseNodeCount = &v
	return s
}

func (s *DecreaseNodesRequest) SetNodeGroupId(v string) *DecreaseNodesRequest {
	s.NodeGroupId = &v
	return s
}

func (s *DecreaseNodesRequest) SetNodeIds(v []*string) *DecreaseNodesRequest {
	s.NodeIds = v
	return s
}

func (s *DecreaseNodesRequest) SetRegionId(v string) *DecreaseNodesRequest {
	s.RegionId = &v
	return s
}

type DecreaseNodesResponseBody struct {
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DecreaseNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DecreaseNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DecreaseNodesResponseBody) SetOperationId(v string) *DecreaseNodesResponseBody {
	s.OperationId = &v
	return s
}

func (s *DecreaseNodesResponseBody) SetRequestId(v string) *DecreaseNodesResponseBody {
	s.RequestId = &v
	return s
}

type DecreaseNodesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DecreaseNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DecreaseNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DecreaseNodesResponse) GoString() string {
	return s.String()
}

func (s *DecreaseNodesResponse) SetHeaders(v map[string]*string) *DecreaseNodesResponse {
	s.Headers = v
	return s
}

func (s *DecreaseNodesResponse) SetStatusCode(v int32) *DecreaseNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DecreaseNodesResponse) SetBody(v *DecreaseNodesResponseBody) *DecreaseNodesResponse {
	s.Body = v
	return s
}

type DeleteClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetClusterId(v string) *DeleteClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterRequest) SetRegionId(v string) *DeleteClusterRequest {
	s.RegionId = &v
	return s
}

type DeleteClusterResponseBody struct {
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) SetOperationId(v string) *DeleteClusterResponseBody {
	s.OperationId = &v
	return s
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) SetHeaders(v map[string]*string) *DeleteClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterResponse) SetStatusCode(v int32) *DeleteClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClusterResponse) SetBody(v *DeleteClusterResponseBody) *DeleteClusterResponse {
	s.Body = v
	return s
}

type GetClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRequest) GoString() string {
	return s.String()
}

func (s *GetClusterRequest) SetClusterId(v string) *GetClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *GetClusterRequest) SetRegionId(v string) *GetClusterRequest {
	s.RegionId = &v
	return s
}

type GetClusterResponseBody struct {
	Cluster   *Cluster `json:"Cluster,omitempty" xml:"Cluster,omitempty"`
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBody) SetCluster(v *Cluster) *GetClusterResponseBody {
	s.Cluster = v
	return s
}

func (s *GetClusterResponseBody) SetRequestId(v string) *GetClusterResponseBody {
	s.RequestId = &v
	return s
}

type GetClusterResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponse) GoString() string {
	return s.String()
}

func (s *GetClusterResponse) SetHeaders(v map[string]*string) *GetClusterResponse {
	s.Headers = v
	return s
}

func (s *GetClusterResponse) SetStatusCode(v int32) *GetClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterResponse) SetBody(v *GetClusterResponseBody) *GetClusterResponse {
	s.Body = v
	return s
}

type GetNodeGroupRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetNodeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *GetNodeGroupRequest) SetClusterId(v string) *GetNodeGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *GetNodeGroupRequest) SetNodeGroupId(v string) *GetNodeGroupRequest {
	s.NodeGroupId = &v
	return s
}

func (s *GetNodeGroupRequest) SetRegionId(v string) *GetNodeGroupRequest {
	s.RegionId = &v
	return s
}

type GetNodeGroupResponseBody struct {
	NodeGroup *NodeGroup `json:"NodeGroup,omitempty" xml:"NodeGroup,omitempty"`
	RequestId *string    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNodeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeGroupResponseBody) SetNodeGroup(v *NodeGroup) *GetNodeGroupResponseBody {
	s.NodeGroup = v
	return s
}

func (s *GetNodeGroupResponseBody) SetRequestId(v string) *GetNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

type GetNodeGroupResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNodeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *GetNodeGroupResponse) SetHeaders(v map[string]*string) *GetNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *GetNodeGroupResponse) SetStatusCode(v int32) *GetNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeGroupResponse) SetBody(v *GetNodeGroupResponseBody) *GetNodeGroupResponse {
	s.Body = v
	return s
}

type GetOperationRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetOperationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOperationRequest) GoString() string {
	return s.String()
}

func (s *GetOperationRequest) SetClusterId(v string) *GetOperationRequest {
	s.ClusterId = &v
	return s
}

func (s *GetOperationRequest) SetOperationId(v string) *GetOperationRequest {
	s.OperationId = &v
	return s
}

func (s *GetOperationRequest) SetRegionId(v string) *GetOperationRequest {
	s.RegionId = &v
	return s
}

type GetOperationResponseBody struct {
	Operation *Operation `json:"Operation,omitempty" xml:"Operation,omitempty"`
	RequestId *string    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOperationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOperationResponseBody) GoString() string {
	return s.String()
}

func (s *GetOperationResponseBody) SetOperation(v *Operation) *GetOperationResponseBody {
	s.Operation = v
	return s
}

func (s *GetOperationResponseBody) SetRequestId(v string) *GetOperationResponseBody {
	s.RequestId = &v
	return s
}

type GetOperationResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOperationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOperationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOperationResponse) GoString() string {
	return s.String()
}

func (s *GetOperationResponse) SetHeaders(v map[string]*string) *GetOperationResponse {
	s.Headers = v
	return s
}

func (s *GetOperationResponse) SetStatusCode(v int32) *GetOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOperationResponse) SetBody(v *GetOperationResponseBody) *GetOperationResponse {
	s.Body = v
	return s
}

type IncreaseNodesRequest struct {
	ApplicationConfigs  []*ApplicationConfig `json:"ApplicationConfigs,omitempty" xml:"ApplicationConfigs,omitempty" type:"Repeated"`
	AutoPayOrder        *bool                `json:"AutoPayOrder,omitempty" xml:"AutoPayOrder,omitempty"`
	ClusterId           *string              `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	IncreaseNodeCount   *int32               `json:"IncreaseNodeCount,omitempty" xml:"IncreaseNodeCount,omitempty"`
	NodeGroupId         *string              `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	PaymentDuration     *int32               `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	PaymentDurationUnit *string              `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	RegionId            *string              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s IncreaseNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s IncreaseNodesRequest) GoString() string {
	return s.String()
}

func (s *IncreaseNodesRequest) SetApplicationConfigs(v []*ApplicationConfig) *IncreaseNodesRequest {
	s.ApplicationConfigs = v
	return s
}

func (s *IncreaseNodesRequest) SetAutoPayOrder(v bool) *IncreaseNodesRequest {
	s.AutoPayOrder = &v
	return s
}

func (s *IncreaseNodesRequest) SetClusterId(v string) *IncreaseNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *IncreaseNodesRequest) SetIncreaseNodeCount(v int32) *IncreaseNodesRequest {
	s.IncreaseNodeCount = &v
	return s
}

func (s *IncreaseNodesRequest) SetNodeGroupId(v string) *IncreaseNodesRequest {
	s.NodeGroupId = &v
	return s
}

func (s *IncreaseNodesRequest) SetPaymentDuration(v int32) *IncreaseNodesRequest {
	s.PaymentDuration = &v
	return s
}

func (s *IncreaseNodesRequest) SetPaymentDurationUnit(v string) *IncreaseNodesRequest {
	s.PaymentDurationUnit = &v
	return s
}

func (s *IncreaseNodesRequest) SetRegionId(v string) *IncreaseNodesRequest {
	s.RegionId = &v
	return s
}

type IncreaseNodesResponseBody struct {
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IncreaseNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IncreaseNodesResponseBody) GoString() string {
	return s.String()
}

func (s *IncreaseNodesResponseBody) SetOperationId(v string) *IncreaseNodesResponseBody {
	s.OperationId = &v
	return s
}

func (s *IncreaseNodesResponseBody) SetRequestId(v string) *IncreaseNodesResponseBody {
	s.RequestId = &v
	return s
}

type IncreaseNodesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IncreaseNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IncreaseNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s IncreaseNodesResponse) GoString() string {
	return s.String()
}

func (s *IncreaseNodesResponse) SetHeaders(v map[string]*string) *IncreaseNodesResponse {
	s.Headers = v
	return s
}

func (s *IncreaseNodesResponse) SetStatusCode(v int32) *IncreaseNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *IncreaseNodesResponse) SetBody(v *IncreaseNodesResponseBody) *IncreaseNodesResponse {
	s.Body = v
	return s
}

type JoinResourceGroupRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s JoinResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *JoinResourceGroupRequest) SetRegionId(v string) *JoinResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *JoinResourceGroupRequest) SetResourceGroupId(v string) *JoinResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *JoinResourceGroupRequest) SetResourceId(v string) *JoinResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *JoinResourceGroupRequest) SetResourceType(v string) *JoinResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type JoinResourceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s JoinResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JoinResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *JoinResourceGroupResponseBody) SetRequestId(v string) *JoinResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type JoinResourceGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *JoinResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s JoinResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *JoinResourceGroupResponse) SetHeaders(v map[string]*string) *JoinResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *JoinResourceGroupResponse) SetStatusCode(v int32) *JoinResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinResourceGroupResponse) SetBody(v *JoinResourceGroupResponseBody) *JoinResourceGroupResponse {
	s.Body = v
	return s
}

type ListNodeGroupsRequest struct {
	ClusterId       *string   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	MaxResults      *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	NodeGroupIds    []*string `json:"NodeGroupIds,omitempty" xml:"NodeGroupIds,omitempty" type:"Repeated"`
	NodeGroupNames  []*string `json:"NodeGroupNames,omitempty" xml:"NodeGroupNames,omitempty" type:"Repeated"`
	NodeGroupStates []*string `json:"NodeGroupStates,omitempty" xml:"NodeGroupStates,omitempty" type:"Repeated"`
	NodeGroupTypes  []*string `json:"NodeGroupTypes,omitempty" xml:"NodeGroupTypes,omitempty" type:"Repeated"`
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListNodeGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsRequest) SetClusterId(v string) *ListNodeGroupsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodeGroupsRequest) SetMaxResults(v int32) *ListNodeGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNodeGroupsRequest) SetNextToken(v string) *ListNodeGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListNodeGroupsRequest) SetNodeGroupIds(v []*string) *ListNodeGroupsRequest {
	s.NodeGroupIds = v
	return s
}

func (s *ListNodeGroupsRequest) SetNodeGroupNames(v []*string) *ListNodeGroupsRequest {
	s.NodeGroupNames = v
	return s
}

func (s *ListNodeGroupsRequest) SetNodeGroupStates(v []*string) *ListNodeGroupsRequest {
	s.NodeGroupStates = v
	return s
}

func (s *ListNodeGroupsRequest) SetNodeGroupTypes(v []*string) *ListNodeGroupsRequest {
	s.NodeGroupTypes = v
	return s
}

func (s *ListNodeGroupsRequest) SetRegionId(v string) *ListNodeGroupsRequest {
	s.RegionId = &v
	return s
}

type ListNodeGroupsResponseBody struct {
	MaxResults *int32       `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string      `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	NodeGroups []*NodeGroup `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
	RequestId  *string      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodeGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponseBody) SetMaxResults(v int32) *ListNodeGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNodeGroupsResponseBody) SetNextToken(v string) *ListNodeGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNodeGroupsResponseBody) SetNodeGroups(v []*NodeGroup) *ListNodeGroupsResponseBody {
	s.NodeGroups = v
	return s
}

func (s *ListNodeGroupsResponseBody) SetRequestId(v string) *ListNodeGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeGroupsResponseBody) SetTotalCount(v int32) *ListNodeGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListNodeGroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodeGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodeGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponse) SetHeaders(v map[string]*string) *ListNodeGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListNodeGroupsResponse) SetStatusCode(v int32) *ListNodeGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeGroupsResponse) SetBody(v *ListNodeGroupsResponseBody) *ListNodeGroupsResponse {
	s.Body = v
	return s
}

type ListNodesRequest struct {
	ClusterId    *string   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	MaxResults   *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	NodeGroupIds []*string `json:"NodeGroupIds,omitempty" xml:"NodeGroupIds,omitempty" type:"Repeated"`
	NodeIds      []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	NodeNames    []*string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty" type:"Repeated"`
	NodeStates   []*string `json:"NodeStates,omitempty" xml:"NodeStates,omitempty" type:"Repeated"`
	PrivateIps   []*string `json:"PrivateIps,omitempty" xml:"PrivateIps,omitempty" type:"Repeated"`
	PublicIps    []*string `json:"PublicIps,omitempty" xml:"PublicIps,omitempty" type:"Repeated"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) SetClusterId(v string) *ListNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodesRequest) SetMaxResults(v int32) *ListNodesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNodesRequest) SetNextToken(v string) *ListNodesRequest {
	s.NextToken = &v
	return s
}

func (s *ListNodesRequest) SetNodeGroupIds(v []*string) *ListNodesRequest {
	s.NodeGroupIds = v
	return s
}

func (s *ListNodesRequest) SetNodeIds(v []*string) *ListNodesRequest {
	s.NodeIds = v
	return s
}

func (s *ListNodesRequest) SetNodeNames(v []*string) *ListNodesRequest {
	s.NodeNames = v
	return s
}

func (s *ListNodesRequest) SetNodeStates(v []*string) *ListNodesRequest {
	s.NodeStates = v
	return s
}

func (s *ListNodesRequest) SetPrivateIps(v []*string) *ListNodesRequest {
	s.PrivateIps = v
	return s
}

func (s *ListNodesRequest) SetPublicIps(v []*string) *ListNodesRequest {
	s.PublicIps = v
	return s
}

func (s *ListNodesRequest) SetRegionId(v string) *ListNodesRequest {
	s.RegionId = &v
	return s
}

type ListNodesResponseBody struct {
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Nodes      []*Node `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) SetMaxResults(v int32) *ListNodesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNodesResponseBody) SetNextToken(v string) *ListNodesResponseBody {
	s.NextToken = &v
	return s
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
		"cn-beijing":            tea.String("emr.aliyuncs.com"),
		"cn-hangzhou":           tea.String("emr.aliyuncs.com"),
		"cn-shanghai":           tea.String("emr.aliyuncs.com"),
		"cn-shenzhen":           tea.String("emr.aliyuncs.com"),
		"ap-southeast-1":        tea.String("emr.aliyuncs.com"),
		"us-west-1":             tea.String("emr.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("emr.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("emr.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("emr.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("emr"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

/**
 * 调用CreateCluster创建集群。
 */
func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationConfigs)) {
		query["ApplicationConfigs"] = request.ApplicationConfigs
	}

	if !tea.BoolValue(util.IsUnset(request.Applications)) {
		query["Applications"] = request.Applications
	}

	if !tea.BoolValue(util.IsUnset(request.BootstrapScripts)) {
		query["BootstrapScripts"] = request.BootstrapScripts
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterType)) {
		query["ClusterType"] = request.ClusterType
	}

	if !tea.BoolValue(util.IsUnset(request.DeployMode)) {
		query["DeployMode"] = request.DeployMode
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.NodeAttributes))) {
		query["NodeAttributes"] = request.NodeAttributes
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroups)) {
		query["NodeGroups"] = request.NodeGroups
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentType)) {
		query["PaymentType"] = request.PaymentType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseVersion)) {
		query["ReleaseVersion"] = request.ReleaseVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityMode)) {
		query["SecurityMode"] = request.SecurityMode
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubscriptionConfig))) {
		query["SubscriptionConfig"] = request.SubscriptionConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCluster"),
		Version:     tea.String("2021-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 缩容节点。
 */
func (client *Client) DecreaseNodesWithOptions(request *DecreaseNodesRequest, runtime *util.RuntimeOptions) (_result *DecreaseNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DecreaseNodeCount)) {
		query["DecreaseNodeCount"] = request.DecreaseNodeCount
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIds)) {
		query["NodeIds"] = request.NodeIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DecreaseNodes"),
		Version:     tea.String("2021-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DecreaseNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DecreaseNodes(request *DecreaseNodesRequest) (_result *DecreaseNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DecreaseNodesResponse{}
	_body, _err := client.DecreaseNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 删除集群。
 */
func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCluster"),
		Version:     tea.String("2021-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 调用GetCluster获取集群详情。
 */
func (client *Client) GetClusterWithOptions(request *GetClusterRequest, runtime *util.RuntimeOptions) (_result *GetClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCluster"),
		Version:     tea.String("2021-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCluster(request *GetClusterRequest) (_result *GetClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClusterResponse{}
	_body, _err := client.GetClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 获取节点组详情。
 */
func (client *Client) GetNodeGroupWithOptions(request *GetNodeGroupRequest, runtime *util.RuntimeOptions) (_result *GetNodeGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNodeGroup"),
		Version:     tea.String("2021-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNodeGroup(request *GetNodeGroupRequest) (_result *GetNodeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNodeGroupResponse{}
	_body, _err := client.GetNodeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 获取操作详情。
 */
func (client *Client) GetOperationWithOptions(request *GetOperationRequest, runtime *util.RuntimeOptions) (_result *GetOperationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OperationId)) {
		query["OperationId"] = request.OperationId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOperation"),
		Version:     tea.String("2021-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOperation(request *GetOperationRequest) (_result *GetOperationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOperationResponse{}
	_body, _err := client.GetOperationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IncreaseNodesWithOptions(request *IncreaseNodesRequest, runtime *util.RuntimeOptions) (_result *IncreaseNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationConfigs)) {
		query["ApplicationConfigs"] = request.ApplicationConfigs
	}

	if !tea.BoolValue(util.IsUnset(request.AutoPayOrder)) {
		query["AutoPayOrder"] = request.AutoPayOrder
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.IncreaseNodeCount)) {
		query["IncreaseNodeCount"] = request.IncreaseNodeCount
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentDuration)) {
		query["PaymentDuration"] = request.PaymentDuration
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentDurationUnit)) {
		query["PaymentDurationUnit"] = request.PaymentDurationUnit
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("IncreaseNodes"),
		Version:     tea.String("2021-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &IncreaseNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IncreaseNodes(request *IncreaseNodesRequest) (_result *IncreaseNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IncreaseNodesResponse{}
	_body, _err := client.IncreaseNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JoinResourceGroupWithOptions(request *JoinResourceGroupRequest, runtime *util.RuntimeOptions) (_result *JoinResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("JoinResourceGroup"),
		Version:     tea.String("2021-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &JoinResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JoinResourceGroup(request *JoinResourceGroupRequest) (_result *JoinResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JoinResourceGroupResponse{}
	_body, _err := client.JoinResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 查询节点组。
 */
func (client *Client) ListNodeGroupsWithOptions(request *ListNodeGroupsRequest, runtime *util.RuntimeOptions) (_result *ListNodeGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupIds)) {
		query["NodeGroupIds"] = request.NodeGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupNames)) {
		query["NodeGroupNames"] = request.NodeGroupNames
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupStates)) {
		query["NodeGroupStates"] = request.NodeGroupStates
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupTypes)) {
		query["NodeGroupTypes"] = request.NodeGroupTypes
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodeGroups"),
		Version:     tea.String("2021-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodeGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodeGroups(request *ListNodeGroupsRequest) (_result *ListNodeGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodeGroupsResponse{}
	_body, _err := client.ListNodeGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 查询节点。
 */
func (client *Client) ListNodesWithOptions(request *ListNodesRequest, runtime *util.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupIds)) {
		query["NodeGroupIds"] = request.NodeGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIds)) {
		query["NodeIds"] = request.NodeIds
	}

	if !tea.BoolValue(util.IsUnset(request.NodeNames)) {
		query["NodeNames"] = request.NodeNames
	}

	if !tea.BoolValue(util.IsUnset(request.NodeStates)) {
		query["NodeStates"] = request.NodeStates
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateIps)) {
		query["PrivateIps"] = request.PrivateIps
	}

	if !tea.BoolValue(util.IsUnset(request.PublicIps)) {
		query["PublicIps"] = request.PublicIps
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodes"),
		Version:     tea.String("2021-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodes(request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
