// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckoutWithAKResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckoutWithAKResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckoutWithAKResponse
	GetStatusCode() *int32
	SetBody(v *CheckoutWithAKResponseBody) *CheckoutWithAKResponse
	GetBody() *CheckoutWithAKResponseBody
}

type CheckoutWithAKResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckoutWithAKResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckoutWithAKResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckoutWithAKResponse) GoString() string {
	return s.String()
}

func (s *CheckoutWithAKResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckoutWithAKResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckoutWithAKResponse) GetBody() *CheckoutWithAKResponseBody {
	return s.Body
}

func (s *CheckoutWithAKResponse) SetHeaders(v map[string]*string) *CheckoutWithAKResponse {
	s.Headers = v
	return s
}

func (s *CheckoutWithAKResponse) SetStatusCode(v int32) *CheckoutWithAKResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckoutWithAKResponse) SetBody(v *CheckoutWithAKResponseBody) *CheckoutWithAKResponse {
	s.Body = v
	return s
}

func (s *CheckoutWithAKResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
