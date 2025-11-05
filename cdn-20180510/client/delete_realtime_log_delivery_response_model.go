// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRealtimeLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRealtimeLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRealtimeLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRealtimeLogDeliveryResponseBody) *DeleteRealtimeLogDeliveryResponse
	GetBody() *DeleteRealtimeLogDeliveryResponseBody
}

type DeleteRealtimeLogDeliveryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRealtimeLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *DeleteRealtimeLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRealtimeLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRealtimeLogDeliveryResponse) GetBody() *DeleteRealtimeLogDeliveryResponseBody {
	return s.Body
}

func (s *DeleteRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *DeleteRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *DeleteRealtimeLogDeliveryResponse) SetStatusCode(v int32) *DeleteRealtimeLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRealtimeLogDeliveryResponse) SetBody(v *DeleteRealtimeLogDeliveryResponseBody) *DeleteRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *DeleteRealtimeLogDeliveryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
