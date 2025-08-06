// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodPlayerCollectDataDemoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodPlayerCollectDataDemoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodPlayerCollectDataDemoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodPlayerCollectDataDemoResponseBody) *DescribeVodPlayerCollectDataDemoResponse
	GetBody() *DescribeVodPlayerCollectDataDemoResponseBody
}

type DescribeVodPlayerCollectDataDemoResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodPlayerCollectDataDemoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodPlayerCollectDataDemoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerCollectDataDemoResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerCollectDataDemoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodPlayerCollectDataDemoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodPlayerCollectDataDemoResponse) GetBody() *DescribeVodPlayerCollectDataDemoResponseBody {
	return s.Body
}

func (s *DescribeVodPlayerCollectDataDemoResponse) SetHeaders(v map[string]*string) *DescribeVodPlayerCollectDataDemoResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodPlayerCollectDataDemoResponse) SetStatusCode(v int32) *DescribeVodPlayerCollectDataDemoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodPlayerCollectDataDemoResponse) SetBody(v *DescribeVodPlayerCollectDataDemoResponseBody) *DescribeVodPlayerCollectDataDemoResponse {
	s.Body = v
	return s
}

func (s *DescribeVodPlayerCollectDataDemoResponse) Validate() error {
	return dara.Validate(s)
}
