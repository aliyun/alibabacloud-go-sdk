// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterStatusResponseBody) *DescribeDBClusterStatusResponse
	GetBody() *DescribeDBClusterStatusResponseBody
}

type DescribeDBClusterStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterStatusResponse) GetBody() *DescribeDBClusterStatusResponseBody {
	return s.Body
}

func (s *DescribeDBClusterStatusResponse) SetHeaders(v map[string]*string) *DescribeDBClusterStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterStatusResponse) SetStatusCode(v int32) *DescribeDBClusterStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterStatusResponse) SetBody(v *DescribeDBClusterStatusResponseBody) *DescribeDBClusterStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterStatusResponse) Validate() error {
	return dara.Validate(s)
}
