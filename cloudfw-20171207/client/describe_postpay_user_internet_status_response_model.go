// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayUserInternetStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePostpayUserInternetStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePostpayUserInternetStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribePostpayUserInternetStatusResponseBody) *DescribePostpayUserInternetStatusResponse
	GetBody() *DescribePostpayUserInternetStatusResponseBody
}

type DescribePostpayUserInternetStatusResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePostpayUserInternetStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePostpayUserInternetStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayUserInternetStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribePostpayUserInternetStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePostpayUserInternetStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePostpayUserInternetStatusResponse) GetBody() *DescribePostpayUserInternetStatusResponseBody {
	return s.Body
}

func (s *DescribePostpayUserInternetStatusResponse) SetHeaders(v map[string]*string) *DescribePostpayUserInternetStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribePostpayUserInternetStatusResponse) SetStatusCode(v int32) *DescribePostpayUserInternetStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePostpayUserInternetStatusResponse) SetBody(v *DescribePostpayUserInternetStatusResponseBody) *DescribePostpayUserInternetStatusResponse {
	s.Body = v
	return s
}

func (s *DescribePostpayUserInternetStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
