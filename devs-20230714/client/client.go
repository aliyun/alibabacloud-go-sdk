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
	// Artifact
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-artifact
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	ResourceVersion *int32          `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
	Spec            *ArtifactSpec   `json:"spec,omitempty" xml:"spec,omitempty"`
	Status          *ArtifactStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// example:
	//
	// 2021-11-19T09:34:38Z
	UpdatedTime *string `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
}

func (s Artifact) String() string {
	return tea.Prettify(s)
}

func (s Artifact) GoString() string {
	return s.String()
}

func (s *Artifact) SetCreatedTime(v string) *Artifact {
	s.CreatedTime = &v
	return s
}

func (s *Artifact) SetDescription(v string) *Artifact {
	s.Description = &v
	return s
}

func (s *Artifact) SetGeneration(v int32) *Artifact {
	s.Generation = &v
	return s
}

func (s *Artifact) SetKind(v string) *Artifact {
	s.Kind = &v
	return s
}

func (s *Artifact) SetLabels(v map[string]*string) *Artifact {
	s.Labels = v
	return s
}

func (s *Artifact) SetName(v string) *Artifact {
	s.Name = &v
	return s
}

func (s *Artifact) SetResourceVersion(v int32) *Artifact {
	s.ResourceVersion = &v
	return s
}

func (s *Artifact) SetSpec(v *ArtifactSpec) *Artifact {
	s.Spec = v
	return s
}

func (s *Artifact) SetStatus(v *ArtifactStatus) *Artifact {
	s.Status = v
	return s
}

func (s *Artifact) SetUid(v string) *Artifact {
	s.Uid = &v
	return s
}

func (s *Artifact) SetUpdatedTime(v string) *Artifact {
	s.UpdatedTime = &v
	return s
}

type ArtifactCode struct {
	Checksum *string `json:"checksum,omitempty" xml:"checksum,omitempty"`
	Url      *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ArtifactCode) String() string {
	return tea.Prettify(s)
}

func (s ArtifactCode) GoString() string {
	return s.String()
}

func (s *ArtifactCode) SetChecksum(v string) *ArtifactCode {
	s.Checksum = &v
	return s
}

func (s *ArtifactCode) SetUrl(v string) *ArtifactCode {
	s.Url = &v
	return s
}

type ArtifactMeta struct {
	// example:
	//
	// CRC-64 code
	Checksum *string `json:"checksum,omitempty" xml:"checksum,omitempty"`
	// example:
	//
	// my-artifact
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ArtifactMeta) String() string {
	return tea.Prettify(s)
}

func (s ArtifactMeta) GoString() string {
	return s.String()
}

func (s *ArtifactMeta) SetChecksum(v string) *ArtifactMeta {
	s.Checksum = &v
	return s
}

func (s *ArtifactMeta) SetName(v string) *ArtifactMeta {
	s.Name = &v
	return s
}

type ArtifactSpec struct {
	// This parameter is required.
	//
	// example:
	//
	// custom.debian10
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FC代码包、工作流yaml
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss://cn-hangzhou/my-bucket/my.zip
	Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s ArtifactSpec) String() string {
	return tea.Prettify(s)
}

func (s ArtifactSpec) GoString() string {
	return s.String()
}

func (s *ArtifactSpec) SetRuntime(v string) *ArtifactSpec {
	s.Runtime = &v
	return s
}

func (s *ArtifactSpec) SetType(v string) *ArtifactSpec {
	s.Type = &v
	return s
}

func (s *ArtifactSpec) SetUri(v string) *ArtifactSpec {
	s.Uri = &v
	return s
}

type ArtifactStatus struct {
	// example:
	//
	// acs:devs:cn-hangzhou:123456:artifacts/my-first-artifact
	Arn *string `json:"arn,omitempty" xml:"arn,omitempty"`
	// example:
	//
	// 2825179536350****
	Checksum *string `json:"checksum,omitempty" xml:"checksum,omitempty"`
	// example:
	//
	// 1024
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ArtifactStatus) String() string {
	return tea.Prettify(s)
}

func (s ArtifactStatus) GoString() string {
	return s.String()
}

func (s *ArtifactStatus) SetArn(v string) *ArtifactStatus {
	s.Arn = &v
	return s
}

func (s *ArtifactStatus) SetChecksum(v string) *ArtifactStatus {
	s.Checksum = &v
	return s
}

func (s *ArtifactStatus) SetSize(v int64) *ArtifactStatus {
	s.Size = &v
	return s
}

type ArtifactTempBucketToken struct {
	Credentials   *ArtifactTempBucketTokenCredentials `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Struct"`
	OssBucketName *string                             `json:"ossBucketName,omitempty" xml:"ossBucketName,omitempty"`
	OssObjectName *string                             `json:"ossObjectName,omitempty" xml:"ossObjectName,omitempty"`
	OssRegion     *string                             `json:"ossRegion,omitempty" xml:"ossRegion,omitempty"`
}

func (s ArtifactTempBucketToken) String() string {
	return tea.Prettify(s)
}

func (s ArtifactTempBucketToken) GoString() string {
	return s.String()
}

func (s *ArtifactTempBucketToken) SetCredentials(v *ArtifactTempBucketTokenCredentials) *ArtifactTempBucketToken {
	s.Credentials = v
	return s
}

func (s *ArtifactTempBucketToken) SetOssBucketName(v string) *ArtifactTempBucketToken {
	s.OssBucketName = &v
	return s
}

func (s *ArtifactTempBucketToken) SetOssObjectName(v string) *ArtifactTempBucketToken {
	s.OssObjectName = &v
	return s
}

func (s *ArtifactTempBucketToken) SetOssRegion(v string) *ArtifactTempBucketToken {
	s.OssRegion = &v
	return s
}

type ArtifactTempBucketTokenCredentials struct {
	AccessKeyId     *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	SecurityToken   *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
}

func (s ArtifactTempBucketTokenCredentials) String() string {
	return tea.Prettify(s)
}

func (s ArtifactTempBucketTokenCredentials) GoString() string {
	return s.String()
}

func (s *ArtifactTempBucketTokenCredentials) SetAccessKeyId(v string) *ArtifactTempBucketTokenCredentials {
	s.AccessKeyId = &v
	return s
}

func (s *ArtifactTempBucketTokenCredentials) SetAccessKeySecret(v string) *ArtifactTempBucketTokenCredentials {
	s.AccessKeySecret = &v
	return s
}

func (s *ArtifactTempBucketTokenCredentials) SetSecurityToken(v string) *ArtifactTempBucketTokenCredentials {
	s.SecurityToken = &v
	return s
}

type BranchFilter struct {
	// example:
	//
	// master
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s BranchFilter) String() string {
	return tea.Prettify(s)
}

func (s BranchFilter) GoString() string {
	return s.String()
}

func (s *BranchFilter) SetName(v string) *BranchFilter {
	s.Name = &v
	return s
}

type BuildCacheConfig struct {
	// example:
	//
	// { 	"3C75C832-0EAD-40D6-8FA1-2BA9171C926B": "~/.npm", 	"D256BB7A-1886-4A19-A75B-A1FDC23D5A00": "~/.cache" }
	KeyPath map[string]interface{} `json:"keyPath,omitempty" xml:"keyPath,omitempty"`
	Paths   []*string              `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
}

func (s BuildCacheConfig) String() string {
	return tea.Prettify(s)
}

func (s BuildCacheConfig) GoString() string {
	return s.String()
}

func (s *BuildCacheConfig) SetKeyPath(v map[string]interface{}) *BuildCacheConfig {
	s.KeyPath = v
	return s
}

func (s *BuildCacheConfig) SetPaths(v []*string) *BuildCacheConfig {
	s.Paths = v
	return s
}

type BuildConfig struct {
	Default *DefaultBuilderConfig `json:"default,omitempty" xml:"default,omitempty"`
}

func (s BuildConfig) String() string {
	return tea.Prettify(s)
}

func (s BuildConfig) GoString() string {
	return s.String()
}

func (s *BuildConfig) SetDefault(v *DefaultBuilderConfig) *BuildConfig {
	s.Default = v
	return s
}

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

type CodeVersionReference struct {
	// example:
	//
	// main
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// 12721ec262d03a93809ba2bbc717963cb298ceca
	CommitID *string `json:"commitID,omitempty" xml:"commitID,omitempty"`
}

func (s CodeVersionReference) String() string {
	return tea.Prettify(s)
}

func (s CodeVersionReference) GoString() string {
	return s.String()
}

func (s *CodeVersionReference) SetBranch(v string) *CodeVersionReference {
	s.Branch = &v
	return s
}

func (s *CodeVersionReference) SetCommitID(v string) *CodeVersionReference {
	s.CommitID = &v
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
	// Connection
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-connection
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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

type DefaultBuilderConfig struct {
	Cache     *BuildCacheConfig `json:"cache,omitempty" xml:"cache,omitempty"`
	Languages []*string         `json:"languages,omitempty" xml:"languages,omitempty" type:"Repeated"`
	Steps     []interface{}     `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
}

func (s DefaultBuilderConfig) String() string {
	return tea.Prettify(s)
}

func (s DefaultBuilderConfig) GoString() string {
	return s.String()
}

func (s *DefaultBuilderConfig) SetCache(v *BuildCacheConfig) *DefaultBuilderConfig {
	s.Cache = v
	return s
}

func (s *DefaultBuilderConfig) SetLanguages(v []*string) *DefaultBuilderConfig {
	s.Languages = v
	return s
}

func (s *DefaultBuilderConfig) SetSteps(v []interface{}) *DefaultBuilderConfig {
	s.Steps = v
	return s
}

type DeleteModelOutput struct {
	ErrCode   *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg    *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteModelOutput) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelOutput) GoString() string {
	return s.String()
}

func (s *DeleteModelOutput) SetErrCode(v string) *DeleteModelOutput {
	s.ErrCode = &v
	return s
}

func (s *DeleteModelOutput) SetErrMsg(v string) *DeleteModelOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeleteModelOutput) SetRequestId(v string) *DeleteModelOutput {
	s.RequestId = &v
	return s
}

func (s *DeleteModelOutput) SetSuccess(v bool) *DeleteModelOutput {
	s.Success = &v
	return s
}

type DeployCustomContainerAsyncOutput struct {
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployCustomContainerAsyncOutput) String() string {
	return tea.Prettify(s)
}

func (s DeployCustomContainerAsyncOutput) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerAsyncOutput) SetData(v string) *DeployCustomContainerAsyncOutput {
	s.Data = &v
	return s
}

func (s *DeployCustomContainerAsyncOutput) SetErrCode(v string) *DeployCustomContainerAsyncOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployCustomContainerAsyncOutput) SetErrMsg(v string) *DeployCustomContainerAsyncOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployCustomContainerAsyncOutput) SetRequestId(v string) *DeployCustomContainerAsyncOutput {
	s.RequestId = &v
	return s
}

func (s *DeployCustomContainerAsyncOutput) SetSuccess(v bool) *DeployCustomContainerAsyncOutput {
	s.Success = &v
	return s
}

type DeployCustomContainerInput struct {
	AccountID             *string                                          `json:"accountID,omitempty" xml:"accountID,omitempty"`
	ConcurrencyConfig     *DeployCustomContainerInputConcurrencyConfig     `json:"concurrencyConfig,omitempty" xml:"concurrencyConfig,omitempty" type:"Struct"`
	Cpu                   *float32                                         `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CustomContainerConfig *DeployCustomContainerInputCustomContainerConfig `json:"customContainerConfig,omitempty" xml:"customContainerConfig,omitempty" type:"Struct"`
	Description           *string                                          `json:"description,omitempty" xml:"description,omitempty"`
	DiskSize              *int32                                           `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnvName               *string                                          `json:"envName,omitempty" xml:"envName,omitempty"`
	EnvironmentVariables  map[string]interface{}                           `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	GpuConfig             *DeployCustomContainerInputGpuConfig             `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty" type:"Struct"`
	HttpTrigger           *DeployCustomContainerInputHttpTrigger           `json:"httpTrigger,omitempty" xml:"httpTrigger,omitempty" type:"Struct"`
	LogConfig             *DeployCustomContainerInputLogConfig             `json:"logConfig,omitempty" xml:"logConfig,omitempty" type:"Struct"`
	MemorySize            *int32                                           `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	ModelConfig           *DeployCustomContainerInputModelConfig           `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Name            *string                                    `json:"name,omitempty" xml:"name,omitempty"`
	NasConfig       *DeployCustomContainerInputNasConfig       `json:"nasConfig,omitempty" xml:"nasConfig,omitempty" type:"Struct"`
	OriginalName    *string                                    `json:"originalName,omitempty" xml:"originalName,omitempty"`
	ProjectName     *string                                    `json:"projectName,omitempty" xml:"projectName,omitempty"`
	ProvisionConfig *DeployCustomContainerInputProvisionConfig `json:"provisionConfig,omitempty" xml:"provisionConfig,omitempty" type:"Struct"`
	Region          *string                                    `json:"region,omitempty" xml:"region,omitempty"`
	ReportStatusURL *string                                    `json:"reportStatusURL,omitempty" xml:"reportStatusURL,omitempty"`
	// This parameter is required.
	Role      *string                              `json:"role,omitempty" xml:"role,omitempty"`
	Timeout   *int32                               `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TraceId   *string                              `json:"traceId,omitempty" xml:"traceId,omitempty"`
	VpcConfig *DeployCustomContainerInputVpcConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty" type:"Struct"`
}

func (s DeployCustomContainerInput) String() string {
	return tea.Prettify(s)
}

func (s DeployCustomContainerInput) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInput) SetAccountID(v string) *DeployCustomContainerInput {
	s.AccountID = &v
	return s
}

func (s *DeployCustomContainerInput) SetConcurrencyConfig(v *DeployCustomContainerInputConcurrencyConfig) *DeployCustomContainerInput {
	s.ConcurrencyConfig = v
	return s
}

func (s *DeployCustomContainerInput) SetCpu(v float32) *DeployCustomContainerInput {
	s.Cpu = &v
	return s
}

func (s *DeployCustomContainerInput) SetCustomContainerConfig(v *DeployCustomContainerInputCustomContainerConfig) *DeployCustomContainerInput {
	s.CustomContainerConfig = v
	return s
}

func (s *DeployCustomContainerInput) SetDescription(v string) *DeployCustomContainerInput {
	s.Description = &v
	return s
}

func (s *DeployCustomContainerInput) SetDiskSize(v int32) *DeployCustomContainerInput {
	s.DiskSize = &v
	return s
}

func (s *DeployCustomContainerInput) SetEnvName(v string) *DeployCustomContainerInput {
	s.EnvName = &v
	return s
}

func (s *DeployCustomContainerInput) SetEnvironmentVariables(v map[string]interface{}) *DeployCustomContainerInput {
	s.EnvironmentVariables = v
	return s
}

func (s *DeployCustomContainerInput) SetGpuConfig(v *DeployCustomContainerInputGpuConfig) *DeployCustomContainerInput {
	s.GpuConfig = v
	return s
}

func (s *DeployCustomContainerInput) SetHttpTrigger(v *DeployCustomContainerInputHttpTrigger) *DeployCustomContainerInput {
	s.HttpTrigger = v
	return s
}

func (s *DeployCustomContainerInput) SetLogConfig(v *DeployCustomContainerInputLogConfig) *DeployCustomContainerInput {
	s.LogConfig = v
	return s
}

func (s *DeployCustomContainerInput) SetMemorySize(v int32) *DeployCustomContainerInput {
	s.MemorySize = &v
	return s
}

func (s *DeployCustomContainerInput) SetModelConfig(v *DeployCustomContainerInputModelConfig) *DeployCustomContainerInput {
	s.ModelConfig = v
	return s
}

func (s *DeployCustomContainerInput) SetName(v string) *DeployCustomContainerInput {
	s.Name = &v
	return s
}

func (s *DeployCustomContainerInput) SetNasConfig(v *DeployCustomContainerInputNasConfig) *DeployCustomContainerInput {
	s.NasConfig = v
	return s
}

func (s *DeployCustomContainerInput) SetOriginalName(v string) *DeployCustomContainerInput {
	s.OriginalName = &v
	return s
}

func (s *DeployCustomContainerInput) SetProjectName(v string) *DeployCustomContainerInput {
	s.ProjectName = &v
	return s
}

func (s *DeployCustomContainerInput) SetProvisionConfig(v *DeployCustomContainerInputProvisionConfig) *DeployCustomContainerInput {
	s.ProvisionConfig = v
	return s
}

func (s *DeployCustomContainerInput) SetRegion(v string) *DeployCustomContainerInput {
	s.Region = &v
	return s
}

func (s *DeployCustomContainerInput) SetReportStatusURL(v string) *DeployCustomContainerInput {
	s.ReportStatusURL = &v
	return s
}

func (s *DeployCustomContainerInput) SetRole(v string) *DeployCustomContainerInput {
	s.Role = &v
	return s
}

func (s *DeployCustomContainerInput) SetTimeout(v int32) *DeployCustomContainerInput {
	s.Timeout = &v
	return s
}

func (s *DeployCustomContainerInput) SetTraceId(v string) *DeployCustomContainerInput {
	s.TraceId = &v
	return s
}

func (s *DeployCustomContainerInput) SetVpcConfig(v *DeployCustomContainerInputVpcConfig) *DeployCustomContainerInput {
	s.VpcConfig = v
	return s
}

type DeployCustomContainerInputConcurrencyConfig struct {
	ReservedConcurrency *int32 `json:"reservedConcurrency,omitempty" xml:"reservedConcurrency,omitempty"`
}

func (s DeployCustomContainerInputConcurrencyConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployCustomContainerInputConcurrencyConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputConcurrencyConfig) SetReservedConcurrency(v int32) *DeployCustomContainerInputConcurrencyConfig {
	s.ReservedConcurrency = &v
	return s
}

type DeployCustomContainerInputCustomContainerConfig struct {
	Command                 []*string                                                               `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
	Entrypoint              []*string                                                               `json:"entrypoint,omitempty" xml:"entrypoint,omitempty" type:"Repeated"`
	HealthCheckConfig       *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig       `json:"healthCheckConfig,omitempty" xml:"healthCheckConfig,omitempty" type:"Struct"`
	Image                   *string                                                                 `json:"image,omitempty" xml:"image,omitempty"`
	InstanceConcurrency     *int32                                                                  `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	InstanceLifecycleConfig *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig `json:"instanceLifecycleConfig,omitempty" xml:"instanceLifecycleConfig,omitempty" type:"Struct"`
	Port                    *int32                                                                  `json:"port,omitempty" xml:"port,omitempty"`
}

func (s DeployCustomContainerInputCustomContainerConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployCustomContainerInputCustomContainerConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputCustomContainerConfig) SetCommand(v []*string) *DeployCustomContainerInputCustomContainerConfig {
	s.Command = v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfig) SetEntrypoint(v []*string) *DeployCustomContainerInputCustomContainerConfig {
	s.Entrypoint = v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfig) SetHealthCheckConfig(v *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) *DeployCustomContainerInputCustomContainerConfig {
	s.HealthCheckConfig = v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfig) SetImage(v string) *DeployCustomContainerInputCustomContainerConfig {
	s.Image = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfig) SetInstanceConcurrency(v int32) *DeployCustomContainerInputCustomContainerConfig {
	s.InstanceConcurrency = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfig) SetInstanceLifecycleConfig(v *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig) *DeployCustomContainerInputCustomContainerConfig {
	s.InstanceLifecycleConfig = v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfig) SetPort(v int32) *DeployCustomContainerInputCustomContainerConfig {
	s.Port = &v
	return s
}

type DeployCustomContainerInputCustomContainerConfigHealthCheckConfig struct {
	FailureThreshold    *int32  `json:"failureThreshold,omitempty" xml:"failureThreshold,omitempty"`
	HttpGetUrl          *string `json:"httpGetUrl,omitempty" xml:"httpGetUrl,omitempty"`
	InitialDelaySeconds *int32  `json:"initialDelaySeconds,omitempty" xml:"initialDelaySeconds,omitempty"`
	PeriodSeconds       *int32  `json:"periodSeconds,omitempty" xml:"periodSeconds,omitempty"`
	SuccessThreshold    *int32  `json:"successThreshold,omitempty" xml:"successThreshold,omitempty"`
	TimeoutSeconds      *int64  `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
}

func (s DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) SetFailureThreshold(v int32) *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig {
	s.FailureThreshold = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) SetHttpGetUrl(v string) *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig {
	s.HttpGetUrl = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) SetInitialDelaySeconds(v int32) *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig {
	s.InitialDelaySeconds = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) SetPeriodSeconds(v int32) *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig {
	s.PeriodSeconds = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) SetSuccessThreshold(v int32) *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig {
	s.SuccessThreshold = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) SetTimeoutSeconds(v int64) *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig {
	s.TimeoutSeconds = &v
	return s
}

type DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig struct {
	Initializer *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer `json:"initializer,omitempty" xml:"initializer,omitempty" type:"Struct"`
	PreStop     *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop     `json:"preStop,omitempty" xml:"preStop,omitempty" type:"Struct"`
}

func (s DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig) SetInitializer(v *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer) *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig {
	s.Initializer = v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig) SetPreStop(v *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop) *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig {
	s.PreStop = v
	return s
}

type DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer struct {
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	Timeout *int32  `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer) String() string {
	return tea.Prettify(s)
}

func (s DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer) SetHandler(v string) *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer {
	s.Handler = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer) SetTimeout(v int32) *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer {
	s.Timeout = &v
	return s
}

type DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop struct {
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	Timeout *int32  `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop) String() string {
	return tea.Prettify(s)
}

func (s DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop) SetHandler(v string) *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop {
	s.Handler = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop) SetTimeout(v int32) *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop {
	s.Timeout = &v
	return s
}

type DeployCustomContainerInputGpuConfig struct {
	GpuMemorySize *int64  `json:"gpuMemorySize,omitempty" xml:"gpuMemorySize,omitempty"`
	GpuType       *string `json:"gpuType,omitempty" xml:"gpuType,omitempty"`
}

func (s DeployCustomContainerInputGpuConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployCustomContainerInputGpuConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputGpuConfig) SetGpuMemorySize(v int64) *DeployCustomContainerInputGpuConfig {
	s.GpuMemorySize = &v
	return s
}

func (s *DeployCustomContainerInputGpuConfig) SetGpuType(v string) *DeployCustomContainerInputGpuConfig {
	s.GpuType = &v
	return s
}

type DeployCustomContainerInputHttpTrigger struct {
	Qualifier     *string                                             `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	TriggerConfig *DeployCustomContainerInputHttpTriggerTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
}

func (s DeployCustomContainerInputHttpTrigger) String() string {
	return tea.Prettify(s)
}

func (s DeployCustomContainerInputHttpTrigger) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputHttpTrigger) SetQualifier(v string) *DeployCustomContainerInputHttpTrigger {
	s.Qualifier = &v
	return s
}

func (s *DeployCustomContainerInputHttpTrigger) SetTriggerConfig(v *DeployCustomContainerInputHttpTriggerTriggerConfig) *DeployCustomContainerInputHttpTrigger {
	s.TriggerConfig = v
	return s
}

type DeployCustomContainerInputHttpTriggerTriggerConfig struct {
	AuthType          *string   `json:"authType,omitempty" xml:"authType,omitempty"`
	DsableURLInternet *bool     `json:"dsableURLInternet,omitempty" xml:"dsableURLInternet,omitempty"`
	Methods           []*string `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
}

func (s DeployCustomContainerInputHttpTriggerTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployCustomContainerInputHttpTriggerTriggerConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputHttpTriggerTriggerConfig) SetAuthType(v string) *DeployCustomContainerInputHttpTriggerTriggerConfig {
	s.AuthType = &v
	return s
}

func (s *DeployCustomContainerInputHttpTriggerTriggerConfig) SetDsableURLInternet(v bool) *DeployCustomContainerInputHttpTriggerTriggerConfig {
	s.DsableURLInternet = &v
	return s
}

func (s *DeployCustomContainerInputHttpTriggerTriggerConfig) SetMethods(v []*string) *DeployCustomContainerInputHttpTriggerTriggerConfig {
	s.Methods = v
	return s
}

type DeployCustomContainerInputLogConfig struct {
	EnableInstanceMetrics *bool   `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	EnableRequestMetrics  *bool   `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	LogBeginRule          *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	Logstore              *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	Project               *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s DeployCustomContainerInputLogConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployCustomContainerInputLogConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputLogConfig) SetEnableInstanceMetrics(v bool) *DeployCustomContainerInputLogConfig {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *DeployCustomContainerInputLogConfig) SetEnableRequestMetrics(v bool) *DeployCustomContainerInputLogConfig {
	s.EnableRequestMetrics = &v
	return s
}

func (s *DeployCustomContainerInputLogConfig) SetLogBeginRule(v string) *DeployCustomContainerInputLogConfig {
	s.LogBeginRule = &v
	return s
}

func (s *DeployCustomContainerInputLogConfig) SetLogstore(v string) *DeployCustomContainerInputLogConfig {
	s.Logstore = &v
	return s
}

func (s *DeployCustomContainerInputLogConfig) SetProject(v string) *DeployCustomContainerInputLogConfig {
	s.Project = &v
	return s
}

type DeployCustomContainerInputModelConfig struct {
	Framework *string `json:"framework,omitempty" xml:"framework,omitempty"`
	// if can be null:
	// true
	MultiModelConfig           []*ModelConfig `json:"multiModelConfig,omitempty" xml:"multiModelConfig,omitempty" type:"Repeated"`
	Prefix                     *string        `json:"prefix,omitempty" xml:"prefix,omitempty"`
	SourceType                 *string        `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	SrcModelScopeModelID       *string        `json:"srcModelScopeModelID,omitempty" xml:"srcModelScopeModelID,omitempty"`
	SrcModelScopeModelRevision *string        `json:"srcModelScopeModelRevision,omitempty" xml:"srcModelScopeModelRevision,omitempty"`
	SrcModelScopeToken         *string        `json:"srcModelScopeToken,omitempty" xml:"srcModelScopeToken,omitempty"`
	SrcOssBucket               *string        `json:"srcOssBucket,omitempty" xml:"srcOssBucket,omitempty"`
	SrcOssPath                 *string        `json:"srcOssPath,omitempty" xml:"srcOssPath,omitempty"`
	SrcOssRegion               *string        `json:"srcOssRegion,omitempty" xml:"srcOssRegion,omitempty"`
}

func (s DeployCustomContainerInputModelConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployCustomContainerInputModelConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputModelConfig) SetFramework(v string) *DeployCustomContainerInputModelConfig {
	s.Framework = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetMultiModelConfig(v []*ModelConfig) *DeployCustomContainerInputModelConfig {
	s.MultiModelConfig = v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetPrefix(v string) *DeployCustomContainerInputModelConfig {
	s.Prefix = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetSourceType(v string) *DeployCustomContainerInputModelConfig {
	s.SourceType = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetSrcModelScopeModelID(v string) *DeployCustomContainerInputModelConfig {
	s.SrcModelScopeModelID = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetSrcModelScopeModelRevision(v string) *DeployCustomContainerInputModelConfig {
	s.SrcModelScopeModelRevision = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetSrcModelScopeToken(v string) *DeployCustomContainerInputModelConfig {
	s.SrcModelScopeToken = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetSrcOssBucket(v string) *DeployCustomContainerInputModelConfig {
	s.SrcOssBucket = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetSrcOssPath(v string) *DeployCustomContainerInputModelConfig {
	s.SrcOssPath = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetSrcOssRegion(v string) *DeployCustomContainerInputModelConfig {
	s.SrcOssRegion = &v
	return s
}

type DeployCustomContainerInputNasConfig struct {
	GroupId     *int64    `json:"groupId,omitempty" xml:"groupId,omitempty"`
	MountPoints []*string `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	UserId      *int64    `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeployCustomContainerInputNasConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployCustomContainerInputNasConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputNasConfig) SetGroupId(v int64) *DeployCustomContainerInputNasConfig {
	s.GroupId = &v
	return s
}

func (s *DeployCustomContainerInputNasConfig) SetMountPoints(v []*string) *DeployCustomContainerInputNasConfig {
	s.MountPoints = v
	return s
}

func (s *DeployCustomContainerInputNasConfig) SetUserId(v int64) *DeployCustomContainerInputNasConfig {
	s.UserId = &v
	return s
}

type DeployCustomContainerInputProvisionConfig struct {
	AlwaysAllocateGPU *bool                                                        `json:"alwaysAllocateGPU,omitempty" xml:"alwaysAllocateGPU,omitempty"`
	ScheduledActions  []*DeployCustomContainerInputProvisionConfigScheduledActions `json:"scheduledActions,omitempty" xml:"scheduledActions,omitempty" type:"Repeated"`
	Target            *int64                                                       `json:"target,omitempty" xml:"target,omitempty"`
}

func (s DeployCustomContainerInputProvisionConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployCustomContainerInputProvisionConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputProvisionConfig) SetAlwaysAllocateGPU(v bool) *DeployCustomContainerInputProvisionConfig {
	s.AlwaysAllocateGPU = &v
	return s
}

func (s *DeployCustomContainerInputProvisionConfig) SetScheduledActions(v []*DeployCustomContainerInputProvisionConfigScheduledActions) *DeployCustomContainerInputProvisionConfig {
	s.ScheduledActions = v
	return s
}

func (s *DeployCustomContainerInputProvisionConfig) SetTarget(v int64) *DeployCustomContainerInputProvisionConfig {
	s.Target = &v
	return s
}

type DeployCustomContainerInputProvisionConfigScheduledActions struct {
	EndTime            *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	ScheduleExpression *string `json:"scheduleExpression,omitempty" xml:"scheduleExpression,omitempty"`
	StartTime          *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Target             *int32  `json:"target,omitempty" xml:"target,omitempty"`
	TimeZone           *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s DeployCustomContainerInputProvisionConfigScheduledActions) String() string {
	return tea.Prettify(s)
}

func (s DeployCustomContainerInputProvisionConfigScheduledActions) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputProvisionConfigScheduledActions) SetEndTime(v string) *DeployCustomContainerInputProvisionConfigScheduledActions {
	s.EndTime = &v
	return s
}

func (s *DeployCustomContainerInputProvisionConfigScheduledActions) SetName(v string) *DeployCustomContainerInputProvisionConfigScheduledActions {
	s.Name = &v
	return s
}

func (s *DeployCustomContainerInputProvisionConfigScheduledActions) SetScheduleExpression(v string) *DeployCustomContainerInputProvisionConfigScheduledActions {
	s.ScheduleExpression = &v
	return s
}

func (s *DeployCustomContainerInputProvisionConfigScheduledActions) SetStartTime(v string) *DeployCustomContainerInputProvisionConfigScheduledActions {
	s.StartTime = &v
	return s
}

func (s *DeployCustomContainerInputProvisionConfigScheduledActions) SetTarget(v int32) *DeployCustomContainerInputProvisionConfigScheduledActions {
	s.Target = &v
	return s
}

func (s *DeployCustomContainerInputProvisionConfigScheduledActions) SetTimeZone(v string) *DeployCustomContainerInputProvisionConfigScheduledActions {
	s.TimeZone = &v
	return s
}

type DeployCustomContainerInputVpcConfig struct {
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	VpcId           *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s DeployCustomContainerInputVpcConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployCustomContainerInputVpcConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputVpcConfig) SetSecurityGroupId(v string) *DeployCustomContainerInputVpcConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *DeployCustomContainerInputVpcConfig) SetVSwitchIds(v []*string) *DeployCustomContainerInputVpcConfig {
	s.VSwitchIds = v
	return s
}

func (s *DeployCustomContainerInputVpcConfig) SetVpcId(v string) *DeployCustomContainerInputVpcConfig {
	s.VpcId = &v
	return s
}

type DeployCustomContainerOutput struct {
	Data    *DeployCustomContainerOutputData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                          `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                          `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployCustomContainerOutput) String() string {
	return tea.Prettify(s)
}

func (s DeployCustomContainerOutput) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerOutput) SetData(v *DeployCustomContainerOutputData) *DeployCustomContainerOutput {
	s.Data = v
	return s
}

func (s *DeployCustomContainerOutput) SetErrCode(v string) *DeployCustomContainerOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployCustomContainerOutput) SetErrMsg(v string) *DeployCustomContainerOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployCustomContainerOutput) SetRequestId(v string) *DeployCustomContainerOutput {
	s.RequestId = &v
	return s
}

func (s *DeployCustomContainerOutput) SetSuccess(v bool) *DeployCustomContainerOutput {
	s.Success = &v
	return s
}

type DeployCustomContainerOutputData struct {
	DeploymentTaskID *string `json:"deploymentTaskID,omitempty" xml:"deploymentTaskID,omitempty"`
	ErrorMessage     *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Finished         *bool   `json:"finished,omitempty" xml:"finished,omitempty"`
	NasConfigStr     *string `json:"nasConfigStr,omitempty" xml:"nasConfigStr,omitempty"`
	ServiceName      *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	TraceID          *string `json:"traceID,omitempty" xml:"traceID,omitempty"`
	UrlInternet      *string `json:"urlInternet,omitempty" xml:"urlInternet,omitempty"`
	UrlIntranet      *string `json:"urlIntranet,omitempty" xml:"urlIntranet,omitempty"`
	VpcConfigStr     *string `json:"vpcConfigStr,omitempty" xml:"vpcConfigStr,omitempty"`
}

func (s DeployCustomContainerOutputData) String() string {
	return tea.Prettify(s)
}

func (s DeployCustomContainerOutputData) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerOutputData) SetDeploymentTaskID(v string) *DeployCustomContainerOutputData {
	s.DeploymentTaskID = &v
	return s
}

func (s *DeployCustomContainerOutputData) SetErrorMessage(v string) *DeployCustomContainerOutputData {
	s.ErrorMessage = &v
	return s
}

func (s *DeployCustomContainerOutputData) SetFinished(v bool) *DeployCustomContainerOutputData {
	s.Finished = &v
	return s
}

func (s *DeployCustomContainerOutputData) SetNasConfigStr(v string) *DeployCustomContainerOutputData {
	s.NasConfigStr = &v
	return s
}

func (s *DeployCustomContainerOutputData) SetServiceName(v string) *DeployCustomContainerOutputData {
	s.ServiceName = &v
	return s
}

func (s *DeployCustomContainerOutputData) SetTraceID(v string) *DeployCustomContainerOutputData {
	s.TraceID = &v
	return s
}

func (s *DeployCustomContainerOutputData) SetUrlInternet(v string) *DeployCustomContainerOutputData {
	s.UrlInternet = &v
	return s
}

func (s *DeployCustomContainerOutputData) SetUrlIntranet(v string) *DeployCustomContainerOutputData {
	s.UrlIntranet = &v
	return s
}

func (s *DeployCustomContainerOutputData) SetVpcConfigStr(v string) *DeployCustomContainerOutputData {
	s.VpcConfigStr = &v
	return s
}

type DeployEnvironmentOptions struct {
	Services []*string `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s DeployEnvironmentOptions) String() string {
	return tea.Prettify(s)
}

func (s DeployEnvironmentOptions) GoString() string {
	return s.String()
}

func (s *DeployEnvironmentOptions) SetServices(v []*string) *DeployEnvironmentOptions {
	s.Services = v
	return s
}

type DeployHuggingFaceModelAsyncOutput struct {
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployHuggingFaceModelAsyncOutput) String() string {
	return tea.Prettify(s)
}

func (s DeployHuggingFaceModelAsyncOutput) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelAsyncOutput) SetData(v string) *DeployHuggingFaceModelAsyncOutput {
	s.Data = &v
	return s
}

func (s *DeployHuggingFaceModelAsyncOutput) SetErrCode(v string) *DeployHuggingFaceModelAsyncOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployHuggingFaceModelAsyncOutput) SetErrMsg(v string) *DeployHuggingFaceModelAsyncOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployHuggingFaceModelAsyncOutput) SetRequestId(v string) *DeployHuggingFaceModelAsyncOutput {
	s.RequestId = &v
	return s
}

func (s *DeployHuggingFaceModelAsyncOutput) SetSuccess(v bool) *DeployHuggingFaceModelAsyncOutput {
	s.Success = &v
	return s
}

type DeployHuggingFaceModelInput struct {
	AccountID            *string                                       `json:"accountID,omitempty" xml:"accountID,omitempty"`
	ConcurrencyConfig    *DeployHuggingFaceModelInputConcurrencyConfig `json:"concurrencyConfig,omitempty" xml:"concurrencyConfig,omitempty" type:"Struct"`
	Cpu                  *float32                                      `json:"cpu,omitempty" xml:"cpu,omitempty"`
	Description          *string                                       `json:"description,omitempty" xml:"description,omitempty"`
	DiskSize             *int32                                        `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnvName              *string                                       `json:"envName,omitempty" xml:"envName,omitempty"`
	EnvironmentVariables map[string]interface{}                        `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	GpuConfig            *DeployHuggingFaceModelInputGpuConfig         `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty" type:"Struct"`
	HttpTrigger          *DeployHuggingFaceModelInputHttpTrigger       `json:"httpTrigger,omitempty" xml:"httpTrigger,omitempty" type:"Struct"`
	ImageName            *string                                       `json:"imageName,omitempty" xml:"imageName,omitempty"`
	InstanceConcurrency  *int32                                        `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	LogConfig            *DeployHuggingFaceModelInputLogConfig         `json:"logConfig,omitempty" xml:"logConfig,omitempty" type:"Struct"`
	MemorySize           *int32                                        `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	ModelConfig          *DeployHuggingFaceModelInputModelConfig       `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Name            *string                                     `json:"name,omitempty" xml:"name,omitempty"`
	NasConfig       *DeployHuggingFaceModelInputNasConfig       `json:"nasConfig,omitempty" xml:"nasConfig,omitempty" type:"Struct"`
	OriginalName    *string                                     `json:"originalName,omitempty" xml:"originalName,omitempty"`
	ProjectName     *string                                     `json:"projectName,omitempty" xml:"projectName,omitempty"`
	ProvisionConfig *DeployHuggingFaceModelInputProvisionConfig `json:"provisionConfig,omitempty" xml:"provisionConfig,omitempty" type:"Struct"`
	Region          *string                                     `json:"region,omitempty" xml:"region,omitempty"`
	ReportStatusURL *string                                     `json:"reportStatusURL,omitempty" xml:"reportStatusURL,omitempty"`
	// This parameter is required.
	Role      *string                               `json:"role,omitempty" xml:"role,omitempty"`
	Timeout   *int32                                `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TraceId   *string                               `json:"traceId,omitempty" xml:"traceId,omitempty"`
	VpcConfig *DeployHuggingFaceModelInputVpcConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty" type:"Struct"`
}

func (s DeployHuggingFaceModelInput) String() string {
	return tea.Prettify(s)
}

func (s DeployHuggingFaceModelInput) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInput) SetAccountID(v string) *DeployHuggingFaceModelInput {
	s.AccountID = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetConcurrencyConfig(v *DeployHuggingFaceModelInputConcurrencyConfig) *DeployHuggingFaceModelInput {
	s.ConcurrencyConfig = v
	return s
}

func (s *DeployHuggingFaceModelInput) SetCpu(v float32) *DeployHuggingFaceModelInput {
	s.Cpu = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetDescription(v string) *DeployHuggingFaceModelInput {
	s.Description = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetDiskSize(v int32) *DeployHuggingFaceModelInput {
	s.DiskSize = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetEnvName(v string) *DeployHuggingFaceModelInput {
	s.EnvName = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetEnvironmentVariables(v map[string]interface{}) *DeployHuggingFaceModelInput {
	s.EnvironmentVariables = v
	return s
}

func (s *DeployHuggingFaceModelInput) SetGpuConfig(v *DeployHuggingFaceModelInputGpuConfig) *DeployHuggingFaceModelInput {
	s.GpuConfig = v
	return s
}

func (s *DeployHuggingFaceModelInput) SetHttpTrigger(v *DeployHuggingFaceModelInputHttpTrigger) *DeployHuggingFaceModelInput {
	s.HttpTrigger = v
	return s
}

func (s *DeployHuggingFaceModelInput) SetImageName(v string) *DeployHuggingFaceModelInput {
	s.ImageName = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetInstanceConcurrency(v int32) *DeployHuggingFaceModelInput {
	s.InstanceConcurrency = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetLogConfig(v *DeployHuggingFaceModelInputLogConfig) *DeployHuggingFaceModelInput {
	s.LogConfig = v
	return s
}

func (s *DeployHuggingFaceModelInput) SetMemorySize(v int32) *DeployHuggingFaceModelInput {
	s.MemorySize = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetModelConfig(v *DeployHuggingFaceModelInputModelConfig) *DeployHuggingFaceModelInput {
	s.ModelConfig = v
	return s
}

func (s *DeployHuggingFaceModelInput) SetName(v string) *DeployHuggingFaceModelInput {
	s.Name = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetNasConfig(v *DeployHuggingFaceModelInputNasConfig) *DeployHuggingFaceModelInput {
	s.NasConfig = v
	return s
}

func (s *DeployHuggingFaceModelInput) SetOriginalName(v string) *DeployHuggingFaceModelInput {
	s.OriginalName = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetProjectName(v string) *DeployHuggingFaceModelInput {
	s.ProjectName = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetProvisionConfig(v *DeployHuggingFaceModelInputProvisionConfig) *DeployHuggingFaceModelInput {
	s.ProvisionConfig = v
	return s
}

func (s *DeployHuggingFaceModelInput) SetRegion(v string) *DeployHuggingFaceModelInput {
	s.Region = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetReportStatusURL(v string) *DeployHuggingFaceModelInput {
	s.ReportStatusURL = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetRole(v string) *DeployHuggingFaceModelInput {
	s.Role = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetTimeout(v int32) *DeployHuggingFaceModelInput {
	s.Timeout = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetTraceId(v string) *DeployHuggingFaceModelInput {
	s.TraceId = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetVpcConfig(v *DeployHuggingFaceModelInputVpcConfig) *DeployHuggingFaceModelInput {
	s.VpcConfig = v
	return s
}

type DeployHuggingFaceModelInputConcurrencyConfig struct {
	ReservedConcurrency *int32 `json:"reservedConcurrency,omitempty" xml:"reservedConcurrency,omitempty"`
}

func (s DeployHuggingFaceModelInputConcurrencyConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployHuggingFaceModelInputConcurrencyConfig) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputConcurrencyConfig) SetReservedConcurrency(v int32) *DeployHuggingFaceModelInputConcurrencyConfig {
	s.ReservedConcurrency = &v
	return s
}

type DeployHuggingFaceModelInputGpuConfig struct {
	GpuMemorySize *int32  `json:"gpuMemorySize,omitempty" xml:"gpuMemorySize,omitempty"`
	GpuType       *string `json:"gpuType,omitempty" xml:"gpuType,omitempty"`
}

func (s DeployHuggingFaceModelInputGpuConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployHuggingFaceModelInputGpuConfig) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputGpuConfig) SetGpuMemorySize(v int32) *DeployHuggingFaceModelInputGpuConfig {
	s.GpuMemorySize = &v
	return s
}

func (s *DeployHuggingFaceModelInputGpuConfig) SetGpuType(v string) *DeployHuggingFaceModelInputGpuConfig {
	s.GpuType = &v
	return s
}

type DeployHuggingFaceModelInputHttpTrigger struct {
	Qualifier     *string                                              `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	TriggerConfig *DeployHuggingFaceModelInputHttpTriggerTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
}

func (s DeployHuggingFaceModelInputHttpTrigger) String() string {
	return tea.Prettify(s)
}

func (s DeployHuggingFaceModelInputHttpTrigger) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputHttpTrigger) SetQualifier(v string) *DeployHuggingFaceModelInputHttpTrigger {
	s.Qualifier = &v
	return s
}

func (s *DeployHuggingFaceModelInputHttpTrigger) SetTriggerConfig(v *DeployHuggingFaceModelInputHttpTriggerTriggerConfig) *DeployHuggingFaceModelInputHttpTrigger {
	s.TriggerConfig = v
	return s
}

type DeployHuggingFaceModelInputHttpTriggerTriggerConfig struct {
	AuthType          *string   `json:"authType,omitempty" xml:"authType,omitempty"`
	DsableURLInternet *bool     `json:"dsableURLInternet,omitempty" xml:"dsableURLInternet,omitempty"`
	Methods           []*string `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
}

func (s DeployHuggingFaceModelInputHttpTriggerTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployHuggingFaceModelInputHttpTriggerTriggerConfig) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputHttpTriggerTriggerConfig) SetAuthType(v string) *DeployHuggingFaceModelInputHttpTriggerTriggerConfig {
	s.AuthType = &v
	return s
}

func (s *DeployHuggingFaceModelInputHttpTriggerTriggerConfig) SetDsableURLInternet(v bool) *DeployHuggingFaceModelInputHttpTriggerTriggerConfig {
	s.DsableURLInternet = &v
	return s
}

func (s *DeployHuggingFaceModelInputHttpTriggerTriggerConfig) SetMethods(v []*string) *DeployHuggingFaceModelInputHttpTriggerTriggerConfig {
	s.Methods = v
	return s
}

type DeployHuggingFaceModelInputLogConfig struct {
	EnableInstanceMetrics *bool   `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	EnableRequestMetrics  *bool   `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	LogBeginRule          *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	Logstore              *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	Project               *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s DeployHuggingFaceModelInputLogConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployHuggingFaceModelInputLogConfig) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputLogConfig) SetEnableInstanceMetrics(v bool) *DeployHuggingFaceModelInputLogConfig {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *DeployHuggingFaceModelInputLogConfig) SetEnableRequestMetrics(v bool) *DeployHuggingFaceModelInputLogConfig {
	s.EnableRequestMetrics = &v
	return s
}

func (s *DeployHuggingFaceModelInputLogConfig) SetLogBeginRule(v string) *DeployHuggingFaceModelInputLogConfig {
	s.LogBeginRule = &v
	return s
}

func (s *DeployHuggingFaceModelInputLogConfig) SetLogstore(v string) *DeployHuggingFaceModelInputLogConfig {
	s.Logstore = &v
	return s
}

func (s *DeployHuggingFaceModelInputLogConfig) SetProject(v string) *DeployHuggingFaceModelInputLogConfig {
	s.Project = &v
	return s
}

type DeployHuggingFaceModelInputModelConfig struct {
	FmkHuggingFaceConfig *DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig `json:"fmkHuggingFaceConfig,omitempty" xml:"fmkHuggingFaceConfig,omitempty" type:"Struct"`
	Framework            *string                                                     `json:"framework,omitempty" xml:"framework,omitempty"`
	// if can be null:
	// true
	MultiModelConfig           []*ModelConfig `json:"multiModelConfig,omitempty" xml:"multiModelConfig,omitempty" type:"Repeated"`
	Prefix                     *string        `json:"prefix,omitempty" xml:"prefix,omitempty"`
	SourceType                 *string        `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	SrcModelScopeModelID       *string        `json:"srcModelScopeModelID,omitempty" xml:"srcModelScopeModelID,omitempty"`
	SrcModelScopeModelRevision *string        `json:"srcModelScopeModelRevision,omitempty" xml:"srcModelScopeModelRevision,omitempty"`
	SrcModelScopeToken         *string        `json:"srcModelScopeToken,omitempty" xml:"srcModelScopeToken,omitempty"`
	SrcOssBucket               *string        `json:"srcOssBucket,omitempty" xml:"srcOssBucket,omitempty"`
	SrcOssPath                 *string        `json:"srcOssPath,omitempty" xml:"srcOssPath,omitempty"`
	SrcOssRegion               *string        `json:"srcOssRegion,omitempty" xml:"srcOssRegion,omitempty"`
}

func (s DeployHuggingFaceModelInputModelConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployHuggingFaceModelInputModelConfig) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputModelConfig) SetFmkHuggingFaceConfig(v *DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig) *DeployHuggingFaceModelInputModelConfig {
	s.FmkHuggingFaceConfig = v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetFramework(v string) *DeployHuggingFaceModelInputModelConfig {
	s.Framework = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetMultiModelConfig(v []*ModelConfig) *DeployHuggingFaceModelInputModelConfig {
	s.MultiModelConfig = v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetPrefix(v string) *DeployHuggingFaceModelInputModelConfig {
	s.Prefix = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetSourceType(v string) *DeployHuggingFaceModelInputModelConfig {
	s.SourceType = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetSrcModelScopeModelID(v string) *DeployHuggingFaceModelInputModelConfig {
	s.SrcModelScopeModelID = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetSrcModelScopeModelRevision(v string) *DeployHuggingFaceModelInputModelConfig {
	s.SrcModelScopeModelRevision = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetSrcModelScopeToken(v string) *DeployHuggingFaceModelInputModelConfig {
	s.SrcModelScopeToken = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetSrcOssBucket(v string) *DeployHuggingFaceModelInputModelConfig {
	s.SrcOssBucket = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetSrcOssPath(v string) *DeployHuggingFaceModelInputModelConfig {
	s.SrcOssPath = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetSrcOssRegion(v string) *DeployHuggingFaceModelInputModelConfig {
	s.SrcOssRegion = &v
	return s
}

type DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig struct {
	Framework *string `json:"framework,omitempty" xml:"framework,omitempty"`
	Task      *string `json:"task,omitempty" xml:"task,omitempty"`
}

func (s DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig) SetFramework(v string) *DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig {
	s.Framework = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig) SetTask(v string) *DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig {
	s.Task = &v
	return s
}

type DeployHuggingFaceModelInputNasConfig struct {
	GroupId     *int32    `json:"groupId,omitempty" xml:"groupId,omitempty"`
	MountPoints []*string `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	UserId      *int32    `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeployHuggingFaceModelInputNasConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployHuggingFaceModelInputNasConfig) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputNasConfig) SetGroupId(v int32) *DeployHuggingFaceModelInputNasConfig {
	s.GroupId = &v
	return s
}

func (s *DeployHuggingFaceModelInputNasConfig) SetMountPoints(v []*string) *DeployHuggingFaceModelInputNasConfig {
	s.MountPoints = v
	return s
}

func (s *DeployHuggingFaceModelInputNasConfig) SetUserId(v int32) *DeployHuggingFaceModelInputNasConfig {
	s.UserId = &v
	return s
}

type DeployHuggingFaceModelInputProvisionConfig struct {
	AlwaysAllocateGPU *bool                                                         `json:"alwaysAllocateGPU,omitempty" xml:"alwaysAllocateGPU,omitempty"`
	ScheduledActions  []*DeployHuggingFaceModelInputProvisionConfigScheduledActions `json:"scheduledActions,omitempty" xml:"scheduledActions,omitempty" type:"Repeated"`
	Target            *int32                                                        `json:"target,omitempty" xml:"target,omitempty"`
}

func (s DeployHuggingFaceModelInputProvisionConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployHuggingFaceModelInputProvisionConfig) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputProvisionConfig) SetAlwaysAllocateGPU(v bool) *DeployHuggingFaceModelInputProvisionConfig {
	s.AlwaysAllocateGPU = &v
	return s
}

func (s *DeployHuggingFaceModelInputProvisionConfig) SetScheduledActions(v []*DeployHuggingFaceModelInputProvisionConfigScheduledActions) *DeployHuggingFaceModelInputProvisionConfig {
	s.ScheduledActions = v
	return s
}

func (s *DeployHuggingFaceModelInputProvisionConfig) SetTarget(v int32) *DeployHuggingFaceModelInputProvisionConfig {
	s.Target = &v
	return s
}

type DeployHuggingFaceModelInputProvisionConfigScheduledActions struct {
	EndTime            *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	ScheduleExpression *string `json:"scheduleExpression,omitempty" xml:"scheduleExpression,omitempty"`
	StartTime          *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Target             *int32  `json:"target,omitempty" xml:"target,omitempty"`
	TimeZone           *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s DeployHuggingFaceModelInputProvisionConfigScheduledActions) String() string {
	return tea.Prettify(s)
}

func (s DeployHuggingFaceModelInputProvisionConfigScheduledActions) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputProvisionConfigScheduledActions) SetEndTime(v string) *DeployHuggingFaceModelInputProvisionConfigScheduledActions {
	s.EndTime = &v
	return s
}

func (s *DeployHuggingFaceModelInputProvisionConfigScheduledActions) SetName(v string) *DeployHuggingFaceModelInputProvisionConfigScheduledActions {
	s.Name = &v
	return s
}

func (s *DeployHuggingFaceModelInputProvisionConfigScheduledActions) SetScheduleExpression(v string) *DeployHuggingFaceModelInputProvisionConfigScheduledActions {
	s.ScheduleExpression = &v
	return s
}

func (s *DeployHuggingFaceModelInputProvisionConfigScheduledActions) SetStartTime(v string) *DeployHuggingFaceModelInputProvisionConfigScheduledActions {
	s.StartTime = &v
	return s
}

func (s *DeployHuggingFaceModelInputProvisionConfigScheduledActions) SetTarget(v int32) *DeployHuggingFaceModelInputProvisionConfigScheduledActions {
	s.Target = &v
	return s
}

func (s *DeployHuggingFaceModelInputProvisionConfigScheduledActions) SetTimeZone(v string) *DeployHuggingFaceModelInputProvisionConfigScheduledActions {
	s.TimeZone = &v
	return s
}

type DeployHuggingFaceModelInputVpcConfig struct {
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	VpcId           *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s DeployHuggingFaceModelInputVpcConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployHuggingFaceModelInputVpcConfig) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputVpcConfig) SetSecurityGroupId(v string) *DeployHuggingFaceModelInputVpcConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *DeployHuggingFaceModelInputVpcConfig) SetVSwitchIds(v []*string) *DeployHuggingFaceModelInputVpcConfig {
	s.VSwitchIds = v
	return s
}

func (s *DeployHuggingFaceModelInputVpcConfig) SetVpcId(v string) *DeployHuggingFaceModelInputVpcConfig {
	s.VpcId = &v
	return s
}

type DeployHuggingFaceModelOutput struct {
	Data    *DeployHuggingFaceModelOutputData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                           `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                           `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployHuggingFaceModelOutput) String() string {
	return tea.Prettify(s)
}

func (s DeployHuggingFaceModelOutput) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelOutput) SetData(v *DeployHuggingFaceModelOutputData) *DeployHuggingFaceModelOutput {
	s.Data = v
	return s
}

func (s *DeployHuggingFaceModelOutput) SetErrCode(v string) *DeployHuggingFaceModelOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployHuggingFaceModelOutput) SetErrMsg(v string) *DeployHuggingFaceModelOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployHuggingFaceModelOutput) SetRequestId(v string) *DeployHuggingFaceModelOutput {
	s.RequestId = &v
	return s
}

func (s *DeployHuggingFaceModelOutput) SetSuccess(v bool) *DeployHuggingFaceModelOutput {
	s.Success = &v
	return s
}

type DeployHuggingFaceModelOutputData struct {
	DeploymentTaskID *string `json:"deploymentTaskID,omitempty" xml:"deploymentTaskID,omitempty"`
	ErrorMessage     *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Finished         *bool   `json:"finished,omitempty" xml:"finished,omitempty"`
	ServiceName      *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	TaskType         *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	TraceID          *string `json:"traceID,omitempty" xml:"traceID,omitempty"`
	UrlInternet      *string `json:"urlInternet,omitempty" xml:"urlInternet,omitempty"`
	UrlIntranet      *string `json:"urlIntranet,omitempty" xml:"urlIntranet,omitempty"`
}

func (s DeployHuggingFaceModelOutputData) String() string {
	return tea.Prettify(s)
}

func (s DeployHuggingFaceModelOutputData) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelOutputData) SetDeploymentTaskID(v string) *DeployHuggingFaceModelOutputData {
	s.DeploymentTaskID = &v
	return s
}

func (s *DeployHuggingFaceModelOutputData) SetErrorMessage(v string) *DeployHuggingFaceModelOutputData {
	s.ErrorMessage = &v
	return s
}

func (s *DeployHuggingFaceModelOutputData) SetFinished(v bool) *DeployHuggingFaceModelOutputData {
	s.Finished = &v
	return s
}

func (s *DeployHuggingFaceModelOutputData) SetServiceName(v string) *DeployHuggingFaceModelOutputData {
	s.ServiceName = &v
	return s
}

func (s *DeployHuggingFaceModelOutputData) SetTaskType(v string) *DeployHuggingFaceModelOutputData {
	s.TaskType = &v
	return s
}

func (s *DeployHuggingFaceModelOutputData) SetTraceID(v string) *DeployHuggingFaceModelOutputData {
	s.TraceID = &v
	return s
}

func (s *DeployHuggingFaceModelOutputData) SetUrlInternet(v string) *DeployHuggingFaceModelOutputData {
	s.UrlInternet = &v
	return s
}

func (s *DeployHuggingFaceModelOutputData) SetUrlIntranet(v string) *DeployHuggingFaceModelOutputData {
	s.UrlIntranet = &v
	return s
}

type DeployModelScopeModelAsyncOutput struct {
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployModelScopeModelAsyncOutput) String() string {
	return tea.Prettify(s)
}

func (s DeployModelScopeModelAsyncOutput) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelAsyncOutput) SetData(v string) *DeployModelScopeModelAsyncOutput {
	s.Data = &v
	return s
}

func (s *DeployModelScopeModelAsyncOutput) SetErrCode(v string) *DeployModelScopeModelAsyncOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployModelScopeModelAsyncOutput) SetErrMsg(v string) *DeployModelScopeModelAsyncOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployModelScopeModelAsyncOutput) SetRequestId(v string) *DeployModelScopeModelAsyncOutput {
	s.RequestId = &v
	return s
}

func (s *DeployModelScopeModelAsyncOutput) SetSuccess(v bool) *DeployModelScopeModelAsyncOutput {
	s.Success = &v
	return s
}

type DeployModelScopeModelInput struct {
	AccountID            *string                                      `json:"accountID,omitempty" xml:"accountID,omitempty"`
	ConcurrencyConfig    *DeployModelScopeModelInputConcurrencyConfig `json:"concurrencyConfig,omitempty" xml:"concurrencyConfig,omitempty" type:"Struct"`
	Cpu                  *float32                                     `json:"cpu,omitempty" xml:"cpu,omitempty"`
	Description          *string                                      `json:"description,omitempty" xml:"description,omitempty"`
	DiskSize             *int32                                       `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnvName              *string                                      `json:"envName,omitempty" xml:"envName,omitempty"`
	EnvironmentVariables map[string]interface{}                       `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	GpuConfig            *DeployModelScopeModelInputGpuConfig         `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty" type:"Struct"`
	HttpTrigger          *DeployModelScopeModelInputHttpTrigger       `json:"httpTrigger,omitempty" xml:"httpTrigger,omitempty" type:"Struct"`
	ImageName            *string                                      `json:"imageName,omitempty" xml:"imageName,omitempty"`
	InstanceConcurrency  *int32                                       `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	LogConfig            *DeployModelScopeModelInputLogConfig         `json:"logConfig,omitempty" xml:"logConfig,omitempty" type:"Struct"`
	MemorySize           *int32                                       `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	ModelConfig          *DeployModelScopeModelInputModelConfig       `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Name            *string                                    `json:"name,omitempty" xml:"name,omitempty"`
	NasConfig       *DeployModelScopeModelInputNasConfig       `json:"nasConfig,omitempty" xml:"nasConfig,omitempty" type:"Struct"`
	OriginalName    *string                                    `json:"originalName,omitempty" xml:"originalName,omitempty"`
	ProjectName     *string                                    `json:"projectName,omitempty" xml:"projectName,omitempty"`
	ProvisionConfig *DeployModelScopeModelInputProvisionConfig `json:"provisionConfig,omitempty" xml:"provisionConfig,omitempty" type:"Struct"`
	Region          *string                                    `json:"region,omitempty" xml:"region,omitempty"`
	ReportStatusURL *string                                    `json:"reportStatusURL,omitempty" xml:"reportStatusURL,omitempty"`
	// This parameter is required.
	Role      *string                              `json:"role,omitempty" xml:"role,omitempty"`
	Timeout   *int32                               `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TraceId   *string                              `json:"traceId,omitempty" xml:"traceId,omitempty"`
	VpcConfig *DeployModelScopeModelInputVpcConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty" type:"Struct"`
}

func (s DeployModelScopeModelInput) String() string {
	return tea.Prettify(s)
}

func (s DeployModelScopeModelInput) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInput) SetAccountID(v string) *DeployModelScopeModelInput {
	s.AccountID = &v
	return s
}

func (s *DeployModelScopeModelInput) SetConcurrencyConfig(v *DeployModelScopeModelInputConcurrencyConfig) *DeployModelScopeModelInput {
	s.ConcurrencyConfig = v
	return s
}

func (s *DeployModelScopeModelInput) SetCpu(v float32) *DeployModelScopeModelInput {
	s.Cpu = &v
	return s
}

func (s *DeployModelScopeModelInput) SetDescription(v string) *DeployModelScopeModelInput {
	s.Description = &v
	return s
}

func (s *DeployModelScopeModelInput) SetDiskSize(v int32) *DeployModelScopeModelInput {
	s.DiskSize = &v
	return s
}

func (s *DeployModelScopeModelInput) SetEnvName(v string) *DeployModelScopeModelInput {
	s.EnvName = &v
	return s
}

func (s *DeployModelScopeModelInput) SetEnvironmentVariables(v map[string]interface{}) *DeployModelScopeModelInput {
	s.EnvironmentVariables = v
	return s
}

func (s *DeployModelScopeModelInput) SetGpuConfig(v *DeployModelScopeModelInputGpuConfig) *DeployModelScopeModelInput {
	s.GpuConfig = v
	return s
}

func (s *DeployModelScopeModelInput) SetHttpTrigger(v *DeployModelScopeModelInputHttpTrigger) *DeployModelScopeModelInput {
	s.HttpTrigger = v
	return s
}

func (s *DeployModelScopeModelInput) SetImageName(v string) *DeployModelScopeModelInput {
	s.ImageName = &v
	return s
}

func (s *DeployModelScopeModelInput) SetInstanceConcurrency(v int32) *DeployModelScopeModelInput {
	s.InstanceConcurrency = &v
	return s
}

func (s *DeployModelScopeModelInput) SetLogConfig(v *DeployModelScopeModelInputLogConfig) *DeployModelScopeModelInput {
	s.LogConfig = v
	return s
}

func (s *DeployModelScopeModelInput) SetMemorySize(v int32) *DeployModelScopeModelInput {
	s.MemorySize = &v
	return s
}

func (s *DeployModelScopeModelInput) SetModelConfig(v *DeployModelScopeModelInputModelConfig) *DeployModelScopeModelInput {
	s.ModelConfig = v
	return s
}

func (s *DeployModelScopeModelInput) SetName(v string) *DeployModelScopeModelInput {
	s.Name = &v
	return s
}

func (s *DeployModelScopeModelInput) SetNasConfig(v *DeployModelScopeModelInputNasConfig) *DeployModelScopeModelInput {
	s.NasConfig = v
	return s
}

func (s *DeployModelScopeModelInput) SetOriginalName(v string) *DeployModelScopeModelInput {
	s.OriginalName = &v
	return s
}

func (s *DeployModelScopeModelInput) SetProjectName(v string) *DeployModelScopeModelInput {
	s.ProjectName = &v
	return s
}

func (s *DeployModelScopeModelInput) SetProvisionConfig(v *DeployModelScopeModelInputProvisionConfig) *DeployModelScopeModelInput {
	s.ProvisionConfig = v
	return s
}

func (s *DeployModelScopeModelInput) SetRegion(v string) *DeployModelScopeModelInput {
	s.Region = &v
	return s
}

func (s *DeployModelScopeModelInput) SetReportStatusURL(v string) *DeployModelScopeModelInput {
	s.ReportStatusURL = &v
	return s
}

func (s *DeployModelScopeModelInput) SetRole(v string) *DeployModelScopeModelInput {
	s.Role = &v
	return s
}

func (s *DeployModelScopeModelInput) SetTimeout(v int32) *DeployModelScopeModelInput {
	s.Timeout = &v
	return s
}

func (s *DeployModelScopeModelInput) SetTraceId(v string) *DeployModelScopeModelInput {
	s.TraceId = &v
	return s
}

func (s *DeployModelScopeModelInput) SetVpcConfig(v *DeployModelScopeModelInputVpcConfig) *DeployModelScopeModelInput {
	s.VpcConfig = v
	return s
}

type DeployModelScopeModelInputConcurrencyConfig struct {
	ReservedConcurrency *int32 `json:"reservedConcurrency,omitempty" xml:"reservedConcurrency,omitempty"`
}

func (s DeployModelScopeModelInputConcurrencyConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployModelScopeModelInputConcurrencyConfig) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInputConcurrencyConfig) SetReservedConcurrency(v int32) *DeployModelScopeModelInputConcurrencyConfig {
	s.ReservedConcurrency = &v
	return s
}

type DeployModelScopeModelInputGpuConfig struct {
	GpuMemorySize *int32  `json:"gpuMemorySize,omitempty" xml:"gpuMemorySize,omitempty"`
	GpuType       *string `json:"gpuType,omitempty" xml:"gpuType,omitempty"`
}

func (s DeployModelScopeModelInputGpuConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployModelScopeModelInputGpuConfig) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInputGpuConfig) SetGpuMemorySize(v int32) *DeployModelScopeModelInputGpuConfig {
	s.GpuMemorySize = &v
	return s
}

func (s *DeployModelScopeModelInputGpuConfig) SetGpuType(v string) *DeployModelScopeModelInputGpuConfig {
	s.GpuType = &v
	return s
}

type DeployModelScopeModelInputHttpTrigger struct {
	Qualifier     *string                                             `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	TriggerConfig *DeployModelScopeModelInputHttpTriggerTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
}

func (s DeployModelScopeModelInputHttpTrigger) String() string {
	return tea.Prettify(s)
}

func (s DeployModelScopeModelInputHttpTrigger) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInputHttpTrigger) SetQualifier(v string) *DeployModelScopeModelInputHttpTrigger {
	s.Qualifier = &v
	return s
}

func (s *DeployModelScopeModelInputHttpTrigger) SetTriggerConfig(v *DeployModelScopeModelInputHttpTriggerTriggerConfig) *DeployModelScopeModelInputHttpTrigger {
	s.TriggerConfig = v
	return s
}

type DeployModelScopeModelInputHttpTriggerTriggerConfig struct {
	AuthType          *string   `json:"authType,omitempty" xml:"authType,omitempty"`
	DsableURLInternet *bool     `json:"dsableURLInternet,omitempty" xml:"dsableURLInternet,omitempty"`
	Methods           []*string `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
}

func (s DeployModelScopeModelInputHttpTriggerTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployModelScopeModelInputHttpTriggerTriggerConfig) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInputHttpTriggerTriggerConfig) SetAuthType(v string) *DeployModelScopeModelInputHttpTriggerTriggerConfig {
	s.AuthType = &v
	return s
}

func (s *DeployModelScopeModelInputHttpTriggerTriggerConfig) SetDsableURLInternet(v bool) *DeployModelScopeModelInputHttpTriggerTriggerConfig {
	s.DsableURLInternet = &v
	return s
}

func (s *DeployModelScopeModelInputHttpTriggerTriggerConfig) SetMethods(v []*string) *DeployModelScopeModelInputHttpTriggerTriggerConfig {
	s.Methods = v
	return s
}

type DeployModelScopeModelInputLogConfig struct {
	EnableInstanceMetrics *bool   `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	EnableRequestMetrics  *bool   `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	LogBeginRule          *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	Logstore              *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	Project               *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s DeployModelScopeModelInputLogConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployModelScopeModelInputLogConfig) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInputLogConfig) SetEnableInstanceMetrics(v bool) *DeployModelScopeModelInputLogConfig {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *DeployModelScopeModelInputLogConfig) SetEnableRequestMetrics(v bool) *DeployModelScopeModelInputLogConfig {
	s.EnableRequestMetrics = &v
	return s
}

func (s *DeployModelScopeModelInputLogConfig) SetLogBeginRule(v string) *DeployModelScopeModelInputLogConfig {
	s.LogBeginRule = &v
	return s
}

func (s *DeployModelScopeModelInputLogConfig) SetLogstore(v string) *DeployModelScopeModelInputLogConfig {
	s.Logstore = &v
	return s
}

func (s *DeployModelScopeModelInputLogConfig) SetProject(v string) *DeployModelScopeModelInputLogConfig {
	s.Project = &v
	return s
}

type DeployModelScopeModelInputModelConfig struct {
	Framework *string `json:"framework,omitempty" xml:"framework,omitempty"`
	// if can be null:
	// true
	MultiModelConfig           []*ModelConfig `json:"multiModelConfig,omitempty" xml:"multiModelConfig,omitempty" type:"Repeated"`
	Prefix                     *string        `json:"prefix,omitempty" xml:"prefix,omitempty"`
	SourceType                 *string        `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	SrcModelScopeModelID       *string        `json:"srcModelScopeModelID,omitempty" xml:"srcModelScopeModelID,omitempty"`
	SrcModelScopeModelRevision *string        `json:"srcModelScopeModelRevision,omitempty" xml:"srcModelScopeModelRevision,omitempty"`
	SrcModelScopeToken         *string        `json:"srcModelScopeToken,omitempty" xml:"srcModelScopeToken,omitempty"`
	SrcOssBucket               *string        `json:"srcOssBucket,omitempty" xml:"srcOssBucket,omitempty"`
	SrcOssPath                 *string        `json:"srcOssPath,omitempty" xml:"srcOssPath,omitempty"`
	SrcOssRegion               *string        `json:"srcOssRegion,omitempty" xml:"srcOssRegion,omitempty"`
}

func (s DeployModelScopeModelInputModelConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployModelScopeModelInputModelConfig) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInputModelConfig) SetFramework(v string) *DeployModelScopeModelInputModelConfig {
	s.Framework = &v
	return s
}

func (s *DeployModelScopeModelInputModelConfig) SetMultiModelConfig(v []*ModelConfig) *DeployModelScopeModelInputModelConfig {
	s.MultiModelConfig = v
	return s
}

func (s *DeployModelScopeModelInputModelConfig) SetPrefix(v string) *DeployModelScopeModelInputModelConfig {
	s.Prefix = &v
	return s
}

func (s *DeployModelScopeModelInputModelConfig) SetSourceType(v string) *DeployModelScopeModelInputModelConfig {
	s.SourceType = &v
	return s
}

func (s *DeployModelScopeModelInputModelConfig) SetSrcModelScopeModelID(v string) *DeployModelScopeModelInputModelConfig {
	s.SrcModelScopeModelID = &v
	return s
}

func (s *DeployModelScopeModelInputModelConfig) SetSrcModelScopeModelRevision(v string) *DeployModelScopeModelInputModelConfig {
	s.SrcModelScopeModelRevision = &v
	return s
}

func (s *DeployModelScopeModelInputModelConfig) SetSrcModelScopeToken(v string) *DeployModelScopeModelInputModelConfig {
	s.SrcModelScopeToken = &v
	return s
}

func (s *DeployModelScopeModelInputModelConfig) SetSrcOssBucket(v string) *DeployModelScopeModelInputModelConfig {
	s.SrcOssBucket = &v
	return s
}

func (s *DeployModelScopeModelInputModelConfig) SetSrcOssPath(v string) *DeployModelScopeModelInputModelConfig {
	s.SrcOssPath = &v
	return s
}

func (s *DeployModelScopeModelInputModelConfig) SetSrcOssRegion(v string) *DeployModelScopeModelInputModelConfig {
	s.SrcOssRegion = &v
	return s
}

type DeployModelScopeModelInputNasConfig struct {
	GroupId     *int32    `json:"groupId,omitempty" xml:"groupId,omitempty"`
	MountPoints []*string `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	UserId      *int32    `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeployModelScopeModelInputNasConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployModelScopeModelInputNasConfig) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInputNasConfig) SetGroupId(v int32) *DeployModelScopeModelInputNasConfig {
	s.GroupId = &v
	return s
}

func (s *DeployModelScopeModelInputNasConfig) SetMountPoints(v []*string) *DeployModelScopeModelInputNasConfig {
	s.MountPoints = v
	return s
}

func (s *DeployModelScopeModelInputNasConfig) SetUserId(v int32) *DeployModelScopeModelInputNasConfig {
	s.UserId = &v
	return s
}

type DeployModelScopeModelInputProvisionConfig struct {
	AlwaysAllocateGPU *bool                                                        `json:"alwaysAllocateGPU,omitempty" xml:"alwaysAllocateGPU,omitempty"`
	ScheduledActions  []*DeployModelScopeModelInputProvisionConfigScheduledActions `json:"scheduledActions,omitempty" xml:"scheduledActions,omitempty" type:"Repeated"`
	Target            *int32                                                       `json:"target,omitempty" xml:"target,omitempty"`
}

func (s DeployModelScopeModelInputProvisionConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployModelScopeModelInputProvisionConfig) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInputProvisionConfig) SetAlwaysAllocateGPU(v bool) *DeployModelScopeModelInputProvisionConfig {
	s.AlwaysAllocateGPU = &v
	return s
}

func (s *DeployModelScopeModelInputProvisionConfig) SetScheduledActions(v []*DeployModelScopeModelInputProvisionConfigScheduledActions) *DeployModelScopeModelInputProvisionConfig {
	s.ScheduledActions = v
	return s
}

func (s *DeployModelScopeModelInputProvisionConfig) SetTarget(v int32) *DeployModelScopeModelInputProvisionConfig {
	s.Target = &v
	return s
}

type DeployModelScopeModelInputProvisionConfigScheduledActions struct {
	EndTime            *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	ScheduleExpression *string `json:"scheduleExpression,omitempty" xml:"scheduleExpression,omitempty"`
	StartTime          *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Target             *int32  `json:"target,omitempty" xml:"target,omitempty"`
	TimeZone           *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s DeployModelScopeModelInputProvisionConfigScheduledActions) String() string {
	return tea.Prettify(s)
}

func (s DeployModelScopeModelInputProvisionConfigScheduledActions) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInputProvisionConfigScheduledActions) SetEndTime(v string) *DeployModelScopeModelInputProvisionConfigScheduledActions {
	s.EndTime = &v
	return s
}

func (s *DeployModelScopeModelInputProvisionConfigScheduledActions) SetName(v string) *DeployModelScopeModelInputProvisionConfigScheduledActions {
	s.Name = &v
	return s
}

func (s *DeployModelScopeModelInputProvisionConfigScheduledActions) SetScheduleExpression(v string) *DeployModelScopeModelInputProvisionConfigScheduledActions {
	s.ScheduleExpression = &v
	return s
}

func (s *DeployModelScopeModelInputProvisionConfigScheduledActions) SetStartTime(v string) *DeployModelScopeModelInputProvisionConfigScheduledActions {
	s.StartTime = &v
	return s
}

func (s *DeployModelScopeModelInputProvisionConfigScheduledActions) SetTarget(v int32) *DeployModelScopeModelInputProvisionConfigScheduledActions {
	s.Target = &v
	return s
}

func (s *DeployModelScopeModelInputProvisionConfigScheduledActions) SetTimeZone(v string) *DeployModelScopeModelInputProvisionConfigScheduledActions {
	s.TimeZone = &v
	return s
}

type DeployModelScopeModelInputVpcConfig struct {
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	VpcId           *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s DeployModelScopeModelInputVpcConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployModelScopeModelInputVpcConfig) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInputVpcConfig) SetSecurityGroupId(v string) *DeployModelScopeModelInputVpcConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *DeployModelScopeModelInputVpcConfig) SetVSwitchIds(v []*string) *DeployModelScopeModelInputVpcConfig {
	s.VSwitchIds = v
	return s
}

func (s *DeployModelScopeModelInputVpcConfig) SetVpcId(v string) *DeployModelScopeModelInputVpcConfig {
	s.VpcId = &v
	return s
}

type DeployModelScopeModelOutput struct {
	Data    *DeployModelScopeModelOutputData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                          `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                          `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployModelScopeModelOutput) String() string {
	return tea.Prettify(s)
}

func (s DeployModelScopeModelOutput) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelOutput) SetData(v *DeployModelScopeModelOutputData) *DeployModelScopeModelOutput {
	s.Data = v
	return s
}

func (s *DeployModelScopeModelOutput) SetErrCode(v string) *DeployModelScopeModelOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployModelScopeModelOutput) SetErrMsg(v string) *DeployModelScopeModelOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployModelScopeModelOutput) SetRequestId(v string) *DeployModelScopeModelOutput {
	s.RequestId = &v
	return s
}

func (s *DeployModelScopeModelOutput) SetSuccess(v bool) *DeployModelScopeModelOutput {
	s.Success = &v
	return s
}

type DeployModelScopeModelOutputData struct {
	DeploymentTaskID *string `json:"deploymentTaskID,omitempty" xml:"deploymentTaskID,omitempty"`
	ErrorMessage     *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Finished         *bool   `json:"finished,omitempty" xml:"finished,omitempty"`
	ServiceName      *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	TaskType         *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	TraceID          *string `json:"traceID,omitempty" xml:"traceID,omitempty"`
	UrlInternet      *string `json:"urlInternet,omitempty" xml:"urlInternet,omitempty"`
	UrlIntranet      *string `json:"urlIntranet,omitempty" xml:"urlIntranet,omitempty"`
}

func (s DeployModelScopeModelOutputData) String() string {
	return tea.Prettify(s)
}

func (s DeployModelScopeModelOutputData) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelOutputData) SetDeploymentTaskID(v string) *DeployModelScopeModelOutputData {
	s.DeploymentTaskID = &v
	return s
}

func (s *DeployModelScopeModelOutputData) SetErrorMessage(v string) *DeployModelScopeModelOutputData {
	s.ErrorMessage = &v
	return s
}

func (s *DeployModelScopeModelOutputData) SetFinished(v bool) *DeployModelScopeModelOutputData {
	s.Finished = &v
	return s
}

func (s *DeployModelScopeModelOutputData) SetServiceName(v string) *DeployModelScopeModelOutputData {
	s.ServiceName = &v
	return s
}

func (s *DeployModelScopeModelOutputData) SetTaskType(v string) *DeployModelScopeModelOutputData {
	s.TaskType = &v
	return s
}

func (s *DeployModelScopeModelOutputData) SetTraceID(v string) *DeployModelScopeModelOutputData {
	s.TraceID = &v
	return s
}

func (s *DeployModelScopeModelOutputData) SetUrlInternet(v string) *DeployModelScopeModelOutputData {
	s.UrlInternet = &v
	return s
}

func (s *DeployModelScopeModelOutputData) SetUrlIntranet(v string) *DeployModelScopeModelOutputData {
	s.UrlIntranet = &v
	return s
}

type DeployOllamaModelAsyncOutput struct {
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployOllamaModelAsyncOutput) String() string {
	return tea.Prettify(s)
}

func (s DeployOllamaModelAsyncOutput) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelAsyncOutput) SetData(v string) *DeployOllamaModelAsyncOutput {
	s.Data = &v
	return s
}

func (s *DeployOllamaModelAsyncOutput) SetErrCode(v string) *DeployOllamaModelAsyncOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployOllamaModelAsyncOutput) SetErrMsg(v string) *DeployOllamaModelAsyncOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployOllamaModelAsyncOutput) SetRequestId(v string) *DeployOllamaModelAsyncOutput {
	s.RequestId = &v
	return s
}

func (s *DeployOllamaModelAsyncOutput) SetSuccess(v bool) *DeployOllamaModelAsyncOutput {
	s.Success = &v
	return s
}

type DeployOllamaModelInput struct {
	AccountID            *string                                  `json:"accountID,omitempty" xml:"accountID,omitempty"`
	ConcurrencyConfig    *DeployOllamaModelInputConcurrencyConfig `json:"concurrencyConfig,omitempty" xml:"concurrencyConfig,omitempty" type:"Struct"`
	Cpu                  *float32                                 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	Description          *string                                  `json:"description,omitempty" xml:"description,omitempty"`
	DiskSize             *int32                                   `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnvName              *string                                  `json:"envName,omitempty" xml:"envName,omitempty"`
	EnvironmentVariables map[string]interface{}                   `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	GpuConfig            *DeployOllamaModelInputGpuConfig         `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty" type:"Struct"`
	HttpTrigger          *DeployOllamaModelInputHttpTrigger       `json:"httpTrigger,omitempty" xml:"httpTrigger,omitempty" type:"Struct"`
	ImageName            *string                                  `json:"imageName,omitempty" xml:"imageName,omitempty"`
	InstanceConcurrency  *int32                                   `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	LogConfig            *DeployOllamaModelInputLogConfig         `json:"logConfig,omitempty" xml:"logConfig,omitempty" type:"Struct"`
	MemorySize           *int32                                   `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	ModelConfig          *DeployOllamaModelInputModelConfig       `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Name            *string                                `json:"name,omitempty" xml:"name,omitempty"`
	NasConfig       *DeployOllamaModelInputNasConfig       `json:"nasConfig,omitempty" xml:"nasConfig,omitempty" type:"Struct"`
	OriginalName    *string                                `json:"originalName,omitempty" xml:"originalName,omitempty"`
	ProjectName     *string                                `json:"projectName,omitempty" xml:"projectName,omitempty"`
	ProvisionConfig *DeployOllamaModelInputProvisionConfig `json:"provisionConfig,omitempty" xml:"provisionConfig,omitempty" type:"Struct"`
	Region          *string                                `json:"region,omitempty" xml:"region,omitempty"`
	ReportStatusURL *string                                `json:"reportStatusURL,omitempty" xml:"reportStatusURL,omitempty"`
	// This parameter is required.
	Role      *string                          `json:"role,omitempty" xml:"role,omitempty"`
	Timeout   *int32                           `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TraceId   *string                          `json:"traceId,omitempty" xml:"traceId,omitempty"`
	VpcConfig *DeployOllamaModelInputVpcConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty" type:"Struct"`
}

func (s DeployOllamaModelInput) String() string {
	return tea.Prettify(s)
}

func (s DeployOllamaModelInput) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInput) SetAccountID(v string) *DeployOllamaModelInput {
	s.AccountID = &v
	return s
}

func (s *DeployOllamaModelInput) SetConcurrencyConfig(v *DeployOllamaModelInputConcurrencyConfig) *DeployOllamaModelInput {
	s.ConcurrencyConfig = v
	return s
}

func (s *DeployOllamaModelInput) SetCpu(v float32) *DeployOllamaModelInput {
	s.Cpu = &v
	return s
}

func (s *DeployOllamaModelInput) SetDescription(v string) *DeployOllamaModelInput {
	s.Description = &v
	return s
}

func (s *DeployOllamaModelInput) SetDiskSize(v int32) *DeployOllamaModelInput {
	s.DiskSize = &v
	return s
}

func (s *DeployOllamaModelInput) SetEnvName(v string) *DeployOllamaModelInput {
	s.EnvName = &v
	return s
}

func (s *DeployOllamaModelInput) SetEnvironmentVariables(v map[string]interface{}) *DeployOllamaModelInput {
	s.EnvironmentVariables = v
	return s
}

func (s *DeployOllamaModelInput) SetGpuConfig(v *DeployOllamaModelInputGpuConfig) *DeployOllamaModelInput {
	s.GpuConfig = v
	return s
}

func (s *DeployOllamaModelInput) SetHttpTrigger(v *DeployOllamaModelInputHttpTrigger) *DeployOllamaModelInput {
	s.HttpTrigger = v
	return s
}

func (s *DeployOllamaModelInput) SetImageName(v string) *DeployOllamaModelInput {
	s.ImageName = &v
	return s
}

func (s *DeployOllamaModelInput) SetInstanceConcurrency(v int32) *DeployOllamaModelInput {
	s.InstanceConcurrency = &v
	return s
}

func (s *DeployOllamaModelInput) SetLogConfig(v *DeployOllamaModelInputLogConfig) *DeployOllamaModelInput {
	s.LogConfig = v
	return s
}

func (s *DeployOllamaModelInput) SetMemorySize(v int32) *DeployOllamaModelInput {
	s.MemorySize = &v
	return s
}

func (s *DeployOllamaModelInput) SetModelConfig(v *DeployOllamaModelInputModelConfig) *DeployOllamaModelInput {
	s.ModelConfig = v
	return s
}

func (s *DeployOllamaModelInput) SetName(v string) *DeployOllamaModelInput {
	s.Name = &v
	return s
}

func (s *DeployOllamaModelInput) SetNasConfig(v *DeployOllamaModelInputNasConfig) *DeployOllamaModelInput {
	s.NasConfig = v
	return s
}

func (s *DeployOllamaModelInput) SetOriginalName(v string) *DeployOllamaModelInput {
	s.OriginalName = &v
	return s
}

func (s *DeployOllamaModelInput) SetProjectName(v string) *DeployOllamaModelInput {
	s.ProjectName = &v
	return s
}

func (s *DeployOllamaModelInput) SetProvisionConfig(v *DeployOllamaModelInputProvisionConfig) *DeployOllamaModelInput {
	s.ProvisionConfig = v
	return s
}

func (s *DeployOllamaModelInput) SetRegion(v string) *DeployOllamaModelInput {
	s.Region = &v
	return s
}

func (s *DeployOllamaModelInput) SetReportStatusURL(v string) *DeployOllamaModelInput {
	s.ReportStatusURL = &v
	return s
}

func (s *DeployOllamaModelInput) SetRole(v string) *DeployOllamaModelInput {
	s.Role = &v
	return s
}

func (s *DeployOllamaModelInput) SetTimeout(v int32) *DeployOllamaModelInput {
	s.Timeout = &v
	return s
}

func (s *DeployOllamaModelInput) SetTraceId(v string) *DeployOllamaModelInput {
	s.TraceId = &v
	return s
}

func (s *DeployOllamaModelInput) SetVpcConfig(v *DeployOllamaModelInputVpcConfig) *DeployOllamaModelInput {
	s.VpcConfig = v
	return s
}

type DeployOllamaModelInputConcurrencyConfig struct {
	ReservedConcurrency *int32 `json:"reservedConcurrency,omitempty" xml:"reservedConcurrency,omitempty"`
}

func (s DeployOllamaModelInputConcurrencyConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployOllamaModelInputConcurrencyConfig) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputConcurrencyConfig) SetReservedConcurrency(v int32) *DeployOllamaModelInputConcurrencyConfig {
	s.ReservedConcurrency = &v
	return s
}

type DeployOllamaModelInputGpuConfig struct {
	GpuMemorySize *int32  `json:"gpuMemorySize,omitempty" xml:"gpuMemorySize,omitempty"`
	GpuType       *string `json:"gpuType,omitempty" xml:"gpuType,omitempty"`
}

func (s DeployOllamaModelInputGpuConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployOllamaModelInputGpuConfig) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputGpuConfig) SetGpuMemorySize(v int32) *DeployOllamaModelInputGpuConfig {
	s.GpuMemorySize = &v
	return s
}

func (s *DeployOllamaModelInputGpuConfig) SetGpuType(v string) *DeployOllamaModelInputGpuConfig {
	s.GpuType = &v
	return s
}

type DeployOllamaModelInputHttpTrigger struct {
	Qualifier     *string                                         `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	TriggerConfig *DeployOllamaModelInputHttpTriggerTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
}

func (s DeployOllamaModelInputHttpTrigger) String() string {
	return tea.Prettify(s)
}

func (s DeployOllamaModelInputHttpTrigger) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputHttpTrigger) SetQualifier(v string) *DeployOllamaModelInputHttpTrigger {
	s.Qualifier = &v
	return s
}

func (s *DeployOllamaModelInputHttpTrigger) SetTriggerConfig(v *DeployOllamaModelInputHttpTriggerTriggerConfig) *DeployOllamaModelInputHttpTrigger {
	s.TriggerConfig = v
	return s
}

type DeployOllamaModelInputHttpTriggerTriggerConfig struct {
	AuthType          *string   `json:"authType,omitempty" xml:"authType,omitempty"`
	DsableURLInternet *bool     `json:"dsableURLInternet,omitempty" xml:"dsableURLInternet,omitempty"`
	Methods           []*string `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
}

func (s DeployOllamaModelInputHttpTriggerTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployOllamaModelInputHttpTriggerTriggerConfig) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputHttpTriggerTriggerConfig) SetAuthType(v string) *DeployOllamaModelInputHttpTriggerTriggerConfig {
	s.AuthType = &v
	return s
}

func (s *DeployOllamaModelInputHttpTriggerTriggerConfig) SetDsableURLInternet(v bool) *DeployOllamaModelInputHttpTriggerTriggerConfig {
	s.DsableURLInternet = &v
	return s
}

func (s *DeployOllamaModelInputHttpTriggerTriggerConfig) SetMethods(v []*string) *DeployOllamaModelInputHttpTriggerTriggerConfig {
	s.Methods = v
	return s
}

type DeployOllamaModelInputLogConfig struct {
	EnableInstanceMetrics *bool   `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	EnableRequestMetrics  *bool   `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	LogBeginRule          *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	Logstore              *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	Project               *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s DeployOllamaModelInputLogConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployOllamaModelInputLogConfig) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputLogConfig) SetEnableInstanceMetrics(v bool) *DeployOllamaModelInputLogConfig {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *DeployOllamaModelInputLogConfig) SetEnableRequestMetrics(v bool) *DeployOllamaModelInputLogConfig {
	s.EnableRequestMetrics = &v
	return s
}

func (s *DeployOllamaModelInputLogConfig) SetLogBeginRule(v string) *DeployOllamaModelInputLogConfig {
	s.LogBeginRule = &v
	return s
}

func (s *DeployOllamaModelInputLogConfig) SetLogstore(v string) *DeployOllamaModelInputLogConfig {
	s.Logstore = &v
	return s
}

func (s *DeployOllamaModelInputLogConfig) SetProject(v string) *DeployOllamaModelInputLogConfig {
	s.Project = &v
	return s
}

type DeployOllamaModelInputModelConfig struct {
	FmkOllamaConfig *DeployOllamaModelInputModelConfigFmkOllamaConfig `json:"fmkOllamaConfig,omitempty" xml:"fmkOllamaConfig,omitempty" type:"Struct"`
	Framework       *string                                           `json:"framework,omitempty" xml:"framework,omitempty"`
	// if can be null:
	// true
	MultiModelConfig           []*ModelConfig `json:"multiModelConfig,omitempty" xml:"multiModelConfig,omitempty" type:"Repeated"`
	Prefix                     *string        `json:"prefix,omitempty" xml:"prefix,omitempty"`
	SourceType                 *string        `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	SrcModelScopeModelID       *string        `json:"srcModelScopeModelID,omitempty" xml:"srcModelScopeModelID,omitempty"`
	SrcModelScopeModelRevision *string        `json:"srcModelScopeModelRevision,omitempty" xml:"srcModelScopeModelRevision,omitempty"`
	SrcModelScopeToken         *string        `json:"srcModelScopeToken,omitempty" xml:"srcModelScopeToken,omitempty"`
	SrcOssBucket               *string        `json:"srcOssBucket,omitempty" xml:"srcOssBucket,omitempty"`
	SrcOssPath                 *string        `json:"srcOssPath,omitempty" xml:"srcOssPath,omitempty"`
	SrcOssRegion               *string        `json:"srcOssRegion,omitempty" xml:"srcOssRegion,omitempty"`
}

func (s DeployOllamaModelInputModelConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployOllamaModelInputModelConfig) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputModelConfig) SetFmkOllamaConfig(v *DeployOllamaModelInputModelConfigFmkOllamaConfig) *DeployOllamaModelInputModelConfig {
	s.FmkOllamaConfig = v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetFramework(v string) *DeployOllamaModelInputModelConfig {
	s.Framework = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetMultiModelConfig(v []*ModelConfig) *DeployOllamaModelInputModelConfig {
	s.MultiModelConfig = v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetPrefix(v string) *DeployOllamaModelInputModelConfig {
	s.Prefix = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetSourceType(v string) *DeployOllamaModelInputModelConfig {
	s.SourceType = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetSrcModelScopeModelID(v string) *DeployOllamaModelInputModelConfig {
	s.SrcModelScopeModelID = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetSrcModelScopeModelRevision(v string) *DeployOllamaModelInputModelConfig {
	s.SrcModelScopeModelRevision = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetSrcModelScopeToken(v string) *DeployOllamaModelInputModelConfig {
	s.SrcModelScopeToken = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetSrcOssBucket(v string) *DeployOllamaModelInputModelConfig {
	s.SrcOssBucket = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetSrcOssPath(v string) *DeployOllamaModelInputModelConfig {
	s.SrcOssPath = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetSrcOssRegion(v string) *DeployOllamaModelInputModelConfig {
	s.SrcOssRegion = &v
	return s
}

type DeployOllamaModelInputModelConfigFmkOllamaConfig struct {
	MinP                           *float32 `json:"minP,omitempty" xml:"minP,omitempty"`
	Mirostat                       *int32   `json:"mirostat,omitempty" xml:"mirostat,omitempty"`
	MirostatEta                    *float32 `json:"mirostatEta,omitempty" xml:"mirostatEta,omitempty"`
	MirostatTau                    *float32 `json:"mirostatTau,omitempty" xml:"mirostatTau,omitempty"`
	ModelName                      *string  `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ModelfileAdapter               *string  `json:"modelfileAdapter,omitempty" xml:"modelfileAdapter,omitempty"`
	ModelfileAdditionalFromsString *string  `json:"modelfileAdditionalFromsString,omitempty" xml:"modelfileAdditionalFromsString,omitempty"`
	ModelfileFullTextPostfix       *string  `json:"modelfileFullTextPostfix,omitempty" xml:"modelfileFullTextPostfix,omitempty"`
	ModelfileParams                *string  `json:"modelfileParams,omitempty" xml:"modelfileParams,omitempty"`
	ModelfileSystem                *string  `json:"modelfileSystem,omitempty" xml:"modelfileSystem,omitempty"`
	ModelfileTemplate              *string  `json:"modelfileTemplate,omitempty" xml:"modelfileTemplate,omitempty"`
	NumCtx                         *int32   `json:"numCtx,omitempty" xml:"numCtx,omitempty"`
	NumPredict                     *int32   `json:"numPredict,omitempty" xml:"numPredict,omitempty"`
	Quantize                       *string  `json:"quantize,omitempty" xml:"quantize,omitempty"`
	RepeatLastN                    *int32   `json:"repeatLastN,omitempty" xml:"repeatLastN,omitempty"`
	RepeatPenalty                  *float32 `json:"repeatPenalty,omitempty" xml:"repeatPenalty,omitempty"`
	Seed                           *int32   `json:"seed,omitempty" xml:"seed,omitempty"`
	SingleModelFile                *string  `json:"singleModelFile,omitempty" xml:"singleModelFile,omitempty"`
	SplitedModelStartFile          *string  `json:"splitedModelStartFile,omitempty" xml:"splitedModelStartFile,omitempty"`
	Stop                           *string  `json:"stop,omitempty" xml:"stop,omitempty"`
	Stream                         *bool    `json:"stream,omitempty" xml:"stream,omitempty"`
	Temperature                    *float32 `json:"temperature,omitempty" xml:"temperature,omitempty"`
	TfsZ                           *float32 `json:"tfsZ,omitempty" xml:"tfsZ,omitempty"`
	TopK                           *int32   `json:"topK,omitempty" xml:"topK,omitempty"`
	TopP                           *float32 `json:"topP,omitempty" xml:"topP,omitempty"`
}

func (s DeployOllamaModelInputModelConfigFmkOllamaConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployOllamaModelInputModelConfigFmkOllamaConfig) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetMinP(v float32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.MinP = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetMirostat(v int32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.Mirostat = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetMirostatEta(v float32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.MirostatEta = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetMirostatTau(v float32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.MirostatTau = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetModelName(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.ModelName = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetModelfileAdapter(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.ModelfileAdapter = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetModelfileAdditionalFromsString(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.ModelfileAdditionalFromsString = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetModelfileFullTextPostfix(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.ModelfileFullTextPostfix = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetModelfileParams(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.ModelfileParams = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetModelfileSystem(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.ModelfileSystem = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetModelfileTemplate(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.ModelfileTemplate = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetNumCtx(v int32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.NumCtx = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetNumPredict(v int32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.NumPredict = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetQuantize(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.Quantize = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetRepeatLastN(v int32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.RepeatLastN = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetRepeatPenalty(v float32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.RepeatPenalty = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetSeed(v int32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.Seed = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetSingleModelFile(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.SingleModelFile = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetSplitedModelStartFile(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.SplitedModelStartFile = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetStop(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.Stop = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetStream(v bool) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.Stream = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetTemperature(v float32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.Temperature = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetTfsZ(v float32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.TfsZ = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetTopK(v int32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.TopK = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetTopP(v float32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.TopP = &v
	return s
}

type DeployOllamaModelInputNasConfig struct {
	GroupId     *int32    `json:"groupId,omitempty" xml:"groupId,omitempty"`
	MountPoints []*string `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	UserId      *int32    `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeployOllamaModelInputNasConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployOllamaModelInputNasConfig) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputNasConfig) SetGroupId(v int32) *DeployOllamaModelInputNasConfig {
	s.GroupId = &v
	return s
}

func (s *DeployOllamaModelInputNasConfig) SetMountPoints(v []*string) *DeployOllamaModelInputNasConfig {
	s.MountPoints = v
	return s
}

func (s *DeployOllamaModelInputNasConfig) SetUserId(v int32) *DeployOllamaModelInputNasConfig {
	s.UserId = &v
	return s
}

type DeployOllamaModelInputProvisionConfig struct {
	AlwaysAllocateGPU *bool                                                    `json:"alwaysAllocateGPU,omitempty" xml:"alwaysAllocateGPU,omitempty"`
	ScheduledActions  []*DeployOllamaModelInputProvisionConfigScheduledActions `json:"scheduledActions,omitempty" xml:"scheduledActions,omitempty" type:"Repeated"`
	Target            *int32                                                   `json:"target,omitempty" xml:"target,omitempty"`
}

func (s DeployOllamaModelInputProvisionConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployOllamaModelInputProvisionConfig) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputProvisionConfig) SetAlwaysAllocateGPU(v bool) *DeployOllamaModelInputProvisionConfig {
	s.AlwaysAllocateGPU = &v
	return s
}

func (s *DeployOllamaModelInputProvisionConfig) SetScheduledActions(v []*DeployOllamaModelInputProvisionConfigScheduledActions) *DeployOllamaModelInputProvisionConfig {
	s.ScheduledActions = v
	return s
}

func (s *DeployOllamaModelInputProvisionConfig) SetTarget(v int32) *DeployOllamaModelInputProvisionConfig {
	s.Target = &v
	return s
}

type DeployOllamaModelInputProvisionConfigScheduledActions struct {
	EndTime            *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	ScheduleExpression *string `json:"scheduleExpression,omitempty" xml:"scheduleExpression,omitempty"`
	StartTime          *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Target             *int32  `json:"target,omitempty" xml:"target,omitempty"`
	TimeZone           *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s DeployOllamaModelInputProvisionConfigScheduledActions) String() string {
	return tea.Prettify(s)
}

func (s DeployOllamaModelInputProvisionConfigScheduledActions) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputProvisionConfigScheduledActions) SetEndTime(v string) *DeployOllamaModelInputProvisionConfigScheduledActions {
	s.EndTime = &v
	return s
}

func (s *DeployOllamaModelInputProvisionConfigScheduledActions) SetName(v string) *DeployOllamaModelInputProvisionConfigScheduledActions {
	s.Name = &v
	return s
}

func (s *DeployOllamaModelInputProvisionConfigScheduledActions) SetScheduleExpression(v string) *DeployOllamaModelInputProvisionConfigScheduledActions {
	s.ScheduleExpression = &v
	return s
}

func (s *DeployOllamaModelInputProvisionConfigScheduledActions) SetStartTime(v string) *DeployOllamaModelInputProvisionConfigScheduledActions {
	s.StartTime = &v
	return s
}

func (s *DeployOllamaModelInputProvisionConfigScheduledActions) SetTarget(v int32) *DeployOllamaModelInputProvisionConfigScheduledActions {
	s.Target = &v
	return s
}

func (s *DeployOllamaModelInputProvisionConfigScheduledActions) SetTimeZone(v string) *DeployOllamaModelInputProvisionConfigScheduledActions {
	s.TimeZone = &v
	return s
}

type DeployOllamaModelInputVpcConfig struct {
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	VpcId           *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s DeployOllamaModelInputVpcConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployOllamaModelInputVpcConfig) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputVpcConfig) SetSecurityGroupId(v string) *DeployOllamaModelInputVpcConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *DeployOllamaModelInputVpcConfig) SetVSwitchIds(v []*string) *DeployOllamaModelInputVpcConfig {
	s.VSwitchIds = v
	return s
}

func (s *DeployOllamaModelInputVpcConfig) SetVpcId(v string) *DeployOllamaModelInputVpcConfig {
	s.VpcId = &v
	return s
}

type DeployOllamaModelOutput struct {
	Data    *DeployOllamaModelOutputData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                      `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                      `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployOllamaModelOutput) String() string {
	return tea.Prettify(s)
}

func (s DeployOllamaModelOutput) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelOutput) SetData(v *DeployOllamaModelOutputData) *DeployOllamaModelOutput {
	s.Data = v
	return s
}

func (s *DeployOllamaModelOutput) SetErrCode(v string) *DeployOllamaModelOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployOllamaModelOutput) SetErrMsg(v string) *DeployOllamaModelOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployOllamaModelOutput) SetRequestId(v string) *DeployOllamaModelOutput {
	s.RequestId = &v
	return s
}

func (s *DeployOllamaModelOutput) SetSuccess(v bool) *DeployOllamaModelOutput {
	s.Success = &v
	return s
}

type DeployOllamaModelOutputData struct {
	DeploymentTaskID *string `json:"deploymentTaskID,omitempty" xml:"deploymentTaskID,omitempty"`
	ErrorMessage     *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Finished         *bool   `json:"finished,omitempty" xml:"finished,omitempty"`
	ModelName        *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ServiceName      *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	TraceID          *string `json:"traceID,omitempty" xml:"traceID,omitempty"`
	UrlInternet      *string `json:"urlInternet,omitempty" xml:"urlInternet,omitempty"`
	UrlIntranet      *string `json:"urlIntranet,omitempty" xml:"urlIntranet,omitempty"`
}

func (s DeployOllamaModelOutputData) String() string {
	return tea.Prettify(s)
}

func (s DeployOllamaModelOutputData) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelOutputData) SetDeploymentTaskID(v string) *DeployOllamaModelOutputData {
	s.DeploymentTaskID = &v
	return s
}

func (s *DeployOllamaModelOutputData) SetErrorMessage(v string) *DeployOllamaModelOutputData {
	s.ErrorMessage = &v
	return s
}

func (s *DeployOllamaModelOutputData) SetFinished(v bool) *DeployOllamaModelOutputData {
	s.Finished = &v
	return s
}

func (s *DeployOllamaModelOutputData) SetModelName(v string) *DeployOllamaModelOutputData {
	s.ModelName = &v
	return s
}

func (s *DeployOllamaModelOutputData) SetServiceName(v string) *DeployOllamaModelOutputData {
	s.ServiceName = &v
	return s
}

func (s *DeployOllamaModelOutputData) SetTraceID(v string) *DeployOllamaModelOutputData {
	s.TraceID = &v
	return s
}

func (s *DeployOllamaModelOutputData) SetUrlInternet(v string) *DeployOllamaModelOutputData {
	s.UrlInternet = &v
	return s
}

func (s *DeployOllamaModelOutputData) SetUrlIntranet(v string) *DeployOllamaModelOutputData {
	s.UrlIntranet = &v
	return s
}

type DeployTensorRtModelAsyncOutput struct {
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployTensorRtModelAsyncOutput) String() string {
	return tea.Prettify(s)
}

func (s DeployTensorRtModelAsyncOutput) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelAsyncOutput) SetData(v string) *DeployTensorRtModelAsyncOutput {
	s.Data = &v
	return s
}

func (s *DeployTensorRtModelAsyncOutput) SetErrCode(v string) *DeployTensorRtModelAsyncOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployTensorRtModelAsyncOutput) SetErrMsg(v string) *DeployTensorRtModelAsyncOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployTensorRtModelAsyncOutput) SetRequestId(v string) *DeployTensorRtModelAsyncOutput {
	s.RequestId = &v
	return s
}

func (s *DeployTensorRtModelAsyncOutput) SetSuccess(v bool) *DeployTensorRtModelAsyncOutput {
	s.Success = &v
	return s
}

type DeployTensorRtModelInput struct {
	AccountID            *string                                    `json:"accountID,omitempty" xml:"accountID,omitempty"`
	ConcurrencyConfig    *DeployTensorRtModelInputConcurrencyConfig `json:"concurrencyConfig,omitempty" xml:"concurrencyConfig,omitempty" type:"Struct"`
	Cpu                  *float32                                   `json:"cpu,omitempty" xml:"cpu,omitempty"`
	Description          *string                                    `json:"description,omitempty" xml:"description,omitempty"`
	DiskSize             *int32                                     `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnvName              *string                                    `json:"envName,omitempty" xml:"envName,omitempty"`
	EnvironmentVariables map[string]interface{}                     `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	GpuConfig            *DeployTensorRtModelInputGpuConfig         `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty" type:"Struct"`
	HttpTrigger          *DeployTensorRtModelInputHttpTrigger       `json:"httpTrigger,omitempty" xml:"httpTrigger,omitempty" type:"Struct"`
	ImageName            *string                                    `json:"imageName,omitempty" xml:"imageName,omitempty"`
	InstanceConcurrency  *int32                                     `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	LogConfig            *DeployTensorRtModelInputLogConfig         `json:"logConfig,omitempty" xml:"logConfig,omitempty" type:"Struct"`
	MemorySize           *int32                                     `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	ModelConfig          *DeployTensorRtModelInputModelConfig       `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Name            *string                                  `json:"name,omitempty" xml:"name,omitempty"`
	NasConfig       *DeployTensorRtModelInputNasConfig       `json:"nasConfig,omitempty" xml:"nasConfig,omitempty" type:"Struct"`
	OriginalName    *string                                  `json:"originalName,omitempty" xml:"originalName,omitempty"`
	ProjectName     *string                                  `json:"projectName,omitempty" xml:"projectName,omitempty"`
	ProvisionConfig *DeployTensorRtModelInputProvisionConfig `json:"provisionConfig,omitempty" xml:"provisionConfig,omitempty" type:"Struct"`
	Region          *string                                  `json:"region,omitempty" xml:"region,omitempty"`
	ReportStatusURL *string                                  `json:"reportStatusURL,omitempty" xml:"reportStatusURL,omitempty"`
	// This parameter is required.
	Role      *string                            `json:"role,omitempty" xml:"role,omitempty"`
	Timeout   *int32                             `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TraceId   *string                            `json:"traceId,omitempty" xml:"traceId,omitempty"`
	VpcConfig *DeployTensorRtModelInputVpcConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty" type:"Struct"`
}

func (s DeployTensorRtModelInput) String() string {
	return tea.Prettify(s)
}

func (s DeployTensorRtModelInput) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInput) SetAccountID(v string) *DeployTensorRtModelInput {
	s.AccountID = &v
	return s
}

func (s *DeployTensorRtModelInput) SetConcurrencyConfig(v *DeployTensorRtModelInputConcurrencyConfig) *DeployTensorRtModelInput {
	s.ConcurrencyConfig = v
	return s
}

func (s *DeployTensorRtModelInput) SetCpu(v float32) *DeployTensorRtModelInput {
	s.Cpu = &v
	return s
}

func (s *DeployTensorRtModelInput) SetDescription(v string) *DeployTensorRtModelInput {
	s.Description = &v
	return s
}

func (s *DeployTensorRtModelInput) SetDiskSize(v int32) *DeployTensorRtModelInput {
	s.DiskSize = &v
	return s
}

func (s *DeployTensorRtModelInput) SetEnvName(v string) *DeployTensorRtModelInput {
	s.EnvName = &v
	return s
}

func (s *DeployTensorRtModelInput) SetEnvironmentVariables(v map[string]interface{}) *DeployTensorRtModelInput {
	s.EnvironmentVariables = v
	return s
}

func (s *DeployTensorRtModelInput) SetGpuConfig(v *DeployTensorRtModelInputGpuConfig) *DeployTensorRtModelInput {
	s.GpuConfig = v
	return s
}

func (s *DeployTensorRtModelInput) SetHttpTrigger(v *DeployTensorRtModelInputHttpTrigger) *DeployTensorRtModelInput {
	s.HttpTrigger = v
	return s
}

func (s *DeployTensorRtModelInput) SetImageName(v string) *DeployTensorRtModelInput {
	s.ImageName = &v
	return s
}

func (s *DeployTensorRtModelInput) SetInstanceConcurrency(v int32) *DeployTensorRtModelInput {
	s.InstanceConcurrency = &v
	return s
}

func (s *DeployTensorRtModelInput) SetLogConfig(v *DeployTensorRtModelInputLogConfig) *DeployTensorRtModelInput {
	s.LogConfig = v
	return s
}

func (s *DeployTensorRtModelInput) SetMemorySize(v int32) *DeployTensorRtModelInput {
	s.MemorySize = &v
	return s
}

func (s *DeployTensorRtModelInput) SetModelConfig(v *DeployTensorRtModelInputModelConfig) *DeployTensorRtModelInput {
	s.ModelConfig = v
	return s
}

func (s *DeployTensorRtModelInput) SetName(v string) *DeployTensorRtModelInput {
	s.Name = &v
	return s
}

func (s *DeployTensorRtModelInput) SetNasConfig(v *DeployTensorRtModelInputNasConfig) *DeployTensorRtModelInput {
	s.NasConfig = v
	return s
}

func (s *DeployTensorRtModelInput) SetOriginalName(v string) *DeployTensorRtModelInput {
	s.OriginalName = &v
	return s
}

func (s *DeployTensorRtModelInput) SetProjectName(v string) *DeployTensorRtModelInput {
	s.ProjectName = &v
	return s
}

func (s *DeployTensorRtModelInput) SetProvisionConfig(v *DeployTensorRtModelInputProvisionConfig) *DeployTensorRtModelInput {
	s.ProvisionConfig = v
	return s
}

func (s *DeployTensorRtModelInput) SetRegion(v string) *DeployTensorRtModelInput {
	s.Region = &v
	return s
}

func (s *DeployTensorRtModelInput) SetReportStatusURL(v string) *DeployTensorRtModelInput {
	s.ReportStatusURL = &v
	return s
}

func (s *DeployTensorRtModelInput) SetRole(v string) *DeployTensorRtModelInput {
	s.Role = &v
	return s
}

func (s *DeployTensorRtModelInput) SetTimeout(v int32) *DeployTensorRtModelInput {
	s.Timeout = &v
	return s
}

func (s *DeployTensorRtModelInput) SetTraceId(v string) *DeployTensorRtModelInput {
	s.TraceId = &v
	return s
}

func (s *DeployTensorRtModelInput) SetVpcConfig(v *DeployTensorRtModelInputVpcConfig) *DeployTensorRtModelInput {
	s.VpcConfig = v
	return s
}

type DeployTensorRtModelInputConcurrencyConfig struct {
	ReservedConcurrency *int32 `json:"reservedConcurrency,omitempty" xml:"reservedConcurrency,omitempty"`
}

func (s DeployTensorRtModelInputConcurrencyConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployTensorRtModelInputConcurrencyConfig) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputConcurrencyConfig) SetReservedConcurrency(v int32) *DeployTensorRtModelInputConcurrencyConfig {
	s.ReservedConcurrency = &v
	return s
}

type DeployTensorRtModelInputGpuConfig struct {
	GpuMemorySize *int32  `json:"gpuMemorySize,omitempty" xml:"gpuMemorySize,omitempty"`
	GpuType       *string `json:"gpuType,omitempty" xml:"gpuType,omitempty"`
}

func (s DeployTensorRtModelInputGpuConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployTensorRtModelInputGpuConfig) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputGpuConfig) SetGpuMemorySize(v int32) *DeployTensorRtModelInputGpuConfig {
	s.GpuMemorySize = &v
	return s
}

func (s *DeployTensorRtModelInputGpuConfig) SetGpuType(v string) *DeployTensorRtModelInputGpuConfig {
	s.GpuType = &v
	return s
}

type DeployTensorRtModelInputHttpTrigger struct {
	Qualifier     *string                                           `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	TriggerConfig *DeployTensorRtModelInputHttpTriggerTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
}

func (s DeployTensorRtModelInputHttpTrigger) String() string {
	return tea.Prettify(s)
}

func (s DeployTensorRtModelInputHttpTrigger) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputHttpTrigger) SetQualifier(v string) *DeployTensorRtModelInputHttpTrigger {
	s.Qualifier = &v
	return s
}

func (s *DeployTensorRtModelInputHttpTrigger) SetTriggerConfig(v *DeployTensorRtModelInputHttpTriggerTriggerConfig) *DeployTensorRtModelInputHttpTrigger {
	s.TriggerConfig = v
	return s
}

type DeployTensorRtModelInputHttpTriggerTriggerConfig struct {
	AuthType          *string   `json:"authType,omitempty" xml:"authType,omitempty"`
	DsableURLInternet *bool     `json:"dsableURLInternet,omitempty" xml:"dsableURLInternet,omitempty"`
	Methods           []*string `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
}

func (s DeployTensorRtModelInputHttpTriggerTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployTensorRtModelInputHttpTriggerTriggerConfig) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputHttpTriggerTriggerConfig) SetAuthType(v string) *DeployTensorRtModelInputHttpTriggerTriggerConfig {
	s.AuthType = &v
	return s
}

func (s *DeployTensorRtModelInputHttpTriggerTriggerConfig) SetDsableURLInternet(v bool) *DeployTensorRtModelInputHttpTriggerTriggerConfig {
	s.DsableURLInternet = &v
	return s
}

func (s *DeployTensorRtModelInputHttpTriggerTriggerConfig) SetMethods(v []*string) *DeployTensorRtModelInputHttpTriggerTriggerConfig {
	s.Methods = v
	return s
}

type DeployTensorRtModelInputLogConfig struct {
	EnableInstanceMetrics *bool   `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	EnableRequestMetrics  *bool   `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	LogBeginRule          *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	Logstore              *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	Project               *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s DeployTensorRtModelInputLogConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployTensorRtModelInputLogConfig) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputLogConfig) SetEnableInstanceMetrics(v bool) *DeployTensorRtModelInputLogConfig {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *DeployTensorRtModelInputLogConfig) SetEnableRequestMetrics(v bool) *DeployTensorRtModelInputLogConfig {
	s.EnableRequestMetrics = &v
	return s
}

func (s *DeployTensorRtModelInputLogConfig) SetLogBeginRule(v string) *DeployTensorRtModelInputLogConfig {
	s.LogBeginRule = &v
	return s
}

func (s *DeployTensorRtModelInputLogConfig) SetLogstore(v string) *DeployTensorRtModelInputLogConfig {
	s.Logstore = &v
	return s
}

func (s *DeployTensorRtModelInputLogConfig) SetProject(v string) *DeployTensorRtModelInputLogConfig {
	s.Project = &v
	return s
}

type DeployTensorRtModelInputModelConfig struct {
	Framework *string `json:"framework,omitempty" xml:"framework,omitempty"`
	// if can be null:
	// true
	MultiModelConfig           []*ModelConfig `json:"multiModelConfig,omitempty" xml:"multiModelConfig,omitempty" type:"Repeated"`
	Prefix                     *string        `json:"prefix,omitempty" xml:"prefix,omitempty"`
	SourceType                 *string        `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	SrcModelScopeModelID       *string        `json:"srcModelScopeModelID,omitempty" xml:"srcModelScopeModelID,omitempty"`
	SrcModelScopeModelRevision *string        `json:"srcModelScopeModelRevision,omitempty" xml:"srcModelScopeModelRevision,omitempty"`
	SrcModelScopeToken         *string        `json:"srcModelScopeToken,omitempty" xml:"srcModelScopeToken,omitempty"`
	SrcOssBucket               *string        `json:"srcOssBucket,omitempty" xml:"srcOssBucket,omitempty"`
	SrcOssPath                 *string        `json:"srcOssPath,omitempty" xml:"srcOssPath,omitempty"`
	SrcOssRegion               *string        `json:"srcOssRegion,omitempty" xml:"srcOssRegion,omitempty"`
}

func (s DeployTensorRtModelInputModelConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployTensorRtModelInputModelConfig) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputModelConfig) SetFramework(v string) *DeployTensorRtModelInputModelConfig {
	s.Framework = &v
	return s
}

func (s *DeployTensorRtModelInputModelConfig) SetMultiModelConfig(v []*ModelConfig) *DeployTensorRtModelInputModelConfig {
	s.MultiModelConfig = v
	return s
}

func (s *DeployTensorRtModelInputModelConfig) SetPrefix(v string) *DeployTensorRtModelInputModelConfig {
	s.Prefix = &v
	return s
}

func (s *DeployTensorRtModelInputModelConfig) SetSourceType(v string) *DeployTensorRtModelInputModelConfig {
	s.SourceType = &v
	return s
}

func (s *DeployTensorRtModelInputModelConfig) SetSrcModelScopeModelID(v string) *DeployTensorRtModelInputModelConfig {
	s.SrcModelScopeModelID = &v
	return s
}

func (s *DeployTensorRtModelInputModelConfig) SetSrcModelScopeModelRevision(v string) *DeployTensorRtModelInputModelConfig {
	s.SrcModelScopeModelRevision = &v
	return s
}

func (s *DeployTensorRtModelInputModelConfig) SetSrcModelScopeToken(v string) *DeployTensorRtModelInputModelConfig {
	s.SrcModelScopeToken = &v
	return s
}

func (s *DeployTensorRtModelInputModelConfig) SetSrcOssBucket(v string) *DeployTensorRtModelInputModelConfig {
	s.SrcOssBucket = &v
	return s
}

func (s *DeployTensorRtModelInputModelConfig) SetSrcOssPath(v string) *DeployTensorRtModelInputModelConfig {
	s.SrcOssPath = &v
	return s
}

func (s *DeployTensorRtModelInputModelConfig) SetSrcOssRegion(v string) *DeployTensorRtModelInputModelConfig {
	s.SrcOssRegion = &v
	return s
}

type DeployTensorRtModelInputNasConfig struct {
	GroupId     *int32                                          `json:"groupId,omitempty" xml:"groupId,omitempty"`
	MountPoints []*DeployTensorRtModelInputNasConfigMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	UserId      *int32                                          `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeployTensorRtModelInputNasConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployTensorRtModelInputNasConfig) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputNasConfig) SetGroupId(v int32) *DeployTensorRtModelInputNasConfig {
	s.GroupId = &v
	return s
}

func (s *DeployTensorRtModelInputNasConfig) SetMountPoints(v []*DeployTensorRtModelInputNasConfigMountPoints) *DeployTensorRtModelInputNasConfig {
	s.MountPoints = v
	return s
}

func (s *DeployTensorRtModelInputNasConfig) SetUserId(v int32) *DeployTensorRtModelInputNasConfig {
	s.UserId = &v
	return s
}

type DeployTensorRtModelInputNasConfigMountPoints struct {
	EnableTLS  *bool   `json:"enableTLS,omitempty" xml:"enableTLS,omitempty"`
	MountDir   *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	ServerAddr *string `json:"serverAddr,omitempty" xml:"serverAddr,omitempty"`
}

func (s DeployTensorRtModelInputNasConfigMountPoints) String() string {
	return tea.Prettify(s)
}

func (s DeployTensorRtModelInputNasConfigMountPoints) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputNasConfigMountPoints) SetEnableTLS(v bool) *DeployTensorRtModelInputNasConfigMountPoints {
	s.EnableTLS = &v
	return s
}

func (s *DeployTensorRtModelInputNasConfigMountPoints) SetMountDir(v string) *DeployTensorRtModelInputNasConfigMountPoints {
	s.MountDir = &v
	return s
}

func (s *DeployTensorRtModelInputNasConfigMountPoints) SetServerAddr(v string) *DeployTensorRtModelInputNasConfigMountPoints {
	s.ServerAddr = &v
	return s
}

type DeployTensorRtModelInputProvisionConfig struct {
	AlwaysAllocateGPU *bool                                                      `json:"alwaysAllocateGPU,omitempty" xml:"alwaysAllocateGPU,omitempty"`
	ScheduledActions  []*DeployTensorRtModelInputProvisionConfigScheduledActions `json:"scheduledActions,omitempty" xml:"scheduledActions,omitempty" type:"Repeated"`
	Target            *int32                                                     `json:"target,omitempty" xml:"target,omitempty"`
}

func (s DeployTensorRtModelInputProvisionConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployTensorRtModelInputProvisionConfig) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputProvisionConfig) SetAlwaysAllocateGPU(v bool) *DeployTensorRtModelInputProvisionConfig {
	s.AlwaysAllocateGPU = &v
	return s
}

func (s *DeployTensorRtModelInputProvisionConfig) SetScheduledActions(v []*DeployTensorRtModelInputProvisionConfigScheduledActions) *DeployTensorRtModelInputProvisionConfig {
	s.ScheduledActions = v
	return s
}

func (s *DeployTensorRtModelInputProvisionConfig) SetTarget(v int32) *DeployTensorRtModelInputProvisionConfig {
	s.Target = &v
	return s
}

type DeployTensorRtModelInputProvisionConfigScheduledActions struct {
	EndTime            *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	ScheduleExpression *string `json:"scheduleExpression,omitempty" xml:"scheduleExpression,omitempty"`
	StartTime          *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Target             *int32  `json:"target,omitempty" xml:"target,omitempty"`
	TimeZone           *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s DeployTensorRtModelInputProvisionConfigScheduledActions) String() string {
	return tea.Prettify(s)
}

func (s DeployTensorRtModelInputProvisionConfigScheduledActions) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputProvisionConfigScheduledActions) SetEndTime(v string) *DeployTensorRtModelInputProvisionConfigScheduledActions {
	s.EndTime = &v
	return s
}

func (s *DeployTensorRtModelInputProvisionConfigScheduledActions) SetName(v string) *DeployTensorRtModelInputProvisionConfigScheduledActions {
	s.Name = &v
	return s
}

func (s *DeployTensorRtModelInputProvisionConfigScheduledActions) SetScheduleExpression(v string) *DeployTensorRtModelInputProvisionConfigScheduledActions {
	s.ScheduleExpression = &v
	return s
}

func (s *DeployTensorRtModelInputProvisionConfigScheduledActions) SetStartTime(v string) *DeployTensorRtModelInputProvisionConfigScheduledActions {
	s.StartTime = &v
	return s
}

func (s *DeployTensorRtModelInputProvisionConfigScheduledActions) SetTarget(v int32) *DeployTensorRtModelInputProvisionConfigScheduledActions {
	s.Target = &v
	return s
}

func (s *DeployTensorRtModelInputProvisionConfigScheduledActions) SetTimeZone(v string) *DeployTensorRtModelInputProvisionConfigScheduledActions {
	s.TimeZone = &v
	return s
}

type DeployTensorRtModelInputVpcConfig struct {
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	VpcId           *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s DeployTensorRtModelInputVpcConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployTensorRtModelInputVpcConfig) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputVpcConfig) SetSecurityGroupId(v string) *DeployTensorRtModelInputVpcConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *DeployTensorRtModelInputVpcConfig) SetVSwitchIds(v []*string) *DeployTensorRtModelInputVpcConfig {
	s.VSwitchIds = v
	return s
}

func (s *DeployTensorRtModelInputVpcConfig) SetVpcId(v string) *DeployTensorRtModelInputVpcConfig {
	s.VpcId = &v
	return s
}

type DeployTensorRtModelOutput struct {
	Data    *DeployTensorRtModelOutputData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                        `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                        `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployTensorRtModelOutput) String() string {
	return tea.Prettify(s)
}

func (s DeployTensorRtModelOutput) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelOutput) SetData(v *DeployTensorRtModelOutputData) *DeployTensorRtModelOutput {
	s.Data = v
	return s
}

func (s *DeployTensorRtModelOutput) SetErrCode(v string) *DeployTensorRtModelOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployTensorRtModelOutput) SetErrMsg(v string) *DeployTensorRtModelOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployTensorRtModelOutput) SetRequestId(v string) *DeployTensorRtModelOutput {
	s.RequestId = &v
	return s
}

func (s *DeployTensorRtModelOutput) SetSuccess(v bool) *DeployTensorRtModelOutput {
	s.Success = &v
	return s
}

type DeployTensorRtModelOutputData struct {
	DeploymentTaskID *string `json:"deploymentTaskID,omitempty" xml:"deploymentTaskID,omitempty"`
	ErrorMessage     *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Finished         *bool   `json:"finished,omitempty" xml:"finished,omitempty"`
	ServiceName      *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	TraceID          *string `json:"traceID,omitempty" xml:"traceID,omitempty"`
	UrlInternet      *string `json:"urlInternet,omitempty" xml:"urlInternet,omitempty"`
	UrlIntranet      *string `json:"urlIntranet,omitempty" xml:"urlIntranet,omitempty"`
}

func (s DeployTensorRtModelOutputData) String() string {
	return tea.Prettify(s)
}

func (s DeployTensorRtModelOutputData) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelOutputData) SetDeploymentTaskID(v string) *DeployTensorRtModelOutputData {
	s.DeploymentTaskID = &v
	return s
}

func (s *DeployTensorRtModelOutputData) SetErrorMessage(v string) *DeployTensorRtModelOutputData {
	s.ErrorMessage = &v
	return s
}

func (s *DeployTensorRtModelOutputData) SetFinished(v bool) *DeployTensorRtModelOutputData {
	s.Finished = &v
	return s
}

func (s *DeployTensorRtModelOutputData) SetServiceName(v string) *DeployTensorRtModelOutputData {
	s.ServiceName = &v
	return s
}

func (s *DeployTensorRtModelOutputData) SetTraceID(v string) *DeployTensorRtModelOutputData {
	s.TraceID = &v
	return s
}

func (s *DeployTensorRtModelOutputData) SetUrlInternet(v string) *DeployTensorRtModelOutputData {
	s.UrlInternet = &v
	return s
}

func (s *DeployTensorRtModelOutputData) SetUrlIntranet(v string) *DeployTensorRtModelOutputData {
	s.UrlIntranet = &v
	return s
}

type DeployVllmModelAsyncOutput struct {
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployVllmModelAsyncOutput) String() string {
	return tea.Prettify(s)
}

func (s DeployVllmModelAsyncOutput) GoString() string {
	return s.String()
}

func (s *DeployVllmModelAsyncOutput) SetData(v string) *DeployVllmModelAsyncOutput {
	s.Data = &v
	return s
}

func (s *DeployVllmModelAsyncOutput) SetErrCode(v string) *DeployVllmModelAsyncOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployVllmModelAsyncOutput) SetErrMsg(v string) *DeployVllmModelAsyncOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployVllmModelAsyncOutput) SetRequestId(v string) *DeployVllmModelAsyncOutput {
	s.RequestId = &v
	return s
}

func (s *DeployVllmModelAsyncOutput) SetSuccess(v bool) *DeployVllmModelAsyncOutput {
	s.Success = &v
	return s
}

type DeployVllmModelInput struct {
	AccountID            *string                                `json:"accountID,omitempty" xml:"accountID,omitempty"`
	ConcurrencyConfig    *DeployVllmModelInputConcurrencyConfig `json:"concurrencyConfig,omitempty" xml:"concurrencyConfig,omitempty" type:"Struct"`
	Cpu                  *float32                               `json:"cpu,omitempty" xml:"cpu,omitempty"`
	Description          *string                                `json:"description,omitempty" xml:"description,omitempty"`
	DiskSize             *int32                                 `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnvName              *string                                `json:"envName,omitempty" xml:"envName,omitempty"`
	EnvironmentVariables map[string]interface{}                 `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	GpuConfig            *DeployVllmModelInputGpuConfig         `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty" type:"Struct"`
	HttpTrigger          *DeployVllmModelInputHttpTrigger       `json:"httpTrigger,omitempty" xml:"httpTrigger,omitempty" type:"Struct"`
	ImageName            *string                                `json:"imageName,omitempty" xml:"imageName,omitempty"`
	InstanceConcurrency  *int32                                 `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	LogConfig            *DeployVllmModelInputLogConfig         `json:"logConfig,omitempty" xml:"logConfig,omitempty" type:"Struct"`
	MemorySize           *int32                                 `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	ModelConfig          *DeployVllmModelInputModelConfig       `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Name            *string                              `json:"name,omitempty" xml:"name,omitempty"`
	NasConfig       *DeployVllmModelInputNasConfig       `json:"nasConfig,omitempty" xml:"nasConfig,omitempty" type:"Struct"`
	OriginalName    *string                              `json:"originalName,omitempty" xml:"originalName,omitempty"`
	ProjectName     *string                              `json:"projectName,omitempty" xml:"projectName,omitempty"`
	ProvisionConfig *DeployVllmModelInputProvisionConfig `json:"provisionConfig,omitempty" xml:"provisionConfig,omitempty" type:"Struct"`
	Region          *string                              `json:"region,omitempty" xml:"region,omitempty"`
	ReportStatusURL *string                              `json:"reportStatusURL,omitempty" xml:"reportStatusURL,omitempty"`
	// This parameter is required.
	Role      *string                        `json:"role,omitempty" xml:"role,omitempty"`
	Timeout   *int32                         `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TraceId   *string                        `json:"traceId,omitempty" xml:"traceId,omitempty"`
	VpcConfig *DeployVllmModelInputVpcConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty" type:"Struct"`
}

func (s DeployVllmModelInput) String() string {
	return tea.Prettify(s)
}

func (s DeployVllmModelInput) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInput) SetAccountID(v string) *DeployVllmModelInput {
	s.AccountID = &v
	return s
}

func (s *DeployVllmModelInput) SetConcurrencyConfig(v *DeployVllmModelInputConcurrencyConfig) *DeployVllmModelInput {
	s.ConcurrencyConfig = v
	return s
}

func (s *DeployVllmModelInput) SetCpu(v float32) *DeployVllmModelInput {
	s.Cpu = &v
	return s
}

func (s *DeployVllmModelInput) SetDescription(v string) *DeployVllmModelInput {
	s.Description = &v
	return s
}

func (s *DeployVllmModelInput) SetDiskSize(v int32) *DeployVllmModelInput {
	s.DiskSize = &v
	return s
}

func (s *DeployVllmModelInput) SetEnvName(v string) *DeployVllmModelInput {
	s.EnvName = &v
	return s
}

func (s *DeployVllmModelInput) SetEnvironmentVariables(v map[string]interface{}) *DeployVllmModelInput {
	s.EnvironmentVariables = v
	return s
}

func (s *DeployVllmModelInput) SetGpuConfig(v *DeployVllmModelInputGpuConfig) *DeployVllmModelInput {
	s.GpuConfig = v
	return s
}

func (s *DeployVllmModelInput) SetHttpTrigger(v *DeployVllmModelInputHttpTrigger) *DeployVllmModelInput {
	s.HttpTrigger = v
	return s
}

func (s *DeployVllmModelInput) SetImageName(v string) *DeployVllmModelInput {
	s.ImageName = &v
	return s
}

func (s *DeployVllmModelInput) SetInstanceConcurrency(v int32) *DeployVllmModelInput {
	s.InstanceConcurrency = &v
	return s
}

func (s *DeployVllmModelInput) SetLogConfig(v *DeployVllmModelInputLogConfig) *DeployVllmModelInput {
	s.LogConfig = v
	return s
}

func (s *DeployVllmModelInput) SetMemorySize(v int32) *DeployVllmModelInput {
	s.MemorySize = &v
	return s
}

func (s *DeployVllmModelInput) SetModelConfig(v *DeployVllmModelInputModelConfig) *DeployVllmModelInput {
	s.ModelConfig = v
	return s
}

func (s *DeployVllmModelInput) SetName(v string) *DeployVllmModelInput {
	s.Name = &v
	return s
}

func (s *DeployVllmModelInput) SetNasConfig(v *DeployVllmModelInputNasConfig) *DeployVllmModelInput {
	s.NasConfig = v
	return s
}

func (s *DeployVllmModelInput) SetOriginalName(v string) *DeployVllmModelInput {
	s.OriginalName = &v
	return s
}

func (s *DeployVllmModelInput) SetProjectName(v string) *DeployVllmModelInput {
	s.ProjectName = &v
	return s
}

func (s *DeployVllmModelInput) SetProvisionConfig(v *DeployVllmModelInputProvisionConfig) *DeployVllmModelInput {
	s.ProvisionConfig = v
	return s
}

func (s *DeployVllmModelInput) SetRegion(v string) *DeployVllmModelInput {
	s.Region = &v
	return s
}

func (s *DeployVllmModelInput) SetReportStatusURL(v string) *DeployVllmModelInput {
	s.ReportStatusURL = &v
	return s
}

func (s *DeployVllmModelInput) SetRole(v string) *DeployVllmModelInput {
	s.Role = &v
	return s
}

func (s *DeployVllmModelInput) SetTimeout(v int32) *DeployVllmModelInput {
	s.Timeout = &v
	return s
}

func (s *DeployVllmModelInput) SetTraceId(v string) *DeployVllmModelInput {
	s.TraceId = &v
	return s
}

func (s *DeployVllmModelInput) SetVpcConfig(v *DeployVllmModelInputVpcConfig) *DeployVllmModelInput {
	s.VpcConfig = v
	return s
}

type DeployVllmModelInputConcurrencyConfig struct {
	ReservedConcurrency *int32 `json:"reservedConcurrency,omitempty" xml:"reservedConcurrency,omitempty"`
}

func (s DeployVllmModelInputConcurrencyConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployVllmModelInputConcurrencyConfig) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputConcurrencyConfig) SetReservedConcurrency(v int32) *DeployVllmModelInputConcurrencyConfig {
	s.ReservedConcurrency = &v
	return s
}

type DeployVllmModelInputGpuConfig struct {
	GpuMemorySize *int32  `json:"gpuMemorySize,omitempty" xml:"gpuMemorySize,omitempty"`
	GpuType       *string `json:"gpuType,omitempty" xml:"gpuType,omitempty"`
}

func (s DeployVllmModelInputGpuConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployVllmModelInputGpuConfig) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputGpuConfig) SetGpuMemorySize(v int32) *DeployVllmModelInputGpuConfig {
	s.GpuMemorySize = &v
	return s
}

func (s *DeployVllmModelInputGpuConfig) SetGpuType(v string) *DeployVllmModelInputGpuConfig {
	s.GpuType = &v
	return s
}

type DeployVllmModelInputHttpTrigger struct {
	Qualifier     *string                                       `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	TriggerConfig *DeployVllmModelInputHttpTriggerTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
}

func (s DeployVllmModelInputHttpTrigger) String() string {
	return tea.Prettify(s)
}

func (s DeployVllmModelInputHttpTrigger) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputHttpTrigger) SetQualifier(v string) *DeployVllmModelInputHttpTrigger {
	s.Qualifier = &v
	return s
}

func (s *DeployVllmModelInputHttpTrigger) SetTriggerConfig(v *DeployVllmModelInputHttpTriggerTriggerConfig) *DeployVllmModelInputHttpTrigger {
	s.TriggerConfig = v
	return s
}

type DeployVllmModelInputHttpTriggerTriggerConfig struct {
	AuthType          *string   `json:"authType,omitempty" xml:"authType,omitempty"`
	DsableURLInternet *bool     `json:"dsableURLInternet,omitempty" xml:"dsableURLInternet,omitempty"`
	Methods           []*string `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
}

func (s DeployVllmModelInputHttpTriggerTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployVllmModelInputHttpTriggerTriggerConfig) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputHttpTriggerTriggerConfig) SetAuthType(v string) *DeployVllmModelInputHttpTriggerTriggerConfig {
	s.AuthType = &v
	return s
}

func (s *DeployVllmModelInputHttpTriggerTriggerConfig) SetDsableURLInternet(v bool) *DeployVllmModelInputHttpTriggerTriggerConfig {
	s.DsableURLInternet = &v
	return s
}

func (s *DeployVllmModelInputHttpTriggerTriggerConfig) SetMethods(v []*string) *DeployVllmModelInputHttpTriggerTriggerConfig {
	s.Methods = v
	return s
}

type DeployVllmModelInputLogConfig struct {
	EnableInstanceMetrics *bool   `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	EnableRequestMetrics  *bool   `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	LogBeginRule          *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	Logstore              *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	Project               *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s DeployVllmModelInputLogConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployVllmModelInputLogConfig) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputLogConfig) SetEnableInstanceMetrics(v bool) *DeployVllmModelInputLogConfig {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *DeployVllmModelInputLogConfig) SetEnableRequestMetrics(v bool) *DeployVllmModelInputLogConfig {
	s.EnableRequestMetrics = &v
	return s
}

func (s *DeployVllmModelInputLogConfig) SetLogBeginRule(v string) *DeployVllmModelInputLogConfig {
	s.LogBeginRule = &v
	return s
}

func (s *DeployVllmModelInputLogConfig) SetLogstore(v string) *DeployVllmModelInputLogConfig {
	s.Logstore = &v
	return s
}

func (s *DeployVllmModelInputLogConfig) SetProject(v string) *DeployVllmModelInputLogConfig {
	s.Project = &v
	return s
}

type DeployVllmModelInputModelConfig struct {
	FmkVllmConfig *DeployVllmModelInputModelConfigFmkVllmConfig `json:"fmkVllmConfig,omitempty" xml:"fmkVllmConfig,omitempty" type:"Struct"`
	Framework     *string                                       `json:"framework,omitempty" xml:"framework,omitempty"`
	// if can be null:
	// true
	MultiModelConfig           []*ModelConfig `json:"multiModelConfig,omitempty" xml:"multiModelConfig,omitempty" type:"Repeated"`
	Prefix                     *string        `json:"prefix,omitempty" xml:"prefix,omitempty"`
	SourceType                 *string        `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	SrcModelScopeModelID       *string        `json:"srcModelScopeModelID,omitempty" xml:"srcModelScopeModelID,omitempty"`
	SrcModelScopeModelRevision *string        `json:"srcModelScopeModelRevision,omitempty" xml:"srcModelScopeModelRevision,omitempty"`
	SrcModelScopeToken         *string        `json:"srcModelScopeToken,omitempty" xml:"srcModelScopeToken,omitempty"`
	SrcOssBucket               *string        `json:"srcOssBucket,omitempty" xml:"srcOssBucket,omitempty"`
	SrcOssPath                 *string        `json:"srcOssPath,omitempty" xml:"srcOssPath,omitempty"`
	SrcOssRegion               *string        `json:"srcOssRegion,omitempty" xml:"srcOssRegion,omitempty"`
}

func (s DeployVllmModelInputModelConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployVllmModelInputModelConfig) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputModelConfig) SetFmkVllmConfig(v *DeployVllmModelInputModelConfigFmkVllmConfig) *DeployVllmModelInputModelConfig {
	s.FmkVllmConfig = v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetFramework(v string) *DeployVllmModelInputModelConfig {
	s.Framework = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetMultiModelConfig(v []*ModelConfig) *DeployVllmModelInputModelConfig {
	s.MultiModelConfig = v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetPrefix(v string) *DeployVllmModelInputModelConfig {
	s.Prefix = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetSourceType(v string) *DeployVllmModelInputModelConfig {
	s.SourceType = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetSrcModelScopeModelID(v string) *DeployVllmModelInputModelConfig {
	s.SrcModelScopeModelID = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetSrcModelScopeModelRevision(v string) *DeployVllmModelInputModelConfig {
	s.SrcModelScopeModelRevision = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetSrcModelScopeToken(v string) *DeployVllmModelInputModelConfig {
	s.SrcModelScopeToken = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetSrcOssBucket(v string) *DeployVllmModelInputModelConfig {
	s.SrcOssBucket = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetSrcOssPath(v string) *DeployVllmModelInputModelConfig {
	s.SrcOssPath = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetSrcOssRegion(v string) *DeployVllmModelInputModelConfig {
	s.SrcOssRegion = &v
	return s
}

type DeployVllmModelInputModelConfigFmkVllmConfig struct {
	ApiKey                    *string  `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	BlockSize                 *int32   `json:"blockSize,omitempty" xml:"blockSize,omitempty"`
	ChatTemplate              *string  `json:"chatTemplate,omitempty" xml:"chatTemplate,omitempty"`
	Dtype                     *string  `json:"dtype,omitempty" xml:"dtype,omitempty"`
	FullTextPostfix           *string  `json:"fullTextPostfix,omitempty" xml:"fullTextPostfix,omitempty"`
	GpuMemoryUtilization      *float32 `json:"gpuMemoryUtilization,omitempty" xml:"gpuMemoryUtilization,omitempty"`
	LoadFormat                *string  `json:"loadFormat,omitempty" xml:"loadFormat,omitempty"`
	MaxModelLen               *int32   `json:"maxModelLen,omitempty" xml:"maxModelLen,omitempty"`
	MaxParallelLoadingWorkers *int32   `json:"maxParallelLoadingWorkers,omitempty" xml:"maxParallelLoadingWorkers,omitempty"`
	Quantization              *string  `json:"quantization,omitempty" xml:"quantization,omitempty"`
	ServedModelName           *string  `json:"servedModelName,omitempty" xml:"servedModelName,omitempty"`
	SwapSpace                 *int32   `json:"swapSpace,omitempty" xml:"swapSpace,omitempty"`
}

func (s DeployVllmModelInputModelConfigFmkVllmConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployVllmModelInputModelConfigFmkVllmConfig) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetApiKey(v string) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.ApiKey = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetBlockSize(v int32) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.BlockSize = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetChatTemplate(v string) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.ChatTemplate = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetDtype(v string) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.Dtype = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetFullTextPostfix(v string) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.FullTextPostfix = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetGpuMemoryUtilization(v float32) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.GpuMemoryUtilization = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetLoadFormat(v string) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.LoadFormat = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetMaxModelLen(v int32) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.MaxModelLen = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetMaxParallelLoadingWorkers(v int32) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.MaxParallelLoadingWorkers = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetQuantization(v string) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.Quantization = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetServedModelName(v string) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.ServedModelName = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetSwapSpace(v int32) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.SwapSpace = &v
	return s
}

type DeployVllmModelInputNasConfig struct {
	GroupId     *int32                                      `json:"groupId,omitempty" xml:"groupId,omitempty"`
	MountPoints []*DeployVllmModelInputNasConfigMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	UserId      *int32                                      `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeployVllmModelInputNasConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployVllmModelInputNasConfig) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputNasConfig) SetGroupId(v int32) *DeployVllmModelInputNasConfig {
	s.GroupId = &v
	return s
}

func (s *DeployVllmModelInputNasConfig) SetMountPoints(v []*DeployVllmModelInputNasConfigMountPoints) *DeployVllmModelInputNasConfig {
	s.MountPoints = v
	return s
}

func (s *DeployVllmModelInputNasConfig) SetUserId(v int32) *DeployVllmModelInputNasConfig {
	s.UserId = &v
	return s
}

type DeployVllmModelInputNasConfigMountPoints struct {
	EnableTLS  *bool   `json:"enableTLS,omitempty" xml:"enableTLS,omitempty"`
	MountDir   *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	ServerAddr *string `json:"serverAddr,omitempty" xml:"serverAddr,omitempty"`
}

func (s DeployVllmModelInputNasConfigMountPoints) String() string {
	return tea.Prettify(s)
}

func (s DeployVllmModelInputNasConfigMountPoints) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputNasConfigMountPoints) SetEnableTLS(v bool) *DeployVllmModelInputNasConfigMountPoints {
	s.EnableTLS = &v
	return s
}

func (s *DeployVllmModelInputNasConfigMountPoints) SetMountDir(v string) *DeployVllmModelInputNasConfigMountPoints {
	s.MountDir = &v
	return s
}

func (s *DeployVllmModelInputNasConfigMountPoints) SetServerAddr(v string) *DeployVllmModelInputNasConfigMountPoints {
	s.ServerAddr = &v
	return s
}

type DeployVllmModelInputProvisionConfig struct {
	AlwaysAllocateGPU *bool                                                  `json:"alwaysAllocateGPU,omitempty" xml:"alwaysAllocateGPU,omitempty"`
	ScheduledActions  []*DeployVllmModelInputProvisionConfigScheduledActions `json:"scheduledActions,omitempty" xml:"scheduledActions,omitempty" type:"Repeated"`
	Target            *int32                                                 `json:"target,omitempty" xml:"target,omitempty"`
}

func (s DeployVllmModelInputProvisionConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployVllmModelInputProvisionConfig) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputProvisionConfig) SetAlwaysAllocateGPU(v bool) *DeployVllmModelInputProvisionConfig {
	s.AlwaysAllocateGPU = &v
	return s
}

func (s *DeployVllmModelInputProvisionConfig) SetScheduledActions(v []*DeployVllmModelInputProvisionConfigScheduledActions) *DeployVllmModelInputProvisionConfig {
	s.ScheduledActions = v
	return s
}

func (s *DeployVllmModelInputProvisionConfig) SetTarget(v int32) *DeployVllmModelInputProvisionConfig {
	s.Target = &v
	return s
}

type DeployVllmModelInputProvisionConfigScheduledActions struct {
	EndTime            *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	ScheduleExpression *string `json:"scheduleExpression,omitempty" xml:"scheduleExpression,omitempty"`
	StartTime          *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Target             *int32  `json:"target,omitempty" xml:"target,omitempty"`
	TimeZone           *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s DeployVllmModelInputProvisionConfigScheduledActions) String() string {
	return tea.Prettify(s)
}

func (s DeployVllmModelInputProvisionConfigScheduledActions) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputProvisionConfigScheduledActions) SetEndTime(v string) *DeployVllmModelInputProvisionConfigScheduledActions {
	s.EndTime = &v
	return s
}

func (s *DeployVllmModelInputProvisionConfigScheduledActions) SetName(v string) *DeployVllmModelInputProvisionConfigScheduledActions {
	s.Name = &v
	return s
}

func (s *DeployVllmModelInputProvisionConfigScheduledActions) SetScheduleExpression(v string) *DeployVllmModelInputProvisionConfigScheduledActions {
	s.ScheduleExpression = &v
	return s
}

func (s *DeployVllmModelInputProvisionConfigScheduledActions) SetStartTime(v string) *DeployVllmModelInputProvisionConfigScheduledActions {
	s.StartTime = &v
	return s
}

func (s *DeployVllmModelInputProvisionConfigScheduledActions) SetTarget(v int32) *DeployVllmModelInputProvisionConfigScheduledActions {
	s.Target = &v
	return s
}

func (s *DeployVllmModelInputProvisionConfigScheduledActions) SetTimeZone(v string) *DeployVllmModelInputProvisionConfigScheduledActions {
	s.TimeZone = &v
	return s
}

type DeployVllmModelInputVpcConfig struct {
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	VpcId           *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s DeployVllmModelInputVpcConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployVllmModelInputVpcConfig) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputVpcConfig) SetSecurityGroupId(v string) *DeployVllmModelInputVpcConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *DeployVllmModelInputVpcConfig) SetVSwitchIds(v []*string) *DeployVllmModelInputVpcConfig {
	s.VSwitchIds = v
	return s
}

func (s *DeployVllmModelInputVpcConfig) SetVpcId(v string) *DeployVllmModelInputVpcConfig {
	s.VpcId = &v
	return s
}

type DeployVllmModelOutput struct {
	Data    *DeployVllmModelOutputData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                    `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                    `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployVllmModelOutput) String() string {
	return tea.Prettify(s)
}

func (s DeployVllmModelOutput) GoString() string {
	return s.String()
}

func (s *DeployVllmModelOutput) SetData(v *DeployVllmModelOutputData) *DeployVllmModelOutput {
	s.Data = v
	return s
}

func (s *DeployVllmModelOutput) SetErrCode(v string) *DeployVllmModelOutput {
	s.ErrCode = &v
	return s
}

func (s *DeployVllmModelOutput) SetErrMsg(v string) *DeployVllmModelOutput {
	s.ErrMsg = &v
	return s
}

func (s *DeployVllmModelOutput) SetRequestId(v string) *DeployVllmModelOutput {
	s.RequestId = &v
	return s
}

func (s *DeployVllmModelOutput) SetSuccess(v bool) *DeployVllmModelOutput {
	s.Success = &v
	return s
}

type DeployVllmModelOutputData struct {
	DeploymentTaskID *string `json:"deploymentTaskID,omitempty" xml:"deploymentTaskID,omitempty"`
	ErrorMessage     *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Finished         *bool   `json:"finished,omitempty" xml:"finished,omitempty"`
	ModelName        *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ServiceName      *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	TraceID          *string `json:"traceID,omitempty" xml:"traceID,omitempty"`
	UrlInternet      *string `json:"urlInternet,omitempty" xml:"urlInternet,omitempty"`
	UrlIntranet      *string `json:"urlIntranet,omitempty" xml:"urlIntranet,omitempty"`
}

func (s DeployVllmModelOutputData) String() string {
	return tea.Prettify(s)
}

func (s DeployVllmModelOutputData) GoString() string {
	return s.String()
}

func (s *DeployVllmModelOutputData) SetDeploymentTaskID(v string) *DeployVllmModelOutputData {
	s.DeploymentTaskID = &v
	return s
}

func (s *DeployVllmModelOutputData) SetErrorMessage(v string) *DeployVllmModelOutputData {
	s.ErrorMessage = &v
	return s
}

func (s *DeployVllmModelOutputData) SetFinished(v bool) *DeployVllmModelOutputData {
	s.Finished = &v
	return s
}

func (s *DeployVllmModelOutputData) SetModelName(v string) *DeployVllmModelOutputData {
	s.ModelName = &v
	return s
}

func (s *DeployVllmModelOutputData) SetServiceName(v string) *DeployVllmModelOutputData {
	s.ServiceName = &v
	return s
}

func (s *DeployVllmModelOutputData) SetTraceID(v string) *DeployVllmModelOutputData {
	s.TraceID = &v
	return s
}

func (s *DeployVllmModelOutputData) SetUrlInternet(v string) *DeployVllmModelOutputData {
	s.UrlInternet = &v
	return s
}

func (s *DeployVllmModelOutputData) SetUrlIntranet(v string) *DeployVllmModelOutputData {
	s.UrlIntranet = &v
	return s
}

type DownloadModelOutput struct {
	Data    *DownloadModelOutputData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                  `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                  `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DownloadModelOutput) String() string {
	return tea.Prettify(s)
}

func (s DownloadModelOutput) GoString() string {
	return s.String()
}

func (s *DownloadModelOutput) SetData(v *DownloadModelOutputData) *DownloadModelOutput {
	s.Data = v
	return s
}

func (s *DownloadModelOutput) SetErrCode(v string) *DownloadModelOutput {
	s.ErrCode = &v
	return s
}

func (s *DownloadModelOutput) SetErrMsg(v string) *DownloadModelOutput {
	s.ErrMsg = &v
	return s
}

func (s *DownloadModelOutput) SetRequestId(v string) *DownloadModelOutput {
	s.RequestId = &v
	return s
}

func (s *DownloadModelOutput) SetSuccess(v bool) *DownloadModelOutput {
	s.Success = &v
	return s
}

type DownloadModelOutputData struct {
	ModelPath *string `json:"modelPath,omitempty" xml:"modelPath,omitempty"`
	TaskType  *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s DownloadModelOutputData) String() string {
	return tea.Prettify(s)
}

func (s DownloadModelOutputData) GoString() string {
	return s.String()
}

func (s *DownloadModelOutputData) SetModelPath(v string) *DownloadModelOutputData {
	s.ModelPath = &v
	return s
}

func (s *DownloadModelOutputData) SetTaskType(v string) *DownloadModelOutputData {
	s.TaskType = &v
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

type EnvironmentBaseline struct {
	ServicesInstances map[string]*ServiceInstance `json:"servicesInstances,omitempty" xml:"servicesInstances,omitempty"`
	Variables         map[string]*Variable        `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s EnvironmentBaseline) String() string {
	return tea.Prettify(s)
}

func (s EnvironmentBaseline) GoString() string {
	return s.String()
}

func (s *EnvironmentBaseline) SetServicesInstances(v map[string]*ServiceInstance) *EnvironmentBaseline {
	s.ServicesInstances = v
	return s
}

func (s *EnvironmentBaseline) SetVariables(v map[string]*Variable) *EnvironmentBaseline {
	s.Variables = v
	return s
}

type EnvironmentChanges struct {
	Services map[string]interface{} `json:"services,omitempty" xml:"services,omitempty"`
}

func (s EnvironmentChanges) String() string {
	return tea.Prettify(s)
}

func (s EnvironmentChanges) GoString() string {
	return s.String()
}

func (s *EnvironmentChanges) SetServices(v map[string]interface{}) *EnvironmentChanges {
	s.Services = v
	return s
}

type EnvironmentDeployment struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// commit by xxx.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// Deployment
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-deployment
	Name   *string                      `json:"name,omitempty" xml:"name,omitempty"`
	Spec   *EnvironmentDeploymentSpec   `json:"spec,omitempty" xml:"spec,omitempty"`
	Status *EnvironmentDeploymentStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s EnvironmentDeployment) String() string {
	return tea.Prettify(s)
}

func (s EnvironmentDeployment) GoString() string {
	return s.String()
}

func (s *EnvironmentDeployment) SetCreatedTime(v string) *EnvironmentDeployment {
	s.CreatedTime = &v
	return s
}

func (s *EnvironmentDeployment) SetDescription(v string) *EnvironmentDeployment {
	s.Description = &v
	return s
}

func (s *EnvironmentDeployment) SetKind(v string) *EnvironmentDeployment {
	s.Kind = &v
	return s
}

func (s *EnvironmentDeployment) SetLabels(v map[string]*string) *EnvironmentDeployment {
	s.Labels = v
	return s
}

func (s *EnvironmentDeployment) SetName(v string) *EnvironmentDeployment {
	s.Name = &v
	return s
}

func (s *EnvironmentDeployment) SetSpec(v *EnvironmentDeploymentSpec) *EnvironmentDeployment {
	s.Spec = v
	return s
}

func (s *EnvironmentDeployment) SetStatus(v *EnvironmentDeploymentStatus) *EnvironmentDeployment {
	s.Status = v
	return s
}

func (s *EnvironmentDeployment) SetUid(v string) *EnvironmentDeployment {
	s.Uid = &v
	return s
}

type EnvironmentDeploymentSpec struct {
	Baseline *EnvironmentSnapshot `json:"baseline,omitempty" xml:"baseline,omitempty"`
	Changes  *EnvironmentChanges  `json:"changes,omitempty" xml:"changes,omitempty"`
	// example:
	//
	// false
	SkipRemoveResources *bool                     `json:"skipRemoveResources,omitempty" xml:"skipRemoveResources,omitempty"`
	Target              *EnvironmentStagedConfigs `json:"target,omitempty" xml:"target,omitempty"`
	WebhookCodeContext  *WebhookCodeContext       `json:"webhookCodeContext,omitempty" xml:"webhookCodeContext,omitempty"`
}

func (s EnvironmentDeploymentSpec) String() string {
	return tea.Prettify(s)
}

func (s EnvironmentDeploymentSpec) GoString() string {
	return s.String()
}

func (s *EnvironmentDeploymentSpec) SetBaseline(v *EnvironmentSnapshot) *EnvironmentDeploymentSpec {
	s.Baseline = v
	return s
}

func (s *EnvironmentDeploymentSpec) SetChanges(v *EnvironmentChanges) *EnvironmentDeploymentSpec {
	s.Changes = v
	return s
}

func (s *EnvironmentDeploymentSpec) SetSkipRemoveResources(v bool) *EnvironmentDeploymentSpec {
	s.SkipRemoveResources = &v
	return s
}

func (s *EnvironmentDeploymentSpec) SetTarget(v *EnvironmentStagedConfigs) *EnvironmentDeploymentSpec {
	s.Target = v
	return s
}

func (s *EnvironmentDeploymentSpec) SetWebhookCodeContext(v *WebhookCodeContext) *EnvironmentDeploymentSpec {
	s.WebhookCodeContext = v
	return s
}

type EnvironmentDeploymentStatus struct {
	FinishedTime *string `json:"finishedTime,omitempty" xml:"finishedTime,omitempty"`
	// example:
	//
	// Running
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
	// example:
	//
	// my-pipeline
	PipelineName       *string            `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	ServiceDeployments map[string]*string `json:"serviceDeployments,omitempty" xml:"serviceDeployments,omitempty"`
}

func (s EnvironmentDeploymentStatus) String() string {
	return tea.Prettify(s)
}

func (s EnvironmentDeploymentStatus) GoString() string {
	return s.String()
}

func (s *EnvironmentDeploymentStatus) SetFinishedTime(v string) *EnvironmentDeploymentStatus {
	s.FinishedTime = &v
	return s
}

func (s *EnvironmentDeploymentStatus) SetPhase(v string) *EnvironmentDeploymentStatus {
	s.Phase = &v
	return s
}

func (s *EnvironmentDeploymentStatus) SetPipelineName(v string) *EnvironmentDeploymentStatus {
	s.PipelineName = &v
	return s
}

func (s *EnvironmentDeploymentStatus) SetServiceDeployments(v map[string]*string) *EnvironmentDeploymentStatus {
	s.ServiceDeployments = v
	return s
}

type EnvironmentSnapshot struct {
	Services map[string]*ServiceInstance `json:"services,omitempty" xml:"services,omitempty"`
}

func (s EnvironmentSnapshot) String() string {
	return tea.Prettify(s)
}

func (s EnvironmentSnapshot) GoString() string {
	return s.String()
}

func (s *EnvironmentSnapshot) SetServices(v map[string]*ServiceInstance) *EnvironmentSnapshot {
	s.Services = v
	return s
}

type EnvironmentSpec struct {
	// example:
	//
	// acs:ram::*******:role/aliyundevsdefaultrole
	RoleArn       *string                   `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	StagedConfigs *EnvironmentStagedConfigs `json:"stagedConfigs,omitempty" xml:"stagedConfigs,omitempty"`
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

func (s *EnvironmentSpec) SetRoleArn(v string) *EnvironmentSpec {
	s.RoleArn = &v
	return s
}

func (s *EnvironmentSpec) SetStagedConfigs(v *EnvironmentStagedConfigs) *EnvironmentSpec {
	s.StagedConfigs = v
	return s
}

func (s *EnvironmentSpec) SetType(v string) *EnvironmentSpec {
	s.Type = &v
	return s
}

type EnvironmentStagedConfigs struct {
	Services  map[string]*ServiceConfig `json:"services,omitempty" xml:"services,omitempty"`
	Variables map[string]*Variable      `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s EnvironmentStagedConfigs) String() string {
	return tea.Prettify(s)
}

func (s EnvironmentStagedConfigs) GoString() string {
	return s.String()
}

func (s *EnvironmentStagedConfigs) SetServices(v map[string]*ServiceConfig) *EnvironmentStagedConfigs {
	s.Services = v
	return s
}

func (s *EnvironmentStagedConfigs) SetVariables(v map[string]*Variable) *EnvironmentStagedConfigs {
	s.Variables = v
	return s
}

type EnvironmentStatus struct {
	LatestEnvironmentDeploymentName *string `json:"latestEnvironmentDeploymentName,omitempty" xml:"latestEnvironmentDeploymentName,omitempty"`
	// example:
	//
	// 1
	ObservedGeneration *int64 `json:"observedGeneration,omitempty" xml:"observedGeneration,omitempty"`
	// example:
	//
	// 2021-11-19T09:34:38Z
	ObservedTime               *string                     `json:"observedTime,omitempty" xml:"observedTime,omitempty"`
	ServicesInstances          map[string]*ServiceInstance `json:"servicesInstances,omitempty" xml:"servicesInstances,omitempty"`
	ServicesWithPendingChanges []*string                   `json:"servicesWithPendingChanges,omitempty" xml:"servicesWithPendingChanges,omitempty" type:"Repeated"`
}

func (s EnvironmentStatus) String() string {
	return tea.Prettify(s)
}

func (s EnvironmentStatus) GoString() string {
	return s.String()
}

func (s *EnvironmentStatus) SetLatestEnvironmentDeploymentName(v string) *EnvironmentStatus {
	s.LatestEnvironmentDeploymentName = &v
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

func (s *EnvironmentStatus) SetServicesInstances(v map[string]*ServiceInstance) *EnvironmentStatus {
	s.ServicesInstances = v
	return s
}

func (s *EnvironmentStatus) SetServicesWithPendingChanges(v []*string) *EnvironmentStatus {
	s.ServicesWithPendingChanges = v
	return s
}

type EventFilterConfig struct {
	Branch *BranchFilter `json:"branch,omitempty" xml:"branch,omitempty"`
}

func (s EventFilterConfig) String() string {
	return tea.Prettify(s)
}

func (s EventFilterConfig) GoString() string {
	return s.String()
}

func (s *EventFilterConfig) SetBranch(v *BranchFilter) *EventFilterConfig {
	s.Branch = v
	return s
}

type FinalizeConfig struct {
	Steps []interface{} `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
}

func (s FinalizeConfig) String() string {
	return tea.Prettify(s)
}

func (s FinalizeConfig) GoString() string {
	return s.String()
}

func (s *FinalizeConfig) SetSteps(v []interface{}) *FinalizeConfig {
	s.Steps = v
	return s
}

type GetModelStatusOutput struct {
	Data      *GetModelStatusOutputData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode   *string                   `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg    *string                   `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	RequestId *string                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetModelStatusOutput) String() string {
	return tea.Prettify(s)
}

func (s GetModelStatusOutput) GoString() string {
	return s.String()
}

func (s *GetModelStatusOutput) SetData(v *GetModelStatusOutputData) *GetModelStatusOutput {
	s.Data = v
	return s
}

func (s *GetModelStatusOutput) SetErrCode(v string) *GetModelStatusOutput {
	s.ErrCode = &v
	return s
}

func (s *GetModelStatusOutput) SetErrMsg(v string) *GetModelStatusOutput {
	s.ErrMsg = &v
	return s
}

func (s *GetModelStatusOutput) SetRequestId(v string) *GetModelStatusOutput {
	s.RequestId = &v
	return s
}

func (s *GetModelStatusOutput) SetSuccess(v bool) *GetModelStatusOutput {
	s.Success = &v
	return s
}

type GetModelStatusOutputData struct {
	CurrentBytes *int64  `json:"currentBytes,omitempty" xml:"currentBytes,omitempty"`
	ErrMessage   *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	FileSize     *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Finished     *bool   `json:"finished,omitempty" xml:"finished,omitempty"`
	FinishedTime *int64  `json:"finishedTime,omitempty" xml:"finishedTime,omitempty"`
	Speed        *int64  `json:"speed,omitempty" xml:"speed,omitempty"`
	StartTime    *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Total        *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetModelStatusOutputData) String() string {
	return tea.Prettify(s)
}

func (s GetModelStatusOutputData) GoString() string {
	return s.String()
}

func (s *GetModelStatusOutputData) SetCurrentBytes(v int64) *GetModelStatusOutputData {
	s.CurrentBytes = &v
	return s
}

func (s *GetModelStatusOutputData) SetErrMessage(v string) *GetModelStatusOutputData {
	s.ErrMessage = &v
	return s
}

func (s *GetModelStatusOutputData) SetFileSize(v int64) *GetModelStatusOutputData {
	s.FileSize = &v
	return s
}

func (s *GetModelStatusOutputData) SetFinished(v bool) *GetModelStatusOutputData {
	s.Finished = &v
	return s
}

func (s *GetModelStatusOutputData) SetFinishedTime(v int64) *GetModelStatusOutputData {
	s.FinishedTime = &v
	return s
}

func (s *GetModelStatusOutputData) SetSpeed(v int64) *GetModelStatusOutputData {
	s.Speed = &v
	return s
}

func (s *GetModelStatusOutputData) SetStartTime(v int64) *GetModelStatusOutputData {
	s.StartTime = &v
	return s
}

func (s *GetModelStatusOutputData) SetTotal(v int64) *GetModelStatusOutputData {
	s.Total = &v
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

type InitializeConfig struct {
	Steps []interface{} `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
}

func (s InitializeConfig) String() string {
	return tea.Prettify(s)
}

func (s InitializeConfig) GoString() string {
	return s.String()
}

func (s *InitializeConfig) SetSteps(v []interface{}) *InitializeConfig {
	s.Steps = v
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

type ModelAsyncTask struct {
	ErrCode      *string     `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg       *string     `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Finished     *bool       `json:"finished,omitempty" xml:"finished,omitempty"`
	FinishedTime *int64      `json:"finishedTime,omitempty" xml:"finishedTime,omitempty"`
	Result       interface{} `json:"result,omitempty" xml:"result,omitempty"`
	StartTime    *int64      `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TaskType     *string     `json:"taskType,omitempty" xml:"taskType,omitempty"`
	UpdateTime   *int64      `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ModelAsyncTask) String() string {
	return tea.Prettify(s)
}

func (s ModelAsyncTask) GoString() string {
	return s.String()
}

func (s *ModelAsyncTask) SetErrCode(v string) *ModelAsyncTask {
	s.ErrCode = &v
	return s
}

func (s *ModelAsyncTask) SetErrMsg(v string) *ModelAsyncTask {
	s.ErrMsg = &v
	return s
}

func (s *ModelAsyncTask) SetFinished(v bool) *ModelAsyncTask {
	s.Finished = &v
	return s
}

func (s *ModelAsyncTask) SetFinishedTime(v int64) *ModelAsyncTask {
	s.FinishedTime = &v
	return s
}

func (s *ModelAsyncTask) SetResult(v interface{}) *ModelAsyncTask {
	s.Result = v
	return s
}

func (s *ModelAsyncTask) SetStartTime(v int64) *ModelAsyncTask {
	s.StartTime = &v
	return s
}

func (s *ModelAsyncTask) SetTaskType(v string) *ModelAsyncTask {
	s.TaskType = &v
	return s
}

func (s *ModelAsyncTask) SetUpdateTime(v int64) *ModelAsyncTask {
	s.UpdateTime = &v
	return s
}

type ModelConfig struct {
	Bucket    *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	Framework *string `json:"framework,omitempty" xml:"framework,omitempty"`
	Model     *string `json:"model,omitempty" xml:"model,omitempty"`
	// if can be null:
	// true
	MultiModelConfig []*ModelConfig `json:"multiModelConfig,omitempty" xml:"multiModelConfig,omitempty" type:"Repeated"`
	Path             *string        `json:"path,omitempty" xml:"path,omitempty"`
	Prefix           *string        `json:"prefix,omitempty" xml:"prefix,omitempty"`
	Region           *string        `json:"region,omitempty" xml:"region,omitempty"`
	Reversion        *string        `json:"reversion,omitempty" xml:"reversion,omitempty"`
	Token            *string        `json:"token,omitempty" xml:"token,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// modelscope
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModelConfig) String() string {
	return tea.Prettify(s)
}

func (s ModelConfig) GoString() string {
	return s.String()
}

func (s *ModelConfig) SetBucket(v string) *ModelConfig {
	s.Bucket = &v
	return s
}

func (s *ModelConfig) SetFramework(v string) *ModelConfig {
	s.Framework = &v
	return s
}

func (s *ModelConfig) SetModel(v string) *ModelConfig {
	s.Model = &v
	return s
}

func (s *ModelConfig) SetMultiModelConfig(v []*ModelConfig) *ModelConfig {
	s.MultiModelConfig = v
	return s
}

func (s *ModelConfig) SetPath(v string) *ModelConfig {
	s.Path = &v
	return s
}

func (s *ModelConfig) SetPrefix(v string) *ModelConfig {
	s.Prefix = &v
	return s
}

func (s *ModelConfig) SetRegion(v string) *ModelConfig {
	s.Region = &v
	return s
}

func (s *ModelConfig) SetReversion(v string) *ModelConfig {
	s.Reversion = &v
	return s
}

func (s *ModelConfig) SetToken(v string) *ModelConfig {
	s.Token = &v
	return s
}

func (s *ModelConfig) SetType(v string) *ModelConfig {
	s.Type = &v
	return s
}

type ModelFile struct {
	IsDir    *bool   `json:"isDir,omitempty" xml:"isDir,omitempty"`
	ModeTime *int64  `json:"modeTime,omitempty" xml:"modeTime,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Path     *string `json:"path,omitempty" xml:"path,omitempty"`
	Size     *int64  `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ModelFile) String() string {
	return tea.Prettify(s)
}

func (s ModelFile) GoString() string {
	return s.String()
}

func (s *ModelFile) SetIsDir(v bool) *ModelFile {
	s.IsDir = &v
	return s
}

func (s *ModelFile) SetModeTime(v int64) *ModelFile {
	s.ModeTime = &v
	return s
}

func (s *ModelFile) SetName(v string) *ModelFile {
	s.Name = &v
	return s
}

func (s *ModelFile) SetPath(v string) *ModelFile {
	s.Path = &v
	return s
}

func (s *ModelFile) SetSize(v int64) *ModelFile {
	s.Size = &v
	return s
}

type ModelFilePreview struct {
	Content           *string `json:"content,omitempty" xml:"content,omitempty"`
	Hash              *string `json:"hash,omitempty" xml:"hash,omitempty"`
	IsCompressedImage *bool   `json:"isCompressedImage,omitempty" xml:"isCompressedImage,omitempty"`
	IsDir             *bool   `json:"isDir,omitempty" xml:"isDir,omitempty"`
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	Path              *string `json:"path,omitempty" xml:"path,omitempty"`
	Size              *int64  `json:"size,omitempty" xml:"size,omitempty"`
	Unpreviewable     *bool   `json:"unpreviewable,omitempty" xml:"unpreviewable,omitempty"`
}

func (s ModelFilePreview) String() string {
	return tea.Prettify(s)
}

func (s ModelFilePreview) GoString() string {
	return s.String()
}

func (s *ModelFilePreview) SetContent(v string) *ModelFilePreview {
	s.Content = &v
	return s
}

func (s *ModelFilePreview) SetHash(v string) *ModelFilePreview {
	s.Hash = &v
	return s
}

func (s *ModelFilePreview) SetIsCompressedImage(v bool) *ModelFilePreview {
	s.IsCompressedImage = &v
	return s
}

func (s *ModelFilePreview) SetIsDir(v bool) *ModelFilePreview {
	s.IsDir = &v
	return s
}

func (s *ModelFilePreview) SetName(v string) *ModelFilePreview {
	s.Name = &v
	return s
}

func (s *ModelFilePreview) SetPath(v string) *ModelFilePreview {
	s.Path = &v
	return s
}

func (s *ModelFilePreview) SetSize(v int64) *ModelFilePreview {
	s.Size = &v
	return s
}

func (s *ModelFilePreview) SetUnpreviewable(v bool) *ModelFilePreview {
	s.Unpreviewable = &v
	return s
}

type ModelProvider struct {
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
	// Toolset
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-toolset
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ModelProvider) String() string {
	return tea.Prettify(s)
}

func (s ModelProvider) GoString() string {
	return s.String()
}

func (s *ModelProvider) SetCreatedTime(v string) *ModelProvider {
	s.CreatedTime = &v
	return s
}

func (s *ModelProvider) SetDescription(v string) *ModelProvider {
	s.Description = &v
	return s
}

func (s *ModelProvider) SetKind(v string) *ModelProvider {
	s.Kind = &v
	return s
}

func (s *ModelProvider) SetLabels(v map[string]*string) *ModelProvider {
	s.Labels = v
	return s
}

func (s *ModelProvider) SetName(v string) *ModelProvider {
	s.Name = &v
	return s
}

func (s *ModelProvider) SetUid(v string) *ModelProvider {
	s.Uid = &v
	return s
}

type ModelProviderAuthorization struct {
	AuthConfig map[string]*string `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// example:
	//
	// apiKey
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModelProviderAuthorization) String() string {
	return tea.Prettify(s)
}

func (s ModelProviderAuthorization) GoString() string {
	return s.String()
}

func (s *ModelProviderAuthorization) SetAuthConfig(v map[string]*string) *ModelProviderAuthorization {
	s.AuthConfig = v
	return s
}

func (s *ModelProviderAuthorization) SetType(v string) *ModelProviderAuthorization {
	s.Type = &v
	return s
}

type ModelProviderSchema struct {
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// example:
	//
	// OpenAPI
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModelProviderSchema) String() string {
	return tea.Prettify(s)
}

func (s ModelProviderSchema) GoString() string {
	return s.String()
}

func (s *ModelProviderSchema) SetDetail(v string) *ModelProviderSchema {
	s.Detail = &v
	return s
}

func (s *ModelProviderSchema) SetType(v string) *ModelProviderSchema {
	s.Type = &v
	return s
}

type ModelProviderSpec struct {
	Authorization *ModelProviderAuthorization `json:"authorization,omitempty" xml:"authorization,omitempty"`
	Schema        *ModelProviderSchema        `json:"schema,omitempty" xml:"schema,omitempty"`
}

func (s ModelProviderSpec) String() string {
	return tea.Prettify(s)
}

func (s ModelProviderSpec) GoString() string {
	return s.String()
}

func (s *ModelProviderSpec) SetAuthorization(v *ModelProviderAuthorization) *ModelProviderSpec {
	s.Authorization = v
	return s
}

func (s *ModelProviderSpec) SetSchema(v *ModelProviderSchema) *ModelProviderSpec {
	s.Schema = v
	return s
}

type ModelTask struct {
	ErrMsg     *string  `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	FileSize   *float64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	FinishTime *float64 `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	Finished   *bool    `json:"finished,omitempty" xml:"finished,omitempty"`
	StartTime  *float64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status     *string  `json:"status,omitempty" xml:"status,omitempty"`
	TaskId     *string  `json:"taskId,omitempty" xml:"taskId,omitempty"`
	Total      *float32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ModelTask) String() string {
	return tea.Prettify(s)
}

func (s ModelTask) GoString() string {
	return s.String()
}

func (s *ModelTask) SetErrMsg(v string) *ModelTask {
	s.ErrMsg = &v
	return s
}

func (s *ModelTask) SetFileSize(v float64) *ModelTask {
	s.FileSize = &v
	return s
}

func (s *ModelTask) SetFinishTime(v float64) *ModelTask {
	s.FinishTime = &v
	return s
}

func (s *ModelTask) SetFinished(v bool) *ModelTask {
	s.Finished = &v
	return s
}

func (s *ModelTask) SetStartTime(v float64) *ModelTask {
	s.StartTime = &v
	return s
}

func (s *ModelTask) SetStatus(v string) *ModelTask {
	s.Status = &v
	return s
}

func (s *ModelTask) SetTaskId(v string) *ModelTask {
	s.TaskId = &v
	return s
}

func (s *ModelTask) SetTotal(v float32) *ModelTask {
	s.Total = &v
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

type OperationModelFileAction struct {
	// This parameter is required.
	Action      *string `json:"action,omitempty" xml:"action,omitempty"`
	Destination *string `json:"destination,omitempty" xml:"destination,omitempty"`
	Source      *string `json:"source,omitempty" xml:"source,omitempty"`
	Target      *string `json:"target,omitempty" xml:"target,omitempty"`
}

func (s OperationModelFileAction) String() string {
	return tea.Prettify(s)
}

func (s OperationModelFileAction) GoString() string {
	return s.String()
}

func (s *OperationModelFileAction) SetAction(v string) *OperationModelFileAction {
	s.Action = &v
	return s
}

func (s *OperationModelFileAction) SetDestination(v string) *OperationModelFileAction {
	s.Destination = &v
	return s
}

func (s *OperationModelFileAction) SetSource(v string) *OperationModelFileAction {
	s.Source = &v
	return s
}

func (s *OperationModelFileAction) SetTarget(v string) *OperationModelFileAction {
	s.Target = &v
	return s
}

type Pipeline struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
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
	// Project
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-project
	Name   *string        `json:"name,omitempty" xml:"name,omitempty"`
	Status *ProjectStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
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

func (s *Project) SetStatus(v *ProjectStatus) *Project {
	s.Status = v
	return s
}

func (s *Project) SetUid(v string) *Project {
	s.Uid = &v
	return s
}

type ProjectSpec struct {
}

func (s ProjectSpec) String() string {
	return tea.Prettify(s)
}

func (s ProjectSpec) GoString() string {
	return s.String()
}

type ProjectStatus struct {
	Services []*ServiceMeta `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s ProjectStatus) String() string {
	return tea.Prettify(s)
}

func (s ProjectStatus) GoString() string {
	return s.String()
}

func (s *ProjectStatus) SetServices(v []*ServiceMeta) *ProjectStatus {
	s.Services = v
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
	// Repository
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-repository
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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

func (s *Repository) SetSpec(v *RepositorySpec) *Repository {
	s.Spec = v
	return s
}

func (s *Repository) SetUid(v string) *Repository {
	s.Uid = &v
	return s
}

type RepositorySourceConfig struct {
	CodeVersion *CodeVersionReference `json:"codeVersion,omitempty" xml:"codeVersion,omitempty"`
	Filter      *EventFilterConfig    `json:"filter,omitempty" xml:"filter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-repository
	RepositoryName *string `json:"repositoryName,omitempty" xml:"repositoryName,omitempty"`
}

func (s RepositorySourceConfig) String() string {
	return tea.Prettify(s)
}

func (s RepositorySourceConfig) GoString() string {
	return s.String()
}

func (s *RepositorySourceConfig) SetCodeVersion(v *CodeVersionReference) *RepositorySourceConfig {
	s.CodeVersion = v
	return s
}

func (s *RepositorySourceConfig) SetFilter(v *EventFilterConfig) *RepositorySourceConfig {
	s.Filter = v
	return s
}

func (s *RepositorySourceConfig) SetRepositoryName(v string) *RepositorySourceConfig {
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

type ServiceBaseline struct {
	ServiceInstance *ServiceInstance `json:"serviceInstance,omitempty" xml:"serviceInstance,omitempty"`
}

func (s ServiceBaseline) String() string {
	return tea.Prettify(s)
}

func (s ServiceBaseline) GoString() string {
	return s.String()
}

func (s *ServiceBaseline) SetServiceInstance(v *ServiceInstance) *ServiceBaseline {
	s.ServiceInstance = v
	return s
}

type ServiceChanges struct {
	// example:
	//
	// {}: 不进行修改
	Merge map[string]interface{} `json:"merge,omitempty" xml:"merge,omitempty"`
}

func (s ServiceChanges) String() string {
	return tea.Prettify(s)
}

func (s ServiceChanges) GoString() string {
	return s.String()
}

func (s *ServiceChanges) SetMerge(v map[string]interface{}) *ServiceChanges {
	s.Merge = v
	return s
}

type ServiceCommandStep struct {
	// example:
	//
	// ./
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// s invoke
	Run *string `json:"run,omitempty" xml:"run,omitempty"`
}

func (s ServiceCommandStep) String() string {
	return tea.Prettify(s)
}

func (s ServiceCommandStep) GoString() string {
	return s.String()
}

func (s *ServiceCommandStep) SetPath(v string) *ServiceCommandStep {
	s.Path = &v
	return s
}

func (s *ServiceCommandStep) SetRun(v string) *ServiceCommandStep {
	s.Run = &v
	return s
}

type ServiceComponentStep struct {
	// example:
	//
	// fc3 invoke
	Component *string `json:"component,omitempty" xml:"component,omitempty"`
}

func (s ServiceComponentStep) String() string {
	return tea.Prettify(s)
}

func (s ServiceComponentStep) GoString() string {
	return s.String()
}

func (s *ServiceComponentStep) SetComponent(v string) *ServiceComponentStep {
	s.Component = &v
	return s
}

type ServiceConfig struct {
	Artifact *ArtifactMeta `json:"artifact,omitempty" xml:"artifact,omitempty"`
	Build    *BuildConfig  `json:"build,omitempty" xml:"build,omitempty"`
	// example:
	//
	// fc3@1.0.0
	Component *string                `json:"component,omitempty" xml:"component,omitempty"`
	Props     map[string]interface{} `json:"props,omitempty" xml:"props,omitempty"`
	Source    *SourceConfig          `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// Function
	Type      *string              `json:"type,omitempty" xml:"type,omitempty"`
	Variables map[string]*Variable `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s ServiceConfig) String() string {
	return tea.Prettify(s)
}

func (s ServiceConfig) GoString() string {
	return s.String()
}

func (s *ServiceConfig) SetArtifact(v *ArtifactMeta) *ServiceConfig {
	s.Artifact = v
	return s
}

func (s *ServiceConfig) SetBuild(v *BuildConfig) *ServiceConfig {
	s.Build = v
	return s
}

func (s *ServiceConfig) SetComponent(v string) *ServiceConfig {
	s.Component = &v
	return s
}

func (s *ServiceConfig) SetProps(v map[string]interface{}) *ServiceConfig {
	s.Props = v
	return s
}

func (s *ServiceConfig) SetSource(v *SourceConfig) *ServiceConfig {
	s.Source = v
	return s
}

func (s *ServiceConfig) SetType(v string) *ServiceConfig {
	s.Type = &v
	return s
}

func (s *ServiceConfig) SetVariables(v map[string]*Variable) *ServiceConfig {
	s.Variables = v
	return s
}

type ServiceDeployment struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// commit by xxx.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// my-environment-deployment
	EnvironmentDeploymentName *string `json:"environmentDeploymentName,omitempty" xml:"environmentDeploymentName,omitempty"`
	// example:
	//
	// Deployment
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-deployment
	Name   *string                  `json:"name,omitempty" xml:"name,omitempty"`
	Status *ServiceDeploymentStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ServiceDeployment) String() string {
	return tea.Prettify(s)
}

func (s ServiceDeployment) GoString() string {
	return s.String()
}

func (s *ServiceDeployment) SetCreatedTime(v string) *ServiceDeployment {
	s.CreatedTime = &v
	return s
}

func (s *ServiceDeployment) SetDescription(v string) *ServiceDeployment {
	s.Description = &v
	return s
}

func (s *ServiceDeployment) SetEnvironmentDeploymentName(v string) *ServiceDeployment {
	s.EnvironmentDeploymentName = &v
	return s
}

func (s *ServiceDeployment) SetKind(v string) *ServiceDeployment {
	s.Kind = &v
	return s
}

func (s *ServiceDeployment) SetLabels(v map[string]*string) *ServiceDeployment {
	s.Labels = v
	return s
}

func (s *ServiceDeployment) SetName(v string) *ServiceDeployment {
	s.Name = &v
	return s
}

func (s *ServiceDeployment) SetStatus(v *ServiceDeploymentStatus) *ServiceDeployment {
	s.Status = v
	return s
}

func (s *ServiceDeployment) SetUid(v string) *ServiceDeployment {
	s.Uid = &v
	return s
}

type ServiceDeploymentSpec struct {
	Baseline *ServiceBaseline `json:"baseline,omitempty" xml:"baseline,omitempty"`
	Changes  *ServiceChanges  `json:"changes,omitempty" xml:"changes,omitempty"`
	// example:
	//
	// false
	SkipRemoveResources *bool            `json:"skipRemoveResources,omitempty" xml:"skipRemoveResources,omitempty"`
	Target              *ServiceBaseline `json:"target,omitempty" xml:"target,omitempty"`
}

func (s ServiceDeploymentSpec) String() string {
	return tea.Prettify(s)
}

func (s ServiceDeploymentSpec) GoString() string {
	return s.String()
}

func (s *ServiceDeploymentSpec) SetBaseline(v *ServiceBaseline) *ServiceDeploymentSpec {
	s.Baseline = v
	return s
}

func (s *ServiceDeploymentSpec) SetChanges(v *ServiceChanges) *ServiceDeploymentSpec {
	s.Changes = v
	return s
}

func (s *ServiceDeploymentSpec) SetSkipRemoveResources(v bool) *ServiceDeploymentSpec {
	s.SkipRemoveResources = &v
	return s
}

func (s *ServiceDeploymentSpec) SetTarget(v *ServiceBaseline) *ServiceDeploymentSpec {
	s.Target = v
	return s
}

type ServiceDeploymentStatus struct {
	FinishedTime *string `json:"finishedTime,omitempty" xml:"finishedTime,omitempty"`
	// example:
	//
	// Running
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
	// example:
	//
	// my-pipeline
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	StartTime    *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// my-task
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s ServiceDeploymentStatus) String() string {
	return tea.Prettify(s)
}

func (s ServiceDeploymentStatus) GoString() string {
	return s.String()
}

func (s *ServiceDeploymentStatus) SetFinishedTime(v string) *ServiceDeploymentStatus {
	s.FinishedTime = &v
	return s
}

func (s *ServiceDeploymentStatus) SetPhase(v string) *ServiceDeploymentStatus {
	s.Phase = &v
	return s
}

func (s *ServiceDeploymentStatus) SetPipelineName(v string) *ServiceDeploymentStatus {
	s.PipelineName = &v
	return s
}

func (s *ServiceDeploymentStatus) SetStartTime(v string) *ServiceDeploymentStatus {
	s.StartTime = &v
	return s
}

func (s *ServiceDeploymentStatus) SetTaskName(v string) *ServiceDeploymentStatus {
	s.TaskName = &v
	return s
}

type ServiceInstance struct {
	Config           *ServiceConfig                   `json:"config,omitempty" xml:"config,omitempty"`
	LatestDeployment *ServiceInstanceLatestDeployment `json:"latestDeployment,omitempty" xml:"latestDeployment,omitempty" type:"Struct"`
	// example:
	//
	// {}
	Outputs   map[string]interface{} `json:"outputs,omitempty" xml:"outputs,omitempty"`
	Variables map[string]*Variable   `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s ServiceInstance) String() string {
	return tea.Prettify(s)
}

func (s ServiceInstance) GoString() string {
	return s.String()
}

func (s *ServiceInstance) SetConfig(v *ServiceConfig) *ServiceInstance {
	s.Config = v
	return s
}

func (s *ServiceInstance) SetLatestDeployment(v *ServiceInstanceLatestDeployment) *ServiceInstance {
	s.LatestDeployment = v
	return s
}

func (s *ServiceInstance) SetOutputs(v map[string]interface{}) *ServiceInstance {
	s.Outputs = v
	return s
}

func (s *ServiceInstance) SetVariables(v map[string]*Variable) *ServiceInstance {
	s.Variables = v
	return s
}

type ServiceInstanceLatestDeployment struct {
	FinishedTime *string `json:"finishedTime,omitempty" xml:"finishedTime,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Running
	Phase     *string `json:"phase,omitempty" xml:"phase,omitempty"`
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ServiceInstanceLatestDeployment) String() string {
	return tea.Prettify(s)
}

func (s ServiceInstanceLatestDeployment) GoString() string {
	return s.String()
}

func (s *ServiceInstanceLatestDeployment) SetFinishedTime(v string) *ServiceInstanceLatestDeployment {
	s.FinishedTime = &v
	return s
}

func (s *ServiceInstanceLatestDeployment) SetName(v string) *ServiceInstanceLatestDeployment {
	s.Name = &v
	return s
}

func (s *ServiceInstanceLatestDeployment) SetPhase(v string) *ServiceInstanceLatestDeployment {
	s.Phase = &v
	return s
}

func (s *ServiceInstanceLatestDeployment) SetStartTime(v string) *ServiceInstanceLatestDeployment {
	s.StartTime = &v
	return s
}

type ServiceMeta struct {
	// example:
	//
	// my-service
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// fc3
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ServiceMeta) String() string {
	return tea.Prettify(s)
}

func (s ServiceMeta) GoString() string {
	return s.String()
}

func (s *ServiceMeta) SetName(v string) *ServiceMeta {
	s.Name = &v
	return s
}

func (s *ServiceMeta) SetType(v string) *ServiceMeta {
	s.Type = &v
	return s
}

type ServicePluginStep struct {
	// example:
	//
	// {"key":"value"}
	Args map[string]interface{} `json:"args,omitempty" xml:"args,omitempty"`
	// example:
	//
	// dingding-robot
	Plugin *string `json:"plugin,omitempty" xml:"plugin,omitempty"`
}

func (s ServicePluginStep) String() string {
	return tea.Prettify(s)
}

func (s ServicePluginStep) GoString() string {
	return s.String()
}

func (s *ServicePluginStep) SetArgs(v map[string]interface{}) *ServicePluginStep {
	s.Args = v
	return s
}

func (s *ServicePluginStep) SetPlugin(v string) *ServicePluginStep {
	s.Plugin = &v
	return s
}

type SourceConfig struct {
	Repository *RepositorySourceConfig `json:"repository,omitempty" xml:"repository,omitempty"`
	Template   *TemplateSourceConfig   `json:"template,omitempty" xml:"template,omitempty"`
}

func (s SourceConfig) String() string {
	return tea.Prettify(s)
}

func (s SourceConfig) GoString() string {
	return s.String()
}

func (s *SourceConfig) SetRepository(v *RepositorySourceConfig) *SourceConfig {
	s.Repository = v
	return s
}

func (s *SourceConfig) SetTemplate(v *TemplateSourceConfig) *SourceConfig {
	s.Template = v
	return s
}

type Task struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
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

type Template struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// It is a template
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// Template
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// example:
	//
	// demo-template
	Name   *string         `json:"name,omitempty" xml:"name,omitempty"`
	Spec   *TemplateSpec   `json:"spec,omitempty" xml:"spec,omitempty"`
	Status *TemplateStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
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

func (s *Template) SetDescription(v string) *Template {
	s.Description = &v
	return s
}

func (s *Template) SetKind(v string) *Template {
	s.Kind = &v
	return s
}

func (s *Template) SetLabels(v map[string]*string) *Template {
	s.Labels = v
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

type TemplateConfig struct {
	// example:
	//
	// {"region":"cn-hangzhou"}
	Parameters map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// example:
	//
	// {"svc1":"svc2"}
	ServiceNameChanges map[string]*string `json:"serviceNameChanges,omitempty" xml:"serviceNameChanges,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// start-springboot
	TemplateName   *string                   `json:"templateName,omitempty" xml:"templateName,omitempty"`
	VariableValues *TemplateVariableValueMap `json:"variableValues,omitempty" xml:"variableValues,omitempty"`
}

func (s TemplateConfig) String() string {
	return tea.Prettify(s)
}

func (s TemplateConfig) GoString() string {
	return s.String()
}

func (s *TemplateConfig) SetParameters(v map[string]*string) *TemplateConfig {
	s.Parameters = v
	return s
}

func (s *TemplateConfig) SetServiceNameChanges(v map[string]*string) *TemplateConfig {
	s.ServiceNameChanges = v
	return s
}

func (s *TemplateConfig) SetTemplateName(v string) *TemplateConfig {
	s.TemplateName = &v
	return s
}

func (s *TemplateConfig) SetVariableValues(v *TemplateVariableValueMap) *TemplateConfig {
	s.VariableValues = v
	return s
}

type TemplateParameterSchema struct {
	// example:
	//
	// defaultValue
	Default interface{} `json:"default,omitempty" xml:"default,omitempty"`
	// example:
	//
	// Parameters for testing
	Description *string   `json:"description,omitempty" xml:"description,omitempty"`
	Enum        []*string `json:"enum,omitempty" xml:"enum,omitempty" type:"Repeated"`
	// example:
	//
	// "^[a-zA-Z._-]+$"
	Pattern *string `json:"pattern,omitempty" xml:"pattern,omitempty"`
	// example:
	//
	// true
	Required      *bool                                 `json:"required,omitempty" xml:"required,omitempty"`
	RoleExtension *TemplateParameterSchemaRoleExtension `json:"roleExtension,omitempty" xml:"roleExtension,omitempty" type:"Struct"`
	// example:
	//
	// false
	Sensitive *bool `json:"sensitive,omitempty" xml:"sensitive,omitempty"`
	// example:
	//
	// demo
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TemplateParameterSchema) String() string {
	return tea.Prettify(s)
}

func (s TemplateParameterSchema) GoString() string {
	return s.String()
}

func (s *TemplateParameterSchema) SetDefault(v interface{}) *TemplateParameterSchema {
	s.Default = v
	return s
}

func (s *TemplateParameterSchema) SetDescription(v string) *TemplateParameterSchema {
	s.Description = &v
	return s
}

func (s *TemplateParameterSchema) SetEnum(v []*string) *TemplateParameterSchema {
	s.Enum = v
	return s
}

func (s *TemplateParameterSchema) SetPattern(v string) *TemplateParameterSchema {
	s.Pattern = &v
	return s
}

func (s *TemplateParameterSchema) SetRequired(v bool) *TemplateParameterSchema {
	s.Required = &v
	return s
}

func (s *TemplateParameterSchema) SetRoleExtension(v *TemplateParameterSchemaRoleExtension) *TemplateParameterSchema {
	s.RoleExtension = v
	return s
}

func (s *TemplateParameterSchema) SetSensitive(v bool) *TemplateParameterSchema {
	s.Sensitive = &v
	return s
}

func (s *TemplateParameterSchema) SetTitle(v string) *TemplateParameterSchema {
	s.Title = &v
	return s
}

func (s *TemplateParameterSchema) SetType(v string) *TemplateParameterSchema {
	s.Type = &v
	return s
}

type TemplateParameterSchemaRoleExtension struct {
	Authorities []*string `json:"authorities,omitempty" xml:"authorities,omitempty" type:"Repeated"`
	Name        *string   `json:"name,omitempty" xml:"name,omitempty"`
	Service     *string   `json:"service,omitempty" xml:"service,omitempty"`
}

func (s TemplateParameterSchemaRoleExtension) String() string {
	return tea.Prettify(s)
}

func (s TemplateParameterSchemaRoleExtension) GoString() string {
	return s.String()
}

func (s *TemplateParameterSchemaRoleExtension) SetAuthorities(v []*string) *TemplateParameterSchemaRoleExtension {
	s.Authorities = v
	return s
}

func (s *TemplateParameterSchemaRoleExtension) SetName(v string) *TemplateParameterSchemaRoleExtension {
	s.Name = &v
	return s
}

func (s *TemplateParameterSchemaRoleExtension) SetService(v string) *TemplateParameterSchemaRoleExtension {
	s.Service = &v
	return s
}

type TemplateRevision struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// It is a template revision
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// TemplateRevision
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// example:
	//
	// demo-template
	Name   *string                 `json:"name,omitempty" xml:"name,omitempty"`
	Spec   *TemplateSpec           `json:"spec,omitempty" xml:"spec,omitempty"`
	Status *TemplateRevisionStatus `json:"status,omitempty" xml:"status,omitempty" type:"Struct"`
	// example:
	//
	// demo-template
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// example:
	//
	// 1455541096***548
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

func (s *TemplateRevision) SetDescription(v string) *TemplateRevision {
	s.Description = &v
	return s
}

func (s *TemplateRevision) SetKind(v string) *TemplateRevision {
	s.Kind = &v
	return s
}

func (s *TemplateRevision) SetLabels(v map[string]*string) *TemplateRevision {
	s.Labels = v
	return s
}

func (s *TemplateRevision) SetName(v string) *TemplateRevision {
	s.Name = &v
	return s
}

func (s *TemplateRevision) SetSpec(v *TemplateSpec) *TemplateRevision {
	s.Spec = v
	return s
}

func (s *TemplateRevision) SetStatus(v *TemplateRevisionStatus) *TemplateRevision {
	s.Status = v
	return s
}

func (s *TemplateRevision) SetTemplateName(v string) *TemplateRevision {
	s.TemplateName = &v
	return s
}

func (s *TemplateRevision) SetUid(v string) *TemplateRevision {
	s.Uid = &v
	return s
}

type TemplateRevisionStatus struct {
	// example:
	//
	// https://registry.serverless-devs.com/details.html?name=template-test&package_type=v3
	PackageUrl *string `json:"packageUrl,omitempty" xml:"packageUrl,omitempty"`
	// example:
	//
	// Published
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
	// example:
	//
	// p-default
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// example:
	//
	// https://cap.console.aliyun.com/template-detail?template=adasdasdaewe-adadqwe
	TemplateUrl *string `json:"templateUrl,omitempty" xml:"templateUrl,omitempty"`
}

func (s TemplateRevisionStatus) String() string {
	return tea.Prettify(s)
}

func (s TemplateRevisionStatus) GoString() string {
	return s.String()
}

func (s *TemplateRevisionStatus) SetPackageUrl(v string) *TemplateRevisionStatus {
	s.PackageUrl = &v
	return s
}

func (s *TemplateRevisionStatus) SetPhase(v string) *TemplateRevisionStatus {
	s.Phase = &v
	return s
}

func (s *TemplateRevisionStatus) SetPipelineName(v string) *TemplateRevisionStatus {
	s.PipelineName = &v
	return s
}

func (s *TemplateRevisionStatus) SetTemplateUrl(v string) *TemplateRevisionStatus {
	s.TemplateUrl = &v
	return s
}

type TemplateServiceConfig struct {
	Artifact *ArtifactMeta `json:"artifact,omitempty" xml:"artifact,omitempty"`
	Build    *BuildConfig  `json:"build,omitempty" xml:"build,omitempty"`
	// example:
	//
	// fc3@1.0.0
	Component *string                `json:"component,omitempty" xml:"component,omitempty"`
	Props     map[string]interface{} `json:"props,omitempty" xml:"props,omitempty"`
	Source    *SourceConfig          `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// Function
	Type      *string                             `json:"type,omitempty" xml:"type,omitempty"`
	Variables map[string]*TemplateParameterSchema `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s TemplateServiceConfig) String() string {
	return tea.Prettify(s)
}

func (s TemplateServiceConfig) GoString() string {
	return s.String()
}

func (s *TemplateServiceConfig) SetArtifact(v *ArtifactMeta) *TemplateServiceConfig {
	s.Artifact = v
	return s
}

func (s *TemplateServiceConfig) SetBuild(v *BuildConfig) *TemplateServiceConfig {
	s.Build = v
	return s
}

func (s *TemplateServiceConfig) SetComponent(v string) *TemplateServiceConfig {
	s.Component = &v
	return s
}

func (s *TemplateServiceConfig) SetProps(v map[string]interface{}) *TemplateServiceConfig {
	s.Props = v
	return s
}

func (s *TemplateServiceConfig) SetSource(v *SourceConfig) *TemplateServiceConfig {
	s.Source = v
	return s
}

func (s *TemplateServiceConfig) SetType(v string) *TemplateServiceConfig {
	s.Type = &v
	return s
}

func (s *TemplateServiceConfig) SetVariables(v map[string]*TemplateParameterSchema) *TemplateServiceConfig {
	s.Variables = v
	return s
}

type TemplateSourceConfig struct {
	// example:
	//
	// https://api.devsapp.cn/v3/packages/start-modelscope-v3/zipball/0.1.6
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	// example:
	//
	// start-springboot-cap
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s TemplateSourceConfig) String() string {
	return tea.Prettify(s)
}

func (s TemplateSourceConfig) GoString() string {
	return s.String()
}

func (s *TemplateSourceConfig) SetDownloadUrl(v string) *TemplateSourceConfig {
	s.DownloadUrl = &v
	return s
}

func (s *TemplateSourceConfig) SetName(v string) *TemplateSourceConfig {
	s.Name = &v
	return s
}

type TemplateSpec struct {
	// This parameter is required.
	//
	// example:
	//
	// CAP
	Author *string `json:"author,omitempty" xml:"author,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AI
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// Apache-1.0
	License *string `json:"license,omitempty" xml:"license,omitempty"`
	// example:
	//
	// demo-package
	PackageName *string `json:"packageName,omitempty" xml:"packageName,omitempty"`
	// This parameter is required.
	Readme        *string                             `json:"readme,omitempty" xml:"readme,omitempty"`
	RegistryToken *string                             `json:"registryToken,omitempty" xml:"registryToken,omitempty"`
	Services      map[string]*TemplateServiceConfig   `json:"services,omitempty" xml:"services,omitempty"`
	Source        *TemplateSpecSource                 `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
	Variables     map[string]*TemplateParameterSchema `json:"variables,omitempty" xml:"variables,omitempty"`
	Version       *string                             `json:"version,omitempty" xml:"version,omitempty"`
}

func (s TemplateSpec) String() string {
	return tea.Prettify(s)
}

func (s TemplateSpec) GoString() string {
	return s.String()
}

func (s *TemplateSpec) SetAuthor(v string) *TemplateSpec {
	s.Author = &v
	return s
}

func (s *TemplateSpec) SetCategory(v string) *TemplateSpec {
	s.Category = &v
	return s
}

func (s *TemplateSpec) SetLicense(v string) *TemplateSpec {
	s.License = &v
	return s
}

func (s *TemplateSpec) SetPackageName(v string) *TemplateSpec {
	s.PackageName = &v
	return s
}

func (s *TemplateSpec) SetReadme(v string) *TemplateSpec {
	s.Readme = &v
	return s
}

func (s *TemplateSpec) SetRegistryToken(v string) *TemplateSpec {
	s.RegistryToken = &v
	return s
}

func (s *TemplateSpec) SetServices(v map[string]*TemplateServiceConfig) *TemplateSpec {
	s.Services = v
	return s
}

func (s *TemplateSpec) SetSource(v *TemplateSpecSource) *TemplateSpec {
	s.Source = v
	return s
}

func (s *TemplateSpec) SetVariables(v map[string]*TemplateParameterSchema) *TemplateSpec {
	s.Variables = v
	return s
}

func (s *TemplateSpec) SetVersion(v string) *TemplateSpec {
	s.Version = &v
	return s
}

type TemplateSpecSource struct {
	Repository *RepositorySourceConfig `json:"repository,omitempty" xml:"repository,omitempty"`
}

func (s TemplateSpecSource) String() string {
	return tea.Prettify(s)
}

func (s TemplateSpecSource) GoString() string {
	return s.String()
}

func (s *TemplateSpecSource) SetRepository(v *RepositorySourceConfig) *TemplateSpecSource {
	s.Repository = v
	return s
}

type TemplateStatus struct {
	LatestDeployment *TemplateStatusLatestDeployment `json:"latestDeployment,omitempty" xml:"latestDeployment,omitempty" type:"Struct"`
	// example:
	//
	// 1.0.0
	LatestVersion *string `json:"latestVersion,omitempty" xml:"latestVersion,omitempty"`
	// example:
	//
	// https://registry.serverless-devs.com/details.html?name=template-test&package_type=v3
	PackageUrl *string `json:"packageUrl,omitempty" xml:"packageUrl,omitempty"`
	// example:
	//
	// Published
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
	// example:
	//
	// https://cap.console.aliyun.com/template-detail?template=adasdasdaewe-adadqwe
	TemplateUrl *string `json:"templateUrl,omitempty" xml:"templateUrl,omitempty"`
}

func (s TemplateStatus) String() string {
	return tea.Prettify(s)
}

func (s TemplateStatus) GoString() string {
	return s.String()
}

func (s *TemplateStatus) SetLatestDeployment(v *TemplateStatusLatestDeployment) *TemplateStatus {
	s.LatestDeployment = v
	return s
}

func (s *TemplateStatus) SetLatestVersion(v string) *TemplateStatus {
	s.LatestVersion = &v
	return s
}

func (s *TemplateStatus) SetPackageUrl(v string) *TemplateStatus {
	s.PackageUrl = &v
	return s
}

func (s *TemplateStatus) SetPhase(v string) *TemplateStatus {
	s.Phase = &v
	return s
}

func (s *TemplateStatus) SetTemplateUrl(v string) *TemplateStatus {
	s.TemplateUrl = &v
	return s
}

type TemplateStatusLatestDeployment struct {
	FinishedTime *string `json:"finishedTime,omitempty" xml:"finishedTime,omitempty"`
	// example:
	//
	// BuildFinished
	Phase        *string `json:"phase,omitempty" xml:"phase,omitempty"`
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	StartTime    *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s TemplateStatusLatestDeployment) String() string {
	return tea.Prettify(s)
}

func (s TemplateStatusLatestDeployment) GoString() string {
	return s.String()
}

func (s *TemplateStatusLatestDeployment) SetFinishedTime(v string) *TemplateStatusLatestDeployment {
	s.FinishedTime = &v
	return s
}

func (s *TemplateStatusLatestDeployment) SetPhase(v string) *TemplateStatusLatestDeployment {
	s.Phase = &v
	return s
}

func (s *TemplateStatusLatestDeployment) SetPipelineName(v string) *TemplateStatusLatestDeployment {
	s.PipelineName = &v
	return s
}

func (s *TemplateStatusLatestDeployment) SetStartTime(v string) *TemplateStatusLatestDeployment {
	s.StartTime = &v
	return s
}

type TemplateVariableValueMap struct {
	Services map[string]map[string]interface{} `json:"services,omitempty" xml:"services,omitempty"`
	// example:
	//
	// {"TEST_KEY":"new_value"}
	Shared map[string]interface{} `json:"shared,omitempty" xml:"shared,omitempty"`
}

func (s TemplateVariableValueMap) String() string {
	return tea.Prettify(s)
}

func (s TemplateVariableValueMap) GoString() string {
	return s.String()
}

func (s *TemplateVariableValueMap) SetServices(v map[string]map[string]interface{}) *TemplateVariableValueMap {
	s.Services = v
	return s
}

func (s *TemplateVariableValueMap) SetShared(v map[string]interface{}) *TemplateVariableValueMap {
	s.Shared = v
	return s
}

type Tool struct {
	Method   *string `json:"method,omitempty" xml:"method,omitempty"`
	Path     *string `json:"path,omitempty" xml:"path,omitempty"`
	ToolId   *string `json:"toolId,omitempty" xml:"toolId,omitempty"`
	ToolName *string `json:"toolName,omitempty" xml:"toolName,omitempty"`
}

func (s Tool) String() string {
	return tea.Prettify(s)
}

func (s Tool) GoString() string {
	return s.String()
}

func (s *Tool) SetMethod(v string) *Tool {
	s.Method = &v
	return s
}

func (s *Tool) SetPath(v string) *Tool {
	s.Path = &v
	return s
}

func (s *Tool) SetToolId(v string) *Tool {
	s.ToolId = &v
	return s
}

func (s *Tool) SetToolName(v string) *Tool {
	s.ToolName = &v
	return s
}

type Toolset struct {
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
	// Toolset
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-toolset
	Name   *string        `json:"name,omitempty" xml:"name,omitempty"`
	Status *ToolsetStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s Toolset) String() string {
	return tea.Prettify(s)
}

func (s Toolset) GoString() string {
	return s.String()
}

func (s *Toolset) SetCreatedTime(v string) *Toolset {
	s.CreatedTime = &v
	return s
}

func (s *Toolset) SetDescription(v string) *Toolset {
	s.Description = &v
	return s
}

func (s *Toolset) SetKind(v string) *Toolset {
	s.Kind = &v
	return s
}

func (s *Toolset) SetLabels(v map[string]*string) *Toolset {
	s.Labels = v
	return s
}

func (s *Toolset) SetName(v string) *Toolset {
	s.Name = &v
	return s
}

func (s *Toolset) SetStatus(v *ToolsetStatus) *Toolset {
	s.Status = v
	return s
}

func (s *Toolset) SetUid(v string) *Toolset {
	s.Uid = &v
	return s
}

type ToolsetAuthorization struct {
	AuthConfig map[string]*string `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// example:
	//
	// apiKey
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ToolsetAuthorization) String() string {
	return tea.Prettify(s)
}

func (s ToolsetAuthorization) GoString() string {
	return s.String()
}

func (s *ToolsetAuthorization) SetAuthConfig(v map[string]*string) *ToolsetAuthorization {
	s.AuthConfig = v
	return s
}

func (s *ToolsetAuthorization) SetType(v string) *ToolsetAuthorization {
	s.Type = &v
	return s
}

type ToolsetSchema struct {
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// example:
	//
	// OpenAPI
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ToolsetSchema) String() string {
	return tea.Prettify(s)
}

func (s ToolsetSchema) GoString() string {
	return s.String()
}

func (s *ToolsetSchema) SetDetail(v string) *ToolsetSchema {
	s.Detail = &v
	return s
}

func (s *ToolsetSchema) SetType(v string) *ToolsetSchema {
	s.Type = &v
	return s
}

type ToolsetSpec struct {
	Authorization *ToolsetAuthorization `json:"authorization,omitempty" xml:"authorization,omitempty"`
	Schema        *ToolsetSchema        `json:"schema,omitempty" xml:"schema,omitempty"`
}

func (s ToolsetSpec) String() string {
	return tea.Prettify(s)
}

func (s ToolsetSpec) GoString() string {
	return s.String()
}

func (s *ToolsetSpec) SetAuthorization(v *ToolsetAuthorization) *ToolsetSpec {
	s.Authorization = v
	return s
}

func (s *ToolsetSpec) SetSchema(v *ToolsetSchema) *ToolsetSpec {
	s.Schema = v
	return s
}

type ToolsetStatus struct {
	Tools map[string]*Tool `json:"tools,omitempty" xml:"tools,omitempty"`
}

func (s ToolsetStatus) String() string {
	return tea.Prettify(s)
}

func (s ToolsetStatus) GoString() string {
	return s.String()
}

func (s *ToolsetStatus) SetTools(v map[string]*Tool) *ToolsetStatus {
	s.Tools = v
	return s
}

type Variable struct {
	// example:
	//
	// false
	Encrypted *bool `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	// example:
	//
	// false
	Sensitive *bool `json:"sensitive,omitempty" xml:"sensitive,omitempty"`
	// example:
	//
	// object_value
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s Variable) String() string {
	return tea.Prettify(s)
}

func (s Variable) GoString() string {
	return s.String()
}

func (s *Variable) SetEncrypted(v bool) *Variable {
	s.Encrypted = &v
	return s
}

func (s *Variable) SetSensitive(v bool) *Variable {
	s.Sensitive = &v
	return s
}

func (s *Variable) SetValue(v interface{}) *Variable {
	s.Value = v
	return s
}

type WebhookCodeContext struct {
	// example:
	//
	// master
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// b1dd9ba168dfef1cb3a1dd608b6054c771a93959
	CommitID *string `json:"commitID,omitempty" xml:"commitID,omitempty"`
	// example:
	//
	// my PR decscription
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// PUSH
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// example:
	//
	// commit message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// merged
	PrType *string `json:"prType,omitempty" xml:"prType,omitempty"`
	// example:
	//
	// https://codeup.aliyun.com/my-namespace/my-repo.git
	RepoUrl *string `json:"repoUrl,omitempty" xml:"repoUrl,omitempty"`
	// example:
	//
	// master
	SourceBranch *string `json:"sourceBranch,omitempty" xml:"sourceBranch,omitempty"`
	// example:
	//
	// release-0.0.1
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// example:
	//
	// master
	TargetBranch *string `json:"targetBranch,omitempty" xml:"targetBranch,omitempty"`
	// example:
	//
	// # FIX
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s WebhookCodeContext) String() string {
	return tea.Prettify(s)
}

func (s WebhookCodeContext) GoString() string {
	return s.String()
}

func (s *WebhookCodeContext) SetBranch(v string) *WebhookCodeContext {
	s.Branch = &v
	return s
}

func (s *WebhookCodeContext) SetCommitID(v string) *WebhookCodeContext {
	s.CommitID = &v
	return s
}

func (s *WebhookCodeContext) SetDescription(v string) *WebhookCodeContext {
	s.Description = &v
	return s
}

func (s *WebhookCodeContext) SetEventType(v string) *WebhookCodeContext {
	s.EventType = &v
	return s
}

func (s *WebhookCodeContext) SetMessage(v string) *WebhookCodeContext {
	s.Message = &v
	return s
}

func (s *WebhookCodeContext) SetPrType(v string) *WebhookCodeContext {
	s.PrType = &v
	return s
}

func (s *WebhookCodeContext) SetRepoUrl(v string) *WebhookCodeContext {
	s.RepoUrl = &v
	return s
}

func (s *WebhookCodeContext) SetSourceBranch(v string) *WebhookCodeContext {
	s.SourceBranch = &v
	return s
}

func (s *WebhookCodeContext) SetTag(v string) *WebhookCodeContext {
	s.Tag = &v
	return s
}

func (s *WebhookCodeContext) SetTargetBranch(v string) *WebhookCodeContext {
	s.TargetBranch = &v
	return s
}

func (s *WebhookCodeContext) SetTitle(v string) *WebhookCodeContext {
	s.Title = &v
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

type DeleteProjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type DeployEnvironmentRequest struct {
	Body *DeployEnvironmentOptions `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployEnvironmentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *DeployEnvironmentRequest) SetBody(v *DeployEnvironmentOptions) *DeployEnvironmentRequest {
	s.Body = v
	return s
}

type DeployEnvironmentResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnvironmentDeployment `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *DeployEnvironmentResponse) SetHeaders(v map[string]*string) *DeployEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *DeployEnvironmentResponse) SetStatusCode(v int32) *DeployEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployEnvironmentResponse) SetBody(v *EnvironmentDeployment) *DeployEnvironmentResponse {
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

type GetEnvironmentDeploymentResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnvironmentDeployment `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEnvironmentDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentDeploymentResponse) GoString() string {
	return s.String()
}

func (s *GetEnvironmentDeploymentResponse) SetHeaders(v map[string]*string) *GetEnvironmentDeploymentResponse {
	s.Headers = v
	return s
}

func (s *GetEnvironmentDeploymentResponse) SetStatusCode(v int32) *GetEnvironmentDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnvironmentDeploymentResponse) SetBody(v *EnvironmentDeployment) *GetEnvironmentDeploymentResponse {
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

type UpdateProjectRequest struct {
	Body *Project `json:"body,omitempty" xml:"body,omitempty"`
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CancelPipelineResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CancelPipelineResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CancelTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CancelTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 创建环境
//
// @param request - CreateEnvironmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnvironmentResponse
func (client *Client) CreateEnvironmentWithOptions(projectName *string, request *CreateEnvironmentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateEnvironmentResponse, _err error) {
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
		Pathname:    tea.String("/2023-07-14/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/environments"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateEnvironmentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateEnvironmentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建环境
//
// @param request - CreateEnvironmentRequest
//
// @return CreateEnvironmentResponse
func (client *Client) CreateEnvironment(projectName *string, request *CreateEnvironmentRequest) (_result *CreateEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEnvironmentResponse{}
	_body, _err := client.CreateEnvironmentWithOptions(projectName, request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreatePipelineResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreatePipelineResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateProjectResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateProjectResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 删除环境
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnvironmentResponse
func (client *Client) DeleteEnvironmentWithOptions(projectName *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteEnvironmentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEnvironment"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/environments/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteEnvironmentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteEnvironmentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除环境
//
// @return DeleteEnvironmentResponse
func (client *Client) DeleteEnvironment(projectName *string, name *string) (_result *DeleteEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEnvironmentResponse{}
	_body, _err := client.DeleteEnvironmentWithOptions(projectName, name, headers, runtime)
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
		BodyType:    tea.String("none"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteProjectResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteProjectResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 手动触发环境部署
//
// @param request - DeployEnvironmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployEnvironmentResponse
func (client *Client) DeployEnvironmentWithOptions(projectName *string, name *string, request *DeployEnvironmentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeployEnvironmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployEnvironment"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/environments/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/deploy"),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeployEnvironmentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeployEnvironmentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 手动触发环境部署
//
// @param request - DeployEnvironmentRequest
//
// @return DeployEnvironmentResponse
func (client *Client) DeployEnvironment(projectName *string, name *string, request *DeployEnvironmentRequest) (_result *DeployEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeployEnvironmentResponse{}
	_body, _err := client.DeployEnvironmentWithOptions(projectName, name, request, headers, runtime)
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
func (client *Client) GetEnvironmentWithOptions(projectName *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEnvironmentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetEnvironment"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/environments/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetEnvironmentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetEnvironmentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取环境信息
//
// @return GetEnvironmentResponse
func (client *Client) GetEnvironment(projectName *string, name *string) (_result *GetEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEnvironmentResponse{}
	_body, _err := client.GetEnvironmentWithOptions(projectName, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询环境部署信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEnvironmentDeploymentResponse
func (client *Client) GetEnvironmentDeploymentWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEnvironmentDeploymentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetEnvironmentDeployment"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/environmentdeployments/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetEnvironmentDeploymentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetEnvironmentDeploymentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询环境部署信息
//
// @return GetEnvironmentDeploymentResponse
func (client *Client) GetEnvironmentDeployment(name *string) (_result *GetEnvironmentDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEnvironmentDeploymentResponse{}
	_body, _err := client.GetEnvironmentDeploymentWithOptions(name, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetPipelineResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetPipelineResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetProjectResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetProjectResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetRepositoryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetRepositoryResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 查询环境列表
//
// @param tmpReq - ListEnvironmentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnvironmentsResponse
func (client *Client) ListEnvironmentsWithOptions(projectName *string, tmpReq *ListEnvironmentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEnvironmentsResponse, _err error) {
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
		Pathname:    tea.String("/2023-07-14/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/environments/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListEnvironmentsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListEnvironmentsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询环境列表
//
// @param request - ListEnvironmentsRequest
//
// @return ListEnvironmentsResponse
func (client *Client) ListEnvironments(projectName *string, request *ListEnvironmentsRequest) (_result *ListEnvironmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEnvironmentsResponse{}
	_body, _err := client.ListEnvironmentsWithOptions(projectName, request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListPipelinesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListPipelinesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListProjectsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListProjectsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTasksResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTasksResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &PutPipelineStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &PutPipelineStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &PutTaskStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &PutTaskStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ResumeTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ResumeTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RetryTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RetryTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StartPipelineResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StartPipelineResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StartTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StartTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
func (client *Client) UpdateEnvironmentWithOptions(projectName *string, name *string, request *UpdateEnvironmentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateEnvironmentResponse, _err error) {
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
		Pathname:    tea.String("/2023-07-14/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/environments/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateEnvironmentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateEnvironmentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新环境（局部更新）
//
// @param request - UpdateEnvironmentRequest
//
// @return UpdateEnvironmentResponse
func (client *Client) UpdateEnvironment(projectName *string, name *string, request *UpdateEnvironmentRequest) (_result *UpdateEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEnvironmentResponse{}
	_body, _err := client.UpdateEnvironmentWithOptions(projectName, name, request, headers, runtime)
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
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProject"),
		Version:     tea.String("2023-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-07-14/projects/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateProjectResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateProjectResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
