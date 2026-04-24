// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertFlowDSLData interface {
	dara.Model
	String() string
	GoString() string
	SetConversionPlan(v *ConvertFlowDSLDataConversionPlan) *ConvertFlowDSLData
	GetConversionPlan() *ConvertFlowDSLDataConversionPlan
	SetFlow(v *ConvertFlowDSLDataFlow) *ConvertFlowDSLData
	GetFlow() *ConvertFlowDSLDataFlow
	SetPluginErrors(v []*ConvertFlowDSLDataPluginErrors) *ConvertFlowDSLData
	GetPluginErrors() []*ConvertFlowDSLDataPluginErrors
	SetToolsetInstallations(v []*ConvertFlowDSLDataToolsetInstallations) *ConvertFlowDSLData
	GetToolsetInstallations() []*ConvertFlowDSLDataToolsetInstallations
}

type ConvertFlowDSLData struct {
	// DSL兼容性分析和转换计划
	ConversionPlan *ConvertFlowDSLDataConversionPlan `json:"conversionPlan,omitempty" xml:"conversionPlan,omitempty" type:"Struct"`
	// 转换后的AgentRun Flow配置信息
	//
	// This parameter is required.
	Flow *ConvertFlowDSLDataFlow `json:"flow,omitempty" xml:"flow,omitempty" type:"Struct"`
	// 插件识别或转换过程中的错误信息
	PluginErrors []*ConvertFlowDSLDataPluginErrors `json:"pluginErrors" xml:"pluginErrors" type:"Repeated"`
	// 需要安装的Toolset配置列表
	ToolsetInstallations []*ConvertFlowDSLDataToolsetInstallations `json:"toolsetInstallations" xml:"toolsetInstallations" type:"Repeated"`
}

func (s ConvertFlowDSLData) String() string {
	return dara.Prettify(s)
}

func (s ConvertFlowDSLData) GoString() string {
	return s.String()
}

func (s *ConvertFlowDSLData) GetConversionPlan() *ConvertFlowDSLDataConversionPlan {
	return s.ConversionPlan
}

func (s *ConvertFlowDSLData) GetFlow() *ConvertFlowDSLDataFlow {
	return s.Flow
}

func (s *ConvertFlowDSLData) GetPluginErrors() []*ConvertFlowDSLDataPluginErrors {
	return s.PluginErrors
}

func (s *ConvertFlowDSLData) GetToolsetInstallations() []*ConvertFlowDSLDataToolsetInstallations {
	return s.ToolsetInstallations
}

func (s *ConvertFlowDSLData) SetConversionPlan(v *ConvertFlowDSLDataConversionPlan) *ConvertFlowDSLData {
	s.ConversionPlan = v
	return s
}

func (s *ConvertFlowDSLData) SetFlow(v *ConvertFlowDSLDataFlow) *ConvertFlowDSLData {
	s.Flow = v
	return s
}

func (s *ConvertFlowDSLData) SetPluginErrors(v []*ConvertFlowDSLDataPluginErrors) *ConvertFlowDSLData {
	s.PluginErrors = v
	return s
}

func (s *ConvertFlowDSLData) SetToolsetInstallations(v []*ConvertFlowDSLDataToolsetInstallations) *ConvertFlowDSLData {
	s.ToolsetInstallations = v
	return s
}

func (s *ConvertFlowDSLData) Validate() error {
	if s.ConversionPlan != nil {
		if err := s.ConversionPlan.Validate(); err != nil {
			return err
		}
	}
	if s.Flow != nil {
		if err := s.Flow.Validate(); err != nil {
			return err
		}
	}
	if s.PluginErrors != nil {
		for _, item := range s.PluginErrors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ToolsetInstallations != nil {
		for _, item := range s.ToolsetInstallations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ConvertFlowDSLDataConversionPlan struct {
	// 节点兼容性问题详情
	Issues []*ConvertFlowDSLDataConversionPlanIssues `json:"issues" xml:"issues" type:"Repeated"`
	// 节点兼容性统计摘要
	//
	// This parameter is required.
	Summary *ConvertFlowDSLDataConversionPlanSummary `json:"summary,omitempty" xml:"summary,omitempty" type:"Struct"`
}

func (s ConvertFlowDSLDataConversionPlan) String() string {
	return dara.Prettify(s)
}

func (s ConvertFlowDSLDataConversionPlan) GoString() string {
	return s.String()
}

func (s *ConvertFlowDSLDataConversionPlan) GetIssues() []*ConvertFlowDSLDataConversionPlanIssues {
	return s.Issues
}

func (s *ConvertFlowDSLDataConversionPlan) GetSummary() *ConvertFlowDSLDataConversionPlanSummary {
	return s.Summary
}

func (s *ConvertFlowDSLDataConversionPlan) SetIssues(v []*ConvertFlowDSLDataConversionPlanIssues) *ConvertFlowDSLDataConversionPlan {
	s.Issues = v
	return s
}

func (s *ConvertFlowDSLDataConversionPlan) SetSummary(v *ConvertFlowDSLDataConversionPlanSummary) *ConvertFlowDSLDataConversionPlan {
	s.Summary = v
	return s
}

func (s *ConvertFlowDSLDataConversionPlan) Validate() error {
	if s.Issues != nil {
		for _, item := range s.Issues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Summary != nil {
		if err := s.Summary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConvertFlowDSLDataConversionPlanIssues struct {
	// 问题描述
	//
	// This parameter is required.
	//
	// example:
	//
	// Tool node requires Toolset installation
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 问题的详细信息（JSON对象）
	Details map[string]interface{} `json:"details" xml:"details"`
	// 问题类型：needs_config, needs_conversion, unsupported
	//
	// This parameter is required.
	//
	// example:
	//
	// needs_config
	IssueType *string `json:"issueType,omitempty" xml:"issueType,omitempty"`
	// 问题严重程度：info, warning, error
	//
	// This parameter is required.
	//
	// example:
	//
	// warning
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// Dify DSL中的节点标识符
	//
	// This parameter is required.
	//
	// example:
	//
	// 1760514991682
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 节点显示名称
	//
	// example:
	//
	// Google Search
	NodeTitle *string `json:"nodeTitle,omitempty" xml:"nodeTitle,omitempty"`
	// 节点类型
	//
	// This parameter is required.
	//
	// example:
	//
	// tool
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
	// 解决建议
	//
	// example:
	//
	// Install Toolset \"google\" before using this flow
	Suggestion *string `json:"suggestion,omitempty" xml:"suggestion,omitempty"`
}

func (s ConvertFlowDSLDataConversionPlanIssues) String() string {
	return dara.Prettify(s)
}

func (s ConvertFlowDSLDataConversionPlanIssues) GoString() string {
	return s.String()
}

func (s *ConvertFlowDSLDataConversionPlanIssues) GetDescription() *string {
	return s.Description
}

func (s *ConvertFlowDSLDataConversionPlanIssues) GetDetails() map[string]interface{} {
	return s.Details
}

func (s *ConvertFlowDSLDataConversionPlanIssues) GetIssueType() *string {
	return s.IssueType
}

func (s *ConvertFlowDSLDataConversionPlanIssues) GetLevel() *string {
	return s.Level
}

func (s *ConvertFlowDSLDataConversionPlanIssues) GetNodeId() *string {
	return s.NodeId
}

func (s *ConvertFlowDSLDataConversionPlanIssues) GetNodeTitle() *string {
	return s.NodeTitle
}

func (s *ConvertFlowDSLDataConversionPlanIssues) GetNodeType() *string {
	return s.NodeType
}

func (s *ConvertFlowDSLDataConversionPlanIssues) GetSuggestion() *string {
	return s.Suggestion
}

func (s *ConvertFlowDSLDataConversionPlanIssues) SetDescription(v string) *ConvertFlowDSLDataConversionPlanIssues {
	s.Description = &v
	return s
}

func (s *ConvertFlowDSLDataConversionPlanIssues) SetDetails(v map[string]interface{}) *ConvertFlowDSLDataConversionPlanIssues {
	s.Details = v
	return s
}

func (s *ConvertFlowDSLDataConversionPlanIssues) SetIssueType(v string) *ConvertFlowDSLDataConversionPlanIssues {
	s.IssueType = &v
	return s
}

func (s *ConvertFlowDSLDataConversionPlanIssues) SetLevel(v string) *ConvertFlowDSLDataConversionPlanIssues {
	s.Level = &v
	return s
}

func (s *ConvertFlowDSLDataConversionPlanIssues) SetNodeId(v string) *ConvertFlowDSLDataConversionPlanIssues {
	s.NodeId = &v
	return s
}

func (s *ConvertFlowDSLDataConversionPlanIssues) SetNodeTitle(v string) *ConvertFlowDSLDataConversionPlanIssues {
	s.NodeTitle = &v
	return s
}

func (s *ConvertFlowDSLDataConversionPlanIssues) SetNodeType(v string) *ConvertFlowDSLDataConversionPlanIssues {
	s.NodeType = &v
	return s
}

func (s *ConvertFlowDSLDataConversionPlanIssues) SetSuggestion(v string) *ConvertFlowDSLDataConversionPlanIssues {
	s.Suggestion = &v
	return s
}

func (s *ConvertFlowDSLDataConversionPlanIssues) Validate() error {
	return dara.Validate(s)
}

type ConvertFlowDSLDataConversionPlanSummary struct {
	// 完全兼容的节点数
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	CompatibleNodes *int32 `json:"compatibleNodes,omitempty" xml:"compatibleNodes,omitempty"`
	// 需要手动配置的节点数
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	NodesNeedConfig *int32 `json:"nodesNeedConfig,omitempty" xml:"nodesNeedConfig,omitempty"`
	// 需要特殊转换处理的节点数
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	NodesNeedConversion *int32 `json:"nodesNeedConversion,omitempty" xml:"nodesNeedConversion,omitempty"`
	// Dify DSL中的总节点数
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	TotalNodes *int32 `json:"totalNodes,omitempty" xml:"totalNodes,omitempty"`
	// 不支持的节点数
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UnsupportedNodes *int32 `json:"unsupportedNodes,omitempty" xml:"unsupportedNodes,omitempty"`
}

func (s ConvertFlowDSLDataConversionPlanSummary) String() string {
	return dara.Prettify(s)
}

func (s ConvertFlowDSLDataConversionPlanSummary) GoString() string {
	return s.String()
}

func (s *ConvertFlowDSLDataConversionPlanSummary) GetCompatibleNodes() *int32 {
	return s.CompatibleNodes
}

func (s *ConvertFlowDSLDataConversionPlanSummary) GetNodesNeedConfig() *int32 {
	return s.NodesNeedConfig
}

func (s *ConvertFlowDSLDataConversionPlanSummary) GetNodesNeedConversion() *int32 {
	return s.NodesNeedConversion
}

func (s *ConvertFlowDSLDataConversionPlanSummary) GetTotalNodes() *int32 {
	return s.TotalNodes
}

func (s *ConvertFlowDSLDataConversionPlanSummary) GetUnsupportedNodes() *int32 {
	return s.UnsupportedNodes
}

func (s *ConvertFlowDSLDataConversionPlanSummary) SetCompatibleNodes(v int32) *ConvertFlowDSLDataConversionPlanSummary {
	s.CompatibleNodes = &v
	return s
}

func (s *ConvertFlowDSLDataConversionPlanSummary) SetNodesNeedConfig(v int32) *ConvertFlowDSLDataConversionPlanSummary {
	s.NodesNeedConfig = &v
	return s
}

func (s *ConvertFlowDSLDataConversionPlanSummary) SetNodesNeedConversion(v int32) *ConvertFlowDSLDataConversionPlanSummary {
	s.NodesNeedConversion = &v
	return s
}

func (s *ConvertFlowDSLDataConversionPlanSummary) SetTotalNodes(v int32) *ConvertFlowDSLDataConversionPlanSummary {
	s.TotalNodes = &v
	return s
}

func (s *ConvertFlowDSLDataConversionPlanSummary) SetUnsupportedNodes(v int32) *ConvertFlowDSLDataConversionPlanSummary {
	s.UnsupportedNodes = &v
	return s
}

func (s *ConvertFlowDSLDataConversionPlanSummary) Validate() error {
	return dara.Validate(s)
}

type ConvertFlowDSLDataFlow struct {
	// 工作流的FnF State Machine定义（YAML格式）
	//
	// This parameter is required.
	Definition *string `json:"definition,omitempty" xml:"definition,omitempty"`
	// 工作流的描述信息
	//
	// example:
	//
	// Converted from external workflow
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 工作流的环境变量配置
	EnvironmentConfiguration *EnvironmentConfiguration `json:"environmentConfiguration,omitempty" xml:"environmentConfiguration,omitempty"`
	// 转换后的工作流名称
	//
	// This parameter is required.
	//
	// example:
	//
	// dify-flow
	FlowName *string `json:"flowName,omitempty" xml:"flowName,omitempty"`
}

func (s ConvertFlowDSLDataFlow) String() string {
	return dara.Prettify(s)
}

func (s ConvertFlowDSLDataFlow) GoString() string {
	return s.String()
}

func (s *ConvertFlowDSLDataFlow) GetDefinition() *string {
	return s.Definition
}

func (s *ConvertFlowDSLDataFlow) GetDescription() *string {
	return s.Description
}

func (s *ConvertFlowDSLDataFlow) GetEnvironmentConfiguration() *EnvironmentConfiguration {
	return s.EnvironmentConfiguration
}

func (s *ConvertFlowDSLDataFlow) GetFlowName() *string {
	return s.FlowName
}

func (s *ConvertFlowDSLDataFlow) SetDefinition(v string) *ConvertFlowDSLDataFlow {
	s.Definition = &v
	return s
}

func (s *ConvertFlowDSLDataFlow) SetDescription(v string) *ConvertFlowDSLDataFlow {
	s.Description = &v
	return s
}

func (s *ConvertFlowDSLDataFlow) SetEnvironmentConfiguration(v *EnvironmentConfiguration) *ConvertFlowDSLDataFlow {
	s.EnvironmentConfiguration = v
	return s
}

func (s *ConvertFlowDSLDataFlow) SetFlowName(v string) *ConvertFlowDSLDataFlow {
	s.FlowName = &v
	return s
}

func (s *ConvertFlowDSLDataFlow) Validate() error {
	if s.EnvironmentConfiguration != nil {
		if err := s.EnvironmentConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConvertFlowDSLDataPluginErrors struct {
	// 相关节点的标识符
	//
	// example:
	//
	// 1760514996015
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 插件提供商名称
	//
	// This parameter is required.
	//
	// example:
	//
	// langgenius
	ProviderName *string `json:"providerName,omitempty" xml:"providerName,omitempty"`
	// 错误原因描述
	//
	// This parameter is required.
	//
	// example:
	//
	// Plugin not found in marketplace
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 工具名称
	//
	// example:
	//
	// google_search
	ToolName *string `json:"toolName,omitempty" xml:"toolName,omitempty"`
}

func (s ConvertFlowDSLDataPluginErrors) String() string {
	return dara.Prettify(s)
}

func (s ConvertFlowDSLDataPluginErrors) GoString() string {
	return s.String()
}

func (s *ConvertFlowDSLDataPluginErrors) GetNodeId() *string {
	return s.NodeId
}

func (s *ConvertFlowDSLDataPluginErrors) GetProviderName() *string {
	return s.ProviderName
}

func (s *ConvertFlowDSLDataPluginErrors) GetReason() *string {
	return s.Reason
}

func (s *ConvertFlowDSLDataPluginErrors) GetToolName() *string {
	return s.ToolName
}

func (s *ConvertFlowDSLDataPluginErrors) SetNodeId(v string) *ConvertFlowDSLDataPluginErrors {
	s.NodeId = &v
	return s
}

func (s *ConvertFlowDSLDataPluginErrors) SetProviderName(v string) *ConvertFlowDSLDataPluginErrors {
	s.ProviderName = &v
	return s
}

func (s *ConvertFlowDSLDataPluginErrors) SetReason(v string) *ConvertFlowDSLDataPluginErrors {
	s.Reason = &v
	return s
}

func (s *ConvertFlowDSLDataPluginErrors) SetToolName(v string) *ConvertFlowDSLDataPluginErrors {
	s.ToolName = &v
	return s
}

func (s *ConvertFlowDSLDataPluginErrors) Validate() error {
	return dara.Validate(s)
}

type ConvertFlowDSLDataToolsetInstallations struct {
	// Toolset描述
	//
	// example:
	//
	// Google Search Plugin
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Toolset名称
	//
	// This parameter is required.
	//
	// example:
	//
	// google
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Toolset规格配置（JSON对象）
	//
	// This parameter is required.
	Spec map[string]interface{} `json:"spec" xml:"spec"`
}

func (s ConvertFlowDSLDataToolsetInstallations) String() string {
	return dara.Prettify(s)
}

func (s ConvertFlowDSLDataToolsetInstallations) GoString() string {
	return s.String()
}

func (s *ConvertFlowDSLDataToolsetInstallations) GetDescription() *string {
	return s.Description
}

func (s *ConvertFlowDSLDataToolsetInstallations) GetName() *string {
	return s.Name
}

func (s *ConvertFlowDSLDataToolsetInstallations) GetSpec() map[string]interface{} {
	return s.Spec
}

func (s *ConvertFlowDSLDataToolsetInstallations) SetDescription(v string) *ConvertFlowDSLDataToolsetInstallations {
	s.Description = &v
	return s
}

func (s *ConvertFlowDSLDataToolsetInstallations) SetName(v string) *ConvertFlowDSLDataToolsetInstallations {
	s.Name = &v
	return s
}

func (s *ConvertFlowDSLDataToolsetInstallations) SetSpec(v map[string]interface{}) *ConvertFlowDSLDataToolsetInstallations {
	s.Spec = v
	return s
}

func (s *ConvertFlowDSLDataToolsetInstallations) Validate() error {
	return dara.Validate(s)
}
