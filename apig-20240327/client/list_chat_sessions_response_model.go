// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatSessionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChatSessionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChatSessionsResponse
	GetStatusCode() *int32
	SetBody(v *ListChatSessionsResponseBody) *ListChatSessionsResponse
	GetBody() *ListChatSessionsResponseBody
}

type ListChatSessionsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChatSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChatSessionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChatSessionsResponse) GoString() string {
	return s.String()
}

func (s *ListChatSessionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChatSessionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChatSessionsResponse) GetBody() *ListChatSessionsResponseBody {
	return s.Body
}

func (s *ListChatSessionsResponse) SetHeaders(v map[string]*string) *ListChatSessionsResponse {
	s.Headers = v
	return s
}

func (s *ListChatSessionsResponse) SetStatusCode(v int32) *ListChatSessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChatSessionsResponse) SetBody(v *ListChatSessionsResponseBody) *ListChatSessionsResponse {
	s.Body = v
	return s
}

func (s *ListChatSessionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
