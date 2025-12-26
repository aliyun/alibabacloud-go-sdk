// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTTSVoiceByIdCustomResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTTSVoiceByIdCustomResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTTSVoiceByIdCustomResponse
	GetStatusCode() *int32
	SetBody(v *GetTTSVoiceByIdCustomResponseBody) *GetTTSVoiceByIdCustomResponse
	GetBody() *GetTTSVoiceByIdCustomResponseBody
}

type GetTTSVoiceByIdCustomResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTTSVoiceByIdCustomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTTSVoiceByIdCustomResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTTSVoiceByIdCustomResponse) GoString() string {
	return s.String()
}

func (s *GetTTSVoiceByIdCustomResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTTSVoiceByIdCustomResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTTSVoiceByIdCustomResponse) GetBody() *GetTTSVoiceByIdCustomResponseBody {
	return s.Body
}

func (s *GetTTSVoiceByIdCustomResponse) SetHeaders(v map[string]*string) *GetTTSVoiceByIdCustomResponse {
	s.Headers = v
	return s
}

func (s *GetTTSVoiceByIdCustomResponse) SetStatusCode(v int32) *GetTTSVoiceByIdCustomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTTSVoiceByIdCustomResponse) SetBody(v *GetTTSVoiceByIdCustomResponseBody) *GetTTSVoiceByIdCustomResponse {
	s.Body = v
	return s
}

func (s *GetTTSVoiceByIdCustomResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
