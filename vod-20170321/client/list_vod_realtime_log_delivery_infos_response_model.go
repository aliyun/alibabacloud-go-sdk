// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodRealtimeLogDeliveryInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVodRealtimeLogDeliveryInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVodRealtimeLogDeliveryInfosResponse
	GetStatusCode() *int32
	SetBody(v *ListVodRealtimeLogDeliveryInfosResponseBody) *ListVodRealtimeLogDeliveryInfosResponse
	GetBody() *ListVodRealtimeLogDeliveryInfosResponseBody
}

type ListVodRealtimeLogDeliveryInfosResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVodRealtimeLogDeliveryInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVodRealtimeLogDeliveryInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryInfosResponse) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVodRealtimeLogDeliveryInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVodRealtimeLogDeliveryInfosResponse) GetBody() *ListVodRealtimeLogDeliveryInfosResponseBody {
	return s.Body
}

func (s *ListVodRealtimeLogDeliveryInfosResponse) SetHeaders(v map[string]*string) *ListVodRealtimeLogDeliveryInfosResponse {
	s.Headers = v
	return s
}

func (s *ListVodRealtimeLogDeliveryInfosResponse) SetStatusCode(v int32) *ListVodRealtimeLogDeliveryInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryInfosResponse) SetBody(v *ListVodRealtimeLogDeliveryInfosResponseBody) *ListVodRealtimeLogDeliveryInfosResponse {
	s.Body = v
	return s
}

func (s *ListVodRealtimeLogDeliveryInfosResponse) Validate() error {
	return dara.Validate(s)
}
