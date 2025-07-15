// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrePayOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePrePayOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePrePayOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreatePrePayOrderResponseBody) *CreatePrePayOrderResponse
	GetBody() *CreatePrePayOrderResponseBody
}

type CreatePrePayOrderResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrePayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrePayOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePrePayOrderResponse) GoString() string {
	return s.String()
}

func (s *CreatePrePayOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePrePayOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePrePayOrderResponse) GetBody() *CreatePrePayOrderResponseBody {
	return s.Body
}

func (s *CreatePrePayOrderResponse) SetHeaders(v map[string]*string) *CreatePrePayOrderResponse {
	s.Headers = v
	return s
}

func (s *CreatePrePayOrderResponse) SetStatusCode(v int32) *CreatePrePayOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrePayOrderResponse) SetBody(v *CreatePrePayOrderResponseBody) *CreatePrePayOrderResponse {
	s.Body = v
	return s
}

func (s *CreatePrePayOrderResponse) Validate() error {
	return dara.Validate(s)
}
