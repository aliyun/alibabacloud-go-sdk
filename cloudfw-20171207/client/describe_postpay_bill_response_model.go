// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayBillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePostpayBillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePostpayBillResponse
	GetStatusCode() *int32
	SetBody(v *DescribePostpayBillResponseBody) *DescribePostpayBillResponse
	GetBody() *DescribePostpayBillResponseBody
}

type DescribePostpayBillResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePostpayBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePostpayBillResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayBillResponse) GoString() string {
	return s.String()
}

func (s *DescribePostpayBillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePostpayBillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePostpayBillResponse) GetBody() *DescribePostpayBillResponseBody {
	return s.Body
}

func (s *DescribePostpayBillResponse) SetHeaders(v map[string]*string) *DescribePostpayBillResponse {
	s.Headers = v
	return s
}

func (s *DescribePostpayBillResponse) SetStatusCode(v int32) *DescribePostpayBillResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePostpayBillResponse) SetBody(v *DescribePostpayBillResponseBody) *DescribePostpayBillResponse {
	s.Body = v
	return s
}

func (s *DescribePostpayBillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
