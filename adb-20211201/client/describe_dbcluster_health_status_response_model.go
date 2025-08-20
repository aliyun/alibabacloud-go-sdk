// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterHealthStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterHealthStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterHealthStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterHealthStatusResponseBody) *DescribeDBClusterHealthStatusResponse
	GetBody() *DescribeDBClusterHealthStatusResponseBody
}

type DescribeDBClusterHealthStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterHealthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterHealthStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterHealthStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterHealthStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterHealthStatusResponse) GetBody() *DescribeDBClusterHealthStatusResponseBody {
	return s.Body
}

func (s *DescribeDBClusterHealthStatusResponse) SetHeaders(v map[string]*string) *DescribeDBClusterHealthStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterHealthStatusResponse) SetStatusCode(v int32) *DescribeDBClusterHealthStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponse) SetBody(v *DescribeDBClusterHealthStatusResponseBody) *DescribeDBClusterHealthStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterHealthStatusResponse) Validate() error {
	return dara.Validate(s)
}
