// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoMediaMetadata interface {
	dara.Model
	String() string
	GoString() string
	SetHeight(v int64) *VideoMediaMetadata
	GetHeight() *int64
	SetVideoMediaAudioStream(v []*VideoMediaAudioStream) *VideoMediaMetadata
	GetVideoMediaAudioStream() []*VideoMediaAudioStream
	SetVideoMediaVideoStream(v []*VideoMediaVideoStream) *VideoMediaMetadata
	GetVideoMediaVideoStream() []*VideoMediaVideoStream
	SetWidth(v int64) *VideoMediaMetadata
	GetWidth() *int64
}

type VideoMediaMetadata struct {
	// The height of the video image. Unit: pixel.
	//
	// example:
	//
	// 1080
	Height *int64 `json:"height,omitempty" xml:"height,omitempty"`
	// The information about the audio stream.
	VideoMediaAudioStream []*VideoMediaAudioStream `json:"video_media_audio_stream,omitempty" xml:"video_media_audio_stream,omitempty" type:"Repeated"`
	// The information about the video stream.
	VideoMediaVideoStream []*VideoMediaVideoStream `json:"video_media_video_stream,omitempty" xml:"video_media_video_stream,omitempty" type:"Repeated"`
	// The width of the video image. Unit: pixel.
	//
	// example:
	//
	// 1920
	Width *int64 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s VideoMediaMetadata) String() string {
	return dara.Prettify(s)
}

func (s VideoMediaMetadata) GoString() string {
	return s.String()
}

func (s *VideoMediaMetadata) GetHeight() *int64 {
	return s.Height
}

func (s *VideoMediaMetadata) GetVideoMediaAudioStream() []*VideoMediaAudioStream {
	return s.VideoMediaAudioStream
}

func (s *VideoMediaMetadata) GetVideoMediaVideoStream() []*VideoMediaVideoStream {
	return s.VideoMediaVideoStream
}

func (s *VideoMediaMetadata) GetWidth() *int64 {
	return s.Width
}

func (s *VideoMediaMetadata) SetHeight(v int64) *VideoMediaMetadata {
	s.Height = &v
	return s
}

func (s *VideoMediaMetadata) SetVideoMediaAudioStream(v []*VideoMediaAudioStream) *VideoMediaMetadata {
	s.VideoMediaAudioStream = v
	return s
}

func (s *VideoMediaMetadata) SetVideoMediaVideoStream(v []*VideoMediaVideoStream) *VideoMediaMetadata {
	s.VideoMediaVideoStream = v
	return s
}

func (s *VideoMediaMetadata) SetWidth(v int64) *VideoMediaMetadata {
	s.Width = &v
	return s
}

func (s *VideoMediaMetadata) Validate() error {
	if s.VideoMediaAudioStream != nil {
		for _, item := range s.VideoMediaAudioStream {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoMediaVideoStream != nil {
		for _, item := range s.VideoMediaVideoStream {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
