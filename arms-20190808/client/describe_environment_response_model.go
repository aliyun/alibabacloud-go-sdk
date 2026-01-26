// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnvironmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnvironmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnvironmentResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnvironmentResponseBody) *DescribeEnvironmentResponse
	GetBody() *DescribeEnvironmentResponseBody
}

type DescribeEnvironmentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnvironmentResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnvironmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnvironmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnvironmentResponse) GetBody() *DescribeEnvironmentResponseBody {
	return s.Body
}

func (s *DescribeEnvironmentResponse) SetHeaders(v map[string]*string) *DescribeEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnvironmentResponse) SetStatusCode(v int32) *DescribeEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnvironmentResponse) SetBody(v *DescribeEnvironmentResponseBody) *DescribeEnvironmentResponse {
	s.Body = v
	return s
}

func (s *DescribeEnvironmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
