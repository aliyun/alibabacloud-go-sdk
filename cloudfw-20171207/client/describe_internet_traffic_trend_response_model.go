// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetTrafficTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInternetTrafficTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInternetTrafficTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInternetTrafficTrendResponseBody) *DescribeInternetTrafficTrendResponse
	GetBody() *DescribeInternetTrafficTrendResponseBody
}

type DescribeInternetTrafficTrendResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInternetTrafficTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInternetTrafficTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetTrafficTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeInternetTrafficTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInternetTrafficTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInternetTrafficTrendResponse) GetBody() *DescribeInternetTrafficTrendResponseBody {
	return s.Body
}

func (s *DescribeInternetTrafficTrendResponse) SetHeaders(v map[string]*string) *DescribeInternetTrafficTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeInternetTrafficTrendResponse) SetStatusCode(v int32) *DescribeInternetTrafficTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponse) SetBody(v *DescribeInternetTrafficTrendResponseBody) *DescribeInternetTrafficTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeInternetTrafficTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
