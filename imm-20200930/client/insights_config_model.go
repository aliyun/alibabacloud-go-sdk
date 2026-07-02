// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsightsConfig interface {
	dara.Model
	String() string
	GoString() string
	SetImage(v *ImageInsightsConfig) *InsightsConfig
	GetImage() *ImageInsightsConfig
	SetLanguage(v string) *InsightsConfig
	GetLanguage() *string
	SetVideo(v *VideoInsightsConfig) *InsightsConfig
	GetVideo() *VideoInsightsConfig
}

type InsightsConfig struct {
	Image *ImageInsightsConfig `json:"Image,omitempty" xml:"Image,omitempty"`
	// The language.
	//
	// example:
	//
	// zh-Hans
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The video content-aware configuration.
	Video *VideoInsightsConfig `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s InsightsConfig) String() string {
	return dara.Prettify(s)
}

func (s InsightsConfig) GoString() string {
	return s.String()
}

func (s *InsightsConfig) GetImage() *ImageInsightsConfig {
	return s.Image
}

func (s *InsightsConfig) GetLanguage() *string {
	return s.Language
}

func (s *InsightsConfig) GetVideo() *VideoInsightsConfig {
	return s.Video
}

func (s *InsightsConfig) SetImage(v *ImageInsightsConfig) *InsightsConfig {
	s.Image = v
	return s
}

func (s *InsightsConfig) SetLanguage(v string) *InsightsConfig {
	s.Language = &v
	return s
}

func (s *InsightsConfig) SetVideo(v *VideoInsightsConfig) *InsightsConfig {
	s.Video = v
	return s
}

func (s *InsightsConfig) Validate() error {
	if s.Image != nil {
		if err := s.Image.Validate(); err != nil {
			return err
		}
	}
	if s.Video != nil {
		if err := s.Video.Validate(); err != nil {
			return err
		}
	}
	return nil
}
