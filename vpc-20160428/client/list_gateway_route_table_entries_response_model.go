// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayRouteTableEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayRouteTableEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayRouteTableEntriesResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayRouteTableEntriesResponseBody) *ListGatewayRouteTableEntriesResponse
	GetBody() *ListGatewayRouteTableEntriesResponseBody
}

type ListGatewayRouteTableEntriesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayRouteTableEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayRouteTableEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteTableEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteTableEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayRouteTableEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayRouteTableEntriesResponse) GetBody() *ListGatewayRouteTableEntriesResponseBody {
	return s.Body
}

func (s *ListGatewayRouteTableEntriesResponse) SetHeaders(v map[string]*string) *ListGatewayRouteTableEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayRouteTableEntriesResponse) SetStatusCode(v int32) *ListGatewayRouteTableEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayRouteTableEntriesResponse) SetBody(v *ListGatewayRouteTableEntriesResponseBody) *ListGatewayRouteTableEntriesResponse {
	s.Body = v
	return s
}

func (s *ListGatewayRouteTableEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
