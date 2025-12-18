// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigDeliveryChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConfigDeliveryChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConfigDeliveryChannelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteConfigDeliveryChannelResponseBody) *DeleteConfigDeliveryChannelResponse
	GetBody() *DeleteConfigDeliveryChannelResponseBody
}

type DeleteConfigDeliveryChannelResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConfigDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConfigDeliveryChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *DeleteConfigDeliveryChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConfigDeliveryChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConfigDeliveryChannelResponse) GetBody() *DeleteConfigDeliveryChannelResponseBody {
	return s.Body
}

func (s *DeleteConfigDeliveryChannelResponse) SetHeaders(v map[string]*string) *DeleteConfigDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *DeleteConfigDeliveryChannelResponse) SetStatusCode(v int32) *DeleteConfigDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConfigDeliveryChannelResponse) SetBody(v *DeleteConfigDeliveryChannelResponseBody) *DeleteConfigDeliveryChannelResponse {
	s.Body = v
	return s
}

func (s *DeleteConfigDeliveryChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
