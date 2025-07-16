// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSpotDiscountHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSpotDiscountHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSpotDiscountHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSpotDiscountHistoryResponseBody) *DescribeSpotDiscountHistoryResponse
	GetBody() *DescribeSpotDiscountHistoryResponseBody
}

type DescribeSpotDiscountHistoryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSpotDiscountHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSpotDiscountHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpotDiscountHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeSpotDiscountHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSpotDiscountHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSpotDiscountHistoryResponse) GetBody() *DescribeSpotDiscountHistoryResponseBody {
	return s.Body
}

func (s *DescribeSpotDiscountHistoryResponse) SetHeaders(v map[string]*string) *DescribeSpotDiscountHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeSpotDiscountHistoryResponse) SetStatusCode(v int32) *DescribeSpotDiscountHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSpotDiscountHistoryResponse) SetBody(v *DescribeSpotDiscountHistoryResponseBody) *DescribeSpotDiscountHistoryResponse {
	s.Body = v
	return s
}

func (s *DescribeSpotDiscountHistoryResponse) Validate() error {
	return dara.Validate(s)
}
