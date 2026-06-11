// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiNetworkConfigSearchEngine interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *AiNetworkConfigSearchEngine
	GetApiKey() *string
	SetContentMode(v string) *AiNetworkConfigSearchEngine
	GetContentMode() *string
	SetCount(v int32) *AiNetworkConfigSearchEngine
	GetCount() *int32
	SetEndpoint(v string) *AiNetworkConfigSearchEngine
	GetEndpoint() *string
	SetIndustry(v string) *AiNetworkConfigSearchEngine
	GetIndustry() *string
	SetOptionArgs(v map[string]*string) *AiNetworkConfigSearchEngine
	GetOptionArgs() map[string]*string
	SetStart(v int32) *AiNetworkConfigSearchEngine
	GetStart() *int32
	SetTimeRange(v string) *AiNetworkConfigSearchEngine
	GetTimeRange() *string
	SetTimeoutMillisecond(v int32) *AiNetworkConfigSearchEngine
	GetTimeoutMillisecond() *int32
	SetType(v string) *AiNetworkConfigSearchEngine
	GetType() *string
}

type AiNetworkConfigSearchEngine struct {
	// Required. The API key to authenticate requests to the search engine service.
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// The search content mode, which determines how the service interprets the query and returns results.
	ContentMode *string `json:"contentMode,omitempty" xml:"contentMode,omitempty"`
	// The number of search results to return. If this parameter is omitted, the service uses a default value.
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The endpoint URL for the search engine service.
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The industry context for tailoring search results to a specific domain.
	Industry *string `json:"industry,omitempty" xml:"industry,omitempty"`
	// Optional. A key-value map for service-specific parameters not covered by the standard configuration.
	OptionArgs map[string]*string `json:"optionArgs,omitempty" xml:"optionArgs,omitempty"`
	// The starting offset for the search results, used for pagination. For example, a value of 10 skips the first 10 results. The default is 0.
	Start *int32 `json:"start,omitempty" xml:"start,omitempty"`
	// The time range for filtering results by their creation or modification date.
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
	// The request timeout in milliseconds. If a request exceeds this time, the service terminates it.
	TimeoutMillisecond *int32 `json:"timeoutMillisecond,omitempty" xml:"timeoutMillisecond,omitempty"`
	// Specifies the search engine service to use.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AiNetworkConfigSearchEngine) String() string {
	return dara.Prettify(s)
}

func (s AiNetworkConfigSearchEngine) GoString() string {
	return s.String()
}

func (s *AiNetworkConfigSearchEngine) GetApiKey() *string {
	return s.ApiKey
}

func (s *AiNetworkConfigSearchEngine) GetContentMode() *string {
	return s.ContentMode
}

func (s *AiNetworkConfigSearchEngine) GetCount() *int32 {
	return s.Count
}

func (s *AiNetworkConfigSearchEngine) GetEndpoint() *string {
	return s.Endpoint
}

func (s *AiNetworkConfigSearchEngine) GetIndustry() *string {
	return s.Industry
}

func (s *AiNetworkConfigSearchEngine) GetOptionArgs() map[string]*string {
	return s.OptionArgs
}

func (s *AiNetworkConfigSearchEngine) GetStart() *int32 {
	return s.Start
}

func (s *AiNetworkConfigSearchEngine) GetTimeRange() *string {
	return s.TimeRange
}

func (s *AiNetworkConfigSearchEngine) GetTimeoutMillisecond() *int32 {
	return s.TimeoutMillisecond
}

func (s *AiNetworkConfigSearchEngine) GetType() *string {
	return s.Type
}

func (s *AiNetworkConfigSearchEngine) SetApiKey(v string) *AiNetworkConfigSearchEngine {
	s.ApiKey = &v
	return s
}

func (s *AiNetworkConfigSearchEngine) SetContentMode(v string) *AiNetworkConfigSearchEngine {
	s.ContentMode = &v
	return s
}

func (s *AiNetworkConfigSearchEngine) SetCount(v int32) *AiNetworkConfigSearchEngine {
	s.Count = &v
	return s
}

func (s *AiNetworkConfigSearchEngine) SetEndpoint(v string) *AiNetworkConfigSearchEngine {
	s.Endpoint = &v
	return s
}

func (s *AiNetworkConfigSearchEngine) SetIndustry(v string) *AiNetworkConfigSearchEngine {
	s.Industry = &v
	return s
}

func (s *AiNetworkConfigSearchEngine) SetOptionArgs(v map[string]*string) *AiNetworkConfigSearchEngine {
	s.OptionArgs = v
	return s
}

func (s *AiNetworkConfigSearchEngine) SetStart(v int32) *AiNetworkConfigSearchEngine {
	s.Start = &v
	return s
}

func (s *AiNetworkConfigSearchEngine) SetTimeRange(v string) *AiNetworkConfigSearchEngine {
	s.TimeRange = &v
	return s
}

func (s *AiNetworkConfigSearchEngine) SetTimeoutMillisecond(v int32) *AiNetworkConfigSearchEngine {
	s.TimeoutMillisecond = &v
	return s
}

func (s *AiNetworkConfigSearchEngine) SetType(v string) *AiNetworkConfigSearchEngine {
	s.Type = &v
	return s
}

func (s *AiNetworkConfigSearchEngine) Validate() error {
	return dara.Validate(s)
}
