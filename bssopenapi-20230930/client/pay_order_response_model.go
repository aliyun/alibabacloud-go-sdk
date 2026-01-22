// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPayOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PayOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PayOrderResponse
	GetStatusCode() *int32
	SetBody(v *PayOrderResponseBody) *PayOrderResponse
	GetBody() *PayOrderResponseBody
}

type PayOrderResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PayOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s PayOrderResponse) GoString() string {
	return s.String()
}

func (s *PayOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PayOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PayOrderResponse) GetBody() *PayOrderResponseBody {
	return s.Body
}

func (s *PayOrderResponse) SetHeaders(v map[string]*string) *PayOrderResponse {
	s.Headers = v
	return s
}

func (s *PayOrderResponse) SetStatusCode(v int32) *PayOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *PayOrderResponse) SetBody(v *PayOrderResponseBody) *PayOrderResponse {
	s.Body = v
	return s
}

func (s *PayOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
