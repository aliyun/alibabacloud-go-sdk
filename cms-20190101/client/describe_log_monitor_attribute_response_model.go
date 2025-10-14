// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogMonitorAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogMonitorAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogMonitorAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogMonitorAttributeResponseBody) *DescribeLogMonitorAttributeResponse
	GetBody() *DescribeLogMonitorAttributeResponseBody
}

type DescribeLogMonitorAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogMonitorAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogMonitorAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogMonitorAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogMonitorAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogMonitorAttributeResponse) GetBody() *DescribeLogMonitorAttributeResponseBody {
	return s.Body
}

func (s *DescribeLogMonitorAttributeResponse) SetHeaders(v map[string]*string) *DescribeLogMonitorAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogMonitorAttributeResponse) SetStatusCode(v int32) *DescribeLogMonitorAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponse) SetBody(v *DescribeLogMonitorAttributeResponseBody) *DescribeLogMonitorAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeLogMonitorAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
