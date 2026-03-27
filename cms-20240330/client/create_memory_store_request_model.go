// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryStoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomExtractionStrategies(v []*CustomExtractionStrategy) *CreateMemoryStoreRequest
	GetCustomExtractionStrategies() []*CustomExtractionStrategy
	SetDescription(v string) *CreateMemoryStoreRequest
	GetDescription() *string
	SetExtractionStrategies(v []*string) *CreateMemoryStoreRequest
	GetExtractionStrategies() []*string
	SetMemoryStoreName(v string) *CreateMemoryStoreRequest
	GetMemoryStoreName() *string
	SetShortTermTtl(v int32) *CreateMemoryStoreRequest
	GetShortTermTtl() *int32
	SetSourceType(v string) *CreateMemoryStoreRequest
	GetSourceType() *string
	SetTraceSourceConfig(v *CreateMemoryStoreRequestTraceSourceConfig) *CreateMemoryStoreRequest
	GetTraceSourceConfig() *CreateMemoryStoreRequestTraceSourceConfig
}

type CreateMemoryStoreRequest struct {
	CustomExtractionStrategies []*CustomExtractionStrategy `json:"customExtractionStrategies,omitempty" xml:"customExtractionStrategies,omitempty" type:"Repeated"`
	// example:
	//
	// Test memory store for demonstration.
	Description          *string   `json:"description,omitempty" xml:"description,omitempty"`
	ExtractionStrategies []*string `json:"extractionStrategies,omitempty" xml:"extractionStrategies,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test-memory-store
	MemoryStoreName *string `json:"memoryStoreName,omitempty" xml:"memoryStoreName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	ShortTermTtl *int32 `json:"shortTermTtl,omitempty" xml:"shortTermTtl,omitempty"`
	// example:
	//
	// None/Trace
	SourceType        *string                                    `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	TraceSourceConfig *CreateMemoryStoreRequestTraceSourceConfig `json:"traceSourceConfig,omitempty" xml:"traceSourceConfig,omitempty" type:"Struct"`
}

func (s CreateMemoryStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryStoreRequest) GoString() string {
	return s.String()
}

func (s *CreateMemoryStoreRequest) GetCustomExtractionStrategies() []*CustomExtractionStrategy {
	return s.CustomExtractionStrategies
}

func (s *CreateMemoryStoreRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateMemoryStoreRequest) GetExtractionStrategies() []*string {
	return s.ExtractionStrategies
}

func (s *CreateMemoryStoreRequest) GetMemoryStoreName() *string {
	return s.MemoryStoreName
}

func (s *CreateMemoryStoreRequest) GetShortTermTtl() *int32 {
	return s.ShortTermTtl
}

func (s *CreateMemoryStoreRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateMemoryStoreRequest) GetTraceSourceConfig() *CreateMemoryStoreRequestTraceSourceConfig {
	return s.TraceSourceConfig
}

func (s *CreateMemoryStoreRequest) SetCustomExtractionStrategies(v []*CustomExtractionStrategy) *CreateMemoryStoreRequest {
	s.CustomExtractionStrategies = v
	return s
}

func (s *CreateMemoryStoreRequest) SetDescription(v string) *CreateMemoryStoreRequest {
	s.Description = &v
	return s
}

func (s *CreateMemoryStoreRequest) SetExtractionStrategies(v []*string) *CreateMemoryStoreRequest {
	s.ExtractionStrategies = v
	return s
}

func (s *CreateMemoryStoreRequest) SetMemoryStoreName(v string) *CreateMemoryStoreRequest {
	s.MemoryStoreName = &v
	return s
}

func (s *CreateMemoryStoreRequest) SetShortTermTtl(v int32) *CreateMemoryStoreRequest {
	s.ShortTermTtl = &v
	return s
}

func (s *CreateMemoryStoreRequest) SetSourceType(v string) *CreateMemoryStoreRequest {
	s.SourceType = &v
	return s
}

func (s *CreateMemoryStoreRequest) SetTraceSourceConfig(v *CreateMemoryStoreRequestTraceSourceConfig) *CreateMemoryStoreRequest {
	s.TraceSourceConfig = v
	return s
}

func (s *CreateMemoryStoreRequest) Validate() error {
	if s.CustomExtractionStrategies != nil {
		for _, item := range s.CustomExtractionStrategies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TraceSourceConfig != nil {
		if err := s.TraceSourceConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMemoryStoreRequestTraceSourceConfig struct {
	IncludeOutput *bool `json:"includeOutput,omitempty" xml:"includeOutput,omitempty"`
	// example:
	//
	// (serviceName : "langchain-rag" or serviceName : "agentscope-code-correction") and hostname = frontend-proxy-999c48c8d-hvk6c
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// test-workspace
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateMemoryStoreRequestTraceSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryStoreRequestTraceSourceConfig) GoString() string {
	return s.String()
}

func (s *CreateMemoryStoreRequestTraceSourceConfig) GetIncludeOutput() *bool {
	return s.IncludeOutput
}

func (s *CreateMemoryStoreRequestTraceSourceConfig) GetQuery() *string {
	return s.Query
}

func (s *CreateMemoryStoreRequestTraceSourceConfig) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateMemoryStoreRequestTraceSourceConfig) SetIncludeOutput(v bool) *CreateMemoryStoreRequestTraceSourceConfig {
	s.IncludeOutput = &v
	return s
}

func (s *CreateMemoryStoreRequestTraceSourceConfig) SetQuery(v string) *CreateMemoryStoreRequestTraceSourceConfig {
	s.Query = &v
	return s
}

func (s *CreateMemoryStoreRequestTraceSourceConfig) SetWorkspace(v string) *CreateMemoryStoreRequestTraceSourceConfig {
	s.Workspace = &v
	return s
}

func (s *CreateMemoryStoreRequestTraceSourceConfig) Validate() error {
	return dara.Validate(s)
}
