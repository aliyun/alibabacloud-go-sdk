// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertMuxConfig interface {
	dara.Model
	String() string
	GoString() string
	SetSegment(v *MediaConvertSegment) *MediaConvertMuxConfig
	GetSegment() *MediaConvertSegment
}

type MediaConvertMuxConfig struct {
	Segment *MediaConvertSegment `json:"Segment,omitempty" xml:"Segment,omitempty"`
}

func (s MediaConvertMuxConfig) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertMuxConfig) GoString() string {
	return s.String()
}

func (s *MediaConvertMuxConfig) GetSegment() *MediaConvertSegment {
	return s.Segment
}

func (s *MediaConvertMuxConfig) SetSegment(v *MediaConvertSegment) *MediaConvertMuxConfig {
	s.Segment = v
	return s
}

func (s *MediaConvertMuxConfig) Validate() error {
	if s.Segment != nil {
		if err := s.Segment.Validate(); err != nil {
			return err
		}
	}
	return nil
}
