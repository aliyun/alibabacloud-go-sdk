// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoInsightsConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCaption(v *VideoInsightsCaptionConfig) *VideoInsightsConfig
	GetCaption() *VideoInsightsCaptionConfig
	SetLabel(v *VideoInsightsLabelConfig) *VideoInsightsConfig
	GetLabel() *VideoInsightsLabelConfig
	SetMultiStream(v *VideoInsightsMultiStreamConfig) *VideoInsightsConfig
	GetMultiStream() *VideoInsightsMultiStreamConfig
}

type VideoInsightsConfig struct {
	// The video synopsis configuration.
	Caption *VideoInsightsCaptionConfig `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The label configuration.
	Label *VideoInsightsLabelConfig `json:"Label,omitempty" xml:"Label,omitempty"`
	// The video multi-stream configuration.
	MultiStream *VideoInsightsMultiStreamConfig `json:"MultiStream,omitempty" xml:"MultiStream,omitempty"`
}

func (s VideoInsightsConfig) String() string {
	return dara.Prettify(s)
}

func (s VideoInsightsConfig) GoString() string {
	return s.String()
}

func (s *VideoInsightsConfig) GetCaption() *VideoInsightsCaptionConfig {
	return s.Caption
}

func (s *VideoInsightsConfig) GetLabel() *VideoInsightsLabelConfig {
	return s.Label
}

func (s *VideoInsightsConfig) GetMultiStream() *VideoInsightsMultiStreamConfig {
	return s.MultiStream
}

func (s *VideoInsightsConfig) SetCaption(v *VideoInsightsCaptionConfig) *VideoInsightsConfig {
	s.Caption = v
	return s
}

func (s *VideoInsightsConfig) SetLabel(v *VideoInsightsLabelConfig) *VideoInsightsConfig {
	s.Label = v
	return s
}

func (s *VideoInsightsConfig) SetMultiStream(v *VideoInsightsMultiStreamConfig) *VideoInsightsConfig {
	s.MultiStream = v
	return s
}

func (s *VideoInsightsConfig) Validate() error {
	if s.Caption != nil {
		if err := s.Caption.Validate(); err != nil {
			return err
		}
	}
	if s.Label != nil {
		if err := s.Label.Validate(); err != nil {
			return err
		}
	}
	if s.MultiStream != nil {
		if err := s.MultiStream.Validate(); err != nil {
			return err
		}
	}
	return nil
}
