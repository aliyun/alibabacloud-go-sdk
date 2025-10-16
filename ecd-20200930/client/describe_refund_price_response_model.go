// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRefundPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRefundPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRefundPriceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRefundPriceResponseBody) *DescribeRefundPriceResponse
	GetBody() *DescribeRefundPriceResponseBody
}

type DescribeRefundPriceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRefundPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRefundPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefundPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeRefundPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRefundPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRefundPriceResponse) GetBody() *DescribeRefundPriceResponseBody {
	return s.Body
}

func (s *DescribeRefundPriceResponse) SetHeaders(v map[string]*string) *DescribeRefundPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeRefundPriceResponse) SetStatusCode(v int32) *DescribeRefundPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRefundPriceResponse) SetBody(v *DescribeRefundPriceResponseBody) *DescribeRefundPriceResponse {
	s.Body = v
	return s
}

func (s *DescribeRefundPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
