// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoiceEnginesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVoiceEnginesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVoiceEnginesResponse
	GetStatusCode() *int32
	SetBody(v *ListVoiceEnginesResponseBody) *ListVoiceEnginesResponse
	GetBody() *ListVoiceEnginesResponseBody
}

type ListVoiceEnginesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVoiceEnginesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVoiceEnginesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceEnginesResponse) GoString() string {
	return s.String()
}

func (s *ListVoiceEnginesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVoiceEnginesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVoiceEnginesResponse) GetBody() *ListVoiceEnginesResponseBody {
	return s.Body
}

func (s *ListVoiceEnginesResponse) SetHeaders(v map[string]*string) *ListVoiceEnginesResponse {
	s.Headers = v
	return s
}

func (s *ListVoiceEnginesResponse) SetStatusCode(v int32) *ListVoiceEnginesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVoiceEnginesResponse) SetBody(v *ListVoiceEnginesResponseBody) *ListVoiceEnginesResponse {
	s.Body = v
	return s
}

func (s *ListVoiceEnginesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
