// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAddonResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAddonResponseBodyData) *GetAddonResponseBody
	GetData() *GetAddonResponseBodyData
	SetRequestId(v string) *GetAddonResponseBody
	GetRequestId() *string
}

type GetAddonResponseBody struct {
	Data *GetAddonResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-A01D6CC3F4B8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAddonResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAddonResponseBody) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBody) GetData() *GetAddonResponseBodyData {
	return s.Data
}

func (s *GetAddonResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAddonResponseBody) SetData(v *GetAddonResponseBodyData) *GetAddonResponseBody {
	s.Data = v
	return s
}

func (s *GetAddonResponseBody) SetRequestId(v string) *GetAddonResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAddonResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAddonResponseBodyData struct {
	// example:
	//
	// resume_vector_alias
	Alias      *string                               `json:"alias,omitempty" xml:"alias,omitempty"`
	Categories []*string                             `json:"categories,omitempty" xml:"categories,omitempty" type:"Repeated"`
	Dashboards []*GetAddonResponseBodyDataDashboards `json:"dashboards,omitempty" xml:"dashboards,omitempty" type:"Repeated"`
	// example:
	//
	// o11y-demo-cn-heyuan
	Description  *string                                 `json:"description,omitempty" xml:"description,omitempty"`
	Environments []*GetAddonResponseBodyDataEnvironments `json:"environments,omitempty" xml:"environments,omitempty" type:"Repeated"`
	// example:
	//
	// icon URL
	Icon     *string   `json:"icon,omitempty" xml:"icon,omitempty"`
	Keywords []*string `json:"keywords,omitempty" xml:"keywords,omitempty" type:"Repeated"`
	// example:
	//
	// zh
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// 2023-09-12 12:30:30
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
	// 10
	Weight *string `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetAddonResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAddonResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBodyData) GetAlias() *string {
	return s.Alias
}

func (s *GetAddonResponseBodyData) GetCategories() []*string {
	return s.Categories
}

func (s *GetAddonResponseBodyData) GetDashboards() []*GetAddonResponseBodyDataDashboards {
	return s.Dashboards
}

func (s *GetAddonResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetAddonResponseBodyData) GetEnvironments() []*GetAddonResponseBodyDataEnvironments {
	return s.Environments
}

func (s *GetAddonResponseBodyData) GetIcon() *string {
	return s.Icon
}

func (s *GetAddonResponseBodyData) GetKeywords() []*string {
	return s.Keywords
}

func (s *GetAddonResponseBodyData) GetLanguage() *string {
	return s.Language
}

func (s *GetAddonResponseBodyData) GetLatestReleaseCreateTime() *string {
	return s.LatestReleaseCreateTime
}

func (s *GetAddonResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetAddonResponseBodyData) GetOnce() *bool {
	return s.Once
}

func (s *GetAddonResponseBodyData) GetScene() *string {
	return s.Scene
}

func (s *GetAddonResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetAddonResponseBodyData) GetWeight() *string {
	return s.Weight
}

func (s *GetAddonResponseBodyData) SetAlias(v string) *GetAddonResponseBodyData {
	s.Alias = &v
	return s
}

func (s *GetAddonResponseBodyData) SetCategories(v []*string) *GetAddonResponseBodyData {
	s.Categories = v
	return s
}

func (s *GetAddonResponseBodyData) SetDashboards(v []*GetAddonResponseBodyDataDashboards) *GetAddonResponseBodyData {
	s.Dashboards = v
	return s
}

func (s *GetAddonResponseBodyData) SetDescription(v string) *GetAddonResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetAddonResponseBodyData) SetEnvironments(v []*GetAddonResponseBodyDataEnvironments) *GetAddonResponseBodyData {
	s.Environments = v
	return s
}

func (s *GetAddonResponseBodyData) SetIcon(v string) *GetAddonResponseBodyData {
	s.Icon = &v
	return s
}

func (s *GetAddonResponseBodyData) SetKeywords(v []*string) *GetAddonResponseBodyData {
	s.Keywords = v
	return s
}

func (s *GetAddonResponseBodyData) SetLanguage(v string) *GetAddonResponseBodyData {
	s.Language = &v
	return s
}

func (s *GetAddonResponseBodyData) SetLatestReleaseCreateTime(v string) *GetAddonResponseBodyData {
	s.LatestReleaseCreateTime = &v
	return s
}

func (s *GetAddonResponseBodyData) SetName(v string) *GetAddonResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetAddonResponseBodyData) SetOnce(v bool) *GetAddonResponseBodyData {
	s.Once = &v
	return s
}

func (s *GetAddonResponseBodyData) SetScene(v string) *GetAddonResponseBodyData {
	s.Scene = &v
	return s
}

func (s *GetAddonResponseBodyData) SetVersion(v string) *GetAddonResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetAddonResponseBodyData) SetWeight(v string) *GetAddonResponseBodyData {
	s.Weight = &v
	return s
}

func (s *GetAddonResponseBodyData) Validate() error {
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

type GetAddonResponseBodyDataDashboards struct {
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// rum_api_dot_metric_set
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// http://iac-service-transfer.oss-cn-hangzhou.aliyuncs.com/78c5_288850010070719968_bdcf7cca781844c8ac4add133791713f
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetAddonResponseBodyDataDashboards) String() string {
	return dara.Prettify(s)
}

func (s GetAddonResponseBodyDataDashboards) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBodyDataDashboards) GetDescription() *string {
	return s.Description
}

func (s *GetAddonResponseBodyDataDashboards) GetName() *string {
	return s.Name
}

func (s *GetAddonResponseBodyDataDashboards) GetUrl() *string {
	return s.Url
}

func (s *GetAddonResponseBodyDataDashboards) SetDescription(v string) *GetAddonResponseBodyDataDashboards {
	s.Description = &v
	return s
}

func (s *GetAddonResponseBodyDataDashboards) SetName(v string) *GetAddonResponseBodyDataDashboards {
	s.Name = &v
	return s
}

func (s *GetAddonResponseBodyDataDashboards) SetUrl(v string) *GetAddonResponseBodyDataDashboards {
	s.Url = &v
	return s
}

func (s *GetAddonResponseBodyDataDashboards) Validate() error {
	return dara.Validate(s)
}

type GetAddonResponseBodyDataEnvironments struct {
	CommonSchemaRefs []*GetAddonResponseBodyDataEnvironmentsCommonSchemaRefs `json:"commonSchemaRefs,omitempty" xml:"commonSchemaRefs,omitempty" type:"Repeated"`
	Dependencies     *GetAddonResponseBodyDataEnvironmentsDependencies       `json:"dependencies,omitempty" xml:"dependencies,omitempty" type:"Struct"`
	// example:
	//
	// Observability integration.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// false
	Enable *bool   `json:"enable,omitempty" xml:"enable,omitempty"`
	Label  *string `json:"label,omitempty" xml:"label,omitempty"`
	// example:
	//
	// cs-default-umodel-1752755202744_k8s.metric.high_level_metric_deployment_cn-hangzhou/c0a686118449c4390b9cc0a07ea3e0e85
	Name     *string                                       `json:"name,omitempty" xml:"name,omitempty"`
	Policies *GetAddonResponseBodyDataEnvironmentsPolicies `json:"policies,omitempty" xml:"policies,omitempty" type:"Struct"`
	// example:
	//
	// CS
	PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty"`
}

func (s GetAddonResponseBodyDataEnvironments) String() string {
	return dara.Prettify(s)
}

func (s GetAddonResponseBodyDataEnvironments) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBodyDataEnvironments) GetCommonSchemaRefs() []*GetAddonResponseBodyDataEnvironmentsCommonSchemaRefs {
	return s.CommonSchemaRefs
}

func (s *GetAddonResponseBodyDataEnvironments) GetDependencies() *GetAddonResponseBodyDataEnvironmentsDependencies {
	return s.Dependencies
}

func (s *GetAddonResponseBodyDataEnvironments) GetDescription() *string {
	return s.Description
}

func (s *GetAddonResponseBodyDataEnvironments) GetEnable() *bool {
	return s.Enable
}

func (s *GetAddonResponseBodyDataEnvironments) GetLabel() *string {
	return s.Label
}

func (s *GetAddonResponseBodyDataEnvironments) GetName() *string {
	return s.Name
}

func (s *GetAddonResponseBodyDataEnvironments) GetPolicies() *GetAddonResponseBodyDataEnvironmentsPolicies {
	return s.Policies
}

func (s *GetAddonResponseBodyDataEnvironments) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetAddonResponseBodyDataEnvironments) SetCommonSchemaRefs(v []*GetAddonResponseBodyDataEnvironmentsCommonSchemaRefs) *GetAddonResponseBodyDataEnvironments {
	s.CommonSchemaRefs = v
	return s
}

func (s *GetAddonResponseBodyDataEnvironments) SetDependencies(v *GetAddonResponseBodyDataEnvironmentsDependencies) *GetAddonResponseBodyDataEnvironments {
	s.Dependencies = v
	return s
}

func (s *GetAddonResponseBodyDataEnvironments) SetDescription(v string) *GetAddonResponseBodyDataEnvironments {
	s.Description = &v
	return s
}

func (s *GetAddonResponseBodyDataEnvironments) SetEnable(v bool) *GetAddonResponseBodyDataEnvironments {
	s.Enable = &v
	return s
}

func (s *GetAddonResponseBodyDataEnvironments) SetLabel(v string) *GetAddonResponseBodyDataEnvironments {
	s.Label = &v
	return s
}

func (s *GetAddonResponseBodyDataEnvironments) SetName(v string) *GetAddonResponseBodyDataEnvironments {
	s.Name = &v
	return s
}

func (s *GetAddonResponseBodyDataEnvironments) SetPolicies(v *GetAddonResponseBodyDataEnvironmentsPolicies) *GetAddonResponseBodyDataEnvironments {
	s.Policies = v
	return s
}

func (s *GetAddonResponseBodyDataEnvironments) SetPolicyType(v string) *GetAddonResponseBodyDataEnvironments {
	s.PolicyType = &v
	return s
}

func (s *GetAddonResponseBodyDataEnvironments) Validate() error {
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

type GetAddonResponseBodyDataEnvironmentsCommonSchemaRefs struct {
	Group   *string `json:"group,omitempty" xml:"group,omitempty"`
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetAddonResponseBodyDataEnvironmentsCommonSchemaRefs) String() string {
	return dara.Prettify(s)
}

func (s GetAddonResponseBodyDataEnvironmentsCommonSchemaRefs) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBodyDataEnvironmentsCommonSchemaRefs) GetGroup() *string {
	return s.Group
}

func (s *GetAddonResponseBodyDataEnvironmentsCommonSchemaRefs) GetVersion() *string {
	return s.Version
}

func (s *GetAddonResponseBodyDataEnvironmentsCommonSchemaRefs) SetGroup(v string) *GetAddonResponseBodyDataEnvironmentsCommonSchemaRefs {
	s.Group = &v
	return s
}

func (s *GetAddonResponseBodyDataEnvironmentsCommonSchemaRefs) SetVersion(v string) *GetAddonResponseBodyDataEnvironmentsCommonSchemaRefs {
	s.Version = &v
	return s
}

func (s *GetAddonResponseBodyDataEnvironmentsCommonSchemaRefs) Validate() error {
	return dara.Validate(s)
}

type GetAddonResponseBodyDataEnvironmentsDependencies struct {
	ClusterTypes []*string        `json:"clusterTypes,omitempty" xml:"clusterTypes,omitempty" type:"Repeated"`
	Features     map[string]*bool `json:"features,omitempty" xml:"features,omitempty"`
	Services     []*string        `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s GetAddonResponseBodyDataEnvironmentsDependencies) String() string {
	return dara.Prettify(s)
}

func (s GetAddonResponseBodyDataEnvironmentsDependencies) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBodyDataEnvironmentsDependencies) GetClusterTypes() []*string {
	return s.ClusterTypes
}

func (s *GetAddonResponseBodyDataEnvironmentsDependencies) GetFeatures() map[string]*bool {
	return s.Features
}

func (s *GetAddonResponseBodyDataEnvironmentsDependencies) GetServices() []*string {
	return s.Services
}

func (s *GetAddonResponseBodyDataEnvironmentsDependencies) SetClusterTypes(v []*string) *GetAddonResponseBodyDataEnvironmentsDependencies {
	s.ClusterTypes = v
	return s
}

func (s *GetAddonResponseBodyDataEnvironmentsDependencies) SetFeatures(v map[string]*bool) *GetAddonResponseBodyDataEnvironmentsDependencies {
	s.Features = v
	return s
}

func (s *GetAddonResponseBodyDataEnvironmentsDependencies) SetServices(v []*string) *GetAddonResponseBodyDataEnvironmentsDependencies {
	s.Services = v
	return s
}

func (s *GetAddonResponseBodyDataEnvironmentsDependencies) Validate() error {
	return dara.Validate(s)
}

type GetAddonResponseBodyDataEnvironmentsPolicies struct {
	// example:
	//
	// true
	AlertDefaultStatus *string `json:"alertDefaultStatus,omitempty" xml:"alertDefaultStatus,omitempty"`
	// example:
	//
	// true
	DefaultInstall *bool `json:"defaultInstall,omitempty" xml:"defaultInstall,omitempty"`
	// example:
	//
	// true
	EnableServiceAccount *bool                                                        `json:"enableServiceAccount,omitempty" xml:"enableServiceAccount,omitempty"`
	MetricCheckRule      *GetAddonResponseBodyDataEnvironmentsPoliciesMetricCheckRule `json:"metricCheckRule,omitempty" xml:"metricCheckRule,omitempty" type:"Struct"`
	// example:
	//
	// true
	NeedRestartAfterIntegration *bool                                                    `json:"needRestartAfterIntegration,omitempty" xml:"needRestartAfterIntegration,omitempty"`
	Protocols                   []*GetAddonResponseBodyDataEnvironmentsPoliciesProtocols `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	// example:
	//
	// cs-default
	TargetAddonName *string `json:"targetAddonName,omitempty" xml:"targetAddonName,omitempty"`
}

func (s GetAddonResponseBodyDataEnvironmentsPolicies) String() string {
	return dara.Prettify(s)
}

func (s GetAddonResponseBodyDataEnvironmentsPolicies) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBodyDataEnvironmentsPolicies) GetAlertDefaultStatus() *string {
	return s.AlertDefaultStatus
}

func (s *GetAddonResponseBodyDataEnvironmentsPolicies) GetDefaultInstall() *bool {
	return s.DefaultInstall
}

func (s *GetAddonResponseBodyDataEnvironmentsPolicies) GetEnableServiceAccount() *bool {
	return s.EnableServiceAccount
}

func (s *GetAddonResponseBodyDataEnvironmentsPolicies) GetMetricCheckRule() *GetAddonResponseBodyDataEnvironmentsPoliciesMetricCheckRule {
	return s.MetricCheckRule
}

func (s *GetAddonResponseBodyDataEnvironmentsPolicies) GetNeedRestartAfterIntegration() *bool {
	return s.NeedRestartAfterIntegration
}

func (s *GetAddonResponseBodyDataEnvironmentsPolicies) GetProtocols() []*GetAddonResponseBodyDataEnvironmentsPoliciesProtocols {
	return s.Protocols
}

func (s *GetAddonResponseBodyDataEnvironmentsPolicies) GetTargetAddonName() *string {
	return s.TargetAddonName
}

func (s *GetAddonResponseBodyDataEnvironmentsPolicies) SetAlertDefaultStatus(v string) *GetAddonResponseBodyDataEnvironmentsPolicies {
	s.AlertDefaultStatus = &v
	return s
}

func (s *GetAddonResponseBodyDataEnvironmentsPolicies) SetDefaultInstall(v bool) *GetAddonResponseBodyDataEnvironmentsPolicies {
	s.DefaultInstall = &v
	return s
}

func (s *GetAddonResponseBodyDataEnvironmentsPolicies) SetEnableServiceAccount(v bool) *GetAddonResponseBodyDataEnvironmentsPolicies {
	s.EnableServiceAccount = &v
	return s
}

func (s *GetAddonResponseBodyDataEnvironmentsPolicies) SetMetricCheckRule(v *GetAddonResponseBodyDataEnvironmentsPoliciesMetricCheckRule) *GetAddonResponseBodyDataEnvironmentsPolicies {
	s.MetricCheckRule = v
	return s
}

func (s *GetAddonResponseBodyDataEnvironmentsPolicies) SetNeedRestartAfterIntegration(v bool) *GetAddonResponseBodyDataEnvironmentsPolicies {
	s.NeedRestartAfterIntegration = &v
	return s
}

func (s *GetAddonResponseBodyDataEnvironmentsPolicies) SetProtocols(v []*GetAddonResponseBodyDataEnvironmentsPoliciesProtocols) *GetAddonResponseBodyDataEnvironmentsPolicies {
	s.Protocols = v
	return s
}

func (s *GetAddonResponseBodyDataEnvironmentsPolicies) SetTargetAddonName(v string) *GetAddonResponseBodyDataEnvironmentsPolicies {
	s.TargetAddonName = &v
	return s
}

func (s *GetAddonResponseBodyDataEnvironmentsPolicies) Validate() error {
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

type GetAddonResponseBodyDataEnvironmentsPoliciesMetricCheckRule struct {
	PromQL []*string `json:"promQL,omitempty" xml:"promQL,omitempty" type:"Repeated"`
}

func (s GetAddonResponseBodyDataEnvironmentsPoliciesMetricCheckRule) String() string {
	return dara.Prettify(s)
}

func (s GetAddonResponseBodyDataEnvironmentsPoliciesMetricCheckRule) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBodyDataEnvironmentsPoliciesMetricCheckRule) GetPromQL() []*string {
	return s.PromQL
}

func (s *GetAddonResponseBodyDataEnvironmentsPoliciesMetricCheckRule) SetPromQL(v []*string) *GetAddonResponseBodyDataEnvironmentsPoliciesMetricCheckRule {
	s.PromQL = v
	return s
}

func (s *GetAddonResponseBodyDataEnvironmentsPoliciesMetricCheckRule) Validate() error {
	return dara.Validate(s)
}

type GetAddonResponseBodyDataEnvironmentsPoliciesProtocols struct {
	// example:
	//
	// Prometheus Metrics
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// icon URL
	Icon  *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// example:
	//
	// rum_api_dot_metric_set
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetAddonResponseBodyDataEnvironmentsPoliciesProtocols) String() string {
	return dara.Prettify(s)
}

func (s GetAddonResponseBodyDataEnvironmentsPoliciesProtocols) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBodyDataEnvironmentsPoliciesProtocols) GetDescription() *string {
	return s.Description
}

func (s *GetAddonResponseBodyDataEnvironmentsPoliciesProtocols) GetIcon() *string {
	return s.Icon
}

func (s *GetAddonResponseBodyDataEnvironmentsPoliciesProtocols) GetLabel() *string {
	return s.Label
}

func (s *GetAddonResponseBodyDataEnvironmentsPoliciesProtocols) GetName() *string {
	return s.Name
}

func (s *GetAddonResponseBodyDataEnvironmentsPoliciesProtocols) SetDescription(v string) *GetAddonResponseBodyDataEnvironmentsPoliciesProtocols {
	s.Description = &v
	return s
}

func (s *GetAddonResponseBodyDataEnvironmentsPoliciesProtocols) SetIcon(v string) *GetAddonResponseBodyDataEnvironmentsPoliciesProtocols {
	s.Icon = &v
	return s
}

func (s *GetAddonResponseBodyDataEnvironmentsPoliciesProtocols) SetLabel(v string) *GetAddonResponseBodyDataEnvironmentsPoliciesProtocols {
	s.Label = &v
	return s
}

func (s *GetAddonResponseBodyDataEnvironmentsPoliciesProtocols) SetName(v string) *GetAddonResponseBodyDataEnvironmentsPoliciesProtocols {
	s.Name = &v
	return s
}

func (s *GetAddonResponseBodyDataEnvironmentsPoliciesProtocols) Validate() error {
	return dara.Validate(s)
}
