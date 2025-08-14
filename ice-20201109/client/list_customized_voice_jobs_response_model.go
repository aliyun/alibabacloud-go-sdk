// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomizedVoiceJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomizedVoiceJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomizedVoiceJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomizedVoiceJobsResponseBody) *ListCustomizedVoiceJobsResponse
	GetBody() *ListCustomizedVoiceJobsResponseBody
}

type ListCustomizedVoiceJobsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomizedVoiceJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomizedVoiceJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomizedVoiceJobsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomizedVoiceJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomizedVoiceJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomizedVoiceJobsResponse) GetBody() *ListCustomizedVoiceJobsResponseBody {
	return s.Body
}

func (s *ListCustomizedVoiceJobsResponse) SetHeaders(v map[string]*string) *ListCustomizedVoiceJobsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomizedVoiceJobsResponse) SetStatusCode(v int32) *ListCustomizedVoiceJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomizedVoiceJobsResponse) SetBody(v *ListCustomizedVoiceJobsResponseBody) *ListCustomizedVoiceJobsResponse {
	s.Body = v
	return s
}

func (s *ListCustomizedVoiceJobsResponse) Validate() error {
	return dara.Validate(s)
}
