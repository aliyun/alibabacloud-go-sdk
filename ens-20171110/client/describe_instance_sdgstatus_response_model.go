// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSDGStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceSDGStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceSDGStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceSDGStatusResponseBody) *DescribeInstanceSDGStatusResponse
	GetBody() *DescribeInstanceSDGStatusResponseBody
}

type DescribeInstanceSDGStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceSDGStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceSDGStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSDGStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSDGStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceSDGStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceSDGStatusResponse) GetBody() *DescribeInstanceSDGStatusResponseBody {
	return s.Body
}

func (s *DescribeInstanceSDGStatusResponse) SetHeaders(v map[string]*string) *DescribeInstanceSDGStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSDGStatusResponse) SetStatusCode(v int32) *DescribeInstanceSDGStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceSDGStatusResponse) SetBody(v *DescribeInstanceSDGStatusResponseBody) *DescribeInstanceSDGStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceSDGStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
