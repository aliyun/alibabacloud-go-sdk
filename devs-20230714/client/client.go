// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Checkout struct {
	// example:
	//
	// +001691d0768ca49e9550beeb59fbc163f33b7e88:refs/remotes/origin/master
	Ref *string `json:"ref,omitempty" xml:"ref,omitempty"`
	// example:
	//
	// https:/your_token/@github.com/buptwzj/test-initRepo4.git
	Remote *string `json:"remote,omitempty" xml:"remote,omitempty"`
}

func (s Checkout) String() string {
	return tea.Prettify(s)
}

func (s Checkout) GoString() string {
	return s.String()
}

func (s *Checkout) SetRef(v string) *Checkout {
	s.Ref = &v
	return s
}

func (s *Checkout) SetRemote(v string) *Checkout {
	s.Remote = &v
	return s
}

type CodeupEventPayload struct {
	OriginalPayload []byte `json:"originalPayload,omitempty" xml:"originalPayload,omitempty"`
}

func (s CodeupEventPayload) String() string {
	return tea.Prettify(s)
}

func (s CodeupEventPayload) GoString() string {
	return s.String()
}

func (s *CodeupEventPayload) SetOriginalPayload(v []byte) *CodeupEventPayload {
	s.OriginalPayload = v
	return s
}

type Condition struct {
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
}

func (s Condition) String() string {
	return tea.Prettify(s)
}

func (s Condition) GoString() string {
	return s.String()
}

func (s *Condition) SetExpression(v string) *Condition {
	s.Expression = &v
	return s
}

type Connection struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// test-description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// example:
	//
	// Connection
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-connection
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	ResourceVersion *int32 `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
	// This parameter is required.
	Spec   *ConnectionSpec   `json:"spec,omitempty" xml:"spec,omitempty"`
	Status *ConnectionStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s Connection) String() string {
	return tea.Prettify(s)
}

func (s Connection) GoString() string {
	return s.String()
}

func (s *Connection) SetCreatedTime(v string) *Connection {
	s.CreatedTime = &v
	return s
}

func (s *Connection) SetDescription(v string) *Connection {
	s.Description = &v
	return s
}

func (s *Connection) SetGeneration(v int32) *Connection {
	s.Generation = &v
	return s
}

func (s *Connection) SetKind(v string) *Connection {
	s.Kind = &v
	return s
}

func (s *Connection) SetLabels(v map[string]*string) *Connection {
	s.Labels = v
	return s
}

func (s *Connection) SetName(v string) *Connection {
	s.Name = &v
	return s
}

func (s *Connection) SetResourceVersion(v int32) *Connection {
	s.ResourceVersion = &v
	return s
}

func (s *Connection) SetSpec(v *ConnectionSpec) *Connection {
	s.Spec = v
	return s
}

func (s *Connection) SetStatus(v *ConnectionStatus) *Connection {
	s.Status = v
	return s
}

func (s *Connection) SetUid(v string) *Connection {
	s.Uid = &v
	return s
}

type ConnectionSpec struct {
	Account      *GitAccount   `json:"account,omitempty" xml:"account,omitempty"`
	GitlabConfig *GitLabConfig `json:"gitlabConfig,omitempty" xml:"gitlabConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// github
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
}

func (s ConnectionSpec) String() string {
	return tea.Prettify(s)
}

func (s ConnectionSpec) GoString() string {
	return s.String()
}

func (s *ConnectionSpec) SetAccount(v *GitAccount) *ConnectionSpec {
	s.Account = v
	return s
}

func (s *ConnectionSpec) SetGitlabConfig(v *GitLabConfig) *ConnectionSpec {
	s.GitlabConfig = v
	return s
}

func (s *ConnectionSpec) SetPlatform(v string) *ConnectionSpec {
	s.Platform = &v
	return s
}

type ConnectionStatus struct {
	Installation *Installation `json:"installation,omitempty" xml:"installation,omitempty"`
}

func (s ConnectionStatus) String() string {
	return tea.Prettify(s)
}

func (s ConnectionStatus) GoString() string {
	return s.String()
}

func (s *ConnectionStatus) SetInstallation(v *Installation) *ConnectionStatus {
	s.Installation = v
	return s
}

type Context struct {
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func (s Context) String() string {
	return tea.Prettify(s)
}

func (s Context) GoString() string {
	return s.String()
}

func (s *Context) SetData(v map[string]interface{}) *Context {
	s.Data = v
	return s
}

type ContextSchema struct {
	// example:
	//
	// [git](https://git-scm.com/) address for [git clone](https://git-scm.com/docs/git-clone).
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// git@gitlab.alibaba-inc.com:serverless/lambda.git
	Hint *string `json:"hint,omitempty" xml:"hint,omitempty"`
	// example:
	//
	// gitRepoUrl
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// true
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ContextSchema) String() string {
	return tea.Prettify(s)
}

func (s ContextSchema) GoString() string {
	return s.String()
}

func (s *ContextSchema) SetDescription(v string) *ContextSchema {
	s.Description = &v
	return s
}

func (s *ContextSchema) SetHint(v string) *ContextSchema {
	s.Hint = &v
	return s
}

func (s *ContextSchema) SetName(v string) *ContextSchema {
	s.Name = &v
	return s
}

func (s *ContextSchema) SetRequired(v bool) *ContextSchema {
	s.Required = &v
	return s
}

func (s *ContextSchema) SetType(v string) *ContextSchema {
	s.Type = &v
	return s
}

type Environment struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// test env
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// example:
	//
	// Environment
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// demo-env
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// demo-project
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// example:
	//
	// 1
	ResourceVersion *int32 `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
	// This parameter is required.
	Spec   *EnvironmentSpec   `json:"spec,omitempty" xml:"spec,omitempty"`
	Status *EnvironmentStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s Environment) String() string {
	return tea.Prettify(s)
}

func (s Environment) GoString() string {
	return s.String()
}

func (s *Environment) SetCreatedTime(v string) *Environment {
	s.CreatedTime = &v
	return s
}

func (s *Environment) SetDescription(v string) *Environment {
	s.Description = &v
	return s
}

func (s *Environment) SetGeneration(v int32) *Environment {
	s.Generation = &v
	return s
}

func (s *Environment) SetKind(v string) *Environment {
	s.Kind = &v
	return s
}

func (s *Environment) SetLabels(v map[string]*string) *Environment {
	s.Labels = v
	return s
}

func (s *Environment) SetName(v string) *Environment {
	s.Name = &v
	return s
}

func (s *Environment) SetProjectName(v string) *Environment {
	s.ProjectName = &v
	return s
}

func (s *Environment) SetResourceVersion(v int32) *Environment {
	s.ResourceVersion = &v
	return s
}

func (s *Environment) SetSpec(v *EnvironmentSpec) *Environment {
	s.Spec = v
	return s
}

func (s *Environment) SetStatus(v *EnvironmentStatus) *Environment {
	s.Status = v
	return s
}

func (s *Environment) SetUid(v string) *Environment {
	s.Uid = &v
	return s
}

type EnvironmentSpec struct {
	// example:
	//
	// dev
	Alias            *string           `json:"alias,omitempty" xml:"alias,omitempty"`
	InfraStackConfig *InfraStackSpec   `json:"infraStackConfig,omitempty" xml:"infraStackConfig,omitempty"`
	IsAutoDeploy     *bool             `json:"isAutoDeploy,omitempty" xml:"isAutoDeploy,omitempty"`
	RepositoryConfig *RepositoryConfig `json:"repositoryConfig,omitempty" xml:"repositoryConfig,omitempty"`
	// example:
	//
	// acs:ram::*******:role/aliyundevsdefaultrole
	RoleArn        *string                        `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	ServiceOverlay *EnvironmentSpecServiceOverlay `json:"serviceOverlay,omitempty" xml:"serviceOverlay,omitempty" type:"Struct"`
	TemplateConfig *TemplateConfig                `json:"templateConfig,omitempty" xml:"templateConfig,omitempty"`
	// example:
	//
	// Testing
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s EnvironmentSpec) String() string {
	return tea.Prettify(s)
}

func (s EnvironmentSpec) GoString() string {
	return s.String()
}

func (s *EnvironmentSpec) SetAlias(v string) *EnvironmentSpec {
	s.Alias = &v
	return s
}

func (s *EnvironmentSpec) SetInfraStackConfig(v *InfraStackSpec) *EnvironmentSpec {
	s.InfraStackConfig = v
	return s
}

func (s *EnvironmentSpec) SetIsAutoDeploy(v bool) *EnvironmentSpec {
	s.IsAutoDeploy = &v
	return s
}

func (s *EnvironmentSpec) SetRepositoryConfig(v *RepositoryConfig) *EnvironmentSpec {
	s.RepositoryConfig = v
	return s
}

func (s *EnvironmentSpec) SetRoleArn(v string) *EnvironmentSpec {
	s.RoleArn = &v
	return s
}

func (s *EnvironmentSpec) SetServiceOverlay(v *EnvironmentSpecServiceOverlay) *EnvironmentSpec {
	s.ServiceOverlay = v
	return s
}

func (s *EnvironmentSpec) SetTemplateConfig(v *TemplateConfig) *EnvironmentSpec {
	s.TemplateConfig = v
	return s
}

func (s *EnvironmentSpec) SetType(v string) *EnvironmentSpec {
	s.Type = &v
	return s
}

type EnvironmentSpecServiceOverlay struct {
	// example:
	//
	// {"fc3":{"timeout":600}}
	Components map[string]interface{} `json:"components,omitempty" xml:"components,omitempty"`
	// example:
	//
	// {"dummyFunction":{"timeout":600}}
	Resources map[string]interface{} `json:"resources,omitempty" xml:"resources,omitempty"`
}

func (s EnvironmentSpecServiceOverlay) String() string {
	return tea.Prettify(s)
}

func (s EnvironmentSpecServiceOverlay) GoString() string {
	return s.String()
}

func (s *EnvironmentSpecServiceOverlay) SetComponents(v map[string]interface{}) *EnvironmentSpecServiceOverlay {
	s.Components = v
	return s
}

func (s *EnvironmentSpecServiceOverlay) SetResources(v map[string]interface{}) *EnvironmentSpecServiceOverlay {
	s.Resources = v
	return s
}

type EnvironmentStatus struct {
	InfraStackStatus    *InfraStackStatus `json:"infraStackStatus,omitempty" xml:"infraStackStatus,omitempty"`
	LatestReleaseDetail *ReleaseDetail    `json:"latestReleaseDetail,omitempty" xml:"latestReleaseDetail,omitempty"`
	// example:
	//
	// 1
	ObservedGeneration *int64 `json:"observedGeneration,omitempty" xml:"observedGeneration,omitempty"`
	// example:
	//
	// 2021-11-19T09:34:38Z
	ObservedTime *string `json:"observedTime,omitempty" xml:"observedTime,omitempty"`
}

func (s EnvironmentStatus) String() string {
	return tea.Prettify(s)
}

func (s EnvironmentStatus) GoString() string {
	return s.String()
}

func (s *EnvironmentStatus) SetInfraStackStatus(v *InfraStackStatus) *EnvironmentStatus {
	s.InfraStackStatus = v
	return s
}

func (s *EnvironmentStatus) SetLatestReleaseDetail(v *ReleaseDetail) *EnvironmentStatus {
	s.LatestReleaseDetail = v
	return s
}

func (s *EnvironmentStatus) SetObservedGeneration(v int64) *EnvironmentStatus {
	s.ObservedGeneration = &v
	return s
}

func (s *EnvironmentStatus) SetObservedTime(v string) *EnvironmentStatus {
	s.ObservedTime = &v
	return s
}

type EventFilterConfig struct {
	PullRequest *PullRequestFilter `json:"pullRequest,omitempty" xml:"pullRequest,omitempty"`
	Push        *PushFilter        `json:"push,omitempty" xml:"push,omitempty"`
}

func (s EventFilterConfig) String() string {
	return tea.Prettify(s)
}

func (s EventFilterConfig) GoString() string {
	return s.String()
}

func (s *EventFilterConfig) SetPullRequest(v *PullRequestFilter) *EventFilterConfig {
	s.PullRequest = v
	return s
}

func (s *EventFilterConfig) SetPush(v *PushFilter) *EventFilterConfig {
	s.Push = v
	return s
}

type EventPayload struct {
	Codeup *CodeupEventPayload `json:"codeup,omitempty" xml:"codeup,omitempty"`
	Gitee  *GiteeEventPayload  `json:"gitee,omitempty" xml:"gitee,omitempty"`
	Github *GithubEventPayload `json:"github,omitempty" xml:"github,omitempty"`
	Gitlab *GitlabEventPayload `json:"gitlab,omitempty" xml:"gitlab,omitempty"`
	Manual *ManualEventPayload `json:"manual,omitempty" xml:"manual,omitempty"`
}

func (s EventPayload) String() string {
	return tea.Prettify(s)
}

func (s EventPayload) GoString() string {
	return s.String()
}

func (s *EventPayload) SetCodeup(v *CodeupEventPayload) *EventPayload {
	s.Codeup = v
	return s
}

func (s *EventPayload) SetGitee(v *GiteeEventPayload) *EventPayload {
	s.Gitee = v
	return s
}

func (s *EventPayload) SetGithub(v *GithubEventPayload) *EventPayload {
	s.Github = v
	return s
}

func (s *EventPayload) SetGitlab(v *GitlabEventPayload) *EventPayload {
	s.Gitlab = v
	return s
}

func (s *EventPayload) SetManual(v *ManualEventPayload) *EventPayload {
	s.Manual = v
	return s
}

type GitAccount struct {
	// example:
	//
	// https://gitee.com/assets/no_portrait.png
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// example:
	//
	// your_displayname
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// 1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// your_username
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// https://gitlab.com
	Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s GitAccount) String() string {
	return tea.Prettify(s)
}

func (s GitAccount) GoString() string {
	return s.String()
}

func (s *GitAccount) SetAvatar(v string) *GitAccount {
	s.Avatar = &v
	return s
}

func (s *GitAccount) SetDisplayName(v string) *GitAccount {
	s.DisplayName = &v
	return s
}

func (s *GitAccount) SetId(v string) *GitAccount {
	s.Id = &v
	return s
}

func (s *GitAccount) SetName(v string) *GitAccount {
	s.Name = &v
	return s
}

func (s *GitAccount) SetUri(v string) *GitAccount {
	s.Uri = &v
	return s
}

type GitEventSnapshot struct {
	// example:
	//
	// main
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// 12721ec262d03a93809ba2bbc717963cb298ceca
	CommitID *string `json:"commitID,omitempty" xml:"commitID,omitempty"`
	// example:
	//
	// 1.0
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s GitEventSnapshot) String() string {
	return tea.Prettify(s)
}

func (s GitEventSnapshot) GoString() string {
	return s.String()
}

func (s *GitEventSnapshot) SetBranch(v string) *GitEventSnapshot {
	s.Branch = &v
	return s
}

func (s *GitEventSnapshot) SetCommitID(v string) *GitEventSnapshot {
	s.CommitID = &v
	return s
}

func (s *GitEventSnapshot) SetTag(v string) *GitEventSnapshot {
	s.Tag = &v
	return s
}

type GitLabConfig struct {
	// example:
	//
	// your-token
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// example:
	//
	// http://gitlab.c16194660f14898a0810408171302ac.cn-shanghai.alicontainer.com/
	Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s GitLabConfig) String() string {
	return tea.Prettify(s)
}

func (s GitLabConfig) GoString() string {
	return s.String()
}

func (s *GitLabConfig) SetToken(v string) *GitLabConfig {
	s.Token = &v
	return s
}

func (s *GitLabConfig) SetUri(v string) *GitLabConfig {
	s.Uri = &v
	return s
}

type GiteeEventPayload struct {
	OriginalPayload []byte `json:"originalPayload,omitempty" xml:"originalPayload,omitempty"`
}

func (s GiteeEventPayload) String() string {
	return tea.Prettify(s)
}

func (s GiteeEventPayload) GoString() string {
	return s.String()
}

func (s *GiteeEventPayload) SetOriginalPayload(v []byte) *GiteeEventPayload {
	s.OriginalPayload = v
	return s
}

type GithubEventPayload struct {
	OriginalPayload []byte `json:"originalPayload,omitempty" xml:"originalPayload,omitempty"`
}

func (s GithubEventPayload) String() string {
	return tea.Prettify(s)
}

func (s GithubEventPayload) GoString() string {
	return s.String()
}

func (s *GithubEventPayload) SetOriginalPayload(v []byte) *GithubEventPayload {
	s.OriginalPayload = v
	return s
}

type GitlabEventPayload struct {
	OriginalPayload []byte `json:"originalPayload,omitempty" xml:"originalPayload,omitempty"`
}

func (s GitlabEventPayload) String() string {
	return tea.Prettify(s)
}

func (s GitlabEventPayload) GoString() string {
	return s.String()
}

func (s *GitlabEventPayload) SetOriginalPayload(v []byte) *GitlabEventPayload {
	s.OriginalPayload = v
	return s
}

type InfraStackResourceState struct {
	ResourceDrifts []*ResourceDrift  `json:"resourceDrifts,omitempty" xml:"resourceDrifts,omitempty" type:"Repeated"`
	Resources      []*ResourceDetail `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
}

func (s InfraStackResourceState) String() string {
	return tea.Prettify(s)
}

func (s InfraStackResourceState) GoString() string {
	return s.String()
}

func (s *InfraStackResourceState) SetResourceDrifts(v []*ResourceDrift) *InfraStackResourceState {
	s.ResourceDrifts = v
	return s
}

func (s *InfraStackResourceState) SetResources(v []*ResourceDetail) *InfraStackResourceState {
	s.Resources = v
	return s
}

type InfraStackSpec struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionID *string `json:"regionID,omitempty" xml:"regionID,omitempty"`
	// example:
	//
	// acs:ram::1234567890:role/devs-role
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// example:
	//
	// serverless.devs.com/fc-builtin
	TemplateName *string                     `json:"templateName,omitempty" xml:"templateName,omitempty"`
	TemplateSpec *InfraStackSpecTemplateSpec `json:"templateSpec,omitempty" xml:"templateSpec,omitempty" type:"Struct"`
	// example:
	//
	// {"testKey":"testValue"}
	TemplateVariables map[string]interface{} `json:"templateVariables,omitempty" xml:"templateVariables,omitempty"`
}

func (s InfraStackSpec) String() string {
	return tea.Prettify(s)
}

func (s InfraStackSpec) GoString() string {
	return s.String()
}

func (s *InfraStackSpec) SetRegionID(v string) *InfraStackSpec {
	s.RegionID = &v
	return s
}

func (s *InfraStackSpec) SetRoleArn(v string) *InfraStackSpec {
	s.RoleArn = &v
	return s
}

func (s *InfraStackSpec) SetTemplateName(v string) *InfraStackSpec {
	s.TemplateName = &v
	return s
}

func (s *InfraStackSpec) SetTemplateSpec(v *InfraStackSpecTemplateSpec) *InfraStackSpec {
	s.TemplateSpec = v
	return s
}

func (s *InfraStackSpec) SetTemplateVariables(v map[string]interface{}) *InfraStackSpec {
	s.TemplateVariables = v
	return s
}

type InfraStackSpecTemplateSpec struct {
	// example:
	//
	// "\nresource \"alicloud_fc_service\" \"default\" {\n  name        = \"xiliu-tf-test-srv4\"\n  description = \"xiliu tf  test service \"\n}\n\n\nresource \"alicloud_oss_bucket\" \"default\" {\n  bucket = \"xiliu-test-tf-bucket4\"\n  acl    = \"private\"\n}\n\noutput \"service_name\" {\n  value = alicloud_fc_service.default.name\n}\n\noutput \"oss_bucket\" {\n  value = alicloud_oss_bucket.default.bucket\n}",
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// terrafrom
	Engine *string `json:"engine,omitempty" xml:"engine,omitempty"`
}

func (s InfraStackSpecTemplateSpec) String() string {
	return tea.Prettify(s)
}

func (s InfraStackSpecTemplateSpec) GoString() string {
	return s.String()
}

func (s *InfraStackSpecTemplateSpec) SetContent(v string) *InfraStackSpecTemplateSpec {
	s.Content = &v
	return s
}

func (s *InfraStackSpecTemplateSpec) SetEngine(v string) *InfraStackSpecTemplateSpec {
	s.Engine = &v
	return s
}

type InfraStackStatus struct {
	// example:
	//
	// Success!
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 0
	ObservedGeneration *int32 `json:"observedGeneration,omitempty" xml:"observedGeneration,omitempty"`
	// example:
	//
	// 2021-10-08T23:14:16Z
	ObservedTime *string `json:"observedTime,omitempty" xml:"observedTime,omitempty"`
	// example:
	//
	// DeploySuccess
	Phase         *string                  `json:"phase,omitempty" xml:"phase,omitempty"`
	ResourceState *InfraStackResourceState `json:"resourceState,omitempty" xml:"resourceState,omitempty"`
	// example:
	//
	// {"vpcId":"vpc-xxx"}
	TemplateOutputs map[string]interface{}          `json:"templateOutputs,omitempty" xml:"templateOutputs,omitempty"`
	TemplateStatus  *InfraStackStatusTemplateStatus `json:"templateStatus,omitempty" xml:"templateStatus,omitempty" type:"Struct"`
}

func (s InfraStackStatus) String() string {
	return tea.Prettify(s)
}

func (s InfraStackStatus) GoString() string {
	return s.String()
}

func (s *InfraStackStatus) SetMessage(v string) *InfraStackStatus {
	s.Message = &v
	return s
}

func (s *InfraStackStatus) SetObservedGeneration(v int32) *InfraStackStatus {
	s.ObservedGeneration = &v
	return s
}

func (s *InfraStackStatus) SetObservedTime(v string) *InfraStackStatus {
	s.ObservedTime = &v
	return s
}

func (s *InfraStackStatus) SetPhase(v string) *InfraStackStatus {
	s.Phase = &v
	return s
}

func (s *InfraStackStatus) SetResourceState(v *InfraStackResourceState) *InfraStackStatus {
	s.ResourceState = v
	return s
}

func (s *InfraStackStatus) SetTemplateOutputs(v map[string]interface{}) *InfraStackStatus {
	s.TemplateOutputs = v
	return s
}

func (s *InfraStackStatus) SetTemplateStatus(v *InfraStackStatusTemplateStatus) *InfraStackStatus {
	s.TemplateStatus = v
	return s
}

type InfraStackStatusTemplateStatus struct {
	// This parameter is required.
	Outputs []*TerraformOutputValue `json:"outputs,omitempty" xml:"outputs,omitempty" type:"Repeated"`
	// This parameter is required.
	Variables []*TerraformInputVariable `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s InfraStackStatusTemplateStatus) String() string {
	return tea.Prettify(s)
}

func (s InfraStackStatusTemplateStatus) GoString() string {
	return s.String()
}

func (s *InfraStackStatusTemplateStatus) SetOutputs(v []*TerraformOutputValue) *InfraStackStatusTemplateStatus {
	s.Outputs = v
	return s
}

func (s *InfraStackStatusTemplateStatus) SetVariables(v []*TerraformInputVariable) *InfraStackStatusTemplateStatus {
	s.Variables = v
	return s
}

type Installation struct {
	// example:
	//
	// https://github.com/login/oauth/authorize?client_id=86059a1b2bb20d3e5fc3&scope=repo,repo:status,delete_repo
	ActionUri *string `json:"actionUri,omitempty" xml:"actionUri,omitempty"`
	// example:
	//
	// Please click \"actionUri\" to complete the OAuth authorization process
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// finished
	Stage *string `json:"stage,omitempty" xml:"stage,omitempty"`
}

func (s Installation) String() string {
	return tea.Prettify(s)
}

func (s Installation) GoString() string {
	return s.String()
}

func (s *Installation) SetActionUri(v string) *Installation {
	s.ActionUri = &v
	return s
}

func (s *Installation) SetMessage(v string) *Installation {
	s.Message = &v
	return s
}

func (s *Installation) SetStage(v string) *Installation {
	s.Stage = &v
	return s
}

type ManualEventPayload struct {
	Branch         *string         `json:"branch,omitempty" xml:"branch,omitempty"`
	CommitID       *string         `json:"commitID,omitempty" xml:"commitID,omitempty"`
	Tag            *string         `json:"tag,omitempty" xml:"tag,omitempty"`
	TemplateConfig *TemplateConfig `json:"templateConfig,omitempty" xml:"templateConfig,omitempty"`
}

func (s ManualEventPayload) String() string {
	return tea.Prettify(s)
}

func (s ManualEventPayload) GoString() string {
	return s.String()
}

func (s *ManualEventPayload) SetBranch(v string) *ManualEventPayload {
	s.Branch = &v
	return s
}

func (s *ManualEventPayload) SetCommitID(v string) *ManualEventPayload {
	s.CommitID = &v
	return s
}

func (s *ManualEventPayload) SetTag(v string) *ManualEventPayload {
	s.Tag = &v
	return s
}

func (s *ManualEventPayload) SetTemplateConfig(v *TemplateConfig) *ManualEventPayload {
	s.TemplateConfig = v
	return s
}

type OAuthCredential struct {
	// This parameter is required.
	//
	// example:
	//
	// 1716176924603
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1716263324603
	Expiration *int64 `json:"expiration,omitempty" xml:"expiration,omitempty"`
	// example:
	//
	// 4d77bfae284770d94ebeed6b0199ebfd65e3943ba4f1e44dc36d792a93ba0d13
	RefreshToken *string `json:"refreshToken,omitempty" xml:"refreshToken,omitempty"`
	// example:
	//
	// user_info projects pull_requests hook gists emails
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4e84246b6b3962cd3d207aad1ea2f911
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// example:
	//
	// bearer
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s OAuthCredential) String() string {
	return tea.Prettify(s)
}

func (s OAuthCredential) GoString() string {
	return s.String()
}

func (s *OAuthCredential) SetCreatedTime(v int64) *OAuthCredential {
	s.CreatedTime = &v
	return s
}

func (s *OAuthCredential) SetExpiration(v int64) *OAuthCredential {
	s.Expiration = &v
	return s
}

func (s *OAuthCredential) SetRefreshToken(v string) *OAuthCredential {
	s.RefreshToken = &v
	return s
}

func (s *OAuthCredential) SetScope(v string) *OAuthCredential {
	s.Scope = &v
	return s
}

func (s *OAuthCredential) SetToken(v string) *OAuthCredential {
	s.Token = &v
	return s
}

func (s *OAuthCredential) SetType(v string) *OAuthCredential {
	s.Type = &v
	return s
}

type Pipeline struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// 2021-11-19T09:34:38Z
	DeletionTime *string `json:"deletionTime,omitempty" xml:"deletionTime,omitempty"`
	// example:
	//
	// Pipeline example.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// example:
	//
	// Pipeline
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-pipeline
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	ResourceVersion *int32          `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
	Spec            *PipelineSpec   `json:"spec,omitempty" xml:"spec,omitempty"`
	Status          *PipelineStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s Pipeline) String() string {
	return tea.Prettify(s)
}

func (s Pipeline) GoString() string {
	return s.String()
}

func (s *Pipeline) SetCreatedTime(v string) *Pipeline {
	s.CreatedTime = &v
	return s
}

func (s *Pipeline) SetDeletionTime(v string) *Pipeline {
	s.DeletionTime = &v
	return s
}

func (s *Pipeline) SetDescription(v string) *Pipeline {
	s.Description = &v
	return s
}

func (s *Pipeline) SetGeneration(v int32) *Pipeline {
	s.Generation = &v
	return s
}

func (s *Pipeline) SetKind(v string) *Pipeline {
	s.Kind = &v
	return s
}

func (s *Pipeline) SetLabels(v map[string]*string) *Pipeline {
	s.Labels = v
	return s
}

func (s *Pipeline) SetName(v string) *Pipeline {
	s.Name = &v
	return s
}

func (s *Pipeline) SetResourceVersion(v int32) *Pipeline {
	s.ResourceVersion = &v
	return s
}

func (s *Pipeline) SetSpec(v *PipelineSpec) *Pipeline {
	s.Spec = v
	return s
}

func (s *Pipeline) SetStatus(v *PipelineStatus) *Pipeline {
	s.Status = v
	return s
}

func (s *Pipeline) SetUid(v string) *Pipeline {
	s.Uid = &v
	return s
}

type PipelineSpec struct {
	Context *Context `json:"context,omitempty" xml:"context,omitempty"`
	// example:
	//
	// my-pipeline-template
	TemplateName *string               `json:"templateName,omitempty" xml:"templateName,omitempty"`
	TemplateSpec *PipelineTemplateSpec `json:"templateSpec,omitempty" xml:"templateSpec,omitempty"`
}

func (s PipelineSpec) String() string {
	return tea.Prettify(s)
}

func (s PipelineSpec) GoString() string {
	return s.String()
}

func (s *PipelineSpec) SetContext(v *Context) *PipelineSpec {
	s.Context = v
	return s
}

func (s *PipelineSpec) SetTemplateName(v string) *PipelineSpec {
	s.TemplateName = &v
	return s
}

func (s *PipelineSpec) SetTemplateSpec(v *PipelineTemplateSpec) *PipelineSpec {
	s.TemplateSpec = v
	return s
}

type PipelineStatus struct {
	LatestExecError *TaskExecError `json:"latestExecError,omitempty" xml:"latestExecError,omitempty"`
	// example:
	//
	// Success
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
}

func (s PipelineStatus) String() string {
	return tea.Prettify(s)
}

func (s PipelineStatus) GoString() string {
	return s.String()
}

func (s *PipelineStatus) SetLatestExecError(v *TaskExecError) *PipelineStatus {
	s.LatestExecError = v
	return s
}

func (s *PipelineStatus) SetPhase(v string) *PipelineStatus {
	s.Phase = &v
	return s
}

type PipelineTemplate struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// 2021-11-19T09:34:38Z
	DeletionTime *string `json:"deletionTime,omitempty" xml:"deletionTime,omitempty"`
	// example:
	//
	// PipelineTemplate example.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// example:
	//
	// PipelineTemplate
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-pipeline-template
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	ResourceVersion *int32                `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
	Spec            *PipelineTemplateSpec `json:"spec,omitempty" xml:"spec,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s PipelineTemplate) String() string {
	return tea.Prettify(s)
}

func (s PipelineTemplate) GoString() string {
	return s.String()
}

func (s *PipelineTemplate) SetCreatedTime(v string) *PipelineTemplate {
	s.CreatedTime = &v
	return s
}

func (s *PipelineTemplate) SetDeletionTime(v string) *PipelineTemplate {
	s.DeletionTime = &v
	return s
}

func (s *PipelineTemplate) SetDescription(v string) *PipelineTemplate {
	s.Description = &v
	return s
}

func (s *PipelineTemplate) SetGeneration(v int32) *PipelineTemplate {
	s.Generation = &v
	return s
}

func (s *PipelineTemplate) SetKind(v string) *PipelineTemplate {
	s.Kind = &v
	return s
}

func (s *PipelineTemplate) SetLabels(v map[string]*string) *PipelineTemplate {
	s.Labels = v
	return s
}

func (s *PipelineTemplate) SetName(v string) *PipelineTemplate {
	s.Name = &v
	return s
}

func (s *PipelineTemplate) SetResourceVersion(v int32) *PipelineTemplate {
	s.ResourceVersion = &v
	return s
}

func (s *PipelineTemplate) SetSpec(v *PipelineTemplateSpec) *PipelineTemplate {
	s.Spec = v
	return s
}

func (s *PipelineTemplate) SetUid(v string) *PipelineTemplate {
	s.Uid = &v
	return s
}

type PipelineTemplateSpec struct {
	Context *Context    `json:"context,omitempty" xml:"context,omitempty"`
	Tasks   []*TaskExec `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
}

func (s PipelineTemplateSpec) String() string {
	return tea.Prettify(s)
}

func (s PipelineTemplateSpec) GoString() string {
	return s.String()
}

func (s *PipelineTemplateSpec) SetContext(v *Context) *PipelineTemplateSpec {
	s.Context = v
	return s
}

func (s *PipelineTemplateSpec) SetTasks(v []*TaskExec) *PipelineTemplateSpec {
	s.Tasks = v
	return s
}

type PipelineTrigger struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// 2021-11-19T09:34:38Z
	DeletionTime *string `json:"deletionTime,omitempty" xml:"deletionTime,omitempty"`
	// example:
	//
	// PipelineTrigger example.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// example:
	//
	// PipelineTrigger
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-pipeline-trigger
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-project
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// example:
	//
	// 1
	ResourceVersion *int32               `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
	Spec            *PipelineTriggerSpec `json:"spec,omitempty" xml:"spec,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s PipelineTrigger) String() string {
	return tea.Prettify(s)
}

func (s PipelineTrigger) GoString() string {
	return s.String()
}

func (s *PipelineTrigger) SetCreatedTime(v string) *PipelineTrigger {
	s.CreatedTime = &v
	return s
}

func (s *PipelineTrigger) SetDeletionTime(v string) *PipelineTrigger {
	s.DeletionTime = &v
	return s
}

func (s *PipelineTrigger) SetDescription(v string) *PipelineTrigger {
	s.Description = &v
	return s
}

func (s *PipelineTrigger) SetGeneration(v int32) *PipelineTrigger {
	s.Generation = &v
	return s
}

func (s *PipelineTrigger) SetKind(v string) *PipelineTrigger {
	s.Kind = &v
	return s
}

func (s *PipelineTrigger) SetLabels(v map[string]*string) *PipelineTrigger {
	s.Labels = v
	return s
}

func (s *PipelineTrigger) SetName(v string) *PipelineTrigger {
	s.Name = &v
	return s
}

func (s *PipelineTrigger) SetProjectName(v string) *PipelineTrigger {
	s.ProjectName = &v
	return s
}

func (s *PipelineTrigger) SetResourceVersion(v int32) *PipelineTrigger {
	s.ResourceVersion = &v
	return s
}

func (s *PipelineTrigger) SetSpec(v *PipelineTriggerSpec) *PipelineTrigger {
	s.Spec = v
	return s
}

func (s *PipelineTrigger) SetUid(v string) *PipelineTrigger {
	s.Uid = &v
	return s
}

type PipelineTriggerEvent struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// 2021-11-19T09:34:38Z
	DeletionTime *string `json:"deletionTime,omitempty" xml:"deletionTime,omitempty"`
	// example:
	//
	// PipelineTriggerEvent example.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// example:
	//
	// PipelineTriggerEvent
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-pipeline-trigger
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	ResourceVersion *int32                      `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
	Spec            *PipelineTriggerEventSpec   `json:"spec,omitempty" xml:"spec,omitempty"`
	Status          *PipelineTriggerEventStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s PipelineTriggerEvent) String() string {
	return tea.Prettify(s)
}

func (s PipelineTriggerEvent) GoString() string {
	return s.String()
}

func (s *PipelineTriggerEvent) SetCreatedTime(v string) *PipelineTriggerEvent {
	s.CreatedTime = &v
	return s
}

func (s *PipelineTriggerEvent) SetDeletionTime(v string) *PipelineTriggerEvent {
	s.DeletionTime = &v
	return s
}

func (s *PipelineTriggerEvent) SetDescription(v string) *PipelineTriggerEvent {
	s.Description = &v
	return s
}

func (s *PipelineTriggerEvent) SetGeneration(v int32) *PipelineTriggerEvent {
	s.Generation = &v
	return s
}

func (s *PipelineTriggerEvent) SetKind(v string) *PipelineTriggerEvent {
	s.Kind = &v
	return s
}

func (s *PipelineTriggerEvent) SetLabels(v map[string]*string) *PipelineTriggerEvent {
	s.Labels = v
	return s
}

func (s *PipelineTriggerEvent) SetName(v string) *PipelineTriggerEvent {
	s.Name = &v
	return s
}

func (s *PipelineTriggerEvent) SetResourceVersion(v int32) *PipelineTriggerEvent {
	s.ResourceVersion = &v
	return s
}

func (s *PipelineTriggerEvent) SetSpec(v *PipelineTriggerEventSpec) *PipelineTriggerEvent {
	s.Spec = v
	return s
}

func (s *PipelineTriggerEvent) SetStatus(v *PipelineTriggerEventStatus) *PipelineTriggerEvent {
	s.Status = v
	return s
}

func (s *PipelineTriggerEvent) SetUid(v string) *PipelineTriggerEvent {
	s.Uid = &v
	return s
}

type PipelineTriggerEventSpec struct {
	Payload     *EventPayload `json:"payload,omitempty" xml:"payload,omitempty"`
	TriggerName *string       `json:"triggerName,omitempty" xml:"triggerName,omitempty"`
}

func (s PipelineTriggerEventSpec) String() string {
	return tea.Prettify(s)
}

func (s PipelineTriggerEventSpec) GoString() string {
	return s.String()
}

func (s *PipelineTriggerEventSpec) SetPayload(v *EventPayload) *PipelineTriggerEventSpec {
	s.Payload = v
	return s
}

func (s *PipelineTriggerEventSpec) SetTriggerName(v string) *PipelineTriggerEventSpec {
	s.TriggerName = &v
	return s
}

type PipelineTriggerEventStatus struct {
	// example:
	//
	// failed to create and start pipeline, error=message
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// my-pipeline
	FiredPipelineName *string        `json:"firedPipelineName,omitempty" xml:"firedPipelineName,omitempty"`
	ReleaseDetail     *ReleaseDetail `json:"releaseDetail,omitempty" xml:"releaseDetail,omitempty"`
	// example:
	//
	// Skipped,Fired,Running,Failed,Success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PipelineTriggerEventStatus) String() string {
	return tea.Prettify(s)
}

func (s PipelineTriggerEventStatus) GoString() string {
	return s.String()
}

func (s *PipelineTriggerEventStatus) SetErrorMessage(v string) *PipelineTriggerEventStatus {
	s.ErrorMessage = &v
	return s
}

func (s *PipelineTriggerEventStatus) SetFiredPipelineName(v string) *PipelineTriggerEventStatus {
	s.FiredPipelineName = &v
	return s
}

func (s *PipelineTriggerEventStatus) SetReleaseDetail(v *ReleaseDetail) *PipelineTriggerEventStatus {
	s.ReleaseDetail = v
	return s
}

func (s *PipelineTriggerEventStatus) SetStatus(v string) *PipelineTriggerEventStatus {
	s.Status = &v
	return s
}

type PipelineTriggerSpec struct {
	// This parameter is required.
	EventFilter *EventFilterConfig `json:"eventFilter,omitempty" xml:"eventFilter,omitempty"`
	// example:
	//
	// acs:ram::1431999****8149:role/aliyundevsdefaultrole
	RoleArn     *string            `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	RunPipeline *RunPipelineConfig `json:"runPipeline,omitempty" xml:"runPipeline,omitempty"`
}

func (s PipelineTriggerSpec) String() string {
	return tea.Prettify(s)
}

func (s PipelineTriggerSpec) GoString() string {
	return s.String()
}

func (s *PipelineTriggerSpec) SetEventFilter(v *EventFilterConfig) *PipelineTriggerSpec {
	s.EventFilter = v
	return s
}

func (s *PipelineTriggerSpec) SetRoleArn(v string) *PipelineTriggerSpec {
	s.RoleArn = &v
	return s
}

func (s *PipelineTriggerSpec) SetRunPipeline(v *RunPipelineConfig) *PipelineTriggerSpec {
	s.RunPipeline = v
	return s
}

type Project struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// test-description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// example:
	//
	// Project
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-project
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	ResourceVersion *int32         `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
	Spec            *ProjectSpec   `json:"spec,omitempty" xml:"spec,omitempty"`
	Status          *ProjectStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// example:
	//
	// 2021-11-19T09:34:38Z
	UpdatedTime *string `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
}

func (s Project) String() string {
	return tea.Prettify(s)
}

func (s Project) GoString() string {
	return s.String()
}

func (s *Project) SetCreatedTime(v string) *Project {
	s.CreatedTime = &v
	return s
}

func (s *Project) SetDescription(v string) *Project {
	s.Description = &v
	return s
}

func (s *Project) SetGeneration(v int32) *Project {
	s.Generation = &v
	return s
}

func (s *Project) SetKind(v string) *Project {
	s.Kind = &v
	return s
}

func (s *Project) SetLabels(v map[string]*string) *Project {
	s.Labels = v
	return s
}

func (s *Project) SetName(v string) *Project {
	s.Name = &v
	return s
}

func (s *Project) SetResourceVersion(v int32) *Project {
	s.ResourceVersion = &v
	return s
}

func (s *Project) SetSpec(v *ProjectSpec) *Project {
	s.Spec = v
	return s
}

func (s *Project) SetStatus(v *ProjectStatus) *Project {
	s.Status = v
	return s
}

func (s *Project) SetUid(v string) *Project {
	s.Uid = &v
	return s
}

func (s *Project) SetUpdatedTime(v string) *Project {
	s.UpdatedTime = &v
	return s
}

type ProjectSpec struct {
	// example:
	//
	// acs:ram::1431999****8149:role/aliyundevsdefaultrole
	RoleArn        *string         `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	TemplateConfig *TemplateConfig `json:"templateConfig,omitempty" xml:"templateConfig,omitempty"`
	// example:
	//
	// 9D72DE01-C732-49C0-8E85-FFD9D695436B
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s ProjectSpec) String() string {
	return tea.Prettify(s)
}

func (s ProjectSpec) GoString() string {
	return s.String()
}

func (s *ProjectSpec) SetRoleArn(v string) *ProjectSpec {
	s.RoleArn = &v
	return s
}

func (s *ProjectSpec) SetTemplateConfig(v *TemplateConfig) *ProjectSpec {
	s.TemplateConfig = v
	return s
}

func (s *ProjectSpec) SetToken(v string) *ProjectSpec {
	s.Token = &v
	return s
}

type ProjectStatus struct {
	LatestReleaseDetail *ReleaseDetail `json:"latestReleaseDetail,omitempty" xml:"latestReleaseDetail,omitempty"`
	// example:
	//
	// 1
	ObservedGeneration *int64 `json:"observedGeneration,omitempty" xml:"observedGeneration,omitempty"`
	// example:
	//
	// 2021-11-19T09:34:38Z
	ObservedTime *string `json:"observedTime,omitempty" xml:"observedTime,omitempty"`
}

func (s ProjectStatus) String() string {
	return tea.Prettify(s)
}

func (s ProjectStatus) GoString() string {
	return s.String()
}

func (s *ProjectStatus) SetLatestReleaseDetail(v *ReleaseDetail) *ProjectStatus {
	s.LatestReleaseDetail = v
	return s
}

func (s *ProjectStatus) SetObservedGeneration(v int64) *ProjectStatus {
	s.ObservedGeneration = &v
	return s
}

func (s *ProjectStatus) SetObservedTime(v string) *ProjectStatus {
	s.ObservedTime = &v
	return s
}

type PullRequestFilter struct {
	// example:
	//
	// feature-.*
	SourceBranch *string `json:"sourceBranch,omitempty" xml:"sourceBranch,omitempty"`
	// example:
	//
	// master
	TargetBranch *string   `json:"targetBranch,omitempty" xml:"targetBranch,omitempty"`
	Types        []*string `json:"types,omitempty" xml:"types,omitempty" type:"Repeated"`
}

func (s PullRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s PullRequestFilter) GoString() string {
	return s.String()
}

func (s *PullRequestFilter) SetSourceBranch(v string) *PullRequestFilter {
	s.SourceBranch = &v
	return s
}

func (s *PullRequestFilter) SetTargetBranch(v string) *PullRequestFilter {
	s.TargetBranch = &v
	return s
}

func (s *PullRequestFilter) SetTypes(v []*string) *PullRequestFilter {
	s.Types = v
	return s
}

type PushFilter struct {
	// example:
	//
	// master
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// prod-.*
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s PushFilter) String() string {
	return tea.Prettify(s)
}

func (s PushFilter) GoString() string {
	return s.String()
}

func (s *PushFilter) SetBranch(v string) *PushFilter {
	s.Branch = &v
	return s
}

func (s *PushFilter) SetTag(v string) *PushFilter {
	s.Tag = &v
	return s
}

type ReleaseDetail struct {
	// example:
	//
	// Running
	BizStatus *string `json:"bizStatus,omitempty" xml:"bizStatus,omitempty"`
	// example:
	//
	// dev
	EnvName *string `json:"envName,omitempty" xml:"envName,omitempty"`
	// example:
	//
	// 2021-11-19T09:34:38Z
	FinishedTime        *string           `json:"finishedTime,omitempty" xml:"finishedTime,omitempty"`
	GitEventSnapshot    *GitEventSnapshot `json:"gitEventSnapshot,omitempty" xml:"gitEventSnapshot,omitempty"`
	LatestTaskExecError *TaskExecError    `json:"latestTaskExecError,omitempty" xml:"latestTaskExecError,omitempty"`
	// example:
	//
	// Triggered manually.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// p-1704952599751-wUOczb
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// example:
	//
	// p-codeup-1714271977972-pa4w
	PipelineTriggerEventName *string `json:"pipelineTriggerEventName,omitempty" xml:"pipelineTriggerEventName,omitempty"`
	// example:
	//
	// {"framework":{"region":"cn-hangzhou","service":{"serviceName":"web-framework","description":"Serverless Devs Web Framework Service","role":"","logConfig":{"project":"","logstore":"","enableRequestMetrics":false,"enableInstanceMetrics":false,"logBeginRule":"None"},"serviceId":"4c9d0e79-16b8-4951-a6b8-169d2037d217","createdTime":"2021-12-07T09:24:08Z","lastModifiedTime":"2022-04-09T09:16:51Z","vpcConfig":{"vpcId":"","vSwitchIds":[],"securityGroupId":"","role":""},"internetAccess":true,"nasConfig":{"userId":-1,"groupId":-1,"mountPoints":[]},"tracingConfig":{},"name":"web-framework"},"function":{"functionId":"e81904f0-24d0-45df-bb53-06da64d01395","functionName":"todolist","description":"Serverless Devs Web Framework Function","runtime":"nodejs12","handler":"index.handler","timeout":60,"initializer":"","initializationTimeout":3,"codeSize":1757093,"codeChecksum":"7185648203525194222","memorySize":256,"environmentVariables":{},"createdTime":"2022-03-15T15:36:22Z","lastModifiedTime":"2022-04-09T09:16:50Z","instanceConcurrency":1,"instanceType":"e1","instanceLifecycleConfig":{"preFreeze":{"handler":"","timeout":3},"preStop":{"handler":"","timeout":3}},"name":"todolist"},"triggers":[{"triggerName":"httpTrigger","description":"","triggerId":"7f2373ce-df1a-4013-b4e5-899989d9b27e","triggerType":"http","triggerConfig":{"methods":["GET","POST"],"authType":"anonymous"},"createdTime":"2022-03-15T15:36:23Z","lastModifiedTime":"2022-04-09T09:16:51Z","name":"httpTrigger"}],"customDomains":[{"domainName":"todolist.web-framework.1835799444022432.cn-hangzhou.fc.devsapp.net","accountId":"1835799444022432","protocol":"HTTP","certConfig":{},"tlsConfig":{},"apiVersion":"2016-08-15","routeConfig":{"routes":[{"path":"/*","serviceName":"web-framework","functionName":"todolist","methods":["GET","POST"]}]},"createdTime":"2022-03-15T15:36:28Z","lastModifiedTime":"2022-04-09T09:17:00Z"}]}}
	ReleaseOutputs         map[string]interface{} `json:"releaseOutputs,omitempty" xml:"releaseOutputs,omitempty"`
	RepositorySnapshot     *RepositorySpec        `json:"repositorySnapshot,omitempty" xml:"repositorySnapshot,omitempty"`
	TemplateConfigSnapshot *TemplateConfig        `json:"templateConfigSnapshot,omitempty" xml:"templateConfigSnapshot,omitempty"`
}

func (s ReleaseDetail) String() string {
	return tea.Prettify(s)
}

func (s ReleaseDetail) GoString() string {
	return s.String()
}

func (s *ReleaseDetail) SetBizStatus(v string) *ReleaseDetail {
	s.BizStatus = &v
	return s
}

func (s *ReleaseDetail) SetEnvName(v string) *ReleaseDetail {
	s.EnvName = &v
	return s
}

func (s *ReleaseDetail) SetFinishedTime(v string) *ReleaseDetail {
	s.FinishedTime = &v
	return s
}

func (s *ReleaseDetail) SetGitEventSnapshot(v *GitEventSnapshot) *ReleaseDetail {
	s.GitEventSnapshot = v
	return s
}

func (s *ReleaseDetail) SetLatestTaskExecError(v *TaskExecError) *ReleaseDetail {
	s.LatestTaskExecError = v
	return s
}

func (s *ReleaseDetail) SetMessage(v string) *ReleaseDetail {
	s.Message = &v
	return s
}

func (s *ReleaseDetail) SetPipelineName(v string) *ReleaseDetail {
	s.PipelineName = &v
	return s
}

func (s *ReleaseDetail) SetPipelineTriggerEventName(v string) *ReleaseDetail {
	s.PipelineTriggerEventName = &v
	return s
}

func (s *ReleaseDetail) SetReleaseOutputs(v map[string]interface{}) *ReleaseDetail {
	s.ReleaseOutputs = v
	return s
}

func (s *ReleaseDetail) SetRepositorySnapshot(v *RepositorySpec) *ReleaseDetail {
	s.RepositorySnapshot = v
	return s
}

func (s *ReleaseDetail) SetTemplateConfigSnapshot(v *TemplateConfig) *ReleaseDetail {
	s.TemplateConfigSnapshot = v
	return s
}

type Repository struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// test-description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// example:
	//
	// Repository
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-repository
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	ResourceVersion *int32 `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
	// This parameter is required.
	Spec *RepositorySpec `json:"spec,omitempty" xml:"spec,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s Repository) String() string {
	return tea.Prettify(s)
}

func (s Repository) GoString() string {
	return s.String()
}

func (s *Repository) SetCreatedTime(v string) *Repository {
	s.CreatedTime = &v
	return s
}

func (s *Repository) SetDescription(v string) *Repository {
	s.Description = &v
	return s
}

func (s *Repository) SetGeneration(v int32) *Repository {
	s.Generation = &v
	return s
}

func (s *Repository) SetKind(v string) *Repository {
	s.Kind = &v
	return s
}

func (s *Repository) SetLabels(v map[string]*string) *Repository {
	s.Labels = v
	return s
}

func (s *Repository) SetName(v string) *Repository {
	s.Name = &v
	return s
}

func (s *Repository) SetResourceVersion(v int32) *Repository {
	s.ResourceVersion = &v
	return s
}

func (s *Repository) SetSpec(v *RepositorySpec) *Repository {
	s.Spec = v
	return s
}

func (s *Repository) SetUid(v string) *Repository {
	s.Uid = &v
	return s
}

type RepositoryConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// master
	BranchName *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
	// example:
	//
	// src/s.yaml
	Manifest *string `json:"manifest,omitempty" xml:"manifest,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-repository
	RepositoryName *string `json:"repositoryName,omitempty" xml:"repositoryName,omitempty"`
}

func (s RepositoryConfig) String() string {
	return tea.Prettify(s)
}

func (s RepositoryConfig) GoString() string {
	return s.String()
}

func (s *RepositoryConfig) SetBranchName(v string) *RepositoryConfig {
	s.BranchName = &v
	return s
}

func (s *RepositoryConfig) SetManifest(v string) *RepositoryConfig {
	s.Manifest = &v
	return s
}

func (s *RepositoryConfig) SetRepositoryName(v string) *RepositoryConfig {
	s.RepositoryName = &v
	return s
}

type RepositorySpec struct {
	// This parameter is required.
	//
	// example:
	//
	// https://github.com/DDofDD/start-springboot-lfgy.git
	CloneUrl *string `json:"cloneUrl,omitempty" xml:"cloneUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// awesome-connection
	ConnectionName *string `json:"connectionName,omitempty" xml:"connectionName,omitempty"`
	// example:
	//
	// my-repo-name
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// 312649
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// my-org-name
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// example:
	//
	// github
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// example:
	//
	// https://github.com/my-org-name/my-repo-name
	WebUrl *string `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
}

func (s RepositorySpec) String() string {
	return tea.Prettify(s)
}

func (s RepositorySpec) GoString() string {
	return s.String()
}

func (s *RepositorySpec) SetCloneUrl(v string) *RepositorySpec {
	s.CloneUrl = &v
	return s
}

func (s *RepositorySpec) SetConnectionName(v string) *RepositorySpec {
	s.ConnectionName = &v
	return s
}

func (s *RepositorySpec) SetDisplayName(v string) *RepositorySpec {
	s.DisplayName = &v
	return s
}

func (s *RepositorySpec) SetId(v int64) *RepositorySpec {
	s.Id = &v
	return s
}

func (s *RepositorySpec) SetOwner(v string) *RepositorySpec {
	s.Owner = &v
	return s
}

func (s *RepositorySpec) SetPlatform(v string) *RepositorySpec {
	s.Platform = &v
	return s
}

func (s *RepositorySpec) SetWebUrl(v string) *RepositorySpec {
	s.WebUrl = &v
	return s
}

type ResourceDetail struct {
	// example:
	//
	// alicloud_fc_trigger.cn-shanghai-fc-stable-diffusion-sd
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// {"config":"{\"methods\":[\"GET\",\"POST\",\"PUT\",\"DELETE\"],\"authType\":\"anonymous\",\"disableURLInternet\":true}","config_mns":null,"function":"sd","id":"fc-stable-diffusion:sd:defaultTrigger","last_modified":"2024-04-17T13:20:53Z","name":"defaultTrigger","name_prefix":null,"role":"","service":"fc-stable-diffusion","source_arn":"","trigger_id":"mock-trigger","type":"http"}
	AttributeValues map[string]interface{} `json:"attributeValues,omitempty" xml:"attributeValues,omitempty"`
	// example:
	//
	// managed
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// example:
	//
	// cn-shanghai-fc-stable-diffusion-sd
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// alicloud_fc_trigger
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ResourceDetail) String() string {
	return tea.Prettify(s)
}

func (s ResourceDetail) GoString() string {
	return s.String()
}

func (s *ResourceDetail) SetAddress(v string) *ResourceDetail {
	s.Address = &v
	return s
}

func (s *ResourceDetail) SetAttributeValues(v map[string]interface{}) *ResourceDetail {
	s.AttributeValues = v
	return s
}

func (s *ResourceDetail) SetMode(v string) *ResourceDetail {
	s.Mode = &v
	return s
}

func (s *ResourceDetail) SetName(v string) *ResourceDetail {
	s.Name = &v
	return s
}

func (s *ResourceDetail) SetType(v string) *ResourceDetail {
	s.Type = &v
	return s
}

type ResourceDrift struct {
	// example:
	//
	// alicloud_fc_trigger.cn-shanghai-fc-stable-diffusion-sd
	Address *string              `json:"address,omitempty" xml:"address,omitempty"`
	Change  *ResourceDriftChange `json:"change,omitempty" xml:"change,omitempty" type:"Struct"`
	// example:
	//
	// managed
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// example:
	//
	// cn-shanghai-fc-stable-diffusion-sd
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// alicloud_fc_trigger
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ResourceDrift) String() string {
	return tea.Prettify(s)
}

func (s ResourceDrift) GoString() string {
	return s.String()
}

func (s *ResourceDrift) SetAddress(v string) *ResourceDrift {
	s.Address = &v
	return s
}

func (s *ResourceDrift) SetChange(v *ResourceDriftChange) *ResourceDrift {
	s.Change = v
	return s
}

func (s *ResourceDrift) SetMode(v string) *ResourceDrift {
	s.Mode = &v
	return s
}

func (s *ResourceDrift) SetName(v string) *ResourceDrift {
	s.Name = &v
	return s
}

func (s *ResourceDrift) SetType(v string) *ResourceDrift {
	s.Type = &v
	return s
}

type ResourceDriftChange struct {
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// example:
	//
	// {"description":"mock deployment","id":"fc-demo-model-scope","internet_access":true,"last_modified":"2024-04-17T13:20:56Z","log_config":[],"name":"fc-demo-model-scope","name_prefix":null,"nas_config":[{"group_id":0,"mount_points":[{"mount_dir":"/mnt/auto","server_addr":"mock-addr.cn-shanghai.nas.aliyuncs.com:/fc-demo-model-scope"}],"user_id":0}],"publish":null,"role":"acs:ram::mock-role:role/aliyunfcdefaultrole","service_id":"mock-service","tags":null,"tracing_config":[],"version":null,"vpc_config":[{"security_group_id":"sg-mock","vpc_id":"vpc-mock","vswitch_ids":["vsw-mock"]}]}
	After interface{} `json:"after,omitempty" xml:"after,omitempty"`
	// example:
	//
	// {"description":"modelscope deployment","id":"fc-demo-model-scope","internet_access":true,"last_modified":"2024-04-17T13:20:56Z","log_config":[],"name":"fc-demo-model-scope","name_prefix":null,"nas_config":[{"group_id":0,"mount_points":[{"mount_dir":"/mnt/auto","server_addr":"mock-addr.cn-shanghai.nas.aliyuncs.com:/fc-demo-model-scope"}],"user_id":0}],"publish":null,"role":"acs:ram::mock-role:role/aliyunfcdefaultrole","service_id":"mock-service","tags":null,"tracing_config":[],"version":null,"vpc_config":[{"security_group_id":"sg-mock","vpc_id":"vpc-mock","vswitch_ids":["vsw-mock"]}]}
	Before interface{} `json:"before,omitempty" xml:"before,omitempty"`
}

func (s ResourceDriftChange) String() string {
	return tea.Prettify(s)
}

func (s ResourceDriftChange) GoString() string {
	return s.String()
}

func (s *ResourceDriftChange) SetActions(v []*string) *ResourceDriftChange {
	s.Actions = v
	return s
}

func (s *ResourceDriftChange) SetAfter(v interface{}) *ResourceDriftChange {
	s.After = v
	return s
}

func (s *ResourceDriftChange) SetBefore(v interface{}) *ResourceDriftChange {
	s.Before = v
	return s
}

type RunAfter struct {
	// example:
	//
	// task-1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RunAfter) String() string {
	return tea.Prettify(s)
}

func (s RunAfter) GoString() string {
	return s.String()
}

func (s *RunAfter) SetName(v string) *RunAfter {
	s.Name = &v
	return s
}

type RunPipelineConfig struct {
	PipelineSpec    *PipelineSpec `json:"pipelineSpec,omitempty" xml:"pipelineSpec,omitempty"`
	Variables       []*Variable   `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
	YamlFileContent *string       `json:"yamlFileContent,omitempty" xml:"yamlFileContent,omitempty"`
	YamlFilePath    *string       `json:"yamlFilePath,omitempty" xml:"yamlFilePath,omitempty"`
}

func (s RunPipelineConfig) String() string {
	return tea.Prettify(s)
}

func (s RunPipelineConfig) GoString() string {
	return s.String()
}

func (s *RunPipelineConfig) SetPipelineSpec(v *PipelineSpec) *RunPipelineConfig {
	s.PipelineSpec = v
	return s
}

func (s *RunPipelineConfig) SetVariables(v []*Variable) *RunPipelineConfig {
	s.Variables = v
	return s
}

func (s *RunPipelineConfig) SetYamlFileContent(v string) *RunPipelineConfig {
	s.YamlFileContent = &v
	return s
}

func (s *RunPipelineConfig) SetYamlFilePath(v string) *RunPipelineConfig {
	s.YamlFilePath = &v
	return s
}

type ServiceSpec struct {
	// This parameter is required.
	//
	// example:
	//
	// my-env
	Environment *string `json:"environment,omitempty" xml:"environment,omitempty"`
	// example:
	//
	// acs:ram::1455541096306548:role/aliyunfcdefaultrole
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// serverless-devs.com/alicloud-multi-functions/service-single-function/deployment
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
	// This parameter is required.
	TemplateVariables map[string]interface{} `json:"templateVariables,omitempty" xml:"templateVariables,omitempty"`
	// example:
	//
	// 1
	TemplateVersion *int32 `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
}

func (s ServiceSpec) String() string {
	return tea.Prettify(s)
}

func (s ServiceSpec) GoString() string {
	return s.String()
}

func (s *ServiceSpec) SetEnvironment(v string) *ServiceSpec {
	s.Environment = &v
	return s
}

func (s *ServiceSpec) SetRoleArn(v string) *ServiceSpec {
	s.RoleArn = &v
	return s
}

func (s *ServiceSpec) SetTemplate(v string) *ServiceSpec {
	s.Template = &v
	return s
}

func (s *ServiceSpec) SetTemplateVariables(v map[string]interface{}) *ServiceSpec {
	s.TemplateVariables = v
	return s
}

func (s *ServiceSpec) SetTemplateVersion(v int32) *ServiceSpec {
	s.TemplateVersion = &v
	return s
}

type Task struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// 2021-11-19T09:34:38Z
	DeletionTime *string `json:"deletionTime,omitempty" xml:"deletionTime,omitempty"`
	// example:
	//
	// Task example.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// example:
	//
	// Task
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-task
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	ResourceVersion *int32      `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
	Spec            *TaskSpec   `json:"spec,omitempty" xml:"spec,omitempty"`
	Status          *TaskStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s Task) String() string {
	return tea.Prettify(s)
}

func (s Task) GoString() string {
	return s.String()
}

func (s *Task) SetCreatedTime(v string) *Task {
	s.CreatedTime = &v
	return s
}

func (s *Task) SetDeletionTime(v string) *Task {
	s.DeletionTime = &v
	return s
}

func (s *Task) SetDescription(v string) *Task {
	s.Description = &v
	return s
}

func (s *Task) SetGeneration(v int32) *Task {
	s.Generation = &v
	return s
}

func (s *Task) SetKind(v string) *Task {
	s.Kind = &v
	return s
}

func (s *Task) SetLabels(v map[string]*string) *Task {
	s.Labels = v
	return s
}

func (s *Task) SetName(v string) *Task {
	s.Name = &v
	return s
}

func (s *Task) SetResourceVersion(v int32) *Task {
	s.ResourceVersion = &v
	return s
}

func (s *Task) SetSpec(v *TaskSpec) *Task {
	s.Spec = v
	return s
}

func (s *Task) SetStatus(v *TaskStatus) *Task {
	s.Status = v
	return s
}

func (s *Task) SetUid(v string) *Task {
	s.Uid = &v
	return s
}

type TaskExec struct {
	Context *Context `json:"context,omitempty" xml:"context,omitempty"`
	// example:
	//
	// task-1
	Name      *string     `json:"name,omitempty" xml:"name,omitempty"`
	RunAfters []*RunAfter `json:"runAfters,omitempty" xml:"runAfters,omitempty" type:"Repeated"`
	// example:
	//
	// serverless-runner
	TaskTemplate *string `json:"taskTemplate,omitempty" xml:"taskTemplate,omitempty"`
}

func (s TaskExec) String() string {
	return tea.Prettify(s)
}

func (s TaskExec) GoString() string {
	return s.String()
}

func (s *TaskExec) SetContext(v *Context) *TaskExec {
	s.Context = v
	return s
}

func (s *TaskExec) SetName(v string) *TaskExec {
	s.Name = &v
	return s
}

func (s *TaskExec) SetRunAfters(v []*RunAfter) *TaskExec {
	s.RunAfters = v
	return s
}

func (s *TaskExec) SetTaskTemplate(v string) *TaskExec {
	s.TaskTemplate = &v
	return s
}

type TaskExecError struct {
	// example:
	//
	// AccessDenied
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 部署辅助函数权限不足，需要添加额外的权限以解决问题。https://help.aliyun.com
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// example:
	//
	// 部署服务[_appcenter-xxx]失败，权限不足
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 1-26d1287xxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 权限不足错误
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s TaskExecError) String() string {
	return tea.Prettify(s)
}

func (s TaskExecError) GoString() string {
	return s.String()
}

func (s *TaskExecError) SetCode(v string) *TaskExecError {
	s.Code = &v
	return s
}

func (s *TaskExecError) SetExtraInfo(v string) *TaskExecError {
	s.ExtraInfo = &v
	return s
}

func (s *TaskExecError) SetMessage(v string) *TaskExecError {
	s.Message = &v
	return s
}

func (s *TaskExecError) SetRequestId(v string) *TaskExecError {
	s.RequestId = &v
	return s
}

func (s *TaskExecError) SetTitle(v string) *TaskExecError {
	s.Title = &v
	return s
}

type TaskInvocation struct {
	// example:
	//
	// c-nkj8shz7xxxx
	InstanceID *string `json:"instanceID,omitempty" xml:"instanceID,omitempty"`
	// example:
	//
	// E099843B-10A2-4936-9964-4E0EE263D564
	InvocationID *string `json:"invocationID,omitempty" xml:"invocationID,omitempty"`
	// example:
	//
	// acs:fc:cn-hangzhou:143xxxx:services/xxx.LATEST/functions/xxx
	InvocationTarget *string `json:"invocationTarget,omitempty" xml:"invocationTarget,omitempty"`
	// example:
	//
	// {"key1":"value1","key2":"value2"}
	Output *string `json:"output,omitempty" xml:"output,omitempty"`
	// example:
	//
	// 1B3058B1-F1C9-457C-B95C-2C250A4B3118
	RequestID *string `json:"requestID,omitempty" xml:"requestID,omitempty"`
	// example:
	//
	// my-sls-logstore
	SlsLogStore *string `json:"slsLogStore,omitempty" xml:"slsLogStore,omitempty"`
	// example:
	//
	// my-sls-project
	SlsProject *string `json:"slsProject,omitempty" xml:"slsProject,omitempty"`
	// example:
	//
	// success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s TaskInvocation) String() string {
	return tea.Prettify(s)
}

func (s TaskInvocation) GoString() string {
	return s.String()
}

func (s *TaskInvocation) SetInstanceID(v string) *TaskInvocation {
	s.InstanceID = &v
	return s
}

func (s *TaskInvocation) SetInvocationID(v string) *TaskInvocation {
	s.InvocationID = &v
	return s
}

func (s *TaskInvocation) SetInvocationTarget(v string) *TaskInvocation {
	s.InvocationTarget = &v
	return s
}

func (s *TaskInvocation) SetOutput(v string) *TaskInvocation {
	s.Output = &v
	return s
}

func (s *TaskInvocation) SetRequestID(v string) *TaskInvocation {
	s.RequestID = &v
	return s
}

func (s *TaskInvocation) SetSlsLogStore(v string) *TaskInvocation {
	s.SlsLogStore = &v
	return s
}

func (s *TaskInvocation) SetSlsProject(v string) *TaskInvocation {
	s.SlsProject = &v
	return s
}

func (s *TaskInvocation) SetStatus(v string) *TaskInvocation {
	s.Status = &v
	return s
}

type TaskSpec struct {
	Context *Context `json:"context,omitempty" xml:"context,omitempty"`
	// example:
	//
	// my-task-template
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s TaskSpec) String() string {
	return tea.Prettify(s)
}

func (s TaskSpec) GoString() string {
	return s.String()
}

func (s *TaskSpec) SetContext(v *Context) *TaskSpec {
	s.Context = v
	return s
}

func (s *TaskSpec) SetTemplateName(v string) *TaskSpec {
	s.TemplateName = &v
	return s
}

type TaskStatus struct {
	ExecutionDetails []*string         `json:"executionDetails,omitempty" xml:"executionDetails,omitempty" type:"Repeated"`
	Invocations      []*TaskInvocation `json:"invocations,omitempty" xml:"invocations,omitempty" type:"Repeated"`
	LatestExecError  *TaskExecError    `json:"latestExecError,omitempty" xml:"latestExecError,omitempty"`
	// example:
	//
	// Success
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
	// example:
	//
	// 123
	StatusGeneration *int64 `json:"statusGeneration,omitempty" xml:"statusGeneration,omitempty"`
}

func (s TaskStatus) String() string {
	return tea.Prettify(s)
}

func (s TaskStatus) GoString() string {
	return s.String()
}

func (s *TaskStatus) SetExecutionDetails(v []*string) *TaskStatus {
	s.ExecutionDetails = v
	return s
}

func (s *TaskStatus) SetInvocations(v []*TaskInvocation) *TaskStatus {
	s.Invocations = v
	return s
}

func (s *TaskStatus) SetLatestExecError(v *TaskExecError) *TaskStatus {
	s.LatestExecError = v
	return s
}

func (s *TaskStatus) SetPhase(v string) *TaskStatus {
	s.Phase = &v
	return s
}

func (s *TaskStatus) SetStatusGeneration(v int64) *TaskStatus {
	s.StatusGeneration = &v
	return s
}

type TaskTemplate struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// 2021-11-19T09:34:38Z
	DeletionTime *string `json:"deletionTime,omitempty" xml:"deletionTime,omitempty"`
	// example:
	//
	// TaskTemplate example.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// example:
	//
	// TaskTemplate
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-task-template
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	ResourceVersion *int32            `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
	Spec            *TaskTemplateSpec `json:"spec,omitempty" xml:"spec,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s TaskTemplate) String() string {
	return tea.Prettify(s)
}

func (s TaskTemplate) GoString() string {
	return s.String()
}

func (s *TaskTemplate) SetCreatedTime(v string) *TaskTemplate {
	s.CreatedTime = &v
	return s
}

func (s *TaskTemplate) SetDeletionTime(v string) *TaskTemplate {
	s.DeletionTime = &v
	return s
}

func (s *TaskTemplate) SetDescription(v string) *TaskTemplate {
	s.Description = &v
	return s
}

func (s *TaskTemplate) SetGeneration(v int32) *TaskTemplate {
	s.Generation = &v
	return s
}

func (s *TaskTemplate) SetKind(v string) *TaskTemplate {
	s.Kind = &v
	return s
}

func (s *TaskTemplate) SetLabels(v map[string]*string) *TaskTemplate {
	s.Labels = v
	return s
}

func (s *TaskTemplate) SetName(v string) *TaskTemplate {
	s.Name = &v
	return s
}

func (s *TaskTemplate) SetResourceVersion(v int32) *TaskTemplate {
	s.ResourceVersion = &v
	return s
}

func (s *TaskTemplate) SetSpec(v *TaskTemplateSpec) *TaskTemplate {
	s.Spec = v
	return s
}

func (s *TaskTemplate) SetUid(v string) *TaskTemplate {
	s.Uid = &v
	return s
}

type TaskTemplateSpec struct {
	Context *Context `json:"context,omitempty" xml:"context,omitempty"`
	// example:
	//
	// build&deploy.
	Description      *string     `json:"description,omitempty" xml:"description,omitempty"`
	ExecuteCondition *Condition  `json:"executeCondition,omitempty" xml:"executeCondition,omitempty"`
	Worker           *TaskWorker `json:"worker,omitempty" xml:"worker,omitempty"`
}

func (s TaskTemplateSpec) String() string {
	return tea.Prettify(s)
}

func (s TaskTemplateSpec) GoString() string {
	return s.String()
}

func (s *TaskTemplateSpec) SetContext(v *Context) *TaskTemplateSpec {
	s.Context = v
	return s
}

func (s *TaskTemplateSpec) SetDescription(v string) *TaskTemplateSpec {
	s.Description = &v
	return s
}

func (s *TaskTemplateSpec) SetExecuteCondition(v *Condition) *TaskTemplateSpec {
	s.ExecuteCondition = v
	return s
}

func (s *TaskTemplateSpec) SetWorker(v *TaskWorker) *TaskTemplateSpec {
	s.Worker = v
	return s
}

type TaskWorker struct {
	// example:
	//
	// serverless-runner
	PresetWorker *string `json:"presetWorker,omitempty" xml:"presetWorker,omitempty"`
}

func (s TaskWorker) String() string {
	return tea.Prettify(s)
}

func (s TaskWorker) GoString() string {
	return s.String()
}

func (s *TaskWorker) SetPresetWorker(v string) *TaskWorker {
	s.PresetWorker = &v
	return s
}

type TemplateConfig struct {
	// example:
	//
	// {"region":"cn-hangzhou"}
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// start-springboot
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s TemplateConfig) String() string {
	return tea.Prettify(s)
}

func (s TemplateConfig) GoString() string {
	return s.String()
}

func (s *TemplateConfig) SetParameters(v map[string]interface{}) *TemplateConfig {
	s.Parameters = v
	return s
}

func (s *TemplateConfig) SetTemplateName(v string) *TemplateConfig {
	s.TemplateName = &v
	return s
}

type TerraformInputVariable struct {
	// example:
	//
	// {"key":"value"}
	DefaultJson *string `json:"defaultJson,omitempty" xml:"defaultJson,omitempty"`
	// example:
	//
	// test variable
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// namePrefix
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// true
	Nullable *bool `json:"nullable,omitempty" xml:"nullable,omitempty"`
	// example:
	//
	// false
	Sensitive *bool `json:"sensitive,omitempty" xml:"sensitive,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TerraformInputVariable) String() string {
	return tea.Prettify(s)
}

func (s TerraformInputVariable) GoString() string {
	return s.String()
}

func (s *TerraformInputVariable) SetDefaultJson(v string) *TerraformInputVariable {
	s.DefaultJson = &v
	return s
}

func (s *TerraformInputVariable) SetDescription(v string) *TerraformInputVariable {
	s.Description = &v
	return s
}

func (s *TerraformInputVariable) SetName(v string) *TerraformInputVariable {
	s.Name = &v
	return s
}

func (s *TerraformInputVariable) SetNullable(v bool) *TerraformInputVariable {
	s.Nullable = &v
	return s
}

func (s *TerraformInputVariable) SetSensitive(v bool) *TerraformInputVariable {
	s.Sensitive = &v
	return s
}

func (s *TerraformInputVariable) SetType(v string) *TerraformInputVariable {
	s.Type = &v
	return s
}

type TerraformOutputValue struct {
	// example:
	//
	// The VPC ID where the resource is located.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpcId
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// false
	Sensitive *bool `json:"sensitive,omitempty" xml:"sensitive,omitempty"`
}

func (s TerraformOutputValue) String() string {
	return tea.Prettify(s)
}

func (s TerraformOutputValue) GoString() string {
	return s.String()
}

func (s *TerraformOutputValue) SetDescription(v string) *TerraformOutputValue {
	s.Description = &v
	return s
}

func (s *TerraformOutputValue) SetName(v string) *TerraformOutputValue {
	s.Name = &v
	return s
}

func (s *TerraformOutputValue) SetSensitive(v bool) *TerraformOutputValue {
	s.Sensitive = &v
	return s
}

type Variable struct {
	// example:
	//
	// object_key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// object_value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s Variable) String() string {
	return tea.Prettify(s)
}

func (s Variable) GoString() string {
	return s.String()
}

func (s *Variable) SetName(v string) *Variable {
	s.Name = &v
	return s
}

func (s *Variable) SetValue(v string) *Variable {
	s.Value = &v
	return s
}

type ActivateConnectionRequest struct {
	Account    *GitAccount      `json:"account,omitempty" xml:"account,omitempty"`
	Credential *OAuthCredential `json:"credential,omitempty" xml:"credential,omitempty"`
}

func (s ActivateConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivateConnectionRequest) GoString() string {
	return s.String()
}

func (s *ActivateConnectionRequest) SetAccount(v *GitAccount) *ActivateConnectionRequest {
	s.Account = v
	return s
}

func (s *ActivateConnectionRequest) SetCredential(v *OAuthCredential) *ActivateConnectionRequest {
	s.Credential = v
	return s
}

type ActivateConnectionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Connection        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivateConnectionResponse) GoString() string {
	return s.String()
}

func (s *ActivateConnectionResponse) SetHeaders(v map[string]*string) *ActivateConnectionResponse {
	s.Headers = v
	return s
}

func (s *ActivateConnectionResponse) SetStatusCode(v int32) *ActivateConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateConnectionResponse) SetBody(v *Connection) *ActivateConnectionResponse {
	s.Body = v
	return s
}

type CancelPipelineResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Pipeline          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelPipelineResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelPipelineResponse) GoString() string {
	return s.String()
}

func (s *CancelPipelineResponse) SetHeaders(v map[string]*string) *CancelPipelineResponse {
	s.Headers = v
	return s
}

func (s *CancelPipelineResponse) SetStatusCode(v int32) *CancelPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelPipelineResponse) SetBody(v *Pipeline) *CancelPipelineResponse {
	s.Body = v
	return s
}

type CancelTaskResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Task              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelTaskResponse) SetHeaders(v map[string]*string) *CancelTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelTaskResponse) SetStatusCode(v int32) *CancelTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelTaskResponse) SetBody(v *Task) *CancelTaskResponse {
	s.Body = v
	return s
}

type CreateConnectionRequest struct {
	Body *Connection `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateConnectionRequest) SetBody(v *Connection) *CreateConnectionRequest {
	s.Body = v
	return s
}

type CreateConnectionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Connection        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateConnectionResponse) SetHeaders(v map[string]*string) *CreateConnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateConnectionResponse) SetStatusCode(v int32) *CreateConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConnectionResponse) SetBody(v *Connection) *CreateConnectionResponse {
	s.Body = v
	return s
}

type CreateEnvironmentRequest struct {
	Body *Environment `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEnvironmentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentRequest) SetBody(v *Environment) *CreateEnvironmentRequest {
	s.Body = v
	return s
}

type CreateEnvironmentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Environment       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentResponse) SetHeaders(v map[string]*string) *CreateEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *CreateEnvironmentResponse) SetStatusCode(v int32) *CreateEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEnvironmentResponse) SetBody(v *Environment) *CreateEnvironmentResponse {
	s.Body = v
	return s
}

type CreatePipelineRequest struct {
	Body *Pipeline `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePipelineRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequest) SetBody(v *Pipeline) *CreatePipelineRequest {
	s.Body = v
	return s
}

type CreatePipelineResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Pipeline          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePipelineResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineResponse) GoString() string {
	return s.String()
}

func (s *CreatePipelineResponse) SetHeaders(v map[string]*string) *CreatePipelineResponse {
	s.Headers = v
	return s
}

func (s *CreatePipelineResponse) SetStatusCode(v int32) *CreatePipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePipelineResponse) SetBody(v *Pipeline) *CreatePipelineResponse {
	s.Body = v
	return s
}

type CreatePipelineTemplateRequest struct {
	Body *PipelineTemplate `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePipelineTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineTemplateRequest) SetBody(v *PipelineTemplate) *CreatePipelineTemplateRequest {
	s.Body = v
	return s
}

type CreatePipelineTemplateResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PipelineTemplate  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePipelineTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreatePipelineTemplateResponse) SetHeaders(v map[string]*string) *CreatePipelineTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreatePipelineTemplateResponse) SetStatusCode(v int32) *CreatePipelineTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePipelineTemplateResponse) SetBody(v *PipelineTemplate) *CreatePipelineTemplateResponse {
	s.Body = v
	return s
}

type CreatePipelineTriggerRequest struct {
	Body *PipelineTrigger `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePipelineTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineTriggerRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineTriggerRequest) SetBody(v *PipelineTrigger) *CreatePipelineTriggerRequest {
	s.Body = v
	return s
}

type CreatePipelineTriggerResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PipelineTrigger   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePipelineTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineTriggerResponse) GoString() string {
	return s.String()
}

func (s *CreatePipelineTriggerResponse) SetHeaders(v map[string]*string) *CreatePipelineTriggerResponse {
	s.Headers = v
	return s
}

func (s *CreatePipelineTriggerResponse) SetStatusCode(v int32) *CreatePipelineTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePipelineTriggerResponse) SetBody(v *PipelineTrigger) *CreatePipelineTriggerResponse {
	s.Body = v
	return s
}

type CreatePipelineTriggerEventRequest struct {
	Body *PipelineTriggerEvent `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePipelineTriggerEventRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineTriggerEventRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineTriggerEventRequest) SetBody(v *PipelineTriggerEvent) *CreatePipelineTriggerEventRequest {
	s.Body = v
	return s
}

type CreatePipelineTriggerEventResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PipelineTriggerEvent `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePipelineTriggerEventResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineTriggerEventResponse) GoString() string {
	return s.String()
}

func (s *CreatePipelineTriggerEventResponse) SetHeaders(v map[string]*string) *CreatePipelineTriggerEventResponse {
	s.Headers = v
	return s
}

func (s *CreatePipelineTriggerEventResponse) SetStatusCode(v int32) *CreatePipelineTriggerEventResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePipelineTriggerEventResponse) SetBody(v *PipelineTriggerEvent) *CreatePipelineTriggerEventResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	Body *Project `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetBody(v *Project) *CreateProjectRequest {
	s.Body = v
	return s
}

type CreateProjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Project           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectResponse) SetHeaders(v map[string]*string) *CreateProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectResponse) SetStatusCode(v int32) *CreateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectResponse) SetBody(v *Project) *CreateProjectResponse {
	s.Body = v
	return s
}

type CreateRepositoryRequest struct {
	Body *Repository `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepositoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryRequest) GoString() string {
	return s.String()
}

func (s *CreateRepositoryRequest) SetBody(v *Repository) *CreateRepositoryRequest {
	s.Body = v
	return s
}

type CreateRepositoryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Repository        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryResponse) GoString() string {
	return s.String()
}

func (s *CreateRepositoryResponse) SetHeaders(v map[string]*string) *CreateRepositoryResponse {
	s.Headers = v
	return s
}

func (s *CreateRepositoryResponse) SetStatusCode(v int32) *CreateRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepositoryResponse) SetBody(v *Repository) *CreateRepositoryResponse {
	s.Body = v
	return s
}

type CreateTaskRequest struct {
	Body *Task `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskRequest) SetBody(v *Task) *CreateTaskRequest {
	s.Body = v
	return s
}

type CreateTaskResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Task              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskResponse) SetHeaders(v map[string]*string) *CreateTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskResponse) SetStatusCode(v int32) *CreateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTaskResponse) SetBody(v *Task) *CreateTaskResponse {
	s.Body = v
	return s
}

type CreateTaskTemplateRequest struct {
	Body *TaskTemplate `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTaskTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskTemplateRequest) SetBody(v *TaskTemplate) *CreateTaskTemplateRequest {
	s.Body = v
	return s
}

type CreateTaskTemplateResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskTemplate      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTaskTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskTemplateResponse) SetHeaders(v map[string]*string) *CreateTaskTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskTemplateResponse) SetStatusCode(v int32) *CreateTaskTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTaskTemplateResponse) SetBody(v *TaskTemplate) *CreateTaskTemplateResponse {
	s.Body = v
	return s
}

type DeleteConnectionRequest struct {
	// example:
	//
	// true
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s DeleteConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConnectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteConnectionRequest) SetForce(v bool) *DeleteConnectionRequest {
	s.Force = &v
	return s
}

type DeleteConnectionResponseBody struct {
	// example:
	//
	// A5152937-1C8A-5260-90FA-520CEF028D2D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConnectionResponseBody) SetRequestId(v string) *DeleteConnectionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteConnectionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConnectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteConnectionResponse) SetHeaders(v map[string]*string) *DeleteConnectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteConnectionResponse) SetStatusCode(v int32) *DeleteConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConnectionResponse) SetBody(v *DeleteConnectionResponseBody) *DeleteConnectionResponse {
	s.Body = v
	return s
}

type DeleteEnvironmentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentResponse) SetHeaders(v map[string]*string) *DeleteEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnvironmentResponse) SetStatusCode(v int32) *DeleteEnvironmentResponse {
	s.StatusCode = &v
	return s
}

type DeletePipelineTemplateResponseBody struct {
	// example:
	//
	// E2977805-E133-5966-878B-6499E6A04D3C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeletePipelineTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelineTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePipelineTemplateResponseBody) SetRequestId(v string) *DeletePipelineTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeletePipelineTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePipelineTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePipelineTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelineTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeletePipelineTemplateResponse) SetHeaders(v map[string]*string) *DeletePipelineTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeletePipelineTemplateResponse) SetStatusCode(v int32) *DeletePipelineTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePipelineTemplateResponse) SetBody(v *DeletePipelineTemplateResponseBody) *DeletePipelineTemplateResponse {
	s.Body = v
	return s
}

type DeletePipelineTriggerResponseBody struct {
	// example:
	//
	// A5152937-1C8A-5260-90FA-520CEF028D2D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeletePipelineTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelineTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePipelineTriggerResponseBody) SetRequestId(v string) *DeletePipelineTriggerResponseBody {
	s.RequestId = &v
	return s
}

type DeletePipelineTriggerResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePipelineTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePipelineTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelineTriggerResponse) GoString() string {
	return s.String()
}

func (s *DeletePipelineTriggerResponse) SetHeaders(v map[string]*string) *DeletePipelineTriggerResponse {
	s.Headers = v
	return s
}

func (s *DeletePipelineTriggerResponse) SetStatusCode(v int32) *DeletePipelineTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePipelineTriggerResponse) SetBody(v *DeletePipelineTriggerResponseBody) *DeletePipelineTriggerResponse {
	s.Body = v
	return s
}

type DeletePipelineTriggerEventResponseBody struct {
	// example:
	//
	// BD835E20-EA35-5EE9-A38E-15F9E4AC0B73
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeletePipelineTriggerEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelineTriggerEventResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePipelineTriggerEventResponseBody) SetRequestId(v string) *DeletePipelineTriggerEventResponseBody {
	s.RequestId = &v
	return s
}

type DeletePipelineTriggerEventResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePipelineTriggerEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePipelineTriggerEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelineTriggerEventResponse) GoString() string {
	return s.String()
}

func (s *DeletePipelineTriggerEventResponse) SetHeaders(v map[string]*string) *DeletePipelineTriggerEventResponse {
	s.Headers = v
	return s
}

func (s *DeletePipelineTriggerEventResponse) SetStatusCode(v int32) *DeletePipelineTriggerEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePipelineTriggerEventResponse) SetBody(v *DeletePipelineTriggerEventResponseBody) *DeletePipelineTriggerEventResponse {
	s.Body = v
	return s
}

type DeleteProjectRequest struct {
	// example:
	//
	// true
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) SetForce(v bool) *DeleteProjectRequest {
	s.Force = &v
	return s
}

type DeleteProjectResponseBody struct {
	// example:
	//
	// D9A4CC0F-132B-5EDC-B252-5E11ADFA4B4E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponseBody) SetRequestId(v string) *DeleteProjectResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponse) SetHeaders(v map[string]*string) *DeleteProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectResponse) SetStatusCode(v int32) *DeleteProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectResponse) SetBody(v *DeleteProjectResponseBody) *DeleteProjectResponse {
	s.Body = v
	return s
}

type DeleteRepositoryResponseBody struct {
	// example:
	//
	// 1EEC6F09-A0DA-5A0E-9C3A-DA90B4346B9A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryResponseBody) SetRequestId(v string) *DeleteRepositoryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRepositoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryResponse) SetHeaders(v map[string]*string) *DeleteRepositoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepositoryResponse) SetStatusCode(v int32) *DeleteRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepositoryResponse) SetBody(v *DeleteRepositoryResponseBody) *DeleteRepositoryResponse {
	s.Body = v
	return s
}

type DeleteTaskTemplateResponseBody struct {
	// example:
	//
	// C08FBF2B-9F8E-5415-9EB1-3DC741805C29
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteTaskTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTaskTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTaskTemplateResponseBody) SetRequestId(v string) *DeleteTaskTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTaskTemplateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTaskTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTaskTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTaskTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteTaskTemplateResponse) SetHeaders(v map[string]*string) *DeleteTaskTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteTaskTemplateResponse) SetStatusCode(v int32) *DeleteTaskTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTaskTemplateResponse) SetBody(v *DeleteTaskTemplateResponseBody) *DeleteTaskTemplateResponse {
	s.Body = v
	return s
}

type FetchConnectionCredentialResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OAuthCredential   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchConnectionCredentialResponse) String() string {
	return tea.Prettify(s)
}

func (s FetchConnectionCredentialResponse) GoString() string {
	return s.String()
}

func (s *FetchConnectionCredentialResponse) SetHeaders(v map[string]*string) *FetchConnectionCredentialResponse {
	s.Headers = v
	return s
}

func (s *FetchConnectionCredentialResponse) SetStatusCode(v int32) *FetchConnectionCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchConnectionCredentialResponse) SetBody(v *OAuthCredential) *FetchConnectionCredentialResponse {
	s.Body = v
	return s
}

type FetchRepositoryCheckoutRequest struct {
	// example:
	//
	// main
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// 8828d0087db4210bb1bfeadba90ae52f2938431d
	Commit *string `json:"commit,omitempty" xml:"commit,omitempty"`
	// example:
	//
	// v1.31.0-alpha.0
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s FetchRepositoryCheckoutRequest) String() string {
	return tea.Prettify(s)
}

func (s FetchRepositoryCheckoutRequest) GoString() string {
	return s.String()
}

func (s *FetchRepositoryCheckoutRequest) SetBranch(v string) *FetchRepositoryCheckoutRequest {
	s.Branch = &v
	return s
}

func (s *FetchRepositoryCheckoutRequest) SetCommit(v string) *FetchRepositoryCheckoutRequest {
	s.Commit = &v
	return s
}

func (s *FetchRepositoryCheckoutRequest) SetTag(v string) *FetchRepositoryCheckoutRequest {
	s.Tag = &v
	return s
}

type FetchRepositoryCheckoutResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Checkout          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchRepositoryCheckoutResponse) String() string {
	return tea.Prettify(s)
}

func (s FetchRepositoryCheckoutResponse) GoString() string {
	return s.String()
}

func (s *FetchRepositoryCheckoutResponse) SetHeaders(v map[string]*string) *FetchRepositoryCheckoutResponse {
	s.Headers = v
	return s
}

func (s *FetchRepositoryCheckoutResponse) SetStatusCode(v int32) *FetchRepositoryCheckoutResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchRepositoryCheckoutResponse) SetBody(v *Checkout) *FetchRepositoryCheckoutResponse {
	s.Body = v
	return s
}

type GetConnectionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Connection        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionResponse) GoString() string {
	return s.String()
}

func (s *GetConnectionResponse) SetHeaders(v map[string]*string) *GetConnectionResponse {
	s.Headers = v
	return s
}

func (s *GetConnectionResponse) SetStatusCode(v int32) *GetConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConnectionResponse) SetBody(v *Connection) *GetConnectionResponse {
	s.Body = v
	return s
}

type GetEnvironmentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Environment       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponse) SetHeaders(v map[string]*string) *GetEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *GetEnvironmentResponse) SetStatusCode(v int32) *GetEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnvironmentResponse) SetBody(v *Environment) *GetEnvironmentResponse {
	s.Body = v
	return s
}

type GetPipelineResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Pipeline          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPipelineResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineResponse) SetHeaders(v map[string]*string) *GetPipelineResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineResponse) SetStatusCode(v int32) *GetPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineResponse) SetBody(v *Pipeline) *GetPipelineResponse {
	s.Body = v
	return s
}

type GetPipelineTemplateResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PipelineTemplate  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPipelineTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineTemplateResponse) SetHeaders(v map[string]*string) *GetPipelineTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineTemplateResponse) SetStatusCode(v int32) *GetPipelineTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineTemplateResponse) SetBody(v *PipelineTemplate) *GetPipelineTemplateResponse {
	s.Body = v
	return s
}

type GetPipelineTriggerResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PipelineTrigger   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPipelineTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineTriggerResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineTriggerResponse) SetHeaders(v map[string]*string) *GetPipelineTriggerResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineTriggerResponse) SetStatusCode(v int32) *GetPipelineTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineTriggerResponse) SetBody(v *PipelineTrigger) *GetPipelineTriggerResponse {
	s.Body = v
	return s
}

type GetPipelineTriggerEventResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PipelineTriggerEvent `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPipelineTriggerEventResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineTriggerEventResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineTriggerEventResponse) SetHeaders(v map[string]*string) *GetPipelineTriggerEventResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineTriggerEventResponse) SetStatusCode(v int32) *GetPipelineTriggerEventResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineTriggerEventResponse) SetBody(v *PipelineTriggerEvent) *GetPipelineTriggerEventResponse {
	s.Body = v
	return s
}

type GetProjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Project           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponse) GoString() string {
	return s.String()
}

func (s *GetProjectResponse) SetHeaders(v map[string]*string) *GetProjectResponse {
	s.Headers = v
	return s
}

func (s *GetProjectResponse) SetStatusCode(v int32) *GetProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectResponse) SetBody(v *Project) *GetProjectResponse {
	s.Body = v
	return s
}

type GetRepositoryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Repository        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryResponse) GoString() string {
	return s.String()
}

func (s *GetRepositoryResponse) SetHeaders(v map[string]*string) *GetRepositoryResponse {
	s.Headers = v
	return s
}

func (s *GetRepositoryResponse) SetStatusCode(v int32) *GetRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepositoryResponse) SetBody(v *Repository) *GetRepositoryResponse {
	s.Body = v
	return s
}

type GetTaskResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Task              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTaskResponse) SetHeaders(v map[string]*string) *GetTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTaskResponse) SetStatusCode(v int32) *GetTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskResponse) SetBody(v *Task) *GetTaskResponse {
	s.Body = v
	return s
}

type GetTaskTemplateResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskTemplate      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetTaskTemplateResponse) SetHeaders(v map[string]*string) *GetTaskTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetTaskTemplateResponse) SetStatusCode(v int32) *GetTaskTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskTemplateResponse) SetBody(v *TaskTemplate) *GetTaskTemplateResponse {
	s.Body = v
	return s
}

type ListConnectionsRequest struct {
	// example:
	//
	// auto-
	Keyword       *string   `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LabelSelector []*string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListConnectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsRequest) GoString() string {
	return s.String()
}

func (s *ListConnectionsRequest) SetKeyword(v string) *ListConnectionsRequest {
	s.Keyword = &v
	return s
}

func (s *ListConnectionsRequest) SetLabelSelector(v []*string) *ListConnectionsRequest {
	s.LabelSelector = v
	return s
}

func (s *ListConnectionsRequest) SetPageNumber(v int64) *ListConnectionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConnectionsRequest) SetPageSize(v int64) *ListConnectionsRequest {
	s.PageSize = &v
	return s
}

type ListConnectionsShrinkRequest struct {
	// example:
	//
	// auto-
	Keyword             *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LabelSelectorShrink *string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListConnectionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListConnectionsShrinkRequest) SetKeyword(v string) *ListConnectionsShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListConnectionsShrinkRequest) SetLabelSelectorShrink(v string) *ListConnectionsShrinkRequest {
	s.LabelSelectorShrink = &v
	return s
}

func (s *ListConnectionsShrinkRequest) SetPageNumber(v int64) *ListConnectionsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConnectionsShrinkRequest) SetPageSize(v int64) *ListConnectionsShrinkRequest {
	s.PageSize = &v
	return s
}

type ListConnectionsResponseBody struct {
	Data []*Connection `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListConnectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBody) SetData(v []*Connection) *ListConnectionsResponseBody {
	s.Data = v
	return s
}

func (s *ListConnectionsResponseBody) SetPageNumber(v int64) *ListConnectionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListConnectionsResponseBody) SetPageSize(v int64) *ListConnectionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListConnectionsResponseBody) SetTotalCount(v int64) *ListConnectionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListConnectionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConnectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsResponse) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponse) SetHeaders(v map[string]*string) *ListConnectionsResponse {
	s.Headers = v
	return s
}

func (s *ListConnectionsResponse) SetStatusCode(v int32) *ListConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConnectionsResponse) SetBody(v *ListConnectionsResponseBody) *ListConnectionsResponse {
	s.Body = v
	return s
}

type ListEnvironmentsRequest struct {
	// example:
	//
	// dev
	Keyword       *string   `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LabelSelector []*string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListEnvironmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsRequest) SetKeyword(v string) *ListEnvironmentsRequest {
	s.Keyword = &v
	return s
}

func (s *ListEnvironmentsRequest) SetLabelSelector(v []*string) *ListEnvironmentsRequest {
	s.LabelSelector = v
	return s
}

func (s *ListEnvironmentsRequest) SetPageNumber(v int64) *ListEnvironmentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEnvironmentsRequest) SetPageSize(v int64) *ListEnvironmentsRequest {
	s.PageSize = &v
	return s
}

type ListEnvironmentsShrinkRequest struct {
	// example:
	//
	// dev
	Keyword             *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LabelSelectorShrink *string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListEnvironmentsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsShrinkRequest) SetKeyword(v string) *ListEnvironmentsShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListEnvironmentsShrinkRequest) SetLabelSelectorShrink(v string) *ListEnvironmentsShrinkRequest {
	s.LabelSelectorShrink = &v
	return s
}

func (s *ListEnvironmentsShrinkRequest) SetPageNumber(v int64) *ListEnvironmentsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEnvironmentsShrinkRequest) SetPageSize(v int64) *ListEnvironmentsShrinkRequest {
	s.PageSize = &v
	return s
}

type ListEnvironmentsResponseBody struct {
	Data []*Environment `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListEnvironmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBody) SetData(v []*Environment) *ListEnvironmentsResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentsResponseBody) SetPageNumber(v int64) *ListEnvironmentsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEnvironmentsResponseBody) SetPageSize(v int64) *ListEnvironmentsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEnvironmentsResponseBody) SetTotalCount(v int64) *ListEnvironmentsResponseBody {
	s.TotalCount = &v
	return s
}

type ListEnvironmentsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnvironmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnvironmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponse) SetHeaders(v map[string]*string) *ListEnvironmentsResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentsResponse) SetStatusCode(v int32) *ListEnvironmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnvironmentsResponse) SetBody(v *ListEnvironmentsResponseBody) *ListEnvironmentsResponse {
	s.Body = v
	return s
}

type ListPipelineTemplatesRequest struct {
	LabelSelector []*string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty" type:"Repeated"`
}

func (s ListPipelineTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineTemplatesRequest) SetLabelSelector(v []*string) *ListPipelineTemplatesRequest {
	s.LabelSelector = v
	return s
}

type ListPipelineTemplatesShrinkRequest struct {
	LabelSelectorShrink *string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty"`
}

func (s ListPipelineTemplatesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineTemplatesShrinkRequest) SetLabelSelectorShrink(v string) *ListPipelineTemplatesShrinkRequest {
	s.LabelSelectorShrink = &v
	return s
}

type ListPipelineTemplatesResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*PipelineTemplate `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ListPipelineTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineTemplatesResponse) SetHeaders(v map[string]*string) *ListPipelineTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineTemplatesResponse) SetStatusCode(v int32) *ListPipelineTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineTemplatesResponse) SetBody(v []*PipelineTemplate) *ListPipelineTemplatesResponse {
	s.Body = v
	return s
}

type ListPipelineTriggerEventsRequest struct {
	// example:
	//
	// demo
	Keyword       *string   `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LabelSelector []*string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListPipelineTriggerEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineTriggerEventsRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineTriggerEventsRequest) SetKeyword(v string) *ListPipelineTriggerEventsRequest {
	s.Keyword = &v
	return s
}

func (s *ListPipelineTriggerEventsRequest) SetLabelSelector(v []*string) *ListPipelineTriggerEventsRequest {
	s.LabelSelector = v
	return s
}

func (s *ListPipelineTriggerEventsRequest) SetPageNumber(v int64) *ListPipelineTriggerEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPipelineTriggerEventsRequest) SetPageSize(v int64) *ListPipelineTriggerEventsRequest {
	s.PageSize = &v
	return s
}

type ListPipelineTriggerEventsShrinkRequest struct {
	// example:
	//
	// demo
	Keyword             *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LabelSelectorShrink *string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListPipelineTriggerEventsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineTriggerEventsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineTriggerEventsShrinkRequest) SetKeyword(v string) *ListPipelineTriggerEventsShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListPipelineTriggerEventsShrinkRequest) SetLabelSelectorShrink(v string) *ListPipelineTriggerEventsShrinkRequest {
	s.LabelSelectorShrink = &v
	return s
}

func (s *ListPipelineTriggerEventsShrinkRequest) SetPageNumber(v int64) *ListPipelineTriggerEventsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPipelineTriggerEventsShrinkRequest) SetPageSize(v int64) *ListPipelineTriggerEventsShrinkRequest {
	s.PageSize = &v
	return s
}

type ListPipelineTriggerEventsResponseBody struct {
	Data []*PipelineTriggerEvent `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPipelineTriggerEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineTriggerEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineTriggerEventsResponseBody) SetData(v []*PipelineTriggerEvent) *ListPipelineTriggerEventsResponseBody {
	s.Data = v
	return s
}

func (s *ListPipelineTriggerEventsResponseBody) SetPageNumber(v int64) *ListPipelineTriggerEventsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPipelineTriggerEventsResponseBody) SetPageSize(v int64) *ListPipelineTriggerEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPipelineTriggerEventsResponseBody) SetTotalCount(v int64) *ListPipelineTriggerEventsResponseBody {
	s.TotalCount = &v
	return s
}

type ListPipelineTriggerEventsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPipelineTriggerEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPipelineTriggerEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineTriggerEventsResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineTriggerEventsResponse) SetHeaders(v map[string]*string) *ListPipelineTriggerEventsResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineTriggerEventsResponse) SetStatusCode(v int32) *ListPipelineTriggerEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineTriggerEventsResponse) SetBody(v *ListPipelineTriggerEventsResponseBody) *ListPipelineTriggerEventsResponse {
	s.Body = v
	return s
}

type ListPipelineTriggersRequest struct {
	// example:
	//
	// demo
	Keyword       *string   `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LabelSelector []*string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListPipelineTriggersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineTriggersRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineTriggersRequest) SetKeyword(v string) *ListPipelineTriggersRequest {
	s.Keyword = &v
	return s
}

func (s *ListPipelineTriggersRequest) SetLabelSelector(v []*string) *ListPipelineTriggersRequest {
	s.LabelSelector = v
	return s
}

func (s *ListPipelineTriggersRequest) SetPageNumber(v int64) *ListPipelineTriggersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPipelineTriggersRequest) SetPageSize(v int64) *ListPipelineTriggersRequest {
	s.PageSize = &v
	return s
}

type ListPipelineTriggersShrinkRequest struct {
	// example:
	//
	// demo
	Keyword             *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LabelSelectorShrink *string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListPipelineTriggersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineTriggersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineTriggersShrinkRequest) SetKeyword(v string) *ListPipelineTriggersShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListPipelineTriggersShrinkRequest) SetLabelSelectorShrink(v string) *ListPipelineTriggersShrinkRequest {
	s.LabelSelectorShrink = &v
	return s
}

func (s *ListPipelineTriggersShrinkRequest) SetPageNumber(v int64) *ListPipelineTriggersShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPipelineTriggersShrinkRequest) SetPageSize(v int64) *ListPipelineTriggersShrinkRequest {
	s.PageSize = &v
	return s
}

type ListPipelineTriggersResponseBody struct {
	Data []*PipelineTrigger `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPipelineTriggersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineTriggersResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineTriggersResponseBody) SetData(v []*PipelineTrigger) *ListPipelineTriggersResponseBody {
	s.Data = v
	return s
}

func (s *ListPipelineTriggersResponseBody) SetPageNumber(v int64) *ListPipelineTriggersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPipelineTriggersResponseBody) SetPageSize(v int64) *ListPipelineTriggersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPipelineTriggersResponseBody) SetTotalCount(v int64) *ListPipelineTriggersResponseBody {
	s.TotalCount = &v
	return s
}

type ListPipelineTriggersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPipelineTriggersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPipelineTriggersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineTriggersResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineTriggersResponse) SetHeaders(v map[string]*string) *ListPipelineTriggersResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineTriggersResponse) SetStatusCode(v int32) *ListPipelineTriggersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineTriggersResponse) SetBody(v *ListPipelineTriggersResponseBody) *ListPipelineTriggersResponse {
	s.Body = v
	return s
}

type ListPipelinesRequest struct {
	LabelSelector []*string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty" type:"Repeated"`
}

func (s ListPipelinesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelinesRequest) GoString() string {
	return s.String()
}

func (s *ListPipelinesRequest) SetLabelSelector(v []*string) *ListPipelinesRequest {
	s.LabelSelector = v
	return s
}

type ListPipelinesShrinkRequest struct {
	LabelSelectorShrink *string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty"`
}

func (s ListPipelinesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelinesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPipelinesShrinkRequest) SetLabelSelectorShrink(v string) *ListPipelinesShrinkRequest {
	s.LabelSelectorShrink = &v
	return s
}

type ListPipelinesResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*Pipeline        `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ListPipelinesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPipelinesResponse) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponse) SetHeaders(v map[string]*string) *ListPipelinesResponse {
	s.Headers = v
	return s
}

func (s *ListPipelinesResponse) SetStatusCode(v int32) *ListPipelinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelinesResponse) SetBody(v []*Pipeline) *ListPipelinesResponse {
	s.Body = v
	return s
}

type ListProjectsRequest struct {
	// example:
	//
	// spring-boot
	Keyword       *string   `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LabelSelector []*string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) SetKeyword(v string) *ListProjectsRequest {
	s.Keyword = &v
	return s
}

func (s *ListProjectsRequest) SetLabelSelector(v []*string) *ListProjectsRequest {
	s.LabelSelector = v
	return s
}

func (s *ListProjectsRequest) SetPageNumber(v int64) *ListProjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsRequest) SetPageSize(v int64) *ListProjectsRequest {
	s.PageSize = &v
	return s
}

type ListProjectsShrinkRequest struct {
	// example:
	//
	// spring-boot
	Keyword             *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LabelSelectorShrink *string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListProjectsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsShrinkRequest) SetKeyword(v string) *ListProjectsShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetLabelSelectorShrink(v string) *ListProjectsShrinkRequest {
	s.LabelSelectorShrink = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetPageNumber(v int64) *ListProjectsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetPageSize(v int64) *ListProjectsShrinkRequest {
	s.PageSize = &v
	return s
}

type ListProjectsResponseBody struct {
	Data []*Project `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) SetData(v []*Project) *ListProjectsResponseBody {
	s.Data = v
	return s
}

func (s *ListProjectsResponseBody) SetPageNumber(v int64) *ListProjectsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsResponseBody) SetPageSize(v int64) *ListProjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProjectsResponseBody) SetTotalCount(v int64) *ListProjectsResponseBody {
	s.TotalCount = &v
	return s
}

type ListProjectsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsResponse) SetHeaders(v map[string]*string) *ListProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsResponse) SetStatusCode(v int32) *ListProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectsResponse) SetBody(v *ListProjectsResponseBody) *ListProjectsResponse {
	s.Body = v
	return s
}

type ListRepositoriesRequest struct {
	// example:
	//
	// repo-start-springboot
	Keyword       *string   `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LabelSelector []*string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListRepositoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoriesRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoriesRequest) SetKeyword(v string) *ListRepositoriesRequest {
	s.Keyword = &v
	return s
}

func (s *ListRepositoriesRequest) SetLabelSelector(v []*string) *ListRepositoriesRequest {
	s.LabelSelector = v
	return s
}

func (s *ListRepositoriesRequest) SetPageNumber(v int64) *ListRepositoriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRepositoriesRequest) SetPageSize(v int64) *ListRepositoriesRequest {
	s.PageSize = &v
	return s
}

type ListRepositoriesShrinkRequest struct {
	// example:
	//
	// repo-start-springboot
	Keyword             *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LabelSelectorShrink *string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListRepositoriesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoriesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoriesShrinkRequest) SetKeyword(v string) *ListRepositoriesShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListRepositoriesShrinkRequest) SetLabelSelectorShrink(v string) *ListRepositoriesShrinkRequest {
	s.LabelSelectorShrink = &v
	return s
}

func (s *ListRepositoriesShrinkRequest) SetPageNumber(v int64) *ListRepositoriesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRepositoriesShrinkRequest) SetPageSize(v int64) *ListRepositoriesShrinkRequest {
	s.PageSize = &v
	return s
}

type ListRepositoriesResponseBody struct {
	Data []*Repository `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListRepositoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoriesResponseBody) SetData(v []*Repository) *ListRepositoriesResponseBody {
	s.Data = v
	return s
}

func (s *ListRepositoriesResponseBody) SetPageNumber(v int64) *ListRepositoriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRepositoriesResponseBody) SetPageSize(v int64) *ListRepositoriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRepositoriesResponseBody) SetTotalCount(v int64) *ListRepositoriesResponseBody {
	s.TotalCount = &v
	return s
}

type ListRepositoriesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepositoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepositoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoriesResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoriesResponse) SetHeaders(v map[string]*string) *ListRepositoriesResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoriesResponse) SetStatusCode(v int32) *ListRepositoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoriesResponse) SetBody(v *ListRepositoriesResponseBody) *ListRepositoriesResponse {
	s.Body = v
	return s
}

type ListTaskTemplatesRequest struct {
	LabelSelector []*string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty" type:"Repeated"`
}

func (s ListTaskTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTaskTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTaskTemplatesRequest) SetLabelSelector(v []*string) *ListTaskTemplatesRequest {
	s.LabelSelector = v
	return s
}

type ListTaskTemplatesShrinkRequest struct {
	LabelSelectorShrink *string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty"`
}

func (s ListTaskTemplatesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTaskTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTaskTemplatesShrinkRequest) SetLabelSelectorShrink(v string) *ListTaskTemplatesShrinkRequest {
	s.LabelSelectorShrink = &v
	return s
}

type ListTaskTemplatesResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*TaskTemplate    `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ListTaskTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTaskTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListTaskTemplatesResponse) SetHeaders(v map[string]*string) *ListTaskTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListTaskTemplatesResponse) SetStatusCode(v int32) *ListTaskTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskTemplatesResponse) SetBody(v []*TaskTemplate) *ListTaskTemplatesResponse {
	s.Body = v
	return s
}

type ListTasksRequest struct {
	LabelSelector []*string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty" type:"Repeated"`
}

func (s ListTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) SetLabelSelector(v []*string) *ListTasksRequest {
	s.LabelSelector = v
	return s
}

type ListTasksShrinkRequest struct {
	LabelSelectorShrink *string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty"`
}

func (s ListTasksShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTasksShrinkRequest) SetLabelSelectorShrink(v string) *ListTasksShrinkRequest {
	s.LabelSelectorShrink = &v
	return s
}

type ListTasksResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*Task            `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ListTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTasksResponse) SetHeaders(v map[string]*string) *ListTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTasksResponse) SetStatusCode(v int32) *ListTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTasksResponse) SetBody(v []*Task) *ListTasksResponse {
	s.Body = v
	return s
}

type PutEnvironmentRequest struct {
	Body *Environment `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// false
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s PutEnvironmentRequest) String() string {
	return tea.Prettify(s)
}

func (s PutEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *PutEnvironmentRequest) SetBody(v *Environment) *PutEnvironmentRequest {
	s.Body = v
	return s
}

func (s *PutEnvironmentRequest) SetForce(v bool) *PutEnvironmentRequest {
	s.Force = &v
	return s
}

type PutEnvironmentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Environment       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s PutEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *PutEnvironmentResponse) SetHeaders(v map[string]*string) *PutEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *PutEnvironmentResponse) SetStatusCode(v int32) *PutEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *PutEnvironmentResponse) SetBody(v *Environment) *PutEnvironmentResponse {
	s.Body = v
	return s
}

type PutPipelineStatusRequest struct {
	Body *Pipeline `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// false
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s PutPipelineStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s PutPipelineStatusRequest) GoString() string {
	return s.String()
}

func (s *PutPipelineStatusRequest) SetBody(v *Pipeline) *PutPipelineStatusRequest {
	s.Body = v
	return s
}

func (s *PutPipelineStatusRequest) SetForce(v bool) *PutPipelineStatusRequest {
	s.Force = &v
	return s
}

type PutPipelineStatusResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Pipeline          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutPipelineStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s PutPipelineStatusResponse) GoString() string {
	return s.String()
}

func (s *PutPipelineStatusResponse) SetHeaders(v map[string]*string) *PutPipelineStatusResponse {
	s.Headers = v
	return s
}

func (s *PutPipelineStatusResponse) SetStatusCode(v int32) *PutPipelineStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *PutPipelineStatusResponse) SetBody(v *Pipeline) *PutPipelineStatusResponse {
	s.Body = v
	return s
}

type PutPipelineTemplateRequest struct {
	Body *PipelineTemplate `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// false
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s PutPipelineTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s PutPipelineTemplateRequest) GoString() string {
	return s.String()
}

func (s *PutPipelineTemplateRequest) SetBody(v *PipelineTemplate) *PutPipelineTemplateRequest {
	s.Body = v
	return s
}

func (s *PutPipelineTemplateRequest) SetForce(v bool) *PutPipelineTemplateRequest {
	s.Force = &v
	return s
}

type PutPipelineTemplateResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PipelineTemplate  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutPipelineTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s PutPipelineTemplateResponse) GoString() string {
	return s.String()
}

func (s *PutPipelineTemplateResponse) SetHeaders(v map[string]*string) *PutPipelineTemplateResponse {
	s.Headers = v
	return s
}

func (s *PutPipelineTemplateResponse) SetStatusCode(v int32) *PutPipelineTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *PutPipelineTemplateResponse) SetBody(v *PipelineTemplate) *PutPipelineTemplateResponse {
	s.Body = v
	return s
}

type PutPipelineTriggerRequest struct {
	Body *PipelineTrigger `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// false
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s PutPipelineTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s PutPipelineTriggerRequest) GoString() string {
	return s.String()
}

func (s *PutPipelineTriggerRequest) SetBody(v *PipelineTrigger) *PutPipelineTriggerRequest {
	s.Body = v
	return s
}

func (s *PutPipelineTriggerRequest) SetForce(v bool) *PutPipelineTriggerRequest {
	s.Force = &v
	return s
}

type PutPipelineTriggerResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PipelineTrigger   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutPipelineTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s PutPipelineTriggerResponse) GoString() string {
	return s.String()
}

func (s *PutPipelineTriggerResponse) SetHeaders(v map[string]*string) *PutPipelineTriggerResponse {
	s.Headers = v
	return s
}

func (s *PutPipelineTriggerResponse) SetStatusCode(v int32) *PutPipelineTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *PutPipelineTriggerResponse) SetBody(v *PipelineTrigger) *PutPipelineTriggerResponse {
	s.Body = v
	return s
}

type PutProjectRequest struct {
	Body *Project `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// true
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s PutProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PutProjectRequest) GoString() string {
	return s.String()
}

func (s *PutProjectRequest) SetBody(v *Project) *PutProjectRequest {
	s.Body = v
	return s
}

func (s *PutProjectRequest) SetForce(v bool) *PutProjectRequest {
	s.Force = &v
	return s
}

type PutProjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Project           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PutProjectResponse) GoString() string {
	return s.String()
}

func (s *PutProjectResponse) SetHeaders(v map[string]*string) *PutProjectResponse {
	s.Headers = v
	return s
}

func (s *PutProjectResponse) SetStatusCode(v int32) *PutProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PutProjectResponse) SetBody(v *Project) *PutProjectResponse {
	s.Body = v
	return s
}

type PutTaskStatusRequest struct {
	Body *Task `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// false
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s PutTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s PutTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *PutTaskStatusRequest) SetBody(v *Task) *PutTaskStatusRequest {
	s.Body = v
	return s
}

func (s *PutTaskStatusRequest) SetForce(v bool) *PutTaskStatusRequest {
	s.Force = &v
	return s
}

type PutTaskStatusResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Task              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s PutTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *PutTaskStatusResponse) SetHeaders(v map[string]*string) *PutTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *PutTaskStatusResponse) SetStatusCode(v int32) *PutTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *PutTaskStatusResponse) SetBody(v *Task) *PutTaskStatusResponse {
	s.Body = v
	return s
}

type PutTaskTemplateRequest struct {
	Body *TaskTemplate `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// false
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s PutTaskTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s PutTaskTemplateRequest) GoString() string {
	return s.String()
}

func (s *PutTaskTemplateRequest) SetBody(v *TaskTemplate) *PutTaskTemplateRequest {
	s.Body = v
	return s
}

func (s *PutTaskTemplateRequest) SetForce(v bool) *PutTaskTemplateRequest {
	s.Force = &v
	return s
}

type PutTaskTemplateResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskTemplate      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutTaskTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s PutTaskTemplateResponse) GoString() string {
	return s.String()
}

func (s *PutTaskTemplateResponse) SetHeaders(v map[string]*string) *PutTaskTemplateResponse {
	s.Headers = v
	return s
}

func (s *PutTaskTemplateResponse) SetStatusCode(v int32) *PutTaskTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *PutTaskTemplateResponse) SetBody(v *TaskTemplate) *PutTaskTemplateResponse {
	s.Body = v
	return s
}

type RefreshConnectionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Connection        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshConnectionResponse) GoString() string {
	return s.String()
}

func (s *RefreshConnectionResponse) SetHeaders(v map[string]*string) *RefreshConnectionResponse {
	s.Headers = v
	return s
}

func (s *RefreshConnectionResponse) SetStatusCode(v int32) *RefreshConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshConnectionResponse) SetBody(v *Connection) *RefreshConnectionResponse {
	s.Body = v
	return s
}

type ResumeTaskResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Task              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeTaskResponse) GoString() string {
	return s.String()
}

func (s *ResumeTaskResponse) SetHeaders(v map[string]*string) *ResumeTaskResponse {
	s.Headers = v
	return s
}

func (s *ResumeTaskResponse) SetStatusCode(v int32) *ResumeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeTaskResponse) SetBody(v *Task) *ResumeTaskResponse {
	s.Body = v
	return s
}

type RetryTaskResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Task              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s RetryTaskResponse) GoString() string {
	return s.String()
}

func (s *RetryTaskResponse) SetHeaders(v map[string]*string) *RetryTaskResponse {
	s.Headers = v
	return s
}

func (s *RetryTaskResponse) SetStatusCode(v int32) *RetryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryTaskResponse) SetBody(v *Task) *RetryTaskResponse {
	s.Body = v
	return s
}

type StartPipelineResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Pipeline          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartPipelineResponse) String() string {
	return tea.Prettify(s)
}

func (s StartPipelineResponse) GoString() string {
	return s.String()
}

func (s *StartPipelineResponse) SetHeaders(v map[string]*string) *StartPipelineResponse {
	s.Headers = v
	return s
}

func (s *StartPipelineResponse) SetStatusCode(v int32) *StartPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *StartPipelineResponse) SetBody(v *Pipeline) *StartPipelineResponse {
	s.Body = v
	return s
}

type StartTaskResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Task              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StartTaskResponse) GoString() string {
	return s.String()
}

func (s *StartTaskResponse) SetHeaders(v map[string]*string) *StartTaskResponse {
	s.Headers = v
	return s
}

func (s *StartTaskResponse) SetStatusCode(v int32) *StartTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartTaskResponse) SetBody(v *Task) *StartTaskResponse {
	s.Body = v
	return s
}

type UpdateEnvironmentRequest struct {
	Body *Environment `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEnvironmentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentRequest) SetBody(v *Environment) *UpdateEnvironmentRequest {
	s.Body = v
	return s
}

type UpdateEnvironmentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Environment       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentResponse) SetHeaders(v map[string]*string) *UpdateEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnvironmentResponse) SetStatusCode(v int32) *UpdateEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEnvironmentResponse) SetBody(v *Environment) *UpdateEnvironmentResponse {
	s.Body = v
	return s
}

type UpdatePipelineTriggerRequest struct {
	Body *PipelineTrigger `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePipelineTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineTriggerRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelineTriggerRequest) SetBody(v *PipelineTrigger) *UpdatePipelineTriggerRequest {
	s.Body = v
	return s
}

type UpdatePipelineTriggerResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PipelineTrigger   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePipelineTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineTriggerResponse) GoString() string {
	return s.String()
}

func (s *UpdatePipelineTriggerResponse) SetHeaders(v map[string]*string) *UpdatePipelineTriggerResponse {
	s.Headers = v
	return s
}

func (s *UpdatePipelineTriggerResponse) SetStatusCode(v int32) *UpdatePipelineTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePipelineTriggerResponse) SetBody(v *PipelineTrigger) *UpdatePipelineTriggerResponse {
	s.Body = v
	return s
}

type UpdateProjectRequest struct {
	Body *Project `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// true
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) SetBody(v *Project) *UpdateProjectRequest {
	s.Body = v
	return s
}

func (s *UpdateProjectRequest) SetForce(v bool) *UpdateProjectRequest {
	s.Force = &v
	return s
}

type UpdateProjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Project           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponse) SetHeaders(v map[string]*string) *UpdateProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectResponse) SetStatusCode(v int32) *UpdateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectResponse) SetBody(v *Project) *UpdateProjectResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("devs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 激活身份绑定,完成OAuth授权
//
// @param request - ActivateConnectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateConnectionResponse
func (client *Client) ActivateConnectionWithOptions(name *string, request *ActivateConnectionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ActivateConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		body["account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.Credential)) {
		body["credential"] = request.Credential
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ActivateConnection"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/connections/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/activate"),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ActivateConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 激活身份绑定,完成OAuth授权
//
// @param request - ActivateConnectionRequest
//
// @return ActivateConnectionResponse
func (client *Client) ActivateConnection(name *string, request *ActivateConnectionRequest) (_result *ActivateConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ActivateConnectionResponse{}
	_body, _err := client.ActivateConnectionWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消流水线
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelPipelineResponse
func (client *Client) CancelPipelineWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelPipelineResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CancelPipeline"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/cancel"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelPipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消流水线
//
// @return CancelPipelineResponse
func (client *Client) CancelPipeline(name *string) (_result *CancelPipelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelPipelineResponse{}
	_body, _err := client.CancelPipelineWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelTaskResponse
func (client *Client) CancelTaskWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CancelTask"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/cancel"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消任务
//
// @return CancelTaskResponse
func (client *Client) CancelTask(name *string) (_result *CancelTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelTaskResponse{}
	_body, _err := client.CancelTaskWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建身份绑定
//
// @param request - CreateConnectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConnectionResponse
func (client *Client) CreateConnectionWithOptions(request *CreateConnectionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateConnection"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/connections"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建身份绑定
//
// @param request - CreateConnectionRequest
//
// @return CreateConnectionResponse
func (client *Client) CreateConnection(request *CreateConnectionRequest) (_result *CreateConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateConnectionResponse{}
	_body, _err := client.CreateConnectionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建环境
//
// @param request - CreateEnvironmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnvironmentResponse
func (client *Client) CreateEnvironmentWithOptions(project *string, request *CreateEnvironmentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateEnvironmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEnvironment"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/projects/" + tea.StringValue(openapiutil.GetEncodeParam(project)) + "/environments"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建环境
//
// @param request - CreateEnvironmentRequest
//
// @return CreateEnvironmentResponse
func (client *Client) CreateEnvironment(project *string, request *CreateEnvironmentRequest) (_result *CreateEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEnvironmentResponse{}
	_body, _err := client.CreateEnvironmentWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建流水线
//
// @param request - CreatePipelineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelineResponse
func (client *Client) CreatePipelineWithOptions(request *CreatePipelineRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePipelineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePipeline"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelines"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建流水线
//
// @param request - CreatePipelineRequest
//
// @return CreatePipelineResponse
func (client *Client) CreatePipeline(request *CreatePipelineRequest) (_result *CreatePipelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePipelineResponse{}
	_body, _err := client.CreatePipelineWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建流水线模板
//
// @param request - CreatePipelineTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelineTemplateResponse
func (client *Client) CreatePipelineTemplateWithOptions(request *CreatePipelineTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePipelineTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePipelineTemplate"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelinetemplates"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePipelineTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建流水线模板
//
// @param request - CreatePipelineTemplateRequest
//
// @return CreatePipelineTemplateResponse
func (client *Client) CreatePipelineTemplate(request *CreatePipelineTemplateRequest) (_result *CreatePipelineTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePipelineTemplateResponse{}
	_body, _err := client.CreatePipelineTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建流水线触发器
//
// @param request - CreatePipelineTriggerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelineTriggerResponse
func (client *Client) CreatePipelineTriggerWithOptions(request *CreatePipelineTriggerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePipelineTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePipelineTrigger"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelinetriggers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePipelineTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建流水线触发器
//
// @param request - CreatePipelineTriggerRequest
//
// @return CreatePipelineTriggerResponse
func (client *Client) CreatePipelineTrigger(request *CreatePipelineTriggerRequest) (_result *CreatePipelineTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePipelineTriggerResponse{}
	_body, _err := client.CreatePipelineTriggerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建流水线触发历史
//
// @param request - CreatePipelineTriggerEventRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelineTriggerEventResponse
func (client *Client) CreatePipelineTriggerEventWithOptions(request *CreatePipelineTriggerEventRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePipelineTriggerEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePipelineTriggerEvent"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelinetriggerevents"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePipelineTriggerEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建流水线触发历史
//
// @param request - CreatePipelineTriggerEventRequest
//
// @return CreatePipelineTriggerEventResponse
func (client *Client) CreatePipelineTriggerEvent(request *CreatePipelineTriggerEventRequest) (_result *CreatePipelineTriggerEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePipelineTriggerEventResponse{}
	_body, _err := client.CreatePipelineTriggerEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建项目
//
// @param request - CreateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProject"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/projects"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建项目
//
// @param request - CreateProjectRequest
//
// @return CreateProjectResponse
func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建仓库绑定
//
// @param request - CreateRepositoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepositoryResponse
func (client *Client) CreateRepositoryWithOptions(request *CreateRepositoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRepositoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRepository"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/repositories"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建仓库绑定
//
// @param request - CreateRepositoryRequest
//
// @return CreateRepositoryResponse
func (client *Client) CreateRepository(request *CreateRepositoryRequest) (_result *CreateRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRepositoryResponse{}
	_body, _err := client.CreateRepositoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建任务
//
// @param request - CreateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTaskResponse
func (client *Client) CreateTaskWithOptions(request *CreateTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTask"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/tasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建任务
//
// @param request - CreateTaskRequest
//
// @return CreateTaskResponse
func (client *Client) CreateTask(request *CreateTaskRequest) (_result *CreateTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTaskResponse{}
	_body, _err := client.CreateTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建任务模板
//
// @param request - CreateTaskTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTaskTemplateResponse
func (client *Client) CreateTaskTemplateWithOptions(request *CreateTaskTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTaskTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTaskTemplate"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/tasktemplates"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTaskTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建任务模板
//
// @param request - CreateTaskTemplateRequest
//
// @return CreateTaskTemplateResponse
func (client *Client) CreateTaskTemplate(request *CreateTaskTemplateRequest) (_result *CreateTaskTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTaskTemplateResponse{}
	_body, _err := client.CreateTaskTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除身份绑定
//
// @param request - DeleteConnectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConnectionResponse
func (client *Client) DeleteConnectionWithOptions(name *string, request *DeleteConnectionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["force"] = request.Force
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConnection"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/connections/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除身份绑定
//
// @param request - DeleteConnectionRequest
//
// @return DeleteConnectionResponse
func (client *Client) DeleteConnection(name *string, request *DeleteConnectionRequest) (_result *DeleteConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConnectionResponse{}
	_body, _err := client.DeleteConnectionWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除环境
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnvironmentResponse
func (client *Client) DeleteEnvironmentWithOptions(project *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteEnvironmentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEnvironment"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/projects/" + tea.StringValue(openapiutil.GetEncodeParam(project)) + "/environments/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除环境
//
// @return DeleteEnvironmentResponse
func (client *Client) DeleteEnvironment(project *string, name *string) (_result *DeleteEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEnvironmentResponse{}
	_body, _err := client.DeleteEnvironmentWithOptions(project, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除流水线模板
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePipelineTemplateResponse
func (client *Client) DeletePipelineTemplateWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeletePipelineTemplateResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePipelineTemplate"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelinetemplates/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePipelineTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除流水线模板
//
// @return DeletePipelineTemplateResponse
func (client *Client) DeletePipelineTemplate(name *string) (_result *DeletePipelineTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePipelineTemplateResponse{}
	_body, _err := client.DeletePipelineTemplateWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除流水线触发器
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePipelineTriggerResponse
func (client *Client) DeletePipelineTriggerWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeletePipelineTriggerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePipelineTrigger"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelinetriggers/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePipelineTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除流水线触发器
//
// @return DeletePipelineTriggerResponse
func (client *Client) DeletePipelineTrigger(name *string) (_result *DeletePipelineTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePipelineTriggerResponse{}
	_body, _err := client.DeletePipelineTriggerWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除流水线触发历史
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePipelineTriggerEventResponse
func (client *Client) DeletePipelineTriggerEventWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeletePipelineTriggerEventResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePipelineTriggerEvent"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelinetriggerevents/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePipelineTriggerEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除流水线触发历史
//
// @return DeletePipelineTriggerEventResponse
func (client *Client) DeletePipelineTriggerEvent(name *string) (_result *DeletePipelineTriggerEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePipelineTriggerEventResponse{}
	_body, _err := client.DeletePipelineTriggerEventWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除项目
//
// @param request - DeleteProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectResponse
func (client *Client) DeleteProjectWithOptions(name *string, request *DeleteProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["force"] = request.Force
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProject"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/projects/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除项目
//
// @param request - DeleteProjectRequest
//
// @return DeleteProjectResponse
func (client *Client) DeleteProject(name *string, request *DeleteProjectRequest) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除仓库绑定
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRepositoryResponse
func (client *Client) DeleteRepositoryWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteRepositoryResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRepository"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/repositories/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除仓库绑定
//
// @return DeleteRepositoryResponse
func (client *Client) DeleteRepository(name *string) (_result *DeleteRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRepositoryResponse{}
	_body, _err := client.DeleteRepositoryWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除任务模板
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTaskTemplateResponse
func (client *Client) DeleteTaskTemplateWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTaskTemplateResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTaskTemplate"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/tasktemplates/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTaskTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除任务模板
//
// @return DeleteTaskTemplateResponse
func (client *Client) DeleteTaskTemplate(name *string) (_result *DeleteTaskTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTaskTemplateResponse{}
	_body, _err := client.DeleteTaskTemplateWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询身份绑定中的凭证信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchConnectionCredentialResponse
func (client *Client) FetchConnectionCredentialWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FetchConnectionCredentialResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("FetchConnectionCredential"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/connections/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/fetchCredential"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FetchConnectionCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询身份绑定中的凭证信息
//
// @return FetchConnectionCredentialResponse
func (client *Client) FetchConnectionCredential(name *string) (_result *FetchConnectionCredentialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FetchConnectionCredentialResponse{}
	_body, _err := client.FetchConnectionCredentialWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询仓库代码拉取所需信息
//
// @param request - FetchRepositoryCheckoutRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchRepositoryCheckoutResponse
func (client *Client) FetchRepositoryCheckoutWithOptions(name *string, request *FetchRepositoryCheckoutRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FetchRepositoryCheckoutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Branch)) {
		query["branch"] = request.Branch
	}

	if !tea.BoolValue(util.IsUnset(request.Commit)) {
		query["commit"] = request.Commit
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FetchRepositoryCheckout"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/repositories/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/fetchCheckout"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FetchRepositoryCheckoutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询仓库代码拉取所需信息
//
// @param request - FetchRepositoryCheckoutRequest
//
// @return FetchRepositoryCheckoutResponse
func (client *Client) FetchRepositoryCheckout(name *string, request *FetchRepositoryCheckoutRequest) (_result *FetchRepositoryCheckoutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FetchRepositoryCheckoutResponse{}
	_body, _err := client.FetchRepositoryCheckoutWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询身份绑定
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConnectionResponse
func (client *Client) GetConnectionWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetConnectionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetConnection"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/connections/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询身份绑定
//
// @return GetConnectionResponse
func (client *Client) GetConnection(name *string) (_result *GetConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetConnectionResponse{}
	_body, _err := client.GetConnectionWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取环境信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEnvironmentResponse
func (client *Client) GetEnvironmentWithOptions(project *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEnvironmentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetEnvironment"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/projects/" + tea.StringValue(openapiutil.GetEncodeParam(project)) + "/environments/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取环境信息
//
// @return GetEnvironmentResponse
func (client *Client) GetEnvironment(project *string, name *string) (_result *GetEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEnvironmentResponse{}
	_body, _err := client.GetEnvironmentWithOptions(project, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询流水线
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineResponse
func (client *Client) GetPipelineWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPipelineResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipeline"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询流水线
//
// @return GetPipelineResponse
func (client *Client) GetPipeline(name *string) (_result *GetPipelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPipelineResponse{}
	_body, _err := client.GetPipelineWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询流水线模板
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineTemplateResponse
func (client *Client) GetPipelineTemplateWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPipelineTemplateResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipelineTemplate"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelinetemplates/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipelineTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询流水线模板
//
// @return GetPipelineTemplateResponse
func (client *Client) GetPipelineTemplate(name *string) (_result *GetPipelineTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPipelineTemplateResponse{}
	_body, _err := client.GetPipelineTemplateWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询流水线触发器
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineTriggerResponse
func (client *Client) GetPipelineTriggerWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPipelineTriggerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipelineTrigger"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelinetriggers/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipelineTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询流水线触发器
//
// @return GetPipelineTriggerResponse
func (client *Client) GetPipelineTrigger(name *string) (_result *GetPipelineTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPipelineTriggerResponse{}
	_body, _err := client.GetPipelineTriggerWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询流水线触发历史
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineTriggerEventResponse
func (client *Client) GetPipelineTriggerEventWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPipelineTriggerEventResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipelineTriggerEvent"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelinetriggerevents/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipelineTriggerEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询流水线触发历史
//
// @return GetPipelineTriggerEventResponse
func (client *Client) GetPipelineTriggerEvent(name *string) (_result *GetPipelineTriggerEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPipelineTriggerEventResponse{}
	_body, _err := client.GetPipelineTriggerEventWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询项目
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetProject"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/projects/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询项目
//
// @return GetProjectResponse
func (client *Client) GetProject(name *string) (_result *GetProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询仓库绑定
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepositoryResponse
func (client *Client) GetRepositoryWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRepositoryResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetRepository"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/repositories/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询仓库绑定
//
// @return GetRepositoryResponse
func (client *Client) GetRepository(name *string) (_result *GetRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRepositoryResponse{}
	_body, _err := client.GetRepositoryWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResponse
func (client *Client) GetTaskWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTask"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务
//
// @return GetTaskResponse
func (client *Client) GetTask(name *string) (_result *GetTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskResponse{}
	_body, _err := client.GetTaskWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询任务模板
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskTemplateResponse
func (client *Client) GetTaskTemplateWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskTemplateResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskTemplate"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/tasktemplates/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务模板
//
// @return GetTaskTemplateResponse
func (client *Client) GetTaskTemplate(name *string) (_result *GetTaskTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskTemplateResponse{}
	_body, _err := client.GetTaskTemplateWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询身份绑定
//
// @param tmpReq - ListConnectionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConnectionsResponse
func (client *Client) ListConnectionsWithOptions(tmpReq *ListConnectionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListConnectionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListConnectionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LabelSelector)) {
		request.LabelSelectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSelector, tea.String("labelSelector"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.LabelSelectorShrink)) {
		query["labelSelector"] = request.LabelSelectorShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConnections"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/connections"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConnectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询身份绑定
//
// @param request - ListConnectionsRequest
//
// @return ListConnectionsResponse
func (client *Client) ListConnections(request *ListConnectionsRequest) (_result *ListConnectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConnectionsResponse{}
	_body, _err := client.ListConnectionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询环境列表
//
// @param tmpReq - ListEnvironmentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnvironmentsResponse
func (client *Client) ListEnvironmentsWithOptions(project *string, tmpReq *ListEnvironmentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEnvironmentsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListEnvironmentsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LabelSelector)) {
		request.LabelSelectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSelector, tea.String("labelSelector"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.LabelSelectorShrink)) {
		query["labelSelector"] = request.LabelSelectorShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEnvironments"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/projects/" + tea.StringValue(openapiutil.GetEncodeParam(project)) + "/environments/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEnvironmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询环境列表
//
// @param request - ListEnvironmentsRequest
//
// @return ListEnvironmentsResponse
func (client *Client) ListEnvironments(project *string, request *ListEnvironmentsRequest) (_result *ListEnvironmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEnvironmentsResponse{}
	_body, _err := client.ListEnvironmentsWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询流水线模板
//
// @param tmpReq - ListPipelineTemplatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelineTemplatesResponse
func (client *Client) ListPipelineTemplatesWithOptions(tmpReq *ListPipelineTemplatesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPipelineTemplatesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListPipelineTemplatesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LabelSelector)) {
		request.LabelSelectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSelector, tea.String("labelSelector"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelSelectorShrink)) {
		query["labelSelector"] = request.LabelSelectorShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPipelineTemplates"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelinetemplates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &ListPipelineTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询流水线模板
//
// @param request - ListPipelineTemplatesRequest
//
// @return ListPipelineTemplatesResponse
func (client *Client) ListPipelineTemplates(request *ListPipelineTemplatesRequest) (_result *ListPipelineTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelineTemplatesResponse{}
	_body, _err := client.ListPipelineTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询流水线触发历史
//
// @param tmpReq - ListPipelineTriggerEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelineTriggerEventsResponse
func (client *Client) ListPipelineTriggerEventsWithOptions(tmpReq *ListPipelineTriggerEventsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPipelineTriggerEventsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListPipelineTriggerEventsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LabelSelector)) {
		request.LabelSelectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSelector, tea.String("labelSelector"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.LabelSelectorShrink)) {
		query["labelSelector"] = request.LabelSelectorShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPipelineTriggerEvents"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelinetriggerevents"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPipelineTriggerEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询流水线触发历史
//
// @param request - ListPipelineTriggerEventsRequest
//
// @return ListPipelineTriggerEventsResponse
func (client *Client) ListPipelineTriggerEvents(request *ListPipelineTriggerEventsRequest) (_result *ListPipelineTriggerEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelineTriggerEventsResponse{}
	_body, _err := client.ListPipelineTriggerEventsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询流水线触发器
//
// @param tmpReq - ListPipelineTriggersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelineTriggersResponse
func (client *Client) ListPipelineTriggersWithOptions(tmpReq *ListPipelineTriggersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPipelineTriggersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListPipelineTriggersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LabelSelector)) {
		request.LabelSelectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSelector, tea.String("labelSelector"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.LabelSelectorShrink)) {
		query["labelSelector"] = request.LabelSelectorShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPipelineTriggers"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelinetriggers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPipelineTriggersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询流水线触发器
//
// @param request - ListPipelineTriggersRequest
//
// @return ListPipelineTriggersResponse
func (client *Client) ListPipelineTriggers(request *ListPipelineTriggersRequest) (_result *ListPipelineTriggersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelineTriggersResponse{}
	_body, _err := client.ListPipelineTriggersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询流水线
//
// @param tmpReq - ListPipelinesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelinesResponse
func (client *Client) ListPipelinesWithOptions(tmpReq *ListPipelinesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPipelinesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListPipelinesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LabelSelector)) {
		request.LabelSelectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSelector, tea.String("labelSelector"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelSelectorShrink)) {
		query["labelSelector"] = request.LabelSelectorShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPipelines"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelines"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &ListPipelinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询流水线
//
// @param request - ListPipelinesRequest
//
// @return ListPipelinesResponse
func (client *Client) ListPipelines(request *ListPipelinesRequest) (_result *ListPipelinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelinesResponse{}
	_body, _err := client.ListPipelinesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询项目
//
// @param tmpReq - ListProjectsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithOptions(tmpReq *ListProjectsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListProjectsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LabelSelector)) {
		request.LabelSelectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSelector, tea.String("labelSelector"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.LabelSelectorShrink)) {
		query["labelSelector"] = request.LabelSelectorShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjects"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/projects"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询项目
//
// @param request - ListProjectsRequest
//
// @return ListProjectsResponse
func (client *Client) ListProjects(request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询仓库绑定
//
// @param tmpReq - ListRepositoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepositoriesResponse
func (client *Client) ListRepositoriesWithOptions(tmpReq *ListRepositoriesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRepositoriesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListRepositoriesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LabelSelector)) {
		request.LabelSelectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSelector, tea.String("labelSelector"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.LabelSelectorShrink)) {
		query["labelSelector"] = request.LabelSelectorShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepositories"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/repositories"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepositoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询仓库绑定
//
// @param request - ListRepositoriesRequest
//
// @return ListRepositoriesResponse
func (client *Client) ListRepositories(request *ListRepositoriesRequest) (_result *ListRepositoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRepositoriesResponse{}
	_body, _err := client.ListRepositoriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询任务模板
//
// @param tmpReq - ListTaskTemplatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskTemplatesResponse
func (client *Client) ListTaskTemplatesWithOptions(tmpReq *ListTaskTemplatesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTaskTemplatesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTaskTemplatesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LabelSelector)) {
		request.LabelSelectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSelector, tea.String("labelSelector"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelSelectorShrink)) {
		query["labelSelector"] = request.LabelSelectorShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTaskTemplates"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/tasktemplates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &ListTaskTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询任务模板
//
// @param request - ListTaskTemplatesRequest
//
// @return ListTaskTemplatesResponse
func (client *Client) ListTaskTemplates(request *ListTaskTemplatesRequest) (_result *ListTaskTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTaskTemplatesResponse{}
	_body, _err := client.ListTaskTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询任务
//
// @param tmpReq - ListTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTasksResponse
func (client *Client) ListTasksWithOptions(tmpReq *ListTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LabelSelector)) {
		request.LabelSelectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSelector, tea.String("labelSelector"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelSelectorShrink)) {
		query["labelSelector"] = request.LabelSelectorShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTasks"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询任务
//
// @param request - ListTasksRequest
//
// @return ListTasksResponse
func (client *Client) ListTasks(request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新环境（全局更新）
//
// @param request - PutEnvironmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutEnvironmentResponse
func (client *Client) PutEnvironmentWithOptions(project *string, name *string, request *PutEnvironmentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutEnvironmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["force"] = request.Force
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutEnvironment"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/projects/" + tea.StringValue(openapiutil.GetEncodeParam(project)) + "/environments/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新环境（全局更新）
//
// @param request - PutEnvironmentRequest
//
// @return PutEnvironmentResponse
func (client *Client) PutEnvironment(project *string, name *string, request *PutEnvironmentRequest) (_result *PutEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutEnvironmentResponse{}
	_body, _err := client.PutEnvironmentWithOptions(project, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新流水线状态
//
// @param request - PutPipelineStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutPipelineStatusResponse
func (client *Client) PutPipelineStatusWithOptions(name *string, request *PutPipelineStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutPipelineStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["force"] = request.Force
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutPipelineStatus"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/status"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutPipelineStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新流水线状态
//
// @param request - PutPipelineStatusRequest
//
// @return PutPipelineStatusResponse
func (client *Client) PutPipelineStatus(name *string, request *PutPipelineStatusRequest) (_result *PutPipelineStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutPipelineStatusResponse{}
	_body, _err := client.PutPipelineStatusWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新替换流水线模板
//
// @param request - PutPipelineTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutPipelineTemplateResponse
func (client *Client) PutPipelineTemplateWithOptions(name *string, request *PutPipelineTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutPipelineTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["force"] = request.Force
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutPipelineTemplate"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelinetemplates/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutPipelineTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新替换流水线模板
//
// @param request - PutPipelineTemplateRequest
//
// @return PutPipelineTemplateResponse
func (client *Client) PutPipelineTemplate(name *string, request *PutPipelineTemplateRequest) (_result *PutPipelineTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutPipelineTemplateResponse{}
	_body, _err := client.PutPipelineTemplateWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新替换流水线触发器
//
// @param request - PutPipelineTriggerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutPipelineTriggerResponse
func (client *Client) PutPipelineTriggerWithOptions(name *string, request *PutPipelineTriggerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutPipelineTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["force"] = request.Force
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutPipelineTrigger"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelinetriggers/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutPipelineTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新替换流水线触发器
//
// @param request - PutPipelineTriggerRequest
//
// @return PutPipelineTriggerResponse
func (client *Client) PutPipelineTrigger(name *string, request *PutPipelineTriggerRequest) (_result *PutPipelineTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutPipelineTriggerResponse{}
	_body, _err := client.PutPipelineTriggerWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新替换应用
//
// @param request - PutProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutProjectResponse
func (client *Client) PutProjectWithOptions(name *string, request *PutProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["force"] = request.Force
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutProject"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/projects/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新替换应用
//
// @param request - PutProjectRequest
//
// @return PutProjectResponse
func (client *Client) PutProject(name *string, request *PutProjectRequest) (_result *PutProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutProjectResponse{}
	_body, _err := client.PutProjectWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新替换任务状态
//
// @param request - PutTaskStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutTaskStatusResponse
func (client *Client) PutTaskStatusWithOptions(name *string, request *PutTaskStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["force"] = request.Force
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutTaskStatus"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/status"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新替换任务状态
//
// @param request - PutTaskStatusRequest
//
// @return PutTaskStatusResponse
func (client *Client) PutTaskStatus(name *string, request *PutTaskStatusRequest) (_result *PutTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutTaskStatusResponse{}
	_body, _err := client.PutTaskStatusWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新替换任务模板
//
// @param request - PutTaskTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutTaskTemplateResponse
func (client *Client) PutTaskTemplateWithOptions(name *string, request *PutTaskTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutTaskTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["force"] = request.Force
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutTaskTemplate"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/tasktemplates/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutTaskTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新替换任务模板
//
// @param request - PutTaskTemplateRequest
//
// @return PutTaskTemplateResponse
func (client *Client) PutTaskTemplate(name *string, request *PutTaskTemplateRequest) (_result *PutTaskTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutTaskTemplateResponse{}
	_body, _err := client.PutTaskTemplateWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检查并刷新身份绑定中的凭证和账号信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshConnectionResponse
func (client *Client) RefreshConnectionWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RefreshConnectionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshConnection"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/connections/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/refresh"),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查并刷新身份绑定中的凭证和账号信息
//
// @return RefreshConnectionResponse
func (client *Client) RefreshConnection(name *string) (_result *RefreshConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RefreshConnectionResponse{}
	_body, _err := client.RefreshConnectionWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 确认并继续执行任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeTaskResponse
func (client *Client) ResumeTaskWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ResumeTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeTask"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/resume"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 确认并继续执行任务
//
// @return ResumeTaskResponse
func (client *Client) ResumeTask(name *string) (_result *ResumeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumeTaskResponse{}
	_body, _err := client.ResumeTaskWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重试执行任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryTaskResponse
func (client *Client) RetryTaskWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RetryTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RetryTask"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/retry"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RetryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重试执行任务
//
// @return RetryTaskResponse
func (client *Client) RetryTask(name *string) (_result *RetryTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RetryTaskResponse{}
	_body, _err := client.RetryTaskWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开始执行流水线
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartPipelineResponse
func (client *Client) StartPipelineWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartPipelineResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartPipeline"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/start"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartPipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开始执行流水线
//
// @return StartPipelineResponse
func (client *Client) StartPipeline(name *string) (_result *StartPipelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartPipelineResponse{}
	_body, _err := client.StartPipelineWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开始执行任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartTaskResponse
func (client *Client) StartTaskWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartTask"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/start"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开始执行任务
//
// @return StartTaskResponse
func (client *Client) StartTask(name *string) (_result *StartTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartTaskResponse{}
	_body, _err := client.StartTaskWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新环境（局部更新）
//
// @param request - UpdateEnvironmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEnvironmentResponse
func (client *Client) UpdateEnvironmentWithOptions(project *string, name *string, request *UpdateEnvironmentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateEnvironmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEnvironment"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/projects/" + tea.StringValue(openapiutil.GetEncodeParam(project)) + "/environments/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新环境（局部更新）
//
// @param request - UpdateEnvironmentRequest
//
// @return UpdateEnvironmentResponse
func (client *Client) UpdateEnvironment(project *string, name *string, request *UpdateEnvironmentRequest) (_result *UpdateEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEnvironmentResponse{}
	_body, _err := client.UpdateEnvironmentWithOptions(project, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新流水线触发器
//
// @param request - UpdatePipelineTriggerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePipelineTriggerResponse
func (client *Client) UpdatePipelineTriggerWithOptions(name *string, request *UpdatePipelineTriggerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdatePipelineTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePipelineTrigger"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/pipelinetriggers/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePipelineTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新流水线触发器
//
// @param request - UpdatePipelineTriggerRequest
//
// @return UpdatePipelineTriggerResponse
func (client *Client) UpdatePipelineTrigger(name *string, request *UpdatePipelineTriggerRequest) (_result *UpdatePipelineTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePipelineTriggerResponse{}
	_body, _err := client.UpdatePipelineTriggerWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新替换应用
//
// @param request - UpdateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectResponse
func (client *Client) UpdateProjectWithOptions(name *string, request *UpdateProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["force"] = request.Force
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProject"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/projects/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新替换应用
//
// @param request - UpdateProjectRequest
//
// @return UpdateProjectResponse
func (client *Client) UpdateProject(name *string, request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
