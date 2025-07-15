// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventSubResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEventSubResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEventSubResponse
	GetStatusCode() *int32
	SetBody(v *ListEventSubResponseBody) *ListEventSubResponse
	GetBody() *ListEventSubResponseBody
}

type ListEventSubResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEventSubResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEventSubResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEventSubResponse) GoString() string {
	return s.String()
}

func (s *ListEventSubResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEventSubResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEventSubResponse) GetBody() *ListEventSubResponseBody {
	return s.Body
}

func (s *ListEventSubResponse) SetHeaders(v map[string]*string) *ListEventSubResponse {
	s.Headers = v
	return s
}

func (s *ListEventSubResponse) SetStatusCode(v int32) *ListEventSubResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEventSubResponse) SetBody(v *ListEventSubResponseBody) *ListEventSubResponse {
	s.Body = v
	return s
}

func (s *ListEventSubResponse) Validate() error {
	return dara.Validate(s)
}
