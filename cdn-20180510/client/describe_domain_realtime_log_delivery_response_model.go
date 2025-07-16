// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealtimeLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainRealtimeLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainRealtimeLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainRealtimeLogDeliveryResponseBody) *DescribeDomainRealtimeLogDeliveryResponse
	GetBody() *DescribeDomainRealtimeLogDeliveryResponseBody
}

type DescribeDomainRealtimeLogDeliveryResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainRealtimeLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealtimeLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainRealtimeLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainRealtimeLogDeliveryResponse) GetBody() *DescribeDomainRealtimeLogDeliveryResponseBody {
	return s.Body
}

func (s *DescribeDomainRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *DescribeDomainRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealtimeLogDeliveryResponse) SetStatusCode(v int32) *DescribeDomainRealtimeLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainRealtimeLogDeliveryResponse) SetBody(v *DescribeDomainRealtimeLogDeliveryResponseBody) *DescribeDomainRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainRealtimeLogDeliveryResponse) Validate() error {
	return dara.Validate(s)
}
