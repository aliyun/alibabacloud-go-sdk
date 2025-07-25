// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmMonitorConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGtmMonitorConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGtmMonitorConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGtmMonitorConfigResponseBody) *DescribeGtmMonitorConfigResponse
	GetBody() *DescribeGtmMonitorConfigResponseBody
}

type DescribeGtmMonitorConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGtmMonitorConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGtmMonitorConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmMonitorConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmMonitorConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGtmMonitorConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGtmMonitorConfigResponse) GetBody() *DescribeGtmMonitorConfigResponseBody {
	return s.Body
}

func (s *DescribeGtmMonitorConfigResponse) SetHeaders(v map[string]*string) *DescribeGtmMonitorConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmMonitorConfigResponse) SetStatusCode(v int32) *DescribeGtmMonitorConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponse) SetBody(v *DescribeGtmMonitorConfigResponseBody) *DescribeGtmMonitorConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeGtmMonitorConfigResponse) Validate() error {
	return dara.Validate(s)
}
