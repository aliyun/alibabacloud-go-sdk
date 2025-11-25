// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetServiceNameListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInternetServiceNameListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInternetServiceNameListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInternetServiceNameListResponseBody) *DescribeInternetServiceNameListResponse
	GetBody() *DescribeInternetServiceNameListResponseBody
}

type DescribeInternetServiceNameListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInternetServiceNameListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInternetServiceNameListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetServiceNameListResponse) GoString() string {
	return s.String()
}

func (s *DescribeInternetServiceNameListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInternetServiceNameListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInternetServiceNameListResponse) GetBody() *DescribeInternetServiceNameListResponseBody {
	return s.Body
}

func (s *DescribeInternetServiceNameListResponse) SetHeaders(v map[string]*string) *DescribeInternetServiceNameListResponse {
	s.Headers = v
	return s
}

func (s *DescribeInternetServiceNameListResponse) SetStatusCode(v int32) *DescribeInternetServiceNameListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInternetServiceNameListResponse) SetBody(v *DescribeInternetServiceNameListResponseBody) *DescribeInternetServiceNameListResponse {
	s.Body = v
	return s
}

func (s *DescribeInternetServiceNameListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
