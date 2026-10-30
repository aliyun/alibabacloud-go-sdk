// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCopilotSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCopilotSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCopilotSessionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCopilotSessionResponseBody) *DeleteCopilotSessionResponse
	GetBody() *DeleteCopilotSessionResponseBody
}

type DeleteCopilotSessionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCopilotSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCopilotSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCopilotSessionResponse) GoString() string {
	return s.String()
}

func (s *DeleteCopilotSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCopilotSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCopilotSessionResponse) GetBody() *DeleteCopilotSessionResponseBody {
	return s.Body
}

func (s *DeleteCopilotSessionResponse) SetHeaders(v map[string]*string) *DeleteCopilotSessionResponse {
	s.Headers = v
	return s
}

func (s *DeleteCopilotSessionResponse) SetStatusCode(v int32) *DeleteCopilotSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCopilotSessionResponse) SetBody(v *DeleteCopilotSessionResponseBody) *DeleteCopilotSessionResponse {
	s.Body = v
	return s
}

func (s *DeleteCopilotSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
