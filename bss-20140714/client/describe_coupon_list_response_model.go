// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCouponListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCouponListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCouponListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCouponListResponseBody) *DescribeCouponListResponse
	GetBody() *DescribeCouponListResponseBody
}

type DescribeCouponListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCouponListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCouponListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCouponListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCouponListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCouponListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCouponListResponse) GetBody() *DescribeCouponListResponseBody {
	return s.Body
}

func (s *DescribeCouponListResponse) SetHeaders(v map[string]*string) *DescribeCouponListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCouponListResponse) SetStatusCode(v int32) *DescribeCouponListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCouponListResponse) SetBody(v *DescribeCouponListResponseBody) *DescribeCouponListResponse {
	s.Body = v
	return s
}

func (s *DescribeCouponListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
