// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateConfigDeliveryChannelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAggregateConfigDeliveryChannelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAggregateConfigDeliveryChannelsResponse
	GetStatusCode() *int32
	SetBody(v *ListAggregateConfigDeliveryChannelsResponseBody) *ListAggregateConfigDeliveryChannelsResponse
	GetBody() *ListAggregateConfigDeliveryChannelsResponseBody
}

type ListAggregateConfigDeliveryChannelsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAggregateConfigDeliveryChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAggregateConfigDeliveryChannelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigDeliveryChannelsResponse) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigDeliveryChannelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAggregateConfigDeliveryChannelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAggregateConfigDeliveryChannelsResponse) GetBody() *ListAggregateConfigDeliveryChannelsResponseBody {
	return s.Body
}

func (s *ListAggregateConfigDeliveryChannelsResponse) SetHeaders(v map[string]*string) *ListAggregateConfigDeliveryChannelsResponse {
	s.Headers = v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponse) SetStatusCode(v int32) *ListAggregateConfigDeliveryChannelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponse) SetBody(v *ListAggregateConfigDeliveryChannelsResponseBody) *ListAggregateConfigDeliveryChannelsResponse {
	s.Body = v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
