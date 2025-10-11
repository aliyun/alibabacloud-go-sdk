// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeliveryChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeliveryChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeliveryChannelResponse
	GetStatusCode() *int32
	SetBody(v *GetDeliveryChannelResponseBody) *GetDeliveryChannelResponse
	GetBody() *GetDeliveryChannelResponseBody
}

type GetDeliveryChannelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeliveryChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeliveryChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeliveryChannelResponse) GetBody() *GetDeliveryChannelResponseBody {
	return s.Body
}

func (s *GetDeliveryChannelResponse) SetHeaders(v map[string]*string) *GetDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *GetDeliveryChannelResponse) SetStatusCode(v int32) *GetDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeliveryChannelResponse) SetBody(v *GetDeliveryChannelResponseBody) *GetDeliveryChannelResponse {
	s.Body = v
	return s
}

func (s *GetDeliveryChannelResponse) Validate() error {
	return dara.Validate(s)
}
