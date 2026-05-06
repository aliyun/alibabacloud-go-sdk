// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVoiceAccessProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVoiceAccessProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVoiceAccessProfileResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVoiceAccessProfileResponseBody) *UpdateVoiceAccessProfileResponse
	GetBody() *UpdateVoiceAccessProfileResponseBody
}

type UpdateVoiceAccessProfileResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVoiceAccessProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVoiceAccessProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVoiceAccessProfileResponse) GoString() string {
	return s.String()
}

func (s *UpdateVoiceAccessProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVoiceAccessProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVoiceAccessProfileResponse) GetBody() *UpdateVoiceAccessProfileResponseBody {
	return s.Body
}

func (s *UpdateVoiceAccessProfileResponse) SetHeaders(v map[string]*string) *UpdateVoiceAccessProfileResponse {
	s.Headers = v
	return s
}

func (s *UpdateVoiceAccessProfileResponse) SetStatusCode(v int32) *UpdateVoiceAccessProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVoiceAccessProfileResponse) SetBody(v *UpdateVoiceAccessProfileResponseBody) *UpdateVoiceAccessProfileResponse {
	s.Body = v
	return s
}

func (s *UpdateVoiceAccessProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
