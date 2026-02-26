// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOrderResponse
	GetStatusCode() *int32
	SetBody(v *OrderResult) *GetOrderResponse
	GetBody() *OrderResult
}

type GetOrderResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OrderResult       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOrderResponse) GoString() string {
	return s.String()
}

func (s *GetOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOrderResponse) GetBody() *OrderResult {
	return s.Body
}

func (s *GetOrderResponse) SetHeaders(v map[string]*string) *GetOrderResponse {
	s.Headers = v
	return s
}

func (s *GetOrderResponse) SetStatusCode(v int32) *GetOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrderResponse) SetBody(v *OrderResult) *GetOrderResponse {
	s.Body = v
	return s
}

func (s *GetOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
