// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCopilotSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCopilotSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCopilotSessionResponse
	GetStatusCode() *int32
	SetBody(v *GetCopilotSessionResponseBody) *GetCopilotSessionResponse
	GetBody() *GetCopilotSessionResponseBody
}

type GetCopilotSessionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCopilotSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCopilotSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCopilotSessionResponse) GoString() string {
	return s.String()
}

func (s *GetCopilotSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCopilotSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCopilotSessionResponse) GetBody() *GetCopilotSessionResponseBody {
	return s.Body
}

func (s *GetCopilotSessionResponse) SetHeaders(v map[string]*string) *GetCopilotSessionResponse {
	s.Headers = v
	return s
}

func (s *GetCopilotSessionResponse) SetStatusCode(v int32) *GetCopilotSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCopilotSessionResponse) SetBody(v *GetCopilotSessionResponseBody) *GetCopilotSessionResponse {
	s.Body = v
	return s
}

func (s *GetCopilotSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
