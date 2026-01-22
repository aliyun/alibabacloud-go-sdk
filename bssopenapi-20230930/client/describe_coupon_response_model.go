// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCouponResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCouponResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCouponResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCouponResponseBody) *DescribeCouponResponse
	GetBody() *DescribeCouponResponseBody
}

type DescribeCouponResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCouponResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCouponResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCouponResponse) GoString() string {
	return s.String()
}

func (s *DescribeCouponResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCouponResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCouponResponse) GetBody() *DescribeCouponResponseBody {
	return s.Body
}

func (s *DescribeCouponResponse) SetHeaders(v map[string]*string) *DescribeCouponResponse {
	s.Headers = v
	return s
}

func (s *DescribeCouponResponse) SetStatusCode(v int32) *DescribeCouponResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCouponResponse) SetBody(v *DescribeCouponResponseBody) *DescribeCouponResponse {
	s.Body = v
	return s
}

func (s *DescribeCouponResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
