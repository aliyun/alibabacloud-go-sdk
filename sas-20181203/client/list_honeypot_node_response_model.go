// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHoneypotNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHoneypotNodeResponse
	GetStatusCode() *int32
	SetBody(v *ListHoneypotNodeResponseBody) *ListHoneypotNodeResponse
	GetBody() *ListHoneypotNodeResponseBody
}

type ListHoneypotNodeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHoneypotNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHoneypotNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotNodeResponse) GoString() string {
	return s.String()
}

func (s *ListHoneypotNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHoneypotNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHoneypotNodeResponse) GetBody() *ListHoneypotNodeResponseBody {
	return s.Body
}

func (s *ListHoneypotNodeResponse) SetHeaders(v map[string]*string) *ListHoneypotNodeResponse {
	s.Headers = v
	return s
}

func (s *ListHoneypotNodeResponse) SetStatusCode(v int32) *ListHoneypotNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHoneypotNodeResponse) SetBody(v *ListHoneypotNodeResponseBody) *ListHoneypotNodeResponse {
	s.Body = v
	return s
}

func (s *ListHoneypotNodeResponse) Validate() error {
	return dara.Validate(s)
}
