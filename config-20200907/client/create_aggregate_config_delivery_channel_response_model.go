// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregateConfigDeliveryChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAggregateConfigDeliveryChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAggregateConfigDeliveryChannelResponse
	GetStatusCode() *int32
	SetBody(v *CreateAggregateConfigDeliveryChannelResponseBody) *CreateAggregateConfigDeliveryChannelResponse
	GetBody() *CreateAggregateConfigDeliveryChannelResponseBody
}

type CreateAggregateConfigDeliveryChannelResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAggregateConfigDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAggregateConfigDeliveryChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateConfigDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *CreateAggregateConfigDeliveryChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAggregateConfigDeliveryChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAggregateConfigDeliveryChannelResponse) GetBody() *CreateAggregateConfigDeliveryChannelResponseBody {
	return s.Body
}

func (s *CreateAggregateConfigDeliveryChannelResponse) SetHeaders(v map[string]*string) *CreateAggregateConfigDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelResponse) SetStatusCode(v int32) *CreateAggregateConfigDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelResponse) SetBody(v *CreateAggregateConfigDeliveryChannelResponseBody) *CreateAggregateConfigDeliveryChannelResponse {
	s.Body = v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
