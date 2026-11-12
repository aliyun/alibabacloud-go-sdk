// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCopilotSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCopilotSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCopilotSessionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCopilotSessionResponseBody) *UpdateCopilotSessionResponse
	GetBody() *UpdateCopilotSessionResponseBody
}

type UpdateCopilotSessionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCopilotSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCopilotSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCopilotSessionResponse) GoString() string {
	return s.String()
}

func (s *UpdateCopilotSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCopilotSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCopilotSessionResponse) GetBody() *UpdateCopilotSessionResponseBody {
	return s.Body
}

func (s *UpdateCopilotSessionResponse) SetHeaders(v map[string]*string) *UpdateCopilotSessionResponse {
	s.Headers = v
	return s
}

func (s *UpdateCopilotSessionResponse) SetStatusCode(v int32) *UpdateCopilotSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCopilotSessionResponse) SetBody(v *UpdateCopilotSessionResponseBody) *UpdateCopilotSessionResponse {
	s.Body = v
	return s
}

func (s *UpdateCopilotSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
