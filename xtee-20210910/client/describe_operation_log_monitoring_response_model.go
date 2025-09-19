// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperationLogMonitoringResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOperationLogMonitoringResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOperationLogMonitoringResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOperationLogMonitoringResponseBody) *DescribeOperationLogMonitoringResponse
	GetBody() *DescribeOperationLogMonitoringResponseBody
}

type DescribeOperationLogMonitoringResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOperationLogMonitoringResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOperationLogMonitoringResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperationLogMonitoringResponse) GoString() string {
	return s.String()
}

func (s *DescribeOperationLogMonitoringResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOperationLogMonitoringResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOperationLogMonitoringResponse) GetBody() *DescribeOperationLogMonitoringResponseBody {
	return s.Body
}

func (s *DescribeOperationLogMonitoringResponse) SetHeaders(v map[string]*string) *DescribeOperationLogMonitoringResponse {
	s.Headers = v
	return s
}

func (s *DescribeOperationLogMonitoringResponse) SetStatusCode(v int32) *DescribeOperationLogMonitoringResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOperationLogMonitoringResponse) SetBody(v *DescribeOperationLogMonitoringResponseBody) *DescribeOperationLogMonitoringResponse {
	s.Body = v
	return s
}

func (s *DescribeOperationLogMonitoringResponse) Validate() error {
	return dara.Validate(s)
}
