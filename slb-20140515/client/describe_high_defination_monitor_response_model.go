// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHighDefinationMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHighDefinationMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHighDefinationMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHighDefinationMonitorResponseBody) *DescribeHighDefinationMonitorResponse
	GetBody() *DescribeHighDefinationMonitorResponseBody
}

type DescribeHighDefinationMonitorResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHighDefinationMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHighDefinationMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighDefinationMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeHighDefinationMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHighDefinationMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHighDefinationMonitorResponse) GetBody() *DescribeHighDefinationMonitorResponseBody {
	return s.Body
}

func (s *DescribeHighDefinationMonitorResponse) SetHeaders(v map[string]*string) *DescribeHighDefinationMonitorResponse {
	s.Headers = v
	return s
}

func (s *DescribeHighDefinationMonitorResponse) SetStatusCode(v int32) *DescribeHighDefinationMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHighDefinationMonitorResponse) SetBody(v *DescribeHighDefinationMonitorResponseBody) *DescribeHighDefinationMonitorResponse {
	s.Body = v
	return s
}

func (s *DescribeHighDefinationMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
