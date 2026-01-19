// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterTaskMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAIDBClusterTaskMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAIDBClusterTaskMetricsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAIDBClusterTaskMetricsResponseBody) *DescribeAIDBClusterTaskMetricsResponse
	GetBody() *DescribeAIDBClusterTaskMetricsResponseBody
}

type DescribeAIDBClusterTaskMetricsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAIDBClusterTaskMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAIDBClusterTaskMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterTaskMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterTaskMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAIDBClusterTaskMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAIDBClusterTaskMetricsResponse) GetBody() *DescribeAIDBClusterTaskMetricsResponseBody {
	return s.Body
}

func (s *DescribeAIDBClusterTaskMetricsResponse) SetHeaders(v map[string]*string) *DescribeAIDBClusterTaskMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponse) SetStatusCode(v int32) *DescribeAIDBClusterTaskMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponse) SetBody(v *DescribeAIDBClusterTaskMetricsResponseBody) *DescribeAIDBClusterTaskMetricsResponse {
	s.Body = v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
