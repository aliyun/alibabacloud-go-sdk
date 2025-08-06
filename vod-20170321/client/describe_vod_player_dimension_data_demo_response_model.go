// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodPlayerDimensionDataDemoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodPlayerDimensionDataDemoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodPlayerDimensionDataDemoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodPlayerDimensionDataDemoResponseBody) *DescribeVodPlayerDimensionDataDemoResponse
	GetBody() *DescribeVodPlayerDimensionDataDemoResponseBody
}

type DescribeVodPlayerDimensionDataDemoResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodPlayerDimensionDataDemoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodPlayerDimensionDataDemoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerDimensionDataDemoResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerDimensionDataDemoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodPlayerDimensionDataDemoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodPlayerDimensionDataDemoResponse) GetBody() *DescribeVodPlayerDimensionDataDemoResponseBody {
	return s.Body
}

func (s *DescribeVodPlayerDimensionDataDemoResponse) SetHeaders(v map[string]*string) *DescribeVodPlayerDimensionDataDemoResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodPlayerDimensionDataDemoResponse) SetStatusCode(v int32) *DescribeVodPlayerDimensionDataDemoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodPlayerDimensionDataDemoResponse) SetBody(v *DescribeVodPlayerDimensionDataDemoResponseBody) *DescribeVodPlayerDimensionDataDemoResponse {
	s.Body = v
	return s
}

func (s *DescribeVodPlayerDimensionDataDemoResponse) Validate() error {
	return dara.Validate(s)
}
