// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDemonstrationForCustomizedVoiceJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDemonstrationForCustomizedVoiceJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDemonstrationForCustomizedVoiceJobResponse
	GetStatusCode() *int32
	SetBody(v *GetDemonstrationForCustomizedVoiceJobResponseBody) *GetDemonstrationForCustomizedVoiceJobResponse
	GetBody() *GetDemonstrationForCustomizedVoiceJobResponseBody
}

type GetDemonstrationForCustomizedVoiceJobResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDemonstrationForCustomizedVoiceJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDemonstrationForCustomizedVoiceJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDemonstrationForCustomizedVoiceJobResponse) GoString() string {
	return s.String()
}

func (s *GetDemonstrationForCustomizedVoiceJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDemonstrationForCustomizedVoiceJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDemonstrationForCustomizedVoiceJobResponse) GetBody() *GetDemonstrationForCustomizedVoiceJobResponseBody {
	return s.Body
}

func (s *GetDemonstrationForCustomizedVoiceJobResponse) SetHeaders(v map[string]*string) *GetDemonstrationForCustomizedVoiceJobResponse {
	s.Headers = v
	return s
}

func (s *GetDemonstrationForCustomizedVoiceJobResponse) SetStatusCode(v int32) *GetDemonstrationForCustomizedVoiceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDemonstrationForCustomizedVoiceJobResponse) SetBody(v *GetDemonstrationForCustomizedVoiceJobResponseBody) *GetDemonstrationForCustomizedVoiceJobResponse {
	s.Body = v
	return s
}

func (s *GetDemonstrationForCustomizedVoiceJobResponse) Validate() error {
	return dara.Validate(s)
}
