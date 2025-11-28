// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTTSVoiceCustomResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTTSVoiceCustomResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTTSVoiceCustomResponse
	GetStatusCode() *int32
	SetBody(v *CreateTTSVoiceCustomResponseBody) *CreateTTSVoiceCustomResponse
	GetBody() *CreateTTSVoiceCustomResponseBody
}

type CreateTTSVoiceCustomResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTTSVoiceCustomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTTSVoiceCustomResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTTSVoiceCustomResponse) GoString() string {
	return s.String()
}

func (s *CreateTTSVoiceCustomResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTTSVoiceCustomResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTTSVoiceCustomResponse) GetBody() *CreateTTSVoiceCustomResponseBody {
	return s.Body
}

func (s *CreateTTSVoiceCustomResponse) SetHeaders(v map[string]*string) *CreateTTSVoiceCustomResponse {
	s.Headers = v
	return s
}

func (s *CreateTTSVoiceCustomResponse) SetStatusCode(v int32) *CreateTTSVoiceCustomResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTTSVoiceCustomResponse) SetBody(v *CreateTTSVoiceCustomResponseBody) *CreateTTSVoiceCustomResponse {
	s.Body = v
	return s
}

func (s *CreateTTSVoiceCustomResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
