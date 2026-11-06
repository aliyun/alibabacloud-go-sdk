// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterResourceUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAIDBClusterResourceUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAIDBClusterResourceUsageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAIDBClusterResourceUsageResponseBody) *DescribeAIDBClusterResourceUsageResponse
	GetBody() *DescribeAIDBClusterResourceUsageResponseBody
}

type DescribeAIDBClusterResourceUsageResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAIDBClusterResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAIDBClusterResourceUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterResourceUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAIDBClusterResourceUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAIDBClusterResourceUsageResponse) GetBody() *DescribeAIDBClusterResourceUsageResponseBody {
	return s.Body
}

func (s *DescribeAIDBClusterResourceUsageResponse) SetHeaders(v map[string]*string) *DescribeAIDBClusterResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponse) SetStatusCode(v int32) *DescribeAIDBClusterResourceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponse) SetBody(v *DescribeAIDBClusterResourceUsageResponseBody) *DescribeAIDBClusterResourceUsageResponse {
	s.Body = v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
