// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrderResponseBody) *CreateOrderResponse
	GetBody() *CreateOrderResponseBody
}

type CreateOrderResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrderResponse) GetBody() *CreateOrderResponseBody {
	return s.Body
}

func (s *CreateOrderResponse) SetHeaders(v map[string]*string) *CreateOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateOrderResponse) SetStatusCode(v int32) *CreateOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrderResponse) SetBody(v *CreateOrderResponseBody) *CreateOrderResponse {
	s.Body = v
	return s
}

func (s *CreateOrderResponse) Validate() error {
	return dara.Validate(s)
}
