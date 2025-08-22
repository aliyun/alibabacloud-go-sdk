// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDcdnKvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDcdnKvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDcdnKvResponse
	GetStatusCode() *int32
	SetBody(v *ListDcdnKvResponseBody) *ListDcdnKvResponse
	GetBody() *ListDcdnKvResponseBody
}

type ListDcdnKvResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDcdnKvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDcdnKvResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDcdnKvResponse) GoString() string {
	return s.String()
}

func (s *ListDcdnKvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDcdnKvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDcdnKvResponse) GetBody() *ListDcdnKvResponseBody {
	return s.Body
}

func (s *ListDcdnKvResponse) SetHeaders(v map[string]*string) *ListDcdnKvResponse {
	s.Headers = v
	return s
}

func (s *ListDcdnKvResponse) SetStatusCode(v int32) *ListDcdnKvResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDcdnKvResponse) SetBody(v *ListDcdnKvResponseBody) *ListDcdnKvResponse {
	s.Body = v
	return s
}

func (s *ListDcdnKvResponse) Validate() error {
	return dara.Validate(s)
}
