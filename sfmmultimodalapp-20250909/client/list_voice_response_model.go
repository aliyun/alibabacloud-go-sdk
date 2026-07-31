// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVoiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVoiceResponse
	GetStatusCode() *int32
	SetBody(v *ListVoiceResponseBody) *ListVoiceResponse
	GetBody() *ListVoiceResponseBody
}

type ListVoiceResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVoiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceResponse) GoString() string {
	return s.String()
}

func (s *ListVoiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVoiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVoiceResponse) GetBody() *ListVoiceResponseBody {
	return s.Body
}

func (s *ListVoiceResponse) SetHeaders(v map[string]*string) *ListVoiceResponse {
	s.Headers = v
	return s
}

func (s *ListVoiceResponse) SetStatusCode(v int32) *ListVoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVoiceResponse) SetBody(v *ListVoiceResponseBody) *ListVoiceResponse {
	s.Body = v
	return s
}

func (s *ListVoiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
