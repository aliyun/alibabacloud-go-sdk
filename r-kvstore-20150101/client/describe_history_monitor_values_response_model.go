// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryMonitorValuesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHistoryMonitorValuesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHistoryMonitorValuesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHistoryMonitorValuesResponseBody) *DescribeHistoryMonitorValuesResponse
	GetBody() *DescribeHistoryMonitorValuesResponseBody
}

type DescribeHistoryMonitorValuesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHistoryMonitorValuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHistoryMonitorValuesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryMonitorValuesResponse) GoString() string {
	return s.String()
}

func (s *DescribeHistoryMonitorValuesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHistoryMonitorValuesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHistoryMonitorValuesResponse) GetBody() *DescribeHistoryMonitorValuesResponseBody {
	return s.Body
}

func (s *DescribeHistoryMonitorValuesResponse) SetHeaders(v map[string]*string) *DescribeHistoryMonitorValuesResponse {
	s.Headers = v
	return s
}

func (s *DescribeHistoryMonitorValuesResponse) SetStatusCode(v int32) *DescribeHistoryMonitorValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHistoryMonitorValuesResponse) SetBody(v *DescribeHistoryMonitorValuesResponseBody) *DescribeHistoryMonitorValuesResponse {
	s.Body = v
	return s
}

func (s *DescribeHistoryMonitorValuesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
