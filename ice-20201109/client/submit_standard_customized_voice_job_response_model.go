// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitStandardCustomizedVoiceJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitStandardCustomizedVoiceJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitStandardCustomizedVoiceJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitStandardCustomizedVoiceJobResponseBody) *SubmitStandardCustomizedVoiceJobResponse
	GetBody() *SubmitStandardCustomizedVoiceJobResponseBody
}

type SubmitStandardCustomizedVoiceJobResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitStandardCustomizedVoiceJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitStandardCustomizedVoiceJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitStandardCustomizedVoiceJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitStandardCustomizedVoiceJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitStandardCustomizedVoiceJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitStandardCustomizedVoiceJobResponse) GetBody() *SubmitStandardCustomizedVoiceJobResponseBody {
	return s.Body
}

func (s *SubmitStandardCustomizedVoiceJobResponse) SetHeaders(v map[string]*string) *SubmitStandardCustomizedVoiceJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitStandardCustomizedVoiceJobResponse) SetStatusCode(v int32) *SubmitStandardCustomizedVoiceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitStandardCustomizedVoiceJobResponse) SetBody(v *SubmitStandardCustomizedVoiceJobResponseBody) *SubmitStandardCustomizedVoiceJobResponse {
	s.Body = v
	return s
}

func (s *SubmitStandardCustomizedVoiceJobResponse) Validate() error {
	return dara.Validate(s)
}
