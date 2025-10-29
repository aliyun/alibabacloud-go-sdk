// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRealtimeLogDeliveryInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLiveRealtimeLogDeliveryInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLiveRealtimeLogDeliveryInfosResponse
	GetStatusCode() *int32
	SetBody(v *ListLiveRealtimeLogDeliveryInfosResponseBody) *ListLiveRealtimeLogDeliveryInfosResponse
	GetBody() *ListLiveRealtimeLogDeliveryInfosResponseBody
}

type ListLiveRealtimeLogDeliveryInfosResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLiveRealtimeLogDeliveryInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLiveRealtimeLogDeliveryInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryInfosResponse) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLiveRealtimeLogDeliveryInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLiveRealtimeLogDeliveryInfosResponse) GetBody() *ListLiveRealtimeLogDeliveryInfosResponseBody {
	return s.Body
}

func (s *ListLiveRealtimeLogDeliveryInfosResponse) SetHeaders(v map[string]*string) *ListLiveRealtimeLogDeliveryInfosResponse {
	s.Headers = v
	return s
}

func (s *ListLiveRealtimeLogDeliveryInfosResponse) SetStatusCode(v int32) *ListLiveRealtimeLogDeliveryInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryInfosResponse) SetBody(v *ListLiveRealtimeLogDeliveryInfosResponseBody) *ListLiveRealtimeLogDeliveryInfosResponse {
	s.Body = v
	return s
}

func (s *ListLiveRealtimeLogDeliveryInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
