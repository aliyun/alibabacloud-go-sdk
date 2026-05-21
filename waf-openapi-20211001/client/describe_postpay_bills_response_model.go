// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayBillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePostpayBillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePostpayBillsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePostpayBillsResponseBody) *DescribePostpayBillsResponse
	GetBody() *DescribePostpayBillsResponseBody
}

type DescribePostpayBillsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePostpayBillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePostpayBillsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayBillsResponse) GoString() string {
	return s.String()
}

func (s *DescribePostpayBillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePostpayBillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePostpayBillsResponse) GetBody() *DescribePostpayBillsResponseBody {
	return s.Body
}

func (s *DescribePostpayBillsResponse) SetHeaders(v map[string]*string) *DescribePostpayBillsResponse {
	s.Headers = v
	return s
}

func (s *DescribePostpayBillsResponse) SetStatusCode(v int32) *DescribePostpayBillsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePostpayBillsResponse) SetBody(v *DescribePostpayBillsResponseBody) *DescribePostpayBillsResponse {
	s.Body = v
	return s
}

func (s *DescribePostpayBillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
