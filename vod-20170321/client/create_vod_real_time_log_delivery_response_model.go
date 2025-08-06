// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVodRealTimeLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVodRealTimeLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVodRealTimeLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *CreateVodRealTimeLogDeliveryResponseBody) *CreateVodRealTimeLogDeliveryResponse
	GetBody() *CreateVodRealTimeLogDeliveryResponseBody
}

type CreateVodRealTimeLogDeliveryResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVodRealTimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVodRealTimeLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVodRealTimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *CreateVodRealTimeLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVodRealTimeLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVodRealTimeLogDeliveryResponse) GetBody() *CreateVodRealTimeLogDeliveryResponseBody {
	return s.Body
}

func (s *CreateVodRealTimeLogDeliveryResponse) SetHeaders(v map[string]*string) *CreateVodRealTimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *CreateVodRealTimeLogDeliveryResponse) SetStatusCode(v int32) *CreateVodRealTimeLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVodRealTimeLogDeliveryResponse) SetBody(v *CreateVodRealTimeLogDeliveryResponseBody) *CreateVodRealTimeLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *CreateVodRealTimeLogDeliveryResponse) Validate() error {
	return dara.Validate(s)
}
