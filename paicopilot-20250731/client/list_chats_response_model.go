// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChatsResponse
	GetStatusCode() *int32
	SetBody(v *ListChatsResponseBody) *ListChatsResponse
	GetBody() *ListChatsResponseBody
}

type ListChatsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChatsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChatsResponse) GoString() string {
	return s.String()
}

func (s *ListChatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChatsResponse) GetBody() *ListChatsResponseBody {
	return s.Body
}

func (s *ListChatsResponse) SetHeaders(v map[string]*string) *ListChatsResponse {
	s.Headers = v
	return s
}

func (s *ListChatsResponse) SetStatusCode(v int32) *ListChatsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChatsResponse) SetBody(v *ListChatsResponseBody) *ListChatsResponse {
	s.Body = v
	return s
}

func (s *ListChatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
