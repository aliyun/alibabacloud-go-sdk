// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAudioFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAudioFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAudioFileResponse
	GetStatusCode() *int32
	SetBody(v *CreateAudioFileResponseBody) *CreateAudioFileResponse
	GetBody() *CreateAudioFileResponseBody
}

type CreateAudioFileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAudioFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAudioFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAudioFileResponse) GoString() string {
	return s.String()
}

func (s *CreateAudioFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAudioFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAudioFileResponse) GetBody() *CreateAudioFileResponseBody {
	return s.Body
}

func (s *CreateAudioFileResponse) SetHeaders(v map[string]*string) *CreateAudioFileResponse {
	s.Headers = v
	return s
}

func (s *CreateAudioFileResponse) SetStatusCode(v int32) *CreateAudioFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAudioFileResponse) SetBody(v *CreateAudioFileResponseBody) *CreateAudioFileResponse {
	s.Body = v
	return s
}

func (s *CreateAudioFileResponse) Validate() error {
	return dara.Validate(s)
}
