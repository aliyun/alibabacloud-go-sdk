// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigDeliveryChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConfigDeliveryChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConfigDeliveryChannelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConfigDeliveryChannelResponseBody) *UpdateConfigDeliveryChannelResponse
	GetBody() *UpdateConfigDeliveryChannelResponseBody
}

type UpdateConfigDeliveryChannelResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConfigDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConfigDeliveryChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *UpdateConfigDeliveryChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConfigDeliveryChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConfigDeliveryChannelResponse) GetBody() *UpdateConfigDeliveryChannelResponseBody {
	return s.Body
}

func (s *UpdateConfigDeliveryChannelResponse) SetHeaders(v map[string]*string) *UpdateConfigDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *UpdateConfigDeliveryChannelResponse) SetStatusCode(v int32) *UpdateConfigDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConfigDeliveryChannelResponse) SetBody(v *UpdateConfigDeliveryChannelResponseBody) *UpdateConfigDeliveryChannelResponse {
	s.Body = v
	return s
}

func (s *UpdateConfigDeliveryChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
