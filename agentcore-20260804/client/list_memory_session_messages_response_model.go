// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemorySessionMessagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMemorySessionMessagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMemorySessionMessagesResponse
	GetStatusCode() *int32
	SetBody(v *ListMemorySessionMessagesResponseBody) *ListMemorySessionMessagesResponse
	GetBody() *ListMemorySessionMessagesResponseBody
}

type ListMemorySessionMessagesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMemorySessionMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMemorySessionMessagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMemorySessionMessagesResponse) GoString() string {
	return s.String()
}

func (s *ListMemorySessionMessagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMemorySessionMessagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMemorySessionMessagesResponse) GetBody() *ListMemorySessionMessagesResponseBody {
	return s.Body
}

func (s *ListMemorySessionMessagesResponse) SetHeaders(v map[string]*string) *ListMemorySessionMessagesResponse {
	s.Headers = v
	return s
}

func (s *ListMemorySessionMessagesResponse) SetStatusCode(v int32) *ListMemorySessionMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMemorySessionMessagesResponse) SetBody(v *ListMemorySessionMessagesResponseBody) *ListMemorySessionMessagesResponse {
	s.Body = v
	return s
}

func (s *ListMemorySessionMessagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
