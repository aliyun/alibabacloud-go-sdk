// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatappMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChatappMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChatappMessageResponse
	GetStatusCode() *int32
	SetBody(v *ListChatappMessageResponseBody) *ListChatappMessageResponse
	GetBody() *ListChatappMessageResponseBody
}

type ListChatappMessageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChatappMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChatappMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChatappMessageResponse) GoString() string {
	return s.String()
}

func (s *ListChatappMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChatappMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChatappMessageResponse) GetBody() *ListChatappMessageResponseBody {
	return s.Body
}

func (s *ListChatappMessageResponse) SetHeaders(v map[string]*string) *ListChatappMessageResponse {
	s.Headers = v
	return s
}

func (s *ListChatappMessageResponse) SetStatusCode(v int32) *ListChatappMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChatappMessageResponse) SetBody(v *ListChatappMessageResponseBody) *ListChatappMessageResponse {
	s.Body = v
	return s
}

func (s *ListChatappMessageResponse) Validate() error {
	return dara.Validate(s)
}
