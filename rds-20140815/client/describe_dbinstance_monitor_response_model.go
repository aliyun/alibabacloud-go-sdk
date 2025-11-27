// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceMonitorResponseBody) *DescribeDBInstanceMonitorResponse
	GetBody() *DescribeDBInstanceMonitorResponseBody
}

type DescribeDBInstanceMonitorResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceMonitorResponse) GetBody() *DescribeDBInstanceMonitorResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceMonitorResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceMonitorResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceMonitorResponse) SetStatusCode(v int32) *DescribeDBInstanceMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceMonitorResponse) SetBody(v *DescribeDBInstanceMonitorResponseBody) *DescribeDBInstanceMonitorResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
