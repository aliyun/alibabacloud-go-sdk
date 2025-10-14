// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridMonitorSLSGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridMonitorSLSGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridMonitorSLSGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridMonitorSLSGroupResponseBody) *DescribeHybridMonitorSLSGroupResponse
	GetBody() *DescribeHybridMonitorSLSGroupResponseBody
}

type DescribeHybridMonitorSLSGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridMonitorSLSGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridMonitorSLSGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorSLSGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorSLSGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridMonitorSLSGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridMonitorSLSGroupResponse) GetBody() *DescribeHybridMonitorSLSGroupResponseBody {
	return s.Body
}

func (s *DescribeHybridMonitorSLSGroupResponse) SetHeaders(v map[string]*string) *DescribeHybridMonitorSLSGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridMonitorSLSGroupResponse) SetStatusCode(v int32) *DescribeHybridMonitorSLSGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupResponse) SetBody(v *DescribeHybridMonitorSLSGroupResponseBody) *DescribeHybridMonitorSLSGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridMonitorSLSGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
