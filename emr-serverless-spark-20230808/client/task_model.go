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
	SetRayActiveDeadlineSeconds(v int64) *Task
	GetRayActiveDeadlineSeconds() *int64
	SetRayBackoffLimit(v int32) *Task
	GetRayBackoffLimit() *int32
	SetRayEntrypoint(v string) *Task
	GetRayEntrypoint() *string
	SetRayEntrypointMemory(v string) *Task
	GetRayEntrypointMemory() *string
	SetRayEntrypointNumCpus(v string) *Task
	GetRayEntrypointNumCpus() *string
	SetRayEntrypointNumGpus(v string) *Task
	GetRayEntrypointNumGpus() *string
	SetRayEntrypointResources(v string) *Task
	GetRayEntrypointResources() *string
	SetRayExtraParam(v string) *Task
	GetRayExtraParam() *string
	SetRayHeadSpec(v *TaskRayHeadSpec) *Task
	GetRayHeadSpec() *TaskRayHeadSpec
	SetRayMetadataJson(v string) *Task
	GetRayMetadataJson() *string
	SetRayNetworkServiceName(v string) *Task
	GetRayNetworkServiceName() *string
	SetRayRuntimeEnvJson(v string) *Task
	GetRayRuntimeEnvJson() *string
	SetRayShutdownAfterJobFinishes(v bool) *Task
	GetRayShutdownAfterJobFinishes() *bool
	SetRaySubmissionMode(v string) *Task
	GetRaySubmissionMode() *string
	SetRayTtlSecondsAfterFinished(v int64) *Task
	GetRayTtlSecondsAfterFinished() *int64
	SetRayVersion(v string) *Task
	GetRayVersion() *string
	SetRayVolumeIds(v []*string) *Task
	GetRayVolumeIds() []*string
	SetRayWorkerSpec(v []*TaskRayWorkerSpec) *Task
	GetRayWorkerSpec() []*TaskRayWorkerSpec
	SetRayWorkingDir(v string) *Task
	GetRayWorkingDir() *string
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
	// The --archives parameter.
	Archives []*string `json:"archives,omitempty" xml:"archives,omitempty" type:"Repeated"`
	// The temporary URL to access the resource file.
	ArtifactUrl *string `json:"artifactUrl,omitempty" xml:"artifactUrl,omitempty"`
	// The business ID.
	//
	// This parameter is required.
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// The business ID of the folder.
	CategoryBizId *string `json:"categoryBizId,omitempty" xml:"categoryBizId,omitempty"`
	// The content of the Spark job.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The UID of the creator.
	//
	// This parameter is required.
	Creator *int64 `json:"creator,omitempty" xml:"creator,omitempty"`
	// The information for directly uploading files to Object Storage Service (OSS).
	Credential *TaskCredential `json:"credential,omitempty" xml:"credential,omitempty" type:"Struct"`
	// The default catalog ID.
	DefaultCatalogId *string `json:"defaultCatalogId,omitempty" xml:"defaultCatalogId,omitempty"`
	// The default database.
	DefaultDatabase *string `json:"defaultDatabase,omitempty" xml:"defaultDatabase,omitempty"`
	// The default queue ID for the task.
	DefaultResourceQueueId *string `json:"defaultResourceQueueId,omitempty" xml:"defaultResourceQueueId,omitempty"`
	// The default SQL session ID.
	DefaultSqlComputeId *string `json:"defaultSqlComputeId,omitempty" xml:"defaultSqlComputeId,omitempty"`
	// The deployment ID of the streaming task.
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// The environment ID.
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The IDs of extra Spark resources.
	ExtraArtifactIds []*string `json:"extraArtifactIds,omitempty" xml:"extraArtifactIds,omitempty" type:"Repeated"`
	// The custom parameters for the spark-submit command.
	ExtraSparkSubmitParams *string `json:"extraSparkSubmitParams,omitempty" xml:"extraSparkSubmitParams,omitempty"`
	// The --files parameter.
	Files []*string `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// Indicates whether to enable fusion.
	Fusion *bool `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// The time when the task was created.
	//
	// This parameter is required.
	GmtCreated *string `json:"gmtCreated,omitempty" xml:"gmtCreated,omitempty"`
	// The time when the task was last modified.
	//
	// This parameter is required.
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Indicates whether the task has been changed since the last commit.
	HasChanged *bool `json:"hasChanged,omitempty" xml:"hasChanged,omitempty"`
	// Indicates whether the task has been committed.
	//
	// This parameter is required.
	HasCommited *bool `json:"hasCommited,omitempty" xml:"hasCommited,omitempty"`
	// Indicates whether the task is a streaming task.
	IsStreaming *bool `json:"isStreaming,omitempty" xml:"isStreaming,omitempty"`
	// The --jars parameter.
	Jars     []*string `json:"jars,omitempty" xml:"jars,omitempty" type:"Repeated"`
	KernelId *string   `json:"kernelId,omitempty" xml:"kernelId,omitempty"`
	// The ID of the resource queue that was used for the last run.
	LastRunResourceQueueId *string `json:"lastRunResourceQueueId,omitempty" xml:"lastRunResourceQueueId,omitempty"`
	// The UID of the user who last updated the task.
	//
	// This parameter is required.
	Modifier *int64 `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// The task name.
	//
	// This parameter is required.
	Name   *string            `json:"name,omitempty" xml:"name,omitempty"`
	Params map[string]*string `json:"params,omitempty" xml:"params,omitempty"`
	// The PySpark dependency pyfiles.
	PyFiles                     []*string            `json:"pyFiles,omitempty" xml:"pyFiles,omitempty" type:"Repeated"`
	RayActiveDeadlineSeconds    *int64               `json:"rayActiveDeadlineSeconds,omitempty" xml:"rayActiveDeadlineSeconds,omitempty"`
	RayBackoffLimit             *int32               `json:"rayBackoffLimit,omitempty" xml:"rayBackoffLimit,omitempty"`
	RayEntrypoint               *string              `json:"rayEntrypoint,omitempty" xml:"rayEntrypoint,omitempty"`
	RayEntrypointMemory         *string              `json:"rayEntrypointMemory,omitempty" xml:"rayEntrypointMemory,omitempty"`
	RayEntrypointNumCpus        *string              `json:"rayEntrypointNumCpus,omitempty" xml:"rayEntrypointNumCpus,omitempty"`
	RayEntrypointNumGpus        *string              `json:"rayEntrypointNumGpus,omitempty" xml:"rayEntrypointNumGpus,omitempty"`
	RayEntrypointResources      *string              `json:"rayEntrypointResources,omitempty" xml:"rayEntrypointResources,omitempty"`
	RayExtraParam               *string              `json:"rayExtraParam,omitempty" xml:"rayExtraParam,omitempty"`
	RayHeadSpec                 *TaskRayHeadSpec     `json:"rayHeadSpec,omitempty" xml:"rayHeadSpec,omitempty" type:"Struct"`
	RayMetadataJson             *string              `json:"rayMetadataJson,omitempty" xml:"rayMetadataJson,omitempty"`
	RayNetworkServiceName       *string              `json:"rayNetworkServiceName,omitempty" xml:"rayNetworkServiceName,omitempty"`
	RayRuntimeEnvJson           *string              `json:"rayRuntimeEnvJson,omitempty" xml:"rayRuntimeEnvJson,omitempty"`
	RayShutdownAfterJobFinishes *bool                `json:"rayShutdownAfterJobFinishes,omitempty" xml:"rayShutdownAfterJobFinishes,omitempty"`
	RaySubmissionMode           *string              `json:"raySubmissionMode,omitempty" xml:"raySubmissionMode,omitempty"`
	RayTtlSecondsAfterFinished  *int64               `json:"rayTtlSecondsAfterFinished,omitempty" xml:"rayTtlSecondsAfterFinished,omitempty"`
	RayVersion                  *string              `json:"rayVersion,omitempty" xml:"rayVersion,omitempty"`
	RayVolumeIds                []*string            `json:"rayVolumeIds,omitempty" xml:"rayVolumeIds,omitempty" type:"Repeated"`
	RayWorkerSpec               []*TaskRayWorkerSpec `json:"rayWorkerSpec,omitempty" xml:"rayWorkerSpec,omitempty" type:"Repeated"`
	RayWorkingDir               *string              `json:"rayWorkingDir,omitempty" xml:"rayWorkingDir,omitempty"`
	SessionClusterId            *string              `json:"sessionClusterId,omitempty" xml:"sessionClusterId,omitempty"`
	// The Spark parameters.
	//
	// example:
	//
	// 100
	SparkArgs *string `json:"sparkArgs,omitempty" xml:"sparkArgs,omitempty"`
	// The list of Spark configurations.
	SparkConf []*SparkConf `json:"sparkConf,omitempty" xml:"sparkConf,omitempty" type:"Repeated"`
	// The number of cores for the Spark driver.
	//
	// This parameter is required.
	SparkDriverCores *int32 `json:"sparkDriverCores,omitempty" xml:"sparkDriverCores,omitempty"`
	// The memory of the Spark driver.
	//
	// This parameter is required.
	SparkDriverMemory *int64 `json:"sparkDriverMemory,omitempty" xml:"sparkDriverMemory,omitempty"`
	// The entrypoint of the Spark main class.
	SparkEntrypoint *string `json:"sparkEntrypoint,omitempty" xml:"sparkEntrypoint,omitempty"`
	// The number of cores for the Spark executor.
	//
	// This parameter is required.
	SparkExecutorCores *int32 `json:"sparkExecutorCores,omitempty" xml:"sparkExecutorCores,omitempty"`
	// The memory of the Spark executor.
	//
	// This parameter is required.
	SparkExecutorMemory *int64 `json:"sparkExecutorMemory,omitempty" xml:"sparkExecutorMemory,omitempty"`
	// The log level for Spark.
	//
	// This parameter is required.
	SparkLogLevel *string `json:"sparkLogLevel,omitempty" xml:"sparkLogLevel,omitempty"`
	// The Spark log path.
	//
	// This parameter is required.
	SparkLogPath *string `json:"sparkLogPath,omitempty" xml:"sparkLogPath,omitempty"`
	// The spark-submit statement.
	SparkSubmitClause *string `json:"sparkSubmitClause,omitempty" xml:"sparkSubmitClause,omitempty"`
	// The Spark version.
	//
	// This parameter is required.
	SparkVersion *string `json:"sparkVersion,omitempty" xml:"sparkVersion,omitempty"`
	// The task tags.
	Tags map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The task timeout duration.
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// The task type.
	//
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

func (s *Task) GetRayActiveDeadlineSeconds() *int64 {
	return s.RayActiveDeadlineSeconds
}

func (s *Task) GetRayBackoffLimit() *int32 {
	return s.RayBackoffLimit
}

func (s *Task) GetRayEntrypoint() *string {
	return s.RayEntrypoint
}

func (s *Task) GetRayEntrypointMemory() *string {
	return s.RayEntrypointMemory
}

func (s *Task) GetRayEntrypointNumCpus() *string {
	return s.RayEntrypointNumCpus
}

func (s *Task) GetRayEntrypointNumGpus() *string {
	return s.RayEntrypointNumGpus
}

func (s *Task) GetRayEntrypointResources() *string {
	return s.RayEntrypointResources
}

func (s *Task) GetRayExtraParam() *string {
	return s.RayExtraParam
}

func (s *Task) GetRayHeadSpec() *TaskRayHeadSpec {
	return s.RayHeadSpec
}

func (s *Task) GetRayMetadataJson() *string {
	return s.RayMetadataJson
}

func (s *Task) GetRayNetworkServiceName() *string {
	return s.RayNetworkServiceName
}

func (s *Task) GetRayRuntimeEnvJson() *string {
	return s.RayRuntimeEnvJson
}

func (s *Task) GetRayShutdownAfterJobFinishes() *bool {
	return s.RayShutdownAfterJobFinishes
}

func (s *Task) GetRaySubmissionMode() *string {
	return s.RaySubmissionMode
}

func (s *Task) GetRayTtlSecondsAfterFinished() *int64 {
	return s.RayTtlSecondsAfterFinished
}

func (s *Task) GetRayVersion() *string {
	return s.RayVersion
}

func (s *Task) GetRayVolumeIds() []*string {
	return s.RayVolumeIds
}

func (s *Task) GetRayWorkerSpec() []*TaskRayWorkerSpec {
	return s.RayWorkerSpec
}

func (s *Task) GetRayWorkingDir() *string {
	return s.RayWorkingDir
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

func (s *Task) SetRayActiveDeadlineSeconds(v int64) *Task {
	s.RayActiveDeadlineSeconds = &v
	return s
}

func (s *Task) SetRayBackoffLimit(v int32) *Task {
	s.RayBackoffLimit = &v
	return s
}

func (s *Task) SetRayEntrypoint(v string) *Task {
	s.RayEntrypoint = &v
	return s
}

func (s *Task) SetRayEntrypointMemory(v string) *Task {
	s.RayEntrypointMemory = &v
	return s
}

func (s *Task) SetRayEntrypointNumCpus(v string) *Task {
	s.RayEntrypointNumCpus = &v
	return s
}

func (s *Task) SetRayEntrypointNumGpus(v string) *Task {
	s.RayEntrypointNumGpus = &v
	return s
}

func (s *Task) SetRayEntrypointResources(v string) *Task {
	s.RayEntrypointResources = &v
	return s
}

func (s *Task) SetRayExtraParam(v string) *Task {
	s.RayExtraParam = &v
	return s
}

func (s *Task) SetRayHeadSpec(v *TaskRayHeadSpec) *Task {
	s.RayHeadSpec = v
	return s
}

func (s *Task) SetRayMetadataJson(v string) *Task {
	s.RayMetadataJson = &v
	return s
}

func (s *Task) SetRayNetworkServiceName(v string) *Task {
	s.RayNetworkServiceName = &v
	return s
}

func (s *Task) SetRayRuntimeEnvJson(v string) *Task {
	s.RayRuntimeEnvJson = &v
	return s
}

func (s *Task) SetRayShutdownAfterJobFinishes(v bool) *Task {
	s.RayShutdownAfterJobFinishes = &v
	return s
}

func (s *Task) SetRaySubmissionMode(v string) *Task {
	s.RaySubmissionMode = &v
	return s
}

func (s *Task) SetRayTtlSecondsAfterFinished(v int64) *Task {
	s.RayTtlSecondsAfterFinished = &v
	return s
}

func (s *Task) SetRayVersion(v string) *Task {
	s.RayVersion = &v
	return s
}

func (s *Task) SetRayVolumeIds(v []*string) *Task {
	s.RayVolumeIds = v
	return s
}

func (s *Task) SetRayWorkerSpec(v []*TaskRayWorkerSpec) *Task {
	s.RayWorkerSpec = v
	return s
}

func (s *Task) SetRayWorkingDir(v string) *Task {
	s.RayWorkingDir = &v
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
	if s.RayHeadSpec != nil {
		if err := s.RayHeadSpec.Validate(); err != nil {
			return err
		}
	}
	if s.RayWorkerSpec != nil {
		for _, item := range s.RayWorkerSpec {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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
	// The AccessKey ID.
	AccessId *string `json:"accessId,omitempty" xml:"accessId,omitempty"`
	// The access URL.
	AccessUrl *string `json:"accessUrl,omitempty" xml:"accessUrl,omitempty"`
	// The expiration time.
	Expire *int64 `json:"expire,omitempty" xml:"expire,omitempty"`
	// The domain name.
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// The path.
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The policy.
	Policy *string `json:"policy,omitempty" xml:"policy,omitempty"`
	// The security token.
	SecurityToken *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
	// The signature.
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
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

type TaskRayHeadSpec struct {
	Cpu                *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
	EnableAutoScaling  *bool   `json:"enableAutoScaling,omitempty" xml:"enableAutoScaling,omitempty"`
	GpuSpec            *string `json:"gpuSpec,omitempty" xml:"gpuSpec,omitempty"`
	IdleTimeoutSeconds *int64  `json:"idleTimeoutSeconds,omitempty" xml:"idleTimeoutSeconds,omitempty"`
	Memory             *string `json:"memory,omitempty" xml:"memory,omitempty"`
	QueueName          *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	Replica            *int32  `json:"replica,omitempty" xml:"replica,omitempty"`
}

func (s TaskRayHeadSpec) String() string {
	return dara.Prettify(s)
}

func (s TaskRayHeadSpec) GoString() string {
	return s.String()
}

func (s *TaskRayHeadSpec) GetCpu() *string {
	return s.Cpu
}

func (s *TaskRayHeadSpec) GetEnableAutoScaling() *bool {
	return s.EnableAutoScaling
}

func (s *TaskRayHeadSpec) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *TaskRayHeadSpec) GetIdleTimeoutSeconds() *int64 {
	return s.IdleTimeoutSeconds
}

func (s *TaskRayHeadSpec) GetMemory() *string {
	return s.Memory
}

func (s *TaskRayHeadSpec) GetQueueName() *string {
	return s.QueueName
}

func (s *TaskRayHeadSpec) GetReplica() *int32 {
	return s.Replica
}

func (s *TaskRayHeadSpec) SetCpu(v string) *TaskRayHeadSpec {
	s.Cpu = &v
	return s
}

func (s *TaskRayHeadSpec) SetEnableAutoScaling(v bool) *TaskRayHeadSpec {
	s.EnableAutoScaling = &v
	return s
}

func (s *TaskRayHeadSpec) SetGpuSpec(v string) *TaskRayHeadSpec {
	s.GpuSpec = &v
	return s
}

func (s *TaskRayHeadSpec) SetIdleTimeoutSeconds(v int64) *TaskRayHeadSpec {
	s.IdleTimeoutSeconds = &v
	return s
}

func (s *TaskRayHeadSpec) SetMemory(v string) *TaskRayHeadSpec {
	s.Memory = &v
	return s
}

func (s *TaskRayHeadSpec) SetQueueName(v string) *TaskRayHeadSpec {
	s.QueueName = &v
	return s
}

func (s *TaskRayHeadSpec) SetReplica(v int32) *TaskRayHeadSpec {
	s.Replica = &v
	return s
}

func (s *TaskRayHeadSpec) Validate() error {
	return dara.Validate(s)
}

type TaskRayWorkerSpec struct {
	Cpu        *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
	GpuSpec    *string `json:"gpuSpec,omitempty" xml:"gpuSpec,omitempty"`
	GroupName  *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	MaxReplica *int32  `json:"maxReplica,omitempty" xml:"maxReplica,omitempty"`
	Memory     *string `json:"memory,omitempty" xml:"memory,omitempty"`
	MinReplica *int32  `json:"minReplica,omitempty" xml:"minReplica,omitempty"`
	QueueName  *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	Replica    *int32  `json:"replica,omitempty" xml:"replica,omitempty"`
}

func (s TaskRayWorkerSpec) String() string {
	return dara.Prettify(s)
}

func (s TaskRayWorkerSpec) GoString() string {
	return s.String()
}

func (s *TaskRayWorkerSpec) GetCpu() *string {
	return s.Cpu
}

func (s *TaskRayWorkerSpec) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *TaskRayWorkerSpec) GetGroupName() *string {
	return s.GroupName
}

func (s *TaskRayWorkerSpec) GetMaxReplica() *int32 {
	return s.MaxReplica
}

func (s *TaskRayWorkerSpec) GetMemory() *string {
	return s.Memory
}

func (s *TaskRayWorkerSpec) GetMinReplica() *int32 {
	return s.MinReplica
}

func (s *TaskRayWorkerSpec) GetQueueName() *string {
	return s.QueueName
}

func (s *TaskRayWorkerSpec) GetReplica() *int32 {
	return s.Replica
}

func (s *TaskRayWorkerSpec) SetCpu(v string) *TaskRayWorkerSpec {
	s.Cpu = &v
	return s
}

func (s *TaskRayWorkerSpec) SetGpuSpec(v string) *TaskRayWorkerSpec {
	s.GpuSpec = &v
	return s
}

func (s *TaskRayWorkerSpec) SetGroupName(v string) *TaskRayWorkerSpec {
	s.GroupName = &v
	return s
}

func (s *TaskRayWorkerSpec) SetMaxReplica(v int32) *TaskRayWorkerSpec {
	s.MaxReplica = &v
	return s
}

func (s *TaskRayWorkerSpec) SetMemory(v string) *TaskRayWorkerSpec {
	s.Memory = &v
	return s
}

func (s *TaskRayWorkerSpec) SetMinReplica(v int32) *TaskRayWorkerSpec {
	s.MinReplica = &v
	return s
}

func (s *TaskRayWorkerSpec) SetQueueName(v string) *TaskRayWorkerSpec {
	s.QueueName = &v
	return s
}

func (s *TaskRayWorkerSpec) SetReplica(v int32) *TaskRayWorkerSpec {
	s.Replica = &v
	return s
}

func (s *TaskRayWorkerSpec) Validate() error {
	return dara.Validate(s)
}
