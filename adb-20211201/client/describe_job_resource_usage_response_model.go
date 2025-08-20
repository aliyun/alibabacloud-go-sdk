// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobResourceUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeJobResourceUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeJobResourceUsageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeJobResourceUsageResponseBody) *DescribeJobResourceUsageResponse
	GetBody() *DescribeJobResourceUsageResponseBody
}

type DescribeJobResourceUsageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeJobResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeJobResourceUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobResourceUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeJobResourceUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeJobResourceUsageResponse) GetBody() *DescribeJobResourceUsageResponseBody {
	return s.Body
}

func (s *DescribeJobResourceUsageResponse) SetHeaders(v map[string]*string) *DescribeJobResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobResourceUsageResponse) SetStatusCode(v int32) *DescribeJobResourceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJobResourceUsageResponse) SetBody(v *DescribeJobResourceUsageResponseBody) *DescribeJobResourceUsageResponse {
	s.Body = v
	return s
}

func (s *DescribeJobResourceUsageResponse) Validate() error {
	return dara.Validate(s)
}
