// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiAccountDeliveryChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMultiAccountDeliveryChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMultiAccountDeliveryChannelResponse
	GetStatusCode() *int32
	SetBody(v *GetMultiAccountDeliveryChannelResponseBody) *GetMultiAccountDeliveryChannelResponse
	GetBody() *GetMultiAccountDeliveryChannelResponseBody
}

type GetMultiAccountDeliveryChannelResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultiAccountDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultiAccountDeliveryChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMultiAccountDeliveryChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMultiAccountDeliveryChannelResponse) GetBody() *GetMultiAccountDeliveryChannelResponseBody {
	return s.Body
}

func (s *GetMultiAccountDeliveryChannelResponse) SetHeaders(v map[string]*string) *GetMultiAccountDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponse) SetStatusCode(v int32) *GetMultiAccountDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponse) SetBody(v *GetMultiAccountDeliveryChannelResponseBody) *GetMultiAccountDeliveryChannelResponse {
	s.Body = v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponse) Validate() error {
	return dara.Validate(s)
}
