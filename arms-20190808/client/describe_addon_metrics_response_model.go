// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddonMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAddonMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAddonMetricsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAddonMetricsResponseBody) *DescribeAddonMetricsResponse
	GetBody() *DescribeAddonMetricsResponseBody
}

type DescribeAddonMetricsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAddonMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAddonMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAddonMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAddonMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAddonMetricsResponse) GetBody() *DescribeAddonMetricsResponseBody {
	return s.Body
}

func (s *DescribeAddonMetricsResponse) SetHeaders(v map[string]*string) *DescribeAddonMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAddonMetricsResponse) SetStatusCode(v int32) *DescribeAddonMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAddonMetricsResponse) SetBody(v *DescribeAddonMetricsResponseBody) *DescribeAddonMetricsResponse {
	s.Body = v
	return s
}

func (s *DescribeAddonMetricsResponse) Validate() error {
	return dara.Validate(s)
}
