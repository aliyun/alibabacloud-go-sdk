// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateConfigDeliveryChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregateConfigDeliveryChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregateConfigDeliveryChannelResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregateConfigDeliveryChannelResponseBody) *GetAggregateConfigDeliveryChannelResponse
	GetBody() *GetAggregateConfigDeliveryChannelResponseBody
}

type GetAggregateConfigDeliveryChannelResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregateConfigDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregateConfigDeliveryChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigDeliveryChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregateConfigDeliveryChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregateConfigDeliveryChannelResponse) GetBody() *GetAggregateConfigDeliveryChannelResponseBody {
	return s.Body
}

func (s *GetAggregateConfigDeliveryChannelResponse) SetHeaders(v map[string]*string) *GetAggregateConfigDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponse) SetStatusCode(v int32) *GetAggregateConfigDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponse) SetBody(v *GetAggregateConfigDeliveryChannelResponseBody) *GetAggregateConfigDeliveryChannelResponse {
	s.Body = v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
