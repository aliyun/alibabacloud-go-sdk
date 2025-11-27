// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyAddonsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddons(v []*ListIntegrationPolicyAddonsResponseBodyAddons) *ListIntegrationPolicyAddonsResponseBody
	GetAddons() []*ListIntegrationPolicyAddonsResponseBodyAddons
	SetRequestId(v string) *ListIntegrationPolicyAddonsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListIntegrationPolicyAddonsResponseBody
	GetTotal() *int64
}

type ListIntegrationPolicyAddonsResponseBody struct {
	Addons []*ListIntegrationPolicyAddonsResponseBodyAddons `json:"addons,omitempty" xml:"addons,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 7E00EF90-CEF9-57C9-9AE9-5AA937D37C03
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 5
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListIntegrationPolicyAddonsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyAddonsResponseBody) GetAddons() []*ListIntegrationPolicyAddonsResponseBodyAddons {
	return s.Addons
}

func (s *ListIntegrationPolicyAddonsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIntegrationPolicyAddonsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListIntegrationPolicyAddonsResponseBody) SetAddons(v []*ListIntegrationPolicyAddonsResponseBodyAddons) *ListIntegrationPolicyAddonsResponseBody {
	s.Addons = v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBody) SetRequestId(v string) *ListIntegrationPolicyAddonsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBody) SetTotal(v int64) *ListIntegrationPolicyAddonsResponseBody {
	s.Total = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBody) Validate() error {
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

type ListIntegrationPolicyAddonsResponseBodyAddons struct {
	// example:
	//
	// MySQL
	Alias      *string                                                    `json:"alias,omitempty" xml:"alias,omitempty"`
	Categories []*string                                                  `json:"categories,omitempty" xml:"categories,omitempty" type:"Repeated"`
	Dashboards []*ListIntegrationPolicyAddonsResponseBodyAddonsDashboards `json:"dashboards,omitempty" xml:"dashboards,omitempty" type:"Repeated"`
	// example:
	//
	// Observability integration.
	Description  *string                                                      `json:"description,omitempty" xml:"description,omitempty"`
	Environments []*ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments `json:"environments,omitempty" xml:"environments,omitempty" type:"Repeated"`
	// example:
	//
	// asert/mysql.svg
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
	// mysql
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// true
	Once *bool `json:"once,omitempty" xml:"once,omitempty"`
	// example:
	//
	// container
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// 0.1.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// 99
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s ListIntegrationPolicyAddonsResponseBodyAddons) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyAddonsResponseBodyAddons) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) GetAlias() *string {
	return s.Alias
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) GetCategories() []*string {
	return s.Categories
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) GetDashboards() []*ListIntegrationPolicyAddonsResponseBodyAddonsDashboards {
	return s.Dashboards
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) GetDescription() *string {
	return s.Description
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) GetEnvironments() []*ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments {
	return s.Environments
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) GetIcon() *string {
	return s.Icon
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) GetKeywords() []*string {
	return s.Keywords
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) GetLanguage() *string {
	return s.Language
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) GetLatestReleaseCreateTime() *string {
	return s.LatestReleaseCreateTime
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) GetName() *string {
	return s.Name
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) GetOnce() *bool {
	return s.Once
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) GetScene() *string {
	return s.Scene
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) GetVersion() *string {
	return s.Version
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) GetWeight() *int32 {
	return s.Weight
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) SetAlias(v string) *ListIntegrationPolicyAddonsResponseBodyAddons {
	s.Alias = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) SetCategories(v []*string) *ListIntegrationPolicyAddonsResponseBodyAddons {
	s.Categories = v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) SetDashboards(v []*ListIntegrationPolicyAddonsResponseBodyAddonsDashboards) *ListIntegrationPolicyAddonsResponseBodyAddons {
	s.Dashboards = v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) SetDescription(v string) *ListIntegrationPolicyAddonsResponseBodyAddons {
	s.Description = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) SetEnvironments(v []*ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments) *ListIntegrationPolicyAddonsResponseBodyAddons {
	s.Environments = v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) SetIcon(v string) *ListIntegrationPolicyAddonsResponseBodyAddons {
	s.Icon = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) SetKeywords(v []*string) *ListIntegrationPolicyAddonsResponseBodyAddons {
	s.Keywords = v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) SetLanguage(v string) *ListIntegrationPolicyAddonsResponseBodyAddons {
	s.Language = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) SetLatestReleaseCreateTime(v string) *ListIntegrationPolicyAddonsResponseBodyAddons {
	s.LatestReleaseCreateTime = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) SetName(v string) *ListIntegrationPolicyAddonsResponseBodyAddons {
	s.Name = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) SetOnce(v bool) *ListIntegrationPolicyAddonsResponseBodyAddons {
	s.Once = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) SetScene(v string) *ListIntegrationPolicyAddonsResponseBodyAddons {
	s.Scene = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) SetVersion(v string) *ListIntegrationPolicyAddonsResponseBodyAddons {
	s.Version = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) SetWeight(v int32) *ListIntegrationPolicyAddonsResponseBodyAddons {
	s.Weight = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddons) Validate() error {
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

type ListIntegrationPolicyAddonsResponseBodyAddonsDashboards struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// MySQL Overview
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// http://xxxxxxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListIntegrationPolicyAddonsResponseBodyAddonsDashboards) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyAddonsResponseBodyAddonsDashboards) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsDashboards) GetDescription() *string {
	return s.Description
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsDashboards) GetName() *string {
	return s.Name
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsDashboards) GetUrl() *string {
	return s.Url
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsDashboards) SetDescription(v string) *ListIntegrationPolicyAddonsResponseBodyAddonsDashboards {
	s.Description = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsDashboards) SetName(v string) *ListIntegrationPolicyAddonsResponseBodyAddonsDashboards {
	s.Name = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsDashboards) SetUrl(v string) *ListIntegrationPolicyAddonsResponseBodyAddonsDashboards {
	s.Url = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsDashboards) Validate() error {
	return dara.Validate(s)
}

type ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments struct {
	Dependencies *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsDependencies `json:"dependencies,omitempty" xml:"dependencies,omitempty" type:"Struct"`
	// example:
	//
	// o11y-demo-cn-heyuan
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// Cloud
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// example:
	//
	// Cloud
	Name     *string                                                            `json:"name,omitempty" xml:"name,omitempty"`
	Policies *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies `json:"policies,omitempty" xml:"policies,omitempty" type:"Struct"`
}

func (s ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments) GetDependencies() *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsDependencies {
	return s.Dependencies
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments) GetDescription() *string {
	return s.Description
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments) GetEnable() *bool {
	return s.Enable
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments) GetLabel() *string {
	return s.Label
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments) GetName() *string {
	return s.Name
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments) GetPolicies() *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies {
	return s.Policies
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments) SetDependencies(v *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsDependencies) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments {
	s.Dependencies = v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments) SetDescription(v string) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments {
	s.Description = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments) SetEnable(v bool) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments {
	s.Enable = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments) SetLabel(v string) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments {
	s.Label = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments) SetName(v string) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments {
	s.Name = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments) SetPolicies(v *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments {
	s.Policies = v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironments) Validate() error {
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

type ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsDependencies struct {
	ClusterTypes []*string        `json:"clusterTypes,omitempty" xml:"clusterTypes,omitempty" type:"Repeated"`
	Features     map[string]*bool `json:"features,omitempty" xml:"features,omitempty"`
	Services     []*string        `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsDependencies) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsDependencies) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsDependencies) GetClusterTypes() []*string {
	return s.ClusterTypes
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsDependencies) GetFeatures() map[string]*bool {
	return s.Features
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsDependencies) GetServices() []*string {
	return s.Services
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsDependencies) SetClusterTypes(v []*string) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsDependencies {
	s.ClusterTypes = v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsDependencies) SetFeatures(v map[string]*bool) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsDependencies {
	s.Features = v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsDependencies) SetServices(v []*string) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsDependencies {
	s.Services = v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsDependencies) Validate() error {
	return dara.Validate(s)
}

type ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies struct {
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
	EnableServiceAccount *bool                                                                             `json:"enableServiceAccount,omitempty" xml:"enableServiceAccount,omitempty"`
	MetricCheckRule      *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesMetricCheckRule `json:"metricCheckRule,omitempty" xml:"metricCheckRule,omitempty" type:"Struct"`
	// example:
	//
	// true
	NeedRestartAfterIntegration *bool                                                                         `json:"needRestartAfterIntegration,omitempty" xml:"needRestartAfterIntegration,omitempty"`
	Protocols                   []*ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	// example:
	//
	// mysql
	TargetAddonName *string `json:"targetAddonName,omitempty" xml:"targetAddonName,omitempty"`
}

func (s ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies) GetAlertDefaultStatus() *string {
	return s.AlertDefaultStatus
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies) GetDefaultInstall() *bool {
	return s.DefaultInstall
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies) GetEnableServiceAccount() *bool {
	return s.EnableServiceAccount
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies) GetMetricCheckRule() *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesMetricCheckRule {
	return s.MetricCheckRule
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies) GetNeedRestartAfterIntegration() *bool {
	return s.NeedRestartAfterIntegration
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies) GetProtocols() []*ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols {
	return s.Protocols
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies) GetTargetAddonName() *string {
	return s.TargetAddonName
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies) SetAlertDefaultStatus(v string) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies {
	s.AlertDefaultStatus = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies) SetDefaultInstall(v bool) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies {
	s.DefaultInstall = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies) SetEnableServiceAccount(v bool) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies {
	s.EnableServiceAccount = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies) SetMetricCheckRule(v *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesMetricCheckRule) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies {
	s.MetricCheckRule = v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies) SetNeedRestartAfterIntegration(v bool) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies {
	s.NeedRestartAfterIntegration = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies) SetProtocols(v []*ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies {
	s.Protocols = v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies) SetTargetAddonName(v string) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies {
	s.TargetAddonName = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPolicies) Validate() error {
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

type ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesMetricCheckRule struct {
	PromQl []*string `json:"promQl,omitempty" xml:"promQl,omitempty" type:"Repeated"`
}

func (s ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesMetricCheckRule) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesMetricCheckRule) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesMetricCheckRule) GetPromQl() []*string {
	return s.PromQl
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesMetricCheckRule) SetPromQl(v []*string) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesMetricCheckRule {
	s.PromQl = v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesMetricCheckRule) Validate() error {
	return dara.Validate(s)
}

type ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols struct {
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// https://img.alixxxx
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// Golang
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// example:
	//
	// golang
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) GetDescription() *string {
	return s.Description
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) GetIcon() *string {
	return s.Icon
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) GetLabel() *string {
	return s.Label
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) GetName() *string {
	return s.Name
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) SetDescription(v string) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols {
	s.Description = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) SetIcon(v string) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols {
	s.Icon = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) SetLabel(v string) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols {
	s.Label = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) SetName(v string) *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols {
	s.Name = &v
	return s
}

func (s *ListIntegrationPolicyAddonsResponseBodyAddonsEnvironmentsPoliciesProtocols) Validate() error {
	return dara.Validate(s)
}
