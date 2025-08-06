// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodRealtimeLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVodRealtimeLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVodRealtimeLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *ListVodRealtimeLogDeliveryResponseBody) *ListVodRealtimeLogDeliveryResponse
	GetBody() *ListVodRealtimeLogDeliveryResponseBody
}

type ListVodRealtimeLogDeliveryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVodRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVodRealtimeLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVodRealtimeLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVodRealtimeLogDeliveryResponse) GetBody() *ListVodRealtimeLogDeliveryResponseBody {
	return s.Body
}

func (s *ListVodRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *ListVodRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *ListVodRealtimeLogDeliveryResponse) SetStatusCode(v int32) *ListVodRealtimeLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryResponse) SetBody(v *ListVodRealtimeLogDeliveryResponseBody) *ListVodRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *ListVodRealtimeLogDeliveryResponse) Validate() error {
	return dara.Validate(s)
}
