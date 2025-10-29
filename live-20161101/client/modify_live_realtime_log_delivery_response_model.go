// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveRealtimeLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLiveRealtimeLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLiveRealtimeLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLiveRealtimeLogDeliveryResponseBody) *ModifyLiveRealtimeLogDeliveryResponse
	GetBody() *ModifyLiveRealtimeLogDeliveryResponseBody
}

type ModifyLiveRealtimeLogDeliveryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLiveRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLiveRealtimeLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *ModifyLiveRealtimeLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLiveRealtimeLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLiveRealtimeLogDeliveryResponse) GetBody() *ModifyLiveRealtimeLogDeliveryResponseBody {
	return s.Body
}

func (s *ModifyLiveRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *ModifyLiveRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *ModifyLiveRealtimeLogDeliveryResponse) SetStatusCode(v int32) *ModifyLiveRealtimeLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLiveRealtimeLogDeliveryResponse) SetBody(v *ModifyLiveRealtimeLogDeliveryResponseBody) *ModifyLiveRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *ModifyLiveRealtimeLogDeliveryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
