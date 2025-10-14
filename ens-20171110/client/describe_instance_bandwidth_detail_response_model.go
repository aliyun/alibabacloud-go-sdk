// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceBandwidthDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceBandwidthDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceBandwidthDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceBandwidthDetailResponseBody) *DescribeInstanceBandwidthDetailResponse
	GetBody() *DescribeInstanceBandwidthDetailResponseBody
}

type DescribeInstanceBandwidthDetailResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceBandwidthDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceBandwidthDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceBandwidthDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceBandwidthDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceBandwidthDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceBandwidthDetailResponse) GetBody() *DescribeInstanceBandwidthDetailResponseBody {
	return s.Body
}

func (s *DescribeInstanceBandwidthDetailResponse) SetHeaders(v map[string]*string) *DescribeInstanceBandwidthDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceBandwidthDetailResponse) SetStatusCode(v int32) *DescribeInstanceBandwidthDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailResponse) SetBody(v *DescribeInstanceBandwidthDetailResponseBody) *DescribeInstanceBandwidthDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceBandwidthDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
