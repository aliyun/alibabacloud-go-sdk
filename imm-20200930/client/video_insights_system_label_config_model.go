// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoInsightsSystemLabelConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *VideoInsightsSystemLabelConfig
	GetEnable() *bool
}

type VideoInsightsSystemLabelConfig struct {
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s VideoInsightsSystemLabelConfig) String() string {
	return dara.Prettify(s)
}

func (s VideoInsightsSystemLabelConfig) GoString() string {
	return s.String()
}

func (s *VideoInsightsSystemLabelConfig) GetEnable() *bool {
	return s.Enable
}

func (s *VideoInsightsSystemLabelConfig) SetEnable(v bool) *VideoInsightsSystemLabelConfig {
	s.Enable = &v
	return s
}

func (s *VideoInsightsSystemLabelConfig) Validate() error {
	return dara.Validate(s)
}
