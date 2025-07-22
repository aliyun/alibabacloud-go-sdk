// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcPeakChannelCntDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRtcPeakChannelCntDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRtcPeakChannelCntDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRtcPeakChannelCntDataResponseBody) *DescribeRtcPeakChannelCntDataResponse
	GetBody() *DescribeRtcPeakChannelCntDataResponseBody
}

type DescribeRtcPeakChannelCntDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRtcPeakChannelCntDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRtcPeakChannelCntDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcPeakChannelCntDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcPeakChannelCntDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRtcPeakChannelCntDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRtcPeakChannelCntDataResponse) GetBody() *DescribeRtcPeakChannelCntDataResponseBody {
	return s.Body
}

func (s *DescribeRtcPeakChannelCntDataResponse) SetHeaders(v map[string]*string) *DescribeRtcPeakChannelCntDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcPeakChannelCntDataResponse) SetStatusCode(v int32) *DescribeRtcPeakChannelCntDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataResponse) SetBody(v *DescribeRtcPeakChannelCntDataResponseBody) *DescribeRtcPeakChannelCntDataResponse {
	s.Body = v
	return s
}

func (s *DescribeRtcPeakChannelCntDataResponse) Validate() error {
	return dara.Validate(s)
}
