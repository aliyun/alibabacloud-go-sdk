// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoInsightsUserDefinedLabelConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *VideoInsightsUserDefinedLabelConfig
	GetEnable() *bool
	SetLabels(v []*InsightsLabel) *VideoInsightsUserDefinedLabelConfig
	GetLabels() []*InsightsLabel
	SetMode(v string) *VideoInsightsUserDefinedLabelConfig
	GetMode() *string
}

type VideoInsightsUserDefinedLabelConfig struct {
	Enable *bool            `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Labels []*InsightsLabel `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Mode   *string          `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s VideoInsightsUserDefinedLabelConfig) String() string {
	return dara.Prettify(s)
}

func (s VideoInsightsUserDefinedLabelConfig) GoString() string {
	return s.String()
}

func (s *VideoInsightsUserDefinedLabelConfig) GetEnable() *bool {
	return s.Enable
}

func (s *VideoInsightsUserDefinedLabelConfig) GetLabels() []*InsightsLabel {
	return s.Labels
}

func (s *VideoInsightsUserDefinedLabelConfig) GetMode() *string {
	return s.Mode
}

func (s *VideoInsightsUserDefinedLabelConfig) SetEnable(v bool) *VideoInsightsUserDefinedLabelConfig {
	s.Enable = &v
	return s
}

func (s *VideoInsightsUserDefinedLabelConfig) SetLabels(v []*InsightsLabel) *VideoInsightsUserDefinedLabelConfig {
	s.Labels = v
	return s
}

func (s *VideoInsightsUserDefinedLabelConfig) SetMode(v string) *VideoInsightsUserDefinedLabelConfig {
	s.Mode = &v
	return s
}

func (s *VideoInsightsUserDefinedLabelConfig) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
