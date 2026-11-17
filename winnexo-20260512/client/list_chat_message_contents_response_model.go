// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatMessageContentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChatMessageContentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChatMessageContentsResponse
	GetStatusCode() *int32
	SetBody(v *ListChatMessageContentsResponseBody) *ListChatMessageContentsResponse
	GetBody() *ListChatMessageContentsResponseBody
}

type ListChatMessageContentsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChatMessageContentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChatMessageContentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChatMessageContentsResponse) GoString() string {
	return s.String()
}

func (s *ListChatMessageContentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChatMessageContentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChatMessageContentsResponse) GetBody() *ListChatMessageContentsResponseBody {
	return s.Body
}

func (s *ListChatMessageContentsResponse) SetHeaders(v map[string]*string) *ListChatMessageContentsResponse {
	s.Headers = v
	return s
}

func (s *ListChatMessageContentsResponse) SetStatusCode(v int32) *ListChatMessageContentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChatMessageContentsResponse) SetBody(v *ListChatMessageContentsResponseBody) *ListChatMessageContentsResponse {
	s.Body = v
	return s
}

func (s *ListChatMessageContentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
