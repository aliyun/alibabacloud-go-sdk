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
	Namespace *string        `json:"namespace,omitempty" xml:"namespace,omitempty"`
	Quota     *ResourceQuota `json:"quota,omitempty" xml:"quota,omitempty"`
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

func (s *DeploymentTarget) SetQuota(v *ResourceQuota) *DeploymentTarget {
	s.Quota = v
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

type Event struct {
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 00000000-0000-0000-0000-000000680003
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// example:
	//
	// 00000000-0000-0000-0000-000000000001
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// example:
	//
	// STATE_TRANSITION_IS_COMPLETED
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// example:
	//
	// 00000000-0000-0000-0000-000000005043
	JobId   *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// default-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// edcef******b4f
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s Event) String() string {
	return tea.Prettify(s)
}

func (s Event) GoString() string {
	return s.String()
}

func (s *Event) SetCreatedAt(v string) *Event {
	s.CreatedAt = &v
	return s
}

func (s *Event) SetDeploymentId(v string) *Event {
	s.DeploymentId = &v
	return s
}

func (s *Event) SetEventId(v string) *Event {
	s.EventId = &v
	return s
}

func (s *Event) SetEventName(v string) *Event {
	s.EventName = &v
	return s
}

func (s *Event) SetJobId(v string) *Event {
	s.JobId = &v
	return s
}

func (s *Event) SetMessage(v string) *Event {
	s.Message = &v
	return s
}

func (s *Event) SetNamespace(v string) *Event {
	s.Namespace = &v
	return s
}

func (s *Event) SetWorkspace(v string) *Event {
	s.Workspace = &v
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

type JobDiagnosis struct {
	DiagnoseId   *string               `json:"diagnoseId,omitempty" xml:"diagnoseId,omitempty"`
	DiagnoseTime *int64                `json:"diagnoseTime,omitempty" xml:"diagnoseTime,omitempty"`
	Namespace    *string               `json:"namespace,omitempty" xml:"namespace,omitempty"`
	RiskLevel    *string               `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	Symptoms     *JobDiagnosisSymptoms `json:"symptoms,omitempty" xml:"symptoms,omitempty"`
	Workspace    *string               `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s JobDiagnosis) String() string {
	return tea.Prettify(s)
}

func (s JobDiagnosis) GoString() string {
	return s.String()
}

func (s *JobDiagnosis) SetDiagnoseId(v string) *JobDiagnosis {
	s.DiagnoseId = &v
	return s
}

func (s *JobDiagnosis) SetDiagnoseTime(v int64) *JobDiagnosis {
	s.DiagnoseTime = &v
	return s
}

func (s *JobDiagnosis) SetNamespace(v string) *JobDiagnosis {
	s.Namespace = &v
	return s
}

func (s *JobDiagnosis) SetRiskLevel(v string) *JobDiagnosis {
	s.RiskLevel = &v
	return s
}

func (s *JobDiagnosis) SetSymptoms(v *JobDiagnosisSymptoms) *JobDiagnosis {
	s.Symptoms = v
	return s
}

func (s *JobDiagnosis) SetWorkspace(v string) *JobDiagnosis {
	s.Workspace = &v
	return s
}

type JobDiagnosisSymptom struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	Recommendation *string `json:"recommendation,omitempty" xml:"recommendation,omitempty"`
}

func (s JobDiagnosisSymptom) String() string {
	return tea.Prettify(s)
}

func (s JobDiagnosisSymptom) GoString() string {
	return s.String()
}

func (s *JobDiagnosisSymptom) SetDescription(v string) *JobDiagnosisSymptom {
	s.Description = &v
	return s
}

func (s *JobDiagnosisSymptom) SetName(v string) *JobDiagnosisSymptom {
	s.Name = &v
	return s
}

func (s *JobDiagnosisSymptom) SetRecommendation(v string) *JobDiagnosisSymptom {
	s.Recommendation = &v
	return s
}

type JobDiagnosisSymptoms struct {
	Autopilot       *JobDiagnosisSymptom   `json:"autopilot,omitempty" xml:"autopilot,omitempty"`
	Others          []*JobDiagnosisSymptom `json:"others,omitempty" xml:"others,omitempty" type:"Repeated"`
	Runtime         []*JobDiagnosisSymptom `json:"runtime,omitempty" xml:"runtime,omitempty" type:"Repeated"`
	Startup         []*JobDiagnosisSymptom `json:"startup,omitempty" xml:"startup,omitempty" type:"Repeated"`
	State           []*JobDiagnosisSymptom `json:"state,omitempty" xml:"state,omitempty" type:"Repeated"`
	Troubleshooting []*JobDiagnosisSymptom `json:"troubleshooting,omitempty" xml:"troubleshooting,omitempty" type:"Repeated"`
}

func (s JobDiagnosisSymptoms) String() string {
	return tea.Prettify(s)
}

func (s JobDiagnosisSymptoms) GoString() string {
	return s.String()
}

func (s *JobDiagnosisSymptoms) SetAutopilot(v *JobDiagnosisSymptom) *JobDiagnosisSymptoms {
	s.Autopilot = v
	return s
}

func (s *JobDiagnosisSymptoms) SetOthers(v []*JobDiagnosisSymptom) *JobDiagnosisSymptoms {
	s.Others = v
	return s
}

func (s *JobDiagnosisSymptoms) SetRuntime(v []*JobDiagnosisSymptom) *JobDiagnosisSymptoms {
	s.Runtime = v
	return s
}

func (s *JobDiagnosisSymptoms) SetStartup(v []*JobDiagnosisSymptom) *JobDiagnosisSymptoms {
	s.Startup = v
	return s
}

func (s *JobDiagnosisSymptoms) SetState(v []*JobDiagnosisSymptom) *JobDiagnosisSymptoms {
	s.State = v
	return s
}

func (s *JobDiagnosisSymptoms) SetTroubleshooting(v []*JobDiagnosisSymptom) *JobDiagnosisSymptoms {
	s.Troubleshooting = v
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
	JobId          *string          `json:"jobId,omitempty" xml:"jobId,omitempty"`
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

func (s *JobStartParameters) SetJobId(v string) *JobStartParameters {
	s.JobId = &v
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
	HealthScore      *int32            `json:"healthScore,omitempty" xml:"healthScore,omitempty"`
	RiskLevel        *string           `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
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

func (s *JobStatus) SetHealthScore(v int32) *JobStatus {
	s.HealthScore = &v
	return s
}

func (s *JobStatus) SetRiskLevel(v string) *JobStatus {
	s.RiskLevel = &v
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

type PeriodicSchedulingPolicy struct {
	IsFinished *bool `json:"isFinished,omitempty" xml:"isFinished,omitempty"`
	// example:
	//
	// 1723195800000
	OnlyOnceTriggerTime *int64 `json:"onlyOnceTriggerTime,omitempty" xml:"onlyOnceTriggerTime,omitempty"`
	// example:
	//
	// true
	OnlyOnceTriggerTimeIsExpired *bool `json:"onlyOnceTriggerTimeIsExpired,omitempty" xml:"onlyOnceTriggerTimeIsExpired,omitempty"`
	// example:
	//
	// DAY
	PeriodicSchedulingLevel  *string  `json:"periodicSchedulingLevel,omitempty" xml:"periodicSchedulingLevel,omitempty"`
	PeriodicSchedulingValues []*int32 `json:"periodicSchedulingValues,omitempty" xml:"periodicSchedulingValues,omitempty" type:"Repeated"`
	// example:
	//
	// 1723199340000
	PeriodicTriggerTime *int64                `json:"periodicTriggerTime,omitempty" xml:"periodicTriggerTime,omitempty"`
	ResourceSetting     *BriefResourceSetting `json:"resourceSetting,omitempty" xml:"resourceSetting,omitempty"`
}

func (s PeriodicSchedulingPolicy) String() string {
	return tea.Prettify(s)
}

func (s PeriodicSchedulingPolicy) GoString() string {
	return s.String()
}

func (s *PeriodicSchedulingPolicy) SetIsFinished(v bool) *PeriodicSchedulingPolicy {
	s.IsFinished = &v
	return s
}

func (s *PeriodicSchedulingPolicy) SetOnlyOnceTriggerTime(v int64) *PeriodicSchedulingPolicy {
	s.OnlyOnceTriggerTime = &v
	return s
}

func (s *PeriodicSchedulingPolicy) SetOnlyOnceTriggerTimeIsExpired(v bool) *PeriodicSchedulingPolicy {
	s.OnlyOnceTriggerTimeIsExpired = &v
	return s
}

func (s *PeriodicSchedulingPolicy) SetPeriodicSchedulingLevel(v string) *PeriodicSchedulingPolicy {
	s.PeriodicSchedulingLevel = &v
	return s
}

func (s *PeriodicSchedulingPolicy) SetPeriodicSchedulingValues(v []*int32) *PeriodicSchedulingPolicy {
	s.PeriodicSchedulingValues = v
	return s
}

func (s *PeriodicSchedulingPolicy) SetPeriodicTriggerTime(v int64) *PeriodicSchedulingPolicy {
	s.PeriodicTriggerTime = &v
	return s
}

func (s *PeriodicSchedulingPolicy) SetResourceSetting(v *BriefResourceSetting) *PeriodicSchedulingPolicy {
	s.ResourceSetting = v
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
	JobParallelism *int64 `json:"jobParallelism,omitempty" xml:"jobParallelism,omitempty"`
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

type ResourceQuota struct {
	Limit *ResourceSpec `json:"limit,omitempty" xml:"limit,omitempty"`
	Used  *ResourceSpec `json:"used,omitempty" xml:"used,omitempty"`
}

func (s ResourceQuota) String() string {
	return tea.Prettify(s)
}

func (s ResourceQuota) GoString() string {
	return s.String()
}

func (s *ResourceQuota) SetLimit(v *ResourceSpec) *ResourceQuota {
	s.Limit = v
	return s
}

func (s *ResourceQuota) SetUsed(v *ResourceSpec) *ResourceQuota {
	s.Used = v
	return s
}

type ResourceSpec struct {
	// example:
	//
	// 1.0
	Cpu *float64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// 4Gi
	Memory *string `json:"memory,omitempty" xml:"memory,omitempty"`
}

func (s ResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s ResourceSpec) GoString() string {
	return s.String()
}

func (s *ResourceSpec) SetCpu(v float64) *ResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ResourceSpec) SetMemory(v string) *ResourceSpec {
	s.Memory = &v
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

type ScheduledPlan struct {
	// example:
	//
	// 1723197248
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
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
	// 00000000-0000-0000-0000-000000000001
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// example:
	//
	// 1723197248
	ModifiedAt *string `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
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
	// test-scheduled-plan
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// default-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// USER_DEFINED
	Origin                     *string                     `json:"origin,omitempty" xml:"origin,omitempty"`
	PeriodicSchedulingPolicies []*PeriodicSchedulingPolicy `json:"periodicSchedulingPolicies,omitempty" xml:"periodicSchedulingPolicies,omitempty" type:"Repeated"`
	// example:
	//
	// f3b4ec1e-85dc-4b1d-9726-1d7f4c37****
	ScheduledPlanId *string `json:"scheduledPlanId,omitempty" xml:"scheduledPlanId,omitempty"`
	// example:
	//
	// FINISHED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	UpdatedByUser *bool `json:"updatedByUser,omitempty" xml:"updatedByUser,omitempty"`
	// example:
	//
	// edcef******b4f
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ScheduledPlan) String() string {
	return tea.Prettify(s)
}

func (s ScheduledPlan) GoString() string {
	return s.String()
}

func (s *ScheduledPlan) SetCreatedAt(v string) *ScheduledPlan {
	s.CreatedAt = &v
	return s
}

func (s *ScheduledPlan) SetCreator(v string) *ScheduledPlan {
	s.Creator = &v
	return s
}

func (s *ScheduledPlan) SetCreatorName(v string) *ScheduledPlan {
	s.CreatorName = &v
	return s
}

func (s *ScheduledPlan) SetDeploymentId(v string) *ScheduledPlan {
	s.DeploymentId = &v
	return s
}

func (s *ScheduledPlan) SetModifiedAt(v string) *ScheduledPlan {
	s.ModifiedAt = &v
	return s
}

func (s *ScheduledPlan) SetModifier(v string) *ScheduledPlan {
	s.Modifier = &v
	return s
}

func (s *ScheduledPlan) SetModifierName(v string) *ScheduledPlan {
	s.ModifierName = &v
	return s
}

func (s *ScheduledPlan) SetName(v string) *ScheduledPlan {
	s.Name = &v
	return s
}

func (s *ScheduledPlan) SetNamespace(v string) *ScheduledPlan {
	s.Namespace = &v
	return s
}

func (s *ScheduledPlan) SetOrigin(v string) *ScheduledPlan {
	s.Origin = &v
	return s
}

func (s *ScheduledPlan) SetPeriodicSchedulingPolicies(v []*PeriodicSchedulingPolicy) *ScheduledPlan {
	s.PeriodicSchedulingPolicies = v
	return s
}

func (s *ScheduledPlan) SetScheduledPlanId(v string) *ScheduledPlan {
	s.ScheduledPlanId = &v
	return s
}

func (s *ScheduledPlan) SetStatus(v string) *ScheduledPlan {
	s.Status = &v
	return s
}

func (s *ScheduledPlan) SetUpdatedByUser(v bool) *ScheduledPlan {
	s.UpdatedByUser = &v
	return s
}

func (s *ScheduledPlan) SetWorkspace(v string) *ScheduledPlan {
	s.Workspace = &v
	return s
}

type ScheduledPlanAppliedInfo struct {
	// example:
	//
	// 00000000-0000-0000-0000-000000000001
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// example:
	//
	// RUNNING
	ExpectedState *string `json:"expectedState,omitempty" xml:"expectedState,omitempty"`
	// example:
	//
	// 1723197248
	ModifiedAt *string `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
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
	// default-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// f3b4ec1e-85dc-4b1d-9726-1d7f4c37****
	ScheduledPlanId *string `json:"scheduledPlanId,omitempty" xml:"scheduledPlanId,omitempty"`
	// example:
	//
	// test-scheduled-plan
	ScheduledPlanName *string `json:"scheduledPlanName,omitempty" xml:"scheduledPlanName,omitempty"`
	// example:
	//
	// WAITING
	StatusState *string `json:"statusState,omitempty" xml:"statusState,omitempty"`
	// example:
	//
	// edcef******b4f
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ScheduledPlanAppliedInfo) String() string {
	return tea.Prettify(s)
}

func (s ScheduledPlanAppliedInfo) GoString() string {
	return s.String()
}

func (s *ScheduledPlanAppliedInfo) SetDeploymentId(v string) *ScheduledPlanAppliedInfo {
	s.DeploymentId = &v
	return s
}

func (s *ScheduledPlanAppliedInfo) SetExpectedState(v string) *ScheduledPlanAppliedInfo {
	s.ExpectedState = &v
	return s
}

func (s *ScheduledPlanAppliedInfo) SetModifiedAt(v string) *ScheduledPlanAppliedInfo {
	s.ModifiedAt = &v
	return s
}

func (s *ScheduledPlanAppliedInfo) SetModifier(v string) *ScheduledPlanAppliedInfo {
	s.Modifier = &v
	return s
}

func (s *ScheduledPlanAppliedInfo) SetModifierName(v string) *ScheduledPlanAppliedInfo {
	s.ModifierName = &v
	return s
}

func (s *ScheduledPlanAppliedInfo) SetNamespace(v string) *ScheduledPlanAppliedInfo {
	s.Namespace = &v
	return s
}

func (s *ScheduledPlanAppliedInfo) SetScheduledPlanId(v string) *ScheduledPlanAppliedInfo {
	s.ScheduledPlanId = &v
	return s
}

func (s *ScheduledPlanAppliedInfo) SetScheduledPlanName(v string) *ScheduledPlanAppliedInfo {
	s.ScheduledPlanName = &v
	return s
}

func (s *ScheduledPlanAppliedInfo) SetStatusState(v string) *ScheduledPlanAppliedInfo {
	s.StatusState = &v
	return s
}

func (s *ScheduledPlanAppliedInfo) SetWorkspace(v string) *ScheduledPlanAppliedInfo {
	s.Workspace = &v
	return s
}

type ScheduledPlanExecutedInfo struct {
	// example:
	//
	// 1723197248
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
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
	// 00000000-0000-0000-0000-000000000001
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// example:
	//
	// 0e6d3bab-2277-4ed1-b573-9de6413d****
	JobResourceUpgradingId *string `json:"jobResourceUpgradingId,omitempty" xml:"jobResourceUpgradingId,omitempty"`
	// example:
	//
	// 1723197248
	ModifiedAt *string `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
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
	// test-scheduled-plan
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// default-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// SCHEDULED_PLAN
	Origin *string `json:"origin,omitempty" xml:"origin,omitempty"`
	// example:
	//
	// f8a2d5d9-9fc5-4273-bfcc-2a3cd354****
	OriginJobId *string                      `json:"originJobId,omitempty" xml:"originJobId,omitempty"`
	Status      *ScheduledPlanExecutedStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// edcef******b4f
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ScheduledPlanExecutedInfo) String() string {
	return tea.Prettify(s)
}

func (s ScheduledPlanExecutedInfo) GoString() string {
	return s.String()
}

func (s *ScheduledPlanExecutedInfo) SetCreatedAt(v string) *ScheduledPlanExecutedInfo {
	s.CreatedAt = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetCreator(v string) *ScheduledPlanExecutedInfo {
	s.Creator = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetCreatorName(v string) *ScheduledPlanExecutedInfo {
	s.CreatorName = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetDeploymentId(v string) *ScheduledPlanExecutedInfo {
	s.DeploymentId = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetJobResourceUpgradingId(v string) *ScheduledPlanExecutedInfo {
	s.JobResourceUpgradingId = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetModifiedAt(v string) *ScheduledPlanExecutedInfo {
	s.ModifiedAt = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetModifier(v string) *ScheduledPlanExecutedInfo {
	s.Modifier = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetModifierName(v string) *ScheduledPlanExecutedInfo {
	s.ModifierName = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetName(v string) *ScheduledPlanExecutedInfo {
	s.Name = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetNamespace(v string) *ScheduledPlanExecutedInfo {
	s.Namespace = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetOrigin(v string) *ScheduledPlanExecutedInfo {
	s.Origin = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetOriginJobId(v string) *ScheduledPlanExecutedInfo {
	s.OriginJobId = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetStatus(v *ScheduledPlanExecutedStatus) *ScheduledPlanExecutedInfo {
	s.Status = v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetWorkspace(v string) *ScheduledPlanExecutedInfo {
	s.Workspace = &v
	return s
}

type ScheduledPlanExecutedStatus struct {
	// example:
	//
	// HOT_UPDATE
	RestartType *string `json:"restartType,omitempty" xml:"restartType,omitempty"`
	// example:
	//
	// UPGRADED
	StatusState *string `json:"statusState,omitempty" xml:"statusState,omitempty"`
}

func (s ScheduledPlanExecutedStatus) String() string {
	return tea.Prettify(s)
}

func (s ScheduledPlanExecutedStatus) GoString() string {
	return s.String()
}

func (s *ScheduledPlanExecutedStatus) SetRestartType(v string) *ScheduledPlanExecutedStatus {
	s.RestartType = &v
	return s
}

func (s *ScheduledPlanExecutedStatus) SetStatusState(v string) *ScheduledPlanExecutedStatus {
	s.StatusState = &v
	return s
}

type Schema struct {
	Columns        []*TableColumn   `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	PrimaryKey     *PrimaryKey      `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
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

func (s *Schema) SetPrimaryKey(v *PrimaryKey) *Schema {
	s.PrimaryKey = v
	return s
}

func (s *Schema) SetWatermarkSpecs(v []*WatermarkSpec) *Schema {
	s.WatermarkSpecs = v
	return s
}

type SessionCluster struct {
	BasicResourceSetting *BasicResourceSetting `json:"basicResourceSetting,omitempty" xml:"basicResourceSetting,omitempty"`
	CreatedAt            *int64                `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
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
	// default-queue
	DeploymentTargetName *string `json:"deploymentTargetName,omitempty" xml:"deploymentTargetName,omitempty"`
	// example:
	//
	// vvr-6.0.7-flink-1.15
	EngineVersion *string `json:"engineVersion,omitempty" xml:"engineVersion,omitempty"`
	// example:
	//
	// {"taskmanager.numberOfTaskSlots":"1"}
	FlinkConf  map[string]interface{} `json:"flinkConf,omitempty" xml:"flinkConf,omitempty"`
	Labels     map[string]interface{} `json:"labels,omitempty" xml:"labels,omitempty"`
	Logging    *Logging               `json:"logging,omitempty" xml:"logging,omitempty"`
	ModifiedAt *int64                 `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
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
	// test-sessionCluster
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// default-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// 1f68a52c-1760-43c6-97fb-afe0674b****
	SessionClusterId *string               `json:"sessionClusterId,omitempty" xml:"sessionClusterId,omitempty"`
	Status           *SessionClusterStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// edcef******b4f
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s SessionCluster) String() string {
	return tea.Prettify(s)
}

func (s SessionCluster) GoString() string {
	return s.String()
}

func (s *SessionCluster) SetBasicResourceSetting(v *BasicResourceSetting) *SessionCluster {
	s.BasicResourceSetting = v
	return s
}

func (s *SessionCluster) SetCreatedAt(v int64) *SessionCluster {
	s.CreatedAt = &v
	return s
}

func (s *SessionCluster) SetCreator(v string) *SessionCluster {
	s.Creator = &v
	return s
}

func (s *SessionCluster) SetCreatorName(v string) *SessionCluster {
	s.CreatorName = &v
	return s
}

func (s *SessionCluster) SetDeploymentTargetName(v string) *SessionCluster {
	s.DeploymentTargetName = &v
	return s
}

func (s *SessionCluster) SetEngineVersion(v string) *SessionCluster {
	s.EngineVersion = &v
	return s
}

func (s *SessionCluster) SetFlinkConf(v map[string]interface{}) *SessionCluster {
	s.FlinkConf = v
	return s
}

func (s *SessionCluster) SetLabels(v map[string]interface{}) *SessionCluster {
	s.Labels = v
	return s
}

func (s *SessionCluster) SetLogging(v *Logging) *SessionCluster {
	s.Logging = v
	return s
}

func (s *SessionCluster) SetModifiedAt(v int64) *SessionCluster {
	s.ModifiedAt = &v
	return s
}

func (s *SessionCluster) SetModifier(v string) *SessionCluster {
	s.Modifier = &v
	return s
}

func (s *SessionCluster) SetModifierName(v string) *SessionCluster {
	s.ModifierName = &v
	return s
}

func (s *SessionCluster) SetName(v string) *SessionCluster {
	s.Name = &v
	return s
}

func (s *SessionCluster) SetNamespace(v string) *SessionCluster {
	s.Namespace = &v
	return s
}

func (s *SessionCluster) SetSessionClusterId(v string) *SessionCluster {
	s.SessionClusterId = &v
	return s
}

func (s *SessionCluster) SetStatus(v *SessionClusterStatus) *SessionCluster {
	s.Status = v
	return s
}

func (s *SessionCluster) SetWorkspace(v string) *SessionCluster {
	s.Workspace = &v
	return s
}

type SessionClusterFailureInfo struct {
	FailedAt *int64  `json:"failedAt,omitempty" xml:"failedAt,omitempty"`
	Message  *string `json:"message,omitempty" xml:"message,omitempty"`
	Reason   *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s SessionClusterFailureInfo) String() string {
	return tea.Prettify(s)
}

func (s SessionClusterFailureInfo) GoString() string {
	return s.String()
}

func (s *SessionClusterFailureInfo) SetFailedAt(v int64) *SessionClusterFailureInfo {
	s.FailedAt = &v
	return s
}

func (s *SessionClusterFailureInfo) SetMessage(v string) *SessionClusterFailureInfo {
	s.Message = &v
	return s
}

func (s *SessionClusterFailureInfo) SetReason(v string) *SessionClusterFailureInfo {
	s.Reason = &v
	return s
}

type SessionClusterRunningInfo struct {
	LastUpdateTime         *int64    `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty"`
	ReferenceDeploymentIds []*string `json:"referenceDeploymentIds,omitempty" xml:"referenceDeploymentIds,omitempty" type:"Repeated"`
	StartedAt              *int64    `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
}

func (s SessionClusterRunningInfo) String() string {
	return tea.Prettify(s)
}

func (s SessionClusterRunningInfo) GoString() string {
	return s.String()
}

func (s *SessionClusterRunningInfo) SetLastUpdateTime(v int64) *SessionClusterRunningInfo {
	s.LastUpdateTime = &v
	return s
}

func (s *SessionClusterRunningInfo) SetReferenceDeploymentIds(v []*string) *SessionClusterRunningInfo {
	s.ReferenceDeploymentIds = v
	return s
}

func (s *SessionClusterRunningInfo) SetStartedAt(v int64) *SessionClusterRunningInfo {
	s.StartedAt = &v
	return s
}

type SessionClusterStatus struct {
	CurrentSessionClusterStatus *string                    `json:"currentSessionClusterStatus,omitempty" xml:"currentSessionClusterStatus,omitempty"`
	Failure                     *SessionClusterFailureInfo `json:"failure,omitempty" xml:"failure,omitempty"`
	Running                     *SessionClusterRunningInfo `json:"running,omitempty" xml:"running,omitempty"`
}

func (s SessionClusterStatus) String() string {
	return tea.Prettify(s)
}

func (s SessionClusterStatus) GoString() string {
	return s.String()
}

func (s *SessionClusterStatus) SetCurrentSessionClusterStatus(v string) *SessionClusterStatus {
	s.CurrentSessionClusterStatus = &v
	return s
}

func (s *SessionClusterStatus) SetFailure(v *SessionClusterFailureInfo) *SessionClusterStatus {
	s.Failure = v
	return s
}

func (s *SessionClusterStatus) SetRunning(v *SessionClusterRunningInfo) *SessionClusterStatus {
	s.Running = v
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

type ApplyScheduledPlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ApplyScheduledPlanHeaders) String() string {
	return tea.Prettify(s)
}

func (s ApplyScheduledPlanHeaders) GoString() string {
	return s.String()
}

func (s *ApplyScheduledPlanHeaders) SetCommonHeaders(v map[string]*string) *ApplyScheduledPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApplyScheduledPlanHeaders) SetWorkspace(v string) *ApplyScheduledPlanHeaders {
	s.Workspace = &v
	return s
}

type ApplyScheduledPlanResponseBody struct {
	Data *ScheduledPlanAppliedInfo `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ApplyScheduledPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyScheduledPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyScheduledPlanResponseBody) SetData(v *ScheduledPlanAppliedInfo) *ApplyScheduledPlanResponseBody {
	s.Data = v
	return s
}

func (s *ApplyScheduledPlanResponseBody) SetErrorCode(v string) *ApplyScheduledPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ApplyScheduledPlanResponseBody) SetErrorMessage(v string) *ApplyScheduledPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ApplyScheduledPlanResponseBody) SetHttpCode(v int32) *ApplyScheduledPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ApplyScheduledPlanResponseBody) SetRequestId(v string) *ApplyScheduledPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyScheduledPlanResponseBody) SetSuccess(v bool) *ApplyScheduledPlanResponseBody {
	s.Success = &v
	return s
}

type ApplyScheduledPlanResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyScheduledPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyScheduledPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyScheduledPlanResponse) GoString() string {
	return s.String()
}

func (s *ApplyScheduledPlanResponse) SetHeaders(v map[string]*string) *ApplyScheduledPlanResponse {
	s.Headers = v
	return s
}

func (s *ApplyScheduledPlanResponse) SetStatusCode(v int32) *ApplyScheduledPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyScheduledPlanResponse) SetBody(v *ApplyScheduledPlanResponseBody) *ApplyScheduledPlanResponse {
	s.Body = v
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

type CreateDeploymentDraftHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateDeploymentDraftHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentDraftHeaders) GoString() string {
	return s.String()
}

func (s *CreateDeploymentDraftHeaders) SetCommonHeaders(v map[string]*string) *CreateDeploymentDraftHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDeploymentDraftHeaders) SetWorkspace(v string) *CreateDeploymentDraftHeaders {
	s.Workspace = &v
	return s
}

type CreateDeploymentDraftRequest struct {
	// This parameter is required.
	Body *DeploymentDraft `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeploymentDraftRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentDraftRequest) GoString() string {
	return s.String()
}

func (s *CreateDeploymentDraftRequest) SetBody(v *DeploymentDraft) *CreateDeploymentDraftRequest {
	s.Body = v
	return s
}

type CreateDeploymentDraftResponseBody struct {
	Data *DeploymentDraft `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateDeploymentDraftResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentDraftResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeploymentDraftResponseBody) SetData(v *DeploymentDraft) *CreateDeploymentDraftResponseBody {
	s.Data = v
	return s
}

func (s *CreateDeploymentDraftResponseBody) SetErrorCode(v string) *CreateDeploymentDraftResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDeploymentDraftResponseBody) SetErrorMessage(v string) *CreateDeploymentDraftResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDeploymentDraftResponseBody) SetHttpCode(v int32) *CreateDeploymentDraftResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateDeploymentDraftResponseBody) SetRequestId(v string) *CreateDeploymentDraftResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeploymentDraftResponseBody) SetSuccess(v bool) *CreateDeploymentDraftResponseBody {
	s.Success = &v
	return s
}

type CreateDeploymentDraftResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeploymentDraftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeploymentDraftResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentDraftResponse) GoString() string {
	return s.String()
}

func (s *CreateDeploymentDraftResponse) SetHeaders(v map[string]*string) *CreateDeploymentDraftResponse {
	s.Headers = v
	return s
}

func (s *CreateDeploymentDraftResponse) SetStatusCode(v int32) *CreateDeploymentDraftResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeploymentDraftResponse) SetBody(v *CreateDeploymentDraftResponseBody) *CreateDeploymentDraftResponse {
	s.Body = v
	return s
}

type CreateDeploymentTargetHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateDeploymentTargetHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentTargetHeaders) GoString() string {
	return s.String()
}

func (s *CreateDeploymentTargetHeaders) SetCommonHeaders(v map[string]*string) *CreateDeploymentTargetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDeploymentTargetHeaders) SetWorkspace(v string) *CreateDeploymentTargetHeaders {
	s.Workspace = &v
	return s
}

type CreateDeploymentTargetRequest struct {
	Body *ResourceSpec `json:"body,omitempty" xml:"body,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-dt
	DeploymentTargetName *string `json:"deploymentTargetName,omitempty" xml:"deploymentTargetName,omitempty"`
}

func (s CreateDeploymentTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentTargetRequest) GoString() string {
	return s.String()
}

func (s *CreateDeploymentTargetRequest) SetBody(v *ResourceSpec) *CreateDeploymentTargetRequest {
	s.Body = v
	return s
}

func (s *CreateDeploymentTargetRequest) SetDeploymentTargetName(v string) *CreateDeploymentTargetRequest {
	s.DeploymentTargetName = &v
	return s
}

type CreateDeploymentTargetResponseBody struct {
	Data *DeploymentTarget `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateDeploymentTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentTargetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeploymentTargetResponseBody) SetData(v *DeploymentTarget) *CreateDeploymentTargetResponseBody {
	s.Data = v
	return s
}

func (s *CreateDeploymentTargetResponseBody) SetErrorCode(v string) *CreateDeploymentTargetResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDeploymentTargetResponseBody) SetErrorMessage(v string) *CreateDeploymentTargetResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDeploymentTargetResponseBody) SetHttpCode(v int32) *CreateDeploymentTargetResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateDeploymentTargetResponseBody) SetRequestId(v string) *CreateDeploymentTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeploymentTargetResponseBody) SetSuccess(v bool) *CreateDeploymentTargetResponseBody {
	s.Success = &v
	return s
}

type CreateDeploymentTargetResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeploymentTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeploymentTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentTargetResponse) GoString() string {
	return s.String()
}

func (s *CreateDeploymentTargetResponse) SetHeaders(v map[string]*string) *CreateDeploymentTargetResponse {
	s.Headers = v
	return s
}

func (s *CreateDeploymentTargetResponse) SetStatusCode(v int32) *CreateDeploymentTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeploymentTargetResponse) SetBody(v *CreateDeploymentTargetResponseBody) *CreateDeploymentTargetResponse {
	s.Body = v
	return s
}

type CreateFolderHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateFolderHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateFolderHeaders) GoString() string {
	return s.String()
}

func (s *CreateFolderHeaders) SetCommonHeaders(v map[string]*string) *CreateFolderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateFolderHeaders) SetWorkspace(v string) *CreateFolderHeaders {
	s.Workspace = &v
	return s
}

type CreateFolderRequest struct {
	Body *Folder `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFolderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFolderRequest) GoString() string {
	return s.String()
}

func (s *CreateFolderRequest) SetBody(v *Folder) *CreateFolderRequest {
	s.Body = v
	return s
}

type CreateFolderResponseBody struct {
	Data *Folder `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFolderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFolderResponseBody) SetData(v *Folder) *CreateFolderResponseBody {
	s.Data = v
	return s
}

func (s *CreateFolderResponseBody) SetErrorCode(v string) *CreateFolderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFolderResponseBody) SetErrorMessage(v string) *CreateFolderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateFolderResponseBody) SetHttpCode(v int32) *CreateFolderResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateFolderResponseBody) SetRequestId(v string) *CreateFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFolderResponseBody) SetSuccess(v bool) *CreateFolderResponseBody {
	s.Success = &v
	return s
}

type CreateFolderResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFolderResponse) GoString() string {
	return s.String()
}

func (s *CreateFolderResponse) SetHeaders(v map[string]*string) *CreateFolderResponse {
	s.Headers = v
	return s
}

func (s *CreateFolderResponse) SetStatusCode(v int32) *CreateFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFolderResponse) SetBody(v *CreateFolderResponseBody) *CreateFolderResponse {
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

type CreateScheduledPlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateScheduledPlanHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledPlanHeaders) GoString() string {
	return s.String()
}

func (s *CreateScheduledPlanHeaders) SetCommonHeaders(v map[string]*string) *CreateScheduledPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateScheduledPlanHeaders) SetWorkspace(v string) *CreateScheduledPlanHeaders {
	s.Workspace = &v
	return s
}

type CreateScheduledPlanRequest struct {
	Body *ScheduledPlan `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScheduledPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledPlanRequest) SetBody(v *ScheduledPlan) *CreateScheduledPlanRequest {
	s.Body = v
	return s
}

type CreateScheduledPlanResponseBody struct {
	Data *ScheduledPlan `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-ABCD-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateScheduledPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduledPlanResponseBody) SetData(v *ScheduledPlan) *CreateScheduledPlanResponseBody {
	s.Data = v
	return s
}

func (s *CreateScheduledPlanResponseBody) SetErrorCode(v string) *CreateScheduledPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateScheduledPlanResponseBody) SetErrorMessage(v string) *CreateScheduledPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateScheduledPlanResponseBody) SetHttpCode(v int32) *CreateScheduledPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateScheduledPlanResponseBody) SetRequestId(v string) *CreateScheduledPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduledPlanResponseBody) SetSuccess(v bool) *CreateScheduledPlanResponseBody {
	s.Success = &v
	return s
}

type CreateScheduledPlanResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScheduledPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScheduledPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduledPlanResponse) SetHeaders(v map[string]*string) *CreateScheduledPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduledPlanResponse) SetStatusCode(v int32) *CreateScheduledPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScheduledPlanResponse) SetBody(v *CreateScheduledPlanResponseBody) *CreateScheduledPlanResponse {
	s.Body = v
	return s
}

type CreateSessionClusterHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateSessionClusterHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateSessionClusterHeaders) GoString() string {
	return s.String()
}

func (s *CreateSessionClusterHeaders) SetCommonHeaders(v map[string]*string) *CreateSessionClusterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSessionClusterHeaders) SetWorkspace(v string) *CreateSessionClusterHeaders {
	s.Workspace = &v
	return s
}

type CreateSessionClusterRequest struct {
	Body *SessionCluster `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSessionClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSessionClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateSessionClusterRequest) SetBody(v *SessionCluster) *CreateSessionClusterRequest {
	s.Body = v
	return s
}

type CreateSessionClusterResponseBody struct {
	Data *SessionCluster `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateSessionClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSessionClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSessionClusterResponseBody) SetData(v *SessionCluster) *CreateSessionClusterResponseBody {
	s.Data = v
	return s
}

func (s *CreateSessionClusterResponseBody) SetErrorCode(v string) *CreateSessionClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateSessionClusterResponseBody) SetErrorMessage(v string) *CreateSessionClusterResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateSessionClusterResponseBody) SetHttpCode(v int32) *CreateSessionClusterResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateSessionClusterResponseBody) SetRequestId(v string) *CreateSessionClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSessionClusterResponseBody) SetSuccess(v bool) *CreateSessionClusterResponseBody {
	s.Success = &v
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

type CreateUdfArtifactHeaders struct {
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

func (s CreateUdfArtifactHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateUdfArtifactHeaders) GoString() string {
	return s.String()
}

func (s *CreateUdfArtifactHeaders) SetCommonHeaders(v map[string]*string) *CreateUdfArtifactHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateUdfArtifactHeaders) SetWorkspace(v string) *CreateUdfArtifactHeaders {
	s.Workspace = &v
	return s
}

type CreateUdfArtifactRequest struct {
	// The resource file of the UDF.
	//
	// This parameter is required.
	Body *UdfArtifact `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUdfArtifactRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUdfArtifactRequest) GoString() string {
	return s.String()
}

func (s *CreateUdfArtifactRequest) SetBody(v *UdfArtifact) *CreateUdfArtifactRequest {
	s.Body = v
	return s
}

type CreateUdfArtifactResponseBody struct {
	// The result of creating an artifact configuration for the UDF.
	Data *CreateUdfArtifactResult `json:"data,omitempty" xml:"data,omitempty"`
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
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateUdfArtifactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUdfArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUdfArtifactResponseBody) SetData(v *CreateUdfArtifactResult) *CreateUdfArtifactResponseBody {
	s.Data = v
	return s
}

func (s *CreateUdfArtifactResponseBody) SetErrorCode(v string) *CreateUdfArtifactResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateUdfArtifactResponseBody) SetErrorMessage(v string) *CreateUdfArtifactResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateUdfArtifactResponseBody) SetHttpCode(v int32) *CreateUdfArtifactResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateUdfArtifactResponseBody) SetRequestId(v string) *CreateUdfArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUdfArtifactResponseBody) SetSuccess(v bool) *CreateUdfArtifactResponseBody {
	s.Success = &v
	return s
}

type CreateUdfArtifactResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUdfArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUdfArtifactResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUdfArtifactResponse) GoString() string {
	return s.String()
}

func (s *CreateUdfArtifactResponse) SetHeaders(v map[string]*string) *CreateUdfArtifactResponse {
	s.Headers = v
	return s
}

func (s *CreateUdfArtifactResponse) SetStatusCode(v int32) *CreateUdfArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUdfArtifactResponse) SetBody(v *CreateUdfArtifactResponseBody) *CreateUdfArtifactResponse {
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

type DeleteCustomConnectorHeaders struct {
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

func (s DeleteCustomConnectorHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomConnectorHeaders) GoString() string {
	return s.String()
}

func (s *DeleteCustomConnectorHeaders) SetCommonHeaders(v map[string]*string) *DeleteCustomConnectorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteCustomConnectorHeaders) SetWorkspace(v string) *DeleteCustomConnectorHeaders {
	s.Workspace = &v
	return s
}

type DeleteCustomConnectorResponseBody struct {
	// If the value of success was true, a list of deployments in which custom connectors were deleted was returned. If the value of success was false, a null value was returned.
	Data []*TableMeta `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s DeleteCustomConnectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomConnectorResponseBody) SetData(v []*TableMeta) *DeleteCustomConnectorResponseBody {
	s.Data = v
	return s
}

func (s *DeleteCustomConnectorResponseBody) SetErrorCode(v string) *DeleteCustomConnectorResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteCustomConnectorResponseBody) SetErrorMessage(v string) *DeleteCustomConnectorResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteCustomConnectorResponseBody) SetHttpCode(v int32) *DeleteCustomConnectorResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteCustomConnectorResponseBody) SetRequestId(v string) *DeleteCustomConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomConnectorResponseBody) SetSuccess(v bool) *DeleteCustomConnectorResponseBody {
	s.Success = &v
	return s
}

type DeleteCustomConnectorResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomConnectorResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomConnectorResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomConnectorResponse) SetHeaders(v map[string]*string) *DeleteCustomConnectorResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomConnectorResponse) SetStatusCode(v int32) *DeleteCustomConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomConnectorResponse) SetBody(v *DeleteCustomConnectorResponseBody) *DeleteCustomConnectorResponse {
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

type DeleteDeploymentDraftHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteDeploymentDraftHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeploymentDraftHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentDraftHeaders) SetCommonHeaders(v map[string]*string) *DeleteDeploymentDraftHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDeploymentDraftHeaders) SetWorkspace(v string) *DeleteDeploymentDraftHeaders {
	s.Workspace = &v
	return s
}

type DeleteDeploymentDraftResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteDeploymentDraftResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeploymentDraftResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentDraftResponseBody) SetErrorCode(v string) *DeleteDeploymentDraftResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDeploymentDraftResponseBody) SetErrorMessage(v string) *DeleteDeploymentDraftResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDeploymentDraftResponseBody) SetHttpCode(v int32) *DeleteDeploymentDraftResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteDeploymentDraftResponseBody) SetRequestId(v string) *DeleteDeploymentDraftResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeploymentDraftResponseBody) SetSuccess(v bool) *DeleteDeploymentDraftResponseBody {
	s.Success = &v
	return s
}

type DeleteDeploymentDraftResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeploymentDraftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDeploymentDraftResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeploymentDraftResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentDraftResponse) SetHeaders(v map[string]*string) *DeleteDeploymentDraftResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeploymentDraftResponse) SetStatusCode(v int32) *DeleteDeploymentDraftResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeploymentDraftResponse) SetBody(v *DeleteDeploymentDraftResponseBody) *DeleteDeploymentDraftResponse {
	s.Body = v
	return s
}

type DeleteDeploymentTargetHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteDeploymentTargetHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeploymentTargetHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentTargetHeaders) SetCommonHeaders(v map[string]*string) *DeleteDeploymentTargetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDeploymentTargetHeaders) SetWorkspace(v string) *DeleteDeploymentTargetHeaders {
	s.Workspace = &v
	return s
}

type DeleteDeploymentTargetResponseBody struct {
	Data *DeploymentTarget `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteDeploymentTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeploymentTargetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentTargetResponseBody) SetData(v *DeploymentTarget) *DeleteDeploymentTargetResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDeploymentTargetResponseBody) SetErrorCode(v string) *DeleteDeploymentTargetResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDeploymentTargetResponseBody) SetErrorMessage(v string) *DeleteDeploymentTargetResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDeploymentTargetResponseBody) SetHttpCode(v int32) *DeleteDeploymentTargetResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteDeploymentTargetResponseBody) SetRequestId(v string) *DeleteDeploymentTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeploymentTargetResponseBody) SetSuccess(v bool) *DeleteDeploymentTargetResponseBody {
	s.Success = &v
	return s
}

type DeleteDeploymentTargetResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeploymentTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDeploymentTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeploymentTargetResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentTargetResponse) SetHeaders(v map[string]*string) *DeleteDeploymentTargetResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeploymentTargetResponse) SetStatusCode(v int32) *DeleteDeploymentTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeploymentTargetResponse) SetBody(v *DeleteDeploymentTargetResponseBody) *DeleteDeploymentTargetResponse {
	s.Body = v
	return s
}

type DeleteFolderHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c96306e2b****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteFolderHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteFolderHeaders) GoString() string {
	return s.String()
}

func (s *DeleteFolderHeaders) SetCommonHeaders(v map[string]*string) *DeleteFolderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteFolderHeaders) SetWorkspace(v string) *DeleteFolderHeaders {
	s.Workspace = &v
	return s
}

type DeleteFolderResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFolderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFolderResponseBody) SetErrorCode(v string) *DeleteFolderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteFolderResponseBody) SetErrorMessage(v string) *DeleteFolderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFolderResponseBody) SetHttpCode(v int32) *DeleteFolderResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteFolderResponseBody) SetRequestId(v string) *DeleteFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFolderResponseBody) SetSuccess(v bool) *DeleteFolderResponseBody {
	s.Success = &v
	return s
}

type DeleteFolderResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFolderResponse) GoString() string {
	return s.String()
}

func (s *DeleteFolderResponse) SetHeaders(v map[string]*string) *DeleteFolderResponse {
	s.Headers = v
	return s
}

func (s *DeleteFolderResponse) SetStatusCode(v int32) *DeleteFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFolderResponse) SetBody(v *DeleteFolderResponseBody) *DeleteFolderResponse {
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

type DeleteScheduledPlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteScheduledPlanHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduledPlanHeaders) GoString() string {
	return s.String()
}

func (s *DeleteScheduledPlanHeaders) SetCommonHeaders(v map[string]*string) *DeleteScheduledPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteScheduledPlanHeaders) SetWorkspace(v string) *DeleteScheduledPlanHeaders {
	s.Workspace = &v
	return s
}

type DeleteScheduledPlanResponseBody struct {
	Data *ScheduledPlan `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteScheduledPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduledPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduledPlanResponseBody) SetData(v *ScheduledPlan) *DeleteScheduledPlanResponseBody {
	s.Data = v
	return s
}

func (s *DeleteScheduledPlanResponseBody) SetErrorCode(v string) *DeleteScheduledPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteScheduledPlanResponseBody) SetErrorMessage(v string) *DeleteScheduledPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteScheduledPlanResponseBody) SetHttpCode(v int32) *DeleteScheduledPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteScheduledPlanResponseBody) SetRequestId(v string) *DeleteScheduledPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScheduledPlanResponseBody) SetSuccess(v bool) *DeleteScheduledPlanResponseBody {
	s.Success = &v
	return s
}

type DeleteScheduledPlanResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScheduledPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScheduledPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduledPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduledPlanResponse) SetHeaders(v map[string]*string) *DeleteScheduledPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduledPlanResponse) SetStatusCode(v int32) *DeleteScheduledPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScheduledPlanResponse) SetBody(v *DeleteScheduledPlanResponseBody) *DeleteScheduledPlanResponse {
	s.Body = v
	return s
}

type DeleteSessionClusterHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteSessionClusterHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteSessionClusterHeaders) GoString() string {
	return s.String()
}

func (s *DeleteSessionClusterHeaders) SetCommonHeaders(v map[string]*string) *DeleteSessionClusterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteSessionClusterHeaders) SetWorkspace(v string) *DeleteSessionClusterHeaders {
	s.Workspace = &v
	return s
}

type DeleteSessionClusterResponseBody struct {
	Data *SessionCluster `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-ABCD-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteSessionClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSessionClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSessionClusterResponseBody) SetData(v *SessionCluster) *DeleteSessionClusterResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSessionClusterResponseBody) SetErrorCode(v string) *DeleteSessionClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteSessionClusterResponseBody) SetErrorMessage(v string) *DeleteSessionClusterResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteSessionClusterResponseBody) SetHttpCode(v int32) *DeleteSessionClusterResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteSessionClusterResponseBody) SetRequestId(v string) *DeleteSessionClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSessionClusterResponseBody) SetSuccess(v bool) *DeleteSessionClusterResponseBody {
	s.Success = &v
	return s
}

type DeleteSessionClusterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSessionClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSessionClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSessionClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteSessionClusterResponse) SetHeaders(v map[string]*string) *DeleteSessionClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteSessionClusterResponse) SetStatusCode(v int32) *DeleteSessionClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSessionClusterResponse) SetBody(v *DeleteSessionClusterResponseBody) *DeleteSessionClusterResponse {
	s.Body = v
	return s
}

type DeleteUdfArtifactHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteUdfArtifactHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteUdfArtifactHeaders) GoString() string {
	return s.String()
}

func (s *DeleteUdfArtifactHeaders) SetCommonHeaders(v map[string]*string) *DeleteUdfArtifactHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteUdfArtifactHeaders) SetWorkspace(v string) *DeleteUdfArtifactHeaders {
	s.Workspace = &v
	return s
}

type DeleteUdfArtifactResponseBody struct {
	Data *DeleteUdfArtifactResult `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-ABCF-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteUdfArtifactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUdfArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUdfArtifactResponseBody) SetData(v *DeleteUdfArtifactResult) *DeleteUdfArtifactResponseBody {
	s.Data = v
	return s
}

func (s *DeleteUdfArtifactResponseBody) SetErrorCode(v string) *DeleteUdfArtifactResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteUdfArtifactResponseBody) SetErrorMessage(v string) *DeleteUdfArtifactResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteUdfArtifactResponseBody) SetHttpCode(v int32) *DeleteUdfArtifactResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteUdfArtifactResponseBody) SetRequestId(v string) *DeleteUdfArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUdfArtifactResponseBody) SetSuccess(v bool) *DeleteUdfArtifactResponseBody {
	s.Success = &v
	return s
}

type DeleteUdfArtifactResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUdfArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUdfArtifactResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUdfArtifactResponse) GoString() string {
	return s.String()
}

func (s *DeleteUdfArtifactResponse) SetHeaders(v map[string]*string) *DeleteUdfArtifactResponse {
	s.Headers = v
	return s
}

func (s *DeleteUdfArtifactResponse) SetStatusCode(v int32) *DeleteUdfArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUdfArtifactResponse) SetBody(v *DeleteUdfArtifactResponseBody) *DeleteUdfArtifactResponse {
	s.Body = v
	return s
}

type DeleteUdfFunctionHeaders struct {
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

func (s DeleteUdfFunctionHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteUdfFunctionHeaders) GoString() string {
	return s.String()
}

func (s *DeleteUdfFunctionHeaders) SetCommonHeaders(v map[string]*string) *DeleteUdfFunctionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteUdfFunctionHeaders) SetWorkspace(v string) *DeleteUdfFunctionHeaders {
	s.Workspace = &v
	return s
}

type DeleteUdfFunctionRequest struct {
	// The name of the class that corresponds to the UDF.
	//
	// This parameter is required.
	//
	// example:
	//
	// Category
	ClassName *string `json:"className,omitempty" xml:"className,omitempty"`
	// The name of the resource that corresponds to the UDF that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-udf
	UdfArtifactName *string `json:"udfArtifactName,omitempty" xml:"udfArtifactName,omitempty"`
}

func (s DeleteUdfFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUdfFunctionRequest) GoString() string {
	return s.String()
}

func (s *DeleteUdfFunctionRequest) SetClassName(v string) *DeleteUdfFunctionRequest {
	s.ClassName = &v
	return s
}

func (s *DeleteUdfFunctionRequest) SetUdfArtifactName(v string) *DeleteUdfFunctionRequest {
	s.UdfArtifactName = &v
	return s
}

type DeleteUdfFunctionResponseBody struct {
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

func (s DeleteUdfFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUdfFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUdfFunctionResponseBody) SetErrorCode(v string) *DeleteUdfFunctionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteUdfFunctionResponseBody) SetErrorMessage(v string) *DeleteUdfFunctionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteUdfFunctionResponseBody) SetHttpCode(v int32) *DeleteUdfFunctionResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteUdfFunctionResponseBody) SetRequestId(v string) *DeleteUdfFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUdfFunctionResponseBody) SetSuccess(v bool) *DeleteUdfFunctionResponseBody {
	s.Success = &v
	return s
}

type DeleteUdfFunctionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUdfFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUdfFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUdfFunctionResponse) GoString() string {
	return s.String()
}

func (s *DeleteUdfFunctionResponse) SetHeaders(v map[string]*string) *DeleteUdfFunctionResponse {
	s.Headers = v
	return s
}

func (s *DeleteUdfFunctionResponse) SetStatusCode(v int32) *DeleteUdfFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUdfFunctionResponse) SetBody(v *DeleteUdfFunctionResponseBody) *DeleteUdfFunctionResponse {
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

type DeployDeploymentDraftAsyncHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeployDeploymentDraftAsyncHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeployDeploymentDraftAsyncHeaders) GoString() string {
	return s.String()
}

func (s *DeployDeploymentDraftAsyncHeaders) SetCommonHeaders(v map[string]*string) *DeployDeploymentDraftAsyncHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeployDeploymentDraftAsyncHeaders) SetWorkspace(v string) *DeployDeploymentDraftAsyncHeaders {
	s.Workspace = &v
	return s
}

type DeployDeploymentDraftAsyncRequest struct {
	// This parameter is required.
	Body *DraftDeployParams `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployDeploymentDraftAsyncRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployDeploymentDraftAsyncRequest) GoString() string {
	return s.String()
}

func (s *DeployDeploymentDraftAsyncRequest) SetBody(v *DraftDeployParams) *DeployDeploymentDraftAsyncRequest {
	s.Body = v
	return s
}

type DeployDeploymentDraftAsyncResponseBody struct {
	Data *DeployDeploymentDraftAsyncResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployDeploymentDraftAsyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployDeploymentDraftAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *DeployDeploymentDraftAsyncResponseBody) SetData(v *DeployDeploymentDraftAsyncResponseBodyData) *DeployDeploymentDraftAsyncResponseBody {
	s.Data = v
	return s
}

func (s *DeployDeploymentDraftAsyncResponseBody) SetErrorCode(v string) *DeployDeploymentDraftAsyncResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeployDeploymentDraftAsyncResponseBody) SetErrorMessage(v string) *DeployDeploymentDraftAsyncResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeployDeploymentDraftAsyncResponseBody) SetHttpCode(v int32) *DeployDeploymentDraftAsyncResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeployDeploymentDraftAsyncResponseBody) SetRequestId(v string) *DeployDeploymentDraftAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployDeploymentDraftAsyncResponseBody) SetSuccess(v bool) *DeployDeploymentDraftAsyncResponseBody {
	s.Success = &v
	return s
}

type DeployDeploymentDraftAsyncResponseBodyData struct {
	// example:
	//
	// b3dcdb25-bf36-457d-92ba-a36077e8****
	TicketId *string `json:"ticketId,omitempty" xml:"ticketId,omitempty"`
}

func (s DeployDeploymentDraftAsyncResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeployDeploymentDraftAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeployDeploymentDraftAsyncResponseBodyData) SetTicketId(v string) *DeployDeploymentDraftAsyncResponseBodyData {
	s.TicketId = &v
	return s
}

type DeployDeploymentDraftAsyncResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployDeploymentDraftAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployDeploymentDraftAsyncResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployDeploymentDraftAsyncResponse) GoString() string {
	return s.String()
}

func (s *DeployDeploymentDraftAsyncResponse) SetHeaders(v map[string]*string) *DeployDeploymentDraftAsyncResponse {
	s.Headers = v
	return s
}

func (s *DeployDeploymentDraftAsyncResponse) SetStatusCode(v int32) *DeployDeploymentDraftAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployDeploymentDraftAsyncResponse) SetBody(v *DeployDeploymentDraftAsyncResponseBody) *DeployDeploymentDraftAsyncResponse {
	s.Body = v
	return s
}

type ExecuteSqlStatementHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 710d6a64d8c34d
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ExecuteSqlStatementHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExecuteSqlStatementHeaders) GoString() string {
	return s.String()
}

func (s *ExecuteSqlStatementHeaders) SetCommonHeaders(v map[string]*string) *ExecuteSqlStatementHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExecuteSqlStatementHeaders) SetWorkspace(v string) *ExecuteSqlStatementHeaders {
	s.Workspace = &v
	return s
}

type ExecuteSqlStatementRequest struct {
	Body *SqlStatementWithContext `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteSqlStatementRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteSqlStatementRequest) GoString() string {
	return s.String()
}

func (s *ExecuteSqlStatementRequest) SetBody(v *SqlStatementWithContext) *ExecuteSqlStatementRequest {
	s.Body = v
	return s
}

type ExecuteSqlStatementResponseBody struct {
	Data *SqlStatementExecuteResult `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteSqlStatementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteSqlStatementResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteSqlStatementResponseBody) SetData(v *SqlStatementExecuteResult) *ExecuteSqlStatementResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteSqlStatementResponseBody) SetErrorCode(v string) *ExecuteSqlStatementResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ExecuteSqlStatementResponseBody) SetErrorMessage(v string) *ExecuteSqlStatementResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ExecuteSqlStatementResponseBody) SetHttpCode(v int32) *ExecuteSqlStatementResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ExecuteSqlStatementResponseBody) SetRequestId(v string) *ExecuteSqlStatementResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteSqlStatementResponseBody) SetSuccess(v bool) *ExecuteSqlStatementResponseBody {
	s.Success = &v
	return s
}

type ExecuteSqlStatementResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteSqlStatementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteSqlStatementResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteSqlStatementResponse) GoString() string {
	return s.String()
}

func (s *ExecuteSqlStatementResponse) SetHeaders(v map[string]*string) *ExecuteSqlStatementResponse {
	s.Headers = v
	return s
}

func (s *ExecuteSqlStatementResponse) SetStatusCode(v int32) *ExecuteSqlStatementResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteSqlStatementResponse) SetBody(v *ExecuteSqlStatementResponseBody) *ExecuteSqlStatementResponse {
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

type GetAppliedScheduledPlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetAppliedScheduledPlanHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAppliedScheduledPlanHeaders) GoString() string {
	return s.String()
}

func (s *GetAppliedScheduledPlanHeaders) SetCommonHeaders(v map[string]*string) *GetAppliedScheduledPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAppliedScheduledPlanHeaders) SetWorkspace(v string) *GetAppliedScheduledPlanHeaders {
	s.Workspace = &v
	return s
}

type GetAppliedScheduledPlanRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5737ef81-d2f1-49cf-8752-30910809****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
}

func (s GetAppliedScheduledPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppliedScheduledPlanRequest) GoString() string {
	return s.String()
}

func (s *GetAppliedScheduledPlanRequest) SetDeploymentId(v string) *GetAppliedScheduledPlanRequest {
	s.DeploymentId = &v
	return s
}

type GetAppliedScheduledPlanResponseBody struct {
	Data *ScheduledPlanAppliedInfo `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetAppliedScheduledPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppliedScheduledPlanResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppliedScheduledPlanResponseBody) SetData(v *ScheduledPlanAppliedInfo) *GetAppliedScheduledPlanResponseBody {
	s.Data = v
	return s
}

func (s *GetAppliedScheduledPlanResponseBody) SetErrorCode(v string) *GetAppliedScheduledPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAppliedScheduledPlanResponseBody) SetErrorMessage(v string) *GetAppliedScheduledPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAppliedScheduledPlanResponseBody) SetHttpCode(v int32) *GetAppliedScheduledPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetAppliedScheduledPlanResponseBody) SetRequestId(v string) *GetAppliedScheduledPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppliedScheduledPlanResponseBody) SetSuccess(v bool) *GetAppliedScheduledPlanResponseBody {
	s.Success = &v
	return s
}

type GetAppliedScheduledPlanResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppliedScheduledPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppliedScheduledPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppliedScheduledPlanResponse) GoString() string {
	return s.String()
}

func (s *GetAppliedScheduledPlanResponse) SetHeaders(v map[string]*string) *GetAppliedScheduledPlanResponse {
	s.Headers = v
	return s
}

func (s *GetAppliedScheduledPlanResponse) SetStatusCode(v int32) *GetAppliedScheduledPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppliedScheduledPlanResponse) SetBody(v *GetAppliedScheduledPlanResponseBody) *GetAppliedScheduledPlanResponse {
	s.Body = v
	return s
}

type GetCatalogsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetCatalogsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogsHeaders) GoString() string {
	return s.String()
}

func (s *GetCatalogsHeaders) SetCommonHeaders(v map[string]*string) *GetCatalogsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCatalogsHeaders) SetWorkspace(v string) *GetCatalogsHeaders {
	s.Workspace = &v
	return s
}

type GetCatalogsRequest struct {
	// example:
	//
	// paimon
	CatalogName *string `json:"catalogName,omitempty" xml:"catalogName,omitempty"`
}

func (s GetCatalogsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogsRequest) GoString() string {
	return s.String()
}

func (s *GetCatalogsRequest) SetCatalogName(v string) *GetCatalogsRequest {
	s.CatalogName = &v
	return s
}

type GetCatalogsResponseBody struct {
	Data []*Catalog `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-ABCD-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetCatalogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCatalogsResponseBody) SetData(v []*Catalog) *GetCatalogsResponseBody {
	s.Data = v
	return s
}

func (s *GetCatalogsResponseBody) SetErrorCode(v string) *GetCatalogsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetCatalogsResponseBody) SetErrorMessage(v string) *GetCatalogsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCatalogsResponseBody) SetHttpCode(v int32) *GetCatalogsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetCatalogsResponseBody) SetRequestId(v string) *GetCatalogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCatalogsResponseBody) SetSuccess(v bool) *GetCatalogsResponseBody {
	s.Success = &v
	return s
}

type GetCatalogsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCatalogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCatalogsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogsResponse) GoString() string {
	return s.String()
}

func (s *GetCatalogsResponse) SetHeaders(v map[string]*string) *GetCatalogsResponse {
	s.Headers = v
	return s
}

func (s *GetCatalogsResponse) SetStatusCode(v int32) *GetCatalogsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCatalogsResponse) SetBody(v *GetCatalogsResponseBody) *GetCatalogsResponse {
	s.Body = v
	return s
}

type GetDatabasesHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetDatabasesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDatabasesHeaders) GoString() string {
	return s.String()
}

func (s *GetDatabasesHeaders) SetCommonHeaders(v map[string]*string) *GetDatabasesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDatabasesHeaders) SetWorkspace(v string) *GetDatabasesHeaders {
	s.Workspace = &v
	return s
}

type GetDatabasesRequest struct {
	// example:
	//
	// paimon-ods
	DatabaseName *string `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
}

func (s GetDatabasesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDatabasesRequest) GoString() string {
	return s.String()
}

func (s *GetDatabasesRequest) SetDatabaseName(v string) *GetDatabasesRequest {
	s.DatabaseName = &v
	return s
}

type GetDatabasesResponseBody struct {
	Data []*Database `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDatabasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatabasesResponseBody) SetData(v []*Database) *GetDatabasesResponseBody {
	s.Data = v
	return s
}

func (s *GetDatabasesResponseBody) SetErrorCode(v string) *GetDatabasesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDatabasesResponseBody) SetErrorMessage(v string) *GetDatabasesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDatabasesResponseBody) SetHttpCode(v int32) *GetDatabasesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDatabasesResponseBody) SetRequestId(v string) *GetDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatabasesResponseBody) SetSuccess(v bool) *GetDatabasesResponseBody {
	s.Success = &v
	return s
}

type GetDatabasesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatabasesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDatabasesResponse) GoString() string {
	return s.String()
}

func (s *GetDatabasesResponse) SetHeaders(v map[string]*string) *GetDatabasesResponse {
	s.Headers = v
	return s
}

func (s *GetDatabasesResponse) SetStatusCode(v int32) *GetDatabasesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatabasesResponse) SetBody(v *GetDatabasesResponseBody) *GetDatabasesResponse {
	s.Body = v
	return s
}

type GetDeployDeploymentDraftResultHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 710d6a64d8****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetDeployDeploymentDraftResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDeployDeploymentDraftResultHeaders) GoString() string {
	return s.String()
}

func (s *GetDeployDeploymentDraftResultHeaders) SetCommonHeaders(v map[string]*string) *GetDeployDeploymentDraftResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeployDeploymentDraftResultHeaders) SetWorkspace(v string) *GetDeployDeploymentDraftResultHeaders {
	s.Workspace = &v
	return s
}

type GetDeployDeploymentDraftResultResponseBody struct {
	Data *AsyncDraftDeployResult `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDeployDeploymentDraftResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeployDeploymentDraftResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeployDeploymentDraftResultResponseBody) SetData(v *AsyncDraftDeployResult) *GetDeployDeploymentDraftResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDeployDeploymentDraftResultResponseBody) SetErrorCode(v string) *GetDeployDeploymentDraftResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDeployDeploymentDraftResultResponseBody) SetErrorMessage(v string) *GetDeployDeploymentDraftResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDeployDeploymentDraftResultResponseBody) SetHttpCode(v int32) *GetDeployDeploymentDraftResultResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDeployDeploymentDraftResultResponseBody) SetRequestId(v string) *GetDeployDeploymentDraftResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeployDeploymentDraftResultResponseBody) SetSuccess(v bool) *GetDeployDeploymentDraftResultResponseBody {
	s.Success = &v
	return s
}

type GetDeployDeploymentDraftResultResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeployDeploymentDraftResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeployDeploymentDraftResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeployDeploymentDraftResultResponse) GoString() string {
	return s.String()
}

func (s *GetDeployDeploymentDraftResultResponse) SetHeaders(v map[string]*string) *GetDeployDeploymentDraftResultResponse {
	s.Headers = v
	return s
}

func (s *GetDeployDeploymentDraftResultResponse) SetStatusCode(v int32) *GetDeployDeploymentDraftResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeployDeploymentDraftResultResponse) SetBody(v *GetDeployDeploymentDraftResultResponseBody) *GetDeployDeploymentDraftResultResponse {
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

type GetDeploymentDraftHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetDeploymentDraftHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentDraftHeaders) GoString() string {
	return s.String()
}

func (s *GetDeploymentDraftHeaders) SetCommonHeaders(v map[string]*string) *GetDeploymentDraftHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeploymentDraftHeaders) SetWorkspace(v string) *GetDeploymentDraftHeaders {
	s.Workspace = &v
	return s
}

type GetDeploymentDraftResponseBody struct {
	Data *DeploymentDraft `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDeploymentDraftResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentDraftResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeploymentDraftResponseBody) SetData(v *DeploymentDraft) *GetDeploymentDraftResponseBody {
	s.Data = v
	return s
}

func (s *GetDeploymentDraftResponseBody) SetErrorCode(v string) *GetDeploymentDraftResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDeploymentDraftResponseBody) SetErrorMessage(v string) *GetDeploymentDraftResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDeploymentDraftResponseBody) SetHttpCode(v int32) *GetDeploymentDraftResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDeploymentDraftResponseBody) SetRequestId(v string) *GetDeploymentDraftResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeploymentDraftResponseBody) SetSuccess(v bool) *GetDeploymentDraftResponseBody {
	s.Success = &v
	return s
}

type GetDeploymentDraftResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeploymentDraftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeploymentDraftResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentDraftResponse) GoString() string {
	return s.String()
}

func (s *GetDeploymentDraftResponse) SetHeaders(v map[string]*string) *GetDeploymentDraftResponse {
	s.Headers = v
	return s
}

func (s *GetDeploymentDraftResponse) SetStatusCode(v int32) *GetDeploymentDraftResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeploymentDraftResponse) SetBody(v *GetDeploymentDraftResponseBody) *GetDeploymentDraftResponse {
	s.Body = v
	return s
}

type GetDeploymentDraftLockHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetDeploymentDraftLockHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentDraftLockHeaders) GoString() string {
	return s.String()
}

func (s *GetDeploymentDraftLockHeaders) SetCommonHeaders(v map[string]*string) *GetDeploymentDraftLockHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeploymentDraftLockHeaders) SetWorkspace(v string) *GetDeploymentDraftLockHeaders {
	s.Workspace = &v
	return s
}

type GetDeploymentDraftLockRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c84d73be-40ad-4627-8bdd-fa1eba51b234
	DeploymentDraftId *string `json:"deploymentDraftId,omitempty" xml:"deploymentDraftId,omitempty"`
}

func (s GetDeploymentDraftLockRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentDraftLockRequest) GoString() string {
	return s.String()
}

func (s *GetDeploymentDraftLockRequest) SetDeploymentDraftId(v string) *GetDeploymentDraftLockRequest {
	s.DeploymentDraftId = &v
	return s
}

type GetDeploymentDraftLockResponseBody struct {
	Data *Lock `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDeploymentDraftLockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentDraftLockResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeploymentDraftLockResponseBody) SetData(v *Lock) *GetDeploymentDraftLockResponseBody {
	s.Data = v
	return s
}

func (s *GetDeploymentDraftLockResponseBody) SetErrorCode(v string) *GetDeploymentDraftLockResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDeploymentDraftLockResponseBody) SetErrorMessage(v string) *GetDeploymentDraftLockResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDeploymentDraftLockResponseBody) SetHttpCode(v int32) *GetDeploymentDraftLockResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDeploymentDraftLockResponseBody) SetRequestId(v string) *GetDeploymentDraftLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeploymentDraftLockResponseBody) SetSuccess(v bool) *GetDeploymentDraftLockResponseBody {
	s.Success = &v
	return s
}

type GetDeploymentDraftLockResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeploymentDraftLockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeploymentDraftLockResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentDraftLockResponse) GoString() string {
	return s.String()
}

func (s *GetDeploymentDraftLockResponse) SetHeaders(v map[string]*string) *GetDeploymentDraftLockResponse {
	s.Headers = v
	return s
}

func (s *GetDeploymentDraftLockResponse) SetStatusCode(v int32) *GetDeploymentDraftLockResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeploymentDraftLockResponse) SetBody(v *GetDeploymentDraftLockResponseBody) *GetDeploymentDraftLockResponse {
	s.Body = v
	return s
}

type GetEventsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetEventsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetEventsHeaders) GoString() string {
	return s.String()
}

func (s *GetEventsHeaders) SetCommonHeaders(v map[string]*string) *GetEventsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetEventsHeaders) SetWorkspace(v string) *GetEventsHeaders {
	s.Workspace = &v
	return s
}

type GetEventsRequest struct {
	// example:
	//
	// 58718c99-3b29-4c5e-93bb-c9fc4ec6****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEventsRequest) GoString() string {
	return s.String()
}

func (s *GetEventsRequest) SetDeploymentId(v string) *GetEventsRequest {
	s.DeploymentId = &v
	return s
}

func (s *GetEventsRequest) SetPageIndex(v int32) *GetEventsRequest {
	s.PageIndex = &v
	return s
}

func (s *GetEventsRequest) SetPageSize(v int32) *GetEventsRequest {
	s.PageSize = &v
	return s
}

type GetEventsResponseBody struct {
	Data []*Event `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 4
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s GetEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEventsResponseBody) GoString() string {
	return s.String()
}

func (s *GetEventsResponseBody) SetData(v []*Event) *GetEventsResponseBody {
	s.Data = v
	return s
}

func (s *GetEventsResponseBody) SetErrorCode(v string) *GetEventsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetEventsResponseBody) SetErrorMessage(v string) *GetEventsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetEventsResponseBody) SetHttpCode(v int32) *GetEventsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetEventsResponseBody) SetPageIndex(v int32) *GetEventsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetEventsResponseBody) SetPageSize(v int32) *GetEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetEventsResponseBody) SetRequestId(v string) *GetEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEventsResponseBody) SetSuccess(v bool) *GetEventsResponseBody {
	s.Success = &v
	return s
}

func (s *GetEventsResponseBody) SetTotalSize(v int32) *GetEventsResponseBody {
	s.TotalSize = &v
	return s
}

type GetEventsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEventsResponse) GoString() string {
	return s.String()
}

func (s *GetEventsResponse) SetHeaders(v map[string]*string) *GetEventsResponse {
	s.Headers = v
	return s
}

func (s *GetEventsResponse) SetStatusCode(v int32) *GetEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEventsResponse) SetBody(v *GetEventsResponseBody) *GetEventsResponse {
	s.Body = v
	return s
}

type GetFolderHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetFolderHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFolderHeaders) GoString() string {
	return s.String()
}

func (s *GetFolderHeaders) SetCommonHeaders(v map[string]*string) *GetFolderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFolderHeaders) SetWorkspace(v string) *GetFolderHeaders {
	s.Workspace = &v
	return s
}

type GetFolderRequest struct {
	// example:
	//
	// 89097
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
}

func (s GetFolderRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFolderRequest) GoString() string {
	return s.String()
}

func (s *GetFolderRequest) SetFolderId(v string) *GetFolderRequest {
	s.FolderId = &v
	return s
}

type GetFolderResponseBody struct {
	// The data structure of the folder.
	Data *Folder `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFolderResponseBody) GoString() string {
	return s.String()
}

func (s *GetFolderResponseBody) SetData(v *Folder) *GetFolderResponseBody {
	s.Data = v
	return s
}

func (s *GetFolderResponseBody) SetErrorCode(v string) *GetFolderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetFolderResponseBody) SetErrorMessage(v string) *GetFolderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetFolderResponseBody) SetHttpCode(v int32) *GetFolderResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetFolderResponseBody) SetRequestId(v string) *GetFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFolderResponseBody) SetSuccess(v bool) *GetFolderResponseBody {
	s.Success = &v
	return s
}

type GetFolderResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFolderResponse) GoString() string {
	return s.String()
}

func (s *GetFolderResponse) SetHeaders(v map[string]*string) *GetFolderResponse {
	s.Headers = v
	return s
}

func (s *GetFolderResponse) SetStatusCode(v int32) *GetFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFolderResponse) SetBody(v *GetFolderResponseBody) *GetFolderResponse {
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

type GetHotUpdateJobResultHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetHotUpdateJobResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHotUpdateJobResultHeaders) GoString() string {
	return s.String()
}

func (s *GetHotUpdateJobResultHeaders) SetCommonHeaders(v map[string]*string) *GetHotUpdateJobResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotUpdateJobResultHeaders) SetWorkspace(v string) *GetHotUpdateJobResultHeaders {
	s.Workspace = &v
	return s
}

type GetHotUpdateJobResultResponseBody struct {
	Data *HotUpdateJobResult `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-ABCF-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetHotUpdateJobResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotUpdateJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotUpdateJobResultResponseBody) SetData(v *HotUpdateJobResult) *GetHotUpdateJobResultResponseBody {
	s.Data = v
	return s
}

func (s *GetHotUpdateJobResultResponseBody) SetErrorCode(v string) *GetHotUpdateJobResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetHotUpdateJobResultResponseBody) SetErrorMessage(v string) *GetHotUpdateJobResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetHotUpdateJobResultResponseBody) SetHttpCode(v int32) *GetHotUpdateJobResultResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetHotUpdateJobResultResponseBody) SetRequestId(v string) *GetHotUpdateJobResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotUpdateJobResultResponseBody) SetSuccess(v bool) *GetHotUpdateJobResultResponseBody {
	s.Success = &v
	return s
}

type GetHotUpdateJobResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotUpdateJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotUpdateJobResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotUpdateJobResultResponse) GoString() string {
	return s.String()
}

func (s *GetHotUpdateJobResultResponse) SetHeaders(v map[string]*string) *GetHotUpdateJobResultResponse {
	s.Headers = v
	return s
}

func (s *GetHotUpdateJobResultResponse) SetStatusCode(v int32) *GetHotUpdateJobResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotUpdateJobResultResponse) SetBody(v *GetHotUpdateJobResultResponseBody) *GetHotUpdateJobResultResponse {
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
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
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

func (s *GetJobResponseBody) SetAccessDeniedDetail(v string) *GetJobResponseBody {
	s.AccessDeniedDetail = &v
	return s
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

type GetJobDiagnosisHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetJobDiagnosisHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetJobDiagnosisHeaders) GoString() string {
	return s.String()
}

func (s *GetJobDiagnosisHeaders) SetCommonHeaders(v map[string]*string) *GetJobDiagnosisHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetJobDiagnosisHeaders) SetWorkspace(v string) *GetJobDiagnosisHeaders {
	s.Workspace = &v
	return s
}

type GetJobDiagnosisResponseBody struct {
	// example:
	//
	// “”
	AccessDeniedDetail *string       `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	Data               *JobDiagnosis `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetJobDiagnosisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobDiagnosisResponseBody) SetAccessDeniedDetail(v string) *GetJobDiagnosisResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetJobDiagnosisResponseBody) SetData(v *JobDiagnosis) *GetJobDiagnosisResponseBody {
	s.Data = v
	return s
}

func (s *GetJobDiagnosisResponseBody) SetErrorCode(v string) *GetJobDiagnosisResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetJobDiagnosisResponseBody) SetErrorMessage(v string) *GetJobDiagnosisResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetJobDiagnosisResponseBody) SetHttpCode(v int32) *GetJobDiagnosisResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetJobDiagnosisResponseBody) SetRequestId(v string) *GetJobDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobDiagnosisResponseBody) SetSuccess(v bool) *GetJobDiagnosisResponseBody {
	s.Success = &v
	return s
}

type GetJobDiagnosisResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobDiagnosisResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *GetJobDiagnosisResponse) SetHeaders(v map[string]*string) *GetJobDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *GetJobDiagnosisResponse) SetStatusCode(v int32) *GetJobDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobDiagnosisResponse) SetBody(v *GetJobDiagnosisResponseBody) *GetJobDiagnosisResponse {
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

type GetLineageInfoHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 710d6a64d8****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetLineageInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetLineageInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetLineageInfoHeaders) SetCommonHeaders(v map[string]*string) *GetLineageInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetLineageInfoHeaders) SetWorkspace(v string) *GetLineageInfoHeaders {
	s.Workspace = &v
	return s
}

type GetLineageInfoRequest struct {
	// The parameters about the lineage information.
	Body *GetLineageInfoParams `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLineageInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLineageInfoRequest) GoString() string {
	return s.String()
}

func (s *GetLineageInfoRequest) SetBody(v *GetLineageInfoParams) *GetLineageInfoRequest {
	s.Body = v
	return s
}

type GetLineageInfoResponseBody struct {
	// The lineage information.
	Data *LineageInfo `json:"data,omitempty" xml:"data,omitempty"`
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

func (s GetLineageInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLineageInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetLineageInfoResponseBody) SetData(v *LineageInfo) *GetLineageInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetLineageInfoResponseBody) SetErrorCode(v string) *GetLineageInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetLineageInfoResponseBody) SetErrorMessage(v string) *GetLineageInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetLineageInfoResponseBody) SetHttpCode(v int32) *GetLineageInfoResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetLineageInfoResponseBody) SetRequestId(v string) *GetLineageInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLineageInfoResponseBody) SetSuccess(v bool) *GetLineageInfoResponseBody {
	s.Success = &v
	return s
}

type GetLineageInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLineageInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLineageInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLineageInfoResponse) GoString() string {
	return s.String()
}

func (s *GetLineageInfoResponse) SetHeaders(v map[string]*string) *GetLineageInfoResponse {
	s.Headers = v
	return s
}

func (s *GetLineageInfoResponse) SetStatusCode(v int32) *GetLineageInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLineageInfoResponse) SetBody(v *GetLineageInfoResponseBody) *GetLineageInfoResponse {
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

type GetSessionClusterHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetSessionClusterHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSessionClusterHeaders) GoString() string {
	return s.String()
}

func (s *GetSessionClusterHeaders) SetCommonHeaders(v map[string]*string) *GetSessionClusterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSessionClusterHeaders) SetWorkspace(v string) *GetSessionClusterHeaders {
	s.Workspace = &v
	return s
}

type GetSessionClusterResponseBody struct {
	Data *SessionCluster `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSessionClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSessionClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetSessionClusterResponseBody) SetData(v *SessionCluster) *GetSessionClusterResponseBody {
	s.Data = v
	return s
}

func (s *GetSessionClusterResponseBody) SetErrorCode(v string) *GetSessionClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSessionClusterResponseBody) SetErrorMessage(v string) *GetSessionClusterResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetSessionClusterResponseBody) SetHttpCode(v int32) *GetSessionClusterResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetSessionClusterResponseBody) SetRequestId(v string) *GetSessionClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSessionClusterResponseBody) SetSuccess(v bool) *GetSessionClusterResponseBody {
	s.Success = &v
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

type GetTablesHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetTablesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTablesHeaders) GoString() string {
	return s.String()
}

func (s *GetTablesHeaders) SetCommonHeaders(v map[string]*string) *GetTablesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTablesHeaders) SetWorkspace(v string) *GetTablesHeaders {
	s.Workspace = &v
	return s
}

type GetTablesRequest struct {
	// example:
	//
	// item
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
}

func (s GetTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTablesRequest) GoString() string {
	return s.String()
}

func (s *GetTablesRequest) SetTableName(v string) *GetTablesRequest {
	s.TableName = &v
	return s
}

type GetTablesResponseBody struct {
	// If the value of success was true, the list and details of tables that meet the condition were returned. If the value of success was false, a null value was returned.
	Data []*Table `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// ECE641B2-AB0B-4174-9C3B-885881558637
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTablesResponseBody) GoString() string {
	return s.String()
}

func (s *GetTablesResponseBody) SetData(v []*Table) *GetTablesResponseBody {
	s.Data = v
	return s
}

func (s *GetTablesResponseBody) SetErrorCode(v string) *GetTablesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTablesResponseBody) SetErrorMessage(v string) *GetTablesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTablesResponseBody) SetHttpCode(v int32) *GetTablesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetTablesResponseBody) SetRequestId(v string) *GetTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTablesResponseBody) SetSuccess(v bool) *GetTablesResponseBody {
	s.Success = &v
	return s
}

type GetTablesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTablesResponse) GoString() string {
	return s.String()
}

func (s *GetTablesResponse) SetHeaders(v map[string]*string) *GetTablesResponse {
	s.Headers = v
	return s
}

func (s *GetTablesResponse) SetStatusCode(v int32) *GetTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTablesResponse) SetBody(v *GetTablesResponseBody) *GetTablesResponse {
	s.Body = v
	return s
}

type GetUdfArtifactsHeaders struct {
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

func (s GetUdfArtifactsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUdfArtifactsHeaders) GoString() string {
	return s.String()
}

func (s *GetUdfArtifactsHeaders) SetCommonHeaders(v map[string]*string) *GetUdfArtifactsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUdfArtifactsHeaders) SetWorkspace(v string) *GetUdfArtifactsHeaders {
	s.Workspace = &v
	return s
}

type GetUdfArtifactsRequest struct {
	// The name of the JAR or Python file that corresponds to the UDF.
	//
	// example:
	//
	// test-udf
	UdfArtifactName *string `json:"udfArtifactName,omitempty" xml:"udfArtifactName,omitempty"`
}

func (s GetUdfArtifactsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUdfArtifactsRequest) GoString() string {
	return s.String()
}

func (s *GetUdfArtifactsRequest) SetUdfArtifactName(v string) *GetUdfArtifactsRequest {
	s.UdfArtifactName = &v
	return s
}

type GetUdfArtifactsResponseBody struct {
	// If the value of success was true, the details of the JAR or Python file that corresponds to the UDF were returned. If the value of success was false, a null value was returned.
	Data []*UdfArtifact `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s GetUdfArtifactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUdfArtifactsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUdfArtifactsResponseBody) SetData(v []*UdfArtifact) *GetUdfArtifactsResponseBody {
	s.Data = v
	return s
}

func (s *GetUdfArtifactsResponseBody) SetErrorCode(v string) *GetUdfArtifactsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUdfArtifactsResponseBody) SetErrorMessage(v string) *GetUdfArtifactsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetUdfArtifactsResponseBody) SetHttpCode(v int32) *GetUdfArtifactsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetUdfArtifactsResponseBody) SetRequestId(v string) *GetUdfArtifactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUdfArtifactsResponseBody) SetSuccess(v bool) *GetUdfArtifactsResponseBody {
	s.Success = &v
	return s
}

type GetUdfArtifactsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUdfArtifactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUdfArtifactsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUdfArtifactsResponse) GoString() string {
	return s.String()
}

func (s *GetUdfArtifactsResponse) SetHeaders(v map[string]*string) *GetUdfArtifactsResponse {
	s.Headers = v
	return s
}

func (s *GetUdfArtifactsResponse) SetStatusCode(v int32) *GetUdfArtifactsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUdfArtifactsResponse) SetBody(v *GetUdfArtifactsResponseBody) *GetUdfArtifactsResponse {
	s.Body = v
	return s
}

type HotUpdateJobHeaders struct {
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

func (s HotUpdateJobHeaders) String() string {
	return tea.Prettify(s)
}

func (s HotUpdateJobHeaders) GoString() string {
	return s.String()
}

func (s *HotUpdateJobHeaders) SetCommonHeaders(v map[string]*string) *HotUpdateJobHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotUpdateJobHeaders) SetWorkspace(v string) *HotUpdateJobHeaders {
	s.Workspace = &v
	return s
}

type HotUpdateJobResponseBody struct {
	// The dynamic update result.
	Data *HotUpdateJobResult `json:"data,omitempty" xml:"data,omitempty"`
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

func (s HotUpdateJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HotUpdateJobResponseBody) GoString() string {
	return s.String()
}

func (s *HotUpdateJobResponseBody) SetData(v *HotUpdateJobResult) *HotUpdateJobResponseBody {
	s.Data = v
	return s
}

func (s *HotUpdateJobResponseBody) SetErrorCode(v string) *HotUpdateJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *HotUpdateJobResponseBody) SetErrorMessage(v string) *HotUpdateJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *HotUpdateJobResponseBody) SetHttpCode(v int32) *HotUpdateJobResponseBody {
	s.HttpCode = &v
	return s
}

func (s *HotUpdateJobResponseBody) SetRequestId(v string) *HotUpdateJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotUpdateJobResponseBody) SetSuccess(v bool) *HotUpdateJobResponseBody {
	s.Success = &v
	return s
}

type HotUpdateJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotUpdateJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotUpdateJobResponse) String() string {
	return tea.Prettify(s)
}

func (s HotUpdateJobResponse) GoString() string {
	return s.String()
}

func (s *HotUpdateJobResponse) SetHeaders(v map[string]*string) *HotUpdateJobResponse {
	s.Headers = v
	return s
}

func (s *HotUpdateJobResponse) SetStatusCode(v int32) *HotUpdateJobResponse {
	s.StatusCode = &v
	return s
}

func (s *HotUpdateJobResponse) SetBody(v *HotUpdateJobResponseBody) *HotUpdateJobResponse {
	s.Body = v
	return s
}

type ListCustomConnectorsHeaders struct {
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

func (s ListCustomConnectorsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListCustomConnectorsHeaders) GoString() string {
	return s.String()
}

func (s *ListCustomConnectorsHeaders) SetCommonHeaders(v map[string]*string) *ListCustomConnectorsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCustomConnectorsHeaders) SetWorkspace(v string) *ListCustomConnectorsHeaders {
	s.Workspace = &v
	return s
}

type ListCustomConnectorsResponseBody struct {
	// If the value of success was true, the list of custom connectors in the namespace was returned. If the value of success was false, a null value was returned.
	Data []*Connector `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s ListCustomConnectorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCustomConnectorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomConnectorsResponseBody) SetData(v []*Connector) *ListCustomConnectorsResponseBody {
	s.Data = v
	return s
}

func (s *ListCustomConnectorsResponseBody) SetErrorCode(v string) *ListCustomConnectorsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListCustomConnectorsResponseBody) SetErrorMessage(v string) *ListCustomConnectorsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListCustomConnectorsResponseBody) SetHttpCode(v int32) *ListCustomConnectorsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListCustomConnectorsResponseBody) SetRequestId(v string) *ListCustomConnectorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomConnectorsResponseBody) SetSuccess(v bool) *ListCustomConnectorsResponseBody {
	s.Success = &v
	return s
}

type ListCustomConnectorsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomConnectorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomConnectorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomConnectorsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomConnectorsResponse) SetHeaders(v map[string]*string) *ListCustomConnectorsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomConnectorsResponse) SetStatusCode(v int32) *ListCustomConnectorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomConnectorsResponse) SetBody(v *ListCustomConnectorsResponseBody) *ListCustomConnectorsResponse {
	s.Body = v
	return s
}

type ListDeploymentDraftsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListDeploymentDraftsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentDraftsHeaders) GoString() string {
	return s.String()
}

func (s *ListDeploymentDraftsHeaders) SetCommonHeaders(v map[string]*string) *ListDeploymentDraftsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDeploymentDraftsHeaders) SetWorkspace(v string) *ListDeploymentDraftsHeaders {
	s.Workspace = &v
	return s
}

type ListDeploymentDraftsRequest struct {
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListDeploymentDraftsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentDraftsRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentDraftsRequest) SetPageIndex(v int32) *ListDeploymentDraftsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDeploymentDraftsRequest) SetPageSize(v int32) *ListDeploymentDraftsRequest {
	s.PageSize = &v
	return s
}

type ListDeploymentDraftsResponseBody struct {
	Data []*DeploymentDraft `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 69
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListDeploymentDraftsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentDraftsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeploymentDraftsResponseBody) SetData(v []*DeploymentDraft) *ListDeploymentDraftsResponseBody {
	s.Data = v
	return s
}

func (s *ListDeploymentDraftsResponseBody) SetErrorCode(v string) *ListDeploymentDraftsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDeploymentDraftsResponseBody) SetErrorMessage(v string) *ListDeploymentDraftsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDeploymentDraftsResponseBody) SetHttpCode(v int32) *ListDeploymentDraftsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListDeploymentDraftsResponseBody) SetPageIndex(v int32) *ListDeploymentDraftsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListDeploymentDraftsResponseBody) SetPageSize(v int32) *ListDeploymentDraftsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentDraftsResponseBody) SetRequestId(v string) *ListDeploymentDraftsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeploymentDraftsResponseBody) SetSuccess(v bool) *ListDeploymentDraftsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDeploymentDraftsResponseBody) SetTotalSize(v int32) *ListDeploymentDraftsResponseBody {
	s.TotalSize = &v
	return s
}

type ListDeploymentDraftsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeploymentDraftsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeploymentDraftsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentDraftsResponse) GoString() string {
	return s.String()
}

func (s *ListDeploymentDraftsResponse) SetHeaders(v map[string]*string) *ListDeploymentDraftsResponse {
	s.Headers = v
	return s
}

func (s *ListDeploymentDraftsResponse) SetStatusCode(v int32) *ListDeploymentDraftsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeploymentDraftsResponse) SetBody(v *ListDeploymentDraftsResponseBody) *ListDeploymentDraftsResponse {
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
	PageSize *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	SortName *string `json:"sortName,omitempty" xml:"sortName,omitempty"`
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

func (s *ListDeploymentsRequest) SetSortName(v string) *ListDeploymentsRequest {
	s.SortName = &v
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
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
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

func (s *ListJobsResponseBody) SetAccessDeniedDetail(v string) *ListJobsResponseBody {
	s.AccessDeniedDetail = &v
	return s
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

type ListScheduledPlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListScheduledPlanHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledPlanHeaders) GoString() string {
	return s.String()
}

func (s *ListScheduledPlanHeaders) SetCommonHeaders(v map[string]*string) *ListScheduledPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListScheduledPlanHeaders) SetWorkspace(v string) *ListScheduledPlanHeaders {
	s.Workspace = &v
	return s
}

type ListScheduledPlanRequest struct {
	// example:
	//
	// 737d0921-c5ac-47fc-9ba9-07a1e0b4****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListScheduledPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledPlanRequest) GoString() string {
	return s.String()
}

func (s *ListScheduledPlanRequest) SetDeploymentId(v string) *ListScheduledPlanRequest {
	s.DeploymentId = &v
	return s
}

func (s *ListScheduledPlanRequest) SetPageIndex(v int32) *ListScheduledPlanRequest {
	s.PageIndex = &v
	return s
}

func (s *ListScheduledPlanRequest) SetPageSize(v int32) *ListScheduledPlanRequest {
	s.PageSize = &v
	return s
}

type ListScheduledPlanResponseBody struct {
	Data []*ScheduledPlan `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 4
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListScheduledPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduledPlanResponseBody) SetData(v []*ScheduledPlan) *ListScheduledPlanResponseBody {
	s.Data = v
	return s
}

func (s *ListScheduledPlanResponseBody) SetErrorCode(v string) *ListScheduledPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListScheduledPlanResponseBody) SetErrorMessage(v string) *ListScheduledPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListScheduledPlanResponseBody) SetHttpCode(v int32) *ListScheduledPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListScheduledPlanResponseBody) SetPageIndex(v int32) *ListScheduledPlanResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListScheduledPlanResponseBody) SetPageSize(v int32) *ListScheduledPlanResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListScheduledPlanResponseBody) SetRequestId(v string) *ListScheduledPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduledPlanResponseBody) SetSuccess(v bool) *ListScheduledPlanResponseBody {
	s.Success = &v
	return s
}

func (s *ListScheduledPlanResponseBody) SetTotalSize(v int32) *ListScheduledPlanResponseBody {
	s.TotalSize = &v
	return s
}

type ListScheduledPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScheduledPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScheduledPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledPlanResponse) GoString() string {
	return s.String()
}

func (s *ListScheduledPlanResponse) SetHeaders(v map[string]*string) *ListScheduledPlanResponse {
	s.Headers = v
	return s
}

func (s *ListScheduledPlanResponse) SetStatusCode(v int32) *ListScheduledPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScheduledPlanResponse) SetBody(v *ListScheduledPlanResponseBody) *ListScheduledPlanResponse {
	s.Body = v
	return s
}

type ListScheduledPlanExecutedHistoryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListScheduledPlanExecutedHistoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledPlanExecutedHistoryHeaders) GoString() string {
	return s.String()
}

func (s *ListScheduledPlanExecutedHistoryHeaders) SetCommonHeaders(v map[string]*string) *ListScheduledPlanExecutedHistoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListScheduledPlanExecutedHistoryHeaders) SetWorkspace(v string) *ListScheduledPlanExecutedHistoryHeaders {
	s.Workspace = &v
	return s
}

type ListScheduledPlanExecutedHistoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 58718c99-3b29-4c5e-93bb-c9fc4ec6****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// example:
	//
	// SCHEDULED_PLAN
	Origin *string `json:"origin,omitempty" xml:"origin,omitempty"`
}

func (s ListScheduledPlanExecutedHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledPlanExecutedHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListScheduledPlanExecutedHistoryRequest) SetDeploymentId(v string) *ListScheduledPlanExecutedHistoryRequest {
	s.DeploymentId = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryRequest) SetOrigin(v string) *ListScheduledPlanExecutedHistoryRequest {
	s.Origin = &v
	return s
}

type ListScheduledPlanExecutedHistoryResponseBody struct {
	Data []*ScheduledPlanExecutedInfo `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success   *bool  `json:"success,omitempty" xml:"success,omitempty"`
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListScheduledPlanExecutedHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledPlanExecutedHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) SetData(v []*ScheduledPlanExecutedInfo) *ListScheduledPlanExecutedHistoryResponseBody {
	s.Data = v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) SetErrorCode(v string) *ListScheduledPlanExecutedHistoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) SetErrorMessage(v string) *ListScheduledPlanExecutedHistoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) SetHttpCode(v int32) *ListScheduledPlanExecutedHistoryResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) SetPageIndex(v int32) *ListScheduledPlanExecutedHistoryResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) SetPageSize(v int32) *ListScheduledPlanExecutedHistoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) SetRequestId(v string) *ListScheduledPlanExecutedHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) SetSuccess(v bool) *ListScheduledPlanExecutedHistoryResponseBody {
	s.Success = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) SetTotalSize(v int32) *ListScheduledPlanExecutedHistoryResponseBody {
	s.TotalSize = &v
	return s
}

type ListScheduledPlanExecutedHistoryResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScheduledPlanExecutedHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScheduledPlanExecutedHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledPlanExecutedHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListScheduledPlanExecutedHistoryResponse) SetHeaders(v map[string]*string) *ListScheduledPlanExecutedHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponse) SetStatusCode(v int32) *ListScheduledPlanExecutedHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponse) SetBody(v *ListScheduledPlanExecutedHistoryResponseBody) *ListScheduledPlanExecutedHistoryResponse {
	s.Body = v
	return s
}

type ListSessionClustersHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListSessionClustersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSessionClustersHeaders) GoString() string {
	return s.String()
}

func (s *ListSessionClustersHeaders) SetCommonHeaders(v map[string]*string) *ListSessionClustersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSessionClustersHeaders) SetWorkspace(v string) *ListSessionClustersHeaders {
	s.Workspace = &v
	return s
}

type ListSessionClustersResponseBody struct {
	Data []*SessionCluster `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-ABCD-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListSessionClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSessionClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListSessionClustersResponseBody) SetData(v []*SessionCluster) *ListSessionClustersResponseBody {
	s.Data = v
	return s
}

func (s *ListSessionClustersResponseBody) SetErrorCode(v string) *ListSessionClustersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSessionClustersResponseBody) SetErrorMessage(v string) *ListSessionClustersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSessionClustersResponseBody) SetHttpCode(v int32) *ListSessionClustersResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListSessionClustersResponseBody) SetRequestId(v string) *ListSessionClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSessionClustersResponseBody) SetSuccess(v bool) *ListSessionClustersResponseBody {
	s.Success = &v
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

type RegisterCustomConnectorHeaders struct {
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

func (s RegisterCustomConnectorHeaders) String() string {
	return tea.Prettify(s)
}

func (s RegisterCustomConnectorHeaders) GoString() string {
	return s.String()
}

func (s *RegisterCustomConnectorHeaders) SetCommonHeaders(v map[string]*string) *RegisterCustomConnectorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RegisterCustomConnectorHeaders) SetWorkspace(v string) *RegisterCustomConnectorHeaders {
	s.Workspace = &v
	return s
}

type RegisterCustomConnectorRequest struct {
	// The URL in which the JAR package of the custom connector is stored. The platform must be able to access this address.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://flink/connector/mysql123
	JarUrl *string `json:"jarUrl,omitempty" xml:"jarUrl,omitempty"`
}

func (s RegisterCustomConnectorRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterCustomConnectorRequest) GoString() string {
	return s.String()
}

func (s *RegisterCustomConnectorRequest) SetJarUrl(v string) *RegisterCustomConnectorRequest {
	s.JarUrl = &v
	return s
}

type RegisterCustomConnectorResponseBody struct {
	// If the value of success was true, a list of deployments in which custom connectors were deleted was returned. If the value of success was false, a null value was returned.
	Data []*Connector `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s RegisterCustomConnectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterCustomConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterCustomConnectorResponseBody) SetData(v []*Connector) *RegisterCustomConnectorResponseBody {
	s.Data = v
	return s
}

func (s *RegisterCustomConnectorResponseBody) SetErrorCode(v string) *RegisterCustomConnectorResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RegisterCustomConnectorResponseBody) SetErrorMessage(v string) *RegisterCustomConnectorResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RegisterCustomConnectorResponseBody) SetHttpCode(v int32) *RegisterCustomConnectorResponseBody {
	s.HttpCode = &v
	return s
}

func (s *RegisterCustomConnectorResponseBody) SetRequestId(v string) *RegisterCustomConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterCustomConnectorResponseBody) SetSuccess(v bool) *RegisterCustomConnectorResponseBody {
	s.Success = &v
	return s
}

type RegisterCustomConnectorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterCustomConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterCustomConnectorResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterCustomConnectorResponse) GoString() string {
	return s.String()
}

func (s *RegisterCustomConnectorResponse) SetHeaders(v map[string]*string) *RegisterCustomConnectorResponse {
	s.Headers = v
	return s
}

func (s *RegisterCustomConnectorResponse) SetStatusCode(v int32) *RegisterCustomConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterCustomConnectorResponse) SetBody(v *RegisterCustomConnectorResponseBody) *RegisterCustomConnectorResponse {
	s.Body = v
	return s
}

type RegisterUdfFunctionHeaders struct {
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

func (s RegisterUdfFunctionHeaders) String() string {
	return tea.Prettify(s)
}

func (s RegisterUdfFunctionHeaders) GoString() string {
	return s.String()
}

func (s *RegisterUdfFunctionHeaders) SetCommonHeaders(v map[string]*string) *RegisterUdfFunctionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RegisterUdfFunctionHeaders) SetWorkspace(v string) *RegisterUdfFunctionHeaders {
	s.Workspace = &v
	return s
}

type RegisterUdfFunctionRequest struct {
	// The name of the class that corresponds to the UDF.
	//
	// This parameter is required.
	//
	// example:
	//
	// orderRank
	ClassName *string `json:"className,omitempty" xml:"className,omitempty"`
	// The name of the UDF. In most cases, the name of the UDF is the same as the class name. You can specify a name for the UDF.
	//
	// This parameter is required.
	//
	// example:
	//
	// orderRank
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// The name of the JAR or Python file that corresponds to the UDF.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-udf
	UdfArtifactName *string `json:"udfArtifactName,omitempty" xml:"udfArtifactName,omitempty"`
}

func (s RegisterUdfFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterUdfFunctionRequest) GoString() string {
	return s.String()
}

func (s *RegisterUdfFunctionRequest) SetClassName(v string) *RegisterUdfFunctionRequest {
	s.ClassName = &v
	return s
}

func (s *RegisterUdfFunctionRequest) SetFunctionName(v string) *RegisterUdfFunctionRequest {
	s.FunctionName = &v
	return s
}

func (s *RegisterUdfFunctionRequest) SetUdfArtifactName(v string) *RegisterUdfFunctionRequest {
	s.UdfArtifactName = &v
	return s
}

type RegisterUdfFunctionResponseBody struct {
	// The information about the UDF.
	Data *UdfFunction `json:"data,omitempty" xml:"data,omitempty"`
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
	// CBC799F0-ABCD-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RegisterUdfFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterUdfFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterUdfFunctionResponseBody) SetData(v *UdfFunction) *RegisterUdfFunctionResponseBody {
	s.Data = v
	return s
}

func (s *RegisterUdfFunctionResponseBody) SetErrorCode(v string) *RegisterUdfFunctionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RegisterUdfFunctionResponseBody) SetErrorMessage(v string) *RegisterUdfFunctionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RegisterUdfFunctionResponseBody) SetHttpCode(v int32) *RegisterUdfFunctionResponseBody {
	s.HttpCode = &v
	return s
}

func (s *RegisterUdfFunctionResponseBody) SetRequestId(v string) *RegisterUdfFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterUdfFunctionResponseBody) SetSuccess(v bool) *RegisterUdfFunctionResponseBody {
	s.Success = &v
	return s
}

type RegisterUdfFunctionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterUdfFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterUdfFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterUdfFunctionResponse) GoString() string {
	return s.String()
}

func (s *RegisterUdfFunctionResponse) SetHeaders(v map[string]*string) *RegisterUdfFunctionResponse {
	s.Headers = v
	return s
}

func (s *RegisterUdfFunctionResponse) SetStatusCode(v int32) *RegisterUdfFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterUdfFunctionResponse) SetBody(v *RegisterUdfFunctionResponseBody) *RegisterUdfFunctionResponse {
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
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
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

func (s *StartJobResponseBody) SetAccessDeniedDetail(v string) *StartJobResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
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

func (s *StartJobWithParamsResponseBody) SetAccessDeniedDetail(v string) *StartJobWithParamsResponseBody {
	s.AccessDeniedDetail = &v
	return s
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

type StartSessionClusterHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s StartSessionClusterHeaders) String() string {
	return tea.Prettify(s)
}

func (s StartSessionClusterHeaders) GoString() string {
	return s.String()
}

func (s *StartSessionClusterHeaders) SetCommonHeaders(v map[string]*string) *StartSessionClusterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartSessionClusterHeaders) SetWorkspace(v string) *StartSessionClusterHeaders {
	s.Workspace = &v
	return s
}

type StartSessionClusterResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-ABCF-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StartSessionClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartSessionClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StartSessionClusterResponseBody) SetErrorCode(v string) *StartSessionClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StartSessionClusterResponseBody) SetErrorMessage(v string) *StartSessionClusterResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StartSessionClusterResponseBody) SetHttpCode(v int32) *StartSessionClusterResponseBody {
	s.HttpCode = &v
	return s
}

func (s *StartSessionClusterResponseBody) SetRequestId(v string) *StartSessionClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSessionClusterResponseBody) SetSuccess(v bool) *StartSessionClusterResponseBody {
	s.Success = &v
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

type StopApplyScheduledPlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s StopApplyScheduledPlanHeaders) String() string {
	return tea.Prettify(s)
}

func (s StopApplyScheduledPlanHeaders) GoString() string {
	return s.String()
}

func (s *StopApplyScheduledPlanHeaders) SetCommonHeaders(v map[string]*string) *StopApplyScheduledPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StopApplyScheduledPlanHeaders) SetWorkspace(v string) *StopApplyScheduledPlanHeaders {
	s.Workspace = &v
	return s
}

type StopApplyScheduledPlanResponseBody struct {
	Data *ScheduledPlanAppliedInfo `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StopApplyScheduledPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopApplyScheduledPlanResponseBody) GoString() string {
	return s.String()
}

func (s *StopApplyScheduledPlanResponseBody) SetData(v *ScheduledPlanAppliedInfo) *StopApplyScheduledPlanResponseBody {
	s.Data = v
	return s
}

func (s *StopApplyScheduledPlanResponseBody) SetErrorCode(v string) *StopApplyScheduledPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StopApplyScheduledPlanResponseBody) SetErrorMessage(v string) *StopApplyScheduledPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopApplyScheduledPlanResponseBody) SetHttpCode(v int32) *StopApplyScheduledPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *StopApplyScheduledPlanResponseBody) SetRequestId(v string) *StopApplyScheduledPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopApplyScheduledPlanResponseBody) SetSuccess(v bool) *StopApplyScheduledPlanResponseBody {
	s.Success = &v
	return s
}

type StopApplyScheduledPlanResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopApplyScheduledPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopApplyScheduledPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s StopApplyScheduledPlanResponse) GoString() string {
	return s.String()
}

func (s *StopApplyScheduledPlanResponse) SetHeaders(v map[string]*string) *StopApplyScheduledPlanResponse {
	s.Headers = v
	return s
}

func (s *StopApplyScheduledPlanResponse) SetStatusCode(v int32) *StopApplyScheduledPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *StopApplyScheduledPlanResponse) SetBody(v *StopApplyScheduledPlanResponseBody) *StopApplyScheduledPlanResponse {
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
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
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

func (s *StopJobResponseBody) SetAccessDeniedDetail(v string) *StopJobResponseBody {
	s.AccessDeniedDetail = &v
	return s
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

type StopSessionClusterHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s StopSessionClusterHeaders) String() string {
	return tea.Prettify(s)
}

func (s StopSessionClusterHeaders) GoString() string {
	return s.String()
}

func (s *StopSessionClusterHeaders) SetCommonHeaders(v map[string]*string) *StopSessionClusterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StopSessionClusterHeaders) SetWorkspace(v string) *StopSessionClusterHeaders {
	s.Workspace = &v
	return s
}

type StopSessionClusterResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StopSessionClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopSessionClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StopSessionClusterResponseBody) SetErrorCode(v string) *StopSessionClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StopSessionClusterResponseBody) SetErrorMessage(v string) *StopSessionClusterResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopSessionClusterResponseBody) SetHttpCode(v int32) *StopSessionClusterResponseBody {
	s.HttpCode = &v
	return s
}

func (s *StopSessionClusterResponseBody) SetRequestId(v string) *StopSessionClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopSessionClusterResponseBody) SetSuccess(v bool) *StopSessionClusterResponseBody {
	s.Success = &v
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

type UpdateDeploymentDraftHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateDeploymentDraftHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeploymentDraftHeaders) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentDraftHeaders) SetCommonHeaders(v map[string]*string) *UpdateDeploymentDraftHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateDeploymentDraftHeaders) SetWorkspace(v string) *UpdateDeploymentDraftHeaders {
	s.Workspace = &v
	return s
}

type UpdateDeploymentDraftRequest struct {
	// This parameter is required.
	Body *DeploymentDraft `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeploymentDraftRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeploymentDraftRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentDraftRequest) SetBody(v *DeploymentDraft) *UpdateDeploymentDraftRequest {
	s.Body = v
	return s
}

type UpdateDeploymentDraftResponseBody struct {
	Data *DeploymentDraft `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateDeploymentDraftResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeploymentDraftResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentDraftResponseBody) SetData(v *DeploymentDraft) *UpdateDeploymentDraftResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDeploymentDraftResponseBody) SetErrorCode(v string) *UpdateDeploymentDraftResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDeploymentDraftResponseBody) SetErrorMessage(v string) *UpdateDeploymentDraftResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDeploymentDraftResponseBody) SetHttpCode(v int32) *UpdateDeploymentDraftResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateDeploymentDraftResponseBody) SetRequestId(v string) *UpdateDeploymentDraftResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeploymentDraftResponseBody) SetSuccess(v bool) *UpdateDeploymentDraftResponseBody {
	s.Success = &v
	return s
}

type UpdateDeploymentDraftResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeploymentDraftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeploymentDraftResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeploymentDraftResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentDraftResponse) SetHeaders(v map[string]*string) *UpdateDeploymentDraftResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeploymentDraftResponse) SetStatusCode(v int32) *UpdateDeploymentDraftResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeploymentDraftResponse) SetBody(v *UpdateDeploymentDraftResponseBody) *UpdateDeploymentDraftResponse {
	s.Body = v
	return s
}

type UpdateDeploymentTargetHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateDeploymentTargetHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeploymentTargetHeaders) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentTargetHeaders) SetCommonHeaders(v map[string]*string) *UpdateDeploymentTargetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateDeploymentTargetHeaders) SetWorkspace(v string) *UpdateDeploymentTargetHeaders {
	s.Workspace = &v
	return s
}

type UpdateDeploymentTargetRequest struct {
	Body *ResourceSpec `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeploymentTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeploymentTargetRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentTargetRequest) SetBody(v *ResourceSpec) *UpdateDeploymentTargetRequest {
	s.Body = v
	return s
}

type UpdateDeploymentTargetResponseBody struct {
	Data *DeploymentTarget `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateDeploymentTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeploymentTargetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentTargetResponseBody) SetData(v *DeploymentTarget) *UpdateDeploymentTargetResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDeploymentTargetResponseBody) SetErrorCode(v string) *UpdateDeploymentTargetResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDeploymentTargetResponseBody) SetErrorMessage(v string) *UpdateDeploymentTargetResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDeploymentTargetResponseBody) SetHttpCode(v int32) *UpdateDeploymentTargetResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateDeploymentTargetResponseBody) SetRequestId(v string) *UpdateDeploymentTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeploymentTargetResponseBody) SetSuccess(v bool) *UpdateDeploymentTargetResponseBody {
	s.Success = &v
	return s
}

type UpdateDeploymentTargetResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeploymentTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeploymentTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeploymentTargetResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentTargetResponse) SetHeaders(v map[string]*string) *UpdateDeploymentTargetResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeploymentTargetResponse) SetStatusCode(v int32) *UpdateDeploymentTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeploymentTargetResponse) SetBody(v *UpdateDeploymentTargetResponseBody) *UpdateDeploymentTargetResponse {
	s.Body = v
	return s
}

type UpdateFolderHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f89a0c1ca8****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateFolderHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateFolderHeaders) GoString() string {
	return s.String()
}

func (s *UpdateFolderHeaders) SetCommonHeaders(v map[string]*string) *UpdateFolderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateFolderHeaders) SetWorkspace(v string) *UpdateFolderHeaders {
	s.Workspace = &v
	return s
}

type UpdateFolderRequest struct {
	// This parameter is required.
	Body *Folder `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFolderRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFolderRequest) GoString() string {
	return s.String()
}

func (s *UpdateFolderRequest) SetBody(v *Folder) *UpdateFolderRequest {
	s.Body = v
	return s
}

type UpdateFolderResponseBody struct {
	Data *Folder `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFolderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFolderResponseBody) SetData(v *Folder) *UpdateFolderResponseBody {
	s.Data = v
	return s
}

func (s *UpdateFolderResponseBody) SetErrorCode(v string) *UpdateFolderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateFolderResponseBody) SetErrorMessage(v string) *UpdateFolderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateFolderResponseBody) SetHttpCode(v int32) *UpdateFolderResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateFolderResponseBody) SetRequestId(v string) *UpdateFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFolderResponseBody) SetSuccess(v bool) *UpdateFolderResponseBody {
	s.Success = &v
	return s
}

type UpdateFolderResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFolderResponse) GoString() string {
	return s.String()
}

func (s *UpdateFolderResponse) SetHeaders(v map[string]*string) *UpdateFolderResponse {
	s.Headers = v
	return s
}

func (s *UpdateFolderResponse) SetStatusCode(v int32) *UpdateFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFolderResponse) SetBody(v *UpdateFolderResponseBody) *UpdateFolderResponse {
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

type UpdateScheduledPlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateScheduledPlanHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduledPlanHeaders) GoString() string {
	return s.String()
}

func (s *UpdateScheduledPlanHeaders) SetCommonHeaders(v map[string]*string) *UpdateScheduledPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateScheduledPlanHeaders) SetWorkspace(v string) *UpdateScheduledPlanHeaders {
	s.Workspace = &v
	return s
}

type UpdateScheduledPlanRequest struct {
	Body *ScheduledPlan `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScheduledPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduledPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduledPlanRequest) SetBody(v *ScheduledPlan) *UpdateScheduledPlanRequest {
	s.Body = v
	return s
}

type UpdateScheduledPlanResponseBody struct {
	Data *ScheduledPlan `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateScheduledPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduledPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScheduledPlanResponseBody) SetData(v *ScheduledPlan) *UpdateScheduledPlanResponseBody {
	s.Data = v
	return s
}

func (s *UpdateScheduledPlanResponseBody) SetErrorCode(v string) *UpdateScheduledPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateScheduledPlanResponseBody) SetErrorMessage(v string) *UpdateScheduledPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateScheduledPlanResponseBody) SetHttpCode(v int32) *UpdateScheduledPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateScheduledPlanResponseBody) SetRequestId(v string) *UpdateScheduledPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateScheduledPlanResponseBody) SetSuccess(v bool) *UpdateScheduledPlanResponseBody {
	s.Success = &v
	return s
}

type UpdateScheduledPlanResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScheduledPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScheduledPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduledPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateScheduledPlanResponse) SetHeaders(v map[string]*string) *UpdateScheduledPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateScheduledPlanResponse) SetStatusCode(v int32) *UpdateScheduledPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScheduledPlanResponse) SetBody(v *UpdateScheduledPlanResponseBody) *UpdateScheduledPlanResponse {
	s.Body = v
	return s
}

type UpdateSessionClusterHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 710d6a64d8c34d
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateSessionClusterHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateSessionClusterHeaders) GoString() string {
	return s.String()
}

func (s *UpdateSessionClusterHeaders) SetCommonHeaders(v map[string]*string) *UpdateSessionClusterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateSessionClusterHeaders) SetWorkspace(v string) *UpdateSessionClusterHeaders {
	s.Workspace = &v
	return s
}

type UpdateSessionClusterRequest struct {
	Body *SessionCluster `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSessionClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSessionClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateSessionClusterRequest) SetBody(v *SessionCluster) *UpdateSessionClusterRequest {
	s.Body = v
	return s
}

type UpdateSessionClusterResponseBody struct {
	Data *SessionCluster `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 1EF03B0C-F44F-47AD-BB48-D002D0F7B8C9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateSessionClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSessionClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSessionClusterResponseBody) SetData(v *SessionCluster) *UpdateSessionClusterResponseBody {
	s.Data = v
	return s
}

func (s *UpdateSessionClusterResponseBody) SetErrorCode(v string) *UpdateSessionClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateSessionClusterResponseBody) SetErrorMessage(v string) *UpdateSessionClusterResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateSessionClusterResponseBody) SetHttpCode(v int32) *UpdateSessionClusterResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateSessionClusterResponseBody) SetRequestId(v string) *UpdateSessionClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSessionClusterResponseBody) SetSuccess(v bool) *UpdateSessionClusterResponseBody {
	s.Success = &v
	return s
}

type UpdateSessionClusterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSessionClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSessionClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSessionClusterResponse) GoString() string {
	return s.String()
}

func (s *UpdateSessionClusterResponse) SetHeaders(v map[string]*string) *UpdateSessionClusterResponse {
	s.Headers = v
	return s
}

func (s *UpdateSessionClusterResponse) SetStatusCode(v int32) *UpdateSessionClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSessionClusterResponse) SetBody(v *UpdateSessionClusterResponseBody) *UpdateSessionClusterResponse {
	s.Body = v
	return s
}

type UpdateUdfArtifactHeaders struct {
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

func (s UpdateUdfArtifactHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateUdfArtifactHeaders) GoString() string {
	return s.String()
}

func (s *UpdateUdfArtifactHeaders) SetCommonHeaders(v map[string]*string) *UpdateUdfArtifactHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateUdfArtifactHeaders) SetWorkspace(v string) *UpdateUdfArtifactHeaders {
	s.Workspace = &v
	return s
}

type UpdateUdfArtifactRequest struct {
	// The details of the JAR file of the UDF.
	//
	// This parameter is required.
	Body *UdfArtifact `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUdfArtifactRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUdfArtifactRequest) GoString() string {
	return s.String()
}

func (s *UpdateUdfArtifactRequest) SetBody(v *UdfArtifact) *UpdateUdfArtifactRequest {
	s.Body = v
	return s
}

type UpdateUdfArtifactResponseBody struct {
	// The result of updating the JAR file of the UDF.
	Data *UpdateUdfArtifactResult `json:"data,omitempty" xml:"data,omitempty"`
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
	// CBC799F0-ABCD-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateUdfArtifactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUdfArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUdfArtifactResponseBody) SetData(v *UpdateUdfArtifactResult) *UpdateUdfArtifactResponseBody {
	s.Data = v
	return s
}

func (s *UpdateUdfArtifactResponseBody) SetErrorCode(v string) *UpdateUdfArtifactResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateUdfArtifactResponseBody) SetErrorMessage(v string) *UpdateUdfArtifactResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateUdfArtifactResponseBody) SetHttpCode(v int32) *UpdateUdfArtifactResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateUdfArtifactResponseBody) SetRequestId(v string) *UpdateUdfArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUdfArtifactResponseBody) SetSuccess(v bool) *UpdateUdfArtifactResponseBody {
	s.Success = &v
	return s
}

type UpdateUdfArtifactResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUdfArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUdfArtifactResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUdfArtifactResponse) GoString() string {
	return s.String()
}

func (s *UpdateUdfArtifactResponse) SetHeaders(v map[string]*string) *UpdateUdfArtifactResponse {
	s.Headers = v
	return s
}

func (s *UpdateUdfArtifactResponse) SetStatusCode(v int32) *UpdateUdfArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUdfArtifactResponse) SetBody(v *UpdateUdfArtifactResponseBody) *UpdateUdfArtifactResponse {
	s.Body = v
	return s
}

type UpdateVariableHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateVariableHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateVariableHeaders) GoString() string {
	return s.String()
}

func (s *UpdateVariableHeaders) SetCommonHeaders(v map[string]*string) *UpdateVariableHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateVariableHeaders) SetWorkspace(v string) *UpdateVariableHeaders {
	s.Workspace = &v
	return s
}

type UpdateVariableRequest struct {
	// This parameter is required.
	Body *Variable `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVariableRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVariableRequest) GoString() string {
	return s.String()
}

func (s *UpdateVariableRequest) SetBody(v *Variable) *UpdateVariableRequest {
	s.Body = v
	return s
}

type UpdateVariableResponseBody struct {
	// example:
	//
	// “”
	AccessDeniedDetail *string   `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	Data               *Variable `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 1EF03B0C-F44F-47AD-BB48-D002D0F7B8C9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateVariableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVariableResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVariableResponseBody) SetAccessDeniedDetail(v string) *UpdateVariableResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateVariableResponseBody) SetData(v *Variable) *UpdateVariableResponseBody {
	s.Data = v
	return s
}

func (s *UpdateVariableResponseBody) SetErrorCode(v string) *UpdateVariableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateVariableResponseBody) SetErrorMessage(v string) *UpdateVariableResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateVariableResponseBody) SetHttpCode(v int32) *UpdateVariableResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateVariableResponseBody) SetRequestId(v string) *UpdateVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVariableResponseBody) SetSuccess(v bool) *UpdateVariableResponseBody {
	s.Success = &v
	return s
}

type UpdateVariableResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVariableResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVariableResponse) GoString() string {
	return s.String()
}

func (s *UpdateVariableResponse) SetHeaders(v map[string]*string) *UpdateVariableResponse {
	s.Headers = v
	return s
}

func (s *UpdateVariableResponse) SetStatusCode(v int32) *UpdateVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVariableResponse) SetBody(v *UpdateVariableResponseBody) *UpdateVariableResponse {
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
// 执行定时计划
//
// @param headers - ApplyScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyScheduledPlanResponse
func (client *Client) ApplyScheduledPlanWithOptions(namespace *string, scheduledPlanId *string, headers *ApplyScheduledPlanHeaders, runtime *util.RuntimeOptions) (_result *ApplyScheduledPlanResponse, _err error) {
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
		Action:      tea.String("ApplyScheduledPlan"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/scheduled-plans/" + tea.StringValue(openapiutil.GetEncodeParam(scheduledPlanId)) + "%3Aapply"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ApplyScheduledPlanResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ApplyScheduledPlanResponse{}
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
// 执行定时计划
//
// @return ApplyScheduledPlanResponse
func (client *Client) ApplyScheduledPlan(namespace *string, scheduledPlanId *string) (_result *ApplyScheduledPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ApplyScheduledPlanHeaders{}
	_result = &ApplyScheduledPlanResponse{}
	_body, _err := client.ApplyScheduledPlanWithOptions(namespace, scheduledPlanId, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateDeploymentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateDeploymentResponse{}
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
// create a deploymentDraft
//
// @param request - CreateDeploymentDraftRequest
//
// @param headers - CreateDeploymentDraftHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeploymentDraftResponse
func (client *Client) CreateDeploymentDraftWithOptions(namespace *string, request *CreateDeploymentDraftRequest, headers *CreateDeploymentDraftHeaders, runtime *util.RuntimeOptions) (_result *CreateDeploymentDraftResponse, _err error) {
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
		Action:      tea.String("CreateDeploymentDraft"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployment-drafts"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateDeploymentDraftResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateDeploymentDraftResponse{}
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
// create a deploymentDraft
//
// @param request - CreateDeploymentDraftRequest
//
// @return CreateDeploymentDraftResponse
func (client *Client) CreateDeploymentDraft(namespace *string, request *CreateDeploymentDraftRequest) (_result *CreateDeploymentDraftResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateDeploymentDraftHeaders{}
	_result = &CreateDeploymentDraftResponse{}
	_body, _err := client.CreateDeploymentDraftWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建deploymentTarget
//
// @param request - CreateDeploymentTargetRequest
//
// @param headers - CreateDeploymentTargetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeploymentTargetResponse
func (client *Client) CreateDeploymentTargetWithOptions(namespace *string, request *CreateDeploymentTargetRequest, headers *CreateDeploymentTargetHeaders, runtime *util.RuntimeOptions) (_result *CreateDeploymentTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeploymentTargetName)) {
		query["deploymentTargetName"] = request.DeploymentTargetName
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
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDeploymentTarget"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployment-targets"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateDeploymentTargetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateDeploymentTargetResponse{}
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
// 创建deploymentTarget
//
// @param request - CreateDeploymentTargetRequest
//
// @return CreateDeploymentTargetResponse
func (client *Client) CreateDeploymentTarget(namespace *string, request *CreateDeploymentTargetRequest) (_result *CreateDeploymentTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateDeploymentTargetHeaders{}
	_result = &CreateDeploymentTargetResponse{}
	_body, _err := client.CreateDeploymentTargetWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// create a folder
//
// @param request - CreateFolderRequest
//
// @param headers - CreateFolderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFolderResponse
func (client *Client) CreateFolderWithOptions(namespace *string, request *CreateFolderRequest, headers *CreateFolderHeaders, runtime *util.RuntimeOptions) (_result *CreateFolderResponse, _err error) {
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
		Action:      tea.String("CreateFolder"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/folder"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateFolderResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateFolderResponse{}
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
// create a folder
//
// @param request - CreateFolderRequest
//
// @return CreateFolderResponse
func (client *Client) CreateFolder(namespace *string, request *CreateFolderRequest) (_result *CreateFolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateFolderHeaders{}
	_result = &CreateFolderResponse{}
	_body, _err := client.CreateFolderWithOptions(namespace, request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateMemberResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateMemberResponse{}
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateSavepointResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateSavepointResponse{}
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
// 创建定时执行计划
//
// @param request - CreateScheduledPlanRequest
//
// @param headers - CreateScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScheduledPlanResponse
func (client *Client) CreateScheduledPlanWithOptions(namespace *string, request *CreateScheduledPlanRequest, headers *CreateScheduledPlanHeaders, runtime *util.RuntimeOptions) (_result *CreateScheduledPlanResponse, _err error) {
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
		Action:      tea.String("CreateScheduledPlan"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/scheduled-plans"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateScheduledPlanResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateScheduledPlanResponse{}
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
// 创建定时执行计划
//
// @param request - CreateScheduledPlanRequest
//
// @return CreateScheduledPlanResponse
func (client *Client) CreateScheduledPlan(namespace *string, request *CreateScheduledPlanRequest) (_result *CreateScheduledPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateScheduledPlanHeaders{}
	_result = &CreateScheduledPlanResponse{}
	_body, _err := client.CreateScheduledPlanWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建session集群
//
// @param request - CreateSessionClusterRequest
//
// @param headers - CreateSessionClusterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSessionClusterResponse
func (client *Client) CreateSessionClusterWithOptions(namespace *string, request *CreateSessionClusterRequest, headers *CreateSessionClusterHeaders, runtime *util.RuntimeOptions) (_result *CreateSessionClusterResponse, _err error) {
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
		Action:      tea.String("CreateSessionCluster"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/sessionclusters"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateSessionClusterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateSessionClusterResponse{}
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
// 创建session集群
//
// @param request - CreateSessionClusterRequest
//
// @return CreateSessionClusterResponse
func (client *Client) CreateSessionCluster(namespace *string, request *CreateSessionClusterRequest) (_result *CreateSessionClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateSessionClusterHeaders{}
	_result = &CreateSessionClusterResponse{}
	_body, _err := client.CreateSessionClusterWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Parses all user-defined function (UDF) methods in your JAR or Python file and creates an artifact configuration for a UDF.
//
// @param request - CreateUdfArtifactRequest
//
// @param headers - CreateUdfArtifactHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUdfArtifactResponse
func (client *Client) CreateUdfArtifactWithOptions(namespace *string, request *CreateUdfArtifactRequest, headers *CreateUdfArtifactHeaders, runtime *util.RuntimeOptions) (_result *CreateUdfArtifactResponse, _err error) {
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
		Action:      tea.String("CreateUdfArtifact"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/udfartifacts"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateUdfArtifactResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateUdfArtifactResponse{}
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
// Parses all user-defined function (UDF) methods in your JAR or Python file and creates an artifact configuration for a UDF.
//
// @param request - CreateUdfArtifactRequest
//
// @return CreateUdfArtifactResponse
func (client *Client) CreateUdfArtifact(namespace *string, request *CreateUdfArtifactRequest) (_result *CreateUdfArtifactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateUdfArtifactHeaders{}
	_result = &CreateUdfArtifactResponse{}
	_body, _err := client.CreateUdfArtifactWithOptions(namespace, request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateVariableResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateVariableResponse{}
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
// Deletes a registered custom connector from a workspace.
//
// @param headers - DeleteCustomConnectorHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomConnectorResponse
func (client *Client) DeleteCustomConnectorWithOptions(namespace *string, connectorName *string, headers *DeleteCustomConnectorHeaders, runtime *util.RuntimeOptions) (_result *DeleteCustomConnectorResponse, _err error) {
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
		Action:      tea.String("DeleteCustomConnector"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/connectors/" + tea.StringValue(openapiutil.GetEncodeParam(connectorName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteCustomConnectorResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteCustomConnectorResponse{}
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
// Deletes a registered custom connector from a workspace.
//
// @return DeleteCustomConnectorResponse
func (client *Client) DeleteCustomConnector(namespace *string, connectorName *string) (_result *DeleteCustomConnectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteCustomConnectorHeaders{}
	_result = &DeleteCustomConnectorResponse{}
	_body, _err := client.DeleteCustomConnectorWithOptions(namespace, connectorName, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteDeploymentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteDeploymentResponse{}
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
// delete a deploymentDraft
//
// @param headers - DeleteDeploymentDraftHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeploymentDraftResponse
func (client *Client) DeleteDeploymentDraftWithOptions(namespace *string, deploymentDraftId *string, headers *DeleteDeploymentDraftHeaders, runtime *util.RuntimeOptions) (_result *DeleteDeploymentDraftResponse, _err error) {
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
		Action:      tea.String("DeleteDeploymentDraft"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployment-drafts/" + tea.StringValue(openapiutil.GetEncodeParam(deploymentDraftId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteDeploymentDraftResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteDeploymentDraftResponse{}
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
// delete a deploymentDraft
//
// @return DeleteDeploymentDraftResponse
func (client *Client) DeleteDeploymentDraft(namespace *string, deploymentDraftId *string) (_result *DeleteDeploymentDraftResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteDeploymentDraftHeaders{}
	_result = &DeleteDeploymentDraftResponse{}
	_body, _err := client.DeleteDeploymentDraftWithOptions(namespace, deploymentDraftId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除deploymentTarget
//
// @param headers - DeleteDeploymentTargetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeploymentTargetResponse
func (client *Client) DeleteDeploymentTargetWithOptions(namespace *string, deploymentTargetName *string, headers *DeleteDeploymentTargetHeaders, runtime *util.RuntimeOptions) (_result *DeleteDeploymentTargetResponse, _err error) {
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
		Action:      tea.String("DeleteDeploymentTarget"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployment-targets/" + tea.StringValue(openapiutil.GetEncodeParam(deploymentTargetName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteDeploymentTargetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteDeploymentTargetResponse{}
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
// 删除deploymentTarget
//
// @return DeleteDeploymentTargetResponse
func (client *Client) DeleteDeploymentTarget(namespace *string, deploymentTargetName *string) (_result *DeleteDeploymentTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteDeploymentTargetHeaders{}
	_result = &DeleteDeploymentTargetResponse{}
	_body, _err := client.DeleteDeploymentTargetWithOptions(namespace, deploymentTargetName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// delete a folder
//
// @param headers - DeleteFolderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFolderResponse
func (client *Client) DeleteFolderWithOptions(namespace *string, folderId *string, headers *DeleteFolderHeaders, runtime *util.RuntimeOptions) (_result *DeleteFolderResponse, _err error) {
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
		Action:      tea.String("DeleteFolder"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/folder/" + tea.StringValue(openapiutil.GetEncodeParam(folderId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteFolderResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteFolderResponse{}
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
// delete a folder
//
// @return DeleteFolderResponse
func (client *Client) DeleteFolder(namespace *string, folderId *string) (_result *DeleteFolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteFolderHeaders{}
	_result = &DeleteFolderResponse{}
	_body, _err := client.DeleteFolderWithOptions(namespace, folderId, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteJobResponse{}
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteMemberResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteMemberResponse{}
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteSavepointResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteSavepointResponse{}
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
// 删除定时执行计划
//
// @param headers - DeleteScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScheduledPlanResponse
func (client *Client) DeleteScheduledPlanWithOptions(namespace *string, scheduledPlanId *string, headers *DeleteScheduledPlanHeaders, runtime *util.RuntimeOptions) (_result *DeleteScheduledPlanResponse, _err error) {
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
		Action:      tea.String("DeleteScheduledPlan"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/scheduled-plans/" + tea.StringValue(openapiutil.GetEncodeParam(scheduledPlanId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteScheduledPlanResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteScheduledPlanResponse{}
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
// 删除定时执行计划
//
// @return DeleteScheduledPlanResponse
func (client *Client) DeleteScheduledPlan(namespace *string, scheduledPlanId *string) (_result *DeleteScheduledPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteScheduledPlanHeaders{}
	_result = &DeleteScheduledPlanResponse{}
	_body, _err := client.DeleteScheduledPlanWithOptions(namespace, scheduledPlanId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除session集群
//
// @param headers - DeleteSessionClusterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSessionClusterResponse
func (client *Client) DeleteSessionClusterWithOptions(namespace *string, sessionClusterName *string, headers *DeleteSessionClusterHeaders, runtime *util.RuntimeOptions) (_result *DeleteSessionClusterResponse, _err error) {
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
		Action:      tea.String("DeleteSessionCluster"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/sessionclusters/" + tea.StringValue(openapiutil.GetEncodeParam(sessionClusterName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteSessionClusterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteSessionClusterResponse{}
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
// 删除session集群
//
// @return DeleteSessionClusterResponse
func (client *Client) DeleteSessionCluster(namespace *string, sessionClusterName *string) (_result *DeleteSessionClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteSessionClusterHeaders{}
	_result = &DeleteSessionClusterResponse{}
	_body, _err := client.DeleteSessionClusterWithOptions(namespace, sessionClusterName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除UdfArtifact
//
// @param headers - DeleteUdfArtifactHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUdfArtifactResponse
func (client *Client) DeleteUdfArtifactWithOptions(namespace *string, udfArtifactName *string, headers *DeleteUdfArtifactHeaders, runtime *util.RuntimeOptions) (_result *DeleteUdfArtifactResponse, _err error) {
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
		Action:      tea.String("DeleteUdfArtifact"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/udfartifacts/" + tea.StringValue(openapiutil.GetEncodeParam(udfArtifactName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteUdfArtifactResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteUdfArtifactResponse{}
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
// 删除UdfArtifact
//
// @return DeleteUdfArtifactResponse
func (client *Client) DeleteUdfArtifact(namespace *string, udfArtifactName *string) (_result *DeleteUdfArtifactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteUdfArtifactHeaders{}
	_result = &DeleteUdfArtifactResponse{}
	_body, _err := client.DeleteUdfArtifactWithOptions(namespace, udfArtifactName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an existing user-defined function (UDF) from a Realtime Compute for Apache Flink workspace.
//
// @param request - DeleteUdfFunctionRequest
//
// @param headers - DeleteUdfFunctionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUdfFunctionResponse
func (client *Client) DeleteUdfFunctionWithOptions(namespace *string, functionName *string, request *DeleteUdfFunctionRequest, headers *DeleteUdfFunctionHeaders, runtime *util.RuntimeOptions) (_result *DeleteUdfFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassName)) {
		query["className"] = request.ClassName
	}

	if !tea.BoolValue(util.IsUnset(request.UdfArtifactName)) {
		query["udfArtifactName"] = request.UdfArtifactName
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
		Action:      tea.String("DeleteUdfFunction"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/udfartifacts/function/" + tea.StringValue(openapiutil.GetEncodeParam(functionName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteUdfFunctionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteUdfFunctionResponse{}
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
// Deletes an existing user-defined function (UDF) from a Realtime Compute for Apache Flink workspace.
//
// @param request - DeleteUdfFunctionRequest
//
// @return DeleteUdfFunctionResponse
func (client *Client) DeleteUdfFunction(namespace *string, functionName *string, request *DeleteUdfFunctionRequest) (_result *DeleteUdfFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteUdfFunctionHeaders{}
	_result = &DeleteUdfFunctionResponse{}
	_body, _err := client.DeleteUdfFunctionWithOptions(namespace, functionName, request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteVariableResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteVariableResponse{}
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
// deploy deploymentDraft async
//
// @param request - DeployDeploymentDraftAsyncRequest
//
// @param headers - DeployDeploymentDraftAsyncHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployDeploymentDraftAsyncResponse
func (client *Client) DeployDeploymentDraftAsyncWithOptions(namespace *string, request *DeployDeploymentDraftAsyncRequest, headers *DeployDeploymentDraftAsyncHeaders, runtime *util.RuntimeOptions) (_result *DeployDeploymentDraftAsyncResponse, _err error) {
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
		Action:      tea.String("DeployDeploymentDraftAsync"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployment-drafts/async-deploy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeployDeploymentDraftAsyncResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeployDeploymentDraftAsyncResponse{}
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
// deploy deploymentDraft async
//
// @param request - DeployDeploymentDraftAsyncRequest
//
// @return DeployDeploymentDraftAsyncResponse
func (client *Client) DeployDeploymentDraftAsync(namespace *string, request *DeployDeploymentDraftAsyncRequest) (_result *DeployDeploymentDraftAsyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeployDeploymentDraftAsyncHeaders{}
	_result = &DeployDeploymentDraftAsyncResponse{}
	_body, _err := client.DeployDeploymentDraftAsyncWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 执行sql语句
//
// @param request - ExecuteSqlStatementRequest
//
// @param headers - ExecuteSqlStatementHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteSqlStatementResponse
func (client *Client) ExecuteSqlStatementWithOptions(namespace *string, request *ExecuteSqlStatementRequest, headers *ExecuteSqlStatementHeaders, runtime *util.RuntimeOptions) (_result *ExecuteSqlStatementResponse, _err error) {
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
		Action:      tea.String("ExecuteSqlStatement"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/sql-statement/execute"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ExecuteSqlStatementResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ExecuteSqlStatementResponse{}
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
// 执行sql语句
//
// @param request - ExecuteSqlStatementRequest
//
// @return ExecuteSqlStatementResponse
func (client *Client) ExecuteSqlStatement(namespace *string, request *ExecuteSqlStatementRequest) (_result *ExecuteSqlStatementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExecuteSqlStatementHeaders{}
	_result = &ExecuteSqlStatementResponse{}
	_body, _err := client.ExecuteSqlStatementWithOptions(namespace, request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &FlinkApiProxyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &FlinkApiProxyResponse{}
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GenerateResourcePlanWithFlinkConfAsyncResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GenerateResourcePlanWithFlinkConfAsyncResponse{}
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
// 获取应用中的执行定时计划
//
// @param request - GetAppliedScheduledPlanRequest
//
// @param headers - GetAppliedScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppliedScheduledPlanResponse
func (client *Client) GetAppliedScheduledPlanWithOptions(namespace *string, request *GetAppliedScheduledPlanRequest, headers *GetAppliedScheduledPlanHeaders, runtime *util.RuntimeOptions) (_result *GetAppliedScheduledPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeploymentId)) {
		query["deploymentId"] = request.DeploymentId
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
		Action:      tea.String("GetAppliedScheduledPlan"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/scheduled-plans%3AgetExecutedScheduledPlan"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetAppliedScheduledPlanResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetAppliedScheduledPlanResponse{}
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
// 获取应用中的执行定时计划
//
// @param request - GetAppliedScheduledPlanRequest
//
// @return GetAppliedScheduledPlanResponse
func (client *Client) GetAppliedScheduledPlan(namespace *string, request *GetAppliedScheduledPlanRequest) (_result *GetAppliedScheduledPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAppliedScheduledPlanHeaders{}
	_result = &GetAppliedScheduledPlanResponse{}
	_body, _err := client.GetAppliedScheduledPlanWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取catalog
//
// @param request - GetCatalogsRequest
//
// @param headers - GetCatalogsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCatalogsResponse
func (client *Client) GetCatalogsWithOptions(namespace *string, request *GetCatalogsRequest, headers *GetCatalogsHeaders, runtime *util.RuntimeOptions) (_result *GetCatalogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogName)) {
		query["catalogName"] = request.CatalogName
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
		Action:      tea.String("GetCatalogs"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/catalogs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetCatalogsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetCatalogsResponse{}
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
// 获取catalog
//
// @param request - GetCatalogsRequest
//
// @return GetCatalogsResponse
func (client *Client) GetCatalogs(namespace *string, request *GetCatalogsRequest) (_result *GetCatalogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCatalogsHeaders{}
	_result = &GetCatalogsResponse{}
	_body, _err := client.GetCatalogsWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取database
//
// @param request - GetDatabasesRequest
//
// @param headers - GetDatabasesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatabasesResponse
func (client *Client) GetDatabasesWithOptions(namespace *string, catalogName *string, request *GetDatabasesRequest, headers *GetDatabasesHeaders, runtime *util.RuntimeOptions) (_result *GetDatabasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["databaseName"] = request.DatabaseName
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
		Action:      tea.String("GetDatabases"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/catalogs/" + tea.StringValue(openapiutil.GetEncodeParam(catalogName)) + "/databases"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDatabasesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDatabasesResponse{}
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
// 获取database
//
// @param request - GetDatabasesRequest
//
// @return GetDatabasesResponse
func (client *Client) GetDatabases(namespace *string, catalogName *string, request *GetDatabasesRequest) (_result *GetDatabasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDatabasesHeaders{}
	_result = &GetDatabasesResponse{}
	_body, _err := client.GetDatabasesWithOptions(namespace, catalogName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// get deploy deploymentDraft result
//
// @param headers - GetDeployDeploymentDraftResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeployDeploymentDraftResultResponse
func (client *Client) GetDeployDeploymentDraftResultWithOptions(namespace *string, ticketId *string, headers *GetDeployDeploymentDraftResultHeaders, runtime *util.RuntimeOptions) (_result *GetDeployDeploymentDraftResultResponse, _err error) {
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
		Action:      tea.String("GetDeployDeploymentDraftResult"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployment-drafts/tickets/" + tea.StringValue(openapiutil.GetEncodeParam(ticketId)) + "/async-deploy"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDeployDeploymentDraftResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDeployDeploymentDraftResultResponse{}
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
// get deploy deploymentDraft result
//
// @return GetDeployDeploymentDraftResultResponse
func (client *Client) GetDeployDeploymentDraftResult(namespace *string, ticketId *string) (_result *GetDeployDeploymentDraftResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDeployDeploymentDraftResultHeaders{}
	_result = &GetDeployDeploymentDraftResultResponse{}
	_body, _err := client.GetDeployDeploymentDraftResultWithOptions(namespace, ticketId, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDeploymentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDeploymentResponse{}
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
// get a deploymentDraft
//
// @param headers - GetDeploymentDraftHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeploymentDraftResponse
func (client *Client) GetDeploymentDraftWithOptions(namespace *string, deploymentDraftId *string, headers *GetDeploymentDraftHeaders, runtime *util.RuntimeOptions) (_result *GetDeploymentDraftResponse, _err error) {
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
		Action:      tea.String("GetDeploymentDraft"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployment-drafts/" + tea.StringValue(openapiutil.GetEncodeParam(deploymentDraftId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDeploymentDraftResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDeploymentDraftResponse{}
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
// get a deploymentDraft
//
// @return GetDeploymentDraftResponse
func (client *Client) GetDeploymentDraft(namespace *string, deploymentDraftId *string) (_result *GetDeploymentDraftResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDeploymentDraftHeaders{}
	_result = &GetDeploymentDraftResponse{}
	_body, _err := client.GetDeploymentDraftWithOptions(namespace, deploymentDraftId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// get deploymentDraft lock
//
// @param request - GetDeploymentDraftLockRequest
//
// @param headers - GetDeploymentDraftLockHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeploymentDraftLockResponse
func (client *Client) GetDeploymentDraftLockWithOptions(namespace *string, request *GetDeploymentDraftLockRequest, headers *GetDeploymentDraftLockHeaders, runtime *util.RuntimeOptions) (_result *GetDeploymentDraftLockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeploymentDraftId)) {
		query["deploymentDraftId"] = request.DeploymentDraftId
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
		Action:      tea.String("GetDeploymentDraftLock"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployment-drafts/getLock"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDeploymentDraftLockResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDeploymentDraftLockResponse{}
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
// get deploymentDraft lock
//
// @param request - GetDeploymentDraftLockRequest
//
// @return GetDeploymentDraftLockResponse
func (client *Client) GetDeploymentDraftLock(namespace *string, request *GetDeploymentDraftLockRequest) (_result *GetDeploymentDraftLockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDeploymentDraftLockHeaders{}
	_result = &GetDeploymentDraftLockResponse{}
	_body, _err := client.GetDeploymentDraftLockWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取运行事件
//
// @param request - GetEventsRequest
//
// @param headers - GetEventsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEventsResponse
func (client *Client) GetEventsWithOptions(namespace *string, request *GetEventsRequest, headers *GetEventsHeaders, runtime *util.RuntimeOptions) (_result *GetEventsResponse, _err error) {
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
		Action:      tea.String("GetEvents"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/events"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetEventsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetEventsResponse{}
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
// 获取运行事件
//
// @param request - GetEventsRequest
//
// @return GetEventsResponse
func (client *Client) GetEvents(namespace *string, request *GetEventsRequest) (_result *GetEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetEventsHeaders{}
	_result = &GetEventsResponse{}
	_body, _err := client.GetEventsWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// get a folder
//
// @param request - GetFolderRequest
//
// @param headers - GetFolderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFolderResponse
func (client *Client) GetFolderWithOptions(namespace *string, request *GetFolderRequest, headers *GetFolderHeaders, runtime *util.RuntimeOptions) (_result *GetFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		query["folderId"] = request.FolderId
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
		Action:      tea.String("GetFolder"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/folder"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetFolderResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetFolderResponse{}
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
// get a folder
//
// @param request - GetFolderRequest
//
// @return GetFolderResponse
func (client *Client) GetFolder(namespace *string, request *GetFolderRequest) (_result *GetFolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFolderHeaders{}
	_result = &GetFolderResponse{}
	_body, _err := client.GetFolderWithOptions(namespace, request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetGenerateResourcePlanResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetGenerateResourcePlanResultResponse{}
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
// 查询动态更新结果
//
// @param headers - GetHotUpdateJobResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotUpdateJobResultResponse
func (client *Client) GetHotUpdateJobResultWithOptions(namespace *string, jobHotUpdateId *string, headers *GetHotUpdateJobResultHeaders, runtime *util.RuntimeOptions) (_result *GetHotUpdateJobResultResponse, _err error) {
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
		Action:      tea.String("GetHotUpdateJobResult"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/jobs/hot-updates/" + tea.StringValue(openapiutil.GetEncodeParam(jobHotUpdateId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetHotUpdateJobResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetHotUpdateJobResultResponse{}
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
// 查询动态更新结果
//
// @return GetHotUpdateJobResultResponse
func (client *Client) GetHotUpdateJobResult(namespace *string, jobHotUpdateId *string) (_result *GetHotUpdateJobResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHotUpdateJobResultHeaders{}
	_result = &GetHotUpdateJobResultResponse{}
	_body, _err := client.GetHotUpdateJobResultWithOptions(namespace, jobHotUpdateId, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetJobResponse{}
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
// 获取作业诊断信息
//
// @param headers - GetJobDiagnosisHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobDiagnosisResponse
func (client *Client) GetJobDiagnosisWithOptions(namespace *string, deploymentId *string, jobId *string, headers *GetJobDiagnosisHeaders, runtime *util.RuntimeOptions) (_result *GetJobDiagnosisResponse, _err error) {
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
		Action:      tea.String("GetJobDiagnosis"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployments/" + tea.StringValue(openapiutil.GetEncodeParam(deploymentId)) + "/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(jobId)) + "/job-diagnoses/lite"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetJobDiagnosisResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetJobDiagnosisResponse{}
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
// 获取作业诊断信息
//
// @return GetJobDiagnosisResponse
func (client *Client) GetJobDiagnosis(namespace *string, deploymentId *string, jobId *string) (_result *GetJobDiagnosisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetJobDiagnosisHeaders{}
	_result = &GetJobDiagnosisResponse{}
	_body, _err := client.GetJobDiagnosisWithOptions(namespace, deploymentId, jobId, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLatestJobStartLogResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLatestJobStartLogResponse{}
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
// Obtains the lineage information of a deployment.
//
// @param request - GetLineageInfoRequest
//
// @param headers - GetLineageInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLineageInfoResponse
func (client *Client) GetLineageInfoWithOptions(request *GetLineageInfoRequest, headers *GetLineageInfoHeaders, runtime *util.RuntimeOptions) (_result *GetLineageInfoResponse, _err error) {
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
		Action:      tea.String("GetLineageInfo"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/meta/v2/lineage"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLineageInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLineageInfoResponse{}
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
// Obtains the lineage information of a deployment.
//
// @param request - GetLineageInfoRequest
//
// @return GetLineageInfoResponse
func (client *Client) GetLineageInfo(request *GetLineageInfoRequest) (_result *GetLineageInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetLineageInfoHeaders{}
	_result = &GetLineageInfoResponse{}
	_body, _err := client.GetLineageInfoWithOptions(request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetMemberResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetMemberResponse{}
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetSavepointResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetSavepointResponse{}
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
// 获取session集群
//
// @param headers - GetSessionClusterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSessionClusterResponse
func (client *Client) GetSessionClusterWithOptions(namespace *string, sessionClusterName *string, headers *GetSessionClusterHeaders, runtime *util.RuntimeOptions) (_result *GetSessionClusterResponse, _err error) {
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
		Action:      tea.String("GetSessionCluster"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/sessionclusters/" + tea.StringValue(openapiutil.GetEncodeParam(sessionClusterName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetSessionClusterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetSessionClusterResponse{}
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
// 获取session集群
//
// @return GetSessionClusterResponse
func (client *Client) GetSessionCluster(namespace *string, sessionClusterName *string) (_result *GetSessionClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSessionClusterHeaders{}
	_result = &GetSessionClusterResponse{}
	_body, _err := client.GetSessionClusterWithOptions(namespace, sessionClusterName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取table
//
// @param request - GetTablesRequest
//
// @param headers - GetTablesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTablesResponse
func (client *Client) GetTablesWithOptions(namespace *string, catalogName *string, databaseName *string, request *GetTablesRequest, headers *GetTablesHeaders, runtime *util.RuntimeOptions) (_result *GetTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["tableName"] = request.TableName
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
		Action:      tea.String("GetTables"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/catalogs/" + tea.StringValue(openapiutil.GetEncodeParam(catalogName)) + "/databases/" + tea.StringValue(openapiutil.GetEncodeParam(databaseName)) + "/tables"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTablesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTablesResponse{}
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
// 获取table
//
// @param request - GetTablesRequest
//
// @return GetTablesResponse
func (client *Client) GetTables(namespace *string, catalogName *string, databaseName *string, request *GetTablesRequest) (_result *GetTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTablesHeaders{}
	_result = &GetTablesResponse{}
	_body, _err := client.GetTablesWithOptions(namespace, catalogName, databaseName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of the JAR or Python file that corresponds to the user-defined function (UDF) that you upload and create.
//
// @param request - GetUdfArtifactsRequest
//
// @param headers - GetUdfArtifactsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUdfArtifactsResponse
func (client *Client) GetUdfArtifactsWithOptions(namespace *string, request *GetUdfArtifactsRequest, headers *GetUdfArtifactsHeaders, runtime *util.RuntimeOptions) (_result *GetUdfArtifactsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UdfArtifactName)) {
		query["udfArtifactName"] = request.UdfArtifactName
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
		Action:      tea.String("GetUdfArtifacts"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/udfartifacts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetUdfArtifactsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetUdfArtifactsResponse{}
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
// Obtains the details of the JAR or Python file that corresponds to the user-defined function (UDF) that you upload and create.
//
// @param request - GetUdfArtifactsRequest
//
// @return GetUdfArtifactsResponse
func (client *Client) GetUdfArtifacts(namespace *string, request *GetUdfArtifactsRequest) (_result *GetUdfArtifactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUdfArtifactsHeaders{}
	_result = &GetUdfArtifactsResponse{}
	_body, _err := client.GetUdfArtifactsWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Dynamically updates parameters or resources of a deployment that is running.
//
// @param headers - HotUpdateJobHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotUpdateJobResponse
func (client *Client) HotUpdateJobWithOptions(namespace *string, jobId *string, headers *HotUpdateJobHeaders, runtime *util.RuntimeOptions) (_result *HotUpdateJobResponse, _err error) {
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
		Action:      tea.String("HotUpdateJob"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(jobId)) + "%3AhotUpdate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &HotUpdateJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &HotUpdateJobResponse{}
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
// Dynamically updates parameters or resources of a deployment that is running.
//
// @return HotUpdateJobResponse
func (client *Client) HotUpdateJob(namespace *string, jobId *string) (_result *HotUpdateJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HotUpdateJobHeaders{}
	_result = &HotUpdateJobResponse{}
	_body, _err := client.HotUpdateJobWithOptions(namespace, jobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a list of existing custom connectors.
//
// @param headers - ListCustomConnectorsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomConnectorsResponse
func (client *Client) ListCustomConnectorsWithOptions(namespace *string, headers *ListCustomConnectorsHeaders, runtime *util.RuntimeOptions) (_result *ListCustomConnectorsResponse, _err error) {
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
		Action:      tea.String("ListCustomConnectors"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/connectors"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListCustomConnectorsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListCustomConnectorsResponse{}
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
// Obtains a list of existing custom connectors.
//
// @return ListCustomConnectorsResponse
func (client *Client) ListCustomConnectors(namespace *string) (_result *ListCustomConnectorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListCustomConnectorsHeaders{}
	_result = &ListCustomConnectorsResponse{}
	_body, _err := client.ListCustomConnectorsWithOptions(namespace, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// list deploymentDrafts
//
// @param request - ListDeploymentDraftsRequest
//
// @param headers - ListDeploymentDraftsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentDraftsResponse
func (client *Client) ListDeploymentDraftsWithOptions(namespace *string, request *ListDeploymentDraftsRequest, headers *ListDeploymentDraftsHeaders, runtime *util.RuntimeOptions) (_result *ListDeploymentDraftsResponse, _err error) {
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
		Action:      tea.String("ListDeploymentDrafts"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployment-drafts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListDeploymentDraftsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListDeploymentDraftsResponse{}
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
// list deploymentDrafts
//
// @param request - ListDeploymentDraftsRequest
//
// @return ListDeploymentDraftsResponse
func (client *Client) ListDeploymentDrafts(namespace *string, request *ListDeploymentDraftsRequest) (_result *ListDeploymentDraftsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListDeploymentDraftsHeaders{}
	_result = &ListDeploymentDraftsResponse{}
	_body, _err := client.ListDeploymentDraftsWithOptions(namespace, request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListDeploymentTargetsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListDeploymentTargetsResponse{}
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

	if !tea.BoolValue(util.IsUnset(request.SortName)) {
		query["sortName"] = request.SortName
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListDeploymentsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListDeploymentsResponse{}
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListEditableNamespaceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListEditableNamespaceResponse{}
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListEngineVersionMetadataResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListEngineVersionMetadataResponse{}
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListJobsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListJobsResponse{}
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListMembersResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListMembersResponse{}
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListSavepointsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListSavepointsResponse{}
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
// 列表定时执行计划
//
// @param request - ListScheduledPlanRequest
//
// @param headers - ListScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScheduledPlanResponse
func (client *Client) ListScheduledPlanWithOptions(namespace *string, request *ListScheduledPlanRequest, headers *ListScheduledPlanHeaders, runtime *util.RuntimeOptions) (_result *ListScheduledPlanResponse, _err error) {
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
		Action:      tea.String("ListScheduledPlan"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/scheduled-plans"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListScheduledPlanResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListScheduledPlanResponse{}
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
// 列表定时执行计划
//
// @param request - ListScheduledPlanRequest
//
// @return ListScheduledPlanResponse
func (client *Client) ListScheduledPlan(namespace *string, request *ListScheduledPlanRequest) (_result *ListScheduledPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListScheduledPlanHeaders{}
	_result = &ListScheduledPlanResponse{}
	_body, _err := client.ListScheduledPlanWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取作业资源变更历史
//
// @param request - ListScheduledPlanExecutedHistoryRequest
//
// @param headers - ListScheduledPlanExecutedHistoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScheduledPlanExecutedHistoryResponse
func (client *Client) ListScheduledPlanExecutedHistoryWithOptions(namespace *string, request *ListScheduledPlanExecutedHistoryRequest, headers *ListScheduledPlanExecutedHistoryHeaders, runtime *util.RuntimeOptions) (_result *ListScheduledPlanExecutedHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeploymentId)) {
		query["deploymentId"] = request.DeploymentId
	}

	if !tea.BoolValue(util.IsUnset(request.Origin)) {
		query["origin"] = request.Origin
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
		Action:      tea.String("ListScheduledPlanExecutedHistory"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/job-resource-upgradings"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListScheduledPlanExecutedHistoryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListScheduledPlanExecutedHistoryResponse{}
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
// 获取作业资源变更历史
//
// @param request - ListScheduledPlanExecutedHistoryRequest
//
// @return ListScheduledPlanExecutedHistoryResponse
func (client *Client) ListScheduledPlanExecutedHistory(namespace *string, request *ListScheduledPlanExecutedHistoryRequest) (_result *ListScheduledPlanExecutedHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListScheduledPlanExecutedHistoryHeaders{}
	_result = &ListScheduledPlanExecutedHistoryResponse{}
	_body, _err := client.ListScheduledPlanExecutedHistoryWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举session集群
//
// @param headers - ListSessionClustersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSessionClustersResponse
func (client *Client) ListSessionClustersWithOptions(namespace *string, headers *ListSessionClustersHeaders, runtime *util.RuntimeOptions) (_result *ListSessionClustersResponse, _err error) {
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
		Action:      tea.String("ListSessionClusters"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/sessionclusters"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListSessionClustersResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListSessionClustersResponse{}
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
// 列举session集群
//
// @return ListSessionClustersResponse
func (client *Client) ListSessionClusters(namespace *string) (_result *ListSessionClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSessionClustersHeaders{}
	_result = &ListSessionClustersResponse{}
	_body, _err := client.ListSessionClustersWithOptions(namespace, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListVariablesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListVariablesResponse{}
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

// Summary:
//
// Registers a custom connector in a namespace. The registered custom connector can be used in SQL statements.
//
// @param request - RegisterCustomConnectorRequest
//
// @param headers - RegisterCustomConnectorHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterCustomConnectorResponse
func (client *Client) RegisterCustomConnectorWithOptions(namespace *string, request *RegisterCustomConnectorRequest, headers *RegisterCustomConnectorHeaders, runtime *util.RuntimeOptions) (_result *RegisterCustomConnectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JarUrl)) {
		query["jarUrl"] = request.JarUrl
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
		Action:      tea.String("RegisterCustomConnector"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/connectors%3Aregister"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RegisterCustomConnectorResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RegisterCustomConnectorResponse{}
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
// Registers a custom connector in a namespace. The registered custom connector can be used in SQL statements.
//
// @param request - RegisterCustomConnectorRequest
//
// @return RegisterCustomConnectorResponse
func (client *Client) RegisterCustomConnector(namespace *string, request *RegisterCustomConnectorRequest) (_result *RegisterCustomConnectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RegisterCustomConnectorHeaders{}
	_result = &RegisterCustomConnectorResponse{}
	_body, _err := client.RegisterCustomConnectorWithOptions(namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Registers specific or all of the user-defined functions (UDFs) that are parsed from the JAR files. The registered functions can be used in SQL statements.
//
// @param request - RegisterUdfFunctionRequest
//
// @param headers - RegisterUdfFunctionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterUdfFunctionResponse
func (client *Client) RegisterUdfFunctionWithOptions(namespace *string, request *RegisterUdfFunctionRequest, headers *RegisterUdfFunctionHeaders, runtime *util.RuntimeOptions) (_result *RegisterUdfFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassName)) {
		query["className"] = request.ClassName
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionName)) {
		query["functionName"] = request.FunctionName
	}

	if !tea.BoolValue(util.IsUnset(request.UdfArtifactName)) {
		query["udfArtifactName"] = request.UdfArtifactName
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
		Action:      tea.String("RegisterUdfFunction"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/udfartifacts/function%3AregisterUdfFunction"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RegisterUdfFunctionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RegisterUdfFunctionResponse{}
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
// Registers specific or all of the user-defined functions (UDFs) that are parsed from the JAR files. The registered functions can be used in SQL statements.
//
// @param request - RegisterUdfFunctionRequest
//
// @return RegisterUdfFunctionResponse
func (client *Client) RegisterUdfFunction(namespace *string, request *RegisterUdfFunctionRequest) (_result *RegisterUdfFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RegisterUdfFunctionHeaders{}
	_result = &RegisterUdfFunctionResponse{}
	_body, _err := client.RegisterUdfFunctionWithOptions(namespace, request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StartJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StartJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StartJobWithParamsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StartJobWithParamsResponse{}
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
// 启动session集群
//
// @param headers - StartSessionClusterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSessionClusterResponse
func (client *Client) StartSessionClusterWithOptions(namespace *string, sessionClusterName *string, headers *StartSessionClusterHeaders, runtime *util.RuntimeOptions) (_result *StartSessionClusterResponse, _err error) {
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
		Action:      tea.String("StartSessionCluster"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/sessionclusters/" + tea.StringValue(openapiutil.GetEncodeParam(sessionClusterName)) + "%3Astart"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StartSessionClusterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StartSessionClusterResponse{}
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
// 启动session集群
//
// @return StartSessionClusterResponse
func (client *Client) StartSessionCluster(namespace *string, sessionClusterName *string) (_result *StartSessionClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StartSessionClusterHeaders{}
	_result = &StartSessionClusterResponse{}
	_body, _err := client.StartSessionClusterWithOptions(namespace, sessionClusterName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止应用执行定时计划
//
// @param headers - StopApplyScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopApplyScheduledPlanResponse
func (client *Client) StopApplyScheduledPlanWithOptions(namespace *string, scheduledPlanId *string, headers *StopApplyScheduledPlanHeaders, runtime *util.RuntimeOptions) (_result *StopApplyScheduledPlanResponse, _err error) {
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
		Action:      tea.String("StopApplyScheduledPlan"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/scheduled-plans/" + tea.StringValue(openapiutil.GetEncodeParam(scheduledPlanId)) + "%3Astop"),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StopApplyScheduledPlanResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StopApplyScheduledPlanResponse{}
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
// 停止应用执行定时计划
//
// @return StopApplyScheduledPlanResponse
func (client *Client) StopApplyScheduledPlan(namespace *string, scheduledPlanId *string) (_result *StopApplyScheduledPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StopApplyScheduledPlanHeaders{}
	_result = &StopApplyScheduledPlanResponse{}
	_body, _err := client.StopApplyScheduledPlanWithOptions(namespace, scheduledPlanId, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StopJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StopJobResponse{}
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
// 停止session集群
//
// @param headers - StopSessionClusterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopSessionClusterResponse
func (client *Client) StopSessionClusterWithOptions(namespace *string, sessionClusterName *string, headers *StopSessionClusterHeaders, runtime *util.RuntimeOptions) (_result *StopSessionClusterResponse, _err error) {
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
		Action:      tea.String("StopSessionCluster"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/sessionclusters/" + tea.StringValue(openapiutil.GetEncodeParam(sessionClusterName)) + "%3Astop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StopSessionClusterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StopSessionClusterResponse{}
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
// 停止session集群
//
// @return StopSessionClusterResponse
func (client *Client) StopSessionCluster(namespace *string, sessionClusterName *string) (_result *StopSessionClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StopSessionClusterHeaders{}
	_result = &StopSessionClusterResponse{}
	_body, _err := client.StopSessionClusterWithOptions(namespace, sessionClusterName, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateDeploymentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateDeploymentResponse{}
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
// update a deploymentDraft
//
// @param request - UpdateDeploymentDraftRequest
//
// @param headers - UpdateDeploymentDraftHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeploymentDraftResponse
func (client *Client) UpdateDeploymentDraftWithOptions(namespace *string, deploymentDraftId *string, request *UpdateDeploymentDraftRequest, headers *UpdateDeploymentDraftHeaders, runtime *util.RuntimeOptions) (_result *UpdateDeploymentDraftResponse, _err error) {
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
		Action:      tea.String("UpdateDeploymentDraft"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployment-drafts/" + tea.StringValue(openapiutil.GetEncodeParam(deploymentDraftId))),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateDeploymentDraftResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateDeploymentDraftResponse{}
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
// update a deploymentDraft
//
// @param request - UpdateDeploymentDraftRequest
//
// @return UpdateDeploymentDraftResponse
func (client *Client) UpdateDeploymentDraft(namespace *string, deploymentDraftId *string, request *UpdateDeploymentDraftRequest) (_result *UpdateDeploymentDraftResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateDeploymentDraftHeaders{}
	_result = &UpdateDeploymentDraftResponse{}
	_body, _err := client.UpdateDeploymentDraftWithOptions(namespace, deploymentDraftId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改deploymentTarget
//
// @param request - UpdateDeploymentTargetRequest
//
// @param headers - UpdateDeploymentTargetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeploymentTargetResponse
func (client *Client) UpdateDeploymentTargetWithOptions(namespace *string, deploymentTargetName *string, request *UpdateDeploymentTargetRequest, headers *UpdateDeploymentTargetHeaders, runtime *util.RuntimeOptions) (_result *UpdateDeploymentTargetResponse, _err error) {
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
		Action:      tea.String("UpdateDeploymentTarget"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/deployment-targets/" + tea.StringValue(openapiutil.GetEncodeParam(deploymentTargetName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateDeploymentTargetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateDeploymentTargetResponse{}
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
// 修改deploymentTarget
//
// @param request - UpdateDeploymentTargetRequest
//
// @return UpdateDeploymentTargetResponse
func (client *Client) UpdateDeploymentTarget(namespace *string, deploymentTargetName *string, request *UpdateDeploymentTargetRequest) (_result *UpdateDeploymentTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateDeploymentTargetHeaders{}
	_result = &UpdateDeploymentTargetResponse{}
	_body, _err := client.UpdateDeploymentTargetWithOptions(namespace, deploymentTargetName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// update a folder
//
// @param request - UpdateFolderRequest
//
// @param headers - UpdateFolderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFolderResponse
func (client *Client) UpdateFolderWithOptions(namespace *string, folderId *string, request *UpdateFolderRequest, headers *UpdateFolderHeaders, runtime *util.RuntimeOptions) (_result *UpdateFolderResponse, _err error) {
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
		Action:      tea.String("UpdateFolder"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/folder/" + tea.StringValue(openapiutil.GetEncodeParam(folderId))),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateFolderResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateFolderResponse{}
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
// update a folder
//
// @param request - UpdateFolderRequest
//
// @return UpdateFolderResponse
func (client *Client) UpdateFolder(namespace *string, folderId *string, request *UpdateFolderRequest) (_result *UpdateFolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateFolderHeaders{}
	_result = &UpdateFolderResponse{}
	_body, _err := client.UpdateFolderWithOptions(namespace, folderId, request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateMemberResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateMemberResponse{}
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
// 更新定时执行计划
//
// @param request - UpdateScheduledPlanRequest
//
// @param headers - UpdateScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScheduledPlanResponse
func (client *Client) UpdateScheduledPlanWithOptions(namespace *string, scheduledPlanId *string, request *UpdateScheduledPlanRequest, headers *UpdateScheduledPlanHeaders, runtime *util.RuntimeOptions) (_result *UpdateScheduledPlanResponse, _err error) {
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
		Action:      tea.String("UpdateScheduledPlan"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/scheduled-plans/" + tea.StringValue(openapiutil.GetEncodeParam(scheduledPlanId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateScheduledPlanResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateScheduledPlanResponse{}
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
// 更新定时执行计划
//
// @param request - UpdateScheduledPlanRequest
//
// @return UpdateScheduledPlanResponse
func (client *Client) UpdateScheduledPlan(namespace *string, scheduledPlanId *string, request *UpdateScheduledPlanRequest) (_result *UpdateScheduledPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateScheduledPlanHeaders{}
	_result = &UpdateScheduledPlanResponse{}
	_body, _err := client.UpdateScheduledPlanWithOptions(namespace, scheduledPlanId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新session集群
//
// @param request - UpdateSessionClusterRequest
//
// @param headers - UpdateSessionClusterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSessionClusterResponse
func (client *Client) UpdateSessionClusterWithOptions(namespace *string, sessionClusterName *string, request *UpdateSessionClusterRequest, headers *UpdateSessionClusterHeaders, runtime *util.RuntimeOptions) (_result *UpdateSessionClusterResponse, _err error) {
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
		Action:      tea.String("UpdateSessionCluster"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/sessionclusters/" + tea.StringValue(openapiutil.GetEncodeParam(sessionClusterName))),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateSessionClusterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateSessionClusterResponse{}
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
// 更新session集群
//
// @param request - UpdateSessionClusterRequest
//
// @return UpdateSessionClusterResponse
func (client *Client) UpdateSessionCluster(namespace *string, sessionClusterName *string, request *UpdateSessionClusterRequest) (_result *UpdateSessionClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateSessionClusterHeaders{}
	_result = &UpdateSessionClusterResponse{}
	_body, _err := client.UpdateSessionClusterWithOptions(namespace, sessionClusterName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the JAR file of the user-defined function (UDF) that you create.
//
// @param request - UpdateUdfArtifactRequest
//
// @param headers - UpdateUdfArtifactHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUdfArtifactResponse
func (client *Client) UpdateUdfArtifactWithOptions(namespace *string, udfArtifactName *string, request *UpdateUdfArtifactRequest, headers *UpdateUdfArtifactHeaders, runtime *util.RuntimeOptions) (_result *UpdateUdfArtifactResponse, _err error) {
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
		Action:      tea.String("UpdateUdfArtifact"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/udfartifacts/" + tea.StringValue(openapiutil.GetEncodeParam(udfArtifactName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateUdfArtifactResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateUdfArtifactResponse{}
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
// Updates the JAR file of the user-defined function (UDF) that you create.
//
// @param request - UpdateUdfArtifactRequest
//
// @return UpdateUdfArtifactResponse
func (client *Client) UpdateUdfArtifact(namespace *string, udfArtifactName *string, request *UpdateUdfArtifactRequest) (_result *UpdateUdfArtifactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateUdfArtifactHeaders{}
	_result = &UpdateUdfArtifactResponse{}
	_body, _err := client.UpdateUdfArtifactWithOptions(namespace, udfArtifactName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新秘钥
//
// @param request - UpdateVariableRequest
//
// @param headers - UpdateVariableHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVariableResponse
func (client *Client) UpdateVariableWithOptions(namespace *string, name *string, request *UpdateVariableRequest, headers *UpdateVariableHeaders, runtime *util.RuntimeOptions) (_result *UpdateVariableResponse, _err error) {
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
		Action:      tea.String("UpdateVariable"),
		Version:     tea.String("2022-07-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/namespaces/" + tea.StringValue(openapiutil.GetEncodeParam(namespace)) + "/variables/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateVariableResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateVariableResponse{}
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
// 更新秘钥
//
// @param request - UpdateVariableRequest
//
// @return UpdateVariableResponse
func (client *Client) UpdateVariable(namespace *string, name *string, request *UpdateVariableRequest) (_result *UpdateVariableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateVariableHeaders{}
	_result = &UpdateVariableResponse{}
	_body, _err := client.UpdateVariableWithOptions(namespace, name, request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ValidateSqlStatementResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ValidateSqlStatementResponse{}
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
