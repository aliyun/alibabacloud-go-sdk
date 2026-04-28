// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoMediaVideoStream interface {
	dara.Model
	String() string
	GoString() string
	SetBitrate(v string) *VideoMediaVideoStream
	GetBitrate() *string
	SetCodeName(v string) *VideoMediaVideoStream
	GetCodeName() *string
	SetDuration(v string) *VideoMediaVideoStream
	GetDuration() *string
	SetFrameCount(v string) *VideoMediaVideoStream
	GetFrameCount() *string
}

type VideoMediaVideoStream struct {
	// The bitrate of the video stream. Unit: bit/s.
	//
	// example:
	//
	// 108420
	Bitrate *string `json:"bitrate,omitempty" xml:"bitrate,omitempty"`
	// The video encoding mode.
	//
	// example:
	//
	// h264
	CodeName *string `json:"code_name,omitempty" xml:"code_name,omitempty"`
	// The duration of the video stream. Unit: seconds.
	//
	// example:
	//
	// 22.88
	Duration *string `json:"duration,omitempty" xml:"duration,omitempty"`
	// The number of video frames.
	//
	// example:
	//
	// 90
	FrameCount *string `json:"frame_count,omitempty" xml:"frame_count,omitempty"`
}

func (s VideoMediaVideoStream) String() string {
	return dara.Prettify(s)
}

func (s VideoMediaVideoStream) GoString() string {
	return s.String()
}

func (s *VideoMediaVideoStream) GetBitrate() *string {
	return s.Bitrate
}

func (s *VideoMediaVideoStream) GetCodeName() *string {
	return s.CodeName
}

func (s *VideoMediaVideoStream) GetDuration() *string {
	return s.Duration
}

func (s *VideoMediaVideoStream) GetFrameCount() *string {
	return s.FrameCount
}

func (s *VideoMediaVideoStream) SetBitrate(v string) *VideoMediaVideoStream {
	s.Bitrate = &v
	return s
}

func (s *VideoMediaVideoStream) SetCodeName(v string) *VideoMediaVideoStream {
	s.CodeName = &v
	return s
}

func (s *VideoMediaVideoStream) SetDuration(v string) *VideoMediaVideoStream {
	s.Duration = &v
	return s
}

func (s *VideoMediaVideoStream) SetFrameCount(v string) *VideoMediaVideoStream {
	s.FrameCount = &v
	return s
}

func (s *VideoMediaVideoStream) Validate() error {
	return dara.Validate(s)
}
