// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggregateConfigDeliveryChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAggregateConfigDeliveryChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAggregateConfigDeliveryChannelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAggregateConfigDeliveryChannelResponseBody) *UpdateAggregateConfigDeliveryChannelResponse
	GetBody() *UpdateAggregateConfigDeliveryChannelResponseBody
}

type UpdateAggregateConfigDeliveryChannelResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAggregateConfigDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAggregateConfigDeliveryChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateConfigDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *UpdateAggregateConfigDeliveryChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAggregateConfigDeliveryChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAggregateConfigDeliveryChannelResponse) GetBody() *UpdateAggregateConfigDeliveryChannelResponseBody {
	return s.Body
}

func (s *UpdateAggregateConfigDeliveryChannelResponse) SetHeaders(v map[string]*string) *UpdateAggregateConfigDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelResponse) SetStatusCode(v int32) *UpdateAggregateConfigDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelResponse) SetBody(v *UpdateAggregateConfigDeliveryChannelResponseBody) *UpdateAggregateConfigDeliveryChannelResponse {
	s.Body = v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
