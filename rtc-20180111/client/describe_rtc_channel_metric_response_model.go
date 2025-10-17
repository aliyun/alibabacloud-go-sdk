// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcChannelMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRtcChannelMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRtcChannelMetricResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRtcChannelMetricResponseBody) *DescribeRtcChannelMetricResponse
	GetBody() *DescribeRtcChannelMetricResponseBody
}

type DescribeRtcChannelMetricResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRtcChannelMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRtcChannelMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcChannelMetricResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRtcChannelMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRtcChannelMetricResponse) GetBody() *DescribeRtcChannelMetricResponseBody {
	return s.Body
}

func (s *DescribeRtcChannelMetricResponse) SetHeaders(v map[string]*string) *DescribeRtcChannelMetricResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcChannelMetricResponse) SetStatusCode(v int32) *DescribeRtcChannelMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRtcChannelMetricResponse) SetBody(v *DescribeRtcChannelMetricResponseBody) *DescribeRtcChannelMetricResponse {
	s.Body = v
	return s
}

func (s *DescribeRtcChannelMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
