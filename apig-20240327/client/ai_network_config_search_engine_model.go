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
	// The API key of the search engine.
	//
	// example:
	//
	// sk-xxx
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// The content mode. This parameter is specific to Quark.
	//
	// example:
	//
	// summary
	ContentMode *string `json:"contentMode,omitempty" xml:"contentMode,omitempty"`
	// The number of results returned per search.
	//
	// example:
	//
	// 5
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The endpoint of the search engine.
	//
	// example:
	//
	// https://cloud-iqs.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The industry filter. This parameter is specific to Quark.
	//
	// example:
	//
	// Internet
	Industry *string `json:"industry,omitempty" xml:"industry,omitempty"`
	// The search engine-specific parameters in key-value pair format.
	OptionArgs map[string]*string `json:"optionArgs,omitempty" xml:"optionArgs,omitempty"`
	// The offset of the search results.
	//
	// example:
	//
	// 0
	Start *int32 `json:"start,omitempty" xml:"start,omitempty"`
	// The time range filter. This parameter is specific to Quark.
	//
	// example:
	//
	// 7d
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
	// The timeout period for the API call, in milliseconds.
	//
	// example:
	//
	// 5000
	TimeoutMillisecond *int32 `json:"timeoutMillisecond,omitempty" xml:"timeoutMillisecond,omitempty"`
	// The search engine type.
	//
	// example:
	//
	// aliyunQuark
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
