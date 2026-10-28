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
	// The name of the associated AgentSpace.
	//
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// The context library configuration.
	Config *GetContextStoreResponseBodyConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The name of the context store.
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
	// My context store
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the region where the context store is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The request ID used to locate the request during troubleshooting.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The status of the context store, such as ACTIVE, INITIALIZING, or FAILED.
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
	// The audit configuration.
	Audit *GetContextStoreResponseBodyConfigAudit `json:"audit,omitempty" xml:"audit,omitempty" type:"Struct"`
	// The extraction policies. You can configure 1 to 10 policies. The strategyVersion parameter indicates the snapshot version of the entire library configuration.
	ExtractionPolicies []*GetContextStoreResponseBodyConfigExtractionPolicies `json:"extractionPolicies,omitempty" xml:"extractionPolicies,omitempty" type:"Repeated"`
	// The extraction policy for the memory type.
	ExtractionPolicy *GetContextStoreResponseBodyConfigExtractionPolicy `json:"extractionPolicy,omitempty" xml:"extractionPolicy,omitempty" type:"Struct"`
	// The system-assigned coordinates of the event Logstore for the memory type. The Logstore is located in the Simple Log Service (SLS) project bound to the AgentSpace.
	InnerSource *GetContextStoreResponseBodyConfigInnerSource `json:"innerSource,omitempty" xml:"innerSource,omitempty" type:"Struct"`
	// The metadata field mapping. The key is the business field and the value is the storage field.
	//
	// example:
	//
	// {"userId":"user_id","sessionId":"session_id"}
	MetadataField map[string]*string `json:"metadataField,omitempty" xml:"metadataField,omitempty"`
	// The interval for experience mining. Valid values: 1h, 6h, 12h, and 1d. Default value: 1d.
	//
	// example:
	//
	// 1d
	MiningInterval *string `json:"miningInterval,omitempty" xml:"miningInterval,omitempty"`
	// The name of the internal running task associated with the server. For continuous libraries, it is used to view consumption records in read-only mode. For manual libraries, it can run limited mining. It is not inferred by the client and cannot replace the ingestion and recall criteria of the native experience library.
	//
	// example:
	//
	// experience-review-library
	MiningTaskName *string `json:"miningTaskName,omitempty" xml:"miningTaskName,omitempty"`
	// The observability resource coordinates (memory type).
	Observability *GetContextStoreResponseBodyConfigObservability `json:"observability,omitempty" xml:"observability,omitempty" type:"Struct"`
	// The location of the memory output dataset (memory type).
	OutputDataset *GetContextStoreResponseBodyConfigOutputDataset `json:"outputDataset,omitempty" xml:"outputDataset,omitempty" type:"Struct"`
	// The scope constraint policy.
	ScopePolicy *GetContextStoreResponseBodyConfigScopePolicy `json:"scopePolicy,omitempty" xml:"scopePolicy,omitempty" type:"Struct"`
	// The list of service names used to locate the trace data source together with source.agentSpace. This parameter cannot be modified in the current phase.
	//
	// example:
	//
	// ["order-service","payment-service"]
	ServiceNames []*string `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
	// The data source configuration provided by the user, which serves only as the root identifier of the data source.
	Source *GetContextStoreResponseBodyConfigSource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
	// The running status of the data source, applicable to the memory type and experience Dataset follow mode. The experience library submits a checkpoint only after it is published and automatically applied successfully. Pausing and restarting retain the checkpoint.
	SourceStatus *GetContextStoreResponseBodyConfigSourceStatus `json:"sourceStatus,omitempty" xml:"sourceStatus,omitempty" type:"Struct"`
	StorageModes []*string                                      `json:"storageModes,omitempty" xml:"storageModes,omitempty" type:"Repeated"`
	// The storage policy for the memory type.
	StoragePolicy *GetContextStoreResponseBodyConfigStoragePolicy `json:"storagePolicy,omitempty" xml:"storagePolicy,omitempty" type:"Struct"`
	// The currently effective strategy version number for the memory type.
	//
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

func (s *GetContextStoreResponseBodyConfig) GetExtractionPolicies() []*GetContextStoreResponseBodyConfigExtractionPolicies {
	return s.ExtractionPolicies
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

func (s *GetContextStoreResponseBodyConfig) GetMiningTaskName() *string {
	return s.MiningTaskName
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

func (s *GetContextStoreResponseBodyConfig) GetStorageModes() []*string {
	return s.StorageModes
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

func (s *GetContextStoreResponseBodyConfig) SetExtractionPolicies(v []*GetContextStoreResponseBodyConfigExtractionPolicies) *GetContextStoreResponseBodyConfig {
	s.ExtractionPolicies = v
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

func (s *GetContextStoreResponseBodyConfig) SetMiningTaskName(v string) *GetContextStoreResponseBodyConfig {
	s.MiningTaskName = &v
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

func (s *GetContextStoreResponseBodyConfig) SetStorageModes(v []*string) *GetContextStoreResponseBodyConfig {
	s.StorageModes = v
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
	if s.ExtractionPolicies != nil {
		for _, item := range s.ExtractionPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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
	// Specifies whether to write item events for dropped recall candidates. Default value: false.
	//
	// example:
	//
	// false
	DroppedCandidates *bool `json:"droppedCandidates,omitempty" xml:"droppedCandidates,omitempty"`
	// The recording mode for recall queries. Valid values: raw (plaintext) and hash (HMAC only). Default value: raw.
	//
	// example:
	//
	// raw
	QueryMode *string `json:"queryMode,omitempty" xml:"queryMode,omitempty"`
	// The number of days to retain audit logs. Valid values: 1 to 180. Default value: 30.
	//
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

type GetContextStoreResponseBodyConfigExtractionPolicies struct {
	// The list of memory categories.
	//
	// example:
	//
	// ["preference","profile"]
	Categories []*string `json:"categories,omitempty" xml:"categories,omitempty" type:"Repeated"`
	// The custom extraction instructions. This parameter is required when preset is set to custom. The length must be 1 to 8,000 characters.
	//
	// example:
	//
	// Extract only user product preferences
	CustomInstructions *string `json:"customInstructions,omitempty" xml:"customInstructions,omitempty"`
	// Specifies whether to apply this policy to new inputs. Default value: true. Disabling the policy does not delete historical memories.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of exclusion rules. Content that matches these rules is not extracted.
	//
	// example:
	//
	// ["Password","ID number"]
	ExcludeRules []*string `json:"excludeRules,omitempty" xml:"excludeRules,omitempty" type:"Repeated"`
	// The extraction model configuration.
	Model *GetContextStoreResponseBodyConfigExtractionPoliciesModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	// The display name of the extraction policy.
	//
	// example:
	//
	// User preferences
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The preset policy. Valid values: fact, toc-profile, tob-digital-twin, and custom. Default value: fact.
	//
	// example:
	//
	// fact
	Preset        *string                                                           `json:"preset,omitempty" xml:"preset,omitempty"`
	StoragePolicy *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy `json:"storagePolicy,omitempty" xml:"storagePolicy,omitempty" type:"Struct"`
	// The stable ID of the extraction policy. The ID must start with a letter and can contain only letters, digits, underscores, and hyphens. It can be up to 64 characters in length. The ID must be unique within the same context library. Modifying the name or instructions does not change this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// preferences
	StrategyId *string `json:"strategyId,omitempty" xml:"strategyId,omitempty"`
}

func (s GetContextStoreResponseBodyConfigExtractionPolicies) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigExtractionPolicies) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicies) GetCategories() []*string {
	return s.Categories
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicies) GetCustomInstructions() *string {
	return s.CustomInstructions
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicies) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicies) GetExcludeRules() []*string {
	return s.ExcludeRules
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicies) GetModel() *GetContextStoreResponseBodyConfigExtractionPoliciesModel {
	return s.Model
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicies) GetName() *string {
	return s.Name
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicies) GetPreset() *string {
	return s.Preset
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicies) GetStoragePolicy() *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy {
	return s.StoragePolicy
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicies) GetStrategyId() *string {
	return s.StrategyId
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicies) SetCategories(v []*string) *GetContextStoreResponseBodyConfigExtractionPolicies {
	s.Categories = v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicies) SetCustomInstructions(v string) *GetContextStoreResponseBodyConfigExtractionPolicies {
	s.CustomInstructions = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicies) SetEnabled(v bool) *GetContextStoreResponseBodyConfigExtractionPolicies {
	s.Enabled = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicies) SetExcludeRules(v []*string) *GetContextStoreResponseBodyConfigExtractionPolicies {
	s.ExcludeRules = v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicies) SetModel(v *GetContextStoreResponseBodyConfigExtractionPoliciesModel) *GetContextStoreResponseBodyConfigExtractionPolicies {
	s.Model = v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicies) SetName(v string) *GetContextStoreResponseBodyConfigExtractionPolicies {
	s.Name = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicies) SetPreset(v string) *GetContextStoreResponseBodyConfigExtractionPolicies {
	s.Preset = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicies) SetStoragePolicy(v *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy) *GetContextStoreResponseBodyConfigExtractionPolicies {
	s.StoragePolicy = v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicies) SetStrategyId(v string) *GetContextStoreResponseBodyConfigExtractionPolicies {
	s.StrategyId = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPolicies) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
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

type GetContextStoreResponseBodyConfigExtractionPoliciesModel struct {
	// The name of the extraction model. Currently, only qwen3.8-flash is supported. Internal platform credentials are used.
	//
	// example:
	//
	// qwen3.8-flash
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetContextStoreResponseBodyConfigExtractionPoliciesModel) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigExtractionPoliciesModel) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigExtractionPoliciesModel) GetName() *string {
	return s.Name
}

func (s *GetContextStoreResponseBodyConfigExtractionPoliciesModel) SetName(v string) *GetContextStoreResponseBodyConfigExtractionPoliciesModel {
	s.Name = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPoliciesModel) Validate() error {
	return dara.Validate(s)
}

type GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy struct {
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

func (s GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy) GetAllowedActions() []*string {
	return s.AllowedActions
}

func (s *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy) GetDedupe() *bool {
	return s.Dedupe
}

func (s *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy) GetHumanEditProtection() *bool {
	return s.HumanEditProtection
}

func (s *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy) GetMergeKey() *string {
	return s.MergeKey
}

func (s *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy) GetMode() *string {
	return s.Mode
}

func (s *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy) GetSimilarityThreshold() *float64 {
	return s.SimilarityThreshold
}

func (s *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy) GetTtlDays() *int32 {
	return s.TtlDays
}

func (s *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy) SetAllowedActions(v []*string) *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy {
	s.AllowedActions = v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy) SetDedupe(v bool) *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy {
	s.Dedupe = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy) SetHumanEditProtection(v bool) *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy {
	s.HumanEditProtection = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy) SetMergeKey(v string) *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy {
	s.MergeKey = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy) SetMode(v string) *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy {
	s.Mode = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy) SetSimilarityThreshold(v float64) *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy {
	s.SimilarityThreshold = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy) SetTtlDays(v int32) *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy {
	s.TtlDays = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigExtractionPoliciesStoragePolicy) Validate() error {
	return dara.Validate(s)
}

type GetContextStoreResponseBodyConfigExtractionPolicy struct {
	// The list of memory categories.
	//
	// example:
	//
	// ["preference","profile"]
	Categories []*string `json:"categories,omitempty" xml:"categories,omitempty" type:"Repeated"`
	// The custom extraction instructions. This parameter is required when preset is set to custom. The length must be 1 to 8,000 characters.
	//
	// example:
	//
	// Extract only user product preferences
	CustomInstructions *string `json:"customInstructions,omitempty" xml:"customInstructions,omitempty"`
	// The list of exclusion rules. Content that matches these rules is not extracted.
	//
	// example:
	//
	// ["Password","ID number"]
	ExcludeRules []*string `json:"excludeRules,omitempty" xml:"excludeRules,omitempty" type:"Repeated"`
	// The extraction model configuration.
	Model *GetContextStoreResponseBodyConfigExtractionPolicyModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	// The preset policy. Valid values: fact, toc-profile, tob-digital-twin, and custom. Default value: fact.
	//
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
	// The name of the extraction model. Currently, only qwen3.8-flash is supported. Internal platform credentials are used.
	//
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
	// The name of the event Logstore.
	//
	// example:
	//
	// memory_events_0a1b2c3d
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The name of the SLS project.
	//
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
	// The name of the audit Logstore.
	//
	// example:
	//
	// memory-audit
	AuditLogstore *string `json:"auditLogstore,omitempty" xml:"auditLogstore,omitempty"`
	// The name of the event Logstore.
	//
	// example:
	//
	// memory_events_0a1b2c3d
	EventsLogstore *string `json:"eventsLogstore,omitempty" xml:"eventsLogstore,omitempty"`
	// The Simple Log Service (SLS) project to which both belong, which is bound to the AgentSpace.
	//
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
	// The AgentSpace to which the dataset belongs.
	//
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// The name of the dataset.
	//
	// example:
	//
	// memory-my-context-store
	DatasetName *string `json:"datasetName,omitempty" xml:"datasetName,omitempty"`
	// The row schema contract of the dataset.
	//
	// example:
	//
	// MemoryRecordV1
	SchemaContract *string `json:"schemaContract,omitempty" xml:"schemaContract,omitempty"`
	// The schema version.
	//
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
	// Specifies that at least one of the scope fields must be non-empty. Valid values: userId, agentId, appId, and runId.
	//
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
	// The AgentSpace to which the trace data source belongs, which is the same as the AgentSpace to which it belonged when it was created.
	//
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// The DatasetBinding v1 shared with optimization tasks, which retains fixed column validation, supplementary column usage, and all/window selection.
	//
	// example:
	//
	// {"version":1,"selection":{"mode":"all"}}
	Binding map[string]interface{} `json:"binding,omitempty" xml:"binding,omitempty"`
	// The data source configuration for the memory type dataset. This parameter is required when the type is memory and source.type is dataset.
	Dataset *GetContextStoreResponseBodyConfigSourceDataset `json:"dataset,omitempty" xml:"dataset,omitempty" type:"Struct"`
	// The name of the source dataset for the experience type in the current AgentSpace.
	//
	// example:
	//
	// reviewed_trajectories
	DatasetName *string `json:"datasetName,omitempty" xml:"datasetName,omitempty"`
	// The input format of the dataset. For the experience base dataset, the source is atif.
	//
	// example:
	//
	// atif
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// The maximum number of records to read from the dataset at a time. Valid values: 1 to 1000.
	//
	// example:
	//
	// 1000
	MaxRecords *int32 `json:"maxRecords,omitempty" xml:"maxRecords,omitempty"`
	// The start time for data backfill in ISO 8601 UTC format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The data source configuration for the memory type trajectory. This parameter is required when the type is memory and source.type is trajectory.
	Trajectory *GetContextStoreResponseBodyConfigSourceTrajectory `json:"trajectory,omitempty" xml:"trajectory,omitempty" type:"Struct"`
	// The source type, which can be trajectory or dataset. For experience, if not specified, the trajectory configuration is used. For Dataset, the datasetName, format, and binding at the current level are used. The default mode is manual. If versionPolicy.mode is set to follow, continuous processing is performed and automatically applied to the experience library. For memory, if not specified, the mode is manual (only written by AddContexts). Otherwise, the trajectory or dataset sub-configuration with the same name is used. This parameter cannot be modified after creation.
	//
	// example:
	//
	// dataset
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The continuous consumption configuration for the native experience library Dataset. The default mode is manual. The follow mode only allows binding.selection.mode to be set to all, and consumes data by append sequence number. This parameter cannot be modified after creation.
	VersionPolicy *GetContextStoreResponseBodyConfigSourceVersionPolicy `json:"versionPolicy,omitempty" xml:"versionPolicy,omitempty" type:"Struct"`
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

func (s *GetContextStoreResponseBodyConfigSource) GetBinding() map[string]interface{} {
	return s.Binding
}

func (s *GetContextStoreResponseBodyConfigSource) GetDataset() *GetContextStoreResponseBodyConfigSourceDataset {
	return s.Dataset
}

func (s *GetContextStoreResponseBodyConfigSource) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetContextStoreResponseBodyConfigSource) GetFormat() *string {
	return s.Format
}

func (s *GetContextStoreResponseBodyConfigSource) GetMaxRecords() *int32 {
	return s.MaxRecords
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

func (s *GetContextStoreResponseBodyConfigSource) GetVersionPolicy() *GetContextStoreResponseBodyConfigSourceVersionPolicy {
	return s.VersionPolicy
}

func (s *GetContextStoreResponseBodyConfigSource) SetAgentSpace(v string) *GetContextStoreResponseBodyConfigSource {
	s.AgentSpace = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSource) SetBinding(v map[string]interface{}) *GetContextStoreResponseBodyConfigSource {
	s.Binding = v
	return s
}

func (s *GetContextStoreResponseBodyConfigSource) SetDataset(v *GetContextStoreResponseBodyConfigSourceDataset) *GetContextStoreResponseBodyConfigSource {
	s.Dataset = v
	return s
}

func (s *GetContextStoreResponseBodyConfigSource) SetDatasetName(v string) *GetContextStoreResponseBodyConfigSource {
	s.DatasetName = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSource) SetFormat(v string) *GetContextStoreResponseBodyConfigSource {
	s.Format = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSource) SetMaxRecords(v int32) *GetContextStoreResponseBodyConfigSource {
	s.MaxRecords = &v
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

func (s *GetContextStoreResponseBodyConfigSource) SetVersionPolicy(v *GetContextStoreResponseBodyConfigSourceVersionPolicy) *GetContextStoreResponseBodyConfigSource {
	s.VersionPolicy = v
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
	if s.VersionPolicy != nil {
		if err := s.VersionPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContextStoreResponseBodyConfigSourceDataset struct {
	// The list of custom field declarations.
	CustomFields []*GetContextStoreResponseBodyConfigSourceDatasetCustomFields `json:"customFields,omitempty" xml:"customFields,omitempty" type:"Repeated"`
	// The name of the dataset in the same AgentSpace. This parameter cannot be modified after creation.
	//
	// example:
	//
	// trajectory-with-crm-profile
	DatasetName *string `json:"datasetName,omitempty" xml:"datasetName,omitempty"`
	// The row filter condition.
	Filter *GetContextStoreResponseBodyConfigSourceDatasetFilter `json:"filter,omitempty" xml:"filter,omitempty" type:"Struct"`
	// The polling interval in seconds. Valid values: 60 to 3600. Default value: 300.
	//
	// example:
	//
	// 300
	PollIntervalSeconds *int32 `json:"pollIntervalSeconds,omitempty" xml:"pollIntervalSeconds,omitempty"`
	// The row schema contract of the dataset. In the current phase, this parameter is fixed to MemorySourceV1.
	//
	// example:
	//
	// MemorySourceV1
	SchemaContract *string `json:"schemaContract,omitempty" xml:"schemaContract,omitempty"`
	// The version consumption policy of the dataset.
	VersionPolicy *GetContextStoreResponseBodyConfigSourceDatasetVersionPolicy `json:"versionPolicy,omitempty" xml:"versionPolicy,omitempty" type:"Struct"`
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
	// The description of the field. This parameter is required when usage is extraction-input and is appended to the extraction prompt.
	//
	// example:
	//
	// Customer tier
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether the field is a sensitive field.
	//
	// example:
	//
	// false
	Sensitive *bool `json:"sensitive,omitempty" xml:"sensitive,omitempty"`
	// The name of the source field.
	//
	// example:
	//
	// customerTier
	SourceField *string `json:"sourceField,omitempty" xml:"sourceField,omitempty"`
	// The target path for writing, such as metadata.customerTier.
	//
	// example:
	//
	// metadata.customerTier
	Target *string `json:"target,omitempty" xml:"target,omitempty"`
	// The usage of the field. Valid values: extraction-input (participates in extraction) and ignore (ignored).
	//
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
	// The subset of the pipeline where clause, which is pushed down to SQL.
	//
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
	// The consumption mode. Valid values: follow (incremental consumption based on __dataset_seq) and fixed (full consumption of the specified ACTIVE version at once).
	//
	// example:
	//
	// follow
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The starting __dataset_seq for the follow mode.
	//
	// example:
	//
	// 0
	StartSeq *int64 `json:"startSeq,omitempty" xml:"startSeq,omitempty"`
	// The dataset version specified for the fixed mode.
	//
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
	// The trajectory filter condition. Multiple conditions are combined with AND, and conditions within a list are combined with OR.
	Filter *GetContextStoreResponseBodyConfigSourceTrajectoryFilter `json:"filter,omitempty" xml:"filter,omitempty" type:"Struct"`
	// The name of the trajectory Logstore. Default value: agent-trajectory. This parameter cannot be modified after creation.
	//
	// example:
	//
	// agent-trajectory
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The polling interval in seconds. Valid values: 60 to 3600. Default value: 300.
	//
	// example:
	//
	// 300
	PollIntervalSeconds *int32 `json:"pollIntervalSeconds,omitempty" xml:"pollIntervalSeconds,omitempty"`
	// The scope field mapping. The value is a JSONPath expression. Only the $.a.b and $.a[0] formats are supported.
	ScopeMapping *GetContextStoreResponseBodyConfigSourceTrajectoryScopeMapping `json:"scopeMapping,omitempty" xml:"scopeMapping,omitempty" type:"Struct"`
	// The start time for trajectory consumption in ISO 8601 UTC format. If not specified, the default value is the creation time.
	//
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
	// The list of agent names. An explicit empty list returns a 400 error.
	//
	// example:
	//
	// ["sales-copilot"]
	AgentNames []*string `json:"agentNames,omitempty" xml:"agentNames,omitempty" type:"Repeated"`
	// Specifies whether to exclude degraded trajectories. Default value: false (includes degraded trajectories).
	//
	// example:
	//
	// false
	ExcludeDegraded *bool `json:"excludeDegraded,omitempty" xml:"excludeDegraded,omitempty"`
	// The minimum step count, which must be greater than or equal to 0.
	//
	// example:
	//
	// 2
	MinStepCount *int32 `json:"minStepCount,omitempty" xml:"minStepCount,omitempty"`
	// The native Simple Log Service (SLS) query statement. This statement is combined with the preceding conditions using the AND operator.
	//
	// example:
	//
	// tool_names:"search_order"
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The list of service names. A single trailing asterisk (*) is supported. An explicit empty list returns a 400 error.
	//
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
	// The JSONPath expression for agentId. Default value: $.agent_name.
	//
	// example:
	//
	// $.agent_name
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The JSONPath expression for appId. Default value: $.service_names[0].
	//
	// example:
	//
	// $.service_names[0]
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// The JSONPath expression for runId. Default value: $.trajectory_id.
	//
	// example:
	//
	// $.trajectory_id
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The JSONPath expression for userId. By default, this field is not mapped.
	//
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

type GetContextStoreResponseBodyConfigSourceVersionPolicy struct {
	// The continuous consumption mode. The value is fixed to follow.
	//
	// example:
	//
	// follow
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The initial consumed append sequence number. The value is a decimal string of a non-negative int64 integer. Default value: 0. The __dataset_seq field is used, and the business timestamp is not used as the consumption checkpoint.
	//
	// example:
	//
	// 0
	StartSeq *string `json:"startSeq,omitempty" xml:"startSeq,omitempty"`
}

func (s GetContextStoreResponseBodyConfigSourceVersionPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigSourceVersionPolicy) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigSourceVersionPolicy) GetMode() *string {
	return s.Mode
}

func (s *GetContextStoreResponseBodyConfigSourceVersionPolicy) GetStartSeq() *string {
	return s.StartSeq
}

func (s *GetContextStoreResponseBodyConfigSourceVersionPolicy) SetMode(v string) *GetContextStoreResponseBodyConfigSourceVersionPolicy {
	s.Mode = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceVersionPolicy) SetStartSeq(v string) *GetContextStoreResponseBodyConfigSourceVersionPolicy {
	s.StartSeq = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceVersionPolicy) Validate() error {
	return dara.Validate(s)
}

type GetContextStoreResponseBodyConfigSourceStatus struct {
	// The processing progress. For memory trajectory, the format is {processTimeNs, watermarkMs}. For dataset, the format is {seq}. For manual, the format is {}. For experience Dataset follow, the seq field uses an exact int64 decimal string.
	Checkpoint map[string]interface{} `json:"checkpoint,omitempty" xml:"checkpoint,omitempty"`
	// The current reserved sequence window for experience Dataset follow, in the format of {after, until, runId}. All sequence numbers are decimal strings. If a failure occurs, the submitted checkpoint is not advanced.
	CurrentWindow map[string]*string `json:"currentWindow,omitempty" xml:"currentWindow,omitempty"`
	// The last desensitized error message. This field is empty when the status is normal.
	//
	// example:
	//
	// Data source read timeout
	LastError *string `json:"lastError,omitempty" xml:"lastError,omitempty"`
	// The run ID of the last successful automatic application to the native experience library. This field is for audit purposes only.
	//
	// example:
	//
	// opr-example
	LastRunId *string `json:"lastRunId,omitempty" xml:"lastRunId,omitempty"`
	// The last window processing time in ISO 8601 UTC format.
	//
	// example:
	//
	// 2026-10-01T08:00:00Z
	LastWindowAt *string `json:"lastWindowAt,omitempty" xml:"lastWindowAt,omitempty"`
	// The number of consecutive retries.
	//
	// example:
	//
	// 0
	RetryCount *int32 `json:"retryCount,omitempty" xml:"retryCount,omitempty"`
	// The state of the data source. Valid values: Validating, Ready, Running, Retrying, Paused, SourceSchemaChanged, and Failed.
	//
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

func (s *GetContextStoreResponseBodyConfigSourceStatus) GetCurrentWindow() map[string]*string {
	return s.CurrentWindow
}

func (s *GetContextStoreResponseBodyConfigSourceStatus) GetLastError() *string {
	return s.LastError
}

func (s *GetContextStoreResponseBodyConfigSourceStatus) GetLastRunId() *string {
	return s.LastRunId
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

func (s *GetContextStoreResponseBodyConfigSourceStatus) SetCurrentWindow(v map[string]*string) *GetContextStoreResponseBodyConfigSourceStatus {
	s.CurrentWindow = v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceStatus) SetLastError(v string) *GetContextStoreResponseBodyConfigSourceStatus {
	s.LastError = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSourceStatus) SetLastRunId(v string) *GetContextStoreResponseBodyConfigSourceStatus {
	s.LastRunId = &v
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
	// The allowed storage actions. Default value: all actions. Valid values: ADD, UPDATE, MERGE, and DELETE.
	//
	// example:
	//
	// ["ADD","UPDATE","MERGE","DELETE"]
	AllowedActions []*string `json:"allowedActions,omitempty" xml:"allowedActions,omitempty" type:"Repeated"`
	// Specifies whether to enable deduplication in event mode. Default value: true.
	//
	// example:
	//
	// true
	Dedupe *bool `json:"dedupe,omitempty" xml:"dedupe,omitempty"`
	// Specifies whether to enable human edit protection. Default value: true. When enabled, automatic extraction does not overwrite manually modified memories.
	//
	// example:
	//
	// true
	HumanEditProtection *bool `json:"humanEditProtection,omitempty" xml:"humanEditProtection,omitempty"`
	// The merge key for upsert operations. Default value: semantic. Valid values: factKey, semantic, and both.
	//
	// example:
	//
	// semantic
	MergeKey *string `json:"mergeKey,omitempty" xml:"mergeKey,omitempty"`
	// The storage mode. Default value: upsert. Valid values: event (event append) and upsert (merge update).
	//
	// example:
	//
	// upsert
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The similarity threshold for semantic merging. Valid values: 0 to 1. Default value: 0.4.
	//
	// example:
	//
	// 0.4
	SimilarityThreshold *float64 `json:"similarityThreshold,omitempty" xml:"similarityThreshold,omitempty"`
	// The number of days before memories expire. A value of 0 indicates that memories do not expire.
	//
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
