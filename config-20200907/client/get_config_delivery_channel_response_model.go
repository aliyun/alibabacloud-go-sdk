// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigDeliveryChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConfigDeliveryChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConfigDeliveryChannelResponse
	GetStatusCode() *int32
	SetBody(v *GetConfigDeliveryChannelResponseBody) *GetConfigDeliveryChannelResponse
	GetBody() *GetConfigDeliveryChannelResponseBody
}

type GetConfigDeliveryChannelResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConfigDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConfigDeliveryChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConfigDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *GetConfigDeliveryChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConfigDeliveryChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConfigDeliveryChannelResponse) GetBody() *GetConfigDeliveryChannelResponseBody {
	return s.Body
}

func (s *GetConfigDeliveryChannelResponse) SetHeaders(v map[string]*string) *GetConfigDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *GetConfigDeliveryChannelResponse) SetStatusCode(v int32) *GetConfigDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConfigDeliveryChannelResponse) SetBody(v *GetConfigDeliveryChannelResponseBody) *GetConfigDeliveryChannelResponse {
	s.Body = v
	return s
}

func (s *GetConfigDeliveryChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
