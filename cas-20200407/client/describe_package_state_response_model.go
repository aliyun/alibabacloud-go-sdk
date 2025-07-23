// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePackageStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePackageStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePackageStateResponse
	GetStatusCode() *int32
	SetBody(v *DescribePackageStateResponseBody) *DescribePackageStateResponse
	GetBody() *DescribePackageStateResponseBody
}

type DescribePackageStateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePackageStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePackageStateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePackageStateResponse) GoString() string {
	return s.String()
}

func (s *DescribePackageStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePackageStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePackageStateResponse) GetBody() *DescribePackageStateResponseBody {
	return s.Body
}

func (s *DescribePackageStateResponse) SetHeaders(v map[string]*string) *DescribePackageStateResponse {
	s.Headers = v
	return s
}

func (s *DescribePackageStateResponse) SetStatusCode(v int32) *DescribePackageStateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePackageStateResponse) SetBody(v *DescribePackageStateResponseBody) *DescribePackageStateResponse {
	s.Body = v
	return s
}

func (s *DescribePackageStateResponse) Validate() error {
	return dara.Validate(s)
}
