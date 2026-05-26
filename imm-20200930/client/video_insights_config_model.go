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
}

type VideoInsightsConfig struct {
	Caption *VideoInsightsCaptionConfig `json:"Caption,omitempty" xml:"Caption,omitempty"`
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

func (s *VideoInsightsConfig) SetCaption(v *VideoInsightsCaptionConfig) *VideoInsightsConfig {
	s.Caption = v
	return s
}

func (s *VideoInsightsConfig) Validate() error {
	if s.Caption != nil {
		if err := s.Caption.Validate(); err != nil {
			return err
		}
	}
	return nil
}
