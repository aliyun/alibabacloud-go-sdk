// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageInsightsConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCaption(v *ImageInsightsCaptionConfig) *ImageInsightsConfig
	GetCaption() *ImageInsightsCaptionConfig
}

type ImageInsightsConfig struct {
	// The image content recognition Caption configuration.
	Caption *ImageInsightsCaptionConfig `json:"Caption,omitempty" xml:"Caption,omitempty"`
}

func (s ImageInsightsConfig) String() string {
	return dara.Prettify(s)
}

func (s ImageInsightsConfig) GoString() string {
	return s.String()
}

func (s *ImageInsightsConfig) GetCaption() *ImageInsightsCaptionConfig {
	return s.Caption
}

func (s *ImageInsightsConfig) SetCaption(v *ImageInsightsCaptionConfig) *ImageInsightsConfig {
	s.Caption = v
	return s
}

func (s *ImageInsightsConfig) Validate() error {
	if s.Caption != nil {
		if err := s.Caption.Validate(); err != nil {
			return err
		}
	}
	return nil
}
