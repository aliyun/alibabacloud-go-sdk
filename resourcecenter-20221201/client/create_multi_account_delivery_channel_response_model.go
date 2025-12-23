// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiAccountDeliveryChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMultiAccountDeliveryChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMultiAccountDeliveryChannelResponse
	GetStatusCode() *int32
	SetBody(v *CreateMultiAccountDeliveryChannelResponseBody) *CreateMultiAccountDeliveryChannelResponse
	GetBody() *CreateMultiAccountDeliveryChannelResponseBody
}

type CreateMultiAccountDeliveryChannelResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMultiAccountDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMultiAccountDeliveryChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiAccountDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *CreateMultiAccountDeliveryChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMultiAccountDeliveryChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMultiAccountDeliveryChannelResponse) GetBody() *CreateMultiAccountDeliveryChannelResponseBody {
	return s.Body
}

func (s *CreateMultiAccountDeliveryChannelResponse) SetHeaders(v map[string]*string) *CreateMultiAccountDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *CreateMultiAccountDeliveryChannelResponse) SetStatusCode(v int32) *CreateMultiAccountDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelResponse) SetBody(v *CreateMultiAccountDeliveryChannelResponseBody) *CreateMultiAccountDeliveryChannelResponse {
	s.Body = v
	return s
}

func (s *CreateMultiAccountDeliveryChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
