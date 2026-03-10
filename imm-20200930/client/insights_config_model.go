// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsightsConfig interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *InsightsConfig
	GetLanguage() *string
}

type InsightsConfig struct {
	// example:
	//
	// zh-Hans
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s InsightsConfig) String() string {
	return dara.Prettify(s)
}

func (s InsightsConfig) GoString() string {
	return s.String()
}

func (s *InsightsConfig) GetLanguage() *string {
	return s.Language
}

func (s *InsightsConfig) SetLanguage(v string) *InsightsConfig {
	s.Language = &v
	return s
}

func (s *InsightsConfig) Validate() error {
	return dara.Validate(s)
}
