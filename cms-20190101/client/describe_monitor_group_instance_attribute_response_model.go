// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorGroupInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMonitorGroupInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMonitorGroupInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMonitorGroupInstanceAttributeResponseBody) *DescribeMonitorGroupInstanceAttributeResponse
	GetBody() *DescribeMonitorGroupInstanceAttributeResponseBody
}

type DescribeMonitorGroupInstanceAttributeResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMonitorGroupInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMonitorGroupInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMonitorGroupInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMonitorGroupInstanceAttributeResponse) GetBody() *DescribeMonitorGroupInstanceAttributeResponseBody {
	return s.Body
}

func (s *DescribeMonitorGroupInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeMonitorGroupInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponse) SetStatusCode(v int32) *DescribeMonitorGroupInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponse) SetBody(v *DescribeMonitorGroupInstanceAttributeResponseBody) *DescribeMonitorGroupInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponse) Validate() error {
	return dara.Validate(s)
}
