// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommonBandwidthPackagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCommonBandwidthPackagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCommonBandwidthPackagesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCommonBandwidthPackagesResponseBody) *DescribeCommonBandwidthPackagesResponse
	GetBody() *DescribeCommonBandwidthPackagesResponseBody
}

type DescribeCommonBandwidthPackagesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCommonBandwidthPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCommonBandwidthPackagesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonBandwidthPackagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCommonBandwidthPackagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCommonBandwidthPackagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCommonBandwidthPackagesResponse) GetBody() *DescribeCommonBandwidthPackagesResponseBody {
	return s.Body
}

func (s *DescribeCommonBandwidthPackagesResponse) SetHeaders(v map[string]*string) *DescribeCommonBandwidthPackagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponse) SetStatusCode(v int32) *DescribeCommonBandwidthPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponse) SetBody(v *DescribeCommonBandwidthPackagesResponseBody) *DescribeCommonBandwidthPackagesResponse {
	s.Body = v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
