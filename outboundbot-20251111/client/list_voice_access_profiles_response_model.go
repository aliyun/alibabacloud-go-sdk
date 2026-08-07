// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoiceAccessProfilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVoiceAccessProfilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVoiceAccessProfilesResponse
	GetStatusCode() *int32
	SetBody(v *ListVoiceAccessProfilesResponseBody) *ListVoiceAccessProfilesResponse
	GetBody() *ListVoiceAccessProfilesResponseBody
}

type ListVoiceAccessProfilesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVoiceAccessProfilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVoiceAccessProfilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceAccessProfilesResponse) GoString() string {
	return s.String()
}

func (s *ListVoiceAccessProfilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVoiceAccessProfilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVoiceAccessProfilesResponse) GetBody() *ListVoiceAccessProfilesResponseBody {
	return s.Body
}

func (s *ListVoiceAccessProfilesResponse) SetHeaders(v map[string]*string) *ListVoiceAccessProfilesResponse {
	s.Headers = v
	return s
}

func (s *ListVoiceAccessProfilesResponse) SetStatusCode(v int32) *ListVoiceAccessProfilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVoiceAccessProfilesResponse) SetBody(v *ListVoiceAccessProfilesResponseBody) *ListVoiceAccessProfilesResponse {
	s.Body = v
	return s
}

func (s *ListVoiceAccessProfilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
