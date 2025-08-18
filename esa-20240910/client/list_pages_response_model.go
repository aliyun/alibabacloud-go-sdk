// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPagesResponse
	GetStatusCode() *int32
	SetBody(v *ListPagesResponseBody) *ListPagesResponse
	GetBody() *ListPagesResponseBody
}

type ListPagesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPagesResponse) GoString() string {
	return s.String()
}

func (s *ListPagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPagesResponse) GetBody() *ListPagesResponseBody {
	return s.Body
}

func (s *ListPagesResponse) SetHeaders(v map[string]*string) *ListPagesResponse {
	s.Headers = v
	return s
}

func (s *ListPagesResponse) SetStatusCode(v int32) *ListPagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPagesResponse) SetBody(v *ListPagesResponseBody) *ListPagesResponse {
	s.Body = v
	return s
}

func (s *ListPagesResponse) Validate() error {
	return dara.Validate(s)
}
