// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewVoiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreviewVoiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreviewVoiceResponse
	GetStatusCode() *int32
	SetBody(v *PreviewVoiceResponseBody) *PreviewVoiceResponse
	GetBody() *PreviewVoiceResponseBody
}

type PreviewVoiceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreviewVoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewVoiceResponse) String() string {
	return dara.Prettify(s)
}

func (s PreviewVoiceResponse) GoString() string {
	return s.String()
}

func (s *PreviewVoiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreviewVoiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreviewVoiceResponse) GetBody() *PreviewVoiceResponseBody {
	return s.Body
}

func (s *PreviewVoiceResponse) SetHeaders(v map[string]*string) *PreviewVoiceResponse {
	s.Headers = v
	return s
}

func (s *PreviewVoiceResponse) SetStatusCode(v int32) *PreviewVoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewVoiceResponse) SetBody(v *PreviewVoiceResponseBody) *PreviewVoiceResponse {
	s.Body = v
	return s
}

func (s *PreviewVoiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
