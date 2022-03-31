// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Application struct {
	// 是否直接跳过plan直接进行发布
	AutoDeploy *string `json:"autoDeploy,omitempty" xml:"autoDeploy,omitempty"`
	// 应用创建时间
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 应用描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 环境变量
	EnvVars map[string]interface{} `json:"envVars,omitempty" xml:"envVars,omitempty"`
	// 指定本地代码源
	LocalSource *string `json:"localSource,omitempty" xml:"localSource,omitempty"`
	// 应用名称，同账号下唯一，创建后不允许变更
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 指定OSS的代码源
	OssSource *ApplicationOssSource `json:"ossSource,omitempty" xml:"ossSource,omitempty" type:"Struct"`
	// 应用参数，schema由应用模板所定义
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// 指定仓库的代码源
	RepoSource *ApplicationRepoSource `json:"repoSource,omitempty" xml:"repoSource,omitempty" type:"Struct"`
	// 指定role进行角色扮演
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// 关联的模板，用于Web应用、模板应用的创建
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
	// 触发配置，不指定表示手动触发
	Trigger *ApplicationTrigger `json:"trigger,omitempty" xml:"trigger,omitempty" type:"Struct"`
	// 阿里云主账号ID，只读
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// 应用更新时间
	UpdatedTime *string `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
	// s.yaml所在目录，不指定则默认使用当前目录
	WorkDir *string `json:"workDir,omitempty" xml:"workDir,omitempty"`
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

func (s *Application) SetLocalSource(v string) *Application {
	s.LocalSource = &v
	return s
}

func (s *Application) SetName(v string) *Application {
	s.Name = &v
	return s
}

func (s *Application) SetOssSource(v *ApplicationOssSource) *Application {
	s.OssSource = v
	return s
}

func (s *Application) SetParameter(v map[string]interface{}) *Application {
	s.Parameter = v
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

func (s *Application) SetUid(v string) *Application {
	s.Uid = &v
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

type ApplicationOssSource struct {
	// OSS Bucket名字
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// OSS Object名字
	ObjectName *string `json:"objectName,omitempty" xml:"objectName,omitempty"`
}

func (s ApplicationOssSource) String() string {
	return tea.Prettify(s)
}

func (s ApplicationOssSource) GoString() string {
	return s.String()
}

func (s *ApplicationOssSource) SetBucketName(v string) *ApplicationOssSource {
	s.BucketName = &v
	return s
}

func (s *ApplicationOssSource) SetObjectName(v string) *ApplicationOssSource {
	s.ObjectName = &v
	return s
}

type ApplicationRepoSource struct {
	// 代码库owner
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// 代码源VCS
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// 代码库名称
	Repo *string `json:"repo,omitempty" xml:"repo,omitempty"`
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
	// 代码分支
	Branch []*string `json:"branch,omitempty" xml:"branch,omitempty" type:"Repeated"`
	// 触发条件
	On *string `json:"on,omitempty" xml:"on,omitempty"`
	// 代码路径。指定时，只有当匹配的path变化才触发
	Paths []*string `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
}

func (s ApplicationTrigger) String() string {
	return tea.Prettify(s)
}

func (s ApplicationTrigger) GoString() string {
	return s.String()
}

func (s *ApplicationTrigger) SetBranch(v []*string) *ApplicationTrigger {
	s.Branch = v
	return s
}

func (s *ApplicationTrigger) SetOn(v string) *ApplicationTrigger {
	s.On = &v
	return s
}

func (s *ApplicationTrigger) SetPaths(v []*string) *ApplicationTrigger {
	s.Paths = v
	return s
}

type Environment struct {
	// A time representing the server time when this object was created. Clients may not set this value. Populated by the system. Read-only.
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// Date and time at which a deletion is requested by the user. Null when the resource has not been requested for deletion. This field is set by the server, not directly settable by a client. Populated by the system. Read-only.
	DeletionTime *string `json:"deletionTime,omitempty" xml:"deletionTime,omitempty"`
	// Human-readable description of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// A sequence number representing a specific generation of the desired state. Populated by the system. Read-only.
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// The kind of the resource
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// Name must be unique within a namespace. Is required when creating resources. Cannot be updated.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specification of the desired behavior of the Environment.
	Spec *EnvironmentSpec `json:"spec,omitempty" xml:"spec,omitempty"`
	// Most recently observed status of the Environment. This data may not be up-to-date. Populated by the system. Read-only.
	Status *EnvironmentStatus `json:"status,omitempty" xml:"status,omitempty"`
	// Main user ID of an Aliyun account
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
	// A time representing the server time when this object was created. Clients may not set this value. Populated by the system. Read-only.
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// The generation of the environment.
	EnvironmentGeneration *int32 `json:"environmentGeneration,omitempty" xml:"environmentGeneration,omitempty"`
	// The name of an environment.
	EnvironmentName *string `json:"environmentName,omitempty" xml:"environmentName,omitempty"`
	// The kind of the resource.
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// Specification of the desired behavior of the Environment.
	Spec *EnvironmentSpec `json:"spec,omitempty" xml:"spec,omitempty"`
	// Most recently observed status of the Environment. This data may not be up-to-date. Populated by the system. Read-only.
	Status *EnvironmentStatus `json:"status,omitempty" xml:"status,omitempty"`
	// Main user ID of an Aliyun account.
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
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
	// A region ID at Aliyun. For example, "cn-hangzhou"
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ARN (Aliyun Resource Name) of the role designated by a user to allow the system to manage his Aliyun resources. If null, use roleArn of role AliyunFCDefaultRole.
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// The name of the template for the Environment
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
	// Variables for specified template
	TemplateVariables map[string]interface{} `json:"templateVariables,omitempty" xml:"templateVariables,omitempty"`
	// The major version of the template. "1" by default.
	TemplateVersion *int32 `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
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
	// A human-readable message indicating details about why the Environment is in this condition
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The most recent generation observed
	ObservedGeneration *int32 `json:"observedGeneration,omitempty" xml:"observedGeneration,omitempty"`
	// Time when the last update of the status is observed
	ObservedTime *string `json:"observedTime,omitempty" xml:"observedTime,omitempty"`
	// Details of current state of the Environment
	Output map[string]interface{} `json:"output,omitempty" xml:"output,omitempty"`
	// A simple, high-level summary of where the Environment is in its lifecycle
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
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
	// A default value (as JSON string) which then makes the variable optional.
	DefaultJson *string `json:"defaultJson,omitempty" xml:"defaultJson,omitempty"`
	// This specifies the input variable's documentation.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The identifier of an input variable. Identifiers can contain letters, digits, underscores (_), and hyphens (-). The first character of an identifier must not be a digit, to avoid ambiguity with literal numbers.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specify if the variable can be null. True by default.
	Nullable *bool `json:"nullable,omitempty" xml:"nullable,omitempty"`
	// Specify if the variable contains sensitive material. False by default.
	Sensitive *bool `json:"sensitive,omitempty" xml:"sensitive,omitempty"`
	// This argument specifies what value types are accepted for the variable.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
	// The description should concisely explain the purpose of the output and what kind of value is expected.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The identifier of an output value. Identifiers can contain letters, digits, underscores (_), and hyphens (-). The first character of an identifier must not be a digit, to avoid ambiguity with literal numbers.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specify if the output value contains sensitive material. False by default.
	Sensitive *bool `json:"sensitive,omitempty" xml:"sensitive,omitempty"`
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

type Release struct {
	// 本次发布的应用快照，只读
	AppConfig map[string]interface{} `json:"appConfig,omitempty" xml:"appConfig,omitempty"`
	// 代码版本，不指定则使用最新的commit
	CodeVersion *ReleaseCodeVersion `json:"codeVersion,omitempty" xml:"codeVersion,omitempty" type:"Struct"`
	// 创建时间，只读
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 本地发布描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 完成时间，只读
	FinishedTime *string `json:"finishedTime,omitempty" xml:"finishedTime,omitempty"`
	// 本次发布的输出，只读
	Output map[string]interface{} `json:"output,omitempty" xml:"output,omitempty"`
	// 流水线配置，不指定则直接部署
	Pipeline *ReleasePipeline `json:"pipeline,omitempty" xml:"pipeline,omitempty" type:"Struct"`
	// 本地发布状态：published: 发布完成 publishing：发布中 failed：发布失败 canceled：已撤销
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 阿里云主账号ID，只读
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// 本次发布版本号，由系统生成，只读
	VersionId *int64 `json:"versionId,omitempty" xml:"versionId,omitempty"`
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

func (s *Release) SetFinishedTime(v string) *Release {
	s.FinishedTime = &v
	return s
}

func (s *Release) SetOutput(v map[string]interface{}) *Release {
	s.Output = v
	return s
}

func (s *Release) SetPipeline(v *ReleasePipeline) *Release {
	s.Pipeline = v
	return s
}

func (s *Release) SetStatus(v string) *Release {
	s.Status = &v
	return s
}

func (s *Release) SetUid(v string) *Release {
	s.Uid = &v
	return s
}

func (s *Release) SetVersionId(v int64) *Release {
	s.VersionId = &v
	return s
}

type ReleaseCodeVersion struct {
	// 代码分支，不指定则使用default分支
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// commit id
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

type ReleasePipeline struct {
	// stage配置
	Stages []*ReleasePipelineStages `json:"stages,omitempty" xml:"stages,omitempty" type:"Repeated"`
}

func (s ReleasePipeline) String() string {
	return tea.Prettify(s)
}

func (s ReleasePipeline) GoString() string {
	return s.String()
}

func (s *ReleasePipeline) SetStages(v []*ReleasePipelineStages) *ReleasePipeline {
	s.Stages = v
	return s
}

type ReleasePipelineStages struct {
	// 执行环境
	Environment *string `json:"environment,omitempty" xml:"environment,omitempty"`
	// stage名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ReleasePipelineStages) String() string {
	return tea.Prettify(s)
}

func (s ReleasePipelineStages) GoString() string {
	return s.String()
}

func (s *ReleasePipelineStages) SetEnvironment(v string) *ReleasePipelineStages {
	s.Environment = &v
	return s
}

func (s *ReleasePipelineStages) SetName(v string) *ReleasePipelineStages {
	s.Name = &v
	return s
}

type Service struct {
	// A time representing the server time when this object was created. Clients may not set this value. Populated by the system. Read-only.
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// Date and time at which a deletion is requested by the user. Null when the resource has not been requested for deletion. This field is set by the server, not directly settable by a client. Populated by the system. Read-only.
	DeletionTime *string `json:"deletionTime,omitempty" xml:"deletionTime,omitempty"`
	// Human-readable description of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// A sequence number representing a specific generation of the desired state. Populated by the system. Read-only.
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// The kind of the resource
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// Name must be unique within a namespace. Is required when creating resources. Cannot be updated.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specification of the desired behavior of the Service.
	Spec *ServiceSpec `json:"spec,omitempty" xml:"spec,omitempty"`
	// Most recently observed status of the Service. This data may not be up-to-date. Populated by the system. Read-only.
	Status *ServiceStatus `json:"status,omitempty" xml:"status,omitempty"`
	// Main user ID of an Aliyun account
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
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
	// A time representing the server time when this object was created. Clients may not set this value. Populated by the system. Read-only.
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// The kind of the resource.
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The generation of the service.
	ServiceGeneration *int32 `json:"serviceGeneration,omitempty" xml:"serviceGeneration,omitempty"`
	// The name of a service.
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// Specification of the desired behavior of the Service.
	Spec *ServiceSpec `json:"spec,omitempty" xml:"spec,omitempty"`
	// Most recently observed status of the Environment. This data may not be up-to-date. Populated by the system. Read-only.
	Status *EnvironmentStatus `json:"status,omitempty" xml:"status,omitempty"`
	// Main user ID of an Aliyun account.
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
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
	// The name of the associated Environment for the Service
	Environment *string `json:"environment,omitempty" xml:"environment,omitempty"`
	// The ARN (Aliyun Resource Name) of the role designated by a user to allow the system to manage his Aliyun resources. If null, use roleArn of role AliyunFCDefaultRole.
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// The name of the template for the Service
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
	// Variables for specified template
	TemplateVariables map[string]interface{} `json:"templateVariables,omitempty" xml:"templateVariables,omitempty"`
	// The major version of the template. "1" by default.
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

type ServiceStatus struct {
	// A human-readable message indicating details about why the Service is in this condition
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The most recent generation observed
	ObservedGeneration *int32 `json:"observedGeneration,omitempty" xml:"observedGeneration,omitempty"`
	// Time when the last update of the status is observed
	ObservedTime *string `json:"observedTime,omitempty" xml:"observedTime,omitempty"`
	// Details of current state of the Service
	Output map[string]interface{} `json:"output,omitempty" xml:"output,omitempty"`
	// A simple, high-level summary of where the Service is in its lifecycle
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
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
	// A machine-readable description of why this operation is in the failure status. If this value is empty there is no information available.
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// A human-readable description of the status of this operation.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// ID of the request. May be null.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Whether the operation is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	// Access key ID
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// Expiration time of the credentials
	ExpirationTime *string `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	// The kind of the credentials
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// Secret access key
	SecretAccessKey *string `json:"secretAccessKey,omitempty" xml:"secretAccessKey,omitempty"`
	// Token
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
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

type Template struct {
	// A time representing the server time when this object was created. Clients may not set this value. Populated by the system. Read-only.
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// Date and time at which a deletion is requested by the user. Null when the resource has not been requested for deletion. This field is set by the server, not directly settable by a client. Populated by the system. Read-only.
	DeletionTime *string `json:"deletionTime,omitempty" xml:"deletionTime,omitempty"`
	// Human-readable description of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// A sequence number representing a specific generation of the desired state. Populated by the system. Read-only.
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// The kind of the resource
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// Name must be unique within a namespace. Is required when creating resources. Cannot be updated.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specification of the desired behavior of the Template.
	Spec *TemplateSpec `json:"spec,omitempty" xml:"spec,omitempty"`
	// Most recently observed status of the Template. This data may not be up-to-date. Populated by the system. Read-only.
	Status *TemplateStatus `json:"status,omitempty" xml:"status,omitempty"`
	// Main user ID of an Aliyun account
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// The major version of the template. "1" by default. You should ONLY increment the major version when the template are not backwards compatible with the previous major version.
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
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
	// A time representing the server time when this object was created. Clients may not set this value. Populated by the system. Read-only.
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// The kind of the resource.
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// Specification of the desired behavior of the Template.
	Spec *TemplateSpec `json:"spec,omitempty" xml:"spec,omitempty"`
	// Most recently observed status of the Template. This data may not be up-to-date. Populated by the system. Read-only.
	Status *TemplateStatus `json:"status,omitempty" xml:"status,omitempty"`
	// The generation of the template.
	TemplateGeneration *int32 `json:"templateGeneration,omitempty" xml:"templateGeneration,omitempty"`
	// The name of a template.
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// The version of a template.
	TemplateVersion *int32 `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
	// Main user ID of an Aliyun account.
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
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
	// The raw content of the template.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The media type of the template content. At the moment, only "application/hcl+terraform" is supported.
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The content of RAM policy  required for this template.
	RamPolicy *string `json:"ramPolicy,omitempty" xml:"ramPolicy,omitempty"`
	// The type of the applicable resource for this template. Must be either "Environment" or "Service".
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
	// A human-readable message indicating details about why the Template is in this condition.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The most recent generation observed.
	ObservedGeneration *int32 `json:"observedGeneration,omitempty" xml:"observedGeneration,omitempty"`
	// Time when the last update of the status is observed.
	ObservedTime *string `json:"observedTime,omitempty" xml:"observedTime,omitempty"`
	// The definition of output values of the template parsed from the template content.
	Outputs []*OutputValue `json:"outputs,omitempty" xml:"outputs,omitempty" type:"Repeated"`
	// A simple, high-level summary of where the Template is in its lifecycle.
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
	// The definition of input variables of the template parsed from the template content.
	Variables []*InputVariable `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
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

type CreateApplicationRequest struct {
	Body *Application `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) SetBody(v *Application) *CreateApplicationRequest {
	s.Body = v
	return s
}

type CreateApplicationResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *Application       `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateApplicationResponse) SetBody(v *Application) *CreateApplicationResponse {
	s.Body = v
	return s
}

type CreateReleaseRequest struct {
	Body *Release `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateReleaseRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateReleaseRequest) GoString() string {
	return s.String()
}

func (s *CreateReleaseRequest) SetBody(v *Release) *CreateReleaseRequest {
	s.Body = v
	return s
}

type CreateReleaseResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *Release           `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateReleaseResponse) SetBody(v *Release) *CreateReleaseResponse {
	s.Body = v
	return s
}

type DeleteApplicationResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *string            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteApplicationResponse) SetBody(v string) *DeleteApplicationResponse {
	s.Body = &v
	return s
}

type DeleteTemplateRequest struct {
	// The major version of the template. "1" by default.
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
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *Status            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteTemplateResponse) SetBody(v *Status) *DeleteTemplateResponse {
	s.Body = v
	return s
}

type GetApplicationResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *Application       `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetApplicationResponse) SetBody(v *Application) *GetApplicationResponse {
	s.Body = v
	return s
}

type GetEnvironmentResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *Environment       `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetEnvironmentResponse) SetBody(v *Environment) *GetEnvironmentResponse {
	s.Body = v
	return s
}

type GetReleaseResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *Release           `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetReleaseResponse) SetBody(v *Release) *GetReleaseResponse {
	s.Body = v
	return s
}

type GetServiceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *Service           `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetServiceResponse) SetBody(v *Service) *GetServiceResponse {
	s.Body = v
	return s
}

type GetTemplateRequest struct {
	// The major version of the template. "1" by default.
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
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *Template          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTemplateResponse) SetBody(v *Template) *GetTemplateResponse {
	s.Body = v
	return s
}

type ListEnvironmentRevisionsRequest struct {
	// The name of an environment.
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    []*EnvironmentRevision `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
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

func (s *ListEnvironmentRevisionsResponse) SetBody(v []*EnvironmentRevision) *ListEnvironmentRevisionsResponse {
	s.Body = v
	return s
}

type ListEnvironmentsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    []*Environment     `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
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

func (s *ListEnvironmentsResponse) SetBody(v []*Environment) *ListEnvironmentsResponse {
	s.Body = v
	return s
}

type ListServiceRevisionsRequest struct {
	// The name of a service.
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
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    []*ServiceRevision `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
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

func (s *ListServiceRevisionsResponse) SetBody(v []*ServiceRevision) *ListServiceRevisionsResponse {
	s.Body = v
	return s
}

type ListServicesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    []*Service         `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
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

func (s *ListServicesResponse) SetBody(v []*Service) *ListServicesResponse {
	s.Body = v
	return s
}

type ListTemplateRevisionsRequest struct {
	// The name of a template.
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// The major version of the template. "1" by default.
	TemplateVersion *int32 `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
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
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    []*TemplateRevision `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
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

func (s *ListTemplateRevisionsResponse) SetBody(v []*TemplateRevision) *ListTemplateRevisionsResponse {
	s.Body = v
	return s
}

type ListTemplatesRequest struct {
	// The type of the applicable resource for this template. Must be either "Environment" or "Service".
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
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    []*Template        `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
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

func (s *ListTemplatesResponse) SetBody(v []*Template) *ListTemplatesResponse {
	s.Body = v
	return s
}

type PutEnvironmentRequest struct {
	// An environment for serverless deployments
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
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *Environment       `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PutEnvironmentResponse) SetBody(v *Environment) *PutEnvironmentResponse {
	s.Body = v
	return s
}

type PutServiceRequest struct {
	// A service for serverless deployments
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
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *Service           `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PutServiceResponse) SetBody(v *Service) *PutServiceResponse {
	s.Body = v
	return s
}

type PutTemplateRequest struct {
	// A custom template
	Body *Template `json:"body,omitempty" xml:"body,omitempty"`
	// The major version of the template. "1" by default.
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
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
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *Template          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PutTemplateResponse) SetBody(v *Template) *PutTemplateResponse {
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
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *Application       `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (client *Client) CreateApplicationWithOptions(request *CreateApplicationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.Body)),
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

func (client *Client) CreateReleaseWithOptions(appName *string, request *CreateReleaseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateReleaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appName = openapiutil.GetEncodeParam(appName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.Body)),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRelease"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/applications/" + tea.StringValue(appName) + "/releases"),
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

func (client *Client) DeleteApplicationWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApplication"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/applications/" + tea.StringValue(name)),
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

func (client *Client) DeleteTemplateWithOptions(name *string, request *DeleteTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	name = openapiutil.GetEncodeParam(name)
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
		Pathname:    tea.String("/apis/serverlessdeployment/v1/templates/" + tea.StringValue(name)),
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

func (client *Client) GetApplicationWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetApplicationResponse, _err error) {
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetApplication"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/applications/" + tea.StringValue(name)),
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

func (client *Client) GetEnvironmentWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEnvironmentResponse, _err error) {
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetEnvironment"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/environments/" + tea.StringValue(name)),
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

func (client *Client) GetReleaseWithOptions(appName *string, versionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetReleaseResponse, _err error) {
	appName = openapiutil.GetEncodeParam(appName)
	versionId = openapiutil.GetEncodeParam(versionId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetRelease"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/applications/" + tea.StringValue(appName) + "/releases/" + tea.StringValue(versionId)),
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

func (client *Client) GetServiceWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetServiceResponse, _err error) {
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetService"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/services/" + tea.StringValue(name)),
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

func (client *Client) GetTemplateWithOptions(name *string, request *GetTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	name = openapiutil.GetEncodeParam(name)
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
		Pathname:    tea.String("/apis/serverlessdeployment/v1/templates/" + tea.StringValue(name)),
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

func (client *Client) PutEnvironmentWithOptions(name *string, request *PutEnvironmentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutEnvironmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.Body)),
	}
	params := &openapi.Params{
		Action:      tea.String("PutEnvironment"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/environments/" + tea.StringValue(name)),
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

func (client *Client) PutServiceWithOptions(name *string, request *PutServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.Body)),
	}
	params := &openapi.Params{
		Action:      tea.String("PutService"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/services/" + tea.StringValue(name)),
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

func (client *Client) PutTemplateWithOptions(name *string, request *PutTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	name = openapiutil.GetEncodeParam(name)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(tea.ToMap(request.Body)),
	}
	params := &openapi.Params{
		Action:      tea.String("PutTemplate"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/templates/" + tea.StringValue(name)),
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

func (client *Client) UpdateApplicationWithOptions(name *string, request *UpdateApplicationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.Body)),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApplication"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/applications/" + tea.StringValue(name)),
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
