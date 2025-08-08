// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDefaultBuilderConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCache(v *BuildCacheConfig) *DefaultBuilderConfig
	GetCache() *BuildCacheConfig
	SetLanguages(v []*string) *DefaultBuilderConfig
	GetLanguages() []*string
	SetRootPath(v string) *DefaultBuilderConfig
	GetRootPath() *string
	SetSteps(v []interface{}) *DefaultBuilderConfig
	GetSteps() []interface{}
}

type DefaultBuilderConfig struct {
	Cache     *BuildCacheConfig `json:"cache,omitempty" xml:"cache,omitempty"`
	Languages []*string         `json:"languages,omitempty" xml:"languages,omitempty" type:"Repeated"`
	// example:
	//
	// ./src
	RootPath *string       `json:"rootPath,omitempty" xml:"rootPath,omitempty"`
	Steps    []interface{} `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
}

func (s DefaultBuilderConfig) String() string {
	return dara.Prettify(s)
}

func (s DefaultBuilderConfig) GoString() string {
	return s.String()
}

func (s *DefaultBuilderConfig) GetCache() *BuildCacheConfig {
	return s.Cache
}

func (s *DefaultBuilderConfig) GetLanguages() []*string {
	return s.Languages
}

func (s *DefaultBuilderConfig) GetRootPath() *string {
	return s.RootPath
}

func (s *DefaultBuilderConfig) GetSteps() []interface{} {
	return s.Steps
}

func (s *DefaultBuilderConfig) SetCache(v *BuildCacheConfig) *DefaultBuilderConfig {
	s.Cache = v
	return s
}

func (s *DefaultBuilderConfig) SetLanguages(v []*string) *DefaultBuilderConfig {
	s.Languages = v
	return s
}

func (s *DefaultBuilderConfig) SetRootPath(v string) *DefaultBuilderConfig {
	s.RootPath = &v
	return s
}

func (s *DefaultBuilderConfig) SetSteps(v []interface{}) *DefaultBuilderConfig {
	s.Steps = v
	return s
}

func (s *DefaultBuilderConfig) Validate() error {
	return dara.Validate(s)
}
