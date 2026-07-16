// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoInsightsLabelConfig interface {
	dara.Model
	String() string
	GoString() string
	SetHighlight(v *VideoInsightsHighlightLabelConfig) *VideoInsightsLabelConfig
	GetHighlight() *VideoInsightsHighlightLabelConfig
	SetSystem(v *VideoInsightsSystemLabelConfig) *VideoInsightsLabelConfig
	GetSystem() *VideoInsightsSystemLabelConfig
	SetUserDefined(v *VideoInsightsUserDefinedLabelConfig) *VideoInsightsLabelConfig
	GetUserDefined() *VideoInsightsUserDefinedLabelConfig
}

type VideoInsightsLabelConfig struct {
	// The highlight label configuration.
	Highlight *VideoInsightsHighlightLabelConfig `json:"Highlight,omitempty" xml:"Highlight,omitempty"`
	// The system label configuration.
	System *VideoInsightsSystemLabelConfig `json:"System,omitempty" xml:"System,omitempty"`
	// The custom label configuration.
	UserDefined *VideoInsightsUserDefinedLabelConfig `json:"UserDefined,omitempty" xml:"UserDefined,omitempty"`
}

func (s VideoInsightsLabelConfig) String() string {
	return dara.Prettify(s)
}

func (s VideoInsightsLabelConfig) GoString() string {
	return s.String()
}

func (s *VideoInsightsLabelConfig) GetHighlight() *VideoInsightsHighlightLabelConfig {
	return s.Highlight
}

func (s *VideoInsightsLabelConfig) GetSystem() *VideoInsightsSystemLabelConfig {
	return s.System
}

func (s *VideoInsightsLabelConfig) GetUserDefined() *VideoInsightsUserDefinedLabelConfig {
	return s.UserDefined
}

func (s *VideoInsightsLabelConfig) SetHighlight(v *VideoInsightsHighlightLabelConfig) *VideoInsightsLabelConfig {
	s.Highlight = v
	return s
}

func (s *VideoInsightsLabelConfig) SetSystem(v *VideoInsightsSystemLabelConfig) *VideoInsightsLabelConfig {
	s.System = v
	return s
}

func (s *VideoInsightsLabelConfig) SetUserDefined(v *VideoInsightsUserDefinedLabelConfig) *VideoInsightsLabelConfig {
	s.UserDefined = v
	return s
}

func (s *VideoInsightsLabelConfig) Validate() error {
	if s.Highlight != nil {
		if err := s.Highlight.Validate(); err != nil {
			return err
		}
	}
	if s.System != nil {
		if err := s.System.Validate(); err != nil {
			return err
		}
	}
	if s.UserDefined != nil {
		if err := s.UserDefined.Validate(); err != nil {
			return err
		}
	}
	return nil
}
