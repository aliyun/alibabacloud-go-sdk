// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrdersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOrdersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOrdersResponse
	GetStatusCode() *int32
	SetBody(v *GetOrdersResponseBody) *GetOrdersResponse
	GetBody() *GetOrdersResponseBody
}

type GetOrdersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrdersResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOrdersResponse) GoString() string {
	return s.String()
}

func (s *GetOrdersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOrdersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOrdersResponse) GetBody() *GetOrdersResponseBody {
	return s.Body
}

func (s *GetOrdersResponse) SetHeaders(v map[string]*string) *GetOrdersResponse {
	s.Headers = v
	return s
}

func (s *GetOrdersResponse) SetStatusCode(v int32) *GetOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrdersResponse) SetBody(v *GetOrdersResponseBody) *GetOrdersResponse {
	s.Body = v
	return s
}

func (s *GetOrdersResponse) Validate() error {
	return dara.Validate(s)
}
