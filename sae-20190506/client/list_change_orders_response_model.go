// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChangeOrdersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChangeOrdersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChangeOrdersResponse
	GetStatusCode() *int32
	SetBody(v *ListChangeOrdersResponseBody) *ListChangeOrdersResponse
	GetBody() *ListChangeOrdersResponseBody
}

type ListChangeOrdersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChangeOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChangeOrdersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChangeOrdersResponse) GoString() string {
	return s.String()
}

func (s *ListChangeOrdersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChangeOrdersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChangeOrdersResponse) GetBody() *ListChangeOrdersResponseBody {
	return s.Body
}

func (s *ListChangeOrdersResponse) SetHeaders(v map[string]*string) *ListChangeOrdersResponse {
	s.Headers = v
	return s
}

func (s *ListChangeOrdersResponse) SetStatusCode(v int32) *ListChangeOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChangeOrdersResponse) SetBody(v *ListChangeOrdersResponseBody) *ListChangeOrdersResponse {
	s.Body = v
	return s
}

func (s *ListChangeOrdersResponse) Validate() error {
	return dara.Validate(s)
}
