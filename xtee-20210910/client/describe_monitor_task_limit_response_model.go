// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorTaskLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMonitorTaskLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMonitorTaskLimitResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMonitorTaskLimitResponseBody) *DescribeMonitorTaskLimitResponse
	GetBody() *DescribeMonitorTaskLimitResponseBody
}

type DescribeMonitorTaskLimitResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMonitorTaskLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMonitorTaskLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorTaskLimitResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitorTaskLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMonitorTaskLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMonitorTaskLimitResponse) GetBody() *DescribeMonitorTaskLimitResponseBody {
	return s.Body
}

func (s *DescribeMonitorTaskLimitResponse) SetHeaders(v map[string]*string) *DescribeMonitorTaskLimitResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitorTaskLimitResponse) SetStatusCode(v int32) *DescribeMonitorTaskLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMonitorTaskLimitResponse) SetBody(v *DescribeMonitorTaskLimitResponseBody) *DescribeMonitorTaskLimitResponse {
	s.Body = v
	return s
}

func (s *DescribeMonitorTaskLimitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
