// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountAccessIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAccountAccessIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAccountAccessIdResponse
	GetStatusCode() *int32
	SetBody(v *ListAccountAccessIdResponseBody) *ListAccountAccessIdResponse
	GetBody() *ListAccountAccessIdResponseBody
}

type ListAccountAccessIdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccountAccessIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccountAccessIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAccountAccessIdResponse) GoString() string {
	return s.String()
}

func (s *ListAccountAccessIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAccountAccessIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAccountAccessIdResponse) GetBody() *ListAccountAccessIdResponseBody {
	return s.Body
}

func (s *ListAccountAccessIdResponse) SetHeaders(v map[string]*string) *ListAccountAccessIdResponse {
	s.Headers = v
	return s
}

func (s *ListAccountAccessIdResponse) SetStatusCode(v int32) *ListAccountAccessIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccountAccessIdResponse) SetBody(v *ListAccountAccessIdResponseBody) *ListAccountAccessIdResponse {
	s.Body = v
	return s
}

func (s *ListAccountAccessIdResponse) Validate() error {
	return dara.Validate(s)
}
