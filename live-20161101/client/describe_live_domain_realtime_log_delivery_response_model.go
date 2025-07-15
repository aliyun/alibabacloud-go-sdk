// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainRealtimeLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainRealtimeLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainRealtimeLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainRealtimeLogDeliveryResponseBody) *DescribeLiveDomainRealtimeLogDeliveryResponse
	GetBody() *DescribeLiveDomainRealtimeLogDeliveryResponseBody
}

type DescribeLiveDomainRealtimeLogDeliveryResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainRealtimeLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponse) GetBody() *DescribeLiveDomainRealtimeLogDeliveryResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponse) SetStatusCode(v int32) *DescribeLiveDomainRealtimeLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponse) SetBody(v *DescribeLiveDomainRealtimeLogDeliveryResponseBody) *DescribeLiveDomainRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponse) Validate() error {
	return dara.Validate(s)
}
