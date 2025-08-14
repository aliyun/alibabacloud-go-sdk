// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectAudioForCustomizedVoiceJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectAudioForCustomizedVoiceJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectAudioForCustomizedVoiceJobResponse
	GetStatusCode() *int32
	SetBody(v *DetectAudioForCustomizedVoiceJobResponseBody) *DetectAudioForCustomizedVoiceJobResponse
	GetBody() *DetectAudioForCustomizedVoiceJobResponseBody
}

type DetectAudioForCustomizedVoiceJobResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectAudioForCustomizedVoiceJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectAudioForCustomizedVoiceJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectAudioForCustomizedVoiceJobResponse) GoString() string {
	return s.String()
}

func (s *DetectAudioForCustomizedVoiceJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectAudioForCustomizedVoiceJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectAudioForCustomizedVoiceJobResponse) GetBody() *DetectAudioForCustomizedVoiceJobResponseBody {
	return s.Body
}

func (s *DetectAudioForCustomizedVoiceJobResponse) SetHeaders(v map[string]*string) *DetectAudioForCustomizedVoiceJobResponse {
	s.Headers = v
	return s
}

func (s *DetectAudioForCustomizedVoiceJobResponse) SetStatusCode(v int32) *DetectAudioForCustomizedVoiceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectAudioForCustomizedVoiceJobResponse) SetBody(v *DetectAudioForCustomizedVoiceJobResponseBody) *DetectAudioForCustomizedVoiceJobResponse {
	s.Body = v
	return s
}

func (s *DetectAudioForCustomizedVoiceJobResponse) Validate() error {
	return dara.Validate(s)
}
