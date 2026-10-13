// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoiceCloneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVoiceCloneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVoiceCloneResponse
	GetStatusCode() *int32
	SetBody(v *ListVoiceCloneResponseBody) *ListVoiceCloneResponse
	GetBody() *ListVoiceCloneResponseBody
}

type ListVoiceCloneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVoiceCloneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVoiceCloneResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceCloneResponse) GoString() string {
	return s.String()
}

func (s *ListVoiceCloneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVoiceCloneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVoiceCloneResponse) GetBody() *ListVoiceCloneResponseBody {
	return s.Body
}

func (s *ListVoiceCloneResponse) SetHeaders(v map[string]*string) *ListVoiceCloneResponse {
	s.Headers = v
	return s
}

func (s *ListVoiceCloneResponse) SetStatusCode(v int32) *ListVoiceCloneResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVoiceCloneResponse) SetBody(v *ListVoiceCloneResponseBody) *ListVoiceCloneResponse {
	s.Body = v
	return s
}

func (s *ListVoiceCloneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
