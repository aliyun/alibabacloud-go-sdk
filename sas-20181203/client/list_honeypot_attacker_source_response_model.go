// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotAttackerSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHoneypotAttackerSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHoneypotAttackerSourceResponse
	GetStatusCode() *int32
	SetBody(v *ListHoneypotAttackerSourceResponseBody) *ListHoneypotAttackerSourceResponse
	GetBody() *ListHoneypotAttackerSourceResponseBody
}

type ListHoneypotAttackerSourceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHoneypotAttackerSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHoneypotAttackerSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotAttackerSourceResponse) GoString() string {
	return s.String()
}

func (s *ListHoneypotAttackerSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHoneypotAttackerSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHoneypotAttackerSourceResponse) GetBody() *ListHoneypotAttackerSourceResponseBody {
	return s.Body
}

func (s *ListHoneypotAttackerSourceResponse) SetHeaders(v map[string]*string) *ListHoneypotAttackerSourceResponse {
	s.Headers = v
	return s
}

func (s *ListHoneypotAttackerSourceResponse) SetStatusCode(v int32) *ListHoneypotAttackerSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHoneypotAttackerSourceResponse) SetBody(v *ListHoneypotAttackerSourceResponseBody) *ListHoneypotAttackerSourceResponse {
	s.Body = v
	return s
}

func (s *ListHoneypotAttackerSourceResponse) Validate() error {
	return dara.Validate(s)
}
