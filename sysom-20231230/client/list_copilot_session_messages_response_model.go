// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCopilotSessionMessagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCopilotSessionMessagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCopilotSessionMessagesResponse
	GetStatusCode() *int32
	SetBody(v *ListCopilotSessionMessagesResponseBody) *ListCopilotSessionMessagesResponse
	GetBody() *ListCopilotSessionMessagesResponseBody
}

type ListCopilotSessionMessagesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCopilotSessionMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCopilotSessionMessagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCopilotSessionMessagesResponse) GoString() string {
	return s.String()
}

func (s *ListCopilotSessionMessagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCopilotSessionMessagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCopilotSessionMessagesResponse) GetBody() *ListCopilotSessionMessagesResponseBody {
	return s.Body
}

func (s *ListCopilotSessionMessagesResponse) SetHeaders(v map[string]*string) *ListCopilotSessionMessagesResponse {
	s.Headers = v
	return s
}

func (s *ListCopilotSessionMessagesResponse) SetStatusCode(v int32) *ListCopilotSessionMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCopilotSessionMessagesResponse) SetBody(v *ListCopilotSessionMessagesResponseBody) *ListCopilotSessionMessagesResponse {
	s.Body = v
	return s
}

func (s *ListCopilotSessionMessagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
