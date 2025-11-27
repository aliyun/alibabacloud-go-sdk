// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceUsageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceUsageResponseBody) *DescribeResourceUsageResponse
	GetBody() *DescribeResourceUsageResponseBody
}

type DescribeResourceUsageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceUsageResponse) GetBody() *DescribeResourceUsageResponseBody {
	return s.Body
}

func (s *DescribeResourceUsageResponse) SetHeaders(v map[string]*string) *DescribeResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceUsageResponse) SetStatusCode(v int32) *DescribeResourceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceUsageResponse) SetBody(v *DescribeResourceUsageResponseBody) *DescribeResourceUsageResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
