// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Artifact struct {
	// This parameter is required.
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	Creator    *int64      `json:"creator,omitempty" xml:"creator,omitempty"`
	Credential *Credential `json:"credential,omitempty" xml:"credential,omitempty"`
	// This parameter is required.
	GmtCreated *string `json:"gmtCreated,omitempty" xml:"gmtCreated,omitempty"`
	// This parameter is required.
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// This parameter is required.
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// This parameter is required.
	Modifier *int64 `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s Artifact) String() string {
	return tea.Prettify(s)
}

func (s Artifact) GoString() string {
	return s.String()
}

func (s *Artifact) SetBizId(v string) *Artifact {
	s.BizId = &v
	return s
}

func (s *Artifact) SetCreator(v int64) *Artifact {
	s.Creator = &v
	return s
}

func (s *Artifact) SetCredential(v *Credential) *Artifact {
	s.Credential = v
	return s
}

func (s *Artifact) SetGmtCreated(v string) *Artifact {
	s.GmtCreated = &v
	return s
}

func (s *Artifact) SetGmtModified(v string) *Artifact {
	s.GmtModified = &v
	return s
}

func (s *Artifact) SetLocation(v string) *Artifact {
	s.Location = &v
	return s
}

func (s *Artifact) SetModifier(v int64) *Artifact {
	s.Modifier = &v
	return s
}

func (s *Artifact) SetName(v string) *Artifact {
	s.Name = &v
	return s
}

type Category struct {
	// This parameter is required.
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	Creator *int64 `json:"creator,omitempty" xml:"creator,omitempty"`
	// This parameter is required.
	GmtCreated *string `json:"gmtCreated,omitempty" xml:"gmtCreated,omitempty"`
	// This parameter is required.
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// This parameter is required.
	Modifier *int64 `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// This parameter is required.
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	ParentBizId *string `json:"parentBizId,omitempty" xml:"parentBizId,omitempty"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Category) String() string {
	return tea.Prettify(s)
}

func (s Category) GoString() string {
	return s.String()
}

func (s *Category) SetBizId(v string) *Category {
	s.BizId = &v
	return s
}

func (s *Category) SetCreator(v int64) *Category {
	s.Creator = &v
	return s
}

func (s *Category) SetGmtCreated(v string) *Category {
	s.GmtCreated = &v
	return s
}

func (s *Category) SetGmtModified(v string) *Category {
	s.GmtModified = &v
	return s
}

func (s *Category) SetModifier(v int64) *Category {
	s.Modifier = &v
	return s
}

func (s *Category) SetName(v string) *Category {
	s.Name = &v
	return s
}

func (s *Category) SetParentBizId(v string) *Category {
	s.ParentBizId = &v
	return s
}

func (s *Category) SetType(v string) *Category {
	s.Type = &v
	return s
}

type Configuration struct {
	ConfigFileName  *string `json:"configFileName,omitempty" xml:"configFileName,omitempty"`
	ConfigItemKey   *string `json:"configItemKey,omitempty" xml:"configItemKey,omitempty"`
	ConfigItemValue *string `json:"configItemValue,omitempty" xml:"configItemValue,omitempty"`
}

func (s Configuration) String() string {
	return tea.Prettify(s)
}

func (s Configuration) GoString() string {
	return s.String()
}

func (s *Configuration) SetConfigFileName(v string) *Configuration {
	s.ConfigFileName = &v
	return s
}

func (s *Configuration) SetConfigItemKey(v string) *Configuration {
	s.ConfigItemKey = &v
	return s
}

func (s *Configuration) SetConfigItemValue(v string) *Configuration {
	s.ConfigItemValue = &v
	return s
}

type ConfigurationOverrides struct {
	Configurations []*ConfigurationOverridesConfigurations `json:"configurations,omitempty" xml:"configurations,omitempty" type:"Repeated"`
}

func (s ConfigurationOverrides) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationOverrides) GoString() string {
	return s.String()
}

func (s *ConfigurationOverrides) SetConfigurations(v []*ConfigurationOverridesConfigurations) *ConfigurationOverrides {
	s.Configurations = v
	return s
}

type ConfigurationOverridesConfigurations struct {
	ConfigFileName  *string `json:"configFileName,omitempty" xml:"configFileName,omitempty"`
	ConfigItemKey   *string `json:"configItemKey,omitempty" xml:"configItemKey,omitempty"`
	ConfigItemValue *string `json:"configItemValue,omitempty" xml:"configItemValue,omitempty"`
}

func (s ConfigurationOverridesConfigurations) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationOverridesConfigurations) GoString() string {
	return s.String()
}

func (s *ConfigurationOverridesConfigurations) SetConfigFileName(v string) *ConfigurationOverridesConfigurations {
	s.ConfigFileName = &v
	return s
}

func (s *ConfigurationOverridesConfigurations) SetConfigItemKey(v string) *ConfigurationOverridesConfigurations {
	s.ConfigItemKey = &v
	return s
}

func (s *ConfigurationOverridesConfigurations) SetConfigItemValue(v string) *ConfigurationOverridesConfigurations {
	s.ConfigItemValue = &v
	return s
}

type Credential struct {
	// This parameter is required.
	AccessId *string `json:"accessId,omitempty" xml:"accessId,omitempty"`
	// This parameter is required.
	Dir *string `json:"dir,omitempty" xml:"dir,omitempty"`
	// This parameter is required.
	Expire *string `json:"expire,omitempty" xml:"expire,omitempty"`
	// This parameter is required.
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// This parameter is required.
	Policy *string `json:"policy,omitempty" xml:"policy,omitempty"`
	// This parameter is required.
	SecurityToken *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
	// This parameter is required.
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s Credential) String() string {
	return tea.Prettify(s)
}

func (s Credential) GoString() string {
	return s.String()
}

func (s *Credential) SetAccessId(v string) *Credential {
	s.AccessId = &v
	return s
}

func (s *Credential) SetDir(v string) *Credential {
	s.Dir = &v
	return s
}

func (s *Credential) SetExpire(v string) *Credential {
	s.Expire = &v
	return s
}

func (s *Credential) SetHost(v string) *Credential {
	s.Host = &v
	return s
}

func (s *Credential) SetPolicy(v string) *Credential {
	s.Policy = &v
	return s
}

func (s *Credential) SetSecurityToken(v string) *Credential {
	s.SecurityToken = &v
	return s
}

func (s *Credential) SetSignature(v string) *Credential {
	s.Signature = &v
	return s
}

type JobDriver struct {
	SparkSubmit *JobDriverSparkSubmit `json:"sparkSubmit,omitempty" xml:"sparkSubmit,omitempty" type:"Struct"`
}

func (s JobDriver) String() string {
	return tea.Prettify(s)
}

func (s JobDriver) GoString() string {
	return s.String()
}

func (s *JobDriver) SetSparkSubmit(v *JobDriverSparkSubmit) *JobDriver {
	s.SparkSubmit = v
	return s
}

type JobDriverSparkSubmit struct {
	EntryPoint            *string   `json:"entryPoint,omitempty" xml:"entryPoint,omitempty"`
	EntryPointArguments   []*string `json:"entryPointArguments,omitempty" xml:"entryPointArguments,omitempty" type:"Repeated"`
	SparkSubmitParameters *string   `json:"sparkSubmitParameters,omitempty" xml:"sparkSubmitParameters,omitempty"`
}

func (s JobDriverSparkSubmit) String() string {
	return tea.Prettify(s)
}

func (s JobDriverSparkSubmit) GoString() string {
	return s.String()
}

func (s *JobDriverSparkSubmit) SetEntryPoint(v string) *JobDriverSparkSubmit {
	s.EntryPoint = &v
	return s
}

func (s *JobDriverSparkSubmit) SetEntryPointArguments(v []*string) *JobDriverSparkSubmit {
	s.EntryPointArguments = v
	return s
}

func (s *JobDriverSparkSubmit) SetSparkSubmitParameters(v string) *JobDriverSparkSubmit {
	s.SparkSubmitParameters = &v
	return s
}

type PrincipalAction struct {
	// example:
	//
	// acs:emr::workspaceId:action/create_queue
	ActionArn *string `json:"actionArn,omitempty" xml:"actionArn,omitempty"`
	// example:
	//
	// acs:emr::workspaceId:user/237593691541622267
	PrincipalArn *string `json:"principalArn,omitempty" xml:"principalArn,omitempty"`
}

func (s PrincipalAction) String() string {
	return tea.Prettify(s)
}

func (s PrincipalAction) GoString() string {
	return s.String()
}

func (s *PrincipalAction) SetActionArn(v string) *PrincipalAction {
	s.ActionArn = &v
	return s
}

func (s *PrincipalAction) SetPrincipalArn(v string) *PrincipalAction {
	s.PrincipalArn = &v
	return s
}

type ReleaseVersionImage struct {
	CpuArchitecture   *string `json:"cpuArchitecture,omitempty" xml:"cpuArchitecture,omitempty"`
	ImageId           *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
	RuntimeEngineType *string `json:"runtimeEngineType,omitempty" xml:"runtimeEngineType,omitempty"`
}

func (s ReleaseVersionImage) String() string {
	return tea.Prettify(s)
}

func (s ReleaseVersionImage) GoString() string {
	return s.String()
}

func (s *ReleaseVersionImage) SetCpuArchitecture(v string) *ReleaseVersionImage {
	s.CpuArchitecture = &v
	return s
}

func (s *ReleaseVersionImage) SetImageId(v string) *ReleaseVersionImage {
	s.ImageId = &v
	return s
}

func (s *ReleaseVersionImage) SetRuntimeEngineType(v string) *ReleaseVersionImage {
	s.RuntimeEngineType = &v
	return s
}

type RunLog struct {
	DriverStartup  *string `json:"driverStartup,omitempty" xml:"driverStartup,omitempty"`
	DriverStdError *string `json:"driverStdError,omitempty" xml:"driverStdError,omitempty"`
	DriverStdOut   *string `json:"driverStdOut,omitempty" xml:"driverStdOut,omitempty"`
	DriverSyslog   *string `json:"driverSyslog,omitempty" xml:"driverSyslog,omitempty"`
}

func (s RunLog) String() string {
	return tea.Prettify(s)
}

func (s RunLog) GoString() string {
	return s.String()
}

func (s *RunLog) SetDriverStartup(v string) *RunLog {
	s.DriverStartup = &v
	return s
}

func (s *RunLog) SetDriverStdError(v string) *RunLog {
	s.DriverStdError = &v
	return s
}

func (s *RunLog) SetDriverStdOut(v string) *RunLog {
	s.DriverStdOut = &v
	return s
}

func (s *RunLog) SetDriverSyslog(v string) *RunLog {
	s.DriverSyslog = &v
	return s
}

type SparkConf struct {
	// This parameter is required.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s SparkConf) String() string {
	return tea.Prettify(s)
}

func (s SparkConf) GoString() string {
	return s.String()
}

func (s *SparkConf) SetKey(v string) *SparkConf {
	s.Key = &v
	return s
}

func (s *SparkConf) SetValue(v string) *SparkConf {
	s.Value = &v
	return s
}

type SqlOutput struct {
	Rows   []*SqlOutputRows `json:"rows,omitempty" xml:"rows,omitempty" type:"Repeated"`
	Schema *SqlOutputSchema `json:"schema,omitempty" xml:"schema,omitempty" type:"Struct"`
}

func (s SqlOutput) String() string {
	return tea.Prettify(s)
}

func (s SqlOutput) GoString() string {
	return s.String()
}

func (s *SqlOutput) SetRows(v []*SqlOutputRows) *SqlOutput {
	s.Rows = v
	return s
}

func (s *SqlOutput) SetSchema(v *SqlOutputSchema) *SqlOutput {
	s.Schema = v
	return s
}

type SqlOutputRows struct {
	// example:
	//
	// null
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s SqlOutputRows) String() string {
	return tea.Prettify(s)
}

func (s SqlOutputRows) GoString() string {
	return s.String()
}

func (s *SqlOutputRows) SetValues(v []*string) *SqlOutputRows {
	s.Values = v
	return s
}

type SqlOutputSchema struct {
	Fields []*SqlOutputSchemaFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
}

func (s SqlOutputSchema) String() string {
	return tea.Prettify(s)
}

func (s SqlOutputSchema) GoString() string {
	return s.String()
}

func (s *SqlOutputSchema) SetFields(v []*SqlOutputSchemaFields) *SqlOutputSchema {
	s.Fields = v
	return s
}

type SqlOutputSchemaFields struct {
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Nullable *bool   `json:"nullable,omitempty" xml:"nullable,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SqlOutputSchemaFields) String() string {
	return tea.Prettify(s)
}

func (s SqlOutputSchemaFields) GoString() string {
	return s.String()
}

func (s *SqlOutputSchemaFields) SetName(v string) *SqlOutputSchemaFields {
	s.Name = &v
	return s
}

func (s *SqlOutputSchemaFields) SetNullable(v bool) *SqlOutputSchemaFields {
	s.Nullable = &v
	return s
}

func (s *SqlOutputSchemaFields) SetType(v string) *SqlOutputSchemaFields {
	s.Type = &v
	return s
}

type Tag struct {
	// 标签key值。
	//
	// example:
	//
	// workflowId
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 标签key值。
	//
	// example:
	//
	// wf-123test
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s Tag) String() string {
	return tea.Prettify(s)
}

func (s Tag) GoString() string {
	return s.String()
}

func (s *Tag) SetKey(v string) *Tag {
	s.Key = &v
	return s
}

func (s *Tag) SetValue(v string) *Tag {
	s.Value = &v
	return s
}

type Task struct {
	Archives    []*string `json:"archives,omitempty" xml:"archives,omitempty" type:"Repeated"`
	ArtifactUrl *string   `json:"artifactUrl,omitempty" xml:"artifactUrl,omitempty"`
	// This parameter is required.
	BizId         *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	CategoryBizId *string `json:"categoryBizId,omitempty" xml:"categoryBizId,omitempty"`
	Content       *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	Creator                *int64    `json:"creator,omitempty" xml:"creator,omitempty"`
	DefaultCatalogId       *string   `json:"defaultCatalogId,omitempty" xml:"defaultCatalogId,omitempty"`
	DefaultDatabase        *string   `json:"defaultDatabase,omitempty" xml:"defaultDatabase,omitempty"`
	DefaultResourceQueueId *string   `json:"defaultResourceQueueId,omitempty" xml:"defaultResourceQueueId,omitempty"`
	DefaultSqlComputeId    *string   `json:"defaultSqlComputeId,omitempty" xml:"defaultSqlComputeId,omitempty"`
	ExtraArtifactIds       []*string `json:"extraArtifactIds,omitempty" xml:"extraArtifactIds,omitempty" type:"Repeated"`
	ExtraSparkSubmitParams *string   `json:"extraSparkSubmitParams,omitempty" xml:"extraSparkSubmitParams,omitempty"`
	Files                  []*string `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// This parameter is required.
	GmtCreated *string `json:"gmtCreated,omitempty" xml:"gmtCreated,omitempty"`
	// This parameter is required.
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	HasChanged  *bool   `json:"hasChanged,omitempty" xml:"hasChanged,omitempty"`
	// This parameter is required.
	HasCommited            *bool     `json:"hasCommited,omitempty" xml:"hasCommited,omitempty"`
	Jars                   []*string `json:"jars,omitempty" xml:"jars,omitempty" type:"Repeated"`
	LastRunResourceQueueId *string   `json:"lastRunResourceQueueId,omitempty" xml:"lastRunResourceQueueId,omitempty"`
	// This parameter is required.
	Modifier *int64 `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// This parameter is required.
	Name    *string   `json:"name,omitempty" xml:"name,omitempty"`
	PyFiles []*string `json:"pyFiles,omitempty" xml:"pyFiles,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	SparkArgs *string      `json:"sparkArgs,omitempty" xml:"sparkArgs,omitempty"`
	SparkConf []*SparkConf `json:"sparkConf,omitempty" xml:"sparkConf,omitempty" type:"Repeated"`
	// This parameter is required.
	SparkDriverCores *int32 `json:"sparkDriverCores,omitempty" xml:"sparkDriverCores,omitempty"`
	// This parameter is required.
	SparkDriverMemory *int64  `json:"sparkDriverMemory,omitempty" xml:"sparkDriverMemory,omitempty"`
	SparkEntrypoint   *string `json:"sparkEntrypoint,omitempty" xml:"sparkEntrypoint,omitempty"`
	// This parameter is required.
	SparkExecutorCores *int32 `json:"sparkExecutorCores,omitempty" xml:"sparkExecutorCores,omitempty"`
	// This parameter is required.
	SparkExecutorMemory *int64 `json:"sparkExecutorMemory,omitempty" xml:"sparkExecutorMemory,omitempty"`
	// This parameter is required.
	SparkLogLevel *string `json:"sparkLogLevel,omitempty" xml:"sparkLogLevel,omitempty"`
	// This parameter is required.
	SparkLogPath *string `json:"sparkLogPath,omitempty" xml:"sparkLogPath,omitempty"`
	// This parameter is required.
	SparkVersion *string            `json:"sparkVersion,omitempty" xml:"sparkVersion,omitempty"`
	Tags         map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Task) String() string {
	return tea.Prettify(s)
}

func (s Task) GoString() string {
	return s.String()
}

func (s *Task) SetArchives(v []*string) *Task {
	s.Archives = v
	return s
}

func (s *Task) SetArtifactUrl(v string) *Task {
	s.ArtifactUrl = &v
	return s
}

func (s *Task) SetBizId(v string) *Task {
	s.BizId = &v
	return s
}

func (s *Task) SetCategoryBizId(v string) *Task {
	s.CategoryBizId = &v
	return s
}

func (s *Task) SetContent(v string) *Task {
	s.Content = &v
	return s
}

func (s *Task) SetCreator(v int64) *Task {
	s.Creator = &v
	return s
}

func (s *Task) SetDefaultCatalogId(v string) *Task {
	s.DefaultCatalogId = &v
	return s
}

func (s *Task) SetDefaultDatabase(v string) *Task {
	s.DefaultDatabase = &v
	return s
}

func (s *Task) SetDefaultResourceQueueId(v string) *Task {
	s.DefaultResourceQueueId = &v
	return s
}

func (s *Task) SetDefaultSqlComputeId(v string) *Task {
	s.DefaultSqlComputeId = &v
	return s
}

func (s *Task) SetExtraArtifactIds(v []*string) *Task {
	s.ExtraArtifactIds = v
	return s
}

func (s *Task) SetExtraSparkSubmitParams(v string) *Task {
	s.ExtraSparkSubmitParams = &v
	return s
}

func (s *Task) SetFiles(v []*string) *Task {
	s.Files = v
	return s
}

func (s *Task) SetGmtCreated(v string) *Task {
	s.GmtCreated = &v
	return s
}

func (s *Task) SetGmtModified(v string) *Task {
	s.GmtModified = &v
	return s
}

func (s *Task) SetHasChanged(v bool) *Task {
	s.HasChanged = &v
	return s
}

func (s *Task) SetHasCommited(v bool) *Task {
	s.HasCommited = &v
	return s
}

func (s *Task) SetJars(v []*string) *Task {
	s.Jars = v
	return s
}

func (s *Task) SetLastRunResourceQueueId(v string) *Task {
	s.LastRunResourceQueueId = &v
	return s
}

func (s *Task) SetModifier(v int64) *Task {
	s.Modifier = &v
	return s
}

func (s *Task) SetName(v string) *Task {
	s.Name = &v
	return s
}

func (s *Task) SetPyFiles(v []*string) *Task {
	s.PyFiles = v
	return s
}

func (s *Task) SetSparkArgs(v string) *Task {
	s.SparkArgs = &v
	return s
}

func (s *Task) SetSparkConf(v []*SparkConf) *Task {
	s.SparkConf = v
	return s
}

func (s *Task) SetSparkDriverCores(v int32) *Task {
	s.SparkDriverCores = &v
	return s
}

func (s *Task) SetSparkDriverMemory(v int64) *Task {
	s.SparkDriverMemory = &v
	return s
}

func (s *Task) SetSparkEntrypoint(v string) *Task {
	s.SparkEntrypoint = &v
	return s
}

func (s *Task) SetSparkExecutorCores(v int32) *Task {
	s.SparkExecutorCores = &v
	return s
}

func (s *Task) SetSparkExecutorMemory(v int64) *Task {
	s.SparkExecutorMemory = &v
	return s
}

func (s *Task) SetSparkLogLevel(v string) *Task {
	s.SparkLogLevel = &v
	return s
}

func (s *Task) SetSparkLogPath(v string) *Task {
	s.SparkLogPath = &v
	return s
}

func (s *Task) SetSparkVersion(v string) *Task {
	s.SparkVersion = &v
	return s
}

func (s *Task) SetTags(v map[string]*string) *Task {
	s.Tags = v
	return s
}

func (s *Task) SetType(v string) *Task {
	s.Type = &v
	return s
}

type TaskInstance struct {
	BizId          *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	Creator        *int64  `json:"creator,omitempty" xml:"creator,omitempty"`
	FenixRunId     *string `json:"fenixRunId,omitempty" xml:"fenixRunId,omitempty"`
	GmtCreated     *string `json:"gmtCreated,omitempty" xml:"gmtCreated,omitempty"`
	TaskBizId      *string `json:"taskBizId,omitempty" xml:"taskBizId,omitempty"`
	TaskInfo       *Task   `json:"taskInfo,omitempty" xml:"taskInfo,omitempty"`
	TaskStatus     *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
	WorkspaceBizId *string `json:"workspaceBizId,omitempty" xml:"workspaceBizId,omitempty"`
}

func (s TaskInstance) String() string {
	return tea.Prettify(s)
}

func (s TaskInstance) GoString() string {
	return s.String()
}

func (s *TaskInstance) SetBizId(v string) *TaskInstance {
	s.BizId = &v
	return s
}

func (s *TaskInstance) SetCreator(v int64) *TaskInstance {
	s.Creator = &v
	return s
}

func (s *TaskInstance) SetFenixRunId(v string) *TaskInstance {
	s.FenixRunId = &v
	return s
}

func (s *TaskInstance) SetGmtCreated(v string) *TaskInstance {
	s.GmtCreated = &v
	return s
}

func (s *TaskInstance) SetTaskBizId(v string) *TaskInstance {
	s.TaskBizId = &v
	return s
}

func (s *TaskInstance) SetTaskInfo(v *Task) *TaskInstance {
	s.TaskInfo = v
	return s
}

func (s *TaskInstance) SetTaskStatus(v string) *TaskInstance {
	s.TaskStatus = &v
	return s
}

func (s *TaskInstance) SetWorkspaceBizId(v string) *TaskInstance {
	s.WorkspaceBizId = &v
	return s
}

type TaskSnapshot struct {
	BizId      *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	Commiter   *int64  `json:"commiter,omitempty" xml:"commiter,omitempty"`
	GmtCreated *string `json:"gmtCreated,omitempty" xml:"gmtCreated,omitempty"`
	Item       *Task   `json:"item,omitempty" xml:"item,omitempty"`
	Message    *string `json:"message,omitempty" xml:"message,omitempty"`
	TaskBizId  *string `json:"taskBizId,omitempty" xml:"taskBizId,omitempty"`
	Version    *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s TaskSnapshot) String() string {
	return tea.Prettify(s)
}

func (s TaskSnapshot) GoString() string {
	return s.String()
}

func (s *TaskSnapshot) SetBizId(v string) *TaskSnapshot {
	s.BizId = &v
	return s
}

func (s *TaskSnapshot) SetCommiter(v int64) *TaskSnapshot {
	s.Commiter = &v
	return s
}

func (s *TaskSnapshot) SetGmtCreated(v string) *TaskSnapshot {
	s.GmtCreated = &v
	return s
}

func (s *TaskSnapshot) SetItem(v *Task) *TaskSnapshot {
	s.Item = v
	return s
}

func (s *TaskSnapshot) SetMessage(v string) *TaskSnapshot {
	s.Message = &v
	return s
}

func (s *TaskSnapshot) SetTaskBizId(v string) *TaskSnapshot {
	s.TaskBizId = &v
	return s
}

func (s *TaskSnapshot) SetVersion(v string) *TaskSnapshot {
	s.Version = &v
	return s
}

type Template struct {
	// This parameter is required.
	Creator *int64 `json:"creator,omitempty" xml:"creator,omitempty"`
	// This parameter is required.
	GmtCreated *string `json:"gmtCreated,omitempty" xml:"gmtCreated,omitempty"`
	// This parameter is required.
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// This parameter is required.
	Modifier  *int64       `json:"modifier,omitempty" xml:"modifier,omitempty"`
	SparkConf []*SparkConf `json:"sparkConf,omitempty" xml:"sparkConf,omitempty" type:"Repeated"`
	// This parameter is required.
	SparkDriverCores *int32 `json:"sparkDriverCores,omitempty" xml:"sparkDriverCores,omitempty"`
	// This parameter is required.
	SparkDriverMemory *int64 `json:"sparkDriverMemory,omitempty" xml:"sparkDriverMemory,omitempty"`
	// This parameter is required.
	SparkExecutorCores *int32 `json:"sparkExecutorCores,omitempty" xml:"sparkExecutorCores,omitempty"`
	// This parameter is required.
	SparkExecutorMemory *int64 `json:"sparkExecutorMemory,omitempty" xml:"sparkExecutorMemory,omitempty"`
	// This parameter is required.
	SparkLogLevel *string `json:"sparkLogLevel,omitempty" xml:"sparkLogLevel,omitempty"`
	// This parameter is required.
	SparkLogPath *string `json:"sparkLogPath,omitempty" xml:"sparkLogPath,omitempty"`
	// This parameter is required.
	SparkVersion *string `json:"sparkVersion,omitempty" xml:"sparkVersion,omitempty"`
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
}

func (s Template) String() string {
	return tea.Prettify(s)
}

func (s Template) GoString() string {
	return s.String()
}

func (s *Template) SetCreator(v int64) *Template {
	s.Creator = &v
	return s
}

func (s *Template) SetGmtCreated(v string) *Template {
	s.GmtCreated = &v
	return s
}

func (s *Template) SetGmtModified(v string) *Template {
	s.GmtModified = &v
	return s
}

func (s *Template) SetModifier(v int64) *Template {
	s.Modifier = &v
	return s
}

func (s *Template) SetSparkConf(v []*SparkConf) *Template {
	s.SparkConf = v
	return s
}

func (s *Template) SetSparkDriverCores(v int32) *Template {
	s.SparkDriverCores = &v
	return s
}

func (s *Template) SetSparkDriverMemory(v int64) *Template {
	s.SparkDriverMemory = &v
	return s
}

func (s *Template) SetSparkExecutorCores(v int32) *Template {
	s.SparkExecutorCores = &v
	return s
}

func (s *Template) SetSparkExecutorMemory(v int64) *Template {
	s.SparkExecutorMemory = &v
	return s
}

func (s *Template) SetSparkLogLevel(v string) *Template {
	s.SparkLogLevel = &v
	return s
}

func (s *Template) SetSparkLogPath(v string) *Template {
	s.SparkLogPath = &v
	return s
}

func (s *Template) SetSparkVersion(v string) *Template {
	s.SparkVersion = &v
	return s
}

func (s *Template) SetTemplateType(v string) *Template {
	s.TemplateType = &v
	return s
}

type TimeRange struct {
	// 时间范围结束时间。
	//
	// example:
	//
	// 1688370894339
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 时间范围开始时间。
	//
	// example:
	//
	// 1688370894339
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s TimeRange) String() string {
	return tea.Prettify(s)
}

func (s TimeRange) GoString() string {
	return s.String()
}

func (s *TimeRange) SetEndTime(v int64) *TimeRange {
	s.EndTime = &v
	return s
}

func (s *TimeRange) SetStartTime(v int64) *TimeRange {
	s.StartTime = &v
	return s
}

type AddMembersRequest struct {
	// This parameter is required.
	MemberArns []*string `json:"memberArns,omitempty" xml:"memberArns,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// w-975bcfda9625****
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s AddMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMembersRequest) GoString() string {
	return s.String()
}

func (s *AddMembersRequest) SetMemberArns(v []*string) *AddMembersRequest {
	s.MemberArns = v
	return s
}

func (s *AddMembersRequest) SetWorkspaceId(v string) *AddMembersRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AddMembersRequest) SetRegionId(v string) *AddMembersRequest {
	s.RegionId = &v
	return s
}

type AddMembersResponseBody struct {
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AddMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddMembersResponseBody) SetRequestId(v string) *AddMembersResponseBody {
	s.RequestId = &v
	return s
}

type AddMembersResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMembersResponse) GoString() string {
	return s.String()
}

func (s *AddMembersResponse) SetHeaders(v map[string]*string) *AddMembersResponse {
	s.Headers = v
	return s
}

func (s *AddMembersResponse) SetStatusCode(v int32) *AddMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMembersResponse) SetBody(v *AddMembersResponseBody) *AddMembersResponse {
	s.Body = v
	return s
}

type CancelJobRunRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CancelJobRunRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelJobRunRequest) GoString() string {
	return s.String()
}

func (s *CancelJobRunRequest) SetRegionId(v string) *CancelJobRunRequest {
	s.RegionId = &v
	return s
}

type CancelJobRunResponseBody struct {
	// example:
	//
	// jr-1a2bc3
	JobRunId *string `json:"jobRunId,omitempty" xml:"jobRunId,omitempty"`
	// 请求ID。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CancelJobRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelJobRunResponseBody) GoString() string {
	return s.String()
}

func (s *CancelJobRunResponseBody) SetJobRunId(v string) *CancelJobRunResponseBody {
	s.JobRunId = &v
	return s
}

func (s *CancelJobRunResponseBody) SetRequestId(v string) *CancelJobRunResponseBody {
	s.RequestId = &v
	return s
}

type CancelJobRunResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelJobRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelJobRunResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelJobRunResponse) GoString() string {
	return s.String()
}

func (s *CancelJobRunResponse) SetHeaders(v map[string]*string) *CancelJobRunResponse {
	s.Headers = v
	return s
}

func (s *CancelJobRunResponse) SetStatusCode(v int32) *CancelJobRunResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelJobRunResponse) SetBody(v *CancelJobRunResponseBody) *CancelJobRunResponse {
	s.Body = v
	return s
}

type GetJobRunRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetJobRunRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobRunRequest) GoString() string {
	return s.String()
}

func (s *GetJobRunRequest) SetRegionId(v string) *GetJobRunRequest {
	s.RegionId = &v
	return s
}

type GetJobRunResponseBody struct {
	JobRun *GetJobRunResponseBodyJobRun `json:"jobRun,omitempty" xml:"jobRun,omitempty" type:"Struct"`
	// 请求ID。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetJobRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobRunResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobRunResponseBody) SetJobRun(v *GetJobRunResponseBodyJobRun) *GetJobRunResponseBody {
	s.JobRun = v
	return s
}

func (s *GetJobRunResponseBody) SetRequestId(v string) *GetJobRunResponseBody {
	s.RequestId = &v
	return s
}

type GetJobRunResponseBodyJobRun struct {
	// 作业代码类型。
	//
	// example:
	//
	// SQL
	CodeType               *string                                            `json:"codeType,omitempty" xml:"codeType,omitempty"`
	ConfigurationOverrides *GetJobRunResponseBodyJobRunConfigurationOverrides `json:"configurationOverrides,omitempty" xml:"configurationOverrides,omitempty" type:"Struct"`
	// 作业结束时间。
	//
	// example:
	//
	// 1684119314000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 运行超时时间。
	//
	// example:
	//
	// 3600
	ExecutionTimeoutSeconds *int32     `json:"executionTimeoutSeconds,omitempty" xml:"executionTimeoutSeconds,omitempty"`
	JobDriver               *JobDriver `json:"jobDriver,omitempty" xml:"jobDriver,omitempty"`
	// 任务实例ID。
	//
	// example:
	//
	// jr-231231
	JobRunId *string `json:"jobRunId,omitempty" xml:"jobRunId,omitempty"`
	Log      *RunLog `json:"log,omitempty" xml:"log,omitempty"`
	// 作业实例名称。
	//
	// example:
	//
	// jobName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// esr-3.3.1
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// 创建用户Uid。
	//
	// example:
	//
	// 1509789347011222
	ResourceOwnerId *string `json:"resourceOwnerId,omitempty" xml:"resourceOwnerId,omitempty"`
	// example:
	//
	// root_queue
	ResourceQueueId *string `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	// 作业状态。
	//
	// example:
	//
	// Running
	State             *string                                       `json:"state,omitempty" xml:"state,omitempty"`
	StateChangeReason *GetJobRunResponseBodyJobRunStateChangeReason `json:"stateChangeReason,omitempty" xml:"stateChangeReason,omitempty" type:"Struct"`
	// 作业提交时间。
	//
	// example:
	//
	// 1684119314000
	SubmitTime *int64 `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
	// 标签。
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// 作业web ui。
	//
	// example:
	//
	// http://spark-ui
	WebUI *string `json:"webUI,omitempty" xml:"webUI,omitempty"`
	// 工作空间id。
	//
	// example:
	//
	// w-1234abcd
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetJobRunResponseBodyJobRun) String() string {
	return tea.Prettify(s)
}

func (s GetJobRunResponseBodyJobRun) GoString() string {
	return s.String()
}

func (s *GetJobRunResponseBodyJobRun) SetCodeType(v string) *GetJobRunResponseBodyJobRun {
	s.CodeType = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetConfigurationOverrides(v *GetJobRunResponseBodyJobRunConfigurationOverrides) *GetJobRunResponseBodyJobRun {
	s.ConfigurationOverrides = v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetEndTime(v int64) *GetJobRunResponseBodyJobRun {
	s.EndTime = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetExecutionTimeoutSeconds(v int32) *GetJobRunResponseBodyJobRun {
	s.ExecutionTimeoutSeconds = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetJobDriver(v *JobDriver) *GetJobRunResponseBodyJobRun {
	s.JobDriver = v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetJobRunId(v string) *GetJobRunResponseBodyJobRun {
	s.JobRunId = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetLog(v *RunLog) *GetJobRunResponseBodyJobRun {
	s.Log = v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetName(v string) *GetJobRunResponseBodyJobRun {
	s.Name = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetReleaseVersion(v string) *GetJobRunResponseBodyJobRun {
	s.ReleaseVersion = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetResourceOwnerId(v string) *GetJobRunResponseBodyJobRun {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetResourceQueueId(v string) *GetJobRunResponseBodyJobRun {
	s.ResourceQueueId = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetState(v string) *GetJobRunResponseBodyJobRun {
	s.State = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetStateChangeReason(v *GetJobRunResponseBodyJobRunStateChangeReason) *GetJobRunResponseBodyJobRun {
	s.StateChangeReason = v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetSubmitTime(v int64) *GetJobRunResponseBodyJobRun {
	s.SubmitTime = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetTags(v []*Tag) *GetJobRunResponseBodyJobRun {
	s.Tags = v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetWebUI(v string) *GetJobRunResponseBodyJobRun {
	s.WebUI = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetWorkspaceId(v string) *GetJobRunResponseBodyJobRun {
	s.WorkspaceId = &v
	return s
}

type GetJobRunResponseBodyJobRunConfigurationOverrides struct {
	Configurations []*Configuration `json:"configurations,omitempty" xml:"configurations,omitempty" type:"Repeated"`
}

func (s GetJobRunResponseBodyJobRunConfigurationOverrides) String() string {
	return tea.Prettify(s)
}

func (s GetJobRunResponseBodyJobRunConfigurationOverrides) GoString() string {
	return s.String()
}

func (s *GetJobRunResponseBodyJobRunConfigurationOverrides) SetConfigurations(v []*Configuration) *GetJobRunResponseBodyJobRunConfigurationOverrides {
	s.Configurations = v
	return s
}

type GetJobRunResponseBodyJobRunStateChangeReason struct {
	// example:
	//
	// ERR-100000
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// connection refused
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetJobRunResponseBodyJobRunStateChangeReason) String() string {
	return tea.Prettify(s)
}

func (s GetJobRunResponseBodyJobRunStateChangeReason) GoString() string {
	return s.String()
}

func (s *GetJobRunResponseBodyJobRunStateChangeReason) SetCode(v string) *GetJobRunResponseBodyJobRunStateChangeReason {
	s.Code = &v
	return s
}

func (s *GetJobRunResponseBodyJobRunStateChangeReason) SetMessage(v string) *GetJobRunResponseBodyJobRunStateChangeReason {
	s.Message = &v
	return s
}

type GetJobRunResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobRunResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobRunResponse) GoString() string {
	return s.String()
}

func (s *GetJobRunResponse) SetHeaders(v map[string]*string) *GetJobRunResponse {
	s.Headers = v
	return s
}

func (s *GetJobRunResponse) SetStatusCode(v int32) *GetJobRunResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobRunResponse) SetBody(v *GetJobRunResponseBody) *GetJobRunResponse {
	s.Body = v
	return s
}

type GrantRoleToUsersRequest struct {
	// example:
	//
	// acs:emr::w-975bcfda9625****:role/Owner
	RoleArn  *string   `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	UserArns []*string `json:"userArns,omitempty" xml:"userArns,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GrantRoleToUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantRoleToUsersRequest) GoString() string {
	return s.String()
}

func (s *GrantRoleToUsersRequest) SetRoleArn(v string) *GrantRoleToUsersRequest {
	s.RoleArn = &v
	return s
}

func (s *GrantRoleToUsersRequest) SetUserArns(v []*string) *GrantRoleToUsersRequest {
	s.UserArns = v
	return s
}

func (s *GrantRoleToUsersRequest) SetRegionId(v string) *GrantRoleToUsersRequest {
	s.RegionId = &v
	return s
}

type GrantRoleToUsersResponseBody struct {
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GrantRoleToUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantRoleToUsersResponseBody) GoString() string {
	return s.String()
}

func (s *GrantRoleToUsersResponseBody) SetRequestId(v string) *GrantRoleToUsersResponseBody {
	s.RequestId = &v
	return s
}

type GrantRoleToUsersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantRoleToUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantRoleToUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantRoleToUsersResponse) GoString() string {
	return s.String()
}

func (s *GrantRoleToUsersResponse) SetHeaders(v map[string]*string) *GrantRoleToUsersResponse {
	s.Headers = v
	return s
}

func (s *GrantRoleToUsersResponse) SetStatusCode(v int32) *GrantRoleToUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantRoleToUsersResponse) SetBody(v *GrantRoleToUsersResponseBody) *GrantRoleToUsersResponse {
	s.Body = v
	return s
}

type ListJobRunsRequest struct {
	// 创建用户Uid。
	//
	// example:
	//
	// 1509789347011222
	Creator *string                    `json:"creator,omitempty" xml:"creator,omitempty"`
	EndTime *ListJobRunsRequestEndTime `json:"endTime,omitempty" xml:"endTime,omitempty" type:"Struct"`
	// 作业id。
	//
	// example:
	//
	// j-xxx
	JobRunId *string `json:"jobRunId,omitempty" xml:"jobRunId,omitempty"`
	// 一次获取的最大记录数。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 作业名称。
	//
	// example:
	//
	// emr-spark-demo-job
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 标记当前开始读取的位置，置空表示从头开始。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// dev_queue
	ResourceQueueId *string                      `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	StartTime       *ListJobRunsRequestStartTime `json:"startTime,omitempty" xml:"startTime,omitempty" type:"Struct"`
	// 作业状态。
	//
	// example:
	//
	// ["Running","Submitted"]
	States []*string `json:"states,omitempty" xml:"states,omitempty" type:"Repeated"`
	// 标签。
	Tags []*ListJobRunsRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ListJobRunsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobRunsRequest) GoString() string {
	return s.String()
}

func (s *ListJobRunsRequest) SetCreator(v string) *ListJobRunsRequest {
	s.Creator = &v
	return s
}

func (s *ListJobRunsRequest) SetEndTime(v *ListJobRunsRequestEndTime) *ListJobRunsRequest {
	s.EndTime = v
	return s
}

func (s *ListJobRunsRequest) SetJobRunId(v string) *ListJobRunsRequest {
	s.JobRunId = &v
	return s
}

func (s *ListJobRunsRequest) SetMaxResults(v int32) *ListJobRunsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListJobRunsRequest) SetName(v string) *ListJobRunsRequest {
	s.Name = &v
	return s
}

func (s *ListJobRunsRequest) SetNextToken(v string) *ListJobRunsRequest {
	s.NextToken = &v
	return s
}

func (s *ListJobRunsRequest) SetRegionId(v string) *ListJobRunsRequest {
	s.RegionId = &v
	return s
}

func (s *ListJobRunsRequest) SetResourceQueueId(v string) *ListJobRunsRequest {
	s.ResourceQueueId = &v
	return s
}

func (s *ListJobRunsRequest) SetStartTime(v *ListJobRunsRequestStartTime) *ListJobRunsRequest {
	s.StartTime = v
	return s
}

func (s *ListJobRunsRequest) SetStates(v []*string) *ListJobRunsRequest {
	s.States = v
	return s
}

func (s *ListJobRunsRequest) SetTags(v []*ListJobRunsRequestTags) *ListJobRunsRequest {
	s.Tags = v
	return s
}

type ListJobRunsRequestEndTime struct {
	// example:
	//
	// 1710432000000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1709740800000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListJobRunsRequestEndTime) String() string {
	return tea.Prettify(s)
}

func (s ListJobRunsRequestEndTime) GoString() string {
	return s.String()
}

func (s *ListJobRunsRequestEndTime) SetEndTime(v int64) *ListJobRunsRequestEndTime {
	s.EndTime = &v
	return s
}

func (s *ListJobRunsRequestEndTime) SetStartTime(v int64) *ListJobRunsRequestEndTime {
	s.StartTime = &v
	return s
}

type ListJobRunsRequestStartTime struct {
	// example:
	//
	// 1710432000000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1709740800000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListJobRunsRequestStartTime) String() string {
	return tea.Prettify(s)
}

func (s ListJobRunsRequestStartTime) GoString() string {
	return s.String()
}

func (s *ListJobRunsRequestStartTime) SetEndTime(v int64) *ListJobRunsRequestStartTime {
	s.EndTime = &v
	return s
}

func (s *ListJobRunsRequestStartTime) SetStartTime(v int64) *ListJobRunsRequestStartTime {
	s.StartTime = &v
	return s
}

type ListJobRunsRequestTags struct {
	// example:
	//
	// tag_key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListJobRunsRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListJobRunsRequestTags) GoString() string {
	return s.String()
}

func (s *ListJobRunsRequestTags) SetKey(v string) *ListJobRunsRequestTags {
	s.Key = &v
	return s
}

func (s *ListJobRunsRequestTags) SetValue(v string) *ListJobRunsRequestTags {
	s.Value = &v
	return s
}

type ListJobRunsShrinkRequest struct {
	// 创建用户Uid。
	//
	// example:
	//
	// 1509789347011222
	Creator       *string `json:"creator,omitempty" xml:"creator,omitempty"`
	EndTimeShrink *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 作业id。
	//
	// example:
	//
	// j-xxx
	JobRunId *string `json:"jobRunId,omitempty" xml:"jobRunId,omitempty"`
	// 一次获取的最大记录数。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 作业名称。
	//
	// example:
	//
	// emr-spark-demo-job
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 标记当前开始读取的位置，置空表示从头开始。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// dev_queue
	ResourceQueueId *string `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	StartTimeShrink *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 作业状态。
	//
	// example:
	//
	// ["Running","Submitted"]
	StatesShrink *string `json:"states,omitempty" xml:"states,omitempty"`
	// 标签。
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListJobRunsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobRunsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListJobRunsShrinkRequest) SetCreator(v string) *ListJobRunsShrinkRequest {
	s.Creator = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetEndTimeShrink(v string) *ListJobRunsShrinkRequest {
	s.EndTimeShrink = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetJobRunId(v string) *ListJobRunsShrinkRequest {
	s.JobRunId = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetMaxResults(v int32) *ListJobRunsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetName(v string) *ListJobRunsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetNextToken(v string) *ListJobRunsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetRegionId(v string) *ListJobRunsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetResourceQueueId(v string) *ListJobRunsShrinkRequest {
	s.ResourceQueueId = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetStartTimeShrink(v string) *ListJobRunsShrinkRequest {
	s.StartTimeShrink = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetStatesShrink(v string) *ListJobRunsShrinkRequest {
	s.StatesShrink = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetTagsShrink(v string) *ListJobRunsShrinkRequest {
	s.TagsShrink = &v
	return s
}

type ListJobRunsResponseBody struct {
	JobRuns []*ListJobRunsResponseBodyJobRuns `json:"jobRuns,omitempty" xml:"jobRuns,omitempty" type:"Repeated"`
	// 本次请求所返回的最大记录条数。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 返回读取到的数据位置，空代表数据已经读取完毕。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 请求ID。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 本次请求条件下的数据总量。
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListJobRunsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobRunsResponseBody) SetJobRuns(v []*ListJobRunsResponseBodyJobRuns) *ListJobRunsResponseBody {
	s.JobRuns = v
	return s
}

func (s *ListJobRunsResponseBody) SetMaxResults(v int32) *ListJobRunsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListJobRunsResponseBody) SetNextToken(v string) *ListJobRunsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListJobRunsResponseBody) SetRequestId(v string) *ListJobRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobRunsResponseBody) SetTotalCount(v int32) *ListJobRunsResponseBody {
	s.TotalCount = &v
	return s
}

type ListJobRunsResponseBodyJobRuns struct {
	// 作业代码类型。
	//
	// example:
	//
	// SQL
	CodeType               *string                                               `json:"codeType,omitempty" xml:"codeType,omitempty"`
	ConfigurationOverrides *ListJobRunsResponseBodyJobRunsConfigurationOverrides `json:"configurationOverrides,omitempty" xml:"configurationOverrides,omitempty" type:"Struct"`
	// 创建用户Uid。
	//
	// example:
	//
	// 1509789347011222
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 作业结束时间。
	//
	// example:
	//
	// 1684119314000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 运行超时时间。
	//
	// example:
	//
	// 3600
	ExecutionTimeoutSeconds *int32     `json:"executionTimeoutSeconds,omitempty" xml:"executionTimeoutSeconds,omitempty"`
	JobDriver               *JobDriver `json:"jobDriver,omitempty" xml:"jobDriver,omitempty"`
	// 任务实例ID。
	//
	// example:
	//
	// jr-231231
	JobRunId *string `json:"jobRunId,omitempty" xml:"jobRunId,omitempty"`
	Log      *RunLog `json:"log,omitempty" xml:"log,omitempty"`
	// 作业实例名称。
	//
	// example:
	//
	// jobName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// esr-native-3.4.0
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// 作业状态。
	//
	// example:
	//
	// Running
	State             *string                                          `json:"state,omitempty" xml:"state,omitempty"`
	StateChangeReason *ListJobRunsResponseBodyJobRunsStateChangeReason `json:"stateChangeReason,omitempty" xml:"stateChangeReason,omitempty" type:"Struct"`
	// 作业提交时间。
	//
	// example:
	//
	// 1684119314000
	SubmitTime *int64 `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
	// 标签。
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// 作业web ui。
	//
	// example:
	//
	// http://spark-ui
	WebUI *string `json:"webUI,omitempty" xml:"webUI,omitempty"`
	// 工作空间id。
	//
	// example:
	//
	// w-1234abcd
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListJobRunsResponseBodyJobRuns) String() string {
	return tea.Prettify(s)
}

func (s ListJobRunsResponseBodyJobRuns) GoString() string {
	return s.String()
}

func (s *ListJobRunsResponseBodyJobRuns) SetCodeType(v string) *ListJobRunsResponseBodyJobRuns {
	s.CodeType = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetConfigurationOverrides(v *ListJobRunsResponseBodyJobRunsConfigurationOverrides) *ListJobRunsResponseBodyJobRuns {
	s.ConfigurationOverrides = v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetCreator(v string) *ListJobRunsResponseBodyJobRuns {
	s.Creator = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetEndTime(v int64) *ListJobRunsResponseBodyJobRuns {
	s.EndTime = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetExecutionTimeoutSeconds(v int32) *ListJobRunsResponseBodyJobRuns {
	s.ExecutionTimeoutSeconds = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetJobDriver(v *JobDriver) *ListJobRunsResponseBodyJobRuns {
	s.JobDriver = v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetJobRunId(v string) *ListJobRunsResponseBodyJobRuns {
	s.JobRunId = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetLog(v *RunLog) *ListJobRunsResponseBodyJobRuns {
	s.Log = v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetName(v string) *ListJobRunsResponseBodyJobRuns {
	s.Name = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetReleaseVersion(v string) *ListJobRunsResponseBodyJobRuns {
	s.ReleaseVersion = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetState(v string) *ListJobRunsResponseBodyJobRuns {
	s.State = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetStateChangeReason(v *ListJobRunsResponseBodyJobRunsStateChangeReason) *ListJobRunsResponseBodyJobRuns {
	s.StateChangeReason = v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetSubmitTime(v int64) *ListJobRunsResponseBodyJobRuns {
	s.SubmitTime = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetTags(v []*Tag) *ListJobRunsResponseBodyJobRuns {
	s.Tags = v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetWebUI(v string) *ListJobRunsResponseBodyJobRuns {
	s.WebUI = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetWorkspaceId(v string) *ListJobRunsResponseBodyJobRuns {
	s.WorkspaceId = &v
	return s
}

type ListJobRunsResponseBodyJobRunsConfigurationOverrides struct {
	Configurations []*Configuration `json:"configurations,omitempty" xml:"configurations,omitempty" type:"Repeated"`
}

func (s ListJobRunsResponseBodyJobRunsConfigurationOverrides) String() string {
	return tea.Prettify(s)
}

func (s ListJobRunsResponseBodyJobRunsConfigurationOverrides) GoString() string {
	return s.String()
}

func (s *ListJobRunsResponseBodyJobRunsConfigurationOverrides) SetConfigurations(v []*Configuration) *ListJobRunsResponseBodyJobRunsConfigurationOverrides {
	s.Configurations = v
	return s
}

type ListJobRunsResponseBodyJobRunsStateChangeReason struct {
	// example:
	//
	// 0
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListJobRunsResponseBodyJobRunsStateChangeReason) String() string {
	return tea.Prettify(s)
}

func (s ListJobRunsResponseBodyJobRunsStateChangeReason) GoString() string {
	return s.String()
}

func (s *ListJobRunsResponseBodyJobRunsStateChangeReason) SetCode(v string) *ListJobRunsResponseBodyJobRunsStateChangeReason {
	s.Code = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRunsStateChangeReason) SetMessage(v string) *ListJobRunsResponseBodyJobRunsStateChangeReason {
	s.Message = &v
	return s
}

type ListJobRunsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobRunsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobRunsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobRunsResponse) GoString() string {
	return s.String()
}

func (s *ListJobRunsResponse) SetHeaders(v map[string]*string) *ListJobRunsResponse {
	s.Headers = v
	return s
}

func (s *ListJobRunsResponse) SetStatusCode(v int32) *ListJobRunsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobRunsResponse) SetBody(v *ListJobRunsResponseBody) *ListJobRunsResponse {
	s.Body = v
	return s
}

type ListReleaseVersionsRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// stable
	ReleaseType *string `json:"releaseType,omitempty" xml:"releaseType,omitempty"`
	// example:
	//
	// esr-2.1 (Spark 3.3.1, Scala 2.12, Java Runtime)
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// example:
	//
	// ONLINE
	ReleaseVersionStatus *string `json:"releaseVersionStatus,omitempty" xml:"releaseVersionStatus,omitempty"`
}

func (s ListReleaseVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListReleaseVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListReleaseVersionsRequest) SetRegionId(v string) *ListReleaseVersionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListReleaseVersionsRequest) SetReleaseType(v string) *ListReleaseVersionsRequest {
	s.ReleaseType = &v
	return s
}

func (s *ListReleaseVersionsRequest) SetReleaseVersion(v string) *ListReleaseVersionsRequest {
	s.ReleaseVersion = &v
	return s
}

func (s *ListReleaseVersionsRequest) SetReleaseVersionStatus(v string) *ListReleaseVersionsRequest {
	s.ReleaseVersionStatus = &v
	return s
}

type ListReleaseVersionsResponseBody struct {
	// 一次获取的最大记录数。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 下一页TOKEN。
	//
	// example:
	//
	// 1
	NextToken       *string                                           `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ReleaseVersions []*ListReleaseVersionsResponseBodyReleaseVersions `json:"releaseVersions,omitempty" xml:"releaseVersions,omitempty" type:"Repeated"`
	// 请求ID。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 记录总数。
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListReleaseVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListReleaseVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListReleaseVersionsResponseBody) SetMaxResults(v int32) *ListReleaseVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListReleaseVersionsResponseBody) SetNextToken(v string) *ListReleaseVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListReleaseVersionsResponseBody) SetReleaseVersions(v []*ListReleaseVersionsResponseBodyReleaseVersions) *ListReleaseVersionsResponseBody {
	s.ReleaseVersions = v
	return s
}

func (s *ListReleaseVersionsResponseBody) SetRequestId(v string) *ListReleaseVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListReleaseVersionsResponseBody) SetTotalCount(v int32) *ListReleaseVersionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListReleaseVersionsResponseBodyReleaseVersions struct {
	// example:
	//
	// Spark 3.3.1
	CommunityVersion *string   `json:"communityVersion,omitempty" xml:"communityVersion,omitempty"`
	CpuArchitectures []*string `json:"cpuArchitectures,omitempty" xml:"cpuArchitectures,omitempty" type:"Repeated"`
	// example:
	//
	// 1716215854101
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// ASI
	IaasType *string `json:"iaasType,omitempty" xml:"iaasType,omitempty"`
	// example:
	//
	// esr-2.1 (Spark 3.3.1, Scala 2.12, Java Runtime)
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// example:
	//
	// 2.12
	ScalaVersion *string `json:"scalaVersion,omitempty" xml:"scalaVersion,omitempty"`
	// example:
	//
	// ONLINE
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// stable
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListReleaseVersionsResponseBodyReleaseVersions) String() string {
	return tea.Prettify(s)
}

func (s ListReleaseVersionsResponseBodyReleaseVersions) GoString() string {
	return s.String()
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetCommunityVersion(v string) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.CommunityVersion = &v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetCpuArchitectures(v []*string) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.CpuArchitectures = v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetGmtCreate(v int64) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.GmtCreate = &v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetIaasType(v string) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.IaasType = &v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetReleaseVersion(v string) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.ReleaseVersion = &v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetScalaVersion(v string) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.ScalaVersion = &v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetState(v string) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.State = &v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetType(v string) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.Type = &v
	return s
}

type ListReleaseVersionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListReleaseVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListReleaseVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListReleaseVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListReleaseVersionsResponse) SetHeaders(v map[string]*string) *ListReleaseVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListReleaseVersionsResponse) SetStatusCode(v int32) *ListReleaseVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListReleaseVersionsResponse) SetBody(v *ListReleaseVersionsResponseBody) *ListReleaseVersionsResponse {
	s.Body = v
	return s
}

type ListSessionClustersRequest struct {
	// 一次获取的最大记录数。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 标记当前开始读取的位置，置空表示从头开始。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// root
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// 作业名称。
	//
	// example:
	//
	// emr-spark-demo-job
	SessionClusterId *string `json:"sessionClusterId,omitempty" xml:"sessionClusterId,omitempty"`
}

func (s ListSessionClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSessionClustersRequest) GoString() string {
	return s.String()
}

func (s *ListSessionClustersRequest) SetMaxResults(v int32) *ListSessionClustersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSessionClustersRequest) SetNextToken(v string) *ListSessionClustersRequest {
	s.NextToken = &v
	return s
}

func (s *ListSessionClustersRequest) SetQueueName(v string) *ListSessionClustersRequest {
	s.QueueName = &v
	return s
}

func (s *ListSessionClustersRequest) SetRegionId(v string) *ListSessionClustersRequest {
	s.RegionId = &v
	return s
}

func (s *ListSessionClustersRequest) SetSessionClusterId(v string) *ListSessionClustersRequest {
	s.SessionClusterId = &v
	return s
}

type ListSessionClustersResponseBody struct {
	// 本次请求所返回的最大记录条数。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 返回读取到的数据位置，空代表数据已经读取完毕。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 请求ID。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId       *string                                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SessionClusters []*ListSessionClustersResponseBodySessionClusters `json:"sessionClusters,omitempty" xml:"sessionClusters,omitempty" type:"Repeated"`
	// 本次请求条件下的数据总量。
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListSessionClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSessionClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListSessionClustersResponseBody) SetMaxResults(v int32) *ListSessionClustersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSessionClustersResponseBody) SetNextToken(v string) *ListSessionClustersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSessionClustersResponseBody) SetRequestId(v string) *ListSessionClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSessionClustersResponseBody) SetSessionClusters(v []*ListSessionClustersResponseBodySessionClusters) *ListSessionClustersResponseBody {
	s.SessionClusters = v
	return s
}

func (s *ListSessionClustersResponseBody) SetTotalCount(v int32) *ListSessionClustersResponseBody {
	s.TotalCount = &v
	return s
}

type ListSessionClustersResponseBodySessionClusters struct {
	ApplicationConfigs     []*ListSessionClustersResponseBodySessionClustersApplicationConfigs   `json:"applicationConfigs,omitempty" xml:"applicationConfigs,omitempty" type:"Repeated"`
	AutoStartConfiguration *ListSessionClustersResponseBodySessionClustersAutoStartConfiguration `json:"autoStartConfiguration,omitempty" xml:"autoStartConfiguration,omitempty" type:"Struct"`
	AutoStopConfiguration  *ListSessionClustersResponseBodySessionClustersAutoStopConfiguration  `json:"autoStopConfiguration,omitempty" xml:"autoStopConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// adhoc_query
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 作业实例名称。
	//
	// example:
	//
	// dev_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// SQL Compute id
	//
	// example:
	//
	// sc-123131
	SessionClusterId *string `json:"sessionClusterId,omitempty" xml:"sessionClusterId,omitempty"`
	// 作业状态。
	//
	// example:
	//
	// Running
	State             *string                                                          `json:"state,omitempty" xml:"state,omitempty"`
	StateChangeReason *ListSessionClustersResponseBodySessionClustersStateChangeReason `json:"stateChangeReason,omitempty" xml:"stateChangeReason,omitempty" type:"Struct"`
	// 任务实例ID。
	//
	// example:
	//
	// 123131
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// test_user
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// 工作空间id。
	//
	// example:
	//
	// w-1234abcd
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListSessionClustersResponseBodySessionClusters) String() string {
	return tea.Prettify(s)
}

func (s ListSessionClustersResponseBodySessionClusters) GoString() string {
	return s.String()
}

func (s *ListSessionClustersResponseBodySessionClusters) SetApplicationConfigs(v []*ListSessionClustersResponseBodySessionClustersApplicationConfigs) *ListSessionClustersResponseBodySessionClusters {
	s.ApplicationConfigs = v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetAutoStartConfiguration(v *ListSessionClustersResponseBodySessionClustersAutoStartConfiguration) *ListSessionClustersResponseBodySessionClusters {
	s.AutoStartConfiguration = v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetAutoStopConfiguration(v *ListSessionClustersResponseBodySessionClustersAutoStopConfiguration) *ListSessionClustersResponseBodySessionClusters {
	s.AutoStopConfiguration = v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetName(v string) *ListSessionClustersResponseBodySessionClusters {
	s.Name = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetQueueName(v string) *ListSessionClustersResponseBodySessionClusters {
	s.QueueName = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetSessionClusterId(v string) *ListSessionClustersResponseBodySessionClusters {
	s.SessionClusterId = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetState(v string) *ListSessionClustersResponseBodySessionClusters {
	s.State = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetStateChangeReason(v *ListSessionClustersResponseBodySessionClustersStateChangeReason) *ListSessionClustersResponseBodySessionClusters {
	s.StateChangeReason = v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetUserId(v string) *ListSessionClustersResponseBodySessionClusters {
	s.UserId = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetUserName(v string) *ListSessionClustersResponseBodySessionClusters {
	s.UserName = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetWorkspaceId(v string) *ListSessionClustersResponseBodySessionClusters {
	s.WorkspaceId = &v
	return s
}

type ListSessionClustersResponseBodySessionClustersApplicationConfigs struct {
	// example:
	//
	// spark-default.conf
	ConfigFileName *string `json:"configFileName,omitempty" xml:"configFileName,omitempty"`
	// example:
	//
	// spark.app.name
	ConfigItemKey *string `json:"configItemKey,omitempty" xml:"configItemKey,omitempty"`
	// example:
	//
	// test_application
	ConfigItemValue *string `json:"configItemValue,omitempty" xml:"configItemValue,omitempty"`
}

func (s ListSessionClustersResponseBodySessionClustersApplicationConfigs) String() string {
	return tea.Prettify(s)
}

func (s ListSessionClustersResponseBodySessionClustersApplicationConfigs) GoString() string {
	return s.String()
}

func (s *ListSessionClustersResponseBodySessionClustersApplicationConfigs) SetConfigFileName(v string) *ListSessionClustersResponseBodySessionClustersApplicationConfigs {
	s.ConfigFileName = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClustersApplicationConfigs) SetConfigItemKey(v string) *ListSessionClustersResponseBodySessionClustersApplicationConfigs {
	s.ConfigItemKey = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClustersApplicationConfigs) SetConfigItemValue(v string) *ListSessionClustersResponseBodySessionClustersApplicationConfigs {
	s.ConfigItemValue = &v
	return s
}

type ListSessionClustersResponseBodySessionClustersAutoStartConfiguration struct {
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s ListSessionClustersResponseBodySessionClustersAutoStartConfiguration) String() string {
	return tea.Prettify(s)
}

func (s ListSessionClustersResponseBodySessionClustersAutoStartConfiguration) GoString() string {
	return s.String()
}

func (s *ListSessionClustersResponseBodySessionClustersAutoStartConfiguration) SetEnable(v bool) *ListSessionClustersResponseBodySessionClustersAutoStartConfiguration {
	s.Enable = &v
	return s
}

type ListSessionClustersResponseBodySessionClustersAutoStopConfiguration struct {
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 45
	IdleTimeoutMinutes *int32 `json:"idleTimeoutMinutes,omitempty" xml:"idleTimeoutMinutes,omitempty"`
}

func (s ListSessionClustersResponseBodySessionClustersAutoStopConfiguration) String() string {
	return tea.Prettify(s)
}

func (s ListSessionClustersResponseBodySessionClustersAutoStopConfiguration) GoString() string {
	return s.String()
}

func (s *ListSessionClustersResponseBodySessionClustersAutoStopConfiguration) SetEnable(v bool) *ListSessionClustersResponseBodySessionClustersAutoStopConfiguration {
	s.Enable = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClustersAutoStopConfiguration) SetIdleTimeoutMinutes(v int32) *ListSessionClustersResponseBodySessionClustersAutoStopConfiguration {
	s.IdleTimeoutMinutes = &v
	return s
}

type ListSessionClustersResponseBodySessionClustersStateChangeReason struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListSessionClustersResponseBodySessionClustersStateChangeReason) String() string {
	return tea.Prettify(s)
}

func (s ListSessionClustersResponseBodySessionClustersStateChangeReason) GoString() string {
	return s.String()
}

func (s *ListSessionClustersResponseBodySessionClustersStateChangeReason) SetCode(v string) *ListSessionClustersResponseBodySessionClustersStateChangeReason {
	s.Code = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClustersStateChangeReason) SetMessage(v string) *ListSessionClustersResponseBodySessionClustersStateChangeReason {
	s.Message = &v
	return s
}

type ListSessionClustersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSessionClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSessionClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSessionClustersResponse) GoString() string {
	return s.String()
}

func (s *ListSessionClustersResponse) SetHeaders(v map[string]*string) *ListSessionClustersResponse {
	s.Headers = v
	return s
}

func (s *ListSessionClustersResponse) SetStatusCode(v int32) *ListSessionClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSessionClustersResponse) SetBody(v *ListSessionClustersResponseBody) *ListSessionClustersResponse {
	s.Body = v
	return s
}

type ListWorkspaceQueuesRequest struct {
	// example:
	//
	// production
	Environment *string `json:"environment,omitempty" xml:"environment,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListWorkspaceQueuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspaceQueuesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspaceQueuesRequest) SetEnvironment(v string) *ListWorkspaceQueuesRequest {
	s.Environment = &v
	return s
}

func (s *ListWorkspaceQueuesRequest) SetRegionId(v string) *ListWorkspaceQueuesRequest {
	s.RegionId = &v
	return s
}

type ListWorkspaceQueuesResponseBody struct {
	// 一次获取的最大记录数。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 下一页TOKEN。
	//
	// example:
	//
	// 1
	NextToken *string                                  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Queues    []*ListWorkspaceQueuesResponseBodyQueues `json:"queues,omitempty" xml:"queues,omitempty" type:"Repeated"`
	// 请求ID。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 记录总数。
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListWorkspaceQueuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspaceQueuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspaceQueuesResponseBody) SetMaxResults(v int32) *ListWorkspaceQueuesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBody) SetNextToken(v string) *ListWorkspaceQueuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBody) SetQueues(v []*ListWorkspaceQueuesResponseBodyQueues) *ListWorkspaceQueuesResponseBody {
	s.Queues = v
	return s
}

func (s *ListWorkspaceQueuesResponseBody) SetRequestId(v string) *ListWorkspaceQueuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBody) SetTotalCount(v int32) *ListWorkspaceQueuesResponseBody {
	s.TotalCount = &v
	return s
}

type ListWorkspaceQueuesResponseBodyQueues struct {
	// 队列允许的操作
	AllowActions []*ListWorkspaceQueuesResponseBodyQueuesAllowActions `json:"allowActions,omitempty" xml:"allowActions,omitempty" type:"Repeated"`
	// example:
	//
	// 237109
	Creator      *string   `json:"creator,omitempty" xml:"creator,omitempty"`
	Environments []*string `json:"environments,omitempty" xml:"environments,omitempty" type:"Repeated"`
	// 队列资源最大容量
	//
	// example:
	//
	// {"cpu": "2","memory": "2Gi"}
	MaxResource *string `json:"maxResource,omitempty" xml:"maxResource,omitempty"`
	// 队列资源最小容量
	//
	// example:
	//
	// {"cpu": "2","memory": "2Gi"}
	MinResource *string `json:"minResource,omitempty" xml:"minResource,omitempty"`
	// 队列Label
	//
	// example:
	//
	// dev_queue
	Properties *string `json:"properties,omitempty" xml:"properties,omitempty"`
	// 队列名称。
	//
	// example:
	//
	// dev_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// 队列架构
	//
	// example:
	//
	// {"arch": "x86"}
	QueueScope *string `json:"queueScope,omitempty" xml:"queueScope,omitempty"`
	// example:
	//
	// RUNNING
	QueueStatus *string `json:"queueStatus,omitempty" xml:"queueStatus,omitempty"`
	// 队列类型
	//
	// example:
	//
	// instance, instanceChildren
	QueueType *string `json:"queueType,omitempty" xml:"queueType,omitempty"`
	// regionId。
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// 队列资源使用容量
	//
	// example:
	//
	// {"cpu": "2","memory": "2Gi"}
	UsedResource *string `json:"usedResource,omitempty" xml:"usedResource,omitempty"`
	// 工作空间id。
	//
	// example:
	//
	// w-1234abcd
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListWorkspaceQueuesResponseBodyQueues) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspaceQueuesResponseBodyQueues) GoString() string {
	return s.String()
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetAllowActions(v []*ListWorkspaceQueuesResponseBodyQueuesAllowActions) *ListWorkspaceQueuesResponseBodyQueues {
	s.AllowActions = v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetCreator(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.Creator = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetEnvironments(v []*string) *ListWorkspaceQueuesResponseBodyQueues {
	s.Environments = v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetMaxResource(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.MaxResource = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetMinResource(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.MinResource = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetProperties(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.Properties = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetQueueName(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.QueueName = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetQueueScope(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.QueueScope = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetQueueStatus(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.QueueStatus = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetQueueType(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.QueueType = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetRegionId(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.RegionId = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetUsedResource(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.UsedResource = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetWorkspaceId(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.WorkspaceId = &v
	return s
}

type ListWorkspaceQueuesResponseBodyQueuesAllowActions struct {
	// 行为 arn。
	//
	// example:
	//
	// acs:emr::workspaceId:action/create_queue
	ActionArn *string `json:"actionArn,omitempty" xml:"actionArn,omitempty"`
	// 权限名称。
	//
	// example:
	//
	// view
	ActionName *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	// action 依赖列表。
	//
	// example:
	//
	// ["view"]
	Dependencies []*string `json:"dependencies,omitempty" xml:"dependencies,omitempty" type:"Repeated"`
	// action 描述。
	//
	// example:
	//
	// 文件目录遍历、文件浏览
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 权限展示名称。
	//
	// example:
	//
	// 文件目录遍历、文件浏览
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s ListWorkspaceQueuesResponseBodyQueuesAllowActions) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspaceQueuesResponseBodyQueuesAllowActions) GoString() string {
	return s.String()
}

func (s *ListWorkspaceQueuesResponseBodyQueuesAllowActions) SetActionArn(v string) *ListWorkspaceQueuesResponseBodyQueuesAllowActions {
	s.ActionArn = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueuesAllowActions) SetActionName(v string) *ListWorkspaceQueuesResponseBodyQueuesAllowActions {
	s.ActionName = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueuesAllowActions) SetDependencies(v []*string) *ListWorkspaceQueuesResponseBodyQueuesAllowActions {
	s.Dependencies = v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueuesAllowActions) SetDescription(v string) *ListWorkspaceQueuesResponseBodyQueuesAllowActions {
	s.Description = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueuesAllowActions) SetDisplayName(v string) *ListWorkspaceQueuesResponseBodyQueuesAllowActions {
	s.DisplayName = &v
	return s
}

type ListWorkspaceQueuesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkspaceQueuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkspaceQueuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspaceQueuesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkspaceQueuesResponse) SetHeaders(v map[string]*string) *ListWorkspaceQueuesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkspaceQueuesResponse) SetStatusCode(v int32) *ListWorkspaceQueuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkspaceQueuesResponse) SetBody(v *ListWorkspaceQueuesResponseBody) *ListWorkspaceQueuesResponse {
	s.Body = v
	return s
}

type ListWorkspacesRequest struct {
	// 一次获取的最大记录数。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// test_workspace
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 标记当前开始读取的位置，置空表示从头开始。
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s ListWorkspacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequest) SetMaxResults(v int32) *ListWorkspacesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesRequest) SetName(v string) *ListWorkspacesRequest {
	s.Name = &v
	return s
}

func (s *ListWorkspacesRequest) SetNextToken(v string) *ListWorkspacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesRequest) SetRegionId(v string) *ListWorkspacesRequest {
	s.RegionId = &v
	return s
}

func (s *ListWorkspacesRequest) SetState(v string) *ListWorkspacesRequest {
	s.State = &v
	return s
}

type ListWorkspacesResponseBody struct {
	// 一次获取的最大记录数。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 下一页TOKEN。
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 请求ID。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 记录总数。
	//
	// example:
	//
	// 200
	TotalCount *int32                                  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	Workspaces []*ListWorkspacesResponseBodyWorkspaces `json:"workspaces,omitempty" xml:"workspaces,omitempty" type:"Repeated"`
}

func (s ListWorkspacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBody) SetMaxResults(v int32) *ListWorkspacesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetNextToken(v string) *ListWorkspacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetRequestId(v string) *ListWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetTotalCount(v int32) *ListWorkspacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetWorkspaces(v []*ListWorkspacesResponseBodyWorkspaces) *ListWorkspacesResponseBody {
	s.Workspaces = v
	return s
}

type ListWorkspacesResponseBodyWorkspaces struct {
	// 是否自动续费(pre付费类型必须)。
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// 自动续费时长(pre付费类型必须)。
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"autoRenewPeriod,omitempty" xml:"autoRenewPeriod,omitempty"`
	// 自动续费周期(pre付费类型必须)。
	//
	// example:
	//
	// YEAR, MONTH, WEEK, DAY, HOUR, MINUTE
	AutoRenewPeriodUnit *string `json:"autoRenewPeriodUnit,omitempty" xml:"autoRenewPeriodUnit,omitempty"`
	// example:
	//
	// 1684115879955
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// dlf catalog 信息。
	//
	// example:
	//
	// default
	DlfCatalogId *string `json:"dlfCatalogId,omitempty" xml:"dlfCatalogId,omitempty"`
	// 订购周期数量(pre付费类型必须)。
	//
	// example:
	//
	// 1
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 1687103999999
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 失败原因。
	//
	// example:
	//
	// out of stock
	FailReason *string `json:"failReason,omitempty" xml:"failReason,omitempty"`
	// 订购周期(pre付费类型必须)。
	//
	// example:
	//
	// YEAR, MONTH, WEEK, DAY, HOUR, MINUTE
	PaymentDurationUnit *string `json:"paymentDurationUnit,omitempty" xml:"paymentDurationUnit,omitempty"`
	// 支付状态。
	//
	// example:
	//
	// PAID/UNPAID
	PaymentStatus *string `json:"paymentStatus,omitempty" xml:"paymentStatus,omitempty"`
	// 付费类型。
	//
	// example:
	//
	// PayAsYouGo or Subscription
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// SERVICE_RELEASE
	ReleaseType *string `json:"releaseType,omitempty" xml:"releaseType,omitempty"`
	// 资源规格。
	//
	// example:
	//
	// 100cu
	ResourceSpec      *string                                                `json:"resourceSpec,omitempty" xml:"resourceSpec,omitempty"`
	StateChangeReason *ListWorkspacesResponseBodyWorkspacesStateChangeReason `json:"stateChangeReason,omitempty" xml:"stateChangeReason,omitempty" type:"Struct"`
	// oss 路径。
	//
	// example:
	//
	// spark-result
	Storage *string `json:"storage,omitempty" xml:"storage,omitempty"`
	// Workspace Id。
	//
	// example:
	//
	// w-******
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// 工作空间名称。
	//
	// example:
	//
	// spark批作业空间-1
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
	// 工作空间状态。
	//
	// example:
	//
	// STARTING,RUNNING,TERMINATED
	WorkspaceStatus *string `json:"workspaceStatus,omitempty" xml:"workspaceStatus,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspaces) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspaces) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetAutoRenew(v bool) *ListWorkspacesResponseBodyWorkspaces {
	s.AutoRenew = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetAutoRenewPeriod(v int32) *ListWorkspacesResponseBodyWorkspaces {
	s.AutoRenewPeriod = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetAutoRenewPeriodUnit(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.AutoRenewPeriodUnit = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetCreateTime(v int64) *ListWorkspacesResponseBodyWorkspaces {
	s.CreateTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetDlfCatalogId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.DlfCatalogId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetDuration(v int32) *ListWorkspacesResponseBodyWorkspaces {
	s.Duration = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetEndTime(v int64) *ListWorkspacesResponseBodyWorkspaces {
	s.EndTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetFailReason(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.FailReason = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetPaymentDurationUnit(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.PaymentDurationUnit = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetPaymentStatus(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.PaymentStatus = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetPaymentType(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.PaymentType = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetRegionId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.RegionId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetReleaseType(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.ReleaseType = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetResourceSpec(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.ResourceSpec = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetStateChangeReason(v *ListWorkspacesResponseBodyWorkspacesStateChangeReason) *ListWorkspacesResponseBodyWorkspaces {
	s.StateChangeReason = v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetStorage(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Storage = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetWorkspaceId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetWorkspaceName(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.WorkspaceName = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetWorkspaceStatus(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.WorkspaceStatus = &v
	return s
}

type ListWorkspacesResponseBodyWorkspacesStateChangeReason struct {
	// example:
	//
	// 0
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspacesStateChangeReason) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspacesStateChangeReason) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspacesStateChangeReason) SetCode(v string) *ListWorkspacesResponseBodyWorkspacesStateChangeReason {
	s.Code = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesStateChangeReason) SetMessage(v string) *ListWorkspacesResponseBodyWorkspacesStateChangeReason {
	s.Message = &v
	return s
}

type ListWorkspacesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkspacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponse) SetHeaders(v map[string]*string) *ListWorkspacesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkspacesResponse) SetStatusCode(v int32) *ListWorkspacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkspacesResponse) SetBody(v *ListWorkspacesResponseBody) *ListWorkspacesResponse {
	s.Body = v
	return s
}

type StartJobRunRequest struct {
	// example:
	//
	// 8e6aae2810c8f67229ca70bb31cd6028
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// SQL
	CodeType               *string                                   `json:"codeType,omitempty" xml:"codeType,omitempty"`
	ConfigurationOverrides *StartJobRunRequestConfigurationOverrides `json:"configurationOverrides,omitempty" xml:"configurationOverrides,omitempty" type:"Struct"`
	// example:
	//
	// 100
	ExecutionTimeoutSeconds *int32     `json:"executionTimeoutSeconds,omitempty" xml:"executionTimeoutSeconds,omitempty"`
	JobDriver               *JobDriver `json:"jobDriver,omitempty" xml:"jobDriver,omitempty"`
	// example:
	//
	// jr-12345
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// spark_job_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// esr-3.3.1
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// example:
	//
	// dev_queue
	ResourceQueueId *string `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	Tags            []*Tag  `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s StartJobRunRequest) String() string {
	return tea.Prettify(s)
}

func (s StartJobRunRequest) GoString() string {
	return s.String()
}

func (s *StartJobRunRequest) SetClientToken(v string) *StartJobRunRequest {
	s.ClientToken = &v
	return s
}

func (s *StartJobRunRequest) SetCodeType(v string) *StartJobRunRequest {
	s.CodeType = &v
	return s
}

func (s *StartJobRunRequest) SetConfigurationOverrides(v *StartJobRunRequestConfigurationOverrides) *StartJobRunRequest {
	s.ConfigurationOverrides = v
	return s
}

func (s *StartJobRunRequest) SetExecutionTimeoutSeconds(v int32) *StartJobRunRequest {
	s.ExecutionTimeoutSeconds = &v
	return s
}

func (s *StartJobRunRequest) SetJobDriver(v *JobDriver) *StartJobRunRequest {
	s.JobDriver = v
	return s
}

func (s *StartJobRunRequest) SetJobId(v string) *StartJobRunRequest {
	s.JobId = &v
	return s
}

func (s *StartJobRunRequest) SetName(v string) *StartJobRunRequest {
	s.Name = &v
	return s
}

func (s *StartJobRunRequest) SetReleaseVersion(v string) *StartJobRunRequest {
	s.ReleaseVersion = &v
	return s
}

func (s *StartJobRunRequest) SetResourceQueueId(v string) *StartJobRunRequest {
	s.ResourceQueueId = &v
	return s
}

func (s *StartJobRunRequest) SetTags(v []*Tag) *StartJobRunRequest {
	s.Tags = v
	return s
}

func (s *StartJobRunRequest) SetRegionId(v string) *StartJobRunRequest {
	s.RegionId = &v
	return s
}

type StartJobRunRequestConfigurationOverrides struct {
	Configurations []*StartJobRunRequestConfigurationOverridesConfigurations `json:"configurations,omitempty" xml:"configurations,omitempty" type:"Repeated"`
}

func (s StartJobRunRequestConfigurationOverrides) String() string {
	return tea.Prettify(s)
}

func (s StartJobRunRequestConfigurationOverrides) GoString() string {
	return s.String()
}

func (s *StartJobRunRequestConfigurationOverrides) SetConfigurations(v []*StartJobRunRequestConfigurationOverridesConfigurations) *StartJobRunRequestConfigurationOverrides {
	s.Configurations = v
	return s
}

type StartJobRunRequestConfigurationOverridesConfigurations struct {
	// example:
	//
	// spark-default.conf
	ConfigFileName *string `json:"configFileName,omitempty" xml:"configFileName,omitempty"`
	// example:
	//
	// spark.app.name
	ConfigItemKey *string `json:"configItemKey,omitempty" xml:"configItemKey,omitempty"`
	// example:
	//
	// test_app
	ConfigItemValue *string `json:"configItemValue,omitempty" xml:"configItemValue,omitempty"`
}

func (s StartJobRunRequestConfigurationOverridesConfigurations) String() string {
	return tea.Prettify(s)
}

func (s StartJobRunRequestConfigurationOverridesConfigurations) GoString() string {
	return s.String()
}

func (s *StartJobRunRequestConfigurationOverridesConfigurations) SetConfigFileName(v string) *StartJobRunRequestConfigurationOverridesConfigurations {
	s.ConfigFileName = &v
	return s
}

func (s *StartJobRunRequestConfigurationOverridesConfigurations) SetConfigItemKey(v string) *StartJobRunRequestConfigurationOverridesConfigurations {
	s.ConfigItemKey = &v
	return s
}

func (s *StartJobRunRequestConfigurationOverridesConfigurations) SetConfigItemValue(v string) *StartJobRunRequestConfigurationOverridesConfigurations {
	s.ConfigItemValue = &v
	return s
}

type StartJobRunResponseBody struct {
	// example:
	//
	// jr-54321
	JobRunId *string `json:"jobRunId,omitempty" xml:"jobRunId,omitempty"`
	// 请求ID。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StartJobRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartJobRunResponseBody) GoString() string {
	return s.String()
}

func (s *StartJobRunResponseBody) SetJobRunId(v string) *StartJobRunResponseBody {
	s.JobRunId = &v
	return s
}

func (s *StartJobRunResponseBody) SetRequestId(v string) *StartJobRunResponseBody {
	s.RequestId = &v
	return s
}

type StartJobRunResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartJobRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartJobRunResponse) String() string {
	return tea.Prettify(s)
}

func (s StartJobRunResponse) GoString() string {
	return s.String()
}

func (s *StartJobRunResponse) SetHeaders(v map[string]*string) *StartJobRunResponse {
	s.Headers = v
	return s
}

func (s *StartJobRunResponse) SetStatusCode(v int32) *StartJobRunResponse {
	s.StatusCode = &v
	return s
}

func (s *StartJobRunResponse) SetBody(v *StartJobRunResponseBody) *StartJobRunResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("emr-serverless-spark"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 添加用户
//
// @param request - AddMembersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMembersResponse
func (client *Client) AddMembersWithOptions(request *AddMembersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MemberArns)) {
		body["memberArns"] = request.MemberArns
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddMembers"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/auth/members"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加用户
//
// @param request - AddMembersRequest
//
// @return AddMembersResponse
func (client *Client) AddMembers(request *AddMembersRequest) (_result *AddMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddMembersResponse{}
	_body, _err := client.AddMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消jobRun作业
//
// @param request - CancelJobRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelJobRunResponse
func (client *Client) CancelJobRunWithOptions(workspaceId *string, jobRunId *string, request *CancelJobRunRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelJobRunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelJobRun"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/jobRuns/" + tea.StringValue(openapiutil.GetEncodeParam(jobRunId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelJobRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消jobRun作业
//
// @param request - CancelJobRunRequest
//
// @return CancelJobRunResponse
func (client *Client) CancelJobRun(workspaceId *string, jobRunId *string, request *CancelJobRunRequest) (_result *CancelJobRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelJobRunResponse{}
	_body, _err := client.CancelJobRunWithOptions(workspaceId, jobRunId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务
//
// @param request - GetJobRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobRunResponse
func (client *Client) GetJobRunWithOptions(workspaceId *string, jobRunId *string, request *GetJobRunRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetJobRunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobRun"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/jobRuns/" + tea.StringValue(openapiutil.GetEncodeParam(jobRunId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务
//
// @param request - GetJobRunRequest
//
// @return GetJobRunResponse
func (client *Client) GetJobRun(workspaceId *string, jobRunId *string, request *GetJobRunRequest) (_result *GetJobRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobRunResponse{}
	_body, _err := client.GetJobRunWithOptions(workspaceId, jobRunId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 给用户授权Role列表
//
// @param request - GrantRoleToUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantRoleToUsersResponse
func (client *Client) GrantRoleToUsersWithOptions(request *GrantRoleToUsersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GrantRoleToUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleArn)) {
		body["roleArn"] = request.RoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.UserArns)) {
		body["userArns"] = request.UserArns
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantRoleToUsers"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/auth/roles/grant"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantRoleToUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 给用户授权Role列表
//
// @param request - GrantRoleToUsersRequest
//
// @return GrantRoleToUsersResponse
func (client *Client) GrantRoleToUsers(request *GrantRoleToUsersRequest) (_result *GrantRoleToUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GrantRoleToUsersResponse{}
	_body, _err := client.GrantRoleToUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询run列表
//
// @param tmpReq - ListJobRunsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobRunsResponse
func (client *Client) ListJobRunsWithOptions(workspaceId *string, tmpReq *ListJobRunsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListJobRunsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListJobRunsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EndTime)) {
		request.EndTimeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EndTime, tea.String("endTime"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.StartTime)) {
		request.StartTimeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StartTime, tea.String("startTime"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.States)) {
		request.StatesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.States, tea.String("states"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Creator)) {
		query["creator"] = request.Creator
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeShrink)) {
		query["endTime"] = request.EndTimeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.JobRunId)) {
		query["jobRunId"] = request.JobRunId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceQueueId)) {
		query["resourceQueueId"] = request.ResourceQueueId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeShrink)) {
		query["startTime"] = request.StartTimeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StatesShrink)) {
		query["states"] = request.StatesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobRuns"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/jobRuns"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJobRunsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询run列表
//
// @param request - ListJobRunsRequest
//
// @return ListJobRunsResponse
func (client *Client) ListJobRuns(workspaceId *string, request *ListJobRunsRequest) (_result *ListJobRunsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListJobRunsResponse{}
	_body, _err := client.ListJobRunsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取发布版本列表
//
// @param request - ListReleaseVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListReleaseVersionsResponse
func (client *Client) ListReleaseVersionsWithOptions(request *ListReleaseVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListReleaseVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseType)) {
		query["releaseType"] = request.ReleaseType
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseVersion)) {
		query["releaseVersion"] = request.ReleaseVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseVersionStatus)) {
		query["releaseVersionStatus"] = request.ReleaseVersionStatus
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListReleaseVersions"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/releaseVersions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListReleaseVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取发布版本列表
//
// @param request - ListReleaseVersionsRequest
//
// @return ListReleaseVersionsResponse
func (client *Client) ListReleaseVersions(request *ListReleaseVersionsRequest) (_result *ListReleaseVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListReleaseVersionsResponse{}
	_body, _err := client.ListReleaseVersionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询run列表
//
// @param request - ListSessionClustersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSessionClustersResponse
func (client *Client) ListSessionClustersWithOptions(workspaceId *string, request *ListSessionClustersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSessionClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		query["queueName"] = request.QueueName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionClusterId)) {
		query["sessionClusterId"] = request.SessionClusterId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSessionClusters"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/sessionClusters"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSessionClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询run列表
//
// @param request - ListSessionClustersRequest
//
// @return ListSessionClustersResponse
func (client *Client) ListSessionClusters(workspaceId *string, request *ListSessionClustersRequest) (_result *ListSessionClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSessionClustersResponse{}
	_body, _err := client.ListSessionClustersWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看工作空间队列列表
//
// @param request - ListWorkspaceQueuesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspaceQueuesResponse
func (client *Client) ListWorkspaceQueuesWithOptions(workspaceId *string, request *ListWorkspaceQueuesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListWorkspaceQueuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		query["environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkspaceQueues"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/queues"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkspaceQueuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看工作空间队列列表
//
// @param request - ListWorkspaceQueuesRequest
//
// @return ListWorkspaceQueuesResponse
func (client *Client) ListWorkspaceQueues(workspaceId *string, request *ListWorkspaceQueuesRequest) (_result *ListWorkspaceQueuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWorkspaceQueuesResponse{}
	_body, _err := client.ListWorkspaceQueuesWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看工作空间列表
//
// @param request - ListWorkspacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspacesWithOptions(request *ListWorkspacesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["state"] = request.State
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkspaces"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看工作空间列表
//
// @param request - ListWorkspacesRequest
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspaces(request *ListWorkspacesRequest) (_result *ListWorkspacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWorkspacesResponse{}
	_body, _err := client.ListWorkspacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启动作业
//
// @param request - StartJobRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartJobRunResponse
func (client *Client) StartJobRunWithOptions(workspaceId *string, request *StartJobRunRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartJobRunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CodeType)) {
		body["codeType"] = request.CodeType
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigurationOverrides)) {
		body["configurationOverrides"] = request.ConfigurationOverrides
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionTimeoutSeconds)) {
		body["executionTimeoutSeconds"] = request.ExecutionTimeoutSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.JobDriver)) {
		body["jobDriver"] = request.JobDriver
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["jobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseVersion)) {
		body["releaseVersion"] = request.ReleaseVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceQueueId)) {
		body["resourceQueueId"] = request.ResourceQueueId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartJobRun"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/jobRuns"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartJobRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动作业
//
// @param request - StartJobRunRequest
//
// @return StartJobRunResponse
func (client *Client) StartJobRun(workspaceId *string, request *StartJobRunRequest) (_result *StartJobRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartJobRunResponse{}
	_body, _err := client.StartJobRunWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
