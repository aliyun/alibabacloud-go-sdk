// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealtimeLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainRealtimeLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainRealtimeLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainRealtimeLogDeliveryResponseBody) *DescribeVodDomainRealtimeLogDeliveryResponse
	GetBody() *DescribeVodDomainRealtimeLogDeliveryResponseBody
}

type DescribeVodDomainRealtimeLogDeliveryResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainRealtimeLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponse) GetBody() *DescribeVodDomainRealtimeLogDeliveryResponseBody {
	return s.Body
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *DescribeVodDomainRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponse) SetStatusCode(v int32) *DescribeVodDomainRealtimeLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponse) SetBody(v *DescribeVodDomainRealtimeLogDeliveryResponseBody) *DescribeVodDomainRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponse) Validate() error {
	return dara.Validate(s)
}
