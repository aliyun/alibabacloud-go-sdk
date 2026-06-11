// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetMemoryStoreResponseBody
	GetCreateTime() *string
	SetCustomExtractionStrategies(v []*CustomExtractionStrategy) *GetMemoryStoreResponseBody
	GetCustomExtractionStrategies() []*CustomExtractionStrategy
	SetDescription(v string) *GetMemoryStoreResponseBody
	GetDescription() *string
	SetExtractionStrategies(v []*string) *GetMemoryStoreResponseBody
	GetExtractionStrategies() []*string
	SetMemoryStoreName(v string) *GetMemoryStoreResponseBody
	GetMemoryStoreName() *string
	SetRegionId(v string) *GetMemoryStoreResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetMemoryStoreResponseBody
	GetRequestId() *string
	SetShortTermStorage(v *GetMemoryStoreResponseBodyShortTermStorage) *GetMemoryStoreResponseBody
	GetShortTermStorage() *GetMemoryStoreResponseBodyShortTermStorage
	SetShortTermTtl(v int32) *GetMemoryStoreResponseBody
	GetShortTermTtl() *int32
	SetSourceType(v string) *GetMemoryStoreResponseBody
	GetSourceType() *string
	SetTraceSourceConfig(v *GetMemoryStoreResponseBodyTraceSourceConfig) *GetMemoryStoreResponseBody
	GetTraceSourceConfig() *GetMemoryStoreResponseBodyTraceSourceConfig
	SetUpdateTime(v string) *GetMemoryStoreResponseBody
	GetUpdateTime() *string
	SetWorkspace(v string) *GetMemoryStoreResponseBody
	GetWorkspace() *string
}

type GetMemoryStoreResponseBody struct {
	// The Unix timestamp (in milliseconds) when the memory store was created.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 1764556182850
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// A list of custom extraction strategies.
	CustomExtractionStrategies []*CustomExtractionStrategy `json:"customExtractionStrategies,omitempty" xml:"customExtractionStrategies,omitempty" type:"Repeated"`
	// The description of the memory store.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The built-in extraction strategies. Valid values are `Episodic`, `Summary`, and `Fact`.
	ExtractionStrategies []*string `json:"extractionStrategies,omitempty" xml:"extractionStrategies,omitempty" type:"Repeated"`
	// The name of the memory store.
	//
	// example:
	//
	// test-memory-store
	MemoryStoreName *string `json:"memoryStoreName,omitempty" xml:"memoryStoreName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-A01D6CC3F4B8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The configuration for short-term memory storage.
	ShortTermStorage *GetMemoryStoreResponseBodyShortTermStorage `json:"shortTermStorage,omitempty" xml:"shortTermStorage,omitempty" type:"Struct"`
	// The short-term memory retention time, in seconds.
	//
	// example:
	//
	// 10
	ShortTermTtl *int32 `json:"shortTermTtl,omitempty" xml:"shortTermTtl,omitempty"`
	// The memory source.
	//
	// example:
	//
	// Trace
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The configuration for the trace source. This parameter is returned only when `sourceType` is set to `Trace`.
	TraceSourceConfig *GetMemoryStoreResponseBodyTraceSourceConfig `json:"traceSourceConfig,omitempty" xml:"traceSourceConfig,omitempty" type:"Struct"`
	// The Unix timestamp (in milliseconds) when the memory store was last updated.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 1764556182850
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// default-cms-xxxxxx-cn-beijing
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetMemoryStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryStoreResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryStoreResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMemoryStoreResponseBody) GetCustomExtractionStrategies() []*CustomExtractionStrategy {
	return s.CustomExtractionStrategies
}

func (s *GetMemoryStoreResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetMemoryStoreResponseBody) GetExtractionStrategies() []*string {
	return s.ExtractionStrategies
}

func (s *GetMemoryStoreResponseBody) GetMemoryStoreName() *string {
	return s.MemoryStoreName
}

func (s *GetMemoryStoreResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetMemoryStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemoryStoreResponseBody) GetShortTermStorage() *GetMemoryStoreResponseBodyShortTermStorage {
	return s.ShortTermStorage
}

func (s *GetMemoryStoreResponseBody) GetShortTermTtl() *int32 {
	return s.ShortTermTtl
}

func (s *GetMemoryStoreResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *GetMemoryStoreResponseBody) GetTraceSourceConfig() *GetMemoryStoreResponseBodyTraceSourceConfig {
	return s.TraceSourceConfig
}

func (s *GetMemoryStoreResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetMemoryStoreResponseBody) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetMemoryStoreResponseBody) SetCreateTime(v string) *GetMemoryStoreResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetMemoryStoreResponseBody) SetCustomExtractionStrategies(v []*CustomExtractionStrategy) *GetMemoryStoreResponseBody {
	s.CustomExtractionStrategies = v
	return s
}

func (s *GetMemoryStoreResponseBody) SetDescription(v string) *GetMemoryStoreResponseBody {
	s.Description = &v
	return s
}

func (s *GetMemoryStoreResponseBody) SetExtractionStrategies(v []*string) *GetMemoryStoreResponseBody {
	s.ExtractionStrategies = v
	return s
}

func (s *GetMemoryStoreResponseBody) SetMemoryStoreName(v string) *GetMemoryStoreResponseBody {
	s.MemoryStoreName = &v
	return s
}

func (s *GetMemoryStoreResponseBody) SetRegionId(v string) *GetMemoryStoreResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetMemoryStoreResponseBody) SetRequestId(v string) *GetMemoryStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryStoreResponseBody) SetShortTermStorage(v *GetMemoryStoreResponseBodyShortTermStorage) *GetMemoryStoreResponseBody {
	s.ShortTermStorage = v
	return s
}

func (s *GetMemoryStoreResponseBody) SetShortTermTtl(v int32) *GetMemoryStoreResponseBody {
	s.ShortTermTtl = &v
	return s
}

func (s *GetMemoryStoreResponseBody) SetSourceType(v string) *GetMemoryStoreResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetMemoryStoreResponseBody) SetTraceSourceConfig(v *GetMemoryStoreResponseBodyTraceSourceConfig) *GetMemoryStoreResponseBody {
	s.TraceSourceConfig = v
	return s
}

func (s *GetMemoryStoreResponseBody) SetUpdateTime(v string) *GetMemoryStoreResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetMemoryStoreResponseBody) SetWorkspace(v string) *GetMemoryStoreResponseBody {
	s.Workspace = &v
	return s
}

func (s *GetMemoryStoreResponseBody) Validate() error {
	if s.CustomExtractionStrategies != nil {
		for _, item := range s.CustomExtractionStrategies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ShortTermStorage != nil {
		if err := s.ShortTermStorage.Validate(); err != nil {
			return err
		}
	}
	if s.TraceSourceConfig != nil {
		if err := s.TraceSourceConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMemoryStoreResponseBodyShortTermStorage struct {
	// The name of the Simple Log Service Logstore.
	//
	// example:
	//
	// memory-store
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The name of the Simple Log Service Project.
	//
	// example:
	//
	// wk_cms_data_warehouse
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s GetMemoryStoreResponseBodyShortTermStorage) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryStoreResponseBodyShortTermStorage) GoString() string {
	return s.String()
}

func (s *GetMemoryStoreResponseBodyShortTermStorage) GetLogstore() *string {
	return s.Logstore
}

func (s *GetMemoryStoreResponseBodyShortTermStorage) GetProject() *string {
	return s.Project
}

func (s *GetMemoryStoreResponseBodyShortTermStorage) SetLogstore(v string) *GetMemoryStoreResponseBodyShortTermStorage {
	s.Logstore = &v
	return s
}

func (s *GetMemoryStoreResponseBodyShortTermStorage) SetProject(v string) *GetMemoryStoreResponseBodyShortTermStorage {
	s.Project = &v
	return s
}

func (s *GetMemoryStoreResponseBodyShortTermStorage) Validate() error {
	return dara.Validate(s)
}

type GetMemoryStoreResponseBodyTraceSourceConfig struct {
	// Indicates whether to include Large Language Model (LLM) output in the memory extraction.
	IncludeOutput *bool `json:"includeOutput,omitempty" xml:"includeOutput,omitempty"`
	// The query to filter traces from Simple Log Service.
	//
	// example:
	//
	// (serviceName : "langchain-rag" or serviceName : "agentscope-code-correction") and hostname = frontend-proxy-999c48c8d-hvk6c
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The workspace where the trace is located.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetMemoryStoreResponseBodyTraceSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryStoreResponseBodyTraceSourceConfig) GoString() string {
	return s.String()
}

func (s *GetMemoryStoreResponseBodyTraceSourceConfig) GetIncludeOutput() *bool {
	return s.IncludeOutput
}

func (s *GetMemoryStoreResponseBodyTraceSourceConfig) GetQuery() *string {
	return s.Query
}

func (s *GetMemoryStoreResponseBodyTraceSourceConfig) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetMemoryStoreResponseBodyTraceSourceConfig) SetIncludeOutput(v bool) *GetMemoryStoreResponseBodyTraceSourceConfig {
	s.IncludeOutput = &v
	return s
}

func (s *GetMemoryStoreResponseBodyTraceSourceConfig) SetQuery(v string) *GetMemoryStoreResponseBodyTraceSourceConfig {
	s.Query = &v
	return s
}

func (s *GetMemoryStoreResponseBodyTraceSourceConfig) SetWorkspace(v string) *GetMemoryStoreResponseBodyTraceSourceConfig {
	s.Workspace = &v
	return s
}

func (s *GetMemoryStoreResponseBodyTraceSourceConfig) Validate() error {
	return dara.Validate(s)
}
