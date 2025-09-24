// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceCoverageTotalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceCoverageTotalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceCoverageTotalResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceCoverageTotalResponseBody) *DescribeResourceCoverageTotalResponse
	GetBody() *DescribeResourceCoverageTotalResponseBody
}

type DescribeResourceCoverageTotalResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceCoverageTotalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceCoverageTotalResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceCoverageTotalResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceCoverageTotalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceCoverageTotalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceCoverageTotalResponse) GetBody() *DescribeResourceCoverageTotalResponseBody {
	return s.Body
}

func (s *DescribeResourceCoverageTotalResponse) SetHeaders(v map[string]*string) *DescribeResourceCoverageTotalResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceCoverageTotalResponse) SetStatusCode(v int32) *DescribeResourceCoverageTotalResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceCoverageTotalResponse) SetBody(v *DescribeResourceCoverageTotalResponseBody) *DescribeResourceCoverageTotalResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceCoverageTotalResponse) Validate() error {
	return dara.Validate(s)
}
