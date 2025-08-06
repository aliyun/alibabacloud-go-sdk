// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodRealtimeLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVodRealtimeLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVodRealtimeLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVodRealtimeLogDeliveryResponseBody) *DeleteVodRealtimeLogDeliveryResponse
	GetBody() *DeleteVodRealtimeLogDeliveryResponseBody
}

type DeleteVodRealtimeLogDeliveryResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVodRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVodRealtimeLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *DeleteVodRealtimeLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVodRealtimeLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVodRealtimeLogDeliveryResponse) GetBody() *DeleteVodRealtimeLogDeliveryResponseBody {
	return s.Body
}

func (s *DeleteVodRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *DeleteVodRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *DeleteVodRealtimeLogDeliveryResponse) SetStatusCode(v int32) *DeleteVodRealtimeLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVodRealtimeLogDeliveryResponse) SetBody(v *DeleteVodRealtimeLogDeliveryResponseBody) *DeleteVodRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *DeleteVodRealtimeLogDeliveryResponse) Validate() error {
	return dara.Validate(s)
}
