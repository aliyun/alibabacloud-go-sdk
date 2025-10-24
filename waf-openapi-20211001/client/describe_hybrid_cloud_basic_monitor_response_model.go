// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudBasicMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridCloudBasicMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridCloudBasicMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridCloudBasicMonitorResponseBody) *DescribeHybridCloudBasicMonitorResponse
	GetBody() *DescribeHybridCloudBasicMonitorResponseBody
}

type DescribeHybridCloudBasicMonitorResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridCloudBasicMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridCloudBasicMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudBasicMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudBasicMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridCloudBasicMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridCloudBasicMonitorResponse) GetBody() *DescribeHybridCloudBasicMonitorResponseBody {
	return s.Body
}

func (s *DescribeHybridCloudBasicMonitorResponse) SetHeaders(v map[string]*string) *DescribeHybridCloudBasicMonitorResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridCloudBasicMonitorResponse) SetStatusCode(v int32) *DescribeHybridCloudBasicMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridCloudBasicMonitorResponse) SetBody(v *DescribeHybridCloudBasicMonitorResponseBody) *DescribeHybridCloudBasicMonitorResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridCloudBasicMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
