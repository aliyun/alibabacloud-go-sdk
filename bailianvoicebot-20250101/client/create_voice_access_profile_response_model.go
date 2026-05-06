// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVoiceAccessProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVoiceAccessProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVoiceAccessProfileResponse
	GetStatusCode() *int32
	SetBody(v *CreateVoiceAccessProfileResponseBody) *CreateVoiceAccessProfileResponse
	GetBody() *CreateVoiceAccessProfileResponseBody
}

type CreateVoiceAccessProfileResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVoiceAccessProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVoiceAccessProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVoiceAccessProfileResponse) GoString() string {
	return s.String()
}

func (s *CreateVoiceAccessProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVoiceAccessProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVoiceAccessProfileResponse) GetBody() *CreateVoiceAccessProfileResponseBody {
	return s.Body
}

func (s *CreateVoiceAccessProfileResponse) SetHeaders(v map[string]*string) *CreateVoiceAccessProfileResponse {
	s.Headers = v
	return s
}

func (s *CreateVoiceAccessProfileResponse) SetStatusCode(v int32) *CreateVoiceAccessProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVoiceAccessProfileResponse) SetBody(v *CreateVoiceAccessProfileResponseBody) *CreateVoiceAccessProfileResponse {
	s.Body = v
	return s
}

func (s *CreateVoiceAccessProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
