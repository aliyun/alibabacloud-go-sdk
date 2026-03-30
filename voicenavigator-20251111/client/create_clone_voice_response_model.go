// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloneVoiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloneVoiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloneVoiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloneVoiceResponseBody) *CreateCloneVoiceResponse
	GetBody() *CreateCloneVoiceResponseBody
}

type CreateCloneVoiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloneVoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloneVoiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloneVoiceResponse) GoString() string {
	return s.String()
}

func (s *CreateCloneVoiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloneVoiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloneVoiceResponse) GetBody() *CreateCloneVoiceResponseBody {
	return s.Body
}

func (s *CreateCloneVoiceResponse) SetHeaders(v map[string]*string) *CreateCloneVoiceResponse {
	s.Headers = v
	return s
}

func (s *CreateCloneVoiceResponse) SetStatusCode(v int32) *CreateCloneVoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloneVoiceResponse) SetBody(v *CreateCloneVoiceResponseBody) *CreateCloneVoiceResponse {
	s.Body = v
	return s
}

func (s *CreateCloneVoiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
