// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableRealtimeLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableRealtimeLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableRealtimeLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *DisableRealtimeLogDeliveryResponseBody) *DisableRealtimeLogDeliveryResponse
	GetBody() *DisableRealtimeLogDeliveryResponseBody
}

type DisableRealtimeLogDeliveryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableRealtimeLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *DisableRealtimeLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableRealtimeLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableRealtimeLogDeliveryResponse) GetBody() *DisableRealtimeLogDeliveryResponseBody {
	return s.Body
}

func (s *DisableRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *DisableRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *DisableRealtimeLogDeliveryResponse) SetStatusCode(v int32) *DisableRealtimeLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableRealtimeLogDeliveryResponse) SetBody(v *DisableRealtimeLogDeliveryResponseBody) *DisableRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *DisableRealtimeLogDeliveryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
