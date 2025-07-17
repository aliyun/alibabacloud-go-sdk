// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddonsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListAddonsResponseBody
	GetCode() *int32
	SetData(v []*ListAddonsResponseBodyData) *ListAddonsResponseBody
	GetData() []*ListAddonsResponseBodyData
	SetMessage(v string) *ListAddonsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAddonsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAddonsResponseBody
	GetSuccess() *bool
}

type ListAddonsResponseBody struct {
	// Status code: 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried add-ons.
	Data []*ListAddonsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 78901766-3806-4E96-8E47-CFEF59E4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the alert rule was deleted. Valid values:
	//
	// 	- `true`: The alert rule was deleted.
	//
	// 	- `false`: The alert rule failed to be deleted.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAddonsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListAddonsResponseBody) GetData() []*ListAddonsResponseBodyData {
	return s.Data
}

func (s *ListAddonsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAddonsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAddonsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAddonsResponseBody) SetCode(v int32) *ListAddonsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAddonsResponseBody) SetData(v []*ListAddonsResponseBodyData) *ListAddonsResponseBody {
	s.Data = v
	return s
}

func (s *ListAddonsResponseBody) SetMessage(v string) *ListAddonsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAddonsResponseBody) SetRequestId(v string) *ListAddonsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAddonsResponseBody) SetSuccess(v bool) *ListAddonsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAddonsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAddonsResponseBodyData struct {
	// The alias of the add-on.
	//
	// example:
	//
	// MySQL
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The tags of the add-on.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The dashboards.
	Dashboards []*ListAddonsResponseBodyDataDashboards `json:"Dashboards,omitempty" xml:"Dashboards,omitempty" type:"Repeated"`
	// The description of the add-on.
	//
	// example:
	//
	// Monitor database indicators with MySQL Exporter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The supported environments.
	Environments []*ListAddonsResponseBodyDataEnvironments `json:"Environments,omitempty" xml:"Environments,omitempty" type:"Repeated"`
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

func (s ListAddonsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBodyData) GetAlias() *string {
	return s.Alias
}

func (s *ListAddonsResponseBodyData) GetCategories() []*string {
	return s.Categories
}

func (s *ListAddonsResponseBodyData) GetDashboards() []*ListAddonsResponseBodyDataDashboards {
	return s.Dashboards
}

func (s *ListAddonsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListAddonsResponseBodyData) GetEnvironments() []*ListAddonsResponseBodyDataEnvironments {
	return s.Environments
}

func (s *ListAddonsResponseBodyData) GetIcon() *string {
	return s.Icon
}

func (s *ListAddonsResponseBodyData) GetKeywords() []*string {
	return s.Keywords
}

func (s *ListAddonsResponseBodyData) GetLanguage() *string {
	return s.Language
}

func (s *ListAddonsResponseBodyData) GetLatestReleaseCreateTime() *string {
	return s.LatestReleaseCreateTime
}

func (s *ListAddonsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListAddonsResponseBodyData) GetOnce() *bool {
	return s.Once
}

func (s *ListAddonsResponseBodyData) GetScene() *string {
	return s.Scene
}

func (s *ListAddonsResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *ListAddonsResponseBodyData) GetWeight() *string {
	return s.Weight
}

func (s *ListAddonsResponseBodyData) SetAlias(v string) *ListAddonsResponseBodyData {
	s.Alias = &v
	return s
}

func (s *ListAddonsResponseBodyData) SetCategories(v []*string) *ListAddonsResponseBodyData {
	s.Categories = v
	return s
}

func (s *ListAddonsResponseBodyData) SetDashboards(v []*ListAddonsResponseBodyDataDashboards) *ListAddonsResponseBodyData {
	s.Dashboards = v
	return s
}

func (s *ListAddonsResponseBodyData) SetDescription(v string) *ListAddonsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListAddonsResponseBodyData) SetEnvironments(v []*ListAddonsResponseBodyDataEnvironments) *ListAddonsResponseBodyData {
	s.Environments = v
	return s
}

func (s *ListAddonsResponseBodyData) SetIcon(v string) *ListAddonsResponseBodyData {
	s.Icon = &v
	return s
}

func (s *ListAddonsResponseBodyData) SetKeywords(v []*string) *ListAddonsResponseBodyData {
	s.Keywords = v
	return s
}

func (s *ListAddonsResponseBodyData) SetLanguage(v string) *ListAddonsResponseBodyData {
	s.Language = &v
	return s
}

func (s *ListAddonsResponseBodyData) SetLatestReleaseCreateTime(v string) *ListAddonsResponseBodyData {
	s.LatestReleaseCreateTime = &v
	return s
}

func (s *ListAddonsResponseBodyData) SetName(v string) *ListAddonsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListAddonsResponseBodyData) SetOnce(v bool) *ListAddonsResponseBodyData {
	s.Once = &v
	return s
}

func (s *ListAddonsResponseBodyData) SetScene(v string) *ListAddonsResponseBodyData {
	s.Scene = &v
	return s
}

func (s *ListAddonsResponseBodyData) SetVersion(v string) *ListAddonsResponseBodyData {
	s.Version = &v
	return s
}

func (s *ListAddonsResponseBodyData) SetWeight(v string) *ListAddonsResponseBodyData {
	s.Weight = &v
	return s
}

func (s *ListAddonsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListAddonsResponseBodyDataDashboards struct {
	// The description of the dashboard.
	//
	// example:
	//
	// MySQL monitors the market information, monitoring the connection information, usage information and other indicators
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

func (s ListAddonsResponseBodyDataDashboards) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBodyDataDashboards) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBodyDataDashboards) GetDescription() *string {
	return s.Description
}

func (s *ListAddonsResponseBodyDataDashboards) GetName() *string {
	return s.Name
}

func (s *ListAddonsResponseBodyDataDashboards) GetUrl() *string {
	return s.Url
}

func (s *ListAddonsResponseBodyDataDashboards) SetDescription(v string) *ListAddonsResponseBodyDataDashboards {
	s.Description = &v
	return s
}

func (s *ListAddonsResponseBodyDataDashboards) SetName(v string) *ListAddonsResponseBodyDataDashboards {
	s.Name = &v
	return s
}

func (s *ListAddonsResponseBodyDataDashboards) SetUrl(v string) *ListAddonsResponseBodyDataDashboards {
	s.Url = &v
	return s
}

func (s *ListAddonsResponseBodyDataDashboards) Validate() error {
	return dara.Validate(s)
}

type ListAddonsResponseBodyDataEnvironments struct {
	// The dependencies of the environment.
	Dependencies *ListAddonsResponseBodyDataEnvironmentsDependencies `json:"Dependencies,omitempty" xml:"Dependencies,omitempty" type:"Struct"`
	// The description of the environment.
	//
	// example:
	//
	// The MySQL service is deployed in a Kubernetes cluster.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the feature is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The tag of the environment.
	//
	// example:
	//
	// Container
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The name of the environment.
	//
	// example:
	//
	// CS
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The control policies in the environment.
	Policies *ListAddonsResponseBodyDataEnvironmentsPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
}

func (s ListAddonsResponseBodyDataEnvironments) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBodyDataEnvironments) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBodyDataEnvironments) GetDependencies() *ListAddonsResponseBodyDataEnvironmentsDependencies {
	return s.Dependencies
}

func (s *ListAddonsResponseBodyDataEnvironments) GetDescription() *string {
	return s.Description
}

func (s *ListAddonsResponseBodyDataEnvironments) GetEnable() *bool {
	return s.Enable
}

func (s *ListAddonsResponseBodyDataEnvironments) GetLabel() *string {
	return s.Label
}

func (s *ListAddonsResponseBodyDataEnvironments) GetName() *string {
	return s.Name
}

func (s *ListAddonsResponseBodyDataEnvironments) GetPolicies() *ListAddonsResponseBodyDataEnvironmentsPolicies {
	return s.Policies
}

func (s *ListAddonsResponseBodyDataEnvironments) SetDependencies(v *ListAddonsResponseBodyDataEnvironmentsDependencies) *ListAddonsResponseBodyDataEnvironments {
	s.Dependencies = v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironments) SetDescription(v string) *ListAddonsResponseBodyDataEnvironments {
	s.Description = &v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironments) SetEnable(v bool) *ListAddonsResponseBodyDataEnvironments {
	s.Enable = &v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironments) SetLabel(v string) *ListAddonsResponseBodyDataEnvironments {
	s.Label = &v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironments) SetName(v string) *ListAddonsResponseBodyDataEnvironments {
	s.Name = &v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironments) SetPolicies(v *ListAddonsResponseBodyDataEnvironmentsPolicies) *ListAddonsResponseBodyDataEnvironments {
	s.Policies = v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironments) Validate() error {
	return dara.Validate(s)
}

type ListAddonsResponseBodyDataEnvironmentsDependencies struct {
	// The supported cluster types.
	ClusterTypes []*string `json:"ClusterTypes,omitempty" xml:"ClusterTypes,omitempty" type:"Repeated"`
	// The feature on which the environment depends.
	Features map[string]*bool `json:"Features,omitempty" xml:"Features,omitempty"`
	// The services.
	Services []*string `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
}

func (s ListAddonsResponseBodyDataEnvironmentsDependencies) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBodyDataEnvironmentsDependencies) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBodyDataEnvironmentsDependencies) GetClusterTypes() []*string {
	return s.ClusterTypes
}

func (s *ListAddonsResponseBodyDataEnvironmentsDependencies) GetFeatures() map[string]*bool {
	return s.Features
}

func (s *ListAddonsResponseBodyDataEnvironmentsDependencies) GetServices() []*string {
	return s.Services
}

func (s *ListAddonsResponseBodyDataEnvironmentsDependencies) SetClusterTypes(v []*string) *ListAddonsResponseBodyDataEnvironmentsDependencies {
	s.ClusterTypes = v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironmentsDependencies) SetFeatures(v map[string]*bool) *ListAddonsResponseBodyDataEnvironmentsDependencies {
	s.Features = v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironmentsDependencies) SetServices(v []*string) *ListAddonsResponseBodyDataEnvironmentsDependencies {
	s.Services = v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironmentsDependencies) Validate() error {
	return dara.Validate(s)
}

type ListAddonsResponseBodyDataEnvironmentsPolicies struct {
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
	// false.
	EnableServiceAccount *bool `json:"EnableServiceAccount,omitempty" xml:"EnableServiceAccount,omitempty"`
	// The metric check rule.
	MetricCheckRule *ListAddonsResponseBodyDataEnvironmentsPoliciesMetricCheckRule `json:"MetricCheckRule,omitempty" xml:"MetricCheckRule,omitempty" type:"Struct"`
	// Indicates whether a restart is required after the installation.
	//
	// example:
	//
	// false
	NeedRestartAfterIntegration *bool `json:"NeedRestartAfterIntegration,omitempty" xml:"NeedRestartAfterIntegration,omitempty"`
	// The supported protocols.
	Protocols []*ListAddonsResponseBodyDataEnvironmentsPoliciesProtocols `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	// The target name of the add-on.
	//
	// example:
	//
	// cloud-rds-mysql
	TargetAddonName *string `json:"TargetAddonName,omitempty" xml:"TargetAddonName,omitempty"`
}

func (s ListAddonsResponseBodyDataEnvironmentsPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBodyDataEnvironmentsPolicies) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBodyDataEnvironmentsPolicies) GetAlertDefaultStatus() *string {
	return s.AlertDefaultStatus
}

func (s *ListAddonsResponseBodyDataEnvironmentsPolicies) GetDefaultInstall() *bool {
	return s.DefaultInstall
}

func (s *ListAddonsResponseBodyDataEnvironmentsPolicies) GetEnableServiceAccount() *bool {
	return s.EnableServiceAccount
}

func (s *ListAddonsResponseBodyDataEnvironmentsPolicies) GetMetricCheckRule() *ListAddonsResponseBodyDataEnvironmentsPoliciesMetricCheckRule {
	return s.MetricCheckRule
}

func (s *ListAddonsResponseBodyDataEnvironmentsPolicies) GetNeedRestartAfterIntegration() *bool {
	return s.NeedRestartAfterIntegration
}

func (s *ListAddonsResponseBodyDataEnvironmentsPolicies) GetProtocols() []*ListAddonsResponseBodyDataEnvironmentsPoliciesProtocols {
	return s.Protocols
}

func (s *ListAddonsResponseBodyDataEnvironmentsPolicies) GetTargetAddonName() *string {
	return s.TargetAddonName
}

func (s *ListAddonsResponseBodyDataEnvironmentsPolicies) SetAlertDefaultStatus(v string) *ListAddonsResponseBodyDataEnvironmentsPolicies {
	s.AlertDefaultStatus = &v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironmentsPolicies) SetDefaultInstall(v bool) *ListAddonsResponseBodyDataEnvironmentsPolicies {
	s.DefaultInstall = &v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironmentsPolicies) SetEnableServiceAccount(v bool) *ListAddonsResponseBodyDataEnvironmentsPolicies {
	s.EnableServiceAccount = &v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironmentsPolicies) SetMetricCheckRule(v *ListAddonsResponseBodyDataEnvironmentsPoliciesMetricCheckRule) *ListAddonsResponseBodyDataEnvironmentsPolicies {
	s.MetricCheckRule = v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironmentsPolicies) SetNeedRestartAfterIntegration(v bool) *ListAddonsResponseBodyDataEnvironmentsPolicies {
	s.NeedRestartAfterIntegration = &v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironmentsPolicies) SetProtocols(v []*ListAddonsResponseBodyDataEnvironmentsPoliciesProtocols) *ListAddonsResponseBodyDataEnvironmentsPolicies {
	s.Protocols = v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironmentsPolicies) SetTargetAddonName(v string) *ListAddonsResponseBodyDataEnvironmentsPolicies {
	s.TargetAddonName = &v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironmentsPolicies) Validate() error {
	return dara.Validate(s)
}

type ListAddonsResponseBodyDataEnvironmentsPoliciesMetricCheckRule struct {
	// The PromQL statements.
	PromQL []*string `json:"PromQL,omitempty" xml:"PromQL,omitempty" type:"Repeated"`
}

func (s ListAddonsResponseBodyDataEnvironmentsPoliciesMetricCheckRule) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBodyDataEnvironmentsPoliciesMetricCheckRule) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBodyDataEnvironmentsPoliciesMetricCheckRule) GetPromQL() []*string {
	return s.PromQL
}

func (s *ListAddonsResponseBodyDataEnvironmentsPoliciesMetricCheckRule) SetPromQL(v []*string) *ListAddonsResponseBodyDataEnvironmentsPoliciesMetricCheckRule {
	s.PromQL = v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironmentsPoliciesMetricCheckRule) Validate() error {
	return dara.Validate(s)
}

type ListAddonsResponseBodyDataEnvironmentsPoliciesProtocols struct {
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

func (s ListAddonsResponseBodyDataEnvironmentsPoliciesProtocols) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBodyDataEnvironmentsPoliciesProtocols) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBodyDataEnvironmentsPoliciesProtocols) GetDescription() *string {
	return s.Description
}

func (s *ListAddonsResponseBodyDataEnvironmentsPoliciesProtocols) GetIcon() *string {
	return s.Icon
}

func (s *ListAddonsResponseBodyDataEnvironmentsPoliciesProtocols) GetLabel() *string {
	return s.Label
}

func (s *ListAddonsResponseBodyDataEnvironmentsPoliciesProtocols) GetName() *string {
	return s.Name
}

func (s *ListAddonsResponseBodyDataEnvironmentsPoliciesProtocols) SetDescription(v string) *ListAddonsResponseBodyDataEnvironmentsPoliciesProtocols {
	s.Description = &v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironmentsPoliciesProtocols) SetIcon(v string) *ListAddonsResponseBodyDataEnvironmentsPoliciesProtocols {
	s.Icon = &v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironmentsPoliciesProtocols) SetLabel(v string) *ListAddonsResponseBodyDataEnvironmentsPoliciesProtocols {
	s.Label = &v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironmentsPoliciesProtocols) SetName(v string) *ListAddonsResponseBodyDataEnvironmentsPoliciesProtocols {
	s.Name = &v
	return s
}

func (s *ListAddonsResponseBodyDataEnvironmentsPoliciesProtocols) Validate() error {
	return dara.Validate(s)
}
