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

type Artifact struct {
	JarArtifact    *JarArtifact    `json:"jarArtifact,omitempty" xml:"jarArtifact,omitempty"`
	Kind           *string         `json:"kind,omitempty" xml:"kind,omitempty"`
	PythonArtifact *PythonArtifact `json:"pythonArtifact,omitempty" xml:"pythonArtifact,omitempty"`
	SqlArtifact    *SqlArtifact    `json:"sqlArtifact,omitempty" xml:"sqlArtifact,omitempty"`
}

func (s Artifact) String() string {
	return tea.Prettify(s)
}

func (s Artifact) GoString() string {
	return s.String()
}

func (s *Artifact) SetJarArtifact(v *JarArtifact) *Artifact {
	s.JarArtifact = v
	return s
}

func (s *Artifact) SetKind(v string) *Artifact {
	s.Kind = &v
	return s
}

func (s *Artifact) SetPythonArtifact(v *PythonArtifact) *Artifact {
	s.PythonArtifact = v
	return s
}

func (s *Artifact) SetSqlArtifact(v *SqlArtifact) *Artifact {
	s.SqlArtifact = v
	return s
}

type AsyncResourcePlanOperationResult struct {
	Message      *string `json:"message,omitempty" xml:"message,omitempty"`
	Plan         *string `json:"plan,omitempty" xml:"plan,omitempty"`
	TicketStatus *string `json:"ticketStatus,omitempty" xml:"ticketStatus,omitempty"`
}

func (s AsyncResourcePlanOperationResult) String() string {
	return tea.Prettify(s)
}

func (s AsyncResourcePlanOperationResult) GoString() string {
	return s.String()
}

func (s *AsyncResourcePlanOperationResult) SetMessage(v string) *AsyncResourcePlanOperationResult {
	s.Message = &v
	return s
}

func (s *AsyncResourcePlanOperationResult) SetPlan(v string) *AsyncResourcePlanOperationResult {
	s.Plan = &v
	return s
}

func (s *AsyncResourcePlanOperationResult) SetTicketStatus(v string) *AsyncResourcePlanOperationResult {
	s.TicketStatus = &v
	return s
}

type BasicResourceSetting struct {
	JobmanagerResourceSettingSpec  *BasicResourceSettingSpec `json:"jobmanagerResourceSettingSpec,omitempty" xml:"jobmanagerResourceSettingSpec,omitempty"`
	Parallelism                    *int64                    `json:"parallelism,omitempty" xml:"parallelism,omitempty"`
	TaskmanagerResourceSettingSpec *BasicResourceSettingSpec `json:"taskmanagerResourceSettingSpec,omitempty" xml:"taskmanagerResourceSettingSpec,omitempty"`
}

func (s BasicResourceSetting) String() string {
	return tea.Prettify(s)
}

func (s BasicResourceSetting) GoString() string {
	return s.String()
}

func (s *BasicResourceSetting) SetJobmanagerResourceSettingSpec(v *BasicResourceSettingSpec) *BasicResourceSetting {
	s.JobmanagerResourceSettingSpec = v
	return s
}

func (s *BasicResourceSetting) SetParallelism(v int64) *BasicResourceSetting {
	s.Parallelism = &v
	return s
}

func (s *BasicResourceSetting) SetTaskmanagerResourceSettingSpec(v *BasicResourceSettingSpec) *BasicResourceSetting {
	s.TaskmanagerResourceSettingSpec = v
	return s
}

type BasicResourceSettingSpec struct {
	Cpu    *float64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	Memory *string  `json:"memory,omitempty" xml:"memory,omitempty"`
}

func (s BasicResourceSettingSpec) String() string {
	return tea.Prettify(s)
}

func (s BasicResourceSettingSpec) GoString() string {
	return s.String()
}

func (s *BasicResourceSettingSpec) SetCpu(v float64) *BasicResourceSettingSpec {
	s.Cpu = &v
	return s
}

func (s *BasicResourceSettingSpec) SetMemory(v string) *BasicResourceSettingSpec {
	s.Memory = &v
	return s
}

type BatchResourceSetting struct {
	BasicResourceSetting *BasicResourceSetting `json:"basicResourceSetting,omitempty" xml:"basicResourceSetting,omitempty"`
	MaxSlot              *int64                `json:"maxSlot,omitempty" xml:"maxSlot,omitempty"`
}

func (s BatchResourceSetting) String() string {
	return tea.Prettify(s)
}

func (s BatchResourceSetting) GoString() string {
	return s.String()
}

func (s *BatchResourceSetting) SetBasicResourceSetting(v *BasicResourceSetting) *BatchResourceSetting {
	s.BasicResourceSetting = v
	return s
}

func (s *BatchResourceSetting) SetMaxSlot(v int64) *BatchResourceSetting {
	s.MaxSlot = &v
	return s
}

type BriefDeploymentTarget struct {
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s BriefDeploymentTarget) String() string {
	return tea.Prettify(s)
}

func (s BriefDeploymentTarget) GoString() string {
	return s.String()
}

func (s *BriefDeploymentTarget) SetMode(v string) *BriefDeploymentTarget {
	s.Mode = &v
	return s
}

func (s *BriefDeploymentTarget) SetName(v string) *BriefDeploymentTarget {
	s.Name = &v
	return s
}

type BriefResourceSetting struct {
	BatchResourceSetting     *BatchResourceSetting     `json:"batchResourceSetting,omitempty" xml:"batchResourceSetting,omitempty"`
	FlinkConf                map[string]interface{}    `json:"flinkConf,omitempty" xml:"flinkConf,omitempty"`
	StreamingResourceSetting *StreamingResourceSetting `json:"streamingResourceSetting,omitempty" xml:"streamingResourceSetting,omitempty"`
}

func (s BriefResourceSetting) String() string {
	return tea.Prettify(s)
}

func (s BriefResourceSetting) GoString() string {
	return s.String()
}

func (s *BriefResourceSetting) SetBatchResourceSetting(v *BatchResourceSetting) *BriefResourceSetting {
	s.BatchResourceSetting = v
	return s
}

func (s *BriefResourceSetting) SetFlinkConf(v map[string]interface{}) *BriefResourceSetting {
	s.FlinkConf = v
	return s
}

func (s *BriefResourceSetting) SetStreamingResourceSetting(v *StreamingResourceSetting) *BriefResourceSetting {
	s.StreamingResourceSetting = v
	return s
}

type Deployment struct {
	Artifact             *Artifact              `json:"artifact,omitempty" xml:"artifact,omitempty"`
	DeploymentHasChanged *bool                  `json:"deploymentHasChanged,omitempty" xml:"deploymentHasChanged,omitempty"`
	DeploymentId         *string                `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	DeploymentTarget     *BriefDeploymentTarget `json:"deploymentTarget,omitempty" xml:"deploymentTarget,omitempty"`
	Description          *string                `json:"description,omitempty" xml:"description,omitempty"`
	EngineVersion        *string                `json:"engineVersion,omitempty" xml:"engineVersion,omitempty"`
	ExecutionMode        *string                `json:"executionMode,omitempty" xml:"executionMode,omitempty"`
	FlinkConf            map[string]interface{} `json:"flinkConf,omitempty" xml:"flinkConf,omitempty"`
	JobSummary           *JobSummary            `json:"jobSummary,omitempty" xml:"jobSummary,omitempty"`
	Logging              *Logging               `json:"logging,omitempty" xml:"logging,omitempty"`
	Name                 *string                `json:"name,omitempty" xml:"name,omitempty"`
	Namespace            *string                `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s Deployment) String() string {
	return tea.Prettify(s)
}

func (s Deployment) GoString() string {
	return s.String()
}

func (s *Deployment) SetArtifact(v *Artifact) *Deployment {
	s.Artifact = v
	return s
}

func (s *Deployment) SetDeploymentHasChanged(v bool) *Deployment {
	s.DeploymentHasChanged = &v
	return s
}

func (s *Deployment) SetDeploymentId(v string) *Deployment {
	s.DeploymentId = &v
	return s
}

func (s *Deployment) SetDeploymentTarget(v *BriefDeploymentTarget) *Deployment {
	s.DeploymentTarget = v
	return s
}

func (s *Deployment) SetDescription(v string) *Deployment {
	s.Description = &v
	return s
}

func (s *Deployment) SetEngineVersion(v string) *Deployment {
	s.EngineVersion = &v
	return s
}

func (s *Deployment) SetExecutionMode(v string) *Deployment {
	s.ExecutionMode = &v
	return s
}

func (s *Deployment) SetFlinkConf(v map[string]interface{}) *Deployment {
	s.FlinkConf = v
	return s
}

func (s *Deployment) SetJobSummary(v *JobSummary) *Deployment {
	s.JobSummary = v
	return s
}

func (s *Deployment) SetLogging(v *Logging) *Deployment {
	s.Logging = v
	return s
}

func (s *Deployment) SetName(v string) *Deployment {
	s.Name = &v
	return s
}

func (s *Deployment) SetNamespace(v string) *Deployment {
	s.Namespace = &v
	return s
}

type DeploymentRestoreStrategy struct {
	AllowNonRestoredState *bool   `json:"allowNonRestoredState,omitempty" xml:"allowNonRestoredState,omitempty"`
	JobStartTimeInMs      *int64  `json:"jobStartTimeInMs,omitempty" xml:"jobStartTimeInMs,omitempty"`
	Kind                  *string `json:"kind,omitempty" xml:"kind,omitempty"`
	SavepointId           *string `json:"savepointId,omitempty" xml:"savepointId,omitempty"`
}

func (s DeploymentRestoreStrategy) String() string {
	return tea.Prettify(s)
}

func (s DeploymentRestoreStrategy) GoString() string {
	return s.String()
}

func (s *DeploymentRestoreStrategy) SetAllowNonRestoredState(v bool) *DeploymentRestoreStrategy {
	s.AllowNonRestoredState = &v
	return s
}

func (s *DeploymentRestoreStrategy) SetJobStartTimeInMs(v int64) *DeploymentRestoreStrategy {
	s.JobStartTimeInMs = &v
	return s
}

func (s *DeploymentRestoreStrategy) SetKind(v string) *DeploymentRestoreStrategy {
	s.Kind = &v
	return s
}

func (s *DeploymentRestoreStrategy) SetSavepointId(v string) *DeploymentRestoreStrategy {
	s.SavepointId = &v
	return s
}

type DeploymentTarget struct {
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s DeploymentTarget) String() string {
	return tea.Prettify(s)
}

func (s DeploymentTarget) GoString() string {
	return s.String()
}

func (s *DeploymentTarget) SetName(v string) *DeploymentTarget {
	s.Name = &v
	return s
}

func (s *DeploymentTarget) SetNamespace(v string) *DeploymentTarget {
	s.Namespace = &v
	return s
}

type EngineVersionMetadata struct {
	EngineVersion *string                         `json:"engineVersion,omitempty" xml:"engineVersion,omitempty"`
	Features      *EngineVersionSupportedFeatures `json:"features,omitempty" xml:"features,omitempty"`
	Status        *string                         `json:"status,omitempty" xml:"status,omitempty"`
}

func (s EngineVersionMetadata) String() string {
	return tea.Prettify(s)
}

func (s EngineVersionMetadata) GoString() string {
	return s.String()
}

func (s *EngineVersionMetadata) SetEngineVersion(v string) *EngineVersionMetadata {
	s.EngineVersion = &v
	return s
}

func (s *EngineVersionMetadata) SetFeatures(v *EngineVersionSupportedFeatures) *EngineVersionMetadata {
	s.Features = v
	return s
}

func (s *EngineVersionMetadata) SetStatus(v string) *EngineVersionMetadata {
	s.Status = &v
	return s
}

type EngineVersionMetadataIndex struct {
	DefaultEngineVersion  *string                  `json:"defaultEngineVersion,omitempty" xml:"defaultEngineVersion,omitempty"`
	EngineVersionMetadata []*EngineVersionMetadata `json:"engineVersionMetadata,omitempty" xml:"engineVersionMetadata,omitempty" type:"Repeated"`
}

func (s EngineVersionMetadataIndex) String() string {
	return tea.Prettify(s)
}

func (s EngineVersionMetadataIndex) GoString() string {
	return s.String()
}

func (s *EngineVersionMetadataIndex) SetDefaultEngineVersion(v string) *EngineVersionMetadataIndex {
	s.DefaultEngineVersion = &v
	return s
}

func (s *EngineVersionMetadataIndex) SetEngineVersionMetadata(v []*EngineVersionMetadata) *EngineVersionMetadataIndex {
	s.EngineVersionMetadata = v
	return s
}

type EngineVersionSupportedFeatures struct {
	SupportNativeSavepoint *bool `json:"supportNativeSavepoint,omitempty" xml:"supportNativeSavepoint,omitempty"`
	UseForSqlDeployments   *bool `json:"useForSqlDeployments,omitempty" xml:"useForSqlDeployments,omitempty"`
}

func (s EngineVersionSupportedFeatures) String() string {
	return tea.Prettify(s)
}

func (s EngineVersionSupportedFeatures) GoString() string {
	return s.String()
}

func (s *EngineVersionSupportedFeatures) SetSupportNativeSavepoint(v bool) *EngineVersionSupportedFeatures {
	s.SupportNativeSavepoint = &v
	return s
}

func (s *EngineVersionSupportedFeatures) SetUseForSqlDeployments(v bool) *EngineVersionSupportedFeatures {
	s.UseForSqlDeployments = &v
	return s
}

type ExpertResourceSetting struct {
	JobmanagerResourceSettingSpec *BasicResourceSettingSpec `json:"jobmanagerResourceSettingSpec,omitempty" xml:"jobmanagerResourceSettingSpec,omitempty"`
	ResourcePlan                  *string                   `json:"resourcePlan,omitempty" xml:"resourcePlan,omitempty"`
}

func (s ExpertResourceSetting) String() string {
	return tea.Prettify(s)
}

func (s ExpertResourceSetting) GoString() string {
	return s.String()
}

func (s *ExpertResourceSetting) SetJobmanagerResourceSettingSpec(v *BasicResourceSettingSpec) *ExpertResourceSetting {
	s.JobmanagerResourceSettingSpec = v
	return s
}

func (s *ExpertResourceSetting) SetResourcePlan(v string) *ExpertResourceSetting {
	s.ResourcePlan = &v
	return s
}

type JarArtifact struct {
	AdditionalDependencies []*string `json:"additionalDependencies,omitempty" xml:"additionalDependencies,omitempty" type:"Repeated"`
	EntryClass             *string   `json:"entryClass,omitempty" xml:"entryClass,omitempty"`
	JarUri                 *string   `json:"jarUri,omitempty" xml:"jarUri,omitempty"`
	MainArgs               *string   `json:"mainArgs,omitempty" xml:"mainArgs,omitempty"`
}

func (s JarArtifact) String() string {
	return tea.Prettify(s)
}

func (s JarArtifact) GoString() string {
	return s.String()
}

func (s *JarArtifact) SetAdditionalDependencies(v []*string) *JarArtifact {
	s.AdditionalDependencies = v
	return s
}

func (s *JarArtifact) SetEntryClass(v string) *JarArtifact {
	s.EntryClass = &v
	return s
}

func (s *JarArtifact) SetJarUri(v string) *JarArtifact {
	s.JarUri = &v
	return s
}

func (s *JarArtifact) SetMainArgs(v string) *JarArtifact {
	s.MainArgs = &v
	return s
}

type Job struct {
	Artifact                 *Artifact                  `json:"artifact,omitempty" xml:"artifact,omitempty"`
	BatchResourceSetting     *BatchResourceSetting      `json:"batchResourceSetting,omitempty" xml:"batchResourceSetting,omitempty"`
	DeploymentId             *string                    `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	DeploymentName           *string                    `json:"deploymentName,omitempty" xml:"deploymentName,omitempty"`
	EndTime                  *int64                     `json:"endTime,omitempty" xml:"endTime,omitempty"`
	EngineVersion            *string                    `json:"engineVersion,omitempty" xml:"engineVersion,omitempty"`
	ExecutionMode            *string                    `json:"executionMode,omitempty" xml:"executionMode,omitempty"`
	FlinkConf                map[string]interface{}     `json:"flinkConf,omitempty" xml:"flinkConf,omitempty"`
	JobId                    *string                    `json:"jobId,omitempty" xml:"jobId,omitempty"`
	Logging                  *Logging                   `json:"logging,omitempty" xml:"logging,omitempty"`
	Metric                   *JobMetric                 `json:"metric,omitempty" xml:"metric,omitempty"`
	Namespace                *string                    `json:"namespace,omitempty" xml:"namespace,omitempty"`
	RestoreStrategy          *DeploymentRestoreStrategy `json:"restoreStrategy,omitempty" xml:"restoreStrategy,omitempty"`
	SessionClusterName       *string                    `json:"sessionClusterName,omitempty" xml:"sessionClusterName,omitempty"`
	StartTime                *int64                     `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status                   *JobStatus                 `json:"status,omitempty" xml:"status,omitempty"`
	StreamingResourceSetting *StreamingResourceSetting  `json:"streamingResourceSetting,omitempty" xml:"streamingResourceSetting,omitempty"`
}

func (s Job) String() string {
	return tea.Prettify(s)
}

func (s Job) GoString() string {
	return s.String()
}

func (s *Job) SetArtifact(v *Artifact) *Job {
	s.Artifact = v
	return s
}

func (s *Job) SetBatchResourceSetting(v *BatchResourceSetting) *Job {
	s.BatchResourceSetting = v
	return s
}

func (s *Job) SetDeploymentId(v string) *Job {
	s.DeploymentId = &v
	return s
}

func (s *Job) SetDeploymentName(v string) *Job {
	s.DeploymentName = &v
	return s
}

func (s *Job) SetEndTime(v int64) *Job {
	s.EndTime = &v
	return s
}

func (s *Job) SetEngineVersion(v string) *Job {
	s.EngineVersion = &v
	return s
}

func (s *Job) SetExecutionMode(v string) *Job {
	s.ExecutionMode = &v
	return s
}

func (s *Job) SetFlinkConf(v map[string]interface{}) *Job {
	s.FlinkConf = v
	return s
}

func (s *Job) SetJobId(v string) *Job {
	s.JobId = &v
	return s
}

func (s *Job) SetLogging(v *Logging) *Job {
	s.Logging = v
	return s
}

func (s *Job) SetMetric(v *JobMetric) *Job {
	s.Metric = v
	return s
}

func (s *Job) SetNamespace(v string) *Job {
	s.Namespace = &v
	return s
}

func (s *Job) SetRestoreStrategy(v *DeploymentRestoreStrategy) *Job {
	s.RestoreStrategy = v
	return s
}

func (s *Job) SetSessionClusterName(v string) *Job {
	s.SessionClusterName = &v
	return s
}

func (s *Job) SetStartTime(v int64) *Job {
	s.StartTime = &v
	return s
}

func (s *Job) SetStatus(v *JobStatus) *Job {
	s.Status = v
	return s
}

func (s *Job) SetStreamingResourceSetting(v *StreamingResourceSetting) *Job {
	s.StreamingResourceSetting = v
	return s
}

type JobFailure struct {
	FailedAt *int64  `json:"failedAt,omitempty" xml:"failedAt,omitempty"`
	Message  *string `json:"message,omitempty" xml:"message,omitempty"`
	Reason   *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s JobFailure) String() string {
	return tea.Prettify(s)
}

func (s JobFailure) GoString() string {
	return s.String()
}

func (s *JobFailure) SetFailedAt(v int64) *JobFailure {
	s.FailedAt = &v
	return s
}

func (s *JobFailure) SetMessage(v string) *JobFailure {
	s.Message = &v
	return s
}

func (s *JobFailure) SetReason(v string) *JobFailure {
	s.Reason = &v
	return s
}

type JobMetric struct {
	TotalCpu        *float64 `json:"totalCpu,omitempty" xml:"totalCpu,omitempty"`
	TotalMemoryByte *int64   `json:"totalMemoryByte,omitempty" xml:"totalMemoryByte,omitempty"`
}

func (s JobMetric) String() string {
	return tea.Prettify(s)
}

func (s JobMetric) GoString() string {
	return s.String()
}

func (s *JobMetric) SetTotalCpu(v float64) *JobMetric {
	s.TotalCpu = &v
	return s
}

func (s *JobMetric) SetTotalMemoryByte(v int64) *JobMetric {
	s.TotalMemoryByte = &v
	return s
}

type JobStatus struct {
	CurrentJobStatus *string           `json:"currentJobStatus,omitempty" xml:"currentJobStatus,omitempty"`
	Failure          *JobFailure       `json:"failure,omitempty" xml:"failure,omitempty"`
	Running          *JobStatusRunning `json:"running,omitempty" xml:"running,omitempty"`
}

func (s JobStatus) String() string {
	return tea.Prettify(s)
}

func (s JobStatus) GoString() string {
	return s.String()
}

func (s *JobStatus) SetCurrentJobStatus(v string) *JobStatus {
	s.CurrentJobStatus = &v
	return s
}

func (s *JobStatus) SetFailure(v *JobFailure) *JobStatus {
	s.Failure = v
	return s
}

func (s *JobStatus) SetRunning(v *JobStatusRunning) *JobStatus {
	s.Running = v
	return s
}

type JobStatusRunning struct {
	ObservedFlinkJobRestarts *int64  `json:"observedFlinkJobRestarts,omitempty" xml:"observedFlinkJobRestarts,omitempty"`
	ObservedFlinkJobStatus   *string `json:"observedFlinkJobStatus,omitempty" xml:"observedFlinkJobStatus,omitempty"`
}

func (s JobStatusRunning) String() string {
	return tea.Prettify(s)
}

func (s JobStatusRunning) GoString() string {
	return s.String()
}

func (s *JobStatusRunning) SetObservedFlinkJobRestarts(v int64) *JobStatusRunning {
	s.ObservedFlinkJobRestarts = &v
	return s
}

func (s *JobStatusRunning) SetObservedFlinkJobStatus(v string) *JobStatusRunning {
	s.ObservedFlinkJobStatus = &v
	return s
}

type JobSummary struct {
	Cancelled  *int32 `json:"cancelled,omitempty" xml:"cancelled,omitempty"`
	Cancelling *int32 `json:"cancelling,omitempty" xml:"cancelling,omitempty"`
	Failed     *int32 `json:"failed,omitempty" xml:"failed,omitempty"`
	Finished   *int32 `json:"finished,omitempty" xml:"finished,omitempty"`
	Running    *int32 `json:"running,omitempty" xml:"running,omitempty"`
	Starting   *int32 `json:"starting,omitempty" xml:"starting,omitempty"`
}

func (s JobSummary) String() string {
	return tea.Prettify(s)
}

func (s JobSummary) GoString() string {
	return s.String()
}

func (s *JobSummary) SetCancelled(v int32) *JobSummary {
	s.Cancelled = &v
	return s
}

func (s *JobSummary) SetCancelling(v int32) *JobSummary {
	s.Cancelling = &v
	return s
}

func (s *JobSummary) SetFailed(v int32) *JobSummary {
	s.Failed = &v
	return s
}

func (s *JobSummary) SetFinished(v int32) *JobSummary {
	s.Finished = &v
	return s
}

func (s *JobSummary) SetRunning(v int32) *JobSummary {
	s.Running = &v
	return s
}

func (s *JobSummary) SetStarting(v int32) *JobSummary {
	s.Starting = &v
	return s
}

type Log4jLogger struct {
	LoggerLevel *string `json:"loggerLevel,omitempty" xml:"loggerLevel,omitempty"`
	LoggerName  *string `json:"loggerName,omitempty" xml:"loggerName,omitempty"`
}

func (s Log4jLogger) String() string {
	return tea.Prettify(s)
}

func (s Log4jLogger) GoString() string {
	return s.String()
}

func (s *Log4jLogger) SetLoggerLevel(v string) *Log4jLogger {
	s.LoggerLevel = &v
	return s
}

func (s *Log4jLogger) SetLoggerName(v string) *Log4jLogger {
	s.LoggerName = &v
	return s
}

type LogReservePolicy struct {
	ExpirationDays *int64 `json:"expirationDays,omitempty" xml:"expirationDays,omitempty"`
	OpenHistory    *bool  `json:"openHistory,omitempty" xml:"openHistory,omitempty"`
}

func (s LogReservePolicy) String() string {
	return tea.Prettify(s)
}

func (s LogReservePolicy) GoString() string {
	return s.String()
}

func (s *LogReservePolicy) SetExpirationDays(v int64) *LogReservePolicy {
	s.ExpirationDays = &v
	return s
}

func (s *LogReservePolicy) SetOpenHistory(v bool) *LogReservePolicy {
	s.OpenHistory = &v
	return s
}

type Logging struct {
	Log4j2ConfigurationTemplate *string           `json:"log4j2ConfigurationTemplate,omitempty" xml:"log4j2ConfigurationTemplate,omitempty"`
	Log4jLoggers                []*Log4jLogger    `json:"log4jLoggers,omitempty" xml:"log4jLoggers,omitempty" type:"Repeated"`
	LogReservePolicy            *LogReservePolicy `json:"logReservePolicy,omitempty" xml:"logReservePolicy,omitempty"`
	LoggingProfile              *string           `json:"loggingProfile,omitempty" xml:"loggingProfile,omitempty"`
}

func (s Logging) String() string {
	return tea.Prettify(s)
}

func (s Logging) GoString() string {
	return s.String()
}

func (s *Logging) SetLog4j2ConfigurationTemplate(v string) *Logging {
	s.Log4j2ConfigurationTemplate = &v
	return s
}

func (s *Logging) SetLog4jLoggers(v []*Log4jLogger) *Logging {
	s.Log4jLoggers = v
	return s
}

func (s *Logging) SetLogReservePolicy(v *LogReservePolicy) *Logging {
	s.LogReservePolicy = v
	return s
}

func (s *Logging) SetLoggingProfile(v string) *Logging {
	s.LoggingProfile = &v
	return s
}

type PythonArtifact struct {
	AdditionalDependencies    []*string `json:"additionalDependencies,omitempty" xml:"additionalDependencies,omitempty" type:"Repeated"`
	AdditionalPythonArchives  []*string `json:"additionalPythonArchives,omitempty" xml:"additionalPythonArchives,omitempty" type:"Repeated"`
	AdditionalPythonLibraries []*string `json:"additionalPythonLibraries,omitempty" xml:"additionalPythonLibraries,omitempty" type:"Repeated"`
	EntryModule               *string   `json:"entryModule,omitempty" xml:"entryModule,omitempty"`
	MainArgs                  *string   `json:"mainArgs,omitempty" xml:"mainArgs,omitempty"`
	PythonArtifactUri         *string   `json:"pythonArtifactUri,omitempty" xml:"pythonArtifactUri,omitempty"`
}

func (s PythonArtifact) String() string {
	return tea.Prettify(s)
}

func (s PythonArtifact) GoString() string {
	return s.String()
}

func (s *PythonArtifact) SetAdditionalDependencies(v []*string) *PythonArtifact {
	s.AdditionalDependencies = v
	return s
}

func (s *PythonArtifact) SetAdditionalPythonArchives(v []*string) *PythonArtifact {
	s.AdditionalPythonArchives = v
	return s
}

func (s *PythonArtifact) SetAdditionalPythonLibraries(v []*string) *PythonArtifact {
	s.AdditionalPythonLibraries = v
	return s
}

func (s *PythonArtifact) SetEntryModule(v string) *PythonArtifact {
	s.EntryModule = &v
	return s
}

func (s *PythonArtifact) SetMainArgs(v string) *PythonArtifact {
	s.MainArgs = &v
	return s
}

func (s *PythonArtifact) SetPythonArtifactUri(v string) *PythonArtifact {
	s.PythonArtifactUri = &v
	return s
}

type Savepoint struct {
	CreatedAt            *int64           `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	DeploymentId         *string          `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	Description          *string          `json:"description,omitempty" xml:"description,omitempty"`
	JobId                *string          `json:"jobId,omitempty" xml:"jobId,omitempty"`
	ModifiedAt           *int64           `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	Namespace            *string          `json:"namespace,omitempty" xml:"namespace,omitempty"`
	NativeFormat         *bool            `json:"nativeFormat,omitempty" xml:"nativeFormat,omitempty"`
	SavepointId          *string          `json:"savepointId,omitempty" xml:"savepointId,omitempty"`
	SavepointLocation    *string          `json:"savepointLocation,omitempty" xml:"savepointLocation,omitempty"`
	SavepointOrigin      *string          `json:"savepointOrigin,omitempty" xml:"savepointOrigin,omitempty"`
	Status               *SavepointStatus `json:"status,omitempty" xml:"status,omitempty"`
	StopWithDrainEnabled *bool            `json:"stopWithDrainEnabled,omitempty" xml:"stopWithDrainEnabled,omitempty"`
}

func (s Savepoint) String() string {
	return tea.Prettify(s)
}

func (s Savepoint) GoString() string {
	return s.String()
}

func (s *Savepoint) SetCreatedAt(v int64) *Savepoint {
	s.CreatedAt = &v
	return s
}

func (s *Savepoint) SetDeploymentId(v string) *Savepoint {
	s.DeploymentId = &v
	return s
}

func (s *Savepoint) SetDescription(v string) *Savepoint {
	s.Description = &v
	return s
}

func (s *Savepoint) SetJobId(v string) *Savepoint {
	s.JobId = &v
	return s
}

func (s *Savepoint) SetModifiedAt(v int64) *Savepoint {
	s.ModifiedAt = &v
	return s
}

func (s *Savepoint) SetNamespace(v string) *Savepoint {
	s.Namespace = &v
	return s
}

func (s *Savepoint) SetNativeFormat(v bool) *Savepoint {
	s.NativeFormat = &v
	return s
}

func (s *Savepoint) SetSavepointId(v string) *Savepoint {
	s.SavepointId = &v
	return s
}

func (s *Savepoint) SetSavepointLocation(v string) *Savepoint {
	s.SavepointLocation = &v
	return s
}

func (s *Savepoint) SetSavepointOrigin(v string) *Savepoint {
	s.SavepointOrigin = &v
	return s
}

func (s *Savepoint) SetStatus(v *SavepointStatus) *Savepoint {
	s.Status = v
	return s
}

func (s *Savepoint) SetStopWithDrainEnabled(v bool) *Savepoint {
	s.StopWithDrainEnabled = &v
	return s
}

type SavepointFailure struct {
	FailedAt *int64  `json:"failedAt,omitempty" xml:"failedAt,omitempty"`
	Message  *string `json:"message,omitempty" xml:"message,omitempty"`
	Reason   *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s SavepointFailure) String() string {
	return tea.Prettify(s)
}

func (s SavepointFailure) GoString() string {
	return s.String()
}

func (s *SavepointFailure) SetFailedAt(v int64) *SavepointFailure {
	s.FailedAt = &v
	return s
}

func (s *SavepointFailure) SetMessage(v string) *SavepointFailure {
	s.Message = &v
	return s
}

func (s *SavepointFailure) SetReason(v string) *SavepointFailure {
	s.Reason = &v
	return s
}

type SavepointStatus struct {
	Failure *SavepointFailure `json:"failure,omitempty" xml:"failure,omitempty"`
	State   *string           `json:"state,omitempty" xml:"state,omitempty"`
}

func (s SavepointStatus) String() string {
	return tea.Prettify(s)
}

func (s SavepointStatus) GoString() string {
	return s.String()
}

func (s *SavepointStatus) SetFailure(v *SavepointFailure) *SavepointStatus {
	s.Failure = v
	return s
}

func (s *SavepointStatus) SetState(v string) *SavepointStatus {
	s.State = &v
	return s
}

type SqlArtifact struct {
	AdditionalDependencies []*string `json:"additionalDependencies,omitempty" xml:"additionalDependencies,omitempty" type:"Repeated"`
	SqlScript              *string   `json:"sqlScript,omitempty" xml:"sqlScript,omitempty"`
}

func (s SqlArtifact) String() string {
	return tea.Prettify(s)
}

func (s SqlArtifact) GoString() string {
	return s.String()
}

func (s *SqlArtifact) SetAdditionalDependencies(v []*string) *SqlArtifact {
	s.AdditionalDependencies = v
	return s
}

func (s *SqlArtifact) SetSqlScript(v string) *SqlArtifact {
	s.SqlScript = &v
	return s
}

type StartJobRequestBody struct {
	DeploymentId        *string                    `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	ResourceSettingSpec *BriefResourceSetting      `json:"resourceSettingSpec,omitempty" xml:"resourceSettingSpec,omitempty"`
	RestoreStrategy     *DeploymentRestoreStrategy `json:"restoreStrategy,omitempty" xml:"restoreStrategy,omitempty"`
}

func (s StartJobRequestBody) String() string {
	return tea.Prettify(s)
}

func (s StartJobRequestBody) GoString() string {
	return s.String()
}

func (s *StartJobRequestBody) SetDeploymentId(v string) *StartJobRequestBody {
	s.DeploymentId = &v
	return s
}

func (s *StartJobRequestBody) SetResourceSettingSpec(v *BriefResourceSetting) *StartJobRequestBody {
	s.ResourceSettingSpec = v
	return s
}

func (s *StartJobRequestBody) SetRestoreStrategy(v *DeploymentRestoreStrategy) *StartJobRequestBody {
	s.RestoreStrategy = v
	return s
}

type StopJobRequestBody struct {
	StopStrategy *string `json:"stopStrategy,omitempty" xml:"stopStrategy,omitempty"`
}

func (s StopJobRequestBody) String() string {
	return tea.Prettify(s)
}

func (s StopJobRequestBody) GoString() string {
	return s.String()
}

func (s *StopJobRequestBody) SetStopStrategy(v string) *StopJobRequestBody {
	s.StopStrategy = &v
	return s
}

type StreamingResourceSetting struct {
	BasicResourceSetting  *BasicResourceSetting  `json:"basicResourceSetting,omitempty" xml:"basicResourceSetting,omitempty"`
	ExpertResourceSetting *ExpertResourceSetting `json:"expertResourceSetting,omitempty" xml:"expertResourceSetting,omitempty"`
	ResourceSettingMode   *string                `json:"resourceSettingMode,omitempty" xml:"resourceSettingMode,omitempty"`
}

func (s StreamingResourceSetting) String() string {
	return tea.Prettify(s)
}

func (s StreamingResourceSetting) GoString() string {
	return s.String()
}

func (s *StreamingResourceSetting) SetBasicResourceSetting(v *BasicResourceSetting) *StreamingResourceSetting {
	s.BasicResourceSetting = v
	return s
}

func (s *StreamingResourceSetting) SetExpertResourceSetting(v *ExpertResourceSetting) *StreamingResourceSetting {
	s.ExpertResourceSetting = v
	return s
}

func (s *StreamingResourceSetting) SetResourceSettingMode(v string) *StreamingResourceSetting {
	s.ResourceSettingMode = &v
	return s
}

type Variable struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Kind        *string `json:"kind,omitempty" xml:"kind,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Value       *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s Variable) String() string {
	return tea.Prettify(s)
}

func (s Variable) GoString() string {
	return s.String()
}

func (s *Variable) SetDescription(v string) *Variable {
	s.Description = &v
	return s
}

func (s *Variable) SetKind(v string) *Variable {
	s.Kind = &v
	return s
}

func (s *Variable) SetName(v string) *Variable {
	s.Name = &v
	return s
}

func (s *Variable) SetValue(v string) *Variable {
	s.Value = &v
	return s
}

type CreateDeploymentHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateDeploymentHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentHeaders) GoString() string {
	return s.String()
}

func (s *CreateDeploymentHeaders) SetCommonHeaders(v map[string]*string) *CreateDeploymentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDeploymentHeaders) SetWorkspace(v string) *CreateDeploymentHeaders {
	s.Workspace = &v
	return s
}

type CreateDeploymentRequest struct {
	Body *Deployment `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeploymentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentRequest) GoString() string {
	return s.String()
}

func (s *CreateDeploymentRequest) SetBody(v *Deployment) *CreateDeploymentRequest {
	s.Body = v
	return s
}

type CreateDeploymentResponseBody struct {
	Data         *Deployment `json:"data,omitempty" xml:"data,omitempty"`
	ErrorCode    *string     `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string     `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32      `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId    *string     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool       `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateDeploymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeploymentResponseBody) SetData(v *Deployment) *CreateDeploymentResponseBody {
	s.Data = v
	return s
}

func (s *CreateDeploymentResponseBody) SetErrorCode(v string) *CreateDeploymentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDeploymentResponseBody) SetErrorMessage(v string) *CreateDeploymentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDeploymentResponseBody) SetHttpCode(v int32) *CreateDeploymentResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateDeploymentResponseBody) SetRequestId(v string) *CreateDeploymentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeploymentResponseBody) SetSuccess(v bool) *CreateDeploymentResponseBody {
	s.Success = &v
	return s
}

type CreateDeploymentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentResponse) GoString() string {
	return s.String()
}

func (s *CreateDeploymentResponse) SetHeaders(v map[string]*string) *CreateDeploymentResponse {
	s.Headers = v
	return s
}

func (s *CreateDeploymentResponse) SetStatusCode(v int32) *CreateDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeploymentResponse) SetBody(v *CreateDeploymentResponseBody) *CreateDeploymentResponse {
	s.Body = v
	return s
}

type CreateSavepointHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateSavepointHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateSavepointHeaders) GoString() string {
	return s.String()
}

func (s *CreateSavepointHeaders) SetCommonHeaders(v map[string]*string) *CreateSavepointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSavepointHeaders) SetWorkspace(v string) *CreateSavepointHeaders {
	s.Workspace = &v
	return s
}

type CreateSavepointRequest struct {
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	Description  *string `json:"description,omitempty" xml:"description,omitempty"`
	NativeFormat *bool   `json:"nativeFormat,omitempty" xml:"nativeFormat,omitempty"`
}

func (s CreateSavepointRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSavepointRequest) GoString() string {
	return s.String()
}

func (s *CreateSavepointRequest) SetDeploymentId(v string) *CreateSavepointRequest {
	s.DeploymentId = &v
	return s
}

func (s *CreateSavepointRequest) SetDescription(v string) *CreateSavepointRequest {
	s.Description = &v
	return s
}

func (s *CreateSavepointRequest) SetNativeFormat(v bool) *CreateSavepointRequest {
	s.NativeFormat = &v
	return s
}

type CreateSavepointResponseBody struct {
	Data         *Savepoint `json:"data,omitempty" xml:"data,omitempty"`
	ErrorCode    *string    `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string    `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32     `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId    *string    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateSavepointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSavepointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSavepointResponseBody) SetData(v *Savepoint) *CreateSavepointResponseBody {
	s.Data = v
	return s
}

func (s *CreateSavepointResponseBody) SetErrorCode(v string) *CreateSavepointResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateSavepointResponseBody) SetErrorMessage(v string) *CreateSavepointResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateSavepointResponseBody) SetHttpCode(v int32) *CreateSavepointResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateSavepointResponseBody) SetRequestId(v string) *CreateSavepointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSavepointResponseBody) SetSuccess(v bool) *CreateSavepointResponseBody {
	s.Success = &v
	return s
}

type CreateSavepointResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSavepointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSavepointResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSavepointResponse) GoString() string {
	return s.String()
}

func (s *CreateSavepointResponse) SetHeaders(v map[string]*string) *CreateSavepointResponse {
	s.Headers = v
	return s
}

func (s *CreateSavepointResponse) SetStatusCode(v int32) *CreateSavepointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSavepointResponse) SetBody(v *CreateSavepointResponseBody) *CreateSavepointResponse {
	s.Body = v
	return s
}

type CreateVariableHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateVariableHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateVariableHeaders) GoString() string {
	return s.String()
}

func (s *CreateVariableHeaders) SetCommonHeaders(v map[string]*string) *CreateVariableHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateVariableHeaders) SetWorkspace(v string) *CreateVariableHeaders {
	s.Workspace = &v
	return s
}

type CreateVariableRequest struct {
	Body *Variable `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVariableRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVariableRequest) GoString() string {
	return s.String()
}

func (s *CreateVariableRequest) SetBody(v *Variable) *CreateVariableRequest {
	s.Body = v
	return s
}

type CreateVariableResponseBody struct {
	Data         *Variable `json:"data,omitempty" xml:"data,omitempty"`
	ErrorCode    *string   `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string   `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32    `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId    *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateVariableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVariableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVariableResponseBody) SetData(v *Variable) *CreateVariableResponseBody {
	s.Data = v
	return s
}

func (s *CreateVariableResponseBody) SetErrorCode(v string) *CreateVariableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateVariableResponseBody) SetErrorMessage(v string) *CreateVariableResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateVariableResponseBody) SetHttpCode(v int32) *CreateVariableResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateVariableResponseBody) SetRequestId(v string) *CreateVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVariableResponseBody) SetSuccess(v bool) *CreateVariableResponseBody {
	s.Success = &v
	return s
}

type CreateVariableResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVariableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVariableResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVariableResponse) GoString() string {
	return s.String()
}

func (s *CreateVariableResponse) SetHeaders(v map[string]*string) *CreateVariableResponse {
	s.Headers = v
	return s
}

func (s *CreateVariableResponse) SetStatusCode(v int32) *CreateVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVariableResponse) SetBody(v *CreateVariableResponseBody) *CreateVariableResponse {
	s.Body = v
	return s
}

type DeleteDeploymentHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteDeploymentHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeploymentHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentHeaders) SetCommonHeaders(v map[string]*string) *DeleteDeploymentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDeploymentHeaders) SetWorkspace(v string) *DeleteDeploymentHeaders {
	s.Workspace = &v
	return s
}

type DeleteDeploymentResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32  `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteDeploymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentResponseBody) SetErrorCode(v string) *DeleteDeploymentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDeploymentResponseBody) SetErrorMessage(v string) *DeleteDeploymentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDeploymentResponseBody) SetHttpCode(v int32) *DeleteDeploymentResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteDeploymentResponseBody) SetRequestId(v string) *DeleteDeploymentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeploymentResponseBody) SetSuccess(v bool) *DeleteDeploymentResponseBody {
	s.Success = &v
	return s
}

type DeleteDeploymentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeploymentResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentResponse) SetHeaders(v map[string]*string) *DeleteDeploymentResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeploymentResponse) SetStatusCode(v int32) *DeleteDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeploymentResponse) SetBody(v *DeleteDeploymentResponseBody) *DeleteDeploymentResponse {
	s.Body = v
	return s
}

type DeleteJobHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteJobHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobHeaders) GoString() string {
	return s.String()
}

func (s *DeleteJobHeaders) SetCommonHeaders(v map[string]*string) *DeleteJobHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteJobHeaders) SetWorkspace(v string) *DeleteJobHeaders {
	s.Workspace = &v
	return s
}

type DeleteJobResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32  `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobResponseBody) SetErrorCode(v string) *DeleteJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteJobResponseBody) SetErrorMessage(v string) *DeleteJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteJobResponseBody) SetHttpCode(v int32) *DeleteJobResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteJobResponseBody) SetRequestId(v string) *DeleteJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteJobResponseBody) SetSuccess(v bool) *DeleteJobResponseBody {
	s.Success = &v
	return s
}

type DeleteJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobResponse) SetHeaders(v map[string]*string) *DeleteJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobResponse) SetStatusCode(v int32) *DeleteJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteJobResponse) SetBody(v *DeleteJobResponseBody) *DeleteJobResponse {
	s.Body = v
	return s
}

type DeleteSavepointHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteSavepointHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteSavepointHeaders) GoString() string {
	return s.String()
}

func (s *DeleteSavepointHeaders) SetCommonHeaders(v map[string]*string) *DeleteSavepointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteSavepointHeaders) SetWorkspace(v string) *DeleteSavepointHeaders {
	s.Workspace = &v
	return s
}

type DeleteSavepointResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32  `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteSavepointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSavepointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSavepointResponseBody) SetErrorCode(v string) *DeleteSavepointResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteSavepointResponseBody) SetErrorMessage(v string) *DeleteSavepointResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteSavepointResponseBody) SetHttpCode(v int32) *DeleteSavepointResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteSavepointResponseBody) SetRequestId(v string) *DeleteSavepointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSavepointResponseBody) SetSuccess(v bool) *DeleteSavepointResponseBody {
	s.Success = &v
	return s
}

type DeleteSavepointResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSavepointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSavepointResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSavepointResponse) GoString() string {
	return s.String()
}

func (s *DeleteSavepointResponse) SetHeaders(v map[string]*string) *DeleteSavepointResponse {
	s.Headers = v
	return s
}

func (s *DeleteSavepointResponse) SetStatusCode(v int32) *DeleteSavepointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSavepointResponse) SetBody(v *DeleteSavepointResponseBody) *DeleteSavepointResponse {
	s.Body = v
	return s
}

type DeleteVariableHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteVariableHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteVariableHeaders) GoString() string {
	return s.String()
}

func (s *DeleteVariableHeaders) SetCommonHeaders(v map[string]*string) *DeleteVariableHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteVariableHeaders) SetWorkspace(v string) *DeleteVariableHeaders {
	s.Workspace = &v
	return s
}

type DeleteVariableResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32  `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteVariableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVariableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVariableResponseBody) SetErrorCode(v string) *DeleteVariableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteVariableResponseBody) SetErrorMessage(v string) *DeleteVariableResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteVariableResponseBody) SetHttpCode(v int32) *DeleteVariableResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteVariableResponseBody) SetRequestId(v string) *DeleteVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVariableResponseBody) SetSuccess(v bool) *DeleteVariableResponseBody {
	s.Success = &v
	return s
}

type DeleteVariableResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVariableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVariableResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVariableResponse) GoString() string {
	return s.String()
}

func (s *DeleteVariableResponse) SetHeaders(v map[string]*string) *DeleteVariableResponse {
	s.Headers = v
	return s
}

func (s *DeleteVariableResponse) SetStatusCode(v int32) *DeleteVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVariableResponse) SetBody(v *DeleteVariableResponseBody) *DeleteVariableResponse {
	s.Body = v
	return s
}

type FlinkApiProxyHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s FlinkApiProxyHeaders) String() string {
	return tea.Prettify(s)
}

func (s FlinkApiProxyHeaders) GoString() string {
	return s.String()
}

func (s *FlinkApiProxyHeaders) SetCommonHeaders(v map[string]*string) *FlinkApiProxyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlinkApiProxyHeaders) SetWorkspace(v string) *FlinkApiProxyHeaders {
	s.Workspace = &v
	return s
}

type FlinkApiProxyRequest struct {
	FlinkApiPath *string `json:"flinkApiPath,omitempty" xml:"flinkApiPath,omitempty"`
	Namespace    *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	ResourceId   *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s FlinkApiProxyRequest) String() string {
	return tea.Prettify(s)
}

func (s FlinkApiProxyRequest) GoString() string {
	return s.String()
}

func (s *FlinkApiProxyRequest) SetFlinkApiPath(v string) *FlinkApiProxyRequest {
	s.FlinkApiPath = &v
	return s
}

func (s *FlinkApiProxyRequest) SetNamespace(v string) *FlinkApiProxyRequest {
	s.Namespace = &v
	return s
}

func (s *FlinkApiProxyRequest) SetResourceId(v string) *FlinkApiProxyRequest {
	s.ResourceId = &v
	return s
}

func (s *FlinkApiProxyRequest) SetResourceType(v string) *FlinkApiProxyRequest {
	s.ResourceType = &v
	return s
}

type FlinkApiProxyResponseBody struct {
	Data         *string `json:"data,omitempty" xml:"data,omitempty"`
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32  `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s FlinkApiProxyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FlinkApiProxyResponseBody) GoString() string {
	return s.String()
}

func (s *FlinkApiProxyResponseBody) SetData(v string) *FlinkApiProxyResponseBody {
	s.Data = &v
	return s
}

func (s *FlinkApiProxyResponseBody) SetErrorCode(v string) *FlinkApiProxyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *FlinkApiProxyResponseBody) SetErrorMessage(v string) *FlinkApiProxyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *FlinkApiProxyResponseBody) SetHttpCode(v int32) *FlinkApiProxyResponseBody {
	s.HttpCode = &v
	return s
}

func (s *FlinkApiProxyResponseBody) SetRequestId(v string) *FlinkApiProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlinkApiProxyResponseBody) SetSuccess(v bool) *FlinkApiProxyResponseBody {
	s.Success = &v
	return s
}

type FlinkApiProxyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FlinkApiProxyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FlinkApiProxyResponse) String() string {
	return tea.Prettify(s)
}

func (s FlinkApiProxyResponse) GoString() string {
	return s.String()
}

func (s *FlinkApiProxyResponse) SetHeaders(v map[string]*string) *FlinkApiProxyResponse {
	s.Headers = v
	return s
}

func (s *FlinkApiProxyResponse) SetStatusCode(v int32) *FlinkApiProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *FlinkApiProxyResponse) SetBody(v *FlinkApiProxyResponseBody) *FlinkApiProxyResponse {
	s.Body = v
	return s
}

type GenerateResourcePlanWithFlinkConfAsyncHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GenerateResourcePlanWithFlinkConfAsyncHeaders) String() string {
	return tea.Prettify(s)
}

func (s GenerateResourcePlanWithFlinkConfAsyncHeaders) GoString() string {
	return s.String()
}

func (s *GenerateResourcePlanWithFlinkConfAsyncHeaders) SetCommonHeaders(v map[string]*string) *GenerateResourcePlanWithFlinkConfAsyncHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncHeaders) SetWorkspace(v string) *GenerateResourcePlanWithFlinkConfAsyncHeaders {
	s.Workspace = &v
	return s
}

type GenerateResourcePlanWithFlinkConfAsyncRequest struct {
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateResourcePlanWithFlinkConfAsyncRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateResourcePlanWithFlinkConfAsyncRequest) GoString() string {
	return s.String()
}

func (s *GenerateResourcePlanWithFlinkConfAsyncRequest) SetBody(v map[string]interface{}) *GenerateResourcePlanWithFlinkConfAsyncRequest {
	s.Body = v
	return s
}

type GenerateResourcePlanWithFlinkConfAsyncResponseBody struct {
	Data         *GenerateResourcePlanWithFlinkConfAsyncResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrorCode    *string                                                 `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                                 `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32                                                  `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId    *string                                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GenerateResourcePlanWithFlinkConfAsyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateResourcePlanWithFlinkConfAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBody) SetData(v *GenerateResourcePlanWithFlinkConfAsyncResponseBodyData) *GenerateResourcePlanWithFlinkConfAsyncResponseBody {
	s.Data = v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBody) SetErrorCode(v string) *GenerateResourcePlanWithFlinkConfAsyncResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBody) SetErrorMessage(v string) *GenerateResourcePlanWithFlinkConfAsyncResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBody) SetHttpCode(v int32) *GenerateResourcePlanWithFlinkConfAsyncResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBody) SetRequestId(v string) *GenerateResourcePlanWithFlinkConfAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBody) SetSuccess(v bool) *GenerateResourcePlanWithFlinkConfAsyncResponseBody {
	s.Success = &v
	return s
}

type GenerateResourcePlanWithFlinkConfAsyncResponseBodyData struct {
	TicketId *string `json:"ticketId,omitempty" xml:"ticketId,omitempty"`
}

func (s GenerateResourcePlanWithFlinkConfAsyncResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateResourcePlanWithFlinkConfAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponseBodyData) SetTicketId(v string) *GenerateResourcePlanWithFlinkConfAsyncResponseBodyData {
	s.TicketId = &v
	return s
}

type GenerateResourcePlanWithFlinkConfAsyncResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateResourcePlanWithFlinkConfAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateResourcePlanWithFlinkConfAsyncResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateResourcePlanWithFlinkConfAsyncResponse) GoString() string {
	return s.String()
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponse) SetHeaders(v map[string]*string) *GenerateResourcePlanWithFlinkConfAsyncResponse {
	s.Headers = v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponse) SetStatusCode(v int32) *GenerateResourcePlanWithFlinkConfAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponse) SetBody(v *GenerateResourcePlanWithFlinkConfAsyncResponseBody) *GenerateResourcePlanWithFlinkConfAsyncResponse {
	s.Body = v
	return s
}

type GetDeploymentHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetDeploymentHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentHeaders) GoString() string {
	return s.String()
}

func (s *GetDeploymentHeaders) SetCommonHeaders(v map[string]*string) *GetDeploymentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeploymentHeaders) SetWorkspace(v string) *GetDeploymentHeaders {
	s.Workspace = &v
	return s
}

type GetDeploymentResponseBody struct {
	Data         *Deployment `json:"data,omitempty" xml:"data,omitempty"`
	ErrorCode    *string     `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string     `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32      `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId    *string     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool       `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDeploymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeploymentResponseBody) SetData(v *Deployment) *GetDeploymentResponseBody {
	s.Data = v
	return s
}

func (s *GetDeploymentResponseBody) SetErrorCode(v string) *GetDeploymentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDeploymentResponseBody) SetErrorMessage(v string) *GetDeploymentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDeploymentResponseBody) SetHttpCode(v int32) *GetDeploymentResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDeploymentResponseBody) SetRequestId(v string) *GetDeploymentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeploymentResponseBody) SetSuccess(v bool) *GetDeploymentResponseBody {
	s.Success = &v
	return s
}

type GetDeploymentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentResponse) GoString() string {
	return s.String()
}

func (s *GetDeploymentResponse) SetHeaders(v map[string]*string) *GetDeploymentResponse {
	s.Headers = v
	return s
}

func (s *GetDeploymentResponse) SetStatusCode(v int32) *GetDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeploymentResponse) SetBody(v *GetDeploymentResponseBody) *GetDeploymentResponse {
	s.Body = v
	return s
}

type GetGenerateResourcePlanResultHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetGenerateResourcePlanResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetGenerateResourcePlanResultHeaders) GoString() string {
	return s.String()
}

func (s *GetGenerateResourcePlanResultHeaders) SetCommonHeaders(v map[string]*string) *GetGenerateResourcePlanResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetGenerateResourcePlanResultHeaders) SetWorkspace(v string) *GetGenerateResourcePlanResultHeaders {
	s.Workspace = &v
	return s
}

type GetGenerateResourcePlanResultResponseBody struct {
	Data         *AsyncResourcePlanOperationResult `json:"data,omitempty" xml:"data,omitempty"`
	ErrorCode    *string                           `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                           `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32                            `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId    *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetGenerateResourcePlanResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGenerateResourcePlanResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetGenerateResourcePlanResultResponseBody) SetData(v *AsyncResourcePlanOperationResult) *GetGenerateResourcePlanResultResponseBody {
	s.Data = v
	return s
}

func (s *GetGenerateResourcePlanResultResponseBody) SetErrorCode(v string) *GetGenerateResourcePlanResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetGenerateResourcePlanResultResponseBody) SetErrorMessage(v string) *GetGenerateResourcePlanResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetGenerateResourcePlanResultResponseBody) SetHttpCode(v int32) *GetGenerateResourcePlanResultResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetGenerateResourcePlanResultResponseBody) SetRequestId(v string) *GetGenerateResourcePlanResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGenerateResourcePlanResultResponseBody) SetSuccess(v bool) *GetGenerateResourcePlanResultResponseBody {
	s.Success = &v
	return s
}

type GetGenerateResourcePlanResultResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGenerateResourcePlanResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGenerateResourcePlanResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGenerateResourcePlanResultResponse) GoString() string {
	return s.String()
}

func (s *GetGenerateResourcePlanResultResponse) SetHeaders(v map[string]*string) *GetGenerateResourcePlanResultResponse {
	s.Headers = v
	return s
}

func (s *GetGenerateResourcePlanResultResponse) SetStatusCode(v int32) *GetGenerateResourcePlanResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGenerateResourcePlanResultResponse) SetBody(v *GetGenerateResourcePlanResultResponseBody) *GetGenerateResourcePlanResultResponse {
	s.Body = v
	return s
}

type GetJobHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetJobHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetJobHeaders) GoString() string {
	return s.String()
}

func (s *GetJobHeaders) SetCommonHeaders(v map[string]*string) *GetJobHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetJobHeaders) SetWorkspace(v string) *GetJobHeaders {
	s.Workspace = &v
	return s
}

type GetJobResponseBody struct {
	Data         *Job    `json:"data,omitempty" xml:"data,omitempty"`
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32  `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) SetData(v *Job) *GetJobResponseBody {
	s.Data = v
	return s
}

func (s *GetJobResponseBody) SetErrorCode(v string) *GetJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetJobResponseBody) SetErrorMessage(v string) *GetJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetJobResponseBody) SetHttpCode(v int32) *GetJobResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetJobResponseBody) SetRequestId(v string) *GetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobResponseBody) SetSuccess(v bool) *GetJobResponseBody {
	s.Success = &v
	return s
}

type GetJobResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponse) GoString() string {
	return s.String()
}

func (s *GetJobResponse) SetHeaders(v map[string]*string) *GetJobResponse {
	s.Headers = v
	return s
}

func (s *GetJobResponse) SetStatusCode(v int32) *GetJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobResponse) SetBody(v *GetJobResponseBody) *GetJobResponse {
	s.Body = v
	return s
}

type GetSavepointHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetSavepointHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSavepointHeaders) GoString() string {
	return s.String()
}

func (s *GetSavepointHeaders) SetCommonHeaders(v map[string]*string) *GetSavepointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSavepointHeaders) SetWorkspace(v string) *GetSavepointHeaders {
	s.Workspace = &v
	return s
}

type GetSavepointResponseBody struct {
	Data         *Savepoint `json:"data,omitempty" xml:"data,omitempty"`
	ErrorCode    *string    `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string    `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32     `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId    *string    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSavepointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSavepointResponseBody) GoString() string {
	return s.String()
}

func (s *GetSavepointResponseBody) SetData(v *Savepoint) *GetSavepointResponseBody {
	s.Data = v
	return s
}

func (s *GetSavepointResponseBody) SetErrorCode(v string) *GetSavepointResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSavepointResponseBody) SetErrorMessage(v string) *GetSavepointResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetSavepointResponseBody) SetHttpCode(v int32) *GetSavepointResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetSavepointResponseBody) SetRequestId(v string) *GetSavepointResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSavepointResponseBody) SetSuccess(v bool) *GetSavepointResponseBody {
	s.Success = &v
	return s
}

type GetSavepointResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSavepointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSavepointResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSavepointResponse) GoString() string {
	return s.String()
}

func (s *GetSavepointResponse) SetHeaders(v map[string]*string) *GetSavepointResponse {
	s.Headers = v
	return s
}

func (s *GetSavepointResponse) SetStatusCode(v int32) *GetSavepointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSavepointResponse) SetBody(v *GetSavepointResponseBody) *GetSavepointResponse {
	s.Body = v
	return s
}

type ListDeploymentTargetsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListDeploymentTargetsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentTargetsHeaders) GoString() string {
	return s.String()
}

func (s *ListDeploymentTargetsHeaders) SetCommonHeaders(v map[string]*string) *ListDeploymentTargetsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDeploymentTargetsHeaders) SetWorkspace(v string) *ListDeploymentTargetsHeaders {
	s.Workspace = &v
	return s
}

type ListDeploymentTargetsRequest struct {
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	PageSize  *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListDeploymentTargetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentTargetsRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentTargetsRequest) SetPageIndex(v int32) *ListDeploymentTargetsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDeploymentTargetsRequest) SetPageSize(v int32) *ListDeploymentTargetsRequest {
	s.PageSize = &v
	return s
}

type ListDeploymentTargetsResponseBody struct {
	Data         []*DeploymentTarget `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrorCode    *string             `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string             `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32              `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	PageIndex    *int32              `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	PageSize     *int32              `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RequestId    *string             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool               `json:"success,omitempty" xml:"success,omitempty"`
	TotalSize    *int32              `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListDeploymentTargetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeploymentTargetsResponseBody) SetData(v []*DeploymentTarget) *ListDeploymentTargetsResponseBody {
	s.Data = v
	return s
}

func (s *ListDeploymentTargetsResponseBody) SetErrorCode(v string) *ListDeploymentTargetsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDeploymentTargetsResponseBody) SetErrorMessage(v string) *ListDeploymentTargetsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDeploymentTargetsResponseBody) SetHttpCode(v int32) *ListDeploymentTargetsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListDeploymentTargetsResponseBody) SetPageIndex(v int32) *ListDeploymentTargetsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListDeploymentTargetsResponseBody) SetPageSize(v int32) *ListDeploymentTargetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentTargetsResponseBody) SetRequestId(v string) *ListDeploymentTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeploymentTargetsResponseBody) SetSuccess(v bool) *ListDeploymentTargetsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDeploymentTargetsResponseBody) SetTotalSize(v int32) *ListDeploymentTargetsResponseBody {
	s.TotalSize = &v
	return s
}

type ListDeploymentTargetsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeploymentTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeploymentTargetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentTargetsResponse) GoString() string {
	return s.String()
}

func (s *ListDeploymentTargetsResponse) SetHeaders(v map[string]*string) *ListDeploymentTargetsResponse {
	s.Headers = v
	return s
}

func (s *ListDeploymentTargetsResponse) SetStatusCode(v int32) *ListDeploymentTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeploymentTargetsResponse) SetBody(v *ListDeploymentTargetsResponseBody) *ListDeploymentTargetsResponse {
	s.Body = v
	return s
}

type ListDeploymentsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListDeploymentsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsHeaders) GoString() string {
	return s.String()
}

func (s *ListDeploymentsHeaders) SetCommonHeaders(v map[string]*string) *ListDeploymentsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDeploymentsHeaders) SetWorkspace(v string) *ListDeploymentsHeaders {
	s.Workspace = &v
	return s
}

type ListDeploymentsRequest struct {
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	PageSize  *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListDeploymentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentsRequest) SetPageIndex(v int32) *ListDeploymentsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDeploymentsRequest) SetPageSize(v int32) *ListDeploymentsRequest {
	s.PageSize = &v
	return s
}

type ListDeploymentsResponseBody struct {
	Data         []*Deployment `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrorCode    *string       `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string       `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32        `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	PageIndex    *int32        `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	PageSize     *int32        `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RequestId    *string       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool         `json:"success,omitempty" xml:"success,omitempty"`
	TotalSize    *int32        `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListDeploymentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponseBody) SetData(v []*Deployment) *ListDeploymentsResponseBody {
	s.Data = v
	return s
}

func (s *ListDeploymentsResponseBody) SetErrorCode(v string) *ListDeploymentsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDeploymentsResponseBody) SetErrorMessage(v string) *ListDeploymentsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDeploymentsResponseBody) SetHttpCode(v int32) *ListDeploymentsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListDeploymentsResponseBody) SetPageIndex(v int32) *ListDeploymentsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListDeploymentsResponseBody) SetPageSize(v int32) *ListDeploymentsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentsResponseBody) SetRequestId(v string) *ListDeploymentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeploymentsResponseBody) SetSuccess(v bool) *ListDeploymentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDeploymentsResponseBody) SetTotalSize(v int32) *ListDeploymentsResponseBody {
	s.TotalSize = &v
	return s
}

type ListDeploymentsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeploymentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeploymentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsResponse) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponse) SetHeaders(v map[string]*string) *ListDeploymentsResponse {
	s.Headers = v
	return s
}

func (s *ListDeploymentsResponse) SetStatusCode(v int32) *ListDeploymentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeploymentsResponse) SetBody(v *ListDeploymentsResponseBody) *ListDeploymentsResponse {
	s.Body = v
	return s
}

type ListEngineVersionMetadataHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListEngineVersionMetadataHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListEngineVersionMetadataHeaders) GoString() string {
	return s.String()
}

func (s *ListEngineVersionMetadataHeaders) SetCommonHeaders(v map[string]*string) *ListEngineVersionMetadataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListEngineVersionMetadataHeaders) SetWorkspace(v string) *ListEngineVersionMetadataHeaders {
	s.Workspace = &v
	return s
}

type ListEngineVersionMetadataResponseBody struct {
	Data         *EngineVersionMetadataIndex `json:"data,omitempty" xml:"data,omitempty"`
	ErrorCode    *string                     `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                     `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32                      `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId    *string                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                       `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListEngineVersionMetadataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEngineVersionMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *ListEngineVersionMetadataResponseBody) SetData(v *EngineVersionMetadataIndex) *ListEngineVersionMetadataResponseBody {
	s.Data = v
	return s
}

func (s *ListEngineVersionMetadataResponseBody) SetErrorCode(v string) *ListEngineVersionMetadataResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListEngineVersionMetadataResponseBody) SetErrorMessage(v string) *ListEngineVersionMetadataResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListEngineVersionMetadataResponseBody) SetHttpCode(v int32) *ListEngineVersionMetadataResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListEngineVersionMetadataResponseBody) SetRequestId(v string) *ListEngineVersionMetadataResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEngineVersionMetadataResponseBody) SetSuccess(v bool) *ListEngineVersionMetadataResponseBody {
	s.Success = &v
	return s
}

type ListEngineVersionMetadataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEngineVersionMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEngineVersionMetadataResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEngineVersionMetadataResponse) GoString() string {
	return s.String()
}

func (s *ListEngineVersionMetadataResponse) SetHeaders(v map[string]*string) *ListEngineVersionMetadataResponse {
	s.Headers = v
	return s
}

func (s *ListEngineVersionMetadataResponse) SetStatusCode(v int32) *ListEngineVersionMetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEngineVersionMetadataResponse) SetBody(v *ListEngineVersionMetadataResponseBody) *ListEngineVersionMetadataResponse {
	s.Body = v
	return s
}

type ListJobsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListJobsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListJobsHeaders) GoString() string {
	return s.String()
}

func (s *ListJobsHeaders) SetCommonHeaders(v map[string]*string) *ListJobsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListJobsHeaders) SetWorkspace(v string) *ListJobsHeaders {
	s.Workspace = &v
	return s
}

type ListJobsRequest struct {
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	PageIndex    *int32  `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	PageSize     *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) SetDeploymentId(v string) *ListJobsRequest {
	s.DeploymentId = &v
	return s
}

func (s *ListJobsRequest) SetPageIndex(v int32) *ListJobsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v int32) *ListJobsRequest {
	s.PageSize = &v
	return s
}

type ListJobsResponseBody struct {
	Data         []*Job  `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32  `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	PageIndex    *int32  `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	PageSize     *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TotalSize    *int32  `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) SetData(v []*Job) *ListJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListJobsResponseBody) SetErrorCode(v string) *ListJobsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListJobsResponseBody) SetErrorMessage(v string) *ListJobsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListJobsResponseBody) SetHttpCode(v int32) *ListJobsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListJobsResponseBody) SetPageIndex(v int32) *ListJobsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListJobsResponseBody) SetPageSize(v int32) *ListJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetSuccess(v bool) *ListJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ListJobsResponseBody) SetTotalSize(v int32) *ListJobsResponseBody {
	s.TotalSize = &v
	return s
}

type ListJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponse) GoString() string {
	return s.String()
}

func (s *ListJobsResponse) SetHeaders(v map[string]*string) *ListJobsResponse {
	s.Headers = v
	return s
}

func (s *ListJobsResponse) SetStatusCode(v int32) *ListJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobsResponse) SetBody(v *ListJobsResponseBody) *ListJobsResponse {
	s.Body = v
	return s
}

type ListSavepointsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListSavepointsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSavepointsHeaders) GoString() string {
	return s.String()
}

func (s *ListSavepointsHeaders) SetCommonHeaders(v map[string]*string) *ListSavepointsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSavepointsHeaders) SetWorkspace(v string) *ListSavepointsHeaders {
	s.Workspace = &v
	return s
}

type ListSavepointsRequest struct {
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	JobId        *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	PageIndex    *int32  `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	PageSize     *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListSavepointsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSavepointsRequest) GoString() string {
	return s.String()
}

func (s *ListSavepointsRequest) SetDeploymentId(v string) *ListSavepointsRequest {
	s.DeploymentId = &v
	return s
}

func (s *ListSavepointsRequest) SetJobId(v string) *ListSavepointsRequest {
	s.JobId = &v
	return s
}

func (s *ListSavepointsRequest) SetPageIndex(v int32) *ListSavepointsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListSavepointsRequest) SetPageSize(v int32) *ListSavepointsRequest {
	s.PageSize = &v
	return s
}

type ListSavepointsResponseBody struct {
	Data         []*Savepoint `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrorCode    *string      `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string      `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32       `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	PageIndex    *int32       `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	PageSize     *int32       `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RequestId    *string      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool        `json:"success,omitempty" xml:"success,omitempty"`
	TotalSize    *int32       `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListSavepointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSavepointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSavepointsResponseBody) SetData(v []*Savepoint) *ListSavepointsResponseBody {
	s.Data = v
	return s
}

func (s *ListSavepointsResponseBody) SetErrorCode(v string) *ListSavepointsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSavepointsResponseBody) SetErrorMessage(v string) *ListSavepointsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSavepointsResponseBody) SetHttpCode(v int32) *ListSavepointsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListSavepointsResponseBody) SetPageIndex(v int32) *ListSavepointsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListSavepointsResponseBody) SetPageSize(v int32) *ListSavepointsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSavepointsResponseBody) SetRequestId(v string) *ListSavepointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSavepointsResponseBody) SetSuccess(v bool) *ListSavepointsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSavepointsResponseBody) SetTotalSize(v int32) *ListSavepointsResponseBody {
	s.TotalSize = &v
	return s
}

type ListSavepointsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSavepointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSavepointsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSavepointsResponse) GoString() string {
	return s.String()
}

func (s *ListSavepointsResponse) SetHeaders(v map[string]*string) *ListSavepointsResponse {
	s.Headers = v
	return s
}

func (s *ListSavepointsResponse) SetStatusCode(v int32) *ListSavepointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSavepointsResponse) SetBody(v *ListSavepointsResponseBody) *ListSavepointsResponse {
	s.Body = v
	return s
}

type ListVariablesHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListVariablesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListVariablesHeaders) GoString() string {
	return s.String()
}

func (s *ListVariablesHeaders) SetCommonHeaders(v map[string]*string) *ListVariablesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListVariablesHeaders) SetWorkspace(v string) *ListVariablesHeaders {
	s.Workspace = &v
	return s
}

type ListVariablesRequest struct {
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	PageSize  *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListVariablesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVariablesRequest) GoString() string {
	return s.String()
}

func (s *ListVariablesRequest) SetPageIndex(v int32) *ListVariablesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListVariablesRequest) SetPageSize(v int32) *ListVariablesRequest {
	s.PageSize = &v
	return s
}

type ListVariablesResponseBody struct {
	Data         []*Variable `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrorCode    *string     `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string     `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32      `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	PageIndex    *int32      `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	PageSize     *int32      `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RequestId    *string     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool       `json:"success,omitempty" xml:"success,omitempty"`
	TotalSize    *int32      `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListVariablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVariablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVariablesResponseBody) SetData(v []*Variable) *ListVariablesResponseBody {
	s.Data = v
	return s
}

func (s *ListVariablesResponseBody) SetErrorCode(v string) *ListVariablesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListVariablesResponseBody) SetErrorMessage(v string) *ListVariablesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListVariablesResponseBody) SetHttpCode(v int32) *ListVariablesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListVariablesResponseBody) SetPageIndex(v int32) *ListVariablesResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListVariablesResponseBody) SetPageSize(v int32) *ListVariablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVariablesResponseBody) SetRequestId(v string) *ListVariablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVariablesResponseBody) SetSuccess(v bool) *ListVariablesResponseBody {
	s.Success = &v
	return s
}

func (s *ListVariablesResponseBody) SetTotalSize(v int32) *ListVariablesResponseBody {
	s.TotalSize = &v
	return s
}

type ListVariablesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVariablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVariablesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVariablesResponse) GoString() string {
	return s.String()
}

func (s *ListVariablesResponse) SetHeaders(v map[string]*string) *ListVariablesResponse {
	s.Headers = v
	return s
}

func (s *ListVariablesResponse) SetStatusCode(v int32) *ListVariablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVariablesResponse) SetBody(v *ListVariablesResponseBody) *ListVariablesResponse {
	s.Body = v
	return s
}

type StartJobHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s StartJobHeaders) String() string {
	return tea.Prettify(s)
}

func (s StartJobHeaders) GoString() string {
	return s.String()
}

func (s *StartJobHeaders) SetCommonHeaders(v map[string]*string) *StartJobHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartJobHeaders) SetWorkspace(v string) *StartJobHeaders {
	s.Workspace = &v
	return s
}

type StartJobRequest struct {
	Body *StartJobRequestBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StartJobRequest) GoString() string {
	return s.String()
}

func (s *StartJobRequest) SetBody(v *StartJobRequestBody) *StartJobRequest {
	s.Body = v
	return s
}

type StartJobResponseBody struct {
	Data         *Job    `json:"data,omitempty" xml:"data,omitempty"`
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32  `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StartJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartJobResponseBody) SetData(v *Job) *StartJobResponseBody {
	s.Data = v
	return s
}

func (s *StartJobResponseBody) SetErrorCode(v string) *StartJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StartJobResponseBody) SetErrorMessage(v string) *StartJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StartJobResponseBody) SetHttpCode(v int32) *StartJobResponseBody {
	s.HttpCode = &v
	return s
}

func (s *StartJobResponseBody) SetRequestId(v string) *StartJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartJobResponseBody) SetSuccess(v bool) *StartJobResponseBody {
	s.Success = &v
	return s
}

type StartJobResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StartJobResponse) GoString() string {
	return s.String()
}

func (s *StartJobResponse) SetHeaders(v map[string]*string) *StartJobResponse {
	s.Headers = v
	return s
}

func (s *StartJobResponse) SetStatusCode(v int32) *StartJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StartJobResponse) SetBody(v *StartJobResponseBody) *StartJobResponse {
	s.Body = v
	return s
}

type StopJobHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s StopJobHeaders) String() string {
	return tea.Prettify(s)
}

func (s StopJobHeaders) GoString() string {
	return s.String()
}

func (s *StopJobHeaders) SetCommonHeaders(v map[string]*string) *StopJobHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StopJobHeaders) SetWorkspace(v string) *StopJobHeaders {
	s.Workspace = &v
	return s
}

type StopJobRequest struct {
	Body *StopJobRequestBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StopJobRequest) GoString() string {
	return s.String()
}

func (s *StopJobRequest) SetBody(v *StopJobRequestBody) *StopJobRequest {
	s.Body = v
	return s
}

type StopJobResponseBody struct {
	Data         *Job    `json:"data,omitempty" xml:"data,omitempty"`
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32  `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StopJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopJobResponseBody) SetData(v *Job) *StopJobResponseBody {
	s.Data = v
	return s
}

func (s *StopJobResponseBody) SetErrorCode(v string) *StopJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StopJobResponseBody) SetErrorMessage(v string) *StopJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopJobResponseBody) SetHttpCode(v int32) *StopJobResponseBody {
	s.HttpCode = &v
	return s
}

func (s *StopJobResponseBody) SetRequestId(v string) *StopJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopJobResponseBody) SetSuccess(v bool) *StopJobResponseBody {
	s.Success = &v
	return s
}

type StopJobResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StopJobResponse) GoString() string {
	return s.String()
}

func (s *StopJobResponse) SetHeaders(v map[string]*string) *StopJobResponse {
	s.Headers = v
	return s
}

func (s *StopJobResponse) SetStatusCode(v int32) *StopJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopJobResponse) SetBody(v *StopJobResponseBody) *StopJobResponse {
	s.Body = v
	return s
}

type UpdateDeploymentHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Workspace     *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateDeploymentHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeploymentHeaders) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentHeaders) SetCommonHeaders(v map[string]*string) *UpdateDeploymentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateDeploymentHeaders) SetWorkspace(v string) *UpdateDeploymentHeaders {
	s.Workspace = &v
	return s
}

type UpdateDeploymentRequest struct {
	Body *Deployment `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeploymentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeploymentRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentRequest) SetBody(v *Deployment) *UpdateDeploymentRequest {
	s.Body = v
	return s
}

type UpdateDeploymentResponseBody struct {
	Data         *Deployment `json:"data,omitempty" xml:"data,omitempty"`
	ErrorCode    *string     `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string     `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpCode     *int32      `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId    *string     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool       `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateDeploymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentResponseBody) SetData(v *Deployment) *UpdateDeploymentResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDeploymentResponseBody) SetErrorCode(v string) *UpdateDeploymentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDeploymentResponseBody) SetErrorMessage(v string) *UpdateDeploymentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDeploymentResponseBody) SetHttpCode(v int32) *UpdateDeploymentResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateDeploymentResponseBody) SetRequestId(v string) *UpdateDeploymentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeploymentResponseBody) SetSuccess(v bool) *UpdateDeploymentResponseBody {
	s.Success = &v
	return s
}

type UpdateDeploymentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeploymentResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentResponse) SetHeaders(v map[string]*string) *UpdateDeploymentResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeploymentResponse) SetStatusCode(v int32) *UpdateDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeploymentResponse) SetBody(v *UpdateDeploymentResponseBody) *UpdateDeploymentResponse {
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
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ververica"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateDeployment(namespace *string, request *CreateDeploymentRequest) (_result *CreateDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateDeploymentHeaders{}
	_result = &CreateDeploymentResponse{}
	_body, _err := client.CreateDeploymentWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDeploymentWithOptions(namespace *string, request *CreateDeploymentRequest, headers *CreateDeploymentHeaders, runtime *util.RuntimeOptions) (_result *CreateDeploymentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDeployment"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployments"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSavepoint(namespace *string, request *CreateSavepointRequest) (_result *CreateSavepointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateSavepointHeaders{}
	_result = &CreateSavepointResponse{}
	_body, _err := client.CreateSavepointWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSavepointWithOptions(namespace *string, request *CreateSavepointRequest, headers *CreateSavepointHeaders, runtime *util.RuntimeOptions) (_result *CreateSavepointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeploymentId)) {
		body["deploymentId"] = request.DeploymentId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.NativeFormat)) {
		body["nativeFormat"] = request.NativeFormat
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSavepoint"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/savepoints"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSavepointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVariable(namespace *string, request *CreateVariableRequest) (_result *CreateVariableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateVariableHeaders{}
	_result = &CreateVariableResponse{}
	_body, _err := client.CreateVariableWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVariableWithOptions(namespace *string, request *CreateVariableRequest, headers *CreateVariableHeaders, runtime *util.RuntimeOptions) (_result *CreateVariableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVariable"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/variables"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDeployment(namespace *string, deploymentId *string) (_result *DeleteDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteDeploymentHeaders{}
	_result = &DeleteDeploymentResponse{}
	_body, _err := client.DeleteDeploymentWithOptions(namespace, deploymentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDeploymentWithOptions(namespace *string, deploymentId *string, headers *DeleteDeploymentHeaders, runtime *util.RuntimeOptions) (_result *DeleteDeploymentResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDeployment"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployments/" + tea.StringValue(openapiutil.GetEncodeParam(deploymentId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteJob(namespace *string, jobId *string) (_result *DeleteJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteJobHeaders{}
	_result = &DeleteJobResponse{}
	_body, _err := client.DeleteJobWithOptions(namespace, jobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteJobWithOptions(namespace *string, jobId *string, headers *DeleteJobHeaders, runtime *util.RuntimeOptions) (_result *DeleteJobResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteJob"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(jobId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSavepoint(namespace *string, savepointId *string) (_result *DeleteSavepointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteSavepointHeaders{}
	_result = &DeleteSavepointResponse{}
	_body, _err := client.DeleteSavepointWithOptions(namespace, savepointId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSavepointWithOptions(namespace *string, savepointId *string, headers *DeleteSavepointHeaders, runtime *util.RuntimeOptions) (_result *DeleteSavepointResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSavepoint"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/savepoints/" + tea.StringValue(openapiutil.GetEncodeParam(savepointId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSavepointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVariable(namespace *string, name *string) (_result *DeleteVariableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteVariableHeaders{}
	_result = &DeleteVariableResponse{}
	_body, _err := client.DeleteVariableWithOptions(namespace, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVariableWithOptions(namespace *string, name *string, headers *DeleteVariableHeaders, runtime *util.RuntimeOptions) (_result *DeleteVariableResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVariable"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/variables/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FlinkApiProxy(request *FlinkApiProxyRequest) (_result *FlinkApiProxyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FlinkApiProxyHeaders{}
	_result = &FlinkApiProxyResponse{}
	_body, _err := client.FlinkApiProxyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FlinkApiProxyWithOptions(request *FlinkApiProxyRequest, headers *FlinkApiProxyHeaders, runtime *util.RuntimeOptions) (_result *FlinkApiProxyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FlinkApiPath)) {
		query["flinkApiPath"] = request.FlinkApiPath
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["resourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["resourceType"] = request.ResourceType
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FlinkApiProxy"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/flink-ui/v2/proxy"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FlinkApiProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateResourcePlanWithFlinkConfAsync(namespace *string, deploymentId *string, request *GenerateResourcePlanWithFlinkConfAsyncRequest) (_result *GenerateResourcePlanWithFlinkConfAsyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GenerateResourcePlanWithFlinkConfAsyncHeaders{}
	_result = &GenerateResourcePlanWithFlinkConfAsyncResponse{}
	_body, _err := client.GenerateResourcePlanWithFlinkConfAsyncWithOptions(namespace, deploymentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateResourcePlanWithFlinkConfAsyncWithOptions(namespace *string, deploymentId *string, request *GenerateResourcePlanWithFlinkConfAsyncRequest, headers *GenerateResourcePlanWithFlinkConfAsyncHeaders, runtime *util.RuntimeOptions) (_result *GenerateResourcePlanWithFlinkConfAsyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateResourcePlanWithFlinkConfAsync"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployments/" + tea.StringValue(openapiutil.GetEncodeParam(deploymentId)) + "/resource-plan%3AasyncGenerate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateResourcePlanWithFlinkConfAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeployment(namespace *string, deploymentId *string) (_result *GetDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDeploymentHeaders{}
	_result = &GetDeploymentResponse{}
	_body, _err := client.GetDeploymentWithOptions(namespace, deploymentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeploymentWithOptions(namespace *string, deploymentId *string, headers *GetDeploymentHeaders, runtime *util.RuntimeOptions) (_result *GetDeploymentResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeployment"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployments/" + tea.StringValue(openapiutil.GetEncodeParam(deploymentId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGenerateResourcePlanResult(namespace *string, ticketId *string) (_result *GetGenerateResourcePlanResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetGenerateResourcePlanResultHeaders{}
	_result = &GetGenerateResourcePlanResultResponse{}
	_body, _err := client.GetGenerateResourcePlanResultWithOptions(namespace, ticketId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGenerateResourcePlanResultWithOptions(namespace *string, ticketId *string, headers *GetGenerateResourcePlanResultHeaders, runtime *util.RuntimeOptions) (_result *GetGenerateResourcePlanResultResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetGenerateResourcePlanResult"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployments/tickets/" + tea.StringValue(openapiutil.GetEncodeParam(ticketId)) + "/resource-plan%3AasyncGenerate"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGenerateResourcePlanResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJob(namespace *string, jobId *string) (_result *GetJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetJobHeaders{}
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(namespace, jobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobWithOptions(namespace *string, jobId *string, headers *GetJobHeaders, runtime *util.RuntimeOptions) (_result *GetJobResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetJob"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(jobId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSavepoint(namespace *string, savepointId *string) (_result *GetSavepointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSavepointHeaders{}
	_result = &GetSavepointResponse{}
	_body, _err := client.GetSavepointWithOptions(namespace, savepointId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSavepointWithOptions(namespace *string, savepointId *string, headers *GetSavepointHeaders, runtime *util.RuntimeOptions) (_result *GetSavepointResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSavepoint"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/savepoints/" + tea.StringValue(openapiutil.GetEncodeParam(savepointId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSavepointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeploymentTargets(namespace *string, request *ListDeploymentTargetsRequest) (_result *ListDeploymentTargetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListDeploymentTargetsHeaders{}
	_result = &ListDeploymentTargetsResponse{}
	_body, _err := client.ListDeploymentTargetsWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeploymentTargetsWithOptions(namespace *string, request *ListDeploymentTargetsRequest, headers *ListDeploymentTargetsHeaders, runtime *util.RuntimeOptions) (_result *ListDeploymentTargetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["pageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeploymentTargets"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployment-targets"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeploymentTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeployments(namespace *string, request *ListDeploymentsRequest) (_result *ListDeploymentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListDeploymentsHeaders{}
	_result = &ListDeploymentsResponse{}
	_body, _err := client.ListDeploymentsWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeploymentsWithOptions(namespace *string, request *ListDeploymentsRequest, headers *ListDeploymentsHeaders, runtime *util.RuntimeOptions) (_result *ListDeploymentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["pageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeployments"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeploymentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEngineVersionMetadata() (_result *ListEngineVersionMetadataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListEngineVersionMetadataHeaders{}
	_result = &ListEngineVersionMetadataResponse{}
	_body, _err := client.ListEngineVersionMetadataWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEngineVersionMetadataWithOptions(headers *ListEngineVersionMetadataHeaders, runtime *util.RuntimeOptions) (_result *ListEngineVersionMetadataResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("ListEngineVersionMetadata"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/engine-version-meta.json"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEngineVersionMetadataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobs(namespace *string, request *ListJobsRequest) (_result *ListJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListJobsHeaders{}
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobsWithOptions(namespace *string, request *ListJobsRequest, headers *ListJobsHeaders, runtime *util.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeploymentId)) {
		query["deploymentId"] = request.DeploymentId
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["pageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobs"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/jobs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSavepoints(namespace *string, request *ListSavepointsRequest) (_result *ListSavepointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSavepointsHeaders{}
	_result = &ListSavepointsResponse{}
	_body, _err := client.ListSavepointsWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSavepointsWithOptions(namespace *string, request *ListSavepointsRequest, headers *ListSavepointsHeaders, runtime *util.RuntimeOptions) (_result *ListSavepointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeploymentId)) {
		query["deploymentId"] = request.DeploymentId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["jobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["pageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSavepoints"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/savepoints"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSavepointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVariables(namespace *string, request *ListVariablesRequest) (_result *ListVariablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListVariablesHeaders{}
	_result = &ListVariablesResponse{}
	_body, _err := client.ListVariablesWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVariablesWithOptions(namespace *string, request *ListVariablesRequest, headers *ListVariablesHeaders, runtime *util.RuntimeOptions) (_result *ListVariablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["pageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVariables"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/variables"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVariablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartJob(namespace *string, request *StartJobRequest) (_result *StartJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StartJobHeaders{}
	_result = &StartJobResponse{}
	_body, _err := client.StartJobWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartJobWithOptions(namespace *string, request *StartJobRequest, headers *StartJobHeaders, runtime *util.RuntimeOptions) (_result *StartJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartJob"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/jobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopJob(namespace *string, jobId *string, request *StopJobRequest) (_result *StopJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StopJobHeaders{}
	_result = &StopJobResponse{}
	_body, _err := client.StopJobWithOptions(namespace, jobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopJobWithOptions(namespace *string, jobId *string, request *StopJobRequest, headers *StopJobHeaders, runtime *util.RuntimeOptions) (_result *StopJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopJob"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(jobId)) + "%3Astop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDeployment(namespace *string, deploymentId *string, request *UpdateDeploymentRequest) (_result *UpdateDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateDeploymentHeaders{}
	_result = &UpdateDeploymentResponse{}
	_body, _err := client.UpdateDeploymentWithOptions(namespace, deploymentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDeploymentWithOptions(namespace *string, deploymentId *string, request *UpdateDeploymentRequest, headers *UpdateDeploymentHeaders, runtime *util.RuntimeOptions) (_result *UpdateDeploymentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Workspace)) {
		realHeaders["workspace"] = util.ToJSONString(headers.Workspace)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDeployment"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployments/" + tea.StringValue(openapiutil.GetEncodeParam(deploymentId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
