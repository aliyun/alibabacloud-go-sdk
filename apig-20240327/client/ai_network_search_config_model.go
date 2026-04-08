// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiNetworkSearchConfig interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultEnable(v bool) *AiNetworkSearchConfig
	GetDefaultEnable() *bool
	SetDefaultLang(v string) *AiNetworkSearchConfig
	GetDefaultLang() *string
	SetNeedReference(v bool) *AiNetworkSearchConfig
	GetNeedReference() *bool
	SetPluginStatus(v *AiPluginStatus) *AiNetworkSearchConfig
	GetPluginStatus() *AiPluginStatus
	SetReferenceFormat(v string) *AiNetworkSearchConfig
	GetReferenceFormat() *string
	SetReferenceLocation(v string) *AiNetworkSearchConfig
	GetReferenceLocation() *string
	SetSearchEngineConfig(v *AiNetworkConfigSearchEngine) *AiNetworkSearchConfig
	GetSearchEngineConfig() *AiNetworkConfigSearchEngine
	SetSearchFrom(v []*AiNetworkConfigSearchEngine) *AiNetworkSearchConfig
	GetSearchFrom() []*AiNetworkConfigSearchEngine
	SetSearchRewrite(v *AiNetworkSearchConfigSearchRewrite) *AiNetworkSearchConfig
	GetSearchRewrite() *AiNetworkSearchConfigSearchRewrite
}

type AiNetworkSearchConfig struct {
	DefaultEnable *bool   `json:"defaultEnable,omitempty" xml:"defaultEnable,omitempty"`
	DefaultLang   *string `json:"defaultLang,omitempty" xml:"defaultLang,omitempty"`
	NeedReference *bool   `json:"needReference,omitempty" xml:"needReference,omitempty"`
	// if can be null:
	// true
	PluginStatus       *AiPluginStatus                     `json:"pluginStatus,omitempty" xml:"pluginStatus,omitempty"`
	ReferenceFormat    *string                             `json:"referenceFormat,omitempty" xml:"referenceFormat,omitempty"`
	ReferenceLocation  *string                             `json:"referenceLocation,omitempty" xml:"referenceLocation,omitempty"`
	SearchEngineConfig *AiNetworkConfigSearchEngine        `json:"searchEngineConfig,omitempty" xml:"searchEngineConfig,omitempty"`
	SearchFrom         []*AiNetworkConfigSearchEngine      `json:"searchFrom,omitempty" xml:"searchFrom,omitempty" type:"Repeated"`
	SearchRewrite      *AiNetworkSearchConfigSearchRewrite `json:"searchRewrite,omitempty" xml:"searchRewrite,omitempty" type:"Struct"`
}

func (s AiNetworkSearchConfig) String() string {
	return dara.Prettify(s)
}

func (s AiNetworkSearchConfig) GoString() string {
	return s.String()
}

func (s *AiNetworkSearchConfig) GetDefaultEnable() *bool {
	return s.DefaultEnable
}

func (s *AiNetworkSearchConfig) GetDefaultLang() *string {
	return s.DefaultLang
}

func (s *AiNetworkSearchConfig) GetNeedReference() *bool {
	return s.NeedReference
}

func (s *AiNetworkSearchConfig) GetPluginStatus() *AiPluginStatus {
	return s.PluginStatus
}

func (s *AiNetworkSearchConfig) GetReferenceFormat() *string {
	return s.ReferenceFormat
}

func (s *AiNetworkSearchConfig) GetReferenceLocation() *string {
	return s.ReferenceLocation
}

func (s *AiNetworkSearchConfig) GetSearchEngineConfig() *AiNetworkConfigSearchEngine {
	return s.SearchEngineConfig
}

func (s *AiNetworkSearchConfig) GetSearchFrom() []*AiNetworkConfigSearchEngine {
	return s.SearchFrom
}

func (s *AiNetworkSearchConfig) GetSearchRewrite() *AiNetworkSearchConfigSearchRewrite {
	return s.SearchRewrite
}

func (s *AiNetworkSearchConfig) SetDefaultEnable(v bool) *AiNetworkSearchConfig {
	s.DefaultEnable = &v
	return s
}

func (s *AiNetworkSearchConfig) SetDefaultLang(v string) *AiNetworkSearchConfig {
	s.DefaultLang = &v
	return s
}

func (s *AiNetworkSearchConfig) SetNeedReference(v bool) *AiNetworkSearchConfig {
	s.NeedReference = &v
	return s
}

func (s *AiNetworkSearchConfig) SetPluginStatus(v *AiPluginStatus) *AiNetworkSearchConfig {
	s.PluginStatus = v
	return s
}

func (s *AiNetworkSearchConfig) SetReferenceFormat(v string) *AiNetworkSearchConfig {
	s.ReferenceFormat = &v
	return s
}

func (s *AiNetworkSearchConfig) SetReferenceLocation(v string) *AiNetworkSearchConfig {
	s.ReferenceLocation = &v
	return s
}

func (s *AiNetworkSearchConfig) SetSearchEngineConfig(v *AiNetworkConfigSearchEngine) *AiNetworkSearchConfig {
	s.SearchEngineConfig = v
	return s
}

func (s *AiNetworkSearchConfig) SetSearchFrom(v []*AiNetworkConfigSearchEngine) *AiNetworkSearchConfig {
	s.SearchFrom = v
	return s
}

func (s *AiNetworkSearchConfig) SetSearchRewrite(v *AiNetworkSearchConfigSearchRewrite) *AiNetworkSearchConfig {
	s.SearchRewrite = v
	return s
}

func (s *AiNetworkSearchConfig) Validate() error {
	if s.PluginStatus != nil {
		if err := s.PluginStatus.Validate(); err != nil {
			return err
		}
	}
	if s.SearchEngineConfig != nil {
		if err := s.SearchEngineConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SearchFrom != nil {
		for _, item := range s.SearchFrom {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SearchRewrite != nil {
		if err := s.SearchRewrite.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AiNetworkSearchConfigSearchRewrite struct {
	Enable             *bool   `json:"enable,omitempty" xml:"enable,omitempty"`
	MaxCount           *int32  `json:"maxCount,omitempty" xml:"maxCount,omitempty"`
	ModelName          *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ServiceId          *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	TimeoutMillisecond *int32  `json:"timeoutMillisecond,omitempty" xml:"timeoutMillisecond,omitempty"`
}

func (s AiNetworkSearchConfigSearchRewrite) String() string {
	return dara.Prettify(s)
}

func (s AiNetworkSearchConfigSearchRewrite) GoString() string {
	return s.String()
}

func (s *AiNetworkSearchConfigSearchRewrite) GetEnable() *bool {
	return s.Enable
}

func (s *AiNetworkSearchConfigSearchRewrite) GetMaxCount() *int32 {
	return s.MaxCount
}

func (s *AiNetworkSearchConfigSearchRewrite) GetModelName() *string {
	return s.ModelName
}

func (s *AiNetworkSearchConfigSearchRewrite) GetServiceId() *string {
	return s.ServiceId
}

func (s *AiNetworkSearchConfigSearchRewrite) GetTimeoutMillisecond() *int32 {
	return s.TimeoutMillisecond
}

func (s *AiNetworkSearchConfigSearchRewrite) SetEnable(v bool) *AiNetworkSearchConfigSearchRewrite {
	s.Enable = &v
	return s
}

func (s *AiNetworkSearchConfigSearchRewrite) SetMaxCount(v int32) *AiNetworkSearchConfigSearchRewrite {
	s.MaxCount = &v
	return s
}

func (s *AiNetworkSearchConfigSearchRewrite) SetModelName(v string) *AiNetworkSearchConfigSearchRewrite {
	s.ModelName = &v
	return s
}

func (s *AiNetworkSearchConfigSearchRewrite) SetServiceId(v string) *AiNetworkSearchConfigSearchRewrite {
	s.ServiceId = &v
	return s
}

func (s *AiNetworkSearchConfigSearchRewrite) SetTimeoutMillisecond(v int32) *AiNetworkSearchConfigSearchRewrite {
	s.TimeoutMillisecond = &v
	return s
}

func (s *AiNetworkSearchConfigSearchRewrite) Validate() error {
	return dara.Validate(s)
}
