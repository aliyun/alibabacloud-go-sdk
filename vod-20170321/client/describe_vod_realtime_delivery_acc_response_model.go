// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodRealtimeDeliveryAccResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodRealtimeDeliveryAccResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodRealtimeDeliveryAccResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodRealtimeDeliveryAccResponseBody) *DescribeVodRealtimeDeliveryAccResponse
	GetBody() *DescribeVodRealtimeDeliveryAccResponseBody
}

type DescribeVodRealtimeDeliveryAccResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodRealtimeDeliveryAccResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodRealtimeDeliveryAccResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodRealtimeDeliveryAccResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodRealtimeDeliveryAccResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodRealtimeDeliveryAccResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodRealtimeDeliveryAccResponse) GetBody() *DescribeVodRealtimeDeliveryAccResponseBody {
	return s.Body
}

func (s *DescribeVodRealtimeDeliveryAccResponse) SetHeaders(v map[string]*string) *DescribeVodRealtimeDeliveryAccResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodRealtimeDeliveryAccResponse) SetStatusCode(v int32) *DescribeVodRealtimeDeliveryAccResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodRealtimeDeliveryAccResponse) SetBody(v *DescribeVodRealtimeDeliveryAccResponseBody) *DescribeVodRealtimeDeliveryAccResponse {
	s.Body = v
	return s
}

func (s *DescribeVodRealtimeDeliveryAccResponse) Validate() error {
	return dara.Validate(s)
}
