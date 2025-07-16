// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRealTimeLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRealTimeLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRealTimeLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *CreateRealTimeLogDeliveryResponseBody) *CreateRealTimeLogDeliveryResponse
	GetBody() *CreateRealTimeLogDeliveryResponseBody
}

type CreateRealTimeLogDeliveryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRealTimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRealTimeLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRealTimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *CreateRealTimeLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRealTimeLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRealTimeLogDeliveryResponse) GetBody() *CreateRealTimeLogDeliveryResponseBody {
	return s.Body
}

func (s *CreateRealTimeLogDeliveryResponse) SetHeaders(v map[string]*string) *CreateRealTimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *CreateRealTimeLogDeliveryResponse) SetStatusCode(v int32) *CreateRealTimeLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRealTimeLogDeliveryResponse) SetBody(v *CreateRealTimeLogDeliveryResponseBody) *CreateRealTimeLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *CreateRealTimeLogDeliveryResponse) Validate() error {
	return dara.Validate(s)
}
