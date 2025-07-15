// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomTextResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomTextResponseBody) *ListCustomTextResponse
	GetBody() *ListCustomTextResponseBody
}

type ListCustomTextResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomTextResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomTextResponse) GoString() string {
	return s.String()
}

func (s *ListCustomTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomTextResponse) GetBody() *ListCustomTextResponseBody {
	return s.Body
}

func (s *ListCustomTextResponse) SetHeaders(v map[string]*string) *ListCustomTextResponse {
	s.Headers = v
	return s
}

func (s *ListCustomTextResponse) SetStatusCode(v int32) *ListCustomTextResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomTextResponse) SetBody(v *ListCustomTextResponseBody) *ListCustomTextResponse {
	s.Body = v
	return s
}

func (s *ListCustomTextResponse) Validate() error {
	return dara.Validate(s)
}
