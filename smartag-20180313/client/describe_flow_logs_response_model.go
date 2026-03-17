// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFlowLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFlowLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFlowLogsResponseBody) *DescribeFlowLogsResponse
	GetBody() *DescribeFlowLogsResponseBody
}

type DescribeFlowLogsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFlowLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFlowLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFlowLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFlowLogsResponse) GetBody() *DescribeFlowLogsResponseBody {
	return s.Body
}

func (s *DescribeFlowLogsResponse) SetHeaders(v map[string]*string) *DescribeFlowLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowLogsResponse) SetStatusCode(v int32) *DescribeFlowLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFlowLogsResponse) SetBody(v *DescribeFlowLogsResponseBody) *DescribeFlowLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeFlowLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
