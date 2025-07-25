// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmMonitorTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudGtmMonitorTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudGtmMonitorTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudGtmMonitorTemplateResponseBody) *DescribeCloudGtmMonitorTemplateResponse
	GetBody() *DescribeCloudGtmMonitorTemplateResponseBody
}

type DescribeCloudGtmMonitorTemplateResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudGtmMonitorTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudGtmMonitorTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmMonitorTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmMonitorTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudGtmMonitorTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudGtmMonitorTemplateResponse) GetBody() *DescribeCloudGtmMonitorTemplateResponseBody {
	return s.Body
}

func (s *DescribeCloudGtmMonitorTemplateResponse) SetHeaders(v map[string]*string) *DescribeCloudGtmMonitorTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponse) SetStatusCode(v int32) *DescribeCloudGtmMonitorTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponse) SetBody(v *DescribeCloudGtmMonitorTemplateResponseBody) *DescribeCloudGtmMonitorTemplateResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponse) Validate() error {
	return dara.Validate(s)
}
