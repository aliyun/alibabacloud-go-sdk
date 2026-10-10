// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *GetContextStoreResponseBody
	GetAgentSpace() *string
	SetConfig(v *GetContextStoreResponseBodyConfig) *GetContextStoreResponseBody
	GetConfig() *GetContextStoreResponseBodyConfig
	SetContextStoreName(v string) *GetContextStoreResponseBody
	GetContextStoreName() *string
	SetContextType(v string) *GetContextStoreResponseBody
	GetContextType() *string
	SetCreateTime(v string) *GetContextStoreResponseBody
	GetCreateTime() *string
	SetDescription(v string) *GetContextStoreResponseBody
	GetDescription() *string
	SetRegionId(v string) *GetContextStoreResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetContextStoreResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetContextStoreResponseBody
	GetStatus() *string
	SetUpdateTime(v string) *GetContextStoreResponseBody
	GetUpdateTime() *string
}

type GetContextStoreResponseBody struct {
	// The name of the AgentSpace to which the context store belongs.
	//
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// The configuration of the context store.
	Config *GetContextStoreResponseBodyConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The context store name.
	//
	// example:
	//
	// my-context-store
	ContextStoreName *string `json:"contextStoreName,omitempty" xml:"contextStoreName,omitempty"`
	// The type of the context store, such as experience or memory.
	//
	// example:
	//
	// experience
	ContextType *string `json:"contextType,omitempty" xml:"contextType,omitempty"`
	// The time when the context store was created, in ISO 8601 UTC format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the context store.
	//
	// example:
	//
	// 我的上下文库
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The region ID of the context store.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The status of the context store. Valid values:
	//
	// - ACTIVE
	//
	// - INITIALIZING
	//
	// - FAILED
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the context store was last updated, in ISO 8601 UTC format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-02T00:00:00Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetContextStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBody) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBody) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *GetContextStoreResponseBody) GetConfig() *GetContextStoreResponseBodyConfig {
	return s.Config
}

func (s *GetContextStoreResponseBody) GetContextStoreName() *string {
	return s.ContextStoreName
}

func (s *GetContextStoreResponseBody) GetContextType() *string {
	return s.ContextType
}

func (s *GetContextStoreResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetContextStoreResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetContextStoreResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetContextStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContextStoreResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetContextStoreResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetContextStoreResponseBody) SetAgentSpace(v string) *GetContextStoreResponseBody {
	s.AgentSpace = &v
	return s
}

func (s *GetContextStoreResponseBody) SetConfig(v *GetContextStoreResponseBodyConfig) *GetContextStoreResponseBody {
	s.Config = v
	return s
}

func (s *GetContextStoreResponseBody) SetContextStoreName(v string) *GetContextStoreResponseBody {
	s.ContextStoreName = &v
	return s
}

func (s *GetContextStoreResponseBody) SetContextType(v string) *GetContextStoreResponseBody {
	s.ContextType = &v
	return s
}

func (s *GetContextStoreResponseBody) SetCreateTime(v string) *GetContextStoreResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetContextStoreResponseBody) SetDescription(v string) *GetContextStoreResponseBody {
	s.Description = &v
	return s
}

func (s *GetContextStoreResponseBody) SetRegionId(v string) *GetContextStoreResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetContextStoreResponseBody) SetRequestId(v string) *GetContextStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContextStoreResponseBody) SetStatus(v string) *GetContextStoreResponseBody {
	s.Status = &v
	return s
}

func (s *GetContextStoreResponseBody) SetUpdateTime(v string) *GetContextStoreResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetContextStoreResponseBody) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContextStoreResponseBodyConfig struct {
	Audit            *GetContextStoreResponseBodyConfigAudit            `json:"audit,omitempty" xml:"audit,omitempty" type:"Struct"`
	ExtractionPolicy *GetContextStoreResponseBodyConfigExtractionPolicy `json:"extractionPolicy,omitempty" xml:"extractionPolicy,omitempty" type:"Struct"`
	InnerSource      *GetContextStoreResponseBodyConfigInnerSource      `json:"innerSource,omitempty" xml:"innerSource,omitempty" type:"Struct"`
	// The metadata field mapping. The key is the business field and the value is the storage field.
	//
	// example:
	//
	// {"userId":"user_id","sessionId":"session_id"}
	MetadataField map[string]*string `json:"metadataField,omitempty" xml:"metadataField,omitempty"`
	// The experience mining interval. Valid values: 1h, 6h, 12h, and 1d. Default value: 1d.
	//
	// example:
	//
	// 1d
	MiningInterval *string                                         `json:"miningInterval,omitempty" xml:"miningInterval,omitempty"`
	Observability  *GetContextStoreResponseBodyConfigObservability `json:"observability,omitempty" xml:"observability,omitempty" type:"Struct"`
	OutputDataset  *GetContextStoreResponseBodyConfigOutputDataset `json:"outputDataset,omitempty" xml:"outputDataset,omitempty" type:"Struct"`
	ScopePolicy    *GetContextStoreResponseBodyConfigScopePolicy   `json:"scopePolicy,omitempty" xml:"scopePolicy,omitempty" type:"Struct"`
	// The list of service names. This works together with source.agentSpace to locate the trace data source. This value cannot be changed in the current version.
	//
	// example:
	//
	// ["order-service","payment-service"]
	ServiceNames []*string `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
	// The datasource config passed in by the user. This serves only as the root identifier of the data source.
	Source        *GetContextStoreResponseBodyConfigSource        `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
	SourceStatus  *GetContextStoreResponseBodyConfigSourceStatus  `json:"sourceStatus,omitempty" xml:"sourceStatus,omitempty" type:"Struct"`
	StoragePolicy *GetContextStoreResponseBodyConfigStoragePolicy `json:"storagePolicy,omitempty" xml:"storagePolicy,omitempty" type:"Struct"`
	// example:
	//
	// 1
	StrategyVersion *int32 `json:"strategyVersion,omitempty" xml:"strategyVersion,omitempty"`
}

func (s GetContextStoreResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfig) GetAudit() *GetContextStoreResponseBodyConfigAudit {
	return s.Audit
}

func (s *GetContextStoreResponseBodyConfig) GetExtractionPolicy() *GetContextStoreResponseBodyConfigExtractionPolicy {
	return s.ExtractionPolicy
}

func (s *GetContextStoreResponseBodyConfig) GetInnerSource() *GetContextStoreResponseBodyConfigInnerSource {
	return s.InnerSource
}

func (s *GetContextStoreResponseBodyConfig) GetMetadataField() map[string]*string {
	return s.MetadataField
}

func (s *GetContextStoreResponseBodyConfig) GetMiningInterval() *string {
	return s.MiningInterval
}

func (s *GetContextStoreResponseBodyConfig) GetObservability() *GetContextStoreResponseBodyConfigObservability {
	return s.Observability
}

func (s *GetContextStoreResponseBodyConfig) GetOutputDataset() *GetContextStoreResponseBodyConfigOutputDataset {
	return s.OutputDataset
}

func (s *GetContextStoreResponseBodyConfig) GetScopePolicy() *GetContextStoreResponseBodyConfigScopePolicy {
	return s.ScopePolicy
}

func (s *GetContextStoreResponseBodyConfig) GetServiceNames() []*string {
	return s.ServiceNames
}

func (s *GetContextStoreResponseBodyConfig) GetSource() *GetContextStoreResponseBodyConfigSource {
	return s.Source
}

func (s *GetContextStoreResponseBodyConfig) GetSourceStatus() *GetContextStoreResponseBodyConfigSourceStatus {
	return s.SourceStatus
}

func (s *GetContextStoreResponseBodyConfig) GetStoragePolicy() *GetContextStoreResponseBodyConfigStoragePolicy {
	return s.StoragePolicy
}

func (s *GetContextStoreResponseBodyConfig) GetStrategyVersion() *int32 {
	return s.StrategyVersion
}

func (s *GetContextStoreResponseBodyConfig) SetAudit(v *GetContextStoreResponseBodyConfigAudit) *GetContextStoreResponseBodyConfig {
	s.Audit = v
	return s
}

func (s *GetContextStoreResponseBodyConfig) SetExtractionPolicy(v *GetContextStoreResponseBodyConfigExtractionPolicy) *GetContextStoreResponseBodyConfig {
	s.ExtractionPolicy = v
	return s
}

func (s *GetContextStoreResponseBodyConfig) SetInnerSource(v *GetContextStoreResponseBodyConfigInnerSource) *GetContextStoreResponseBodyConfig {
	s.InnerSource = v
	return s
}

func (s *GetContextStoreResponseBodyConfig) SetMetadataField(v map[string]*string) *GetContextStoreResponseBodyConfig {
	s.MetadataField = v
	return s
}

func (s *GetContextStoreResponseBodyConfig) SetMiningInterval(v string) *GetContextStoreResponseBodyConfig {
	s.MiningInterval = &v
	return s
}

func (s *GetContextStoreResponseBodyConfig) SetObservability(v *GetContextStoreResponseBodyConfigObservability) *GetContextStoreResponseBodyConfig {
	s.Observability = v
	return s
}

func (s *GetContextStoreResponseBodyConfig) SetOutputDataset(v *GetContextStoreResponseBodyConfigOutputDataset) *GetContextStoreResponseBodyConfig {
	s.OutputDataset = v
	return s
}

func (s *GetContextStoreResponseBodyConfig) SetScopePolicy(v *GetContextStoreResponseBodyConfigScopePolicy) *GetContextStoreResponseBodyConfig {
	s.ScopePolicy = v
	return s
}

func (s *GetContextStoreResponseBodyConfig) SetServiceNames(v []*string) *GetContextStoreResponseBodyConfig {
	s.ServiceNames = v
	return s
}

func (s *GetContextStoreResponseBodyConfig) SetSource(v *GetContextStoreResponseBodyConfigSource) *GetContextStoreResponseBodyConfig {
	s.Source = v
	return s
}

func (s *GetContextStoreResponseBodyConfig) SetSourceStatus(v *GetContextStoreResponseBodyConfigSourceStatus) *GetContextStoreResponseBodyConfig {
	s.SourceStatus = v
	return s
}

func (s *GetContextStoreResponseBodyConfig) SetStoragePolicy(v *GetContextStoreResponseBodyConfigStoragePolicy) *GetContextStoreResponseBodyConfig {
	s.StoragePolicy = v
	return s
}

func (s *GetContextStoreResponseBodyConfig) SetStrategyVersion(v int32) *GetContextStoreResponseBodyConfig {
	s.StrategyVersion = &v
	return s
}

func (s *GetContextStoreResponseBodyConfig) Validate() error {
	if s.Audit != nil {
		if err := s.Audit.Validate(); err != nil {
			return err
		}
	}
	if s.ExtractionPolicy != nil {
		if err := s.ExtractionPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.InnerSource != nil {
		if err := s.InnerSource.Validate(); err != nil {
			return err
		}
	}
	if s.Observability != nil {
		if err := s.Observability.Validate(); err != nil {
			return err
		}
	}
	if s.OutputDataset != nil {
		if err := s.OutputDataset.Validate(); err != nil {
			return err
		}
	}
	if s.ScopePolicy != nil {
		if err := s.ScopePolicy.Validate(); err != nil {
			return err
		}
	}
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	if s.SourceStatus != nil {
		if err := s.SourceStatus.Validate(); err != nil {
			return err
		}
	}
	if s.StoragePolicy != nil {
		if err := s.StoragePolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContextStoreResponseBodyConfigAudit struct {
	// example:
	//
	// false
	DroppedCandidates *bool `json:"droppedCandidates,omitempty" xml:"droppedCandidates,omitempty"`
	// example:
	//
	// raw
	QueryMode *string `json:"queryMode,omitempty" xml:"queryMode,omitempty"`
	// example:
	//
	// 30
	RetentionDays *int32 `json:"retentionDays,omitempty" xml:"retentionDays,omitempty"`
}

func (s GetContextStoreResponseBodyConfigAudit) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigAudit) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigAudit) GetDroppedCandidates() *bool {
	return s.DroppedCandidates
}

func (s *GetContextStoreResponseBodyConfigAudit) GetQueryMode() *string {
	return s.QueryMode
}

func (s *GetContextStoreResponseBodyConfigAudit) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *GetContextStoreResponseBodyConfigAudit) SetDroppedCandidates(v bool) *GetContextStoreResponseBodyConfigAudit {
	s.DroppedCandidates = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigAudit) SetQueryMode(v string) *GetContextStoreResponseBodyConfigAudit {
	s.QueryMode = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigAudit) SetRetentionDays(v int32) *GetContextStoreResponseBodyConfigAudit {
	s.RetentionDays = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigAudit) Validate() error {
	return dara.Validate(s)
}

type GetContextStoreResponseBodyConfigExtractionPolicy struct {
	// example:
	//
	// ["preference","profile"]
	Categories []*string `json:"categories,omitempty" xml:"categories,omitempty" type:"Repeated"`
	// example:
	//
	// 只抽取用户的产品偏好
	CustomInstructions *string `json:"customInstructions,omitempty" xml:"customInstructions,omitempty"`
	// example:
	//
	// ["密码","证件号"]
	ExcludeRules []*string                                               `json:"excludeRules,omitempty" xml:"excludeRules,omitempty" type:"Repeated"`
	Model        *GetContextStoreResponseBodyConfigExtractionPolicyModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	// example:
	//
	// fact
	Preset *string `json:"preset,omitempty" xml:"preset,omitempty"`
}

func (s GetContextStoreResponseBodyConfigExtractionPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigExtractionPolicy) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicy) GetCategories() []*string {
	return s.Categories
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicy) GetCustomInstructions() *string {
	return s.CustomInstructions
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicy) GetExcludeRules() []*string {
	return s.ExcludeRules
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicy) GetModel() *GetContextStoreResponseBodyConfigExtractionPolicyModel {
	return s.Model
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicy) GetPreset() *string {
	return s.Preset
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicy) SetCategories(v []*string) *GetContextStoreResponseBodyConfigExtractionPolicy {
	s.Categories = v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicy) SetCustomInstructions(v string) *GetContextStoreResponseBodyConfigExtractionPolicy {
	s.CustomInstructions = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicy) SetExcludeRules(v []*string) *GetContextStoreResponseBodyConfigExtractionPolicy {
	s.ExcludeRules = v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicy) SetModel(v *GetContextStoreResponseBodyConfigExtractionPolicyModel) *GetContextStoreResponseBodyConfigExtractionPolicy {
	s.Model = v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicy) SetPreset(v string) *GetContextStoreResponseBodyConfigExtractionPolicy {
	s.Preset = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicy) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContextStoreResponseBodyConfigExtractionPolicyModel struct {
	// example:
	//
	// qwen3.8-flash
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetContextStoreResponseBodyConfigExtractionPolicyModel) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigExtractionPolicyModel) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicyModel) GetName() *string {
	return s.Name
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicyModel) SetName(v string) *GetContextStoreResponseBodyConfigExtractionPolicyModel {
	s.Name = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicyModel) Validate() error {
	return dara.Validate(s)
}

type GetContextStoreResponseBodyConfigInnerSource struct {
	// example:
	//
	// memory_events_0a1b2c3d
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// agentloop-xxx
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s GetContextStoreResponseBodyConfigInnerSource) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigInnerSource) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigInnerSource) GetLogstore() *string {
	return s.Logstore
}

func (s *GetContextStoreResponseBodyConfigInnerSource) GetProject() *string {
	return s.Project
}

func (s *GetContextStoreResponseBodyConfigInnerSource) SetLogstore(v string) *GetContextStoreResponseBodyConfigInnerSource {
	s.Logstore = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigInnerSource) SetProject(v string) *GetContextStoreResponseBodyConfigInnerSource {
	s.Project = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigInnerSource) Validate() error {
	return dara.Validate(s)
}

type GetContextStoreResponseBodyConfigObservability struct {
	// example:
	//
	// memory-audit
	AuditLogstore *string `json:"auditLogstore,omitempty" xml:"auditLogstore,omitempty"`
	// example:
	//
	// memory_events_0a1b2c3d
	EventsLogstore *string `json:"eventsLogstore,omitempty" xml:"eventsLogstore,omitempty"`
	// example:
	//
	// agentloop-xxx
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s GetContextStoreResponseBodyConfigObservability) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigObservability) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigObservability) GetAuditLogstore() *string {
	return s.AuditLogstore
}

func (s *GetContextStoreResponseBodyConfigObservability) GetEventsLogstore() *string {
	return s.EventsLogstore
}

func (s *GetContextStoreResponseBodyConfigObservability) GetProject() *string {
	return s.Project
}

func (s *GetContextStoreResponseBodyConfigObservability) SetAuditLogstore(v string) *GetContextStoreResponseBodyConfigObservability {
	s.AuditLogstore = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigObservability) SetEventsLogstore(v string) *GetContextStoreResponseBodyConfigObservability {
	s.EventsLogstore = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigObservability) SetProject(v string) *GetContextStoreResponseBodyConfigObservability {
	s.Project = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigObservability) Validate() error {
	return dara.Validate(s)
}

type GetContextStoreResponseBodyConfigOutputDataset struct {
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// example:
	//
	// memory-my-context-store
	DatasetName *string `json:"datasetName,omitempty" xml:"datasetName,omitempty"`
	// example:
	//
	// MemoryRecordV1
	SchemaContract *string `json:"schemaContract,omitempty" xml:"schemaContract,omitempty"`
	// example:
	//
	// 1
	SchemaVersion *int32 `json:"schemaVersion,omitempty" xml:"schemaVersion,omitempty"`
}

func (s GetContextStoreResponseBodyConfigOutputDataset) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigOutputDataset) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigOutputDataset) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *GetContextStoreResponseBodyConfigOutputDataset) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetContextStoreResponseBodyConfigOutputDataset) GetSchemaContract() *string {
	return s.SchemaContract
}

func (s *GetContextStoreResponseBodyConfigOutputDataset) GetSchemaVersion() *int32 {
	return s.SchemaVersion
}

func (s *GetContextStoreResponseBodyConfigOutputDataset) SetAgentSpace(v string) *GetContextStoreResponseBodyConfigOutputDataset {
	s.AgentSpace = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigOutputDataset) SetDatasetName(v string) *GetContextStoreResponseBodyConfigOutputDataset {
	s.DatasetName = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigOutputDataset) SetSchemaContract(v string) *GetContextStoreResponseBodyConfigOutputDataset {
	s.SchemaContract = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigOutputDataset) SetSchemaVersion(v int32) *GetContextStoreResponseBodyConfigOutputDataset {
	s.SchemaVersion = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigOutputDataset) Validate() error {
	return dara.Validate(s)
}

type GetContextStoreResponseBodyConfigScopePolicy struct {
	// example:
	//
	// ["userId"]
	RequiredAnyOf []*string `json:"requiredAnyOf,omitempty" xml:"requiredAnyOf,omitempty" type:"Repeated"`
}

func (s GetContextStoreResponseBodyConfigScopePolicy) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigScopePolicy) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigScopePolicy) GetRequiredAnyOf() []*string {
	return s.RequiredAnyOf
}

func (s *GetContextStoreResponseBodyConfigScopePolicy) SetRequiredAnyOf(v []*string) *GetContextStoreResponseBodyConfigScopePolicy {
	s.RequiredAnyOf = v
	return s
}

func (s *GetContextStoreResponseBodyConfigScopePolicy) Validate() error {
	return dara.Validate(s)
}

type GetContextStoreResponseBodyConfigSource struct {
	// The AgentSpace where the trace data source resides. This is the same as the AgentSpace specified during creation.
	//
	// example:
	//
	// my-agent-space
	AgentSpace *string                                         `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	Dataset    *GetContextStoreResponseBodyConfigSourceDataset `json:"dataset,omitempty" xml:"dataset,omitempty" type:"Struct"`
	// The start time for data backfill, in ISO 8601 UTC format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	StartTime  *string                                            `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Trajectory *GetContextStoreResponseBodyConfigSourceTrajectory `json:"trajectory,omitempty" xml:"trajectory,omitempty" type:"Struct"`
	// example:
	//
	// trajectory
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetContextStoreResponseBodyConfigSource) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigSource) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigSource) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *GetContextStoreResponseBodyConfigSource) GetDataset() *GetContextStoreResponseBodyConfigSourceDataset {
	return s.Dataset
}

func (s *GetContextStoreResponseBodyConfigSource) GetStartTime() *string {
	return s.StartTime
}

func (s *GetContextStoreResponseBodyConfigSource) GetTrajectory() *GetContextStoreResponseBodyConfigSourceTrajectory {
	return s.Trajectory
}

func (s *GetContextStoreResponseBodyConfigSource) GetType() *string {
	return s.Type
}

func (s *GetContextStoreResponseBodyConfigSource) SetAgentSpace(v string) *GetContextStoreResponseBodyConfigSource {
	s.AgentSpace = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSource) SetDataset(v *GetContextStoreResponseBodyConfigSourceDataset) *GetContextStoreResponseBodyConfigSource {
	s.Dataset = v
	return s
}

func (s *GetContextStoreResponseBodyConfigSource) SetStartTime(v string) *GetContextStoreResponseBodyConfigSource {
	s.StartTime = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSource) SetTrajectory(v *GetContextStoreResponseBodyConfigSourceTrajectory) *GetContextStoreResponseBodyConfigSource {
	s.Trajectory = v
	return s
}

func (s *GetContextStoreResponseBodyConfigSource) SetType(v string) *GetContextStoreResponseBodyConfigSource {
	s.Type = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSource) Validate() error {
	if s.Dataset != nil {
		if err := s.Dataset.Validate(); err != nil {
			return err
		}
	}
	if s.Trajectory != nil {
		if err := s.Trajectory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContextStoreResponseBodyConfigSourceDataset struct {
	CustomFields []*GetContextStoreResponseBodyConfigSourceDatasetCustomFields `json:"customFields,omitempty" xml:"customFields,omitempty" type:"Repeated"`
	// example:
	//
	// trajectory-with-crm-profile
	DatasetName *string                                               `json:"datasetName,omitempty" xml:"datasetName,omitempty"`
	Filter      *GetContextStoreResponseBodyConfigSourceDatasetFilter `json:"filter,omitempty" xml:"filter,omitempty" type:"Struct"`
	// example:
	//
	// 300
	PollIntervalSeconds *int32 `json:"pollIntervalSeconds,omitempty" xml:"pollIntervalSeconds,omitempty"`
	// example:
	//
	// MemorySourceV1
	SchemaContract *string                                                      `json:"schemaContract,omitempty" xml:"schemaContract,omitempty"`
	VersionPolicy  *GetContextStoreResponseBodyConfigSourceDatasetVersionPolicy `json:"versionPolicy,omitempty" xml:"versionPolicy,omitempty" type:"Struct"`
}

func (s GetContextStoreResponseBodyConfigSourceDataset) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigSourceDataset) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigSourceDataset) GetCustomFields() []*GetContextStoreResponseBodyConfigSourceDatasetCustomFields {
	return s.CustomFields
}

func (s *GetContextStoreResponseBodyConfigSourceDataset) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetContextStoreResponseBodyConfigSourceDataset) GetFilter() *GetContextStoreResponseBodyConfigSourceDatasetFilter {
	return s.Filter
}

func (s *GetContextStoreResponseBodyConfigSourceDataset) GetPollIntervalSeconds() *int32 {
	return s.PollIntervalSeconds
}

func (s *GetContextStoreResponseBodyConfigSourceDataset) GetSchemaContract() *string {
	return s.SchemaContract
}

func (s *GetContextStoreResponseBodyConfigSourceDataset) GetVersionPolicy() *GetContextStoreResponseBodyConfigSourceDatasetVersionPolicy {
	return s.VersionPolicy
}

func (s *GetContextStoreResponseBodyConfigSourceDataset) SetCustomFields(v []*GetContextStoreResponseBodyConfigSourceDatasetCustomFields) *GetContextStoreResponseBodyConfigSourceDataset {
	s.CustomFields = v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceDataset) SetDatasetName(v string) *GetContextStoreResponseBodyConfigSourceDataset {
	s.DatasetName = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceDataset) SetFilter(v *GetContextStoreResponseBodyConfigSourceDatasetFilter) *GetContextStoreResponseBodyConfigSourceDataset {
	s.Filter = v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceDataset) SetPollIntervalSeconds(v int32) *GetContextStoreResponseBodyConfigSourceDataset {
	s.PollIntervalSeconds = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceDataset) SetSchemaContract(v string) *GetContextStoreResponseBodyConfigSourceDataset {
	s.SchemaContract = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceDataset) SetVersionPolicy(v *GetContextStoreResponseBodyConfigSourceDatasetVersionPolicy) *GetContextStoreResponseBodyConfigSourceDataset {
	s.VersionPolicy = v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceDataset) Validate() error {
	if s.CustomFields != nil {
		for _, item := range s.CustomFields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	if s.VersionPolicy != nil {
		if err := s.VersionPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContextStoreResponseBodyConfigSourceDatasetCustomFields struct {
	// example:
	//
	// 客户等级
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// false
	Sensitive *bool `json:"sensitive,omitempty" xml:"sensitive,omitempty"`
	// example:
	//
	// customerTier
	SourceField *string `json:"sourceField,omitempty" xml:"sourceField,omitempty"`
	// example:
	//
	// metadata.customerTier
	Target *string `json:"target,omitempty" xml:"target,omitempty"`
	// example:
	//
	// extraction-input
	Usage *string `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s GetContextStoreResponseBodyConfigSourceDatasetCustomFields) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigSourceDatasetCustomFields) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetCustomFields) GetDescription() *string {
	return s.Description
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetCustomFields) GetSensitive() *bool {
	return s.Sensitive
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetCustomFields) GetSourceField() *string {
	return s.SourceField
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetCustomFields) GetTarget() *string {
	return s.Target
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetCustomFields) GetUsage() *string {
	return s.Usage
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetCustomFields) SetDescription(v string) *GetContextStoreResponseBodyConfigSourceDatasetCustomFields {
	s.Description = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetCustomFields) SetSensitive(v bool) *GetContextStoreResponseBodyConfigSourceDatasetCustomFields {
	s.Sensitive = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetCustomFields) SetSourceField(v string) *GetContextStoreResponseBodyConfigSourceDatasetCustomFields {
	s.SourceField = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetCustomFields) SetTarget(v string) *GetContextStoreResponseBodyConfigSourceDatasetCustomFields {
	s.Target = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetCustomFields) SetUsage(v string) *GetContextStoreResponseBodyConfigSourceDatasetCustomFields {
	s.Usage = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetCustomFields) Validate() error {
	return dara.Validate(s)
}

type GetContextStoreResponseBodyConfigSourceDatasetFilter struct {
	// example:
	//
	// appId = \\"crm-service\\"
	Where *string `json:"where,omitempty" xml:"where,omitempty"`
}

func (s GetContextStoreResponseBodyConfigSourceDatasetFilter) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigSourceDatasetFilter) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetFilter) GetWhere() *string {
	return s.Where
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetFilter) SetWhere(v string) *GetContextStoreResponseBodyConfigSourceDatasetFilter {
	s.Where = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetFilter) Validate() error {
	return dara.Validate(s)
}

type GetContextStoreResponseBodyConfigSourceDatasetVersionPolicy struct {
	// example:
	//
	// follow
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// example:
	//
	// 0
	StartSeq *int64 `json:"startSeq,omitempty" xml:"startSeq,omitempty"`
	// example:
	//
	// v3
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetContextStoreResponseBodyConfigSourceDatasetVersionPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigSourceDatasetVersionPolicy) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetVersionPolicy) GetMode() *string {
	return s.Mode
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetVersionPolicy) GetStartSeq() *int64 {
	return s.StartSeq
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetVersionPolicy) GetVersion() *string {
	return s.Version
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetVersionPolicy) SetMode(v string) *GetContextStoreResponseBodyConfigSourceDatasetVersionPolicy {
	s.Mode = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetVersionPolicy) SetStartSeq(v int64) *GetContextStoreResponseBodyConfigSourceDatasetVersionPolicy {
	s.StartSeq = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetVersionPolicy) SetVersion(v string) *GetContextStoreResponseBodyConfigSourceDatasetVersionPolicy {
	s.Version = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceDatasetVersionPolicy) Validate() error {
	return dara.Validate(s)
}

type GetContextStoreResponseBodyConfigSourceTrajectory struct {
	Filter *GetContextStoreResponseBodyConfigSourceTrajectoryFilter `json:"filter,omitempty" xml:"filter,omitempty" type:"Struct"`
	// example:
	//
	// agent-trajectory
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// 300
	PollIntervalSeconds *int32                                                         `json:"pollIntervalSeconds,omitempty" xml:"pollIntervalSeconds,omitempty"`
	ScopeMapping        *GetContextStoreResponseBodyConfigSourceTrajectoryScopeMapping `json:"scopeMapping,omitempty" xml:"scopeMapping,omitempty" type:"Struct"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-10-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetContextStoreResponseBodyConfigSourceTrajectory) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigSourceTrajectory) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectory) GetFilter() *GetContextStoreResponseBodyConfigSourceTrajectoryFilter {
	return s.Filter
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectory) GetLogstore() *string {
	return s.Logstore
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectory) GetPollIntervalSeconds() *int32 {
	return s.PollIntervalSeconds
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectory) GetScopeMapping() *GetContextStoreResponseBodyConfigSourceTrajectoryScopeMapping {
	return s.ScopeMapping
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectory) GetStartTime() *string {
	return s.StartTime
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectory) SetFilter(v *GetContextStoreResponseBodyConfigSourceTrajectoryFilter) *GetContextStoreResponseBodyConfigSourceTrajectory {
	s.Filter = v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectory) SetLogstore(v string) *GetContextStoreResponseBodyConfigSourceTrajectory {
	s.Logstore = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectory) SetPollIntervalSeconds(v int32) *GetContextStoreResponseBodyConfigSourceTrajectory {
	s.PollIntervalSeconds = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectory) SetScopeMapping(v *GetContextStoreResponseBodyConfigSourceTrajectoryScopeMapping) *GetContextStoreResponseBodyConfigSourceTrajectory {
	s.ScopeMapping = v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectory) SetStartTime(v string) *GetContextStoreResponseBodyConfigSourceTrajectory {
	s.StartTime = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectory) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	if s.ScopeMapping != nil {
		if err := s.ScopeMapping.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContextStoreResponseBodyConfigSourceTrajectoryFilter struct {
	// example:
	//
	// ["sales-copilot"]
	AgentNames []*string `json:"agentNames,omitempty" xml:"agentNames,omitempty" type:"Repeated"`
	// example:
	//
	// false
	ExcludeDegraded *bool `json:"excludeDegraded,omitempty" xml:"excludeDegraded,omitempty"`
	// example:
	//
	// 2
	MinStepCount *int32 `json:"minStepCount,omitempty" xml:"minStepCount,omitempty"`
	// example:
	//
	// tool_names:"search_order"
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// ["crm-service","app-*"]
	ServiceNames []*string `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
}

func (s GetContextStoreResponseBodyConfigSourceTrajectoryFilter) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigSourceTrajectoryFilter) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectoryFilter) GetAgentNames() []*string {
	return s.AgentNames
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectoryFilter) GetExcludeDegraded() *bool {
	return s.ExcludeDegraded
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectoryFilter) GetMinStepCount() *int32 {
	return s.MinStepCount
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectoryFilter) GetQuery() *string {
	return s.Query
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectoryFilter) GetServiceNames() []*string {
	return s.ServiceNames
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectoryFilter) SetAgentNames(v []*string) *GetContextStoreResponseBodyConfigSourceTrajectoryFilter {
	s.AgentNames = v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectoryFilter) SetExcludeDegraded(v bool) *GetContextStoreResponseBodyConfigSourceTrajectoryFilter {
	s.ExcludeDegraded = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectoryFilter) SetMinStepCount(v int32) *GetContextStoreResponseBodyConfigSourceTrajectoryFilter {
	s.MinStepCount = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectoryFilter) SetQuery(v string) *GetContextStoreResponseBodyConfigSourceTrajectoryFilter {
	s.Query = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectoryFilter) SetServiceNames(v []*string) *GetContextStoreResponseBodyConfigSourceTrajectoryFilter {
	s.ServiceNames = v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectoryFilter) Validate() error {
	return dara.Validate(s)
}

type GetContextStoreResponseBodyConfigSourceTrajectoryScopeMapping struct {
	// example:
	//
	// $.agent_name
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// $.service_names[0]
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// $.trajectory_id
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// example:
	//
	// $.trajectory_extensions.user_id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetContextStoreResponseBodyConfigSourceTrajectoryScopeMapping) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigSourceTrajectoryScopeMapping) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectoryScopeMapping) GetAgentId() *string {
	return s.AgentId
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectoryScopeMapping) GetAppId() *string {
	return s.AppId
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectoryScopeMapping) GetRunId() *string {
	return s.RunId
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectoryScopeMapping) GetUserId() *string {
	return s.UserId
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectoryScopeMapping) SetAgentId(v string) *GetContextStoreResponseBodyConfigSourceTrajectoryScopeMapping {
	s.AgentId = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectoryScopeMapping) SetAppId(v string) *GetContextStoreResponseBodyConfigSourceTrajectoryScopeMapping {
	s.AppId = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectoryScopeMapping) SetRunId(v string) *GetContextStoreResponseBodyConfigSourceTrajectoryScopeMapping {
	s.RunId = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectoryScopeMapping) SetUserId(v string) *GetContextStoreResponseBodyConfigSourceTrajectoryScopeMapping {
	s.UserId = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceTrajectoryScopeMapping) Validate() error {
	return dara.Validate(s)
}

type GetContextStoreResponseBodyConfigSourceStatus struct {
	Checkpoint map[string]interface{} `json:"checkpoint,omitempty" xml:"checkpoint,omitempty"`
	// example:
	//
	// 读取数据源超时
	LastError *string `json:"lastError,omitempty" xml:"lastError,omitempty"`
	// example:
	//
	// 2026-10-01T08:00:00Z
	LastWindowAt *string `json:"lastWindowAt,omitempty" xml:"lastWindowAt,omitempty"`
	// example:
	//
	// 0
	RetryCount *int32 `json:"retryCount,omitempty" xml:"retryCount,omitempty"`
	// example:
	//
	// Running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s GetContextStoreResponseBodyConfigSourceStatus) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigSourceStatus) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigSourceStatus) GetCheckpoint() map[string]interface{} {
	return s.Checkpoint
}

func (s *GetContextStoreResponseBodyConfigSourceStatus) GetLastError() *string {
	return s.LastError
}

func (s *GetContextStoreResponseBodyConfigSourceStatus) GetLastWindowAt() *string {
	return s.LastWindowAt
}

func (s *GetContextStoreResponseBodyConfigSourceStatus) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *GetContextStoreResponseBodyConfigSourceStatus) GetState() *string {
	return s.State
}

func (s *GetContextStoreResponseBodyConfigSourceStatus) SetCheckpoint(v map[string]interface{}) *GetContextStoreResponseBodyConfigSourceStatus {
	s.Checkpoint = v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceStatus) SetLastError(v string) *GetContextStoreResponseBodyConfigSourceStatus {
	s.LastError = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceStatus) SetLastWindowAt(v string) *GetContextStoreResponseBodyConfigSourceStatus {
	s.LastWindowAt = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceStatus) SetRetryCount(v int32) *GetContextStoreResponseBodyConfigSourceStatus {
	s.RetryCount = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceStatus) SetState(v string) *GetContextStoreResponseBodyConfigSourceStatus {
	s.State = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceStatus) Validate() error {
	return dara.Validate(s)
}

type GetContextStoreResponseBodyConfigStoragePolicy struct {
	// example:
	//
	// ["ADD","UPDATE","MERGE","DELETE"]
	AllowedActions []*string `json:"allowedActions,omitempty" xml:"allowedActions,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Dedupe *bool `json:"dedupe,omitempty" xml:"dedupe,omitempty"`
	// example:
	//
	// true
	HumanEditProtection *bool `json:"humanEditProtection,omitempty" xml:"humanEditProtection,omitempty"`
	// example:
	//
	// semantic
	MergeKey *string `json:"mergeKey,omitempty" xml:"mergeKey,omitempty"`
	// example:
	//
	// upsert
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// example:
	//
	// 0.4
	SimilarityThreshold *float64 `json:"similarityThreshold,omitempty" xml:"similarityThreshold,omitempty"`
	// example:
	//
	// 0
	TtlDays *int32 `json:"ttlDays,omitempty" xml:"ttlDays,omitempty"`
}

func (s GetContextStoreResponseBodyConfigStoragePolicy) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigStoragePolicy) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigStoragePolicy) GetAllowedActions() []*string {
	return s.AllowedActions
}

func (s *GetContextStoreResponseBodyConfigStoragePolicy) GetDedupe() *bool {
	return s.Dedupe
}

func (s *GetContextStoreResponseBodyConfigStoragePolicy) GetHumanEditProtection() *bool {
	return s.HumanEditProtection
}

func (s *GetContextStoreResponseBodyConfigStoragePolicy) GetMergeKey() *string {
	return s.MergeKey
}

func (s *GetContextStoreResponseBodyConfigStoragePolicy) GetMode() *string {
	return s.Mode
}

func (s *GetContextStoreResponseBodyConfigStoragePolicy) GetSimilarityThreshold() *float64 {
	return s.SimilarityThreshold
}

func (s *GetContextStoreResponseBodyConfigStoragePolicy) GetTtlDays() *int32 {
	return s.TtlDays
}

func (s *GetContextStoreResponseBodyConfigStoragePolicy) SetAllowedActions(v []*string) *GetContextStoreResponseBodyConfigStoragePolicy {
	s.AllowedActions = v
	return s
}

func (s *GetContextStoreResponseBodyConfigStoragePolicy) SetDedupe(v bool) *GetContextStoreResponseBodyConfigStoragePolicy {
	s.Dedupe = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigStoragePolicy) SetHumanEditProtection(v bool) *GetContextStoreResponseBodyConfigStoragePolicy {
	s.HumanEditProtection = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigStoragePolicy) SetMergeKey(v string) *GetContextStoreResponseBodyConfigStoragePolicy {
	s.MergeKey = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigStoragePolicy) SetMode(v string) *GetContextStoreResponseBodyConfigStoragePolicy {
	s.Mode = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigStoragePolicy) SetSimilarityThreshold(v float64) *GetContextStoreResponseBodyConfigStoragePolicy {
	s.SimilarityThreshold = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigStoragePolicy) SetTtlDays(v int32) *GetContextStoreResponseBodyConfigStoragePolicy {
	s.TtlDays = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigStoragePolicy) Validate() error {
	return dara.Validate(s)
}
