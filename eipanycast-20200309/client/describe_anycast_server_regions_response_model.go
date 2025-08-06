// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnycastServerRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAnycastServerRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAnycastServerRegionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAnycastServerRegionsResponseBody) *DescribeAnycastServerRegionsResponse
	GetBody() *DescribeAnycastServerRegionsResponseBody
}

type DescribeAnycastServerRegionsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAnycastServerRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAnycastServerRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnycastServerRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAnycastServerRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAnycastServerRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAnycastServerRegionsResponse) GetBody() *DescribeAnycastServerRegionsResponseBody {
	return s.Body
}

func (s *DescribeAnycastServerRegionsResponse) SetHeaders(v map[string]*string) *DescribeAnycastServerRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAnycastServerRegionsResponse) SetStatusCode(v int32) *DescribeAnycastServerRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAnycastServerRegionsResponse) SetBody(v *DescribeAnycastServerRegionsResponseBody) *DescribeAnycastServerRegionsResponse {
	s.Body = v
	return s
}

func (s *DescribeAnycastServerRegionsResponse) Validate() error {
	return dara.Validate(s)
}
