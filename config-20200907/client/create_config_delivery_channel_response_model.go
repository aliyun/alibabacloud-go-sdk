// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigDeliveryChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateConfigDeliveryChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateConfigDeliveryChannelResponse
	GetStatusCode() *int32
	SetBody(v *CreateConfigDeliveryChannelResponseBody) *CreateConfigDeliveryChannelResponse
	GetBody() *CreateConfigDeliveryChannelResponseBody
}

type CreateConfigDeliveryChannelResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConfigDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConfigDeliveryChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *CreateConfigDeliveryChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateConfigDeliveryChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateConfigDeliveryChannelResponse) GetBody() *CreateConfigDeliveryChannelResponseBody {
	return s.Body
}

func (s *CreateConfigDeliveryChannelResponse) SetHeaders(v map[string]*string) *CreateConfigDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *CreateConfigDeliveryChannelResponse) SetStatusCode(v int32) *CreateConfigDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConfigDeliveryChannelResponse) SetBody(v *CreateConfigDeliveryChannelResponseBody) *CreateConfigDeliveryChannelResponse {
	s.Body = v
	return s
}

func (s *CreateConfigDeliveryChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
