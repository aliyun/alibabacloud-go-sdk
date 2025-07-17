// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentAddonsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListEnvironmentAddonsResponseBody
	GetCode() *int32
	SetData(v *ListEnvironmentAddonsResponseBodyData) *ListEnvironmentAddonsResponseBody
	GetData() *ListEnvironmentAddonsResponseBodyData
	SetMessage(v string) *ListEnvironmentAddonsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEnvironmentAddonsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEnvironmentAddonsResponseBody
	GetSuccess() *bool
}

type ListEnvironmentAddonsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the operation.
	Data *ListEnvironmentAddonsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 32940175-181B-4B93-966E-4BB69176****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEnvironmentAddonsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentAddonsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListEnvironmentAddonsResponseBody) GetData() *ListEnvironmentAddonsResponseBodyData {
	return s.Data
}

func (s *ListEnvironmentAddonsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEnvironmentAddonsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnvironmentAddonsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEnvironmentAddonsResponseBody) SetCode(v int32) *ListEnvironmentAddonsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBody) SetData(v *ListEnvironmentAddonsResponseBodyData) *ListEnvironmentAddonsResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentAddonsResponseBody) SetMessage(v string) *ListEnvironmentAddonsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBody) SetRequestId(v string) *ListEnvironmentAddonsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBody) SetSuccess(v bool) *ListEnvironmentAddonsResponseBody {
	s.Success = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEnvironmentAddonsResponseBodyData struct {
	// The queried add-ons.
	Addons []*ListEnvironmentAddonsResponseBodyDataAddons `json:"Addons,omitempty" xml:"Addons,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListEnvironmentAddonsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentAddonsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvironmentAddonsResponseBodyData) GetAddons() []*ListEnvironmentAddonsResponseBodyDataAddons {
	return s.Addons
}

func (s *ListEnvironmentAddonsResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListEnvironmentAddonsResponseBodyData) SetAddons(v []*ListEnvironmentAddonsResponseBodyDataAddons) *ListEnvironmentAddonsResponseBodyData {
	s.Addons = v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyData) SetTotal(v int64) *ListEnvironmentAddonsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListEnvironmentAddonsResponseBodyDataAddons struct {
	// The alias of the add-on.
	//
	// example:
	//
	// MySQL
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The tags of the add-on.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The dashboards.
	Dashboards []*ListEnvironmentAddonsResponseBodyDataAddonsDashboards `json:"Dashboards,omitempty" xml:"Dashboards,omitempty" type:"Repeated"`
	// The description of the add-on.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The supported environments.
	Environments []*ListEnvironmentAddonsResponseBodyDataAddonsEnvironments `json:"Environments,omitempty" xml:"Environments,omitempty" type:"Repeated"`
	// The URL of the icon.
	//
	// example:
	//
	// http://xxxx
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The collection of keywords.
	Keywords []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	// The language.
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The time when the instance was last created.
	//
	// example:
	//
	// 2023-09-22T16:56:29+08:00
	LatestReleaseCreateTime *string `json:"LatestReleaseCreateTime,omitempty" xml:"LatestReleaseCreateTime,omitempty"`
	// The name of the add-on.
	//
	// example:
	//
	// mysql
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the add-on can be installed only once.
	//
	// example:
	//
	// false
	Once *bool `json:"Once,omitempty" xml:"Once,omitempty"`
	// The scenario.
	//
	// example:
	//
	// database
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The version of the agent.
	//
	// example:
	//
	// 0.0.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The weight.
	//
	// example:
	//
	// 857
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListEnvironmentAddonsResponseBodyDataAddons) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentAddonsResponseBodyDataAddons) GoString() string {
	return s.String()
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) GetAlias() *string {
	return s.Alias
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) GetCategories() []*string {
	return s.Categories
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) GetDashboards() []*ListEnvironmentAddonsResponseBodyDataAddonsDashboards {
	return s.Dashboards
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) GetDescription() *string {
	return s.Description
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) GetEnvironments() []*ListEnvironmentAddonsResponseBodyDataAddonsEnvironments {
	return s.Environments
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) GetIcon() *string {
	return s.Icon
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) GetKeywords() []*string {
	return s.Keywords
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) GetLanguage() *string {
	return s.Language
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) GetLatestReleaseCreateTime() *string {
	return s.LatestReleaseCreateTime
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) GetName() *string {
	return s.Name
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) GetOnce() *bool {
	return s.Once
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) GetScene() *string {
	return s.Scene
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) GetVersion() *string {
	return s.Version
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) GetWeight() *string {
	return s.Weight
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) SetAlias(v string) *ListEnvironmentAddonsResponseBodyDataAddons {
	s.Alias = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) SetCategories(v []*string) *ListEnvironmentAddonsResponseBodyDataAddons {
	s.Categories = v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) SetDashboards(v []*ListEnvironmentAddonsResponseBodyDataAddonsDashboards) *ListEnvironmentAddonsResponseBodyDataAddons {
	s.Dashboards = v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) SetDescription(v string) *ListEnvironmentAddonsResponseBodyDataAddons {
	s.Description = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) SetEnvironments(v []*ListEnvironmentAddonsResponseBodyDataAddonsEnvironments) *ListEnvironmentAddonsResponseBodyDataAddons {
	s.Environments = v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) SetIcon(v string) *ListEnvironmentAddonsResponseBodyDataAddons {
	s.Icon = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) SetKeywords(v []*string) *ListEnvironmentAddonsResponseBodyDataAddons {
	s.Keywords = v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) SetLanguage(v string) *ListEnvironmentAddonsResponseBodyDataAddons {
	s.Language = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) SetLatestReleaseCreateTime(v string) *ListEnvironmentAddonsResponseBodyDataAddons {
	s.LatestReleaseCreateTime = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) SetName(v string) *ListEnvironmentAddonsResponseBodyDataAddons {
	s.Name = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) SetOnce(v bool) *ListEnvironmentAddonsResponseBodyDataAddons {
	s.Once = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) SetScene(v string) *ListEnvironmentAddonsResponseBodyDataAddons {
	s.Scene = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) SetVersion(v string) *ListEnvironmentAddonsResponseBodyDataAddons {
	s.Version = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) SetWeight(v string) *ListEnvironmentAddonsResponseBodyDataAddons {
	s.Weight = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddons) Validate() error {
	return dara.Validate(s)
}

type ListEnvironmentAddonsResponseBodyDataAddonsDashboards struct {
	// The description of the dashboard.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the dashboard.
	//
	// example:
	//
	// mysql-overview
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The URL of the dashboard.
	//
	// example:
	//
	// http://xxxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListEnvironmentAddonsResponseBodyDataAddonsDashboards) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentAddonsResponseBodyDataAddonsDashboards) GoString() string {
	return s.String()
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsDashboards) GetDescription() *string {
	return s.Description
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsDashboards) GetName() *string {
	return s.Name
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsDashboards) GetUrl() *string {
	return s.Url
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsDashboards) SetDescription(v string) *ListEnvironmentAddonsResponseBodyDataAddonsDashboards {
	s.Description = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsDashboards) SetName(v string) *ListEnvironmentAddonsResponseBodyDataAddonsDashboards {
	s.Name = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsDashboards) SetUrl(v string) *ListEnvironmentAddonsResponseBodyDataAddonsDashboards {
	s.Url = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsDashboards) Validate() error {
	return dara.Validate(s)
}

type ListEnvironmentAddonsResponseBodyDataAddonsEnvironments struct {
	// The dependencies of the environment.
	Dependencies *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsDependencies `json:"Dependencies,omitempty" xml:"Dependencies,omitempty" type:"Struct"`
	// The description of the environment.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the feature is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The tag of the environment.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The name of the environment.
	//
	// example:
	//
	// CS
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The control policies in the environment.
	Policies *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
}

func (s ListEnvironmentAddonsResponseBodyDataAddonsEnvironments) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentAddonsResponseBodyDataAddonsEnvironments) GoString() string {
	return s.String()
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironments) GetDependencies() *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsDependencies {
	return s.Dependencies
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironments) GetDescription() *string {
	return s.Description
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironments) GetEnable() *bool {
	return s.Enable
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironments) GetLabel() *string {
	return s.Label
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironments) GetName() *string {
	return s.Name
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironments) GetPolicies() *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies {
	return s.Policies
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironments) SetDependencies(v *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsDependencies) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironments {
	s.Dependencies = v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironments) SetDescription(v string) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironments {
	s.Description = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironments) SetEnable(v bool) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironments {
	s.Enable = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironments) SetLabel(v string) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironments {
	s.Label = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironments) SetName(v string) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironments {
	s.Name = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironments) SetPolicies(v *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironments {
	s.Policies = v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironments) Validate() error {
	return dara.Validate(s)
}

type ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsDependencies struct {
	// The cluster type.
	ClusterTypes []*string `json:"ClusterTypes,omitempty" xml:"ClusterTypes,omitempty" type:"Repeated"`
	// The feature that can be installed in the environment.
	Features map[string]*bool `json:"Features,omitempty" xml:"Features,omitempty"`
	// The services.
	Services []*string `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
}

func (s ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsDependencies) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsDependencies) GoString() string {
	return s.String()
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsDependencies) GetClusterTypes() []*string {
	return s.ClusterTypes
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsDependencies) GetFeatures() map[string]*bool {
	return s.Features
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsDependencies) GetServices() []*string {
	return s.Services
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsDependencies) SetClusterTypes(v []*string) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsDependencies {
	s.ClusterTypes = v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsDependencies) SetFeatures(v map[string]*bool) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsDependencies {
	s.Features = v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsDependencies) SetServices(v []*string) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsDependencies {
	s.Services = v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsDependencies) Validate() error {
	return dara.Validate(s)
}

type ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies struct {
	// The default alert status.
	//
	// example:
	//
	// default
	AlertDefaultStatus *string `json:"AlertDefaultStatus,omitempty" xml:"AlertDefaultStatus,omitempty"`
	// The default installation status.
	//
	// example:
	//
	// false
	DefaultInstall *bool `json:"DefaultInstall,omitempty" xml:"DefaultInstall,omitempty"`
	// Indicates whether a service account is enabled.
	//
	// example:
	//
	// true
	EnableServiceAccount *bool `json:"EnableServiceAccount,omitempty" xml:"EnableServiceAccount,omitempty"`
	// The metric check rule.
	MetricCheckRule *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesMetricCheckRule `json:"MetricCheckRule,omitempty" xml:"MetricCheckRule,omitempty" type:"Struct"`
	// Indicates whether a restart is required after the installation.
	//
	// example:
	//
	// true
	NeedRestartAfterIntegration *bool `json:"NeedRestartAfterIntegration,omitempty" xml:"NeedRestartAfterIntegration,omitempty"`
	// The supported protocols.
	Protocols []*ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesProtocols `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	// The target name of the add-on.
	//
	// example:
	//
	// cloud-rds-mysql
	TargetAddonName *string `json:"TargetAddonName,omitempty" xml:"TargetAddonName,omitempty"`
}

func (s ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies) GoString() string {
	return s.String()
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies) GetAlertDefaultStatus() *string {
	return s.AlertDefaultStatus
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies) GetDefaultInstall() *bool {
	return s.DefaultInstall
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies) GetEnableServiceAccount() *bool {
	return s.EnableServiceAccount
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies) GetMetricCheckRule() *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesMetricCheckRule {
	return s.MetricCheckRule
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies) GetNeedRestartAfterIntegration() *bool {
	return s.NeedRestartAfterIntegration
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies) GetProtocols() []*ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesProtocols {
	return s.Protocols
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies) GetTargetAddonName() *string {
	return s.TargetAddonName
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies) SetAlertDefaultStatus(v string) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies {
	s.AlertDefaultStatus = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies) SetDefaultInstall(v bool) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies {
	s.DefaultInstall = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies) SetEnableServiceAccount(v bool) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies {
	s.EnableServiceAccount = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies) SetMetricCheckRule(v *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesMetricCheckRule) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies {
	s.MetricCheckRule = v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies) SetNeedRestartAfterIntegration(v bool) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies {
	s.NeedRestartAfterIntegration = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies) SetProtocols(v []*ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesProtocols) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies {
	s.Protocols = v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies) SetTargetAddonName(v string) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies {
	s.TargetAddonName = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPolicies) Validate() error {
	return dara.Validate(s)
}

type ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesMetricCheckRule struct {
	// The PromQL statements.
	PromQL []*string `json:"PromQL,omitempty" xml:"PromQL,omitempty" type:"Repeated"`
}

func (s ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesMetricCheckRule) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesMetricCheckRule) GoString() string {
	return s.String()
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesMetricCheckRule) GetPromQL() []*string {
	return s.PromQL
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesMetricCheckRule) SetPromQL(v []*string) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesMetricCheckRule {
	s.PromQL = v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesMetricCheckRule) Validate() error {
	return dara.Validate(s)
}

type ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesProtocols struct {
	// The description of the protocol.
	//
	// example:
	//
	// ARMS
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the protocol icon.
	//
	// example:
	//
	// http://xxxxxxx
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The tag of the protocol.
	//
	// example:
	//
	// ARMS
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The name of the protocol.
	//
	// example:
	//
	// arms
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesProtocols) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesProtocols) GoString() string {
	return s.String()
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesProtocols) GetDescription() *string {
	return s.Description
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesProtocols) GetIcon() *string {
	return s.Icon
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesProtocols) GetLabel() *string {
	return s.Label
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesProtocols) GetName() *string {
	return s.Name
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesProtocols) SetDescription(v string) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesProtocols {
	s.Description = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesProtocols) SetIcon(v string) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesProtocols {
	s.Icon = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesProtocols) SetLabel(v string) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesProtocols {
	s.Label = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesProtocols) SetName(v string) *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesProtocols {
	s.Name = &v
	return s
}

func (s *ListEnvironmentAddonsResponseBodyDataAddonsEnvironmentsPoliciesProtocols) Validate() error {
	return dara.Validate(s)
}
