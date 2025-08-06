// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodMultiUsageDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodMultiUsageDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodMultiUsageDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodMultiUsageDataResponseBody) *DescribeVodMultiUsageDataResponse
	GetBody() *DescribeVodMultiUsageDataResponseBody
}

type DescribeVodMultiUsageDataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodMultiUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodMultiUsageDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodMultiUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodMultiUsageDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodMultiUsageDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodMultiUsageDataResponse) GetBody() *DescribeVodMultiUsageDataResponseBody {
	return s.Body
}

func (s *DescribeVodMultiUsageDataResponse) SetHeaders(v map[string]*string) *DescribeVodMultiUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodMultiUsageDataResponse) SetStatusCode(v int32) *DescribeVodMultiUsageDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodMultiUsageDataResponse) SetBody(v *DescribeVodMultiUsageDataResponseBody) *DescribeVodMultiUsageDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodMultiUsageDataResponse) Validate() error {
	return dara.Validate(s)
}
