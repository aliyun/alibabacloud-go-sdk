// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHealthStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHealthStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHealthStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHealthStatusResponseBody) *DescribeHealthStatusResponse
	GetBody() *DescribeHealthStatusResponseBody
}

type DescribeHealthStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHealthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHealthStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHealthStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHealthStatusResponse) GetBody() *DescribeHealthStatusResponseBody {
	return s.Body
}

func (s *DescribeHealthStatusResponse) SetHeaders(v map[string]*string) *DescribeHealthStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeHealthStatusResponse) SetStatusCode(v int32) *DescribeHealthStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHealthStatusResponse) SetBody(v *DescribeHealthStatusResponseBody) *DescribeHealthStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeHealthStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
