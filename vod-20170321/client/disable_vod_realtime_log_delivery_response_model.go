// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableVodRealtimeLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableVodRealtimeLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableVodRealtimeLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *DisableVodRealtimeLogDeliveryResponseBody) *DisableVodRealtimeLogDeliveryResponse
	GetBody() *DisableVodRealtimeLogDeliveryResponseBody
}

type DisableVodRealtimeLogDeliveryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableVodRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableVodRealtimeLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableVodRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *DisableVodRealtimeLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableVodRealtimeLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableVodRealtimeLogDeliveryResponse) GetBody() *DisableVodRealtimeLogDeliveryResponseBody {
	return s.Body
}

func (s *DisableVodRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *DisableVodRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *DisableVodRealtimeLogDeliveryResponse) SetStatusCode(v int32) *DisableVodRealtimeLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableVodRealtimeLogDeliveryResponse) SetBody(v *DisableVodRealtimeLogDeliveryResponseBody) *DisableVodRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *DisableVodRealtimeLogDeliveryResponse) Validate() error {
	return dara.Validate(s)
}
