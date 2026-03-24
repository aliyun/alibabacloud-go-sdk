// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddonMeta interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *AddonMeta
	GetAlias() *string
	SetCategories(v []*string) *AddonMeta
	GetCategories() []*string
	SetDashboards(v []*AddonMetaDashboards) *AddonMeta
	GetDashboards() []*AddonMetaDashboards
	SetDescription(v string) *AddonMeta
	GetDescription() *string
	SetEnvironments(v []*AddonMetaEnvironments) *AddonMeta
	GetEnvironments() []*AddonMetaEnvironments
	SetIcon(v string) *AddonMeta
	GetIcon() *string
	SetKeywords(v []*string) *AddonMeta
	GetKeywords() []*string
	SetLanguage(v string) *AddonMeta
	GetLanguage() *string
	SetLatestReleaseCreateTime(v string) *AddonMeta
	GetLatestReleaseCreateTime() *string
	SetName(v string) *AddonMeta
	GetName() *string
	SetOnce(v bool) *AddonMeta
	GetOnce() *bool
	SetScene(v string) *AddonMeta
	GetScene() *string
	SetVersion(v string) *AddonMeta
	GetVersion() *string
	SetWeight(v int32) *AddonMeta
	GetWeight() *int32
}

type AddonMeta struct {
	// 组件别名，显示名称
	//
	// example:
	//
	// ECS 监控
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// 组件分类信息
	Categories []*string `json:"categories,omitempty" xml:"categories,omitempty" type:"Repeated"`
	// 组件示意图列表
	Dashboards []*AddonMetaDashboards `json:"dashboards,omitempty" xml:"dashboards,omitempty" type:"Repeated"`
	// 描述信息。
	//
	// example:
	//
	// The out-of-the-box and comprehensive ECS observe dashboards and alarm rules. Based on AliYun CloudMonitor agentless metrics, exporter agent metrics, host audit logs, host events and other data.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 支持的环境类型列表
	Environments []*AddonMetaEnvironments `json:"environments,omitempty" xml:"environments,omitempty" type:"Repeated"`
	// 组件图标。
	//
	// example:
	//
	// assets/logos/ecs.svg
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 关键词列表
	Keywords []*string `json:"keywords,omitempty" xml:"keywords,omitempty" type:"Repeated"`
	// 语言，取值：
	//
	// - zh：中文（默认值）
	//
	// - en：英文
	//
	// example:
	//
	// zh
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 该组件上一次接入时间
	//
	// example:
	//
	// 2025-10-25 09:12:12
	LatestReleaseCreateTime *string `json:"latestReleaseCreateTime,omitempty" xml:"latestReleaseCreateTime,omitempty"`
	// 组件名称
	//
	// example:
	//
	// cloud-acs-ecs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Policy 下是否只能安装一次
	//
	// example:
	//
	// true/false
	Once *bool `json:"once,omitempty" xml:"once,omitempty"`
	// 场景
	//
	// example:
	//
	// feature
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// 版本号
	//
	// example:
	//
	// 0.0.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// 组件排序权重
	//
	// example:
	//
	// 1000
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s AddonMeta) String() string {
	return dara.Prettify(s)
}

func (s AddonMeta) GoString() string {
	return s.String()
}

func (s *AddonMeta) GetAlias() *string {
	return s.Alias
}

func (s *AddonMeta) GetCategories() []*string {
	return s.Categories
}

func (s *AddonMeta) GetDashboards() []*AddonMetaDashboards {
	return s.Dashboards
}

func (s *AddonMeta) GetDescription() *string {
	return s.Description
}

func (s *AddonMeta) GetEnvironments() []*AddonMetaEnvironments {
	return s.Environments
}

func (s *AddonMeta) GetIcon() *string {
	return s.Icon
}

func (s *AddonMeta) GetKeywords() []*string {
	return s.Keywords
}

func (s *AddonMeta) GetLanguage() *string {
	return s.Language
}

func (s *AddonMeta) GetLatestReleaseCreateTime() *string {
	return s.LatestReleaseCreateTime
}

func (s *AddonMeta) GetName() *string {
	return s.Name
}

func (s *AddonMeta) GetOnce() *bool {
	return s.Once
}

func (s *AddonMeta) GetScene() *string {
	return s.Scene
}

func (s *AddonMeta) GetVersion() *string {
	return s.Version
}

func (s *AddonMeta) GetWeight() *int32 {
	return s.Weight
}

func (s *AddonMeta) SetAlias(v string) *AddonMeta {
	s.Alias = &v
	return s
}

func (s *AddonMeta) SetCategories(v []*string) *AddonMeta {
	s.Categories = v
	return s
}

func (s *AddonMeta) SetDashboards(v []*AddonMetaDashboards) *AddonMeta {
	s.Dashboards = v
	return s
}

func (s *AddonMeta) SetDescription(v string) *AddonMeta {
	s.Description = &v
	return s
}

func (s *AddonMeta) SetEnvironments(v []*AddonMetaEnvironments) *AddonMeta {
	s.Environments = v
	return s
}

func (s *AddonMeta) SetIcon(v string) *AddonMeta {
	s.Icon = &v
	return s
}

func (s *AddonMeta) SetKeywords(v []*string) *AddonMeta {
	s.Keywords = v
	return s
}

func (s *AddonMeta) SetLanguage(v string) *AddonMeta {
	s.Language = &v
	return s
}

func (s *AddonMeta) SetLatestReleaseCreateTime(v string) *AddonMeta {
	s.LatestReleaseCreateTime = &v
	return s
}

func (s *AddonMeta) SetName(v string) *AddonMeta {
	s.Name = &v
	return s
}

func (s *AddonMeta) SetOnce(v bool) *AddonMeta {
	s.Once = &v
	return s
}

func (s *AddonMeta) SetScene(v string) *AddonMeta {
	s.Scene = &v
	return s
}

func (s *AddonMeta) SetVersion(v string) *AddonMeta {
	s.Version = &v
	return s
}

func (s *AddonMeta) SetWeight(v int32) *AddonMeta {
	s.Weight = &v
	return s
}

func (s *AddonMeta) Validate() error {
	if s.Dashboards != nil {
		for _, item := range s.Dashboards {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Environments != nil {
		for _, item := range s.Environments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddonMetaDashboards struct {
	// 描述信息
	//
	// example:
	//
	// 描述信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 示意图名称
	//
	// example:
	//
	// ECS 监控概览大盘
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 示意图 URL
	//
	// example:
	//
	// assets/dashboards/ecs.png
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s AddonMetaDashboards) String() string {
	return dara.Prettify(s)
}

func (s AddonMetaDashboards) GoString() string {
	return s.String()
}

func (s *AddonMetaDashboards) GetDescription() *string {
	return s.Description
}

func (s *AddonMetaDashboards) GetName() *string {
	return s.Name
}

func (s *AddonMetaDashboards) GetUrl() *string {
	return s.Url
}

func (s *AddonMetaDashboards) SetDescription(v string) *AddonMetaDashboards {
	s.Description = &v
	return s
}

func (s *AddonMetaDashboards) SetName(v string) *AddonMetaDashboards {
	s.Name = &v
	return s
}

func (s *AddonMetaDashboards) SetUrl(v string) *AddonMetaDashboards {
	s.Url = &v
	return s
}

func (s *AddonMetaDashboards) Validate() error {
	return dara.Validate(s)
}

type AddonMetaEnvironments struct {
	// 绑定的CommonSchema 列表
	CommonSchemaRefs []*AddonMetaEnvironmentsCommonSchemaRefs `json:"commonSchemaRefs,omitempty" xml:"commonSchemaRefs,omitempty" type:"Repeated"`
	// 依赖描述信息
	Dependencies *AddonMetaEnvironmentsDependencies `json:"dependencies,omitempty" xml:"dependencies,omitempty" type:"Struct"`
	// 环境类型的描述
	//
	// example:
	//
	// 支持容器集群的工作覆盖监控
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 是否启用
	//
	// example:
	//
	// true/false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// 环境类型显示名称
	//
	// example:
	//
	// 容器环境
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// 环境类型名称
	//
	// example:
	//
	// CS/ECS/Cloud/Client
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 组件的控制策略组合信息
	Policies *AddonMetaEnvironmentsPolicies `json:"policies,omitempty" xml:"policies,omitempty" type:"Struct"`
	// 策略类型
	//
	// example:
	//
	// ECS
	PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty"`
}

func (s AddonMetaEnvironments) String() string {
	return dara.Prettify(s)
}

func (s AddonMetaEnvironments) GoString() string {
	return s.String()
}

func (s *AddonMetaEnvironments) GetCommonSchemaRefs() []*AddonMetaEnvironmentsCommonSchemaRefs {
	return s.CommonSchemaRefs
}

func (s *AddonMetaEnvironments) GetDependencies() *AddonMetaEnvironmentsDependencies {
	return s.Dependencies
}

func (s *AddonMetaEnvironments) GetDescription() *string {
	return s.Description
}

func (s *AddonMetaEnvironments) GetEnable() *bool {
	return s.Enable
}

func (s *AddonMetaEnvironments) GetLabel() *string {
	return s.Label
}

func (s *AddonMetaEnvironments) GetName() *string {
	return s.Name
}

func (s *AddonMetaEnvironments) GetPolicies() *AddonMetaEnvironmentsPolicies {
	return s.Policies
}

func (s *AddonMetaEnvironments) GetPolicyType() *string {
	return s.PolicyType
}

func (s *AddonMetaEnvironments) SetCommonSchemaRefs(v []*AddonMetaEnvironmentsCommonSchemaRefs) *AddonMetaEnvironments {
	s.CommonSchemaRefs = v
	return s
}

func (s *AddonMetaEnvironments) SetDependencies(v *AddonMetaEnvironmentsDependencies) *AddonMetaEnvironments {
	s.Dependencies = v
	return s
}

func (s *AddonMetaEnvironments) SetDescription(v string) *AddonMetaEnvironments {
	s.Description = &v
	return s
}

func (s *AddonMetaEnvironments) SetEnable(v bool) *AddonMetaEnvironments {
	s.Enable = &v
	return s
}

func (s *AddonMetaEnvironments) SetLabel(v string) *AddonMetaEnvironments {
	s.Label = &v
	return s
}

func (s *AddonMetaEnvironments) SetName(v string) *AddonMetaEnvironments {
	s.Name = &v
	return s
}

func (s *AddonMetaEnvironments) SetPolicies(v *AddonMetaEnvironmentsPolicies) *AddonMetaEnvironments {
	s.Policies = v
	return s
}

func (s *AddonMetaEnvironments) SetPolicyType(v string) *AddonMetaEnvironments {
	s.PolicyType = &v
	return s
}

func (s *AddonMetaEnvironments) Validate() error {
	if s.CommonSchemaRefs != nil {
		for _, item := range s.CommonSchemaRefs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Dependencies != nil {
		if err := s.Dependencies.Validate(); err != nil {
			return err
		}
	}
	if s.Policies != nil {
		if err := s.Policies.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddonMetaEnvironmentsCommonSchemaRefs struct {
	// CommonSchema 的分组名称
	//
	// example:
	//
	// acs-ecs
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// CommonSchema 的分组版本
	//
	// example:
	//
	// 0.1.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s AddonMetaEnvironmentsCommonSchemaRefs) String() string {
	return dara.Prettify(s)
}

func (s AddonMetaEnvironmentsCommonSchemaRefs) GoString() string {
	return s.String()
}

func (s *AddonMetaEnvironmentsCommonSchemaRefs) GetGroup() *string {
	return s.Group
}

func (s *AddonMetaEnvironmentsCommonSchemaRefs) GetVersion() *string {
	return s.Version
}

func (s *AddonMetaEnvironmentsCommonSchemaRefs) SetGroup(v string) *AddonMetaEnvironmentsCommonSchemaRefs {
	s.Group = &v
	return s
}

func (s *AddonMetaEnvironmentsCommonSchemaRefs) SetVersion(v string) *AddonMetaEnvironmentsCommonSchemaRefs {
	s.Version = &v
	return s
}

func (s *AddonMetaEnvironmentsCommonSchemaRefs) Validate() error {
	return dara.Validate(s)
}

type AddonMetaEnvironmentsDependencies struct {
	// 支持的集群类型
	ClusterTypes []*string `json:"clusterTypes,omitempty" xml:"clusterTypes,omitempty" type:"Repeated"`
	// 探针依赖描述，组件名称。新版已由 collectors 字段替换
	Features map[string]*bool `json:"features,omitempty" xml:"features,omitempty"`
	// 依赖的服务列表
	Services []*string `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s AddonMetaEnvironmentsDependencies) String() string {
	return dara.Prettify(s)
}

func (s AddonMetaEnvironmentsDependencies) GoString() string {
	return s.String()
}

func (s *AddonMetaEnvironmentsDependencies) GetClusterTypes() []*string {
	return s.ClusterTypes
}

func (s *AddonMetaEnvironmentsDependencies) GetFeatures() map[string]*bool {
	return s.Features
}

func (s *AddonMetaEnvironmentsDependencies) GetServices() []*string {
	return s.Services
}

func (s *AddonMetaEnvironmentsDependencies) SetClusterTypes(v []*string) *AddonMetaEnvironmentsDependencies {
	s.ClusterTypes = v
	return s
}

func (s *AddonMetaEnvironmentsDependencies) SetFeatures(v map[string]*bool) *AddonMetaEnvironmentsDependencies {
	s.Features = v
	return s
}

func (s *AddonMetaEnvironmentsDependencies) SetServices(v []*string) *AddonMetaEnvironmentsDependencies {
	s.Services = v
	return s
}

func (s *AddonMetaEnvironmentsDependencies) Validate() error {
	return dara.Validate(s)
}

type AddonMetaEnvironmentsPolicies struct {
	// 告警规则默认安装后是否启用
	//
	// example:
	//
	// RUNNING
	AlertDefaultStatus *string `json:"alertDefaultStatus,omitempty" xml:"alertDefaultStatus,omitempty"`
	// 默认模式，即无需绑定实体的接入模式。
	//
	// example:
	//
	// true/false
	BindDefaultPolicy *bool `json:"bindDefaultPolicy,omitempty" xml:"bindDefaultPolicy,omitempty"`
	// 绑定的目标实体信息
	BindEntity *AddonMetaEnvironmentsPoliciesBindEntity `json:"bindEntity,omitempty" xml:"bindEntity,omitempty" type:"Struct"`
	// 是否默认安装
	//
	// example:
	//
	// true/false
	DefaultInstall *bool `json:"defaultInstall,omitempty" xml:"defaultInstall,omitempty"`
	// 是否启用内部授权Token分配
	//
	// example:
	//
	// true/false
	EnableServiceAccount *bool `json:"enableServiceAccount,omitempty" xml:"enableServiceAccount,omitempty"`
	// 组件接入后的数据检查规则
	MetricCheckRule *AddonMetaEnvironmentsPoliciesMetricCheckRule `json:"metricCheckRule,omitempty" xml:"metricCheckRule,omitempty" type:"Struct"`
	// 是否需要在接入后提示重启工作负载
	//
	// example:
	//
	// true/false
	NeedRestartAfterIntegration *bool `json:"needRestartAfterIntegration,omitempty" xml:"needRestartAfterIntegration,omitempty"`
	// 支持的客户端协议信息列表
	Protocols []*AddonMetaEnvironmentsPoliciesProtocols `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	// 跳转的目标组件名称
	//
	// example:
	//
	// cloud-acs-ecs
	TargetAddonName *string `json:"targetAddonName,omitempty" xml:"targetAddonName,omitempty"`
}

func (s AddonMetaEnvironmentsPolicies) String() string {
	return dara.Prettify(s)
}

func (s AddonMetaEnvironmentsPolicies) GoString() string {
	return s.String()
}

func (s *AddonMetaEnvironmentsPolicies) GetAlertDefaultStatus() *string {
	return s.AlertDefaultStatus
}

func (s *AddonMetaEnvironmentsPolicies) GetBindDefaultPolicy() *bool {
	return s.BindDefaultPolicy
}

func (s *AddonMetaEnvironmentsPolicies) GetBindEntity() *AddonMetaEnvironmentsPoliciesBindEntity {
	return s.BindEntity
}

func (s *AddonMetaEnvironmentsPolicies) GetDefaultInstall() *bool {
	return s.DefaultInstall
}

func (s *AddonMetaEnvironmentsPolicies) GetEnableServiceAccount() *bool {
	return s.EnableServiceAccount
}

func (s *AddonMetaEnvironmentsPolicies) GetMetricCheckRule() *AddonMetaEnvironmentsPoliciesMetricCheckRule {
	return s.MetricCheckRule
}

func (s *AddonMetaEnvironmentsPolicies) GetNeedRestartAfterIntegration() *bool {
	return s.NeedRestartAfterIntegration
}

func (s *AddonMetaEnvironmentsPolicies) GetProtocols() []*AddonMetaEnvironmentsPoliciesProtocols {
	return s.Protocols
}

func (s *AddonMetaEnvironmentsPolicies) GetTargetAddonName() *string {
	return s.TargetAddonName
}

func (s *AddonMetaEnvironmentsPolicies) SetAlertDefaultStatus(v string) *AddonMetaEnvironmentsPolicies {
	s.AlertDefaultStatus = &v
	return s
}

func (s *AddonMetaEnvironmentsPolicies) SetBindDefaultPolicy(v bool) *AddonMetaEnvironmentsPolicies {
	s.BindDefaultPolicy = &v
	return s
}

func (s *AddonMetaEnvironmentsPolicies) SetBindEntity(v *AddonMetaEnvironmentsPoliciesBindEntity) *AddonMetaEnvironmentsPolicies {
	s.BindEntity = v
	return s
}

func (s *AddonMetaEnvironmentsPolicies) SetDefaultInstall(v bool) *AddonMetaEnvironmentsPolicies {
	s.DefaultInstall = &v
	return s
}

func (s *AddonMetaEnvironmentsPolicies) SetEnableServiceAccount(v bool) *AddonMetaEnvironmentsPolicies {
	s.EnableServiceAccount = &v
	return s
}

func (s *AddonMetaEnvironmentsPolicies) SetMetricCheckRule(v *AddonMetaEnvironmentsPoliciesMetricCheckRule) *AddonMetaEnvironmentsPolicies {
	s.MetricCheckRule = v
	return s
}

func (s *AddonMetaEnvironmentsPolicies) SetNeedRestartAfterIntegration(v bool) *AddonMetaEnvironmentsPolicies {
	s.NeedRestartAfterIntegration = &v
	return s
}

func (s *AddonMetaEnvironmentsPolicies) SetProtocols(v []*AddonMetaEnvironmentsPoliciesProtocols) *AddonMetaEnvironmentsPolicies {
	s.Protocols = v
	return s
}

func (s *AddonMetaEnvironmentsPolicies) SetTargetAddonName(v string) *AddonMetaEnvironmentsPolicies {
	s.TargetAddonName = &v
	return s
}

func (s *AddonMetaEnvironmentsPolicies) Validate() error {
	if s.BindEntity != nil {
		if err := s.BindEntity.Validate(); err != nil {
			return err
		}
	}
	if s.MetricCheckRule != nil {
		if err := s.MetricCheckRule.Validate(); err != nil {
			return err
		}
	}
	if s.Protocols != nil {
		for _, item := range s.Protocols {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddonMetaEnvironmentsPoliciesBindEntity struct {
	// 是否是组模式
	//
	// example:
	//
	// true/false
	EntityGroupMode *bool `json:"entityGroupMode,omitempty" xml:"entityGroupMode,omitempty"`
	// 实体类型
	//
	// example:
	//
	// acs.ecs.instance
	EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
	// 是否是单实体模式
	//
	// example:
	//
	// true/false
	SingleEntityMode *bool `json:"singleEntityMode,omitempty" xml:"singleEntityMode,omitempty"`
	// 实体中提取VPC ID 信息的字段
	//
	// example:
	//
	// vpcId
	VpcIdFieldKey *string `json:"vpcIdFieldKey,omitempty" xml:"vpcIdFieldKey,omitempty"`
}

func (s AddonMetaEnvironmentsPoliciesBindEntity) String() string {
	return dara.Prettify(s)
}

func (s AddonMetaEnvironmentsPoliciesBindEntity) GoString() string {
	return s.String()
}

func (s *AddonMetaEnvironmentsPoliciesBindEntity) GetEntityGroupMode() *bool {
	return s.EntityGroupMode
}

func (s *AddonMetaEnvironmentsPoliciesBindEntity) GetEntityType() *string {
	return s.EntityType
}

func (s *AddonMetaEnvironmentsPoliciesBindEntity) GetSingleEntityMode() *bool {
	return s.SingleEntityMode
}

func (s *AddonMetaEnvironmentsPoliciesBindEntity) GetVpcIdFieldKey() *string {
	return s.VpcIdFieldKey
}

func (s *AddonMetaEnvironmentsPoliciesBindEntity) SetEntityGroupMode(v bool) *AddonMetaEnvironmentsPoliciesBindEntity {
	s.EntityGroupMode = &v
	return s
}

func (s *AddonMetaEnvironmentsPoliciesBindEntity) SetEntityType(v string) *AddonMetaEnvironmentsPoliciesBindEntity {
	s.EntityType = &v
	return s
}

func (s *AddonMetaEnvironmentsPoliciesBindEntity) SetSingleEntityMode(v bool) *AddonMetaEnvironmentsPoliciesBindEntity {
	s.SingleEntityMode = &v
	return s
}

func (s *AddonMetaEnvironmentsPoliciesBindEntity) SetVpcIdFieldKey(v string) *AddonMetaEnvironmentsPoliciesBindEntity {
	s.VpcIdFieldKey = &v
	return s
}

func (s *AddonMetaEnvironmentsPoliciesBindEntity) Validate() error {
	return dara.Validate(s)
}

type AddonMetaEnvironmentsPoliciesMetricCheckRule struct {
	// 检测规则 PromQL
	PromQL []*string `json:"promQL,omitempty" xml:"promQL,omitempty" type:"Repeated"`
}

func (s AddonMetaEnvironmentsPoliciesMetricCheckRule) String() string {
	return dara.Prettify(s)
}

func (s AddonMetaEnvironmentsPoliciesMetricCheckRule) GoString() string {
	return s.String()
}

func (s *AddonMetaEnvironmentsPoliciesMetricCheckRule) GetPromQL() []*string {
	return s.PromQL
}

func (s *AddonMetaEnvironmentsPoliciesMetricCheckRule) SetPromQL(v []*string) *AddonMetaEnvironmentsPoliciesMetricCheckRule {
	s.PromQL = v
	return s
}

func (s *AddonMetaEnvironmentsPoliciesMetricCheckRule) Validate() error {
	return dara.Validate(s)
}

type AddonMetaEnvironmentsPoliciesProtocols struct {
	// 协议描述
	//
	// example:
	//
	// 使用 Prometheus 协议写入指标数据
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 协议显示icon
	//
	// example:
	//
	// assets/logos/ecs.svg
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 协议显示名称
	//
	// example:
	//
	// Prometheus 协议
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// 协议名称
	//
	// example:
	//
	// Prometheus
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s AddonMetaEnvironmentsPoliciesProtocols) String() string {
	return dara.Prettify(s)
}

func (s AddonMetaEnvironmentsPoliciesProtocols) GoString() string {
	return s.String()
}

func (s *AddonMetaEnvironmentsPoliciesProtocols) GetDescription() *string {
	return s.Description
}

func (s *AddonMetaEnvironmentsPoliciesProtocols) GetIcon() *string {
	return s.Icon
}

func (s *AddonMetaEnvironmentsPoliciesProtocols) GetLabel() *string {
	return s.Label
}

func (s *AddonMetaEnvironmentsPoliciesProtocols) GetName() *string {
	return s.Name
}

func (s *AddonMetaEnvironmentsPoliciesProtocols) SetDescription(v string) *AddonMetaEnvironmentsPoliciesProtocols {
	s.Description = &v
	return s
}

func (s *AddonMetaEnvironmentsPoliciesProtocols) SetIcon(v string) *AddonMetaEnvironmentsPoliciesProtocols {
	s.Icon = &v
	return s
}

func (s *AddonMetaEnvironmentsPoliciesProtocols) SetLabel(v string) *AddonMetaEnvironmentsPoliciesProtocols {
	s.Label = &v
	return s
}

func (s *AddonMetaEnvironmentsPoliciesProtocols) SetName(v string) *AddonMetaEnvironmentsPoliciesProtocols {
	s.Name = &v
	return s
}

func (s *AddonMetaEnvironmentsPoliciesProtocols) Validate() error {
	return dara.Validate(s)
}
