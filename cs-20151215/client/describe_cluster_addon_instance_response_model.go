// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterAddonInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterAddonInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterAddonInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterAddonInstanceResponseBody) *DescribeClusterAddonInstanceResponse
	GetBody() *DescribeClusterAddonInstanceResponseBody
}

type DescribeClusterAddonInstanceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterAddonInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterAddonInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAddonInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterAddonInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterAddonInstanceResponse) GetBody() *DescribeClusterAddonInstanceResponseBody {
	return s.Body
}

func (s *DescribeClusterAddonInstanceResponse) SetHeaders(v map[string]*string) *DescribeClusterAddonInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterAddonInstanceResponse) SetStatusCode(v int32) *DescribeClusterAddonInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterAddonInstanceResponse) SetBody(v *DescribeClusterAddonInstanceResponseBody) *DescribeClusterAddonInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterAddonInstanceResponse) Validate() error {
	return dara.Validate(s)
}
