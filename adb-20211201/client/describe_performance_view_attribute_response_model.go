// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePerformanceViewAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePerformanceViewAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePerformanceViewAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribePerformanceViewAttributeResponseBody) *DescribePerformanceViewAttributeResponse
	GetBody() *DescribePerformanceViewAttributeResponseBody
}

type DescribePerformanceViewAttributeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePerformanceViewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePerformanceViewAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePerformanceViewAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribePerformanceViewAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePerformanceViewAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePerformanceViewAttributeResponse) GetBody() *DescribePerformanceViewAttributeResponseBody {
	return s.Body
}

func (s *DescribePerformanceViewAttributeResponse) SetHeaders(v map[string]*string) *DescribePerformanceViewAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribePerformanceViewAttributeResponse) SetStatusCode(v int32) *DescribePerformanceViewAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePerformanceViewAttributeResponse) SetBody(v *DescribePerformanceViewAttributeResponseBody) *DescribePerformanceViewAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribePerformanceViewAttributeResponse) Validate() error {
	return dara.Validate(s)
}
