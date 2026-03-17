// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHealthCheckAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHealthCheckAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHealthCheckAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHealthCheckAttributeResponseBody) *DescribeHealthCheckAttributeResponse
	GetBody() *DescribeHealthCheckAttributeResponseBody
}

type DescribeHealthCheckAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHealthCheckAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHealthCheckAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthCheckAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHealthCheckAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHealthCheckAttributeResponse) GetBody() *DescribeHealthCheckAttributeResponseBody {
	return s.Body
}

func (s *DescribeHealthCheckAttributeResponse) SetHeaders(v map[string]*string) *DescribeHealthCheckAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeHealthCheckAttributeResponse) SetStatusCode(v int32) *DescribeHealthCheckAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHealthCheckAttributeResponse) SetBody(v *DescribeHealthCheckAttributeResponseBody) *DescribeHealthCheckAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeHealthCheckAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
