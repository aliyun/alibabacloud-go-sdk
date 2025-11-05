// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRealtimeLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRealtimeLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRealtimeLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *ListRealtimeLogDeliveryResponseBody) *ListRealtimeLogDeliveryResponse
	GetBody() *ListRealtimeLogDeliveryResponseBody
}

type ListRealtimeLogDeliveryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRealtimeLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRealtimeLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRealtimeLogDeliveryResponse) GetBody() *ListRealtimeLogDeliveryResponseBody {
	return s.Body
}

func (s *ListRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *ListRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *ListRealtimeLogDeliveryResponse) SetStatusCode(v int32) *ListRealtimeLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRealtimeLogDeliveryResponse) SetBody(v *ListRealtimeLogDeliveryResponseBody) *ListRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *ListRealtimeLogDeliveryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
