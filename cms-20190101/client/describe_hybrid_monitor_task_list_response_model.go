// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridMonitorTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridMonitorTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridMonitorTaskListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridMonitorTaskListResponseBody) *DescribeHybridMonitorTaskListResponse
	GetBody() *DescribeHybridMonitorTaskListResponseBody
}

type DescribeHybridMonitorTaskListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridMonitorTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridMonitorTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorTaskListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridMonitorTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridMonitorTaskListResponse) GetBody() *DescribeHybridMonitorTaskListResponseBody {
	return s.Body
}

func (s *DescribeHybridMonitorTaskListResponse) SetHeaders(v map[string]*string) *DescribeHybridMonitorTaskListResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridMonitorTaskListResponse) SetStatusCode(v int32) *DescribeHybridMonitorTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponse) SetBody(v *DescribeHybridMonitorTaskListResponseBody) *DescribeHybridMonitorTaskListResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridMonitorTaskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
