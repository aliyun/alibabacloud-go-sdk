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

type Application struct {
	AutoDeploy  *string                `json:"autoDeploy,omitempty" xml:"autoDeploy,omitempty"`
	CreatedTime *string                `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Description *string                `json:"description,omitempty" xml:"description,omitempty"`
	EnvVars     map[string]interface{} `json:"envVars,omitempty" xml:"envVars,omitempty"`
	LastRelease map[string]interface{} `json:"lastRelease,omitempty" xml:"lastRelease,omitempty"`
	Name        *string                `json:"name,omitempty" xml:"name,omitempty"`
	Output      map[string]interface{} `json:"output,omitempty" xml:"output,omitempty"`
	Parameters  map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	RepoSource  *ApplicationRepoSource `json:"repoSource,omitempty" xml:"repoSource,omitempty" type:"Struct"`
	RoleArn     *string                `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	Template    *string                `json:"template,omitempty" xml:"template,omitempty"`
	Trigger     *ApplicationTrigger    `json:"trigger,omitempty" xml:"trigger,omitempty" type:"Struct"`
	UpdatedTime *string                `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
	WorkDir     *string                `json:"workDir,omitempty" xml:"workDir,omitempty"`
}

func (s Application) String() string {
	return tea.Prettify(s)
}

func (s Application) GoString() string {
	return s.String()
}

func (s *Application) SetAutoDeploy(v string) *Application {
	s.AutoDeploy = &v
	return s
}

func (s *Application) SetCreatedTime(v string) *Application {
	s.CreatedTime = &v
	return s
}

func (s *Application) SetDescription(v string) *Application {
	s.Description = &v
	return s
}

func (s *Application) SetEnvVars(v map[string]interface{}) *Application {
	s.EnvVars = v
	return s
}

func (s *Application) SetLastRelease(v map[string]interface{}) *Application {
	s.LastRelease = v
	return s
}

func (s *Application) SetName(v string) *Application {
	s.Name = &v
	return s
}

func (s *Application) SetOutput(v map[string]interface{}) *Application {
	s.Output = v
	return s
}

func (s *Application) SetParameters(v map[string]interface{}) *Application {
	s.Parameters = v
	return s
}

func (s *Application) SetRepoSource(v *ApplicationRepoSource) *Application {
	s.RepoSource = v
	return s
}

func (s *Application) SetRoleArn(v string) *Application {
	s.RoleArn = &v
	return s
}

func (s *Application) SetTemplate(v string) *Application {
	s.Template = &v
	return s
}

func (s *Application) SetTrigger(v *ApplicationTrigger) *Application {
	s.Trigger = v
	return s
}

func (s *Application) SetUpdatedTime(v string) *Application {
	s.UpdatedTime = &v
	return s
}

func (s *Application) SetWorkDir(v string) *Application {
	s.WorkDir = &v
	return s
}

type ApplicationRepoSource struct {
	Owner    *string `json:"owner,omitempty" xml:"owner,omitempty"`
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	Repo     *string `json:"repo,omitempty" xml:"repo,omitempty"`
}

func (s ApplicationRepoSource) String() string {
	return tea.Prettify(s)
}

func (s ApplicationRepoSource) GoString() string {
	return s.String()
}

func (s *ApplicationRepoSource) SetOwner(v string) *ApplicationRepoSource {
	s.Owner = &v
	return s
}

func (s *ApplicationRepoSource) SetProvider(v string) *ApplicationRepoSource {
	s.Provider = &v
	return s
}

func (s *ApplicationRepoSource) SetRepo(v string) *ApplicationRepoSource {
	s.Repo = &v
	return s
}

type ApplicationTrigger struct {
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	Commit *string `json:"commit,omitempty" xml:"commit,omitempty"`
	On     *string `json:"on,omitempty" xml:"on,omitempty"`
}

func (s ApplicationTrigger) String() string {
	return tea.Prettify(s)
}

func (s ApplicationTrigger) GoString() string {
	return s.String()
}

func (s *ApplicationTrigger) SetBranch(v string) *ApplicationTrigger {
	s.Branch = &v
	return s
}

func (s *ApplicationTrigger) SetCommit(v string) *ApplicationTrigger {
	s.Commit = &v
	return s
}

func (s *ApplicationTrigger) SetOn(v string) *ApplicationTrigger {
	s.On = &v
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
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Hint        *string `json:"hint,omitempty" xml:"hint,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Required    *bool   `json:"required,omitempty" xml:"required,omitempty"`
	Type        *string `json:"type,omitempty" xml:"type,omitempty"`
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
	CreatedTime  *string            `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	DeletionTime *string            `json:"deletionTime,omitempty" xml:"deletionTime,omitempty"`
	Description  *string            `json:"description,omitempty" xml:"description,omitempty"`
	Generation   *int32             `json:"generation,omitempty" xml:"generation,omitempty"`
	Kind         *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Name         *string            `json:"name,omitempty" xml:"name,omitempty"`
	Spec         *EnvironmentSpec   `json:"spec,omitempty" xml:"spec,omitempty"`
	Status       *EnvironmentStatus `json:"status,omitempty" xml:"status,omitempty"`
	Uid          *string            `json:"uid,omitempty" xml:"uid,omitempty"`
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

func (s *Environment) SetDeletionTime(v string) *Environment {
	s.DeletionTime = &v
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

func (s *Environment) SetName(v string) *Environment {
	s.Name = &v
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

type EnvironmentRevision struct {
	CreatedTime           *string            `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	EnvironmentGeneration *int32             `json:"environmentGeneration,omitempty" xml:"environmentGeneration,omitempty"`
	EnvironmentName       *string            `json:"environmentName,omitempty" xml:"environmentName,omitempty"`
	Kind                  *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Spec                  *EnvironmentSpec   `json:"spec,omitempty" xml:"spec,omitempty"`
	Status                *EnvironmentStatus `json:"status,omitempty" xml:"status,omitempty"`
	Uid                   *string            `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s EnvironmentRevision) String() string {
	return tea.Prettify(s)
}

func (s EnvironmentRevision) GoString() string {
	return s.String()
}

func (s *EnvironmentRevision) SetCreatedTime(v string) *EnvironmentRevision {
	s.CreatedTime = &v
	return s
}

func (s *EnvironmentRevision) SetEnvironmentGeneration(v int32) *EnvironmentRevision {
	s.EnvironmentGeneration = &v
	return s
}

func (s *EnvironmentRevision) SetEnvironmentName(v string) *EnvironmentRevision {
	s.EnvironmentName = &v
	return s
}

func (s *EnvironmentRevision) SetKind(v string) *EnvironmentRevision {
	s.Kind = &v
	return s
}

func (s *EnvironmentRevision) SetSpec(v *EnvironmentSpec) *EnvironmentRevision {
	s.Spec = v
	return s
}

func (s *EnvironmentRevision) SetStatus(v *EnvironmentStatus) *EnvironmentRevision {
	s.Status = v
	return s
}

func (s *EnvironmentRevision) SetUid(v string) *EnvironmentRevision {
	s.Uid = &v
	return s
}

type EnvironmentSpec struct {
	Region            *string                `json:"region,omitempty" xml:"region,omitempty"`
	RoleArn           *string                `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	Template          *string                `json:"template,omitempty" xml:"template,omitempty"`
	TemplateVariables map[string]interface{} `json:"templateVariables,omitempty" xml:"templateVariables,omitempty"`
	TemplateVersion   *int32                 `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
}

func (s EnvironmentSpec) String() string {
	return tea.Prettify(s)
}

func (s EnvironmentSpec) GoString() string {
	return s.String()
}

func (s *EnvironmentSpec) SetRegion(v string) *EnvironmentSpec {
	s.Region = &v
	return s
}

func (s *EnvironmentSpec) SetRoleArn(v string) *EnvironmentSpec {
	s.RoleArn = &v
	return s
}

func (s *EnvironmentSpec) SetTemplate(v string) *EnvironmentSpec {
	s.Template = &v
	return s
}

func (s *EnvironmentSpec) SetTemplateVariables(v map[string]interface{}) *EnvironmentSpec {
	s.TemplateVariables = v
	return s
}

func (s *EnvironmentSpec) SetTemplateVersion(v int32) *EnvironmentSpec {
	s.TemplateVersion = &v
	return s
}

type EnvironmentStatus struct {
	Message            *string                `json:"message,omitempty" xml:"message,omitempty"`
	ObservedGeneration *int32                 `json:"observedGeneration,omitempty" xml:"observedGeneration,omitempty"`
	ObservedTime       *string                `json:"observedTime,omitempty" xml:"observedTime,omitempty"`
	Output             map[string]interface{} `json:"output,omitempty" xml:"output,omitempty"`
	Phase              *string                `json:"phase,omitempty" xml:"phase,omitempty"`
}

func (s EnvironmentStatus) String() string {
	return tea.Prettify(s)
}

func (s EnvironmentStatus) GoString() string {
	return s.String()
}

func (s *EnvironmentStatus) SetMessage(v string) *EnvironmentStatus {
	s.Message = &v
	return s
}

func (s *EnvironmentStatus) SetObservedGeneration(v int32) *EnvironmentStatus {
	s.ObservedGeneration = &v
	return s
}

func (s *EnvironmentStatus) SetObservedTime(v string) *EnvironmentStatus {
	s.ObservedTime = &v
	return s
}

func (s *EnvironmentStatus) SetOutput(v map[string]interface{}) *EnvironmentStatus {
	s.Output = v
	return s
}

func (s *EnvironmentStatus) SetPhase(v string) *EnvironmentStatus {
	s.Phase = &v
	return s
}

type InputVariable struct {
	DefaultJson *string `json:"defaultJson,omitempty" xml:"defaultJson,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Nullable    *bool   `json:"nullable,omitempty" xml:"nullable,omitempty"`
	Sensitive   *bool   `json:"sensitive,omitempty" xml:"sensitive,omitempty"`
	Type        *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s InputVariable) String() string {
	return tea.Prettify(s)
}

func (s InputVariable) GoString() string {
	return s.String()
}

func (s *InputVariable) SetDefaultJson(v string) *InputVariable {
	s.DefaultJson = &v
	return s
}

func (s *InputVariable) SetDescription(v string) *InputVariable {
	s.Description = &v
	return s
}

func (s *InputVariable) SetName(v string) *InputVariable {
	s.Name = &v
	return s
}

func (s *InputVariable) SetNullable(v bool) *InputVariable {
	s.Nullable = &v
	return s
}

func (s *InputVariable) SetSensitive(v bool) *InputVariable {
	s.Sensitive = &v
	return s
}

func (s *InputVariable) SetType(v string) *InputVariable {
	s.Type = &v
	return s
}

type OutputValue struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Sensitive   *bool   `json:"sensitive,omitempty" xml:"sensitive,omitempty"`
}

func (s OutputValue) String() string {
	return tea.Prettify(s)
}

func (s OutputValue) GoString() string {
	return s.String()
}

func (s *OutputValue) SetDescription(v string) *OutputValue {
	s.Description = &v
	return s
}

func (s *OutputValue) SetName(v string) *OutputValue {
	s.Name = &v
	return s
}

func (s *OutputValue) SetSensitive(v bool) *OutputValue {
	s.Sensitive = &v
	return s
}

type Pipeline struct {
	CreatedTime     *string            `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	DeletionTime    *string            `json:"deletionTime,omitempty" xml:"deletionTime,omitempty"`
	Description     *string            `json:"description,omitempty" xml:"description,omitempty"`
	Generation      *int32             `json:"generation,omitempty" xml:"generation,omitempty"`
	Kind            *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels          map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	Name            *string            `json:"name,omitempty" xml:"name,omitempty"`
	ResourceVersion *int32             `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
	Spec            *PipelineSpec      `json:"spec,omitempty" xml:"spec,omitempty"`
	Status          *PipelineStatus    `json:"status,omitempty" xml:"status,omitempty"`
	Uid             *string            `json:"uid,omitempty" xml:"uid,omitempty"`
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
	Context      *Context              `json:"context,omitempty" xml:"context,omitempty"`
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
	Phase           *string        `json:"phase,omitempty" xml:"phase,omitempty"`
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
	CreatedTime     *string               `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	DeletionTime    *string               `json:"deletionTime,omitempty" xml:"deletionTime,omitempty"`
	Description     *string               `json:"description,omitempty" xml:"description,omitempty"`
	Generation      *int32                `json:"generation,omitempty" xml:"generation,omitempty"`
	Kind            *string               `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels          map[string]*string    `json:"labels,omitempty" xml:"labels,omitempty"`
	Name            *string               `json:"name,omitempty" xml:"name,omitempty"`
	ResourceVersion *int32                `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
	Spec            *PipelineTemplateSpec `json:"spec,omitempty" xml:"spec,omitempty"`
	Uid             *string               `json:"uid,omitempty" xml:"uid,omitempty"`
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

type Release struct {
	AppConfig   map[string]interface{} `json:"appConfig,omitempty" xml:"appConfig,omitempty"`
	CodeVersion *ReleaseCodeVersion    `json:"codeVersion,omitempty" xml:"codeVersion,omitempty" type:"Struct"`
	CreatedTime *string                `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Description *string                `json:"description,omitempty" xml:"description,omitempty"`
	Output      map[string]interface{} `json:"output,omitempty" xml:"output,omitempty"`
	Status      *string                `json:"status,omitempty" xml:"status,omitempty"`
	VersionId   *int64                 `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s Release) String() string {
	return tea.Prettify(s)
}

func (s Release) GoString() string {
	return s.String()
}

func (s *Release) SetAppConfig(v map[string]interface{}) *Release {
	s.AppConfig = v
	return s
}

func (s *Release) SetCodeVersion(v *ReleaseCodeVersion) *Release {
	s.CodeVersion = v
	return s
}

func (s *Release) SetCreatedTime(v string) *Release {
	s.CreatedTime = &v
	return s
}

func (s *Release) SetDescription(v string) *Release {
	s.Description = &v
	return s
}

func (s *Release) SetOutput(v map[string]interface{}) *Release {
	s.Output = v
	return s
}

func (s *Release) SetStatus(v string) *Release {
	s.Status = &v
	return s
}

func (s *Release) SetVersionId(v int64) *Release {
	s.VersionId = &v
	return s
}

type ReleaseCodeVersion struct {
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	Commit *string `json:"commit,omitempty" xml:"commit,omitempty"`
}

func (s ReleaseCodeVersion) String() string {
	return tea.Prettify(s)
}

func (s ReleaseCodeVersion) GoString() string {
	return s.String()
}

func (s *ReleaseCodeVersion) SetBranch(v string) *ReleaseCodeVersion {
	s.Branch = &v
	return s
}

func (s *ReleaseCodeVersion) SetCommit(v string) *ReleaseCodeVersion {
	s.Commit = &v
	return s
}

type RepoSource struct {
	Owner    *string `json:"owner,omitempty" xml:"owner,omitempty"`
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	Repo     *string `json:"repo,omitempty" xml:"repo,omitempty"`
}

func (s RepoSource) String() string {
	return tea.Prettify(s)
}

func (s RepoSource) GoString() string {
	return s.String()
}

func (s *RepoSource) SetOwner(v string) *RepoSource {
	s.Owner = &v
	return s
}

func (s *RepoSource) SetProvider(v string) *RepoSource {
	s.Provider = &v
	return s
}

func (s *RepoSource) SetRepo(v string) *RepoSource {
	s.Repo = &v
	return s
}

type RunAfter struct {
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

type Service struct {
	CreatedTime  *string        `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	DeletionTime *string        `json:"deletionTime,omitempty" xml:"deletionTime,omitempty"`
	Description  *string        `json:"description,omitempty" xml:"description,omitempty"`
	Generation   *int32         `json:"generation,omitempty" xml:"generation,omitempty"`
	Kind         *string        `json:"kind,omitempty" xml:"kind,omitempty"`
	Name         *string        `json:"name,omitempty" xml:"name,omitempty"`
	Spec         *ServiceSpec   `json:"spec,omitempty" xml:"spec,omitempty"`
	Status       *ServiceStatus `json:"status,omitempty" xml:"status,omitempty"`
	Uid          *string        `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s Service) String() string {
	return tea.Prettify(s)
}

func (s Service) GoString() string {
	return s.String()
}

func (s *Service) SetCreatedTime(v string) *Service {
	s.CreatedTime = &v
	return s
}

func (s *Service) SetDeletionTime(v string) *Service {
	s.DeletionTime = &v
	return s
}

func (s *Service) SetDescription(v string) *Service {
	s.Description = &v
	return s
}

func (s *Service) SetGeneration(v int32) *Service {
	s.Generation = &v
	return s
}

func (s *Service) SetKind(v string) *Service {
	s.Kind = &v
	return s
}

func (s *Service) SetName(v string) *Service {
	s.Name = &v
	return s
}

func (s *Service) SetSpec(v *ServiceSpec) *Service {
	s.Spec = v
	return s
}

func (s *Service) SetStatus(v *ServiceStatus) *Service {
	s.Status = v
	return s
}

func (s *Service) SetUid(v string) *Service {
	s.Uid = &v
	return s
}

type ServiceRevision struct {
	CreatedTime       *string            `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Kind              *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	ServiceGeneration *int32             `json:"serviceGeneration,omitempty" xml:"serviceGeneration,omitempty"`
	ServiceName       *string            `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	Spec              *ServiceSpec       `json:"spec,omitempty" xml:"spec,omitempty"`
	Status            *EnvironmentStatus `json:"status,omitempty" xml:"status,omitempty"`
	Uid               *string            `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ServiceRevision) String() string {
	return tea.Prettify(s)
}

func (s ServiceRevision) GoString() string {
	return s.String()
}

func (s *ServiceRevision) SetCreatedTime(v string) *ServiceRevision {
	s.CreatedTime = &v
	return s
}

func (s *ServiceRevision) SetKind(v string) *ServiceRevision {
	s.Kind = &v
	return s
}

func (s *ServiceRevision) SetServiceGeneration(v int32) *ServiceRevision {
	s.ServiceGeneration = &v
	return s
}

func (s *ServiceRevision) SetServiceName(v string) *ServiceRevision {
	s.ServiceName = &v
	return s
}

func (s *ServiceRevision) SetSpec(v *ServiceSpec) *ServiceRevision {
	s.Spec = v
	return s
}

func (s *ServiceRevision) SetStatus(v *EnvironmentStatus) *ServiceRevision {
	s.Status = v
	return s
}

func (s *ServiceRevision) SetUid(v string) *ServiceRevision {
	s.Uid = &v
	return s
}

type ServiceSpec struct {
	Environment       *string                `json:"environment,omitempty" xml:"environment,omitempty"`
	RoleArn           *string                `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	Template          *string                `json:"template,omitempty" xml:"template,omitempty"`
	TemplateVariables map[string]interface{} `json:"templateVariables,omitempty" xml:"templateVariables,omitempty"`
	TemplateVersion   *int32                 `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
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

type ServiceStatus struct {
	Message            *string                `json:"message,omitempty" xml:"message,omitempty"`
	ObservedGeneration *int32                 `json:"observedGeneration,omitempty" xml:"observedGeneration,omitempty"`
	ObservedTime       *string                `json:"observedTime,omitempty" xml:"observedTime,omitempty"`
	Output             map[string]interface{} `json:"output,omitempty" xml:"output,omitempty"`
	Phase              *string                `json:"phase,omitempty" xml:"phase,omitempty"`
}

func (s ServiceStatus) String() string {
	return tea.Prettify(s)
}

func (s ServiceStatus) GoString() string {
	return s.String()
}

func (s *ServiceStatus) SetMessage(v string) *ServiceStatus {
	s.Message = &v
	return s
}

func (s *ServiceStatus) SetObservedGeneration(v int32) *ServiceStatus {
	s.ObservedGeneration = &v
	return s
}

func (s *ServiceStatus) SetObservedTime(v string) *ServiceStatus {
	s.ObservedTime = &v
	return s
}

func (s *ServiceStatus) SetOutput(v map[string]interface{}) *ServiceStatus {
	s.Output = v
	return s
}

func (s *ServiceStatus) SetPhase(v string) *ServiceStatus {
	s.Phase = &v
	return s
}

type Status struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s Status) String() string {
	return tea.Prettify(s)
}

func (s Status) GoString() string {
	return s.String()
}

func (s *Status) SetCode(v string) *Status {
	s.Code = &v
	return s
}

func (s *Status) SetMessage(v string) *Status {
	s.Message = &v
	return s
}

func (s *Status) SetRequestId(v string) *Status {
	s.RequestId = &v
	return s
}

func (s *Status) SetSuccess(v bool) *Status {
	s.Success = &v
	return s
}

type StsCredentials struct {
	AccessKeyId     *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	ExpirationTime  *string `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	Kind            *string `json:"kind,omitempty" xml:"kind,omitempty"`
	SecretAccessKey *string `json:"secretAccessKey,omitempty" xml:"secretAccessKey,omitempty"`
	Token           *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s StsCredentials) String() string {
	return tea.Prettify(s)
}

func (s StsCredentials) GoString() string {
	return s.String()
}

func (s *StsCredentials) SetAccessKeyId(v string) *StsCredentials {
	s.AccessKeyId = &v
	return s
}

func (s *StsCredentials) SetExpirationTime(v string) *StsCredentials {
	s.ExpirationTime = &v
	return s
}

func (s *StsCredentials) SetKind(v string) *StsCredentials {
	s.Kind = &v
	return s
}

func (s *StsCredentials) SetSecretAccessKey(v string) *StsCredentials {
	s.SecretAccessKey = &v
	return s
}

func (s *StsCredentials) SetToken(v string) *StsCredentials {
	s.Token = &v
	return s
}

type Task struct {
	CreatedTime     *string            `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	DeletionTime    *string            `json:"deletionTime,omitempty" xml:"deletionTime,omitempty"`
	Description     *string            `json:"description,omitempty" xml:"description,omitempty"`
	Generation      *int32             `json:"generation,omitempty" xml:"generation,omitempty"`
	Kind            *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels          map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	Name            *string            `json:"name,omitempty" xml:"name,omitempty"`
	ResourceVersion *int32             `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
	Spec            *TaskSpec          `json:"spec,omitempty" xml:"spec,omitempty"`
	Status          *TaskStatus        `json:"status,omitempty" xml:"status,omitempty"`
	Uid             *string            `json:"uid,omitempty" xml:"uid,omitempty"`
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
	Context      *Context    `json:"context,omitempty" xml:"context,omitempty"`
	Name         *string     `json:"name,omitempty" xml:"name,omitempty"`
	RunAfters    []*RunAfter `json:"runAfters,omitempty" xml:"runAfters,omitempty" type:"Repeated"`
	TaskTemplate *string     `json:"taskTemplate,omitempty" xml:"taskTemplate,omitempty"`
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
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
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
	InstanceID       *string `json:"instanceID,omitempty" xml:"instanceID,omitempty"`
	InvocationID     *string `json:"invocationID,omitempty" xml:"invocationID,omitempty"`
	InvocationTarget *string `json:"invocationTarget,omitempty" xml:"invocationTarget,omitempty"`
	Output           *string `json:"output,omitempty" xml:"output,omitempty"`
	RequestID        *string `json:"requestID,omitempty" xml:"requestID,omitempty"`
	SlsLogStore      *string `json:"slsLogStore,omitempty" xml:"slsLogStore,omitempty"`
	SlsProject       *string `json:"slsProject,omitempty" xml:"slsProject,omitempty"`
	Status           *string `json:"status,omitempty" xml:"status,omitempty"`
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
	Context      *Context `json:"context,omitempty" xml:"context,omitempty"`
	TemplateName *string  `json:"templateName,omitempty" xml:"templateName,omitempty"`
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
	Phase            *string           `json:"phase,omitempty" xml:"phase,omitempty"`
	StatusGeneration *int64            `json:"statusGeneration,omitempty" xml:"statusGeneration,omitempty"`
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
	CreatedTime     *string            `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	DeletionTime    *string            `json:"deletionTime,omitempty" xml:"deletionTime,omitempty"`
	Description     *string            `json:"description,omitempty" xml:"description,omitempty"`
	Generation      *int32             `json:"generation,omitempty" xml:"generation,omitempty"`
	Kind            *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels          map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	Name            *string            `json:"name,omitempty" xml:"name,omitempty"`
	ResourceVersion *int32             `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
	Spec            *TaskTemplateSpec  `json:"spec,omitempty" xml:"spec,omitempty"`
	Uid             *string            `json:"uid,omitempty" xml:"uid,omitempty"`
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
	Context          *Context    `json:"context,omitempty" xml:"context,omitempty"`
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

type Template struct {
	CreatedTime  *string         `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	DeletionTime *string         `json:"deletionTime,omitempty" xml:"deletionTime,omitempty"`
	Description  *string         `json:"description,omitempty" xml:"description,omitempty"`
	Generation   *int32          `json:"generation,omitempty" xml:"generation,omitempty"`
	Kind         *string         `json:"kind,omitempty" xml:"kind,omitempty"`
	Name         *string         `json:"name,omitempty" xml:"name,omitempty"`
	Spec         *TemplateSpec   `json:"spec,omitempty" xml:"spec,omitempty"`
	Status       *TemplateStatus `json:"status,omitempty" xml:"status,omitempty"`
	Uid          *string         `json:"uid,omitempty" xml:"uid,omitempty"`
	Version      *int32          `json:"version,omitempty" xml:"version,omitempty"`
}

func (s Template) String() string {
	return tea.Prettify(s)
}

func (s Template) GoString() string {
	return s.String()
}

func (s *Template) SetCreatedTime(v string) *Template {
	s.CreatedTime = &v
	return s
}

func (s *Template) SetDeletionTime(v string) *Template {
	s.DeletionTime = &v
	return s
}

func (s *Template) SetDescription(v string) *Template {
	s.Description = &v
	return s
}

func (s *Template) SetGeneration(v int32) *Template {
	s.Generation = &v
	return s
}

func (s *Template) SetKind(v string) *Template {
	s.Kind = &v
	return s
}

func (s *Template) SetName(v string) *Template {
	s.Name = &v
	return s
}

func (s *Template) SetSpec(v *TemplateSpec) *Template {
	s.Spec = v
	return s
}

func (s *Template) SetStatus(v *TemplateStatus) *Template {
	s.Status = v
	return s
}

func (s *Template) SetUid(v string) *Template {
	s.Uid = &v
	return s
}

func (s *Template) SetVersion(v int32) *Template {
	s.Version = &v
	return s
}

type TemplateRevision struct {
	CreatedTime        *string         `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Kind               *string         `json:"kind,omitempty" xml:"kind,omitempty"`
	Spec               *TemplateSpec   `json:"spec,omitempty" xml:"spec,omitempty"`
	Status             *TemplateStatus `json:"status,omitempty" xml:"status,omitempty"`
	TemplateGeneration *int32          `json:"templateGeneration,omitempty" xml:"templateGeneration,omitempty"`
	TemplateName       *string         `json:"templateName,omitempty" xml:"templateName,omitempty"`
	TemplateVersion    *int32          `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
	Uid                *string         `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s TemplateRevision) String() string {
	return tea.Prettify(s)
}

func (s TemplateRevision) GoString() string {
	return s.String()
}

func (s *TemplateRevision) SetCreatedTime(v string) *TemplateRevision {
	s.CreatedTime = &v
	return s
}

func (s *TemplateRevision) SetKind(v string) *TemplateRevision {
	s.Kind = &v
	return s
}

func (s *TemplateRevision) SetSpec(v *TemplateSpec) *TemplateRevision {
	s.Spec = v
	return s
}

func (s *TemplateRevision) SetStatus(v *TemplateStatus) *TemplateRevision {
	s.Status = v
	return s
}

func (s *TemplateRevision) SetTemplateGeneration(v int32) *TemplateRevision {
	s.TemplateGeneration = &v
	return s
}

func (s *TemplateRevision) SetTemplateName(v string) *TemplateRevision {
	s.TemplateName = &v
	return s
}

func (s *TemplateRevision) SetTemplateVersion(v int32) *TemplateRevision {
	s.TemplateVersion = &v
	return s
}

func (s *TemplateRevision) SetUid(v string) *TemplateRevision {
	s.Uid = &v
	return s
}

type TemplateSpec struct {
	Content     *string `json:"content,omitempty" xml:"content,omitempty"`
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	RamPolicy   *string `json:"ramPolicy,omitempty" xml:"ramPolicy,omitempty"`
	Type        *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TemplateSpec) String() string {
	return tea.Prettify(s)
}

func (s TemplateSpec) GoString() string {
	return s.String()
}

func (s *TemplateSpec) SetContent(v string) *TemplateSpec {
	s.Content = &v
	return s
}

func (s *TemplateSpec) SetContentType(v string) *TemplateSpec {
	s.ContentType = &v
	return s
}

func (s *TemplateSpec) SetRamPolicy(v string) *TemplateSpec {
	s.RamPolicy = &v
	return s
}

func (s *TemplateSpec) SetType(v string) *TemplateSpec {
	s.Type = &v
	return s
}

type TemplateStatus struct {
	Message            *string          `json:"message,omitempty" xml:"message,omitempty"`
	ObservedGeneration *int32           `json:"observedGeneration,omitempty" xml:"observedGeneration,omitempty"`
	ObservedTime       *string          `json:"observedTime,omitempty" xml:"observedTime,omitempty"`
	Outputs            []*OutputValue   `json:"outputs,omitempty" xml:"outputs,omitempty" type:"Repeated"`
	Phase              *string          `json:"phase,omitempty" xml:"phase,omitempty"`
	Variables          []*InputVariable `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s TemplateStatus) String() string {
	return tea.Prettify(s)
}

func (s TemplateStatus) GoString() string {
	return s.String()
}

func (s *TemplateStatus) SetMessage(v string) *TemplateStatus {
	s.Message = &v
	return s
}

func (s *TemplateStatus) SetObservedGeneration(v int32) *TemplateStatus {
	s.ObservedGeneration = &v
	return s
}

func (s *TemplateStatus) SetObservedTime(v string) *TemplateStatus {
	s.ObservedTime = &v
	return s
}

func (s *TemplateStatus) SetOutputs(v []*OutputValue) *TemplateStatus {
	s.Outputs = v
	return s
}

func (s *TemplateStatus) SetPhase(v string) *TemplateStatus {
	s.Phase = &v
	return s
}

func (s *TemplateStatus) SetVariables(v []*InputVariable) *TemplateStatus {
	s.Variables = v
	return s
}

type TriggerConfig struct {
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	Commit *string `json:"commit,omitempty" xml:"commit,omitempty"`
	On     *string `json:"on,omitempty" xml:"on,omitempty"`
}

func (s TriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s TriggerConfig) GoString() string {
	return s.String()
}

func (s *TriggerConfig) SetBranch(v string) *TriggerConfig {
	s.Branch = &v
	return s
}

func (s *TriggerConfig) SetCommit(v string) *TriggerConfig {
	s.Commit = &v
	return s
}

func (s *TriggerConfig) SetOn(v string) *TriggerConfig {
	s.On = &v
	return s
}

type CancelTaskResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Task              `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateApplicationRequest struct {
	AutoDeploy  *bool              `json:"autoDeploy,omitempty" xml:"autoDeploy,omitempty"`
	Description *string            `json:"description,omitempty" xml:"description,omitempty"`
	EnvVars     map[string]*string `json:"envVars,omitempty" xml:"envVars,omitempty"`
	Name        *string            `json:"name,omitempty" xml:"name,omitempty"`
	Parameters  map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	RepoSource  *RepoSource        `json:"repoSource,omitempty" xml:"repoSource,omitempty"`
	RoleArn     *string            `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	Template    *string            `json:"template,omitempty" xml:"template,omitempty"`
	Trigger     *TriggerConfig     `json:"trigger,omitempty" xml:"trigger,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) SetAutoDeploy(v bool) *CreateApplicationRequest {
	s.AutoDeploy = &v
	return s
}

func (s *CreateApplicationRequest) SetDescription(v string) *CreateApplicationRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationRequest) SetEnvVars(v map[string]*string) *CreateApplicationRequest {
	s.EnvVars = v
	return s
}

func (s *CreateApplicationRequest) SetName(v string) *CreateApplicationRequest {
	s.Name = &v
	return s
}

func (s *CreateApplicationRequest) SetParameters(v map[string]*string) *CreateApplicationRequest {
	s.Parameters = v
	return s
}

func (s *CreateApplicationRequest) SetRepoSource(v *RepoSource) *CreateApplicationRequest {
	s.RepoSource = v
	return s
}

func (s *CreateApplicationRequest) SetRoleArn(v string) *CreateApplicationRequest {
	s.RoleArn = &v
	return s
}

func (s *CreateApplicationRequest) SetTemplate(v string) *CreateApplicationRequest {
	s.Template = &v
	return s
}

func (s *CreateApplicationRequest) SetTrigger(v *TriggerConfig) *CreateApplicationRequest {
	s.Trigger = v
	return s
}

type CreateApplicationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Application       `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponse) SetHeaders(v map[string]*string) *CreateApplicationResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationResponse) SetStatusCode(v int32) *CreateApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationResponse) SetBody(v *Application) *CreateApplicationResponse {
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Pipeline          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PipelineTemplate  `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateReleaseRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s CreateReleaseRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateReleaseRequest) GoString() string {
	return s.String()
}

func (s *CreateReleaseRequest) SetDescription(v string) *CreateReleaseRequest {
	s.Description = &v
	return s
}

type CreateReleaseResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Release           `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateReleaseResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateReleaseResponse) GoString() string {
	return s.String()
}

func (s *CreateReleaseResponse) SetHeaders(v map[string]*string) *CreateReleaseResponse {
	s.Headers = v
	return s
}

func (s *CreateReleaseResponse) SetStatusCode(v int32) *CreateReleaseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateReleaseResponse) SetBody(v *Release) *CreateReleaseResponse {
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Task              `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TaskTemplate      `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteApplicationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponse) SetHeaders(v map[string]*string) *DeleteApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationResponse) SetStatusCode(v int32) *DeleteApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationResponse) SetBody(v string) *DeleteApplicationResponse {
	s.Body = &v
	return s
}

type DeleteEnvironmentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePipelineTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteTaskTemplateResponseBody struct {
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTaskTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteTemplateRequest struct {
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DeleteTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateRequest) SetVersion(v int32) *DeleteTemplateRequest {
	s.Version = &v
	return s
}

type DeleteTemplateResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Status            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponse) SetHeaders(v map[string]*string) *DeleteTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplateResponse) SetStatusCode(v int32) *DeleteTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTemplateResponse) SetBody(v *Status) *DeleteTemplateResponse {
	s.Body = v
	return s
}

type GetApplicationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Application       `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationResponse) SetHeaders(v map[string]*string) *GetApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationResponse) SetStatusCode(v int32) *GetApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationResponse) SetBody(v *Application) *GetApplicationResponse {
	s.Body = v
	return s
}

type GetEnvironmentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Environment       `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Pipeline          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PipelineTemplate  `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetReleaseResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Release           `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetReleaseResponse) String() string {
	return tea.Prettify(s)
}

func (s GetReleaseResponse) GoString() string {
	return s.String()
}

func (s *GetReleaseResponse) SetHeaders(v map[string]*string) *GetReleaseResponse {
	s.Headers = v
	return s
}

func (s *GetReleaseResponse) SetStatusCode(v int32) *GetReleaseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReleaseResponse) SetBody(v *Release) *GetReleaseResponse {
	s.Body = v
	return s
}

type GetServiceResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Service           `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponse) GoString() string {
	return s.String()
}

func (s *GetServiceResponse) SetHeaders(v map[string]*string) *GetServiceResponse {
	s.Headers = v
	return s
}

func (s *GetServiceResponse) SetStatusCode(v int32) *GetServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceResponse) SetBody(v *Service) *GetServiceResponse {
	s.Body = v
	return s
}

type GetTaskResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Task              `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TaskTemplate      `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetTemplateRequest struct {
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) SetVersion(v int32) *GetTemplateRequest {
	s.Version = &v
	return s
}

type GetTemplateResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Template          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateResponse) SetHeaders(v map[string]*string) *GetTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateResponse) SetStatusCode(v int32) *GetTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateResponse) SetBody(v *Template) *GetTemplateResponse {
	s.Body = v
	return s
}

type ListApplicationsRequest struct {
	CurrentPage *string `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	FilterName  *string `json:"filterName,omitempty" xml:"filterName,omitempty"`
	PageSize    *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Template    *string `json:"template,omitempty" xml:"template,omitempty"`
}

func (s ListApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsRequest) SetCurrentPage(v string) *ListApplicationsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListApplicationsRequest) SetFilterName(v string) *ListApplicationsRequest {
	s.FilterName = &v
	return s
}

func (s *ListApplicationsRequest) SetPageSize(v string) *ListApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsRequest) SetTemplate(v string) *ListApplicationsRequest {
	s.Template = &v
	return s
}

type ListApplicationsResponseBody struct {
	CurrentPage *string        `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	Result      []*Application `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount  *string        `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBody) SetCurrentPage(v string) *ListApplicationsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListApplicationsResponseBody) SetResult(v []*Application) *ListApplicationsResponseBody {
	s.Result = v
	return s
}

func (s *ListApplicationsResponseBody) SetTotalCount(v string) *ListApplicationsResponseBody {
	s.TotalCount = &v
	return s
}

type ListApplicationsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponse) SetHeaders(v map[string]*string) *ListApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsResponse) SetStatusCode(v int32) *ListApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationsResponse) SetBody(v *ListApplicationsResponseBody) *ListApplicationsResponse {
	s.Body = v
	return s
}

type ListEnvironmentRevisionsRequest struct {
	EnvironmentName *string `json:"environmentName,omitempty" xml:"environmentName,omitempty"`
}

func (s ListEnvironmentRevisionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentRevisionsRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentRevisionsRequest) SetEnvironmentName(v string) *ListEnvironmentRevisionsRequest {
	s.EnvironmentName = &v
	return s
}

type ListEnvironmentRevisionsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*EnvironmentRevision `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s ListEnvironmentRevisionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentRevisionsResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentRevisionsResponse) SetHeaders(v map[string]*string) *ListEnvironmentRevisionsResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentRevisionsResponse) SetStatusCode(v int32) *ListEnvironmentRevisionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnvironmentRevisionsResponse) SetBody(v []*EnvironmentRevision) *ListEnvironmentRevisionsResponse {
	s.Body = v
	return s
}

type ListEnvironmentsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*Environment     `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
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

func (s *ListEnvironmentsResponse) SetBody(v []*Environment) *ListEnvironmentsResponse {
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
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*PipelineTemplate `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*Pipeline        `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
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

type ListServiceRevisionsRequest struct {
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s ListServiceRevisionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceRevisionsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceRevisionsRequest) SetServiceName(v string) *ListServiceRevisionsRequest {
	s.ServiceName = &v
	return s
}

type ListServiceRevisionsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*ServiceRevision `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s ListServiceRevisionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceRevisionsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceRevisionsResponse) SetHeaders(v map[string]*string) *ListServiceRevisionsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceRevisionsResponse) SetStatusCode(v int32) *ListServiceRevisionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceRevisionsResponse) SetBody(v []*ServiceRevision) *ListServiceRevisionsResponse {
	s.Body = v
	return s
}

type ListServicesResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*Service         `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s ListServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponse) GoString() string {
	return s.String()
}

func (s *ListServicesResponse) SetHeaders(v map[string]*string) *ListServicesResponse {
	s.Headers = v
	return s
}

func (s *ListServicesResponse) SetStatusCode(v int32) *ListServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServicesResponse) SetBody(v []*Service) *ListServicesResponse {
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*TaskTemplate    `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*Task            `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
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

type ListTemplateRevisionsRequest struct {
	TemplateName    *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	TemplateVersion *int32  `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
}

func (s ListTemplateRevisionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateRevisionsRequest) GoString() string {
	return s.String()
}

func (s *ListTemplateRevisionsRequest) SetTemplateName(v string) *ListTemplateRevisionsRequest {
	s.TemplateName = &v
	return s
}

func (s *ListTemplateRevisionsRequest) SetTemplateVersion(v int32) *ListTemplateRevisionsRequest {
	s.TemplateVersion = &v
	return s
}

type ListTemplateRevisionsResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*TemplateRevision `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s ListTemplateRevisionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateRevisionsResponse) GoString() string {
	return s.String()
}

func (s *ListTemplateRevisionsResponse) SetHeaders(v map[string]*string) *ListTemplateRevisionsResponse {
	s.Headers = v
	return s
}

func (s *ListTemplateRevisionsResponse) SetStatusCode(v int32) *ListTemplateRevisionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTemplateRevisionsResponse) SetBody(v []*TemplateRevision) *ListTemplateRevisionsResponse {
	s.Body = v
	return s
}

type ListTemplatesRequest struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) SetType(v string) *ListTemplatesRequest {
	s.Type = &v
	return s
}

type ListTemplatesResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*Template        `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s ListTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponse) SetHeaders(v map[string]*string) *ListTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListTemplatesResponse) SetStatusCode(v int32) *ListTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTemplatesResponse) SetBody(v []*Template) *ListTemplatesResponse {
	s.Body = v
	return s
}

type PutEnvironmentRequest struct {
	Body *Environment `json:"body,omitempty" xml:"body,omitempty"`
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

type PutEnvironmentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Environment       `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Body  *Pipeline `json:"body,omitempty" xml:"body,omitempty"`
	Force *bool     `json:"force,omitempty" xml:"force,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Pipeline          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Body  *PipelineTemplate `json:"body,omitempty" xml:"body,omitempty"`
	Force *bool             `json:"force,omitempty" xml:"force,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PipelineTemplate  `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type PutServiceRequest struct {
	Body *Service `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s PutServiceRequest) GoString() string {
	return s.String()
}

func (s *PutServiceRequest) SetBody(v *Service) *PutServiceRequest {
	s.Body = v
	return s
}

type PutServiceResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Service           `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s PutServiceResponse) GoString() string {
	return s.String()
}

func (s *PutServiceResponse) SetHeaders(v map[string]*string) *PutServiceResponse {
	s.Headers = v
	return s
}

func (s *PutServiceResponse) SetStatusCode(v int32) *PutServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *PutServiceResponse) SetBody(v *Service) *PutServiceResponse {
	s.Body = v
	return s
}

type PutTaskStatusRequest struct {
	Body  *Task `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Task              `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Body  *TaskTemplate `json:"body,omitempty" xml:"body,omitempty"`
	Force *bool         `json:"force,omitempty" xml:"force,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TaskTemplate      `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type PutTemplateRequest struct {
	Body    *Template `json:"body,omitempty" xml:"body,omitempty"`
	Version *int32    `json:"version,omitempty" xml:"version,omitempty"`
}

func (s PutTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s PutTemplateRequest) GoString() string {
	return s.String()
}

func (s *PutTemplateRequest) SetBody(v *Template) *PutTemplateRequest {
	s.Body = v
	return s
}

func (s *PutTemplateRequest) SetVersion(v int32) *PutTemplateRequest {
	s.Version = &v
	return s
}

type PutTemplateResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Template          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s PutTemplateResponse) GoString() string {
	return s.String()
}

func (s *PutTemplateResponse) SetHeaders(v map[string]*string) *PutTemplateResponse {
	s.Headers = v
	return s
}

func (s *PutTemplateResponse) SetStatusCode(v int32) *PutTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *PutTemplateResponse) SetBody(v *Template) *PutTemplateResponse {
	s.Body = v
	return s
}

type ResumeTaskResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Task              `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Task              `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Pipeline          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Task              `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UpdateApplicationRequest struct {
	Body *Application `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRequest) SetBody(v *Application) *UpdateApplicationRequest {
	s.Body = v
	return s
}

type UpdateApplicationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Application       `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponse) SetHeaders(v map[string]*string) *UpdateApplicationResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationResponse) SetStatusCode(v int32) *UpdateApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationResponse) SetBody(v *Application) *UpdateApplicationResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("serverless"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CancelTaskWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CancelTask"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/cancel"),
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

func (client *Client) CreateApplicationWithOptions(request *CreateApplicationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoDeploy)) {
		body["autoDeploy"] = request.AutoDeploy
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EnvVars)) {
		body["envVars"] = request.EnvVars
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RepoSource)) {
		body["repoSource"] = request.RepoSource
	}

	if !tea.BoolValue(util.IsUnset(request.RoleArn)) {
		body["roleArn"] = request.RoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.Template)) {
		body["template"] = request.Template
	}

	if !tea.BoolValue(util.IsUnset(request.Trigger)) {
		body["trigger"] = request.Trigger
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApplication"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/applications"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApplication(request *CreateApplicationRequest) (_result *CreateApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateApplicationResponse{}
	_body, _err := client.CreateApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/pipelines"),
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
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/pipelinetemplates"),
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

func (client *Client) CreateReleaseWithOptions(appName *string, request *CreateReleaseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateReleaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["description"] = request.Description
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRelease"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/applications/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/releases"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateReleaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRelease(appName *string, request *CreateReleaseRequest) (_result *CreateReleaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateReleaseResponse{}
	_body, _err := client.CreateReleaseWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/tasks"),
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
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/tasktemplates"),
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

func (client *Client) DeleteApplicationWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApplication"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/applications/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("string"),
	}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApplication(name *string) (_result *DeleteApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DeleteApplicationWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEnvironmentWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteEnvironmentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEnvironment"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/environments/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
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

func (client *Client) DeleteEnvironment(name *string) (_result *DeleteEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEnvironmentResponse{}
	_body, _err := client.DeleteEnvironmentWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePipelineTemplateWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeletePipelineTemplateResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePipelineTemplate"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/pipelinetemplates/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
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

func (client *Client) DeleteTaskTemplateWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTaskTemplateResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTaskTemplate"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/tasktemplates/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
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

func (client *Client) DeleteTemplateWithOptions(name *string, request *DeleteTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTemplate"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/templates/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTemplate(name *string, request *DeleteTemplateRequest) (_result *DeleteTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetApplicationWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetApplicationResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetApplication"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/applications/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApplication(name *string) (_result *GetApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetApplicationResponse{}
	_body, _err := client.GetApplicationWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEnvironmentWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEnvironmentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetEnvironment"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/environments/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
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

func (client *Client) GetEnvironment(name *string) (_result *GetEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEnvironmentResponse{}
	_body, _err := client.GetEnvironmentWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPipelineWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPipelineResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipeline"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
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

func (client *Client) GetPipelineTemplateWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPipelineTemplateResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipelineTemplate"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/pipelinetemplates/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
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

func (client *Client) GetReleaseWithOptions(appName *string, versionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetReleaseResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetRelease"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/applications/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/releases/" + tea.StringValue(openapiutil.GetEncodeParam(versionId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetReleaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRelease(appName *string, versionId *string) (_result *GetReleaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetReleaseResponse{}
	_body, _err := client.GetReleaseWithOptions(appName, versionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetService"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/services/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetService(name *string) (_result *GetServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceResponse{}
	_body, _err := client.GetServiceWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTaskWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTask"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
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

func (client *Client) GetTaskTemplateWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskTemplateResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskTemplate"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/tasktemplates/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
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

func (client *Client) GetTemplateWithOptions(name *string, request *GetTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTemplate"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/templates/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTemplate(name *string, request *GetTemplateRequest) (_result *GetTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApplicationsWithOptions(request *ListApplicationsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["currentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.FilterName)) {
		query["filterName"] = request.FilterName
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Template)) {
		query["template"] = request.Template
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApplications"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/applications"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApplications(request *ListApplicationsRequest) (_result *ListApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListApplicationsResponse{}
	_body, _err := client.ListApplicationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEnvironmentRevisionsWithOptions(request *ListEnvironmentRevisionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEnvironmentRevisionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvironmentName)) {
		query["environmentName"] = request.EnvironmentName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEnvironmentRevisions"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/environmentrevisions/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &ListEnvironmentRevisionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEnvironmentRevisions(request *ListEnvironmentRevisionsRequest) (_result *ListEnvironmentRevisionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEnvironmentRevisionsResponse{}
	_body, _err := client.ListEnvironmentRevisionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEnvironmentsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEnvironmentsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListEnvironments"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/environments/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &ListEnvironmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEnvironments() (_result *ListEnvironmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEnvironmentsResponse{}
	_body, _err := client.ListEnvironmentsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/pipelinetemplates"),
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
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/pipelines"),
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

func (client *Client) ListServiceRevisionsWithOptions(request *ListServiceRevisionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListServiceRevisionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["serviceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceRevisions"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/servicerevisions/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &ListServiceRevisionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServiceRevisions(request *ListServiceRevisionsRequest) (_result *ListServiceRevisionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServiceRevisionsResponse{}
	_body, _err := client.ListServiceRevisionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServicesWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListServices"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/services/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServices() (_result *ListServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServicesResponse{}
	_body, _err := client.ListServicesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/tasktemplates"),
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
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/tasks"),
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

func (client *Client) ListTemplateRevisionsWithOptions(request *ListTemplateRevisionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTemplateRevisionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["templateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["templateVersion"] = request.TemplateVersion
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTemplateRevisions"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/templaterevisions/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &ListTemplateRevisionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTemplateRevisions(request *ListTemplateRevisionsRequest) (_result *ListTemplateRevisionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTemplateRevisionsResponse{}
	_body, _err := client.ListTemplateRevisionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTemplatesWithOptions(request *ListTemplatesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTemplates"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/templates/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &ListTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTemplates(request *ListTemplatesRequest) (_result *ListTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTemplatesResponse{}
	_body, _err := client.ListTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutEnvironmentWithOptions(name *string, request *PutEnvironmentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutEnvironmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutEnvironment"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/environments/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
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

func (client *Client) PutEnvironment(name *string, request *PutEnvironmentRequest) (_result *PutEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutEnvironmentResponse{}
	_body, _err := client.PutEnvironmentWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/status"),
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
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/pipelinetemplates/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
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

func (client *Client) PutServiceWithOptions(name *string, request *PutServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutService"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/services/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutService(name *string, request *PutServiceRequest) (_result *PutServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutServiceResponse{}
	_body, _err := client.PutServiceWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/status"),
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
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/tasktemplates/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
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

func (client *Client) PutTemplateWithOptions(name *string, request *PutTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutTemplate"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/templates/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutTemplate(name *string, request *PutTemplateRequest) (_result *PutTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutTemplateResponse{}
	_body, _err := client.PutTemplateWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeTaskWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ResumeTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeTask"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/resume"),
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

func (client *Client) RetryTaskWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RetryTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RetryTask"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/retry"),
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

func (client *Client) StartPipelineWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartPipelineResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartPipeline"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/start"),
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

func (client *Client) StartTaskWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartTask"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/start"),
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

func (client *Client) UpdateApplicationWithOptions(name *string, request *UpdateApplicationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApplication"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/applications/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApplication(name *string, request *UpdateApplicationRequest) (_result *UpdateApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateApplicationResponse{}
	_body, _err := client.UpdateApplicationWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
