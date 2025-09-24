// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrdersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryOrdersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryOrdersResponse
	GetStatusCode() *int32
	SetBody(v *QueryOrdersResponseBody) *QueryOrdersResponse
	GetBody() *QueryOrdersResponseBody
}

type QueryOrdersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrdersResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryOrdersResponse) GoString() string {
	return s.String()
}

func (s *QueryOrdersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryOrdersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryOrdersResponse) GetBody() *QueryOrdersResponseBody {
	return s.Body
}

func (s *QueryOrdersResponse) SetHeaders(v map[string]*string) *QueryOrdersResponse {
	s.Headers = v
	return s
}

func (s *QueryOrdersResponse) SetStatusCode(v int32) *QueryOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrdersResponse) SetBody(v *QueryOrdersResponseBody) *QueryOrdersResponse {
	s.Body = v
	return s
}

func (s *QueryOrdersResponse) Validate() error {
	return dara.Validate(s)
}
