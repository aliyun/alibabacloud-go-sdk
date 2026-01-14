// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationMonitorResponseBody) *DescribeApplicationMonitorResponse
	GetBody() *DescribeApplicationMonitorResponseBody
}

type DescribeApplicationMonitorResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationMonitorResponse) GetBody() *DescribeApplicationMonitorResponseBody {
	return s.Body
}

func (s *DescribeApplicationMonitorResponse) SetHeaders(v map[string]*string) *DescribeApplicationMonitorResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationMonitorResponse) SetStatusCode(v int32) *DescribeApplicationMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationMonitorResponse) SetBody(v *DescribeApplicationMonitorResponseBody) *DescribeApplicationMonitorResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
