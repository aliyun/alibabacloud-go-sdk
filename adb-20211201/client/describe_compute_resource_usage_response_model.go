// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComputeResourceUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeComputeResourceUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeComputeResourceUsageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeComputeResourceUsageResponseBody) *DescribeComputeResourceUsageResponse
	GetBody() *DescribeComputeResourceUsageResponseBody
}

type DescribeComputeResourceUsageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComputeResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComputeResourceUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeComputeResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeComputeResourceUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeComputeResourceUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeComputeResourceUsageResponse) GetBody() *DescribeComputeResourceUsageResponseBody {
	return s.Body
}

func (s *DescribeComputeResourceUsageResponse) SetHeaders(v map[string]*string) *DescribeComputeResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeComputeResourceUsageResponse) SetStatusCode(v int32) *DescribeComputeResourceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComputeResourceUsageResponse) SetBody(v *DescribeComputeResourceUsageResponseBody) *DescribeComputeResourceUsageResponse {
	s.Body = v
	return s
}

func (s *DescribeComputeResourceUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
