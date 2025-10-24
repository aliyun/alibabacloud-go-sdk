// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudProcessMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridCloudProcessMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridCloudProcessMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridCloudProcessMonitorResponseBody) *DescribeHybridCloudProcessMonitorResponse
	GetBody() *DescribeHybridCloudProcessMonitorResponseBody
}

type DescribeHybridCloudProcessMonitorResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridCloudProcessMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridCloudProcessMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudProcessMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudProcessMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridCloudProcessMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridCloudProcessMonitorResponse) GetBody() *DescribeHybridCloudProcessMonitorResponseBody {
	return s.Body
}

func (s *DescribeHybridCloudProcessMonitorResponse) SetHeaders(v map[string]*string) *DescribeHybridCloudProcessMonitorResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridCloudProcessMonitorResponse) SetStatusCode(v int32) *DescribeHybridCloudProcessMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridCloudProcessMonitorResponse) SetBody(v *DescribeHybridCloudProcessMonitorResponseBody) *DescribeHybridCloudProcessMonitorResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridCloudProcessMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
