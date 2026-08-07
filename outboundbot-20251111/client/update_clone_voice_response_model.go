// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloneVoiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloneVoiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloneVoiceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloneVoiceResponseBody) *UpdateCloneVoiceResponse
	GetBody() *UpdateCloneVoiceResponseBody
}

type UpdateCloneVoiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloneVoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloneVoiceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloneVoiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloneVoiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloneVoiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloneVoiceResponse) GetBody() *UpdateCloneVoiceResponseBody {
	return s.Body
}

func (s *UpdateCloneVoiceResponse) SetHeaders(v map[string]*string) *UpdateCloneVoiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloneVoiceResponse) SetStatusCode(v int32) *UpdateCloneVoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloneVoiceResponse) SetBody(v *UpdateCloneVoiceResponseBody) *UpdateCloneVoiceResponse {
	s.Body = v
	return s
}

func (s *UpdateCloneVoiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
