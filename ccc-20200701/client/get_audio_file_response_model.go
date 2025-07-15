// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAudioFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAudioFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAudioFileResponse
	GetStatusCode() *int32
	SetBody(v *GetAudioFileResponseBody) *GetAudioFileResponse
	GetBody() *GetAudioFileResponseBody
}

type GetAudioFileResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAudioFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAudioFileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAudioFileResponse) GoString() string {
	return s.String()
}

func (s *GetAudioFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAudioFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAudioFileResponse) GetBody() *GetAudioFileResponseBody {
	return s.Body
}

func (s *GetAudioFileResponse) SetHeaders(v map[string]*string) *GetAudioFileResponse {
	s.Headers = v
	return s
}

func (s *GetAudioFileResponse) SetStatusCode(v int32) *GetAudioFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAudioFileResponse) SetBody(v *GetAudioFileResponseBody) *GetAudioFileResponse {
	s.Body = v
	return s
}

func (s *GetAudioFileResponse) Validate() error {
	return dara.Validate(s)
}
