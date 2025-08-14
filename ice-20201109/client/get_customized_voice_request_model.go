// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomizedVoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVoiceId(v string) *GetCustomizedVoiceRequest
	GetVoiceId() *string
}

type GetCustomizedVoiceRequest struct {
	// The voice ID.
	//
	// example:
	//
	// xiaozhuan
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
}

func (s GetCustomizedVoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomizedVoiceRequest) GoString() string {
	return s.String()
}

func (s *GetCustomizedVoiceRequest) GetVoiceId() *string {
	return s.VoiceId
}

func (s *GetCustomizedVoiceRequest) SetVoiceId(v string) *GetCustomizedVoiceRequest {
	s.VoiceId = &v
	return s
}

func (s *GetCustomizedVoiceRequest) Validate() error {
	return dara.Validate(s)
}
