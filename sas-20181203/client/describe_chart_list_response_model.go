// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChartListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeChartListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeChartListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeChartListResponseBody) *DescribeChartListResponse
	GetBody() *DescribeChartListResponseBody
}

type DescribeChartListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChartListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChartListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeChartListResponse) GoString() string {
	return s.String()
}

func (s *DescribeChartListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeChartListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeChartListResponse) GetBody() *DescribeChartListResponseBody {
	return s.Body
}

func (s *DescribeChartListResponse) SetHeaders(v map[string]*string) *DescribeChartListResponse {
	s.Headers = v
	return s
}

func (s *DescribeChartListResponse) SetStatusCode(v int32) *DescribeChartListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChartListResponse) SetBody(v *DescribeChartListResponseBody) *DescribeChartListResponse {
	s.Body = v
	return s
}

func (s *DescribeChartListResponse) Validate() error {
	return dara.Validate(s)
}
