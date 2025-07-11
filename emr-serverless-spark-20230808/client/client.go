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
	BizId         *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	CatagoryBizId *string `json:"catagoryBizId,omitempty" xml:"catagoryBizId,omitempty"`
	// This parameter is required.
	Creator    *int64      `json:"creator,omitempty" xml:"creator,omitempty"`
	Credential *Credential `json:"credential,omitempty" xml:"credential,omitempty"`
	FullPath   []*string   `json:"fullPath,omitempty" xml:"fullPath,omitempty" type:"Repeated"`
	// This parameter is required.
	GmtCreated *string `json:"gmtCreated,omitempty" xml:"gmtCreated,omitempty"`
	// This parameter is required.
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// This parameter is required.
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// This parameter is required.
	Modifier     *int64  `json:"modifier,omitempty" xml:"modifier,omitempty"`
	ModifierName *string `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
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

func (s *Artifact) SetCatagoryBizId(v string) *Artifact {
	s.CatagoryBizId = &v
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

func (s *Artifact) SetFullPath(v []*string) *Artifact {
	s.FullPath = v
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

func (s *Artifact) SetModifierName(v string) *Artifact {
	s.ModifierName = &v
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

type KerberosConf struct {
	Creator          *string   `json:"creator,omitempty" xml:"creator,omitempty"`
	Enabled          *bool     `json:"enabled,omitempty" xml:"enabled,omitempty"`
	GmtCreated       *string   `json:"gmtCreated,omitempty" xml:"gmtCreated,omitempty"`
	GmtModified      *string   `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	KerberosConfId   *string   `json:"kerberosConfId,omitempty" xml:"kerberosConfId,omitempty"`
	Keytabs          []*string `json:"keytabs,omitempty" xml:"keytabs,omitempty" type:"Repeated"`
	Krb5Conf         *string   `json:"krb5Conf,omitempty" xml:"krb5Conf,omitempty"`
	Name             *string   `json:"name,omitempty" xml:"name,omitempty"`
	NetworkServiceId *string   `json:"networkServiceId,omitempty" xml:"networkServiceId,omitempty"`
	WorkspaceId      *string   `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s KerberosConf) String() string {
	return tea.Prettify(s)
}

func (s KerberosConf) GoString() string {
	return s.String()
}

func (s *KerberosConf) SetCreator(v string) *KerberosConf {
	s.Creator = &v
	return s
}

func (s *KerberosConf) SetEnabled(v bool) *KerberosConf {
	s.Enabled = &v
	return s
}

func (s *KerberosConf) SetGmtCreated(v string) *KerberosConf {
	s.GmtCreated = &v
	return s
}

func (s *KerberosConf) SetGmtModified(v string) *KerberosConf {
	s.GmtModified = &v
	return s
}

func (s *KerberosConf) SetKerberosConfId(v string) *KerberosConf {
	s.KerberosConfId = &v
	return s
}

func (s *KerberosConf) SetKeytabs(v []*string) *KerberosConf {
	s.Keytabs = v
	return s
}

func (s *KerberosConf) SetKrb5Conf(v string) *KerberosConf {
	s.Krb5Conf = &v
	return s
}

func (s *KerberosConf) SetName(v string) *KerberosConf {
	s.Name = &v
	return s
}

func (s *KerberosConf) SetNetworkServiceId(v string) *KerberosConf {
	s.NetworkServiceId = &v
	return s
}

func (s *KerberosConf) SetWorkspaceId(v string) *KerberosConf {
	s.WorkspaceId = &v
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
	Creator                *int64          `json:"creator,omitempty" xml:"creator,omitempty"`
	Credential             *TaskCredential `json:"credential,omitempty" xml:"credential,omitempty" type:"Struct"`
	DefaultCatalogId       *string         `json:"defaultCatalogId,omitempty" xml:"defaultCatalogId,omitempty"`
	DefaultDatabase        *string         `json:"defaultDatabase,omitempty" xml:"defaultDatabase,omitempty"`
	DefaultResourceQueueId *string         `json:"defaultResourceQueueId,omitempty" xml:"defaultResourceQueueId,omitempty"`
	DefaultSqlComputeId    *string         `json:"defaultSqlComputeId,omitempty" xml:"defaultSqlComputeId,omitempty"`
	DeploymentId           *string         `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	EnvironmentId          *string         `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	ExtraArtifactIds       []*string       `json:"extraArtifactIds,omitempty" xml:"extraArtifactIds,omitempty" type:"Repeated"`
	ExtraSparkSubmitParams *string         `json:"extraSparkSubmitParams,omitempty" xml:"extraSparkSubmitParams,omitempty"`
	Files                  []*string       `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	Fusion                 *bool           `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// This parameter is required.
	GmtCreated *string `json:"gmtCreated,omitempty" xml:"gmtCreated,omitempty"`
	// This parameter is required.
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	HasChanged  *bool   `json:"hasChanged,omitempty" xml:"hasChanged,omitempty"`
	// This parameter is required.
	HasCommited            *bool     `json:"hasCommited,omitempty" xml:"hasCommited,omitempty"`
	IsStreaming            *bool     `json:"isStreaming,omitempty" xml:"isStreaming,omitempty"`
	Jars                   []*string `json:"jars,omitempty" xml:"jars,omitempty" type:"Repeated"`
	KernelId               *string   `json:"kernelId,omitempty" xml:"kernelId,omitempty"`
	LastRunResourceQueueId *string   `json:"lastRunResourceQueueId,omitempty" xml:"lastRunResourceQueueId,omitempty"`
	// This parameter is required.
	Modifier *int64 `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// This parameter is required.
	Name             *string            `json:"name,omitempty" xml:"name,omitempty"`
	Params           map[string]*string `json:"params,omitempty" xml:"params,omitempty"`
	PyFiles          []*string          `json:"pyFiles,omitempty" xml:"pyFiles,omitempty" type:"Repeated"`
	SessionClusterId *string            `json:"sessionClusterId,omitempty" xml:"sessionClusterId,omitempty"`
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
	SparkLogPath      *string `json:"sparkLogPath,omitempty" xml:"sparkLogPath,omitempty"`
	SparkSubmitClause *string `json:"sparkSubmitClause,omitempty" xml:"sparkSubmitClause,omitempty"`
	// This parameter is required.
	SparkVersion *string            `json:"sparkVersion,omitempty" xml:"sparkVersion,omitempty"`
	Tags         map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
	Timeout      *int32             `json:"timeout,omitempty" xml:"timeout,omitempty"`
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

func (s *Task) SetCredential(v *TaskCredential) *Task {
	s.Credential = v
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

func (s *Task) SetDeploymentId(v string) *Task {
	s.DeploymentId = &v
	return s
}

func (s *Task) SetEnvironmentId(v string) *Task {
	s.EnvironmentId = &v
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

func (s *Task) SetFusion(v bool) *Task {
	s.Fusion = &v
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

func (s *Task) SetIsStreaming(v bool) *Task {
	s.IsStreaming = &v
	return s
}

func (s *Task) SetJars(v []*string) *Task {
	s.Jars = v
	return s
}

func (s *Task) SetKernelId(v string) *Task {
	s.KernelId = &v
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

func (s *Task) SetParams(v map[string]*string) *Task {
	s.Params = v
	return s
}

func (s *Task) SetPyFiles(v []*string) *Task {
	s.PyFiles = v
	return s
}

func (s *Task) SetSessionClusterId(v string) *Task {
	s.SessionClusterId = &v
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

func (s *Task) SetSparkSubmitClause(v string) *Task {
	s.SparkSubmitClause = &v
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

func (s *Task) SetTimeout(v int32) *Task {
	s.Timeout = &v
	return s
}

func (s *Task) SetType(v string) *Task {
	s.Type = &v
	return s
}

type TaskCredential struct {
	AccessId      *string `json:"accessId,omitempty" xml:"accessId,omitempty"`
	AccessUrl     *string `json:"accessUrl,omitempty" xml:"accessUrl,omitempty"`
	Expire        *int64  `json:"expire,omitempty" xml:"expire,omitempty"`
	Host          *string `json:"host,omitempty" xml:"host,omitempty"`
	Path          *string `json:"path,omitempty" xml:"path,omitempty"`
	Policy        *string `json:"policy,omitempty" xml:"policy,omitempty"`
	SecurityToken *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
	Signature     *string `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s TaskCredential) String() string {
	return tea.Prettify(s)
}

func (s TaskCredential) GoString() string {
	return s.String()
}

func (s *TaskCredential) SetAccessId(v string) *TaskCredential {
	s.AccessId = &v
	return s
}

func (s *TaskCredential) SetAccessUrl(v string) *TaskCredential {
	s.AccessUrl = &v
	return s
}

func (s *TaskCredential) SetExpire(v int64) *TaskCredential {
	s.Expire = &v
	return s
}

func (s *TaskCredential) SetHost(v string) *TaskCredential {
	s.Host = &v
	return s
}

func (s *TaskCredential) SetPath(v string) *TaskCredential {
	s.Path = &v
	return s
}

func (s *TaskCredential) SetPolicy(v string) *TaskCredential {
	s.Policy = &v
	return s
}

func (s *TaskCredential) SetSecurityToken(v string) *TaskCredential {
	s.SecurityToken = &v
	return s
}

func (s *TaskCredential) SetSignature(v string) *TaskCredential {
	s.Signature = &v
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
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	Creator             *int64  `json:"creator,omitempty" xml:"creator,omitempty"`
	DisplaySparkVersion *string `json:"displaySparkVersion,omitempty" xml:"displaySparkVersion,omitempty"`
	Fusion              *bool   `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// This parameter is required.
	GmtCreated *string `json:"gmtCreated,omitempty" xml:"gmtCreated,omitempty"`
	// This parameter is required.
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	IsDefault   *bool   `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	// This parameter is required.
	Modifier  *int64       `json:"modifier,omitempty" xml:"modifier,omitempty"`
	Name      *string      `json:"name,omitempty" xml:"name,omitempty"`
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

func (s *Template) SetBizId(v string) *Template {
	s.BizId = &v
	return s
}

func (s *Template) SetCreator(v int64) *Template {
	s.Creator = &v
	return s
}

func (s *Template) SetDisplaySparkVersion(v string) *Template {
	s.DisplaySparkVersion = &v
	return s
}

func (s *Template) SetFusion(v bool) *Template {
	s.Fusion = &v
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

func (s *Template) SetIsDefault(v bool) *Template {
	s.IsDefault = &v
	return s
}

func (s *Template) SetModifier(v int64) *Template {
	s.Modifier = &v
	return s
}

func (s *Template) SetName(v string) *Template {
	s.Name = &v
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// w-975bcfda9625****
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// The region ID.
	//
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
	// The request ID.
	//
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
	// The region ID.
	//
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
	// The job ID.
	//
	// example:
	//
	// jr-1a2bc3
	JobRunId *string `json:"jobRunId,omitempty" xml:"jobRunId,omitempty"`
	// The request ID.
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

type CreateLivyComputeRequest struct {
	// example:
	//
	// Token
	AuthType               *string                                         `json:"authType,omitempty" xml:"authType,omitempty"`
	AutoStartConfiguration *CreateLivyComputeRequestAutoStartConfiguration `json:"autoStartConfiguration,omitempty" xml:"autoStartConfiguration,omitempty" type:"Struct"`
	AutoStopConfiguration  *CreateLivyComputeRequestAutoStopConfiguration  `json:"autoStopConfiguration,omitempty" xml:"autoStopConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// 1
	CpuLimit *string `json:"cpuLimit,omitempty" xml:"cpuLimit,omitempty"`
	// example:
	//
	// esr-4.3.0 (Spark 3.5.2, Scala 2.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// example:
	//
	// true
	EnablePublic *bool `json:"enablePublic,omitempty" xml:"enablePublic,omitempty"`
	// example:
	//
	// ev-ctfq0fem1hkhgv4hapng
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// example:
	//
	// false
	Fusion *bool `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// example:
	//
	// {
	//
	//   "sparkDefaultsConf": "spark.driver.cores     1\\nspark.driver.memory    4g\\nspark.executor.cores   1\\nspark.executor.memory  4g\\n",
	//
	//   "sparkBlackListConf": "spark.driver.cores\\nspark.driver.memory",
	//
	//   "livyConf": "livy.server.session.timeout  1h\\n",
	//
	//   "livyClientConf": "livy.rsc.sql.num-rows  1000\\n"
	//
	// }
	LivyServerConf *string `json:"livyServerConf,omitempty" xml:"livyServerConf,omitempty"`
	// example:
	//
	// 0.8.0
	LivyVersion *string `json:"livyVersion,omitempty" xml:"livyVersion,omitempty"`
	// example:
	//
	// 4Gi
	MemoryLimit *string `json:"memoryLimit,omitempty" xml:"memoryLimit,omitempty"`
	// example:
	//
	// testGateway
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// test
	NetworkName *string `json:"networkName,omitempty" xml:"networkName,omitempty"`
	// example:
	//
	// root_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// example:
	//
	// esr-4.3.0 (Spark 3.5.2, Scala 2.12, Java Runtime)
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateLivyComputeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLivyComputeRequest) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeRequest) SetAuthType(v string) *CreateLivyComputeRequest {
	s.AuthType = &v
	return s
}

func (s *CreateLivyComputeRequest) SetAutoStartConfiguration(v *CreateLivyComputeRequestAutoStartConfiguration) *CreateLivyComputeRequest {
	s.AutoStartConfiguration = v
	return s
}

func (s *CreateLivyComputeRequest) SetAutoStopConfiguration(v *CreateLivyComputeRequestAutoStopConfiguration) *CreateLivyComputeRequest {
	s.AutoStopConfiguration = v
	return s
}

func (s *CreateLivyComputeRequest) SetCpuLimit(v string) *CreateLivyComputeRequest {
	s.CpuLimit = &v
	return s
}

func (s *CreateLivyComputeRequest) SetDisplayReleaseVersion(v string) *CreateLivyComputeRequest {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *CreateLivyComputeRequest) SetEnablePublic(v bool) *CreateLivyComputeRequest {
	s.EnablePublic = &v
	return s
}

func (s *CreateLivyComputeRequest) SetEnvironmentId(v string) *CreateLivyComputeRequest {
	s.EnvironmentId = &v
	return s
}

func (s *CreateLivyComputeRequest) SetFusion(v bool) *CreateLivyComputeRequest {
	s.Fusion = &v
	return s
}

func (s *CreateLivyComputeRequest) SetLivyServerConf(v string) *CreateLivyComputeRequest {
	s.LivyServerConf = &v
	return s
}

func (s *CreateLivyComputeRequest) SetLivyVersion(v string) *CreateLivyComputeRequest {
	s.LivyVersion = &v
	return s
}

func (s *CreateLivyComputeRequest) SetMemoryLimit(v string) *CreateLivyComputeRequest {
	s.MemoryLimit = &v
	return s
}

func (s *CreateLivyComputeRequest) SetName(v string) *CreateLivyComputeRequest {
	s.Name = &v
	return s
}

func (s *CreateLivyComputeRequest) SetNetworkName(v string) *CreateLivyComputeRequest {
	s.NetworkName = &v
	return s
}

func (s *CreateLivyComputeRequest) SetQueueName(v string) *CreateLivyComputeRequest {
	s.QueueName = &v
	return s
}

func (s *CreateLivyComputeRequest) SetReleaseVersion(v string) *CreateLivyComputeRequest {
	s.ReleaseVersion = &v
	return s
}

func (s *CreateLivyComputeRequest) SetRegionId(v string) *CreateLivyComputeRequest {
	s.RegionId = &v
	return s
}

type CreateLivyComputeRequestAutoStartConfiguration struct {
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s CreateLivyComputeRequestAutoStartConfiguration) String() string {
	return tea.Prettify(s)
}

func (s CreateLivyComputeRequestAutoStartConfiguration) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeRequestAutoStartConfiguration) SetEnable(v bool) *CreateLivyComputeRequestAutoStartConfiguration {
	s.Enable = &v
	return s
}

type CreateLivyComputeRequestAutoStopConfiguration struct {
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 300
	IdleTimeoutMinutes *int64 `json:"idleTimeoutMinutes,omitempty" xml:"idleTimeoutMinutes,omitempty"`
}

func (s CreateLivyComputeRequestAutoStopConfiguration) String() string {
	return tea.Prettify(s)
}

func (s CreateLivyComputeRequestAutoStopConfiguration) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeRequestAutoStopConfiguration) SetEnable(v bool) *CreateLivyComputeRequestAutoStopConfiguration {
	s.Enable = &v
	return s
}

func (s *CreateLivyComputeRequestAutoStopConfiguration) SetIdleTimeoutMinutes(v int64) *CreateLivyComputeRequestAutoStopConfiguration {
	s.IdleTimeoutMinutes = &v
	return s
}

type CreateLivyComputeResponseBody struct {
	// example:
	//
	// 1000000
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateLivyComputeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateLivyComputeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLivyComputeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeResponseBody) SetCode(v string) *CreateLivyComputeResponseBody {
	s.Code = &v
	return s
}

func (s *CreateLivyComputeResponseBody) SetData(v *CreateLivyComputeResponseBodyData) *CreateLivyComputeResponseBody {
	s.Data = v
	return s
}

func (s *CreateLivyComputeResponseBody) SetMessage(v string) *CreateLivyComputeResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLivyComputeResponseBody) SetRequestId(v string) *CreateLivyComputeResponseBody {
	s.RequestId = &v
	return s
}

type CreateLivyComputeResponseBodyData struct {
	// example:
	//
	// lc-i8xogcdfa4fk3yn1
	LivyComputeId *string `json:"livyComputeId,omitempty" xml:"livyComputeId,omitempty"`
}

func (s CreateLivyComputeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateLivyComputeResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeResponseBodyData) SetLivyComputeId(v string) *CreateLivyComputeResponseBodyData {
	s.LivyComputeId = &v
	return s
}

type CreateLivyComputeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLivyComputeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLivyComputeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLivyComputeResponse) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeResponse) SetHeaders(v map[string]*string) *CreateLivyComputeResponse {
	s.Headers = v
	return s
}

func (s *CreateLivyComputeResponse) SetStatusCode(v int32) *CreateLivyComputeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLivyComputeResponse) SetBody(v *CreateLivyComputeResponseBody) *CreateLivyComputeResponse {
	s.Body = v
	return s
}

type CreateLivyComputeTokenRequest struct {
	AutoExpireConfiguration *CreateLivyComputeTokenRequestAutoExpireConfiguration `json:"autoExpireConfiguration,omitempty" xml:"autoExpireConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// mytoken
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// gs3fy75w4o7hqe5s
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateLivyComputeTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLivyComputeTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeTokenRequest) SetAutoExpireConfiguration(v *CreateLivyComputeTokenRequestAutoExpireConfiguration) *CreateLivyComputeTokenRequest {
	s.AutoExpireConfiguration = v
	return s
}

func (s *CreateLivyComputeTokenRequest) SetName(v string) *CreateLivyComputeTokenRequest {
	s.Name = &v
	return s
}

func (s *CreateLivyComputeTokenRequest) SetToken(v string) *CreateLivyComputeTokenRequest {
	s.Token = &v
	return s
}

func (s *CreateLivyComputeTokenRequest) SetRegionId(v string) *CreateLivyComputeTokenRequest {
	s.RegionId = &v
	return s
}

type CreateLivyComputeTokenRequestAutoExpireConfiguration struct {
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 7
	ExpireDays *int32 `json:"expireDays,omitempty" xml:"expireDays,omitempty"`
}

func (s CreateLivyComputeTokenRequestAutoExpireConfiguration) String() string {
	return tea.Prettify(s)
}

func (s CreateLivyComputeTokenRequestAutoExpireConfiguration) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeTokenRequestAutoExpireConfiguration) SetEnable(v bool) *CreateLivyComputeTokenRequestAutoExpireConfiguration {
	s.Enable = &v
	return s
}

func (s *CreateLivyComputeTokenRequestAutoExpireConfiguration) SetExpireDays(v int32) *CreateLivyComputeTokenRequestAutoExpireConfiguration {
	s.ExpireDays = &v
	return s
}

type CreateLivyComputeTokenResponseBody struct {
	// example:
	//
	// 1000000
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateLivyComputeTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateLivyComputeTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLivyComputeTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeTokenResponseBody) SetCode(v string) *CreateLivyComputeTokenResponseBody {
	s.Code = &v
	return s
}

func (s *CreateLivyComputeTokenResponseBody) SetData(v *CreateLivyComputeTokenResponseBodyData) *CreateLivyComputeTokenResponseBody {
	s.Data = v
	return s
}

func (s *CreateLivyComputeTokenResponseBody) SetMessage(v string) *CreateLivyComputeTokenResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLivyComputeTokenResponseBody) SetRequestId(v string) *CreateLivyComputeTokenResponseBody {
	s.RequestId = &v
	return s
}

type CreateLivyComputeTokenResponseBodyData struct {
	// Token ID。
	//
	// example:
	//
	// lctk-xxxxxxxx
	TokenId *string `json:"tokenId,omitempty" xml:"tokenId,omitempty"`
}

func (s CreateLivyComputeTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateLivyComputeTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeTokenResponseBodyData) SetTokenId(v string) *CreateLivyComputeTokenResponseBodyData {
	s.TokenId = &v
	return s
}

type CreateLivyComputeTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLivyComputeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLivyComputeTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLivyComputeTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeTokenResponse) SetHeaders(v map[string]*string) *CreateLivyComputeTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateLivyComputeTokenResponse) SetStatusCode(v int32) *CreateLivyComputeTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLivyComputeTokenResponse) SetBody(v *CreateLivyComputeTokenResponseBody) *CreateLivyComputeTokenResponse {
	s.Body = v
	return s
}

type CreateProcessDefinitionWithScheduleRequest struct {
	// The email address to receive alerts.
	//
	// example:
	//
	// foo_bar@spark.alert.invalid.com
	AlertEmailAddress *string `json:"alertEmailAddress,omitempty" xml:"alertEmailAddress,omitempty"`
	// The description of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// ods batch workflow
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The execution policy
	//
	// This parameter is required.
	//
	// example:
	//
	// PARALLEL
	ExecutionType *string                                                   `json:"executionType,omitempty" xml:"executionType,omitempty"`
	GlobalParams  []*CreateProcessDefinitionWithScheduleRequestGlobalParams `json:"globalParams,omitempty" xml:"globalParams,omitempty" type:"Repeated"`
	// The name of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// ods_batch_workflow
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The code of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// SS
	ProductNamespace *string `json:"productNamespace,omitempty" xml:"productNamespace,omitempty"`
	// Specifies whether to publish the workflow.
	//
	// example:
	//
	// true
	Publish *bool `json:"publish,omitempty" xml:"publish,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The resource queue.
	//
	// example:
	//
	// root_queue
	ResourceQueue *string `json:"resourceQueue,omitempty" xml:"resourceQueue,omitempty"`
	// The number of retries.
	//
	// example:
	//
	// 1
	RetryTimes *int32 `json:"retryTimes,omitempty" xml:"retryTimes,omitempty"`
	// The ID of the Alibaba Cloud account used by the user who creates the workflow.
	//
	// example:
	//
	// 113***************
	RunAs *string `json:"runAs,omitempty" xml:"runAs,omitempty"`
	// The scheduling settings.
	Schedule *CreateProcessDefinitionWithScheduleRequestSchedule `json:"schedule,omitempty" xml:"schedule,omitempty" type:"Struct"`
	// The tags.
	Tags map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The descriptions of all nodes in the workflow.
	//
	// This parameter is required.
	TaskDefinitionJson []*CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson `json:"taskDefinitionJson,omitempty" xml:"taskDefinitionJson,omitempty" type:"Repeated"`
	// The node parallelism.
	//
	// example:
	//
	// 1
	TaskParallelism *int32 `json:"taskParallelism,omitempty" xml:"taskParallelism,omitempty"`
	// The dependencies of all nodes in the workflow. preTaskCode specifies the ID of an upstream node, and postTaskCode specifies the ID of a downstream node. The ID of each node is unique. If a node does not have an upstream node, set preTaskCode to 0.
	//
	// This parameter is required.
	TaskRelationJson []*CreateProcessDefinitionWithScheduleRequestTaskRelationJson `json:"taskRelationJson,omitempty" xml:"taskRelationJson,omitempty" type:"Repeated"`
	// The default timeout period of the workflow.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleRequest) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetAlertEmailAddress(v string) *CreateProcessDefinitionWithScheduleRequest {
	s.AlertEmailAddress = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetDescription(v string) *CreateProcessDefinitionWithScheduleRequest {
	s.Description = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetExecutionType(v string) *CreateProcessDefinitionWithScheduleRequest {
	s.ExecutionType = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetGlobalParams(v []*CreateProcessDefinitionWithScheduleRequestGlobalParams) *CreateProcessDefinitionWithScheduleRequest {
	s.GlobalParams = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetName(v string) *CreateProcessDefinitionWithScheduleRequest {
	s.Name = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetProductNamespace(v string) *CreateProcessDefinitionWithScheduleRequest {
	s.ProductNamespace = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetPublish(v bool) *CreateProcessDefinitionWithScheduleRequest {
	s.Publish = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetRegionId(v string) *CreateProcessDefinitionWithScheduleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetResourceQueue(v string) *CreateProcessDefinitionWithScheduleRequest {
	s.ResourceQueue = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetRetryTimes(v int32) *CreateProcessDefinitionWithScheduleRequest {
	s.RetryTimes = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetRunAs(v string) *CreateProcessDefinitionWithScheduleRequest {
	s.RunAs = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetSchedule(v *CreateProcessDefinitionWithScheduleRequestSchedule) *CreateProcessDefinitionWithScheduleRequest {
	s.Schedule = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetTags(v map[string]*string) *CreateProcessDefinitionWithScheduleRequest {
	s.Tags = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetTaskDefinitionJson(v []*CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) *CreateProcessDefinitionWithScheduleRequest {
	s.TaskDefinitionJson = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetTaskParallelism(v int32) *CreateProcessDefinitionWithScheduleRequest {
	s.TaskParallelism = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetTaskRelationJson(v []*CreateProcessDefinitionWithScheduleRequestTaskRelationJson) *CreateProcessDefinitionWithScheduleRequest {
	s.TaskRelationJson = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetTimeout(v int32) *CreateProcessDefinitionWithScheduleRequest {
	s.Timeout = &v
	return s
}

type CreateProcessDefinitionWithScheduleRequestGlobalParams struct {
	Direct *string `json:"direct,omitempty" xml:"direct,omitempty"`
	Prop   *string `json:"prop,omitempty" xml:"prop,omitempty"`
	Type   *string `json:"type,omitempty" xml:"type,omitempty"`
	Value  *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleRequestGlobalParams) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleRequestGlobalParams) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleRequestGlobalParams) SetDirect(v string) *CreateProcessDefinitionWithScheduleRequestGlobalParams {
	s.Direct = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestGlobalParams) SetProp(v string) *CreateProcessDefinitionWithScheduleRequestGlobalParams {
	s.Prop = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestGlobalParams) SetType(v string) *CreateProcessDefinitionWithScheduleRequestGlobalParams {
	s.Type = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestGlobalParams) SetValue(v string) *CreateProcessDefinitionWithScheduleRequestGlobalParams {
	s.Value = &v
	return s
}

type CreateProcessDefinitionWithScheduleRequestSchedule struct {
	// The CRON expression that is used for scheduling.
	//
	// example:
	//
	// 0 0 0 	- 	- ?
	Crontab *string `json:"crontab,omitempty" xml:"crontab,omitempty"`
	// The end time of the scheduling.
	//
	// example:
	//
	// 2025-12-23 16:13:27
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The start time of the scheduling.
	//
	// example:
	//
	// 2024-12-23 16:13:27
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The ID of the time zone.
	//
	// example:
	//
	// Asia/Shanghai
	TimezoneId *string `json:"timezoneId,omitempty" xml:"timezoneId,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleRequestSchedule) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleRequestSchedule) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleRequestSchedule) SetCrontab(v string) *CreateProcessDefinitionWithScheduleRequestSchedule {
	s.Crontab = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestSchedule) SetEndTime(v string) *CreateProcessDefinitionWithScheduleRequestSchedule {
	s.EndTime = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestSchedule) SetStartTime(v string) *CreateProcessDefinitionWithScheduleRequestSchedule {
	s.StartTime = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestSchedule) SetTimezoneId(v string) *CreateProcessDefinitionWithScheduleRequestSchedule {
	s.TimezoneId = &v
	return s
}

type CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson struct {
	// The email address to receive alerts.
	//
	// example:
	//
	// foo_bar@spark.alert.invalid.com
	AlertEmailAddress *string `json:"alertEmailAddress,omitempty" xml:"alertEmailAddress,omitempty"`
	// The node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 36************
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// The node description.
	//
	// example:
	//
	// ods transform task
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to send alerts when the node fails.
	//
	// example:
	//
	// false
	FailAlertEnable *bool `json:"failAlertEnable,omitempty" xml:"failAlertEnable,omitempty"`
	// The number of retries when the node fails.
	//
	// example:
	//
	// 1
	FailRetryTimes *int32 `json:"failRetryTimes,omitempty" xml:"failRetryTimes,omitempty"`
	// The name of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// ods_transform_task
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specifies whether to send alerts when the node is started.
	//
	// example:
	//
	// false
	StartAlertEnable *bool `json:"startAlertEnable,omitempty" xml:"startAlertEnable,omitempty"`
	// The tags.
	Tags map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The job parameters.
	//
	// This parameter is required.
	TaskParams *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams `json:"taskParams,omitempty" xml:"taskParams,omitempty" type:"Struct"`
	// The type of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// MigrateData
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// The timeout period of the callback. Unit: seconds.
	//
	// example:
	//
	// 1200
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetAlertEmailAddress(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.AlertEmailAddress = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetCode(v int64) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.Code = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetDescription(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.Description = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetFailAlertEnable(v bool) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.FailAlertEnable = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetFailRetryTimes(v int32) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.FailRetryTimes = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetName(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.Name = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetStartAlertEnable(v bool) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.StartAlertEnable = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetTags(v map[string]*string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.Tags = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetTaskParams(v *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.TaskParams = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetTaskType(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.TaskType = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetTimeout(v int32) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.Timeout = &v
	return s
}

type CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams struct {
	// The displayed version of the Spark engine.
	//
	// example:
	//
	// esr-4.0.0 (Spark 3.5.2, Scala 2.12)
	DisplaySparkVersion *string `json:"displaySparkVersion,omitempty" xml:"displaySparkVersion,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-crhq2h5lhtgju93buhkg
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Specifies whether to enable Fusion engine for acceleration.
	//
	// example:
	//
	// false
	Fusion      *bool                                                                                `json:"fusion,omitempty" xml:"fusion,omitempty"`
	LocalParams []*CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams `json:"localParams,omitempty" xml:"localParams,omitempty" type:"Repeated"`
	// The name of the resource queue on which the job runs.
	//
	// This parameter is required.
	//
	// example:
	//
	// root_queue
	ResourceQueueId *string `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	// The configurations of the Spark job.
	SparkConf []*CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf `json:"sparkConf,omitempty" xml:"sparkConf,omitempty" type:"Repeated"`
	// The number of driver cores of the Spark job.
	//
	// example:
	//
	// 1
	SparkDriverCores *int32 `json:"sparkDriverCores,omitempty" xml:"sparkDriverCores,omitempty"`
	// The size of driver memory of the Spark job.
	//
	// example:
	//
	// 4g
	SparkDriverMemory *int64 `json:"sparkDriverMemory,omitempty" xml:"sparkDriverMemory,omitempty"`
	// The number of executor cores of the Spark job.
	//
	// example:
	//
	// 1
	SparkExecutorCores *int32 `json:"sparkExecutorCores,omitempty" xml:"sparkExecutorCores,omitempty"`
	// The size of executor memory of the Spark job.
	//
	// example:
	//
	// 4g
	SparkExecutorMemory *int64 `json:"sparkExecutorMemory,omitempty" xml:"sparkExecutorMemory,omitempty"`
	// The level of the Spark log.
	//
	// example:
	//
	// INFO
	SparkLogLevel *string `json:"sparkLogLevel,omitempty" xml:"sparkLogLevel,omitempty"`
	// The path where the operational logs of the Spark job are stored.
	SparkLogPath *string `json:"sparkLogPath,omitempty" xml:"sparkLogPath,omitempty"`
	// The version of the Spark engine.
	//
	// example:
	//
	// esr-4.0.0 (Spark 3.5.2, Scala 2.12)
	SparkVersion *string `json:"sparkVersion,omitempty" xml:"sparkVersion,omitempty"`
	// The ID of the data development job.
	//
	// This parameter is required.
	//
	// example:
	//
	// TSK-d87******************
	TaskBizId *string `json:"taskBizId,omitempty" xml:"taskBizId,omitempty"`
	// The type of the Spark job.
	//
	// example:
	//
	// VPC
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// w-d8********
	WorkspaceBizId *string `json:"workspaceBizId,omitempty" xml:"workspaceBizId,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetDisplaySparkVersion(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.DisplaySparkVersion = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetEnvironmentId(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.EnvironmentId = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetFusion(v bool) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.Fusion = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetLocalParams(v []*CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.LocalParams = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetResourceQueueId(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.ResourceQueueId = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkConf(v []*CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkConf = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkDriverCores(v int32) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkDriverCores = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkDriverMemory(v int64) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkDriverMemory = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkExecutorCores(v int32) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkExecutorCores = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkExecutorMemory(v int64) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkExecutorMemory = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkLogLevel(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkLogLevel = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkLogPath(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkLogPath = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkVersion(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkVersion = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetTaskBizId(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.TaskBizId = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetType(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.Type = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetWorkspaceBizId(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.WorkspaceBizId = &v
	return s
}

type CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams struct {
	Direct *string `json:"direct,omitempty" xml:"direct,omitempty"`
	Prop   *string `json:"prop,omitempty" xml:"prop,omitempty"`
	Type   *string `json:"type,omitempty" xml:"type,omitempty"`
	Value  *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) SetDirect(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams {
	s.Direct = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) SetProp(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams {
	s.Prop = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) SetType(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams {
	s.Type = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) SetValue(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams {
	s.Value = &v
	return s
}

type CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf struct {
	// The key of the SparkConf object.
	//
	// example:
	//
	// spark.dynamicAllocation.enabled
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the SparkConf object.
	//
	// example:
	//
	// true
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) SetKey(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf {
	s.Key = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) SetValue(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf {
	s.Value = &v
	return s
}

type CreateProcessDefinitionWithScheduleRequestTaskRelationJson struct {
	// The name of the node topology. You can enter a workflow name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ods batch workflow
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the downstream node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 28************
	PostTaskCode *int64 `json:"postTaskCode,omitempty" xml:"postTaskCode,omitempty"`
	// The version of the downstream node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PostTaskVersion *int32 `json:"postTaskVersion,omitempty" xml:"postTaskVersion,omitempty"`
	// The ID of the upstream node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16************
	PreTaskCode *int64 `json:"preTaskCode,omitempty" xml:"preTaskCode,omitempty"`
	// The version of the upstream node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PreTaskVersion *int32 `json:"preTaskVersion,omitempty" xml:"preTaskVersion,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleRequestTaskRelationJson) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleRequestTaskRelationJson) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskRelationJson) SetName(v string) *CreateProcessDefinitionWithScheduleRequestTaskRelationJson {
	s.Name = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskRelationJson) SetPostTaskCode(v int64) *CreateProcessDefinitionWithScheduleRequestTaskRelationJson {
	s.PostTaskCode = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskRelationJson) SetPostTaskVersion(v int32) *CreateProcessDefinitionWithScheduleRequestTaskRelationJson {
	s.PostTaskVersion = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskRelationJson) SetPreTaskCode(v int64) *CreateProcessDefinitionWithScheduleRequestTaskRelationJson {
	s.PreTaskCode = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskRelationJson) SetPreTaskVersion(v int32) *CreateProcessDefinitionWithScheduleRequestTaskRelationJson {
	s.PreTaskVersion = &v
	return s
}

type CreateProcessDefinitionWithScheduleShrinkRequest struct {
	// The email address to receive alerts.
	//
	// example:
	//
	// foo_bar@spark.alert.invalid.com
	AlertEmailAddress *string `json:"alertEmailAddress,omitempty" xml:"alertEmailAddress,omitempty"`
	// The description of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// ods batch workflow
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The execution policy
	//
	// This parameter is required.
	//
	// example:
	//
	// PARALLEL
	ExecutionType      *string `json:"executionType,omitempty" xml:"executionType,omitempty"`
	GlobalParamsShrink *string `json:"globalParams,omitempty" xml:"globalParams,omitempty"`
	// The name of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// ods_batch_workflow
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The code of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// SS
	ProductNamespace *string `json:"productNamespace,omitempty" xml:"productNamespace,omitempty"`
	// Specifies whether to publish the workflow.
	//
	// example:
	//
	// true
	Publish *bool `json:"publish,omitempty" xml:"publish,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The resource queue.
	//
	// example:
	//
	// root_queue
	ResourceQueue *string `json:"resourceQueue,omitempty" xml:"resourceQueue,omitempty"`
	// The number of retries.
	//
	// example:
	//
	// 1
	RetryTimes *int32 `json:"retryTimes,omitempty" xml:"retryTimes,omitempty"`
	// The ID of the Alibaba Cloud account used by the user who creates the workflow.
	//
	// example:
	//
	// 113***************
	RunAs *string `json:"runAs,omitempty" xml:"runAs,omitempty"`
	// The scheduling settings.
	ScheduleShrink *string `json:"schedule,omitempty" xml:"schedule,omitempty"`
	// The tags.
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The descriptions of all nodes in the workflow.
	//
	// This parameter is required.
	TaskDefinitionJsonShrink *string `json:"taskDefinitionJson,omitempty" xml:"taskDefinitionJson,omitempty"`
	// The node parallelism.
	//
	// example:
	//
	// 1
	TaskParallelism *int32 `json:"taskParallelism,omitempty" xml:"taskParallelism,omitempty"`
	// The dependencies of all nodes in the workflow. preTaskCode specifies the ID of an upstream node, and postTaskCode specifies the ID of a downstream node. The ID of each node is unique. If a node does not have an upstream node, set preTaskCode to 0.
	//
	// This parameter is required.
	TaskRelationJsonShrink *string `json:"taskRelationJson,omitempty" xml:"taskRelationJson,omitempty"`
	// The default timeout period of the workflow.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetAlertEmailAddress(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.AlertEmailAddress = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetDescription(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetExecutionType(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.ExecutionType = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetGlobalParamsShrink(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.GlobalParamsShrink = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetName(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetProductNamespace(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.ProductNamespace = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetPublish(v bool) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.Publish = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetRegionId(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetResourceQueue(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.ResourceQueue = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetRetryTimes(v int32) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.RetryTimes = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetRunAs(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.RunAs = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetScheduleShrink(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.ScheduleShrink = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetTagsShrink(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetTaskDefinitionJsonShrink(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.TaskDefinitionJsonShrink = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetTaskParallelism(v int32) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.TaskParallelism = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetTaskRelationJsonShrink(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.TaskRelationJsonShrink = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetTimeout(v int32) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.Timeout = &v
	return s
}

type CreateProcessDefinitionWithScheduleResponseBody struct {
	// The code that is returned by the backend server.
	//
	// example:
	//
	// 1400009
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *CreateProcessDefinitionWithScheduleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Indicates whether the request failed.
	//
	// example:
	//
	// false
	Failed *string `json:"failed,omitempty" xml:"failed,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The description of the returned code.
	//
	// example:
	//
	// No permission for resource action
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) SetCode(v int32) *CreateProcessDefinitionWithScheduleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) SetData(v *CreateProcessDefinitionWithScheduleResponseBodyData) *CreateProcessDefinitionWithScheduleResponseBody {
	s.Data = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) SetFailed(v string) *CreateProcessDefinitionWithScheduleResponseBody {
	s.Failed = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) SetHttpStatusCode(v int32) *CreateProcessDefinitionWithScheduleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) SetMsg(v string) *CreateProcessDefinitionWithScheduleResponseBody {
	s.Msg = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) SetRequestId(v string) *CreateProcessDefinitionWithScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponseBody) SetSuccess(v string) *CreateProcessDefinitionWithScheduleResponseBody {
	s.Success = &v
	return s
}

type CreateProcessDefinitionWithScheduleResponseBodyData struct {
	// The workflow ID.
	//
	// example:
	//
	// 160************
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// The serial number of the workflow.
	//
	// example:
	//
	// 12342
	Id *int32 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleResponseBodyData) SetCode(v int64) *CreateProcessDefinitionWithScheduleResponseBodyData {
	s.Code = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponseBodyData) SetId(v int32) *CreateProcessDefinitionWithScheduleResponseBodyData {
	s.Id = &v
	return s
}

type CreateProcessDefinitionWithScheduleResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProcessDefinitionWithScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleResponse) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleResponse) SetHeaders(v map[string]*string) *CreateProcessDefinitionWithScheduleResponse {
	s.Headers = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponse) SetStatusCode(v int32) *CreateProcessDefinitionWithScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleResponse) SetBody(v *CreateProcessDefinitionWithScheduleResponseBody) *CreateProcessDefinitionWithScheduleResponse {
	s.Body = v
	return s
}

type CreateSessionClusterRequest struct {
	// The Spark configurations.
	ApplicationConfigs []*CreateSessionClusterRequestApplicationConfigs `json:"applicationConfigs,omitempty" xml:"applicationConfigs,omitempty" type:"Repeated"`
	// Specifies whether to enable automatic startup.
	//
	// 	- true
	//
	// 	- false
	AutoStartConfiguration *CreateSessionClusterRequestAutoStartConfiguration `json:"autoStartConfiguration,omitempty" xml:"autoStartConfiguration,omitempty" type:"Struct"`
	// The automatic termination configuration.
	AutoStopConfiguration *CreateSessionClusterRequestAutoStopConfiguration `json:"autoStopConfiguration,omitempty" xml:"autoStopConfiguration,omitempty" type:"Struct"`
	// The version of the Spark engine.
	//
	// example:
	//
	// esr-3.3.1
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// The ID of the Python environment. This parameter takes effect only for notebook sessions.
	//
	// example:
	//
	// env-cpv569tlhtgndjl86t40
	EnvId *string `json:"envId,omitempty" xml:"envId,omitempty"`
	// Specifies whether to enable Fusion engine for acceleration.
	//
	// example:
	//
	// false
	Fusion *bool `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// The session type.
	//
	// 	- SQL
	//
	// 	- NOTEBOOK
	//
	// example:
	//
	// SQL
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// spark_job_name
	Name                  *string `json:"name,omitempty" xml:"name,omitempty"`
	PublicEndpointEnabled *bool   `json:"publicEndpointEnabled,omitempty" xml:"publicEndpointEnabled,omitempty"`
	// The queue name.
	//
	// example:
	//
	// root_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// The version number of Spark.
	//
	// example:
	//
	// esr-3.3.1
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateSessionClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSessionClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateSessionClusterRequest) SetApplicationConfigs(v []*CreateSessionClusterRequestApplicationConfigs) *CreateSessionClusterRequest {
	s.ApplicationConfigs = v
	return s
}

func (s *CreateSessionClusterRequest) SetAutoStartConfiguration(v *CreateSessionClusterRequestAutoStartConfiguration) *CreateSessionClusterRequest {
	s.AutoStartConfiguration = v
	return s
}

func (s *CreateSessionClusterRequest) SetAutoStopConfiguration(v *CreateSessionClusterRequestAutoStopConfiguration) *CreateSessionClusterRequest {
	s.AutoStopConfiguration = v
	return s
}

func (s *CreateSessionClusterRequest) SetDisplayReleaseVersion(v string) *CreateSessionClusterRequest {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *CreateSessionClusterRequest) SetEnvId(v string) *CreateSessionClusterRequest {
	s.EnvId = &v
	return s
}

func (s *CreateSessionClusterRequest) SetFusion(v bool) *CreateSessionClusterRequest {
	s.Fusion = &v
	return s
}

func (s *CreateSessionClusterRequest) SetKind(v string) *CreateSessionClusterRequest {
	s.Kind = &v
	return s
}

func (s *CreateSessionClusterRequest) SetName(v string) *CreateSessionClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateSessionClusterRequest) SetPublicEndpointEnabled(v bool) *CreateSessionClusterRequest {
	s.PublicEndpointEnabled = &v
	return s
}

func (s *CreateSessionClusterRequest) SetQueueName(v string) *CreateSessionClusterRequest {
	s.QueueName = &v
	return s
}

func (s *CreateSessionClusterRequest) SetReleaseVersion(v string) *CreateSessionClusterRequest {
	s.ReleaseVersion = &v
	return s
}

func (s *CreateSessionClusterRequest) SetRegionId(v string) *CreateSessionClusterRequest {
	s.RegionId = &v
	return s
}

type CreateSessionClusterRequestApplicationConfigs struct {
	// The name of the configuration file.
	//
	// example:
	//
	// spark-defaults.conf
	ConfigFileName *string `json:"configFileName,omitempty" xml:"configFileName,omitempty"`
	// The key of SparkConf.
	//
	// example:
	//
	// spark.app.name
	ConfigItemKey *string `json:"configItemKey,omitempty" xml:"configItemKey,omitempty"`
	// The value of SparkConf.
	//
	// example:
	//
	// test
	ConfigItemValue *string `json:"configItemValue,omitempty" xml:"configItemValue,omitempty"`
}

func (s CreateSessionClusterRequestApplicationConfigs) String() string {
	return tea.Prettify(s)
}

func (s CreateSessionClusterRequestApplicationConfigs) GoString() string {
	return s.String()
}

func (s *CreateSessionClusterRequestApplicationConfigs) SetConfigFileName(v string) *CreateSessionClusterRequestApplicationConfigs {
	s.ConfigFileName = &v
	return s
}

func (s *CreateSessionClusterRequestApplicationConfigs) SetConfigItemKey(v string) *CreateSessionClusterRequestApplicationConfigs {
	s.ConfigItemKey = &v
	return s
}

func (s *CreateSessionClusterRequestApplicationConfigs) SetConfigItemValue(v string) *CreateSessionClusterRequestApplicationConfigs {
	s.ConfigItemValue = &v
	return s
}

type CreateSessionClusterRequestAutoStartConfiguration struct {
	// Specifies whether to enable automatic startup.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s CreateSessionClusterRequestAutoStartConfiguration) String() string {
	return tea.Prettify(s)
}

func (s CreateSessionClusterRequestAutoStartConfiguration) GoString() string {
	return s.String()
}

func (s *CreateSessionClusterRequestAutoStartConfiguration) SetEnable(v bool) *CreateSessionClusterRequestAutoStartConfiguration {
	s.Enable = &v
	return s
}

type CreateSessionClusterRequestAutoStopConfiguration struct {
	// Specifies whether to enable automatic termination.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The idle timeout period. The session is automatically terminated when the idle timeout period is exceeded.
	//
	// example:
	//
	// 60
	IdleTimeoutMinutes *int `json:"idleTimeoutMinutes,omitempty" xml:"idleTimeoutMinutes,omitempty"`
}

func (s CreateSessionClusterRequestAutoStopConfiguration) String() string {
	return tea.Prettify(s)
}

func (s CreateSessionClusterRequestAutoStopConfiguration) GoString() string {
	return s.String()
}

func (s *CreateSessionClusterRequestAutoStopConfiguration) SetEnable(v bool) *CreateSessionClusterRequestAutoStopConfiguration {
	s.Enable = &v
	return s
}

func (s *CreateSessionClusterRequestAutoStopConfiguration) SetIdleTimeoutMinutes(v int) *CreateSessionClusterRequestAutoStopConfiguration {
	s.IdleTimeoutMinutes = &v
	return s
}

type CreateSessionClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// w-******
	SessionClusterId *string `json:"sessionClusterId,omitempty" xml:"sessionClusterId,omitempty"`
}

func (s CreateSessionClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSessionClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSessionClusterResponseBody) SetRequestId(v string) *CreateSessionClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSessionClusterResponseBody) SetSessionClusterId(v string) *CreateSessionClusterResponseBody {
	s.SessionClusterId = &v
	return s
}

type CreateSessionClusterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSessionClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSessionClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSessionClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateSessionClusterResponse) SetHeaders(v map[string]*string) *CreateSessionClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateSessionClusterResponse) SetStatusCode(v int32) *CreateSessionClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSessionClusterResponse) SetBody(v *CreateSessionClusterResponseBody) *CreateSessionClusterResponse {
	s.Body = v
	return s
}

type CreateSqlStatementRequest struct {
	// The SQL code. You can specify one or more SQL statements.
	//
	// example:
	//
	// SHOW TABLES
	CodeContent *string `json:"codeContent,omitempty" xml:"codeContent,omitempty"`
	// The default Data Lake Formation (DLF) catalog ID.
	//
	// example:
	//
	// default_catalog
	DefaultCatalog *string `json:"defaultCatalog,omitempty" xml:"defaultCatalog,omitempty"`
	// The name of the default database.
	//
	// example:
	//
	// default
	DefaultDatabase *string `json:"defaultDatabase,omitempty" xml:"defaultDatabase,omitempty"`
	// The maximum number of entries to return. Valid values: 1 to 10000.
	//
	// example:
	//
	// 1000
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The SQL session ID. You can create an SQL session in the workspace created in EMR Serverless Spark.
	//
	// example:
	//
	// sc-dfahdfjafhajd****
	SqlComputeId *string `json:"sqlComputeId,omitempty" xml:"sqlComputeId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateSqlStatementRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSqlStatementRequest) GoString() string {
	return s.String()
}

func (s *CreateSqlStatementRequest) SetCodeContent(v string) *CreateSqlStatementRequest {
	s.CodeContent = &v
	return s
}

func (s *CreateSqlStatementRequest) SetDefaultCatalog(v string) *CreateSqlStatementRequest {
	s.DefaultCatalog = &v
	return s
}

func (s *CreateSqlStatementRequest) SetDefaultDatabase(v string) *CreateSqlStatementRequest {
	s.DefaultDatabase = &v
	return s
}

func (s *CreateSqlStatementRequest) SetLimit(v int32) *CreateSqlStatementRequest {
	s.Limit = &v
	return s
}

func (s *CreateSqlStatementRequest) SetSqlComputeId(v string) *CreateSqlStatementRequest {
	s.SqlComputeId = &v
	return s
}

func (s *CreateSqlStatementRequest) SetRegionId(v string) *CreateSqlStatementRequest {
	s.RegionId = &v
	return s
}

type CreateSqlStatementResponseBody struct {
	// The data returned.
	Data *CreateSqlStatementResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateSqlStatementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSqlStatementResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSqlStatementResponseBody) SetData(v *CreateSqlStatementResponseBodyData) *CreateSqlStatementResponseBody {
	s.Data = v
	return s
}

func (s *CreateSqlStatementResponseBody) SetRequestId(v string) *CreateSqlStatementResponseBody {
	s.RequestId = &v
	return s
}

type CreateSqlStatementResponseBodyData struct {
	// The interactive query ID.
	//
	// example:
	//
	// st-1231dfafadfa***
	StatementId *string `json:"statementId,omitempty" xml:"statementId,omitempty"`
}

func (s CreateSqlStatementResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateSqlStatementResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSqlStatementResponseBodyData) SetStatementId(v string) *CreateSqlStatementResponseBodyData {
	s.StatementId = &v
	return s
}

type CreateSqlStatementResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSqlStatementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSqlStatementResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSqlStatementResponse) GoString() string {
	return s.String()
}

func (s *CreateSqlStatementResponse) SetHeaders(v map[string]*string) *CreateSqlStatementResponse {
	s.Headers = v
	return s
}

func (s *CreateSqlStatementResponse) SetStatusCode(v int32) *CreateSqlStatementResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSqlStatementResponse) SetBody(v *CreateSqlStatementResponseBody) *CreateSqlStatementResponse {
	s.Body = v
	return s
}

type CreateWorkspaceRequest struct {
	// Specifies whether to enable auto-renewal. This parameter is required only if the paymentType parameter is set to Pre.
	//
	// example:
	//
	// false
	AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// The auto-renewal duration. This parameter is required only if the paymentType parameter is set to Pre.
	//
	// example:
	//
	// 100
	AutoRenewPeriod *string `json:"autoRenewPeriod,omitempty" xml:"autoRenewPeriod,omitempty"`
	// The unit of the auto-renewal duration. This parameter is required only if the paymentType parameter is set to Pre.
	//
	// example:
	//
	// month
	AutoRenewPeriodUnit *string `json:"autoRenewPeriodUnit,omitempty" xml:"autoRenewPeriodUnit,omitempty"`
	// Specifies whether to automatically start a session.
	//
	// example:
	//
	// false
	AutoStartSessionCluster *bool `json:"autoStartSessionCluster,omitempty" xml:"autoStartSessionCluster,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 8e6aae2810c8f67229ca70bb31cd****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The information of the Data Lake Formation (DLF) catalog.
	//
	// example:
	//
	// 123xxxxx
	DlfCatalogId *string `json:"dlfCatalogId,omitempty" xml:"dlfCatalogId,omitempty"`
	// The version of DLF.
	//
	// example:
	//
	// dlf1.0
	DlfType *string `json:"dlfType,omitempty" xml:"dlfType,omitempty"`
	// The subscription period. This parameter is required only if the paymentType parameter is set to Pre.
	//
	// example:
	//
	// 12452
	Duration *string `json:"duration,omitempty" xml:"duration,omitempty"`
	// The name of the Object Storage Service (OSS) bucket.
	//
	// example:
	//
	// oss://test-bucket/
	OssBucket *string `json:"ossBucket,omitempty" xml:"ossBucket,omitempty"`
	// The unit of the subscription duration.
	//
	// example:
	//
	// 1000
	PaymentDurationUnit *string `json:"paymentDurationUnit,omitempty" xml:"paymentDurationUnit,omitempty"`
	// The billing method. Valid values:
	//
	// 	- PayAsYouGo
	//
	// 	- Pre
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The name of the role used to run Spark jobs.
	//
	// example:
	//
	// AliyunEMRSparkJobRunDefaultRole
	RamRoleName *string `json:"ramRoleName,omitempty" xml:"ramRoleName,omitempty"`
	// The type of the version.
	//
	// example:
	//
	// pro
	ReleaseType *string `json:"releaseType,omitempty" xml:"releaseType,omitempty"`
	// The resource specifications.
	ResourceSpec *CreateWorkspaceRequestResourceSpec `json:"resourceSpec,omitempty" xml:"resourceSpec,omitempty" type:"Struct"`
	// if can be null:
	// false
	Tag []*CreateWorkspaceRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
	// The name of the workspace.
	//
	// example:
	//
	// default
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequest) SetAutoRenew(v string) *CreateWorkspaceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateWorkspaceRequest) SetAutoRenewPeriod(v string) *CreateWorkspaceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateWorkspaceRequest) SetAutoRenewPeriodUnit(v string) *CreateWorkspaceRequest {
	s.AutoRenewPeriodUnit = &v
	return s
}

func (s *CreateWorkspaceRequest) SetAutoStartSessionCluster(v bool) *CreateWorkspaceRequest {
	s.AutoStartSessionCluster = &v
	return s
}

func (s *CreateWorkspaceRequest) SetClientToken(v string) *CreateWorkspaceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateWorkspaceRequest) SetDlfCatalogId(v string) *CreateWorkspaceRequest {
	s.DlfCatalogId = &v
	return s
}

func (s *CreateWorkspaceRequest) SetDlfType(v string) *CreateWorkspaceRequest {
	s.DlfType = &v
	return s
}

func (s *CreateWorkspaceRequest) SetDuration(v string) *CreateWorkspaceRequest {
	s.Duration = &v
	return s
}

func (s *CreateWorkspaceRequest) SetOssBucket(v string) *CreateWorkspaceRequest {
	s.OssBucket = &v
	return s
}

func (s *CreateWorkspaceRequest) SetPaymentDurationUnit(v string) *CreateWorkspaceRequest {
	s.PaymentDurationUnit = &v
	return s
}

func (s *CreateWorkspaceRequest) SetPaymentType(v string) *CreateWorkspaceRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateWorkspaceRequest) SetRamRoleName(v string) *CreateWorkspaceRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateWorkspaceRequest) SetReleaseType(v string) *CreateWorkspaceRequest {
	s.ReleaseType = &v
	return s
}

func (s *CreateWorkspaceRequest) SetResourceSpec(v *CreateWorkspaceRequestResourceSpec) *CreateWorkspaceRequest {
	s.ResourceSpec = v
	return s
}

func (s *CreateWorkspaceRequest) SetTag(v []*CreateWorkspaceRequestTag) *CreateWorkspaceRequest {
	s.Tag = v
	return s
}

func (s *CreateWorkspaceRequest) SetWorkspaceName(v string) *CreateWorkspaceRequest {
	s.WorkspaceName = &v
	return s
}

func (s *CreateWorkspaceRequest) SetRegionId(v string) *CreateWorkspaceRequest {
	s.RegionId = &v
	return s
}

type CreateWorkspaceRequestResourceSpec struct {
	// The maximum resource quota for a workspace.
	//
	// example:
	//
	// 1000
	Cu *string `json:"cu,omitempty" xml:"cu,omitempty"`
}

func (s CreateWorkspaceRequestResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequestResourceSpec) SetCu(v string) *CreateWorkspaceRequestResourceSpec {
	s.Cu = &v
	return s
}

type CreateWorkspaceRequestTag struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateWorkspaceRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequestTag) SetKey(v string) *CreateWorkspaceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateWorkspaceRequestTag) SetValue(v string) *CreateWorkspaceRequestTag {
	s.Value = &v
	return s
}

type CreateWorkspaceResponseBody struct {
	// The operation ID.
	//
	// example:
	//
	// op-******
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// w-******
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBody) SetOperationId(v string) *CreateWorkspaceResponseBody {
	s.OperationId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetRequestId(v string) *CreateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetWorkspaceId(v string) *CreateWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

type CreateWorkspaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponse) SetHeaders(v map[string]*string) *CreateWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkspaceResponse) SetStatusCode(v int32) *CreateWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkspaceResponse) SetBody(v *CreateWorkspaceResponseBody) *CreateWorkspaceResponse {
	s.Body = v
	return s
}

type DeleteLivyComputeRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DeleteLivyComputeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLivyComputeRequest) GoString() string {
	return s.String()
}

func (s *DeleteLivyComputeRequest) SetRegionId(v string) *DeleteLivyComputeRequest {
	s.RegionId = &v
	return s
}

type DeleteLivyComputeResponseBody struct {
	// example:
	//
	// 1000000
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteLivyComputeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLivyComputeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLivyComputeResponseBody) SetCode(v string) *DeleteLivyComputeResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteLivyComputeResponseBody) SetMessage(v string) *DeleteLivyComputeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteLivyComputeResponseBody) SetRequestId(v string) *DeleteLivyComputeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLivyComputeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLivyComputeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLivyComputeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLivyComputeResponse) GoString() string {
	return s.String()
}

func (s *DeleteLivyComputeResponse) SetHeaders(v map[string]*string) *DeleteLivyComputeResponse {
	s.Headers = v
	return s
}

func (s *DeleteLivyComputeResponse) SetStatusCode(v int32) *DeleteLivyComputeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLivyComputeResponse) SetBody(v *DeleteLivyComputeResponseBody) *DeleteLivyComputeResponse {
	s.Body = v
	return s
}

type DeleteLivyComputeTokenRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DeleteLivyComputeTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLivyComputeTokenRequest) GoString() string {
	return s.String()
}

func (s *DeleteLivyComputeTokenRequest) SetRegionId(v string) *DeleteLivyComputeTokenRequest {
	s.RegionId = &v
	return s
}

type DeleteLivyComputeTokenResponseBody struct {
	// example:
	//
	// 1000000
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteLivyComputeTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLivyComputeTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLivyComputeTokenResponseBody) SetCode(v string) *DeleteLivyComputeTokenResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteLivyComputeTokenResponseBody) SetMessage(v string) *DeleteLivyComputeTokenResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteLivyComputeTokenResponseBody) SetRequestId(v string) *DeleteLivyComputeTokenResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLivyComputeTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLivyComputeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLivyComputeTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLivyComputeTokenResponse) GoString() string {
	return s.String()
}

func (s *DeleteLivyComputeTokenResponse) SetHeaders(v map[string]*string) *DeleteLivyComputeTokenResponse {
	s.Headers = v
	return s
}

func (s *DeleteLivyComputeTokenResponse) SetStatusCode(v int32) *DeleteLivyComputeTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLivyComputeTokenResponse) SetBody(v *DeleteLivyComputeTokenResponseBody) *DeleteLivyComputeTokenResponse {
	s.Body = v
	return s
}

type EditWorkspaceQueueRequest struct {
	Environments []*string                              `json:"environments,omitempty" xml:"environments,omitempty" type:"Repeated"`
	ResourceSpec *EditWorkspaceQueueRequestResourceSpec `json:"resourceSpec,omitempty" xml:"resourceSpec,omitempty" type:"Struct"`
	// example:
	//
	// w-975bcfda9625****
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// example:
	//
	// dev_queue
	WorkspaceQueueName *string `json:"workspaceQueueName,omitempty" xml:"workspaceQueueName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s EditWorkspaceQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s EditWorkspaceQueueRequest) GoString() string {
	return s.String()
}

func (s *EditWorkspaceQueueRequest) SetEnvironments(v []*string) *EditWorkspaceQueueRequest {
	s.Environments = v
	return s
}

func (s *EditWorkspaceQueueRequest) SetResourceSpec(v *EditWorkspaceQueueRequestResourceSpec) *EditWorkspaceQueueRequest {
	s.ResourceSpec = v
	return s
}

func (s *EditWorkspaceQueueRequest) SetWorkspaceId(v string) *EditWorkspaceQueueRequest {
	s.WorkspaceId = &v
	return s
}

func (s *EditWorkspaceQueueRequest) SetWorkspaceQueueName(v string) *EditWorkspaceQueueRequest {
	s.WorkspaceQueueName = &v
	return s
}

func (s *EditWorkspaceQueueRequest) SetRegionId(v string) *EditWorkspaceQueueRequest {
	s.RegionId = &v
	return s
}

type EditWorkspaceQueueRequestResourceSpec struct {
	// example:
	//
	// 1000
	Cu *int64 `json:"cu,omitempty" xml:"cu,omitempty"`
}

func (s EditWorkspaceQueueRequestResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s EditWorkspaceQueueRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *EditWorkspaceQueueRequestResourceSpec) SetCu(v int64) *EditWorkspaceQueueRequestResourceSpec {
	s.Cu = &v
	return s
}

type EditWorkspaceQueueResponseBody struct {
	// 请求ID。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EditWorkspaceQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditWorkspaceQueueResponseBody) GoString() string {
	return s.String()
}

func (s *EditWorkspaceQueueResponseBody) SetRequestId(v string) *EditWorkspaceQueueResponseBody {
	s.RequestId = &v
	return s
}

type EditWorkspaceQueueResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EditWorkspaceQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditWorkspaceQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s EditWorkspaceQueueResponse) GoString() string {
	return s.String()
}

func (s *EditWorkspaceQueueResponse) SetHeaders(v map[string]*string) *EditWorkspaceQueueResponse {
	s.Headers = v
	return s
}

func (s *EditWorkspaceQueueResponse) SetStatusCode(v int32) *EditWorkspaceQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *EditWorkspaceQueueResponse) SetBody(v *EditWorkspaceQueueResponseBody) *EditWorkspaceQueueResponse {
	s.Body = v
	return s
}

type GetCuHoursRequest struct {
	// The end time of the query time range.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-01-08 00:00:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The start time of the query time range.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-01-01 00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetCuHoursRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCuHoursRequest) GoString() string {
	return s.String()
}

func (s *GetCuHoursRequest) SetEndTime(v string) *GetCuHoursRequest {
	s.EndTime = &v
	return s
}

func (s *GetCuHoursRequest) SetStartTime(v string) *GetCuHoursRequest {
	s.StartTime = &v
	return s
}

type GetCuHoursResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// {
	//
	//     "cuHours": "{2025-01-09 00:00:00=2.033333, 2025-01-09 01:00:00=2.033333, 2025-01-09 02:00:00=2.033333, 2025-01-09 03:00:00=2.033333, 2025-01-09 04:00:00=2.033333, 2025-01-09 05:00:00=2.033333, 2025-01-09 06:00:00=2.033333, 2025-01-09 07:00:00=2.033333, 2025-01-09 08:00:00=2.033333, 2025-01-09 09:00:00=1.933333, 2025-01-09 10:00:00=2.133333, 2025-01-09 11:00:00=3.100000, 2025-01-09 12:00:00=2.900000}"
	//
	// }
	Data *GetCuHoursResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetCuHoursResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCuHoursResponseBody) GoString() string {
	return s.String()
}

func (s *GetCuHoursResponseBody) SetData(v *GetCuHoursResponseBodyData) *GetCuHoursResponseBody {
	s.Data = v
	return s
}

func (s *GetCuHoursResponseBody) SetRequestId(v string) *GetCuHoursResponseBody {
	s.RequestId = &v
	return s
}

type GetCuHoursResponseBodyData struct {
	// The number of CU-hours consumed by a queue during a specified cycle. The value is an estimated value. Refer to your Alibaba Cloud bill for the actual number of consumed CU-hours.
	//
	// example:
	//
	// {2025-01-09 00:00:00=2.033333, 2025-01-09 01:00:00=2.033333, 2025-01-09 02:00:00=2.033333, 2025-01-09 03:00:00=2.033333, 2025-01-09 04:00:00=2.033333, 2025-01-09 05:00:00=2.033333, 2025-01-09 06:00:00=2.033333, 2025-01-09 07:00:00=2.033333, 2025-01-09 08:00:00=2.033333, 2025-01-09 09:00:00=1.933333, 2025-01-09 10:00:00=2.133333, 2025-01-09 11:00:00=3.100000, 2025-01-09 12:00:00=2.900000}
	CuHours *string `json:"cuHours,omitempty" xml:"cuHours,omitempty"`
}

func (s GetCuHoursResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCuHoursResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCuHoursResponseBodyData) SetCuHours(v string) *GetCuHoursResponseBodyData {
	s.CuHours = &v
	return s
}

type GetCuHoursResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCuHoursResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCuHoursResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCuHoursResponse) GoString() string {
	return s.String()
}

func (s *GetCuHoursResponse) SetHeaders(v map[string]*string) *GetCuHoursResponse {
	s.Headers = v
	return s
}

func (s *GetCuHoursResponse) SetStatusCode(v int32) *GetCuHoursResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCuHoursResponse) SetBody(v *GetCuHoursResponseBody) *GetCuHoursResponse {
	s.Body = v
	return s
}

type GetDoctorApplicationRequest struct {
	// The language of diagnostic information.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"locale,omitempty" xml:"locale,omitempty"`
	// The query time.
	//
	// example:
	//
	// 2024-01-01
	QueryTime *string `json:"queryTime,omitempty" xml:"queryTime,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetDoctorApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDoctorApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorApplicationRequest) SetLocale(v string) *GetDoctorApplicationRequest {
	s.Locale = &v
	return s
}

func (s *GetDoctorApplicationRequest) SetQueryTime(v string) *GetDoctorApplicationRequest {
	s.QueryTime = &v
	return s
}

func (s *GetDoctorApplicationRequest) SetRegionId(v string) *GetDoctorApplicationRequest {
	s.RegionId = &v
	return s
}

type GetDoctorApplicationResponseBody struct {
	// The data returned.
	Data *GetDoctorApplicationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s GetDoctorApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDoctorApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetDoctorApplicationResponseBody) SetData(v *GetDoctorApplicationResponseBodyData) *GetDoctorApplicationResponseBody {
	s.Data = v
	return s
}

type GetDoctorApplicationResponseBodyData struct {
	// The diagnostics list.
	Suggestions []*string `json:"suggestions,omitempty" xml:"suggestions,omitempty" type:"Repeated"`
}

func (s GetDoctorApplicationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDoctorApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDoctorApplicationResponseBodyData) SetSuggestions(v []*string) *GetDoctorApplicationResponseBodyData {
	s.Suggestions = v
	return s
}

type GetDoctorApplicationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDoctorApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDoctorApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDoctorApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetDoctorApplicationResponse) SetHeaders(v map[string]*string) *GetDoctorApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetDoctorApplicationResponse) SetStatusCode(v int32) *GetDoctorApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDoctorApplicationResponse) SetBody(v *GetDoctorApplicationResponseBody) *GetDoctorApplicationResponse {
	s.Body = v
	return s
}

type GetJobRunRequest struct {
	// The region ID.
	//
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
	// The details of the job.
	JobRun *GetJobRunResponseBodyJobRun `json:"jobRun,omitempty" xml:"jobRun,omitempty" type:"Struct"`
	// The request ID.
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
	// The code type of the job. Valid values:
	//
	// 	- SQL
	//
	// 	- JAR
	//
	// 	- PYTHON
	//
	// example:
	//
	// SQL
	CodeType *string `json:"codeType,omitempty" xml:"codeType,omitempty"`
	// The configurations of the Spark jobs.
	ConfigurationOverrides *GetJobRunResponseBodyJobRunConfigurationOverrides `json:"configurationOverrides,omitempty" xml:"configurationOverrides,omitempty" type:"Struct"`
	// The version of the Spark engine.
	//
	// example:
	//
	// esr-4.0.0 (Spark 3.5.2, Scala 2.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// The end time of the job.
	//
	// example:
	//
	// 1684119314000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-cpv569tlhtgndjl8****
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The timeout period of the job.
	//
	// example:
	//
	// 3600
	ExecutionTimeoutSeconds *int32 `json:"executionTimeoutSeconds,omitempty" xml:"executionTimeoutSeconds,omitempty"`
	// Indicates whether the Fusion engine is used for acceleration.
	//
	// example:
	//
	// false
	Fusion *bool `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// The information about Spark Driver.
	JobDriver *JobDriver `json:"jobDriver,omitempty" xml:"jobDriver,omitempty"`
	// The job ID.
	//
	// example:
	//
	// jr-231231
	JobRunId *string `json:"jobRunId,omitempty" xml:"jobRunId,omitempty"`
	// The path where the operational logs are stored.
	Log *RunLog `json:"log,omitempty" xml:"log,omitempty"`
	// The job name.
	//
	// example:
	//
	// jobName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The version of the Spark engine on which the job runs.
	//
	// example:
	//
	// esr-3.3.1
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// The ID of the user who created the job.
	//
	// example:
	//
	// 1509789347011222
	ResourceOwnerId *string `json:"resourceOwnerId,omitempty" xml:"resourceOwnerId,omitempty"`
	// The name of the queue on which the job runs.
	//
	// example:
	//
	// root_queue
	ResourceQueueId *string `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	// The job state.
	//
	// example:
	//
	// Running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The reason of the job status change.
	StateChangeReason *GetJobRunResponseBodyJobRunStateChangeReason `json:"stateChangeReason,omitempty" xml:"stateChangeReason,omitempty" type:"Struct"`
	// The time when the job was submitted.
	//
	// example:
	//
	// 1684119314000
	SubmitTime *int64 `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
	// The tags of the job.
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The web UI of the job.
	//
	// example:
	//
	// http://spark-ui
	WebUI *string `json:"webUI,omitempty" xml:"webUI,omitempty"`
	// The workspace ID.
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

func (s *GetJobRunResponseBodyJobRun) SetDisplayReleaseVersion(v string) *GetJobRunResponseBodyJobRun {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetEndTime(v int64) *GetJobRunResponseBodyJobRun {
	s.EndTime = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetEnvironmentId(v string) *GetJobRunResponseBodyJobRun {
	s.EnvironmentId = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetExecutionTimeoutSeconds(v int32) *GetJobRunResponseBodyJobRun {
	s.ExecutionTimeoutSeconds = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetFusion(v bool) *GetJobRunResponseBodyJobRun {
	s.Fusion = &v
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
	// The configurations.
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
	// The error code.
	//
	// example:
	//
	// ERR-100000
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error message.
	//
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

type GetLivyComputeRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetLivyComputeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLivyComputeRequest) GoString() string {
	return s.String()
}

func (s *GetLivyComputeRequest) SetRegionId(v string) *GetLivyComputeRequest {
	s.RegionId = &v
	return s
}

type GetLivyComputeResponseBody struct {
	// example:
	//
	// 1000000
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetLivyComputeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetLivyComputeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLivyComputeResponseBody) GoString() string {
	return s.String()
}

func (s *GetLivyComputeResponseBody) SetCode(v string) *GetLivyComputeResponseBody {
	s.Code = &v
	return s
}

func (s *GetLivyComputeResponseBody) SetData(v *GetLivyComputeResponseBodyData) *GetLivyComputeResponseBody {
	s.Data = v
	return s
}

func (s *GetLivyComputeResponseBody) SetMessage(v string) *GetLivyComputeResponseBody {
	s.Message = &v
	return s
}

func (s *GetLivyComputeResponseBody) SetRequestId(v string) *GetLivyComputeResponseBody {
	s.RequestId = &v
	return s
}

type GetLivyComputeResponseBodyData struct {
	// example:
	//
	// Token
	AuthType              *string                                              `json:"authType,omitempty" xml:"authType,omitempty"`
	AutoStopConfiguration *GetLivyComputeResponseBodyDataAutoStopConfiguration `json:"autoStopConfiguration,omitempty" xml:"autoStopConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// lc-xxxxxxxxxxxxx
	ComputeId *string `json:"computeId,omitempty" xml:"computeId,omitempty"`
	// example:
	//
	// 1
	CpuLimit *string `json:"cpuLimit,omitempty" xml:"cpuLimit,omitempty"`
	// example:
	//
	// alice
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// example:
	//
	// esr-4.3.0 (Spark 3.5.2, Scala 2.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// example:
	//
	// true
	EnablePublic *bool `json:"enablePublic,omitempty" xml:"enablePublic,omitempty"`
	// example:
	//
	// emr-spark-livy-gateway-cn-hangzhou.data.aliyun.com/api/v1/workspace/w-xxxxxxxxx/livycompute/lc-xxxxxxxxxxx
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// emr-spark-livy-gateway-cn-hangzhou-internal.aliyun.com/api/v1/workspace/w-xxxxxxxxx/livycompute/lc-xxxxxxxxxxx
	EndpointInner *string `json:"endpointInner,omitempty" xml:"endpointInner,omitempty"`
	// example:
	//
	// ev-cq31c7tlhtgm9nrrlj4g
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// example:
	//
	// false
	Fusion *bool `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// example:
	//
	// 1749456094000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// {
	//
	//   "sparkDefaultsConf": "spark.driver.cores     1\\nspark.driver.memory    4g\\nspark.executor.cores   1\\nspark.executor.memory  4g\\n",
	//
	//   "sparkBlackListConf": "spark.driver.cores\\nspark.driver.memory",
	//
	//   "livyConf": "livy.server.session.timeout  1h\\n",
	//
	//   "livyClientConf": "livy.rsc.sql.num-rows  1000\\n"
	//
	// }
	LivyServerConf *string `json:"livyServerConf,omitempty" xml:"livyServerConf,omitempty"`
	// example:
	//
	// 0.8.0
	LivyVersion *string `json:"livyVersion,omitempty" xml:"livyVersion,omitempty"`
	// example:
	//
	// 4Gi
	MemoryLimit *string `json:"memoryLimit,omitempty" xml:"memoryLimit,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// test
	NetworkName *string `json:"networkName,omitempty" xml:"networkName,omitempty"`
	// example:
	//
	// root_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// example:
	//
	// 10000001
	RamUserId *string `json:"ramUserId,omitempty" xml:"ramUserId,omitempty"`
	// example:
	//
	// esr-4.3.0 (Spark 3.5.2, Scala 2.12, Java Runtime)
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// example:
	//
	// 1749456094000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetLivyComputeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetLivyComputeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLivyComputeResponseBodyData) SetAuthType(v string) *GetLivyComputeResponseBodyData {
	s.AuthType = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetAutoStopConfiguration(v *GetLivyComputeResponseBodyDataAutoStopConfiguration) *GetLivyComputeResponseBodyData {
	s.AutoStopConfiguration = v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetComputeId(v string) *GetLivyComputeResponseBodyData {
	s.ComputeId = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetCpuLimit(v string) *GetLivyComputeResponseBodyData {
	s.CpuLimit = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetCreatedBy(v string) *GetLivyComputeResponseBodyData {
	s.CreatedBy = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetDisplayReleaseVersion(v string) *GetLivyComputeResponseBodyData {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetEnablePublic(v bool) *GetLivyComputeResponseBodyData {
	s.EnablePublic = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetEndpoint(v string) *GetLivyComputeResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetEndpointInner(v string) *GetLivyComputeResponseBodyData {
	s.EndpointInner = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetEnvironmentId(v string) *GetLivyComputeResponseBodyData {
	s.EnvironmentId = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetFusion(v bool) *GetLivyComputeResponseBodyData {
	s.Fusion = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetGmtCreate(v int64) *GetLivyComputeResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetLivyServerConf(v string) *GetLivyComputeResponseBodyData {
	s.LivyServerConf = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetLivyVersion(v string) *GetLivyComputeResponseBodyData {
	s.LivyVersion = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetMemoryLimit(v string) *GetLivyComputeResponseBodyData {
	s.MemoryLimit = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetName(v string) *GetLivyComputeResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetNetworkName(v string) *GetLivyComputeResponseBodyData {
	s.NetworkName = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetQueueName(v string) *GetLivyComputeResponseBodyData {
	s.QueueName = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetRamUserId(v string) *GetLivyComputeResponseBodyData {
	s.RamUserId = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetReleaseVersion(v string) *GetLivyComputeResponseBodyData {
	s.ReleaseVersion = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetStartTime(v int64) *GetLivyComputeResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetStatus(v string) *GetLivyComputeResponseBodyData {
	s.Status = &v
	return s
}

type GetLivyComputeResponseBodyDataAutoStopConfiguration struct {
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 300
	IdleTimeoutMinutes *int64 `json:"idleTimeoutMinutes,omitempty" xml:"idleTimeoutMinutes,omitempty"`
}

func (s GetLivyComputeResponseBodyDataAutoStopConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetLivyComputeResponseBodyDataAutoStopConfiguration) GoString() string {
	return s.String()
}

func (s *GetLivyComputeResponseBodyDataAutoStopConfiguration) SetEnable(v bool) *GetLivyComputeResponseBodyDataAutoStopConfiguration {
	s.Enable = &v
	return s
}

func (s *GetLivyComputeResponseBodyDataAutoStopConfiguration) SetIdleTimeoutMinutes(v int64) *GetLivyComputeResponseBodyDataAutoStopConfiguration {
	s.IdleTimeoutMinutes = &v
	return s
}

type GetLivyComputeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLivyComputeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLivyComputeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLivyComputeResponse) GoString() string {
	return s.String()
}

func (s *GetLivyComputeResponse) SetHeaders(v map[string]*string) *GetLivyComputeResponse {
	s.Headers = v
	return s
}

func (s *GetLivyComputeResponse) SetStatusCode(v int32) *GetLivyComputeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLivyComputeResponse) SetBody(v *GetLivyComputeResponseBody) *GetLivyComputeResponse {
	s.Body = v
	return s
}

type GetLivyComputeTokenRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetLivyComputeTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLivyComputeTokenRequest) GoString() string {
	return s.String()
}

func (s *GetLivyComputeTokenRequest) SetRegionId(v string) *GetLivyComputeTokenRequest {
	s.RegionId = &v
	return s
}

type GetLivyComputeTokenResponseBody struct {
	// example:
	//
	// 1000000
	Code *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetLivyComputeTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 484D9DDA-300D-525E-AF7A-0CCCA5C64A7A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetLivyComputeTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLivyComputeTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetLivyComputeTokenResponseBody) SetCode(v string) *GetLivyComputeTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetLivyComputeTokenResponseBody) SetData(v *GetLivyComputeTokenResponseBodyData) *GetLivyComputeTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetLivyComputeTokenResponseBody) SetMessage(v string) *GetLivyComputeTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetLivyComputeTokenResponseBody) SetRequestId(v string) *GetLivyComputeTokenResponseBody {
	s.RequestId = &v
	return s
}

type GetLivyComputeTokenResponseBodyData struct {
	AutoExpireConfiguration *GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration `json:"autoExpireConfiguration,omitempty" xml:"autoExpireConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// 1749456094000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// alice
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// example:
	//
	// 1749457994000
	ExpireTime *int64 `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// example:
	//
	// 1749456098000
	LastUsedTime *int64 `json:"lastUsedTime,omitempty" xml:"lastUsedTime,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// d25561157a635bb
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// Token ID。
	//
	// example:
	//
	// lctk-xxxxxxxxxx
	TokenId *string `json:"tokenId,omitempty" xml:"tokenId,omitempty"`
}

func (s GetLivyComputeTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetLivyComputeTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLivyComputeTokenResponseBodyData) SetAutoExpireConfiguration(v *GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration) *GetLivyComputeTokenResponseBodyData {
	s.AutoExpireConfiguration = v
	return s
}

func (s *GetLivyComputeTokenResponseBodyData) SetCreateTime(v int64) *GetLivyComputeTokenResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetLivyComputeTokenResponseBodyData) SetCreatedBy(v string) *GetLivyComputeTokenResponseBodyData {
	s.CreatedBy = &v
	return s
}

func (s *GetLivyComputeTokenResponseBodyData) SetExpireTime(v int64) *GetLivyComputeTokenResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetLivyComputeTokenResponseBodyData) SetLastUsedTime(v int64) *GetLivyComputeTokenResponseBodyData {
	s.LastUsedTime = &v
	return s
}

func (s *GetLivyComputeTokenResponseBodyData) SetName(v string) *GetLivyComputeTokenResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetLivyComputeTokenResponseBodyData) SetToken(v string) *GetLivyComputeTokenResponseBodyData {
	s.Token = &v
	return s
}

func (s *GetLivyComputeTokenResponseBodyData) SetTokenId(v string) *GetLivyComputeTokenResponseBodyData {
	s.TokenId = &v
	return s
}

type GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration struct {
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 7
	ExpireDays *int32 `json:"expireDays,omitempty" xml:"expireDays,omitempty"`
}

func (s GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration) GoString() string {
	return s.String()
}

func (s *GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration) SetEnable(v bool) *GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration {
	s.Enable = &v
	return s
}

func (s *GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration) SetExpireDays(v int32) *GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration {
	s.ExpireDays = &v
	return s
}

type GetLivyComputeTokenResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLivyComputeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLivyComputeTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLivyComputeTokenResponse) GoString() string {
	return s.String()
}

func (s *GetLivyComputeTokenResponse) SetHeaders(v map[string]*string) *GetLivyComputeTokenResponse {
	s.Headers = v
	return s
}

func (s *GetLivyComputeTokenResponse) SetStatusCode(v int32) *GetLivyComputeTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLivyComputeTokenResponse) SetBody(v *GetLivyComputeTokenResponseBody) *GetLivyComputeTokenResponse {
	s.Body = v
	return s
}

type GetSessionClusterRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetSessionClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSessionClusterRequest) GoString() string {
	return s.String()
}

func (s *GetSessionClusterRequest) SetRegionId(v string) *GetSessionClusterRequest {
	s.RegionId = &v
	return s
}

type GetSessionClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The session object.
	SessionCluster *GetSessionClusterResponseBodySessionCluster `json:"sessionCluster,omitempty" xml:"sessionCluster,omitempty" type:"Struct"`
}

func (s GetSessionClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSessionClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetSessionClusterResponseBody) SetRequestId(v string) *GetSessionClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSessionClusterResponseBody) SetSessionCluster(v *GetSessionClusterResponseBodySessionCluster) *GetSessionClusterResponseBody {
	s.SessionCluster = v
	return s
}

type GetSessionClusterResponseBodySessionCluster struct {
	// The Spark configurations.
	ApplicationConfigs []*GetSessionClusterResponseBodySessionClusterApplicationConfigs `json:"applicationConfigs,omitempty" xml:"applicationConfigs,omitempty" type:"Repeated"`
	// Indicates whether automatic startup is enabled.
	AutoStartConfiguration *GetSessionClusterResponseBodySessionClusterAutoStartConfiguration `json:"autoStartConfiguration,omitempty" xml:"autoStartConfiguration,omitempty" type:"Struct"`
	// Indicates whether automatic termination is enabled.
	AutoStopConfiguration *GetSessionClusterResponseBodySessionClusterAutoStopConfiguration `json:"autoStopConfiguration,omitempty" xml:"autoStopConfiguration,omitempty" type:"Struct"`
	ConnectionToken       *string                                                           `json:"connectionToken,omitempty" xml:"connectionToken,omitempty"`
	// The version of the Spark engine.
	//
	// example:
	//
	// esr-2.2(Java Runtime)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// The domain name to which the Spark UI of the session belongs.
	//
	// example:
	//
	// your.domain.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The internal endpoint.
	//
	// example:
	//
	// emr-spark-gateway-cn-hangzhou-internal.data.aliyuncs.com
	DomainInner *string `json:"domainInner,omitempty" xml:"domainInner,omitempty"`
	// The ID of the job that is associated with the session.
	//
	// example:
	//
	// TSK-xxxxxxxx
	DraftId *string `json:"draftId,omitempty" xml:"draftId,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-cpv569tlhtgndjl86t40
	EnvId *string `json:"envId,omitempty" xml:"envId,omitempty"`
	// The additional metadata of the session.
	//
	// example:
	//
	// {"extraInfoKey":"extraInfoValue"}
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// Indicates whether the Fusion engine is used for acceleration.
	//
	// example:
	//
	// false
	Fusion *bool `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-09-01 06:23:01
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The type of the job. This parameter is required and cannot be modified after the deployment is created. Valid values:
	//
	// 	- SQLSCRIPT
	//
	// 	- JAR
	//
	// 	- PYTHON
	//
	// example:
	//
	// SQL
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The name of the session.
	//
	// example:
	//
	// test
	Name                  *string `json:"name,omitempty" xml:"name,omitempty"`
	PublicEndpointEnabled *bool   `json:"publicEndpointEnabled,omitempty" xml:"publicEndpointEnabled,omitempty"`
	// The queue name.
	//
	// example:
	//
	// jobName
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// The version of Serverless Spark.
	//
	// example:
	//
	// esr-2.2(Java Runtime)
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	SessionClusterId *string `json:"sessionClusterId,omitempty" xml:"sessionClusterId,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2024-09-01 06:23:01
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The job status.
	//
	// 	- Starting
	//
	// 	- Running
	//
	// 	- Stopping
	//
	// 	- Stopped
	//
	// 	- Error
	//
	// example:
	//
	// Running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The reason of the job status change.
	StateChangeReason *GetSessionClusterResponseBodySessionClusterStateChangeReason `json:"stateChangeReason,omitempty" xml:"stateChangeReason,omitempty" type:"Struct"`
	// The user ID.
	//
	// example:
	//
	// jr-231231
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The name of the account that is used to create the session.
	//
	// example:
	//
	// user1
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// The Spark UI of the session.
	//
	// example:
	//
	// https://spark-ui/link
	WebUI *string `json:"webUI,omitempty" xml:"webUI,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// w-1234abcd
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetSessionClusterResponseBodySessionCluster) String() string {
	return tea.Prettify(s)
}

func (s GetSessionClusterResponseBodySessionCluster) GoString() string {
	return s.String()
}

func (s *GetSessionClusterResponseBodySessionCluster) SetApplicationConfigs(v []*GetSessionClusterResponseBodySessionClusterApplicationConfigs) *GetSessionClusterResponseBodySessionCluster {
	s.ApplicationConfigs = v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetAutoStartConfiguration(v *GetSessionClusterResponseBodySessionClusterAutoStartConfiguration) *GetSessionClusterResponseBodySessionCluster {
	s.AutoStartConfiguration = v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetAutoStopConfiguration(v *GetSessionClusterResponseBodySessionClusterAutoStopConfiguration) *GetSessionClusterResponseBodySessionCluster {
	s.AutoStopConfiguration = v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetConnectionToken(v string) *GetSessionClusterResponseBodySessionCluster {
	s.ConnectionToken = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetDisplayReleaseVersion(v string) *GetSessionClusterResponseBodySessionCluster {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetDomain(v string) *GetSessionClusterResponseBodySessionCluster {
	s.Domain = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetDomainInner(v string) *GetSessionClusterResponseBodySessionCluster {
	s.DomainInner = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetDraftId(v string) *GetSessionClusterResponseBodySessionCluster {
	s.DraftId = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetEnvId(v string) *GetSessionClusterResponseBodySessionCluster {
	s.EnvId = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetExtra(v string) *GetSessionClusterResponseBodySessionCluster {
	s.Extra = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetFusion(v bool) *GetSessionClusterResponseBodySessionCluster {
	s.Fusion = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetGmtCreate(v int64) *GetSessionClusterResponseBodySessionCluster {
	s.GmtCreate = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetKind(v string) *GetSessionClusterResponseBodySessionCluster {
	s.Kind = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetName(v string) *GetSessionClusterResponseBodySessionCluster {
	s.Name = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetPublicEndpointEnabled(v bool) *GetSessionClusterResponseBodySessionCluster {
	s.PublicEndpointEnabled = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetQueueName(v string) *GetSessionClusterResponseBodySessionCluster {
	s.QueueName = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetReleaseVersion(v string) *GetSessionClusterResponseBodySessionCluster {
	s.ReleaseVersion = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetSessionClusterId(v string) *GetSessionClusterResponseBodySessionCluster {
	s.SessionClusterId = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetStartTime(v int64) *GetSessionClusterResponseBodySessionCluster {
	s.StartTime = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetState(v string) *GetSessionClusterResponseBodySessionCluster {
	s.State = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetStateChangeReason(v *GetSessionClusterResponseBodySessionClusterStateChangeReason) *GetSessionClusterResponseBodySessionCluster {
	s.StateChangeReason = v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetUserId(v string) *GetSessionClusterResponseBodySessionCluster {
	s.UserId = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetUserName(v string) *GetSessionClusterResponseBodySessionCluster {
	s.UserName = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetWebUI(v string) *GetSessionClusterResponseBodySessionCluster {
	s.WebUI = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionCluster) SetWorkspaceId(v string) *GetSessionClusterResponseBodySessionCluster {
	s.WorkspaceId = &v
	return s
}

type GetSessionClusterResponseBodySessionClusterApplicationConfigs struct {
	// The name of the configuration file.
	//
	// example:
	//
	// spark-defaults.conf
	ConfigFileName *string `json:"configFileName,omitempty" xml:"configFileName,omitempty"`
	// The key of the configuration.
	//
	// example:
	//
	// spark.app.name
	ConfigItemKey *string `json:"configItemKey,omitempty" xml:"configItemKey,omitempty"`
	// The configuration value.
	//
	// example:
	//
	// test
	ConfigItemValue *string `json:"configItemValue,omitempty" xml:"configItemValue,omitempty"`
}

func (s GetSessionClusterResponseBodySessionClusterApplicationConfigs) String() string {
	return tea.Prettify(s)
}

func (s GetSessionClusterResponseBodySessionClusterApplicationConfigs) GoString() string {
	return s.String()
}

func (s *GetSessionClusterResponseBodySessionClusterApplicationConfigs) SetConfigFileName(v string) *GetSessionClusterResponseBodySessionClusterApplicationConfigs {
	s.ConfigFileName = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionClusterApplicationConfigs) SetConfigItemKey(v string) *GetSessionClusterResponseBodySessionClusterApplicationConfigs {
	s.ConfigItemKey = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionClusterApplicationConfigs) SetConfigItemValue(v string) *GetSessionClusterResponseBodySessionClusterApplicationConfigs {
	s.ConfigItemValue = &v
	return s
}

type GetSessionClusterResponseBodySessionClusterAutoStartConfiguration struct {
	// Indicates whether automatic startup is enabled.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s GetSessionClusterResponseBodySessionClusterAutoStartConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetSessionClusterResponseBodySessionClusterAutoStartConfiguration) GoString() string {
	return s.String()
}

func (s *GetSessionClusterResponseBodySessionClusterAutoStartConfiguration) SetEnable(v bool) *GetSessionClusterResponseBodySessionClusterAutoStartConfiguration {
	s.Enable = &v
	return s
}

type GetSessionClusterResponseBodySessionClusterAutoStopConfiguration struct {
	// Indicates whether automatic termination is enabled.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The idle timeout period. The session is automatically terminated when the idle timeout period is exceeded.
	//
	// example:
	//
	// 60
	IdleTimeoutMinutes *int32 `json:"idleTimeoutMinutes,omitempty" xml:"idleTimeoutMinutes,omitempty"`
}

func (s GetSessionClusterResponseBodySessionClusterAutoStopConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetSessionClusterResponseBodySessionClusterAutoStopConfiguration) GoString() string {
	return s.String()
}

func (s *GetSessionClusterResponseBodySessionClusterAutoStopConfiguration) SetEnable(v bool) *GetSessionClusterResponseBodySessionClusterAutoStopConfiguration {
	s.Enable = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionClusterAutoStopConfiguration) SetIdleTimeoutMinutes(v int32) *GetSessionClusterResponseBodySessionClusterAutoStopConfiguration {
	s.IdleTimeoutMinutes = &v
	return s
}

type GetSessionClusterResponseBodySessionClusterStateChangeReason struct {
	// The status change code.
	//
	// example:
	//
	// 1000000
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The status change message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetSessionClusterResponseBodySessionClusterStateChangeReason) String() string {
	return tea.Prettify(s)
}

func (s GetSessionClusterResponseBodySessionClusterStateChangeReason) GoString() string {
	return s.String()
}

func (s *GetSessionClusterResponseBodySessionClusterStateChangeReason) SetCode(v string) *GetSessionClusterResponseBodySessionClusterStateChangeReason {
	s.Code = &v
	return s
}

func (s *GetSessionClusterResponseBodySessionClusterStateChangeReason) SetMessage(v string) *GetSessionClusterResponseBodySessionClusterStateChangeReason {
	s.Message = &v
	return s
}

type GetSessionClusterResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSessionClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSessionClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSessionClusterResponse) GoString() string {
	return s.String()
}

func (s *GetSessionClusterResponse) SetHeaders(v map[string]*string) *GetSessionClusterResponse {
	s.Headers = v
	return s
}

func (s *GetSessionClusterResponse) SetStatusCode(v int32) *GetSessionClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSessionClusterResponse) SetBody(v *GetSessionClusterResponseBody) *GetSessionClusterResponse {
	s.Body = v
	return s
}

type GetSqlStatementRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetSqlStatementRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSqlStatementRequest) GoString() string {
	return s.String()
}

func (s *GetSqlStatementRequest) SetRegionId(v string) *GetSqlStatementRequest {
	s.RegionId = &v
	return s
}

type GetSqlStatementResponseBody struct {
	// The response parameters.
	Data *GetSqlStatementResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetSqlStatementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSqlStatementResponseBody) GoString() string {
	return s.String()
}

func (s *GetSqlStatementResponseBody) SetData(v *GetSqlStatementResponseBodyData) *GetSqlStatementResponseBody {
	s.Data = v
	return s
}

func (s *GetSqlStatementResponseBody) SetRequestId(v string) *GetSqlStatementResponseBody {
	s.RequestId = &v
	return s
}

type GetSqlStatementResponseBodyData struct {
	// The list of time that is consumed by SQL queries.
	ExecutionTime []*int64 `json:"executionTime,omitempty" xml:"executionTime,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// ERROR-102
	SqlErrorCode *string `json:"sqlErrorCode,omitempty" xml:"sqlErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// error message
	SqlErrorMessage *string `json:"sqlErrorMessage,omitempty" xml:"sqlErrorMessage,omitempty"`
	// The query results.
	SqlOutputs []*GetSqlStatementResponseBodyDataSqlOutputs `json:"sqlOutputs,omitempty" xml:"sqlOutputs,omitempty" type:"Repeated"`
	// The query status.
	//
	// Valid values:
	//
	// 	- running
	//
	// 	- available
	//
	// 	- cancelled
	//
	// 	- error
	//
	// 	- cancelling
	//
	// example:
	//
	// running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The query ID.
	//
	// example:
	//
	// st-1231311abadfaa
	StatementId *string `json:"statementId,omitempty" xml:"statementId,omitempty"`
}

func (s GetSqlStatementResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSqlStatementResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSqlStatementResponseBodyData) SetExecutionTime(v []*int64) *GetSqlStatementResponseBodyData {
	s.ExecutionTime = v
	return s
}

func (s *GetSqlStatementResponseBodyData) SetSqlErrorCode(v string) *GetSqlStatementResponseBodyData {
	s.SqlErrorCode = &v
	return s
}

func (s *GetSqlStatementResponseBodyData) SetSqlErrorMessage(v string) *GetSqlStatementResponseBodyData {
	s.SqlErrorMessage = &v
	return s
}

func (s *GetSqlStatementResponseBodyData) SetSqlOutputs(v []*GetSqlStatementResponseBodyDataSqlOutputs) *GetSqlStatementResponseBodyData {
	s.SqlOutputs = v
	return s
}

func (s *GetSqlStatementResponseBodyData) SetState(v string) *GetSqlStatementResponseBodyData {
	s.State = &v
	return s
}

func (s *GetSqlStatementResponseBodyData) SetStatementId(v string) *GetSqlStatementResponseBodyData {
	s.StatementId = &v
	return s
}

type GetSqlStatementResponseBodyDataSqlOutputs struct {
	// The queried data, which is a string in the JSON format.
	//
	// example:
	//
	// [{\\"values\\":[\\"test_db\\",\\"test_table\\",false]}
	Rows         *string `json:"rows,omitempty" xml:"rows,omitempty"`
	RowsFilePath *string `json:"rowsFilePath,omitempty" xml:"rowsFilePath,omitempty"`
	// The information about the schema, which is a string in the JSON format.
	//
	// example:
	//
	// {\\"type\\":\\"struct\\",\\"fields\\":[{\\"name\\":\\"namespace\\",\\"type\\":\\"string\\",\\"nullable\\":false,\\"metadata\\":{}},{\\"name\\":\\"tableName\\",\\"type\\":\\"string\\",\\"nullable\\":false,\\"metadata\\":{}},{\\"name\\":\\"isTemporary\\",\\"type\\":\\"boolean\\",\\"nullable\\":false,\\"metadata\\":{}}]}
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
}

func (s GetSqlStatementResponseBodyDataSqlOutputs) String() string {
	return tea.Prettify(s)
}

func (s GetSqlStatementResponseBodyDataSqlOutputs) GoString() string {
	return s.String()
}

func (s *GetSqlStatementResponseBodyDataSqlOutputs) SetRows(v string) *GetSqlStatementResponseBodyDataSqlOutputs {
	s.Rows = &v
	return s
}

func (s *GetSqlStatementResponseBodyDataSqlOutputs) SetRowsFilePath(v string) *GetSqlStatementResponseBodyDataSqlOutputs {
	s.RowsFilePath = &v
	return s
}

func (s *GetSqlStatementResponseBodyDataSqlOutputs) SetSchema(v string) *GetSqlStatementResponseBodyDataSqlOutputs {
	s.Schema = &v
	return s
}

type GetSqlStatementResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSqlStatementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSqlStatementResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSqlStatementResponse) GoString() string {
	return s.String()
}

func (s *GetSqlStatementResponse) SetHeaders(v map[string]*string) *GetSqlStatementResponse {
	s.Headers = v
	return s
}

func (s *GetSqlStatementResponse) SetStatusCode(v int32) *GetSqlStatementResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSqlStatementResponse) SetBody(v *GetSqlStatementResponseBody) *GetSqlStatementResponse {
	s.Body = v
	return s
}

type GetTemplateRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId      *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	TemplateBizId *string `json:"templateBizId,omitempty" xml:"templateBizId,omitempty"`
	// The template type.
	//
	// Valid values:
	//
	// 	- TASK
	//
	// 	- SESSION
	//
	// example:
	//
	// TASK
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) SetRegionId(v string) *GetTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateBizId(v string) *GetTemplateRequest {
	s.TemplateBizId = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateType(v string) *GetTemplateRequest {
	s.TemplateType = &v
	return s
}

type GetTemplateResponseBody struct {
	// The returned data.
	Data *Template `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// 040003
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// InvalidUser.NotFound
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 484D9DDA-300D-525E-AF7A-0CCCA5C64A7A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) SetData(v *Template) *GetTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetTemplateResponseBody) SetErrorCode(v string) *GetTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTemplateResponseBody) SetErrorMessage(v string) *GetTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTemplateResponseBody) SetHttpStatusCode(v string) *GetTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetSuccess(v bool) *GetTemplateResponseBody {
	s.Success = &v
	return s
}

type GetTemplateResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetTemplateResponse) SetBody(v *GetTemplateResponseBody) *GetTemplateResponse {
	s.Body = v
	return s
}

type GrantRoleToUsersRequest struct {
	// The Alibaba Cloud Resource Name (ARN) of the RAM role.
	//
	// example:
	//
	// acs:emr::w-975bcfda9625****:role/Owner
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// The user ARNs.
	UserArns []*string `json:"userArns,omitempty" xml:"userArns,omitempty" type:"Repeated"`
	// The region ID.
	//
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
	// The request ID.
	//
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
	// The ID of the user who created the job.
	//
	// example:
	//
	// 1509789347011222
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The range of end time.
	EndTime    *ListJobRunsRequestEndTime `json:"endTime,omitempty" xml:"endTime,omitempty" type:"Struct"`
	IsWorkflow *string                    `json:"isWorkflow,omitempty" xml:"isWorkflow,omitempty"`
	// The job run ID.
	//
	// example:
	//
	// jd-b6d003f1930f****
	JobRunDeploymentId *string `json:"jobRunDeploymentId,omitempty" xml:"jobRunDeploymentId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// j-xxx
	JobRunId *string `json:"jobRunId,omitempty" xml:"jobRunId,omitempty"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The minimum running duration of the job. Unit: ms.
	//
	// example:
	//
	// 60000
	MinDuration *int64 `json:"minDuration,omitempty" xml:"minDuration,omitempty"`
	// The job name.
	//
	// example:
	//
	// emr-spark-demo-job
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The pagination token that is used in the request to retrieve a new page of results.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The name of the resource queue on which the Spark jobs run.
	//
	// example:
	//
	// dev_queue
	ResourceQueueId *string `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	// The range of start time.
	StartTime *ListJobRunsRequestStartTime `json:"startTime,omitempty" xml:"startTime,omitempty" type:"Struct"`
	// The job states.
	//
	// example:
	//
	// ["Running","Submitted"]
	States []*string `json:"states,omitempty" xml:"states,omitempty" type:"Repeated"`
	// The tags of the job.
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

func (s *ListJobRunsRequest) SetIsWorkflow(v string) *ListJobRunsRequest {
	s.IsWorkflow = &v
	return s
}

func (s *ListJobRunsRequest) SetJobRunDeploymentId(v string) *ListJobRunsRequest {
	s.JobRunDeploymentId = &v
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

func (s *ListJobRunsRequest) SetMinDuration(v int64) *ListJobRunsRequest {
	s.MinDuration = &v
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
	// The end of the end time range.
	//
	// example:
	//
	// 1710432000000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The beginning of the end time range.
	//
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
	// The end of the start time range.
	//
	// example:
	//
	// 1710432000000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The beginning of the start time range.
	//
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
	// The key of tag N.
	//
	// example:
	//
	// tag_key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of tag N.
	//
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
	// The ID of the user who created the job.
	//
	// example:
	//
	// 1509789347011222
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The range of end time.
	EndTimeShrink *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	IsWorkflow    *string `json:"isWorkflow,omitempty" xml:"isWorkflow,omitempty"`
	// The job run ID.
	//
	// example:
	//
	// jd-b6d003f1930f****
	JobRunDeploymentId *string `json:"jobRunDeploymentId,omitempty" xml:"jobRunDeploymentId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// j-xxx
	JobRunId *string `json:"jobRunId,omitempty" xml:"jobRunId,omitempty"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The minimum running duration of the job. Unit: ms.
	//
	// example:
	//
	// 60000
	MinDuration *int64 `json:"minDuration,omitempty" xml:"minDuration,omitempty"`
	// The job name.
	//
	// example:
	//
	// emr-spark-demo-job
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The pagination token that is used in the request to retrieve a new page of results.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The name of the resource queue on which the Spark jobs run.
	//
	// example:
	//
	// dev_queue
	ResourceQueueId *string `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	// The range of start time.
	StartTimeShrink *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The job states.
	//
	// example:
	//
	// ["Running","Submitted"]
	StatesShrink *string `json:"states,omitempty" xml:"states,omitempty"`
	// The tags of the job.
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

func (s *ListJobRunsShrinkRequest) SetIsWorkflow(v string) *ListJobRunsShrinkRequest {
	s.IsWorkflow = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetJobRunDeploymentId(v string) *ListJobRunsShrinkRequest {
	s.JobRunDeploymentId = &v
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

func (s *ListJobRunsShrinkRequest) SetMinDuration(v int64) *ListJobRunsShrinkRequest {
	s.MinDuration = &v
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
	// The Spark jobs.
	JobRuns []*ListJobRunsResponseBodyJobRuns `json:"jobRuns,omitempty" xml:"jobRuns,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries returned.
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
	// The code type of the job. Valid values:
	//
	// SQL
	//
	// JAR
	//
	// PYTHON
	//
	// example:
	//
	// SQL
	CodeType *string `json:"codeType,omitempty" xml:"codeType,omitempty"`
	// The advanced configurations of Spark.
	ConfigurationOverrides *ListJobRunsResponseBodyJobRunsConfigurationOverrides `json:"configurationOverrides,omitempty" xml:"configurationOverrides,omitempty" type:"Struct"`
	// The ID of the user who created the job.
	//
	// example:
	//
	// 1509789347011222
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The number of CUs consumed during a specified cycle of a task. The value is an estimated value. Refer to your Alibaba Cloud bill for the actual number of consumed CUs.
	//
	// example:
	//
	// 2.059
	CuHours *float64 `json:"cuHours,omitempty" xml:"cuHours,omitempty"`
	// The version of Spark on which the jobs run.
	//
	// example:
	//
	// esr-3.0.0 (Spark 3.4.3, Scala 2.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// The end time of the job.
	//
	// example:
	//
	// 1684119314000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The timeout period of the job.
	//
	// example:
	//
	// 3600
	ExecutionTimeoutSeconds *int32 `json:"executionTimeoutSeconds,omitempty" xml:"executionTimeoutSeconds,omitempty"`
	// Indicates whether the Fusion engine is used for acceleration.
	//
	// example:
	//
	// true
	Fusion *bool `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// The information about Spark Driver.
	JobDriver *JobDriver `json:"jobDriver,omitempty" xml:"jobDriver,omitempty"`
	// The job ID.
	//
	// example:
	//
	// jr-231231
	JobRunId *string `json:"jobRunId,omitempty" xml:"jobRunId,omitempty"`
	// The path where the operational logs are stored.
	Log *RunLog `json:"log,omitempty" xml:"log,omitempty"`
	// The total amount of memory allocated to the job multiplied by the running duration (seconds).
	//
	// example:
	//
	// 33030784
	MbSeconds *int64 `json:"mbSeconds,omitempty" xml:"mbSeconds,omitempty"`
	// The job name.
	//
	// example:
	//
	// jobName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The version of Spark on which the jobs run.
	//
	// example:
	//
	// esr-native-3.4.0
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// The job state.
	//
	// example:
	//
	// Running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The reason of the job status change.
	StateChangeReason *ListJobRunsResponseBodyJobRunsStateChangeReason `json:"stateChangeReason,omitempty" xml:"stateChangeReason,omitempty" type:"Struct"`
	// The time when the job was submitted.
	//
	// example:
	//
	// 1684119314000
	SubmitTime *int64 `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
	// The tags of the job.
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The total number of CPU cores allocated to the job multiplied by the running duration (seconds).
	//
	// example:
	//
	// 8236
	VcoreSeconds *int64 `json:"vcoreSeconds,omitempty" xml:"vcoreSeconds,omitempty"`
	// The web UI of the job.
	//
	// example:
	//
	// http://spark-ui
	WebUI *string `json:"webUI,omitempty" xml:"webUI,omitempty"`
	// The workspace ID.
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

func (s *ListJobRunsResponseBodyJobRuns) SetCuHours(v float64) *ListJobRunsResponseBodyJobRuns {
	s.CuHours = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetDisplayReleaseVersion(v string) *ListJobRunsResponseBodyJobRuns {
	s.DisplayReleaseVersion = &v
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

func (s *ListJobRunsResponseBodyJobRuns) SetFusion(v bool) *ListJobRunsResponseBodyJobRuns {
	s.Fusion = &v
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

func (s *ListJobRunsResponseBodyJobRuns) SetMbSeconds(v int64) *ListJobRunsResponseBodyJobRuns {
	s.MbSeconds = &v
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

func (s *ListJobRunsResponseBodyJobRuns) SetVcoreSeconds(v int64) *ListJobRunsResponseBodyJobRuns {
	s.VcoreSeconds = &v
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
	// The SparkConf objects.
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
	// The error code.
	//
	// example:
	//
	// 0
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error message.
	//
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

type ListKyuubiServicesResponseBody struct {
	Data *ListKyuubiServicesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListKyuubiServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListKyuubiServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListKyuubiServicesResponseBody) SetData(v *ListKyuubiServicesResponseBodyData) *ListKyuubiServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListKyuubiServicesResponseBody) SetRequestId(v string) *ListKyuubiServicesResponseBody {
	s.RequestId = &v
	return s
}

type ListKyuubiServicesResponseBodyData struct {
	KyuubiServices []*ListKyuubiServicesResponseBodyDataKyuubiServices `json:"kyuubiServices,omitempty" xml:"kyuubiServices,omitempty" type:"Repeated"`
}

func (s ListKyuubiServicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListKyuubiServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListKyuubiServicesResponseBodyData) SetKyuubiServices(v []*ListKyuubiServicesResponseBodyDataKyuubiServices) *ListKyuubiServicesResponseBodyData {
	s.KyuubiServices = v
	return s
}

type ListKyuubiServicesResponseBodyDataKyuubiServices struct {
	// example:
	//
	// 4C16G
	ComputeInstance *string `json:"computeInstance,omitempty" xml:"computeInstance,omitempty"`
	// example:
	//
	// 2025-03-11T08:21:58Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 103*******
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// kyuubi-cn-hangzhou-internal.spark.emr.aliyuncs.com
	InnerEndpoint *string `json:"innerEndpoint,omitempty" xml:"innerEndpoint,omitempty"`
	// example:
	//
	// kyuubi.conf.key=value1
	//
	// kyuubi.conf.key1=value2
	KyuubiConfigs *string `json:"kyuubiConfigs,omitempty" xml:"kyuubiConfigs,omitempty"`
	// KyuubiServer ID。
	//
	// example:
	//
	// kb-070104e7631242448d12a1377c309f30
	KyuubiServiceId *string `json:"kyuubiServiceId,omitempty" xml:"kyuubiServiceId,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// emr-spark-kyuubi-gateway-cn-hangzhou.aliyuncs.com
	PublicEndpoint *string `json:"publicEndpoint,omitempty" xml:"publicEndpoint,omitempty"`
	// example:
	//
	// dev_queue
	Queue *string `json:"queue,omitempty" xml:"queue,omitempty"`
	// example:
	//
	// esr-4.2.0 (Spark 3.5.2, Scala 2.12)
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// example:
	//
	// 3
	Replica *int32 `json:"replica,omitempty" xml:"replica,omitempty"`
	// example:
	//
	// spark.conf.key=value1
	//
	// spark.conf.key1=value2
	SparkConfigs *string `json:"sparkConfigs,omitempty" xml:"sparkConfigs,omitempty"`
	// example:
	//
	// 2024-11-23 09:22:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// Running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s ListKyuubiServicesResponseBodyDataKyuubiServices) String() string {
	return tea.Prettify(s)
}

func (s ListKyuubiServicesResponseBodyDataKyuubiServices) GoString() string {
	return s.String()
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetComputeInstance(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.ComputeInstance = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetCreateTime(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.CreateTime = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetCreator(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.Creator = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetInnerEndpoint(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.InnerEndpoint = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetKyuubiConfigs(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.KyuubiConfigs = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetKyuubiServiceId(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.KyuubiServiceId = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetName(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.Name = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetPublicEndpoint(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.PublicEndpoint = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetQueue(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.Queue = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetReleaseVersion(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.ReleaseVersion = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetReplica(v int32) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.Replica = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetSparkConfigs(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.SparkConfigs = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetStartTime(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.StartTime = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetState(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.State = &v
	return s
}

type ListKyuubiServicesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKyuubiServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKyuubiServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListKyuubiServicesResponse) GoString() string {
	return s.String()
}

func (s *ListKyuubiServicesResponse) SetHeaders(v map[string]*string) *ListKyuubiServicesResponse {
	s.Headers = v
	return s
}

func (s *ListKyuubiServicesResponse) SetStatusCode(v int32) *ListKyuubiServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKyuubiServicesResponse) SetBody(v *ListKyuubiServicesResponseBody) *ListKyuubiServicesResponse {
	s.Body = v
	return s
}

type ListKyuubiSparkApplicationsRequest struct {
	// The ID of the application that is submitted by using a Kyuubi gateway.
	//
	// example:
	//
	// spark-339f844005b6404c95f9f7c7a13b****
	ApplicationId *string `json:"applicationId,omitempty" xml:"applicationId,omitempty"`
	// The name of the Spark application that is submitted by using a Kyuubi gateway.
	//
	// example:
	//
	// kyuubi-connection-spark-sql-anonymous-fa9a5e73-b4b1-474a-b****
	ApplicationName *string `json:"applicationName,omitempty" xml:"applicationName,omitempty"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 20
	MaxResults  *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	MinDuration *int64 `json:"minDuration,omitempty" xml:"minDuration,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 1
	NextToken       *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OrderBy         []*string `json:"orderBy,omitempty" xml:"orderBy,omitempty" type:"Repeated"`
	ResourceQueueId *string   `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	Sort            *string   `json:"sort,omitempty" xml:"sort,omitempty"`
	// The range of start time.
	StartTime *ListKyuubiSparkApplicationsRequestStartTime `json:"startTime,omitempty" xml:"startTime,omitempty" type:"Struct"`
}

func (s ListKyuubiSparkApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListKyuubiSparkApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListKyuubiSparkApplicationsRequest) SetApplicationId(v string) *ListKyuubiSparkApplicationsRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListKyuubiSparkApplicationsRequest) SetApplicationName(v string) *ListKyuubiSparkApplicationsRequest {
	s.ApplicationName = &v
	return s
}

func (s *ListKyuubiSparkApplicationsRequest) SetMaxResults(v int32) *ListKyuubiSparkApplicationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListKyuubiSparkApplicationsRequest) SetMinDuration(v int64) *ListKyuubiSparkApplicationsRequest {
	s.MinDuration = &v
	return s
}

func (s *ListKyuubiSparkApplicationsRequest) SetNextToken(v string) *ListKyuubiSparkApplicationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListKyuubiSparkApplicationsRequest) SetOrderBy(v []*string) *ListKyuubiSparkApplicationsRequest {
	s.OrderBy = v
	return s
}

func (s *ListKyuubiSparkApplicationsRequest) SetResourceQueueId(v string) *ListKyuubiSparkApplicationsRequest {
	s.ResourceQueueId = &v
	return s
}

func (s *ListKyuubiSparkApplicationsRequest) SetSort(v string) *ListKyuubiSparkApplicationsRequest {
	s.Sort = &v
	return s
}

func (s *ListKyuubiSparkApplicationsRequest) SetStartTime(v *ListKyuubiSparkApplicationsRequestStartTime) *ListKyuubiSparkApplicationsRequest {
	s.StartTime = v
	return s
}

type ListKyuubiSparkApplicationsRequestStartTime struct {
	// The end of the start time range.
	//
	// example:
	//
	// 1710432000000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The beginning of the start time range.
	//
	// example:
	//
	// 1709740800000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListKyuubiSparkApplicationsRequestStartTime) String() string {
	return tea.Prettify(s)
}

func (s ListKyuubiSparkApplicationsRequestStartTime) GoString() string {
	return s.String()
}

func (s *ListKyuubiSparkApplicationsRequestStartTime) SetEndTime(v int64) *ListKyuubiSparkApplicationsRequestStartTime {
	s.EndTime = &v
	return s
}

func (s *ListKyuubiSparkApplicationsRequestStartTime) SetStartTime(v int64) *ListKyuubiSparkApplicationsRequestStartTime {
	s.StartTime = &v
	return s
}

type ListKyuubiSparkApplicationsShrinkRequest struct {
	// The ID of the application that is submitted by using a Kyuubi gateway.
	//
	// example:
	//
	// spark-339f844005b6404c95f9f7c7a13b****
	ApplicationId *string `json:"applicationId,omitempty" xml:"applicationId,omitempty"`
	// The name of the Spark application that is submitted by using a Kyuubi gateway.
	//
	// example:
	//
	// kyuubi-connection-spark-sql-anonymous-fa9a5e73-b4b1-474a-b****
	ApplicationName *string `json:"applicationName,omitempty" xml:"applicationName,omitempty"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 20
	MaxResults  *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	MinDuration *int64 `json:"minDuration,omitempty" xml:"minDuration,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 1
	NextToken       *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OrderByShrink   *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	ResourceQueueId *string `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	Sort            *string `json:"sort,omitempty" xml:"sort,omitempty"`
	// The range of start time.
	StartTimeShrink *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListKyuubiSparkApplicationsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListKyuubiSparkApplicationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetApplicationId(v string) *ListKyuubiSparkApplicationsShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetApplicationName(v string) *ListKyuubiSparkApplicationsShrinkRequest {
	s.ApplicationName = &v
	return s
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetMaxResults(v int32) *ListKyuubiSparkApplicationsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetMinDuration(v int64) *ListKyuubiSparkApplicationsShrinkRequest {
	s.MinDuration = &v
	return s
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetNextToken(v string) *ListKyuubiSparkApplicationsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetOrderByShrink(v string) *ListKyuubiSparkApplicationsShrinkRequest {
	s.OrderByShrink = &v
	return s
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetResourceQueueId(v string) *ListKyuubiSparkApplicationsShrinkRequest {
	s.ResourceQueueId = &v
	return s
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetSort(v string) *ListKyuubiSparkApplicationsShrinkRequest {
	s.Sort = &v
	return s
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetStartTimeShrink(v string) *ListKyuubiSparkApplicationsShrinkRequest {
	s.StartTimeShrink = &v
	return s
}

type ListKyuubiSparkApplicationsResponseBody struct {
	// The details of the applications.
	Applications []*ListKyuubiSparkApplicationsResponseBodyApplications `json:"applications,omitempty" xml:"applications,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListKyuubiSparkApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListKyuubiSparkApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKyuubiSparkApplicationsResponseBody) SetApplications(v []*ListKyuubiSparkApplicationsResponseBodyApplications) *ListKyuubiSparkApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBody) SetMaxResults(v int32) *ListKyuubiSparkApplicationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBody) SetNextToken(v string) *ListKyuubiSparkApplicationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBody) SetRequestId(v string) *ListKyuubiSparkApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBody) SetTotalCount(v int32) *ListKyuubiSparkApplicationsResponseBody {
	s.TotalCount = &v
	return s
}

type ListKyuubiSparkApplicationsResponseBodyApplications struct {
	// The ID of the application that is submitted by using a Kyuubi gateway.
	//
	// example:
	//
	// spark-339f844005b6404c95f9f7c7a13b****
	ApplicationId *string `json:"applicationId,omitempty" xml:"applicationId,omitempty"`
	// The name of the Spark application that is submitted by using a Kyuubi gateway.
	//
	// example:
	//
	// kyuubi-connection-spark-sql-anonymous-fa9a5e73-b4b1-474a-b****
	ApplicationName *string `json:"applicationName,omitempty" xml:"applicationName,omitempty"`
	// The number of CUs consumed during a specified cycle of a task. The value is an estimated value. Refer to your Alibaba Cloud bill for the actual number of consumed CUs.
	//
	// example:
	//
	// 0.238302
	CuHours *float64 `json:"cuHours,omitempty" xml:"cuHours,omitempty"`
	// The time when the task ended.
	//
	// example:
	//
	// 2025-02-12 20:02:02
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The total amount of memory allocated to the job multiplied by the running duration (seconds).
	//
	// example:
	//
	// 3513900
	MbSeconds *int64 `json:"mbSeconds,omitempty" xml:"mbSeconds,omitempty"`
	// The name of the resource queue on which the Spark jobs run.
	//
	// example:
	//
	// dev_queue
	ResourceQueueId *string `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	// The time when the task started.
	//
	// example:
	//
	// 2025-02-12 19:59:16
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The status of the Spark application.
	//
	// 	- STARTING
	//
	// 	- RUNNING
	//
	// 	- TERMINATED
	//
	// example:
	//
	// STARTING
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The total number of CPU cores allocated to the job multiplied by the running duration (seconds).
	//
	// example:
	//
	// 780
	VcoreSeconds *int64 `json:"vcoreSeconds,omitempty" xml:"vcoreSeconds,omitempty"`
	// The URL of the web UI for the Spark application.
	WebUI *string `json:"webUI,omitempty" xml:"webUI,omitempty"`
}

func (s ListKyuubiSparkApplicationsResponseBodyApplications) String() string {
	return tea.Prettify(s)
}

func (s ListKyuubiSparkApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetApplicationId(v string) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetApplicationName(v string) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.ApplicationName = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetCuHours(v float64) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.CuHours = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetEndTime(v string) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.EndTime = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetMbSeconds(v int64) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.MbSeconds = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetResourceQueueId(v string) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.ResourceQueueId = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetStartTime(v string) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.StartTime = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetState(v string) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.State = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetVcoreSeconds(v int64) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.VcoreSeconds = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponseBodyApplications) SetWebUI(v string) *ListKyuubiSparkApplicationsResponseBodyApplications {
	s.WebUI = &v
	return s
}

type ListKyuubiSparkApplicationsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKyuubiSparkApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKyuubiSparkApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListKyuubiSparkApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListKyuubiSparkApplicationsResponse) SetHeaders(v map[string]*string) *ListKyuubiSparkApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListKyuubiSparkApplicationsResponse) SetStatusCode(v int32) *ListKyuubiSparkApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKyuubiSparkApplicationsResponse) SetBody(v *ListKyuubiSparkApplicationsResponseBody) *ListKyuubiSparkApplicationsResponse {
	s.Body = v
	return s
}

type ListKyuubiTokenRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListKyuubiTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s ListKyuubiTokenRequest) GoString() string {
	return s.String()
}

func (s *ListKyuubiTokenRequest) SetRegionId(v string) *ListKyuubiTokenRequest {
	s.RegionId = &v
	return s
}

type ListKyuubiTokenResponseBody struct {
	Data *ListKyuubiTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListKyuubiTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListKyuubiTokenResponseBody) GoString() string {
	return s.String()
}

func (s *ListKyuubiTokenResponseBody) SetData(v *ListKyuubiTokenResponseBodyData) *ListKyuubiTokenResponseBody {
	s.Data = v
	return s
}

func (s *ListKyuubiTokenResponseBody) SetRequestId(v string) *ListKyuubiTokenResponseBody {
	s.RequestId = &v
	return s
}

type ListKyuubiTokenResponseBodyData struct {
	Tokens []*ListKyuubiTokenResponseBodyDataTokens `json:"tokens,omitempty" xml:"tokens,omitempty" type:"Repeated"`
}

func (s ListKyuubiTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListKyuubiTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListKyuubiTokenResponseBodyData) SetTokens(v []*ListKyuubiTokenResponseBodyDataTokens) *ListKyuubiTokenResponseBodyData {
	s.Tokens = v
	return s
}

type ListKyuubiTokenResponseBodyDataTokens struct {
	// example:
	//
	// 2025-02-11T02:23:02Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// test_user
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// example:
	//
	// 1740366769165
	ExpireTime *int64 `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// example:
	//
	// 1740366232121
	LastUsedTime *int64 `json:"lastUsedTime,omitempty" xml:"lastUsedTime,omitempty"`
	// example:
	//
	// dev_serveless_spark
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// f14c1347-dcfd-4082-b101-77aa96b5de36
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// Token ID。
	//
	// example:
	//
	// f14c1347-dcfd-4082-b101-77aa96b5de36
	TokenId *string `json:"tokenId,omitempty" xml:"tokenId,omitempty"`
}

func (s ListKyuubiTokenResponseBodyDataTokens) String() string {
	return tea.Prettify(s)
}

func (s ListKyuubiTokenResponseBodyDataTokens) GoString() string {
	return s.String()
}

func (s *ListKyuubiTokenResponseBodyDataTokens) SetCreateTime(v int64) *ListKyuubiTokenResponseBodyDataTokens {
	s.CreateTime = &v
	return s
}

func (s *ListKyuubiTokenResponseBodyDataTokens) SetCreatedBy(v string) *ListKyuubiTokenResponseBodyDataTokens {
	s.CreatedBy = &v
	return s
}

func (s *ListKyuubiTokenResponseBodyDataTokens) SetExpireTime(v int64) *ListKyuubiTokenResponseBodyDataTokens {
	s.ExpireTime = &v
	return s
}

func (s *ListKyuubiTokenResponseBodyDataTokens) SetLastUsedTime(v int64) *ListKyuubiTokenResponseBodyDataTokens {
	s.LastUsedTime = &v
	return s
}

func (s *ListKyuubiTokenResponseBodyDataTokens) SetName(v string) *ListKyuubiTokenResponseBodyDataTokens {
	s.Name = &v
	return s
}

func (s *ListKyuubiTokenResponseBodyDataTokens) SetToken(v string) *ListKyuubiTokenResponseBodyDataTokens {
	s.Token = &v
	return s
}

func (s *ListKyuubiTokenResponseBodyDataTokens) SetTokenId(v string) *ListKyuubiTokenResponseBodyDataTokens {
	s.TokenId = &v
	return s
}

type ListKyuubiTokenResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKyuubiTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKyuubiTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s ListKyuubiTokenResponse) GoString() string {
	return s.String()
}

func (s *ListKyuubiTokenResponse) SetHeaders(v map[string]*string) *ListKyuubiTokenResponse {
	s.Headers = v
	return s
}

func (s *ListKyuubiTokenResponse) SetStatusCode(v int32) *ListKyuubiTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKyuubiTokenResponse) SetBody(v *ListKyuubiTokenResponseBody) *ListKyuubiTokenResponse {
	s.Body = v
	return s
}

type ListLivyComputeRequest struct {
	// example:
	//
	// ev-cq31c7tlhtgm9nrrlj4g
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListLivyComputeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLivyComputeRequest) GoString() string {
	return s.String()
}

func (s *ListLivyComputeRequest) SetEnvironmentId(v string) *ListLivyComputeRequest {
	s.EnvironmentId = &v
	return s
}

func (s *ListLivyComputeRequest) SetRegionId(v string) *ListLivyComputeRequest {
	s.RegionId = &v
	return s
}

type ListLivyComputeResponseBody struct {
	// example:
	//
	// 1000000
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListLivyComputeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListLivyComputeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLivyComputeResponseBody) GoString() string {
	return s.String()
}

func (s *ListLivyComputeResponseBody) SetCode(v string) *ListLivyComputeResponseBody {
	s.Code = &v
	return s
}

func (s *ListLivyComputeResponseBody) SetData(v *ListLivyComputeResponseBodyData) *ListLivyComputeResponseBody {
	s.Data = v
	return s
}

func (s *ListLivyComputeResponseBody) SetMessage(v string) *ListLivyComputeResponseBody {
	s.Message = &v
	return s
}

func (s *ListLivyComputeResponseBody) SetRequestId(v string) *ListLivyComputeResponseBody {
	s.RequestId = &v
	return s
}

type ListLivyComputeResponseBodyData struct {
	LivyComputes []*ListLivyComputeResponseBodyDataLivyComputes `json:"livyComputes,omitempty" xml:"livyComputes,omitempty" type:"Repeated"`
}

func (s ListLivyComputeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListLivyComputeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLivyComputeResponseBodyData) SetLivyComputes(v []*ListLivyComputeResponseBodyDataLivyComputes) *ListLivyComputeResponseBodyData {
	s.LivyComputes = v
	return s
}

type ListLivyComputeResponseBodyDataLivyComputes struct {
	// example:
	//
	// lc-xxxxxxxxxxxx
	ComputeId *string `json:"computeId,omitempty" xml:"computeId,omitempty"`
	// example:
	//
	// alice
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// example:
	//
	// emr-spark-livy-gateway-cn-hangzhou.data.aliyun.com/api/v1/workspace/w-xxxxxxxxx/livycompute/lc-xxxxxxxxxxx
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// emr-spark-livy-gateway-cn-hangzhou-internal.aliyun.com/api/v1/workspace/w-xxxxxxxxx/livycompute/lc-xxxxxxxxxxx
	EndpointInner *string `json:"endpointInner,omitempty" xml:"endpointInner,omitempty"`
	// example:
	//
	// 1749456094000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// root_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// example:
	//
	// 1749456094000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListLivyComputeResponseBodyDataLivyComputes) String() string {
	return tea.Prettify(s)
}

func (s ListLivyComputeResponseBodyDataLivyComputes) GoString() string {
	return s.String()
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) SetComputeId(v string) *ListLivyComputeResponseBodyDataLivyComputes {
	s.ComputeId = &v
	return s
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) SetCreatedBy(v string) *ListLivyComputeResponseBodyDataLivyComputes {
	s.CreatedBy = &v
	return s
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) SetEndpoint(v string) *ListLivyComputeResponseBodyDataLivyComputes {
	s.Endpoint = &v
	return s
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) SetEndpointInner(v string) *ListLivyComputeResponseBodyDataLivyComputes {
	s.EndpointInner = &v
	return s
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) SetGmtCreate(v int64) *ListLivyComputeResponseBodyDataLivyComputes {
	s.GmtCreate = &v
	return s
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) SetName(v string) *ListLivyComputeResponseBodyDataLivyComputes {
	s.Name = &v
	return s
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) SetQueueName(v string) *ListLivyComputeResponseBodyDataLivyComputes {
	s.QueueName = &v
	return s
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) SetStartTime(v int64) *ListLivyComputeResponseBodyDataLivyComputes {
	s.StartTime = &v
	return s
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) SetStatus(v string) *ListLivyComputeResponseBodyDataLivyComputes {
	s.Status = &v
	return s
}

type ListLivyComputeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLivyComputeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLivyComputeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLivyComputeResponse) GoString() string {
	return s.String()
}

func (s *ListLivyComputeResponse) SetHeaders(v map[string]*string) *ListLivyComputeResponse {
	s.Headers = v
	return s
}

func (s *ListLivyComputeResponse) SetStatusCode(v int32) *ListLivyComputeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLivyComputeResponse) SetBody(v *ListLivyComputeResponseBody) *ListLivyComputeResponse {
	s.Body = v
	return s
}

type ListLivyComputeTokenRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListLivyComputeTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLivyComputeTokenRequest) GoString() string {
	return s.String()
}

func (s *ListLivyComputeTokenRequest) SetRegionId(v string) *ListLivyComputeTokenRequest {
	s.RegionId = &v
	return s
}

type ListLivyComputeTokenResponseBody struct {
	// example:
	//
	// 1000000
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListLivyComputeTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListLivyComputeTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLivyComputeTokenResponseBody) GoString() string {
	return s.String()
}

func (s *ListLivyComputeTokenResponseBody) SetCode(v string) *ListLivyComputeTokenResponseBody {
	s.Code = &v
	return s
}

func (s *ListLivyComputeTokenResponseBody) SetData(v *ListLivyComputeTokenResponseBodyData) *ListLivyComputeTokenResponseBody {
	s.Data = v
	return s
}

func (s *ListLivyComputeTokenResponseBody) SetMessage(v string) *ListLivyComputeTokenResponseBody {
	s.Message = &v
	return s
}

func (s *ListLivyComputeTokenResponseBody) SetRequestId(v string) *ListLivyComputeTokenResponseBody {
	s.RequestId = &v
	return s
}

type ListLivyComputeTokenResponseBodyData struct {
	Tokens []*ListLivyComputeTokenResponseBodyDataTokens `json:"tokens,omitempty" xml:"tokens,omitempty" type:"Repeated"`
}

func (s ListLivyComputeTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListLivyComputeTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLivyComputeTokenResponseBodyData) SetTokens(v []*ListLivyComputeTokenResponseBodyDataTokens) *ListLivyComputeTokenResponseBodyData {
	s.Tokens = v
	return s
}

type ListLivyComputeTokenResponseBodyDataTokens struct {
	// example:
	//
	// 1749456094000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// alice
	Createdby *string `json:"createdby,omitempty" xml:"createdby,omitempty"`
	// example:
	//
	// 1749456994000
	ExpireTime *int64 `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// example:
	//
	// 1749456098000
	LastUsedTime *int64 `json:"lastUsedTime,omitempty" xml:"lastUsedTime,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 5d37843fb6f1e8
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// Token ID。
	//
	// example:
	//
	// lctk-xxxxxxxxxxx
	TokenId *string `json:"tokenId,omitempty" xml:"tokenId,omitempty"`
}

func (s ListLivyComputeTokenResponseBodyDataTokens) String() string {
	return tea.Prettify(s)
}

func (s ListLivyComputeTokenResponseBodyDataTokens) GoString() string {
	return s.String()
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) SetCreateTime(v int64) *ListLivyComputeTokenResponseBodyDataTokens {
	s.CreateTime = &v
	return s
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) SetCreatedby(v string) *ListLivyComputeTokenResponseBodyDataTokens {
	s.Createdby = &v
	return s
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) SetExpireTime(v int64) *ListLivyComputeTokenResponseBodyDataTokens {
	s.ExpireTime = &v
	return s
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) SetLastUsedTime(v int64) *ListLivyComputeTokenResponseBodyDataTokens {
	s.LastUsedTime = &v
	return s
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) SetName(v string) *ListLivyComputeTokenResponseBodyDataTokens {
	s.Name = &v
	return s
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) SetToken(v string) *ListLivyComputeTokenResponseBodyDataTokens {
	s.Token = &v
	return s
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) SetTokenId(v string) *ListLivyComputeTokenResponseBodyDataTokens {
	s.TokenId = &v
	return s
}

type ListLivyComputeTokenResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLivyComputeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLivyComputeTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLivyComputeTokenResponse) GoString() string {
	return s.String()
}

func (s *ListLivyComputeTokenResponse) SetHeaders(v map[string]*string) *ListLivyComputeTokenResponse {
	s.Headers = v
	return s
}

func (s *ListLivyComputeTokenResponse) SetStatusCode(v int32) *ListLivyComputeTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLivyComputeTokenResponse) SetBody(v *ListLivyComputeTokenResponseBody) *ListLivyComputeTokenResponse {
	s.Body = v
	return s
}

type ListLogContentsRequest struct {
	// Full path of the file.
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// Length of the log.
	//
	// example:
	//
	// 9999
	Length *int32 `json:"length,omitempty" xml:"length,omitempty"`
	// Start line for query.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListLogContentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLogContentsRequest) GoString() string {
	return s.String()
}

func (s *ListLogContentsRequest) SetFileName(v string) *ListLogContentsRequest {
	s.FileName = &v
	return s
}

func (s *ListLogContentsRequest) SetLength(v int32) *ListLogContentsRequest {
	s.Length = &v
	return s
}

func (s *ListLogContentsRequest) SetOffset(v int32) *ListLogContentsRequest {
	s.Offset = &v
	return s
}

func (s *ListLogContentsRequest) SetRegionId(v string) *ListLogContentsRequest {
	s.RegionId = &v
	return s
}

type ListLogContentsResponseBody struct {
	// Log content.
	ListLogContent *ListLogContentsResponseBodyListLogContent `json:"listLogContent,omitempty" xml:"listLogContent,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListLogContentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLogContentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogContentsResponseBody) SetListLogContent(v *ListLogContentsResponseBodyListLogContent) *ListLogContentsResponseBody {
	s.ListLogContent = v
	return s
}

func (s *ListLogContentsResponseBody) SetRequestId(v string) *ListLogContentsResponseBody {
	s.RequestId = &v
	return s
}

type ListLogContentsResponseBodyListLogContent struct {
	// List of log line contents.
	Contents []*ListLogContentsResponseBodyListLogContentContents `json:"contents,omitempty" xml:"contents,omitempty" type:"Repeated"`
	// Total number of log lines.
	//
	// example:
	//
	// 10
	TotalLength *int64 `json:"totalLength,omitempty" xml:"totalLength,omitempty"`
}

func (s ListLogContentsResponseBodyListLogContent) String() string {
	return tea.Prettify(s)
}

func (s ListLogContentsResponseBodyListLogContent) GoString() string {
	return s.String()
}

func (s *ListLogContentsResponseBodyListLogContent) SetContents(v []*ListLogContentsResponseBodyListLogContentContents) *ListLogContentsResponseBodyListLogContent {
	s.Contents = v
	return s
}

func (s *ListLogContentsResponseBodyListLogContent) SetTotalLength(v int64) *ListLogContentsResponseBodyListLogContent {
	s.TotalLength = &v
	return s
}

type ListLogContentsResponseBodyListLogContentContents struct {
	// Log line content.
	//
	// example:
	//
	// spark pi is 3.14\\n
	LineContent *string `json:"LineContent,omitempty" xml:"LineContent,omitempty"`
}

func (s ListLogContentsResponseBodyListLogContentContents) String() string {
	return tea.Prettify(s)
}

func (s ListLogContentsResponseBodyListLogContentContents) GoString() string {
	return s.String()
}

func (s *ListLogContentsResponseBodyListLogContentContents) SetLineContent(v string) *ListLogContentsResponseBodyListLogContentContents {
	s.LineContent = &v
	return s
}

type ListLogContentsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLogContentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLogContentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLogContentsResponse) GoString() string {
	return s.String()
}

func (s *ListLogContentsResponse) SetHeaders(v map[string]*string) *ListLogContentsResponse {
	s.Headers = v
	return s
}

func (s *ListLogContentsResponse) SetStatusCode(v int32) *ListLogContentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogContentsResponse) SetBody(v *ListLogContentsResponseBody) *ListLogContentsResponse {
	s.Body = v
	return s
}

type ListReleaseVersionsRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The type of the version.
	//
	// Valid values:
	//
	// 	- stable
	//
	// 	- Beta
	//
	// example:
	//
	// stable
	ReleaseType *string `json:"releaseType,omitempty" xml:"releaseType,omitempty"`
	// The version of EMR Serverless Spark.
	//
	// example:
	//
	// esr-2.1 (Spark 3.3.1, Scala 2.12, Java Runtime)
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// The status of the version.
	//
	// Valid values:
	//
	// 	- ONLINE
	//
	// 	- OFFLINE
	//
	// example:
	//
	// ONLINE
	ReleaseVersionStatus *string `json:"releaseVersionStatus,omitempty" xml:"releaseVersionStatus,omitempty"`
	ServiceFilter        *string `json:"serviceFilter,omitempty" xml:"serviceFilter,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// w-d2d82aa09155****
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
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

func (s *ListReleaseVersionsRequest) SetServiceFilter(v string) *ListReleaseVersionsRequest {
	s.ServiceFilter = &v
	return s
}

func (s *ListReleaseVersionsRequest) SetWorkspaceId(v string) *ListReleaseVersionsRequest {
	s.WorkspaceId = &v
	return s
}

type ListReleaseVersionsResponseBody struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The versions.
	ReleaseVersions []*ListReleaseVersionsResponseBodyReleaseVersions `json:"releaseVersions,omitempty" xml:"releaseVersions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries returned.
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
	// The version number of open source Spark.
	//
	// example:
	//
	// Spark 3.3.1
	CommunityVersion *string `json:"communityVersion,omitempty" xml:"communityVersion,omitempty"`
	// The CPU architectures.
	CpuArchitectures []*string `json:"cpuArchitectures,omitempty" xml:"cpuArchitectures,omitempty" type:"Repeated"`
	// The version number.
	//
	// example:
	//
	// esr-2.1 (Spark 3.3.1, Scala 2.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// Indicates whether the Fusion engine is used for acceleration.
	//
	// example:
	//
	// true
	Fusion *bool `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1716215854101
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The type of the Infrastructure as a Service (IaaS) layer.
	//
	// example:
	//
	// ASI
	IaasType *string `json:"iaasType,omitempty" xml:"iaasType,omitempty"`
	// The version number.
	//
	// example:
	//
	// esr-2.1 (Spark 3.3.1, Scala 2.12, Java Runtime)
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// The version of Scala.
	//
	// example:
	//
	// 2.12
	ScalaVersion *string `json:"scalaVersion,omitempty" xml:"scalaVersion,omitempty"`
	// The status of the version.
	//
	// example:
	//
	// ONLINE
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The type of the version.
	//
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

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetDisplayReleaseVersion(v string) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetFusion(v bool) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.Fusion = &v
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
	// The session type.
	//
	// Valid values:
	//
	// 	- NOTEBOOK
	//
	// 	- THRIFT
	//
	// 	- SQL
	//
	// example:
	//
	// SQL
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token that is used in the request to retrieve a new page of results.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The name of the queue.
	//
	// example:
	//
	// root
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The name of the job.
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

func (s *ListSessionClustersRequest) SetKind(v string) *ListSessionClustersRequest {
	s.Kind = &v
	return s
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
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The sessions.
	SessionClusters []*ListSessionClustersResponseBodySessionClusters `json:"sessionClusters,omitempty" xml:"sessionClusters,omitempty" type:"Repeated"`
	// The total number of entries returned.
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
	// The session configurations, which are equivalent to the configurations of the Spark job.
	ApplicationConfigs []*ListSessionClustersResponseBodySessionClustersApplicationConfigs `json:"applicationConfigs,omitempty" xml:"applicationConfigs,omitempty" type:"Repeated"`
	// The automatic startup configurations.
	AutoStartConfiguration *ListSessionClustersResponseBodySessionClustersAutoStartConfiguration `json:"autoStartConfiguration,omitempty" xml:"autoStartConfiguration,omitempty" type:"Struct"`
	// The configurations of automatic termination.
	AutoStopConfiguration *ListSessionClustersResponseBodySessionClustersAutoStopConfiguration `json:"autoStopConfiguration,omitempty" xml:"autoStopConfiguration,omitempty" type:"Struct"`
	// The version of the Spark engine.
	//
	// example:
	//
	// esr-4.0.0 (Spark 3.5.2, Scala 2.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// The public endpoint of the Thrift server.
	//
	// example:
	//
	// emr-spark-gateway-cn-hangzhou.data.aliyun.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The internal endpoint of the Thrift server.
	//
	// example:
	//
	// emr-spark-gateway-cn-hangzhou-internal.data.aliyuncs.com
	DomainInner *string `json:"domainInner,omitempty" xml:"domainInner,omitempty"`
	// The ID of the job that is associated with the session.
	//
	// example:
	//
	// TSK-xxxxxxxxx
	DraftId *string `json:"draftId,omitempty" xml:"draftId,omitempty"`
	// The additional metadata of the session.
	//
	// example:
	//
	// {"extraInfoKey":"extraInfoValue"}
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// Indicates whether the Fusion engine is used for acceleration.
	//
	// example:
	//
	// false
	Fusion *bool `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1732267598000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The session type.
	//
	// Valid values:
	//
	// 	- NOTEBOOK
	//
	// 	- THRIFT
	//
	// 	- SQL
	//
	// example:
	//
	// SQL
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The name of the session.
	//
	// example:
	//
	// adhoc_query
	Name                  *string `json:"name,omitempty" xml:"name,omitempty"`
	PublicEndpointEnabled *bool   `json:"publicEndpointEnabled,omitempty" xml:"publicEndpointEnabled,omitempty"`
	// The name of the queue that is used to run the session.
	//
	// example:
	//
	// dev_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// The version of EMR Serverless Spark.
	//
	// example:
	//
	// esr-2.1
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// The session ID.
	//
	// example:
	//
	// sc-123131
	SessionClusterId *string `json:"sessionClusterId,omitempty" xml:"sessionClusterId,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1732267598000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The status of the session.
	//
	// 	- Starting
	//
	// 	- Running
	//
	// 	- Stopping
	//
	// 	- Stopped
	//
	// 	- Error
	//
	// example:
	//
	// Running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The details of the most recent status change of the session.
	StateChangeReason *ListSessionClustersResponseBodySessionClustersStateChangeReason `json:"stateChangeReason,omitempty" xml:"stateChangeReason,omitempty" type:"Struct"`
	// The user ID.
	//
	// example:
	//
	// 123131
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The username.
	//
	// example:
	//
	// test_user
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// The Spark UI of the session.
	//
	// example:
	//
	// http://spark-ui-xxxx
	WebUI *string `json:"webUI,omitempty" xml:"webUI,omitempty"`
	// The workspace ID.
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

func (s *ListSessionClustersResponseBodySessionClusters) SetDisplayReleaseVersion(v string) *ListSessionClustersResponseBodySessionClusters {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetDomain(v string) *ListSessionClustersResponseBodySessionClusters {
	s.Domain = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetDomainInner(v string) *ListSessionClustersResponseBodySessionClusters {
	s.DomainInner = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetDraftId(v string) *ListSessionClustersResponseBodySessionClusters {
	s.DraftId = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetExtra(v string) *ListSessionClustersResponseBodySessionClusters {
	s.Extra = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetFusion(v bool) *ListSessionClustersResponseBodySessionClusters {
	s.Fusion = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetGmtCreate(v int64) *ListSessionClustersResponseBodySessionClusters {
	s.GmtCreate = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetKind(v string) *ListSessionClustersResponseBodySessionClusters {
	s.Kind = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetName(v string) *ListSessionClustersResponseBodySessionClusters {
	s.Name = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetPublicEndpointEnabled(v bool) *ListSessionClustersResponseBodySessionClusters {
	s.PublicEndpointEnabled = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetQueueName(v string) *ListSessionClustersResponseBodySessionClusters {
	s.QueueName = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetReleaseVersion(v string) *ListSessionClustersResponseBodySessionClusters {
	s.ReleaseVersion = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetSessionClusterId(v string) *ListSessionClustersResponseBodySessionClusters {
	s.SessionClusterId = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetStartTime(v int64) *ListSessionClustersResponseBodySessionClusters {
	s.StartTime = &v
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

func (s *ListSessionClustersResponseBodySessionClusters) SetWebUI(v string) *ListSessionClustersResponseBodySessionClusters {
	s.WebUI = &v
	return s
}

func (s *ListSessionClustersResponseBodySessionClusters) SetWorkspaceId(v string) *ListSessionClustersResponseBodySessionClusters {
	s.WorkspaceId = &v
	return s
}

type ListSessionClustersResponseBodySessionClustersApplicationConfigs struct {
	// The name of the configuration file.
	//
	// example:
	//
	// spark-default.conf
	ConfigFileName *string `json:"configFileName,omitempty" xml:"configFileName,omitempty"`
	// The key of the configuration.
	//
	// example:
	//
	// spark.app.name
	ConfigItemKey *string `json:"configItemKey,omitempty" xml:"configItemKey,omitempty"`
	// The configuration value.
	//
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
	// Indicates whether automatic startup is enabled.
	//
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
	// Indicates whether automatic termination is enabled.
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The idle timeout period. The session is automatically terminated when the idle timeout period is exceeded.
	//
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
	// The status change code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The status change message.
	//
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
	// The environment type.
	//
	// Valid values:
	//
	// 	- dev
	//
	// 	- production
	//
	// example:
	//
	// production
	Environment *string `json:"environment,omitempty" xml:"environment,omitempty"`
	// The region ID.
	//
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
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The list of queues.
	Queues []*ListWorkspaceQueuesResponseBodyQueues `json:"queues,omitempty" xml:"queues,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries returned.
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
	// The operations allowed for the queue.
	AllowActions []*ListWorkspaceQueuesResponseBodyQueuesAllowActions `json:"allowActions,omitempty" xml:"allowActions,omitempty" type:"Repeated"`
	// The time when the workspace was created.
	//
	// example:
	//
	// 1684115879955
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the user who created the queue.
	//
	// example:
	//
	// 237109
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The environment types of the queue.
	Environments []*string `json:"environments,omitempty" xml:"environments,omitempty" type:"Repeated"`
	// The maximum capacity of resources that can be used in the queue.
	//
	// example:
	//
	// {"cpu": "2","memory": "2Gi"}
	MaxResource *string `json:"maxResource,omitempty" xml:"maxResource,omitempty"`
	// The minimum capacity of resources that can be used in the queue.
	//
	// example:
	//
	// {"cpu": "2","memory": "2Gi"}
	MinResource *string `json:"minResource,omitempty" xml:"minResource,omitempty"`
	// The billing method. Valid values:
	//
	// 	- PayAsYouGo
	//
	// 	- Pre
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The queue label.
	//
	// example:
	//
	// dev_queue
	Properties *string `json:"properties,omitempty" xml:"properties,omitempty"`
	// The name of the queue.
	//
	// example:
	//
	// dev_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// The queue architecture.
	//
	// example:
	//
	// {"arch": "x86"}
	QueueScope *string `json:"queueScope,omitempty" xml:"queueScope,omitempty"`
	// The status of the queue.
	//
	// example:
	//
	// RUNNING
	QueueStatus *string `json:"queueStatus,omitempty" xml:"queueStatus,omitempty"`
	// The type of the queue. Valid values:
	//
	// 	- instance
	//
	// 	- instanceChildren
	//
	// example:
	//
	// instance, instanceChildren
	QueueType *string `json:"queueType,omitempty" xml:"queueType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The capacity of resources that are used in the queue.
	//
	// example:
	//
	// {"cpu": "2","memory": "2Gi"}
	UsedResource *string `json:"usedResource,omitempty" xml:"usedResource,omitempty"`
	// The workspace ID.
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

func (s *ListWorkspaceQueuesResponseBodyQueues) SetCreateTime(v int64) *ListWorkspaceQueuesResponseBodyQueues {
	s.CreateTime = &v
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

func (s *ListWorkspaceQueuesResponseBodyQueues) SetPaymentType(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.PaymentType = &v
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
	// The Alibaba Cloud Resource Name (ARN) of a behavior.
	//
	// example:
	//
	// acs:emr::workspaceId:action/create_queue
	ActionArn *string `json:"actionArn,omitempty" xml:"actionArn,omitempty"`
	// The name of the permission.
	//
	// example:
	//
	// view
	ActionName *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	// The dependencies of the operation.
	//
	// example:
	//
	// ["view"]
	Dependencies []*string `json:"dependencies,omitempty" xml:"dependencies,omitempty" type:"Repeated"`
	// The description of the operation.
	//
	// example:
	//
	// 文件目录遍历、文件浏览
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name of the permission.
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
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The name of the workspace. Fuzzy match is supported.
	//
	// example:
	//
	// test_workspace
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The state of the workspace.
	//
	// example:
	//
	// running
	State *string                     `json:"state,omitempty" xml:"state,omitempty"`
	Tag   []*ListWorkspacesRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
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

func (s *ListWorkspacesRequest) SetTag(v []*ListWorkspacesRequestTag) *ListWorkspacesRequest {
	s.Tag = v
	return s
}

type ListWorkspacesRequestTag struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListWorkspacesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesRequestTag) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequestTag) SetKey(v string) *ListWorkspacesRequestTag {
	s.Key = &v
	return s
}

func (s *ListWorkspacesRequestTag) SetValue(v string) *ListWorkspacesRequestTag {
	s.Value = &v
	return s
}

type ListWorkspacesShrinkRequest struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The name of the workspace. Fuzzy match is supported.
	//
	// example:
	//
	// test_workspace
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The state of the workspace.
	//
	// example:
	//
	// running
	State     *string `json:"state,omitempty" xml:"state,omitempty"`
	TagShrink *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ListWorkspacesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesShrinkRequest) SetMaxResults(v int32) *ListWorkspacesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetName(v string) *ListWorkspacesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetNextToken(v string) *ListWorkspacesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetRegionId(v string) *ListWorkspacesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetState(v string) *ListWorkspacesShrinkRequest {
	s.State = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetTagShrink(v string) *ListWorkspacesShrinkRequest {
	s.TagShrink = &v
	return s
}

type ListWorkspacesResponseBody struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// The queried workspaces.
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
	// Specifies whether to enable auto-renewal. This parameter is required only if the paymentType parameter is set to Pre.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// The auto-renewal duration. This parameter is required only if the paymentType parameter is set to Pre.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"autoRenewPeriod,omitempty" xml:"autoRenewPeriod,omitempty"`
	// The unit of the auto-renewal duration. This parameter is required only if the paymentType parameter is set to Pre.
	//
	// example:
	//
	// YEAR, MONTH, WEEK, DAY, HOUR, MINUTE
	AutoRenewPeriodUnit *string `json:"autoRenewPeriodUnit,omitempty" xml:"autoRenewPeriodUnit,omitempty"`
	// The time when the workflow was created.
	//
	// example:
	//
	// 1684115879955
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The information of the Data Lake Formation (DLF) catalog.
	//
	// example:
	//
	// default
	DlfCatalogId *string `json:"dlfCatalogId,omitempty" xml:"dlfCatalogId,omitempty"`
	// The version of DLF.
	//
	// example:
	//
	// 1.0
	DlfType *string `json:"dlfType,omitempty" xml:"dlfType,omitempty"`
	// The subscription period. This parameter is required only if the paymentType parameter is set to Pre.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// The end of the end time range.
	//
	// example:
	//
	// 1687103999999
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The failure reason.
	//
	// example:
	//
	// out of stock
	FailReason *string `json:"failReason,omitempty" xml:"failReason,omitempty"`
	// The unit of the subscription duration.
	//
	// example:
	//
	// YEAR, MONTH, WEEK, DAY, HOUR, MINUTE
	PaymentDurationUnit *string `json:"paymentDurationUnit,omitempty" xml:"paymentDurationUnit,omitempty"`
	// The status of the payment.
	//
	// example:
	//
	// PAID/UNPAID
	PaymentStatus *string `json:"paymentStatus,omitempty" xml:"paymentStatus,omitempty"`
	// The billing method. Valid values:
	//
	// - PayAsYouGo
	//
	// - Pre
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The information about the subscription quota.
	PrePaidQuota *ListWorkspacesResponseBodyWorkspacesPrePaidQuota `json:"prePaidQuota,omitempty" xml:"prePaidQuota,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The reason why the workspace is released.
	//
	// example:
	//
	// SERVICE_RELEASE
	ReleaseType *string `json:"releaseType,omitempty" xml:"releaseType,omitempty"`
	// The resource specifications.
	//
	// example:
	//
	// 100cu
	ResourceSpec *string `json:"resourceSpec,omitempty" xml:"resourceSpec,omitempty"`
	// The reason of the job status change.
	StateChangeReason *ListWorkspacesResponseBodyWorkspacesStateChangeReason `json:"stateChangeReason,omitempty" xml:"stateChangeReason,omitempty" type:"Struct"`
	// The OSS path.
	//
	// example:
	//
	// spark-result
	Storage *string                                     `json:"storage,omitempty" xml:"storage,omitempty"`
	Tags    []*ListWorkspacesResponseBodyWorkspacesTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The workspace ID.
	//
	// example:
	//
	// w-******
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// spark-1
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
	// The workspace status.
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

func (s *ListWorkspacesResponseBodyWorkspaces) SetDlfType(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.DlfType = &v
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

func (s *ListWorkspacesResponseBodyWorkspaces) SetPrePaidQuota(v *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) *ListWorkspacesResponseBodyWorkspaces {
	s.PrePaidQuota = v
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

func (s *ListWorkspacesResponseBodyWorkspaces) SetTags(v []*ListWorkspacesResponseBodyWorkspacesTags) *ListWorkspacesResponseBodyWorkspaces {
	s.Tags = v
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

type ListWorkspacesResponseBodyWorkspacesPrePaidQuota struct {
	// The amount of resources that are allocated by a subscription quota.
	//
	// example:
	//
	// {\\"cpu\\":\\"1\\",\\"memory\\":\\"4Gi\\",\\"cu\\":\\"1\\"}
	AllocatedResource *string `json:"allocatedResource,omitempty" xml:"allocatedResource,omitempty"`
	// Indicates whether auto-renewal is enabled for the subscription quota.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AutoRenewal *bool `json:"autoRenewal,omitempty" xml:"autoRenewal,omitempty"`
	// The creation time of the subscription quota.
	//
	// example:
	//
	// 1745683200000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The expiration time of the subscription quota.
	//
	// example:
	//
	// 1740537153000
	ExpireTime *int64 `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The ID of the instance that is generated when you purchase the subscription quota.
	//
	// example:
	//
	// i-abc12345
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The maximum amount of resources that can be used in a subscription quota.
	//
	// example:
	//
	// {\\"cpu\\":\\"1\\",\\"memory\\":\\"4Gi\\",\\"cu\\":\\"1\\"}
	MaxResource *string `json:"maxResource,omitempty" xml:"maxResource,omitempty"`
	OrderId     *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// The status of the subscription quota. Valid values:
	//
	// 	- NORMAL
	//
	// 	- WAIT_FOR_EXPIRE
	//
	// 	- EXPIRED
	//
	// example:
	//
	// NORMAL
	PaymentStatus *string `json:"paymentStatus,omitempty" xml:"paymentStatus,omitempty"`
	// The amount of resources that are used.
	//
	// example:
	//
	// {\\"cpu\\":\\"0\\",\\"memory\\":\\"0Gi\\",\\"cu\\":\\"0\\"}
	UsedResource *string `json:"usedResource,omitempty" xml:"usedResource,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspacesPrePaidQuota) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspacesPrePaidQuota) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) SetAllocatedResource(v string) *ListWorkspacesResponseBodyWorkspacesPrePaidQuota {
	s.AllocatedResource = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) SetAutoRenewal(v bool) *ListWorkspacesResponseBodyWorkspacesPrePaidQuota {
	s.AutoRenewal = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) SetCreateTime(v int64) *ListWorkspacesResponseBodyWorkspacesPrePaidQuota {
	s.CreateTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) SetExpireTime(v int64) *ListWorkspacesResponseBodyWorkspacesPrePaidQuota {
	s.ExpireTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) SetInstanceId(v string) *ListWorkspacesResponseBodyWorkspacesPrePaidQuota {
	s.InstanceId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) SetMaxResource(v string) *ListWorkspacesResponseBodyWorkspacesPrePaidQuota {
	s.MaxResource = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) SetOrderId(v string) *ListWorkspacesResponseBodyWorkspacesPrePaidQuota {
	s.OrderId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) SetPaymentStatus(v string) *ListWorkspacesResponseBodyWorkspacesPrePaidQuota {
	s.PaymentStatus = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) SetUsedResource(v string) *ListWorkspacesResponseBodyWorkspacesPrePaidQuota {
	s.UsedResource = &v
	return s
}

type ListWorkspacesResponseBodyWorkspacesStateChangeReason struct {
	// The error code.
	//
	// example:
	//
	// 0
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error message.
	//
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

type ListWorkspacesResponseBodyWorkspacesTags struct {
	TagKey   *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspacesTags) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspacesTags) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspacesTags) SetTagKey(v string) *ListWorkspacesResponseBodyWorkspacesTags {
	s.TagKey = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesTags) SetTagValue(v string) *ListWorkspacesResponseBodyWorkspacesTags {
	s.TagValue = &v
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

type RefreshLivyComputeTokenRequest struct {
	AutoExpireConfiguration *RefreshLivyComputeTokenRequestAutoExpireConfiguration `json:"autoExpireConfiguration,omitempty" xml:"autoExpireConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// fe86812667f04v343
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s RefreshLivyComputeTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshLivyComputeTokenRequest) GoString() string {
	return s.String()
}

func (s *RefreshLivyComputeTokenRequest) SetAutoExpireConfiguration(v *RefreshLivyComputeTokenRequestAutoExpireConfiguration) *RefreshLivyComputeTokenRequest {
	s.AutoExpireConfiguration = v
	return s
}

func (s *RefreshLivyComputeTokenRequest) SetName(v string) *RefreshLivyComputeTokenRequest {
	s.Name = &v
	return s
}

func (s *RefreshLivyComputeTokenRequest) SetToken(v string) *RefreshLivyComputeTokenRequest {
	s.Token = &v
	return s
}

func (s *RefreshLivyComputeTokenRequest) SetRegionId(v string) *RefreshLivyComputeTokenRequest {
	s.RegionId = &v
	return s
}

type RefreshLivyComputeTokenRequestAutoExpireConfiguration struct {
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 7
	ExpireDays *int32 `json:"expireDays,omitempty" xml:"expireDays,omitempty"`
}

func (s RefreshLivyComputeTokenRequestAutoExpireConfiguration) String() string {
	return tea.Prettify(s)
}

func (s RefreshLivyComputeTokenRequestAutoExpireConfiguration) GoString() string {
	return s.String()
}

func (s *RefreshLivyComputeTokenRequestAutoExpireConfiguration) SetEnable(v bool) *RefreshLivyComputeTokenRequestAutoExpireConfiguration {
	s.Enable = &v
	return s
}

func (s *RefreshLivyComputeTokenRequestAutoExpireConfiguration) SetExpireDays(v int32) *RefreshLivyComputeTokenRequestAutoExpireConfiguration {
	s.ExpireDays = &v
	return s
}

type RefreshLivyComputeTokenResponseBody struct {
	// example:
	//
	// 1000000
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RefreshLivyComputeTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshLivyComputeTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshLivyComputeTokenResponseBody) SetCode(v string) *RefreshLivyComputeTokenResponseBody {
	s.Code = &v
	return s
}

func (s *RefreshLivyComputeTokenResponseBody) SetMessage(v string) *RefreshLivyComputeTokenResponseBody {
	s.Message = &v
	return s
}

func (s *RefreshLivyComputeTokenResponseBody) SetRequestId(v string) *RefreshLivyComputeTokenResponseBody {
	s.RequestId = &v
	return s
}

type RefreshLivyComputeTokenResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshLivyComputeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshLivyComputeTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshLivyComputeTokenResponse) GoString() string {
	return s.String()
}

func (s *RefreshLivyComputeTokenResponse) SetHeaders(v map[string]*string) *RefreshLivyComputeTokenResponse {
	s.Headers = v
	return s
}

func (s *RefreshLivyComputeTokenResponse) SetStatusCode(v int32) *RefreshLivyComputeTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshLivyComputeTokenResponse) SetBody(v *RefreshLivyComputeTokenResponseBody) *RefreshLivyComputeTokenResponse {
	s.Body = v
	return s
}

type StartJobRunRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 8e6aae2810c8f67229ca70bb31cd6028
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The code type of the job. Valid values:
	//
	// 	- SQL
	//
	// 	- JAR
	//
	// 	- PYTHON
	//
	// example:
	//
	// SQL
	CodeType *string `json:"codeType,omitempty" xml:"codeType,omitempty"`
	// The advanced configurations of Spark.
	ConfigurationOverrides *StartJobRunRequestConfigurationOverrides `json:"configurationOverrides,omitempty" xml:"configurationOverrides,omitempty" type:"Struct"`
	// The version of the Spark engine.
	//
	// example:
	//
	// esr-3.3.1
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// The timeout period of the job.
	//
	// example:
	//
	// 100
	ExecutionTimeoutSeconds *int32 `json:"executionTimeoutSeconds,omitempty" xml:"executionTimeoutSeconds,omitempty"`
	// Specifies whether to enable Fusion engine for acceleration.
	//
	// example:
	//
	// false
	Fusion *bool `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// The information about Spark Driver.
	JobDriver *JobDriver `json:"jobDriver,omitempty" xml:"jobDriver,omitempty"`
	// The job ID.
	//
	// example:
	//
	// jr-12345
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// spark_job_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The version number of Spark.
	//
	// example:
	//
	// esr-3.3.1
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// The name of the resource queue on which the Spark job runs.
	//
	// example:
	//
	// dev_queue
	ResourceQueueId *string `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	// The tags of the job.
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The region ID.
	//
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

func (s *StartJobRunRequest) SetDisplayReleaseVersion(v string) *StartJobRunRequest {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *StartJobRunRequest) SetExecutionTimeoutSeconds(v int32) *StartJobRunRequest {
	s.ExecutionTimeoutSeconds = &v
	return s
}

func (s *StartJobRunRequest) SetFusion(v bool) *StartJobRunRequest {
	s.Fusion = &v
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
	// The SparkConf objects.
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
	// The configuration file of SparkConf.
	//
	// example:
	//
	// spark-default.conf
	ConfigFileName *string `json:"configFileName,omitempty" xml:"configFileName,omitempty"`
	// The key of SparkConf.
	//
	// example:
	//
	// spark.app.name
	ConfigItemKey *string `json:"configItemKey,omitempty" xml:"configItemKey,omitempty"`
	// The value of SparkConf.
	//
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
	// The job ID.
	//
	// example:
	//
	// jr-54321
	JobRunId *string `json:"jobRunId,omitempty" xml:"jobRunId,omitempty"`
	// The request ID.
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

type StartLivyComputeRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s StartLivyComputeRequest) String() string {
	return tea.Prettify(s)
}

func (s StartLivyComputeRequest) GoString() string {
	return s.String()
}

func (s *StartLivyComputeRequest) SetRegionId(v string) *StartLivyComputeRequest {
	s.RegionId = &v
	return s
}

type StartLivyComputeResponseBody struct {
	// example:
	//
	// 1000000
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StartLivyComputeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartLivyComputeResponseBody) GoString() string {
	return s.String()
}

func (s *StartLivyComputeResponseBody) SetCode(v string) *StartLivyComputeResponseBody {
	s.Code = &v
	return s
}

func (s *StartLivyComputeResponseBody) SetMessage(v string) *StartLivyComputeResponseBody {
	s.Message = &v
	return s
}

func (s *StartLivyComputeResponseBody) SetRequestId(v string) *StartLivyComputeResponseBody {
	s.RequestId = &v
	return s
}

type StartLivyComputeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartLivyComputeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartLivyComputeResponse) String() string {
	return tea.Prettify(s)
}

func (s StartLivyComputeResponse) GoString() string {
	return s.String()
}

func (s *StartLivyComputeResponse) SetHeaders(v map[string]*string) *StartLivyComputeResponse {
	s.Headers = v
	return s
}

func (s *StartLivyComputeResponse) SetStatusCode(v int32) *StartLivyComputeResponse {
	s.StatusCode = &v
	return s
}

func (s *StartLivyComputeResponse) SetBody(v *StartLivyComputeResponseBody) *StartLivyComputeResponse {
	s.Body = v
	return s
}

type StartProcessInstanceRequest struct {
	Action   *string `json:"action,omitempty" xml:"action,omitempty"`
	Comments *string `json:"comments,omitempty" xml:"comments,omitempty"`
	Email    *string `json:"email,omitempty" xml:"email,omitempty"`
	Interval *string `json:"interval,omitempty" xml:"interval,omitempty"`
	// Specifies whether to run the workflow in the production environment.
	//
	// example:
	//
	// false
	IsProd *bool `json:"isProd,omitempty" xml:"isProd,omitempty"`
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12***********
	ProcessDefinitionCode *int64 `json:"processDefinitionCode,omitempty" xml:"processDefinitionCode,omitempty"`
	// The code of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// SS
	ProductNamespace *string `json:"productNamespace,omitempty" xml:"productNamespace,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The queue on which the workflow runs.
	//
	// example:
	//
	// root_queue
	RuntimeQueue *string `json:"runtimeQueue,omitempty" xml:"runtimeQueue,omitempty"`
	// The hash code of the version.
	//
	// example:
	//
	// dh*********
	VersionHashCode *string `json:"versionHashCode,omitempty" xml:"versionHashCode,omitempty"`
	// The version number of the workflow.
	//
	// example:
	//
	// 1
	VersionNumber *int32 `json:"versionNumber,omitempty" xml:"versionNumber,omitempty"`
}

func (s StartProcessInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceRequest) SetAction(v string) *StartProcessInstanceRequest {
	s.Action = &v
	return s
}

func (s *StartProcessInstanceRequest) SetComments(v string) *StartProcessInstanceRequest {
	s.Comments = &v
	return s
}

func (s *StartProcessInstanceRequest) SetEmail(v string) *StartProcessInstanceRequest {
	s.Email = &v
	return s
}

func (s *StartProcessInstanceRequest) SetInterval(v string) *StartProcessInstanceRequest {
	s.Interval = &v
	return s
}

func (s *StartProcessInstanceRequest) SetIsProd(v bool) *StartProcessInstanceRequest {
	s.IsProd = &v
	return s
}

func (s *StartProcessInstanceRequest) SetProcessDefinitionCode(v int64) *StartProcessInstanceRequest {
	s.ProcessDefinitionCode = &v
	return s
}

func (s *StartProcessInstanceRequest) SetProductNamespace(v string) *StartProcessInstanceRequest {
	s.ProductNamespace = &v
	return s
}

func (s *StartProcessInstanceRequest) SetRegionId(v string) *StartProcessInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartProcessInstanceRequest) SetRuntimeQueue(v string) *StartProcessInstanceRequest {
	s.RuntimeQueue = &v
	return s
}

func (s *StartProcessInstanceRequest) SetVersionHashCode(v string) *StartProcessInstanceRequest {
	s.VersionHashCode = &v
	return s
}

func (s *StartProcessInstanceRequest) SetVersionNumber(v int32) *StartProcessInstanceRequest {
	s.VersionNumber = &v
	return s
}

type StartProcessInstanceResponseBody struct {
	// The code that is returned by the backend server.
	//
	// example:
	//
	// 1400009
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// {\\"sessionBizId\\": \\"sc-dc85644dba1c8c63\\", \\"bizId\\": \\"st-aeed3b0d4f87418a9a9dcbd757477658\\", \\"gmtCreated\\": \\"Thu Sep 12 02:28:45 UTC 2024\\"}
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// Indicates whether the workflow fails to be run manually.
	//
	// example:
	//
	// false
	Failed *bool `json:"failed,omitempty" xml:"failed,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The description of the returned code.
	//
	// example:
	//
	// No permission for resource action
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StartProcessInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceResponseBody) SetCode(v int32) *StartProcessInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *StartProcessInstanceResponseBody) SetData(v interface{}) *StartProcessInstanceResponseBody {
	s.Data = v
	return s
}

func (s *StartProcessInstanceResponseBody) SetFailed(v bool) *StartProcessInstanceResponseBody {
	s.Failed = &v
	return s
}

func (s *StartProcessInstanceResponseBody) SetHttpStatusCode(v int32) *StartProcessInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartProcessInstanceResponseBody) SetMsg(v string) *StartProcessInstanceResponseBody {
	s.Msg = &v
	return s
}

func (s *StartProcessInstanceResponseBody) SetRequestId(v string) *StartProcessInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartProcessInstanceResponseBody) SetSuccess(v bool) *StartProcessInstanceResponseBody {
	s.Success = &v
	return s
}

type StartProcessInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartProcessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartProcessInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceResponse) SetHeaders(v map[string]*string) *StartProcessInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartProcessInstanceResponse) SetStatusCode(v int32) *StartProcessInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartProcessInstanceResponse) SetBody(v *StartProcessInstanceResponseBody) *StartProcessInstanceResponse {
	s.Body = v
	return s
}

type StartSessionClusterRequest struct {
	// The queue name.
	//
	// example:
	//
	// root_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// The session ID.
	//
	// example:
	//
	// sc-xxxxxxxxxxx
	SessionClusterId *string `json:"sessionClusterId,omitempty" xml:"sessionClusterId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s StartSessionClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s StartSessionClusterRequest) GoString() string {
	return s.String()
}

func (s *StartSessionClusterRequest) SetQueueName(v string) *StartSessionClusterRequest {
	s.QueueName = &v
	return s
}

func (s *StartSessionClusterRequest) SetSessionClusterId(v string) *StartSessionClusterRequest {
	s.SessionClusterId = &v
	return s
}

func (s *StartSessionClusterRequest) SetRegionId(v string) *StartSessionClusterRequest {
	s.RegionId = &v
	return s
}

type StartSessionClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// w-******
	SessionClusterId *string `json:"sessionClusterId,omitempty" xml:"sessionClusterId,omitempty"`
}

func (s StartSessionClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartSessionClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StartSessionClusterResponseBody) SetRequestId(v string) *StartSessionClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSessionClusterResponseBody) SetSessionClusterId(v string) *StartSessionClusterResponseBody {
	s.SessionClusterId = &v
	return s
}

type StartSessionClusterResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartSessionClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartSessionClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s StartSessionClusterResponse) GoString() string {
	return s.String()
}

func (s *StartSessionClusterResponse) SetHeaders(v map[string]*string) *StartSessionClusterResponse {
	s.Headers = v
	return s
}

func (s *StartSessionClusterResponse) SetStatusCode(v int32) *StartSessionClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *StartSessionClusterResponse) SetBody(v *StartSessionClusterResponseBody) *StartSessionClusterResponse {
	s.Body = v
	return s
}

type StopLivyComputeRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s StopLivyComputeRequest) String() string {
	return tea.Prettify(s)
}

func (s StopLivyComputeRequest) GoString() string {
	return s.String()
}

func (s *StopLivyComputeRequest) SetRegionId(v string) *StopLivyComputeRequest {
	s.RegionId = &v
	return s
}

type StopLivyComputeResponseBody struct {
	// example:
	//
	// 1000000
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 484D9DDA-300D-525E-AF7A-0CCCA5C64A7A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StopLivyComputeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopLivyComputeResponseBody) GoString() string {
	return s.String()
}

func (s *StopLivyComputeResponseBody) SetCode(v string) *StopLivyComputeResponseBody {
	s.Code = &v
	return s
}

func (s *StopLivyComputeResponseBody) SetMessage(v string) *StopLivyComputeResponseBody {
	s.Message = &v
	return s
}

func (s *StopLivyComputeResponseBody) SetRequestId(v string) *StopLivyComputeResponseBody {
	s.RequestId = &v
	return s
}

type StopLivyComputeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopLivyComputeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopLivyComputeResponse) String() string {
	return tea.Prettify(s)
}

func (s StopLivyComputeResponse) GoString() string {
	return s.String()
}

func (s *StopLivyComputeResponse) SetHeaders(v map[string]*string) *StopLivyComputeResponse {
	s.Headers = v
	return s
}

func (s *StopLivyComputeResponse) SetStatusCode(v int32) *StopLivyComputeResponse {
	s.StatusCode = &v
	return s
}

func (s *StopLivyComputeResponse) SetBody(v *StopLivyComputeResponseBody) *StopLivyComputeResponse {
	s.Body = v
	return s
}

type StopSessionClusterRequest struct {
	// The queue name.
	//
	// example:
	//
	// root_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// The session ID.
	//
	// example:
	//
	// sc-xxxxxxxxxxxx
	SessionClusterId *string `json:"sessionClusterId,omitempty" xml:"sessionClusterId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s StopSessionClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s StopSessionClusterRequest) GoString() string {
	return s.String()
}

func (s *StopSessionClusterRequest) SetQueueName(v string) *StopSessionClusterRequest {
	s.QueueName = &v
	return s
}

func (s *StopSessionClusterRequest) SetSessionClusterId(v string) *StopSessionClusterRequest {
	s.SessionClusterId = &v
	return s
}

func (s *StopSessionClusterRequest) SetRegionId(v string) *StopSessionClusterRequest {
	s.RegionId = &v
	return s
}

type StopSessionClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// w-******
	SessionClusterId *string `json:"sessionClusterId,omitempty" xml:"sessionClusterId,omitempty"`
}

func (s StopSessionClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopSessionClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StopSessionClusterResponseBody) SetRequestId(v string) *StopSessionClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopSessionClusterResponseBody) SetSessionClusterId(v string) *StopSessionClusterResponseBody {
	s.SessionClusterId = &v
	return s
}

type StopSessionClusterResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopSessionClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopSessionClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s StopSessionClusterResponse) GoString() string {
	return s.String()
}

func (s *StopSessionClusterResponse) SetHeaders(v map[string]*string) *StopSessionClusterResponse {
	s.Headers = v
	return s
}

func (s *StopSessionClusterResponse) SetStatusCode(v int32) *StopSessionClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *StopSessionClusterResponse) SetBody(v *StopSessionClusterResponseBody) *StopSessionClusterResponse {
	s.Body = v
	return s
}

type TerminateSqlStatementRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s TerminateSqlStatementRequest) String() string {
	return tea.Prettify(s)
}

func (s TerminateSqlStatementRequest) GoString() string {
	return s.String()
}

func (s *TerminateSqlStatementRequest) SetRegionId(v string) *TerminateSqlStatementRequest {
	s.RegionId = &v
	return s
}

type TerminateSqlStatementResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s TerminateSqlStatementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TerminateSqlStatementResponseBody) GoString() string {
	return s.String()
}

func (s *TerminateSqlStatementResponseBody) SetRequestId(v string) *TerminateSqlStatementResponseBody {
	s.RequestId = &v
	return s
}

type TerminateSqlStatementResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TerminateSqlStatementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TerminateSqlStatementResponse) String() string {
	return tea.Prettify(s)
}

func (s TerminateSqlStatementResponse) GoString() string {
	return s.String()
}

func (s *TerminateSqlStatementResponse) SetHeaders(v map[string]*string) *TerminateSqlStatementResponse {
	s.Headers = v
	return s
}

func (s *TerminateSqlStatementResponse) SetStatusCode(v int32) *TerminateSqlStatementResponse {
	s.StatusCode = &v
	return s
}

func (s *TerminateSqlStatementResponse) SetBody(v *TerminateSqlStatementResponseBody) *TerminateSqlStatementResponse {
	s.Body = v
	return s
}

type UpdateLivyComputeRequest struct {
	// example:
	//
	// Token
	AuthType               *string                                         `json:"authType,omitempty" xml:"authType,omitempty"`
	AutoStartConfiguration *UpdateLivyComputeRequestAutoStartConfiguration `json:"autoStartConfiguration,omitempty" xml:"autoStartConfiguration,omitempty" type:"Struct"`
	AutoStopConfiguration  *UpdateLivyComputeRequestAutoStopConfiguration  `json:"autoStopConfiguration,omitempty" xml:"autoStopConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// 1
	CpuLimit *string `json:"cpuLimit,omitempty" xml:"cpuLimit,omitempty"`
	// example:
	//
	// esr-4.3.0 (Spark 3.5.2, Scala 2.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	EnablePublic          *bool   `json:"enablePublic,omitempty" xml:"enablePublic,omitempty"`
	// example:
	//
	// ev-cq146allhtgkulp5smk0
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// example:
	//
	// false
	Fusion *bool `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// example:
	//
	// {
	//
	//   "sparkDefaultsConf": "spark.driver.cores     1\\nspark.driver.memory    4g\\nspark.executor.cores   1\\nspark.executor.memory  4g\\n",
	//
	//   "sparkBlackListConf": "spark.driver.cores\\nspark.driver.memory",
	//
	//   "livyConf": "livy.server.session.timeout  1h\\n",
	//
	//   "livyClientConf": "livy.rsc.sql.num-rows  1000\\n"
	//
	// }
	LivyServerConf *string `json:"livyServerConf,omitempty" xml:"livyServerConf,omitempty"`
	// example:
	//
	// 0.8.0
	LivyVersion *string `json:"livyVersion,omitempty" xml:"livyVersion,omitempty"`
	// example:
	//
	// 4Gi
	MemoryLimit *string `json:"memoryLimit,omitempty" xml:"memoryLimit,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// test
	NetworkName *string `json:"networkName,omitempty" xml:"networkName,omitempty"`
	// example:
	//
	// root_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// example:
	//
	// esr-4.3.0 (Spark 3.5.2, Scala 2.12, Java Runtime)
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s UpdateLivyComputeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLivyComputeRequest) GoString() string {
	return s.String()
}

func (s *UpdateLivyComputeRequest) SetAuthType(v string) *UpdateLivyComputeRequest {
	s.AuthType = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetAutoStartConfiguration(v *UpdateLivyComputeRequestAutoStartConfiguration) *UpdateLivyComputeRequest {
	s.AutoStartConfiguration = v
	return s
}

func (s *UpdateLivyComputeRequest) SetAutoStopConfiguration(v *UpdateLivyComputeRequestAutoStopConfiguration) *UpdateLivyComputeRequest {
	s.AutoStopConfiguration = v
	return s
}

func (s *UpdateLivyComputeRequest) SetCpuLimit(v string) *UpdateLivyComputeRequest {
	s.CpuLimit = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetDisplayReleaseVersion(v string) *UpdateLivyComputeRequest {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetEnablePublic(v bool) *UpdateLivyComputeRequest {
	s.EnablePublic = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetEnvironmentId(v string) *UpdateLivyComputeRequest {
	s.EnvironmentId = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetFusion(v bool) *UpdateLivyComputeRequest {
	s.Fusion = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetLivyServerConf(v string) *UpdateLivyComputeRequest {
	s.LivyServerConf = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetLivyVersion(v string) *UpdateLivyComputeRequest {
	s.LivyVersion = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetMemoryLimit(v string) *UpdateLivyComputeRequest {
	s.MemoryLimit = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetName(v string) *UpdateLivyComputeRequest {
	s.Name = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetNetworkName(v string) *UpdateLivyComputeRequest {
	s.NetworkName = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetQueueName(v string) *UpdateLivyComputeRequest {
	s.QueueName = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetReleaseVersion(v string) *UpdateLivyComputeRequest {
	s.ReleaseVersion = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetRegionId(v string) *UpdateLivyComputeRequest {
	s.RegionId = &v
	return s
}

type UpdateLivyComputeRequestAutoStartConfiguration struct {
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s UpdateLivyComputeRequestAutoStartConfiguration) String() string {
	return tea.Prettify(s)
}

func (s UpdateLivyComputeRequestAutoStartConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateLivyComputeRequestAutoStartConfiguration) SetEnable(v bool) *UpdateLivyComputeRequestAutoStartConfiguration {
	s.Enable = &v
	return s
}

type UpdateLivyComputeRequestAutoStopConfiguration struct {
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 300
	IdleTimeoutMinutes *int64 `json:"idleTimeoutMinutes,omitempty" xml:"idleTimeoutMinutes,omitempty"`
}

func (s UpdateLivyComputeRequestAutoStopConfiguration) String() string {
	return tea.Prettify(s)
}

func (s UpdateLivyComputeRequestAutoStopConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateLivyComputeRequestAutoStopConfiguration) SetEnable(v bool) *UpdateLivyComputeRequestAutoStopConfiguration {
	s.Enable = &v
	return s
}

func (s *UpdateLivyComputeRequestAutoStopConfiguration) SetIdleTimeoutMinutes(v int64) *UpdateLivyComputeRequestAutoStopConfiguration {
	s.IdleTimeoutMinutes = &v
	return s
}

type UpdateLivyComputeResponseBody struct {
	// example:
	//
	// 1000000
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateLivyComputeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLivyComputeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLivyComputeResponseBody) SetCode(v string) *UpdateLivyComputeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLivyComputeResponseBody) SetMessage(v string) *UpdateLivyComputeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLivyComputeResponseBody) SetRequestId(v string) *UpdateLivyComputeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLivyComputeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLivyComputeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLivyComputeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLivyComputeResponse) GoString() string {
	return s.String()
}

func (s *UpdateLivyComputeResponse) SetHeaders(v map[string]*string) *UpdateLivyComputeResponse {
	s.Headers = v
	return s
}

func (s *UpdateLivyComputeResponse) SetStatusCode(v int32) *UpdateLivyComputeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLivyComputeResponse) SetBody(v *UpdateLivyComputeResponseBody) *UpdateLivyComputeResponse {
	s.Body = v
	return s
}

type UpdateProcessDefinitionWithScheduleRequest struct {
	// The email address to receive alerts.
	//
	// example:
	//
	// foo_bar@spark.alert.invalid.com
	AlertEmailAddress *string `json:"alertEmailAddress,omitempty" xml:"alertEmailAddress,omitempty"`
	// The description of the workflow.
	//
	// example:
	//
	// ods batch workflow
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The execution policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// PARALLEL
	ExecutionType *string                                                   `json:"executionType,omitempty" xml:"executionType,omitempty"`
	GlobalParams  []*UpdateProcessDefinitionWithScheduleRequestGlobalParams `json:"globalParams,omitempty" xml:"globalParams,omitempty" type:"Repeated"`
	// The name of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// ods_batch_workflow
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The code of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// SS
	ProductNamespace *string `json:"productNamespace,omitempty" xml:"productNamespace,omitempty"`
	// Specifies whether to publish the workflow.
	//
	// example:
	//
	// true
	Publish *bool `json:"publish,omitempty" xml:"publish,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The status of the workflow.
	//
	// example:
	//
	// ONLINE
	ReleaseState *string `json:"releaseState,omitempty" xml:"releaseState,omitempty"`
	// The resource queue.
	//
	// example:
	//
	// root_queue
	ResourceQueue *string `json:"resourceQueue,omitempty" xml:"resourceQueue,omitempty"`
	// The number of retries.
	//
	// example:
	//
	// 1
	RetryTimes *int32 `json:"retryTimes,omitempty" xml:"retryTimes,omitempty"`
	// The execution user.
	//
	// example:
	//
	// 113***************
	RunAs *string `json:"runAs,omitempty" xml:"runAs,omitempty"`
	// The scheduling settings.
	Schedule *UpdateProcessDefinitionWithScheduleRequestSchedule `json:"schedule,omitempty" xml:"schedule,omitempty" type:"Struct"`
	// The tags.
	Tags map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The descriptions of all nodes in the workflow.
	//
	// This parameter is required.
	TaskDefinitionJson []*UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson `json:"taskDefinitionJson,omitempty" xml:"taskDefinitionJson,omitempty" type:"Repeated"`
	// The node parallelism.
	//
	// example:
	//
	// 1
	TaskParallelism *int32 `json:"taskParallelism,omitempty" xml:"taskParallelism,omitempty"`
	// The dependencies of all nodes in the workflow. preTaskCode specifies the ID of an upstream node, and postTaskCode specifies the ID of a downstream node. The ID of each node is unique. If a node does not have an upstream node, set preTaskCode to 0.
	//
	// This parameter is required.
	TaskRelationJson []*UpdateProcessDefinitionWithScheduleRequestTaskRelationJson `json:"taskRelationJson,omitempty" xml:"taskRelationJson,omitempty" type:"Repeated"`
	// The default timeout period of the workflow.
	//
	// example:
	//
	// 300
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleRequest) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetAlertEmailAddress(v string) *UpdateProcessDefinitionWithScheduleRequest {
	s.AlertEmailAddress = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetDescription(v string) *UpdateProcessDefinitionWithScheduleRequest {
	s.Description = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetExecutionType(v string) *UpdateProcessDefinitionWithScheduleRequest {
	s.ExecutionType = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetGlobalParams(v []*UpdateProcessDefinitionWithScheduleRequestGlobalParams) *UpdateProcessDefinitionWithScheduleRequest {
	s.GlobalParams = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetName(v string) *UpdateProcessDefinitionWithScheduleRequest {
	s.Name = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetProductNamespace(v string) *UpdateProcessDefinitionWithScheduleRequest {
	s.ProductNamespace = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetPublish(v bool) *UpdateProcessDefinitionWithScheduleRequest {
	s.Publish = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetRegionId(v string) *UpdateProcessDefinitionWithScheduleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetReleaseState(v string) *UpdateProcessDefinitionWithScheduleRequest {
	s.ReleaseState = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetResourceQueue(v string) *UpdateProcessDefinitionWithScheduleRequest {
	s.ResourceQueue = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetRetryTimes(v int32) *UpdateProcessDefinitionWithScheduleRequest {
	s.RetryTimes = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetRunAs(v string) *UpdateProcessDefinitionWithScheduleRequest {
	s.RunAs = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetSchedule(v *UpdateProcessDefinitionWithScheduleRequestSchedule) *UpdateProcessDefinitionWithScheduleRequest {
	s.Schedule = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetTags(v map[string]*string) *UpdateProcessDefinitionWithScheduleRequest {
	s.Tags = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetTaskDefinitionJson(v []*UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) *UpdateProcessDefinitionWithScheduleRequest {
	s.TaskDefinitionJson = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetTaskParallelism(v int32) *UpdateProcessDefinitionWithScheduleRequest {
	s.TaskParallelism = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetTaskRelationJson(v []*UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) *UpdateProcessDefinitionWithScheduleRequest {
	s.TaskRelationJson = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetTimeout(v int32) *UpdateProcessDefinitionWithScheduleRequest {
	s.Timeout = &v
	return s
}

type UpdateProcessDefinitionWithScheduleRequestGlobalParams struct {
	Direct *string `json:"direct,omitempty" xml:"direct,omitempty"`
	Prop   *string `json:"prop,omitempty" xml:"prop,omitempty"`
	Type   *string `json:"type,omitempty" xml:"type,omitempty"`
	Value  *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleRequestGlobalParams) String() string {
	return tea.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleRequestGlobalParams) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleRequestGlobalParams) SetDirect(v string) *UpdateProcessDefinitionWithScheduleRequestGlobalParams {
	s.Direct = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestGlobalParams) SetProp(v string) *UpdateProcessDefinitionWithScheduleRequestGlobalParams {
	s.Prop = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestGlobalParams) SetType(v string) *UpdateProcessDefinitionWithScheduleRequestGlobalParams {
	s.Type = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestGlobalParams) SetValue(v string) *UpdateProcessDefinitionWithScheduleRequestGlobalParams {
	s.Value = &v
	return s
}

type UpdateProcessDefinitionWithScheduleRequestSchedule struct {
	// The CRON expression that is used for scheduling.
	//
	// example:
	//
	// 0 0 0 	- 	- ?
	Crontab *string `json:"crontab,omitempty" xml:"crontab,omitempty"`
	// The end time of the scheduling.
	//
	// example:
	//
	// 2025-12-23 16:13:27
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The start time of the scheduling.
	//
	// example:
	//
	// 2024-12-23 16:13:27
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The ID of the time zone.
	//
	// example:
	//
	// Asia/Shanghai
	TimezoneId *string `json:"timezoneId,omitempty" xml:"timezoneId,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleRequestSchedule) String() string {
	return tea.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleRequestSchedule) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleRequestSchedule) SetCrontab(v string) *UpdateProcessDefinitionWithScheduleRequestSchedule {
	s.Crontab = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestSchedule) SetEndTime(v string) *UpdateProcessDefinitionWithScheduleRequestSchedule {
	s.EndTime = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestSchedule) SetStartTime(v string) *UpdateProcessDefinitionWithScheduleRequestSchedule {
	s.StartTime = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestSchedule) SetTimezoneId(v string) *UpdateProcessDefinitionWithScheduleRequestSchedule {
	s.TimezoneId = &v
	return s
}

type UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson struct {
	// The email address to receive alerts.
	//
	// example:
	//
	// foo_bar@spark.alert.invalid.com
	AlertEmailAddress *string `json:"alertEmailAddress,omitempty" xml:"alertEmailAddress,omitempty"`
	// The node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23************
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// The node description.
	//
	// example:
	//
	// ods transform task
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to send alerts when the node fails.
	//
	// example:
	//
	// true
	FailAlertEnable *bool `json:"failAlertEnable,omitempty" xml:"failAlertEnable,omitempty"`
	// The number of retries when the node fails.
	//
	// example:
	//
	// 1
	FailRetryTimes *int32 `json:"failRetryTimes,omitempty" xml:"failRetryTimes,omitempty"`
	// The name of the job.
	//
	// This parameter is required.
	//
	// example:
	//
	// ods_transform_task
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specifies whether to send alerts when the node is started.
	//
	// example:
	//
	// true
	StartAlertEnable *bool `json:"startAlertEnable,omitempty" xml:"startAlertEnable,omitempty"`
	// The tags of the job.
	Tags map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The job parameters.
	//
	// This parameter is required.
	TaskParams *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams `json:"taskParams,omitempty" xml:"taskParams,omitempty" type:"Struct"`
	// The type of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// EMR-SERVERLESS-SPARK
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// The default timeout period of the node.
	//
	// example:
	//
	// 30
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) String() string {
	return tea.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetAlertEmailAddress(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.AlertEmailAddress = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetCode(v int64) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.Code = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetDescription(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.Description = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetFailAlertEnable(v bool) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.FailAlertEnable = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetFailRetryTimes(v int32) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.FailRetryTimes = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetName(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.Name = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetStartAlertEnable(v bool) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.StartAlertEnable = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetTags(v map[string]*string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.Tags = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetTaskParams(v *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.TaskParams = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetTaskType(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.TaskType = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetTimeout(v int32) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.Timeout = &v
	return s
}

type UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams struct {
	// The displayed version of the Spark engine.
	//
	// example:
	//
	// esr-4.0.0 (Spark 3.5.2, Scala 2.12)
	DisplaySparkVersion *string `json:"displaySparkVersion,omitempty" xml:"displaySparkVersion,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// ev-h*************
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Specifies whether to enable Fusion engine for acceleration.
	//
	// example:
	//
	// false
	Fusion      *bool                                                                                `json:"fusion,omitempty" xml:"fusion,omitempty"`
	LocalParams []*UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams `json:"localParams,omitempty" xml:"localParams,omitempty" type:"Repeated"`
	// The name of the queue on which the job runs.
	//
	// This parameter is required.
	//
	// example:
	//
	// root_queue
	ResourceQueueId *string `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	// The configurations of the Spark jobs.
	SparkConf []*UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf `json:"sparkConf,omitempty" xml:"sparkConf,omitempty" type:"Repeated"`
	// The number of driver cores of the Spark job.
	//
	// example:
	//
	// 1
	SparkDriverCores *int32 `json:"sparkDriverCores,omitempty" xml:"sparkDriverCores,omitempty"`
	// The size of driver memory of the Spark job.
	//
	// example:
	//
	// 4g
	SparkDriverMemory *int64 `json:"sparkDriverMemory,omitempty" xml:"sparkDriverMemory,omitempty"`
	// The number of executor cores of the Spark job.
	//
	// example:
	//
	// 1
	SparkExecutorCores *int32 `json:"sparkExecutorCores,omitempty" xml:"sparkExecutorCores,omitempty"`
	// The size of executor memory of the Spark job.
	//
	// example:
	//
	// 4g
	SparkExecutorMemory *int64 `json:"sparkExecutorMemory,omitempty" xml:"sparkExecutorMemory,omitempty"`
	// The level of the Spark log.
	//
	// example:
	//
	// INFO
	SparkLogLevel *string `json:"sparkLogLevel,omitempty" xml:"sparkLogLevel,omitempty"`
	// The path where the operational logs of the Spark job are stored.
	SparkLogPath *string `json:"sparkLogPath,omitempty" xml:"sparkLogPath,omitempty"`
	// The version of the Spark engine.
	//
	// example:
	//
	// esr-4.0.0 (Spark 3.5.2, Scala 2.12)
	SparkVersion *string `json:"sparkVersion,omitempty" xml:"sparkVersion,omitempty"`
	// The ID of the data development job.
	//
	// This parameter is required.
	//
	// example:
	//
	// TSK-d87******************
	TaskBizId *string `json:"taskBizId,omitempty" xml:"taskBizId,omitempty"`
	// The type of the Spark job.
	//
	// example:
	//
	// SQL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// w-d8********
	WorkspaceBizId *string `json:"workspaceBizId,omitempty" xml:"workspaceBizId,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) String() string {
	return tea.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetDisplaySparkVersion(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.DisplaySparkVersion = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetEnvironmentId(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.EnvironmentId = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetFusion(v bool) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.Fusion = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetLocalParams(v []*UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.LocalParams = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetResourceQueueId(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.ResourceQueueId = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkConf(v []*UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkConf = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkDriverCores(v int32) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkDriverCores = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkDriverMemory(v int64) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkDriverMemory = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkExecutorCores(v int32) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkExecutorCores = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkExecutorMemory(v int64) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkExecutorMemory = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkLogLevel(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkLogLevel = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkLogPath(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkLogPath = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkVersion(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkVersion = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetTaskBizId(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.TaskBizId = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetType(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.Type = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetWorkspaceBizId(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.WorkspaceBizId = &v
	return s
}

type UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams struct {
	Direct *string `json:"direct,omitempty" xml:"direct,omitempty"`
	Prop   *string `json:"prop,omitempty" xml:"prop,omitempty"`
	Type   *string `json:"type,omitempty" xml:"type,omitempty"`
	Value  *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) String() string {
	return tea.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) SetDirect(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams {
	s.Direct = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) SetProp(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams {
	s.Prop = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) SetType(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams {
	s.Type = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) SetValue(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams {
	s.Value = &v
	return s
}

type UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf struct {
	// The key of the SparkConf object.
	//
	// example:
	//
	// spark.dynamicAllocation.enabled
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the SparkConf object.
	//
	// example:
	//
	// true
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) String() string {
	return tea.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) SetKey(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf {
	s.Key = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) SetValue(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf {
	s.Value = &v
	return s
}

type UpdateProcessDefinitionWithScheduleRequestTaskRelationJson struct {
	// The name of the node topology. You can enter a workflow name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ods batch workflow
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the downstream node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 19************
	PostTaskCode *int64 `json:"postTaskCode,omitempty" xml:"postTaskCode,omitempty"`
	// The version of the downstream node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PostTaskVersion *int32 `json:"postTaskVersion,omitempty" xml:"postTaskVersion,omitempty"`
	// The ID of the upstream node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16************
	PreTaskCode *int64 `json:"preTaskCode,omitempty" xml:"preTaskCode,omitempty"`
	// The version of the upstream node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PreTaskVersion *int32 `json:"preTaskVersion,omitempty" xml:"preTaskVersion,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) String() string {
	return tea.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) SetName(v string) *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson {
	s.Name = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) SetPostTaskCode(v int64) *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson {
	s.PostTaskCode = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) SetPostTaskVersion(v int32) *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson {
	s.PostTaskVersion = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) SetPreTaskCode(v int64) *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson {
	s.PreTaskCode = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) SetPreTaskVersion(v int32) *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson {
	s.PreTaskVersion = &v
	return s
}

type UpdateProcessDefinitionWithScheduleShrinkRequest struct {
	// The email address to receive alerts.
	//
	// example:
	//
	// foo_bar@spark.alert.invalid.com
	AlertEmailAddress *string `json:"alertEmailAddress,omitempty" xml:"alertEmailAddress,omitempty"`
	// The description of the workflow.
	//
	// example:
	//
	// ods batch workflow
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The execution policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// PARALLEL
	ExecutionType      *string `json:"executionType,omitempty" xml:"executionType,omitempty"`
	GlobalParamsShrink *string `json:"globalParams,omitempty" xml:"globalParams,omitempty"`
	// The name of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// ods_batch_workflow
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The code of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// SS
	ProductNamespace *string `json:"productNamespace,omitempty" xml:"productNamespace,omitempty"`
	// Specifies whether to publish the workflow.
	//
	// example:
	//
	// true
	Publish *bool `json:"publish,omitempty" xml:"publish,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The status of the workflow.
	//
	// example:
	//
	// ONLINE
	ReleaseState *string `json:"releaseState,omitempty" xml:"releaseState,omitempty"`
	// The resource queue.
	//
	// example:
	//
	// root_queue
	ResourceQueue *string `json:"resourceQueue,omitempty" xml:"resourceQueue,omitempty"`
	// The number of retries.
	//
	// example:
	//
	// 1
	RetryTimes *int32 `json:"retryTimes,omitempty" xml:"retryTimes,omitempty"`
	// The execution user.
	//
	// example:
	//
	// 113***************
	RunAs *string `json:"runAs,omitempty" xml:"runAs,omitempty"`
	// The scheduling settings.
	ScheduleShrink *string `json:"schedule,omitempty" xml:"schedule,omitempty"`
	// The tags.
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The descriptions of all nodes in the workflow.
	//
	// This parameter is required.
	TaskDefinitionJsonShrink *string `json:"taskDefinitionJson,omitempty" xml:"taskDefinitionJson,omitempty"`
	// The node parallelism.
	//
	// example:
	//
	// 1
	TaskParallelism *int32 `json:"taskParallelism,omitempty" xml:"taskParallelism,omitempty"`
	// The dependencies of all nodes in the workflow. preTaskCode specifies the ID of an upstream node, and postTaskCode specifies the ID of a downstream node. The ID of each node is unique. If a node does not have an upstream node, set preTaskCode to 0.
	//
	// This parameter is required.
	TaskRelationJsonShrink *string `json:"taskRelationJson,omitempty" xml:"taskRelationJson,omitempty"`
	// The default timeout period of the workflow.
	//
	// example:
	//
	// 300
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetAlertEmailAddress(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.AlertEmailAddress = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetDescription(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetExecutionType(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.ExecutionType = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetGlobalParamsShrink(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.GlobalParamsShrink = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetName(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetProductNamespace(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.ProductNamespace = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetPublish(v bool) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.Publish = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetRegionId(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetReleaseState(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.ReleaseState = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetResourceQueue(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.ResourceQueue = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetRetryTimes(v int32) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.RetryTimes = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetRunAs(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.RunAs = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetScheduleShrink(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.ScheduleShrink = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetTagsShrink(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetTaskDefinitionJsonShrink(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.TaskDefinitionJsonShrink = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetTaskParallelism(v int32) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.TaskParallelism = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetTaskRelationJsonShrink(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.TaskRelationJsonShrink = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetTimeout(v int32) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.Timeout = &v
	return s
}

type UpdateProcessDefinitionWithScheduleResponseBody struct {
	// The code that is returned by the backend server.
	//
	// example:
	//
	// 1400009
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *UpdateProcessDefinitionWithScheduleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Indicates whether the request failed.
	//
	// example:
	//
	// false
	Failed *string `json:"failed,omitempty" xml:"failed,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The description of the returned code.
	//
	// example:
	//
	// No permission for resource action
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) SetCode(v int32) *UpdateProcessDefinitionWithScheduleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) SetData(v *UpdateProcessDefinitionWithScheduleResponseBodyData) *UpdateProcessDefinitionWithScheduleResponseBody {
	s.Data = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) SetFailed(v string) *UpdateProcessDefinitionWithScheduleResponseBody {
	s.Failed = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) SetHttpStatusCode(v int32) *UpdateProcessDefinitionWithScheduleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) SetMsg(v string) *UpdateProcessDefinitionWithScheduleResponseBody {
	s.Msg = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) SetRequestId(v string) *UpdateProcessDefinitionWithScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) SetSuccess(v string) *UpdateProcessDefinitionWithScheduleResponseBody {
	s.Success = &v
	return s
}

type UpdateProcessDefinitionWithScheduleResponseBodyData struct {
	// The email address to receive alerts.
	//
	// example:
	//
	// foo_bar@spark.alert.invalid.com
	AlertEmailAddress *string `json:"alertEmailAddress,omitempty" xml:"alertEmailAddress,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// alicloud_ack_one_cluster
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// The workflow ID.
	//
	// example:
	//
	// 12***********
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The time when the workflow was created.
	//
	// example:
	//
	// 2024-09-05T02:03:19Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The CRON expression that is used for scheduling.
	//
	// example:
	//
	// 0 0 0 	- 	- ?
	Crontab *string `json:"crontab,omitempty" xml:"crontab,omitempty"`
	// The node description.
	//
	// example:
	//
	// 1
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The end of the end time range.
	//
	// example:
	//
	// 1710432000000
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The execution policy.
	//
	// example:
	//
	// SERIAL
	ExecutionType *string `json:"executionType,omitempty" xml:"executionType,omitempty"`
	// The serial number of the workflow.
	//
	// example:
	//
	// 123223
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the workflow.
	//
	// example:
	//
	// ods_batch_workflow
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the project to which the workflow belongs.
	//
	// example:
	//
	// w-********
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// The status of the workflow.
	//
	// example:
	//
	// ONLINE
	ReleaseState *string `json:"releaseState,omitempty" xml:"releaseState,omitempty"`
	// The start time of the scheduling.
	//
	// example:
	//
	// 0
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The ID of the time zone.
	//
	// example:
	//
	// Asia/Shanghai
	TimezoneId *string `json:"timezoneId,omitempty" xml:"timezoneId,omitempty"`
	// The time when the workflow was updated.
	//
	// example:
	//
	// 2024-03-05T06:24:27Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The ID of the user that is used to initiate a scheduling.
	//
	// example:
	//
	// 113*********
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The name of the user that is used to initiate a scheduling.
	//
	// example:
	//
	// w-********
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
	// The hash code of the version.
	//
	// example:
	//
	// dwerf*********
	VersionHashCode *string `json:"versionHashCode,omitempty" xml:"versionHashCode,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetAlertEmailAddress(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.AlertEmailAddress = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetBizId(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.BizId = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetCode(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.Code = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetCreateTime(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetCrontab(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.Crontab = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetDescription(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetEndTime(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetExecutionType(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.ExecutionType = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetId(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetName(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetProjectName(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetReleaseState(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.ReleaseState = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetStartTime(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetTimezoneId(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.TimezoneId = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetUpdateTime(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetUserId(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.UserId = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetUserName(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.UserName = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetVersion(v int32) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.Version = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetVersionHashCode(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.VersionHashCode = &v
	return s
}

type UpdateProcessDefinitionWithScheduleResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProcessDefinitionWithScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleResponse) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleResponse) SetHeaders(v map[string]*string) *UpdateProcessDefinitionWithScheduleResponse {
	s.Headers = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponse) SetStatusCode(v int32) *UpdateProcessDefinitionWithScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponse) SetBody(v *UpdateProcessDefinitionWithScheduleResponseBody) *UpdateProcessDefinitionWithScheduleResponse {
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
// Adds a RAM user or RAM role to a workspace as a member.
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
// Adds a RAM user or RAM role to a workspace as a member.
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
// Terminates a Spark job.
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
// Terminates a Spark job.
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
// 创建Livy compute
//
// @param request - CreateLivyComputeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLivyComputeResponse
func (client *Client) CreateLivyComputeWithOptions(workspaceBizId *string, request *CreateLivyComputeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateLivyComputeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthType)) {
		body["authType"] = request.AuthType
	}

	if !tea.BoolValue(util.IsUnset(request.AutoStartConfiguration)) {
		body["autoStartConfiguration"] = request.AutoStartConfiguration
	}

	if !tea.BoolValue(util.IsUnset(request.AutoStopConfiguration)) {
		body["autoStopConfiguration"] = request.AutoStopConfiguration
	}

	if !tea.BoolValue(util.IsUnset(request.CpuLimit)) {
		body["cpuLimit"] = request.CpuLimit
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayReleaseVersion)) {
		body["displayReleaseVersion"] = request.DisplayReleaseVersion
	}

	if !tea.BoolValue(util.IsUnset(request.EnablePublic)) {
		body["enablePublic"] = request.EnablePublic
	}

	if !tea.BoolValue(util.IsUnset(request.EnvironmentId)) {
		body["environmentId"] = request.EnvironmentId
	}

	if !tea.BoolValue(util.IsUnset(request.Fusion)) {
		body["fusion"] = request.Fusion
	}

	if !tea.BoolValue(util.IsUnset(request.LivyServerConf)) {
		body["livyServerConf"] = request.LivyServerConf
	}

	if !tea.BoolValue(util.IsUnset(request.LivyVersion)) {
		body["livyVersion"] = request.LivyVersion
	}

	if !tea.BoolValue(util.IsUnset(request.MemoryLimit)) {
		body["memoryLimit"] = request.MemoryLimit
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkName)) {
		body["networkName"] = request.NetworkName
	}

	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		body["queueName"] = request.QueueName
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseVersion)) {
		body["releaseVersion"] = request.ReleaseVersion
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLivyCompute"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/interactive/v1/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceBizId)) + "/livycompute"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLivyComputeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Livy compute
//
// @param request - CreateLivyComputeRequest
//
// @return CreateLivyComputeResponse
func (client *Client) CreateLivyCompute(workspaceBizId *string, request *CreateLivyComputeRequest) (_result *CreateLivyComputeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLivyComputeResponse{}
	_body, _err := client.CreateLivyComputeWithOptions(workspaceBizId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Livy Compute的token
//
// @param request - CreateLivyComputeTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLivyComputeTokenResponse
func (client *Client) CreateLivyComputeTokenWithOptions(workspaceBizId *string, livyComputeId *string, request *CreateLivyComputeTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateLivyComputeTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoExpireConfiguration)) {
		body["autoExpireConfiguration"] = request.AutoExpireConfiguration
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLivyComputeToken"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/interactive/v1/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceBizId)) + "/livycompute/" + tea.StringValue(openapiutil.GetEncodeParam(livyComputeId)) + "/token"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLivyComputeTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Livy Compute的token
//
// @param request - CreateLivyComputeTokenRequest
//
// @return CreateLivyComputeTokenResponse
func (client *Client) CreateLivyComputeToken(workspaceBizId *string, livyComputeId *string, request *CreateLivyComputeTokenRequest) (_result *CreateLivyComputeTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLivyComputeTokenResponse{}
	_body, _err := client.CreateLivyComputeTokenWithOptions(workspaceBizId, livyComputeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a workflow.
//
// @param tmpReq - CreateProcessDefinitionWithScheduleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProcessDefinitionWithScheduleResponse
func (client *Client) CreateProcessDefinitionWithScheduleWithOptions(bizId *string, tmpReq *CreateProcessDefinitionWithScheduleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProcessDefinitionWithScheduleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateProcessDefinitionWithScheduleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.GlobalParams)) {
		request.GlobalParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GlobalParams, tea.String("globalParams"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Schedule)) {
		request.ScheduleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Schedule, tea.String("schedule"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("tags"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TaskDefinitionJson)) {
		request.TaskDefinitionJsonShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskDefinitionJson, tea.String("taskDefinitionJson"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TaskRelationJson)) {
		request.TaskRelationJsonShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskRelationJson, tea.String("taskRelationJson"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertEmailAddress)) {
		query["alertEmailAddress"] = request.AlertEmailAddress
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionType)) {
		query["executionType"] = request.ExecutionType
	}

	if !tea.BoolValue(util.IsUnset(request.GlobalParamsShrink)) {
		query["globalParams"] = request.GlobalParamsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProductNamespace)) {
		query["productNamespace"] = request.ProductNamespace
	}

	if !tea.BoolValue(util.IsUnset(request.Publish)) {
		query["publish"] = request.Publish
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceQueue)) {
		query["resourceQueue"] = request.ResourceQueue
	}

	if !tea.BoolValue(util.IsUnset(request.RetryTimes)) {
		query["retryTimes"] = request.RetryTimes
	}

	if !tea.BoolValue(util.IsUnset(request.RunAs)) {
		query["runAs"] = request.RunAs
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleShrink)) {
		query["schedule"] = request.ScheduleShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskDefinitionJsonShrink)) {
		query["taskDefinitionJson"] = request.TaskDefinitionJsonShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskParallelism)) {
		query["taskParallelism"] = request.TaskParallelism
	}

	if !tea.BoolValue(util.IsUnset(request.TaskRelationJsonShrink)) {
		query["taskRelationJson"] = request.TaskRelationJsonShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		query["timeout"] = request.Timeout
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProcessDefinitionWithSchedule"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/projects/" + tea.StringValue(openapiutil.GetEncodeParam(bizId)) + "/process-definition"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProcessDefinitionWithScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workflow.
//
// @param request - CreateProcessDefinitionWithScheduleRequest
//
// @return CreateProcessDefinitionWithScheduleResponse
func (client *Client) CreateProcessDefinitionWithSchedule(bizId *string, request *CreateProcessDefinitionWithScheduleRequest) (_result *CreateProcessDefinitionWithScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProcessDefinitionWithScheduleResponse{}
	_body, _err := client.CreateProcessDefinitionWithScheduleWithOptions(bizId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a session.
//
// @param request - CreateSessionClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSessionClusterResponse
func (client *Client) CreateSessionClusterWithOptions(workspaceId *string, request *CreateSessionClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSessionClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationConfigs)) {
		body["applicationConfigs"] = request.ApplicationConfigs
	}

	if !tea.BoolValue(util.IsUnset(request.AutoStartConfiguration)) {
		body["autoStartConfiguration"] = request.AutoStartConfiguration
	}

	if !tea.BoolValue(util.IsUnset(request.AutoStopConfiguration)) {
		body["autoStopConfiguration"] = request.AutoStopConfiguration
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayReleaseVersion)) {
		body["displayReleaseVersion"] = request.DisplayReleaseVersion
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		body["envId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Fusion)) {
		body["fusion"] = request.Fusion
	}

	if !tea.BoolValue(util.IsUnset(request.Kind)) {
		body["kind"] = request.Kind
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PublicEndpointEnabled)) {
		body["publicEndpointEnabled"] = request.PublicEndpointEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		body["queueName"] = request.QueueName
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseVersion)) {
		body["releaseVersion"] = request.ReleaseVersion
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSessionCluster"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/sessionClusters"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSessionClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a session.
//
// @param request - CreateSessionClusterRequest
//
// @return CreateSessionClusterResponse
func (client *Client) CreateSessionCluster(workspaceId *string, request *CreateSessionClusterRequest) (_result *CreateSessionClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSessionClusterResponse{}
	_body, _err := client.CreateSessionClusterWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an SQL query task.
//
// @param request - CreateSqlStatementRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSqlStatementResponse
func (client *Client) CreateSqlStatementWithOptions(workspaceId *string, request *CreateSqlStatementRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSqlStatementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CodeContent)) {
		body["codeContent"] = request.CodeContent
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultCatalog)) {
		body["defaultCatalog"] = request.DefaultCatalog
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultDatabase)) {
		body["defaultDatabase"] = request.DefaultDatabase
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.SqlComputeId)) {
		body["sqlComputeId"] = request.SqlComputeId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSqlStatement"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/interactive/v1/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/statement"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSqlStatementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an SQL query task.
//
// @param request - CreateSqlStatementRequest
//
// @return CreateSqlStatementResponse
func (client *Client) CreateSqlStatement(workspaceId *string, request *CreateSqlStatementRequest) (_result *CreateSqlStatementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSqlStatementResponse{}
	_body, _err := client.CreateSqlStatementWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a workspace.
//
// @param request - CreateWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkspaceResponse
func (client *Client) CreateWorkspaceWithOptions(request *CreateWorkspaceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		body["autoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewPeriod)) {
		body["autoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewPeriodUnit)) {
		body["autoRenewPeriodUnit"] = request.AutoRenewPeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.AutoStartSessionCluster)) {
		body["autoStartSessionCluster"] = request.AutoStartSessionCluster
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DlfCatalogId)) {
		body["dlfCatalogId"] = request.DlfCatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DlfType)) {
		body["dlfType"] = request.DlfType
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucket)) {
		body["ossBucket"] = request.OssBucket
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentDurationUnit)) {
		body["paymentDurationUnit"] = request.PaymentDurationUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentType)) {
		body["paymentType"] = request.PaymentType
	}

	if !tea.BoolValue(util.IsUnset(request.RamRoleName)) {
		body["ramRoleName"] = request.RamRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseType)) {
		body["releaseType"] = request.ReleaseType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceSpec)) {
		body["resourceSpec"] = request.ResourceSpec
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceName)) {
		body["workspaceName"] = request.WorkspaceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWorkspace"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workspace.
//
// @param request - CreateWorkspaceRequest
//
// @return CreateWorkspaceResponse
func (client *Client) CreateWorkspace(request *CreateWorkspaceRequest) (_result *CreateWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CreateWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除livy compute
//
// @param request - DeleteLivyComputeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLivyComputeResponse
func (client *Client) DeleteLivyComputeWithOptions(workspaceBizId *string, livyComputeId *string, request *DeleteLivyComputeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteLivyComputeResponse, _err error) {
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
		Action:      tea.String("DeleteLivyCompute"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/interactive/v1/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceBizId)) + "/livycompute/" + tea.StringValue(openapiutil.GetEncodeParam(livyComputeId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLivyComputeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除livy compute
//
// @param request - DeleteLivyComputeRequest
//
// @return DeleteLivyComputeResponse
func (client *Client) DeleteLivyCompute(workspaceBizId *string, livyComputeId *string, request *DeleteLivyComputeRequest) (_result *DeleteLivyComputeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLivyComputeResponse{}
	_body, _err := client.DeleteLivyComputeWithOptions(workspaceBizId, livyComputeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除Livy Compute的token
//
// @param request - DeleteLivyComputeTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLivyComputeTokenResponse
func (client *Client) DeleteLivyComputeTokenWithOptions(workspaceBizId *string, livyComputeId *string, tokenId *string, request *DeleteLivyComputeTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteLivyComputeTokenResponse, _err error) {
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
		Action:      tea.String("DeleteLivyComputeToken"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/interactive/v1/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceBizId)) + "/livycompute/" + tea.StringValue(openapiutil.GetEncodeParam(livyComputeId)) + "/token/" + tea.StringValue(openapiutil.GetEncodeParam(tokenId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLivyComputeTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Livy Compute的token
//
// @param request - DeleteLivyComputeTokenRequest
//
// @return DeleteLivyComputeTokenResponse
func (client *Client) DeleteLivyComputeToken(workspaceBizId *string, livyComputeId *string, tokenId *string, request *DeleteLivyComputeTokenRequest) (_result *DeleteLivyComputeTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLivyComputeTokenResponse{}
	_body, _err := client.DeleteLivyComputeTokenWithOptions(workspaceBizId, livyComputeId, tokenId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the queue of a workspace.
//
// @param request - EditWorkspaceQueueRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditWorkspaceQueueResponse
func (client *Client) EditWorkspaceQueueWithOptions(request *EditWorkspaceQueueRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EditWorkspaceQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Environments)) {
		body["environments"] = request.Environments
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceSpec)) {
		body["resourceSpec"] = request.ResourceSpec
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["workspaceId"] = request.WorkspaceId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceQueueName)) {
		body["workspaceQueueName"] = request.WorkspaceQueueName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EditWorkspaceQueue"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/queues/action/edit"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &EditWorkspaceQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the queue of a workspace.
//
// @param request - EditWorkspaceQueueRequest
//
// @return EditWorkspaceQueueResponse
func (client *Client) EditWorkspaceQueue(request *EditWorkspaceQueueRequest) (_result *EditWorkspaceQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EditWorkspaceQueueResponse{}
	_body, _err := client.EditWorkspaceQueueWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of CU-hours consumed by a queue during a specified cycle.
//
// @param request - GetCuHoursRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCuHoursResponse
func (client *Client) GetCuHoursWithOptions(workspaceId *string, queue *string, request *GetCuHoursRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCuHoursResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCuHours"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/metric/cuHours/" + tea.StringValue(openapiutil.GetEncodeParam(queue))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCuHoursResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of CU-hours consumed by a queue during a specified cycle.
//
// @param request - GetCuHoursRequest
//
// @return GetCuHoursResponse
func (client *Client) GetCuHours(workspaceId *string, queue *string, request *GetCuHoursRequest) (_result *GetCuHoursResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCuHoursResponse{}
	_body, _err := client.GetCuHoursWithOptions(workspaceId, queue, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains job analysis information on E-MapReduce (EMR) Doctor.
//
// @param request - GetDoctorApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDoctorApplicationResponse
func (client *Client) GetDoctorApplicationWithOptions(workspaceId *string, runId *string, request *GetDoctorApplicationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDoctorApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Locale)) {
		query["locale"] = request.Locale
	}

	if !tea.BoolValue(util.IsUnset(request.QueryTime)) {
		query["queryTime"] = request.QueryTime
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDoctorApplication"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/runs/" + tea.StringValue(openapiutil.GetEncodeParam(runId)) + "/action/getDoctorApplication"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDoctorApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains job analysis information on E-MapReduce (EMR) Doctor.
//
// @param request - GetDoctorApplicationRequest
//
// @return GetDoctorApplicationResponse
func (client *Client) GetDoctorApplication(workspaceId *string, runId *string, request *GetDoctorApplicationRequest) (_result *GetDoctorApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDoctorApplicationResponse{}
	_body, _err := client.GetDoctorApplicationWithOptions(workspaceId, runId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the job details.
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
// Obtain the job details.
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
// 获取livy compute
//
// @param request - GetLivyComputeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLivyComputeResponse
func (client *Client) GetLivyComputeWithOptions(workspaceBizId *string, livyComputeId *string, request *GetLivyComputeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLivyComputeResponse, _err error) {
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
		Action:      tea.String("GetLivyCompute"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/interactive/v1/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceBizId)) + "/livycompute/" + tea.StringValue(openapiutil.GetEncodeParam(livyComputeId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLivyComputeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取livy compute
//
// @param request - GetLivyComputeRequest
//
// @return GetLivyComputeResponse
func (client *Client) GetLivyCompute(workspaceBizId *string, livyComputeId *string, request *GetLivyComputeRequest) (_result *GetLivyComputeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLivyComputeResponse{}
	_body, _err := client.GetLivyComputeWithOptions(workspaceBizId, livyComputeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取livy compute token
//
// @param request - GetLivyComputeTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLivyComputeTokenResponse
func (client *Client) GetLivyComputeTokenWithOptions(workspaceBizId *string, livyComputeId *string, tokenId *string, request *GetLivyComputeTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLivyComputeTokenResponse, _err error) {
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
		Action:      tea.String("GetLivyComputeToken"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/interactive/v1/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceBizId)) + "/livycompute/" + tea.StringValue(openapiutil.GetEncodeParam(livyComputeId)) + "/token/" + tea.StringValue(openapiutil.GetEncodeParam(tokenId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLivyComputeTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取livy compute token
//
// @param request - GetLivyComputeTokenRequest
//
// @return GetLivyComputeTokenResponse
func (client *Client) GetLivyComputeToken(workspaceBizId *string, livyComputeId *string, tokenId *string, request *GetLivyComputeTokenRequest) (_result *GetLivyComputeTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLivyComputeTokenResponse{}
	_body, _err := client.GetLivyComputeTokenWithOptions(workspaceBizId, livyComputeId, tokenId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a session.
//
// @param request - GetSessionClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSessionClusterResponse
func (client *Client) GetSessionClusterWithOptions(workspaceId *string, sessionClusterId *string, request *GetSessionClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSessionClusterResponse, _err error) {
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
		Action:      tea.String("GetSessionCluster"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/sessionClusters/" + tea.StringValue(openapiutil.GetEncodeParam(sessionClusterId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSessionClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a session.
//
// @param request - GetSessionClusterRequest
//
// @return GetSessionClusterResponse
func (client *Client) GetSessionCluster(workspaceId *string, sessionClusterId *string, request *GetSessionClusterRequest) (_result *GetSessionClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSessionClusterResponse{}
	_body, _err := client.GetSessionClusterWithOptions(workspaceId, sessionClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of an SQL query task.
//
// @param request - GetSqlStatementRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSqlStatementResponse
func (client *Client) GetSqlStatementWithOptions(workspaceId *string, statementId *string, request *GetSqlStatementRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSqlStatementResponse, _err error) {
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
		Action:      tea.String("GetSqlStatement"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/interactive/v1/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/statement/" + tea.StringValue(openapiutil.GetEncodeParam(statementId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSqlStatementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of an SQL query task.
//
// @param request - GetSqlStatementRequest
//
// @return GetSqlStatementResponse
func (client *Client) GetSqlStatement(workspaceId *string, statementId *string, request *GetSqlStatementRequest) (_result *GetSqlStatementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSqlStatementResponse{}
	_body, _err := client.GetSqlStatementWithOptions(workspaceId, statementId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries task templates.
//
// @param request - GetTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateResponse
func (client *Client) GetTemplateWithOptions(workspaceBizId *string, request *GetTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateBizId)) {
		query["templateBizId"] = request.TemplateBizId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["templateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTemplate"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/interactive/v1/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceBizId)) + "/template"),
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

// Summary:
//
// Queries task templates.
//
// @param request - GetTemplateRequest
//
// @return GetTemplateResponse
func (client *Client) GetTemplate(workspaceBizId *string, request *GetTemplateRequest) (_result *GetTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(workspaceBizId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Assigns a specified role to users.
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
// Assigns a specified role to users.
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
// Queries a list of Spark jobs.
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

	if !tea.BoolValue(util.IsUnset(request.IsWorkflow)) {
		query["isWorkflow"] = request.IsWorkflow
	}

	if !tea.BoolValue(util.IsUnset(request.JobRunDeploymentId)) {
		query["jobRunDeploymentId"] = request.JobRunDeploymentId
	}

	if !tea.BoolValue(util.IsUnset(request.JobRunId)) {
		query["jobRunId"] = request.JobRunId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.MinDuration)) {
		query["minDuration"] = request.MinDuration
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
// Queries a list of Spark jobs.
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
// # ListKyuubiServices
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKyuubiServicesResponse
func (client *Client) ListKyuubiServicesWithOptions(workspaceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListKyuubiServicesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListKyuubiServices"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/kyuubi/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListKyuubiServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListKyuubiServices
//
// @return ListKyuubiServicesResponse
func (client *Client) ListKyuubiServices(workspaceId *string) (_result *ListKyuubiServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListKyuubiServicesResponse{}
	_body, _err := client.ListKyuubiServicesWithOptions(workspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the applications that are submitted by using a Kyuubi gateway.
//
// @param tmpReq - ListKyuubiSparkApplicationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKyuubiSparkApplicationsResponse
func (client *Client) ListKyuubiSparkApplicationsWithOptions(workspaceId *string, kyuubiServiceId *string, tmpReq *ListKyuubiSparkApplicationsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListKyuubiSparkApplicationsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListKyuubiSparkApplicationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OrderBy)) {
		request.OrderByShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderBy, tea.String("orderBy"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.StartTime)) {
		request.StartTimeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StartTime, tea.String("startTime"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["applicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["applicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.MinDuration)) {
		query["minDuration"] = request.MinDuration
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrderByShrink)) {
		query["orderBy"] = request.OrderByShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceQueueId)) {
		query["resourceQueueId"] = request.ResourceQueueId
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeShrink)) {
		query["startTime"] = request.StartTimeShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListKyuubiSparkApplications"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/kyuubi/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(kyuubiServiceId)) + "/applications"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListKyuubiSparkApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the applications that are submitted by using a Kyuubi gateway.
//
// @param request - ListKyuubiSparkApplicationsRequest
//
// @return ListKyuubiSparkApplicationsResponse
func (client *Client) ListKyuubiSparkApplications(workspaceId *string, kyuubiServiceId *string, request *ListKyuubiSparkApplicationsRequest) (_result *ListKyuubiSparkApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListKyuubiSparkApplicationsResponse{}
	_body, _err := client.ListKyuubiSparkApplicationsWithOptions(workspaceId, kyuubiServiceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出compute的token
//
// @param request - ListKyuubiTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKyuubiTokenResponse
func (client *Client) ListKyuubiTokenWithOptions(workspaceId *string, kyuubiServiceId *string, request *ListKyuubiTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListKyuubiTokenResponse, _err error) {
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
		Action:      tea.String("ListKyuubiToken"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/kyuubiService/" + tea.StringValue(openapiutil.GetEncodeParam(kyuubiServiceId)) + "/token"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListKyuubiTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出compute的token
//
// @param request - ListKyuubiTokenRequest
//
// @return ListKyuubiTokenResponse
func (client *Client) ListKyuubiToken(workspaceId *string, kyuubiServiceId *string, request *ListKyuubiTokenRequest) (_result *ListKyuubiTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListKyuubiTokenResponse{}
	_body, _err := client.ListKyuubiTokenWithOptions(workspaceId, kyuubiServiceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出livy compute
//
// @param request - ListLivyComputeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLivyComputeResponse
func (client *Client) ListLivyComputeWithOptions(workspaceBizId *string, request *ListLivyComputeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListLivyComputeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvironmentId)) {
		query["environmentId"] = request.EnvironmentId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLivyCompute"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/interactive/v1/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceBizId)) + "/livycompute"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLivyComputeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出livy compute
//
// @param request - ListLivyComputeRequest
//
// @return ListLivyComputeResponse
func (client *Client) ListLivyCompute(workspaceBizId *string, request *ListLivyComputeRequest) (_result *ListLivyComputeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLivyComputeResponse{}
	_body, _err := client.ListLivyComputeWithOptions(workspaceBizId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出livy compute token
//
// @param request - ListLivyComputeTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLivyComputeTokenResponse
func (client *Client) ListLivyComputeTokenWithOptions(workspaceBizId *string, livyComputeId *string, request *ListLivyComputeTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListLivyComputeTokenResponse, _err error) {
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
		Action:      tea.String("ListLivyComputeToken"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/interactive/v1/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceBizId)) + "/livycompute/" + tea.StringValue(openapiutil.GetEncodeParam(livyComputeId)) + "/token"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLivyComputeTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出livy compute token
//
// @param request - ListLivyComputeTokenRequest
//
// @return ListLivyComputeTokenResponse
func (client *Client) ListLivyComputeToken(workspaceBizId *string, livyComputeId *string, request *ListLivyComputeTokenRequest) (_result *ListLivyComputeTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLivyComputeTokenResponse{}
	_body, _err := client.ListLivyComputeTokenWithOptions(workspaceBizId, livyComputeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Log Content
//
// @param request - ListLogContentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogContentsResponse
func (client *Client) ListLogContentsWithOptions(workspaceId *string, request *ListLogContentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListLogContentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.Length)) {
		query["length"] = request.Length
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLogContents"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/action/listLogContents"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLogContentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Log Content
//
// @param request - ListLogContentsRequest
//
// @return ListLogContentsResponse
func (client *Client) ListLogContents(workspaceId *string, request *ListLogContentsRequest) (_result *ListLogContentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLogContentsResponse{}
	_body, _err := client.ListLogContentsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of published versions of E-MapReduce (EMR) Serverless Spark.
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

	if !tea.BoolValue(util.IsUnset(request.ServiceFilter)) {
		query["serviceFilter"] = request.ServiceFilter
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
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
// Queries the list of published versions of E-MapReduce (EMR) Serverless Spark.
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
// Queries the list of sessions.
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
	if !tea.BoolValue(util.IsUnset(request.Kind)) {
		query["kind"] = request.Kind
	}

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
// Queries the list of sessions.
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
// Queries the list of queues in a Spark workspace.
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
// Queries the list of queues in a Spark workspace.
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
// Queries a list of workspaces.
//
// @param tmpReq - ListWorkspacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspacesWithOptions(tmpReq *ListWorkspacesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListWorkspacesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("tag"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["tag"] = request.TagShrink
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
// Queries a list of workspaces.
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
// 更新Livy Compute的token
//
// @param request - RefreshLivyComputeTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshLivyComputeTokenResponse
func (client *Client) RefreshLivyComputeTokenWithOptions(workspaceBizId *string, livyComputeId *string, tokenId *string, request *RefreshLivyComputeTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RefreshLivyComputeTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoExpireConfiguration)) {
		body["autoExpireConfiguration"] = request.AutoExpireConfiguration
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshLivyComputeToken"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/interactive/v1/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceBizId)) + "/livycompute/" + tea.StringValue(openapiutil.GetEncodeParam(livyComputeId)) + "/token/" + tea.StringValue(openapiutil.GetEncodeParam(tokenId))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshLivyComputeTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Livy Compute的token
//
// @param request - RefreshLivyComputeTokenRequest
//
// @return RefreshLivyComputeTokenResponse
func (client *Client) RefreshLivyComputeToken(workspaceBizId *string, livyComputeId *string, tokenId *string, request *RefreshLivyComputeTokenRequest) (_result *RefreshLivyComputeTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RefreshLivyComputeTokenResponse{}
	_body, _err := client.RefreshLivyComputeTokenWithOptions(workspaceBizId, livyComputeId, tokenId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a Spark job.
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

	if !tea.BoolValue(util.IsUnset(request.DisplayReleaseVersion)) {
		body["displayReleaseVersion"] = request.DisplayReleaseVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionTimeoutSeconds)) {
		body["executionTimeoutSeconds"] = request.ExecutionTimeoutSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.Fusion)) {
		body["fusion"] = request.Fusion
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
// Starts a Spark job.
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

// Summary:
//
// 启动livy compute
//
// @param request - StartLivyComputeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartLivyComputeResponse
func (client *Client) StartLivyComputeWithOptions(workspaceBizId *string, livyComputeId *string, request *StartLivyComputeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartLivyComputeResponse, _err error) {
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
		Action:      tea.String("StartLivyCompute"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/interactive/v1/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceBizId)) + "/livycompute/" + tea.StringValue(openapiutil.GetEncodeParam(livyComputeId)) + "/start"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartLivyComputeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动livy compute
//
// @param request - StartLivyComputeRequest
//
// @return StartLivyComputeResponse
func (client *Client) StartLivyCompute(workspaceBizId *string, livyComputeId *string, request *StartLivyComputeRequest) (_result *StartLivyComputeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartLivyComputeResponse{}
	_body, _err := client.StartLivyComputeWithOptions(workspaceBizId, livyComputeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Manually runs a workflow.
//
// @param request - StartProcessInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartProcessInstanceResponse
func (client *Client) StartProcessInstanceWithOptions(bizId *string, request *StartProcessInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartProcessInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		query["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.Comments)) {
		query["comments"] = request.Comments
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IsProd)) {
		query["isProd"] = request.IsProd
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessDefinitionCode)) {
		query["processDefinitionCode"] = request.ProcessDefinitionCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductNamespace)) {
		query["productNamespace"] = request.ProductNamespace
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuntimeQueue)) {
		query["runtimeQueue"] = request.RuntimeQueue
	}

	if !tea.BoolValue(util.IsUnset(request.VersionHashCode)) {
		query["versionHashCode"] = request.VersionHashCode
	}

	if !tea.BoolValue(util.IsUnset(request.VersionNumber)) {
		query["versionNumber"] = request.VersionNumber
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartProcessInstance"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/projects/" + tea.StringValue(openapiutil.GetEncodeParam(bizId)) + "/executors/start-process-instance"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartProcessInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manually runs a workflow.
//
// @param request - StartProcessInstanceRequest
//
// @return StartProcessInstanceResponse
func (client *Client) StartProcessInstance(bizId *string, request *StartProcessInstanceRequest) (_result *StartProcessInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartProcessInstanceResponse{}
	_body, _err := client.StartProcessInstanceWithOptions(bizId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a session.
//
// @param request - StartSessionClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSessionClusterResponse
func (client *Client) StartSessionClusterWithOptions(workspaceId *string, request *StartSessionClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartSessionClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		body["queueName"] = request.QueueName
	}

	if !tea.BoolValue(util.IsUnset(request.SessionClusterId)) {
		body["sessionClusterId"] = request.SessionClusterId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartSessionCluster"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/sessionClusters/action/startSessionCluster"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartSessionClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a session.
//
// @param request - StartSessionClusterRequest
//
// @return StartSessionClusterResponse
func (client *Client) StartSessionCluster(workspaceId *string, request *StartSessionClusterRequest) (_result *StartSessionClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartSessionClusterResponse{}
	_body, _err := client.StartSessionClusterWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止livy compute
//
// @param request - StopLivyComputeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopLivyComputeResponse
func (client *Client) StopLivyComputeWithOptions(workspaceBizId *string, livyComputeId *string, request *StopLivyComputeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopLivyComputeResponse, _err error) {
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
		Action:      tea.String("StopLivyCompute"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/interactive/v1/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceBizId)) + "/livycompute/" + tea.StringValue(openapiutil.GetEncodeParam(livyComputeId)) + "/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopLivyComputeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止livy compute
//
// @param request - StopLivyComputeRequest
//
// @return StopLivyComputeResponse
func (client *Client) StopLivyCompute(workspaceBizId *string, livyComputeId *string, request *StopLivyComputeRequest) (_result *StopLivyComputeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopLivyComputeResponse{}
	_body, _err := client.StopLivyComputeWithOptions(workspaceBizId, livyComputeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a session.
//
// @param request - StopSessionClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopSessionClusterResponse
func (client *Client) StopSessionClusterWithOptions(workspaceId *string, request *StopSessionClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopSessionClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		body["queueName"] = request.QueueName
	}

	if !tea.BoolValue(util.IsUnset(request.SessionClusterId)) {
		body["sessionClusterId"] = request.SessionClusterId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopSessionCluster"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/sessionClusters/action/stopSessionCluster"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopSessionClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a session.
//
// @param request - StopSessionClusterRequest
//
// @return StopSessionClusterResponse
func (client *Client) StopSessionCluster(workspaceId *string, request *StopSessionClusterRequest) (_result *StopSessionClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopSessionClusterResponse{}
	_body, _err := client.StopSessionClusterWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Terminates an SQL query task.
//
// @param request - TerminateSqlStatementRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TerminateSqlStatementResponse
func (client *Client) TerminateSqlStatementWithOptions(workspaceId *string, statementId *string, request *TerminateSqlStatementRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TerminateSqlStatementResponse, _err error) {
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
		Action:      tea.String("TerminateSqlStatement"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/interactive/v1/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/statement/" + tea.StringValue(openapiutil.GetEncodeParam(statementId)) + "/terminate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TerminateSqlStatementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Terminates an SQL query task.
//
// @param request - TerminateSqlStatementRequest
//
// @return TerminateSqlStatementResponse
func (client *Client) TerminateSqlStatement(workspaceId *string, statementId *string, request *TerminateSqlStatementRequest) (_result *TerminateSqlStatementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TerminateSqlStatementResponse{}
	_body, _err := client.TerminateSqlStatementWithOptions(workspaceId, statementId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新livy compute
//
// @param request - UpdateLivyComputeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLivyComputeResponse
func (client *Client) UpdateLivyComputeWithOptions(workspaceBizId *string, livyComputeId *string, request *UpdateLivyComputeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateLivyComputeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthType)) {
		body["authType"] = request.AuthType
	}

	if !tea.BoolValue(util.IsUnset(request.AutoStartConfiguration)) {
		body["autoStartConfiguration"] = request.AutoStartConfiguration
	}

	if !tea.BoolValue(util.IsUnset(request.AutoStopConfiguration)) {
		body["autoStopConfiguration"] = request.AutoStopConfiguration
	}

	if !tea.BoolValue(util.IsUnset(request.CpuLimit)) {
		body["cpuLimit"] = request.CpuLimit
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayReleaseVersion)) {
		body["displayReleaseVersion"] = request.DisplayReleaseVersion
	}

	if !tea.BoolValue(util.IsUnset(request.EnablePublic)) {
		body["enablePublic"] = request.EnablePublic
	}

	if !tea.BoolValue(util.IsUnset(request.EnvironmentId)) {
		body["environmentId"] = request.EnvironmentId
	}

	if !tea.BoolValue(util.IsUnset(request.Fusion)) {
		body["fusion"] = request.Fusion
	}

	if !tea.BoolValue(util.IsUnset(request.LivyServerConf)) {
		body["livyServerConf"] = request.LivyServerConf
	}

	if !tea.BoolValue(util.IsUnset(request.LivyVersion)) {
		body["livyVersion"] = request.LivyVersion
	}

	if !tea.BoolValue(util.IsUnset(request.MemoryLimit)) {
		body["memoryLimit"] = request.MemoryLimit
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkName)) {
		body["networkName"] = request.NetworkName
	}

	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		body["queueName"] = request.QueueName
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseVersion)) {
		body["releaseVersion"] = request.ReleaseVersion
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLivyCompute"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/interactive/v1/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceBizId)) + "/livycompute/" + tea.StringValue(openapiutil.GetEncodeParam(livyComputeId))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLivyComputeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新livy compute
//
// @param request - UpdateLivyComputeRequest
//
// @return UpdateLivyComputeResponse
func (client *Client) UpdateLivyCompute(workspaceBizId *string, livyComputeId *string, request *UpdateLivyComputeRequest) (_result *UpdateLivyComputeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLivyComputeResponse{}
	_body, _err := client.UpdateLivyComputeWithOptions(workspaceBizId, livyComputeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the workflow and time-based scheduling configurations.
//
// @param tmpReq - UpdateProcessDefinitionWithScheduleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProcessDefinitionWithScheduleResponse
func (client *Client) UpdateProcessDefinitionWithScheduleWithOptions(bizId *string, code *string, tmpReq *UpdateProcessDefinitionWithScheduleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProcessDefinitionWithScheduleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateProcessDefinitionWithScheduleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.GlobalParams)) {
		request.GlobalParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GlobalParams, tea.String("globalParams"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Schedule)) {
		request.ScheduleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Schedule, tea.String("schedule"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("tags"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TaskDefinitionJson)) {
		request.TaskDefinitionJsonShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskDefinitionJson, tea.String("taskDefinitionJson"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TaskRelationJson)) {
		request.TaskRelationJsonShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskRelationJson, tea.String("taskRelationJson"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertEmailAddress)) {
		query["alertEmailAddress"] = request.AlertEmailAddress
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionType)) {
		query["executionType"] = request.ExecutionType
	}

	if !tea.BoolValue(util.IsUnset(request.GlobalParamsShrink)) {
		query["globalParams"] = request.GlobalParamsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProductNamespace)) {
		query["productNamespace"] = request.ProductNamespace
	}

	if !tea.BoolValue(util.IsUnset(request.Publish)) {
		query["publish"] = request.Publish
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseState)) {
		query["releaseState"] = request.ReleaseState
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceQueue)) {
		query["resourceQueue"] = request.ResourceQueue
	}

	if !tea.BoolValue(util.IsUnset(request.RetryTimes)) {
		query["retryTimes"] = request.RetryTimes
	}

	if !tea.BoolValue(util.IsUnset(request.RunAs)) {
		query["runAs"] = request.RunAs
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleShrink)) {
		query["schedule"] = request.ScheduleShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskDefinitionJsonShrink)) {
		query["taskDefinitionJson"] = request.TaskDefinitionJsonShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskParallelism)) {
		query["taskParallelism"] = request.TaskParallelism
	}

	if !tea.BoolValue(util.IsUnset(request.TaskRelationJsonShrink)) {
		query["taskRelationJson"] = request.TaskRelationJsonShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		query["timeout"] = request.Timeout
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProcessDefinitionWithSchedule"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/projects/" + tea.StringValue(openapiutil.GetEncodeParam(bizId)) + "/process-definition/" + tea.StringValue(openapiutil.GetEncodeParam(code))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProcessDefinitionWithScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the workflow and time-based scheduling configurations.
//
// @param request - UpdateProcessDefinitionWithScheduleRequest
//
// @return UpdateProcessDefinitionWithScheduleResponse
func (client *Client) UpdateProcessDefinitionWithSchedule(bizId *string, code *string, request *UpdateProcessDefinitionWithScheduleRequest) (_result *UpdateProcessDefinitionWithScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProcessDefinitionWithScheduleResponse{}
	_body, _err := client.UpdateProcessDefinitionWithScheduleWithOptions(bizId, code, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
