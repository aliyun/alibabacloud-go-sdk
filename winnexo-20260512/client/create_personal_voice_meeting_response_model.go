// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalVoiceMeetingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePersonalVoiceMeetingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePersonalVoiceMeetingResponse
	GetStatusCode() *int32
	SetBody(v *CreatePersonalVoiceMeetingResponseBody) *CreatePersonalVoiceMeetingResponse
	GetBody() *CreatePersonalVoiceMeetingResponseBody
}

type CreatePersonalVoiceMeetingResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePersonalVoiceMeetingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePersonalVoiceMeetingResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalVoiceMeetingResponse) GoString() string {
	return s.String()
}

func (s *CreatePersonalVoiceMeetingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePersonalVoiceMeetingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePersonalVoiceMeetingResponse) GetBody() *CreatePersonalVoiceMeetingResponseBody {
	return s.Body
}

func (s *CreatePersonalVoiceMeetingResponse) SetHeaders(v map[string]*string) *CreatePersonalVoiceMeetingResponse {
	s.Headers = v
	return s
}

func (s *CreatePersonalVoiceMeetingResponse) SetStatusCode(v int32) *CreatePersonalVoiceMeetingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePersonalVoiceMeetingResponse) SetBody(v *CreatePersonalVoiceMeetingResponseBody) *CreatePersonalVoiceMeetingResponse {
	s.Body = v
	return s
}

func (s *CreatePersonalVoiceMeetingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
