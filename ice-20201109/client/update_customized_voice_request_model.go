// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomizedVoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDemoAudioMediaId(v string) *UpdateCustomizedVoiceRequest
	GetDemoAudioMediaId() *string
	SetVoiceId(v string) *UpdateCustomizedVoiceRequest
	GetVoiceId() *string
}

type UpdateCustomizedVoiceRequest struct {
	// The media asset ID of the sample audio file.
	//
	// example:
	//
	// ****4d5e829d498aaf966b119348****
	DemoAudioMediaId *string `json:"DemoAudioMediaId,omitempty" xml:"DemoAudioMediaId,omitempty"`
	// The voice ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// xiaozhuan
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
}

func (s UpdateCustomizedVoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomizedVoiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomizedVoiceRequest) GetDemoAudioMediaId() *string {
	return s.DemoAudioMediaId
}

func (s *UpdateCustomizedVoiceRequest) GetVoiceId() *string {
	return s.VoiceId
}

func (s *UpdateCustomizedVoiceRequest) SetDemoAudioMediaId(v string) *UpdateCustomizedVoiceRequest {
	s.DemoAudioMediaId = &v
	return s
}

func (s *UpdateCustomizedVoiceRequest) SetVoiceId(v string) *UpdateCustomizedVoiceRequest {
	s.VoiceId = &v
	return s
}

func (s *UpdateCustomizedVoiceRequest) Validate() error {
	return dara.Validate(s)
}
