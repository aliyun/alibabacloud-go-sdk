// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveRealtimeLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveRealtimeLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveRealtimeLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveRealtimeLogDeliveryResponseBody) *DeleteLiveRealtimeLogDeliveryResponse
	GetBody() *DeleteLiveRealtimeLogDeliveryResponseBody
}

type DeleteLiveRealtimeLogDeliveryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveRealtimeLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveRealtimeLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveRealtimeLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveRealtimeLogDeliveryResponse) GetBody() *DeleteLiveRealtimeLogDeliveryResponseBody {
	return s.Body
}

func (s *DeleteLiveRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *DeleteLiveRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveRealtimeLogDeliveryResponse) SetStatusCode(v int32) *DeleteLiveRealtimeLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveRealtimeLogDeliveryResponse) SetBody(v *DeleteLiveRealtimeLogDeliveryResponseBody) *DeleteLiveRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveRealtimeLogDeliveryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
