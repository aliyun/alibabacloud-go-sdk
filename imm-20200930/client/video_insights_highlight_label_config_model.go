// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoInsightsHighlightLabelConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *VideoInsightsHighlightLabelConfig
	GetEnable() *bool
	SetLabels(v []*InsightsLabel) *VideoInsightsHighlightLabelConfig
	GetLabels() []*InsightsLabel
}

type VideoInsightsHighlightLabelConfig struct {
	// Specifies whether highlight labels are supported.
	//
	// example:
	//
	// false
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The list of labels.
	Labels []*InsightsLabel `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s VideoInsightsHighlightLabelConfig) String() string {
	return dara.Prettify(s)
}

func (s VideoInsightsHighlightLabelConfig) GoString() string {
	return s.String()
}

func (s *VideoInsightsHighlightLabelConfig) GetEnable() *bool {
	return s.Enable
}

func (s *VideoInsightsHighlightLabelConfig) GetLabels() []*InsightsLabel {
	return s.Labels
}

func (s *VideoInsightsHighlightLabelConfig) SetEnable(v bool) *VideoInsightsHighlightLabelConfig {
	s.Enable = &v
	return s
}

func (s *VideoInsightsHighlightLabelConfig) SetLabels(v []*InsightsLabel) *VideoInsightsHighlightLabelConfig {
	s.Labels = v
	return s
}

func (s *VideoInsightsHighlightLabelConfig) Validate() error {
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
