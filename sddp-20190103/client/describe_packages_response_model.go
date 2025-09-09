// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePackagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePackagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePackagesResponse
	GetStatusCode() *int32
	SetBody(v *DescribePackagesResponseBody) *DescribePackagesResponse
	GetBody() *DescribePackagesResponseBody
}

type DescribePackagesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePackagesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePackagesResponse) GoString() string {
	return s.String()
}

func (s *DescribePackagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePackagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePackagesResponse) GetBody() *DescribePackagesResponseBody {
	return s.Body
}

func (s *DescribePackagesResponse) SetHeaders(v map[string]*string) *DescribePackagesResponse {
	s.Headers = v
	return s
}

func (s *DescribePackagesResponse) SetStatusCode(v int32) *DescribePackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePackagesResponse) SetBody(v *DescribePackagesResponseBody) *DescribePackagesResponse {
	s.Body = v
	return s
}

func (s *DescribePackagesResponse) Validate() error {
	return dara.Validate(s)
}
