// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataLimitsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataLimitsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataLimitsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataLimitsResponseBody) *DescribeDataLimitsResponse
	GetBody() *DescribeDataLimitsResponseBody
}

type DescribeDataLimitsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataLimitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataLimitsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataLimitsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataLimitsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataLimitsResponse) GetBody() *DescribeDataLimitsResponseBody {
	return s.Body
}

func (s *DescribeDataLimitsResponse) SetHeaders(v map[string]*string) *DescribeDataLimitsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataLimitsResponse) SetStatusCode(v int32) *DescribeDataLimitsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataLimitsResponse) SetBody(v *DescribeDataLimitsResponseBody) *DescribeDataLimitsResponse {
	s.Body = v
	return s
}

func (s *DescribeDataLimitsResponse) Validate() error {
	return dara.Validate(s)
}
