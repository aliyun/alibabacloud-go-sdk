// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoMediaAudioStream interface {
	dara.Model
	String() string
	GoString() string
	SetBitRate(v string) *VideoMediaAudioStream
	GetBitRate() *string
	SetCodeName(v string) *VideoMediaAudioStream
	GetCodeName() *string
	SetDuration(v string) *VideoMediaAudioStream
	GetDuration() *string
}

type VideoMediaAudioStream struct {
	// The bitrate of the audio stream. Unit: bit/s.
	//
	// example:
	//
	// 129280
	BitRate *string `json:"bit_rate,omitempty" xml:"bit_rate,omitempty"`
	// The audio encoding mode.
	//
	// example:
	//
	// aac
	CodeName *string `json:"code_name,omitempty" xml:"code_name,omitempty"`
	// The duration of the audio stream. Unit: seconds.
	//
	// example:
	//
	// 7704.573000
	Duration *string `json:"duration,omitempty" xml:"duration,omitempty"`
}

func (s VideoMediaAudioStream) String() string {
	return dara.Prettify(s)
}

func (s VideoMediaAudioStream) GoString() string {
	return s.String()
}

func (s *VideoMediaAudioStream) GetBitRate() *string {
	return s.BitRate
}

func (s *VideoMediaAudioStream) GetCodeName() *string {
	return s.CodeName
}

func (s *VideoMediaAudioStream) GetDuration() *string {
	return s.Duration
}

func (s *VideoMediaAudioStream) SetBitRate(v string) *VideoMediaAudioStream {
	s.BitRate = &v
	return s
}

func (s *VideoMediaAudioStream) SetCodeName(v string) *VideoMediaAudioStream {
	s.CodeName = &v
	return s
}

func (s *VideoMediaAudioStream) SetDuration(v string) *VideoMediaAudioStream {
	s.Duration = &v
	return s
}

func (s *VideoMediaAudioStream) Validate() error {
	return dara.Validate(s)
}
