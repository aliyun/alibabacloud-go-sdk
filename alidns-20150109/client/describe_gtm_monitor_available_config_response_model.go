// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmMonitorAvailableConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGtmMonitorAvailableConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGtmMonitorAvailableConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGtmMonitorAvailableConfigResponseBody) *DescribeGtmMonitorAvailableConfigResponse
	GetBody() *DescribeGtmMonitorAvailableConfigResponseBody
}

type DescribeGtmMonitorAvailableConfigResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGtmMonitorAvailableConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGtmMonitorAvailableConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmMonitorAvailableConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmMonitorAvailableConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGtmMonitorAvailableConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGtmMonitorAvailableConfigResponse) GetBody() *DescribeGtmMonitorAvailableConfigResponseBody {
	return s.Body
}

func (s *DescribeGtmMonitorAvailableConfigResponse) SetHeaders(v map[string]*string) *DescribeGtmMonitorAvailableConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponse) SetStatusCode(v int32) *DescribeGtmMonitorAvailableConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponse) SetBody(v *DescribeGtmMonitorAvailableConfigResponseBody) *DescribeGtmMonitorAvailableConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
