// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMessagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMessagesResponse
	GetStatusCode() *int32
	SetBody(v *ListMessagesResponseBody) *ListMessagesResponse
	GetBody() *ListMessagesResponseBody
}

type ListMessagesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMessagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMessagesResponse) GoString() string {
	return s.String()
}

func (s *ListMessagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMessagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMessagesResponse) GetBody() *ListMessagesResponseBody {
	return s.Body
}

func (s *ListMessagesResponse) SetHeaders(v map[string]*string) *ListMessagesResponse {
	s.Headers = v
	return s
}

func (s *ListMessagesResponse) SetStatusCode(v int32) *ListMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMessagesResponse) SetBody(v *ListMessagesResponseBody) *ListMessagesResponse {
	s.Body = v
	return s
}

func (s *ListMessagesResponse) Validate() error {
	return dara.Validate(s)
}
