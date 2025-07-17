// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRefreshesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceRefreshesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceRefreshesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceRefreshesResponseBody) *DescribeInstanceRefreshesResponse
	GetBody() *DescribeInstanceRefreshesResponseBody
}

type DescribeInstanceRefreshesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceRefreshesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceRefreshesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRefreshesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRefreshesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceRefreshesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceRefreshesResponse) GetBody() *DescribeInstanceRefreshesResponseBody {
	return s.Body
}

func (s *DescribeInstanceRefreshesResponse) SetHeaders(v map[string]*string) *DescribeInstanceRefreshesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceRefreshesResponse) SetStatusCode(v int32) *DescribeInstanceRefreshesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceRefreshesResponse) SetBody(v *DescribeInstanceRefreshesResponseBody) *DescribeInstanceRefreshesResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceRefreshesResponse) Validate() error {
	return dara.Validate(s)
}
