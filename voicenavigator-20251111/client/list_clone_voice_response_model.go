// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloneVoiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloneVoiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloneVoiceResponse
	GetStatusCode() *int32
	SetBody(v *ListCloneVoiceResponseBody) *ListCloneVoiceResponse
	GetBody() *ListCloneVoiceResponseBody
}

type ListCloneVoiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloneVoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloneVoiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloneVoiceResponse) GoString() string {
	return s.String()
}

func (s *ListCloneVoiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloneVoiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloneVoiceResponse) GetBody() *ListCloneVoiceResponseBody {
	return s.Body
}

func (s *ListCloneVoiceResponse) SetHeaders(v map[string]*string) *ListCloneVoiceResponse {
	s.Headers = v
	return s
}

func (s *ListCloneVoiceResponse) SetStatusCode(v int32) *ListCloneVoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloneVoiceResponse) SetBody(v *ListCloneVoiceResponseBody) *ListCloneVoiceResponse {
	s.Body = v
	return s
}

func (s *ListCloneVoiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
