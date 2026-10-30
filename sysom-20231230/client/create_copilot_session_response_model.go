// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCopilotSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCopilotSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCopilotSessionResponse
	GetStatusCode() *int32
	SetBody(v *CreateCopilotSessionResponseBody) *CreateCopilotSessionResponse
	GetBody() *CreateCopilotSessionResponseBody
}

type CreateCopilotSessionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCopilotSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCopilotSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCopilotSessionResponse) GoString() string {
	return s.String()
}

func (s *CreateCopilotSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCopilotSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCopilotSessionResponse) GetBody() *CreateCopilotSessionResponseBody {
	return s.Body
}

func (s *CreateCopilotSessionResponse) SetHeaders(v map[string]*string) *CreateCopilotSessionResponse {
	s.Headers = v
	return s
}

func (s *CreateCopilotSessionResponse) SetStatusCode(v int32) *CreateCopilotSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCopilotSessionResponse) SetBody(v *CreateCopilotSessionResponseBody) *CreateCopilotSessionResponse {
	s.Body = v
	return s
}

func (s *CreateCopilotSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
