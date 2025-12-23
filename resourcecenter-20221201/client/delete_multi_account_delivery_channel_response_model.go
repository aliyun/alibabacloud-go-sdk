// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultiAccountDeliveryChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMultiAccountDeliveryChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMultiAccountDeliveryChannelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMultiAccountDeliveryChannelResponseBody) *DeleteMultiAccountDeliveryChannelResponse
	GetBody() *DeleteMultiAccountDeliveryChannelResponseBody
}

type DeleteMultiAccountDeliveryChannelResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMultiAccountDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMultiAccountDeliveryChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiAccountDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *DeleteMultiAccountDeliveryChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMultiAccountDeliveryChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMultiAccountDeliveryChannelResponse) GetBody() *DeleteMultiAccountDeliveryChannelResponseBody {
	return s.Body
}

func (s *DeleteMultiAccountDeliveryChannelResponse) SetHeaders(v map[string]*string) *DeleteMultiAccountDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *DeleteMultiAccountDeliveryChannelResponse) SetStatusCode(v int32) *DeleteMultiAccountDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMultiAccountDeliveryChannelResponse) SetBody(v *DeleteMultiAccountDeliveryChannelResponseBody) *DeleteMultiAccountDeliveryChannelResponse {
	s.Body = v
	return s
}

func (s *DeleteMultiAccountDeliveryChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
