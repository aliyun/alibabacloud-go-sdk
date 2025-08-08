// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReceiversResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListReceiversResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListReceiversResponse
	GetStatusCode() *int32
	SetBody(v *ListReceiversResponseBody) *ListReceiversResponse
	GetBody() *ListReceiversResponseBody
}

type ListReceiversResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListReceiversResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListReceiversResponse) String() string {
	return dara.Prettify(s)
}

func (s ListReceiversResponse) GoString() string {
	return s.String()
}

func (s *ListReceiversResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListReceiversResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListReceiversResponse) GetBody() *ListReceiversResponseBody {
	return s.Body
}

func (s *ListReceiversResponse) SetHeaders(v map[string]*string) *ListReceiversResponse {
	s.Headers = v
	return s
}

func (s *ListReceiversResponse) SetStatusCode(v int32) *ListReceiversResponse {
	s.StatusCode = &v
	return s
}

func (s *ListReceiversResponse) SetBody(v *ListReceiversResponseBody) *ListReceiversResponse {
	s.Body = v
	return s
}

func (s *ListReceiversResponse) Validate() error {
	return dara.Validate(s)
}
