// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeliveryChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDeliveryChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDeliveryChannelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDeliveryChannelResponseBody) *UpdateDeliveryChannelResponse
	GetBody() *UpdateDeliveryChannelResponseBody
}

type UpdateDeliveryChannelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeliveryChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeliveryChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDeliveryChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDeliveryChannelResponse) GetBody() *UpdateDeliveryChannelResponseBody {
	return s.Body
}

func (s *UpdateDeliveryChannelResponse) SetHeaders(v map[string]*string) *UpdateDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeliveryChannelResponse) SetStatusCode(v int32) *UpdateDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeliveryChannelResponse) SetBody(v *UpdateDeliveryChannelResponseBody) *UpdateDeliveryChannelResponse {
	s.Body = v
	return s
}

func (s *UpdateDeliveryChannelResponse) Validate() error {
	return dara.Validate(s)
}
