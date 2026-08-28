// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiStatisticsConfig interface {
	dara.Model
	String() string
	GoString() string
	SetLogRequestContent(v bool) *AiStatisticsConfig
	GetLogRequestContent() *bool
	SetLogResponseContent(v bool) *AiStatisticsConfig
	GetLogResponseContent() *bool
	SetPathFieldConfigs(v []*AiStatisticsConfigPathFieldConfigs) *AiStatisticsConfig
	GetPathFieldConfigs() []*AiStatisticsConfigPathFieldConfigs
}

type AiStatisticsConfig struct {
	// **[Deprecated]*	- Specifies whether to record request content (controls whether question-related attributes are generated). This parameter is deprecated in the new version.
	//
	// example:
	//
	// true
	LogRequestContent *bool `json:"logRequestContent,omitempty" xml:"logRequestContent,omitempty"`
	// **[Deprecated]*	- Specifies whether to record response content (controls whether answer-related attributes are generated). This parameter is deprecated in the new version.
	//
	// example:
	//
	// true
	LogResponseContent *bool `json:"logResponseContent,omitempty" xml:"logResponseContent,omitempty"`
	// The list of AI request log field collection configurations, configured by API path.
	PathFieldConfigs []*AiStatisticsConfigPathFieldConfigs `json:"pathFieldConfigs,omitempty" xml:"pathFieldConfigs,omitempty" type:"Repeated"`
}

func (s AiStatisticsConfig) String() string {
	return dara.Prettify(s)
}

func (s AiStatisticsConfig) GoString() string {
	return s.String()
}

func (s *AiStatisticsConfig) GetLogRequestContent() *bool {
	return s.LogRequestContent
}

func (s *AiStatisticsConfig) GetLogResponseContent() *bool {
	return s.LogResponseContent
}

func (s *AiStatisticsConfig) GetPathFieldConfigs() []*AiStatisticsConfigPathFieldConfigs {
	return s.PathFieldConfigs
}

func (s *AiStatisticsConfig) SetLogRequestContent(v bool) *AiStatisticsConfig {
	s.LogRequestContent = &v
	return s
}

func (s *AiStatisticsConfig) SetLogResponseContent(v bool) *AiStatisticsConfig {
	s.LogResponseContent = &v
	return s
}

func (s *AiStatisticsConfig) SetPathFieldConfigs(v []*AiStatisticsConfigPathFieldConfigs) *AiStatisticsConfig {
	s.PathFieldConfigs = v
	return s
}

func (s *AiStatisticsConfig) Validate() error {
	if s.PathFieldConfigs != nil {
		for _, item := range s.PathFieldConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AiStatisticsConfigPathFieldConfigs struct {
	// The AI request log field configuration groups for the API path, passed in as a Map. The Map keys are fixed to basic and custom, and the values are arrays of log field configurations for the corresponding groups. basic indicates basic log fields, and custom indicates custom log fields. For the current API path, fieldPaths represents the complete desired state of field configurations and does not support incremental appending or diff merging.
	//
	// If pathFieldConfigs is not passed, is null, or is an empty array, the existing log field configurations are not updated. If a non-empty array is passed, the system performs a desired state replacement based on the complete set of Paths in the request, and historical Path configurations not included in the request are deleted.
	//
	// For example, to add a custom field test to the /v1/chat/completions API path on top of existing configurations, the caller must use a "read-merge-write back in full" approach:
	//
	// 1. Read all current Path configurations.
	//
	// 2. Retain the complete basic array and custom array for the target API path /v1/chat/completions.
	//
	// 3. Append test to the current custom array.
	//
	// 4. Keep configurations for other API paths unchanged.
	//
	// 5. Submit the merged complete pathFieldConfigs.
	FieldPaths map[string]*AiStatisticsPathField `json:"fieldPaths,omitempty" xml:"fieldPaths,omitempty"`
	// The API path.
	//
	// example:
	//
	// /v1/chat/completions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s AiStatisticsConfigPathFieldConfigs) String() string {
	return dara.Prettify(s)
}

func (s AiStatisticsConfigPathFieldConfigs) GoString() string {
	return s.String()
}

func (s *AiStatisticsConfigPathFieldConfigs) GetFieldPaths() map[string]*AiStatisticsPathField {
	return s.FieldPaths
}

func (s *AiStatisticsConfigPathFieldConfigs) GetPath() *string {
	return s.Path
}

func (s *AiStatisticsConfigPathFieldConfigs) SetFieldPaths(v map[string]*AiStatisticsPathField) *AiStatisticsConfigPathFieldConfigs {
	s.FieldPaths = v
	return s
}

func (s *AiStatisticsConfigPathFieldConfigs) SetPath(v string) *AiStatisticsConfigPathFieldConfigs {
	s.Path = &v
	return s
}

func (s *AiStatisticsConfigPathFieldConfigs) Validate() error {
	return dara.Validate(s)
}
