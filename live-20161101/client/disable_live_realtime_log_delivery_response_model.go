// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableLiveRealtimeLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableLiveRealtimeLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableLiveRealtimeLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *DisableLiveRealtimeLogDeliveryResponseBody) *DisableLiveRealtimeLogDeliveryResponse
	GetBody() *DisableLiveRealtimeLogDeliveryResponseBody
}

type DisableLiveRealtimeLogDeliveryResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableLiveRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableLiveRealtimeLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableLiveRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *DisableLiveRealtimeLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableLiveRealtimeLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableLiveRealtimeLogDeliveryResponse) GetBody() *DisableLiveRealtimeLogDeliveryResponseBody {
	return s.Body
}

func (s *DisableLiveRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *DisableLiveRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *DisableLiveRealtimeLogDeliveryResponse) SetStatusCode(v int32) *DisableLiveRealtimeLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableLiveRealtimeLogDeliveryResponse) SetBody(v *DisableLiveRealtimeLogDeliveryResponseBody) *DisableLiveRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *DisableLiveRealtimeLogDeliveryResponse) Validate() error {
	return dara.Validate(s)
}
