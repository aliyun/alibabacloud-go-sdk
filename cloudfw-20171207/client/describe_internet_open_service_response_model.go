// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetOpenServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInternetOpenServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInternetOpenServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInternetOpenServiceResponseBody) *DescribeInternetOpenServiceResponse
	GetBody() *DescribeInternetOpenServiceResponseBody
}

type DescribeInternetOpenServiceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInternetOpenServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInternetOpenServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInternetOpenServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInternetOpenServiceResponse) GetBody() *DescribeInternetOpenServiceResponseBody {
	return s.Body
}

func (s *DescribeInternetOpenServiceResponse) SetHeaders(v map[string]*string) *DescribeInternetOpenServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInternetOpenServiceResponse) SetStatusCode(v int32) *DescribeInternetOpenServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInternetOpenServiceResponse) SetBody(v *DescribeInternetOpenServiceResponseBody) *DescribeInternetOpenServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeInternetOpenServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
