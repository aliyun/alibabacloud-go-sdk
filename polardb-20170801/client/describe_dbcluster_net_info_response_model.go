// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterNetInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterNetInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterNetInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterNetInfoResponseBody) *DescribeDBClusterNetInfoResponse
	GetBody() *DescribeDBClusterNetInfoResponseBody
}

type DescribeDBClusterNetInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterNetInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterNetInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNetInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterNetInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterNetInfoResponse) GetBody() *DescribeDBClusterNetInfoResponseBody {
	return s.Body
}

func (s *DescribeDBClusterNetInfoResponse) SetHeaders(v map[string]*string) *DescribeDBClusterNetInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterNetInfoResponse) SetStatusCode(v int32) *DescribeDBClusterNetInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponse) SetBody(v *DescribeDBClusterNetInfoResponseBody) *DescribeDBClusterNetInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterNetInfoResponse) Validate() error {
	return dara.Validate(s)
}
