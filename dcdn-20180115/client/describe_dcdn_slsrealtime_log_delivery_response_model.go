// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSLSRealtimeLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnSLSRealtimeLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnSLSRealtimeLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnSLSRealtimeLogDeliveryResponseBody) *DescribeDcdnSLSRealtimeLogDeliveryResponse
	GetBody() *DescribeDcdnSLSRealtimeLogDeliveryResponseBody
}

type DescribeDcdnSLSRealtimeLogDeliveryResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnSLSRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnSLSRealtimeLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSLSRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponse) GetBody() *DescribeDcdnSLSRealtimeLogDeliveryResponseBody {
	return s.Body
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *DescribeDcdnSLSRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponse) SetStatusCode(v int32) *DescribeDcdnSLSRealtimeLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponse) SetBody(v *DescribeDcdnSLSRealtimeLogDeliveryResponseBody) *DescribeDcdnSLSRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponse) Validate() error {
	return dara.Validate(s)
}
