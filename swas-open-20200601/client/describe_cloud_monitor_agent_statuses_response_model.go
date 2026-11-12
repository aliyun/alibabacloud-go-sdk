// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudMonitorAgentStatusesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudMonitorAgentStatusesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudMonitorAgentStatusesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudMonitorAgentStatusesResponseBody) *DescribeCloudMonitorAgentStatusesResponse
	GetBody() *DescribeCloudMonitorAgentStatusesResponseBody
}

type DescribeCloudMonitorAgentStatusesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudMonitorAgentStatusesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudMonitorAgentStatusesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudMonitorAgentStatusesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudMonitorAgentStatusesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudMonitorAgentStatusesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudMonitorAgentStatusesResponse) GetBody() *DescribeCloudMonitorAgentStatusesResponseBody {
	return s.Body
}

func (s *DescribeCloudMonitorAgentStatusesResponse) SetHeaders(v map[string]*string) *DescribeCloudMonitorAgentStatusesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudMonitorAgentStatusesResponse) SetStatusCode(v int32) *DescribeCloudMonitorAgentStatusesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudMonitorAgentStatusesResponse) SetBody(v *DescribeCloudMonitorAgentStatusesResponseBody) *DescribeCloudMonitorAgentStatusesResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudMonitorAgentStatusesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
