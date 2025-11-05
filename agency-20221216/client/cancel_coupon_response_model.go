// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCouponResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelCouponResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelCouponResponse
	GetStatusCode() *int32
	SetBody(v *CancelCouponResponseBody) *CancelCouponResponse
	GetBody() *CancelCouponResponseBody
}

type CancelCouponResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelCouponResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelCouponResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelCouponResponse) GoString() string {
	return s.String()
}

func (s *CancelCouponResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelCouponResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelCouponResponse) GetBody() *CancelCouponResponseBody {
	return s.Body
}

func (s *CancelCouponResponse) SetHeaders(v map[string]*string) *CancelCouponResponse {
	s.Headers = v
	return s
}

func (s *CancelCouponResponse) SetStatusCode(v int32) *CancelCouponResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelCouponResponse) SetBody(v *CancelCouponResponseBody) *CancelCouponResponse {
	s.Body = v
	return s
}

func (s *CancelCouponResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
