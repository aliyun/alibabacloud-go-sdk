// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBVersionInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBVersionInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBVersionInfosResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBVersionInfosResponseBody) *DescribeDBVersionInfosResponse
	GetBody() *DescribeDBVersionInfosResponseBody
}

type DescribeDBVersionInfosResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBVersionInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBVersionInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBVersionInfosResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBVersionInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBVersionInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBVersionInfosResponse) GetBody() *DescribeDBVersionInfosResponseBody {
	return s.Body
}

func (s *DescribeDBVersionInfosResponse) SetHeaders(v map[string]*string) *DescribeDBVersionInfosResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBVersionInfosResponse) SetStatusCode(v int32) *DescribeDBVersionInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBVersionInfosResponse) SetBody(v *DescribeDBVersionInfosResponseBody) *DescribeDBVersionInfosResponse {
	s.Body = v
	return s
}

func (s *DescribeDBVersionInfosResponse) Validate() error {
	return dara.Validate(s)
}
