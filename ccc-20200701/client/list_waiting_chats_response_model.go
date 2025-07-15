// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWaitingChatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWaitingChatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWaitingChatsResponse
	GetStatusCode() *int32
	SetBody(v *ListWaitingChatsResponseBody) *ListWaitingChatsResponse
	GetBody() *ListWaitingChatsResponseBody
}

type ListWaitingChatsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWaitingChatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWaitingChatsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingChatsResponse) GoString() string {
	return s.String()
}

func (s *ListWaitingChatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWaitingChatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWaitingChatsResponse) GetBody() *ListWaitingChatsResponseBody {
	return s.Body
}

func (s *ListWaitingChatsResponse) SetHeaders(v map[string]*string) *ListWaitingChatsResponse {
	s.Headers = v
	return s
}

func (s *ListWaitingChatsResponse) SetStatusCode(v int32) *ListWaitingChatsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWaitingChatsResponse) SetBody(v *ListWaitingChatsResponseBody) *ListWaitingChatsResponse {
	s.Body = v
	return s
}

func (s *ListWaitingChatsResponse) Validate() error {
	return dara.Validate(s)
}
