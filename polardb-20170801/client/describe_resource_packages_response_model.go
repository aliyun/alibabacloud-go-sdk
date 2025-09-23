// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourcePackagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourcePackagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourcePackagesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourcePackagesResponseBody) *DescribeResourcePackagesResponse
	GetBody() *DescribeResourcePackagesResponseBody
}

type DescribeResourcePackagesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourcePackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourcePackagesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePackagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourcePackagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourcePackagesResponse) GetBody() *DescribeResourcePackagesResponseBody {
	return s.Body
}

func (s *DescribeResourcePackagesResponse) SetHeaders(v map[string]*string) *DescribeResourcePackagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourcePackagesResponse) SetStatusCode(v int32) *DescribeResourcePackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourcePackagesResponse) SetBody(v *DescribeResourcePackagesResponseBody) *DescribeResourcePackagesResponse {
	s.Body = v
	return s
}

func (s *DescribeResourcePackagesResponse) Validate() error {
	return dara.Validate(s)
}
