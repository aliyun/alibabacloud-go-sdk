// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCouponItemListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCouponItemListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCouponItemListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCouponItemListResponseBody) *DescribeCouponItemListResponse
	GetBody() *DescribeCouponItemListResponseBody
}

type DescribeCouponItemListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCouponItemListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCouponItemListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCouponItemListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCouponItemListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCouponItemListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCouponItemListResponse) GetBody() *DescribeCouponItemListResponseBody {
	return s.Body
}

func (s *DescribeCouponItemListResponse) SetHeaders(v map[string]*string) *DescribeCouponItemListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCouponItemListResponse) SetStatusCode(v int32) *DescribeCouponItemListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCouponItemListResponse) SetBody(v *DescribeCouponItemListResponseBody) *DescribeCouponItemListResponse {
	s.Body = v
	return s
}

func (s *DescribeCouponItemListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
