// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCopilotMessagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCopilotMessagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCopilotMessagesResponse
	GetStatusCode() *int32
	SetBody(v *ListCopilotMessagesResponseBody) *ListCopilotMessagesResponse
	GetBody() *ListCopilotMessagesResponseBody
}

type ListCopilotMessagesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCopilotMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCopilotMessagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCopilotMessagesResponse) GoString() string {
	return s.String()
}

func (s *ListCopilotMessagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCopilotMessagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCopilotMessagesResponse) GetBody() *ListCopilotMessagesResponseBody {
	return s.Body
}

func (s *ListCopilotMessagesResponse) SetHeaders(v map[string]*string) *ListCopilotMessagesResponse {
	s.Headers = v
	return s
}

func (s *ListCopilotMessagesResponse) SetStatusCode(v int32) *ListCopilotMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCopilotMessagesResponse) SetBody(v *ListCopilotMessagesResponseBody) *ListCopilotMessagesResponse {
	s.Body = v
	return s
}

func (s *ListCopilotMessagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
