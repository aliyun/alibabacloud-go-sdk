// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayUserVpcStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePostpayUserVpcStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePostpayUserVpcStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribePostpayUserVpcStatusResponseBody) *DescribePostpayUserVpcStatusResponse
	GetBody() *DescribePostpayUserVpcStatusResponseBody
}

type DescribePostpayUserVpcStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePostpayUserVpcStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePostpayUserVpcStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayUserVpcStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribePostpayUserVpcStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePostpayUserVpcStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePostpayUserVpcStatusResponse) GetBody() *DescribePostpayUserVpcStatusResponseBody {
	return s.Body
}

func (s *DescribePostpayUserVpcStatusResponse) SetHeaders(v map[string]*string) *DescribePostpayUserVpcStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribePostpayUserVpcStatusResponse) SetStatusCode(v int32) *DescribePostpayUserVpcStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePostpayUserVpcStatusResponse) SetBody(v *DescribePostpayUserVpcStatusResponseBody) *DescribePostpayUserVpcStatusResponse {
	s.Body = v
	return s
}

func (s *DescribePostpayUserVpcStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
