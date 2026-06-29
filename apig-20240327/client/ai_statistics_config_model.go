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
	// Specifies whether to log request content (controls whether question-related attributes are generated).
	//
	// example:
	//
	// true
	LogRequestContent *bool `json:"logRequestContent,omitempty" xml:"logRequestContent,omitempty"`
	// Specifies whether to log response content (controls whether answer-related attributes are generated).
	//
	// example:
	//
	// true
	LogResponseContent *bool `json:"logResponseContent,omitempty" xml:"logResponseContent,omitempty"`
	// The list of custom field collection configurations, configured by API path.
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
	// The field collection configuration.
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
