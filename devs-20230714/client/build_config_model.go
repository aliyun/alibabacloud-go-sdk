// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildConfig interface {
	dara.Model
	String() string
	GoString() string
	SetDefault(v *DefaultBuilderConfig) *BuildConfig
	GetDefault() *DefaultBuilderConfig
}

type BuildConfig struct {
	Default *DefaultBuilderConfig `json:"default,omitempty" xml:"default,omitempty"`
}

func (s BuildConfig) String() string {
	return dara.Prettify(s)
}

func (s BuildConfig) GoString() string {
	return s.String()
}

func (s *BuildConfig) GetDefault() *DefaultBuilderConfig {
	return s.Default
}

func (s *BuildConfig) SetDefault(v *DefaultBuilderConfig) *BuildConfig {
	s.Default = v
	return s
}

func (s *BuildConfig) Validate() error {
	return dara.Validate(s)
}
