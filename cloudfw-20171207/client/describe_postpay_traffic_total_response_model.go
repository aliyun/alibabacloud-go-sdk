// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayTrafficTotalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePostpayTrafficTotalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePostpayTrafficTotalResponse
	GetStatusCode() *int32
	SetBody(v *DescribePostpayTrafficTotalResponseBody) *DescribePostpayTrafficTotalResponse
	GetBody() *DescribePostpayTrafficTotalResponseBody
}

type DescribePostpayTrafficTotalResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePostpayTrafficTotalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePostpayTrafficTotalResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayTrafficTotalResponse) GoString() string {
	return s.String()
}

func (s *DescribePostpayTrafficTotalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePostpayTrafficTotalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePostpayTrafficTotalResponse) GetBody() *DescribePostpayTrafficTotalResponseBody {
	return s.Body
}

func (s *DescribePostpayTrafficTotalResponse) SetHeaders(v map[string]*string) *DescribePostpayTrafficTotalResponse {
	s.Headers = v
	return s
}

func (s *DescribePostpayTrafficTotalResponse) SetStatusCode(v int32) *DescribePostpayTrafficTotalResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePostpayTrafficTotalResponse) SetBody(v *DescribePostpayTrafficTotalResponseBody) *DescribePostpayTrafficTotalResponse {
	s.Body = v
	return s
}

func (s *DescribePostpayTrafficTotalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
