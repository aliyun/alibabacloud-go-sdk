// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRealtimeLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRealtimeLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRealtimeLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRealtimeLogDeliveryResponseBody) *ModifyRealtimeLogDeliveryResponse
	GetBody() *ModifyRealtimeLogDeliveryResponseBody
}

type ModifyRealtimeLogDeliveryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRealtimeLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *ModifyRealtimeLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRealtimeLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRealtimeLogDeliveryResponse) GetBody() *ModifyRealtimeLogDeliveryResponseBody {
	return s.Body
}

func (s *ModifyRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *ModifyRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *ModifyRealtimeLogDeliveryResponse) SetStatusCode(v int32) *ModifyRealtimeLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRealtimeLogDeliveryResponse) SetBody(v *ModifyRealtimeLogDeliveryResponseBody) *ModifyRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *ModifyRealtimeLogDeliveryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
