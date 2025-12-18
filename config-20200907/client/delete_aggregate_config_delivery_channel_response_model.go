// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAggregateConfigDeliveryChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAggregateConfigDeliveryChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAggregateConfigDeliveryChannelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAggregateConfigDeliveryChannelResponseBody) *DeleteAggregateConfigDeliveryChannelResponse
	GetBody() *DeleteAggregateConfigDeliveryChannelResponseBody
}

type DeleteAggregateConfigDeliveryChannelResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAggregateConfigDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAggregateConfigDeliveryChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregateConfigDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *DeleteAggregateConfigDeliveryChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAggregateConfigDeliveryChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAggregateConfigDeliveryChannelResponse) GetBody() *DeleteAggregateConfigDeliveryChannelResponseBody {
	return s.Body
}

func (s *DeleteAggregateConfigDeliveryChannelResponse) SetHeaders(v map[string]*string) *DeleteAggregateConfigDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *DeleteAggregateConfigDeliveryChannelResponse) SetStatusCode(v int32) *DeleteAggregateConfigDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAggregateConfigDeliveryChannelResponse) SetBody(v *DeleteAggregateConfigDeliveryChannelResponseBody) *DeleteAggregateConfigDeliveryChannelResponse {
	s.Body = v
	return s
}

func (s *DeleteAggregateConfigDeliveryChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
