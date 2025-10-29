// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRealtimeLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLiveRealtimeLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLiveRealtimeLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *ListLiveRealtimeLogDeliveryResponseBody) *ListLiveRealtimeLogDeliveryResponse
	GetBody() *ListLiveRealtimeLogDeliveryResponseBody
}

type ListLiveRealtimeLogDeliveryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLiveRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLiveRealtimeLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLiveRealtimeLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLiveRealtimeLogDeliveryResponse) GetBody() *ListLiveRealtimeLogDeliveryResponseBody {
	return s.Body
}

func (s *ListLiveRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *ListLiveRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *ListLiveRealtimeLogDeliveryResponse) SetStatusCode(v int32) *ListLiveRealtimeLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryResponse) SetBody(v *ListLiveRealtimeLogDeliveryResponseBody) *ListLiveRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *ListLiveRealtimeLogDeliveryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
