// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsights interface {
	dara.Model
	String() string
	GoString() string
	SetImage(v *ImageInsight) *Insights
	GetImage() *ImageInsight
	SetVideo(v *VideoInsight) *Insights
	GetVideo() *VideoInsight
}

type Insights struct {
	// if can be null:
	// true
	Image *ImageInsight `json:"Image,omitempty" xml:"Image,omitempty"`
	// if can be null:
	// true
	Video *VideoInsight `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s Insights) String() string {
	return dara.Prettify(s)
}

func (s Insights) GoString() string {
	return s.String()
}

func (s *Insights) GetImage() *ImageInsight {
	return s.Image
}

func (s *Insights) GetVideo() *VideoInsight {
	return s.Video
}

func (s *Insights) SetImage(v *ImageInsight) *Insights {
	s.Image = v
	return s
}

func (s *Insights) SetVideo(v *VideoInsight) *Insights {
	s.Video = v
	return s
}

func (s *Insights) Validate() error {
	return dara.Validate(s)
}
