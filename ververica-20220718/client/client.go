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
	JarArtifact *JarArtifact `json:"jarArtifact,omitempty" xml:"jarArtifact,omitempty"`
	// example:
	//
	// SQLSCRIPT
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

type AsyncDraftDeployResult struct {
	ArtifactValidationDetail *ValidateStatementResult `json:"artifactValidationDetail,omitempty" xml:"artifactValidationDetail,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// example:
	//
	// "Validation error: SQL validate failed"
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// FINISHED
	TicketStatus *string `json:"ticketStatus,omitempty" xml:"ticketStatus,omitempty"`
}

func (s AsyncDraftDeployResult) String() string {
	return tea.Prettify(s)
}

func (s AsyncDraftDeployResult) GoString() string {
	return s.String()
}

func (s *AsyncDraftDeployResult) SetArtifactValidationDetail(v *ValidateStatementResult) *AsyncDraftDeployResult {
	s.ArtifactValidationDetail = v
	return s
}

func (s *AsyncDraftDeployResult) SetDeploymentId(v string) *AsyncDraftDeployResult {
	s.DeploymentId = &v
	return s
}

func (s *AsyncDraftDeployResult) SetMessage(v string) *AsyncDraftDeployResult {
	s.Message = &v
	return s
}

func (s *AsyncDraftDeployResult) SetSuccess(v bool) *AsyncDraftDeployResult {
	s.Success = &v
	return s
}

func (s *AsyncDraftDeployResult) SetTicketStatus(v string) *AsyncDraftDeployResult {
	s.TicketStatus = &v
	return s
}

type AsyncResourcePlanOperationResult struct {
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// {\"ssgProfiles\":[{\"name\":\"default\",\"cpu\":1.13,\"heap\":\"1 gb\",\"offHeap\":\"32 mb\",\"managed\":{},\"extended\":{}}],\"nodes\":[{\"id\":1,\"type\":\"StreamExecTableSourceScan\",\"desc\":\"Source: datagen_source[78]\",\"profile\":{\"group\":\"default\",\"parallelism\":1,\"maxParallelism\":32768,\"minParallelism\":1}},{\"id\":2,\"type\":\"StreamExecSink\",\"desc\":\"Sink: blackhole_sink[79]\",\"profile\":{\"group\":\"default\",\"parallelism\":1,\"maxParallelism\":32768,\"minParallelism\":1}}],\"edges\":[{\"source\":1,\"target\":2,\"mode\":\"PIPELINED\",\"strategy\":\"FORWARD\"}],\"vertices\":{\"717c7b8afebbfb7137f6f0f99beb2a94\":[1,2]}}
	Plan *string `json:"plan,omitempty" xml:"plan,omitempty"`
	// example:
	//
	// FINISHED
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
	JobmanagerResourceSettingSpec *BasicResourceSettingSpec `json:"jobmanagerResourceSettingSpec,omitempty" xml:"jobmanagerResourceSettingSpec,omitempty"`
	// example:
	//
	// 4
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
	// example:
	//
	// 2.0
	Cpu *float64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// 4Gi
	Memory *string `json:"memory,omitempty" xml:"memory,omitempty"`
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
	// example:
	//
	// 10
	MaxSlot *int64 `json:"maxSlot,omitempty" xml:"maxSlot,omitempty"`
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

type Catalog struct {
	Name       *string                `json:"name,omitempty" xml:"name,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
}

func (s Catalog) String() string {
	return tea.Prettify(s)
}

func (s Catalog) GoString() string {
	return s.String()
}

func (s *Catalog) SetName(v string) *Catalog {
	s.Name = &v
	return s
}

func (s *Catalog) SetProperties(v map[string]interface{}) *Catalog {
	s.Properties = v
	return s
}

type Connector struct {
	Creator          *string     `json:"creator,omitempty" xml:"creator,omitempty"`
	CreatorName      *string     `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	Dependencies     []*string   `json:"dependencies,omitempty" xml:"dependencies,omitempty" type:"Repeated"`
	Lookup           *bool       `json:"lookup,omitempty" xml:"lookup,omitempty"`
	Modifier         *string     `json:"modifier,omitempty" xml:"modifier,omitempty"`
	ModifierName     *string     `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	Name             *string     `json:"name,omitempty" xml:"name,omitempty"`
	Properties       []*Property `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
	Sink             *bool       `json:"sink,omitempty" xml:"sink,omitempty"`
	Source           *bool       `json:"source,omitempty" xml:"source,omitempty"`
	SupportedFormats []*string   `json:"supportedFormats,omitempty" xml:"supportedFormats,omitempty" type:"Repeated"`
	Type             *string     `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Connector) String() string {
	return tea.Prettify(s)
}

func (s Connector) GoString() string {
	return s.String()
}

func (s *Connector) SetCreator(v string) *Connector {
	s.Creator = &v
	return s
}

func (s *Connector) SetCreatorName(v string) *Connector {
	s.CreatorName = &v
	return s
}

func (s *Connector) SetDependencies(v []*string) *Connector {
	s.Dependencies = v
	return s
}

func (s *Connector) SetLookup(v bool) *Connector {
	s.Lookup = &v
	return s
}

func (s *Connector) SetModifier(v string) *Connector {
	s.Modifier = &v
	return s
}

func (s *Connector) SetModifierName(v string) *Connector {
	s.ModifierName = &v
	return s
}

func (s *Connector) SetName(v string) *Connector {
	s.Name = &v
	return s
}

func (s *Connector) SetProperties(v []*Property) *Connector {
	s.Properties = v
	return s
}

func (s *Connector) SetSink(v bool) *Connector {
	s.Sink = &v
	return s
}

func (s *Connector) SetSource(v bool) *Connector {
	s.Source = &v
	return s
}

func (s *Connector) SetSupportedFormats(v []*string) *Connector {
	s.SupportedFormats = v
	return s
}

func (s *Connector) SetType(v string) *Connector {
	s.Type = &v
	return s
}

type CreateUdfArtifactResult struct {
	CollidingClasses []*UdfClass  `json:"collidingClasses,omitempty" xml:"collidingClasses,omitempty" type:"Repeated"`
	CreateSuccess    *bool        `json:"createSuccess,omitempty" xml:"createSuccess,omitempty"`
	Message          *string      `json:"message,omitempty" xml:"message,omitempty"`
	UdfArtifact      *UdfArtifact `json:"udfArtifact,omitempty" xml:"udfArtifact,omitempty"`
}

func (s CreateUdfArtifactResult) String() string {
	return tea.Prettify(s)
}

func (s CreateUdfArtifactResult) GoString() string {
	return s.String()
}

func (s *CreateUdfArtifactResult) SetCollidingClasses(v []*UdfClass) *CreateUdfArtifactResult {
	s.CollidingClasses = v
	return s
}

func (s *CreateUdfArtifactResult) SetCreateSuccess(v bool) *CreateUdfArtifactResult {
	s.CreateSuccess = &v
	return s
}

func (s *CreateUdfArtifactResult) SetMessage(v string) *CreateUdfArtifactResult {
	s.Message = &v
	return s
}

func (s *CreateUdfArtifactResult) SetUdfArtifact(v *UdfArtifact) *CreateUdfArtifactResult {
	s.UdfArtifact = v
	return s
}

type Database struct {
	Comment    *string                `json:"comment,omitempty" xml:"comment,omitempty"`
	Name       *string                `json:"name,omitempty" xml:"name,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
}

func (s Database) String() string {
	return tea.Prettify(s)
}

func (s Database) GoString() string {
	return s.String()
}

func (s *Database) SetComment(v string) *Database {
	s.Comment = &v
	return s
}

func (s *Database) SetName(v string) *Database {
	s.Name = &v
	return s
}

func (s *Database) SetProperties(v map[string]interface{}) *Database {
	s.Properties = v
	return s
}

type DeleteUdfArtifactResult struct {
	DeleteSuccess     *bool       `json:"deleteSuccess,omitempty" xml:"deleteSuccess,omitempty"`
	Message           *string     `json:"message,omitempty" xml:"message,omitempty"`
	ReferencedClasses []*UdfClass `json:"referencedClasses,omitempty" xml:"referencedClasses,omitempty" type:"Repeated"`
}

func (s DeleteUdfArtifactResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteUdfArtifactResult) GoString() string {
	return s.String()
}

func (s *DeleteUdfArtifactResult) SetDeleteSuccess(v bool) *DeleteUdfArtifactResult {
	s.DeleteSuccess = &v
	return s
}

func (s *DeleteUdfArtifactResult) SetMessage(v string) *DeleteUdfArtifactResult {
	s.Message = &v
	return s
}

func (s *DeleteUdfArtifactResult) SetReferencedClasses(v []*UdfClass) *DeleteUdfArtifactResult {
	s.ReferencedClasses = v
	return s
}

type Deployment struct {
	Artifact             *Artifact             `json:"artifact,omitempty" xml:"artifact,omitempty"`
	BatchResourceSetting *BatchResourceSetting `json:"batchResourceSetting,omitempty" xml:"batchResourceSetting,omitempty"`
	CreatedAt            *string               `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 27846363877456****
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// ****@streamcompute.onaliyun.com
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// example:
	//
	// true
	DeploymentHasChanged *bool `json:"deploymentHasChanged,omitempty" xml:"deploymentHasChanged,omitempty"`
	// example:
	//
	// 00000000-0000-0000-0000-000000000001
	DeploymentId     *string                `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	DeploymentTarget *BriefDeploymentTarget `json:"deploymentTarget,omitempty" xml:"deploymentTarget,omitempty"`
	// example:
	//
	// this is a deployment description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// vvr-6.0.0-flink-1.15
	EngineVersion *string `json:"engineVersion,omitempty" xml:"engineVersion,omitempty"`
	// example:
	//
	// STREAMING | BATCH
	ExecutionMode *string `json:"executionMode,omitempty" xml:"executionMode,omitempty"`
	// example:
	//
	// {"taskmanager.numberOfTaskSlots":"1"}
	FlinkConf      map[string]interface{} `json:"flinkConf,omitempty" xml:"flinkConf,omitempty"`
	JobSummary     *JobSummary            `json:"jobSummary,omitempty" xml:"jobSummary,omitempty"`
	Labels         map[string]interface{} `json:"labels,omitempty" xml:"labels,omitempty"`
	LocalVariables []*LocalVariable       `json:"localVariables,omitempty" xml:"localVariables,omitempty" type:"Repeated"`
	Logging        *Logging               `json:"logging,omitempty" xml:"logging,omitempty"`
	ModifiedAt     *string                `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	// example:
	//
	// 27846363877456****
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// ****@streamcompute.onaliyun.com
	ModifierName *string `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	// example:
	//
	// deploymentName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// default-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// 00000000-0000-0000-0000-000000000003
	ReferencedDeploymentDraftId *string                   `json:"referencedDeploymentDraftId,omitempty" xml:"referencedDeploymentDraftId,omitempty"`
	StreamingResourceSetting    *StreamingResourceSetting `json:"streamingResourceSetting,omitempty" xml:"streamingResourceSetting,omitempty"`
	// example:
	//
	// edcef******b4f
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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

func (s *Deployment) SetBatchResourceSetting(v *BatchResourceSetting) *Deployment {
	s.BatchResourceSetting = v
	return s
}

func (s *Deployment) SetCreatedAt(v string) *Deployment {
	s.CreatedAt = &v
	return s
}

func (s *Deployment) SetCreator(v string) *Deployment {
	s.Creator = &v
	return s
}

func (s *Deployment) SetCreatorName(v string) *Deployment {
	s.CreatorName = &v
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

func (s *Deployment) SetLabels(v map[string]interface{}) *Deployment {
	s.Labels = v
	return s
}

func (s *Deployment) SetLocalVariables(v []*LocalVariable) *Deployment {
	s.LocalVariables = v
	return s
}

func (s *Deployment) SetLogging(v *Logging) *Deployment {
	s.Logging = v
	return s
}

func (s *Deployment) SetModifiedAt(v string) *Deployment {
	s.ModifiedAt = &v
	return s
}

func (s *Deployment) SetModifier(v string) *Deployment {
	s.Modifier = &v
	return s
}

func (s *Deployment) SetModifierName(v string) *Deployment {
	s.ModifierName = &v
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

func (s *Deployment) SetReferencedDeploymentDraftId(v string) *Deployment {
	s.ReferencedDeploymentDraftId = &v
	return s
}

func (s *Deployment) SetStreamingResourceSetting(v *StreamingResourceSetting) *Deployment {
	s.StreamingResourceSetting = v
	return s
}

func (s *Deployment) SetWorkspace(v string) *Deployment {
	s.Workspace = &v
	return s
}

type DeploymentDraft struct {
	Artifact  *Artifact `json:"artifact,omitempty" xml:"artifact,omitempty"`
	CreatedAt *int64    `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 27846363877456****
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// ****@streamcompute.onaliyun.com
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// example:
	//
	// 00000000-0000-0000-0000-00000012****
	DeploymentDraftId *string `json:"deploymentDraftId,omitempty" xml:"deploymentDraftId,omitempty"`
	// example:
	//
	// vvr-6.0.7-flink-1.15
	EngineVersion *string `json:"engineVersion,omitempty" xml:"engineVersion,omitempty"`
	// example:
	//
	// STREAMING
	ExecutionMode  *string                `json:"executionMode,omitempty" xml:"executionMode,omitempty"`
	Labels         map[string]interface{} `json:"labels,omitempty" xml:"labels,omitempty"`
	LocalVariables []*LocalVariable       `json:"localVariables,omitempty" xml:"localVariables,omitempty" type:"Repeated"`
	Lock           *Lock                  `json:"lock,omitempty" xml:"lock,omitempty"`
	ModifiedAt     *int64                 `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	// example:
	//
	// 27846363877456****
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// ****@streamcompute.onaliyun.com
	ModifierName *string `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	// example:
	//
	// test-draft
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// default-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// 00000000-0000-0000-0000-00000013****
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// 00000000-0000-0000-0000-0000012312****
	ReferencedDeploymentId *string `json:"referencedDeploymentId,omitempty" xml:"referencedDeploymentId,omitempty"`
	// example:
	//
	// edcef******b4f
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeploymentDraft) String() string {
	return tea.Prettify(s)
}

func (s DeploymentDraft) GoString() string {
	return s.String()
}

func (s *DeploymentDraft) SetArtifact(v *Artifact) *DeploymentDraft {
	s.Artifact = v
	return s
}

func (s *DeploymentDraft) SetCreatedAt(v int64) *DeploymentDraft {
	s.CreatedAt = &v
	return s
}

func (s *DeploymentDraft) SetCreator(v string) *DeploymentDraft {
	s.Creator = &v
	return s
}

func (s *DeploymentDraft) SetCreatorName(v string) *DeploymentDraft {
	s.CreatorName = &v
	return s
}

func (s *DeploymentDraft) SetDeploymentDraftId(v string) *DeploymentDraft {
	s.DeploymentDraftId = &v
	return s
}

func (s *DeploymentDraft) SetEngineVersion(v string) *DeploymentDraft {
	s.EngineVersion = &v
	return s
}

func (s *DeploymentDraft) SetExecutionMode(v string) *DeploymentDraft {
	s.ExecutionMode = &v
	return s
}

func (s *DeploymentDraft) SetLabels(v map[string]interface{}) *DeploymentDraft {
	s.Labels = v
	return s
}

func (s *DeploymentDraft) SetLocalVariables(v []*LocalVariable) *DeploymentDraft {
	s.LocalVariables = v
	return s
}

func (s *DeploymentDraft) SetLock(v *Lock) *DeploymentDraft {
	s.Lock = v
	return s
}

func (s *DeploymentDraft) SetModifiedAt(v int64) *DeploymentDraft {
	s.ModifiedAt = &v
	return s
}

func (s *DeploymentDraft) SetModifier(v string) *DeploymentDraft {
	s.Modifier = &v
	return s
}

func (s *DeploymentDraft) SetModifierName(v string) *DeploymentDraft {
	s.ModifierName = &v
	return s
}

func (s *DeploymentDraft) SetName(v string) *DeploymentDraft {
	s.Name = &v
	return s
}

func (s *DeploymentDraft) SetNamespace(v string) *DeploymentDraft {
	s.Namespace = &v
	return s
}

func (s *DeploymentDraft) SetParentId(v string) *DeploymentDraft {
	s.ParentId = &v
	return s
}

func (s *DeploymentDraft) SetReferencedDeploymentId(v string) *DeploymentDraft {
	s.ReferencedDeploymentId = &v
	return s
}

func (s *DeploymentDraft) SetWorkspace(v string) *DeploymentDraft {
	s.Workspace = &v
	return s
}

type DeploymentRestoreStrategy struct {
	// example:
	//
	// TRUE
	AllowNonRestoredState *bool `json:"allowNonRestoredState,omitempty" xml:"allowNonRestoredState,omitempty"`
	// example:
	//
	// 1660293803155
	JobStartTimeInMs *int64 `json:"jobStartTimeInMs,omitempty" xml:"jobStartTimeInMs,omitempty"`
	// example:
	//
	// LATEST_STATE
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// example:
	//
	// 354dde66-a3ae-463e-967a-0b4107fd****
	SavepointId *string `json:"savepointId,omitempty" xml:"savepointId,omitempty"`
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
	// example:
	//
	// deployment target
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// namespace
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

type DraftDeployParams struct {
	DeploymentDraftId *string                `json:"deploymentDraftId,omitempty" xml:"deploymentDraftId,omitempty"`
	DeploymentTarget  *BriefDeploymentTarget `json:"deploymentTarget,omitempty" xml:"deploymentTarget,omitempty"`
	// example:
	//
	// false
	SkipValidate *bool `json:"skipValidate,omitempty" xml:"skipValidate,omitempty"`
}

func (s DraftDeployParams) String() string {
	return tea.Prettify(s)
}

func (s DraftDeployParams) GoString() string {
	return s.String()
}

func (s *DraftDeployParams) SetDeploymentDraftId(v string) *DraftDeployParams {
	s.DeploymentDraftId = &v
	return s
}

func (s *DraftDeployParams) SetDeploymentTarget(v *BriefDeploymentTarget) *DraftDeployParams {
	s.DeploymentTarget = v
	return s
}

func (s *DraftDeployParams) SetSkipValidate(v bool) *DraftDeployParams {
	s.SkipValidate = &v
	return s
}

type DraftDeployResult struct {
	ArtifactValidationDetail *ValidateStatementResult `json:"artifactValidationDetail,omitempty" xml:"artifactValidationDetail,omitempty"`
	// example:
	//
	// 58718c99-3b29-4c5e-93bb-c9fc4ec6****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DraftDeployResult) String() string {
	return tea.Prettify(s)
}

func (s DraftDeployResult) GoString() string {
	return s.String()
}

func (s *DraftDeployResult) SetArtifactValidationDetail(v *ValidateStatementResult) *DraftDeployResult {
	s.ArtifactValidationDetail = v
	return s
}

func (s *DraftDeployResult) SetDeploymentId(v string) *DraftDeployResult {
	s.DeploymentId = &v
	return s
}

func (s *DraftDeployResult) SetMessage(v string) *DraftDeployResult {
	s.Message = &v
	return s
}

func (s *DraftDeployResult) SetSuccess(v bool) *DraftDeployResult {
	s.Success = &v
	return s
}

type Edge struct {
	ColumnLineage []*Relation `json:"columnLineage,omitempty" xml:"columnLineage,omitempty" type:"Repeated"`
	TableLineage  []*Relation `json:"tableLineage,omitempty" xml:"tableLineage,omitempty" type:"Repeated"`
}

func (s Edge) String() string {
	return tea.Prettify(s)
}

func (s Edge) GoString() string {
	return s.String()
}

func (s *Edge) SetColumnLineage(v []*Relation) *Edge {
	s.ColumnLineage = v
	return s
}

func (s *Edge) SetTableLineage(v []*Relation) *Edge {
	s.TableLineage = v
	return s
}

type EditableNamespace struct {
	Namespace   *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Role        *string `json:"Role,omitempty" xml:"Role,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s EditableNamespace) String() string {
	return tea.Prettify(s)
}

func (s EditableNamespace) GoString() string {
	return s.String()
}

func (s *EditableNamespace) SetNamespace(v string) *EditableNamespace {
	s.Namespace = &v
	return s
}

func (s *EditableNamespace) SetRole(v string) *EditableNamespace {
	s.Role = &v
	return s
}

func (s *EditableNamespace) SetWorkspaceId(v string) *EditableNamespace {
	s.WorkspaceId = &v
	return s
}

type EngineVersionMetadata struct {
	// This parameter is required.
	//
	// example:
	//
	// vvr-6.0.0-flink-1.15
	EngineVersion *string                         `json:"engineVersion,omitempty" xml:"engineVersion,omitempty"`
	Features      *EngineVersionSupportedFeatures `json:"features,omitempty" xml:"features,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// STABLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
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
	// example:
	//
	// vvr-6.0.1-flink-1.15
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
	// example:
	//
	// true
	SupportNativeSavepoint *bool `json:"supportNativeSavepoint,omitempty" xml:"supportNativeSavepoint,omitempty"`
	// example:
	//
	// true
	UseForSqlDeployments *bool `json:"useForSqlDeployments,omitempty" xml:"useForSqlDeployments,omitempty"`
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

type ErrorDetails struct {
	ColumnNumber     *string   `json:"columnNumber,omitempty" xml:"columnNumber,omitempty"`
	EndColumnNumber  *string   `json:"endColumnNumber,omitempty" xml:"endColumnNumber,omitempty"`
	EndLineNumber    *string   `json:"endLineNumber,omitempty" xml:"endLineNumber,omitempty"`
	InvalidflinkConf []*string `json:"invalidflinkConf,omitempty" xml:"invalidflinkConf,omitempty" type:"Repeated"`
	LineNumber       *string   `json:"lineNumber,omitempty" xml:"lineNumber,omitempty"`
	Message          *string   `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ErrorDetails) String() string {
	return tea.Prettify(s)
}

func (s ErrorDetails) GoString() string {
	return s.String()
}

func (s *ErrorDetails) SetColumnNumber(v string) *ErrorDetails {
	s.ColumnNumber = &v
	return s
}

func (s *ErrorDetails) SetEndColumnNumber(v string) *ErrorDetails {
	s.EndColumnNumber = &v
	return s
}

func (s *ErrorDetails) SetEndLineNumber(v string) *ErrorDetails {
	s.EndLineNumber = &v
	return s
}

func (s *ErrorDetails) SetInvalidflinkConf(v []*string) *ErrorDetails {
	s.InvalidflinkConf = v
	return s
}

func (s *ErrorDetails) SetLineNumber(v string) *ErrorDetails {
	s.LineNumber = &v
	return s
}

func (s *ErrorDetails) SetMessage(v string) *ErrorDetails {
	s.Message = &v
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

type Folder struct {
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 00000000-0000-0000-0000-0000012312****
	FolderId   *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	ModifiedAt *int64  `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// default-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// 00000000-0000-0000-0000-0000012390****
	ParentId  *string      `json:"parentId,omitempty" xml:"parentId,omitempty"`
	SubFolder []*SubFolder `json:"subFolder,omitempty" xml:"subFolder,omitempty" type:"Repeated"`
	// example:
	//
	// edcef******b4f
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s Folder) String() string {
	return tea.Prettify(s)
}

func (s Folder) GoString() string {
	return s.String()
}

func (s *Folder) SetCreatedAt(v int64) *Folder {
	s.CreatedAt = &v
	return s
}

func (s *Folder) SetFolderId(v string) *Folder {
	s.FolderId = &v
	return s
}

func (s *Folder) SetModifiedAt(v int64) *Folder {
	s.ModifiedAt = &v
	return s
}

func (s *Folder) SetName(v string) *Folder {
	s.Name = &v
	return s
}

func (s *Folder) SetNamespace(v string) *Folder {
	s.Namespace = &v
	return s
}

func (s *Folder) SetParentId(v string) *Folder {
	s.ParentId = &v
	return s
}

func (s *Folder) SetSubFolder(v []*SubFolder) *Folder {
	s.SubFolder = v
	return s
}

func (s *Folder) SetWorkspace(v string) *Folder {
	s.Workspace = &v
	return s
}

type GetLineageInfoParams struct {
	Depth         *int64  `json:"depth,omitempty" xml:"depth,omitempty"`
	Direction     *string `json:"direction,omitempty" xml:"direction,omitempty"`
	Id            *string `json:"id,omitempty" xml:"id,omitempty"`
	IdType        *string `json:"idType,omitempty" xml:"idType,omitempty"`
	IsColumnLevel *bool   `json:"isColumnLevel,omitempty" xml:"isColumnLevel,omitempty"`
	IsTemporary   *bool   `json:"isTemporary,omitempty" xml:"isTemporary,omitempty"`
	Namespace     *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	Workspace     *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetLineageInfoParams) String() string {
	return tea.Prettify(s)
}

func (s GetLineageInfoParams) GoString() string {
	return s.String()
}

func (s *GetLineageInfoParams) SetDepth(v int64) *GetLineageInfoParams {
	s.Depth = &v
	return s
}

func (s *GetLineageInfoParams) SetDirection(v string) *GetLineageInfoParams {
	s.Direction = &v
	return s
}

func (s *GetLineageInfoParams) SetId(v string) *GetLineageInfoParams {
	s.Id = &v
	return s
}

func (s *GetLineageInfoParams) SetIdType(v string) *GetLineageInfoParams {
	s.IdType = &v
	return s
}

func (s *GetLineageInfoParams) SetIsColumnLevel(v bool) *GetLineageInfoParams {
	s.IsColumnLevel = &v
	return s
}

func (s *GetLineageInfoParams) SetIsTemporary(v bool) *GetLineageInfoParams {
	s.IsTemporary = &v
	return s
}

func (s *GetLineageInfoParams) SetNamespace(v string) *GetLineageInfoParams {
	s.Namespace = &v
	return s
}

func (s *GetLineageInfoParams) SetWorkspace(v string) *GetLineageInfoParams {
	s.Workspace = &v
	return s
}

type HotUpdateJobFailureInfo struct {
	FailureSeverity *string `json:"failureSeverity,omitempty" xml:"failureSeverity,omitempty"`
	Message         *string `json:"message,omitempty" xml:"message,omitempty"`
	Reason          *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s HotUpdateJobFailureInfo) String() string {
	return tea.Prettify(s)
}

func (s HotUpdateJobFailureInfo) GoString() string {
	return s.String()
}

func (s *HotUpdateJobFailureInfo) SetFailureSeverity(v string) *HotUpdateJobFailureInfo {
	s.FailureSeverity = &v
	return s
}

func (s *HotUpdateJobFailureInfo) SetMessage(v string) *HotUpdateJobFailureInfo {
	s.Message = &v
	return s
}

func (s *HotUpdateJobFailureInfo) SetReason(v string) *HotUpdateJobFailureInfo {
	s.Reason = &v
	return s
}

type HotUpdateJobParams struct {
	RescaleJobParam      *RescaleJobParam      `json:"rescaleJobParam,omitempty" xml:"rescaleJobParam,omitempty"`
	UpdateJobConfigParam *UpdateJobConfigParam `json:"updateJobConfigParam,omitempty" xml:"updateJobConfigParam,omitempty"`
}

func (s HotUpdateJobParams) String() string {
	return tea.Prettify(s)
}

func (s HotUpdateJobParams) GoString() string {
	return s.String()
}

func (s *HotUpdateJobParams) SetRescaleJobParam(v *RescaleJobParam) *HotUpdateJobParams {
	s.RescaleJobParam = v
	return s
}

func (s *HotUpdateJobParams) SetUpdateJobConfigParam(v *UpdateJobConfigParam) *HotUpdateJobParams {
	s.UpdateJobConfigParam = v
	return s
}

type HotUpdateJobResult struct {
	HotUpdateParams       *HotUpdateJobParams   `json:"hotUpdateParams,omitempty" xml:"hotUpdateParams,omitempty"`
	JobHotUpdateId        *string               `json:"jobHotUpdateId,omitempty" xml:"jobHotUpdateId,omitempty"`
	JobId                 *string               `json:"jobId,omitempty" xml:"jobId,omitempty"`
	Status                *HotUpdateJobStatus   `json:"status,omitempty" xml:"status,omitempty"`
	TargetResourceSetting *BriefResourceSetting `json:"targetResourceSetting,omitempty" xml:"targetResourceSetting,omitempty"`
}

func (s HotUpdateJobResult) String() string {
	return tea.Prettify(s)
}

func (s HotUpdateJobResult) GoString() string {
	return s.String()
}

func (s *HotUpdateJobResult) SetHotUpdateParams(v *HotUpdateJobParams) *HotUpdateJobResult {
	s.HotUpdateParams = v
	return s
}

func (s *HotUpdateJobResult) SetJobHotUpdateId(v string) *HotUpdateJobResult {
	s.JobHotUpdateId = &v
	return s
}

func (s *HotUpdateJobResult) SetJobId(v string) *HotUpdateJobResult {
	s.JobId = &v
	return s
}

func (s *HotUpdateJobResult) SetStatus(v *HotUpdateJobStatus) *HotUpdateJobResult {
	s.Status = v
	return s
}

func (s *HotUpdateJobResult) SetTargetResourceSetting(v *BriefResourceSetting) *HotUpdateJobResult {
	s.TargetResourceSetting = v
	return s
}

type HotUpdateJobStatus struct {
	Failure   *HotUpdateJobFailureInfo `json:"failure,omitempty" xml:"failure,omitempty"`
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Status    *string                  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s HotUpdateJobStatus) String() string {
	return tea.Prettify(s)
}

func (s HotUpdateJobStatus) GoString() string {
	return s.String()
}

func (s *HotUpdateJobStatus) SetFailure(v *HotUpdateJobFailureInfo) *HotUpdateJobStatus {
	s.Failure = v
	return s
}

func (s *HotUpdateJobStatus) SetRequestId(v string) *HotUpdateJobStatus {
	s.RequestId = &v
	return s
}

func (s *HotUpdateJobStatus) SetStatus(v string) *HotUpdateJobStatus {
	s.Status = &v
	return s
}

type JarArtifact struct {
	AdditionalDependencies []*string `json:"additionalDependencies,omitempty" xml:"additionalDependencies,omitempty" type:"Repeated"`
	// example:
	//
	// org.apapche.flink.test
	EntryClass *string `json:"entryClass,omitempty" xml:"entryClass,omitempty"`
	// example:
	//
	// https://oss//bucket//test.jar
	JarUri   *string `json:"jarUri,omitempty" xml:"jarUri,omitempty"`
	MainArgs *string `json:"mainArgs,omitempty" xml:"mainArgs,omitempty"`
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
	Artifact             *Artifact             `json:"artifact,omitempty" xml:"artifact,omitempty"`
	BatchResourceSetting *BatchResourceSetting `json:"batchResourceSetting,omitempty" xml:"batchResourceSetting,omitempty"`
	CreatedAt            *string               `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 27846363877456****
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// ****@streamcompute.onaliyun.com
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// example:
	//
	// 354dde66-a3ae-463e-967a-0b4107fd****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// example:
	//
	// flinktest
	DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty"`
	// example:
	//
	// 1660277235
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// vvr-4.0.14-flink-1.13
	EngineVersion *string `json:"engineVersion,omitempty" xml:"engineVersion,omitempty"`
	// example:
	//
	// BATCH
	ExecutionMode *string                `json:"executionMode,omitempty" xml:"executionMode,omitempty"`
	FlinkConf     map[string]interface{} `json:"flinkConf,omitempty" xml:"flinkConf,omitempty"`
	// example:
	//
	// 354dde66-a3ae-463e-967a-0b4107fd****
	JobId          *string          `json:"jobId,omitempty" xml:"jobId,omitempty"`
	LocalVariables []*LocalVariable `json:"localVariables,omitempty" xml:"localVariables,omitempty" type:"Repeated"`
	Logging        *Logging         `json:"logging,omitempty" xml:"logging,omitempty"`
	Metric         *JobMetric       `json:"metric,omitempty" xml:"metric,omitempty"`
	ModifiedAt     *string          `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	// example:
	//
	// 27846363877456****
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// ****@streamcompute.onaliyun.com
	ModifierName *string `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	// example:
	//
	// namespacetest
	Namespace       *string                    `json:"namespace,omitempty" xml:"namespace,omitempty"`
	RestoreStrategy *DeploymentRestoreStrategy `json:"restoreStrategy,omitempty" xml:"restoreStrategy,omitempty"`
	// example:
	//
	// preview
	SessionClusterName *string `json:"sessionClusterName,omitempty" xml:"sessionClusterName,omitempty"`
	// example:
	//
	// 1660190835
	StartTime                *int64                    `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status                   *JobStatus                `json:"status,omitempty" xml:"status,omitempty"`
	StreamingResourceSetting *StreamingResourceSetting `json:"streamingResourceSetting,omitempty" xml:"streamingResourceSetting,omitempty"`
	UserFlinkConf            map[string]interface{}    `json:"userFlinkConf,omitempty" xml:"userFlinkConf,omitempty"`
	// example:
	//
	// edcef******b4f
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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

func (s *Job) SetCreatedAt(v string) *Job {
	s.CreatedAt = &v
	return s
}

func (s *Job) SetCreator(v string) *Job {
	s.Creator = &v
	return s
}

func (s *Job) SetCreatorName(v string) *Job {
	s.CreatorName = &v
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

func (s *Job) SetLocalVariables(v []*LocalVariable) *Job {
	s.LocalVariables = v
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

func (s *Job) SetModifiedAt(v string) *Job {
	s.ModifiedAt = &v
	return s
}

func (s *Job) SetModifier(v string) *Job {
	s.Modifier = &v
	return s
}

func (s *Job) SetModifierName(v string) *Job {
	s.ModifierName = &v
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

func (s *Job) SetUserFlinkConf(v map[string]interface{}) *Job {
	s.UserFlinkConf = v
	return s
}

func (s *Job) SetWorkspace(v string) *Job {
	s.Workspace = &v
	return s
}

type JobFailure struct {
	// example:
	//
	// 1660120062
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

type JobInfo struct {
	Id         *string                `json:"id,omitempty" xml:"id,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
}

func (s JobInfo) String() string {
	return tea.Prettify(s)
}

func (s JobInfo) GoString() string {
	return s.String()
}

func (s *JobInfo) SetId(v string) *JobInfo {
	s.Id = &v
	return s
}

func (s *JobInfo) SetProperties(v map[string]interface{}) *JobInfo {
	s.Properties = v
	return s
}

type JobMetric struct {
	// example:
	//
	// 2
	TotalCpu *float64 `json:"totalCpu,omitempty" xml:"totalCpu,omitempty"`
	// example:
	//
	// 4096
	TotalMemoryByte *int64 `json:"totalMemoryByte,omitempty" xml:"totalMemoryByte,omitempty"`
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

type JobStartParameters struct {
	DeploymentId   *string          `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	LocalVariables []*LocalVariable `json:"localVariables,omitempty" xml:"localVariables,omitempty" type:"Repeated"`
	// example:
	//
	// default-queue
	ResourceQueueName *string                    `json:"resourceQueueName,omitempty" xml:"resourceQueueName,omitempty"`
	RestoreStrategy   *DeploymentRestoreStrategy `json:"restoreStrategy,omitempty" xml:"restoreStrategy,omitempty"`
}

func (s JobStartParameters) String() string {
	return tea.Prettify(s)
}

func (s JobStartParameters) GoString() string {
	return s.String()
}

func (s *JobStartParameters) SetDeploymentId(v string) *JobStartParameters {
	s.DeploymentId = &v
	return s
}

func (s *JobStartParameters) SetLocalVariables(v []*LocalVariable) *JobStartParameters {
	s.LocalVariables = v
	return s
}

func (s *JobStartParameters) SetResourceQueueName(v string) *JobStartParameters {
	s.ResourceQueueName = &v
	return s
}

func (s *JobStartParameters) SetRestoreStrategy(v *DeploymentRestoreStrategy) *JobStartParameters {
	s.RestoreStrategy = v
	return s
}

type JobStatus struct {
	// example:
	//
	// RUNNING
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
	// example:
	//
	// 4
	ObservedFlinkJobRestarts *int64 `json:"observedFlinkJobRestarts,omitempty" xml:"observedFlinkJobRestarts,omitempty"`
	// example:
	//
	// RUNNING
	ObservedFlinkJobStatus *string `json:"observedFlinkJobStatus,omitempty" xml:"observedFlinkJobStatus,omitempty"`
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
	// example:
	//
	// 1
	Cancelled *int32 `json:"cancelled,omitempty" xml:"cancelled,omitempty"`
	// example:
	//
	// 1
	Cancelling *int32 `json:"cancelling,omitempty" xml:"cancelling,omitempty"`
	// example:
	//
	// 1
	Failed *int32 `json:"failed,omitempty" xml:"failed,omitempty"`
	// example:
	//
	// 1
	Finished *int32 `json:"finished,omitempty" xml:"finished,omitempty"`
	// example:
	//
	// 1
	Running *int32 `json:"running,omitempty" xml:"running,omitempty"`
	// example:
	//
	// 1
	Starting *int32 `json:"starting,omitempty" xml:"starting,omitempty"`
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

type LineageColumn struct {
	ColumnName       *string `json:"columnName,omitempty" xml:"columnName,omitempty"`
	ColumnNativeType *string `json:"columnNativeType,omitempty" xml:"columnNativeType,omitempty"`
	ColumnType       *string `json:"columnType,omitempty" xml:"columnType,omitempty"`
	CreatedAt        *int64  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Creator          *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Description      *string `json:"description,omitempty" xml:"description,omitempty"`
	Id               *string `json:"id,omitempty" xml:"id,omitempty"`
	ModifiedAt       *int64  `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	Modifier         *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	Nullable         *bool   `json:"nullable,omitempty" xml:"nullable,omitempty"`
}

func (s LineageColumn) String() string {
	return tea.Prettify(s)
}

func (s LineageColumn) GoString() string {
	return s.String()
}

func (s *LineageColumn) SetColumnName(v string) *LineageColumn {
	s.ColumnName = &v
	return s
}

func (s *LineageColumn) SetColumnNativeType(v string) *LineageColumn {
	s.ColumnNativeType = &v
	return s
}

func (s *LineageColumn) SetColumnType(v string) *LineageColumn {
	s.ColumnType = &v
	return s
}

func (s *LineageColumn) SetCreatedAt(v int64) *LineageColumn {
	s.CreatedAt = &v
	return s
}

func (s *LineageColumn) SetCreator(v string) *LineageColumn {
	s.Creator = &v
	return s
}

func (s *LineageColumn) SetDescription(v string) *LineageColumn {
	s.Description = &v
	return s
}

func (s *LineageColumn) SetId(v string) *LineageColumn {
	s.Id = &v
	return s
}

func (s *LineageColumn) SetModifiedAt(v int64) *LineageColumn {
	s.ModifiedAt = &v
	return s
}

func (s *LineageColumn) SetModifier(v string) *LineageColumn {
	s.Modifier = &v
	return s
}

func (s *LineageColumn) SetNullable(v bool) *LineageColumn {
	s.Nullable = &v
	return s
}

type LineageInfo struct {
	Edges    *Edge      `json:"edges,omitempty" xml:"edges,omitempty"`
	JobInfos []*JobInfo `json:"jobInfos,omitempty" xml:"jobInfos,omitempty" type:"Repeated"`
	Nodes    []*Node    `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
}

func (s LineageInfo) String() string {
	return tea.Prettify(s)
}

func (s LineageInfo) GoString() string {
	return s.String()
}

func (s *LineageInfo) SetEdges(v *Edge) *LineageInfo {
	s.Edges = v
	return s
}

func (s *LineageInfo) SetJobInfos(v []*JobInfo) *LineageInfo {
	s.JobInfos = v
	return s
}

func (s *LineageInfo) SetNodes(v []*Node) *LineageInfo {
	s.Nodes = v
	return s
}

type LineageTable struct {
	Columns    []*LineageColumn       `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	Id         *string                `json:"id,omitempty" xml:"id,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
	TableName  *string                `json:"tableName,omitempty" xml:"tableName,omitempty"`
	With       map[string]interface{} `json:"with,omitempty" xml:"with,omitempty"`
}

func (s LineageTable) String() string {
	return tea.Prettify(s)
}

func (s LineageTable) GoString() string {
	return s.String()
}

func (s *LineageTable) SetColumns(v []*LineageColumn) *LineageTable {
	s.Columns = v
	return s
}

func (s *LineageTable) SetId(v string) *LineageTable {
	s.Id = &v
	return s
}

func (s *LineageTable) SetProperties(v map[string]interface{}) *LineageTable {
	s.Properties = v
	return s
}

func (s *LineageTable) SetTableName(v string) *LineageTable {
	s.TableName = &v
	return s
}

func (s *LineageTable) SetWith(v map[string]interface{}) *LineageTable {
	s.With = v
	return s
}

type LocalVariable struct {
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// datagen
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s LocalVariable) String() string {
	return tea.Prettify(s)
}

func (s LocalVariable) GoString() string {
	return s.String()
}

func (s *LocalVariable) SetName(v string) *LocalVariable {
	s.Name = &v
	return s
}

func (s *LocalVariable) SetValue(v string) *LocalVariable {
	s.Value = &v
	return s
}

type Lock struct {
	HolderId   *string `json:"holderId,omitempty" xml:"holderId,omitempty"`
	HolderName *string `json:"holderName,omitempty" xml:"holderName,omitempty"`
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	Namespace  *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	Workspace  *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s Lock) String() string {
	return tea.Prettify(s)
}

func (s Lock) GoString() string {
	return s.String()
}

func (s *Lock) SetHolderId(v string) *Lock {
	s.HolderId = &v
	return s
}

func (s *Lock) SetHolderName(v string) *Lock {
	s.HolderName = &v
	return s
}

func (s *Lock) SetId(v string) *Lock {
	s.Id = &v
	return s
}

func (s *Lock) SetNamespace(v string) *Lock {
	s.Namespace = &v
	return s
}

func (s *Lock) SetWorkspace(v string) *Lock {
	s.Workspace = &v
	return s
}

type Log4jLogger struct {
	// example:
	//
	// ERROR
	LoggerLevel *string `json:"loggerLevel,omitempty" xml:"loggerLevel,omitempty"`
	// example:
	//
	// StdOutErrConsoleAppender
	LoggerName *string `json:"loggerName,omitempty" xml:"loggerName,omitempty"`
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
	// example:
	//
	// 7
	ExpirationDays *int64 `json:"expirationDays,omitempty" xml:"expirationDays,omitempty"`
	// example:
	//
	// true
	OpenHistory *bool `json:"openHistory,omitempty" xml:"openHistory,omitempty"`
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
	// example:
	//
	// xml格式文本
	Log4j2ConfigurationTemplate *string           `json:"log4j2ConfigurationTemplate,omitempty" xml:"log4j2ConfigurationTemplate,omitempty"`
	Log4jLoggers                []*Log4jLogger    `json:"log4jLoggers,omitempty" xml:"log4jLoggers,omitempty" type:"Repeated"`
	LogReservePolicy            *LogReservePolicy `json:"logReservePolicy,omitempty" xml:"logReservePolicy,omitempty"`
	// example:
	//
	// oss
	LoggingProfile *string `json:"loggingProfile,omitempty" xml:"loggingProfile,omitempty"`
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

type Member struct {
	// This parameter is required.
	//
	// example:
	//
	// user: 181319557522****
	Member *string `json:"member,omitempty" xml:"member,omitempty"`
	// example:
	//
	// VIEWER
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s Member) String() string {
	return tea.Prettify(s)
}

func (s Member) GoString() string {
	return s.String()
}

func (s *Member) SetMember(v string) *Member {
	s.Member = &v
	return s
}

func (s *Member) SetRole(v string) *Member {
	s.Role = &v
	return s
}

type MetadataInfo struct {
	Key     *string `json:"key,omitempty" xml:"key,omitempty"`
	Virtual *bool   `json:"virtual,omitempty" xml:"virtual,omitempty"`
}

func (s MetadataInfo) String() string {
	return tea.Prettify(s)
}

func (s MetadataInfo) GoString() string {
	return s.String()
}

func (s *MetadataInfo) SetKey(v string) *MetadataInfo {
	s.Key = &v
	return s
}

func (s *MetadataInfo) SetVirtual(v bool) *MetadataInfo {
	s.Virtual = &v
	return s
}

type Node struct {
	CatalogName  *string         `json:"catalogName,omitempty" xml:"catalogName,omitempty"`
	Connector    *string         `json:"connector,omitempty" xml:"connector,omitempty"`
	DatabaseName *string         `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	Id           *string         `json:"id,omitempty" xml:"id,omitempty"`
	IsTemporary  *bool           `json:"isTemporary,omitempty" xml:"isTemporary,omitempty"`
	Tables       []*LineageTable `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
}

func (s Node) String() string {
	return tea.Prettify(s)
}

func (s Node) GoString() string {
	return s.String()
}

func (s *Node) SetCatalogName(v string) *Node {
	s.CatalogName = &v
	return s
}

func (s *Node) SetConnector(v string) *Node {
	s.Connector = &v
	return s
}

func (s *Node) SetDatabaseName(v string) *Node {
	s.DatabaseName = &v
	return s
}

func (s *Node) SetId(v string) *Node {
	s.Id = &v
	return s
}

func (s *Node) SetIsTemporary(v bool) *Node {
	s.IsTemporary = &v
	return s
}

func (s *Node) SetTables(v []*LineageTable) *Node {
	s.Tables = v
	return s
}

type PrimaryKey struct {
	Columns        []*string `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	ConstraintName *string   `json:"constraintName,omitempty" xml:"constraintName,omitempty"`
}

func (s PrimaryKey) String() string {
	return tea.Prettify(s)
}

func (s PrimaryKey) GoString() string {
	return s.String()
}

func (s *PrimaryKey) SetColumns(v []*string) *PrimaryKey {
	s.Columns = v
	return s
}

func (s *PrimaryKey) SetConstraintName(v string) *PrimaryKey {
	s.ConstraintName = &v
	return s
}

type Property struct {
	DefaultValue  *string `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	DefinesFormat *bool   `json:"definesFormat,omitempty" xml:"definesFormat,omitempty"`
	Description   *string `json:"description,omitempty" xml:"description,omitempty"`
	Key           *string `json:"key,omitempty" xml:"key,omitempty"`
	Required      *bool   `json:"required,omitempty" xml:"required,omitempty"`
	Sensitive     *bool   `json:"sensitive,omitempty" xml:"sensitive,omitempty"`
}

func (s Property) String() string {
	return tea.Prettify(s)
}

func (s Property) GoString() string {
	return s.String()
}

func (s *Property) SetDefaultValue(v string) *Property {
	s.DefaultValue = &v
	return s
}

func (s *Property) SetDefinesFormat(v bool) *Property {
	s.DefinesFormat = &v
	return s
}

func (s *Property) SetDescription(v string) *Property {
	s.Description = &v
	return s
}

func (s *Property) SetKey(v string) *Property {
	s.Key = &v
	return s
}

func (s *Property) SetRequired(v bool) *Property {
	s.Required = &v
	return s
}

func (s *Property) SetSensitive(v bool) *Property {
	s.Sensitive = &v
	return s
}

type PythonArtifact struct {
	AdditionalDependencies    []*string `json:"additionalDependencies,omitempty" xml:"additionalDependencies,omitempty" type:"Repeated"`
	AdditionalPythonArchives  []*string `json:"additionalPythonArchives,omitempty" xml:"additionalPythonArchives,omitempty" type:"Repeated"`
	AdditionalPythonLibraries []*string `json:"additionalPythonLibraries,omitempty" xml:"additionalPythonLibraries,omitempty" type:"Repeated"`
	EntryModule               *string   `json:"entryModule,omitempty" xml:"entryModule,omitempty"`
	MainArgs                  *string   `json:"mainArgs,omitempty" xml:"mainArgs,omitempty"`
	// example:
	//
	// https://oss//bucket//test.py
	PythonArtifactUri *string `json:"pythonArtifactUri,omitempty" xml:"pythonArtifactUri,omitempty"`
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

type Relation struct {
	Destination *string `json:"destination,omitempty" xml:"destination,omitempty"`
	JobId       *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	Source      *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s Relation) String() string {
	return tea.Prettify(s)
}

func (s Relation) GoString() string {
	return s.String()
}

func (s *Relation) SetDestination(v string) *Relation {
	s.Destination = &v
	return s
}

func (s *Relation) SetJobId(v string) *Relation {
	s.JobId = &v
	return s
}

func (s *Relation) SetSource(v string) *Relation {
	s.Source = &v
	return s
}

type RescaleJobParam struct {
	JobParallelism    *int64                 `json:"jobParallelism,omitempty" xml:"jobParallelism,omitempty"`
	VertexParallelism map[string]interface{} `json:"vertexParallelism,omitempty" xml:"vertexParallelism,omitempty"`
}

func (s RescaleJobParam) String() string {
	return tea.Prettify(s)
}

func (s RescaleJobParam) GoString() string {
	return s.String()
}

func (s *RescaleJobParam) SetJobParallelism(v int64) *RescaleJobParam {
	s.JobParallelism = &v
	return s
}

func (s *RescaleJobParam) SetVertexParallelism(v map[string]interface{}) *RescaleJobParam {
	s.VertexParallelism = v
	return s
}

type Savepoint struct {
	// example:
	//
	// 1659066711
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 1d716b22-6aad-4be2-85c2-50cfc757****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	Description  *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 5af678c0-7db0-4650-94c2-d2604f0a****
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// 1659069473
	ModifiedAt *int64 `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	// example:
	//
	// namespacetest
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// TRUE
	NativeFormat *bool `json:"nativeFormat,omitempty" xml:"nativeFormat,omitempty"`
	// example:
	//
	// 354dde66-a3ae-463e-967a-0b4107fd****
	SavepointId *string `json:"savepointId,omitempty" xml:"savepointId,omitempty"`
	// example:
	//
	// https://oss/bucket/flink/flink-jobs/namespaces/vvp-team/deployments/5a19a71b-1c42-4f34-94fd-86cf60782c81/checkpoints/sp-3285
	SavepointLocation *string `json:"savepointLocation,omitempty" xml:"savepointLocation,omitempty"`
	// example:
	//
	// USER_REQUEST
	SavepointOrigin *string          `json:"savepointOrigin,omitempty" xml:"savepointOrigin,omitempty"`
	Status          *SavepointStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// TRUE
	StopWithDrainEnabled *bool `json:"stopWithDrainEnabled,omitempty" xml:"stopWithDrainEnabled,omitempty"`
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
	// example:
	//
	// 1655006835
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
	// example:
	//
	// COMPLETED
	State *string `json:"state,omitempty" xml:"state,omitempty"`
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

type Schema struct {
	Columns        []*TableColumn   `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	PrimaryKey     []*PrimaryKey    `json:"primaryKey,omitempty" xml:"primaryKey,omitempty" type:"Repeated"`
	WatermarkSpecs []*WatermarkSpec `json:"watermarkSpecs,omitempty" xml:"watermarkSpecs,omitempty" type:"Repeated"`
}

func (s Schema) String() string {
	return tea.Prettify(s)
}

func (s Schema) GoString() string {
	return s.String()
}

func (s *Schema) SetColumns(v []*TableColumn) *Schema {
	s.Columns = v
	return s
}

func (s *Schema) SetPrimaryKey(v []*PrimaryKey) *Schema {
	s.PrimaryKey = v
	return s
}

func (s *Schema) SetWatermarkSpecs(v []*WatermarkSpec) *Schema {
	s.WatermarkSpecs = v
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

type SqlStatementExecuteResult struct {
	ErrorDetails   *ErrorDetails `json:"errorDetails,omitempty" xml:"errorDetails,omitempty"`
	ExecuteSuccess *bool         `json:"executeSuccess,omitempty" xml:"executeSuccess,omitempty"`
	Statement      *string       `json:"statement,omitempty" xml:"statement,omitempty"`
}

func (s SqlStatementExecuteResult) String() string {
	return tea.Prettify(s)
}

func (s SqlStatementExecuteResult) GoString() string {
	return s.String()
}

func (s *SqlStatementExecuteResult) SetErrorDetails(v *ErrorDetails) *SqlStatementExecuteResult {
	s.ErrorDetails = v
	return s
}

func (s *SqlStatementExecuteResult) SetExecuteSuccess(v bool) *SqlStatementExecuteResult {
	s.ExecuteSuccess = &v
	return s
}

func (s *SqlStatementExecuteResult) SetStatement(v string) *SqlStatementExecuteResult {
	s.Statement = &v
	return s
}

type SqlStatementValidationResult struct {
	ErrorDetails     *ErrorDetails `json:"errorDetails,omitempty" xml:"errorDetails,omitempty"`
	Message          *string       `json:"message,omitempty" xml:"message,omitempty"`
	Success          *bool         `json:"success,omitempty" xml:"success,omitempty"`
	ValidationResult *string       `json:"validationResult,omitempty" xml:"validationResult,omitempty"`
}

func (s SqlStatementValidationResult) String() string {
	return tea.Prettify(s)
}

func (s SqlStatementValidationResult) GoString() string {
	return s.String()
}

func (s *SqlStatementValidationResult) SetErrorDetails(v *ErrorDetails) *SqlStatementValidationResult {
	s.ErrorDetails = v
	return s
}

func (s *SqlStatementValidationResult) SetMessage(v string) *SqlStatementValidationResult {
	s.Message = &v
	return s
}

func (s *SqlStatementValidationResult) SetSuccess(v bool) *SqlStatementValidationResult {
	s.Success = &v
	return s
}

func (s *SqlStatementValidationResult) SetValidationResult(v string) *SqlStatementValidationResult {
	s.ValidationResult = &v
	return s
}

type SqlStatementWithContext struct {
	AdditionalDependencies []*string `json:"additionalDependencies,omitempty" xml:"additionalDependencies,omitempty" type:"Repeated"`
	// This parameter is required.
	BatchMode          *bool                  `json:"batchMode,omitempty" xml:"batchMode,omitempty"`
	FlinkConfiguration map[string]interface{} `json:"flinkConfiguration,omitempty" xml:"flinkConfiguration,omitempty"`
	// This parameter is required.
	Statement   *string `json:"statement,omitempty" xml:"statement,omitempty"`
	VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty"`
}

func (s SqlStatementWithContext) String() string {
	return tea.Prettify(s)
}

func (s SqlStatementWithContext) GoString() string {
	return s.String()
}

func (s *SqlStatementWithContext) SetAdditionalDependencies(v []*string) *SqlStatementWithContext {
	s.AdditionalDependencies = v
	return s
}

func (s *SqlStatementWithContext) SetBatchMode(v bool) *SqlStatementWithContext {
	s.BatchMode = &v
	return s
}

func (s *SqlStatementWithContext) SetFlinkConfiguration(v map[string]interface{}) *SqlStatementWithContext {
	s.FlinkConfiguration = v
	return s
}

func (s *SqlStatementWithContext) SetStatement(v string) *SqlStatementWithContext {
	s.Statement = &v
	return s
}

func (s *SqlStatementWithContext) SetVersionName(v string) *SqlStatementWithContext {
	s.VersionName = &v
	return s
}

type StartJobRequestBody struct {
	// example:
	//
	// 5a19a71b-1c42-4f34-94fd-86cf60782c81
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
	// This parameter is required.
	//
	// example:
	//
	// NONE
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
	// example:
	//
	// EXPERT
	ResourceSettingMode *string `json:"resourceSettingMode,omitempty" xml:"resourceSettingMode,omitempty"`
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

type SubFolder struct {
	// example:
	//
	// a579aec9-1d5e-3382-9d65-9887ff6cfaff
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 95c0787c-408f-4e1f-88ba-ef0a84a2c2ee
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
}

func (s SubFolder) String() string {
	return tea.Prettify(s)
}

func (s SubFolder) GoString() string {
	return s.String()
}

func (s *SubFolder) SetFolderId(v string) *SubFolder {
	s.FolderId = &v
	return s
}

func (s *SubFolder) SetName(v string) *SubFolder {
	s.Name = &v
	return s
}

func (s *SubFolder) SetParentId(v string) *SubFolder {
	s.ParentId = &v
	return s
}

type Table struct {
	Comment       *string                `json:"comment,omitempty" xml:"comment,omitempty"`
	Name          *string                `json:"name,omitempty" xml:"name,omitempty"`
	PartitionKeys []*string              `json:"partitionKeys,omitempty" xml:"partitionKeys,omitempty" type:"Repeated"`
	Properties    map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
	Schema        *Schema                `json:"schema,omitempty" xml:"schema,omitempty"`
}

func (s Table) String() string {
	return tea.Prettify(s)
}

func (s Table) GoString() string {
	return s.String()
}

func (s *Table) SetComment(v string) *Table {
	s.Comment = &v
	return s
}

func (s *Table) SetName(v string) *Table {
	s.Name = &v
	return s
}

func (s *Table) SetPartitionKeys(v []*string) *Table {
	s.PartitionKeys = v
	return s
}

func (s *Table) SetProperties(v map[string]interface{}) *Table {
	s.Properties = v
	return s
}

func (s *Table) SetSchema(v *Schema) *Table {
	s.Schema = v
	return s
}

type TableColumn struct {
	Expression   *string       `json:"expression,omitempty" xml:"expression,omitempty"`
	MetadataInfo *MetadataInfo `json:"metadataInfo,omitempty" xml:"metadataInfo,omitempty"`
	Name         *string       `json:"name,omitempty" xml:"name,omitempty"`
	Nullable     *bool         `json:"nullable,omitempty" xml:"nullable,omitempty"`
	Type         *string       `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TableColumn) String() string {
	return tea.Prettify(s)
}

func (s TableColumn) GoString() string {
	return s.String()
}

func (s *TableColumn) SetExpression(v string) *TableColumn {
	s.Expression = &v
	return s
}

func (s *TableColumn) SetMetadataInfo(v *MetadataInfo) *TableColumn {
	s.MetadataInfo = v
	return s
}

func (s *TableColumn) SetName(v string) *TableColumn {
	s.Name = &v
	return s
}

func (s *TableColumn) SetNullable(v bool) *TableColumn {
	s.Nullable = &v
	return s
}

func (s *TableColumn) SetType(v string) *TableColumn {
	s.Type = &v
	return s
}

type TableMeta struct {
	CatalogName  *string `json:"catalogName,omitempty" xml:"catalogName,omitempty"`
	DatabaseName *string `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	TableName    *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
}

func (s TableMeta) String() string {
	return tea.Prettify(s)
}

func (s TableMeta) GoString() string {
	return s.String()
}

func (s *TableMeta) SetCatalogName(v string) *TableMeta {
	s.CatalogName = &v
	return s
}

func (s *TableMeta) SetDatabaseName(v string) *TableMeta {
	s.DatabaseName = &v
	return s
}

func (s *TableMeta) SetTableName(v string) *TableMeta {
	s.TableName = &v
	return s
}

type UdfArtifact struct {
	ArtifactType      *string     `json:"artifactType,omitempty" xml:"artifactType,omitempty"`
	CreatedAt         *int64      `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Creator           *string     `json:"creator,omitempty" xml:"creator,omitempty"`
	DependencyJarUris []*string   `json:"dependencyJarUris,omitempty" xml:"dependencyJarUris,omitempty" type:"Repeated"`
	JarUrl            *string     `json:"jarUrl,omitempty" xml:"jarUrl,omitempty"`
	ModifiedAt        *int64      `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	Name              *string     `json:"name,omitempty" xml:"name,omitempty"`
	Namespace         *string     `json:"namespace,omitempty" xml:"namespace,omitempty"`
	UdfClasses        []*UdfClass `json:"udfClasses,omitempty" xml:"udfClasses,omitempty" type:"Repeated"`
}

func (s UdfArtifact) String() string {
	return tea.Prettify(s)
}

func (s UdfArtifact) GoString() string {
	return s.String()
}

func (s *UdfArtifact) SetArtifactType(v string) *UdfArtifact {
	s.ArtifactType = &v
	return s
}

func (s *UdfArtifact) SetCreatedAt(v int64) *UdfArtifact {
	s.CreatedAt = &v
	return s
}

func (s *UdfArtifact) SetCreator(v string) *UdfArtifact {
	s.Creator = &v
	return s
}

func (s *UdfArtifact) SetDependencyJarUris(v []*string) *UdfArtifact {
	s.DependencyJarUris = v
	return s
}

func (s *UdfArtifact) SetJarUrl(v string) *UdfArtifact {
	s.JarUrl = &v
	return s
}

func (s *UdfArtifact) SetModifiedAt(v int64) *UdfArtifact {
	s.ModifiedAt = &v
	return s
}

func (s *UdfArtifact) SetName(v string) *UdfArtifact {
	s.Name = &v
	return s
}

func (s *UdfArtifact) SetNamespace(v string) *UdfArtifact {
	s.Namespace = &v
	return s
}

func (s *UdfArtifact) SetUdfClasses(v []*UdfClass) *UdfArtifact {
	s.UdfClasses = v
	return s
}

type UdfClass struct {
	ClassName       *string   `json:"className,omitempty" xml:"className,omitempty"`
	ClassType       *string   `json:"classType,omitempty" xml:"classType,omitempty"`
	FunctionNames   []*string `json:"functionNames,omitempty" xml:"functionNames,omitempty" type:"Repeated"`
	UdfArtifactName *string   `json:"udfArtifactName,omitempty" xml:"udfArtifactName,omitempty"`
}

func (s UdfClass) String() string {
	return tea.Prettify(s)
}

func (s UdfClass) GoString() string {
	return s.String()
}

func (s *UdfClass) SetClassName(v string) *UdfClass {
	s.ClassName = &v
	return s
}

func (s *UdfClass) SetClassType(v string) *UdfClass {
	s.ClassType = &v
	return s
}

func (s *UdfClass) SetFunctionNames(v []*string) *UdfClass {
	s.FunctionNames = v
	return s
}

func (s *UdfClass) SetUdfArtifactName(v string) *UdfClass {
	s.UdfArtifactName = &v
	return s
}

type UdfFunction struct {
	ClassName       *string `json:"className,omitempty" xml:"className,omitempty"`
	FunctionName    *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	UdfArtifactName *string `json:"udfArtifactName,omitempty" xml:"udfArtifactName,omitempty"`
}

func (s UdfFunction) String() string {
	return tea.Prettify(s)
}

func (s UdfFunction) GoString() string {
	return s.String()
}

func (s *UdfFunction) SetClassName(v string) *UdfFunction {
	s.ClassName = &v
	return s
}

func (s *UdfFunction) SetFunctionName(v string) *UdfFunction {
	s.FunctionName = &v
	return s
}

func (s *UdfFunction) SetUdfArtifactName(v string) *UdfFunction {
	s.UdfArtifactName = &v
	return s
}

type UpdateJobConfigParam struct {
	NewFlinkConf map[string]interface{} `json:"newFlinkConf,omitempty" xml:"newFlinkConf,omitempty"`
}

func (s UpdateJobConfigParam) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobConfigParam) GoString() string {
	return s.String()
}

func (s *UpdateJobConfigParam) SetNewFlinkConf(v map[string]interface{}) *UpdateJobConfigParam {
	s.NewFlinkConf = v
	return s
}

type UpdateUdfArtifactResult struct {
	CollidingClasses []*UdfClass  `json:"collidingClasses,omitempty" xml:"collidingClasses,omitempty" type:"Repeated"`
	Message          *string      `json:"message,omitempty" xml:"message,omitempty"`
	MissingClasses   []*UdfClass  `json:"missingClasses,omitempty" xml:"missingClasses,omitempty" type:"Repeated"`
	UdfArtifact      *UdfArtifact `json:"udfArtifact,omitempty" xml:"udfArtifact,omitempty"`
	UpdateSuccess    *bool        `json:"updateSuccess,omitempty" xml:"updateSuccess,omitempty"`
}

func (s UpdateUdfArtifactResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateUdfArtifactResult) GoString() string {
	return s.String()
}

func (s *UpdateUdfArtifactResult) SetCollidingClasses(v []*UdfClass) *UpdateUdfArtifactResult {
	s.CollidingClasses = v
	return s
}

func (s *UpdateUdfArtifactResult) SetMessage(v string) *UpdateUdfArtifactResult {
	s.Message = &v
	return s
}

func (s *UpdateUdfArtifactResult) SetMissingClasses(v []*UdfClass) *UpdateUdfArtifactResult {
	s.MissingClasses = v
	return s
}

func (s *UpdateUdfArtifactResult) SetUdfArtifact(v *UdfArtifact) *UpdateUdfArtifactResult {
	s.UdfArtifact = v
	return s
}

func (s *UpdateUdfArtifactResult) SetUpdateSuccess(v bool) *UpdateUdfArtifactResult {
	s.UpdateSuccess = &v
	return s
}

type ValidateStatementResult struct {
	ErrorDetails *ValidationErrorDetails `json:"errorDetails,omitempty" xml:"errorDetails,omitempty"`
	// example:
	//
	// "there have some errors""
	ValidationResult *string `json:"validationResult,omitempty" xml:"validationResult,omitempty"`
}

func (s ValidateStatementResult) String() string {
	return tea.Prettify(s)
}

func (s ValidateStatementResult) GoString() string {
	return s.String()
}

func (s *ValidateStatementResult) SetErrorDetails(v *ValidationErrorDetails) *ValidateStatementResult {
	s.ErrorDetails = v
	return s
}

func (s *ValidateStatementResult) SetValidationResult(v string) *ValidateStatementResult {
	s.ValidationResult = &v
	return s
}

type ValidationErrorDetails struct {
	ColumnNumber    *string `json:"columnNumber,omitempty" xml:"columnNumber,omitempty"`
	EndColumnNumber *string `json:"endColumnNumber,omitempty" xml:"endColumnNumber,omitempty"`
	EndLineNumber   *string `json:"endLineNumber,omitempty" xml:"endLineNumber,omitempty"`
	LineNumber      *string `json:"lineNumber,omitempty" xml:"lineNumber,omitempty"`
	Message         *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ValidationErrorDetails) String() string {
	return tea.Prettify(s)
}

func (s ValidationErrorDetails) GoString() string {
	return s.String()
}

func (s *ValidationErrorDetails) SetColumnNumber(v string) *ValidationErrorDetails {
	s.ColumnNumber = &v
	return s
}

func (s *ValidationErrorDetails) SetEndColumnNumber(v string) *ValidationErrorDetails {
	s.EndColumnNumber = &v
	return s
}

func (s *ValidationErrorDetails) SetEndLineNumber(v string) *ValidationErrorDetails {
	s.EndLineNumber = &v
	return s
}

func (s *ValidationErrorDetails) SetLineNumber(v string) *ValidationErrorDetails {
	s.LineNumber = &v
	return s
}

func (s *ValidationErrorDetails) SetMessage(v string) *ValidationErrorDetails {
	s.Message = &v
	return s
}

type Variable struct {
	// example:
	//
	// This is a variable description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Plain
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// variableName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// variableValue
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
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

type WatermarkSpec struct {
	Column              *string `json:"column,omitempty" xml:"column,omitempty"`
	WatermarkExpression *string `json:"watermarkExpression,omitempty" xml:"watermarkExpression,omitempty"`
}

func (s WatermarkSpec) String() string {
	return tea.Prettify(s)
}

func (s WatermarkSpec) GoString() string {
	return s.String()
}

func (s *WatermarkSpec) SetColumn(v string) *WatermarkSpec {
	s.Column = &v
	return s
}

func (s *WatermarkSpec) SetWatermarkExpression(v string) *WatermarkSpec {
	s.WatermarkExpression = &v
	return s
}

type CreateDeploymentHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// The content of the deployment.
	//
	// This parameter is required.
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
	// 	- If the value of success was true, the deployment that you created was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data *Deployment `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateMemberHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca84d539167d4d
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateMemberHeaders) GoString() string {
	return s.String()
}

func (s *CreateMemberHeaders) SetCommonHeaders(v map[string]*string) *CreateMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateMemberHeaders) SetWorkspace(v string) *CreateMemberHeaders {
	s.Workspace = &v
	return s
}

type CreateMemberRequest struct {
	// The mappings between the ID and permissions of the member.
	Body *Member `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMemberRequest) GoString() string {
	return s.String()
}

func (s *CreateMemberRequest) SetBody(v *Member) *CreateMemberRequest {
	s.Body = v
	return s
}

type CreateMemberResponseBody struct {
	// 	- If the value of success was false, a null value was returned.
	//
	// 	- If the value of success was true, the authorization information was returned.
	Data *Member `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The status code returned. The value was fixed to 200. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F989CA70-2925-5A94-92B7-20F5762B71C8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMemberResponseBody) SetData(v *Member) *CreateMemberResponseBody {
	s.Data = v
	return s
}

func (s *CreateMemberResponseBody) SetErrorCode(v string) *CreateMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateMemberResponseBody) SetErrorMessage(v string) *CreateMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateMemberResponseBody) SetHttpCode(v int32) *CreateMemberResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateMemberResponseBody) SetRequestId(v string) *CreateMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMemberResponseBody) SetSuccess(v bool) *CreateMemberResponseBody {
	s.Success = &v
	return s
}

type CreateMemberResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMemberResponse) GoString() string {
	return s.String()
}

func (s *CreateMemberResponse) SetHeaders(v map[string]*string) *CreateMemberResponse {
	s.Headers = v
	return s
}

func (s *CreateMemberResponse) SetStatusCode(v int32) *CreateMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMemberResponse) SetBody(v *CreateMemberResponseBody) *CreateMemberResponse {
	s.Body = v
	return s
}

type CreateSavepointHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// The deployment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 58718c99-3b29-4c5e-93bb-c9fc4ec6****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// The description of the savepoint.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to use the native format mode. Valid values:
	//
	// 	- true: The native format mode is used.
	//
	// 	- false: The native format mode is not used.
	//
	// example:
	//
	// true
	NativeFormat *bool `json:"nativeFormat,omitempty" xml:"nativeFormat,omitempty"`
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
	// 	- If the value of success was true, the savepoint that was created was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data *Savepoint `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSavepointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// The parameter that is used to create the variable.
	//
	// This parameter is required.
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
	// 	- If the value of success was true, the variable that you created was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data *Variable `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-ABCD-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DeleteMemberHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 710d6a64d8c34d
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteMemberHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMemberHeaders) SetCommonHeaders(v map[string]*string) *DeleteMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMemberHeaders) SetWorkspace(v string) *DeleteMemberHeaders {
	s.Workspace = &v
	return s
}

type DeleteMemberResponseBody struct {
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The status code returned. The value was fixed to 200. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMemberResponseBody) SetErrorCode(v string) *DeleteMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteMemberResponseBody) SetErrorMessage(v string) *DeleteMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteMemberResponseBody) SetHttpCode(v int32) *DeleteMemberResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteMemberResponseBody) SetRequestId(v string) *DeleteMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMemberResponseBody) SetSuccess(v bool) *DeleteMemberResponseBody {
	s.Success = &v
	return s
}

type DeleteMemberResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteMemberResponse) SetHeaders(v map[string]*string) *DeleteMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteMemberResponse) SetStatusCode(v int32) *DeleteMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMemberResponse) SetBody(v *DeleteMemberResponseBody) *DeleteMemberResponse {
	s.Body = v
	return s
}

type DeleteSavepointHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSavepointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// The path of the Flink UI.
	//
	// This parameter is required.
	//
	// example:
	//
	// /jobs/4df35f8e54554b23bf7dcd38a151****
	FlinkApiPath *string `json:"flinkApiPath,omitempty" xml:"flinkApiPath,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// default-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5a27a3aa-c5b9-4dc1-8c86-be57d2d6****
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- jobs
	//
	// 	- sessionclusters
	//
	// This parameter is required.
	//
	// example:
	//
	// jobs
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
	// 	- If the value of success was true, the result of the proxy request was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	//
	// example:
	//
	// { "jobs": [ { "jid": "4df35f8e54554b23bf7dcd38a151****", "name": "69d001d5-419a-4bfc-9c2e-849cacd3****", "state": "RUNNING", "start-time": 1659154942068, "end-time": -1, "duration": 188161756, "last-modification": 1659154968305, "tasks": { "total": 2, "created": 0, "scheduled": 0, "deploying": 0, "running": 2, "finished": 0, "canceling": 0, "canceled": 0, "failed": 0, "reconciling": 0, "initializing": 0 } } ] }
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlinkApiProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// The Flink configuration that is used to generate a resource plan.
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
	// 	- If the value of success was true, the asynchronous generation result was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data *GenerateResourcePlanWithFlinkConfAsyncResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	// The ID of the ticket for you to query the asynchronous generation result.
	//
	// example:
	//
	// b3dcdb25-bf36-457d-92ba-a36077e8****
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
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateResourcePlanWithFlinkConfAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// 	- If the value of success was true, the details of the deployment were returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data *Deployment `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// 	- If the value of success was true, the asynchronous generation result was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data *AsyncResourcePlanOperationResult `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGenerateResourcePlanResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// 	- If the value of success was true, the details of the job was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data *Job `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetLatestJobStartLogHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetLatestJobStartLogHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetLatestJobStartLogHeaders) GoString() string {
	return s.String()
}

func (s *GetLatestJobStartLogHeaders) SetCommonHeaders(v map[string]*string) *GetLatestJobStartLogHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetLatestJobStartLogHeaders) SetWorkspace(v string) *GetLatestJobStartLogHeaders {
	s.Workspace = &v
	return s
}

type GetLatestJobStartLogResponseBody struct {
	// If the value of success was false, the latest logs of the deployment were returned. If the value of success was true, a null value was returned.
	//
	// example:
	//
	// "[main] INFO  org.apache.flink.runtime.entrypoint.ClusterEntrypoint        [] - --------------------------------------------------------------------------------\\n2024-05-22 11:46:39,871 [main] INFO  org.apache.flink.runtime.entrypoint.ClusterEntrypoint"
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// If the value of success was false, an error code was returned. If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// If the value of success was false, an error message was returned. If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The status code returned. The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetLatestJobStartLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLatestJobStartLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetLatestJobStartLogResponseBody) SetData(v string) *GetLatestJobStartLogResponseBody {
	s.Data = &v
	return s
}

func (s *GetLatestJobStartLogResponseBody) SetErrorCode(v string) *GetLatestJobStartLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetLatestJobStartLogResponseBody) SetErrorMessage(v string) *GetLatestJobStartLogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetLatestJobStartLogResponseBody) SetHttpCode(v int32) *GetLatestJobStartLogResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetLatestJobStartLogResponseBody) SetRequestId(v string) *GetLatestJobStartLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLatestJobStartLogResponseBody) SetSuccess(v bool) *GetLatestJobStartLogResponseBody {
	s.Success = &v
	return s
}

type GetLatestJobStartLogResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLatestJobStartLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLatestJobStartLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLatestJobStartLogResponse) GoString() string {
	return s.String()
}

func (s *GetLatestJobStartLogResponse) SetHeaders(v map[string]*string) *GetLatestJobStartLogResponse {
	s.Headers = v
	return s
}

func (s *GetLatestJobStartLogResponse) SetStatusCode(v int32) *GetLatestJobStartLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLatestJobStartLogResponse) SetBody(v *GetLatestJobStartLogResponseBody) *GetLatestJobStartLogResponse {
	s.Body = v
	return s
}

type GetMemberHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMemberHeaders) GoString() string {
	return s.String()
}

func (s *GetMemberHeaders) SetCommonHeaders(v map[string]*string) *GetMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMemberHeaders) SetWorkspace(v string) *GetMemberHeaders {
	s.Workspace = &v
	return s
}

type GetMemberResponseBody struct {
	// 	- If the value of success was false, a null value was returned.
	//
	// 	- If the value of success was true, the authorization information was returned.
	Data *Member `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The status code returned. The value was fixed to 200. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMemberResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemberResponseBody) SetData(v *Member) *GetMemberResponseBody {
	s.Data = v
	return s
}

func (s *GetMemberResponseBody) SetErrorCode(v string) *GetMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMemberResponseBody) SetErrorMessage(v string) *GetMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMemberResponseBody) SetHttpCode(v int32) *GetMemberResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetMemberResponseBody) SetRequestId(v string) *GetMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemberResponseBody) SetSuccess(v bool) *GetMemberResponseBody {
	s.Success = &v
	return s
}

type GetMemberResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMemberResponse) GoString() string {
	return s.String()
}

func (s *GetMemberResponse) SetHeaders(v map[string]*string) *GetMemberResponse {
	s.Headers = v
	return s
}

func (s *GetMemberResponse) SetStatusCode(v int32) *GetMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemberResponse) SetBody(v *GetMemberResponseBody) *GetMemberResponse {
	s.Body = v
	return s
}

type GetSavepointHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// 	- If the value of success was true, the savepoint information was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data *Savepoint `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSavepointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// The page number. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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
	// 	- If the value of success was true, a list of clusters in which the deployment is deployed was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data []*DeploymentTarget `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeploymentTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// The ID of the user who creates the deployment.
	//
	// example:
	//
	// 183899668*******
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The execution mode of the deployment.
	//
	// Valid values:
	//
	// 	- BATCH
	//
	// 	- STREAMING
	//
	// example:
	//
	// STREAMING
	ExecutionMode *string `json:"executionMode,omitempty" xml:"executionMode,omitempty"`
	// The tag key.
	//
	// example:
	//
	// key
	LabelKey *string `json:"labelKey,omitempty" xml:"labelKey,omitempty"`
	// The tag value. Separate multiple values with semicolon (;).
	//
	// example:
	//
	// value1,value2
	LabelValueArray *string `json:"labelValueArray,omitempty" xml:"labelValueArray,omitempty"`
	// The ID of the user who modifies the deployment.
	//
	// example:
	//
	// 183899668*******
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// The name of the deployment.
	//
	// example:
	//
	// vvp_ds_0522
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The page number. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The latest status of the deployment.
	//
	// Valid values:
	//
	// 	- CANCELLED
	//
	// 	- FAILED
	//
	// 	- RUNNING
	//
	// 	- TRANSITIONING
	//
	// 	- FINISHED
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListDeploymentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentsRequest) SetCreator(v string) *ListDeploymentsRequest {
	s.Creator = &v
	return s
}

func (s *ListDeploymentsRequest) SetExecutionMode(v string) *ListDeploymentsRequest {
	s.ExecutionMode = &v
	return s
}

func (s *ListDeploymentsRequest) SetLabelKey(v string) *ListDeploymentsRequest {
	s.LabelKey = &v
	return s
}

func (s *ListDeploymentsRequest) SetLabelValueArray(v string) *ListDeploymentsRequest {
	s.LabelValueArray = &v
	return s
}

func (s *ListDeploymentsRequest) SetModifier(v string) *ListDeploymentsRequest {
	s.Modifier = &v
	return s
}

func (s *ListDeploymentsRequest) SetName(v string) *ListDeploymentsRequest {
	s.Name = &v
	return s
}

func (s *ListDeploymentsRequest) SetPageIndex(v int32) *ListDeploymentsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDeploymentsRequest) SetPageSize(v int32) *ListDeploymentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentsRequest) SetStatus(v string) *ListDeploymentsRequest {
	s.Status = &v
	return s
}

type ListDeploymentsResponseBody struct {
	// 	- If the value of success was true, the list of all deployments was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data []*Deployment `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeploymentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListEditableNamespaceRequest struct {
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	PageIndex *string `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	PageSize  *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	RegionId    *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListEditableNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEditableNamespaceRequest) GoString() string {
	return s.String()
}

func (s *ListEditableNamespaceRequest) SetNamespace(v string) *ListEditableNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *ListEditableNamespaceRequest) SetPageIndex(v string) *ListEditableNamespaceRequest {
	s.PageIndex = &v
	return s
}

func (s *ListEditableNamespaceRequest) SetPageSize(v string) *ListEditableNamespaceRequest {
	s.PageSize = &v
	return s
}

func (s *ListEditableNamespaceRequest) SetRegionId(v string) *ListEditableNamespaceRequest {
	s.RegionId = &v
	return s
}

func (s *ListEditableNamespaceRequest) SetWorkspaceId(v string) *ListEditableNamespaceRequest {
	s.WorkspaceId = &v
	return s
}

type ListEditableNamespaceResponseBody struct {
	Data      *ListEditableNamespaceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	HttpCode  *int32                                 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	Message   *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Reason    *string                                `json:"reason,omitempty" xml:"reason,omitempty"`
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListEditableNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEditableNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *ListEditableNamespaceResponseBody) SetData(v *ListEditableNamespaceResponseBodyData) *ListEditableNamespaceResponseBody {
	s.Data = v
	return s
}

func (s *ListEditableNamespaceResponseBody) SetHttpCode(v int32) *ListEditableNamespaceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListEditableNamespaceResponseBody) SetMessage(v string) *ListEditableNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *ListEditableNamespaceResponseBody) SetReason(v string) *ListEditableNamespaceResponseBody {
	s.Reason = &v
	return s
}

func (s *ListEditableNamespaceResponseBody) SetRequestId(v string) *ListEditableNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEditableNamespaceResponseBody) SetSuccess(v bool) *ListEditableNamespaceResponseBody {
	s.Success = &v
	return s
}

type ListEditableNamespaceResponseBodyData struct {
	EditableNamespaces []*EditableNamespace `json:"editableNamespaces,omitempty" xml:"editableNamespaces,omitempty" type:"Repeated"`
	PageIndex          *string              `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	PageSize           *string              `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total              *string              `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListEditableNamespaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEditableNamespaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEditableNamespaceResponseBodyData) SetEditableNamespaces(v []*EditableNamespace) *ListEditableNamespaceResponseBodyData {
	s.EditableNamespaces = v
	return s
}

func (s *ListEditableNamespaceResponseBodyData) SetPageIndex(v string) *ListEditableNamespaceResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListEditableNamespaceResponseBodyData) SetPageSize(v string) *ListEditableNamespaceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListEditableNamespaceResponseBodyData) SetTotal(v string) *ListEditableNamespaceResponseBodyData {
	s.Total = &v
	return s
}

type ListEditableNamespaceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEditableNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEditableNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEditableNamespaceResponse) GoString() string {
	return s.String()
}

func (s *ListEditableNamespaceResponse) SetHeaders(v map[string]*string) *ListEditableNamespaceResponse {
	s.Headers = v
	return s
}

func (s *ListEditableNamespaceResponse) SetStatusCode(v int32) *ListEditableNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEditableNamespaceResponse) SetBody(v *ListEditableNamespaceResponseBody) *ListEditableNamespaceResponse {
	s.Body = v
	return s
}

type ListEngineVersionMetadataHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// 	- If the value of success was true, the engine versions that are supported by Realtime Compute for Apache Flink were returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data *EngineVersionMetadataIndex `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEngineVersionMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// The deployment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 58718c99-3b29-4c5e-93bb-c9fc4ec6****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// The page number. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The collation.
	//
	// Valid values:
	//
	// 	- gmt_create
	//
	// 	- job_id
	//
	// 	- status
	//
	// example:
	//
	// gmt_create
	SortName *string `json:"sortName,omitempty" xml:"sortName,omitempty"`
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

func (s *ListJobsRequest) SetSortName(v string) *ListJobsRequest {
	s.SortName = &v
	return s
}

type ListJobsResponseBody struct {
	// 	- If the value of success was true, all jobs that meet the condition were returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data []*Job `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListMembersHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListMembersHeaders) GoString() string {
	return s.String()
}

func (s *ListMembersHeaders) SetCommonHeaders(v map[string]*string) *ListMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListMembersHeaders) SetWorkspace(v string) *ListMembersHeaders {
	s.Workspace = &v
	return s
}

type ListMembersRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMembersRequest) GoString() string {
	return s.String()
}

func (s *ListMembersRequest) SetPageIndex(v int32) *ListMembersRequest {
	s.PageIndex = &v
	return s
}

func (s *ListMembersRequest) SetPageSize(v int32) *ListMembersRequest {
	s.PageSize = &v
	return s
}

type ListMembersResponseBody struct {
	// 	- If the value of success was false, a null value was returned.
	//
	// 	- If the value of success was true, the authorization information was returned.
	Data []*Member `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The status code returned. The value was fixed to 200. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 50
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListMembersResponseBody) SetData(v []*Member) *ListMembersResponseBody {
	s.Data = v
	return s
}

func (s *ListMembersResponseBody) SetErrorCode(v string) *ListMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListMembersResponseBody) SetErrorMessage(v string) *ListMembersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListMembersResponseBody) SetHttpCode(v int32) *ListMembersResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListMembersResponseBody) SetPageIndex(v int32) *ListMembersResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListMembersResponseBody) SetPageSize(v int32) *ListMembersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListMembersResponseBody) SetRequestId(v string) *ListMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMembersResponseBody) SetSuccess(v bool) *ListMembersResponseBody {
	s.Success = &v
	return s
}

func (s *ListMembersResponseBody) SetTotalSize(v int32) *ListMembersResponseBody {
	s.TotalSize = &v
	return s
}

type ListMembersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMembersResponse) GoString() string {
	return s.String()
}

func (s *ListMembersResponse) SetHeaders(v map[string]*string) *ListMembersResponse {
	s.Headers = v
	return s
}

func (s *ListMembersResponse) SetStatusCode(v int32) *ListMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMembersResponse) SetBody(v *ListMembersResponseBody) *ListMembersResponse {
	s.Body = v
	return s
}

type ListSavepointsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// The deployment ID. This parameter is optional.
	//
	// example:
	//
	// 88a8fc49-e090-430a-85d8-3ee8c79c****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// The job ID. This parameter is optional.
	//
	// example:
	//
	// 99a8fc49-e090-430a-85d8-3ee8c79c****
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// The page number. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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
	// 	- If the value of success was true, a list of savepoints was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data []*Savepoint `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSavepointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// The page number. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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
	// 	- If the value of success was true, a list of variables was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data []*Variable `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-ABCF-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVariablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// The parameter that is used to start the job.
	//
	// This parameter is required.
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
	// 	- If the value of success was true, the job that you created was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data *Job `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type StartJobWithParamsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s StartJobWithParamsHeaders) String() string {
	return tea.Prettify(s)
}

func (s StartJobWithParamsHeaders) GoString() string {
	return s.String()
}

func (s *StartJobWithParamsHeaders) SetCommonHeaders(v map[string]*string) *StartJobWithParamsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartJobWithParamsHeaders) SetWorkspace(v string) *StartJobWithParamsHeaders {
	s.Workspace = &v
	return s
}

type StartJobWithParamsRequest struct {
	// The parameter that is used to start the job.
	Body *JobStartParameters `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartJobWithParamsRequest) String() string {
	return tea.Prettify(s)
}

func (s StartJobWithParamsRequest) GoString() string {
	return s.String()
}

func (s *StartJobWithParamsRequest) SetBody(v *JobStartParameters) *StartJobWithParamsRequest {
	s.Body = v
	return s
}

type StartJobWithParamsResponseBody struct {
	// The details of the job of the deployment returned.
	Data *Job `json:"data,omitempty" xml:"data,omitempty"`
	// If the value of success was false, an error code was returned. If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// If the value of success was false, an error message was returned. If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The status code returned. The value was fixed to 200. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StartJobWithParamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartJobWithParamsResponseBody) GoString() string {
	return s.String()
}

func (s *StartJobWithParamsResponseBody) SetData(v *Job) *StartJobWithParamsResponseBody {
	s.Data = v
	return s
}

func (s *StartJobWithParamsResponseBody) SetErrorCode(v string) *StartJobWithParamsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StartJobWithParamsResponseBody) SetErrorMessage(v string) *StartJobWithParamsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StartJobWithParamsResponseBody) SetHttpCode(v int32) *StartJobWithParamsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *StartJobWithParamsResponseBody) SetRequestId(v string) *StartJobWithParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartJobWithParamsResponseBody) SetSuccess(v bool) *StartJobWithParamsResponseBody {
	s.Success = &v
	return s
}

type StartJobWithParamsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartJobWithParamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartJobWithParamsResponse) String() string {
	return tea.Prettify(s)
}

func (s StartJobWithParamsResponse) GoString() string {
	return s.String()
}

func (s *StartJobWithParamsResponse) SetHeaders(v map[string]*string) *StartJobWithParamsResponse {
	s.Headers = v
	return s
}

func (s *StartJobWithParamsResponse) SetStatusCode(v int32) *StartJobWithParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *StartJobWithParamsResponse) SetBody(v *StartJobWithParamsResponseBody) *StartJobWithParamsResponse {
	s.Body = v
	return s
}

type StopJobHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// The parameter that is used to stop the job.
	//
	// This parameter is required.
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
	// 	- If the value of success was true, the job that you stopped was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data *Job `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	// The information about the deployment that you want to update.
	//
	// This parameter is required.
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
	// 	- If the value of success was true, the information about the deployment after the update was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data *Deployment `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type UpdateMemberHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMemberHeaders) SetCommonHeaders(v map[string]*string) *UpdateMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMemberHeaders) SetWorkspace(v string) *UpdateMemberHeaders {
	s.Workspace = &v
	return s
}

type UpdateMemberRequest struct {
	// The permission information about the member.
	Body *Member `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemberRequest) SetBody(v *Member) *UpdateMemberRequest {
	s.Body = v
	return s
}

type UpdateMemberResponseBody struct {
	// If the value of success was true, the member that was created was returned. If the value of success was false, a null value was returned.
	Data *Member `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The status code returned. The value was fixed to 200. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemberResponseBody) SetData(v *Member) *UpdateMemberResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMemberResponseBody) SetErrorCode(v string) *UpdateMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateMemberResponseBody) SetErrorMessage(v string) *UpdateMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateMemberResponseBody) SetHttpCode(v int32) *UpdateMemberResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateMemberResponseBody) SetRequestId(v string) *UpdateMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMemberResponseBody) SetSuccess(v bool) *UpdateMemberResponseBody {
	s.Success = &v
	return s
}

type UpdateMemberResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateMemberResponse) SetHeaders(v map[string]*string) *UpdateMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateMemberResponse) SetStatusCode(v int32) *UpdateMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMemberResponse) SetBody(v *UpdateMemberResponseBody) *UpdateMemberResponse {
	s.Body = v
	return s
}

type ValidateSqlStatementHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ValidateSqlStatementHeaders) String() string {
	return tea.Prettify(s)
}

func (s ValidateSqlStatementHeaders) GoString() string {
	return s.String()
}

func (s *ValidateSqlStatementHeaders) SetCommonHeaders(v map[string]*string) *ValidateSqlStatementHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ValidateSqlStatementHeaders) SetWorkspace(v string) *ValidateSqlStatementHeaders {
	s.Workspace = &v
	return s
}

type ValidateSqlStatementRequest struct {
	// The content of the code that you want to verify.
	//
	// This parameter is required.
	Body *SqlStatementWithContext `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateSqlStatementRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateSqlStatementRequest) GoString() string {
	return s.String()
}

func (s *ValidateSqlStatementRequest) SetBody(v *SqlStatementWithContext) *ValidateSqlStatementRequest {
	s.Body = v
	return s
}

type ValidateSqlStatementResponseBody struct {
	// The returned data, which represents the details of SQL validation results.
	Data *SqlStatementValidationResult `json:"data,omitempty" xml:"data,omitempty"`
	// If the value of success was false, an error code was returned. If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// If the value of success was false, an error message was returned. If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The status code returned. The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-****-1D30-8A4F-882ED4DD5E02
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ValidateSqlStatementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateSqlStatementResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateSqlStatementResponseBody) SetData(v *SqlStatementValidationResult) *ValidateSqlStatementResponseBody {
	s.Data = v
	return s
}

func (s *ValidateSqlStatementResponseBody) SetErrorCode(v string) *ValidateSqlStatementResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ValidateSqlStatementResponseBody) SetErrorMessage(v string) *ValidateSqlStatementResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ValidateSqlStatementResponseBody) SetHttpCode(v int32) *ValidateSqlStatementResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ValidateSqlStatementResponseBody) SetRequestId(v string) *ValidateSqlStatementResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateSqlStatementResponseBody) SetSuccess(v bool) *ValidateSqlStatementResponseBody {
	s.Success = &v
	return s
}

type ValidateSqlStatementResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateSqlStatementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateSqlStatementResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateSqlStatementResponse) GoString() string {
	return s.String()
}

func (s *ValidateSqlStatementResponse) SetHeaders(v map[string]*string) *ValidateSqlStatementResponse {
	s.Headers = v
	return s
}

func (s *ValidateSqlStatementResponse) SetStatusCode(v int32) *ValidateSqlStatementResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateSqlStatementResponse) SetBody(v *ValidateSqlStatementResponseBody) *ValidateSqlStatementResponse {
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

// Summary:
//
// Creates a deployment.
//
// @param request - CreateDeploymentRequest
//
// @param headers - CreateDeploymentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeploymentResponse
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

// Summary:
//
// Creates a deployment.
//
// @param request - CreateDeploymentRequest
//
// @return CreateDeploymentResponse
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

// Summary:
//
// Adds a user to a namespace as a member and grants permissions to the user.
//
// @param request - CreateMemberRequest
//
// @param headers - CreateMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemberResponse
func (client *Client) CreateMemberWithOptions(namespace *string, request *CreateMemberRequest, headers *CreateMemberHeaders, runtime *util.RuntimeOptions) (_result *CreateMemberResponse, _err error) {
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
		Action:      tea.String("CreateMember"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gateway/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/members"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a user to a namespace as a member and grants permissions to the user.
//
// @param request - CreateMemberRequest
//
// @return CreateMemberResponse
func (client *Client) CreateMember(namespace *string, request *CreateMemberRequest) (_result *CreateMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateMemberHeaders{}
	_result = &CreateMemberResponse{}
	_body, _err := client.CreateMemberWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a savepoint.
//
// @param request - CreateSavepointRequest
//
// @param headers - CreateSavepointHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSavepointResponse
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

// Summary:
//
// Creates a savepoint.
//
// @param request - CreateSavepointRequest
//
// @return CreateSavepointResponse
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

// Summary:
//
// Creates a variable.
//
// @param request - CreateVariableRequest
//
// @param headers - CreateVariableHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVariableResponse
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

// Summary:
//
// Creates a variable.
//
// @param request - CreateVariableRequest
//
// @return CreateVariableResponse
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

// Summary:
//
// Deletes a deployment based on the deployment ID.
//
// @param headers - DeleteDeploymentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeploymentResponse
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

// Summary:
//
// Deletes a deployment based on the deployment ID.
//
// @return DeleteDeploymentResponse
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

// Summary:
//
// Deletes the information about a job that is not in the running state in a deployment.
//
// @param headers - DeleteJobHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteJobResponse
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

// Summary:
//
// Deletes the information about a job that is not in the running state in a deployment.
//
// @return DeleteJobResponse
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

// Summary:
//
// Revokes the permissions from a member.
//
// @param headers - DeleteMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemberResponse
func (client *Client) DeleteMemberWithOptions(namespace *string, member *string, headers *DeleteMemberHeaders, runtime *util.RuntimeOptions) (_result *DeleteMemberResponse, _err error) {
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
		Action:      tea.String("DeleteMember"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gateway/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/members/" + tea.StringValue(openapiutil.GetEncodeParam(member))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the permissions from a member.
//
// @return DeleteMemberResponse
func (client *Client) DeleteMember(namespace *string, member *string) (_result *DeleteMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteMemberHeaders{}
	_result = &DeleteMemberResponse{}
	_body, _err := client.DeleteMemberWithOptions(namespace, member, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a savepoint.
//
// @param headers - DeleteSavepointHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSavepointResponse
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

// Summary:
//
// Deletes a savepoint.
//
// @return DeleteSavepointResponse
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

// Summary:
//
// Deletes a variable.
//
// @param headers - DeleteVariableHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVariableResponse
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

// Summary:
//
// Deletes a variable.
//
// @return DeleteVariableResponse
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

// Summary:
//
// Provides a Flink request proxy.
//
// @param request - FlinkApiProxyRequest
//
// @param headers - FlinkApiProxyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlinkApiProxyResponse
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

// Summary:
//
// Provides a Flink request proxy.
//
// @param request - FlinkApiProxyRequest
//
// @return FlinkApiProxyResponse
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

// Summary:
//
// Submits a ticket that applies for asynchronous generation of the fine-grained resources. This operation returns the ID of the ticket for you to query the asynchronous generation result.
//
// @param request - GenerateResourcePlanWithFlinkConfAsyncRequest
//
// @param headers - GenerateResourcePlanWithFlinkConfAsyncHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateResourcePlanWithFlinkConfAsyncResponse
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

// Summary:
//
// Submits a ticket that applies for asynchronous generation of the fine-grained resources. This operation returns the ID of the ticket for you to query the asynchronous generation result.
//
// @param request - GenerateResourcePlanWithFlinkConfAsyncRequest
//
// @return GenerateResourcePlanWithFlinkConfAsyncResponse
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

// Summary:
//
// Obtains the details of a deployment.
//
// @param headers - GetDeploymentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeploymentResponse
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

// Summary:
//
// Obtains the details of a deployment.
//
// @return GetDeploymentResponse
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

// Summary:
//
// Obtains the asynchronous generation result of fine-grained resources based on the ID of the ticket that applies for an asynchronous generation.
//
// @param headers - GetGenerateResourcePlanResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGenerateResourcePlanResultResponse
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

// Summary:
//
// Obtains the asynchronous generation result of fine-grained resources based on the ID of the ticket that applies for an asynchronous generation.
//
// @return GetGenerateResourcePlanResultResponse
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

// Summary:
//
// Obtains the details of a job.
//
// @param headers - GetJobHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResponse
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

// Summary:
//
// Obtains the details of a job.
//
// @return GetJobResponse
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

// Summary:
//
// Obtains the latest startup logs of a job.
//
// @param headers - GetLatestJobStartLogHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLatestJobStartLogResponse
func (client *Client) GetLatestJobStartLogWithOptions(namespace *string, deploymentId *string, headers *GetLatestJobStartLogHeaders, runtime *util.RuntimeOptions) (_result *GetLatestJobStartLogResponse, _err error) {
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
		Action:      tea.String("GetLatestJobStartLog"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployments/" + tea.StringValue(openapiutil.GetEncodeParam(deploymentId)) + "/latest_jobmanager_start_log"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLatestJobStartLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the latest startup logs of a job.
//
// @return GetLatestJobStartLogResponse
func (client *Client) GetLatestJobStartLog(namespace *string, deploymentId *string) (_result *GetLatestJobStartLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetLatestJobStartLogHeaders{}
	_result = &GetLatestJobStartLogResponse{}
	_body, _err := client.GetLatestJobStartLogWithOptions(namespace, deploymentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the permissions of a member.
//
// @param headers - GetMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemberResponse
func (client *Client) GetMemberWithOptions(namespace *string, member *string, headers *GetMemberHeaders, runtime *util.RuntimeOptions) (_result *GetMemberResponse, _err error) {
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
		Action:      tea.String("GetMember"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gateway/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/members/" + tea.StringValue(openapiutil.GetEncodeParam(member))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the permissions of a member.
//
// @return GetMemberResponse
func (client *Client) GetMember(namespace *string, member *string) (_result *GetMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMemberHeaders{}
	_result = &GetMemberResponse{}
	_body, _err := client.GetMemberWithOptions(namespace, member, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries details of a savepoint and checkpoint.
//
// @param headers - GetSavepointHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSavepointResponse
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

// Summary:
//
// Queries details of a savepoint and checkpoint.
//
// @return GetSavepointResponse
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

// Summary:
//
// Obtains a list of clusters in which deployments can be deployed. The cluster can be a session cluster or a per-job cluster.
//
// @param request - ListDeploymentTargetsRequest
//
// @param headers - ListDeploymentTargetsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentTargetsResponse
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

// Summary:
//
// Obtains a list of clusters in which deployments can be deployed. The cluster can be a session cluster or a per-job cluster.
//
// @param request - ListDeploymentTargetsRequest
//
// @return ListDeploymentTargetsResponse
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

// Summary:
//
// Obtains information about all deployments.
//
// @param request - ListDeploymentsRequest
//
// @param headers - ListDeploymentsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentsResponse
func (client *Client) ListDeploymentsWithOptions(namespace *string, request *ListDeploymentsRequest, headers *ListDeploymentsHeaders, runtime *util.RuntimeOptions) (_result *ListDeploymentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Creator)) {
		query["creator"] = request.Creator
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionMode)) {
		query["executionMode"] = request.ExecutionMode
	}

	if !tea.BoolValue(util.IsUnset(request.LabelKey)) {
		query["labelKey"] = request.LabelKey
	}

	if !tea.BoolValue(util.IsUnset(request.LabelValueArray)) {
		query["labelValueArray"] = request.LabelValueArray
	}

	if !tea.BoolValue(util.IsUnset(request.Modifier)) {
		query["modifier"] = request.Modifier
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["pageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
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

// Summary:
//
// Obtains information about all deployments.
//
// @param request - ListDeploymentsRequest
//
// @return ListDeploymentsResponse
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

// Summary:
//
// 列出有编辑权限的项目空间。
//
// @param request - ListEditableNamespaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEditableNamespaceResponse
func (client *Client) ListEditableNamespaceWithOptions(request *ListEditableNamespaceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEditableNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["pageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEditableNamespace"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gateway/v2/namespaces/editable"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEditableNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出有编辑权限的项目空间。
//
// @param request - ListEditableNamespaceRequest
//
// @return ListEditableNamespaceResponse
func (client *Client) ListEditableNamespace(request *ListEditableNamespaceRequest) (_result *ListEditableNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEditableNamespaceResponse{}
	_body, _err := client.ListEditableNamespaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a list of engine versions that are supported by Realtime Compute for Apache Flink.
//
// @param headers - ListEngineVersionMetadataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEngineVersionMetadataResponse
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

// Summary:
//
// Obtains a list of engine versions that are supported by Realtime Compute for Apache Flink.
//
// @return ListEngineVersionMetadataResponse
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

// Summary:
//
// Queries the information about all jobs in a deployment.
//
// @param request - ListJobsRequest
//
// @param headers - ListJobsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsResponse
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

	if !tea.BoolValue(util.IsUnset(request.SortName)) {
		query["sortName"] = request.SortName
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

// Summary:
//
// Queries the information about all jobs in a deployment.
//
// @param request - ListJobsRequest
//
// @return ListJobsResponse
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

// Summary:
//
// Queries the mappings between the ID and permissions of a member in a specific namespace.
//
// @param request - ListMembersRequest
//
// @param headers - ListMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMembersResponse
func (client *Client) ListMembersWithOptions(namespace *string, request *ListMembersRequest, headers *ListMembersHeaders, runtime *util.RuntimeOptions) (_result *ListMembersResponse, _err error) {
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
		Action:      tea.String("ListMembers"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gateway/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/members"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the mappings between the ID and permissions of a member in a specific namespace.
//
// @param request - ListMembersRequest
//
// @return ListMembersResponse
func (client *Client) ListMembers(namespace *string, request *ListMembersRequest) (_result *ListMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListMembersHeaders{}
	_result = &ListMembersResponse{}
	_body, _err := client.ListMembersWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a list of savepoints or checkpoints.
//
// @param request - ListSavepointsRequest
//
// @param headers - ListSavepointsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSavepointsResponse
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

// Summary:
//
// Obtains a list of savepoints or checkpoints.
//
// @param request - ListSavepointsRequest
//
// @return ListSavepointsResponse
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

// Summary:
//
// Obtains a list of variables.
//
// @param request - ListVariablesRequest
//
// @param headers - ListVariablesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVariablesResponse
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

// Summary:
//
// Obtains a list of variables.
//
// @param request - ListVariablesRequest
//
// @return ListVariablesResponse
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

// Deprecated: OpenAPI StartJob is deprecated
//
// Summary:
//
// Creates and starts a job.
//
// @param request - StartJobRequest
//
// @param headers - StartJobHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartJobResponse
// Deprecated
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

// Deprecated: OpenAPI StartJob is deprecated
//
// Summary:
//
// Creates and starts a job.
//
// @param request - StartJobRequest
//
// @return StartJobResponse
// Deprecated
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

// Summary:
//
// Starts a job.
//
// @param request - StartJobWithParamsRequest
//
// @param headers - StartJobWithParamsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartJobWithParamsResponse
func (client *Client) StartJobWithParamsWithOptions(namespace *string, request *StartJobWithParamsRequest, headers *StartJobWithParamsHeaders, runtime *util.RuntimeOptions) (_result *StartJobWithParamsResponse, _err error) {
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
		Action:      tea.String("StartJobWithParams"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/jobs%3Astart"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartJobWithParamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a job.
//
// @param request - StartJobWithParamsRequest
//
// @return StartJobWithParamsResponse
func (client *Client) StartJobWithParams(namespace *string, request *StartJobWithParamsRequest) (_result *StartJobWithParamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StartJobWithParamsHeaders{}
	_result = &StartJobWithParamsResponse{}
	_body, _err := client.StartJobWithParamsWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a job.
//
// @param request - StopJobRequest
//
// @param headers - StopJobHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopJobResponse
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

// Summary:
//
// Stops a job.
//
// @param request - StopJobRequest
//
// @return StopJobResponse
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

// Summary:
//
// Updates information about a deployment.
//
// @param request - UpdateDeploymentRequest
//
// @param headers - UpdateDeploymentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeploymentResponse
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

// Summary:
//
// Updates information about a deployment.
//
// @param request - UpdateDeploymentRequest
//
// @return UpdateDeploymentResponse
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

// Summary:
//
// Updates the permissions of one or more members in a specific namespace.
//
// @param request - UpdateMemberRequest
//
// @param headers - UpdateMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemberResponse
func (client *Client) UpdateMemberWithOptions(namespace *string, request *UpdateMemberRequest, headers *UpdateMemberHeaders, runtime *util.RuntimeOptions) (_result *UpdateMemberResponse, _err error) {
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
		Action:      tea.String("UpdateMember"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gateway/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/members"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the permissions of one or more members in a specific namespace.
//
// @param request - UpdateMemberRequest
//
// @return UpdateMemberResponse
func (client *Client) UpdateMember(namespace *string, request *UpdateMemberRequest) (_result *UpdateMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateMemberHeaders{}
	_result = &UpdateMemberResponse{}
	_body, _err := client.UpdateMemberWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the code of an SQL deployment.
//
// @param request - ValidateSqlStatementRequest
//
// @param headers - ValidateSqlStatementHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateSqlStatementResponse
func (client *Client) ValidateSqlStatementWithOptions(namespace *string, request *ValidateSqlStatementRequest, headers *ValidateSqlStatementHeaders, runtime *util.RuntimeOptions) (_result *ValidateSqlStatementResponse, _err error) {
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
		Action:      tea.String("ValidateSqlStatement"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/sql-statement/validate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateSqlStatementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the code of an SQL deployment.
//
// @param request - ValidateSqlStatementRequest
//
// @return ValidateSqlStatementResponse
func (client *Client) ValidateSqlStatement(namespace *string, request *ValidateSqlStatementRequest) (_result *ValidateSqlStatementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ValidateSqlStatementHeaders{}
	_result = &ValidateSqlStatementResponse{}
	_body, _err := client.ValidateSqlStatementWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
