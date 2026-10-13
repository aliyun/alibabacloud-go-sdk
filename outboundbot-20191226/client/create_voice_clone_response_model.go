// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVoiceCloneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVoiceCloneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVoiceCloneResponse
	GetStatusCode() *int32
	SetBody(v *CreateVoiceCloneResponseBody) *CreateVoiceCloneResponse
	GetBody() *CreateVoiceCloneResponseBody
}

type CreateVoiceCloneResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVoiceCloneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVoiceCloneResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVoiceCloneResponse) GoString() string {
	return s.String()
}

func (s *CreateVoiceCloneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVoiceCloneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVoiceCloneResponse) GetBody() *CreateVoiceCloneResponseBody {
	return s.Body
}

func (s *CreateVoiceCloneResponse) SetHeaders(v map[string]*string) *CreateVoiceCloneResponse {
	s.Headers = v
	return s
}

func (s *CreateVoiceCloneResponse) SetStatusCode(v int32) *CreateVoiceCloneResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVoiceCloneResponse) SetBody(v *CreateVoiceCloneResponseBody) *CreateVoiceCloneResponse {
	s.Body = v
	return s
}

func (s *CreateVoiceCloneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
