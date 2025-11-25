// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayEnabledProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePostpayEnabledProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePostpayEnabledProtectionResponse
	GetStatusCode() *int32
	SetBody(v *DescribePostpayEnabledProtectionResponseBody) *DescribePostpayEnabledProtectionResponse
	GetBody() *DescribePostpayEnabledProtectionResponseBody
}

type DescribePostpayEnabledProtectionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePostpayEnabledProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePostpayEnabledProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayEnabledProtectionResponse) GoString() string {
	return s.String()
}

func (s *DescribePostpayEnabledProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePostpayEnabledProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePostpayEnabledProtectionResponse) GetBody() *DescribePostpayEnabledProtectionResponseBody {
	return s.Body
}

func (s *DescribePostpayEnabledProtectionResponse) SetHeaders(v map[string]*string) *DescribePostpayEnabledProtectionResponse {
	s.Headers = v
	return s
}

func (s *DescribePostpayEnabledProtectionResponse) SetStatusCode(v int32) *DescribePostpayEnabledProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePostpayEnabledProtectionResponse) SetBody(v *DescribePostpayEnabledProtectionResponseBody) *DescribePostpayEnabledProtectionResponse {
	s.Body = v
	return s
}

func (s *DescribePostpayEnabledProtectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
