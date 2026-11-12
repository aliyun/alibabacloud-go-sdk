// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMonitorDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMonitorDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMonitorDataResponseBody) *DescribeMonitorDataResponse
	GetBody() *DescribeMonitorDataResponseBody
}

type DescribeMonitorDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMonitorDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMonitorDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitorDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMonitorDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMonitorDataResponse) GetBody() *DescribeMonitorDataResponseBody {
	return s.Body
}

func (s *DescribeMonitorDataResponse) SetHeaders(v map[string]*string) *DescribeMonitorDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitorDataResponse) SetStatusCode(v int32) *DescribeMonitorDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMonitorDataResponse) SetBody(v *DescribeMonitorDataResponseBody) *DescribeMonitorDataResponse {
	s.Body = v
	return s
}

func (s *DescribeMonitorDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
