// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeliveryChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDeliveryChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDeliveryChannelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDeliveryChannelResponseBody) *DeleteDeliveryChannelResponse
	GetBody() *DeleteDeliveryChannelResponseBody
}

type DeleteDeliveryChannelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDeliveryChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeliveryChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDeliveryChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDeliveryChannelResponse) GetBody() *DeleteDeliveryChannelResponseBody {
	return s.Body
}

func (s *DeleteDeliveryChannelResponse) SetHeaders(v map[string]*string) *DeleteDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeliveryChannelResponse) SetStatusCode(v int32) *DeleteDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeliveryChannelResponse) SetBody(v *DeleteDeliveryChannelResponseBody) *DeleteDeliveryChannelResponse {
	s.Body = v
	return s
}

func (s *DeleteDeliveryChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
