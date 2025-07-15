// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveRealTimeLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLiveRealTimeLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLiveRealTimeLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *CreateLiveRealTimeLogDeliveryResponseBody) *CreateLiveRealTimeLogDeliveryResponse
	GetBody() *CreateLiveRealTimeLogDeliveryResponseBody
}

type CreateLiveRealTimeLogDeliveryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLiveRealTimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLiveRealTimeLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveRealTimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveRealTimeLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLiveRealTimeLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLiveRealTimeLogDeliveryResponse) GetBody() *CreateLiveRealTimeLogDeliveryResponseBody {
	return s.Body
}

func (s *CreateLiveRealTimeLogDeliveryResponse) SetHeaders(v map[string]*string) *CreateLiveRealTimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *CreateLiveRealTimeLogDeliveryResponse) SetStatusCode(v int32) *CreateLiveRealTimeLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLiveRealTimeLogDeliveryResponse) SetBody(v *CreateLiveRealTimeLogDeliveryResponseBody) *CreateLiveRealTimeLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *CreateLiveRealTimeLogDeliveryResponse) Validate() error {
	return dara.Validate(s)
}
