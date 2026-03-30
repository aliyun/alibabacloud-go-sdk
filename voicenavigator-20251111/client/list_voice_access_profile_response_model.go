// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoiceAccessProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVoiceAccessProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVoiceAccessProfileResponse
	GetStatusCode() *int32
	SetBody(v *ListVoiceAccessProfileResponseBody) *ListVoiceAccessProfileResponse
	GetBody() *ListVoiceAccessProfileResponseBody
}

type ListVoiceAccessProfileResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVoiceAccessProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVoiceAccessProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceAccessProfileResponse) GoString() string {
	return s.String()
}

func (s *ListVoiceAccessProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVoiceAccessProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVoiceAccessProfileResponse) GetBody() *ListVoiceAccessProfileResponseBody {
	return s.Body
}

func (s *ListVoiceAccessProfileResponse) SetHeaders(v map[string]*string) *ListVoiceAccessProfileResponse {
	s.Headers = v
	return s
}

func (s *ListVoiceAccessProfileResponse) SetStatusCode(v int32) *ListVoiceAccessProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVoiceAccessProfileResponse) SetBody(v *ListVoiceAccessProfileResponseBody) *ListVoiceAccessProfileResponse {
	s.Body = v
	return s
}

func (s *ListVoiceAccessProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
