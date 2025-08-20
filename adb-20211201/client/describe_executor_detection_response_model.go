// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExecutorDetectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExecutorDetectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExecutorDetectionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExecutorDetectionResponseBody) *DescribeExecutorDetectionResponse
	GetBody() *DescribeExecutorDetectionResponseBody
}

type DescribeExecutorDetectionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExecutorDetectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExecutorDetectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExecutorDetectionResponse) GoString() string {
	return s.String()
}

func (s *DescribeExecutorDetectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExecutorDetectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExecutorDetectionResponse) GetBody() *DescribeExecutorDetectionResponseBody {
	return s.Body
}

func (s *DescribeExecutorDetectionResponse) SetHeaders(v map[string]*string) *DescribeExecutorDetectionResponse {
	s.Headers = v
	return s
}

func (s *DescribeExecutorDetectionResponse) SetStatusCode(v int32) *DescribeExecutorDetectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExecutorDetectionResponse) SetBody(v *DescribeExecutorDetectionResponseBody) *DescribeExecutorDetectionResponse {
	s.Body = v
	return s
}

func (s *DescribeExecutorDetectionResponse) Validate() error {
	return dara.Validate(s)
}
