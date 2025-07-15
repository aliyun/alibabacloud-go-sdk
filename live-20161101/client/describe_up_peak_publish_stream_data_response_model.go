// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpPeakPublishStreamDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUpPeakPublishStreamDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUpPeakPublishStreamDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUpPeakPublishStreamDataResponseBody) *DescribeUpPeakPublishStreamDataResponse
	GetBody() *DescribeUpPeakPublishStreamDataResponseBody
}

type DescribeUpPeakPublishStreamDataResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUpPeakPublishStreamDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUpPeakPublishStreamDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpPeakPublishStreamDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUpPeakPublishStreamDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUpPeakPublishStreamDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUpPeakPublishStreamDataResponse) GetBody() *DescribeUpPeakPublishStreamDataResponseBody {
	return s.Body
}

func (s *DescribeUpPeakPublishStreamDataResponse) SetHeaders(v map[string]*string) *DescribeUpPeakPublishStreamDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeUpPeakPublishStreamDataResponse) SetStatusCode(v int32) *DescribeUpPeakPublishStreamDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataResponse) SetBody(v *DescribeUpPeakPublishStreamDataResponseBody) *DescribeUpPeakPublishStreamDataResponse {
	s.Body = v
	return s
}

func (s *DescribeUpPeakPublishStreamDataResponse) Validate() error {
	return dara.Validate(s)
}
