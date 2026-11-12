// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCopilotSessionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCopilotSessionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCopilotSessionsResponse
	GetStatusCode() *int32
	SetBody(v *ListCopilotSessionsResponseBody) *ListCopilotSessionsResponse
	GetBody() *ListCopilotSessionsResponseBody
}

type ListCopilotSessionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCopilotSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCopilotSessionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCopilotSessionsResponse) GoString() string {
	return s.String()
}

func (s *ListCopilotSessionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCopilotSessionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCopilotSessionsResponse) GetBody() *ListCopilotSessionsResponseBody {
	return s.Body
}

func (s *ListCopilotSessionsResponse) SetHeaders(v map[string]*string) *ListCopilotSessionsResponse {
	s.Headers = v
	return s
}

func (s *ListCopilotSessionsResponse) SetStatusCode(v int32) *ListCopilotSessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCopilotSessionsResponse) SetBody(v *ListCopilotSessionsResponseBody) *ListCopilotSessionsResponse {
	s.Body = v
	return s
}

func (s *ListCopilotSessionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
