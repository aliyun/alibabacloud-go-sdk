// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRealtimeLogDeliveryInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRealtimeLogDeliveryInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRealtimeLogDeliveryInfosResponse
	GetStatusCode() *int32
	SetBody(v *ListRealtimeLogDeliveryInfosResponseBody) *ListRealtimeLogDeliveryInfosResponse
	GetBody() *ListRealtimeLogDeliveryInfosResponseBody
}

type ListRealtimeLogDeliveryInfosResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRealtimeLogDeliveryInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRealtimeLogDeliveryInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeLogDeliveryInfosResponse) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRealtimeLogDeliveryInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRealtimeLogDeliveryInfosResponse) GetBody() *ListRealtimeLogDeliveryInfosResponseBody {
	return s.Body
}

func (s *ListRealtimeLogDeliveryInfosResponse) SetHeaders(v map[string]*string) *ListRealtimeLogDeliveryInfosResponse {
	s.Headers = v
	return s
}

func (s *ListRealtimeLogDeliveryInfosResponse) SetStatusCode(v int32) *ListRealtimeLogDeliveryInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRealtimeLogDeliveryInfosResponse) SetBody(v *ListRealtimeLogDeliveryInfosResponseBody) *ListRealtimeLogDeliveryInfosResponse {
	s.Body = v
	return s
}

func (s *ListRealtimeLogDeliveryInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
