// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListShareReceiversResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListShareReceiversResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListShareReceiversResponse
	GetStatusCode() *int32
	SetBody(v *ListShareReceiversResponseBody) *ListShareReceiversResponse
	GetBody() *ListShareReceiversResponseBody
}

type ListShareReceiversResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListShareReceiversResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListShareReceiversResponse) String() string {
	return dara.Prettify(s)
}

func (s ListShareReceiversResponse) GoString() string {
	return s.String()
}

func (s *ListShareReceiversResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListShareReceiversResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListShareReceiversResponse) GetBody() *ListShareReceiversResponseBody {
	return s.Body
}

func (s *ListShareReceiversResponse) SetHeaders(v map[string]*string) *ListShareReceiversResponse {
	s.Headers = v
	return s
}

func (s *ListShareReceiversResponse) SetStatusCode(v int32) *ListShareReceiversResponse {
	s.StatusCode = &v
	return s
}

func (s *ListShareReceiversResponse) SetBody(v *ListShareReceiversResponseBody) *ListShareReceiversResponse {
	s.Body = v
	return s
}

func (s *ListShareReceiversResponse) Validate() error {
	return dara.Validate(s)
}
