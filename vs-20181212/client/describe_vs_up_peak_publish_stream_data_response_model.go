// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsUpPeakPublishStreamDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsUpPeakPublishStreamDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsUpPeakPublishStreamDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsUpPeakPublishStreamDataResponseBody) *DescribeVsUpPeakPublishStreamDataResponse
	GetBody() *DescribeVsUpPeakPublishStreamDataResponseBody
}

type DescribeVsUpPeakPublishStreamDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsUpPeakPublishStreamDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsUpPeakPublishStreamDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsUpPeakPublishStreamDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsUpPeakPublishStreamDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsUpPeakPublishStreamDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsUpPeakPublishStreamDataResponse) GetBody() *DescribeVsUpPeakPublishStreamDataResponseBody {
	return s.Body
}

func (s *DescribeVsUpPeakPublishStreamDataResponse) SetHeaders(v map[string]*string) *DescribeVsUpPeakPublishStreamDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataResponse) SetStatusCode(v int32) *DescribeVsUpPeakPublishStreamDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataResponse) SetBody(v *DescribeVsUpPeakPublishStreamDataResponseBody) *DescribeVsUpPeakPublishStreamDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVsUpPeakPublishStreamDataResponse) Validate() error {
	return dara.Validate(s)
}
