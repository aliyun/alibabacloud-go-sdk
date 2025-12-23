// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiAccountDeliveryChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMultiAccountDeliveryChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMultiAccountDeliveryChannelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMultiAccountDeliveryChannelResponseBody) *UpdateMultiAccountDeliveryChannelResponse
	GetBody() *UpdateMultiAccountDeliveryChannelResponseBody
}

type UpdateMultiAccountDeliveryChannelResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMultiAccountDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMultiAccountDeliveryChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiAccountDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *UpdateMultiAccountDeliveryChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMultiAccountDeliveryChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMultiAccountDeliveryChannelResponse) GetBody() *UpdateMultiAccountDeliveryChannelResponseBody {
	return s.Body
}

func (s *UpdateMultiAccountDeliveryChannelResponse) SetHeaders(v map[string]*string) *UpdateMultiAccountDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelResponse) SetStatusCode(v int32) *UpdateMultiAccountDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelResponse) SetBody(v *UpdateMultiAccountDeliveryChannelResponseBody) *UpdateMultiAccountDeliveryChannelResponse {
	s.Body = v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
