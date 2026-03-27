// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddonsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddons(v []*ListAddonsResponseBodyAddons) *ListAddonsResponseBody
	GetAddons() []*ListAddonsResponseBodyAddons
	SetRequestId(v string) *ListAddonsResponseBody
	GetRequestId() *string
}

type ListAddonsResponseBody struct {
	Addons []*ListAddonsResponseBodyAddons `json:"addons,omitempty" xml:"addons,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAddonsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBody) GetAddons() []*ListAddonsResponseBodyAddons {
	return s.Addons
}

func (s *ListAddonsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAddonsResponseBody) SetAddons(v []*ListAddonsResponseBodyAddons) *ListAddonsResponseBody {
	s.Addons = v
	return s
}

func (s *ListAddonsResponseBody) SetRequestId(v string) *ListAddonsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAddonsResponseBody) Validate() error {
	if s.Addons != nil {
		for _, item := range s.Addons {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAddonsResponseBodyAddons struct {
	// example:
	//
	// resume_vector_alias
	Alias      *string                                   `json:"alias,omitempty" xml:"alias,omitempty"`
	Categories []*string                                 `json:"categories,omitempty" xml:"categories,omitempty" type:"Repeated"`
	Dashboards []*ListAddonsResponseBodyAddonsDashboards `json:"dashboards,omitempty" xml:"dashboards,omitempty" type:"Repeated"`
	// example:
	//
	// workspace api monitor test
	Description  *string                                     `json:"description,omitempty" xml:"description,omitempty"`
	Environments []*ListAddonsResponseBodyAddonsEnvironments `json:"environments,omitempty" xml:"environments,omitempty" type:"Repeated"`
	// example:
	//
	// http://xxxxxxx
	Icon     *string   `json:"icon,omitempty" xml:"icon,omitempty"`
	Keywords []*string `json:"keywords,omitempty" xml:"keywords,omitempty" type:"Repeated"`
	// example:
	//
	// zh
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// 2025-09-03T03:15:56Z
	LatestReleaseCreateTime *string `json:"latestReleaseCreateTime,omitempty" xml:"latestReleaseCreateTime,omitempty"`
	// example:
	//
	// rum_api_dashboard_explorer_link_metric_set
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// true
	Once *bool `json:"once,omitempty" xml:"once,omitempty"`
	// example:
	//
	// middleware
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// *
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// 99
	Weight *string `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s ListAddonsResponseBodyAddons) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBodyAddons) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBodyAddons) GetAlias() *string {
	return s.Alias
}

func (s *ListAddonsResponseBodyAddons) GetCategories() []*string {
	return s.Categories
}

func (s *ListAddonsResponseBodyAddons) GetDashboards() []*ListAddonsResponseBodyAddonsDashboards {
	return s.Dashboards
}

func (s *ListAddonsResponseBodyAddons) GetDescription() *string {
	return s.Description
}

func (s *ListAddonsResponseBodyAddons) GetEnvironments() []*ListAddonsResponseBodyAddonsEnvironments {
	return s.Environments
}

func (s *ListAddonsResponseBodyAddons) GetIcon() *string {
	return s.Icon
}

func (s *ListAddonsResponseBodyAddons) GetKeywords() []*string {
	return s.Keywords
}

func (s *ListAddonsResponseBodyAddons) GetLanguage() *string {
	return s.Language
}

func (s *ListAddonsResponseBodyAddons) GetLatestReleaseCreateTime() *string {
	return s.LatestReleaseCreateTime
}

func (s *ListAddonsResponseBodyAddons) GetName() *string {
	return s.Name
}

func (s *ListAddonsResponseBodyAddons) GetOnce() *bool {
	return s.Once
}

func (s *ListAddonsResponseBodyAddons) GetScene() *string {
	return s.Scene
}

func (s *ListAddonsResponseBodyAddons) GetVersion() *string {
	return s.Version
}

func (s *ListAddonsResponseBodyAddons) GetWeight() *string {
	return s.Weight
}

func (s *ListAddonsResponseBodyAddons) SetAlias(v string) *ListAddonsResponseBodyAddons {
	s.Alias = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetCategories(v []*string) *ListAddonsResponseBodyAddons {
	s.Categories = v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetDashboards(v []*ListAddonsResponseBodyAddonsDashboards) *ListAddonsResponseBodyAddons {
	s.Dashboards = v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetDescription(v string) *ListAddonsResponseBodyAddons {
	s.Description = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetEnvironments(v []*ListAddonsResponseBodyAddonsEnvironments) *ListAddonsResponseBodyAddons {
	s.Environments = v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetIcon(v string) *ListAddonsResponseBodyAddons {
	s.Icon = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetKeywords(v []*string) *ListAddonsResponseBodyAddons {
	s.Keywords = v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetLanguage(v string) *ListAddonsResponseBodyAddons {
	s.Language = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetLatestReleaseCreateTime(v string) *ListAddonsResponseBodyAddons {
	s.LatestReleaseCreateTime = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetName(v string) *ListAddonsResponseBodyAddons {
	s.Name = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetOnce(v bool) *ListAddonsResponseBodyAddons {
	s.Once = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetScene(v string) *ListAddonsResponseBodyAddons {
	s.Scene = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetVersion(v string) *ListAddonsResponseBodyAddons {
	s.Version = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetWeight(v string) *ListAddonsResponseBodyAddons {
	s.Weight = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) Validate() error {
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

type ListAddonsResponseBodyAddonsDashboards struct {
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// rum_view_link_rum_api
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// http://xxxxxxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListAddonsResponseBodyAddonsDashboards) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBodyAddonsDashboards) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBodyAddonsDashboards) GetDescription() *string {
	return s.Description
}

func (s *ListAddonsResponseBodyAddonsDashboards) GetName() *string {
	return s.Name
}

func (s *ListAddonsResponseBodyAddonsDashboards) GetUrl() *string {
	return s.Url
}

func (s *ListAddonsResponseBodyAddonsDashboards) SetDescription(v string) *ListAddonsResponseBodyAddonsDashboards {
	s.Description = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsDashboards) SetName(v string) *ListAddonsResponseBodyAddonsDashboards {
	s.Name = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsDashboards) SetUrl(v string) *ListAddonsResponseBodyAddonsDashboards {
	s.Url = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsDashboards) Validate() error {
	return dara.Validate(s)
}

type ListAddonsResponseBodyAddonsEnvironments struct {
	CommonSchemaRefs []*ListAddonsResponseBodyAddonsEnvironmentsCommonSchemaRefs `json:"commonSchemaRefs,omitempty" xml:"commonSchemaRefs,omitempty" type:"Repeated"`
	Dependencies     *ListAddonsResponseBodyAddonsEnvironmentsDependencies       `json:"dependencies,omitempty" xml:"dependencies,omitempty" type:"Struct"`
	// example:
	//
	// o11y-demo-cn-heyuan
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// false
	Enable *bool   `json:"enable,omitempty" xml:"enable,omitempty"`
	Label  *string `json:"label,omitempty" xml:"label,omitempty"`
	// example:
	//
	// CS
	Name     *string                                           `json:"name,omitempty" xml:"name,omitempty"`
	Policies *ListAddonsResponseBodyAddonsEnvironmentsPolicies `json:"policies,omitempty" xml:"policies,omitempty" type:"Struct"`
	// example:
	//
	// CS
	PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty"`
}

func (s ListAddonsResponseBodyAddonsEnvironments) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBodyAddonsEnvironments) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBodyAddonsEnvironments) GetCommonSchemaRefs() []*ListAddonsResponseBodyAddonsEnvironmentsCommonSchemaRefs {
	return s.CommonSchemaRefs
}

func (s *ListAddonsResponseBodyAddonsEnvironments) GetDependencies() *ListAddonsResponseBodyAddonsEnvironmentsDependencies {
	return s.Dependencies
}

func (s *ListAddonsResponseBodyAddonsEnvironments) GetDescription() *string {
	return s.Description
}

func (s *ListAddonsResponseBodyAddonsEnvironments) GetEnable() *bool {
	return s.Enable
}

func (s *ListAddonsResponseBodyAddonsEnvironments) GetLabel() *string {
	return s.Label
}

func (s *ListAddonsResponseBodyAddonsEnvironments) GetName() *string {
	return s.Name
}

func (s *ListAddonsResponseBodyAddonsEnvironments) GetPolicies() *ListAddonsResponseBodyAddonsEnvironmentsPolicies {
	return s.Policies
}

func (s *ListAddonsResponseBodyAddonsEnvironments) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListAddonsResponseBodyAddonsEnvironments) SetCommonSchemaRefs(v []*ListAddonsResponseBodyAddonsEnvironmentsCommonSchemaRefs) *ListAddonsResponseBodyAddonsEnvironments {
	s.CommonSchemaRefs = v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironments) SetDependencies(v *ListAddonsResponseBodyAddonsEnvironmentsDependencies) *ListAddonsResponseBodyAddonsEnvironments {
	s.Dependencies = v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironments) SetDescription(v string) *ListAddonsResponseBodyAddonsEnvironments {
	s.Description = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironments) SetEnable(v bool) *ListAddonsResponseBodyAddonsEnvironments {
	s.Enable = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironments) SetLabel(v string) *ListAddonsResponseBodyAddonsEnvironments {
	s.Label = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironments) SetName(v string) *ListAddonsResponseBodyAddonsEnvironments {
	s.Name = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironments) SetPolicies(v *ListAddonsResponseBodyAddonsEnvironmentsPolicies) *ListAddonsResponseBodyAddonsEnvironments {
	s.Policies = v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironments) SetPolicyType(v string) *ListAddonsResponseBodyAddonsEnvironments {
	s.PolicyType = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironments) Validate() error {
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

type ListAddonsResponseBodyAddonsEnvironmentsCommonSchemaRefs struct {
	// example:
	//
	// acs-ecs
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// example:
	//
	// 0.1.4
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListAddonsResponseBodyAddonsEnvironmentsCommonSchemaRefs) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBodyAddonsEnvironmentsCommonSchemaRefs) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsCommonSchemaRefs) GetGroup() *string {
	return s.Group
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsCommonSchemaRefs) GetVersion() *string {
	return s.Version
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsCommonSchemaRefs) SetGroup(v string) *ListAddonsResponseBodyAddonsEnvironmentsCommonSchemaRefs {
	s.Group = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsCommonSchemaRefs) SetVersion(v string) *ListAddonsResponseBodyAddonsEnvironmentsCommonSchemaRefs {
	s.Version = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsCommonSchemaRefs) Validate() error {
	return dara.Validate(s)
}

type ListAddonsResponseBodyAddonsEnvironmentsDependencies struct {
	ClusterTypes []*string        `json:"clusterTypes,omitempty" xml:"clusterTypes,omitempty" type:"Repeated"`
	Features     map[string]*bool `json:"features,omitempty" xml:"features,omitempty"`
	Services     []*string        `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s ListAddonsResponseBodyAddonsEnvironmentsDependencies) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBodyAddonsEnvironmentsDependencies) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsDependencies) GetClusterTypes() []*string {
	return s.ClusterTypes
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsDependencies) GetFeatures() map[string]*bool {
	return s.Features
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsDependencies) GetServices() []*string {
	return s.Services
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsDependencies) SetClusterTypes(v []*string) *ListAddonsResponseBodyAddonsEnvironmentsDependencies {
	s.ClusterTypes = v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsDependencies) SetFeatures(v map[string]*bool) *ListAddonsResponseBodyAddonsEnvironmentsDependencies {
	s.Features = v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsDependencies) SetServices(v []*string) *ListAddonsResponseBodyAddonsEnvironmentsDependencies {
	s.Services = v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsDependencies) Validate() error {
	return dara.Validate(s)
}

type ListAddonsResponseBodyAddonsEnvironmentsPolicies struct {
	// example:
	//
	// RUNNING
	AlertDefaultStatus *string `json:"alertDefaultStatus,omitempty" xml:"alertDefaultStatus,omitempty"`
	// example:
	//
	// true
	DefaultInstall *bool `json:"defaultInstall,omitempty" xml:"defaultInstall,omitempty"`
	// example:
	//
	// true
	EnableServiceAccount *bool                                                            `json:"enableServiceAccount,omitempty" xml:"enableServiceAccount,omitempty"`
	MetricCheckRule      *ListAddonsResponseBodyAddonsEnvironmentsPoliciesMetricCheckRule `json:"metricCheckRule,omitempty" xml:"metricCheckRule,omitempty" type:"Struct"`
	// example:
	//
	// true
	NeedRestartAfterIntegration *bool                                                        `json:"needRestartAfterIntegration,omitempty" xml:"needRestartAfterIntegration,omitempty"`
	Protocols                   []*ListAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	// example:
	//
	// cloud-acs-ecs
	TargetAddonName *string `json:"targetAddonName,omitempty" xml:"targetAddonName,omitempty"`
}

func (s ListAddonsResponseBodyAddonsEnvironmentsPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBodyAddonsEnvironmentsPolicies) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPolicies) GetAlertDefaultStatus() *string {
	return s.AlertDefaultStatus
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPolicies) GetDefaultInstall() *bool {
	return s.DefaultInstall
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPolicies) GetEnableServiceAccount() *bool {
	return s.EnableServiceAccount
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPolicies) GetMetricCheckRule() *ListAddonsResponseBodyAddonsEnvironmentsPoliciesMetricCheckRule {
	return s.MetricCheckRule
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPolicies) GetNeedRestartAfterIntegration() *bool {
	return s.NeedRestartAfterIntegration
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPolicies) GetProtocols() []*ListAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols {
	return s.Protocols
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPolicies) GetTargetAddonName() *string {
	return s.TargetAddonName
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPolicies) SetAlertDefaultStatus(v string) *ListAddonsResponseBodyAddonsEnvironmentsPolicies {
	s.AlertDefaultStatus = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPolicies) SetDefaultInstall(v bool) *ListAddonsResponseBodyAddonsEnvironmentsPolicies {
	s.DefaultInstall = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPolicies) SetEnableServiceAccount(v bool) *ListAddonsResponseBodyAddonsEnvironmentsPolicies {
	s.EnableServiceAccount = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPolicies) SetMetricCheckRule(v *ListAddonsResponseBodyAddonsEnvironmentsPoliciesMetricCheckRule) *ListAddonsResponseBodyAddonsEnvironmentsPolicies {
	s.MetricCheckRule = v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPolicies) SetNeedRestartAfterIntegration(v bool) *ListAddonsResponseBodyAddonsEnvironmentsPolicies {
	s.NeedRestartAfterIntegration = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPolicies) SetProtocols(v []*ListAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) *ListAddonsResponseBodyAddonsEnvironmentsPolicies {
	s.Protocols = v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPolicies) SetTargetAddonName(v string) *ListAddonsResponseBodyAddonsEnvironmentsPolicies {
	s.TargetAddonName = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPolicies) Validate() error {
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

type ListAddonsResponseBodyAddonsEnvironmentsPoliciesMetricCheckRule struct {
	PromQL []*string `json:"promQL,omitempty" xml:"promQL,omitempty" type:"Repeated"`
}

func (s ListAddonsResponseBodyAddonsEnvironmentsPoliciesMetricCheckRule) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBodyAddonsEnvironmentsPoliciesMetricCheckRule) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPoliciesMetricCheckRule) GetPromQL() []*string {
	return s.PromQL
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPoliciesMetricCheckRule) SetPromQL(v []*string) *ListAddonsResponseBodyAddonsEnvironmentsPoliciesMetricCheckRule {
	s.PromQL = v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPoliciesMetricCheckRule) Validate() error {
	return dara.Validate(s)
}

type ListAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols struct {
	// example:
	//
	// Support OpenTelemetry Protocal
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// http://xxxxxxx
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// OpenTelemetry
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// example:
	//
	// openTelemetry
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) GetDescription() *string {
	return s.Description
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) GetIcon() *string {
	return s.Icon
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) GetLabel() *string {
	return s.Label
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) GetName() *string {
	return s.Name
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) SetDescription(v string) *ListAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols {
	s.Description = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) SetIcon(v string) *ListAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols {
	s.Icon = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) SetLabel(v string) *ListAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols {
	s.Label = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) SetName(v string) *ListAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols {
	s.Name = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) Validate() error {
	return dara.Validate(s)
}
