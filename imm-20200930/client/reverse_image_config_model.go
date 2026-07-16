// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReverseImageConfig interface {
	dara.Model
	String() string
	GoString() string
	SetImage(v *ImageReverseImageConfig) *ReverseImageConfig
	GetImage() *ImageReverseImageConfig
	SetVideo(v *VideoReverseImageConfig) *ReverseImageConfig
	GetVideo() *VideoReverseImageConfig
}

type ReverseImageConfig struct {
	// The image-to-image search configuration.
	Image *ImageReverseImageConfig `json:"Image,omitempty" xml:"Image,omitempty"`
	// The image-to-video search configuration.
	Video *VideoReverseImageConfig `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s ReverseImageConfig) String() string {
	return dara.Prettify(s)
}

func (s ReverseImageConfig) GoString() string {
	return s.String()
}

func (s *ReverseImageConfig) GetImage() *ImageReverseImageConfig {
	return s.Image
}

func (s *ReverseImageConfig) GetVideo() *VideoReverseImageConfig {
	return s.Video
}

func (s *ReverseImageConfig) SetImage(v *ImageReverseImageConfig) *ReverseImageConfig {
	s.Image = v
	return s
}

func (s *ReverseImageConfig) SetVideo(v *VideoReverseImageConfig) *ReverseImageConfig {
	s.Video = v
	return s
}

func (s *ReverseImageConfig) Validate() error {
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
