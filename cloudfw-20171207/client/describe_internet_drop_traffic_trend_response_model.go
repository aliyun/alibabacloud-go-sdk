// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetDropTrafficTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInternetDropTrafficTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInternetDropTrafficTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInternetDropTrafficTrendResponseBody) *DescribeInternetDropTrafficTrendResponse
	GetBody() *DescribeInternetDropTrafficTrendResponseBody
}

type DescribeInternetDropTrafficTrendResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInternetDropTrafficTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInternetDropTrafficTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetDropTrafficTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeInternetDropTrafficTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInternetDropTrafficTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInternetDropTrafficTrendResponse) GetBody() *DescribeInternetDropTrafficTrendResponseBody {
	return s.Body
}

func (s *DescribeInternetDropTrafficTrendResponse) SetHeaders(v map[string]*string) *DescribeInternetDropTrafficTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeInternetDropTrafficTrendResponse) SetStatusCode(v int32) *DescribeInternetDropTrafficTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendResponse) SetBody(v *DescribeInternetDropTrafficTrendResponseBody) *DescribeInternetDropTrafficTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeInternetDropTrafficTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
