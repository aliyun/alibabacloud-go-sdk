// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTask interface {
	dara.Model
	String() string
	GoString() string
	SetArchives(v []*string) *Task
	GetArchives() []*string
	SetArtifactUrl(v string) *Task
	GetArtifactUrl() *string
	SetBizId(v string) *Task
	GetBizId() *string
	SetCategoryBizId(v string) *Task
	GetCategoryBizId() *string
	SetContent(v string) *Task
	GetContent() *string
	SetCreator(v int64) *Task
	GetCreator() *int64
	SetCredential(v *TaskCredential) *Task
	GetCredential() *TaskCredential
	SetDefaultCatalogId(v string) *Task
	GetDefaultCatalogId() *string
	SetDefaultDatabase(v string) *Task
	GetDefaultDatabase() *string
	SetDefaultResourceQueueId(v string) *Task
	GetDefaultResourceQueueId() *string
	SetDefaultSqlComputeId(v string) *Task
	GetDefaultSqlComputeId() *string
	SetDeploymentId(v string) *Task
	GetDeploymentId() *string
	SetEnvironmentId(v string) *Task
	GetEnvironmentId() *string
	SetExtraArtifactIds(v []*string) *Task
	GetExtraArtifactIds() []*string
	SetExtraSparkSubmitParams(v string) *Task
	GetExtraSparkSubmitParams() *string
	SetFiles(v []*string) *Task
	GetFiles() []*string
	SetFusion(v bool) *Task
	GetFusion() *bool
	SetGmtCreated(v string) *Task
	GetGmtCreated() *string
	SetGmtModified(v string) *Task
	GetGmtModified() *string
	SetHasChanged(v bool) *Task
	GetHasChanged() *bool
	SetHasCommited(v bool) *Task
	GetHasCommited() *bool
	SetIsStreaming(v bool) *Task
	GetIsStreaming() *bool
	SetJars(v []*string) *Task
	GetJars() []*string
	SetKernelId(v string) *Task
	GetKernelId() *string
	SetLastRunResourceQueueId(v string) *Task
	GetLastRunResourceQueueId() *string
	SetModifier(v int64) *Task
	GetModifier() *int64
	SetName(v string) *Task
	GetName() *string
	SetParams(v map[string]*string) *Task
	GetParams() map[string]*string
	SetPyFiles(v []*string) *Task
	GetPyFiles() []*string
	SetSessionClusterId(v string) *Task
	GetSessionClusterId() *string
	SetSparkArgs(v string) *Task
	GetSparkArgs() *string
	SetSparkConf(v []*SparkConf) *Task
	GetSparkConf() []*SparkConf
	SetSparkDriverCores(v int32) *Task
	GetSparkDriverCores() *int32
	SetSparkDriverMemory(v int64) *Task
	GetSparkDriverMemory() *int64
	SetSparkEntrypoint(v string) *Task
	GetSparkEntrypoint() *string
	SetSparkExecutorCores(v int32) *Task
	GetSparkExecutorCores() *int32
	SetSparkExecutorMemory(v int64) *Task
	GetSparkExecutorMemory() *int64
	SetSparkLogLevel(v string) *Task
	GetSparkLogLevel() *string
	SetSparkLogPath(v string) *Task
	GetSparkLogPath() *string
	SetSparkSubmitClause(v string) *Task
	GetSparkSubmitClause() *string
	SetSparkVersion(v string) *Task
	GetSparkVersion() *string
	SetTags(v map[string]*string) *Task
	GetTags() map[string]*string
	SetTimeout(v int32) *Task
	GetTimeout() *int32
	SetType(v string) *Task
	GetType() *string
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
	return dara.Prettify(s)
}

func (s Task) GoString() string {
	return s.String()
}

func (s *Task) GetArchives() []*string {
	return s.Archives
}

func (s *Task) GetArtifactUrl() *string {
	return s.ArtifactUrl
}

func (s *Task) GetBizId() *string {
	return s.BizId
}

func (s *Task) GetCategoryBizId() *string {
	return s.CategoryBizId
}

func (s *Task) GetContent() *string {
	return s.Content
}

func (s *Task) GetCreator() *int64 {
	return s.Creator
}

func (s *Task) GetCredential() *TaskCredential {
	return s.Credential
}

func (s *Task) GetDefaultCatalogId() *string {
	return s.DefaultCatalogId
}

func (s *Task) GetDefaultDatabase() *string {
	return s.DefaultDatabase
}

func (s *Task) GetDefaultResourceQueueId() *string {
	return s.DefaultResourceQueueId
}

func (s *Task) GetDefaultSqlComputeId() *string {
	return s.DefaultSqlComputeId
}

func (s *Task) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *Task) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *Task) GetExtraArtifactIds() []*string {
	return s.ExtraArtifactIds
}

func (s *Task) GetExtraSparkSubmitParams() *string {
	return s.ExtraSparkSubmitParams
}

func (s *Task) GetFiles() []*string {
	return s.Files
}

func (s *Task) GetFusion() *bool {
	return s.Fusion
}

func (s *Task) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *Task) GetGmtModified() *string {
	return s.GmtModified
}

func (s *Task) GetHasChanged() *bool {
	return s.HasChanged
}

func (s *Task) GetHasCommited() *bool {
	return s.HasCommited
}

func (s *Task) GetIsStreaming() *bool {
	return s.IsStreaming
}

func (s *Task) GetJars() []*string {
	return s.Jars
}

func (s *Task) GetKernelId() *string {
	return s.KernelId
}

func (s *Task) GetLastRunResourceQueueId() *string {
	return s.LastRunResourceQueueId
}

func (s *Task) GetModifier() *int64 {
	return s.Modifier
}

func (s *Task) GetName() *string {
	return s.Name
}

func (s *Task) GetParams() map[string]*string {
	return s.Params
}

func (s *Task) GetPyFiles() []*string {
	return s.PyFiles
}

func (s *Task) GetSessionClusterId() *string {
	return s.SessionClusterId
}

func (s *Task) GetSparkArgs() *string {
	return s.SparkArgs
}

func (s *Task) GetSparkConf() []*SparkConf {
	return s.SparkConf
}

func (s *Task) GetSparkDriverCores() *int32 {
	return s.SparkDriverCores
}

func (s *Task) GetSparkDriverMemory() *int64 {
	return s.SparkDriverMemory
}

func (s *Task) GetSparkEntrypoint() *string {
	return s.SparkEntrypoint
}

func (s *Task) GetSparkExecutorCores() *int32 {
	return s.SparkExecutorCores
}

func (s *Task) GetSparkExecutorMemory() *int64 {
	return s.SparkExecutorMemory
}

func (s *Task) GetSparkLogLevel() *string {
	return s.SparkLogLevel
}

func (s *Task) GetSparkLogPath() *string {
	return s.SparkLogPath
}

func (s *Task) GetSparkSubmitClause() *string {
	return s.SparkSubmitClause
}

func (s *Task) GetSparkVersion() *string {
	return s.SparkVersion
}

func (s *Task) GetTags() map[string]*string {
	return s.Tags
}

func (s *Task) GetTimeout() *int32 {
	return s.Timeout
}

func (s *Task) GetType() *string {
	return s.Type
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

func (s *Task) Validate() error {
	if s.Credential != nil {
		if err := s.Credential.Validate(); err != nil {
			return err
		}
	}
	if s.SparkConf != nil {
		for _, item := range s.SparkConf {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s TaskCredential) GoString() string {
	return s.String()
}

func (s *TaskCredential) GetAccessId() *string {
	return s.AccessId
}

func (s *TaskCredential) GetAccessUrl() *string {
	return s.AccessUrl
}

func (s *TaskCredential) GetExpire() *int64 {
	return s.Expire
}

func (s *TaskCredential) GetHost() *string {
	return s.Host
}

func (s *TaskCredential) GetPath() *string {
	return s.Path
}

func (s *TaskCredential) GetPolicy() *string {
	return s.Policy
}

func (s *TaskCredential) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *TaskCredential) GetSignature() *string {
	return s.Signature
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

func (s *TaskCredential) Validate() error {
	return dara.Validate(s)
}
