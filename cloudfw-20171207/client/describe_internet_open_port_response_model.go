// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetOpenPortResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInternetOpenPortResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInternetOpenPortResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInternetOpenPortResponseBody) *DescribeInternetOpenPortResponse
	GetBody() *DescribeInternetOpenPortResponseBody
}

type DescribeInternetOpenPortResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInternetOpenPortResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInternetOpenPortResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenPortResponse) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenPortResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInternetOpenPortResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInternetOpenPortResponse) GetBody() *DescribeInternetOpenPortResponseBody {
	return s.Body
}

func (s *DescribeInternetOpenPortResponse) SetHeaders(v map[string]*string) *DescribeInternetOpenPortResponse {
	s.Headers = v
	return s
}

func (s *DescribeInternetOpenPortResponse) SetStatusCode(v int32) *DescribeInternetOpenPortResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInternetOpenPortResponse) SetBody(v *DescribeInternetOpenPortResponseBody) *DescribeInternetOpenPortResponse {
	s.Body = v
	return s
}

func (s *DescribeInternetOpenPortResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
