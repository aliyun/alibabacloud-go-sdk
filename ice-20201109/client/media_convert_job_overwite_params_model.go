// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertJobOverwiteParams interface {
	dara.Model
	String() string
	GoString() string
	SetSubtitles(v []*MediaConvertJobOverwiteParamsSubtitles) *MediaConvertJobOverwiteParams
	GetSubtitles() []*MediaConvertJobOverwiteParamsSubtitles
}

type MediaConvertJobOverwiteParams struct {
	Subtitles []*MediaConvertJobOverwiteParamsSubtitles `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Repeated"`
}

func (s MediaConvertJobOverwiteParams) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertJobOverwiteParams) GoString() string {
	return s.String()
}

func (s *MediaConvertJobOverwiteParams) GetSubtitles() []*MediaConvertJobOverwiteParamsSubtitles {
	return s.Subtitles
}

func (s *MediaConvertJobOverwiteParams) SetSubtitles(v []*MediaConvertJobOverwiteParamsSubtitles) *MediaConvertJobOverwiteParams {
	s.Subtitles = v
	return s
}

func (s *MediaConvertJobOverwiteParams) Validate() error {
	if s.Subtitles != nil {
		for _, item := range s.Subtitles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MediaConvertJobOverwiteParamsSubtitles struct {
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
}

func (s MediaConvertJobOverwiteParamsSubtitles) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertJobOverwiteParamsSubtitles) GoString() string {
	return s.String()
}

func (s *MediaConvertJobOverwiteParamsSubtitles) GetCodec() *string {
	return s.Codec
}

func (s *MediaConvertJobOverwiteParamsSubtitles) SetCodec(v string) *MediaConvertJobOverwiteParamsSubtitles {
	s.Codec = &v
	return s
}

func (s *MediaConvertJobOverwiteParamsSubtitles) Validate() error {
	return dara.Validate(s)
}
