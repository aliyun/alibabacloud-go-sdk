// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayWebsocketEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayWebsocketEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayWebsocketEventsResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayWebsocketEventsResponseBody) *ListGatewayWebsocketEventsResponse
	GetBody() *ListGatewayWebsocketEventsResponseBody
}

type ListGatewayWebsocketEventsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayWebsocketEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayWebsocketEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayWebsocketEventsResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayWebsocketEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayWebsocketEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayWebsocketEventsResponse) GetBody() *ListGatewayWebsocketEventsResponseBody {
	return s.Body
}

func (s *ListGatewayWebsocketEventsResponse) SetHeaders(v map[string]*string) *ListGatewayWebsocketEventsResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayWebsocketEventsResponse) SetStatusCode(v int32) *ListGatewayWebsocketEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayWebsocketEventsResponse) SetBody(v *ListGatewayWebsocketEventsResponseBody) *ListGatewayWebsocketEventsResponse {
	s.Body = v
	return s
}

func (s *ListGatewayWebsocketEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
