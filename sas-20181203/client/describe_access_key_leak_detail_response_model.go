// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessKeyLeakDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccessKeyLeakDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccessKeyLeakDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccessKeyLeakDetailResponseBody) *DescribeAccessKeyLeakDetailResponse
	GetBody() *DescribeAccessKeyLeakDetailResponseBody
}

type DescribeAccessKeyLeakDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccessKeyLeakDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccessKeyLeakDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessKeyLeakDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessKeyLeakDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccessKeyLeakDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccessKeyLeakDetailResponse) GetBody() *DescribeAccessKeyLeakDetailResponseBody {
	return s.Body
}

func (s *DescribeAccessKeyLeakDetailResponse) SetHeaders(v map[string]*string) *DescribeAccessKeyLeakDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponse) SetStatusCode(v int32) *DescribeAccessKeyLeakDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponse) SetBody(v *DescribeAccessKeyLeakDetailResponseBody) *DescribeAccessKeyLeakDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponse) Validate() error {
	return dara.Validate(s)
}
