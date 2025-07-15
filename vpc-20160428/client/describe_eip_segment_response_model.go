// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEipSegmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEipSegmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEipSegmentResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEipSegmentResponseBody) *DescribeEipSegmentResponse
	GetBody() *DescribeEipSegmentResponseBody
}

type DescribeEipSegmentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEipSegmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEipSegmentResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipSegmentResponse) GoString() string {
	return s.String()
}

func (s *DescribeEipSegmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEipSegmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEipSegmentResponse) GetBody() *DescribeEipSegmentResponseBody {
	return s.Body
}

func (s *DescribeEipSegmentResponse) SetHeaders(v map[string]*string) *DescribeEipSegmentResponse {
	s.Headers = v
	return s
}

func (s *DescribeEipSegmentResponse) SetStatusCode(v int32) *DescribeEipSegmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEipSegmentResponse) SetBody(v *DescribeEipSegmentResponseBody) *DescribeEipSegmentResponse {
	s.Body = v
	return s
}

func (s *DescribeEipSegmentResponse) Validate() error {
	return dara.Validate(s)
}
